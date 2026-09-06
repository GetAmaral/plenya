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
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
	"github.com/plenya/api/internal/utils"
)

type PatientPortalHandler struct {
	svc          *services.PatientPortalService
	dashboard    *services.PatientDashboardService
	continuum    *services.PatientContinuumService
	appointments *services.PatientAppointmentService
	messages     *services.PatientMessagesService
	labs         *services.PatientLabsService
	scores       *services.PatientScoresService
	profile      *services.PatientProfileService
	documents    *services.PatientDocumentsService
	workouts     *services.PatientWorkoutsService
	checkIns     *services.PatientCheckInsService
	notifPrefs   *services.NotificationPreferencesService
	notification *services.NotificationService
}

func NewPatientPortalHandler(
	svc *services.PatientPortalService,
	dashboard *services.PatientDashboardService,
	continuum *services.PatientContinuumService,
	appointments *services.PatientAppointmentService,
	messages *services.PatientMessagesService,
	labs *services.PatientLabsService,
	scores *services.PatientScoresService,
	profile *services.PatientProfileService,
	documents *services.PatientDocumentsService,
	workouts *services.PatientWorkoutsService,
	checkIns *services.PatientCheckInsService,
	notifPrefs *services.NotificationPreferencesService,
	notification *services.NotificationService,
) *PatientPortalHandler {
	return &PatientPortalHandler{
		svc:          svc,
		dashboard:    dashboard,
		continuum:    continuum,
		appointments: appointments,
		messages:     messages,
		labs:         labs,
		scores:       scores,
		profile:      profile,
		documents:    documents,
		workouts:     workouts,
		checkIns:     checkIns,
		notifPrefs:   notifPrefs,
		notification: notification,
	}
}

// ListDocuments GET /api/v1/patient/me/documents
func (h *PatientPortalHandler) ListDocuments(c *fiber.Ctx) error {
	rows, err := h.documents.List(middleware.GetPatientID(c))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(rows)
}

// DownloadDocument GET /api/v1/patient/me/documents/:id/download
func (h *PatientPortalHandler) DownloadDocument(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	doc, fullPath, err := h.documents.GetForDownload(middleware.GetPatientID(c), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	c.Set("Content-Type", doc.ContentType)
	c.Set("Content-Disposition", `attachment; filename="`+doc.FileName+`"`)
	return c.SendFile(fullPath)
}

// StaffDownloadDocument GET /api/v1/patients/:id/documents/:docId/download
//
// O espelho, para o staff, do download que só o portal do paciente tinha. Sem isto o PDF do plano
// de devolutiva existia e a tela do médico não tinha porta para ele: a paciente baixava a própria
// devolutiva e quem a escreveu, não.
//
// O paciente vem da URL (e não da sessão, como no portal), então `GetForDownload` continua sendo o
// que amarra documento e paciente — pedir o documento de outro paciente é 404, não vazamento.
func (h *PatientPortalHandler) StaffDownloadDocument(c *fiber.Ctx) error {
	pid, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid patient id"})
	}
	docID, err := uuid.Parse(c.Params("docId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid document id"})
	}
	doc, fullPath, err := h.documents.GetForDownload(pid, docID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	c.Set("Content-Type", doc.ContentType)
	// `inline` e não `attachment`: o médico quer OLHAR antes de mandar, e o nome acentuado precisa
	// do filename* em UTF-8 (mesmo padrão da receita e do pedido de exames).
	c.Set("Content-Disposition", utils.ContentDisposition(doc.FileName, "documento.pdf"))
	return c.SendFile(fullPath)
}

// StaffListDocuments GET /api/v1/patients/:id/documents
func (h *PatientPortalHandler) StaffListDocuments(c *fiber.Ctx) error {
	pid, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid patient id"})
	}
	rows, err := h.documents.List(pid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(rows)
}

