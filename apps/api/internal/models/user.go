package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/crypto"
)

// Role define os tipos de usuários no sistema
type Role string

const (
	RoleAdmin            Role = "admin"
	RoleDoctor           Role = "doctor"
	RoleNurse            Role = "nurse"
	RolePatient          Role = "patient"
	RoleNutritionist     Role = "nutritionist"
	RolePsychologist     Role = "psychologist"
	RolePhysicalEducator Role = "physicalEducator"
	RoleSecretary        Role = "secretary"
	RoleManager          Role = "manager"
)

// User representa um usuário no sistema
// @Description Usuário base do sistema (profissionais e pacientes)
type User struct {
	// ID único do usuário
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Nome completo do usuário
	// @minLength 3
	// @maxLength 200
	// @example Dr. João da Silva
	Name string `gorm:"type:varchar(200);not null" json:"name" validate:"required,min=3,max=200"`

	// Email único do usuário
	// @example usuario@example.com
	Email string `gorm:"type:varchar(255);not null;unique;index" json:"email" validate:"required,email"`

	// Hash da senha (NULLABLE para OAuth users)
	PasswordHash *string `gorm:"type:text" json:"-"`

	// Roles do usuário no sistema (array JSON)
	// @example ["doctor","nutritionist"]
	Roles datatypes.JSON `gorm:"type:jsonb;not null;default:'[\"patient\"]'" json:"roles"`

	// Se 2FA está habilitado
	TwoFactorEnabled bool `gorm:"type:boolean;default:false" json:"twoFactorEnabled"`

	// Secret do 2FA (nunca exposto no JSON)
	TwoFactorSecret string `gorm:"type:text" json:"-"`

	// OAuth provider (google, apple, null para tradicional)
	// @enum google,apple
	OAuthProvider *string `gorm:"type:varchar(20);index" json:"-"`

	// OAuth provider user ID (único por provider)
	OAuthProviderID *string `gorm:"type:varchar(255);index" json:"-"`

	// OAuth profile picture URL
	OAuthPictureURL *string `gorm:"type:text" json:"oauthPictureURL,omitempty"`

	// ID do paciente selecionado (contexto atual de trabalho)
	// @example 550e8400-e29b-41d4-a716-446655440000
	SelectedPatientID *uuid.UUID `gorm:"type:uuid;index" json:"selectedPatientId,omitempty"`

	// Preferências do usuário (viewport do mindmap, etc.)
	// @example {"mindmapViewport":{"x":0,"y":0,"zoom":0.7}}
	Preferences datatypes.JSON `gorm:"type:jsonb;default:'{}'" json:"preferences,omitempty"`

	// CPF do usuário (criptografado, LGPD)
	// @example 123.456.789-00
	CPF *string `gorm:"type:text;unique" json:"-"`

	// Dados profissionais (obrigatórios para role doctor)
	// Número do CRM
	// @example 123456
	CRM *string `gorm:"type:varchar(20);index" json:"crm,omitempty" validate:"omitempty,max=20"`

	// UF do CRM
	// @example SP
	CRMUF *string `gorm:"type:varchar(2)" json:"crmUF,omitempty" validate:"omitempty,len=2"`

	// Registro de Qualificação de Especialista
	// @example RQE-12345
	RQE *string `gorm:"type:varchar(20)" json:"rqe,omitempty"`

	// Especialidade médica
	// @example Cardiologia
	Specialty *string `gorm:"type:varchar(100)" json:"specialty,omitempty"`

	// Endereço profissional completo
	// @example Rua das Flores, 123 - Centro - São Paulo/SP - CEP 01234-567
	ProfessionalAddress *string `gorm:"type:text" json:"professionalAddress,omitempty"`

	// Telefone profissional
	// @example (11) 98765-4321
	ProfessionalPhone *string `gorm:"type:varchar(20)" json:"professionalPhone,omitempty"`

	// Certificado A1 (criptografado, nunca exposto no JSON)
	// Armazena o arquivo .pfx em Base64 criptografado com AES-256
	CertificatePFX *string `gorm:"type:text" json:"-"`

	// Senha do certificado A1 (criptografada, nunca exposta no JSON)
	CertificatePassword *string `gorm:"type:text" json:"-"`

	// Data de expiração do certificado A1
	// @example 2027-12-31T23:59:59Z
	CertificateExpiry *time.Time `gorm:"type:date" json:"certificateExpiry,omitempty"`

	// Número de série do certificado A1
	// @example 1234567890ABCDEF
	CertificateSerial *string `gorm:"type:varchar(100)" json:"certificateSerial,omitempty"`

	// CPF do titular do certificado (extraído do certificado ICP-Brasil)
	// @example 123.456.789-00
	CertificateCPF *string `gorm:"type:varchar(14)" json:"certificateCPF,omitempty"`

	// Nome do titular do certificado (extraído do certificado ICP-Brasil)
	// @example João da Silva
	CertificateName *string `gorm:"type:varchar(200)" json:"certificateName,omitempty"`

	// Se o certificado está ativo (desativa automaticamente se CPF do usuário mudar)
	// @example true
	CertificateActive bool `gorm:"type:boolean;default:false" json:"certificateActive"`

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

// BeforeCreate hook to generate UUID v7
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// BeforeSave hook - criptografa o CPF antes de salvar
func (u *User) BeforeSave(tx *gorm.DB) error {
	// Criptografar CPF se fornecido e não criptografado
	if u.CPF != nil && *u.CPF != "" && !isEncrypted(*u.CPF) {
		encrypted, err := crypto.EncryptWithDefaultKey(*u.CPF)
		if err != nil {
			return err
		}
		u.CPF = &encrypted
	}
	return nil
}

// AfterFind hook - descriptografa o CPF após buscar do banco
func (u *User) AfterFind(tx *gorm.DB) error {
	// Descriptografar CPF se estiver criptografado
	if u.CPF != nil && *u.CPF != "" && isEncrypted(*u.CPF) {
		decrypted, err := crypto.DecryptWithDefaultKey(*u.CPF)
		if err != nil {
			return err
		}
		u.CPF = &decrypted
	}
	return nil
}

// isEncrypted já definido em patient.go (função compartilhada no package)

// GetRoles retorna os roles do usuário como []string
func (u *User) GetRoles() []string {
	var roles []string
	if len(u.Roles) == 0 {
		return []string{}
	}
	if err := json.Unmarshal(u.Roles, &roles); err != nil {
		return []string{}
	}
	return roles
}

// IsGranted verifica se o usuário tem um role específico
func (u *User) IsGranted(role Role) bool {
	roles := u.GetRoles()
	roleStr := string(role)
	for _, r := range roles {
		if r == roleStr {
			return true
		}
	}
	return false
}

// SetRoles define os roles do usuário a partir de []string
func (u *User) SetRoles(roles []string) error {
	data, err := json.Marshal(roles)
	if err != nil {
		return err
	}
	u.Roles = datatypes.JSON(data)
	return nil
}
