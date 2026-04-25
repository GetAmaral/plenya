// Package services — ContinuumDashboardService alimenta o panorama da equipe
// em /continuum com 3 visões:
//
//   - PerPatient: lista inscrições ativas com progresso (% concluído, próximo
//     marco, # atrasados, semana atual). Visão "kanban de pacientes".
//   - PerWeek: agregação dia × especialidade × tipo (consultas/boxes) pra
//     uma semana específica — heatmap operacional.
//   - Alerts: items missed + items pending nos próximos N dias — lista
//     priorizada pra ação imediata.
package services

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

type ContinuumDashboardService struct {
	db *gorm.DB
}

func NewContinuumDashboardService(db *gorm.DB) *ContinuumDashboardService {
	return &ContinuumDashboardService{db: db}
}

// === Visão A: Por paciente ===

type PerPatientRow struct {
	ContinuumID         uuid.UUID  `json:"continuumId"`
	PatientID           uuid.UUID  `json:"patientId"`
	PatientName         string     `json:"patientName"`
	StartDate           time.Time  `json:"startDate"`
	EndDate             time.Time  `json:"endDate"`
	DurationWeeks       int        `json:"durationWeeks"`
	CurrentWeek         int        `json:"currentWeek"`
	TotalItems          int        `json:"totalItems"`
	CompletedItems      int        `json:"completedItems"`
	ScheduledItems      int        `json:"scheduledItems"`
	MissedItems         int        `json:"missedItems"`
	PendingItems        int        `json:"pendingItems"`
	NextItemTitle       *string    `json:"nextItemTitle,omitempty"`
	NextItemDate        *time.Time `json:"nextItemDate,omitempty"`
	CoordinatorDoctorID *uuid.UUID `json:"coordinatorDoctorId,omitempty"`
	CoordinatorName     *string    `json:"coordinatorName,omitempty"`
}

// PerPatient retorna uma linha por inscrição ativa com agregados pra exibição
// estilo kanban. Calcula currentWeek a partir de start_date vs now.
func (s *ContinuumDashboardService) PerPatient() ([]PerPatientRow, error) {
	now := time.Now().UTC()

	type baseRow struct {
		ContinuumID         uuid.UUID
		PatientID           uuid.UUID
		PatientName         string
		StartDate           time.Time
		EndDate             time.Time
		CoordinatorDoctorID *uuid.UUID
		CoordinatorName     *string
	}
	var bases []baseRow
	err := s.db.Table("patient_continuums pc").
		Select(`pc.id as continuum_id, pc.patient_id, p.name as patient_name,
			pc.start_date, pc.end_date,
			pc.coordinator_doctor_id, u.name as coordinator_name`).
		Joins("JOIN patients p ON p.id = pc.patient_id").
		Joins("LEFT JOIN users u ON u.id = pc.coordinator_doctor_id").
		Where("pc.status = ? AND pc.deleted_at IS NULL", models.ContinuumActive).
		Order("pc.start_date ASC").
		Scan(&bases).Error
	if err != nil {
		return nil, err
	}
	if len(bases) == 0 {
		return []PerPatientRow{}, nil
	}

	// Agregação de items por inscrição.
	ids := make([]uuid.UUID, 0, len(bases))
	for _, b := range bases {
		ids = append(ids, b.ContinuumID)
	}
	type itemAgg struct {
		ContinuumID uuid.UUID
		Status      models.ContinuumItemStatus
		Count       int
	}
	var aggs []itemAgg
	if err := s.db.Table("patient_continuum_items").
		Select("continuum_id, status, COUNT(*) as count").
		Where("continuum_id IN ?", ids).
		Group("continuum_id, status").
		Scan(&aggs).Error; err != nil {
		return nil, err
	}
	// Próximo item pendente/agendado por inscrição.
	type nextRow struct {
		ContinuumID  uuid.UUID
		Title        string
		ExpectedDate time.Time
	}
	var nexts []nextRow
	if err := s.db.Raw(`
		SELECT DISTINCT ON (continuum_id) continuum_id, title, expected_date
		  FROM patient_continuum_items
		 WHERE continuum_id IN ?
		   AND status IN ('pending','scheduled')
		   AND expected_date >= ?
		 ORDER BY continuum_id, expected_date ASC`, ids, now).
		Scan(&nexts).Error; err != nil {
		return nil, err
	}
	nextByID := make(map[uuid.UUID]nextRow, len(nexts))
	for _, n := range nexts {
		nextByID[n.ContinuumID] = n
	}

	// Indexa agregados por (continuumID, status).
	statusByID := make(map[uuid.UUID]map[models.ContinuumItemStatus]int)
	for _, a := range aggs {
		m := statusByID[a.ContinuumID]
		if m == nil {
			m = map[models.ContinuumItemStatus]int{}
			statusByID[a.ContinuumID] = m
		}
		m[a.Status] = a.Count
	}

	out := make([]PerPatientRow, 0, len(bases))
	for _, b := range bases {
		st := statusByID[b.ContinuumID]
		total := 0
		for _, c := range st {
			total += c
		}
		durationWeeks := int(b.EndDate.Sub(b.StartDate).Hours()/24/7 + 0.5)
		if durationWeeks < 1 {
			durationWeeks = 1
		}
		currentWeek := int(now.Sub(b.StartDate).Hours()/24/7) + 1
		if currentWeek < 1 {
			currentWeek = 1
		}
		if currentWeek > durationWeeks {
			currentWeek = durationWeeks
		}
		row := PerPatientRow{
			ContinuumID:         b.ContinuumID,
			PatientID:           b.PatientID,
			PatientName:         b.PatientName,
			StartDate:           b.StartDate,
			EndDate:             b.EndDate,
			DurationWeeks:       durationWeeks,
			CurrentWeek:         currentWeek,
			TotalItems:          total,
			CompletedItems:      st[models.ContinuumItemCompleted],
			ScheduledItems:      st[models.ContinuumItemScheduled],
			MissedItems:         st[models.ContinuumItemMissed],
			PendingItems:        st[models.ContinuumItemPending],
			CoordinatorDoctorID: b.CoordinatorDoctorID,
			CoordinatorName:     b.CoordinatorName,
		}
		if n, ok := nextByID[b.ContinuumID]; ok {
			t := n.Title
			d := n.ExpectedDate
			row.NextItemTitle = &t
			row.NextItemDate = &d
		}
		out = append(out, row)
	}
	return out, nil
}

