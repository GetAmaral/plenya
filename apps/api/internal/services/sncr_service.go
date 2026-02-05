package services

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// SNCRProvider é a interface para provedores SNCR
type SNCRProvider interface {
	RequestPrescriptionNumber(prescription *models.Prescription) (string, error)
	MarkAsUsed(sncrNumber string) error
	GetStatus(sncrNumber string) (string, error)
}

// ============================================================================
// STUB Provider (Fake - Usar até jun/2026)
// ============================================================================

type SNCRStubProvider struct {
	db *gorm.DB
}

func NewSNCRStubProvider(db *gorm.DB) *SNCRStubProvider {
	return &SNCRStubProvider{db: db}
}

func (s *SNCRStubProvider) RequestPrescriptionNumber(prescription *models.Prescription) (string, error) {
	// Gera número FAKE: BR-STUB-2026-00000001
	var count int64
	s.db.Model(&models.Prescription{}).
		Where("sncr_number LIKE ?", "BR-STUB-%").
		Count(&count)

	number := fmt.Sprintf("BR-STUB-%d-%08d", time.Now().Year(), count+1)
	return number, nil
}

func (s *SNCRStubProvider) MarkAsUsed(sncrNumber string) error {
	status := "used"
	return s.db.Model(&models.Prescription{}).
		Where("sncr_number = ?", sncrNumber).
		Updates(map[string]interface{}{
			"is_used":     true,
			"sncr_status": status,
			"dispensed_at": time.Now(),
		}).Error
}

func (s *SNCRStubProvider) GetStatus(sncrNumber string) (string, error) {
	var prescription models.Prescription
	err := s.db.Select("sncr_status, is_used").
		Where("sncr_number = ?", sncrNumber).
		First(&prescription).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("SNCR number not found")
		}
		return "", err
	}

	if prescription.SNCRStatus == nil {
		return "active", nil
	}
	return *prescription.SNCRStatus, nil
}

// ============================================================================
// PRODUCTION Provider (API ANVISA - jun/2026+)
// ============================================================================

type SNCRProductionProvider struct {
	apiURL     string
	apiKey     string
	httpClient *http.Client
}

func NewSNCRProductionProvider(apiURL, apiKey string) *SNCRProductionProvider {
	return &SNCRProductionProvider{
		apiURL:     apiURL,
		apiKey:     apiKey,
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (s *SNCRProductionProvider) RequestPrescriptionNumber(prescription *models.Prescription) (string, error) {
	// TODO: Implementar quando ANVISA publicar API (mai/2026)
	// POST https://sncr.anvisa.gov.br/api/v1/prescriptions/register
	// Headers: Authorization: Bearer {apiKey}
	// Body: {
	//   "medicationName": "...",
	//   "category": "c1",
	//   "patientCPF": "...",
	//   "doctorCRM": "...",
	//   ...
	// }
	// Response: { "sncrNumber": "BR-2026-XXXXXXXX" }

	return "", errors.New("SNCR Production API not available yet (expected: mai/2026)")
}

func (s *SNCRProductionProvider) MarkAsUsed(sncrNumber string) error {
	// TODO: Implementar quando ANVISA publicar API
	// PUT https://sncr.anvisa.gov.br/api/v1/prescriptions/{sncrNumber}/dispense
	return errors.New("SNCR Production API not available yet")
}

func (s *SNCRProductionProvider) GetStatus(sncrNumber string) (string, error) {
	// TODO: Implementar quando ANVISA publicar API
	// GET https://sncr.anvisa.gov.br/api/v1/prescriptions/{sncrNumber}/status
	return "", errors.New("SNCR Production API not available yet")
}

// ============================================================================
// SNCR Service Principal
// ============================================================================

type SNCRConfig struct {
	Enabled        bool
	ProductionMode bool
	APIURL         string
	APIKey         string
}

type SNCRService struct {
	provider SNCRProvider
	enabled  bool
}

func NewSNCRService(config SNCRConfig, db *gorm.DB) *SNCRService {
	var provider SNCRProvider

	if config.ProductionMode {
		provider = NewSNCRProductionProvider(config.APIURL, config.APIKey)
	} else {
		provider = NewSNCRStubProvider(db)
	}

	return &SNCRService{
		provider: provider,
		enabled:  config.Enabled,
	}
}

// RequestNumber solicita número SNCR para prescrição controlada
func (s *SNCRService) RequestNumber(prescription *models.Prescription) (string, error) {
	if !s.enabled {
		return "", nil // SNCR desabilitado
	}

	// Verificar se ALGUM medicamento é controlado (C1 ou C5)
	hasControlledMedication := false
	for _, med := range prescription.Medications {
		if med.Category == models.MedCategoryC1 || med.Category == models.MedCategoryC5 {
			hasControlledMedication = true
			break
		}
	}

	if !hasControlledMedication {
		return "", nil // Não precisa de SNCR
	}

	return s.provider.RequestPrescriptionNumber(prescription)
}

// MarkAsUsed marca prescrição como dispensada
func (s *SNCRService) MarkAsUsed(sncrNumber string) error {
	if !s.enabled || sncrNumber == "" {
		return nil
	}
	return s.provider.MarkAsUsed(sncrNumber)
}

// GetStatus consulta status no SNCR
func (s *SNCRService) GetStatus(sncrNumber string) (string, error) {
	if !s.enabled || sncrNumber == "" {
		return "active", nil
	}
	return s.provider.GetStatus(sncrNumber)
}
