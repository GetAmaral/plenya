# P3 — Sub-especificação de execução (decisões fixadas)

> Complementa `docs/emr/plano-ux-medico-consultorio.md` (plano-mãe). Escopo escolhido pelo
> usuário em 2026-06-02: **frentes 1 + 2 + 3** (a 4, Memed/CFM, fica de fora — bloqueada em
> credencial externa). Base do levantamento: workflow `p3-understand` (4 leitores paralelos).
> Estado de partida: prod @`0d51fdf0`, goose v7, próxima migration = **00008**.

## Ordem de implementação
1. **Frente 1 — Results inbox** (independente). Migration `00008`.
2. **Frente 2 — Documentos assináveis** (provê `SignDocumentPDF` + `IssuedDocument`). Migration `00009`.
3. **Frente 3 — Plano AGIR + relatório** (relatório reusa `SignDocumentPDF`/`IssuedDocument` da frente 2). Migration `00010`.

Commit por frente, direto no master. **Push único ao final** (após QA verde das 3) → um ciclo de
deploy só (o clone do web é lento; evitar 3 rebuilds). Migrations 00008–00010 aplicam em prod no deploy.

---

## Frente 1 — Results inbox (exames a revisar)

**Mecânica de criticidade (descoberta-chave):** NÃO há `reference_min/max`/`abnormal_flag` em
`lab_results`. A criticidade vem do campo `level *int` (0=pior/vermelho … 5=melhor/emerald, 6=N/A;
NULL=não classificado), preenchido por `LabResultBatchService.ClassifyBatchResults` (que casa o
valor contra `ScoreLevel.EvaluatesTrue` via `ScoreItem` aplicável). **Crítico = level 0/1; anormal = level≤2.**

**Decisões:**
- Granularidade da ciência = **por BATCH** (1 PDF = 1 coleta). Migration `00008` faz `ALTER lab_result_batches ADD`:
  `reviewed_at timestamptz`, `reviewed_by_user_id uuid`, `is_critical boolean NOT NULL DEFAULT false`,
  `worst_level smallint` (= MIN(level) dos results; menor = pior). Backfill por UPDATE a partir de `lab_results.level`.
  Índice parcial `WHERE reviewed_at IS NULL AND deleted_at IS NULL`.
- `ClassifyBatchResults` passa a **recomputar** `worst_level`/`is_critical` ao final.
- Fila **cross-patient** (NÃO selectedPatient): novo método `ListPendingReview(limit,offset)` —
  `WHERE result_date IS NOT NULL AND reviewed_at IS NULL AND deleted_at IS NULL`, order
  `is_critical DESC, worst_level ASC NULLS LAST, result_date DESC`, preload `Patient` + `LabResults.LabTestDefinition`.
  `CountPendingReview() (total, critical int)` p/ badge. `AcknowledgeBatch(batchID, userID)` seta reviewed_at/by.
  `GetByIDForReview(batchID)` (detalhe cross-patient, sem gate selectedPatient).
- Rotas: novo grupo `/api/v1/lab-result-inbox` = `Auth + RequireClinician + AuditLog`:
  `GET /` (lista), `GET /count` (badge), `GET /:id` (detalhe), `POST /:id/acknowledge` (ciência).
  RequireClinician (revisar exame é ato clínico; exclui secretary/manager).
- DTO `LabInboxItemResponse { id, patientId, patientName, laboratoryName, collectionDate, resultDate,
  status, isCritical, worstLevel, abnormalCount, totalResults, reviewedAt }`.
- Frontend: `lib/api/lab-inbox.ts` (`useLabInbox`, `useLabReviewCount({enabled})`, `useLabInboxDetail`,
  `useAcknowledgeBatch`); página `app/(authenticated)/lab-results/revisar/page.tsx` (críticos no topo,
  expandir result table colorida por level, "Dar ciência" + link "Abrir prontuário"); sidebar item
  "Exames a revisar" no grupo Clínico com `badgeKey:'lab-review'` (estender union + ramo em NavItemRow).

**Gotcha latente a corrigir se tocar:** `patient_labs_service.go:65,100` usa coluna `batch_id` (real é
`lab_result_batch_id`) — query falha silenciosa. Fora do caminho da frente; só corrigir se mexer no arquivo.

