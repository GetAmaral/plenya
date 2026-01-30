package dto

import "github.com/plenya/api/internal/models"

// CreatePatientRequest representa o payload de criação de paciente
type CreatePatientRequest struct {
	Name         string        `json:"name" validate:"required,min=3,max=200"`
	CPF          string        `json:"cpf" validate:"required,len=11"`
	BirthDate    string        `json:"birthDate" validate:"required"` // formato: YYYY-MM-DD
	Gender       models.Gender `json:"gender" validate:"required,oneof=male female other"`
	Phone        *string       `json:"phone,omitempty" validate:"omitempty,min=10,max=20"`
	Address      *string       `json:"address,omitempty" validate:"omitempty,max=500"`
	Municipality *string       `json:"municipality,omitempty" validate:"omitempty,max=100"`
	State        *string       `json:"state,omitempty" validate:"omitempty,len=2"`
	MotherName   *string       `json:"motherName,omitempty" validate:"omitempty,min=3,max=200"`
	FatherName   *string       `json:"fatherName,omitempty" validate:"omitempty,min=3,max=200"`
	Height       *float64      `json:"height,omitempty" validate:"omitempty,gt=0"`
	Weight       *float64      `json:"weight,omitempty" validate:"omitempty,gt=0"`
}

// UpdatePatientRequest representa o payload de atualização de paciente
type UpdatePatientRequest struct {
	Name         *string        `json:"name,omitempty" validate:"omitempty,min=3,max=200"`
	BirthDate    *string        `json:"birthDate,omitempty"` // formato: YYYY-MM-DD
	Gender       *models.Gender `json:"gender,omitempty" validate:"omitempty,oneof=male female other"`
	Phone        *string        `json:"phone,omitempty" validate:"omitempty,min=10,max=20"`
	Address      *string        `json:"address,omitempty" validate:"omitempty,max=500"`
	Municipality *string        `json:"municipality,omitempty" validate:"omitempty,max=100"`
	State        *string        `json:"state,omitempty" validate:"omitempty,len=2"`
	MotherName   *string        `json:"motherName,omitempty" validate:"omitempty,min=3,max=200"`
	FatherName   *string        `json:"fatherName,omitempty" validate:"omitempty,min=3,max=200"`
	Height       *float64       `json:"height,omitempty" validate:"omitempty,gt=0"`
	Weight       *float64       `json:"weight,omitempty" validate:"omitempty,gt=0"`
}

// PatientResponse representa um paciente na resposta
type PatientResponse struct {
	ID           string        `json:"id"`
	UserID       string        `json:"userId"`
	Name         string        `json:"name"`
	BirthDate    string        `json:"birthDate"`
	Gender       models.Gender `json:"gender"`
	Phone        *string       `json:"phone,omitempty"`
	Address      *string       `json:"address,omitempty"`
	Municipality *string       `json:"municipality,omitempty"`
	State        *string       `json:"state,omitempty"`
	MotherName   *string       `json:"motherName,omitempty"`
	FatherName   *string       `json:"fatherName,omitempty"`
	Height       *float64      `json:"height,omitempty"`
	Weight       *float64      `json:"weight,omitempty"`
	CreatedAt    string        `json:"createdAt"`
	UpdatedAt    string        `json:"updatedAt"`
	// CPF não é retornado por segurança
}
