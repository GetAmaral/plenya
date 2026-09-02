package services

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/pdfdoc"
)

// janelaDeCoalescencia — dentro dela, gravações seguidas do MESMO autor humano editando à mão
// atualizam a última revisão em vez de criar outra.
//
// Não é economia prematura. O editor salva nas fronteiras naturais (colapsar cartão, trocar de
// slide), e uma sessão de autoria de 90 minutos gera facilmente centenas de gravações. Um deck de
// 20 slides com réguas passa de 60 KB de JSON: sem coalescer, uma sessão vira ~100 MB, e o ano
// vira ~100 GB. Com esta janela cai para ~0,5 GB/ano, e nada do que importa se perde — o que se
// perde são estados intermediários de digitação do mesmo autor no mesmo minuto.
const janelaDeCoalescencia = 2 * time.Minute

// ErrPlanRevisionConflict — a gravação partiu de uma revisão que não é mais a corrente.
//
// Existe por causa do segundo escritor: enquanto só o médico editava, sobrescrever era aceitável.
// Com o assistente escrevendo no mesmo `content`, um salvamento em voo (carregado antes do turno)
// apagaria em silêncio o que a ferramenta acabou de aplicar.
var ErrPlanRevisionConflict = errors.New("o plano mudou desde que esta edição começou")

type PatientPlanRevisionService struct {
	db *gorm.DB
}

func NewPatientPlanRevisionService(db *gorm.DB) *PatientPlanRevisionService {
	return &PatientPlanRevisionService{db: db}
}

// RecordRevisionInput — o que uma gravação precisa declarar.
type RecordRevisionInput struct {
	Plan        *models.PatientPlan
	Title       string
	Content     []pdfdoc.DeckSlide
	AuthorKind  models.PatientPlanAuthorKind
	CreatedByID uuid.UUID
	Reason      models.PatientPlanRevisionReason

	Ops             any
	MessageID       *uuid.UUID
	DossierID       *uuid.UUID
	AIModel         string
	AIPromptVersion string
	IsPublication   bool
}

// Record grava uma revisão e devolve o novo `revision_seq` do plano.
//
// Roda DENTRO da transação do chamador, de propósito: a revisão e a atualização do plano têm que
// ser atômicas, senão sobra revisão sem plano ou plano sem histórico.
//
// Três comportamentos que valem por documentação:
//
//  1. EDIÇÃO cujo conteúdo é idêntico ao da última revisão não vira linha (devolve o seq atual);
//  2. edição humana à mão dentro da janela coalesce com a anterior do mesmo autor;
//  3. tudo o mais — o assistente aplicando, aceite de sugestão, restauração e publicação — sempre
//     cria linha, porque é evento auditável.
func (s *PatientPlanRevisionService) Record(tx *gorm.DB, in RecordRevisionInput) (int, error) {
	if in.Plan == nil {
		return 0, errors.New("revisão sem plano")
	}
	hash, err := HashDeckContent(in.Content)
	if err != nil {
		return 0, err
	}

	var ultima models.PatientPlanRevision
	temUltima := true
	if err := tx.Where("plan_id = ?", in.Plan.ID).
		Order("seq DESC").First(&ultima).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, err
		}
		temUltima = false
	}

	// 1. Nada mudou — mas só vale para edição à mão.
	//
	// Publicar, restaurar, aplicar o que o assistente escreveu e aceitar sugestão são EVENTOS: têm
	// que deixar linha mesmo quando o conteúdo resultante é idêntico ao anterior. Publicar sem ter
	// mexido em nada é o caso comum (escreve, confere, publica) e foi exatamente onde este atalho
	// engoliu a revisão de publicação na primeira versão deste código.
	if temUltima && eventoDeConteudo(in.Reason) &&
		ultima.ContentHash == hash && ultima.Title == in.Title {
		return in.Plan.RevisionSeq, nil
	}

	// 2. Coalescência: mesma pessoa, editando à mão, no mesmo intervalo curto.
	if temUltima && podeCoalescer(&ultima, in) {
		ultima.Title = in.Title
		ultima.Content = in.Content
		ultima.ContentHash = hash
		ultima.CreatedAt = time.Now()
		if err := tx.Model(&ultima).
			Select("title", "content", "content_hash", "created_at").
			Updates(&ultima).Error; err != nil {
			return 0, err
		}
		return ultima.Seq, nil
	}

	// 3. Linha nova.
	seq := 1
	if temUltima {
		seq = ultima.Seq + 1
	}
	rev := models.PatientPlanRevision{
		PlanID:          in.Plan.ID,
		Seq:             seq,
		PlanVersion:     in.Plan.Version,
		Title:           in.Title,
		Content:         in.Content,
		ContentHash:     hash,
		AuthorKind:      in.AuthorKind,
		CreatedByID:     in.CreatedByID,
		Reason:          in.Reason,
		Ops:             in.Ops,
		MessageID:       in.MessageID,
		DossierID:       in.DossierID,
		AIModel:         in.AIModel,
		AIPromptVersion: in.AIPromptVersion,
		IsPublication:   in.IsPublication,
	}
	if in.IsPublication {
		// Calculado ANTES de inserir: a própria revisão de publicação não escreve conteúdo, e
		// incluí-la só adicionaria ruído.
		if caminhos, err := s.AITouchedPaths(tx, in.Plan.ID); err == nil && len(caminhos) > 0 {
			rev.AITouchedPaths = caminhos
		}
	}
	if err := tx.Create(&rev).Error; err != nil {
		return 0, err
	}
	if in.IsPublication {
		if err := tx.Model(&models.PatientPlan{}).Where("id = ?", in.Plan.ID).
			Update("published_revision_id", rev.ID).Error; err != nil {
			return 0, err
		}
		// Escreve TAMBÉM no struct do chamador, e não só na linha.
		//
		// Sem esta linha o campo nunca sobreviveu: `Publish` chama `Record` e depois faz
		// `tx.Save(plan)` com o struct carregado ANTES deste UPDATE, e o Save do GORM grava todos
		// os campos — devolvendo a coluna a NULL na mesma transação. Os sete planos publicados em
		// dev estavam todos com `published_revision_id` nulo, ou seja, "o que o paciente leu em
		// janeiro" nunca virou consulta, que era o objetivo declarado da migration 00092.
		in.Plan.PublishedRevisionID = &rev.ID
	}
	return seq, nil
}

