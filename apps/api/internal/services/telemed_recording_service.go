// Package services — TelemedRecordingService persiste e expõe os artefatos
// (gravação + transcrição) de uma teleconsulta. Os eventos chegam de forma
// assíncrona via webhook do Daily (POST /webhooks/daily), mapeados por
// daily_room_name → appointment.
//
// Decisão "referência + link sob demanda": o MP4 fica no storage do Daily —
// guardamos só metadados e geramos o link assinado on-demand. A transcrição
// (WebVTT, poucos KB) é baixada, parseada em diálogo rotulado (médico×paciente)
// e persistida.
package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

var (
	// ErrTelemedRecordingNotFound — não há linha de gravação/transcrição p/ a consulta.
	ErrTelemedRecordingNotFound = errors.New("telemed recording not found")
	// ErrTelemedNoRecording — consulta sem gravação pronta (nada pra baixar).
	ErrTelemedNoRecording = errors.New("telemed recording not available")
)

type TelemedRecordingService struct {
	db          *gorm.DB
	dailyCoSvc  *DailyCoService
	uploadsRoot string
}

func NewTelemedRecordingService(db *gorm.DB, dailyCoSvc *DailyCoService, uploadsRoot string) *TelemedRecordingService {
	return &TelemedRecordingService{db: db, dailyCoSvc: dailyCoSvc, uploadsRoot: uploadsRoot}
}

// upsertByRoom resolve (ou cria) a linha a partir do nome da sala Daily.
// Mapeia room_name → appointment pra preencher appointment_id + patient_id.
// Erro gorm.ErrRecordNotFound = sala não corresponde a nenhuma consulta (caller loga+ignora).
func (s *TelemedRecordingService) upsertByRoom(ctx context.Context, roomName string) (*models.TelemedRecording, error) {
	roomName = strings.TrimSpace(roomName)
	if roomName == "" {
		return nil, errors.New("telemed recording: empty room name")
	}
	var appt models.Appointment
	if err := s.db.WithContext(ctx).
		Select("id", "patient_id").
		Where("daily_room_name = ?", roomName).
		Order("created_at DESC").
		First(&appt).Error; err != nil {
		return nil, err
	}

	var rec models.TelemedRecording
	err := s.db.WithContext(ctx).Where("appointment_id = ?", appt.ID).First(&rec).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		patientID := appt.PatientID
		rec = models.TelemedRecording{
			AppointmentID: &appt.ID,
			PatientID:     &patientID,
			DailyRoomName: roomName,
		}
		if cerr := s.db.WithContext(ctx).Create(&rec).Error; cerr != nil {
			return nil, cerr
		}
		return &rec, nil
	}
	if err != nil {
		return nil, err
	}
	return &rec, nil
}

// MarkRecordingStarted — evento recording.started.
func (s *TelemedRecordingService) MarkRecordingStarted(ctx context.Context, roomName, recordingID string, startedAt time.Time) error {
	rec, err := s.upsertByRoom(ctx, roomName)
	if err != nil {
		return err
	}
	if recordingID != "" {
		rec.RecordingID = &recordingID
	}
	// Não regride status terminal.
	if rec.RecordingStatus == models.RecordingStatusPending || rec.RecordingStatus == "" {
		rec.RecordingStatus = models.RecordingStatusStarted
	}
	if rec.RecordingStartedAt == nil && !startedAt.IsZero() {
		rec.RecordingStartedAt = &startedAt
	}
	return s.db.WithContext(ctx).Save(rec).Error
}

