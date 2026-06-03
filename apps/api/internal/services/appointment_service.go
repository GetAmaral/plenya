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
	ErrAppointmentNotFound       = errors.New("appointment not found")
	ErrAppointmentConflict       = errors.New("appointment time slot already booked")
	ErrAppointmentNotTelemed     = errors.New("appointment is not telemedicine")
	ErrAppointmentRoomNotReady   = errors.New("appointment room not ready yet")
	ErrAppointmentOutsideWindow  = errors.New("appointment is outside the access window")
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
	telemedLobbySvc   *TelemedLobbyService
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

// WithTelemedLobby injeta TelemedLobbyService — chamado quando AppointmentService
// já existe mas o lobby foi criado depois (ordem de DI no main).
// Quando setado, createDailyRoom também gera lobby token automaticamente.
func (s *AppointmentService) WithTelemedLobby(svc *TelemedLobbyService) *AppointmentService {
	s.telemedLobbySvc = svc
	return s
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
	if !doctor.IsGranted(models.RoleDoctor) &&
		!doctor.IsGranted(models.RoleNutritionist) &&
		!doctor.IsGranted(models.RolePsychologist) &&
		!doctor.IsGranted(models.RolePhysicalEducator) {
		return nil, errors.New("user is not a professional with agenda")
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

	// Continuum (Fase 3): se vier ContinuumItemID, valida que existe e
	// pertence a um Continuum desse paciente — evita ancoragem cruzada.
	var continuumItemID *uuid.UUID
	if req.ContinuumItemID != nil && *req.ContinuumItemID != "" {
		cid, err := uuid.Parse(*req.ContinuumItemID)
		if err != nil {
			return nil, errors.New("invalid continuumItemId")
		}
		var item models.PatientContinuumItem
		err = s.db.
			Joins("JOIN patient_continuums pc ON pc.id = patient_continuum_items.continuum_id").
			Where("patient_continuum_items.id = ? AND pc.patient_id = ?", cid, patientID).
			First(&item).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("continuum item not found for this patient")
			}
			return nil, err
		}
		continuumItemID = &cid
		appointment.ContinuumItemID = &cid
	}

	if err := s.db.Create(&appointment).Error; err != nil {
		// EXCLUDE constraint violation → 409.
		if isAppointmentOverlapErr(err) {
			return nil, ErrAppointmentConflict
		}
		return nil, err
	}

	// Anexa relations já carregadas pra toDTO popular nested patient/doctor
	// sem round-trip extra.
	appointment.Patient = patient
	appointment.Doctor = doctor

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

	// Continuum: propaga ancoragem síncrona (rápido — só um UPDATE) pra que
	// a timeline reflita imediatamente após criar a consulta.
	if continuumItemID != nil {
		if err := s.db.Model(&models.PatientContinuumItem{}).
			Where("id = ?", *continuumItemID).
			Updates(map[string]any{
				"status":         models.ContinuumItemScheduled,
				"appointment_id": apptID,
			}).Error; err != nil {
			log.Printf("⚠️  [APPT] continuum anchor item=%s: %v", *continuumItemID, err)
		}
	}

	return s.toDTO(&appointment), nil
}

// GetByID busca uma consulta por ID. Patient sempre escopado a SelectedPatientID;
// staff (admin/secretary/manager/doctor) acessa qualquer consulta.
func (s *AppointmentService) GetByID(appointmentID, userID uuid.UUID, userRole models.Role) (*dto.AppointmentResponse, error) {
	query := s.db.Preload("Patient").Preload("Doctor").Where("id = ?", appointmentID)

	if userRole == models.RolePatient {
		var user models.User
		if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
			return nil, err
		}
		if user.SelectedPatientID == nil {
			return nil, ErrAppointmentNotFound
		}
		query = query.Where("patient_id = ?", *user.SelectedPatientID)
	}

	var appointment models.Appointment
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

