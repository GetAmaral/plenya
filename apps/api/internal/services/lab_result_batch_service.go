package services

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/utils"
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

			// Os mesmos campos que `AddResultInternal` mapeia.
			//
			// Criar o lote de uma vez perdia sete deles, e o silêncio custava caro: `Matched` não
			// tem default no model, então TODO resultado criado por aqui nascia `matched=false` e a
			// tela mostrava "não catalogado" para um lote inteiro que estava catalogado. A faixa de
			// referência, que o dossiê lê para a régua, ia junto para o ralo.
			matched := labTestDefID != nil // casou, se veio com definição
			if resReq.Matched != nil {
				matched = *resReq.Matched
			}
			source := "manual"
			if resReq.Source != nil && *resReq.Source != "" {
				source = *resReq.Source
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
				Matched:             matched,
				Source:              source,
				MatchReason:         resReq.MatchReason,
				Specimen:            resReq.Specimen,
				Method:              resReq.Method,
				ReferenceRange:      resReq.ReferenceRange,
				CollectionDate:      resReq.CollectionDate,
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
		Specimen:            req.Specimen,
		Method:              req.Method,
		ReferenceRange:      req.ReferenceRange,
		CollectionDate:      req.CollectionDate,
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
		Specimen:            req.Specimen,
		Method:              req.Method,
		ReferenceRange:      req.ReferenceRange,
		CollectionDate:      req.CollectionDate,
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

// applyUnitConversion leva o resultado para a unidade do exame e REGISTRA o que aconteceu.
//
// Antes esta função desistia em silêncio quando o par (exame, unidade) não estava na tabela
// curada: devolvia o valor original sem erro e sem marca. O resultado ficava gravado na unidade
// do laudo e, lá na frente, o motor do escore comparava esse número contra uma escala em outra
// grandeza. Agora ou converte, ou diz por que não deu.
func (s *LabResultBatchService) applyUnitConversion(result *models.LabResult) error {
	if result.ResultNumeric == nil || result.Unit == nil {
		return nil
	}
	if result.LabTestDefinitionID == nil {
		return nil
	}

	// Valores como vieram do laudo, sempre, mesmo que nada mude depois.
	result.ResultNumericOriginal = result.ResultNumeric
	result.UnitOriginal = result.Unit

	var labTestDef models.LabTestDefinition
	if err := s.db.First(&labTestDef, *result.LabTestDefinitionID).Error; err != nil {
		return nil
	}

	conv := labTestDef.ConverteParaUnidadePrincipal(
		s.db, *result.ResultNumeric, *result.Unit, s.faixaPlausivel(labTestDef.Code))

	if conv.Status == models.ConversaoAplicada {
		valor, unidade := conv.Valor, conv.Unidade
		result.ResultNumeric = &valor
		result.Unit = &unidade
	}

	status := string(conv.Status)
	result.UnitConversionStatus = &status
	if conv.Motivo != "" && conv.Status == models.ConversaoPendente {
		motivo := conv.Motivo
		result.UnitConversionNote = &motivo
	} else {
		result.UnitConversionNote = nil
	}

	return nil
}

// faixaPlausivel devolve um teste de sanidade para o valor convertido, montado a partir das
// faixas que o próprio catálogo do escore define para aquele exame.
//
// Serve contra rótulo errado, não contra paciente doente: a margem é de dez vezes para cada lado
// do que as faixas cobrem, então um valor clinicamente péssimo passa e só um erro de ordem de
// grandeza é barrado. Sem faixas cadastradas não há como julgar, e aí aceita.
func (s *LabResultBatchService) faixaPlausivel(code string) func(float64) bool {
	if code == "" {
		return nil
	}

	var limites struct {
		Menor *float64
		Maior *float64
	}
	err := s.db.Table("score_levels AS sl").
		Select(`MIN(LEAST(NULLIF(sl.lower_limit,'')::double precision, NULLIF(sl.upper_limit,'')::double precision)) AS menor,
		        MAX(GREATEST(NULLIF(sl.lower_limit,'')::double precision, NULLIF(sl.upper_limit,'')::double precision)) AS maior`).
		Joins("JOIN score_items si ON si.id = sl.score_item_id").
		Where("si.lab_test_code = ? AND sl.deleted_at IS NULL AND si.deleted_at IS NULL", code).
		Scan(&limites).Error
	if err != nil || limites.Menor == nil || limites.Maior == nil {
		return nil
	}

	menor, maior := *limites.Menor, *limites.Maior
	if maior <= 0 || maior < menor {
		return nil
	}
	piso, teto := menor/10, maior*10
	if menor <= 0 {
		// Escala que passa por zero ou é negativa (T-score da densitometria): sem piso útil.
		piso = menor*10 - 10
	}

	return func(v float64) bool { return v >= piso && v <= teto }
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
// Devolve o caminho e o NOME do arquivo. O laudo original é reenviado por WhatsApp como qualquer
// outro documento, e saía chamado "laudo.pdf" para todo paciente e toda data — na pasta de quem
// recebe, o segundo já sobrescrevia o primeiro.
func (s *LabResultBatchService) GetPDFPath(batchID, userID uuid.UUID) (path, fileName string, err error) {
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return "", "", err
	}
	if user.SelectedPatientID == nil {
		return "", "", ErrNoPatientSelected
	}
	var batch models.LabResultBatch
	if err := s.db.First(&batch, batchID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", "", ErrLabResultBatchNotFound
		}
		return "", "", err
	}
	if batch.PatientID != *user.SelectedPatientID {
		return "", "", ErrPatientMismatch
	}
	var job models.ProcessingJob
	if err := s.db.Where("lab_result_batch_id = ? AND pdf_path <> ''", batchID).
		Order("created_at DESC").First(&job).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", "", ErrLabResultBatchPDFNotFound
		}
		return "", "", err
	}
	// O arquivo pode ter sumido do disco (uploads não-persistente em deploys antigos).
	if _, err := os.Stat(job.PDFPath); err != nil {
		return "", "", ErrLabResultBatchPDFNotFound
	}
	// A data do arquivo é a da COLETA, não a do upload: é ela que o médico procura, e um laudo
	// antigo digitalizado hoje sairia com a data de hoje.
	dia := batch.CollectionDate
	if dia.IsZero() {
		dia = batch.CreatedAt
	}
	// O nome do paciente é COSMÉTICO aqui, então a falha em lê-lo não pode derrubar a rota: quem
	// chama isto também é o Reinterpret, que descarta o nome e mesmo assim quebraria. Unscoped
	// porque o laudo de um paciente arquivado continua sendo dele; sem nome, DocumentFileName
	// cai em "Paciente".
	var patient models.Patient
	nome := ""
	if err := s.db.Unscoped().First(&patient, batch.PatientID).Error; err == nil {
		nome = patient.Name
	}
	return job.PDFPath, utils.DocumentFileName(nome, "Laudo", dia, batch.ID), nil
}

