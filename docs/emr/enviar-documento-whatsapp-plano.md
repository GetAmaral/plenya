# Enviar documento clínico ao paciente por WhatsApp (EMR)

**Status:** aprovado (2026-06-29), em implementação.
**Origem:** fluxo atual é burro — pra mandar um pedido de exames/receita/documento ao paciente,
o usuário precisa **baixar o PDF** e **reanexar manualmente** no compositor de conversa. Os botões
de "enviar" nunca foram cabeados; o backend já tem quase tudo (link público + template
`documento_disponivel` aprovado; anexo de mídia consertado em `ee7b9de4`).

## Decisão de UX (aprovada pelo usuário)

1. **Botão "Enviar por WhatsApp" nos cartões de documento** (pedido de exames, emitido, receita),
   **smart por janela de 24h:**
   - Janela **aberta** (inbound do paciente < 24h): oferece **Enviar arquivo** (mídia inline) **e**
     **Enviar link** (template).
   - Janela **fechada** (≥ 24h ou sem inbound): só **Enviar link** (template reabre a conversa).
2. **No compositor de conversa**, no menu "Anexar", adicionar **"Anexar arquivo do EMR"** — picker
   dos documentos já gerados do paciente (pedidos, receitas, emitidos, patient_documents) sem
   download/reupload.

## Trade-off dos dois modos
- **Arquivo (mídia inline):** chega como PDF no chat do paciente; **só dentro da janela 24h** (session
  message — Meta). Caminho consertado em `ee7b9de4` (Content-Type real da parte).
- **Link (template `documento_disponivel`):** manda link seguro do PDF; **funciona sempre** (reabre
  conversa); LGPD-friendly (arquivo fica no servidor, token expira/revoga). Params do template:
  `{{1}}` primeiro nome · `{{2}}` tipo do doc · `{{3}}` link.

## Arquitetura — unificar em PatientDocument

Tudo resolve para um **`PatientDocument`** (tem `FilePath` em disco + serve share-link):
- Emitido / receita → já têm `PatientDocumentID`. Usar direto.
- Pedido de exames → **materializar** um `PatientDocument` a partir do PDF on-demand, idempotente.
- patient_document → é ele mesmo.

**Idempotência da materialização:** migration goose adiciona `patient_documents.source_ref`
(`varchar`, ex.: `lab_request:<uuid>`) com índice único parcial (NULLs distintos). Evita duplicar
o doc materializado a cada envio.

## Backend (apps/api)

1. **Migration goose** — `patient_documents.source_ref varchar(80)` + `uniqueIndex` parcial.
2. **`ClinicalDocumentService`** (novo ou método em PatientDocumentsService):
   - `ResolveToPatientDocument(patientID, docType, docId) (*PatientDocument, error)` —
     docType ∈ {`lab_request`,`issued_document`,`prescription`,`patient_document`}; materializa p/
     lab_request (idempotente via source_ref).
   - label legível por docType p/ `{{2}}` do template ("Pedido de exames", "Receita", etc.).
3. **Endpoint de envio** — `POST /patients/:id/clinical-documents/send-whatsapp`
   body `{ docType, docId, mode: "file"|"link" }`:
   - resolve → PatientDocument.
   - `mode=file`: lê bytes do FilePath → reusa o caminho de mídia (UploadMedia + SendMediaMessage)
     via ConversationService; **valida janela 24h** (422 se fechada → front cai pro link); persiste
     LeadActivity (aparece na conversa).
   - `mode=link`: mint share-link do doc + `SendWhatsAppTemplate("documento_disponivel",
     [primeiroNome, label, url])`. Persiste activity (corpo = texto renderizado, via fix `c842473f`).
4. **Estado da janela** — `GET /patients/:id/whatsapp-window` →
   `{ phone, optIn, windowOpen, lastInboundAt }`. Os cartões consultam pra decidir os botões.
5. **Listagem p/ o picker** — `GET /patients/:id/clinical-documents` → unifica
   lab_requests + prescriptions + issued_documents + patient_documents em
   `[{ docType, docId, title, filename, createdAt, signed }]`.

## Frontend (apps/web)

6. Hook/cliente em `lib/api/clinical-documents.ts`: `sendDocumentWhatsApp`, `useWhatsAppWindow`,
   `useClinicalDocuments`.
7. Botão **"Enviar por WhatsApp"** (menu smart) nos cartões:
   - `app/(authenticated)/lab-requests/page.tsx` (`LabRequestCard`, ~l.944, só quando há PDF).
   - `components/clinical/issued-documents-card.tsx` (~l.281, doc assinado).
   - `app/(authenticated)/prescriptions/page.tsx` (`PrescriptionActions`, ~l.337, assinada).
   - Componente compartilhado `components/clinical/send-document-whatsapp-button.tsx` (consulta
     janela, mostra arquivo+link ou só link, dispara, toasts).
8. Compositor `components/conversations/conversation-composer.tsx`: ao lado de "Anexar", botão
   **"Do EMR"** → modal `attach-emr-document-dialog.tsx` lista `useClinicalDocuments(patientId)` →
   seleção envia via `send-whatsapp` (modo arquivo se janela aberta, senão link). Só aparece quando
   a conversa tem paciente.

## Guardas / regras
- Paciente precisa de phone + `WhatsAppOptIn` (≠ unsubscribed) — senão 422 amigável.
- Modo arquivo respeita janela 24h (motor já existe em `sendWhatsApp`).
- Reusa correções desta sessão: `ee7b9de4` (mídia), `c842473f` (corpo template), `83d85a75` (log).
- Sem deploy sem ordem explícita; migration roda no deploy do `api`.

## Fases
- **F1 (backend):** ✅ migration 00059 + resolve/materialize + send-whatsapp + whatsapp-window + list. Build verde, migration aplicada no dev.
- **F2 (frontend):** ✅ `send-document-whatsapp-button.tsx` smart nos 3 cartões (lab-requests / issued-documents-card / prescriptions) + `lib/api/clinical-documents.ts`.
- **F3 (frontend):** ✅ botão "Do EMR" no `conversation-composer.tsx` + `attach-emr-document-dialog.tsx` (só conversa de paciente).
- Deploy: `api` primeiro (migration + endpoints), depois `web`. Autorizado 2026-06-29 ("construir F3 e deployar tudo junto").
