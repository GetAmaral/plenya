package handlers

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/pdfdoc"
	"github.com/plenya/api/internal/services"
)

// PatientPlanHandler — plano de paciente (a devolutiva). Por ora expõe só o dossiê: o insumo
// derivado do prontuário para montar o plano sem recomeçar do zero.
type PatientPlanHandler struct {
	dossier   *services.PatientPlanDossierService
	plans     *services.PatientPlanService
	validator *validator.Validate
}

func NewPatientPlanHandler(dossier *services.PatientPlanDossierService, plans *services.PatientPlanService) *PatientPlanHandler {
	return &PatientPlanHandler{dossier: dossier, plans: plans, validator: validator.New()}
}

// GetDossier godoc
// @Summary Dossiê do plano de paciente
// @Description Insumo derivado do prontuário para montar a devolutiva: uma régua por exame (escala
// @Description do escore aplicável ao paciente + histórico dele), achados separados em "está bem" e
// @Description "está se movendo" e ordenados por pontos perdidos, mais plano de cuidado, último
// @Description pedido de exames e prescrições ativas.
// @Tags patient-plans
// @Accept json
// @Produce json
// @Param id path string true "ID do paciente (UUID)"
// @Success 200 {object} dto.PlanDossierResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/patients/{id}/plan-dossier [get]
func (h *PatientPlanHandler) GetDossier(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid patient id", Message: err.Error()})
	}
	out, err := h.dossier.Build(patientID)
	if err != nil {
		if errors.Is(err, services.ErrPatientNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "paciente não encontrado", Message: err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "falha ao montar o dossiê do plano", Message: err.Error()})
	}
	return c.JSON(out)
}

// ---- Plano de devolutiva ----

// List godoc
// @Summary Lista os planos de devolutiva do paciente
// @Tags patient-plans
// @Produce json
// @Param id path string true "ID do paciente (UUID)"
// @Success 200 {array} dto.PatientPlanResponse
// @Failure 400 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/patients/{id}/plans [get]
func (h *PatientPlanHandler) List(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid patient id", Message: err.Error()})
	}
	out, err := h.plans.ListByPatient(patientID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "falha ao listar planos", Message: err.Error()})
	}
	return c.JSON(out)
}

// Get godoc
// @Summary Carrega um plano de devolutiva
// @Tags patient-plans
// @Produce json
// @Param id path string true "ID do paciente (UUID)"
// @Param planId path string true "ID do plano (UUID)"
// @Success 200 {object} dto.PatientPlanResponse
// @Failure 404 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/patients/{id}/plans/{planId} [get]
func (h *PatientPlanHandler) Get(c *fiber.Ctx) error {
	patientID, planID, resp, ok := h.ids(c)
	if !ok {
		return resp
	}
	out, err := h.plans.Get(planID, patientID)
	if err != nil {
		return h.fail(c, err, "falha ao carregar o plano")
	}
	return c.JSON(out)
}

// Create godoc
// @Summary Cria um plano de devolutiva
// @Tags patient-plans
// @Accept json
// @Produce json
// @Param id path string true "ID do paciente (UUID)"
// @Param request body dto.SavePatientPlanRequest true "Título e slides"
// @Success 201 {object} dto.PatientPlanResponse
// @Failure 400 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/patients/{id}/plans [post]
func (h *PatientPlanHandler) Create(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid patient id", Message: err.Error()})
	}
	var req dto.SavePatientPlanRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid request body", Message: err.Error()})
	}
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "validation failed", Details: formatValidationErrors(err)})
	}
	out, err := h.plans.Create(patientID, middleware.GetUserID(c), &req)
	if err != nil {
		return h.fail(c, err, "falha ao criar o plano")
	}
	return c.Status(fiber.StatusCreated).JSON(out)
}

// Update godoc
// @Summary Reescreve um plano (volta para rascunho)
// @Tags patient-plans
// @Accept json
// @Produce json
// @Param id path string true "ID do paciente (UUID)"
// @Param planId path string true "ID do plano (UUID)"
// @Param request body dto.SavePatientPlanRequest true "Título e slides"
// @Success 200 {object} dto.PatientPlanResponse
// @Failure 404 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/patients/{id}/plans/{planId} [put]
func (h *PatientPlanHandler) Update(c *fiber.Ctx) error {
	patientID, planID, resp, ok := h.ids(c)
	if !ok {
		return resp
	}
	var req dto.SavePatientPlanRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid request body", Message: err.Error()})
	}
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "validation failed", Details: formatValidationErrors(err)})
	}
	out, err := h.plans.Update(planID, patientID, middleware.GetUserID(c), &req)
	if err != nil {
		return h.fail(c, err, "falha ao salvar o plano")
	}
	return c.JSON(out)
}

