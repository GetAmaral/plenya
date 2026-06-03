package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/plenya/api/internal/config"
)

func TestParseDiarizedVTT(t *testing.T) {
	vtt := strings.Join([]string{
		"WEBVTT",
		"",
		"1",
		"00:00:01.000 --> 00:00:04.000",
		"<v Speaker 0>Bom dia, como você está se sentindo hoje?",
		"",
		"2",
		"00:00:04.500 --> 00:00:07.000",
		"<v Speaker 1>Bom dia, doutor. Estou um pouco cansado.",
		"",
		"3",
		"00:00:07.500 --> 00:00:10.000",
		"<v Speaker 0>Entendo. Vamos conversar sobre isso.",
		"",
	}, "\n")

	got := parseDiarizedVTT([]byte(vtt))
	lines := strings.Split(got, "\n")
	if len(lines) != 3 {
		t.Fatalf("esperava 3 turnos, veio %d:\n%s", len(lines), got)
	}
	if !strings.Contains(lines[0], "Falante 1") || !strings.Contains(lines[0], "Bom dia, como você está") {
		t.Errorf("linha 0 inesperada: %q", lines[0])
	}
	if !strings.Contains(lines[1], "Falante 2") || !strings.Contains(lines[1], "Estou um pouco cansado") {
		t.Errorf("linha 1 inesperada: %q", lines[1])
	}
	// Speaker 0 reaparece → mesmo "Falante 1" (rótulos estáveis).
	if !strings.Contains(lines[2], "Falante 1") {
		t.Errorf("linha 2 deveria ser Falante 1: %q", lines[2])
	}
	if !strings.HasPrefix(lines[0], "[00:01]") {
		t.Errorf("linha 0 deveria começar com timestamp [00:01]: %q", lines[0])
	}
}

func TestParseDiarizedVTT_MergesConsecutiveSameSpeaker(t *testing.T) {
	vtt := strings.Join([]string{
		"WEBVTT",
		"",
		"00:00:01.000 --> 00:00:03.000",
		"<v Speaker 0>Primeira parte.",
		"",
		"00:00:03.000 --> 00:00:05.000",
		"<v Speaker 0>Segunda parte do mesmo falante.",
		"",
	}, "\n")
	got := parseDiarizedVTT([]byte(vtt))
	if strings.Count(got, "\n") != 0 {
		t.Errorf("turnos consecutivos do mesmo falante deveriam fundir numa linha:\n%s", got)
	}
	if !strings.Contains(got, "Primeira parte.") || !strings.Contains(got, "Segunda parte") {
		t.Errorf("texto fundido faltando: %q", got)
	}
}

func TestParseDiarizedVTT_NoDiarization(t *testing.T) {
	// Sem tags de falante → vira um único turno "Falante" sem número.
	vtt := "WEBVTT\n\n00:00:00.000 --> 00:00:02.000\nApenas um texto sem diarização.\n"
	got := parseDiarizedVTT([]byte(vtt))
	if !strings.Contains(got, "Apenas um texto") {
		t.Errorf("esperava o texto cru: %q", got)
	}
}

func TestVerifyWebhookSignature(t *testing.T) {
	rawSecret := []byte("plenya-daily-webhook-secret-0123456789")
	b64secret := base64.StdEncoding.EncodeToString(rawSecret)
	svc := NewDailyCoService(&config.Config{
		DailyCo: config.DailyCoConfig{WebhookSecret: b64secret},
	})

	body := []byte(`{"type":"recording.ready-to-download","payload":{"room_name":"plenya-abc"}}`)
	ts := strconv.FormatInt(time.Now().UTC().Unix(), 10)

	mac := hmac.New(sha256.New, rawSecret)
	mac.Write([]byte(ts + "." + string(body)))
	sig := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	if err := svc.VerifyWebhookSignature(sig, ts, body); err != nil {
		t.Fatalf("assinatura válida foi rejeitada: %v", err)
	}
	// Corpo adulterado → falha.
	if err := svc.VerifyWebhookSignature(sig, ts, append(body, '!')); err == nil {
		t.Error("corpo adulterado deveria falhar a verificação")
	}
	// Timestamp velho → replay rejeitado.
	oldTs := strconv.FormatInt(time.Now().UTC().Add(-10*time.Minute).Unix(), 10)
	macOld := hmac.New(sha256.New, rawSecret)
	macOld.Write([]byte(oldTs + "." + string(body)))
	sigOld := base64.StdEncoding.EncodeToString(macOld.Sum(nil))
	if err := svc.VerifyWebhookSignature(sigOld, oldTs, body); err == nil {
		t.Error("timestamp velho deveria ser rejeitado (replay)")
	}
	// Sem segredo → fail-closed.
	svcNoSecret := NewDailyCoService(&config.Config{})
	if err := svcNoSecret.VerifyWebhookSignature(sig, ts, body); err == nil {
		t.Error("sem segredo deveria falhar fechado")
	}
}
