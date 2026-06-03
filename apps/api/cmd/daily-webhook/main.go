// Command daily-webhook registra (ou lista) o webhook do domínio Daily.co que
// entrega os eventos de gravação + transcrição da teleconsulta.
//
// Setup out-of-band (uma vez por ambiente):
//
//	docker compose exec -w /app api go run ./cmd/daily-webhook \
//	    -url https://api.plenyasaude.com.br/api/v1/webhooks/daily
//
// Ao criar, o Daily faz um POST de teste {"test":"test"} na URL (precisa 200) e
// retorna o segredo HMAC. Copie-o pra DAILY_CO_WEBHOOK_SECRET (Coolify) e redeploy.
//
//	-list   lista os webhooks já registrados no domínio (não cria)
package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/plenya/api/internal/config"
)

const dailyAPIBaseURL = "https://api.daily.co/v1"

func main() {
	url := flag.String("url", os.Getenv("DAILY_WEBHOOK_URL"), "URL pública do webhook (POST /api/v1/webhooks/daily)")
	list := flag.Bool("list", false, "lista webhooks existentes em vez de criar")
	flag.Parse()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}
	apiKey := cfg.DailyCo.APIKey
	if apiKey == "" {
		log.Fatal("DAILY_CO_API_KEY ausente")
	}

	client := &http.Client{Timeout: 20 * time.Second}
	doReq := func(method, path string, body any) ([]byte, int) {
		var reader io.Reader
		if body != nil {
			raw, _ := json.Marshal(body)
			reader = bytes.NewReader(raw)
		}
		req, _ := http.NewRequest(method, dailyAPIBaseURL+path, reader)
		req.Header.Set("Authorization", "Bearer "+apiKey)
		if body != nil {
			req.Header.Set("Content-Type", "application/json")
		}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("http %s %s: %v", method, path, err)
		}
		defer resp.Body.Close()
		out, _ := io.ReadAll(resp.Body)
		return out, resp.StatusCode
	}

	if *list {
		out, code := doReq(http.MethodGet, "/webhooks", nil)
		fmt.Printf("status=%d\n%s\n", code, string(out))
		return
	}

	if *url == "" {
		log.Fatal("informe -url https://api.plenyasaude.com.br/api/v1/webhooks/daily")
	}

	body := map[string]any{
		"url": *url,
		"eventTypes": []string{
			"recording.started",
			"recording.ready-to-download",
			"recording.error",
			"transcript.started",
			"transcript.ready-to-download",
			"transcript.error",
		},
		// circuit-breaker (default): para após 3 falhas consecutivas; nosso handler
		// responde 200 imediatamente, então não deve disparar.
		"retryType": "circuit-breaker",
	}
	out, code := doReq(http.MethodPost, "/webhooks", body)
	fmt.Printf("status=%d\n%s\n", code, string(out))
	if code >= 200 && code < 300 {
		var parsed struct {
			UUID string `json:"uuid"`
			HMAC string `json:"hmac"`
		}
		if json.Unmarshal(out, &parsed) == nil && parsed.HMAC != "" {
			fmt.Printf("\n✅ Webhook criado.\n   uuid: %s\n   👉 Copie o segredo HMAC pra DAILY_CO_WEBHOOK_SECRET (Coolify) e redeploy:\n   %s\n", parsed.UUID, parsed.HMAC)
		}
	}
}
