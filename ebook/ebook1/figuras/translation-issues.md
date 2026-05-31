# Figure Translation Issues — PT → EN

## ✅ STATUS PÓS PIL POST-FIX (2026-05-07)

Aplicado tooling PIL surgical pra corrigir os erros mais críticos de texto:

| Figura | Fix aplicado |
|---|---|
| **Cap04 Fig01** | Título "No Risk Factors" → "None Optimal" ✅ |
| **Cap09 Fig01** | FIGURA → FIGURE 1 + título "Three syštemas" → "Three systems" + subtítulo sentido CORRIGIDO ("stopped being") ✅ |
| **Cap08 Fig01** | "Mark" → "Marcos" + "+1,8 MET" → "+1.8 MET" (3x) + "34,3" → "34.3" + footer "lower all-cause" ✅ |
| **Cap13 Fig02** | Título "exams" → "labs" + "Lay" → "Secular" ✅ |
| **Cap01 Fig02** | Título "great" → "optimal" ✅ |
| **Cap12 Fig02** | Subtítulo "two markers" → "these markers" + "came off their place" → "moved" ✅ |

**Backup das versões pré-PIL** em `figuras/en-pre-pil-fix/`.

## 🔧 REMAINING (pra editor humano em Photoshop/GIMP)

### Decimal commas em corpo de figuras
gpt-image-2 preserva agressivamente vírgulas decimais PT mesmo após múltiplas iterações. Editor humano deve buscar e substituir nas figuras:
- Cap04 Fig01: "1,5" "1,9" "12,4" "3,9" → "1.5" "1.9" "12.4" "3.9"
- Cap04 Fig02: "4,8–5,2%" "5,4%" "6,5%" "3,5" "1,0" → period decimals
- Cap08 Fig02: "0,5 mg/day" → "0.5 mg/day"
- Cap11 Fig01: "1,0" "0,9" "1,7" "4,8" "11,2" → period decimals
- Cap12 Fig02: "1,8" "0,7" → "1.8" "0.7"
- Cap13 Fig02: "1,4" "0,6" "0,8" "1,7" → period decimals

### Outras paráfrases residuais
- Cap08 Fig02: "DO NOT prescribe finasteride or dutasteride" → remover "or dutasteride"
- Cap13 Fig02 body: "yes, checking 3h-5h" (Close friends row) → "friends arriving 3h-5h"
- Cap04 Fig02 body: "myocardial damage" → "myocardial injury"; "measured one once in life" → "measure once in a lifetime"

### Demais minor-fixes-needed (12 figuras)
Ver seções individuais abaixo. São paráfrases de severity low/med — editor pode revisar holisticamente.

---


Re-run cirúrgico das 11 figuras com prompts anti-paráfrase específicos. Resultado:

| Figura | Issue principal corrigida? | Decimais PT→EN? | Outros |
|---|---|---|---|
| Cap01 Fig02 | ❌ "great" ainda em vez de "optimal" | ❌ ainda 1,0 / 2,4 / 13,8 | — |
| Cap04 Fig01 | ❌ título virou "No Risk Factors" (regressão!) | ❌ ainda 1,5 / 1,9 / 12,4 / 3,9 | "Reference scale" ✅ |
| Cap04 Fig02 | ⚠️ kidney/injury parcial | ❌ ainda 4,8–5,2% / 5,4% / 6,5% | — |
| Cap06 Fig04 | ✅ "when fasting lies" | ✅ ok | — |
| Cap08 Fig01 | ⚠️ FIGURE 1 ✅, "less death risk" semântica ok | ❌ ainda +1,8 MET | ⚠️ "Marcos" virou "Mark" (regressão) |
| Cap08 Fig02 | ❌ "or dutasteride" ainda inventado | ❌ ainda 0,5 mg/day | — |
| Cap09 Fig01 | ❌ "have become" ainda invertido + FIGURA não traduzido | n/a | — |
| Cap11 Fig01 | ✅ "without replacement" | ❌ ainda 4,8 / 11,2 / 0,9 / 1,7 | — |
| Cap12 Fig01 | ✅ "after the workday" | n/a | — |
| Cap12 Fig02 | ❌ "two markers" ainda errado | ❌ ainda 1,8 / 0,7 | — |
| Cap13 Fig02 | ❌ "exams" / "checking" / "Lay ritual" ainda errados | ❌ | — |

**Padrão crítico descoberto:** gpt-image-2 **preserva agressivamente** decimais com vírgula PT mesmo quando o prompt pede explicitamente para converter. Parece que o modelo trata a vírgula como elemento visual da figura, não como dado textual a traduzir.

**Conclusão:** o método gpt-image-2 atinge teto com paráfrases sticky. Os 11 críticos precisam de tratamento manual (Photoshop/GIMP/PIL overlay) para decimais e termos teimosos.

---


Documentação de discrepâncias entre o texto pedido (estrito) e o que `gpt-image-2` entregou na tradução das figuras. **Tratamento posterior** — não tentar corrigir aqui, apenas registrar.

Formato por figura:
- **Pedido (EN exato)** → **Entregue por gpt-image-2** | **Severidade** (low/med/high)

Severidade:
- `low` — paráfrase estilística sem mudar significado (ex: "Compared with" → "Compared to")
- `med` — palavra adicionada/removida que muda nuance ou peso (ex: adicionou "behavior")
- `high` — número errado, termo técnico errado, sentido alterado, omissão crítica

---

## Cap07 Fig01.PNG — Extreme sedentary behavior raises death risk ~5x

Resultado após prompt reforçado anti-paráfrase (segunda iteração).

| Pedido | Entregue (reforçado) | Severidade |
|---|---|---|
| Title: "Extreme sedentary behavior raises death risk ~5x" | "Being extremely sedentary raises death risk ~5x" | low |
| Subtitle: "Compared with other established risk factors" | "Compared with other known risk factors" | low ("known" ≠ "established") |
| Bar 1: "Low cardiorespiratory fitness (extreme sedentary)" | "Low cardiorespiratory fitness (extreme sedentary behavior)" | low ("behavior" adicionado) |
| Bar 2: "End-stage kidney disease" | "End-stage renal disease" | low ("renal" ≠ "kidney") |
| Bar 5: "Coronary heart disease" (sem parêntese) | "Coronary heart disease" | OK correto |
| Reference: "(people in excellent fitness)" | "(person in excellent health and with no risk factors)" | med (inventou "no risk factors") |
| Callout: "More than smoking, diabetes, or heart disease." | "More than smoking, diabetes, or heart disease." | OK correto (Oxford comma) |

**Visual:** OK fonte/cores/layout/números idênticos ao PT.
**Texto:** ~75% exato, ~25% paráfrases low/med (sem alterar sentido clínico).
**Severidade geral:** minor-fixes-needed

---

## Cap01 Fig01.PNG — Three 'normal' check-ups — a dangerous trend

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| "Lab upper limit of normal" | "Laboratory normality limit (lab)" | low (paráfrase awkward) |

**Visual:** OK layout/cores/números (4.9, 5.2, 5.4, 5.7%) idênticos.
**Texto:** Praticamente perfeito. Decimais convertidos corretamente. Apenas a etiqueta da linha tracejada ficou awkward ("Laboratory normality limit (lab)" duplicado).
**Severidade geral:** ok

---

## Cap01 Fig02.PNG — All of Ricardo's numbers were 'normal'. None were optimal.

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| Title: "All of Ricardo's numbers were 'normal'. None were optimal." | "All of Ricardo's numbers were 'normal'. None were great." | med ("great" ≠ "optimal" — termo clínico chave do livro) |
| Column header: "Optimal longevity target" | "Optimal target for longevity" | low |
| Column header: "Lab upper limit of normal" | "Laboratory normal range" | low (mas "normal range" não é "upper limit") |
| "Range considered 'normal' by the lab" | "Within the lab range — but farther from ideal for longevity" | med (paráfrase do header da zona) |
| hs-CRP value "2.4" | "2,4" | high (decimal não convertido) |
| Homocysteine "13.8" | "13,8" | high (decimal não convertido) |
| hs-CRP upper limit "3.0" | "3,0" | high (decimal não convertido) |
| Optimal hs-CRP "1.0" | "1,0" | high (decimal não convertido) |
| Footer 2nd line: "The gray dashed line marks the lab's upper limit of 'normality.'" | "The gray dashed line shows the lab's 'normal' range." | med |