---

## Frente 2 — Documentos assináveis (atestado/declaração/laudo)

**Reuso:** `SignatureService.SignPrescriptionPDF(pdfBytes, doctorID)` **já é genérico** (PAdES/ICP via
digitorus/pdfsign). Cert do médico vive em `User` (CertificatePFX/Password criptografados), carregado por
`CertificateService.GetActiveCertificate`. PDF da prescrição via **gofpdf** (codeberg.org/go-pdf/fpdf) +
QR (skip2/go-qrcode); fontes OpenSans + letterhead `/app/PlenyaA4-150dpi.png` + selo `/app/icp-brasil-seal.png`
(confirmados no container). `pdf_service.go:addSignatureSection` (l.122) é o **melhor molde** de bloco de
assinatura (selo+QR+ITI+CFM 2.299/2021 + fallback manual). `lab_request_pdf_service.go` é o molde da
**degradação graciosa** (sem cert → regenera PDF não assinado, hash do não-assinado, não bloqueia).

**Decisões:**
- Generalizar assinatura: extrair `signPDF(pdfBytes, doctorID, info SignInfo)` interno; `SignPrescriptionPDF`
  vira wrapper (não quebra prescrição/lab_request); novo `SignDocumentPDF(pdfBytes, doctorID, reason, name)`.
- Novo model **`IssuedDocument`** (`issued_documents`, migration `00009`) — distinto de `PatientDocument`
  (que é metadado de UPLOAD). Campos: id(uuidv7), patient_id, appointment_id*, doctor_id, type
  (`certificate|declaration|report`), title, body text, purpose*, days_off* (atestado), includes_cid bool,
  cid_code*, cid_consent bool, status(`draft|signed`), signed_at*, signed_pdf_hash*, certificate_serial*,
  qr_code_data*, has_digital_signature bool, patient_document_id* (FK p/ a cópia publicada no portal),
  issued_at, timestamps + soft delete.
- Fluxo de emissão (`DocumentPDFService.GenerateSignedDocument(docID)`): carrega IssuedDocument+Patient+Doctor →
  valida → gera PDF gofpdf (cabeçalho médico + corpo por type + QR → `https://plenya.com.br/documentos/validar/<id>`)
  → checa `Doctor.CertificateActive`: se ativo `SignDocumentPDF`; senão bloco "assinatura manual" + has_digital_signature=false
  → **publica via `PatientDocumentsService.CreateFromBytes(Type=certificate, bytes)`** (resolve o `/uploads` morto:
  vai pra patient-docs + download authed + aparece no portal) → grava patient_document_id + hash/serial/qr/signed_at/status=signed.
- Rotas: `issued-documents` grupo `Auth + RequireAnyStaff + AuditLog`: `GET /:id`, `GET /:id/pdf` (download staff
  via patient_document_id → SendFile). `POST /` e `POST /:id/sign` com `RequireDoctor()` inline. Patient-scoped:
  `patients.Get("/:id/issued-documents", ...)`. Público (sem auth): `GET /documents/validate/:id` (recalcula hash, dados mascarados).
- includes_cid: form `includeCid + cidCode + cidConsent`; backend exige `cid_consent=true` se `includes_cid=true`. Default SEM CID.
- Frontend: `lib/api/issued-documents.ts`; card "Documentos emitidos" na capa + ação "Emitir documento" no workspace
  (dialog: tipo, título, corpo, dias de afastamento p/ atestado, toggle CID+consentimento); após assinar
  `window.open` do download staff (apiClient blob); página pública `app/documentos/validar/[id]/page.tsx`.

**Gotcha CRÍTICO:** `/uploads` NÃO é mais servido estático (removido por segurança H1). NUNCA retornar
`/uploads/...` como URL clicável — usar o download authed (CreateFromBytes + DownloadDocument / SendFile).

---

## Frente 3 — Plano AGIR + relatório longitudinal

**Descobertas-chave:** AGIR **não é persistido** — o radar é derivado em runtime no frontend a partir de
`ItemResults[].Item.MethodPillars` (M2M). A **M2M `score_item_method_pillars` está quase vazia** (≈30 de
~800 items mapeados) e os pilares seedados (13) divergem do canônico de `apps/site/lib/agir-structure.ts` (14 em G).
Letra AGIR = `MethodLetter.Code` string literal `'A'/'G'/'I'/'R'`.

