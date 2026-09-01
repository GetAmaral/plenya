package services

import (
	"encoding/json"
	"testing"
)

// caminhosDasOps é a metade extratora de AITouchedPaths. A outra metade precisa de banco; esta não,
// e é onde moram as decisões que não são óbvias.
func TestCaminhosDasOps(t *testing.T) {
	casos := []struct {
		nome     string
		ops      string
		esperado []string
	}{
		{
			nome:     "op aplicada conta",
			ops:      `[{"op":{"op":"edit","slideId":"s1","path":"title"},"decision":"aplicar"}]`,
			esperado: []string{"s1:title"},
		},
		{
			// Uma op que virou sugestão não escreveu nada no conteúdo. Contá-la marcaria como
			// "gerado por IA" um trecho que a IA propôs e o médico nunca aceitou.
			nome:     "op que virou sugestao nao conta",
			ops:      `[{"op":{"op":"edit","slideId":"s1","path":"summary.lines[0].value"},"decision":"sugerir"}]`,
			esperado: nil,
		},
		{
			nome:     "op recusada nao conta",
			ops:      `[{"op":{"op":"edit","slideId":"s1","path":"rulers[0].segments"},"decision":"recusar"}]`,
			esperado: nil,
		},
		{
			// Revisão de aceite de sugestão grava a op sem campo `decision`: ela JÁ é a decisão.
			nome:     "op sem decision conta",
			ops:      `[{"op":{"op":"edit","slideId":"s2","path":"rulers[0].sub"}}]`,
			esperado: []string{"s2:rulers[0].sub"},
		},
		{
			// Op estrutural não tem `path`. Sem o fallback ela sumiria da trilha, e perder um
			// slide em silêncio é pior que um número errado.
			nome:     "op estrutural usa o verbo no lugar do caminho",
			ops:      `[{"op":{"op":"remove","slideId":"s3"}}]`,
			esperado: []string{"s3:remove"},
		},
	}
	for _, c := range casos {
		t.Run(c.nome, func(t *testing.T) {
			var ops any
			if err := json.Unmarshal([]byte(c.ops), &ops); err != nil {
				t.Fatal(err)
			}
			got := caminhosDasOps(ops)
			if len(got) != len(c.esperado) {
				t.Fatalf("esperava %v, veio %v", c.esperado, got)
			}
			for i := range got {
				if got[i] != c.esperado[i] {
					t.Fatalf("esperava %v, veio %v", c.esperado, got)
				}
			}
		})
	}
}