// MarkRecordingReady — evento recording.ready-to-download. Guarda metadados; NÃO baixa o MP4.
func (s *TelemedRecordingService) MarkRecordingReady(ctx context.Context, roomName, recordingID string, durationSeconds int, s3Key string, readyAt time.Time) error {
	rec, err := s.upsertByRoom(ctx, roomName)
	if err != nil {
		return err
	}
	if recordingID != "" {
		rec.RecordingID = &recordingID
	}
	rec.RecordingStatus = models.RecordingStatusFinished
	if durationSeconds > 0 {
		d := durationSeconds
		rec.RecordingDurationSeconds = &d
	}
	if s3Key != "" {
		rec.RecordingS3Key = &s3Key
	}
	if readyAt.IsZero() {
		readyAt = time.Now().UTC()
	}
	rec.RecordingReadyAt = &readyAt
	return s.db.WithContext(ctx).Save(rec).Error
}

// MarkRecordingError — evento recording.error.
func (s *TelemedRecordingService) MarkRecordingError(ctx context.Context, roomName, errMsg string) error {
	rec, err := s.upsertByRoom(ctx, roomName)
	if err != nil {
		return err
	}
	rec.RecordingStatus = models.RecordingStatusError
	if errMsg != "" {
		rec.RecordingError = &errMsg
	}
	return s.db.WithContext(ctx).Save(rec).Error
}

// MarkTranscriptStarted — evento transcript.started.
func (s *TelemedRecordingService) MarkTranscriptStarted(ctx context.Context, roomName, transcriptID string) error {
	rec, err := s.upsertByRoom(ctx, roomName)
	if err != nil {
		return err
	}
	if transcriptID != "" {
		rec.TranscriptID = &transcriptID
	}
	if rec.TranscriptStatus == models.TranscriptStatusNone || rec.TranscriptStatus == "" {
		rec.TranscriptStatus = models.TranscriptStatusInProgress
	}
	return s.db.WithContext(ctx).Save(rec).Error
}

// MarkTranscriptReady — evento transcript.ready-to-download. Baixa o WebVTT,
// parseia em diálogo rotulado e persiste. Idempotente (não re-baixa se já finished).
func (s *TelemedRecordingService) MarkTranscriptReady(ctx context.Context, roomName, transcriptID string, readyAt time.Time) error {
	rec, err := s.upsertByRoom(ctx, roomName)
	if err != nil {
		return err
	}
	// Idempotência: já finalizado com o mesmo id → nada a fazer.
	if rec.TranscriptStatus == models.TranscriptStatusFinished && rec.TranscriptID != nil && *rec.TranscriptID == transcriptID {
		return nil
	}
	if transcriptID != "" {
		rec.TranscriptID = &transcriptID
	}
	if readyAt.IsZero() {
		readyAt = time.Now().UTC()
	}
	rec.TranscriptReadyAt = &readyAt

	// Baixa o VTT via access-link e parseia. Falha aqui não perde o evento:
	// marca finished mesmo assim e guarda o erro pra reprocesso manual.
	if s.dailyCoSvc != nil && transcriptID != "" {
		link, lerr := s.dailyCoSvc.GetTranscriptAccessLink(ctx, transcriptID)
		if lerr != nil {
			msg := lerr.Error()
			rec.TranscriptError = &msg
		} else if raw, ferr := s.dailyCoSvc.FetchSignedURL(ctx, link); ferr != nil {
			msg := ferr.Error()
			rec.TranscriptError = &msg
		} else {
			if path, perr := s.saveVTT(rec, transcriptID, raw); perr == nil {
				rec.TranscriptVTTPath = &path
			}
			dialog := parseDiarizedVTT(raw)
			if dialog != "" {
				rec.TranscriptText = &dialog
			}
			rec.TranscriptError = nil
		}
	}
	rec.TranscriptStatus = models.TranscriptStatusFinished
	return s.db.WithContext(ctx).Save(rec).Error
}

// MarkTranscriptError — evento transcript.error.
func (s *TelemedRecordingService) MarkTranscriptError(ctx context.Context, roomName, errMsg string) error {
	rec, err := s.upsertByRoom(ctx, roomName)
	if err != nil {
		return err
	}
	rec.TranscriptStatus = models.TranscriptStatusFailed
	if errMsg != "" {
		rec.TranscriptError = &errMsg
	}
	return s.db.WithContext(ctx).Save(rec).Error
}

