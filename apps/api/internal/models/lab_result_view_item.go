package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// LabResultViewItem representa um exame dentro de uma view
// @Description Item de exame com ordem específica dentro de uma view
type LabResultViewItem struct {
	// @example 019c1a1e-0579-7f3b-a1bd-4767008e844c
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// @example 019c1a1e-0579-7f3b-a1bd-4767008e844c
	LabResultViewID uuid.UUID `gorm:"type:uuid;not null;index" json:"labResultViewId" validate:"required"`

	// @example 019c1a1e-0579-7f3b-a1bd-4767008e844c
	LabTestDefinitionID uuid.UUID `gorm:"type:uuid;not null;index" json:"labTestDefinitionId" validate:"required"`

	// @minimum 0
	// @maximum 9999
	Order int `gorm:"type:integer;not null;default:0;index" json:"order" validate:"gte=0,lte=9999"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	LabResultView     *LabResultView     `gorm:"foreignKey:LabResultViewID;constraint:OnDelete:CASCADE" json:"labResultView,omitempty"`
	LabTestDefinition *LabTestDefinition `gorm:"foreignKey:LabTestDefinitionID;constraint:OnDelete:RESTRICT" json:"labTestDefinition,omitempty"`
}

// BeforeCreate hook para gerar UUID v7
func (lrvi *LabResultViewItem) BeforeCreate(tx *gorm.DB) error {
	if lrvi.ID == uuid.Nil {
		lrvi.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// TableName especifica o nome da tabela
func (lrvi *LabResultViewItem) TableName() string {
	return "lab_result_view_items"
}
