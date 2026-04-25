// Package services — AppointmentService gerencia o ciclo de vida de consultas.
//
// Calendar V1 (Bloco E) refactor:
//   - Create/Update/Cancel disparam side-effects async via goSafe:
//       1. Sincroniza Google Calendar (create/patch/delete event)
//       2. Cria/deleta sala Daily.co quando type=telemedicine
//       3. Envia notificações (email + WhatsApp) via AppointmentNotificationService
//   - Erros nos asyncs SÃO LOGADOS mas NÃO bloqueiam o request HTTP — o
//     appointment fica criado mesmo se Google/Daily/email falharem.
//   - Confirm é novo método: usado pela IA Claude (Bloco F) quando paciente
//     responde "ok" no WhatsApp do reminder.
package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

var (
	ErrAppointmentNotFound = errors.New("appointment not found")
	ErrAppointmentConflict = errors.New("appointment time slot already booked")
)

// AppointmentService coordena o ciclo de vida de Appointments + integrações.
//
// Os campos googleCalendarSvc, dailyCoSvc, notifSvc e cfg são opcionais — quando
// nil, o service degrada (não chama integrações). Isso facilita testes e
// inicialização parcial em dev.
type AppointmentService struct {
	db                *gorm.DB
	googleCalendarSvc *GoogleCalendarService
	dailyCoSvc        *DailyCoService
	notifSvc          *AppointmentNotificationService
	cfg               *config.Config
}

func NewAppointmentService(
	db *gorm.DB,
	googleCalendarSvc *GoogleCalendarService,
	dailyCoSvc *DailyCoService,
	notifSvc *AppointmentNotificationService,
	cfg *config.Config,
) *AppointmentService {
	return &AppointmentService{
		db:                db,
		googleCalendarSvc: googleCalendarSvc,
		dailyCoSvc:        dailyCoSvc,
		notifSvc:          notifSvc,
		cfg:               cfg,
	}
}

