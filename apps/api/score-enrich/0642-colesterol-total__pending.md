# ScoreItem: Colesterol Total

**ID:** `c77cedd3-2800-75e0-891b-a0c51a47ccc1`
**FullName:** Colesterol Total (Exames - Laboratoriais)
**Unit:** mg/dL

**Preparation Metadata:**
- Quality Grade: **EXCELLENT**
- Total Chunks: 30 de 11 artigos
- Avg Similarity: 0.632

---

## Contexto

Você é um especialista em medicina funcional integrativa e está contribuindo com o **Escore Plenya** — um escore completo de análise de saúde que avalia todos os aspectos da saúde, performance e longevidade humana. Cada ScoreItem representa um parâmetro clínico, laboratorial, genético, comportamental ou histórico que compõe esse escore.

Seu papel é gerar conteúdo clínico de alta qualidade para enriquecer cada parâmetro do escore com relevância clínica, orientação ao paciente e conduta prática.

**Regras inegociáveis:**
- Use **apenas** o conhecimento médico real consolidado e os dados presentes nos chunks científicos abaixo
- **Não alucine, não invente** dados, estudos, estatísticas ou referências que não estejam nos chunks ou no seu conhecimento médico estabelecido
- Se um dado específico não constar nos chunks e não for do seu conhecimento consolidado, **não o inclua**
- Seja preciso: prefira omitir a inventar

## Instrução

Com base nos chunks científicos abaixo, gere as respostas em formato JSON.

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `c77cedd3-2800-75e0-891b-a0c51a47ccc1`.**

```json
{
  "score_item_id": "c77cedd3-2800-75e0-891b-a0c51a47ccc1",
  "clinical_relevance": "Texto técnico para médicos (1000-5000 chars): definição fisiológica precisa, valores de referência e interpretação, fisiopatologia resumida, dados epidemiológicos com números concretos, estratificação de risco baseada em evidências.",
  "points": 1,
  "patient_explanation": "Texto simples para pacientes (500-1000 chars): o que é este parâmetro sem jargões, por que é importante para a saúde, o que valores alterados podem significar. Tom tranquilizador e educativo.",
  "conduct": "Conduta clínica em Markdown (1000-5000 chars): investigação complementar necessária, critérios de encaminhamento a especialistas, intervenções baseadas em evidências. Use bullet points, seções e negrito."
}
```

**Regras para `points` (1-50):**
- Baixo impacto clínico: 1-9 pts
- Alto impacto clínico: 10-19 pts
- Alto impacto em mortalidade: 20-50 pts
- Critérios: gravidade/mortalidade (40%), prevalência (30%), intervencionabilidade (30%)

---

### Contexto Científico

**ScoreItem:** Colesterol Total (Exames - Laboratoriais)
**Unidade:** mg/dL

**30 chunks de 11 artigos (avg similarity: 0.632)**

### Chunk 1/30
**Article:** 2024 Guidelines of the Polish Society of Laboratory Diagnostics and the Polish Lipid Association on laboratory diagnostics of lipid metabolism disorders (2024)
**Journal:** Archives of Medical Science
**Section:** results | **Similarity:** 0.723

50%Moderate< 100< 2.6Low< 115< 3.0Alarming levels:Suspected homozygous familial hypercholesterolaemia:In untreated individuals> 500> 13.0In treated individuals> 300> 7.8Suspected heterozygous familial hypercholesterolaemia> 190> 4.9Unit conversion: [mg/dl] × 0.026 = [mmol/l].less, as with the use of other formulas, the accu-racy of non-HDL-C calculation depends on the bi-ological and analytical variability of TC and HDL-C concentrations. However, the biological variability 
of HDL-C levels is much lower than that of other lipid parameters, especially TG. In addition, HDL-C concentrations are much lower than the TC levels, which minimises their eﬀect on changes in calcu-lated non-HDL-C concentrations.Reporting of resultsAlongside the calculated non-HDL-C level, a laboratory report should include information on 

Arch Med Sci 2, March / 2024 369the desirable (target) values with regard to cardio-vascular risk (Table VI).

---

### Chunk 2/30
**Article:** Cardiologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.663

em apresentar aterosclerose aos 50 anos.
- A heterogeneidade das partículas (estudo dos “11 tipos de LDL”) implica impacto aterogênico variável.
- Avaliação deve considerar modificações das lipoproteínas e o contexto clínico e metabólico.
### 2. Exames laboratoriais como desfechos substitutos e individualização
- Números isolados (p.ex., LDL < 100; CT < 200) não definem saúde nem garantem desfechos.
- Evitar tratar pela média estatística; cada indivíduo é um “exemplar genômico único”.
- Equilíbrio entre medicina tradicional e funcional integrativa: valorizar hábitos, sintomas, risco e imagem quando necessário.
### 3. Razão triglicerídeos/HDL como inferência prática de risco
- Regra prática: triglicerídeos aproximadamente 2,5 vezes o HDL sugerem maior proporção de LDL aterogênico.
- Classificação prática: 
  - Risco baixo em faixas como TG ~100–125 e HDL ~50.
  - Acima disso: risco médio a alto, conforme contexto.

---

### Chunk 3/30
**Article:** Cardiologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.662

