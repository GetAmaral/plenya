package models

import "strings"

// PlaceholderEmailSuffix — suffixo usado quando criamos User+Patient a partir de
// claim do Escore Light feito SÓ via WhatsApp (sem email real). Esse email não
// recebe notificações reais (boas-vindas, etc) — mantemos pra satisfazer a
// constraint UNIQUE NOT NULL de User.Email.
//
// Formato: "wa-5511999998888@placeholder.plenyasaude.com.br"
//
// Use IsPlaceholderEmail pra checar; NÃO fazer strings.Contains hardcoded em outro
// ponto do código (causa drift se o suffix mudar).
const PlaceholderEmailSuffix = "@placeholder.plenyasaude.com.br"

// IsPlaceholderEmail retorna true se o email é o placeholder gerado pelo CRM.
func IsPlaceholderEmail(email string) bool {
	return strings.HasSuffix(strings.ToLower(strings.TrimSpace(email)), PlaceholderEmailSuffix)
}

// BuildPlaceholderEmail gera o email placeholder a partir de phone E.164 (sem +).
// Ex: "+5511999998888" → "wa-5511999998888@placeholder.plenyasaude.com.br"
func BuildPlaceholderEmail(phone string) string {
	return "wa-" + strings.TrimPrefix(phone, "+") + PlaceholderEmailSuffix
}

// NormalizePhoneBR aceita formatos BR comuns e devolve E.164 (+55XXXXXXXXXXX).
// Retorna "" se inválido (caller decide se mantém input original ou rejeita).
//
// Aceita:
//   - "(11) 99999-9999"     → "+5511999999999"
//   - "11999999999"         → "+5511999999999"
//   - "+5511999999999"      → "+5511999999999" (já normalizado, no-op)
//   - "5511999999999"       → "+5511999999999"
//   - "0055 11 99999-9999"  → "+5511999999999" (IDD brasileiro)
//   - "1133334444"          → "+551133334444"  (fixo, 10 dígitos)
//
// Rejeita (retorna ""):
//   - menos de 10 ou mais de 13 dígitos efetivos (após strip de IDD)
//   - DDI diferente de 55
func NormalizePhoneBR(raw string) string {
	if raw == "" {
		return ""
	}
	digits := stripNonDigits(raw)

	// IDD brasileiro "00 55 ..." — strip os 2 primeiros zeros, deixa começar com 55
	if len(digits) >= 4 && strings.HasPrefix(digits, "0055") {
		digits = digits[2:]
	}

	// Se já vem com DDI 55, descarta os 2 primeiros pra padronizar
	if len(digits) == 12 || len(digits) == 13 {
		if strings.HasPrefix(digits, "55") {
			digits = digits[2:]
		} else {
			return "" // DDI estrangeiro — fora de escopo
		}
	}

	if len(digits) != 10 && len(digits) != 11 {
		return ""
	}

	// DDD válido (11-99) — checagem mínima
	dd := digits[:2]
	if dd[0] < '1' || dd[0] > '9' {
		return ""
	}

	return "+55" + digits
}

func stripNonDigits(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		if r >= '0' && r <= '9' {
			b.WriteRune(r)
		}
	}
	return b.String()
}
