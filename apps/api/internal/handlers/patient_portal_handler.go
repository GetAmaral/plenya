// Package handlers — PatientPortalHandler expõe endpoints da área do paciente:
//   - auth: login email+senha, magic link, esqueci-senha, consume invite
//   - convite: equipe cria/lista convites pra um Patient
//   - me: dados do paciente logado
//
// Convenções de rota:
//
//	POST /api/v1/auth/patient/login                — público
//	POST /api/v1/auth/patient/magic-link           — público (envia)
//	POST /api/v1/auth/patient/magic-link/consume   — público (troca por JWT)
//	POST /api/v1/auth/patient/invite/consume       — público (troca por JWT)
//	POST /api/v1/auth/patient/forgot               — público (alias de magic-link)
//
//	POST /api/v1/patients/:id/portal-invite        — staff (admin/manager/secretary)
//	GET  /api/v1/patients/:id/portal-invite        — staff (status do último convite)
//
//	GET  /api/v1/patient/me                        — paciente
//	POST /api/v1/patient/me/password               — paciente
package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

type PatientPortalHandler struct {
	svc       *services.PatientPortalService
	dashboard *services.PatientDashboardService
	continuum *services.PatientContinuumService
}

func NewPatientPortalHandler(
	svc *services.PatientPortalService,
	dashboard *services.PatientDashboardService,
	continuum *services.PatientContinuumService,
) *PatientPortalHandler {
	return &PatientPortalHandler{svc: svc, dashboard: dashboard, continuum: continuum}
}

// WithContinuum injeta a dep de PatientContinuumService quando ele só fica
// disponível mais tarde no setup (ordem de DI no main).
func (h *PatientPortalHandler) WithContinuum(c *services.PatientContinuumService) *PatientPortalHandler {
	h.continuum = c
	return h
}

