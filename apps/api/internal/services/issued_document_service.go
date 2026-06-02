package services

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

var (
	ErrIssuedDocumentNotFound = errors.New("issued document not found")
	ErrIssuedDocAlreadySigned = errors.New("issued document already signed (immutable)")
	ErrCIDConsentRequired     = errors.New("incluir CID exige consentimento do paciente")
	ErrDoctorMissingCRM       = errors.New("médico sem CRM cadastrado")
)

// IssuedDocumentService — emite/assina/publica documentos clínicos (atestado/declaração/laudo).
// Recursos path-scoped por paciente (patientID vem da rota, não do selectedPatient).
type IssuedDocumentService struct {
	db          *gorm.DB
	pdfService  *DocumentPDFService
	signature   *SignatureService
	documents   *PatientDocumentsService
	uploadsRoot string
}

func NewIssuedDocumentService(
	db *gorm.DB,
	pdfService *DocumentPDFService,
	signature *SignatureService,
	documents *PatientDocumentsService,
	uploadsRoot string,
) *IssuedDocumentService {
	return &IssuedDocumentService{
		db:          db,
		pdfService:  pdfService,
		signature:   signature,
		documents:   documents,
		uploadsRoot: uploadsRoot,
	}
}

func (s *IssuedDocumentService) validationURL(id uuid.UUID) string {
	return fmt.Sprintf("https://plenya.com.br/documentos/validar/%s", id)
}

// Create cria um documento em rascunho (draft).
func (s *IssuedDocumentService) Create(patientID, userID uuid.UUID, req *dto.CreateIssuedDocumentRequest) (*dto.IssuedDocumentResponse, error) {
	var patient models.Patient
	if err := s.db.First(&patient, patientID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPatientNotFound
		}
		return nil, err
	}

	if req.IncludesCid && !req.CidConsent {
		return nil, ErrCIDConsentRequired
	}

	var appointmentID *uuid.UUID
	if req.AppointmentID != nil && *req.AppointmentID != "" {
		parsed, err := uuid.Parse(*req.AppointmentID)
		if err != nil {
			return nil, errors.New("invalid appointment id")
		}
		appointmentID = &parsed
	}

	doc := &models.IssuedDocument{
		PatientID:      patientID,
		AppointmentID:  appointmentID,
		DoctorID:       userID,
		Type:           models.IssuedDocumentType(req.Type),
		Title:          req.Title,
		Body:           req.Body,
		Purpose:        req.Purpose,
		DaysOff:        req.DaysOff,
		IncludesCID:    req.IncludesCid,
		CIDCode:        req.CidCode,
		CIDConsent:     req.CidConsent,
		Status:         models.IssuedDocDraft,
		IssuedByUserID: userID,
	}
	if err := s.db.Create(doc).Error; err != nil {
		return nil, err
	}
	return s.GetByID(doc.ID)
}

// GetByID retorna o documento (com médico).
func (s *IssuedDocumentService) GetByID(id uuid.UUID) (*dto.IssuedDocumentResponse, error) {
	var doc models.IssuedDocument
	if err := s.db.Preload("Doctor").First(&doc, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrIssuedDocumentNotFound
		}
		return nil, err
	}
	return s.toDTO(&doc), nil
}

// ListByPatient lista documentos emitidos do paciente (mais recentes primeiro).
func (s *IssuedDocumentService) ListByPatient(patientID uuid.UUID) ([]*dto.IssuedDocumentResponse, error) {
	var docs []models.IssuedDocument
	if err := s.db.Preload("Doctor").
		Where("patient_id = ?", patientID).
		Order("issued_at DESC").
		Find(&docs).Error; err != nil {
		return nil, err
	}
	out := make([]*dto.IssuedDocumentResponse, len(docs))
	for i := range docs {
		out[i] = s.toDTO(&docs[i])
	}
	return out, nil
}

