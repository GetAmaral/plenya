package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/plenya/api/internal/models"
)

// PatientWorkoutsService expõe planos de treino + execução (sessions/logs)
// no escopo do paciente logado. Espelha o pattern de PatientLabsService:
// nunca aceita patientID como input — o middleware injeta.
type PatientWorkoutsService struct {
	db *gorm.DB
}

func NewPatientWorkoutsService(db *gorm.DB) *PatientWorkoutsService {
	return &PatientWorkoutsService{db: db}
}

// PlanSummary é um resumo enxuto pra lista no app.
type PlanSummary struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	DisplayTitle    string    `json:"displayTitle"`
	Objective       string    `json:"objective"`
	Intensity       string    `json:"intensity"`
	WeeklyFrequency int       `json:"weeklyFrequency"`
	IsActive        bool      `json:"isActive"`
	TotalSessions   int       `json:"totalSessions"`
	CreatedAt       time.Time `json:"createdAt"`
}

// ListPlans retorna planos ativos do paciente (mais recentes primeiro).
// Esconde inativos pra reduzir ruído no app — histórico fica na web.
func (s *PatientWorkoutsService) ListPlans(patientID uuid.UUID) ([]PlanSummary, error) {
	var rows []models.WorkoutPlan
	err := s.db.
		Where("patient_id = ? AND is_active = ?", patientID, true).
		Order("created_at DESC").
		Find(&rows).Error
	if err != nil {
		return nil, err
	}

	out := make([]PlanSummary, 0, len(rows))
	for i := range rows {
		var sessions int64
		s.db.Model(&models.WorkoutPlanSession{}).
			Where("plan_id = ?", rows[i].ID).Count(&sessions)

		out = append(out, PlanSummary{
			ID:              rows[i].ID,
			Name:            rows[i].Name,
			DisplayTitle:    rows[i].DisplayTitle,
			Objective:       string(rows[i].Objective),
			Intensity:       string(rows[i].Intensity),
			WeeklyFrequency: rows[i].WeeklyFrequency,
			IsActive:        rows[i].IsActive,
			TotalSessions:   int(sessions),
			CreatedAt:       rows[i].CreatedAt,
		})
	}
	return out, nil
}

// GetPlan devolve o plano com sessões + exercícios (incluindo o Exercise denormalizado
// pra UI poder mostrar GIF/instruções sem fetch extra).
func (s *PatientWorkoutsService) GetPlan(patientID, planID uuid.UUID) (*models.WorkoutPlan, error) {
	var plan models.WorkoutPlan
	err := s.db.
		Preload("Sessions.Exercises.Exercise").
		Where("id = ? AND patient_id = ?", planID, patientID).
		First(&plan).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("plano não encontrado")
	}
	if err != nil {
		return nil, err
	}
	return &plan, nil
}

// CreateSessionInput é o payload pra abrir uma execução de treino.
type CreateSessionInput struct {
	PlanID        uuid.UUID `json:"planId" validate:"required"`
	PlanSessionID uuid.UUID `json:"planSessionId" validate:"required"`
	ScheduledDate time.Time `json:"scheduledDate" validate:"required"`
}

// StartSession abre uma WorkoutSession (sem CompletedAt) — paciente vai
// adicionar logs durante a execução e finalizar em CompleteSession.
func (s *PatientWorkoutsService) StartSession(patientID uuid.UUID, in CreateSessionInput) (*models.WorkoutSession, error) {
	// Garante que o plan/session pertencem ao paciente
	var plan models.WorkoutPlan
	if err := s.db.Where("id = ? AND patient_id = ?", in.PlanID, patientID).First(&plan).Error; err != nil {
		return nil, errors.New("plano inválido")
	}
	var ps models.WorkoutPlanSession
	if err := s.db.Where("id = ? AND plan_id = ?", in.PlanSessionID, in.PlanID).First(&ps).Error; err != nil {
		return nil, errors.New("sessão inválida pro plano")
	}

	session := &models.WorkoutSession{
		PatientID:     patientID,
		PlanID:        in.PlanID,
		PlanSessionID: in.PlanSessionID,
		ScheduledDate: in.ScheduledDate,
	}
	if err := s.db.Create(session).Error; err != nil {
		return nil, err
	}
	return session, nil
}

