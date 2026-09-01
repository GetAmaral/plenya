package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/repository"
)

var ErrNoSnapshotForReport = errors.New("paciente sem escore calculado — calcule o escore antes de gerar o relatório")

// CarePlanReportService gera o relatório longitudinal AGIR (Escore + biomarcadores
// Normal-vs-Ótimo + plano por pilar), converte HTML→PDF (go-rod), assina (ICP-Brasil ou
// degrada para assinatura manual) e publica no portal do paciente como IssuedDocument(report).
type CarePlanReportService struct {
	db           *gorm.DB
	snapshotRepo *repository.ScoreSnapshotRepository
	carePlan     *CarePlanService
	signature    *SignatureService
	documents    *PatientDocumentsService
}

func NewCarePlanReportService(
	db *gorm.DB,
	snapshotRepo *repository.ScoreSnapshotRepository,
	carePlan *CarePlanService,
	signature *SignatureService,
	documents *PatientDocumentsService,
) *CarePlanReportService {
	return &CarePlanReportService{db: db, snapshotRepo: snapshotRepo, carePlan: carePlan, signature: signature, documents: documents}
}

// GenerateAndPublish compõe, assina e publica o relatório. Retorna o IssuedDocument criado.
func (s *CarePlanReportService) GenerateAndPublish(patientID, doctorID uuid.UUID) (string, error) {
	var patient models.Patient
	if err := s.db.First(&patient, patientID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", ErrPatientNotFound
		}
		return "", err
	}

	snapshot, err := s.snapshotRepo.GetLatestByPatientID(patientID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", ErrNoSnapshotForReport
		}
		return "", err
	}

	items, err := s.carePlan.ListActiveModels(patientID)
	if err != nil {
		return "", err
	}

	var doctor models.User
	if err := s.db.First(&doctor, doctorID).Error; err != nil {
		return "", err
	}

	// 1. IssuedDocument(report) em draft → pega o ID p/ o QR de validação.
	doc := &models.IssuedDocument{
		PatientID:      patientID,
		DoctorID:       doctorID,
		Type:           models.IssuedDocReport,
		Title:          "Relatório de Saúde, Performance e Longevidade",
		Body:           "Relatório longitudinal AGIR.",
		Status:         models.IssuedDocDraft,
		IssuedByUserID: doctorID,
	}
	if err := s.db.Create(doc).Error; err != nil {
		return "", err
	}

	validationURL := fmt.Sprintf("https://app.plenyasaude.com.br/documentos/validar/%s", doc.ID)
	reportNow := time.Now()

	render := func(digital bool) ([]byte, error) {
		return renderCarePlanReportBytes(&patient, snapshot, items, &doctor, validationURL, digital, reportNow)
	}

	out, err := signOrDegrade(s.signature, &doctor, doctorID,
		"Assinatura Digital de Relatório Médico", "Plenya EMR - Relatório AGIR", render)
	if err != nil {
		return "", fmt.Errorf("erro ao gerar PDF do relatório: %w", err)
	}
	finalPDF, hash, hasDigital, certSerial := out.Bytes, out.Hash, out.Digital, out.CertSerial

	// 2. Publicar no portal (idempotente por doc.ID — retry após falha do Updates não duplica).
	uploadedBy := doctorID
	idemKey := "report:" + doc.ID.String()
	patientDoc, err := s.documents.CreateFromBytes(CreateFromBytesInput{
		PatientID:         patientID,
		Bytes:             finalPDF,
		Filename:          fmt.Sprintf("relatorio_agir_%s.pdf", doc.ID),
		Title:             doc.Title,
		Type:              models.DocumentTypeReport,
		Source:            models.DocumentSourceStaffUpload,
		UploadedBy:        &uploadedBy,
		OriginWAMessageID: &idemKey,
	})
	if err != nil {
		return "", fmt.Errorf("erro ao publicar relatório: %w", err)
	}

	// 3. Persistir metadados.
	now := time.Now()
	updates := map[string]interface{}{
		"status":                models.IssuedDocSigned,
		"has_digital_signature": hasDigital,
		"signed_at":             now,
		"signed_pdf_hash":       hash,
		"signed_pdf_path":       patientDoc.FilePath,
		"qr_code_data":          validationURL,
		"patient_document_id":   patientDoc.ID,
	}
	if certSerial != nil {
		updates["certificate_serial"] = *certSerial
	}
	if err := s.db.Model(doc).Updates(updates).Error; err != nil {
		return "", err
	}

	return doc.ID.String(), nil
}

var reportLevelLabel = map[int]string{
	0: "Crítico", 1: "Muito alterado", 2: "Subótimo", 3: "Limítrofe", 4: "Bom", 5: "Ótimo", 6: "N/A",
}
var agirNames = map[string]string{"A": "Atividade", "G": "Gestão", "I": "Integração", "R": "Ritmo"}
