package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// PatientCheckInsService gerencia HealthCheckIn no escopo do paciente logado.
// Sem rate limit explícito — middleware de grupo /patient/me já cobre.
type PatientCheckInsService struct {
	db *gorm.DB
}

func NewPatientCheckInsService(db *gorm.DB) *PatientCheckInsService {
	return &PatientCheckInsService{db: db}
}

// CreateCheckInInput é o payload do paciente.
type CreateCheckInInput struct {
	Energy       int     `json:"energy" validate:"required,min=1,max=5"`
	Pain         int     `json:"pain" validate:"min=0,max=5"`
	PainLocation *string `json:"painLocation,omitempty"`
	Mood         int     `json:"mood" validate:"required,min=1,max=5"`
	SleepHours   float64 `json:"sleepHours" validate:"min=0,max=24"`
	SleepQuality int     `json:"sleepQuality" validate:"required,min=1,max=5"`
	Stress       int     `json:"stress" validate:"required,min=1,max=5"`
	Notes        *string `json:"notes,omitempty"`
}

// Create insere um novo check-in.
func (s *PatientCheckInsService) Create(patientID uuid.UUID, in CreateCheckInInput) (*models.HealthCheckIn, error) {
	row := &models.HealthCheckIn{
		PatientID:    patientID,
		Energy:       in.Energy,
		Pain:         in.Pain,
		PainLocation: in.PainLocation,
		Mood:         in.Mood,
		SleepHours:   in.SleepHours,
		SleepQuality: in.SleepQuality,
		Stress:       in.Stress,
		Notes:        in.Notes,
	}
	if err := s.db.Create(row).Error; err != nil {
		return nil, err
	}
	return row, nil
}

// List devolve check-ins recentes (default 30 últimos, max 90 dias).
// Profissional consome via PatientCheckInsService.ListForStaff (web).
func (s *PatientCheckInsService) List(patientID uuid.UUID, limit int) ([]models.HealthCheckIn, error) {
	if limit <= 0 || limit > 100 {
		limit = 30
	}
	var rows []models.HealthCheckIn
	err := s.db.
		Where("patient_id = ? AND created_at >= ?", patientID, time.Now().AddDate(0, 0, -90)).
		Order("created_at DESC").
		Limit(limit).
		Find(&rows).Error
	return rows, err
}

// LatestToday retorna o último check-in de hoje (ou nil) — usado pelo card
// de home no app: se já preencheu hoje, mostra estado "feito ✓" em vez do form.
func (s *PatientCheckInsService) LatestToday(patientID uuid.UUID) (*models.HealthCheckIn, error) {
	startOfDay := time.Now().Truncate(24 * time.Hour)
	var row models.HealthCheckIn
	err := s.db.
		Where("patient_id = ? AND created_at >= ?", patientID, startOfDay).
		Order("created_at DESC").
		First(&row).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &row, nil
}

// ListForStaff devolve check-ins do paciente pra leitura de profissional.
// Sem limite de tempo (continuum quer histórico completo).
func (s *PatientCheckInsService) ListForStaff(patientID uuid.UUID, limit int) ([]models.HealthCheckIn, error) {
	if limit <= 0 || limit > 365 {
		limit = 90
	}
	var rows []models.HealthCheckIn
	err := s.db.
		Where("patient_id = ?", patientID).
		Order("created_at DESC").
		Limit(limit).
		Find(&rows).Error
	return rows, err
}
