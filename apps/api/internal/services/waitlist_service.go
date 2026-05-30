package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

var ErrWaitlistEntryNotFound = errors.New("waitlist entry not found")

type WaitlistService struct {
	db *gorm.DB
}

func NewWaitlistService(db *gorm.DB) *WaitlistService {
	return &WaitlistService{db: db}
}

// Create registra um paciente na lista de espera.
func (s *WaitlistService) Create(createdBy uuid.UUID, req *dto.CreateWaitlistEntryRequest) (*dto.WaitlistEntryResponse, error) {
	patientID, err := uuid.Parse(req.PatientID)
	if err != nil {
		return nil, errors.New("invalid patient id")
	}

	entry := models.WaitlistEntry{
		PatientID:       patientID,
		PreferredType:   req.PreferredType,
		Notes:           req.Notes,
		Status:          models.WaitlistWaiting,
		CreatedByUserID: createdBy,
	}
	if req.DoctorID != nil && *req.DoctorID != "" {
		did, err := uuid.Parse(*req.DoctorID)
		if err != nil {
			return nil, errors.New("invalid doctor id")
		}
		entry.DoctorID = &did
	}

	if err := s.db.Create(&entry).Error; err != nil {
		return nil, err
	}
	return s.reload(entry.ID)
}

// List lista entradas, por padrão só as que aguardam (status=waiting).
// statusFilter vazio => waiting; "all" => todas.
func (s *WaitlistService) List(statusFilter string) ([]dto.WaitlistEntryResponse, error) {
	var entries []models.WaitlistEntry
	query := s.db.Preload("Patient").Preload("Doctor").Order("created_at ASC")

	switch statusFilter {
	case "", "waiting":
		query = query.Where("status = ?", models.WaitlistWaiting)
	case "all":
		// sem filtro
	default:
		query = query.Where("status = ?", statusFilter)
	}

	if err := query.Find(&entries).Error; err != nil {
		return nil, err
	}
	result := make([]dto.WaitlistEntryResponse, len(entries))
	for i := range entries {
		result[i] = *s.toDTO(&entries[i])
	}
	return result, nil
}

// Update altera status/observações (ex.: converter para scheduled ou cancelar).
func (s *WaitlistService) Update(id uuid.UUID, req *dto.UpdateWaitlistEntryRequest) (*dto.WaitlistEntryResponse, error) {
	var entry models.WaitlistEntry
	if err := s.db.Where("id = ?", id).First(&entry).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrWaitlistEntryNotFound
		}
		return nil, err
	}

	if req.Status != nil {
		entry.Status = *req.Status
	}
	if req.Notes != nil {
		entry.Notes = req.Notes
	}
	if req.ScheduledAppointmentID != nil && *req.ScheduledAppointmentID != "" {
		aid, err := uuid.Parse(*req.ScheduledAppointmentID)
		if err != nil {
			return nil, errors.New("invalid appointment id")
		}
		entry.ScheduledAppointmentID = &aid
	}

	if err := s.db.Save(&entry).Error; err != nil {
		return nil, err
	}
	return s.reload(entry.ID)
}

// Delete remove (soft delete) uma entrada.
func (s *WaitlistService) Delete(id uuid.UUID) error {
	result := s.db.Where("id = ?", id).Delete(&models.WaitlistEntry{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrWaitlistEntryNotFound
	}
	return nil
}

func (s *WaitlistService) reload(id uuid.UUID) (*dto.WaitlistEntryResponse, error) {
	var entry models.WaitlistEntry
	if err := s.db.Preload("Patient").Preload("Doctor").Where("id = ?", id).First(&entry).Error; err != nil {
		return nil, err
	}
	return s.toDTO(&entry), nil
}

func (s *WaitlistService) toDTO(e *models.WaitlistEntry) *dto.WaitlistEntryResponse {
	resp := &dto.WaitlistEntryResponse{
		ID:              e.ID.String(),
		PatientID:       e.PatientID.String(),
		PreferredType:   e.PreferredType,
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
