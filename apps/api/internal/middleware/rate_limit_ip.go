package middleware

import (
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

// IPRateLimiter — sliding window in-memory por IP.
// Adequado para single-instance. Em multi-instância, trocar por Redis.
//
// Thread-safe via sync.Mutex. Garbage collection: na cada request, remove
// timestamps expirados do bucket do IP.
type IPRateLimiter struct {
	mu       sync.Mutex
	window   time.Duration
	limit    int
	buckets  map[string][]time.Time
	lastSwep time.Time
}

func NewIPRateLimiter(limit int, window time.Duration) *IPRateLimiter {
	return &IPRateLimiter{
		window:  window,
		limit:   limit,
		buckets: make(map[string][]time.Time),
	}
}

// Allow retorna true se o IP ainda pode fazer mais uma request.
// false se atingiu o limite. Atualiza o bucket em qualquer caso (apenas se Allow=true).
func (l *IPRateLimiter) Allow(ip string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	cutoff := now.Add(-l.window)

	// Sweep periódico de buckets vazios (a cada 5 min)
	if now.Sub(l.lastSwep) > 5*time.Minute {
		for k, v := range l.buckets {
			pruned := pruneTimestamps(v, cutoff)
			if len(pruned) == 0 {
				delete(l.buckets, k)
			} else {
				l.buckets[k] = pruned
			}
		}
		l.lastSwep = now
	}

	// Bucket do IP atual
	timestamps := pruneTimestamps(l.buckets[ip], cutoff)
	if len(timestamps) >= l.limit {
		l.buckets[ip] = timestamps
		return false
	}
	l.buckets[ip] = append(timestamps, now)
	return true
}

// RetryAfter retorna em quanto tempo o IP poderá tentar de novo.
// Só útil quando Allow retornou false.
func (l *IPRateLimiter) RetryAfter(ip string) time.Duration {
	l.mu.Lock()
	defer l.mu.Unlock()
	timestamps := l.buckets[ip]
	if len(timestamps) == 0 {
		return 0
	}
	oldest := timestamps[0]
	return l.window - time.Since(oldest)
}

func pruneTimestamps(stamps []time.Time, cutoff time.Time) []time.Time {
	out := stamps[:0]
	for _, t := range stamps {
		if t.After(cutoff) {
			out = append(out, t)
		}
	}
	return out
}

// Middleware retorna um handler Fiber que aplica o rate limit com base
// no IP do cliente. Usa X-Forwarded-For se presente (para reverse proxy).
func (l *IPRateLimiter) Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ip := clientIP(c)
		if !l.Allow(ip) {
			retry := l.RetryAfter(ip)
			c.Set("Retry-After", formatRetryAfter(retry))
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error":   "rate limit exceeded",
				"message": "Muitos pedidos. Tente novamente em " + retry.Round(time.Minute).String() + ".",
			})
		}
		return c.Next()
	}
}

// clientIP — extrai IP do cliente respeitando X-Forwarded-For (primeiro hop).
func clientIP(c *fiber.Ctx) string {
	if xff := c.Get("X-Forwarded-For"); xff != "" {
		// Pega o primeiro IP da cadeia (cliente original)
		for i := 0; i < len(xff); i++ {
			if xff[i] == ',' || xff[i] == ' ' {
				if i > 0 {
					return xff[:i]
				}
				break
			}
		}
		return xff
	}
	return c.IP()
}

func formatRetryAfter(d time.Duration) string {
	secs := int64(d.Seconds())
	if secs < 1 {
		secs = 1
	}
	// Retry-After em segundos
	return formatInt(secs)
}

func formatInt(n int64) string {
	if n == 0 {
		return "0"
	}
	neg := n < 0
	if neg {
		n = -n
	}
	var buf [20]byte
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
	}
	if neg {
		i--
		buf[i] = '-'
	}
	return string(buf[i:])
}
