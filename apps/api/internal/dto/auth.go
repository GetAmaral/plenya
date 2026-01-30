package dto

import "github.com/plenya/api/internal/models"

// RegisterRequest representa o payload de registro de usuário
type RegisterRequest struct {
	Email    string           `json:"email" validate:"required,email"`
	Password string           `json:"password" validate:"required,min=8"`
	Role     models.UserRole  `json:"role" validate:"required,oneof=admin doctor nurse patient"`
}

// LoginRequest representa o payload de login
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// RefreshRequest representa o payload de refresh token
type RefreshRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

// AuthResponse representa a resposta de autenticação
type AuthResponse struct {
	AccessToken  string      `json:"accessToken"`
	RefreshToken string      `json:"refreshToken"`
	User         UserDTO     `json:"user"`
}

// UserDTO representa um usuário na resposta
type UserDTO struct {
	ID                string           `json:"id"`
	Email             string           `json:"email"`
	Role              models.UserRole  `json:"role"`
	TwoFactorEnabled  bool             `json:"twoFactorEnabled"`
	SelectedPatientID *string          `json:"selectedPatientId,omitempty"`
	SelectedPatient   *PatientResponse `json:"selectedPatient,omitempty"`
	CreatedAt         string           `json:"createdAt"`
}

// UpdateSelectedPatientRequest representa o payload para atualizar paciente selecionado
type UpdateSelectedPatientRequest struct {
	PatientID string `json:"patientId" validate:"required,uuid"`
}
