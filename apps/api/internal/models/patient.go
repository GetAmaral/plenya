package models

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/crypto"
)

// Gender define os gêneros disponíveis (sexo biológico)
type Gender string

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
	GenderOther  Gender = "other"
)

// SocialGender define a identidade de gênero social
type SocialGender string

const (
	SocialGenderMale       SocialGender = "male"
	SocialGenderFemale     SocialGender = "female"
	SocialGenderNonBinary  SocialGender = "non_binary"
	SocialGenderTransMale  SocialGender = "trans_male"
	SocialGenderTransFemale SocialGender = "trans_female"
	SocialGenderOther      SocialGender = "other"
	SocialGenderPreferNotToSay SocialGender = "prefer_not_to_say"
)

// BloodType define os tipos sanguíneos
type BloodType string

const (
	BloodTypeAPositive  BloodType = "A+"
	BloodTypeANegative  BloodType = "A-"
	BloodTypeBPositive  BloodType = "B+"
	BloodTypeBNegative  BloodType = "B-"
	BloodTypeABPositive BloodType = "AB+"
	BloodTypeABNegative BloodType = "AB-"
	BloodTypeOPositive  BloodType = "O+"
	BloodTypeONegative  BloodType = "O-"
)

// MaritalStatus define os estados civis
type MaritalStatus string

