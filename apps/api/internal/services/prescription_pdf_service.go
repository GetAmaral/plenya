package services

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/plenya/api/internal/models"
)

type PrescriptionPDFService struct {
	db               *gorm.DB
	signatureService *SignatureService
	sncrService      *SNCRService
	documents        *PatientDocumentsService
	uploadsPath      string
}

func NewPrescriptionPDFService(
	db *gorm.DB,
	signatureService *SignatureService,
	sncrService *SNCRService,
	documents *PatientDocumentsService,
	uploadsPath string,
) *PrescriptionPDFService {
	return &PrescriptionPDFService{
		db:               db,
		signatureService: signatureService,
		sncrService:      sncrService,
		documents:        documents,
		uploadsPath:      uploadsPath,
	}
}

// GetDB returns the database instance (for handler access)
func (s *PrescriptionPDFService) GetDB() *gorm.DB {
	return s.db
}

// VerifySignature delega ao SignatureService a verificação criptográfica da assinatura PAdES.
func (s *PrescriptionPDFService) VerifySignature(pdfBytes []byte) (*SignatureVerification, error) {
	return s.signatureService.VerifySignature(pdfBytes)
}

// GenerateSignedPrescriptionPDF gera o PDF da prescrição e define o modo de assinatura:
//
//   - CONTROLADO (C1/C5): SEMPRE modo MANUAL. O sistema gera a receita para o médico
//     IMPRIMIR, CARIMBAR e ASSINAR à mão. Sem número SNCR (a integração ainda não abriu,
//     ver docs/emr/lei-receita-cfm-anvisa-2026.md) e sem assinatura digital fingindo validade.
//   - NÃO CONTROLADO: assina digitalmente com ICP-Brasil se o médico tiver certificado ativo;
//     se não tiver (ou a assinatura falhar), degrada graciosamente para o modo manual.
func (s *PrescriptionPDFService) GenerateSignedPrescriptionPDF(
	prescriptionID uuid.UUID,
) (pdfURL string, err error) {
	// 1. Buscar prescrição com relações
	var prescription models.Prescription
	if err := s.db.Preload("Patient").Preload("Doctor").Preload("Medications").
		Preload("Formulas", func(db *gorm.DB) *gorm.DB { return db.Order("display_order") }).
		Preload("Formulas.Components", func(db *gorm.DB) *gorm.DB { return db.Order("display_order") }).
		First(&prescription, prescriptionID).Error; err != nil {
		return "", err
	}

	// 2. Validar dados obrigatórios
	if prescription.Doctor.CRM == nil {
		return "", fmt.Errorf("médico sem CRM cadastrado")
	}
	// CPF do paciente NÃO é exigido: a recepção cadastra paciente só pelo nome, e a receita
	// identifica o paciente pelo nome (o CPF entra mascarado quando existe). Exigir CPF aqui
	// derrubava a emissão de receita de paciente novo, sem que a lei peça isso.
	if strings.TrimSpace(prescription.Patient.Name) == "" {
		return "", fmt.Errorf("paciente sem nome")
	}
	if len(prescription.Medications) == 0 && len(prescription.Formulas) == 0 {
		return "", fmt.Errorf("prescrição sem medicamentos")
	}

	// 3. Decidir o modo de assinatura.
	hasControlled := prescriptionHasControlled(&prescription)
	// Controlado => manual obrigatório. Não controlado => digital se houver cert ativo.
	manual := hasControlled || !prescription.Doctor.CertificateActive

	// 4. Gerar o PDF base (pipeline pdfdoc; layout muda conforme o modo).
	// No modo digital, pré-setar SignedAt para o selo ICP-Brasil exibir o carimbo temporal.
	if !manual {
		nowSig := time.Now()
		prescription.SignedAt = &nowSig
	}
	unsignedPDF, err := renderPrescriptionBytes(&prescription, manual)
	if err != nil {
		return "", fmt.Errorf("erro ao gerar PDF: %v", err)
	}

	var finalPDF []byte
	var signatureHash string
	mode := "manual"

	if manual {
		finalPDF = unsignedPDF
		sum := sha256.Sum256(unsignedPDF)
		signatureHash = hex.EncodeToString(sum[:])
	} else {
		signed, sigHash, sErr := s.signatureService.SignPrescriptionPDF(unsignedPDF, prescription.DoctorID)
		if sErr != nil {
			// Degradação graciosa (padrão IssuedDocument): regenera em modo manual.
			prescription.SignedAt = nil
			manualPDF, bErr := renderPrescriptionBytes(&prescription, true)
			if bErr != nil {
				return "", fmt.Errorf("erro ao regenerar PDF: %v", bErr)
			}
			finalPDF = manualPDF
			sum := sha256.Sum256(manualPDF)
			signatureHash = hex.EncodeToString(sum[:])
		} else {
			finalPDF = signed
			signatureHash = sigHash
			mode = "digital"
		}
	}

	// 5. Publicar como PatientDocument (download autenticado; o /uploads estático foi removido
	//    porque vazava PDFs sem auth). Re-assinatura substitui o documento anterior.
	now := time.Now()
	if prescription.PatientDocumentID != nil {
		_ = s.documents.Delete(*prescription.PatientDocumentID)
	}
	uploadedBy := prescription.DoctorID
	patientDoc, err := s.documents.CreateFromBytes(CreateFromBytesInput{
		PatientID:  prescription.PatientID,
		Bytes:      finalPDF,
		Filename:   fmt.Sprintf("receita_%s.pdf", prescriptionID),
		Title:      documentTitleFor(&prescription) + " - " + now.Format("02/01/2006"),
		Type:       models.DocumentTypePrescription,
		Source:     models.DocumentSourceStaffUpload,
		UploadedBy: &uploadedBy,
	})
	if err != nil {
		return "", fmt.Errorf("erro ao publicar prescrição: %v", err)
	}

	// 6. Atualizar prescrição com metadados
	prescription.PatientDocumentID = &patientDoc.ID
	prescription.SignedPDFPath = &patientDoc.FilePath // relativo ao uploadsPath
	prescription.SignedPDFHash = &signatureHash
	prescription.SignedAt = &now
	prescription.SignatureMode = &mode
	if mode == "digital" && prescription.Doctor.CertificateSerial != nil {
		prescription.CertificateSerial = prescription.Doctor.CertificateSerial
	} else {
		prescription.CertificateSerial = nil
	}
	// A URL de validação que foi impressa no QR fica gravada junto: o campo tinha ficado só com
	// escrita de nil desde que o PDF passou a montar o QR sozinho, e ficava permanentemente nulo
	// na resposta da API enquanto o papel trazia o código.
	if mode == "digital" {
		url := fmt.Sprintf("https://app.plenyasaude.com.br/prescriptions/validate/%s", prescription.ID)
		prescription.QRCodeData = &url
	} else {
		// Modo manual não tem validação digital por QR.
		prescription.QRCodeData = nil
	}

	// Omit das associações: `prescription` veio com Patient, Doctor e Medications carregados, e
	// Save por default faz upsert dos filhos — gravaria de volta paciente e usuário a cada
	// assinatura. Aqui só interessam os metadados da assinatura.
	if err := s.db.Omit(clause.Associations).Save(&prescription).Error; err != nil {
		return "", fmt.Errorf("erro ao atualizar prescrição: %v", err)
	}

	// URL de download autenticado (a web baixa via fetch+blob com o Bearer token).
	return fmt.Sprintf("/api/v1/prescriptions/%s/download", prescriptionID), nil
}

