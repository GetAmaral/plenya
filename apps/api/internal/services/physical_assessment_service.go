package services

import (
	"errors"
	"math"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

var (
	ErrPhysicalAssessmentNotFound = errors.New("physical assessment not found")
)

type PhysicalAssessmentService struct {
	db       *gorm.DB
	clinical *ClinicalDataService
}

func NewPhysicalAssessmentService(db *gorm.DB, clinical *ClinicalDataService) *PhysicalAssessmentService {
	return &PhysicalAssessmentService{db: db, clinical: clinical}
}

// Create cria uma avaliação física com dados diretos e calcula ACSM automaticamente
func (s *PhysicalAssessmentService) Create(userID uuid.UUID, req *dto.CreatePhysicalAssessmentRequest) (*dto.PhysicalAssessmentResponse, error) {
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

	assessmentDate, err := time.Parse("2006-01-02", req.AssessmentDate)
	if err != nil {
		return nil, errors.New("invalid assessment date, expected YYYY-MM-DD")
	}

	var patient models.Patient
	if err := s.db.First(&patient, patientID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPatientNotFound
		}
		return nil, err
	}
	patient.CalculateAge()

	assessment := models.PhysicalAssessment{
		PatientID:      patientID,
		CreatedByID:    userID,
		AssessmentDate: assessmentDate,

		Weight:             req.Weight,
		Height:             req.Height,
		WaistCircumference: req.WaistCircumference,
		BodyFatPercent:     req.BodyFatPercent,
		LeanMass:           req.LeanMass,
		SystolicBP:         req.SystolicBP,
		DiastolicBP:        req.DiastolicBP,
		RestingHeartRate:    req.RestingHeartRate,
		LDL:                req.LDL,
		HDL:                req.HDL,
		TotalCholesterol:   req.TotalCholesterol,
		Triglycerides:      req.Triglycerides,
		FastingGlucose:     req.FastingGlucose,
		HbA1c:              req.HbA1c,
		FamilyHistory:      req.FamilyHistory,
		CardiovascularDisease: req.CardiovascularDisease,
		DiabetesType:       req.DiabetesType,
		Symptoms:           req.Symptoms,
		ClinicalAlert:      req.ClinicalAlert,
		FrontPhotoURL:      req.FrontPhotoURL,
		SidePhotoURL:       req.SidePhotoURL,
	}

	if req.SmokingStatus != nil {
		ss := models.SmokingStatus(*req.SmokingStatus)
		assessment.SmokingStatus = &ss
	}
	if req.PhysicalActivityLevel != nil {
		pal := models.PhysicalActivityLevel(*req.PhysicalActivityLevel)
		assessment.PhysicalActivityLevel = &pal
	}

	// Auto-fill campos não fornecidos a partir da anamnese e labs
	s.autoFill(&assessment, patientID, patient.Gender)

	// Cálculos derivados
	s.computeDerived(&assessment)

	// ACSM
	riskLevel, riskCount, riskFactors, protectiveFactors := calculateACSMRisk(&assessment, &patient)
	assessment.ACSMRiskLevel = &riskLevel
	assessment.ACSMRiskFactorsCount = riskCount
	assessment.ACSMRiskFactors = riskFactors
	assessment.ACSMProtectiveFactors = protectiveFactors
	recommendation := generateACSMRecommendation(riskLevel)
	assessment.ACSMRecommendation = &recommendation
	assessment.ACSMTags = computeACSMTags(&assessment, &patient)

	if err := s.db.Create(&assessment).Error; err != nil {
		return nil, err
	}

	return s.toDTO(&assessment, patient.Age), nil
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

	var patient models.Patient
	s.db.First(&patient, assessment.PatientID)
	patient.CalculateAge()

	return s.toDTO(&assessment, patient.Age), nil
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

	var patient models.Patient
	s.db.First(&patient, *user.SelectedPatientID)
	patient.CalculateAge()

	result := make([]dto.PhysicalAssessmentResponse, len(assessments))
	for i, a := range assessments {
		result[i] = *s.toDTO(&a, patient.Age)
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

// autoFill preenche campos não fornecidos a partir da anamnese e labs do paciente
func (s *PhysicalAssessmentService) autoFill(a *models.PhysicalAssessment, patientID uuid.UUID, gender models.Gender) {
	c := s.clinical

	// Antropometria (da anamnese)
	if a.Weight == nil {
		a.Weight = c.LatestAnamnesisNumeric(patientID, models.AnamCodePeso)
	}
	if a.Height == nil {
		a.Height = c.LatestAnamnesisNumeric(patientID, models.AnamCodeAltura)
	}
	if a.WaistCircumference == nil {
		code := models.AnamCodeAbdominalHomem
		if gender == models.GenderFemale {
			code = models.AnamCodeAbdominalMulher
		}
		a.WaistCircumference = c.LatestAnamnesisNumeric(patientID, code)
	}
	if a.BodyFatPercent == nil {
		code := models.AnamCodeGorduraHomem
		if gender == models.GenderFemale {
			code = models.AnamCodeGorduraMulher
		}
		a.BodyFatPercent = c.LatestAnamnesisNumeric(patientID, code)
	}
	if a.LeanMass == nil {
		a.LeanMass = c.LatestAnamnesisNumeric(patientID, models.AnamCodeMassaMuscular)
	}

	// Fatores de risco (da anamnese via selected_level)
	if a.SmokingStatus == nil {
		if level := c.LatestAnamnesisLevel(patientID, models.AnamCodeTabaco); level != nil {
			ss := SmokingStatusFromLevel(*level)
			a.SmokingStatus = &ss
		}
	}
	if a.PhysicalActivityLevel == nil {
		if level := c.LatestAnamnesisLevel(patientID, models.AnamCodeAtividadeFisica); level != nil {
			pal := ActivityLevelFromLevel(*level)
			a.PhysicalActivityLevel = &pal
		}
	}
	if a.FamilyHistory == nil {
		if level := c.LatestAnamnesisLevel(patientID, models.AnamCodeDoencaCVFamiliar); level != nil {
			has := BoolFromRiskLevel(*level, 2)
			a.FamilyHistory = &has
		}
	}
	if a.CardiovascularDisease == nil {
		if level := c.LatestAnamnesisLevel(patientID, models.AnamCodeDoencaCV); level != nil {
			has := BoolFromRiskLevel(*level, 3)
			a.CardiovascularDisease = &has
		}
	}
	if a.Symptoms == nil {
		if level := c.LatestAnamnesisLevel(patientID, models.AnamCodeSintomas); level != nil && *level <= 2 {
			text := c.LatestAnamnesisText(patientID, models.AnamCodeSintomas)
			if text != nil {
				a.Symptoms = text
			} else {
				fallback := "Sintomas presentes"
				a.Symptoms = &fallback
			}
		}
	}

	// Laboratorial (dos LabResults)
	if a.LDL == nil {
		a.LDL = c.LatestLabNumeric(patientID, models.LabCodeLDL)
	}
	if a.HDL == nil {
		a.HDL = c.LatestLabNumeric(patientID, models.LabCodeHDL)
	}
	if a.TotalCholesterol == nil {
		a.TotalCholesterol = c.LatestLabNumeric(patientID, models.LabCodeColesterolTotal)
	}
	if a.Triglycerides == nil {
		a.Triglycerides = c.LatestLabNumeric(patientID, models.LabCodeTriglicerideos)
	}
	if a.FastingGlucose == nil {
		a.FastingGlucose = c.LatestLabNumeric(patientID, models.LabCodeGlicemiaJejum)
	}
	if a.HbA1c == nil {
		a.HbA1c = c.LatestLabNumeric(patientID, models.LabCodeHbA1c)
	}
}

// computeDerived calcula BMI, BRI e massa magra a partir dos dados disponíveis
func (s *PhysicalAssessmentService) computeDerived(a *models.PhysicalAssessment) {
	if a.Weight != nil && a.Height != nil && *a.Height > 0 {
		bmi := CalculateBMI(*a.Weight, *a.Height)
		a.BMI = &bmi
	}
	if a.WaistCircumference != nil && a.Height != nil && *a.Height > 0 {
		bri := CalculateBRI(*a.WaistCircumference, *a.Height)
		a.BRI = &bri
	}
	if a.LeanMass == nil && a.BodyFatPercent != nil && a.Weight != nil {
		lm := *a.Weight * (1 - *a.BodyFatPercent/100)
		a.LeanMass = &lm
	}
}

// --- ACSM ---

func calculateACSMRisk(a *models.PhysicalAssessment, patient *models.Patient) (models.ACSMRiskLevel, int, []string, []string) {
	riskFactors := []string{}
	protectiveFactors := []string{}

	// 1. Idade (M>=45, F>=55)
	if (patient.Gender == models.GenderMale && patient.Age >= 45) ||
		(patient.Gender == models.GenderFemale && patient.Age >= 55) {
		riskFactors = append(riskFactors, "idade")
	}
	// 2. Histórico familiar CV
	if a.FamilyHistory != nil && *a.FamilyHistory {
		riskFactors = append(riskFactors, "historico_familiar_cv")
	}
	// 3. Tabagismo
	if a.SmokingStatus != nil && *a.SmokingStatus == models.SmokingCurrent {
		riskFactors = append(riskFactors, "tabagismo")
	}
	// 4. Sedentarismo
	if a.PhysicalActivityLevel != nil &&
		(*a.PhysicalActivityLevel == models.ActivitySedentary || *a.PhysicalActivityLevel == models.ActivityInsufficient) {
		riskFactors = append(riskFactors, "sedentarismo")
	}
	// 5. Obesidade (BMI >= 30 ou cintura M>102/F>88)
	isObese := (a.BMI != nil && *a.BMI >= 30)
	if a.WaistCircumference != nil {
		if (patient.Gender == models.GenderMale && *a.WaistCircumference > 102) ||
			(patient.Gender == models.GenderFemale && *a.WaistCircumference > 88) {
			isObese = true
		}
	}
	if isObese {
		riskFactors = append(riskFactors, "obesidade")
	}
	// 6. Dislipidemia (LDL >= 130 ou HDL < 40)
	if (a.LDL != nil && *a.LDL >= 130) || (a.HDL != nil && *a.HDL < 40) {
		riskFactors = append(riskFactors, "dislipidemia")
	}
	// 7. Pré-diabetes (Glicemia >= 100 ou HbA1c >= 5.7)
	if (a.FastingGlucose != nil && *a.FastingGlucose >= 100) || (a.HbA1c != nil && *a.HbA1c >= 5.7) {
		riskFactors = append(riskFactors, "pre_diabetes")
	}
	// Fator protetor: HDL >= 60
	if a.HDL != nil && *a.HDL >= 60 {
		protectiveFactors = append(protectiveFactors, "hdl_alto")
	}

	riskCount := len(riskFactors) - len(protectiveFactors)
	if riskCount < 0 {
		riskCount = 0
	}

	// Alto: sintomas ou doença CV pessoal
	hasSymptoms := a.Symptoms != nil && *a.Symptoms != ""
	hasCV := a.CardiovascularDisease != nil && *a.CardiovascularDisease
	if hasSymptoms || hasCV {
		return models.ACSMRiskHigh, riskCount, riskFactors, protectiveFactors
	}
	if riskCount <= 1 {
		return models.ACSMRiskLow, riskCount, riskFactors, protectiveFactors
	}
	return models.ACSMRiskModerate, riskCount, riskFactors, protectiveFactors
}

func computeACSMTags(a *models.PhysicalAssessment, patient *models.Patient) []string {
	tags := []string{}

	switch {
	case patient.Age < 18:
		tags = append(tags, "<18 anos")
	case patient.Age <= 29:
		tags = append(tags, "18-29 anos")
	case patient.Age <= 39:
		tags = append(tags, "30-39 anos")
	case patient.Age <= 49:
		tags = append(tags, "40-49 anos")
	case patient.Age <= 59:
		tags = append(tags, "50-59 anos")
	default:
		tags = append(tags, "60+ anos")
	}

	switch patient.Gender {
	case models.GenderMale:
		tags = append(tags, "MASCULINO")
	case models.GenderFemale:
		tags = append(tags, "FEMININO")
	}

	hasSymptoms := a.Symptoms != nil && *a.Symptoms != ""
	hasCV := a.CardiovascularDisease != nil && *a.CardiovascularDisease
	hasRiskFactors := a.ACSMRiskFactorsCount >= 2

	switch {
	case hasRiskFactors && (hasSymptoms || hasCV):
		tags = append(tags, "RYSH")
	case hasRiskFactors:
		tags = append(tags, "RYNH")
	case hasSymptoms || hasCV:
		tags = append(tags, "NYSH")
	default:
		tags = append(tags, "NYNH")
	}

	if a.PhysicalActivityLevel != nil &&
		(*a.PhysicalActivityLevel == models.ActivitySedentary || *a.PhysicalActivityLevel == models.ActivityInsufficient) {
		tags = append(tags, "SEDENTARIO")
	} else {
		tags = append(tags, "ATIVO")
	}

	if a.BMI != nil {
		switch {
		case *a.BMI < 18.5:
			tags = append(tags, "BAIXO_PESO")
		case *a.BMI < 25:
			tags = append(tags, "PESO_NORMAL")
		case *a.BMI < 30:
			tags = append(tags, "SOBREPESO")
		case *a.BMI < 35:
			tags = append(tags, "OBESIDADE_I")
		case *a.BMI < 40:
			tags = append(tags, "OBESIDADE_II")
		default:
			tags = append(tags, "OBESIDADE_III")
		}
	}

	if a.BodyFatPercent != nil {
		bf := *a.BodyFatPercent
		if patient.Gender == models.GenderMale {
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

	if a.ACSMRiskLevel != nil {
		switch *a.ACSMRiskLevel {
		case models.ACSMRiskLow:
			tags = append(tags, "RISCO_BAIXO")
		case models.ACSMRiskModerate:
			tags = append(tags, "RISCO_MODERADO")
		case models.ACSMRiskHigh:
			tags = append(tags, "RISCO_ALTO")
		}
	}

	return tags
}

// --- Fórmulas ---

func CalculateBMI(weightKg, heightCm float64) float64 {
	heightM := heightCm / 100
	if heightM <= 0 {
		return 0
	}
	return weightKg / (heightM * heightM)
}

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

func CalculateMaxHR(age int) int {
	return 220 - age
}

func CalculateKarvonenZones(maxHR, restHR int) []dto.HRZone {
	reserve := maxHR - restHR
	return []dto.HRZone{
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
}

// --- DTO ---

func (s *PhysicalAssessmentService) toDTO(pa *models.PhysicalAssessment, patientAge int) *dto.PhysicalAssessmentResponse {
	resp := &dto.PhysicalAssessmentResponse{
		ID:                    pa.ID.String(),
		PatientID:             pa.PatientID.String(),
		CreatedByID:           pa.CreatedByID.String(),
		AssessmentDate:        pa.AssessmentDate.Format("2006-01-02"),
		Weight:                pa.Weight,
		Height:                pa.Height,
		WaistCircumference:    pa.WaistCircumference,
		BMI:                   pa.BMI,
		BRI:                   pa.BRI,
		BodyFatPercent:        pa.BodyFatPercent,
		LeanMass:              pa.LeanMass,
		SystolicBP:            pa.SystolicBP,
		DiastolicBP:           pa.DiastolicBP,
		RestingHeartRate:       pa.RestingHeartRate,
		LDL:                   pa.LDL,
		HDL:                   pa.HDL,
		TotalCholesterol:      pa.TotalCholesterol,
		Triglycerides:         pa.Triglycerides,
		FastingGlucose:        pa.FastingGlucose,
		HbA1c:                 pa.HbA1c,
		FamilyHistory:         pa.FamilyHistory,
		CardiovascularDisease: pa.CardiovascularDisease,
		DiabetesType:          pa.DiabetesType,
		Symptoms:              pa.Symptoms,
		ClinicalAlert:         pa.ClinicalAlert,
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
	if pa.SmokingStatus != nil {
		ss := string(*pa.SmokingStatus)
		resp.SmokingStatus = &ss
	}
	if pa.PhysicalActivityLevel != nil {
		pal := string(*pa.PhysicalActivityLevel)
		resp.PhysicalActivityLevel = &pal
	}

	maxHR := CalculateMaxHR(patientAge)
	resp.MaxHR = &maxHR
	restHR := 70
	if pa.RestingHeartRate != nil {
		restHR = *pa.RestingHeartRate
	}
	resp.HRZones = CalculateKarvonenZones(maxHR, restHR)

	return resp
}

func generateACSMRecommendation(level models.ACSMRiskLevel) string {
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
