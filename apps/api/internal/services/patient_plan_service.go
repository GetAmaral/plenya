package services

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/dto"
	"github.com/plenya/api/internal/models"
	"github.com/plenya/api/internal/pdfdoc"
	"github.com/plenya/api/internal/utils"
)

var (
	ErrPatientPlanNotFound = errors.New("plano do paciente não encontrado")
	ErrPatientPlanEmpty    = errors.New("plano sem slides — não há o que publicar")
)

// ErrPatientPlanOverflow — algum slide transborda a moldura de 1920×1080.
//
// Não é aviso, é bloqueio de publicação. O slide tem altura fixa e `overflow:hidden`: conteúdo
// demais não empurra nada nem gera erro, simplesmente SOME do PDF. Publicar assim entrega ao
// paciente um documento com pedaço faltando, e ninguém percebe.
type ErrPatientPlanOverflow struct {
	Slides []pdfdoc.DeckOverflow
}

func (e *ErrPatientPlanOverflow) Error() string {
	parts := make([]string, 0, len(e.Slides))
	for _, s := range e.Slides {
		parts = append(parts, fmt.Sprintf("slide %02d (%s) passa %.0fpx", s.Slide, s.Title, maxF(s.Right, s.Bottom)))
	}
	return "conteúdo não cabe no slide: " + strings.Join(parts, "; ")
}

func maxF(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// PatientPlanService — o plano de devolutiva: rascunho, edição e publicação.
type PatientPlanService struct {
	db        *gorm.DB
	documents *PatientDocumentsService
}

func NewPatientPlanService(db *gorm.DB, documents *PatientDocumentsService) *PatientPlanService {
	return &PatientPlanService{db: db, documents: documents}
}

func (s *PatientPlanService) ListByPatient(patientID uuid.UUID) ([]dto.PatientPlanResponse, error) {
	var rows []models.PatientPlan
	if err := s.db.Where("patient_id = ?", patientID).
		Order("created_at DESC").Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]dto.PatientPlanResponse, len(rows))
	for i := range rows {
		out[i] = *toPatientPlanDTO(&rows[i])
	}
	return out, nil
}

func (s *PatientPlanService) Get(planID, patientID uuid.UUID) (*dto.PatientPlanResponse, error) {
	plan, err := s.load(planID, patientID)
	if err != nil {
		return nil, err
	}
	return toPatientPlanDTO(plan), nil
}

func (s *PatientPlanService) Create(patientID, authorID uuid.UUID, req *dto.SavePatientPlanRequest) (*dto.PatientPlanResponse, error) {
	title := strings.TrimSpace(req.Title)
	if title == "" {
		title = "Seus exames"
	}
	plan := models.PatientPlan{
		PatientID:    patientID,
		Title:        title,
		Status:       models.PatientPlanDraft,
		Version:      1,
		Content:      req.Content,
		AuthorUserID: authorID,
	}
	if plan.Content == nil {
		plan.Content = []pdfdoc.DeckSlide{}
	}
	if req.SourceSnapshotID != nil && *req.SourceSnapshotID != "" {
		id, err := uuid.Parse(*req.SourceSnapshotID)
		if err != nil {
			return nil, errors.New("sourceSnapshotId inválido")
		}
		plan.SourceSnapshotID = &id
	}
	if err := s.db.Create(&plan).Error; err != nil {
		return nil, err
	}
	return toPatientPlanDTO(&plan), nil
}

// Update reescreve o conteúdo do plano. Publicar de novo é o que fecha uma versão nova.
func (s *PatientPlanService) Update(planID, patientID uuid.UUID, req *dto.SavePatientPlanRequest) (*dto.PatientPlanResponse, error) {
	plan, err := s.load(planID, patientID)
	if err != nil {
		return nil, err
	}
	if t := strings.TrimSpace(req.Title); t != "" {
		plan.Title = t
	}
	if req.Content != nil {
		plan.Content = req.Content
	}
	// Editar um plano publicado volta ele para rascunho, e os ponteiros de documento saem junto: o
	// conteúdo mudou, então aqueles PDFs não representam mais este plano. Os arquivos continuam no
	// portal (o paciente já os recebeu) e `PublishedAt` fica como registro de que houve uma
	// publicação — o que a tela mostra é "rascunho, vN publicada em <data>".
	plan.Status = models.PatientPlanDraft
	plan.Document16x9ID = nil
	plan.DocumentA4ID = nil
	if err := s.db.Save(plan).Error; err != nil {
		return nil, err
	}
	return toPatientPlanDTO(plan), nil
}

