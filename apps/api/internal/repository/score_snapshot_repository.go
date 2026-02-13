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
		Preload("GroupResults.Group").
		Preload("ItemResults", func(db *gorm.DB) *gorm.DB {
			return db.Order("actual_points DESC")
		}).
		Preload("ItemResults.Item").
		Preload("ItemResults.Item.Subgroup").
		Preload("ItemResults.Group").
		Preload("ItemResults.LevelMatched").
		Preload("ItemResults.LabResult").
		Preload("ItemResults.LabResult.LabResultBatch").
		Preload("ItemResults.AnamnesisItem").
		Preload("ItemResults.AnamnesisItem.Anamnesis").
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
		Preload("GroupResults.Group").
		Preload("ItemResults", func(db *gorm.DB) *gorm.DB {
			return db.Order("actual_points DESC")
		}).
		Preload("ItemResults.Item").
		Preload("ItemResults.Item.Subgroup").
		Preload("ItemResults.Group").
		Preload("ItemResults.LevelMatched").
		Preload("ItemResults.LabResult").
		Preload("ItemResults.LabResult.LabResultBatch").
		Preload("ItemResults.AnamnesisItem").
		Preload("ItemResults.AnamnesisItem.Anamnesis").
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