// ListAppointmentsParams — filtros estruturados pra List.
type ListAppointmentsParams struct {
	PatientID *uuid.UUID
	DoctorIDs []uuid.UUID // suporta calendário multi-médico (admin/secretary)
	Status    *models.AppointmentStatus
	DateFrom  *time.Time
	DateTo    *time.Time
	Limit     int
	Offset    int
}

// List lista consultas com filtros.
// - Patient: sempre escopado a SelectedPatientID (segurança).
// - Staff (admin/secretary/manager/doctor): livre, usa filtros do params.
//   Doctor sem filtro vê todas as consultas; pode passar DoctorIDs=[self.ID]
//   pra ver só as próprias.
func (s *AppointmentService) List(userID uuid.UUID, userRole models.Role, params ListAppointmentsParams) ([]dto.AppointmentResponse, error) {
	if params.Limit <= 0 || params.Limit > 500 {
		params.Limit = 100
	}

	query := s.db.Preload("Patient").Preload("Doctor").
		Limit(params.Limit).Offset(params.Offset).
		Order("scheduled_at ASC")

	if userRole == models.RolePatient {
		var user models.User
		if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
			return nil, err
		}
		if user.SelectedPatientID == nil {
			return []dto.AppointmentResponse{}, nil
		}
		query = query.Where("patient_id = ?", *user.SelectedPatientID)
	} else {
		if params.PatientID != nil {
			query = query.Where("patient_id = ?", *params.PatientID)
		}
		if len(params.DoctorIDs) > 0 {
			query = query.Where("doctor_id IN ?", params.DoctorIDs)
		}
	}

	if params.Status != nil {
		query = query.Where("status = ?", *params.Status)
	}
	if params.DateFrom != nil {
		query = query.Where("scheduled_at >= ?", *params.DateFrom)
	}
	if params.DateTo != nil {
		query = query.Where("scheduled_at < ?", *params.DateTo)
	}

	var appointments []models.Appointment
	if err := query.Find(&appointments).Error; err != nil {
		return nil, err
	}

	result := make([]dto.AppointmentResponse, len(appointments))
	for i := range appointments {
		result[i] = *s.toDTO(&appointments[i])
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

	// Continuum: se a consulta tem item ancorado e o status virou completed,
	// marca o item como completed também (síncrono, query simples).
	if appointment.ContinuumItemID != nil &&
		req.Status != nil && *req.Status == models.AppointmentCompleted {
		refType := "appointment"
		now := time.Now().UTC()
		if err := s.db.Model(&models.PatientContinuumItem{}).
			Where("id = ?", *appointment.ContinuumItemID).
			Updates(map[string]any{
				"status":             models.ContinuumItemCompleted,
				"completed_at":       now,
				"completed_ref_type": refType,
				"completed_ref_id":   appointment.ID,
			}).Error; err != nil {
			log.Printf("⚠️  [APPT] continuum complete item=%s: %v", *appointment.ContinuumItemID, err)
		}
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

// CheckIn registra a chegada do paciente no balcão (status -> checked_in).
// Usado pela recepção. Idempotente: se já está checked_in/in_progress, no-op.
//
// Transições válidas: scheduled|confirmed -> checked_in.
// Bloqueia consultas canceladas/concluídas/no-show.
func (s *AppointmentService) CheckIn(ctx context.Context, appointmentID, actorUserID uuid.UUID) (*dto.AppointmentResponse, error) {
	var appt models.Appointment
	if err := s.db.WithContext(ctx).Preload("Patient").Preload("Doctor").
		Where("id = ?", appointmentID).First(&appt).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAppointmentNotFound
		}
		return nil, err
	}
	// Idempotente — já chegou (ou já está em atendimento).
	if appt.Status == models.AppointmentCheckedIn || appt.Status == models.AppointmentInProgress {
		return s.toDTO(&appt), nil
	}
	if appt.Status != models.AppointmentScheduled && appt.Status != models.AppointmentConfirmed {
		return nil, fmt.Errorf("appointment %s status=%s não pode receber check-in", appointmentID, appt.Status)
	}

	now := time.Now().UTC()
	appt.Status = models.AppointmentCheckedIn
	appt.CheckedInAt = &now
	if err := s.db.WithContext(ctx).Save(&appt).Error; err != nil {
		return nil, err
	}
	return s.toDTO(&appt), nil
}

