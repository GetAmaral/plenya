package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

// AnonymousLabPDFService - extrai valores de PDFs de exames para o Escore Light.
//
// Reusa todo o pipeline do EMR (OCR + Cleaner + AIService) mas:
//   - É público (sem auth, sem Patient)
//   - NÃO persiste o PDF (LGPD: arquivo é descartado após extração)
//   - Filtra match apenas para o subset de items Light com lab_test_code
//   - Operação síncrona (usuário aguarda na tela ~10s)
type AnonymousLabPDFService struct {
	db          *gorm.DB
	ocrService  *OCRService
	textCleaner *PDFTextCleaner
	aiService   *AIService
}

func NewAnonymousLabPDFService(db *gorm.DB, ocr *OCRService, cleaner *PDFTextCleaner, ai *AIService) *AnonymousLabPDFService {
	return &AnonymousLabPDFService{
		db:          db,
		ocrService:  ocr,
		textCleaner: cleaner,
		aiService:   ai,
	}
}

// LightLabMatch — um exame extraído e mapeado para um item Light.
type LightLabMatch struct {
	ScoreItemID  string  `json:"scoreItemId"`
	ItemName     string  `json:"itemName"`     // nome humano do item Light
	NumericValue float64 `json:"numericValue"` // valor numérico final (já convertido se necessário)
	OriginalRaw  string  `json:"originalRaw"`  // resultado bruto do PDF (debug/transparência)
	Unit         string  `json:"unit,omitempty"`
}

// LightLabUnmatched — exame extraído do PDF que NÃO casou com nenhum item Light.
type LightLabUnmatched struct {
	NomeExame string `json:"nomeExame"`
	Resultado string `json:"resultado"`
	Unidade   string `json:"unidade,omitempty"`
}

// LightLabExtractionResult — payload retornado ao frontend do Light.
type LightLabExtractionResult struct {
	Matched   []LightLabMatch     `json:"matched"`
	Unmatched []LightLabUnmatched `json:"unmatched"`
	Warnings  []string            `json:"warnings,omitempty"`
}

