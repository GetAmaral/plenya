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

// fatalJobError marca uma falha determinística — repetir a mesma chamada daria o mesmo
// erro. markJobFailed encerra o job na hora em vez de queimar as 3 tentativas (e os
// tokens) mostrando o mesmo erro ao usuário três vezes.
type fatalJobError struct{ err error }

func (e fatalJobError) Error() string { return e.err.Error() }
func (e fatalJobError) Unwrap() error { return e.err }

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
		return fmt.Errorf("Não foi possível extrair o texto do PDF. Confirme que o arquivo não está protegido por senha nem corrompido. (%v)", err)
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

	// Step 4: Analisando com IA. Inclui o CABEÇALHO do OCR bruto (o cleaner remove o
	// letterhead, onde fica o nome do laboratório e a data) para a IA conseguir extrair
	// laboratorio/dataColeta, além dos exames do texto limpo.
	s.updateProgress(job, models.StepAnalyzingWithAI, "Analisando conteúdo com IA")
	fmt.Printf("🤖 [Job %s] Calling Claude API for extraction...\n", job.ID)
	forAI := "## CABEÇALHO DO LAUDO (use para laboratorio e dataColeta)\n" +
		firstRunes(ocrText, 1200) +
		"\n\n## CONTEÚDO (use para os exames)\n" + cleanedText
	jsonStr, err := s.aiService.InterpretLabResult(forAI)
	if err != nil {
		// Truncagem é determinística: repetir a mesma chamada dá o mesmo resultado.
		if errors.Is(err, ErrAITruncated) {
			return fatalJobError{errors.New(
				"Laudo extenso demais para ser interpretado de uma vez. Divida o PDF em partes menores (por exemplo, um arquivo por página/painel) e envie novamente.")}
		}
		return fmt.Errorf("Falha ao interpretar o laudo com a IA. Tente novamente em alguns minutos. (%v)", err)
	}
	fmt.Printf("✅ [Job %s] AI extracted data (JSON length: %d chars)\n", job.ID, len(jsonStr))

	// Step 5: Análise concluída - contar exames
	var parsed struct {
		Exames []interface{} `json:"exames"`
	}
	if err := json.Unmarshal([]byte(jsonStr), &parsed); err != nil {
		return fmt.Errorf("A IA devolveu uma resposta inválida ao ler o laudo. Tente novamente. (%v)", err)
	}
	examCount := len(parsed.Exames)
	if examCount == 0 {
		// Extração VAZIA não é sucesso: força falha/retry em vez de marcar "completed"
		// silenciosamente (causas: OCR sem texto, hiccup da API, saldo de créditos esgotado).
		return fmt.Errorf("A IA não identificou nenhum exame neste PDF. Verifique se o arquivo é um laudo com resultados legíveis (PDF digitalizado sem texto pode não ser lido).")
	}
	s.updateProgress(job, models.StepAIComplete, fmt.Sprintf("Conteúdo analisado pela IA - %d exames identificados", examCount))

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

