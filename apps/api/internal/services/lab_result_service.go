package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

var (
	ErrLabResultNotFound = errors.New("lab result not found")
)

type LabResultService struct {
	db *gorm.DB
}

func NewLabResultService(db *gorm.DB) *LabResultService {
	return &LabResultService{db: db}
}

// Create cria um novo resultado de exame
func (s *LabResultService) Create(doctorID uuid.UUID, req *dto.CreateLabResultRequest) (*dto.LabResultResponse, error) {
	// Parse patient ID
	patientID, err := uuid.Parse(req.PatientID)
	if err != nil {
		return nil, errors.New("invalid patient id")
	}

	// Verificar se paciente existe
	var patient models.Patient
	if err := s.db.First(&patient, patientID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPatientNotFound
		}
		return nil, err
	}

	// Parse ordered date
	orderedAt, err := time.Parse(time.RFC3339, req.OrderedAt)
	if err != nil {
		return nil, errors.New("invalid ordered date format, expected RFC3339")
	}

	// Criar lab result
	labResult := models.LabResult{
		PatientID:          patientID,
		RequestingDoctorID: doctorID,
		TestName:           req.TestName,
		TestType:           req.TestType,
		RequestDate:        orderedAt,
		Status:             req.Status,
		Result:             req.Result,
		Unit:               req.Unit,
		ReferenceRange:     req.ReferenceRange,
		Interpretation:     req.Interpretation,
	}

	// Parse performed date if provided (collection date)
	if req.PerformedAt != nil {
		performedAt, err := time.Parse(time.RFC3339, *req.PerformedAt)
		if err != nil {
			return nil, errors.New("invalid performed date format, expected RFC3339")
		}
		labResult.CollectionDate = &performedAt
	}

	if err := s.db.Create(&labResult).Error; err != nil {
		return nil, err
	}

	return s.toDTO(&labResult), nil
}

// GetByID busca um resultado de exame por ID
func (s *LabResultService) GetByID(labResultID, userID uuid.UUID, userRole models.UserRole) (*dto.LabResultResponse, error) {
	var labResult models.LabResult
	query := s.db.Where("id = ?", labResultID)

	// Pacientes só podem ver seus próprios resultados
	if userRole == models.RolePatient {
		var patient models.Patient
		if err := s.db.Where("user_id = ?", userID).First(&patient).Error; err != nil {
			return nil, ErrUnauthorized
		}
		query = query.Where("patient_id = ?", patient.ID)
	}

	if err := query.First(&labResult).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrLabResultNotFound
		}
		return nil, err
	}

	return s.toDTO(&labResult), nil
}

// List lista resultados de exames com filtros
func (s *LabResultService) List(userID uuid.UUID, userRole models.UserRole, patientID *uuid.UUID, limit, offset int) ([]dto.LabResultResponse, error) {
	var labResults []models.LabResult
	query := s.db.Limit(limit).Offset(offset).Order("ordered_at DESC")

	// Filtros baseados em role
	if userRole == models.RolePatient {
		var patient models.Patient
		if err := s.db.Where("user_id = ?", userID).First(&patient).Error; err != nil {
			return nil, ErrUnauthorized
		}
		query = query.Where("patient_id = ?", patient.ID)
	} else if userRole == models.RoleDoctor {
		// Doctors veem seus próprios exames solicitados
		query = query.Where("doctor_id = ?", userID)
	} else {
		// Admin/Nurse podem filtrar por paciente
		if patientID != nil {
			query = query.Where("patient_id = ?", *patientID)
		}
	}

	if err := query.Find(&labResults).Error; err != nil {
		return nil, err
	}

	result := make([]dto.LabResultResponse, len(labResults))
	for i, lr := range labResults {
		result[i] = *s.toDTO(&lr)
	}

	return result, nil
}

// Update atualiza um resultado de exame
func (s *LabResultService) Update(labResultID, userID uuid.UUID, userRole models.UserRole, req *dto.UpdateLabResultRequest) (*dto.LabResultResponse, error) {
	var labResult models.LabResult
	query := s.db.Where("id = ?", labResultID)

	if err := query.First(&labResult).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrLabResultNotFound
		}
		return nil, err
	}

	// Permissões: apenas o médico que solicitou ou admin/nurse
	if userRole == models.RoleDoctor && labResult.RequestingDoctorID != userID {
		return nil, ErrUnauthorized
	} else if userRole == models.RolePatient {
		return nil, ErrUnauthorized
	}

	// Atualizar campos
	if req.PerformedAt != nil {
		performedAt, err := time.Parse(time.RFC3339, *req.PerformedAt)
		if err != nil {
			return nil, errors.New("invalid performed date format, expected RFC3339")
		}
		labResult.CollectionDate = &performedAt
		// Auto-set result date if status is completed
		if req.Status != nil && *req.Status == models.LabResultCompleted {
			now := time.Now()
			labResult.ResultDate = &now
		}
	}
	if req.Status != nil {
		labResult.Status = *req.Status
	}
	if req.Result != nil {
		labResult.Result = req.Result
	}
	if req.Unit != nil {
		labResult.Unit = req.Unit
	}
	if req.ReferenceRange != nil {
		labResult.ReferenceRange = req.ReferenceRange
	}
	if req.Interpretation != nil {
		labResult.Interpretation = req.Interpretation
	}

	if err := s.db.Save(&labResult).Error; err != nil {
		return nil, err
	}

	return s.toDTO(&labResult), nil
}

// Delete faz soft delete de um resultado de exame
func (s *LabResultService) Delete(labResultID, userID uuid.UUID, userRole models.UserRole) error {
	// Apenas admin pode deletar resultados de exames
	if userRole != models.RoleAdmin {
		return ErrUnauthorized
	}

	result := s.db.Where("id = ?", labResultID).Delete(&models.LabResult{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrLabResultNotFound
	}

	return nil
}

// toDTO converte LabResult para LabResultResponse
func (s *LabResultService) toDTO(labResult *models.LabResult) *dto.LabResultResponse {
	resp := &dto.LabResultResponse{
		ID:             labResult.ID.String(),
		PatientID:      labResult.PatientID.String(),
		DoctorID:       labResult.RequestingDoctorID.String(),
		TestName:       labResult.TestName,
		TestType:       labResult.TestType,
		OrderedAt:      labResult.RequestDate.Format(time.RFC3339),
		Status:         labResult.Status,
		Result:         labResult.Result,
		Unit:           labResult.Unit,
		ReferenceRange: labResult.ReferenceRange,
		Interpretation: labResult.Interpretation,
		CreatedAt:      labResult.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      labResult.UpdatedAt.Format(time.RFC3339),
	}

	// Use CollectionDate as PerformedAt
	if labResult.CollectionDate != nil {
		pa := labResult.CollectionDate.Format(time.RFC3339)
		resp.PerformedAt = &pa
	}

	return resp
}
