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

			// Aplicar conversão de unidade ANTES de criar
			if err := s.applyUnitConversion(&result); err != nil {
				// Log error mas não falha
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

	// Classificar resultados automaticamente
	if err := s.ClassifyBatchResults(batch.ID); err != nil {
		// Log error mas não falha a criação do batch
		// TODO: Log warning
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
	err := s.db.Preload("LabResults.LabTestDefinition").First(&batch, batchID).Error
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
		Preload("LabResults.LabTestDefinition")

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
					needsConversion := false
					if reqResult.LabTestDefinitionID != nil {
						parsed, err := uuid.Parse(*reqResult.LabTestDefinitionID)
						if err != nil {
							return errors.New("invalid lab test definition id")
						}
						existingResult.LabTestDefinitionID = &parsed
						needsConversion = true
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
						needsConversion = true
					}
					if reqResult.Unit != nil {
						existingResult.Unit = reqResult.Unit
						needsConversion = true
					}
					if reqResult.Interpretation != nil {
						existingResult.Interpretation = reqResult.Interpretation
					}
					if reqResult.Level != nil {
						existingResult.Level = reqResult.Level
					}

					// Aplicar conversão de unidade se mudou valor/unidade/definição
					if needsConversion {
						if err := s.applyUnitConversion(existingResult); err != nil {
							// Log error mas não falha
						}
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

					// Aplicar conversão de unidade ANTES de criar
					if err := s.applyUnitConversion(&newResult); err != nil {
						// Log error mas não falha
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

	// Classificar resultados automaticamente após update
	if err := s.ClassifyBatchResults(batchID); err != nil {
		// Log error mas não falha o update
		// TODO: Log warning
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

	// Aplicar conversão de unidade ANTES de criar
	if err := s.applyUnitConversion(&result); err != nil {
		// Log error mas não falha
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

	// Aplicar conversão de unidade ANTES de criar
	if err := s.applyUnitConversion(&result); err != nil {
		// Log error mas não falha
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
	needsConversion := false
	if req.LabTestDefinitionID != nil {
		parsed, err := uuid.Parse(*req.LabTestDefinitionID)
		if err != nil {
			return nil, errors.New("invalid lab test definition id")
		}
		result.LabTestDefinitionID = &parsed
		needsConversion = true
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
		needsConversion = true
	}
	if req.Unit != nil {
		result.Unit = req.Unit
		needsConversion = true
	}
	if req.Interpretation != nil {
		result.Interpretation = req.Interpretation
	}

	// Aplicar conversão de unidade se mudou valor/unidade/definição
	if needsConversion {
		if err := s.applyUnitConversion(&result); err != nil {
			// Log error mas não falha
		}
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

// applyUnitConversion aplica conversão de unidade se disponível
// Armazena valores originais e converte para unidade padrão
func (s *LabResultBatchService) applyUnitConversion(result *models.LabResult) error {
	// Se não tiver valor numérico ou unidade, skip
	if result.ResultNumeric == nil || result.Unit == nil {
		return nil
	}

	// Se não tiver LabTestDefinitionID, skip (sem como fazer conversão)
	if result.LabTestDefinitionID == nil {
		return nil
	}

	// 1. Armazenar valores ORIGINAIS
	result.ResultNumericOriginal = result.ResultNumeric
	result.UnitOriginal = result.Unit

	// 2. Buscar LabTestDefinition
	var labTestDef models.LabTestDefinition
	if err := s.db.First(&labTestDef, *result.LabTestDefinitionID).Error; err != nil {
		// Se não encontrou, manter original (não falhar)
		return nil
	}

	// 3. Tentar converter
	convertedValue, convertedUnit, wasConverted, err := labTestDef.ConvertToMainUnit(
		s.db,
		*result.ResultNumeric,
		*result.Unit,
	)

	if err != nil {
		// Se erro na conversão, manter original (não falhar)
		return nil
	}

	// 4. Se converteu, atualizar campos convertidos
	if wasConverted {
		result.ResultNumeric = &convertedValue
		result.Unit = &convertedUnit
	}
	// Se não converteu (wasConverted = false), mantém valores originais

	return nil
}

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
		PDFContentJSON:     batch.PDFContentJSON,
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
		PDFContentJSON:     baseResp.PDFContentJSON,
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

	// Incluir objeto LabTestDefinition se estiver preloaded
	var labTestDef *dto.LabTestDefinitionResponse
	if result.LabTestDefinition != nil {
		labTestDef = &dto.LabTestDefinitionResponse{
			ID:       result.LabTestDefinition.ID.String(),
			Name:     result.LabTestDefinition.Name,
			Code:     result.LabTestDefinition.Code,
			Category: string(result.LabTestDefinition.Category),
			Unit:     result.LabTestDefinition.Unit,
		}
	}

	return &dto.LabResultInBatchResponse{
		ID:                    result.ID.String(),
		LabResultBatchID:      result.LabResultBatchID.String(),
		LabTestDefinitionID:   labTestDefID,
		LabTestDefinition:     labTestDef,
		TestName:              result.TestName,
		TestType:              result.TestType,
		ResultText:            result.ResultText,
		ResultNumeric:         result.ResultNumeric,
		Unit:                  result.Unit,
		ResultNumericOriginal: result.ResultNumericOriginal,
		UnitOriginal:          result.UnitOriginal,
		Interpretation:        result.Interpretation,
		Level:                 result.Level,
		CreatedAt:             result.CreatedAt.Format(time.RFC3339),
		UpdatedAt:             result.UpdatedAt.Format(time.RFC3339),
	}
}

// UpdateAttachments atualiza o campo attachments de um batch
func (s *LabResultBatchService) UpdateAttachments(batchID uuid.UUID, attachments *string) error {
	return s.db.Model(&models.LabResultBatch{}).
		Where("id = ?", batchID).
		Update("attachments", attachments).Error
}

// ClassifyBatchResults classifica automaticamente os resultados de um batch baseado nos ScoreItems
// Esta função roda após salvamento de batch (manual ou importação de PDF)
func (s *LabResultBatchService) ClassifyBatchResults(batchID uuid.UUID) error {
	// 1. Buscar batch com LabResults e LabTestDefinitions preloaded
	var batch models.LabResultBatch
	if err := s.db.Preload("LabResults.LabTestDefinition").First(&batch, batchID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrLabResultBatchNotFound
		}
		return err
	}

	// 2. Buscar paciente
	var patient models.Patient
	if err := s.db.First(&patient, batch.PatientID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrPatientNotFound
		}
		return err
	}

	// 3. Para cada LabResult
	for i := range batch.LabResults {
		result := &batch.LabResults[i]

		// Se não tiver LabTestDefinitionID, skip
		if result.LabTestDefinitionID == nil {
			continue
		}

		// Se não tiver ResultNumeric, skip (não há como classificar)
		if result.ResultNumeric == nil {
			continue
		}

		// Buscar LabTestDefinition se não estiver preloaded
		if result.LabTestDefinition == nil {
			var labTestDef models.LabTestDefinition
			if err := s.db.First(&labTestDef, *result.LabTestDefinitionID).Error; err != nil {
				// Log error mas não falha - continua para próximo result
				continue
			}
			result.LabTestDefinition = &labTestDef
		}

		// 4. Buscar ScoreItems relacionados ao LabTestDefinition.Code
		var scoreItems []models.ScoreItem
		if err := s.db.
			Preload("Levels", func(db *gorm.DB) *gorm.DB {
				return db.Order("level ASC") // Ordenar por level (0-5)
			}).
			Where("lab_test_code = ?", result.LabTestDefinition.Code).
			Find(&scoreItems).Error; err != nil {
			// Log error mas não falha
			continue
		}

		if len(scoreItems) == 0 {
			// Nenhum ScoreItem configurado para este exame - skip
			continue
		}

		// 5. Filtrar ScoreItems por patient (gender, age, menopause)
		var applicableItems []models.ScoreItem
		for _, item := range scoreItems {
			if item.AppliesToPatient(&patient) {
				applicableItems = append(applicableItems, item)
			}
		}

		if len(applicableItems) == 0 {
			// Nenhum ScoreItem se aplica a este paciente - skip
			continue
		}

		// Se sobrar mais de 1, pegar o primeiro (já filtrado)
		selectedItem := applicableItems[0]

		// 6. Avaliar níveis (0-5) até encontrar o que satisfaz
		foundLevel := false
		for _, level := range selectedItem.Levels {
			if level.EvaluatesTrue(*result.ResultNumeric) {
				// Encontrou o nível! Salvar no LabResult
				result.Level = &level.Level
				if err := s.db.Model(result).Update("level", level.Level).Error; err != nil {
					// Log error mas não falha
					continue
				}
				foundLevel = true
				break
			}
		}

		// 7. Se nenhum nível satisfez, warning (erro de logística dos níveis)
		if !foundLevel && len(selectedItem.Levels) > 0 {
			// TODO: Log warning - nenhum level satisfez o valor (erro de configuração)
			// Por ora, deixar Level como nil (não forçar 0)
		}
	}

	return nil
}
