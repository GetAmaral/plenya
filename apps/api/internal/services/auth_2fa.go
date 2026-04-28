package services

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"net/url"
	"strings"
	"time"
)

// Implementação minimalista RFC 6238 (TOTP) e RFC 4226 (HOTP) sem dependência
// externa. Usa SHA-1 (compatível com Google Authenticator/Authy/1Password) e
// 6 dígitos com janela de 30s. Aceita ±1 step de drift (90s window total).
//
// Decisão: implementação inline pra não puxar pquerna/otp + go.mod tidy num
// patch de segurança. Se precisar de funcionalidades avançadas (backup codes,
// counter-based), trocar pra biblioteca depois.

const (
	totpDigits   = 6
	totpStepSecs = 30
)

// GenerateTOTPSecret — gera secret aleatório base32 (160 bits / 32 chars
// base32 sem padding). Compatível com app autenticador padrão.
func GenerateTOTPSecret() (string, error) {
	b := make([]byte, 20) // 160 bits
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	enc := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(b)
	return enc, nil
}

// BuildOTPAuthURL — monta otpauth://totp/... pra QR code do app.
// issuer aparece no app antes do email. accountName é geralmente o email.
func BuildOTPAuthURL(issuer, accountName, secret string) string {
	label := url.PathEscape(issuer + ":" + accountName)
	q := url.Values{}
	q.Set("secret", secret)
	q.Set("issuer", issuer)
	q.Set("algorithm", "SHA1")
	q.Set("digits", "6")
	q.Set("period", "30")
	return fmt.Sprintf("otpauth://totp/%s?%s", label, q.Encode())
}

// ValidateTOTP — valida code de 6 dígitos com tolerância de ±1 step.
// Retorna true se code bateu em algum dos steps [now-1, now, now+1].
func ValidateTOTP(secret, code string) bool {
	code = strings.TrimSpace(code)
	if len(code) != totpDigits {
		return false
	}
	key, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(strings.ToUpper(secret))
	if err != nil {
		return false
	}
	now := time.Now().Unix() / totpStepSecs
	for _, drift := range []int64{-1, 0, 1} {
		expected := generateHOTP(key, uint64(now+drift))
		if hmac.Equal([]byte(expected), []byte(code)) {
			return true
		}
	}
	return false
}

// generateHOTP — RFC 4226. counter é o step (now/30 pra TOTP).
func generateHOTP(key []byte, counter uint64) string {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, counter)

	mac := hmac.New(sha1.New, key)
	mac.Write(buf)
	sum := mac.Sum(nil)

	offset := sum[len(sum)-1] & 0x0F
	bin := (uint32(sum[offset])&0x7F)<<24 |
		uint32(sum[offset+1])<<16 |
		uint32(sum[offset+2])<<8 |
		uint32(sum[offset+3])

	mod := uint32(1)
	for i := 0; i < totpDigits; i++ {
		mod *= 10
	}
	code := bin % mod
	return fmt.Sprintf("%0*d", totpDigits, code)
}
