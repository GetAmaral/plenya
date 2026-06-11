package services

import (
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// LabTestMatcher resolve um NOME de exame (texto livre/OCR) para a LabTestDefinition do catálogo.
// Núcleo reutilizável de: (a) anexar TUSS no render do pedido, (b) dedup ao importar pedido externo.
// Reaproveita normalizeTestName/containsSubstring (processing_job_service.go).
type LabTestMatcher struct{ db *gorm.DB }

func NewLabTestMatcher(db *gorm.DB) *LabTestMatcher { return &LabTestMatcher{db: db} }

// MatchIndex é o índice normalizado (nome→definição) construído uma vez por uso.
type MatchIndex struct {
	byName map[string]*models.LabTestDefinition
}

// BuildIndex carrega as definições requisitáveis ativas e indexa por Name/ShortName/AltNames.
func (m *LabTestMatcher) BuildIndex() (*MatchIndex, error) {
	var defs []models.LabTestDefinition
	if err := m.db.Where("is_active = ? AND is_requestable = ?", true, true).Find(&defs).Error; err != nil {
		return nil, err
	}
	idx := &MatchIndex{byName: make(map[string]*models.LabTestDefinition, len(defs)*2)}
	for i := range defs {
		d := &defs[i]
		add := func(s string) {
			n := normalizeTestName(s)
			if n == "" {
				return
			}
			if _, exists := idx.byName[n]; !exists {
				idx.byName[n] = d
			}
		}
		add(d.Name)
		if d.ShortName != nil {
			add(*d.ShortName)
		}
		for _, a := range d.AltNames {
			add(a)
		}
	}
	return idx, nil
}

// Resolve devolve a definição que casa com a linha (texto livre), ou nil.
// Prioriza match EXATO normalizado; cai para parcial (≥5 chars) escolhendo o nome indexado mais
// longo que contém / é contido — reduz super-match espúrio (relevante p/ o dedup).
func (idx *MatchIndex) Resolve(line string) *models.LabTestDefinition {
	n := normalizeTestName(line)
	if n == "" {
		return nil
	}
	if d, ok := idx.byName[n]; ok {
		return d
	}
	if len(n) < 5 {
		return nil
	}
	var best *models.LabTestDefinition
	bestLen := 0
	for name, d := range idx.byName {
		if len(name) < 5 {
			continue
		}
		if containsSubstring(n, name) || containsSubstring(name, n) {
			if len(name) > bestLen {
				best, bestLen = d, len(name)
			}
		}
	}
	return best
}
