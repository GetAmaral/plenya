package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
	"github.com/plenya/api/internal/utils"
)

type PrescriptionHandler struct {
	prescriptionService    *services.PrescriptionService
	prescriptionPDFService *services.PrescriptionPDFService
	validator              *validator.Validate
}

func NewPrescriptionHandler(
	prescriptionService *services.PrescriptionService,
	prescriptionPDFService *services.PrescriptionPDFService,
) *PrescriptionHandler {
	return &PrescriptionHandler{
		prescriptionService:    prescriptionService,
		prescriptionPDFService: prescriptionPDFService,
		validator:              validator.New(),
	}
}

// Create godoc
// @Summary Create prescription
// @Description Creates a new prescription
// @Tags Prescriptions
// @Accept json
// @Produce json
// @Param body body dto.CreatePrescriptionRequest true "Prescription data"
// @Security BearerAuth
// @Success 201 {object} dto.PrescriptionResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /prescriptions [post]
func (h *PrescriptionHandler) Create(c *fiber.Ctx) error {
	var req dto.CreatePrescriptionRequest
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

	doctorID := middleware.GetUserID(c)

	resp, err := h.prescriptionService.Create(doctorID, &req)
	if err != nil {
		if errors.Is(err, services.ErrPatientNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "patient not found",
				Message: "no patient found with the given id",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

// GetByID godoc
// @Summary Get prescription by ID
// @Description Returns a single prescription
// @Tags Prescriptions
// @Produce json
// @Param id path string true "Prescription UUID"
// @Security BearerAuth
// @Success 200 {object} dto.PrescriptionResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /prescriptions/{id} [get]
func (h *PrescriptionHandler) GetByID(c *fiber.Ctx) error {
	prescriptionID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid prescription id",
			Message: "prescription id must be a valid UUID",
		})
	}

	userID := middleware.GetUserID(c)
	userRole := middleware.GetPrimaryRole(c)

	resp, err := h.prescriptionService.GetByID(prescriptionID, userID, userRole)
	if err != nil {
		if errors.Is(err, services.ErrPrescriptionNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "prescription not found",
				Message: "no prescription found with the given id",
			})
		}
		if errors.Is(err, services.ErrUnauthorized) {
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
				Error:   "unauthorized",
				Message: "you do not have permission to view this prescription",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(resp)
}

