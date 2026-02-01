package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// UserRole define os tipos de usuários no sistema
type UserRole string

const (
	RoleAdmin   UserRole = "admin"
	RoleDoctor  UserRole = "doctor"
	RoleNurse   UserRole = "nurse"
	RolePatient UserRole = "patient"
)

// User representa um usuário no sistema
// @Description Usuário base do sistema (profissionais e pacientes)
type User struct {
	// ID único do usuário
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`

	// Email único do usuário
	// @example usuario@example.com
	Email string `gorm:"type:varchar(255);not null;unique;index" json:"email" validate:"required,email"`

	// Hash da senha (nunca exposto no JSON)
	PasswordHash string `gorm:"type:text;not null" json:"-"`

	// Role do usuário no sistema
	// @enum admin,doctor,nurse,patient
	Role UserRole `gorm:"type:varchar(20);not null;check:role IN ('admin','doctor','nurse','patient')" json:"role" validate:"required,oneof=admin doctor nurse patient"`

	// Se 2FA está habilitado
	TwoFactorEnabled bool `gorm:"type:boolean;default:false" json:"twoFactorEnabled"`

	// Secret do 2FA (nunca exposto no JSON)
	TwoFactorSecret string `gorm:"type:text" json:"-"`

	// ID do paciente selecionado (contexto atual de trabalho)
	// @example 550e8400-e29b-41d4-a716-446655440000
	SelectedPatientID *uuid.UUID `gorm:"type:uuid;index" json:"selectedPatientId,omitempty"`

	// Preferências do usuário (viewport do mindmap, etc.)
	// @example {"mindmapViewport":{"x":0,"y":0,"zoom":0.7}}
	Preferences datatypes.JSON `gorm:"type:jsonb;default:'{}'" json:"preferences,omitempty"`

	// Data de criação
	CreatedAt time.Time `gorm:"not null;autoCreateTime" json:"createdAt"`

	// Data de atualização
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime" json:"updatedAt"`

	// Data de deleção (soft delete)
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relações
	SelectedPatient *Patient `gorm:"foreignKey:SelectedPatientID;constraint:OnDelete:SET NULL" json:"selectedPatient,omitempty"`
}

// TableName especifica o nome da tabela
func (User) TableName() string {
	return "users"
}

// BeforeSave hook executado antes de salvar
func (u *User) BeforeSave(tx *gorm.DB) error {
	// Validar role
	validRoles := map[UserRole]bool{
		RoleAdmin:   true,
		RoleDoctor:  true,
		RoleNurse:   true,
		RolePatient: true,
	}

	if !validRoles[u.Role] {
		return gorm.ErrInvalidData
	}

	return nil
}
