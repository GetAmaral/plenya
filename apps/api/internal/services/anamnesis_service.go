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
	ErrAnamnesisNotFound = errors.New("anamnesis not found")
)

type AnamnesisService struct {
	db *gorm.DB
}

func NewAnamnesisService(db *gorm.DB) *AnamnesisService {
	return &AnamnesisService{db: db}
}

// Create cria uma nova anamnese
func (s *AnamnesisService) Create(authorID uuid.UUID, req *dto.CreateAnamnesisRequest) (*dto.AnamnesisResponse, error) {
	// CRITICAL SECURITY: Get user's selected patient
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, authorID).Error; err != nil {
		return nil, err
	}

	// If no selected patient, cannot create anamnesis
	if user.SelectedPatientID == nil {
		return nil, errors.New("no patient selected - please select a patient first")
	}

	// Parse patient ID from request
	var patientID uuid.UUID
	if req.PatientID != "" {
		pid, err := uuid.Parse(req.PatientID)
		if err != nil {
			return nil, errors.New("invalid patient id")
		}
		// SECURITY: Validate that patientID matches selectedPatient
		if pid != *user.SelectedPatientID {
			return nil, errors.New("patient id does not match selected patient")
		}
		patientID = pid
	} else {
		// Auto-fill with selectedPatient
		patientID = *user.SelectedPatientID
	}

	// Verificar se o paciente existe
	var patient models.Patient
	if err := s.db.First(&patient, patientID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPatientNotFound
		}
		return nil, err
	}

	// Parse consultation date
	consultationDate, err := time.Parse(time.RFC3339, req.ConsultationDate)
	if err != nil {
		return nil, errors.New("invalid consultation date format, expected RFC3339")
	}

	// Parse anamnesis template ID if provided
	var anamnesisTemplateID *uuid.UUID
	if req.AnamnesisTemplateID != nil && *req.AnamnesisTemplateID != "" {
		templateID, err := uuid.Parse(*req.AnamnesisTemplateID)
		if err != nil {
			return nil, errors.New("invalid anamnesis template id")
		}
		anamnesisTemplateID = &templateID
	}

	// Criar anamnese
	anamnesis := models.Anamnesis{
		PatientID:           patientID,
		AuthorID:            authorID,
		AnamnesisTemplateID: anamnesisTemplateID,
		ConsultationDate:    consultationDate,
		Content:             req.Content,
		Summary:             req.Summary,
		Visibility:          models.AnamnesisVisibility(req.Visibility),
		Notes:               req.Notes,
	}

	// Start transaction
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(&anamnesis).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Create anamnesis items
	if len(req.Items) > 0 {
		items := make([]models.AnamnesisItem, len(req.Items))
		for i, itemReq := range req.Items {
			scoreItemID, err := uuid.Parse(itemReq.ScoreItemID)
			if err != nil {
				tx.Rollback()
				return nil, errors.New("invalid score item id")
			}

			items[i] = models.AnamnesisItem{
				AnamnesisID:  anamnesis.ID,
				ScoreItemID:  scoreItemID,
				TextValue:    itemReq.TextValue,
				NumericValue: itemReq.NumericValue,
				Order:        itemReq.Order,
			}
		}

		if err := tx.Create(&items).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		anamnesis.Items = items
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return s.toDTO(&anamnesis), nil
}

// GetByID busca uma anamnese por ID
func (s *AnamnesisService) GetByID(anamnesisID, userID uuid.UUID, userRole models.UserRole) (*dto.AnamnesisResponse, error) {
	// CRITICAL SECURITY: Get user's selected patient
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}

	// If no selected patient, cannot access any anamnesis
	if user.SelectedPatientID == nil {
		return nil, ErrAnamnesisNotFound
	}

	// ALWAYS filter by selectedPatient (all roles including admin)
	var anamnesis models.Anamnesis
	query := s.db.Preload("Items", func(db *gorm.DB) *gorm.DB {
		return db.Order("\"order\" ASC")
	}).Where("id = ?", anamnesisID).Where("patient_id = ?", *user.SelectedPatientID)

	if err := query.First(&anamnesis).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAnamnesisNotFound
		}
		return nil, err
	}

	return s.toDTO(&anamnesis), nil
}

// List lista anamneses com filtros
func (s *AnamnesisService) List(userID uuid.UUID, userRole models.UserRole, patientID *uuid.UUID, limit, offset int) ([]dto.AnamnesisResponse, error) {
	// CRITICAL SECURITY: Get user's selected patient
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}

	// If no selected patient, return empty list (security measure)
	if user.SelectedPatientID == nil {
		return []dto.AnamnesisResponse{}, nil
	}

	// ALWAYS filter by selectedPatient (all roles including admin)
	var anamneses []models.Anamnesis
	query := s.db.Preload("Items", func(db *gorm.DB) *gorm.DB {
		return db.Order("\"order\" ASC")
	}).Limit(limit).Offset(offset).Order("consultation_date DESC")
	query = query.Where("patient_id = ?", *user.SelectedPatientID)

	if err := query.Find(&anamneses).Error; err != nil {
		return nil, err
	}

	result := make([]dto.AnamnesisResponse, len(anamneses))
	for i, a := range anamneses {
		result[i] = *s.toDTO(&a)
	}

	return result, nil
}

