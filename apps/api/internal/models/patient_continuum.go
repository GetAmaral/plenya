package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// ContinuumStatus enumera os estados de uma inscrição de paciente em um Continuum.
type ContinuumStatus string

const (
	ContinuumActive    ContinuumStatus = "active"
	ContinuumPaused    ContinuumStatus = "paused"
	ContinuumCompleted ContinuumStatus = "completed"
	ContinuumCancelled ContinuumStatus = "cancelled"
)

// PatientContinuum é a inscrição de um paciente em um programa Continuum.
// Snapshot imutável do template é guardado em TemplateSnapshot (JSONB).
type PatientContinuum struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	PatientID  uuid.UUID `gorm:"type:uuid;not null;index" json:"patientId"`
	TemplateID uuid.UUID `gorm:"type:uuid;not null" json:"templateId"`

	// Snapshot imutável do template no momento da inscrição. Evita que mudanças
	// posteriores no template oficial alterem inscrições ativas.
	TemplateSnapshot datatypes.JSON `gorm:"type:jsonb;not null" json:"templateSnapshot"`

	Status ContinuumStatus `gorm:"type:varchar(20);not null;default:'active';check:status IN ('active','paused','completed','cancelled')" json:"status"`

	StartDate time.Time `gorm:"type:timestamp;not null" json:"startDate"`
	EndDate   time.Time `gorm:"type:timestamp;not null" json:"endDate"`

	// Médico coordenador do plano integrado (recebe alertas de atraso).
	CoordinatorDoctorID *uuid.UUID `gorm:"type:uuid;index" json:"coordinatorDoctorId,omitempty"`

	IntegratedPlanMarkdown  string     `gorm:"type:text" json:"integratedPlanMarkdown"`
	IntegratedPlanUpdatedAt *time.Time `gorm:"type:timestamp" json:"integratedPlanUpdatedAt,omitempty"`
	IntegratedPlanUpdatedBy *uuid.UUID `gorm:"type:uuid" json:"integratedPlanUpdatedBy,omitempty"`

	WhatsappGroupName       *string `gorm:"type:varchar(160)" json:"whatsappGroupName,omitempty"`
	WhatsappGroupInviteLink *string `gorm:"type:varchar(255)" json:"whatsappGroupInviteLink,omitempty"`

	Notes string `gorm:"type:text" json:"notes"`

	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Patient            *Patient                  `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"patient,omitempty"`
	CoordinatorDoctor  *User                     `gorm:"foreignKey:CoordinatorDoctorID" json:"coordinatorDoctor,omitempty"`
	Items              []PatientContinuumItem    `gorm:"foreignKey:ContinuumID" json:"items,omitempty"`
}

func (PatientContinuum) TableName() string { return "patient_continuums" }

func (c *PatientContinuum) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// ContinuumItemStatus enumera o ciclo de vida de cada item da timeline.
type ContinuumItemStatus string

const (
	ContinuumItemPending   ContinuumItemStatus = "pending"   // criado, ainda sem appointment ancorado
	ContinuumItemScheduled ContinuumItemStatus = "scheduled" // ancorado (appointment criado mas não realizado)
	ContinuumItemCompleted ContinuumItemStatus = "completed" // realizado
	ContinuumItemMissed    ContinuumItemStatus = "missed"    // passou do LateAfterDate sem ser concluído (cron seta)
	ContinuumItemCancelled ContinuumItemStatus = "cancelled" // cancelado explicitamente pela equipe
	ContinuumItemSkipped   ContinuumItemStatus = "skipped"   // pulado intencionalmente (paciente decidiu não fazer)
)

