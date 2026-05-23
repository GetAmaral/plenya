package services

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	openai "github.com/sashabaranov/go-openai"

	"github.com/plenya/api/internal/config"
)

// TranscriptionService transcreve áudio (nota de voz do WhatsApp) via OpenAI Whisper.
// Sem API key configurada vira no-op (mesmo padrão de fallback dos outros serviços).
type TranscriptionService struct {
	client *openai.Client
}

func NewTranscriptionService(cfg *config.Config) *TranscriptionService {
	s := &TranscriptionService{}
	if cfg.OpenAI.APIKey != "" {
		c := openai.DefaultConfig(cfg.OpenAI.APIKey)
		if cfg.OpenAI.APIURL != "" && cfg.OpenAI.APIURL != "https://api.openai.com/v1" {
			c.BaseURL = cfg.OpenAI.APIURL
		}
		s.client = openai.NewClientWithConfig(c)
	}
	return s
}

// Enabled indica se a transcrição está disponível (API key presente).
func (s *TranscriptionService) Enabled() bool { return s.client != nil }

// Transcribe transcreve bytes de áudio em PT-BR. Retorna ("", nil) se desabilitado.
func (s *TranscriptionService) Transcribe(ctx context.Context, data []byte, mime string) (string, error) {
	if s.client == nil {
		return "", nil
	}
	resp, err := s.client.CreateTranscription(ctx, openai.AudioRequest{
		Model:    openai.Whisper1,
		Reader:   bytes.NewReader(data),
		FilePath: "audio" + audioExt(mime), // nome só pra Whisper inferir o formato
		Language: "pt",
	})
	if err != nil {
		return "", fmt.Errorf("transcription: whisper: %w", err)
	}
	return strings.TrimSpace(resp.Text), nil
}

// audioExt deriva a extensão do MIME (Whisper aceita ogg/mp3/m4a/wav/webm…).
func audioExt(mime string) string {
	switch {
	case strings.Contains(mime, "ogg"):
		return ".ogg"
	case strings.Contains(mime, "mpeg"):
		return ".mp3"
	case strings.Contains(mime, "mp4"), strings.Contains(mime, "aac"):
		return ".m4a"
	case strings.Contains(mime, "wav"):
		return ".wav"
	case strings.Contains(mime, "webm"):
		return ".webm"
	default:
		return ".ogg"
	}
}
