package services

import (
	"github.com/plenya/api/internal/config"
)

// MobileMinVersion define a versão mínima suportada por plataforma.
type MobileMinVersion struct {
	IOS     string `json:"ios"`
	Android string `json:"android"`
}

// MobileKillSwitch permite bloquear acesso de todos os apps remotamente.
type MobileKillSwitch struct {
	Enabled bool   `json:"enabled"`
	Message string `json:"message,omitempty"`
}

// MobileSSLPins distribui o pin atual + backup para rotação sem release.
type MobileSSLPins struct {
	Current string `json:"current"`
	Backup  string `json:"backup"`
}

// MobileSupportContact lista canais de suporte exibidos no app.
type MobileSupportContact struct {
	Email    string `json:"email"`
	WhatsApp string `json:"whatsapp,omitempty"`
}

// MobileConfig é a resposta de GET /api/v1/mobile/config.
type MobileConfig struct {
	MinVersion     MobileMinVersion     `json:"minVersion"`
	KillSwitch     MobileKillSwitch     `json:"killSwitch"`
	SSLPins        MobileSSLPins        `json:"sslPins"`
	FeatureFlags   map[string]bool      `json:"featureFlags"`
	SupportContact MobileSupportContact `json:"supportContact"`
}

// MobileConfigService produz a config do app. Hoje serve valores estáticos
// derivados de config; pode evoluir para ler de DB / PostHog no futuro.
type MobileConfigService struct {
	cfg *config.Config
}

// NewMobileConfigService instancia o service.
func NewMobileConfigService(cfg *config.Config) *MobileConfigService {
	return &MobileConfigService{cfg: cfg}
}

// Get retorna a config atual.
func (s *MobileConfigService) Get() MobileConfig {
	return MobileConfig{
		MinVersion: MobileMinVersion{
			IOS:     "0.1.0",
			Android: "0.1.0",
		},
		KillSwitch: MobileKillSwitch{Enabled: false},
		SSLPins: MobileSSLPins{
			Current: "",
			Backup:  "",
		},
		FeatureFlags: map[string]bool{
			"mobile.training":  false,
			"mobile.crm":       false,
			"mobile.scoreEdit": false,
		},
		SupportContact: MobileSupportContact{
			Email: "suporte@plenyasaude.com.br",
		},
	}
}