// eventoDeConteudo diz se a gravação é uma edição comum, cujo resultado é o que importa, em
// oposição a um evento cuja OCORRÊNCIA é o que importa registrar.
func eventoDeConteudo(r models.PatientPlanRevisionReason) bool {
	return r == models.PlanRevisionEdit
}

func podeCoalescer(ultima *models.PatientPlanRevision, in RecordRevisionInput) bool {
	return ultima.AuthorKind == models.PlanAuthorHuman &&
		in.AuthorKind == models.PlanAuthorHuman &&
		ultima.Reason == models.PlanRevisionEdit &&
		in.Reason == models.PlanRevisionEdit &&
		!ultima.IsPublication &&
		ultima.CreatedByID == in.CreatedByID &&
		time.Since(ultima.CreatedAt) < janelaDeCoalescencia
}

// CheckExpected compara o token de concorrência que o cliente mandou com a revisão corrente.
// `nil` passa: cliente antigo que ainda não manda o token não é bloqueado.
func CheckExpected(plan *models.PatientPlan, expected *int) error {
	if expected == nil || *expected == plan.RevisionSeq {
		return nil
	}
	return fmt.Errorf("%w: esperava a edição %d, a atual é %d", ErrPlanRevisionConflict, *expected, plan.RevisionSeq)
}

// HashDeckContent produz o sha256 do conteúdo. Serializa com `json.Marshal`, que ordena as chaves
// de struct pela ordem de declaração — determinístico para o mesmo binário, que é o que se precisa
// aqui (o hash só é comparado com outro hash gerado pela mesma versão do código).
func HashDeckContent(slides []pdfdoc.DeckSlide) (string, error) {
	if slides == nil {
		slides = []pdfdoc.DeckSlide{}
	}
	b, err := json.Marshal(slides)
	if err != nil {
		return "", err
	}
	soma := sha256.Sum256(b)
	return hex.EncodeToString(soma[:]), nil
}