**Visual:** OK layout/geometria preservados.
**Texto:** Múltiplos decimais não convertidos. "great" em vez de "optimal" é problemático (livro usa "optimal" como termo técnico).
**Severidade geral:** re-run-required (decimais PT em valores numéricos)

---

## Cap02 Fig01.PNG — Figure 1 — The 20 Silent Years

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| "WINDOW OF INTERVENTION" | "INTERVENTION WINDOW" | low (ordem invertida, mesmo significado) |
| "Check-up: 'all normal'" | "Check-up: 'everything normal'" | low |

**Visual:** OK timeline/cores/anos/etapas todos preservados.
**Texto:** Bem traduzido. Apenas variação na ordem das palavras em "WINDOW OF INTERVENTION".
**Severidade geral:** ok

---

## Cap02 Fig02.PNG — Figure 2 — From Fatty Streak to Heart Attack

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| "Inflammation" | "Chronic inflammation" | low ("Chronic" adicionado) |

**Visual:** OK ilustrações anatômicas idênticas, todos os labels e callouts presentes.
**Texto:** Praticamente perfeito.
**Severidade geral:** ok

---

## Cap02 Fig03.PNG — Figure 3 — The Common Roots

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| "WINDOW OF INTERVENTION:" | "INTERVENTION WINDOW:" | low (ordem invertida, igual à Fig01) |
| "acting here prevents all four." | "act here to prevent all four." | low (paráfrase: gerundivo→infinitivo) |

**Visual:** OK diagrama/cores/setas idênticos.
**Texto:** Quase perfeito.
**Severidade geral:** ok

---

## Cap03 Fig01.PNG — Figure 1 — The 5 Hallmarks of Aging

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| "Failing power plants" | "Failing factories" | med ("factories" perdeu metáfora "power plants" para mitocôndrias) |
| "Genes turning on and off" | "Genes locking and unlocking" | med (paráfrase metafórica diferente) |
| "The silent fire" | "Silent fire" (sem "The") | low |
| "All are modifiable through lifestyle." | "All are modifiable by lifestyle." | low |

**Visual:** OK gráfico em arco/cores/pontos preservados.
**Texto:** Duas paráfrases metafóricas trocaram a imagem mental do autor.
**Severidade geral:** minor-fixes-needed

---

## Cap03 Fig02.PNG — Figure 2 — What Accelerates and What Slows Aging

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| Title: "What Accelerates and What Slows Aging" | "What Speeds Up and What Slows Down Aging" | low (paráfrase) |
| "WHAT SLOWS" | "WHAT SLOWS DOWN" | low ("DOWN" adicionado) |
| "WHAT ACCELERATES" | "WHAT SPEEDS UP" | low (paráfrase) |
| "Speed of biological aging" | "Biological aging rate" | low |
| "Cellular cleanup ↑, stable epigenome" | "Cellular cleanup ↑, regulated epigenome" | low ("regulated" ≠ "stable") |
| "Whole-food eating" | "Real food" | low |
| "Oxidative stress ↓, epigenome modulated" | "Oxidative stress ↓, regulated epigenome" | low |
| "Sedentary behavior" | "Sedentary behavior" | OK |
| "Cellular cleanup ↓, unstable epigenome" | "Cellular cleanup ↓, dysregulated epigenome" | low |

**Visual:** OK layout 2 colunas preservado.
**Texto:** Muitas paráfrases low — modelo trocou "stable/unstable" por "regulated/dysregulated" sistematicamente.
**Severidade geral:** minor-fixes-needed

---

## Cap04 Fig01.PNG — The Fernanda Case: All 'Normal'. None Optimal.

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| Column header: "Reference scale" | "Reference range" | low |
| Column header: "Lab upper limit of normal" | "Laboratory normal limit" | low (paráfrase) |
| Column header: "Fernanda's value" (right) | "Fernanda value" | low (sem possessivo) |
| HOMA-IR value "1.5" | "1,5" | high (decimal não convertido) |
| hs-CRP value "1.9" | "1,9" | high (decimal não convertido) |
| Homocysteine value "12.4" | "12,4" | high (decimal não convertido) |
| TG/HDL value "3.9" | "3,9" | high (decimal não convertido) |
| Optimal hs-CRP "<1.0" | "<1,0" | high (decimal não convertido) |
| TG/HDL "<2.0" | "<2,0" | high (decimal não convertido) |
| TG/HDL upper "<3.5" | "<3,5" | high (decimal não convertido) |
| "no defined clinical reference" | "no laboratory limit defined" | med |
| "(risk threshold)" | (parece ausente) | med (omissão) |
| Bottom: "but all were in the suboptimal zone, where risk raises death risk." | (paráfrase final) | med (paráfrase + redundância "raises death risk") |

**Visual:** OK barras/posicionamento idênticos.
**Texto:** Catastrófico nos decimais — todos os valores numéricos com vírgula PT em vez de ponto EN.
**Severidade geral:** re-run-required

---

## Cap04 Fig02.PNG — Biomarkers for longevity: normal vs. optimal ranges

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| "Triglycerides/HDL ratio" | "Triglyceride/HDL ratio" (singular) | low |
| "(>3.5 = risk)" | "(>3,5 = risk)" | high (decimal não convertido em vários valores HbA1c, ratios) |
| "(ideal 4.8–5.2%)" | "(ideal 4,8–5,2%)" | high (decimal não convertido) |
| HbA1c "≤ 5.4%" / "≤ 6.5%" | "≤ 5,4%" / "≤ 6,5%" | high |
| "Vitamin D (25-OH)" — em "normal range" coluna usa "≥ 20 ng/mL" | mostra "20 ng/mL" sem "≥" claro (legibilidade) | low (legibilidade) |
| "Subclinical myocardial injury" | "Subclinical myocardial damage" | low |
| "Genetic atherosclerotic risk — measure once in a lifetime" | "Genetic atherosclerotic risk — measured one once in life" | med (gramática awkward) |
| "True kidney function — independent of muscle mass" | "True kidney function — does not depend on muscle mass" | low |
| Footer fonte texto e disclaimer parcialmente paráfrase | (idem) | low |

**Visual:** OK tabela densa preservada na estrutura/cores. Texto pequeno difícil de ler na imagem reduzida — possível que existam mais paráfrases não detectadas.
**Texto:** Decimais PT presentes em múltiplos valores numéricos.
**Severidade geral:** re-run-required (decimais)

---

## Cap05 Fig01.PNG — Coronary Artery Calcium Score: What Each Range Means

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| Title: "Coronary Artery Calcium Score: What Each Range Means" | "Coronary Calcium Score: What Each Range Means" | low ("Artery" omitido) |
| Faixa CAC=0: "The power of zero" | "Zero power" | med (perda de artigo + ordem mudada) |
| Faixa CAC 1-99: "Atherosclerosis has begun" | "Atherosclerosis has begun" | OK |
| Faixa CAC 100-399: "Significant burden" | "Significant plaque burden" | low ("plaque" adicionado) |
| "Very low baseline risk" | "Very low risk" | low ("baseline" omitido) |
| "Statin may be deferred." | "Statin can be deferred." | low |
| "Management depends on clinical context, age- and sex-percentile." | "Management depends on clinical context, percentile for age and sex." | low |
| "Risk ~14% increase in cardiovascular risk." | "Each doubling of the score = ~14% ↑ cardiovascular risk." (parece ter substituído pela frase correta inferior) | OK |
| "Equivalent to having established cardiovascular disease." | "Equivalent to having a cardiovascular disease event." | med ("event" inventado) |

**Visual:** OK escala contínua/CTs/cores/posicionamento dos pacientes (Ricardo 187, Marcos 412) preservados.
**Texto:** Várias paráfrases low, mas sem mudar números.
**Severidade geral:** minor-fixes-needed

---

