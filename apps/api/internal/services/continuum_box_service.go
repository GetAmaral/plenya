// Package services — ContinuumBoxService gerencia o fluxo logístico de
// PatientContinuumBox: status (planned → preparing → shipped → delivered),
// tracking code, transportadora, endereço snapshot.
//
// Inscrição (Fase 2 Enroll) já cria o box com status=planned + contents copiado
// do template. Aqui só atualizamos status e tracking conforme equipe avança.
package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

type ContinuumBoxService struct {
	db *gorm.DB
}

func NewContinuumBoxService(db *gorm.DB) *ContinuumBoxService {
	return &ContinuumBoxService{db: db}
}

// BoxFilter — filtros pra listagem cross-paciente em /continuum/boxes.
type BoxFilter struct {
	Statuses []models.PatientContinuumBoxStatus
	Limit    int
	Offset   int
}

// BoxListItem — projeção pra UI: box + dados básicos do paciente + item.
type BoxListItem struct {
	ID              uuid.UUID                       `json:"id"`
	ContinuumItemID uuid.UUID                       `json:"continuumItemId"`
	Name            string                          `json:"name"`
	Contents        string                          `json:"contents"`
	Status          models.PatientContinuumBoxStatus `json:"status"`
	PreparedAt      *time.Time                      `json:"preparedAt,omitempty"`
	ShippedAt       *time.Time                      `json:"shippedAt,omitempty"`
	DeliveredAt     *time.Time                      `json:"deliveredAt,omitempty"`
	TrackingCode    *string                         `json:"trackingCode,omitempty"`
	Carrier         *string                         `json:"carrier,omitempty"`
	AddressSnapshot string                          `json:"addressSnapshot,omitempty"`
	Notes           string                          `json:"notes,omitempty"`
	CreatedAt       time.Time                       `json:"createdAt"`
	UpdatedAt       time.Time                       `json:"updatedAt"`

	PatientID    uuid.UUID `json:"patientId"`
	PatientName  string    `json:"patientName"`
	ExpectedDate time.Time `json:"expectedDate"`
	WeekOffset   int       `json:"weekOffset"`
	ContinuumID  uuid.UUID `json:"continuumId"`
}

// List retorna boxes cross-paciente com filtros — usado pela equipe de logística
// pra ver o que precisa preparar/despachar.
func (s *ContinuumBoxService) List(filter BoxFilter) ([]BoxListItem, error) {
	if filter.Limit <= 0 || filter.Limit > 500 {
		filter.Limit = 100
	}
	q := s.db.Table("patient_continuum_boxes pcb").
		Select(`pcb.id, pcb.continuum_item_id, pcb.name, pcb.contents, pcb.status,
			pcb.prepared_at, pcb.shipped_at, pcb.delivered_at,
			pcb.tracking_code, pcb.carrier, pcb.address_snapshot, pcb.notes,
			pcb.created_at, pcb.updated_at,
			pc.patient_id, p.name as patient_name,
			pci.expected_date, pci.week_offset,
			pc.id as continuum_id`).
		Joins("JOIN patient_continuum_items pci ON pci.id = pcb.continuum_item_id").
		Joins("JOIN patient_continuums pc ON pc.id = pci.continuum_id").
		Joins("JOIN patients p ON p.id = pc.patient_id").
		Where("pc.deleted_at IS NULL").
		Order("pci.expected_date ASC").
		Limit(filter.Limit).Offset(filter.Offset)
	if len(filter.Statuses) > 0 {
		q = q.Where("pcb.status IN ?", filter.Statuses)
	}
	rows := []BoxListItem{}
	if err := q.Scan(&rows).Error; err != nil {
		return nil, err
	}
	return rows, nil
}

// GetByID retorna um box específico (com mesma projeção de List, fonte única).
func (s *ContinuumBoxService) GetByID(id uuid.UUID) (*BoxListItem, error) {
	rows, err := s.List(BoxFilter{Limit: 1})
	if err != nil {
		return nil, err
	}
	for _, r := range rows {
		if r.ID == id {
			return &r, nil
		}
	}
	// Fallback: query direto (pra not-found semântico).
	var box models.PatientContinuumBox
	if err := s.db.First(&box, "id = ?", id).Error; err != nil {
		return nil, err
	}
	// Re-busca com join.
	var hits []BoxListItem
	if err := s.db.Table("patient_continuum_boxes pcb").
		Select(`pcb.id, pcb.continuum_item_id, pcb.name, pcb.contents, pcb.status,
			pcb.prepared_at, pcb.shipped_at, pcb.delivered_at,
			pcb.tracking_code, pcb.carrier, pcb.address_snapshot, pcb.notes,
			pcb.created_at, pcb.updated_at,
			pc.patient_id, p.name as patient_name,
			pci.expected_date, pci.week_offset, pc.id as continuum_id`).
		Joins("JOIN patient_continuum_items pci ON pci.id = pcb.continuum_item_id").
		Joins("JOIN patient_continuums pc ON pc.id = pci.continuum_id").
		Joins("JOIN patients p ON p.id = pc.patient_id").
		Where("pcb.id = ?", id).
		Scan(&hits).Error; err != nil {
		return nil, err
	}
	if len(hits) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &hits[0], nil
}

