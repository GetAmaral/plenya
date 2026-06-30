package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// ClinicalDocumentService — "Enviar por WhatsApp" de documentos clínicos (pedido de exames,
// documento emitido, receita, doc de prontuário) ao paciente, sem o round-trip de baixar+reanexar.
//
// Tudo resolve para um PatientDocument (tem arquivo em disco + serve share-link). Emitido/receita
// já materializam um; pedido de exames é materializado on-demand (idempotente via source_ref).
type ClinicalDocumentService struct {
	db            *gorm.DB
	patientDocs   *PatientDocumentsService
	conversations *ConversationService
	uploadsRoot   string
}

func NewClinicalDocumentService(db *gorm.DB, patientDocs *PatientDocumentsService, conversations *ConversationService, uploadsRoot string) *ClinicalDocumentService {
	return &ClinicalDocumentService{db: db, patientDocs: patientDocs, conversations: conversations, uploadsRoot: uploadsRoot}
}

// ClinicalDocType — origem do documento clínico a enviar.
type ClinicalDocType string

const (
	ClinicalDocLabRequest   ClinicalDocType = "lab_request"
	ClinicalDocIssued       ClinicalDocType = "issued_document"
	ClinicalDocPrescription ClinicalDocType = "prescription"
	ClinicalDocPatientDoc   ClinicalDocType = "patient_document"
)

// templateDocumentoDisponivel — template aprovado na Meta p/ mandar o link do documento.
// Params: {{1}} primeiro nome · {{2}} tipo do doc · {{3}} link seguro.
const templateDocumentoDisponivel = "documento_disponivel"

var (
	ErrClinicalDocUnknownType = errors.New("clinical-doc: tipo de documento desconhecido")
	ErrClinicalDocNoFile      = errors.New("clinical-doc: documento ainda não tem PDF (gere/assine antes)")
)

// DocLabel — rótulo legível do tipo, usado no {{2}} do template.
func DocLabel(t ClinicalDocType) string {
	switch t {
	case ClinicalDocLabRequest:
		return "Pedido de exames"
	case ClinicalDocIssued:
		return "Documento"
	case ClinicalDocPrescription:
		return "Receita"
	default:
		return "Documento"
	}
}

// ResolveToPatientDocument resolve qualquer documento clínico para um PatientDocument do paciente,
// materializando uma cópia compartilhável quando a origem não é um patient_document (pedido de
// exames). Valida ownership pelo patientID.
func (s *ClinicalDocumentService) ResolveToPatientDocument(patientID uuid.UUID, docType ClinicalDocType, docID, actorID uuid.UUID) (*models.PatientDocument, error) {
	switch docType {
	case ClinicalDocPatientDoc:
		doc, _, err := s.patientDocs.GetForDownload(patientID, docID)
		return doc, err

	case ClinicalDocIssued:
		var d models.IssuedDocument
		if err := s.db.Where("id = ? AND patient_id = ?", docID, patientID).First(&d).Error; err != nil {
			return nil, errClinicalNotFound(err)
		}
		if d.PatientDocumentID == nil {
			return nil, ErrClinicalDocNoFile
		}
		doc, _, err := s.patientDocs.GetForDownload(patientID, *d.PatientDocumentID)
		return doc, err

	case ClinicalDocPrescription:
		var p models.Prescription
		if err := s.db.Where("id = ? AND patient_id = ?", docID, patientID).First(&p).Error; err != nil {
			return nil, errClinicalNotFound(err)
		}
		if p.PatientDocumentID == nil {
			return nil, ErrClinicalDocNoFile
		}
		doc, _, err := s.patientDocs.GetForDownload(patientID, *p.PatientDocumentID)
		return doc, err

	case ClinicalDocLabRequest:
		return s.materializeLabRequest(patientID, docID, actorID)

	default:
		return nil, ErrClinicalDocUnknownType
	}
}

