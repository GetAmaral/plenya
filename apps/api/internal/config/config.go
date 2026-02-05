package config

import (
	"fmt"
	"os"
	"strconv"

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
			Model:  getEnv("CLAUDE_MODEL", "claude-sonnet-4-5-20250929"),
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

// getEnvAsBool obtém variável de ambiente como bool ou retorna valor padrão
func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}
