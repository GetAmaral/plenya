package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type Config struct {
	Server        ServerConfig
	Database      DatabaseConfig
	JWT           JWTConfig
	Security      SecurityConfig
	OAuth         OAuthConfig
	SNCR          SNCRConfig
	Signature     SignatureConfig
	Claude        ClaudeConfig
	OpenAI        OpenAIConfig
	VoyageAI      VoyageAIConfig
	Email         EmailConfig
	MailIngest    MailIngestConfig
	WhatsApp      WhatsAppConfig
	CRM           CRMConfig
	Site          SiteConfig
	Google        GoogleConfig
	DailyCo       DailyCoConfig
	PatientPortal PatientPortalConfig
	MagicLink     MagicLinkConfig
	Dev           DevConfig
	Turnstile     TurnstileConfig
	ReceptionBot  ReceptionBotConfig
}

// ReceptionBotConfig — recepcionista virtual (Fase 2, modo automático).
// Enabled é o kill switch global do auto-envio: com false, a IA nunca envia sozinha
// (o Copiloto continua funcionando). DefaultMode é o modo de conversas sem override.
type ReceptionBotConfig struct {
	Enabled         bool   // RECEPTION_BOT_ENABLED — liga o auto-envio (default false)
	DefaultMode     string // RECEPTION_BOT_DEFAULT_MODE — off|copilot|auto (default off)
	FallbackMinutes int    // RECEPTION_BOT_FALLBACK_MINUTES — min sem resposta humana antes do bot (default 5)
	MaxMsgsPerHour  int    // RECEPTION_BOT_MAX_MSGS_HOUR — anti-spam por conversa (default 6)
	ConsultDoctorID string // RECEPTION_CONSULT_DOCTOR_ID — médico da Consulta Plenya p/ ofertar horários (vazio = 1º doctor)
}

// TurnstileConfig — M10 — Cloudflare Turnstile CAPTCHA (lead form público).
// Quando Secret vazio em dev, validação é pulada com warning (não bloqueia local).
// Em prod, Secret obrigatório (Validate avisa, mas não bloqueia boot pra evitar
// regressão se ainda não configurado no Coolify).
type TurnstileConfig struct {
	Secret  string // TURNSTILE_SECRET — site secret do Cloudflare
	SiteKey string // TURNSTILE_SITE_KEY — público (informativo, frontend usa NEXT_PUBLIC_)
}

// PatientPortalConfig — área do paciente (minha.plenyasaude.com.br).
// PublicURL é usado pra montar links em emails de convite, magic link, etc.
type PatientPortalConfig struct {
	PublicURL string // PATIENT_PORTAL_URL — ex: https://minha.plenyasaude.com.br
}

// GoogleConfig — credenciais OAuth Google + Calendar API.
//
// SETUP OPERACIONAL (Calendar V1, Bloco H pendente):
//  1. Criar projeto "Plenya EMR" em console.cloud.google.com
//  2. Habilitar Google Calendar API
//  3. OAuth consent screen: External, scope `calendar.events.owned` + `userinfo.email`
//  4. Credentials → OAuth Client ID Web Application
//  5. Authorized redirect URI: <RedirectURL>
//  6. Domain verification em search.google.com/search-console
//
// Diferente de OAuthConfig.GoogleClientID/Secret (esse é pra "Sign in with Google"
// do app móvel/web). GoogleConfig é APENAS pra integração Calendar do médico.
// Mantemos separados pra permitir scopes/projects diferentes (Calendar exige
// CASA security assessment quando >100 users; Sign-in não).
type GoogleConfig struct {
	ClientID     string // GOOGLE_CLIENT_ID
	ClientSecret string // GOOGLE_CLIENT_SECRET
	RedirectURL  string // GOOGLE_REDIRECT_URL — callback completo (deve bater com Console)
}