// EnsureSlideIDs preenche o id dos slides que não têm.
//
// Deck escrito à mão pela skill e plano criado antes da migration chegam sem id. Preencher ao
// carregar (e não só na migration) é o que garante que toda operação tenha alvo estável, inclusive
// em conteúdo que entrou depois por um PUT vindo de fora da tela.
func EnsureSlideIDs(slides []pdfdoc.DeckSlide) ([]pdfdoc.DeckSlide, bool) {
	mudou := false
	for i := range slides {
		if slides[i].ID == "" {
			slides[i].ID = uuid.Must(uuid.NewV7()).String()
			mudou = true
		}
	}
	return slides, mudou
}

// AITouchedPaths devolve os caminhos cujo ÚLTIMO escritor foi o assistente, dentro da versão
// publicada corrente.
//
// A lógica é a de um "quem tocou por último": percorre as revisões da versão em ordem cronológica
// e, para cada caminho, guarda quem escreveu por último. Sobram os caminhos onde a ferramenta
// escreveu e o médico não voltou.
//
// Isso NÃO aparece para o paciente. É registro interno, e existe para responder duas perguntas que
// de outro jeito seriam arqueologia: "esta devolutiva tem frase gerada que ninguém revisou?" e,
// somando muitos planos, "a revisão está de fato acontecendo?". A evidência sobre revisão de texto
// redigido por IA é de que ela falha com frequência; a única forma de saber se está falhando aqui
// é medir.
// Percorre a cadeia INTEIRA do plano, não só a versão corrente. Escopar por `plan_version` foi
// erro meu, e errava para o lado perigoso: a partir da segunda publicação o cálculo só enxergaria o
// que mudou desde a anterior, e um deck escrito pela ferramenta na v1 e republicado sem alteração
// apareceria como se nada nele tivesse vindo dela. A pergunta é sobre o conteúdo que o paciente
// tem na mão, não sobre o intervalo entre duas publicações.
func (s *PatientPlanRevisionService) AITouchedPaths(tx *gorm.DB, planID uuid.UUID) ([]string, error) {
	var revs []models.PatientPlanRevision
	if err := tx.Where("plan_id = ?", planID).
		Order("seq").Find(&revs).Error; err != nil {
		return nil, err
	}

	// caminho -> quem escreveu por último
	ultimoAutor := map[string]models.PatientPlanAuthorKind{}
	for _, r := range revs {
		if r.Ops == nil {
			// Edição à mão pela tela reescreve o conteúdo inteiro e não declara caminhos, então
			// não dá para saber o que ela tocou.
			//
			// Ela NÃO limpa a atribuição. Salvar não é ler: o médico pode ter mexido só no slide 3
			// e apertado salvar com os outros dezenove como a ferramenta os deixou. Assumir que
			// tudo passou pela revisão dele é exatamente a suposição que a evidência de viés de
			// automação desmente, e é a que faz a coluna mentir para o lado perigoso.
			//
			// A consequência é que este número super-reporta: um trecho reescrito à mão continua
			// contando como da ferramenta até uma op declarada dizer o contrário. Errar para cima
			// é o lado certo de errar aqui.
			continue
		}
		for _, caminho := range caminhosDasOps(r.Ops) {
			ultimoAutor[caminho] = r.AuthorKind
		}
	}

	var out []string
	for caminho, autor := range ultimoAutor {
		if autor == models.PlanAuthorAssistant {
			out = append(out, caminho)
		}
	}
	sort.Strings(out)
	return out, nil
}

// caminhosDasOps extrai "slideId:path" de um `ops` gravado como JSON solto.
func caminhosDasOps(ops any) []string {
	bruto, err := json.Marshal(ops)
	if err != nil {
		return nil
	}
	var lista []struct {
		Op struct {
			SlideID string `json:"slideId"`
			Path    string `json:"path"`
			Op      string `json:"op"`
		} `json:"op"`
		Decision string `json:"decision"`
	}
	if err := json.Unmarshal(bruto, &lista); err != nil {
		return nil
	}
	var out []string
	for _, l := range lista {
		// Só o que de fato entrou no conteúdo conta; sugestão e recusa não escreveram nada.
		if l.Decision != "" && l.Decision != string(TriageApply) {
			continue
		}
		if l.Op.Path == "" {
			out = append(out, l.Op.SlideID+":"+l.Op.Op)
			continue
		}
		out = append(out, l.Op.SlideID+":"+l.Op.Path)
	}
	return out
}

