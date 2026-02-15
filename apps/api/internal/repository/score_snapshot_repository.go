package repository

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// ScoreSnapshotRepository handles database operations for patient score snapshots
type ScoreSnapshotRepository struct {
	db *gorm.DB
}

// NewScoreSnapshotRepository creates a new score snapshot repository instance
func NewScoreSnapshotRepository(db *gorm.DB) *ScoreSnapshotRepository {
	return &ScoreSnapshotRepository{db: db}
}

// GetByID retrieves a snapshot by ID without relations
func (r *ScoreSnapshotRepository) GetByID(id uuid.UUID) (*models.PatientScoreSnapshot, error) {
	var snapshot models.PatientScoreSnapshot
	err := r.db.First(&snapshot, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("snapshot not found")
		}
		return nil, err
	}
	return &snapshot, nil
}

// GetByIDWithRelations retrieves a snapshot by ID with all related data preloaded
func (r *ScoreSnapshotRepository) GetByIDWithRelations(id uuid.UUID) (*models.PatientScoreSnapshot, error) {
	var snapshot models.PatientScoreSnapshot
	err := r.db.
		Preload("Patient").
		Preload("CalculatedByUser").
		Preload("GroupResults", func(db *gorm.DB) *gorm.DB {
			return db.Order("score_percentage DESC")
		}).
		Preload("GroupResults.Group", "deleted_at IS NULL").
		Preload("ItemResults", func(db *gorm.DB) *gorm.DB {
			return db.Order("actual_points DESC")
		}).
		Preload("ItemResults.Item", "deleted_at IS NULL").
		Preload("ItemResults.Item.Subgroup", "deleted_at IS NULL").
		Preload("ItemResults.Group", "deleted_at IS NULL").
		Preload("ItemResults.LevelMatched", "deleted_at IS NULL").
		Preload("ItemResults.LabResult", "deleted_at IS NULL").
		Preload("ItemResults.LabResult.LabResultBatch", "deleted_at IS NULL").
		Preload("ItemResults.AnamnesisItem", func(db *gorm.DB) *gorm.DB {
			return db.Unscoped() // Load even if soft deleted (historical data)
		}).
		Preload("ItemResults.AnamnesisItem.Anamnesis", "deleted_at IS NULL"). // Filter Anamnesis - only load if not deleted
		First(&snapshot, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("snapshot not found")
		}
		return nil, err
	}
	return &snapshot, nil
}

// GetAllByPatientID retrieves all snapshots for a patient, ordered by date DESC
func (r *ScoreSnapshotRepository) GetAllByPatientID(patientID uuid.UUID) ([]models.PatientScoreSnapshot, error) {
	var snapshots []models.PatientScoreSnapshot
	err := r.db.
		Where("patient_id = ?", patientID).
		Preload("CalculatedByUser").
		Order("calculated_at DESC").
		Find(&snapshots).Error
	return snapshots, err
}

// GetLatestByPatientID retrieves the most recent snapshot for a patient with all relations
func (r *ScoreSnapshotRepository) GetLatestByPatientID(patientID uuid.UUID) (*models.PatientScoreSnapshot, error) {
	var snapshot models.PatientScoreSnapshot
	err := r.db.
		Where("patient_id = ?", patientID).
		Preload("Patient").
		Preload("CalculatedByUser").
		Preload("GroupResults", func(db *gorm.DB) *gorm.DB {
			return db.Order("score_percentage DESC")
		}).
		Preload("GroupResults.Group", "deleted_at IS NULL").
		Preload("ItemResults", func(db *gorm.DB) *gorm.DB {
			return db.Order("actual_points DESC")
		}).
		Preload("ItemResults.Item", "deleted_at IS NULL").
		Preload("ItemResults.Item.Subgroup", "deleted_at IS NULL").
		Preload("ItemResults.Group", "deleted_at IS NULL").
		Preload("ItemResults.LevelMatched", "deleted_at IS NULL").
		Preload("ItemResults.LabResult", "deleted_at IS NULL").
		Preload("ItemResults.LabResult.LabResultBatch", "deleted_at IS NULL").
		Preload("ItemResults.AnamnesisItem", func(db *gorm.DB) *gorm.DB {
			return db.Unscoped() // Load even if soft deleted (historical data)
		}).
		Preload("ItemResults.AnamnesisItem.Anamnesis", "deleted_at IS NULL"). // Filter Anamnesis - only load if not deleted
		Order("calculated_at DESC").
		First(&snapshot).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no snapshots found for patient")
		}
		return nil, err
	}
	return &snapshot, nil
}

// GetTodaySnapshotByPatientID retrieves the snapshot for today (if exists)
func (r *ScoreSnapshotRepository) GetTodaySnapshotByPatientID(patientID uuid.UUID) (*models.PatientScoreSnapshot, error) {
	var snapshot models.PatientScoreSnapshot
	err := r.db.
		Where("patient_id = ?", patientID).
		Where("DATE(calculated_at) = CURRENT_DATE").
		First(&snapshot).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Return nil, nil if not found (not an error)
		}
		return nil, err
	}
	return &snapshot, nil
}

// Create creates a new snapshot (used by service layer in transaction)
func (r *ScoreSnapshotRepository) Create(snapshot *models.PatientScoreSnapshot) error {
	return r.db.Create(snapshot).Error
}

// CreateWithTransaction creates a new snapshot within a transaction
func (r *ScoreSnapshotRepository) CreateWithTransaction(tx *gorm.DB, snapshot *models.PatientScoreSnapshot) error {
	return tx.Create(snapshot).Error
}

// DeleteSnapshot soft deletes a snapshot by ID
func (r *ScoreSnapshotRepository) DeleteSnapshot(id uuid.UUID) error {
	result := r.db.Delete(&models.PatientScoreSnapshot{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("snapshot not found")
	}
	return nil
}
