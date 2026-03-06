package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

var (
	ErrWorkoutPlanNotFound = errors.New("workout plan not found")
	ErrPublicCodeNotFound  = errors.New("workout plan not found for public code")
)

type WorkoutPlanService struct {
	db *gorm.DB
}

func NewWorkoutPlanService(db *gorm.DB) *WorkoutPlanService {
	return &WorkoutPlanService{db: db}
}

// Create cria um plano de treino com sessões e exercícios
func (s *WorkoutPlanService) Create(userID uuid.UUID, req *dto.CreateWorkoutPlanRequest) (*dto.WorkoutPlanResponse, error) {
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}
	if user.SelectedPatientID == nil {
		return nil, errors.New("no patient selected")
	}

	patientID := *user.SelectedPatientID
	if req.PatientID != "" {
		pid, err := uuid.Parse(req.PatientID)
		if err != nil {
			return nil, errors.New("invalid patient id")
		}
		if pid != patientID {
			return nil, errors.New("patient id does not match selected patient")
		}
	}

	plan := models.WorkoutPlan{
		PatientID:       patientID,
		CreatedByID:     userID,
		Name:            req.Name,
		Objective:       req.Objective,
		Intensity:       req.Intensity,
		DurationMinutes: req.DurationMinutes,
		WeeklyFrequency: req.WeeklyFrequency,
		IsActive:        true,
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&plan).Error; err != nil {
			return err
		}

		for _, sessReq := range req.Sessions {
			session := models.WorkoutPlanSession{
				PlanID: plan.ID,
				Name:   sessReq.Name,
				Order:  sessReq.Order,
			}
			if err := tx.Create(&session).Error; err != nil {
				return err
			}

			for _, exReq := range sessReq.Exercises {
				exerciseID, err := uuid.Parse(exReq.ExerciseID)
				if err != nil {
					return errors.New("invalid exercise id: " + exReq.ExerciseID)
				}
				exercise := models.WorkoutSessionExercise{
					SessionID:               session.ID,
					ExerciseID:              exerciseID,
					Phase:                   exReq.Phase,
					Order:                   exReq.Order,
					Sets:                    exReq.Sets,
					Reps:                    exReq.Reps,
					Cadence:                 exReq.Cadence,
					RestBetweenSetsSec:      exReq.RestBetweenSetsSec,
					RestBetweenExercisesSec: exReq.RestBetweenExercisesSec,
					Notes:                   exReq.Notes,
				}
				if err := tx.Create(&exercise).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return s.GetByID(plan.ID, userID)
}

// GetByID busca um plano por ID
func (s *WorkoutPlanService) GetByID(id, userID uuid.UUID) (*dto.WorkoutPlanResponse, error) {
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}
	if user.SelectedPatientID == nil {
		return nil, ErrWorkoutPlanNotFound
	}

	var plan models.WorkoutPlan
	if err := s.db.Where("id = ? AND patient_id = ?", id, *user.SelectedPatientID).
		Preload("Sessions", func(db *gorm.DB) *gorm.DB { return db.Order("\"order\" ASC") }).
		Preload("Sessions.Exercises", func(db *gorm.DB) *gorm.DB { return db.Order("phase ASC, \"order\" ASC") }).
		Preload("Sessions.Exercises.Exercise").
		First(&plan).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrWorkoutPlanNotFound
		}
		return nil, err
	}

	return planToDTO(&plan), nil
}

// GetByPublicCode busca um plano por código público (sem auth)
func (s *WorkoutPlanService) GetByPublicCode(code string) (*dto.WorkoutPlanResponse, error) {
	var plan models.WorkoutPlan
	if err := s.db.Where("public_code = ? AND is_active = true", code).
		Preload("Sessions", func(db *gorm.DB) *gorm.DB { return db.Order("\"order\" ASC") }).
		Preload("Sessions.Exercises", func(db *gorm.DB) *gorm.DB { return db.Order("phase ASC, \"order\" ASC") }).
		Preload("Sessions.Exercises.Exercise").
		First(&plan).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPublicCodeNotFound
		}
		return nil, err
	}

	return planToDTO(&plan), nil
}

