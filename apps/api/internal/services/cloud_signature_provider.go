package services

import (
	"bytes"
	"crypto"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// ============================================================================
// Assinatura ICP-Brasil em NUVEM (e-CPF em nuvem / PSC/broker)
//
// A chave privada do médico fica num HSM certificado do provedor (VIDaaS/BirdID/SafeID/
// SerproID), exposto por um broker de API REST (IntegraICP/Certillion — o mesmo motor que
// a Prescrição Eletrônica do CFM usa por baixo). O backend NUNCA detém a chave: dispara a
// assinatura por API e o titular confirma a sessão (push/biometria/OTP no app do provedor).
// Não existe assinatura ICP-Brasil 100% "headless" — sempre há um ato de vontade do titular;
// por isso o fluxo é: StartAuthorization (titular confirma) -> credencial com janela de tempo
// -> SignHash em lote dentro da janela.
//
// O contrato REST abaixo segue o padrão "assinar hash -> assinatura crua" (IntegraICP v3 /
// Certillion), que casa com crypto.Signer e com a pipeline PAdES já existente (sign.Sign).
// Gated off por default (ICP_CLOUD_ENABLED=false) até haver contrato/credencial do PSC —
// mesmo padrão do SNCRProductionProvider. Os caminhos/payloads exatos devem ser confirmados
// na documentação do PSC escolhido antes de ligar em produção.
// ============================================================================

var (
	ErrCloudSigningDisabled = errors.New("assinatura em nuvem desabilitada (ICP_CLOUD_ENABLED=false)")
	ErrCloudAuthExpired     = errors.New("credencial de assinatura em nuvem ausente ou expirada — reautorize")
	ErrCloudProvider        = errors.New("erro no provedor de assinatura em nuvem")
)

// CloudSignatureConfig — espelho dos campos cloud de config.SignatureConfig (mapeado em main.go,
// para não acoplar o pacote services ao pacote config — mesmo padrão de SNCRConfig).
type CloudSignatureConfig struct {
	Enabled  bool
	Provider string
	BaseURL  string
	APIKey   string
}

// CloudSignatureProvider abstrai um PSC/broker ICP-Brasil de assinatura em nuvem.
type CloudSignatureProvider interface {
	// Name identifica o provedor ativo.
	Name() string
	// Enabled indica se a assinatura em nuvem está operacional (config + base URL).
	Enabled() bool
	// StartAuthorization inicia a sessão de autorização do titular (push/OTP no app do PSC).
	// credentialRef é o id da credencial/chave em nuvem do médico. Retorna o token autorizado
	// e sua expiração — dentro dessa janela o backend pode assinar em lote.
	StartAuthorization(credentialRef string) (token string, expiresAt time.Time, err error)
	// SignHash assina o digest (hash já aplicado) com a chave em nuvem do titular.
	// hashAlg identifica o algoritmo do digest (ex.: crypto.SHA256). Retorna a assinatura crua
	// (PKCS#1 v1.5 sobre o DigestInfo), pronta para embutir no PAdES.
	SignHash(credentialRef, credentialToken string, digest []byte, hashAlg crypto.Hash) (signature []byte, err error)
}

// NewCloudSignatureProvider escolhe a implementação conforme a config.
func NewCloudSignatureProvider(cfg CloudSignatureConfig) CloudSignatureProvider {
	if !cfg.Enabled || cfg.BaseURL == "" {
		return &disabledCloudProvider{}
	}
	// IntegraICP e Certillion expõem o mesmo padrão REST "assinar hash"; um provider genérico
	// cobre ambos via base URL. Diferenças finas de payload ficam confinadas aqui.
	return newRESTCloudProvider(cfg.Provider, cfg.BaseURL, cfg.APIKey)
}

// ----------------------------------------------------------------------------
// Provider desabilitado (default)
// ----------------------------------------------------------------------------

type disabledCloudProvider struct{}

func (p *disabledCloudProvider) Name() string  { return "disabled" }
func (p *disabledCloudProvider) Enabled() bool { return false }
func (p *disabledCloudProvider) StartAuthorization(string) (string, time.Time, error) {
	return "", time.Time{}, ErrCloudSigningDisabled
}
func (p *disabledCloudProvider) SignHash(string, string, []byte, crypto.Hash) ([]byte, error) {
	return nil, ErrCloudSigningDisabled
}

// ----------------------------------------------------------------------------
// Provider REST genérico (IntegraICP / Certillion)
// ----------------------------------------------------------------------------

type restCloudProvider struct {
	name    string
	baseURL string
	apiKey  string
	http    *http.Client
}

func newRESTCloudProvider(name, baseURL, apiKey string) *restCloudProvider {
	if name == "" {
		name = "integraicp"
	}
	return &restCloudProvider{
		name:    name,
		baseURL: trimTrailingSlash(baseURL),
		apiKey:  apiKey,
		http:    &http.Client{Timeout: 20 * time.Second},
	}
}

func (p *restCloudProvider) Name() string  { return p.name }
func (p *restCloudProvider) Enabled() bool { return p.baseURL != "" }

type cloudAuthResponse struct {
	Token     string `json:"token"`
	ExpiresAt string `json:"expiresAt"` // RFC3339
	ExpiresIn int    `json:"expiresIn"` // segundos (fallback se ExpiresAt vazio)
}

func (p *restCloudProvider) StartAuthorization(credentialRef string) (string, time.Time, error) {
	// CONTRATO (confirmar na doc do PSC): POST {base}/credentials/{ref}/authorize
	// O PSC dispara push/OTP no app do titular; a resposta traz o token de credencial e a janela.
	body := map[string]any{"credentialRef": credentialRef}
	var out cloudAuthResponse
	if err := p.do("POST", fmt.Sprintf("/credentials/%s/authorize", credentialRef), body, &out); err != nil {
		return "", time.Time{}, err
	}
	if out.Token == "" {
		return "", time.Time{}, fmt.Errorf("%w: resposta sem token", ErrCloudProvider)
	}
	exp := time.Now().Add(time.Duration(out.ExpiresIn) * time.Second)
	if out.ExpiresAt != "" {
		if t, err := time.Parse(time.RFC3339, out.ExpiresAt); err == nil {
			exp = t
		}
	}
	if out.ExpiresIn == 0 && out.ExpiresAt == "" {
		// Janela padrão conservadora quando o PSC não informa (sessão por consulta).
		exp = time.Now().Add(30 * time.Minute)
	}
	return out.Token, exp, nil
}

type cloudSignRequest struct {
	CredentialRef string   `json:"credentialRef"`
	Token         string   `json:"token"`
	HashAlgorithm string   `json:"hashAlgorithm"`
	Hashes        []string `json:"hashes"` // base64 dos digests
}

type cloudSignResponse struct {
	Signatures []string `json:"signatures"` // base64 das assinaturas (mesma ordem)
}

func (p *restCloudProvider) SignHash(credentialRef, credentialToken string, digest []byte, hashAlg crypto.Hash) ([]byte, error) {
	if credentialToken == "" {
		return nil, ErrCloudAuthExpired
	}
	// CONTRATO (confirmar na doc do PSC): POST {base}/signatures
	req := cloudSignRequest{
		CredentialRef: credentialRef,
		Token:         credentialToken,
		HashAlgorithm: hashAlgName(hashAlg),
		Hashes:        []string{base64.StdEncoding.EncodeToString(digest)},
	}
	var out cloudSignResponse
	if err := p.do("POST", "/signatures", req, &out); err != nil {
		return nil, err
	}
	if len(out.Signatures) == 0 {
		return nil, fmt.Errorf("%w: resposta sem assinatura", ErrCloudProvider)
	}
	sig, err := base64.StdEncoding.DecodeString(out.Signatures[0])
	if err != nil {
		return nil, fmt.Errorf("%w: assinatura base64 inválida: %v", ErrCloudProvider, err)
	}
	return sig, nil
}

func (p *restCloudProvider) do(method, path string, body any, out any) error {
	var rdr io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return err
		}
		rdr = bytes.NewReader(b)
	}
	httpReq, err := http.NewRequest(method, p.baseURL+path, rdr)
	if err != nil {
		return err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json")
	if p.apiKey != "" {
		httpReq.Header.Set("Authorization", "Bearer "+p.apiKey)
	}
	resp, err := p.http.Do(httpReq)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrCloudProvider, err)
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("%w: status %d: %s", ErrCloudProvider, resp.StatusCode, string(raw))
	}
	if out != nil {
		if err := json.Unmarshal(raw, out); err != nil {
			return fmt.Errorf("%w: resposta inválida: %v", ErrCloudProvider, err)
		}
	}
	return nil
}

