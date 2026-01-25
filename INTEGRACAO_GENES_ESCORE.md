# Guia de Integração - Painéis Genéticos no Escore Plenya

## Visão Geral

Sistema de estratificação de risco genético com **80 genes/variantes** integrado ao **Escore Plenya de Saúde Performance e Longevidade**.

**Arquivos Criados:**
1. `TABELA_GENES_80_COMPLETA.md` - Documentação completa com fontes científicas (80 genes) ✅
2. `genes_estratificacao_risco.csv` - Tabela CSV pronta para integração com 80 genes (formato idêntico aos exames laboratoriais) ✅
3. `PAINEIS_GENETICOS_BRASIL.md` - Guia de laboratórios e painéis disponíveis no Brasil
4. `INTEGRACAO_GENES_ESCORE.md` - Este guia de integração (atualizado para 80 genes) ✅

---

## Estrutura dos Dados

### Formato CSV

```csv
Nome;Unidade | Conversão;20;Nível 0;Nível 1;Nível 2;Nível 3;Nível 4;Nível 5
```

**Exemplo:**
```csv
FTO rs9939609 (Obesidade);Genótipo | A=risco T=proteção;20;AA + IMC>35;AA + IMC 30-35;AA + IMC 25-30;AT qualquer IMC;TT + IMC>25;TT + IMC<25
```

### Níveis de Risco (0-5)

- **Nível 0 (Crítico):** Mutações patogênicas + fenótipo manifesto
- **Nível 1 (Alto Risco):** Genótipo de risco homozigoto + fatores agravantes
- **Nível 2 (Risco Moderado-Alto):** Genótipo de risco + fenótipo leve
- **Nível 3 (Risco Leve):** Heterozigoto de risco ou genótipo protetor + fatores de risco
- **Nível 4 (Neutro/Basal):** Genótipo mais comum ou protetor sem otimização
- **Nível 5 (Ótimo):** Genótipo protetor + fenótipo ótimo

---

## Categorias de Genes (80 genes total)

### 1. Obesidade e Metabolismo (4 genes)
- **FTO rs9939609:** Alelo A (risco) OR 1.6 vs TT - "gene da obesidade"
- **MC4R rs17782313:** Alelo C (risco) OR 1.925 vs TT - receptor melanocortina
- **LEPR rs1137101:** Arg/Arg (risco) - receptor leptina, resistência leptínica
- **POMC rs6713532:** CC (risco) - produção hormônios saciedade

### 2. Diabetes e Resistência Insulínica (14 genes)
- **TCF7L2 rs7903146:** Alelo T (risco) OR 1.46 - principal gene diabetes tipo 2
- **PPARG Pro12Ala:** Alelo Ala (proteção) - sensibilidade insulínica
- **KCNJ11 rs5219:** Genótipo TT OR 3.09 - canal potássio células β
- **IRS1 rs1801278:** Arg/Arg OR 17.61 - substrato receptor insulina
- **SLC30A8 rs13266634:** Alelo C (risco) - transportador zinco
- **HHEX rs1111875:** CC (risco) - desenvolvimento pâncreas
- **IGF2BP2 rs4402960:** TT (risco) - secreção insulina
- **CDKAL1 rs7754840:** CC (risco) - função células β
- **GCK:** Mutações causam MODY2 (diabetes gestacional)
- **HNF1A:** Mutações causam MODY3 (80% necessitam insulina aos 25a)
- **HNF4A:** Mutações causam MODY1
- **HNF1B:** Mutações causam MODY5 (+ alterações renais)
- **INS VNTR:** Classe I/I (risco) - diabetes tipo 1
- **ABCC8 rs757110:** TT (risco) - canal ATP células β

### 3. Metabolismo Lipídico (14 genes)
**Colesterol e Lipoproteínas:**
- **APOE ε2/ε3/ε4:** ε4/ε4 OR 14.9 Alzheimer + ↑ LDL
- **LDLR rs688:** TT OR 13.2 para DCV - receptor LDL
- **PCSK9 R46L:** Proteção (↓34% LDL, ↓47% mortalidade CV)

**Triglicerídeos:**
- **APOA5 rs662799:** GG (risco) - TG 38.5% ↑
- **LPL rs328:** Portador X (S447X) - proteção (↓ TG)
- **APOC2:** Deficiência - hipertrigliceridemia grave
- **GPIHBP1:** Mutação - quilomicronemia (TG >1000)
- **LMF1:** Mutação - hipertrigliceridemia parcial

**HDL:**
- **APOA1 rs670:** AA (↓ HDL) - principal proteína HDL
- **LIPC rs1800588:** CC (lipase hepática baixa) - HDL alto + TG alto
- **ABCA1 rs9282541:** GG (↓ HDL) - efluxo colesterol
- **LCAT rs5923:** TT (↓ LCAT) - maturação HDL

**Absorção de Gordura:**
- **FADS1 rs174546:** TT ↓ conversão ômega-3
- **FADS2 rs174575:** Minor allele ↓ DHA
- **FABP2 Ala54Thr:** Thr/Thr ↑ absorção gordura (2x afinidade)

### 4. Homocisteína e Folato (1 gene)
- **MTHFR C677T:** TT 55-65% ↓ atividade - hiperhomocisteinemia

