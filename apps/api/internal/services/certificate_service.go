package services

import (
	"crypto"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"software.sslmate.com/src/go-pkcs12"

	internalCrypto "github.com/plenya/api/internal/crypto"
	"github.com/plenya/api/internal/models"
)

var (
	ErrInvalidCertificate     = errors.New("invalid certificate or password")
	ErrCertificateExpired     = errors.New("certificate has expired")
	ErrNotICPBrasil           = errors.New("certificate is not ICP-Brasil")
	ErrNoCertificate          = errors.New("user has no certificate")
	ErrCertificateExpiredUser = errors.New("user's certificate has expired")
	ErrCPFMismatch            = errors.New("certificate CPF does not match user CPF")
)

type CertificateService struct {
	db            *gorm.DB
	encryptionKey string
}

func NewCertificateService(db *gorm.DB, encryptionKey string) *CertificateService {
	return &CertificateService{
		db:            db,
		encryptionKey: encryptionKey,
	}
}

// CPFValidationResult indica o resultado da validação de CPF
type CPFValidationResult struct {
	Valid             bool
	UserCPF           string
	CertificateCPF    string
	RequiresConfirmation bool
}

// UploadCertificate faz upload e valida certificado A1 do médico
// Retorna CPFValidationResult para indicar se há divergência de CPF
func (s *CertificateService) UploadCertificate(
	doctorID uuid.UUID,
	pfxBytes []byte,
	password string,
) (*CPFValidationResult, error) {
	// 1. Validar PFX (parsear com pkcs12)
	_, cert, err := pkcs12.Decode(pfxBytes, password)
	if err != nil {
		return nil, ErrInvalidCertificate
	}

	// 2. Verificar validade temporal
	now := time.Now()
	if now.After(cert.NotAfter) {
		return nil, ErrCertificateExpired
	}
	if now.Before(cert.NotBefore) {
		return nil, errors.New("certificate not yet valid")
	}

	// 3. Verificar se é certificado de assinatura digital
	// (KeyUsage deve incluir DigitalSignature)
	if cert.KeyUsage&x509.KeyUsageDigitalSignature == 0 {
		return nil, errors.New("certificate does not have digital signature capability")
	}

	// 4. Verificar cadeia ICP-Brasil (básico)
	// Nota: Verificação completa da cadeia seria mais complexa
	// Por enquanto, apenas verificamos se o certificado tem issuer válido
	if !s.isICPBrasilCertificate(cert) {
		return nil, ErrNotICPBrasil
	}

	// 5. Criptografar PFX e senha (AES-256-GCM)
	pfxBase64 := base64.StdEncoding.EncodeToString(pfxBytes)
	encryptedPFX, err := internalCrypto.EncryptString(pfxBase64, s.encryptionKey)
	if err != nil {
		return nil, err
	}

	encryptedPassword, err := internalCrypto.EncryptString(password, s.encryptionKey)
	if err != nil {
		return nil, err
	}

	// 6. Extrair CPF e nome do titular do certificado
	certificateCPF := s.extractCPFFromCertificate(cert)
	certificateName := s.extractNameFromCertificate(cert)

	// 7. Buscar usuário
	var user models.User
	if err := s.db.First(&user, doctorID).Error; err != nil {
		return nil, err
	}

	// 8. Validar CPF (se ambos disponíveis)
	validationResult := &CPFValidationResult{
		Valid:             true,
		CertificateCPF:    certificateCPF,
		RequiresConfirmation: false,
	}

	// Normalizar CPFs (remover pontuação para comparação)
	normalizedCertCPF := normalizeCPF(certificateCPF)

	if user.CPF != nil && *user.CPF != "" {
		normalizedUserCPF := normalizeCPF(*user.CPF)
		validationResult.UserCPF = *user.CPF

		// Comparar CPFs normalizados
		if normalizedUserCPF != "" && normalizedCertCPF != "" {
			if normalizedUserCPF != normalizedCertCPF {
				validationResult.Valid = false
				validationResult.RequiresConfirmation = true

				// Desativar certificado por divergência de CPF
				user.CertificateActive = false

				// Retornar erro para que o handler possa tratar
				return validationResult, ErrCPFMismatch
			}
		}
	}

	// 9. Salvar certificado no User
	user.CertificatePFX = &encryptedPFX
	user.CertificatePassword = &encryptedPassword
	user.CertificateExpiry = &cert.NotAfter
	serialStr := cert.SerialNumber.String()
	user.CertificateSerial = &serialStr

	if certificateCPF != "" {
		user.CertificateCPF = &certificateCPF
	}
	if certificateName != "" {
		user.CertificateName = &certificateName
	}

	// Ativar certificado se válido
	user.CertificateActive = true

	if err := s.db.Save(&user).Error; err != nil {
		return nil, err
	}

	return validationResult, nil
}

