package services

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// PatientProfileService — perfil + LGPD do paciente no portal.
type PatientProfileService struct {
	db *gorm.DB
}

func NewPatientProfileService(db *gorm.DB) *PatientProfileService {
	return &PatientProfileService{db: db}
}

// UpdatableProfile contém SÓ os campos que o paciente pode editar.
// Nome, CPF, RG, BirthDate, Gender etc são imutáveis aqui (paciente fala
// com a equipe se houver erro nos dados de identificação).
type UpdatableProfile struct {
	Phone         *string `json:"phone,omitempty"`
	Email         *string `json:"email,omitempty"`
	Address       *string `json:"address,omitempty"`
	Municipality  *string `json:"municipality,omitempty"`
	State         *string `json:"state,omitempty"`
	EmergencyPhone *string `json:"emergencyPhone,omitempty"`
}

// UpdateProfile aplica patch parcial. Campos nil são preservados.
func (s *PatientProfileService) UpdateProfile(patientID uuid.UUID, p UpdatableProfile) error {
	updates := map[string]any{}
	if p.Phone != nil {
		v := strings.TrimSpace(*p.Phone)
		if v != "" {
			if normalized := models.NormalizePhoneBR(v); normalized != "" {
				v = normalized
			}
		}
		updates["phone"] = v
	}
	if p.Email != nil {
		v := strings.ToLower(strings.TrimSpace(*p.Email))
		if v != "" && !strings.Contains(v, "@") {
			return errors.New("email inválido")
		}
		updates["email"] = v
	}
	if p.Address != nil {
		updates["address"] = strings.TrimSpace(*p.Address)
	}
	if p.Municipality != nil {
		updates["municipality"] = strings.TrimSpace(*p.Municipality)
	}
	if p.State != nil {
		v := strings.ToUpper(strings.TrimSpace(*p.State))
		if v != "" && len(v) != 2 {
			return errors.New("UF inválida (2 letras)")
		}
		updates["state"] = v
	}
	if p.EmergencyPhone != nil {
		updates["emergency_phone"] = strings.TrimSpace(*p.EmergencyPhone)
	}
	if len(updates) == 0 {
		return nil
	}
	// UpdateColumns pula BeforeSave hook do Patient (que valida Gender,
	// criptografa CPF, etc — esses campos não estão no patch parcial).
	return s.db.Model(&models.Patient{}).
		Where("id = ?", patientID).
		UpdateColumns(updates).Error
}

// ============================================================
// LGPD — exportação e exclusão (V1: solicitação humana, não delete auto)
// ============================================================

// LGPDExport monta um JSON com todos os dados do paciente que temos.
// Resposta síncrona pra V1 (pacientes não têm volume gigante). V2 vai
// pra background + email com link de download.
func (s *PatientProfileService) LGPDExport(patientID uuid.UUID) (map[string]any, error) {
	var patient models.Patient
	if err := s.db.First(&patient, "id = ?", patientID).Error; err != nil {
		return nil, err
	}

	export := map[string]any{
		"_meta": map[string]any{
			"exportType":      "LGPD-portability",
			"exportVersion":   "1.0",
			"exportedAt":      time.Now().UTC(),
			"controller":      "Plenya Saúde Ltda.",
			"controllerEmail": "contato@plenyasaude.com.br",
			"dpoEmail":        "dpo@plenyasaude.com.br",
		},
		"patient": patient,
	}

	// Anamnesis
	var anamneses []map[string]any
	_ = s.db.Table("anamneses").Where("patient_id = ?", patientID).Find(&anamneses).Error
	export["anamnesis"] = anamneses

	// Lab batches + values
	var labBatches []models.LabResultBatch
	_ = s.db.Where("patient_id = ?", patientID).Find(&labBatches).Error
	export["labBatches"] = labBatches

	// Prescriptions
	var prescriptions []models.Prescription
	_ = s.db.Where("patient_id = ?", patientID).Preload("Medications").Find(&prescriptions).Error
	export["prescriptions"] = prescriptions

	// Appointments
	var appointments []models.Appointment
	_ = s.db.Where("patient_id = ?", patientID).Find(&appointments).Error
	export["appointments"] = appointments

	// Score snapshots
	var snapshots []models.PatientScoreSnapshot
	_ = s.db.Where("patient_id = ?", patientID).
		Preload("GroupResults.Group").
		Find(&snapshots).Error
	export["scoreSnapshots"] = snapshots

	// Light sessions claimed
	var lightSessions []models.AnonymousScoreSession
	_ = s.db.Preload("Snapshot.GroupResults.Group").
		Where("claimed_by_patient_id = ?", patientID).
		Find(&lightSessions).Error
	export["lightSessions"] = lightSessions

	// Continuum
	var continuums []models.PatientContinuum
	_ = s.db.Where("patient_id = ?", patientID).Preload("Items").Find(&continuums).Error
	export["continuums"] = continuums

	// Lead activities (mensagens com a equipe)
	var activities []models.LeadActivity
	_ = s.db.Where("patient_id = ?", patientID).Find(&activities).Error
	export["messagesAndActivities"] = activities

	// Physical assessments
	var assessments []models.PhysicalAssessment
	_ = s.db.Where("patient_id = ?", patientID).Find(&assessments).Error
	export["physicalAssessments"] = assessments

	return export, nil
}

// LGPDDeleteRequestInput agrupa input do POST /lgpd/account-delete-request.
type LGPDDeleteRequestInput struct {
	PatientID uuid.UUID
	UserID    uuid.UUID
	Reason    string
}

// RequestAccountDeletion cria notificação pra DPO/admin processar manualmente.
// V1 não deleta dados — exclusão automática traz risco enorme em EMR (LGPD permite
// retenção pra cumprir regulamentações sanitárias por 20 anos). DPO avalia caso a caso.
func (s *PatientProfileService) RequestAccountDeletion(in LGPDDeleteRequestInput, notification *NotificationService) error {
	if notification == nil {
		return errors.New("notification service indisponível")
	}
	var patient models.Patient
	if err := s.db.Select("name").First(&patient, "id = ?", in.PatientID).Error; err != nil {
		return err
	}
	// Targets: admins
	var admins []models.User
	_ = s.db.Where("roles::text LIKE ?", "%admin%").Limit(10).Find(&admins).Error

	title := fmt.Sprintf("Solicitação de exclusão LGPD: %s", patient.Name)
	msg := fmt.Sprintf("Paciente solicitou exclusão de conta. Motivo: %s", in.Reason)
	if strings.TrimSpace(in.Reason) == "" {
		msg = "Paciente solicitou exclusão de conta. (sem motivo informado)"
	}
	actionURL := fmt.Sprintf("/patients/%s", in.PatientID.String())
	actionText := "Abrir paciente"
	for i := range admins {
		_ = notification.CreateNotification(
			admins[i].ID,
			models.NotificationGeneral,
			title,
			msg,
			&in.PatientID,
			nil,
			&actionURL,
			&actionText,
		)
	}
	return nil
}
