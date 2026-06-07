package services

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// ============================================================
// RelationshipFactService — memória SEMÂNTICA (fatos sociais) da Lívia
// ============================================================
//
// Fatos atômicos sociais por dono, com validade temporal (valid_until nil = ativo). Operações
// ADD/UPDATE convergem em SetFact (1 ativo por key); DELETE em CloseFact. Só social — nada
// clínico (LGPD/CFM, §3 do plano). Alimentado pela IA (job) e pela equipe (tela 360).

type RelationshipFactService struct {
	db *gorm.DB
}

func NewRelationshipFactService(db *gorm.DB) *RelationshipFactService {
	return &RelationshipFactService{db: db}
}

// ListActive devolve os fatos ativos (valid_until nil, não sensíveis) do dono, por categoria/key.
func (s *RelationshipFactService) ListActive(ctx context.Context, ownerType string, ownerID uuid.UUID) ([]models.RelationshipFact, error) {
	if !isValidOwnerType(ownerType) || ownerID == uuid.Nil {
		return nil, ErrConversationOwnerInvalid
	}
	var facts []models.RelationshipFact
	err := s.db.WithContext(ctx).
		Where("owner_type = ? AND owner_id = ?", ownerType, ownerID).
		Where("valid_until IS NULL").
		Where("sensitive = ?", false).
		Order("category ASC, key ASC").
		Find(&facts).Error
	return facts, err
}

// SetFact aplica um ADD/UPDATE: garante 1 ativo por key. Se já existir ativo com o mesmo valor,
// é NOOP (changed=false). Se o valor mudou, fecha o anterior (valid_until=now) e insere o novo.
func (s *RelationshipFactService) SetFact(
	ctx context.Context,
	ownerType string, ownerID uuid.UUID,
	category, key, value, source string,
	addedBy *uuid.UUID, confidence *float32,
) (changed bool, err error) {
	key = strings.TrimSpace(key)
	value = strings.TrimSpace(value)
	if key == "" || value == "" {
		return false, nil
	}
	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var existing models.RelationshipFact
		findErr := tx.Where("owner_type = ? AND owner_id = ? AND key = ? AND valid_until IS NULL", ownerType, ownerID, key).
			First(&existing).Error
		now := time.Now().UTC()
		if findErr == nil {
			if strings.EqualFold(strings.TrimSpace(existing.Value), value) {
				return nil // NOOP — mesmo valor
			}
			if uErr := tx.Model(&models.RelationshipFact{}).Where("id = ?", existing.ID).
				Update("valid_until", now).Error; uErr != nil {
				return uErr
			}
		} else if findErr != gorm.ErrRecordNotFound {
			return findErr
		}
		nf := &models.RelationshipFact{
			OwnerType:  ownerType,
			OwnerID:    ownerID,
			Category:   category,
			Key:        key,
			Value:      value,
			Source:     source,
			AddedBy:    addedBy,
			Confidence: confidence,
		}
		if cErr := tx.Create(nf).Error; cErr != nil {
			return cErr
		}
		changed = true
		return nil
	})
	return changed, err
}

// CloseFact fecha (DELETE lógico) o fato ativo de uma key, preservando o histórico.
func (s *RelationshipFactService) CloseFact(ctx context.Context, ownerType string, ownerID uuid.UUID, key string) error {
	return s.db.WithContext(ctx).Model(&models.RelationshipFact{}).
		Where("owner_type = ? AND owner_id = ? AND key = ? AND valid_until IS NULL", ownerType, ownerID, strings.TrimSpace(key)).
		Update("valid_until", time.Now().UTC()).Error
}

// AddManual cria um fato pela equipe (source=staff). Reusa SetFact (1 ativo por key).
func (s *RelationshipFactService) AddManual(
	ctx context.Context,
	ownerType string, ownerID uuid.UUID,
	category, key, value string, addedBy uuid.UUID,
) (bool, error) {
	if !isValidOwnerType(ownerType) || ownerID == uuid.Nil {
		return false, ErrConversationOwnerInvalid
	}
	return s.SetFact(ctx, ownerType, ownerID, category, key, value, models.RelationshipSourceStaff, &addedBy, nil)
}

// UpdateValueByID edita o valor de um fato existente (equipe). Atualiza in-place a key ativa.
func (s *RelationshipFactService) UpdateValueByID(ctx context.Context, id uuid.UUID, value string, addedBy uuid.UUID) error {
	value = strings.TrimSpace(value)
	if value == "" {
		return ErrConversationOwnerInvalid
	}
	return s.db.WithContext(ctx).Model(&models.RelationshipFact{}).
		Where("id = ? AND valid_until IS NULL", id).
		Updates(map[string]any{"value": value, "source": models.RelationshipSourceStaff, "added_by": addedBy}).Error
}

// CloseByID fecha (DELETE lógico) um fato ativo por id (equipe).
func (s *RelationshipFactService) CloseByID(ctx context.Context, id uuid.UUID) error {
	return s.db.WithContext(ctx).Model(&models.RelationshipFact{}).
		Where("id = ? AND valid_until IS NULL", id).
		Update("valid_until", time.Now().UTC()).Error
}
