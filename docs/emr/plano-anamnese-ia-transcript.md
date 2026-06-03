# Plano — Resumo/Anamnese por IA a partir do transcript (AI scribe estilo PLAUD)

> **Status:** EM IMPLANTAÇÃO em dev (2026-06-03). Decisão do usuário: construir **anamnese +
> SOAP** (templates selecionáveis). Estudo completo: workflow `estudo-resumo-anamnese-ia`
> (síntese arquivada). Reusa Claude (`ai_service.go`, padrão `tool_use` do `InterpretLabResult`)
> + o `TelemedRecording.transcript_text` já existente.

## Princípio
A nota gerada por IA é **rascunho, não ato assinável**. Grava num campo novo do `TelemedRecording`
(NÃO na `ClinicalNote`). O médico revisa lado-a-lado com o transcript, insere na nota SOAP e
assina. Igual ao fluxo do botão "Inserir no Subjetivo". CFM/LGPD: human-in-the-loop obrigatório,
marcação "gerado por IA (modelo, data), revisado pelo médico".

## Backend
- **config:** `Claude.NoteModel` (env `CLAUDE_NOTE_MODEL`, default `claude-sonnet-4-6`; geração
  pede mais qualidade que o haiku default). Configurável; se indisponível, cair pro `s.model`.
- **`ai_service.go`:** `GenerateClinicalNoteFromTranscript(transcript, format, model) (string,error)`
  espelhando `InterpretLabResult` (tool_use + tool_choice forçado). Schema por formato
  (anamnese: 11 seções; soap: 4) + `itens_ambiguos_para_revisao_medica` + `papeis`. Prompt com
  **regras anti-alucinação** (só o que está no transcript; teleconsulta = SEM exame físico →
  Objetivo default "não avaliado nesta teleconsulta"; nunca inventar valor/dose; "refere/relata"
  pro que o paciente disse; negações preservadas). temperatura baixa.
- **Model `TelemedRecording` (migration 00013):** `generated_note_json text`,
  `generated_note_format varchar(16)`, `generated_note_status varchar(16) default 'none'`
  (none|generating|done|failed), `generated_note_model varchar(64)`, `generated_note_at timestamptz`,
  `generated_note_error text`.
- **`telemed_recording_service.go`:** `GenerateNote(ctx, appointmentID, format)` — exige
  transcript_text; chama AIService; grava json/format/model/at/status; parseia o JSON bruto em
  **seções ordenadas** (chave+titulo+texto+soapTarget) pro DTO. Defs de seção (ordem/título/alvo
  SOAP) no Go. `GetByAppointment`/`toDTO` passam a incluir a nota gerada se houver.
- **DTO:** `GeneratedNote{format, sections[], itensAmbiguos[], papeis}` +
  `GeneratedNoteSection{chave,titulo,texto,soapTarget}`. Incluído no TelemedRecordingResponse.
- **Handler/Rota:** `POST /appointments/:id/telemed-recording/generate-note` (RequireClinician),
  body `{format: "anamnese"|"soap"}`. Síncrono (geração ~5-15s; spinner no front).
- **Wire:** injetar `aiService` no `NewTelemedRecordingService`.

### Mapeamento seção → campo SOAP (pro "Inserir na nota")
- Anamnese: queixa/hda/antecedentes/medicações/alergias/hábitos/revisão → **subjective**;
  objetivo_relatado → **objective**; avaliacao → **assessment**; plano → **plan**.
- SOAP: subjetivo/objetivo/avaliacao/plano → 1:1.

## Frontend (`apps/web`)
- `calendar-api.ts`: estende `TelemedRecording` com `generatedNote*`; `useGenerateTelemedNote(id)`.
- `consultation-workspace.tsx`: botão **"Gerar com IA"** + seletor de formato (Anamnese | SOAP);
  painel com o rascunho (seções), banner "gerado por IA — revise antes de assinar", lista de
  `itens_ambiguos` em destaque, e **"Inserir na nota"** (agrupa seções por soapTarget e anexa aos
  campos). Só aparece com type=telemedicine + transcrição pronta + nota não-assinada.

## LGPD/CFM (bloqueia só go-live em prod, não o dev)
- DPA Anthropic (cláusulas-padrão ANPD + ZDR + sem-treino).
- Atualizar termo de telemedicina: menção a "IA de apoio" + direito de recusa (CFM 2.454/2026,
  vigora 26/08/2026). Marcação "gerado por IA" + guarda rascunho+final (auditoria).
- Base LGPD = tutela da saúde; minimização (mandar só o transcript, sem CPF/identificadores).

## Verificação
- build/vet/gofmt; migration 00013 up/down; teste unitário do builder de schema/seções.
- **Teste real em dev** (dev TEM CLAUDE_API_KEY): gerar a partir de um transcript fictício →
  conferir seções, "não avaliado" no Objetivo, itens ambíguos, e o "Inserir na nota".
- Playwright: botão gera, painel renderiza, inserir preenche os campos SOAP, 0 erro.

## Não-fazer agora
- Não popular AnamnesisItem codificados por ScoreItem (risco numérico) — adiar.
- Não auto-inserir/auto-assinar. Não resumo livre tipo carta (pior anti-alucinação).
- Sem evidência-por-afirmação clicável no MVP (fica como evolução; manter itens_ambiguos).
