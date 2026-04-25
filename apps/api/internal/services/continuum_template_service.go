// Package services — ContinuumTemplateService gerencia CRUD + clone dos
// templates de programa Continuum (Semestral, Anual, Trimestral...).
//
// Templates são editáveis a qualquer momento. Inscrições já abertas
// (PatientContinuum) NÃO são afetadas — elas guardam snapshot JSONB do
// template no momento da criação.
package services

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

type ContinuumTemplateService struct {
	db *gorm.DB
}

func NewContinuumTemplateService(db *gorm.DB) *ContinuumTemplateService {
	return &ContinuumTemplateService{db: db}
}

// List retorna todos os templates não-arquivados (ou todos, se includeArchived=true).
func (s *ContinuumTemplateService) List(includeArchived bool) ([]models.ContinuumTemplate, error) {
	var rows []models.ContinuumTemplate
	q := s.db.Model(&models.ContinuumTemplate{}).Order("created_at DESC")
	if !includeArchived {
		q = q.Where("status = ?", "active")
	}
	if err := q.Find(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}

// Get retorna um template com seus items ordenados por (week_offset, position).
func (s *ContinuumTemplateService) Get(id uuid.UUID) (*models.ContinuumTemplate, error) {
	var t models.ContinuumTemplate
	if err := s.db.
		Preload("Items", func(db *gorm.DB) *gorm.DB {
			return db.Order("week_offset ASC, position ASC")
		}).
		First(&t, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

// Create persiste um novo template e seus items numa transação.
func (s *ContinuumTemplateService) Create(payload dto.ContinuumTemplatePayload, createdBy uuid.UUID) (*models.ContinuumTemplate, error) {
	t := models.ContinuumTemplate{
		Name:            payload.Name,
		Description:     payload.Description,
		DurationWeeks:   payload.DurationWeeks,
		Status:          firstNonEmpty(payload.Status, "active"),
		CreatedByUserID: createdBy,
	}
	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&t).Error; err != nil {
			return err
		}
		return s.replaceItems(tx, t.ID, payload.Items)
	})
	if err != nil {
		return nil, err
	}
	return s.Get(t.ID)
}

// Update atualiza campos do template e substitui os items por inteiro
// (semântica "replace all" — mais simples que diff/upsert e suficiente pro UX
// de editor em tabela).
func (s *ContinuumTemplateService) Update(id uuid.UUID, payload dto.ContinuumTemplatePayload) (*models.ContinuumTemplate, error) {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		var t models.ContinuumTemplate
		if err := tx.First(&t, "id = ?", id).Error; err != nil {
			return err
		}
		t.Name = payload.Name
		t.Description = payload.Description
		t.DurationWeeks = payload.DurationWeeks
		if payload.Status != "" {
			t.Status = payload.Status
		}
		if err := tx.Save(&t).Error; err != nil {
			return err
		}
		return s.replaceItems(tx, t.ID, payload.Items)
	})
	if err != nil {
		return nil, err
	}
	return s.Get(id)
}

// Delete soft-deleta o template e seus items (cascade).
func (s *ContinuumTemplateService) Delete(id uuid.UUID) error {
	return s.db.Delete(&models.ContinuumTemplate{}, "id = ?", id).Error
}

// Clone cria um novo template copiando todos os items do source. Útil pra
// derivar Trimestral a partir de Semestral.
func (s *ContinuumTemplateService) Clone(sourceID uuid.UUID, newName string, createdBy uuid.UUID) (*models.ContinuumTemplate, error) {
	src, err := s.Get(sourceID)
	if err != nil {
		return nil, err
	}
	clone := models.ContinuumTemplate{
		Name:            newName,
		Description:     src.Description,
		DurationWeeks:   src.DurationWeeks,
		Status:          "active",
		CreatedByUserID: createdBy,
	}
	err = s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&clone).Error; err != nil {
			return err
		}
		for _, it := range src.Items {
			ni := models.ContinuumTemplateItem{
				TemplateID:         clone.ID,
				Type:               it.Type,
				Specialty:          it.Specialty,
				Title:              it.Title,
				Description:        it.Description,
				WeekOffset:         it.WeekOffset,
				ExpectedOffsetDays: it.ExpectedOffsetDays,
				LateAfterDays:      it.LateAfterDays,
				BoxTemplateID:      it.BoxTemplateID,
				Position:           it.Position,
			}
			if err := tx.Create(&ni).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return s.Get(clone.ID)
}

// replaceItems apaga todos os items do template e recria a partir do payload.
// Roda dentro de transação. Itens com payload.ID setado são ignorados (todos
// recriados — o ID do payload é só pra UX do frontend).
func (s *ContinuumTemplateService) replaceItems(tx *gorm.DB, templateID uuid.UUID, items []dto.ContinuumTemplateItemPayload) error {
	if err := tx.Where("template_id = ?", templateID).Delete(&models.ContinuumTemplateItem{}).Error; err != nil {
		return err
	}
	for _, p := range items {
		row, err := buildTemplateItem(templateID, p)
		if err != nil {
			return err
		}
		if err := tx.Create(row).Error; err != nil {
			return err
		}
	}
	return nil
}

func buildTemplateItem(templateID uuid.UUID, p dto.ContinuumTemplateItemPayload) (*models.ContinuumTemplateItem, error) {
	row := &models.ContinuumTemplateItem{
		TemplateID:         templateID,
		Type:               models.ContinuumItemType(p.Type),
		Title:              p.Title,
		Description:        p.Description,
		WeekOffset:         p.WeekOffset,
		ExpectedOffsetDays: p.ExpectedOffsetDays,
		LateAfterDays:      p.LateAfterDays,
		Position:           p.Position,
	}
	if p.Specialty != nil && *p.Specialty != "" {
		spec := models.ContinuumItemSpecialty(*p.Specialty)
		row.Specialty = &spec
	}
	if p.BoxTemplateID != nil && *p.BoxTemplateID != "" {
		bid, err := uuid.Parse(*p.BoxTemplateID)
		if err != nil {
			return nil, fmt.Errorf("invalid boxTemplateId: %w", err)
		}
		row.BoxTemplateID = &bid
	}
	if row.Type == models.ContinuumItemAppointment && row.Specialty == nil {
		return nil, errors.New("appointment item requires specialty")
	}
	if row.Type == models.ContinuumItemBox && row.BoxTemplateID == nil {
		return nil, errors.New("box item requires boxTemplateId")
	}
	return row, nil
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}