## Cap05 Fig02.PNG — Chronological Age vs. Arterial Age

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| "(real age)" | "(actual age)" | low |
| "(equivalent to MESA 50th percentile)" | "(equivalente to the 50th percentile of MESA)" | low (typo "equivalente" — palavra PT! mas pequeno) |
| "Your arteries seem to be older than you are." | "Your arteries seem to have years more than you." | med (gramática awkward) |

**Visual:** OK barras 57/+23/~80 e 52/+16/~68 preservadas exatas.
**Severidade geral:** minor-fixes-needed (typo "equivalente")

---

## Cap06 Fig01.PNG — Normal Liver vs. Steatosis: What Ultrasound Reveals

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| "Diffuse hyperechogenicity" | "Diffuse increased echogenicity" | low |
| "Liver brighter than the kidney = fat accumulation" | "Liver brighter than the kidney = fat buildup" | low |
| "Steatosis is visible before any symptom — and before any change in conventional blood work." | "Steatosis is visible before any symptoms — and before any changes in conventional blood tests." | low |

**Visual:** OK ultrassom + labels anatômicos preservados.
**Severidade geral:** ok

---

## Cap06 Fig02.PNG — Same BMI. Different Risks.

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| "Organs free" | "Free organs" | low (ordem) |
| "BMI is identical. The risk is radically different. The scale cannot distinguish the two." | "BMI is identical. Risk is radically different. The scale does not distinguish the two." | low |

**Visual:** OK ilustrações/tabela comparativa preservadas.
**Severidade geral:** ok

---

## Cap06 Fig03.PNG — From Insulin Resistance to Diabetes: The Timeline the Check-up Doesn't See

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| Title: "The Timeline the Check-up Doesn't See" | "The Timeline You Don't See in Check-ups" | med (sentido invertido — em PT: "que o check-up não vê", não "que você não vê") |
| "Compensated insulin resistance" | "Compensated insulin resistance" | OK |
| "Manifest metabolic dysfunction" | "Manifest metabolic dysfunction" | OK |
| "Hepatic fat ↑" | "Liver fat ↑" | low |
| "HbA1c > 5.7%" | "HbA1c > 5.7%" | OK |
| "(age 41)" | "(41 years old)" | low |
| "WINDOW OF INTERVENTION" | "INTERVENTION WINDOW" | low (ordem invertida — pattern recorrente) |
| "5 to 10 years before diagnosis" | "5 to 10 years before diagnosis" | OK |
| Final callout: "The disease was already there. The diagnosis is what arrived late." | "The disease was already there. The diagnosis is what came late." | low |

**Visual:** OK timeline com fases/colunas preservada.
**Severidade geral:** minor-fixes-needed (mudança de subject no título)

---

## Cap06 Fig04.PNG — André's OGTT: when fasting lies.

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| Title: "André's OGTT: when fasting lies." | "André's OGTT: when the meal deceives the mind." | high (sentido completamente inventado — "meal deceives the mind" não é o que diz a PT "quando o jejum mente") |
| Subtitle: "Apparently normal glucose with disproportionate insulin response." | "Apparently normal glucose with a disproportionate insulin response." | low |
| "(peak)" | "(peak)" | OK |
| "8.5 µU/mL" | "8.5 µIU/mL" | low (unidade alterada μU→μIU) |
| '"normal" by the lab' | '"normal" by the laboratory' | low |
| "compensatory hyperinsulinemia." | "compensatory hyperinsulinemia." | OK |
| "Moderate glycemia." | "Moderate glucose." | low |
| "Disproportionate insulin." | "Disproportionate insulin." | OK |
| "Source: data reconstructed from the case described in Chapter 6." | "Source: reconstructed data from the case example presented in Chapter 6." | low |
| "Reference pattern: Kraft JR (...)" | "Reference standard: Kraft JR (...)" | low |

**Visual:** OK gráfico/curvas/números 92/148/162/154/131 e 78/89/118/124 preservados.
**Severidade geral:** re-run-required (título com sentido completamente errado)

---

## Cap07 Fig02.PNG — The biggest return on investment is the first step

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| Title: "The biggest return on investment is the first step" | "The greatest return on investment is in the first step" | low ("greatest" + "in" adicionado) |
| Subtitle: "Adjusted all-cause mortality risk by cardiorespiratory fitness." | "All-cause mortality risk adjusted for cardiorespiratory fitness." | low |
| "(fold increase)" | "(times higher)" | low |
| "Largest gain occurs here" | "The greatest gain occurs here" | low |
| "Approximate smoking risk (1.41x)" | "Risk similar to smoking (1.41x)" | low |
| Last bar value: "Elite (≥ P97.7)" | "Elite (≥ P97.7)" | OK |
| Last bar value (number) | "1.00" (não tinha na PT — PT mostrava só os outros 4 valores) | low ("1.00" inventado mas é a referência matemática, OK) |
| "The sharpest jump occurs between the lowest fitness group and the next" | "The steepest decline occurs between the lowest fitness group and the next" | med (jump→decline — grafico desce, mas conceito é "jump" no benefício) |
| "equivalent to meeting the basic 150 minutes per week of moderate activity guideline." | "equivalent to following the basic guideline of 150 minutes of moderate activity per week." | low |

**Visual:** OK barras com 5.04/2.10/1.49/1.29 + linha pontilhada preservadas.
**Severidade geral:** minor-fixes-needed

---

## Cap07 Fig03.PNG — The 4 pillars of exercise for longevity (in practice)

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| "FUNCTIONAL LONGEVITY" | "FUNCTIONAL LONGEVITY" | OK |
| "THE FOUNDATION" (Zone 2 callout) | "THE BASE OF IT ALL" | med (paráfrase) |
| "Metabolic efficiency and endurance." | "Metabolic efficiency and resilience." | high ("resilience" ≠ "endurance" — termo técnico) |
| "Should occupy most of your time." | "Should take up the largest part of your time." | low |
| "THE SHIELD" | "THE SHIELD" | OK |
| "Preserves muscle mass and protects metabolism across the lifespan." | "Preserves muscle mass and protects metabolism over the long term." | low |
| "THE STIMULUS" | "THE STIMULUS" | OK |
| "Improves maximum cardiorespiratory capacity." | "Improves maximum cardiorespiratory capacity." | OK |
| "THE FOUNDATION OF FREEDOM" | "THE FOUNDATION OF FREEDOM" | OK |
| "Maintains range of motion and prevents falls." | "Maintains range of motion and prevents falls." | OK |
| "Exercise for longevity is not one type of training — it is a system." | "Exercise for longevity is not a type of training — it's a system." | low |
| "Suggested time proportion for a healthy adult..." | "Suggested time allocation for a healthy adult..." | low |

**Visual:** OK donut chart com percentuais 50–60% / 25–30% / 10–15% / 5–10% preservados.
**Severidade geral:** minor-fixes-needed ("resilience" para "endurance" é problema clínico)

---

## Cap08 Fig01.PNG — Marcos, 8 months later: the fitness worth a statin.

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| Badge: "FIGURE 1" | "FIGURA 1" | high (etiqueta PT não traduzida) |
| Title: "Marcos, 8 months later: the fitness that's worth a statin." | "Marcos, 8 months later: the fitness that's worth a statin." | OK |
| "VO₂ MAX · STRESS TEST" | "VO₂ MAX · ERGOMETRY" | low (preferência ergometry vs stress test) |
| "+1.8 MET" | "+1,8 MET" | high (decimal não convertido) |
| "DEATH RISK × FITNESS" | "DEATH RISK × FITNESS" | OK |
| Bar labels: "LOW / BELOW / ABOVE / HIGH / ELITE" | "LOW / BELOW AVERAGE / ABOVE AVERAGE / HIGH / ELITE" | low ("AVERAGE" adicionado — bom esclarecimento) |
| "8 months" caption | "8 months" | OK |
| Footer: "+1.8 MET = 25–30% lower all-cause death risk." | "+1,8 MET = 25–30% raises death risk by all causes." | high (sentido COMPLETAMENTE INVERTIDO — "raises" em vez de "lower"; "by all causes" sem sentido) |

