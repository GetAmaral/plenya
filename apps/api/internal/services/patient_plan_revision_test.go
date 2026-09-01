package services

import (
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/pdfdoc"
)

func slidesDeTeste(titulo string) []pdfdoc.DeckSlide {
	return []pdfdoc.DeckSlide{
		{ID: "s1", Kind: pdfdoc.DeckCover, Title: titulo},
		{ID: "s2", Kind: pdfdoc.DeckClosing, Title: "fecho"},
	}
}

// O hash é o que decide se uma gravação vira linha. Precisa ser estável para o mesmo conteúdo e
// mudar para qualquer diferença, inclusive de ordem: reordenar slide é mudança de verdade.
func TestHashDeckContent(t *testing.T) {
	a, err := HashDeckContent(slidesDeTeste("um"))
	if err != nil {
		t.Fatal(err)
	}
	b, _ := HashDeckContent(slidesDeTeste("um"))
	if a != b {
		t.Error("mesmo conteúdo deu hash diferente; gravação sem mudança viraria revisão")
	}

	c, _ := HashDeckContent(slidesDeTeste("dois"))
	if a == c {
		t.Error("conteúdo diferente deu o mesmo hash")
	}

	invertido := slidesDeTeste("um")
	invertido[0], invertido[1] = invertido[1], invertido[0]
	d, _ := HashDeckContent(invertido)
	if a == d {
		t.Error("reordenar slides não mudou o hash; reordenação é mudança")
	}

	vazio, err := HashDeckContent(nil)
	if err != nil || vazio == "" {
		t.Errorf("conteúdo nil deveria hashear como lista vazia, veio %q / %v", vazio, err)
	}
	if len(a) != 64 {
		t.Errorf("hash com %d caracteres, esperava 64 (sha256 em hex)", len(a))
	}
}

// Slide sem id não pode existir depois de carregado: é o alvo de toda operação e sugestão.
func TestEnsureSlideIDs(t *testing.T) {
	entrada := []pdfdoc.DeckSlide{
		{Kind: pdfdoc.DeckCover},
		{ID: "ja-tinha", Kind: pdfdoc.DeckTableKind},
		{Kind: pdfdoc.DeckClosing},
	}
	saida, mudou := EnsureSlideIDs(entrada)

	if !mudou {
		t.Error("havia slide sem id e a função disse que nada mudou")
	}
	if saida[1].ID != "ja-tinha" {
		t.Errorf("id existente foi sobrescrito: %q", saida[1].ID)
	}
	for i, s := range saida {
		if s.ID == "" {
			t.Errorf("slide %d continua sem id", i)
		}
	}
	if saida[0].ID == saida[2].ID {
		t.Error("dois slides receberam o mesmo id")
	}

	// Rodar de novo não pode trocar id: quem já foi endereçado continua endereçável.
	antes := saida[0].ID
	denovo, mudou2 := EnsureSlideIDs(saida)
	if mudou2 {
		t.Error("segunda passada disse que mudou algo")
	}
	if denovo[0].ID != antes {
		t.Errorf("id trocou entre passadas: %q -> %q", antes, denovo[0].ID)
	}
}

// O token de concorrência é o que impede um salvamento em voo de apagar o que o assistente
// acabou de escrever. Ausente passa, para não quebrar cliente que ainda não manda.
func TestCheckExpected(t *testing.T) {
	plan := &models.PatientPlan{RevisionSeq: 7}

	if err := CheckExpected(plan, nil); err != nil {
		t.Errorf("sem token deveria passar, veio %v", err)
	}
	sete := 7
	if err := CheckExpected(plan, &sete); err != nil {
		t.Errorf("token igual deveria passar, veio %v", err)
	}
	cinco := 5
	err := CheckExpected(plan, &cinco)
	if err == nil {
		t.Fatal("token velho passou; um autosave atrasado apagaria escrita que o cliente não viu")
	}
	if !errorIs(err, ErrPlanRevisionConflict) {
		t.Errorf("erro %v não é ErrPlanRevisionConflict, o handler não devolveria 409", err)
	}
}

