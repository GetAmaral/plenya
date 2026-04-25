package services

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
)

// CampaignService — CRUD de campanhas de marketing + helpers de URL/QR.
type CampaignService struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewCampaignService(db *gorm.DB, cfg *config.Config) *CampaignService {
	return &CampaignService{db: db, cfg: cfg}
}

// ============================================================
// DTOs
// ============================================================

// CreateCampaignInput — payload de criação.
type CreateCampaignInput struct {
	Name        string  `json:"name"        validate:"required,min=2,max=160"`
	Slug        string  `json:"slug"        validate:"omitempty,max=120"`
	Description *string `json:"description,omitempty"`
	LandingPath string  `json:"landingPath" validate:"required,startswith=/"`
	UTMSource   string  `json:"utmSource"   validate:"required,max=80"`
	UTMMedium   string  `json:"utmMedium"   validate:"required,max=80"`
	UTMCampaign string  `json:"utmCampaign" validate:"omitempty,max=120"`
	UTMTerm     *string `json:"utmTerm,omitempty" validate:"omitempty,max=120"`
}

// UpdateCampaignInput — todos campos opcionais.
type UpdateCampaignInput struct {
	Name        *string                `json:"name,omitempty"        validate:"omitempty,min=2,max=160"`
	Slug        *string                `json:"slug,omitempty"        validate:"omitempty,max=120"`
	Description *string                `json:"description,omitempty"`
	LandingPath *string                `json:"landingPath,omitempty" validate:"omitempty,startswith=/"`
	UTMSource   *string                `json:"utmSource,omitempty"   validate:"omitempty,max=80"`
	UTMMedium   *string                `json:"utmMedium,omitempty"   validate:"omitempty,max=80"`
	UTMCampaign *string                `json:"utmCampaign,omitempty" validate:"omitempty,max=120"`
	UTMTerm     *string                `json:"utmTerm,omitempty"     validate:"omitempty,max=120"`
	Status      *models.CampaignStatus `json:"status,omitempty"      validate:"omitempty,oneof=active archived"`
}

// CampaignStats — métricas agregadas por campanha (usado na listagem do CRM).
type CampaignStats struct {
	SessionsCount int64 `json:"sessionsCount"`
	LeadsCount    int64 `json:"leadsCount"`
}

// CampaignDTO — projeção pública da campanha + URL final + métricas opcionais.
type CampaignDTO struct {
	models.Campaign
	URL   string         `json:"url"`
	Stats *CampaignStats `json:"stats,omitempty"`
}

// ============================================================
// Métodos públicos
// ============================================================

// List devolve todas as campanhas (active+archived) ordenadas por mais recente.
// Inclui métricas agregadas (sessions/leads) por utm_campaign.
func (s *CampaignService) List(includeArchived bool) ([]CampaignDTO, error) {
	q := s.db.Model(&models.Campaign{}).Order("created_at DESC")
	if !includeArchived {
		q = q.Where("status = ?", models.CampaignStatusActive)
	}
	var rows []models.Campaign
	if err := q.Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]CampaignDTO, 0, len(rows))
	for _, c := range rows {
		stats, _ := s.statsFor(c.UTMCampaign)
		out = append(out, CampaignDTO{Campaign: c, URL: s.BuildURL(&c), Stats: stats})
	}
	return out, nil
}

// GetByID retorna uma campanha + URL + stats.
func (s *CampaignService) GetByID(id uuid.UUID) (*CampaignDTO, error) {
	var c models.Campaign
	if err := s.db.First(&c, "id = ?", id).Error; err != nil {
		return nil, err
	}
	stats, _ := s.statsFor(c.UTMCampaign)
	return &CampaignDTO{Campaign: c, URL: s.BuildURL(&c), Stats: stats}, nil
}

// Create cria uma campanha. Slug e utm_campaign default para Slugify(Name) se vazios.
func (s *CampaignService) Create(in CreateCampaignInput, createdBy *uuid.UUID) (*CampaignDTO, error) {
	c := models.Campaign{
		Name:            strings.TrimSpace(in.Name),
		Slug:            in.Slug,
		Description:     normalizeOptional(in.Description),
		LandingPath:     strings.TrimSpace(in.LandingPath),
		UTMSource:       strings.TrimSpace(in.UTMSource),
		UTMMedium:       strings.TrimSpace(in.UTMMedium),
		UTMCampaign:     strings.TrimSpace(in.UTMCampaign),
		UTMTerm:         normalizeOptional(in.UTMTerm),
		Status:          models.CampaignStatusActive,
		CreatedByUserID: createdBy,
	}
	if err := s.db.Create(&c).Error; err != nil {
		if isUniqueViolation(err) {
			return nil, fmt.Errorf("já existe campanha com slug '%s'", c.Slug)
		}
		return nil, err
	}
	return &CampaignDTO{Campaign: c, URL: s.BuildURL(&c)}, nil
}

