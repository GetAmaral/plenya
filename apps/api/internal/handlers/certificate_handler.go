package handlers

import (
	"io"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/middleware"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/services"
)

type CertificateHandler struct {
	db          *gorm.DB
	certService *services.CertificateService
}

func NewCertificateHandler(db *gorm.DB, certService *services.CertificateService) *CertificateHandler {
	return &CertificateHandler{
		db:          db,
		certService: certService,
	}
}

// UploadCertificate godoc
// @Summary Upload digital certificate A1 for doctor
// @Description Admin uploads certificate A1 (.pfx) for a doctor with password
// @Tags Certificates
// @Accept multipart/form-data
// @Produce json
// @Param doctorId formData string true "Doctor UUID"
// @Param password formData string true "Certificate password"
// @Param certificate formData file true "Certificate file (.pfx or .p12)"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "success and message"
// @Failure 400 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse "Forbidden - admin only"
// @Failure 500 {object} dto.ErrorResponse
// @Router /admin/certificates/upload [post]
func (h *CertificateHandler) UploadCertificate(c *fiber.Ctx) error {
	// Verificar se é admin
	roles := middleware.GetUserRoles(c)
	isAdmin := false
	for _, role := range roles {
		if role == string(models.RoleAdmin) {
			isAdmin = true
			break
		}
	}

	if !isAdmin {
		return c.Status(403).JSON(dto.ErrorResponse{Error: "forbidden - admin only"})
	}

	// Parse multipart form
	doctorIDStr := c.FormValue("doctorId")
	password := c.FormValue("password")

	if doctorIDStr == "" || password == "" {
		return c.Status(400).JSON(dto.ErrorResponse{Error: "doctorId and password are required"})
	}

	doctorID, err := uuid.Parse(doctorIDStr)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{Error: "invalid doctorId format"})
	}

	// Get certificate file
	file, err := c.FormFile("certificate")
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{Error: "missing certificate file"})
	}

	// Verificar extensão
	if file.Header.Get("Content-Type") != "application/x-pkcs12" &&
		file.Header.Get("Content-Type") != "application/octet-stream" {
		// Aceitar se for .pfx ou .p12
		filename := file.Filename
		if len(filename) < 4 {
			return c.Status(400).JSON(dto.ErrorResponse{Error: "invalid certificate file"})
		}
		ext := filename[len(filename)-4:]
		if ext != ".pfx" && ext != ".p12" {
			return c.Status(400).JSON(dto.ErrorResponse{Error: "certificate must be .pfx or .p12 file"})
		}
	}

	// Ler bytes do arquivo
	fileContent, err := file.Open()
	if err != nil {
		return c.Status(500).JSON(dto.ErrorResponse{Error: "failed to open file"})
	}
	defer fileContent.Close()

	pfxBytes, err := io.ReadAll(fileContent)
	if err != nil {
		return c.Status(500).JSON(dto.ErrorResponse{Error: "failed to read file"})
	}

	// Upload via service
	validationResult, err := h.certService.UploadCertificate(doctorID, pfxBytes, password)
	if err != nil {
		// Se for erro de CPF divergente, retornar warning com dados para confirmação
		if err == services.ErrCPFMismatch && validationResult != nil {
			return c.Status(409).JSON(fiber.Map{
				"success":             false,
				"error":               "CPF_MISMATCH",
				"message":             "O CPF do certificado não corresponde ao CPF do usuário",
				"requiresConfirmation": true,
				"userCPF":             validationResult.UserCPF,
				"certificateCPF":      validationResult.CertificateCPF,
			})
		}

		// Outros erros
		return c.Status(400).JSON(dto.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Certificado enviado com sucesso",
	})
}

// ListCertificates godoc
// @Summary List all certificates
// @Description Admin sees all certificates, doctors see only their own
// @Tags Certificates
// @Produce json
// @Security BearerAuth
// @Success 200 {array} map[string]interface{} "List of certificates"
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /certificates [get]
func (h *CertificateHandler) ListCertificates(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)
	roles := middleware.GetUserRoles(c)

	// Check if admin
	isAdmin := false
	for _, role := range roles {
		if role == string(models.RoleAdmin) {
			isAdmin = true
			break
		}
	}

	var users []models.User
	query := h.db.Select("id, name, email, cpf, certificate_expiry, certificate_serial, certificate_cpf, certificate_name, certificate_active").
		Where("certificate_expiry IS NOT NULL")

	// Non-admin can only see their own certificate
	if !isAdmin {
		query = query.Where("id = ?", userID)
	}

	if err := query.Find(&users).Error; err != nil {
		return c.Status(500).JSON(dto.ErrorResponse{Error: "failed to fetch certificates"})
	}

	// Map to response
	certificates := make([]fiber.Map, len(users))
	for i, user := range users {
		daysUntilExpiry := 0
		needsRenewal := false
		isExpired := false

		if user.CertificateExpiry != nil {
			daysUntilExpiry = int(user.CertificateExpiry.Sub(c.Context().Time()).Hours() / 24)
			needsRenewal = daysUntilExpiry <= 30
			isExpired = daysUntilExpiry < 0
		}

		certificates[i] = fiber.Map{
			"userId":            user.ID,
			"userName":          user.Name,
			"userEmail":         user.Email,
			"userCPF":           user.CPF,
			"validUntil":        user.CertificateExpiry,
			"certificateSerial": user.CertificateSerial,
			"certificateCPF":    user.CertificateCPF,
			"certificateName":   user.CertificateName,
			"certificateActive": user.CertificateActive,
			"daysUntilExpiry":   daysUntilExpiry,
			"needsRenewal":      needsRenewal,
			"isExpired":         isExpired,
		}
	}

	return c.JSON(fiber.Map{
		"data": certificates,
	})
}

