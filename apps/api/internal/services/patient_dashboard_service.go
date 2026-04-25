package services

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// PatientDashboardService monta o agregado da home do portal do paciente.
//
// Filosofia: 1 endpoint, 1 round-trip, vários SELECTs leves em paralelo.
// Não duplicamos cálculos pesados (snapshots etc) — só lemos o que já existe.
type PatientDashboardService struct {
	db *gorm.DB
}

func NewPatientDashboardService(db *gorm.DB) *PatientDashboardService {
	return &PatientDashboardService{db: db}
}

// PatientDashboard é o payload da home.
type PatientDashboard struct {
	NextAppointment   *DashboardAppointment   `json:"nextAppointment,omitempty"`
	Continuum         *DashboardContinuum     `json:"continuum,omitempty"`
	LastScore         *DashboardScore         `json:"lastScore,omitempty"`
	BoxesCount        DashboardBoxesCount     `json:"boxesCount"`
	UnreadMessages    int                     `json:"unreadMessages"`
	PendingActions    []DashboardPendingAction `json:"pendingActions"`
}

type DashboardAppointment struct {
	ID                  uuid.UUID  `json:"id"`
	ScheduledAt         time.Time  `json:"scheduledAt"`
	DurationMinutes     int        `json:"durationMinutes"`
	Type                string     `json:"type"`
	Status              string     `json:"status"`
	DoctorID            uuid.UUID  `json:"doctorId"`
	DoctorName          string     `json:"doctorName"`
	IsTelemedicine      bool       `json:"isTelemedicine"`
	PatientConfirmedAt  *time.Time `json:"patientConfirmedAt,omitempty"`
	MinutesUntilStart   int64      `json:"minutesUntilStart"` // negativo se já passou (sala aberta)
}

type DashboardContinuum struct {
	ID                   uuid.UUID `json:"id"`
	TemplateName         string    `json:"templateName"`
	StartDate            time.Time `json:"startDate"`
	EndDate              time.Time `json:"endDate"`
	CurrentWeek          int       `json:"currentWeek"`
	TotalWeeks           int       `json:"totalWeeks"`
	ItemsCompleted       int       `json:"itemsCompleted"`
	ItemsTotal           int       `json:"itemsTotal"`
	ItemsLate            int       `json:"itemsLate"`
	NextItemTitle        *string   `json:"nextItemTitle,omitempty"`
	NextItemExpectedDate *time.Time `json:"nextItemExpectedDate,omitempty"`
}

type DashboardScore struct {
	ID                   uuid.UUID `json:"id"`
	Source               string    `json:"source"` // "light" ou "complete"
	CreatedAt            time.Time `json:"createdAt"`
	TotalScorePercentage float64   `json:"totalScorePercentage"`
	DeltaVsPrevious      *float64  `json:"deltaVsPrevious,omitempty"`
}

type DashboardBoxesCount struct {
	Preparing int `json:"preparing"`
	Shipped   int `json:"shipped"`
	Delivered int `json:"delivered"`
}

type DashboardPendingAction struct {
	Kind        string `json:"kind"` // "confirm_appointment", "lab_request_pending", etc
	Description string `json:"description"`
	Link        string `json:"link"`
}

// Build agrega tudo num payload só. Cada bloco falha silenciosamente
// (loga e retorna nil) pra que erro num bloco não derrube a home toda.
func (s *PatientDashboardService) Build(patientID uuid.UUID) (*PatientDashboard, error) {
	out := &PatientDashboard{
		PendingActions: []DashboardPendingAction{},
	}

	// Próxima consulta — primeiro scheduled/confirmed >= agora, ordenado asc.
	out.NextAppointment = s.nextAppointment(patientID)

	// Continuum ativo + agregados.
	out.Continuum = s.activeContinuum(patientID)

	// Último escore (Light OU Completo, o mais recente).
	out.LastScore = s.lastScore(patientID)

	// Boxes em fluxo (preparing/shipped/delivered).
	out.BoxesCount = s.boxesCount(patientID)

	// Pendências — gera com base nos blocos acima (consultas a confirmar etc).
	out.PendingActions = s.pendingActions(patientID, out.NextAppointment)

	// V1 — mensagens não-lidas vai zerar até Fase 6 implementar contador do lado paciente.
	out.UnreadMessages = 0

	return out, nil
}

