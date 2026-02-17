package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BillingCycle represents the billing frequency
type BillingCycle string

const (
	BillingCycleMonthly   BillingCycle = "monthly"
	BillingCycleQuarterly BillingCycle = "quarterly"
	BillingCycleYearly    BillingCycle = "yearly"
	BillingCycleOneTime   BillingCycle = "one_time"
)

// SubscriptionPlan represents a subscription plan definition (catalog)
// @Description Subscription plan catalog defining available plans for patients
type SubscriptionPlan struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id" example:"019d7f90-1234-7890-abcd-ef1234567890"`
	Name        string    `gorm:"type:varchar(255);not null;index:idx_plan_name" json:"name" binding:"required" example:"Plano Premium"`
	Description string    `gorm:"type:text" json:"description" example:"Acompanhamento completo com todas as funcionalidades"`

	// Features and pricing
	Features    string  `gorm:"type:jsonb" json:"features" example:"{\"consultas\": 12, \"exames\": \"completos\", \"whatsapp\": true}"`
	Price       float64 `gorm:"type:decimal(10,2);not null;check:price >= 0" json:"price" binding:"required,min=0" example:"299.90"`
	Currency    string  `gorm:"type:varchar(3);not null;default:'BRL'" json:"currency" example:"BRL"`
	BillingCycle BillingCycle `gorm:"type:varchar(20);not null;check:billing_cycle IN ('monthly','quarterly','yearly','one_time')" json:"billingCycle" binding:"required" example:"monthly"`

	// Method association
	MethodID *uuid.UUID `gorm:"type:uuid;index:idx_plan_method" json:"methodId" example:"019d7f90-1234-7890-abcd-ef1234567890"`
	Method   *Method    `gorm:"foreignKey:MethodID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"method,omitempty"`

	// Configuration
	IsActive         bool `gorm:"type:boolean;not null;default:true;index:idx_plan_is_active" json:"isActive" example:"true"`
	TrialPeriodDays  int  `gorm:"type:integer;not null;default:0;check:trial_period_days >= 0" json:"trialPeriodDays" example:"7"`
	Order            int  `gorm:"type:integer;not null;default:0;index:idx_plan_order" json:"order" example:"1"`

	// Timestamps
	CreatedAt time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"createdAt" example:"2024-02-15T10:30:00Z"`
	UpdatedAt time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updatedAt" example:"2024-02-15T10:30:00Z"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty" swaggertype:"string" example:"2024-02-15T10:30:00Z"`
}

// TableName specifies the table name for SubscriptionPlan
func (SubscriptionPlan) TableName() string {
	return "subscription_plans"
}

// BeforeCreate hook generates UUID v7 if not set
func (sp *SubscriptionPlan) BeforeCreate(tx *gorm.DB) error {
	if sp.ID == uuid.Nil {
		sp.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
