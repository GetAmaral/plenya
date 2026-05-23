package services

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

// M7 — validação de uploads por magic bytes.
//
// Problema: validar só Content-Type/extensão é trivial de burlar — atacante
// envia .pdf com payload PHP/PE32 dentro. A defesa é ler os primeiros 512 bytes
// e deixar http.DetectContentType inspecionar a assinatura real do conteúdo.
//
// Whitelist é por content-type "real" (detectado), não por extensão informada.
//
// Limitações:
//   - HEIC tem múltiplos brand boxes que http.DetectContentType pode não cobrir
//     (Go stdlib detecta image/heic mas com cobertura limitada). Aceitamos
//     fallback "application/octet-stream" pra mobile uploader desde que extensão
//     bata e tamanho seja razoável (caller decide). Aqui retornamos um boolean
//     leniencia.
//   - Stream consumido é restituído pro caller via reader recriado.

// DetectedMime — resultado da inspeção de magic bytes.
type DetectedMime struct {
	ContentType string // ex: "application/pdf", "image/jpeg"
	Lenient     bool   // true se octet-stream (caller decide se aceita por ext)
}

// AllowedMimeSet padrão de imagem+pdf (uploads gerais e documentos do paciente).
var AllowedMimeSetImagesPDF = map[string]bool{
	"application/pdf": true,
	"image/jpeg":      true,
	"image/png":       true,
	"image/webp":      true,
	"image/heic":      true,
	"image/heif":      true,
}

// AllowedMimeSet só PDF/JPG/PNG (documentos clínicos).
var AllowedMimeSetClinical = map[string]bool{
	"application/pdf": true,
	"image/jpeg":      true,
	"image/png":       true,
}

// DetectFileMime lê os primeiros 512 bytes do arquivo multipart e retorna o
// content-type detectado. NÃO valida contra whitelist — caller faz.
//
// Não consome o file: usa file.Open() + Close() — caller pode reabrir pra
// salvar o conteúdo completo.
func DetectFileMime(fh *multipart.FileHeader) (DetectedMime, error) {
	if fh == nil {
		return DetectedMime{}, errors.New("file header nil")
	}
	src, err := fh.Open()
	if err != nil {
		return DetectedMime{}, fmt.Errorf("open uploaded file: %w", err)
	}
	defer src.Close()

	buf := make([]byte, 512)
	n, err := io.ReadFull(src, buf)
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		return DetectedMime{}, fmt.Errorf("read magic bytes: %w", err)
	}
	ct := http.DetectContentType(buf[:n])
	// Strip charset suffix (e.g. "text/plain; charset=utf-8").
	if i := strings.Index(ct, ";"); i >= 0 {
		ct = strings.TrimSpace(ct[:i])
	}
	return DetectedMime{
		ContentType: ct,
		Lenient:     ct == "application/octet-stream",
	}, nil
}

// DetectBytesMime inspeciona magic bytes de um buffer já em memória (ex: mídia
// baixada da Meta) e retorna o content-type detectado. Mesmo critério do
// DetectFileMime, mas sem multipart.
func DetectBytesMime(data []byte) DetectedMime {
	n := len(data)
	if n > 512 {
		n = 512
	}
	ct := http.DetectContentType(data[:n])
	if i := strings.Index(ct, ";"); i >= 0 {
		ct = strings.TrimSpace(ct[:i])
	}
	return DetectedMime{ContentType: ct, Lenient: ct == "application/octet-stream"}
}

// ValidateUploadMagicBytes faz DetectFileMime + valida contra whitelist.
// Retorna o content-type detectado pra caller persistir (defense in depth).
//
// Se Lenient (octet-stream) e a extensão informada está numa whitelist auxiliar,
// aceita. Isso cobre HEIC/HEIF que stdlib às vezes detecta como octet-stream.
func ValidateUploadMagicBytes(
	fh *multipart.FileHeader,
	whitelist map[string]bool,
	lenientExtensions map[string]bool,
) (string, error) {
	det, err := DetectFileMime(fh)
	if err != nil {
		return "", err
	}
	if whitelist[det.ContentType] {
		return det.ContentType, nil
	}
	// Lenient fallback: extensão sane + lenientExtensions whitelist.
	if det.Lenient && lenientExtensions != nil {
		ext := strings.ToLower(extOf(fh.Filename))
		if lenientExtensions[ext] {
			return det.ContentType, nil
		}
	}
	return det.ContentType, fmt.Errorf("content-type detectado %q não permitido", det.ContentType)
}

func extOf(name string) string {
	for i := len(name) - 1; i >= 0; i-- {
		if name[i] == '.' {
			return name[i:]
		}
		if name[i] == '/' || name[i] == '\\' {
			break
		}
	}
	return ""
}
