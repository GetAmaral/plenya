package services

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

var (
	ErrInviteNotFound  = errors.New("invite not found")
	ErrInviteExpired   = errors.New("invite expired or already used")
	ErrPatientNoEmail  = errors.New("patient has no email — cannot send invite by email")
	ErrPatientNotFoundForUser = errors.New("user has no patient record")
)

// PatientPortalService cuida de:
//  - convites enviados pela equipe (gera+envia+consome)
//  - magic link "puro" pra paciente já cadastrado (login passwordless)
//  - "esqueci a senha" reusando magic link
//  - troca de senha do paciente
type PatientPortalService struct {
	db           *gorm.DB
	cfg          *config.Config
	authService  *AuthService
	emailService *EmailService
	whatsapp     *WhatsAppService
}

func NewPatientPortalService(
	db *gorm.DB,
	cfg *config.Config,
	authService *AuthService,
	emailService *EmailService,
	whatsapp *WhatsAppService,
) *PatientPortalService {
	return &PatientPortalService{
		db:           db,
		cfg:          cfg,
		authService:  authService,
		emailService: emailService,
		whatsapp:     whatsapp,
	}
}

// ============================================================
// Invites — equipe convida paciente pra acessar o portal
// ============================================================

// PortalURL retorna a URL base do portal do paciente. Default: meu.plenyasaude.com.br.
func (s *PatientPortalService) PortalURL() string {
	if s.cfg.PatientPortal.PublicURL != "" {
		return strings.TrimRight(s.cfg.PatientPortal.PublicURL, "/")
	}
	return "https://meu.plenyasaude.com.br"
}

// CreateInviteInput é o payload do CreateInvite.
type CreateInviteInput struct {
	PatientID  uuid.UUID
	InvitedBy  uuid.UUID
	SendEmail  bool
	SendWA     bool
}

// CreateInviteResult é retornado pra UI da equipe — inclui o link gerado pra
// quem queira copiar (caso o paciente perca o email/WA).
type CreateInviteResult struct {
	InviteID  uuid.UUID `json:"inviteId"`
	Link      string    `json:"link"`
	ExpiresAt time.Time `json:"expiresAt"`
	SentEmail bool      `json:"sentEmail"`
	SentWA    bool      `json:"sentWA"`
}

// CreateInvite cria um token de convite e dispara nos canais escolhidos.
// Erros de envio não são fatais — o link sempre é retornado pra UI.
func (s *PatientPortalService) CreateInvite(in CreateInviteInput) (*CreateInviteResult, error) {
	var patient models.Patient
	if err := s.db.First(&patient, "id = ?", in.PatientID).Error; err != nil {
		return nil, fmt.Errorf("patient not found: %w", err)
	}

	// Gera token aleatório (256 bits)
	token, err := generateRandomToken(32)
	if err != nil {
		return nil, err
	}

	invite := models.PatientPortalInvite{
		PatientID:    in.PatientID,
		InvitedBy:    in.InvitedBy,
		Token:        token,
		ChannelEmail: in.SendEmail,
		ChannelWA:    in.SendWA,
	}
	if err := s.db.Create(&invite).Error; err != nil {
		return nil, fmt.Errorf("failed to create invite: %w", err)
	}

	link := fmt.Sprintf("%s/auth/invite?token=%s", s.PortalURL(), token)

	result := &CreateInviteResult{
		InviteID:  invite.ID,
		Link:      link,
		ExpiresAt: invite.ExpiresAt,
	}

	if in.SendEmail && s.emailService != nil && patient.Email != nil && *patient.Email != "" && !models.IsPlaceholderEmail(*patient.Email) {
		if err := s.emailService.SendPortalInvite(*patient.Email, patient.Name, link); err == nil {
			result.SentEmail = true
		}
	}
	if in.SendWA && s.whatsapp != nil && patient.Phone != nil && *patient.Phone != "" {
		if err := s.whatsapp.SendMagicLink(*patient.Phone, link); err == nil {
			result.SentWA = true
		}
	}

	return result, nil
}

// GetInviteStatus retorna o último convite do paciente (pra UI mostrar status).
func (s *PatientPortalService) GetInviteStatus(patientID uuid.UUID) (*models.PatientPortalInvite, error) {
	var invite models.PatientPortalInvite
	err := s.db.Where("patient_id = ?", patientID).
		Order("created_at DESC").
		First(&invite).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &invite, nil
}

// ConsumeInviteResult é retornado quando paciente abre o link de convite.
type ConsumeInviteResult struct {
	*dto.AuthResponse
	PatientID uuid.UUID `json:"patientId"`
}

