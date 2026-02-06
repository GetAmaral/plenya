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
	ErrLabResultBatchNotFound = errors.New("lab result batch not found")
	ErrNoPatientSelected      = errors.New("no patient selected - please select a patient first")
	ErrPatientMismatch        = errors.New("patient id does not match selected patient")
)

type LabResultBatchService struct {
	db *gorm.DB
}

func NewLabResultBatchService(db *gorm.DB) *LabResultBatchService {
	return &LabResultBatchService{db: db}
}

// Create cria um novo lote de resultados com múltiplos results (transação atômica)
func (s *LabResultBatchService) Create(userID uuid.UUID, req *dto.CreateLabResultBatchRequest) (*dto.LabResultBatchDetailResponse, error) {
	// CRITICAL SECURITY: Get user's selected patient
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}

	if user.SelectedPatientID == nil {
		return nil, ErrNoPatientSelected
	}

	patientID := *user.SelectedPatientID

	// Verificar se paciente existe
	var patient models.Patient
	if err := s.db.First(&patient, patientID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPatientNotFound
		}
		return nil, err
	}

	// Parse dates
	collectionDate, err := time.Parse(time.RFC3339, req.CollectionDate)
	if err != nil {
		return nil, errors.New("invalid collection date format, expected RFC3339")
	}

	var resultDate *time.Time
	if req.ResultDate != nil {
		parsed, err := time.Parse(time.RFC3339, *req.ResultDate)
		if err != nil {
			return nil, errors.New("invalid result date format, expected RFC3339")
		}
		resultDate = &parsed
	}

	// Parse optional UUIDs
	var labRequestID *uuid.UUID
	if req.LabRequestID != nil {
		parsed, err := uuid.Parse(*req.LabRequestID)
		if err != nil {
			return nil, errors.New("invalid lab request id")
		}
		labRequestID = &parsed
	}

	var requestingDoctorID *uuid.UUID
	if req.RequestingDoctorID != nil {
		parsed, err := uuid.Parse(*req.RequestingDoctorID)
		if err != nil {
			return nil, errors.New("invalid requesting doctor id")
		}
		requestingDoctorID = &parsed
	}

	// Transação: criar batch + results
	var batch models.LabResultBatch
	err = s.db.Transaction(func(tx *gorm.DB) error {
		// Criar batch
		batch = models.LabResultBatch{
			PatientID:          patientID,
			LabRequestID:       labRequestID,
			RequestingDoctorID: requestingDoctorID,
			LaboratoryName:     req.LaboratoryName,
			CollectionDate:     collectionDate,
			ResultDate:         resultDate,
			Status:             req.Status,
			Observations:       req.Observations,
			Attachments:        req.Attachments,
		}

		if err := tx.Create(&batch).Error; err != nil {
			return err
		}

		// Criar results
		for _, resReq := range req.LabResults {
			var labTestDefID *uuid.UUID
			if resReq.LabTestDefinitionID != nil {
				parsed, err := uuid.Parse(*resReq.LabTestDefinitionID)
				if err != nil {
					return errors.New("invalid lab test definition id")
				}
				labTestDefID = &parsed
			}

			result := models.LabResult{
				LabResultBatchID:    batch.ID,
				LabTestDefinitionID: labTestDefID,
				TestName:            resReq.TestName,
				TestType:            resReq.TestType,
				ResultText:          resReq.ResultText,
				ResultNumeric:       resReq.ResultNumeric,
				Unit:                resReq.Unit,
				Interpretation:      resReq.Interpretation,
				Level:               resReq.Level,
			}

			if err := tx.Create(&result).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Retornar com preload
	return s.GetByID(batch.ID, userID)
}

// GetByID busca um batch por ID com preload de results
func (s *LabResultBatchService) GetByID(batchID, userID uuid.UUID) (*dto.LabResultBatchDetailResponse, error) {
	// CRITICAL SECURITY: Get user's selected patient
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}

	if user.SelectedPatientID == nil {
		return nil, ErrNoPatientSelected
	}

	var batch models.LabResultBatch
	err := s.db.Preload("LabResults").First(&batch, batchID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrLabResultBatchNotFound
		}
		return nil, err
	}

	// SECURITY: Verify batch belongs to selected patient
	if batch.PatientID != *user.SelectedPatientID {
		return nil, ErrPatientMismatch
	}

	// Convert to response
	return s.toDetailResponse(&batch), nil
}

