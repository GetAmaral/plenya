package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

var ErrRecallNotFound = errors.New("recall not found")

const recallDateLayout = "2006-01-02"

type RecallService struct {
	db *gorm.DB
}

func NewRecallService(db *gorm.DB) *RecallService {
	return &RecallService{db: db}
}

// Create registra um retorno previsto.
func (s *RecallService) Create(createdBy uuid.UUID, req *dto.CreateRecallRequest) (*dto.RecallResponse, error) {
	patientID, err := uuid.Parse(req.PatientID)
	if err != nil {
		return nil, errors.New("invalid patient id")
	}
	due, err := time.Parse(recallDateLayout, req.DueDate)
	if err != nil {
		return nil, errors.New("invalid dueDate format, expected YYYY-MM-DD")
	}

	recall := models.AppointmentRecall{
		PatientID:       patientID,
		DueDate:         due,
		Reason:          req.Reason,
		Notes:           req.Notes,
		Status:          models.RecallPending,
		CreatedByUserID: createdBy,
	}
	if req.DoctorID != nil && *req.DoctorID != "" {
		did, err := uuid.Parse(*req.DoctorID)
		if err != nil {
			return nil, errors.New("invalid doctor id")
		}
		recall.DoctorID = &did
	}
	if req.SourceAppointmentID != nil && *req.SourceAppointmentID != "" {
		said, err := uuid.Parse(*req.SourceAppointmentID)
		if err != nil {
			return nil, errors.New("invalid source appointment id")
		}
		recall.SourceAppointmentID = &said
	}

	if err := s.db.Create(&recall).Error; err != nil {
		return nil, err
	}
	return s.reload(recall.ID)
}

// List lista retornos. Default: pending. statusFilter "all" => todos.
// dueBefore (opcional) filtra retornos com due_date <= dueBefore (retornos a vencer).
func (s *RecallService) List(statusFilter string, dueBefore *time.Time) ([]dto.RecallResponse, error) {
	var recalls []models.AppointmentRecall
	query := s.db.Preload("Patient").Preload("Doctor").Order("due_date ASC")

	switch statusFilter {
	case "", "pending":
		query = query.Where("status = ?", models.RecallPending)
	case "all":
		// sem filtro de status
	default:
		query = query.Where("status = ?", statusFilter)
	}
	if dueBefore != nil {
		query = query.Where("due_date <= ?", *dueBefore)
	}

	if err := query.Find(&recalls).Error; err != nil {
		return nil, err
	}
	result := make([]dto.RecallResponse, len(recalls))
	for i := range recalls {
		result[i] = *s.toDTO(&recalls[i])
	}
	return result, nil
}

// Update altera status/data/observações (ex.: agendar = scheduled, dispensar).
func (s *RecallService) Update(id uuid.UUID, req *dto.UpdateRecallRequest) (*dto.RecallResponse, error) {
	var recall models.AppointmentRecall
	if err := s.db.Where("id = ?", id).First(&recall).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecallNotFound
		}
		return nil, err
	}

	if req.Status != nil {
		recall.Status = *req.Status
	}
	if req.DueDate != nil && *req.DueDate != "" {
		due, err := time.Parse(recallDateLayout, *req.DueDate)
		if err != nil {
			return nil, errors.New("invalid dueDate format, expected YYYY-MM-DD")
		}
		recall.DueDate = due
	}
	if req.Reason != nil {
		recall.Reason = req.Reason
	}
	if req.Notes != nil {
		recall.Notes = req.Notes
	}
	if req.ScheduledAppointmentID != nil && *req.ScheduledAppointmentID != "" {
		aid, err := uuid.Parse(*req.ScheduledAppointmentID)
		if err != nil {
			return nil, errors.New("invalid appointment id")
		}
		recall.ScheduledAppointmentID = &aid
	}

	if err := s.db.Save(&recall).Error; err != nil {
		return nil, err
	}
	return s.reload(recall.ID)
}

// Delete remove (soft delete) um retorno.
func (s *RecallService) Delete(id uuid.UUID) error {
	result := s.db.Where("id = ?", id).Delete(&models.AppointmentRecall{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrRecallNotFound
	}
	return nil
}

func (s *RecallService) reload(id uuid.UUID) (*dto.RecallResponse, error) {
	var recall models.AppointmentRecall
	if err := s.db.Preload("Patient").Preload("Doctor").Where("id = ?", id).First(&recall).Error; err != nil {
		return nil, err
	}
	return s.toDTO(&recall), nil
}

func (s *RecallService) toDTO(e *models.AppointmentRecall) *dto.RecallResponse {
	resp := &dto.RecallResponse{
		ID:              e.ID.String(),
		PatientID:       e.PatientID.String(),
		DueDate:         e.DueDate.Format(recallDateLayout),
		Reason:          e.Reason,
		Notes:           e.Notes,
		Status:          e.Status,
		CreatedByUserID: e.CreatedByUserID.String(),
		CreatedAt:       e.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       e.UpdatedAt.Format(time.RFC3339),
	}
	if e.DoctorID != nil {
		did := e.DoctorID.String()
		resp.DoctorID = &did
	}
	if e.SourceAppointmentID != nil {
		said := e.SourceAppointmentID.String()
		resp.SourceAppointmentID = &said
	}
	if e.ScheduledAppointmentID != nil {
		aid := e.ScheduledAppointmentID.String()
		resp.ScheduledAppointmentID = &aid
	}
	if e.Patient.ID != uuid.Nil {
		resp.Patient = &dto.AppointmentParty{
			ID:    e.Patient.ID.String(),
			Name:  e.Patient.Name,
			Email: e.Patient.Email,
			Phone: e.Patient.Phone,
		}
	}
	if e.Doctor != nil && e.Doctor.ID != uuid.Nil {
		email := e.Doctor.Email
		resp.Doctor = &dto.AppointmentParty{
			ID:    e.Doctor.ID.String(),
			Name:  e.Doctor.Name,
			Email: &email,
		}
	}
	return resp
}
