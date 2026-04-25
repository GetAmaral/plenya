// Package services — CalendarSlotService implementa a lógica de slot picker
// pra agenda médica (Calendar V1, Bloco D).
//
// Algoritmo (intersection):
//
//	available_slots = working_hours(day)
//	                ∖ doctor_absences(day)
//	                ∖ existing_appointments(doctor, day)
//	                ∖ google_freebusy(doctor, day)
//
// Decisões:
//
//  1. Date é interpretado como "dia do médico" — TZ default America/Sao_Paulo.
//     V1 hardcoded; V2 lê User.Timezone (campo a ser adicionado).
//
//  2. SlotDuration vem de WorkingHours (default 30). OverrideDuration permite
//     mudar pro tipo de consulta (ex: initial_assessment 90min usa override).
//
//  3. Free/busy Google é OPCIONAL — se médico não tem CalendarCredential ativa,
//     pulamos esse step (não bloqueia). Se Google falhar (rate limit, network),
//     LOGAMOS warning e degradamos pra "sem free/busy" — pior caso é mostrar
//     slots que conflitam com Google pessoal (médico vê na própria agenda).
//
//  4. Free/busy consulta DOIS calendars: o dedicado "Plenya" (redundante mas
//     garante consistência) E o "primary" (eventos pessoais do médico).
//
//  5. Algoritmo O(slots × busy_intervals) — esperamos <100 slots/dia × ~10
//     busy intervals → trivial.
package services

