package services

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
)

var (
	ErrPrescriptionNotFound = errors.New("prescription not found")
	ErrPrescriptionSigned   = errors.New("prescription already signed and cannot be changed")
)

type PrescriptionService struct {
	db *gorm.DB
}

func NewPrescriptionService(db *gorm.DB) *PrescriptionService {
	return &PrescriptionService{db: db}
}

// validityDaysByCategory é o fallback de validade para medicamento digitado à mão (sem vínculo
// com o catálogo). Quando o item veio do catálogo, quem manda é a coluna validity_days da
// definição — que a curadoria pode ajustar sem recompilar a API.
var validityDaysByCategory = map[models.MedicationCategory]int{
	models.MedCategoryAntibiotic: 10,
	models.MedCategoryC1:         30,
	models.MedCategoryC5:         30,
	models.MedCategoryGLP1:       90,
	models.MedCategorySimple:     30,
	models.MedCategoryAB:         30,
}

const defaultValidityDays = 30

// calculateValidUntil devolve a data de validade MAIS RESTRITIVA entre os medicamentos.
// `catalogDays` mapeia medication_definition_id → validity_days do catálogo; o que não estiver
// lá cai na tabela por categoria.
func calculateValidUntil(prescriptionDate time.Time, medications []dto.MedicationRequest, catalogDays map[uuid.UUID]int) time.Time {
	if len(medications) == 0 {
		return prescriptionDate.AddDate(0, 0, defaultValidityDays)
	}

	minDays := 0
	for _, med := range medications {
		days := defaultValidityDays
		if d, ok := validityDaysByCategory[med.Category]; ok {
			days = d
		}
		if med.MedicationDefinitionID != nil {
			if id, err := uuid.Parse(*med.MedicationDefinitionID); err == nil {
				if d, ok := catalogDays[id]; ok && d > 0 {
					days = d
				}
			}
		}
		if minDays == 0 || days < minDays {
			minDays = days
		}
	}

	if minDays == 0 {
		minDays = defaultValidityDays
	}

	return prescriptionDate.AddDate(0, 0, minDays)
}

// catalogValidityDays lê a validade das definições do catálogo referenciadas pela receita.
// Falha de leitura não derruba a prescrição: o mapa volta vazio e a validade cai na categoria.
func (s *PrescriptionService) catalogValidityDays(tx *gorm.DB, medications []dto.MedicationRequest) map[uuid.UUID]int {
	ids := make([]uuid.UUID, 0, len(medications))
	for _, med := range medications {
		if med.MedicationDefinitionID == nil {
			continue
		}
		if id, err := uuid.Parse(*med.MedicationDefinitionID); err == nil {
			ids = append(ids, id)
		}
	}
	out := map[uuid.UUID]int{}
	if len(ids) == 0 {
		return out
	}

	var rows []struct {
		ID           uuid.UUID
		ValidityDays int
	}
	if err := tx.Model(&models.MedicationDefinition{}).
		Select("id, validity_days").
		Where("id IN ?", ids).
		Scan(&rows).Error; err != nil {
		return out
	}
	for _, r := range rows {
		out[r.ID] = r.ValidityDays
	}
	return out
}

// distinctControlledSubstances conta SUBSTÂNCIAS de controle especial, não itens: a Portaria
// 344/98 limita o receituário a 3 substâncias. Contar itens recusava receita legítima (duas
// apresentações da mesma substância) e aceitava receita irregular (mesma substância com nomes
// comerciais diferentes).
func distinctControlledSubstances(medications []dto.MedicationRequest) int {
	seen := map[string]struct{}{}
	for _, med := range medications {
		if med.Category != models.MedCategoryC1 {
			continue
		}
		key := strings.ToLower(strings.TrimSpace(med.ActiveIngredient))
		if key == "" {
			key = strings.ToLower(strings.TrimSpace(med.MedicationName))
		}
		if key == "" {
			continue
		}
		seen[key] = struct{}{}
	}
	return len(seen)
}

