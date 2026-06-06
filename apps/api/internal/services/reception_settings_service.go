package services

// Serviço da configuração GLOBAL do Atendimento IA (singleton reception_settings) + a lógica
// da janela de horário do Auto e a resolução do modo global efetivo no instante.
// Plano: docs/emr/plano-atendimento-ia-global.md.

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	_ "time/tzdata" // embute o tzdata p/ LoadLocation funcionar sem depender do SO

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// brLocation é o fuso da clínica (janela de horário é interpretada nele).
var brLocation = func() *time.Location {
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		return time.UTC
	}
	return loc
}()

type ReceptionSettingsService struct {
	db *gorm.DB
}

func NewReceptionSettingsService(db *gorm.DB) *ReceptionSettingsService {
	return &ReceptionSettingsService{db: db}
}

// Get carrega o singleton (cria com defaults se não existir).
func (s *ReceptionSettingsService) Get(ctx context.Context) (*models.ReceptionSettings, error) {
	var st models.ReceptionSettings
	err := s.db.WithContext(ctx).First(&st, "id = ?", models.ReceptionSettingsID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		st = models.ReceptionSettings{
			ID:                     models.ReceptionSettingsID,
			BaselineMode:           models.ConversationAutomationOff,
			Schedule:               datatypes.JSON([]byte("{}")),
			DebounceSeconds:        30,
			PreviewSeconds:         10,
			CopilotFallbackMinutes: 10,
			OffAlertMinutes:        20,
		}
		if err := s.db.WithContext(ctx).Create(&st).Error; err != nil {
			return nil, err
		}
		return &st, nil
	}
	if err != nil {
		return nil, err
	}
	return &st, nil
}

// UpdateReceptionSettingsInput — campos opcionais (ponteiro = "mexer"); nil = manter.
type UpdateReceptionSettingsInput struct {
	BaselineMode           *string
	ScheduleEnabled        *bool
	Schedule               *models.WeeklySchedule
	DebounceSeconds        *int
	PreviewSeconds         *int
	CopilotFallbackMinutes *int
	OffAlertMinutes        *int
}

func validMode(m string) bool {
	return m == models.ConversationAutomationOff ||
		m == models.ConversationAutomationCopilot ||
		m == models.ConversationAutomationAuto
}

func (s *ReceptionSettingsService) Update(ctx context.Context, in UpdateReceptionSettingsInput, updatedBy uuid.UUID) (*models.ReceptionSettings, error) {
	st, err := s.Get(ctx)
	if err != nil {
		return nil, err
	}
	if in.BaselineMode != nil {
		if !validMode(*in.BaselineMode) {
			return nil, errors.New("baselineMode inválido (off|copilot|auto)")
		}
		st.BaselineMode = *in.BaselineMode
	}
	if in.ScheduleEnabled != nil {
		st.ScheduleEnabled = *in.ScheduleEnabled
	}
	if in.Schedule != nil {
		if err := validateSchedule(*in.Schedule); err != nil {
			return nil, err
		}
		b, _ := json.Marshal(*in.Schedule)
		st.Schedule = datatypes.JSON(b)
	}
	if in.DebounceSeconds != nil {
		if *in.DebounceSeconds < 0 || *in.DebounceSeconds > 3600 {
			return nil, errors.New("debounceSeconds fora do intervalo (0-3600)")
		}
		st.DebounceSeconds = *in.DebounceSeconds
	}
	if in.PreviewSeconds != nil {
		if *in.PreviewSeconds < 1 || *in.PreviewSeconds > 600 {
			return nil, errors.New("previewSeconds fora do intervalo (1-600)")
		}
		st.PreviewSeconds = *in.PreviewSeconds
	}
	if in.CopilotFallbackMinutes != nil {
		if *in.CopilotFallbackMinutes < 1 || *in.CopilotFallbackMinutes > 1440 {
			return nil, errors.New("copilotFallbackMinutes fora do intervalo (1-1440)")
		}
		st.CopilotFallbackMinutes = *in.CopilotFallbackMinutes
	}
	if in.OffAlertMinutes != nil {
		if *in.OffAlertMinutes < 1 || *in.OffAlertMinutes > 1440 {
			return nil, errors.New("offAlertMinutes fora do intervalo (1-1440)")
		}
		st.OffAlertMinutes = *in.OffAlertMinutes
	}
	st.UpdatedBy = &updatedBy
	if err := s.db.WithContext(ctx).Save(st).Error; err != nil {
		return nil, err
	}
	return st, nil
}

