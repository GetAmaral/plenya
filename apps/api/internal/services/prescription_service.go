package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

var (
	ErrPrescriptionNotFound = errors.New("prescription not found")
)

type PrescriptionService struct {
	db *gorm.DB
}

func NewPrescriptionService(db *gorm.DB) *PrescriptionService {
	return &PrescriptionService{db: db}
}

// calculateValidUntil calcula a data de validade baseada nas categorias dos medicamentos
// Retorna a data MAIS RESTRITIVA (menor validade) entre todos os medicamentos
func calculateValidUntil(prescriptionDate time.Time, medications []dto.MedicationRequest) time.Time {
	if len(medications) == 0 {
		return prescriptionDate.AddDate(0, 0, 30) // 30 dias padrão
	}

	// Mapear categorias para dias de validade
	categoryDays := map[models.MedicationCategory]int{
		models.MedCategoryAntibiotic: 10,  // 10 dias
		models.MedCategoryC1:         30,  // 30 dias
		models.MedCategoryC5:         30,  // 30 dias
		models.MedCategoryGLP1:       90,  // 90 dias
		models.MedCategorySimple:     30,  // 30 dias
	}

	// Encontrar a validade mais restritiva (menor número de dias)
	minDays := 365 // Começar com um valor alto
	for _, med := range medications {
		if days, ok := categoryDays[med.Category]; ok {
			if days < minDays {
				minDays = days
			}
		}
	}

	if minDays == 365 {
		minDays = 30 // Fallback
	}

	return prescriptionDate.AddDate(0, 0, minDays)
}

// Create cria uma nova prescrição com múltiplos medicamentos
func (s *PrescriptionService) Create(doctorID uuid.UUID, req *dto.CreatePrescriptionRequest) (*dto.PrescriptionResponse, error) {
	// CRITICAL SECURITY: Get user's selected patient
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, doctorID).Error; err != nil {
		return nil, err
	}

	// If no selected patient, cannot create prescription
	if user.SelectedPatientID == nil {
		return nil, errors.New("no patient selected - please select a patient first")
	}

	// Parse patient ID from request
	var patientID uuid.UUID
	if req.PatientID != "" {
		pid, err := uuid.Parse(req.PatientID)
		if err != nil {
			return nil, errors.New("invalid patient id")
		}
		// SECURITY: Validate that patientID matches selectedPatient
		if pid != *user.SelectedPatientID {
			return nil, errors.New("patient id does not match selected patient")
		}
		patientID = pid
	} else {
		// Auto-fill with selectedPatient
		patientID = *user.SelectedPatientID
	}

	// Verificar se o paciente existe
	var patient models.Patient
	if err := s.db.First(&patient, patientID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPatientNotFound
		}
		return nil, err
	}

	// Validar medications array
	if len(req.Medications) == 0 {
		return nil, errors.New("prescription must have at least one medication")
	}
	if len(req.Medications) > 10 {
		return nil, errors.New("prescription cannot have more than 10 medications")
	}

	// Validar max 3 medicamentos controlados C1
	c1Count := 0
	for _, med := range req.Medications {
		if med.Category == models.MedCategoryC1 {
			c1Count++
		}
	}
	if c1Count > 3 {
		return nil, errors.New("prescription cannot have more than 3 controlled (C1) medications")
	}

	// Parse prescription date
	prescriptionDate, err := time.Parse(time.RFC3339, req.PrescriptionDate)
	if err != nil {
		return nil, errors.New("invalid prescription date format, expected RFC3339")
	}

	// Calcular ValidUntil baseado nas categorias dos medicamentos (mais restritiva)
	validUntil := calculateValidUntil(prescriptionDate, req.Medications)

	// Criar prescrição (sem medicamentos - serão criados separadamente)
	prescription := models.Prescription{
		PatientID:           patientID,
		DoctorID:            doctorID,
		GeneralInstructions: req.GeneralInstructions,
		Status:              models.PrescriptionActive,
		PrescriptionDate:    prescriptionDate,
		ValidUntil:          validUntil,
		IsUsed:              false,
	}

	if err := s.db.Create(&prescription).Error; err != nil {
		return nil, err
	}

	// Criar medicamentos da prescrição
	for _, medReq := range req.Medications {
		// Parse medication definition ID se fornecido
		var medicationDefID *uuid.UUID
		if medReq.MedicationDefinitionID != nil {
			medID, err := uuid.Parse(*medReq.MedicationDefinitionID)
			if err != nil {
				// Rollback da prescrição em caso de erro
				s.db.Delete(&prescription)
				return nil, errors.New("invalid medication definition id")
			}
			medicationDefID = &medID
		}

		medication := models.PrescriptionMedication{
			PrescriptionID:         prescription.ID,
			MedicationDefinitionID: medicationDefID,
			MedicationName:         medReq.MedicationName,
			ActiveIngredient:       medReq.ActiveIngredient,
			Category:               medReq.Category,
			Concentration:          medReq.Concentration,
			Dosage:                 medReq.Dosage,
			Frequency:              medReq.Frequency,
			Route:                  medReq.Route,
			Duration:               medReq.Duration,
			Quantity:               medReq.Quantity,
			QuantityInWords:        medReq.QuantityInWords,
			Instructions:           medReq.Instructions,
		}

		if err := s.db.Create(&medication).Error; err != nil {
			// Rollback da prescrição e medicamentos já criados em caso de erro
			s.db.Delete(&prescription)
			return nil, err
		}
	}

	// Recarregar prescrição com medicamentos
	if err := s.db.Preload("Medications").First(&prescription, prescription.ID).Error; err != nil {
		return nil, err
	}

	return s.toDTO(&prescription), nil
}

