package dto

// RegisterRequest representa o payload de registro de usuário
type RegisterRequest struct {
	Email    string   `json:"email" validate:"required,email"`
	Password string   `json:"password" validate:"required,min=8"`
	Roles    []string `json:"roles" validate:"required,min=1,dive,oneof=admin doctor nurse patient nutritionist psychologist physicalEducator secretary manager"`
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
	ID                string                 `json:"id"`
	Name              string                 `json:"name"`
	Email             string                 `json:"email"`
	CPF               *string                `json:"cpf,omitempty"`
	Roles             []string               `json:"roles"`
	TwoFactorEnabled  bool                   `json:"twoFactorEnabled"`
	SelectedPatientID *string                `json:"selectedPatientId,omitempty"`
	SelectedPatient   *PatientResponse       `json:"selectedPatient,omitempty"`
	Preferences       map[string]interface{} `json:"preferences,omitempty"`
	CreatedAt         string                 `json:"createdAt"`
}

// UpdateSelectedPatientRequest representa o payload para atualizar paciente selecionado
type UpdateSelectedPatientRequest struct {
	PatientID string `json:"patientId" validate:"required,uuid"`
}

// GoogleOAuthRequest representa o payload de login OAuth do Google
type GoogleOAuthRequest struct {
	IDToken string `json:"idToken" validate:"required"`
}

// AppleOAuthRequest representa o payload de login OAuth da Apple
type AppleOAuthRequest struct {
	IDToken string     `json:"idToken" validate:"required"`
	User    *AppleUser `json:"user,omitempty"`
}

// AppleUser representa dados do usuário retornados pela Apple (apenas no primeiro login)
type AppleUser struct {
	Name *AppleUserName `json:"name,omitempty"`
}

// AppleUserName representa o nome do usuário da Apple
type AppleUserName struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