func errorIs(err, alvo error) bool {
	for err != nil {
		if err == alvo {
			return true
		}
		u, ok := err.(interface{ Unwrap() error })
		if !ok {
			return false
		}
		err = u.Unwrap()
	}
	return false
}

// A coalescência é o que separa 0,5 GB/ano de 100 GB/ano, mas ela não pode engolir evento
// auditável: o que a ferramenta escreveu, o aceite de sugestão, a restauração e a publicação
// sempre viram linha própria.
func TestPodeCoalescer(t *testing.T) {
	autor := uuid.New()
	outro := uuid.New()
	agora := func(d time.Duration) time.Time { return time.Now().Add(-d) }

	base := func() *models.PatientPlanRevision {
		return &models.PatientPlanRevision{
			AuthorKind:  models.PlanAuthorHuman,
			Reason:      models.PlanRevisionEdit,
			CreatedByID: autor,
			CreatedAt:   agora(30 * time.Second),
		}
	}
	entradaHumana := RecordRevisionInput{
		AuthorKind:  models.PlanAuthorHuman,
		CreatedByID: autor,
		Reason:      models.PlanRevisionEdit,
	}

	if !podeCoalescer(base(), entradaHumana) {
		t.Error("mesma pessoa editando à mão em 30s deveria coalescer")
	}

	velha := base()
	velha.CreatedAt = agora(5 * time.Minute)
	if podeCoalescer(velha, entradaHumana) {
		t.Error("fora da janela não pode coalescer")
	}

	outroAutor := base()
	outroAutor.CreatedByID = outro
	if podeCoalescer(outroAutor, entradaHumana) {
		t.Error("autor diferente não pode coalescer: perderia de quem foi a edição anterior")
	}

	publicada := base()
	publicada.IsPublication = true
	if podeCoalescer(publicada, entradaHumana) {
		t.Error("revisão de publicação é imutável; editar depois dela tem que virar linha nova")
	}

	for _, r := range []models.PatientPlanRevisionReason{
		models.PlanRevisionAIApply,
		models.PlanRevisionSuggestionAccept,
		models.PlanRevisionRestore,
		models.PlanRevisionPublish,
	} {
		anterior := base()
		anterior.Reason = r
		if podeCoalescer(anterior, entradaHumana) {
			t.Errorf("revisão anterior de motivo %q não pode ser sobrescrita por edição humana", r)
		}
		entradaIA := entradaHumana
		entradaIA.Reason = r
		if podeCoalescer(base(), entradaIA) {
			t.Errorf("gravação de motivo %q é evento auditável e não pode coalescer", r)
		}
	}

	entradaIA := RecordRevisionInput{
		AuthorKind:  models.PlanAuthorAssistant,
		CreatedByID: autor,
		Reason:      models.PlanRevisionEdit,
	}
	if podeCoalescer(base(), entradaIA) {
		t.Error("escrita do assistente não pode ser absorvida por uma revisão humana: some a autoria")
	}
}

// Publicar sem ter mudado nada é o caso comum: escreve, confere, publica. O atalho de "conteúdo
// idêntico" não pode engolir esse registro — foi o que aconteceu na primeira versão deste código,
// e o efeito era `published_revision_id` nulo, ou seja, nenhuma forma de recuperar o que o
// paciente leu naquela versão.
func TestEventoDeConteudo(t *testing.T) {
	if !eventoDeConteudo(models.PlanRevisionEdit) {
		t.Error("edição é gravação de conteúdo: idêntica não deve virar linha")
	}
	for _, r := range []models.PatientPlanRevisionReason{
		models.PlanRevisionPublish,
		models.PlanRevisionRestore,
		models.PlanRevisionAIApply,
		models.PlanRevisionSuggestionAccept,
	} {
		if eventoDeConteudo(r) {
			t.Errorf("%q é evento: tem que virar linha mesmo com conteúdo idêntico", r)
		}
	}
}
