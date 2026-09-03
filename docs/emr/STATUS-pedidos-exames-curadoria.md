# ⭐ STATUS — Curadoria de Pedidos de Exames + Painéis + Score (ÂNCORA DE RETOMADA)

> **LER PRIMEIRO ao retomar.** Atualizado 2026-06-16.
> ✅ **DEPLOYADO EM PROD (2026-06-16).** Toda a curadoria de painéis (18 painéis, 437 vínculos, 528
> exames requisitáveis) + a integração de score (25 exames revisados, 17 pontuam) que eram **dev-only**
> foram subidos pro prod via seed UPSERT (dry-run em clone → transação única → verificado: prod ≡ dev,
> paciente intacto). teste/teste2 apagados. Backup pré-deploy guardado (local + VPS). App redeployada
> (Coolify concurrent_builds=1, um app por vez). Commits de doc/yml no master.

## Ajuste 2026-06-17 (UX dos painéis — dev-only, prod pendente)
- **Nomes:** removido o prefixo "Painel " (e "Plenya ") de todos — ex.: "Painel Plenya Completo" → **Completo**.
- **Ordem:** Completo · Inicial · Acompanhamento, depois os 15 complementos em **ordem alfabética**
  (via `display_order` 1..18; backend ordena por `display_order, name`; front idem).
- **Dentro do painel** (`apps/web/lib/lab-request-apply.ts`): **laboratório** primeiro (alfabético,
  contíguo, pagina automático a 40/pág), depois **imagem** (`category=imaging`) alfabético **um por
  página** (linha em branco = quebra de página no render `pdfdoc/exam_request.go`).
- **Isolamento:** `/lab-requests` lista só do `selectedPatient` (era `getAllLabRequests` global — bug).
- **Novo paciente** vira `selectedPatient` automaticamente também no cadastro completo `patients/new`.
- yml regenerado do banco. **Prod ainda com nomes "Painel X"** → seed UPSERT (rename + display_order) sob ordem.

## Ajuste 2026-09-03 (Doppler de tireoide — dev-only, prod pendente)
- **Acrescentado** `PLNUSGTIRDOP01` — "Ultrassonografia com Doppler colorido de tireoide",
  **TUSS 40901386**, ao lado da USG de tireoide nos 4 templates que a pedem (Acompanhamento,
  Completo, Inicial, Tireoide). Não substitui: no TUSS são dois procedimentos, e a operadora
  cobra assim.
- **O TUSS foi conferido em duas fontes**, e há duas armadilhas registradas no seed:
  não existe código específico de "Doppler de tireoide" (o certo é o genérico de órgão isolado);
  `40901378` é vaso cervical VENOSO, não tireoide; e `40901203`, o da USG que já estava no
  catálogo, é "Órgãos superficiais (tireóide ou escroto ou pênis ou crânio)" e não inclui Doppler.
- `page_break_before = false`: quem abre o bloco de imagem é a USG de tireoide, e o Doppler
  pertence à mesma página. Conferido pelo caminho real (`applyTemplate` sobre o payload da API):
  a página de USG sai com a USG, o Doppler, o abdome e a próstata.
- **Vai por migration goose `00097_usg_doppler_tireoide.sql`**, não por seed em `docs/emr/`.
  Dado de catálogo neste projeto é migration (precedente: `00056` insere em
  `lab_test_definitions`; `00081`/`00082` são catálogos inteiros) — assim chega em prod no deploy
  do `api`, em vez de depender de alguém lembrar de rodar um script. Provado: `up` sobre o dado já
  existente é no-op, `down` devolve a ordem original e apaga a definição, `up` de novo reproduz o
  estado byte a byte.

> Nota sobre a linha de 2026-06-17 acima: "imagem alfabético um por página" descreve uma regra que
> não vive mais no código. `apps/web/lib/lab-request-apply.ts` diz explicitamente que **a ordem e a
> paginação vêm do template (dado), não de agrupamento no código** — `display_order` e
> `page_break_before`. Foi por isso que o ajuste acima é de dado, não de código.

## Próxima tarefa (o que estávamos fazendo)
**PLANO DE SCORE em execução item a item (2026-06-15).** #1–#11 pontuam/ajustados; #12–16 (cluster
HAS Secundária) fora; **falta revalidar #17–#25** (cortisol salivar, 17-OHP, pregnenolona, estrona,
anti-dsDNA, CA 19-9, CA 15-3, TTPa, EPO — retomar por **cortisol salivar**). Detalhe em
`plano-score-exames-faltantes.md` (§Progresso) + refactor de sono em `plano-refactor-score-sono.md`.
Inclui ajustes além dos "faltantes": **SIBO 63→26**, **refactor objetivo de sono** (IAH/SRI/TST/
Eficiência), **Rx tórax com 3 achados** (nódulo+ICT+calcif. aórtica), **TI-RADS leve**.

