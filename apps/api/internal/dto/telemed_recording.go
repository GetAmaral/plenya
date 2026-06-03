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

	// Nota clínica gerada por IA (AI scribe) — rascunho revisável, não assinável.
	GeneratedNoteStatus string         `json:"generatedNoteStatus"`
	GeneratedNoteFormat *string        `json:"generatedNoteFormat,omitempty"`
	GeneratedNoteModel  *string        `json:"generatedNoteModel,omitempty"`
	GeneratedNoteAt     *time.Time     `json:"generatedNoteAt,omitempty"`
	GeneratedNoteError  *string        `json:"generatedNoteError,omitempty"`
	GeneratedNote       *GeneratedNote `json:"generatedNote,omitempty"`

	UpdatedAt time.Time `json:"updatedAt"`
}

// GeneratedNote — nota clínica estruturada gerada por IA, já parseada em seções
// ordenadas prontas pra exibir/inserir. soapTarget liga cada seção ao campo da nota.
type GeneratedNote struct {
	Format        string                 `json:"format"`
	Sections      []GeneratedNoteSection `json:"sections"`
	ItensAmbiguos []string               `json:"itensAmbiguos,omitempty"`
	Papeis        map[string]string      `json:"papeis,omitempty"`
}

type GeneratedNoteSection struct {
	Chave      string `json:"chave"`
	Titulo     string `json:"titulo"`
	Texto      string `json:"texto"`
	SoapTarget string `json:"soapTarget"` // subjective|objective|assessment|plan
}

// TelemedRecordingDownloadResponse — link assinado temporário pro MP4 da gravação.
type TelemedRecordingDownloadResponse struct {
	DownloadURL string `json:"downloadUrl"`
}
