package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/services"
)

// DocumentShareHandler — link PÚBLICO por documento (sem portal). Um token JWT assinado escopa o
// acesso a UM documento específico e expira; o endpoint público verifica o token e serve o PDF
// assinado inline (abre direto no navegador). Pensado para enviar ao paciente por WhatsApp.
type DocumentShareHandler struct {
	documents *services.PatientDocumentsService
	secret    []byte
	ttl       time.Duration
}

// NewDocumentShareHandler — ttl<=0 cai em 30 dias (documento clínico costuma ser guardado).
func NewDocumentShareHandler(documents *services.PatientDocumentsService, secret string, ttl time.Duration) *DocumentShareHandler {
	if ttl <= 0 {
		ttl = 30 * 24 * time.Hour
	}
	return &DocumentShareHandler{documents: documents, secret: []byte(secret), ttl: ttl}
}

type docShareClaims struct {
	DocumentID string `json:"did"`
	jwt.RegisteredClaims
}

func (h *DocumentShareHandler) mintToken(docID uuid.UUID) (string, time.Time, error) {
	exp := time.Now().Add(h.ttl)
	claims := docShareClaims{
		DocumentID: docID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "doc-share",
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	signed, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(h.secret)
	return signed, exp, err
}

// Mint (staff): POST /patients/:id/documents/:docId/share-link → gera o link público do documento.
func (h *DocumentShareHandler) Mint(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "patient id inválido"})
	}
	docID, err := uuid.Parse(c.Params("docId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "document id inválido"})
	}
	// Confirma que o documento pertence ao paciente antes de gerar o link.
	if _, _, err := h.documents.GetForDownload(patientID, docID); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "documento não encontrado"})
	}
	token, exp, err := h.mintToken(docID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "falha ao gerar link"})
	}
	return c.JSON(fiber.Map{
		"url":       c.BaseURL() + "/api/v1/documents/shared/" + token,
		"expiresAt": exp,
	})
}

// Serve (público): GET /documents/shared/:token → serve o PDF assinado inline.
func (h *DocumentShareHandler) Serve(c *fiber.Ctx) error {
	claims := &docShareClaims{}
	tok, err := jwt.ParseWithClaims(c.Params("token"), claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("alg inesperado")
		}
		return h.secret, nil
	})
	if err != nil || !tok.Valid || claims.Subject != "doc-share" {
		return c.Status(fiber.StatusUnauthorized).SendString("Link inválido ou expirado.")
	}
	docID, err := uuid.Parse(claims.DocumentID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Link inválido.")
	}
	doc, full, err := h.documents.GetForShare(docID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Documento não encontrado.")
	}
	c.Set("Content-Type", doc.ContentType)
	c.Set("Content-Disposition", `inline; filename="`+doc.FileName+`"`)
	return c.SendFile(full)
}
