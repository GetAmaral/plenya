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

// DossierPerson é uma pessoa importante para exibição na tela 360.
type DossierPerson struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Relation string     `json:"relation"`
	Birthday *time.Time `json:"birthday,omitempty"`
	Notes    string     `json:"notes"`
	Source   string     `json:"source"`
}

// DossierEvent é um evento/data de relacionamento para exibição na tela 360.
type DossierEvent struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Title     string    `json:"title"`
	EventDate time.Time `json:"eventDate"`
	Recurring bool      `json:"recurring"`
	Status    string    `json:"status"`
	Source    string    `json:"source"`
}

// DossierView é a visão 360 social de uma pessoa.
type DossierView struct {
	OwnerType         string          `json:"ownerType"`
	OwnerID           string          `json:"ownerId"`
	IsPatient         bool            `json:"isPatient"`
	RollingSummary    string          `json:"rollingSummary"`
	SummaryUpdatedAt  *time.Time      `json:"summaryUpdatedAt,omitempty"`
	RelationshipStage string          `json:"relationshipStage"`
	Facts             []DossierFact   `json:"facts"`
	People            []DossierPerson `json:"people"`
	Events            []DossierEvent  `json:"events"`

	// Flags derivadas (calculadas, não digitadas) — Fase D.
	ContinuumActive       bool       `json:"continuumActive"`
	Frequent              bool       `json:"frequent"`
	AppointmentsCompleted int        `json:"appointmentsCompleted"`
	LastConsultAt         *time.Time `json:"lastConsultAt,omitempty"`
}

// frequentThreshold — nº de consultas realizadas a partir do qual a pessoa é "frequente".
const frequentThreshold = 3

// relationshipFlags são as flags derivadas (calculadas) de um dono.
type relationshipFlags struct {
	ContinuumActive       bool
	AppointmentsCompleted int
	LastConsultAt         *time.Time
}

func (f relationshipFlags) frequent() bool { return f.AppointmentsCompleted >= frequentThreshold }

