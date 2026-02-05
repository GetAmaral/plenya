package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrInvalidPassword   = errors.New("invalid password")
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

// List lista usuários com filtros e paginação
func (s *UserService) List(role *string, limit, offset int) (*dto.UserListResponse, error) {
	var users []models.User
	var total int64

	query := s.db.Model(&models.User{})

	// Filter by role if provided (usando JSONB contains)
	if role != nil && *role != "" {
		query = query.Where("roles @> ?", `["`+*role+`"]`)
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	// Get paginated results
	if err := query.Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&users).Error; err != nil {
		return nil, err
	}

	// Convert to DTOs
	data := make([]dto.UserResponse, len(users))
	for i, user := range users {
		data[i] = *s.toDTO(&user)
	}

	return &dto.UserListResponse{
		Data:  data,
		Total: total,
		Page:  (offset / limit) + 1,
		Limit: limit,
	}, nil
}

// GetByID busca um usuário por ID
func (s *UserService) GetByID(userID uuid.UUID) (*dto.UserResponse, error) {
	var user models.User
	if err := s.db.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return s.toDTO(&user), nil
}

// Create cria um novo usuário
func (s *UserService) Create(req *dto.CreateUserRequest) (*dto.UserResponse, error) {
	// Check if email already exists
	var count int64
	if err := s.db.Model(&models.User{}).Where("email = ?", req.Email).Count(&count).Error; err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, ErrEmailAlreadyExists
	}

	// Hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Sanitize CPF (remove non-numeric characters)
	var sanitizedCPF *string
	if req.CPF != nil && *req.CPF != "" {
		cleaned := sanitizeCPF(*req.CPF)
		sanitizedCPF = &cleaned
	}

	// Create user
	passwordStr := string(passwordHash)
	user := models.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: &passwordStr,
		CPF:          sanitizedCPF,
	}

	// Set roles
	if err := user.SetRoles(req.Roles); err != nil {
		return nil, err
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return s.toDTO(&user), nil
}

// Update atualiza um usuário
func (s *UserService) Update(userID uuid.UUID, req *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	var user models.User
	if err := s.db.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	// Check email uniqueness if changing
	if req.Email != nil && *req.Email != user.Email {
		var count int64
		if err := s.db.Model(&models.User{}).
			Where("email = ? AND id != ?", *req.Email, userID).
			Count(&count).Error; err != nil {
			return nil, err
		}
		if count > 0 {
			return nil, ErrEmailAlreadyExists
		}
		user.Email = *req.Email
	}

	// Update fields
	if req.Name != nil {
		user.Name = *req.Name
	}

	if req.CPF != nil {
		oldCPF := ""
		if user.CPF != nil {
			oldCPF = *user.CPF
		}

		if *req.CPF == "" {
			user.CPF = nil
		} else {
			cleaned := sanitizeCPF(*req.CPF)
			user.CPF = &cleaned
		}

		// Se CPF mudou e usuário tem certificado, verificar divergência
		newCPF := ""
		if user.CPF != nil {
			newCPF = *user.CPF
		}

		if oldCPF != newCPF && user.CertificateCPF != nil {
			// Normalizar CPFs para comparação
			normalizedNewCPF := sanitizeCPF(newCPF)
			normalizedCertCPF := sanitizeCPF(*user.CertificateCPF)

			// Se CPF do usuário for diferente do certificado, desativar
			if normalizedNewCPF != normalizedCertCPF {
				user.CertificateActive = false
			}
		}
	}

	if req.Password != nil {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		passwordStr := string(passwordHash)
		user.PasswordHash = &passwordStr
	}

	// Update roles if provided
	if req.Roles != nil && len(req.Roles) > 0 {
		if err := user.SetRoles(req.Roles); err != nil {
			return nil, err
		}
	}

	if err := s.db.Save(&user).Error; err != nil {
		return nil, err
	}

	return s.toDTO(&user), nil
}

// Delete faz soft delete de um usuário
func (s *UserService) Delete(userID uuid.UUID) error {
	result := s.db.Delete(&models.User{}, userID)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrUserNotFound
	}

	return nil
}

// toDTO converte User para UserResponse
func (s *UserService) toDTO(user *models.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:                  user.ID.String(),
		Name:                user.Name,
		Email:               user.Email,
		CPF:                 user.CPF,
		Roles:               user.GetRoles(),
		TwoFactorEnabled:    user.TwoFactorEnabled,
		ProfessionalPhone:   user.ProfessionalPhone,
		ProfessionalAddress: user.ProfessionalAddress,
		Specialty:           user.Specialty,
		CRM:                 user.CRM,
		CRMUF:               user.CRMUF,
		RQE:                 user.RQE,
		CreatedAt:           user.CreatedAt.Format(time.RFC3339),
		UpdatedAt:           user.UpdatedAt.Format(time.RFC3339),
	}
}
