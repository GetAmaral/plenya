package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

var (
	ErrFitnessTestNotFound = errors.New("fitness test not found")
)

type FitnessTestService struct {
	db *gorm.DB
}

func NewFitnessTestService(db *gorm.DB) *FitnessTestService {
	return &FitnessTestService{db: db}
}

// Create cria um resultado de teste de fitness com classificacoes automaticas
func (s *FitnessTestService) Create(userID uuid.UUID, req *dto.CreateFitnessTestRequest) (*dto.FitnessTestResponse, error) {
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

	result := models.FitnessTestResult{
		PatientID:      patientID,
		CreatedByID:    userID,
		AssessmentDate: assessmentDate,
		AbdominalReps:  req.AbdominalReps,
		PushupReps:     req.PushupReps,
		PlankSeconds:   req.PlankSeconds,
		BurpeeCycles:   req.BurpeeCycles,
		FRTReps:        req.FRTReps,
		Notes:          req.Notes,
	}

	// Classify each test
	age := patient.Age
	gender := patient.Gender

	if req.AbdominalReps != nil {
		level := classifyTest(abdominalNorms, gender, age, *req.AbdominalReps)
		result.AbdominalLevel = &level
	}
	if req.PushupReps != nil {
		level := classifyTest(pushupNorms, gender, age, *req.PushupReps)
		result.PushupLevel = &level
	}
	if req.PlankSeconds != nil {
		level := classifyTest(plankNorms, gender, age, *req.PlankSeconds)
		result.PlankLevel = &level
	}
	if req.BurpeeCycles != nil {
		level := classifyTest(burpeeNorms, gender, age, *req.BurpeeCycles)
		result.BurpeeLevel = &level
	}
	if req.FRTReps != nil {
		level := classifyTest(frtNorms, gender, age, *req.FRTReps)
		result.FRTLevel = &level
	}

	score, classification := calculateOverallScore(&result)
	result.OverallScore = score
	result.OverallClassification = classification

	if err := s.db.Create(&result).Error; err != nil {
		return nil, err
	}

	return s.toDTO(&result), nil
}

// GetByID busca um teste de fitness por ID
func (s *FitnessTestService) GetByID(id, userID uuid.UUID) (*dto.FitnessTestResponse, error) {
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}
	if user.SelectedPatientID == nil {
		return nil, ErrFitnessTestNotFound
	}

	var result models.FitnessTestResult
	if err := s.db.Where("id = ? AND patient_id = ?", id, *user.SelectedPatientID).
		First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrFitnessTestNotFound
		}
		return nil, err
	}

	return s.toDTO(&result), nil
}

// List lista testes de fitness do paciente selecionado
func (s *FitnessTestService) List(userID uuid.UUID, limit, offset int) ([]dto.FitnessTestResponse, error) {
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}
	if user.SelectedPatientID == nil {
		return []dto.FitnessTestResponse{}, nil
	}

	var results []models.FitnessTestResult
	if err := s.db.Where("patient_id = ?", *user.SelectedPatientID).
		Order("assessment_date DESC").
		Limit(limit).Offset(offset).
		Find(&results).Error; err != nil {
		return nil, err
	}

	resp := make([]dto.FitnessTestResponse, len(results))
	for i, r := range results {
		resp[i] = *s.toDTO(&r)
	}
	return resp, nil
}

