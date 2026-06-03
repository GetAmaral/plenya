package dto

import "time"

// TelemedRecordingResponse — status de gravação + transcrição de uma teleconsulta,
// para o card do prontuário/consulta. O MP4 é referenciado (download sob demanda
// via endpoint próprio), nunca embutido aqui. A transcrição já vem parseada em
// diálogo legível (TranscriptText).
//
// @Description Status de gravação e transcrição de uma teleconsulta
type TelemedRecordingResponse struct {
	ID            string  `json:"id"`
	AppointmentID *string `json:"appointmentId,omitempty"`
	PatientID     *string `json:"patientId,omitempty"`

	RecordingStatus          string     `json:"recordingStatus"`
	HasRecording             bool       `json:"hasRecording"`
	RecordingReadyAt         *time.Time `json:"recordingReadyAt,omitempty"`
	RecordingDurationSeconds *int       `json:"recordingDurationSeconds,omitempty"`
	RecordingError           *string    `json:"recordingError,omitempty"`

	TranscriptStatus  string     `json:"transcriptStatus"`
	HasTranscript     bool       `json:"hasTranscript"`
	TranscriptReadyAt *time.Time `json:"transcriptReadyAt,omitempty"`
	TranscriptText    *string    `json:"transcriptText,omitempty"`
	TranscriptError   *string    `json:"transcriptError,omitempty"`

	UpdatedAt time.Time `json:"updatedAt"`
}

// TelemedRecordingDownloadResponse — link assinado temporário pro MP4 da gravação.
type TelemedRecordingDownloadResponse struct {
	DownloadURL string `json:"downloadUrl"`
}