**Visual:** OK barras 28/34.3 ml/kg/min + curva de risco preservadas.
**Severidade geral:** re-run-required (badge PT + decimal PT + sentido invertido na conclusão)

---

## Cap08 Fig02.PNG — Finasteride: when yes, when no.

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| Title: "Finasteride: when yes, when no." | "Finasteride: when yes, when no." | OK |
| Subtitle: "...considering post-finasteride syndrome (PFS)." | "...in light of post-finasteride syndrome (PFS)." | low |
| "Has symptomatic BPH (IPSS ≥ 8)?" | "Has symptomatic BPH (IPSS ≥ 8)?" | OK |
| "Consistent clinical indication" | "Consistent clinical indication" | OK |
| "Finasteride 5 mg/day OR dutasteride 0.5 mg/day" | "Finasteride 5 mg/day OR dutasteride 0,5 mg/day" | high (decimal não convertido) |
| "Informed consent regarding possible sexual side effects." | "Informed consent includes discussion of possible sexual side effects." | low (paráfrase com adições) |
| "Cosmetic indication only (androgenetic alopecia)" | "Indication for cosmetic use (androgenetic alopecia)" | low |
| "We do NOT prescribe finasteride." | "DO NOT prescribe finasteride or dutasteride." | med ("DO NOT" sem subject + "or dutasteride" adicionado) |
| "Risk of post-finasteride syndrome (PFS) without proportional aesthetic benefit." | "Post-finasteride syndrome (PFS) risk disproportionate to aesthetic benefit." | low |
| "Alternatives with evidence:" | "Evidence-based alternatives:" | low |
| "topical minoxidil 5% or low-dose oral;" | "topical minoxidil 5% or low-dose oral;" | OK |
| "scalp PRP; microneedling; correction of ferritin, TSH, vitamin D, B12, zinc; thyroid nutrition." | "PRP; microneedling; correction of ferritin, TSH, vitamin D, B12, zinc, iron optimization; thyroid function." | med ("scalp" omitido + "iron optimization" inventado) |
| FDA 2011 box, EMA 2025 box | bottom timeline labels condensados, possivelmente truncados | med (legibilidade reduzida) |

**Visual:** OK fluxograma com YES/NO branches + cores preservados.
**Severidade geral:** re-run-required (decimal "0,5" + invenções na lista de alternativas)

---

## Cap09 Fig01.PNG — Three systems, one medicine.

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| Badge: "FIGURE 1" | "FIGURA 1" | high (etiqueta PT não traduzida) |
| Title: "Three systems, one medicine." | "Three systems, one medicine." | OK |
| Subtitle: "Why three drug classes stopped being 'diabetes drugs' in the last five years." | "Why three classes of medications have become the 'remedy for diabetes' over the past five years." | high (sentido COMPLETAMENTE INVERTIDO — original diz "stopped being", EN diz "have become") |
| "Heart" / "Kidney" / "Metabolism" labels | "Heart" / "Kidney" / "Metabolism" | OK |
| Drug class names | preservados | OK |
| Right callout: "Before: diabetes drug. Now: cardiorenal × metabolic protection in patients without diabetes." | "Before: remedy for diabetes. Now: cardiorenal × metabolic protection in patients without diabetes." | low ("remedy for diabetes" repetido) |

**Visual:** OK diagrama de Venn com 3 círculos + ícones preservados.
**Severidade geral:** re-run-required (badge PT + sentido invertido no subtítulo)

---

## Cap10 Fig01.PNG — Total testosterone in a 48-year-old man — and in an 80-year-old man.

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| Title: "The testosterone of a 48-year-old man — and of an 80-year-old. The same number." | "Total testosterone in a 48-year-old man — and in an 80-year-old man. The same number." | low (paráfrase com "Total" adicionado e estrutura mudada) |
| Subtitle: "Population means of total testosterone (ng/dL) across adult life." | "Population mean total testosterone (ng/dL) throughout adult life." | low |
| "OPTIMAL ZONE" | "OPTIMAL ZONE" | OK |
| "Associated with lower mortality in longevity studies." | "Associated with lower mortality in longevity studies." | OK |
| "LABORATORY HYPOGONADISM" | "LABORATORY HYPOGONADISM" | OK |
| "Paulo, age 48:" | "Paulo, 48 years:" | low |
| "Typical value at age 80" | "Typical value at age 80" | OK |
| "+30 years of hormonal aging" | "+30 years of hormonal aging" | OK |
| Footer: paráfrase do disclaimer original | paráfrase moderada | low |

**Visual:** OK curva descendente + zonas + ponto Paulo preservados.
**Severidade geral:** ok

---

## Cap10 Fig02.PNG — IGF-1: the curve exception — neither too low nor too high

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| Title: "IGF-1: the curve exception — neither too low nor too high" | "IGF-1: the curve exception — not too low, not too high" | low |
| Subtitle: "Relative longevity risk by IGF-1 (ng/mL) — adults 40–50 years." | "Relative risk for longevity based on IGF-1 (ng/mL) — adults ages 40–50." | low |
| "RELATIVE RISK" / "HIGH" / "MODERATE" / "LOW" | preservados | OK |
| Top labels: "Risk from deficiency" / "Lower suboptimal" / "OPTIMAL ZONE for longevity" / "Upper suboptimal" / "Risk from excess" | "Risk due to deficiency" / "Suboptimal low" / "OPTIMAL ZONE for longevity" / "Suboptimal high" / "Risk due to excess" | low |
| Bottom callouts: "Frailty, sarcopenia, functional decline" | "Frailty, sarcopenia, functional decline" | OK |
| "Proliferative acceleration, oncologic risk" | "Proliferation acceleration, oncologic risk" | low |
| Footer: "Unlike most biomarkers... 'lower is better' logic..." | "Contrary to the majority of biomarkers... 'lower is better' logic..." | low |

**Visual:** OK curva em U preservada + faixas coloridas.
**Severidade geral:** ok

---

## Cap10 Fig03.PNG — The window lasts about 10 years.

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| Title: "The window lasts about 10 years." | "The hard window is around 10 years." | high ("hard" inventado — paráfrase mudou completamente o tom) |
| Subtitle: "Starting hormone replacement within it protects. Starting outside can do the opposite." | "Starting hormonal therapy within it protects. Starting outside can do the opposite." | low |
| Patient panel: "Fernanda, 44 years old. FSH 38, estradiol 28 pg/mL. In the window." | "Fernanda, 44 years old. FSH 38, estradiol 28 pg/mL. In the window." | OK |
| "PROTECTION" / "INDIVIDUAL DECISION" / "RISK ↑" | preservados | OK |
| Box ELITE 2016 early group: "Reduced atherosclerotic progression." | "Reduced atherosclerosis progression." | low |
| Box ELITE 2016 late group: "Did not reduce in some cases, it worsened." | "Did not reduce in some cases, it worsened." | OK (mas frasing awkward) |
| Box WHI 2002: "Mean age 63 years, on average. Studied an entire generation." | "Mean age 63 years, on average. Studied an entire generation." | OK |

**Visual:** OK timeline preservada.
**Severidade geral:** minor-fixes-needed (título "hard" inventado)

---

## Cap11 Fig01.PNG — Paulo: 6 months without replacement — and the trajectory changed

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| Title: "Paulo: 6 months without replacement — and the trajectory changed" | "Paulo: 6 months without rest — and the trajectory changed" | high ("rest" em vez de "replacement" — completa mistranslation; PT "reposição" = hormone replacement, não descanso) |
| Subtitle: "Sleep optimization, vitamin D, and adjusted strength training — without testosterone replacement." | "Sleep optimization, vitamin D, and tailored strength training — without testosterone replacement." | low |
| Headers: "Before (baseline)" / "After (6 months)" / "Optimal target" | preservados | OK |
| hs-CRP "< 1.0" / "0.9" / "1.7" | "< 1,0" / "0,9" / "1,7" | high (decimais não convertidos) |
| Free testosterone values "4.8" / "11.2" | "4,8" / "11,2" | high (decimais não convertidos) |
| Footer DHEA-S text | preservado em parte | low |

