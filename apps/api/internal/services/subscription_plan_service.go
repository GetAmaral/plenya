package services

import (
	"github.com/google/uuid"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
)

type SubscriptionPlanService struct {
	repo *repository.SubscriptionPlanRepository
}

func NewSubscriptionPlanService(repo *repository.SubscriptionPlanRepository) *SubscriptionPlanService {
	return &SubscriptionPlanService{repo: repo}
}

// CreatePlanDTO represents the data needed to create a subscription plan
type CreatePlanDTO struct {
	Name            string                  `json:"name" binding:"required,min=3,max=255"`
	Description     string                  `json:"description"`
	Features        string                  `json:"features"` // JSON string
	Price           float64                 `json:"price" binding:"required,min=0"`
	Currency        string                  `json:"currency" binding:"required,len=3"`
	BillingCycle    models.BillingCycle     `json:"billingCycle" binding:"required,oneof=monthly quarterly yearly one_time"`
	MethodID        *uuid.UUID              `json:"methodId"`
	IsActive        bool                    `json:"isActive"`
	TrialPeriodDays int                     `json:"trialPeriodDays" binding:"min=0"`
	Order           int                     `json:"order" binding:"min=0"`
}

// UpdatePlanDTO represents the data needed to update a subscription plan
type UpdatePlanDTO struct {
	Name            *string                 `json:"name" binding:"omitempty,min=3,max=255"`
	Description     *string                 `json:"description"`
	Features        *string                 `json:"features"`
	Price           *float64                `json:"price" binding:"omitempty,min=0"`
	Currency        *string                 `json:"currency" binding:"omitempty,len=3"`
	BillingCycle    *models.BillingCycle    `json:"billingCycle" binding:"omitempty,oneof=monthly quarterly yearly one_time"`
	MethodID        *uuid.UUID              `json:"methodId"`
	IsActive        *bool                   `json:"isActive"`
	TrialPeriodDays *int                    `json:"trialPeriodDays" binding:"omitempty,min=0"`
	Order           *int                    `json:"order" binding:"omitempty,min=0"`
}

// GetAll returns all subscription plans
func (s *SubscriptionPlanService) GetAll() ([]models.SubscriptionPlan, error) {
	return s.repo.GetAll()
}

// GetActive returns only active subscription plans
func (s *SubscriptionPlanService) GetActive() ([]models.SubscriptionPlan, error) {
	return s.repo.GetActive()
}

// GetByID returns a subscription plan by ID
func (s *SubscriptionPlanService) GetByID(id uuid.UUID) (*models.SubscriptionPlan, error) {
	return s.repo.GetByID(id)
}

// GetByMethodID returns all plans for a specific method
func (s *SubscriptionPlanService) GetByMethodID(methodID uuid.UUID) ([]models.SubscriptionPlan, error) {
	return s.repo.GetByMethodID(methodID)
}

// Create creates a new subscription plan
func (s *SubscriptionPlanService) Create(dto CreatePlanDTO) (*models.SubscriptionPlan, error) {
	plan := &models.SubscriptionPlan{
		Name:            dto.Name,
		Description:     dto.Description,
		Features:        dto.Features,
		Price:           dto.Price,
		Currency:        dto.Currency,
		BillingCycle:    dto.BillingCycle,
		MethodID:        dto.MethodID,
		IsActive:        dto.IsActive,
		TrialPeriodDays: dto.TrialPeriodDays,
		Order:           dto.Order,
	}

	if err := s.repo.Create(plan); err != nil {
		return nil, err
	}

	return plan, nil
}

// Update updates an existing subscription plan
func (s *SubscriptionPlanService) Update(id uuid.UUID, dto UpdatePlanDTO) (*models.SubscriptionPlan, error) {
	plan, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update only provided fields
	if dto.Name != nil {
		plan.Name = *dto.Name
	}
	if dto.Description != nil {
		plan.Description = *dto.Description
	}
	if dto.Features != nil {
		plan.Features = *dto.Features
	}
	if dto.Price != nil {
		plan.Price = *dto.Price
	}
	if dto.Currency != nil {
		plan.Currency = *dto.Currency
	}
	if dto.BillingCycle != nil {
		plan.BillingCycle = *dto.BillingCycle
	}
	if dto.MethodID != nil {
		plan.MethodID = dto.MethodID
	}
	if dto.IsActive != nil {
		plan.IsActive = *dto.IsActive
	}
	if dto.TrialPeriodDays != nil {
		plan.TrialPeriodDays = *dto.TrialPeriodDays
	}
	if dto.Order != nil {
		plan.Order = *dto.Order
	}

	if err := s.repo.Update(plan); err != nil {
		return nil, err
	}

	return plan, nil
}

// Delete soft deletes a subscription plan
func (s *SubscriptionPlanService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}

// Activate sets plan as active
func (s *SubscriptionPlanService) Activate(id uuid.UUID) error {
	return s.repo.Activate(id)
}

// Deactivate sets plan as inactive
func (s *SubscriptionPlanService) Deactivate(id uuid.UUID) error {
	return s.repo.Deactivate(id)
}

// UpdateOrder updates the order of multiple plans
func (s *SubscriptionPlanService) UpdateOrder(orderMap map[uuid.UUID]int) error {
	return s.repo.UpdateOrder(orderMap)
}
