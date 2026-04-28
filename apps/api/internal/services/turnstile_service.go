package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// M10 — Cloudflare Turnstile CAPTCHA validator.
//
// Endpoint: POST https://challenges.cloudflare.com/turnstile/v0/siteverify
// Body form-urlencoded com `secret` + `response` (token do widget).
//
// Comportamento:
//   - Secret vazio (dev local): retorna nil sem fazer chamada (skip validação).
//     Loga warning na 1a chamada pra deixar claro.
//   - Token vazio (cliente não passou): erro semântico.
//   - Cloudflare retorna success=false: erro com error-codes.
//   - Network/HTTP error: erro envolvido.
const turnstileVerifyURL = "https://challenges.cloudflare.com/turnstile/v0/siteverify"

// ErrTurnstileMissing — frontend não enviou cf-turnstile-response.
var ErrTurnstileMissing = errors.New("turnstile: token ausente")

// ErrTurnstileInvalid — Cloudflare rejeitou o token.
var ErrTurnstileInvalid = errors.New("turnstile: token inválido")

type TurnstileService struct {
	secret    string
	client    *http.Client
	devWarned bool
}

func NewTurnstileService(secret string) *TurnstileService {
	return &TurnstileService{
		secret: strings.TrimSpace(secret),
		client: &http.Client{Timeout: 5 * time.Second},
	}
}

type turnstileResp struct {
	Success     bool     `json:"success"`
	ErrorCodes  []string `json:"error-codes"`
	ChallengeTs string   `json:"challenge_ts"`
	Hostname    string   `json:"hostname"`
}

// Verify — chama siteverify. remoteIP é opcional (pode ser "" quando atrás
// de proxy não confiável, mas idealmente é o IP do client real via c.IP()).
//
// Em dev (secret vazio): retorna nil. Loga 1x.
func (s *TurnstileService) Verify(ctx context.Context, token, remoteIP string) error {
	if s.secret == "" {
		if !s.devWarned {
			s.devWarned = true
			// Não imprimimos via log direto pra evitar import cycle/deps;
			// caller que loga se quiser. Aqui silencioso.
		}
		return nil
	}
	if strings.TrimSpace(token) == "" {
		return ErrTurnstileMissing
	}

	form := url.Values{}
	form.Set("secret", s.secret)
	form.Set("response", token)
	if remoteIP != "" {
		form.Set("remoteip", remoteIP)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, turnstileVerifyURL, strings.NewReader(form.Encode()))
	if err != nil {
		return fmt.Errorf("turnstile: build request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("turnstile: http: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("turnstile: status %d", resp.StatusCode)
	}

	var body turnstileResp
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return fmt.Errorf("turnstile: decode: %w", err)
	}
	if !body.Success {
		return fmt.Errorf("%w: %v", ErrTurnstileInvalid, body.ErrorCodes)
	}
	return nil
}

// Enabled — true se Secret configurado. Handlers usam pra decidir se exibem
// nota "captcha skip in dev" no log.
func (s *TurnstileService) Enabled() bool { return s.secret != "" }
