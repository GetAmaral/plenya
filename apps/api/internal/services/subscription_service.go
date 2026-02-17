package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
)

// DTOs for API
type CreateSubscriptionDTO struct {
	PatientID          string   `json:"patientId" validate:"required,uuid"`
	SubscriptionPlanID string   `json:"subscriptionPlanId" validate:"required,uuid"`
	Status             string   `json:"status" validate:"required,oneof=active inactive cancelled expired suspended trial"`
	AutoRenew          bool     `json:"autoRenew"`
	StartDate          string   `json:"startDate" validate:"required"` // ISO 8601
	EndDate            *string  `json:"endDate,omitempty"`
	TrialEndDate       *string  `json:"trialEndDate,omitempty"`
	DiscountPercent    float64  `json:"discountPercent" validate:"gte=0,lte=100"`
	DiscountReason     *string  `json:"discountReason,omitempty"`
	CustomPrice        *float64 `json:"customPrice,omitempty" validate:"omitempty,gte=0"`
	CustomTrialDays    *int     `json:"customTrialDays,omitempty" validate:"omitempty,gte=0"`
	Notes              *string  `json:"notes,omitempty"`
}

type UpdateSubscriptionDTO struct {
	Status             *string  `json:"status,omitempty" validate:"omitempty,oneof=active inactive cancelled expired suspended trial"`
	AutoRenew          *bool    `json:"autoRenew,omitempty"`
	EndDate            *string  `json:"endDate,omitempty"`
	TrialEndDate       *string  `json:"trialEndDate,omitempty"`
	NextBillingDate    *string  `json:"nextBillingDate,omitempty"`
	DiscountPercent    *float64 `json:"discountPercent,omitempty" validate:"omitempty,gte=0,lte=100"`
	DiscountReason     *string  `json:"discountReason,omitempty"`
	CustomPrice        *float64 `json:"customPrice,omitempty" validate:"omitempty,gte=0"`
	CustomTrialDays    *int     `json:"customTrialDays,omitempty" validate:"omitempty,gte=0"`
	CancellationReason *string  `json:"cancellationReason,omitempty"`
	Notes              *string  `json:"notes,omitempty"`
}

type CancelSubscriptionDTO struct {
	CancellationReason *string `json:"cancellationReason,omitempty"`
}

type RenewSubscriptionDTO struct {
	EndDate         *string `json:"endDate,omitempty"`
	NextBillingDate *string `json:"nextBillingDate,omitempty"`
}

type SubscriptionService struct {
	repo *repository.SubscriptionRepository
}

func NewSubscriptionService(repo *repository.SubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{repo: repo}
}

// GetByID busca um subscription por ID
func (s *SubscriptionService) GetByID(id string) (*models.PatientSubscription, error) {
	return s.repo.GetByID(id)
}

// GetByPatientID busca todos os subscriptions de um paciente
func (s *SubscriptionService) GetByPatientID(patientID string) ([]models.PatientSubscription, error) {
	return s.repo.GetByPatientID(patientID)
}

// GetActive busca todos os subscriptions ativos
func (s *SubscriptionService) GetActive() ([]models.PatientSubscription, error) {
	return s.repo.GetActive()
}

// GetAll busca todos os subscriptions
func (s *SubscriptionService) GetAll() ([]models.PatientSubscription, error) {
	return s.repo.GetAll()
}

// Create cria um novo subscription
func (s *SubscriptionService) Create(dto *CreateSubscriptionDTO) (*models.PatientSubscription, error) {
	patientUUID, err := uuid.Parse(dto.PatientID)
	if err != nil {
		return nil, err
	}

	planUUID, err := uuid.Parse(dto.SubscriptionPlanID)
	if err != nil {
		return nil, err
	}

	// Parse dates
	startDate, err := time.Parse("2006-01-02", dto.StartDate)
	if err != nil {
		return nil, err
	}

	var endDate *time.Time
	if dto.EndDate != nil {
		parsed, err := time.Parse("2006-01-02", *dto.EndDate)
		if err != nil {
			return nil, err
		}
		endDate = &parsed
	}

	var trialEndDate *time.Time
	if dto.TrialEndDate != nil {
		parsed, err := time.Parse("2006-01-02", *dto.TrialEndDate)
		if err != nil {
			return nil, err
		}
		trialEndDate = &parsed
	}

	subscription := &models.PatientSubscription{
		PatientID:          patientUUID,
		SubscriptionPlanID: planUUID,
		Status:             models.SubscriptionStatus(dto.Status),
		AutoRenew:          dto.AutoRenew,
		StartDate:          startDate,
		EndDate:            endDate,
		TrialEndDate:       trialEndDate,
		DiscountPercent:    dto.DiscountPercent,
		DiscountReason:     dto.DiscountReason,
		CustomPrice:        dto.CustomPrice,
		CustomTrialDays:    dto.CustomTrialDays,
		Notes:              dto.Notes,
	}

	// BeforeCreate hook will create plan snapshot
	if err := s.repo.Create(subscription); err != nil {
		return nil, err
	}

	return subscription, nil
}