// Create cria uma nova consulta + dispara side-effects async.
//
// Validação de conflito SQL é mantida como defense-in-depth — a EXCLUDE
// constraint Postgres (Bloco A) é a fonte de verdade, mas devolvemos 409
// semântico antes de bater no DB.
func (s *AppointmentService) Create(userID uuid.UUID, userRole models.Role, req *dto.CreateAppointmentRequest) (*dto.AppointmentResponse, error) {
	// Resolve patient. Pra paciente self-booking: usa SelectedPatientID e exige
	// que o req.PatientID (se vier) bata. Pra staff (admin/secretary/manager/doctor):
	// usa req.PatientID direto, sem amarrar ao SelectedPatientID do usuário logado.
	var patientID uuid.UUID
	if userRole == models.RolePatient {
		var user models.User
		if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
			return nil, err
		}
		if user.SelectedPatientID == nil {
			return nil, errors.New("no patient selected - please select a patient first")
		}
		if req.PatientID != "" {
			pid, err := uuid.Parse(req.PatientID)
			if err != nil {
				return nil, errors.New("invalid patient id")
			}
			if pid != *user.SelectedPatientID {
				return nil, errors.New("patient id does not match selected patient")
			}
			patientID = pid
		} else {
			patientID = *user.SelectedPatientID
		}
	} else {
		if req.PatientID == "" {
			return nil, errors.New("patientId is required")
		}
		pid, err := uuid.Parse(req.PatientID)
		if err != nil {
			return nil, errors.New("invalid patient id")
		}
		patientID = pid
	}

	doctorID, err := uuid.Parse(req.DoctorID)
	if err != nil {
		return nil, errors.New("invalid doctor id")
	}

	var patient models.Patient
	if err := s.db.First(&patient, patientID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPatientNotFound
		}
		return nil, err
	}

	var doctor models.User
	if err := s.db.First(&doctor, doctorID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("doctor not found")
		}
		return nil, err
	}
	if !doctor.IsGranted(models.RoleDoctor) {
		return nil, errors.New("user is not a doctor")
	}

	scheduledAt, err := time.Parse(time.RFC3339, req.ScheduledAt)
	if err != nil {
		return nil, errors.New("invalid scheduled date format, expected RFC3339")
	}

	// Pre-flight conflict check (EXCLUDE constraint é a guarda real).
	endTime := scheduledAt.Add(time.Duration(req.DurationMinutes) * time.Minute)
	var conflict models.Appointment
	err = s.db.Where("doctor_id = ? AND status NOT IN (?, ?)", doctorID, models.AppointmentCancelled, models.AppointmentNoShow).
		Where("scheduled_at < ? AND (scheduled_at + INTERVAL '1 minute' * duration_minutes) > ?", endTime, scheduledAt).
		First(&conflict).Error
	if err == nil {
		return nil, ErrAppointmentConflict
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	appointment := models.Appointment{
		PatientID:       patientID,
		DoctorID:        doctorID,
		ScheduledAt:     scheduledAt,
		DurationMinutes: req.DurationMinutes,
		Type:            req.Type,
		Status:          models.AppointmentScheduled,
		Reason:          req.Reason,
		PatientNotes:    req.PatientNotes,
	}

	if err := s.db.Create(&appointment).Error; err != nil {
		// EXCLUDE constraint violation → 409.
		if isAppointmentOverlapErr(err) {
			return nil, ErrAppointmentConflict
		}
		return nil, err
	}

	// Async side-effects — não bloqueiam request.
	apptID := appointment.ID
	doctorIDCopy := doctorID
	apptType := appointment.Type
	scheduled := appointment.ScheduledAt
	duration := appointment.DurationMinutes
	patientName := patient.Name

	goSafe("appt_sync_google", func() {
		s.syncGoogleEventCreate(apptID, doctorIDCopy, apptType, scheduled, duration, patientName)
	})

	if apptType == models.AppointmentTelemedicine {
		goSafe("appt_create_daily", func() {
			s.createDailyRoom(apptID, scheduled, duration)
		})
	}

	goSafe("appt_send_confirmation", func() {
		if s.notifSvc == nil {
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if err := s.notifSvc.SendConfirmation(ctx, apptID); err != nil {
			log.Printf("⚠️  [APPT] SendConfirmation apt=%s: %v", apptID, err)
		}
	})

	return s.toDTO(&appointment), nil
}

// GetByID busca uma consulta por ID
func (s *AppointmentService) GetByID(appointmentID, userID uuid.UUID, userRole models.Role) (*dto.AppointmentResponse, error) {
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}

	if user.SelectedPatientID == nil {
		return nil, ErrAppointmentNotFound
	}

	var appointment models.Appointment
	query := s.db.Where("id = ?", appointmentID).Where("patient_id = ?", *user.SelectedPatientID)

	if err := query.First(&appointment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAppointmentNotFound
		}
		return nil, err
	}

	return s.toDTO(&appointment), nil
}

// GetByIDNoScope busca um appointment ignorando selectedPatient — usado por jobs
// internos (cron) e pela IA (que precisa lookup por patient do inbound, não do
// user logado). Não exposto via HTTP — caller é trusted.
func (s *AppointmentService) GetByIDNoScope(id uuid.UUID) (*models.Appointment, error) {
	var appt models.Appointment
	if err := s.db.Where("id = ?", id).First(&appt).Error; err != nil {
		return nil, err
	}
	return &appt, nil
}

// List lista consultas com filtros
func (s *AppointmentService) List(userID uuid.UUID, userRole models.Role, patientID, doctorID *uuid.UUID, status *models.AppointmentStatus, limit, offset int) ([]dto.AppointmentResponse, error) {
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}

	if user.SelectedPatientID == nil {
		return []dto.AppointmentResponse{}, nil
	}

	var appointments []models.Appointment
	query := s.db.Limit(limit).Offset(offset).Order("scheduled_at DESC")
	query = query.Where("patient_id = ?", *user.SelectedPatientID)

	if status != nil {
		query = query.Where("status = ?", *status)
	}

	if err := query.Find(&appointments).Error; err != nil {
		return nil, err
	}

	result := make([]dto.AppointmentResponse, len(appointments))
	for i, a := range appointments {
		result[i] = *s.toDTO(&a)
	}

	return result, nil
}

