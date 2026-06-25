package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/google/uuid"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
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
	step := models.StepUploadingPDF
	message := "Enviando PDF ao nosso servidor"

	job := &models.ProcessingJob{
		LabResultBatchID: batchID,
		Type:             models.ProcessingJobTypePDFExtraction,
		PDFPath:          pdfPath,
		Status:           models.ProcessingJobPending,
		ProgressStep:     &step,
		ProgressMessage:  &message,
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

// updateProgress - atualiza step e mensagem do job (para comunicação com frontend)
func (s *ProcessingJobService) updateProgress(job *models.ProcessingJob, step int, message string) {
	job.ProgressStep = &step
	job.ProgressMessage = &message
	s.db.Save(job)
	fmt.Printf("📊 [Job %s] Step %d: %s\n", job.ID, step, message)
}

// processJob - workflow completo: OCR → IA → Save
func (s *ProcessingJobService) processJob(job *models.ProcessingJob) error {
	// Step 2: Upload completo (Step 1 seria no handler de upload)
	s.updateProgress(job, models.StepUploadComplete, "Upload completo")

	// Step 3: Extraindo texto do PDF
	s.updateProgress(job, models.StepExtractingText, "Extraindo conteúdo do PDF para texto")
	fmt.Printf("🔍 [Job %s] Starting OCR extraction...\n", job.ID)
	ocrText, err := s.ocrService.ExtractText(job.PDFPath)
	if err != nil {
		return fmt.Errorf("OCR failed: %v", err)
	}

	// Salvar texto extraído no job
	job.ExtractedText = &ocrText
	s.db.Save(job)
	fmt.Printf("✅ [Job %s] OCR extracted %d chars\n", job.ID, len(ocrText))

	// Limpar e processar o texto (remover ruído)
	fmt.Printf("🧹 [Job %s] Cleaning extracted text...\n", job.ID)
	cleanedText := s.textCleaner.CleanText(ocrText)
	stats := s.textCleaner.GetCompressionStats(ocrText, cleanedText)
	fmt.Printf("✅ [Job %s] Text cleaned: %d → %d chars (%.1f%% reduction)\n",
		job.ID, stats["originalChars"], stats["cleanedChars"], stats["reductionPct"])

	// Salvar textos full e cleaned no batch
	if err := s.savePDFContentToBatch(job.LabResultBatchID, ocrText, cleanedText); err != nil {
		fmt.Printf("⚠️  [Job %s] Failed to save PDF content to batch: %v\n", job.ID, err)
		// Não falha o job, apenas loga o erro
	}

	// Step 4: Analisando com IA
	s.updateProgress(job, models.StepAnalyzingWithAI, "Analisando conteúdo com IA")
	fmt.Printf("🤖 [Job %s] Calling Claude API for extraction...\n", job.ID)
	jsonStr, err := s.aiService.InterpretLabResult(cleanedText)
	if err != nil {
		return fmt.Errorf("AI interpretation failed: %v", err)
	}
	fmt.Printf("✅ [Job %s] AI extracted data (JSON length: %d chars)\n", job.ID, len(jsonStr))

	// Step 5: Análise concluída - contar exames
	var parsed struct {
		Exames []interface{} `json:"exames"`
	}
	if err := json.Unmarshal([]byte(jsonStr), &parsed); err == nil {
		examCount := len(parsed.Exames)
		message := fmt.Sprintf("Conteúdo analisado pela IA - %d exames identificados", examCount)
		s.updateProgress(job, models.StepAIComplete, message)
	} else {
		s.updateProgress(job, models.StepAIComplete, "Conteúdo analisado pela IA")
	}

	// Step 6: Salvando resultados
	s.updateProgress(job, models.StepSavingResults, "Salvando resultados no prontuário")
	if err := s.savePDFContentJSON(job.LabResultBatchID, jsonStr); err != nil {
		return fmt.Errorf("failed to save PDF content JSON: %v", err)
	}

	// Preencher laboratório + data de coleta do lote a partir do que a IA extraiu do PDF
	// (o PDF é a fonte de verdade num import; o usuário ainda pode editar depois).
	s.applyExtractedMetadata(job.LabResultBatchID, jsonStr)

	// Criar LabResults a partir do JSON extraído
	fmt.Printf("🔗 [Job %s] Creating lab results from extracted data...\n", job.ID)
	matchedCount, unmatchedCount, err := s.createLabResultsFromJSON(job.LabResultBatchID, jsonStr)
	if err != nil {
		fmt.Printf("⚠️  [Job %s] Failed to create lab results: %v\n", job.ID, err)
		// Não falha o job, apenas loga o erro
	} else {
		fmt.Printf("✅ [Job %s] Created %d matched + %d unmatched lab results\n", job.ID, matchedCount, unmatchedCount)
	}

	// Step 7: Classificar automaticamente (atribui Level 0-5 + criticidade via ScoreItems).
	// Sem isso, o lote importava mas ficava "não classificado" até o usuário clicar em
	// "Classificar" manualmente. Best-effort: não falha o job se a classificação falhar.
	s.updateProgress(job, models.StepSavingResults, "Classificando resultados")
	if err := s.labResultBatchService.ClassifyBatchResults(job.LabResultBatchID); err != nil {
		fmt.Printf("⚠️  [Job %s] Classify failed: %v\n", job.ID, err)
	} else {
		fmt.Printf("✅ [Job %s] Results classified\n", job.ID)
	}

	fmt.Printf("✅ [Job %s] Processing completed successfully\n", job.ID)
	return nil
}

// applyExtractedMetadata preenche laboratório + data de coleta do lote a partir do JSON
// extraído pela IA. Best-effort: só atualiza os campos que vieram preenchidos/parseáveis.
func (s *ProcessingJobService) applyExtractedMetadata(batchID uuid.UUID, jsonStr string) {
	var extracted dto.PDFExtractionResponse
	if err := json.Unmarshal([]byte(jsonStr), &extracted); err != nil {
		return
	}
	updates := map[string]interface{}{}
	if extracted.Laboratorio != nil {
		if lab := strings.TrimSpace(*extracted.Laboratorio); lab != "" {
			updates["laboratory_name"] = lab
		}
	}
	if extracted.DataColeta != nil {
		if d, err := parseFlexibleDate(strings.TrimSpace(*extracted.DataColeta)); err == nil {
			updates["collection_date"] = d
		}
	}
	if len(updates) == 0 {
		return
	}
	if err := s.db.Model(&models.LabResultBatch{}).Where("id = ?", batchID).Updates(updates).Error; err != nil {
		fmt.Printf("⚠️  [Job] apply extracted metadata batch=%s: %v\n", batchID, err)
	}
}

// ReprocessResultsFromJSON recria os LabResults de um lote a partir do PDFContentJSON
// já armazenado (sem refazer OCR/IA) e re-classifica. Útil para um botão "reprocessar"
// e para recuperar um lote cujos resultados foram perdidos. Retorna (matched, unmatched).
func (s *ProcessingJobService) ReprocessResultsFromJSON(batchID uuid.UUID) (int, int, error) {
	var batch models.LabResultBatch
	if err := s.db.First(&batch, batchID).Error; err != nil {
		return 0, 0, fmt.Errorf("batch %s: %w", batchID, err)
	}
	if batch.PDFContentJSON == nil || *batch.PDFContentJSON == "" {
		return 0, 0, fmt.Errorf("batch %s sem PDFContentJSON armazenado", batchID)
	}
	matched, unmatched, err := s.createLabResultsFromJSON(batchID, *batch.PDFContentJSON)
	if err != nil {
		return 0, 0, err
	}
	if err := s.labResultBatchService.ClassifyBatchResults(batchID); err != nil {
		fmt.Printf("⚠️  [Reprocess %s] classify: %v\n", batchID, err)
	}
	return matched, unmatched, nil
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

// savePDFContentJSON - salva JSON extraído pela IA no batch
func (s *ProcessingJobService) savePDFContentJSON(
	batchID uuid.UUID,
	jsonContent string,
) error {
	return s.db.Model(&models.LabResultBatch{}).
		Where("id = ?", batchID).
		Update("pdf_content_json", jsonContent).Error
}

// createLabResultsFromJSON - cria LabResults a partir do JSON extraído pela IA
// Retorna: (matchedCount, unmatchedCount, error)
func (s *ProcessingJobService) createLabResultsFromJSON(
	batchID uuid.UUID,
	jsonContent string,
) (int, int, error) {
	// Parse JSON
	var extracted dto.PDFExtractionResponse
	if err := json.Unmarshal([]byte(jsonContent), &extracted); err != nil {
		return 0, 0, fmt.Errorf("failed to parse JSON: %v", err)
	}

	// Buscar todas as definições de testes para matching
	var testDefinitions []models.LabTestDefinition
	if err := s.db.Find(&testDefinitions).Error; err != nil {
		return 0, 0, fmt.Errorf("failed to load test definitions: %v", err)
	}

	// Criar mapa para busca rápida
	testDefMap := s.buildTestDefinitionMap(testDefinitions)

	matchedCount := 0
	unmatchedCount := 0

	// Deletar resultados existentes do batch (se houver)
	s.db.Where("lab_result_batch_id = ?", batchID).Delete(&models.LabResult{})

	// Criar LabResult para cada exame extraído
	for _, exam := range extracted.Exames {
		// Tentar fazer match com definição de teste
		testDefID := s.matchTestDefinition(exam.NomeExame, testDefMap)

		// Tentar converter resultado para numérico primeiro
		var resultNumeric *float64
		var resultText *string
		if numeric, err := parseNumericResult(exam.Resultado); err == nil {
			// Se for numérico, salvar apenas como numérico (não duplicar no texto)
			resultNumeric = &numeric
			resultText = nil
		} else {
			// Se NÃO for numérico, salvar apenas como texto
			resultText = &exam.Resultado
			resultNumeric = nil
		}

		// Se tiver match com definição, deixar TestName e TestType vazios
		testName := exam.NomeExame
		testType := "other"
		if testDefID != nil {
			testName = "" // Vazio quando linkado
			testType = "" // Vazio quando linkado
		}

		// Criar LabResult usando AddResultInternal para aplicar conversão de unidades
		matched := testDefID != nil
		req := &dto.CreateLabResultInBatchRequest{
			TestName:      testName,
			TestType:      testType,
			ResultNumeric: resultNumeric,
			ResultText:    resultText,
			Unit:          exam.Unidade,
			Matched:       &matched,
		}

		if testDefID != nil {
			testDefIDStr := testDefID.String()
			req.LabTestDefinitionID = &testDefIDStr
		}

		// Usar AddResultInternal que aplica conversão automática
		if _, err := s.labResultBatchService.AddResultInternal(batchID, req); err != nil {
			fmt.Printf("⚠️  Failed to create lab result for '%s': %v\n", exam.NomeExame, err)
			continue
		}

		if testDefID != nil {
			matchedCount++
		} else {
			unmatchedCount++
		}
	}

	return matchedCount, unmatchedCount, nil
}

// buildTestDefinitionMap - cria mapa de nomes normalizados para IDs de definições
func (s *ProcessingJobService) buildTestDefinitionMap(testDefs []models.LabTestDefinition) map[string]uuid.UUID {
	defMap := make(map[string]uuid.UUID)

	for _, def := range testDefs {
		// Adicionar nome principal
		normalizedName := normalizeTestName(def.Name)
		defMap[normalizedName] = def.ID

		// Adicionar nome curto
		if def.ShortName != nil && *def.ShortName != "" {
			normalizedShort := normalizeTestName(*def.ShortName)
			defMap[normalizedShort] = def.ID
		}

		// Adicionar nomes alternativos (altNames)
		for _, altName := range def.AltNames {
			if altName != "" {
				normalizedAlt := normalizeTestName(altName)
				defMap[normalizedAlt] = def.ID
			}
		}
	}

	return defMap
}

// matchTestDefinition - busca definição de teste usando nome normalizado
func (s *ProcessingJobService) matchTestDefinition(
	examName string,
	defMap map[string]uuid.UUID,
) *uuid.UUID {
	normalizedName := normalizeTestName(examName)

	// Busca exata
	if id, found := defMap[normalizedName]; found {
		return &id
	}

	// Busca parcial (se nome tem pelo menos 5 caracteres)
	if len(normalizedName) >= 5 {
		for defName, id := range defMap {
			// Se o nome do exame contém o nome da definição OU vice-versa
			if len(defName) >= 5 && (containsSubstring(normalizedName, defName) || containsSubstring(defName, normalizedName)) {
				return &id
			}
		}
	}

	return nil
}

// normalizeTestName - normaliza nome de teste para matching
func normalizeTestName(name string) string {
	// Converter para minúsculas
	name = strings.ToLower(name)

	// Remover acentos
	name = removeAccentsFromString(name)

	// Substituir hífens, vírgulas e outros separadores por espaços ANTES de remover caracteres especiais
	name = strings.Map(func(r rune) rune {
		if r == '-' || r == ',' || r == '/' || r == '(' || r == ')' || r == ':' {
			return ' '
		}
		return r
	}, name)

	// Remover caracteres especiais mantendo apenas letras, dígitos e espaços
	name = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == ' ' {
			return r
		}
		return -1
	}, name)

	// Remover espaços duplicados e normalizar
	name = strings.Join(strings.Fields(name), " ")

	return strings.TrimSpace(name)
}

