# Plano — Integrar ao Escore os exames requisitáveis que ainda não pontuam

**Status:** ✅ revisão dos 25 CONCLUÍDA (dev) · **Criado:** 2026-06-15 · **Origem:** QA da curadoria de painéis de pedido de exames.

**Placar final: 17 exames pontuam · 8 fora.**
- **Pontuam (17):** #1 ômega-3 · #2 glicemia + #3 insulina (consolidados) · #4 sangue oculto · #5 H.pylori · #6 PSG + #7 HSAT + #8 actigrafia (via achados de sono IAH/SRI/TST/eficiência) · #9 SIBO · #10 USG tireoide (TI-RADS) · #11 Rx tórax (3 achados) · #17 cortisol salivar (ritmo+excesso) · #19 pregnenolona · #20 estrona · #21 anti-dsDNA · #22 CA 19-9 · #23 CA 15-3.
- **Fora (8):** #12 aldosterona · #13 renina · #14 metanefrinas · #15 catecolaminas · #16 cortisol pós-dexa · #18 17-OHP · #24 TTPa · #25 EPO (diagnósticos dirigidos).
- Levels (topo→5) + site_legend auditados. Embeddings 17/17 (itens novos).
- ✅ **DEPLOYADO EM PROD (2026-06-16)** junto com a curadoria de painéis inteira (que também era dev-only). Seed via UPSERT FK-safe (dump→staging→upsert, `session_replication_role=replica`), dry-run em clone do prod restaurado (config ficou byte-equivalente, paciente intacto), aplicado em transação única. PROD agora ≡ DEV: 18 painéis / 437 vínculos / 665 defs / 528 requisitáveis / 1230 score items / 4570 levels / 2674 pilares. teste/teste2 apagados. Embeddings dos 19 novos gerados no prod. Backup prod pré-deploy: `plenya_db_PROD_FULL_20260615-231539.dump` (local + VPS `~/db-backups/`). App redeployada (Coolify `concurrent_builds=1`).

## ⚠️ Mecânica correta dos levels + checklist por item (aprendido tarde, 2026-06-15)
**`GetLevelMultiplier` é tabela FIXA pelo nº absoluto do level** (não `level/máx`): 0→0.0 · 1→0.2 · 2→0.4 · 3→0.6 · 4→0.8 · **5→1.0 · 6→1.0**. Convenção: level 0 = pior, **melhor TEM que estar no 5/6**. Binário = level 0 (ruim) e 5 (bom). Itens que criei com topo <5 estavam subpontuando o caso perfeito (bug).

**Definition of done por item de score:**
1. `score_item`: `clinical_relevance` + `patient_explanation` + `conduct` (feito).
2. `score_levels`: levels com **melhor no 5/6, pior no 0**; **`site_legend` em CADA nível** (legenda curta paciente, sem travessão). ← eu esquecia isso.
3. Embeddings: **lote no fim** via `EmbeddingQueueService.QueueAllScoreItems()` → worker gera vetor → article links **auto**. (Links manuais = curadoria à parte.)

**🔧 DÉBITO RETROATIVO (re-review um a um em curso):** corrigir levels (topo→5) + preencher `site_legend` em TODOS os itens que criei. Depois rodar QueueAllScoreItems.
- ✅ **TODOS corrigidos (levels topo→5 + `site_legend` em cada nível):** #4 sangue oculto · #5 H.pylori · #6 IAH · #7 SRI · #8 Eficiência · TST · #9 SIBO (ΔH₂/CH₄/H₂S) · #11 Rx (nódulo/ICT/CAA) · #17 cortisol (ritmo/excesso) · #19 pregnenolona · #20 estrona · #21 anti-dsDNA · ômega-3 · TI-RADS. Auditoria: **0 níveis sem site_legend**.
- ✅ **Embeddings GERADOS (15/15)** — enfileirados na `embedding_queue` + worker em background processou. `auto-link-articles-rag` rodado (idempotente).
- ⚠️ **Article links = 0 nos nossos itens** por **lacuna de corpus**: só 4/1043 artigos tangenciam esses temas; nenhum dedicado a ômega-3 index/IAH/SRI/TI-RADS/H.pylori/sangue oculto/anti-dsDNA/cortisol. Não é config — falta conteúdo. Fix: `cmd/fetch-missing-articles` (PubMed) p/ popular o corpus desses temas → re-rodar auto-link. **Não afeta o cálculo do score.** Pendente sob ordem.