// Update atualiza uma anamnese
func (s *AnamnesisService) Update(anamnesisID, authorID uuid.UUID, userRole models.UserRole, req *dto.UpdateAnamnesisRequest) (*dto.AnamnesisResponse, error) {
	var anamnesis models.Anamnesis
	query := s.db.Where("id = ?", anamnesisID)

	// Apenas o autor que criou pode editar (ou admin)
	if userRole != models.RoleAdmin {
		query = query.Where("author_id = ?", authorID)
	}

	if err := query.First(&anamnesis).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAnamnesisNotFound
		}
		return nil, err
	}

	// Start transaction
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Atualizar campos
	if req.AnamnesisTemplateID != nil {
		if *req.AnamnesisTemplateID == "" {
			anamnesis.AnamnesisTemplateID = nil
		} else {
			templateID, err := uuid.Parse(*req.AnamnesisTemplateID)
			if err != nil {
				tx.Rollback()
				return nil, errors.New("invalid anamnesis template id")
			}
			anamnesis.AnamnesisTemplateID = &templateID
		}
	}
	if req.ConsultationDate != nil {
		consultationDate, err := time.Parse(time.RFC3339, *req.ConsultationDate)
		if err != nil {
			tx.Rollback()
			return nil, errors.New("invalid consultation date format, expected RFC3339")
		}
		anamnesis.ConsultationDate = consultationDate
	}
	if req.Content != nil {
		anamnesis.Content = req.Content
	}
	if req.Summary != nil {
		anamnesis.Summary = req.Summary
	}
	if req.Visibility != nil {
		anamnesis.Visibility = models.AnamnesisVisibility(*req.Visibility)
	}
	if req.Notes != nil {
		anamnesis.Notes = req.Notes
	}

	if err := tx.Save(&anamnesis).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Update items (replace all)
	if req.Items != nil {
		// Delete old items
		if err := tx.Where("anamnesis_id = ?", anamnesisID).Delete(&models.AnamnesisItem{}).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		// Create new items
		if len(req.Items) > 0 {
			items := make([]models.AnamnesisItem, len(req.Items))
			for i, itemReq := range req.Items {
				scoreItemID, err := uuid.Parse(itemReq.ScoreItemID)
				if err != nil {
					tx.Rollback()
					return nil, errors.New("invalid score item id")
				}

				items[i] = models.AnamnesisItem{
					AnamnesisID:  anamnesis.ID,
					ScoreItemID:  scoreItemID,
					TextValue:    itemReq.TextValue,
					NumericValue: itemReq.NumericValue,
					Order:        itemReq.Order,
				}
			}

			if err := tx.Create(&items).Error; err != nil {
				tx.Rollback()
				return nil, err
			}

			anamnesis.Items = items
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	// Reload with items
	if err := s.db.Preload("Items", func(db *gorm.DB) *gorm.DB {
		return db.Order("\"order\" ASC")
	}).First(&anamnesis, anamnesisID).Error; err != nil {
		return nil, err
	}

	return s.toDTO(&anamnesis), nil
}

// Delete faz soft delete de uma anamnese
func (s *AnamnesisService) Delete(anamnesisID, authorID uuid.UUID, userRole models.UserRole) error {
	query := s.db.Where("id = ?", anamnesisID)

	// Apenas o autor que criou pode deletar (ou admin)
	if userRole != models.RoleAdmin {
		query = query.Where("author_id = ?", authorID)
	}

	result := query.Delete(&models.Anamnesis{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrAnamnesisNotFound
	}

	return nil
}

// toDTO converte Anamnesis para AnamnesisResponse
func (s *AnamnesisService) toDTO(anamnesis *models.Anamnesis) *dto.AnamnesisResponse {
	response := &dto.AnamnesisResponse{
		ID:               anamnesis.ID.String(),
		PatientID:        anamnesis.PatientID.String(),
		AuthorID:         anamnesis.AuthorID.String(),
		ConsultationDate: anamnesis.ConsultationDate.Format(time.RFC3339),
		Content:          anamnesis.Content,
		Summary:          anamnesis.Summary,
		Visibility:       string(anamnesis.Visibility),
		Notes:            anamnesis.Notes,
		CreatedAt:        anamnesis.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        anamnesis.UpdatedAt.Format(time.RFC3339),
	}

	// Add template ID if present
	if anamnesis.AnamnesisTemplateID != nil {
		templateID := anamnesis.AnamnesisTemplateID.String()
		response.AnamnesisTemplateID = &templateID
	}

	// Convert items
	if len(anamnesis.Items) > 0 {
		response.Items = make([]dto.AnamnesisItemResponse, len(anamnesis.Items))
		for i, item := range anamnesis.Items {
			response.Items[i] = dto.AnamnesisItemResponse{
				ID:           item.ID.String(),
				ScoreItemID:  item.ScoreItemID.String(),
				TextValue:    item.TextValue,
				NumericValue: item.NumericValue,
				Order:        item.Order,
				CreatedAt:    item.CreatedAt.Format(time.RFC3339),
				UpdatedAt:    item.UpdatedAt.Format(time.RFC3339),
			}
		}
	}

	return response
}