// StartConsultation marca o início do atendimento (status -> in_progress).
// Normalmente acionado pelo médico ao iniciar. Idempotente se já in_progress.
//
// Transições válidas: scheduled|confirmed|checked_in -> in_progress.
// Se não houve check-in prévio (atendimento direto), CheckedInAt fica nulo e o
// tempo de espera não é computado.
func (s *AppointmentService) StartConsultation(ctx context.Context, appointmentID, actorUserID uuid.UUID) (*dto.AppointmentResponse, error) {
	var appt models.Appointment
	if err := s.db.WithContext(ctx).Preload("Patient").Preload("Doctor").
		Where("id = ?", appointmentID).First(&appt).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAppointmentNotFound
		}
		return nil, err
	}
	if appt.Status == models.AppointmentInProgress {
		return s.toDTO(&appt), nil
	}
	if appt.Status != models.AppointmentScheduled &&
		appt.Status != models.AppointmentConfirmed &&
		appt.Status != models.AppointmentCheckedIn {
		return nil, fmt.Errorf("appointment %s status=%s não pode iniciar atendimento", appointmentID, appt.Status)
	}

	now := time.Now().UTC()
	appt.Status = models.AppointmentInProgress
	appt.StartedAt = &now
	if err := s.db.WithContext(ctx).Save(&appt).Error; err != nil {
		return nil, err
	}
	return s.toDTO(&appt), nil
}

// TelemedConsentTermText — termo canônico de consentimento de telemedicina (CFM 2.314/2022
// + 2.454/2026 sobre IA de apoio), com cláusulas de gravação, transcrição, uso de IA de apoio
// e transferência internacional de dados (LGPD). É o texto exibido ao registrar e gravado em
// Appointment.TelemedConsentText (prova do que foi apresentado). Fonte versionada:
// docs/emr/termo-consentimento-telemedicina.md. Alterar aqui ao revisar o termo.
const TelemedConsentTermText = `TERMO DE CONSENTIMENTO LIVRE E ESCLARECIDO PARA ATENDIMENTO POR TELEMEDICINA

Atendimento prestado por Plenya Serviços de Saúde Ltda. (CNPJ 66.991.259/0001-50), nos termos da Resolução CFM nº 2.314/2022.

Declaro estar ciente e de acordo com o seguinte:

1. Este atendimento será realizado por telemedicina, de forma remota, por meio de vídeo em plataforma segura, sem a presença física do médico no mesmo ambiente.
2. Fui informado(a) de que tenho o direito de optar pelo atendimento presencial a qualquer momento, sem qualquer prejuízo ao meu cuidado.
3. Compreendo que a telemedicina possui limitações — em especial quanto ao exame físico — e que o médico poderá recomendar avaliação presencial, exames complementares ou encaminhamento sempre que julgar necessário para a minha segurança.
4. Mantêm-se integralmente o sigilo médico e a proteção dos meus dados pessoais e de saúde, tratados conforme a legislação aplicável (LGPD).
5. Posso interromper o atendimento a qualquer momento e esclarecer dúvidas com o médico.
6. Em caso de emergência ou urgência, devo procurar imediatamente um serviço de pronto atendimento presencial.
7. Autorizo a gravação desta teleconsulta, que integrará o meu prontuário e ficará sujeita ao mesmo sigilo e proteção de dados.
8. Fui informado(a) de que a gravação poderá ser transcrita e de que ferramentas de inteligência artificial poderão ser utilizadas como apoio ao médico para organizar e resumir as informações da consulta no prontuário. Esses recursos são apenas de apoio: o médico permanece integralmente responsável pelo conteúdo, que é por ele revisado e validado. Posso recusar o uso dessas ferramentas a qualquer momento, sem qualquer prejuízo ao meu atendimento.
9. Para viabilizar o vídeo, a gravação e a transcrição, meus dados pessoais e de saúde poderão ser processados por prestadores de serviço de tecnologia contratados pela Plenya, inclusive localizados no exterior, sempre sob contrato com cláusulas de proteção de dados e o mesmo dever de sigilo, nos termos da LGPD. O arquivamento do meu prontuário permanece no Brasil.

Declaro que fui devidamente esclarecido(a), tive a oportunidade de tirar dúvidas e consinto livremente com a realização do atendimento por telemedicina, inclusive com a gravação, a transcrição e o uso de ferramentas de inteligência artificial de apoio acima descritos.`

