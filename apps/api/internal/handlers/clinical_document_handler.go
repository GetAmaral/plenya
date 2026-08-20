package handlers

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/services"
)

// ClinicalDocumentHandler — "Enviar por WhatsApp" de documentos clínicos ao paciente (pedido de
// exames, emitido, receita, doc de prontuário), sem o round-trip de baixar+reanexar. Dois modos:
// arquivo (mídia inline, exige janela 24h) e link (template documento_disponivel, reabre conversa).
type ClinicalDocumentHandler struct {
	svc         *services.ClinicalDocumentService
	shareSecret []byte
	shareTTL    time.Duration
}

func NewClinicalDocumentHandler(svc *services.ClinicalDocumentService, shareSecret string, shareTTL time.Duration) *ClinicalDocumentHandler {
	if shareTTL <= 0 {
		shareTTL = 30 * 24 * time.Hour
	}
	return &ClinicalDocumentHandler{svc: svc, shareSecret: []byte(shareSecret), shareTTL: shareTTL}
}

// mintShareURL monta o link público do documento com o MESMO claim (doc-share/did) que o endpoint
// público GET /documents/shared/:token valida.
func (h *ClinicalDocumentHandler) mintShareURL(baseURL string, docID uuid.UUID) (string, error) {
	now := time.Now()
	claims := docShareClaims{
		DocumentID: docID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "doc-share",
			ExpiresAt: jwt.NewNumericDate(now.Add(h.shareTTL)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
	signed, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(h.shareSecret)
	if err != nil {
		return "", err
	}
	return baseURL + "/api/v1/documents/shared/" + signed, nil
}

type sendDocWhatsAppRequest struct {
	DocType string `json:"docType"` // lab_request | issued_document | prescription | patient_document
	DocID   string `json:"docId"`
	Mode    string `json:"mode"` // file | link
}

// SendWhatsApp: POST /patients/:id/clinical-documents/send-whatsapp
func (h *ClinicalDocumentHandler) SendWhatsApp(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "patient id inválido"})
	}
	var body sendDocWhatsAppRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "payload inválido"})
	}
	docID, err := uuid.Parse(body.DocID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "docId inválido"})
	}
	if body.Mode != "file" && body.Mode != "link" {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "mode deve ser file ou link"})
	}
	docType := services.ClinicalDocType(body.DocType)
	actorID := middleware.GetUserID(c)

	doc, err := h.svc.ResolveToPatientDocument(patientID, docType, docID, actorID)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrClinicalDocUnknownType):
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
		case errors.Is(err, services.ErrClinicalDocNoFile):
			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.ErrorResponse{Error: err.Error()})
		default:
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: err.Error()})
		}
	}

	switch body.Mode {
	case "file":
		activity, serr := h.svc.SendFile(c.UserContext(), patientID, doc, actorID)
		if serr != nil {
			return sendDocError(c, serr)
		}
		return c.Status(fiber.StatusCreated).JSON(activity)
	default: // link
		url, merr := h.mintShareURL(c.BaseURL(), doc.ID)
		if merr != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "falha ao gerar link"})
		}
		activity, serr := h.svc.SendLink(c.UserContext(), patientID, services.DocLabel(docType), url, actorID)
		if serr != nil {
			return sendDocError(c, serr)
		}
		return c.Status(fiber.StatusCreated).JSON(activity)
	}
}

type sendDocEmailRequest struct {
	DocType string `json:"docType"` // lab_request | issued_document | prescription | patient_document
	DocID   string `json:"docId"`
}

// SendEmail: POST /patients/:id/clinical-documents/send-email
//
// Manda o LINK seguro do documento por e-mail. Vai link e não anexo porque o PDF traz dado clínico
// identificável, e-mail comum atravessa servidor que não controlamos e fica na caixa do paciente
// para sempre. É o mesmo link assinado e com prazo que o envio por WhatsApp usa.
//
// @Summary Envia documento clínico por e-mail
// @Tags clinical-documents
// @Security BearerAuth
// @Param id path string true "Patient ID"
// @Param request body sendDocEmailRequest true "documento a enviar"
// @Success 201 {object} models.LeadActivity
// @Failure 422 {object} dto.ErrorResponse
// @Router /patients/{id}/clinical-documents/send-email [post]
func (h *ClinicalDocumentHandler) SendEmail(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "patient id inválido"})
	}
	var body sendDocEmailRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "payload inválido"})
	}
	docID, err := uuid.Parse(body.DocID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "docId inválido"})
	}
	docType := services.ClinicalDocType(body.DocType)
	actorID := middleware.GetUserID(c)

	doc, err := h.svc.ResolveToPatientDocument(patientID, docType, docID, actorID)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrClinicalDocUnknownType):
			return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: err.Error()})
		case errors.Is(err, services.ErrClinicalDocNoFile):
			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.ErrorResponse{Error: err.Error()})
		default:
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: err.Error()})
		}
	}

	url, merr := h.mintShareURL(c.BaseURL(), doc.ID)
	if merr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "falha ao gerar link"})
	}

	atividade, serr := h.svc.SendEmail(patientID, services.DocLabel(docType), url, h.shareTTL, actorID)
	if serr != nil {
		if errors.Is(serr, services.ErrClinicalDocNoEmail) {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.ErrorResponse{
				Error: "paciente não tem e-mail cadastrado", Message: "no_email",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: serr.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(atividade)
}

// sendDocError mapeia os erros de envio do ConversationService. Janela fechada vira 422 com
// code=window_closed pra o front cair no modo link.
func sendDocError(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, services.ErrConversationWindowClosed):
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.ErrorResponse{
			Error: "janela de 24h fechada — use o link", Message: "window_closed",
		})
	case errors.Is(err, services.ErrConversationNoChannel):
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.ErrorResponse{Error: "paciente sem telefone WhatsApp"})
	case errors.Is(err, services.ErrConversationOptedOut):
		return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.ErrorResponse{Error: "paciente com opt-out de WhatsApp"})
	default:
		return c.Status(fiber.StatusBadGateway).JSON(dto.ErrorResponse{Error: "falha ao enviar", Message: err.Error()})
	}
}

// DocumentChannels: GET /patients/:id/clinical-documents/channels
//
// Diz por onde dá para mandar documento a este paciente. O botão nasce desabilitado com o motivo
// em vez de o médico clicar e tomar um erro.
func (h *ClinicalDocumentHandler) DocumentChannels(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "patient id inválido"})
	}
	st, err := h.svc.DocumentChannels(patientID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	return c.JSON(st)
}

// ListDocuments: GET /patients/:id/clinical-documents — picker "Anexar arquivo do EMR".
func (h *ClinicalDocumentHandler) ListDocuments(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "patient id inválido"})
	}
	items, err := h.svc.ListForPatient(patientID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: err.Error()})
	}
	return c.JSON(fiber.Map{"data": items})
}
