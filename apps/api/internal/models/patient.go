package models

import (
	"encoding/base64"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/crypto"
)

// Gender define os gêneros disponíveis
type Gender string

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
	GenderOther  Gender = "other"
)

// Patient representa um paciente no sistema
// @Description Paciente com dados pessoais e médicos
type Patient struct {
	// ID único do paciente
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// ID do usuário associado
	UserID uuid.UUID `gorm:"type:uuid;not null;unique;index" json:"userId"`

	// Nome completo do paciente
	// @minLength 3
	// @maxLength 200
	// @example João da Silva
	Name string `gorm:"type:varchar(200);not null" json:"name" validate:"required,min=3,max=200"`

	// CPF do paciente (será criptografado antes de salvar)
	// @pattern ^\d{3}\.\d{3}\.\d{3}-\d{2}$
	// @example 123.456.789-00
	CPF string `gorm:"type:text;not null;unique" json:"-" validate:"required"`

	// Data de nascimento
	// @example 1990-01-01
	BirthDate time.Time `gorm:"type:date;not null" json:"birthDate" validate:"required"`

	// Gênero do paciente
	// @enum male,female,other
	// @example male
	Gender Gender `gorm:"type:varchar(10);not null;check:gender IN ('male','female','other')" json:"gender" validate:"required,oneof=male female other"`

	// Telefone de contato (opcional)
	// @example (11) 98765-4321
	Phone *string `gorm:"type:varchar(20)" json:"phone,omitempty" validate:"omitempty,min=10"`

	// Endereço residencial (opcional)
	Address *string `gorm:"type:varchar(500)" json:"address,omitempty"`

	// Município de residência (opcional)
	// @example São Paulo
	Municipality *string `gorm:"type:varchar(100);index" json:"municipality,omitempty"`

	// Estado/UF de residência (opcional)
	// @example SP
	State *string `gorm:"type:varchar(2);index" json:"state,omitempty" validate:"omitempty,len=2"`

	// Nome da mãe (opcional)
	// @example Maria da Silva
	MotherName *string `gorm:"type:varchar(200)" json:"motherName,omitempty"`

	// Nome do pai (opcional)
	// @example José da Silva
	FatherName *string `gorm:"type:varchar(200)" json:"fatherName,omitempty"`

	// Altura em centímetros (opcional)
	// @example 180.0
	Height *float64 `gorm:"type:decimal(5,2)" json:"height,omitempty"`

	// Peso em kilogramas (opcional)
	// @example 68.6
	Weight *float64 `gorm:"type:decimal(5,2)" json:"weight,omitempty"`

	// Data de criação
	CreatedAt time.Time `gorm:"not null;autoCreateTime" json:"createdAt"`

	// Data de atualização
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime" json:"updatedAt"`

	// Data de deleção (soft delete)
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relações
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
}

// TableName especifica o nome da tabela
func (Patient) TableName() string {
	return "patients"
}

// BeforeSave hook - criptografa o CPF antes de salvar
func (p *Patient) BeforeSave(tx *gorm.DB) error {
	// Validar gender
	validGenders := map[Gender]bool{
		GenderMale:   true,
		GenderFemale: true,
		GenderOther:  true,
	}

	if !validGenders[p.Gender] {
		return gorm.ErrInvalidData
	}

	// Criptografar CPF se ainda não estiver criptografado
	if p.CPF != "" && !isEncrypted(p.CPF) {
		encrypted, err := crypto.Encrypt(p.CPF)
		if err != nil {
			return err
		}
		p.CPF = encrypted
	}

	return nil
}

// BeforeCreate hook to generate UUID v7
func (p *Patient) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// AfterFind hook - descriptografa o CPF após buscar do banco
func (p *Patient) AfterFind(tx *gorm.DB) error {
	// Descriptografar CPF se estiver criptografado
	if p.CPF != "" && isEncrypted(p.CPF) {
		decrypted, err := crypto.Decrypt(p.CPF)
		if err != nil {
			// Se falhar ao descriptografar, manter criptografado
			// Isso evita erros se o formato mudar
			return nil
		}
		p.CPF = decrypted
	}

	return nil
}

// isEncrypted verifica se uma string está no formato base64 (indicando criptografia)
func isEncrypted(s string) bool {
	// CPF descriptografado tem 11 caracteres
	// Se tiver mais e for base64 válido, provavelmente está criptografado
	if len(s) <= 11 {
		return false
	}

	_, err := base64.StdEncoding.DecodeString(s)
	return err == nil
}
