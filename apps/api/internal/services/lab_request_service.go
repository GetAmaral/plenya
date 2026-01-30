package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
)

// LabRequestService handles business logic for lab requests
type LabRequestService struct {
	repo *repository.LabRequestRepository
}

// NewLabRequestService creates a new lab request service
func NewLabRequestService(repo *repository.LabRequestRepository) *LabRequestService {
	return &LabRequestService{repo: repo}
}

// CreateLabRequest creates a new lab request
func (s *LabRequestService) CreateLabRequest(req *models.LabRequest) error {
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