func (s *PatientPlanService) Delete(planID, patientID uuid.UUID) error {
	res := s.db.Where("id = ? AND patient_id = ?", planID, patientID).Delete(&models.PatientPlan{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return ErrPatientPlanNotFound
	}
	return nil
}

// Preview devolve o HTML do deck sem publicar nada — é o que a tela de montagem mostra.
func (s *PatientPlanService) Preview(planID, patientID uuid.UUID) (string, error) {
	plan, err := s.load(planID, patientID)
	if err != nil {
		return "", err
	}
	return pdfdoc.DeckHTML(s.deck(plan))
}

// CheckOverflow mede se algum slide transborda, sem publicar.
func (s *PatientPlanService) CheckOverflow(planID, patientID uuid.UUID) ([]pdfdoc.DeckOverflow, error) {
	plan, err := s.load(planID, patientID)
	if err != nil {
		return nil, err
	}
	if len(plan.Content) == 0 {
		return nil, ErrPatientPlanEmpty
	}
	return pdfdoc.CheckDeckOverflow(s.deck(plan))
}

// Publish gera os dois PDFs e publica no portal do paciente.
//
// SÍNCRONO de propósito, apesar de o Chromium ser serializado por um mutex global: medido, um deck
// de 8 slides sai em ~1,0 s no 16:9 e ~0,2 s no A4, então uma publicação segura a fila da receita
// por poucos segundos. Vale reavaliar se aparecer deck muito maior ou uso concorrente de verdade;
// hoje uma fila de jobs só para isso seria peso sem retorno.
func (s *PatientPlanService) Publish(planID, patientID, publisherID uuid.UUID) (*dto.PatientPlanResponse, error) {
	plan, err := s.load(planID, patientID)
	if err != nil {
		return nil, err
	}
	if len(plan.Content) == 0 {
		return nil, ErrPatientPlanEmpty
	}

	var patient models.Patient
	if err := s.db.First(&patient, patientID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPatientNotFound
		}
		return nil, err
	}

	deck := s.deck(plan)

	// Barreira: o que não cabe no slide some do PDF em silêncio. Melhor recusar e dizer onde.
	over, err := pdfdoc.CheckDeckOverflow(deck)
	if err != nil {
		return nil, fmt.Errorf("erro ao medir os slides: %w", err)
	}
	if len(over) > 0 {
		return nil, &ErrPatientPlanOverflow{Slides: over}
	}

	// Publicar de novo é uma versão nova: o documento antigo continua no portal, e o `source_ref`
	// da versão nova é diferente, então nada é sobrescrito nem deduplicado por engano.
	version := plan.Version
	if plan.PublishedAt != nil {
		version++
	}
	now := time.Now()

	type output struct {
		paper    pdfdoc.DeckPaper
		kind     string
		suffix   string
		title    string
		assignTo **uuid.UUID
	}
	var doc169, docA4 *uuid.UUID
	outputs := []output{
		{pdfdoc.DeckPaper169, "Plano", "16x9", plan.Title, &doc169},
		{pdfdoc.DeckPaperA4, "PlanoImpressao", "a4", plan.Title + " (impressão)", &docA4},
	}

	// Renderiza os DOIS antes de publicar qualquer um. Publicando dentro do mesmo laço, uma falha
	// no A4 deixava o 16:9 já visível no portal enquanto o médico via "falha ao publicar" e o plano
	// continuava rascunho, sem registro do documento órfão.
	rendered := make([][]byte, len(outputs))
	for i, o := range outputs {
		bytesPDF, rErr := pdfdoc.RenderDeck(deck, o.paper)
		if rErr != nil {
			return nil, fmt.Errorf("erro ao gerar o PDF %s do plano: %w", o.paper, rErr)
		}
		rendered[i] = bytesPDF
	}

	for i, o := range outputs {
		bytesPDF := rendered[i]
		// Nome legível: o PDF sai do EMR e vai parar na pasta de downloads do paciente.
		filename := utils.DocumentFileName(patient.Name, o.kind, now, plan.ID)
		sourceRef := fmt.Sprintf("patient_plan:%s:v%d:%s", plan.ID, version, o.suffix)
		doc, dErr := s.documents.CreateFromBytes(CreateFromBytesInput{
			PatientID:  patientID,
			Bytes:      bytesPDF,
			Filename:   filename,
			Title:      o.title,
			Type:       models.DocumentTypeReport,
			Source:     models.DocumentSourceStaffUpload,
			UploadedBy: &publisherID,
			SourceRef:  &sourceRef,
		})
		if dErr != nil {
			return nil, fmt.Errorf("erro ao publicar o plano no portal: %w", dErr)
		}
		id := doc.ID
		*o.assignTo = &id
	}

	plan.Status = models.PatientPlanPublished
	plan.Version = version
	plan.PublishedAt = &now
	plan.Document16x9ID = doc169
	plan.DocumentA4ID = docA4
	if err := s.db.Save(plan).Error; err != nil {
		return nil, err
	}
	return toPatientPlanDTO(plan), nil
}

func (s *PatientPlanService) deck(plan *models.PatientPlan) pdfdoc.Deck {
	return pdfdoc.Deck{Title: plan.Title, Slides: plan.Content}
}

func (s *PatientPlanService) load(planID, patientID uuid.UUID) (*models.PatientPlan, error) {
	var plan models.PatientPlan
	// Escopado pelo paciente SEMPRE: nenhuma tela de prontuário lê por id solto.
	err := s.db.Where("id = ? AND patient_id = ?", planID, patientID).First(&plan).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrPatientPlanNotFound
	}
	if err != nil {
		return nil, err
	}
	return &plan, nil
}

func toPatientPlanDTO(p *models.PatientPlan) *dto.PatientPlanResponse {
	r := &dto.PatientPlanResponse{
		ID:           p.ID.String(),
		PatientID:    p.PatientID.String(),
		Title:        p.Title,
		Status:       string(p.Status),
		Version:      p.Version,
		Content:      p.Content,
		AuthorUserID: p.AuthorUserID.String(),
		CreatedAt:    p.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    p.UpdatedAt.Format(time.RFC3339),
	}
	if r.Content == nil {
		r.Content = []pdfdoc.DeckSlide{}
	}
	if p.SourceSnapshotID != nil {
		v := p.SourceSnapshotID.String()
		r.SourceSnapshotID = &v
	}
	if p.PublishedAt != nil {
		v := p.PublishedAt.In(saoPaulo()).Format(time.RFC3339)
		r.PublishedAt = &v
	}
	if p.Document16x9ID != nil {
		v := p.Document16x9ID.String()
		r.Document16x9ID = &v
	}
	if p.DocumentA4ID != nil {
		v := p.DocumentA4ID.String()
		r.DocumentA4ID = &v
	}
	return r
}
