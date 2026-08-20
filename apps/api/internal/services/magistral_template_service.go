package services

import (
	"errors"
	"math"
	"sort"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

var ErrTemplateNotFound = errors.New("formula template not found")

// ErrPatientNotSelected — a sugestão só responde sobre o paciente aberto na tela.
var ErrPatientNotSelected = errors.New("patient id does not match selected patient")

// MagistralTemplateService cuida das fórmulas-base e das sugestões de dose.
//
// A busca do prontuário (peso, exame) vive aqui; o cálculo vive em magistral_dose_rules.go, puro
// e testável. A separação é de propósito: o que decide dose não toca banco.
type MagistralTemplateService struct {
	db       *gorm.DB
	clinical *ClinicalDataService
}

func NewMagistralTemplateService(db *gorm.DB) *MagistralTemplateService {
	return &MagistralTemplateService{db: db, clinical: NewClinicalDataService(db)}
}

func (s *MagistralTemplateService) preload(q *gorm.DB) *gorm.DB {
	return q.
		Preload("Components", func(db *gorm.DB) *gorm.DB { return db.Order("display_order") }).
		Preload("Components.Rule").
		Preload("Components.Rule.Bands", func(db *gorm.DB) *gorm.DB { return db.Order("display_order") })
}

// List — fórmulas-base ativas, mais usadas primeiro. Busca por nome OU indicação: o médico
// procura tanto por "fórmula do sono" quanto por "insônia".
func (s *MagistralTemplateService) List(query string, limit int) ([]models.MagistralFormulaTemplate, error) {
	if limit <= 0 || limit > 100 {
		limit = 30
	}
	q := s.preload(s.db.Where("is_active = true AND deleted_at IS NULL"))

	if term := strings.TrimSpace(query); term != "" {
		like := "%" + strings.ToLower(term) + "%"
		q = q.Where(`lower(public.immutable_unaccent(name)) LIKE lower(public.immutable_unaccent(?))
		             OR lower(public.immutable_unaccent(coalesce(indication,''))) LIKE lower(public.immutable_unaccent(?))
		             OR lower(public.immutable_unaccent(coalesce(indication_bullets,''))) LIKE lower(public.immutable_unaccent(?))`,
			like, like, like)
	}

	var out []models.MagistralFormulaTemplate
	err := q.Order("usage_count DESC").Order("name").Limit(limit).Find(&out).Error
	return out, err
}

func (s *MagistralTemplateService) GetByID(id uuid.UUID) (*models.MagistralFormulaTemplate, error) {
	var t models.MagistralFormulaTemplate
	if err := s.preload(s.db).First(&t, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrTemplateNotFound
		}
		return nil, err
	}
	return &t, nil
}

