package services

import (
	"errors"
	"os"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

var (
	ErrLabResultBatchNotFound    = errors.New("lab result batch not found")
	ErrNoPatientSelected         = errors.New("no patient selected - please select a patient first")
	ErrPatientMismatch           = errors.New("patient id does not match selected patient")
	ErrLabResultBatchPDFNotFound = errors.New("lab result batch has no original PDF")
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

			// Processar cada resultado do request (upsert-only: NUNCA deletamos em
			// massa aqui — remoção de exame individual é feita pela rota dedicada
			// DELETE /:batchId/results/:resultId. Deletar pelo que falta no payload
			// já causou perda silenciosa de exames classificados ao editar o lote.)
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
				} else {
					// Sem ID: CREATE novo resultado (guarda nil — campos são ponteiros omitempty)
					newResult := models.LabResult{
						LabResultBatchID: batchID,
						ResultText:       reqResult.ResultText,
						ResultNumeric:    reqResult.ResultNumeric,
						Unit:             reqResult.Unit,
						Interpretation:   reqResult.Interpretation,
						Level:            reqResult.Level,
					}
					if reqResult.TestName != nil {
						newResult.TestName = *reqResult.TestName
					}
					if reqResult.TestType != nil {
						newResult.TestType = *reqResult.TestType
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

	matched := labTestDefID != nil // default: casou se tem definição
	if req.Matched != nil {
		matched = *req.Matched
	}
	source := "manual"
	if req.Source != nil && *req.Source != "" {
		source = *req.Source
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
		Source:              source,
		MatchReason:         req.MatchReason,
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
	matched := labTestDefID != nil // default: casou se tem definição
	if req.Matched != nil {
		matched = *req.Matched
	}
	source := "manual"
	if req.Source != nil && *req.Source != "" {
		source = *req.Source
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
		Source:              source,
		MatchReason:         req.MatchReason,
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

	var reviewedAt *string
	if batch.ReviewedAt != nil {
		date := batch.ReviewedAt.Format(time.RFC3339)
		reviewedAt = &date
	}

	var pdfJobs int64
	s.db.Model(&models.ProcessingJob{}).
		Where("lab_result_batch_id = ? AND pdf_path <> ''", batch.ID).
		Count(&pdfJobs)

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
		IsCritical:         batch.IsCritical,
		WorstLevel:         batch.WorstLevel,
		ReviewedAt:         reviewedAt,
		HasPDF:             pdfJobs > 0,
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
		IsCritical:         baseResp.IsCritical,
		WorstLevel:         baseResp.WorstLevel,
		ReviewedAt:         baseResp.ReviewedAt,
		HasPDF:             baseResp.HasPDF,
		LabResults:         results,
		CreatedAt:          baseResp.CreatedAt,
		UpdatedAt:          baseResp.UpdatedAt,
	}
}

// GetPDFPath devolve o caminho do PDF original do lote (último ProcessingJob com PDF),
// verificando que o lote pertence ao paciente selecionado.
func (s *LabResultBatchService) GetPDFPath(batchID, userID uuid.UUID) (string, error) {
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return "", err
	}
	if user.SelectedPatientID == nil {
		return "", ErrNoPatientSelected
	}
	var batch models.LabResultBatch
	if err := s.db.First(&batch, batchID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", ErrLabResultBatchNotFound
		}
		return "", err
	}
	if batch.PatientID != *user.SelectedPatientID {
		return "", ErrPatientMismatch
	}
	var job models.ProcessingJob
	if err := s.db.Where("lab_result_batch_id = ? AND pdf_path <> ''", batchID).
		Order("created_at DESC").First(&job).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", ErrLabResultBatchPDFNotFound
		}
		return "", err
	}
	// O arquivo pode ter sumido do disco (uploads não-persistente em deploys antigos).
	if _, err := os.Stat(job.PDFPath); err != nil {
		return "", ErrLabResultBatchPDFNotFound
	}
	return job.PDFPath, nil
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
		Matched:               result.Matched,
		Source:                result.Source,
		MatchReason:           result.MatchReason,
		ClassifyReason:        result.ClassifyReason,
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

	// 3. Para cada LabResult: define o nível quando classificável; senão, registra o MOTIVO.
	for i := range batch.LabResults {
		result := &batch.LabResults[i]

		var levelToSet *int
		var reason *string
		setReason := func(r string) { rr := r; reason = &rr }

		switch {
		case result.LabTestDefinitionID == nil:
			setReason("Exame não catalogado no sistema")
		case result.ResultNumeric == nil:
			setReason("Resultado não numérico (qualitativo/texto)")
		default:
			if result.LabTestDefinition == nil {
				var def models.LabTestDefinition
				if err := s.db.First(&def, *result.LabTestDefinitionID).Error; err == nil {
					result.LabTestDefinition = &def
				}
			}
			if result.LabTestDefinition == nil {
				setReason("Definição do exame indisponível")
				break
			}

			var scoreItems []models.ScoreItem
			s.db.Preload("Levels", func(db *gorm.DB) *gorm.DB {
				return db.Order("level ASC")
			}).Where("lab_test_code = ?", result.LabTestDefinition.Code).Find(&scoreItems)

			if len(scoreItems) == 0 {
				setReason("Exame não entra no escore (sem item de score configurado)")
				break
			}

			var applicable []models.ScoreItem
			for _, it := range scoreItems {
				if it.AppliesToPatient(&patient) {
					applicable = append(applicable, it)
				}
			}
			if len(applicable) == 0 {
				setReason("Não se aplica a este paciente (sexo/idade/menopausa)")
				break
			}

			for _, lvl := range applicable[0].Levels {
				if lvl.EvaluatesTrue(*result.ResultNumeric) {
					l := lvl.Level
					levelToSet = &l
					break
				}
			}
			if levelToSet == nil {
				setReason("Valor fora das faixas de classificação configuradas")
			}
		}

		// Persiste level + classify_reason juntos (map garante gravar NULL quando nil).
		result.Level = levelToSet
		result.ClassifyReason = reason
		if err := s.db.Model(result).Updates(map[string]interface{}{
			"level":           levelToSet,
			"classify_reason": reason,
		}).Error; err != nil {
			continue // best-effort
		}
	}

	// 8. Recomputa criticidade do lote (worst_level/is_critical) p/ a Results inbox.
	return s.recomputeBatchCriticality(batchID)
}