### 5. Hipertensão e Sistema Cardiovascular (7 genes)
- **ACE I/D rs4646994:** DD OR 4.032 - enzima conversora angiotensina
- **AGT rs699 M235T:** TT OR 1.72 - angiotensinogênio
- **ADD1 rs4961 Gly460Trp:** Trp/Trp - hipertensão sal-sensível
- **NOS3 rs1799983 Glu298Asp:** Asp/Asp ↓ produção NO
- **AGTR1 rs5186 A1166C:** CC (risco) - receptor angiotensina II
- **CYP11B2 rs1799998:** CC (↑ aldosterona) - síntese aldosterona
- **GNB3 rs5443 C825T:** TT (risco) - transdução sinal renal

### 6. Intolerâncias Alimentares (2 genes)
- **MCM6 rs4988235:** CC intolerância lactose (72.28% alelo C no Brasil NE)
- **HLA-DQ2/DQ8:** DQ2.5/DQ2.2 OR 1:10 para doença celíaca (98.4% celíacos BR)

### 7. Metabolismo de Álcool (2 genes)
- **ADH1B rs1229984 Arg48His:** His/His 70-80x ↑ atividade - álcool → acetaldeído
- **ALDH2 rs671 Glu504Lys:** Lys/Lys deficiência completa - acetaldeído → acetato

### 8. Metabolismo de Cafeína (1 gene)
- **CYP1A2 rs762551:** CC metabolizadores lentos (16%) - ↑ risco IAM com café

### 9. Vitaminas e Antioxidantes (3 genes)
- **VDR FokI rs2228570:** ff ↑ risco deficiência vitamina D
- **BCO1 rs6564851:** GG ↓ 49% conversão beta-caroteno → vitamina A
- **SLC23A1 rs33972313:** AA ↓ 40-50% vitamina C plasmática

### 10. Enzimas Antioxidantes e Detoxificação (10 genes)
**Antioxidantes Endógenos:**
- **SOD2 rs4880 Ala16Val:** CC associado a NASH
- **CAT rs1001179:** TT ↑ risco esteatose hepática
- **GPX1 rs1050450 Pro198Leu:** Leu/Leu ↓ atividade glutationa peroxidase

**Sistema Glutationa S-Transferase:**
- **GSTM1 nulo:** 48-65% populações, OR 1.46 câncer pulmão (tabagismo)
- **GSTT1 nulo:** 15-31% europeus, OR 2.50 câncer gástrico
- **GSTP1 rs1695 Ile105Val:** Val/Val ↓ detoxificação xenobióticos

**Citocromos P450:**
- **EPHX1 rs1051740 Tyr113His:** His/His ↓ hidrolase epóxido
- **CYP1A1 rs4646903 MspI:** CC ↑ ativação PAH (hidrocarbonetos aromáticos)
- **CYP2A6 rs1801272:** Deficiência - metabolismo nicotina
- **NAT2:** Acetilador lento - toxicidade isoniazida + aminas aromáticas

### 11. Marcadores Inflamatórios (4 genes)
- **TNF rs1800629 -308G>A:** AA OR 1.4 para SOP - ↑ TNF-α
- **IL6 rs1800795 -174G>C:** CC modulação IL-6 (variável por etnia)
- **IL1B rs16944 -511C>T:** TT ↑ IL-1β - inflamação sistêmica
- **CRP rs1130864:** AA ↑ proteína C reativa basal

### 12. Demência e Alzheimer (4 genes)
- **APOE ε2/ε3/ε4:** ε4/ε4 OR 14.9 para Alzheimer (já listado em Lipídios)
- **APP A673T rs63750847:** Proteção rara (<0.1%) - mutação islandesa
- **PSEN1 E280A:** AD familial ~100% penetrância (sintomas 35-50a)
- **PSEN2:** <5% AD precoce familial

### 13. ⭐ Sport e Performance Física (4 genes - NOVO)
- **ACTN3 rs1815739 R577X:** RR (poder/velocidade), XX (resistência) - "gene do velocista"
- **PPARA rs4253778:** GG (endurance) - oxidação lipídica muscular
- **PPARGC1A rs8192678 Gly482Ser:** Gly/Gly (atletas) - biogênese mitocondrial
- **COL5A1 rs12722:** TT (lesões tendão) - colágeno tipo V

### 14. ⭐ Doença de Parkinson (5 genes - NOVO)
- **SNCA rs356219:** GG (risco) - alfa-sinucleína
- **LRRK2 G2019S rs34637584:** Mutação (49% penetrância aos 80a) - 41% marroquinos
- **PARK2 (Parkin):** Mutações bialélicas - Parkinson precoce AR (18% DP <40a)
- **PINK1:** Mutações bialélicas - Parkinson precoce AR
- **PARK7 (DJ-1):** Mutações bialélicas - Parkinson precoce AR (raro)

### 15. ⭐ Demência Frontotemporal (3 genes - NOVO)
- **MAPT rs1467967 H1/H2:** H1/H1 (risco DFT, PSP) - haplótipos tau
- **GRN rs5848:** TT ↓ progranulina - 5-10% DFT familial
- **C9orf72:** Expansão GGGGCC (61.5% DFT/ELA familial) - 30 repeats patogênico