// DailyCoConfig — credenciais Daily.co pra teleconsulta embedada.
//
// SETUP OPERACIONAL:
//  1. Criar conta em daily.co (free tier: 10k min/mês)
//  2. Dashboard → API Keys → criar key
//  3. Domain: subdomínio escolhido (ex: "plenya" → URLs https://plenya.daily.co/...)
//
// Quando APIKey vazio, DailyCoService retorna error explícito sem crashar
// (caller decide se bloqueia appointment ou cria sem teleconsulta).
type DailyCoConfig struct {
	APIKey string // DAILY_CO_API_KEY
	Domain string // DAILY_CO_DOMAIN — subdomínio (ex: "plenya")
	// WebhookSecret — segredo HMAC (base64) retornado por POST /v1/webhooks ao
	// registrar o webhook do domínio. Usado pra verificar a assinatura
	// X-Webhook-Signature dos eventos recording.*/transcript.*. Vazio = webhook
	// rejeitado em prod (fail-closed); rode cmd/daily-webhook pra obter o segredo.
	WebhookSecret string // DAILY_CO_WEBHOOK_SECRET
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
	// LeadReplyFromAddress — From usado especificamente pra responder Leads (Bloco 3 CRM).
	// Default: contato@plenyasaude.com.br. Diferente de FromAddress (transacional, geralmente
	// noreply@) porque respostas humanas precisam cair em caixa monitorada — fecha o ciclo
	// pelo worker IMAP IDLE, que ingesta inbound e cria LeadActivity message_received.
	LeadReplyFromAddress string
}

