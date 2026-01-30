package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
)

// LabRequestService handles business logic for lab requests
type LabRequestService struct {
	repo *repository.LabRequestRepository
	db   *gorm.DB
}

// NewLabRequestService creates a new lab request service
func NewLabRequestService(repo *repository.LabRequestRepository, db *gorm.DB) *LabRequestService {
	return &LabRequestService{
		repo: repo,
		db:   db,
	}
}

// CreateLabRequest creates a new lab request
func (s *LabRequestService) CreateLabRequest(userID uuid.UUID, req *models.LabRequest) error {
	// CRITICAL SECURITY: Get user's selected patient
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return err
	}

	// If no selected patient, cannot create lab request
	if user.SelectedPatientID == nil {
		return errors.New("no patient selected - please select a patient first")
	}

	// If patientID is provided, validate it matches selectedPatient
	if req.PatientID != uuid.Nil {
		if req.PatientID != *user.SelectedPatientID {
			return errors.New("patient id does not match selected patient")
		}
	} else {
		// Auto-fill with selectedPatient
		req.PatientID = *user.SelectedPatientID
	}

	// Auto-fill doctorID if not provided
	if req.DoctorID == nil {
		req.DoctorID = &userID
	}

	return s.repo.CreateLabRequest(req)
}

// GetLabRequestByID retrieves a lab request by ID
func (s *LabRequestService) GetLabRequestByID(id uuid.UUID) (*models.LabRequest, error) {
	return s.repo.GetLabRequestByID(id)
}

// GetLabRequestsByPatientID retrieves all lab requests for a patient
func (s *LabRequestService) GetLabRequestsByPatientID(patientID uuid.UUID) ([]models.LabRequest, error) {
	return s.repo.GetLabRequestsByPatientID(patientID)
}

// GetLabRequestsByDate retrieves all lab requests for a specific date
func (s *LabRequestService) GetLabRequestsByDate(date time.Time) ([]models.LabRequest, error) {
	return s.repo.GetLabRequestsByDate(date)
}

// GetLabRequestsByDateRange retrieves lab requests within a date range
func (s *LabRequestService) GetLabRequestsByDateRange(startDate, endDate time.Time) ([]models.LabRequest, error) {
	return s.repo.GetLabRequestsByDateRange(startDate, endDate)
}

// GetAllLabRequests retrieves all lab requests with pagination
func (s *LabRequestService) GetAllLabRequests(limit, offset int) ([]models.LabRequest, int64, error) {
	return s.repo.GetAllLabRequests(limit, offset)
}

// UpdateLabRequest updates an existing lab request
func (s *LabRequestService) UpdateLabRequest(id uuid.UUID, req *models.LabRequest) error {
	existing, err := s.repo.GetLabRequestByID(id)
	if err != nil {
		return err
	}

	// Update fields
	existing.PatientID = req.PatientID
	existing.Date = req.Date
	existing.Exams = req.Exams
	existing.Notes = req.Notes
	existing.DoctorID = req.DoctorID

	return s.repo.UpdateLabRequest(existing)
}

// DeleteLabRequest soft deletes a lab request
func (s *LabRequestService) DeleteLabRequest(id uuid.UUID) error {
	return s.repo.DeleteLabRequest(id)
}