// StaffUploadDocument POST /api/v1/patients/:id/documents (multipart)
func (h *PatientPortalHandler) StaffUploadDocument(c *fiber.Ctx) error {
	pid, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid patient id"})
	}
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "arquivo obrigatório (campo 'file')"})
	}
	docType := c.FormValue("type", string(models.DocumentTypeOther))
	title := c.FormValue("title")
	description := c.FormValue("description")
	doc, err := h.documents.Create(services.CreateDocumentInput{
		PatientID:   pid,
		UploadedBy:  middleware.GetUserID(c),
		Type:        models.PatientDocumentType(docType),
		Title:       title,
		Description: description,
		File:        file,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(doc)
}

// StaffDeleteDocument DELETE /api/v1/patient-documents/:id
func (h *PatientPortalHandler) StaffDeleteDocument(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := h.documents.Delete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// UpdateProfile PUT /api/v1/patient/me/profile
func (h *PatientPortalHandler) UpdateProfile(c *fiber.Ctx) error {
	var body services.UpdatableProfile
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	if err := h.profile.UpdateProfile(middleware.GetPatientID(c), body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// LGPDExport GET /api/v1/patient/me/lgpd/export
func (h *PatientPortalHandler) LGPDExport(c *fiber.Ctx) error {
	out, err := h.profile.LGPDExport(middleware.GetPatientID(c))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	c.Set("Content-Disposition", `attachment; filename="plenya-meus-dados.json"`)
	return c.JSON(out)
}

type lgpdDeleteBody struct {
	Reason string `json:"reason"`
}

// LGPDRequestDelete POST /api/v1/patient/me/lgpd/account-delete-request
func (h *PatientPortalHandler) LGPDRequestDelete(c *fiber.Ctx) error {
	var body lgpdDeleteBody
	_ = c.BodyParser(&body)
	if err := h.profile.RequestAccountDeletion(services.LGPDDeleteRequestInput{
		PatientID: middleware.GetPatientID(c),
		UserID:    middleware.GetUserID(c),
		Reason:    body.Reason,
	}, h.notification); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// ListScores GET /api/v1/patient/me/scores
func (h *PatientPortalHandler) ListScores(c *fiber.Ctx) error {
	out, err := h.scores.ListAll(middleware.GetPatientID(c))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(out)
}

// GetCompleteSnapshot GET /api/v1/patient/me/score-snapshots/:id
func (h *PatientPortalHandler) GetCompleteSnapshot(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	out, err := h.scores.GetCompleteSnapshot(middleware.GetPatientID(c), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(out)
}

// GetLatestSnapshot GET /api/v1/patient/me/score-snapshots/latest
// Escore COMPLETO mais recente do paciente (com pilares AGIR) para o radar do
// portal. 204 quando o paciente não tem escore completo (só Light, sem pilares).
func (h *PatientPortalHandler) GetLatestSnapshot(c *fiber.Ctx) error {
	out, err := h.scores.GetLatestCompleteSnapshot(middleware.GetPatientID(c))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if out == nil {
		return c.JSON(nil) // 200 + null: paciente sem escore completo
	}
	return c.JSON(out)
}

// ============================================================
// Labs / Prescriptions / Physical Assessments
// ============================================================

func (h *PatientPortalHandler) ListLabBatches(c *fiber.Ctx) error {
	out, err := h.labs.ListBatches(middleware.GetPatientID(c))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(out)
}

func (h *PatientPortalHandler) GetLabBatch(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	out, err := h.labs.GetBatch(middleware.GetPatientID(c), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(out)
}

func (h *PatientPortalHandler) ListLabRequests(c *fiber.Ctx) error {
	out, err := h.labs.ListLabRequests(middleware.GetPatientID(c))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(out)
}

func (h *PatientPortalHandler) ListPrescriptions(c *fiber.Ctx) error {
	out, err := h.labs.ListPrescriptions(middleware.GetPatientID(c))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(out)
}

func (h *PatientPortalHandler) ListBoxes(c *fiber.Ctx) error {
	out, err := h.labs.ListBoxes(middleware.GetPatientID(c))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(out)
}

func (h *PatientPortalHandler) ListPhysicalAssessments(c *fiber.Ctx) error {
	out, err := h.labs.ListPhysicalAssessments(middleware.GetPatientID(c))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(out)
}

func (h *PatientPortalHandler) GetPhysicalAssessmentHTML(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	html, err := h.labs.GetPhysicalAssessmentHTML(middleware.GetPatientID(c), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(html)
}

// ListMessages GET /api/v1/patient/me/messages?before=ISO
func (h *PatientPortalHandler) ListMessages(c *fiber.Ctx) error {
	patientID := middleware.GetPatientID(c)
	var before *time.Time
	if b := c.Query("before"); b != "" {
		if t, err := time.Parse(time.RFC3339, b); err == nil {
			before = &t
		}
	}
	rows, err := h.messages.List(patientID, 100, before)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(rows)
}

type sendMessageBody struct {
	Content string `json:"content"`
}

// SendMessage POST /api/v1/patient/me/messages
func (h *PatientPortalHandler) SendMessage(c *fiber.Ctx) error {
	patientID := middleware.GetPatientID(c)
	userID := middleware.GetUserID(c)
	var body sendMessageBody
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	row, err := h.messages.Send(patientID, userID, body.Content)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(row)
}

// MessagesUnreadCount GET /api/v1/patient/me/messages/unread-count
func (h *PatientPortalHandler) MessagesUnreadCount(c *fiber.Ctx) error {
	patientID := middleware.GetPatientID(c)
	userID := middleware.GetUserID(c)
	count, err := h.messages.UnreadCount(patientID, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"unread": count})
}

// MessagesMarkRead POST /api/v1/patient/me/messages/read
func (h *PatientPortalHandler) MessagesMarkRead(c *fiber.Ctx) error {
	patientID := middleware.GetPatientID(c)
	userID := middleware.GetUserID(c)
	if err := h.messages.MarkRead(patientID, userID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// ListAppointments GET /api/v1/patient/me/appointments?range=upcoming|past
func (h *PatientPortalHandler) ListAppointments(c *fiber.Ctx) error {
	patientID := middleware.GetPatientID(c)
	rangeKind := c.Query("range", "upcoming")
	rows, err := h.appointments.List(patientID, rangeKind)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(rows)
}

// GetAppointment GET /api/v1/patient/me/appointments/:id
func (h *PatientPortalHandler) GetAppointment(c *fiber.Ctx) error {
	patientID := middleware.GetPatientID(c)
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	row, err := h.appointments.GetByID(patientID, id)
	if err != nil {
		if errors.Is(err, services.ErrAppointmentNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "consulta não encontrada"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(row)
}

// ConfirmAppointment POST /api/v1/patient/me/appointments/:id/confirm
func (h *PatientPortalHandler) ConfirmAppointment(c *fiber.Ctx) error {
	patientID := middleware.GetPatientID(c)
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := h.appointments.Confirm(patientID, id); err != nil {
		if errors.Is(err, services.ErrAppointmentNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "consulta não encontrada"})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

type requestActionBody struct {
	Reason string `json:"reason"`
}

// RequestCancelAppointment POST /api/v1/patient/me/appointments/:id/request-cancel
func (h *PatientPortalHandler) RequestCancelAppointment(c *fiber.Ctx) error {
	return h.requestAction(c, "cancel")
}

// RequestRescheduleAppointment POST /api/v1/patient/me/appointments/:id/request-reschedule
func (h *PatientPortalHandler) RequestRescheduleAppointment(c *fiber.Ctx) error {
	return h.requestAction(c, "reschedule")
}

func (h *PatientPortalHandler) requestAction(c *fiber.Ctx, action string) error {
	patientID := middleware.GetPatientID(c)
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var body requestActionBody
	_ = c.BodyParser(&body)
	userID := middleware.GetUserID(c)
	if err := h.appointments.RequestAction(services.RequestActionInput{
		PatientID:     patientID,
		AppointmentID: id,
		Action:        action,
		Reason:        body.Reason,
		UserID:        userID,
	}); err != nil {
		if errors.Is(err, services.ErrAppointmentNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "consulta não encontrada"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
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
	// M5 — passa IP pra notificação por email. c.IP() já respeita TrustedProxies.
	if err := h.svc.SetPasswordWithIP(userID, body.Password, c.IP()); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// Compile-time check do role enum (silencia unused import quando middleware muda)
var _ = models.RolePatient
