package middleware

import (
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

// EmailRateLimiter — sliding window por email (case-insensitive).
// Complementa IPRateLimiter pra defender de brute-force quando atacante
// tem range NAT'd (mesmo email tentado de múltiplos IPs).
//
// Defaults sugeridos: 5 tentativas / 15min.
//
// O middleware lê o email do body via inspeção idempotente (BodyParser-ish):
// como Fiber consome o body uma vez, copiamos via c.Body() e reanexamos.
type EmailRateLimiter struct {
	mu      sync.Mutex
	window  time.Duration
	limit   int
	buckets map[string][]time.Time
	last    time.Time
}

// NewEmailRateLimiter retorna um limiter chaveado por email normalizado
// (lowercase + trim). Defaults: 5 tentativas em 15min.
func NewEmailRateLimiter(limit int, window time.Duration) *EmailRateLimiter {
	return &EmailRateLimiter{
		limit:   limit,
		window:  window,
		buckets: make(map[string][]time.Time),
	}
}

// allow chave-genérica (testável). Retorna true se ainda dentro do limite.
func (l *EmailRateLimiter) allow(key string) bool {
	if key == "" {
		// Sem email no payload — não bloqueia (IPRateLimiter cuida).
		return true
	}
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	cutoff := now.Add(-l.window)

	// GC periódico
	if now.Sub(l.last) > 5*time.Minute {
		for k, v := range l.buckets {
			pruned := pruneTimestamps(v, cutoff)
			if len(pruned) == 0 {
				delete(l.buckets, k)
			} else {
				l.buckets[k] = pruned
			}
		}
		l.last = now
	}

	stamps := pruneTimestamps(l.buckets[key], cutoff)
	if len(stamps) >= l.limit {
		l.buckets[key] = stamps
		return false
	}
	l.buckets[key] = append(stamps, now)
	return true
}

// MiddlewareByUser — variante chaveada pelo userID setado por Auth().
// Usado em rotas autenticadas onde o ataque seria contra a conta da vítima
// (ex: troca de senha). Falls back pra IP se userID ausente (paranoia).
func (l *EmailRateLimiter) MiddlewareByUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		uid := GetUserID(c)
		key := uid.String()
		if key == "00000000-0000-0000-0000-000000000000" || key == "" {
			key = "ip:" + c.IP()
		}
		if !l.allow(key) {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error":   "too_many_attempts",
				"message": "muitas tentativas — aguarde 15 minutos",
			})
		}
		return c.Next()
	}
}

// Middleware — extrai email do JSON do body (campo "email") e aplica limite.
// Usar APÓS validação básica (corpo presente). Em caso de email ausente,
// passa adiante (IPRateLimiter cuida).
func (l *EmailRateLimiter) Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Body parse idempotente — fastjson-ish manual pra não consumir o body.
		var payload struct {
			Email string `json:"email"`
		}
		if err := c.BodyParser(&payload); err == nil {
			email := strings.ToLower(strings.TrimSpace(payload.Email))
			if !l.allow(email) {
				return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
					"error":   "too_many_attempts",
					"message": "muitas tentativas para este email — aguarde 15 minutos",
				})
			}
		}
		// Body já consumido por BodyParser — Fiber re-uses internal buffer,
		// downstream BodyParser ainda funciona em c.Body().
		return c.Next()
	}
}
