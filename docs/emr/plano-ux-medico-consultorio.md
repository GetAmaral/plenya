# Plano — UX do médico no consultório (EMR Plenya)

## Context

A feature da Recepção/Secretária está concluída e no ar. O próximo foco é o **dia‑a‑dia clínico do médico** (Dr. Getúlio, nefrologista; clínica concierge de Saúde/Performance/Longevidade). Uma análise de 10 agentes sobre o code base + pesquisa de padrões EMR (internacional, regulatório BR e longevidade‑concierge) revelou que o **backend é rico mas a UX do ato clínico está quebrada no centro**:

- **Não existe espaço de consulta.** `/appointments/[id]` é só agenda. Os campos `Appointment.DoctorNotes` e `Appointment.Diagnosis` existem no DB + API (`PUT /api/v1/appointments/:id`) mas **não têm UI** — hoje o médico só conseguiria documentar via `curl`. Prescrever/pedir exame são seções top‑level que **perdem o contexto do paciente**.
- **Não há nota de evolução por visita** (SOAP). Anamnese é intake one‑shot, não evolução de retorno.
- **Falta o esqueleto clínico:** sem problem list, sem CID‑10 codificado (`Diagnosis` é texto livre), **sem alergias** (zero `allerg` nos models → risco direto ao prescrever), sem medicações‑em‑uso/reconciliação, sem sinais vitais por consulta (`PhysicalAssessment` é artefato do módulo Treino/ACSM).
- **O prontuário abre em demografia, não na síntese.** `/patients/[id]` só mostra dados cadastrais; não há at‑a‑glance (escore, AGIR, alergias, últimos labs).
- **Sem caixa de exames a revisar** (sem fila/ciência auditada/flag de crítico). **A home `/dashboard` é 100% mock** (dados hardcoded, botões que não navegam).
- **Ativos proprietários siloados:** Escore/AGIR/Continuum deveriam ser a **capa** do prontuário (tese de marca "Normal vs Ótimo" + tendência longitudinal), não abas separadas.

**Resultado pretendido:** uma UX em que o médico abre a consulta numa tela única que mostra a síntese do paciente e permite documentar a evolução, prescrever, pedir exame e finalizar **sem perder contexto**; o prontuário abre na síntese proprietária; e o EMR ganha o esqueleto clínico (alergias/problemas/CID/vitais/meds) + caixa de resultados + documentos assináveis, dentro das regras do CFM/Anvisa/LGPD.

**Decisões do usuário (2026‑06):** começar pelo **Workspace de consulta**; modelar a nota como **novo `ClinicalNote`** (assinável/imutável); **incluir o regulatório** (atestado assinável + Memed/CFM) no roadmap.

## Princípio de arquitetura (vale para todas as fases)

- **`Appointment` JÁ é o "encounter".** Tem ciclo `scheduled→confirmed→checked_in→in_progress→completed`, telemed (Daily.co), recall, `AnamnesisID`, `ContinuumItemID`. **NÃO criar um modelo `Encounter` genérico** (duplicaria ciclo de vida + recepção + telemed + calendário e exigiria migration destrutiva). Todo artefato clínico novo se vincula à visita por `appointment_id *uuid` **nullable** (permite também criação fora de consulta — ex.: lab batch ingerido por OCR).
- **Reuso agressivo de primitivos existentes** (Tiptap `RichTextEditor` da anamnese, `MedicationSearch`, `DailyCoEmbed`, `ScheduleRecallDialog`, `useLatestHealthScore`, `RadarAgir`, pivot de labs, `middleware.AuditLog`, `SignatureService` ICP‑Brasil, `useAppointments` da Recepção).
- **Convenções Plenya obrigatórias:** todo model com `BeforeCreate` UUID v7; fluxo `model → dto → service (toDTO) → handler (fino) → rota em cmd/server/main.go`; **migrations goose escritas à mão** em `apps/api/database/migrations/` (próxima é `00005_…`), aditivas, com `IF NOT EXISTS`/`ADD COLUMN IF NOT EXISTS` e `Down` (molde da `00003`); writes passam por `middleware.AuditLog`; rodar `pnpm generate` após editar models; todo formulário usa `useFormNavigation({ formRef })`; páginas de paciente usam `useRequireSelectedPatient()`. **Commit sempre direto no master.**
- **Guardrails regulatórios/voz (anti‑transplante americano):** validade de receita = assinatura **ICP‑Brasil** (já implementada) + entrega ao paciente (QR/validador ITI), **sem pipe EHR→farmácia**; CID‑10 (2019) agora, campo extensível para **CID‑11 só em jan/2027** (não migrar); LGPD do ato clínico = base **"tutela da saúde" (art. 11, II, f)**, **sem gate de consentimento** (consentimento só para telemed CFM 2.314/2022 e para incluir CID em atestado); **não construir TISS/TUSS**; **não reconstruir numeração SNCR** (integrar plataforma homologada); nota clínica **imutável pós‑assinatura** (adendo; guarda 20 anos/NGS2). Voz: prosa clínica conectiva PT‑BR, **sem "medicina preditiva", sem superlativos (CFM 2.336/2023), sem citar marcas, sem promessa de resultado, sem bloat** (hierarquia atenção→positivos→resumo).

