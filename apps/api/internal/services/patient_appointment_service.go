package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// PatientAppointmentService expõe ações de consulta no escopo do paciente logado.
//
// Diferente de AppointmentService (acessa todas as consultas, escopo staff),
// este sempre filtra por patientID e bloqueia ações destrutivas — paciente NÃO
// cancela direto. Solicitação vai pra secretária via NotificationService.
//
// HIGH H9 — quando dailyCoSvc é injetado, GetByID(withRoom=true) gera um
// meeting_token escopado ao paciente (is_owner=false, sem screenshare) e
// retorna a URL completa em DailyJoinURL. Sala é privacy=private — sem
// token, Daily.co rejeita o acesso.
type PatientAppointmentService struct {
	db           *gorm.DB
	notification *NotificationService
	dailyCoSvc   *DailyCoService
}

func NewPatientAppointmentService(db *gorm.DB, notification *NotificationService) *PatientAppointmentService {
	return &PatientAppointmentService{db: db, notification: notification}
}

// WithDailyCo injeta DailyCoService — chamado em main.go após inicialização.
func (s *PatientAppointmentService) WithDailyCo(daily *DailyCoService) *PatientAppointmentService {
	s.dailyCoSvc = daily
	return s
}

// PatientAppointmentView é o payload retornado pro paciente.
// Sem campos internos (ExternalCalendarEventID, ConfirmationSentAt, etc).
// DailyJoinURL (HIGH H9) só é exposto quando dentro da janela permitida E
// vem com meeting_token escopado ao paciente.
type PatientAppointmentView struct {
	ID                 uuid.UUID  `json:"id"`
	ScheduledAt        time.Time  `json:"scheduledAt"`
	DurationMinutes    int        `json:"durationMinutes"`
	Type               string     `json:"type"`
	Status             string     `json:"status"`
	DoctorID           uuid.UUID  `json:"doctorId"`
	DoctorName         string     `json:"doctorName"`
	IsTelemedicine     bool       `json:"isTelemedicine"`
	DailyJoinURL       *string    `json:"dailyJoinUrl,omitempty"`
	Notes              string     `json:"notes,omitempty"`
	PatientConfirmedAt *time.Time `json:"patientConfirmedAt,omitempty"`
	ConfirmedAt        *time.Time `json:"confirmedAt,omitempty"`
	CancelledAt        *time.Time `json:"cancelledAt,omitempty"`
	MinutesUntilStart  int64      `json:"minutesUntilStart"`
}

// List retorna consultas do paciente. range="upcoming" devolve futuras+hoje;
// range="past" devolve histórico desc. Default upcoming.
func (s *PatientAppointmentService) List(patientID uuid.UUID, rangeKind string) ([]PatientAppointmentView, error) {
	var rows []models.Appointment
	q := s.db.Where("patient_id = ?", patientID)

	switch rangeKind {
	case "past":
		q = q.Where("scheduled_at < ?", time.Now().UTC()).
			Order("scheduled_at DESC")
	default: // upcoming
		q = q.Where("scheduled_at >= ?", time.Now().UTC().Add(-2*time.Hour)).
			Order("scheduled_at ASC")
	}

	if err := q.Find(&rows).Error; err != nil {
		return nil, err
	}

	out := make([]PatientAppointmentView, 0, len(rows))
	for i := range rows {
		out = append(out, *s.toView(&rows[i], false))
	}
	return out, nil
}

// GetByID retorna detalhe + DailyJoinURL (com meeting_token, HIGH H9) se
// dentro da janela permitida (-30min até +2h do horário marcado). Erro se
// appointment não pertence ao paciente.
func (s *PatientAppointmentService) GetByID(patientID, appointmentID uuid.UUID) (*PatientAppointmentView, error) {
	var appt models.Appointment
	err := s.db.Where("id = ? AND patient_id = ?", appointmentID, patientID).First(&appt).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAppointmentNotFound
		}
		return nil, err
	}
	return s.toView(&appt, true), nil
}

// Confirm marca PatientConfirmedAt + status=confirmed (se ainda scheduled).
func (s *PatientAppointmentService) Confirm(patientID, appointmentID uuid.UUID) error {
	var appt models.Appointment
	err := s.db.Where("id = ? AND patient_id = ?", appointmentID, patientID).First(&appt).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrAppointmentNotFound
		}
		return err
	}
	if appt.Status == models.AppointmentCancelled || appt.Status == models.AppointmentCompleted {
		return errors.New("não é possível confirmar esta consulta")
	}
	now := time.Now().UTC()
	updates := map[string]any{
		"patient_confirmed_at": now,
	}
	if appt.Status == models.AppointmentScheduled {
		updates["status"] = models.AppointmentConfirmed
		updates["confirmed_at"] = now
	}
	return s.db.Model(&models.Appointment{}).Where("id = ?", appointmentID).Updates(updates).Error
}

// RequestActionInput é o payload de request-cancel / request-reschedule.
type RequestActionInput struct {
	PatientID     uuid.UUID
	AppointmentID uuid.UUID
	Action        string // "cancel" | "reschedule"
	Reason        string
	UserID        uuid.UUID // pra audit
}