roporção de LDL aterogênico.
- Classificação prática: 
  - Risco baixo em faixas como TG ~100–125 e HDL ~50.
  - Acima disso: risco médio a alto, conforme contexto.
- TG elevado (excluídas condições genéticas raras) indica consumo exagerado de carboidratos e desbalanço com individualidade (idade, metabolismo, exercício, polimorfismos).
### 4. Cadeia causal metabólica e estratégia de intervenção
- Fluxo típico: excesso de carboidratos → ↑ TG, CT, LDL, insulina de jejum, HbA1c.
- Maior impacto do colesterol endógeno derivado da dieta sobre produção hepática, não do colesterol alimentar direto.
- Foco terapêutico: reduzir a causa inicial (excesso de carboidratos), não apenas baixar números.
### 5. Exames de LDL oxidada e anti-LDL oxidada: interpretação e limitações
- LDL oxidada: medida direta do antígeno; mais fidedigna.
- Anti-LDL oxidada: anticorpos que podem não acompanhar a LDL ox por depuração/metabolização.

---

### Chunk 4/30
**Article:** Dislipdemia (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.661

maior que TG, mas ainda pode haver LDL de baixa qualidade, reforçando limitações do uso isolado da razão.
- Colesterol remanescente (IDL, VLDL, quilomícrons) estimado por total – (HDL+LDL): exemplo 220 – (40+150) = 30, destacando fração mais aterogênica.
- Oxidação de LDL: anti-LDL oxidado ideal até 25; valor de 27,5 indica maior estresse oxidativo e motiva estratégias para aumentar tamanho pico e reduzir oxidação.
**Benefício absoluto de fármacos/alternativas em prevenção primária é modesto; intervenções podem ser valiosas em prevenção secundária com metas moderadas e combinações.**
- Redução relativa média de risco com estatina ~20%; em baixo risco (2% em 10 anos), redução absoluta ~0,4% e NNT ~250 (4 por 1.000), benefício muito pequeno; em risco extremo (40%), redução absoluta 8% (40%→32%).

---

### Chunk 5/30
**Article:** 2024 Guidelines of the Polish Society of Laboratory Diagnostics and the Polish Lipid Association on laboratory diagnostics of lipid metabolism disorders (2024)
**Journal:** Archives of Medical Science
**Section:** results | **Similarity:** 0.651

by NCEP, is ±9% while the one used by COBJwDL is ±8% and this value is also recommended by PSLD/PoLA (2024).Reporting of resultsAlongside the TC level, a laboratory report should include information on the desirable (tar-get) values with regard to cardiovascular risk (Ta-
ble II).HDL cholesterolHigh density lipoproteins (HDL), unlike other lipoproteins, are characterised by a low lipid and high protein content. HDL transport about 25% of the cholesterol present in the blood, and its content 
in the particles of these lipoproteins varies consid-erably. Therefore, plasma HDL-C level provides in-direct and inaccurate information on HDL content in the blood. Nevertheless, HDL-C measurement re-mains a basic test for the assessment of HDL con-
tent in the blood, as methods of direct measure-ment of the number of HDL particles (HDL-P), and their individual subfraction (measured with e.g.

---

### Chunk 6/30
**Article:** Cardiologia VII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.650

o cardiovascular absoluto em 10 anos revela que um colesterol total elevado, como 310 mg/dL, pode não justificar o uso de estatinas quando o risco basal do paciente é baixo.**
- Uma paciente de 32 anos, apesar de ter um colesterol total de 310 mg/dL e LDL de 200 mg/dL, apresentou um risco cardiovascular em 10 anos de apenas 1,22%.
- A redução de risco relativa padrão de 20% proporcionada pela estatina resultaria em uma redução de risco absoluta de apenas 0,24%, diminuindo seu risco final para 0,98%, um benefício considerado insignificante.
- Em contraste, um paciente de 50 anos com colesterol total de 180 mg/dL, mas com outros fatores de risco (HDL de 29 mg/dL, pressão de 18 por 10), teve um risco cardiovascular em 10 anos calculado em 68%.
- Para este paciente de alto risco, a mesma redução relativa de 20% da estatina se traduz em uma redução absoluta de 13%, baixando seu risco para 56%.

---

### Chunk 7/30
**Article:** 2024 Guidelines of the Polish Society of Laboratory Diagnostics and the Polish Lipid Association on laboratory diagnostics of lipid metabolism disorders (2024)
**Journal:** Archives of Medical Science
**Section:** other | **Similarity:** 0.648

te: < 100 mg/dl (2.6 mmol/l)Low: < 115 mg/dl (3.0 mmol/l)> 500 mg/dl (13 mmol/l)  suspected hoFH> 190 mg/dl (4.9 mmol/l)  suspected heFHNon-HDL cholesterol Fasting and non-fasting:Cardiovascular risk:   Extreme: < 70 mg/dl (1.8 mmol/l)Very high: < 85 mg/dl (2.2 mmol/l)High: < 100 mg/dl (2.6 mmol/l)Low/Moderate: < 130 mg/dl (3.4 mmol/l)Apolipoprotein B (apoB)Fasting and non-fasting:Cardiovascular risk:Extreme: < 55 mg/dlVery high: < 65 mg/dlHigh: < 80 mg/dlLow/Moderate: < 100 mg/dlLipoprotein (a) [Lp(a)]Fasting and non-fasting:< 30 mg/dl (75 nmol/l) 3050 mg/dl (75125 nmol/l)  moderate risk> 50 mg/dl (125 nmol/l)  high risk > 180 mg/dl (450 nmol/l)  very high cardiovascular riskFH  familial hypercholesterolemia, F  female, M  male. At TG levels > 200 mg/dl (2.3 mmol/l), the LDL-C is not calculated. The equivalent indicator of cardiovascular risk is then the non-HDL-C or apoB level. When alarming values are detected, urgent medical consultation is indicated.

---

### Chunk 8/30
**Article:** 2024 Guidelines of the Polish Society of Laboratory Diagnostics and the Polish Lipid Association on laboratory diagnostics of lipid metabolism disorders (2024)
**Journal:** Archives of Medical Science
**Section:** methods | **Similarity:** 0.647

Desirable TC serum/plasma levels [4, 10]
ParameterTC [mg/dl]TC [mmol/l]Desirable levels fasting and non-fasting< 190< 4.9Unit conversion: [mg/dl] × 0.026 = [mmol/l].

364 Arch Med Sci 2, March / 2024
and do not provide suﬃcient new data to be rec-ommended. It appears that the best method (not available in practice) to assess HDL functionalities is to assess their cholesterol eﬄux capacity (CEC). This is an in vitro test that measures the ability of HDL to promote cholesterol removal from cho-lesterol donor cells such as macrophages. CEC is a predictor of cardiovascular risk independent of HDL-C concentration [34, 35].From a practical point of view, HDL-C level is not currently recommended as a treatment target 
or predictor of cardiovascular risk or for use in 
monitoring the treatment of lipid disorders [4, 5]. HDL-C level is used in the calculation of LDL-C and non-HDL-C levels. Methods of determinationThe HDL-C level is measured in serum or plas-ma.

---

### Chunk 9/30
**Article:** Cardiologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.646

ção triglicerídeos/HDL como inferência prática de risco: TG aproximadamente 2,5 vezes o HDL sugere maior presença de partículas aterogênicas de LDL.
- Ênfase de que números isolados (LDL total, LDL oxidada, subfracionamento) não são “bala de prata”; é preciso avaliar o conjunto (inflamação, oxidação, glicação, metilação).
- Sugere que mesmo com LDL oxidada baixa, podem existir outras modificações (LDL glicada, inflamada) com risco aterosclerótico.
- Destaca que reduzir apenas o número de colesterol sem abordar a cadeia causal (excesso de carboidratos, resistência insulínica) é insuficiente.
- Não há queixa específica do paciente registrada; o conteúdo é educacional, incluindo interpretação de exames e impacto de polimorfismos genéticos no metabolismo lipídico e risco cardiovascular.
## Objetivo:
- Referência a estudos indicando que cerca de 50% dos indivíduos com LDL “normal” (≤100 mg/dL) apresentam aterosclerose aos 50 anos.

---

### Chunk 10/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.646

, apoiando a personalização baseada em perfis de risco cumulativos, não em um único marcador.
**Achados Adicionais**
- O cardiômetro de mortalidade cardiovascular iniciou monitoramento em 11/6/2024, contextualizando a urgência de intervenções contínuas.
- Diretrizes de tratamento costumam usar um limiar de 7,5% de risco para estatinas, mas exemplos práticos ilustram faixas de 2–4% em que a decisão deve considerar CAC e ApoB.
- Quase 40% dos indivíduos com LDL muito elevado podem não apresentar aterosclerose, reforçando heterogeneidade do risco.
- Colesterol total de 300 frequentemente acompanha LDL >190, mas a decisão terapêutica deve ser guiada por risco global.
- Revisões de 2019 na Annals of Internal Medicine sobre carne vermelha/processada fornecem contexto adicional para dietas cardiometabólicas.
- Editorial de 2020 sobre hipercolesterolemia familiar difundiu o conceito “Power of Zero”, ampliando o uso do CAC na estratificação.

---

### Chunk 11/30
**Article:** 2024 Guidelines of the Polish Society of Laboratory Diagnostics and the Polish Lipid Association on laboratory diagnostics of lipid metabolism disorders (2024)
**Journal:** Archives of Medical Science
**Section:** results | **Similarity:** 0.640

(2019) and the Europe-
an Federation of Clinical Chemistry and Laboratory 
Medicine (EFLM) of 2016, which was based on data 
indicating that a slight postprandial increase in the 
TG level (to 27 mg/dl (0.3 mmol/l) does not result in 
signiﬁcant changes in other measurements, or in the lipid proﬁle assessment compared to the test-ing of fasting samples. Small diﬀerences in the in-terpretation of results pertain to TG and non-HDL-C. It is recommended that lipid proﬁle be repeated in the fasting sample if the non-fasting TG level is  > 400 mg/dl (4.5 mmol/l) [2, 10].
RecommendationsRoutine lipid profile testing, primarily LDL-C and TC, does not require fasting samples. Re-testing for correct 
measurement of LDL cholesterol using material collect-ed in the fasting state should be considered if non-fast-
ing TG is > 400 mg/dl (4.5 mmol/l).The levels of individual lipid proﬁle compo-nents are characterised by intra-individual vari-ability of 510% for TC and > 20% for TG.

---

### Chunk 12/30
**Article:** 2024 Guidelines of the Polish Society of Laboratory Diagnostics and the Polish Lipid Association on laboratory diagnostics of lipid metabolism disorders (2024)
**Journal:** Archives of Medical Science
**Section:** results | **Similarity:** 0.637

and their remnants and 
a signiﬁcant risk of ASCVD dependent on them, 
were considered hypertriglyceridaemia (Table I) [26, 27]. Such a division is also currently recom-mended by PSLD/PoLA experts (2024).

Arch Med Sci 2, March / 2024 363
RecommendationsIncreased plasma/serum TG levels are associated with signiﬁcant cardiovascular risk dependent, among oth-
ers, on the accumulation of TRL and their remnants. 
Moderately increased fasting TG levels > 150 mg/
dl (1.7 mmol/l) are an indication for treatment of hy-
pertriglyceridaemia, which should aim for TG levels  < 100 mg/dl (1.1 mmol/l).Total cholesterolCholesterol is one of the best-known lipids, which results, among others, from its direct con-nection with the development of atherosclerosis.

---

### Chunk 13/30
**Article:** Cardiologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.635

lipídico e risco cardiovascular.
## Objetivo:
- Referência a estudos indicando que cerca de 50% dos indivíduos com LDL “normal” (≤100 mg/dL) apresentam aterosclerose aos 50 anos.
- Proposta de avaliação laboratorial:
  - Colesterol total, HDL, triglicerídeos, LDL (com possibilidade de subfracionamento).
  - Insulina de jejum, glicemia de jejum, hemoglobina glicada.
  - LDL oxidada direta; considerar anticorpos anti-LDL oxidada quando a direta não estiver disponível (menos fidedigno, porém informativo sobre resposta imune).
  - Subfracionamento de LDL (tamanho/densidade das partículas), reconhecendo limitações.
  - Apolipoproteínas: ApoA (predominante em HDL) e ApoB (predominante em LDL); maior razão ApoA/ApoB sugere melhor perfil de risco.
- Considerar angiotomografia de coronárias com escore de cálcio para avaliar placas e calcificações quando exames laboratoriais sugerirem alto risco, reconhecendo possíveis discordâncias entre exames séricos e imagem.

---

### Chunk 14/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.630

co da dislipidemia.
- [ ] 2. Analisar criticamente os estudos sobre o Inclisiran (Cibrava), focando na diferença entre desfechos substitutos (redução de LDL) e desfechos clínicos duros (mortalidade, infarto, AVC).
- [ ] 3. Utilizar o site `the-nnt.com` para pesquisar o NNT e NNH de outros medicamentos prescritos na prática clínica.
- [ ] 4. Refletir sobre a influência da indústria farmacêutica e dos interesses financeiros na prescrição de novos medicamentos.
- [ ] 5. Estudar a diferença entre risco relativo e risco absoluto para interpretar criticamente os achados de estudos científicos.
- [ ] 6. Pesquisar sobre os diferentes tipos de ômega-3 (EPA, DHA, ALA) e suas formulações para entender como a qualidade do suplemento afeta os resultados.
- [ ] 7. Investigar os exames de subfracionamento de LDL e a relação ApoA/ApoB como ferramentas de avaliação de risco cardiovascular mais precisas que o LDL total.
- [ ] 8.

---

### Chunk 15/30
**Article:** 2024 Guidelines of the Polish Society of Laboratory Diagnostics and the Polish Lipid Association on laboratory diagnostics of lipid metabolism disorders (2024)
**Journal:** Archives of Medical Science
**Section:** introduction | **Similarity:** 0.625

ffective monitoring (after therapy 
initiation), and in the consequence to avoid the first and recurrent cardio-
vascular events.Key words: lipid disorders, low-density lipoprotein cholesterol,  high-density lipoprotein cholesterol (HDL-C), triglicerides, non-HDL-C, guidelines, laboratory diagnostics.
Editors choice 

358 Arch Med Sci 2, March / 2024IntroductionThe lipid proﬁle routinely performed to assess cardiovascular risk involves the measurement/calculation of serum/plasma levels of total cho-
lesterol (TC), high-density lipoprotein cholester-ol (HDL-C), low-density lipoprotein cholesterol (LDL-C), triglycerides (TG) and non-HDL choles-terol (non-HDL-C), although LDL-C level is still the most important factor in both the diagnosis and 
monitoring of the course and treatment of lipid 
disorders and the prediction of cardiovascular in-
cidents [14].

---

### Chunk 16/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XIV (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.623

- [ ] 2. Para pacientes com síndrome metabólica, considerar estratégia Mediterrânea organizada; preferir versão com restrição calórica combinada com atividade física, monitorando resultados aos 6 e 12 meses.
- [ ] 3. Solicitar perfil lipídico ampliado em casos selecionados: HDL, triglicerídeos, insulina, PCR, LDL oxidado e, quando indicado, subfracionamento de LDL; evitar decisões baseadas apenas em LDL total.
- [ ] 4. Planejar exames de risco cardiovascular conforme necessidade: score de cálcio coronariano e angiotomografia de coronárias (incluindo avaliação de placas moles) quando o contexto clínico justificar.
- [ ] 5. Revisar e atualizar protocolos internos sobre álcool: remover recomendações de consumo “cardioprotetor”; educar pacientes sobre riscos de câncer e piora do sono; avaliar predisposição/sensibilidade individual.
- [ ] 6.

---

### Chunk 17/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.622

acientes com doença cardiovascular prévia, o ômega-3 mostrou uma redução absoluta de risco de infarto de 3% (30 vidas salvas a cada 1.000 pessoas).
    - **Vieses do Estudo:** O estudo não controlou tipo, qualidade ou dose do ômega-3, baseou-se em autorrelato e usou uma população (UK Biobank) mais saudável que a média, o que limita a generalização dos resultados.
*   **Complexidade do Colesterol e Perfil Lipídico**
    - **Paradoxo do HDL Alto:** Níveis muito altos de HDL podem, paradoxalmente, aumentar o risco cardiovascular, mostrando que a relação não é linear.
    - **Qualidade vs. Quantidade do LDL:** A qualidade das partículas de LDL (tamanho e densidade) é um preditor de risco mais importante que a quantidade total. Partículas pequenas e densas (small dense LDL) são mais aterogênicas.

---

### Chunk 18/30
**Article:** Dislipidemias II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.620

dade, inatividade física, pressão arterial e dislipidemia.
- A inflamação é indicada por marcadores como a Proteína C-Reativa (PCR), onde um valor de 5 já é considerado elevado, e a Interleucina 6 (exemplo de paciente com 8.45).
- A resposta anti-inflamatória é medida pela Interleucina 10, com um valor de corte de 3,5 (abaixo disso é um risco) e um exemplo de paciente com 6.44.
- A saúde do endotélio é avaliada pelo óxido nítrico, que deve estar em 8.8; um paciente com 4.8 já apresenta disfunção sistêmica.
- A LDL oxidada é um marcador crítico, com um limite saudável de 133, enquanto pacientes de alto risco podem apresentar valores extremos, como 1000.
**Achados Adicionais**
- Existem cinco parâmetros clássicos para definir a síndrome metabólica.
- O alvo de LDL para pacientes de alto risco, segundo as diretrizes atuais, é de 50.

---

### Chunk 19/30
**Article:** 2024 Guidelines of the Polish Society of Laboratory Diagnostics and the Polish Lipid Association on laboratory diagnostics of lipid metabolism disorders (2024)
**Journal:** Archives of Medical Science
**Section:** abstract | **Similarity:** 0.618

34. Clin Chem 1993; 39: 1127.23. Yang N, Wang M, Liu J, Liu J, Hao Y, Zhao D; Ccc-Acs In-
vestigators. The level of remnant cholesterol and impli-
cations for lipid-lowering strategy in hospitalized pa-
tients with acute coronary syndrome in China: ﬁndings 
from the improving care for cardiovascular disease in 
China-Acute Coronary Syndrome Project. Metabolites 
2022; 12: 898. 24. Doi T, Langsted A, Nordestgaard BG. Elevated remnant 
cholesterol reclassiﬁes risk of ischemic heart disease 
and myocardial infarction. J Am Coll Cardiol 2022; 79: 
2383-97. 25. Quispe R, Martin SS, Michos ED, et al. Remnant choles-
terol predicts cardiovascular disease beyond LDL and 
ApoB: a primary prevention study. Eur Heart J 2021; 42: 
4324-32. 26. Moulin P, Dufour R, Averna M, et al. Identiﬁcation and di-
agnosis of patients with familial chylomicronaemia syn-
drome (FCS): Expert panel recommendations and propos-al of an FCS score. Atherosclerosis 2018; 275: 265-72.27. Ginsberg HN, Packard CJ, Chapman MJ, et al. Triglycer-ide-rich lipoproteins and their remnants: metabolic 
insights, role in atherosclerotic cardiovascular disease, and emerging therapeutic strategies  a consensus 
statement from the European Atherosclerosis Society. 
Eur Heart J 2021; 42: 4791-806.28. Rynkiewicz A, Cybulska B, Banach M, et al. Management of familial heterozygous hypercholesterolemia: position paper of the Polish Lipid Expert Forum. J Clin Lipidol 2013; 7: 217-21.29. Marx N, Federici M, Schütt K, et al.; ESC Scientiﬁc Docu-ment Group. 2023 ESC Guidelines for the management 
of cardiovascular disease in patients with diabetes. Eur 
Heart J 2023; 44: 4043-140. 30. Visseren FLJ, Mach F, Smulders YM, et al.; ESC National 
Cardiac Societies; ESC Scientiﬁc Document Group. 2021 ESC Guidelines on cardiovascular disease prevention in clinical practice. Eur Heart J 2021; 42: 3227-337. 31. Li LH, Dutkiewicz EP, Huang YC, et al. Analitycal methods for cholesterol quantiﬁcation. J Food Drug Ann 2019; 27: 375-86.32. Lopes-Virella MF, Stone P, Ellis S, Colwell JA. Cholesterol 
determination in high-density lipoproteins separated by three diﬀerent methods. Clin Chem 1977; 23: 882-4.33. Allain CC, Poon LS, Chan CS, et al. Enzymatic determi-nation of total serum cholesterol. Clin Chem 1974; 20: 
470-5.34. Ganjali S, Mahdipour E, Aghaee-Bakhtiari SH, et al. Com-
positional and functional properties of high-density li-
poproteins in relation to coronary in-stent restenosis. 
Arch Med Sci 2021; 19: 57-72. 35. Otocka-Kmiecik A, Mikhailidis DP, Nicholls SJ, Davidson M, Rysz J, Banach M. Dysfunctional HDL: a novel import-ant diagnostic and therapeutic target in cardiovascular disease? Prog Lipid Res 2012; 51: 314-24.36. Warnick GR, Nauck M, Rifai N. Evolution of methods for measurement of HDL-cholesterol: from ultracentrif-
ugation to homogeneous assays. Clin Chem 2001; 47: 
1579-96.37. Camont L, Chapman MJ, Kontush A. Pendal activities of 
HDL subpopulations and their relevance to cardiovascu-
lar disease. Trends Mol Med 2011; 17: 596-605.38. Martin SS, Jones SR, Toth PP. High-density lipoprotein 
subfractions: current views and clinical practice applica-tions. Trends Mol Med 2014; 26: 328-36.39. Kosmas CE, Martinez I, Sourlas A, et al. High-densi-
ty lipoprotein (HDL) functionality and its relevance to atherosclerotic cardiovascular disease. Drugs Context 
2018; 7: 212-25.40. Movvo R, Rader DJ. Laboratory assessment of HDL het-
erogeneity and function. Clin Chem 2008; 54: 788-801.41. Rosenson RS, Brewer HB, Chapman MJ, et al. HDL mea-
sures, particle heterogeneity, proposed nomenclature, 
and relation to atherosclerotic cardiovascular events. 
Clin Chem 2011; 57: 392-410.42. Sean Davidson W. HDL-C vs HDL-P: how changing one 
letter could make a diﬀerence in understanding the role 
of high-density lipoprotein in disease. Clin Chem 2014; 
60: e1-3.43. Kidawa M, Gluba-Brzózka A, Zielinska M, et al. Choles-
terol subfraction analysis in patients with acute coro-
nary syndrome. Curr Vasc Pharmacol 2019; 17: 365-75.44. Rizzo M, Otvos J, Nikolic D, et al. Subfractions and sub-
populations of HDL: an update. Curr Med Chem 2014; 
21: 2881-91.45. Sonmez A, Nikolic D, Dogru T, et al. Low- and high-den-sity lipoprotein subclasses in subjects with nonalcoholic fatty liver disease. J Clin Lipidol 2015; 9: 576-82.46. Annema W, von Eckardstein A. Dysfunctional high-densi-ty lipoproteins in coronary heart disease: implication for diagnostics and therapy. Translat Res 2016; 173: 30-57.47. Otocka-Kmiecik A, Mikhailidis DP, Nicholls SJ, et al. Dys-
functional HDL: a novel important diagnostic and ther-

---

### Chunk 20/30
**Article:** 2024 Guidelines of the Polish Society of Laboratory Diagnostics and the Polish Lipid Association on laboratory diagnostics of lipid metabolism disorders (2024)
**Journal:** Archives of Medical Science
**Section:** discussion | **Similarity:** 0.618

. Curr Atheroscler Rep 2017; 19: 31.19. Boekholdt SM, Arsenault BJ, Mora S, et al. Association 
of LDL cholesterol, non-HDL cholesterol, and apolipopro-tein B levels with risk of cardiovascular events among 
patients treated with statins: a meta-analysis. JAMA 
2012; 307: 1302-9. 20. Park JK, Bafna S, Forrest IS, et al. Phenome-wide Men-
delian randomization study of plasma triglyceride levels 
and 2600 disease traits. Elife 2023; 12: e80560.21. Trinder P. Determination of glucose in blood using glu-
cose oxidase with an alternative oxygen acceptor. Ann 
Clin Biochem 1969; 6: 24-7.22. Siedel J, Schmuck R, Staepels J, et al. Long term stable, liquid ready-to-use monoreagent for the enzymatic as-say of serum or plasma triglycerides (GPO-PAP-method). AACC Meeting Abstract 34. Clin Chem 1993; 39: 1127.23. Yang N, Wang M, Liu J, Liu J, Hao Y, Zhao D; Ccc-Acs In-
vestigators.

---

### Chunk 21/30
**Article:** Dislipdemia (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.612

ominância aterogênica de Apo B.
- Score de cálcio: 0 indica baixíssimo risco e sustenta não usar estatina em prevenção primária; faixas 0–100 (baixo), 100–300 (moderado), >300 (alto), com interpretação dependente da idade (ex.: 38 aos 36 anos é alto, aos 70 seria baixo).
- Classificação de risco em 10 anos: 0–5% baixo, 5–10% moderado, 10–20% alto, >20% muito alto; diretriz atual para LDL em baixo risco sugere <130 mg/dL (por vezes <100), mas a decisão deve integrar cálcio e partículas.
- Em estudo antigo (1984), mortalidade mínima ocorreu com colesterol total 200–240 mg/dL, com mortalidade semelhante acima de 240 versus abaixo de 200 em indivíduos sem DCV, desafiando metas rígidas de colesterol total isolado.

---

### Chunk 22/30
**Article:** Dislipdemia (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.612

quais ~250 mil ocorrem antes dos 60 anos.
- Brasil: 410 mil mortes/ano por DCV; 14 milhões com alguma DCV; 36% dos óbitos ≥55 anos decorrem de doença cardio-circulatória.
- Fisiopatologia: fluxo arterial torna-se turbilhonado ~50% de estenose, aumentando estresse de parede e risco de ruptura de placa.
**Achados Adicionais**
- LDL alvo em baixo risco: diretriz sugere <130 mg/dL (por vezes <100), mas o número isolado não determina benefício sem DCV prévia e sem avaliação de partículas/cálcio.
- LDL basal em estudo: 190 mg/dL; colesterol total 275 e HDL 31, ilustrando perfis iniciais elevados; em pacientes com cálcio zero, LDL >240 não alterou mortalidade/infarto com estatina.

---

### Chunk 23/30
**Article:** Cardiologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.612

pidemiologia.
- Limiar epidemiológico citado: abaixo de 2% (RR ≈ 1,02) não se considera robusto; um RR de 1,07 é insuficiente para conclusões firmes.
- Ômega‑3 em 1 grama/dia é dose apontada como insuficiente para efeitos terapêuticos, sugerindo necessidade de doses maiores ou mudança dietética.
- Estudo sobre estatinas (2016) menciona potenciais disfunções e ausência de benefícios em certos mecanismos, reforçando evidência heterogênea.
**Achados Adicionais**
- Corpus de estudos sobre gorduras (Annals 2014): 32 observacionais de consumo, 17 com biomarcadores sanguíneos e 27 ensaios prospectivos randomizados com suplementação de diferentes ácidos graxos.

---

## Teaching Note

Data e Hora: 2025-11-20 20:41:16
Local: [Inserir Local]
Aula: Cardiologia Metabólica Funcional Integrativa
## Visão Geral
A aula desconstruiu a teoria do colesterol como principal causa de doenças cardiovasculares por meio de análise crítica de estudos históricos e recentes.

---

### Chunk 24/30
**Article:** Cardiologia VII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.611

estatinas e prevenção cardiovascular.
- Ponto 1: O objetivo da prevenção deve ser desfechos clínicos (infarto, AVC, morte), não a alteração de um exame de LDL.
- Ponto 3 e 4: A eficácia da estatina não varia com o nível de colesterol; não há evidências para ajustar doses visando metas de LDL.
- Ponto 7: O risco deve ser determinado por calculadoras, não apenas pelo nível de colesterol.
- Ponto 8 e 9: Exceções das diretrizes (LDL > 190), mas com falta de evidências que sustentem tratamento agressivo e indiscriminado nesse grupo.
- Ponto 10 e 11: LDL é um preditor ruim isolado: metade das pessoas com LDL normal têm aterosclerose, e metade das com hipercolesterolemia familiar não têm.
- Ponto 12: Sugestão do uso do escore de cálcio coronariano como ferramenta superior para estratificação de risco, apesar de não constar nas diretrizes do SUS por ser exame de tomografia.

---

### Chunk 25/30
**Article:** Total cholesterol/HDL-cholesterol ratio discordance with LDL-cholesterol and non-HDL-cholesterol and incidence of atherosclerotic cardiovascular disease in primary prevention: The ARIC study (2020)
**Journal:** European Journal of Preventive Cardiology
**Section:** discussion | **Similarity:** 0.611

2008;372:224–33. [PubMed: 18640459] 
11. Mora S, Otvos JD, Rifai N, Rosenson RS, Buring JE, Ridker PM. Lipoprotein particle profiles by nuclear magnetic resonance compared with standard lipids and apolipoproteins in predicting incident cardiovascular disease in women. Circulation 2009;119:931–9. [PubMed: 19204302] 
12. Prospective Studies C, Lewington S, Whitlock G, Clarke R, Sherliker P, Emberson J, Halsey J, Qizilbash N, Peto R, Collins R. Blood cholesterol and vascular mortality by age, sex, and blood pressure: a meta-analysis of individual data from 61 prospective studies with 55,000 vascular deaths. Lancet 2007;370:1829–39. [PubMed: 18061058] 
13. Ingelsson E, Schaefer EJ, Contois JH, McNamara JR, Sullivan L, Keyes MJ, Pencina MJ, Schoonmaker C, Wilson PW, D’Agostino RB, Vasan RS. Clinical utility of different lipid measures for prediction of coronary heart disease in men and women. JAMA 2007;298:776–85. [PubMed: 17699011] 
14.

---

### Chunk 26/30
**Article:** 2024 Guidelines of the Polish Society of Laboratory Diagnostics and the Polish Lipid Association on laboratory diagnostics of lipid metabolism disorders (2024)
**Journal:** Archives of Medical Science
**Section:** other | **Similarity:** 0.610

if the pre-analytical 
requirements are met, tends to underestimate LDL-C levels at low values < 70 mg/dl (1.8 mmol/l). 
GPxFigure 6. Dysfunctional HDL particles: A  HDL modiﬁed by myeloperoxidase, B  inﬂammatory HDL SAA  serum amyloid A, PON-1  paraoxonase-1, GPx  glutathione peroxidase, RCT  reverse cholesterol transport,  ABCA1  ATP-binding membrane cassette transporter A1.ABMPOApo A-IPON-1GPxApo A-IApo A-II↓ Macrophage RCT↓ ABCA1Apo A-IISAACeruloplasminApo A-IApo A-IIApo JSAACeruloplasminApo A-IPON-1

Arch Med Sci 2, March / 2024 367One of the recent modiﬁcations proposed to Friedewald formula is he Martin-Hopkins equation (2013) [59]: LDL-C = TC  HDL-C  TG/x (in mg/dl), where: x  is the TG-VLDL-C ratio based on TG and non-HDL-C levels; values are available in spe-
cial tables or online calculators (https://ldlcalcula-tor.com).

---

### Chunk 27/30
**Article:** Cardiologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.610

iovascular isolado; ajustar dieta, atividade e exames conforme genótipo.
- LDLR: mutações causam hipercolesterolemia familiar; maior peso a score de cálcio e controle de inflamação/glicação na condução.
- CETP: transfere ésteres de colesterol da HDL; variações podem elevar HDL disfuncional; HDL alto não implica necessariamente proteção.
- ABCG5/ABCG8: transportadores de esteróis; polimorfismos aumentam colesterol e predisposição a ateroma; ajustar colesterol dietético e gorduras saturadas; atenção a respostas paradoxais em low carb/cetogênica.
- HMGCR: via do mevalonato; polimorfismos associam-se a pior resposta a estatinas e menor produção de ubiquinona; considerar suplementação de coQ10/ubiquinol.
- LIPC: lipase hepática dual; fenótipo de HDL baixo e LDL/CT altos; melhor resposta com controle de gordura saturada e treino de resistência.

---

### Chunk 28/30
**Article:** Cardiologia I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.608

- Critérios de triagem clínica: histórico dietético, triglicerídeos elevados, esteatose hepática, aumento de ácido úrico, circunferência abdominal.
### 6. História do colesterol, diretrizes e estatinas
- 1784: identificação do colesterol; 1870 (Virchow): presença em placas ateroscleróticas.
- 1921: modelos com coelhos herbívoros alimentados com colesterol levaram a interpretações inadequadas.
- 1925: 80–90% do colesterol é produzido endogenamente.
- 1949 (Goffman): tipos de colesterol via ultracentrífuga.
- 1953 (Ancel Keys): Estudo dos “Sete Países” com seleção de países compatíveis com sua hipótese; forte influência em diretrizes.
- 1954 (OMS) e 1961 (AHA): comissões e guidelines; anos 1970 popularizam HDL “bom” e LDL “ruim”.
- 1976–1977: politização da teoria do colesterol (McGovern; Congresso).
- 1979: correlação LDL-DCV não implica causalidade; 1980: consolidação do dogma.

---

### Chunk 29/30
**Article:** 2024 Guidelines of the Polish Society of Laboratory Diagnostics and the Polish Lipid Association on laboratory diagnostics of lipid metabolism disorders (2024)
**Journal:** Archives of Medical Science
**Section:** results | **Similarity:** 0.608

 s modiﬁcation. 1Based on the PSC/PoLA 2024 Guidelines [81].Lipid proﬁle  laboratory reportLipid proﬁle includes a battery of blood serum or plasma tests and calculations aimed at identiﬁ-cation of dyslipidemia as a cardiovascular risk fac-tor, deﬁning the recommendations and  treatment  
monitoring, including: total cholesterol (TC) level, HDL cholesterol level (HDL-C), LDL cholesterol level (LDL-C), non-HDL cholesterol level (non-HDL-C),
 triglyceride (TG) level,  lipoprotein (a) level [Lp(a)] (determined at least once in life  see PCS/PoLA 2024 recommenda-
tions [81]), apolipoprotein B (apoB) level  as indicated.In addition  to the results of measurements and calculations, a lipid proﬁle laboratory report  (Table IX), should include information on how the LDL-C level was determined (calculated/deter-mined), as well as the target (desirable) and alarm-Table IX.

---

### Chunk 30/30
**Article:** 2024 Guidelines of the Polish Society of Laboratory Diagnostics and the Polish Lipid Association on laboratory diagnostics of lipid metabolism disorders (2024)
**Journal:** Archives of Medical Science
**Section:** results | **Similarity:** 0.608

may provide a strong justiﬁcation for its preferential use. Unfortunately, most clinical labo-ratories continue to use the Friedewald equation, which is fraught with many ﬂaws and often un-derestimates results, so there is an urgent need for an improved education on the subject, as well 
as eﬀorts to implement new formulas [63, 64]. The LDL-C level can be determined using di-rect (homogeneous) methods. Current third-gen-
eration methods involve the use of reagents 
containing various detergents, surfactants, car-bohydrate derivatives or other agents that block 
or dissolve individual lipoprotein fractions, se-lectively making LDL-C available for cholesterol esterase and oxidase. These methods allow the use of automated analysers. Due to considerable methodological variability, direct methods for the determination of LDL-C vary in terms of the accuracy (traceability to the reference method) and precision of assays [62].

---