// List lista batches do selectedPatient (com paginação e filtros)
// Retorna DetailResponse para incluir labResults (necessário para visualização pivot)
func (s *LabResultBatchService) List(userID uuid.UUID, status *string, limit, offset int) ([]*dto.LabResultBatchDetailResponse, error) {
	// CRITICAL SECURITY: Get user's selected patient
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}

	if user.SelectedPatientID == nil {
		return nil, ErrNoPatientSelected
	}

	query := s.db.Model(&models.LabResultBatch{}).
		Where("patient_id = ?", *user.SelectedPatientID).
		Preload("LabResults")

	if status != nil && *status != "" {
		query = query.Where("status = ?", *status)
	}

	var batches []models.LabResultBatch
	err := query.Order("collection_date DESC").
		Limit(limit).
		Offset(offset).
		Find(&batches).Error

	if err != nil {
		return nil, err
	}

	// Usar toDetailResponse para incluir labResults
	responses := make([]*dto.LabResultBatchDetailResponse, len(batches))
	for i, batch := range batches {
		responses[i] = s.toDetailResponse(&batch)
	}

	return responses, nil
}

// Update atualiza batch e sincroniza resultados (create/update/delete)
func (s *LabResultBatchService) Update(batchID, userID uuid.UUID, req *dto.UpdateLabResultBatchRequest) (*dto.LabResultBatchResponse, error) {
	// CRITICAL SECURITY: Get user's selected patient
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}

	if user.SelectedPatientID == nil {
		return nil, ErrNoPatientSelected
	}

	var batch models.LabResultBatch
	if err := s.db.Preload("LabResults").First(&batch, batchID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrLabResultBatchNotFound
		}
		return nil, err
	}

	// SECURITY: Verify batch belongs to selected patient
	if batch.PatientID != *user.SelectedPatientID {
		return nil, ErrPatientMismatch
	}

	// Usar transação para atomicidade
	err := s.db.Transaction(func(tx *gorm.DB) error {
		// 1. Update batch metadata
		if req.LaboratoryName != nil {
			batch.LaboratoryName = *req.LaboratoryName
		}
		if req.CollectionDate != nil {
			parsed, err := time.Parse(time.RFC3339, *req.CollectionDate)
			if err != nil {
				return errors.New("invalid collection date format")
			}
			batch.CollectionDate = parsed
		}
		if req.ResultDate != nil {
			parsed, err := time.Parse(time.RFC3339, *req.ResultDate)
			if err != nil {
				return errors.New("invalid result date format")
			}
			batch.ResultDate = &parsed
		}
		if req.Status != nil {
			batch.Status = *req.Status
		}
		if req.Observations != nil {
			batch.Observations = req.Observations
		}
		if req.Attachments != nil {
			batch.Attachments = req.Attachments
		}
		if req.LabRequestID != nil {
			parsed, err := uuid.Parse(*req.LabRequestID)
			if err != nil {
				return errors.New("invalid lab request id")
			}
			batch.LabRequestID = &parsed
		}
		if req.RequestingDoctorID != nil {
			parsed, err := uuid.Parse(*req.RequestingDoctorID)
			if err != nil {
				return errors.New("invalid requesting doctor id")
			}
			batch.RequestingDoctorID = &parsed
		}

		if err := tx.Save(&batch).Error; err != nil {
			return err
		}

		// 2. Sincronizar resultados (se fornecidos)
		if req.LabResults != nil {
			// Criar map de IDs existentes para comparação
			existingResultsMap := make(map[uuid.UUID]*models.LabResult)
			for i := range batch.LabResults {
				existingResultsMap[batch.LabResults[i].ID] = &batch.LabResults[i]
			}

			// Criar map de IDs do request para detectar quais manter
			requestedIDs := make(map[uuid.UUID]bool)

			// Processar cada resultado do request
			for _, reqResult := range req.LabResults {
				if reqResult.ID != nil && *reqResult.ID != "" {
					// Tem ID: UPDATE resultado existente
					resultID, err := uuid.Parse(*reqResult.ID)
					if err != nil {
						return errors.New("invalid result id: " + *reqResult.ID)
					}

					existingResult, exists := existingResultsMap[resultID]
					if !exists {
						return errors.New("result not found: " + resultID.String())
					}

					// Atualizar campos
					if reqResult.LabTestDefinitionID != nil {
						parsed, err := uuid.Parse(*reqResult.LabTestDefinitionID)
						if err != nil {
							return errors.New("invalid lab test definition id")
						}
						existingResult.LabTestDefinitionID = &parsed
					}
					if reqResult.TestName != nil {
						existingResult.TestName = *reqResult.TestName
					}
					if reqResult.TestType != nil {
						existingResult.TestType = *reqResult.TestType
					}
					if reqResult.ResultText != nil {
						existingResult.ResultText = reqResult.ResultText
					}
					if reqResult.ResultNumeric != nil {
						existingResult.ResultNumeric = reqResult.ResultNumeric
					}
					if reqResult.Unit != nil {
						existingResult.Unit = reqResult.Unit
					}
					if reqResult.Interpretation != nil {
						existingResult.Interpretation = reqResult.Interpretation
					}
					if reqResult.Level != nil {
						existingResult.Level = reqResult.Level
					}

					if err := tx.Save(existingResult).Error; err != nil {
						return err
					}

					requestedIDs[resultID] = true
				} else {
					// Sem ID: CREATE novo resultado
					newResult := models.LabResult{
						LabResultBatchID: batchID,
						TestName:         *reqResult.TestName,
						TestType:         *reqResult.TestType,
						ResultText:       reqResult.ResultText,
						ResultNumeric:    reqResult.ResultNumeric,
						Unit:             reqResult.Unit,
						Interpretation:   reqResult.Interpretation,
						Level:            reqResult.Level,
					}

					if reqResult.LabTestDefinitionID != nil {
						parsed, err := uuid.Parse(*reqResult.LabTestDefinitionID)
						if err != nil {
							return errors.New("invalid lab test definition id")
						}
						newResult.LabTestDefinitionID = &parsed
					}

					if err := tx.Create(&newResult).Error; err != nil {
						return err
					}
				}
			}

			// 3. DELETE resultados que foram removidos (existiam mas não estão no request)
			for id := range existingResultsMap {
				if !requestedIDs[id] {
					if err := tx.Delete(&models.LabResult{}, id).Error; err != nil {
						return err
					}
				}
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Recarregar batch com resultados atualizados
	if err := s.db.Preload("LabResults").First(&batch, batchID).Error; err != nil {
		return nil, err
	}

	resp := s.toResponse(&batch)
	return &resp, nil
}

// Delete soft deleta um batch (admin only)
func (s *LabResultBatchService) Delete(batchID, userID uuid.UUID, userRole string) error {
	if userRole != "admin" {
		return errors.New("only admins can delete lab result batches")
	}

	var batch models.LabResultBatch
	if err := s.db.First(&batch, batchID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrLabResultBatchNotFound
		}
		return err
	}

	return s.db.Delete(&batch).Error
}

// AddResultInternal adiciona um result a um batch (uso interno, sem verificação de usuário)
// IMPORTANTE: Usar apenas em contextos de processamento em background (jobs)
func (s *LabResultBatchService) AddResultInternal(batchID uuid.UUID, req *dto.CreateLabResultInBatchRequest) (*dto.LabResultInBatchResponse, error) {
	var labTestDefID *uuid.UUID
	if req.LabTestDefinitionID != nil {
		parsed, err := uuid.Parse(*req.LabTestDefinitionID)
		if err != nil {
			return nil, errors.New("invalid lab test definition id")
		}
		labTestDefID = &parsed
	}

	matched := true
	if req.Matched != nil {
		matched = *req.Matched
	}

	result := models.LabResult{
		LabResultBatchID:    batchID,
		LabTestDefinitionID: labTestDefID,
		TestName:            req.TestName,
		TestType:            req.TestType,
		ResultText:          req.ResultText,
		ResultNumeric:       req.ResultNumeric,
		Unit:                req.Unit,
		Interpretation:      req.Interpretation,
		Matched:             matched,
	}

	if err := s.db.Create(&result).Error; err != nil {
		return nil, err
	}

	return s.toLabResultResponse(&result), nil
}

// AddResult adiciona um result a um batch existente
func (s *LabResultBatchService) AddResult(batchID, userID uuid.UUID, req *dto.CreateLabResultInBatchRequest) (*dto.LabResultInBatchResponse, error) {
	// CRITICAL SECURITY: Get user's selected patient
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}

	if user.SelectedPatientID == nil {
		return nil, ErrNoPatientSelected
	}

	var batch models.LabResultBatch
	if err := s.db.First(&batch, batchID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrLabResultBatchNotFound
		}
		return nil, err
	}

	// SECURITY: Verify batch belongs to selected patient
	if batch.PatientID != *user.SelectedPatientID {
		return nil, ErrPatientMismatch
	}

	var labTestDefID *uuid.UUID
	if req.LabTestDefinitionID != nil {
		parsed, err := uuid.Parse(*req.LabTestDefinitionID)
		if err != nil {
			return nil, errors.New("invalid lab test definition id")
		}
		labTestDefID = &parsed
	}

	// Default matched to true (manual entry is considered matched)
	matched := true
	if req.Matched != nil {
		matched = *req.Matched
	}

	result := models.LabResult{
		LabResultBatchID:    batchID,
		LabTestDefinitionID: labTestDefID,
		TestName:            req.TestName,
		TestType:            req.TestType,
		ResultText:          req.ResultText,
		ResultNumeric:       req.ResultNumeric,
		Unit:                req.Unit,
		Interpretation:      req.Interpretation,
		Matched:             matched,
	}

	if err := s.db.Create(&result).Error; err != nil {
		return nil, err
	}

	return s.toLabResultResponse(&result), nil
}

