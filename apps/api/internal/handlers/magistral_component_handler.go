package handlers

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/services"
)

// MagistralComponentHandler expõe o catálogo magistral: busca, curadoria e a checagem de fórmula
// (compatibilidade + tamanho de cápsula) que a tela de manipulado chama enquanto o médico digita.
type MagistralComponentHandler struct {
	service   *services.MagistralComponentService
	validator *validator.Validate
}

func NewMagistralComponentHandler(service *services.MagistralComponentService) *MagistralComponentHandler {
	return &MagistralComponentHandler{service: service, validator: validator.New()}
}

// Search godoc
// @Summary Search magistral components
// @Description Busca componentes do catálogo magistral por nome ou sinônimo (mín. 2 caracteres)
// @Tags MagistralComponents
// @Produce json
// @Param q query string true "Termo de busca"
// @Param limit query int false "Limite" default(20)
// @Security BearerAuth
// @Success 200 {array} models.MagistralComponent
// @Failure 500 {object} dto.ErrorResponse
// @Router /magistral-components/search [get]
func (h *MagistralComponentHandler) Search(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "20"))
	items, err := h.service.Search(c.Query("q"), limit)
	if err != nil {
		return internalError(c, err)
	}
	return c.JSON(items)
}

// Get godoc
// @Summary Get magistral component
// @Tags MagistralComponents
// @Produce json
// @Param id path string true "Component UUID"
// @Security BearerAuth
// @Success 200 {object} models.MagistralComponent
// @Failure 404 {object} dto.ErrorResponse
// @Router /magistral-components/{id} [get]
func (h *MagistralComponentHandler) Get(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid id"})
	}
	item, err := h.service.GetByID(id)
	if err != nil {
		if errors.Is(err, services.ErrMagistralComponentNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "component not found"})
		}
		return internalError(c, err)
	}
	return c.JSON(item)
}

// Create godoc
// @Summary Create magistral component
// @Tags MagistralComponents
// @Accept json
// @Produce json
// @Param body body dto.MagistralComponentRequest true "Component"
// @Security BearerAuth
// @Success 201 {object} models.MagistralComponent
// @Failure 400 {object} dto.ErrorResponse
// @Router /magistral-components [post]
func (h *MagistralComponentHandler) Create(c *fiber.Ctx) error {
	var req dto.MagistralComponentRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid request body"})
	}
	// As tags de validação do DTO existiam e nunca rodavam: dava para criar componente sem nome.
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	item, err := h.service.Upsert(nil, &req, middleware.GetUserID(c))
	if err != nil {
		return internalError(c, err)
	}
	return c.Status(fiber.StatusCreated).JSON(item)
}

// Update godoc
// @Summary Update magistral component
// @Tags MagistralComponents
// @Accept json
// @Produce json
// @Param id path string true "Component UUID"
// @Param body body dto.MagistralComponentRequest true "Component"
// @Security BearerAuth
// @Success 200 {object} models.MagistralComponent
// @Failure 404 {object} dto.ErrorResponse
// @Router /magistral-components/{id} [put]
func (h *MagistralComponentHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid id"})
	}
	var req dto.MagistralComponentRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid request body"})
	}
	// As tags de validação do DTO existiam e nunca rodavam: dava para criar componente sem nome.
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	item, err := h.service.Upsert(&id, &req, middleware.GetUserID(c))
	if err != nil {
		if errors.Is(err, services.ErrMagistralComponentNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "component not found"})
		}
		return internalError(c, err)
	}
	return c.JSON(item)
}

// SaveDefaultDose godoc
// @Summary Save default dose for a substance
// @Description Curadoria oportunista: guarda a dose que o médico acabou de prescrever como padrão da substância
// @Tags MagistralComponents
// @Accept json
// @Produce json
// @Param body body dto.MagistralDefaultDoseRequest true "Substância e dose"
// @Security BearerAuth
// @Success 200 {object} models.MagistralComponent
// @Failure 400 {object} dto.ErrorResponse
// @Router /magistral-components/default-dose [post]
func (h *MagistralComponentHandler) SaveDefaultDose(c *fiber.Ctx) error {
	var req dto.MagistralDefaultDoseRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid request body"})
	}
	// As tags de validação do DTO existiam e nunca rodavam: dava para criar componente sem nome.
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	item, err := h.service.SaveDefaultDose(req.Substance, req.Dose, req.Unit, middleware.GetUserID(c))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	return c.JSON(item)
}