// saveVTT grava o arquivo .vtt cru sob uploadsRoot/telemed-transcripts/{appointmentID}/{transcriptID}.vtt.
func (s *TelemedRecordingService) saveVTT(rec *models.TelemedRecording, transcriptID string, raw []byte) (string, error) {
	apptDir := "orphan"
	if rec.AppointmentID != nil {
		apptDir = rec.AppointmentID.String()
	}
	safeID := sanitizeFileToken(transcriptID)
	rel := filepath.Join("telemed-transcripts", apptDir, safeID+".vtt")
	full := filepath.Join(s.uploadsRoot, rel)
	if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
		return "", err
	}
	if err := os.WriteFile(full, raw, 0o644); err != nil {
		return "", err
	}
	return rel, nil
}

// GetByAppointment retorna o DTO de status. ErrTelemedRecordingNotFound se não há linha.
func (s *TelemedRecordingService) GetByAppointment(ctx context.Context, appointmentID uuid.UUID) (*dto.TelemedRecordingResponse, error) {
	var rec models.TelemedRecording
	if err := s.db.WithContext(ctx).Where("appointment_id = ?", appointmentID).First(&rec).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrTelemedRecordingNotFound
		}
		return nil, err
	}
	return s.toDTO(&rec), nil
}

// RecordingDownloadLink gera o link assinado de download do MP4 (sob demanda).
func (s *TelemedRecordingService) RecordingDownloadLink(ctx context.Context, appointmentID uuid.UUID) (string, error) {
	var rec models.TelemedRecording
	if err := s.db.WithContext(ctx).Where("appointment_id = ?", appointmentID).First(&rec).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", ErrTelemedRecordingNotFound
		}
		return "", err
	}
	if rec.RecordingID == nil || rec.RecordingStatus != models.RecordingStatusFinished {
		return "", ErrTelemedNoRecording
	}
	if s.dailyCoSvc == nil {
		return "", ErrDailyNotConfigured
	}
	return s.dailyCoSvc.GetRecordingAccessLink(ctx, *rec.RecordingID, 3600)
}

func (s *TelemedRecordingService) toDTO(rec *models.TelemedRecording) *dto.TelemedRecordingResponse {
	out := &dto.TelemedRecordingResponse{
		ID:                       rec.ID.String(),
		RecordingStatus:          rec.RecordingStatus,
		HasRecording:             rec.RecordingID != nil && rec.RecordingStatus == models.RecordingStatusFinished,
		RecordingReadyAt:         rec.RecordingReadyAt,
		RecordingDurationSeconds: rec.RecordingDurationSeconds,
		RecordingError:           rec.RecordingError,
		TranscriptStatus:         rec.TranscriptStatus,
		HasTranscript:            rec.TranscriptText != nil && *rec.TranscriptText != "",
		TranscriptReadyAt:        rec.TranscriptReadyAt,
		TranscriptText:           rec.TranscriptText,
		TranscriptError:          rec.TranscriptError,
		UpdatedAt:                rec.UpdatedAt,
	}
	if rec.AppointmentID != nil {
		id := rec.AppointmentID.String()
		out.AppointmentID = &id
	}
	if rec.PatientID != nil {
		id := rec.PatientID.String()
		out.PatientID = &id
	}
	return out
}

// ============================================================
// WebVTT parsing (diarizado) → diálogo legível
// ============================================================

var (
	vttVoiceRe         = regexp.MustCompile(`<v\s+([^>]+)>`)
	vttTagRe           = regexp.MustCompile(`</?[a-zA-Z][^>]*>`)
	vttTimestampRe     = regexp.MustCompile(`(\d{2}):(\d{2}):(\d{2})[.,]\d{3}\s*-->`)
	vttSpeakerPrefixRe = regexp.MustCompile(`^(?:Speaker|Falante)\s+(\d+)\s*[:\-]\s*`)
)

