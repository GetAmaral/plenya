package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

var (
	ErrProcessingJobNotFound = errors.New("processing job not found")
)

type ProcessingJobService struct {
	db                    *gorm.DB
	ocrService            *OCRService
	aiService             *AIService
	labTestDefService     *LabTestDefinitionService
	labResultBatchService *LabResultBatchService
	textCleaner           *PDFTextCleaner
	preMatchingService    *PreMatchingService
}

func NewProcessingJobService(
	db *gorm.DB,
	ocrService *OCRService,
	aiService *AIService,
	labTestDefService *LabTestDefinitionService,
	labResultBatchService *LabResultBatchService,
) *ProcessingJobService {
	return &ProcessingJobService{
		db:                    db,
		ocrService:            ocrService,
		aiService:             aiService,
		labTestDefService:     labTestDefService,
		labResultBatchService: labResultBatchService,
		textCleaner:           NewPDFTextCleaner(),
		preMatchingService:    NewPreMatchingService(db),
	}
}

// Create cria um novo job de processamento
func (s *ProcessingJobService) Create(batchID uuid.UUID, pdfPath string) (*dto.ProcessingJobResponse, error) {
	job := &models.ProcessingJob{
		LabResultBatchID: batchID,
		Type:             models.ProcessingJobTypePDFExtraction,
		PDFPath:          pdfPath,
		Status:           models.ProcessingJobPending,
		Attempts:         0,
		MaxAttempts:      3,
	}

	if err := s.db.Create(job).Error; err != nil {
		return nil, err
	}

	return s.toDTO(job), nil
}

// GetByID busca job por ID
func (s *ProcessingJobService) GetByID(jobID uuid.UUID) (*dto.ProcessingJobResponse, error) {
	var job models.ProcessingJob
	if err := s.db.First(&job, jobID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrProcessingJobNotFound
		}
		return nil, err
	}
	return s.toDTO(&job), nil
}

// PollAndProcess - padrão FOR UPDATE SKIP LOCKED para processar próximo job
// Retorna nil se não há jobs pendentes (comportamento normal)
func (s *ProcessingJobService) PollAndProcess() error {
	var job models.ProcessingJob

	// 1. Lock próximo job disponível (transação)
	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.
			Where("status = ?", models.ProcessingJobPending).
			Where("attempts < max_attempts").
			Where("deleted_at IS NULL").
			Order("created_at ASC").
			Limit(1).
			Clauses(clause.Locking{Strength: "UPDATE", Options: "SKIP LOCKED"}).
			First(&job).Error; err != nil {
			return err
		}

		// Atualizar status para processing
		now := time.Now()
		job.Status = models.ProcessingJobProcessing
		job.StartedAt = &now
		job.Attempts++

		return tx.Save(&job).Error
	})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Sem jobs pendentes (normal)
			return nil
		}
		return fmt.Errorf("failed to lock job: %v", err)
	}

	// 2. Processar job (fora da transação para não bloquear)
	if err := s.processJob(&job); err != nil {
		s.markJobFailed(&job, err)
		return err
	}

	s.markJobCompleted(&job)
	return nil
}

