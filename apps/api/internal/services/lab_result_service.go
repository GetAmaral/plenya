package services

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

var (
	ErrLabResultNotFound = errors.New("lab result not found")
)

// LabResultService is DEPRECATED - Use LabResultBatchService instead
// This service is kept for backward compatibility but should not be used for new code
type LabResultService struct {
	db *gorm.DB
}

func NewLabResultService(db *gorm.DB) *LabResultService {
	return &LabResultService{db: db}
}

// DEPRECATED: Use LabResultBatchService.Create instead
// This method is maintained for backward compatibility only
func (s *LabResultService) Create(doctorID uuid.UUID, req *dto.CreateLabResultRequest) (*dto.LabResultResponse, error) {
	return nil, errors.New("DEPRECATED: Use LabResultBatchService.Create instead - single results must be created within a batch")
}

// List lista resultados (mantido para compatibilidade)
func (s *LabResultService) List(userID uuid.UUID, limit, offset int) ([]dto.LabResultResponse, error) {
	// CRITICAL SECURITY: Get user's selected patient
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}

	if user.SelectedPatientID == nil {
		return nil, errors.New("no patient selected - please select a patient first")
	}

	// Query results via batches
	var results []models.LabResult
	err := s.db.Joins("JOIN lab_result_batches ON lab_results.lab_result_batch_id = lab_result_batches.id").
		Where("lab_result_batches.patient_id = ?", *user.SelectedPatientID).
		Order("lab_results.created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&results).Error

	if err != nil {
		return nil, err
	}

	// Convert to DTOs
	dtos := make([]dto.LabResultResponse, len(results))
	for i, result := range results {
		// Need to fetch batch to get patient info
		var batch models.LabResultBatch
		if err := s.db.First(&batch, result.LabResultBatchID).Error; err != nil {
			continue
		}

		dtos[i] = dto.LabResultResponse{
			ID:             result.ID.String(),
			PatientID:      batch.PatientID.String(),
			DoctorID:       "",
			TestName:       result.TestName,
			TestType:       result.TestType,
			OrderedAt:      batch.CollectionDate.Format("2006-01-02T15:04:05Z07:00"),
			PerformedAt:    nil,
			Result:         result.ResultText,
			Unit:           result.Unit,
			Interpretation: result.Interpretation,
			CreatedAt:      result.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:      result.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		}

		if batch.RequestingDoctorID != nil {
			doctorID := batch.RequestingDoctorID.String()
			dtos[i].DoctorID = doctorID
		}

		if batch.ResultDate != nil {
			performedAt := batch.ResultDate.Format("2006-01-02T15:04:05Z07:00")
			dtos[i].PerformedAt = &performedAt
		}
	}

	return dtos, nil
}

// GetByID busca um resultado por ID
func (s *LabResultService) GetByID(resultID, userID uuid.UUID) (*dto.LabResultResponse, error) {
	// CRITICAL SECURITY: Get user's selected patient
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}

	if user.SelectedPatientID == nil {
		return nil, errors.New("no patient selected - please select a patient first")
	}

	var result models.LabResult
	err := s.db.Joins("JOIN lab_result_batches ON lab_results.lab_result_batch_id = lab_result_batches.id").
		Where("lab_results.id = ? AND lab_result_batches.patient_id = ?", resultID, *user.SelectedPatientID).
		First(&result).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrLabResultNotFound
		}
		return nil, err
	}

	// Fetch batch
	var batch models.LabResultBatch
	if err := s.db.First(&batch, result.LabResultBatchID).Error; err != nil {
		return nil, err
	}

	dto := &dto.LabResultResponse{
		ID:             result.ID.String(),
		PatientID:      batch.PatientID.String(),
		DoctorID:       "",
		TestName:       result.TestName,
		TestType:       result.TestType,
		OrderedAt:      batch.CollectionDate.Format("2006-01-02T15:04:05Z07:00"),
		PerformedAt:    nil,
		Result:         result.ResultText,
		Unit:           result.Unit,
		Interpretation: result.Interpretation,
		CreatedAt:      result.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:      result.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}

	if batch.RequestingDoctorID != nil {
		doctorID := batch.RequestingDoctorID.String()
		dto.DoctorID = doctorID
	}

	if batch.ResultDate != nil {
		performedAt := batch.ResultDate.Format("2006-01-02T15:04:05Z07:00")
		dto.PerformedAt = &performedAt
	}

	return dto, nil
}

// DEPRECATED: Use LabResultBatchService.Update instead
func (s *LabResultService) Update(resultID, userID uuid.UUID, req *dto.UpdateLabResultRequest) (*dto.LabResultResponse, error) {
	return nil, errors.New("DEPRECATED: Use LabResultBatchService.UpdateResult instead")
}

// DEPRECATED: Use LabResultBatchService.Delete instead
func (s *LabResultService) Delete(resultID, userID uuid.UUID, userRole string) error {
	return errors.New("DEPRECATED: Use LabResultBatchService.DeleteResult instead")
}
