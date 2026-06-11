package services

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// ImportedExam — um exame extraído de um pedido externo (foto/PDF), com o match no nosso catálogo.
type ImportedExam struct {
	Raw      string `json:"raw"`                // nome como veio do documento
	Matched  bool   `json:"matched"`            // casou com o catálogo?
	Code     string `json:"code,omitempty"`     // código interno da definição
	Name     string `json:"name,omitempty"`     // nome canônico
	TussCode string `json:"tussCode,omitempty"` // TUSS, se houver
}

// LabRequestImportService extrai exames de um pedido externo (OCR → IA → match) para dedup.
type LabRequestImportService struct {
	db      *gorm.DB
	ocr     *OCRService
	ai      *AIService
	cleaner *PDFTextCleaner
}

func NewLabRequestImportService(db *gorm.DB, ocr *OCRService, ai *AIService, cleaner *PDFTextCleaner) *LabRequestImportService {
	return &LabRequestImportService{db: db, ocr: ocr, ai: ai, cleaner: cleaner}
}

// ExtractExams: arquivo (PDF ou imagem) → OCR → limpeza → IA extrai nomes → casa no catálogo.
func (s *LabRequestImportService) ExtractExams(filePath string, isImage bool) ([]ImportedExam, error) {
	var text string
	var err error
	if isImage {
		text, err = s.ocr.ExtractTextFromImage(filePath)
	} else {
		text, err = s.ocr.ExtractText(filePath)
	}
	if err != nil {
		return nil, fmt.Errorf("OCR: %w", err)
	}
	cleaned := s.cleaner.CleanText(text)

	names, err := s.ai.ExtractRequestedExams(cleaned)
	if err != nil {
		return nil, fmt.Errorf("extração IA: %w", err)
	}

	idx, err := NewLabTestMatcher(s.db).BuildIndex()
	if err != nil {
		return nil, fmt.Errorf("índice: %w", err)
	}

	out := make([]ImportedExam, 0, len(names))
	seen := map[string]bool{}
	for _, n := range names {
		n = strings.TrimSpace(n)
		if n == "" {
			continue
		}
		item := ImportedExam{Raw: n}
		if d := idx.Resolve(n); d != nil {
			if seen[d.Code] {
				continue // dedup dentro do próprio import
			}
			seen[d.Code] = true
			item.Matched = true
			item.Code = d.Code
			item.Name = d.Name
			if d.TussCode != nil {
				item.TussCode = *d.TussCode
			}
		}
		out = append(out, item)
	}
	return out, nil
}
