// Package handlers — GoogleOAuthHandler implementa o fluxo OAuth de
// integração Google Calendar (Calendar V1, Bloco B).
//
// Fluxo:
//
//  1. GET /api/v1/integrations/google/auth-url
//     Médico autenticado → retorna URL pro consent screen com state JWT (CSRF).
//
//  2. GET /api/v1/integrations/google/callback?code=X&state=Y
//     SEM middleware Auth (Google redirect não traz JWT).
//     - Valida state JWT → extrai user_id
//     - Troca code por tokens
//     - Cria/upsert CalendarCredential (tokens cifrados)
//     - Cria calendar dedicado "Plenya"
//     - Redireciona pro frontend com ?google_connected=true
//
//  3. POST /api/v1/integrations/google/disconnect
//     Revoga refresh token + marca CalendarCredential.RevokedAt.
//
//  4. GET /api/v1/integrations/google/status
//     Retorna {connected, googleEmail, lastSyncAt}.
package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/crypto"
	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

// stateClaims — JWT curto (5min) usado como CSRF token no fluxo OAuth.
// JWT > random nonce em DB porque é stateless e não polui banco.
type stateClaims struct {
	UserID string `json:"uid"`
	Nonce  string `json:"n"`
	jwt.RegisteredClaims
}

const oauthStateTTL = 5 * time.Minute

type GoogleOAuthHandler struct {
	cfg       *config.Config
	googleSvc *services.GoogleCalendarService
	db        *gorm.DB
}

func NewGoogleOAuthHandler(cfg *config.Config, googleSvc *services.GoogleCalendarService, db *gorm.DB) *GoogleOAuthHandler {
	return &GoogleOAuthHandler{
		cfg:       cfg,
		googleSvc: googleSvc,
		db:        db,
	}
}

// AuthURL retorna a URL de consent OAuth Google + JWT state.
//
// GET /api/v1/integrations/google/auth-url
// Auth: qualquer usuário autenticado (médico tipicamente).
// Resp: {url: "https://accounts.google.com/o/oauth2/v2/auth?..."}
func (h *GoogleOAuthHandler) AuthURL(c *fiber.Ctx) error {
	if !h.googleSvc.IsConfigured() {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"error":   "google_not_configured",
			"message": "GOOGLE_CLIENT_ID/SECRET ausentes — setup operacional pendente",
		})
	}

	userID := middleware.GetUserID(c)

	// Gera nonce aleatório.
	nonceBytes := make([]byte, 16)
	if _, err := rand.Read(nonceBytes); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "nonce_generation_failed"})
	}
	nonce := hex.EncodeToString(nonceBytes)

	// Assina JWT state com mesmo secret JWT do app (HS256).
	claims := stateClaims{
		UserID: userID.String(),
		Nonce:  nonce,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(oauthStateTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "plenya-google-oauth",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(h.cfg.JWT.Secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "state_sign_failed"})
	}

	authURL := h.googleSvc.BuildAuthURL(signed)
	return c.JSON(fiber.Map{"url": authURL})
}