// MailIngestConfig — config do worker IMAP que ingesta emails inbound (Stalwart self-hosted)
// e os espelha como LeadActivity. Quando Enabled=false, o worker faz no-op.
type MailIngestConfig struct {
	Enabled         bool     // MAIL_INGEST_ENABLED
	IMAPHost        string   // MAIL_INGEST_IMAP_HOST (ex: mail.plenyasaude.com.br)
	IMAPPort        int      // MAIL_INGEST_IMAP_PORT (default: 993)
	Username        string   // MAIL_INGEST_USERNAME (ex: contato@plenyasaude.com.br)
	Password        string   // MAIL_INGEST_PASSWORD
	Folders         []string // MAIL_INGEST_FOLDERS (csv) — default: "INBOX,Sent Items"
	MaxAttachmentMB int      // MAIL_INGEST_MAX_ATTACHMENT_MB (default: 50)
	AttachmentDir   string   // MAIL_INGEST_ATTACHMENT_DIR (default: /app/uploads/email-attachments)
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
	// CoexistenceEnabled trata ecos (mensagens enviadas pelo app do celular) que
	// chegam no webhook quando o número roda em coexistence. WHATSAPP_COEXISTENCE.
	CoexistenceEnabled bool
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
	// CORSOrigin — single-origin legado (mantido para retrocompat). Quando
	// CORSOrigins é populado, CORSOrigin não é usado pelo middleware.
	CORSOrigin string
	// CORSOrigins — M1: lista de origens permitidas (multi-domínio).
	// Lido de CORS_ORIGINS (CSV). Em prod tipicamente:
	//   "https://app.plenyasaude.com.br,https://minha.plenyasaude.com.br"
	// Em dev cai no default ["http://localhost:3000"].
	// Validate() rejeita "*" em produção.
	CORSOrigins []string
	// TrustedProxies — CIDRs/IPs autorizados a setar X-Forwarded-For/X-Real-IP.
	// Em prod: IP da rede do reverse proxy (Coolify/Traefik). Em dev: loopback.
	// Lido de TRUSTED_PROXIES (CSV). Usado pelo Fiber EnableTrustedProxyCheck.
	TrustedProxies []string
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

// MagicLinkConfig — secret separado pra magic links do Score Light (e similares).
// Default: cai no JWTConfig.Secret se vazio (com warning), pra retrocompat.
// Idealmente prod usa MAGIC_LINK_SECRET dedicado pra que vazamento de um não
// quebre o outro.
type MagicLinkConfig struct {
	Secret string        // MAGIC_LINK_SECRET
	TTL    time.Duration // MAGIC_LINK_TTL — default 7d (link de claim do Escore chega por email; precisa durar dias, não minutos)
}

type SecurityConfig struct {
	EncryptionKey string
	OAuthTokenKey string // OAUTH_TOKEN_KEY — chave dedicada pra cifrar OAuth tokens (Calendar etc.). Default: EncryptionKey.
	// BlindIndexKey — M4 — chave HMAC pra gerar blind index de CPF (e futuros).
	// Permite buscar Patient by CPF sem descriptografar tabela inteira.
	// Em prod: BLIND_INDEX_KEY (hex 32 bytes). Em dev cai pra EncryptionKey
	// com warning. NUNCA logar valor.
	BlindIndexKey   string
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

// SignatureConfig — assinatura ICP-Brasil (hardening + e-CPF em nuvem).
//
// CARIMBO DE TEMPO (RFC 3161 / PAdES-T): quando TSAURL vazio, assina em PAdES básico
// (AD-RB — válido para CFM/ITI). Quando configurado com uma ACT credenciada ICP-Brasil,
// embute carimbo de tempo (prova a data independente do relógio do servidor; recomendado
// para controlados). TSAs ICP-Brasil costumam exigir credenciais.
//
// E-CPF EM NUVEM (PSC/broker): quando CloudEnabled e o médico tiver certificado em nuvem
// vinculado, a assinatura é disparada via API do PSC (Certillion/IntegraICP/VIDaaS/BirdID/
// SafeID). A chave privada nunca sai do HSM do provedor; o titular autoriza por push/OTP.
// Gated off por default (padrão SNCRProductionProvider) até haver contrato/credencial do PSC.
type SignatureConfig struct {
	TSAURL      string // ICP_TSA_URL — endpoint da ACT (RFC 3161). Vazio = sem carimbo de tempo.
	TSAUsername string // ICP_TSA_USER
	TSAPassword string // ICP_TSA_PASS

	CloudEnabled  bool   // ICP_CLOUD_ENABLED — liga o provedor de assinatura em nuvem
	CloudProvider string // ICP_CLOUD_PROVIDER — integraicp (default) | certillion
	CloudBaseURL  string // ICP_CLOUD_BASE_URL — base da API REST do PSC/broker
	CloudAPIKey   string // ICP_CLOUD_API_KEY — credencial do app integrador (Bearer)
}

type ClaudeConfig struct {
	APIKey string
	Model  string
	// NoteModel — modelo usado p/ gerar nota clínica/anamnese a partir do transcript
	// da teleconsulta (mais capaz que o Model default). Configurável; se indisponível
	// na conta, o serviço cai pro Model.
	NoteModel string // CLAUDE_NOTE_MODEL
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
			// M1 — multi-domínio. Quando CORS_ORIGINS vazio, mantemos
			// retrocompat: usa CORSOrigin como única origem permitida.
			CORSOrigins: parseCSVOrDefault(getEnv("CORS_ORIGINS", ""),
				[]string{getEnv("CORS_ORIGIN", "http://localhost:3000")}),
			// Default em dev: loopback. Em prod, configurar TRUSTED_PROXIES com
			// CIDR/IP do reverse proxy (ex: "172.16.0.0/12" pra rede docker do Coolify,
			// ou IP específico do Traefik). Sem isso, c.IP() do Fiber confia no header
			// X-Forwarded-For vindo de qualquer hop — risco de spoofing.
			TrustedProxies: parseCSVOrDefault(getEnv("TRUSTED_PROXIES", ""), []string{"127.0.0.1", "::1"}),
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
			OAuthTokenKey:   getEnv("OAUTH_TOKEN_KEY", ""), // se vazio, fallback para EncryptionKey (warning na boot)
			BlindIndexKey:   getEnv("BLIND_INDEX_KEY", ""), // M4 — fallback EncryptionKey
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
			// Default false: o SNCR ainda não é obrigatório (prorrogado p/ 30/09/2026) e a API
			// de emissores não foi publicada. Receita de controlado sai impressa p/ assinatura
			// manual. Ligar (com PRODUCTION_MODE + API_URL/KEY) quando a Anvisa abrir a integração.
			Enabled:        getEnvAsBool("SNCR_ENABLED", false),
			ProductionMode: getEnvAsBool("SNCR_PRODUCTION_MODE", false),
			APIURL:         getEnv("SNCR_API_URL", "https://sncr.anvisa.gov.br/api/v1"),
			APIKey:         getEnv("SNCR_API_KEY", ""),
		},
		Signature: SignatureConfig{
			TSAURL:        getEnv("ICP_TSA_URL", ""),
			TSAUsername:   getEnv("ICP_TSA_USER", ""),
			TSAPassword:   getEnv("ICP_TSA_PASS", ""),
			CloudEnabled:  getEnvAsBool("ICP_CLOUD_ENABLED", false),
			CloudProvider: getEnv("ICP_CLOUD_PROVIDER", "integraicp"),
			CloudBaseURL:  getEnv("ICP_CLOUD_BASE_URL", ""),
			CloudAPIKey:   getEnv("ICP_CLOUD_API_KEY", ""),
		},
		Claude: ClaudeConfig{
			APIKey:    getEnv("CLAUDE_API_KEY", ""),
			Model:     getEnv("CLAUDE_MODEL", "claude-3-5-haiku-20241022"),
			NoteModel: getEnv("CLAUDE_NOTE_MODEL", "claude-sonnet-4-6"),
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
			Provider:             getEnv("EMAIL_PROVIDER", "smtp"),
			ResendAPIKey:         getEnv("RESEND_API_KEY", ""),
			SMTPHost:             getEnv("SMTP_HOST", ""),
			SMTPPort:             getEnv("SMTP_PORT", "587"),
			SMTPUser:             getEnv("SMTP_USER", ""),
			SMTPPass:             getEnv("SMTP_PASSWORD", ""),
			FromAddress:          getEnv("EMAIL_FROM", "no-reply@plenyasaude.com.br"),
			FromName:             getEnv("EMAIL_FROM_NAME", "Plenya"),
			LeadReplyFromAddress: getEnv("EMAIL_LEAD_REPLY_FROM", "contato@plenyasaude.com.br"),
		},
		MailIngest: MailIngestConfig{
			Enabled:         getEnvAsBool("MAIL_INGEST_ENABLED", false),
			IMAPHost:        getEnv("MAIL_INGEST_IMAP_HOST", "mail.plenyasaude.com.br"),
			IMAPPort:        getEnvAsInt("MAIL_INGEST_IMAP_PORT", 993),
			Username:        getEnv("MAIL_INGEST_USERNAME", ""),
			Password:        getEnv("MAIL_INGEST_PASSWORD", ""),
			Folders:         parseCSVOrDefault(getEnv("MAIL_INGEST_FOLDERS", ""), []string{"INBOX", "Sent Items"}),
			MaxAttachmentMB: getEnvAsInt("MAIL_INGEST_MAX_ATTACHMENT_MB", 50),
			AttachmentDir:   getEnv("MAIL_INGEST_ATTACHMENT_DIR", "/app/uploads/email-attachments"),
		},
		WhatsApp: WhatsAppConfig{
			AppSecret:          getEnv("WHATSAPP_APP_SECRET", ""),
			AccessToken:        getEnv("WHATSAPP_ACCESS_TOKEN", ""),
			PhoneNumberID:      getEnv("WHATSAPP_PHONE_NUMBER_ID", ""),
			WebhookVerifyToken: getEnv("WHATSAPP_WEBHOOK_VERIFY_TOKEN", ""),
			TemplateMagicLink:  getEnv("WHATSAPP_TEMPLATE_MAGIC_LINK", "magic_link"),
			TemplateLeadAlert:  getEnv("WHATSAPP_TEMPLATE_LEAD_ALERT", "lead_alert"),
			GraphAPIVersion:    getEnv("WHATSAPP_GRAPH_API_VERSION", "v18.0"),
			CoexistenceEnabled: getEnv("WHATSAPP_COEXISTENCE", "true") != "false",
		},
		CRM: CRMConfig{
			AdminURL:          getEnv("CRM_ADMIN_URL", "http://localhost:3000"),
			LeadNotifyUserIDs: parseCSV(getEnv("CRM_LEAD_NOTIFY_USER_IDS", "")),
		},
		ReceptionBot: ReceptionBotConfig{
			Enabled:         getEnvAsBool("RECEPTION_BOT_ENABLED", false),
			DefaultMode:     getEnv("RECEPTION_BOT_DEFAULT_MODE", "off"),
			FallbackMinutes: getEnvAsInt("RECEPTION_BOT_FALLBACK_MINUTES", 5),
			MaxMsgsPerHour:  getEnvAsInt("RECEPTION_BOT_MAX_MSGS_HOUR", 6),
			ConsultDoctorID: getEnv("RECEPTION_CONSULT_DOCTOR_ID", ""),
		},
		Site: SiteConfig{
			PublicURL: getEnv("SITE_PUBLIC_URL", "http://localhost:3002"),
		},
		Google: GoogleConfig{
			ClientID:     getEnv("GOOGLE_CLIENT_ID", ""),
			ClientSecret: getEnv("GOOGLE_CLIENT_SECRET", ""),
			RedirectURL:  getEnv("GOOGLE_REDIRECT_URL", "https://app.plenyasaude.com.br/api/v1/integrations/google/callback"),
		},
		DailyCo: DailyCoConfig{
			APIKey:        getEnv("DAILY_CO_API_KEY", ""),
			Domain:        getEnv("DAILY_CO_DOMAIN", ""),
			WebhookSecret: getEnv("DAILY_CO_WEBHOOK_SECRET", ""),
		},
		PatientPortal: PatientPortalConfig{
			PublicURL: getEnv("PATIENT_PORTAL_URL", "http://localhost:3000"),
		},
		MagicLink: MagicLinkConfig{
			Secret: getEnv("MAGIC_LINK_SECRET", ""),
			TTL:    parseDurationOrDefault(getEnv("MAGIC_LINK_TTL", ""), 7*24*time.Hour),
		},
		Dev: DevConfig{
			BypassAuth: getEnvAsBool("DEV_BYPASS_AUTH", false),
		},
		Turnstile: TurnstileConfig{
			Secret:  getEnv("TURNSTILE_SECRET", ""),
			SiteKey: getEnv("TURNSTILE_SITE_KEY", ""),
		},
	}

	// Fallbacks com warnings (não bloqueia boot pra retrocompat)
	if cfg.Security.OAuthTokenKey == "" {
		cfg.Security.OAuthTokenKey = cfg.Security.EncryptionKey
		log.Println("⚠️  OAUTH_TOKEN_KEY não setado — usando ENCRYPTION_KEY (defina chave dedicada em produção)")
	}
	if cfg.Security.BlindIndexKey == "" {
		cfg.Security.BlindIndexKey = cfg.Security.EncryptionKey
		if cfg.Server.Environment == "production" {
			log.Println("⚠️  BLIND_INDEX_KEY não setado em PROD — usando ENCRYPTION_KEY (rotação fica acoplada; defina chave dedicada)")
		} else {
			log.Println("⚠️  BLIND_INDEX_KEY não setado — usando ENCRYPTION_KEY (dev fallback)")
		}
	}
	if cfg.MagicLink.Secret == "" {
		cfg.MagicLink.Secret = cfg.JWT.Secret
		log.Println("⚠️  MAGIC_LINK_SECRET não setado — usando JWT_SECRET (defina secret dedicado em produção)")
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

	// M1 — proíbe wildcard CORS em produção. AllowOriginsFunc com retorno
	// true incondicional é equivalente a "*" e elimina a defesa CSRF que o
	// browser oferece. Em dev qualquer "*" passa (conveniência).
	if c.Server.Environment == "production" {
		for _, o := range c.Server.CORSOrigins {
			if o == "*" {
				return fmt.Errorf("CORS_ORIGINS cannot contain '*' in production (set explicit hosts)")
			}
		}
		if c.Server.CORSOrigin == "*" {
			return fmt.Errorf("CORS_ORIGIN cannot be '*' in production")
		}
	}

	// CRITICAL C1 — Bloqueia DEV_BYPASS_AUTH em produção. Se essa flag chegar
	// junto com ENVIRONMENT=production, abortamos o boot pra evitar que um
	// deploy acidentalmente exponha login automatizado de admin.
	if c.Dev.BypassAuth && c.Server.Environment == "production" {
		return fmt.Errorf("DEV_BYPASS_AUTH cannot be enabled in production environment (set DEV_BYPASS_AUTH=false or remove)")
	}

	return nil
}

// parseDurationOrDefault — parse de duration com fallback (não exporta).
func parseDurationOrDefault(s string, def time.Duration) time.Duration {
	if s == "" {
		return def
	}
	d, err := time.ParseDuration(s)
	if err != nil {
		return def
	}
	return d
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

// parseCSVOrDefault retorna parseCSV(s) se não vazio, senão def.
func parseCSVOrDefault(s string, def []string) []string {
	if v := parseCSV(s); len(v) > 0 {
		return v
	}
	return def
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
