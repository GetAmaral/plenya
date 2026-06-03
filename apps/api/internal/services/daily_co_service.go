// Package services — DailyCoService cria/deleta salas Daily.co pra
// teleconsulta embedada no EMR (Calendar V1, Bloco C).
//
// API: https://docs.daily.co/reference/rest-api
//
// Decisões:
//
//  1. HIGH H9 (resolvido): salas são criadas com privacy="private" e cada
//     participante recebe um meeting_token curto via CreateMeetingToken.
//     Médico = is_owner=true + screenshare; paciente = is_owner=false sem
//     screenshare. URL crua não funciona — Daily exige o token no querystring.
//
//  2. exp (expiration) no momento da criação evita salas órfãs ocupando quota
//     se appointment for cancelado e DeleteRoom falhar.
//
//  3. Quando APIKey vazio, retorna ErrDailyNotConfigured — caller (appointment
//     service) decide se cria appointment sem URL ou bloqueia o tipo telemedicine.
//
//  4. Não usamos SDK (não existe oficial Go) — REST API é trivial.
package services

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/plenya/api/internal/config"
)

const (
	dailyAPIBaseURL        = "https://api.daily.co/v1"
	dailyHTTPClientTimeout = 15 * time.Second

	// Transcrição nativa do Daily (Deepgram). Configurável aqui num único lugar.
	// nova-3 tem PT-BR GA com o maior ganho de WER da rodada; se o backend de
	// transcrição do Daily rejeitar nova-3, trocar por "nova-2-general".
	dailyTranscriptionModel    = "nova-3-general"
	dailyTranscriptionLanguage = "pt-BR"
)

// Erros públicos.
var (
	// ErrDailyNotConfigured = APIKey vazio. Caller decide se cria appointment sem sala.
	ErrDailyNotConfigured = errors.New("daily.co not configured (DAILY_CO_API_KEY missing)")

	// ErrDailyConflict = nome de sala já existe (raro — usamos prefixo + UUID).
	ErrDailyConflict = errors.New("daily.co room name conflict")
)

type DailyCoService struct {
	cfg    *config.Config
	client *http.Client
}

func NewDailyCoService(cfg *config.Config) *DailyCoService {
	return &DailyCoService{
		cfg: cfg,
		client: &http.Client{
			Timeout: dailyHTTPClientTimeout,
		},
	}
}

// DailyRoom — sala criada no Daily.co.
type DailyRoom struct {
	Name string // ex: "plenya-abc123"
	URL  string // ex: "https://plenya.daily.co/plenya-abc123"
}

// IsConfigured retorna true se APIKey está setado.
func (s *DailyCoService) IsConfigured() bool {
	return s.cfg.DailyCo.APIKey != ""
}

