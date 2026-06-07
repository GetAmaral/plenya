package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// ============================================================
// RelationshipProfileService — memória de longo prazo (SOCIAL) da Lívia
// ============================================================
//
// CRUD fino do relationship_profiles (1 por dono lead|patient). Só dado social/relacional:
// nada clínico entra aqui (LGPD/CFM — ver docs/emr/plano-livia-memoria-dossie-360.md §0/§3.1).
// O resumo rolante (rolling_summary) é mantido pelo job e injetado no prompt do recepcionista.
// Reutilizado pela tela 360 (fases seguintes).

// RelationshipProfileService acessa o perfil de relacionamento por dono.
type RelationshipProfileService struct {
	db *gorm.DB
}

// NewRelationshipProfileService cria o serviço.
func NewRelationshipProfileService(db *gorm.DB) *RelationshipProfileService {
	return &RelationshipProfileService{db: db}
}

// Get retorna o perfil do dono, ou (nil, nil) quando ainda não existe.
func (s *RelationshipProfileService) Get(ctx context.Context, ownerType string, ownerID uuid.UUID) (*models.RelationshipProfile, error) {
	if !isValidOwnerType(ownerType) || ownerID == uuid.Nil {
		return nil, ErrConversationOwnerInvalid
	}
	var p models.RelationshipProfile
	err := s.db.WithContext(ctx).
		Where("owner_type = ? AND owner_id = ?", ownerType, ownerID).
		First(&p).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

// GetOrCreate retorna o perfil do dono, criando uma linha vazia se ainda não existir.
func (s *RelationshipProfileService) GetOrCreate(ctx context.Context, ownerType string, ownerID uuid.UUID) (*models.RelationshipProfile, error) {
	existing, err := s.Get(ctx, ownerType, ownerID)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return existing, nil
	}
	p := &models.RelationshipProfile{OwnerType: ownerType, OwnerID: ownerID}
	if err := s.db.WithContext(ctx).Create(p).Error; err != nil {
		// Corrida: outra goroutine criou primeiro → relê.
		if again, gErr := s.Get(ctx, ownerType, ownerID); gErr == nil && again != nil {
			return again, nil
		}
		return nil, err
	}
	return p, nil
}

// SaveSummary persiste o resumo rolante (social) regerado e o contador de mensagens resumidas.
func (s *RelationshipProfileService) SaveSummary(ctx context.Context, id uuid.UUID, summary string, msgCount int) error {
	now := time.Now().UTC()
	return s.db.WithContext(ctx).Model(&models.RelationshipProfile{}).
		Where("id = ?", id).
		Updates(map[string]any{
			"rolling_summary":    summary,
			"summary_updated_at": now,
			"summary_msg_count":  msgCount,
		}).Error
}