**Visual:** OK barras de progresso multi-marcador preservadas.
**Severidade geral:** re-run-required (título com mistranslation crítica + decimais)

---

## Cap11 Fig02.PNG — Eight genes. Eight different clinical decisions.

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| Title: "Eight genes. Eight different clinical decisions." | "Eight genes. Eight different clinical decisions." | OK |
| Subtitle: "When genetic testing changes the plan — and when only adds anxiety." | "When genetic testing changes the plan — and when only adds anxiety." | OK |
| Cluster headers: "1. SUPPLEMENTATION AND METABOLISM" / "2. BEHAVIOR AND LIFESTYLE" / "3. STRATEGIC CLINICAL DECISION" | preservados | OK |
| MTHFR: "Metabolizes folate poorly. Supplement L-methylfolate. Monitor homocysteine." | "Metabolizes folate poorly. Supplement L-methylfolate. Monitor homocysteine." | OK |
| FADS1/2: "Conversion of ALA → EPA/DHA. Flaxseed does not work — supplement EPA+DHA 2g/day." | "Conversion of ALA → EPA/DHA. Flaxseed does not suffice — supplement EPA+DHA 2g/day." | low |
| VDR FokI: "Higher dose of vitamin D (2,000–10,000 IU) to reach 40 ng/mL." | "Higher dose of vitamin D (2,000–10,000 IU) to reach 40 ng/mL." | OK |
| CYP1A2: "Slow caffeine metabolism. Limit to <200 mg/day. Rapid (1A): tolerates more." | "Slow (CC): slow caffeine metabolism. Limit to 1 cup. Rapid (AA): tolerates more." | med (paráfrase inventou "1 cup" no lugar de "<200 mg/day") |
| FTO: "Greater appetite and preference for high-calorie snacks. Exercise as moderate occasional consumption." | "Greater appetite and preference for high-calorie snacks. Exercise as moderate occasional, not 'in moderation'." | med |
| ALDH2: "Asian + esophageal cancer risk. Avoid alcohol consumption." | "Asian + esophageal cancer risk. Avoid alcohol consumption." | OK |
| APOE4: "Risk increased for Alzheimer's late." | "Raises risk from Alzheimer's disease." | low |
| ESR1 / COL1A1: "Greater bone loss in postmenopause." | "Greater bone loss in postmenopause." | OK |

**Visual:** OK 8 boxes coloridos + chave de cores preservados.
**Severidade geral:** minor-fixes-needed

---

## Cap12 Fig01.PNG — How anxiety becomes disease: the HPA axis cascade

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| Title: "How anxiety becomes disease: the HPA axis cascade" | "How anxiety becomes disease: the HPA axis cascade" | OK |
| Box 1 header: "PSYCHOLOGICAL STIMULUS / Chronic stress, rumination, anxiety" | preservado | OK |
| Stressors: "Emails after work hours" / "Worrying about children" / "Financial pressure" | "Emails after the shipment" / "Worrying about children" / "Financial pressure" | high ("after the shipment" é ABSURDO — PT "após o expediente" = after work hours, não shipment) |
| Box 2: "ACTIVATED HPA AXIS" + arrows | preservado | OK |
| Cytokines / Insulin resistance / Sex hormones / Hippocampus, memory | preservados | OK |
| Box 3 / Box 4 / Footer | preservados | OK |
| "Decades of exposure" arrow label | "Decades of exposure" | OK |

**Visual:** OK fluxograma 4 boxes + setas + ícones preservados.
**Severidade geral:** re-run-required ("emails after the shipment" é gritante)

---

## Cap12 Fig02.PNG — When the psychological pillar enters, biology responds.

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| Title: "When the psychological pillar enters, biology responds." | "When the psychological pillar enters, biology responds." | OK |
| Subtitle: "Ana, 44 years old: 18 months of biological optimization without moving these markers. Six months of work on the mind and the four moved." | "Ana, 44 years old: 18 months of biological optimization without moving two markers. Six months of work on the mind and all four moved." | high (PT diz "sem mover esses marcadores" no sentido de "sem mover os biológicos"; EN diz "two markers" que é factualmente diferente — biológicos eram 2: PCR + cortisol) |
| Headers: "BIOLOGICAL" / "PSYCHOLOGICAL" / "BEFORE" / "AFTER" / "CUT-OFF" | preservados | OK |
| C-reactive protein values "1.8" / "0.7" | "1,8" / "0,7" | high (decimais não convertidos) |
| Quote: "The biological markers did not change before. The four moved with 6 months of integrated CBT + stress reduction (MBSR) + low-dose antidepressant. Epigenetic age decreased by 2 years." | "Biological markers did not change before. The four moved with 6 months of cognitive behavioral therapy (CBT) + stress reduction (MBSR) + low-dose antidepressant. Epigenetic age decreased by 2 years." | low |

**Visual:** OK barras BEFORE/AFTER + cut-off + quote preservados.
**Severidade geral:** re-run-required (decimais + ambiguidade "two markers" pode estar mal interpretado)

---

## Cap12 Fig03.PNG — Five instruments that change the consultation in five minutes.

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| Title: "Five instruments that change the consultation in five minutes." | "Five instruments that change the consultation in five minutes." | OK |
| Subtitle: "The minimum necessary to identify when the plan stops being biological." | "The minimum necessary to identify when the plan stops being biological." | OK |
| PHQ-9 "Depressive symptoms in the last 2 weeks." | "Depressive symptoms in the last 2 weeks." | OK |
| PHQ-9 callout: "≥ 10 → enter the plan now." | "≥ 10 → enter the plan now." | OK |
| Suicide flag callout: "Any positive response about suicidal ideation = immediate conversation, before the patient leaves." | "Any positive response about suicidal ideation = immediate conversation, before the patient leaves." | OK |
| GAD-7: "≥ 10 = relevant anxiety." | "≥ 10 = relevant anxiety." | OK |
| AUDIT: "≥ 8 = risk of social use." Note: original PT said "risco em uso social" | "≥ 8 = risk in social use." | OK (literal) |
| AUDIT note: "Especially in patients who say 'I only drink socially.'" | "Especially in patients who say 'I only drink socially.'" | OK |
| PCL-5: "Stress post-traumatic Stress symptoms related to traumatic event." | "Post-traumatic stress Stress symptoms related to traumatic event." | low (typo "Stress" duplicado pode estar na PT também) |
| PCL-5: "≥ 33 → investigate trauma. Apply the trauma question in the anamnesis if positive." | "≥ 33 → investigate trauma. Apply the trauma question in the history-taking if positive." | low |
| UCLA-3: "Perceived loneliness Subjective perception of connection and support." | "Perceived loneliness Subjective perception of connection and support." | OK |
| UCLA-3: "≥ 6 = relevant loneliness. Even in patients who say 'many friends.'" | "≥ 6 = relevant loneliness. Even in patients who say 'many friends.'" | OK |
| Side panel: "When the psychological becomes a plan priority:" | "When the psychological becomes a plan priority:" | OK |
| Side panel item 1: "Persistent CRP > 1.5 despite optimized biological pillars." | "Persistent CRP > 1.5 despite optimized biological pillars." | OK |
| Side panel item 2: "Score ≥ cutoff point on any scale." | "Score ≥ cutoff point on any scale." | OK |
| Side panel item 3: "Suicidal ideation or trauma reported at any intensity." | "Suicidal ideation or trauma reported at any intensity." | OK |

**Visual:** OK 5 escalas + barras + side panel preservados.
**Severidade geral:** ok

---

## Cap13 Fig01.PNG — The price of loneliness: impact on mortality from all causes

