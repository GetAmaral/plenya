package repository

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// LabResultValueRepository handles database operations for lab result values
type LabResultValueRepository struct {
	db *gorm.DB
}

// NewLabResultValueRepository creates a new lab result value repository instance
func NewLabResultValueRepository(db *gorm.DB) *LabResultValueRepository {
	return &LabResultValueRepository{db: db}
}

// CreateLabResultValue creates a new lab result value
func (r *LabResultValueRepository) CreateLabResultValue(value *models.LabResultValue) error {
	return r.db.Create(value).Error
}

// CreateLabResultValues creates multiple lab result values in a transaction
func (r *LabResultValueRepository) CreateLabResultValues(values []models.LabResultValue) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, value := range values {
			if err := tx.Create(&value).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// GetLabResultValueByID retrieves a lab result value by ID
func (r *LabResultValueRepository) GetLabResultValueByID(id uuid.UUID) (*models.LabResultValue, error) {
	var value models.LabResultValue
	err := r.db.
		Preload("LabResult").
		Preload("LabTestDefinition").
		First(&value, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("lab result value not found")
		}
		return nil, err
	}
	return &value, nil
}

// GetValuesByLabResult retrieves all values for a specific lab result
func (r *LabResultValueRepository) GetValuesByLabResult(labResultID uuid.UUID) ([]models.LabResultValue, error) {
	var values []models.LabResultValue
	err := r.db.
		Where("lab_result_id = ?", labResultID).
		Preload("LabTestDefinition").
		Order("lab_test_definition_id ASC").
		Find(&values).Error
	return values, err
}

// GetValuesByPatient retrieves all lab result values for a patient
func (r *LabResultValueRepository) GetValuesByPatient(patientID uuid.UUID) ([]models.LabResultValue, error) {
	var values []models.LabResultValue
	err := r.db.
		Joins("JOIN lab_results ON lab_results.id = lab_result_values.lab_result_id").
		Where("lab_results.patient_id = ?", patientID).
		Preload("LabResult").
		Preload("LabTestDefinition").
		Order("lab_result_values.created_at DESC").
		Find(&values).Error
	return values, err
}

// GetLatestValueForTest retrieves the most recent value for a specific test and patient
func (r *LabResultValueRepository) GetLatestValueForTest(patientID uuid.UUID, labTestDefID uuid.UUID) (*models.LabResultValue, error) {
	var value models.LabResultValue
	err := r.db.
		Joins("JOIN lab_results ON lab_results.id = lab_result_values.lab_result_id").
		Where("lab_results.patient_id = ?", patientID).
		Where("lab_result_values.lab_test_definition_id = ?", labTestDefID).
		Preload("LabResult").
		Preload("LabTestDefinition").
		Order("lab_result_values.created_at DESC").
		First(&value).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no lab result value found for this test and patient")
		}
		return nil, err
	}
	return &value, nil
}

// UpdateLabResultValue updates an existing lab result value
func (r *LabResultValueRepository) UpdateLabResultValue(value *models.LabResultValue) error {
	return r.db.Save(value).Error
}

// DeleteLabResultValue soft deletes a lab result value
func (r *LabResultValueRepository) DeleteLabResultValue(id uuid.UUID) error {
	result := r.db.Delete(&models.LabResultValue{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("lab result value not found")
	}
	return nil
}

// DeleteValuesByLabResult deletes all values for a specific lab result
func (r *LabResultValueRepository) DeleteValuesByLabResult(labResultID uuid.UUID) error {
	return r.db.Where("lab_result_id = ?", labResultID).Delete(&models.LabResultValue{}).Error
}
