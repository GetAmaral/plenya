package services

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"

	"github.com/plenya/api/internal/crypto"
)

// WhatsAppMediaService guarda mídia de conversa (WhatsApp) CIFRADA em repouso.
//
// Uso: mídia que NÃO vai pro prontuário do paciente — áudio/nota de voz, sticker,
// e qualquer arquivo de Lead (que não tem prontuário). Arquivos de paciente
// (foto/PDF) vão pro PatientDocumentsService (mídias do prontuário, não cifrado,
// porque o interpretador de exames lê do disco).
//
// Layout: <uploadsRoot>/whatsapp-media/<ownerType>/<ownerID>/<uuidv7>.<ext>
// Conteúdo: base64 do AES-256-GCM (crypto.Encrypt) — o arquivo em disco é cifra,
// não o binário cru.
type WhatsAppMediaService struct {
	uploadsRoot string
}

func NewWhatsAppMediaService(uploadsRoot string) *WhatsAppMediaService {
	return &WhatsAppMediaService{uploadsRoot: uploadsRoot}
}

// StoreInbound cifra os bytes e grava no bucket da conversa.
// ownerType deve ser "lead" ou "patient". Retorna o storage key (path relativo).
func (s *WhatsAppMediaService) StoreInbound(ownerType string, ownerID uuid.UUID, data []byte, mime string) (string, int64, error) {
	if len(data) == 0 {
		return "", 0, errors.New("whatsapp media: bytes vazios")
	}
	if ownerType != "lead" && ownerType != "patient" {
		return "", 0, fmt.Errorf("whatsapp media: ownerType inválido %q", ownerType)
	}

	dir := filepath.Join(s.uploadsRoot, "whatsapp-media", ownerType, ownerID.String())
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", 0, err
	}

	fileID := uuid.Must(uuid.NewV7())
	name := fileID.String() + extFromMime(mime)
	fullPath := filepath.Join(dir, name)

	cipher, err := crypto.Encrypt(data, crypto.GetDefaultKey())
	if err != nil {
		return "", 0, fmt.Errorf("whatsapp media: encrypt: %w", err)
	}
	if err := os.WriteFile(fullPath, []byte(cipher), 0o600); err != nil {
		return "", 0, err
	}

	relPath := filepath.Join("whatsapp-media", ownerType, ownerID.String(), name)
	return relPath, int64(len(data)), nil
}

// OpenDecrypted lê o arquivo cifrado pelo storage key e devolve o binário em claro.
func (s *WhatsAppMediaService) OpenDecrypted(storageKey string) ([]byte, error) {
	full := filepath.Join(s.uploadsRoot, storageKey)
	// Defesa anti path-traversal: full precisa estar dentro de uploadsRoot.
	abs, _ := filepath.Abs(full)
	rootAbs, _ := filepath.Abs(s.uploadsRoot)
	if !strings.HasPrefix(abs, rootAbs+string(os.PathSeparator)) {
		return nil, errors.New("whatsapp media: caminho inválido")
	}
	cipher, err := os.ReadFile(full)
	if err != nil {
		return nil, err
	}
	plain, err := crypto.Decrypt(string(cipher), crypto.GetDefaultKey())
	if err != nil {
		return nil, fmt.Errorf("whatsapp media: decrypt: %w", err)
	}
	return plain, nil
}

// extFromMime devolve uma extensão de arquivo a partir do MIME (best-effort).
func extFromMime(mime string) string {
	switch strings.ToLower(strings.TrimSpace(mime)) {
	case "image/jpeg":
		return ".jpg"
	case "image/png":
		return ".png"
	case "image/webp":
		return ".webp"
	case "application/pdf":
		return ".pdf"
	case "audio/ogg":
		return ".ogg"
	case "audio/mpeg":
		return ".mp3"
	case "audio/mp4", "audio/aac":
		return ".m4a"
	case "audio/amr":
		return ".amr"
	case "video/mp4":
		return ".mp4"
	case "image/webp+sticker", "image/sticker":
		return ".webp"
	default:
		return ".bin"
	}
}
