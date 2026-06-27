package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

// WhatsAppTemplateHandler — catálogo de templates WhatsApp (admin): listar, sincronizar da Meta,
// editar metadados nossos (propósito, variáveis, cabeamento, toggle).
type WhatsAppTemplateHandler struct {
	svc *services.WhatsAppTemplateService
}

func NewWhatsAppTemplateHandler(svc *services.WhatsAppTemplateService) *WhatsAppTemplateHandler {
	return &WhatsAppTemplateHandler{svc: svc}
}

// GET /admin/whatsapp-templates
func (h *WhatsAppTemplateHandler) List(c *fiber.Ctx) error {
	items, err := h.svc.List()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": items})
}

// GET /conversations/whatsapp-templates — só os enviáveis (APPROVED + enabled), pro compositor.
func (h *WhatsAppTemplateHandler) ListSendable(c *fiber.Ctx) error {
	items, err := h.svc.ListSendable()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"data": items})
}

// POST /admin/whatsapp-templates/sync — força sincronização com a Meta.
func (h *WhatsAppTemplateHandler) Sync(c *fiber.Ctx) error {
	n, err := h.svc.SyncFromMeta(c.UserContext())
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"synced": n})
}

// PATCH /admin/whatsapp-templates/:id — edita os campos nossos.
func (h *WhatsAppTemplateHandler) Update(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "id inválido"})
	}
	var body struct {
		Purpose        *string                       `json:"purpose"`
		WiringNotes    *string                       `json:"wiringNotes"`
		Enabled        *bool                         `json:"enabled"`
		ChatSelectable *bool                         `json:"chatSelectable"`
		Variables      *[]models.WhatsAppTemplateVar `json:"variables"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "payload inválido"})
	}
	t, err := h.svc.UpdateMetaFields(id, services.UpdateTemplateInput{
		Purpose: body.Purpose, WiringNotes: body.WiringNotes, Enabled: body.Enabled,
		ChatSelectable: body.ChatSelectable, Variables: body.Variables,
	})
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(t)
}
