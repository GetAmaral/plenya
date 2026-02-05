package services

import (
	"strings"
	"unicode"
)

// sanitizeCPF remove tudo que não é número do CPF
// "123.456.789-00" -> "12345678900"
func sanitizeCPF(cpf string) string {
	var result strings.Builder
	for _, r := range cpf {
		if unicode.IsDigit(r) {
			result.WriteRune(r)
		}
	}
	return result.String()
}
