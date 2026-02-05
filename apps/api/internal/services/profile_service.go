package services

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

type ProfileService struct {
	db *gorm.DB
}

func NewProfileService(db *gorm.DB) *ProfileService {
	return &ProfileService{db: db}
}

// GetProfile retorna o perfil do usuário
func (s *ProfileService) GetProfile(userID uuid.UUID) (*dto.UserResponse, error) {
	var user models.User
	if err := s.db.First(&user, userID).Error; err != nil {
		return nil, err
	}

	return s.toDTO(&user), nil
}

// UpdateProfile atualiza o perfil do usuário (dados não sensíveis)
func (s *ProfileService) UpdateProfile(userID uuid.UUID, req *dto.UpdateProfileRequest) (*dto.UserResponse, error) {
	var user models.User
	if err := s.db.First(&user, userID).Error; err != nil {
		return nil, err
	}

	// Atualizar apenas campos permitidos
	updates := make(map[string]interface{})

	if req.Name != nil {
		updates["name"] = *req.Name
	}

	if req.CPF != nil {
		if *req.CPF == "" {
			updates["cpf"] = nil
		} else {
			// Sanitizar CPF (remover pontos, traços, etc)
			newCPF := sanitizeCPF(*req.CPF)
			updates["cpf"] = newCPF

			// Se CPF mudou e usuário tem certificado, verificar divergência
			if user.CertificateCPF != nil {
				oldCPF := ""
				if user.CPF != nil {
					oldCPF = *user.CPF
				}

				// Se CPF mudou
				if oldCPF != newCPF {
					// Normalizar CPFs para comparação
					normalizedNewCPF := sanitizeCPF(newCPF)
					normalizedCertCPF := sanitizeCPF(*user.CertificateCPF)

					// Se CPF do usuário for diferente do certificado, desativar
					if normalizedNewCPF != normalizedCertCPF {
						updates["certificate_active"] = false
					}
				}
			}
		}
	}

	if req.ProfessionalPhone != nil {
		if *req.ProfessionalPhone == "" {
			updates["professional_phone"] = nil
		} else {
			updates["professional_phone"] = *req.ProfessionalPhone
		}
	}

	if req.ProfessionalAddress != nil {
		if *req.ProfessionalAddress == "" {
			updates["professional_address"] = nil
		} else {
			updates["professional_address"] = *req.ProfessionalAddress
		}
	}

	if req.Specialty != nil {
		if *req.Specialty == "" {
			updates["specialty"] = nil
		} else {
			updates["specialty"] = *req.Specialty
		}
	}

	if req.CRM != nil {
		if *req.CRM == "" {
			updates["crm"] = nil
		} else {
			updates["crm"] = *req.CRM
		}
	}

	if req.CRMUF != nil {
		if *req.CRMUF == "" {
			updates["crmuf"] = nil
		} else {
			updates["crmuf"] = *req.CRMUF
		}
	}

	if req.RQE != nil {
		if *req.RQE == "" {
			updates["rqe"] = nil
		} else {
			updates["rqe"] = *req.RQE
		}
	}

	if len(updates) > 0 {
		if err := s.db.Model(&user).Updates(updates).Error; err != nil {
			return nil, err
		}

		// Recarregar usuário atualizado
		if err := s.db.First(&user, userID).Error; err != nil {
			return nil, err
		}
	}

	return s.toDTO(&user), nil
}

// toDTO converte User para UserResponse
func (s *ProfileService) toDTO(user *models.User) *dto.UserResponse {
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
