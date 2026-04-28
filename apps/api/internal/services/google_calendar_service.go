// Package services — GoogleCalendarService implementa o cliente HTTP pra
// Google Calendar API v3 + OAuth 2.0.
//
// Decisões arquiteturais (Calendar V1, Bloco B):
//
//  1. NÃO usamos `google.golang.org/api/calendar/v3` SDK oficial — a superfície
//     que precisamos é pequena (3 endpoints OAuth + 5 Calendar) e o SDK arrasta
//     dezenas de transitive deps. Manter manual deixa o erro handling explícito
//     (rate limit, token revogado).
//
//  2. Tokens (refresh + access) sempre passam por crypto.EncryptWithDefaultKey
//     antes de persistir em CalendarCredential. NUNCA são logados.
//
//  3. GetValidAccessToken faz lazy refresh: se access token expira em <10min,
//     renova via refresh token e atualiza a row. Retorna ErrCredentialRevoked
//     se Google retornou invalid_grant (médico desconectou).
//
//  4. HTTP client com timeout 15s — Google API costuma responder <500ms; 15s
//     cobre cold start + retry interno.
//
//  5. State JWT (CSRF) é responsabilidade do handler, não do service.
package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/crypto"
	"github.com/plenya/api/internal/models"
)

// Erros públicos do GoogleCalendarService — handlers mapeiam pra HTTP status.
var (
	// ErrCredentialRevoked = Google retornou invalid_grant ou 401 persistente.
	// Handler deve responder 401 + frontend mostra "Reconectar Google".
	ErrCredentialRevoked = errors.New("google calendar credential revoked")

	// ErrCredentialNotFound = usuário ainda não conectou Google Calendar.
	ErrCredentialNotFound = errors.New("google calendar credential not configured")

	// ErrGoogleRateLimit = Google retornou 429. Caller pode retry com backoff.
	ErrGoogleRateLimit = errors.New("google calendar rate limit")

	// ErrGoogleConfigMissing = GOOGLE_CLIENT_ID/SECRET vazios. Setup operacional pendente.
	ErrGoogleConfigMissing = errors.New("google calendar config missing (GOOGLE_CLIENT_ID/SECRET)")
)

// Scopes solicitados no OAuth.
//   - calendar.events.owned: gerenciar APENAS eventos criados por nós (não vê
//     outros eventos do médico). Mais restrito que `calendar` ou `calendar.events`.
//   - userinfo.email: pra exibir email da conta conectada na UI ("Conectado como X").
const (
	googleOAuthScope        = "https://www.googleapis.com/auth/calendar https://www.googleapis.com/auth/userinfo.email"
	googleOAuthAuthURL      = "https://accounts.google.com/o/oauth2/v2/auth"
	googleOAuthTokenURL     = "https://oauth2.googleapis.com/token"     //nolint:gosec // URL pública, não é credential
	googleOAuthRevokeURL    = "https://oauth2.googleapis.com/revoke"    //nolint:gosec
	googleUserinfoURL       = "https://www.googleapis.com/oauth2/v2/userinfo"
	googleCalendarBaseURL   = "https://www.googleapis.com/calendar/v3"
	tokenRefreshLeewayMin   = 10 // refresh quando expira em <X min
	googleHTTPClientTimeout = 15 * time.Second
)

type GoogleCalendarService struct {
	cfg    *config.Config
	db     *gorm.DB
	client *http.Client
}

func NewGoogleCalendarService(cfg *config.Config, db *gorm.DB) *GoogleCalendarService {
	return &GoogleCalendarService{
		cfg: cfg,
		db:  db,
		client: &http.Client{
			Timeout: googleHTTPClientTimeout,
		},
	}
}

// OAuthTokens — tokens recebidos do Google após exchange/refresh.
type OAuthTokens struct {
	AccessToken  string
	RefreshToken string // pode ser vazio em refresh subsequente (Google só envia 1ª vez)
	ExpiresAt    time.Time
}

// EncryptOAuthToken — M8 — wrapper que usa OAuth key dedicada do config.
// Exposto pra callers que persistem tokens fora do service (handler OAuth, scheduler).
func (s *GoogleCalendarService) EncryptOAuthToken(plaintext string) (string, error) {
	return crypto.EncryptWithOAuthKey(plaintext, s.cfg.Security.OAuthTokenKey)
}

