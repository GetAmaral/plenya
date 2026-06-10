package dto

import "time"

// ClinicalTimelineSource identifica a origem de uma entrada do histórico de um item clínico.
const (
	ClinicalSourceEscoreLight = "escore_light" // Escore Plenya Light (sessão anônima, claimed)
	ClinicalSourcePreConsulta = "pre_consulta" // Formulário de preparação pré-consulta
	ClinicalSourceAnamnese    = "anamnese"     // Anamnese (itens registrados pela equipe na consulta)
)

// ClinicalTimelineEntry é um ponto na linha do tempo de um ScoreItem para um paciente,
// independente da origem (Escore Light, pré-consulta ou anamnese/consulta).
//
// @Description Entrada do histórico longitudinal de um item clínico
type ClinicalTimelineEntry struct {
	// Data do registro (consultation_date da anamnese, submit do prep, criação da sessão light)
	Date time.Time `json:"date"`

	// Origem: escore_light | pre_consulta | anamnese
	Source string `json:"source"`

	// Rótulo legível da origem (ex.: "Escore Light", "Pré-consulta", "Anamnese")
	SourceLabel string `json:"sourceLabel"`

	// Nível classificado (0-6), quando houver
	SelectedLevel *int `json:"selectedLevel,omitempty"`

	// Valor numérico medido, quando houver
	NumericValue *float64 `json:"numericValue,omitempty"`

	// Valor textual, quando houver
	TextValue *string `json:"textValue,omitempty"`
}

// ClinicalItemHistoryResponse é a resposta do histórico de um item clínico.
//
// @Description Histórico longitudinal de um item clínico (todas as origens)
type ClinicalItemHistoryResponse struct {
	ScoreItemID string                  `json:"scoreItemId"`
	Entries     []ClinicalTimelineEntry `json:"entries"`
}