// RegisterTelemedConsent carimba o consentimento de telemedicina no prontuário (CFM 2.314/2022).
// Só para Type=telemedicine. Idempotente (não re-carimba se já registrado). Mode default 'verbal'.
func (s *AppointmentService) RegisterTelemedConsent(ctx context.Context, appointmentID, actorUserID uuid.UUID, mode string) (*dto.AppointmentResponse, error) {
	var appt models.Appointment
	if err := s.db.WithContext(ctx).Preload("Patient").Preload("Doctor").
		Where("id = ?", appointmentID).First(&appt).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAppointmentNotFound
		}
		return nil, err
	}
	if appt.Type != models.AppointmentTelemedicine {
		return nil, ErrAppointmentNotTelemed
	}
	if appt.TelemedConsentAt != nil {
		return s.toDTO(&appt), nil // idempotente
	}
	if mode != "written" {
		mode = "verbal"
	}

	now := time.Now().UTC()
	text := TelemedConsentTermText
	appt.TelemedConsentAt = &now
	appt.TelemedConsentText = &text
	appt.TelemedConsentByUserID = &actorUserID
	appt.TelemedConsentMode = &mode
	if err := s.db.WithContext(ctx).Save(&appt).Error; err != nil {
		return nil, err
	}
	return s.toDTO(&appt), nil
}

