# Refactor do Score de Sono — adicionar medidas objetivas

**Status:** ✅ núcleo O1–O4 APLICADO no dev (2026-06-15) · **Criado:** 2026-06-15 · **Origem:** curadoria de exames #6–#8 (sono).

## ✅ Aplicado (dev)
Subgrupo novo **`Sono › Medidas objetivas`** + 4 definições-achado não-requisitáveis numéricas + 4 score items + níveis + pilares:
| item | code | pts | níveis (mult) | pilares |
|---|---|---|---|---|
| **IAH** | PLNSLEEPIAH | 18 | `<5`=1 · `5–15`=.67 · `15–30`=.33 · `≥30`=0 (AASM) | Distúrbios do Sono · Cardiovascular |
| **SRI** (regularidade) | PLNSLEEPSRI | 18 | `≥87`=1 · `78–87`=.67 · `70–78`=.33 · `<70`=0 (quintis UK Biobank) | Cronobiologia e Ritmo · Distúrbios do Sono |
| **TST** (duração objetiva) | PLNSLEEPTST | 16 | U: `7–8h`=1 · `6–7`=.8 · `8–9`=.6 · `5–6`=.4 · `<5`=.2 · `>9`=0 | Distúrbios do Sono · Vitalidade e Disposição |
| **Eficiência** | PLNSLEEPEFF | 12 | `≥90`=1 · `85–90`=.67 · `80–85`=.33 · `<80`=0 | Distúrbios do Sono |

Defaults aplicados do "manda ver": sobreposição **(a)** (subjetivo intacto), **O5 não incluído**, **minúcia §5 não enxugada**. Decisões §4/§5 e O5 seguem em aberto p/ ajuste. Fecha #6 (PSG→IAH), #7 (HSAT→IAH), #8 (Actigrafia→SRI/TST/Eficiência). Exames PSG/HSAT/Actigrafia seguem requisitáveis (laudo); achados entram manualmente ou por wearable.


## Diagnóstico do estado atual
O score de sono hoje é **grande e quase 100% subjetivo** (anamnese). Levantamento do banco (dev):

- **Grupo `Sono`** (~50 itens, subgrupos Atual/Histórico): mistura itens clínicos relevantes com **muita minúcia de ambiente/higiene** de baixo valor.
  - Clínicos de peso: `Uso de BiPAP/CPAP` 22 · `Tempo de sono` 20 · `Apneias` 18 · `Hora de dormir` 18 · `Roncos` 14 · `Qualidade percebida` 14 · `Quarto` 14 · `Regularidade no acordar` 12 · `Frequência` 12 · `Insônia` 12 · `Sudorese noturna` 12 · `Tempo tela noturna` 12 · `Duração/Motivo interrupções` 12+12.
  - Minúcia ambiental (2–4 pts cada, ~20 itens): Pijamas 2 · Travesseiros 3 · Lençóis 3 · Colchão 4 · Cama 4 · Odores 4 · Janelas 4 · Temperatura 6 · Iluminação 8 · Barulho 8 · Telas 8 · Dieta noturna 8 · etc.
- **`Escala de Epworth`** (14 pts) — está no grupo **Cognição**, não em Sono.
- **Genética circadiana** (~14 itens × 1 pt): CLOCK, CRY1/2, PER2/3, MTNR1A/B — subgrupo Genética › Circadiano & Sono.
- **Zero medida objetiva.** Os 3 exames de sono (PSG/HSAT/Actigrafia) são requisitáveis mas `result_type='text'` (laudo livre) e **não pontuam**.

**Conclusão:** falta a camada objetiva, e o peso subjetivo está inflado por itens de higiene de quarto. O refactor **adiciona objetivo** e (opcional) **enxuga** a minúcia.

## Princípio do refactor
1. **Camada objetiva nova** = "achados" estruturados (definições **não-requisitáveis**, numéricas), no padrão das componentes da curva glicêmica (`GLICOSE 0 MIN`). Cada achado é alimentado pelo exame **ou wearable** que o produz; o exame (PSG/HSAT/Actigrafia) continua só pra **pedir**.
2. **Só métricas validadas e ligadas a desfecho.** Nada de % de estágio (REM/profundo) de wearable — validade 50–86%, ruído alto (meta-análises 2024). Sono/vigília, duração, eficiência e regularidade de wearable são aceitáveis; estágio não.
3. **Subjetivo continua** (todos respondem anamnese); objetivo **refina** quando existir. Tratar sobreposição (ver §4).

## Itens objetivos propostos (núcleo)
> Cutoffs marcados como **(confirmar)** precisam de 1 verificação online dirigida antes de implementar (não chutar).

