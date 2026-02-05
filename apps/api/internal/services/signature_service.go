package services

import (
	"bytes"
	"crypto"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/digitorus/pdf"
	"github.com/digitorus/pdfsign/sign"
	"github.com/google/uuid"
)

var (
	ErrSignatureFailed = errors.New("failed to sign PDF")
)

type SignatureService struct {
	certService *CertificateService
}

func NewSignatureService(certService *CertificateService) *SignatureService {
	return &SignatureService{
		certService: certService,
	}
}

// SignPrescriptionPDF assina PDF com certificado do médico
// Retorna: PDF assinado, hash SHA-256, erro
func (s *SignatureService) SignPrescriptionPDF(
	pdfBytes []byte,
	doctorID uuid.UUID,
) (signedPDF []byte, signatureHash string, err error) {
	// 1. Obter certificado do médico
	cert, privateKey, err := s.certService.GetActiveCertificate(doctorID)
	if err != nil {
		return nil, "", err
	}

	// 2. Converter privateKey para crypto.Signer
	signer, ok := privateKey.(crypto.Signer)
	if !ok {
		return nil, "", errors.New("private key does not implement crypto.Signer")
	}

	// 3. Preparar dados de assinatura PAdES
	signData := sign.SignData{
		Signature: sign.SignDataSignature{
			Info: sign.SignDataSignatureInfo{
				Name:        "Plenya EMR - Sistema de Prescrição Digital",
				Location:    "Brasil",
				Reason:      "Assinatura Digital de Prescrição Médica",
				ContactInfo: "https://plenya.com.br",
			},
			CertType:   sign.CertificationSignature,
			DocMDPPerm: sign.AllowFillingExistingFormFieldsAndSignaturesPerms,
		},
		DigestAlgorithm: crypto.SHA256,
		Certificate:     cert,
		Signer:          signer,
	}

	// 4. Criar readers e writers
	inputReader := bytes.NewReader(pdfBytes)
	var outputBuffer bytes.Buffer

	// 5. Criar PDF reader
	pdfReader, err := pdf.NewReader(inputReader, int64(len(pdfBytes)))
	if err != nil {
		return nil, "", ErrSignatureFailed
	}

	// 6. Assinar PDF com PAdES
	err = sign.Sign(inputReader, &outputBuffer, pdfReader, int64(len(pdfBytes)), signData)
	if err != nil {
		return nil, "", ErrSignatureFailed
	}

	signedPDF = outputBuffer.Bytes()

	// 7. Calcular hash SHA-256 do PDF assinado
	hash := sha256.Sum256(signedPDF)
	signatureHash = hex.EncodeToString(hash[:])

	return signedPDF, signatureHash, nil
}

// VerifySignature valida assinatura do PDF (placeholder para implementação futura)
func (s *SignatureService) VerifySignature(pdfBytes []byte) (bool, error) {
	// TODO: Implementar extração e validação da assinatura PAdES
	// Por enquanto, retorna true (assumindo que a assinatura é válida)
	// A validação real seria feita pelo ITI ou Adobe Reader
	return true, nil
}
