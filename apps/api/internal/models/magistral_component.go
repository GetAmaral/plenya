package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// MagistralComponent é uma substância do repertório magistral: nome, faixa de dose sugerida,
// densidade aparente e sinalizadores que alimentam compatibilidade e palatabilidade.
//
// Todo campo clínico é PONTEIRO de propósito: NULL significa "não cadastrado", e a tela precisa
// distinguir isso de zero. Dose sugerida ausente é silêncio; dose sugerida zero seria erro.
// @Description Componente do catálogo magistral
type MagistralComponent struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Nome da substância (DCB quando existir)
	// @example Melatonina
	Name string `gorm:"type:varchar(200);not null" json:"name" validate:"required,min=2,max=200"`

	// Sinônimos separados por vírgula (busca)
	Synonyms string `gorm:"type:text;not null;default:''" json:"synonyms"`

	CAS     *string `gorm:"type:varchar(20)" json:"cas,omitempty"`
	DCBCode *string `gorm:"type:varchar(20)" json:"dcbCode,omitempty"`

	// Unidade em que a substância costuma ser prescrita
	DefaultUnit string `gorm:"type:varchar(20);not null;default:'mg'" json:"defaultUnit"`

	// Faixa de dose por unidade aviada. NULL = sem sugestão.
	UsualDose *float64 `gorm:"type:numeric(14,4)" json:"usualDose,omitempty"`
	MinDose   *float64 `gorm:"type:numeric(14,4)" json:"minDose,omitempty"`
	MaxDose   *float64 `gorm:"type:numeric(14,4)" json:"maxDose,omitempty"`
	// DoseBasis diz se a faixa acima é do DIA ou de UMA tomada. Sem isto o painel comparava a
	// dose de uma cápsula contra faixa diária numas substâncias e contra faixa por tomada em
	// outras — a mesma coluna significando duas coisas.
	DoseBasis string `gorm:"type:varchar(10);not null;default:'por_dia'" json:"doseBasis"`

	// DefaultCategory — a categoria de receita que a substância carrega. Sem isto, digitar
	// "testosterona" numa fórmula magistral emitia receita simples por default, quando a
	// Portaria 344/98 pede Controle Especial em duas vias.
	DefaultCategory MedicationCategory `gorm:"type:varchar(20);not null;default:'simple'" json:"defaultCategory"`
	// RegulatoryNote — restrição de finalidade ou exigência que acompanha a substância.
	RegulatoryNote *string `gorm:"type:text" json:"regulatoryNote,omitempty"`

	// AssayInterference — como a substância atrapalha exame laboratorial, e o que fazer antes da
	// coleta. Importa duas vezes aqui: a receita sai deste sistema e o exame que ela corrompe
	// volta para ele, alimentando as regras de dose dinâmica.
	AssayInterference *string `gorm:"type:text" json:"assayInterference,omitempty"`
	// Dose diária a partir da qual a interferência é descrita. Nulo com texto preenchido
	// significa que interfere em qualquer dose.
	AssayInterferenceDose *float64 `gorm:"type:numeric(14,4)" json:"assayInterferenceDose,omitempty"`

	// Nutriente correspondente no Anexo IV da IN 28. Várias substâncias apontam para o mesmo de
	// propósito: é assim que P5P e vitamina B6 na mesma fórmula SOMAM antes de comparar com o
	// teto, que foi exatamente o erro achado no formulário das parceiras.
	In28Nutrient *string `gorm:"type:varchar(400)" json:"in28Nutrient,omitempty"`
	// Quantas unidades do Anexo IV valem uma unidade desta substância.
	In28Factor float64 `gorm:"type:numeric(14,6);not null;default:1" json:"in28Factor"`

	// Densidade aparente (g/mL). NULL = a calculadora de cápsula não opina.
	BulkDensity *float64 `gorm:"type:numeric(8,4)" json:"bulkDensity,omitempty"`
	// De onde veio a densidade: 'medida' (publicada para a substância), 'classe' (aproximação
	// pela classe do pó) ou o nome da ficha técnica. A tela precisa distinguir estimativa de
	// medida — o cálculo é o mesmo, a confiança não.
	DensitySource *string `gorm:"type:varchar(60)" json:"densitySource,omitempty"`

	EutecticFormer     bool `gorm:"not null;default:false" json:"eutecticFormer"`
	Hygroscopic        bool `gorm:"not null;default:false" json:"hygroscopic"`
	Oxidizing          bool `gorm:"not null;default:false" json:"oxidizing"`
	OxidationSensitive bool `gorm:"not null;default:false" json:"oxidationSensitive"`
	Photosensitive     bool `gorm:"not null;default:false" json:"photosensitive"`

	// 0 sem amargor · 1 leve · 2 marcante · 3 intolerável em sachê. NULL = não avaliado.
	Bitterness *int  `gorm:"type:smallint" json:"bitterness,omitempty"`
	SachetOK   *bool `json:"sachetOk,omitempty"`

	Notes  *string `gorm:"type:text" json:"notes,omitempty"`
	Source string  `gorm:"type:varchar(30);not null;default:'manual'" json:"source"`

	// Marca do insumo, quando o que se prescreve é a marca (CavaQ10, Morosil, Exsynutriment).
	Brand *string `gorm:"type:varchar(60)" json:"brand,omitempty"`

	// Percentual de elemento (ou de ativo) no insumo: bisglicinato de magnésio 30%,
	// selenometionina 1%. É o que permite converter dose do ELEMENTO em massa do INSUMO — e é a
	// massa do insumo que ocupa a cápsula.
	ElementalPercent *float64 `gorm:"type:numeric(7,3)" json:"elementalPercent,omitempty"`
	CorrectionNote   *string  `gorm:"type:varchar(300)" json:"correctionNote,omitempty"`

	// Forma que o prescritor usa no lugar desta. A tela sugere a troca em vez de depender da
	// memória de quem prescreve.
	PreferredAlternativeID *uuid.UUID          `gorm:"type:uuid" json:"preferredAlternativeId,omitempty"`
	PreferenceNote         *string             `gorm:"type:varchar(300)" json:"preferenceNote,omitempty"`
	PreferredAlternative   *MagistralComponent `gorm:"foreignKey:PreferredAlternativeID" json:"preferredAlternative,omitempty"`

	// Para que serve, em prosa clínica curta. É o que faz a busca responder "o que eu uso para
	// sono?" em vez de exigir que se saiba o nome da substância de antemão.
	Indications *string `gorm:"type:text" json:"indications,omitempty"`

	// Posologia em texto, com a origem embutida ("aula X da pós", "formulário magistral").
	// Complementa a faixa numérica, que não cabe posologia ("1x ao dia, com a refeição").
	DoseReference *string `gorm:"type:text" json:"doseReference,omitempty"`

	// Os mesmos dois campos em TÓPICOS, uma linha por item. É o que a tela mostra por default;
	// o texto corrido fica a um clique, para conferência. Texto puro, não jsonb: o campo é
	// editado à mão num textarea e não há por que repetir a cicatriz da migration 00060.
	IndicationBullets *string `gorm:"type:text" json:"indicationBullets,omitempty"`
	DoseBullets       *string `gorm:"type:text" json:"doseBullets,omitempty"`

	// pending (nunca enriquecido) · suggested (veio do RAG/pesquisa, aguarda conferência) ·
	// confirmed (o médico conferiu). Curadoria manual carimba confirmed.
	EvidenceStatus string     `gorm:"type:varchar(12);not null;default:'pending'" json:"evidenceStatus"`
	EnrichedAt     *time.Time `gorm:"type:timestamptz" json:"enrichedAt,omitempty"`

	ReviewedBy *uuid.UUID `gorm:"type:uuid" json:"reviewedBy,omitempty"`
	LastReview *time.Time `gorm:"type:timestamptz" json:"lastReview,omitempty"`

	// Quantas vezes já foi prescrita — ordena a busca pelo repertório real do prescritor.
	UsageCount int  `gorm:"type:integer;not null;default:0" json:"usageCount"`
	IsActive   bool `gorm:"not null;default:true" json:"isActive"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (MagistralComponent) TableName() string { return "magistral_components" }

// MagistralComponentArticle liga um componente a um trecho de artigo/aula do RAG.
// Evidência é ANEXADA, nunca gerada: o trecho fica guardado para leitura humana e nenhum cálculo
// do sistema consome article_id.
// @Description Evidência (artigo/aula do RAG) ligada a um componente magistral
type MagistralComponentArticle struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	ComponentID uuid.UUID  `gorm:"type:uuid;not null;index" json:"componentId"`
	ArticleID   uuid.UUID  `gorm:"type:uuid;not null" json:"articleId"`
	EmbeddingID *uuid.UUID `gorm:"type:uuid" json:"embeddingId,omitempty"`
	ChunkIndex  *int       `gorm:"type:integer" json:"chunkIndex,omitempty"`
	Similarity  *float64   `gorm:"type:numeric(6,4)" json:"similarity,omitempty"`
	Excerpt     string     `gorm:"type:text;not null;default:''" json:"excerpt"`
	// Fixado pelo médico. A sugestão automática nasce não-fixada.
	Pinned   bool       `gorm:"not null;default:false" json:"pinned"`
	PinnedBy *uuid.UUID `gorm:"type:uuid" json:"pinnedBy,omitempty"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Article *Article `gorm:"foreignKey:ArticleID" json:"article,omitempty"`
}