// GetByID busca uma prescrição por ID
func (s *PrescriptionService) GetByID(prescriptionID, userID uuid.UUID, userRole models.Role) (*dto.PrescriptionResponse, error) {
	// CRITICAL SECURITY: Get user's selected patient
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}

	// If no selected patient, cannot access any prescription
	if user.SelectedPatientID == nil {
		return nil, ErrPrescriptionNotFound
	}

	// ALWAYS filter by selectedPatient (all roles including admin)
	var prescription models.Prescription
	query := s.db.Preload("Medications").Where("id = ?", prescriptionID).Where("patient_id = ?", *user.SelectedPatientID)

	if err := query.First(&prescription).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPrescriptionNotFound
		}
		return nil, err
	}

	return s.toDTO(&prescription), nil
}

// List lista prescrições com filtros
func (s *PrescriptionService) List(userID uuid.UUID, userRole models.Role, patientID *uuid.UUID, status *models.PrescriptionStatus, limit, offset int) ([]dto.PrescriptionResponse, error) {
	// CRITICAL SECURITY: Get user's selected patient
	var user models.User
	if err := s.db.Select("selected_patient_id").First(&user, userID).Error; err != nil {
		return nil, err
	}

	// If no selected patient, return empty list (security measure)
	if user.SelectedPatientID == nil {
		return []dto.PrescriptionResponse{}, nil
	}

	// ALWAYS filter by selectedPatient (all roles including admin)
	var prescriptions []models.Prescription
	query := s.db.Preload("Medications").Limit(limit).Offset(offset).Order("prescription_date DESC")
	query = query.Where("patient_id = ?", *user.SelectedPatientID)

	// Filtro por status
	if status != nil {
		query = query.Where("status = ?", *status)
	}

	if err := query.Find(&prescriptions).Error; err != nil {
		return nil, err
	}

	result := make([]dto.PrescriptionResponse, len(prescriptions))
	for i, p := range prescriptions {
		result[i] = *s.toDTO(&p)
	}

	return result, nil
}