// GetCertificateStatus godoc
// @Summary Get certificate status for logged-in doctor
// @Description Returns certificate expiry and validity status for the authenticated doctor
// @Tags Certificates
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "hasCertificate, validUntil, daysUntilExpiry, needsRenewal"
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /certificates/status [get]
func (h *CertificateHandler) GetCertificateStatus(c *fiber.Ctx) error {
	doctorID := middleware.GetUserID(c)

	var user models.User
	if err := h.db.Select("certificate_expiry, certificate_serial").
		First(&user, doctorID).Error; err != nil {
		return c.Status(500).JSON(dto.ErrorResponse{Error: "failed to fetch user"})
	}

	// Se não tem certificado
	if user.CertificateExpiry == nil {
		return c.JSON(fiber.Map{
			"hasCertificate": false,
		})
	}

	// Calcular dias até expirar
	daysUntilExpiry := int(user.CertificateExpiry.Sub(c.Context().Time()).Hours() / 24)

	return c.JSON(fiber.Map{
		"hasCertificate":      true,
		"validUntil":          user.CertificateExpiry,
		"certificateSerial":   user.CertificateSerial,
		"daysUntilExpiry":     daysUntilExpiry,
		"needsRenewal":        daysUntilExpiry <= 30,
	})
}

// ToggleCertificateActive godoc
// @Summary Toggle certificate active status
// @Description Ativa ou desativa o certificado de um usuário (admin only)
// @Tags Certificates
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param userId path string true "User ID"
// @Success 200 {object} map[string]interface{} "success message"
// @Failure 400 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /admin/certificates/{userId}/toggle [patch]
func (h *CertificateHandler) ToggleCertificateActive(c *fiber.Ctx) error {
	// Verificar se é admin
	roles := middleware.GetUserRoles(c)
	isAdmin := false
	for _, role := range roles {
		if role == string(models.RoleAdmin) {
			isAdmin = true
			break
		}
	}

	if !isAdmin {
		return c.Status(403).JSON(dto.ErrorResponse{Error: "forbidden - admin only"})
	}

	// Parse userId
	userIDStr := c.Params("userId")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{Error: "invalid userId format"})
	}

	// Buscar usuário
	var user models.User
	if err := h.db.Select("id, certificate_active, certificate_expiry").
		First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(dto.ErrorResponse{Error: "user not found"})
		}
		return c.Status(500).JSON(dto.ErrorResponse{Error: "failed to fetch user"})
	}

	// Verificar se tem certificado
	if user.CertificateExpiry == nil {
		return c.Status(400).JSON(dto.ErrorResponse{Error: "user has no certificate"})
	}

	// Toggle estado
	newStatus := !user.CertificateActive
	if err := h.db.Model(&user).Update("certificate_active", newStatus).Error; err != nil {
		return c.Status(500).JSON(dto.ErrorResponse{Error: "failed to update certificate status"})
	}

	statusText := "desativado"
	if newStatus {
		statusText = "ativado"
	}

	return c.JSON(fiber.Map{
		"success":           true,
		"certificateActive": newStatus,
		"message":           "Certificado " + statusText + " com sucesso",
	})
}

// DeleteCertificate godoc
// @Summary Delete certificate
// @Description Admin can delete any certificate, doctors can delete only their own
// @Tags Certificates
// @Produce json
// @Param userId path string true "User ID"
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "success message"
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse "Forbidden"
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /certificates/{userId} [delete]
func (h *CertificateHandler) DeleteCertificate(c *fiber.Ctx) error {
	userIDParam := c.Params("userId")
	targetUserID, err := uuid.Parse(userIDParam)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{Error: "invalid userId format"})
	}

	currentUserID := middleware.GetUserID(c)
	roles := middleware.GetUserRoles(c)

	// Check if admin
	isAdmin := false
	for _, role := range roles {
		if role == string(models.RoleAdmin) {
			isAdmin = true
			break
		}
	}

	// Non-admin can only delete their own certificate
	if !isAdmin && currentUserID != targetUserID {
		return c.Status(403).JSON(dto.ErrorResponse{Error: "forbidden - can only delete your own certificate"})
	}

	// Find user
	var user models.User
	if err := h.db.First(&user, targetUserID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(dto.ErrorResponse{Error: "user not found"})
		}
		return c.Status(500).JSON(dto.ErrorResponse{Error: "failed to fetch user"})
	}

	// Clear certificate fields
	user.CertificatePFX = nil
	user.CertificatePassword = nil
	user.CertificateExpiry = nil
	user.CertificateSerial = nil

	if err := h.db.Save(&user).Error; err != nil {
		return c.Status(500).JSON(dto.ErrorResponse{Error: "failed to delete certificate"})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Certificado removido com sucesso",
	})
}
