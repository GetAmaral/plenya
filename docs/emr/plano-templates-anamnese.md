# Plano — Catálogo de Templates de Anamnese (Escore Plenya) — fase MACRO

## Context

O EMR já tem o motor de templates de anamnese **pronto**: `AnamnesisTemplate` → `AnamnesisTemplateItem` (FK `score_item_id` + `order`), CRUD completo, handler/rotas (`/api/v1/anamnesis-templates`), e um **builder drag‑drop** no front (`apps/web/app/(authenticated)/anamnesis-templates/page.tsx` + `components/anamnesis/anamnesis-template-item-selector.tsx`). Um template é uma **coleção curada de referências a score items** (não cópia) — e o mesmo score item pode aparecer em vários templates (FK reuso). O `Area` é enum (`Medicina`/`Nutricao`/`Psicologia`/`Educacao Fisica`). **Logo: criar os templates é tarefa de DADOS (curadoria + seed), não de código.**

Hoje só existem 2 templates de placeholder, idênticos em dev e prod: **"Teste"** (35 itens, seed de dev) e **"Medica Inicial"** (0 itens). Servem de sucata — serão aposentados.

O objetivo desta fase é **fechar o catálogo macro** (quais templates, nomes, faixas, propósito), ancorado no material canônico Plenya (Continuum, Método AGIR, Escore Plenya), nas normas dos conselhos BR (CFM/CFN/CFP/CONFEF) e na prática world‑class de longevidade. O **conteúdo de cada template** (quais score items entram) é a **próxima fase**, discutida depois desta aprovação.

### Decisões do usuário (2026‑06‑06)
- **Manter** a entrada Continuum "concisa" + um complemento pós‑fechamento (2 templates), renomeando para nome clínico seguro (CFM desaprova "degustação"/venda no rótulo de consulta).
- Nutri/psico/educação física: **2 templates cada** (Avaliação Inicial + Acompanhamento); re‑escore trimestral fica com o médico.
- **Separar** as trilhas: consultas Plenya avulsas (fora do Continuum) ≠ templates Continuum.

### Princípios que governam o catálogo
- **Nomenclatura por conselho (default adotado):** nutricionista → **"Avaliação Nutricional"** (CFN 380/2005); psicólogo → **"Avaliação Psicológica"** (CFP 07/2003); educador físico → **"Avaliação Física", nunca "Anamnese"** (CONFEF 218/2010, risco de não‑conformidade). "Consulta" só para o médico.
- **Template clínico ≠ formulário de preparação do paciente.** Estes templates são do **clínico** (preencher durante/ao redor da consulta, captando score items). O form de preparação (curto, 5–8 min, motor Escore Light, pós‑agendamento — `docs/atendimento/fluxo-pre-consulta.md`) é outra coisa e fica fora desta fase.
- **Entrada + Complemento (Continuum) particionam o mesmo conjunto que a Inicial avulsa.** Entrada capta o subconjunto prioritário em 45 min; Complemento capta o restante dos itens médicos após o fechamento. Juntos ≡ "Avaliação Médica Inicial" avulsa.
- **Voz/regulatório:** prosa clínica PT‑BR; sem "medicina preditiva", sem superlativos (CFM 2.336/2023), sem rótulo comercial em consulta; nota/anamnese é ato clínico (base LGPD "tutela da saúde", sem gate de consentimento).

---

## Catálogo recomendado — 13 templates em 2 trilhas

### Trilha A — Plenya avulsa (clínica, fora do Continuum) — *Area = Medicina*
| # | Template | Faixa | Propósito |
|---|----------|-------|-----------|
| A1 | **Avaliação Médica Inicial** | 60–80 min | Intake médico completo (equivale a Entrada+Complemento do Continuum somados). |
| A2 | **Acompanhamento Médico** | 45–60 min | Retorno: subconjunto dinâmico (evolução, ajustes, itens que mudam). |
| A3 | **Revisão de Exames** *(recomendado adicionar)* | 30–45 min | Consulta dedicada à interpretação de exames/biomarcadores. Padrão em clínicas de longevidade; já há infra de results‑inbox no roadmap. |

### Trilha B — Continuum (programa)
**Médico** *(Area = Medicina)*
| # | Template | Faixa | Propósito |
|---|----------|-------|-----------|
| B1 | **Avaliação Médica de Entrada (Continuum)** | 45 min | A "entrada/encantamento" CONCISA, antes da apresentação do deck. Capta o subconjunto médico prioritário do Escore. (Rótulo clínico, não "degustação".) |
| B2 | **Complemento da Avaliação Médica (Continuum)** | 30–45 min | Pós‑fechamento: completa os itens médicos do Escore não cobertos na Entrada. |
| B3 | **Acompanhamento Médico (Continuum)** | 30–45 min | Encontro recorrente (cadência mensal/rotação AGIR). |
| B4 | **Reavaliação Médica Trimestral (Continuum)** | 45–60 min | Re‑escore trimestral dos itens médicos. |

