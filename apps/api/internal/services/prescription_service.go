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
	ErrPrescriptionNotFound = errors.New("prescription not found")
)

type PrescriptionService struct {
	db *gorm.DB
}

func NewPrescriptionService(db *gorm.DB) *PrescriptionService {
	return &PrescriptionService{db: db}
}

// Create cria uma nova prescrição
func (s *PrescriptionService) Create(doctorID uuid.UUID, req *dto.CreatePrescriptionRequest) (*dto.PrescriptionResponse, error) {
	// Parse patient ID
	patientID, err := uuid.Parse(req.PatientID)
	if err != nil {
		return nil, errors.New("invalid patient id")
	}

	// Verificar se o paciente existe
	var patient models.Patient
	if err := s.db.First(&patient, patientID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPatientNotFound
		}
		return nil, err
	}

	// Parse prescription date
	prescriptionDate, err := time.Parse(time.RFC3339, req.PrescriptionDate)
	if err != nil {
		return nil, errors.New("invalid prescription date format, expected RFC3339")
	}

	// Parse dates opcionais
	var startDate, endDate *time.Time
	if req.StartDate != nil {
		sd, err := time.Parse("2006-01-02", *req.StartDate)
		if err != nil {
			return nil, errors.New("invalid start date format, expected YYYY-MM-DD")
		}
		startDate = &sd
	}
	if req.EndDate != nil {
		ed, err := time.Parse("2006-01-02", *req.EndDate)
		if err != nil {
			return nil, errors.New("invalid end date format, expected YYYY-MM-DD")
		}
		endDate = &ed
	}

	// Criar prescrição
	prescription := models.Prescription{
		PatientID:        patientID,
		DoctorID:         doctorID,
		MedicationName:   req.MedicationName,
		ActiveIngredient: req.ActiveIngredient,
		Dosage:           req.Dosage,
		Frequency:        req.Frequency,
		Route:            req.Route,
		Duration:         req.Duration,
		Quantity:         req.Quantity,
		Instructions:     req.Instructions,
		Status:           models.PrescriptionActive,
		PrescriptionDate: prescriptionDate,
		StartDate:        startDate,
		EndDate:          endDate,
		Notes:            req.Notes,
	}

	if err := s.db.Create(&prescription).Error; err != nil {
		return nil, err
	}

	return s.toDTO(&prescription), nil
}

// GetByID busca uma prescrição por ID
func (s *PrescriptionService) GetByID(prescriptionID, userID uuid.UUID, userRole models.UserRole) (*dto.PrescriptionResponse, error) {
	var prescription models.Prescription
	query := s.db.Where("id = ?", prescriptionID)

	// Pacientes só podem ver suas próprias prescrições
	if userRole == models.RolePatient {
		var patient models.Patient
		if err := s.db.Where("user_id = ?", userID).First(&patient).Error; err != nil {
			return nil, ErrUnauthorized
		}
		query = query.Where("patient_id = ?", patient.ID)
	}

	if err := query.First(&prescription).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPrescriptionNotFound
		}
		return nil, err
	}

	return s.toDTO(&prescription), nil
}

// List lista prescrições com filtros
func (s *PrescriptionService) List(userID uuid.UUID, userRole models.UserRole, patientID *uuid.UUID, status *models.PrescriptionStatus, limit, offset int) ([]dto.PrescriptionResponse, error) {
	var prescriptions []models.Prescription
	query := s.db.Limit(limit).Offset(offset).Order("prescription_date DESC")

	// Pacientes só podem ver suas próprias prescrições
	if userRole == models.RolePatient {
		var patient models.Patient
		if err := s.db.Where("user_id = ?", userID).First(&patient).Error; err != nil {
			return nil, ErrUnauthorized
		}
		query = query.Where("patient_id = ?", patient.ID)
	} else if patientID != nil {
		query = query.Where("patient_id = ?", *patientID)
	}

	// Filtro por status
	if status != nil {
		query = query.Where("status = ?", *status)
	}

	if err := query.Find(&prescriptions).Error; err != nil {
		return nil, err
	}

	result := make([]dto.PrescriptionResponse, len(prescriptions))
	for i, p := range prescriptions {
		result[i] = *s.toDTO(&p)
	}

	return result, nil
}

// Update atualiza uma prescrição
func (s *PrescriptionService) Update(prescriptionID, doctorID uuid.UUID, userRole models.UserRole, req *dto.UpdatePrescriptionRequest) (*dto.PrescriptionResponse, error) {
	var prescription models.Prescription
	query := s.db.Where("id = ?", prescriptionID)

	// Apenas o médico que criou pode editar (ou admin)
	if userRole != models.RoleAdmin {
		query = query.Where("doctor_id = ?", doctorID)
	}

	if err := query.First(&prescription).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPrescriptionNotFound
		}
		return nil, err
	}

	// Atualizar campos
	if req.Status != nil {
		prescription.Status = *req.Status
	}
	if req.Instructions != nil {
		prescription.Instructions = req.Instructions
	}
	if req.Notes != nil {
		prescription.Notes = req.Notes
	}

	if err := s.db.Save(&prescription).Error; err != nil {
		return nil, err
	}

	return s.toDTO(&prescription), nil
}

// Delete faz soft delete de uma prescrição
func (s *PrescriptionService) Delete(prescriptionID, doctorID uuid.UUID, userRole models.UserRole) error {
	query := s.db.Where("id = ?", prescriptionID)

	// Apenas o médico que criou pode deletar (ou admin)
	if userRole != models.RoleAdmin {
		query = query.Where("doctor_id = ?", doctorID)
	}

	result := query.Delete(&models.Prescription{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrPrescriptionNotFound
	}

	return nil
}

// toDTO converte Prescription para PrescriptionResponse
func (s *PrescriptionService) toDTO(prescription *models.Prescription) *dto.PrescriptionResponse {
	resp := &dto.PrescriptionResponse{
		ID:               prescription.ID.String(),
		PatientID:        prescription.PatientID.String(),
		DoctorID:         prescription.DoctorID.String(),
		MedicationName:   prescription.MedicationName,
		ActiveIngredient: prescription.ActiveIngredient,
		Dosage:           prescription.Dosage,
		Frequency:        prescription.Frequency,
		Route:            prescription.Route,
		Duration:         prescription.Duration,
		Quantity:         prescription.Quantity,
		Instructions:     prescription.Instructions,
		Status:           prescription.Status,
		PrescriptionDate: prescription.PrescriptionDate.Format(time.RFC3339),
		Notes:            prescription.Notes,
		CreatedAt:        prescription.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        prescription.UpdatedAt.Format(time.RFC3339),
	}

	if prescription.StartDate != nil {
		sd := prescription.StartDate.Format("2006-01-02")
		resp.StartDate = &sd
	}
	if prescription.EndDate != nil {
		ed := prescription.EndDate.Format("2006-01-02")
		resp.EndDate = &ed
	}

	return resp
}
