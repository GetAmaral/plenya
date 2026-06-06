package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// ReceptionSettingsID é o UUID fixo da linha singleton de configuração do Atendimento IA.
var ReceptionSettingsID = uuid.MustParse("019e9301-0000-7000-a000-0000000a1a01")

// ReceptionSettings é a configuração GLOBAL (runtime) do Atendimento IA — controle único
// (singleton). O modo baseline vale para conversas sem flag; a janela de horário, quando
// habilitada e ativa, eleva o efetivo para Auto. Os tempos X/Y/Z controlam o job.
// Plano: docs/emr/plano-atendimento-ia-global.md.
//
// @Description Configuração global do Atendimento IA (recepcionista virtual)
type ReceptionSettings struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Modo baseline global: off | copilot | auto. Vale quando a conversa não tem flag e a
	// janela de horário não está elevando para Auto.
	// @enum off,copilot,auto
	BaselineMode string `gorm:"type:varchar(10);not null;default:'off'" json:"baselineMode"`

	// Liga/desliga a janela de horário do Auto.
	ScheduleEnabled bool `gorm:"not null;default:false" json:"scheduleEnabled"`

	// Janela semanal (estilo Google): { "mon": [{"start":"18:00","end":"08:00"}], ... }.
	// Faixas podem cruzar a meia-noite (start > end). Dias: mon,tue,wed,thu,fri,sat,sun.
	Schedule datatypes.JSON `gorm:"type:jsonb;not null;default:'{}'" json:"schedule" swaggertype:"object"`

	// X — debounce/atraso do Auto, em segundos (conta do último inbound; reseta a cada msg).
	DebounceSeconds int `gorm:"not null;default:30" json:"debounceSeconds"`

	// X2 — janela de preview do Auto, em segundos: tempo que o rascunho fica visível no
	// compositor (countdown) antes do envio automático. Separado do debounce.
	PreviewSeconds int `gorm:"not null;default:10" json:"previewSeconds"`

	// Y — minutos sem resposta no Copiloto antes de escalar para Auto.
	CopilotFallbackMinutes int `gorm:"not null;default:10" json:"copilotFallbackMinutes"`

	// Z — minutos sem resposta no Off (por herança) antes de alertar o admin.
	OffAlertMinutes int `gorm:"not null;default:20" json:"offAlertMinutes"`

	UpdatedBy *uuid.UUID `gorm:"type:uuid" json:"updatedBy,omitempty"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updatedAt"`
}

// TableName specifies the table name
func (ReceptionSettings) TableName() string { return "reception_settings" }

// BeforeCreate garante o UUID singleton fixo.
func (s *ReceptionSettings) BeforeCreate(tx *gorm.DB) error {
	if s.ID == uuid.Nil {
		s.ID = ReceptionSettingsID
	}
	return nil
}

// ScheduleTimeRange é uma faixa de horário "HH:MM". Start > End significa cruzar a meia-noite.
type ScheduleTimeRange struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

// WeeklySchedule mapeia dia da semana (mon..sun) → faixas de horário ativas do Auto.
type WeeklySchedule map[string][]ScheduleTimeRange

// ParsedSchedule desserializa o JSONB da janela semanal.
func (s *ReceptionSettings) ParsedSchedule() (WeeklySchedule, error) {
	ws := WeeklySchedule{}
	if len(s.Schedule) == 0 {
		return ws, nil
	}
	if err := json.Unmarshal(s.Schedule, &ws); err != nil {
		return nil, err
	}
	return ws, nil
}
