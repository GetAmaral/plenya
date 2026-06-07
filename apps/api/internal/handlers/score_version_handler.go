package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/services"
)

// ScoreVersionHandler — versões públicas do escore (Triagem/Light) + geração do config do site.
type ScoreVersionHandler struct {
	service  *services.ScoreVersionService
	validate *validator.Validate
}

func NewScoreVersionHandler(service *services.ScoreVersionService, validate *validator.Validate) *ScoreVersionHandler {
	return &ScoreVersionHandler{service: service, validate: validate}
}

// List GET /api/v1/score-versions
func (h *ScoreVersionHandler) List(c *fiber.Ctx) error {
	out, err := h.service.List()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(out)
}

// Get GET /api/v1/score-versions/:id (com itens)
func (h *ScoreVersionHandler) Get(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	out, err := h.service.GetWithItems(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(out)
}

// Create POST /api/v1/score-versions
func (h *ScoreVersionHandler) Create(c *fiber.Ctx) error {
	var dto services.CreateScoreVersionDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	if err := h.validate.Struct(dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	out, err := h.service.Create(dto)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(out)
}

// Update PUT /api/v1/score-versions/:id
func (h *ScoreVersionHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var dto services.UpdateScoreVersionDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	out, err := h.service.Update(id, dto)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(out)
}

// Delete DELETE /api/v1/score-versions/:id
func (h *ScoreVersionHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := h.service.Delete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// SetItems PUT /api/v1/score-versions/:id/items — substitui os itens da version.
func (h *ScoreVersionHandler) SetItems(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var dto services.SetVersionItemsDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	if err := h.service.SetItems(id, dto.Items); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// GetConfig GET /api/v1/score-light/config/:slug — PÚBLICO. Config estático para o site (gerador).
// Só serve versões context='public' (Triagem/Light). Versões de preparação pré-consulta
// (context='patient_prep') são privadas e atendidas pelo endpoint do paciente logado.
func (h *ScoreVersionHandler) GetConfig(c *fiber.Ctx) error {
	slug := c.Params("slug")
	v, err := h.service.GetBySlug(slug)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	if v.Context != "public" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "score version not found"})
	}
	out, err := h.service.BuildConfig(slug)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(out)
}