// Update atualiza uma consulta. Quando ScheduledAt ou Status mudam, dispara
// sincronização Google Calendar + notificação de reagendamento async.
func (s *AppointmentService) Update(appointmentID, userID uuid.UUID, userRole models.Role, req *dto.UpdateAppointmentRequest) (*dto.AppointmentResponse, error) {
	var appointment models.Appointment
	query := s.db.Where("id = ?", appointmentID)

	if err := query.First(&appointment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAppointmentNotFound
		}
		return nil, err
	}

	// Permissões: médico pode editar sua consulta, paciente não pode mudar status
	if userRole == models.RolePatient {
		var patient models.Patient
		if err := s.db.Where("user_id = ?", userID).First(&patient).Error; err != nil {
			return nil, ErrUnauthorized
		}
		if appointment.PatientID != patient.ID {
			return nil, ErrUnauthorized
		}
		req.Status = nil
	} else if userRole == models.RoleDoctor && appointment.DoctorID != userID {
		return nil, ErrUnauthorized
	}

	now := time.Now()
	oldScheduledAt := appointment.ScheduledAt
	scheduledChanged := false
	if req.ScheduledAt != nil {
		scheduledAt, err := time.Parse(time.RFC3339, *req.ScheduledAt)
		if err != nil {
			return nil, errors.New("invalid scheduled date format, expected RFC3339")
		}
		if !scheduledAt.Equal(appointment.ScheduledAt) {
			scheduledChanged = true
			appointment.ScheduledAt = scheduledAt
		}
	}
	if req.Status != nil {
		appointment.Status = *req.Status
		switch *req.Status {
		case models.AppointmentConfirmed:
			if appointment.ConfirmedAt == nil {
				appointment.ConfirmedAt = &now
			}
		case models.AppointmentCompleted:
			if appointment.CompletedAt == nil {
				appointment.CompletedAt = &now
			}
		case models.AppointmentCancelled:
			if appointment.CancelledAt == nil {
				appointment.CancelledAt = &now
			}
		}
	}
	if req.DoctorNotes != nil {
		appointment.DoctorNotes = req.DoctorNotes
	}
	if req.Diagnosis != nil {
		appointment.Diagnosis = req.Diagnosis
	}

	if err := s.db.Save(&appointment).Error; err != nil {
		if isAppointmentOverlapErr(err) {
			return nil, ErrAppointmentConflict
		}
		return nil, err
	}

	// Async: se horário mudou, sincroniza Google + notifica paciente.
	if scheduledChanged {
		apptID := appointment.ID
		doctorID := appointment.DoctorID
		apptType := appointment.Type
		newScheduled := appointment.ScheduledAt
		duration := appointment.DurationMinutes
		eventID := appointment.ExternalCalendarEventID

		goSafe("appt_sync_google_patch", func() {
			s.syncGoogleEventPatch(apptID, doctorID, eventID, apptType, newScheduled, duration)
		})

		goSafe("appt_send_reschedule", func() {
			if s.notifSvc == nil {
				return
			}
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			if err := s.notifSvc.SendReschedule(ctx, apptID, oldScheduledAt); err != nil {
				log.Printf("⚠️  [APPT] SendReschedule apt=%s: %v", apptID, err)
			}
		})
	}

	return s.toDTO(&appointment), nil
}

// Cancel cancela uma consulta + dispara cleanup async (Google + Daily) + notif.
func (s *AppointmentService) Cancel(appointmentID, userID uuid.UUID, userRole models.Role, req *dto.CancelAppointmentRequest) (*dto.AppointmentResponse, error) {
	var appointment models.Appointment
	if err := s.db.Where("id = ?", appointmentID).First(&appointment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAppointmentNotFound
		}
		return nil, err
	}

	if userRole == models.RolePatient {
		var patient models.Patient
		if err := s.db.Where("user_id = ?", userID).First(&patient).Error; err != nil {
			return nil, ErrUnauthorized
		}
		if appointment.PatientID != patient.ID {
			return nil, ErrUnauthorized
		}
	} else if userRole == models.RoleDoctor && appointment.DoctorID != userID {
		return nil, ErrUnauthorized
	}

	now := time.Now()
	appointment.Status = models.AppointmentCancelled
	appointment.CancelledAt = &now
	appointment.CancellationReason = &req.Reason

	if err := s.db.Save(&appointment).Error; err != nil {
		return nil, err
	}

	apptID := appointment.ID
	doctorID := appointment.DoctorID
	eventID := appointment.ExternalCalendarEventID
	roomName := appointment.DailyRoomName

	if eventID != nil && *eventID != "" {
		goSafe("appt_delete_google", func() {
			s.deleteGoogleEvent(apptID, doctorID, *eventID)
		})
	}

	if roomName != nil && *roomName != "" {
		goSafe("appt_delete_daily", func() {
			s.deleteDailyRoom(apptID, *roomName)
		})
	}

	goSafe("appt_send_cancellation", func() {
		if s.notifSvc == nil {
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if err := s.notifSvc.SendCancellation(ctx, apptID); err != nil {
			log.Printf("⚠️  [APPT] SendCancellation apt=%s: %v", apptID, err)
		}
	})

	return s.toDTO(&appointment), nil
}

