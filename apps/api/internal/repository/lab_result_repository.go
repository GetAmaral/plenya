package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// LabResultRepository handles database operations for lab results
type LabResultRepository struct {
	db *gorm.DB
}

// NewLabResultRepository creates a new lab result repository instance
func NewLabResultRepository(db *gorm.DB) *LabResultRepository {
	return &LabResultRepository{db: db}
}

// GetHistoricalResultsByLabTestCode retrieves a map of [LabTestCode]→LabResult (most recent)
// Searches entire patient history using DISTINCT ON (PostgreSQL feature)
// Returns only results with non-null ResultNumeric and non-null LabTestDefinition.Code
func (r *LabResultRepository) GetHistoricalResultsByLabTestCode(patientID uuid.UUID) (map[string]models.LabResult, error) {
	var results []models.LabResult

	// Optimized query with DISTINCT ON (PostgreSQL)
	// Returns the most recent result for each LabTestCode
	//
	// Query explanation:
	// 1. JOIN lab_results with lab_result_batches to get patient_id and collection_date
	// 2. JOIN with lab_test_definitions to get the test code
	// 3. Filter by patient_id and ensure we have numeric results and test codes
	// 4. DISTINCT ON (ltd.code) ensures we get only one result per test code
	// 5. ORDER BY ltd.code, lrb.collection_date DESC, lr.created_at DESC ensures we get the most recent
	err := r.db.
		Raw(`
			SELECT DISTINCT ON (ltd.code)
				lr.id,
				lr.lab_result_batch_id,
				lr.lab_test_definition_id,
				lr.test_name,
				lr.test_type,
				lr.result_text,
				lr.result_numeric,
				lr.unit,
				lr.result_numeric_original,
				lr.unit_original,
				lr.interpretation,
				lr.level,
				lr.matched,
				lr.created_at,
				lr.updated_at,
				lr.deleted_at
			FROM lab_results lr
			INNER JOIN lab_result_batches lrb ON lr.lab_result_batch_id = lrb.id
			INNER JOIN lab_test_definitions ltd ON lr.lab_test_definition_id = ltd.id
			WHERE lrb.patient_id = ?
				AND lr.result_numeric IS NOT NULL
				AND ltd.code IS NOT NULL
				AND lr.deleted_at IS NULL
				AND lrb.deleted_at IS NULL
			ORDER BY ltd.code, lrb.collection_date DESC, lr.created_at DESC
		`, patientID).
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	// Preload LabTestDefinition for each result to get the code (excluding soft-deleted)
	for i := range results {
		if results[i].LabTestDefinitionID != nil {
			var labTestDef models.LabTestDefinition
			if err := r.db.Where("deleted_at IS NULL").First(&labTestDef, "id = ?", *results[i].LabTestDefinitionID).Error; err == nil {
				results[i].LabTestDefinition = &labTestDef
			}
		}
	}

	// Convert to map [LabTestCode]→LabResult
	resultMap := make(map[string]models.LabResult)
	for _, result := range results {
		if result.LabTestDefinition != nil && result.LabTestDefinition.Code != "" {
			resultMap[result.LabTestDefinition.Code] = result
		}
	}

	return resultMap, nil
}