// RequestAction cria notificação interna pra equipe (humano-no-loop).
// NÃO cancela/reagenda direto — secretária decide na conversa com paciente.
func (s *PatientAppointmentService) RequestAction(in RequestActionInput) error {
	var appt models.Appointment
	err := s.db.Where("id = ? AND patient_id = ?", in.AppointmentID, in.PatientID).
		Preload("Patient").
		First(&appt).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrAppointmentNotFound
		}
		return err
	}

	verb := "remarcar"
	if in.Action == "cancel" {
		verb = "cancelar"
	}

	patientName := "paciente"
	if appt.Patient.Name != "" {
		patientName = appt.Patient.Name
	}

	title := fmt.Sprintf("%s solicita %s consulta", patientName, verb)
	msg := fmt.Sprintf(
		"Consulta marcada para %s. Motivo: %s",
		appt.ScheduledAt.Format("02/01/2006 15:04"),
		strings.TrimSpace(in.Reason),
	)
	if strings.TrimSpace(in.Reason) == "" {
		msg = fmt.Sprintf("Consulta marcada para %s. (sem motivo informado)",
			appt.ScheduledAt.Format("02/01/2006 15:04"))
	}

	actionURL := fmt.Sprintf("/appointments/%s", appt.ID.String())
	actionText := "Abrir consulta"

	// Despacha pra todos os staff que recebem alertas de leads (reuso convenção).
	// O médico responsável também é notificado.
	targets := s.notifyTargets(&appt)

	for _, uid := range targets {
		_ = s.notification.CreateNotification(
			uid,
			models.NotificationGeneral,
			title,
			msg,
			&appt.PatientID,
			nil,
			&actionURL,
			&actionText,
		)
	}
	return nil
}

// notifyTargets resolve os user IDs que recebem a notificação.
// V1: doctor da consulta + admin/manager/secretary disponíveis (top N).
func (s *PatientAppointmentService) notifyTargets(appt *models.Appointment) []uuid.UUID {
	out := []uuid.UUID{appt.DoctorID}
	var rows []models.User
	_ = s.db.Where("roles::text LIKE ? OR roles::text LIKE ? OR roles::text LIKE ?",
		"%admin%", "%manager%", "%secretary%").
		Limit(10).
		Find(&rows).Error
	for i := range rows {
		out = append(out, rows[i].ID)
	}
	return out
}

// ============================================================
// helpers
// ============================================================

func (s *PatientAppointmentService) toView(appt *models.Appointment, withRoom bool) *PatientAppointmentView {
	doctorName := ""
	var doctor models.User
	if err := s.db.Select("name").First(&doctor, "id = ?", appt.DoctorID).Error; err == nil {
		doctorName = doctor.Name
	}

	delta := time.Until(appt.ScheduledAt)
	view := &PatientAppointmentView{
		ID:                 appt.ID,
		ScheduledAt:        appt.ScheduledAt,
		DurationMinutes:    appt.DurationMinutes,
		Type:               string(appt.Type),
		Status:             string(appt.Status),
		DoctorID:           appt.DoctorID,
		DoctorName:         doctorName,
		IsTelemedicine:     appt.Type == models.AppointmentTelemedicine,
		Notes:              derefStr(appt.PatientNotes),
		PatientConfirmedAt: appt.PatientConfirmedAt,
		ConfirmedAt:        appt.ConfirmedAt,
		CancelledAt:        appt.CancelledAt,
		MinutesUntilStart:  int64(delta.Minutes()),
	}

	// HIGH H9 — DailyJoinURL com meeting_token só dentro da janela permitida
	// (-30min até +2h após horário). Sala é privacy=private; URL crua não
	// funciona — precisamos do token escopado ao paciente.
	if withRoom && view.IsTelemedicine && appt.DailyRoomURL != nil {
		minutesUntil := view.MinutesUntilStart
		if minutesUntil <= 30 && minutesUntil >= -120 {
			s.attachJoinURL(appt, view)
		}
	}
	return view
}

// attachJoinURL gera meeting_token (paciente, no-owner, sem screenshare) e
// popula view.DailyJoinURL. Em caso de erro, log + mantém URL nil (frontend
// mostra "sala temporariamente indisponível").
func (s *PatientAppointmentService) attachJoinURL(appt *models.Appointment, view *PatientAppointmentView) {
	if appt.DailyRoomName == nil || *appt.DailyRoomName == "" {
		// Salas pré-H9 podem ter URL mas sem name persistido.
		// Backcompat: expõe URL crua se Daily não configurado.
		if s.dailyCoSvc == nil || !s.dailyCoSvc.IsConfigured() {
			view.DailyJoinURL = appt.DailyRoomURL
		}
		return
	}
	if s.dailyCoSvc == nil || !s.dailyCoSvc.IsConfigured() {
		view.DailyJoinURL = appt.DailyRoomURL
		return
	}

	// Carrega nome do paciente pra display name.
	var patient models.Patient
	_ = s.db.Select("name").First(&patient, "id = ?", appt.PatientID).Error
	displayName := patientFirstName(patient.Name)

	// Janela do token: até +duration+15min (igual lobby).
	duration := appt.DurationMinutes
	if duration <= 0 {
		duration = 60
	}
	exp := appt.ScheduledAt.Add(time.Duration(duration)*time.Minute + 15*time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	tok, err := s.dailyCoSvc.CreateMeetingToken(ctx, MeetingTokenParams{
		RoomName:          *appt.DailyRoomName,
		UserName:          displayName,
		IsOwner:           false,
		ExpiresAt:         exp,
		EnableScreenshare: false,
	})
	if err != nil {
		log.Printf("⚠️  [PATIENT APPT] CreateMeetingToken apt=%s: %v", appt.ID, err)
		return
	}
	joinURL := BuildJoinURL(*appt.DailyRoomURL, tok)
	view.DailyJoinURL = &joinURL
}

// patientFirstName extrai primeiro nome (display opaco) do nome completo.
func patientFirstName(full string) string {
	full = strings.TrimSpace(full)
	if full == "" {
		return "paciente"
	}
	parts := strings.Fields(full)
	return parts[0]
}

func derefStr(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

