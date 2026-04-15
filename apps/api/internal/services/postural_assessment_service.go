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
	ErrPosturalAssessmentNotFound = errors.New("postural assessment not found")
)

type PosturalAssessmentService struct {
	db *gorm.DB
}

func NewPosturalAssessmentService(db *gorm.DB) *PosturalAssessmentService {
	return &PosturalAssessmentService{db: db}
}

// Create cria uma avaliacao postural com medicoes e calcula score automaticamente
func (s *PosturalAssessmentService) Create(userID uuid.UUID, req *dto.CreatePosturalAssessmentRequest) (*dto.PosturalAssessmentResponse, error) {
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

	assessment := models.PosturalAssessment{
		PatientID:            patientID,
		CreatedByID:          userID,
		AssessmentDate:       assessmentDate,
		ViewType:             req.ViewType,
		ShoulderDeviation:    req.ShoulderDeviation,
		HipDeviation:         req.HipDeviation,
		HeadLateralDeviation: req.HeadLateralDeviation,
		FHP:                  req.FHP,
		ThoracicKyphosis:     req.ThoracicKyphosis,
		LumbarLordosis:       req.LumbarLordosis,
		KneeAngle:            req.KneeAngle,
		PhotoURL:             req.PhotoURL,
		Notes:                req.Notes,
	}

	if req.PhysicalAssessmentID != nil {
		paID, err := uuid.Parse(*req.PhysicalAssessmentID)
		if err != nil {
			return nil, errors.New("invalid physical assessment id")
		}
		assessment.PhysicalAssessmentID = &paID
	}

	computePosturalScore(&assessment)

	if err := s.db.Create(&assessment).Error; err != nil {
		return nil, err
	}

	return s.toDTO(&assessment), nil
}

// GetByID busca uma avaliacao postural por ID
func (s *PosturalAssessmentService) GetByID(id, userID uuid.UUID) (*dto.PosturalAssessmentResponse, error) {
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}
	if user.SelectedPatientID == nil {
		return nil, ErrPosturalAssessmentNotFound
	}

	var assessment models.PosturalAssessment
	if err := s.db.Where("id = ? AND patient_id = ?", id, *user.SelectedPatientID).
		First(&assessment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPosturalAssessmentNotFound
		}
		return nil, err
	}

	return s.toDTO(&assessment), nil
}

// List lista avaliacoes posturais do paciente selecionado
func (s *PosturalAssessmentService) List(userID uuid.UUID, limit, offset int) ([]dto.PosturalAssessmentResponse, error) {
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}
	if user.SelectedPatientID == nil {
		return []dto.PosturalAssessmentResponse{}, nil
	}

	var assessments []models.PosturalAssessment
	if err := s.db.Where("patient_id = ?", *user.SelectedPatientID).
		Order("assessment_date DESC").
		Limit(limit).Offset(offset).
		Find(&assessments).Error; err != nil {
		return nil, err
	}

	result := make([]dto.PosturalAssessmentResponse, len(assessments))
	for i, a := range assessments {
		result[i] = *s.toDTO(&a)
	}
	return result, nil
}

// Delete faz soft delete de uma avaliacao postural
func (s *PosturalAssessmentService) Delete(id uuid.UUID, userRole models.Role) error {
	if userRole != models.RoleAdmin {
		return ErrUnauthorized
	}
	result := s.db.Where("id = ?", id).Delete(&models.PosturalAssessment{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrPosturalAssessmentNotFound
	}
	return nil
}

// --- Scoring ---

// computePosturalScore calcula o score postural baseado nas medicoes
func computePosturalScore(pa *models.PosturalAssessment) {
	totalPenalty := 0
	severeCount := 0

	// Standard deviations: shoulder, hip, head lateral, knee (thresholds: 2, 5, 10)
	standardMeasurements := []*float64{
		pa.ShoulderDeviation,
		pa.HipDeviation,
		pa.HeadLateralDeviation,
		pa.KneeAngle,
	}
	for _, m := range standardMeasurements {
		p, severe := penaltyStandard(m, 2, 5, 10)
		totalPenalty += p
		if severe {
			severeCount++
		}
	}

	// FHP: thresholds 3, 8, 15
	{
		p, severe := penaltyStandard(pa.FHP, 3, 8, 15)
		totalPenalty += p
		if severe {
			severeCount++
		}
	}

	// Thoracic Kyphosis: thresholds 5, 15, 30
	{
		p, severe := penaltyStandard(pa.ThoracicKyphosis, 5, 15, 30)
		totalPenalty += p
		if severe {
			severeCount++
		}
	}

	// Lumbar Lordosis: thresholds 5, 15, 30
	{
		p, severe := penaltyStandard(pa.LumbarLordosis, 5, 15, 30)
		totalPenalty += p
		if severe {
			severeCount++
		}
	}

	score := 100 - totalPenalty
	if score < 0 {
		score = 0
	}

	pa.PosturalScore = score
	pa.SevereDeviations = severeCount

	switch {
	case score >= 90:
		pa.PosturalClassification = "Excelente"
	case score >= 70:
		pa.PosturalClassification = "Boa"
	case score >= 50:
		pa.PosturalClassification = "Regular"
	default:
		pa.PosturalClassification = "Necessita Atencao"
	}
}

// penaltyStandard returns (penalty, isSevere) for a measurement given three thresholds
func penaltyStandard(value *float64, t1, t2, t3 float64) (int, bool) {
	if value == nil {
		return 0, false
	}
	abs := math.Abs(*value)
	switch {
	case abs <= t1:
		return 0, false
	case abs <= t2:
		return 5, false
	case abs <= t3:
		return 15, false
	default:
		return 25, true
	}
}

// --- DTO ---

func (s *PosturalAssessmentService) toDTO(pa *models.PosturalAssessment) *dto.PosturalAssessmentResponse {
	resp := &dto.PosturalAssessmentResponse{
		ID:                     pa.ID.String(),
		PatientID:              pa.PatientID.String(),
		CreatedByID:            pa.CreatedByID.String(),
		AssessmentDate:         pa.AssessmentDate.Format("2006-01-02"),
		ViewType:               pa.ViewType,
		ShoulderDeviation:      pa.ShoulderDeviation,
		HipDeviation:           pa.HipDeviation,
		HeadLateralDeviation:   pa.HeadLateralDeviation,
		FHP:                    pa.FHP,
		ThoracicKyphosis:       pa.ThoracicKyphosis,
		LumbarLordosis:         pa.LumbarLordosis,
		KneeAngle:              pa.KneeAngle,
		PhotoURL:               pa.PhotoURL,
		PosturalScore:          pa.PosturalScore,
		PosturalClassification: pa.PosturalClassification,
		SevereDeviations:       pa.SevereDeviations,
		Notes:                  pa.Notes,
		CreatedAt:              pa.CreatedAt.Format(time.RFC3339),
		UpdatedAt:              pa.UpdatedAt.Format(time.RFC3339),
	}

	if pa.PhysicalAssessmentID != nil {
		paID := pa.PhysicalAssessmentID.String()
		resp.PhysicalAssessmentID = &paID
	}

	return resp
}
