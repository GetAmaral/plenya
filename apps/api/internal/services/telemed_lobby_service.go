package services

import (
	"context"
	"errors"
	"fmt"
	"log"
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

// Janela de acesso à teleconsulta — tempo antes do horário programado e depois
// do fim programado em que a sala pode ser aberta. Vale tanto pro staff
// (GetTelemedJoinURL) quanto pro paciente (lobby Resolve), e baliza também a
// expiração do token do lobby e da sala Daily — todos derivam daqui pra não
// haver descompasso (janela "aberta" com sala/token já expirados).
const (
	TelemedWindowBefore = 3 * time.Hour
	TelemedWindowAfter  = 3 * time.Hour
)

// TelemedLobbyService gera/valida tokens de entrada pública na sala Daily.co.
//
// HIGH H9 — quando dailyCoSvc é injetado, Resolve troca a URL crua da sala por
// uma URL com meeting_token (?t=...) escopado ao paciente (is_owner=false,
// enable_screenshare=false, exp=closesAt). Isso permite que a sala seja
// privacy=private no Daily.co — URL sem token retorna 401.
type TelemedLobbyService struct {
	db         *gorm.DB
	cfg        *config.Config
	dailyCoSvc *DailyCoService
}

func NewTelemedLobbyService(db *gorm.DB, cfg *config.Config) *TelemedLobbyService {
	return &TelemedLobbyService{db: db, cfg: cfg}
}

// WithDailyCo injeta o DailyCoService — chamado em main.go quando ambos
// services existem. Quando nil, Resolve degrada e não emite o token (paciente
// vê "sala temporariamente indisponível").
func (s *TelemedLobbyService) WithDailyCo(daily *DailyCoService) *TelemedLobbyService {
	s.dailyCoSvc = daily
	return s
}

// PublicLinkURL retorna a URL pública /sala/[token] absoluta.
// Usado por email + WhatsApp templates.
func (s *TelemedLobbyService) PublicLinkURL(token string) string {
	base := s.cfg.PatientPortal.PublicURL
	if base == "" {
		base = "https://minha.plenyasaude.com.br"
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

	// Janela de acesso: abre TelemedWindowBefore antes e fecha TelemedWindowAfter
	// após o fim programado. O token do lobby expira no fechamento da janela
	// (mesmo instante de closesAt no Resolve) — assim o link público não vale
	// além da janela.
	durationMin := appt.DurationMinutes
	if durationMin <= 0 {
		durationMin = 60 // sane default
	}
	expiresAt := appt.ScheduledAt.Add(time.Duration(durationMin)*time.Minute + TelemedWindowAfter)
	t := models.TelemedLobbyToken{
		AppointmentID: appt.ID,
		Token:         token,
		ExpiresAt:     expiresAt,
	}
	if err := s.db.Create(&t).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

// LobbyView é o payload retornado pra página standalone /sala/[token].
// Inclui apenas o necessário pra renderizar a entrada — sem PHI.
//
// HIGH H9 — DailyJoinURL substitui DailyRoomURL: vem com meeting_token
// embutido (?t=...) escopado ao paciente. Sala é privacy=private — URL sem
// token retorna 401 mesmo se vazada.
type LobbyView struct {
	AppointmentID    uuid.UUID `json:"appointmentId"`
	PatientFirstName string    `json:"patientFirstName"` // só primeiro nome (opaco)
	DoctorName       string    `json:"doctorName"`
	ScheduledAt      time.Time `json:"scheduledAt"`
	DurationMinutes  int       `json:"durationMinutes"`
	OpensAt          time.Time `json:"opensAt"`  // -TelemedWindowBefore do scheduledAt
	ClosesAt         time.Time `json:"closesAt"` // +duration+TelemedWindowAfter do scheduledAt
	DailyJoinURL     *string   `json:"dailyJoinUrl,omitempty"` // só dentro da janela, com meeting_token
	IsOpen           bool      `json:"isOpen"`
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

	// Janela: -TelemedWindowBefore até +duration+TelemedWindowAfter do scheduledAt.
	durationMin := appt.DurationMinutes
	if durationMin <= 0 {
		durationMin = 60
	}
	opensAt := appt.ScheduledAt.Add(-TelemedWindowBefore)
	closesAt := appt.ScheduledAt.Add(time.Duration(durationMin)*time.Minute + TelemedWindowAfter)
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
	if isOpen && appt.DailyRoomURL != nil && appt.DailyRoomName != nil &&
		s.dailyCoSvc != nil && s.dailyCoSvc.IsConfigured() {
		// HIGH H9: gera meeting_token escopado ao paciente. Token vale até
		// closesAt — vazamento depois disso é inócuo.
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		tok, terr := s.dailyCoSvc.CreateMeetingToken(ctx, MeetingTokenParams{
			RoomName:          *appt.DailyRoomName,
			UserName:          view.PatientFirstName,
			IsOwner:           false,
			ExpiresAt:         closesAt,
			EnableScreenshare: false,
		})
		if terr != nil {
			// Degradação grácil: paciente vê "sala temporariamente indisponível".
			log.Printf("⚠️  [LOBBY] CreateMeetingToken apt=%s: %v", appt.ID, terr)
		} else {
			joinURL := BuildJoinURL(*appt.DailyRoomURL, tok)
			view.DailyJoinURL = &joinURL
		}
	} else if isOpen && appt.DailyRoomURL != nil &&
		(s.dailyCoSvc == nil || !s.dailyCoSvc.IsConfigured()) {
		// Backcompat: se Daily não configurado (dev), expõe URL crua.
		view.DailyJoinURL = appt.DailyRoomURL
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
