package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// AppointmentStatus define os status possíveis de uma consulta
type AppointmentStatus string

const (
	AppointmentScheduled  AppointmentStatus = "scheduled"   // Agendada
	AppointmentConfirmed  AppointmentStatus = "confirmed"   // Confirmada
	AppointmentCheckedIn  AppointmentStatus = "checked_in"  // Paciente chegou, aguardando atendimento
	AppointmentInProgress AppointmentStatus = "in_progress" // Em atendimento
	AppointmentCompleted  AppointmentStatus = "completed"   // Realizada
	AppointmentCancelled  AppointmentStatus = "cancelled"   // Cancelada
	AppointmentNoShow     AppointmentStatus = "no_show"     // Paciente não compareceu
)

// AppointmentType define os tipos de consulta.
//
// Plano Calendar V1 (Abril/2026): enum mudou de
//   routine|follow_up|urgent|emergency
// para
//   initial_assessment|follow_up|telemedicine|procedure|results_review
//
// Cada tipo tem duração default sugerida (UI controla isso, model permite
// qualquer DurationMinutes):
//   - InitialAssessment: 90min — primeira consulta com anamnese completa
//   - FollowUp:          30min — retorno
//   - Telemedicine:      30min — consulta remota (cria sala Daily.co)
//   - Procedure:         60min — procedimento clínico
//   - ResultsReview:     45min — revisão de exames
//
// Migration de dados existentes (em database.go): rows com type='routine',
// 'urgent' ou 'emergency' são reclassificadas como 'follow_up' (mapping
// conservador antes do CHECK constraint novo entrar).
type AppointmentType string

const (
	AppointmentInitialAssessment AppointmentType = "initial_assessment" // Avaliação Inicial (90min)
	AppointmentFollowUp          AppointmentType = "follow_up"          // Retorno (30min)
	AppointmentTelemedicine      AppointmentType = "telemedicine"       // Teleconsulta (30min, Daily.co)
	AppointmentProcedure         AppointmentType = "procedure"          // Procedimento (60min)
	AppointmentResultsReview     AppointmentType = "results_review"     // Revisão de Exames (45min)
)

