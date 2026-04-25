package handlers

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/services"
)

// CampaignHandler — endpoints autenticados (admin) de campanhas de marketing.
type CampaignHandler struct {
	service   *services.CampaignService
	validator *validator.Validate
}

func NewCampaignHandler(service *services.CampaignService) *CampaignHandler {
	return &CampaignHandler{
		service:   service,
		validator: validator.New(),
	}
}

// List GET /campaigns?includeArchived=true
//
// @Summary Listar campanhas
// @Tags    campaigns
// @Produce json
// @Param   includeArchived query bool false "Inclui arquivadas (default: false)"
// @Success 200 {array} services.CampaignDTO
// @Router  /campaigns [get]
func (h *CampaignHandler) List(c *fiber.Ctx) error {
	includeArchived := c.QueryBool("includeArchived", false)
	campaigns, err := h.service.List(includeArchived)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "list campaigns: " + err.Error()})
	}
	return c.JSON(campaigns)
}

// Get GET /campaigns/:id
//
// @Summary Detalhe de campanha
// @Tags    campaigns
// @Produce json
// @Param   id path string true "Campaign ID (uuid)"
// @Success 200 {object} services.CampaignDTO
// @Router  /campaigns/{id} [get]
func (h *CampaignHandler) Get(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	dto, err := h.service.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "campaign not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(dto)
}

// Create POST /campaigns
//
// @Summary Criar campanha
// @Tags    campaigns
// @Accept  json
// @Produce json
// @Param   body body services.CreateCampaignInput true "Dados"
// @Success 201 {object} services.CampaignDTO
// @Router  /campaigns [post]
func (h *CampaignHandler) Create(c *fiber.Ctx) error {
	var in services.CreateCampaignInput
	if err := c.BodyParser(&in); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	if err := h.validator.Struct(in); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{"error": "validation failed", "details": err.Error()})
	}
	actor := middleware.GetUserID(c)
	var actorPtr *uuid.UUID
	if actor != uuid.Nil {
		actorPtr = &actor
	}
	dto, err := h.service.Create(in, actorPtr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(dto)
}

// Update PATCH /campaigns/:id
//
// @Summary Atualizar campanha
// @Tags    campaigns
// @Accept  json
// @Produce json
// @Param   id   path string                          true "Campaign ID"
// @Param   body body services.UpdateCampaignInput    true "Patch"
// @Success 200 {object} services.CampaignDTO
// @Router  /campaigns/{id} [patch]
func (h *CampaignHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var in services.UpdateCampaignInput
	if err := c.BodyParser(&in); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	if err := h.validator.Struct(in); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{"error": "validation failed", "details": err.Error()})
	}
	dto, err := h.service.Update(id, in)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "campaign not found"})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(dto)
}

// Archive POST /campaigns/:id/archive
//
// @Summary Arquivar campanha
// @Tags    campaigns
// @Param   id path string true "Campaign ID"
// @Success 204
// @Router  /campaigns/{id}/archive [post]
func (h *CampaignHandler) Archive(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := h.service.Archive(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// Delete DELETE /campaigns/:id
//
// @Summary Remover campanha (soft delete)
// @Tags    campaigns
// @Param   id path string true "Campaign ID"
// @Success 204
// @Router  /campaigns/{id} [delete]
func (h *CampaignHandler) Delete(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := h.service.Delete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// QRCode GET /campaigns/:id/qrcode?size=512
//
// Retorna PNG do QR code da URL final da campanha.
//
// @Summary QR code da URL final da campanha
// @Tags    campaigns
// @Produce image/png
// @Param   id   path  string true  "Campaign ID"
// @Param   size query int    false "Tamanho em px (128–1024, default 512)"
// @Success 200 {file} png
// @Router  /campaigns/{id}/qrcode [get]
func (h *CampaignHandler) QRCode(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	size := c.QueryInt("size", 512)
	png, filename, err := h.service.QRCodePNG(id, size)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "campaign not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	c.Set("Content-Type", "image/png")
	c.Set("Content-Disposition", `inline; filename="`+filename+`"`)
	return c.Send(png)
}