// removeAccentsFromString - remove acentos de uma string
func removeAccentsFromString(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, s)
	return result
}

// containsSubstring - verifica se s contém substr (ambos já normalizados)
func containsSubstring(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 && len(s) >= len(substr) &&
		(s == substr ||
			strings.HasPrefix(s, substr) ||
			strings.HasSuffix(s, substr) ||
			strings.Contains(s, substr))
}

// parseNumericResult - tenta converter resultado textual para numérico
func parseNumericResult(result string) (float64, error) {
	// Remover espaços
	result = strings.TrimSpace(result)

	// Substituir vírgula por ponto (padrão brasileiro)
	result = strings.ReplaceAll(result, ",", ".")

	// Remover caracteres não numéricos (exceto ponto e sinal)
	cleaned := ""
	for _, r := range result {
		if unicode.IsDigit(r) || r == '.' || r == '-' || r == '+' {
			cleaned += string(r)
		}
	}

	// Tentar parsear
	if cleaned == "" {
		return 0, fmt.Errorf("no numeric value found")
	}

	return strconv.ParseFloat(cleaned, 64)
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
		ProgressStep:     job.ProgressStep,
		ProgressMessage:  job.ProgressMessage,
		ErrorMessage:     job.ErrorMessage,
		Attempts:         job.Attempts,
		CreatedAt:        job.CreatedAt,
		StartedAt:        job.StartedAt,
		CompletedAt:      job.CompletedAt,
	}
}