// Save cria ou substitui uma fórmula-base inteira (componentes e regras junto), numa transação.
func (s *MagistralTemplateService) Save(id *uuid.UUID, req *dto.FormulaTemplateRequest, userID uuid.UUID) (*models.MagistralFormulaTemplate, error) {
	if len(req.Components) == 0 {
		return nil, errors.New("a fórmula-base precisa de pelo menos um componente")
	}
	if len(req.Components) > maxComponentsPerFormula {
		return nil, errors.New("máximo de 20 componentes por fórmula")
	}

	tpl := models.MagistralFormulaTemplate{CreatedBy: &userID}
	if id != nil {
		existing, err := s.GetByID(*id)
		if err != nil {
			return nil, err
		}
		tpl = *existing
	}

	tpl.Name = strings.TrimSpace(req.Name)
	tpl.Indication = req.Indication
	tpl.IndicationBullets = req.IndicationBullets
	tpl.PharmaceuticalForm = strings.TrimSpace(req.PharmaceuticalForm)
	tpl.UsageType = req.UsageType
	if tpl.UsageType == "" {
		tpl.UsageType = models.FormulaUsageInternal
	}
	tpl.Route = strings.TrimSpace(req.Route)
	tpl.Vehicle = strings.TrimSpace(req.Vehicle)
	if req.QuantityToDispense > 0 {
		tpl.QuantityToDispense = req.QuantityToDispense
	}
	if strings.TrimSpace(req.QuantityUnit) != "" {
		tpl.QuantityUnit = strings.TrimSpace(req.QuantityUnit)
	}
	tpl.Posology = strings.TrimSpace(req.Posology)
	tpl.Duration = req.Duration
	tpl.Instructions = req.Instructions
	tpl.Notes = req.Notes
	tpl.ReviewedBy = &userID
	now := time.Now()
	tpl.LastReview = &now

	err := s.db.Transaction(func(tx *gorm.DB) error {
		if id == nil {
			if err := tx.Create(&tpl).Error; err != nil {
				return err
			}
		} else {
			// Omit das associações: Save com filhos carregados os ressuscitaria depois do delete.
			if err := tx.Omit(clause.Associations).Save(&tpl).Error; err != nil {
				return err
			}
			var oldIDs []uuid.UUID
			if err := tx.Model(&models.MagistralFormulaTemplateComponent{}).Unscoped().
				Where("template_id = ?", tpl.ID).Pluck("id", &oldIDs).Error; err != nil {
				return err
			}
			if len(oldIDs) > 0 {
				if err := tx.Unscoped().Where("template_component_id IN ?", oldIDs).
					Delete(&models.MagistralFormulaTemplateRule{}).Error; err != nil {
					return err
				}
				if err := tx.Unscoped().Where("template_id = ?", tpl.ID).
					Delete(&models.MagistralFormulaTemplateComponent{}).Error; err != nil {
					return err
				}
			}
		}

		for i, c := range req.Components {
			comp := models.MagistralFormulaTemplateComponent{
				TemplateID:   tpl.ID,
				DisplayOrder: i,
				Substance:    strings.TrimSpace(c.Substance),
				Quantity:     c.Quantity,
				Unit:         strings.TrimSpace(c.Unit),
				Category:     c.Category,
				Note:         strings.TrimSpace(c.Note),
				AsElemental:  c.AsElemental,
			}
			if comp.Category == "" {
				comp.Category = models.MedCategorySimple
			}
			if c.MagistralComponentID != nil && *c.MagistralComponentID != "" {
				cid, err := uuid.Parse(*c.MagistralComponentID)
				if err != nil {
					return errors.New("magistralComponentId inválido")
				}
				comp.MagistralComponentID = &cid
			}
			if err := tx.Create(&comp).Error; err != nil {
				return err
			}

			if c.Rule == nil {
				continue
			}
			rule, err := buildRule(comp.ID, c.Rule)
			if err != nil {
				return err
			}
			if err := tx.Create(rule).Error; err != nil {
				return err
			}
		}

		return s.preload(tx).First(&tpl, "id = ?", tpl.ID).Error
	})
	if err != nil {
		return nil, err
	}
	return &tpl, nil
}