// Sign gera o PDF, assina (ICP-Brasil quando há cert; senão assinatura manual),
// publica como PatientDocument (portal) e grava os metadados de assinatura.
func (s *IssuedDocumentService) Sign(id, userID uuid.UUID) (*dto.IssuedDocumentResponse, error) {
	var doc models.IssuedDocument
	if err := s.db.First(&doc, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrIssuedDocumentNotFound
		}
		return nil, err
	}
	if doc.Status == models.IssuedDocSigned {
		return nil, ErrIssuedDocAlreadySigned
	}
	if doc.IncludesCID && !doc.CIDConsent {
		return nil, ErrCIDConsentRequired
	}

	var patient models.Patient
	if err := s.db.First(&patient, doc.PatientID).Error; err != nil {
		return nil, err
	}
	var doctor models.User
	if err := s.db.First(&doctor, doc.DoctorID).Error; err != nil {
		return nil, err
	}
	if doctor.CRM == nil || *doctor.CRM == "" {
		return nil, ErrDoctorMissingCRM
	}

	validationURL := s.validationURL(doc.ID)
	hasDigital := doctor.CertificateActive

	// 1. Gerar PDF (bloco de assinatura conforme disponibilidade do cert)
	pdfBytes, err := s.pdfService.BuildDocumentPDF(&doc, &patient, &doctor, hasDigital, validationURL)
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar PDF: %w", err)
	}

	var finalPDF []byte
	var hash string
	var certSerial *string

	if hasDigital {
		reason := "Assinatura Digital de Documento Médico"
		signed, sigHash, sErr := s.signature.SignDocumentPDF(pdfBytes, doc.DoctorID, reason, "Plenya EMR - Documentos Clínicos")
		if sErr != nil {
			// Degradação graciosa (padrão lab_request): regenera sem assinatura digital.
			hasDigital = false
			unsigned, bErr := s.pdfService.BuildDocumentPDF(&doc, &patient, &doctor, false, validationURL)
			if bErr != nil {
				return nil, fmt.Errorf("erro ao regenerar PDF: %w", bErr)
			}
			finalPDF = unsigned
			sum := sha256.Sum256(unsigned)
			hash = hex.EncodeToString(sum[:])
		} else {
			finalPDF = signed
			hash = sigHash
			certSerial = doctor.CertificateSerial
		}
	} else {
		finalPDF = pdfBytes
		sum := sha256.Sum256(pdfBytes)
		hash = hex.EncodeToString(sum[:])
	}

	// 2. Publicar como PatientDocument (portal + download autenticado; resolve /uploads não-estático).
	// Chave de idempotência por documento: um retry após falha do Updates reaproveita a publicação
	// existente em vez de duplicar/orfanar (CreateFromBytes deduplica por OriginWAMessageID).
	uploadedBy := doc.DoctorID
	filename := fmt.Sprintf("%s_%s.pdf", doc.Type, doc.ID)
	idemKey := "issued:" + doc.ID.String()
	patientDoc, err := s.documents.CreateFromBytes(CreateFromBytesInput{
		PatientID:         doc.PatientID,
		Bytes:             finalPDF,
		Filename:          filename,
		Title:             doc.Title,
		Type:              patientDocTypeFromIssued(doc.Type),
		Source:            models.DocumentSourceStaffUpload,
		UploadedBy:        &uploadedBy,
		OriginWAMessageID: &idemKey,
	})
	if err != nil {
		return nil, fmt.Errorf("erro ao publicar documento: %w", err)
	}

	// 3. Persistir metadados de assinatura
	now := time.Now()
	qr := validationURL
	updates := map[string]interface{}{
		"status":                models.IssuedDocSigned,
		"has_digital_signature": hasDigital,
		"signed_at":             now,
		"signed_pdf_hash":       hash,
		"signed_pdf_path":       patientDoc.FilePath,
		"qr_code_data":          qr,
		"patient_document_id":   patientDoc.ID,
	}
	if certSerial != nil {
		updates["certificate_serial"] = *certSerial
	}
	if err := s.db.Model(&doc).Updates(updates).Error; err != nil {
		return nil, err
	}
	return s.GetByID(doc.ID)
}

// GetForDownload resolve o PDF assinado (via PatientDocument) para download autenticado de staff.
func (s *IssuedDocumentService) GetForDownload(id uuid.UUID) (fullPath, fileName, contentType string, err error) {
	var doc models.IssuedDocument
	if e := s.db.First(&doc, id).Error; e != nil {
		if errors.Is(e, gorm.ErrRecordNotFound) {
			return "", "", "", ErrIssuedDocumentNotFound
		}
		return "", "", "", e
	}
	if doc.PatientDocumentID == nil {
		return "", "", "", errors.New("documento ainda não assinado")
	}
	pdoc, full, e := s.documents.GetForDownload(doc.PatientID, *doc.PatientDocumentID)
	if e != nil {
		return "", "", "", e
	}
	return full, pdoc.FileName, pdoc.ContentType, nil
}

