package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	webpush "github.com/SherClockHolmes/webpush-go"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// WebPushService dispara notificações Web Push (protocolo VAPID) pros navegadores
// inscritos da equipe. Implementa PushSender, então entra no mesmo dispatchPush do
// NotificationService em paralelo ao Expo (mobile) — ver CompositePushSender.
//
// Quando as chaves VAPID não estão configuradas (PublicKey/PrivateKey vazios), o
// serviço fica desligado: Send/Subscribe viram no-op graceful e nada quebra.
type WebPushService struct {
	db         *gorm.DB
	publicKey  string
	privateKey string
	subject    string
}

func NewWebPushService(db *gorm.DB, publicKey, privateKey, subject string) *WebPushService {
	return &WebPushService{
		db:         db,
		publicKey:  publicKey,
		privateKey: privateKey,
		// A webpush-go prefixa "mailto:" sozinha em subjects que não são https.
		// Se o env já vier "mailto:email", viraria "mailto:mailto:email" e a Apple
		// rejeita o JWT (BadJwtToken). Normalizamos pra e-mail puro (ou deixamos a
		// URL https intacta) — robusto a qualquer formato de env.
		subject: normalizeVAPIDSubject(subject),
	}
}

// normalizeVAPIDSubject remove um prefixo "mailto:" redundante (case-insensitive),
// preservando subjects https:// como estão.
func normalizeVAPIDSubject(s string) string {
	t := strings.TrimSpace(s)
	if strings.HasPrefix(strings.ToLower(t), "https:") {
		return t
	}
	for strings.HasPrefix(strings.ToLower(t), "mailto:") {
		t = t[len("mailto:"):]
	}
	return t
}

// endpointHost devolve só o host do endpoint, pro log identificar o serviço de push
// (web.push.apple.com, fcm.googleapis.com) sem despejar o token do aparelho.
func endpointHost(endpoint string) string {
	if u, err := url.Parse(endpoint); err == nil && u.Host != "" {
		return u.Host
	}
	return "?"
}

// Enabled indica se há chaves VAPID configuradas.
func (s *WebPushService) Enabled() bool {
	return s.publicKey != "" && s.privateKey != ""
}

// PublicKey devolve a chave pública VAPID (exposta ao frontend pra subscribe).
func (s *WebPushService) PublicKey() string { return s.publicKey }

// webPushPayload é o JSON que chega no service worker (sw.js).
type webPushPayload struct {
	Title string         `json:"title"`
	Body  string         `json:"body"`
	URL   string         `json:"url,omitempty"`
	Data  map[string]any `json:"data,omitempty"`
}

// Send envia push para todas as inscrições do usuário. Inscrições com endpoint
// expirado (404/410) são apagadas. Falha em uma não aborta as outras.
func (s *WebPushService) Send(userID uuid.UUID, payload PushPayload) error {
	if !s.Enabled() {
		return nil
	}

	var subs []models.WebPushSubscription
	if err := s.db.Where("user_id = ?", userID).Find(&subs).Error; err != nil {
		return err
	}
	if len(subs) == 0 {
		return nil
	}

	body, err := json.Marshal(webPushPayload{
		Title: payload.Title,
		Body:  payload.Body,
		URL:   payload.URL,
		Data:  payload.Data,
	})
	if err != nil {
		return err
	}

	var firstErr error
	for _, sub := range subs {
		if err := s.sendOne(sub, body); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}

func (s *WebPushService) sendOne(sub models.WebPushSubscription, body []byte) error {
	resp, err := webpush.SendNotification(body, &webpush.Subscription{
		Endpoint: sub.Endpoint,
		Keys:     webpush.Keys{P256dh: sub.P256dh, Auth: sub.Auth},
	}, &webpush.Options{
		Subscriber:      s.subject,
		VAPIDPublicKey:  s.publicKey,
		VAPIDPrivateKey: s.privateKey,
		TTL:             60,
		Urgency:         webpush.UrgencyHigh,
	})
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	switch {
	case resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusGone:
		// Endpoint morto — limpa a inscrição pra não tentar de novo.
		log.Printf("[webpush] inscrição expirada (%d), removendo: user=%s host=%s", resp.StatusCode, sub.UserID, endpointHost(sub.Endpoint))
		_ = s.db.Delete(&models.WebPushSubscription{}, "id = ?", sub.ID).Error
		return nil
	case resp.StatusCode >= 400:
		return fmt.Errorf("web push error %d (endpoint=%.40s)", resp.StatusCode, sub.Endpoint)
	}

	// Log de entrega: sem isto, "o push não chegou" vira adivinhação — não dá pra separar
	// "o servidor não mandou" de "o aparelho não mostrou". Duas investigações já se perderam
	// aqui. Ver docs/emr/plano-webpush-notificacoes.md.
	log.Printf("[webpush] entregue %d: user=%s host=%s", resp.StatusCode, sub.UserID, endpointHost(sub.Endpoint))

	// Sucesso — marca atividade (best-effort).
	_ = s.db.Model(&models.WebPushSubscription{}).
		Where("id = ?", sub.ID).
		Update("last_seen_at", gorm.Expr("now()")).Error
	return nil
}

// WebPushSubscribeInput é o que o navegador manda (PushSubscription serializado).
type WebPushSubscribeInput struct {
	Endpoint string
	P256dh   string
	Auth     string
	Label    string
	UA       string
}

// Subscribe registra (ou atualiza) uma inscrição para o usuário, idempotente por endpoint.
func (s *WebPushService) Subscribe(userID uuid.UUID, in WebPushSubscribeInput) (*models.WebPushSubscription, error) {
	if in.Endpoint == "" || in.P256dh == "" || in.Auth == "" {
		return nil, errors.New("inscrição inválida: endpoint/keys ausentes")
	}

	var existing models.WebPushSubscription
	err := s.db.Where("endpoint = ?", in.Endpoint).First(&existing).Error
	switch {
	case err == nil:
		// Endpoint já existe — pode ter trocado de dono (mesmo navegador, outro login).
		updates := map[string]any{
			"user_id":      userID,
			"p256dh":       in.P256dh,
			"auth":         in.Auth,
			"last_seen_at": gorm.Expr("now()"),
		}
		if in.Label != "" {
			updates["device_label"] = in.Label
		}
		if in.UA != "" {
			updates["user_agent"] = in.UA
		}
		if err := s.db.Model(&existing).Updates(updates).Error; err != nil {
			return nil, err
		}
		return &existing, nil
	case errors.Is(err, gorm.ErrRecordNotFound):
		sub := &models.WebPushSubscription{
			UserID:   userID,
			Endpoint: in.Endpoint,
			P256dh:   in.P256dh,
			Auth:     in.Auth,
		}
		if in.Label != "" {
			sub.DeviceLabel = &in.Label
		}
		if in.UA != "" {
			sub.UserAgent = &in.UA
		}
		if err := s.db.Create(sub).Error; err != nil {
			return nil, err
		}
		return sub, nil
	default:
		return nil, err
	}
}

// Unsubscribe remove a inscrição de um endpoint (ao desativar avisos no navegador).
func (s *WebPushService) Unsubscribe(endpoint string) error {
	if endpoint == "" {
		return nil
	}
	return s.db.Where("endpoint = ?", endpoint).Delete(&models.WebPushSubscription{}).Error
}