// buildRule valida a regra antes de gravar. Piso e teto são obrigatórios — é a trava que impede
// dado errado no prontuário de virar dose absurda.
func buildRule(componentID uuid.UUID, req *dto.DoseRuleRequest) (*models.MagistralFormulaTemplateRule, error) {
	if req.MinDose <= 0 || req.MaxDose <= 0 {
		return nil, errors.New("a regra de dose precisa de piso e teto maiores que zero")
	}
	if req.MaxDose < req.MinDose {
		return nil, errors.New("o teto da regra não pode ser menor que o piso")
	}

	rule := &models.MagistralFormulaTemplateRule{
		TemplateComponentID: componentID,
		Kind:                req.Kind,
		PerKg:               req.PerKg,
		LabCode:             req.LabCode,
		LabOperator:         req.LabOperator,
		LabThreshold:        req.LabThreshold,
		DoseIfTrue:          req.DoseIfTrue,
		DoseIfFalse:         req.DoseIfFalse,
		FixedDose:           req.FixedDose,
		LabUnit:             req.LabUnit,
		RoundTo:             req.RoundTo,
		MinDose:             req.MinDose,
		MaxDose:             req.MaxDose,
		MaxDataAgeDays:      req.MaxDataAgeDays,
		Note:                strings.TrimSpace(req.Note),
	}
	if rule.MaxDataAgeDays <= 0 {
		rule.MaxDataAgeDays = 365
	}

	switch req.Kind {
	case models.DoseRuleFixed:
		if req.FixedDose == nil {
			return nil, errors.New("regra fixa precisa da dose")
		}
	case models.DoseRulePerKg:
		if req.PerKg == nil {
			return nil, errors.New("regra por peso precisa da dose por kg")
		}
	case models.DoseRuleLabThreshold:
		if req.LabCode == nil || req.LabOperator == nil || req.LabThreshold == nil || req.DoseIfTrue == nil {
			return nil, errors.New("regra por exame precisa de exame, comparação, limiar e dose")
		}
	case models.DoseRuleLabBand:
		if req.LabCode == nil || len(req.Bands) == 0 {
			return nil, errors.New("regra por faixa precisa do exame e de pelo menos uma faixa")
		}
		bands, err := buildBands(req.Bands)
		if err != nil {
			return nil, err
		}
		rule.Bands = bands
	default:
		return nil, errors.New("tipo de regra inválido")
	}
	return rule, nil
}

// buildBands valida e ordena as faixas. Duas travas:
//
//   - faixa precisa de piso OU teto (faixa sem os dois pega qualquer valor: é regra fixa
//     disfarçada, e o médico leria "conforme o exame" onde o exame não muda nada);
//   - faixas não podem se sobrepor — com sobreposição, a dose passa a depender da ORDEM de
//     cadastro, que é a última coisa que alguém confere.
//
// Buraco entre faixas é permitido de propósito: "só doso abaixo de 30" é conduta legítima, e o
// motor responde com o motivo quando o valor cai no buraco.
func buildBands(reqs []dto.DoseBandRequest) ([]models.MagistralFormulaTemplateRuleBand, error) {
	out := make([]models.MagistralFormulaTemplateRuleBand, 0, len(reqs))
	for i, b := range reqs {
		if b.Dose <= 0 {
			return nil, errors.New("toda faixa precisa de dose maior que zero")
		}
		if b.LowerBound == nil && b.UpperBound == nil {
			return nil, errors.New("faixa sem piso e sem teto pega qualquer valor: use regra fixa")
		}
		if b.LowerBound != nil && b.UpperBound != nil && *b.LowerBound >= *b.UpperBound {
			return nil, errors.New("o piso da faixa precisa ser menor que o teto")
		}
		out = append(out, models.MagistralFormulaTemplateRuleBand{
			DisplayOrder: i,
			LowerBound:   b.LowerBound,
			UpperBound:   b.UpperBound,
			Dose:         b.Dose,
			Label:        strings.TrimSpace(b.Label),
		})
	}

	sort.Slice(out, func(i, j int) bool {
		li, lj := math.Inf(-1), math.Inf(-1)
		if out[i].LowerBound != nil {
			li = *out[i].LowerBound
		}
		if out[j].LowerBound != nil {
			lj = *out[j].LowerBound
		}
		return li < lj
	})
	for i := range out {
		out[i].DisplayOrder = i
		if i == 0 {
			continue
		}
		prevTop := math.Inf(1)
		if out[i-1].UpperBound != nil {
			prevTop = *out[i-1].UpperBound
		}
		thisFloor := math.Inf(-1)
		if out[i].LowerBound != nil {
			thisFloor = *out[i].LowerBound
		}
		if thisFloor < prevTop {
			return nil, errors.New("as faixas se sobrepõem: cada valor de exame precisa cair em uma faixa só")
		}
	}
	return out, nil
}

func (s *MagistralTemplateService) Delete(id uuid.UUID) error {
	return s.db.Delete(&models.MagistralFormulaTemplate{}, "id = ?", id).Error
}