// Update atualiza um subscription
func (s *SubscriptionService) Update(id string, dto *UpdateSubscriptionDTO) (*models.PatientSubscription, error) {
	subscription, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if dto.Status != nil {
		subscription.Status = models.SubscriptionStatus(*dto.Status)
	}

	if dto.AutoRenew != nil {
		subscription.AutoRenew = *dto.AutoRenew
	}

	if dto.EndDate != nil {
		parsed, err := time.Parse("2006-01-02", *dto.EndDate)
		if err != nil {
			return nil, err
		}
		subscription.EndDate = &parsed
	}

	if dto.TrialEndDate != nil {
		parsed, err := time.Parse("2006-01-02", *dto.TrialEndDate)
		if err != nil {
			return nil, err
		}
		subscription.TrialEndDate = &parsed
	}

	if dto.NextBillingDate != nil {
		parsed, err := time.Parse("2006-01-02", *dto.NextBillingDate)
		if err != nil {
			return nil, err
		}
		subscription.NextBillingDate = &parsed
	}

	if dto.DiscountPercent != nil {
		subscription.DiscountPercent = *dto.DiscountPercent
	}

	if dto.DiscountReason != nil {
		subscription.DiscountReason = dto.DiscountReason
	}

	if dto.CustomPrice != nil {
		subscription.CustomPrice = dto.CustomPrice
	}

	if dto.CustomTrialDays != nil {
		subscription.CustomTrialDays = dto.CustomTrialDays
	}

	if dto.CancellationReason != nil {
		subscription.CancellationReason = dto.CancellationReason
	}

	if dto.Notes != nil {
		subscription.Notes = dto.Notes
	}

	if err := s.repo.Update(subscription); err != nil {
		return nil, err
	}

	return subscription, nil
}

// Delete soft delete de um subscription
func (s *SubscriptionService) Delete(id string) error {
	return s.repo.Delete(id)
}

// Cancel cancela um subscription
func (s *SubscriptionService) Cancel(id string, dto *CancelSubscriptionDTO) (*models.PatientSubscription, error) {
	subscription, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	subscription.Status = models.SubscriptionCancelled
	subscription.CancelledAt = &now
	subscription.AutoRenew = false

	if dto.CancellationReason != nil {
		subscription.CancellationReason = dto.CancellationReason
	}

	if err := s.repo.Update(subscription); err != nil {
		return nil, err
	}

	return subscription, nil
}

// Renew renova um subscription
func (s *SubscriptionService) Renew(id string, dto *RenewSubscriptionDTO) (*models.PatientSubscription, error) {
	subscription, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	subscription.Status = models.SubscriptionActive
	subscription.RenewalCount++

	// Update end date if provided
	if dto.EndDate != nil {
		parsed, err := time.Parse("2006-01-02", *dto.EndDate)
		if err != nil {
			return nil, err
		}
		subscription.EndDate = &parsed
	}

	// Update next billing date if provided
	if dto.NextBillingDate != nil {
		parsed, err := time.Parse("2006-01-02", *dto.NextBillingDate)
		if err != nil {
			return nil, err
		}
		subscription.NextBillingDate = &parsed
	}

	if err := s.repo.Update(subscription); err != nil {
		return nil, err
	}

	return subscription, nil
}

// GetExpiringTrials busca subscriptions em trial que expiram em N dias
func (s *SubscriptionService) GetExpiringTrials(daysUntilExpiry int) ([]models.PatientSubscription, error) {
	return s.repo.GetExpiringTrials(daysUntilExpiry)
}

// GetUpcomingRenewals busca subscriptions com renovação próxima em N dias
func (s *SubscriptionService) GetUpcomingRenewals(daysUntilRenewal int) ([]models.PatientSubscription, error) {
	return s.repo.GetUpcomingRenewals(daysUntilRenewal)
}

// GetExpired busca subscriptions expirados
func (s *SubscriptionService) GetExpired() ([]models.PatientSubscription, error) {
	return s.repo.GetExpired()
}

// GetActiveByPatientID returns the active subscription for a patient with full method hierarchy
func (s *SubscriptionService) GetActiveByPatientID(patientID string) (*models.PatientSubscription, error) {
	return s.repo.GetActiveByPatientID(patientID)
}

// Suspend suspends an active subscription
func (s *SubscriptionService) Suspend(id string) (*models.PatientSubscription, error) {
	subscription, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	subscription.Status = models.SubscriptionSuspended

	if err := s.repo.Update(subscription); err != nil {
		return nil, err
	}

	return subscription, nil
}

// Activate activates a suspended or inactive subscription
func (s *SubscriptionService) Activate(id string) (*models.PatientSubscription, error) {
	subscription, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	subscription.Status = models.SubscriptionActive

	if err := s.repo.Update(subscription); err != nil {
		return nil, err
	}

	return subscription, nil
}
