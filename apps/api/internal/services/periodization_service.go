package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

var ErrPeriodizationNotFound = errors.New("periodization not found")

type PeriodizationService struct {
	db *gorm.DB
}

func NewPeriodizationService(db *gorm.DB) *PeriodizationService {
	return &PeriodizationService{db: db}
}

// Create cria uma periodização com mesociclos
func (s *PeriodizationService) Create(userID uuid.UUID, req *dto.CreatePeriodizationRequest) (*dto.PeriodizationResponse, error) {
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}
	if user.SelectedPatientID == nil {
		return nil, errors.New("no patient selected")
	}

	patientID := *user.SelectedPatientID
	if req.PatientID != "" {
		pid, err := uuid.Parse(req.PatientID)
		if err != nil {
			return nil, errors.New("invalid patient id")
		}
		if pid != patientID {
			return nil, errors.New("patient id does not match selected patient")
		}
	}

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return nil, errors.New("invalid start date, expected YYYY-MM-DD")
	}

	periodization := models.WorkoutPeriodization{
		PatientID:   patientID,
		CreatedByID: userID,
		Framework:   req.Framework,
		StartDate:   startDate,
		TotalWeeks:  req.TotalWeeks,
		Objective:   req.Objective,
	}

	if err := s.db.Create(&periodization).Error; err != nil {
		return nil, err
	}

	return s.toDTO(&periodization), nil
}

// GetByID busca uma periodização por ID
func (s *PeriodizationService) GetByID(id, userID uuid.UUID) (*dto.PeriodizationResponse, error) {
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}
	if user.SelectedPatientID == nil {
		return nil, ErrPeriodizationNotFound
	}

	var periodization models.WorkoutPeriodization
	if err := s.db.Where("id = ? AND patient_id = ?", id, *user.SelectedPatientID).
		Preload("Mesocycles", func(db *gorm.DB) *gorm.DB { return db.Order("\"order\" ASC") }).
		First(&periodization).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPeriodizationNotFound
		}
		return nil, err
	}

	return s.toDTO(&periodization), nil
}

// List lista periodizações do paciente selecionado
func (s *PeriodizationService) List(userID uuid.UUID, limit, offset int) ([]dto.PeriodizationResponse, error) {
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}
	if user.SelectedPatientID == nil {
		return []dto.PeriodizationResponse{}, nil
	}

	var periodizations []models.WorkoutPeriodization
	if err := s.db.Where("patient_id = ?", *user.SelectedPatientID).
		Preload("Mesocycles", func(db *gorm.DB) *gorm.DB { return db.Order("\"order\" ASC") }).
		Order("start_date DESC").
		Limit(limit).Offset(offset).
		Find(&periodizations).Error; err != nil {
		return nil, err
	}

	result := make([]dto.PeriodizationResponse, len(periodizations))
	for i, p := range periodizations {
		result[i] = *s.toDTO(&p)
	}
	return result, nil
}

// Delete faz soft delete de uma periodização
func (s *PeriodizationService) Delete(id uuid.UUID, userRole models.Role) error {
	if userRole != models.RoleAdmin {
		return ErrUnauthorized
	}
	result := s.db.Where("id = ?", id).Delete(&models.WorkoutPeriodization{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrPeriodizationNotFound
	}
	return nil
}

func (s *PeriodizationService) toDTO(p *models.WorkoutPeriodization) *dto.PeriodizationResponse {
	resp := &dto.PeriodizationResponse{
		ID:                      p.ID.String(),
		PatientID:               p.PatientID.String(),
		CreatedByID:             p.CreatedByID.String(),
		Framework:               p.Framework,
		StartDate:               p.StartDate.Format("2006-01-02"),
		TotalWeeks:              p.TotalWeeks,
		Objective:               p.Objective,
		ScientificJustification: p.ScientificJustification,
		DisplayTitle:            p.DisplayTitle,
		CreatedAt:               p.CreatedAt.Format(time.RFC3339),
		UpdatedAt:               p.UpdatedAt.Format(time.RFC3339),
	}

	for _, m := range p.Mesocycles {
		resp.Mesocycles = append(resp.Mesocycles, dto.MesocycleResponse{
			ID:                 m.ID.String(),
			PeriodizationID:    m.PeriodizationID.String(),
			Order:              m.Order,
			Phase:              m.Phase,
			DurationWeeks:      m.DurationWeeks,
			VolumePercent:      m.VolumePercent,
			IntensityPercent:   m.IntensityPercent,
			PhysiologicalFocus: m.PhysiologicalFocus,
			StartDate:          m.StartDate.Format("2006-01-02"),
			EndDate:            m.EndDate.Format("2006-01-02"),
		})
	}

	return resp
}