---

## P0 — Workspace de Consulta + Nota de Evolução + Home real  *(a fatia‑1, o que o médico sente primeiro)*

**Objetivo:** desbloquear o ato de documentar e dar ao médico uma tela única de atendimento, além de trocar a home mock por agenda + pendências reais.

### Backend
- **Novo model `ClinicalNote`** (`apps/api/internal/models/clinical_note.go`) → tabela `clinical_notes`:
  - `id` (UUIDv7), `appointment_id *uuid` (idx), `patient_id` (idx), `author_id`, `layout` varchar check `soap|apso`, `subjective/objective/assessment/plan` text + `*_html` text (Tiptap, sanitizado com DOMPurify no display como na anamnese), `status` varchar check `draft|signed` default `draft`, `signed_at`, `amendment_of_id *uuid` (self‑FK; correção de nota assinada = nova nota apontando para a original), `visibility` (reusa enum `AnamnesisVisibility`: `all|medicalOnly|psychOnly|authorOnly`), timestamps + soft delete.
  - Hooks: `BeforeCreate` UUIDv7. Regra de serviço: nota `signed` é **read‑only** (correção só via adendo).
  - Migration goose `00005_clinical_notes.sql` (aditiva, molde `00003`).
- **Service/DTO/Handler/Rotas:** `clinical_note_service.go` (`Create`/`GetByID`/`ListByPatient`/`Update`/`Sign`/`Amend`, `toDTO`, checagem de `selectedPatient` como a anamnese), `dto/clinical_note.go`, `handlers/clinical_note_handler.go`, rotas em `cmd/server/main.go` sob `RequireAnyStaff` espelhando a autorização de appointment (médico só edita a própria nota). Writes via `AuditLog`.
- **Estender `medical_record_aggregate_service.go`:** adicionar o tipo `clinical_note` ao `UNION ALL` (mais um subquery) **e** expor flag `documentada` no entry de `appointment` (existe `clinical_note` com aquele `appointment_id`). Sem migration (só SQL).
- `pnpm generate` para tipos TS/Zod.

### Frontend
- **Workspace em `apps/web/app/(authenticated)/appointments/[id]/page.tsx`** quando `status ∈ {checked_in, in_progress}` — layout de 3 zonas:
  - **Esquerda — Síntese:** Escore + barras AGIR (`useLatestHealthScore`), últimos labs relevantes, motivo/queixa, banner de alergia (placeholder até P2). Embed Daily.co no topo quando `type=telemedicine` (já existe).
  - **Centro — Nota de evolução SOAP:** editor rich‑text reusando o `RichTextEditor` da anamnese (`apps/web/components/anamnesis/`), pré‑carregado com a estrutura S/O/A/P (não em branco). `useFormNavigation`.
  - **Direita — Ações clínicas:** Prescrever, Pedir exame, Calcular escore, Agendar retorno (`ScheduleRecallDialog`) + **Finalizar consulta**.
  - **Ao abrir:** setar `selectedPatient` automaticamente (mutation já existente) — assim os fluxos de prescrição/exame resolvem `patientId` pelo JWT sem mudança de backend. `PatientContextBar` visível.
  - **"Iniciar consulta"** (hook de start da Recepção) e **"Finalizar consulta"**: cria/assina a `ClinicalNote` e faz `PUT status=completed` (ordem: criar nota → PUT; se o PUT falhar, a nota persiste e o link é idempotente).
  - **Ações embutidas:** abrir os fluxos existentes `/prescriptions/new` e `/lab-requests` como `Dialog`/sheet ou deep‑link que retorna ao appointment — **não duplicar** os fluxos.