// Confirm marca a consulta como confirmada. actorUserID nil = ação automática
// da IA (paciente respondeu "ok" no WhatsApp). Quando setado, é o staff/admin
// confirmando manualmente via UI.
//
// Idempotente — se já está confirmed, no-op.
func (s *AppointmentService) Confirm(ctx context.Context, appointmentID uuid.UUID, actorUserID *uuid.UUID) error {
	var appt models.Appointment
	if err := s.db.WithContext(ctx).Where("id = ?", appointmentID).First(&appt).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrAppointmentNotFound
		}
		return err
	}
	if appt.Status == models.AppointmentConfirmed {
		return nil
	}
	// Não confirma consultas já canceladas/concluídas/no-show.
	if appt.Status == models.AppointmentCancelled ||
		appt.Status == models.AppointmentCompleted ||
		appt.Status == models.AppointmentNoShow {
		return fmt.Errorf("appointment %s status=%s não pode ser confirmado", appointmentID, appt.Status)
	}

	now := time.Now().UTC()
	updates := map[string]any{
		"status":       models.AppointmentConfirmed,
		"confirmed_at": now,
	}
	if err := s.db.WithContext(ctx).Model(&models.Appointment{}).
		Where("id = ?", appointmentID).
		Updates(updates).Error; err != nil {
		return err
	}

	// Audit trail: cria LeadActivity{note_added, internal} no patient.
	// Quando confirmação vem da IA (actor=nil), grava metadata.ai=true.
	patientID := appt.PatientID
	content := "Consulta confirmada"
	meta := map[string]any{
		"appointment_id": appointmentID.String(),
	}
	if actorUserID == nil {
		meta["ai_auto_confirmed"] = true
		content = "Consulta confirmada automaticamente via IA (resposta WhatsApp do paciente)"
	}

	metaJSON, _ := json.Marshal(meta)
	activity := &models.LeadActivity{
		PatientID:   &patientID,
		Type:        models.LeadActivityNoteAdded,
		Channel:     models.LeadChannelInternal,
		Content:     &content,
		Metadata:    datatypes.JSON(metaJSON),
		ActorUserID: actorUserID,
	}
	if err := s.db.WithContext(ctx).Create(activity).Error; err != nil {
		// Audit é best-effort — não desfaz confirmação.
		log.Printf("⚠️  [APPT] Confirm audit log apt=%s: %v", appointmentID, err)
	}
	return nil
}