type vttTurn struct {
	speaker int
	start   string // mm:ss
	text    string
}

// parseDiarizedVTT converte o WebVTT diarizado do Daily/Deepgram num diálogo
// legível com rótulos estáveis "Falante N" e marca de tempo no início de cada
// turno. Funde cues consecutivos do mesmo falante. Robusto a ausência de
// diarização (tudo vira um único falante) e a tags VTT residuais.
func parseDiarizedVTT(raw []byte) string {
	text := strings.ReplaceAll(string(raw), "\r\n", "\n")
	blocks := strings.Split(text, "\n\n")

	speakerMap := map[string]int{}
	nextSpeaker := 1
	var turns []vttTurn

	for _, b := range blocks {
		lines := strings.Split(strings.TrimSpace(b), "\n")
		var start string
		var textLines []string
		for _, ln := range lines {
			ln = strings.TrimSpace(ln)
			if ln == "" || ln == "WEBVTT" || strings.HasPrefix(ln, "NOTE") {
				continue
			}
			if strings.Contains(ln, "-->") {
				if m := vttTimestampRe.FindStringSubmatch(ln); m != nil {
					start = mmssFromHMS(m[1], m[2], m[3])
				}
				continue
			}
			if isAllDigits(ln) { // linha de número do cue
				continue
			}
			textLines = append(textLines, ln)
		}
		if len(textLines) == 0 {
			continue
		}
		joined := strings.Join(textLines, " ")

		speakerLabel := ""
		if m := vttVoiceRe.FindStringSubmatch(joined); m != nil {
			speakerLabel = strings.TrimSpace(m[1])
		} else if m := vttSpeakerPrefixRe.FindStringSubmatch(joined); m != nil {
			speakerLabel = "Speaker " + m[1]
			joined = joined[len(m[0]):]
		}
		// Remove quaisquer tags VTT (<v ...>, </v>, <c>, etc.).
		joined = vttTagRe.ReplaceAllString(joined, "")
		joined = strings.TrimSpace(joined)
		if joined == "" {
			continue
		}

		sp := 0
		if speakerLabel != "" {
			if v, ok := speakerMap[speakerLabel]; ok {
				sp = v
			} else {
				speakerMap[speakerLabel] = nextSpeaker
				sp = nextSpeaker
				nextSpeaker++
			}
		}
		turns = append(turns, vttTurn{speaker: sp, start: start, text: joined})
	}

	var b strings.Builder
	lastSp := -1
	for _, t := range turns {
		if t.speaker == lastSp && b.Len() > 0 {
			b.WriteString(" " + t.text)
			continue
		}
		if b.Len() > 0 {
			b.WriteString("\n")
		}
		label := "Falante"
		if t.speaker > 0 {
			label = fmt.Sprintf("Falante %d", t.speaker)
		}
		if t.start != "" {
			b.WriteString(fmt.Sprintf("[%s] %s: %s", t.start, label, t.text))
		} else {
			b.WriteString(fmt.Sprintf("%s: %s", label, t.text))
		}
		lastSp = t.speaker
	}
	return strings.TrimSpace(b.String())
}

func isAllDigits(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// mmssFromHMS converte HH:MM:SS em MM:SS (some a hora quando 00; senão H:MM:SS).
func mmssFromHMS(hh, mm, ss string) string {
	if hh == "00" {
		return mm + ":" + ss
	}
	h := strings.TrimLeft(hh, "0")
	if h == "" {
		h = "0"
	}
	return h + ":" + mm + ":" + ss
}

// sanitizeFileToken mantém só [a-zA-Z0-9-_] pra compor nome de arquivo seguro.
func sanitizeFileToken(s string) string {
	var b strings.Builder
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9', r == '-', r == '_':
			b.WriteRune(r)
		}
	}
	out := b.String()
	if out == "" {
		out = "transcript"
	}
	if len(out) > 80 {
		out = out[:80]
	}
	return out
}