**Decisão central (evita a dependência da M2M esparsa):** `CarePlanItem` ancora por **`letter_code` A/G/I/R**
(sempre presente, canônico, estável) — o médico escolhe a letra. Âncora a biomarcador é **opcional**:
`score_item_id*` e/ou `lab_test_code*`. Sem reseed de pilares neste round; o radar segue usando o `buildAgir` já vivo na capa.

**Decisões:**
- Model **`CarePlanItem`** (`care_plan_items`, migration `00010`): patient_id, `letter_code` varchar(1) check
  `IN ('A','G','I','R')`, score_item_id*, lab_test_code*, recommendation text (req), rationale* text, priority
  (`high|medium|low` default medium), status (`active|achieved|suspended` default active), target* text,
  source_snapshot_id*, author_user_id, author_role, timestamps + soft delete. (Sem container CarePlan — "plano"
  = conjunto de items ativos do paciente; mirror de problem_list_items.)
- Service/DTO/handler path-scoped `patients.Get/Post/Put/Delete("/:id/care-plan-items...")` — writes com
  `RequireClinician()` inline (multidisciplinar: cada clínico adiciona à sua letra). ListByPatient agrupável por letra.
- Endpoint `GET /api/v1/score-items/search?q=` (molde `cid_service.Search`: `unaccent(name) ILIKE`) p/ autocomplete da âncora.
- Frontend: extrair `buildAgir` de `patients/[id]/page.tsx` p/ `lib/agir.ts` (compartilhado). `CarePlanCard` agrupado
  por A·G·I·R (4 cores do MethodLetter) na capa + workspace; dialog add/edit (select letra + recommendation +
  priority + âncora opcional via score-item autocomplete). RadarAgir reusado.
- **Relatório longitudinal** (peça maior, paciente-facing): `CarePlanReportHTMLService` (molde
  `assessment_html_service.go` — strings.Builder + CSS @media print) compõe Escore + radar AGIR + Normal-vs-Ótimo
  (biomarcadores por level) + plano agrupado por letra → `ScorePDFService.generatePDFFromHTML` (go-rod/Chromium) →
  `SignDocumentPDF` (frente 2) → `IssuedDocument(type=report)` publicado no portal via CreateFromBytes.
  **Publish é ação manual e deliberada do médico** (botão "Gerar/assinar relatório") — conteúdo paciente-facing
  exige revisão de voz com o Getúlio (sem maneirismos de IA, sem preditiva, sem marcas, sem promessa) antes de publicar.

---

## Convenções (do P0–P2, valem aqui)
- Migrations goose aditivas, `IF NOT EXISTS`, CHECK `chk_<tabela>_<campo>`, índice parcial p/ ativos +
  índice em deleted_at, bloco Down com DROP; FKs só nos models GORM (migrations não declaram FK). Aplicar
  `docker compose exec -w /app api go run ./cmd/migrate up`; testar down/up. Registrar models no AutoMigrate
  (`internal/database/database.go`, antes da l.338).
- Fluxo `model→dto→service(toDTO)→handler(fino)→rota main.go`. Compilar no container (`go build ./...`).
- **🪲 GORM inicialismo:** campo com "ID"/maiúsculas consecutivas → tag `gorm:"column:..."` (SpO2→spo2, CIDCode→cid_code).
- Frontend: hooks LOCAIS em `lib/api/*.ts`; `next build` (Docker) REJEITA import não-declarado (usar `isomorphic-dompurify`).
  Cards em `components/clinical/`; `<select>` nativo `selectCls`; toast sonner; padrão dos cards P2.
- QA: Playwright dev (localhost:3000, bypass=admin@plenya.com; paciente com escore = João
  `019b998c-bfd6-7ae8-8e02-5bd4de22650c`) + verificação no DB via psql. Build verde NÃO pega mismatch de coluna.
- Deploy: push → Coolify auto-deploy 3 apps. Web clone LENTO (~5–15min, sem log — NÃO re-disparar).
  Verificar tags `ssh plenya "docker ps"`, migrations PG prod (container `mb511beqjtgd7nsjlnngh3m6`), smoke público.
