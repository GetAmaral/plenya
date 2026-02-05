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
	ErrAnamnesisNotFound   = errors.New("anamnesis not found")
	ErrAnamnesisRestricted = errors.New("you don't have permission to view this anamnesis content")
)

type AnamnesisService struct {
	db *gorm.DB
}

func NewAnamnesisService(db *gorm.DB) *AnamnesisService {
	return &AnamnesisService{db: db}
}

// CanViewFullContent verifica se o usuário tem permissão para ver o conteúdo completo da anamnese
func (s *AnamnesisService) CanViewFullContent(anamnesis *models.Anamnesis, userID uuid.UUID, userRole models.Role) bool {
	// Admin sempre tem acesso total
	if userRole == models.RoleAdmin {
		return true
	}

	switch anamnesis.Visibility {
	case models.VisibilityAll:
		return true
	case models.VisibilityAuthorOnly:
		return anamnesis.AuthorID == userID
	case models.VisibilityMedicalOnly:
		return userRole == models.RoleDoctor
	case models.VisibilityPsychOnly:
		// TODO: Adicionar RolePsychologist quando implementado
		// Por enquanto, ninguém além de admin pode ver (admin já retornou true acima)
		return false
	default:
		return false
	}
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
		ContentHtml:         req.ContentHtml,
		Summary:             req.Summary,
		SummaryHtml:         req.SummaryHtml,
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

	return s.toDTO(&anamnesis, false), nil
}

// GetByID busca uma anamnese por ID
func (s *AnamnesisService) GetByID(anamnesisID, userID uuid.UUID, userRole models.Role) (*dto.AnamnesisResponse, error) {
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
	query := s.db.
		Preload("Author").
		Preload("AnamnesisTemplate").
		Preload("Items", func(db *gorm.DB) *gorm.DB {
			return db.Order("\"order\" ASC")
		}).
		Where("id = ?", anamnesisID).
		Where("patient_id = ?", *user.SelectedPatientID)

	if err := query.First(&anamnesis).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAnamnesisNotFound
		}
		return nil, err
	}

	// Check if user has permission to view full content
	// For GetByID (detail/edit), deny access if restricted
	if !s.CanViewFullContent(&anamnesis, userID, userRole) {
		return nil, ErrAnamnesisRestricted
	}

	return s.toDTO(&anamnesis, false), nil
}

// List lista anamneses com filtros
func (s *AnamnesisService) List(userID uuid.UUID, userRole models.Role, patientID *uuid.UUID, limit, offset int) ([]dto.AnamnesisResponse, error) {
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
	query := s.db.
		Preload("Author").
		Preload("AnamnesisTemplate").
		Preload("Items", func(db *gorm.DB) *gorm.DB {
			return db.Order("\"order\" ASC").
				Preload("ScoreItem.Subgroup.Group").
				Preload("ScoreItem.Levels", func(db *gorm.DB) *gorm.DB {
					return db.Order("level ASC")
				})
		}).
		Limit(limit).
		Offset(offset).
		Order("consultation_date DESC").
		Where("patient_id = ?", *user.SelectedPatientID)

	if err := query.Find(&anamneses).Error; err != nil {
		return nil, err
	}

	result := make([]dto.AnamnesisResponse, len(anamneses))
	for i, a := range anamneses {
		// Check if content should be restricted
		restricted := !s.CanViewFullContent(&a, userID, userRole)
		result[i] = *s.toDTO(&a, restricted)
	}

	return result, nil
}