## Progresso (2026-06-15)
- **Passaram a pontuar / foram ajustados (#1–#11):** #1 ômega-3 (novo) · #2+3 glicemia+insulina (consolidados, já pontuavam) · #4 sangue oculto (consolidado+binarizado, 12) · #5 H. pylori (novo, 8) · #6+7+8 sono (camada objetiva nova: IAH 18/SRI 18/TST 16/Eficiência 12) · #9 SIBO (revisado 63→26) · #10 TI-RADS tireoide (novo leve, 6) · #11 Rx tórax (3 achados novos, 5/5/5).
- **FORA do escore — cluster HAS Secundária, aprovado (#12–16):** Aldosterona, Renina, Metanefrinas, Catecolaminas, Cortisol pós-dexa. Diagnósticos dirigidos (achar/excluir hiperaldosteronismo/feocromocitoma/Cushing), não graduáveis; seguem como exames de investigação no painel HAS Secundária.
- **#17 Cortisol salivar → PONTUA (2 achados, contra a triagem inicial):** `Cortisol salivar — ritmo diurno` (8; Hormonal+Estresse) + `Cortisol salivar — excesso/hipercortisolismo` (8; Hormonal+Metabólico+CV). Eixos ortogonais, níveis qualitativos. Evidência: ritmo achatado → mortalidade (Whitehall II HR 1,30); excesso subclínico → CV/metabólico/ósseo/mortalidade. Lição: nem todo "diagnóstico dirigido" é fora — o cortisol tem dimensão de longevidade real.
- **FORA confirmado (#18):** 17-OHP — rastreio de CAH/21-hidroxilase (limiar diagnóstico), sem gradiente de longevidade (≠ cortisol). Segue como exame de investigação.
- **#19 Pregnenolona → PONTUA leve (decisão do Dr., contra minha recomendação de fora):** 4 pts, **curva em U** (pico 125–175 ng/dL=1.0; deficiência <33=0; >248=0.2), pilares Hormonal+Função Cognitiva. Evidência fraca + cutoffs exploratórios (faixa "ótima" é convenção funcional) → peso baixo. **TUSS `40317056` corrigido** na definição (faltava desde a criação do exame) + yml ressincronizado (3 painéis).
- **#20 Estrona → PONTUA leve, só pós-menopausa:** 5 pts, `female`+`post_menopause`, maior=pior (≤50=1.0 … >200=0), pilares Hormonal+Rastreio Oncológico. E1 alta na pós-menopausa = aromatização adiposa → risco mama RE+/endométrio. Estradiol já cobre o eixo (7 itens); estrona entra só no nicho pós-menopausa. Cutoffs aproximados (ensaio-dependentes).
- **#21 Anti-dsDNA → PONTUA graduado por atividade lúpica (decisão do Dr.):** 12 pts (espelha TRAb), `negativo <30 IU/mL = nota máxima` · 30–75=.67 · 75–200=.33 · >200 (atividade alta)=0. Pilares Imune e Inflamatório + Renal (nefrite). Cortes ensaio-dependentes.
- **#22 CA 19-9 → PONTUA** (14 pts; normal <37 U/mL no level 5; GI+Rastreio Onco) e **#23 CA 15-3 → PONTUA** (14 pts; normal <30 U/mL; Rastreio Onco). Correção da triagem: o sistema **já scoreia marcadores tumorais** (CEA/CA-125/AFP/PSA), então tumorais entram graduados por elevação (coerência).
- **#24 TTPa → FORA** (cutoff em segundos é reagente-dependente, sem normalização tipo INR; natureza diagnóstica; trombofilia já coberta por F2/F5/D-dímero).
- **#25 Eritropoietina → FORA** (só interpretável junto da Hb; diagnóstico contextual de anemia/policitemia, sem direção de score isolada).
- ✅ **REVISÃO DOS 25 COMPLETA.** Embeddings dos novos: 17/17 (15 + CA 19-9/CA 15-3). Article links seguem fora (lacuna de corpus, decisão do Dr.).
- **Tudo dev-only.** Commit + deploy prod pendentes sob ordem.


## ✅ Decisões travadas (aplicadas no dev)
> Mecânica confirmada: `score_levels.level` = **rank de qualidade** (0=pior … 5=melhor); multiplicador implícito = `level/maxlevel`. Item lab vive em grupo **Exames › Laboratoriais**; itens de especialidade ficam em **zero `score_version`** (contam no escore EMR completo, fora de Light/Triagem — padrão de 227/246 itens lab).

| # | exame | code | pontos | níveis | pilares AGIR | aplicado |
|---|---|---|---|---|---|---|
| 1 | **Índice de ômega-3 (EPA+DHA)** | PLNOMEGA301 | **16** | 6 (banda-alvo 8–11% ótima; U suave: `>11%`=0.8; progressivo na subida `<4`=0 … `6–8`=0.6) | Cardiovascular · Nutrologia e Micronutrientes · Função Cognitiva | ✅ dev |
| 2+3 | **Glicemia de jejum + Insulina (jejum)** — consolidação (opção B), **sem item novo** | PLN9AF0BCF5 / PLN24EA4ACE | (já 35 / 15) | (já existentes) | (já existentes) | ✅ dev |

| 4 | **Sangue oculto nas fezes (imunoquímico)** | PLN771B34C5 | **12** | binário: Positivo=**level 0** · Negativo=**level 5** (1.0) | Rastreio Oncológico · Gastrointestinal | ✅ dev (levels corrigidos) |
| 5 | **H. pylori (teste respiratório)** | PLNHPYLOR01 | **8** | 2 binário (Positivo=0 · Negativo=1) | Gastrointestinal · Rastreio Oncológico | ✅ dev |
| 6+7+8 | **Sono — refactor objetivo** (PSG/HSAT→IAH; Actigrafia/wearable→SRI/TST/Eficiência) | PLNSLEEPIAH/SRI/TST/EFF | 18/18/16/12 | ver doc | Sono/Crono/CV/Vitalidade | ✅ dev |
| 9 | **SIBO — revisão de pontuação** (já era scoreado; reweight) | SIBO_TEST (8 componentes) | **63→26** | Δ H₂(10)+CH₄ pico(10)+H₂S pico(6); 5 zerados | Gastrointestinal | ✅ dev |
| 10 | **Nódulo de tireoide (ACR TI-RADS)** — estrutural leve (anti-sobre-diagnóstico) | PLNTIRADS01 | **6** | 4 cat: TR1=1.0 · TR2–3=.8 · TR4=.2 · TR5=0 | Rastreio Oncológico | ✅ dev |
| 11 | **Rx Tórax — 3 achados** (nódulo/massa · ICT · calcif. arco aórtico) | PLNRXTNOD01/CTR01/AAC01 | **5/5/5** | ver doc | Pulmonar·Rastreio Onco / CV / CV | ✅ dev |

**#11 (Rx tórax — refeito com achados prognósticos):** 1ª versão (só nódulo, 1/5 da TC) foi expandida após estudar o que o Rx mostra de relevante. 3 achados gradáveis, **5 pts cada**: (1) **Nódulo/massa** (cortes realistas de Rx: Ausente/<8mm · 8–30mm · >30mm; Pulmonar+Rastreio Onco); (2) **Índice cardiotorácico** (cardiomegalia >55% → CV mortalidade; Cardiovascular); (3) **Calcificação do arco aórtico** (preditor forte/independente de eventos CV — Guangzhou Biobank 28a; Cardiovascular). Achados não-requisitáveis (Rx segue orderável). Lit.: CTR cutoff 0,5 (PMC8125954); AAC (Atherosclerosis 2009, Lancet WP 2022). **Grupo ④ concluído.**

**#10 (USG tireoide):** eixo **funcional** da tireoide já completíssimo (TSH/T4L/T3L/T3R/TRAb/anti-Tg/anti-TPO). Faltava só o estrutural. Nódulo é comum e quase sempre indolente → item **leve** que só penaliza TR4/TR5 (risco real: TR4 ~5–19%, TR5 >40%), TR1–3 sem perda. Achado não-requisitável (USG segue orderável). Anti-sobre-diagnóstico (não transplantar rastreio agressivo americano).

**#9 (revisão SIBO):** já era scoreado via 8 componentes do teste expirado = **63 pts** (desproporcional + redundante: Basal+Pico+Delta por gás). Revisado p/ **26 pts** alinhado ao **Consenso Norte-Americano 2017**: H₂ pontua pela **elevação Δ≥20** (critério SIBO), CH₄ pelo **absoluto ≥10** (critério IMO, via pico), H₂S leve (sem consenso). Os 5 redundantes (H₂ basal/pico, CH₄ basal, Δ CH₄, H₂S basal) → **0 pts** (informativos). Corrigido tb o H₂S que pulava o nível 3.

**#6/#7/#8 (refactor de sono):** score de sono era ~69 itens **100% subjetivos** (anamnese), zero objetivo. Adicionada camada objetiva em `Sono › Medidas objetivas` (4 achados não-requisitáveis alimentados por exame OU wearable). Não pontuar estágio (REM/profundo) de wearable (validade 50–86%). Detalhe completo + decisões em aberto (§4 sobreposição, §5 enxugar minúcia, O5 ODI/SpO₂): **`docs/emr/plano-refactor-score-sono.md`**.

**#4 (correção de duplicação que eu criei + binarização):** ao desenhar os painéis eu criei `PLNSANGOCF1` sem ver que já existia o item pontuado `PLN771B34C5` ("Sangue Oculto (Hemoglobina)", 32 pts, 5 níveis guaiaco, pilar **errado "Renal"**). Fix: `PLN771B34C5` vira a única definição (requisitável + pontuada, nome/tuss do FIT, painéis Completo+Gastrointestinal); `PLNSANGOCF1` aposentado. Níveis → **binário** positivo/negativo. Pontos **32→12**. Pilar **Renal→Rastreio Oncológico+Gastrointestinal**. yml ressincronizado (2 blocos). Lição: **conferir o que já existe no banco ANTES de criar exame novo no painel** (evitar duplicação). Lit.: FIT = padrão CRC; laudo BR qualitativo; corte 10–20 µg Hb/g.

**#2/#3 (consolidação, não duplicação):** havia 2 definições por analito — uma só-pedido (`PLNGLIJEJ01`/`PLNINSJEJ01`, requisitável, sem score) e uma só-score (`PLN9AF0BCF5`/`PLN24EA4ACE`, não-requisitável, 35/15 pts, com `alt_names` ricos). Risco: matcher de OCR poderia rotear um resultado avulso pra def sem score (ambiguidade de nome). **Fix:** a def **pontuada** virou a requisitável (herda nome/tuss de pedido + os 5/2 painéis), e a só-pedido foi **aposentada** (`deleted_at`). Direção escolhida preserva a constante Go `LabCodeGlicemiaJejum="PLN9AF0BCF5"` (zero recompile). Resultado: 1 definição por analito, serve pra pedir **e** pontuar, sem ambiguidade. yml ressincronizado (7 trocas de `codigo`).

Literatura #1: HS-Omega-3 Index — zonas alto `<4%`/interm. `4–8%`/baixo `>8%`; alvo cardioprotetor banda **8–11%** (menor mortalidade total/MACE; sangramento e FA minimizados na banda); sinal de FA em U acima da banda. Refs: AJCN (CHD risk factor), PMC3942733, Atherosclerosis 10-coortes (PMID 28511049), PMC12122841 (FA double-edged), PMC10602979.


## Contexto
Após a expansão do catálogo (156 exames requisitáveis, 18 painéis), **25 exames requisitáveis não têm
`ScoreItem`** — o resultado deles não entra no Escore Plenya. Este plano define **quais devem pontuar,
de que forma, e quais NÃO devem** (testes diagnósticos dirigidos não se graduam).

> Fonte da verdade é o banco (dev). Mecânica do escore: `ScoreItem` (com `lab_test_code` → exame,
> `points`, gênero/idade) + `ScoreLevels` (faixas com multiplicador; convenção nível alto = bom,
> mult(0)=0) + distribuição AGIR (M2M pilares) + inclusão em `score_version` + campos de site se
> for paciente-facing. Ver [[genetica_pontos_niveis_status]] e `plano-agir-distribuicao-scoreitems.md`.

## Triagem dos 25 (decisão por exame)

### A) Pontuar como biomarcador graduado — `ScoreItem` numérico + níveis
| Exame | Racional | Pilar AGIR sugerido |
|---|---|---|
| **Índice de ômega-3 (EPA+DHA)** | biomarcador graduável de longevidade; alvo > 8% (HS-Omega-3 Index). Níveis: <4 ruim · 4–8 intermediário · >8 ótimo | Cardiovascular / Nutrologia (G) |

**Reconciliar (já pontuam por componente, NÃO duplicar):**
- **Glicemia de jejum** e **Insulina (jejum)** — o escore já lê `GLICOSE 0 MIN` / `INSULINA 0 MIN`
  (componentes da curva). Decisão: apontar o exame avulso para o MESMO item de score do componente
  0 min, ou deixar como está (sem score próprio). Não criar item duplicado.

### B) Pontuar como flag binário (normal/alterado)
| Exame | Racional | Pilar |
|---|---|---|
| **Sangue oculto nas fezes (imunoquímico)** | positivo = sinal de alerta (rastreio CRC) | Rastreio Oncológico (G) |
| **H. pylori (teste respiratório)** | positivo = fator de risco GI tratável | Gastrointestinal (G) |

### C) Pontuar via ACHADOS ESTRUTURADOS do laudo (não o exame em si)
O exame é imagem/funcional; o que pontua é o achado numérico/categórico extraído.
| Exame | Achado(s) a criar como item | Pilar |
|---|---|---|
| **Polissonografia / HSAT** | **IAH** (índice apneia-hipopneia), SpO₂ nadir, eficiência do sono | Sono (ACTS-S) / Ritmo |
| **Actigrafia** | tempo total de sono, regularidade circadiana | Sono / Ritmo |
| **SIBO (H₂/metano)** | **já tem itens de score** (H₂ basal, Δ metano, etc.) — só falta conferir/linkar o exame | Gastrointestinal |
| **USG de tireoide** | nódulo / classificação TI-RADS (opcional, baixa prioridade) | Tireoide |
| **Rx tórax PA+perfil** | achado (baixa prioridade) | Pulmonar |

### D) NÃO pontuar — diagnóstico dirigido, não marcador de otimização
Interpretam-se clinicamente (achar/excluir doença); não fazem sentido como pontuação graduada.
- **HAS secundária:** Aldosterona, Renina, Metanefrinas urinárias, Catecolaminas fracionadas,
  Cortisol pós-dexametasona, Cortisol salivar (ritmo), 17-OH-progesterona, Pregnenolona.
- **Autoimune dirigido:** Anti-DNA nativo (anti-dsDNA).
- **Marcadores tumorais:** CA 19-9, CA 15-3 (elevação = investigar, não pontuar; rastreio em
  assintomático já é controverso).
- **Coagulação dirigida:** TTPa.
- **Outros dirigidos/nicho:** Eritropoietina, Estrona.

> Resultado da triagem: dos 25, **~3 viram score numérico/flag** (ômega-3, sangue oculto, H. pylori),
> **~5 viram itens de achado estruturado** (sono + SIBO já existente), **2 reconciliam** com a curva
> (glicemia/insulina), e **~15 ficam fora** (diagnósticos dirigidos). O escore NÃO deve inchar com
> testes diagnósticos.

## Mecânica de implementação (por item que vai pontuar)
1. **`ScoreItem`** — criar com `lab_test_code` (liga ao exame), `points` (peso, calibrar pelos itens
   análogos já pontuados), `gender`/`age_range` se aplicável, `name` clínico + campos de site
   (`site_question`, linguagem paciente) se for paciente-facing.
2. **`ScoreLevels`** — faixas com `operator`/cutoff + `multiplier` (convenção nível alto = bom;
   mult(0)=0). Ex. ômega-3: 4 níveis (<4 / 4–6 / 6–8 / >8).
3. **AGIR** — vincular ao(s) pilar(es) na distribuição M2M (`plano-agir-distribuicao-scoreitems.md`).
4. **`score_version`** — incluir nos versions pertinentes (Plenya completo; provavelmente fora da
   triagem/light).
5. **Hooks** — `BeforeCreate` (UUID v7), `BeforeUpdate` (`LastReview`).

## Fases
- **Fase 1 (rápida, alto valor):** ômega-3 (numérico) + sangue oculto + H. pylori (flag). 3 itens.
- **Fase 2 (sono):** IAH/eficiência da polissonografia/HSAT/actigrafia + conferir SIBO. Exige definir
  os campos de resultado estruturado do sono.
- **Fase 3 (reconciliação):** glicemia/insulina avulsas ↔ componentes da curva (decisão do Dr.).
- **Achados D:** nenhuma ação (ficam fora por design).

## Verificação
- `migrate`/seed em dev → conferir que o item pontua num paciente dummy (lançar resultado, ver o radar).
- Recalcular gate das versions (não quebrar contagens light/triagem).
- Dry-run em clone do prod → aplicar prod (banco direto, idempotente) → commit no master.

## Decisões pendentes do Dr. (antes de executar)
1. **Pontos/peso** de cada item novo (ômega-3, flags) — calibrar com o Dr.
2. **Glicemia/insulina avulsas:** apontar pro score da curva ou deixar sem score próprio?
3. **Sono:** quais achados estruturados queremos pontuar (IAH só, ou + SpO₂/eficiência)?
4. Confirmar a lista **D (fora do score)** — algum deles o Dr. quer pontuar mesmo assim?
