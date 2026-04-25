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

	doc := &models.PatientDocument{
		ID:          docID,
		PatientID:   in.PatientID,
		UploadedBy:  in.UploadedBy,
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
	if !strings.HasPrefix(abs, rootAbs) {
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