// validateMedications aplica as regras comuns a Create e Update.
func validateMedications(medications []dto.MedicationRequest) error {
	if len(medications) == 0 {
		return errors.New("prescription must have at least one medication")
	}
	if len(medications) > 10 {
		return errors.New("prescription cannot have more than 10 medications")
	}
	if distinctControlledSubstances(medications) > 3 {
		return errors.New("prescription cannot have more than 3 controlled (C1) substances")
	}
	return nil
}

// buildMedications converte o request em models, validando os IDs do catálogo antes de gravar
// qualquer coisa — assim um UUID inválido não deixa prescrição pela metade no banco.
func buildMedications(prescriptionID uuid.UUID, medications []dto.MedicationRequest) ([]models.PrescriptionMedication, error) {
	out := make([]models.PrescriptionMedication, 0, len(medications))
	for _, medReq := range medications {
		var medicationDefID *uuid.UUID
		if medReq.MedicationDefinitionID != nil && *medReq.MedicationDefinitionID != "" {
			medID, err := uuid.Parse(*medReq.MedicationDefinitionID)
			if err != nil {
				return nil, errors.New("invalid medication definition id")
			}
			medicationDefID = &medID
		}

		out = append(out, models.PrescriptionMedication{
			PrescriptionID:         prescriptionID,
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
		})
	}
	return out, nil
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

	// Tipo da receita: ausente = industrializado (compatível com todo cliente anterior).
	prescriptionType := req.Type
	if prescriptionType == "" {
		prescriptionType = models.PrescriptionCommercial
	}

	// Cada tipo preenche uma tabela filha diferente; misturar os dois numa receita só produziria
	// um PDF com dois receituários dentro.
	switch prescriptionType {
	case models.PrescriptionCommercial:
		if len(req.Formulas) > 0 {
			return nil, errors.New("receita de industrializado não aceita fórmulas manipuladas")
		}
		if err := validateMedications(req.Medications); err != nil {
			return nil, err
		}
	case models.PrescriptionCompounded:
		if len(req.Medications) > 0 {
			return nil, errors.New("receita de manipulado não aceita medicamentos industrializados")
		}
		if err := validateFormulas(req.Formulas); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("tipo de prescrição inválido")
	}

	// Parse prescription date
	prescriptionDate, err := time.Parse(time.RFC3339, req.PrescriptionDate)
	if err != nil {
		return nil, errors.New("invalid prescription date format, expected RFC3339")
	}

	prescription := models.Prescription{
		PatientID:           patientID,
		DoctorID:            doctorID,
		Type:                prescriptionType,
		GeneralInstructions: req.GeneralInstructions,
		Status:              models.PrescriptionActive,
		PrescriptionDate:    prescriptionDate,
		IsUsed:              false,
	}

	// Prescrição e filhos numa transação só. Antes era um loop de Creates soltos cujo "rollback"
	// era um soft delete manual da prescrição — que deixava os medicamentos já criados órfãos no
	// banco quando o erro caía no meio.
	if err := s.db.Transaction(func(tx *gorm.DB) error {
		if prescriptionType == models.PrescriptionCompounded {
			formulas, err := buildFormulas(uuid.Nil, req.Formulas)
			if err != nil {
				return err
			}
			prescription.ValidUntil = calculateValidUntilFormulas(prescriptionDate, formulas)

			if err := tx.Create(&prescription).Error; err != nil {
				return err
			}
			// Curadoria oportunista: substância nova entra no catálogo magistral como esboço e
			// as conhecidas ganham +1 no contador de uso — é o que faz a busca ir ficando
			// parecida com o repertório real sem ninguém abrir tela de cadastro.
			if err := NewMagistralComponentService(s.db).EnsureComponents(tx, formulas); err != nil {
				return err
			}
			// Os componentes vêm aninhados: o GORM grava fórmula e componentes juntos.
			templates := NewMagistralTemplateService(s.db)
			for i := range formulas {
				formulas[i].PrescriptionID = prescription.ID
				// A fórmula-base sobe na lista conforme vira receita de verdade.
				if formulas[i].TemplateID != nil {
					if err := templates.IncrementUsage(tx, *formulas[i].TemplateID); err != nil {
						return err
					}
				}
			}
			if err := tx.Create(&formulas).Error; err != nil {
				return err
			}
			return s.reloadPrescription(tx, &prescription)
		}

		prescription.ValidUntil = calculateValidUntil(prescriptionDate, req.Medications, s.catalogValidityDays(tx, req.Medications))

		if err := tx.Create(&prescription).Error; err != nil {
			return err
		}

		medications, err := buildMedications(prescription.ID, req.Medications)
		if err != nil {
			return err
		}
		if err := tx.Create(&medications).Error; err != nil {
			return err
		}

		return s.reloadPrescription(tx, &prescription)
	}); err != nil {
		return nil, err
	}

	return s.toDTO(&prescription), nil
}

