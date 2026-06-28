package services

import (
	"encoding/json"
	"sort"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

// ClinicalTimelineService monta a linha do tempo longitudinal de um ScoreItem para um
// paciente, fundindo (read-side) as três fontes que registram resposta de ScoreItem com a
// mesma forma (nível/numérico/texto):
//
//   - Escore Light  → anonymous_score_items   (sessão claimed pelo paciente)
//   - Pré-consulta  → consultation_prep_responses (preps submetidos do paciente)
//   - Anamnese      → anamnesis_items          (anamneses do paciente)
//
// Não há migration nem tabela nova: cada fonte permanece como está; aqui só lemos e
// ordenamos por data (mais recente primeiro). Ver docs/emr/estudo-simplificacao-consulta-anamnese.md (§3, abordagem B).
type ClinicalTimelineService struct {
	db *gorm.DB
}

func NewClinicalTimelineService(db *gorm.DB) *ClinicalTimelineService {
	return &ClinicalTimelineService{db: db}
}

// rawTimelineRow é o formato cru lido de cada fonte (Scan).
type rawTimelineRow struct {
	Date          time.Time
	SelectedLevel *int
	NumericValue  *float64
	TextValue     *string
	// Só a fonte anamnese seleciona este campo (jsonb::text); demais ficam nil.
	ScaleResponsesJSON *string
}

// parseScaleResponses desserializa o jsonb scale_responses (lido como texto) em struct.
func parseScaleResponses(s *string) *models.ScaleResponseData {
	if s == nil || *s == "" {
		return nil
	}
	var sr models.ScaleResponseData
	if err := json.Unmarshal([]byte(*s), &sr); err != nil {
		return nil
	}
	return &sr
}

// GetItemHistory retorna o histórico de um ScoreItem para um paciente, de todas as origens,
// ordenado da entrada mais recente para a mais antiga.
func (s *ClinicalTimelineService) GetItemHistory(patientID, scoreItemID uuid.UUID) ([]dto.ClinicalTimelineEntry, error) {
	entries := make([]dto.ClinicalTimelineEntry, 0, 8)

	// 1) Anamnese — itens registrados pela equipe (consultation_date da anamnese).
	var anamRows []rawTimelineRow
	if err := s.db.Raw(`
		SELECT a.consultation_date AS date, ai.selected_level, ai.numeric_value, ai.text_value,
		       ai.scale_responses::text AS scale_responses_json
		FROM anamnesis_items ai
		JOIN anamnesis a ON a.id = ai.anamnesis_id AND a.deleted_at IS NULL
		WHERE ai.score_item_id = ? AND a.patient_id = ? AND ai.deleted_at IS NULL
	`, scoreItemID, patientID).Scan(&anamRows).Error; err != nil {
		return nil, err
	}
	for _, r := range anamRows {
		entries = append(entries, dto.ClinicalTimelineEntry{
			Date:           r.Date,
			Source:         dto.ClinicalSourceAnamnese,
			SourceLabel:    "Anamnese",
			SelectedLevel:  r.SelectedLevel,
			NumericValue:   r.NumericValue,
			TextValue:      r.TextValue,
			ScaleResponses: parseScaleResponses(r.ScaleResponsesJSON),
		})
	}

	// 2) Pré-consulta — respostas de preps SUBMETIDOS (drafts não entram no histórico).
	var prepRows []rawTimelineRow
	if err := s.db.Raw(`
		SELECT COALESCE(p.submitted_at, p.created_at) AS date, r.selected_level, r.numeric_value, r.text_value
		FROM consultation_prep_responses r
		JOIN consultation_preps p ON p.id = r.prep_id AND p.deleted_at IS NULL
		WHERE r.score_item_id = ? AND p.patient_id = ? AND r.deleted_at IS NULL AND p.status = 'submitted'
	`, scoreItemID, patientID).Scan(&prepRows).Error; err != nil {
		return nil, err
	}
	for _, r := range prepRows {
		entries = append(entries, dto.ClinicalTimelineEntry{
			Date:          r.Date,
			Source:        dto.ClinicalSourcePreConsulta,
			SourceLabel:   "Pré-consulta",
			SelectedLevel: r.SelectedLevel,
			NumericValue:  r.NumericValue,
			TextValue:     r.TextValue,
		})
	}

	// 3) Escore Light — itens de sessões anônimas vinculadas (claimed) a este paciente.
	var lightRows []rawTimelineRow
	if err := s.db.Raw(`
		SELECT s.created_at AS date, i.selected_level, i.numeric_value, i.text_value
		FROM anonymous_score_items i
		JOIN anonymous_score_sessions s ON s.id = i.session_id AND s.deleted_at IS NULL
		WHERE i.score_item_id = ? AND s.claimed_by_patient_id = ? AND i.deleted_at IS NULL
	`, scoreItemID, patientID).Scan(&lightRows).Error; err != nil {
		return nil, err
	}
	for _, r := range lightRows {
		entries = append(entries, dto.ClinicalTimelineEntry{
			Date:          r.Date,
			Source:        dto.ClinicalSourceEscoreLight,
			SourceLabel:   "Escore Light",
			SelectedLevel: r.SelectedLevel,
			NumericValue:  r.NumericValue,
			TextValue:     r.TextValue,
		})
	}

	// Mais recente primeiro.
	sort.SliceStable(entries, func(i, j int) bool {
		return entries[i].Date.After(entries[j].Date)
	})

	return entries, nil
}
