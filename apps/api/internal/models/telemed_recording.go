package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Status da gravação cloud do Daily.
const (
	RecordingStatusPending  = "pending"  // token pediu start, ainda sem evento
	RecordingStatusStarted  = "started"  // recording.started recebido
	RecordingStatusFinished = "finished" // recording.ready-to-download recebido
	RecordingStatusError    = "error"    // recording.error
)

// Status da transcrição (Deepgram via Daily).
const (
	TranscriptStatusNone       = "none"        // nada iniciado
	TranscriptStatusInProgress = "in_progress" // transcript.started
	TranscriptStatusFinished   = "finished"    // transcript.ready-to-download + VTT baixado
	TranscriptStatusFailed     = "failed"      // transcript.error
)

// TelemedRecording — artefatos (gravação + transcrição) de uma teleconsulta.
//
// Uma linha por consulta (upsert por AppointmentID). Os eventos chegam de forma
// assíncrona via webhook do Daily (POST /webhooks/daily), mapeados por
// DailyRoomName. Decisão "referência + link sob demanda": o MP4 NÃO é baixado —
// guardamos RecordingID + metadados e geramos o link assinado quando o médico
// pede o download. A transcrição (WebVTT, poucos KB) é baixada, parseada em
// diálogo rotulado (médico×paciente) e persistida em TranscriptText.
//
// @Description Gravação + transcrição de uma teleconsulta (Daily.co + Deepgram)
type TelemedRecording struct {
	ID            uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	AppointmentID *uuid.UUID `gorm:"type:uuid;uniqueIndex;column:appointment_id" json:"appointmentId,omitempty"`
	PatientID     *uuid.UUID `gorm:"type:uuid;index" json:"patientId,omitempty"`
	// DailyRoomName — chave de join dos webhooks (payload.room_name → appointment).
	DailyRoomName string `gorm:"type:varchar(255);index" json:"dailyRoomName"`

	// Gravação (cloud MP4 no storage do Daily — referenciada, não baixada).
	RecordingID              *string    `gorm:"type:varchar(64);column:recording_id" json:"recordingId,omitempty"`
	RecordingStatus          string     `gorm:"type:varchar(16);not null;default:'pending'" json:"recordingStatus"`
	RecordingStartedAt       *time.Time `gorm:"type:timestamptz" json:"recordingStartedAt,omitempty"`
	RecordingReadyAt         *time.Time `gorm:"type:timestamptz" json:"recordingReadyAt,omitempty"`
	RecordingDurationSeconds *int       `gorm:"type:int" json:"recordingDurationSeconds,omitempty"`
	RecordingS3Key           *string    `gorm:"type:varchar(500);column:recording_s3_key" json:"-"`
	RecordingError           *string    `gorm:"type:text" json:"recordingError,omitempty"`

	// Transcrição (Deepgram nova-3 pt-BR, diarizada; WebVTT baixado + parseado).
	TranscriptID      *string    `gorm:"type:varchar(64);column:transcript_id" json:"transcriptId,omitempty"`
	TranscriptStatus  string     `gorm:"type:varchar(16);not null;default:'none'" json:"transcriptStatus"`
	TranscriptReadyAt *time.Time `gorm:"type:timestamptz" json:"transcriptReadyAt,omitempty"`
	TranscriptVTTPath *string    `gorm:"type:varchar(500);column:transcript_vtt_path" json:"-"` // relativo ao uploadsRoot
	TranscriptText    *string    `gorm:"type:text" json:"transcriptText,omitempty"`             // diálogo rotulado, legível
	TranscriptError   *string    `gorm:"type:text" json:"transcriptError,omitempty"`

	CreatedAt time.Time      `gorm:"not null;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (TelemedRecording) TableName() string { return "telemed_recordings" }

func (r *TelemedRecording) BeforeCreate(tx *gorm.DB) error {
	if r.ID == uuid.Nil {
		r.ID = uuid.Must(uuid.NewV7())
	}
	if r.RecordingStatus == "" {
		r.RecordingStatus = RecordingStatusPending
	}
	if r.TranscriptStatus == "" {
		r.TranscriptStatus = TranscriptStatusNone
	}
	return nil
}
