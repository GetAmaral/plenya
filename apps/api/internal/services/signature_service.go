package services

import (
	"bytes"
	"crypto"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/digitorus/pdf"
	"github.com/digitorus/pdfsign/sign"
	"github.com/digitorus/pdfsign/verify"
	"github.com/google/uuid"
)

var (
	ErrSignatureFailed = errors.New("failed to sign PDF")
)

// SignatureOptions — config de assinatura (carimbo de tempo RFC 3161 opcional).
type SignatureOptions struct {
	TSAURL      string // ACT credenciada (vazio = PAdES básico AD-RB, sem carimbo de tempo)
	TSAUsername string
	TSAPassword string
}

type SignatureService struct {
	certService *CertificateService
	opts        SignatureOptions
}

func NewSignatureService(certService *CertificateService, opts SignatureOptions) *SignatureService {
	return &SignatureService{
		certService: certService,
		opts:        opts,
	}
}

// SignPrescriptionPDF assina PDF de prescrição com certificado do médico.
// Mantido como wrapper p/ não quebrar prescription_pdf_service / lab_request_pdf_service.
// Retorna: PDF assinado, hash SHA-256, erro.
func (s *SignatureService) SignPrescriptionPDF(
	pdfBytes []byte,
	doctorID uuid.UUID,
) (signedPDF []byte, signatureHash string, err error) {
	return s.signPDF(pdfBytes, doctorID,
		"Assinatura Digital de Prescrição Médica",
		"Plenya EMR - Sistema de Prescrição Digital")
}

// SignDocumentPDF assina qualquer PDF clínico (atestado/declaração/laudo/relatório) com o
// certificado ICP-Brasil do médico. `reason` e `name` parametrizam os metadados PAdES.
func (s *SignatureService) SignDocumentPDF(
	pdfBytes []byte,
	doctorID uuid.UUID,
	reason string,
	name string,
) (signedPDF []byte, signatureHash string, err error) {
	if reason == "" {
		reason = "Assinatura Digital de Documento Médico"
	}
	if name == "" {
		name = "Plenya EMR - Documentos Clínicos"
	}
	return s.signPDF(pdfBytes, doctorID, reason, name)
}

// signPDF é o assinador PAdES/ICP-Brasil genérico (núcleo compartilhado).
func (s *SignatureService) signPDF(
	pdfBytes []byte,
	doctorID uuid.UUID,
	reason string,
	name string,
) (signedPDF []byte, signatureHash string, err error) {
	// 1. Obter o assinador do médico (A1 local ou e-CPF em nuvem, transparente).
	cert, signer, _, err := s.certService.GetSigner(doctorID)
	if err != nil {
		return nil, "", err
	}

	// 2. Preparar dados de assinatura PAdES
	signData := sign.SignData{
		Signature: sign.SignDataSignature{
			Info: sign.SignDataSignatureInfo{
				Name:        name,
				Location:    "Brasil",
				Reason:      reason,
				ContactInfo: "https://plenya.com.br",
			},
			CertType:   sign.CertificationSignature,
			DocMDPPerm: sign.AllowFillingExistingFormFieldsAndSignaturesPerms,
		},
		DigestAlgorithm: crypto.SHA256,
		Certificate:     cert,
		Signer:          signer,
	}

	// 2b. Carimbo de tempo (RFC 3161 / PAdES-T) quando uma ACT credenciada está configurada.
	if s.opts.TSAURL != "" {
		signData.TSA = sign.TSA{
			URL:      s.opts.TSAURL,
			Username: s.opts.TSAUsername,
			Password: s.opts.TSAPassword,
		}
	}

	// 4. Criar readers e writers
	inputReader := bytes.NewReader(pdfBytes)
	var outputBuffer bytes.Buffer

	// 5. Criar PDF reader
	pdfReader, err := pdf.NewReader(inputReader, int64(len(pdfBytes)))
	if err != nil {
		return nil, "", ErrSignatureFailed
	}

	// 6. Assinar PDF com PAdES
	err = sign.Sign(inputReader, &outputBuffer, pdfReader, int64(len(pdfBytes)), signData)
	if err != nil {
		return nil, "", ErrSignatureFailed
	}

	signedPDF = outputBuffer.Bytes()

	// 7. Calcular hash SHA-256 do PDF assinado
	hash := sha256.Sum256(signedPDF)
	signatureHash = hex.EncodeToString(hash[:])

	return signedPDF, signatureHash, nil
}

// SignatureVerification — resultado da verificação criptográfica da assinatura PAdES.
type SignatureVerification struct {
	Signed       bool         `json:"signed"`       // há ao menos uma assinatura no PDF
	Valid        bool         `json:"valid"`        // todas as assinaturas íntegras (hash confere)
	TrustedAll   bool         `json:"trustedAll"`   // todas com emissor confiável (cadeia até raiz)
	HasTimestamp bool         `json:"hasTimestamp"` // ao menos uma com carimbo de tempo
	SignerCount  int          `json:"signerCount"`
	Signers      []SignerInfo `json:"signers"`
	Error        string       `json:"error,omitempty"`
}

// SignerInfo — dados de um signatário extraídos da assinatura.
type SignerInfo struct {
	Name           string `json:"name,omitempty"`
	ValidSignature bool   `json:"validSignature"`
	TrustedIssuer  bool   `json:"trustedIssuer"`
	Revoked        bool   `json:"revoked"`
	HasTimestamp   bool   `json:"hasTimestamp"`
}

// VerifySignature valida a assinatura PAdES do PDF de verdade (integridade criptográfica +
// dados do signatário), via github.com/digitorus/pdfsign/verify. Substitui o antigo stub.
//
// Nota: TrustedIssuer depende das raízes confiáveis do verificador; certificados ICP-Brasil
// podem aparecer como não-confiáveis aqui mesmo sendo válidos. A validação oficial de
// confiança/revogação para terceiros é o validador do ITI (validar.iti.gov.br). Aqui o ganho
// é confirmar a INTEGRIDADE (Valid), não só re-hashear o arquivo.
func (s *SignatureService) VerifySignature(pdfBytes []byte) (*SignatureVerification, error) {
	resp, err := verify.Verify(bytes.NewReader(pdfBytes), int64(len(pdfBytes)))
	if err != nil {
		// PDF sem assinatura ou ilegível — não é erro de servidor; é resultado "não assinado".
		return &SignatureVerification{Signed: false, Valid: false, Error: err.Error()}, nil
	}

	out := &SignatureVerification{
		Signed:      len(resp.Signers) > 0,
		Valid:       len(resp.Signers) > 0,
		TrustedAll:  len(resp.Signers) > 0,
		SignerCount: len(resp.Signers),
	}
	if resp.Error != "" {
		out.Error = resp.Error
	}
	for _, sg := range resp.Signers {
		hasTS := sg.TimeStamp != nil
		if hasTS {
			out.HasTimestamp = true
		}
		out.Signers = append(out.Signers, SignerInfo{
			Name:           sg.Name,
			ValidSignature: sg.ValidSignature,
			TrustedIssuer:  sg.TrustedIssuer,
			Revoked:        sg.RevokedCertificate,
			HasTimestamp:   hasTS,
		})
		if !sg.ValidSignature {
			out.Valid = false
		}
		if !sg.TrustedIssuer {
			out.TrustedAll = false
		}
	}
	return out, nil
}
