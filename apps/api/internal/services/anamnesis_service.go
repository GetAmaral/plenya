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
	ErrAnamnesisNotFound = errors.New("anamnesis not found")
)

type AnamnesisService struct {
	db *gorm.DB
}

func NewAnamnesisService(db *gorm.DB) *AnamnesisService {
	return &AnamnesisService{db: db}
}

// Create cria uma nova anamnese
func (s *AnamnesisService) Create(doctorID uuid.UUID, req *dto.CreateAnamnesisRequest) (*dto.AnamnesisResponse, error) {
	// CRITICAL SECURITY: Get user's selected patient
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, doctorID).Error; err != nil {
		return nil, err
	}

	// If no selected patient, cannot create anamnesis
	if user.SelectedPatientID == nil {
		return nil, errors.New("no patient selected - please select a patient first")
	}

	// Parse patient ID from request
	var patientID uuid.UUID
	if req.PatientID != "" {
		pid, err := uuid.Parse(req.PatientID)
		if err != nil {
			return nil, errors.New("invalid patient id")
		}
		// SECURITY: Validate that patientID matches selectedPatient
		if pid != *user.SelectedPatientID {
			return nil, errors.New("patient id does not match selected patient")
		}
		patientID = pid
	} else {
		// Auto-fill with selectedPatient
		patientID = *user.SelectedPatientID
	}

	// Verificar se o paciente existe
	var patient models.Patient
	if err := s.db.First(&patient, patientID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPatientNotFound
		}
		return nil, err
	}

	// Parse consultation date
	consultationDate, err := time.Parse(time.RFC3339, req.ConsultationDate)
	if err != nil {
		return nil, errors.New("invalid consultation date format, expected RFC3339")
	}

	// Criar anamnese
	anamnesis := models.Anamnesis{
		PatientID:               patientID,
		DoctorID:                doctorID,
		ChiefComplaint:          req.ChiefComplaint,
		HistoryOfPresentIllness: req.HistoryOfPresentIllness,
		PastMedicalHistory:      req.PastMedicalHistory,
		CurrentMedications:      req.CurrentMedications,
		Allergies:               req.Allergies,
		FamilyHistory:           req.FamilyHistory,
		SocialHistory:           req.SocialHistory,
		ReviewOfSystems:         req.ReviewOfSystems,
		PhysicalExamination:     req.PhysicalExamination,
		Assessment:              req.Assessment,
		Plan:                    req.Plan,
		ConsultationDate:        consultationDate,
		Notes:                   req.Notes,
	}

	if err := s.db.Create(&anamnesis).Error; err != nil {
		return nil, err
	}

	return s.toDTO(&anamnesis), nil
}

// GetByID busca uma anamnese por ID
func (s *AnamnesisService) GetByID(anamnesisID, userID uuid.UUID, userRole models.UserRole) (*dto.AnamnesisResponse, error) {
	// CRITICAL SECURITY: Get user's selected patient
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}

	// If no selected patient, cannot access any anamnesis
	if user.SelectedPatientID == nil {
		return nil, ErrAnamnesisNotFound
	}

	// ALWAYS filter by selectedPatient (all roles including admin)
	var anamnesis models.Anamnesis
	query := s.db.Where("id = ?", anamnesisID).Where("patient_id = ?", *user.SelectedPatientID)

	if err := query.First(&anamnesis).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAnamnesisNotFound
		}
		return nil, err
	}

	return s.toDTO(&anamnesis), nil
}

// List lista anamneses com filtros
func (s *AnamnesisService) List(userID uuid.UUID, userRole models.UserRole, patientID *uuid.UUID, limit, offset int) ([]dto.AnamnesisResponse, error) {
	// CRITICAL SECURITY: Get user's selected patient
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}

	// If no selected patient, return empty list (security measure)
	if user.SelectedPatientID == nil {
		return []dto.AnamnesisResponse{}, nil
	}

	// ALWAYS filter by selectedPatient (all roles including admin)
	var anamneses []models.Anamnesis
	query := s.db.Limit(limit).Offset(offset).Order("consultation_date DESC")
	query = query.Where("patient_id = ?", *user.SelectedPatientID)

	if err := query.Find(&anamneses).Error; err != nil {
		return nil, err
	}

	result := make([]dto.AnamnesisResponse, len(anamneses))
	for i, a := range anamneses {
		result[i] = *s.toDTO(&a)
	}

	return result, nil
}

