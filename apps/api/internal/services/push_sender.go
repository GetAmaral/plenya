package services

import "github.com/google/uuid"

// CompositePushSender encadeia vários PushSender (ex: Expo mobile + Web Push).
// Cada um recebe o mesmo payload; falha em um não impede os demais. Mantém o
// contrato PushSender, então o NotificationService não muda.
type CompositePushSender struct {
	senders []PushSender
}

// NewCompositePushSender monta o fan-out, ignorando senders nil.
func NewCompositePushSender(senders ...PushSender) *CompositePushSender {
	filtered := make([]PushSender, 0, len(senders))
	for _, s := range senders {
		if s != nil {
			filtered = append(filtered, s)
		}
	}
	return &CompositePushSender{senders: filtered}
}

// Send dispara em todos os canais. Retorna o primeiro erro (best-effort).
func (c *CompositePushSender) Send(userID uuid.UUID, payload PushPayload) error {
	var firstErr error
	for _, s := range c.senders {
		if err := s.Send(userID, payload); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}
