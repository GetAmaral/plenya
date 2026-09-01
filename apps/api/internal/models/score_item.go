package models

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ScoreItem represents a specific clinical parameter (e.g., "Hemoglobina - Homens", "FEVE")
// @Description Item de escore - parâmetro clínico específico
type ScoreItem struct {
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Código de referência para LabTestDefinition (pode ser compartilhado entre múltiplos score_items)
	// Liga ao lab_test_definitions.code para associar resultados de exames
	// @example PLN585CE3E3, PLNF66C0E48, GLUCOSE_FASTING
	LabTestCode *string `gorm:"type:varchar(100);index;column:lab_test_code" json:"labTestCode,omitempty" validate:"omitempty,max=100"`

	// Código para hardcoding de AnamneseItem - gerado automaticamente a partir do nome
	// Apenas preenchido quando lab_test_code é NULL (itens de anamnese, não de exame laboratorial)
	// Facilita referência estável por código em vez de UUID
	// @example ALTURA_CM, PESO_KG, TABAGISMO_ATUAL, COXA_CM_HOMEM
	AnamneseItemCode *string `gorm:"type:varchar(200);index;column:anamnese_item_code" json:"anamneseItemCode,omitempty" validate:"omitempty,max=200"`

	// @minLength 2
	// @maxLength 300
	// @example Hemoglobina - Homens
	Name string `gorm:"type:varchar(300);not null" json:"name" validate:"required,min=2,max=300"`

	// Nome completo computado (Group - Subgroup - ParentItem - Name)
	// @example Hemograma - Série Vermelha - Hemoglobina - Homens
	FullName string `gorm:"-" json:"fullName,omitempty"`

	// @example g/dL
	Unit *string `gorm:"type:varchar(50)" json:"unit,omitempty" validate:"omitempty,max=50"`

	// @example 1 g/dL = 10 g/L
	UnitConversion *string `gorm:"type:text" json:"unitConversion,omitempty"`

	// Gênero aplicável (not_applicable, male, female)
	// @enum not_applicable,male,female
	// @example male
	Gender *string `gorm:"type:varchar(20);default:'not_applicable';check:gender IN ('not_applicable','male','female')" json:"gender,omitempty" validate:"omitempty,oneof=not_applicable male female"`

	// Idade mínima aplicável (anos)
	// @minimum 0
	// @maximum 150
	// @example 18
	AgeRangeMin *int `gorm:"type:integer;check:age_range_min >= 0 AND age_range_min <= 150" json:"ageRangeMin,omitempty" validate:"omitempty,gte=0,lte=150"`

	// Idade máxima aplicável (anos)
	// @minimum 0
	// @maximum 150
	// @example 65
	AgeRangeMax *int `gorm:"type:integer;check:age_range_max >= 0 AND age_range_max <= 150" json:"ageRangeMax,omitempty" validate:"omitempty,gte=0,lte=150"`

	// RequiresLabCode / RequiresMin / RequiresMax — o item só se aplica quando o resultado mais
	// recente de OUTRO exame cai na faixa dada.
	//
	// Existe porque há item cuja escala só é interpretável dentro de um contexto: a razão
	// %Free PSA marca ≤10% como o pior nível, mas isso só discrimina quando o PSA TOTAL está entre
	// 4 e 10 ng/mL. Com PSA total baixo, uma razão de 8,8% não é achado nenhum — e virava o item
	// número 1 da devolutiva, com 28 pontos, num paciente sem nada na próstata.
	RequiresLabCode *string  `gorm:"type:varchar(100);index;column:requires_lab_code" json:"requiresLabCode,omitempty"`
	RequiresMin     *float64 `gorm:"column:requires_min" json:"requiresMin,omitempty"`
	RequiresMax     *float64 `gorm:"column:requires_max" json:"requiresMax,omitempty"`

	// Indica se o score_item é aplicável apenas para mulheres pós-menopausa
	// @example true
	PostMenopause *bool `gorm:"type:boolean" json:"postMenopause,omitempty"`

	// Restringe o item por uso de terapia de reposição hormonal: true = só para quem repõe,
	// false = só para quem não repõe, NULL = não filtra por isso. Existe porque o mesmo
	// analito tem faixas diferentes com e sem TRH (estradiol pós-menopausa é o caso).
	// @example false
	HormoneTherapy *bool `gorm:"type:boolean" json:"hormoneTherapy,omitempty"`

	// Se true, um template de anamnese que inclui este item deve pré-selecionar o nível 5
	// por padrão (útil para histórico de doença / uso de medicação: padrão = "sem doença" / "sem uso").
	// @example false
	DefaultLevel5 bool `gorm:"type:boolean;not null;default:false" json:"defaultLevel5"`

	// Relevância clínica - explicação técnica para profissionais de saúde
	// @example Valores baixos de hemoglobina indicam anemia, que pode estar associada a fadiga, redução da capacidade funcional e aumento do risco cardiovascular
	ClinicalRelevance *string `gorm:"type:text" json:"clinicalRelevance,omitempty"`

	// Explicação para o paciente - linguagem simples e acessível
	// @example Hemoglobina é a proteína que transporta oxigênio no sangue. Valores baixos podem causar cansaço e falta de ar
	PatientExplanation *string `gorm:"type:text" json:"patientExplanation,omitempty"`

	// Conduta clínica recomendada
	// @example Investigar causa da anemia (deficiência de ferro, B12, folato, doença crônica). Suplementação conforme indicação. Encaminhar ao hematologista se Hb < 10 g/dL ou causa não esclarecida
	Conduct *string `gorm:"type:text" json:"conduct,omitempty"`

	// Data da última revisão dos campos clínicos ou artigos associados
	// @example 2026-01-25T10:30:00Z
	LastReview *time.Time `gorm:"type:timestamptz" json:"lastReview,omitempty"`

	// @minimum 0
	// @maximum 100
	// @example 20
	Points *float64 `gorm:"type:double precision" json:"points,omitempty" validate:"omitempty,gte=0,lte=100"`

	// @minimum 0
	// @maximum 9999
	// @example 1
	Order int `gorm:"type:integer;not null;default:0;index:idx_score_item_order" json:"order" validate:"gte=0,lte=9999"`

	// ── Campos de exibição no SITE público (leigo) — fonte do gerador de score_versions ──
	// (O "conjunto Light" deixou de ser o flag is_light_version; agora é a score_version slug="light".)

	// Tipo de input renderizado no site público (≠ inferência usada na anamnese/exame do EMR).
	// @enum level_choice,numeric_classifier,boolean,text,scale_0_3
	// @example level_choice
	SiteRenderType *string `gorm:"type:varchar(40)" json:"siteRenderType,omitempty" validate:"omitempty,oneof=level_choice numeric_classifier boolean text scale_0_3"`

	// Pergunta reescrita para o site público (leigo). Substitui LightQuestion (migrado na 00021).
	// @example Você fuma atualmente?
	SiteQuestion *string `gorm:"type:text" json:"siteQuestion,omitempty"`

	// Explicação leiga genérica do item para o site público (≠ PatientExplanation clínica).
	// @example Fumar acelera o envelhecimento dos vasos e aumenta o risco cardiovascular.
	SiteExplanation *string `gorm:"type:text" json:"siteExplanation,omitempty"`

	// Ordem de exibição do item no formulário de preparação pré-consulta (subset curado,
	// independente de Order/LightOrder). Item entra no formulário quando PrepOrder != nil.
	// @minimum 0
	// @maximum 9999
	// @example 5
	PrepOrder *int `gorm:"type:integer;index:idx_score_item_prep_order" json:"prepOrder,omitempty" validate:"omitempty,gte=0,lte=9999"`

	// Foreign Keys
	// @example 550e8400-e29b-41d4-a716-446655440000
	SubgroupID uuid.UUID `gorm:"type:uuid;not null;index:idx_score_item_subgroup" json:"subgroupId" validate:"required"`

	// Self-referencing for hierarchical items (optional)
	// @example 550e8400-e29b-41d4-a716-446655440000
	ParentItemID *uuid.UUID `gorm:"type:uuid;index:idx_score_item_parent" json:"parentItemId,omitempty"`

	// Relationships
	Subgroup   *ScoreSubgroup `gorm:"foreignKey:SubgroupID;constraint:OnDelete:CASCADE" json:"subgroup,omitempty"`
	ParentItem *ScoreItem     `gorm:"foreignKey:ParentItemID;constraint:OnDelete:SET NULL" json:"parentItem,omitempty"`
	ChildItems []ScoreItem    `gorm:"foreignKey:ParentItemID;constraint:OnDelete:SET NULL" json:"childItems,omitempty"`
	Levels     []ScoreLevel   `gorm:"foreignKey:ItemID;constraint:OnDelete:CASCADE" json:"levels,omitempty"`

	// Many-to-many relationship with Articles
	// @items.type object
	Articles []Article `gorm:"many2many:article_score_items;" json:"articles,omitempty"`

	// Many-to-many relationship with MethodPillars
	// @items.type object
	MethodPillars []MethodPillar `gorm:"many2many:score_item_method_pillars;" json:"methodPillars,omitempty"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for ScoreItem
func (ScoreItem) TableName() string {
	return "score_items"
}

// BeforeCreate hook to generate UUID v7
func (si *ScoreItem) BeforeCreate(tx *gorm.DB) error {
	if si.ID == uuid.Nil {
		si.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// BeforeUpdate hook to update LastReview and invalidate embeddings when relevant fields change
func (si *ScoreItem) BeforeUpdate(tx *gorm.DB) error {
	// Update LastReview when clinical fields change
	if tx.Statement.Changed("ClinicalRelevance") ||
		tx.Statement.Changed("PatientExplanation") ||
		tx.Statement.Changed("Conduct") {
		now := time.Now()
		si.LastReview = &now
	}

	// Invalidate embedding when any field that affects semantic meaning changes.
	//
	// IMPORTANTE: a invalidação NÃO acontece aqui. O caminho de update real
	// (ScoreRepository.UpdateScoreItem) usa `Model(&ScoreItem{}).Updates(map)` com um
	// model VAZIO, então neste hook `si.ID == uuid.Nil` e `Changed(...)` é não-confiável
	// (compara contra zero-values). Disparar `invalidate_embedding` com ID nulo era um
	// no-op que ainda por cima ABORTAVA a transação em prod (ON CONFLICT/SQLSTATE 25P02),
	// derrubando qualquer edição de score item (ex.: indentar). A invalidação correta,
	// com o ID real e tolerante a erro, é feita em ScoreService.UpdateItem.
	//
	// O guard `si.ID != uuid.Nil` mantém o comportamento válido caso algum caminho futuro
	// atualize via struct completa (Save), sem reintroduzir o abort no caminho atual.
	if si.ID != uuid.Nil {
		embeddingFields := []string{
			"Name", "ClinicalRelevance", "PatientExplanation", "Conduct",
			"Gender", "AgeRangeMin", "AgeRangeMax", "PostMenopause", "Unit",
		}
		needsReembedding := false
		for _, field := range embeddingFields {
			if tx.Statement.Changed(field) {
				needsReembedding = true
				break
			}
		}
		if needsReembedding {
			tx.Exec(`UPDATE score_item_embeddings SET is_stale = true WHERE score_item_id = ?`, si.ID)
			tx.Exec(`SELECT invalidate_embedding('score_item', ?, 'Field update via BeforeUpdate hook', 0)`, si.ID)
			tx.Exec(`SELECT invalidate_preparation(?, 'ScoreItem field update')`, si.ID)
		}
	}

	return nil
}

// RequirementMet diz se a condição de contexto do item está satisfeita.
//
// `labValues` é o resultado mais recente de cada exame do paciente, por código. Item sem condição
// passa sempre. Item COM condição e sem o exame de referência medido **não passa**: sem o PSA total
// não dá para saber se a razão livre/total quer dizer alguma coisa, e afirmar que quer é pior do
// que omitir.
func (si *ScoreItem) RequirementMet(labValues map[string]float64) bool {
	if si.RequiresLabCode == nil || *si.RequiresLabCode == "" {
		return true
	}
	v, ok := labValues[*si.RequiresLabCode]
	if !ok {
		return false
	}
	if si.RequiresMin != nil && v < *si.RequiresMin {
		return false
	}
	if si.RequiresMax != nil && v > *si.RequiresMax {
		return false
	}
	return true
}

// UnitMatches diz se a escala deste item e o exame falam da MESMA grandeza.
//
// Três itens do sedimento urinário têm escala em `células/campo` (contagem por campo de
// microscópio) enquanto o laboratório reporta `/µL` (concentração). Classificar assim compara
// números que não se comparam: 0,5/µL cai na faixa "≤10 células/campo" e sai como nível ÓTIMO.
// Falha silenciosa, e a régua da devolutiva a mostraria ao paciente como notícia boa.
//
// Unidade vazia dos dois lados não bloqueia: item categórico não tem unidade.
// `sinonimosDoExame` são os pares (unidade principal, unidade secundária) com fator de conversão
// 1 que o catálogo registra PARA ESTE EXAME, vindos de `lab_test_unit_conversions`. Eles cobrem o
// que nenhuma regra de string alcança porque depende do analito: `mEq/L` só é igual a `mmol/L` em
// íon monovalente (Na⁺, K⁺, HCO₃⁻); em cálcio seria o dobro. Passar nil desliga essa consulta e
// deixa só as equivalências mecânicas.
func (si *ScoreItem) UnitMatches(unidadeDoExame string, sinonimosDoExame [][2]string) bool {
	item := ""
	if si.Unit != nil {
		item = *si.Unit
	}
	return MesmaGrandeza(item, unidadeDoExame, sinonimosDoExame)
}

// MesmaGrandeza diz se duas unidades expressam a mesma coisa, sem conversão de escala. É o teste
// que a guarda do escore e o passo 1 do conversor compartilham — duplicar isso foi o que deixou
// `mm` e `mm/hr` (VHS) passando num lugar e sendo recusados no outro.
//
// Unidade vazia de qualquer lado não bloqueia: item categórico não tem unidade, e laudo sem
// unidade declarada não é motivo para deixar de avaliar.
func MesmaGrandeza(a, b string, sinonimosDoExame [][2]string) bool {
	na, nb := NormalizaUnidade(a), NormalizaUnidade(b)
	if na == "" || nb == "" {
		return true
	}
	if na == nb {
		return true
	}
	if unidadesEquivalentes[[2]string{na, nb}] || unidadesEquivalentes[[2]string{nb, na}] {
		return true
	}
	for _, par := range sinonimosDoExame {
		x, y := NormalizaUnidade(par[0]), NormalizaUnidade(par[1])
		if (x == na && y == nb) || (x == nb && y == na) {
			return true
		}
	}
	return false
}

// unidadesEquivalentes lista pares que a normalização não junta mas que valem o MESMO número,
// por definição e para QUALQUER analito. Só entra aqui igualdade exata e independente de peso
// molecular ou valência: o que depende do analito mora em `lab_test_unit_conversions`.
var unidadesEquivalentes = map[[2]string]bool{
	// 1 mIU/L = 10⁻³ IU/L = 10⁻⁶ IU/mL = 1 µIU/mL. Aparece no TSH.
	{"miu/l", "uiu/ml"}: true,
	// 1 mIU/mL = 10⁻³ IU/mL = 1 IU/L. Aparece no anti-HBs, FSH, LH.
	{"miu/ml", "iu/l"}: true,
	// VHS reportada só como "mm": o laboratório omite o /h, a grandeza é a mesma.
	{"mm/h", "mm"}: true,
}

// numeradorEmUI casa um numerador que é unidade internacional em qualquer grafia (`U`, `UI`,
// `IU`), com prefixo SI opcional: `U`, `mUI`, `µIU`, `kU`.
var numeradorEmUI = regexp.MustCompile(`^(m|u|k|n|p)?(?:ui|iu|u)$`)

// NormalizaUnidade reduz uma unidade à forma canônica de comparação. O catálogo e os laudos
// escrevem a mesma grandeza de vários jeitos (`mcg/dL`, `µg/dL`, `ug/dL`) e comparar string crua
// acusaria divergência onde não há — o que silenciaria item correto.
//
// Só faz equivalências EXATAS, nunca conversão de escala: `mm³` é por definição 1 µL, `UI` é a
// grafia portuguesa de `IU`, `mc` é a grafia ASCII de `µ`. `mg/dL` e `g/dL` continuam diferentes,
// que é justamente o que a guarda precisa pegar.
func NormalizaUnidade(u string) string {
	s := strings.ToLower(strings.TrimSpace(u))

	// Qualificador textual: "mg/g de creatinina" é mg/g.
	if i := strings.Index(s, " de "); i > 0 {
		s = s[:i]
	}
	s = strings.ReplaceAll(s, " ", "")

	// Micro: sinal U+00B5, mu grego U+03BC e a abreviação `mc` são todos µ.
	s = strings.ReplaceAll(s, "\u00b5", "u")
	s = strings.ReplaceAll(s, "\u03bc", "u")
	s = strings.ReplaceAll(s, "mc", "u")

	// Unidade internacional: `UI` (pt), `IU` (en) e `U` sozinho são a mesma coisa nos ensaios
	// em que aparecem — o laudo escreve `U/mL` onde o catálogo escreve `UI/mL`
	// (anti-transglutaminase) e `µU/mL` onde escreve `µUI/mL` (TSH, insulina).
	//
	// A troca é no NUMERADOR inteiro, com prefixo SI opcional, nunca por substring: um
	// `strings.ReplaceAll("ui","iu")` cego transformaria `µUI/mL` em `uiu/ml` e `µIU/mL` em
	// `iuu/ml`, separando duas grafias da mesma unidade. E estragaria `ug/dL` e `/µL`, que não
	// têm nada a ver com unidade internacional.
	if i := strings.Index(s, "/"); i > 0 {
		if num := numeradorEmUI.ReplaceAllString(s[:i], "${1}iu"); num != s[:i] {
			s = num + s[i:]
		}
	}

	// Vírgula decimal e expoente por extenso: a TFG sai como `mL/min/1,73m2`, `mL/min/1.73m²`
	// e `mL/min/1,73 m²` conforme o laboratório. É a mesma superfície corporal nas três.
	s = strings.ReplaceAll(s, ",", ".")
	s = strings.ReplaceAll(s, "²", "2")

	// Hora por extenso ou abreviada (VHS sai como `mm/h`, `mm/hr` e `mm/Hora`).
	s = strings.ReplaceAll(s, "/hora", "/h")
	s = strings.ReplaceAll(s, "/hr", "/h")

	// 1 dL = 100 mL, por definição: o laudo de chumbo escreve `µg/100 mL`.
	s = strings.ReplaceAll(s, "/100ml", "/dl")

	// 1 mm³ = 1 µL, exato.
	s = strings.ReplaceAll(s, "mm\u00b3", "ul")
	s = strings.ReplaceAll(s, "mm3", "ul")

	// Multiplicadores escritos por extenso, como o hemograma faz.
	for _, mil := range []string{"x10\u00b3", "x10^3", "10\u00b3", "mil"} {
		s = strings.ReplaceAll(s, mil+"/", "k/")
	}
	for _, milhao := range []string{"milh\u00f5es", "milhoes"} {
		s = strings.ReplaceAll(s, milhao+"/", "m/")
	}

	return s
}

// AppliesToPatient verifica se este ScoreItem se aplica ao paciente baseado em gênero, idade e menopausa
func (si *ScoreItem) AppliesToPatient(patient *Patient) bool {
	return si.MotivoDeNaoAplicar(patient) == ""
}

// MotivoDeNaoAplicar diz, em português de prontuário, por que este item não vale para este
// paciente — ou "" se vale. É a MESMA lógica de AppliesToPatient, e de propósito: antes o motivo
// era montado à parte, olhando só se o item tinha sexo declarado, e saía "sexo female requerido
// (paciente: female)" em 43 itens de uma paciente. A frase se contradizia porque a causa real era
// outra (menopausa, TRH, faixa etária) e ninguém tinha como saber lendo o prontuário.
func (si *ScoreItem) MotivoDeNaoAplicar(patient *Patient) string {
	if si.Gender != nil && *si.Gender != "not_applicable" && *si.Gender != string(patient.Gender) {
		return "Item não aplicável: é de paciente do sexo " + sexoPorExtenso(*si.Gender)
	}

	// Idade desconhecida não é idade zero. `CalculateAge` deixa Age em 0 quando não há data de
	// nascimento, e a recepção cadastra paciente só pelo nome de propósito. Comparar 0 com a
	// faixa faz o item "até 29 anos" valer para um senhor de 70 e o "a partir de 50" sumir:
	// o paciente acaba pontuado na faixa etária errada, não só descrito errado.
	if si.AgeRangeMin != nil || si.AgeRangeMax != nil {
		if patient.BirthDate.IsZero() {
			return "Item não avaliado: depende da idade, e a data de nascimento não está no cadastro"
		}
		if si.AgeRangeMin != nil && patient.Age < *si.AgeRangeMin {
			return fmt.Sprintf("Item não aplicável: vale a partir de %d anos (paciente tem %d)", *si.AgeRangeMin, patient.Age)
		}
		if si.AgeRangeMax != nil && patient.Age > *si.AgeRangeMax {
			return fmt.Sprintf("Item não aplicável: vale até %d anos (paciente tem %d)", *si.AgeRangeMax, patient.Age)
		}
	}

	// Menopausa e reposição hormonal só filtram mulheres, e item que não declara a condição não
	// filtra por ela. Quando o item declara e a paciente não tem o dado, o item não é "não
	// aplicável": é indeterminado, e dizer isso é o que faz a recepção preencher o cadastro.
	if patient.Gender == "female" && si.PostMenopause != nil {
		if patient.Menopause == nil {
			return "Item não avaliado: depende do estado de menopausa, que não está registrado no cadastro"
		}
		if *si.PostMenopause != *patient.Menopause {
			if *si.PostMenopause {
				return "Item não aplicável: é de paciente na pós-menopausa"
			}
			return "Item não aplicável: é de paciente antes da menopausa"
		}
	}

	if patient.Gender == "female" && si.HormoneTherapy != nil {
		if patient.HormoneTherapy == nil {
			return "Item não avaliado: depende de saber se faz reposição hormonal, que não está registrado no cadastro"
		}
		if *si.HormoneTherapy != *patient.HormoneTherapy {
			if *si.HormoneTherapy {
				return "Item não aplicável: é de paciente em reposição hormonal"
			}
			return "Item não aplicável: é de paciente sem reposição hormonal"
		}
	}

	return ""
}

func sexoPorExtenso(g string) string {
	switch g {
	case "male":
		return "masculino"
	case "female":
		return "feminino"
	default:
		return g
	}
}

// GetFullName retorna o nome completo do ScoreItem no formato:
// "ItemName (Group - Subgroup - ParentItem)" (se houver parent e relações carregadas)
// "ItemName (Group - Subgroup)" (se não houver parent mas relações carregadas)
// "ItemName" (se as relações não estiverem carregadas)
func (si *ScoreItem) GetFullName() string {
	// Se as relações não estiverem carregadas, retorna apenas o nome do item
	if si.Subgroup == nil || si.Subgroup.Group == nil {
		return si.Name
	}

	groupName := si.Subgroup.Group.Name
	subgroupName := si.Subgroup.Name

	// Se houver ParentItem, inclui no nome
	if si.ParentItemID != nil && si.ParentItem != nil {
		return si.Name + " (" + groupName + " - " + subgroupName + " - " + si.ParentItem.Name + ")"
	}

	return si.Name + " (" + groupName + " - " + subgroupName + ")"
}
