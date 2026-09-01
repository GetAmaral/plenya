package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
)

// ScoreSnapshotService handles business logic for patient score snapshots
type ScoreSnapshotService struct {
	snapshotRepo  *repository.ScoreSnapshotRepository
	scoreRepo     *repository.ScoreRepository
	labResultRepo *repository.LabResultRepository
	anamnesisRepo *repository.AnamnesisRepository
	db            *gorm.DB
}

// NewScoreSnapshotService creates a new score snapshot service instance
func NewScoreSnapshotService(
	snapshotRepo *repository.ScoreSnapshotRepository,
	scoreRepo *repository.ScoreRepository,
	labResultRepo *repository.LabResultRepository,
	anamnesisRepo *repository.AnamnesisRepository,
	db *gorm.DB,
) *ScoreSnapshotService {
	return &ScoreSnapshotService{
		snapshotRepo:  snapshotRepo,
		scoreRepo:     scoreRepo,
		labResultRepo: labResultRepo,
		anamnesisRepo: anamnesisRepo,
		db:            db,
	}
}

// ============================================================
// DTOs (Data Transfer Objects)
// ============================================================

// CalculateSnapshotDTO represents the request to calculate a new snapshot
type CalculateSnapshotDTO struct {
	PatientID uuid.UUID `json:"patientId" validate:"required"`
	Notes     *string   `json:"notes,omitempty"`
}

// ============================================================
// Public Methods
// ============================================================