// EffectiveGlobalMode resolve o modo global no instante t: dentro da janela do Auto (se
// habilitada) → "auto"; senão → baseline.
func (s *ReceptionSettingsService) EffectiveGlobalMode(st *models.ReceptionSettings, t time.Time) string {
	if st.ScheduleEnabled {
		if sched, err := st.ParsedSchedule(); err == nil && isAutoActiveAt(sched, t) {
			return models.ConversationAutomationAuto
		}
	}
	return st.BaselineMode
}

// AutoActiveNow informa se a janela do Auto está ativa agora (para o banner de estado).
func (s *ReceptionSettingsService) AutoActiveNow(st *models.ReceptionSettings) bool {
	if !st.ScheduleEnabled {
		return false
	}
	sched, err := st.ParsedSchedule()
	if err != nil {
		return false
	}
	return isAutoActiveAt(sched, time.Now())
}

// ===== helpers da janela de horário =====

func weekdayKey(d time.Weekday) string {
	switch d {
	case time.Monday:
		return "mon"
	case time.Tuesday:
		return "tue"
	case time.Wednesday:
		return "wed"
	case time.Thursday:
		return "thu"
	case time.Friday:
		return "fri"
	case time.Saturday:
		return "sat"
	default:
		return "sun"
	}
}

// parseHHMM converte "HH:MM" em minutos do dia (0-1439).
func parseHHMM(s string) (int, bool) {
	parts := strings.SplitN(strings.TrimSpace(s), ":", 2)
	if len(parts) != 2 {
		return 0, false
	}
	h, err1 := strconv.Atoi(parts[0])
	m, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil || h < 0 || h > 23 || m < 0 || m > 59 {
		return 0, false
	}
	return h*60 + m, true
}

// isAutoActiveAt trata faixas que cruzam a meia-noite (start > end): a parte da noite conta no
// próprio dia (hm >= start) e a "madrugada" conta como spill do dia anterior (hm < end).
func isAutoActiveAt(sched models.WeeklySchedule, t time.Time) bool {
	t = t.In(brLocation)
	hm := t.Hour()*60 + t.Minute()
	today := weekdayKey(t.Weekday())
	yest := weekdayKey(t.AddDate(0, 0, -1).Weekday())

	for _, r := range sched[today] {
		st, ok1 := parseHHMM(r.Start)
		en, ok2 := parseHHMM(r.End)
		if !ok1 || !ok2 {
			continue
		}
		switch {
		case st < en:
			if hm >= st && hm < en {
				return true
			}
		case st > en: // cruza a meia-noite — parte da noite
			if hm >= st {
				return true
			}
		default: // start == end → dia todo
			return true
		}
	}
	for _, r := range sched[yest] {
		st, ok1 := parseHHMM(r.Start)
		en, ok2 := parseHHMM(r.End)
		if !ok1 || !ok2 {
			continue
		}
		if st > en && hm < en { // madrugada herdada do dia anterior
			return true
		}
	}
	return false
}

func validateSchedule(sched models.WeeklySchedule) error {
	valid := map[string]bool{"mon": true, "tue": true, "wed": true, "thu": true, "fri": true, "sat": true, "sun": true}
	for day, ranges := range sched {
		if !valid[day] {
			return fmt.Errorf("dia inválido: %s (use mon..sun)", day)
		}
		for _, r := range ranges {
			if _, ok := parseHHMM(r.Start); !ok {
				return fmt.Errorf("horário inválido em %s: %q", day, r.Start)
			}
			if _, ok := parseHHMM(r.End); !ok {
				return fmt.Errorf("horário inválido em %s: %q", day, r.End)
			}
		}
	}
	return nil
}
