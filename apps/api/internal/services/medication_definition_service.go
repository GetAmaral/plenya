package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

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