- **Home real `apps/web/app/(authenticated)/dashboard/page.tsx`:** remover o mock (`@tremor` hardcoded) e usar `useAppointments` (mesma query da Recepção, filtro `doctorIds/dateFrom/dateTo/status`, polling 15s): agenda de hoje (cards que navegam de verdade), bloco **"Em atendimento agora"** (`in_progress`) com link direto pro workspace, próximas consultas, e pendências derivadas (consultas `completed` sem nota; retornos/recalls).
- **Timeline `patients/[id]/prontuario/page.tsx`:** badge "Nota registrada" vs "Sem nota" usando o flag `documentada` do agregado; clique abre a `ClinicalNote`.

**Arquivos‑chave P0:** `models/clinical_note.go` (novo), `database/migrations/00005_clinical_notes.sql` (novo), `services/clinical_note_service.go` (novo), `dto/clinical_note.go` (novo), `handlers/clinical_note_handler.go` (novo), `cmd/server/main.go`, `services/medical_record_aggregate_service.go`, `appointments/[id]/page.tsx`, `dashboard/page.tsx`, `patients/[id]/prontuario/page.tsx`, `lib/api/calendar-api.ts`, `components/anamnesis/` (RichTextEditor).

**Regulatório P0:** nota = ato clínico (base "tutela da saúde", **sem** consentimento); nota assinável e imutável (NGS2/20 anos).

---

## P1 — Capa AGIR do prontuário + navegação

**Objetivo:** o prontuário abre na síntese proprietária; a navegação deixa de afogar as ferramentas do médico.

- **Reescrever `apps/web/app/(authenticated)/patients/[id]/page.tsx`** para abrir na **capa**: hero com Escore Plenya + radar AGIR (`RadarAgir`), faixa horizontal de biomarcadores‑chave em **Normal vs Ótimo** (valor do paciente plotado, cor pelo `LevelNumber` 0‑6 dos item results do snapshot, mini‑sparkline da tendência via `ScoreEvolutionChart`), e 3 cards na ordem **Atenção → Positivos → Resumo**. Demografia desce para drawer/aba "Dados cadastrais" (nada perdido). Estado vazio elegante (CTA "Calcular Escore") quando não há snapshot. Reusa `RadarAgir.tsx`, `ScoreEvolutionChart.tsx`, `ScoreLevelBadge.tsx`, `useLatestHealthScore`, pivot de labs colorido por level.
- **`PatientSummaryService`** (backend, **sem tabela nova**): compõe o at‑a‑glance num request (último snapshot + AGIR, últimos N labs, última consulta/nota, e — quando existirem em P2 — alergias/problemas/meds ativos), no molde read‑only do `medical_record_aggregate_service`.
- **Sidebar agrupada** (`apps/web/components/layout/collapsible-sidebar.tsx`, array `navigation[]`): seções colapsáveis **Clínico** (prontuário, consultas, escores, exames, prescrições, anamnese, continuum) · **Operação** (recepção, calendário, conversas, leads, campanhas) · **Treino** · **Config/Admin**. Só reestrutura o array + render; mantém `staffOnly/requiredRoles/badgeKey`.

---

## P2 — Esqueleto clínico + CDS de alergia

**Objetivo:** o EMR ganha o esqueleto que todo prontuário sério tem; prescrever passa a alertar alergia.