// GetActiveCertificate retorna certificado do médico se válido
func (s *CertificateService) GetActiveCertificate(
	doctorID uuid.UUID,
) (*x509.Certificate, crypto.PrivateKey, error) {
	var user models.User
	if err := s.db.Select(
		"certificate_pfx, certificate_password, certificate_expiry, certificate_active",
	).First(&user, doctorID).Error; err != nil {
		return nil, nil, err
	}

	// Verificar se tem certificado
	if user.CertificatePFX == nil || user.CertificatePassword == nil {
		return nil, nil, ErrNoCertificate
	}

	// Verificar se certificado está ativo
	if !user.CertificateActive {
		return nil, nil, errors.New("certificado desativado")
	}

	// Verificar validade
	if user.CertificateExpiry != nil && time.Now().After(*user.CertificateExpiry) {
		return nil, nil, ErrCertificateExpiredUser
	}

	// Descriptografar PFX
	pfxBase64, err := internalCrypto.DecryptString(*user.CertificatePFX, s.encryptionKey)
	if err != nil {
		return nil, nil, err
	}

	pfxBytes, err := base64.StdEncoding.DecodeString(pfxBase64)
	if err != nil {
		return nil, nil, err
	}

	// Descriptografar senha
	password, err := internalCrypto.DecryptString(*user.CertificatePassword, s.encryptionKey)
	if err != nil {
		return nil, nil, err
	}

	// Parsear certificado
	privateKey, cert, err := pkcs12.Decode(pfxBytes, password)
	if err != nil {
		return nil, nil, ErrInvalidCertificate
	}

	// Verificar validade novamente (double-check)
	if time.Now().After(cert.NotAfter) {
		return nil, nil, ErrCertificateExpired
	}

	return cert, privateKey, nil
}

// isICPBrasilCertificate verifica se o certificado é ICP-Brasil
// Verificação básica: checa se o issuer contém "ICP-Brasil" ou CAs conhecidas
func (s *CertificateService) isICPBrasilCertificate(cert *x509.Certificate) bool {
	// Lista de Autoridades Certificadoras ICP-Brasil oficiais
	// Ref: https://www.gov.br/iti/pt-br/assuntos/repositorio
	icpBrasilIssuers := []string{
		"AC Raiz",
		"ICP-Brasil",
		"Autoridade Certificadora",
		"Certisign",
		"Serasa",
		"Serpro",
		"Caixa",
		"VALID",
		"Soluti",
		"SafeWeb",
		"Fenacon",
		"Sincor",
		"ACBR",
		"Prodest",
		"Digitalsign",
		"AC SERPRO",
		"AC SERASA",
		"AC CERTISIGN",
		"AC VALID",
		"AC SOLUTI",
	}

	issuerStr := cert.Issuer.String()

	for _, knownIssuer := range icpBrasilIssuers {
		if contains(issuerStr, knownIssuer) {
			return true
		}
	}

	return false
}

