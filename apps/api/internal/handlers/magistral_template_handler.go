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

// MagistralTemplateHandler expõe as fórmulas-base e a sugestão de dose.
//
// A sugestão é um endpoint SEPARADO da criação de receita, de propósito: quem prescreve nunca
// recebe dose calculada por tabela; recebe um número com a base escrita e decide.
type MagistralTemplateHandler struct {
	service   *services.MagistralTemplateService
	validator *validator.Validate
}

func NewMagistralTemplateHandler(service *services.MagistralTemplateService) *MagistralTemplateHandler {
	return &MagistralTemplateHandler{service: service, validator: validator.New()}
}

// List godoc
// @Summary List magistral formula templates
// @Description Fórmulas-base ativas. `q` busca por nome ou indicação.
// @Tags MagistralTemplates
// @Produce json
// @Param q query string false "Busca por nome ou indicação"
// @Security BearerAuth
// @Success 200 {array} models.MagistralFormulaTemplate
// @Router /magistral-templates [get]
func (h *MagistralTemplateHandler) List(c *fiber.Ctx) error {
	items, err := h.service.List(c.Query("q"), c.QueryInt("limit", 30))
	if err != nil {
		return internalError(c, err)
	}
	return c.JSON(items)
}

// Get godoc
// @Summary Get a formula template
// @Tags MagistralTemplates
// @Produce json
// @Param id path string true "Template UUID"
// @Security BearerAuth
// @Success 200 {object} models.MagistralFormulaTemplate
// @Failure 404 {object} dto.ErrorResponse
// @Router /magistral-templates/{id} [get]
func (h *MagistralTemplateHandler) Get(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid id"})
	}
	item, err := h.service.GetByID(id)
	if err != nil {
		if errors.Is(err, services.ErrTemplateNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "template not found"})
		}
		return internalError(c, err)
	}
	return c.JSON(item)
}

// Create godoc
// @Summary Create a formula template
// @Tags MagistralTemplates
// @Accept json
// @Produce json
// @Param body body dto.FormulaTemplateRequest true "Fórmula-base"
// @Security BearerAuth
// @Success 201 {object} models.MagistralFormulaTemplate
// @Failure 400 {object} dto.ErrorResponse
// @Router /magistral-templates [post]
func (h *MagistralTemplateHandler) Create(c *fiber.Ctx) error {
	var req dto.FormulaTemplateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid request body"})
	}
	// As tags de validação do DTO existiam e nunca rodavam: dava para criar componente sem nome.
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	item, err := h.service.Save(nil, &req, middleware.GetUserID(c))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(item)
}

// Update godoc
// @Summary Update a formula template
// @Tags MagistralTemplates
// @Accept json
// @Produce json
// @Param id path string true "Template UUID"
// @Param body body dto.FormulaTemplateRequest true "Fórmula-base"
// @Security BearerAuth
// @Success 200 {object} models.MagistralFormulaTemplate
// @Failure 404 {object} dto.ErrorResponse
// @Router /magistral-templates/{id} [put]
func (h *MagistralTemplateHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid id"})
	}
	var req dto.FormulaTemplateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid request body"})
	}
	// As tags de validação do DTO existiam e nunca rodavam: dava para criar componente sem nome.
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	item, err := h.service.Save(&id, &req, middleware.GetUserID(c))
	if err != nil {
		if errors.Is(err, services.ErrTemplateNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "template not found"})
		}
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	return c.JSON(item)
}

// Delete godoc
// @Summary Delete a formula template
// @Tags MagistralTemplates
// @Produce json
// @Param id path string true "Template UUID"
// @Security BearerAuth
// @Success 204 "No Content"
// @Router /magistral-templates/{id} [delete]
func (h *MagistralTemplateHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid id"})
	}
	if err := h.service.Delete(id); err != nil {
		return internalError(c, err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// Suggest godoc
// @Summary Suggest doses for a patient
// @Description Aplica as regras da fórmula-base ao paciente e devolve SUGESTÕES com a base escrita. Não escreve na receita.
// @Tags MagistralTemplates
// @Produce json
// @Param id path string true "Template UUID"
// @Param patientId query string true "Patient UUID"
// @Security BearerAuth
// @Success 200 {object} dto.DoseSuggestionResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /magistral-templates/{id}/suggest [get]
func (h *MagistralTemplateHandler) Suggest(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid id"})
	}
	patientID, err := uuid.Parse(c.Query("patientId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "patientId inválido"})
	}
	resp, err := h.service.Suggest(id, patientID, middleware.GetUserID(c))
	if err != nil {
		if errors.Is(err, services.ErrTemplateNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "template not found"})
		}
		if errors.Is(err, services.ErrPatientNotSelected) {
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
				Error: "a sugestão só responde sobre o paciente selecionado",
			})
		}
		return internalError(c, err)
	}
	return c.JSON(resp)
}