// ValidatePublic confere integridade (hash) do PDF assinado e devolve dados mascarados (público).
func (s *IssuedDocumentService) ValidatePublic(id uuid.UUID) (map[string]interface{}, error) {
	var doc models.IssuedDocument
	if err := s.db.Preload("Doctor").Preload("Patient").First(&doc, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrIssuedDocumentNotFound
		}
		return nil, err
	}

	pdfIntact := false
	if doc.SignedPDFPath != nil && doc.SignedPDFHash != nil {
		full := filepath.Join(s.uploadsRoot, *doc.SignedPDFPath)
		if b, err := os.ReadFile(full); err == nil {
			sum := sha256.Sum256(b)
			pdfIntact = hex.EncodeToString(sum[:]) == *doc.SignedPDFHash
		}
	}

	crm := ""
	if doc.Doctor.CRM != nil && doc.Doctor.CRMUF != nil {
		crm = "CRM-" + *doc.Doctor.CRMUF + " " + *doc.Doctor.CRM
	}

	return map[string]interface{}{
		"valid":     doc.Status == models.IssuedDocSigned && pdfIntact,
		"pdfIntact": pdfIntact,
		"document": map[string]interface{}{
			"id":                  doc.ID,
			"type":                doc.Type,
			"title":               doc.Title,
			"status":              doc.Status,
			"hasDigitalSignature": doc.HasDigitalSignature,
			"issuedAt":            doc.IssuedAt,
			"signedAt":            doc.SignedAt,
		},
		"patient": map[string]interface{}{
			"name": maskName(doc.Patient.Name),
		},
		"doctor": map[string]interface{}{
			"name": doc.Doctor.Name,
			"crm":  crm,
		},
	}, nil
}

// Delete remove um documento ainda em rascunho (assinado é imutável).
func (s *IssuedDocumentService) Delete(id uuid.UUID) error {
	var doc models.IssuedDocument
	if err := s.db.First(&doc, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrIssuedDocumentNotFound
		}
		return err
	}
	if doc.Status == models.IssuedDocSigned {
		return ErrIssuedDocAlreadySigned
	}
	return s.db.Delete(&doc).Error
}

func patientDocTypeFromIssued(t models.IssuedDocumentType) models.PatientDocumentType {
	switch t {
	case models.IssuedDocCertificate:
		return models.DocumentTypeCertificate
	case models.IssuedDocDeclaration:
		return models.DocumentTypeDeclaration
	case models.IssuedDocReport:
		return models.DocumentTypeReport
	default:
		return models.DocumentTypeOther
	}
}

// maskName mostra o primeiro nome + iniciais (validação pública não expõe nome completo).
func maskName(name string) string {
	if name == "" {
		return ""
	}
	parts := []rune(name)
	if len(parts) <= 3 {
		return name
	}
	// mostra os 3 primeiros caracteres + ***
	return string(parts[:3]) + "***"
}

func (s *IssuedDocumentService) toDTO(d *models.IssuedDocument) *dto.IssuedDocumentResponse {
	var appointmentID, signedAt, patientDocID *string
	if d.AppointmentID != nil {
		v := d.AppointmentID.String()
		appointmentID = &v
	}
	if d.SignedAt != nil {
		v := d.SignedAt.Format(time.RFC3339)
		signedAt = &v
	}
	if d.PatientDocumentID != nil {
		v := d.PatientDocumentID.String()
		patientDocID = &v
	}
	doctorName := ""
	if d.Doctor.ID != uuid.Nil {
		doctorName = d.Doctor.Name
	}
	return &dto.IssuedDocumentResponse{
		ID:                  d.ID.String(),
		PatientID:           d.PatientID.String(),
		AppointmentID:       appointmentID,
		DoctorID:            d.DoctorID.String(),
		DoctorName:          doctorName,
		Type:                string(d.Type),
		Title:               d.Title,
		Body:                d.Body,
		Purpose:             d.Purpose,
		DaysOff:             d.DaysOff,
		IncludesCid:         d.IncludesCID,
		CidCode:             d.CIDCode,
		CidConsent:          d.CIDConsent,
		Status:              string(d.Status),
		HasDigitalSignature: d.HasDigitalSignature,
		SignedAt:            signedAt,
		SignedPdfHash:       d.SignedPDFHash,
		CertificateSerial:   d.CertificateSerial,
		QrCodeData:          d.QRCodeData,
		PatientDocumentID:   patientDocID,
		IssuedByUserID:      d.IssuedByUserID.String(),
		IssuedAt:            d.IssuedAt.Format(time.RFC3339),
		CreatedAt:           d.CreatedAt.Format(time.RFC3339),
	}
}