// hashAlgName mapeia crypto.Hash para o rótulo esperado pelo PSC.
func hashAlgName(h crypto.Hash) string {
	switch h {
	case crypto.SHA256:
		return "SHA256"
	case crypto.SHA384:
		return "SHA384"
	case crypto.SHA512:
		return "SHA512"
	default:
		return "SHA256"
	}
}

func trimTrailingSlash(s string) string {
	for len(s) > 0 && s[len(s)-1] == '/' {
		s = s[:len(s)-1]
	}
	return s
}

// ----------------------------------------------------------------------------
// cloudRemoteSigner — crypto.Signer cuja operação de chave privada acontece no PSC.
// Permite reusar a pipeline PAdES (sign.Sign) sem que a chave saia do HSM do provedor.
// ----------------------------------------------------------------------------

type cloudRemoteSigner struct {
	provider        CloudSignatureProvider
	credentialRef   string
	credentialToken string
	pub             crypto.PublicKey
}

func (s *cloudRemoteSigner) Public() crypto.PublicKey { return s.pub }

func (s *cloudRemoteSigner) Sign(_ io.Reader, digest []byte, opts crypto.SignerOpts) ([]byte, error) {
	return s.provider.SignHash(s.credentialRef, s.credentialToken, digest, opts.HashFunc())
}
