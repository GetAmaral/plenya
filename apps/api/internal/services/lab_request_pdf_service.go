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

	// 2. Verificar se o médico tem certificado ativo ANTES de gerar o PDF
	var willBeDigitallySigned bool
	if labRequest.Doctor != nil && labRequest.Doctor.CertificateActive {
		// Médico tem certificado ativo - PDF será assinado digitalmente
		willBeDigitallySigned = true

		// Pré-configurar metadados da assinatura ANTES de gerar o PDF
		now := time.Now()
		labRequest.SignedAt = &now

		// Copiar informações do certificado do médico para o lab request
		if labRequest.Doctor.CertificateSerial != nil {
			labRequest.CertificateSerial = labRequest.Doctor.CertificateSerial
		}
	}

	// 3. Gerar QR Code data
	qrCodeData := fmt.Sprintf("https://plenya.com.br/lab-requests/validate/%s", labRequestID)
	labRequest.QRCodeData = &qrCodeData

	// 4. Gerar PDF (agora com SignedAt já configurado se for assinado digitalmente)
	pdfURL, err = s.pdfService.GenerateLabRequestPDF(&labRequest)
	if err != nil {
		return "", fmt.Errorf("erro ao gerar PDF: %v", err)
	}

	// 5. Se será assinado digitalmente, ler PDF e assinar
	pdfPath := fmt.Sprintf("/app/uploads/lab-requests/%s", pdfURL[len("/uploads/lab-requests/"):])

	if willBeDigitallySigned {
		unsignedPDF, err := os.ReadFile(pdfPath)
		if err != nil {
			return "", fmt.Errorf("erro ao ler PDF: %v", err)
		}

		// Assinar PDF com certificado do médico
		signedPDF, signatureHash, err := s.signatureService.SignPrescriptionPDF(
			unsignedPDF,
			*labRequest.DoctorID,
		)
		if err != nil {
			// Se falhar assinatura, reverter SignedAt e gerar PDF sem assinatura
			labRequest.SignedAt = nil
			labRequest.CertificateSerial = nil

			// Regerar PDF sem assinatura digital
			pdfURL, err = s.pdfService.GenerateLabRequestPDF(&labRequest)
			if err != nil {
				return "", fmt.Errorf("erro ao gerar PDF sem assinatura: %v", err)
			}

			// Calcular hash do PDF não assinado
			hash := sha256.Sum256(unsignedPDF)
			signatureHash = hex.EncodeToString(hash[:])

			labRequest.SignedPDFPath = &pdfPath
			labRequest.SignedPDFHash = &signatureHash
		} else {
			// Sobrescrever PDF com versão assinada
			if err := os.WriteFile(pdfPath, signedPDF, 0644); err != nil {
				return "", fmt.Errorf("erro ao salvar PDF assinado: %v", err)
			}

			labRequest.SignedPDFPath = &pdfPath
			labRequest.SignedPDFHash = &signatureHash
		}
	} else {
		// PDF não assinado - calcular hash do arquivo
		unsignedPDF, err := os.ReadFile(pdfPath)
		if err == nil {
			hash := sha256.Sum256(unsignedPDF)
			signatureHash := hex.EncodeToString(hash[:])
			labRequest.SignedPDFPath = &pdfPath
			labRequest.SignedPDFHash = &signatureHash
		}
	}

	// 6. Atualizar PdfURL
	labRequest.PdfURL = &pdfURL

	// 7. Salvar metadados no banco
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