// Delete faz soft delete de um teste de fitness
func (s *FitnessTestService) Delete(id uuid.UUID, userRole models.Role) error {
	if userRole != models.RoleAdmin {
		return ErrUnauthorized
	}
	res := s.db.Where("id = ?", id).Delete(&models.FitnessTestResult{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return ErrFitnessTestNotFound
	}
	return nil
}

// toDTO converte model para response DTO
func (s *FitnessTestService) toDTO(f *models.FitnessTestResult) *dto.FitnessTestResponse {
	return &dto.FitnessTestResponse{
		ID:                    f.ID.String(),
		PatientID:             f.PatientID.String(),
		CreatedByID:           f.CreatedByID.String(),
		AssessmentDate:        f.AssessmentDate.Format("2006-01-02"),
		AbdominalReps:         f.AbdominalReps,
		PushupReps:            f.PushupReps,
		PlankSeconds:          f.PlankSeconds,
		BurpeeCycles:          f.BurpeeCycles,
		FRTReps:               f.FRTReps,
		AbdominalLevel:        f.AbdominalLevel,
		PushupLevel:           f.PushupLevel,
		PlankLevel:            f.PlankLevel,
		BurpeeLevel:           f.BurpeeLevel,
		FRTLevel:              f.FRTLevel,
		OverallScore:          f.OverallScore,
		OverallClassification: f.OverallClassification,
		Notes:                 f.Notes,
		CreatedAt:             f.CreatedAt.Format(time.RFC3339),
		UpdatedAt:             f.UpdatedAt.Format(time.RFC3339),
	}
}

// --- Normative tables ---

// normEntry holds age bracket thresholds: [ageMin, ageMax, excelente, bom, medio, abaixo]
type normEntry struct {
	ageMin, ageMax                       int
	excelente, bom, medio, abaixo int
}

// genderedNorms holds male and female normative tables
type genderedNorms struct {
	male   []normEntry
	female []normEntry
}

var abdominalNorms = genderedNorms{
	male: []normEntry{
		{20, 29, 42, 38, 33, 29},
		{30, 39, 36, 32, 27, 22},
		{40, 49, 31, 26, 22, 17},
		{50, 59, 26, 22, 18, 13},
		{60, 69, 20, 16, 12, 8},
	},
	female: []normEntry{
		{20, 29, 36, 32, 25, 21},
		{30, 39, 29, 25, 20, 15},
		{40, 49, 25, 20, 15, 9},
		{50, 59, 19, 14, 10, 6},
		{60, 69, 16, 12, 6, 4},
	},
}

var pushupNorms = genderedNorms{
	male: []normEntry{
		{20, 29, 36, 29, 22, 17},
		{30, 39, 30, 22, 17, 12},
		{40, 49, 25, 17, 13, 10},
		{50, 59, 21, 13, 10, 7},
		{60, 69, 18, 11, 8, 5},
	},
	female: []normEntry{
		{20, 29, 30, 21, 15, 10},
		{30, 39, 27, 20, 13, 8},
		{40, 49, 24, 15, 11, 5},
		{50, 59, 21, 11, 7, 2},
		{60, 69, 17, 12, 5, 2},
	},
}

// Plank norms are the same for both genders
var plankNorms = genderedNorms{
	male: []normEntry{
		{20, 29, 120, 90, 60, 30},
		{30, 39, 105, 75, 45, 25},
		{40, 49, 90, 60, 35, 20},
		{50, 59, 75, 50, 30, 15},
		{60, 69, 60, 40, 20, 10},
	},
	female: []normEntry{
		{20, 29, 120, 90, 60, 30},
		{30, 39, 105, 75, 45, 25},
		{40, 49, 90, 60, 35, 20},
		{50, 59, 75, 50, 30, 15},
		{60, 69, 60, 40, 20, 10},
	},
}

var burpeeNorms = genderedNorms{
	male: []normEntry{
		{20, 29, 45, 38, 30, 20},
		{30, 39, 40, 33, 25, 17},
		{40, 49, 35, 28, 20, 13},
		{50, 59, 30, 23, 16, 10},
		{60, 69, 25, 18, 12, 7},
	},
	female: []normEntry{
		{20, 29, 38, 30, 23, 15},
		{30, 39, 33, 25, 18, 12},
		{40, 49, 28, 20, 14, 8},
		{50, 59, 23, 16, 10, 5},
		{60, 69, 18, 12, 7, 3},
	},
}

var frtNorms = genderedNorms{
	male: []normEntry{
		{20, 29, 28, 23, 18, 13},
		{30, 39, 25, 20, 15, 10},
		{40, 49, 22, 17, 12, 8},
		{50, 59, 18, 14, 10, 6},
		{60, 69, 15, 11, 7, 4},
	},
	female: []normEntry{
		{20, 29, 22, 18, 14, 10},
		{30, 39, 19, 15, 11, 7},
		{40, 49, 16, 12, 8, 5},
		{50, 59, 13, 10, 6, 3},
		{60, 69, 10, 7, 4, 2},
	},
}

// classifyTest classifies a raw value against gendered normative tables
func classifyTest(norms genderedNorms, gender models.Gender, age, value int) string {
	table := norms.male
	if gender == models.GenderFemale {
		table = norms.female
	}

	// Find age bracket; clamp to nearest if out of range
	var entry *normEntry
	for i := range table {
		if age >= table[i].ageMin && age <= table[i].ageMax {
			entry = &table[i]
			break
		}
	}
	if entry == nil {
		// Age below lowest bracket: use first; above highest: use last
		if age < table[0].ageMin {
			entry = &table[0]
		} else {
			entry = &table[len(table)-1]
		}
	}

	switch {
	case value >= entry.excelente:
		return "excelente"
	case value >= entry.bom:
		return "bom"
	case value >= entry.medio:
		return "medio"
	case value >= entry.abaixo:
		return "abaixo"
	default:
		return "fraco"
	}
}

// levelToScore converts a classification level to a numeric score
func levelToScore(level string) int {
	switch level {
	case "excelente":
		return 100
	case "bom":
		return 75
	case "medio":
		return 50
	case "abaixo":
		return 25
	default: // "fraco"
		return 0
	}
}

// calculateOverallScore computes the average score and overall classification
func calculateOverallScore(f *models.FitnessTestResult) (int, string) {
	total := 0
	count := 0

	if f.AbdominalLevel != nil {
		total += levelToScore(*f.AbdominalLevel)
		count++
	}
	if f.PushupLevel != nil {
		total += levelToScore(*f.PushupLevel)
		count++
	}
	if f.PlankLevel != nil {
		total += levelToScore(*f.PlankLevel)
		count++
	}
	if f.BurpeeLevel != nil {
		total += levelToScore(*f.BurpeeLevel)
		count++
	}
	if f.FRTLevel != nil {
		total += levelToScore(*f.FRTLevel)
		count++
	}

	if count == 0 {
		return 0, "Sem dados"
	}

	score := total / count

	var classification string
	switch {
	case score >= 90:
		classification = "Excelente"
	case score >= 70:
		classification = "Bom"
	case score >= 50:
		classification = "Medio"
	case score >= 30:
		classification = "Abaixo da Media"
	default:
		classification = "Fraco"
	}

	return score, classification
}
