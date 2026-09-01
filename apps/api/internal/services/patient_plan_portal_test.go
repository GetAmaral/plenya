package services

import (
	"os"
	"strings"
	"testing"
)

// O portal do paciente só pode ver plano PUBLICADO, e o filtro tem que estar na CONSULTA.
//
// A tentação é listar tudo e filtrar depois, no handler ou no front. Aqui não: um plano em rascunho
// é a leitura clínica ainda sendo escrita — meio texto, hipótese que vai cair, número que ainda vai
// mudar. O paciente lendo isso é pior do que não ter tela nenhuma.
//
// Não há harness de banco nos testes deste pacote, então a garantia possível é estrutural: se
// alguém tirar o filtro da cláusula Where, este teste cai antes de o rascunho chegar à tela do
// paciente.
func TestConsultaDoPortalFiltraStatusNaPropriaQuery(t *testing.T) {
	src, err := os.ReadFile("patient_plan_service.go")
	if err != nil {
		t.Fatalf("lendo a fonte: %v", err)
	}

	for _, fn := range []string{
		"func (s *PatientPlanService) ListPublished(",
		"func (s *PatientPlanService) GetPublished(",
	} {
		corpo := corpoDaFuncao(t, string(src), fn)
		if !strings.Contains(corpo, "status = ?") || !strings.Contains(corpo, "models.PatientPlanPublished") {
			t.Errorf("%s tem que filtrar por status publicado na consulta:\n%s", fn, corpo)
		}
		if !strings.Contains(corpo, "patient_id = ?") {
			t.Errorf("%s tem que escopar pelo paciente na consulta:\n%s", fn, corpo)
		}
	}

	// Rascunho responde "não encontrado", nunca "existe mas você não pode ver": quem sabe o id não
	// pode descobrir, pela resposta diferente, que há um plano sendo escrito.
	get := corpoDaFuncao(t, string(src), "func (s *PatientPlanService) GetPublished(")
	if !strings.Contains(get, "ErrPatientPlanNotFound") {
		t.Errorf("GetPublished tem que devolver 'não encontrado' para rascunho:\n%s", get)
	}
}

// corpoDaFuncao devolve o texto entre a assinatura e a próxima declaração de função.
func corpoDaFuncao(t *testing.T, src, assinatura string) string {
	t.Helper()
	i := strings.Index(src, assinatura)
	if i < 0 {
		t.Fatalf("função não encontrada: %s", assinatura)
	}
	rest := src[i+len(assinatura):]
	if j := strings.Index(rest, "\nfunc "); j >= 0 {
		return rest[:j]
	}
	return rest
}
