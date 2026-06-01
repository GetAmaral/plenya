package services

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/crypto"
	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

// digitsOnly extrai apenas os dígitos de uma string (para busca por
// telefone/CPF, que podem vir formatados).
func digitsOnly(s string) string {
	var b strings.Builder
	for _, r := range s {
		if r >= '0' && r <= '9' {
			b.WriteByte(byte(r))
		}
	}
	return b.String()
}

var (
	ErrPatientNotFound      = errors.New("patient not found")
	ErrPatientAlreadyExists = errors.New("patient already exists for this user")
	ErrPatientCPFExists     = errors.New("patient with this CPF already exists")
	ErrUnauthorized         = errors.New("unauthorized")
)

type PatientService struct {
	db *gorm.DB
}

func NewPatientService(db *gorm.DB) *PatientService {
	return &PatientService{db: db}
}

// Create cria um novo paciente
func (s *PatientService) Create(userID uuid.UUID, req *dto.CreatePatientRequest) (*dto.PatientResponse, error) {
	// Verificar se já existe paciente para este usuário
	var existing models.Patient
	if err := s.db.Where("user_id = ?", userID).First(&existing).Error; err == nil {
		return nil, ErrPatientAlreadyExists
	}

	// Dedupe por CPF: um CPF = um paciente. Sem isso, a violação de unique no
	// banco vira 500 cru no balcão; aqui devolvemos erro tratável (409 amigável).
	if req.CPF != nil && *req.CPF != "" {
		if found, _ := s.FindByCPF(*req.CPF); found != nil {
			return nil, ErrPatientCPFExists
		}
	}

	// Parse da data de nascimento (opcional no cadastro rápido). Ausente => zero,
	// e CalculateAge trata a data zerada (Age=0, AgeText="").
	var birthDate time.Time
	if req.BirthDate != "" {
		parsed, err := time.Parse("2006-01-02", req.BirthDate)
		if err != nil {
			return nil, errors.New("invalid birth date format, expected YYYY-MM-DD")
		}
		birthDate = parsed
	}

	// Gênero default 'other' quando omitido (cadastro rápido).
	gender := req.Gender
	if gender == "" {
		gender = models.GenderOther
	}

	// Criar paciente
	patient := models.Patient{
		UserID:           userID,
		Name:             req.Name,
		CPF:              req.CPF,              // Será criptografado pelo hook
		RG:               req.RG,               // Será criptografado pelo hook
		BirthDate:        birthDate,
		Gender:           gender,
		SocialGender:     req.SocialGender,
		Email:            req.Email,
		Phone:            req.Phone,
		Address:          req.Address,
		Municipality:     req.Municipality,
		State:            req.State,
		MotherName:       req.MotherName,
		FatherName:       req.FatherName,
		Height:           req.Height,
		Weight:           req.Weight,
		BloodType:        req.BloodType,
		MaritalStatus:    req.MaritalStatus,
		Occupation:       req.Occupation,
		EmergencyContact: req.EmergencyContact,
		EmergencyPhone:   req.EmergencyPhone,
	}

	if err := s.db.Create(&patient).Error; err != nil {
		return nil, err
	}

	return s.toDTO(&patient), nil
}

// GetByID busca um paciente por ID.
//
// CRITICAL C2 — RBAC clínica única:
//   - Role 'patient' nunca deveria chegar aqui (middleware RequireAnyStaff
//     bloqueia em /patients/*). Mantemos o filtro por user_id como defense-in-depth
//     pra cobrir callers internos que não passam pelo middleware (ex: future scripts).
//   - Todos staff (admin/manager/secretary/doctor/nurse/psychologist/nutritionist/
//     physical_educator) podem ver qualquer paciente da clínica — decisão de
//     produto: clínica única por instalação. Multi-tenant entre clínicas é V2.
func (s *PatientService) GetByID(patientID, userID uuid.UUID, userRole models.Role) (*dto.PatientResponse, error) {
	var patient models.Patient
	query := s.db.Where("id = ?", patientID)

	// Defense-in-depth: caso passe um patient role direto.
	if userRole == models.RolePatient {
		query = query.Where("user_id = ?", userID)
	}

	if err := query.First(&patient).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPatientNotFound
		}
		return nil, err
	}

	return s.toDTO(&patient), nil
}