| # | achado | unidade | fonte | cutoffs | pilares | pts |
|---|---|---|---|---|---|---|
| O1 | **IAH (índice apneia-hipopneia)** | eventos/h | PSG, HSAT | AASM: `<5` normal · `5–14` leve · `15–29` moderado · `≥30` grave (menor=melhor) | Distúrbios do Sono · Cardiovascular | 18 |
| O2 | **Índice de regularidade do sono (SRI)** | 0–100 | Actigrafia, wearable | bandas por quintil UK Biobank **(confirmar nº exatos)**; maior=melhor | Cronobiologia e Ritmo · Distúrbios do Sono | 18–20 |
| O3 | **Tempo total de sono objetivo (TST)** | h | Actigrafia, wearable | curva **U**, ótimo ~7–8h; `<5`/`>9` ruins (longo penaliza mais/h) | Distúrbios do Sono · Vitalidade e Disposição | 16 |
| O4 | **Eficiência do sono** | % | PSG, actigrafia, wearable | `≥85` ótimo · `75–85` · `65–75` · `<65` (convenção clínica, **confirmar**) | Distúrbios do Sono | 12 |
| O5 *(opcional)* | **ODI / SpO₂ noturno (dessaturação)** | eventos/h ou % | PSG, oximetria | hipoxemia noturna; sobrepõe IAH | Distúrbios do Sono · Cardiovascular · Pulmonar | 10 |

**Justificativa de peso:** SRI é o preditor de mortalidade **mais forte** (mais que duração) → peso alto. IAH é o diagnóstico objetivo de AOS (CV/mortalidade) → 18. Duração U-shape forte mas já há subjetivo → 16. Eficiência complementar → 12.

## §4 Sobreposição objetivo × subjetivo (decisão do Dr.)
Os objetivos espelham subjetivos existentes:
- O1 IAH ↔ `Apneias`(18) + `Roncos`(14)
- O2 SRI ↔ `Regularidade no dormir`(0) + `Regularidade no acordar`(12)
- O3 TST ↔ `Tempo de sono`(20)
- O4 Eficiência ↔ `Duração/Motivo interrupções`(12+12) + `Qualidade percebida`(14)

Opções:
- **(a)** Ambos pontuam (subjetivo = sintoma/hábito; objetivo = medido). Simples; aceita leve dupla-contagem no domínio sono. *(recomendado p/ começar, com sanity-check do peso total de sono)*
- **(b)** Objetivo **supera** o subjetivo quando presente (motor escolhe o objetivo) — mais limpo, exige lógica de exclusão mútua.
- **(c)** Objetivo entra cheio e reduz-se o peso do subjetivo correspondente.

## §5 Enxugar a minúcia subjetiva (opcional, decisão do Dr.)
~20 itens de ambiente de quarto (Pijamas/Travesseiros/Lençóis/Colchão/Odores/Janelas…) somam peso desproporcional de higiene. Opções: manter / zerar pontos (vira informativo) / agrupar num único item "Ambiente e higiene do sono". **Fora do escopo do "adicionar objetivo"; só sinalizado.**

## §6 Mecânica de implementação (por achado objetivo)
1. Definição **não-requisitável** numérica (ex.: code `PLNSLEEP_IAH`), `result_type='numeric'`, unidade.
2. `ScoreItem` keyed nessa definição + `ScoreLevels` (cutoffs acima) + pilares (M2M) + grupo `Exames›Laboratoriais` **ou** novo subgrupo `Sono › Medidas objetivas` (decidir).
3. PSG/HSAT/Actigrafia seguem orderáveis; ao laudar, preenche-se o(s) achado(s). Wearable: entrada manual ou import futuro.
4. Sem versão (EMR-completo). Hooks UUID v7 / LastReview.

## Decisões pendentes do Dr.
1. Quais objetivos entram: **O1–O4** (núcleo) e **O5** (sim/não)?
2. Sobreposição: **(a)/(b)/(c)**?
3. Onde vivem: subgrupo novo **`Sono › Medidas objetivas`** ou junto em `Laboratoriais`?
4. Enxugar a minúcia (§5) agora, depois, ou nunca?
5. Pesos finais de O1–O5.

## Fontes
- IAH/AASM: Sleep Foundation (AHI), Cleveland Clinic (AHI ranges).
- SRI mortalidade: UK Biobank, *Sleep* 2024 (zsad253) "Sleep regularity stronger predictor than duration"; eLife 88359.
- Duração U-shape: JAHA 2017 (JAHA.117.005947, PMID 28889101); Sci Rep srep21480.
- Wearables validade: meta-análise PMC11874098; JMIR mHealth e50983 (11 trackers); Sensors 24/6532.
