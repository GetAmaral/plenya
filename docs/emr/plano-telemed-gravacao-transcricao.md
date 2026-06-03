# Plano — Gravação + Transcrição da Teleconsulta (Daily.co + Deepgram nova-3)

> **Status:** IMPLEMENTADO + QA em dev (2026-06-02). Backend completo (model 00012,
> webhook HMAC, service, endpoints), frontend (bloco "Gravação e transcrição" no card de
> teleconsulta), testes unitários (VTT parser + HMAC) verdes, e2e do webhook verificado
> (recording.ready + transcript.ready → upsert + DTO, idempotente, degradação grácil sem
> DAILY_CO_API_KEY). **Falta só o setup out-of-band em prod** (ver no fim): habilitar
> gravação/transcrição no plano/domínio Daily, registrar o webhook (`cmd/daily-webhook`) e
> pôr `DAILY_CO_WEBHOOK_SECRET` no Coolify. Decisões do
> usuário (Getúlio):
> - **Gravação:** "referência + link sob demanda" — o MP4 fica no storage gerenciado do Daily;
>   guardamos só metadados e geramos link assinado de download quando preciso (auditado). NÃO
>   baixamos o vídeo.
> - **Disparo:** "auto ao entrar, só com consentimento" — o token do médico (owner) auto-inicia
>   gravação + transcrição ao entrar na sala, **somente** se `telemed_consent_at` já estiver
>   carimbado. Sem consentimento → não grava, não transcreve.
> - **Transcrição:** **Daily-nativo (Deepgram nova-3, realtime, `extra.diarize`)** — diarização
>   médico×paciente embutida, pt-BR, entregue como WebVTT via webhook. Sem download de MP4, sem
>   ffmpeg, sem conta nova. (Comparativo completo de 6 motores na pesquisa de 2026-06-02; OpenAI
>   `gpt-4o-transcribe-diarize` fica como fallback batch documentado, reusando `OPENAI_API_KEY`.)
>
> Base regulatória: o termo de consentimento de telemedicina (CFM 2.314/2022) já inclui a
> cláusula de gravação (#7). LGPD: dado processado fora do BR (EUA) → transferência internacional;
> postura = DPA/BAA + retenção mínima + cláusula no termo (já coletada) + registro no RIPD.

## Arquitetura

O Daily Prebuilt é embutido como **iframe** e a sala abre em **janela separada** (`window.open`
→ `win.location.href = joinUrl`). Não há SDK `daily-js`/`callObject` no client. Logo, o controle
de gravação/transcrição é **via propriedades de token + sala** (server-side), e a entrega dos
artefatos é **assíncrona via webhook do Daily** para o backend Go.

### Fluxo
1. Médico registra o consentimento de telemedicina (já existe).
2. Médico clica "Abrir sala" → backend cria/garante a sala Daily com `enable_recording: "cloud"`
   + transcrição-storage + settings pt-BR/diarize, e cunha o **token do owner** com
   `start_cloud_recording: true` + `auto_start_transcription: true` (somente se consentimento OK).
3. Ao entrar, o Daily auto-inicia gravação cloud + transcrição Deepgram nova-3 pt-BR diarizada.
   O Prebuilt mostra o indicador de gravação ao paciente (transparência LGPD).
4. Ao fim da chamada, o Daily dispara webhooks → nosso `POST /api/v1/webhooks/daily`:
   - `recording.ready-to-download` → grava metadados (recording_id, duração, s3_key) no
     `TelemedRecording`. NÃO baixa o MP4.
   - `transcript.ready-to-download` → `GET /transcript/:id/access-link` → baixa o **.vtt**
     (poucos KB) → parseia em diálogo rotulado (médico/paciente) → grava `transcript_text` +
     guarda o `.vtt` em uploads.
5. No prontuário/consulta: status "Gravação disponível" + "Baixar gravação" (link assinado sob
   demanda, auditado) e "Transcrição disponível" + visualização do diálogo (com botão copiar).

### Backend
- **Config** (`internal/config`): `DailyCoConfig.WebhookSecret` (env `DAILY_CO_WEBHOOK_SECRET`).
- **`daily_co_service.go`**:
  - `CreateRoom`: adicionar props de sala `enable_recording: "cloud"`,
    `enable_transcription_storage: true`, `auto_transcription_settings` (language `pt-BR`, model
    `nova-3-general`, `punctuate: true`, `extra.diarize: true`) — **shapes a confirmar na doc**.
  - `MeetingTokenParams` + `CreateMeetingToken`: novos campos `StartCloudRecording`,
    `AutoStartTranscription` → vão em `properties`.
  - Novos métodos REST: `GetRecordingAccessLink(recordingID)` (`GET /recordings/:id/access-link`),
    `GetTranscriptAccessLink(transcriptID)` (`GET /transcript/:id/access-link`),
    `DownloadTranscriptVTT(url)`.
- **Model `TelemedRecording`** (`telemed_recordings`, UUIDv7, `appointment_id *uuid` idx,
  `patient_id` idx, `daily_room_name` idx = chave de join dos webhooks):
  - Gravação: `recording_id *string`, `recording_status` (pending|recording|finished|error),
    `recording_started_at/ready_at *time`, `recording_duration_seconds *int`,
    `recording_s3_key *string`, `recording_error *string`.
  - Transcrição: `transcript_id *string`, `transcript_status` (none|in_progress|finished|failed),
    `transcript_ready_at *time`, `transcript_vtt_path *string`, `transcript_text *string`,
    `transcript_error *string`.
  - Hook `BeforeCreate` UUIDv7. Upsert por `appointment_id` (1 sessão = 1 linha; múltiplos
    start/stop sobrescrevem — aceitável no MVP).
- **Migration goose `00012_telemed_recordings.sql`** (aditiva, `IF NOT EXISTS`, com Down).
- **`telemed_recording_service.go`**: `UpsertFromRecordingWebhook`, `UpsertFromTranscriptWebhook`
  (baixa+parseia VTT), `GetByAppointment`, `RecordingDownloadLink` (auditado), `toDTO`.
- **`webhook` (novo handler `daily_webhook_handler.go`)**:
  - `POST /api/v1/webhooks/daily` (público, rate-limited como o do WhatsApp).
  - Bootstrap: se body == `{"test":"test"}` → 200 imediato (ping de criação do webhook).
  - Verificação HMAC: headers `X-Webhook-Signature` + `X-Webhook-Timestamp`; segredo
    **base64-decoded**; HMAC-SHA256 sobre `timestamp + "." + rawBody`; saída base64;
    `hmac.Equal`. Rejeita timestamp velho (replay, janela ~5 min). **Fail-closed** se o segredo
    não estiver configurado (loga claro).
  - Idempotente por `event.id` + estado terminal do artefato (não re-baixa transcript já finished).
  - Dispatch: `recording.*` e `transcript.*` → mapeia `payload.room_name` → appointment.
  - Ack 200 imediato, processa (download VTT) de forma resiliente.
- **Rotas** (`cmd/server/main.go`): `POST /api/v1/webhooks/daily` (público) +
  `GET /api/v1/appointments/:id/telemed-recording` (clinician) +
  `GET /api/v1/appointments/:id/telemed-recording/download` (clinician, auditado).
- **`GetTelemedJoinURL`** (token do médico): se `appt.TelemedConsentAt != nil`, set
  `StartCloudRecording: true` + `AutoStartTranscription: true`. Senão, false.
- **Setup helper `cmd/daily-webhook/main.go`** (out-of-band): `POST /v1/webhooks` apontando
  para `https://api.plenyasaude.com.br/api/v1/webhooks/daily` com os eventTypes; imprime o `hmac`
  retornado → operador põe `DAILY_CO_WEBHOOK_SECRET` no Coolify e redeploya.

### Frontend (`apps/web`)
- `lib/api/calendar-api.ts`: tipo `TelemedRecording` + `useTelemedRecording(appointmentId)`
  (poll enquanto pendente) + `useTelemedRecordingDownload(id)`.
- `appointments/[id]/page.tsx`: no card de Teleconsulta, novo bloco "Gravação e transcrição"
  (status pendente/pronto; botão "Baixar gravação"; visualização do diálogo transcrito + copiar).
  Quando consentimento OK, microcopy "Esta teleconsulta será gravada e transcrita."

### Verificação
- `go build ./...` exit 0; `migrate up` aplica 00012 e `down` reverte limpo; `pnpm generate`.
- HMAC: teste unitário do verify com vetor conhecido (timestamp+body+secret base64).
- QA dev: simular payloads `recording.ready-to-download` / `transcript.ready-to-download`
  (POST assinado) → confere upsert + parse VTT + DTO. Em prod, depende de habilitar transcrição
  no plano/domínio Daily + registrar o webhook (setup helper) — passos out-of-band documentados.

### Dependências out-of-band (antes do go-live em prod)
1. Habilitar **gravação + transcrição no plano/domínio Daily** (verificar billing; webhooks
   exigem cartão na conta Daily).
2. Rodar o setup helper → registrar webhook → pôr `DAILY_CO_WEBHOOK_SECRET` no Coolify.
3. DPA/BAA Daily+Deepgram; retenção mínima; cláusula de transferência internacional (já no termo);
   registro no RIPD.

### Não-fazer agora
- Não baixar/armazenar o MP4 (decisão "referência + link").
- Não fatiar (chunking) — Daily-nativo é streaming, sem arquivo; fatiar só quebraria diarização.
- Não auto-injetar a transcrição na nota SOAP assinada (mostrar como artefato + copiar manual).
- Não construir o fallback batch OpenAI agora (documentado como evolução).
