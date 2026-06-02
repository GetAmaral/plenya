package services

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

var ErrCarePlanItemNotFound = errors.New("care plan item not found")

// CarePlanService — plano de cuidado AGIR do paciente (escopado por patientID).
type CarePlanService struct {
	db *gorm.DB
}

func NewCarePlanService(db *gorm.DB) *CarePlanService { return &CarePlanService{db: db} }

// priorityOrder ordena high<medium<low no SQL.
const carePlanPriorityOrder = "CASE priority WHEN 'high' THEN 0 WHEN 'medium' THEN 1 ELSE 2 END"

// ListByPatient lista itens do plano (ordenados por letra A·G·I·R, prioridade, recência).
func (s *CarePlanService) ListByPatient(patientID uuid.UUID, includeInactive bool) ([]dto.CarePlanItemResponse, error) {
	q := s.db.Preload("AuthorUser").Where("patient_id = ?", patientID)
	if !includeInactive {
		q = q.Where("status = ?", models.CarePlanStatusActive)
	}
	var rows []models.CarePlanItem
	if err := q.
		Order("letter_code ASC").
		Order(carePlanPriorityOrder).
		Order("created_at DESC").
		Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]dto.CarePlanItemResponse, len(rows))
	for i := range rows {
		out[i] = *s.toDTO(&rows[i])
	}
	return out, nil
}

// ListActiveModels — itens ativos como models (consumido pelo relatório longitudinal).
func (s *CarePlanService) ListActiveModels(patientID uuid.UUID) ([]models.CarePlanItem, error) {
	var rows []models.CarePlanItem
	err := s.db.Where("patient_id = ? AND status = ?", patientID, models.CarePlanStatusActive).
		Order("letter_code ASC").
		Order(carePlanPriorityOrder).
		Find(&rows).Error
	return rows, err
}

func (s *CarePlanService) Create(patientID, authorID uuid.UUID, authorRole string, req *dto.CreateCarePlanItemRequest) (*dto.CarePlanItemResponse, error) {
	item := models.CarePlanItem{
		PatientID:      patientID,
		LetterCode:     strings.ToUpper(req.LetterCode),
		Recommendation: strings.TrimSpace(req.Recommendation),
		Rationale:      req.Rationale,
		Target:         req.Target,
		LabTestCode:    req.LabTestCode,
		Priority:       models.CarePlanPriorityMedium,
		Status:         models.CarePlanStatusActive,
		AuthorUserID:   authorID,
		AuthorRole:     authorRole,
	}
	if req.Priority != "" {
		item.Priority = models.CarePlanPriority(req.Priority)
	}
	if req.ScoreItemID != nil && *req.ScoreItemID != "" {
		sid, err := uuid.Parse(*req.ScoreItemID)
		if err != nil {
			return nil, errors.New("invalid scoreItemId")
		}
		item.ScoreItemID = &sid
	}
	if req.SourceSnapshotID != nil && *req.SourceSnapshotID != "" {
		ssid, err := uuid.Parse(*req.SourceSnapshotID)
		if err != nil {
			return nil, errors.New("invalid sourceSnapshotId")
		}
		item.SourceSnapshotID = &ssid
	}
	if err := s.db.Create(&item).Error; err != nil {
		return nil, err
	}
	return s.loadDTO(item.ID)
}

func (s *CarePlanService) Update(itemID, patientID uuid.UUID, req *dto.UpdateCarePlanItemRequest) (*dto.CarePlanItemResponse, error) {
	var item models.CarePlanItem
	if err := s.db.Where("id = ? AND patient_id = ?", itemID, patientID).First(&item).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrCarePlanItemNotFound
		}
		return nil, err
	}
	if req.LetterCode != nil {
		item.LetterCode = strings.ToUpper(*req.LetterCode)
	}
	if req.Recommendation != nil {
		item.Recommendation = strings.TrimSpace(*req.Recommendation)
	}
	if req.Rationale != nil {
		item.Rationale = req.Rationale
	}
	if req.Target != nil {
		item.Target = req.Target
	}
	if req.LabTestCode != nil {
		if *req.LabTestCode == "" {
			item.LabTestCode = nil
		} else {
			item.LabTestCode = req.LabTestCode
		}
	}
	if req.Priority != nil {
		item.Priority = models.CarePlanPriority(*req.Priority)
	}
	if req.Status != nil {
		item.Status = models.CarePlanItemStatus(*req.Status)
	}
	if err := s.db.Save(&item).Error; err != nil {
		return nil, err
	}
	return s.loadDTO(item.ID)
}

func (s *CarePlanService) Delete(itemID, patientID uuid.UUID) error {
	res := s.db.Where("id = ? AND patient_id = ?", itemID, patientID).Delete(&models.CarePlanItem{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return ErrCarePlanItemNotFound
	}
	return nil
}

func (s *CarePlanService) loadDTO(id uuid.UUID) (*dto.CarePlanItemResponse, error) {
	var item models.CarePlanItem
	if err := s.db.Preload("AuthorUser").First(&item, id).Error; err != nil {
		return nil, err
	}
	return s.toDTO(&item), nil
}

func (s *CarePlanService) toDTO(c *models.CarePlanItem) *dto.CarePlanItemResponse {
	r := &dto.CarePlanItemResponse{
		ID:             c.ID.String(),
		PatientID:      c.PatientID.String(),
		LetterCode:     c.LetterCode,
		Recommendation: c.Recommendation,
		Rationale:      c.Rationale,
		Target:         c.Target,
		LabTestCode:    c.LabTestCode,
		Priority:       string(c.Priority),
		Status:         string(c.Status),
		AuthorUserID:   c.AuthorUserID.String(),
		AuthorRole:     c.AuthorRole,
		CreatedAt:      c.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      c.UpdatedAt.Format(time.RFC3339),
	}
	if c.ScoreItemID != nil {
		v := c.ScoreItemID.String()
		r.ScoreItemID = &v
	}
	if c.SourceSnapshotID != nil {
		v := c.SourceSnapshotID.String()
		r.SourceSnapshotID = &v
	}
	if c.AuthorUser.ID != uuid.Nil {
		r.AuthorName = c.AuthorUser.Name
	}
	return r
}