// List godoc
// @Summary List prescriptions
// @Description Lists prescriptions (optionally filtered by patient/status)
// @Tags Prescriptions
// @Produce json
// @Param patientId query string false "Filter by patient UUID"
// @Param status query string false "Filter by status"
// @Param limit query int false "Page size"
// @Param offset query int false "Offset"
// @Security BearerAuth
// @Success 200 {array} dto.PrescriptionResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /prescriptions [get]
func (h *PrescriptionHandler) List(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "20"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))
	if limit > 100 {
		limit = 100
	}

	userID := middleware.GetUserID(c)
	userRole := middleware.GetPrimaryRole(c)

	// Parse optional filters
	var patientID *uuid.UUID
	var status *models.PrescriptionStatus

	if patientIDStr := c.Query("patientId"); patientIDStr != "" {
		pid, err := uuid.Parse(patientIDStr)
		if err == nil {
			patientID = &pid
		}
	}
	if statusStr := c.Query("status"); statusStr != "" {
		s := models.PrescriptionStatus(statusStr)
		status = &s
	}

	resp, err := h.prescriptionService.List(userID, userRole, patientID, status, limit, offset)
	if err != nil {
		if errors.Is(err, services.ErrUnauthorized) {
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
				Error:   "unauthorized",
				Message: "you do not have permission to list prescriptions",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(resp)
}

// Update godoc
// @Summary Update prescription
// @Description Updates an existing prescription
// @Tags Prescriptions
// @Accept json
// @Produce json
// @Param id path string true "Prescription UUID"
// @Param body body dto.UpdatePrescriptionRequest true "Fields to update"
// @Security BearerAuth
// @Success 200 {object} dto.PrescriptionResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 409 {object} dto.ErrorResponse "receita assinada não pode ser alterada"
// @Failure 500 {object} dto.ErrorResponse
// @Router /prescriptions/{id} [put]
func (h *PrescriptionHandler) Update(c *fiber.Ctx) error {
	prescriptionID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid prescription id",
			Message: "prescription id must be a valid UUID",
		})
	}

	var req dto.UpdatePrescriptionRequest
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

	doctorID := middleware.GetUserID(c)
	userRole := middleware.GetPrimaryRole(c)

	resp, err := h.prescriptionService.Update(prescriptionID, doctorID, userRole, &req)
	if err != nil {
		if errors.Is(err, services.ErrPrescriptionNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "prescription not found",
				Message: "no prescription found with the given id",
			})
		}
		if errors.Is(err, services.ErrUnauthorized) {
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
				Error:   "unauthorized",
				Message: "you do not have permission to update this prescription",
			})
		}
		if errors.Is(err, services.ErrPrescriptionSigned) {
			return c.Status(fiber.StatusConflict).JSON(dto.ErrorResponse{
				Error:   "prescription already signed",
				Message: "receita assinada não pode ser alterada; duplique e emita uma nova",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.JSON(resp)
}

// Delete godoc
// @Summary Delete prescription draft
// @Description Soft-deletes a prescription. Signed prescriptions are refused (409).
// @Tags Prescriptions
// @Produce json
// @Param id path string true "Prescription UUID"
// @Security BearerAuth
// @Success 204 "No Content"
// @Failure 400 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 409 {object} dto.ErrorResponse "receita assinada não pode ser excluída"
// @Failure 500 {object} dto.ErrorResponse
// @Router /prescriptions/{id} [delete]
func (h *PrescriptionHandler) Delete(c *fiber.Ctx) error {
	prescriptionID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid prescription id",
			Message: "prescription id must be a valid UUID",
		})
	}

	doctorID := middleware.GetUserID(c)
	userRole := middleware.GetPrimaryRole(c)

	err = h.prescriptionService.Delete(prescriptionID, doctorID, userRole)
	if err != nil {
		if errors.Is(err, services.ErrPrescriptionNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "prescription not found",
				Message: "no prescription found with the given id",
			})
		}
		if errors.Is(err, services.ErrUnauthorized) {
			return c.Status(fiber.StatusForbidden).JSON(dto.ErrorResponse{
				Error:   "unauthorized",
				Message: "you do not have permission to delete this prescription",
			})
		}
		if errors.Is(err, services.ErrPrescriptionSigned) {
			return c.Status(fiber.StatusConflict).JSON(dto.ErrorResponse{
				Error:   "prescription already signed",
				Message: "receita assinada não pode ser excluída; use o status cancelada",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// SignAndGenerate godoc
// @Summary Sign prescription and generate signed PDF
// @Description Signs prescription with doctor's digital certificate and generates PDF with QR code
// @Tags Prescriptions
// @Produce json
// @Param id path string true "Prescription UUID"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "signedPdfUrl, sncrNumber"
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /prescriptions/{id}/sign [post]
func (h *PrescriptionHandler) SignAndGenerate(c *fiber.Ctx) error {
	prescriptionID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid prescription id",
			Message: "prescription id must be a valid UUID",
		})
	}

	doctorID := middleware.GetUserID(c)

	// Validar que prescrição pertence ao médico
	var prescription models.Prescription
	if err := h.prescriptionPDFService.GetDB().
		Where("id = ? AND doctor_id = ?", prescriptionID, doctorID).
		First(&prescription).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
				Error:   "prescription not found",
				Message: "no prescription found with the given id or you don't have permission",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "internal server error",
			Message: err.Error(),
		})
	}

	// Gerar PDF assinado
	pdfURL, err := h.prescriptionPDFService.GenerateSignedPrescriptionPDF(prescriptionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error:   "failed to generate signed PDF",
			Message: err.Error(),
		})
	}

	// Buscar prescrição atualizada para pegar o modo de assinatura efetivo.
	h.prescriptionPDFService.GetDB().First(&prescription, prescriptionID)

	mode := ""
	if prescription.SignatureMode != nil {
		mode = *prescription.SignatureMode
	}
	// Mensagem para o frontend orientar o médico no modo manual (controlado ou sem certificado).
	message := "Prescrição assinada digitalmente (ICP-Brasil)."
	if mode == "manual" {
		message = "Receita gerada para impressão. Imprima, carimbe e assine à mão."
	}

	return c.JSON(fiber.Map{
		"signedPdfUrl":  pdfURL,
		"signatureMode": mode,
		"message":       message,
		"sncrNumber":    prescription.SNCRNumber,
	})
}