// Callback processa o redirect Google. Sem middleware Auth (state JWT valida).
//
// GET /api/v1/integrations/google/callback?code=X&state=Y
//
// Em sucesso, redireciona pro frontend:
//   {SITE_PUBLIC_URL}/configuracoes/integracoes?google_connected=true
//
// Em erro, redireciona com ?google_error=<reason>
func (h *GoogleOAuthHandler) Callback(c *fiber.Ctx) error {
	if !h.googleSvc.IsConfigured() {
		return h.redirectError(c, "not_configured")
	}

	// Google pode passar `error` em vez de `code` (usuário recusou consent).
	if errParam := c.Query("error"); errParam != "" {
		log.Printf("⚠️  Google OAuth callback denied: %s", errParam)
		return h.redirectError(c, "consent_denied")
	}

	code := c.Query("code")
	state := c.Query("state")
	if code == "" || state == "" {
		return h.redirectError(c, "missing_params")
	}

	// Valida state JWT.
	parsed, err := jwt.ParseWithClaims(state, &stateClaims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(h.cfg.JWT.Secret), nil
	})
	if err != nil || !parsed.Valid {
		log.Printf("⚠️  Google OAuth: invalid state JWT: %v", err)
		return h.redirectError(c, "invalid_state")
	}
	claims, ok := parsed.Claims.(*stateClaims)
	if !ok {
		return h.redirectError(c, "invalid_state_claims")
	}
	userID, err := uuid.Parse(claims.UserID)
	if err != nil {
		return h.redirectError(c, "invalid_user_id")
	}

	ctx := c.Context()

	// Troca code por tokens.
	tokens, err := h.googleSvc.ExchangeCode(ctx, code)
	if err != nil {
		log.Printf("⚠️  Google OAuth exchange failed: %v", err)
		return h.redirectError(c, "exchange_failed")
	}

	// Pega email da conta Google.
	email, err := h.googleSvc.GetUserEmail(ctx, tokens.AccessToken)
	if err != nil {
		log.Printf("⚠️  Google OAuth userinfo failed: %v", err)
		return h.redirectError(c, "userinfo_failed")
	}

	// Cria calendar dedicado.
	calendarID, err := h.googleSvc.CreateDedicatedCalendar(ctx, tokens.AccessToken)
	if err != nil {
		log.Printf("⚠️  Google OAuth create calendar failed: %v", err)
		return h.redirectError(c, "create_calendar_failed")
	}

	// Cifra tokens antes de persistir.
	encAccess, err := crypto.EncryptWithDefaultKey(tokens.AccessToken)
	if err != nil {
		return h.redirectError(c, "encrypt_failed")
	}
	encRefresh, err := crypto.EncryptWithDefaultKey(tokens.RefreshToken)
	if err != nil {
		return h.redirectError(c, "encrypt_failed")
	}

	// Upsert CalendarCredential — se médico já tinha credential revogada,
	// reaproveita a row (mantém ID estável).
	var cred models.CalendarCredential
	now := time.Now()
	res := h.db.WithContext(ctx).
		Where("user_id = ? AND provider = ?", userID, models.CalendarProviderGoogle).
		First(&cred)
	if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		log.Printf("⚠️  Google OAuth lookup credential: %v", res.Error)
		return h.redirectError(c, "db_lookup_failed")
	}

	cred.UserID = userID
	cred.Provider = string(models.CalendarProviderGoogle)
	cred.GoogleAccountEmail = email
	cred.EncryptedAccessToken = encAccess
	cred.EncryptedRefreshToken = encRefresh
	cred.AccessTokenExpiresAt = tokens.ExpiresAt
	cred.DedicatedCalendarID = calendarID
	cred.RevokedAt = nil
	cred.LastSyncAt = &now

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		if err := h.db.WithContext(ctx).Create(&cred).Error; err != nil {
			log.Printf("⚠️  Google OAuth create credential: %v", err)
			return h.redirectError(c, "db_create_failed")
		}
	} else {
		if err := h.db.WithContext(ctx).Save(&cred).Error; err != nil {
			log.Printf("⚠️  Google OAuth save credential: %v", err)
			return h.redirectError(c, "db_save_failed")
		}
	}

	log.Printf("✅ Google Calendar conectado para user %s (email=%s, calendarID=%s)", userID, email, calendarID)

	return c.Redirect(h.frontendURL("/configuracoes/integracoes?google_connected=true"))
}

// Disconnect revoga o token + marca CalendarCredential.RevokedAt.
//
// POST /api/v1/integrations/google/disconnect
// Auth: usuário dono da credential.
func (h *GoogleOAuthHandler) Disconnect(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)
	ctx := c.Context()

	var cred models.CalendarCredential
	err := h.db.WithContext(ctx).
		Where("user_id = ? AND provider = ?", userID, models.CalendarProviderGoogle).
		First(&cred).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "not_connected",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Tenta revogar no Google (best effort — se falhar, ainda marca local como revogado).
	if cred.IsActive() {
		refresh, decErr := crypto.DecryptWithDefaultKey(cred.EncryptedRefreshToken)
		if decErr == nil {
			if revErr := h.googleSvc.RevokeToken(ctx, refresh); revErr != nil {
				log.Printf("⚠️  Google revoke (best-effort) failed for user %s: %v", userID, revErr)
			}
		} else {
			log.Printf("⚠️  Google revoke skipped: decrypt refresh failed: %v", decErr)
		}
	}

	now := time.Now()
	cred.RevokedAt = &now
	if err := h.db.WithContext(ctx).Save(&cred).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"disconnected": true})
}

// Status retorna o status da credential do usuário corrente.
//
// GET /api/v1/integrations/google/status
// Auth: usuário corrente.
// Resp: {connected: bool, googleEmail?, lastSyncAt?, revokedAt?}
func (h *GoogleOAuthHandler) Status(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)
	ctx := c.Context()

	var cred models.CalendarCredential
	err := h.db.WithContext(ctx).
		Where("user_id = ? AND provider = ?", userID, models.CalendarProviderGoogle).
		First(&cred).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(fiber.Map{
				"connected":  false,
				"configured": h.googleSvc.IsConfigured(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"connected":    cred.IsActive(),
		"configured":   h.googleSvc.IsConfigured(),
		"googleEmail":  cred.GoogleAccountEmail,
		"lastSyncAt":   cred.LastSyncAt,
		"revokedAt":    cred.RevokedAt,
		"calendarId":   cred.DedicatedCalendarID,
	})
}

// redirectError monta URL de erro pro frontend e redireciona.
// Permite UI mostrar toast contextual com a razão.
func (h *GoogleOAuthHandler) redirectError(c *fiber.Ctx, reason string) error {
	target := fmt.Sprintf("/configuracoes/integracoes?google_error=%s", url.QueryEscape(reason))
	return c.Redirect(h.frontendURL(target))
}

// frontendURL prefixa SITE_PUBLIC_URL (CRM_ADMIN_URL fallback).
func (h *GoogleOAuthHandler) frontendURL(path string) string {
	base := h.cfg.CRM.AdminURL
	if base == "" {
		base = h.cfg.Site.PublicURL
	}
	return base + path
}
