package models

import (
	"crypto/rand"
	"errors"
	"math/big"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AnonymousScoreSession representa uma sessão anônima do Escore Plenya Light.
// É preenchida pelo paciente no site público sem autenticação. Pode opcionalmente ser
// vinculada (claimed) a um Patient via magic link.
//
// @Description Sessão anônima do Escore Plenya Light (formulário público)
type AnonymousScoreSession struct {
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Código público alfanumérico (12 chars, case-sensitive) — usado em URLs do site
	// @example A7f9X2k3LpQv
	PublicCode string `gorm:"type:varchar(12);uniqueIndex;not null" json:"publicCode"`

	// Idade declarada pelo paciente
	// @minimum 0
	// @maximum 150
	// @example 47
	Age int `gorm:"type:integer;not null;check:age >= 0 AND age <= 150" json:"age" validate:"required,gte=0,lte=150"`

	// Gênero biológico declarado (male, female, other)
	// @enum male,female,other
	// @example female
	Gender string `gorm:"type:varchar(20);not null;check:gender IN ('male','female','other')" json:"gender" validate:"required,oneof=male female other"`

	// Status de pós-menopausa (apenas relevante para gênero female)
	// @example true
	PostMenopause *bool `gorm:"type:boolean" json:"postMenopause,omitempty"`

	// Altura declarada em cm (opcional)
	// @example 170
	Height *float64 `gorm:"type:double precision" json:"height,omitempty"`

	// Peso declarado em kg (opcional)
	// @example 72.5
	Weight *float64 `gorm:"type:double precision" json:"weight,omitempty"`

	// Vinculação opcional a um Patient (preenchido após claim via magic link)
	// @example 550e8400-e29b-41d4-a716-446655440000
	ClaimedByPatientID *uuid.UUID `gorm:"type:uuid;index:idx_anon_session_patient" json:"claimedByPatientId,omitempty"`

	// Timestamp do claim
	ClaimedAt *time.Time `gorm:"type:timestamp" json:"claimedAt,omitempty"`

	// Email informado durante o claim (usado para login posterior via magic link)
	// @example paciente@exemplo.com
	Email *string `gorm:"type:varchar(255);index:idx_anon_session_email" json:"email,omitempty"`

	// Data de expiração — sessões não-claimed são limpas após 90 dias.
	// Sessões claimed têm ExpiresAt = NULL (não expiram).
	ExpiresAt *time.Time `gorm:"type:timestamp;index:idx_anon_session_expires" json:"expiresAt,omitempty"`

	// Timestamps
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	ClaimedByPatient *Patient                   `gorm:"foreignKey:ClaimedByPatientID;constraint:OnDelete:SET NULL" json:"claimedByPatient,omitempty"`
	Items            []AnonymousScoreItem       `gorm:"foreignKey:SessionID;constraint:OnDelete:CASCADE" json:"items,omitempty"`
	Snapshot         *AnonymousScoreSnapshot    `gorm:"foreignKey:SessionID;constraint:OnDelete:CASCADE" json:"snapshot,omitempty"`
}

// TableName specifies the table name
func (AnonymousScoreSession) TableName() string {
	return "anonymous_score_sessions"
}

// BeforeCreate hook generates UUID v7, PublicCode (com retry em colisão) e ExpiresAt (90d)
func (s *AnonymousScoreSession) BeforeCreate(tx *gorm.DB) error {
	if s.ID == uuid.Nil {
		s.ID = uuid.Must(uuid.NewV7())
	}
	if s.PublicCode == "" {
		// Até 5 tentativas — entropy 62^12 torna colisão estatisticamente irrelevante,
		// mas o uniqueIndex protege contra o caso degenerado.
		for attempt := 0; attempt < 5; attempt++ {
			code, err := generateAnonSessionCode(12)
			if err != nil {
				return err
			}
			var count int64
			if err := tx.Model(&AnonymousScoreSession{}).Where("public_code = ?", code).Count(&count).Error; err != nil {
				return err
			}
			if count == 0 {
				s.PublicCode = code
				break
			}
		}
		if s.PublicCode == "" {
			return errAnonSessionCodeCollision
		}
	}
	if s.ExpiresAt == nil && s.ClaimedByPatientID == nil {
		t := time.Now().Add(90 * 24 * time.Hour)
		s.ExpiresAt = &t
	}
	return nil
}

var errAnonSessionCodeCollision = errors.New("failed to generate unique public code after 5 attempts")

// generateAnonSessionCode gera um código alfanumérico aleatório (case-sensitive)
func generateAnonSessionCode(length int) (string, error) {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, length)
	for i := range result {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[n.Int64()]
	}
	return string(result), nil
}

// ToPatientForApplies retorna um pseudo-Patient compatível com ScoreItem.AppliesToPatient
// (reaproveita a lógica de filtragem demográfica existente sem precisar de Patient real).
func (s *AnonymousScoreSession) ToPatientForApplies() *Patient {
	return &Patient{
		Age:       s.Age,
		Gender:    Gender(s.Gender),
		Menopause: s.PostMenopause,
	}
}