// CreateRoom cria sala no Daily.co.
//
//   - namePrefix: prefixo legível (ex: appointment ID curto). Sufixo UUID
//     curto é adicionado pra unicidade.
//   - expiresAt: quando sala auto-expira. Recomendado: scheduled_at + 4h.
//
// Properties:
//   - exp: timestamp UNIX de expiração (sala deletada automaticamente)
//   - eject_at_room_exp: usuários são removidos quando expira
//   - enable_screenshare/chat: sempre true (sem custo extra no free tier)
func (s *DailyCoService) CreateRoom(ctx context.Context, namePrefix string, expiresAt time.Time) (*DailyRoom, error) {
	if !s.IsConfigured() {
		return nil, ErrDailyNotConfigured
	}

	// Sanitiza prefix: lowercase, somente [a-z0-9-]. Daily exige.
	cleanPrefix := sanitizeDailyName(namePrefix)
	if cleanPrefix == "" {
		cleanPrefix = "plenya"
	}
	// Sufixo curto random (8 chars) pra evitar conflito sem revelar UUIDs longos.
	suffix := randHex(4) // 8 chars
	name := fmt.Sprintf("%s-%s", cleanPrefix, suffix)
	if len(name) > 60 {
		// Daily permite até 64 chars mas damos folga.
		name = name[:60]
	}

	// HIGH H9 — endurecimento total:
	// • privacy="private": sala não é acessível por URL crua. Cada
	//   participante precisa de um meeting_token (gerado via
	//   CreateMeetingToken). Token carrega is_owner, user_name, exp e
	//   permissões granulares (enable_screenshare).
	// • enable_screenshare=true: liberado para todos (médico e paciente). O
	//   paciente também pode mostrar a tela (ex: exibir um exame na tela dele).
	// • enable_chat=true: chat textual é útil pra trocar links/anotações
	//   durante a consulta.
	// Gravação + transcrição (follow-up telemed): a sala é CAPAZ de gravar/transcrever,
	// mas só INICIA de fato quando o token do owner (médico) traz start_cloud_recording
	// + auto_start_transcription — e isso só acontece com consentimento registrado
	// (ver GetTelemedJoinURL). Aqui apenas habilitamos a capacidade na sala:
	//   • enable_recording="cloud": gravação composta MP4 no storage gerenciado do Daily.
	//   • enable_transcription_storage=true: salva o WebVTT da transcrição (default é só
	//     caption ao vivo, sem persistir).
	//   • auto_transcription_settings: opções usadas quando auto_start_transcription liga
	//     — pt-BR, Deepgram nova-3, pontuação e diarização (rotula médico×paciente).
	body := map[string]any{
		"name":    name,
		"privacy": "private",
		"properties": map[string]any{
			"exp":                          expiresAt.UTC().Unix(),
			"eject_at_room_exp":            true,
			"enable_screenshare":           true,
			"enable_chat":                  true,
			"enable_recording":             "cloud",
			"enable_transcription_storage": true,
			"auto_transcription_settings": map[string]any{
				"language":  dailyTranscriptionLanguage,
				"model":     dailyTranscriptionModel,
				"punctuate": true,
				"extra": map[string]any{
					"diarize":      true,
					"smart_format": true,
				},
			},
		},
	}

	resp, err := s.do(ctx, http.MethodPost, dailyAPIBaseURL+"/rooms", body)
	if err != nil {
		return nil, fmt.Errorf("daily create room http: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusConflict {
		return nil, ErrDailyConflict
	}
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, errors.New("daily.co: invalid api key (401)")
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		errBody, _ := io.ReadAll(resp.Body)
		log.Printf("⚠️  Daily create room: status=%d body=%s", resp.StatusCode, string(errBody))
		return nil, fmt.Errorf("daily create room: status %d", resp.StatusCode)
	}

	var parsed struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return nil, fmt.Errorf("daily create room decode: %w", err)
	}
	return &DailyRoom{Name: parsed.Name, URL: parsed.URL}, nil
}

// MeetingTokenParams — parâmetros pra gerar um meeting_token Daily.co.
//
// Os campos viram propriedades do token (POST /meeting-tokens).
//
//   - RoomName: nome da sala que o token autoriza (escopo).
//   - UserName: display name que aparece no Daily prebuilt UI.
//   - IsOwner: true só pro médico — owner pode ejetar participantes,
//     iniciar gravação, etc.
//   - ExpiresAt: token deixa de valer após esse instante (Unix segundos).
//     Usar janela curta (closesAt do lobby) — token vazado fica inútil
//     rapidamente.
//   - EnableScreenshare: liberado pra médico e paciente (ambos podem mostrar a tela).
type MeetingTokenParams struct {
	RoomName          string
	UserName          string
	IsOwner           bool
	ExpiresAt         time.Time
	EnableScreenshare bool

	// StartCloudRecording — auto-inicia a gravação cloud quando este participante
	// entra. Só no token do médico (owner) e só com consentimento de telemedicina
	// registrado. Exige a sala (ou o token) com enable_recording="cloud".
	StartCloudRecording bool
	// AutoStartTranscription — auto-inicia a transcrição (Deepgram, settings da sala)
	// quando o OWNER entra. Idem: só com consentimento.
	AutoStartTranscription bool
	// DisableRecordingUI — esconde o botão de gravar do Prebuilt pra este
	// participante. Usado no token do paciente (não deve ligar/desligar gravação).
	DisableRecordingUI bool
}