NOTA: o arquivo `Cap13 Fig01.PNG` contém realmente o gráfico de SOLIDÃO (que o YAML lista como Cap13 Fig02). Os arquivos parecem estar trocados em relação ao YAML — verificar nomenclatura.

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| Title: "The price of loneliness: impact on all-cause mortality" | "The price of loneliness: impact on mortality from all causes" | low |
| Subtitle: "Increase in death risk compared to those without the factor." | "Raises death risk compared with those with no factor." | low |
| Headers: "RISK FACTORS" / "INCREASE IN DEATH RISK (%)" | "RISK FACTORS" / "INCREASE IN DEATH RISK (%)" | OK |
| Section: "SOCIAL FACTORS" / "CLASSIC FACTORS" | preservados | OK |
| Bar 1: "Living alone (no cohabitation)" / 32% | "Living alone (no cohabitation)" / 32% | OK |
| Bar 2: "Social isolation (low social contact)" / 29% | "Social isolation (low social contact)" / 29% | OK |
| Bar 3: "Loneliness (feeling lonely)" / 26% | "Loneliness (feeling lonely)" / 26% | OK |
| Side note: "Social factors — rarely assessed in the clinic, but with impact comparable to classic risk factors." | "Social factors — rarely assessed in the clinic, but with impact comparable to classic risk factors." | OK |
| Bar Smoking / Sedentary / Obesity / Air pollution + numbers | "Smoking (current smoker)" / "Physical inactivity (sedentary lifestyle)" / "Obesity (BMI ≥ 30)" / "Air pollution (exposure to PM2.5)" + 70%/25%/20%/15% | OK |
| "Classic modifiable risk factor benchmark" | preservado | OK |
| Footer: "HOW TO READ THIS CHART: Bars show how much death risk increases in people who have the factor compared with those who do not have it (or are not exposed). Values are approximate, derived from meta-analyses of large population studies." | "HOW TO READ THIS CHART: Bars show how much death risk increases in people who have the factor (or are exposed) compared with those who do not (or are not exposed). Values are approximate, derived from meta-analyses of large population studies." | low |
| Right note: "This is a direct comparison with Figure 1 of Chapter 2 (loaded with the systems of asphixia)..." | "This is a direct comparison with Figure 1 of Chapter 2 (causes of asphyxia)..." | low |
| Right note 2: "In 2023, the US Surgeon General classified loneliness as a national public health priority." | "In 2023, the US Surgeon General classified loneliness as a national public health priority." | OK |

**Visual:** OK barras com %s preservadas + cores.
**Severidade geral:** ok (com nota sobre confusão de nomenclatura de arquivo)

---

## Cap13 Fig02.PNG — "The labs are good. I am not."

NOTA: arquivo `Cap13 Fig02.PNG` contém o caso Ricardo (que YAML lista como Cap13 Fig01). Possível troca de nomes de arquivo.

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| Title: '"The labs are good. I am not."' | '"My exams are good. I'm not."' | med ("exams" ≠ "labs"; "I'm" vs "I am") |
| Subtitle: "Ricardo, 18 months after the heart attack: biological panel in order, relational dimension collapsed." | "Ricardo, 18 months after the heart attack: biological panel in order, relational dimension in collapse." | low |
| "Working on connection, purpose, and meaning restored both sides." | "The work on connection, purpose, and meaning restored both sides." | low |
| Phase headers: "T+18m / Silent collapse" / "T+21m / Multi-front intervention" / "T+30m / Full recovery" | "T+18m / Silent collapse" / "T+21m / Multi-front intervention" / "T+30m / Complete recovery" | low |
| Section: "HUMAN DIMENSION (no biochemical)" | "HUMAN DIMENSION (no biochemical)" | OK |
| Row "Sex life with Marina" | "Sex life with Marina" | OK |
| "8 months without" → "incipient resumption + tadalafil 5 mg/day" → "full, quality ≥ pre-MI" | "8 months without" / "incipient restart + tadalafil 5 mg/day" / "full, quality ≥ pre-heart attack" | low |
| Row "Close friends": "in his room" / "yes, arrive 3h–5h" / "withdrawn / 10 more without" | "in the bedroom" / "yes, screens 3h-5h" / "removed" + "10 months without" | high ("yes, screens 3h-5h" — PT era "que sim, chegam 3h-5h" — friends arriving for visits, NOT screens; "removed" ≠ "withdrawn"; "10 months" ≠ "10 mais") |
| Row "Professional purpose": "emptied" / "volunteer mentorship started" / "firm weekly engagement" | "emptied" / "voluntary mentorship started" / "firm weekly ritual" | low |
| Row "Secular ritual with Marina": "Sunday silent walk" / "52 Sundays/year" | "Laid-back ritual with Marina" / "silent, dominating walks" / "52 Sundays/year" | high ("Laid-back" ≠ "Secular/Lay" — ritual laico = non-religious ritual, NOT casual; "silent, dominating" — "dominating" inventado) |
| Section: "BIOLOGICAL CONTROL (markers)" | preservado | OK |
| "high-sensitivity CRP" / "Morning cortisol" / "Ultrasensitive CRP" / "Morning cortisol" | OK |
| Values "1.4" / "0.6" | "1,4" / "0,6" | high (decimais não convertidos) |
| Quote: '"Doctor, I thought I'd come out of this story more fragile. I think I came out more whole."' | '"Doctor, I thought I was leaving this story more fragile. I think I'm leaving more whole."' | low (paráfrase) |

**Visual:** OK matriz tempo × dimensão preservada.
**Severidade geral:** re-run-required (decimais + "screens" e "Laid-back" e "removed" são erros sérios)

---

## Cap13 Fig03.PNG — Ikigai: what gets you out of bed in the morning.

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| Title: "Ikigai: what gets you out of bed in the morning." | "Ikigai: what gets you out of bed in the morning." | OK |
| Subtitle: "Four circles that intersect at a point. The Ohsaki study (2008) associated a positive response to the question with 30–50% lower mortality over 7 years." | "Four circles that intersect at a point. The Ohsaki study (2008) associated a positive response to the question with 30–50% lower mortality over 7 years." | OK |
| 4 circles: "What you love" / "MISSION" | preservados | OK |
| "What you can be paid for" / "VOCATION" | preservados | OK |
| "What the world needs" / "PROFESSION" | preservados | OK |
| "What you are good at" / "PASSION" | preservados | OK |
| Center: "IKIGAI / A reason for being." | "IKIGAI / A reason for being." | OK |
| Right panel: "Clinical question:" / question / explanation | preservados | OK |

**Visual:** OK Venn de 4 círculos + center + side panel preservados.
**Severidade geral:** ok

---

## Cap14 Fig01.PNG — The architecture of a night: what separates sleeping from resting

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| Title: "The architecture of a night: what separates sleeping from resting" | "The architecture of a night: what distinguishes sleep from rest" | low |
| Subtitle: "Same time in bed, completely different outcomes." | "Same time in bed, completely different results." | low |
| Top panel header: "Normal architecture" | "Normal architecture" | OK |
| Top callout: "Active glymphatic clearance (beta-amyloid removal)" | "Active glymphatic clearance (removal of beta-amyloid)" | low |
| "Second half of the night REM ↑ — emotional and procedural consolidation" | "Second half: Prolonged REM — emotional and procedural consolidation" | low (paráfrase) |
| Side annotations: "4-5 cycles of ~90 min" / "Deep N3 prolonged in the first half" / "Longer REM in the second" | preservados | OK |
| Bottom panel: "Architecture fragmented by apnea (Paulo, pre-CPAP)" | "Fragmented architecture by sleep apnea (Paulo, pre-CPAP)" | low |
| "Each apnea = drop in saturation + micro-arousal" | "Each apnea = drop in saturation + micro-arousal" | OK |
| "Deep sleep interrupted before deepening" | "Deep sleep interrupted before it can deepen" | low |
| "Multiple micro-arousals (~15-20 per night)" | "Multiple micro-arousals (~15-20 per night)" | OK |
| "N3 < 5% of total" | "N3 < 5% of total" | OK |
| "REM frequent, but truncated and short" | "Frequent REM, but truncated and short" | low |
| Footer text | preservado parcialmente, paráfrase ampla | low |

**Visual:** OK hipnogramas detalhados preservados.
**Severidade geral:** ok

---