// IncrementUsage — a fórmula-base sobe na lista conforme vira receita de verdade.
func (s *MagistralTemplateService) IncrementUsage(tx *gorm.DB, templateID uuid.UUID) error {
	return tx.Model(&models.MagistralFormulaTemplate{}).Where("id = ?", templateID).
		UpdateColumn("usage_count", gorm.Expr("usage_count + 1")).Error
}

// --- sugestão de dose ---

// Suggest aplica as regras da fórmula-base ao paciente e devolve SUGESTÕES.
//
// Nada aqui escreve na receita: a tela recebe os números e o médico decide. O payload de criação
// de prescrição não tem noção de regra, então um erro deste caminho não chega a documento
// assinado.
func (s *MagistralTemplateService) Suggest(templateID, patientID, userID uuid.UUID) (*dto.DoseSuggestionResponse, error) {
	// SEGURANÇA: a sugestão lê peso, exame e a última dose assinada do paciente — é prontuário.
	// Sem esta trava, qualquer usuário de equipe lia o prontuário de qualquer paciente passando o
	// UUID na query. Toda tela com dado de paciente escopa pelo paciente selecionado; esta
	// escapou por ser um endpoint de cálculo.
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}
	if user.SelectedPatientID == nil {
		return nil, ErrPatientNotSelected
	}
	if patientID != *user.SelectedPatientID {
		return nil, ErrPatientNotSelected
	}

	tpl, err := s.GetByID(templateID)
	if err != nil {
		return nil, err
	}

	weight := s.latestWeight(patientID)
	now := time.Now()

	resp := &dto.DoseSuggestionResponse{TemplateID: tpl.ID.String(), TemplateName: tpl.Name}
	if weight != nil {
		resp.WeightKg = &weight.Value
		resp.WeightSource = describeSource(weight)
		if weight.Date != nil {
			d := weight.Date.Format("2006-01-02")
			resp.WeightDate = &d
		}
	}

	// A posologia da própria fórmula-base diz quantas tomadas o dia tem: é o que converte a dose
	// diária da regra na dose de uma cápsula.
	tomadas := DosesPorDia(tpl.Posology)
	resp.DosesPerDay = tomadas

	for _, c := range tpl.Components {
		in := SuggestInput{Component: c, Rule: c.Rule, Weight: weight, Now: now, DosesPerDay: tomadas}
		if c.Rule != nil && c.Rule.LabCode != nil &&
			(c.Rule.Kind == models.DoseRuleLabThreshold || c.Rule.Kind == models.DoseRuleLabBand) {
			in.Lab = s.latestLab(patientID, *c.Rule.LabCode)
		}

		sug := SuggestDose(in)
		item := dto.DoseSuggestionItem{
			TemplateComponentID: c.ID.String(),
			Substance:           sug.Substance,
			Unit:                sug.Unit,
			BaseDose:            sug.BaseDose,
			Suggested:           sug.Suggested,
			Basis:               sug.Basis,
			Clamped:             sug.Clamped,
			Reason:              sug.Reason,
		}
		for _, b := range sug.Bands {
			item.Bands = append(item.Bands, dto.DoseBandView{
				LowerBound: b.LowerBound, UpperBound: b.UpperBound,
				Dose: b.Dose, Label: b.Label, Active: b.Active,
			})
		}
		// "Resposta do paciente" não é dado estruturado no EMR. O que existe e vale tanto quanto:
		// a dose que o próprio médico assinou da última vez para esta base.
		if last := s.lastSignedDose(patientID, templateID, c.Substance); last != nil {
			item.LastPrescribed = last
		}
		resp.Items = append(resp.Items, item)
	}
	return resp, nil
}

