package repository

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// LabRequestRepository handles database operations for lab requests
type LabRequestRepository struct {
	db *gorm.DB
}

// NewLabRequestRepository creates a new lab request repository instance
func NewLabRequestRepository(db *gorm.DB) *LabRequestRepository {
	return &LabRequestRepository{db: db}
}

// CreateLabRequest creates a new lab request
func (r *LabRequestRepository) CreateLabRequest(req *models.LabRequest) error {
	return r.db.Create(req).Error
}

// GetLabRequestByID retrieves a lab request by ID
func (r *LabRequestRepository) GetLabRequestByID(id uuid.UUID) (*models.LabRequest, error) {
	var req models.LabRequest
	err := r.db.
		Preload("Patient").
		Preload("Doctor").
		First(&req, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("lab request not found")
		}
		return nil, err
	}
	return &req, nil
}

// GetLabRequestsByPatientID retrieves all lab requests for a patient
func (r *LabRequestRepository) GetLabRequestsByPatientID(patientID uuid.UUID) ([]models.LabRequest, error) {
	var reqs []models.LabRequest
	err := r.db.
		Where("patient_id = ?", patientID).
		Preload("Doctor").
		Order("date DESC, created_at DESC").
		Find(&reqs).Error
	return reqs, err
}

// GetLabRequestsByDate retrieves all lab requests for a specific date
func (r *LabRequestRepository) GetLabRequestsByDate(date time.Time) ([]models.LabRequest, error) {
	var reqs []models.LabRequest
	err := r.db.
		Where("date = ?", date).
		Preload("Patient").
		Preload("Doctor").
		Order("created_at DESC").
		Find(&reqs).Error
	return reqs, err
}

// GetLabRequestsByDateRange retrieves lab requests within a date range
func (r *LabRequestRepository) GetLabRequestsByDateRange(startDate, endDate time.Time) ([]models.LabRequest, error) {
	var reqs []models.LabRequest
	err := r.db.
		Where("date BETWEEN ? AND ?", startDate, endDate).
		Preload("Patient").
		Preload("Doctor").
		Order("date DESC, created_at DESC").
		Find(&reqs).Error
	return reqs, err
}

// GetAllLabRequests retrieves all lab requests with pagination
func (r *LabRequestRepository) GetAllLabRequests(limit, offset int) ([]models.LabRequest, int64, error) {
	var reqs []models.LabRequest
	var total int64

	// Count total
	if err := r.db.Model(&models.LabRequest{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated results
	err := r.db.
		Preload("Patient").
		Preload("Doctor").
		Order("date DESC, created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&reqs).Error

	return reqs, total, err
}

// UpdateLabRequest updates an existing lab request
func (r *LabRequestRepository) UpdateLabRequest(req *models.LabRequest) error {
	return r.db.Save(req).Error
}

// DeleteLabRequest soft deletes a lab request
func (r *LabRequestRepository) DeleteLabRequest(id uuid.UUID) error {
	result := r.db.Delete(&models.LabRequest{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("lab request not found")
	}
	return nil
}