### 16. ⭐ Saúde Óssea e Colágeno (3 genes - NOVO)
- **COL1A1 rs1800012 Sp1:** TT OR 2.45 osteoporose - colágeno tipo I
- **ESR1 rs2234693 PvuII:** CC (risco osteoporose) - receptor estrogênio α
- **ALPL:** Mutações bialélicas - hipofosfatasia (fosfatase alcalina baixa)

---

## Fluxo de Integração no Sistema

### 1. Coleta de Dados Genéticos

**Entrada do Paciente:**
- Upload de resultado do painel genético (PDF do laboratório)
- Seleção manual de genótipos via interface
- Importação de arquivo VCF (futuro)

**Laboratórios Brasileiros Compatíveis:**
- DB Molecular (Diagnósticos do Brasil) - 95 genes nutrigenética
- Fleury Genômica - APOE, Alzheimer 4 genes
- Genera - >100 variantes DTC
- GnTech - APOE isolado
- Sabin Genômica
- Dasa Genômica
- Anaclin Gene - 663 genes endócrino-metabólico
- Mendelics - MODY 56 genes

### 2. Cálculo de Risco Genético

**Para cada gene:**
1. Identificar genótipo do paciente (ex: FTO rs9939609 = AA)
2. Consultar dados fenotípicos relevantes (ex: IMC = 32)
3. Mapear para o nível de risco apropriado (ex: AA + IMC 30-35 → Nível 1)
4. Atribuir pontuação: Nível 0 = 0 pontos, Nível 5 = 100 pontos

**Fórmula de Pontuação:**
```
Pontuação = (Nível / 5) × 100
Nível 0 → 0 pontos
Nível 1 → 20 pontos
Nível 2 → 40 pontos
Nível 3 → 60 pontos
Nível 4 → 80 pontos
Nível 5 → 100 pontos
```

### 3. Integração com Exames Laboratoriais

**Exemplo - Diabetes:**

**Exames Lab:**
- Glicemia jejum: 110 mg/dL → Nível 2 (40 pontos)
- HbA1c: 6.0% → Nível 3 (60 pontos)
- Insulina jejum: 15 µU/mL → Nível 2 (40 pontos)

**Genética:**
- TCF7L2 rs7903146: TT → Nível 1 (20 pontos)
- KCNJ11 rs5219: CT → Nível 3 (60 pontos)
- SLC30A8 rs13266634: CC → Nível 1 (20 pontos)

**Escore Diabetes Final:**
```
Média Lab = (40 + 60 + 40) / 3 = 46.7
Média Genética = (20 + 60 + 20) / 3 = 33.3
Escore Final = (Média Lab × 0.7) + (Média Genética × 0.3)
             = (46.7 × 0.7) + (33.3 × 0.3)
             = 32.7 + 10.0
             = 42.7 pontos (Risco Moderado-Alto)
```

**Peso sugerido:** 70% laboratorial, 30% genético (genética é fator de risco, não diagnóstico)

### 4. Contexto Clínico Obrigatório

**Genes que EXIGEM fenótipo para estratificação:**
- **FTO, MC4R:** Necessitam IMC atual
- **TCF7L2, KCNJ11, IRS1:** Necessitam glicemia/HbA1c
- **APOE, LDLR, PCSK9:** Necessitam perfil lipídico
- **ACE, AGT, NOS3:** Necessitam pressão arterial
- **MTHFR:** Necessita homocisteína
- **ADH1B, ALDH2:** Necessitam histórico de consumo
- **CYP1A2:** Necessita consumo de cafeína

**Genes independentes de fenótipo (estratificação direta):**
- **HLA-DQ2/DQ8:** Risco celíaca (mas sintomas aumentam nível)
- **GSTM1/GSTT1:** Capacidade detoxificação (mas exposição importa)
- **APP A673T:** Proteção Alzheimer (raro)
- **PSEN1/PSEN2:** Mutações patogênicas (AD familial)

---

## Recomendações Clínicas por Nível

### Nível 0-1 (Crítico/Alto Risco)
- **Ação:** Encaminhamento especialista obrigatório
- **Exames:** Repetir a cada 3-6 meses
- **Intervenção:** Farmacológica + nutricional + lifestyle agressivo
- **Exemplo:** APOE ε4/ε4 + LDL 200 mg/dL → Cardiologista + estatina

### Nível 2 (Risco Moderado-Alto)
- **Ação:** Acompanhamento médico próximo
- **Exames:** Repetir a cada 6 meses
- **Intervenção:** Nutricional intensiva + suplementação direcionada
- **Exemplo:** MTHFR TT + Hcy 12 µmol/L → L-metilfolato 1000 mcg

### Nível 3 (Risco Leve)
- **Ação:** Monitoramento anual
- **Exames:** Repetir anualmente
- **Intervenção:** Orientação nutricional + lifestyle
- **Exemplo:** FTO AT + IMC 26 → Dieta hipocalórica + exercício

### Nível 4-5 (Neutro/Ótimo)
- **Ação:** Prevenção e manutenção
- **Exames:** Conforme protocolo de rotina
- **Intervenção:** Manutenção de hábitos saudáveis
- **Exemplo:** PCSK9 R46L → Acompanhamento padrão

