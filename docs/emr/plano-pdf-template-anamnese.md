# Plano — PDF de exibição dos templates de anamnese (para revisão)

**Status:** proposta · **Data:** 2026-06-23

## Objetivo
Gerar um PDF por **template de anamnese** mostrando todo o conteúdo do template (grupos →
subgrupos → itens, com níveis/perguntas), pra o Dr. Getúlio **avaliar e revisar cada template**.
Inspirado no PDF do score completo que já existe.

## Como o "PDF do score" funciona hoje (a referência)
Há **dois mecanismos** no projeto:
1. **Server-side (Go + Chromium)** — `internal/services/score_pdf_service.go`: carrega a árvore
   do score, renderiza um template Go HTML (`internal/templates/score_poster.html`) e converte
   via `internal/pdfdoc/render.go` (`pdfdoc.RenderHTML`, Chromium headless/rod, fila+browser
   compartilhados). Endpoint `GET /api/v1/score-groups/poster-pdf`. Gera o **pôster** (60×200cm)
   e também tem opção A4 (`a4PDFOptions`). É um PDF vetorial baixável.
2. **Front-side (print)** — páginas `app/(authenticated)/scores/print/page.tsx` e `scores/poster/`
   que usam `window.print()` + `@media print`. O usuário "salva como PDF" pelo navegador.

Dados do template já estão prontos: `AnamnesisTemplateRepository.GetByID(id, withItems=true)`
faz `Preload(Items → ScoreItem.Subgroup.Group + ScoreItem.Levels)` — tudo que o PDF precisa.

## Decisão central: onde renderizar (front print vs server-side Go)
A diferença que importa: **as perguntas das escalas** (PHQ-9, GAD-7, Dubois, ASEX…) vivem no
registro TS `@plenya/domain` (`SCALE_REGISTRY`). O Go **não** enxerga esse registro.

- **Opção A — Página de print no front** (`/anamnesis-templates/[id]/print`, `window.print`):
  reaproveita os componentes e o `SCALE_REGISTRY` → consegue mostrar as **sub-perguntas das
  escalas**, preview no navegador, "salvar como PDF". Mais rica e mais rápida de construir.
  Espelha o padrão `/scores/print`. Contra: é "imprimir → PDF" pelo navegador, não um arquivo
  gerado por endpoint.
- **Opção B — Endpoint server-side** (`GET /anamnesis-templates/:id/pdf`, Go + `pdfdoc`):
  PDF vetorial baixável, espelha exatamente o `score_pdf_service`. Contra: o Go não acessa o
  `SCALE_REGISTRY` → as sub-perguntas de escala teriam que ser **omitidas** (mostra a escala
  como um item de nível normal) ou **duplicadas em Go**. Mais trabalho e duplicação.

**Recomendação: Opção A** (print no front). O objetivo é revisão visual; a página de print
mostra o conteúdo completo (incluindo escalas), dá preview imediato e reusa o que já temos.
Se você quiser um arquivo gerado/baixável por endpoint (ex.: anexar em e-mail, gerar em lote
no servidor), aí vale a Opção B.

## Layout proposto (A4, retrato, paginado — 1 review por folha[s])
Cabeçalho: nome do template, área/trilha, contagem de itens, data. Marca Plenya discreta.
Corpo, por **grupo → subgrupo → item**:
- Nome do item (+ unidade, se houver).
- **Modalidade de resposta**:
  - níveis → lista N0…N5 com nome e faixa (cores como no score: vermelho→verde);
  - numérico → unidade + faixas dos níveis;
  - **escala** (PHQ-9 etc.) → lista as sub-perguntas + opções (via `SCALE_REGISTRY`).
- **Aplicabilidade**: filtros de gênero / faixa etária / pós-menopausa (deixa claro itens
  condicionais).
- (Opcional, modo "verboso") relevância clínica / conduta / explicação ao paciente.

Respeita a supressão pai/filho e o filtro demográfico já usados no
`AnamnesisTemplateItemsRenderer` (pra o PDF refletir o que o profissional vê).

## Implementação (Opção A)
1. **Rota** `app/(authenticated)/anamnesis-templates/[id]/print/page.tsx` — busca o template
   (`useAnamnesisTemplate(id)`), monta a visão de revisão (read-only, não o renderer de
   preenchimento), com CSS `@media print` (page-break por grupo). Botão "Imprimir / Salvar PDF".
2. **Componente** `AnamnesisTemplateReview` — organiza itens por grupo/subgrupo (reusa a lógica
   de `organizeTemplateItems`) e, para itens de escala, lê o `SCALE_REGISTRY` para listar
   sub-perguntas. Reusa cores de nível (`LEVEL_*` do renderer).
3. **Botão "PDF/Imprimir"** na lista de templates (`anamnesis-templates/page.tsx`), abrindo a
   rota de print em nova aba.
4. (Opcional) toggle compacto/verboso na própria página de print.

## Implementação (Opção B, se preferir endpoint)
1. `internal/templates/anamnesis_template.html` (Go template, A4).
2. `AnamnesisTemplatePDFService`: `GetByID(id, true)` → organiza por grupo/subgrupo → HTML →
   `pdfdoc.RenderHTML(html, a4PDFOptions())`.
3. `GET /api/v1/anamnesis-templates/:id/pdf` (clinician) → download.
4. Botão no front (apiClient blob). Escalas: sem sub-perguntas (ou duplicar textos em Go).

## Decisões pra você
1. **Opção A (print no front, recomendada) ou B (endpoint server-side)?**
2. **Conteúdo**: compacto (nome + níveis/modalidade + aplicabilidade) ou verboso (também
   relevância clínica/conduta)? Recomendo compacto, com toggle pra verboso.
3. **Em lote?** Por ora 1 template por vez; dá pra fazer um "todos" depois.