// ConsumeInvite valida o token e gera tokens JWT pro paciente entrar no portal.
// Marca o convite como aceito (single-use) e garante que o User existe com role=patient.
// Se o User ainda não existe (caso muito raro — Patient criado sem UserID), cria.
func (s *PatientPortalService) ConsumeInvite(token string) (*ConsumeInviteResult, error) {
	var invite models.PatientPortalInvite
	if err := s.db.Where("token = ?", token).First(&invite).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrInviteNotFound
		}
		return nil, err
	}
	if !invite.IsValid() {
		return nil, ErrInviteExpired
	}

	var patient models.Patient
	if err := s.db.First(&patient, "id = ?", invite.PatientID).Error; err != nil {
		return nil, fmt.Errorf("patient not found: %w", err)
	}

	var user models.User
	if err := s.db.First(&user, "id = ?", patient.UserID).Error; err != nil {
		return nil, fmt.Errorf("user not found for patient: %w", err)
	}

	// Garante role=patient
	roles := user.GetRoles()
	hasPatient := false
	for _, r := range roles {
		if r == string(models.RolePatient) {
			hasPatient = true
			break
		}
	}
	if !hasPatient {
		_ = user.SetRoles(append(roles, string(models.RolePatient)))
		if err := s.db.Save(&user).Error; err != nil {
			return nil, err
		}
	}

	// Marca convite como aceito
	now := time.Now().UTC()
	invite.AcceptedAt = &now
	if err := s.db.Save(&invite).Error; err != nil {
		return nil, err
	}

	authResp, err := s.authService.GenerateTokensForUser(&user)
	if err != nil {
		return nil, err
	}

	return &ConsumeInviteResult{
		AuthResponse: authResp,
		PatientID:    patient.ID,
	}, nil
}

// ============================================================
// Magic link puro — paciente já cadastrado quer entrar sem senha
// ============================================================

// RequestMagicLink envia um magic link pro email do paciente. Não revela existência
// do User (sempre retorna nil — anti-enumeração).
func (s *PatientPortalService) RequestMagicLink(email string, ipHash *string) error {
	email = strings.ToLower(strings.TrimSpace(email))
	if email == "" {
		return errors.New("email required")
	}

	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		// Anti-enumeração — silencia not-found. Outros erros propagam.
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}

	// Só permite magic link pra Users com role=patient (defesa: profissional não deve
	// usar essa porta — tem que ir pelo fluxo do EMR).
	if !user.IsGranted(models.RolePatient) {
		return nil
	}

	token, err := generateRandomToken(32)
	if err != nil {
		return err
	}

	link := models.PatientMagicLink{
		UserID: user.ID,
		Token:  token,
		IPHash: ipHash,
	}
	if err := s.db.Create(&link).Error; err != nil {
		return err
	}

	url := fmt.Sprintf("%s/auth/magic?token=%s", s.PortalURL(), token)
	// Best-effort — não fatal se falhar (usuário pode pedir de novo)
	if s.emailService != nil {
		_ = s.emailService.SendPatientMagicLink(email, user.Name, url)
	}
	return nil
}

// ConsumeMagicLink valida o token e devolve sessão JWT.
func (s *PatientPortalService) ConsumeMagicLink(token string) (*dto.AuthResponse, error) {
	var link models.PatientMagicLink
	if err := s.db.Where("token = ?", token).First(&link).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrInviteNotFound
		}
		return nil, err
	}
	if !link.IsValid() {
		return nil, ErrInviteExpired
	}

	var user models.User
	if err := s.db.First(&user, "id = ?", link.UserID).Error; err != nil {
		return nil, err
	}
	if !user.IsGranted(models.RolePatient) {
		return nil, ErrInviteExpired
	}

	now := time.Now().UTC()
	link.UsedAt = &now
	if err := s.db.Save(&link).Error; err != nil {
		return nil, err
	}

	return s.authService.GenerateTokensForUser(&user)
}

// ============================================================
// Login email + senha (paciente)
// ============================================================

// LoginPassword autentica paciente por email+senha. Reusa fluxo do AuthService mas
// rejeita Users sem role=patient (não deixa profissional logar pelo portal).
func (s *PatientPortalService) LoginPassword(email, password string) (*dto.AuthResponse, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	if email == "" || password == "" {
		return nil, ErrInvalidCredentials
	}

	var user models.User
	if err := s.db.Preload("SelectedPatient").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, ErrInvalidCredentials
	}
	if user.PasswordHash == nil {
		// Não tem senha definida — pede magic link
		return nil, errors.New("conta sem senha definida — use o link mágico")
	}
	if !user.IsGranted(models.RolePatient) {
		return nil, ErrInvalidCredentials
	}
	if err := bcrypt.CompareHashAndPassword([]byte(*user.PasswordHash), []byte(password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	return s.authService.GenerateTokensForUser(&user)
}

// SetPassword permite paciente definir senha (quando ainda não tem) ou trocar.
// Sem validação de senha atual — usado depois de magic link ou consume invite.
// Pra trocar com senha atual, usa AuthService.ChangePassword.
func (s *PatientPortalService) SetPassword(userID uuid.UUID, password string) error {
	if len(password) < 8 {
		return errors.New("senha deve ter ao menos 8 caracteres")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	hashStr := string(hash)
	return s.db.Model(&models.User{}).Where("id = ?", userID).
		UpdateColumn("password_hash", hashStr).Error
}

// ============================================================
// Helpers
// ============================================================

func generateRandomToken(byteLen int) (string, error) {
	b := make([]byte, byteLen)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
