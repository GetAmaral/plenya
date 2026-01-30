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
	ErrPatientNotFound     = errors.New("patient not found")
	ErrPatientAlreadyExists = errors.New("patient already exists for this user")
	ErrUnauthorized        = errors.New("unauthorized")
)

type PatientService struct {
	db *gorm.DB
}

func NewPatientService(db *gorm.DB) *PatientService {
	return &PatientService{db: db}
}

// Create cria um novo paciente
func (s *PatientService) Create(userID uuid.UUID, req *dto.CreatePatientRequest) (*dto.PatientResponse, error) {
	// Verificar se já existe paciente para este usuário
	var existing models.Patient
	if err := s.db.Where("user_id = ?", userID).First(&existing).Error; err == nil {
		return nil, ErrPatientAlreadyExists
	}

	// Parse da data de nascimento
	birthDate, err := time.Parse("2006-01-02", req.BirthDate)
	if err != nil {
		return nil, errors.New("invalid birth date format, expected YYYY-MM-DD")
	}

	// Criar paciente
	patient := models.Patient{
		UserID:       userID,
		Name:         req.Name,
		CPF:          req.CPF, // Será criptografado pelo hook
		BirthDate:    birthDate,
		Gender:       req.Gender,
		Phone:        req.Phone,
		Address:      req.Address,
		Municipality: req.Municipality,
		State:        req.State,
		MotherName:   req.MotherName,
		FatherName:   req.FatherName,
		Height:       req.Height,
		Weight:       req.Weight,
	}

	if err := s.db.Create(&patient).Error; err != nil {
		return nil, err
	}

	return s.toDTO(&patient), nil
}

// GetByID busca um paciente por ID
func (s *PatientService) GetByID(patientID, userID uuid.UUID, userRole models.UserRole) (*dto.PatientResponse, error) {
	var patient models.Patient
	query := s.db.Where("id = ?", patientID)

	// Pacientes só podem ver seus próprios dados
	if userRole == models.RolePatient {
		query = query.Where("user_id = ?", userID)
	}

	if err := query.First(&patient).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPatientNotFound
		}
		return nil, err
	}

	return s.toDTO(&patient), nil
}

// List lista todos os pacientes (com paginação)
func (s *PatientService) List(userID uuid.UUID, userRole models.UserRole, limit, offset int) ([]dto.PatientResponse, error) {
	var patients []models.Patient
	query := s.db.Limit(limit).Offset(offset).Order("created_at DESC")

	// Pacientes só podem ver seus próprios dados
	if userRole == models.RolePatient {
		query = query.Where("user_id = ?", userID)
	}

	if err := query.Find(&patients).Error; err != nil {
		return nil, err
	}

	result := make([]dto.PatientResponse, len(patients))
	for i, p := range patients {
		result[i] = *s.toDTO(&p)
	}

	return result, nil
}

// Update atualiza um paciente
func (s *PatientService) Update(patientID, userID uuid.UUID, userRole models.UserRole, req *dto.UpdatePatientRequest) (*dto.PatientResponse, error) {
	var patient models.Patient
	query := s.db.Where("id = ?", patientID)

	// Pacientes só podem atualizar seus próprios dados
	if userRole == models.RolePatient {
		query = query.Where("user_id = ?", userID)
	}

	if err := query.First(&patient).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPatientNotFound
		}
		return nil, err
	}

	// Atualizar campos
	if req.Name != nil {
		patient.Name = *req.Name
	}
	if req.BirthDate != nil {
		birthDate, err := time.Parse("2006-01-02", *req.BirthDate)
		if err != nil {
			return nil, errors.New("invalid birth date format, expected YYYY-MM-DD")
		}
		patient.BirthDate = birthDate
	}
	if req.Gender != nil {
		patient.Gender = *req.Gender
	}
	if req.Phone != nil {
		patient.Phone = req.Phone
	}
	if req.Address != nil {
		patient.Address = req.Address
	}
	if req.Municipality != nil {
		patient.Municipality = req.Municipality
	}
	if req.State != nil {
		patient.State = req.State
	}
	if req.MotherName != nil {
		patient.MotherName = req.MotherName
	}
	if req.FatherName != nil {
		patient.FatherName = req.FatherName
	}
	if req.Height != nil {
		patient.Height = req.Height
	}
	if req.Weight != nil {
		patient.Weight = req.Weight
	}

	if err := s.db.Save(&patient).Error; err != nil {
		return nil, err
	}

	return s.toDTO(&patient), nil
}

// Delete faz soft delete de um paciente
func (s *PatientService) Delete(patientID, userID uuid.UUID, userRole models.UserRole) error {
	query := s.db.Where("id = ?", patientID)

	// Pacientes não podem deletar seus próprios dados
	if userRole == models.RolePatient {
		return ErrUnauthorized
	}

	result := query.Delete(&models.Patient{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrPatientNotFound
	}

	return nil
}

// toDTO converte Patient para PatientResponse
func (s *PatientService) toDTO(patient *models.Patient) *dto.PatientResponse {
	return &dto.PatientResponse{
		ID:           patient.ID.String(),
		UserID:       patient.UserID.String(),
		Name:         patient.Name,
		BirthDate:    patient.BirthDate.Format("2006-01-02"),
		Gender:       patient.Gender,
		Phone:        patient.Phone,
		Address:      patient.Address,
		Municipality: patient.Municipality,
		State:        patient.State,
		MotherName:   patient.MotherName,
		FatherName:   patient.FatherName,
		Height:       patient.Height,
		Weight:       patient.Weight,
		CreatedAt:    patient.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    patient.UpdatedAt.Format(time.RFC3339),
	}
}