---

## Interações Gene-Nutriente e Estilo de Vida (Nutrigenômica)

### Alta Prioridade para Intervenção Nutricional e Lifestyle

| Gene | Genótipo Risco | Intervenção | Fundamento | Evidência |
|------|----------------|-------------|------------|-----------|
| **VITAMINAS E MINERAIS** |||||
| MTHFR | TT | L-metilfolato 400-1000 mcg + B12 + B6 | ↓ 55-65% atividade enzimática | ↓ Hcy 20-30% |
| VDR | ff | Vitamina D3 4000-5000 UI + K2 MK-7 | ↓ resposta ao colecalciferol | Níveis 25-OH-D |
| BCO1 | GG | Vitamina A pré-formada 3000-5000 UI | ↓ 49% conversão beta-caroteno | Retinol sérico |
| SLC23A1 | AA | Vitamina C 500-1000 mg 2x/dia | ↓ 40-50% absorção/reabsorção | Vit C plasmática |
| **ÔMEGA-3 E LIPÍDIOS** |||||
| FADS1/2 | TT/minor | EPA/DHA 2000-3000 mg/dia (não ALA!) | ↓ conversão ALA→EPA/DHA | Índice ômega-3 |
| APOE | ε4/ε4 | ↓ gordura saturada <7%, ↑ ômega-3 | ↑↑ resposta a gordura saturada | ↓ LDL 15-20% |
| FABP2 | Thr/Thr | ↓ gordura saturada <10%, ↑ insaturada | 2x ↑ afinidade ácidos graxos | ↓ TG, ↓ RI |
| APOA5 | GG | ↓ gordura total <25%, ↓ álcool | TG 38.5% ↑ vs AA | ↓ TG 30-40% |
| **OBESIDADE** |||||
| FTO | AA | ↓ densidade calórica, ↑ proteína 25-30% | ↑ apetite, ↑ preferência gordura | ↓ peso 2-3 kg |
| MC4R | CC | Controle porções, fibras solúveis 30g/d | ↑ ingestão calórica espontânea | ↓ IMC 1-2 |
| LEPR | Arg/Arg | Dieta low-carb, jejum intermitente | Resistência leptínica | ↓ leptina 15% |
| POMC | CC | Refeições frequentes, ↑ saciedade | ↓ hormônios saciedade | Controle fome |
| **HIPERTENSÃO** |||||
| ADD1 | Trp/Trp | ↓ sódio <1500 mg/dia, ↑ potássio | Hipertensão sal-sensível | ↓ PA 10-15 mmHg |
| ACE | DD | ↓ sódio <2000 mg/dia + IECA | ↑ angiotensina II | ↓ PA 5-10 mmHg |
| CYP11B2 | CC | ↓ sódio, ↑ magnésio 400 mg | ↑ aldosterona | ↓ retenção Na+ |
| **DETOXIFICAÇÃO** |||||
| GSTM1/T1 | Nulo/Nulo | N-acetilcisteína 600 mg 2x/d + vegetais crucíferos | ↓ glutationa conjugação | ↓ oxidação 20% |
| GSTP1 | Val/Val | Sulforafano 30 mg/d (brócolis germinado) | ↓ detox xenobióticos | Indução GST |
| SOD2 | CC | Manganês 5 mg + CoQ10 200 mg | ↓ SOD mitocondrial | ↓ estresse oxidativo |
| GPX1 | Leu/Leu | Selênio 200 mcg (castanha-do-pará 2-3/d) | ↓ glutationa peroxidase | ↑ atividade GPX |
| NAT2 | Acetilador lento | Evitar carnes processadas + aminas | Acúmulo metabólitos tóxicos | ↓ risco CA |
| **CAFÉ E ÁLCOOL** |||||
| CYP1A2 | CC | ↓ café <200 mg cafeína/dia (1-2 xíc) | Metabolizador lento | ↓ risco IAM |
| ADH1B | Arg/Arg | Abstinência ou consumo mínimo | Metabolizador lento álcool | ↓ alcoolismo |
| ALDH2 | Lys/+ | Abstinência COMPLETA álcool | Acúmulo acetaldeído | ↓ CA 7-10x |
| **SPORT E PERFORMANCE** |||||
| ACTN3 | XX | Treino endurance (corrida, ciclismo longa dist) | Deficiência α-actinin-3 tipo II | Performance resist |
| ACTN3 | RR | Treino poder/velocidade (HIIT, sprint, peso) | Fibras tipo II completas | Performance poder |
| PPARA | GG | Treino aeróbico >60min, dieta ↑ gordura | ↑ oxidação lipídica muscular | VO2max +8% |
| PPARGC1A | Gly/Gly | Treino intervalado, CoQ10 200 mg | ↑ biogênese mitocondrial | Mitocôndrias +15% |
| COL5A1 | TT | Vitamina C 1-2g + colágeno tipo I 10g + aquecimento | ↑ risco lesão tendão | ↓ lesões 30% |
| **OSSO** |||||
| COL1A1 | TT | Vitamina D 4000 UI + K2 + cálcio 1200 mg + peso | ↓ colágeno tipo I ósseo | ↑ DMO 3-5% |
| ESR1 | CC | Isoflavonas 50 mg + exercício resistido | ↓ resposta estrogênio | ↑ DMO 2-3% |
| **INFLAMAÇÃO** |||||
| TNF | AA | Dieta anti-inflamatória (mediterrânea) | ↑ TNF-α basal | ↓ PCR 30% |
| IL6 | CC | Ômega-3 EPA 2g + cúrcuma 500 mg | ↑ IL-6 (variável) | ↓ IL-6 20% |
| IL1B | TT | Ômega-3 + antioxidantes + evitar açúcar | ↑ IL-1β inflamatória | ↓ inflamação 25% |
| CRP | AA | Dieta low-glycemic + perda peso | ↑ CRP basal | ↓ PCR 40-50% |

