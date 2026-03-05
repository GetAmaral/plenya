package services

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

// LabTestDefinition.Code constants (fonte: lab_test_definitions.code)
const (
	LabCodeLDL             = "PLNACA1492A"
	LabCodeHDL             = "PLN53A449CA"
	LabCodeColesterolTotal = "PLN919303A4"
	LabCodeTriglicerideos  = "PLNA0C5545F"
	LabCodeGlicemiaJejum   = "PLN9AF0BCF5"
	LabCodeHbA1c           = "PLN3FC5EDA6"
)

// ScoreItem.AnamneseItemCode constants (fonte: score_items.anamnese_item_code)
const (
	AnamCodePeso            = "PESO"
	AnamCodeAltura          = "ALTURA"
	AnamCodeIMC             = "IMC"
	AnamCodeBRI             = "BRI"
	AnamCodeAbdominalHomem  = "ABDOMINAL_HOMEM"
	AnamCodeAbdominalMulher = "ABDOMINAL_MULHER"
	AnamCodeGorduraHomem    = "GORDURA_CORPORAL_HOMEM"
	AnamCodeGorduraMulher   = "GORDURA_CORPORAL_MULHER"
	AnamCodeMassaMuscular   = "MASSA_MUSCULAR_ESQUELETICA"
	AnamCodeTaxaMetabolica  = "TAXA_METABOLICA_BASAL"

	AnamCodeTabaco                       = "TABACO"
	AnamCodeDoencaCV                     = "DOENCA_CARDIOVASCULAR"
	AnamCodeDoencaCVFamiliar             = "DOENCA_CARDIOVASCULAR_2"
	AnamCodeHistFamiliarComposicao       = "HISTORICO_FAMILIAR_DE_COMPOSICAO_CORPORAL"
	AnamCodeSintomas                     = "OUTROS_SINTOMAS"
	AnamCodeAtividadeFisica              = "ATIVIDADE_FISICA"
	AnamCodeExercicioFisico              = "EXERCICIO_FISICO"
)

var (
	ErrPhysicalAssessmentNotFound = errors.New("physical assessment not found")
)

// ACSMInputData contém dados coletados para cálculo ACSM
type ACSMInputData struct {
	Age              int
	Gender           models.Gender
	Weight           *float64
	Height           *float64
	BMI              *float64
	BRI              *float64
	WaistCm          *float64
	BodyFatPercent   *float64
	LDL              *float64
	HDL              *float64
	TotalChol        *float64
	Triglycerides    *float64
	FastingGlucose   *float64
	HbA1c            *float64
	SmokingStatus    string
	ActivityLevel    string
	FamilyHistoryCV  string
	PersonalCV       string
	Symptoms         string
}

type PhysicalAssessmentService struct {
	db *gorm.DB
}

func NewPhysicalAssessmentService(db *gorm.DB) *PhysicalAssessmentService {
	return &PhysicalAssessmentService{db: db}
}

