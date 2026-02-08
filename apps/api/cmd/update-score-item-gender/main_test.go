package main

import (
	"testing"
)

func TestDetectGender(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Casos masculinos
		{
			name:     "Homem singular minúsculo",
			input:    "Hemoglobina - homem",
			expected: "male",
		},
		{
			name:     "Homens plural maiúsculo",
			input:    "Hemoglobina - HOMENS",
			expected: "male",
		},
		{
			name:     "Masculino explícito",
			input:    "Testosterona (sexo masculino)",
			expected: "male",
		},
		{
			name:     "Para homens",
			input:    "Níveis normais para homens adultos",
			expected: "male",
		},
		{
			name:     "No homem",
			input:    "Próstata no homem",
			expected: "male",
		},

		// Casos femininos
		{
			name:     "Mulher singular minúsculo",
			input:    "Estradiol - mulher",
			expected: "female",
		},
		{
			name:     "Mulheres plural maiúsculo",
			input:    "Progesterona - MULHERES",
			expected: "female",
		},
		{
			name:     "Feminino explícito",
			input:    "Níveis hormonais (sexo feminino)",
			expected: "female",
		},
		{
			name:     "Para mulheres",
			input:    "Valores de referência para mulheres",
			expected: "female",
		},
		{
			name:     "Na mulher",
			input:    "Ciclo menstrual na mulher",
			expected: "female",
		},

		// Casos não aplicáveis
		{
			name:     "Sem indicação de gênero",
			input:    "Glicose em jejum",
			expected: "not_applicable",
		},
		{
			name:     "Hemoglobina glicada sem gênero",
			input:    "HbA1c",
			expected: "not_applicable",
		},
		{
			name:     "Colesterol total",
			input:    "Colesterol Total",
			expected: "not_applicable",
		},
		{
			name:     "TSH",
			input:    "Hormônio Tireoestimulante (TSH)",
			expected: "not_applicable",
		},

		// Casos edge
		{
			name:     "Texto vazio",
			input:    "",
			expected: "not_applicable",
		},
		{
			name:     "Case misto",
			input:    "Hemoglobina - HoMeM",
			expected: "male",
		},
		{
			name:     "Com acentos",
			input:    "Testosterona em homens adultos",
			expected: "male",
		},

		// Casos reais do Plenya
		{
			name:     "Hemoglobina - Homens (real)",
			input:    "Hemoglobina - Homens",
			expected: "male",
		},
		{
			name:     "Hemoglobina - Mulheres (real)",
			input:    "Hemoglobina - Mulheres",
			expected: "female",
		},
		{
			name:     "FEVE sem gênero",
			input:    "FEVE (Fração de Ejeção do Ventrículo Esquerdo)",
			expected: "not_applicable",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := detectGender(tt.input)
			if result != tt.expected {
				t.Errorf("detectGender(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

func BenchmarkDetectGender(b *testing.B) {
	inputs := []string{
		"Hemoglobina - Homens",
		"Hemoglobina - Mulheres",
		"Glicose em jejum",
		"Testosterona (sexo masculino)",
		"Progesterona - mulheres",
		"Colesterol Total",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, input := range inputs {
			detectGender(input)
		}
	}
}
