// Package crypto — blind index helpers (M4).
//
// Blind index = HMAC-SHA256(plaintext, blindKey). Permite buscar por valor
// criptografado sem descriptografar (que requer ler todas as rows). Casos:
//
//   - Patient.CPFBlindIndex: PatientService.GetByCPF lookups O(log N) via
//     `WHERE cpf_blind_index = ?` em vez de descriptografar tabela inteira.
//
// Por que HMAC e não SHA-256 cru: SHA-256 plain de CPF é trivialmente
// reversível por força bruta (10^11 CPFs válidos, GPU rainbow table). HMAC
// com chave secreta torna a brute-force inviável sem a chave.
//
// A chave é separada da chave de criptografia (defense in depth: vazamento
// de uma não compromete a outra). Em prod: BLIND_INDEX_KEY (32 bytes hex).
// Em dev: cai pra ENCRYPTION_KEY com warning (pra não quebrar local).
package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
)

var blindIndexKey string

// SetBlindIndexKey — chamado no boot. Aceita string vazia em dev (warning logado
// no caller); BlindIndexCPF retorna erro nesse caso.
func SetBlindIndexKey(key string) {
	blindIndexKey = key
}

// GetBlindIndexKey — exposto pra testes.
func GetBlindIndexKey() string { return blindIndexKey }

// BlindIndexCPF — HMAC-SHA256(normalizedCPF, blindKey) → hex.
// Normaliza CPF antes (remove pontuação, lower) pra que "123.456.789-00" e
// "12345678900" gerem o mesmo índice.
func BlindIndexCPF(plain string) (string, error) {
	if blindIndexKey == "" {
		return "", errors.New("blind index key not set — call crypto.SetBlindIndexKey")
	}
	norm := normalizeCPF(plain)
	if norm == "" {
		return "", nil // CPF vazio → índice vazio (não quebra o save)
	}
	mac := hmac.New(sha256.New, []byte(blindIndexKey))
	mac.Write([]byte(norm))
	return hex.EncodeToString(mac.Sum(nil)), nil
}

// MustBlindIndexCPF — wrapper que loga e retorna "" em erro (pra hooks que
// não devem abortar Save por blind index não configurado em dev).
func MustBlindIndexCPF(plain string) string {
	idx, err := BlindIndexCPF(plain)
	if err != nil {
		return ""
	}
	return idx
}

// normalizeCPF — remove pontuação, retorna só dígitos.
func normalizeCPF(s string) string {
	var b strings.Builder
	for _, r := range s {
		if r >= '0' && r <= '9' {
			b.WriteRune(r)
		}
	}
	return b.String()
}
