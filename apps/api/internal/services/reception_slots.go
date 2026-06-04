package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// Recepcionista virtual Fase 3 — provedor de horários reais.
//
// BuildUpcomingSlotsText busca os próximos horários disponíveis de um médico (Dr. Getúlio
// por padrão = primeiro user com role doctor) e devolve um bloco PT-BR pronto pro prompt do
// cérebro. Best-effort: qualquer falha retorna "" (o bot segue sem citar horários).

const (
	receptionSlotsDays     = 10 // janela de busca (dias à frente)
	receptionSlotsMax      = 6  // máximo de horários ofertados
	receptionSlotsPerDay   = 2  // máximo por dia (variedade)
	receptionConsultMins   = 60 // duração da Consulta Plenya
	receptionSlotsTimezone = SlotPickerTimezone
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

// BuildUpcomingSlotsText devolve, por exemplo:
//
//	terça, 10/06 às 14:00
//	quarta, 11/06 às 09:00
func BuildUpcomingSlotsText(ctx context.Context, db *gorm.DB, slotSvc *CalendarSlotService) string {
	if db == nil || slotSvc == nil {
		return ""
	}

	doctorID, ok := defaultConsultDoctorID(ctx, db)
	if !ok {
		return ""
	}

	loc, err := time.LoadLocation(receptionSlotsTimezone)
	if err != nil {
		loc = time.UTC
	}

	dur := receptionConsultMins
	var lines []string
	now := time.Now()
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
	return strings.Join(lines, "\n")
}

// defaultConsultDoctorID retorna o médico padrão da Consulta Plenya (primeiro user com role
// doctor, por nome). Calendar V1 não tem "médico default" formal; usamos isso até existir.
func defaultConsultDoctorID(ctx context.Context, db *gorm.DB) (uuid.UUID, bool) {
	var doctor models.User
	err := db.WithContext(ctx).
		Where(`roles @> ?::jsonb`, `["doctor"]`).
		Order("name ASC").
		First(&doctor).Error
	if err != nil {
		return uuid.Nil, false
	}
	return doctor.ID, true
}
