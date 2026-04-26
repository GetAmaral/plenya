package services

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// NotificationPreferencesService faz get-or-create + patch das prefs do usuário.
// Padrão usado por todos os endpoints de push trigger pra decidir se manda ou não.
type NotificationPreferencesService struct {
	db *gorm.DB
}

func NewNotificationPreferencesService(db *gorm.DB) *NotificationPreferencesService {
	return &NotificationPreferencesService{db: db}
}

// Get devolve as prefs (ou cria com defaults se não existirem).
// Default = todos os toggles em true, lembrete às 07:00.
func (s *NotificationPreferencesService) Get(userID uuid.UUID) (*models.NotificationPreferences, error) {
	var prefs models.NotificationPreferences
	err := s.db.Where("user_id = ?", userID).First(&prefs).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		prefs = models.NotificationPreferences{
			UserID:              userID,
			AppointmentReminder: true,
			MessageAlert:        true,
			WorkoutReminder:     true,
			WorkoutReminderTime: "07:00",
		}
		if err := s.db.Create(&prefs).Error; err != nil {
			return nil, err
		}
		return &prefs, nil
	}
	if err != nil {
		return nil, err
	}
	return &prefs, nil
}

// PatchInput é o payload do PATCH /notification-preferences.
// Campos nullable = não-mexa-nesse.
type PatchInput struct {
	AppointmentReminder *bool   `json:"appointmentReminder,omitempty"`
	MessageAlert        *bool   `json:"messageAlert,omitempty"`
	WorkoutReminder     *bool   `json:"workoutReminder,omitempty"`
	WorkoutReminderTime *string `json:"workoutReminderTime,omitempty"`
}

// Patch aplica updates parciais.
func (s *NotificationPreferencesService) Patch(userID uuid.UUID, in PatchInput) (*models.NotificationPreferences, error) {
	prefs, err := s.Get(userID)
	if err != nil {
		return nil, err
	}
	if in.AppointmentReminder != nil {
		prefs.AppointmentReminder = *in.AppointmentReminder
	}
	if in.MessageAlert != nil {
		prefs.MessageAlert = *in.MessageAlert
	}
	if in.WorkoutReminder != nil {
		prefs.WorkoutReminder = *in.WorkoutReminder
	}
	if in.WorkoutReminderTime != nil {
		prefs.WorkoutReminderTime = *in.WorkoutReminderTime
	}
	if err := s.db.Save(prefs).Error; err != nil {
		return nil, err
	}
	return prefs, nil
}
