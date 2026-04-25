package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// PatientLabsService — visão paciente de lab batches, lab requests, prescrições
// e avaliações físicas. Sempre filtra por patientID — paciente nunca vê de outros.
type PatientLabsService struct {
	db *gorm.DB
}

func NewPatientLabsService(db *gorm.DB) *PatientLabsService {
	return &PatientLabsService{db: db}
}

// ============================================================
// Lab Batches (resultados)
// ============================================================

type LabBatchSummary struct {
	ID               uuid.UUID  `json:"id"`
	LaboratoryName   string     `json:"laboratoryName"`
	CollectionDate   time.Time  `json:"collectionDate"`
	ResultDate       *time.Time `json:"resultDate,omitempty"`
	Status           string     `json:"status"`
	RequestingDoctor *string    `json:"requestingDoctor,omitempty"`
	ResultsCount     int        `json:"resultsCount"`
}

func (s *PatientLabsService) ListBatches(patientID uuid.UUID) ([]LabBatchSummary, error) {
	var batches []models.LabResultBatch
	err := s.db.Where("patient_id = ?", patientID).
		Order("collection_date DESC").
		Find(&batches).Error
	if err != nil {
		return nil, err
	}

	out := make([]LabBatchSummary, 0, len(batches))
	for i := range batches {
		b := &batches[i]
		summary := LabBatchSummary{
			ID:             b.ID,
			LaboratoryName: b.LaboratoryName,
			CollectionDate: b.CollectionDate,
			ResultDate:     b.ResultDate,
			Status:         string(b.Status),
		}
		if b.RequestingDoctorID != nil {
			var doctor models.User
			if err := s.db.Select("name").First(&doctor, "id = ?", *b.RequestingDoctorID).Error; err == nil {
				name := doctor.Name
				summary.RequestingDoctor = &name
			}
		}
		var n int64
		_ = s.db.Model(&models.LabResult{}).
			Where("batch_id = ?", b.ID).
			Count(&n).Error
		summary.ResultsCount = int(n)
		out = append(out, summary)
	}
	return out, nil
}

type LabValueView struct {
	ID        uuid.UUID `json:"id"`
	TestName  string    `json:"testName"`
	TestCode  string    `json:"testCode"`
	Value     *float64  `json:"value,omitempty"`
	ValueText string    `json:"valueText,omitempty"`
	Unit      string    `json:"unit,omitempty"`
}

type LabBatchDetail struct {
	LabBatchSummary
	Observations *string        `json:"observations,omitempty"`
	Values       []LabValueView `json:"values"`
}

func (s *PatientLabsService) GetBatch(patientID, batchID uuid.UUID) (*LabBatchDetail, error) {
	var batch models.LabResultBatch
	err := s.db.Where("id = ? AND patient_id = ?", batchID, patientID).First(&batch).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("batch não encontrado")
		}
		return nil, err
	}

	// LabResult > LabResultValue → traz tudo numa só query
	var results []models.LabResult
	_ = s.db.Where("batch_id = ?", batchID).Find(&results).Error

	resultIDs := make([]uuid.UUID, 0, len(results))
	for i := range results {
		resultIDs = append(resultIDs, results[i].ID)
	}

	var values []models.LabResultValue
	if len(resultIDs) > 0 {
		_ = s.db.Where("lab_result_id IN ?", resultIDs).
			Preload("LabTestDefinition").
			Order("created_at ASC").
			Find(&values).Error
	}

	out := &LabBatchDetail{
		LabBatchSummary: LabBatchSummary{
			ID:             batch.ID,
			LaboratoryName: batch.LaboratoryName,
			CollectionDate: batch.CollectionDate,
			ResultDate:     batch.ResultDate,
			Status:         string(batch.Status),
			ResultsCount:   len(results),
		},
		Observations: batch.Observations,
		Values:       []LabValueView{},
	}
	if batch.RequestingDoctorID != nil {
		var doctor models.User
		if err := s.db.Select("name").First(&doctor, "id = ?", *batch.RequestingDoctorID).Error; err == nil {
			name := doctor.Name
			out.RequestingDoctor = &name
		}
	}

	for i := range values {
		v := &values[i]
		view := LabValueView{
			ID:        v.ID,
			Value:     v.NumericValue,
			ValueText: derefStr(v.TextValue),
			Unit:      derefStr(v.Unit),
			TestName:  v.LabTestDefinition.Name,
			TestCode:  v.LabTestDefinition.Code,
		}
		out.Values = append(out.Values, view)
	}
	return out, nil
}

// ============================================================
// Lab Requests (pedidos de exame pendentes pra coletar)
// ============================================================

type LabRequestSummary struct {
	ID         uuid.UUID  `json:"id"`
	IssuedAt   time.Time  `json:"issuedAt"`
	Date       time.Time  `json:"date"`
	DoctorName string     `json:"doctorName,omitempty"`
	Exams      string     `json:"exams"`
	SignedAt   *time.Time `json:"signedAt,omitempty"`
	PdfURL     *string    `json:"pdfUrl,omitempty"`
}