// reloadPrescription recarrega a receita com os filhos na ordem em que são impressos.
func (s *PrescriptionService) reloadPrescription(tx *gorm.DB, prescription *models.Prescription) error {
	return tx.
		Preload("Medications").
		Preload("Formulas", func(db *gorm.DB) *gorm.DB { return db.Order("display_order") }).
		Preload("Formulas.Components", func(db *gorm.DB) *gorm.DB { return db.Order("display_order") }).
		First(prescription, prescription.ID).Error
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
	query := s.db.
		Preload("Medications").
		Preload("Formulas", func(db *gorm.DB) *gorm.DB { return db.Order("display_order") }).
		Preload("Formulas.Components", func(db *gorm.DB) *gorm.DB { return db.Order("display_order") }).
		Where("id = ?", prescriptionID).Where("patient_id = ?", *user.SelectedPatientID)

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
	query := s.db.
		Preload("Medications").
		Preload("Formulas", func(db *gorm.DB) *gorm.DB { return db.Order("display_order") }).
		Preload("Formulas.Components", func(db *gorm.DB) *gorm.DB { return db.Order("display_order") }).
		Limit(limit).Offset(offset).Order("prescription_date DESC")
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

	if err := s.db.Transaction(func(tx *gorm.DB) error {
		query := tx.Preload("Medications").Preload("Formulas.Components").Where("id = ?", prescriptionID)

		// Apenas o médico que criou pode editar (ou admin)
		if userRole != models.RoleAdmin {
			query = query.Where("doctor_id = ?", doctorID)
		}

		if err := query.First(&prescription).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrPrescriptionNotFound
			}
			return err
		}

		// Receita assinada é documento emitido: alterar o CONTEÚDO depois da assinatura
		// invalidaria o PDF e o hash.
		//
		// Mudar só o status é outra coisa, e precisa continuar possível: é assim que uma receita
		// emitida vira "cancelada" ou "concluída". A primeira versão desta guarda recusava tudo,
		// e o próprio erro do Delete manda usar o status cancelada — uma transição que nenhum
		// caminho conseguia fazer.
		if isSigned(&prescription) {
			if mudaConteudo(req) {
				return ErrPrescriptionSigned
			}
		}

		if req.Status != nil {
			prescription.Status = *req.Status
		}
		if req.GeneralInstructions != nil {
			prescription.GeneralInstructions = req.GeneralInstructions
		}

		// O tipo é imutável: trocar industrializado por manipulado em cima de uma receita
		// existente misturaria os dois receituários.
		if req.Medications != nil && prescription.Type == models.PrescriptionCompounded {
			return errors.New("receita de manipulado não aceita medicamentos industrializados")
		}
		if req.Formulas != nil && prescription.Type != models.PrescriptionCompounded {
			return errors.New("receita de industrializado não aceita fórmulas manipuladas")
		}

		// Se forneceu fórmulas, substituir todas
		if req.Formulas != nil {
			if err := validateFormulas(*req.Formulas); err != nil {
				return err
			}

			formulas, err := buildFormulas(prescription.ID, *req.Formulas)
			if err != nil {
				return err
			}

			// Componentes primeiro: a FK é ON DELETE CASCADE, mas o delete das fórmulas com
			// Unscoped não dispara o cascade sobre linhas soft-deletadas.
			var oldIDs []uuid.UUID
			if err := tx.Model(&models.PrescriptionFormula{}).Unscoped().
				Where("prescription_id = ?", prescriptionID).Pluck("id", &oldIDs).Error; err != nil {
				return err
			}
			if len(oldIDs) > 0 {
				if err := tx.Unscoped().Where("formula_id IN ?", oldIDs).
					Delete(&models.PrescriptionFormulaComponent{}).Error; err != nil {
					return err
				}
				if err := tx.Unscoped().Where("prescription_id = ?", prescriptionID).
					Delete(&models.PrescriptionFormula{}).Error; err != nil {
					return err
				}
			}
			if err := NewMagistralComponentService(s.db).EnsureComponents(tx, formulas); err != nil {
				return err
			}
			if err := tx.Create(&formulas).Error; err != nil {
				return err
			}

			prescription.ValidUntil = calculateValidUntilFormulas(prescription.PrescriptionDate, formulas)
			// A associação carregada aponta para as fórmulas apagadas; sem limpar, o Save
			// abaixo tentaria ressuscitá-las.
			prescription.Formulas = nil
		}

		// Se forneceu medications, substituir todos
		if req.Medications != nil {
			if err := validateMedications(*req.Medications); err != nil {
				return err
			}

			medications, err := buildMedications(prescription.ID, *req.Medications)
			if err != nil {
				return err
			}

			// Hard delete: o histórico do que foi prescrito vive no PDF assinado, e rascunho
			// editado não precisa acumular linhas soft-deletadas com o mesmo prescription_id.
			if err := tx.Unscoped().Where("prescription_id = ?", prescriptionID).
				Delete(&models.PrescriptionMedication{}).Error; err != nil {
				return err
			}
			if err := tx.Create(&medications).Error; err != nil {
				return err
			}

			prescription.ValidUntil = calculateValidUntil(
				prescription.PrescriptionDate,
				*req.Medications,
				s.catalogValidityDays(tx, *req.Medications),
			)
		}

		// Omit das associações: Save com associação carregada faz upsert dos filhos, e o GORM
		// ressuscitaria justamente os itens que acabamos de apagar (a receita voltava com os
		// antigos + os novos).
		if err := tx.Omit(clause.Associations).Save(&prescription).Error; err != nil {
			return err
		}

		return s.reloadPrescription(tx, &prescription)
	}); err != nil {
		return nil, err
	}

	return s.toDTO(&prescription), nil
}