// GetTelemedJoinURL gera meeting_token de owner pro staff (médico/secretaria)
// e retorna a URL completa pra abrir a sala.
//
// HIGH H9 — sala é privacy=private. Owner = pode ejetar, screenshare habilitado.
// Token é gerado fresh a cada chamada (no caching) — janela curta (até closesAt).
//
// Erros:
//   - ErrAppointmentNotFound — id inválido
//   - ErrAppointmentNotTelemed — type != telemedicine
//   - ErrAppointmentRoomNotReady — sala ainda não criada (async em flight)
//   - ErrAppointmentOutsideWindow — fora de [-TelemedWindowBefore, +duration+TelemedWindowAfter]
//   - ErrDailyNotConfigured — Daily.co API key não setada
func (s *AppointmentService) GetTelemedJoinURL(ctx context.Context, appointmentID, userID uuid.UUID, userRole models.Role) (string, error) {
	var appt models.Appointment
	q := s.db.WithContext(ctx).Preload("Doctor").Where("id = ?", appointmentID)
	if err := q.First(&appt).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", ErrAppointmentNotFound
		}
		return "", err
	}

	// Doctor só pode acessar suas próprias consultas; admin/manager/secretary
	// passam direto. Patient role NÃO chega aqui (rota é staff-only) — defesa
	// extra mesmo assim.
	if userRole == models.RolePatient {
		return "", ErrUnauthorized
	}
	if userRole == models.RoleDoctor && appt.DoctorID != userID {
		return "", ErrUnauthorized
	}

	if appt.Type != models.AppointmentTelemedicine {
		return "", ErrAppointmentNotTelemed
	}
	// Janela: -TelemedWindowBefore até +duration+TelemedWindowAfter do scheduledAt.
	// Bloqueia geração de token muito antes/depois — limita superfície de abuso.
	duration := appt.DurationMinutes
	if duration <= 0 {
		duration = 60
	}
	now := time.Now().UTC()
	opensAt := appt.ScheduledAt.Add(-TelemedWindowBefore)
	closesAt := appt.ScheduledAt.Add(time.Duration(duration)*time.Minute + TelemedWindowAfter)
	if now.Before(opensAt) || now.After(closesAt) {
		return "", ErrAppointmentOutsideWindow
	}

	if s.dailyCoSvc == nil || !s.dailyCoSvc.IsConfigured() {
		// Backcompat dev: sem Daily, devolve URL crua se já existir.
		if appt.DailyRoomURL != nil {
			return *appt.DailyRoomURL, nil
		}
		return "", ErrDailyNotConfigured
	}

	// Sala sob demanda: se a criação assíncrona (no create do appointment) ainda
	// não terminou — ou falhou — cria agora, sincronamente. Elimina a race de
	// clicar "Abrir sala" logo após criar a consulta (antes devolvia 409).
	if appt.DailyRoomURL == nil || appt.DailyRoomName == nil {
		name, url, cerr := s.ensureDailyRoomSync(ctx, appt.ID, appt.ScheduledAt, duration)
		if cerr != nil {
			return "", cerr
		}
		appt.DailyRoomName = &name
		appt.DailyRoomURL = &url
	}

	doctorName := strings.TrimSpace(appt.Doctor.Name)
	if doctorName == "" {
		doctorName = "Médico"
	}
	// Gravação + transcrição auto-iniciam SÓ com consentimento de telemedicina
	// registrado (CFM 2.314/2022, cláusula de gravação). Sem consentimento, o
	// token do médico não liga nada — a sala continua só de vídeo.
	withConsent := appt.TelemedConsentAt != nil
	tok, err := s.dailyCoSvc.CreateMeetingToken(ctx, MeetingTokenParams{
		RoomName:               *appt.DailyRoomName,
		UserName:               doctorName,
		IsOwner:                true,
		ExpiresAt:              closesAt,
		EnableScreenshare:      true,
		StartCloudRecording:    withConsent,
		AutoStartTranscription: withConsent,
	})
	if err != nil {
		return "", err
	}
	return BuildJoinURL(*appt.DailyRoomURL, tok), nil
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