func (s *PatientLabsService) ListLabRequests(patientID uuid.UUID) ([]LabRequestSummary, error) {
	var rows []models.LabRequest
	err := s.db.Where("patient_id = ?", patientID).
		Order("created_at DESC").
		Find(&rows).Error
	if err != nil {
		return nil, err
	}
	out := make([]LabRequestSummary, 0, len(rows))
	for i := range rows {
		r := &rows[i]
		summary := LabRequestSummary{
			ID:       r.ID,
			IssuedAt: r.CreatedAt,
			Date:     r.Date,
			Exams:    r.Exams,
			SignedAt: r.SignedAt,
			PdfURL:   r.PdfURL,
		}
		if r.DoctorID != nil {
			var doctor models.User
			if err := s.db.Select("name").First(&doctor, "id = ?", *r.DoctorID).Error; err == nil {
				summary.DoctorName = doctor.Name
			}
		}
		out = append(out, summary)
	}
	return out, nil
}

// ============================================================
// Prescrições
// ============================================================

type PrescriptionSummary struct {
	ID         uuid.UUID  `json:"id"`
	IssuedAt   time.Time  `json:"issuedAt"`
	Status     string     `json:"status"`
	DoctorName string     `json:"doctorName,omitempty"`
	ValidUntil time.Time  `json:"validUntil"`
	MedsCount  int        `json:"medsCount"`
	SignedAt   *time.Time `json:"signedAt,omitempty"`
	PdfPath    *string    `json:"pdfPath,omitempty"`
}

func (s *PatientLabsService) ListPrescriptions(patientID uuid.UUID) ([]PrescriptionSummary, error) {
	var rows []models.Prescription
	err := s.db.Where("patient_id = ?", patientID).
		Order("created_at DESC").
		Find(&rows).Error
	if err != nil {
		return nil, err
	}
	out := make([]PrescriptionSummary, 0, len(rows))
	for i := range rows {
		r := &rows[i]
		summary := PrescriptionSummary{
			ID:         r.ID,
			IssuedAt:   r.PrescriptionDate,
			Status:     string(r.Status),
			ValidUntil: r.ValidUntil,
			SignedAt:   r.SignedAt,
			PdfPath:    r.SignedPDFPath,
		}
		var doctor models.User
		if err := s.db.Select("name").First(&doctor, "id = ?", r.DoctorID).Error; err == nil {
			summary.DoctorName = doctor.Name
		}
		var n int64
		_ = s.db.Model(&models.PrescriptionMedication{}).
			Where("prescription_id = ?", r.ID).
			Count(&n).Error
		summary.MedsCount = int(n)
		out = append(out, summary)
	}
	return out, nil
}

// ============================================================
// Avaliações físicas
// ============================================================

type PhysicalAssessmentSummary struct {
	ID             uuid.UUID `json:"id"`
	AssessmentDate time.Time `json:"assessmentDate"`
	AssessorName   string    `json:"assessorName,omitempty"`
	Weight         *float64  `json:"weight,omitempty"`
	Height         *float64  `json:"height,omitempty"`
	BMI            *float64  `json:"bmi,omitempty"`
}

func (s *PatientLabsService) ListPhysicalAssessments(patientID uuid.UUID) ([]PhysicalAssessmentSummary, error) {
	var rows []models.PhysicalAssessment
	err := s.db.Where("patient_id = ?", patientID).
		Order("assessment_date DESC").
		Find(&rows).Error
	if err != nil {
		return nil, err
	}
	out := make([]PhysicalAssessmentSummary, 0, len(rows))
	for i := range rows {
		r := &rows[i]
		summary := PhysicalAssessmentSummary{
			ID:             r.ID,
			AssessmentDate: r.AssessmentDate,
			Weight:         r.Weight,
			Height:         r.Height,
			BMI:            r.BMI,
		}
		var assessor models.User
		if err := s.db.Select("name").First(&assessor, "id = ?", r.CreatedByID).Error; err == nil {
			summary.AssessorName = assessor.Name
		}
		out = append(out, summary)
	}
	return out, nil
}

// GetPhysicalAssessmentHTML retorna o HtmlContent de uma avaliação,
// já filtrado por patientID. Pra render iframe.
func (s *PatientLabsService) GetPhysicalAssessmentHTML(patientID, assessmentID uuid.UUID) (string, error) {
	var row models.PhysicalAssessment
	err := s.db.Where("id = ? AND patient_id = ?", assessmentID, patientID).
		Select("html_content").
		First(&row).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("avaliação não encontrada")
		}
		return "", err
	}
	if row.HtmlContent == nil {
		return "", nil
	}
	return *row.HtmlContent, nil
}
