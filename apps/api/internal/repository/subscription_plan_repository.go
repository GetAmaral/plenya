package repository

import (
	"github.com/google/uuid"
	"github.com/plenya/api/internal/models"
	"gorm.io/gorm"
)

type SubscriptionPlanRepository struct {
	db *gorm.DB
}

func NewSubscriptionPlanRepository(db *gorm.DB) *SubscriptionPlanRepository {
	return &SubscriptionPlanRepository{db: db}
}

// GetAll returns all subscription plans ordered by order field
func (r *SubscriptionPlanRepository) GetAll() ([]models.SubscriptionPlan, error) {
	var plans []models.SubscriptionPlan
	err := r.db.Preload("Method").Order("\"order\" ASC").Find(&plans).Error
	return plans, err
}

// GetActive returns only active subscription plans
func (r *SubscriptionPlanRepository) GetActive() ([]models.SubscriptionPlan, error) {
	var plans []models.SubscriptionPlan
	err := r.db.Preload("Method").
		Where("is_active = ?", true).
		Order("\"order\" ASC").
		Find(&plans).Error
	return plans, err
}

// GetByID returns a subscription plan by ID
func (r *SubscriptionPlanRepository) GetByID(id uuid.UUID) (*models.SubscriptionPlan, error) {
	var plan models.SubscriptionPlan
	err := r.db.Preload("Method").First(&plan, id).Error
	if err != nil {
		return nil, err
	}
	return &plan, nil
}

// GetByMethodID returns all plans for a specific method
func (r *SubscriptionPlanRepository) GetByMethodID(methodID uuid.UUID) ([]models.SubscriptionPlan, error) {
	var plans []models.SubscriptionPlan
	err := r.db.Preload("Method").
		Where("method_id = ?", methodID).
		Order("\"order\" ASC").
		Find(&plans).Error
	return plans, err
}

// Create creates a new subscription plan
func (r *SubscriptionPlanRepository) Create(plan *models.SubscriptionPlan) error {
	return r.db.Create(plan).Error
}

// Update updates an existing subscription plan
func (r *SubscriptionPlanRepository) Update(plan *models.SubscriptionPlan) error {
	return r.db.Save(plan).Error
}

// Delete soft deletes a subscription plan
func (r *SubscriptionPlanRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.SubscriptionPlan{}, id).Error
}

// Activate sets plan as active
func (r *SubscriptionPlanRepository) Activate(id uuid.UUID) error {
	return r.db.Model(&models.SubscriptionPlan{}).
		Where("id = ?", id).
		Update("is_active", true).Error
}

// Deactivate sets plan as inactive
func (r *SubscriptionPlanRepository) Deactivate(id uuid.UUID) error {
	return r.db.Model(&models.SubscriptionPlan{}).
		Where("id = ?", id).
		Update("is_active", false).Error
}

// UpdateOrder updates the order of multiple plans
func (r *SubscriptionPlanRepository) UpdateOrder(orderMap map[uuid.UUID]int) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		for id, order := range orderMap {
			if err := tx.Model(&models.SubscriptionPlan{}).
				Where("id = ?", id).
				Update("order", order).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