Novos models (goose `00006…`, todos UUIDv7, todos com `appointment_id`/`patient_id` conforme o caso):
- **`PatientAllergy`** (`allergies`): `substance`, `substance_type` (drug|food|environmental|other), `reaction`, `severity` (mild|moderate|severe|anaphylaxis), `status` (active|inactive), `no_known_allergies` bool (flag explícita), `recorded_by_id`. Index parcial em `(patient_id) WHERE status='active'`.
- **`PatientProblem`** (`problem_list_items`): `description`, `cid_code` varchar(16) **nullable**, `cid_version` varchar default `'CID-10'`, `status` (active|resolved|inactive), `onset_date`, `resolved_date`, `onset_appointment_id *uuid`, `recorded_by_id`.
- **`CIDCode`** (`cid_codes`): catálogo `code` PK, `description`, `chapter`, `version` default `'10'`, índice trigram/GIN p/ autocomplete PT‑BR. **Seed CID‑10 2019** via migration de dados.
- **`ConsultationVitals`** (`consultation_vitals`): `appointment_id` (1:1), `systolic_bp/diastolic_bp/heart_rate/resp_rate/temperature/spo2/weight/height/waist_circumference`, `bmi` (computado em `BeforeSave`), `measured_by_id`, `measured_at`. **Distinto** de `physical_assessments` (ACSM/treino — não tocar).
- **`MedicationInUse`** (`medications_in_use`): `medication_name`, `active_ingredient` (idx p/ CDS), `dosage/frequency`, `source` (prescribed_here|external|patient_reported), `source_prescription_id *uuid`, `status` (active|suspended|stopped), `reconciled_appointment_id`. **Distinto** de `Prescription` (ato de emitir).

UI/integração:
- Captura inline no **workspace** (vitais, problemas com autocomplete CID, reconciliação de meds) — não telas soltas.
- Banner de alergia na **capa** + na abertura do paciente.
- **CDS de alergia ao prescrever:** `AllergyService.CheckAgainstMedication(patientID, activeIngredient)` consumido pelo `prescription_service` antes de assinar (cross‑check por **princípio ativo** já normalizado em `prescription_medication`/`medication_definition`; alerta **não‑bloqueante**, médico confirma; respeita `no_known_allergies`).
- `Appointment.Diagnosis` (texto livre) evolui para a problem list codificada, mantendo o campo legado (não dropar).

**Regulatório P2:** `cid_version` acomoda CID‑11 (2027) sem migrar; CID em atestado só com consentimento (flag no documento, em P3).

---

## P3 — Caixa de resultados + documentos assináveis + controlados + plano AGIR/relatório

**Objetivo:** fechar o ciclo (revisão de exames, documentos legais, controlados) e materializar o diferencial longevidade‑concierge. *(Pode ser sequenciado em sub‑entregas.)*

- **Results inbox:** estender `lab_result_batches` (`acknowledged_at`, `acknowledged_by_user_id`, `is_critical`) e `lab_results` (`reference_min/max`, `abnormal_flag` normal|low|high|critical_low|critical_high — laboratorial, distinto do `level` de escore; backfill a partir de `lab_test_definitions`). Fila `/lab-results/inbox` (não‑revisados, críticos no topo) com **ciência auditada** via `AuditLog`. Reusa `lab_result_batch_service` e o pivot existente.
- **Atestado/declaração/laudo assináveis (baixo custo — pipe já existe):** generalizar `SignatureService.SignPrescriptionPDF → SignDocumentPDF(pdfBytes, doctorID)` (PAdES/ICP‑Brasil já genérico) e gerar `PatientDocument` (type `certificate`/`declaration`/`report`) com `signed_pdf_path/hash/qr_code_data/signed_at/certificate_serial` + `includes_cid` (gated por consentimento). Hoje atestado é só upload — passa a ser gerado e assinado no EMR com QR (validável no ITI/ATESTA CFM).
- **Controlados (dependência externa):** trocar o **stub `sncr_service`** por provider **Memed/CFM Prescrição Eletrônica** atrás da interface já existente (Portaria 344/98 + SNCR RDC 873/2024). **Não reconstruir numeração.** *Dependência out‑of‑band: conta/credencial Memed ou CFM (e‑CPF Médico em nuvem) — mapear antes desta sub‑fase.*
- **Telemed compliance:** carimbar consentimento (CFM 2.314/2022) como `telemed_consent_at/text` no `Appointment`, capturado no início da teleconsulta dentro do workspace (doc lado‑a‑lado). Opcional: plugar `transcription_service` ao recording do Daily.co.
- **Plano multidomínio AGIR + relatório (diferencial de marca):** `CarePlanItem` (pilar `method_pillar_id` + âncora `score_item_id`/`lab_test_code` + `recommendation` + `author_id/role` + `status`), agrupado por letra A·G·I·R, cada recomendação ancorada ao biomarcador (o M2M `score_item_method_pillars` já existe). Relatório longitudinal ao paciente da **mesma fonte** (Escore + AGIR + Normal vs Ótimo + plano), curado e **assinado** pelo médico antes de publicar no portal (reusa `signature_service` + `assessment_html_service` + `patient-portal`). QoL: renovar/copy‑forward de receita + posologia favorita; recomendação de exame a partir do escore. Recall **clínico‑relacional** (reteste/retorno), nunca win‑back.

