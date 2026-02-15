package services

import (
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"

	"github.com/google/uuid"
)

// DTOs for API
type CreateMethodDTO struct {
	Name        string  `json:"name" validate:"required,min=3,max=200"`
	ShortName   string  `json:"shortName" validate:"required,min=2,max=20"`
	Description *string `json:"description,omitempty"`
	Version     *string `json:"version,omitempty"`
	Color       *string `json:"color,omitempty" validate:"omitempty,hexcolor"`
	Order       int     `json:"order" validate:"gte=0,lte=9999"`
	IsDefault   bool    `json:"isDefault"`
}

type UpdateMethodDTO struct {
	Name        *string `json:"name,omitempty" validate:"omitempty,min=3,max=200"`
	ShortName   *string `json:"shortName,omitempty" validate:"omitempty,min=2,max=20"`
	Description *string `json:"description,omitempty"`
	Version     *string `json:"version,omitempty"`
	Color       *string `json:"color,omitempty" validate:"omitempty,hexcolor"`
	Order       *int    `json:"order,omitempty" validate:"omitempty,gte=0,lte=9999"`
	IsDefault   *bool   `json:"isDefault,omitempty"`
}

type CreateMethodLetterDTO struct {
	Code               string  `json:"code" validate:"required,min=1,max=10"`
	Name               string  `json:"name" validate:"required,min=3,max=300"`
	Description        *string `json:"description,omitempty"`
	ClinicalRelevance  *string `json:"clinicalRelevance,omitempty"`
	PatientExplanation *string `json:"patientExplanation,omitempty"`
	Conduct            *string `json:"conduct,omitempty"`
	Color              *string `json:"color,omitempty" validate:"omitempty,hexcolor"`
	Icon               *string `json:"icon,omitempty"`
	Order              int     `json:"order" validate:"gte=0,lte=9999"`
}

type UpdateMethodLetterDTO struct {
	Code               *string `json:"code,omitempty" validate:"omitempty,min=1,max=10"`
	Name               *string `json:"name,omitempty" validate:"omitempty,min=3,max=300"`
	Description        *string `json:"description,omitempty"`
	ClinicalRelevance  *string `json:"clinicalRelevance,omitempty"`
	PatientExplanation *string `json:"patientExplanation,omitempty"`
	Conduct            *string `json:"conduct,omitempty"`
	Color              *string `json:"color,omitempty" validate:"omitempty,hexcolor"`
	Icon               *string `json:"icon,omitempty"`
	Order              *int    `json:"order,omitempty" validate:"omitempty,gte=0,lte=9999"`
}

type CreateMethodPillarDTO struct {
	Name               string  `json:"name" validate:"required,min=3,max=300"`
	Description        *string `json:"description,omitempty"`
	ClinicalRelevance  *string `json:"clinicalRelevance,omitempty"`
	PatientExplanation *string `json:"patientExplanation,omitempty"`
	Conduct            *string `json:"conduct,omitempty"`
	Order              int     `json:"order" validate:"gte=0,lte=9999"`
}

type UpdateMethodPillarDTO struct {
	Name               *string `json:"name,omitempty" validate:"omitempty,min=3,max=300"`
	Description        *string `json:"description,omitempty"`
	ClinicalRelevance  *string `json:"clinicalRelevance,omitempty"`
	PatientExplanation *string `json:"patientExplanation,omitempty"`
	Conduct            *string `json:"conduct,omitempty"`
	Order              *int    `json:"order,omitempty" validate:"omitempty,gte=0,lte=9999"`
}

type AssignItemToPillarDTO struct {
	ItemID string `json:"itemId" validate:"required,uuid"`
}

type MethodService struct {
	repo *repository.MethodRepository
}

func NewMethodService(repo *repository.MethodRepository) *MethodService {
	return &MethodService{repo: repo}
}

// Method operations
func (s *MethodService) GetAllMethods() ([]models.Method, error) {
	return s.repo.GetAllMethods()
}

func (s *MethodService) GetMethodByID(methodID string) (*models.Method, error) {
	return s.repo.GetMethodByID(methodID)
}

func (s *MethodService) GetMethodWithTree(methodID string) (*models.Method, error) {
	return s.repo.GetMethodWithTree(methodID)
}

func (s *MethodService) GetAllMethodsWithTree() ([]models.Method, error) {
	return s.repo.GetAllMethodsWithTree()
}

func (s *MethodService) CreateMethod(dto *CreateMethodDTO) (*models.Method, error) {
	method := &models.Method{
		Name:        dto.Name,
		ShortName:   dto.ShortName,
		Description: dto.Description,
		Version:     dto.Version,
		Color:       dto.Color,
		Order:       dto.Order,
		IsDefault:   dto.IsDefault,
	}

	if err := s.repo.CreateMethod(method); err != nil {
		return nil, err
	}

	return method, nil
}

func (s *MethodService) UpdateMethod(methodID string, dto *UpdateMethodDTO) (*models.Method, error) {
	method, err := s.repo.GetMethodByID(methodID)
	if err != nil {
		return nil, err
	}

	if dto.Name != nil {
		method.Name = *dto.Name
	}
	if dto.ShortName != nil {
		method.ShortName = *dto.ShortName
	}
	if dto.Description != nil {
		method.Description = dto.Description
	}
	if dto.Version != nil {
		method.Version = dto.Version
	}
	if dto.Color != nil {
		method.Color = dto.Color
	}
	if dto.Order != nil {
		method.Order = *dto.Order
	}
	if dto.IsDefault != nil {
		method.IsDefault = *dto.IsDefault
	}

	if err := s.repo.UpdateMethod(method); err != nil {
		return nil, err
	}

	return method, nil
}

