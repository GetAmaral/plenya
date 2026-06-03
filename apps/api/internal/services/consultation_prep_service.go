package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

// ConsultationPrepService gerencia a preparação pré-consulta do paciente (formulário curado
// pós-agendamento). Espelha o fluxo do Escore Light, porém vinculado ao Patient.
type ConsultationPrepService struct {
	db *gorm.DB
}

func NewConsultationPrepService(db *gorm.DB) *ConsultationPrepService {
	return &ConsultationPrepService{db: db}
}

// Get retorna a preparação do paciente para a consulta indicada (ou a sem-consulta quando
// appointmentID é nil). Retorna nil quando não existe.
func (s *ConsultationPrepService) Get(patientID uuid.UUID, appointmentID *uuid.UUID) (*models.ConsultationPrep, error) {
	return s.find(s.db, patientID, appointmentID)
}

func (s *ConsultationPrepService) find(tx *gorm.DB, patientID uuid.UUID, appointmentID *uuid.UUID) (*models.ConsultationPrep, error) {
	q := tx.Preload("Responses").Where("patient_id = ?", patientID)
	if appointmentID != nil {
		q = q.Where("appointment_id = ?", *appointmentID)
	} else {
		q = q.Where("appointment_id IS NULL")
	}
	var prep models.ConsultationPrep
	err := q.Order("created_at DESC").First(&prep).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &prep, nil
}

// Submit faz upsert da preparação + respostas do paciente, numa transação. Quando req.Submit
// é true, finaliza (status submitted + SubmittedAt). As respostas são substituídas integralmente.
func (s *ConsultationPrepService) Submit(patientID uuid.UUID, req dto.SubmitPrepRequest) (*models.ConsultationPrep, error) {
	var apptID *uuid.UUID
	if req.AppointmentID != nil && *req.AppointmentID != "" {
		id, err := uuid.Parse(*req.AppointmentID)
		if err != nil {
			return nil, errors.New("appointmentId inválido")
		}
		apptID = &id
	}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		existing, err := s.find(tx, patientID, apptID)
		if err != nil {
			return err
		}
		now := time.Now()
		if existing == nil {
			existing = &models.ConsultationPrep{
				PatientID:     patientID,
				AppointmentID: apptID,
				Status:        models.ConsultationPrepDraft,
			}
			if err := tx.Create(existing).Error; err != nil {
				return err
			}
		}

		existing.ChiefComplaint = req.ChiefComplaint
		if req.ConsentVersion != nil {
			existing.ConsentVersion = req.ConsentVersion
			existing.ConsentTimestamp = &now
		}
		if req.Submit {
			existing.Status = models.ConsultationPrepSubmitted
			existing.SubmittedAt = &now
		}
		if err := tx.Save(existing).Error; err != nil {
			return err
		}

		// Substitui as respostas integralmente.
		if err := tx.Where("prep_id = ?", existing.ID).Delete(&models.ConsultationPrepResponse{}).Error; err != nil {
			return err
		}
		for _, r := range req.Responses {
			if r.NumericValue == nil && r.SelectedLevel == nil && (r.TextValue == nil || *r.TextValue == "") {
				continue
			}
			sid, err := uuid.Parse(r.ScoreItemID)
			if err != nil {
				return errors.New("scoreItemId inválido")
			}
			resp := models.ConsultationPrepResponse{
				PrepID:        existing.ID,
				ScoreItemID:   sid,
				NumericValue:  r.NumericValue,
				SelectedLevel: r.SelectedLevel,
				TextValue:     r.TextValue,
			}
			if err := tx.Create(&resp).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return s.Get(patientID, apptID)
}

// ToConsultationPrepView mapeia o model para o DTO de leitura.
func ToConsultationPrepView(p *models.ConsultationPrep) dto.ConsultationPrepView {
	v := dto.ConsultationPrepView{
		ID:             p.ID.String(),
		PatientID:      p.PatientID.String(),
		Status:         p.Status,
		ChiefComplaint: p.ChiefComplaint,
		CreatedAt:      p.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      p.UpdatedAt.Format(time.RFC3339),
		Responses:      make([]dto.PrepResponseView, 0, len(p.Responses)),
	}
	if p.AppointmentID != nil {
		s := p.AppointmentID.String()
		v.AppointmentID = &s
	}
	if p.SubmittedAt != nil {
		s := p.SubmittedAt.Format(time.RFC3339)
		v.SubmittedAt = &s
	}
	for _, r := range p.Responses {
		v.Responses = append(v.Responses, dto.PrepResponseView{
			ScoreItemID:   r.ScoreItemID.String(),
			NumericValue:  r.NumericValue,
			SelectedLevel: r.SelectedLevel,
			TextValue:     r.TextValue,
		})
	}
	return v
}
