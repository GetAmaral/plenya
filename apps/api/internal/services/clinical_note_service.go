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
	ErrClinicalNoteNotFound   = errors.New("clinical note not found")
	ErrClinicalNoteSigned     = errors.New("clinical note is signed and cannot be edited (create an amendment instead)")
	ErrClinicalNoteRestricted = errors.New("you don't have permission to view this clinical note")
)

// ClinicalNoteService — nota de evolução por consulta (SOAP/APSO).
// Espelha o padrão de AnamnesisService (selectedPatient, visibilidade por
// papel, toDTO), mas a nota é ancorada ao Appointment e imutável após assinada.
type ClinicalNoteService struct {
	db *gorm.DB
}

func NewClinicalNoteService(db *gorm.DB) *ClinicalNoteService {
	return &ClinicalNoteService{db: db}
}

// canView replica a regra de visibilidade da anamnese.
func (s *ClinicalNoteService) canView(note *models.ClinicalNote, userID uuid.UUID, userRole models.Role) bool {
	if userRole == models.RoleAdmin {
		return true
	}
	switch note.Visibility {
	case models.VisibilityAll:
		return true
	case models.VisibilityAuthorOnly:
		return note.AuthorID == userID
	case models.VisibilityMedicalOnly:
		return userRole == models.RoleDoctor
	case models.VisibilityPsychOnly:
		return false
	default:
		return false
	}
}

// resolvePatient retorna o selectedPatient do usuário, validando contra o
// patientID do request (mesma trava de segurança da anamnese).
func (s *ClinicalNoteService) resolvePatient(userID uuid.UUID, reqPatientID string) (uuid.UUID, error) {
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return uuid.Nil, err
	}
	if user.SelectedPatientID == nil {
		return uuid.Nil, errors.New("no patient selected - please select a patient first")
	}
	if reqPatientID != "" {
		pid, err := uuid.Parse(reqPatientID)
		if err != nil {
			return uuid.Nil, errors.New("invalid patient id")
		}
		if pid != *user.SelectedPatientID {
			return uuid.Nil, errors.New("patient id does not match selected patient")
		}
		return pid, nil
	}
	return *user.SelectedPatientID, nil
}

// Create cria uma nova nota de evolução (rascunho, ou já assinada se req.Sign).
func (s *ClinicalNoteService) Create(authorID uuid.UUID, req *dto.CreateClinicalNoteRequest) (*dto.ClinicalNoteResponse, error) {
	// O paciente da nota é determinado pela CONSULTA quando ancorada a ela (a
	// consulta é o contexto de autorização do médico que está documentando a
	// visita). Sem appointment (nota avulsa), cai no selectedPatient.
	var patientID uuid.UUID
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
		patientID = appt.PatientID
		appointmentID = &aid
	} else {
		pid, err := s.resolvePatient(authorID, req.PatientID)
		if err != nil {
			return nil, err
		}
		patientID = pid
	}

	visibility := models.VisibilityAll
	if req.Visibility != "" {
		visibility = models.AnamnesisVisibility(req.Visibility)
	}

	note := models.ClinicalNote{
		AppointmentID:       appointmentID,
		PatientID:           patientID,
		AuthorID:            authorID,
		ClinicalHistory:     req.ClinicalHistory,
		ClinicalHistoryHtml: sanitizeRichHTMLPtr(req.ClinicalHistoryHtml),
		Conduct:             req.Conduct,
		ConductHtml:         sanitizeRichHTMLPtr(req.ConductHtml),
		Visibility:          visibility,
		Status:              models.ClinicalNoteStatusDraft,
	}
	if req.Sign {
		now := time.Now()
		note.Status = models.ClinicalNoteStatusSigned
		note.SignedAt = &now
	}

	if err := s.db.Create(&note).Error; err != nil {
		return nil, err
	}
	return s.loadDTO(note.ID, authorID)
}

// GetByID busca uma nota por ID (escopada ao selectedPatient).
func (s *ClinicalNoteService) GetByID(noteID, userID uuid.UUID, userRole models.Role) (*dto.ClinicalNoteResponse, error) {
	note, err := s.fetchScoped(noteID, userID)
	if err != nil {
		return nil, err
	}
	if !s.canView(note, userID, userRole) {
		return nil, ErrClinicalNoteRestricted
	}
	return s.toDTO(note), nil
}

