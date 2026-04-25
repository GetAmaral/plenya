package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

const expoPushAPIURL = "https://exp.host/--/api/v2/push/send"

// PushService dispara push notifications via Expo Push API.
//
// Funciona com tokens Expo (formato ExponentPushToken[xxx]) — a Expo
// internamente roteia pra APNs/FCM. Quando migrarmos pra FCM/APNs nativo
// (>50k devices), trocar a impl mantendo o contrato Send.
type PushService struct {
	db         *gorm.DB
	httpClient *http.Client
}

func NewPushService(db *gorm.DB) *PushService {
	return &PushService{
		db: db,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// PushPayload é o que vai pro device. URL é o deep link/path interno
// que o usePushNotificationRouter vai abrir ao tocar.
type PushPayload struct {
	Title string         `json:"title"`
	Body  string         `json:"body"`
	URL   string         `json:"url,omitempty"`
	Data  map[string]any `json:"data,omitempty"`
	Badge *int           `json:"badge,omitempty"`
}

type expoMessage struct {
	To       string         `json:"to"`
	Title    string         `json:"title"`
	Body     string         `json:"body"`
	Data     map[string]any `json:"data,omitempty"`
	Sound    string         `json:"sound,omitempty"`
	Badge    *int           `json:"badge,omitempty"`
	Priority string         `json:"priority,omitempty"`
}

type expoTicket struct {
	Status  string         `json:"status"`
	ID      string         `json:"id,omitempty"`
	Message string         `json:"message,omitempty"`
	Details map[string]any `json:"details,omitempty"`
}

type expoResponse struct {
	Data   []expoTicket `json:"data"`
	Errors []struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
}

// Send envia push para um único usuário. Faz fan-out por todos os
// device tokens ativos do usuário, descarta tokens inválidos
// (DeviceNotRegistered) automaticamente.
func (s *PushService) Send(userID uuid.UUID, payload PushPayload) error {
	var tokens []models.DeviceToken
	if err := s.db.Where("user_id = ?", userID).Find(&tokens).Error; err != nil {
		return err
	}
	if len(tokens) == 0 {
		return nil
	}

	messages := make([]expoMessage, 0, len(tokens))
	for _, t := range tokens {
		data := payload.Data
		if payload.URL != "" {
			if data == nil {
				data = make(map[string]any, 1)
			}
			data["url"] = payload.URL
		}
		messages = append(messages, expoMessage{
			To:       t.Token,
			Title:    payload.Title,
			Body:     payload.Body,
			Data:     data,
			Sound:    "default",
			Badge:    payload.Badge,
			Priority: "high",
		})
	}

	body, err := json.Marshal(messages)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, expoPushAPIURL, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Encoding", "gzip, deflate")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		return fmt.Errorf("expo push api error %d: %s", resp.StatusCode, string(respBody))
	}

	var parsed expoResponse
	if err := json.Unmarshal(respBody, &parsed); err != nil {
		return err
	}
	if len(parsed.Errors) > 0 {
		return errors.New(parsed.Errors[0].Message)
	}

	// Tokens inválidos vêm como ticket status=error code=DeviceNotRegistered.
	// Removemos eles do banco pra não tentar de novo.
	for i, ticket := range parsed.Data {
		if i >= len(tokens) {
			break
		}
		if ticket.Status == "error" {
			code, _ := ticket.Details["error"].(string)
			if code == "DeviceNotRegistered" || code == "InvalidCredentials" {
				_ = s.db.Delete(&models.DeviceToken{}, "id = ?", tokens[i].ID).Error
			}
		}
	}

	return nil
}

// SendToMany faz fan-out pra múltiplos usuários — útil quando vários
// profissionais devem ser notificados (lead inbound assigned, etc).
func (s *PushService) SendToMany(userIDs []uuid.UUID, payload PushPayload) {
	for _, uid := range userIDs {
		if err := s.Send(uid, payload); err != nil {
			// log + continue: falha em 1 não deve abortar os outros
			fmt.Printf("[push] failed userId=%s: %v\n", uid, err)
		}
	}
}

// SanitizeBody sanitiza body de notificação removendo PHI explícito
// (CPF, RG, números longos). Use para conteúdo gerado a partir de payloads
// vindos do paciente (mensagem WA, email).
func SanitizeBody(s string) string {
	out := s
	for _, pii := range []string{"CPF:", "RG:"} {
		if i := strings.Index(out, pii); i >= 0 {
			out = out[:i] + "[…]"
		}
	}
	if len(out) > 140 {
		out = out[:137] + "…"
	}
	return out
}
