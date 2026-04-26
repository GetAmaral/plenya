package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// HealthCheckIn é o "como você está hoje?" diário do paciente. Card único na
// home do app paciente, ~10s de UX. Vira timeline pro profissional ler no
// continuum (gráfico simples + ícones de alerta quando há picos).
//
// Sem unique constraint (paciente pode logar 2x/dia se quiser, ex.: matutino
// + noturno) — o web agrega por dia.
//
// @Description Check-in diário de bem-estar do paciente
type HealthCheckIn struct {
	// ID único
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Paciente
	PatientID uuid.UUID `gorm:"type:uuid;not null;index" json:"patientId"`

	// Energia 1-5 (1 = exausto, 5 = ótimo)
	Energy int `gorm:"type:int;not null;check:energy >= 1 AND energy <= 5" json:"energy" validate:"required,min=1,max=5"`

	// Dor 0-5 (0 = nenhuma, 5 = severa)
	Pain int `gorm:"type:int;not null;check:pain >= 0 AND pain <= 5" json:"pain" validate:"min=0,max=5"`

	// Localização da dor (free text quando Pain > 0)
	PainLocation *string `gorm:"type:varchar(200)" json:"painLocation,omitempty"`

	// Humor 1-5 (1 = péssimo, 5 = ótimo)
	Mood int `gorm:"type:int;not null;check:mood >= 1 AND mood <= 5" json:"mood" validate:"required,min=1,max=5"`

	// Horas dormidas (0-24, com decimal)
	SleepHours float64 `gorm:"type:numeric(4,1);not null;check:sleep_hours >= 0 AND sleep_hours <= 24" json:"sleepHours" validate:"min=0,max=24"`

	// Qualidade do sono 1-5
	SleepQuality int `gorm:"type:int;not null;check:sleep_quality >= 1 AND sleep_quality <= 5" json:"sleepQuality" validate:"required,min=1,max=5"`

	// Estresse 1-5
	Stress int `gorm:"type:int;not null;check:stress >= 1 AND stress <= 5" json:"stress" validate:"required,min=1,max=5"`

	// Notas opcionais
	Notes *string `gorm:"type:text" json:"notes,omitempty"`

	// Timestamps
	CreatedAt time.Time      `gorm:"not null;autoCreateTime;index" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relação
	Patient Patient `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"-"`
}

func (HealthCheckIn) TableName() string { return "health_check_ins" }

func (h *HealthCheckIn) BeforeCreate(tx *gorm.DB) error {
	if h.ID == uuid.Nil {
		h.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}