// LogSetInput é um set executado.
type LogSetInput struct {
	PlanExerciseID uuid.UUID `json:"planExerciseId" validate:"required"`
	ExerciseID     uuid.UUID `json:"exerciseId" validate:"required"`
	SetNumber      int       `json:"setNumber" validate:"required,min=1,max=20"`
	Reps           *int      `json:"reps,omitempty"`
	Weight         *float64  `json:"weight,omitempty"`
	DurationSec    *int      `json:"durationSec,omitempty"`
	RPE            *int      `json:"rpe,omitempty"`
	Notes          *string   `json:"notes,omitempty"`
}

// LogSet adiciona/substitui o registro de UM set numa session existente.
// Race-safe via UNIQUE INDEX (session_id, plan_exercise_id, set_number) +
// ON CONFLICT DO UPDATE — dois clicks simultâneos não geram duplicata.
func (s *PatientWorkoutsService) LogSet(patientID, sessionID uuid.UUID, in LogSetInput) (*models.WorkoutSessionExerciseLog, error) {
	var sess models.WorkoutSession
	if err := s.db.Where("id = ? AND patient_id = ?", sessionID, patientID).First(&sess).Error; err != nil {
		return nil, errors.New("sessão não encontrada")
	}

	log := &models.WorkoutSessionExerciseLog{
		SessionID:      sessionID,
		PlanExerciseID: in.PlanExerciseID,
		ExerciseID:     in.ExerciseID,
		SetNumber:      in.SetNumber,
		Reps:           in.Reps,
		Weight:         in.Weight,
		DurationSec:    in.DurationSec,
		RPE:            in.RPE,
		Notes:          in.Notes,
	}
	err := s.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "session_id"}, {Name: "plan_exercise_id"}, {Name: "set_number"},
		},
		DoUpdates: clause.AssignmentColumns([]string{
			"reps", "weight", "duration_sec", "rpe", "notes", "updated_at",
		}),
	}).Create(log).Error
	if err != nil {
		return nil, err
	}
	// log.ID pode ter sido reescrito pelo ON CONFLICT — recarrega pra garantir
	// que devolvemos a row existente (com ID original).
	var out models.WorkoutSessionExerciseLog
	if err := s.db.Where(
		"session_id = ? AND plan_exercise_id = ? AND set_number = ?",
		sessionID, in.PlanExerciseID, in.SetNumber,
	).First(&out).Error; err != nil {
		return nil, err
	}
	return &out, nil
}

// CompleteSession marca CompletedAt e aceita notas finais opcionais.
func (s *PatientWorkoutsService) CompleteSession(patientID, sessionID uuid.UUID, notes *string) (*models.WorkoutSession, error) {
	var sess models.WorkoutSession
	if err := s.db.Where("id = ? AND patient_id = ?", sessionID, patientID).First(&sess).Error; err != nil {
		return nil, errors.New("sessão não encontrada")
	}
	now := time.Now().UTC()
	sess.CompletedAt = &now
	if notes != nil {
		sess.Notes = notes
	}
	if err := s.db.Save(&sess).Error; err != nil {
		return nil, err
	}
	return &sess, nil
}

// ListSessions retorna histórico do paciente (com logs preloaded).
// Ordenado por scheduled_date desc, default 30 últimas.
func (s *PatientWorkoutsService) ListSessions(patientID uuid.UUID, limit int) ([]models.WorkoutSession, error) {
	if limit <= 0 || limit > 100 {
		limit = 30
	}
	var rows []models.WorkoutSession
	err := s.db.
		Preload("Logs").
		Preload("PlanSession").
		Where("patient_id = ?", patientID).
		Order("scheduled_date DESC, created_at DESC").
		Limit(limit).
		Find(&rows).Error
	return rows, err
}

// GetSession devolve uma session específica (com logs).
func (s *PatientWorkoutsService) GetSession(patientID, sessionID uuid.UUID) (*models.WorkoutSession, error) {
	var sess models.WorkoutSession
	err := s.db.
		Preload("Logs.Exercise").
		Preload("PlanSession.Exercises.Exercise").
		Where("id = ? AND patient_id = ?", sessionID, patientID).
		First(&sess).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("sessão não encontrada")
	}
	if err != nil {
		return nil, err
	}
	return &sess, nil
}
