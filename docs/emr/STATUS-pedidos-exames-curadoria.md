# ⭐ STATUS — Curadoria de Pedidos de Exames + Painéis + Score (ÂNCORA DE RETOMADA)

> **LER PRIMEIRO ao retomar.** Atualizado 2026-06-15. Tudo **só no dev** (banco + yml + planos em doc).
> **Nada commitado, nada em prod.** A mecânica da feature JÁ está em prod (commits antigos, ver abaixo);
> o que é dev-only são os **DADOS** (painéis/exames/vínculos) + os planos.

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