// UpdateResult atualiza um result individual
func (s *LabResultBatchService) UpdateResult(batchID, resultID, userID uuid.UUID, req *dto.UpdateLabResultInBatchRequest) (*dto.LabResultInBatchResponse, error) {
	// CRITICAL SECURITY: Get user's selected patient
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}

	if user.SelectedPatientID == nil {
		return nil, ErrNoPatientSelected
	}

	var batch models.LabResultBatch
	if err := s.db.First(&batch, batchID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrLabResultBatchNotFound
		}
		return nil, err
	}

	// SECURITY: Verify batch belongs to selected patient
	if batch.PatientID != *user.SelectedPatientID {
		return nil, ErrPatientMismatch
	}

	var result models.LabResult
	if err := s.db.Where("id = ? AND lab_result_batch_id = ?", resultID, batchID).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrLabResultNotFound
		}
		return nil, err
	}

	// Update fields
	if req.LabTestDefinitionID != nil {
		parsed, err := uuid.Parse(*req.LabTestDefinitionID)
		if err != nil {
			return nil, errors.New("invalid lab test definition id")
		}
		result.LabTestDefinitionID = &parsed
	}
	if req.TestName != nil {
		result.TestName = *req.TestName
	}
	if req.TestType != nil {
		result.TestType = *req.TestType
	}
	if req.ResultText != nil {
		result.ResultText = req.ResultText
	}
	if req.ResultNumeric != nil {
		result.ResultNumeric = req.ResultNumeric
	}
	if req.Unit != nil {
		result.Unit = req.Unit
	}
	if req.Interpretation != nil {
		result.Interpretation = req.Interpretation
	}

	if err := s.db.Save(&result).Error; err != nil {
		return nil, err
	}

	return s.toLabResultResponse(&result), nil
}

