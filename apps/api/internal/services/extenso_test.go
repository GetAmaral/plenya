package services

import "testing"

func TestValorPorExtenso(t *testing.T) {
	cases := map[int64]string{
		85000:  "oitocentos e cinquenta reais",
		100:    "um real",
		120000: "mil e duzentos reais",
		125050: "mil duzentos e cinquenta reais e cinquenta centavos",
		35000:  "trezentos e cinquenta reais",
		50:     "cinquenta centavos",
	}
	for cents, want := range cases {
		if got := valorPorExtenso(cents); got != want {
			t.Errorf("valorPorExtenso(%d) = %q; quero %q", cents, got, want)
		}
	}
}
