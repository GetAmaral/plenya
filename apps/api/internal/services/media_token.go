package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Token de mídia: assinatura HMAC curta e ESCOPADA a UMA mídia de conversa.
//
// Motivo: o <audio>/<video> do navegador não envia header Authorization, e o
// Safari/iOS não toca mídia de blob: URL — precisa de uma URL direta (com Range).
// Em vez de pôr o JWT de acesso na URL (amplo + vaza em logs), emitimos um token
// dedicado que só libera a leitura daquele arquivo, por tempo curto.
//
// Formato: base64url("type:id:activity:exp") + "." + base64url(hmacSHA256).

// SignMediaToken gera o token assinado pra uma mídia específica.
func SignMediaToken(secret, ownerType, ownerID, activityID string, ttl time.Duration) string {
	exp := time.Now().Add(ttl).Unix()
	payload := fmt.Sprintf("%s:%s:%s:%d", ownerType, ownerID, activityID, exp)
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(payload))
	return base64.RawURLEncoding.EncodeToString([]byte(payload)) + "." +
		base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}

// VerifyMediaToken valida assinatura, expiração e escopo (type/id/activity têm que
// bater com a rota). UUIDs não contêm ":", então o split é seguro.
func VerifyMediaToken(secret, token, ownerType, ownerID, activityID string) bool {
	parts := strings.SplitN(token, ".", 2)
	if len(parts) != 2 {
		return false
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return false
	}
	sig, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return false
	}
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	if !hmac.Equal(sig, mac.Sum(nil)) {
		return false
	}
	fields := strings.Split(string(payload), ":")
	if len(fields) != 4 || fields[0] != ownerType || fields[1] != ownerID || fields[2] != activityID {
		return false
	}
	exp, err := strconv.ParseInt(fields[3], 10, 64)
	if err != nil || time.Now().Unix() > exp {
		return false
	}
	return true
}