// Update atualiza uma prescrição
func (s *PrescriptionService) Update(prescriptionID, doctorID uuid.UUID, userRole models.Role, req *dto.UpdatePrescriptionRequest) (*dto.PrescriptionResponse, error) {
	var prescription models.Prescription
	query := s.db.Preload("Medications").Where("id = ?", prescriptionID)

	// Apenas o médico que criou pode editar (ou admin)
	if userRole != models.RoleAdmin {
		query = query.Where("doctor_id = ?", doctorID)
	}

	if err := query.First(&prescription).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPrescriptionNotFound
		}
		return nil, err
	}

	// Atualizar campos da prescrição
	if req.Status != nil {
		prescription.Status = *req.Status
	}
	if req.GeneralInstructions != nil {
		prescription.GeneralInstructions = req.GeneralInstructions
	}

	// Se forneceu medications, substituir todos
	if req.Medications != nil {
		// Validar array
		if len(*req.Medications) == 0 {
			return nil, errors.New("prescription must have at least one medication")
		}
		if len(*req.Medications) > 10 {
			return nil, errors.New("prescription cannot have more than 10 medications")
		}

		// Validar max 3 medicamentos controlados C1
		c1Count := 0
		for _, med := range *req.Medications {
			if med.Category == models.MedCategoryC1 {
				c1Count++
			}
		}
		if c1Count > 3 {
			return nil, errors.New("prescription cannot have more than 3 controlled (C1) medications")
		}

		// Deletar medicamentos antigos
		if err := s.db.Where("prescription_id = ?", prescriptionID).Delete(&models.PrescriptionMedication{}).Error; err != nil {
			return nil, err
		}

		// Criar novos medicamentos
		for _, medReq := range *req.Medications {
			// Parse medication definition ID se fornecido
			var medicationDefID *uuid.UUID
			if medReq.MedicationDefinitionID != nil {
				medID, err := uuid.Parse(*medReq.MedicationDefinitionID)
				if err != nil {
					return nil, errors.New("invalid medication definition id")
				}
				medicationDefID = &medID
			}

			medication := models.PrescriptionMedication{
				PrescriptionID:         prescription.ID,
				MedicationDefinitionID: medicationDefID,
				MedicationName:         medReq.MedicationName,
				ActiveIngredient:       medReq.ActiveIngredient,
				Category:               medReq.Category,
				Concentration:          medReq.Concentration,
				Dosage:                 medReq.Dosage,
				Frequency:              medReq.Frequency,
				Route:                  medReq.Route,
				Duration:               medReq.Duration,
				Quantity:               medReq.Quantity,
				QuantityInWords:        medReq.QuantityInWords,
				Instructions:           medReq.Instructions,
			}

			if err := s.db.Create(&medication).Error; err != nil {
				return nil, err
			}
		}

		// Recalcular ValidUntil baseado nos novos medicamentos
		validUntil := calculateValidUntil(prescription.PrescriptionDate, *req.Medications)
		prescription.ValidUntil = validUntil
	}

	if err := s.db.Save(&prescription).Error; err != nil {
		return nil, err
	}

	// Recarregar com medications atualizados
	if err := s.db.Preload("Medications").First(&prescription, prescriptionID).Error; err != nil {
		return nil, err
	}

	return s.toDTO(&prescription), nil
}

// Delete faz soft delete de uma prescrição
func (s *PrescriptionService) Delete(prescriptionID, doctorID uuid.UUID, userRole models.Role) error {
	query := s.db.Where("id = ?", prescriptionID)

	// Apenas o médico que criou pode deletar (ou admin)
	if userRole != models.RoleAdmin {
		query = query.Where("doctor_id = ?", doctorID)
	}

	result := query.Delete(&models.Prescription{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrPrescriptionNotFound
	}

	return nil
}

// toDTO converte Prescription para PrescriptionResponse
func (s *PrescriptionService) toDTO(prescription *models.Prescription) *dto.PrescriptionResponse {
	resp := &dto.PrescriptionResponse{
		ID:                  prescription.ID.String(),
		PatientID:           prescription.PatientID.String(),
		DoctorID:            prescription.DoctorID.String(),
		GeneralInstructions: prescription.GeneralInstructions,
		Status:              prescription.Status,
		PrescriptionDate:    prescription.PrescriptionDate.Format(time.RFC3339),
		ValidUntil:          prescription.ValidUntil.Format("2006-01-02"),
		SNCRNumber:          prescription.SNCRNumber,
		SNCRStatus:          prescription.SNCRStatus,
		SignedPDFPath:       prescription.SignedPDFPath,
		SignedPDFHash:       prescription.SignedPDFHash,
		QRCodeData:          prescription.QRCodeData,
		CertificateSerial:   prescription.CertificateSerial,
		IsUsed:              prescription.IsUsed,
		CreatedAt:           prescription.CreatedAt.Format(time.RFC3339),
		UpdatedAt:           prescription.UpdatedAt.Format(time.RFC3339),
	}

	// Converter medications array
	medications := make([]dto.MedicationResponse, len(prescription.Medications))
	for i, med := range prescription.Medications {
		medResp := dto.MedicationResponse{
			ID:               med.ID.String(),
			MedicationName:   med.MedicationName,
			ActiveIngredient: med.ActiveIngredient,
			Category:         med.Category,
			Concentration:    med.Concentration,
			Dosage:           med.Dosage,
			Frequency:        med.Frequency,
			Route:            med.Route,
			Duration:         med.Duration,
			Quantity:         med.Quantity,
			QuantityInWords:  med.QuantityInWords,
			Instructions:     med.Instructions,
		}

		// Medication Definition ID
		if med.MedicationDefinitionID != nil {
			medDefID := med.MedicationDefinitionID.String()
			medResp.MedicationDefinitionID = &medDefID
		}

		medications[i] = medResp
	}
	resp.Medications = medications

	// Signed At
	if prescription.SignedAt != nil {
		signedAt := prescription.SignedAt.Format(time.RFC3339)
		resp.SignedAt = &signedAt
	}

	// Dispensed At
	if prescription.DispensedAt != nil {
		dispensedAt := prescription.DispensedAt.Format(time.RFC3339)
		resp.DispensedAt = &dispensedAt
	}

	return resp
}