// MyContinuum GET /api/v1/patient/me/continuum
//
// Retorna o Continuum ativo do paciente (o mais recente). Inclui plano integrado
// markdown + items ordenados. Snapshot do template + revisões do plano NÃO entram
// (são internas da equipe).
func (h *PatientPortalHandler) MyContinuum(c *fiber.Ctx) error {
	patientID := middleware.GetPatientID(c)
	rows, err := h.continuum.GetByPatient(patientID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	// Ativo > pausado > completado > cancelado (mais "atual")
	var active *uuid.UUID
	for i := range rows {
		if rows[i].Status == "active" {
			active = &rows[i].ID
			break
		}
	}
	if active == nil && len(rows) > 0 {
		// Sem ativo, devolve o mais recente.
		active = &rows[0].ID
	}
	if active == nil {
		return c.JSON(fiber.Map{"continuum": nil})
	}
	full, err := h.continuum.GetByID(*active)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	// Sanitiza: zera snapshot bruto (só usa name extraída) + remove campos de "equipe"
	out := fiber.Map{
		"id":                       full.ID,
		"templateName":             services.ExtractTemplateName(full.TemplateSnapshot),
		"status":                   full.Status,
		"startDate":                full.StartDate,
		"endDate":                  full.EndDate,
		"integratedPlanMarkdown":   full.IntegratedPlanMarkdown,
		"integratedPlanUpdatedAt":  full.IntegratedPlanUpdatedAt,
		"items":                    full.Items,
		"coordinatorDoctor":        full.CoordinatorDoctor,
	}
	return c.JSON(fiber.Map{"continuum": out})
}


// Dashboard GET /api/v1/patient/me/dashboard
func (h *PatientPortalHandler) Dashboard(c *fiber.Ctx) error {
	patientID := middleware.GetPatientID(c)
	out, err := h.dashboard.Build(patientID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(out)
}

// ============================================================
// Auth público
// ============================================================

type loginPasswordBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *PatientPortalHandler) LoginPassword(c *fiber.Ctx) error {
	var body loginPasswordBody
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	resp, err := h.svc.LoginPassword(body.Email, body.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}

type magicLinkBody struct {
	Email string `json:"email"`
}

func (h *PatientPortalHandler) RequestMagicLink(c *fiber.Ctx) error {
	var body magicLinkBody
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	// Anti-enumeração: sempre 204 mesmo se email não existir
	if err := h.svc.RequestMagicLink(body.Email, nil); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

type consumeTokenBody struct {
	Token string `json:"token"`
}

func (h *PatientPortalHandler) ConsumeMagicLink(c *fiber.Ctx) error {
	var body consumeTokenBody
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	resp, err := h.svc.ConsumeMagicLink(body.Token)
	if err != nil {
		if errors.Is(err, services.ErrInviteNotFound) || errors.Is(err, services.ErrInviteExpired) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "link inválido ou expirado"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}

func (h *PatientPortalHandler) ConsumeInvite(c *fiber.Ctx) error {
	var body consumeTokenBody
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	resp, err := h.svc.ConsumeInvite(body.Token)
	if err != nil {
		if errors.Is(err, services.ErrInviteNotFound) || errors.Is(err, services.ErrInviteExpired) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "convite inválido ou expirado"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(resp)
}

// ============================================================
// Convite (staff)
// ============================================================

type createInviteBody struct {
	SendEmail bool `json:"sendEmail"`
	SendWA    bool `json:"sendWA"`
}

// CreateInvite POST /api/v1/patients/:id/portal-invite
func (h *PatientPortalHandler) CreateInvite(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid patient id"})
	}
	var body createInviteBody
	if err := c.BodyParser(&body); err != nil {
		// Default: só email
		body.SendEmail = true
	}
	if !body.SendEmail && !body.SendWA {
		body.SendEmail = true
	}
	invitedBy := middleware.GetUserID(c)
	res, err := h.svc.CreateInvite(services.CreateInviteInput{
		PatientID: patientID,
		InvitedBy: invitedBy,
		SendEmail: body.SendEmail,
		SendWA:    body.SendWA,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(res)
}

// GetInviteStatus GET /api/v1/patients/:id/portal-invite
func (h *PatientPortalHandler) GetInviteStatus(c *fiber.Ctx) error {
	patientID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid patient id"})
	}
	invite, err := h.svc.GetInviteStatus(patientID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if invite == nil {
		return c.JSON(fiber.Map{"status": "none"})
	}
	status := "pending"
	if invite.AcceptedAt != nil {
		status = "accepted"
	} else if !invite.IsValid() {
		status = "expired"
	}
	return c.JSON(fiber.Map{
		"status":     status,
		"id":         invite.ID,
		"createdAt":  invite.CreatedAt,
		"expiresAt":  invite.ExpiresAt,
		"acceptedAt": invite.AcceptedAt,
	})
}

// ============================================================
// /patient/me — endpoints autenticados pelo middleware RequirePatient
// ============================================================

// Me GET /api/v1/patient/me — retorna dados básicos do paciente logado
func (h *PatientPortalHandler) Me(c *fiber.Ctx) error {
	patient := middleware.GetPatient(c)
	userID := middleware.GetUserID(c)
	return c.JSON(fiber.Map{
		"userId":  userID,
		"patient": patient,
	})
}

type setPasswordBody struct {
	Password string `json:"password"`
}

// SetPassword POST /api/v1/patient/me/password — define ou troca a senha (sem exigir atual)
//
// Pra trocar com senha atual em mãos, usar AuthService.ChangePassword (existente, exposto
// em outra rota staff). Esta rota é defensiva: paciente acabou de entrar via magic link
// ou consume invite e quer definir uma senha.
func (h *PatientPortalHandler) SetPassword(c *fiber.Ctx) error {
	var body setPasswordBody
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	userID := middleware.GetUserID(c)
	if err := h.svc.SetPassword(userID, body.Password); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// Compile-time check do role enum (silencia unused import quando middleware muda)
var _ = models.RolePatient
