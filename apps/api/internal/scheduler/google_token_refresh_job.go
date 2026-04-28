// Package scheduler — GoogleTokenRefreshJob (Calendar V1, Bloco E).
//
// Roda a cada 30min checando CalendarCredentials cujo access_token_expires_at é
// menor que now+10min. Pra cada um:
//
//  1. Decifra refresh token cifrado.
//  2. Chama Google OAuth refresh.
//  3. Em sucesso → re-cifra novo access token + atualiza expiry + LastSyncAt.
//  4. Em invalid_grant → marca RevokedAt + cria notificação in-app pro médico
//     ("Reconectar Google Calendar").
//
// Idempotência: se token recém-renovado, não faz nada (filtro by expiry).
//
// LGPD: tokens NUNCA logados. Apenas IDs + status agregados.
package scheduler

import (
	"context"
	"errors"
	"log"
	"time"

	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

const (
	googleRefreshInterval        = 30 * time.Minute
	googleRefreshLeeway          = 10 * time.Minute
	googleRefreshTimeoutPerToken = 30 * time.Second
	googleRefreshTimeoutBatch    = 10 * time.Minute
)

type GoogleTokenRefreshJob struct {
	db        *gorm.DB
	googleSvc *services.GoogleCalendarService
	notifSvc  *services.NotificationService
}

func NewGoogleTokenRefreshJob(
	db *gorm.DB,
	googleSvc *services.GoogleCalendarService,
	notifSvc *services.NotificationService,
) *GoogleTokenRefreshJob {
	return &GoogleTokenRefreshJob{
		db:        db,
		googleSvc: googleSvc,
		notifSvc:  notifSvc,
	}
}

// Run executa uma passada do job.
func (j *GoogleTokenRefreshJob) Run() error {
	if j.googleSvc == nil || !j.googleSvc.IsConfigured() {
		// Sem credenciais Google config — degrade silencioso.
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), googleRefreshTimeoutBatch)
	defer cancel()

	expiryThreshold := time.Now().UTC().Add(googleRefreshLeeway)

	var creds []models.CalendarCredential
	err := j.db.WithContext(ctx).
		Where("provider = ?", models.CalendarProviderGoogle).
		Where("revoked_at IS NULL").
		Where("access_token_expires_at < ?", expiryThreshold).
		Find(&creds).Error
	if err != nil {
		log.Printf("⚠️  [GOOGLE REFRESH] query: %v", err)
		return err
	}
	if len(creds) == 0 {
		return nil
	}

	log.Printf("🔁 [GOOGLE REFRESH] %d credentials a renovar", len(creds))
	for i := range creds {
		j.refreshOne(ctx, &creds[i])
	}
	return nil
}

// refreshOne processa uma credencial. Erros são logados (sem token leak).
func (j *GoogleTokenRefreshJob) refreshOne(ctx context.Context, cred *models.CalendarCredential) {
	subCtx, cancel := context.WithTimeout(ctx, googleRefreshTimeoutPerToken)
	defer cancel()

	// M8 — usa OAuth key dedicada (com fallback retrocompat pra default key).
	refreshToken, err := j.googleSvc.DecryptOAuthToken(cred.EncryptedRefreshToken)
	if err != nil {
		log.Printf("⚠️  [GOOGLE REFRESH] decrypt cred=%s: %v", cred.ID, err)
		return
	}

	tokens, err := j.googleSvc.RefreshAccessToken(subCtx, refreshToken)
	if err != nil {
		if errors.Is(err, services.ErrCredentialRevoked) {
			j.handleRevoked(subCtx, cred)
			return
		}
		log.Printf("⚠️  [GOOGLE REFRESH] refresh cred=%s user=%s: %v", cred.ID, cred.UserID, err)
		return
	}

	encAccess, err := j.googleSvc.EncryptOAuthToken(tokens.AccessToken)
	if err != nil {
		log.Printf("⚠️  [GOOGLE REFRESH] encrypt access cred=%s: %v", cred.ID, err)
		return
	}
	updates := map[string]any{
		"encrypted_access_token":  encAccess,
		"access_token_expires_at": tokens.ExpiresAt,
		"last_sync_at":            time.Now().UTC(),
	}
	// Google só devolve refresh_token na 1a vez — preserva o anterior.
	if tokens.RefreshToken != "" && tokens.RefreshToken != refreshToken {
		encRefresh, err := j.googleSvc.EncryptOAuthToken(tokens.RefreshToken)
		if err == nil {
			updates["encrypted_refresh_token"] = encRefresh
		}
	}
	if err := j.db.WithContext(subCtx).Model(cred).Updates(updates).Error; err != nil {
		log.Printf("⚠️  [GOOGLE REFRESH] persist cred=%s: %v", cred.ID, err)
	}
}

// handleRevoked marca credencial como revogada e notifica o médico.
func (j *GoogleTokenRefreshJob) handleRevoked(ctx context.Context, cred *models.CalendarCredential) {
	now := time.Now().UTC()
	if err := j.db.WithContext(ctx).Model(cred).Update("revoked_at", now).Error; err != nil {
		log.Printf("⚠️  [GOOGLE REFRESH] mark revoked cred=%s: %v", cred.ID, err)
	}
	log.Printf("🚫 [GOOGLE REFRESH] cred=%s user=%s marcada como revogada", cred.ID, cred.UserID)

	// Notificação in-app pro médico. Reusa enum NotificationLeadAssigned como
	// melhor encaixe atual ("acionável, com action URL"). V2 cria type específico.
	if j.notifSvc == nil {
		return
	}
	err := j.notifSvc.CreateConversationNotification(
		cred.UserID,
		nil, nil,
		models.NotificationLeadAssigned,
		"Reconectar Google Calendar",
		"Sua conexão com o Google expirou. Vá em Configurações > Integrações para reconectar.",
		"/configuracoes/integracoes",
	)
	if err != nil {
		log.Printf("⚠️  [GOOGLE REFRESH] notif cred=%s: %v", cred.ID, err)
	}
}

// Start arranca o ticker em goroutine.
func (j *GoogleTokenRefreshJob) Start() {
	log.Printf("🔁 [GOOGLE REFRESH] iniciado (interval=%s, leeway=%s)", googleRefreshInterval, googleRefreshLeeway)
	go func() {
		// Run imediatamente (catch-up após restart).
		if err := j.Run(); err != nil {
			log.Printf("⚠️  [GOOGLE REFRESH] erro inicial: %v", err)
		}
		ticker := time.NewTicker(googleRefreshInterval)
		defer ticker.Stop()
		for range ticker.C {
			if err := j.Run(); err != nil {
				log.Printf("⚠️  [GOOGLE REFRESH] erro: %v", err)
			}
		}
	}()
}