// firstRunes devolve os primeiros n runes de s (truncamento seguro p/ UTF-8).
func firstRunes(s string, n int) string {
	r := []rune(s)
	if len(r) <= n {
		return s
	}
	return string(r[:n])
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
			// Ignora data no futuro (extração provavelmente errada).
			if !d.After(time.Now().Add(24 * time.Hour)) {
				updates["collection_date"] = d
			}
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

	// Criar mapa para busca rápida + mapa de espécime por definição (desambiguação no match)
	testDefMap := s.buildTestDefinitionMap(testDefinitions)
	specByID := make(map[uuid.UUID]string, len(testDefinitions))
	for _, def := range testDefinitions {
		if def.SpecimenType != nil && *def.SpecimenType != "" {
			specByID[def.ID] = *def.SpecimenType
		}
	}

	matchedCount := 0
	unmatchedCount := 0

	// Deletar resultados existentes do batch (se houver)
	s.db.Where("lab_result_batch_id = ?", batchID).Delete(&models.LabResult{})

	// Criar LabResult para cada exame extraído
	for _, exam := range extracted.Exames {
		// Espécime normalizado (Sangue/Urina/...) — usado no match e gravado no result
		specimen := normalizeSpecimen(exam.Material)
		// Tentar fazer match com definição de teste (desempatando por espécime)
		testDefID := s.matchTestDefinition(exam.NomeExame, specimen, testDefMap, specByID)

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
		source := "pdf"
		var matchReason *string
		if !matched {
			r := "Não encontrado no catálogo de exames"
			matchReason = &r
		}
		req := &dto.CreateLabResultInBatchRequest{
			TestName:       testName,
			TestType:       testType,
			ResultNumeric:  resultNumeric,
			ResultText:     resultText,
			Unit:           exam.Unidade,
			Matched:        &matched,
			Source:         &source,
			MatchReason:    matchReason,
			Specimen:       specimen,
			Method:         exam.Metodo,
			ReferenceRange: exam.ValorReferencia,
			CollectionDate: parseExamDate(exam.DataColetaExame),
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

// matchTestDefinition - busca a melhor definição para um exame. Quando há mais de um candidato
// e o resultado tem espécime conhecido, prefere o candidato cujo specimen_type casa — desambigua
// exames de mesmo nome em espécimes diferentes (ex.: glicose sangue vs urina). `specByID` mapeia
// id da definição -> specimen_type; `specimen` é o espécime normalizado do resultado.
func (s *ProcessingJobService) matchTestDefinition(
	examName string,
	specimen *string,
	defMap map[string]uuid.UUID,
	specByID map[uuid.UUID]string,
) *uuid.UUID {
	normalizedName := normalizeTestName(examName)

	// Busca exata (chave única → 1 id)
	if id, found := defMap[normalizedName]; found {
		// Se o resultado tem espécime e o match exato CONFLITA com ele (ex.: "Hemácias" da
		// urina batendo na definição de hemácias do SANGUE), procura uma definição do espécime
		// certo entre os candidatos por substring antes de aceitar o exato.
		if specimen != nil && *specimen != "" {
			if st, ok := specByID[id]; ok && st != "" && !strings.EqualFold(st, *specimen) {
				if best := preferBySpecimen(gatherSubstringCandidates(normalizedName, defMap), specimen, specByID); best != nil {
					return best
				}
			}
		}
		return &id
	}

	// Busca parcial (substring): coleta TODOS os candidatos e desempata por espécime.
	if cands := gatherSubstringCandidates(normalizedName, defMap); len(cands) > 0 {
		if best := preferBySpecimen(cands, specimen, specByID); best != nil {
			return best
		}
		return &cands[0]
	}

	// Busca fuzzy (fallback): tolera typos do laudo (ex.: "trigliceridios" vs
	// "trigliceridos"). Salvaguardas contra falso-positivo: só nomes >=8 chars,
	// distância de edição <=1 (<=2 p/ nomes longos), e match ÚNICO na menor distância.
	if len(normalizedName) >= 8 {
		maxDist := 1
		if len(normalizedName) > 15 {
			maxDist = 2
		}
		bestDist := maxDist + 1
		var bestID uuid.UUID
		tie := false
		for defName, id := range defMap {
			if len(defName) < 8 {
				continue
			}
			if dl := len(defName) - len(normalizedName); dl > maxDist || dl < -maxDist {
				continue // poda: diferença de tamanho já excede o limite
			}
			d := levenshtein(normalizedName, defName)
			if d < bestDist {
				bestDist, bestID, tie = d, id, false
			} else if d == bestDist {
				tie = true
			}
		}
		if bestDist <= maxDist && !tie {
			return &bestID
		}
	}

	return nil
}

// gatherSubstringCandidates - coleta todas as definições cujo nome é substring do nome do exame
// (ou vice-versa), deduplicadas por id. Base para a desambiguação por espécime.
func gatherSubstringCandidates(normalizedName string, defMap map[string]uuid.UUID) []uuid.UUID {
	var cands []uuid.UUID
	if len(normalizedName) < 5 {
		return cands
	}
	seen := map[uuid.UUID]bool{}
	for defName, id := range defMap {
		if len(defName) >= 5 && (containsSubstring(normalizedName, defName) || containsSubstring(defName, normalizedName)) {
			if !seen[id] {
				cands = append(cands, id)
				seen[id] = true
			}
		}
	}
	return cands
}

// preferBySpecimen - entre candidatos de match, devolve o ÚNICO cujo specimen_type casa com o
// espécime do resultado. nil se não há espécime, nenhum casa, ou há empate (mantém ambíguo).
func preferBySpecimen(cands []uuid.UUID, specimen *string, specByID map[uuid.UUID]string) *uuid.UUID {
	if specimen == nil || *specimen == "" {
		return nil
	}
	var matches []uuid.UUID
	for _, id := range cands {
		if specByID != nil && strings.EqualFold(specByID[id], *specimen) {
			matches = append(matches, id)
		}
	}
	if len(matches) == 1 {
		return &matches[0]
	}
	return nil
}

// levenshtein - distância de edição entre duas strings (já normalizadas).
func levenshtein(a, b string) int {
	ra, rb := []rune(a), []rune(b)
	la, lb := len(ra), len(rb)
	if la == 0 {
		return lb
	}
	if lb == 0 {
		return la
	}
	prev := make([]int, lb+1)
	curr := make([]int, lb+1)
	for j := 0; j <= lb; j++ {
		prev[j] = j
	}
	for i := 1; i <= la; i++ {
		curr[0] = i
		for j := 1; j <= lb; j++ {
			cost := 1
			if ra[i-1] == rb[j-1] {
				cost = 0
			}
			curr[j] = min(min(prev[j]+1, curr[j-1]+1), prev[j-1]+cost)
		}
		prev, curr = curr, prev
	}
	return prev[lb]
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

// normalizeSpecimen - normaliza o material/espécime extraído do laudo para uma categoria
// canônica (Sangue/Urina/Fezes/Saliva/Líquor), usada na disambiguação de matching e no
// prontuário. Família "sangue" agrega sangue/soro/plasma. Desconhecido preserva o original.
func normalizeSpecimen(material *string) *string {
	if material == nil {
		return nil
	}
	m := removeAccentsFromString(strings.ToLower(strings.TrimSpace(*material)))
	if m == "" {
		return nil
	}
	canon := func(s string) *string { return &s }
	switch {
	case strings.Contains(m, "urin"):
		return canon("Urina")
	case strings.Contains(m, "fezes") || strings.Contains(m, "fecal"):
		return canon("Fezes")
	case strings.Contains(m, "saliva"):
		return canon("Saliva")
	case strings.Contains(m, "liquor") || strings.Contains(m, "lcr"):
		return canon("Líquor")
	case strings.Contains(m, "sangue") || strings.Contains(m, "soro") || strings.Contains(m, "plasm") || strings.Contains(m, "serum"):
		return canon("Sangue")
	default:
		return material // desconhecido: preserva o que veio
	}
}

// parseExamDate - converte data de coleta por-exame (YYYY-MM-DD) em time.Time; nil se
// ausente/inválida. Armazena ao MEIO-DIA UTC para que a exibição em qualquer fuso (±12h)
// mantenha o dia correto (meia-noite UTC voltaria 1 dia em São Paulo, -3).
func parseExamDate(d *string) *time.Time {
	if d == nil || strings.TrimSpace(*d) == "" {
		return nil
	}
	t, err := time.Parse("2006-01-02", strings.TrimSpace(*d))
	if err != nil {
		return nil
	}
	t = t.Add(12 * time.Hour)
	return &t
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

	var fatal fatalJobError
	if errors.As(err, &fatal) {
		job.Attempts = job.MaxAttempts
	}

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
