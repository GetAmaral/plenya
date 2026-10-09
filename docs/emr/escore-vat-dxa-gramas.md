# Escore — Gordura visceral (VAT) por DXA em gramas e cm³

> Estudo de 2026-10-09 para criar o item de escore. Status: **proposta, aguardando aprovação do
> Dr. Getúlio**. Nada aplicado em catálogo ainda (vai por migration goose quando aprovado).

## Por que um item novo

O escore só tem gordura visceral como **área (cm²)**, nos itens de anamnese
`GORDURA_VISCERAL_MULHER` / `GORDURA_VISCERAL_HOMEM` (mulher ≤40 ótimo … >150). Área vem de
tomografia (corte L4-L5) ou de DXA de outro fabricante. O DXA GE Lunar com **CoreScan** (Prodigy e
iDXA) entrega **massa (g) e volume (cm³)**, que não entram naquela escala. Hoje o número fica fora
do escore em silêncio.

## Massa ↔ volume

O CoreScan converte massa em volume por **densidade constante de 0,94 g/cm³**
(da Rosa et al., *Healthcare* 2021: "VAT mass (g) was automatically transposed into volume (cm3)
using a constant correction factor (0.94 g/mL)").

    cm³ = g ÷ 0,94        g = cm³ × 0,94

Conferido num laudo real: 1.005 g ÷ 0,94 = 1.069 cm³, laudo traz 1.065 cm³.
**Consequência de desenho:** um só item em gramas; resultado em cm³ entra por conversão de unidade.

## Evidência (todos com GE Lunar + CoreScan, salvo indicação)

| Estudo | População | Desfecho | Mulheres | Homens |
|---|---|---|---|---|
| Miazgowski, *PLoS One* 2017 (Prodigy) | 421 poloneses saudáveis, 20-30 a | referência | mediana 204 g · P80 405 · **P90 515 g** | mediana 391 g · P70 677 · **P90 1.259 g** |
| Meredith-Jones, *IJO* 2021 (Prodigy, CoreScan 16) | 1.482 neozelandeses, 18-66 a | risco cardiometabólico | **700 g** (<40 a) · **800 g** (≥40 a) | **1.000 g** (<40 a) · **1.200 g** (≥40 a) |
| SPAH, ACR 2024 (iDXA) | 707 brasileiros ≥65 a, baixa renda | SM (NCEP) | **600,5 g** (AUC 0,73) | **642,5 g** (AUC 0,74) |
| da Rosa, *Healthcare* 2021 (iDXA) | 270 militares brasileiros | SM | — | **1.086 g / 1.025 cm³** (AUC 0,80; S 77%, E 70%) |
| Bestaoui, *PLoS One* 2025 (iDXA) | 301 argelinos | SM (IDF) | **1.082 g** (AUC 0,855) | **1.369 g** (AUC 0,868) |
| Lundblad (Tromsø), *J Obes* 2021 (Prodigy, EnCore 17) | 3.675 noruegueses, 40-84 a | SM | **1.134 g** (Youden 0,30) | **1.859 g** (Youden 0,33) |
| Bi, *EJCN* 2015 (iDXA) | 229 mulheres obesas (EUA) | SM, prob. ≥50% | 1.320-1.713 cm³ (= 1.241-1.610 g) | — |

Notas: cortes de Youden têm discriminação modesta (0,3-0,6). Tromsø não achou ganho do VAT% (VAT ÷
gordura androide) sobre a massa para síndrome metabólica: **não propor item de VAT%**.
DXA de fabricantes diferentes **não** são comparáveis: a escala vale para GE Lunar/CoreScan.

## Proposta de faixas (convenção do escore: `between` meio-aberto `(lower, upper]`)

### Mulher (`gender = female`)

| Nível | Faixa (g) | Faixa (cm³) | Âncora |
|---|---|---|---|
| **5** ótimo | ≤ 500 | ≤ 532 | ≈ P90 de mulheres jovens saudáveis (515 g); abaixo de todos os cortes de risco, inclusive o do idoso brasileiro (600 g) |
| **4** bom | 500-800 | 532-851 | abaixo do corte de risco ≥40 a (800 g, Otago) |
| **2** médio | 800-1.100 | 851-1.170 | entre o corte de risco de Otago (800) e os de síndrome metabólica (1.082 Argélia, 1.134 Tromsø) |
| **0** ruim | > 1.100 | > 1.170 | no ou acima dos cortes de síndrome metabólica |

### Homem (`gender = male`)

| Nível | Faixa (g) | Faixa (cm³) | Âncora |
|---|---|---|---|
| **5** ótimo | ≤ 700 | ≤ 745 | ≈ P70 de homens jovens saudáveis (677 g); perto do corte do idoso brasileiro (642 g) |
| **4** bom | 700-1.100 | 745-1.170 | abaixo do corte do militar brasileiro (1.086 g) e de Otago ≥40 a (1.200 g) |
| **2** médio | 1.100-1.400 | 1.170-1.489 | entre militar/Otago (1.086-1.200) e Argélia (1.369) |
| **0** ruim | > 1.400 | > 1.489 | acima do corte de SM da Argélia; Tromsø (1.859) é mais tolerante |

Caso real que motivou: mulher, 52 a, 1.005 g / 1.065 cm³ → **nível 2** ("no limite"), coerente com
VAT% de 49% (acima do corte de Tromsø) e massa abaixo de 1.082-1.134.

## Implementação (quando aprovado)

1. Migration goose: `lab_test_definitions` filho de `PLND1D39793` (Densitometria corpo inteiro),
   numérico, unidade `g`; conversão cm³→g (×0,94) em `lab_test_unit_conversions`.
2. Dois `score_items` (female/male) com os níveis acima; `patient_gloss` curto ("gordura entre os órgãos").
3. Deploy do `api` (só com ordem), `reconvert-lab-units`, reclassificar, `recalc-scores`.

## Fontes
- Miazgowski 2017 — https://www.ncbi.nlm.nih.gov/pmc/articles/PMC5500349/
- Meredith-Jones 2021 — https://www.nature.com/articles/s41366-021-00743-3
- SPAH 2024 — https://acrabstracts.org/abstract/visceral-adipose-tissue-measured-by-dxa-predicts-metabolic-syndrome-in-low-income-community-dwelling-elderly-insights-from-the-sao-paulo-ageing-health-spah-study/
- da Rosa 2021 — https://www.ncbi.nlm.nih.gov/pmc/articles/PMC8304366/
- Bestaoui 2025 — https://pmc.ncbi.nlm.nih.gov/articles/PMC12419631/
- Lundblad 2021 — https://www.ncbi.nlm.nih.gov/pmc/articles/PMC8147540/
- Bi 2015 — https://www.nature.com/articles/ejcn2014227