// Check godoc
// @Summary Check magistral formula
// @Description Avisos de compatibilidade/palatabilidade e cálculo de tamanho de cápsula. Avisa, não bloqueia.
// @Tags MagistralComponents
// @Accept json
// @Produce json
// @Param body body dto.MagistralCheckRequest true "Fórmula"
// @Security BearerAuth
// @Success 200 {object} dto.MagistralCheckResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /magistral-components/check [post]
func (h *MagistralComponentHandler) Check(c *fiber.Ctx) error {
	var req dto.MagistralCheckRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid request body"})
	}
	resp, err := h.service.CheckFormula(&req)
	if err != nil {
		return internalError(c, err)
	}
	return c.JSON(resp)
}

// ListEvidence godoc
// @Summary List RAG evidence for a component
// @Description Trechos de aulas/artigos ligados ao componente. Material de leitura; nenhum cálculo usa isto.
// @Tags MagistralComponents
// @Produce json
// @Param id path string true "Component UUID"
// @Security BearerAuth
// @Success 200 {array} models.MagistralComponentArticle
// @Router /magistral-components/{id}/evidence [get]
func (h *MagistralComponentHandler) ListEvidence(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid id"})
	}
	items, err := h.service.ListEvidence(id)
	if err != nil {
		return internalError(c, err)
	}
	return c.JSON(items)
}

// PinEvidence godoc
// @Summary Pin or unpin an evidence excerpt
// @Tags MagistralComponents
// @Produce json
// @Param id path string true "Evidence UUID"
// @Param pinned query bool false "true para fixar, false para soltar" default(true)
// @Security BearerAuth
// @Success 204 "No Content"
// @Router /magistral-components/evidence/{id}/pin [post]
func (h *MagistralComponentHandler) PinEvidence(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid id"})
	}
	pinned := c.Query("pinned", "true") == "true"
	if err := h.service.PinEvidence(id, pinned, middleware.GetUserID(c)); err != nil {
		return internalError(c, err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// ConfirmEvidence godoc
// @Summary Confirm the RAG-extracted content of a component
// @Description Aceite do médico: o que veio do RAG deixa de ser sugestão e passa a conferido.
// @Tags MagistralComponents
// @Produce json
// @Param id path string true "Component UUID"
// @Security BearerAuth
// @Success 204 "No Content"
// @Router /magistral-components/{id}/confirm [post]
func (h *MagistralComponentHandler) ConfirmEvidence(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid id"})
	}
	if err := h.service.ConfirmEvidence(id, middleware.GetUserID(c)); err != nil {
		return internalError(c, err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// ListIncompatibilities godoc
// @Summary List curated incompatibilities
// @Tags MagistralComponents
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.MagistralIncompatibility
// @Router /magistral-components/incompatibilities [get]
func (h *MagistralComponentHandler) ListIncompatibilities(c *fiber.Ctx) error {
	items, err := h.service.ListIncompatibilities()
	if err != nil {
		return internalError(c, err)
	}
	return c.JSON(items)
}

// CreateIncompatibility godoc
// @Summary Create curated incompatibility
// @Tags MagistralComponents
// @Accept json
// @Produce json
// @Param body body dto.MagistralIncompatibilityRequest true "Par"
// @Security BearerAuth
// @Success 201 {object} models.MagistralIncompatibility
// @Failure 400 {object} dto.ErrorResponse
// @Router /magistral-components/incompatibilities [post]
func (h *MagistralComponentHandler) CreateIncompatibility(c *fiber.Ctx) error {
	var req dto.MagistralIncompatibilityRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid request body"})
	}
	// As tags de validação do DTO existiam e nunca rodavam: dava para criar componente sem nome.
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	item, err := h.service.CreateIncompatibility(&req, middleware.GetUserID(c))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(item)
}

// DeleteIncompatibility godoc
// @Summary Delete curated incompatibility
// @Tags MagistralComponents
// @Produce json
// @Param id path string true "Incompatibility UUID"
// @Security BearerAuth
// @Success 204 "No Content"
// @Router /magistral-components/incompatibilities/{id} [delete]
func (h *MagistralComponentHandler) DeleteIncompatibility(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid id"})
	}
	if err := h.service.DeleteIncompatibility(id); err != nil {
		return internalError(c, err)
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func internalError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
		Error:   "internal server error",
		Message: err.Error(),
	})
}