// DeleteResult deleta um result (admin only)
func (s *LabResultBatchService) DeleteResult(batchID, resultID, userID uuid.UUID, userRole string) error {
	if userRole != "admin" {
		return errors.New("only admins can delete lab results")
	}

	var result models.LabResult
	if err := s.db.Where("id = ? AND lab_result_batch_id = ?", resultID, batchID).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrLabResultNotFound
		}
		return err
	}

	return s.db.Delete(&result).Error
}

// Helper functions

func (s *LabResultBatchService) toResponse(batch *models.LabResultBatch) dto.LabResultBatchResponse {
	var labRequestID, requestingDoctorID, resultDate *string

	if batch.LabRequestID != nil {
		id := batch.LabRequestID.String()
		labRequestID = &id
	}
	if batch.RequestingDoctorID != nil {
		id := batch.RequestingDoctorID.String()
		requestingDoctorID = &id
	}
	if batch.ResultDate != nil {
		date := batch.ResultDate.Format(time.RFC3339)
		resultDate = &date
	}

	return dto.LabResultBatchResponse{
		ID:                 batch.ID.String(),
		PatientID:          batch.PatientID.String(),
		LabRequestID:       labRequestID,
		RequestingDoctorID: requestingDoctorID,
		LaboratoryName:     batch.LaboratoryName,
		CollectionDate:     batch.CollectionDate.Format(time.RFC3339),
		ResultDate:         resultDate,
		Status:             batch.Status,
		Observations:       batch.Observations,
		Attachments:        batch.Attachments,
		ResultCount:        len(batch.LabResults),
		CreatedAt:          batch.CreatedAt.Format(time.RFC3339),
		UpdatedAt:          batch.UpdatedAt.Format(time.RFC3339),
	}
}

