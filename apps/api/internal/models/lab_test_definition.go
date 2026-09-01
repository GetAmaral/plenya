package models

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/google/uuid"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"gorm.io/gorm"
)

// LabTestCategory define categorias de exames laboratoriais
type LabTestCategory string

const (
	LabTestCategoryHematology   LabTestCategory = "hematology"   // Hemograma, coagulação
	LabTestCategoryBiochemistry LabTestCategory = "biochemistry" // Glicose, ureia, creatinina
	LabTestCategoryHormones     LabTestCategory = "hormones"     // TSH, T4, testosterona
	LabTestCategoryImmunology   LabTestCategory = "immunology"   // Sorologias, autoimunes
	LabTestCategoryMicrobiology LabTestCategory = "microbiology" // Culturas, antibiogramas
	LabTestCategoryUrine        LabTestCategory = "urine"        // EAS, urocultura
	LabTestCategoryImaging      LabTestCategory = "imaging"      // Raio-X, TC, RM
	LabTestCategoryFunctional   LabTestCategory = "functional"   // Medicina funcional
	LabTestCategoryGenetics     LabTestCategory = "genetics"     // Testes genéticos
	LabTestCategoryOther        LabTestCategory = "other"        // Outros
)

// LabTestDefinition representa a definição de um exame laboratorial ou parâmetro
// @Description Definição estruturada de exames e seus parâmetros
type LabTestDefinition struct {
	// ID único da definição
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Código interno único (usado para referência programática)
	// @example HEMOGRAMA_COMPLETO, HGB, GLUCOSE_FASTING
	Code string `gorm:"type:varchar(100);not null;unique;index" json:"code" validate:"required,max=100"`

	// Nome do exame/parâmetro
	// @example Hemograma Completo, Hemoglobina, Glicemia de Jejum
	Name string `gorm:"type:varchar(300);not null" json:"name" validate:"required,max=300"`

	// Nome curto/abreviação (opcional) - usado para exibição
	// @example Hemograma, Hb, Gli
	ShortName *string `gorm:"type:varchar(50)" json:"shortName,omitempty"`

	// Nomes alternativos para matching em PDFs (sempre inclui Name também)
	// Array de strings com variações do nome encontradas em laudos
	// NUNCA incluir siglas (siglas vão em ShortName)
	// @example ["Hemoglobina", "Hemoglobina total", "Hemoglobina - Homens", "Hemoglobina - Mulheres"]
	AltNames []string `gorm:"type:jsonb;serializer:json" json:"altNames,omitempty"`

	// Código TUSS (Terminologia Unificada da Saúde Suplementar)
	// Usado para faturamento e solicitação no Brasil
	// @example 40304485
	TussCode *string `gorm:"type:varchar(20);index" json:"tussCode,omitempty"`

	// Código LOINC (Logical Observation Identifiers Names and Codes)
	// Padrão internacional para identificação de exames
	// @example 718-7 (Hemoglobin)
	LoincCode *string `gorm:"type:varchar(20);index" json:"loincCode,omitempty"`

	// Categoria do exame
	// @enum hematology,biochemistry,hormones,immunology,microbiology,urine,imaging,functional,genetics,other
	Category LabTestCategory `gorm:"type:varchar(30);not null;index" json:"category" validate:"required"`

	// Aplicabilidade por sexo biológico do paciente: 'all' (default), 'male', 'female'.
	// Filtra exames sexo-específicos (PSA, CA-125, mamografia...) ao carregar um template.
	// @enum all,male,female
	SexApplicability string `gorm:"type:varchar(10);not null;default:'all'" json:"sexApplicability"`

	// Indica se o exame pode ser solicitado individualmente
	// true: pode ser solicitado (ex: Hemograma Completo, Glicemia)
	// false: só aparece como resultado de outro exame (ex: Hemoglobina, Bilirrubina Indireta)
	IsRequestable bool `gorm:"type:boolean;not null;default:true;index" json:"isRequestable"`

	// ID do exame pai (hierarquia)
	// Ex: Hemoglobina tem parentTestId = Hemograma Completo
	// Ex: Bilirrubina Indireta tem parentTestId = Bilirrubina Total e Frações
	ParentTestID *uuid.UUID `gorm:"type:uuid;index" json:"parentTestId,omitempty"`

	// Unidade de medida padrão
	// @example g/dL, mg/dL, mU/L, %
	Unit *string `gorm:"type:varchar(50)" json:"unit,omitempty"`

	// Fórmula de conversão entre unidades (se aplicável)
	// @example 1 g/dL = 10 g/L
	UnitConversion *string `gorm:"type:text" json:"unitConversion,omitempty"`

	// Tipo de resultado
	// @enum numeric, text, boolean, categorical
	ResultType string `gorm:"type:varchar(20);not null;default:'numeric'" json:"resultType" validate:"required"`

	// Método de coleta/realização
	// @example Sangue venoso com jejum de 8-12h
	CollectionMethod *string `gorm:"type:text" json:"collectionMethod,omitempty"`

	// Tempo de jejum necessário (em horas)
	FastingHours *int `gorm:"type:integer" json:"fastingHours,omitempty"`

	// Material biológico
	// @example Sangue total, Soro, Urina, Fezes
	SpecimenType *string `gorm:"type:varchar(100)" json:"specimenType,omitempty"`

	// Descrição/observações sobre o exame
	Description *string `gorm:"type:text" json:"description,omitempty"`

	// Indicações clínicas principais
	ClinicalIndications *string `gorm:"type:text" json:"clinicalIndications,omitempty"`

	// Justificativa clínica padrão que acompanha o exame no pedido.
	// Carrega automaticamente quando o exame é selecionado em qualquer painel
	// (uso primário: exames de alto custo que exigem fundamentação p/ autorização).
	RequestJustification *string `gorm:"type:text" json:"requestJustification,omitempty"`

	// Significância clínica detalhada (200-400 palavras)
	// Mecanismos fisiológicos, aplicações clínicas, interpretação de valores alterados
	ClinicalSignificance *string `gorm:"type:text" json:"clinicalSignificance,omitempty"`

	// Contexto de longevidade (100-200 palavras)
	// Relação com envelhecimento saudável, marcadores de longevidade, implicações preventivas
	LongevityContext *string `gorm:"type:text" json:"longevityContext,omitempty"`

	// Recomendações clínicas (150-300 palavras)
	// Quando solicitar, interpretação de resultados, fatores que afetam valores, intervenções
	ClinicalRecommendations *string `gorm:"type:text" json:"clinicalRecommendations,omitempty"`

	// Ordem de exibição (para organizar parâmetros de um exame)
	DisplayOrder int `gorm:"type:integer;not null;default:0" json:"displayOrder"`

	// PageBreakBefore — campo TRANSIENTE (não persiste): preenchido só quando o exame é devolvido
	// dentro de um template (vem do join lab_request_template_tests.page_break_before). Indica que,
	// ao aplicar o template, este exame começa em nova página (linha em branco antes).
	PageBreakBefore bool `gorm:"-" json:"pageBreakBefore,omitempty"`

	// Status (ativo/inativo)
	IsActive bool `gorm:"type:boolean;not null;default:true;index" json:"isActive"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relacionamentos

	// Exame pai (se for um parâmetro/subexame)
	ParentTest *LabTestDefinition `gorm:"foreignKey:ParentTestID;constraint:OnDelete:SET NULL" json:"parentTest,omitempty"`

	// Subexames/parâmetros (se for um exame composto)
	SubTests []LabTestDefinition `gorm:"foreignKey:ParentTestID;constraint:OnDelete:SET NULL" json:"subTests,omitempty"`

	// ScoreItems relacionados via código (oneToMany: labTestDefinition.code -> scoreItem.labTestCode)
	// Múltiplos ScoreItems podem ter o mesmo código (ex: variantes por gênero/idade)
	ScoreItems []ScoreItem `gorm:"foreignKey:LabTestCode;references:Code" json:"scoreItems,omitempty"`
}

// TableName especifica o nome da tabela
func (LabTestDefinition) TableName() string {
	return "lab_test_definitions"
}

// BeforeCreate hook to generate UUID v7
func (ltd *LabTestDefinition) BeforeCreate(tx *gorm.DB) error {
	if ltd.ID == uuid.Nil {
		ltd.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// BeforeSave hook to normalize AltNames (lower + unaccent)
func (ltd *LabTestDefinition) BeforeSave(tx *gorm.DB) error {
	// Normalizar AltNames se não for nil/vazio
	if len(ltd.AltNames) > 0 {
		normalized := make([]string, 0, len(ltd.AltNames))
		for _, name := range ltd.AltNames {
			normalizedName := normalizeString(name)
			if normalizedName != "" {
				normalized = append(normalized, normalizedName)
			}
		}
		ltd.AltNames = normalized
	}
	return nil
}

// normalizeString aplica trim + lower + unaccent em uma string
func normalizeString(s string) string {
	// 1. Trim espaços
	s = strings.TrimSpace(s)

	// 2. Lower case
	s = strings.ToLower(s)

	// 3. Unaccent (remover acentos)
	s = removeAccents(s)

	return s
}

// removeAccents remove acentos de uma string (equivalente ao unaccent do PostgreSQL)
func removeAccents(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, s)
	return result
}

// StatusDeConversao diz o que aconteceu com a unidade de um resultado, para o registro poder
// ser marcado. Sem isto, um resultado que não converteu é indistinguível de um que já chegou
// certo — foi assim que 124 resultados ficaram gravados na unidade do laudo sem ninguém saber.
type StatusDeConversao string

const (
	// ConversaoDesnecessaria: o laudo já veio na unidade do exame (ou numa grafia dela).
	ConversaoDesnecessaria StatusDeConversao = "ok"
	// ConversaoAplicada: o valor foi convertido, por regra curada ou por aritmética de prefixo.
	ConversaoAplicada StatusDeConversao = "convertido"
	// ConversaoPendente: as unidades diferem e não há como converter com segurança. O valor
	// fica como veio e o registro entra na fila de revisão.
	ConversaoPendente StatusDeConversao = "revisar"
)

// ResultadoDaConversao é o que ConverteParaUnidadePrincipal devolve.
type ResultadoDaConversao struct {
	Valor   float64
	Unidade string
	Status  StatusDeConversao
	Motivo  string
}

// ConverteParaUnidadePrincipal leva um valor para a unidade do exame, em quatro camadas, da mais
// confiável para a menos:
//
//  1. grafia — `mcg/dL` e `µg/dL` são a mesma unidade, não há o que converter;
//  2. tabela curada (`lab_test_unit_conversions`) — é onde mora o que depende do analito, como
//     `mEq/L` = `mmol/L` só em íon monovalente;
//  3. aritmética de prefixo SI — `pg/mL` para `ng/mL` é dividir por mil, e isso não precisa de
//     curadoria nenhuma;
//  4. desiste, mas AVISA.
//
// `plausivel` é a rede de segurança da camada 3: aritmética correta em cima de rótulo errado
// produz número confiantemente errado. Um resultado de hemácias marcado `/mm³` cujo valor está
// de fato em `M/µL` viraria 4,17 milionésimos e sairia como anemia catastrófica. Quando a
// conversão cai fora do que aquele exame pode valer, o valor fica como está e vai para revisão.
// Passar nil desliga a checagem.
func (ltd *LabTestDefinition) ConverteParaUnidadePrincipal(
	db *gorm.DB,
	valorOriginal float64,
	unidadeOriginal string,
	plausivel func(float64) bool,
) ResultadoDaConversao {
	principal := ""
	if ltd.Unit != nil {
		principal = *ltd.Unit
	}
	semConversao := ResultadoDaConversao{Valor: valorOriginal, Unidade: unidadeOriginal}

	// Sem unidade de um dos lados não há o que comparar.
	if strings.TrimSpace(principal) == "" || strings.TrimSpace(unidadeOriginal) == "" {
		semConversao.Status = ConversaoDesnecessaria
		return semConversao
	}

	// 1. Mesma unidade, escrita de outro jeito. Mesmo teste que a guarda do escore usa.
	if MesmaGrandeza(unidadeOriginal, principal, nil) {
		return ResultadoDaConversao{Valor: valorOriginal, Unidade: principal, Status: ConversaoDesnecessaria}
	}

	// 2. Regra curada para este exame.
	if db == nil {
		return ltd.porAritmetica(valorOriginal, unidadeOriginal, principal, plausivel)
	}
	var conversao LabTestUnitConversion
	err := db.Where(
		"lab_test_definition_id = ? AND LOWER(TRIM(secondary_unit)) = ? AND deleted_at IS NULL",
		ltd.ID,
		strings.ToLower(strings.TrimSpace(unidadeOriginal)),
	).First(&conversao).Error
	if err == nil {
		return ResultadoDaConversao{
			Valor:   conversao.ConvertToMain(valorOriginal),
			Unidade: conversao.MainUnit,
			Status:  ConversaoAplicada,
			Motivo:  "regra curada do catálogo",
		}
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		semConversao.Status = ConversaoPendente
		semConversao.Motivo = "falha ao consultar as conversões do exame"
		return semConversao
	}

	// 3. Aritmética de prefixo.
	return ltd.porAritmetica(valorOriginal, unidadeOriginal, principal, plausivel)
}

// porAritmetica é a camada 3 (e o desfecho 4) isolada, para poder ser exercitada sem banco.
func (ltd *LabTestDefinition) porAritmetica(
	valorOriginal float64,
	unidadeOriginal, principal string,
	plausivel func(float64) bool,
) ResultadoDaConversao {
	semConversao := ResultadoDaConversao{Valor: valorOriginal, Unidade: unidadeOriginal}

	if fator, ok := FatorEntreUnidades(unidadeOriginal, principal); ok {
		convertido := valorOriginal * fator
		if plausivel == nil || plausivel(convertido) {
			return ResultadoDaConversao{
				Valor:   convertido,
				Unidade: principal,
				Status:  ConversaoAplicada,
				Motivo:  "aritmética de prefixo",
			}
		}
		semConversao.Status = ConversaoPendente
		semConversao.Motivo = fmt.Sprintf(
			"laudo em %s e exame em %s: a conversão daria %s, fora do que este exame pode valer — o rótulo da unidade é que parece errado",
			unidadeOriginal, principal, formataNumero(convertido))
		return semConversao
	}

	// 4. Não dá para converter com segurança.
	semConversao.Status = ConversaoPendente
	semConversao.Motivo = fmt.Sprintf(
		"laudo em %s e exame em %s: grandezas diferentes, sem regra de conversão cadastrada",
		unidadeOriginal, principal)
	return semConversao
}

// formataNumero imprime sem zero à toa nem notação científica.
func formataNumero(v float64) string {
	if math.Abs(v-math.Round(v)) < 1e-9 {
		return strconv.FormatInt(int64(math.Round(v)), 10)
	}
	return strings.Replace(strconv.FormatFloat(v, 'g', 4, 64), ".", ",", 1)
}