---

## Alertas Clínicos Automáticos

### Combinações de Alto Risco

**Implementar alertas quando:**

**Cardiovascular e Alzheimer:**
1. **APOE ε4/ε4 + LDL >160 mg/dL**
   - Alerta: "⚠️ CRÍTICO: Risco cardiovascular e Alzheimer extremamente elevado - considerar estatina + EPA/DHA 2g/dia"

2. **APOE ε4/ε4 + APOA5 GG + TG >200**
   - Alerta: "⚠️ Risco duplo CV e demência - estatina + fenofibrato + dieta baixa gordura saturada"

**Metabolismo e Obesidade:**
3. **FTO AA + MC4R CC + LEPR Arg/Arg + IMC >30**
   - Alerta: "⚠️ Risco genético triplo obesidade - intervenção intensiva: dieta hipocalórica + exercício 5x/sem + considerar farmacoterapia"

4. **PPARGC1A Ser/Ser + sedentarismo + síndrome metabólica**
   - Alerta: "⚠️ Genética desfavorável + estilo de vida - exercício aeróbico obrigatório (mitocôndrias)"

**Diabetes:**
5. **TCF7L2 TT + KCNJ11 TT + HHEX CC + HbA1c >6.0%**
   - Alerta: "⚠️ Alto risco progressão diabetes tipo 2 - considerar metformina preventiva + TOTG"

6. **Mutação MODY (GCK/HNF1A/HNF4A/HNF1B) + diabetes gestacional**
   - Alerta: "⚠️ MODY confirmado - aconselhamento genético + teste familiares + ajuste terapêutico"

**Cardiovascular:**
7. **ACE DD + AGT TT + AGTR1 CC + HAS estágio 2**
   - Alerta: "⚠️ Risco genético triplo hipertensão - IECA obrigatório + monitoramento rigoroso"

8. **APOA1 AA + ABCA1 GG + HDL <35 mg/dL**
   - Alerta: "⚠️ HDL criticamente baixo com genética desfavorável - considerar niacina ou fibratos"

**Homocisteína:**
9. **MTHFR TT + Homocisteína >12 µmol/L**
   - Alerta: "⚠️ Suplementar L-metilfolato 1000 mcg + B12 metilcobalamina 1000 mcg + B6 50 mg"

**Detoxificação:**
10. **GSTM1 nulo + GSTT1 nulo + GSTP1 Val/Val + Tabagismo**
    - Alerta: "⚠️ CRÍTICO: Risco câncer extremo - cessação tabágica URGENTE + antioxidantes + N-acetilcisteína"

11. **NAT2 acetilador lento + uso isoniazida**
    - Alerta: "⚠️ Ajustar dose isoniazida (hepatotoxicidade) + monitorar TGO/TGP mensalmente"

**Álcool:**
12. **ALDH2 Lys/+ + Consumo álcool (qualquer quantidade)**
    - Alerta: "🚨 CRÍTICO: Abstinência OBRIGATÓRIA - risco 7-10x câncer esôfago/cabeça/pescoço"

13. **ADH1B Arg/Arg + ALDH2 Glu/Lys + consumo álcool moderado**
    - Alerta: "⚠️ Metabolismo lento + acúmulo acetaldeído - abstinência ou consumo mínimo"

**Intolerâncias:**
14. **HLA-DQ2.5/DQ2.2 + Sintomas gastrointestinais + anemia ferropriva**
    - Alerta: "⚠️ Risco 1:10 doença celíaca - dosagem anti-transglutaminase + biópsia duodenal"

15. **MCM6 CC + sintomas graves lactose + consumo diário**
    - Alerta: "⚠️ Intolerância lactose confirmada geneticamente - dieta sem lactose + cálcio alternativo"

**Demência:**
16. **PSEN1 E280A ou PSEN2 mutação patogênica + idade >30a**
    - Alerta: "🚨 CRÍTICO: Alzheimer familial confirmado - aconselhamento genético + ressonância anual + ensaios clínicos"

17. **APOE ε4/ε4 + MAPT H1/H1 + declínio cognitivo**
    - Alerta: "⚠️ Alto risco demência mista (Alzheimer + FTD) - avaliação neuropsicológica + neurologia"

18. **C9orf72 expansão + idade >50a assintomático**
    - Alerta: "⚠️ Alto risco DFT/ELA (61.5% penetrância familial) - monitoramento neurológico anual"