// recomputeBatchCriticality recalcula worst_level/is_critical do lote a partir do level dos results.
// worst_level = menor level (0 = pior); is_critical = worst_level ∈ {0,1}. Persiste no batch.
func (s *LabResultBatchService) recomputeBatchCriticality(batchID uuid.UUID) error {
	var minLevel *int
	if err := s.db.Model(&models.LabResult{}).
		Where("lab_result_batch_id = ? AND level IS NOT NULL", batchID).
		Select("MIN(level)").
		Scan(&minLevel).Error; err != nil {
		return err
	}

	isCritical := minLevel != nil && (*minLevel == 0 || *minLevel == 1)

	return s.db.Model(&models.LabResultBatch{}).
		Where("id = ?", batchID).
		Updates(map[string]interface{}{
			"worst_level": minLevel,
			"is_critical": isCritical,
		}).Error
}

// ListPendingReview lista lotes a revisar de TODOS os pacientes (cross-patient, NÃO selectedPatient).
// Pendente = resultados chegaram (result_date != NULL) e ninguém deu ciência (reviewed_at == NULL).
// Ordem: críticos primeiro, depois pior level, depois mais recentes.
func (s *LabResultBatchService) ListPendingReview(limit, offset int) ([]*dto.LabInboxItemResponse, error) {
	if limit <= 0 || limit > 200 {
		limit = 100
	}
	var batches []models.LabResultBatch
	err := s.db.
		Preload("Patient").
		Preload("LabResults").
		Where("result_date IS NOT NULL AND reviewed_at IS NULL").
		Order("is_critical DESC").
		Order("worst_level ASC NULLS LAST").
		Order("result_date DESC").
		Limit(limit).
		Offset(offset).
		Find(&batches).Error
	if err != nil {
		return nil, err
	}

	out := make([]*dto.LabInboxItemResponse, len(batches))
	for i := range batches {
		out[i] = s.toInboxItem(&batches[i])
	}
	return out, nil
}

