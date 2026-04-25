package services

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/config"
	"github.com/plenya/api/internal/models"
)

var (
	ErrLobbyTokenNotFound = errors.New("lobby token not found")
	ErrLobbyTokenExpired  = errors.New("lobby token expired")
)

// TelemedLobbyService gera/valida tokens de entrada pública na sala Daily.co.
type TelemedLobbyService struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewTelemedLobbyService(db *gorm.DB, cfg *config.Config) *TelemedLobbyService {
	return &TelemedLobbyService{db: db, cfg: cfg}
}

// PublicLinkURL retorna a URL pública /sala/[token] absoluta.
// Usado por email + WhatsApp templates.
func (s *TelemedLobbyService) PublicLinkURL(token string) string {
	base := s.cfg.PatientPortal.PublicURL
	if base == "" {
		base = "https://meu.plenyasaude.com.br"
	}
	return fmt.Sprintf("%s/sala/%s", strings.TrimRight(base, "/"), token)
}

// EnsureTokenForAppointmentID carrega o appointment e ensura token.
// Wrapper conveniente pra callers que só têm o ID.
func (s *TelemedLobbyService) EnsureTokenForAppointmentID(appointmentID uuid.UUID) (*models.TelemedLobbyToken, error) {
	var appt models.Appointment
	if err := s.db.First(&appt, "id = ?", appointmentID).Error; err != nil {
		return nil, err
	}
	return s.EnsureTokenForAppointment(&appt)
}

// EnsureTokenForAppointment cria (ou reaproveita) um token vinculado a appointment.
// Idempotente: se já existe token válido, retorna ele.
func (s *TelemedLobbyService) EnsureTokenForAppointment(appt *models.Appointment) (*models.TelemedLobbyToken, error) {
	if appt.Type != models.AppointmentTelemedicine {
		return nil, errors.New("appointment não é telemedicina")
	}

	// Reusa se já tem válido
	var existing models.TelemedLobbyToken
	err := s.db.Where("appointment_id = ?", appt.ID).
		Order("created_at DESC").
		First(&existing).Error
	if err == nil && existing.IsValid() {
		return &existing, nil
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	token, err := generateRandomToken(32)
	if err != nil {
		return nil, err
	}

	t := models.TelemedLobbyToken{
		AppointmentID: appt.ID,
		Token:         token,
		// Expira 4h após scheduledAt — cobre atraso e a sessão completa
		ExpiresAt: appt.ScheduledAt.Add(4 * time.Hour),
	}
	if err := s.db.Create(&t).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

// LobbyView é o payload retornado pra página standalone /sala/[token].
// Inclui apenas o necessário pra renderizar a entrada — sem PHI.
type LobbyView struct {
	AppointmentID   uuid.UUID `json:"appointmentId"`
	PatientFirstName string   `json:"patientFirstName"` // só primeiro nome (opaco)
	DoctorName      string    `json:"doctorName"`
	ScheduledAt     time.Time `json:"scheduledAt"`
	DurationMinutes int       `json:"durationMinutes"`
	OpensAt         time.Time `json:"opensAt"`  // -30min do scheduledAt
	ClosesAt        time.Time `json:"closesAt"` // +2h do scheduledAt
	DailyRoomURL    *string   `json:"dailyRoomUrl,omitempty"` // só dentro da janela
	IsOpen          bool      `json:"isOpen"`
}

// Resolve valida o token e retorna LobbyView.
func (s *TelemedLobbyService) Resolve(token string) (*LobbyView, error) {
	if strings.TrimSpace(token) == "" {
		return nil, ErrLobbyTokenNotFound
	}
	var t models.TelemedLobbyToken
	if err := s.db.Where("token = ?", token).First(&t).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrLobbyTokenNotFound
		}
		return nil, err
	}
	if !t.IsValid() {
		return nil, ErrLobbyTokenExpired
	}

	var appt models.Appointment
	if err := s.db.Preload("Patient").First(&appt, "id = ?", t.AppointmentID).Error; err != nil {
		return nil, fmt.Errorf("appointment not found: %w", err)
	}

	doctorName := ""
	var doctor models.User
	if err := s.db.Select("name").First(&doctor, "id = ?", appt.DoctorID).Error; err == nil {
		doctorName = doctor.Name
	}

	opensAt := appt.ScheduledAt.Add(-30 * time.Minute)
	closesAt := appt.ScheduledAt.Add(2 * time.Hour)
	now := time.Now().UTC()
	isOpen := now.After(opensAt) && now.Before(closesAt)

	view := &LobbyView{
		AppointmentID:    appt.ID,
		PatientFirstName: firstName(appt.Patient.Name),
		DoctorName:       doctorName,
		ScheduledAt:      appt.ScheduledAt,
		DurationMinutes:  appt.DurationMinutes,
		OpensAt:          opensAt,
		ClosesAt:         closesAt,
		IsOpen:           isOpen,
	}
	if isOpen && appt.DailyRoomURL != nil {
		view.DailyRoomURL = appt.DailyRoomURL
	}
	return view, nil
}

// MarkUsed marca primeiro acesso (best-effort, não falha o fluxo).
func (s *TelemedLobbyService) MarkUsed(token string) error {
	now := time.Now().UTC()
	return s.db.Model(&models.TelemedLobbyToken{}).
		Where("token = ? AND used_at IS NULL", token).
		Update("used_at", now).Error
}

func firstName(full string) string {
	full = strings.TrimSpace(full)
	if full == "" {
		return "paciente"
	}
	parts := strings.Fields(full)
	return parts[0]
}
