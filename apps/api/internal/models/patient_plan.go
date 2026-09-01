package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/pdfdoc"
)

// PatientPlan — a devolutiva de resultados do paciente ("o deck"): a leitura dos exames em voz de
// paciente, o que está bem, o que está se movendo e o que vai ser feito.
//
// É a fonte ÚNICA do conteúdo, com três saídas: a tela do portal, o PDF 16:9 (apresentar, mandar)
// e o PDF A4 paisagem (imprimir). As três saem do MESMO `Content`.
//
// Não se confunde com `CarePlanItem`: aquele é a lista estruturada de condutas por pilar AGIR, e
// continua sendo a fonte das condutas. O plano é a APRESENTAÇÃO delas junto com os achados.
type PatientPlanStatus string

const (
	PatientPlanDraft     PatientPlanStatus = "draft"
	PatientPlanPublished PatientPlanStatus = "published"
)

// @Description Plano de devolutiva do paciente
type PatientPlan struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	PatientID uuid.UUID `gorm:"type:uuid;not null;index" json:"patientId"`

	Title  string            `gorm:"type:varchar(300);not null" json:"title"`
	Status PatientPlanStatus `gorm:"type:varchar(20);not null;default:'draft'" json:"status"`

	// Version — a publicação. Cada publicação vira um documento novo no portal, e o número compõe
	// o `source_ref` que impede o mesmo plano de ser copiado duas vezes para o prontuário.
	Version int `gorm:"not null;default:1" json:"version"`

	// Content — os slides, na ordem. JSONB porque são heterogêneos e ninguém consulta um slide
	// isolado. O banco tem CHECK garantindo que é uma LISTA: gravar objeto aqui já derrubou uma
	// tela inteira antes (migration 00060).
	Content []pdfdoc.DeckSlide `gorm:"type:jsonb;serializer:json;not null;default:'[]'" json:"content"`

	SourceSnapshotID *uuid.UUID `gorm:"type:uuid" json:"sourceSnapshotId,omitempty"`
	AuthorUserID     uuid.UUID  `gorm:"type:uuid;not null" json:"authorUserId"`

	PublishedAt    *time.Time `gorm:"type:timestamptz" json:"publishedAt,omitempty"`
	Document16x9ID *uuid.UUID `gorm:"type:uuid;column:document_16x9_id" json:"document16x9Id,omitempty"`
	DocumentA4ID   *uuid.UUID `gorm:"type:uuid;column:document_a4_id" json:"documentA4Id,omitempty"`

	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Patient    Patient `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"-"`
	AuthorUser User    `gorm:"foreignKey:AuthorUserID;constraint:OnDelete:RESTRICT" json:"authorUser,omitempty"`
}

func (PatientPlan) TableName() string { return "patient_plans" }

func (p *PatientPlan) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