// Delete godoc
// @Summary Apaga um plano
// @Tags patient-plans
// @Param id path string true "ID do paciente (UUID)"
// @Param planId path string true "ID do plano (UUID)"
// @Success 204
// @Failure 404 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/patients/{id}/plans/{planId} [delete]
func (h *PatientPlanHandler) Delete(c *fiber.Ctx) error {
	patientID, planID, resp, ok := h.ids(c)
	if !ok {
		return resp
	}
	if err := h.plans.Delete(planID, patientID); err != nil {
		return h.fail(c, err, "falha ao apagar o plano")
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// Preview godoc
// @Summary Prévia do deck em HTML
// @Description O MESMO HTML que vira os dois PDFs. É o que a tela de montagem mostra e a base da tela do portal.
// @Tags patient-plans
// @Produce html
// @Param id path string true "ID do paciente (UUID)"
// @Param planId path string true "ID do plano (UUID)"
// @Success 200 {string} string "HTML do deck"
// @Failure 404 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/patients/{id}/plans/{planId}/preview [get]
func (h *PatientPlanHandler) Preview(c *fiber.Ctx) error {
	patientID, planID, resp, ok := h.ids(c)
	if !ok {
		return resp
	}
	html, err := h.plans.Preview(planID, patientID)
	if err != nil {
		return h.fail(c, err, "falha ao montar a prévia")
	}
	c.Set(fiber.HeaderContentType, "text/html; charset=utf-8")
	return c.SendString(html)
}

// CheckOverflow godoc
// @Summary Mede quais slides transbordam
// @Description Lista vazia = pode publicar. O slide tem altura fixa e overflow:hidden, então
// @Description conteúdo demais some do PDF sem erro nenhum; esta medição é o que impede isso.
// @Tags patient-plans
// @Produce json
// @Param id path string true "ID do paciente (UUID)"
// @Param planId path string true "ID do plano (UUID)"
// @Success 200 {object} map[string][]pdfdoc.DeckOverflow
// @Failure 404 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/patients/{id}/plans/{planId}/overflow [get]
func (h *PatientPlanHandler) CheckOverflow(c *fiber.Ctx) error {
	patientID, planID, resp, ok := h.ids(c)
	if !ok {
		return resp
	}
	over, err := h.plans.CheckOverflow(planID, patientID)
	if err != nil {
		return h.fail(c, err, "falha ao medir os slides")
	}
	if over == nil {
		over = []pdfdoc.DeckOverflow{}
	}
	return c.JSON(fiber.Map{"slides": over})
}

// Publish godoc
// @Summary Publica o plano no portal do paciente
// @Description Gera o PDF 16:9 (apresentar/mandar) e o A4 paisagem (imprimir) a partir do MESMO
// @Description conteúdo e publica os dois. Recusa com 422 se algum slide transbordar.
// @Tags patient-plans
// @Produce json
// @Param id path string true "ID do paciente (UUID)"
// @Param planId path string true "ID do plano (UUID)"
// @Success 200 {object} dto.PatientPlanResponse
// @Failure 422 {object} dto.PatientPlanOverflowResponse
// @Failure 404 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/patients/{id}/plans/{planId}/publish [post]
func (h *PatientPlanHandler) Publish(c *fiber.Ctx) error {
	patientID, planID, resp, ok := h.ids(c)
	if !ok {
		return resp
	}
	out, err := h.plans.Publish(planID, patientID, middleware.GetUserID(c))
	if err != nil {
		// Slide que transborda não é erro de servidor: é conteúdo que não cabe, e quem escreveu
		// precisa saber exatamente onde para poder cortar.
		var overflow *services.ErrPatientPlanOverflow
		if errors.As(err, &overflow) {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.PatientPlanOverflowResponse{
				Error:  "conteúdo não cabe no slide",
				Slides: overflow.Slides,
			})
		}
		return h.fail(c, err, "falha ao publicar o plano")
	}
	return c.JSON(out)
}

// ids extrai os dois ids do path. Quando algum é inválido JÁ RESPONDE 400 e devolve ok=false; o
// chamador só precisa devolver `resp`.
//
// O `ok` existe porque `c.Status(...).JSON(...)` devolve **nil** quando consegue escrever a
// resposta. Usar esse retorno como sinal de erro (o jeito que parece óbvio) faz o handler seguir em
// frente com uuid.Nil depois de já ter escrito um 400 — e o que chega ao cliente é um 404 de
// "não encontrado" no lugar do 400 de "id inválido".
func (h *PatientPlanHandler) ids(c *fiber.Ctx) (patientID, planID uuid.UUID, resp error, ok bool) {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return uuid.Nil, uuid.Nil, c.Status(fiber.StatusBadRequest).
			JSON(dto.ErrorResponse{Error: "invalid patient id", Message: err.Error()}), false
	}
	planID, err = uuid.Parse(c.Params("planId"))
	if err != nil {
		return uuid.Nil, uuid.Nil, c.Status(fiber.StatusBadRequest).
			JSON(dto.ErrorResponse{Error: "invalid plan id", Message: err.Error()}), false
	}
	return patientID, planID, nil, true
}

