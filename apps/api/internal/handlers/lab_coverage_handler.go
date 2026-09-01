package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/services"
)

// LabCoverageHandler — quando o paciente fez cada exame.
type LabCoverageHandler struct {
	coverage *services.LabCoverageService
}

func NewLabCoverageHandler(coverage *services.LabCoverageService) *LabCoverageHandler {
	return &LabCoverageHandler{coverage: coverage}
}

// GetCoverage godoc
// @Summary Quando o paciente fez cada exame
// @Description Resolve o painel pelos analitos FILHOS: o laboratório não reporta "hemograma
// @Description completo", reporta hemoglobina e plaquetas. Cruzar o protocolo olhando só o painel
// @Description diz "nunca feito" e manda repetir exame de quem acabou de fazer.
// @Tags lab-coverage
// @Produce json
// @Param id path string true "ID do paciente (UUID)"
// @Param onlyRequestable query bool false "Só o que dá para pedir (padrão: true)"
// @Param doneOnly query bool false "Só o que já foi feito (padrão: false)"
// @Success 200 {object} dto.LabCoverageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Security BearerAuth
// @Router /api/v1/patients/{id}/lab-coverage [get]
func (h *LabCoverageHandler) GetCoverage(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "invalid patient id", Message: err.Error()})
	}
	onlyRequestable := c.Query("onlyRequestable") != "false"
	doneOnly := c.Query("doneOnly") == "true"
	out, err := h.coverage.Build(patientID, onlyRequestable, doneOnly)
	if err != nil {
		if errors.Is(err, services.ErrPatientNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "paciente não encontrado", Message: err.Error()})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "falha ao montar a cobertura de exames", Message: err.Error()})
	}
	return c.JSON(out)
}
