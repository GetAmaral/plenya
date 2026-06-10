# Estudo — Simplificação Consulta × Anamnese (em avaliação)

> Status: **ESTUDO / a avaliar** (2026-06-09). Não implementado. Origem: Getúlio achou que
> "anamnese deveria estar dentro da consulta", e ao discutir concluímos que há **entidades
> demais** e que SOAP/APSO **atrapalham mais do que ajudam**.

## 1. O que existe hoje (e por que confunde)

Quatro entidades penduradas no paciente, três delas escrevendo **o mesmo tipo de dado**
(respostas de `ScoreItem`) em momentos/autores diferentes:

| Entidade | Quem | Quando | Conteúdo |
|---|---|---|---|
| `ConsultationPrep` (+`...Response`) | paciente | antes | itens `prep_order` + queixa |
| `Anamnesis` (+`AnamnesisItem`) | equipe | intake | narrativa (`content/summary`) **+** itens ScoreItem → alimentam o Escore; tem `visibility` |
| `ClinicalNote` | médico | cada visita | **SOAP** S/O/A/P + toggle SOAP/APSO; assinável/imutável |
| `Appointment` | recepção | evento | tem `AnamnesisID` (link **opcional, mão única**), `PrepFormVersionID`, status, telemed |

**Colisão de nome:** "anamnese" hoje significa duas coisas — a **narrativa** (texto livre) e os
**itens estruturados** (respostas de ScoreItem). Isso é parte da confusão.

Problemas concretos:
- `ClinicalNote` tem 4 campos SOAP + toggle SOAP/APSO → ninguém usa direito, atrapalha.
- A narrativa vive em dois lugares (`Anamnesis.content` e `ClinicalNote.Subjective`).
- Os itens estruturados (que viram Escore) **não aparecem na tela da consulta** — o médico sai
  pra `/anamnesis`, troca de paciente, etc.
- Não há **visão longitudinal por item** (o valor de "Água" ao longo do tempo, venha do Escore
  Light, da pré-consulta ou de consulta anterior).

## 2. Direção alvo (a avaliar)

### 2.1 Registro da consulta = 2 campos de texto livre
Aposentar SOAP/APSO. A nota da consulta passa a ter:
- **Anamnese** (texto da consulta / narrativa)
- **Conduta** (plano)

Mantém: assinatura/imutabilidade, adendo (`AmendmentOfID`), `visibility`, rich-text (texto+HTML),
vínculo ao `Appointment`. O AI-scribe da telemed já gera formato "anamnese" → encaixa direto.

### 2.2 Itens clínicos (ScoreItem) = entidade separada, linkada e visível na consulta
Os "anamnese itens" continuam como entidade própria (são o que alimenta o Escore), **mas**:
- ficam **linkados à consulta** (via `AnamnesisID` no appointment, ou item → appointment direto);
- **aparecem na tela da consulta** com a UX nova (densa, mobile-first — ver §4);
- cada item mostra seu **histórico** de forma fácil, **independente da origem** (Escore Light,
  pré-consulta, consulta anterior).

### 2.3 Destino da entidade `Anamnesis`
Como a narrativa migra pro registro da consulta (§2.1), a `Anamnesis` deixa de precisar de
`content/summary` e vira essencialmente o **container datado dos itens** de um encontro. Duas
opções (decidir):
- **(a)** Manter `Anamnesis` como bucket de itens (só esvaziar a narrativa); menor mudança.
- **(b)** Linkar `AnamnesisItem` direto no `Appointment` e aposentar a `Anamnesis` standalone;
  menos entidade, mais refactor. Perde o caso "anamnese sem consulta" (ex.: enfermagem) — checar
  se isso importa.

## 3. Histórico por item — "estude melhores formas"

Todas as fontes escrevem resposta de ScoreItem (nível + valor numérico + texto). Queremos **uma
linha do tempo por item**. Três caminhos:

| Abordagem | Como | Prós | Contras |
|---|---|---|---|
| **A. Tabela unificada** | uma `clinical_item_response` que TODAS as fontes gravam | timeline trivial | migration grande + reescrever 4 fluxos; arriscado |
| **B. Agregador read-side** (recomendado) | `ClinicalTimelineService` lê `AnamnesisItem` + `ConsultationPrepResponse` + sessões do Escore Light (já importadas) e **funde** por `ScoreItemID`+paciente, ordenado por data, com `source` | **sem migration**; fontes ficam como estão; bate com "itens podem existir como entidade separada" | precisa de UNION/ordenação no serviço; cache |
| **C. Event log** | append-only de eventos de item | auditoria forte | overkill agora |

**Recomendação: B.** É um *read-model*. Endpoint tipo
`GET /patients/:id/score-items/:itemId/history` → `[{date, source, level, numericValue, text}]`,
com `source ∈ {escore_light, pre_consulta, anamnese, consulta}`.

**UX do histórico (inline, por item):** na linha do item, um expander "Histórico" abre uma
mini-timeline: valor atual em destaque + 2-3 anteriores com **badge de origem** + data (ex.:
`Pré-consulta · 02/06 · N2` / `Escore Light · 14/05 · N0`). Para itens numéricos, um sparkline
opcional. Sem sair da tela da consulta.

## 4. UI/UX mobile dos itens (mock primeiro)