---

## Riscos & mitigações
- **`selectedPatient` é estado global do médico (server‑side).** Setá‑lo ao abrir o workspace troca o contexto — manter `PatientContextBar` visível indicando o paciente ativo; o `AnamnesisService` já valida `patientId == selectedPatient`, então é o caminho seguro.
- **Finalizar = 2 chamadas (criar nota + PUT).** Ordem nota→PUT; falha do PUT preserva a nota; link idempotente no retry. Sem transação distribuída em P0.
- **Confundir `ConsultationVitals` com `PhysicalAssessment`** (treino/ACSM). Fronteira explícita: vital de consulta é leve/por‑visita; ACSM permanece do educador físico. Não migrar dados entre eles.
- **CDS de alergia com falso negativo** (nome comercial vs princípio ativo). Cross‑check por `active_ingredient` normalizado; alerta não‑bloqueante; `no_known_allergies` explícito.
- **Drift de schema prod** (existe `00004_reconcile_prod_schema`). Toda migration nova aditiva, `IF NOT EXISTS`, com `Down`; validar `migrate status` antes do `up`.
- **Editar nota assinada** (viola NGS2/20 anos). `status=signed` torna read‑only no service; correção só via `amendment_of_id`.

## Anti‑padrões (não fazer)
Modelo `Encounter` genérico; transplantar tela de EHR americano (campos vazios como entrada; pipe EHR→farmácia/Surescripts); reconstruir numeração SNCR; transformar Anamnese em evolução; migrar p/ CID‑11 agora ou hardcodar códigos fora de `clinical_codes.go`; gate de consentimento LGPD no ato clínico; construir TISS/TUSS; bloat de prontuário sem hierarquia; over‑engineering de imagem/genômica/wearables; voz "medicina preditiva"/superlativos; citar marcas; depoimentos/NPS/win‑back agressivo; remover `DoctorNotes`/`Diagnosis` (depreciar como legado, não dropar).

## Verificação (end‑to‑end)
- **Backend:** `docker compose exec -w /app api go build ./...` (exit 0); `docker compose exec -w /app api go run ./cmd/migrate up` aplica `00005…` e `down` reverte limpo sem tocar `appointments/anamnesis/physical_assessments`; `migrate status` consistente; `pnpm generate` produz os tipos sem diff inesperado fora dos gerados.
- **Fluxo P0 (dev, bypass auth em localhost:3000):** abrir uma consulta `in_progress` → paciente auto‑selecionado + capa com Escore/AGIR; escrever evolução SOAP no editor pré‑estruturado; "Finalizar consulta" cria a `ClinicalNote`, vincula e move status para `completed` numa ação; `PUT /api/v1/appointments/:id` aceita o link; a consulta aparece na timeline com badge "Nota registrada"; `/dashboard` mostra agenda real (sem dado hardcoded) com card "Em atendimento" clicável.
- **QA visual/runtime:** Playwright local (localhost:3000 bypass + prod público) nas telas tocadas — workspace, dashboard, prontuário, capa — confirmando 0 erro de runtime (build verde não pega regressão de runtime; dedupe react‑query após qualquer bump).
- **Por fase:** P2 → alerta de alergia dispara ao prescrever medicamento com princípio ativo conflitante (não‑bloqueante); P3 → atestado gerado assina com ICP‑Brasil e valida pelo QR; results inbox registra ciência auditada (`AuditLog`).
- **Conteúdo voltado ao paciente** (relatório/portal em P3): revisar voz (sem maneirismos de IA, sem preditiva, sem marcas, sem promessa) com o Getúlio antes de publicar.

## Pós‑aprovação (housekeeping)
Copiar este plano para `docs/emr/plano-ux-medico-consultorio.md` (versionado) e adicionar ponteiro na memória (`MEMORY.md`), conforme a regra de persistir plano aprovado em arquivo. Implementar **fase a fase**, commitando direto no master, validando build + QA por fase antes de seguir.