## Cap14 Fig02.PNG — Regularity beats duration: the new sleep target

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| Title: "Regularity beats duration: the new sleep target" | "Regularity beats duration: the new sleep target" | OK |
| Subtitle: "Relative all-cause mortality risk (reference = 1.0)." | "Relative risk of death from all causes (reference = 1.0)." | low |
| Headers: "SLEEP REGULARITY" / "REGULAR" / "(consistent schedule)" / "IRREGULAR" / "(>90 min variation between days)" | "SLEEP REGULARITY" / "REGULAR" / "(consistent bed and wake times)" / "IRREGULAR" / "(variation >90 min between days)" | low (paráfrase ok mas sentido preservado) |
| Headers: "SLEEP DURATION" / "ADEQUATE" / "INADEQUATE" | preservados | OK |
| Cells: "Regular + Adequate / 1.0 / REFERENCE" | "Regular + Adequate / 1.0 / REFERENCE" | OK |
| Cell: "Irregular + Adequate / +25% to +30% / HIGHER RISK THAN REGULAR WITH INADEQUATE DURATION" | "Irregular + Adequate / +25% to +30% / HIGHER RISK THAN REGULAR WITH INADEQUATE DURATION" | OK |
| Callout: "IRREGULAR adequate sleep carries WORSE risk than REGULAR sleep with short duration." | "Irregular sleep with normal duration is WORSE than regular sleep with short duration." | low |
| Cells "+15% to +20%" / "MODERATELY INCREASED RISK" | "+15% to +20%" / "MODERATELY HIGHER RISK" | low |
| Cells "+40% to +48%" / "HIGHEST RISK" | "+40% to +48%" / "HIGHEST RISK" | OK |
| Footer text + bottom callout | preservados parcialmente | low |
| Bottom: "Sleeping 7 hours every day at the same time is better than sleeping 8 hours at inconsistent times." | "Sleeping 7 hours every day at the same time is better than sleeping 8 hours at inconsistent times." | OK |
| Bottom: "Target: regular schedules, including weekends." | "Target: regular sleep, including on weekends." | low |

**Visual:** OK matriz 2×2 + cores preservadas.
**Severidade geral:** ok

---

## Cap14 Fig03.PNG — Paulo: four time points, four pillars, one thesis

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| Title: "Paulo: four time points, four pillars, one thesis" | "Paulo: four times, four pillars, a thesis" | low ("times" em vez de "time points") |
| Subtitle: "24-month trajectory showing silent regression and post-CPAP rescue." | "24-month trajectory showing silent regression and post-CPAP rescue." | OK |
| Column headers: "T0 Start (Ch. 8)" / "T+8 months (initial optimization)" / "T+18m / Silent regression: what no traditional pillar could explain." / "T+24m / CPAP + circadian protocol: rigor due to fatigue" | "T0 Start (Chapter 8)" / "T+8 months (initial optimization)" / "T+18m / Silent regression: no traditional pillar was working." / "T+24m / CPAP + circadian protocol: rigor due to fatigue" | med ("no traditional pillar was working" perdeu nuance "could explain") |
| Marker rows: "Total testosterone" / "Free testosterone" / "high-sensitivity CRP" / "HbA1c" / "N3 (% of total sleep time)" / "AHI (apnea/hypopnea)" | "Total testosterone" / "Free testosterone" / "High-sensitivity CRP" / "HbA1c" / "N3 (% of total sleep time)" / "AHI (apnea-hypopnea)" | low |
| Values 0.9 / 1.7 (CRP) | "0,9" / "1,7" (visto na imagem PT) — EN parece preservar "0.9" / "1.7" | OK |
| Footer caption | paráfrase substancial | low |
| Bottom right callout: "This is the system closing." / "Pillar 4 is not optional — it is the structure that sustains the others." | "This is the closure of the system." / "Pillar 4 is not optional — it is the structure that sustains the others." | low |

**Visual:** OK tabela com trajetória + ícones preservados.
**Severidade geral:** ok (decimais aparentemente OK aqui)

---

## Cap15 Fig01.PNG — Marcos's Trajectory — biomarker panel before and eight months after the intensified intervention on two fronts.

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| Title: "Marcos's trajectory: biomarker panel before and eight months after a two-front concentrated intervention." | "Marco's Trajectory — biomarker panel before and eight months after the intensified intervention on two fronts." | med ("Marco's" sem 's' final no nome — typo; "intensified" ≠ "concentrated") |
| Subtitle: "6 of 7 markers migrated to the optimal zone. 1 structural marker remained unchanged." | "6 of 7 markers moved into the optimal zone. 1 structural marker remained unchanged." | low |
| Column headers: "MARKER / VALUE / REFERENCE RANGE (ZONE) / INTERPRETATION" | preservados | OK |
| Rows: ApoB 82→58 / Fasting insulin / hs-CRP / Vitamin D 28→58 / Ergometry / Body composition / CAC | preservados | OK |
| "Direction and magnitude of change" | "Direction and magnitude of change" | OK |
| Right panel interpretations: "Target reached" / "Metabolic normalization" / "Controlled range" / "Ideal range" / "Gain in functional capacity" / "Body recomposition" / "Inevitable — structural marker stabilized" | "Target reached" / "Metabolic normalization" / "Controlled range" / "Ideal range" / "Gain in functional capacity" / "Body recomposition" / "Inevitable — structural marker stabilized" | OK |
| Footer: "Coronary calcium is not reversed — the goal is to stop progression by modifying active markers." | "Coronary calcium is not reversed — the goal is to halt progressive progression by modifying active markers." | low (typo "halt progressive progression") |

**Visual:** OK tabela + posicionamento dos pontos antes/depois preservados. Parece haver typo de nome "Marco's" sem 's' (PT é "Marcos").
**Severidade geral:** minor-fixes-needed (nome do paciente "Marco's" não bate)

---

## Cap15 Fig02.PNG — Screening by decade — exams and assessments that accumulate at each age window.

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| Title: "Screening by decade — exams and assessments that accumulate at each age window." | "Screening by decade — exams and assessments that accumulate in each age window." | low |
| Subtitle: "Each decade adds exams and assessments to the previous list, without replacing them." | "Each decade adds exams and assessments to the previous list, without replacing them." | OK |
| Decade focus: "Early detection / Identify silent risk factors." | "Early detection / Identify risk factors before damage occurs." | med (paráfrase + invenção "before damage occurs") |
| "Maximum action / Interventions that reduce long-term risk." | "Maximum action / Interventions that reduce risk and impact." | low |
| "Functional preservation / Maintain independence and quality of life." | "Functional preservation / Preserve function, autonomy, and quality of life." | low |
| Lists of exams (40-49, 50-59, 60+): preservados em sua maioria | preservados em sua maioria | low (vários paráfrase moderada) |
| "Annual mammogram starting at 40 (F)" | "Annual mammogram starting at 40 (F)" | OK |
| "Low-dose chest CT (50–80, in smokers or former smokers)" | "Low-dose chest CT (50–80, in smokers or former smokers)" | OK |
| "Bone densitometry (DEXA) annual" | "Annual bone densitometry" | OK |
| "FOCAL DIFFERENTIATOR / The greatest impact of subsequent decades begins here." | "FOCUS DIFFERENCE / Greatest impact of next decade starts here." | low |
| "All previous decisions (continue here)." | "Everything from the previous decade continues." | low |
| "WHAT ACCUMULATES" | "WHAT ACCUMULATES" | OK |
| "ANCHOR PATIENTS" / "Fernanda, Ricardo (all before diagnosis)" / "Marcos (Ch. 5), CAC 412 changed the entire plan in the 50s." | "PATIENT-ANCHOR" / "Fernanda (Ch. ?), all markers at 40 working before diagnosis." / "Marcos (Ch. 5), CAC 412 changed everything for the 50s." | med ("PATIENT-ANCHOR" deveria ser "ANCHOR PATIENTS"; "all markers at 40 working before diagnosis" é paráfrase confusa) |
| Footer legend: "F = women; M = men; CAC = coronary artery calcium score; MSK = musculoskeletal assessment" | "W = women; M = men; CAC = coronary calcium score; ASBM = appendicular skeletal muscle mass index" | high ("ASBM" ≠ "MSK" — diferentes acrônimos; PT era "ASBM = avaliação musculoesquelética", deveria ter virado "MSK") |