Dores do print atual (iPhone): renderer em modo "full" (desktop) — `container py-8 space-y-8`,
cards `p-5/space-y-6`, e **textarea "Observações (opcional)" sempre aberta** mesmo vazia (maior
desperdício vertical).

Plano (mock → aprovar → implementar):
1. **"Observações" colapsada** atrás de um "+ observação" (só abre quando quer).
2. **Modo compact forçado no mobile** (`p-3/space-y-2`).
3. **Níveis como chips que quebram linha** (não blocos `min-h-80px` empilhados).
4. **Grupos em accordion** — ver cabeçalhos, expandir um por vez.
5. **Cabeçalho apertado** no mobile (`py-8`→`py-4`).
6. **Histórico por item** inline (§3).

## 4.1 Implementado nesta rodada (2026-06-09) — UI mobile + histórico

Aprovado o mock, foi implementado (dev, **não deployado**):

**Backend (histórico por item, abordagem B — sem migration):**
- `dto/clinical_timeline.go`, `services/clinical_timeline_service.go` (funde
  `anamnesis_items` + `consultation_prep_responses` submetidos + `anonymous_score_items`
  claimed, ordena por data desc), `handlers/clinical_timeline_handler.go`.
- Rota: `GET /api/v1/patients/:id/score-items/:scoreItemId/history` (grupo `patients`,
  RequireAnyStaff + AuditLog). Compila OK.

**Frontend (redesign do modo `compact` = form regular/mobile; fullscreen tablet intacto):**
- `lib/api/clinical-timeline.ts` (`useItemHistory`, lazy on-expand).
- `components/anamnesis/AnamnesisItemHistory.tsx` (timeline com badge de origem + sparkline).
- `AnamnesisTemplateItemsRenderer.tsx`: novo branch `compact` — accordion de grupos, linha
  densa com chip do nível atual, item expansível (chips de nível, observação colapsada atrás
  de "+ adicionar observação", histórico inline). Não-compact (tablet) inalterado.
- `anamnesis/page.tsx`: container apertado no mobile.

Pendente de revisão visual em produção (decisão do Getúlio).

## 4.2 Implementado — SOAP → 2 campos (2026-06-09)

Getúlio: "não tem nota SOAP, sistema em beta, não precisa armazenar nada" → troca limpa,
sem migração de dados. Implementado (dev, **não deployado**):

- **Migration `00037_clinical_note_simplify.sql`**: dropa `layout` + S/O/A/P (×8 colunas),
  adiciona `clinical_history(_html)` + `conduct(_html)`. Aplicada no dev (goose v37).
- **Model `clinical_note.go`**: remove `ClinicalNoteLayout` + 4 seções; adiciona
  `ClinicalHistory(Html)` + `Conduct(Html)`.
- **DTO + service**: Create/Update/toDTO nos 2 campos (sem layout). Compila OK.
- **Frontend `clinical-notes.ts`**: tipos/payloads para `clinicalHistory*`/`conduct*`.
- **`consultation-workspace.tsx`**: remove toggle SOAP/APSO; 2 editores ("História clínica e
  evolução" + "Conduta"). AI-scribe: subjetivo+objetivo+avaliação → História; plano → Conduta.
  Transcrição → História clínica.

> ⚠️ Follow-up: rodar `pnpm generate` (OpenAPI/tipos) — model mudou. Não é necessário pro
> funcionamento (front usa tipos manuais), e o gerado já tem quebras pré-existentes.

## 4.3 Implementado — Anamnese embutida na consulta (2026-06-09)

Decisão §2.3 = **(a)**: `Anamnesis` segue como container datado de itens, **linkada** ao
appointment via `Appointment.AnamnesisID` (sem aposentar a entidade). Implementado (dev, **não
deployado**):

**Backend:**
- `AnamnesisService`: refatorado `Create` → extrai `createForPatient`; novos
  `GetByAppointment` (anamnese vinculada, centrado na consulta) e `CreateForAppointment`
  (cria + seta `Appointment.AnamnesisID`, idempotente).
- `AnamnesisHandler`: `GetByAppointment` + `CreateForAppointment`.
- Rotas: `GET/POST /api/v1/appointments/:id/anamnesis`. Compila, sobe, rota responde.

**Frontend:**
- `lib/api/appointment-anamnesis.ts` (hooks get/create/update).
- `components/consultations/consultation-anamnesis-panel.tsx`: painel colapsável no centro do
  workspace — sem anamnese, oferece modelos e "iniciar"; com anamnese, renderiza o renderer
  denso (itens + histórico) com **autosave debounced**.
- `consultation-workspace.tsx`: painel embutido abaixo da nota (coluna central).

Agora os itens da anamnese (com histórico por origem) aparecem **dentro** da consulta; o médico
não navega pra `/anamnesis`.

## 5. Decisões em aberto
- §2.3: manter `Anamnesis` como bucket (a) ou linkar item→appointment e aposentar (b)?
- Migração das notas SOAP existentes/assinadas: notas assinadas são imutáveis (legal/NGS2 20
  anos). Provável: **manter colunas antigas read-only** (notas SOAP antigas exibem 4 seções) e
  **novas notas** usam Anamnese+Conduta. Confirmar com a régua legal.
- Onde a pré-consulta/Escore Light "entram" como itens da consulta: importar pra `AnamnesisItem`
  do encontro, ou só exibir via timeline (B) sem materializar?