// PlanRevisionSummary — uma linha do histórico. Sem o `content`, que é o volume todo: a lista
// carregaria dezenas de decks inteiros para desenhar dezenas de linhas.
type PlanRevisionSummary struct {
	ID             uuid.UUID                        `json:"id"`
	Seq            int                              `json:"seq"`
	PlanVersion    int                              `json:"planVersion"`
	Title          string                           `json:"title"`
	AuthorKind     models.PatientPlanAuthorKind     `json:"authorKind"`
	AuthorName     string                           `json:"authorName"`
	Reason         models.PatientPlanRevisionReason `json:"reason"`
	IsPublication  bool                             `json:"isPublication"`
	Slides         int                              `json:"slides"`
	ChangedPaths   []string                         `json:"changedPaths,omitempty"`
	AITouchedPaths []string                         `json:"aiTouchedPaths,omitempty"`
	AIModel        string                           `json:"aiModel,omitempty"`
	CreatedAt      time.Time                        `json:"createdAt"`
}

// ListRevisions devolve o histórico do plano, do mais recente para o mais antigo.
func (s *PatientPlanRevisionService) ListRevisions(planID uuid.UUID, limite int) ([]PlanRevisionSummary, error) {
	if limite <= 0 || limite > 200 {
		limite = 100
	}
	var revs []models.PatientPlanRevision
	if err := s.db.Preload("CreatedBy").Where("plan_id = ?", planID).
		Order("seq DESC").Limit(limite).Find(&revs).Error; err != nil {
		return nil, err
	}
	out := make([]PlanRevisionSummary, 0, len(revs))
	for i := range revs {
		r := &revs[i]
		linha := PlanRevisionSummary{
			ID: r.ID, Seq: r.Seq, PlanVersion: r.PlanVersion, Title: r.Title,
			AuthorKind: r.AuthorKind, AuthorName: r.CreatedBy.Name, Reason: r.Reason,
			IsPublication: r.IsPublication, Slides: len(r.Content),
			AIModel: r.AIModel, CreatedAt: r.CreatedAt,
		}
		if r.Ops != nil {
			linha.ChangedPaths = caminhosDasOps(r.Ops)
		}
		if r.AITouchedPaths != nil {
			bruto, err := json.Marshal(r.AITouchedPaths)
			if err == nil {
				_ = json.Unmarshal(bruto, &linha.AITouchedPaths)
			}
		}
		out = append(out, linha)
	}
	return out, nil
}

// Restore devolve o rascunho ao estado de uma revisão.
//
// Não apaga nada: restaurar GRAVA uma revisão nova com o conteúdo antigo. Desfazer que destrói
// histórico é o mesmo defeito que esta tabela existe para consertar, e "restaurei a v3 por engano,
// volta" tem que continuar sendo possível.
//
// O autor é `human` mesmo quando a revisão restaurada foi escrita pela ferramenta: a decisão de
// trazer aquele texto de volta é do médico, e é a decisão que a trilha precisa registrar. A
// atribuição de quem escreveu o texto continua na revisão de origem, e `ai_touched_paths` continua
// contando, porque restaurar não declara caminhos.
func (s *PatientPlanRevisionService) Restore(plan *models.PatientPlan, revisionID, userID uuid.UUID) (int, error) {
	var alvo models.PatientPlanRevision
	if err := s.db.Where("id = ? AND plan_id = ?", revisionID, plan.ID).First(&alvo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, fmt.Errorf("revisão não encontrada neste plano")
		}
		return 0, err
	}
	var seq int
	err := s.db.Transaction(func(tx *gorm.DB) error {
		plan.Title = alvo.Title
		plan.Content = alvo.Content
		// Restaurar devolve o plano a rascunho: o que está no portal continua sendo a versão
		// publicada até alguém publicar de novo.
		plan.Status = models.PatientPlanDraft
		novo, err := s.Record(tx, RecordRevisionInput{
			Plan: plan, Title: plan.Title, Content: plan.Content,
			AuthorKind: models.PlanAuthorHuman, CreatedByID: userID,
			Reason: models.PlanRevisionRestore,
		})
		if err != nil {
			return err
		}
		seq = novo
		plan.RevisionSeq = novo
		return tx.Save(plan).Error
	})
	return seq, err
}
