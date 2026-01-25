package repository

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// LabTestDefinitionRepository handles database operations for lab test definitions
type LabTestDefinitionRepository struct {
	db *gorm.DB
}

// NewLabTestDefinitionRepository creates a new lab test definition repository instance
func NewLabTestDefinitionRepository(db *gorm.DB) *LabTestDefinitionRepository {
	return &LabTestDefinitionRepository{db: db}
}

// ============================================================
// LabTestDefinition Operations
// ============================================================

// CreateLabTestDefinition creates a new lab test definition
func (r *LabTestDefinitionRepository) CreateLabTestDefinition(def *models.LabTestDefinition) error {
	return r.db.Create(def).Error
}

// GetLabTestDefinitionByID retrieves a lab test definition by ID
func (r *LabTestDefinitionRepository) GetLabTestDefinitionByID(id uuid.UUID) (*models.LabTestDefinition, error) {
	var def models.LabTestDefinition
	err := r.db.
		Preload("ParentTest").
		Preload("SubTests", func(db *gorm.DB) *gorm.DB {
			return db.Where("is_active = ?", true).Order("display_order ASC, name ASC")
		}).
		Preload("ScoreMappings", func(db *gorm.DB) *gorm.DB {
			return db.Where("is_active = ?", true)
		}).
		First(&def, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("lab test definition not found")
		}
		return nil, err
	}
	return &def, nil
}

// GetLabTestDefinitionByCode retrieves a lab test definition by code
func (r *LabTestDefinitionRepository) GetLabTestDefinitionByCode(code string) (*models.LabTestDefinition, error) {
	var def models.LabTestDefinition
	err := r.db.
		Preload("ParentTest").
		Preload("SubTests", func(db *gorm.DB) *gorm.DB {
			return db.Where("is_active = ?", true).Order("display_order ASC, name ASC")
		}).
		Preload("ScoreMappings", func(db *gorm.DB) *gorm.DB {
			return db.Where("is_active = ?", true)
		}).
		First(&def, "code = ?", code).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("lab test definition not found")
		}
		return nil, err
	}
	return &def, nil
}

// GetAllLabTestDefinitions retrieves all lab test definitions
func (r *LabTestDefinitionRepository) GetAllLabTestDefinitions() ([]models.LabTestDefinition, error) {
	var defs []models.LabTestDefinition
	err := r.db.
		Where("is_active = ?", true).
		Order("category ASC, display_order ASC, name ASC").
		Find(&defs).Error
	return defs, err
}

// GetRequestableLabTests retrieves only requestable lab tests (those that can be ordered)
func (r *LabTestDefinitionRepository) GetRequestableLabTests(category *models.LabTestCategory) ([]models.LabTestDefinition, error) {
	query := r.db.Where("is_requestable = ? AND is_active = ?", true, true)

	if category != nil {
		query = query.Where("category = ?", *category)
	}

	var defs []models.LabTestDefinition
	err := query.Order("category ASC, display_order ASC, name ASC").Find(&defs).Error
	return defs, err
}

// GetSubTests retrieves all sub-tests for a parent test
func (r *LabTestDefinitionRepository) GetSubTests(parentID uuid.UUID) ([]models.LabTestDefinition, error) {
	var defs []models.LabTestDefinition
	err := r.db.
		Where("parent_test_id = ? AND is_active = ?", parentID, true).
		Order("display_order ASC, name ASC").
		Find(&defs).Error
	return defs, err
}

// SearchLabTestDefinitions searches lab tests by name or code
func (r *LabTestDefinitionRepository) SearchLabTestDefinitions(searchTerm string) ([]models.LabTestDefinition, error) {
	var defs []models.LabTestDefinition
	err := r.db.
		Where("(name ILIKE ? OR code ILIKE ? OR short_name ILIKE ?) AND is_active = ?",
			"%"+searchTerm+"%", "%"+searchTerm+"%", "%"+searchTerm+"%", true).
		Order("is_requestable DESC, name ASC").
		Limit(50).
		Find(&defs).Error
	return defs, err
}

// UpdateLabTestDefinition updates an existing lab test definition
func (r *LabTestDefinitionRepository) UpdateLabTestDefinition(def *models.LabTestDefinition) error {
	return r.db.Save(def).Error
}

// DeleteLabTestDefinition soft deletes a lab test definition
func (r *LabTestDefinitionRepository) DeleteLabTestDefinition(id uuid.UUID) error {
	result := r.db.Delete(&models.LabTestDefinition{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("lab test definition not found")
	}
	return nil
}

// ============================================================
// LabTestScoreMapping Operations
// ============================================================

// CreateLabTestScoreMapping creates a new mapping
func (r *LabTestDefinitionRepository) CreateLabTestScoreMapping(mapping *models.LabTestScoreMapping) error {
	return r.db.Create(mapping).Error
}

// GetLabTestScoreMappingByID retrieves a mapping by ID
func (r *LabTestDefinitionRepository) GetLabTestScoreMappingByID(id uuid.UUID) (*models.LabTestScoreMapping, error) {
	var mapping models.LabTestScoreMapping
	err := r.db.
		Preload("LabTest").
		Preload("ScoreItem").
		First(&mapping, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("lab test score mapping not found")
		}
		return nil, err
	}
	return &mapping, nil
}

// GetMappingsForLabTest retrieves all score mappings for a lab test
func (r *LabTestDefinitionRepository) GetMappingsForLabTest(labTestID uuid.UUID) ([]models.LabTestScoreMapping, error) {
	var mappings []models.LabTestScoreMapping
	err := r.db.
		Where("lab_test_id = ? AND is_active = ?", labTestID, true).
		Preload("ScoreItem").
		Find(&mappings).Error
	return mappings, err
}

// GetMappingForPatient retrieves the appropriate mapping for a specific patient
func (r *LabTestDefinitionRepository) GetMappingForPatient(labTestID uuid.UUID, gender models.Gender, age int) (*models.LabTestScoreMapping, error) {
	var mapping models.LabTestScoreMapping
	err := r.db.
		Where("lab_test_id = ? AND is_active = ?", labTestID, true).
		Where("(gender IS NULL OR gender = ?)", gender).
		Where("(min_age IS NULL OR min_age <= ?)", age).
		Where("(max_age IS NULL OR max_age >= ?)", age).
		Preload("ScoreItem").
		Preload("ScoreItem.Levels", func(db *gorm.DB) *gorm.DB {
			return db.Order("level ASC")
		}).
		First(&mapping).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no matching score mapping found for patient")
		}
		return nil, err
	}
	return &mapping, nil
}

// UpdateLabTestScoreMapping updates a mapping
func (r *LabTestDefinitionRepository) UpdateLabTestScoreMapping(mapping *models.LabTestScoreMapping) error {
	return r.db.Save(mapping).Error
}

// DeleteLabTestScoreMapping soft deletes a mapping
func (r *LabTestDefinitionRepository) DeleteLabTestScoreMapping(id uuid.UUID) error {
	result := r.db.Delete(&models.LabTestScoreMapping{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("lab test score mapping not found")
	}
	return nil
}