**Nutricionista** *(Area = Nutricao)*
| # | Template | Faixa | Propósito |
|---|----------|-------|-----------|
| B5 | **Avaliação Nutricional Inicial (Continuum)** | 45–60 min | Avaliação nutricional de entrada (CFN). |
| B6 | **Acompanhamento Nutricional (Continuum)** | ~30 min | Retorno recorrente. |

**Psicólogo** *(Area = Psicologia)*
| # | Template | Faixa | Propósito |
|---|----------|-------|-----------|
| B7 | **Avaliação Psicológica Inicial (Continuum)** | ~50 min | Avaliação psicológica de entrada (CFP). |
| B8 | **Acompanhamento Psicológico (Continuum)** | ~50 min | Retorno recorrente. |

**Educador físico** *(Area = Educacao Fisica)*
| # | Template | Faixa | Propósito |
|---|----------|-------|-----------|
| B9 | **Avaliação Física Inicial (Continuum)** | ~60 min | Avaliação física de entrada (CONFEF — "Avaliação", nunca "Anamnese"). |
| B10 | **Acompanhamento (Educação Física) (Continuum)** | 30–45 min | Retorno recorrente. |

> O `Name` carrega trilha (Plenya/Continuum) + tipo (Inicial/Acompanhamento/…); o `Area` carrega a disciplina. **Sem mudança de schema.** Faixas de duração são descritivas (não há campo de duração no template hoje).
>
> *Fora deste catálogo (não são templates de anamnese):* a "Reunião de Equipe / Plano Unificado" da semana 0–1 é artefato de **plano de cuidado** (CarePlanItem/relatório), não anamnese.

---

## Bridge para a fase de CONTEÚDO (próxima, não agora)

A partição dos ~916–1211 score items entre os templates segue o mapa **disciplina ↔ letra AGIR** (M2M `score_item_method_pillars` já existe):
- **Médico** → **G** (Gestão Clínica e Metabólica, 23 pilares / ~1027 itens) + **R** (Ritmo, lidera) + Genética + Exames/labs + Histórico/familiar/medicações.
- **Educador físico** → **A** (Atividade física) + composição corporal + testes práticos/fitness.
- **Nutricionista** → **A** (Alimentação/Suplementação) + nutrologia/micronutrientes + antropometria.
- **Psicólogo** → **I** (Integração mente‑corpo).
- **Sobreposição esperada** (mesmo item em vários templates): peso/altura/composição (médico+nutri+edfísica); sono (médico+psico). Suportado nativamente.

Entrada (B1) + Complemento (B2) particionam o conjunto médico; B1 = subconjunto prioritário, B2 = o resto. Acompanhamentos = subconjunto que evolui no tempo. Reavaliação trimestral = itens re‑medidos.

---

## Implementação (executar SÓ após definir o conteúdo de cada template)
- **Curadoria** dos score items por template (fase de conteúdo, item a item, com o Dr.).
- **Seed idempotente** dos 13 templates: `INSERT` em `anamnesis_templates` (name+area) + `anamnesis_template_items` (`score_item_id`,`order`), com UUIDv7 determinístico para dev≡prod. Dev = **banco direto** (psql), nunca API HTTP.
- **Aposentar** os 2 placeholders ("Teste", "Medica Inicial") em dev e prod (soft delete) ao publicar o catálogo real.
- **Prod:** método dry‑run‑em‑prod com rollback (mesmo de AGIR/Bioma/genética) — aplicar txn na prod, conferir hashes, rollback, depois aplicar de verdade. Backup durável antes.
- Sem `pnpm generate`/migration (nenhum model muda). Opcional futuro: campo de "trilha/cadência" no template + vínculo template↔tipo de Appointment (fora de escopo).

## Verificação
- Catálogo macro **aprovado** pelo Dr. (esta fase) antes de curar conteúdo.
- Fase de conteúdo: por template, conferir cobertura (a união Entrada+Complemento ≡ Inicial avulsa; nada essencial fora; sobreposições intencionais), nomes conformes aos conselhos, e que cada item existe e está ativo no banco.
- Pós‑seed: builder UI lista os 13 templates com itens na ordem certa; abrir uma anamnese a partir de cada um popula os itens; QA visual Playwright (localhost:3000, bypass).

## Housekeeping pós‑aprovação
Copiar este plano para `docs/emr/plano-templates-anamnese.md` (versionado) e adicionar ponteiro na memória (`MEMORY.md`), conforme a regra de persistir plano aprovado em arquivo.