// contains verifica se uma string contém outra (case-insensitive)
func contains(str, substr string) bool {
	str = toLower(str)
	substr = toLower(substr)
	for i := 0; i <= len(str)-len(substr); i++ {
		if str[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func toLower(s string) string {
	b := make([]byte, len(s))
	for i := range s {
		c := s[i]
		if 'A' <= c && c <= 'Z' {
			c += 'a' - 'A'
		}
		b[i] = c
	}
	return string(b)
}

// extractCPFFromCertificate extrai o CPF do titular do certificado ICP-Brasil
// O CPF pode estar em vários lugares no certificado
func (s *CertificateService) extractCPFFromCertificate(cert *x509.Certificate) string {
	// 1. Tentar extrair do SerialNumber do Subject
	// Formato comum: "CPF:12345678900" ou apenas "12345678900"
	if cert.Subject.SerialNumber != "" {
		serial := cert.Subject.SerialNumber

		// Remove prefixo "CPF:" se existir
		if len(serial) >= 4 && toLower(serial[:4]) == "cpf:" {
			serial = serial[4:]
		}

		// Verificar se tem 11 dígitos numéricos
		if cpf := extractDigits(serial); len(cpf) == 11 {
			return formatCPF(cpf)
		}
	}

	// 2. Tentar extrair do CommonName se tiver formato "Nome:CPF"
	if cert.Subject.CommonName != "" {
		commonName := cert.Subject.CommonName
		// Procurar por ":" e tentar extrair CPF após os dois pontos
		if idx := indexByte(commonName, ':'); idx != -1 {
			afterColon := commonName[idx+1:]
			if cpf := extractDigits(afterColon); len(cpf) == 11 {
				return formatCPF(cpf)
			}
		}
	}

	// 3. Percorrer todos os campos do Subject procurando CPF
	for _, name := range cert.Subject.Names {
		oidStr := name.Type.String()

		// OID 2.5.4.5 é serialNumber (padrão X.509)
		// OID 2.16.76.1.3.1 é CPF específico ICP-Brasil
		// OID 2.16.76.1.3.9 também pode conter CPF
		if oidStr == "2.5.4.5" || oidStr == "2.16.76.1.3.1" || oidStr == "2.16.76.1.3.9" {
			if str, ok := name.Value.(string); ok {
				// Remove prefixo "CPF:" se existir
				if len(str) >= 4 && toLower(str[:4]) == "cpf:" {
					str = str[4:]
				}

				// Extrair apenas dígitos
				if cpf := extractDigits(str); len(cpf) == 11 {
					return formatCPF(cpf)
				}
			}
		}
	}

	return ""
}

// extractDigits extrai apenas os dígitos de uma string
func extractDigits(s string) string {
	digits := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			digits += string(s[i])
		}
	}
	return digits
}

// extractNameFromCertificate extrai o nome do titular do certificado
func (s *CertificateService) extractNameFromCertificate(cert *x509.Certificate) string {
	// O nome está no CommonName
	if cert.Subject.CommonName != "" {
		// Certificados ICP-Brasil geralmente têm o formato:
		// "Nome do Titular:CPF" ou apenas "Nome do Titular"
		name := cert.Subject.CommonName

		// Remover CPF se estiver junto
		if idx := indexByte(name, ':'); idx != -1 {
			name = name[:idx]
		}

		return name
	}

	return ""
}

// formatCPF formata CPF no padrão 123.456.789-00
func formatCPF(cpf string) string {
	if len(cpf) != 11 {
		return cpf
	}
	return cpf[:3] + "." + cpf[3:6] + "." + cpf[6:9] + "-" + cpf[9:]
}

// indexByte retorna o índice do primeiro byte em s que é igual a c, ou -1 se não encontrado
func indexByte(s string, c byte) int {
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			return i
		}
	}
	return -1
}

// normalizeCPF remove pontuação do CPF para comparação
// "123.456.789-00" -> "12345678900"
func normalizeCPF(cpf string) string {
	normalized := ""
	for i := 0; i < len(cpf); i++ {
		c := cpf[i]
		if c >= '0' && c <= '9' {
			normalized += string(c)
		}
	}
	return normalized
}