// Update aplica patch parcial.
func (s *CampaignService) Update(id uuid.UUID, in UpdateCampaignInput) (*CampaignDTO, error) {
	var c models.Campaign
	if err := s.db.First(&c, "id = ?", id).Error; err != nil {
		return nil, err
	}
	if in.Name != nil {
		c.Name = strings.TrimSpace(*in.Name)
	}
	if in.Slug != nil {
		c.Slug = *in.Slug
	}
	if in.Description != nil {
		c.Description = normalizeOptional(in.Description)
	}
	if in.LandingPath != nil {
		c.LandingPath = strings.TrimSpace(*in.LandingPath)
	}
	if in.UTMSource != nil {
		c.UTMSource = strings.TrimSpace(*in.UTMSource)
	}
	if in.UTMMedium != nil {
		c.UTMMedium = strings.TrimSpace(*in.UTMMedium)
	}
	if in.UTMCampaign != nil {
		c.UTMCampaign = strings.TrimSpace(*in.UTMCampaign)
	}
	if in.UTMTerm != nil {
		c.UTMTerm = normalizeOptional(in.UTMTerm)
	}
	if in.Status != nil {
		c.Status = *in.Status
	}
	if err := s.db.Save(&c).Error; err != nil {
		if isUniqueViolation(err) {
			return nil, fmt.Errorf("já existe campanha com slug '%s'", c.Slug)
		}
		return nil, err
	}
	stats, _ := s.statsFor(c.UTMCampaign)
	return &CampaignDTO{Campaign: c, URL: s.BuildURL(&c), Stats: stats}, nil
}

// Archive marca como arquivada (soft — preserva histórico de atribuição).
func (s *CampaignService) Archive(id uuid.UUID) error {
	return s.db.Model(&models.Campaign{}).Where("id = ?", id).
		Update("status", models.CampaignStatusArchived).Error
}

// Delete remove de fato (cuidado — usar Archive na maioria dos casos).
func (s *CampaignService) Delete(id uuid.UUID) error {
	return s.db.Delete(&models.Campaign{}, "id = ?", id).Error
}

// BuildURL monta a URL pública final com os UTMs preenchidos.
func (s *CampaignService) BuildURL(c *models.Campaign) string {
	base := strings.TrimRight(s.cfg.Site.PublicURL, "/")
	if base == "" {
		base = "https://plenyasaude.com.br"
	}
	q := url.Values{}
	q.Set("utm_source", c.UTMSource)
	q.Set("utm_medium", c.UTMMedium)
	q.Set("utm_campaign", c.UTMCampaign)
	if c.UTMTerm != nil && *c.UTMTerm != "" {
		q.Set("utm_term", *c.UTMTerm)
	}
	return fmt.Sprintf("%s%s?%s", base, c.LandingPath, q.Encode())
}

// QRCodePNG devolve o PNG (bytes) do QR code da URL final.
func (s *CampaignService) QRCodePNG(id uuid.UUID, sizePx int) ([]byte, string, error) {
	if sizePx < 128 {
		sizePx = 128
	}
	if sizePx > 1024 {
		sizePx = 1024
	}
	var c models.Campaign
	if err := s.db.First(&c, "id = ?", id).Error; err != nil {
		return nil, "", err
	}
	full := s.BuildURL(&c)
	png, err := qrcode.Encode(full, qrcode.Medium, sizePx)
	if err != nil {
		return nil, "", fmt.Errorf("qr encode: %w", err)
	}
	filename := fmt.Sprintf("plenya-campanha-%s.png", c.Slug)
	return png, filename, nil
}

// ============================================================
// Helpers
// ============================================================

func (s *CampaignService) statsFor(utmCampaign string) (*CampaignStats, error) {
	if utmCampaign == "" {
		return &CampaignStats{}, nil
	}
	var sessions int64
	if err := s.db.Model(&models.AnonymousScoreSession{}).
		Where("utm_campaign = ?", utmCampaign).Count(&sessions).Error; err != nil {
		return nil, err
	}
	var leads int64
	if err := s.db.Model(&models.Lead{}).
		Where("utm_campaign = ?", utmCampaign).Count(&leads).Error; err != nil {
		return nil, err
	}
	return &CampaignStats{SessionsCount: sessions, LeadsCount: leads}, nil
}

func normalizeOptional(p *string) *string {
	if p == nil {
		return nil
	}
	s := strings.TrimSpace(*p)
	if s == "" {
		return nil
	}
	return &s
}

// Sentinel para handler distinguir 404 de outros erros.
var ErrCampaignNotFound = errors.New("campaign not found")