// FindByCPF — M4 — lookup performante via CPFBlindIndex (HMAC).
// Retorna nil sem erro se não achar (caller distingue notfound de erro).
//
// Uso: deduplicação na criação (um CPF = um Patient), conversão de Lead via
// CPF informado em formulário, etc.
func (s *PatientService) FindByCPF(cpfPlain string) (*models.Patient, error) {
	idx, err := crypto.BlindIndexCPF(cpfPlain)
	if err != nil || idx == "" {
		return nil, err
	}
	var p models.Patient
	if err := s.db.Where("cpf_blind_index = ?", idx).First(&p).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

// List lista todos os pacientes (com paginação)
func (s *PatientService) List(userID uuid.UUID, userRole models.Role, limit, offset int, search string) ([]dto.PatientResponse, error) {
	var patients []models.Patient
	query := s.db.Limit(limit).Offset(offset).Order("created_at DESC")

	// Pacientes só podem ver seus próprios dados
	if userRole == models.RolePatient {
		query = query.Where("user_id = ?", userID)
	}

	// Busca global (recepção / Cmd+K). Nome e telefone são texto puro → ILIKE.
	// CPF é criptografado (não dá ILIKE): se o termo tiver 11 dígitos, casa exato
	// via blind index. 11 dígitos também pode ser celular (DDD+9), então nesse
	// caso buscamos por nome OU telefone OU CPF ao mesmo tempo.
	if search = strings.TrimSpace(search); search != "" {
		digits := digitsOnly(search)
		namePat := "%" + search + "%"
		// unaccent() nos dois lados do nome: "Joao" casa "João" (extensão unaccent).
		switch {
		case len(digits) == 11:
			cpfIdx, _ := crypto.BlindIndexCPF(digits)
			query = query.Where(
				"unaccent(name) ILIKE unaccent(?) OR phone ILIKE ? OR cpf_blind_index = ?",
				namePat, "%"+digits+"%", cpfIdx,
			)
		case len(digits) >= 3:
			query = query.Where("unaccent(name) ILIKE unaccent(?) OR phone ILIKE ?", namePat, "%"+digits+"%")
		default:
			query = query.Where("unaccent(name) ILIKE unaccent(?)", namePat)
		}
	}

	if err := query.Find(&patients).Error; err != nil {
		return nil, err
	}

	result := make([]dto.PatientResponse, len(patients))
	for i, p := range patients {
		result[i] = *s.toDTO(&p)
	}

	return result, nil
}

// Update atualiza um paciente
func (s *PatientService) Update(patientID, userID uuid.UUID, userRole models.Role, req *dto.UpdatePatientRequest) (*dto.PatientResponse, error) {
	var patient models.Patient
	query := s.db.Where("id = ?", patientID)

	// Pacientes só podem atualizar seus próprios dados
	if userRole == models.RolePatient {
		query = query.Where("user_id = ?", userID)
	}

	if err := query.First(&patient).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPatientNotFound
		}
		return nil, err
	}

	// Atualizar campos
	if req.Name != nil {
		patient.Name = *req.Name
	}
	if req.CPF != nil {
		patient.CPF = req.CPF
	}
	if req.RG != nil {
		patient.RG = req.RG
	}
	if req.BirthDate != nil {
		birthDate, err := time.Parse("2006-01-02", *req.BirthDate)
		if err != nil {
			return nil, errors.New("invalid birth date format, expected YYYY-MM-DD")
		}
		patient.BirthDate = birthDate
	}
	if req.Gender != nil {
		patient.Gender = *req.Gender
	}
	if req.SocialGender != nil {
		patient.SocialGender = req.SocialGender
	}
	if req.Email != nil {
		patient.Email = req.Email
	}
	if req.Phone != nil {
		patient.Phone = req.Phone
	}
	if req.Address != nil {
		patient.Address = req.Address
	}
	if req.Municipality != nil {
		patient.Municipality = req.Municipality
	}
	if req.State != nil {
		patient.State = req.State
	}
	if req.MotherName != nil {
		patient.MotherName = req.MotherName
	}
	if req.FatherName != nil {
		patient.FatherName = req.FatherName
	}
	if req.Height != nil {
		patient.Height = req.Height
	}
	if req.Weight != nil {
		patient.Weight = req.Weight
	}
	if req.BloodType != nil {
		patient.BloodType = req.BloodType
	}
	if req.MaritalStatus != nil {
		patient.MaritalStatus = req.MaritalStatus
	}
	if req.Occupation != nil {
		patient.Occupation = req.Occupation
	}
	if req.EmergencyContact != nil {
		patient.EmergencyContact = req.EmergencyContact
	}
	if req.EmergencyPhone != nil {
		patient.EmergencyPhone = req.EmergencyPhone
	}

	if err := s.db.Save(&patient).Error; err != nil {
		return nil, err
	}

	return s.toDTO(&patient), nil
}

// Delete faz soft delete de um paciente
func (s *PatientService) Delete(patientID, userID uuid.UUID, userRole models.Role) error {
	query := s.db.Where("id = ?", patientID)

	// Pacientes não podem deletar seus próprios dados
	if userRole == models.RolePatient {
		return ErrUnauthorized
	}

	result := query.Delete(&models.Patient{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrPatientNotFound
	}

	return nil
}

// toDTO converte Patient para PatientResponse
func (s *PatientService) toDTO(patient *models.Patient) *dto.PatientResponse {
	// Cadastro rápido sem data de nascimento grava a data zero (0001-01-01).
	// Devolvemos string vazia em vez de "0001-01-01" para o frontend.
	birthDate := ""
	if !patient.BirthDate.IsZero() {
		birthDate = patient.BirthDate.Format("2006-01-02")
	}

	return &dto.PatientResponse{
		ID:               patient.ID.String(),
		UserID:           patient.UserID.String(),
		Name:             patient.Name,
		CPF:              patient.CPF,              // Já foi descriptografado pelo AfterFind hook
		RG:               patient.RG,               // Já foi descriptografado pelo AfterFind hook
		BirthDate:        birthDate,
		Gender:           patient.Gender,
		SocialGender:     patient.SocialGender,
		Age:              patient.Age,
		AgeText:          patient.AgeText,
		Email:            patient.Email,
		Phone:            patient.Phone,
		Address:          patient.Address,
		Municipality:     patient.Municipality,
		State:            patient.State,
		MotherName:       patient.MotherName,
		FatherName:       patient.FatherName,
		Height:           patient.Height,
		Weight:           patient.Weight,
		BloodType:        patient.BloodType,
		MaritalStatus:    patient.MaritalStatus,
		Occupation:       patient.Occupation,
		EmergencyContact: patient.EmergencyContact,
		EmergencyPhone:   patient.EmergencyPhone,
		CreatedAt:        patient.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        patient.UpdatedAt.Format(time.RFC3339),
	}
}
