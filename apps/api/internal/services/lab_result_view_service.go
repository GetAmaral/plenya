package services

import (
	"github.com/google/uuid"
	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
)

type LabResultViewService struct {
	repo *repository.LabResultViewRepository
}

func NewLabResultViewService(repo *repository.LabResultViewRepository) *LabResultViewService {
	return &LabResultViewService{repo: repo}
}

// Create cria uma nova view
func (s *LabResultViewService) Create(req *dto.CreateLabResultViewRequest) (*models.LabResultView, error) {
	view := &models.LabResultView{
		Name:         req.Name,
		Description:  req.Description,
		DisplayOrder: req.DisplayOrder,
		IsActive:     true,
	}

	if err := s.repo.Create(view); err != nil {
		return nil, err
	}

	return view, nil
}

// GetByID busca uma view por ID
func (s *LabResultViewService) GetByID(id uuid.UUID, withItems bool) (*models.LabResultView, error) {
	return s.repo.GetByID(id, withItems)
}

// GetAll busca todas as views
func (s *LabResultViewService) GetAll(withItems bool, includeInactive bool) ([]models.LabResultView, error) {
	return s.repo.GetAll(withItems, includeInactive)
}

// Update atualiza uma view
func (s *LabResultViewService) Update(id uuid.UUID, req *dto.UpdateLabResultViewRequest) (*models.LabResultView, error) {
	view, err := s.repo.GetByID(id, false)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		view.Name = *req.Name
	}
	if req.Description != nil {
		view.Description = req.Description
	}
	if req.IsActive != nil {
		view.IsActive = *req.IsActive
	}
	if req.DisplayOrder != nil {
		view.DisplayOrder = *req.DisplayOrder
	}

	if err := s.repo.Update(view); err != nil {
		return nil, err
	}

	return view, nil
}

// Delete deleta uma view
func (s *LabResultViewService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}

// UpdateItems atualiza os items de uma view
func (s *LabResultViewService) UpdateItems(viewID uuid.UUID, req *dto.UpdateLabResultViewItemsRequest) error {
	// Parse UUIDs e criar models
	items := make([]models.LabResultViewItem, len(req.Items))
	for i, itemData := range req.Items {
		testID, err := uuid.Parse(itemData.LabTestDefinitionID)
		if err != nil {
			return err
		}

		items[i] = models.LabResultViewItem{
			LabTestDefinitionID: testID,
			Order:               itemData.Order,
		}
	}

	return s.repo.UpdateItems(viewID, items)
}

// Search busca views por nome ou descrição
func (s *LabResultViewService) Search(query string, withItems bool) ([]models.LabResultView, error) {
	return s.repo.Search(query, withItems)
}