// PatientContinuumItem é uma instância de marco/entrega da inscrição. Gerada
// no Enroll a partir de cada ContinuumTemplateItem.
type PatientContinuumItem struct {
	ID          uuid.UUID         `gorm:"type:uuid;primaryKey" json:"id"`
	ContinuumID uuid.UUID         `gorm:"type:uuid;not null;index" json:"continuumId"`
	Type        ContinuumItemType `gorm:"type:varchar(30);not null;check:type IN ('appointment','box','reassessment','milestone','custom')" json:"type"`

	Specialty *ContinuumItemSpecialty `gorm:"type:varchar(30)" json:"specialty,omitempty"`

	Title       string `gorm:"type:varchar(160);not null" json:"title"`
	Description string `gorm:"type:text" json:"description"`

	WeekOffset    int       `gorm:"not null" json:"weekOffset"`
	ExpectedDate  time.Time `gorm:"type:timestamp;not null;index:idx_continuum_item_expected" json:"expectedDate"`
	LateAfterDate time.Time `gorm:"type:timestamp;not null" json:"lateAfterDate"`

	Status ContinuumItemStatus `gorm:"type:varchar(20);not null;default:'pending';check:status IN ('pending','scheduled','completed','missed','cancelled','skipped')" json:"status"`

	// Quando ancorado em appointment real (Type=appointment).
	AppointmentID *uuid.UUID `gorm:"type:uuid;index" json:"appointmentId,omitempty"`

	// Quando o item é Type=box, referencia o registro logístico do box.
	BoxID *uuid.UUID `gorm:"type:uuid" json:"boxId,omitempty"`

	CompletedAt *time.Time `gorm:"type:timestamp" json:"completedAt,omitempty"`

	// Referência genérica ao "evento" que satisfez o marco — pra reassessment
	// pode apontar pra patient_score_snapshot, pra milestone pode ser anamnesis.
	CompletedRefType *string    `gorm:"type:varchar(40)" json:"completedRefType,omitempty"`
	CompletedRefID   *uuid.UUID `gorm:"type:uuid" json:"completedRefId,omitempty"`

	Position int `gorm:"not null;default:0" json:"position"`

	CreatedAt time.Time `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime" json:"updatedAt"`
}

func (PatientContinuumItem) TableName() string { return "patient_continuum_items" }

func (i *PatientContinuumItem) BeforeCreate(tx *gorm.DB) error {
	if i.ID == uuid.Nil {
		i.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// PatientContinuumBoxStatus enumera o fluxo logístico de um Box Plenya.
type PatientContinuumBoxStatus string

const (
	BoxPlanned   PatientContinuumBoxStatus = "planned"
	BoxPreparing PatientContinuumBoxStatus = "preparing"
	BoxShipped   PatientContinuumBoxStatus = "shipped"
	BoxDelivered PatientContinuumBoxStatus = "delivered"
	BoxCancelled PatientContinuumBoxStatus = "cancelled"
)

// PatientContinuumBox é o registro logístico de um Box (1:1 com PatientContinuumItem
// quando Type=box). Snapshot do conteúdo + endereço pra que mudanças posteriores
// no paciente não alterem o histórico do box já enviado.
type PatientContinuumBox struct {
	ID              uuid.UUID                 `gorm:"type:uuid;primaryKey" json:"id"`
	ContinuumItemID uuid.UUID                 `gorm:"type:uuid;not null;index" json:"continuumItemId"`
	BoxTemplateID   *uuid.UUID                `gorm:"type:uuid" json:"boxTemplateId,omitempty"`
	Name            string                    `gorm:"type:varchar(120);not null" json:"name"`
	Contents        string                    `gorm:"type:text" json:"contents"`
	Status          PatientContinuumBoxStatus `gorm:"type:varchar(20);not null;default:'planned';check:status IN ('planned','preparing','shipped','delivered','cancelled')" json:"status"`

	PreparedAt   *time.Time `gorm:"type:timestamp" json:"preparedAt,omitempty"`
	ShippedAt    *time.Time `gorm:"type:timestamp" json:"shippedAt,omitempty"`
	DeliveredAt  *time.Time `gorm:"type:timestamp" json:"deliveredAt,omitempty"`
	TrackingCode *string    `gorm:"type:varchar(60)" json:"trackingCode,omitempty"`
	Carrier      *string    `gorm:"type:varchar(40)" json:"carrier,omitempty"` // "Correios", "Loggi"...

	AddressSnapshot string `gorm:"type:text" json:"addressSnapshot"`
	Notes           string `gorm:"type:text" json:"notes"`

	CreatedAt time.Time `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime" json:"updatedAt"`
}

func (PatientContinuumBox) TableName() string { return "patient_continuum_boxes" }

func (b *PatientContinuumBox) BeforeCreate(tx *gorm.DB) error {
	if b.ID == uuid.Nil {
		b.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// IntegratedPlanRevision é o histórico imutável de revisões do plano integrado
// (markdown colaborativo). Cada save no plano cria uma revisão.
type IntegratedPlanRevision struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	ContinuumID uuid.UUID `gorm:"type:uuid;not null;index" json:"continuumId"`
	Content     string    `gorm:"type:text;not null" json:"content"`
	UpdatedByID uuid.UUID `gorm:"type:uuid;not null" json:"updatedById"`
	CreatedAt   time.Time `gorm:"not null;autoCreateTime;index" json:"createdAt"`

	UpdatedBy *User `gorm:"foreignKey:UpdatedByID" json:"updatedBy,omitempty"`
}

func (IntegratedPlanRevision) TableName() string { return "integrated_plan_revisions" }

func (r *IntegratedPlanRevision) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