// Update atualiza uma anamnese
func (s *AnamnesisService) Update(anamnesisID, doctorID uuid.UUID, userRole models.UserRole, req *dto.UpdateAnamnesisRequest) (*dto.AnamnesisResponse, error) {
	var anamnesis models.Anamnesis
	query := s.db.Where("id = ?", anamnesisID)

	// Apenas o médico que criou pode editar (ou admin)
	if userRole != models.RoleAdmin {
		query = query.Where("doctor_id = ?", doctorID)
	}

	if err := query.First(&anamnesis).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAnamnesisNotFound
		}
		return nil, err
	}

	// Atualizar campos
	if req.ChiefComplaint != nil {
		anamnesis.ChiefComplaint = *req.ChiefComplaint
	}
	if req.HistoryOfPresentIllness != nil {
		anamnesis.HistoryOfPresentIllness = *req.HistoryOfPresentIllness
	}
	if req.PastMedicalHistory != nil {
		anamnesis.PastMedicalHistory = req.PastMedicalHistory
	}
	if req.CurrentMedications != nil {
		anamnesis.CurrentMedications = req.CurrentMedications
	}
	if req.Allergies != nil {
		anamnesis.Allergies = req.Allergies
	}
	if req.FamilyHistory != nil {
		anamnesis.FamilyHistory = req.FamilyHistory
	}
	if req.SocialHistory != nil {
		anamnesis.SocialHistory = req.SocialHistory
	}
	if req.ReviewOfSystems != nil {
		anamnesis.ReviewOfSystems = req.ReviewOfSystems
	}
	if req.PhysicalExamination != nil {
		anamnesis.PhysicalExamination = req.PhysicalExamination
	}
	if req.Assessment != nil {
		anamnesis.Assessment = req.Assessment
	}
	if req.Plan != nil {
		anamnesis.Plan = req.Plan
	}
	if req.ConsultationDate != nil {
		consultationDate, err := time.Parse(time.RFC3339, *req.ConsultationDate)
		if err != nil {
			return nil, errors.New("invalid consultation date format, expected RFC3339")
		}
		anamnesis.ConsultationDate = consultationDate
	}
	if req.Notes != nil {
		anamnesis.Notes = req.Notes
	}

	if err := s.db.Save(&anamnesis).Error; err != nil {
		return nil, err
	}

	return s.toDTO(&anamnesis), nil
}

// Delete faz soft delete de uma anamnese
func (s *AnamnesisService) Delete(anamnesisID, doctorID uuid.UUID, userRole models.UserRole) error {
	query := s.db.Where("id = ?", anamnesisID)

	// Apenas o médico que criou pode deletar (ou admin)
	if userRole != models.RoleAdmin {
		query = query.Where("doctor_id = ?", doctorID)
	}

	result := query.Delete(&models.Anamnesis{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrAnamnesisNotFound
	}

	return nil
}

// toDTO converte Anamnesis para AnamnesisResponse
func (s *AnamnesisService) toDTO(anamnesis *models.Anamnesis) *dto.AnamnesisResponse {
	return &dto.AnamnesisResponse{
		ID:                      anamnesis.ID.String(),
		PatientID:               anamnesis.PatientID.String(),
		DoctorID:                anamnesis.DoctorID.String(),
		ChiefComplaint:          anamnesis.ChiefComplaint,
		HistoryOfPresentIllness: anamnesis.HistoryOfPresentIllness,
		PastMedicalHistory:      anamnesis.PastMedicalHistory,
		CurrentMedications:      anamnesis.CurrentMedications,
		Allergies:               anamnesis.Allergies,
		FamilyHistory:           anamnesis.FamilyHistory,
		SocialHistory:           anamnesis.SocialHistory,
		ReviewOfSystems:         anamnesis.ReviewOfSystems,
		PhysicalExamination:     anamnesis.PhysicalExamination,
		Assessment:              anamnesis.Assessment,
		Plan:                    anamnesis.Plan,
		ConsultationDate:        anamnesis.ConsultationDate.Format(time.RFC3339),
		Notes:                   anamnesis.Notes,
		CreatedAt:               anamnesis.CreatedAt.Format(time.RFC3339),
		UpdatedAt:               anamnesis.UpdatedAt.Format(time.RFC3339),
	}
}
