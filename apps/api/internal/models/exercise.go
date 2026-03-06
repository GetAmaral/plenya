package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Exercise representa um exercício físico do catálogo
// @Description Exercício com nome, músculos-alvo, equipamentos e GIF
type Exercise struct {
	// ID único do exercício
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// ID externo (do exercises.json original)
	ExternalId string `gorm:"type:varchar(20);uniqueIndex;not null" json:"externalId"`

	// Nome em inglês
	Name string `gorm:"type:varchar(200);not null" json:"name"`

	// Nome em português
	NamePt string `gorm:"type:varchar(200)" json:"namePt"`

	// Partes do corpo (EN)
	BodyParts []string `gorm:"type:jsonb;serializer:json" json:"bodyParts"`

	// Partes do corpo (PT)
	BodyPartsPt []string `gorm:"type:jsonb;serializer:json" json:"bodyPartsPt"`

	// Músculos-alvo (EN)
	TargetMuscles []string `gorm:"type:jsonb;serializer:json" json:"targetMuscles"`

	// Músculos-alvo (PT)
	TargetMusclesPt []string `gorm:"type:jsonb;serializer:json" json:"targetMusclesPt"`

	// Equipamentos (EN)
	Equipments []string `gorm:"type:jsonb;serializer:json" json:"equipments"`

	// Equipamentos (PT)
	EquipmentsPt []string `gorm:"type:jsonb;serializer:json" json:"equipmentsPt"`

	// Músculos secundários (EN)
	SecondaryMuscles []string `gorm:"type:jsonb;serializer:json" json:"secondaryMuscles"`

	// Instruções (EN)
	Instructions []string `gorm:"type:jsonb;serializer:json" json:"instructions"`

	// Instruções (PT)
	InstructionsPt []string `gorm:"type:jsonb;serializer:json" json:"instructionsPt"`

	// URL do GIF local (/uploads/exercises/{id}.gif)
	GifUrl string `gorm:"type:text" json:"gifUrl"`

	// URL do GIF fallback (CDN original)
	GifUrlFallback string `gorm:"type:text" json:"gifUrlFallback"`

	// Ativo
	IsActive bool `gorm:"not null;default:true" json:"isActive"`

	// Timestamps
	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName especifica o nome da tabela
func (Exercise) TableName() string {
	return "exercises"
}

// BeforeCreate hook to generate UUID v7
func (e *Exercise) BeforeCreate(tx *gorm.DB) error {
	if e.ID == uuid.Nil {
		e.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