// Download godoc
// @Summary Download autenticado do PDF da prescrição
// @Description Entrega o PDF da prescrição via PatientDocument (substitui o /uploads estático).
// @Tags Prescriptions
// @Produce application/pdf
// @Param id path string true "Prescription UUID"
// @Security BearerAuth
// @Success 200 {file} file
// @Failure 404 {object} dto.ErrorResponse
// @Router /prescriptions/{id}/download [get]
// Preview godoc
// @Summary Rascunho da receita, sem assinar
// @Description PDF da receita como ela vai sair, no layout de assinatura manual. Não grava nada e
// @Description não marca a receita como assinada: é o que o médico confere ANTES de assinar.
// @Tags Prescriptions
// @Produce application/pdf
// @Param id path string true "Prescription UUID"
// @Success 200 {file} binary
// @Failure 404 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/prescriptions/{id}/preview [get]
func (h *PrescriptionHandler) Preview(c *fiber.Ctx) error {
	prescriptionID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid prescription id",
			Message: "prescription id must be a valid UUID",
		})
	}
	pdf, fileName, err := h.prescriptionPDFService.PreviewPrescriptionPDF(prescriptionID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
			Error:   "rascunho indisponível",
			Message: err.Error(),
		})
	}
	c.Set("Content-Type", "application/pdf")
	// Mesmo padrão de nome dos documentos assinados. O rascunho também é salvo e reenviado.
	c.Set("Content-Disposition", utils.ContentDisposition(fileName, "receita-rascunho.pdf"))
	return c.Send(pdf)
}

func (h *PrescriptionHandler) Download(c *fiber.Ctx) error {
	prescriptionID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{
			Error:   "invalid prescription id",
			Message: "prescription id must be a valid UUID",
		})
	}

	full, fileName, contentType, err := h.prescriptionPDFService.GetForDownload(prescriptionID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{
			Error:   "pdf indisponível",
			Message: err.Error(),
		})
	}

	c.Set("Content-Type", contentType)
	// Mesmo padrão do pedido de exames: filename= em ASCII puro para cliente antigo e
	// filename*=UTF-8 com o nome real acentuado. Sem isso, "José" quebrava o cabeçalho.
	c.Set("Content-Disposition", utils.ContentDisposition(fileName, "receita.pdf"))
	return c.SendFile(full)
}

