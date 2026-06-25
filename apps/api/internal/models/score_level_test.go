package models

import "testing"

func sptr(s string) *string { return &s }

// TestEvaluatesTrue cobre todos os operadores. Caso-chave: "<" e "<=" guardam o limite
// no UpperLimit (convenção dos dados); a regressão era ler o LowerLimit (vazio) → nunca true.
func TestEvaluatesTrue(t *testing.T) {
	cases := []struct {
		name  string
		op    string
		lower *string
		upper *string
		value float64
		want  bool
	}{
		// "<" com limite no UpperLimit (regressão: LDL <70, Colesterol <120)
		{"< upper: abaixo casa", "<", nil, sptr("120"), 111, true},
		{"< upper: no limite não casa", "<", nil, sptr("120"), 120, false},
		{"< upper: acima não casa", "<", nil, sptr("120"), 130, false},
		// "<=" com limite no UpperLimit
		{"<= upper: no limite casa", "<=", nil, sptr("70"), 70, true},
		{"<= upper: acima não casa", "<=", nil, sptr("70"), 71, false},
		// fallback defensivo: "<" só com LowerLimit
		{"< fallback lower", "<", sptr("5"), nil, 4, true},
		{"< fallback lower não casa", "<", sptr("5"), nil, 5, false},
		// ">" e ">=" com limite no LowerLimit (continuam funcionando)
		{"> lower: acima casa", ">", sptr("120"), nil, 130, true},
		{"> lower: no limite não casa", ">", sptr("120"), nil, 120, false},
		{">= lower: no limite casa", ">=", sptr("60"), nil, 60, true},
		{">= lower: abaixo não casa", ">=", sptr("60"), nil, 59, false},
		// between inclusivo
		{"between dentro", "between", sptr("120"), sptr("159"), 130, true},
		{"between borda inferior", "between", sptr("120"), sptr("159"), 120, true},
		{"between borda superior", "between", sptr("120"), sptr("159"), 159, true},
		{"between fora acima", "between", sptr("120"), sptr("159"), 160, false},
		{"between fora abaixo", "between", sptr("120"), sptr("159"), 119, false},
		// "="
		{"= casa", "=", sptr("5"), nil, 5, true},
		{"= não casa", "=", sptr("5"), nil, 6, false},
		// limites ausentes não estouram
		{"< sem limite", "<", nil, nil, 10, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			sl := ScoreLevel{Operator: c.op, LowerLimit: c.lower, UpperLimit: c.upper}
			if got := sl.EvaluatesTrue(c.value); got != c.want {
				t.Errorf("op=%q lower=%v upper=%v value=%v: got %v, want %v",
					c.op, c.lower, c.upper, c.value, got, c.want)
			}
		})
	}
}
