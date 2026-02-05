package dto

// CreateUserRequest representa o payload de criação de usuário
type CreateUserRequest struct {
	Name     string   `json:"name" validate:"required,min=3,max=200"`
	Email    string   `json:"email" validate:"required,email"`
	Password string   `json:"password" validate:"required,min=8,max=72"` // bcrypt max 72 bytes
	CPF      *string  `json:"cpf,omitempty"` // Formato: 123.456.789-00 ou sem formatação
	Roles    []string `json:"roles" validate:"required,min=1,dive,oneof=admin doctor nurse patient nutritionist psychologist physicalEducator secretary manager"`
}

// UpdateUserRequest representa o payload de atualização de usuário
type UpdateUserRequest struct {
	Name     *string  `json:"name,omitempty" validate:"omitempty,min=3,max=200"`
	Email    *string  `json:"email,omitempty" validate:"omitempty,email"`
	Password *string  `json:"password,omitempty" validate:"omitempty,min=8,max=72"`
	CPF      *string  `json:"cpf,omitempty"` // Formato: 123.456.789-00 ou sem formatação
	Roles    []string `json:"roles,omitempty" validate:"omitempty,min=1,dive,oneof=admin doctor nurse patient nutritionist psychologist physicalEducator secretary manager"`
}

// UserResponse representa um usuário na resposta
type UserResponse struct {
	ID                  string   `json:"id"`
	Name                string   `json:"name"`
	Email               string   `json:"email"`
	CPF                 *string  `json:"cpf,omitempty"`
	Roles               []string `json:"roles"`
	TwoFactorEnabled    bool     `json:"twoFactorEnabled"`
	ProfessionalPhone   *string  `json:"professionalPhone,omitempty"`
	ProfessionalAddress *string  `json:"professionalAddress,omitempty"`
	Specialty           *string  `json:"specialty,omitempty"`
	CRM                 *string  `json:"crm,omitempty"`
	CRMUF               *string  `json:"crmUF,omitempty"`
	RQE                 *string  `json:"rqe,omitempty"`
	CreatedAt           string   `json:"createdAt"`
	UpdatedAt           string   `json:"updatedAt"`
}

// UserListResponse representa a resposta paginada de usuários
type UserListResponse struct {
	Data  []UserResponse `json:"data"`
	Total int64          `json:"total"`
	Page  int            `json:"page"`
	Limit int            `json:"limit"`
}

// UpdateProfileRequest representa o payload de atualização do perfil
type UpdateProfileRequest struct {
	Name                *string `json:"name,omitempty" validate:"omitempty,min=3,max=200"`
	CPF                 *string `json:"cpf,omitempty"` // Formato: 123.456.789-00 ou sem formatação
	ProfessionalPhone   *string `json:"professionalPhone,omitempty" validate:"omitempty,max=20"`
	ProfessionalAddress *string `json:"professionalAddress,omitempty"`
	Specialty           *string `json:"specialty,omitempty" validate:"omitempty,max=100"`
	CRM                 *string `json:"crm,omitempty" validate:"omitempty,max=20"`
	CRMUF               *string `json:"crmUF,omitempty" validate:"omitempty,len=2"`
	RQE                 *string `json:"rqe,omitempty" validate:"omitempty,max=20"`
}
