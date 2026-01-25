package services

import (
	"github.com/google/uuid"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
)

// LabResultValueService handles business logic for lab result values
type LabResultValueService struct {
	repo *repository.LabResultValueRepository
}

// NewLabResultValueService creates a new lab result value service
func NewLabResultValueService(repo *repository.LabResultValueRepository) *LabResultValueService {
	return &LabResultValueService{repo: repo}
}

// CreateLabResultValue creates a new lab result value
func (s *LabResultValueService) CreateLabResultValue(value *models.LabResultValue) error {
	return s.repo.CreateLabResultValue(value)
}

// CreateLabResultValues creates multiple lab result values in a transaction
func (s *LabResultValueService) CreateLabResultValues(values []models.LabResultValue) error {
	return s.repo.CreateLabResultValues(values)
}

// GetLabResultValueByID retrieves a lab result value by ID
func (s *LabResultValueService) GetLabResultValueByID(id uuid.UUID) (*models.LabResultValue, error) {
	return s.repo.GetLabResultValueByID(id)
}

// GetValuesByLabResult retrieves all values for a specific lab result
func (s *LabResultValueService) GetValuesByLabResult(labResultID uuid.UUID) ([]models.LabResultValue, error) {
	return s.repo.GetValuesByLabResult(labResultID)
}

// GetValuesByPatient retrieves all lab result values for a patient
func (s *LabResultValueService) GetValuesByPatient(patientID uuid.UUID) ([]models.LabResultValue, error) {
	return s.repo.GetValuesByPatient(patientID)
}

// GetLatestValueForTest retrieves the most recent value for a specific test and patient
func (s *LabResultValueService) GetLatestValueForTest(patientID uuid.UUID, labTestDefID uuid.UUID) (*models.LabResultValue, error) {
	return s.repo.GetLatestValueForTest(patientID, labTestDefID)
}

// GetAbnormalValues retrieves all abnormal values for a patient
func (s *LabResultValueService) GetAbnormalValues(patientID uuid.UUID) ([]models.LabResultValue, error) {
	return s.repo.GetAbnormalValues(patientID)
}

// GetCriticalValues retrieves all critical values for a patient
func (s *LabResultValueService) GetCriticalValues(patientID uuid.UUID) ([]models.LabResultValue, error) {
	return s.repo.GetCriticalValues(patientID)
}

// UpdateLabResultValue updates an existing lab result value
func (s *LabResultValueService) UpdateLabResultValue(id uuid.UUID, value *models.LabResultValue) error {
	existing, err := s.repo.GetLabResultValueByID(id)
	if err != nil {
		return err
	}

	existing.LabResultID = value.LabResultID
	existing.LabTestDefinitionID = value.LabTestDefinitionID
	existing.NumericValue = value.NumericValue
	existing.TextValue = value.TextValue
	existing.BooleanValue = value.BooleanValue
	existing.Unit = value.Unit
	existing.ReferenceRange = value.ReferenceRange
	existing.IsAbnormal = value.IsAbnormal
	existing.IsCritical = value.IsCritical
	existing.Notes = value.Notes

	return s.repo.UpdateLabResultValue(existing)
}

// DeleteLabResultValue soft deletes a lab result value
func (s *LabResultValueService) DeleteLabResultValue(id uuid.UUID) error {
	return s.repo.DeleteLabResultValue(id)
}

// DeleteValuesByLabResult deletes all values for a specific lab result
func (s *LabResultValueService) DeleteValuesByLabResult(labResultID uuid.UUID) error {
	return s.repo.DeleteValuesByLabResult(labResultID)
}