// ValidatePublic godoc
// @Summary Validate prescription (public - no auth)
// @Description Public endpoint for validating prescription authenticity via QR code
// @Tags Prescriptions
// @Produce json
// @Param id path string true "Prescription UUID"
// @Success 200 {object} map[string]interface{} "validation result with prescription details"
// @Failure 404 {object} map[string]interface{}
// @Router /prescriptions/validate/{id} [get]
func (h *PrescriptionHandler) ValidatePublic(c *fiber.Ctx) error {
	prescriptionID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"valid": false,
			"error": "invalid prescription id",
		})
	}

	var prescription models.Prescription
	if err := h.prescriptionPDFService.GetDB().
		Preload("Patient").
		Preload("Doctor").
		Preload("Medications").
		First(&prescription, prescriptionID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"valid": false,
			"error": "prescription not found",
		})
	}

	// Verificar validade
	isExpired := time.Now().After(prescription.ValidUntil)
	isUsed := prescription.IsUsed

	sigMode := ""
	if prescription.SignatureMode != nil {
		sigMode = *prescription.SignatureMode
	}

	// Ler o PDF uma vez: confere integridade (re-hash) e, no modo digital, valida a
	// assinatura PAdES de verdade (cadeia/integridade criptográfica), não só o hash.
	pdfIntact := true
	var sigVerification *services.SignatureVerification
	if prescription.SignedPDFPath != nil && prescription.SignedPDFHash != nil {
		pdfBytes, readErr := h.prescriptionPDFService.ReadSignedPDF(&prescription)
		if readErr == nil {
			currentHash := sha256.Sum256(pdfBytes)
			if hex.EncodeToString(currentHash[:]) != *prescription.SignedPDFHash {
				pdfIntact = false
			}
			if sigMode == "digital" {
				sigVerification, _ = h.prescriptionPDFService.VerifySignature(pdfBytes)
			}
		}
	}

	// No modo digital, a validade exige assinatura PAdES íntegra. No modo manual, a
	// validade jurídica vem da assinatura/carimbo físico (não há assinatura digital a aferir).
	signatureValid := true
	if sigVerification != nil {
		signatureValid = sigVerification.Valid
	}

	// Converter medications para resposta
	medications := make([]fiber.Map, len(prescription.Medications))
	for i, med := range prescription.Medications {
		medications[i] = fiber.Map{
			"id":               med.ID,
			"name":             med.MedicationName,
			"activeIngredient": med.ActiveIngredient,
			"category":         med.Category,
			"concentration":    med.Concentration,
			"quantity":         med.Quantity,
			"quantityInWords":  med.QuantityInWords,
		}
	}

	return c.JSON(fiber.Map{
		"valid":         !isExpired && !isUsed && pdfIntact && signatureValid,
		"pdfIntact":     pdfIntact,
		"signatureMode": sigMode,
		"prescription": fiber.Map{
			"id":               prescription.ID,
			"prescriptionDate": prescription.PrescriptionDate,
			"validUntil":       prescription.ValidUntil,
			"isExpired":        isExpired,
			"isUsed":           isUsed,
			"sncrNumber":       prescription.SNCRNumber,
			"medicationCount":  len(prescription.Medications),
		},
		"patient": fiber.Map{
			"name": prescription.Patient.Name,
			"cpf":  maskCPF(prescription.Patient.CPF),
		},
		"doctor": fiber.Map{
			"name": prescription.Doctor.Name,
			"crm":  formatCRM(prescription.Doctor.CRM, prescription.Doctor.CRMUF),
		},
		"medications": medications,
		"signature": fiber.Map{
			"mode":              sigMode,
			"signedAt":          prescription.SignedAt,
			"certificateSerial": prescription.CertificateSerial,
			// Não expõe URL do PDF no validador público (download é autenticado, evita
			// vazar o PDF do paciente sem auth). A validação confere autenticidade, não baixa.
			"signedPdfUrl": nil,
			"verification": sigVerification,
		},
	})
}

// Helper functions

func maskCPF(cpf *string) string {
	if cpf == nil || len(*cpf) != 11 {
		return "***.***.***-**"
	}
	c := *cpf
	return fmt.Sprintf("***.***.%s-%s", c[6:9], c[9:11])
}

func formatCRM(crm *string, uf *string) string {
	if crm == nil {
		return "CRM não informado"
	}
	if uf == nil {
		return fmt.Sprintf("CRM %s", *crm)
	}
	return fmt.Sprintf("CRM-%s %s", *uf, *crm)
}