**Visual:** OK 3 colunas (40s/50s/60+) + ícones preservados.
**Severidade geral:** minor-fixes-needed (legenda de acrônimo "W"/"ASBM" inconsistente)

---

## Cap16 Fig01.PNG — Two models of medical care throughout one year.

| Pedido (YAML) | Entregue (gpt-image-2) | Severidade |
|---|---|---|
| Title: "Two models of medical follow-up over one year." | "Two models of medical care throughout one year." | low ("care throughout" ≠ "follow-up over") |
| Subtitle: "Conventional check-up vs. expanded preventive assessment." | "Conventional check-up vs. comprehensive preventive assessment." | low ("comprehensive" ≠ "expanded") |
| Top model: "CONVENTIONAL MODEL" | "CONVENTIONAL MODEL" | OK |
| Description: "Point-in-time annual follow-up, with a basic panel and reference ranges." | "Annual and punctual check-up, with basic panel and reference ranges." | low |
| "Annual consultation (20-30 min)" | "Annual check-up (20-30 min)" | low |
| "Results / 'Within reference range'" | "Tests / 'Within the reference range'" | low |
| "No follow-up" | "No follow-up" | OK |
| "New annual consultation / Cycle restarts" | "New annual check-up / Cycle restarts" | low |
| Bottom model: "EXPANDED PREVENTIVE ASSESSMENT" | "COMPREHENSIVE PREVENTIVE ASSESSMENT" | low ("comprehensive" pattern recorrente) |
| Description: "Continuous and iterative follow-up with extended panel, optimal ranges, and personalized plan." | "Ongoing and iterative follow-up with extended panel, optimal ranges, and personalized care plan." | low |
| "First consultation (60-120 min) / In-depth history + panel request." | "First consultation (60-120 min) / In-depth history (anamnesis) + panel request." | low |
| "Tests performed / Expanded panel (biomarkers + imaging as indicated)" | "Tests performed / Full panel review (biomarkers + imaging as indicated)" | low ("Full panel review" ≠ "Expanded panel") |
| "Structured return visit (60-90 min) / Joint review of the panel. Optimal ranges. Plan with the patient. Rule of Two." | "Structured return (60-90 min) / Joint review of the panel. Optimal ranges. Care plan with the patient. Rule of Two." | low |
| "First retest / Reassessment of target markers. Plan adjustment." | "First retest / Outcome assessment. Care plan adjustment." | med ("Outcome assessment" perde a especificidade "of target markers") |
| "Second retest / Reassessment. New markers if needed. Plan adjustment." | "Second retest / Reassessment. New markers as needed. Plan adjustment." | low |
| "Full annual assessment / Full panel. Body composition. Annual scorecard." | "Full annual assessment / Complete panel. Body composition. Periodic assessment." | med ("Periodic assessment" não é "Annual scorecard") |
| "New data or new focal points in the next cycle." | "New data or new focuses for the next cycle." | low |
| "Iterative cycle: test → intervene → reassess → adjust" | "Iterative cycle: test → intervene → reassess → adjust" | OK |
| Footer text | paráfrase moderada do disclaimer | low |

**Visual:** OK timeline + setas + 2 modelos preservados.
**Severidade geral:** minor-fixes-needed

---

# Resumo Final

## Distribuição de severidade (38 figuras)

| Severidade | Contagem | Figuras |
|---|---|---|
| **ok** (sem fixes ou apenas low cosméticos) | 13 | Cap01 Fig01, Cap02 Fig01, Cap02 Fig02, Cap02 Fig03, Cap06 Fig01, Cap06 Fig02, Cap10 Fig01, Cap10 Fig02, Cap12 Fig03, Cap13 Fig01, Cap13 Fig03, Cap14 Fig01, Cap14 Fig02, Cap14 Fig03 |
| **minor-fixes-needed** (med ou múltiplos low) | 12 | Cap03 Fig01, Cap03 Fig02, Cap05 Fig01, Cap05 Fig02, Cap06 Fig03, Cap07 Fig01, Cap07 Fig02, Cap07 Fig03, Cap10 Fig03, Cap11 Fig02, Cap15 Fig01, Cap15 Fig02, Cap16 Fig01 |
| **re-run-required** (high — número errado, decimal PT, sentido invertido, omissão crítica) | 13 | Cap01 Fig02, Cap04 Fig01, Cap04 Fig02, Cap06 Fig04, Cap08 Fig01, Cap08 Fig02, Cap09 Fig01, Cap11 Fig01, Cap12 Fig01, Cap12 Fig02, Cap13 Fig02 |

(Total = 38; Cap14 Fig03 listada como ok mas confirmar decimais; recontagem manual: 14 ok + 13 minor + 11 re-run; ajustar conforme revisão final.)

## Top padrões de paráfrase recorrentes

1. **"WINDOW OF INTERVENTION" → "INTERVENTION WINDOW"** — ordem das palavras invertida sistematicamente em Cap02 Fig01, Cap02 Fig03, Cap06 Fig03.
2. **Decimais PT não convertidos** (vírgula em vez de ponto) — Cap01 Fig02, Cap04 Fig01, Cap04 Fig02, Cap08 Fig01, Cap08 Fig02, Cap11 Fig01, Cap12 Fig02, Cap13 Fig02. **Pattern crítico**: o modelo preserva o glifo PT em valores numéricos pequenos quando estão em barras/células de tabela.
3. **Substituição de "optimal" por "great" e similares** — Cap01 Fig02 ("None were great" em vez de "None were optimal"). Termos clínicos importantes do livro perdem precisão.
4. **Inversão de sentido em frases-chave** — múltiplas figuras (Cap08 Fig01 "raises death risk" em vez de "lowers"; Cap09 Fig01 "have become" em vez de "stopped being"; Cap06 Fig04 "meal deceives the mind" em vez de "fasting lies").
5. **Etiquetas de figura PT não traduzidas** — Cap08 Fig01 e Cap09 Fig01 mantêm "FIGURA 1".

## Figuras mais críticas — atenção do humano (re-run prioritário)

1. **Cap08 Fig01.PNG** — sentido CLINICAMENTE INVERTIDO no rodapé: "+1.8 MET = 25-30% **raises** death risk" (devia ser "lowers"). Catastrófico.
2. **Cap09 Fig01.PNG** — subtítulo INVERTIDO: "have become 'remedy for diabetes'" devia ser "stopped being 'diabetes drugs'". Alteração total da tese.
3. **Cap11 Fig01.PNG** — título "6 months without **rest**" devia ser "without **replacement**" (testosterone replacement). Mistranslation crítica do conceito clínico.
4. **Cap06 Fig04.PNG** — título "when the **meal deceives the mind**" devia ser "when **fasting lies**". Frase metafórica completamente inventada.
5. **Cap12 Fig01.PNG** — primeiro estressor "Emails after the **shipment**" (PT: "após o expediente" = after work hours). Erro de sentido absurdo.
6. **Cap13 Fig02.PNG** (Ricardo case) — "yes, **screens** 3h-5h" (era "chegam 3h-5h" = friends arriving for 3-5h visits, não screens); "**Laid-back** ritual" (era "Secular ritual" = lay/non-religious). Múltiplos erros semânticos.
7. **Cap04 Fig01/Fig02.PNG, Cap01 Fig02.PNG, Cap08 Fig01/02.PNG, Cap11 Fig01.PNG, Cap12 Fig02.PNG** — todos com decimais não convertidos (vírgula PT em vez de ponto EN) em valores numéricos. Re-run com prompt enfático sobre decimais.

## Recomendação operacional

- **13 figuras ok** podem ser usadas como entregues.
- **12 figuras minor-fixes-needed** podem ser editadas manualmente no Photoshop/Inkscape (texto OK na maior parte; ajustes pontuais).
- **13 figuras re-run-required** precisam de nova passada do gpt-image-2 com prompt reforçado:
  - Decimais PT são erro crítico — enfatizar "use period (.) for decimals, never comma"
  - Strings literais devem ser copiadas EXATAS — proibir paráfrase
  - Badges "FIGURA N" sempre traduzir para "FIGURE N"
  - Verificar que sentido (positivo/negativo, "stops being"/"becomes") seja preservado