func (s *LabResultBatchService) toDetailResponse(batch *models.LabResultBatch) *dto.LabResultBatchDetailResponse {
	baseResp := s.toResponse(batch)

	results := make([]dto.LabResultInBatchResponse, len(batch.LabResults))
	for i, result := range batch.LabResults {
		results[i] = *s.toLabResultResponse(&result)
	}

	return &dto.LabResultBatchDetailResponse{
		ID:                 baseResp.ID,
		PatientID:          baseResp.PatientID,
		LabRequestID:       baseResp.LabRequestID,
		RequestingDoctorID: baseResp.RequestingDoctorID,
		LaboratoryName:     baseResp.LaboratoryName,
		CollectionDate:     baseResp.CollectionDate,
		ResultDate:         baseResp.ResultDate,
		Status:             baseResp.Status,
		Observations:       baseResp.Observations,
		Attachments:        baseResp.Attachments,
		ResultCount:        baseResp.ResultCount,
		LabResults:         results,
		CreatedAt:          baseResp.CreatedAt,
		UpdatedAt:          baseResp.UpdatedAt,
	}
}

func (s *LabResultBatchService) toLabResultResponse(result *models.LabResult) *dto.LabResultInBatchResponse {
	var labTestDefID *string
	if result.LabTestDefinitionID != nil {
		id := result.LabTestDefinitionID.String()
		labTestDefID = &id
	}

	return &dto.LabResultInBatchResponse{
		ID:                  result.ID.String(),
		LabResultBatchID:    result.LabResultBatchID.String(),
		LabTestDefinitionID: labTestDefID,
		TestName:            result.TestName,
		TestType:            result.TestType,
		ResultText:          result.ResultText,
		ResultNumeric:       result.ResultNumeric,
		Unit:                result.Unit,
		Interpretation:      result.Interpretation,
		Level:               result.Level,
		CreatedAt:           result.CreatedAt.Format(time.RFC3339),
		UpdatedAt:           result.UpdatedAt.Format(time.RFC3339),
	}
}

// UpdateAttachments atualiza o campo attachments de um batch
func (s *LabResultBatchService) UpdateAttachments(batchID uuid.UUID, attachments *string) error {
	return s.db.Model(&models.LabResultBatch{}).
		Where("id = ?", batchID).
		Update("attachments", attachments).Error
}
