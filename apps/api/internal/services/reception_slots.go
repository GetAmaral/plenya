package services

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// Recepcionista virtual Fase 3 — provedor de horários reais.
//
// BuildUpcomingSlotsText busca os próximos horários disponíveis do médico da Consulta Plenya
// e devolve um bloco PT-BR pronto pro prompt do cérebro. Best-effort: qualquer falha retorna ""
// (o bot segue sem citar horários).
//
// Cache: a disponibilidade é a mesma para todas as conversas num dado instante. Como
// GenerateReceptionReply roda a cada copiloto e a cada candidato do job (1/min), e cada
// chamada faz até `receptionSlotsDays` chamadas ao CalendarSlotService (que pode bater no
// free/busy do Google), cacheamos o texto por `receptionSlotsCacheTTL` por médico.

const (
	receptionSlotsDays     = 10 // janela de busca (dias à frente)
	receptionSlotsMax      = 6  // máximo de horários ofertados
	receptionSlotsPerDay   = 2  // máximo por dia (variedade)
	receptionConsultMins   = 60 // duração da Consulta Plenya
	receptionSlotsTimezone = SlotPickerTimezone
	receptionSlotsCacheTTL = 10 * time.Minute
)

var ptWeekday = map[time.Weekday]string{
	time.Sunday:    "domingo",
	time.Monday:    "segunda",
	time.Tuesday:   "terça",
	time.Wednesday: "quarta",
	time.Thursday:  "quinta",
	time.Friday:    "sexta",
	time.Saturday:  "sábado",
}

type slotsCacheEntry struct {
	text      string
	expiresAt time.Time
}

var (
	slotsCacheMu sync.Mutex
	slotsCache   = map[uuid.UUID]slotsCacheEntry{}
)

// BuildUpcomingSlotsText devolve, por exemplo:
//
//	terça, 10/06 às 14:00
//	quarta, 11/06 às 09:00
//
// configuredDoctorID vazio = primeiro user com role doctor (ver defaultConsultDoctorID).
func BuildUpcomingSlotsText(ctx context.Context, db *gorm.DB, slotSvc *CalendarSlotService, configuredDoctorID string) string {
	if db == nil || slotSvc == nil {
		return ""
	}

	doctorID, ok := resolveConsultDoctorID(ctx, db, configuredDoctorID)
	if !ok {
		return ""
	}

	// Cache hit?
	now := time.Now()
	slotsCacheMu.Lock()
	if e, found := slotsCache[doctorID]; found && now.Before(e.expiresAt) {
		txt := e.text
		slotsCacheMu.Unlock()
		return txt
	}
	slotsCacheMu.Unlock()

	loc, err := time.LoadLocation(receptionSlotsTimezone)
	if err != nil {
		loc = time.UTC
	}

	dur := receptionConsultMins
	var lines []string
	for d := 0; d < receptionSlotsDays && len(lines) < receptionSlotsMax; d++ {
		day := now.AddDate(0, 0, d)
		slots, err := slotSvc.ListAvailable(ctx, ListSlotsParams{
			DoctorID:         doctorID,
			Date:             day,
			AppointmentType:  models.AppointmentInitialAssessment,
			OverrideDuration: &dur,
		})
		if err != nil || len(slots) == 0 {
			continue
		}
		perDay := 0
		for _, sl := range slots {
			if len(lines) >= receptionSlotsMax || perDay >= receptionSlotsPerDay {
				break
			}
			t := sl.StartUTC.In(loc)
			if t.Before(now) {
				continue
			}
			lines = append(lines, fmt.Sprintf("%s, %s às %s",
				ptWeekday[t.Weekday()], t.Format("02/01"), t.Format("15:04")))
			perDay++
		}
	}

	text := strings.Join(lines, "\n")
	slotsCacheMu.Lock()
	slotsCache[doctorID] = slotsCacheEntry{text: text, expiresAt: now.Add(receptionSlotsCacheTTL)}
	slotsCacheMu.Unlock()
	return text
}

// minToHHMM formata minutos-desde-meia-noite em "HH:MM".
func minToHHMM(m int) string {
	if m >= 1440 {
		m = 1439 // 24:00 vira 23:59 pra não confundir
	}
	return fmt.Sprintf("%02d:%02d", m/60, m%60)
}

// BuildBusinessHoursText resume o horário de FUNCIONAMENTO da clínica (working_hours do médico
// de consulta) num bloco curto pro prompt, pra IA saber quando a clínica atende. Ex.:
// "seg 08:00–12:00 e 14:00–18:00; ter 08:00–18:00; ...; dom fechado". "" se não houver dados.
func BuildBusinessHoursText(ctx context.Context, db *gorm.DB, configuredDoctorID string) string {
	doctorID, ok := resolveConsultDoctorID(ctx, db, configuredDoctorID)
	if !ok {
		return ""
	}
	var blocks []models.WorkingHours
	if err := db.WithContext(ctx).
		Where("doctor_id = ?", doctorID).
		Order("weekday ASC, start_minute ASC").
		Find(&blocks).Error; err != nil || len(blocks) == 0 {
		return ""
	}
	byDay := map[int][]string{}
	for _, b := range blocks {
		byDay[b.Weekday] = append(byDay[b.Weekday], fmt.Sprintf("%s–%s", minToHHMM(b.StartMinute), minToHHMM(b.EndMinute)))
	}
	names := []string{"dom", "seg", "ter", "qua", "qui", "sex", "sáb"}
	var parts []string
	for wd := 0; wd < 7; wd++ {
		if rngs, ok := byDay[wd]; ok {
			parts = append(parts, fmt.Sprintf("%s %s", names[wd], strings.Join(rngs, " e ")))
		} else {
			parts = append(parts, names[wd]+" fechado")
		}
	}
	return strings.Join(parts, "; ")
}

// resolveConsultDoctorID usa o médico configurado (RECEPTION_CONSULT_DOCTOR_ID) quando válido
// e ainda for um doctor; senão cai no primeiro user com role doctor (por nome). Evita oferecer
// a agenda do médico errado quando há mais de um doctor.
func resolveConsultDoctorID(ctx context.Context, db *gorm.DB, configuredDoctorID string) (uuid.UUID, bool) {
	if configuredDoctorID != "" {
		if id, err := uuid.Parse(configuredDoctorID); err == nil {
			var u models.User
			if err := db.WithContext(ctx).
				Where("id = ? AND roles @> ?::jsonb", id, `["doctor"]`).
				First(&u).Error; err == nil {
				return u.ID, true
			}
		}
	}
	var doctor models.User
	if err := db.WithContext(ctx).
		Where(`roles @> ?::jsonb`, `["doctor"]`).
		Order("name ASC").
		First(&doctor).Error; err != nil {
		return uuid.Nil, false
	}
	return doctor.ID, true
}
