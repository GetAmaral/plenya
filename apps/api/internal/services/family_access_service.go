package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
)

// FamilyAccessService — paciente cria convites granulares pra familiares verem dados.
//
// V1 (consumo): paciente original ainda é o dono. Familiar consumiu o convite vira
// um User normal e usa um endpoint adicional do portal pra alternar entre suas
// próprias contas + as compartilhadas. V1 do FRONTEND só lista grants no perfil
// do paciente. A integração full do "switch entre contas" fica pra V2 (precisa
// de mexida no auth-store + middleware de seleção).
type FamilyAccessService struct {
	db           *gorm.DB
	cfg          *config.Config
	emailService *EmailService
}

func NewFamilyAccessService(db *gorm.DB, cfg *config.Config, emailService *EmailService) *FamilyAccessService {
	return &FamilyAccessService{db: db, cfg: cfg, emailService: emailService}
}

// CreateInviteInput é o payload do CreateInvite.
type CreateFamilyInviteInput struct {
	PatientID    uuid.UUID
	GranteeEmail string
	GranteeLabel string
	Scope        models.FamilyAccessScope
}

// CreateInvite cria um grant pendente + dispara email com link.
func (s *FamilyAccessService) CreateInvite(in CreateFamilyInviteInput) (*models.FamilyAccessGrant, error) {
	in.GranteeEmail = strings.ToLower(strings.TrimSpace(in.GranteeEmail))
	if in.GranteeEmail == "" || !strings.Contains(in.GranteeEmail, "@") {
		return nil, errors.New("email inválido")
	}
	if strings.TrimSpace(in.GranteeLabel) == "" {
		return nil, errors.New("rótulo obrigatório (ex: 'Esposa Mariana')")
	}

	scopeJSON, err := json.Marshal(in.Scope)
	if err != nil {
		return nil, err
	}

	token, err := generateRandomToken(32)
	if err != nil {
		return nil, err
	}

	grant := models.FamilyAccessGrant{
		PatientID:    in.PatientID,
		GranteeEmail: in.GranteeEmail,
		GranteeLabel: in.GranteeLabel,
		Token:        token,
		ScopeJSON:    datatypes.JSON(scopeJSON),
		Status:       models.FamilyGrantInvited,
	}
	if err := s.db.Create(&grant).Error; err != nil {
		return nil, err
	}

	// Best-effort: dispara email.
	link := s.inviteLink(token)
	if s.emailService != nil {
		var patient models.Patient
		_ = s.db.Select("name").First(&patient, "id = ?", in.PatientID).Error
		_ = s.emailService.SendFamilyInvite(in.GranteeEmail, in.GranteeLabel, patient.Name, link)
	}

	return &grant, nil
}

func (s *FamilyAccessService) inviteLink(token string) string {
	base := s.cfg.PatientPortal.PublicURL
	if base == "" {
		base = "https://meu.plenyasaude.com.br"
	}
	return fmt.Sprintf("%s/auth/familia?token=%s", strings.TrimRight(base, "/"), token)
}

// List retorna grants ativos/pending do paciente.
func (s *FamilyAccessService) List(patientID uuid.UUID) ([]models.FamilyAccessGrant, error) {
	var rows []models.FamilyAccessGrant
	err := s.db.Where("patient_id = ? AND status != ?", patientID, models.FamilyGrantRevoked).
		Order("created_at DESC").
		Find(&rows).Error
	return rows, err
}

// Revoke marca grant revogado (idempotente).
func (s *FamilyAccessService) Revoke(patientID, grantID uuid.UUID) error {
	now := time.Now().UTC()
	res := s.db.Model(&models.FamilyAccessGrant{}).
		Where("id = ? AND patient_id = ?", grantID, patientID).
		Updates(map[string]any{
			"status":     models.FamilyGrantRevoked,
			"revoked_at": now,
		})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("grant não encontrado")
	}
	return nil
}

// UpdateScope permite paciente alterar permissões do grant (ex: tirar mensagens).
func (s *FamilyAccessService) UpdateScope(patientID, grantID uuid.UUID, scope models.FamilyAccessScope) error {
	scopeJSON, err := json.Marshal(scope)
	if err != nil {
		return err
	}
	res := s.db.Model(&models.FamilyAccessGrant{}).
		Where("id = ? AND patient_id = ?", grantID, patientID).
		Update("scope_json", datatypes.JSON(scopeJSON))
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("grant não encontrado")
	}
	return nil
}

// ConsumeInvite — familiar abre o link, sistema vincula o User existente ou retorna
// instrução pra criar conta. V1 simplificado: precisa que o GranteeEmail seja
// igual ao email de um User patient existente.
func (s *FamilyAccessService) ConsumeInvite(token string, granteeUserID uuid.UUID) (*models.FamilyAccessGrant, error) {
	var grant models.FamilyAccessGrant
	if err := s.db.Where("token = ?", token).First(&grant).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("convite não encontrado")
		}
		return nil, err
	}
	if grant.Status != models.FamilyGrantInvited {
		return nil, errors.New("convite já consumido ou revogado")
	}
	if time.Now().UTC().After(grant.ExpiresAt) {
		return nil, errors.New("convite expirado")
	}

	// Confirma que userID corresponde ao email do grant
	var user models.User
	if err := s.db.First(&user, "id = ?", granteeUserID).Error; err != nil {
		return nil, err
	}
	if !strings.EqualFold(user.Email, grant.GranteeEmail) {
		return nil, errors.New("email do convite não bate com sua conta")
	}

	now := time.Now().UTC()
	grant.GranteeUserID = &granteeUserID
	grant.Status = models.FamilyGrantActive
	grant.AcceptedAt = &now
	if err := s.db.Save(&grant).Error; err != nil {
		return nil, err
	}
	return &grant, nil
}

// ListGrantedToMe — pra UI "ver outras pessoas que você acompanha".
// Retorna grants ativos onde o User logado é o grantee.
func (s *FamilyAccessService) ListGrantedToMe(userID uuid.UUID) ([]models.FamilyAccessGrant, error) {
	var rows []models.FamilyAccessGrant
	err := s.db.Where("grantee_user_id = ? AND status = ?", userID, models.FamilyGrantActive).
		Order("created_at DESC").
		Find(&rows).Error
	return rows, err
}
