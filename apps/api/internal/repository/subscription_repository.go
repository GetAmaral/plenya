package repository

import (
	"github.com/plenya/api/internal/models"
	"gorm.io/gorm"
)

type SubscriptionRepository struct {
	db *gorm.DB
}

func NewSubscriptionRepository(db *gorm.DB) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

// GetByID busca um subscription por ID
func (r *SubscriptionRepository) GetByID(id string) (*models.PatientSubscription, error) {
	var subscription models.PatientSubscription
	if err := r.db.Preload("Patient").Preload("SubscriptionPlan").Preload("SubscriptionPlan.Method").Where("id = ?", id).First(&subscription).Error; err != nil {
		return nil, err
	}
	return &subscription, nil
}

// GetByPatientID busca todos os subscriptions de um paciente
func (r *SubscriptionRepository) GetByPatientID(patientID string) ([]models.PatientSubscription, error) {
	var subscriptions []models.PatientSubscription
	if err := r.db.Preload("SubscriptionPlan").Preload("SubscriptionPlan.Method").Where("patient_id = ?", patientID).Order("created_at DESC").Find(&subscriptions).Error; err != nil {
		return nil, err
	}
	return subscriptions, nil
}

// GetActive busca todos os subscriptions ativos
func (r *SubscriptionRepository) GetActive() ([]models.PatientSubscription, error) {
	var subscriptions []models.PatientSubscription
	if err := r.db.Preload("Patient").Preload("SubscriptionPlan").Preload("SubscriptionPlan.Method").Where("status = ?", models.SubscriptionActive).Order("created_at DESC").Find(&subscriptions).Error; err != nil {
		return nil, err
	}
	return subscriptions, nil
}

// GetAll busca todos os subscriptions (para admin dashboard)
func (r *SubscriptionRepository) GetAll() ([]models.PatientSubscription, error) {
	var subscriptions []models.PatientSubscription
	if err := r.db.Preload("Patient").Preload("SubscriptionPlan").Preload("SubscriptionPlan.Method").Order("created_at DESC").Find(&subscriptions).Error; err != nil {
		return nil, err
	}
	return subscriptions, nil
}

// Create cria um novo subscription
func (r *SubscriptionRepository) Create(subscription *models.PatientSubscription) error {
	return r.db.Create(subscription).Error
}

// Update atualiza um subscription
func (r *SubscriptionRepository) Update(subscription *models.PatientSubscription) error {
	return r.db.Save(subscription).Error
}

// Delete soft delete de um subscription
func (r *SubscriptionRepository) Delete(id string) error {
	return r.db.Where("id = ?", id).Delete(&models.PatientSubscription{}).Error
}

// GetExpiringTrials busca subscriptions em trial que expiram em N dias
func (r *SubscriptionRepository) GetExpiringTrials(daysUntilExpiry int) ([]models.PatientSubscription, error) {
	var subscriptions []models.PatientSubscription
	query := `
		SELECT * FROM patient_subscriptions
		WHERE status = ?
		AND trial_end_date IS NOT NULL
		AND trial_end_date <= CURRENT_DATE + INTERVAL '? days'
		AND trial_end_date >= CURRENT_DATE
		AND deleted_at IS NULL
	`
	if err := r.db.Raw(query, models.SubscriptionTrial, daysUntilExpiry).Scan(&subscriptions).Error; err != nil {
		return nil, err
	}
	return subscriptions, nil
}

// GetUpcomingRenewals busca subscriptions com renovação próxima em N dias
func (r *SubscriptionRepository) GetUpcomingRenewals(daysUntilRenewal int) ([]models.PatientSubscription, error) {
	var subscriptions []models.PatientSubscription
	query := `
		SELECT * FROM patient_subscriptions
		WHERE status = ?
		AND auto_renew = true
		AND next_billing_date IS NOT NULL
		AND next_billing_date <= CURRENT_DATE + INTERVAL '? days'
		AND next_billing_date >= CURRENT_DATE
		AND deleted_at IS NULL
	`
	if err := r.db.Raw(query, models.SubscriptionActive, daysUntilRenewal).Scan(&subscriptions).Error; err != nil {
		return nil, err
	}
	return subscriptions, nil
}

// GetExpired busca subscriptions expirados (end_date no passado e status ainda active)
func (r *SubscriptionRepository) GetExpired() ([]models.PatientSubscription, error) {
	var subscriptions []models.PatientSubscription
	if err := r.db.Where("status = ? AND end_date < CURRENT_DATE", models.SubscriptionActive).Find(&subscriptions).Error; err != nil {
		return nil, err
	}
	return subscriptions, nil
}

// GetActiveByPatientID returns the active subscription for a patient with full method hierarchy
// Considers both 'active' and 'trial' status as active subscriptions
func (r *SubscriptionRepository) GetActiveByPatientID(patientID string) (*models.PatientSubscription, error) {
	var subscription models.PatientSubscription
	err := r.db.
		Preload("SubscriptionPlan").
		Preload("SubscriptionPlan.Method").
		Preload("SubscriptionPlan.Method.Letters", func(db *gorm.DB) *gorm.DB {
			return db.Where("deleted_at IS NULL").Order("\"order\" ASC")
		}).
		Preload("SubscriptionPlan.Method.Letters.Pillars", func(db *gorm.DB) *gorm.DB {
			return db.Where("deleted_at IS NULL").Order("\"order\" ASC")
		}).
		Where("patient_id = ? AND status IN ?", patientID, []models.SubscriptionStatus{
			models.SubscriptionActive,
			models.SubscriptionTrial,
		}).
		First(&subscription).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // No active subscription is not an error
		}
		return nil, err
	}
	return &subscription, nil
}
