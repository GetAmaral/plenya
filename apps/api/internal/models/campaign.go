package models

import (
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CampaignStatus indica o estado operacional de uma campanha de marketing.
type CampaignStatus string

const (
	CampaignStatusActive   CampaignStatus = "active"
	CampaignStatusArchived CampaignStatus = "archived"
)

// Campaign representa uma ação de marketing rastreável via UTM.
//
// Cada campanha gera uma URL final tipo
// https://plenyasaude.com.br{LandingPath}?utm_source={UTMSource}&utm_medium={UTMMedium}&utm_campaign={UTMCampaign}[&utm_term={UTMTerm}]
// que pode ser distribuída diretamente, virar QR code, ou compor links de bio/anúncio.
//
// Quando o usuário aterra na URL, o site captura os utm_* e propaga até a sessão do
// Escore (AnonymousScoreSession) e, no claim, até o Lead — fechando o ciclo de atribuição.
//
// @Description Campanha de marketing rastreável (UTM + QR code)
type Campaign struct {
	// @example 550e8400-e29b-41d4-a716-446655440000
	ID uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	// Nome legível da campanha
	// @example Stories Instagram — Janeiro 2026
	Name string `gorm:"type:varchar(160);not null" json:"name" validate:"required,min=2,max=160"`

	// Slug auto-gerado a partir do nome (ou definido manualmente). Usado como utm_campaign default.
	// @example stories-instagram-janeiro-2026
	Slug string `gorm:"type:varchar(120);uniqueIndex;not null" json:"slug" validate:"required,min=2,max=120"`

	// Descrição interna (briefing, audiência, etc.)
	Description *string `gorm:"type:text" json:"description,omitempty"`

	// Caminho do site público para o qual o link aponta (sem domínio nem query).
	// @example /escore-plenya/painel
	LandingPath string `gorm:"type:varchar(255);not null;default:'/escore-plenya/painel'" json:"landingPath" validate:"required,startswith=/"`

	// utm_source — origem (ex: instagram, whatsapp, qr-cartao-cca, podcast-x)
	// @example instagram
	UTMSource string `gorm:"type:varchar(80);not null;index:idx_campaigns_utm_source" json:"utmSource" validate:"required,max=80"`

	// utm_medium — meio (ex: stories, post, qr, email, paid-ads, organic)
	// @example stories
	UTMMedium string `gorm:"type:varchar(80);not null" json:"utmMedium" validate:"required,max=80"`

	// utm_campaign — identificador da campanha (default = slug)
	// @example stories-instagram-janeiro-2026
	UTMCampaign string `gorm:"type:varchar(120);not null" json:"utmCampaign" validate:"required,max=120"`

	// utm_term — termo opcional (variação A/B, criativo, etc.)
	UTMTerm *string `gorm:"type:varchar(120)" json:"utmTerm,omitempty"`

	// Status operacional
	// @enum active,archived
	// @example active
	Status CampaignStatus `gorm:"type:varchar(20);not null;default:'active';index:idx_campaigns_status" json:"status" validate:"required,oneof=active archived"`

	// Quem criou a campanha
	CreatedByUserID *uuid.UUID `gorm:"type:uuid;index:idx_campaigns_created_by" json:"createdByUserId,omitempty"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	CreatedBy *User `gorm:"foreignKey:CreatedByUserID;constraint:OnDelete:SET NULL" json:"createdBy,omitempty"`
}

// TableName especifica o nome da tabela.
func (Campaign) TableName() string { return "campaigns" }

// BeforeCreate hook: UUID v7, defaults, slug auto-gerado se vazio.
func (c *Campaign) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.Must(uuid.NewV7())
	}
	if c.Status == "" {
		c.Status = CampaignStatusActive
	}
	if c.Slug == "" {
		c.Slug = Slugify(c.Name)
	}
	if c.UTMCampaign == "" {
		c.UTMCampaign = c.Slug
	}
	if c.LandingPath == "" {
		c.LandingPath = "/escore-plenya/painel"
	}
	return nil
}

// BeforeSave hook: normaliza slug e utm_campaign ao salvar.
func (c *Campaign) BeforeSave(tx *gorm.DB) error {
	c.Slug = Slugify(c.Slug)
	c.UTMCampaign = strings.TrimSpace(c.UTMCampaign)
	c.UTMSource = strings.TrimSpace(c.UTMSource)
	c.UTMMedium = strings.TrimSpace(c.UTMMedium)
	return nil
}

var slugifyRe = regexp.MustCompile(`[^a-z0-9]+`)

// Slugify converte um texto em slug ASCII lowercase com hífens.
func Slugify(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	// Remove acentos comuns PT-BR (subset prático — não cobre todos os caracteres Unicode).
	r := strings.NewReplacer(
		"á", "a", "à", "a", "ã", "a", "â", "a", "ä", "a",
		"é", "e", "è", "e", "ê", "e", "ë", "e",
		"í", "i", "ì", "i", "î", "i", "ï", "i",
		"ó", "o", "ò", "o", "õ", "o", "ô", "o", "ö", "o",
		"ú", "u", "ù", "u", "û", "u", "ü", "u",
		"ç", "c", "ñ", "n",
	)
	s = r.Replace(s)
	s = slugifyRe.ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	return s
}
