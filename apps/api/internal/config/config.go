package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	Security SecurityConfig
	OAuth    OAuthConfig
	SNCR     SNCRConfig
	Claude   ClaudeConfig
	OpenAI   OpenAIConfig
	VoyageAI VoyageAIConfig
	Email    EmailConfig
	WhatsApp WhatsAppConfig
	CRM      CRMConfig
	Site     SiteConfig
	Dev      DevConfig
}

// EmailConfig — Provider escolhe a implementação. Quando vazio (ou Provider="smtp" com SMTPHost vazio),
// EmailService loga no stdout (dev fallback).
type EmailConfig struct {
	Provider     string // "smtp" | "resend"; default: "smtp"
	ResendAPIKey string // RESEND_API_KEY (necessário se Provider="resend")
	SMTPHost     string
	SMTPPort     string
	SMTPUser     string
	SMTPPass     string
	FromAddress  string
	FromName     string // nome amigável no header From (ex: "Plenya")
}

// WhatsAppConfig — credenciais Meta Cloud API. Quando PhoneNumberID vazio,
// WhatsAppService loga no stdout (dev fallback).
type WhatsAppConfig struct {
	AppSecret          string // WHATSAPP_APP_SECRET — usado pra validar HMAC do webhook
	AccessToken        string // WHATSAPP_ACCESS_TOKEN — Permanent Access Token (System User)
	PhoneNumberID      string // WHATSAPP_PHONE_NUMBER_ID — número Plenya WhatsApp Business
	WebhookVerifyToken string // WHATSAPP_WEBHOOK_VERIFY_TOKEN — secret aleatório do challenge inicial
	TemplateMagicLink  string // WHATSAPP_TEMPLATE_MAGIC_LINK — nome do template (default: "magic_link")
	TemplateLeadAlert  string // WHATSAPP_TEMPLATE_LEAD_ALERT — notificação interna (default: "lead_alert")
	GraphAPIVersion    string // WHATSAPP_GRAPH_API_VERSION — ex: "v18.0" (default)
}

// CRMConfig — config operacional do CRM (Phase 2).
type CRMConfig struct {
	AdminURL          string   // CRM_ADMIN_URL — URL base do EMR (ex: https://app.plenyasaude.com.br)
	LeadNotifyUserIDs []string // CRM_LEAD_NOTIFY_USER_IDS (csv) — IDs hardcoded de quem recebe notif (fallback)
}

// SiteConfig URLs do site público — usado para montar links em emails transacionais
type SiteConfig struct {
	PublicURL string // ex: https://plenyasaude.com.br
}