// Delete faz soft delete de uma prescrição
func (s *PrescriptionService) Delete(prescriptionID, doctorID uuid.UUID, userRole models.Role) error {
	var prescription models.Prescription
	query := s.db.Where("id = ?", prescriptionID)

	// Apenas o médico que criou pode deletar (ou admin)
	if userRole != models.RoleAdmin {
		query = query.Where("doctor_id = ?", doctorID)
	}

	if err := query.First(&prescription).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrPrescriptionNotFound
		}
		return err
	}

	// Receita assinada não se apaga: o documento foi emitido, publicado ao paciente e pode ter
	// sido dispensado. Para encerrar uma receita válida existe o status "cancelled".
	if isSigned(&prescription) {
		return ErrPrescriptionSigned
	}

	return s.db.Delete(&models.Prescription{}, "id = ?", prescriptionID).Error
}

// mudaConteudo diz se o pedido mexe em algo além do status.
func mudaConteudo(req *dto.UpdatePrescriptionRequest) bool {
	return req.Medications != nil || req.Formulas != nil || req.GeneralInstructions != nil
}

// isSigned diz se a receita já virou documento emitido — assinatura digital ou impressão para
// assinatura à mão, que também gera PDF e publica o documento ao paciente.
func isSigned(p *models.Prescription) bool {
	return p.SignedAt != nil || (p.SignedPDFPath != nil && *p.SignedPDFPath != "")
}

// toDTO converte Prescription para PrescriptionResponse
func (s *PrescriptionService) toDTO(prescription *models.Prescription) *dto.PrescriptionResponse {
	resp := &dto.PrescriptionResponse{
		ID:                  prescription.ID.String(),
		PatientID:           prescription.PatientID.String(),
		DoctorID:            prescription.DoctorID.String(),
		Type:                prescription.Type,
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
		SignatureMode:       prescription.SignatureMode,
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
	resp.Formulas = formulasToDTO(prescription.Formulas)

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
