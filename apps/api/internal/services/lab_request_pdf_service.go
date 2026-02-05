package services

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

type LabRequestPDFService struct {
	db               *gorm.DB
	pdfService       *PDFService
	signatureService *SignatureService
	uploadsPath      string
}

func NewLabRequestPDFService(
	db *gorm.DB,
	pdfService *PDFService,
	signatureService *SignatureService,
	uploadsPath string,
) *LabRequestPDFService {
	return &LabRequestPDFService{
		db:               db,
		pdfService:       pdfService,
		signatureService: signatureService,
		uploadsPath:      uploadsPath,
	}
}

// GenerateSignedLabRequestPDF gera e assina PDF de pedido de exames
func (s *LabRequestPDFService) GenerateSignedLabRequestPDF(
	labRequestID uuid.UUID,
) (pdfURL string, err error) {
	// 1. Buscar lab request com relações
	var labRequest models.LabRequest
	if err := s.db.Preload("Patient").Preload("Doctor").
		First(&labRequest, labRequestID).Error; err != nil {
		return "", err
	}

	// 2. Gerar PDF base usando PDFService existente
	pdfURL, err = s.pdfService.GenerateLabRequestPDF(&labRequest)
	if err != nil {
		return "", fmt.Errorf("erro ao gerar PDF: %v", err)
	}

	// 3. Ler PDF gerado
	pdfPath := fmt.Sprintf("/app/uploads/lab-requests/%s", pdfURL[len("/uploads/lab-requests/"):])
	unsignedPDF, err := os.ReadFile(pdfPath)
	if err != nil {
		return "", fmt.Errorf("erro ao ler PDF: %v", err)
	}

	// 4. Gerar QR Code
	qrCodeData := fmt.Sprintf("https://plenya.com.br/lab-requests/validate/%s", labRequestID)
	labRequest.QRCodeData = &qrCodeData

	// 5. Adicionar QR Code ao PDF (temporariamente salvamos sem QR, 
	//    em uma versão futura podemos integrar diretamente no PDFService)
	// Por enquanto, vamos apenas salvar o QR code data e assinatura

	// 6. Assinar PDF (se médico tiver certificado)
	var signedPDF []byte
	var signatureHash string

	if labRequest.DoctorID != nil {
		signedPDF, signatureHash, err = s.signatureService.SignPrescriptionPDF(
			unsignedPDF,
			*labRequest.DoctorID,
		)
		if err != nil {
			// Se falhar assinatura (médico sem certificado), usa PDF não assinado
			signedPDF = unsignedPDF
			// Calcula hash do PDF não assinado
			hash := sha256.Sum256(unsignedPDF)
			signatureHash = hex.EncodeToString(hash[:])
		} else {
			// Sobrescrever PDF com versão assinada
			if err := os.WriteFile(pdfPath, signedPDF, 0644); err != nil {
				return "", fmt.Errorf("erro ao salvar PDF assinado: %v", err)
			}
		}
	} else {
		signedPDF = unsignedPDF
		hash := sha256.Sum256(unsignedPDF)
		signatureHash = hex.EncodeToString(hash[:])
	}

	// 7. Atualizar LabRequest com metadados
	now := time.Now()
	labRequest.SignedPDFPath = &pdfPath
	labRequest.SignedPDFHash = &signatureHash
	labRequest.SignedAt = &now

	// Atualizar PdfURL também
	labRequest.PdfURL = &pdfURL

	if labRequest.Doctor != nil && labRequest.Doctor.CertificateSerial != nil {
		labRequest.CertificateSerial = labRequest.Doctor.CertificateSerial
	}

	if err := s.db.Save(&labRequest).Error; err != nil {
		return "", fmt.Errorf("erro ao atualizar lab request: %v", err)
	}

	return pdfURL, nil
}

// AddQRCodeToPDF adiciona QR code a um PDF existente (placeholder para implementação futura)
func (s *LabRequestPDFService) AddQRCodeToPDF(pdfBytes []byte, qrCodeData string) ([]byte, error) {
	// Gerar QR Code
	qrCode, err := qrcode.Encode(qrCodeData, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}

	// TODO: Reabrir PDF e adicionar QR Code
	// Por enquanto, retorna o PDF original
	_ = qrCode // evita warning de variável não usada

	return pdfBytes, nil
}
