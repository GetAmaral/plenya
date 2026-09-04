package services

// AssembleDraft — criar o rascunho SEM chamar modelo nenhum.
//
// Mesmo caminho de `GenerateDraft` (cria o plano, congela o dossiê, confere, grava revisão), com o
// miolo trocado: onde a geração faz sete chamadas ao modelo e custa ~US$ 0,21, aqui `MontaDeckLocal`
// monta o deck em memória, de graça e em milissegundos.
//
// O que muda na saída, e é intencional:
//
//   - autoria `system`, não `assistant`. Nenhuma palavra veio de modelo, então nada deste deck pode
//     aparecer em `ai_touched_paths` na publicação. Marcar como assistente seria mentira auditável.
//   - `punch` vazio na maioria dos slides, trocado por UM aviso de lacuna em vez de um por slide: é
//     a leitura clínica que a conversa com o médico acrescenta depois.
//   - a conferência NUMÉRICA não roda. Todo número veio do dossiê por construção; o que sobra de
//     texto livre são as palavras do próprio médico no plano de cuidado, e acusar a dose que ELE
//     escreveu como "número sem origem" seria o aviso trocar de dono.

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/pdfdoc"
	"gorm.io/gorm"
)

// AssembleDraft monta e grava o rascunho mecânico.
func (s *PatientPlanAssistantService) AssembleDraft(in GenerateDraftInput) (*GenerateDraftResult, error) {
	titulo := strings.TrimSpace(in.Title)
	if titulo == "" {
		titulo = "Seus exames"
	}
	criado, err := s.plans.Create(in.PatientID, in.UserID, &dto.SavePatientPlanRequest{
		Title: titulo, Content: []pdfdoc.DeckSlide{},
	})
	if err != nil {
		return nil, err
	}
	planID, err := uuid.Parse(criado.ID)
	if err != nil {
		return nil, err
	}
	limpaSeFalhar := func(e error) (*GenerateDraftResult, error) {
		_ = s.plans.Delete(planID, in.PatientID)
		return nil, e
	}

	dossie, dossieRow, err := s.dossiers.Current(planID)
	if err != nil {
		return limpaSeFalhar(err)
	}
	if dossieVazio(dossie) {
		return limpaSeFalhar(errors.New("este paciente não tem exame nem anamnese suficientes para uma devolutiva"))
	}
	secoes := montaArco(dossie)
	if semMaterial(secoes, dossie) {
		return limpaSeFalhar(errors.New("este paciente não tem exame nem anamnese suficientes para uma devolutiva"))
	}

	slides := MontaDeckLocal(dossie)
	if len(slides) == 0 {
		return limpaSeFalhar(errors.New("não sobrou material para montar o deck"))
	}

	plan, err := s.plans.load(planID, in.PatientID)
	if err != nil {
		return limpaSeFalhar(err)
	}
	if err := s.db.Transaction(func(tx *gorm.DB) error {
		plan.Content = slides
		plan.Status = models.PatientPlanDraft
		seq, errR := s.revisions.Record(tx, RecordRevisionInput{
			Plan: plan, Title: plan.Title, Content: plan.Content,
			// `system`: quem escreveu foi o servidor, a partir do prontuário.
			AuthorKind: models.PlanAuthorSystem, CreatedByID: in.UserID,
			Reason: models.PlanRevisionEdit, DossierID: &dossieRow.ID,
		})
		if errR != nil {
			return errR
		}
		plan.RevisionSeq = seq
		return tx.Save(plan).Error
	}); err != nil {
		return limpaSeFalhar(err)
	}

	// Sem reparo: reparar é reescrever prosa, e prosa é o que este caminho não escreve. O que
	// estourar vai para a tela e é o médico quem corta, na mesma conversa em que escreve o punch.
	estouro, _ := pdfdoc.CheckDeckOverflow(pdfdoc.Deck{Title: plan.Title, Slides: plan.Content})
	avisos := avisosDoMontado(plan.Content)
	avisos = append(avisos, lacunasDoProntuario(dossie)...)

	final, err := s.plans.Get(planID, in.PatientID)
	if err != nil {
		return nil, err
	}
	return &GenerateDraftResult{
		Plan: final, Warnings: avisos, Overflow: estouro,
		Reply: "Rascunho montado do prontuário, sem modelo. O texto de leitura clínica (os punches e os títulos como afirmação) é o que falta, e é o que se escreve na conversa.",
	}, nil
}

// avisosDoMontado troca os N avisos de "slide sem punch" por um só.
//
// Um por slide seria ruído: no deck montado a ausência é a regra, não o desvio. Um aviso com a
// contagem diz a mesma coisa e ainda diz o tamanho do trabalho que sobrou.
func avisosDoMontado(slides []pdfdoc.DeckSlide) []dto.PlanGenWarning {
	var out []dto.PlanGenWarning
	semPunch := 0
	for _, w := range confereDeck(slides) {
		if w.Reason == motivoSemPunch {
			semPunch++
			continue
		}
		out = append(out, w)
	}
	// Os cartões da conduta são o `rationale` do plano de cuidado, copiado palavra por palavra.
	// Isso é fidelidade, e é também o risco: aquele texto foi escrito para o PRONTUÁRIO, e traz
	// diretriz, sigla e número de recomendação que o paciente não lê ("a ESE/ENSAT 2023 pede
	// confirmação da independência do ACTH"). Quem publica precisa saber disso antes, não depois.
	condutas := 0
	for _, s := range slides {
		if s.Kind == pdfdoc.DeckPlanStep && len(s.Cards) > 0 {
			condutas++
		}
	}
	if condutas > 0 {
		out = append(out, dto.PlanGenWarning{
			// `estilo`, não `lacuna`: a tela põe lacuna na caixa do "o prontuário não tinha, e por
			// isso não entrou", que diz explicitamente NÃO ser aviso de defeito. Este é o oposto —
			// é o pedido de conferir um texto que já está no deck e vai para o paciente.
			Kind: dto.PlanGenWarningEstilo,
			Reason: fmt.Sprintf("%d slides do plano trazem o texto do plano de cuidado como está no prontuário: "+
				"confira se ele fala com o paciente antes de publicar", condutas),
		})
	}
	if semPunch > 0 {
		out = append(out, dto.PlanGenWarning{
			Kind: dto.PlanGenWarningEstilo,
			Reason: fmt.Sprintf("%d slides estão sem punch e com título descritivo: o código monta o que é dado, "+
				"a leitura clínica (por que o achado está assim, o que ele cobra) entra na conversa", semPunch),
		})
	}
	return out
}
