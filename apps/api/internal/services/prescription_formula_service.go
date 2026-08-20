package services

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

// Regras da receita magistral. Ficam aqui, separadas do CRUD da prescrição, porque são as que
// mudam quando a frente de manipulados evoluir (catálogo de componentes, cálculo de cápsula,
// compatibilidade) — e porque assim dá para testá-las sem banco.

const (
	maxFormulasPerPrescription = 10
	// 20 componentes por fórmula é limite de LAYOUT, não clínico: o paginador do PDF só reflui
	// entre blocos, e uma fórmula é um bloco atômico. Um bloco maior que a página transborda em
	// silêncio, então o limite existe para a receita nunca sair cortada.
	maxComponentsPerFormula = 20
)

// distinctControlledSubstancesInFormulas conta substâncias C1 distintas na receita inteira
// (mesma regra da 344/98 aplicada ao manipulado: o limite é por substância, não por fórmula).
func distinctControlledSubstancesInFormulas(formulas []dto.FormulaRequest) int {
	seen := map[string]struct{}{}
	for _, f := range formulas {
		for _, c := range f.Components {
			if c.Category != models.MedCategoryC1 {
				continue
			}
			key := strings.ToLower(strings.TrimSpace(c.Substance))
			if key != "" {
				seen[key] = struct{}{}
			}
		}
	}
	return len(seen)
}

// formulaHasControlled indica se a fórmula tem componente de controle especial.
func formulaHasControlled(f dto.FormulaRequest) bool {
	for _, c := range f.Components {
		if models.IsControlled(c.Category) {
			return true
		}
	}
	return false
}

// validateFormulas aplica as regras da receita de manipulado.
//
// Deliberadamente NÃO recusa misturar fórmula controlada com fórmula comum na mesma receita: a
// 344/98 quer receituário exclusivo para o controlado, mas quem decide separar é o médico na
// hora de imprimir. Recusar aqui travaria a receita inteira por causa de uma fórmula.
func validateFormulas(formulas []dto.FormulaRequest) error {
	if len(formulas) == 0 {
		return errors.New("prescrição de manipulado precisa de pelo menos uma fórmula")
	}
	if len(formulas) > maxFormulasPerPrescription {
		return fmt.Errorf("prescrição não pode ter mais de %d fórmulas", maxFormulasPerPrescription)
	}

	for i, f := range formulas {
		if len(f.Components) == 0 {
			return fmt.Errorf("fórmula %d precisa de pelo menos um componente", i+1)
		}
		if len(f.Components) > maxComponentsPerFormula {
			return fmt.Errorf("fórmula %d não pode ter mais de %d componentes", i+1, maxComponentsPerFormula)
		}
		if strings.TrimSpace(f.PharmaceuticalForm) == "" {
			return fmt.Errorf("fórmula %d precisa da forma farmacêutica", i+1)
		}
		if f.QuantityToDispense <= 0 || strings.TrimSpace(f.QuantityUnit) == "" {
			return fmt.Errorf("fórmula %d precisa da quantidade a aviar", i+1)
		}
		// Extenso é exigência do receituário de controle especial — e só dele.
		if formulaHasControlled(f) && strings.TrimSpace(f.QuantityInWords) == "" {
			return fmt.Errorf("fórmula %d tem componente controlado e precisa da quantidade por extenso", i+1)
		}
	}

	if distinctControlledSubstancesInFormulas(formulas) > 3 {
		return errors.New("receita de controle especial aceita no máximo 3 substâncias diferentes")
	}
	return nil
}