// === Visão B: Por semana ===

type PerWeekItem struct {
	ID            uuid.UUID                  `json:"id"`
	ContinuumID   uuid.UUID                  `json:"continuumId"`
	Type          models.ContinuumItemType   `json:"type"`
	Specialty     *models.ContinuumItemSpecialty `json:"specialty,omitempty"`
	Title         string                     `json:"title"`
	ExpectedDate  time.Time                  `json:"expectedDate"`
	Status        models.ContinuumItemStatus `json:"status"`
	PatientID     uuid.UUID                  `json:"patientId"`
	PatientName   string                     `json:"patientName"`
	AppointmentID *uuid.UUID                 `json:"appointmentId,omitempty"`
	BoxID         *uuid.UUID                 `json:"boxId,omitempty"`
}

// PerWeek retorna todos items na janela [from, from+7d) ordenados por data.
// Frontend agrupa por dia + especialidade pra heatmap visual.
func (s *ContinuumDashboardService) PerWeek(weekStart time.Time) ([]PerWeekItem, error) {
	weekEnd := weekStart.AddDate(0, 0, 7)
	var rows []PerWeekItem
	err := s.db.Table("patient_continuum_items pci").
		Select(`pci.id, pci.continuum_id, pci.type, pci.specialty, pci.title,
			pci.expected_date, pci.status,
			pc.patient_id, p.name as patient_name,
			pci.appointment_id, pci.box_id`).
		Joins("JOIN patient_continuums pc ON pc.id = pci.continuum_id").
		Joins("JOIN patients p ON p.id = pc.patient_id").
		Where("pc.status = ? AND pc.deleted_at IS NULL", models.ContinuumActive).
		Where("pci.expected_date >= ? AND pci.expected_date < ?", weekStart, weekEnd).
		Order("pci.expected_date ASC").
		Scan(&rows).Error
	return rows, err
}

// === Visão C: Alertas ===

type AlertRow struct {
	ID            uuid.UUID                  `json:"id"`
	ContinuumID   uuid.UUID                  `json:"continuumId"`
	PatientID     uuid.UUID                  `json:"patientId"`
	PatientName   string                     `json:"patientName"`
	Type          models.ContinuumItemType   `json:"type"`
	Specialty     *models.ContinuumItemSpecialty `json:"specialty,omitempty"`
	Title         string                     `json:"title"`
	Status        models.ContinuumItemStatus `json:"status"`
	ExpectedDate  time.Time                  `json:"expectedDate"`
	LateAfterDate time.Time                  `json:"lateAfterDate"`
	Severity      string                     `json:"severity"` // missed | due-soon
	AppointmentID *uuid.UUID                 `json:"appointmentId,omitempty"`
}

// Alerts retorna items que exigem ação imediata: missed primeiro (severity
// alta), depois pending nos próximos `dueSoonDays` dias (severity média).
// Limit alto pra dashboard mas paginar quando passar 200.
func (s *ContinuumDashboardService) Alerts(dueSoonDays int) ([]AlertRow, error) {
	if dueSoonDays <= 0 {
		dueSoonDays = 7
	}
	now := time.Now().UTC()
	cutoff := now.AddDate(0, 0, dueSoonDays)
	var rows []AlertRow
	err := s.db.Table("patient_continuum_items pci").
		Select(`pci.id, pci.continuum_id, pc.patient_id, p.name as patient_name,
			pci.type, pci.specialty, pci.title, pci.status,
			pci.expected_date, pci.late_after_date,
			CASE WHEN pci.status = 'missed' THEN 'missed' ELSE 'due-soon' END AS severity,
			pci.appointment_id`).
		Joins("JOIN patient_continuums pc ON pc.id = pci.continuum_id").
		Joins("JOIN patients p ON p.id = pc.patient_id").
		Where("pc.status = ? AND pc.deleted_at IS NULL", models.ContinuumActive).
		Where(`(pci.status = 'missed') OR (pci.status = 'pending' AND pci.expected_date <= ?)`, cutoff).
		Order(`CASE WHEN pci.status = 'missed' THEN 0 ELSE 1 END, pci.expected_date ASC`).
		Limit(200).
		Scan(&rows).Error
	return rows, err
}