func (h *PatientPlanHandler) fail(c *fiber.Ctx, err error, msg string) error {
	switch {
	case errors.Is(err, services.ErrPatientPlanNotFound):
		return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "plano não encontrado", Message: err.Error()})
	case errors.Is(err, services.ErrPatientNotFound):
		return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "paciente não encontrado", Message: err.Error()})
	case errors.Is(err, services.ErrPatientPlanEmpty):
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "plano vazio", Message: err.Error()})
	case errors.Is(err, services.ErrPlanRevisionConflict):
		// 409 e não 500: a tela precisa distinguir "deu erro" de "alguém escreveu antes de você",
		// que tem tratamento próprio (recarregar e mostrar o que mudou).
		return c.Status(fiber.StatusConflict).JSON(dto.ErrorResponse{Error: "o plano mudou", Message: err.Error()})
	}
	return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: msg, Message: err.Error()})
}

// PublishReport godoc
// @Summary Publica o relatório A4 assinado do plano
// @Description Terceiro modo do MESMO conteúdo: os slides achatados no documento fluido da
// @Description papelaria, assinado com ICP-Brasil. O deck 16:9/A4 é peça de comunicação e não leva
// @Description assinatura; este é o documento clínico.
// @Tags patient-plans
// @Produce json
// @Param id path string true "ID do paciente (UUID)"
// @Param planId path string true "ID do plano (UUID)"
// @Success 201 {object} map[string]string
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/patients/{id}/plans/{planId}/report [post]
func (h *PatientPlanHandler) PublishReport(c *fiber.Ctx) error {
	patientID, planID, resp, ok := h.ids(c)
	if !ok {
		return resp
	}
	docID, err := h.plans.PublishReport(planID, patientID, middleware.GetUserID(c))
	if err != nil {
		return h.fail(c, err, "falha ao publicar o relatório")
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"issuedDocumentId": docID})
}

// ---- Portal do paciente ----

// MyPlans godoc
// @Summary Meus planos de devolutiva (portal)
// @Description Só os publicados. Rascunho é trabalho em andamento do médico.
// @Tags patient-portal
// @Produce json
// @Success 200 {array} dto.PatientPlanResponse
// @Security BearerAuth
// @Router /api/v1/patient/me/plans [get]
func (h *PatientPlanHandler) MyPlans(c *fiber.Ctx) error {
	out, err := h.plans.ListPublished(middleware.GetPatientID(c))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "falha ao listar", Message: err.Error()})
	}
	return c.JSON(out)
}

// MyPlan godoc
// @Summary Um plano publicado meu (portal)
// @Tags patient-portal
// @Produce json
// @Param id path string true "ID do plano (UUID)"
// @Success 200 {object} dto.PatientPlanResponse
// @Failure 404 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/patient/me/plans/{id} [get]
func (h *PatientPlanHandler) MyPlan(c *fiber.Ctx) error {
	planID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid plan id", Message: err.Error()})
	}
	out, gErr := h.plans.GetPublished(planID, middleware.GetPatientID(c))
	if gErr != nil {
		return h.fail(c, gErr, "falha ao carregar o plano")
	}
	return c.JSON(out)
}
