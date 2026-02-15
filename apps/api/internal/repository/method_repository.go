package repository

import (
	"github.com/plenya/api/internal/models"

	"gorm.io/gorm"
)

type MethodRepository struct {
	db *gorm.DB
}

func NewMethodRepository(db *gorm.DB) *MethodRepository {
	return &MethodRepository{db: db}
}

// GetMethodWithTree returns method with full hierarchy (letters → pillars → items)
func (r *MethodRepository) GetMethodWithTree(methodID string) (*models.Method, error) {
	var method models.Method
	err := r.db.
		Preload("Letters", func(db *gorm.DB) *gorm.DB {
			return db.Order("\"order\" ASC")
		}).
		Preload("Letters.Pillars", func(db *gorm.DB) *gorm.DB {
			return db.Order("\"order\" ASC")
		}).
		Preload("Letters.Pillars.ScoreItems", func(db *gorm.DB) *gorm.DB {
			return db.Order("\"order\" ASC")
		}).
		Preload("Letters.Pillars.ScoreItems.Levels", func(db *gorm.DB) *gorm.DB {
			return db.Order("level ASC")
		}).
		First(&method, "id = ?", methodID).Error

	return &method, err
}

// GetAllMethodsWithTree returns all methods with full hierarchy
func (r *MethodRepository) GetAllMethodsWithTree() ([]models.Method, error) {
	var methods []models.Method
	err := r.db.
		Preload("Letters", func(db *gorm.DB) *gorm.DB {
			return db.Order("\"order\" ASC")
		}).
		Preload("Letters.Pillars", func(db *gorm.DB) *gorm.DB {
			return db.Order("\"order\" ASC")
		}).
		Preload("Letters.Pillars.ScoreItems", func(db *gorm.DB) *gorm.DB {
			return db.Order("\"order\" ASC")
		}).
		Preload("Letters.Pillars.ScoreItems.Levels").
		Order("\"order\" ASC").
		Find(&methods).Error

	return methods, err
}

// GetAllMethods returns all methods without relationships
func (r *MethodRepository) GetAllMethods() ([]models.Method, error) {
	var methods []models.Method
	err := r.db.Order("\"order\" ASC").Find(&methods).Error
	return methods, err
}

// GetMethodByID returns a method by ID
func (r *MethodRepository) GetMethodByID(methodID string) (*models.Method, error) {
	var method models.Method
	err := r.db.First(&method, "id = ?", methodID).Error
	return &method, err
}

// CreateMethod creates a new method
func (r *MethodRepository) CreateMethod(method *models.Method) error {
	return r.db.Create(method).Error
}

// UpdateMethod updates an existing method
func (r *MethodRepository) UpdateMethod(method *models.Method) error {
	return r.db.Save(method).Error
}

// DeleteMethod soft deletes a method
func (r *MethodRepository) DeleteMethod(methodID string) error {
	return r.db.Delete(&models.Method{}, "id = ?", methodID).Error
}

// GetMethodLetters returns all letters for a method
func (r *MethodRepository) GetMethodLetters(methodID string) ([]models.MethodLetter, error) {
	var letters []models.MethodLetter
	err := r.db.Where("method_id = ?", methodID).Order("\"order\" ASC").Find(&letters).Error
	return letters, err
}

// GetMethodLetterByID returns a letter by ID
func (r *MethodRepository) GetMethodLetterByID(letterID string) (*models.MethodLetter, error) {
	var letter models.MethodLetter
	err := r.db.First(&letter, "id = ?", letterID).Error
	return &letter, err
}

// CreateMethodLetter creates a new letter
func (r *MethodRepository) CreateMethodLetter(letter *models.MethodLetter) error {
	return r.db.Create(letter).Error
}

// UpdateMethodLetter updates an existing letter
func (r *MethodRepository) UpdateMethodLetter(letter *models.MethodLetter) error {
	return r.db.Save(letter).Error
}

// DeleteMethodLetter soft deletes a letter
func (r *MethodRepository) DeleteMethodLetter(letterID string) error {
	return r.db.Delete(&models.MethodLetter{}, "id = ?", letterID).Error
}

// GetLetterPillars returns all pillars for a letter
func (r *MethodRepository) GetLetterPillars(letterID string) ([]models.MethodPillar, error) {
	var pillars []models.MethodPillar
	err := r.db.Where("letter_id = ?", letterID).Order("\"order\" ASC").Find(&pillars).Error
	return pillars, err
}

// GetMethodPillarByID returns a pillar by ID
func (r *MethodRepository) GetMethodPillarByID(pillarID string) (*models.MethodPillar, error) {
	var pillar models.MethodPillar
	err := r.db.Preload("ScoreItems").Preload("ScoreItems.Levels").First(&pillar, "id = ?", pillarID).Error
	return &pillar, err
}

// CreateMethodPillar creates a new pillar
func (r *MethodRepository) CreateMethodPillar(pillar *models.MethodPillar) error {
	return r.db.Create(pillar).Error
}

// UpdateMethodPillar updates an existing pillar
func (r *MethodRepository) UpdateMethodPillar(pillar *models.MethodPillar) error {
	return r.db.Save(pillar).Error
}

// DeleteMethodPillar soft deletes a pillar
func (r *MethodRepository) DeleteMethodPillar(pillarID string) error {
	return r.db.Delete(&models.MethodPillar{}, "id = ?", pillarID).Error
}

// AssignScoreItemToPillar adds a score item to a pillar (M:N)
func (r *MethodRepository) AssignScoreItemToPillar(itemID, pillarID string) error {
	return r.db.Exec(
		"INSERT INTO score_item_method_pillars (score_item_id, method_pillar_id) VALUES (?, ?) ON CONFLICT DO NOTHING",
		itemID, pillarID,
	).Error
}

// UnassignScoreItemFromPillar removes association
func (r *MethodRepository) UnassignScoreItemFromPillar(itemID, pillarID string) error {
	return r.db.Exec(
		"DELETE FROM score_item_method_pillars WHERE score_item_id = ? AND method_pillar_id = ?",
		itemID, pillarID,
	).Error
}

// GetUnassignedScoreItems returns items not in any pillar
func (r *MethodRepository) GetUnassignedScoreItems() ([]models.ScoreItem, error) {
	var items []models.ScoreItem
	err := r.db.
		Preload("Levels").
		Preload("Subgroup").
		Preload("Subgroup.Group").
		Where("id NOT IN (SELECT DISTINCT score_item_id FROM score_item_method_pillars)").
		Order("\"order\" ASC").
		Find(&items).Error

	return items, err
}

// GetAllScoreItemsWithPillars returns all score items with their method pillars
func (r *MethodRepository) GetAllScoreItemsWithPillars() ([]models.ScoreItem, error) {
	var items []models.ScoreItem
	err := r.db.
		Preload("Levels").
		Preload("Subgroup").
		Preload("Subgroup.Group").
		Preload("MethodPillars").
		Order("\"order\" ASC").
		Find(&items).Error

	return items, err
}
