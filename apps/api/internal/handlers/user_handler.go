package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/services"
)

type UserHandler struct {
	userService *services.UserService
	validator   *validator.Validate
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
		validator:   validator.New(),
	}
}

// GetUsers lista todos os usuários (admin only)
// @Summary List users
// @Description Get all users with optional filters
// @Tags Users
// @Accept json
// @Produce json
// @Param role query string false "Filter by role"
// @Param limit query int false "Limit" default(20)
// @Param offset query int false "Offset" default(0)
// @Success 200 {object} dto.UserListResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Router /api/v1/users [get]
func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	// Parse query params
	limit, _ := strconv.Atoi(c.Query("limit", "20"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))
	role := c.Query("role")

	var rolePtr *string
	if role != "" {
		rolePtr = &role
	}

	// Get users
	result, err := h.userService.List(rolePtr, limit, offset)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: "Failed to get users: " + err.Error(),
		})
	}

	return c.JSON(result)
}

// GetUser busca um usuário por ID (admin only)
// @Summary Get user by ID
// @Description Get a single user by ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid user ID",
		})
	}

	user, err := h.userService.GetByID(userID)
	if err != nil {
		if err == services.ErrUserNotFound {
			return c.Status(http.StatusNotFound).JSON(dto.ErrorResponse{
				Error: "User not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: "Failed to get user: " + err.Error(),
		})
	}

	return c.JSON(user)
}

// CreateUser cria um novo usuário (admin only)
// @Summary Create user
// @Description Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body dto.CreateUserRequest true "User data"
// @Success 201 {object} dto.UserResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 409 {object} dto.ErrorResponse
// @Router /api/v1/users [post]
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req dto.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid request body: " + err.Error(),
		})
	}

	// Validate request
	if err := h.validator.Struct(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Validation error: " + err.Error(),
		})
	}

	user, err := h.userService.Create(&req)
	if err != nil {
		if err == services.ErrEmailAlreadyExists {
			return c.Status(http.StatusConflict).JSON(dto.ErrorResponse{
				Error: "Email already exists",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: "Failed to create user: " + err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(user)
}

// UpdateUser atualiza um usuário (admin only)
// @Summary Update user
// @Description Update an existing user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body dto.UpdateUserRequest true "User data"
// @Success 200 {object} dto.UserResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 409 {object} dto.ErrorResponse
// @Router /api/v1/users/{id} [put]
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid user ID",
		})
	}

	var req dto.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid request body: " + err.Error(),
		})
	}

	// Validate request
	if err := h.validator.Struct(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Validation error: " + err.Error(),
		})
	}

	user, err := h.userService.Update(userID, &req)
	if err != nil {
		if err == services.ErrUserNotFound {
			return c.Status(http.StatusNotFound).JSON(dto.ErrorResponse{
				Error: "User not found",
			})
		}
		if err == services.ErrEmailAlreadyExists {
			return c.Status(http.StatusConflict).JSON(dto.ErrorResponse{
				Error: "Email already exists",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: "Failed to update user: " + err.Error(),
		})
	}

	return c.JSON(user)
}

// DeleteUser deleta um usuário (admin only)
// @Summary Delete user
// @Description Soft delete a user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 204
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 403 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /api/v1/users/{id} [delete]
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Invalid user ID",
		})
	}

	// Don't allow deleting yourself
	currentUserID := c.Locals("userID").(uuid.UUID)
	if userID == currentUserID {
		return c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{
			Error: "Cannot delete yourself",
		})
	}

	if err := h.userService.Delete(userID); err != nil {
		if err == services.ErrUserNotFound {
			return c.Status(http.StatusNotFound).JSON(dto.ErrorResponse{
				Error: "User not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
			Error: "Failed to delete user: " + err.Error(),
		})
	}

	return c.SendStatus(http.StatusNoContent)
}

// RegisterRoutes registra as rotas de usuários (admin only)
func (h *UserHandler) RegisterRoutes(router fiber.Router, requireAdmin func(*fiber.Ctx) error) {
	users := router.Group("/users", requireAdmin)
	users.Get("/", h.GetUsers)
	users.Get("/:id", h.GetUser)
	users.Post("/", h.CreateUser)
	users.Put("/:id", h.UpdateUser)
	users.Delete("/:id", h.DeleteUser)
}
