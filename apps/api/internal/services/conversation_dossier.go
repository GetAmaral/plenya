package services

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/plenya/api/internal/models"
)

// ============================================================
// Dossiê 360 (social) — leitura + edição manual pela equipe
// ============================================================
//
// Visão estruturada do que se sabe SOCIALMENTE de uma pessoa (resumo + fatos), montada a partir
// de relationship_profiles + relationship_facts. Nada clínico (LGPD/CFM, §0/§4 do plano). Usada
// pelo painel "Dossiê" na conversa. Alimentação: IA (job) + equipe (estes métodos).

// DossierFact é um fato social para exibição na tela 360.
type DossierFact struct {
	ID        string    `json:"id"`
	Category  string    `json:"category"`
	Key       string    `json:"key"`
	Label     string    `json:"label"`
	Value     string    `json:"value"`
	Source    string    `json:"source"` // ai|staff|form|consulta
	UpdatedAt time.Time `json:"updatedAt"`
}

// DossierView é a visão 360 social de uma pessoa.
type DossierView struct {
	OwnerType         string        `json:"ownerType"`
	OwnerID           string        `json:"ownerId"`
	IsPatient         bool          `json:"isPatient"`
	RollingSummary    string        `json:"rollingSummary"`
	SummaryUpdatedAt  *time.Time    `json:"summaryUpdatedAt,omitempty"`
	RelationshipStage string        `json:"relationshipStage"`
	Facts             []DossierFact `json:"facts"`
}

// GetDossier monta a visão 360 social (perfil + fatos ativos) de uma pessoa.
func (s *ConversationService) GetDossier(ctx context.Context, ownerType string, ownerID uuid.UUID) (*DossierView, error) {
	if !isValidOwnerType(ownerType) || ownerID == uuid.Nil {
		return nil, ErrConversationOwnerInvalid
	}
	view := &DossierView{
		OwnerType: ownerType,
		OwnerID:   ownerID.String(),
		IsPatient: ownerType == string(models.ConversationOwnerPatient),
		Facts:     []DossierFact{},
	}

	if prof, err := NewRelationshipProfileService(s.db).Get(ctx, ownerType, ownerID); err == nil && prof != nil {
		view.RollingSummary = prof.RollingSummary
		view.SummaryUpdatedAt = prof.SummaryUpdatedAt
		view.RelationshipStage = prof.RelationshipStage
	}

	facts, err := NewRelationshipFactService(s.db).ListActive(ctx, ownerType, ownerID)
	if err != nil {
		return nil, err
	}
	for _, f := range facts {
		view.Facts = append(view.Facts, DossierFact{
			ID:        f.ID.String(),
			Category:  f.Category,
			Key:       f.Key,
			Label:     factLabel(f.Key),
			Value:     f.Value,
			Source:    f.Source,
			UpdatedAt: f.UpdatedAt,
		})
	}
	return view, nil
}

// AddDossierFact registra um fato social pela equipe e devolve o dossiê atualizado.
func (s *ConversationService) AddDossierFact(ctx context.Context, ownerType string, ownerID uuid.UUID, category, key, value string, addedBy uuid.UUID) (*DossierView, error) {
	if !isValidOwnerType(ownerType) || ownerID == uuid.Nil {
		return nil, ErrConversationOwnerInvalid
	}
	cat := normalizeFactCategory(category)
	k := normalizeFactKey(key)
	if k == "" {
		return nil, ErrConversationOwnerInvalid
	}
	if _, err := NewRelationshipFactService(s.db).AddManual(ctx, ownerType, ownerID, cat, k, value, addedBy); err != nil {
		return nil, err
	}
	return s.GetDossier(ctx, ownerType, ownerID)
}

// UpdateDossierFact edita o valor de um fato (equipe) e devolve o dossiê atualizado.
func (s *ConversationService) UpdateDossierFact(ctx context.Context, ownerType string, ownerID uuid.UUID, factID uuid.UUID, value string, addedBy uuid.UUID) (*DossierView, error) {
	if err := NewRelationshipFactService(s.db).UpdateValueByID(ctx, factID, value, addedBy); err != nil {
		return nil, err
	}
	return s.GetDossier(ctx, ownerType, ownerID)
}

// DeleteDossierFact fecha (DELETE lógico) um fato (equipe) e devolve o dossiê atualizado.
func (s *ConversationService) DeleteDossierFact(ctx context.Context, ownerType string, ownerID uuid.UUID, factID uuid.UUID) (*DossierView, error) {
	if err := NewRelationshipFactService(s.db).CloseByID(ctx, factID); err != nil {
		return nil, err
	}
	return s.GetDossier(ctx, ownerType, ownerID)
}
