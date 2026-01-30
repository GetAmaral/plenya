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
	ErrAppointmentNotFound = errors.New("appointment not found")
	ErrAppointmentConflict = errors.New("appointment time slot already booked")
)

type AppointmentService struct {
	db *gorm.DB
}

func NewAppointmentService(db *gorm.DB) *AppointmentService {
	return &AppointmentService{db: db}
}

// Create cria uma nova consulta
func (s *AppointmentService) Create(userID uuid.UUID, req *dto.CreateAppointmentRequest) (*dto.AppointmentResponse, error) {
	// CRITICAL SECURITY: Get user's selected patient
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}

	// If no selected patient, cannot create appointment
	if user.SelectedPatientID == nil {
		return nil, errors.New("no patient selected - please select a patient first")
	}

	// Parse patient ID from request
	var patientID uuid.UUID
	if req.PatientID != "" {
		pid, err := uuid.Parse(req.PatientID)
		if err != nil {
			return nil, errors.New("invalid patient id")
		}
		// SECURITY: Validate that patientID matches selectedPatient
		if pid != *user.SelectedPatientID {
			return nil, errors.New("patient id does not match selected patient")
		}
		patientID = pid
	} else {
		// Auto-fill with selectedPatient
		patientID = *user.SelectedPatientID
	}

	doctorID, err := uuid.Parse(req.DoctorID)
	if err != nil {
		return nil, errors.New("invalid doctor id")
	}

	// Verificar se paciente existe
	var patient models.Patient
	if err := s.db.First(&patient, patientID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPatientNotFound
		}
		return nil, err
	}

	// Verificar se médico existe e é doctor
	var doctor models.User
	if err := s.db.First(&doctor, doctorID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("doctor not found")
		}
		return nil, err
	}
	if doctor.Role != models.RoleDoctor {
		return nil, errors.New("user is not a doctor")
	}

	// Parse scheduled date
	scheduledAt, err := time.Parse(time.RFC3339, req.ScheduledAt)
	if err != nil {
		return nil, errors.New("invalid scheduled date format, expected RFC3339")
	}

	// Verificar conflito de horário do médico
	endTime := scheduledAt.Add(time.Duration(req.DurationMinutes) * time.Minute)
	var conflict models.Appointment
	err = s.db.Where("doctor_id = ? AND status NOT IN (?, ?)", doctorID, models.AppointmentCancelled, models.AppointmentNoShow).
		Where("scheduled_at < ? AND (scheduled_at + INTERVAL '1 minute' * duration_minutes) > ?", endTime, scheduledAt).
		First(&conflict).Error

	if err == nil {
		return nil, ErrAppointmentConflict
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// Criar appointment
	appointment := models.Appointment{
		PatientID:       patientID,
		DoctorID:        doctorID,
		ScheduledAt:     scheduledAt,
		DurationMinutes: req.DurationMinutes,
		Type:            req.Type,
		Status:          models.AppointmentScheduled,
		Reason:          req.Reason,
		PatientNotes:    req.PatientNotes,
	}

	if err := s.db.Create(&appointment).Error; err != nil {
		return nil, err
	}

	return s.toDTO(&appointment), nil
}

// GetByID busca uma consulta por ID
func (s *AppointmentService) GetByID(appointmentID, userID uuid.UUID, userRole models.UserRole) (*dto.AppointmentResponse, error) {
	// CRITICAL SECURITY: Get user's selected patient
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}

	// If no selected patient, cannot access any appointment
	if user.SelectedPatientID == nil {
		return nil, ErrAppointmentNotFound
	}

	// ALWAYS filter by selectedPatient (all roles including admin)
	var appointment models.Appointment
	query := s.db.Where("id = ?", appointmentID).Where("patient_id = ?", *user.SelectedPatientID)

	if err := query.First(&appointment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAppointmentNotFound
		}
		return nil, err
	}

	return s.toDTO(&appointment), nil
}

// List lista consultas com filtros
func (s *AppointmentService) List(userID uuid.UUID, userRole models.UserRole, patientID, doctorID *uuid.UUID, status *models.AppointmentStatus, limit, offset int) ([]dto.AppointmentResponse, error) {
	// CRITICAL SECURITY: Get user's selected patient
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}

	// If no selected patient, return empty list (security measure)
	if user.SelectedPatientID == nil {
		return []dto.AppointmentResponse{}, nil
	}

	// ALWAYS filter by selectedPatient (all roles including admin)
	var appointments []models.Appointment
	query := s.db.Limit(limit).Offset(offset).Order("scheduled_at DESC")
	query = query.Where("patient_id = ?", *user.SelectedPatientID)

	// Filtro por status
	if status != nil {
		query = query.Where("status = ?", *status)
	}

	if err := query.Find(&appointments).Error; err != nil {
		return nil, err
	}

	result := make([]dto.AppointmentResponse, len(appointments))
	for i, a := range appointments {
		result[i] = *s.toDTO(&a)
	}

	return result, nil
}

