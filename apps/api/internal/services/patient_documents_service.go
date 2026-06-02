package services

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/plenya/api/internal/models"
)

// PatientDocumentsService — staff faz upload, paciente lê/baixa.
type PatientDocumentsService struct {
	db          *gorm.DB
	uploadsRoot string
}

func NewPatientDocumentsService(db *gorm.DB, uploadsRoot string) *PatientDocumentsService {
	return &PatientDocumentsService{db: db, uploadsRoot: uploadsRoot}
}

// allowedContentType: PDF, jpeg, png. Sem .doc/.exe/etc.
var allowedDocContentTypes = map[string]bool{
	"application/pdf": true,
	"image/jpeg":      true,
	"image/png":       true,
}

const maxDocSize = 20 * 1024 * 1024 // 20MB

// CreateInput é o payload do CreateDocument (uso staff).
type CreateDocumentInput struct {
	PatientID   uuid.UUID
	UploadedBy  uuid.UUID
	Type        models.PatientDocumentType
	Title       string
	Description string
	IssuedAt    *time.Time
	File        *multipart.FileHeader
}

func (s *PatientDocumentsService) Create(in CreateDocumentInput) (*models.PatientDocument, error) {
	if in.File == nil {
		return nil, errors.New("arquivo obrigatório")
	}
	if in.File.Size > maxDocSize {
		return nil, errors.New("arquivo > 20MB")
	}
	contentType := in.File.Header.Get("Content-Type")
	if !allowedDocContentTypes[contentType] {
		return nil, fmt.Errorf("tipo %q não permitido (PDF, JPG ou PNG)", contentType)
	}
	// M7 — magic bytes. PDF clínico precisa ter assinatura real (%PDF-),
	// imagem precisa ser JPEG/PNG real. Sem leniência aqui (clínico).
	detected, err := ValidateUploadMagicBytes(in.File, AllowedMimeSetClinical, nil)
	if err != nil {
		return nil, fmt.Errorf("conteúdo inválido: %w", err)
	}
	contentType = detected
	if strings.TrimSpace(in.Title) == "" {
		return nil, errors.New("título obrigatório")
	}

	// Sub-pasta por paciente
	dir := filepath.Join(s.uploadsRoot, "patient-docs", in.PatientID.String())
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, err
	}

	// Nome único: <docID>-<sanitizedFilename>
	docID := uuid.Must(uuid.NewV7())
	safeName := sanitizeDocFilename(in.File.Filename)
	storedName := docID.String() + "-" + safeName
	fullPath := filepath.Join(dir, storedName)

	src, err := in.File.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	dst, err := os.Create(fullPath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()
	if _, err := io.Copy(dst, src); err != nil {
		return nil, err
	}

	relPath := filepath.Join("patient-docs", in.PatientID.String(), storedName)

	issued := time.Now().UTC()
	if in.IssuedAt != nil {
		issued = *in.IssuedAt
	}

	uploadedBy := in.UploadedBy
	doc := &models.PatientDocument{
		ID:          docID,
		PatientID:   in.PatientID,
		UploadedBy:  &uploadedBy,
		Source:      models.DocumentSourceStaffUpload,
		Type:        in.Type,
		Title:       in.Title,
		Description: in.Description,
		FilePath:    relPath,
		FileName:    in.File.Filename,
		ContentType: contentType,
		SizeBytes:   in.File.Size,
		IssuedAt:    issued,
	}
	if err := s.db.Create(doc).Error; err != nil {
		_ = os.Remove(fullPath)
		return nil, err
	}
	return doc, nil
}

// CreateFromBytesInput é o payload pra criar documento de prontuário a partir de
// bytes já em memória (ex: mídia recebida por WhatsApp inbound). Diferente do
// Create (staff/multipart), aqui UploadedBy é opcional (inbound não tem usuário).
type CreateFromBytesInput struct {
	PatientID         uuid.UUID
	Bytes             []byte
	MIME              string // MIME informado pela origem (ex: Meta)
	Filename          string
	Title             string
	Type              models.PatientDocumentType
	Source            models.PatientDocumentSource
	OriginWAMessageID *string    // idempotência do webhook
	UploadedBy        *uuid.UUID // staff responsável (ex: documento gerado pelo médico); nil em inbound
}