// ExtractFromPDF — pipeline completo:
//  1. OCR
//  2. Limpeza
//  3. AI Claude extrai exames
//  4. Match contra subset Light
//  5. Conversão de unidade quando necessário
//
// IMPORTANTE: o caller deve garantir a deleção do arquivo após esta chamada
// (apagamos defensivamente aqui também, mas o handler é o owner do path).
func (s *AnonymousLabPDFService) ExtractFromPDF(pdfPath string) (*LightLabExtractionResult, error) {
	if _, err := os.Stat(pdfPath); err != nil {
		return nil, fmt.Errorf("PDF não encontrado: %w", err)
	}

	// Apaga defensivamente ao sair (LGPD)
	defer func() {
		_ = os.Remove(pdfPath)
	}()

	// 1) OCR
	ocrText, err := s.ocrService.ExtractText(pdfPath)
	if err != nil {
		return nil, fmt.Errorf("OCR falhou: %w", err)
	}
	if strings.TrimSpace(ocrText) == "" {
		return nil, errors.New("PDF não tem texto extraível")
	}

	// 2) Limpeza
	cleaned := s.textCleaner.CleanText(ocrText)

	// 3) IA → JSON estruturado
	jsonStr, err := s.aiService.InterpretLabResult(cleaned)
	if err != nil {
		return nil, fmt.Errorf("interpretação por IA falhou: %w", err)
	}

	var extracted dto.PDFExtractionResponse
	if err := json.Unmarshal([]byte(jsonStr), &extracted); err != nil {
		return nil, fmt.Errorf("falha ao decodificar JSON da IA: %w", err)
	}

	// 4) Carrega items Light com lab_test_code + suas LabTestDefinition (para conversão de unidade)
	type lightLabRef struct {
		ScoreItemID    uuid.UUID
		ScoreItemName  string
		LabTestCode    string
		LabTestDefID   uuid.UUID
		LabTestName    string
		LabTestUnit    *string
		LabTestAltLst  []string
		ScoreItemLight bool
	}

	var lightItems []models.ScoreItem
	if err := s.db.
		Where("lab_test_code IS NOT NULL AND deleted_at IS NULL").
		Where("id IN (?)", s.db.
			Table("score_version_items AS svi").
			Select("svi.score_item_id").
			Joins("JOIN score_versions sv ON sv.id = svi.version_id").
			Where("sv.slug = ? AND sv.deleted_at IS NULL", "light")).
		Find(&lightItems).Error; err != nil {
		return nil, fmt.Errorf("falha ao carregar items Light: %w", err)
	}

	// Carrega definitions referenciadas
	codes := make([]string, 0, len(lightItems))
	for _, it := range lightItems {
		if it.LabTestCode != nil && *it.LabTestCode != "" {
			codes = append(codes, *it.LabTestCode)
		}
	}
	var defs []models.LabTestDefinition
	if len(codes) > 0 {
		if err := s.db.Where("code IN ?", codes).Find(&defs).Error; err != nil {
			return nil, fmt.Errorf("falha ao carregar lab definitions: %w", err)
		}
	}
	defByCode := make(map[string]*models.LabTestDefinition, len(defs))
	for i := range defs {
		defByCode[defs[i].Code] = &defs[i]
	}

	// Constrói índice de matching: nome normalizado → ScoreItem (Light)
	type lightCandidate struct {
		Item *models.ScoreItem
		Def  *models.LabTestDefinition
	}
	matchIndex := make(map[string]*lightCandidate)
	for i := range lightItems {
		it := &lightItems[i]
		def := defByCode[strings.TrimSpace(*it.LabTestCode)]
		if def == nil {
			continue
		}
		cand := &lightCandidate{Item: it, Def: def}

		// Indexa pelo nome do score item, nome da definição e altNames
		matchIndex[normalizeTestName(it.Name)] = cand
		matchIndex[normalizeTestName(def.Name)] = cand
		if def.ShortName != nil {
			matchIndex[normalizeTestName(*def.ShortName)] = cand
		}
		for _, alt := range def.AltNames {
			if alt != "" {
				matchIndex[normalizeTestName(alt)] = cand
			}
		}
	}

	result := &LightLabExtractionResult{
		Matched:   []LightLabMatch{},
		Unmatched: []LightLabUnmatched{},
	}

	// 5) Match cada exame extraído contra o subset Light
	seen := make(map[uuid.UUID]bool) // evita duplicar matches do mesmo item Light
	for _, exam := range extracted.Exames {
		nomeNorm := normalizeTestName(exam.NomeExame)
		cand, ok := matchIndex[nomeNorm]
		if !ok {
			// fuzzy: tenta substring com nomes >=5 chars
			if len(nomeNorm) >= 5 {
				for k, c := range matchIndex {
					if strings.Contains(nomeNorm, k) || strings.Contains(k, nomeNorm) {
						cand = c
						ok = true
						break
					}
				}
			}
		}
		if !ok {
			result.Unmatched = append(result.Unmatched, LightLabUnmatched{
				NomeExame: exam.NomeExame,
				Resultado: exam.Resultado,
				Unidade:   strDeref(exam.Unidade),
			})
			continue
		}
		if seen[cand.Item.ID] {
			continue // já temos um valor para esse item Light
		}

		// Parse numérico (já existe parseNumericResult em processing_job_service.go)
		numeric, err := parseNumericResult(exam.Resultado)
		if err != nil {
			// Se vem texto não-numérico, tenta extrair primeiro número
			numeric, err = extractFirstNumber(exam.Resultado)
			if err != nil {
				result.Warnings = append(result.Warnings,
					fmt.Sprintf("Não consegui ler o valor de '%s'", exam.NomeExame))
				continue
			}
		}

		// Conversão de unidade se necessário (best-effort)
		// Nota: lab_test_definitions.unit_conversion é texto livre, não regra parsável.
		// Por hora, não convertemos automaticamente — confiamos no usuário usar o laudo padrão.
		// Phase 2: parser de unit_conversion para casos comuns (ng/mL ↔ pg/mL, mg/L ↔ µg/L, etc).

		seen[cand.Item.ID] = true
		result.Matched = append(result.Matched, LightLabMatch{
			ScoreItemID:  cand.Item.ID.String(),
			ItemName:     cand.Item.Name,
			NumericValue: numeric,
			OriginalRaw:  exam.Resultado,
			Unit:         strDeref(exam.Unidade),
		})
	}

	return result, nil
}

// extractFirstNumber tenta achar o primeiro número (inteiro ou decimal) no texto.
// Útil quando o resultado vem como "9.2 mg/dL" ou "Negativo (1.2)".
func extractFirstNumber(s string) (float64, error) {
	s = strings.ReplaceAll(s, ",", ".")
	var current strings.Builder
	for i := 0; i < len(s); i++ {
		c := s[i]
		isDigit := c >= '0' && c <= '9'
		isSep := c == '.' || c == '-' || c == '+'
		if isDigit || (isSep && current.Len() > 0) || (c == '-' && current.Len() == 0) {
			current.WriteByte(c)
		} else if current.Len() > 0 {
			if v, err := strconv.ParseFloat(current.String(), 64); err == nil {
				return v, nil
			}
			current.Reset()
		}
	}
	if current.Len() > 0 {
		if v, err := strconv.ParseFloat(current.String(), 64); err == nil {
			return v, nil
		}
	}
	return 0, fmt.Errorf("nenhum número em %q", s)
}

func strDeref(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

// SaveTempPDF — helper para o handler salvar o upload em /tmp.
// O caller deve chamar os.Remove() após processar (já fazemos defensivamente em ExtractFromPDF).
func SaveTempPDF(data []byte) (string, error) {
	tmpDir := "/tmp/light-lab-uploads"
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create temp dir: %w", err)
	}
	name := fmt.Sprintf("upload-%s.pdf", uuid.NewString())
	path := filepath.Join(tmpDir, name)
	if err := os.WriteFile(path, data, 0600); err != nil {
		return "", fmt.Errorf("failed to write temp file: %w", err)
	}
	return path, nil
}
