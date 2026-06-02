package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ClinicalNoteLayout define o layout da nota de evolução.
//   - soap: Subjetivo · Objetivo · Avaliação · Plano (ordem clássica)
//   - apso: Avaliação · Plano · Subjetivo · Objetivo (A/P no topo, leitura rápida)
//
// É só ordem de exibição no frontend — os 4 campos são os mesmos.
type ClinicalNoteLayout string

const (
	ClinicalNoteLayoutSOAP ClinicalNoteLayout = "soap"
	ClinicalNoteLayoutAPSO ClinicalNoteLayout = "apso"
)

// ClinicalNoteStatus define o ciclo de vida da nota.
//   - draft: editável pelo autor
//   - signed: imutável (correção só por adendo — guarda 20 anos / NGS2)
type ClinicalNoteStatus string

const (
	ClinicalNoteStatusDraft  ClinicalNoteStatus = "draft"
	ClinicalNoteStatusSigned ClinicalNoteStatus = "signed"
)

// ClinicalNote é a NOTA DE EVOLUÇÃO por consulta (SOAP/APSO).
//
// Distinta da Anamnese (que é o intake inicial one-shot, com Items ligados a
// ScoreItem): a ClinicalNote é a evolução de cada visita, ancorada ao
// Appointment via AppointmentID. Reaproveita o enum AnamnesisVisibility para
// governança por papel e o mesmo padrão rich-text (texto plano + HTML Tiptap,
// sanitizado com DOMPurify no display).
//
// Imutabilidade: após assinada (Status=signed), a nota é read-only no service.
// Correção é feita criando uma NOVA nota com AmendmentOfID apontando para a
// original (adendo), nunca editando a assinada.
//
// @Description Nota clínica de evolução por consulta (SOAP/APSO)
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

	// Layout de exibição (soap|apso)
	Layout ClinicalNoteLayout `gorm:"type:varchar(4);not null;default:'soap';check:layout IN ('soap','apso')" json:"layout"`

	// Seções SOAP — texto plano (busca/indexação) + HTML (exibição Tiptap)
	Subjective     *string `gorm:"type:text" json:"subjective,omitempty"`
	SubjectiveHtml *string `gorm:"type:text" json:"subjectiveHtml,omitempty"`
	Objective      *string `gorm:"type:text" json:"objective,omitempty"`
	ObjectiveHtml  *string `gorm:"type:text" json:"objectiveHtml,omitempty"`
	Assessment     *string `gorm:"type:text" json:"assessment,omitempty"`
	AssessmentHtml *string `gorm:"type:text" json:"assessmentHtml,omitempty"`
	Plan           *string `gorm:"type:text" json:"plan,omitempty"`
	PlanHtml       *string `gorm:"type:text" json:"planHtml,omitempty"`

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