// derivedRelationshipFlags calcula Continuum/assinatura ativa + consultas realizadas + última
// consulta. Só faz sentido para pacientes (leads não têm consultas/assinatura). Best-effort.
func (s *ConversationService) derivedRelationshipFlags(ctx context.Context, ownerType string, ownerID uuid.UUID) relationshipFlags {
	var f relationshipFlags
	if ownerType != string(models.ConversationOwnerPatient) {
		return f
	}

	var continuum int64
	s.db.WithContext(ctx).Model(&models.PatientContinuum{}).
		Where("patient_id = ? AND status = ?", ownerID, models.ContinuumActive).Count(&continuum)
	if continuum == 0 {
		var sub int64
		s.db.WithContext(ctx).Model(&models.PatientSubscription{}).
			Where("patient_id = ? AND status = ?", ownerID, models.SubscriptionActive).Count(&sub)
		f.ContinuumActive = sub > 0
	} else {
		f.ContinuumActive = true
	}

	var count int64
	s.db.WithContext(ctx).Model(&models.Appointment{}).
		Where("patient_id = ? AND status = ?", ownerID, models.AppointmentCompleted).Count(&count)
	f.AppointmentsCompleted = int(count)

	if count > 0 {
		var last models.Appointment
		if err := s.db.WithContext(ctx).
			Where("patient_id = ? AND status = ?", ownerID, models.AppointmentCompleted).
			Order("scheduled_at DESC").First(&last).Error; err == nil {
			t := last.ScheduledAt
			f.LastConsultAt = &t
		}
	}
	return f
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
		People:    []DossierPerson{},
		Events:    []DossierEvent{},
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

	people, err := NewRelationshipPersonService(s.db).ListByOwner(ctx, ownerType, ownerID)
	if err != nil {
		return nil, err
	}
	for _, p := range people {
		view.People = append(view.People, DossierPerson{
			ID: p.ID.String(), Name: p.Name, Relation: p.Relation,
			Birthday: p.Birthday, Notes: p.Notes, Source: p.Source,
		})
	}

	events, err := NewRelationshipEventService(s.db).ListByOwner(ctx, ownerType, ownerID)
	if err != nil {
		return nil, err
	}
	for _, e := range events {
		view.Events = append(view.Events, DossierEvent{
			ID: e.ID.String(), Type: e.Type, Title: e.Title, EventDate: e.EventDate,
			Recurring: e.Recurring, Status: e.Status, Source: e.Source,
		})
	}

	flags := s.derivedRelationshipFlags(ctx, ownerType, ownerID)
	view.ContinuumActive = flags.ContinuumActive
	view.AppointmentsCompleted = flags.AppointmentsCompleted
	view.LastConsultAt = flags.LastConsultAt
	view.Frequent = flags.frequent()
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

// ---------- Pessoas importantes (equipe) ----------

func (s *ConversationService) AddDossierPerson(ctx context.Context, ownerType string, ownerID uuid.UUID, name, relation string, birthday *time.Time, notes string, addedBy uuid.UUID) (*DossierView, error) {
	if _, err := NewRelationshipPersonService(s.db).Upsert(ctx, ownerType, ownerID, name, relation, birthday, notes, models.RelationshipSourceStaff, &addedBy); err != nil {
		return nil, err
	}
	return s.GetDossier(ctx, ownerType, ownerID)
}

func (s *ConversationService) UpdateDossierPerson(ctx context.Context, ownerType string, ownerID, personID uuid.UUID, name, relation string, birthday *time.Time, notes string) (*DossierView, error) {
	if err := NewRelationshipPersonService(s.db).UpdateByID(ctx, personID, name, relation, birthday, notes); err != nil {
		return nil, err
	}
	return s.GetDossier(ctx, ownerType, ownerID)
}

func (s *ConversationService) DeleteDossierPerson(ctx context.Context, ownerType string, ownerID, personID uuid.UUID) (*DossierView, error) {
	if err := NewRelationshipPersonService(s.db).DeleteByID(ctx, personID); err != nil {
		return nil, err
	}
	return s.GetDossier(ctx, ownerType, ownerID)
}

// ---------- Eventos / avisos (equipe) ----------

func (s *ConversationService) AddDossierEvent(ctx context.Context, ownerType string, ownerID uuid.UUID, eventType, title string, eventDate time.Time, recurring bool, leadTimeDays int, notes string, addedBy uuid.UUID) (*DossierView, error) {
	if _, err := NewRelationshipEventService(s.db).Upsert(ctx, ownerType, ownerID, nil, normalizeEventType(eventType), title, eventDate, recurring, leadTimeDays, models.RelationshipSourceStaff, notes, &addedBy); err != nil {
		return nil, err
	}
	return s.GetDossier(ctx, ownerType, ownerID)
}

func (s *ConversationService) UpdateDossierEventStatus(ctx context.Context, ownerType string, ownerID, eventID uuid.UUID, status string) (*DossierView, error) {
	if err := NewRelationshipEventService(s.db).UpdateStatus(ctx, eventID, status); err != nil {
		return nil, err
	}
	return s.GetDossier(ctx, ownerType, ownerID)
}

func (s *ConversationService) DeleteDossierEvent(ctx context.Context, ownerType string, ownerID, eventID uuid.UUID) (*DossierView, error) {
	if err := NewRelationshipEventService(s.db).DeleteByID(ctx, eventID); err != nil {
		return nil, err
	}
	return s.GetDossier(ctx, ownerType, ownerID)
}

// UpcomingEventItem é um item achatado da lista "Próximos" da Recepção (cross-owner).
type UpcomingEventItem struct {
	ID        string    `json:"id"`
	OwnerType string    `json:"ownerType"`
	OwnerID   string    `json:"ownerId"`
	OwnerName string    `json:"ownerName"`
	Type         string    `json:"type"`
	Title        string    `json:"title"`
	NextDate     time.Time `json:"nextDate"`
	DaysUntil    int       `json:"daysUntil"`
	LeadTimeDays int       `json:"leadTimeDays"`
	Status       string    `json:"status"`
}

// UpcomingEvents lista os próximos eventos de relacionamento (todos os donos) p/ o painel da Recepção.
func (s *ConversationService) UpcomingEvents(ctx context.Context, withinDays int, now time.Time) ([]UpcomingEventItem, error) {
	ups, err := NewRelationshipEventService(s.db).ListUpcomingAll(ctx, withinDays, now)
	if err != nil {
		return nil, err
	}
	out := make([]UpcomingEventItem, 0, len(ups))
	nameCache := map[string]string{}
	for _, u := range ups {
		ck := u.Event.OwnerType + ":" + u.Event.OwnerID.String()
		name, ok := nameCache[ck]
		if !ok {
			name = s.resolveOwnerName(ctx, u.Event.OwnerType, u.Event.OwnerID)
			nameCache[ck] = name
		}
		out = append(out, UpcomingEventItem{
			ID: u.Event.ID.String(), OwnerType: u.Event.OwnerType, OwnerID: u.Event.OwnerID.String(),
			OwnerName: name, Type: u.Event.Type, Title: u.Event.Title,
			NextDate: u.NextDate, DaysUntil: u.DaysUntil, LeadTimeDays: u.Event.LeadTimeDays, Status: u.Event.Status,
		})
	}
	return out, nil
}

// resolveOwnerName devolve o nome de exibição do dono (lead/paciente). Best-effort.
func (s *ConversationService) resolveOwnerName(ctx context.Context, ownerType string, ownerID uuid.UUID) string {
	switch ownerType {
	case string(models.ConversationOwnerPatient):
		var p models.Patient
		if err := s.db.WithContext(ctx).Select("name").First(&p, "id = ?", ownerID).Error; err == nil {
			return p.Name
		}
	case string(models.ConversationOwnerLead):
		var l models.Lead
		if err := s.db.WithContext(ctx).Select("name").First(&l, "id = ?", ownerID).Error; err == nil && l.Name != nil {
			return *l.Name
		}
	}
	return ""
}
