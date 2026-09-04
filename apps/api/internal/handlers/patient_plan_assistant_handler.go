package handlers

import (
	"errors"
	"strings"

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

// AssembleDraft godoc
// @Summary Monta o rascunho da devolutiva só com código, sem custo de modelo
// @Description Mesma estrutura da geração (arco, réguas hidratadas do escore, mini-série do
// @Description histórico, plano a partir das condutas registradas, "para levar" a partir da
// @Description receita), montada pelo servidor a partir do prontuário congelado. Não chama modelo:
// @Description custo zero e resposta em milissegundos. O que fica em branco é o que é leitura
// @Description clínica (o punch, e o título como afirmação), e vem depois pela conversa.
// @Tags patient-plans
// @Accept json
// @Produce json
// @Param id path string true "ID do paciente (UUID)"
// @Param body body dto.GeneratePlanRequest false "Título opcional"
// @Success 201 {object} services.GenerateDraftResult
// @Failure 400 {object} dto.ErrorResponse
// @Failure 422 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/patients/{id}/plans/assemble [post]
func (h *PatientPlanAssistantHandler) AssembleDraft(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(dto.ErrorResponse{Error: "invalid patient id", Message: err.Error()})
	}
	userID, ok := c.Locals("userID").(uuid.UUID)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{Error: "não autenticado"})
	}
	var req dto.GeneratePlanRequest
	_ = c.BodyParser(&req)

	out, err := h.assistant.AssembleDraft(services.GenerateDraftInput{
		PatientID: patientID, UserID: userID, Title: req.Title,
	})
	if err != nil {
		if strings.Contains(err.Error(), "não tem exame nem anamnese") {
			return c.Status(fiber.StatusUnprocessableEntity).
				JSON(dto.ErrorResponse{Error: "prontuário insuficiente", Message: err.Error()})
		}
		return h.fail(c, err, "falha ao montar o rascunho")
	}
	return c.Status(fiber.StatusCreated).JSON(out)
}

// GenerateDraft godoc
// @Summary Gera o rascunho da devolutiva a partir do prontuário compilado
// @Description Cria um plano novo já ESCRITO: congela o dossiê, o modelo redige o deck seguindo o
// @Description arco e a gramática dos decks reais, e o servidor confere cada número contra o
// @Description dossiê antes de gravar. Números não encontrados viram aviso no slide exato, não
// @Description bloqueiam. A revisão fica registrada como escrita pelo assistente.
// @Tags patient-plans
// @Accept json
// @Produce json
// @Param id path string true "ID do paciente (UUID)"
// @Param body body dto.GeneratePlanRequest false "Título e instrução opcionais"
// @Success 201 {object} services.GenerateDraftResult
// @Failure 400 {object} dto.ErrorResponse
// @Failure 422 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/patients/{id}/plans/generate [post]
func (h *PatientPlanAssistantHandler) GenerateDraft(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(dto.ErrorResponse{Error: "invalid patient id", Message: err.Error()})
	}
	userID, ok := c.Locals("userID").(uuid.UUID)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{Error: "não autenticado"})
	}
	var req dto.GeneratePlanRequest
	// Corpo é opcional: gerar sem instrução é o caso comum.
	_ = c.BodyParser(&req)

	out, err := h.assistant.GenerateDraft(services.GenerateDraftInput{
		PatientID: patientID, UserID: userID,
		Title: req.Title, Instruction: req.Instruction,
	})
	if err != nil {
		// 422 e não 500 quando o paciente simplesmente não tem prontuário: não é defeito, é a
		// resposta certa, e a tela precisa dizer isso em vez de "erro".
		if strings.Contains(err.Error(), "não tem exame nem anamnese") {
			return c.Status(fiber.StatusUnprocessableEntity).
				JSON(dto.ErrorResponse{Error: "prontuário insuficiente", Message: err.Error()})
		}
		// Pelo `fail` e não com 500 fixo: chave ausente é 503, modelo fora do ar é 502, e o
		// frontend só mostra a mensagem — com 500 em tudo, falta de configuração lia como bug.
		return h.fail(c, err, "falha ao gerar o rascunho")
	}
	return c.Status(fiber.StatusCreated).JSON(out)
}
