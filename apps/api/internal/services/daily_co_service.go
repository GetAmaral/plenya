// Package services — DailyCoService cria/deleta salas Daily.co pra
// teleconsulta embedada no EMR (Calendar V1, Bloco C).
//
// API: https://docs.daily.co/reference/rest-api
//
// Decisões:
//
//  1. Sala "private" + token de acesso fica V2 (free tier permite, mas sala
//     pública por URL difícil-de-adivinhar é suficiente pro MVP médico-paciente).
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
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/plenya/api/internal/config"
)

const (
	dailyAPIBaseURL          = "https://api.daily.co/v1"
	dailyHTTPClientTimeout   = 15 * time.Second
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

	body := map[string]any{
		"name": name,
		"properties": map[string]any{
			"exp":                expiresAt.UTC().Unix(),
			"eject_at_room_exp":  true,
			"enable_screenshare": true,
			"enable_chat":        true,
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