**Parkinson:**
19. **LRRK2 G2019S + idade >60a assintomático**
    - Alerta: "⚠️ Risco Parkinson 28% aos 70a, 49% aos 80a - avaliação neurológica anual + DAT-scan preventivo"

20. **PARK2/PINK1/DJ-1 bialélico + sintomas parkinsonianos <40a**
    - Alerta: "⚠️ Parkinson precoce AR - neurologia + considerar teste genético familiares"

**Sport e Lesões:**
21. **COL5A1 TT + lesões recorrentes tendão/ligamento**
    - Alerta: "⚠️ Risco genético lesão tecido conjuntivo - reforço muscular + vitamina C 1-2g + colágeno tipo I"

22. **ACTN3 XX + treinamento velocidade/força sem resposta**
    - Alerta: "ℹ️ Genética favorável resistência, desfavorável poder - ajustar treino para endurance"

**Osso e Inflamação:**
23. **COL1A1 TT + ESR1 CC + osteoporose + fraturas**
    - Alerta: "⚠️ Alto risco genético osteoporose - vitamina D 4000 UI + K2 + cálcio + considerar bifosfonatos"

24. **TNF AA + IL6 CC + IL1B TT + CRP AA + doenças autoimunes**
    - Alerta: "⚠️ Perfil inflamatório genético extremo - dieta anti-inflamatória + ômega-3 EPA 2g + cúrcuma"

**Vitaminas:**
25. **FADS1 TT + FADS2 minor/minor + ômega-3 baixo + dieta vegetariana**
    - Alerta: "⚠️ Conversão ALA→EPA/DHA prejudicada - suplementar EPA/DHA 2000-3000 mg/dia (não ALA)"

26. **BCO1 GG + retinol <30 µg/dL + dieta vegetariana estrita**
    - Alerta: "⚠️ Conversão beta-caroteno prejudicada - vitamina A pré-formada 3000-5000 UI"

---

## Relatório para o Paciente

### Modelo de Relatório Genético

**ESCORE PLENYA - PERFIL GENÉTICO**

**Paciente:** [Nome]
**Data:** [Data]
**Painel:** [Lab] - [Nome do Painel]

---

**RESUMO EXECUTIVO**

Escore Genético Global: **XX/100 pontos**
- Obesidade/Metabolismo: XX/100
- Diabetes: XX/100
- Cardiovascular: XX/100
- Detoxificação: XX/100
- Vitaminas: XX/100

---

**GENÓTIPOS DE ALTO RISCO** (Níveis 0-2)

| Gene | Genótipo | Significado | Recomendação |
|------|----------|-------------|--------------|
| FTO rs9939609 | AA | Alto risco obesidade | Dieta hipocalórica + exercício 5x/sem |
| APOE | ε4/ε4 | Alto risco Alzheimer | EPA/DHA 2g/dia + estatina |

---

**GENÓTIPOS PROTETORES** (Nível 5)

| Gene | Genótipo | Benefício |
|------|----------|-----------|
| PCSK9 | R46L | Proteção cardiovascular natural |

---

**RECOMENDAÇÕES NUTRICIONAIS PERSONALIZADAS**

1. **MTHFR TT:** Suplementar L-metilfolato 800 mcg/dia
2. **FADS1 TT:** Consumir EPA/DHA 2000 mg/dia (peixe gordo 3x/sem ou suplemento)
3. **CYP1A2 CC:** Limitar café a 1 xícara/dia

---

**MONITORAMENTO RECOMENDADO**

- Glicemia jejum + HbA1c: a cada 6 meses (risco diabetes)
- Perfil lipídico completo: a cada 6 meses (risco cardiovascular)
- Homocisteína: a cada 6 meses (MTHFR TT)

---

## Considerações Éticas e Legais (LGPD)

### Armazenamento de Dados Genéticos

**Classificação:** Dados sensíveis (Art. 5º, II, LGPD)

**Requisitos:**
1. Consentimento explícito e específico
2. Criptografia obrigatória (pgcrypto)
3. Audit log de TODOS os acessos
4. Retenção mínima: 20 anos (dados médicos)
5. Direito ao esquecimento (exceto casos legais)

**Tabela Database:**
```sql
CREATE TABLE genetic_tests (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    patient_id UUID NOT NULL REFERENCES patients(id),
    test_date DATE NOT NULL,
    laboratory VARCHAR(100),
    panel_name VARCHAR(200),
    raw_data BYTEA, -- Encrypted VCF/PDF
    genotypes JSONB, -- Encrypted genotypes
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Genotypes JSONB structure (encrypted):
{
  "FTO_rs9939609": "AA",
  "APOE": "e3/e4",
  "MTHFR_C677T": "CT",
  ...
}
```

### Consentimento Informado

**Texto obrigatório:**
> "Autorizo o armazenamento e análise dos meus dados genéticos para fins de estratificação de risco de saúde e recomendações nutricionais personalizadas. Estou ciente de que:
> 1. Dados genéticos são informações sensíveis protegidas pela LGPD
> 2. Meus dados serão criptografados e acessados apenas por profissionais autorizados
> 3. Posso solicitar exclusão dos dados a qualquer momento (direito ao esquecimento)
> 4. Resultados genéticos não são diagnóstico, apenas fatores de risco
> 5. Aconselhamento genético profissional pode ser necessário para mutações patogênicas"