// createDailyRoom cria sala Daily.co pra teleconsulta (chamado async no create
// do appointment). Best-effort — erros só logam. A criação real fica em
// ensureDailyRoomSync, reusada também sob demanda no GetTelemedJoinURL.
func (s *AppointmentService) createDailyRoom(apptID uuid.UUID, scheduled time.Time, duration int) {
	if s.dailyCoSvc == nil || !s.dailyCoSvc.IsConfigured() {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if _, _, err := s.ensureDailyRoomSync(ctx, apptID, scheduled, duration); err != nil {
		log.Printf("⚠️  [APPT] daily create apt=%s: %v", apptID, err)
	}
}

// ensureDailyRoomSync garante (sincronamente) que o appointment tem sala Daily +
// lobby token. Idempotente: se já existe sala, devolve a existente sem recriar
// (re-lê do banco pra cobrir a goroutine async que pode ter populado no meio).
// Retorna (roomName, roomURL).
func (s *AppointmentService) ensureDailyRoomSync(ctx context.Context, apptID uuid.UUID, scheduled time.Time, duration int) (string, string, error) {
	if s.dailyCoSvc == nil || !s.dailyCoSvc.IsConfigured() {
		return "", "", ErrDailyNotConfigured
	}

	// Re-lê: a criação async pode ter populado entre o load do caller e agora.
	var cur models.Appointment
	if err := s.db.WithContext(ctx).
		Select("daily_room_url", "daily_room_name").
		Where("id = ?", apptID).First(&cur).Error; err == nil {
		if cur.DailyRoomURL != nil && cur.DailyRoomName != nil {
			return *cur.DailyRoomName, *cur.DailyRoomURL, nil
		}
	}

	// Prefix legível usando primeiros 8 chars do UUID.
	prefix := "plenya-" + apptID.String()[:8]
	if duration <= 0 {
		duration = 60
	}
	// Sala precisa sobreviver até o fim da janela de acesso (closesAt =
	// scheduled + duration + TelemedWindowAfter) + margem, senão a janela diz
	// "aberta" mas a sala já foi auto-deletada pelo Daily.
	exp := scheduled.Add(time.Duration(duration)*time.Minute + TelemedWindowAfter + 30*time.Minute)
	// Daily recusa criar sala com exp no passado (400 invalid-request-error).
	// Consulta imediata/backdated cairia nisso — garante piso no futuro.
	if min := time.Now().Add(TelemedWindowAfter + 30*time.Minute); exp.Before(min) {
		exp = min
	}

	room, err := s.dailyCoSvc.CreateRoom(ctx, prefix, exp)
	if err != nil {
		return "", "", err
	}
	if err := s.db.WithContext(ctx).Model(&models.Appointment{}).
		Where("id = ?", apptID).
		Updates(map[string]any{
			"daily_room_url":  room.URL,
			"daily_room_name": room.Name,
		}).Error; err != nil {
		log.Printf("⚠️  [APPT] daily persist apt=%s: %v", apptID, err)
	}

	// Gera lobby token público (link /sala/[token] que email/WA inclui pra
	// paciente entrar sem login). Idempotente — Ensure reaproveita se existir.
	if s.telemedLobbySvc != nil {
		if _, err := s.telemedLobbySvc.EnsureTokenForAppointmentID(apptID); err != nil {
			log.Printf("⚠️  [APPT] telemed lobby token apt=%s: %v", apptID, err)
		}
	}
	return room.Name, room.URL, nil
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
	if appointment.ContinuumItemID != nil {
		cid := appointment.ContinuumItemID.String()
		resp.ContinuumItemID = &cid
	}
	if appointment.ConfirmedAt != nil {
		ca := appointment.ConfirmedAt.Format(time.RFC3339)
		resp.ConfirmedAt = &ca
	}
	if appointment.CheckedInAt != nil {
		ci := appointment.CheckedInAt.Format(time.RFC3339)
		resp.CheckedInAt = &ci
	}
	if appointment.StartedAt != nil {
		st := appointment.StartedAt.Format(time.RFC3339)
		resp.StartedAt = &st
	}
	if appointment.CompletedAt != nil {
		ca := appointment.CompletedAt.Format(time.RFC3339)
		resp.CompletedAt = &ca
	}
	if appointment.CancelledAt != nil {
		ca := appointment.CancelledAt.Format(time.RFC3339)
		resp.CancelledAt = &ca
	}
	if appointment.TelemedConsentAt != nil {
		tc := appointment.TelemedConsentAt.Format(time.RFC3339)
		resp.TelemedConsentAt = &tc
	}
	resp.TelemedConsentMode = appointment.TelemedConsentMode

	// Nested patient/doctor (preloaded). Apenas projeção mínima — nada sensível.
	if appointment.Patient.ID != uuid.Nil {
		resp.Patient = &dto.AppointmentParty{
			ID:    appointment.Patient.ID.String(),
			Name:  appointment.Patient.Name,
			Email: appointment.Patient.Email,
			Phone: appointment.Patient.Phone,
		}
	}
	if appointment.Doctor.ID != uuid.Nil {
		email := appointment.Doctor.Email
		resp.Doctor = &dto.AppointmentParty{
			ID:    appointment.Doctor.ID.String(),
			Name:  appointment.Doctor.Name,
			Email: &email,
		}
	}

	return resp
}
