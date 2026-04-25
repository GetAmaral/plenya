package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ContinuumBoxTemplate é o template de um Box Plenya (boas-vindas, mês 1,
// reposição...). Reutilizável entre Continuums via ContinuumTemplateItem.BoxTemplateID.
//
// Contents é markdown livre — descreve o que entra na caixa "padrão". A equipe
// de logística pode customizar o snapshot por paciente em PatientContinuumBox.
type ContinuumBoxTemplate struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name            string    `gorm:"type:varchar(120);not null" json:"name" validate:"required"`
	Description     string    `gorm:"type:text" json:"description"`
	Contents        string    `gorm:"type:text" json:"contents"` // markdown — itens "padrão"
	Notes           string    `gorm:"type:text" json:"notes"`    // instruções pra logística
	Status          string    `gorm:"type:varchar(20);not null;default:'active';check:status IN ('active','archived')" json:"status"`
	CreatedByUserID uuid.UUID `gorm:"type:uuid;not null" json:"createdByUserId"`

	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (ContinuumBoxTemplate) TableName() string { return "continuum_box_templates" }

func (b *ContinuumBoxTemplate) BeforeCreate(tx *gorm.DB) error {
	if b.ID == uuid.Nil {
		b.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