// Update atualiza uma anamnese
func (s *AnamnesisService) Update(anamnesisID, authorID uuid.UUID, userRole models.Role, req *dto.UpdateAnamnesisRequest) (*dto.AnamnesisResponse, error) {
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
	if req.ContentHtml != nil {
		anamnesis.ContentHtml = req.ContentHtml
	}
	if req.Summary != nil {
		anamnesis.Summary = req.Summary
	}
	if req.SummaryHtml != nil {
		anamnesis.SummaryHtml = req.SummaryHtml
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

	return s.toDTO(&anamnesis, false), nil
}

// Delete faz soft delete de uma anamnese
func (s *AnamnesisService) Delete(anamnesisID, authorID uuid.UUID, userRole models.Role) error {
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
// Se restricted=true, oculta conteúdo sensível
func (s *AnamnesisService) toDTO(anamnesis *models.Anamnesis, restricted bool) *dto.AnamnesisResponse {
	response := &dto.AnamnesisResponse{
		ID:               anamnesis.ID.String(),
		PatientID:        anamnesis.PatientID.String(),
		AuthorID:         anamnesis.AuthorID.String(),
		ConsultationDate: anamnesis.ConsultationDate.Format(time.RFC3339),
		Visibility:       string(anamnesis.Visibility),
		CreatedAt:        anamnesis.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        anamnesis.UpdatedAt.Format(time.RFC3339),
	}

	// If content is restricted, hide sensitive data
	if restricted {
		restrictedMsg := "Conteúdo restrito"
		response.Content = &restrictedMsg
		response.ContentHtml = &restrictedMsg
		response.Summary = &restrictedMsg
		response.SummaryHtml = &restrictedMsg
		response.Notes = nil
		// Don't include items for restricted content
	} else {
		response.Content = anamnesis.Content
		response.ContentHtml = anamnesis.ContentHtml
		response.Summary = anamnesis.Summary
		response.SummaryHtml = anamnesis.SummaryHtml
		response.Notes = anamnesis.Notes
	}

	// Add template ID if present (only if not restricted)
	if !restricted && anamnesis.AnamnesisTemplateID != nil {
		templateID := anamnesis.AnamnesisTemplateID.String()
		response.AnamnesisTemplateID = &templateID
	}

	// Add Author if preloaded
	if anamnesis.Author.ID != uuid.Nil {
		// Get primary role (first one or most privileged)
		roles := anamnesis.Author.GetRoles()
		primaryRole := ""
		if len(roles) > 0 {
			primaryRole = roles[0]
		}

		response.Author = &dto.AuthorBrief{
			ID:    anamnesis.Author.ID.String(),
			Name:  anamnesis.Author.Name,
			Email: anamnesis.Author.Email,
			Role:  primaryRole,
		}
	}

	// Add AnamnesisTemplate if preloaded (only if not restricted)
	if !restricted && anamnesis.AnamnesisTemplate != nil && anamnesis.AnamnesisTemplate.ID != uuid.Nil {
		response.AnamnesisTemplate = &dto.AnamnesisTemplateBrief{
			ID:   anamnesis.AnamnesisTemplate.ID.String(),
			Name: anamnesis.AnamnesisTemplate.Name,
			Area: string(anamnesis.AnamnesisTemplate.Area),
		}
	}

	// Convert items (only if not restricted)
	if !restricted && len(anamnesis.Items) > 0 {
		response.Items = make([]dto.AnamnesisItemResponse, len(anamnesis.Items))
		for i, item := range anamnesis.Items {
			itemResponse := dto.AnamnesisItemResponse{
				ID:           item.ID.String(),
				ScoreItemID:  item.ScoreItemID.String(),
				TextValue:    item.TextValue,
				NumericValue: item.NumericValue,
				Order:        item.Order,
				CreatedAt:    item.CreatedAt.Format(time.RFC3339),
				UpdatedAt:    item.UpdatedAt.Format(time.RFC3339),
			}

			// Include ScoreItem if preloaded
			if item.ScoreItem != nil {
				scoreItem := &dto.ScoreItemBrief{
					ID:                 item.ScoreItem.ID.String(),
					Name:               item.ScoreItem.Name,
					Unit:               item.ScoreItem.Unit,
					UnitConversion:     item.ScoreItem.UnitConversion,
					ClinicalRelevance:  item.ScoreItem.ClinicalRelevance,
					PatientExplanation: item.ScoreItem.PatientExplanation,
					Conduct:            item.ScoreItem.Conduct,
					Points:             item.ScoreItem.Points,
					Order:              item.ScoreItem.Order,
					SubgroupID:         item.ScoreItem.SubgroupID.String(),
					CreatedAt:          item.ScoreItem.CreatedAt.Format(time.RFC3339),
					UpdatedAt:          item.ScoreItem.UpdatedAt.Format(time.RFC3339),
				}

				if item.ScoreItem.LastReview != nil {
					lr := item.ScoreItem.LastReview.Format(time.RFC3339)
					scoreItem.LastReview = &lr
				}

				if item.ScoreItem.ParentItemID != nil {
					pid := item.ScoreItem.ParentItemID.String()
					scoreItem.ParentItemID = &pid
				}

				// Include Subgroup if preloaded
				if item.ScoreItem.Subgroup != nil {
					subgroup := &dto.ScoreSubgroupBrief{
						ID:        item.ScoreItem.Subgroup.ID.String(),
						Name:      item.ScoreItem.Subgroup.Name,
						Order:     item.ScoreItem.Subgroup.Order,
						GroupID:   item.ScoreItem.Subgroup.GroupID.String(),
						CreatedAt: item.ScoreItem.Subgroup.CreatedAt.Format(time.RFC3339),
						UpdatedAt: item.ScoreItem.Subgroup.UpdatedAt.Format(time.RFC3339),
					}

					// Include Group if preloaded
					if item.ScoreItem.Subgroup.Group != nil {
						subgroup.Group = &dto.ScoreGroupBrief{
							ID:        item.ScoreItem.Subgroup.Group.ID.String(),
							Name:      item.ScoreItem.Subgroup.Group.Name,
							Order:     item.ScoreItem.Subgroup.Group.Order,
							CreatedAt: item.ScoreItem.Subgroup.Group.CreatedAt.Format(time.RFC3339),
							UpdatedAt: item.ScoreItem.Subgroup.Group.UpdatedAt.Format(time.RFC3339),
						}
					}

					scoreItem.Subgroup = subgroup
				}

				// Include Levels if preloaded
				if len(item.ScoreItem.Levels) > 0 {
					scoreItem.Levels = make([]dto.ScoreLevelBrief, len(item.ScoreItem.Levels))
					for j, level := range item.ScoreItem.Levels {
						levelBrief := dto.ScoreLevelBrief{
							ID:           level.ID.String(),
							Level:        level.Level,
							Name:         level.Name,
							LowerLimit:   level.LowerLimit,
							UpperLimit:   level.UpperLimit,
							Operator:     level.Operator,
							ClinicalRelevance: level.ClinicalRelevance,
							PatientExplanation: level.PatientExplanation,
							Conduct:      level.Conduct,
							ItemID:       level.ItemID.String(),
							CreatedAt:    level.CreatedAt.Format(time.RFC3339),
							UpdatedAt:    level.UpdatedAt.Format(time.RFC3339),
						}

						if level.LastReview != nil {
							lr := level.LastReview.Format(time.RFC3339)
							levelBrief.LastReview = &lr
						}

						scoreItem.Levels[j] = levelBrief
					}
				}

				itemResponse.ScoreItem = scoreItem
			}

			response.Items[i] = itemResponse
		}
	}

	return response
}