// Update atualiza uma consulta
func (s *AppointmentService) Update(appointmentID, userID uuid.UUID, userRole models.UserRole, req *dto.UpdateAppointmentRequest) (*dto.AppointmentResponse, error) {
	var appointment models.Appointment
	query := s.db.Where("id = ?", appointmentID)

	if err := query.First(&appointment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAppointmentNotFound
		}
		return nil, err
	}

	// Permissões: médico pode editar sua consulta, paciente não pode mudar status
	if userRole == models.RolePatient {
		var patient models.Patient
		if err := s.db.Where("user_id = ?", userID).First(&patient).Error; err != nil {
			return nil, ErrUnauthorized
		}
		if appointment.PatientID != patient.ID {
			return nil, ErrUnauthorized
		}
		// Pacientes não podem mudar status
		req.Status = nil
	} else if userRole == models.RoleDoctor && appointment.DoctorID != userID {
		return nil, ErrUnauthorized
	}

	// Atualizar campos
	now := time.Now()
	if req.ScheduledAt != nil {
		scheduledAt, err := time.Parse(time.RFC3339, *req.ScheduledAt)
		if err != nil {
			return nil, errors.New("invalid scheduled date format, expected RFC3339")
		}
		appointment.ScheduledAt = scheduledAt
	}
	if req.Status != nil {
		appointment.Status = *req.Status
		// Atualizar timestamps baseado no status
		switch *req.Status {
		case models.AppointmentConfirmed:
			if appointment.ConfirmedAt == nil {
				appointment.ConfirmedAt = &now
			}
		case models.AppointmentCompleted:
			if appointment.CompletedAt == nil {
				appointment.CompletedAt = &now
			}
		case models.AppointmentCancelled:
			if appointment.CancelledAt == nil {
				appointment.CancelledAt = &now
			}
		}
	}
	if req.DoctorNotes != nil {
		appointment.DoctorNotes = req.DoctorNotes
	}
	if req.Diagnosis != nil {
		appointment.Diagnosis = req.Diagnosis
	}

	if err := s.db.Save(&appointment).Error; err != nil {
		return nil, err
	}

	return s.toDTO(&appointment), nil
}

// Cancel cancela uma consulta
func (s *AppointmentService) Cancel(appointmentID, userID uuid.UUID, userRole models.UserRole, req *dto.CancelAppointmentRequest) (*dto.AppointmentResponse, error) {
	var appointment models.Appointment
	if err := s.db.Where("id = ?", appointmentID).First(&appointment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAppointmentNotFound
		}
		return nil, err
	}

	// Verificar permissões
	if userRole == models.RolePatient {
		var patient models.Patient
		if err := s.db.Where("user_id = ?", userID).First(&patient).Error; err != nil {
			return nil, ErrUnauthorized
		}
		if appointment.PatientID != patient.ID {
			return nil, ErrUnauthorized
		}
	} else if userRole == models.RoleDoctor && appointment.DoctorID != userID {
		return nil, ErrUnauthorized
	}

	// Cancelar
	now := time.Now()
	appointment.Status = models.AppointmentCancelled
	appointment.CancelledAt = &now
	appointment.CancellationReason = &req.Reason

	if err := s.db.Save(&appointment).Error; err != nil {
		return nil, err
	}

	return s.toDTO(&appointment), nil
}

// Delete faz soft delete de uma consulta
func (s *AppointmentService) Delete(appointmentID uuid.UUID, userRole models.UserRole) error {
	if userRole != models.RoleAdmin {
		return ErrUnauthorized
	}

	result := s.db.Where("id = ?", appointmentID).Delete(&models.Appointment{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrAppointmentNotFound
	}

	return nil
}

// toDTO converte Appointment para AppointmentResponse
func (s *AppointmentService) toDTO(appointment *models.Appointment) *dto.AppointmentResponse {
	resp := &dto.AppointmentResponse{
		ID:              appointment.ID.String(),
		PatientID:       appointment.PatientID.String(),
		DoctorID:        appointment.DoctorID.String(),
		ScheduledAt:     appointment.ScheduledAt.Format(time.RFC3339),
		DurationMinutes: appointment.DurationMinutes,
		Type:            appointment.Type,
		Status:          appointment.Status,
		Reason:          appointment.Reason,
		PatientNotes:    appointment.PatientNotes,
		DoctorNotes:     appointment.DoctorNotes,
		Diagnosis:       appointment.Diagnosis,
		CancellationReason: appointment.CancellationReason,
		CreatedAt:       appointment.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       appointment.UpdatedAt.Format(time.RFC3339),
	}

	if appointment.AnamnesisID != nil {
		aid := appointment.AnamnesisID.String()
		resp.AnamnesisID = &aid
	}
	if appointment.ConfirmedAt != nil {
		ca := appointment.ConfirmedAt.Format(time.RFC3339)
		resp.ConfirmedAt = &ca
	}
	if appointment.CompletedAt != nil {
		ca := appointment.CompletedAt.Format(time.RFC3339)
		resp.CompletedAt = &ca
	}
	if appointment.CancelledAt != nil {
		ca := appointment.CancelledAt.Format(time.RFC3339)
		resp.CancelledAt = &ca
	}

	return resp
}