// buildFormulas converte o request em models, resolvendo os UUIDs ANTES de gravar (um id
// inválido no meio da lista não pode deixar meia receita no banco).
func buildFormulas(prescriptionID uuid.UUID, formulas []dto.FormulaRequest) ([]models.PrescriptionFormula, error) {
	parse := func(s *string, campo string) (*uuid.UUID, error) {
		if s == nil || *s == "" {
			return nil, nil
		}
		id, err := uuid.Parse(*s)
		if err != nil {
			return nil, fmt.Errorf("%s inválido", campo)
		}
		return &id, nil
	}

	out := make([]models.PrescriptionFormula, 0, len(formulas))
	for i, f := range formulas {
		templateID, err := parse(f.TemplateID, "templateId")
		if err != nil {
			return nil, err
		}

		components := make([]models.PrescriptionFormulaComponent, 0, len(f.Components))
		for j, c := range f.Components {
			magistralID, err := parse(c.MagistralComponentID, "magistralComponentId")
			if err != nil {
				return nil, err
			}
			medDefID, err := parse(c.MedicationDefinitionID, "medicationDefinitionId")
			if err != nil {
				return nil, err
			}
			components = append(components, models.PrescriptionFormulaComponent{
				MagistralComponentID:   magistralID,
				MedicationDefinitionID: medDefID,
				DisplayOrder:           j,
				Substance:              strings.TrimSpace(c.Substance),
				Quantity:               c.Quantity,
				Unit:                   strings.TrimSpace(c.Unit),
				Category:               c.Category,
				Note:                   strings.TrimSpace(c.Note),
				SuggestedQuantity:      c.SuggestedQuantity,
				AsElemental:            c.AsElemental,
			})
		}

		usage := f.UsageType
		if usage == "" {
			usage = models.FormulaUsageInternal
		}

		out = append(out, models.PrescriptionFormula{
			PrescriptionID:     prescriptionID,
			TemplateID:         templateID,
			DisplayOrder:       i,
			Name:               strings.TrimSpace(f.Name),
			PharmaceuticalForm: strings.TrimSpace(f.PharmaceuticalForm),
			UsageType:          usage,
			Route:              strings.TrimSpace(f.Route),
			Vehicle:            strings.TrimSpace(f.Vehicle),
			QuantityToDispense: f.QuantityToDispense,
			QuantityUnit:       strings.TrimSpace(f.QuantityUnit),
			QuantityInWords:    strings.TrimSpace(f.QuantityInWords),
			Posology:           strings.TrimSpace(f.Posology),
			Duration:           f.Duration,
			Instructions:       f.Instructions,
			// Denormalizado no momento da escrita: é o que decide modo de assinatura e rótulo.
			HighestCategory: models.HighestCategoryOf(components),
			Components:      components,
		})
	}
	return out, nil
}

// calculateValidUntilFormulas — mesma regra do industrializado (validade mais restritiva da
// receita), lendo a categoria mais restritiva de cada fórmula.
func calculateValidUntilFormulas(prescriptionDate time.Time, formulas []models.PrescriptionFormula) time.Time {
	if len(formulas) == 0 {
		return prescriptionDate.AddDate(0, 0, defaultValidityDays)
	}
	minDays := 0
	for _, f := range formulas {
		days := defaultValidityDays
		if d, ok := validityDaysByCategory[f.HighestCategory]; ok {
			days = d
		}
		if minDays == 0 || days < minDays {
			minDays = days
		}
	}
	return prescriptionDate.AddDate(0, 0, minDays)
}

// formulasToDTO converte as fórmulas do model para a resposta da API.
func formulasToDTO(formulas []models.PrescriptionFormula) []dto.FormulaResponse {
	out := make([]dto.FormulaResponse, 0, len(formulas))
	for _, f := range formulas {
		components := make([]dto.FormulaComponentResponse, 0, len(f.Components))
		for _, c := range f.Components {
			comp := dto.FormulaComponentResponse{
				ID:                c.ID.String(),
				Substance:         c.Substance,
				Quantity:          c.Quantity,
				Unit:              c.Unit,
				Category:          c.Category,
				Note:              c.Note,
				SuggestedQuantity: c.SuggestedQuantity,
				AsElemental:       c.AsElemental,
			}
			if c.MagistralComponentID != nil {
				s := c.MagistralComponentID.String()
				comp.MagistralComponentID = &s
			}
			if c.MedicationDefinitionID != nil {
				s := c.MedicationDefinitionID.String()
				comp.MedicationDefinitionID = &s
			}
			components = append(components, comp)
		}

		resp := dto.FormulaResponse{
			ID:                 f.ID.String(),
			Name:               f.Name,
			PharmaceuticalForm: f.PharmaceuticalForm,
			UsageType:          f.UsageType,
			Route:              f.Route,
			Vehicle:            f.Vehicle,
			QuantityToDispense: f.QuantityToDispense,
			QuantityUnit:       f.QuantityUnit,
			QuantityInWords:    f.QuantityInWords,
			Posology:           f.Posology,
			Duration:           f.Duration,
			Instructions:       f.Instructions,
			HighestCategory:    f.HighestCategory,
			Components:         components,
		}
		if f.TemplateID != nil {
			s := f.TemplateID.String()
			resp.TemplateID = &s
		}
		out = append(out, resp)
	}
	return out
}