// CreateFromBytes salva um arquivo (bytes) como documento de prontuário.
// Idempotente por OriginWAMessageID: se já existe doc com o mesmo wa_message_id,
// retorna o existente sem regravar (Meta reentrega webhooks).
func (s *PatientDocumentsService) CreateFromBytes(in CreateFromBytesInput) (*models.PatientDocument, error) {
	if len(in.Bytes) == 0 {
		return nil, errors.New("arquivo vazio")
	}
	if int64(len(in.Bytes)) > maxDocSize {
		return nil, errors.New("arquivo > 20MB")
	}

	// Idempotência: webhook reentregue não duplica documento.
	if in.OriginWAMessageID != nil && *in.OriginWAMessageID != "" {
		var existing models.PatientDocument
		err := s.db.Where("origin_wa_message_id = ? AND patient_id = ?", *in.OriginWAMessageID, in.PatientID).First(&existing).Error
		if err == nil {
			return &existing, nil
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	// Magic bytes — só PDF/JPG/PNG entram no prontuário (clínico). Outros tipos
	// (docx, áudio) não devem chegar aqui: o caller roteia pro bucket da conversa.
	detected := DetectBytesMime(in.Bytes)
	if !allowedDocContentTypes[detected.ContentType] {
		return nil, fmt.Errorf("conteúdo %q não permitido no prontuário (PDF, JPG ou PNG)", detected.ContentType)
	}
	contentType := detected.ContentType

	docType := in.Type
	if docType == "" {
		docType = models.DocumentTypeOther
	}
	title := strings.TrimSpace(in.Title)
	if title == "" {
		title = "Arquivo recebido por WhatsApp"
	}

	dir := filepath.Join(s.uploadsRoot, "patient-docs", in.PatientID.String())
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, err
	}

	docID := uuid.Must(uuid.NewV7())
	safeName := sanitizeDocFilename(in.Filename)
	storedName := docID.String() + "-" + safeName
	fullPath := filepath.Join(dir, storedName)

	if err := os.WriteFile(fullPath, in.Bytes, 0o644); err != nil {
		return nil, err
	}
	relPath := filepath.Join("patient-docs", in.PatientID.String(), storedName)

	source := in.Source
	if source == "" {
		source = models.DocumentSourceWhatsApp
	}
	doc := &models.PatientDocument{
		ID:                docID,
		PatientID:         in.PatientID,
		UploadedBy:        in.UploadedBy, // staff responsável (doc gerado) ou nil (inbound)
		Source:            source,
		OriginWAMessageID: in.OriginWAMessageID,
		Type:              docType,
		Title:             title,
		FilePath:          relPath,
		FileName:          in.Filename,
		ContentType:       contentType,
		SizeBytes:         int64(len(in.Bytes)),
		IssuedAt:          time.Now().UTC(),
	}
	if err := s.db.Create(doc).Error; err != nil {
		_ = os.Remove(fullPath)
		// Corrida de idempotência: outra entrega do mesmo wa_message_id inseriu primeiro.
		// O uniqueIndex barra a duplicata — devolvemos o documento existente.
		if isUniqueViolation(err) && in.OriginWAMessageID != nil && *in.OriginWAMessageID != "" {
			var existing models.PatientDocument
			if e := s.db.Where("origin_wa_message_id = ? AND patient_id = ?", *in.OriginWAMessageID, in.PatientID).First(&existing).Error; e == nil {
				return &existing, nil
			}
		}
		return nil, err
	}
	return doc, nil
}

// List devolve docs do paciente (lista usada por staff E pelo paciente — mesmo escopo).
func (s *PatientDocumentsService) List(patientID uuid.UUID) ([]models.PatientDocument, error) {
	var rows []models.PatientDocument
	err := s.db.Where("patient_id = ?", patientID).
		Preload("UploadedByUser").
		Order("issued_at DESC").
		Find(&rows).Error
	return rows, err
}

// GetForDownload retorna doc + caminho absoluto após validar ownership.
func (s *PatientDocumentsService) GetForDownload(patientID, docID uuid.UUID) (*models.PatientDocument, string, error) {
	var doc models.PatientDocument
	err := s.db.Where("id = ? AND patient_id = ?", docID, patientID).First(&doc).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", errors.New("documento não encontrado")
		}
		return nil, "", err
	}
	full := filepath.Join(s.uploadsRoot, doc.FilePath)
	// Defesa anti path-traversal: garante que full está dentro de uploadsRoot
	abs, _ := filepath.Abs(full)
	rootAbs, _ := filepath.Abs(s.uploadsRoot)
	// Separador no prefixo evita bypass por irmão (/app/uploads-evil).
	if abs != rootAbs && !strings.HasPrefix(abs, rootAbs+string(os.PathSeparator)) {
		return nil, "", errors.New("caminho inválido")
	}
	if _, err := os.Stat(full); err != nil {
		return nil, "", errors.New("arquivo não encontrado no storage")
	}
	return &doc, full, nil
}

// Delete remove doc + arquivo (uso staff).
func (s *PatientDocumentsService) Delete(docID uuid.UUID) error {
	var doc models.PatientDocument
	if err := s.db.First(&doc, "id = ?", docID).Error; err != nil {
		return err
	}
	full := filepath.Join(s.uploadsRoot, doc.FilePath)
	if err := s.db.Delete(&doc).Error; err != nil {
		return err
	}
	_ = os.Remove(full) // best-effort
	return nil
}

func sanitizeDocFilename(name string) string {
	name = filepath.Base(name)
	name = strings.ReplaceAll(name, " ", "_")
	// Remove caracteres perigosos
	var b strings.Builder
	for _, r := range name {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') ||
			r == '_' || r == '-' || r == '.' {
			b.WriteRune(r)
		}
	}
	out := b.String()
	if out == "" {
		out = "documento"
	}
	if len(out) > 100 {
		out = out[:100]
	}
	return out
}
