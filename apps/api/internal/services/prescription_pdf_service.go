package services

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"time"

	gofpdf "codeberg.org/go-pdf/fpdf"
	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
	"gorm.io/gorm"

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
		First(&prescription, prescriptionID).Error; err != nil {
		return "", err
	}

	// 2. Validar dados obrigatórios
	if prescription.Doctor.CRM == nil {
		return "", fmt.Errorf("médico sem CRM cadastrado")
	}
	if prescription.Patient.CPF == nil || *prescription.Patient.CPF == "" {
		return "", fmt.Errorf("paciente sem CPF")
	}
	if len(prescription.Medications) == 0 {
		return "", fmt.Errorf("prescrição sem medicamentos")
	}

	// 3. Decidir o modo de assinatura.
	hasControlled := prescriptionHasControlled(&prescription)
	// Controlado => manual obrigatório. Não controlado => digital se houver cert ativo.
	manual := hasControlled || !prescription.Doctor.CertificateActive

	// 4. Gerar o PDF base (layout muda conforme o modo).
	unsignedPDF, err := s.generatePDFContent(&prescription, manual)
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
			manualPDF, bErr := s.generatePDFContent(&prescription, true)
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
		Title:      "Receita médica - " + now.Format("02/01/2006"),
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
	// Modo manual não tem validação digital por QR.
	if mode != "digital" {
		prescription.QRCodeData = nil
	}

	if err := s.db.Save(&prescription).Error; err != nil {
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

// prescriptionHasControlled indica se há medicamento de controle especial (C1/C5).
func prescriptionHasControlled(prescription *models.Prescription) bool {
	for _, med := range prescription.Medications {
		if med.Category == models.MedCategoryC1 || med.Category == models.MedCategoryC5 {
			return true
		}
	}
	return false
}

// generatePDFContent gera o conteúdo do PDF. Quando manual=true, monta a receita para
// impressão/assinatura à mão (bloco de assinatura+carimbo, sem QR/selo digital).
func (s *PrescriptionPDFService) generatePDFContent(
	prescription *models.Prescription,
	manual bool,
) ([]byte, error) {
	hasControlled := prescriptionHasControlled(prescription)

	// Create PDF (A4 portrait)
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetAuthor("Plenya EMR - Prescrição", true)
	pdf.SetCreator("Plenya EMR", true)
	pdf.SetTitle("Receita Médica", true)

	// Load fonts
	regularFont, err := os.ReadFile("/usr/share/fonts/opensans/OpenSans-Regular.ttf")
	if err != nil {
		return nil, err
	}
	boldFont, err := os.ReadFile("/usr/share/fonts/opensans/OpenSans-Bold.ttf")
	if err != nil {
		return nil, err
	}

	pdf.AddUTF8FontFromBytes("OpenSans", "", regularFont)
	pdf.AddUTF8FontFromBytes("OpenSans", "B", boldFont)
	pdf.SetFont("OpenSans", "", 10)

	pdf.AddPage()

	// Add letterhead background
	pdf.ImageOptions(
		"/app/PlenyaA4-150dpi.png",
		0, 0, 210, 297,
		false,
		gofpdf.ImageOptions{ImageType: "PNG"},
		0, "",
	)

	y := 50.0

	// === HEADER ===
	title := "PRESCRIÇÃO MÉDICA DIGITAL"
	if manual {
		if hasControlled {
			title = "RECEITA DE CONTROLE ESPECIAL"
		} else {
			title = "RECEITA MÉDICA"
		}
	}
	pdf.SetFont("OpenSans", "B", 16)
	pdf.SetXY(20, y)
	pdf.Cell(170, 10, title)
	y += 15

	// Datas
	pdf.SetFont("OpenSans", "", 10)
	pdf.SetXY(20, y)
	pdf.Cell(85, 6, fmt.Sprintf("Emissão: %s", prescription.PrescriptionDate.Format("02/01/2006")))
	pdf.SetXY(105, y)
	pdf.Cell(85, 6, fmt.Sprintf("Validade: %s", prescription.ValidUntil.Format("02/01/2006")))
	y += 12

	// === DADOS DO MÉDICO ===
	pdf.SetFont("OpenSans", "B", 11)
	pdf.SetXY(20, y)
	pdf.Cell(170, 6, "MÉDICO PRESCRITOR")
	y += 6

	pdf.SetFont("OpenSans", "", 10)
	doctorInfo := prescription.Doctor.Name + "\n"
	doctorInfo += fmt.Sprintf("CRM-%s %s", *prescription.Doctor.CRMUF, *prescription.Doctor.CRM)

	if prescription.Doctor.Specialty != nil {
		doctorInfo += fmt.Sprintf(" - %s", *prescription.Doctor.Specialty)
	}
	if prescription.Doctor.ProfessionalAddress != nil {
		doctorInfo += "\n" + *prescription.Doctor.ProfessionalAddress
	}
	if prescription.Doctor.ProfessionalPhone != nil {
		doctorInfo += "\nTel: " + *prescription.Doctor.ProfessionalPhone
	}

	pdf.SetXY(20, y)
	pdf.MultiCell(170, 5, doctorInfo, "", "", false)
	y += 25

	// === DADOS DO PACIENTE ===
	pdf.SetFont("OpenSans", "B", 11)
	pdf.SetXY(20, y)
	pdf.Cell(170, 6, "PACIENTE")
	y += 6

	pdf.SetFont("OpenSans", "", 10)
	patientInfo := prescription.Patient.Name + "\n"
	if prescription.Patient.CPF != nil && *prescription.Patient.CPF != "" {
		cpfFormatted := formatCPF(*prescription.Patient.CPF)
		patientInfo += fmt.Sprintf("CPF: %s", cpfFormatted)
	}

	if prescription.Patient.Address != nil {
		patientInfo += "\n" + *prescription.Patient.Address
	}

	pdf.SetXY(20, y)
	pdf.MultiCell(170, 5, patientInfo, "", "", false)
	y += 20

	// === MEDICAMENTOS PRESCRITOS ===
	pdf.SetFont("OpenSans", "B", 11)
	pdf.SetXY(20, y)
	pdf.Cell(170, 6, "MEDICAMENTOS PRESCRITOS")
	y += 8

	for i, med := range prescription.Medications {
		pdf.SetFont("OpenSans", "B", 10)
		pdf.SetXY(20, y)
		pdf.Cell(170, 6, fmt.Sprintf("%d. %s", i+1, med.MedicationName))
		y += 6

		pdf.SetFont("OpenSans", "", 10)
		medicationInfo := fmt.Sprintf("   Princípio ativo: %s", med.ActiveIngredient) + "\n"
		medicationInfo += fmt.Sprintf("   Concentração: %s", med.Concentration) + "\n"
		medicationInfo += fmt.Sprintf("   Quantidade: %d (%s)", med.Quantity, med.QuantityInWords) + "\n"
		medicationInfo += fmt.Sprintf("   Posologia: %s %s", med.Dosage, med.Frequency) + "\n"
		medicationInfo += fmt.Sprintf("   Via: %s", med.Route) + "\n"
		medicationInfo += fmt.Sprintf("   Duração: %d dias", med.Duration)

		if med.Instructions != nil && *med.Instructions != "" {
			medicationInfo += "\n   Instruções específicas: " + *med.Instructions
		}

		pdf.SetXY(20, y)
		pdf.MultiCell(170, 5, medicationInfo, "", "", false)
		y += 40
	}

	// Instruções gerais (se houver)
	if prescription.GeneralInstructions != nil && *prescription.GeneralInstructions != "" {
		pdf.SetFont("OpenSans", "B", 10)
		pdf.SetXY(20, y)
		pdf.Cell(170, 6, "Instruções Gerais:")
		y += 6
		pdf.SetFont("OpenSans", "", 10)
		pdf.SetXY(20, y)
		pdf.MultiCell(170, 5, *prescription.GeneralInstructions, "", "", false)
		y += 15
	}

	if manual {
		s.addManualSignatureBlock(pdf, y, hasControlled)
	} else {
		s.addDigitalSignatureBlock(pdf, prescription, y)
	}

	// Gerar bytes do PDF
	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// addManualSignatureBlock — receita para impressão: linha de assinatura + carimbo.
// Para controlados, acrescenta a nota de dispensação física (Portaria SVS/MS 344/98).
func (s *PrescriptionPDFService) addManualSignatureBlock(pdf *gofpdf.Fpdf, y float64, hasControlled bool) {
	if y > 245 {
		y = 245
	}
	y += 16
	pdf.Line(60, y, 150, y)
	pdf.SetFont("OpenSans", "", 9)
	pdf.SetXY(25, y+2)
	pdf.Cell(120, 5, "Assinatura e Carimbo do Médico")
	y += 10

	pdf.SetFont("OpenSans", "", 7)
	pdf.SetTextColor(60, 60, 60)
	pdf.SetXY(20, y)
	pdf.Cell(170, 3, "Receita para impressão, assinatura e carimbo do médico. Sem assinatura digital.")
	if hasControlled {
		y += 3
		pdf.SetXY(20, y)
		pdf.MultiCell(170, 3,
			"Medicamento sujeito a controle especial (Portaria SVS/MS 344/98): dispensação mediante "+
				"receituário/Notificação de Receita em via física, conforme a regulamentação vigente.",
			"", "", false)
	}
	pdf.SetTextColor(0, 0, 0)
}

// addDigitalSignatureBlock — QR de validação + selo "assinado digitalmente" (modo digital).
func (s *PrescriptionPDFService) addDigitalSignatureBlock(pdf *gofpdf.Fpdf, prescription *models.Prescription, y float64) {
	qrCodeData := fmt.Sprintf("https://plenya.com.br/prescriptions/validate/%s", prescription.ID)
	qrCode, _ := qrcode.Encode(qrCodeData, qrcode.Medium, 256)

	qrPath := fmt.Sprintf("/tmp/qr_%s.png", prescription.ID)
	os.WriteFile(qrPath, qrCode, 0644)
	defer os.Remove(qrPath)

	pdf.Image(qrPath, 20, y, 40, 40, false, "", 0, "")

	pdf.SetFont("OpenSans", "", 9)
	pdf.SetXY(65, y)
	qrText := "Validar prescrição:\n" + qrCodeData + "\n\n"
	qrText += "Documento assinado digitalmente com certificado ICP-Brasil (PAdES).\n"
	qrText += "Verificar assinatura: https://validar.iti.gov.br"
	pdf.MultiCell(125, 4, qrText, "", "", false)

	// Persistido pelo Save final em GenerateSignedPrescriptionPDF.
	prescription.QRCodeData = &qrCodeData
}

// formatCPF já definido em certificate_service.go (função compartilhada no package)
