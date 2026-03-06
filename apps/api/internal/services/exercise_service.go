package services

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

var ErrExerciseNotFound = errors.New("exercise not found")

type ExerciseService struct {
	db *gorm.DB
}

func NewExerciseService(db *gorm.DB) *ExerciseService {
	return &ExerciseService{db: db}
}

// List busca exercícios com filtros JSONB
func (s *ExerciseService) List(search string, bodyParts, targetMuscles, equipments []string, limit, offset int) (*dto.ExerciseListResponse, error) {
	query := s.db.Model(&models.Exercise{}).Where("is_active = true")

	if search != "" {
		query = query.Where("(name ILIKE ? OR name_pt ILIKE ?)", "%"+search+"%", "%"+search+"%")
	}

	// JSONB array containment filters (@> operator)
	for _, bp := range bodyParts {
		query = query.Where("body_parts @> ?", `["`+bp+`"]`)
	}
	for _, tm := range targetMuscles {
		query = query.Where("target_muscles @> ?", `["`+tm+`"]`)
	}
	for _, eq := range equipments {
		query = query.Where("equipments @> ?", `["`+eq+`"]`)
	}

	var total int64
	query.Count(&total)

	var exercises []models.Exercise
	if err := query.Order("name ASC").Limit(limit).Offset(offset).Find(&exercises).Error; err != nil {
		return nil, err
	}

	items := make([]dto.ExerciseResponse, len(exercises))
	for i, e := range exercises {
		items[i] = exerciseToDTO(&e)
	}

	return &dto.ExerciseListResponse{
		Exercises: items,
		Total:     total,
	}, nil
}

// GetByID busca exercício por ID
func (s *ExerciseService) GetByID(id uuid.UUID) (*dto.ExerciseResponse, error) {
	var exercise models.Exercise
	if err := s.db.Where("id = ? AND is_active = true", id).First(&exercise).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrExerciseNotFound
		}
		return nil, err
	}
	resp := exerciseToDTO(&exercise)
	return &resp, nil
}

func exerciseToDTO(e *models.Exercise) dto.ExerciseResponse {
	return dto.ExerciseResponse{
		ID:               e.ID.String(),
		ExternalId:       e.ExternalId,
		Name:             e.Name,
		NamePt:           e.NamePt,
		BodyParts:        e.BodyParts,
		BodyPartsPt:      e.BodyPartsPt,
		TargetMuscles:    e.TargetMuscles,
		TargetMusclesPt:  e.TargetMusclesPt,
		Equipments:       e.Equipments,
		EquipmentsPt:     e.EquipmentsPt,
		SecondaryMuscles: e.SecondaryMuscles,
		Instructions:     e.Instructions,
		InstructionsPt:   e.InstructionsPt,
		GifUrl:           e.GifUrl,
		GifUrlFallback:   e.GifUrlFallback,
		IsActive:         e.IsActive,
	}
}
