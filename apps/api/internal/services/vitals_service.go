package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

// VitalsService — sinais vitais por consulta (escopados por patientID via path).
type VitalsService struct {
	db *gorm.DB
}

func NewVitalsService(db *gorm.DB) *VitalsService {
	return &VitalsService{db: db}
}

// ListByPatient retorna as aferições do paciente (mais recentes primeiro).
// Se appointmentID != nil, filtra pela consulta.
func (s *VitalsService) ListByPatient(patientID uuid.UUID, appointmentID *uuid.UUID, limit int) ([]dto.ConsultationVitalsResponse, error) {
	if limit <= 0 || limit > 100 {
		limit = 50
	}
	q := s.db.Preload("MeasuredByUser").Where("patient_id = ?", patientID)
	if appointmentID != nil {
		q = q.Where("appointment_id = ?", *appointmentID)
	}
	var rows []models.ConsultationVitals
	if err := q.Order("measured_at DESC").Limit(limit).Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]dto.ConsultationVitalsResponse, len(rows))
	for i := range rows {
		out[i] = *s.toDTO(&rows[i])
	}
	return out, nil
}

func (s *VitalsService) Create(patientID, measuredBy uuid.UUID, req *dto.CreateConsultationVitalsRequest) (*dto.ConsultationVitalsResponse, error) {
	var appointmentID *uuid.UUID
	if req.AppointmentID != nil && *req.AppointmentID != "" {
		aid, err := uuid.Parse(*req.AppointmentID)
		if err != nil {
			return nil, errors.New("invalid appointment id")
		}
		var appt models.Appointment
		if err := s.db.Select("id", "patient_id").First(&appt, aid).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrAppointmentNotFound
			}
			return nil, err
		}
		if appt.PatientID != patientID {
			return nil, errors.New("appointment does not belong to this patient")
		}
		appointmentID = &aid
	}

	measuredAt := time.Now()
	if req.MeasuredAt != nil && *req.MeasuredAt != "" {
		t, err := time.Parse(time.RFC3339, *req.MeasuredAt)
		if err != nil {
			return nil, errors.New("invalid measuredAt format, expected RFC3339")
		}
		measuredAt = t
	}

	v := models.ConsultationVitals{
		AppointmentID:      appointmentID,
		PatientID:          patientID,
		SystolicBP:         req.SystolicBP,
		DiastolicBP:        req.DiastolicBP,
		HeartRate:          req.HeartRate,
		RespRate:           req.RespRate,
		Temperature:        req.Temperature,
		SpO2:               req.SpO2,
		Weight:             req.Weight,
		Height:             req.Height,
		WaistCircumference: req.WaistCircumference,
		MeasuredByUserID:   measuredBy,
		MeasuredAt:         measuredAt,
	}
	if err := s.db.Create(&v).Error; err != nil {
		return nil, err
	}
	var loaded models.ConsultationVitals
	if err := s.db.Preload("MeasuredByUser").First(&loaded, v.ID).Error; err != nil {
		return nil, err
	}
	return s.toDTO(&loaded), nil
}

func (s *VitalsService) toDTO(v *models.ConsultationVitals) *dto.ConsultationVitalsResponse {
	r := &dto.ConsultationVitalsResponse{
		ID:                 v.ID.String(),
		PatientID:          v.PatientID.String(),
		SystolicBP:         v.SystolicBP,
		DiastolicBP:        v.DiastolicBP,
		HeartRate:          v.HeartRate,
		RespRate:           v.RespRate,
		Temperature:        v.Temperature,
		SpO2:               v.SpO2,
		Weight:             v.Weight,
		Height:             v.Height,
		WaistCircumference: v.WaistCircumference,
		BMI:                v.BMI,
		MeasuredByUserID:   v.MeasuredByUserID.String(),
		MeasuredAt:         v.MeasuredAt.Format(time.RFC3339),
		CreatedAt:          v.CreatedAt.Format(time.RFC3339),
	}
	if v.AppointmentID != nil {
		aid := v.AppointmentID.String()
		r.AppointmentID = &aid
	}
	if v.MeasuredByUser.ID != uuid.Nil {
		r.MeasuredByName = v.MeasuredByUser.Name
	}
	return r
}
