package services

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

// OCRService - serviço de extração de texto via OCR
type OCRService struct{}

// NewOCRService cria uma nova instância do serviço OCR
func NewOCRService() *OCRService {
	return &OCRService{}
}

// ExtractText - estratégia híbrida: pdftotext → Tesseract fallback
// Primeiro tenta pdftotext (rápido, PDFs nativos com texto)
// Se falhar ou retornar pouco texto, usa Tesseract (scanned PDFs)
func (s *OCRService) ExtractText(pdfPath string) (string, error) {
	// 1. Tentar pdftotext primeiro (MUITO mais rápido para PDFs nativos)
	text, err := s.extractWithPdfToText(pdfPath)
	if err == nil && len(strings.TrimSpace(text)) > 100 {
		// Sucesso com pdftotext e tem conteúdo razoável
		return text, nil
	}

	// 2. Fallback: Tesseract OCR (para PDFs escaneados)
	return s.extractWithTesseract(pdfPath)
}

// extractWithPdfToText - extração rápida com poppler-utils
func (s *OCRService) extractWithPdfToText(pdfPath string) (string, error) {
	cmd := exec.Command("pdftotext", "-layout", pdfPath, "-")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("pdftotext failed: %v - %s", err, stderr.String())
	}

	return out.String(), nil
}

// extractWithTesseract - OCR completo com ImageMagick + Tesseract
func (s *OCRService) extractWithTesseract(pdfPath string) (string, error) {
	// Criar diretório temporário para imagens
	tmpDir := filepath.Join("/tmp", "ocr-"+uuid.NewString())
	defer os.RemoveAll(tmpDir)

	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create temp dir: %v", err)
	}

	// Converter PDF → imagens PNG (300 DPI para boa qualidade OCR)
	convertCmd := exec.Command(
		"convert",
		"-density", "300",       // Alta resolução para OCR
		pdfPath,
		"-quality", "90",        // Qualidade da imagem
		filepath.Join(tmpDir, "page_%d.png"),
	)

	var convertStderr bytes.Buffer
	convertCmd.Stderr = &convertStderr

	if err := convertCmd.Run(); err != nil {
		return "", fmt.Errorf("imagemagick convert failed: %v - %s", err, convertStderr.String())
	}

	// OCR cada página
	files, err := filepath.Glob(filepath.Join(tmpDir, "page_*.png"))
	if err != nil {
		return "", fmt.Errorf("failed to list images: %v", err)
	}

	if len(files) == 0 {
		return "", fmt.Errorf("no images generated from PDF")
	}

	var allText strings.Builder

	for i, imgPath := range files {
		// Tesseract com português
		cmd := exec.Command("tesseract", imgPath, "stdout", "-l", "por")
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr

		if err := cmd.Run(); err != nil {
			// Log mas continua com próximas páginas
			fmt.Printf("⚠️  Tesseract failed for page %d: %v - %s\n", i+1, err, stderr.String())
			continue
		}

		allText.WriteString(out.String())
		allText.WriteString("\n\n--- PÁGINA " + fmt.Sprintf("%d", i+1) + " ---\n\n")
	}

	result := allText.String()
	if len(strings.TrimSpace(result)) < 50 {
		return "", fmt.Errorf("extracted text too short (< 50 chars), OCR may have failed")
	}

	return result, nil
}
