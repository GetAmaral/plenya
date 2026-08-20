package services

import (
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

func med(category models.MedicationCategory, ingredient, name string) dto.MedicationRequest {
	return dto.MedicationRequest{
		Category:         category,
		ActiveIngredient: ingredient,
		MedicationName:   name,
	}
}

// A Portaria 344/98 limita o receituário de controle especial a 3 SUBSTÂNCIAS. Contar itens
// (o comportamento antigo) recusava receita legítima e aceitava receita irregular.
func TestDistinctControlledSubstances(t *testing.T) {
	cases := []struct {
		name string
		meds []dto.MedicationRequest
		want int
	}{
		{"sem controlado", []dto.MedicationRequest{med(models.MedCategorySimple, "losartana", "Losartana")}, 0},
		{
			"mesma substância em 2 apresentações conta 1",
			[]dto.MedicationRequest{
				med(models.MedCategoryC1, "clonazepam", "Rivotril 0,5mg"),
				med(models.MedCategoryC1, "Clonazepam", "Rivotril 2mg"),
			},
			1,
		},
		{
			"três substâncias distintas",
			[]dto.MedicationRequest{
				med(models.MedCategoryC1, "clonazepam", "A"),
				med(models.MedCategoryC1, "sertralina", "B"),
				med(models.MedCategoryC1, "quetiapina", "C"),
			},
			3,
		},
		{
			"sem princípio ativo cai no nome comercial",
			[]dto.MedicationRequest{
				med(models.MedCategoryC1, "", "Rivotril"),
				med(models.MedCategoryC1, "", "rivotril"),
			},
			1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := distinctControlledSubstances(tc.meds); got != tc.want {
				t.Errorf("esperava %d substância(s), veio %d", tc.want, got)
			}
		})
	}
}

func TestValidateMedications(t *testing.T) {
	if err := validateMedications(nil); err == nil {
		t.Error("receita sem medicamento deveria falhar")
	}

	four := []dto.MedicationRequest{
		med(models.MedCategoryC1, "clonazepam", "A"),
		med(models.MedCategoryC1, "sertralina", "B"),
		med(models.MedCategoryC1, "quetiapina", "C"),
		med(models.MedCategoryC1, "zolpidem", "D"),
	}
	if err := validateMedications(four); err == nil {
		t.Error("4 substâncias C1 deveriam ser recusadas")
	}
	if err := validateMedications(four[:3]); err != nil {
		t.Errorf("3 substâncias C1 são permitidas, veio erro: %v", err)
	}
}

// A validade é a MAIS restritiva da receita, e o catálogo (curável) tem precedência sobre a
// tabela por categoria compilada no binário.
func TestCalculateValidUntil(t *testing.T) {
	base := time.Date(2026, 8, 1, 12, 0, 0, 0, time.UTC)
	days := func(t time.Time) int { return int(t.Sub(base).Hours() / 24) }

	if got := days(calculateValidUntil(base, nil, nil)); got != 30 {
		t.Errorf("receita vazia deveria valer 30 dias, veio %d", got)
	}

	mixed := []dto.MedicationRequest{
		med(models.MedCategorySimple, "losartana", "A"),
		med(models.MedCategoryAntibiotic, "amoxicilina", "B"),
		med(models.MedCategoryGLP1, "semaglutida", "C"),
	}
	if got := days(calculateValidUntil(base, mixed, nil)); got != 10 {
		t.Errorf("antimicrobiano deveria puxar a validade para 10 dias, veio %d", got)
	}

	id := uuid.Must(uuid.NewV7())
	idStr := id.String()
	fromCatalog := []dto.MedicationRequest{{
		Category:               models.MedCategorySimple,
		ActiveIngredient:       "losartana",
		MedicationDefinitionID: &idStr,
	}}
	if got := days(calculateValidUntil(base, fromCatalog, map[uuid.UUID]int{id: 60})); got != 60 {
		t.Errorf("validade do catálogo deveria vencer a da categoria, veio %d", got)
	}
	if got := days(calculateValidUntil(base, fromCatalog, map[uuid.UUID]int{id: 0})); got != 30 {
		t.Errorf("validade zerada no catálogo deveria cair na categoria, veio %d", got)
	}
}
