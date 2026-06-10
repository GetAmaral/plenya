package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ClinicalNoteStatus define o ciclo de vida da nota.
//   - draft: editável pelo autor
//   - signed: imutável (correção só por adendo — guarda 20 anos / NGS2)
type ClinicalNoteStatus string

const (
	ClinicalNoteStatusDraft  ClinicalNoteStatus = "draft"
	ClinicalNoteStatusSigned ClinicalNoteStatus = "signed"
)

// ClinicalNote é a NOTA DE EVOLUÇÃO por consulta.
//
// Modelo simplificado (sem SOAP/APSO): dois campos de texto livre — "História
// clínica e evolução" (ClinicalHistory) e "Conduta" (Conduct). SOAP não é
// exigido por lei; a CFM Res. 1.638/2002 exige o CONTEÚDO (anamnese, exame,
// hipótese, conduta, evolução), não o formato — e parte do "objetivo" já é
// estruturada nos cards do workspace (vitais, CID-10, alergias, medicações).
//
// Distinta da Anamnese (intake com Items ligados a ScoreItem): a ClinicalNote é
// a evolução de cada visita, ancorada ao Appointment via AppointmentID.
// Reaproveita o enum AnamnesisVisibility e o padrão rich-text (texto plano +
// HTML Tiptap, sanitizado com DOMPurify no display).
//
// Imutabilidade: após assinada (Status=signed), a nota é read-only no service.
// Correção é feita criando uma NOVA nota com AmendmentOfID apontando para a
// original (adendo), nunca editando a assinada.
//
// @Description Nota clínica de evolução por consulta (História clínica + Conduta)
type ClinicalNote struct {
	// ID único da nota
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Consulta à qual a nota pertence (vínculo de "mesma visita"). Nullable
	// para permitir nota avulsa, mas no fluxo de consulta é sempre setado.
	AppointmentID *uuid.UUID `gorm:"type:uuid;index" json:"appointmentId,omitempty"`

	// Paciente da nota
	PatientID uuid.UUID `gorm:"type:uuid;not null;index" json:"patientId"`

	// Profissional autor da nota
	AuthorID uuid.UUID `gorm:"type:uuid;not null;index" json:"authorId"`

	// História clínica e evolução — narrativa da consulta (queixa, HDA, exame,
	// hipótese, evolução). Texto plano (busca/indexação) + HTML (exibição Tiptap).
	ClinicalHistory     *string `gorm:"type:text" json:"clinicalHistory,omitempty"`
	ClinicalHistoryHtml *string `gorm:"type:text" json:"clinicalHistoryHtml,omitempty"`

	// Conduta — plano (prescrição, exames, orientações, retorno).
	Conduct     *string `gorm:"type:text" json:"conduct,omitempty"`
	ConductHtml *string `gorm:"type:text" json:"conductHtml,omitempty"`

	// Status (draft|signed). Signed = imutável.
	Status ClinicalNoteStatus `gorm:"type:varchar(10);not null;default:'draft';check:status IN ('draft','signed')" json:"status"`

	// Quando a nota foi assinada (status -> signed)
	SignedAt *time.Time `gorm:"type:timestamptz" json:"signedAt,omitempty"`

	// Quando preenchido, esta nota é um ADENDO/correção da nota indicada
	// (que está assinada e permanece imutável).
	AmendmentOfID *uuid.UUID `gorm:"type:uuid;index" json:"amendmentOfId,omitempty"`

	// Visibilidade (reusa o enum da anamnese: all|medicalOnly|psychOnly|authorOnly)
	Visibility AnamnesisVisibility `gorm:"type:varchar(20);not null;default:'all';check:visibility IN ('all','medicalOnly','psychOnly','authorOnly')" json:"visibility"`

	// Título computado para exibição no frontend (não persistido)
	DisplayTitle string `gorm:"-" json:"displayTitle"`

	// Timestamps
	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relações
	Patient     Patient      `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"patient,omitempty"`
	Author      User         `gorm:"foreignKey:AuthorID;constraint:OnDelete:RESTRICT" json:"author,omitempty"`
	Appointment *Appointment `gorm:"foreignKey:AppointmentID" json:"appointment,omitempty"`
}

// TableName especifica o nome da tabela
func (ClinicalNote) TableName() string {
	return "clinical_notes"
}

// GetTitle retorna um título legível para a nota
func (n *ClinicalNote) GetTitle() string {
	label := "Evolução"
	if n.AmendmentOfID != nil {
		label = "Adendo"
	}
	return fmt.Sprintf("%s - %s", label, n.CreatedAt.Format("02/01/2006 15:04"))
}

// AfterFind popula DisplayTitle após carregar do banco
func (n *ClinicalNote) AfterFind(tx *gorm.DB) error {
	n.DisplayTitle = n.GetTitle()
	return nil
}

// BeforeCreate hook to generate UUID v7
func (n *ClinicalNote) BeforeCreate(tx *gorm.DB) error {
	if n.ID == uuid.Nil {
		n.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