// GetByAppointment retorna a nota mais recente vinculada a uma consulta
// (usada pelo workspace ao abrir). Retorna ErrClinicalNoteNotFound se não houver.
func (s *ClinicalNoteService) GetByAppointment(appointmentID, userID uuid.UUID, userRole models.Role) (*dto.ClinicalNoteResponse, error) {
	// Centrado na consulta: o staff que abre o appointment vê a nota dele.
	var note models.ClinicalNote
	err := s.db.
		Preload("Author").
		Where("appointment_id = ?", appointmentID).
		Order("created_at DESC").
		First(&note).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrClinicalNoteNotFound
		}
		return nil, err
	}
	if !s.canView(&note, userID, userRole) {
		return nil, ErrClinicalNoteRestricted
	}
	return s.toDTO(&note), nil
}

// ListByPatient lista as notas do selectedPatient (mais recentes primeiro).
func (s *ClinicalNoteService) ListByPatient(userID uuid.UUID, userRole models.Role, limit, offset int) ([]dto.ClinicalNoteResponse, error) {
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}
	if user.SelectedPatientID == nil {
		return []dto.ClinicalNoteResponse{}, nil
	}
	var notes []models.ClinicalNote
	if err := s.db.
		Preload("Author").
		Where("patient_id = ?", *user.SelectedPatientID).
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&notes).Error; err != nil {
		return nil, err
	}
	out := make([]dto.ClinicalNoteResponse, 0, len(notes))
	for i := range notes {
		if !s.canView(&notes[i], userID, userRole) {
			continue
		}
		out = append(out, *s.toDTO(&notes[i]))
	}
	return out, nil
}

// Update atualiza uma nota — só permitido em rascunho (signed é imutável).
func (s *ClinicalNoteService) Update(noteID, authorID uuid.UUID, userRole models.Role, req *dto.UpdateClinicalNoteRequest) (*dto.ClinicalNoteResponse, error) {
	note, err := s.fetchEditable(noteID, authorID, userRole)
	if err != nil {
		return nil, err
	}
	if note.Status == models.ClinicalNoteStatusSigned {
		return nil, ErrClinicalNoteSigned
	}

	if req.ClinicalHistory != nil {
		note.ClinicalHistory = req.ClinicalHistory
	}
	if req.ClinicalHistoryHtml != nil {
		note.ClinicalHistoryHtml = sanitizeRichHTMLPtr(req.ClinicalHistoryHtml)
	}
	if req.Conduct != nil {
		note.Conduct = req.Conduct
	}
	if req.ConductHtml != nil {
		note.ConductHtml = sanitizeRichHTMLPtr(req.ConductHtml)
	}
	if req.Visibility != nil {
		note.Visibility = models.AnamnesisVisibility(*req.Visibility)
	}

	if err := s.db.Save(note).Error; err != nil {
		return nil, err
	}
	return s.loadDTO(note.ID, authorID)
}

// Sign assina a nota (draft -> signed). Idempotente: nota já assinada retorna como está.
func (s *ClinicalNoteService) Sign(noteID, authorID uuid.UUID, userRole models.Role) (*dto.ClinicalNoteResponse, error) {
	note, err := s.fetchEditable(noteID, authorID, userRole)
	if err != nil {
		return nil, err
	}
	if note.Status != models.ClinicalNoteStatusSigned {
		now := time.Now()
		note.Status = models.ClinicalNoteStatusSigned
		note.SignedAt = &now
		if err := s.db.Save(note).Error; err != nil {
			return nil, err
		}
	}
	return s.loadDTO(note.ID, authorID)
}

