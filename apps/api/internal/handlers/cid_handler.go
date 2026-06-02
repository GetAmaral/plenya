package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/services"
)

type CIDHandler struct {
	service *services.CIDService
}

func NewCIDHandler(service *services.CIDService) *CIDHandler {
	return &CIDHandler{service: service}
}

// Search GET /api/v1/cid-codes?q=&limit=
func (h *CIDHandler) Search(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "20"))
	rows, err := h.service.Search(c.Query("q"), limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "internal server error", Message: err.Error()})
	}
	return c.JSON(rows)
}
