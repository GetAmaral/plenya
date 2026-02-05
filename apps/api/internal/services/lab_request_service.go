package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
)

// LabRequestService handles business logic for lab requests
type LabRequestService struct {
	repo       *repository.LabRequestRepository
	db         *gorm.DB
	pdfService *PDFService
}

// NewLabRequestService creates a new lab request service
func NewLabRequestService(repo *repository.LabRequestRepository, db *gorm.DB) *LabRequestService {
	return &LabRequestService{
		repo:       repo,
		db:         db,
		pdfService: NewPDFService(),
	}
}

// CreateLabRequest creates a new lab request
func (s *LabRequestService) CreateLabRequest(userID uuid.UUID, req *models.LabRequest) error {
	// CRITICAL SECURITY: Get user's selected patient
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return err
	}

	// If no selected patient, cannot create lab request
	if user.SelectedPatientID == nil {
		return errors.New("no patient selected - please select a patient first")
	}

	// If patientID is provided, validate it matches selectedPatient
	if req.PatientID != uuid.Nil {
		if req.PatientID != *user.SelectedPatientID {
			return errors.New("patient id does not match selected patient")
		}
	} else {
		// Auto-fill with selectedPatient
		req.PatientID = *user.SelectedPatientID
	}

	// Auto-fill doctorID if not provided
	if req.DoctorID == nil {
		req.DoctorID = &userID
	}

	return s.repo.CreateLabRequest(req)
}

// GetLabRequestByID retrieves a lab request by ID
func (s *LabRequestService) GetLabRequestByID(id uuid.UUID) (*models.LabRequest, error) {
	return s.repo.GetLabRequestByID(id)
}

// GetLabRequestWithRelations retrieves a lab request with Patient and Doctor relationships
func (s *LabRequestService) GetLabRequestWithRelations(id uuid.UUID) (*models.LabRequest, error) {
	var req models.LabRequest
	if err := s.db.Preload("Patient").Preload("Doctor").First(&req, id).Error; err != nil {
		return nil, err
	}
	return &req, nil
}

// GetLabRequestsByPatientID retrieves all lab requests for a patient
func (s *LabRequestService) GetLabRequestsByPatientID(patientID uuid.UUID) ([]models.LabRequest, error) {
	return s.repo.GetLabRequestsByPatientID(patientID)
}

// GetLabRequestsByDate retrieves all lab requests for a specific date
func (s *LabRequestService) GetLabRequestsByDate(date time.Time) ([]models.LabRequest, error) {
	return s.repo.GetLabRequestsByDate(date)
}

// GetLabRequestsByDateRange retrieves lab requests within a date range
func (s *LabRequestService) GetLabRequestsByDateRange(startDate, endDate time.Time) ([]models.LabRequest, error) {
	return s.repo.GetLabRequestsByDateRange(startDate, endDate)
}

// GetAllLabRequests retrieves all lab requests with pagination
func (s *LabRequestService) GetAllLabRequests(limit, offset int) ([]models.LabRequest, int64, error) {
	return s.repo.GetAllLabRequests(limit, offset)
}

// UpdateLabRequest updates an existing lab request
func (s *LabRequestService) UpdateLabRequest(id uuid.UUID, req *models.LabRequest) error {
	existing, err := s.repo.GetLabRequestByID(id)
	if err != nil {
		return err
	}

	// Update fields
	existing.PatientID = req.PatientID
	existing.Date = req.Date
	existing.Exams = req.Exams
	existing.Notes = req.Notes
	existing.DoctorID = req.DoctorID
	existing.LabRequestTemplateID = req.LabRequestTemplateID
	existing.PdfURL = req.PdfURL

	// Campos de assinatura digital
	existing.SignedPDFPath = req.SignedPDFPath
	existing.SignedPDFHash = req.SignedPDFHash
	existing.QRCodeData = req.QRCodeData
	existing.SignedAt = req.SignedAt

	return s.repo.UpdateLabRequest(existing)
}

// DeleteLabRequest soft deletes a lab request
func (s *LabRequestService) DeleteLabRequest(id uuid.UUID) error {
	return s.repo.DeleteLabRequest(id)
}

// GenerateLabRequestPDF generates a PDF for the lab request
func (s *LabRequestService) GenerateLabRequestPDF(id uuid.UUID) (string, error) {
	// Get lab request with all relationships (Patient, Doctor)
	req, err := s.repo.GetLabRequestByID(id)
	if err != nil {
		return "", err
	}

	// CRÍTICO: Verificar se o médico tem certificado ativo ANTES de gerar o PDF
	// Isso determina se o PDF terá seção de assinatura digital ou manual
	if req.Doctor != nil && req.Doctor.CertificateActive {
		// Médico tem certificado ativo - marcar para assinatura digital
		now := time.Now()
		req.SignedAt = &now

		// Copiar serial do certificado
		if req.Doctor.CertificateSerial != nil {
			req.CertificateSerial = req.Doctor.CertificateSerial
		}
	} else {
		// Garantir que SignedAt está nil para PDF sem assinatura digital
		req.SignedAt = nil
		req.CertificateSerial = nil
	}

	// Generate PDF (agora com SignedAt correto)
	pdfURL, err := s.pdfService.GenerateLabRequestPDF(req)
	if err != nil {
		return "", err
	}

	// Atualizar lab request no banco com os metadados
	if err := s.repo.UpdateLabRequest(req); err != nil {
		return "", err
	}

	return pdfURL, nil
}