// latestWeight — consulta primeiro, depois avaliação física, por último o cadastro. Devolve
// SEMPRE a procedência e a data: dose calculada sobre peso antigo é risco, e quem lê precisa ver
// de onde o número saiu.
func (s *MagistralTemplateService) latestWeight(patientID uuid.UUID) *MeasurementInput {
	var vital struct {
		Weight *float64
		Date   *time.Time
	}
	// consultation_vitals já guarda patient_id direto; não precisa passar pela consulta.
	s.db.Table("consultation_vitals").
		Select("weight, created_at AS date").
		Where("patient_id = ? AND weight IS NOT NULL AND deleted_at IS NULL", patientID).
		Order("created_at DESC").Limit(1).Scan(&vital)
	if vital.Weight != nil {
		return &MeasurementInput{Value: *vital.Weight, Date: vital.Date, Source: "peso da consulta", Unit: "kg"}
	}

	var assessment struct {
		Weight *float64
		Date   *time.Time
	}
	s.db.Table("physical_assessments").
		Select("weight, assessment_date AS date").
		Where("patient_id = ? AND weight IS NOT NULL AND deleted_at IS NULL", patientID).
		Order("assessment_date DESC").Limit(1).Scan(&assessment)
	if assessment.Weight != nil {
		return &MeasurementInput{Value: *assessment.Weight, Date: assessment.Date, Source: "avaliação física", Unit: "kg"}
	}

	var patient struct {
		Weight *float64
		Date   *time.Time
	}
	s.db.Table("patients").Select("weight, updated_at AS date").
		Where("id = ? AND weight IS NOT NULL", patientID).Scan(&patient)
	if patient.Weight != nil {
		return &MeasurementInput{Value: *patient.Weight, Date: patient.Date, Source: "cadastro do paciente", Unit: "kg"}
	}
	return nil
}

// latestLab — valor, data da coleta, nome e unidade do exame. O nome entra na frase que o médico
// lê ("25-OH-vitamina D = 22 ng/mL de 12/07").
func (s *MagistralTemplateService) latestLab(patientID uuid.UUID, code string) *MeasurementInput {
	var row struct {
		Value *float64
		Date  *time.Time
		Name  string
		Unit  string
	}
	s.db.Table("lab_results lr").
		// A unidade que vale é a do RESULTADO; a da definição é só o fallback. 31% dos
		// resultados numéricos estão gravados em unidade diferente da definição do exame, e é
		// contra o número gravado que a regra vai comparar.
		Select("lr.result_numeric AS value, b.collection_date AS date, d.name, coalesce(nullif(lr.unit, ''), d.unit) AS unit").
		Joins("JOIN lab_test_definitions d ON d.id = lr.lab_test_definition_id").
		Joins("JOIN lab_result_batches b ON b.id = lr.lab_result_batch_id").
		Where("b.patient_id = ? AND d.code = ?", patientID, code).
		Where("lr.result_numeric IS NOT NULL AND lr.deleted_at IS NULL").
		Order("b.collection_date DESC").Limit(1).Scan(&row)
	if row.Value == nil {
		return nil
	}
	return &MeasurementInput{Value: *row.Value, Date: row.Date, Source: row.Name, Unit: row.Unit}
}

// lastSignedDose — a última dose que o médico ASSINOU desta substância nesta base. É o mais
// perto de "resposta do paciente" que o EMR tem hoje, sem inventar escala de resposta.
func (s *MagistralTemplateService) lastSignedDose(patientID, templateID uuid.UUID, substance string) *float64 {
	var row struct{ Quantity *float64 }
	s.db.Table("prescription_formula_components pfc").
		Select("pfc.quantity").
		Joins("JOIN prescription_formulas pf ON pf.id = pfc.formula_id").
		Joins("JOIN prescriptions p ON p.id = pf.prescription_id").
		Where("p.patient_id = ? AND pf.template_id = ? AND p.signed_at IS NOT NULL", patientID, templateID).
		Where("lower(public.immutable_unaccent(pfc.substance)) = lower(public.immutable_unaccent(?))", substance).
		Where("pfc.deleted_at IS NULL AND pf.deleted_at IS NULL AND p.deleted_at IS NULL").
		Order("p.signed_at DESC").Limit(1).Scan(&row)
	return row.Quantity
}
