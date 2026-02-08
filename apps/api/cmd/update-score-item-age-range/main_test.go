package main

import (
	"testing"
)

func TestDetectAgeRange(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectMin   *int
		expectMax   *int
		expectFound bool
	}{
		// Range "X-Y anos" ou "X a Y anos"
		{
			name:        "Range 18-65 anos",
			input:       "Hemoglobina - 18-65 anos",
			expectMin:   intPtr(18),
			expectMax:   intPtr(65),
			expectFound: true,
		},
		{
			name:        "Range 18 a 65 anos",
			input:       "Colesterol - 18 a 65 anos",
			expectMin:   intPtr(18),
			expectMax:   intPtr(65),
			expectFound: true,
		},
		{
			name:        "Range maiúsculo",
			input:       "TSH - 20-60 ANOS",
			expectMin:   intPtr(20),
			expectMax:   intPtr(60),
			expectFound: true,
		},
		{
			name:        "Range sem espaços",
			input:       "Glicose 30-70anos",
			expectMin:   intPtr(30),
			expectMax:   intPtr(70),
			expectFound: true,
		},

		// Range entre parênteses "(X-Y anos)"
		{
			name:        "Range em parênteses",
			input:       "Testosterona (18-50 anos)",
			expectMin:   intPtr(18),
			expectMax:   intPtr(50),
			expectFound: true,
		},
		{
			name:        "Range em parênteses com 'a'",
			input:       "Estradiol (20 a 45 anos)",
			expectMin:   intPtr(20),
			expectMax:   intPtr(45),
			expectFound: true,
		},

		// Padrão "< X anos"
		{
			name:        "Menor que 18",
			input:       "Hemoglobina < 18 anos",
			expectMin:   intPtr(0),
			expectMax:   intPtr(18),
			expectFound: true,
		},
		{
			name:        "Menor que sem espaço",
			input:       "Glicose <18 anos",
			expectMin:   intPtr(0),
			expectMax:   intPtr(18),
			expectFound: true,
		},
		{
			name:        "Menor que por extenso",
			input:       "Colesterol menor que 20 anos",
			expectMin:   intPtr(0),
			expectMax:   intPtr(20),
			expectFound: true,
		},
		{
			name:        "Abaixo de",
			input:       "TSH abaixo de 15 anos",
			expectMin:   intPtr(0),
			expectMax:   intPtr(15),
			expectFound: true,
		},

		// Padrão "> X anos"
		{
			name:        "Maior que 65",
			input:       "Hemoglobina > 65 anos",
			expectMin:   intPtr(65),
			expectMax:   intPtr(150),
			expectFound: true,
		},
		{
			name:        "Maior que sem espaço",
			input:       "PSA >50 anos",
			expectMin:   intPtr(50),
			expectMax:   intPtr(150),
			expectFound: true,
		},
		{
			name:        "Maior que por extenso",
			input:       "Colesterol maior que 60 anos",
			expectMin:   intPtr(60),
			expectMax:   intPtr(150),
			expectFound: true,
		},
		{
			name:        "Acima de",
			input:       "Glicose acima de 70 anos",
			expectMin:   intPtr(70),
			expectMax:   intPtr(150),
			expectFound: true,
		},

		// Padrão "X+ anos"
		{
			name:        "18+ anos",
			input:       "Testosterona 18+ anos",
			expectMin:   intPtr(18),
			expectMax:   intPtr(150),
			expectFound: true,
		},
		{
			name:        "65 +",
			input:       "PSA 65 + anos",
			expectMin:   intPtr(65),
			expectMax:   intPtr(150),
			expectFound: true,
		},
		{
			name:        "50+",
			input:       "Colonoscopia 50+",
			expectMin:   intPtr(50),
			expectMax:   intPtr(150),
			expectFound: true,
		},

		// Casos que DEVEM ser IGNORADOS
		{
			name:        "Pré-menopausa (ignorar)",
			input:       "Estradiol - Pré-menopausa",
			expectMin:   nil,
			expectMax:   nil,
			expectFound: false,
		},
		{
			name:        "Pós-menopausa (ignorar)",
			input:       "FSH - Pós-menopausa",
			expectMin:   nil,
			expectMax:   nil,
			expectFound: false,
		},
		{
			name:        "Gestação (ignorar)",
			input:       "hCG - Gestação",
			expectMin:   nil,
			expectMax:   nil,
			expectFound: false,
		},
		{
			name:        "Gestante (ignorar)",
			input:       "TSH - Gestante",
			expectMin:   nil,
			expectMax:   nil,
			expectFound: false,
		},
		{
			name:        "Fase folicular (ignorar)",
			input:       "Estradiol - Fase folicular",
			expectMin:   nil,
			expectMax:   nil,
			expectFound: false,
		},
		{
			name:        "Lactante (ignorar)",
			input:       "Prolactina - Lactante",
			expectMin:   nil,
			expectMax:   nil,
			expectFound: false,
		},

		// Casos sem range detectável
		{
			name:        "Sem indicação de idade",
			input:       "Glicose em jejum",
			expectMin:   nil,
			expectMax:   nil,
			expectFound: false,
		},
		{
			name:        "Hemoglobina geral",
			input:       "Hemoglobina",
			expectMin:   nil,
			expectMax:   nil,
			expectFound: false,
		},
		{
			name:        "Colesterol total",
			input:       "Colesterol Total",
			expectMin:   nil,
			expectMax:   nil,
			expectFound: false,
		},

		// Edge cases
		{
			name:        "Range inválido (min > max)",
			input:       "Teste 80-20 anos",
			expectMin:   nil,
			expectMax:   nil,
			expectFound: false,
		},
		{
			name:        "Idade > 150 (inválido)",
			input:       "Teste 200-300 anos",
			expectMin:   nil,
			expectMax:   nil,
			expectFound: false,
		},
		{
			name:        "String vazia",
			input:       "",
			expectMin:   nil,
			expectMax:   nil,
			expectFound: false,
		},

		// Casos reais do Plenya
		{
			name:        "Hemoglobina Homens (real)",
			input:       "Hemoglobina - Homens",
			expectMin:   nil,
			expectMax:   nil,
			expectFound: false,
		},
		{
			name:        "Adultos 20-60 anos (real)",
			input:       "Colesterol LDL - Adultos 20-60 anos",
			expectMin:   intPtr(20),
			expectMax:   intPtr(60),
			expectFound: true,
		},
		{
			name:        "Idosos 65+ (real)",
			input:       "PSA - Homens 65+",
			expectMin:   intPtr(65),
			expectMax:   intPtr(150),
			expectFound: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := detectAgeRange(tt.input)

			// Verificar se foi detectado
			if result.Detected != tt.expectFound {
				t.Errorf("detectAgeRange(%q).Detected = %v, expected %v",
					tt.input, result.Detected, tt.expectFound)
			}

			// Verificar Min
			if !ptrEqual(result.Min, tt.expectMin) {
				t.Errorf("detectAgeRange(%q).Min = %v, expected %v",
					tt.input, ptrToString(result.Min), ptrToString(tt.expectMin))
			}

			// Verificar Max
			if !ptrEqual(result.Max, tt.expectMax) {
				t.Errorf("detectAgeRange(%q).Max = %v, expected %v",
					tt.input, ptrToString(result.Max), ptrToString(tt.expectMax))
			}
		})
	}
}

// Helpers
func intPtr(i int) *int {
	return &i
}

func ptrEqual(a, b *int) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

func ptrToString(p *int) string {
	if p == nil {
		return "nil"
	}
	return string(rune(*p))
}

func BenchmarkDetectAgeRange(b *testing.B) {
	inputs := []string{
		"Hemoglobina - 18-65 anos",
		"Colesterol < 20 anos",
		"PSA > 50 anos",
		"Testosterona 18+ anos",
		"Estradiol - Pré-menopausa",
		"Glicose em jejum",
		"Adultos 20-60 anos",
		"Idosos (65 a 80 anos)",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, input := range inputs {
			detectAgeRange(input)
		}
	}
}