// List lista planos do paciente selecionado
func (s *WorkoutPlanService) List(userID uuid.UUID, limit, offset int) ([]dto.WorkoutPlanResponse, error) {
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}
	if user.SelectedPatientID == nil {
		return []dto.WorkoutPlanResponse{}, nil
	}

	var plans []models.WorkoutPlan
	if err := s.db.Where("patient_id = ?", *user.SelectedPatientID).
		Preload("Sessions", func(db *gorm.DB) *gorm.DB { return db.Order("\"order\" ASC") }).
		Preload("Sessions.Exercises", func(db *gorm.DB) *gorm.DB { return db.Order("phase ASC, \"order\" ASC") }).
		Preload("Sessions.Exercises.Exercise").
		Order("created_at DESC").
		Limit(limit).Offset(offset).
		Find(&plans).Error; err != nil {
		return nil, err
	}

	result := make([]dto.WorkoutPlanResponse, len(plans))
	for i, p := range plans {
		result[i] = *planToDTO(&p)
	}
	return result, nil
}

// Update atualiza um plano de treino
func (s *WorkoutPlanService) Update(id, userID uuid.UUID, req *dto.UpdateWorkoutPlanRequest) (*dto.WorkoutPlanResponse, error) {
	var plan models.WorkoutPlan
	if err := s.db.Where("id = ?", id).First(&plan).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrWorkoutPlanNotFound
		}
		return nil, err
	}

	if req.Name != nil {
		plan.Name = *req.Name
	}
	if req.Objective != nil {
		plan.Objective = *req.Objective
	}
	if req.Intensity != nil {
		plan.Intensity = *req.Intensity
	}
	if req.DurationMinutes != nil {
		plan.DurationMinutes = *req.DurationMinutes
	}
	if req.WeeklyFrequency != nil {
		plan.WeeklyFrequency = *req.WeeklyFrequency
	}
	if req.IsActive != nil {
		plan.IsActive = *req.IsActive
	}

	if err := s.db.Save(&plan).Error; err != nil {
		return nil, err
	}

	return s.GetByID(plan.ID, userID)
}

// Delete faz soft delete de um plano
func (s *WorkoutPlanService) Delete(id uuid.UUID, userRole models.Role) error {
	if userRole != models.RoleAdmin {
		return ErrUnauthorized
	}
	result := s.db.Where("id = ?", id).Delete(&models.WorkoutPlan{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrWorkoutPlanNotFound
	}
	return nil
}

func planToDTO(p *models.WorkoutPlan) *dto.WorkoutPlanResponse {
	resp := &dto.WorkoutPlanResponse{
		ID:              p.ID.String(),
		PatientID:       p.PatientID.String(),
		CreatedByID:     p.CreatedByID.String(),
		Name:            p.Name,
		Objective:       p.Objective,
		Intensity:       p.Intensity,
		DurationMinutes: p.DurationMinutes,
		WeeklyFrequency: p.WeeklyFrequency,
		PublicCode:      p.PublicCode,
		HtmlContent:     p.HtmlContent,
		IsActive:        p.IsActive,
		DisplayTitle:    p.DisplayTitle,
		CreatedAt:       p.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       p.UpdatedAt.Format(time.RFC3339),
	}

	for _, sess := range p.Sessions {
		sessResp := dto.WorkoutPlanSessionResponse{
			ID:     sess.ID.String(),
			PlanID: sess.PlanID.String(),
			Name:   sess.Name,
			Order:  sess.Order,
		}
		for _, ex := range sess.Exercises {
			exResp := dto.WorkoutSessionExerciseResponse{
				ID:                      ex.ID.String(),
				SessionID:               ex.SessionID.String(),
				ExerciseID:              ex.ExerciseID.String(),
				Phase:                   ex.Phase,
				Order:                   ex.Order,
				Sets:                    ex.Sets,
				Reps:                    ex.Reps,
				Cadence:                 ex.Cadence,
				RestBetweenSetsSec:      ex.RestBetweenSetsSec,
				RestBetweenExercisesSec: ex.RestBetweenExercisesSec,
				Notes:                   ex.Notes,
			}
			if ex.Exercise.ID != uuid.Nil {
				exDTO := exerciseToDTO(&ex.Exercise)
				exResp.Exercise = &exDTO
			}
			sessResp.Exercises = append(sessResp.Exercises, exResp)
		}
		resp.Sessions = append(resp.Sessions, sessResp)
	}

	return resp
}
