package services

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
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
	if err := tx.Create(&rev).Error; err != nil {
		return 0, err
	}
	if in.IsPublication {
		if err := tx.Model(&models.PatientPlan{}).Where("id = ?", in.Plan.ID).
			Update("published_revision_id", rev.ID).Error; err != nil {
			return 0, err
		}
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
