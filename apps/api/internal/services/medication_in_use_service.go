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

var ErrMedicationInUseNotFound = errors.New("medication in use not found")

// MedicationInUseService — medicações em uso do paciente (reconciliação).
type MedicationInUseService struct {
	db *gorm.DB
}

func NewMedicationInUseService(db *gorm.DB) *MedicationInUseService {
	return &MedicationInUseService{db: db}
}

func (s *MedicationInUseService) ListByPatient(patientID uuid.UUID, includeInactive bool) ([]dto.MedicationInUseResponse, error) {
	q := s.db.Preload("RecordedByUser").Where("patient_id = ?", patientID)
	if !includeInactive {
		q = q.Where("status = ?", models.MedInUseActive)
	}
	var rows []models.MedicationInUse
	if err := q.Order("status ASC").Order("medication_name ASC").Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]dto.MedicationInUseResponse, len(rows))
	for i := range rows {
		out[i] = *s.toDTO(&rows[i])
	}
	return out, nil
}

func (s *MedicationInUseService) Create(patientID, recordedBy uuid.UUID, req *dto.CreateMedicationInUseRequest) (*dto.MedicationInUseResponse, error) {
	source := models.MedSourcePatientReported
	if req.Source != "" {
		source = models.MedicationSource(req.Source)
	}
	m := models.MedicationInUse{
		PatientID:        patientID,
		MedicationName:   strings.TrimSpace(req.MedicationName),
		ActiveIngredient: strings.TrimSpace(req.ActiveIngredient),
		Dosage:           req.Dosage,
		Frequency:        req.Frequency,
		Route:            req.Route,
		Source:           source,
		Status:           models.MedInUseActive,
		Notes:            req.Notes,
		RecordedByUserID: recordedBy,
	}
	if req.SourcePrescriptionID != nil && *req.SourcePrescriptionID != "" {
		pid, err := uuid.Parse(*req.SourcePrescriptionID)
		if err != nil {
			return nil, errors.New("invalid sourcePrescriptionId")
		}
		m.SourcePrescriptionID = &pid
	}
	if req.StartDate != nil && *req.StartDate != "" {
		t, err := parseFlexibleDate(*req.StartDate)
		if err != nil {
			return nil, errors.New("invalid startDate")
		}
		m.StartDate = &t
	}
	if err := s.db.Create(&m).Error; err != nil {
		return nil, err
	}
	return s.loadDTO(m.ID)
}

func (s *MedicationInUseService) Update(medID, patientID uuid.UUID, req *dto.UpdateMedicationInUseRequest) (*dto.MedicationInUseResponse, error) {
	var m models.MedicationInUse
	if err := s.db.Where("id = ? AND patient_id = ?", medID, patientID).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMedicationInUseNotFound
		}
		return nil, err
	}
	if req.MedicationName != nil {
		m.MedicationName = strings.TrimSpace(*req.MedicationName)
	}
	if req.ActiveIngredient != nil {
		m.ActiveIngredient = strings.TrimSpace(*req.ActiveIngredient)
	}
	if req.Dosage != nil {
		m.Dosage = *req.Dosage
	}
	if req.Frequency != nil {
		m.Frequency = *req.Frequency
	}
	if req.Route != nil {
		m.Route = *req.Route
	}
	if req.Status != nil {
		m.Status = models.MedicationInUseStatus(*req.Status)
		if m.Status == models.MedInUseStopped && m.EndDate == nil && req.EndDate == nil {
			now := time.Now()
			m.EndDate = &now
		}
	}
	if req.EndDate != nil && *req.EndDate != "" {
		t, err := parseFlexibleDate(*req.EndDate)
		if err != nil {
			return nil, errors.New("invalid endDate")
		}
		m.EndDate = &t
	}
	if req.Notes != nil {
		m.Notes = req.Notes
	}
	if err := s.db.Save(&m).Error; err != nil {
		return nil, err
	}
	return s.loadDTO(m.ID)
}

func (s *MedicationInUseService) Delete(medID, patientID uuid.UUID) error {
	res := s.db.Where("id = ? AND patient_id = ?", medID, patientID).Delete(&models.MedicationInUse{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return ErrMedicationInUseNotFound
	}
	return nil
}

func (s *MedicationInUseService) loadDTO(id uuid.UUID) (*dto.MedicationInUseResponse, error) {
	var m models.MedicationInUse
	if err := s.db.Preload("RecordedByUser").First(&m, id).Error; err != nil {
		return nil, err
	}
	return s.toDTO(&m), nil
}

func (s *MedicationInUseService) toDTO(m *models.MedicationInUse) *dto.MedicationInUseResponse {
	r := &dto.MedicationInUseResponse{
		ID:               m.ID.String(),
		PatientID:        m.PatientID.String(),
		MedicationName:   m.MedicationName,
		ActiveIngredient: m.ActiveIngredient,
		Dosage:           m.Dosage,
		Frequency:        m.Frequency,
		Route:            m.Route,
		Source:           string(m.Source),
		Status:           string(m.Status),
		StartDate:        fmtDatePtr(m.StartDate),
		EndDate:          fmtDatePtr(m.EndDate),
		Notes:            m.Notes,
		RecordedByUserID: m.RecordedByUserID.String(),
		CreatedAt:        m.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        m.UpdatedAt.Format(time.RFC3339),
	}
	if m.SourcePrescriptionID != nil {
		pid := m.SourcePrescriptionID.String()
		r.SourcePrescriptionID = &pid
	}
	if m.ReconciledAppointmentID != nil {
		aid := m.ReconciledAppointmentID.String()
		r.ReconciledAppointmentID = &aid
	}
	if m.RecordedByUser.ID != uuid.Nil {
		r.RecordedByName = m.RecordedByUser.Name
	}
	return r
}