// processJob - workflow completo: OCR → IA → Save
func (s *ProcessingJobService) processJob(job *models.ProcessingJob) error {
	// 1. OCR: extrair texto do PDF
	fmt.Printf("🔍 [Job %s] Starting OCR extraction...\n", job.ID)
	ocrText, err := s.ocrService.ExtractText(job.PDFPath)
	if err != nil {
		return fmt.Errorf("OCR failed: %v", err)
	}

	// Salvar texto extraído no job
	job.ExtractedText = &ocrText
	s.db.Save(job)
	fmt.Printf("✅ [Job %s] OCR extracted %d chars\n", job.ID, len(ocrText))

	// 1.5. Limpar e processar o texto (remover ruído)
	fmt.Printf("🧹 [Job %s] Cleaning extracted text...\n", job.ID)
	cleanedText := s.textCleaner.CleanText(ocrText)
	stats := s.textCleaner.GetCompressionStats(ocrText, cleanedText)
	fmt.Printf("✅ [Job %s] Text cleaned: %d → %d chars (%.1f%% reduction)\n",
		job.ID, stats["originalChars"], stats["cleanedChars"], stats["reductionPct"])

	// 1.6. Salvar textos full e cleaned no batch
	if err := s.savePDFContentToBatch(job.LabResultBatchID, ocrText, cleanedText); err != nil {
		fmt.Printf("⚠️  [Job %s] Failed to save PDF content to batch: %v\n", job.ID, err)
		// Não falha o job, apenas loga o erro
	}

	// 1.7. PRÉ-MATCHING: Identificar e extrair exames conhecidos (sem IA)
	fmt.Printf("🔍 [Job %s] Pre-matching tests with altNames...\n", job.ID)
	preMatchResult, err := s.preMatchingService.PreMatch(cleanedText)
	if err != nil {
		fmt.Printf("⚠️  [Job %s] Pre-matching failed: %v\n", job.ID, err)
		// Continua sem pre-matching
		preMatchResult = &PreMatchingResult{
			RemainingText: cleanedText,
		}
	} else {
		fmt.Printf("✅ [Job %s] Pre-matched: %d tests found, %d estimated remaining\n",
			job.ID, preMatchResult.MatchedCount, preMatchResult.RemainingTestsEst)

		// Salvar exames matched (sem passar pela IA)
		if len(preMatchResult.MatchedTests) > 0 {
			if err := s.savePreMatchedTests(job.LabResultBatchID, preMatchResult); err != nil {
				fmt.Printf("⚠️  [Job %s] Failed to save pre-matched tests: %v\n", job.ID, err)
			}
		}

		// Atualizar metadata do batch (lab name, dates) se encontrado
		if preMatchResult.LaboratoryName != "" || preMatchResult.CollectionDate != "" {
			s.updateBatchMetadata(job.LabResultBatchID, preMatchResult)
		}
	}

	// Usar texto restante (após pre-matching) para IA
	textForAI := preMatchResult.RemainingText

	// Salvar pdfContentNeedAI no batch
	s.db.Model(&models.LabResultBatch{}).
		Where("id = ?", job.LabResultBatchID).
		Update("pdf_content_need_ai", textForAI)

	// Se não restou nada para IA (tudo foi matched), finalizar aqui
	if strings.TrimSpace(textForAI) == "" || preMatchResult.RemainingTestsEst == 0 {
		fmt.Printf("✅ [Job %s] All tests pre-matched! No AI needed.\n", job.ID)
		fmt.Printf("✅ [Job %s] Processing completed successfully\n", job.ID)
		return nil
	}

	fmt.Printf("📊 [Job %s] Sending to AI: %d chars (%d estimated tests)\n",
		job.ID, len(textForAI), preMatchResult.RemainingTestsEst)

	// 2. Buscar definições de exames para matching
	// OTIMIZAÇÃO: Enviar apenas exames mais comuns para reduzir tamanho do prompt
	fmt.Printf("📋 [Job %s] Loading common lab test definitions...\n", job.ID)
	testDefs, err := s.labTestDefService.GetAllLabTestDefinitions()
	if err != nil {
		return fmt.Errorf("failed to load test definitions: %v", err)
	}

	// Filtrar apenas exames comuns (hematologia, bioquímica) para reduzir tamanho
	// Isso reduz de 311 para ~100 exames mais relevantes
	testSummaries := make([]dto.LabTestSummary, 0, len(testDefs))
	for _, def := range testDefs {
		// Incluir apenas categorias mais comuns em laudos médicos
		category := string(def.Category)
		if category == "hematology" || category == "biochemistry" ||
		   category == "immunology" || category == "urinalysis" ||
		   category == "hormones" {
			testSummaries = append(testSummaries, dto.LabTestSummary{
				ID:        def.ID,
				Code:      def.Code,
				Name:      def.Name,
				ShortName: def.ShortName,
				TussCode:  def.TussCode,
				LoincCode: def.LoincCode,
				Category:  category,
				Unit:      def.Unit,
			})
		}
	}
	fmt.Printf("📋 [Job %s] Loaded %d common test definitions (filtered from %d total)\n",
		job.ID, len(testSummaries), len(testDefs))

	// 3. Claude API: interpretar laudo (usando texto que precisa de IA)
	fmt.Printf("🤖 [Job %s] Calling Claude API for interpretation...\n", job.ID)
	aiResp, err := s.aiService.InterpretLabResult(textForAI, testSummaries)
	if err != nil {
		return fmt.Errorf("AI interpretation failed: %v", err)
	}

	// Salvar resposta da IA (JSON)
	aiRespJSON, _ := json.Marshal(aiResp)
	aiRespStr := string(aiRespJSON)
	job.AIResponse = &aiRespStr
	s.db.Save(job)
	fmt.Printf("✅ [Job %s] AI extracted %d lab results\n", job.ID, len(aiResp.LabResults))

	// 4. Salvar resultados no batch
	fmt.Printf("💾 [Job %s] Saving results to batch...\n", job.ID)
	if err := s.saveResultsToBatch(job.LabResultBatchID, aiResp); err != nil {
		return fmt.Errorf("failed to save results: %v", err)
	}

	fmt.Printf("✅ [Job %s] Processing completed successfully\n", job.ID)
	return nil
}

