package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/cmed"
	"github.com/plenya/api/internal/models"
)

var (
	ErrMedicationDefinitionNotFound = errors.New("medication definition not found")
)

type MedicationDefinitionService struct {
	db *gorm.DB
}

func NewMedicationDefinitionService(db *gorm.DB) *MedicationDefinitionService {
	return &MedicationDefinitionService{db: db}
}

// List lista todas as definições de medicamentos com filtros
func (s *MedicationDefinitionService) List(category *models.MedicationCategory, limit, offset int) ([]models.MedicationDefinition, int64, error) {
	var medications []models.MedicationDefinition
	var total int64

	query := s.db.Model(&models.MedicationDefinition{})

	// Filter by category if provided
	if category != nil && *category != "" {
		query = query.Where("category = ?", *category)
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated results
	if err := query.Order("common_name ASC").
		Limit(limit).
		Offset(offset).
		Find(&medications).Error; err != nil {
		return nil, 0, err
	}

	return medications, total, nil
}

// Search busca medicamentos por nome ou princípio ativo
func (s *MedicationDefinitionService) Search(query string, limit int) ([]models.MedicationDefinition, error) {
	var medications []models.MedicationDefinition

	if limit <= 0 || limit > 50 {
		limit = 20
	}

	// Busca em DOIS NÍVEIS. O catálogo da CMED tem ~26 mil apresentações, mas só ~10,8 mil
	// combinações (produto, concentração, forma) — que é o que o médico de fato escolhe.
	// Buscar linha a linha devolveria "dipirona" 60 vezes, variando só laboratório e tamanho
	// de caixa. Aqui devolvemos um representante por combinação; a apresentação exata sai
	// depois, por ListPresentations.
	//
	// O DISTINCT ON obriga o ORDER BY interno a começar pela chave do grupo, então o ranking
	// por relevância fica na consulta externa.
	sql := `
WITH termo AS (SELECT lower(public.immutable_unaccent(?)) AS t)
SELECT g.* FROM (
  SELECT DISTINCT ON (
           lower(public.immutable_unaccent(m.common_name)),
           coalesce(m.concentration, ''),
           coalesce(m.pharmaceutical_form, '')
         ) m.*
    FROM medication_definitions m, termo
   WHERE m.deleted_at IS NULL
     AND m.is_active
     AND m.is_prescribable
     AND ( lower(public.immutable_unaccent(m.common_name))       LIKE '%' || termo.t || '%'
        OR lower(public.immutable_unaccent(m.active_ingredient)) LIKE '%' || termo.t || '%' )
   ORDER BY lower(public.immutable_unaccent(m.common_name)),
            coalesce(m.concentration, ''),
            coalesce(m.pharmaceutical_form, ''),
            (m.product_type = 'Genérico') DESC,
            m.pmc_price ASC NULLS LAST
   LIMIT 200
) g, termo
ORDER BY
  CASE WHEN lower(public.immutable_unaccent(g.common_name))       LIKE termo.t || '%' THEN 0
       WHEN lower(public.immutable_unaccent(g.active_ingredient)) LIKE termo.t || '%' THEN 1
       ELSE 2 END,
  g.common_name, g.concentration
LIMIT ?`

	if err := s.db.Raw(sql, query, limit).Scan(&medications).Error; err != nil {
		return nil, err
	}
	return medications, nil
}

// ListPresentations devolve as apresentações concretas de uma combinação escolhida na busca
// (laboratórios e tamanhos de embalagem). É o segundo nível: só é consultado quando o médico
// quer imprimir a caixa exata, e por isso não polui o autocomplete.
func (s *MedicationDefinitionService) ListPresentations(product, concentration, form string, limit int) ([]models.MedicationDefinition, error) {
	var medications []models.MedicationDefinition

	if limit <= 0 || limit > 100 {
		limit = 25
	}

	sql := `
SELECT * FROM medication_definitions
 WHERE deleted_at IS NULL AND is_active AND is_prescribable
   AND lower(public.immutable_unaccent(common_name)) = lower(public.immutable_unaccent(?))
   AND coalesce(concentration, '') = ?
   AND coalesce(pharmaceutical_form, '') = ?
 ORDER BY (product_type = 'Genérico') DESC, pmc_price ASC NULLS LAST, laboratory
 LIMIT ?`

	if err := s.db.Raw(sql, product, concentration, form, limit).Scan(&medications).Error; err != nil {
		return nil, err
	}
	return medications, nil
}

// ReviewQueueItem é uma SUBSTÂNCIA pendente de conferência, não uma apresentação. O import
// marca ~5,9 mil apresentações como needs_review, mas elas se resumem a ~1.078 substâncias —
// conferir por substância é a única forma de a fila ter fim.
type ReviewQueueItem struct {
	ActiveIngredient string                    `json:"activeIngredient"`
	Category         models.MedicationCategory `json:"category"`
	CategorySource   string                    `json:"categorySource"`
	Stripe           *string                   `json:"stripe,omitempty"`
	TherapeuticClass *string                   `json:"therapeuticClass,omitempty"`
	Presentations    int                       `json:"presentations"`
	SampleProducts   string                    `json:"sampleProducts"`
	UsedByPatients   bool                      `json:"usedByPatients"`
}

// ReviewQueue devolve as substâncias que o import não conseguiu classificar com segurança,
// na ordem em que vale a pena conferir:
//
//  1. o que os pacientes JÁ usam ou já receberam (é o que vai ser prescrito de novo);
//  2. o risco de SUBESTIMAR controle — 'simple' vindo de palpite é o caso perigoso, porque a
//     CMED não publicou tarja nenhuma e o sistema assumiu receita simples;
//  3. presença no mercado (mais apresentações = mais chance de aparecer numa receita).
//
// Tarja preta ('a_b') fica no fim: já sai do receituário, então errar ali não gera receita
// indevida.
func (s *MedicationDefinitionService) ReviewQueue(limit, offset int) ([]ReviewQueueItem, int64, error) {
	if limit <= 0 || limit > 100 {
		limit = 25
	}

	var total int64
	if err := s.db.Raw(`
		SELECT count(DISTINCT active_ingredient)
		  FROM medication_definitions
		 WHERE needs_review AND deleted_at IS NULL AND is_active`).Scan(&total).Error; err != nil {
		return nil, 0, err
	}

	var items []ReviewQueueItem
	err := s.db.Raw(`
WITH conhecidas AS (
    SELECT DISTINCT lower(public.immutable_unaccent(active_ingredient)) AS ing
      FROM medications_in_use WHERE deleted_at IS NULL AND active_ingredient IS NOT NULL
     UNION
    SELECT DISTINCT lower(public.immutable_unaccent(active_ingredient))
      FROM prescription_medications WHERE deleted_at IS NULL
)
SELECT m.active_ingredient,
       min(m.category::text)        AS category,
       min(m.category_source)       AS category_source,
       min(m.stripe)                AS stripe,
       min(m.therapeutic_class)     AS therapeutic_class,
       count(*)                     AS presentations,
       string_agg(DISTINCT m.common_name, ', ' ORDER BY m.common_name) AS sample_products,
       bool_or(c.ing IS NOT NULL)   AS used_by_patients
  FROM medication_definitions m
  LEFT JOIN conhecidas c ON c.ing = lower(public.immutable_unaccent(m.active_ingredient))
 WHERE m.needs_review AND m.deleted_at IS NULL AND m.is_active
 GROUP BY m.active_ingredient
 ORDER BY bool_or(c.ing IS NOT NULL) DESC,
          (min(m.category::text) = 'simple' AND min(m.category_source) = 'cmed_fallback') DESC,
          (min(m.category::text) = 'a_b') ASC,
          count(*) DESC,
          m.active_ingredient
 LIMIT ? OFFSET ?`, limit, offset).Scan(&items).Error
	if err != nil {
		return nil, 0, err
	}
	return items, total, nil
}

// CurateSubstance grava a decisão do médico para TODAS as apresentações de uma substância.
// Confirmar a categoria deduzida e corrigi-la são a mesma operação: as duas tiram a linha da
// fila e carimbam curated_at, que é o que impede o reimport mensal de desfazer a decisão.
func (s *MedicationDefinitionService) CurateSubstance(
	activeIngredient string,
	category models.MedicationCategory,
	controlList *string,
	isPrescribable *bool,
	curatedBy uuid.UUID,
) (int64, error) {
	if activeIngredient == "" {
		return 0, errors.New("substância é obrigatória")
	}

	rules := cmed.RulesFor(category)
	updates := map[string]any{
		"category":                   category,
		"category_source":            models.MedCategorySourceManual,
		"needs_review":               false,
		"curated_at":                 time.Now().UTC(),
		"curated_by":                 curatedBy,
		"control_list":               controlList,
		"validity_days":              rules.ValidityDays,
		"max_per_prescription":       rules.MaxPerPrescription,
		"max_treatment_days":         rules.MaxTreatmentDays,
		"requires_digital_signature": rules.RequiresDigitalSignature,
		"requires_sncr":              rules.RequiresSNCR,
		"updated_at":                 time.Now().UTC(),
	}
	if isPrescribable != nil {
		updates["is_prescribable"] = *isPrescribable
	}

	res := s.db.Model(&models.MedicationDefinition{}).
		Where("lower(public.immutable_unaccent(active_ingredient)) = lower(public.immutable_unaccent(?))", activeIngredient).
		Where("deleted_at IS NULL").
		Updates(updates)

	return res.RowsAffected, res.Error
}

// GetByID busca uma definição por ID
func (s *MedicationDefinitionService) GetByID(id uuid.UUID) (*models.MedicationDefinition, error) {
	var medication models.MedicationDefinition
	if err := s.db.First(&medication, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMedicationDefinitionNotFound
		}
		return nil, err
	}

	return &medication, nil
}

// Create cria uma nova definição de medicamento
func (s *MedicationDefinitionService) Create(medication *models.MedicationDefinition) error {
	return s.db.Create(medication).Error
}

// Update atualiza uma definição de medicamento
func (s *MedicationDefinitionService) Update(id uuid.UUID, medication *models.MedicationDefinition, curatedBy *uuid.UUID) error {
	// Toda edição manual carimba curated_at: é essa marca que faz o reimport mensal da CMED
	// parar de sobrescrever os campos clínicos desta linha. Sem ela, a correção do médico
	// duraria até a próxima atualização da lista.
	now := time.Now().UTC()
	medication.CuratedAt = &now
	medication.CuratedBy = curatedBy
	medication.CategorySource = models.MedCategorySourceManual
	medication.NeedsReview = false

	result := s.db.Model(&models.MedicationDefinition{}).
		Where("id = ?", id).
		Updates(medication)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrMedicationDefinitionNotFound
	}

	return nil
}

// Delete deleta uma definição de medicamento (soft delete)
func (s *MedicationDefinitionService) Delete(id uuid.UUID) error {
	result := s.db.Delete(&models.MedicationDefinition{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrMedicationDefinitionNotFound
	}

	return nil
}
