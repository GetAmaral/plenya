package services

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

var ErrAllergyNotFound = errors.New("allergy not found")

// AllergyService — alergias do paciente (escopadas por patientID via path).
type AllergyService struct {
	db *gorm.DB
}

func NewAllergyService(db *gorm.DB) *AllergyService {
	return &AllergyService{db: db}
}

// ListByPatient retorna as alergias do paciente. Ativas primeiro, depois por
// gravidade (anafilaxia → leve) e mais recentes.
func (s *AllergyService) ListByPatient(patientID uuid.UUID, includeInactive bool) ([]dto.PatientAllergyResponse, error) {
	q := s.db.Preload("RecordedByUser").Where("patient_id = ?", patientID)
	if !includeInactive {
		q = q.Where("status = ?", models.AllergyStatusActive)
	}
	var rows []models.PatientAllergy
	if err := q.
		Order("status ASC").
		Order("CASE severity WHEN 'anaphylaxis' THEN 0 WHEN 'severe' THEN 1 WHEN 'moderate' THEN 2 ELSE 3 END").
		Order("created_at DESC").
		Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]dto.PatientAllergyResponse, len(rows))
	for i := range rows {
		out[i] = *s.toDTO(&rows[i])
	}
	return out, nil
}

func (s *AllergyService) Create(patientID, recordedBy uuid.UUID, req *dto.CreatePatientAllergyRequest) (*dto.PatientAllergyResponse, error) {
	if !req.NoKnownAllergies && strings.TrimSpace(req.Substance) == "" {
		return nil, errors.New("substance is required unless noKnownAllergies is set")
	}
	substanceType := models.AllergyTypeDrug
	if req.SubstanceType != "" {
		substanceType = models.AllergySubstanceType(req.SubstanceType)
	}
	severity := models.AllergySeverityModerate
	if req.Severity != "" {
		severity = models.AllergySeverity(req.Severity)
	}
	a := models.PatientAllergy{
		PatientID:        patientID,
		Substance:        strings.TrimSpace(req.Substance),
		SubstanceType:    substanceType,
		Reaction:         req.Reaction,
		Severity:         severity,
		Status:           models.AllergyStatusActive,
		NoKnownAllergies: req.NoKnownAllergies,
		Notes:            req.Notes,
		RecordedByUserID: recordedBy,
	}
	if err := s.db.Create(&a).Error; err != nil {
		return nil, err
	}
	return s.loadDTO(a.ID)
}

func (s *AllergyService) Update(allergyID, patientID uuid.UUID, req *dto.UpdatePatientAllergyRequest) (*dto.PatientAllergyResponse, error) {
	var a models.PatientAllergy
	if err := s.db.Where("id = ? AND patient_id = ?", allergyID, patientID).First(&a).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAllergyNotFound
		}
		return nil, err
	}
	if req.Substance != nil {
		a.Substance = strings.TrimSpace(*req.Substance)
	}
	if req.SubstanceType != nil {
		a.SubstanceType = models.AllergySubstanceType(*req.SubstanceType)
	}
	if req.Reaction != nil {
		a.Reaction = req.Reaction
	}
	if req.Severity != nil {
		a.Severity = models.AllergySeverity(*req.Severity)
	}
	if req.Status != nil {
		a.Status = models.AllergyStatus(*req.Status)
	}
	if req.Notes != nil {
		a.Notes = req.Notes
	}
	if err := s.db.Save(&a).Error; err != nil {
		return nil, err
	}
	return s.loadDTO(a.ID)
}

func (s *AllergyService) Delete(allergyID, patientID uuid.UUID) error {
	res := s.db.Where("id = ? AND patient_id = ?", allergyID, patientID).Delete(&models.PatientAllergy{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return ErrAllergyNotFound
	}
	return nil
}

func (s *AllergyService) loadDTO(id uuid.UUID) (*dto.PatientAllergyResponse, error) {
	var a models.PatientAllergy
	if err := s.db.Preload("RecordedByUser").First(&a, id).Error; err != nil {
		return nil, err
	}
	return s.toDTO(&a), nil
}

func (s *AllergyService) toDTO(a *models.PatientAllergy) *dto.PatientAllergyResponse {
	r := &dto.PatientAllergyResponse{
		ID:               a.ID.String(),
		PatientID:        a.PatientID.String(),
		Substance:        a.Substance,
		SubstanceType:    string(a.SubstanceType),
		Reaction:         a.Reaction,
		Severity:         string(a.Severity),
		Status:           string(a.Status),
		NoKnownAllergies: a.NoKnownAllergies,
		Notes:            a.Notes,
		RecordedByUserID: a.RecordedByUserID.String(),
		CreatedAt:        a.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        a.UpdatedAt.Format(time.RFC3339),
	}
	if a.RecordedByUser.ID != uuid.Nil {
		r.RecordedByName = a.RecordedByUser.Name
	}
	return r
}