// materializeLabRequest garante um PatientDocument a partir do PDF (assinado) do pedido de exames.
// Idempotente por source_ref — não duplica em envios seguintes.
func (s *ClinicalDocumentService) materializeLabRequest(patientID, labRequestID, actorID uuid.UUID) (*models.PatientDocument, error) {
	var lr models.LabRequest
	if err := s.db.Where("id = ? AND patient_id = ?", labRequestID, patientID).First(&lr).Error; err != nil {
		return nil, errClinicalNotFound(err)
	}
	if lr.PdfURL == nil || *lr.PdfURL == "" {
		return nil, ErrClinicalDocNoFile
	}
	rel := strings.TrimPrefix(*lr.PdfURL, "/uploads/")
	full := filepath.Join(s.uploadsRoot, rel)
	bytes, err := os.ReadFile(full)
	if err != nil {
		return nil, ErrClinicalDocNoFile
	}
	sourceRef := "lab_request:" + labRequestID.String()
	uploadedBy := actorID
	return s.patientDocs.CreateFromBytes(CreateFromBytesInput{
		PatientID:  patientID,
		Bytes:      bytes,
		MIME:       "application/pdf",
		Filename:   "Pedido de exames.pdf",
		Title:      "Pedido de exames",
		Type:       models.DocumentTypeReferral,
		Source:     models.DocumentSourceStaffUpload,
		UploadedBy: &uploadedBy,
		SourceRef:  &sourceRef,
	})
}

// SendFile manda o PDF como mídia inline no WhatsApp (session message). Reusa ConversationService:
// valida janela 24h (ErrConversationWindowClosed → o caller cai pro link) e persiste a activity.
func (s *ClinicalDocumentService) SendFile(ctx context.Context, patientID uuid.UUID, doc *models.PatientDocument, actorID uuid.UUID) (*models.LeadActivity, error) {
	return s.conversations.SendMessage(ctx, SendMessageInput{
		UserID:    actorID,
		OwnerType: "patient",
		OwnerID:   patientID,
		Channel:   models.LeadChannelWhatsApp,
		Attachments: []OutboundAttachment{{
			Path:     doc.FilePath,
			Filename: doc.FileName,
		}},
	})
}

// SendLink manda o link seguro do documento via template documento_disponivel (reabre conversa
// fora da janela 24h). url é o link público já montado (com BaseURL) pelo handler.
func (s *ClinicalDocumentService) SendLink(ctx context.Context, patientID uuid.UUID, label, url string, actorID uuid.UUID) (*models.LeadActivity, error) {
	firstName := s.patientFirstName(patientID)
	return s.conversations.SendWhatsAppTemplate(ctx, SendWhatsAppTemplateInput{
		UserID:       actorID,
		OwnerType:    "patient",
		OwnerID:      patientID,
		TemplateName: templateDocumentoDisponivel,
		Language:     "pt_BR",
		Params:       []string{firstName, label, url},
	})
}

func (s *ClinicalDocumentService) patientFirstName(patientID uuid.UUID) string {
	var p models.Patient
	if err := s.db.Select("name").First(&p, "id = ?", patientID).Error; err != nil {
		return ""
	}
	if fields := strings.Fields(p.Name); len(fields) > 0 {
		return fields[0]
	}
	return p.Name
}

// WhatsAppWindowState — estado da janela de 24h do paciente, p/ os cartões decidirem os botões.
type WhatsAppWindowState struct {
	HasPhone      bool       `json:"hasPhone"`
	WindowOpen    bool       `json:"windowOpen"`
	LastInboundAt *time.Time `json:"lastInboundAt,omitempty"`
}