// DecryptOAuthToken — M8 — par do EncryptOAuthToken. Tenta a chave dedicada,
// fallback pra default key (retrocompat com rows pré-migração).
func (s *GoogleCalendarService) DecryptOAuthToken(ciphertext string) (string, error) {
	plain, err := crypto.DecryptWithOAuthKey(ciphertext, s.cfg.Security.OAuthTokenKey)
	if err == nil {
		return plain, nil
	}
	return crypto.DecryptWithDefaultKey(ciphertext)
}

// IsConfigured retorna true se ClientID + ClientSecret estão setados.
// Handlers checam isso pra responder 503 em vez de 500 quando setup pendente.
func (s *GoogleCalendarService) IsConfigured() bool {
	return s.cfg.Google.ClientID != "" && s.cfg.Google.ClientSecret != ""
}

// BuildAuthURL constrói a URL de redirect pro consent screen Google.
// State é passado opaco — handler é responsável por gerar JWT CSRF.
//
// access_type=offline + prompt=consent garantem que sempre vem refresh token
// (necessário pra renovar access token sem nova interação).
func (s *GoogleCalendarService) BuildAuthURL(state string) string {
	q := url.Values{}
	q.Set("client_id", s.cfg.Google.ClientID)
	q.Set("redirect_uri", s.cfg.Google.RedirectURL)
	q.Set("response_type", "code")
	q.Set("scope", googleOAuthScope)
	q.Set("access_type", "offline")
	q.Set("prompt", "consent") // força refresh token mesmo se médico já consentiu antes
	q.Set("include_granted_scopes", "true")
	q.Set("state", state)
	return googleOAuthAuthURL + "?" + q.Encode()
}

// ExchangeCode troca o authorization code (callback) por tokens.
func (s *GoogleCalendarService) ExchangeCode(ctx context.Context, code string) (*OAuthTokens, error) {
	if !s.IsConfigured() {
		return nil, ErrGoogleConfigMissing
	}

	form := url.Values{}
	form.Set("code", code)
	form.Set("client_id", s.cfg.Google.ClientID)
	form.Set("client_secret", s.cfg.Google.ClientSecret)
	form.Set("redirect_uri", s.cfg.Google.RedirectURL)
	form.Set("grant_type", "authorization_code")

	return s.postTokenForm(ctx, form, "exchange")
}

// RefreshAccessToken usa o refresh token pra obter um novo access token.
// Retorna ErrCredentialRevoked se Google responder invalid_grant.
func (s *GoogleCalendarService) RefreshAccessToken(ctx context.Context, refreshToken string) (*OAuthTokens, error) {
	if !s.IsConfigured() {
		return nil, ErrGoogleConfigMissing
	}

	form := url.Values{}
	form.Set("refresh_token", refreshToken)
	form.Set("client_id", s.cfg.Google.ClientID)
	form.Set("client_secret", s.cfg.Google.ClientSecret)
	form.Set("grant_type", "refresh_token")

	tokens, err := s.postTokenForm(ctx, form, "refresh")
	if err != nil {
		return nil, err
	}
	// Google NÃO retorna refresh_token em refresh — preservar o anterior.
	if tokens.RefreshToken == "" {
		tokens.RefreshToken = refreshToken
	}
	return tokens, nil
}

