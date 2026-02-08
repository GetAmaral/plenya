package dto

import "github.com/plenya/api/internal/models"

// CreatePatientRequest representa o payload de criação de paciente
type CreatePatientRequest struct {
	Name             string               `json:"name" validate:"required,min=3,max=200"`
	CPF              *string              `json:"cpf,omitempty" validate:"omitempty,len=11"`
	RG               *string              `json:"rg,omitempty" validate:"omitempty,max=20"`
	BirthDate        string               `json:"birthDate" validate:"required"` // formato: YYYY-MM-DD
	Gender           models.Gender        `json:"gender" validate:"required,oneof=male female other"`
	SocialGender     *models.SocialGender `json:"socialGender,omitempty" validate:"omitempty,oneof=male female non_binary trans_male trans_female other prefer_not_to_say"`
	Email            *string              `json:"email,omitempty" validate:"omitempty,email,max=200"`
	Phone            *string              `json:"phone,omitempty" validate:"omitempty,min=10,max=20"`
	Address          *string              `json:"address,omitempty" validate:"omitempty,max=500"`
	Municipality     *string              `json:"municipality,omitempty" validate:"omitempty,max=100"`
	State            *string              `json:"state,omitempty" validate:"omitempty,len=2"`
	MotherName       *string              `json:"motherName,omitempty" validate:"omitempty,min=3,max=200"`
	FatherName       *string              `json:"fatherName,omitempty" validate:"omitempty,min=3,max=200"`
	Height           *float64             `json:"height,omitempty" validate:"omitempty,gt=0"`
	Weight           *float64             `json:"weight,omitempty" validate:"omitempty,gt=0"`
	BloodType        *models.BloodType    `json:"bloodType,omitempty" validate:"omitempty,oneof=A+ A- B+ B- AB+ AB- O+ O-"`
	MaritalStatus    *models.MaritalStatus `json:"maritalStatus,omitempty" validate:"omitempty,oneof=single married divorced widowed other"`
	Occupation       *string              `json:"occupation,omitempty" validate:"omitempty,max=200"`
	EmergencyContact *string              `json:"emergencyContact,omitempty" validate:"omitempty,min=3,max=200"`
	EmergencyPhone   *string              `json:"emergencyPhone,omitempty" validate:"omitempty,min=10,max=20"`
}

// UpdatePatientRequest representa o payload de atualização de paciente
type UpdatePatientRequest struct {
	Name             *string               `json:"name,omitempty" validate:"omitempty,min=3,max=200"`
	CPF              *string               `json:"cpf,omitempty" validate:"omitempty,len=11"`
	RG               *string               `json:"rg,omitempty" validate:"omitempty,max=20"`
	BirthDate        *string               `json:"birthDate,omitempty"` // formato: YYYY-MM-DD
	Gender           *models.Gender        `json:"gender,omitempty" validate:"omitempty,oneof=male female other"`
	SocialGender     *models.SocialGender  `json:"socialGender,omitempty" validate:"omitempty,oneof=male female non_binary trans_male trans_female other prefer_not_to_say"`
	Email            *string               `json:"email,omitempty" validate:"omitempty,email,max=200"`
	Phone            *string               `json:"phone,omitempty" validate:"omitempty,min=10,max=20"`
	Address          *string               `json:"address,omitempty" validate:"omitempty,max=500"`
	Municipality     *string               `json:"municipality,omitempty" validate:"omitempty,max=100"`
	State            *string               `json:"state,omitempty" validate:"omitempty,len=2"`
	MotherName       *string               `json:"motherName,omitempty" validate:"omitempty,min=3,max=200"`
	FatherName       *string               `json:"fatherName,omitempty" validate:"omitempty,min=3,max=200"`
	Height           *float64              `json:"height,omitempty" validate:"omitempty,gt=0"`
	Weight           *float64              `json:"weight,omitempty" validate:"omitempty,gt=0"`
	BloodType        *models.BloodType     `json:"bloodType,omitempty" validate:"omitempty,oneof=A+ A- B+ B- AB+ AB- O+ O-"`
	MaritalStatus    *models.MaritalStatus `json:"maritalStatus,omitempty" validate:"omitempty,oneof=single married divorced widowed other"`
	Occupation       *string               `json:"occupation,omitempty" validate:"omitempty,max=200"`
	EmergencyContact *string               `json:"emergencyContact,omitempty" validate:"omitempty,min=3,max=200"`
	EmergencyPhone   *string               `json:"emergencyPhone,omitempty" validate:"omitempty,min=10,max=20"`
}

// PatientResponse representa um paciente na resposta
type PatientResponse struct {
	ID               string                `json:"id"`
	UserID           string                `json:"userId"`
	Name             string                `json:"name"`
	CPF              *string               `json:"cpf,omitempty"` // Retornado descriptografado para profissionais de saúde
	RG               *string               `json:"rg,omitempty"`  // Retornado descriptografado para profissionais de saúde
	BirthDate        string                `json:"birthDate"`
	Gender           models.Gender         `json:"gender"`
	SocialGender     *models.SocialGender  `json:"socialGender,omitempty"`
	Age              int                   `json:"age"`
	AgeText          string                `json:"ageText"`
	Email            *string               `json:"email,omitempty"`
	Phone            *string               `json:"phone,omitempty"`
	Address          *string               `json:"address,omitempty"`
	Municipality     *string               `json:"municipality,omitempty"`
	State            *string               `json:"state,omitempty"`
	MotherName       *string               `json:"motherName,omitempty"`
	FatherName       *string               `json:"fatherName,omitempty"`
	Height           *float64              `json:"height,omitempty"`
	Weight           *float64              `json:"weight,omitempty"`
	BloodType        *models.BloodType     `json:"bloodType,omitempty"`
	MaritalStatus    *models.MaritalStatus `json:"maritalStatus,omitempty"`
	Occupation       *string               `json:"occupation,omitempty"`
	EmergencyContact *string               `json:"emergencyContact,omitempty"`
	EmergencyPhone   *string               `json:"emergencyPhone,omitempty"`
	CreatedAt        string                `json:"createdAt"`
	UpdatedAt        string                `json:"updatedAt"`
}