type ServerConfig struct {
	Port        string
	Environment string
	CORSOrigin  string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type JWTConfig struct {
	Secret        string
	AccessExpiry  string
	RefreshExpiry string
}

type SecurityConfig struct {
	EncryptionKey   string
	RateLimitReqs   int
	RateLimitWindow int
}

type OAuthConfig struct {
	GoogleClientID     string
	GoogleClientSecret string
	AppleClientID      string
	AppleTeamID        string
	AppleKeyID         string
	ApplePrivateKey    string
}

type SNCRConfig struct {
	Enabled        bool
	ProductionMode bool
	APIURL         string
	APIKey         string
}

type ClaudeConfig struct {
	APIKey string
	Model  string
}

type OpenAIConfig struct {
	APIKey         string
	EmbeddingModel string
	APIURL         string
}

type VoyageAIConfig struct {
	APIKey string
	Model  string // Default: voyage-multilingual-2
	APIURL string // Default: https://api.voyageai.com/v1
}

type DevConfig struct {
	BypassAuth  bool      // DEV_BYPASS_AUTH
	AdminUserID uuid.UUID // Populated at runtime after DB init
	AdminEmail  string    // Populated at runtime after DB init
	AdminRoles  []string  // Populated at runtime after DB init
}

// Load carrega as configurações do ambiente
func Load() (*Config, error) {
	// Tentar carregar .env (ignora erro se não existir em produção)
	_ = godotenv.Load()

	cfg := &Config{
		Server: ServerConfig{
			Port:        getEnv("PORT", "3001"),
			Environment: getEnv("ENVIRONMENT", "development"),
			CORSOrigin:  getEnv("CORS_ORIGIN", "http://localhost:3000"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "plenya_user"),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", "plenya_db"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
		JWT: JWTConfig{
			Secret:        getEnv("JWT_SECRET", ""),
			AccessExpiry:  getEnv("JWT_ACCESS_EXPIRY", "15m"),
			RefreshExpiry: getEnv("JWT_REFRESH_EXPIRY", "168h"),
		},
		Security: SecurityConfig{
			EncryptionKey:   getEnv("ENCRYPTION_KEY", ""),
			RateLimitReqs:   getEnvAsInt("RATE_LIMIT_REQUESTS", 100),
			RateLimitWindow: getEnvAsInt("RATE_LIMIT_WINDOW", 60),
		},
		OAuth: OAuthConfig{
			GoogleClientID:     getEnv("OAUTH_GOOGLE_CLIENT_ID", ""),
			GoogleClientSecret: getEnv("OAUTH_GOOGLE_CLIENT_SECRET", ""),
			AppleClientID:      getEnv("OAUTH_APPLE_CLIENT_ID", ""),
			AppleTeamID:        getEnv("OAUTH_APPLE_TEAM_ID", ""),
			AppleKeyID:         getEnv("OAUTH_APPLE_KEY_ID", ""),
			ApplePrivateKey:    getEnv("OAUTH_APPLE_PRIVATE_KEY", ""),
		},
		SNCR: SNCRConfig{
			Enabled:        getEnvAsBool("SNCR_ENABLED", true),
			ProductionMode: getEnvAsBool("SNCR_PRODUCTION_MODE", false),
			APIURL:         getEnv("SNCR_API_URL", "https://sncr.anvisa.gov.br/api/v1"),
			APIKey:         getEnv("SNCR_API_KEY", ""),
		},
		Claude: ClaudeConfig{
			APIKey: getEnv("CLAUDE_API_KEY", ""),
			Model:  getEnv("CLAUDE_MODEL", "claude-3-5-haiku-20241022"),
		},
		OpenAI: OpenAIConfig{
			APIKey:         getEnv("OPENAI_API_KEY", ""),
			EmbeddingModel: getEnv("OPENAI_EMBEDDING_MODEL", "text-embedding-3-large"),
			APIURL:         getEnv("OPENAI_API_URL", "https://api.openai.com/v1"),
		},
		VoyageAI: VoyageAIConfig{
			APIKey: getEnv("VOYAGE_API_KEY", ""),
			Model:  getEnv("VOYAGE_MODEL", "voyage-multilingual-2"),
			APIURL: getEnv("VOYAGE_API_URL", "https://api.voyageai.com/v1"),
		},
		Email: EmailConfig{
			Provider:     getEnv("EMAIL_PROVIDER", "smtp"),
			ResendAPIKey: getEnv("RESEND_API_KEY", ""),
			SMTPHost:     getEnv("SMTP_HOST", ""),
			SMTPPort:     getEnv("SMTP_PORT", "587"),
			SMTPUser:     getEnv("SMTP_USER", ""),
			SMTPPass:     getEnv("SMTP_PASSWORD", ""),
			FromAddress:  getEnv("EMAIL_FROM", "no-reply@plenyasaude.com.br"),
			FromName:     getEnv("EMAIL_FROM_NAME", "Plenya"),
		},
		WhatsApp: WhatsAppConfig{
			AppSecret:          getEnv("WHATSAPP_APP_SECRET", ""),
			AccessToken:        getEnv("WHATSAPP_ACCESS_TOKEN", ""),
			PhoneNumberID:      getEnv("WHATSAPP_PHONE_NUMBER_ID", ""),
			WebhookVerifyToken: getEnv("WHATSAPP_WEBHOOK_VERIFY_TOKEN", ""),
			TemplateMagicLink:  getEnv("WHATSAPP_TEMPLATE_MAGIC_LINK", "magic_link"),
			TemplateLeadAlert:  getEnv("WHATSAPP_TEMPLATE_LEAD_ALERT", "lead_alert"),
			GraphAPIVersion:    getEnv("WHATSAPP_GRAPH_API_VERSION", "v18.0"),
		},
		CRM: CRMConfig{
			AdminURL:          getEnv("CRM_ADMIN_URL", "http://localhost:3000"),
			LeadNotifyUserIDs: parseCSV(getEnv("CRM_LEAD_NOTIFY_USER_IDS", "")),
		},
		Site: SiteConfig{
			PublicURL: getEnv("SITE_PUBLIC_URL", "http://localhost:3002"),
		},
		Dev: DevConfig{
			BypassAuth: getEnvAsBool("DEV_BYPASS_AUTH", false),
		},
	}

	// Validar configurações críticas
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

// Validate verifica se as configurações obrigatórias estão presentes
func (c *Config) Validate() error {
	if c.Database.Password == "" {
		return fmt.Errorf("DB_PASSWORD is required")
	}

	if c.JWT.Secret == "" {
		return fmt.Errorf("JWT_SECRET is required")
	}

	if c.Security.EncryptionKey == "" {
		return fmt.Errorf("ENCRYPTION_KEY is required")
	}

	return nil
}

// GetDSN retorna a string de conexão do PostgreSQL
func (c *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.Name, c.SSLMode,
	)
}

// getEnv obtém variável de ambiente ou retorna valor padrão
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt obtém variável de ambiente como int ou retorna valor padrão
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// parseCSV separa string CSV em slice, removendo espaços e itens vazios.
func parseCSV(s string) []string {
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

// getEnvAsBool obtém variável de ambiente como bool ou retorna valor padrão
func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}
