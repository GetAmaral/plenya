package handlers

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/services"
)

// O assistente da devolutiva: conversa, sugestões e aceite.
//
// A rota de mensagem é síncrona e leva de dez a vinte segundos. É por isso que ela aceita
// `clientMessageId`: fechar a aba depois de o modelo responder e reenviar não pode duplicar o turno.
type PatientPlanAssistantHandler struct {
	assistant *services.PatientPlanAssistantService
	validator *validator.Validate
}

func NewPatientPlanAssistantHandler(a *services.PatientPlanAssistantService, v *validator.Validate) *PatientPlanAssistantHandler {
	return &PatientPlanAssistantHandler{assistant: a, validator: v}
}

func (h *PatientPlanAssistantHandler) ids(c *fiber.Ctx) (uuid.UUID, uuid.UUID, error) {
	pid, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return uuid.Nil, uuid.Nil, errors.New("id de paciente inválido")
	}
	plan, err := uuid.Parse(c.Params("planId"))
	if err != nil {
		return uuid.Nil, uuid.Nil, errors.New("id de plano inválido")
	}
	return pid, plan, nil
}

func (h *PatientPlanAssistantHandler) fail(c *fiber.Ctx, err error, msg string) error {
	switch {
	case errors.Is(err, services.ErrPatientPlanNotFound):
		return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "plano não encontrado", Message: err.Error()})
	case errors.Is(err, services.ErrPlanDossierNotFound):
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "sem dossiê congelado", Message: err.Error()})
	case errors.Is(err, services.ErrPlanRevisionConflict):
		return c.Status(fiber.StatusConflict).JSON(dto.ErrorResponse{Error: "o plano mudou", Message: err.Error()})
	case errors.Is(err, services.ErrPlanAssistantDuplicate):
		return c.Status(fiber.StatusConflict).JSON(dto.ErrorResponse{Error: "turno duplicado", Message: err.Error()})
	case errors.Is(err, services.ErrAINotConfigured):
		return c.Status(fiber.StatusServiceUnavailable).JSON(dto.ErrorResponse{Error: "assistente indisponível", Message: err.Error()})
	case errors.Is(err, services.ErrAIUpstream), errors.Is(err, services.ErrAITruncated):
		return c.Status(fiber.StatusBadGateway).JSON(dto.ErrorResponse{Error: "falha na chamada do assistente", Message: err.Error()})
	}
	return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: msg, Message: err.Error()})
}

// SendMessage godoc
// @Summary Um turno da conversa que edita o rascunho
// @Description Alteração de TEXTO entra direto e reversível; alteração que toque número, unidade,
// @Description dose ou régua vira SUGESTÃO com a origem do número anexada. A classificação sai do
// @Description diff de numerais, não do nome do campo.
// @Tags patient-plans
// @Accept json
// @Produce json
// @Success 200 {object} dto.PlanAssistantTurn
// @Failure 409 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/patients/{id}/plans/{planId}/assistant/messages [post]
func (h *PatientPlanAssistantHandler) SendMessage(c *fiber.Ctx) error {
	patientID, planID, err := h.ids(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	var req dto.SendPlanMessageRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "corpo inválido", Message: err.Error()})
	}
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "validação falhou", Details: formatValidationErrors(err)})
	}
	out, err := h.assistant.SendMessage(services.SendPlanMessageInput{
		PlanID: planID, PatientID: patientID, UserID: middleware.GetUserID(c),
		Body: req.Body, ClientMessageID: req.ClientMessageID, ExpectedRevision: req.ExpectedRevision,
	})
	if err != nil {
		return h.fail(c, err, "falha no turno do assistente")
	}
	return c.JSON(out)
}

// History godoc
// @Summary A conversa do plano
// @Tags patient-plans
// @Produce json
// @Success 200 {array} dto.PlanMessage
// @Security BearerAuth
// @Router /api/v1/patients/{id}/plans/{planId}/assistant/messages [get]
func (h *PatientPlanAssistantHandler) History(c *fiber.Ctx) error {
	patientID, planID, err := h.ids(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	out, err := h.assistant.History(planID, patientID)
	if err != nil {
		return h.fail(c, err, "falha ao ler a conversa")
	}
	return c.JSON(out)
}

// ListSuggestions godoc
// @Summary As sugestões pendentes do plano
// @Tags patient-plans
// @Produce json
// @Success 200 {array} dto.PlanSuggestion
// @Security BearerAuth
// @Router /api/v1/patients/{id}/plans/{planId}/suggestions [get]
func (h *PatientPlanAssistantHandler) ListSuggestions(c *fiber.Ctx) error {
	patientID, planID, err := h.ids(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	out, err := h.assistant.ListSuggestions(planID, patientID)
	if err != nil {
		return h.fail(c, err, "falha ao listar as sugestões")
	}
	return c.JSON(out)
}

// ResolveSuggestions godoc
// @Summary Aceita ou recusa sugestões
// @Description Resultado parcial é resposta legítima: sugestão cujo slide mudou depois não é
// @Description aplicada, e volta em `skipped` com o motivo. É isso que impede o painel de vinte
// @Description minutos atrás de apagar o que o médico acabou de escrever.
// @Tags patient-plans
// @Accept json
// @Produce json
// @Success 200 {object} dto.PlanResolveResult
// @Security BearerAuth
// @Router /api/v1/patients/{id}/plans/{planId}/suggestions/resolve [post]
func (h *PatientPlanAssistantHandler) ResolveSuggestions(c *fiber.Ctx) error {
	patientID, planID, err := h.ids(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	var req dto.ResolveSuggestionsRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "corpo inválido", Message: err.Error()})
	}
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "validação falhou", Details: formatValidationErrors(err)})
	}
	ids := make([]uuid.UUID, 0, len(req.SuggestionIDs))
	for _, s := range req.SuggestionIDs {
		if id, err := uuid.Parse(s); err == nil {
			ids = append(ids, id)
		}
	}
	out, err := h.assistant.ResolveSuggestions(services.ResolveSuggestionsInput{
		PlanID: planID, PatientID: patientID, UserID: middleware.GetUserID(c),
		Action: req.Action, SuggestionIDs: ids, SlideID: req.SlideID,
		ExpectedRevision: req.ExpectedRevision,
	})
	if err != nil {
		return h.fail(c, err, "falha ao resolver as sugestões")
	}
	// 207 quando parte foi pulada: o cliente precisa distinguir "tudo entrou" de "entrou o que deu".
	if len(out.Skipped) > 0 {
		return c.Status(fiber.StatusMultiStatus).JSON(out)
	}
	return c.JSON(out)
}