---

## Próximos Passos - Implementação

### Fase 1: Backend (Go)

**1.1 Database Schema:**
```sql
-- Tabela de testes genéticos
CREATE TABLE genetic_tests (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    patient_id UUID NOT NULL REFERENCES patients(id),
    test_date DATE NOT NULL,
    laboratory VARCHAR(100), -- DB Molecular, Fleury, Genera, etc
    panel_name VARCHAR(200), -- Nome do painel
    panel_type VARCHAR(50), -- nutrigenetica, metabolismo, sport, demencia
    raw_data BYTEA, -- PDF/VCF encrypted
    genotypes JSONB, -- Genótipos encrypted
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Catálogo de 80 variantes genéticas
CREATE TABLE genetic_variants (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    gene_name VARCHAR(50) NOT NULL, -- FTO, APOE, MTHFR, etc
    variant_name VARCHAR(100), -- rs9939609, C677T, ε2/ε3/ε4, etc
    category VARCHAR(50), -- obesidade, diabetes, cardiovascular, sport, parkinson, etc
    chromosome VARCHAR(10),
    position BIGINT,
    risk_allele VARCHAR(50),
    protective_allele VARCHAR(50),
    clinical_significance TEXT,
    population_frequency JSONB, -- Frequências por população
    references JSONB -- Links científicos
);

-- Níveis de risco (0-5) para cada gene
CREATE TABLE genetic_risk_levels (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    variant_id UUID REFERENCES genetic_variants(id),
    level_number INT CHECK (level_number BETWEEN 0 AND 5),
    genotype VARCHAR(50), -- AA, AT, TT, ε4/ε4, etc
    phenotype_condition TEXT, -- IMC>35, LDL>190, etc
    description TEXT, -- Descrição do nível
    recommendations TEXT -- Recomendações específicas
);

-- Resultados genéticos dos pacientes
CREATE TABLE patient_genetic_results (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    patient_id UUID REFERENCES patients(id),
    genetic_test_id UUID REFERENCES genetic_tests(id),
    variant_id UUID REFERENCES genetic_variants(id),
    genotype VARCHAR(50), -- Genótipo do paciente
    risk_level INT, -- Nível calculado (0-5)
    score INT, -- Pontuação (0-100)
    phenotype_data JSONB, -- Dados fenotípicos usados no cálculo
    calculated_at TIMESTAMP DEFAULT NOW()
);
```

**1.2 API Endpoints:**
- `POST /api/v1/genetic-tests` - Upload teste genético (PDF/manual)
- `GET /api/v1/patients/:id/genetic-profile` - Perfil completo (80 genes)
- `GET /api/v1/patients/:id/genetic-score` - Escore genético por categoria
- `GET /api/v1/patients/:id/genetic-alerts` - Alertas clínicos automáticos
- `GET /api/v1/patients/:id/genetic-recommendations` - Recomendações personalizadas
- `GET /api/v1/genetic-variants` - Catálogo de 80 variantes
- `POST /api/v1/genetic-results/calculate` - Recalcular risco com novos exames

**1.3 Lógica de Cálculo:**
```go
// Algoritmo de estratificação de risco
func CalculateGeneticRisk(patient Patient, genotype string, phenotypeData map[string]interface{}) (level int, score int) {
    // 1. Buscar regra de estratificação para o gene
    rules := getRiskRules(genotype)

    // 2. Avaliar fenótipo (exames laboratoriais, IMC, PA, etc)
    phenotypeMatch := matchPhenotype(phenotypeData, rules)

    // 3. Determinar nível de risco (0-5)
    level = determineRiskLevel(genotype, phenotypeMatch)

    // 4. Calcular pontuação (0-100)
    score = (level * 100) / 5

    // 5. Gerar alertas se necessário
    checkAlerts(patient, genotype, level, phenotypeData)

    return level, score
}
```

### Fase 2: Frontend Web

**2.1 Páginas Principais:**
1. **Upload de Painel Genético:**
   - Upload PDF (parsing automático ou manual)
   - Seleção de laboratório (DB Molecular, Fleury, Genera, etc)
   - Entrada manual de genótipos (80 genes)
   - Validação de genótipos

2. **Dashboard Genético:**
   - **Cards por Categoria (16 categorias):**
     - Obesidade/Metabolismo (4 genes)
     - Diabetes (14 genes)
     - Lipídios (14 genes)
     - Cardiovascular (7 genes)
     - Sport/Performance (4 genes)
     - Parkinson (5 genes)
     - Demência FTD (3 genes)
     - Alzheimer (4 genes)
     - Osso (3 genes)
     - Detoxificação (10 genes)
     - Inflamação (4 genes)
     - Vitaminas (3 genes)
     - Homocisteína (1 gene)
     - Intolerâncias (2 genes)
     - Álcool (2 genes)
     - Cafeína (1 gene)

   - **Score por Categoria:** Radial chart com 16 eixos
   - **Top 10 Genes de Risco:** Lista com níveis 0-2
   - **Top 10 Genes Protetores:** Lista com nível 5
   - **Alertas Críticos:** Banner vermelho com alertas de combinações