// WhatsAppWindow calcula se o paciente respondeu nas últimas 24h (janela aberta p/ session message).
func (s *ClinicalDocumentService) WhatsAppWindow(patientID uuid.UUID) (*WhatsAppWindowState, error) {
	var p models.Patient
	if err := s.db.Select("id", "phone").First(&p, "id = ?", patientID).Error; err != nil {
		return nil, errClinicalNotFound(err)
	}
	st := &WhatsAppWindowState{HasPhone: p.Phone != nil && *p.Phone != ""}

	var last models.LeadActivity
	err := s.db.
		Where("patient_id = ? AND channel = ? AND type = ?", patientID, models.LeadChannelWhatsApp, models.LeadActivityMessageReceived).
		Order("created_at DESC").
		First(&last).Error
	if err == nil {
		t := last.CreatedAt
		st.LastInboundAt = &t
		st.WindowOpen = time.Since(t) < 24*time.Hour
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return st, nil
}

// ClinicalDocItem — item do picker "Anexar arquivo do EMR".
type ClinicalDocItem struct {
	DocType   ClinicalDocType `json:"docType"`
	DocID     uuid.UUID       `json:"docId"`
	Title     string          `json:"title"`
	Filename  string          `json:"filename"`
	CreatedAt time.Time       `json:"createdAt"`
	Signed    bool            `json:"signed"`
}

// ListForPatient unifica os documentos enviáveis do paciente (com PDF pronto) p/ o picker.
func (s *ClinicalDocumentService) ListForPatient(patientID uuid.UUID) ([]ClinicalDocItem, error) {
	out := []ClinicalDocItem{}

	// Pedidos de exames com PDF gerado.
	var labs []models.LabRequest
	if err := s.db.Where("patient_id = ? AND pdf_url IS NOT NULL AND pdf_url <> ''", patientID).
		Order("created_at DESC").Find(&labs).Error; err != nil {
		return nil, err
	}
	for _, lr := range labs {
		out = append(out, ClinicalDocItem{
			DocType:   ClinicalDocLabRequest,
			DocID:     lr.ID,
			Title:     "Pedido de exames · " + lr.CreatedAt.Format("02/01/2006"),
			Filename:  "Pedido de exames.pdf",
			CreatedAt: lr.CreatedAt,
			Signed:    lr.SignedAt != nil,
		})
	}

	// Documentos emitidos publicados (assinados).
	var issued []models.IssuedDocument
	if err := s.db.Where("patient_id = ? AND patient_document_id IS NOT NULL", patientID).
		Order("created_at DESC").Find(&issued).Error; err != nil {
		return nil, err
	}
	for _, d := range issued {
		out = append(out, ClinicalDocItem{
			DocType:   ClinicalDocIssued,
			DocID:     d.ID,
			Title:     d.Title,
			Filename:  d.Title + ".pdf",
			CreatedAt: d.CreatedAt,
			Signed:    true,
		})
	}

	// Receitas publicadas.
	var presc []models.Prescription
	if err := s.db.Where("patient_id = ? AND patient_document_id IS NOT NULL", patientID).
		Order("created_at DESC").Find(&presc).Error; err != nil {
		return nil, err
	}
	for _, p := range presc {
		out = append(out, ClinicalDocItem{
			DocType:   ClinicalDocPrescription,
			DocID:     p.ID,
			Title:     p.GetTitle(),
			Filename:  "Receita.pdf",
			CreatedAt: p.CreatedAt,
			Signed:    true,
		})
	}

	// Documentos de prontuário (exclui os materializados via source_ref — já representados acima).
	var docs []models.PatientDocument
	if err := s.db.Where("patient_id = ? AND source_ref IS NULL", patientID).
		Order("created_at DESC").Find(&docs).Error; err != nil {
		return nil, err
	}
	for _, d := range docs {
		out = append(out, ClinicalDocItem{
			DocType:   ClinicalDocPatientDoc,
			DocID:     d.ID,
			Title:     d.Title,
			Filename:  d.FileName,
			CreatedAt: d.CreatedAt,
			Signed:    true,
		})
	}

	return out, nil
}

func errClinicalNotFound(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("clinical-doc: documento não encontrado")
	}
	return err
}