// Create cria uma avaliação física e calcula ACSM automaticamente
func (s *PhysicalAssessmentService) Create(userID uuid.UUID, req *dto.CreatePhysicalAssessmentRequest) (*dto.PhysicalAssessmentResponse, error) {
	// Get selected patient
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}
	if user.SelectedPatientID == nil {
		return nil, errors.New("no patient selected")
	}

	patientID := *user.SelectedPatientID
	if req.PatientID != "" {
		pid, err := uuid.Parse(req.PatientID)
		if err != nil {
			return nil, errors.New("invalid patient id")
		}
		if pid != patientID {
			return nil, errors.New("patient id does not match selected patient")
		}
	}

	anamnesisID, err := uuid.Parse(req.AnamnesisID)
	if err != nil {
		return nil, errors.New("invalid anamnesis id")
	}

	assessmentDate, err := time.Parse("2006-01-02", req.AssessmentDate)
	if err != nil {
		return nil, errors.New("invalid assessment date, expected YYYY-MM-DD")
	}

	// Verify patient exists
	var patient models.Patient
	if err := s.db.First(&patient, patientID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPatientNotFound
		}
		return nil, err
	}

	// Verify anamnesis exists and belongs to patient
	var anamnesis models.Anamnesis
	if err := s.db.Preload("Items").Preload("Items.ScoreItem").
		Where("id = ? AND patient_id = ?", anamnesisID, patientID).
		First(&anamnesis).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAnamnesisNotFound
		}
		return nil, err
	}

	// Gather ACSM data from existing sources
	data := s.gatherACSMData(&patient, anamnesis.Items)

	// Calculate ACSM risk
	riskLevel, riskCount, riskFactors, protectiveFactors := CalculateACSMRisk(data)

	// Compute tags
	tags := ComputeACSMTags(data)

	// Generate recommendation text
	recommendation := generateACSMRecommendation(riskLevel, riskFactors, protectiveFactors)

	assessment := models.PhysicalAssessment{
		PatientID:             patientID,
		CreatedByID:           userID,
		AssessmentDate:        assessmentDate,
		AnamnesisID:           anamnesisID,
		ACSMRiskLevel:         &riskLevel,
		ACSMRiskFactorsCount:  riskCount,
		ACSMRiskFactors:       riskFactors,
		ACSMProtectiveFactors: protectiveFactors,
		ACSMRecommendation:    &recommendation,
		ACSMTags:              tags,
	}

	if err := s.db.Create(&assessment).Error; err != nil {
		return nil, err
	}

	return s.toDTO(&assessment, &data), nil
}

// GetByID busca uma avaliação física por ID
func (s *PhysicalAssessmentService) GetByID(id, userID uuid.UUID) (*dto.PhysicalAssessmentResponse, error) {
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}
	if user.SelectedPatientID == nil {
		return nil, ErrPhysicalAssessmentNotFound
	}

	var assessment models.PhysicalAssessment
	if err := s.db.Where("id = ? AND patient_id = ?", id, *user.SelectedPatientID).
		First(&assessment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPhysicalAssessmentNotFound
		}
		return nil, err
	}

	// Load patient and anamnesis for resolved data
	var patient models.Patient
	s.db.First(&patient, assessment.PatientID)

	var anamnesis models.Anamnesis
	s.db.Preload("Items").Preload("Items.ScoreItem").First(&anamnesis, assessment.AnamnesisID)

	data := s.gatherACSMData(&patient, anamnesis.Items)

	return s.toDTO(&assessment, &data), nil
}

// List lista avaliações físicas do paciente selecionado
func (s *PhysicalAssessmentService) List(userID uuid.UUID, limit, offset int) ([]dto.PhysicalAssessmentResponse, error) {
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}
	if user.SelectedPatientID == nil {
		return []dto.PhysicalAssessmentResponse{}, nil
	}

	var assessments []models.PhysicalAssessment
	if err := s.db.Where("patient_id = ?", *user.SelectedPatientID).
		Order("assessment_date DESC").
		Limit(limit).Offset(offset).
		Find(&assessments).Error; err != nil {
		return nil, err
	}

	result := make([]dto.PhysicalAssessmentResponse, len(assessments))
	for i, a := range assessments {
		result[i] = *s.toDTO(&a, nil)
	}
	return result, nil
}

