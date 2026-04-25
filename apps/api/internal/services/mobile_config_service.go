package services

import (
	"os"
	"strconv"
	"strings"

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

// MobileConfigService produz a config do app. Lê de env vars MOBILE_*
// pra permitir mudança sem deploy (kill switch, pin rotation, feature flags).
type MobileConfigService struct {
	cfg *config.Config
}

// NewMobileConfigService instancia o service.
func NewMobileConfigService(cfg *config.Config) *MobileConfigService {
	return &MobileConfigService{cfg: cfg}
}

// Get retorna a config atual lida das env vars MOBILE_*.
func (s *MobileConfigService) Get() MobileConfig {
	return MobileConfig{
		MinVersion: MobileMinVersion{
			IOS:     envOrDefault("MOBILE_MIN_VERSION_IOS", "0.1.0"),
			Android: envOrDefault("MOBILE_MIN_VERSION_ANDROID", "0.1.0"),
		},
		KillSwitch: MobileKillSwitch{
			Enabled: envBool("MOBILE_KILL_SWITCH"),
			Message: os.Getenv("MOBILE_KILL_SWITCH_MESSAGE"),
		},
		SSLPins: MobileSSLPins{
			Current: os.Getenv("MOBILE_SSL_PIN_CURRENT"),
			Backup:  os.Getenv("MOBILE_SSL_PIN_BACKUP"),
		},
		FeatureFlags: parseFlags(os.Getenv("MOBILE_FEATURE_FLAGS")),
		SupportContact: MobileSupportContact{
			Email:    envOrDefault("MOBILE_SUPPORT_EMAIL", "suporte@plenyasaude.com.br"),
			WhatsApp: os.Getenv("MOBILE_SUPPORT_WHATSAPP"),
		},
	}
}

func envOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func envBool(key string) bool {
	v, _ := strconv.ParseBool(os.Getenv(key))
	return v
}

// parseFlags interpreta MOBILE_FEATURE_FLAGS no formato "flag1=true,flag2=false".
// Defaults são definidos aqui — qualquer flag não declarada na env mantém o
// default (mais seguro do que false implícito).
func parseFlags(raw string) map[string]bool {
	out := map[string]bool{
		"mobile.training":  true,
		"mobile.crm":       true,
		"mobile.scoreEdit": false,
	}
	if raw == "" {
		return out
	}
	for _, kv := range strings.Split(raw, ",") {
		parts := strings.SplitN(strings.TrimSpace(kv), "=", 2)
		if len(parts) != 2 {
			continue
		}
		v, _ := strconv.ParseBool(strings.TrimSpace(parts[1]))
		out[strings.TrimSpace(parts[0])] = v
	}
	return out
}
