package services

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// ============================================================
// Rede de relacionamento + eventos/avisos (Fase C)
// ============================================================
//
// Pessoas importantes (cônjuge, filhos, quem indicou…) + datas/eventos para o time interagir
// proativamente. Só social (LGPD/CFM). Alimentado por IA (job) e equipe (tela 360). Ver
// docs/emr/plano-livia-memoria-dossie-360.md §2.3/§2.4/§5.

// ---------- Pessoas importantes ----------

type RelationshipPersonService struct{ db *gorm.DB }

func NewRelationshipPersonService(db *gorm.DB) *RelationshipPersonService {
	return &RelationshipPersonService{db: db}
}

func (s *RelationshipPersonService) ListByOwner(ctx context.Context, ownerType string, ownerID uuid.UUID) ([]models.ImportantPerson, error) {
	if !isValidOwnerType(ownerType) || ownerID == uuid.Nil {
		return nil, ErrConversationOwnerInvalid
	}
	var people []models.ImportantPerson
	err := s.db.WithContext(ctx).
		Where("owner_type = ? AND owner_id = ?", ownerType, ownerID).
		Order("created_at ASC").Find(&people).Error
	return people, err
}

// Upsert cria ou atualiza (dedupe por owner + nome, case-insensitive). Atualiza relação/aniversário/
// notas quando vierem preenchidos; não apaga o que já existe com vazio.
func (s *RelationshipPersonService) Upsert(
	ctx context.Context,
	ownerType string, ownerID uuid.UUID,
	name, relation string, birthday *time.Time, notes, source string, addedBy *uuid.UUID,
) (*models.ImportantPerson, error) {
	name = strings.TrimSpace(name)
	if name == "" || !isValidOwnerType(ownerType) || ownerID == uuid.Nil {
		return nil, ErrConversationOwnerInvalid
	}
	var existing models.ImportantPerson
	err := s.db.WithContext(ctx).
		Where("owner_type = ? AND owner_id = ? AND lower(name) = lower(?)", ownerType, ownerID, name).
		First(&existing).Error
	if err == nil {
		updates := map[string]any{}
		if r := strings.TrimSpace(relation); r != "" && r != existing.Relation {
			updates["relation"] = r
		}
		if birthday != nil {
			updates["birthday"] = birthday
		}
		if n := strings.TrimSpace(notes); n != "" && n != existing.Notes {
			updates["notes"] = n
		}
		if len(updates) > 0 {
			_ = s.db.WithContext(ctx).Model(&existing).Updates(updates).Error
		}
		return &existing, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	p := &models.ImportantPerson{
		OwnerType: ownerType, OwnerID: ownerID,
		Name: name, Relation: strings.TrimSpace(relation), Birthday: birthday,
		Notes: strings.TrimSpace(notes), Source: source, AddedBy: addedBy,
	}
	if cErr := s.db.WithContext(ctx).Create(p).Error; cErr != nil {
		return nil, cErr
	}
	return p, nil
}

func (s *RelationshipPersonService) UpdateByID(ctx context.Context, id uuid.UUID, name, relation string, birthday *time.Time, notes string) error {
	updates := map[string]any{
		"name":     strings.TrimSpace(name),
		"relation": strings.TrimSpace(relation),
		"notes":    strings.TrimSpace(notes),
		"birthday": birthday,
	}
	return s.db.WithContext(ctx).Model(&models.ImportantPerson{}).Where("id = ?", id).Updates(updates).Error
}

func (s *RelationshipPersonService) DeleteByID(ctx context.Context, id uuid.UUID) error {
	return s.db.WithContext(ctx).Where("id = ?", id).Delete(&models.ImportantPerson{}).Error
}

// ---------- Eventos / avisos ----------

type RelationshipEventService struct{ db *gorm.DB }

func NewRelationshipEventService(db *gorm.DB) *RelationshipEventService {
	return &RelationshipEventService{db: db}
}

func (s *RelationshipEventService) ListByOwner(ctx context.Context, ownerType string, ownerID uuid.UUID) ([]models.RelationshipEvent, error) {
	if !isValidOwnerType(ownerType) || ownerID == uuid.Nil {
		return nil, ErrConversationOwnerInvalid
	}
	var events []models.RelationshipEvent
	err := s.db.WithContext(ctx).
		Where("owner_type = ? AND owner_id = ?", ownerType, ownerID).
		Where("status <> ?", models.RelationshipEventDismissed).
		Order("event_date ASC").Find(&events).Error
	return events, err
}

// Upsert cria um evento (dedupe por owner + tipo + título). Não duplica eventos já registrados.
func (s *RelationshipEventService) Upsert(
	ctx context.Context,
	ownerType string, ownerID uuid.UUID, relatedPersonID *uuid.UUID,
	eventType, title string, eventDate time.Time, recurring bool, leadTimeDays int,
	source, notes string, addedBy *uuid.UUID,
) (*models.RelationshipEvent, error) {
	title = strings.TrimSpace(title)
	if title == "" || eventDate.IsZero() || !isValidOwnerType(ownerType) || ownerID == uuid.Nil {
		return nil, ErrConversationOwnerInvalid
	}
	if leadTimeDays <= 0 {
		leadTimeDays = 7
	}
	var existing models.RelationshipEvent
	err := s.db.WithContext(ctx).
		Where("owner_type = ? AND owner_id = ? AND type = ? AND lower(title) = lower(?)", ownerType, ownerID, eventType, title).
		Where("status <> ?", models.RelationshipEventDismissed).
		First(&existing).Error
	if err == nil {
		return &existing, nil // já existe — NOOP
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	e := &models.RelationshipEvent{
		OwnerType: ownerType, OwnerID: ownerID, RelatedPersonID: relatedPersonID,
		Type: eventType, Title: title, EventDate: eventDate, Recurring: recurring,
		LeadTimeDays: leadTimeDays, Status: models.RelationshipEventPending,
		Source: source, Notes: strings.TrimSpace(notes), AddedBy: addedBy,
	}
	if cErr := s.db.WithContext(ctx).Create(e).Error; cErr != nil {
		return nil, cErr
	}
	return e, nil
}

func (s *RelationshipEventService) UpdateStatus(ctx context.Context, id uuid.UUID, status string) error {
	switch status {
	case models.RelationshipEventPending, models.RelationshipEventAcknowledged, models.RelationshipEventDone, models.RelationshipEventDismissed:
	default:
		return ErrConversationOwnerInvalid
	}
	return s.db.WithContext(ctx).Model(&models.RelationshipEvent{}).Where("id = ?", id).Update("status", status).Error
}

func (s *RelationshipEventService) DeleteByID(ctx context.Context, id uuid.UUID) error {
	return s.db.WithContext(ctx).Where("id = ?", id).Delete(&models.RelationshipEvent{}).Error
}

// ListUpcomingAll devolve, em todos os donos, os eventos cuja PRÓXIMA ocorrência cai dentro de
// `withinDays`, ordenados por proximidade. Calcula recorrência (aniversários) em Go. Escala de
// clínica: a tabela é pequena, varrer e calcular em memória é suficiente.
func (s *RelationshipEventService) ListUpcomingAll(ctx context.Context, withinDays int, now time.Time) ([]UpcomingEvent, error) {
	if withinDays <= 0 {
		withinDays = 14
	}
	var events []models.RelationshipEvent
	if err := s.db.WithContext(ctx).
		Where("status IN (?, ?)", models.RelationshipEventPending, models.RelationshipEventAcknowledged).
		Find(&events).Error; err != nil {
		return nil, err
	}
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	var out []UpcomingEvent
	for _, e := range events {
		next := nextOccurrence(e.EventDate, e.Recurring, today)
		if next.IsZero() {
			continue
		}
		days := int(next.Sub(today).Hours() / 24)
		if days < 0 || days > withinDays {
			continue
		}
		out = append(out, UpcomingEvent{Event: e, NextDate: next, DaysUntil: days})
	}
	// Ordena por proximidade (insertion simples; lista pequena).
	for i := 1; i < len(out); i++ {
		for j := i; j > 0 && out[j].DaysUntil < out[j-1].DaysUntil; j-- {
			out[j], out[j-1] = out[j-1], out[j]
		}
	}
	return out, nil
}

// UpcomingEvent é um evento com a próxima ocorrência calculada (para job + painel).
type UpcomingEvent struct {
	Event     models.RelationshipEvent
	NextDate  time.Time
	DaysUntil int
}

// nextOccurrence calcula a próxima ocorrência de um evento. Recorrente (aniversário): a data
// mês/dia no ano corrente; se já passou, no próximo ano. Não recorrente: a própria data.
func nextOccurrence(eventDate time.Time, recurring bool, today time.Time) time.Time {
	if !recurring {
		d := time.Date(eventDate.Year(), eventDate.Month(), eventDate.Day(), 0, 0, 0, 0, today.Location())
		return d
	}
	month, day := eventDate.Month(), eventDate.Day()
	// 29/02 em ano não bissexto → 28/02.
	mkDate := func(year int) time.Time {
		d := day
		if month == time.February && day == 29 && !isLeap(year) {
			d = 28
		}
		return time.Date(year, month, d, 0, 0, 0, 0, today.Location())
	}
	cand := mkDate(today.Year())
	if cand.Before(today) {
		cand = mkDate(today.Year() + 1)
	}
	return cand
}

func isLeap(y int) bool { return (y%4 == 0 && y%100 != 0) || y%400 == 0 }