// Delete faz soft delete de uma consulta
func (s *AppointmentService) Delete(appointmentID uuid.UUID, userRole models.Role) error {
	if userRole != models.RoleAdmin {
		return ErrUnauthorized
	}

	result := s.db.Where("id = ?", appointmentID).Delete(&models.Appointment{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrAppointmentNotFound
	}

	return nil
}

// ============================================================
// Async helpers — Google Calendar + Daily.co
// ============================================================

// syncGoogleEventCreate cria evento no calendar dedicado do médico se ele tem
// CalendarCredential ativa. Persiste ExternalCalendarEventID após sucesso.
//
// Privacidade: title é "Plenya — {tipo} · {Iniciais}" — sem nome completo.
func (s *AppointmentService) syncGoogleEventCreate(
	apptID, doctorID uuid.UUID,
	apptType models.AppointmentType,
	scheduled time.Time,
	duration int,
	patientName string,
) {
	if s.googleCalendarSvc == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Verifica se médico tem credencial.
	var cred models.CalendarCredential
	err := s.db.WithContext(ctx).
		Where("user_id = ? AND provider = ?", doctorID, models.CalendarProviderGoogle).
		First(&cred).Error
	if err != nil {
		// Sem credencial: silencioso (médico não conectou Google).
		return
	}
	if !cred.IsActive() {
		return
	}

	accessToken, err := s.googleCalendarSvc.GetValidAccessToken(ctx, doctorID)
	if err != nil {
		log.Printf("⚠️  [APPT] sync google: GetValidAccessToken doctor=%s: %v", doctorID, err)
		return
	}

	summary := buildEventSummary(apptType, patientName)
	end := scheduled.Add(time.Duration(duration) * time.Minute)
	out, err := s.googleCalendarSvc.CreateEvent(ctx, accessToken, EventInput{
		CalendarID:  cred.DedicatedCalendarID,
		Summary:     summary,
		Description: "",
		Location:    "Plenya Saúde",
		StartUTC:    scheduled.UTC(),
		EndUTC:      end.UTC(),
		Visibility:  "private",
	})
	if err != nil {
		log.Printf("⚠️  [APPT] sync google CreateEvent apt=%s: %v", apptID, err)
		return
	}
	if err := s.db.WithContext(ctx).Model(&models.Appointment{}).
		Where("id = ?", apptID).
		Update("external_calendar_event_id", out.EventID).Error; err != nil {
		log.Printf("⚠️  [APPT] sync google persist event_id apt=%s: %v", apptID, err)
	}
}

// syncGoogleEventPatch atualiza evento existente (horário mudou). Se eventID for
// nil, tenta criar um novo (caso o appt foi criado antes do médico conectar Google).
func (s *AppointmentService) syncGoogleEventPatch(
	apptID, doctorID uuid.UUID,
	eventID *string,
	apptType models.AppointmentType,
	scheduled time.Time,
	duration int,
) {
	if s.googleCalendarSvc == nil {
		return
	}
	if eventID == nil || *eventID == "" {
		// Sem evento prévio — vira create.
		s.syncGoogleEventCreate(apptID, doctorID, apptType, scheduled, duration, "")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var cred models.CalendarCredential
	if err := s.db.WithContext(ctx).
		Where("user_id = ? AND provider = ?", doctorID, models.CalendarProviderGoogle).
		First(&cred).Error; err != nil {
		return
	}
	if !cred.IsActive() {
		return
	}
	accessToken, err := s.googleCalendarSvc.GetValidAccessToken(ctx, doctorID)
	if err != nil {
		log.Printf("⚠️  [APPT] patch google: token doctor=%s: %v", doctorID, err)
		return
	}

	summary := buildEventSummary(apptType, "")
	end := scheduled.Add(time.Duration(duration) * time.Minute)
	if err := s.googleCalendarSvc.PatchEvent(ctx, accessToken, cred.DedicatedCalendarID, *eventID, EventInput{
		CalendarID:  cred.DedicatedCalendarID,
		Summary:     summary,
		Description: "",
		Location:    "Plenya Saúde",
		StartUTC:    scheduled.UTC(),
		EndUTC:      end.UTC(),
		Visibility:  "private",
	}); err != nil {
		log.Printf("⚠️  [APPT] patch google apt=%s: %v", apptID, err)
	}
}

// deleteGoogleEvent remove evento. Idempotente (404/410 são tratados em DeleteEvent).
func (s *AppointmentService) deleteGoogleEvent(apptID, doctorID uuid.UUID, eventID string) {
	if s.googleCalendarSvc == nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var cred models.CalendarCredential
	if err := s.db.WithContext(ctx).
		Where("user_id = ? AND provider = ?", doctorID, models.CalendarProviderGoogle).
		First(&cred).Error; err != nil {
		return
	}
	if !cred.IsActive() {
		return
	}
	accessToken, err := s.googleCalendarSvc.GetValidAccessToken(ctx, doctorID)
	if err != nil {
		log.Printf("⚠️  [APPT] delete google: token doctor=%s: %v", doctorID, err)
		return
	}
	if err := s.googleCalendarSvc.DeleteEvent(ctx, accessToken, cred.DedicatedCalendarID, eventID); err != nil {
		log.Printf("⚠️  [APPT] delete google apt=%s: %v", apptID, err)
	}
}

// createDailyRoom cria sala Daily.co pra teleconsulta. Persiste URL+Name.
//
// Expiração: scheduled + duration + 1h buffer (paciente atrasado, etc).
func (s *AppointmentService) createDailyRoom(apptID uuid.UUID, scheduled time.Time, duration int) {
	if s.dailyCoSvc == nil || !s.dailyCoSvc.IsConfigured() {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Prefix legível usando primeiros 8 chars do UUID.
	prefix := "plenya-" + apptID.String()[:8]
	exp := scheduled.Add(time.Duration(duration)*time.Minute + 1*time.Hour)

	room, err := s.dailyCoSvc.CreateRoom(ctx, prefix, exp)
	if err != nil {
		log.Printf("⚠️  [APPT] daily create apt=%s: %v", apptID, err)
		return
	}
	if err := s.db.WithContext(ctx).Model(&models.Appointment{}).
		Where("id = ?", apptID).
		Updates(map[string]any{
			"daily_room_url":  room.URL,
			"daily_room_name": room.Name,
		}).Error; err != nil {
		log.Printf("⚠️  [APPT] daily persist apt=%s: %v", apptID, err)
	}
}

// deleteDailyRoom remove sala (idempotente).
func (s *AppointmentService) deleteDailyRoom(apptID uuid.UUID, name string) {
	if s.dailyCoSvc == nil || !s.dailyCoSvc.IsConfigured() {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := s.dailyCoSvc.DeleteRoom(ctx, name); err != nil {
		log.Printf("⚠️  [APPT] daily delete apt=%s: %v", apptID, err)
	}
}

// buildEventSummary monta título opaco pro Google Calendar.
//
// "Plenya — {tipo}" se patientName vazio, ou "Plenya — {tipo} · {Iniciais}".
// LGPD: NUNCA inclui nome completo, CPF, dados sensíveis.
func buildEventSummary(t models.AppointmentType, patientName string) string {
	label := appointmentTypeLabel(t)
	initials := patientInitials(patientName)
	if initials == "" {
		return "Plenya — " + label
	}
	return "Plenya — " + label + " · " + initials
}

// patientInitials extrai iniciais ("Maria Silva Santos" → "M.S.S."). Limita a 4.
func patientInitials(name string) string {
	parts := strings.Fields(strings.TrimSpace(name))
	if len(parts) == 0 {
		return ""
	}
	var b strings.Builder
	for i, p := range parts {
		if i >= 4 {
			break
		}
		r := []rune(p)
		if len(r) == 0 {
			continue
		}
		b.WriteRune(r[0])
		b.WriteRune('.')
	}
	return strings.ToUpper(b.String())
}

// isAppointmentOverlapErr detecta violação da EXCLUDE constraint
// `appointments_no_overlap` (Bloco A). Postgres SQLState 23P01.
func isAppointmentOverlapErr(err error) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	return strings.Contains(msg, "appointments_no_overlap") ||
		strings.Contains(msg, "23P01") ||
		strings.Contains(msg, "exclusion constraint")
}

// toDTO converte Appointment para AppointmentResponse
func (s *AppointmentService) toDTO(appointment *models.Appointment) *dto.AppointmentResponse {
	resp := &dto.AppointmentResponse{
		ID:                 appointment.ID.String(),
		PatientID:          appointment.PatientID.String(),
		DoctorID:           appointment.DoctorID.String(),
		ScheduledAt:        appointment.ScheduledAt.Format(time.RFC3339),
		DurationMinutes:    appointment.DurationMinutes,
		Type:               appointment.Type,
		Status:             appointment.Status,
		Reason:             appointment.Reason,
		PatientNotes:       appointment.PatientNotes,
		DoctorNotes:        appointment.DoctorNotes,
		Diagnosis:          appointment.Diagnosis,
		CancellationReason: appointment.CancellationReason,
		CreatedAt:          appointment.CreatedAt.Format(time.RFC3339),
		UpdatedAt:          appointment.UpdatedAt.Format(time.RFC3339),
	}

	if appointment.AnamnesisID != nil {
		aid := appointment.AnamnesisID.String()
		resp.AnamnesisID = &aid
	}
	if appointment.ConfirmedAt != nil {
		ca := appointment.ConfirmedAt.Format(time.RFC3339)
		resp.ConfirmedAt = &ca
	}
	if appointment.CompletedAt != nil {
		ca := appointment.CompletedAt.Format(time.RFC3339)
		resp.CompletedAt = &ca
	}
	if appointment.CancelledAt != nil {
		ca := appointment.CancelledAt.Format(time.RFC3339)
		resp.CancelledAt = &ca
	}

	return resp
}