// Delete faz soft delete de uma avaliação física
func (s *PhysicalAssessmentService) Delete(id uuid.UUID, userRole models.Role) error {
	if userRole != models.RoleAdmin {
		return ErrUnauthorized
	}
	result := s.db.Where("id = ?", id).Delete(&models.PhysicalAssessment{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrPhysicalAssessmentNotFound
	}
	return nil
}

// gatherACSMData coleta dados de AnamnesisItems e LabResults
func (s *PhysicalAssessmentService) gatherACSMData(patient *models.Patient, anamnesisItems []models.AnamnesisItem) ACSMInputData {
	data := ACSMInputData{
		Age:    patient.Age,
		Gender: patient.Gender,
	}

	// From AnamnesisItems via ScoreItem.AnamneseItemCode
	for _, item := range anamnesisItems {
		if item.ScoreItem == nil || item.ScoreItem.AnamneseItemCode == nil {
			continue
		}
		code := *item.ScoreItem.AnamneseItemCode

		switch code {
		case AnamCodePeso:
			data.Weight = item.NumericValue
		case AnamCodeAltura:
			data.Height = item.NumericValue
		case AnamCodeIMC:
			data.BMI = item.NumericValue
		case AnamCodeBRI:
			data.BRI = item.NumericValue
		case AnamCodeAbdominalHomem:
			if patient.Gender == models.GenderMale {
				data.WaistCm = item.NumericValue
			}
		case AnamCodeAbdominalMulher:
			if patient.Gender == models.GenderFemale {
				data.WaistCm = item.NumericValue
			}
		case AnamCodeGorduraHomem:
			if patient.Gender == models.GenderMale {
				data.BodyFatPercent = item.NumericValue
			}
		case AnamCodeGorduraMulher:
			if patient.Gender == models.GenderFemale {
				data.BodyFatPercent = item.NumericValue
			}
		case AnamCodeTabaco:
			if item.TextValue != nil {
				data.SmokingStatus = *item.TextValue
			}
		case AnamCodeAtividadeFisica, AnamCodeExercicioFisico:
			if item.TextValue != nil {
				data.ActivityLevel = *item.TextValue
			}
		case AnamCodeDoencaCVFamiliar:
			if item.TextValue != nil {
				data.FamilyHistoryCV = *item.TextValue
			}
		case AnamCodeDoencaCV:
			if item.TextValue != nil {
				data.PersonalCV = *item.TextValue
			}
		case AnamCodeSintomas:
			if item.TextValue != nil {
				data.Symptoms = *item.TextValue
			}
		}
	}

	// Calculate BMI if not provided
	if data.BMI == nil && data.Weight != nil && data.Height != nil && *data.Height > 0 {
		bmi := CalculateBMI(*data.Weight, *data.Height)
		data.BMI = &bmi
	}

	// Calculate BRI if not provided
	if data.BRI == nil && data.WaistCm != nil && data.Height != nil && *data.Height > 0 {
		bri := CalculateBRI(*data.WaistCm, *data.Height)
		data.BRI = &bri
	}

	// From LabResults via LabTestDefinition.Code (most recent per code)
	s.loadLatestLabValue(patient.ID, LabCodeLDL, &data.LDL)
	s.loadLatestLabValue(patient.ID, LabCodeHDL, &data.HDL)
	s.loadLatestLabValue(patient.ID, LabCodeColesterolTotal, &data.TotalChol)
	s.loadLatestLabValue(patient.ID, LabCodeTriglicerideos, &data.Triglycerides)
	s.loadLatestLabValue(patient.ID, LabCodeGlicemiaJejum, &data.FastingGlucose)
	s.loadLatestLabValue(patient.ID, LabCodeHbA1c, &data.HbA1c)

	return data
}

// loadLatestLabValue busca o valor mais recente de um lab test por código
func (s *PhysicalAssessmentService) loadLatestLabValue(patientID uuid.UUID, labTestCode string, target **float64) {
	var result struct {
		ResultNumeric *float64
	}
	err := s.db.Raw(`
		SELECT lr.result_numeric
		FROM lab_results lr
		JOIN lab_test_definitions ltd ON lr.lab_test_definition_id = ltd.id
		JOIN lab_result_batches lrb ON lr.lab_result_batch_id = lrb.id
		WHERE lrb.patient_id = ? AND ltd.code = ? AND lr.result_numeric IS NOT NULL AND lr.deleted_at IS NULL
		ORDER BY lrb.collection_date DESC
		LIMIT 1
	`, patientID, labTestCode).Scan(&result).Error

	if err == nil && result.ResultNumeric != nil {
		*target = result.ResultNumeric
	}
}

// CalculateACSMRisk calcula a estratificação de risco ACSM 2018
func CalculateACSMRisk(data ACSMInputData) (models.ACSMRiskLevel, int, []string, []string) {
	riskFactors := []string{}
	protectiveFactors := []string{}

	// 1. Idade (M>=45, F>=55)
	if (data.Gender == models.GenderMale && data.Age >= 45) ||
		(data.Gender == models.GenderFemale && data.Age >= 55) {
		riskFactors = append(riskFactors, "idade")
	}

	// 2. Histórico familiar CV
	if hasPositiveValue(data.FamilyHistoryCV) {
		riskFactors = append(riskFactors, "historico_familiar_cv")
	}

	// 3. Tabagismo
	if hasPositiveValue(data.SmokingStatus) {
		riskFactors = append(riskFactors, "tabagismo")
	}

	// 4. Sedentarismo
	if isSedentary(data.ActivityLevel) {
		riskFactors = append(riskFactors, "sedentarismo")
	}

	// 5. Obesidade (BMI >= 30 ou cintura M>102/F>88)
	isObese := false
	if data.BMI != nil && *data.BMI >= 30 {
		isObese = true
	}
	if data.WaistCm != nil {
		if (data.Gender == models.GenderMale && *data.WaistCm > 102) ||
			(data.Gender == models.GenderFemale && *data.WaistCm > 88) {
			isObese = true
		}
	}
	if isObese {
		riskFactors = append(riskFactors, "obesidade")
	}

	// 6. Dislipidemia (LDL >= 130 ou HDL < 40)
	isDyslipidemic := false
	if data.LDL != nil && *data.LDL >= 130 {
		isDyslipidemic = true
	}
	if data.HDL != nil && *data.HDL < 40 {
		isDyslipidemic = true
	}
	if isDyslipidemic {
		riskFactors = append(riskFactors, "dislipidemia")
	}

	// 7. Pré-diabetes (Glicemia >= 100 ou HbA1c >= 5.7)
	isPreDiabetic := false
	if data.FastingGlucose != nil && *data.FastingGlucose >= 100 {
		isPreDiabetic = true
	}
	if data.HbA1c != nil && *data.HbA1c >= 5.7 {
		isPreDiabetic = true
	}
	if isPreDiabetic {
		riskFactors = append(riskFactors, "pre_diabetes")
	}

	// Fator protetor: HDL >= 60
	if data.HDL != nil && *data.HDL >= 60 {
		protectiveFactors = append(protectiveFactors, "hdl_alto")
	}

	// Net risk count
	riskCount := len(riskFactors) - len(protectiveFactors)
	if riskCount < 0 {
		riskCount = 0
	}

	// High: sintomas ou doença CV pessoal
	if hasPositiveValue(data.Symptoms) || hasPositiveValue(data.PersonalCV) {
		return models.ACSMRiskHigh, riskCount, riskFactors, protectiveFactors
	}

	// Low: <=1 | Moderate: >=2
	if riskCount <= 1 {
		return models.ACSMRiskLow, riskCount, riskFactors, protectiveFactors
	}
	return models.ACSMRiskModerate, riskCount, riskFactors, protectiveFactors
}

// ComputeACSMTags replica funcoes_tags_v2.py
func ComputeACSMTags(data ACSMInputData) []string {
	tags := []string{}

	// Age bracket
	switch {
	case data.Age < 18:
		tags = append(tags, "<18 anos")
	case data.Age <= 29:
		tags = append(tags, "18-29 anos")
	case data.Age <= 39:
		tags = append(tags, "30-39 anos")
	case data.Age <= 49:
		tags = append(tags, "40-49 anos")
	case data.Age <= 59:
		tags = append(tags, "50-59 anos")
	default:
		tags = append(tags, "60+ anos")
	}

	// Gender
	switch data.Gender {
	case models.GenderMale:
		tags = append(tags, "MASCULINO")
	case models.GenderFemale:
		tags = append(tags, "FEMININO")
	}

	// Risk matrix (RY/NY x RH/NH)
	hasRiskFactors := false
	_, riskCount, _, _ := CalculateACSMRisk(data)
	if riskCount >= 2 {
		hasRiskFactors = true
	}
	hasSymptoms := hasPositiveValue(data.Symptoms) || hasPositiveValue(data.PersonalCV)

	switch {
	case hasRiskFactors && hasSymptoms:
		tags = append(tags, "RYSH")
	case hasRiskFactors && !hasSymptoms:
		tags = append(tags, "RYNH")
	case !hasRiskFactors && hasSymptoms:
		tags = append(tags, "NYSH")
	default:
		tags = append(tags, "NYNH")
	}

	// Activity level
	if isSedentary(data.ActivityLevel) {
		tags = append(tags, "SEDENTARIO")
	} else {
		tags = append(tags, "ATIVO")
	}

	// BMI class
	if data.BMI != nil {
		switch {
		case *data.BMI < 18.5:
			tags = append(tags, "BAIXO_PESO")
		case *data.BMI < 25:
			tags = append(tags, "PESO_NORMAL")
		case *data.BMI < 30:
			tags = append(tags, "SOBREPESO")
		case *data.BMI < 35:
			tags = append(tags, "OBESIDADE_I")
		case *data.BMI < 40:
			tags = append(tags, "OBESIDADE_II")
		default:
			tags = append(tags, "OBESIDADE_III")
		}
	}

	// BF% class
	if data.BodyFatPercent != nil {
		bf := *data.BodyFatPercent
		if data.Gender == models.GenderMale {
			switch {
			case bf < 6:
				tags = append(tags, "GC_ESSENCIAL")
			case bf < 14:
				tags = append(tags, "GC_ATLETICO")
			case bf < 18:
				tags = append(tags, "GC_FITNESS")
			case bf < 25:
				tags = append(tags, "GC_ACEITAVEL")
			default:
				tags = append(tags, "GC_OBESIDADE")
			}
		} else {
			switch {
			case bf < 14:
				tags = append(tags, "GC_ESSENCIAL")
			case bf < 21:
				tags = append(tags, "GC_ATLETICO")
			case bf < 25:
				tags = append(tags, "GC_FITNESS")
			case bf < 32:
				tags = append(tags, "GC_ACEITAVEL")
			default:
				tags = append(tags, "GC_OBESIDADE")
			}
		}
	}

	// Final risk badge
	riskLevel, _, _, _ := CalculateACSMRisk(data)
	switch riskLevel {
	case models.ACSMRiskLow:
		tags = append(tags, "RISCO_BAIXO")
	case models.ACSMRiskModerate:
		tags = append(tags, "RISCO_MODERADO")
	case models.ACSMRiskHigh:
		tags = append(tags, "RISCO_ALTO")
	}

	return tags
}

// CalculateBMI calcula o IMC
func CalculateBMI(weightKg, heightCm float64) float64 {
	heightM := heightCm / 100
	if heightM <= 0 {
		return 0
	}
	return weightKg / (heightM * heightM)
}

// CalculateBRI calcula o Body Roundness Index
func CalculateBRI(waistCm, heightCm float64) float64 {
	waistM := waistCm / 100
	heightM := heightCm / 100
	if heightM <= 0 {
		return 0
	}
	ratio := waistM / (heightM * math.Pi)
	inner := 1 - (ratio * ratio)
	if inner < 0 {
		inner = 0
	}
	return 364.2 - 365.5*math.Sqrt(inner)
}

// CalculateMaxHR calcula a frequência cardíaca máxima
func CalculateMaxHR(age int) int {
	return 220 - age
}

// CalculateKarvonenZones calcula as zonas de FC pelo método Karvonen
func CalculateKarvonenZones(maxHR, restHR int) []dto.HRZone {
	reserve := maxHR - restHR
	zones := []dto.HRZone{
		{Zone: 1, Name: "Recuperação", Percent: "50-60%",
			MinBPM: restHR + int(float64(reserve)*0.50),
			MaxBPM: restHR + int(float64(reserve)*0.60)},
		{Zone: 2, Name: "Aeróbico Leve", Percent: "60-70%",
			MinBPM: restHR + int(float64(reserve)*0.60),
			MaxBPM: restHR + int(float64(reserve)*0.70)},
		{Zone: 3, Name: "Aeróbico Moderado", Percent: "70-80%",
			MinBPM: restHR + int(float64(reserve)*0.70),
			MaxBPM: restHR + int(float64(reserve)*0.80)},
		{Zone: 4, Name: "Limiar Anaeróbico", Percent: "80-90%",
			MinBPM: restHR + int(float64(reserve)*0.80),
			MaxBPM: restHR + int(float64(reserve)*0.90)},
		{Zone: 5, Name: "VO2max", Percent: "90-100%",
			MinBPM: restHR + int(float64(reserve)*0.90),
			MaxBPM: maxHR},
	}
	return zones
}

// toDTO converte PhysicalAssessment para PhysicalAssessmentResponse
func (s *PhysicalAssessmentService) toDTO(pa *models.PhysicalAssessment, data *ACSMInputData) *dto.PhysicalAssessmentResponse {
	resp := &dto.PhysicalAssessmentResponse{
		ID:                    pa.ID.String(),
		PatientID:             pa.PatientID.String(),
		CreatedByID:           pa.CreatedByID.String(),
		AssessmentDate:        pa.AssessmentDate.Format("2006-01-02"),
		AnamnesisID:           pa.AnamnesisID.String(),
		ACSMRiskFactorsCount:  pa.ACSMRiskFactorsCount,
		ACSMRiskFactors:       pa.ACSMRiskFactors,
		ACSMProtectiveFactors: pa.ACSMProtectiveFactors,
		ACSMRecommendation:    pa.ACSMRecommendation,
		ACSMTags:              pa.ACSMTags,
		FrontPhotoURL:         pa.FrontPhotoURL,
		SidePhotoURL:          pa.SidePhotoURL,
		AIRecommendation:      pa.AIRecommendation,
		DisplayTitle:          pa.DisplayTitle,
		CreatedAt:             pa.CreatedAt.Format(time.RFC3339),
		UpdatedAt:             pa.UpdatedAt.Format(time.RFC3339),
	}

	if pa.ACSMRiskLevel != nil {
		level := string(*pa.ACSMRiskLevel)
		resp.ACSMRiskLevel = &level
	}

	// Resolved data (computed, not persisted)
	if data != nil {
		resolved := &dto.ResolvedData{
			Weight:         data.Weight,
			Height:         data.Height,
			BMI:            data.BMI,
			BRI:            data.BRI,
			WaistCm:        data.WaistCm,
			BodyFatPercent: data.BodyFatPercent,
			LDL:            data.LDL,
			HDL:            data.HDL,
			TotalChol:      data.TotalChol,
			Triglycerides:  data.Triglycerides,
			FastingGlucose: data.FastingGlucose,
			HbA1c:          data.HbA1c,
		}

		maxHR := CalculateMaxHR(data.Age)
		resolved.MaxHR = &maxHR

		// Default rest HR of 70 if not available
		resolved.HRZones = CalculateKarvonenZones(maxHR, 70)

		resp.ResolvedData = resolved
	}

	return resp
}

// Helpers

func hasPositiveValue(s string) bool {
	if s == "" {
		return false
	}
	lower := strings.ToLower(s)
	positives := []string{"sim", "yes", "true", "1", "positivo", "presente", "ativo", "fumante"}
	for _, p := range positives {
		if strings.Contains(lower, p) {
			return true
		}
	}
	// Check if numeric value > 0
	if v, err := strconv.ParseFloat(s, 64); err == nil && v > 0 {
		return true
	}
	return false
}

func isSedentary(activityLevel string) bool {
	if activityLevel == "" {
		return true // Default to sedentary if unknown
	}
	lower := strings.ToLower(activityLevel)
	sedentaryTerms := []string{"sedentario", "sedentário", "inativo", "nenhum", "não", "nao", "never", "0"}
	for _, t := range sedentaryTerms {
		if strings.Contains(lower, t) {
			return true
		}
	}
	return false
}

func generateACSMRecommendation(level models.ACSMRiskLevel, riskFactors, protectiveFactors []string) string {
	switch level {
	case models.ACSMRiskLow:
		return "Risco baixo. Paciente pode iniciar programa de exercícios de intensidade leve a moderada sem necessidade de avaliação médica prévia. Recomendado acompanhamento periódico."
	case models.ACSMRiskModerate:
		return "Risco moderado. Recomendada avaliação médica antes de iniciar programa de exercícios de alta intensidade. Exercícios de intensidade leve a moderada podem ser iniciados com supervisão profissional."
	case models.ACSMRiskHigh:
		return "Risco alto. Avaliação médica obrigatória antes de iniciar qualquer programa de exercícios. Necessário teste de esforço e liberação cardiológica. Exercícios apenas sob supervisão direta de profissional qualificado."
	default:
		return "Estratificação de risco não calculada. Dados insuficientes."
	}
}