// ============================================================
// Próxima consulta
// ============================================================

func (s *PatientDashboardService) nextAppointment(patientID uuid.UUID) *DashboardAppointment {
	var appt models.Appointment
	err := s.db.
		Where("patient_id = ? AND status IN ?", patientID, []string{"scheduled", "confirmed"}).
		Where("scheduled_at >= ?", time.Now().UTC().Add(-15*time.Minute)).
		Order("scheduled_at ASC").
		First(&appt).Error
	if err != nil {
		return nil
	}

	doctorName := ""
	var doctor models.User
	if err := s.db.Select("name").First(&doctor, "id = ?", appt.DoctorID).Error; err == nil {
		doctorName = doctor.Name
	}

	delta := time.Until(appt.ScheduledAt)
	return &DashboardAppointment{
		ID:                appt.ID,
		ScheduledAt:       appt.ScheduledAt,
		DurationMinutes:   appt.DurationMinutes,
		Type:              string(appt.Type),
		Status:            string(appt.Status),
		DoctorID:          appt.DoctorID,
		DoctorName:        doctorName,
		IsTelemedicine:    appt.Type == models.AppointmentTelemedicine,
		MinutesUntilStart: int64(delta.Minutes()),
	}
}

// ============================================================
// Continuum ativo
// ============================================================

func (s *PatientDashboardService) activeContinuum(patientID uuid.UUID) *DashboardContinuum {
	var c models.PatientContinuum
	err := s.db.
		Where("patient_id = ? AND status = ?", patientID, "active").
		Order("start_date DESC").
		First(&c).Error
	if err != nil {
		return nil
	}

	out := &DashboardContinuum{
		ID:           c.ID,
		StartDate:    c.StartDate,
		EndDate:      c.EndDate,
		TemplateName: ExtractTemplateName(c.TemplateSnapshot),
	}

	// Semanas
	totalDays := c.EndDate.Sub(c.StartDate).Hours() / 24
	out.TotalWeeks = int((totalDays + 6) / 7)
	since := time.Since(c.StartDate).Hours() / 24
	if since < 0 {
		since = 0
	}
	out.CurrentWeek = int(since/7) + 1
	if out.CurrentWeek > out.TotalWeeks && out.TotalWeeks > 0 {
		out.CurrentWeek = out.TotalWeeks
	}

	// Items: total, completed, late, próximo pendente
	type itemAgg struct {
		Status       string
		Count        int
		Title        *string
		ExpectedDate *time.Time
	}

	type itemRow struct {
		ID           uuid.UUID  `gorm:"column:id"`
		Title        string     `gorm:"column:title"`
		Status       string     `gorm:"column:status"`
		ExpectedDate time.Time  `gorm:"column:expected_date"`
	}
	var rows []itemRow
	_ = s.db.Raw(`
		SELECT id, title, status, expected_date
		FROM patient_continuum_items
		WHERE continuum_id = ?
		ORDER BY expected_date ASC
	`, c.ID).Scan(&rows).Error

	now := time.Now().UTC()
	out.ItemsTotal = len(rows)
	for _, r := range rows {
		switch r.Status {
		case "completed":
			out.ItemsCompleted++
		case "missed":
			out.ItemsLate++
		case "pending", "scheduled":
			if out.NextItemTitle == nil && !r.ExpectedDate.Before(now.AddDate(0, 0, -1)) {
				t := r.Title
				d := r.ExpectedDate
				out.NextItemTitle = &t
				out.NextItemExpectedDate = &d
			}
		}
	}

	return out
}