// savePDFContentToBatch - salva textos full e cleaned no batch
func (s *ProcessingJobService) savePDFContentToBatch(
	batchID uuid.UUID,
	fullText, cleanedText string,
) error {
	return s.db.Model(&models.LabResultBatch{}).
		Where("id = ?", batchID).
		Updates(map[string]interface{}{
			"pdf_content_full":  fullText,
			"pdf_content_short": cleanedText,
		}).Error
}

// savePreMatchedTests - salva exames identificados no pré-matching
func (s *ProcessingJobService) savePreMatchedTests(
	batchID uuid.UUID,
	preMatchResult *PreMatchingResult,
) error {
	// Buscar batch para obter userID
	var batch models.LabResultBatch
	if err := s.db.Preload("Patient").First(&batch, batchID).Error; err != nil {
		return err
	}

	userID := batch.PatientID

	// Criar LabResult para cada match
	for _, match := range preMatchResult.MatchedTests {
		idStr := match.LabTestDefinitionID.String()
		matched := true

		createReq := &dto.CreateLabResultInBatchRequest{
			LabTestDefinitionID: &idStr,
			TestName:            match.TestName,
			TestType:            "pending", // Será atualizado
			ResultNumeric:       match.ResultNumeric,
			ResultText:          match.ResultText,
			Unit:                match.Unit,
			ReferenceRange:      match.ReferenceRange,
			Matched:             &matched,
		}

		if _, err := s.labResultBatchService.AddResult(batchID, userID, createReq); err != nil {
			fmt.Printf("⚠️  Failed to save pre-matched test '%s': %v\n", match.TestName, err)
			// Continua salvando os outros
		}
	}

	return nil
}

// updateBatchMetadata - atualiza metadata do batch com info extraída
func (s *ProcessingJobService) updateBatchMetadata(
	batchID uuid.UUID,
	preMatchResult *PreMatchingResult,
) error {
	updates := make(map[string]interface{})

	if preMatchResult.LaboratoryName != "" {
		updates["laboratory_name"] = preMatchResult.LaboratoryName
	}

	if preMatchResult.CollectionDate != "" {
		// Parse ISO date
		if collDate, err := time.Parse("2006-01-02", preMatchResult.CollectionDate); err == nil {
			updates["collection_date"] = collDate
		}
	}

	if len(updates) > 0 {
		return s.db.Model(&models.LabResultBatch{}).
			Where("id = ?", batchID).
			Updates(updates).Error
	}

	return nil
}

