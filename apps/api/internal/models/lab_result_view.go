package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// LabResultView representa uma view personalizada de exames laboratoriais
// @Description View customizada para ordenação de exames na tabela pivot
type LabResultView struct {
	// @example 019c1a1e-0579-7f3b-a1bd-4767008e844c
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// @minLength 2
	// @maxLength 200
	Name string `gorm:"type:varchar(200);not null" json:"name" validate:"required,min=2,max=200"`

	Description *string `gorm:"type:text" json:"description,omitempty"`

	// @example true
	IsActive bool `gorm:"type:boolean;not null;default:true;index" json:"isActive"`

	// @minimum 0
	// @maximum 9999
	DisplayOrder int `gorm:"type:integer;not null;default:0;index" json:"displayOrder" validate:"gte=0,lte=9999"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Items []LabResultViewItem `gorm:"foreignKey:LabResultViewID;constraint:OnDelete:CASCADE" json:"items,omitempty"`
}

// BeforeCreate hook para gerar UUID v7
func (lrv *LabResultView) BeforeCreate(tx *gorm.DB) error {
	if lrv.ID == uuid.Nil {
		lrv.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// TableName especifica o nome da tabela
func (lrv *LabResultView) TableName() string {
	return "lab_result_views"
}