// ExtractTemplateName lê {"name":"..."} do JSONB sem precisar parsear tudo.
// Exportado pra reuso (handler do portal usa).
func ExtractTemplateName(snapshot []byte) string {
	if len(snapshot) == 0 {
		return ""
	}
	// Hack barato: se a struct mudar muito, vira proper unmarshal.
	type minimal struct {
		Name string `json:"name"`
	}
	var m minimal
	if err := json.Unmarshal(snapshot, &m); err == nil {
		return m.Name
	}
	return ""
}

// ============================================================
// Último escore (Light OU Completo)
// ============================================================

func (s *PatientDashboardService) lastScore(patientID uuid.UUID) *DashboardScore {
	type row struct {
		ID                   uuid.UUID `gorm:"column:id"`
		Source               string    `gorm:"column:source"`
		CreatedAt            time.Time `gorm:"column:created_at"`
		TotalScorePercentage float64   `gorm:"column:total_score_percentage"`
	}

	var rows []row
	// UNION dos dois — mais recentes primeiro, pegamos top 2 pra calcular delta.
	q := `
		SELECT s.id::uuid as id,
		       'complete' as source,
		       s.created_at,
		       s.total_score_percentage
		  FROM patient_score_snapshots s
		 WHERE s.patient_id = ? AND s.deleted_at IS NULL

		UNION ALL

		SELECT snap.id::uuid as id,
		       'light' as source,
		       snap.created_at,
		       snap.total_score_percentage
		  FROM anonymous_score_snapshots snap
		  JOIN anonymous_score_sessions sess ON sess.id = snap.session_id
		 WHERE sess.claimed_by_patient_id = ?

		ORDER BY created_at DESC
		LIMIT 2
	`
	if err := s.db.Raw(q, patientID, patientID).Scan(&rows).Error; err != nil || len(rows) == 0 {
		return nil
	}

	out := &DashboardScore{
		ID:                   rows[0].ID,
		Source:               rows[0].Source,
		CreatedAt:            rows[0].CreatedAt,
		TotalScorePercentage: rows[0].TotalScorePercentage,
	}
	if len(rows) > 1 {
		delta := rows[0].TotalScorePercentage - rows[1].TotalScorePercentage
		out.DeltaVsPrevious = &delta
	}
	return out
}

// ============================================================
// Boxes em fluxo
// ============================================================

func (s *PatientDashboardService) boxesCount(patientID uuid.UUID) DashboardBoxesCount {
	type row struct {
		Status string
		N      int
	}
	var rows []row
	_ = s.db.Raw(`
		SELECT b.status, COUNT(*)::int as n
		  FROM patient_continuum_boxes b
		  JOIN patient_continuums c ON c.id = b.continuum_id
		 WHERE c.patient_id = ?
		   AND b.status IN ('preparing','shipped','delivered')
		 GROUP BY b.status
	`, patientID).Scan(&rows).Error

	out := DashboardBoxesCount{}
	for _, r := range rows {
		switch r.Status {
		case "preparing":
			out.Preparing = r.N
		case "shipped":
			out.Shipped = r.N
		case "delivered":
			out.Delivered = r.N
		}
	}
	return out
}

// ============================================================
// Pendências
// ============================================================

func (s *PatientDashboardService) pendingActions(patientID uuid.UUID, next *DashboardAppointment) []DashboardPendingAction {
	out := []DashboardPendingAction{}

	// Confirmação de consulta
	if next != nil && next.PatientConfirmedAt == nil && next.MinutesUntilStart >= 0 && next.MinutesUntilStart < 60*24*2 {
		out = append(out, DashboardPendingAction{
			Kind:        "confirm_appointment",
			Description: "Confirme presença na sua próxima consulta",
			Link:        "/consultas/" + next.ID.String(),
		})
	}

	// Pedido de exame pendente
	var pendingLabs int64
	_ = s.db.Model(&models.LabRequest{}).
		Where("patient_id = ? AND status IN ?", patientID, []string{"draft", "issued"}).
		Count(&pendingLabs).Error
	if pendingLabs > 0 {
		out = append(out, DashboardPendingAction{
			Kind:        "lab_request_pending",
			Description: "Você tem pedido(s) de exame para coletar",
			Link:        "/exames",
		})
	}

	return out
}