## O que é esta feature
Sistema de **templates de pedido de exames** no EMR: 3 painéis macro (Completo/Inicial/Acompanhamento)
+ complementos empilháveis, com **filtro automático por sexo**, **TUSS resolvido no render do PDF**,
**import/dedup de pedido externo**, e **justificativa por exame** (`request_justification`, carrega ao
selecionar; linha `#` no texto livre). **Mecânica 100% em prod** (commits `68f5303d`, `7ae82314`,
`7b3c0357`, `0539fab5`; migrations `00038_lab_test_sex_applicability`, `00039_lab_test_request_justification`).

## Estado dos painéis (dev) — 18 painéis, 156 exames requisitáveis
| ord | painel | itens |
|---|---|---|
| 1 | Completo | 131 | 
| 2 | Inicial | 97 |
| 3 | Acompanhamento | 66 |
| 4 | Trombose *(ex-"Risco Vascular", renomeado)* | 3 |
| 5 | Hematológico | 13 |
| 6 | Inflamação | 5 |
| 7 | Oncologia | 11 |
| 8 | Pesquisa | 4 |
| 9 | Risco Cardiovascular | 17 |
| 10 | Gastrointestinal | 11 |
| 11 | Autoimune | 6 |
| 12 | Tireoide | 8 |
| 13 | Metais Pesados | 8 |
| 14 | Hepático | 6 |
| 15 | Hormônios Esteroides | 13 |
| 16 | Sorologias Virais | 6 |
| 17 | HAS Secundária | 23 |
| 18 | Sono | 9 |

## Onde vive o estado (fonte da verdade = banco dev)
- **Banco dev** (`plenya_db`): `lab_request_templates`, `lab_request_template_tests`, `lab_test_definitions`.
- **Backup legível:** `docs/emr/templates-pedido-exames.yml` — regenerado do banco ao fim de cada bloco
  (script de regeneração foi rodado via psql dump + python; 18 painéis, ~437 vínculos, 2 justificativas).
- **Curadoria/decisões:** `docs/emr/templates-pedido-exames-curadoria.md` (histórico, parcial/antigo).
- **Plano de score:** `docs/emr/plano-score-exames-faltantes.md` (RASO — a aprofundar item a item).

## Exames criados NESTA sessão (17) — novos no catálogo
sangue oculto imunoquímico (40303250) · H. pylori respiratório (40307786) · Aldosterona (40316050) ·
Renina (40316432) · Metanefrinas urinárias 24h (40311163) · Catecolaminas fracionadas (40311058) ·
Polissonografia PSG (40103528) · **Índice de ômega-3 (sem TUSS)** · Cortisol pós-dexa (40316190) ·
Poligrafia respiratória HSAT (sem TUSS) · Actigrafia (sem TUSS) · 17-OH-progesterona (40316017) ·
**Pregnenolona (sem TUSS)** · Cortisol salivar ritmo (40317374) · Anti-dsDNA (40306062) ·
CA 19-9 (40316378) · CA 15-3 (40316378). Painéis novos: **HAS Secundária**, **Sono**.

## Triagem do score (ponto de partida da discussão a aprofundar)
25 exames requisitáveis sem `ScoreItem`. Triagem PRELIMINAR (a revisar item a item):
- **A) Pontuar numérico:** Índice de ômega-3 (alvo >8%).
- **B) Flag binário:** Sangue oculto (positivo), H. pylori (positivo).
- **C) Achado estruturado:** sono (IAH/eficiência da PSG/HSAT/Actigrafia); SIBO já tem itens de score.
- **Reconciliar:** Glicemia/Insulina avulsas já pontuam via componentes da curva (não duplicar).
- **D) FORA (diagnóstico dirigido):** HAS secundária (aldo/renina/metanefrinas/catecolaminas/
  cortisóis/17-OHP/pregnenolona), anti-dsDNA, CA 19-9/15-3, TTPa, EPO, Estrona.
> Getúlio achou isso raso — aprofundar pesos, cutoffs, pilares, literatura, e revalidar a lista D.

## Regras de processo aplicadas nesta curadoria (manter ao retomar)
- **Cadência item-a-item, de 5 em 5; NUNCA avançar de bloco antes de o Getúlio fechar o atual.**
- **Banco + yml SEMPRE sincronizados** — regenerar o yml do banco ao fim de cada bloco.
- **TUSS e preços: SEMPRE verificar online**, nunca chutar (vários TUSS corrigidos no caminho).
- **Não transplantar padrões americanos** (Tier 3 — ADMA/GlycA/NMR/Galleri/relógio epigenético — ficou FORA).
- Justificativa de exame = `request_justification` (longevidade p/ TC tórax; MESA/Agatston p/ CAC).

## Pendências (NÃO autorizadas — só sob ordem)
1. **Commit** da curadoria (yml + doc + plano) no master.
2. **Deploy pra prod** dos DADOS: dry-run em clone → aplicar seed no banco prod → deletar `teste`/`teste2`.
3. **Aprofundar + executar o plano de score** (próxima discussão).