// CountPendingReview retorna (total, críticos) de lotes a revisar — usado no badge do sidebar.
func (s *LabResultBatchService) CountPendingReview() (int64, int64, error) {
	var total, critical int64
	if err := s.db.Model(&models.LabResultBatch{}).
		Where("result_date IS NOT NULL AND reviewed_at IS NULL").
		Count(&total).Error; err != nil {
		return 0, 0, err
	}
	if err := s.db.Model(&models.LabResultBatch{}).
		Where("result_date IS NOT NULL AND reviewed_at IS NULL AND is_critical = true").
		Count(&critical).Error; err != nil {
		return 0, 0, err
	}
	return total, critical, nil
}

// GetByIDForReview busca o detalhe de um lote SEM gate de selectedPatient (a inbox é cross-patient:
// qualquer clínico pode revisar). Inclui results + definições + paciente.
func (s *LabResultBatchService) GetByIDForReview(batchID uuid.UUID) (*dto.LabResultBatchDetailResponse, error) {
	var batch models.LabResultBatch
	err := s.db.
		Preload("Patient").
		Preload("LabResults.LabTestDefinition").
		First(&batch, batchID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrLabResultBatchNotFound
		}
		return nil, err
	}
	return s.toDetailResponse(&batch), nil
}

// AcknowledgeBatch registra a ciência clínica de um lote (cross-patient). Idempotente.
func (s *LabResultBatchService) AcknowledgeBatch(batchID, userID uuid.UUID) (*dto.LabResultBatchDetailResponse, error) {
	var batch models.LabResultBatch
	if err := s.db.First(&batch, batchID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrLabResultBatchNotFound
		}
		return nil, err
	}

	if batch.ReviewedAt == nil {
		now := time.Now()
		if err := s.db.Model(&batch).Updates(map[string]interface{}{
			"reviewed_at":         now,
			"reviewed_by_user_id": userID,
		}).Error; err != nil {
			return nil, err
		}
	}
	return s.GetByIDForReview(batchID)
}

// toInboxItem monta o item resumido da fila "Exames a revisar".
func (s *LabResultBatchService) toInboxItem(batch *models.LabResultBatch) *dto.LabInboxItemResponse {
	var resultDate, reviewedAt *string
	if batch.ResultDate != nil {
		d := batch.ResultDate.Format(time.RFC3339)
		resultDate = &d
	}
	if batch.ReviewedAt != nil {
		d := batch.ReviewedAt.Format(time.RFC3339)
		reviewedAt = &d
	}

	abnormal := 0
	for i := range batch.LabResults {
		if l := batch.LabResults[i].Level; l != nil && *l <= 2 {
			abnormal++
		}
	}

	return &dto.LabInboxItemResponse{
		ID:             batch.ID.String(),
		PatientID:      batch.PatientID.String(),
		PatientName:    batch.Patient.Name,
		LaboratoryName: batch.LaboratoryName,
		CollectionDate: batch.CollectionDate.Format(time.RFC3339),
		ResultDate:     resultDate,
		Status:         batch.Status,
		IsCritical:     batch.IsCritical,
		WorstLevel:     batch.WorstLevel,
		AbnormalCount:  abnormal,
		TotalResults:   len(batch.LabResults),
		ReviewedAt:     reviewedAt,
	}
}