// CreateMeetingToken gera um meeting_token escopado a uma sala + participante.
//
// Retorna o token cru (string opaca). Caller monta a URL final com
// BuildJoinURL.
//
// Erros tratados igual ao CreateRoom (401 → invalid api key; status != 2xx).
func (s *DailyCoService) CreateMeetingToken(ctx context.Context, p MeetingTokenParams) (string, error) {
	if !s.IsConfigured() {
		return "", ErrDailyNotConfigured
	}
	if strings.TrimSpace(p.RoomName) == "" {
		return "", errors.New("daily meeting token: empty room name")
	}

	// Daily exige timestamp UNIX em "exp" e somente nas propriedades do token.
	// Sanity: se ExpiresAt for zero ou no passado, força janela curta de 1h pra
	// evitar token "permanente" por bug de caller.
	exp := p.ExpiresAt
	if exp.IsZero() || exp.Before(time.Now().UTC()) {
		exp = time.Now().UTC().Add(1 * time.Hour)
	}

	displayName := strings.TrimSpace(p.UserName)
	if displayName == "" {
		displayName = "Plenya"
	}

	props := map[string]any{
		"room_name":          p.RoomName,
		"user_name":          displayName,
		"is_owner":           p.IsOwner,
		"exp":                exp.UTC().Unix(),
		"enable_screenshare": p.EnableScreenshare,
	}
	if p.StartCloudRecording {
		// Garante a capacidade no próprio token (defesa: salas antigas podem não
		// ter sido criadas com enable_recording) + auto-inicia ao entrar.
		props["enable_recording"] = "cloud"
		props["start_cloud_recording"] = true
	}
	if p.AutoStartTranscription {
		props["auto_start_transcription"] = true
	}
	if p.DisableRecordingUI {
		props["enable_recording_ui"] = false
	}
	body := map[string]any{"properties": props}

	resp, err := s.do(ctx, http.MethodPost, dailyAPIBaseURL+"/meeting-tokens", body)
	if err != nil {
		return "", fmt.Errorf("daily create meeting token http: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return "", errors.New("daily.co: invalid api key (401)")
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		errBody, _ := io.ReadAll(resp.Body)
		log.Printf("⚠️  Daily create meeting token: status=%d body=%s", resp.StatusCode, string(errBody))
		return "", fmt.Errorf("daily create meeting token: status %d", resp.StatusCode)
	}

	var parsed struct {
		Token string `json:"token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return "", fmt.Errorf("daily create meeting token decode: %w", err)
	}
	if parsed.Token == "" {
		return "", errors.New("daily create meeting token: empty token in response")
	}
	return parsed.Token, nil
}

// BuildJoinURL anexa o meeting_token na URL da sala como ?t=<token>.
// Daily prebuilt aceita esse param e usa o token automaticamente.
func BuildJoinURL(roomURL, token string) string {
	if token == "" {
		return roomURL
	}
	if strings.Contains(roomURL, "?") {
		return roomURL + "&t=" + token
	}
	return roomURL + "?t=" + token
}

// DeleteRoom remove a sala. Idempotente — 404 é tratado como sucesso.
func (s *DailyCoService) DeleteRoom(ctx context.Context, name string) error {
	if !s.IsConfigured() {
		return ErrDailyNotConfigured
	}
	if name == "" {
		return errors.New("daily delete room: empty name")
	}

	resp, err := s.do(ctx, http.MethodDelete, dailyAPIBaseURL+"/rooms/"+name, nil)
	if err != nil {
		return fmt.Errorf("daily delete room http: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK, http.StatusNoContent, http.StatusNotFound:
		return nil
	case http.StatusUnauthorized:
		return errors.New("daily.co: invalid api key (401)")
	default:
		return fmt.Errorf("daily delete room: status %d", resp.StatusCode)
	}
}

// ============================================================
// Gravação + transcrição: access-links e download (follow-up telemed)
// ============================================================

// GetRecordingAccessLink gera um link assinado de download do MP4 da gravação
// (GET /recordings/:id/access-link). NÃO baixamos o vídeo — só repassamos esse
// link temporário sob demanda (decisão "referência + link sob demanda").
// validForSecs default 3600 (1h); máx 12h.
func (s *DailyCoService) GetRecordingAccessLink(ctx context.Context, recordingID string, validForSecs int) (string, error) {
	if !s.IsConfigured() {
		return "", ErrDailyNotConfigured
	}
	if strings.TrimSpace(recordingID) == "" {
		return "", errors.New("daily recording access-link: empty recording id")
	}
	if validForSecs <= 0 {
		validForSecs = 3600
	}
	url := fmt.Sprintf("%s/recordings/%s/access-link?valid_for_secs=%d", dailyAPIBaseURL, recordingID, validForSecs)
	resp, err := s.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("daily recording access-link http: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		errBody, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("daily recording access-link: status %d body=%s", resp.StatusCode, string(errBody))
	}
	var parsed struct {
		DownloadLink string `json:"download_link"`
		Link         string `json:"link"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return "", fmt.Errorf("daily recording access-link decode: %w", err)
	}
	if parsed.DownloadLink != "" {
		return parsed.DownloadLink, nil
	}
	if parsed.Link != "" {
		return parsed.Link, nil
	}
	return "", errors.New("daily recording access-link: empty link in response")
}

// GetTranscriptAccessLink gera o link assinado pro arquivo WebVTT da transcrição
// (GET /transcript/:id/access-link, válido por 1h). O texto é pequeno — baixamos
// e persistimos (diferente do MP4, que é só referenciado).
func (s *DailyCoService) GetTranscriptAccessLink(ctx context.Context, transcriptID string) (string, error) {
	if !s.IsConfigured() {
		return "", ErrDailyNotConfigured
	}
	if strings.TrimSpace(transcriptID) == "" {
		return "", errors.New("daily transcript access-link: empty transcript id")
	}
	url := fmt.Sprintf("%s/transcript/%s/access-link", dailyAPIBaseURL, transcriptID)
	resp, err := s.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", fmt.Errorf("daily transcript access-link http: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		errBody, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("daily transcript access-link: status %d body=%s", resp.StatusCode, string(errBody))
	}
	var parsed struct {
		Link         string `json:"link"`
		DownloadLink string `json:"download_link"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return "", fmt.Errorf("daily transcript access-link decode: %w", err)
	}
	if parsed.Link != "" {
		return parsed.Link, nil
	}
	if parsed.DownloadLink != "" {
		return parsed.DownloadLink, nil
	}
	return "", errors.New("daily transcript access-link: empty link in response")
}

// FetchSignedURL baixa o conteúdo de uma URL assinada (S3) — sem Authorization
// (o link já é autenticado pelo querystring). Usado pra puxar o WebVTT.
func (s *DailyCoService) FetchSignedURL(ctx context.Context, signedURL string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, signedURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("fetch signed url: status %d", resp.StatusCode)
	}
	// Limite defensivo: VTT de teleconsulta tem KBs; 16MB cobre folga.
	return io.ReadAll(io.LimitReader(resp.Body, 16<<20))
}

// VerifyWebhookSignature valida a assinatura HMAC dos webhooks do Daily.
//
// Receita oficial: HMAC-SHA256 sobre (timestamp + "." + corpo_cru), com o segredo
// DECODIFICADO de base64 como chave, saída em base64, comparada com o header
// X-Webhook-Signature. timestamp é o header X-Webhook-Timestamp.
//
// Rejeita timestamp muito antigo (replay). Falha fechado se o segredo não está
// configurado (caller decide o comportamento por ambiente).
func (s *DailyCoService) VerifyWebhookSignature(signatureHeader, timestampHeader string, body []byte) error {
	secret := strings.TrimSpace(s.cfg.DailyCo.WebhookSecret)
	if secret == "" {
		return errors.New("daily webhook secret not configured")
	}
	if signatureHeader == "" || timestampHeader == "" {
		return errors.New("daily webhook: missing signature/timestamp headers")
	}
	// Replay protection: rejeita timestamps fora de ±5min.
	if ts, err := strconv.ParseInt(strings.TrimSpace(timestampHeader), 10, 64); err == nil {
		skew := time.Now().UTC().Unix() - ts
		if skew < 0 {
			skew = -skew
		}
		if skew > 300 {
			return fmt.Errorf("daily webhook: stale timestamp (skew=%ds)", skew)
		}
	}
	key, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		// Alguns segredos podem não ser base64 — usa cru como fallback.
		key = []byte(secret)
	}
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(timestampHeader + "." + string(body)))
	expected := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	if !hmac.Equal([]byte(expected), []byte(strings.TrimSpace(signatureHeader))) {
		return errors.New("daily webhook: signature mismatch")
	}
	return nil
}

// do é o helper HTTP — adiciona Authorization Bearer + Content-Type JSON.
func (s *DailyCoService) do(ctx context.Context, method, url string, body any) (*http.Response, error) {
	var reader io.Reader
	if body != nil {
		raw, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reader = bytes.NewReader(raw)
	}
	req, err := http.NewRequestWithContext(ctx, method, url, reader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+s.cfg.DailyCo.APIKey)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return s.client.Do(req)
}

// randHex retorna n bytes random como string hexadecimal (2*n chars).
// Usado pra gerar sufixo único de room name. Falha → fallback "00...".
func randHex(n int) string {
	buf := make([]byte, n)
	if _, err := rand.Read(buf); err != nil {
		return strings.Repeat("0", n*2)
	}
	return hex.EncodeToString(buf)
}

// sanitizeDailyName converte uma string em formato aceito pelo Daily:
// lowercase, apenas [a-z0-9-]. Espaços → "-", outros chars removidos.
func sanitizeDailyName(s string) string {
	s = strings.ToLower(s)
	var b strings.Builder
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9':
			b.WriteRune(r)
		case r == ' ' || r == '_' || r == '-':
			b.WriteRune('-')
		}
	}
	out := b.String()
	// Colapsa múltiplos hífens.
	for strings.Contains(out, "--") {
		out = strings.ReplaceAll(out, "--", "-")
	}
	return strings.Trim(out, "-")
}
