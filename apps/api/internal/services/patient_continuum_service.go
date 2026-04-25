// Package services — PatientContinuumService gerencia o ciclo de vida de uma
// inscrição de paciente em um programa Continuum: Enroll (snapshot template +
// gera N items pré-agendados), GetByPatient, GetByID, ListItems, UpdateItem.
//
// Snapshot do template é guardado em PatientContinuum.TemplateSnapshot (JSONB)
// — mudanças posteriores no template oficial NÃO afetam inscrições já abertas.
package services

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

type PatientContinuumService struct {
	db *gorm.DB
}

func NewPatientContinuumService(db *gorm.DB) *PatientContinuumService {
	return &PatientContinuumService{db: db}
}

// EnrollParams são os parâmetros pra inscrever um paciente em um template.
type EnrollParams struct {
	PatientID           uuid.UUID
	TemplateID          uuid.UUID
	StartDate           time.Time
	CoordinatorDoctorID *uuid.UUID
	Notes               string
}

// Enroll cria PatientContinuum + N PatientContinuumItem + N PatientContinuumBox
// (pra items type=box) numa transação. Snapshot do template é serializado em
// JSONB. Datas calculadas a partir de StartDate + WeekOffset*7 + ExpectedOffsetDays.
func (s *PatientContinuumService) Enroll(p EnrollParams) (*models.PatientContinuum, error) {
	var template models.ContinuumTemplate
	if err := s.db.
		Preload("Items", func(db *gorm.DB) *gorm.DB {
			return db.Order("week_offset ASC, position ASC")
		}).
		First(&template, "id = ?", p.TemplateID).Error; err != nil {
		return nil, err
	}
	if template.Status != "active" {
		return nil, errors.New("template is not active")
	}

	// Snapshot serializado — congela o template no momento da inscrição.
	snapshotBytes, err := json.Marshal(template)
	if err != nil {
		return nil, err
	}

	startDate := p.StartDate.UTC().Truncate(24 * time.Hour)
	endDate := startDate.AddDate(0, 0, template.DurationWeeks*7)

	enrollment := models.PatientContinuum{
		PatientID:           p.PatientID,
		TemplateID:          p.TemplateID,
		TemplateSnapshot:    datatypes.JSON(snapshotBytes),
		Status:              models.ContinuumActive,
		StartDate:           startDate,
		EndDate:             endDate,
		CoordinatorDoctorID: p.CoordinatorDoctorID,
		Notes:               p.Notes,
	}

	err = s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&enrollment).Error; err != nil {
			return err
		}
		for _, ti := range template.Items {
			expected := startDate.AddDate(0, 0, ti.WeekOffset*7+ti.ExpectedOffsetDays)
			late := expected.AddDate(0, 0, ti.LateAfterDays)
			item := models.PatientContinuumItem{
				ContinuumID:   enrollment.ID,
				Type:          ti.Type,
				Specialty:     ti.Specialty,
				Title:         ti.Title,
				Description:   ti.Description,
				WeekOffset:    ti.WeekOffset,
				ExpectedDate:  expected,
				LateAfterDate: late,
				Status:        models.ContinuumItemPending,
				Position:      ti.Position,
			}
			if err := tx.Create(&item).Error; err != nil {
				return err
			}
			// Pra items type=box, criar PatientContinuumBox pré-snapshot.
			if ti.Type == models.ContinuumItemBox && ti.BoxTemplateID != nil {
				var box models.ContinuumBoxTemplate
				if err := tx.First(&box, "id = ?", *ti.BoxTemplateID).Error; err != nil {
					return err
				}
				pcBox := models.PatientContinuumBox{
					ContinuumItemID: item.ID,
					BoxTemplateID:   ti.BoxTemplateID,
					Name:            box.Name,
					Contents:        box.Contents,
					Status:          models.BoxPlanned,
				}
				if err := tx.Create(&pcBox).Error; err != nil {
					return err
				}
				// Linkar de volta no item.
				if err := tx.Model(&item).Update("box_id", pcBox.ID).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return s.GetByID(enrollment.ID)
}

// GetByPatient retorna todas as inscrições de um paciente (ativas, completadas, etc),
// ordenadas por start_date desc.
func (s *PatientContinuumService) GetByPatient(patientID uuid.UUID) ([]models.PatientContinuum, error) {
	var rows []models.PatientContinuum
	err := s.db.
		Where("patient_id = ?", patientID).
		Order("start_date DESC").
		Find(&rows).Error
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// GetByID retorna uma inscrição com items (ordenados) + paciente + coordenador.
func (s *PatientContinuumService) GetByID(id uuid.UUID) (*models.PatientContinuum, error) {
	var c models.PatientContinuum
	err := s.db.
		Preload("Patient").
		Preload("CoordinatorDoctor").
		Preload("Items", func(db *gorm.DB) *gorm.DB {
			return db.Order("expected_date ASC, position ASC")
		}).
		First(&c, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// UpdateItemStatus atualiza status + opcionalmente AppointmentID/CompletedAt
// (usado pelo handler quando equipe muda manualmente o status, e pela
// integração com AppointmentService — Fase 3).
type UpdateItemPatch struct {
	Status           *models.ContinuumItemStatus
	AppointmentID   *uuid.UUID
	CompletedAt     *time.Time
	CompletedRefType *string
	CompletedRefID  *uuid.UUID
}

func (s *PatientContinuumService) UpdateItem(itemID uuid.UUID, patch UpdateItemPatch) (*models.PatientContinuumItem, error) {
	var item models.PatientContinuumItem
	if err := s.db.First(&item, "id = ?", itemID).Error; err != nil {
		return nil, err
	}
	if patch.Status != nil {
		item.Status = *patch.Status
	}
	if patch.AppointmentID != nil {
		item.AppointmentID = patch.AppointmentID
	}
	if patch.CompletedAt != nil {
		item.CompletedAt = patch.CompletedAt
	}
	if patch.CompletedRefType != nil {
		item.CompletedRefType = patch.CompletedRefType
	}
	if patch.CompletedRefID != nil {
		item.CompletedRefID = patch.CompletedRefID
	}
	if err := s.db.Save(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

// Cancel marca a inscrição como cancelled (mantém histórico).
func (s *PatientContinuumService) Cancel(id uuid.UUID) error {
	return s.db.Model(&models.PatientContinuum{}).
		Where("id = ?", id).
		Update("status", models.ContinuumCancelled).Error
}

// ListActive retorna todas as inscrições ativas (usado pelo cron de atrasos
// e pelo dashboard da Fase 7).
func (s *PatientContinuumService) ListActive() ([]models.PatientContinuum, error) {
	var rows []models.PatientContinuum
	err := s.db.
		Preload("Patient").
		Where("status = ?", models.ContinuumActive).
		Order("start_date ASC").
		Find(&rows).Error
	return rows, err
}

// === Plano integrado (Fase 4) ===

// UpdateIntegratedPlan persiste novo conteúdo markdown E cria revisão imutável
// no histórico. Salva atualizadoPor + atualizadoEm no PatientContinuum.
//
// Concorrência: V1 usa "last write wins" — sem optimistic locking. Histórico de
// revisões garante que nada é perdido de fato (mesmo edição "sobrescrita" tem
// sua versão preservada). UI mostra "última edição há X" pra alertar conflito.
func (s *PatientContinuumService) UpdateIntegratedPlan(continuumID, updatedByID uuid.UUID, content string) (*models.PatientContinuum, error) {
	now := time.Now().UTC()
	err := s.db.Transaction(func(tx *gorm.DB) error {
		// Cria revisão (audit trail).
		rev := models.IntegratedPlanRevision{
			ContinuumID: continuumID,
			Content:     content,
			UpdatedByID: updatedByID,
		}
		if err := tx.Create(&rev).Error; err != nil {
			return err
		}
		// Atualiza campo current.
		return tx.Model(&models.PatientContinuum{}).
			Where("id = ?", continuumID).
			Updates(map[string]any{
				"integrated_plan_markdown":   content,
				"integrated_plan_updated_at": now,
				"integrated_plan_updated_by": updatedByID,
			}).Error
	})
	if err != nil {
		return nil, err
	}
	return s.GetByID(continuumID)
}

// ListPlanRevisions retorna o histórico cronológico (mais recente primeiro)
// das revisões do plano integrado, com info do autor.
func (s *PatientContinuumService) ListPlanRevisions(continuumID uuid.UUID) ([]models.IntegratedPlanRevision, error) {
	var rows []models.IntegratedPlanRevision
	err := s.db.
		Preload("UpdatedBy").
		Where("continuum_id = ?", continuumID).
		Order("created_at DESC").
		Find(&rows).Error
	return rows, err
}