3. **Perfil Genético Completo:**
   - Tabela com 80 genes
   - Filtros por categoria
   - Genótipo + Nível de risco + Pontuação
   - Link para referências científicas

4. **Recomendações Personalizadas:**
   - **Nutrição:** Tabela gene-nutriente com intervenções
   - **Suplementação:** Lista priorizada por risco
   - **Lifestyle:** Exercício, sono, exposições
   - **Monitoramento:** Exames recomendados + periodicidade

5. **Relatório PDF:**
   - Resumo executivo (1 página)
   - Genótipos de alto risco (tabela)
   - Genótipos protetores (tabela)
   - Recomendações nutricionais
   - Plano de monitoramento
   - Referências científicas

**2.2 Componentes UI:**
```typescript
// Componente de Score Genético
<GeneticScoreCard
  category="Diabetes"
  genes={14}
  averageScore={42.7}
  riskLevel="Moderado-Alto"
  topRiskGenes={[
    { gene: "TCF7L2", genotype: "TT", level: 1 },
    { gene: "KCNJ11", genotype: "TT", level: 1 }
  ]}
/>

// Componente de Alerta
<GeneticAlert
  severity="critical"
  title="ALDH2 Lys/+ + Consumo Álcool"
  message="Abstinência obrigatória - risco 7-10x câncer esôfago"
  action="Marcar consulta com gastroenterologista"
/>

// Radar Chart 16 Categorias
<GeneticRadarChart
  categories={[
    { name: "Obesidade", score: 65 },
    { name: "Diabetes", score: 42 },
    { name: "Cardiovascular", score: 55 },
    { name: "Sport", score: 80 },
    // ... 12 mais
  ]}
/>
```

### Fase 3: Algoritmo de Pontuação Integrado

**3.1 Integração com Exames Laboratoriais:**
```typescript
// Exemplo: Escore Diabetes
const diabetesScore = calculateDiabetesScore({
  // Exames laboratoriais (peso 70%)
  labs: {
    glicemiaJejum: 110, // Nível 2 (40 pontos)
    hba1c: 6.0,         // Nível 3 (60 pontos)
    insulinaJejum: 15   // Nível 2 (40 pontos)
  },
  // Genética (peso 30%)
  genetics: {
    TCF7L2: "TT",      // Nível 1 (20 pontos)
    KCNJ11: "CT",      // Nível 3 (60 pontos)
    PPARG: "Pro/Pro",  // Nível 2 (40 pontos)
    HHEX: "CC"         // Nível 1 (20 pontos)
  }
});

// Resultado: 44.3 pontos (Risco Moderado-Alto)
```

**3.2 Sistema de Alertas:**
- Monitorar 26 combinações de alto risco
- Enviar notificações ao médico
- Sugerir condutas automáticas
- Agendar consultas prioritárias

### Fase 4: Relatórios e Insights

**4.1 Relatório para Paciente:**
- Linguagem acessível (evitar jargão)
- Foco em ações práticas
- Gráficos coloridos e intuitivos
- QR code com link para vídeos educativos

**4.2 Relatório para Médico:**
- Referências científicas completas (>200 estudos)
- Odds ratios e frequências populacionais
- Recomendações baseadas em evidência
- Links para guidelines (SBC, SBD, ADA, ESC)

**4.3 Timeline de Evolução:**
- Comparar escores genéticos ao longo do tempo
- Mostrar impacto de intervenções no fenótipo
- Alertar sobre deterioração de risco
- Celebrar melhorias (gamificação)

---

## Referências Científicas

Todas as **80 variantes genéticas** documentadas possuem referências científicas verificáveis em:
- `TABELA_GENES_80_COMPLETA.md` - Documentação completa com >200 estudos científicos
- Fontes: Nature, NEJM, JAMA, BMC, PMC/PubMed, Frontiers, MDPI, Springer, BMJ, Lancet

**Categorias documentadas:**
- 16 categorias principais
- 80 genes/variantes com SNPs específicos
- Frequências alélicas por população (incluindo dados brasileiros)
- Odds ratios para doenças
- Intervenções nutricionais baseadas em evidência
- Protocolos de monitoramento clínico

**Laboratórios Brasileiros Compatíveis:**
- DB Molecular - 95 genes nutrigenética
- Fleury Genômica - APOE + painéis neurogenética
- Genera - >100 variantes DTC (direct-to-consumer)
- GnTech - APOE isolado
- Sabin Genômica - Painéis customizados
- Dasa Genômica - Oncogenética + nutrigenética
- Anaclin Gene - 663 genes endócrino-metabólico
- Mendelics - MODY 56 genes + neurogenética

**Próxima Expansão (Futuro):**
- Farmacogenética: CYP2D6, CYP2C19, CYP3A4/5, VKORC1, SLCO1B1
- Oncogenética: BRCA1/2, TP53, PALB2, CHEK2, ATM
- Cardio avançado: SCN5A, KCNH2, MYBPC3, MYH7
- Monogênicas: FH (LDLR/APOB/PCSK9), Hemocromatose (HFE)

**Última atualização:** Janeiro 2026 - Versão 2.0 (80 genes)