import (
	"context"
	"errors"
	"log"
	"sort"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// SlotPickerTimezone — V1 hardcoded. V2: ler de User.Timezone.
const slotPickerTimezone = "America/Sao_Paulo"

type CalendarSlotService struct {
	db        *gorm.DB
	googleSvc *GoogleCalendarService
}

func NewCalendarSlotService(db *gorm.DB, googleSvc *GoogleCalendarService) *CalendarSlotService {
	return &CalendarSlotService{
		db:        db,
		googleSvc: googleSvc,
	}
}

// Slot — intervalo livre disponível pra agendamento.
type Slot struct {
	StartUTC time.Time `json:"startUtc"`
	EndUTC   time.Time `json:"endUtc"`
}

// ListSlotsParams — input do slot picker.
type ListSlotsParams struct {
	DoctorID         uuid.UUID
	Date             time.Time              // qualquer time no dia desejado (em qualquer TZ)
	AppointmentType  models.AppointmentType // pra computar duração default
	OverrideDuration *int                   // se setado, sobrescreve SlotDuration do WorkingHours
}

// defaultDurationFor retorna a duração default pra cada tipo (Plano V1).
// Caller pode passar OverrideDuration pra customizar.
func defaultDurationFor(t models.AppointmentType) int {
	switch t {
	case models.AppointmentInitialAssessment:
		return 90
	case models.AppointmentFollowUp:
		return 30
	case models.AppointmentTelemedicine:
		return 30
	case models.AppointmentProcedure:
		return 60
	case models.AppointmentResultsReview:
		return 45
	default:
		return 30
	}
}

// ListAvailable retorna slots livres pro médico no dia.
//
// Comportamento por step:
//  1. WorkingHours não cobre o weekday → retorna [] (médico não atende).
//  2. DoctorAbsence cobre o dia → retorna [] (férias/feriado).
//  3. Gera slots candidatos dentro de cada bloco WorkingHours.
//  4. Marca como busy slots que sobrepõem Appointments ativos.
//  5. (Opcional) Marca como busy slots que sobrepõem Google free/busy.
//  6. Retorna slots NÃO-busy ordenados ASC.
func (s *CalendarSlotService) ListAvailable(ctx context.Context, p ListSlotsParams) ([]Slot, error) {
	// Carrega TZ do médico (V1 hardcoded). Se loadLocation falhar (image sem
	// tzdata), cai pra UTC — slots vão estar em UTC e UI converte.
	loc, err := time.LoadLocation(slotPickerTimezone)
	if err != nil {
		loc = time.UTC
	}

	// Normaliza Date pra início do dia no TZ do médico.
	localDate := p.Date.In(loc)
	dayStart := time.Date(localDate.Year(), localDate.Month(), localDate.Day(), 0, 0, 0, 0, loc)
	dayEnd := dayStart.Add(24 * time.Hour)
	weekday := int(dayStart.Weekday()) // 0=Domingo

	// Resolve duração efetiva.
	effectiveDuration := defaultDurationFor(p.AppointmentType)
	if p.OverrideDuration != nil && *p.OverrideDuration > 0 {
		effectiveDuration = *p.OverrideDuration
	}

	// Step 1: WorkingHours pro weekday.
	var hours []models.WorkingHours
	if err := s.db.WithContext(ctx).
		Where("doctor_id = ? AND weekday = ?", p.DoctorID, weekday).
		Order("start_minute ASC").
		Find(&hours).Error; err != nil {
		return nil, err
	}
	if len(hours) == 0 {
		return []Slot{}, nil
	}

	// Step 2: DoctorAbsence sobrepondo o dia.
	// Sobreposição se: absence.StartAt < dayEnd AND absence.EndAt > dayStart.
	var absences []models.DoctorAbsence
	if err := s.db.WithContext(ctx).
		Where("doctor_id = ? AND start_at < ? AND end_at > ?", p.DoctorID, dayEnd, dayStart).
		Find(&absences).Error; err != nil {
		return nil, err
	}

	// Step 3: gera slots candidatos.
	candidates := make([]Slot, 0, 40)
	for _, h := range hours {
		blockStart := dayStart.Add(time.Duration(h.StartMinute) * time.Minute)
		blockEnd := dayStart.Add(time.Duration(h.EndMinute) * time.Minute)
		// Slot stride: usa effectiveDuration (não SlotDuration do row).
		// Justificativa: pra appointment tipo initial_assessment 90min num bloco
		// 9-12, queremos 9:00, 10:30 — não 9:00, 9:30, 10:00...
		stride := time.Duration(effectiveDuration) * time.Minute
		for t := blockStart; t.Add(stride).Compare(blockEnd) <= 0; t = t.Add(stride) {
			slot := Slot{
				StartUTC: t.UTC(),
				EndUTC:   t.Add(stride).UTC(),
			}
			if !slotOverlapsAbsences(slot, absences) {
				candidates = append(candidates, slot)
			}
		}
	}
	if len(candidates) == 0 {
		return []Slot{}, nil
	}

	// Step 4: Appointments existentes do médico nesse dia (status ativo).
	var existing []models.Appointment
	if err := s.db.WithContext(ctx).
		Where("doctor_id = ? AND scheduled_at >= ? AND scheduled_at < ? AND status NOT IN ?",
			p.DoctorID, dayStart, dayEnd, []models.AppointmentStatus{
				models.AppointmentCancelled,
				models.AppointmentNoShow,
			}).
		Find(&existing).Error; err != nil {
		return nil, err
	}
	busyFromAppointments := make([]busyInterval, 0, len(existing))
	for _, a := range existing {
		busyFromAppointments = append(busyFromAppointments, busyInterval{
			Start: a.ScheduledAt.UTC(),
			End:   a.ScheduledAt.Add(time.Duration(a.DurationMinutes) * time.Minute).UTC(),
		})
	}

	// Step 5: free/busy Google (best effort).
	googleBusy := s.fetchGoogleFreeBusy(ctx, p.DoctorID, dayStart.UTC(), dayEnd.UTC())

	// Step 6: filtra candidates contra busy intervals.
	allBusy := append(busyFromAppointments, googleBusy...)
	out := make([]Slot, 0, len(candidates))
	for _, c := range candidates {
		if !slotOverlapsBusy(c, allBusy) {
			out = append(out, c)
		}
	}

	// Garante ordenação ASC.
	sort.Slice(out, func(i, j int) bool {
		return out[i].StartUTC.Before(out[j].StartUTC)
	})

	return out, nil
}

// busyInterval é uso interno — equivale a FreeBusyInterval do Google service
// mas evita acoplar tipos.
type busyInterval struct {
	Start time.Time
	End   time.Time
}

// fetchGoogleFreeBusy faz lookup credential e chama Google. Em qualquer falha,
// retorna [] e logamos warning (degradação graciosa).
func (s *CalendarSlotService) fetchGoogleFreeBusy(ctx context.Context, doctorID uuid.UUID, startUTC, endUTC time.Time) []busyInterval {
	var cred models.CalendarCredential
	err := s.db.WithContext(ctx).
		Where("user_id = ? AND provider = ?", doctorID, models.CalendarProviderGoogle).
		First(&cred).Error
	if err != nil {
		// Sem credential ou outro erro DB: pula sem reclamar.
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("ℹ️  slot picker: google credential lookup error for doctor %s: %v", doctorID, err)
		}
		return nil
	}
	if !cred.IsActive() {
		return nil
	}

	access, err := s.googleSvc.GetValidAccessToken(ctx, doctorID)
	if err != nil {
		log.Printf("ℹ️  slot picker: google access token error for doctor %s: %v (degrading without freebusy)", doctorID, err)
		return nil
	}

	result, err := s.googleSvc.FreeBusy(ctx, access, FreeBusyInput{
		// Consulta calendar dedicado (eventos Plenya — redundante mas consistente)
		// E o primary (eventos pessoais do médico).
		CalendarIDs: []string{cred.DedicatedCalendarID, "primary"},
		StartUTC:    startUTC,
		EndUTC:      endUTC,
	})
	if err != nil {
		log.Printf("ℹ️  slot picker: google freebusy failed for doctor %s: %v (degrading)", doctorID, err)
		return nil
	}

	out := make([]busyInterval, 0, 16)
	for _, intervals := range result {
		for _, b := range intervals {
			out = append(out, busyInterval{Start: b.Start, End: b.End})
		}
	}
	return out
}

// slotOverlapsBusy retorna true se o slot intersecciona QUALQUER busy interval.
// Convenção: [start, end) — slot.End == busy.Start NÃO é overlap.
func slotOverlapsBusy(slot Slot, busy []busyInterval) bool {
	for _, b := range busy {
		if slot.StartUTC.Before(b.End) && slot.EndUTC.After(b.Start) {
			return true
		}
	}
	return false
}

// slotOverlapsAbsences — mesma lógica para absences (struct diferente).
func slotOverlapsAbsences(slot Slot, absences []models.DoctorAbsence) bool {
	for _, a := range absences {
		if slot.StartUTC.Before(a.EndAt) && slot.EndUTC.After(a.StartAt) {
			return true
		}
	}
	return false
}
