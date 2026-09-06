package handlers

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/services"
	"github.com/plenya/api/internal/utils"
)

type PaymentHandler struct {
	paymentService *services.PaymentService
	validator      *validator.Validate
}

func NewPaymentHandler(paymentService *services.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		paymentService: paymentService,
		validator:      validator.New(),
	}
}

// Create POST /api/v1/payments
func (h *PaymentHandler) Create(c *fiber.Ctx) error {
	var req dto.CreatePaymentRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid request body",
			Message: err.Error(),
		})
	}
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "validation failed",
			Details: formatValidationErrors(err),
		})
	}
	userID := middleware.GetUserID(c)
	resp, err := h.paymentService.Create(userID, &req)
	if err != nil {
		if errors.Is(err, services.ErrPaymentAlreadyExists) {
			return c.Status(fiber.StatusConflict).JSON(dto.ErrorResponse{
				Error:   "payment already exists",
				Message: "já existe um pagamento registrado para esta consulta",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(resp)
}

// List GET /api/v1/payments?patientId=&from=&to=
func (h *PaymentHandler) List(c *fiber.Ctx) error {
	var patientID *uuid.UUID
	if pid := c.Query("patientId"); pid != "" {
		if id, err := uuid.Parse(pid); err == nil {
			patientID = &id
		}
	}
	var from, to *time.Time
	if v := c.Query("from"); v != "" {
		if t, err := time.Parse(time.RFC3339, v); err == nil {
			from = &t
		}
	}
	if v := c.Query("to"); v != "" {
		if t, err := time.Parse(time.RFC3339, v); err == nil {
			to = &t
		}
	}
	resp, err := h.paymentService.List(patientID, from, to)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}
	return c.JSON(resp)
}

// Refund POST /api/v1/payments/:id/refund
func (h *PaymentHandler) Refund(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid payment id"})
	}
	var req dto.RefundPaymentRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid request body",
			Message: err.Error(),
		})
	}
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "validation failed",
			Details: formatValidationErrors(err),
		})
	}
	resp, err := h.paymentService.Refund(id, &req)
	if err != nil {
		if errors.Is(err, services.ErrPaymentNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "payment not found"})
		}
		if errors.Is(err, services.ErrPaymentAlreadyRefunded) {
			return c.Status(fiber.StatusConflict).JSON(dto.ErrorResponse{
				Error:   "payment already refunded",
				Message: "este pagamento já foi estornado",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}
	return c.JSON(resp)
}

// Receipt GET /api/v1/payments/:id/receipt → PDF
func (h *PaymentHandler) Receipt(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid payment id"})
	}
	pdf, filename, err := h.paymentService.GenerateReceipt(id)
	if err != nil {
		if errors.Is(err, services.ErrPaymentNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "payment not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}
	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", utils.ContentDisposition(filename, "recibo.pdf"))
	return c.Send(pdf)
}

// ListPrices GET /api/v1/consultation-prices
func (h *PaymentHandler) ListPrices(c *fiber.Ctx) error {
	resp, err := h.paymentService.ListPrices()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}
	return c.JSON(resp)
}

// UpsertPrice PUT /api/v1/consultation-prices
func (h *PaymentHandler) UpsertPrice(c *fiber.Ctx) error {
	var req dto.ConsultationPriceRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid request body",
			Message: err.Error(),
		})
	}
	if err := h.validator.Struct(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "validation failed",
			Details: formatValidationErrors(err),
		})
	}
	resp, err := h.paymentService.UpsertPrice(&req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}
	return c.JSON(resp)
}