// DeletePDFResultsForReinterpret apaga (soft delete) os resultados que vieram do PDF de um
// lote, preservando os lançados à mão. Usado antes de mandar a IA reler o laudo: sem isso a
// releitura duplicaria cada exame. Retorna quantos foram removidos.
func (s *LabResultBatchService) DeletePDFResultsForReinterpret(batchID uuid.UUID) (int64, error) {
	tx := s.db.Where("lab_result_batch_id = ? AND source = ?", batchID, "pdf").
		Delete(&models.LabResult{})
	return tx.RowsAffected, tx.Error
}

// RestorePDFResultsDeletedSince desfaz o soft delete acima. Usado quando a releitura não
// chega nem a ser enfileirada: sem isso o lote ficaria vazio por causa de um erro nosso.
// O corte por `since` evita ressuscitar exames que o usuário apagou antes, de propósito.
func (s *LabResultBatchService) RestorePDFResultsDeletedSince(batchID uuid.UUID, since time.Time) error {
	return s.db.Unscoped().
		Model(&models.LabResult{}).
		Where("lab_result_batch_id = ? AND source = ? AND deleted_at >= ?", batchID, "pdf", since).
		Update("deleted_at", nil).Error
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

	var collectionDate *string
	if result.CollectionDate != nil {
		cd := result.CollectionDate.Format(time.RFC3339)
		collectionDate = &cd
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
		Specimen:              result.Specimen,
		Method:                result.Method,
		ReferenceRange:        result.ReferenceRange,
		CollectionDate:        collectionDate,
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

	// Sinônimos de unidade do catálogo, para não recusar `mEq/L` contra `mmol/L` no sódio.
	catalogo := carregaCatalogoDeExames(s.db)

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

		hasText := result.ResultText != nil && strings.TrimSpace(*result.ResultText) != ""

		switch {
		case result.LabTestDefinitionID == nil:
			setReason("Exame não catalogado no sistema")
		case result.ResultNumeric == nil && !hasText:
			setReason("Resultado vazio")
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

			item := pickScoringItem(applicable)
			if item == nil {
				setReason("Exame não entra no escore (item sem faixas configuradas)")
				break
			}

			classifyNumeric := func(value float64) {
				for _, lvl := range item.Levels {
					if lvl.EvaluatesTrue(value) {
						l := lvl.Level
						levelToSet = &l
						return
					}
				}
			}

			if result.ResultNumeric != nil {
				// A escala do item e o laudo precisam falar da mesma grandeza. Três itens do
				// sedimento urinário estão em `células/campo` e o laboratório reporta `/µL`:
				// 0,5/µL cai na faixa "≤10 células/campo" e o resultado fica gravado como
				// nível ÓTIMO, que é o que aparece na tela e na régua da devolutiva.
				unidadeDoLaudo := ""
				if result.Unit != nil {
					unidadeDoLaudo = *result.Unit
				}
				if !item.UnitMatches(unidadeDoLaudo, catalogo.sinonimosDe(item.LabTestCode)) {
					setReason(fmt.Sprintf(
						"Não classificado: a faixa do escore está em %s e o resultado veio em %s",
						*item.Unit, unidadeDoLaudo))
					break
				}

				classifyNumeric(*result.ResultNumeric)
				if levelToSet == nil {
					setReason("Valor fora das faixas de classificação configuradas")
				}
				break
			}

			// Laudo ainda sem o resultado: não é pendência de classificação, é exame em curso.
			if isPendingLabText(*result.ResultText) {
				setReason("Resultado ainda não liberado pelo laboratório")
				break
			}

			// Resultado em texto (sorologia, cultura, sedimento): casa com o NOME do nível.
			if l := matchQualitativeLevel(item.Levels, *result.ResultText); l != nil {
				levelToSet = l
				break
			}

			// "Superior a 1.000,0", "< 5", "maior que 100": o laudo não deu um número puro,
			// mas deu um número. Classifica por ele em vez de desistir.
			if value, ok := numericFromComparativeText(*result.ResultText); ok {
				classifyNumeric(value)
			}
			if levelToSet == nil {
				setReason("Resultado em texto não bate com nenhum nível configurado")
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

// pickScoringItem escolhe qual ScoreItem manda quando o mesmo exame tem vários. Um laudo
// como o IGF-1 tem itens por faixa etária MAIS um item guarda-chuva sem faixa nenhuma;
// pegar o primeiro da lista deixava o resultado sem nível sempre que o guarda-chuva vinha
// antes. Regra: só entram itens com faixas; entre eles vence o mais específico (com recorte
// de idade, e o recorte mais estreito).
func pickScoringItem(items []models.ScoreItem) *models.ScoreItem {
	var best *models.ScoreItem
	bestWidth := math.MaxInt32

	for i := range items {
		it := &items[i]
		if len(it.Levels) == 0 {
			continue
		}

		width := math.MaxInt32 - 1 // sem recorte de idade: menos específico que qualquer recorte
		if it.AgeRangeMin != nil || it.AgeRangeMax != nil {
			min, max := 0, 150
			if it.AgeRangeMin != nil {
				min = *it.AgeRangeMin
			}
			if it.AgeRangeMax != nil {
				max = *it.AgeRangeMax
			}
			width = max - min
		}

		if best == nil || width < bestWidth {
			best, bestWidth = it, width
		}
	}

	return best
}

// qualitativeSynonyms mapeia como os laudos escrevem um resultado qualitativo para o
// vocabulário usado nos níveis do escore.
var qualitativeSynonyms = map[string]string{
	"positivo":       "reagente",
	"positiva":       "reagente",
	"detectavel":     "reagente",
	"detectado":      "reagente",
	"presente":       "reagente",
	"negativo":       "nao reagente",
	"negativa":       "nao reagente",
	"nao detectavel": "nao reagente",
	"nao detectado":  "nao reagente",
	"indetectavel":   "nao reagente",
	"ausente":        "nao reagente",
	"nao reagente":   "nao reagente",
}

// pendingLabTexts são os avisos que o laboratório imprime no lugar do resultado.
var pendingLabTexts = []string{
	"em andamento", "aguardando", "nao realizado", "material insuficiente",
	"amostra insuficiente", "a liberar", "pendente", "em analise",
}

// isPendingLabText reconhece laudo sem resultado ainda — não é falha de classificação.
func isPendingLabText(text string) bool {
	value := normalizeQualitative(text)
	for _, marker := range pendingLabTexts {
		if strings.Contains(value, marker) {
			return true
		}
	}
	return false
}

// normalizeQualitative deixa o texto comparável: minúsculas, sem acento, sem hífen, sem o
// que estiver entre parênteses e com espaços colapsados. "Não-reagente", "nao reagente" e
// "Negativo (<15)" passam a ser comparáveis entre si.
func normalizeQualitative(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	s = removeAccentsFromString(s)
	s = stripParenthesized(s)
	s = strings.NewReplacer("-", " ", "/", " / ").Replace(s)
	return strings.Join(strings.Fields(s), " ")
}

// stripParenthesized remove trechos entre parênteses ("Negativo (<15)" → "Negativo").
func stripParenthesized(s string) string {
	var b strings.Builder
	depth := 0
	for _, r := range s {
		switch r {
		case '(':
			depth++
		case ')':
			if depth > 0 {
				depth--
			}
		default:
			if depth == 0 {
				b.WriteRune(r)
			}
		}
	}
	return b.String()
}

// sameWord compara palavras tolerando plural ("ausente" ≡ "ausentes").
func sameWord(a, b string) bool {
	return a == b || a == b+"s" || a+"s" == b
}

// containsPhrase procura a sequência de palavras `phrase` dentro de `tokens`, respeitando
// limite de palavra — "amostra nao reagente para hiv" contém "nao reagente", mas "reagente"
// sozinho NUNCA casa com um texto que diz "nao reagente" (a checagem é por sequência).
func containsPhrase(tokens, phrase []string) bool {
	if len(phrase) == 0 || len(phrase) > len(tokens) {
		return false
	}
	for i := 0; i+len(phrase) <= len(tokens); i++ {
		ok := true
		for j := range phrase {
			if !sameWord(tokens[i+j], phrase[j]) {
				ok = false
				break
			}
		}
		if ok {
			// "reagente" dentro de "nao reagente" seria o oposto do resultado: recusa.
			if i > 0 && tokens[i-1] == "nao" && (len(phrase) == 0 || phrase[0] != "nao") {
				continue
			}
			return true
		}
	}
	return false
}

// levelNameAlternatives quebra o nome do nível nas formas que o laudo pode usar:
// "Límpido/Cristalino" → ["limpido", "cristalino"].
func levelNameAlternatives(levelName string) []string {
	normalized := normalizeQualitative(levelName)
	var out []string
	for _, part := range strings.Split(normalized, "/") {
		if p := strings.TrimSpace(part); p != "" {
			out = append(out, p)
		}
	}
	return out
}

// matchQualitativeLevel classifica resultado em texto pelo NOME do nível. Sorologia é o caso
// típico ("Reagente" 0 / "Não-reagente" 5), mas vale também para sedimento ("Ausentes"),
// cultura ("Negativa") e aspecto da urina ("Límpido/Cristalino"). Empate entre níveis
// diferentes NÃO classifica: rótulo clínico errado é pior que rótulo nenhum.
func matchQualitativeLevel(levels []models.ScoreLevel, text string) *int {
	value := normalizeQualitative(text)
	if value == "" {
		return nil
	}
	tokens := strings.Fields(value)

	search := func(target string, targetTokens []string) *int {
		var found *int
		for i := range levels {
			for _, alt := range levelNameAlternatives(levels[i].Name) {
				altTokens := strings.Fields(alt)
				if !(target == alt || containsPhrase(tokens, altTokens) || containsPhrase(targetTokens, altTokens)) {
					continue
				}
				if found != nil && *found != levels[i].Level {
					return nil // ambíguo
				}
				l := levels[i].Level
				found = &l
				break
			}
		}
		return found
	}

	if l := search(value, tokens); l != nil {
		return l
	}

	// Sinônimos: o laudo diz "Amostra NEGATIVA", o nível se chama "Não-reagente".
	for raw, canonical := range qualitativeSynonyms {
		if !containsPhrase(tokens, strings.Fields(raw)) {
			continue
		}
		if l := search(canonical, strings.Fields(canonical)); l != nil {
			return l
		}
	}
	return nil
}

// comparativePrefixes são as formas com que o laudo entrega um número sem entregar o número
// ("Superior a 1.000,0"). O valor do limite é o melhor palpite honesto para classificar.
var comparativePrefixes = []string{
	"superior a", "maior que", "maior do que", "acima de", ">=", ">",
	"inferior a", "menor que", "menor do que", "abaixo de", "<=", "<",
}

// numericFromComparativeText extrai o número de um resultado comparativo.
func numericFromComparativeText(text string) (float64, bool) {
	value := normalizeQualitative(text)
	for _, prefix := range comparativePrefixes {
		if !strings.HasPrefix(value, prefix) {
			continue
		}
		rest := strings.TrimSpace(strings.TrimPrefix(value, prefix))
		if rest == "" {
			continue
		}
		// Corta a unidade que às vezes vem colada ("1.000,0 mUI/mL").
		if fields := strings.Fields(rest); len(fields) > 0 {
			rest = fields[0]
		}
		if n, err := parseNumericResult(rest); err == nil {
			return n, true
		}
	}
	return 0, false
}