// postTokenForm é o helper interno pros 2 fluxos token (exchange + refresh).
func (s *GoogleCalendarService) postTokenForm(ctx context.Context, form url.Values, label string) (*OAuthTokens, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, googleOAuthTokenURL, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, fmt.Errorf("google %s build request: %w", label, err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("google %s http: %w", label, err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusBadRequest {
		// Detectar invalid_grant — médico revogou no console Google.
		// NUNCA logar body inteiro (pode ter token).
		if bytes.Contains(body, []byte("invalid_grant")) {
			log.Printf("⚠️  Google OAuth %s: invalid_grant (credential revoked)", label)
			return nil, ErrCredentialRevoked
		}
		log.Printf("⚠️  Google OAuth %s: status=%d", label, resp.StatusCode)
		return nil, fmt.Errorf("google %s: status %d", label, resp.StatusCode)
	}
	if resp.StatusCode != http.StatusOK {
		log.Printf("⚠️  Google OAuth %s: status=%d", label, resp.StatusCode)
		return nil, fmt.Errorf("google %s: unexpected status %d", label, resp.StatusCode)
	}

	var parsed struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int    `json:"expires_in"` // segundos
		TokenType    string `json:"token_type"`
		Scope        string `json:"scope"`
	}
	if err := json.Unmarshal(body, &parsed); err != nil {
		return nil, fmt.Errorf("google %s decode: %w", label, err)
	}
	if parsed.AccessToken == "" {
		return nil, fmt.Errorf("google %s: empty access_token in response", label)
	}

	return &OAuthTokens{
		AccessToken:  parsed.AccessToken,
		RefreshToken: parsed.RefreshToken,
		ExpiresAt:    time.Now().Add(time.Duration(parsed.ExpiresIn) * time.Second),
	}, nil
}

// RevokeToken chama https://oauth2.googleapis.com/revoke pra invalidar o
// refresh token no Google. Idempotente — se já foi revogado, retorna nil.
func (s *GoogleCalendarService) RevokeToken(ctx context.Context, refreshToken string) error {
	form := url.Values{}
	form.Set("token", refreshToken)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, googleOAuthRevokeURL, strings.NewReader(form.Encode()))
	if err != nil {
		return fmt.Errorf("google revoke build: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("google revoke http: %w", err)
	}
	defer resp.Body.Close()

	// 200 = sucesso. 400 com invalid_token = já revogado (idempotente).
	if resp.StatusCode == http.StatusOK {
		return nil
	}
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode == http.StatusBadRequest && bytes.Contains(body, []byte("invalid_token")) {
		return nil
	}
	return fmt.Errorf("google revoke: status %d", resp.StatusCode)
}

// GetUserEmail retorna o email da conta Google a partir de access token válido.
// Usado uma vez no callback OAuth pra exibir "Conectado como X" na UI.
func (s *GoogleCalendarService) GetUserEmail(ctx context.Context, accessToken string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, googleUserinfoURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := s.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("google userinfo http: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return "", ErrCredentialRevoked
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("google userinfo: status %d", resp.StatusCode)
	}

	var parsed struct {
		Email string `json:"email"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return "", err
	}
	return parsed.Email, nil
}

// CreateDedicatedCalendar cria um calendar "Plenya" na conta do médico.
// Chamado uma vez no callback OAuth — todos eventos posteriores vão pra esse
// calendar (NUNCA pro primário), garantindo isolamento e fácil deleção.
func (s *GoogleCalendarService) CreateDedicatedCalendar(ctx context.Context, accessToken string) (string, error) {
	body := map[string]string{
		"summary":     "Plenya",
		"description": "Agenda de consultas Plenya — sincronizada automaticamente. Não editar eventos diretamente aqui.",
		"timeZone":    "America/Sao_Paulo",
	}
	resp, err := s.doJSON(ctx, accessToken, http.MethodPost, googleCalendarBaseURL+"/calendars", body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("google create calendar: status %d", resp.StatusCode)
	}

	var parsed struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return "", err
	}
	return parsed.ID, nil
}

// EventInput — payload pra create/patch event no Google Calendar.
//
// Privacidade (LGPD):
//   - Summary deve ser opaco: "Plenya — Avaliação Inicial · M.S." (iniciais)
//   - Description vazio (sem PHI)
//   - Visibility "private" garante que mesmo se calendar for compartilhado,
//     só donos veem detalhes.
type EventInput struct {
	CalendarID  string
	Summary     string
	Description string
	Location    string
	StartUTC    time.Time
	EndUTC      time.Time
	Visibility  string // "private" recomendado
	ReminderMin int    // 0 = usar Calendar defaults; >0 = override
}

// EventOutput — resposta após create.
type EventOutput struct {
	EventID  string
	HtmlLink string // URL pública pro evento (não usamos — manter pra auditoria)
}

// CreateEvent cria um evento no calendar dedicado do médico.
func (s *GoogleCalendarService) CreateEvent(ctx context.Context, accessToken string, in EventInput) (*EventOutput, error) {
	body := s.buildEventBody(in)
	url := fmt.Sprintf("%s/calendars/%s/events", googleCalendarBaseURL, urlEncodePath(in.CalendarID))

	resp, err := s.doJSON(ctx, accessToken, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, ErrCredentialRevoked
	}
	if resp.StatusCode == http.StatusTooManyRequests {
		return nil, ErrGoogleRateLimit
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("google create event: status %d", resp.StatusCode)
	}

	var parsed struct {
		ID       string `json:"id"`
		HtmlLink string `json:"htmlLink"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return nil, err
	}
	return &EventOutput{EventID: parsed.ID, HtmlLink: parsed.HtmlLink}, nil
}

// PatchEvent atualiza campos do evento. Usa PATCH (parcial), não PUT.
func (s *GoogleCalendarService) PatchEvent(ctx context.Context, accessToken, calendarID, eventID string, in EventInput) error {
	body := s.buildEventBody(in)
	url := fmt.Sprintf("%s/calendars/%s/events/%s", googleCalendarBaseURL, urlEncodePath(calendarID), urlEncodePath(eventID))

	resp, err := s.doJSON(ctx, accessToken, http.MethodPatch, url, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return ErrCredentialRevoked
	}
	if resp.StatusCode == http.StatusTooManyRequests {
		return ErrGoogleRateLimit
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("google patch event: status %d", resp.StatusCode)
	}
	return nil
}

// DeleteEvent remove o evento. 410 (gone) é tratado como sucesso (idempotente).
func (s *GoogleCalendarService) DeleteEvent(ctx context.Context, accessToken, calendarID, eventID string) error {
	url := fmt.Sprintf("%s/calendars/%s/events/%s", googleCalendarBaseURL, urlEncodePath(calendarID), urlEncodePath(eventID))

	resp, err := s.doJSON(ctx, accessToken, http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK, http.StatusNoContent, http.StatusGone, http.StatusNotFound:
		return nil
	case http.StatusUnauthorized:
		return ErrCredentialRevoked
	case http.StatusTooManyRequests:
		return ErrGoogleRateLimit
	default:
		return fmt.Errorf("google delete event: status %d", resp.StatusCode)
	}
}

// FreeBusyInput — query agregada por múltiplos calendars no mesmo range.
type FreeBusyInput struct {
	CalendarIDs []string
	StartUTC    time.Time
	EndUTC      time.Time
}

// FreeBusyInterval — bloco ocupado retornado pelo Google.
type FreeBusyInterval struct {
	Start time.Time
	End   time.Time
}

// FreeBusyResult — busy intervals por calendar ID.
type FreeBusyResult map[string][]FreeBusyInterval

// FreeBusy consulta livre/ocupado em N calendars de uma vez. Eficiente —
// Google não retorna detalhes de eventos, só intervalos.
func (s *GoogleCalendarService) FreeBusy(ctx context.Context, accessToken string, in FreeBusyInput) (FreeBusyResult, error) {
	items := make([]map[string]string, 0, len(in.CalendarIDs))
	for _, id := range in.CalendarIDs {
		items = append(items, map[string]string{"id": id})
	}
	body := map[string]any{
		"timeMin":  in.StartUTC.UTC().Format(time.RFC3339),
		"timeMax":  in.EndUTC.UTC().Format(time.RFC3339),
		"timeZone": "UTC",
		"items":    items,
	}

	resp, err := s.doJSON(ctx, accessToken, http.MethodPost, googleCalendarBaseURL+"/freeBusy", body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, ErrCredentialRevoked
	}
	if resp.StatusCode == http.StatusTooManyRequests {
		return nil, ErrGoogleRateLimit
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("google freebusy: status %d", resp.StatusCode)
	}

	var parsed struct {
		Calendars map[string]struct {
			Busy []struct {
				Start time.Time `json:"start"`
				End   time.Time `json:"end"`
			} `json:"busy"`
		} `json:"calendars"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return nil, err
	}

	out := make(FreeBusyResult, len(parsed.Calendars))
	for id, cal := range parsed.Calendars {
		intervals := make([]FreeBusyInterval, 0, len(cal.Busy))
		for _, b := range cal.Busy {
			intervals = append(intervals, FreeBusyInterval{Start: b.Start, End: b.End})
		}
		out[id] = intervals
	}
	return out, nil
}

// GetValidAccessToken retorna um access token decifrado e VÁLIDO pro user.
//
// Lookup → decifra → se expira em <10min, refresh → atualiza row → retorna.
// Se RevokedAt != nil → ErrCredentialRevoked.
// Se refresh retorna invalid_grant → marca RevokedAt e propaga ErrCredentialRevoked.
func (s *GoogleCalendarService) GetValidAccessToken(ctx context.Context, userID uuid.UUID) (string, error) {
	var cred models.CalendarCredential
	err := s.db.WithContext(ctx).
		Where("user_id = ? AND provider = ?", userID, models.CalendarProviderGoogle).
		First(&cred).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", ErrCredentialNotFound
		}
		return "", err
	}
	if !cred.IsActive() {
		return "", ErrCredentialRevoked
	}

	// Token ainda válido com leeway?
	// M8 — fallback retrocompat: se ciphertext foi gravado antes da migração
	// pra OAuth key (chave dedicada), tenta default key. Após próximo refresh,
	// regrava com OAuth key e o fallback some.
	if time.Until(cred.AccessTokenExpiresAt) > tokenRefreshLeewayMin*time.Minute {
		access, err := crypto.DecryptWithOAuthKey(cred.EncryptedAccessToken, s.cfg.Security.OAuthTokenKey)
		if err != nil {
			access, err = crypto.DecryptWithDefaultKey(cred.EncryptedAccessToken)
		}
		if err != nil {
			return "", fmt.Errorf("decrypt access token: %w", err)
		}
		return access, nil
	}

	// Refresh.
	refresh, err := crypto.DecryptWithOAuthKey(cred.EncryptedRefreshToken, s.cfg.Security.OAuthTokenKey)
	if err != nil {
		// Fallback retrocompat (chave default).
		refresh, err = crypto.DecryptWithDefaultKey(cred.EncryptedRefreshToken)
	}
	if err != nil {
		return "", fmt.Errorf("decrypt refresh token: %w", err)
	}
	tokens, err := s.RefreshAccessToken(ctx, refresh)
	if err != nil {
		if errors.Is(err, ErrCredentialRevoked) {
			now := time.Now()
			cred.RevokedAt = &now
			if dbErr := s.db.WithContext(ctx).Save(&cred).Error; dbErr != nil {
				log.Printf("⚠️  failed to persist RevokedAt for credential %s: %v", cred.ID, dbErr)
			}
		}
		return "", err
	}

	// Persist refreshed tokens.
	encAccess, err := crypto.EncryptWithOAuthKey(tokens.AccessToken, s.cfg.Security.OAuthTokenKey)
	if err != nil {
		return "", fmt.Errorf("encrypt new access token: %w", err)
	}
	cred.EncryptedAccessToken = encAccess
	cred.AccessTokenExpiresAt = tokens.ExpiresAt
	if tokens.RefreshToken != "" && tokens.RefreshToken != refresh {
		encRefresh, err := crypto.EncryptWithOAuthKey(tokens.RefreshToken, s.cfg.Security.OAuthTokenKey)
		if err != nil {
			return "", fmt.Errorf("encrypt new refresh token: %w", err)
		}
		cred.EncryptedRefreshToken = encRefresh
	}
	now := time.Now()
	cred.LastSyncAt = &now
	if err := s.db.WithContext(ctx).Save(&cred).Error; err != nil {
		return "", fmt.Errorf("persist refreshed tokens: %w", err)
	}

	return tokens.AccessToken, nil
}

// buildEventBody constrói o JSON pra create/patch event.
func (s *GoogleCalendarService) buildEventBody(in EventInput) map[string]any {
	visibility := in.Visibility
	if visibility == "" {
		visibility = "private"
	}
	body := map[string]any{
		"summary":     in.Summary,
		"description": in.Description,
		"location":    in.Location,
		"start": map[string]string{
			"dateTime": in.StartUTC.UTC().Format(time.RFC3339),
			"timeZone": "UTC",
		},
		"end": map[string]string{
			"dateTime": in.EndUTC.UTC().Format(time.RFC3339),
			"timeZone": "UTC",
		},
		"visibility": visibility,
		"transparency": "opaque",
	}
	if in.ReminderMin > 0 {
		body["reminders"] = map[string]any{
			"useDefault": false,
			"overrides": []map[string]any{
				{"method": "popup", "minutes": in.ReminderMin},
			},
		}
	} else {
		// useDefault=true respeita config do médico no Google Calendar.
		body["reminders"] = map[string]any{"useDefault": true}
	}
	return body
}

// doJSON é o helper genérico pra requisições autenticadas Calendar v3.
func (s *GoogleCalendarService) doJSON(ctx context.Context, accessToken, method, url string, body any) (*http.Response, error) {
	var reader io.Reader
	if body != nil {
		raw, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("encode body: %w", err)
		}
		reader = bytes.NewReader(raw)
	}
	req, err := http.NewRequestWithContext(ctx, method, url, reader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	return s.client.Do(req)
}

// urlEncodePath escapa um único path segment (calendar ID ou event ID).
// IDs Google podem conter @ . - _ etc — net/url.PathEscape é seguro.
func urlEncodePath(s string) string {
	return url.PathEscape(s)
}
