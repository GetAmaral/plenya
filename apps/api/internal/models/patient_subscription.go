package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SubscriptionStatus define os status possíveis de uma assinatura
type SubscriptionStatus string

const (
	SubscriptionActive    SubscriptionStatus = "active"    // Ativo
	SubscriptionInactive  SubscriptionStatus = "inactive"  // Inativo (pausado)
	SubscriptionCancelled SubscriptionStatus = "cancelled" // Cancelado
	SubscriptionExpired   SubscriptionStatus = "expired"   // Expirado
	SubscriptionSuspended SubscriptionStatus = "suspended" // Suspenso (inadimplência)
	SubscriptionTrial     SubscriptionStatus = "trial"     // Em período trial
)

// PatientSubscription representa a assinatura concreta de um paciente a um plano
// @Description Patient subscription instance linked to a subscription plan
type PatientSubscription struct {
	// Primary Key
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id" example:"019d7f90-1234-7890-abcd-ef1234567890"`

	// Foreign Keys
	PatientID          uuid.UUID  `gorm:"type:uuid;not null;index:idx_patient_subscriptions_patient" json:"patientId" binding:"required" example:"019d7f90-1234-7890-abcd-ef1234567890"`
	SubscriptionPlanID uuid.UUID  `gorm:"type:uuid;not null;index:idx_patient_subscriptions_plan" json:"subscriptionPlanId" binding:"required" example:"019d7f90-1234-7890-abcd-ef1234567890"`

	// Plan snapshot - preserves plan data at subscription time (for history)
	PlanSnapshot string `gorm:"type:jsonb;not null" json:"planSnapshot" example:"{\"name\":\"Plano Premium\",\"price\":299.90,\"features\":{...}}"`

	// Status
	Status    SubscriptionStatus `gorm:"type:varchar(20);not null;index:idx_patient_subscriptions_status;check:status IN ('active','inactive','cancelled','expired','suspended','trial')" json:"status" binding:"required" example:"active"`
	AutoRenew bool               `gorm:"type:boolean;not null;default:true" json:"autoRenew" example:"true"`

	// Dates
	StartDate       time.Time  `gorm:"type:date;not null" json:"startDate" binding:"required" example:"2026-02-15"`
	EndDate         *time.Time `gorm:"type:date" json:"endDate,omitempty" example:"2027-02-15"`
	TrialEndDate    *time.Time `gorm:"type:date" json:"trialEndDate,omitempty" example:"2026-03-15"`
	NextBillingDate *time.Time `gorm:"type:date" json:"nextBillingDate,omitempty" example:"2026-03-15"`
	CancelledAt     *time.Time `gorm:"type:timestamp" json:"cancelledAt,omitempty" example:"2026-12-31T23:59:59Z"`

	// Customization (overrides plan defaults)
	DiscountPercent   float64 `gorm:"type:decimal(5,2);not null;default:0;check:discount_percent >= 0 AND discount_percent <= 100" json:"discountPercent" binding:"min=0,max=100" example:"10"`
	DiscountReason    *string `gorm:"type:text" json:"discountReason,omitempty" example:"Promoção de lançamento"`
	CustomPrice       *float64 `gorm:"type:decimal(10,2);check:custom_price >= 0" json:"customPrice,omitempty" example:"249.90"`
	CustomTrialDays   *int     `gorm:"type:integer;check:custom_trial_days >= 0" json:"customTrialDays,omitempty" example:"14"`

	// Management
	CancellationReason *string `gorm:"type:text" json:"cancellationReason,omitempty" example:"Paciente solicitou cancelamento por questões financeiras"`
	Notes              *string `gorm:"type:text" json:"notes,omitempty" example:"Paciente VIP - atendimento prioritário"`
	RenewalCount       int     `gorm:"type:integer;not null;default:0" json:"renewalCount" example:"3"`

	// Timestamps
	CreatedAt time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"createdAt" example:"2024-02-15T10:30:00Z"`
	UpdatedAt time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updatedAt" example:"2024-02-15T10:30:00Z"`
	DeletedAt gorm.DeletedAt `gorm:"index:idx_patient_subscriptions_deleted_at" json:"deletedAt,omitempty" swaggertype:"string" example:"2024-02-15T10:30:00Z"`

	// Relationships
	Patient          Patient          `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"patient,omitempty"`
	SubscriptionPlan SubscriptionPlan `gorm:"foreignKey:SubscriptionPlanID;constraint:OnDelete:RESTRICT" json:"subscriptionPlan,omitempty"`
}

func (PatientSubscription) TableName() string {
	return "patient_subscriptions"
}

// BeforeCreate hook - gera UUID v7 e cria snapshot do plano
func (ps *PatientSubscription) BeforeCreate(tx *gorm.DB) error {
	// Gerar UUID v7
	if ps.ID == uuid.Nil {
		ps.ID = uuid.Must(uuid.NewV7())
	}

	// Criar snapshot do plano se não fornecido
	if ps.PlanSnapshot == "" {
		var plan SubscriptionPlan
		if err := tx.Where("id = ?", ps.SubscriptionPlanID).First(&plan).Error; err != nil {
			return err
		}

		// Serializar plano para JSON (snapshot)
		snapshotData := map[string]interface{}{
			"id":              plan.ID,
			"name":            plan.Name,
			"description":     plan.Description,
			"features":        plan.Features,
			"price":           plan.Price,
			"currency":        plan.Currency,
			"billingCycle":    plan.BillingCycle,
			"methodId":        plan.MethodID,
			"trialPeriodDays": plan.TrialPeriodDays,
		}

		snapshotJSON, err := json.Marshal(snapshotData)
		if err != nil {
			return err
		}
		ps.PlanSnapshot = string(snapshotJSON)
	}

	// Validações
	return ps.validate()
}

// BeforeSave hook - validações antes de salvar
func (ps *PatientSubscription) BeforeSave(tx *gorm.DB) error {
	return ps.validate()
}

// validate valida os campos do subscription
func (ps *PatientSubscription) validate() error {
	// Validar status
	validStatuses := map[SubscriptionStatus]bool{
		SubscriptionActive:    true,
		SubscriptionInactive:  true,
		SubscriptionCancelled: true,
		SubscriptionExpired:   true,
		SubscriptionSuspended: true,
		SubscriptionTrial:     true,
	}
	if !validStatuses[ps.Status] {
		return gorm.ErrInvalidData
	}

	// Validar discount (0-100)
	if ps.DiscountPercent < 0 || ps.DiscountPercent > 100 {
		return gorm.ErrInvalidData
	}

	// Validar custom_price se fornecido
	if ps.CustomPrice != nil && *ps.CustomPrice < 0 {
		return gorm.ErrInvalidData
	}

	// Validar custom_trial_days se fornecido
	if ps.CustomTrialDays != nil && *ps.CustomTrialDays < 0 {
		return gorm.ErrInvalidData
	}

	return nil
}