// GetForDownload resolve o PDF da prescrição (via PatientDocument) para download autenticado.
func (s *PrescriptionPDFService) GetForDownload(prescriptionID uuid.UUID) (fullPath, fileName, contentType string, err error) {
	var p models.Prescription
	if e := s.db.First(&p, prescriptionID).Error; e != nil {
		return "", "", "", e
	}
	if p.PatientDocumentID == nil {
		return "", "", "", errors.New("prescrição ainda não gerada/assinada")
	}
	pdoc, full, e := s.documents.GetForDownload(p.PatientID, *p.PatientDocumentID)
	if e != nil {
		return "", "", "", e
	}
	return full, pdoc.FileName, pdoc.ContentType, nil
}

// ReadSignedPDF lê os bytes do PDF assinado (para validação pública por hash/assinatura).
// Prefere o PatientDocument; cai pro caminho legado (absoluto) em prescrições antigas.
func (s *PrescriptionPDFService) ReadSignedPDF(p *models.Prescription) ([]byte, error) {
	if p.PatientDocumentID != nil {
		_, full, err := s.documents.GetForDownload(p.PatientID, *p.PatientDocumentID)
		if err != nil {
			return nil, err
		}
		return os.ReadFile(full)
	}
	if p.SignedPDFPath == nil {
		return nil, errors.New("sem PDF assinado")
	}
	return os.ReadFile(*p.SignedPDFPath)
}

// documentTitleFor — o paciente vê este título no portal e no WhatsApp; manipulado e
// industrializado chegam misturados na mesma lista de documentos.
func documentTitleFor(p *models.Prescription) string {
	if p.Type == models.PrescriptionCompounded {
		return "Receita de manipulado"
	}
	return "Receita médica"
}

// prescriptionHasControlled indica se a receita tem item de controle especial. É o ÚNICO ponto
// que decide modo manual e rótulo de Controle Especial, nos dois tipos de receita.
func prescriptionHasControlled(prescription *models.Prescription) bool {
	for _, med := range prescription.Medications {
		if models.IsControlled(med.Category) {
			return true
		}
	}
	for _, f := range prescription.Formulas {
		if models.IsControlled(f.HighestCategory) {
			return true
		}
	}
	return false
}

// formatCPF já definido em certificate_service.go (função compartilhada no package)