// UpdateBoxPatch — campos opcionalmente atualizáveis pelo handler.
type UpdateBoxPatch struct {
	Status       *models.PatientContinuumBoxStatus
	TrackingCode *string
	Carrier      *string
	Notes        *string
	Contents     *string
	Address      *string
}

// Update aplica patch + auto-preenche timestamps de transição (preparedAt/
// shippedAt/deliveredAt) na primeira vez que cada estado é atingido.
//
// Quando status vira delivered E o item Continuum vinculado ainda está pending,
// também marca o item como completed (paciente recebeu o box → marco cumprido).
func (s *ContinuumBoxService) Update(id uuid.UUID, patch UpdateBoxPatch) (*BoxListItem, error) {
	var box models.PatientContinuumBox
	if err := s.db.First(&box, "id = ?", id).Error; err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	if patch.Status != nil {
		newStatus := *patch.Status
		// Auto-fill timestamp de transição (idempotente — não sobrescreve).
		switch newStatus {
		case models.BoxPreparing:
			if box.PreparedAt == nil {
				box.PreparedAt = &now
			}
		case models.BoxShipped:
			if box.ShippedAt == nil {
				box.ShippedAt = &now
			}
			if box.PreparedAt == nil {
				box.PreparedAt = &now
			}
		case models.BoxDelivered:
			if box.DeliveredAt == nil {
				box.DeliveredAt = &now
			}
			if box.ShippedAt == nil {
				box.ShippedAt = &now
			}
			if box.PreparedAt == nil {
				box.PreparedAt = &now
			}
		}
		box.Status = newStatus
	}
	if patch.TrackingCode != nil {
		if *patch.TrackingCode == "" {
			box.TrackingCode = nil
		} else {
			tc := *patch.TrackingCode
			box.TrackingCode = &tc
		}
	}
	if patch.Carrier != nil {
		if *patch.Carrier == "" {
			box.Carrier = nil
		} else {
			c := *patch.Carrier
			box.Carrier = &c
		}
	}
	if patch.Notes != nil {
		box.Notes = *patch.Notes
	}
	if patch.Contents != nil {
		box.Contents = *patch.Contents
	}
	if patch.Address != nil {
		box.AddressSnapshot = *patch.Address
	}

	if err := s.db.Save(&box).Error; err != nil {
		return nil, err
	}

	// Propagação pro item Continuum: se virou delivered E o item ainda está
	// pending, marca como completed (box é evento "auto-cumprido" na entrega).
	if patch.Status != nil && *patch.Status == models.BoxDelivered {
		var item models.PatientContinuumItem
		if err := s.db.First(&item, "id = ?", box.ContinuumItemID).Error; err == nil {
			if item.Status == models.ContinuumItemPending ||
				item.Status == models.ContinuumItemMissed ||
				item.Status == models.ContinuumItemScheduled {
				refType := "box"
				_ = s.db.Model(&models.PatientContinuumItem{}).
					Where("id = ?", item.ID).
					Updates(map[string]any{
						"status":             models.ContinuumItemCompleted,
						"completed_at":       now,
						"completed_ref_type": refType,
						"completed_ref_id":   box.ID,
					}).Error
			}
		}
	}

	return s.GetByID(id)
}

// CountByStatus retorna agregação pra dashboard ("3 a despachar essa semana").
func (s *ContinuumBoxService) CountByStatus() (map[models.PatientContinuumBoxStatus]int64, error) {
	type row struct {
		Status models.PatientContinuumBoxStatus
		Count  int64
	}
	var rows []row
	err := s.db.Table("patient_continuum_boxes pcb").
		Select("pcb.status, COUNT(*) as count").
		Joins("JOIN patient_continuum_items pci ON pci.id = pcb.continuum_item_id").
		Joins("JOIN patient_continuums pc ON pc.id = pci.continuum_id").
		Where("pc.deleted_at IS NULL").
		Group("pcb.status").
		Scan(&rows).Error
	if err != nil {
		return nil, err
	}
	out := make(map[models.PatientContinuumBoxStatus]int64, len(rows))
	for _, r := range rows {
		out[r.Status] = r.Count
	}
	return out, nil
}

// GetByItemID retorna o box associado a um PatientContinuumItem (1:1 quando
// item.Type=box). Usado pelo frontend da timeline pra exibir status do box.
func (s *ContinuumBoxService) GetByItemID(itemID uuid.UUID) (*BoxListItem, error) {
	var hits []BoxListItem
	err := s.db.Table("patient_continuum_boxes pcb").
		Select(`pcb.id, pcb.continuum_item_id, pcb.name, pcb.contents, pcb.status,
			pcb.prepared_at, pcb.shipped_at, pcb.delivered_at,
			pcb.tracking_code, pcb.carrier, pcb.address_snapshot, pcb.notes,
			pcb.created_at, pcb.updated_at,
			pc.patient_id, p.name as patient_name,
			pci.expected_date, pci.week_offset, pc.id as continuum_id`).
		Joins("JOIN patient_continuum_items pci ON pci.id = pcb.continuum_item_id").
		Joins("JOIN patient_continuums pc ON pc.id = pci.continuum_id").
		Joins("JOIN patients p ON p.id = pc.patient_id").
		Where("pcb.continuum_item_id = ?", itemID).
		Scan(&hits).Error
	if err != nil {
		return nil, err
	}
	if len(hits) == 0 {
		return nil, errors.New("box not found for item")
	}
	return &hits[0], nil
}
