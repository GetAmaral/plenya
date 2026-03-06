package services

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// ClinicalDataService centraliza o acesso a dados clínicos do paciente
// obtidos de AnamnesisItems (via score_items.anamnese_item_code) e
// LabResults (via lab_test_definitions.code).
//
// Todos os serviços que precisam de dados de anamnese ou laboratório
// devem usar este serviço ao invés de queries diretas.
type ClinicalDataService struct {
	db *gorm.DB
}

func NewClinicalDataService(db *gorm.DB) *ClinicalDataService {
	return &ClinicalDataService{db: db}
}

// --- Anamnesis ---

// LatestAnamnesisNumeric retorna o numeric_value mais recente de um AnamnesisItem
// identificado pelo anamnese_item_code do ScoreItem vinculado.
func (s *ClinicalDataService) LatestAnamnesisNumeric(patientID uuid.UUID, code models.AnamneseItemCode) *float64 {
	var item struct {
		NumericValue *float64
	}
	s.db.Model(&models.AnamnesisItem{}).
		Select("anamnesis_items.numeric_value").
		Joins("JOIN score_items ON anamnesis_items.score_item_id = score_items.id").
		Joins("JOIN anamnesis ON anamnesis_items.anamnesis_id = anamnesis.id").
		Where("anamnesis.patient_id = ? AND score_items.anamnese_item_code = ?", patientID, code).
		Where("anamnesis_items.numeric_value IS NOT NULL").
		Where("anamnesis_items.deleted_at IS NULL AND anamnesis.deleted_at IS NULL").
		Order("anamnesis.consultation_date DESC").
		Limit(1).
		Scan(&item)
	return item.NumericValue
}

// LatestAnamnesisLevel retorna o selected_level mais recente de um AnamnesisItem.
func (s *ClinicalDataService) LatestAnamnesisLevel(patientID uuid.UUID, code models.AnamneseItemCode) *int {
	var item struct {
		SelectedLevel *int
	}
	s.db.Model(&models.AnamnesisItem{}).
		Select("anamnesis_items.selected_level").
		Joins("JOIN score_items ON anamnesis_items.score_item_id = score_items.id").
		Joins("JOIN anamnesis ON anamnesis_items.anamnesis_id = anamnesis.id").
		Where("anamnesis.patient_id = ? AND score_items.anamnese_item_code = ?", patientID, code).
		Where("anamnesis_items.selected_level IS NOT NULL").
		Where("anamnesis_items.deleted_at IS NULL AND anamnesis.deleted_at IS NULL").
		Order("anamnesis.consultation_date DESC").
		Limit(1).
		Scan(&item)
	return item.SelectedLevel
}

// LatestAnamnesisText retorna o text_value mais recente de um AnamnesisItem.
func (s *ClinicalDataService) LatestAnamnesisText(patientID uuid.UUID, code models.AnamneseItemCode) *string {
	var item struct {
		TextValue *string
	}
	s.db.Model(&models.AnamnesisItem{}).
		Select("anamnesis_items.text_value").
		Joins("JOIN score_items ON anamnesis_items.score_item_id = score_items.id").
		Joins("JOIN anamnesis ON anamnesis_items.anamnesis_id = anamnesis.id").
		Where("anamnesis.patient_id = ? AND score_items.anamnese_item_code = ?", patientID, code).
		Where("anamnesis_items.text_value IS NOT NULL AND anamnesis_items.text_value != ''").
		Where("anamnesis_items.deleted_at IS NULL AND anamnesis.deleted_at IS NULL").
		Order("anamnesis.consultation_date DESC").
		Limit(1).
		Scan(&item)
	return item.TextValue
}

// --- Lab Results ---

// LatestLabNumeric retorna o result_numeric mais recente de um LabResult
// identificado pelo code do LabTestDefinition.
func (s *ClinicalDataService) LatestLabNumeric(patientID uuid.UUID, code models.LabTestCode) *float64 {
	var result struct {
		ResultNumeric *float64
	}
	s.db.Model(&models.LabResult{}).
		Select("lab_results.result_numeric").
		Joins("JOIN lab_test_definitions ON lab_results.lab_test_definition_id = lab_test_definitions.id").
		Joins("JOIN lab_result_batches ON lab_results.lab_result_batch_id = lab_result_batches.id").
		Where("lab_result_batches.patient_id = ? AND lab_test_definitions.code = ?", patientID, code).
		Where("lab_results.result_numeric IS NOT NULL AND lab_results.deleted_at IS NULL").
		Order("lab_result_batches.collection_date DESC").
		Limit(1).
		Scan(&result)
	return result.ResultNumeric
}

// --- Helpers de alto nível ---

// SmokingStatusFromLevel converte selected_level do ScoreItem TABACO para SmokingStatus.
// Levels: 0-1 = current, 2-3 = former, 4-5 = never
func SmokingStatusFromLevel(level int) models.SmokingStatus {
	switch {
	case level <= 1:
		return models.SmokingCurrent
	case level <= 3:
		return models.SmokingFormer
	default:
		return models.SmokingNever
	}
}

// ActivityLevelFromLevel converte selected_level do ScoreItem ATIVIDADE_FISICA para PhysicalActivityLevel.
// Levels: 0 = sedentary, 1 = insufficient, 2-3 = moderate, 4-5 = active
func ActivityLevelFromLevel(level int) models.PhysicalActivityLevel {
	switch {
	case level == 0:
		return models.ActivitySedentary
	case level == 1:
		return models.ActivityInsufficient
	case level <= 3:
		return models.ActivityModerate
	default:
		return models.ActivityActive
	}
}

// BoolFromRiskLevel converte selected_level de itens como DOENCA_CARDIOVASCULAR para bool.
// cutoff define o limite: level <= cutoff → true.
func BoolFromRiskLevel(level int, cutoff int) bool {
	return level <= cutoff
}
