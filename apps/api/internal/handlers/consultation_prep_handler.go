package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

// ConsultationPrepHandler expõe o formulário de preparação pré-consulta e o upload de exames
// pelo paciente (portal). Todas as rotas são patient-scoped (middleware.RequirePatient).
type ConsultationPrepHandler struct {
	prep  *services.ConsultationPrepService
	score *services.AnonymousScoreService
	docs  *services.PatientDocumentsService
}

func NewConsultationPrepHandler(prep *services.ConsultationPrepService, score *services.AnonymousScoreService, docs *services.PatientDocumentsService) *ConsultationPrepHandler {
	return &ConsultationPrepHandler{prep: prep, score: score, docs: docs}
}

// GetConfig GET /api/v1/patient/me/prep/config
// Retorna o subset curado de ScoreItems (PrepOrder != nil), no mesmo formato da Triagem.
func (h *ConsultationPrepHandler) GetConfig(c *fiber.Ctx) error {
	cfg, err := h.score.BuildPrepConfig()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(cfg)
}

// GetPrep GET /api/v1/patient/me/prep?appointmentId=...
// Retorna a preparação do paciente para a consulta (ou nil quando ainda não existe).
func (h *ConsultationPrepHandler) GetPrep(c *fiber.Ctx) error {
	patientID := middleware.GetPatientID(c)
	var apptID *uuid.UUID
	if raw := c.Query("appointmentId"); raw != "" {
		id, err := uuid.Parse(raw)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "appointmentId inválido"})
		}
		apptID = &id
	}
	prep, err := h.prep.Get(patientID, apptID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if prep == nil {
		return c.JSON(nil)
	}
	view := services.ToConsultationPrepView(prep)
	return c.JSON(view)
}

// SubmitPrep POST /api/v1/patient/me/prep
// Salva rascunho (submit=false) ou finaliza (submit=true) a preparação do paciente.
func (h *ConsultationPrepHandler) SubmitPrep(c *fiber.Ctx) error {
	patientID := middleware.GetPatientID(c)
	var req dto.SubmitPrepRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "payload inválido"})
	}
	prep, err := h.prep.Submit(patientID, req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	view := services.ToConsultationPrepView(prep)
	return c.Status(fiber.StatusOK).JSON(view)
}

// UploadExam POST /api/v1/patient/me/documents (multipart)
// Permite o paciente enviar exames anteriores pelo portal (Source=portal).
func (h *ConsultationPrepHandler) UploadExam(c *fiber.Ctx) error {
	patientID := middleware.GetPatientID(c)
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "arquivo obrigatório (campo 'file')"})
	}
	title := c.FormValue("title")
	if title == "" {
		title = file.Filename
	}
	docType := c.FormValue("type", string(models.DocumentTypeReport))
	doc, err := h.docs.Create(services.CreateDocumentInput{
		PatientID:   patientID,
		UploadedBy:  middleware.GetUserID(c),
		Source:      models.DocumentSourcePortal,
		Type:        models.PatientDocumentType(docType),
		Title:       title,
		Description: c.FormValue("description"),
		File:        file,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(doc)
}
