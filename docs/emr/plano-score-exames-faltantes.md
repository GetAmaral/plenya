# Plano — Integrar ao Escore os exames requisitáveis que ainda não pontuam

**Status:** proposto · **Criado:** 2026-06-15 · **Origem:** QA da curadoria de painéis de pedido de exames.

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
