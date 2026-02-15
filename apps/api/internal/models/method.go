package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Method represents a clinical methodology/protocol system (e.g., AGIR)
// @Description Metodologia clínica (ex: AGIR, DASH, MIND)
type Method struct {
	// Primary Key
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Nome completo (ex: "AGIR - Protocolo de Saúde Integrativa")
	// @minLength 3
	// @maxLength 200
	// @example AGIR - Protocolo de Saúde Integrativa
	Name string `gorm:"type:varchar(200);not null;uniqueIndex:idx_method_name" json:"name" validate:"required,min=3,max=200"`

	// Nome curto/sigla (ex: "AGIR")
	// @minLength 2
	// @maxLength 20
	// @example AGIR
	ShortName string `gorm:"type:varchar(20);not null;uniqueIndex:idx_method_short_name" json:"shortName" validate:"required,min=2,max=20"`

	// Descrição da metodologia
	// @example Alimentação e Atividade Física, Gestão metabólica, Integração mente-corpo, Ritmo Circadiano
	Description *string `gorm:"type:text" json:"description,omitempty"`

	// Versão (ex: "1.0", "2024.1")
	// @maxLength 20
	Version *string `gorm:"type:varchar(20)" json:"version,omitempty" validate:"omitempty,max=20"`

	// Cor tema (para UI)
	// @example #3B82F6
	// @pattern ^#[0-9A-Fa-f]{6}$
	Color *string `gorm:"type:varchar(7)" json:"color,omitempty" validate:"omitempty,hexcolor"`

	// Ordem de exibição
	// @minimum 0
	// @maximum 9999
	Order int `gorm:"type:integer;not null;default:0;index:idx_method_order" json:"order" validate:"gte=0,lte=9999"`

	// Metodologia padrão do sistema (apenas uma pode ser default)
	// @example true
	IsDefault bool `gorm:"type:boolean;not null;default:false;index:idx_method_is_default" json:"isDefault"`

	// Relationships
	Letters []MethodLetter `gorm:"foreignKey:MethodID;constraint:OnDelete:CASCADE" json:"letters,omitempty"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Method) TableName() string {
	return "methods"
}

func (m *Method) BeforeCreate(tx *gorm.DB) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// BeforeSave ensures only one method can be default at a time
func (m *Method) BeforeSave(tx *gorm.DB) error {
	// If this method is being set as default, unset all other defaults
	if m.IsDefault {
		// Update all other methods to IsDefault=false
		if err := tx.Model(&Method{}).
			Where("id != ?", m.ID).
			Where("is_default = ?", true).
			Update("is_default", false).Error; err != nil {
			return err
		}
	}
	return nil
}