// Appointment representa uma consulta/agendamento
// @Description Consulta médica com data, hora, médico e paciente
type Appointment struct {
	// ID único da consulta
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// ID do paciente
	PatientID uuid.UUID `gorm:"type:uuid;not null;index" json:"patientId"`

	// ID do médico
	DoctorID uuid.UUID `gorm:"type:uuid;not null;index" json:"doctorId"`

	// Data e hora da consulta
	ScheduledAt time.Time `gorm:"type:timestamptz;not null;index" json:"scheduledAt"`

	// Duração estimada em minutos
	DurationMinutes int `gorm:"type:int;not null;default:30" json:"durationMinutes"`

	// EndAt = ScheduledAt + DurationMinutes. Coluna materializada, mantida pelo
	// hook BeforeSave. Existe porque a EXCLUDE constraint anti-overlap precisa de
	// uma expressão IMMUTABLE — `tstzrange(scheduled_at, end_at)`. Calcular o fim
	// inline (`scheduled_at + interval`) NÃO é IMMUTABLE sobre timestamptz (a
	// soma com interval pode cruzar DST), então o Postgres rejeita em index
	// expression. Materializar o fim numa coluna resolve.
	EndAt time.Time `gorm:"type:timestamptz;not null;index" json:"endAt"`

	// Tipo de consulta. CHECK constraint garante apenas valores válidos no DB.
	// Validação de payload Go via tag `validate:"oneof=..."` (parsed pelo
	// validator no service/handler layer).
	Type AppointmentType `gorm:"type:varchar(20);not null;check:type IN ('initial_assessment','follow_up','telemedicine','procedure','results_review')" json:"type" validate:"oneof=initial_assessment follow_up telemedicine procedure results_review"`

	// Status da consulta. CHECK constraint garante valores válidos no DB.
	// A tag `check:` do GORM só cria a constraint na criação da tabela; alterar
	// o conjunto de valores numa tabela existente exige um bloco pré-AutoMigrate
	// em database.go que recria chk_appointments_status.
	Status AppointmentStatus `gorm:"type:varchar(20);not null;default:'scheduled';check:status IN ('scheduled','confirmed','checked_in','in_progress','completed','cancelled','no_show')" json:"status"`

	// Motivo da consulta
	Reason string `gorm:"type:text;not null" json:"reason" validate:"required"`

	// Observações do paciente
	PatientNotes *string `gorm:"type:text" json:"patientNotes,omitempty"`

	// Observações do médico (preenchido após a consulta)
	DoctorNotes *string `gorm:"type:text" json:"doctorNotes,omitempty"`

	// Diagnóstico (preenchido após a consulta)
	Diagnosis *string `gorm:"type:text" json:"diagnosis,omitempty"`

	// Link para anamnese (se criada durante a consulta)
	AnamnesisID *uuid.UUID `gorm:"type:uuid;index" json:"anamnesisId,omitempty"`

	// Link para item do Continuum (quando a consulta ancora um marco do programa).
	// AppointmentService propaga status:
	//  - Create com ContinuumItemID setado → item.Status = scheduled
	//  - Status muda pra completed → item.Status = completed
	// Cancel não muda o item automaticamente (equipe pode reagendar).
	ContinuumItemID *uuid.UUID `gorm:"type:uuid;index" json:"continuumItemId,omitempty"`

	// Data de confirmação
	ConfirmedAt *time.Time `gorm:"type:timestamptz" json:"confirmedAt,omitempty"`

	// Quando a recepção registrou a chegada do paciente (status -> checked_in).
	// Base, junto com StartedAt, para o cálculo do tempo de espera no balcão.
	CheckedInAt *time.Time `gorm:"type:timestamptz" json:"checkedInAt,omitempty"`

	// Quando o atendimento começou de fato (status -> in_progress).
	// StartedAt - CheckedInAt = tempo de espera do paciente.
	StartedAt *time.Time `gorm:"type:timestamptz" json:"startedAt,omitempty"`

	// Data de conclusão
	CompletedAt *time.Time `gorm:"type:timestamptz" json:"completedAt,omitempty"`

	// Data de cancelamento
	CancelledAt *time.Time `gorm:"type:timestamptz" json:"cancelledAt,omitempty"`

	// Motivo do cancelamento
	CancellationReason *string `gorm:"type:text" json:"cancellationReason,omitempty"`

	// ID do evento criado no Google Calendar dedicado do médico.
	// Populado async pelo GoogleCalendarService.CreateEvent após o appointment
	// ser persistido. Quando NULL, ainda não sincronizou (ou médico não tem
	// CalendarCredential ativa).
	ExternalCalendarEventID *string `gorm:"type:varchar(255);index" json:"externalCalendarEventId,omitempty"`

	// URL da sala Daily.co (somente Type=telemedicine).
	// Embedada inline em /appointments/[id] no frontend.
	DailyRoomURL *string `gorm:"type:varchar(255)" json:"dailyRoomUrl,omitempty"`

	// Nome único da sala Daily.co (necessário pra deletar a sala via API
	// quando appointment é cancelado).
	DailyRoomName *string `gorm:"type:varchar(255)" json:"dailyRoomName,omitempty"`

	// Quando a notificação de confirmação (email + WA template) foi enviada.
	// Usado pra evitar reenvio se o sync async retentar.
	ConfirmationSentAt *time.Time `gorm:"type:timestamptz" json:"confirmationSentAt,omitempty"`

	// Quando o reminder T-24h foi enviado (cron AppointmentReminderJob).
	// NULL = ainda não enviado; setado = job já processou.
	ReminderSentAt *time.Time `gorm:"type:timestamptz" json:"reminderSentAt,omitempty"`

	// Quando o push T-1h foi enviado pro paciente (cron AppointmentPushReminderJob).
	// NULL = ainda não enviado; setado = job já processou. Independente do WA T-24h.
	PushReminder1hSentAt *time.Time `gorm:"type:timestamptz" json:"pushReminder1hSentAt,omitempty"`

	// Quando o paciente clicou "Confirmar presença" no portal
	// (minha.plenyasaude.com.br/consultas/[id]). Diferente de ConfirmedAt,
	// que é a confirmação geral via Confirm() (também pode vir da IA).
	PatientConfirmedAt *time.Time `gorm:"type:timestamptz" json:"patientConfirmedAt,omitempty"`

	// Título computado para exibição no frontend (não persistido)
	DisplayTitle string `gorm:"-" json:"displayTitle"`

	// Data de criação
	CreatedAt time.Time `gorm:"not null;autoCreateTime" json:"createdAt"`

	// Data de atualização
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime" json:"updatedAt"`

	// Data de deleção (soft delete)
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relações
	Patient   Patient    `gorm:"foreignKey:PatientID;constraint:OnDelete:CASCADE" json:"patient,omitempty"`
	Doctor    User       `gorm:"foreignKey:DoctorID;constraint:OnDelete:RESTRICT" json:"doctor,omitempty"`
	Anamnesis *Anamnesis `gorm:"foreignKey:AnamnesisID" json:"anamnesis,omitempty"`
}

// TableName especifica o nome da tabela
func (Appointment) TableName() string {
	return "appointments"
}

// GetTitle retorna um título legível para a consulta
func (a *Appointment) GetTitle() string {
	typeLabels := map[AppointmentType]string{
		AppointmentInitialAssessment: "Avaliação Inicial",
		AppointmentFollowUp:          "Retorno",
		AppointmentTelemedicine:      "Teleconsulta",
		AppointmentProcedure:         "Procedimento",
		AppointmentResultsReview:     "Revisão de Exames",
	}
	label, ok := typeLabels[a.Type]
	if !ok {
		label = "Consulta"
	}
	return fmt.Sprintf("%s - %s", label, a.ScheduledAt.Format("02/01/2006 15:04"))
}

// AfterFind popula DisplayTitle após carregar do banco
func (a *Appointment) AfterFind(tx *gorm.DB) error {
	a.DisplayTitle = a.GetTitle()
	return nil
}

// BeforeCreate hook to generate UUID v7
func (a *Appointment) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.Must(uuid.NewV7())
	}
	return nil
}

// BeforeSave mantém EndAt = ScheduledAt + DurationMinutes em sincronia. Roda em
// Create e Update (Save). EndAt alimenta a EXCLUDE constraint anti-overlap, que
// precisa de tstzrange(scheduled_at, end_at) — ver comentário no campo EndAt.
func (a *Appointment) BeforeSave(tx *gorm.DB) error {
	dur := a.DurationMinutes
	if dur <= 0 {
		dur = 30
	}
	a.EndAt = a.ScheduledAt.Add(time.Duration(dur) * time.Minute)
	return nil
}