const (
	MaritalStatusSingle   MaritalStatus = "single"
	MaritalStatusMarried  MaritalStatus = "married"
	MaritalStatusDivorced MaritalStatus = "divorced"
	MaritalStatusWidowed  MaritalStatus = "widowed"
	MaritalStatusOther    MaritalStatus = "other"
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

	// CPF do paciente (NULLABLE - pode ser preenchido depois)
	// @pattern ^\d{3}\.\d{3}\.\d{3}-\d{2}$
	// @example 123.456.789-00
	CPF *string `gorm:"type:text;unique" json:"-"`

	// Data de nascimento
	// @example 1990-01-01
	BirthDate time.Time `gorm:"type:date;not null" json:"birthDate" validate:"required"`

	// Gênero biológico do paciente
	// @enum male,female,other
	// @example male
	Gender Gender `gorm:"type:varchar(10);not null;check:gender IN ('male','female','other')" json:"gender" validate:"required,oneof=male female other"`

	// Identidade de gênero social (opcional)
	// @enum male,female,non_binary,trans_male,trans_female,other,prefer_not_to_say
	// @example female
	SocialGender *SocialGender `gorm:"type:varchar(20);check:social_gender IN ('male','female','non_binary','trans_male','trans_female','other','prefer_not_to_say')" json:"socialGender,omitempty" validate:"omitempty,oneof=male female non_binary trans_male trans_female other prefer_not_to_say"`

	// Indica se a paciente está na menopausa (apenas para gender=female)
	// @example true
	Menopause *bool `gorm:"type:boolean" json:"menopause,omitempty"`

	// Idade em anos (calculado automaticamente)
	// @example 35
	Age int `gorm:"type:int;not null;default:0" json:"age"`

	// Idade formatada (calculado automaticamente)
	// @example 35a6m15d
	AgeText string `gorm:"type:varchar(20);not null;default:''" json:"ageText"`

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

	// RG - Registro Geral (opcional, criptografado)
	// @example 12.345.678-9
	RG *string `gorm:"type:text" json:"-"`

	// Email pessoal (opcional)
	// @example joao@email.com
	Email *string `gorm:"type:varchar(200)" json:"email,omitempty" validate:"omitempty,email"`

	// Tipo sanguíneo (opcional)
	// @enum A+,A-,B+,B-,AB+,AB-,O+,O-
	// @example A+
	BloodType *BloodType `gorm:"type:varchar(3);check:blood_type IN ('A+','A-','B+','B-','AB+','AB-','O+','O-')" json:"bloodType,omitempty" validate:"omitempty,oneof=A+ A- B+ B- AB+ AB- O+ O-"`

	// Estado civil (opcional)
	// @enum single,married,divorced,widowed,other
	// @example married
	MaritalStatus *MaritalStatus `gorm:"type:varchar(10);check:marital_status IN ('single','married','divorced','widowed','other')" json:"maritalStatus,omitempty" validate:"omitempty,oneof=single married divorced widowed other"`

	// Profissão/ocupação (opcional)
	// @example Engenheiro
	Occupation *string `gorm:"type:varchar(200)" json:"occupation,omitempty"`

	// Nome do contato de emergência (opcional)
	// @example Maria da Silva
	EmergencyContact *string `gorm:"type:varchar(200)" json:"emergencyContact,omitempty"`

	// Telefone do contato de emergência (opcional)
	// @example (11) 99999-9999
	EmergencyPhone *string `gorm:"type:varchar(20)" json:"emergencyPhone,omitempty" validate:"omitempty,min=10"`

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

// CalculateAge calcula a idade em anos, meses e dias
func (p *Patient) CalculateAge() {
	now := time.Now()
	birthDate := p.BirthDate

	// Calcula anos completos
	years := now.Year() - birthDate.Year()

	// Ajusta se ainda não fez aniversário este ano
	if now.Month() < birthDate.Month() ||
		(now.Month() == birthDate.Month() && now.Day() < birthDate.Day()) {
		years--
	}

	// Calcula meses completos
	months := int(now.Month()) - int(birthDate.Month())
	if months < 0 {
		months += 12
	}
	if now.Day() < birthDate.Day() {
		months--
		if months < 0 {
			months += 12
		}
	}

	// Calcula dias completos
	days := now.Day() - birthDate.Day()
	if days < 0 {
		// Pega o número de dias do mês anterior
		prevMonth := now.AddDate(0, -1, 0)
		daysInPrevMonth := time.Date(prevMonth.Year(), prevMonth.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
		days += daysInPrevMonth
	}

	// Atualiza campos
	p.Age = years
	p.AgeText = fmt.Sprintf("%da%dm%dd", years, months, days)
}

// BeforeSave hook - criptografa o CPF e recalcula idade
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

	// Validar socialGender se fornecido
	if p.SocialGender != nil {
		validSocialGenders := map[SocialGender]bool{
			SocialGenderMale:           true,
			SocialGenderFemale:         true,
			SocialGenderNonBinary:      true,
			SocialGenderTransMale:      true,
			SocialGenderTransFemale:    true,
			SocialGenderOther:          true,
			SocialGenderPreferNotToSay: true,
		}

		if !validSocialGenders[*p.SocialGender] {
			return gorm.ErrInvalidData
		}
	}

	// Criptografar CPF se fornecido e não criptografado
	if p.CPF != nil && *p.CPF != "" && !isEncrypted(*p.CPF) {
		encrypted, err := crypto.EncryptWithDefaultKey(*p.CPF)
		if err != nil {
			return err
		}
		p.CPF = &encrypted
	}

	// Criptografar RG se fornecido e não criptografado
	if p.RG != nil && *p.RG != "" && !isEncrypted(*p.RG) {
		encrypted, err := crypto.EncryptWithDefaultKey(*p.RG)
		if err != nil {
			return err
		}
		p.RG = &encrypted
	}

	// Validar BloodType se fornecido
	if p.BloodType != nil {
		validBloodTypes := map[BloodType]bool{
			BloodTypeAPositive:  true,
			BloodTypeANegative:  true,
			BloodTypeBPositive:  true,
			BloodTypeBNegative:  true,
			BloodTypeABPositive: true,
			BloodTypeABNegative: true,
			BloodTypeOPositive:  true,
			BloodTypeONegative:  true,
		}

		if !validBloodTypes[*p.BloodType] {
			return gorm.ErrInvalidData
		}
	}

	// Validar MaritalStatus se fornecido
	if p.MaritalStatus != nil {
		validMaritalStatuses := map[MaritalStatus]bool{
			MaritalStatusSingle:   true,
			MaritalStatusMarried:  true,
			MaritalStatusDivorced: true,
			MaritalStatusWidowed:  true,
			MaritalStatusOther:    true,
		}

		if !validMaritalStatuses[*p.MaritalStatus] {
			return gorm.ErrInvalidData
		}
	}

	// Recalcula idade antes de salvar
	p.CalculateAge()

	return nil
}

// BeforeCreate hook to generate UUID v7
func (p *Patient) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.Must(uuid.NewV7())
	}
	// Calcula idade inicial
	p.CalculateAge()
	return nil
}

// AfterFind hook - descriptografa o CPF e RG após buscar do banco
func (p *Patient) AfterFind(tx *gorm.DB) error {
	// Descriptografar CPF se estiver criptografado
	if p.CPF != nil && *p.CPF != "" && isEncrypted(*p.CPF) {
		decrypted, err := crypto.DecryptWithDefaultKey(*p.CPF)
		if err != nil {
			// Se falhar ao descriptografar, manter criptografado
			// Isso evita erros se o formato mudar
			return nil
		}
		p.CPF = &decrypted
	}

	// Descriptografar RG se estiver criptografado
	if p.RG != nil && *p.RG != "" && isEncrypted(*p.RG) {
		decrypted, err := crypto.DecryptWithDefaultKey(*p.RG)
		if err != nil {
			// Se falhar ao descriptografar, manter criptografado
			return nil
		}
		p.RG = &decrypted
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