// Amend cria uma nova nota (rascunho) como adendo de uma nota assinada.
func (s *ClinicalNoteService) Amend(originalID, authorID uuid.UUID, userRole models.Role, req *dto.CreateClinicalNoteRequest) (*dto.ClinicalNoteResponse, error) {
	original, err := s.fetchScoped(originalID, authorID)
	if err != nil {
		return nil, err
	}
	if original.Status != models.ClinicalNoteStatusSigned {
		return nil, errors.New("only a signed note can be amended")
	}
	req.PatientID = original.PatientID.String()
	if original.AppointmentID != nil {
		aid := original.AppointmentID.String()
		req.AppointmentID = &aid
	}
	req.Sign = false
	resp, err := s.Create(authorID, req)
	if err != nil {
		return nil, err
	}
	// Marca o vínculo de adendo.
	if err := s.db.Model(&models.ClinicalNote{}).
		Where("id = ?", resp.ID).
		Update("amendment_of_id", original.ID).Error; err != nil {
		return nil, err
	}
	noteID, _ := uuid.Parse(resp.ID)
	return s.loadDTO(noteID, authorID)
}

// Delete faz soft delete (autor ou admin; rascunho).
func (s *ClinicalNoteService) Delete(noteID, authorID uuid.UUID, userRole models.Role) error {
	note, err := s.fetchEditable(noteID, authorID, userRole)
	if err != nil {
		return err
	}
	if note.Status == models.ClinicalNoteStatusSigned {
		return ErrClinicalNoteSigned
	}
	return s.db.Delete(&models.ClinicalNote{}, noteID).Error
}

// fetchScoped carrega a nota escopada ao selectedPatient do usuário.
func (s *ClinicalNoteService) fetchScoped(noteID, userID uuid.UUID) (*models.ClinicalNote, error) {
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}
	if user.SelectedPatientID == nil {
		return nil, ErrClinicalNoteNotFound
	}
	var note models.ClinicalNote
	if err := s.db.
		Preload("Author").
		Where("id = ?", noteID).
		Where("patient_id = ?", *user.SelectedPatientID).
		First(&note).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrClinicalNoteNotFound
		}
		return nil, err
	}
	return &note, nil
}

// fetchEditable carrega a nota garantindo que o usuário pode editá-la (autor ou admin).
func (s *ClinicalNoteService) fetchEditable(noteID, authorID uuid.UUID, userRole models.Role) (*models.ClinicalNote, error) {
	note, err := s.fetchScoped(noteID, authorID)
	if err != nil {
		return nil, err
	}
	if userRole != models.RoleAdmin && note.AuthorID != authorID {
		return nil, ErrUnauthorized
	}
	return note, nil
}

func (s *ClinicalNoteService) loadDTO(noteID, userID uuid.UUID) (*dto.ClinicalNoteResponse, error) {
	var note models.ClinicalNote
	if err := s.db.Preload("Author").First(&note, noteID).Error; err != nil {
		return nil, err
	}
	return s.toDTO(&note), nil
}

func (s *ClinicalNoteService) toDTO(note *models.ClinicalNote) *dto.ClinicalNoteResponse {
	resp := &dto.ClinicalNoteResponse{
		ID:                  note.ID.String(),
		PatientID:           note.PatientID.String(),
		AuthorID:            note.AuthorID.String(),
		ClinicalHistory:     note.ClinicalHistory,
		ClinicalHistoryHtml: note.ClinicalHistoryHtml,
		Conduct:             note.Conduct,
		ConductHtml:         note.ConductHtml,
		Status:              string(note.Status),
		Visibility:          string(note.Visibility),
		DisplayTitle:        note.GetTitle(),
		CreatedAt:           note.CreatedAt.Format(time.RFC3339),
		UpdatedAt:           note.UpdatedAt.Format(time.RFC3339),
	}
	if note.AppointmentID != nil {
		aid := note.AppointmentID.String()
		resp.AppointmentID = &aid
	}
	if note.AmendmentOfID != nil {
		amid := note.AmendmentOfID.String()
		resp.AmendmentOfID = &amid
	}
	if note.SignedAt != nil {
		sa := note.SignedAt.Format(time.RFC3339)
		resp.SignedAt = &sa
	}
	if note.Author.ID != uuid.Nil {
		roles := note.Author.GetRoles()
		primaryRole := ""
		if len(roles) > 0 {
			primaryRole = roles[0]
		}
		resp.Author = &dto.AuthorBrief{
			ID:    note.Author.ID.String(),
			Name:  note.Author.Name,
			Email: note.Author.Email,
			Role:  primaryRole,
		}
	}
	return resp
}