// CalculateSnapshot calculates a new score snapshot for a patient
// If a snapshot already exists for today, it will be updated instead of creating a new one
func (s *ScoreSnapshotService) CalculateSnapshot(dto CalculateSnapshotDTO, calculatedByUserID uuid.UUID) (*models.PatientScoreSnapshot, error) {
	// 1. Load patient and validate existence
	var patient models.Patient
	if err := s.db.First(&patient, "id = ?", dto.PatientID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("patient not found")
		}
		return nil, err
	}

	// 2. Check if snapshot already exists for today
	existingSnapshot, err := s.snapshotRepo.GetTodaySnapshotByPatientID(dto.PatientID)
	if err != nil {
		return nil, fmt.Errorf("failed to check for existing snapshot: %w", err)
	}

	// 3. Load complete score structure (all groups with subgroups, items, and levels)
	scoreGroups, err := s.scoreRepo.GetAllScoreGroupTrees()
	if err != nil {
		return nil, fmt.Errorf("failed to load score structure: %w", err)
	}

	// 4. Load patient's historical data (ENTIRE HISTORY - most recent values)
	labResultsByCode, err := s.labResultRepo.GetHistoricalResultsByLabTestCode(dto.PatientID)
	if err != nil {
		return nil, fmt.Errorf("failed to load lab results: %w", err)
	}

	anamnesisItemsByScoreItemID, err := s.anamnesisRepo.GetHistoricalItemsByScoreItemID(dto.PatientID)
	if err != nil {
		return nil, fmt.Errorf("failed to load anamnesis items: %w", err)
	}

	// Nomes e sinônimos de unidade do catálogo de exames, lidos uma vez para todos os itens.
	catalogo := carregaCatalogoDeExames(s.db)

	// 5. Create or update snapshot in transaction
	var snapshot *models.PatientScoreSnapshot
	err = s.db.Transaction(func(tx *gorm.DB) error {
		// If snapshot exists for today, reuse it and delete old results
		if existingSnapshot != nil {
			snapshot = existingSnapshot

			// Hard delete old group results (Unscoped to bypass soft delete)
			if err := tx.Unscoped().Where("snapshot_id = ?", snapshot.ID).Delete(&models.PatientScoreGroupResult{}).Error; err != nil {
				return fmt.Errorf("failed to delete old group results: %w", err)
			}

			// Hard delete old item results (Unscoped to bypass soft delete)
			if err := tx.Unscoped().Where("snapshot_id = ?", snapshot.ID).Delete(&models.PatientScoreItemResult{}).Error; err != nil {
				return fmt.Errorf("failed to delete old item results: %w", err)
			}

			// Reset snapshot values
			snapshot.CalculatedByUserID = calculatedByUserID
			snapshot.CalculatedAt = time.Now()
			snapshot.TotalActualPoints = 0
			snapshot.TotalPossiblePoints = 0
			snapshot.TotalScorePercentage = 0
			snapshot.ItemsEvaluatedCount = 0
			snapshot.ItemsNotEvaluatedCount = 0
			snapshot.Notes = dto.Notes
			// Defensivo: um recálculo do staff nunca deve carregar origem de import.
			snapshot.Source = nil
			snapshot.SourceSessionID = nil
		} else {
			// Initialize new snapshot
			snapshot = &models.PatientScoreSnapshot{
				PatientID:              dto.PatientID,
				CalculatedByUserID:     calculatedByUserID,
				CalculatedAt:           time.Now(),
				TotalActualPoints:      0,
				TotalPossiblePoints:    0,
				TotalScorePercentage:   0,
				ItemsEvaluatedCount:    0,
				ItemsNotEvaluatedCount: 0,
				Notes:                  dto.Notes,
			}
		}

		// Maps to aggregate by group
		groupResultsMap := make(map[uuid.UUID]*models.PatientScoreGroupResult)
		var allItemResults []models.PatientScoreItemResult

		// Process each group → subgroup → item
		for _, group := range scoreGroups {
			// Initialize group result
			groupResult := &models.PatientScoreGroupResult{
				GroupID:             group.ID,
				ActualPoints:        0,
				PossiblePoints:      0,
				ScorePercentage:     0,
				ItemsEvaluatedCount: 0,
			}
			groupResultsMap[group.ID] = groupResult

			// `Subgroups.Items` já traz os filhos (o Preload não filtra por parent_item_id),
			// e o laço interno percorre `ChildItems` de novo — sem o `seen` todo item com pai
			// entrava DUAS vezes no snapshot e no denominador. Mesma guarda que
			// anonymous_score_service e score_version_service já usam.
			seen := make(map[uuid.UUID]bool)

			// Process subgroups
			for _, subgroup := range group.Subgroups {
				// Process items
				for _, item := range subgroup.Items {
					if seen[item.ID] {
						continue
					}
					seen[item.ID] = true

					// Evaluate score item
					itemResult := s.evaluateScoreItem(
						&patient,
						&item,
						group.ID,
						labResultsByCode,
						anamnesisItemsByScoreItemID,
						catalogo,
					)

					allItemResults = append(allItemResults, itemResult)

					// Update group aggregations
					if itemResult.Status == models.EvaluationStatusEvaluated {
						groupResult.ActualPoints += itemResult.ActualPoints
						groupResult.PossiblePoints += itemResult.MaxPoints
						groupResult.ItemsEvaluatedCount++

						// Update snapshot totals
						snapshot.TotalActualPoints += itemResult.ActualPoints
						snapshot.TotalPossiblePoints += itemResult.MaxPoints
						snapshot.ItemsEvaluatedCount++
					} else {
						snapshot.ItemsNotEvaluatedCount++
					}

					// Process child items (if any)
					for _, childItem := range item.ChildItems {
						if seen[childItem.ID] {
							continue
						}
						seen[childItem.ID] = true

						childItemResult := s.evaluateScoreItem(
							&patient,
							&childItem,
							group.ID,
							labResultsByCode,
							anamnesisItemsByScoreItemID,
							catalogo,
						)

						allItemResults = append(allItemResults, childItemResult)

						if childItemResult.Status == models.EvaluationStatusEvaluated {
							groupResult.ActualPoints += childItemResult.ActualPoints
							groupResult.PossiblePoints += childItemResult.MaxPoints
							groupResult.ItemsEvaluatedCount++

							snapshot.TotalActualPoints += childItemResult.ActualPoints
							snapshot.TotalPossiblePoints += childItemResult.MaxPoints
							snapshot.ItemsEvaluatedCount++
						} else {
							snapshot.ItemsNotEvaluatedCount++
						}
					}
				}
			}

			// Calculate group percentage
			if groupResult.PossiblePoints > 0 {
				groupResult.ScorePercentage = (groupResult.ActualPoints / groupResult.PossiblePoints) * 100
			}
		}

		// Calculate total percentage
		if snapshot.TotalPossiblePoints > 0 {
			snapshot.TotalScorePercentage = (snapshot.TotalActualPoints / snapshot.TotalPossiblePoints) * 100
		}

		// Save snapshot (create new or update existing)
		if existingSnapshot != nil {
			// Update existing snapshot
			if err := tx.Save(snapshot).Error; err != nil {
				return fmt.Errorf("failed to update snapshot: %w", err)
			}
		} else {
			// Create new snapshot
			if err := tx.Create(snapshot).Error; err != nil {
				return fmt.Errorf("failed to create snapshot: %w", err)
			}
		}

		// Save group results
		for _, groupResult := range groupResultsMap {
			groupResult.SnapshotID = snapshot.ID
			if err := tx.Create(groupResult).Error; err != nil {
				return fmt.Errorf("failed to create group result: %w", err)
			}
		}

		// Save item results
		for i := range allItemResults {
			allItemResults[i].SnapshotID = snapshot.ID
			if err := tx.Create(&allItemResults[i]).Error; err != nil {
				return fmt.Errorf("failed to create item result: %w", err)
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Reload snapshot with all relations
	return s.snapshotRepo.GetByIDWithRelations(snapshot.ID)
}

// GetSnapshotsByPatientID retrieves all snapshots for a patient
func (s *ScoreSnapshotService) GetSnapshotsByPatientID(patientID uuid.UUID) ([]models.PatientScoreSnapshot, error) {
	return s.snapshotRepo.GetAllByPatientID(patientID)
}

// GetLatestSnapshotByPatientID retrieves the most recent snapshot for a patient
func (s *ScoreSnapshotService) GetLatestSnapshotByPatientID(patientID uuid.UUID) (*models.PatientScoreSnapshot, error) {
	return s.snapshotRepo.GetLatestByPatientID(patientID)
}

// GetSnapshotByID retrieves a snapshot by ID
func (s *ScoreSnapshotService) GetSnapshotByID(id uuid.UUID) (*models.PatientScoreSnapshot, error) {
	return s.snapshotRepo.GetByIDWithRelations(id)
}

// DeleteSnapshot deletes a snapshot
func (s *ScoreSnapshotService) DeleteSnapshot(id uuid.UUID) error {
	return s.snapshotRepo.DeleteSnapshot(id)
}

// ============================================================
// Private Methods
// ============================================================

// motivoDoContexto explica, em português de prontuário, por que o item de contexto não se
// aplica. O código cru (`PLN1BFE6CA3`) não diz nada a quem lê o escore, então o nome do exame
// vem de `nomes` (catálogo) e, na falta dele, do resultado do próprio paciente.
//
// Só é chamada quando `RequirementMet` reprovou, o que exige `RequiresLabCode` preenchido e ao
// menos um de min/max — por isso não há ramo para os dois nulos.
func motivoDoContexto(item *models.ScoreItem, labResultsByCode map[string]models.LabResult, catalogo catalogoDeExames) string {
	codigo := *item.RequiresLabCode
	nome := codigo
	if n := catalogo.nomeDe(codigo); n != "" {
		nome = n
	}
	ref, temResultado := labResultsByCode[codigo]
	if temResultado && ref.TestName != "" {
		nome = ref.TestName
	}

	switch {
	case !temResultado:
		return fmt.Sprintf("Item não avaliado: só é interpretável junto com %s, que não foi medido", nome)
	case ref.ResultNumeric == nil:
		// O exame foi feito, mas veio qualitativo ou em texto: dizer "não foi medido" seria
		// mentira, e mandaria pedir de novo um exame que já está no prontuário.
		return fmt.Sprintf("Item não avaliado: %s veio sem valor numérico, e a interpretação depende do número", nome)
	}

	var faixa string
	switch {
	case item.RequiresMin != nil && item.RequiresMax != nil:
		faixa = fmt.Sprintf("entre %s e %s", formatNumberPT(*item.RequiresMin), formatNumberPT(*item.RequiresMax))
	case item.RequiresMin != nil:
		faixa = fmt.Sprintf("a partir de %s", formatNumberPT(*item.RequiresMin))
	default:
		faixa = fmt.Sprintf("até %s", formatNumberPT(*item.RequiresMax))
	}
	return fmt.Sprintf("Item não aplicável: só discrimina com %s %s (medido: %s)",
		nome, faixa, formatNumberPT(*ref.ResultNumeric))
}

// valoresPorCodigo reduz os resultados do paciente a "código → valor numérico", que é o contexto
// que `ScoreItem.RequirementMet` consulta.
func valoresPorCodigo(porCodigo map[string]models.LabResult) map[string]float64 {
	out := make(map[string]float64, len(porCodigo))
	for code, r := range porCodigo {
		if r.ResultNumeric != nil {
			out[code] = *r.ResultNumeric
		}
	}
	return out
}

// evaluateScoreItem evaluates a single ScoreItem for the patient
func (s *ScoreSnapshotService) evaluateScoreItem(
	patient *models.Patient,
	item *models.ScoreItem,
	groupID uuid.UUID,
	labResultsByCode map[string]models.LabResult,
	anamnesisItemsByScoreItemID map[uuid.UUID]models.AnamnesisItem,
	catalogo catalogoDeExames,
) models.PatientScoreItemResult {
	result := models.PatientScoreItemResult{
		ItemID:       item.ID,
		GroupID:      groupID,
		Status:       models.EvaluationStatusNoDataAvailable,
		MaxPoints:    0,
		ActualPoints: 0,
	}

	// Set MaxPoints from item
	if item.Points != nil {
		result.MaxPoints = *item.Points
	}

	// 1a. Contexto de outro exame: item cuja escala só vale dentro de uma faixa de outro
	// resultado (a razão %Free PSA só discrimina com PSA total entre 4 e 10 ng/mL). Fora do
	// contexto ele não é "ruim", é *inaplicável* — pontuar ali tira ponto do paciente por um
	// número que não quer dizer nada.
	if !item.RequirementMet(valoresPorCodigo(labResultsByCode)) {
		result.Status = models.EvaluationStatusNotApplicable
		reason := motivoDoContexto(item, labResultsByCode, catalogo)
		result.NotEvaluatedReason = &reason
		return result
	}

	// 1. Check if item applies to this patient
	if !item.AppliesToPatient(patient) {
		result.Status = models.EvaluationStatusNotApplicable
		reason := fmt.Sprintf("Item não aplicável ao paciente")
		if item.Gender != nil && *item.Gender != "not_applicable" {
			reason = fmt.Sprintf("Item não aplicável: sexo %s requerido (paciente: %s)", *item.Gender, patient.Gender)
		} else if item.AgeRangeMin != nil || item.AgeRangeMax != nil {
			reason = fmt.Sprintf("Item não aplicável: fora da faixa etária")
		}
		result.NotEvaluatedReason = &reason
		return result
	}

	// 2. Find value and determine source
	var dataSource *models.DataSource
	var labResultID *uuid.UUID
	var anamnesisItemID *uuid.UUID
	var matchedLevel *models.ScoreLevel
	var valueUsed *float64
	// Preenchido quando o exame do paciente foi descartado por falar de outra grandeza. Vira o
	// motivo final só se nada mais (anamnese) tiver servido.
	var motivoUnidade string

	// Try to find value from LabResult (if LabTestCode is set)
	if item.LabTestCode != nil && *item.LabTestCode != "" {
		if labResult, found := labResultsByCode[*item.LabTestCode]; found {
			// Escala numa grandeza diferente da do exame: os números não se comparam, e
			// classificar assim é inventar resultado. Três itens do sedimento urinário têm
			// escala em `células/campo` enquanto o laboratório reporta `/µL` — 0,5/µL cai na
			// faixa "≤10 células/campo" e sai como nível ÓTIMO sem que ninguém tenha olhado
			// nada. Descarta o exame sem abortar o item: se houver anamnese para o mesmo
			// item, ela ainda vale.
			unidadeExame := ""
			if labResult.Unit != nil {
				unidadeExame = *labResult.Unit
			}
			if !item.UnitMatches(unidadeExame, catalogo.sinonimosDe(item.LabTestCode)) {
				motivoUnidade = fmt.Sprintf(
					"Item não avaliado: a escala está em %s e o exame veio em %s (grandezas diferentes)",
					*item.Unit, unidadeExame)
			} else if labResult.ResultNumeric != nil {
				// Exame numérico: avalia o valor contra os limites dos níveis.
				valueUsed = labResult.ResultNumeric
				ds := models.DataSourceLabResult
				dataSource = &ds
				labResultID = &labResult.ID

				// Sort levels by level number ASC (0→6) - most critical first
				levels := item.Levels
				SortLevelsAsc(levels)

				for i := range levels {
					if levels[i].EvaluatesTrue(*valueUsed) {
						matchedLevel = &levels[i]
						break // Use FIRST match (most critical level that applies)
					}
				}

				if matchedLevel == nil {
					reason := fmt.Sprintf("Valor %.2f não corresponde a nenhum nível definido", *valueUsed)
					result.NotEvaluatedReason = &reason
					return result
				}
			} else if labResult.Level != nil {
				// Exame qualitativo (ex.: genótipo, categórico): o valor não é numérico.
				// Usa o nível pré-casado na entrada/ingestão (lab_results.level) para
				// selecionar o score level pelo número — direção alto=bom / 0=ruim.
				ds := models.DataSourceLabResult
				dataSource = &ds
				labResultID = &labResult.ID

				for i := range item.Levels {
					if item.Levels[i].Level == *labResult.Level {
						matchedLevel = &item.Levels[i]
						break
					}
				}

				if matchedLevel == nil {
					reason := fmt.Sprintf("Nível %d não existe na configuração do item", *labResult.Level)
					result.NotEvaluatedReason = &reason
					return result
				}
			}
			// Sem result_numeric e sem level: não seta dataSource → tenta anamnese / no_data.
		}
	}

	// If not found in lab results, try AnamnesisItem.
	// Dois modos (mesma semântica do escore-light/anônimo):
	//  1. SelectedLevel preenchido → encontra o level pelo número (caminho normal: o
	//     front grava o nível classificado em selected_level, não em numeric_value).
	//  2. NumericValue preenchido (sem nível) → avalia o valor contra os limites.
	if dataSource == nil {
		if anamnesisItem, found := anamnesisItemsByScoreItemID[item.ID]; found {
			if anamnesisItem.SelectedLevel != nil {
				levelNumber := *anamnesisItem.SelectedLevel
				for i := range item.Levels {
					if item.Levels[i].Level == levelNumber {
						matchedLevel = &item.Levels[i]
						break
					}
				}
				if matchedLevel == nil {
					reason := fmt.Sprintf("Nível %d não existe na configuração do item", levelNumber)
					result.NotEvaluatedReason = &reason
					return result
				}
			} else if anamnesisItem.NumericValue != nil {
				levels := item.Levels
				SortLevelsAsc(levels)
				for i := range levels {
					if levels[i].EvaluatesTrue(*anamnesisItem.NumericValue) {
						matchedLevel = &levels[i]
						break
					}
				}
				if matchedLevel == nil {
					reason := fmt.Sprintf("Valor %.2f não corresponde a nenhum nível definido", *anamnesisItem.NumericValue)
					result.NotEvaluatedReason = &reason
					return result
				}
			}

			if matchedLevel != nil {
				ds := models.DataSourceAnamnesisItem
				dataSource = &ds
				anamnesisItemID = &anamnesisItem.ID
				valueUsed = anamnesisItem.NumericValue
			}
		}
	}

	// If no value found in entire history, return no_data_available
	if dataSource == nil {
		if motivoUnidade != "" {
			result.Status = models.EvaluationStatusNotApplicable
			result.NotEvaluatedReason = &motivoUnidade
			return result
		}
		reason := "Sem dados disponíveis em todo o histórico do paciente"
		result.NotEvaluatedReason = &reason
		return result
	}

	// 3. Calculate proportional points
	proportionalMultiplier := GetLevelMultiplier(matchedLevel.Level)
	actualPoints := result.MaxPoints * proportionalMultiplier

	// 4. Populate result
	result.Status = models.EvaluationStatusEvaluated
	result.DataSource = dataSource
	result.LabResultID = labResultID
	result.AnamnesisItemID = anamnesisItemID
	result.ValueUsed = valueUsed
	result.LevelMatchedID = &matchedLevel.ID
	result.LevelNumber = &matchedLevel.Level
	result.ActualPoints = actualPoints

	return result
}

// SortLevelsAsc ordena ScoreLevels por Level ASC (mais crítico primeiro).
// Helper compartilhado entre score_snapshot_service e anonymous_score_service.
func SortLevelsAsc(levels []models.ScoreLevel) {
	for i := 0; i < len(levels); i++ {
		for j := i + 1; j < len(levels); j++ {
			if levels[j].Level < levels[i].Level {
				levels[i], levels[j] = levels[j], levels[i]
			}
		}
	}
}

// GetLevelMultiplier returns the proportional multiplier for a given level (0-6)
func GetLevelMultiplier(level int) float64 {
	switch level {
	case 6:
		return 1.0 // 100%
	case 5:
		return 1.0 // 100%
	case 4:
		return 0.8 // 80%
	case 3:
		return 0.6 // 60%
	case 2:
		return 0.4 // 40%
	case 1:
		return 0.2 // 20%
	case 0:
		return 0.0 // 0%
	default:
		return 0.0
	}
}
