package services

import (
	"os"
	"strings"
	"testing"
)

// O portal do paciente só pode ver o que foi PUBLICADO, e o filtro tem que estar na CONSULTA.
//
// A tentação é listar tudo e filtrar depois, no handler ou no front. Aqui não: um plano em rascunho
// é a leitura clínica ainda sendo escrita — meio texto, hipótese que vai cair, número que ainda vai
// mudar. O paciente lendo isso é pior do que não ter tela nenhuma.
//
// O portão é `published_content IS NOT NULL`, não `status`. A diferença importa nos dois sentidos:
// um plano nunca publicado não tem cópia congelada e não aparece; e o médico voltar um plano
// publicado para rascunho, para ajustar uma frase, NÃO pode apagar da tela a devolutiva que o
// paciente já recebeu — foi por isso que o filtro deixou de ser por status.
//
// Não há harness de banco nos testes deste pacote, então a garantia possível é estrutural: se
// alguém tirar o filtro da cláusula Where, este teste cai antes de um rascunho chegar ao paciente.
func TestConsultaDoPortalFiltraOPublicadoNaPropriaQuery(t *testing.T) {
	src, err := os.ReadFile("patient_plan_service.go")
	if err != nil {
		t.Fatalf("lendo a fonte: %v", err)
	}

	for _, fn := range []string{
		"func (s *PatientPlanService) ListPublished(",
		"func (s *PatientPlanService) GetPublished(",
	} {
		corpo := corpoDaFuncao(t, string(src), fn)
		if !strings.Contains(corpo, "published_content IS NOT NULL") {
			t.Errorf("%s tem que exigir conteúdo publicado na consulta:\n%s", fn, corpo)
		}
		if !strings.Contains(corpo, "published_at IS NOT NULL") {
			t.Errorf("%s tem que exigir publicação na consulta:\n%s", fn, corpo)
		}
		if !strings.Contains(corpo, "patient_id = ?") {
			t.Errorf("%s tem que escopar pelo paciente na consulta:\n%s", fn, corpo)
		}
		// O que sai é a cópia congelada, nunca o rascunho vivo.
		if !strings.Contains(corpo, "toPublishedPlanDTO") {
			t.Errorf("%s tem que devolver o conteúdo publicado, não o content em edição:\n%s", fn, corpo)
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
