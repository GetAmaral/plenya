package services

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/jung-kurt/gofpdf/v2"
	"github.com/skip2/go-qrcode"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

type PrescriptionPDFService struct {
	db               *gorm.DB
	signatureService *SignatureService
	sncrService      *SNCRService
	uploadsPath      string
}

func NewPrescriptionPDFService(
	db *gorm.DB,
	signatureService *SignatureService,
	sncrService *SNCRService,
	uploadsPath string,
) *PrescriptionPDFService {
	return &PrescriptionPDFService{
		db:               db,
		signatureService: signatureService,
		sncrService:      sncrService,
		uploadsPath:      uploadsPath,
	}
}

// GetDB returns the database instance (for handler access)
func (s *PrescriptionPDFService) GetDB() *gorm.DB {
	return s.db
}

// GenerateSignedPrescriptionPDF - Fluxo completo de geração e assinatura
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

	// 3. Verificar se precisa de SNCR (se tem algum medicamento controlado)
	needsSNCR := false
	for _, med := range prescription.Medications {
		if med.Category == models.MedCategoryC1 || med.Category == models.MedCategoryC5 {
			needsSNCR = true
			break
		}
	}

	// 4. Solicitar número SNCR (se controlado e ainda não tem)
	if needsSNCR && prescription.SNCRNumber == nil {
		// Para SNCR, usamos o primeiro medicamento controlado como referência
		// (O SNCR é por prescrição, não por medicamento)
		sncrNumber, err := s.sncrService.RequestNumber(&prescription)
		if err != nil {
			return "", fmt.Errorf("erro ao solicitar número SNCR: %v", err)
		}
		if sncrNumber != "" {
			prescription.SNCRNumber = &sncrNumber
			status := "active"
			prescription.SNCRStatus = &status
			s.db.Save(&prescription)
		}
	}

	// 4. Gerar PDF base (não assinado)
	unsignedPDF, err := s.generatePDFContent(&prescription)
	if err != nil {
		return "", fmt.Errorf("erro ao gerar PDF: %v", err)
	}

	// 5. Assinar PDF com certificado do médico
	signedPDF, signatureHash, err := s.signatureService.SignPrescriptionPDF(
		unsignedPDF,
		prescription.DoctorID,
	)
	if err != nil {
		return "", fmt.Errorf("erro ao assinar PDF: %v", err)
	}

	// 6. Salvar PDF assinado
	prescriptionsDir := filepath.Join(s.uploadsPath, "prescriptions")
	os.MkdirAll(prescriptionsDir, 0755)

	filename := fmt.Sprintf("prescription_%s_signed.pdf", prescriptionID)
	pdfPath := filepath.Join(prescriptionsDir, filename)

	if err := os.WriteFile(pdfPath, signedPDF, 0644); err != nil {
		return "", fmt.Errorf("erro ao salvar PDF: %v", err)
	}

	// 7. Atualizar prescrição com metadados
	now := time.Now()
	prescription.SignedPDFPath = &pdfPath
	prescription.SignedPDFHash = &signatureHash
	prescription.SignedAt = &now

	if prescription.Doctor.CertificateSerial != nil {
		prescription.CertificateSerial = prescription.Doctor.CertificateSerial
	}

	if err := s.db.Save(&prescription).Error; err != nil {
		return "", fmt.Errorf("erro ao atualizar prescrição: %v", err)
	}

	return fmt.Sprintf("/uploads/prescriptions/%s", filename), nil
}

// generatePDFContent gera o conteúdo do PDF conforme CFM/ANVISA
func (s *PrescriptionPDFService) generatePDFContent(
	prescription *models.Prescription,
) ([]byte, error) {
	// Create PDF (A4 portrait)
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetAuthor("Plenya EMR - Prescrição Digital", true)
	pdf.SetCreator("Plenya EMR", true)
	pdf.SetTitle("Prescrição Médica Digital", true)

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

	// === HEADER: PRESCRIÇÃO DIGITAL ===
	pdf.SetFont("OpenSans", "B", 16)
	pdf.SetXY(20, y)
	pdf.Cell(170, 10, "PRESCRIÇÃO MÉDICA DIGITAL")
	y += 15

	// SNCR Number (se existir)
	if prescription.SNCRNumber != nil {
		pdf.SetFont("OpenSans", "B", 10)
		pdf.SetXY(20, y)
		pdf.Cell(170, 6, fmt.Sprintf("SNCR: %s", *prescription.SNCRNumber))
		y += 8
	}

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

	// Iterar pelos medicamentos
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

	// === QR CODE ===
	qrCodeData := fmt.Sprintf("https://plenya.com.br/prescriptions/validate/%s", prescription.ID)
	qrCode, _ := qrcode.Encode(qrCodeData, qrcode.Medium, 256)

	// Salvar QR temporário
	qrPath := fmt.Sprintf("/tmp/qr_%s.png", prescription.ID)
	os.WriteFile(qrPath, qrCode, 0644)
	defer os.Remove(qrPath)

	// Adicionar QR ao PDF
	pdf.Image(qrPath, 20, y, 40, 40, false, "", 0, "")

	// Texto ao lado do QR
	pdf.SetFont("OpenSans", "", 9)
	pdf.SetXY(65, y)
	qrText := "Validar prescrição:\n" + qrCodeData + "\n\n"
	qrText += "Documento assinado digitalmente\ncom certificado ICP-Brasil"
	pdf.MultiCell(125, 4, qrText, "", "", false)

	// Salvar QR code data na prescrição
	prescription.QRCodeData = &qrCodeData
	s.db.Save(prescription)

	// Gerar bytes do PDF
	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// formatCPF já definido em certificate_service.go (função compartilhada no package)