// saveResultsToBatch - salva resultados interpretados no batch
func (s *ProcessingJobService) saveResultsToBatch(
	batchID uuid.UUID,
	aiResp *dto.AILabResultExtractionResponse,
) error {
	// Buscar batch para obter userID (owner)
	var batch models.LabResultBatch
	if err := s.db.Preload("Patient.User").First(&batch, batchID).Error; err != nil {
		return fmt.Errorf("batch not found: %v", err)
	}

	// Parse datas
	collectionDate, err := time.Parse("2006-01-02", aiResp.CollectionDate)
	if err != nil {
		return fmt.Errorf("invalid collection date: %v", err)
	}

	var resultDate *time.Time
	if aiResp.ResultDate != "" {
		rd, err := time.Parse("2006-01-02", aiResp.ResultDate)
		if err != nil {
			return fmt.Errorf("invalid result date: %v", err)
		}
		resultDate = &rd
	}

	// Atualizar metadata do batch (se ainda não preenchido)
	updateBatch := false
	if batch.LaboratoryName == "" {
		batch.LaboratoryName = aiResp.LaboratoryName
		updateBatch = true
	}
	if batch.CollectionDate.IsZero() {
		batch.CollectionDate = collectionDate
		updateBatch = true
	}
	if batch.ResultDate == nil && resultDate != nil {
		batch.ResultDate = resultDate
		updateBatch = true
	}

	if updateBatch {
		s.db.Save(&batch)
	}

	// O userID é o mesmo que o patientID (patients são users)
	userID := batch.PatientID

	// Limpar resultados existentes antes de adicionar os novos do PDF
	// (remove placeholders ou resultados antigos)
	s.db.Where("lab_result_batch_id = ?", batchID).Delete(&models.LabResult{})

	for _, aiResult := range aiResp.LabResults {
		var labTestDefIDStr *string
		if aiResult.LabTestDefinitionID != nil {
			idStr := aiResult.LabTestDefinitionID.String()
			labTestDefIDStr = &idStr
		}

		createReq := &dto.CreateLabResultInBatchRequest{
			LabTestDefinitionID: labTestDefIDStr,
			TestName:            aiResult.TestName,
			TestType:            aiResult.TestType,
			ResultText:          aiResult.ResultText,
			ResultNumeric:       aiResult.ResultNumeric,
			Unit:                aiResult.Unit,
			ReferenceRange:      aiResult.ReferenceRange,
			Interpretation:      aiResult.Interpretation,
			Matched:             &aiResult.Matched,
		}

		if _, err := s.labResultBatchService.AddResult(batchID, userID, createReq); err != nil {
			return fmt.Errorf("failed to add result '%s': %v", aiResult.TestName, err)
		}
	}

	return nil
}

// markJobCompleted - marca job como concluído
func (s *ProcessingJobService) markJobCompleted(job *models.ProcessingJob) {
	now := time.Now()
	job.Status = models.ProcessingJobCompleted
	job.CompletedAt = &now
	s.db.Save(job)
}

// markJobFailed - marca job como falho (ou pending para retry se < max attempts)
func (s *ProcessingJobService) markJobFailed(job *models.ProcessingJob, err error) {
	errMsg := err.Error()
	job.ErrorMessage = &errMsg

	if job.Attempts >= job.MaxAttempts {
		// Atingiu max attempts, marcar como failed definitivo
		job.Status = models.ProcessingJobFailed
		now := time.Now()
		job.CompletedAt = &now
		fmt.Printf("❌ [Job %s] Failed after %d attempts: %v\n", job.ID, job.Attempts, err)
	} else {
		// Ainda há tentativas, marcar como pending para retry
		job.Status = models.ProcessingJobPending
		fmt.Printf("⚠️  [Job %s] Failed (attempt %d/%d), will retry: %v\n",
			job.ID, job.Attempts, job.MaxAttempts, err)
	}

	s.db.Save(job)
}

// toDTO converte model para DTO
func (s *ProcessingJobService) toDTO(job *models.ProcessingJob) *dto.ProcessingJobResponse {
	return &dto.ProcessingJobResponse{
		ID:               job.ID,
		LabResultBatchID: job.LabResultBatchID,
		Type:             job.Type,
		Status:           job.Status,
		ErrorMessage:     job.ErrorMessage,
		Attempts:         job.Attempts,
		CreatedAt:        job.CreatedAt,
		StartedAt:        job.StartedAt,
		CompletedAt:      job.CompletedAt,
	}
}
