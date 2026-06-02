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

var ErrProblemNotFound = errors.New("problem not found")

// parseFlexibleDate aceita YYYY-MM-DD ou RFC3339. Compartilhado no pacote services.
func parseFlexibleDate(s string) (time.Time, error) {
	s = strings.TrimSpace(s)
	if t, err := time.Parse("2006-01-02", s); err == nil {
		return t, nil
	}
	return time.Parse(time.RFC3339, s)
}

func fmtDatePtr(t *time.Time) *string {
	if t == nil {
		return nil
	}
	s := t.Format("2006-01-02")
	return &s
}

// ProblemService — lista de problemas do paciente (escopada por patientID).
type ProblemService struct {
	db *gorm.DB
}

func NewProblemService(db *gorm.DB) *ProblemService {
	return &ProblemService{db: db}
}

func (s *ProblemService) ListByPatient(patientID uuid.UUID, includeResolved bool) ([]dto.PatientProblemResponse, error) {
	q := s.db.Preload("RecordedByUser").Where("patient_id = ?", patientID)
	if !includeResolved {
		q = q.Where("status = ?", models.ProblemStatusActive)
	}
	var rows []models.PatientProblem
	if err := q.Order("status ASC").Order("created_at DESC").Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]dto.PatientProblemResponse, len(rows))
	for i := range rows {
		out[i] = *s.toDTO(&rows[i])
	}
	return out, nil
}

func (s *ProblemService) Create(patientID, recordedBy uuid.UUID, req *dto.CreatePatientProblemRequest) (*dto.PatientProblemResponse, error) {
	p := models.PatientProblem{
		PatientID:        patientID,
		Description:      strings.TrimSpace(req.Description),
		CIDCode:          req.CIDCode,
		CIDVersion:       "CID-10",
		Status:           models.ProblemStatusActive,
		Notes:            req.Notes,
		RecordedByUserID: recordedBy,
	}
	if req.CIDVersion != "" {
		p.CIDVersion = req.CIDVersion
	}
	if req.OnsetDate != nil && *req.OnsetDate != "" {
		t, err := parseFlexibleDate(*req.OnsetDate)
		if err != nil {
			return nil, errors.New("invalid onsetDate")
		}
		p.OnsetDate = &t
	}
	if req.OnsetAppointmentID != nil && *req.OnsetAppointmentID != "" {
		aid, err := uuid.Parse(*req.OnsetAppointmentID)
		if err != nil {
			return nil, errors.New("invalid onsetAppointmentId")
		}
		p.OnsetAppointmentID = &aid
	}
	if err := s.db.Create(&p).Error; err != nil {
		return nil, err
	}
	return s.loadDTO(p.ID)
}

func (s *ProblemService) Update(problemID, patientID uuid.UUID, req *dto.UpdatePatientProblemRequest) (*dto.PatientProblemResponse, error) {
	var p models.PatientProblem
	if err := s.db.Where("id = ? AND patient_id = ?", problemID, patientID).First(&p).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrProblemNotFound
		}
		return nil, err
	}
	if req.Description != nil {
		p.Description = strings.TrimSpace(*req.Description)
	}
	if req.CIDCode != nil {
		if *req.CIDCode == "" {
			p.CIDCode = nil
		} else {
			p.CIDCode = req.CIDCode
		}
	}
	if req.Status != nil {
		p.Status = models.ProblemStatus(*req.Status)
		// Ao resolver sem data explícita, carimba hoje.
		if p.Status == models.ProblemStatusResolved && p.ResolvedDate == nil && req.ResolvedDate == nil {
			now := time.Now()
			p.ResolvedDate = &now
		}
	}
	if req.OnsetDate != nil && *req.OnsetDate != "" {
		t, err := parseFlexibleDate(*req.OnsetDate)
		if err != nil {
			return nil, errors.New("invalid onsetDate")
		}
		p.OnsetDate = &t
	}
	if req.ResolvedDate != nil && *req.ResolvedDate != "" {
		t, err := parseFlexibleDate(*req.ResolvedDate)
		if err != nil {
			return nil, errors.New("invalid resolvedDate")
		}
		p.ResolvedDate = &t
	}
	if req.Notes != nil {
		p.Notes = req.Notes
	}
	if err := s.db.Save(&p).Error; err != nil {
		return nil, err
	}
	return s.loadDTO(p.ID)
}

func (s *ProblemService) Delete(problemID, patientID uuid.UUID) error {
	res := s.db.Where("id = ? AND patient_id = ?", problemID, patientID).Delete(&models.PatientProblem{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return ErrProblemNotFound
	}
	return nil
}

func (s *ProblemService) loadDTO(id uuid.UUID) (*dto.PatientProblemResponse, error) {
	var p models.PatientProblem
	if err := s.db.Preload("RecordedByUser").First(&p, id).Error; err != nil {
		return nil, err
	}
	return s.toDTO(&p), nil
}

func (s *ProblemService) toDTO(p *models.PatientProblem) *dto.PatientProblemResponse {
	r := &dto.PatientProblemResponse{
		ID:               p.ID.String(),
		PatientID:        p.PatientID.String(),
		Description:      p.Description,
		CIDCode:          p.CIDCode,
		CIDVersion:       p.CIDVersion,
		Status:           string(p.Status),
		OnsetDate:        fmtDatePtr(p.OnsetDate),
		ResolvedDate:     fmtDatePtr(p.ResolvedDate),
		Notes:            p.Notes,
		RecordedByUserID: p.RecordedByUserID.String(),
		CreatedAt:        p.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        p.UpdatedAt.Format(time.RFC3339),
	}
	if p.OnsetAppointmentID != nil {
		aid := p.OnsetAppointmentID.String()
		r.OnsetAppointmentID = &aid
	}
	if p.RecordedByUser.ID != uuid.Nil {
		r.RecordedByName = p.RecordedByUser.Name
	}
	return r
}