func (s *MethodService) DeleteMethod(methodID string) error {
	return s.repo.DeleteMethod(methodID)
}

// Method Letter operations
func (s *MethodService) GetMethodLetters(methodID string) ([]models.MethodLetter, error) {
	return s.repo.GetMethodLetters(methodID)
}

func (s *MethodService) GetMethodLetterByID(letterID string) (*models.MethodLetter, error) {
	return s.repo.GetMethodLetterByID(letterID)
}

func (s *MethodService) CreateMethodLetter(methodID string, dto *CreateMethodLetterDTO) (*models.MethodLetter, error) {
	methodUUID, err := uuid.Parse(methodID)
	if err != nil {
		return nil, err
	}

	letter := &models.MethodLetter{
		Code:               dto.Code,
		Name:               dto.Name,
		Description:        dto.Description,
		ClinicalRelevance:  dto.ClinicalRelevance,
		PatientExplanation: dto.PatientExplanation,
		Conduct:            dto.Conduct,
		Color:              dto.Color,
		Icon:               dto.Icon,
		Order:              dto.Order,
		MethodID:           methodUUID,
	}

	if err := s.repo.CreateMethodLetter(letter); err != nil {
		return nil, err
	}

	return letter, nil
}

func (s *MethodService) UpdateMethodLetter(letterID string, dto *UpdateMethodLetterDTO) (*models.MethodLetter, error) {
	letter, err := s.repo.GetMethodLetterByID(letterID)
	if err != nil {
		return nil, err
	}

	if dto.Code != nil {
		letter.Code = *dto.Code
	}
	if dto.Name != nil {
		letter.Name = *dto.Name
	}
	if dto.Description != nil {
		letter.Description = dto.Description
	}
	if dto.ClinicalRelevance != nil {
		letter.ClinicalRelevance = dto.ClinicalRelevance
	}
	if dto.PatientExplanation != nil {
		letter.PatientExplanation = dto.PatientExplanation
	}
	if dto.Conduct != nil {
		letter.Conduct = dto.Conduct
	}
	if dto.Color != nil {
		letter.Color = dto.Color
	}
	if dto.Icon != nil {
		letter.Icon = dto.Icon
	}
	if dto.Order != nil {
		letter.Order = *dto.Order
	}

	if err := s.repo.UpdateMethodLetter(letter); err != nil {
		return nil, err
	}

	return letter, nil
}

func (s *MethodService) DeleteMethodLetter(letterID string) error {
	return s.repo.DeleteMethodLetter(letterID)
}

// Method Pillar operations
func (s *MethodService) GetLetterPillars(letterID string) ([]models.MethodPillar, error) {
	return s.repo.GetLetterPillars(letterID)
}

func (s *MethodService) GetMethodPillarByID(pillarID string) (*models.MethodPillar, error) {
	return s.repo.GetMethodPillarByID(pillarID)
}

func (s *MethodService) CreateMethodPillar(letterID string, dto *CreateMethodPillarDTO) (*models.MethodPillar, error) {
	letterUUID, err := uuid.Parse(letterID)
	if err != nil {
		return nil, err
	}

	pillar := &models.MethodPillar{
		Name:               dto.Name,
		Description:        dto.Description,
		ClinicalRelevance:  dto.ClinicalRelevance,
		PatientExplanation: dto.PatientExplanation,
		Conduct:            dto.Conduct,
		Order:              dto.Order,
		LetterID:           letterUUID,
	}

	if err := s.repo.CreateMethodPillar(pillar); err != nil {
		return nil, err
	}

	return pillar, nil
}

func (s *MethodService) UpdateMethodPillar(pillarID string, dto *UpdateMethodPillarDTO) (*models.MethodPillar, error) {
	pillar, err := s.repo.GetMethodPillarByID(pillarID)
	if err != nil {
		return nil, err
	}

	if dto.Name != nil {
		pillar.Name = *dto.Name
	}
	if dto.Description != nil {
		pillar.Description = dto.Description
	}
	if dto.ClinicalRelevance != nil {
		pillar.ClinicalRelevance = dto.ClinicalRelevance
	}
	if dto.PatientExplanation != nil {
		pillar.PatientExplanation = dto.PatientExplanation
	}
	if dto.Conduct != nil {
		pillar.Conduct = dto.Conduct
	}
	if dto.Order != nil {
		pillar.Order = *dto.Order
	}

	if err := s.repo.UpdateMethodPillar(pillar); err != nil {
		return nil, err
	}

	return pillar, nil
}

func (s *MethodService) DeleteMethodPillar(pillarID string) error {
	return s.repo.DeleteMethodPillar(pillarID)
}

// Score Item assignment operations
func (s *MethodService) AssignScoreItemToPillar(pillarID string, dto *AssignItemToPillarDTO) error {
	return s.repo.AssignScoreItemToPillar(dto.ItemID, pillarID)
}

func (s *MethodService) UnassignScoreItemFromPillar(pillarID string, dto *AssignItemToPillarDTO) error {
	return s.repo.UnassignScoreItemFromPillar(dto.ItemID, pillarID)
}

func (s *MethodService) GetUnassignedScoreItems() ([]models.ScoreItem, error) {
	return s.repo.GetUnassignedScoreItems()
}

func (s *MethodService) GetAllScoreItemsWithPillars() ([]models.ScoreItem, error) {
	return s.repo.GetAllScoreItemsWithPillars()
}