func (MagistralComponentArticle) TableName() string { return "magistral_component_articles" }

func (m *MagistralComponentArticle) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

func (m *MagistralComponent) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// BeforeUpdate carimba a revisão — mesmo padrão de ScoreItem/ScoreLevel: dado clínico editado
// sem data de revisão é dado que ninguém sabe se ainda vale.
func (m *MagistralComponent) BeforeUpdate(tx *gorm.DB) error {
	now := time.Now()
	m.LastReview = &now
	return nil
}

// IncompatibilitySeverity — info observa, warn desaconselha, avoid é "não associe".
type IncompatibilitySeverity string

const (
	IncompatInfo  IncompatibilitySeverity = "info"
	IncompatWarn  IncompatibilitySeverity = "warn"
	IncompatAvoid IncompatibilitySeverity = "avoid"
)

// MagistralIncompatibility é um par de substâncias que não convivem bem, com o mecanismo.
// @Description Incompatibilidade entre dois componentes magistrais
type MagistralIncompatibility struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	ComponentAID uuid.UUID `gorm:"type:uuid;not null" json:"componentAId" validate:"required"`
	ComponentBID uuid.UUID `gorm:"type:uuid;not null" json:"componentBId" validate:"required"`

	Severity  IncompatibilitySeverity `gorm:"type:varchar(10);not null;default:'warn'" json:"severity"`
	Mechanism string                  `gorm:"type:varchar(200);not null;default:''" json:"mechanism"`
	Note      *string                 `gorm:"type:text" json:"note,omitempty"`
	Source    string                  `gorm:"type:varchar(200);not null;default:''" json:"source"`

	ReviewedBy *uuid.UUID `gorm:"type:uuid" json:"reviewedBy,omitempty"`
	LastReview *time.Time `gorm:"type:timestamptz" json:"lastReview,omitempty"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	ComponentA *MagistralComponent `gorm:"foreignKey:ComponentAID" json:"componentA,omitempty"`
	ComponentB *MagistralComponent `gorm:"foreignKey:ComponentBID" json:"componentB,omitempty"`
}

func (MagistralIncompatibility) TableName() string { return "magistral_incompatibilities" }

func (m *MagistralIncompatibility) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.Must(uuid.NewV7())
	}
	// O par é simétrico: normalizar a ordem é o que faz o índice único impedir duplicata.
	if m.ComponentBID.String() < m.ComponentAID.String() {
		m.ComponentAID, m.ComponentBID = m.ComponentBID, m.ComponentAID
	}
	return nil
}
