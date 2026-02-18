# ScoreItem: Doppler Aorta - PSV Artéria Renal

**ID:** `c77cedd3-2800-7068-927c-db125905c0d3`
**FullName:** Doppler Aorta - PSV Artéria Renal (Exames - Imagem)
**Unit:** cm/s

**Preparation Metadata:**
- Quality Grade: **GOOD**
- Total Chunks: 30 de 17 artigos
- Avg Similarity: 0.507

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `c77cedd3-2800-7068-927c-db125905c0d3`.**

```json
{
  "score_item_id": "c77cedd3-2800-7068-927c-db125905c0d3",
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

**ScoreItem:** Doppler Aorta - PSV Artéria Renal (Exames - Imagem)
**Unidade:** cm/s

**30 chunks de 17 artigos (avg similarity: 0.507)**

### Chunk 1/30
**Article:** Critical analysis of renal duplex ultrasound parameters in detecting significant renal artery stenosis (2012)
**Journal:** Journal of Vascular Surgery
**Section:** abstract | **Similarity:** 0.704

Large study of 313 patients evaluating Doppler parameters for renal artery stenosis. Mean renal-aortic ratios for normal, <60%, and ≥60% stenosis were 2.2, 2.9, and 4.5 respectively. RAR >3.5 demonstrated high diagnostic accuracy for detecting hemodynamically significant stenosis.

---

### Chunk 2/30
**Article:** Doppler ultrasound and renal artery stenosis: An overview (2013)
**Journal:** Journal of Ultrasound
**Section:** abstract | **Similarity:** 0.695

Overview of Doppler ultrasound techniques for diagnosing renal artery stenosis, the most common cause of secondary hypertension. Discusses renal-aortic ratio as a reliable parameter that normalizes individual hemodynamic variations, improving diagnostic specificity compared to peak systolic velocity alone.

---

### Chunk 3/30
**Article:** Doppler Renal Assessment, Protocols, and Interpretation (2024)
**Journal:** StatPearls
**Section:** abstract | **Similarity:** 0.672

Comprehensive review of renal Doppler ultrasound techniques, including renal-aortic ratio (RAR) for detecting renal artery stenosis. The RAR compares intrastenotic flow velocity in renal arteries with aortic reference values, with RAR >3.5 predicting ≥60% stenosis with 84-91% sensitivity and 95-97% specificity.

---

### Chunk 4/30
**Article:** Optimal Peak Systolic Velocity Thresholds for Predicting Internal Carotid Artery Stenosis Greater than or Equal to 50%, 60%, 70%, and 80% (2016)
**Journal:** Journal of Stroke and Cerebrovascular Diseases
**Section:** abstract | **Similarity:** 0.567

The research established optimal peak systolic velocity measurements for detecting various degrees of internal carotid artery stenosis. Testing 127 arterial specimens, researchers identified specific thresholds: 130 cm/s, 160 cm/s, 200 cm/s, and 270 cm/s for detecting stenosis at increasing severity levels (≥50%, ≥60%, ≥70%, and ≥80% respectively). The findings demonstrated high diagnostic accuracies across all measured thresholds, with sensitivity and specificity values exceeding 85% for each threshold category.

---

### Chunk 5/30
**Article:** Correlation between Ultrasound Peak Systolic Velocity and Angiography for Grading Internal Carotid Artery Stenosis (2024)
**Journal:** Journal of Clinical Medicine
**Section:** abstract | **Similarity:** 0.541

The study evaluated how peak systolic velocity (PSV) measured via duplex ultrasound correlates with angiography findings for assessing internal carotid artery stenosis. Researchers analyzed 47 stenotic lesions using both ultrasound and digital subtraction angiography, applying NASCET and ECST classification protocols. Key findings indicated that a PSV threshold of 200 cm/s was found to be the best criterion for identifying severe NASCET stenoses (≥70%), while a threshold of 180 cm/s was the best for ECST stenoses (≥80%). However, PSV demonstrated limited reliability for moderate stenoses, suggesting complementary imaging techniques should be employed in such cases.

---

### Chunk 6/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.527

ardiovasculardiseaseandantithrombotictherapy.EurHeartJ.2013;34:1708–1713,1713a–1713b.697.JamesS,BudajA,AylwardP,etal.Ticagrelorversusclopidogrelinacutecoronarysyndromesinrelationtorenalfunction:resultsfromthePlateletInhibitionandPatientOutcomes(PLATO)trial.Circulation.2010;122:1056–1067.698.HerringtonWG,StaplinN.InpatientswithcoronarydiseaseandCKD,
addinganinvasivestrategytoMTdidnotimproveoutcomes.AnnInternMed.2020;173:JC16.699.SarnakMJ,AmannK,BangaloreS,etal.Chronickidneydiseaseand
coronaryarterydisease:JACCstate-of-the-artreview.JAmCollCardiol.2019;74:1823–1838.700.CharytanDM,WallentinL,LagerqvistB,etal.Earlyangiographyin
patientswithchronickidneydisease:acollaborativesystematicreview.ClinJAmSocNephrol.2009;4:1032–1043.701.HastingsRS,HochmanJS,DzavikV,etal.Effectoflaterevascularization
ofatotallyoccludedcoronaryarteryaftermyocardialinfarctiononmortalityratesinpatientswithrenalimpairment.AmJCardiol.2012;110:954–960.702.JohnstonN,JernbergT,LagerqvistB,etal.Earlyinvasivetrea

---

### Chunk 7/30
**Article:** Dislipdemia (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.527

quais ~250 mil ocorrem antes dos 60 anos.
- Brasil: 410 mil mortes/ano por DCV; 14 milhões com alguma DCV; 36% dos óbitos ≥55 anos decorrem de doença cardio-circulatória.
- Fisiopatologia: fluxo arterial torna-se turbilhonado ~50% de estenose, aumentando estresse de parede e risco de ruptura de placa.
**Achados Adicionais**
- LDL alvo em baixo risco: diretriz sugere <130 mg/dL (por vezes <100), mas o número isolado não determina benefício sem DCV prévia e sem avaliação de partículas/cálcio.
- LDL basal em estudo: 190 mg/dL; colesterol total 275 e HDL 31, ilustrando perfis iniciais elevados; em pacientes com cálcio zero, LDL >240 não alterou mortalidade/infarto com estatina.

---

### Chunk 8/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.500

nsiderado alterado (normal > 15 ml/s).
- O tratamento medicamentoso para HPB inclui doses como 0,4 mg de Tansulosina ou 2 a 4 mg de Doxazosina para relaxar a musculatura e melhorar o fluxo urinário.
### Achados Adicionais Chave
- A ressonância magnética multiparamétrica 3 Tesla é uma ferramenta de alta definição, recomendada a partir dos 40 anos para homens com histórico familiar ou a partir dos 50 anos como rotina para uma avaliação prostática precisa.
- A dosagem hormonal salivar oferece uma medição precisa dos hormônios livres, com faixas de referência para testosterona (47-150), estradiol (0.6-3) e o quociente estrogênico (0.04-1.67), que avalia o equilíbrio hormonal.
- A escala Gleason, que vai de 1 a 10, mede a agressividade do câncer de próstata, com tumores classificados como 8, 9 ou 10 sendo considerados os mais agressivos.
- A hiperplasia prostática é uma condição que afeta apenas 3 espécies: humanos, cães e macacos.

---

### Chunk 9/30
**Article:** Hipertensão Arterial Sistêmica (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.496

nal (sistema renina-angiotensina-aldosterona - SRAA), renal, microbiota e endotélio.
    *   **Mecanismos de Controle:** Rápidos (neurais), médio prazo (hormonais, alvo principal dos fármacos) e longo prazo (controle de volemia pelo rim).
*   **Diagnóstico e Classificação:**
    *   A medição em consultório é criticada; recomenda-se MAPA (Monitorização Ambulatorial) ou MRPA (Monitorização Residencial) para um diagnóstico preciso.
    *   **Valores de Referência para Diagnóstico:** ≥ 140/90 mmHg (consultório), ≥ 130/80 mmHg (MAPA 24h), ≥ 135/85 mmHg (MRPA).
    *   **Classificação (a partir de 18 anos):**
        *   **Ótima:** < 120/80 mmHg.
        *   **Normal:** 120-129 / 80-84 mmHg.
        *   **Pré-hipertensão:** 130-139 / 85-89 mmHg.
        *   **Hipertensão Estágio 1:** 140-159 / 90-99 mmHg.
        *   **Hipertensão Estágio 2:** 160-179 / 100-109 mmHg.
        *   **Hipertensão Estágio 3:** ≥ 180/110 mmHg.
### 3.

---

### Chunk 10/30
**Article:** Peak systolic velocity ratio for evaluation of internal carotid artery stenosis correlated with plaque morphology: substudy results of the ANTIQUE study (2023)
**Journal:** Frontiers in Neurology
**Section:** abstract | **Similarity:** 0.494

This study evaluated how effectively four duplex sonography measurements assess carotid artery narrowing severity when compared against computed tomography angiography. The research examined 143 patients with significant stenosis, analyzing peak systolic velocity (PSV), velocity ratios, end-diastolic velocity, and B-mode imaging. Results demonstrated that the PSV ICA/CCA ratio showed the highest correlation with CTA, followed by PSV and other parameters. The study found that plaque composition significantly influenced measurement accuracy, with calcified plaques producing substantially weaker correlations than softer plaque types with smooth surfaces.

---

### Chunk 11/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.487

ustold,butvery
old,andincorporatesmoreandmorepeopleovertheageof80.Octo-andnonagenariansoftendemonstratedistinctpatternsofdiseasecomplexity.Thesefeaturesincludemulti-
morbidityoftenaccompaniedbypolypharmacy,frailty,
cognitiveimpairment,andgerontopsychiatricdisordersamongothers.Often,severalofthesefeaturescoexistespe-
ciallyinolderadultswithCKD.ImplicationsforagingadultswithCKDareimportantinbothdiagnosisandtreatment.Theinterpretationoflabora-toryresults(specicallySCr)usedinthestagingsystemshouldfactorinanolderadult’shabitusgiventhefrequencyofsarcopenia.Acreatinine-basedeGFR(eGFRcr)willover-estimateGFRintheelderly(andothers)withsarcopenia
leadingtodrugoverdosing.UrineACRatthesametimewill
befalselyhighduetothefalselylowcreatinineinthede-
nominator.Furthermore,thepresenceoffrailtymayalter
treatmenttargetsrecommendedforyoungerpeoplewithCKD,astheymaynotnecessarilybetransferabletoolderadults.StrictBP-lowering,forexample,maycomewiththe
riskofdizziness,falls,andfracturesinolderadults,manyof
w

---

### Chunk 12/30
**Article:** Urea levels and cardiovascular disease in patients with chronic kidney disease (2022)
**Journal:** Nephrology Dialysis Transplantation
**Section:** other | **Similarity:** 0.485

20.3(5.1)0%Ageatbaseline(years)69[61–77]68[60–76]69[61–77]69[61–77]0.130%Men,n(%)666566670.710Smoking,n(%)0.030.8Never-smoker,n(%)40.644.539.637.7Currentsmoker,n(%)12.611.711.814.4Formersmoker,n(%)46.843.848.547.9eGFRatbaseline(mL/min/1.73m)33.5(11.6)43.5(9.9)32.6(8.9)24.5(7.0)<0.0010%Albumin-orprotein-to-creatinineratio<0.0018.0A1(normaltomildlyincreased),n(%)28.642.127.016.9A2(moderatelyincreased),n(%)31.831.833.729.7A3(severelyincreased),n(%)39.626.139.253.4Bodymassindex(kg/m)28.8(5.8)28.3(5.2)28.7(5.9)29.5(6.3)<0.0012.0%Diabetes,n(%)44.836.843.953.6<0.0010.2Systolicbloodpressure(mmHg)142(20)142(20)142(21)143(20)0.322.3%Historyofcardiovasculardisease,n(%)53.947.354.659.6<0.0011.3Anaemia,n(%)38.321.135.857.8<0.0010.3Serumbicarbonate(mmol/L)25.0(3.4)25.8(3.1)24.9(3.3)24.1(3.6)<0.0016.9%Serumalbumin(g/L)40.4(4.5)40.6(4.4)40.5(4.2)39.9(4.9)0.00915.2%High-sensitivityC-reactiveprotein(mg/L)2.5[1.1–5.9]2.2[1.1–5.0]2.5[1.1–5.4]2.9[1.2–7.1]<0.00117.6%Historyofacutekidneyinj

---

### Chunk 13/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.484

 10460–8945–5930–44<30105+90–104
60–8945–5930–44<30105+90–10460–8945–5930–44<30105+90–10460–89
45–59Urine albumin-creatinine ratio, mg/gUrine albumin-creatinine ratio, mg/gAll-cause mortality: 11 cohorts692 802 participants; 97 006 eventsMyocardial infarction: 10 cohorts649 365 participants; 17 926 eventsKidney failure with replacement therapy: 5 cohorts 630 370 participants; 4306 eventsHeart failure: 9 cohorts641 298 participants; 27 406 eventsCardiovascular mortality: 11 cohorts692 322 participants, 25 322 eventsStroke: 9 cohorts662 605 participants; 16 909 eventsAcute kidney injury: 5 cohorts630 370 participants; 24 062 events
607 102 participants; 37 278 eventsHospitalization: 3 cohorts630 489 participants; 464 894 eventsPeripheral artery disease: 6 cohorts642 624 participants; 3943 events30–44<30300+1.01.31.62.50.91.21.42.8ref1.31.52.0ref1.21.41.81.21.51.92.51.21.41.51.91.72.22.53.31.61.92.33.3
2.32.63.44.42.12.63.13.33.64.05.57.15.13.04.95.01.01.41.8

---

### Chunk 14/30
**Article:** Hipertensão Arterial Sistêmica (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.481

# Hipertensão Arterial Sistêmica

**Source:** https://web.plaud.ai/share/1dd61764908881487::YXdzOnVzLXdlc3QtMg

---

## Lecture

Data e Hora: 2025-11-20 20:43:21
Local: [Inserir Local]
Instrutor: Tullius Pervi (Túlio)
## 📝 Resumo
A aula, ministrada pelo Dr. Tullius Pervi (Túlio), cardiologista, oferece uma visão abrangente sobre a hipertensão arterial sistêmica, abordando-a como uma doença crônica, multifatorial e frequentemente assintomática, que constitui um grave problema de saúde pública. A apresentação detalha a definição, diagnóstico, prevalência e fatores de risco, como idade, obesidade e genética. É destacada a importância de um diagnóstico preciso, diferenciando a hipertensão primária da secundária e criticando a falta de rastreio para causas como apneia do sono. A fisiopatologia complexa é explicada, envolvendo múltiplos sistemas.

---

### Chunk 15/30
**Article:** Comparing Carotid Artery Velocities with Current ASCVD Risk Stratification: A Novel Approach to Simpler Risk Assessment (2024)
**Journal:** Journal of Epidemiology and Global Health
**Section:** abstract | **Similarity:** 0.481

This prospective study examined 1,636 participants aged 40–75 years without prior cardiovascular events to explore whether carotid artery blood flow measurements could simplify ASCVD risk assessment. Researchers used duplex ultrasonography to measure flow velocities and compared results against standard 2022 USPSTF risk guidelines. The investigation revealed that end diastolic velocity (EDV) of common carotid artery (CCA) and the peak systolic velocity (PSV) of internal carotid artery (ICA) were inversely and nonlinearly associated with cardiovascular event risk. Optimal cutoff velocities were identified: approximately 23.75 cm/s for CCA-EDV, 81.75 cm/s for ICA-PSV, and 26.75 cm/s for ICA-EDV. Analysis showed U-shaped relationships suggesting these measurements could complement existing risk assessment methods for primary cardiovascular prevention.

---

### Chunk 16/30
**Article:** Dislipidemias II (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.481

m DCV), condições gestacionais (pré-termo, hipertensivas, diabetes gestacional), autoimunidade, tratamento de câncer de mama e deficiências hormonais (climatério/menopausa), frequentemente subvalorizadas nos protocolos. O palestrante defende abordagem multidisciplinar e estruturada de estilo de vida, especialmente em hipertensão limítrofe, apoiada por nutricionistas e educação para adesão.
O uso de estatinas é discutido criticamente: reconhece-se benefício anti-inflamatório local no pós-angioplastia (lesão de parede e fragilidade do stent), porém questiona-se o uso indiscriminado, sobretudo em prevenção primária, citando meta-análise que desafia a hipótese lipídica e alertando para vieses na interpretação de risco relativo vs. absoluto. Em UTI, menciona-se aumento de delírio e a necessidade de evitar “receita de bolo” (anticoagulação, IBP, estatina automática).

---

### Chunk 17/30
**Article:** Dislipidemias II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.479

dade, inatividade física, pressão arterial e dislipidemia.
- A inflamação é indicada por marcadores como a Proteína C-Reativa (PCR), onde um valor de 5 já é considerado elevado, e a Interleucina 6 (exemplo de paciente com 8.45).
- A resposta anti-inflamatória é medida pela Interleucina 10, com um valor de corte de 3,5 (abaixo disso é um risco) e um exemplo de paciente com 6.44.
- A saúde do endotélio é avaliada pelo óxido nítrico, que deve estar em 8.8; um paciente com 4.8 já apresenta disfunção sistêmica.
- A LDL oxidada é um marcador crítico, com um limite saudável de 133, enquanto pacientes de alto risco podem apresentar valores extremos, como 1000.
**Achados Adicionais**
- Existem cinco parâmetros clássicos para definir a síndrome metabólica.
- O alvo de LDL para pacientes de alto risco, segundo as diretrizes atuais, é de 50.

---

### Chunk 18/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.479

saltos abruptos (ex: de 2.5 para 4 ou 5 em dois meses) exigem investigação, mesmo dentro da faixa de normalidade.
**O diagnóstico e manejo da Hiperplasia Prostática Benigna (HPB) dependem mais dos sintomas obstrutivos, como resíduo pós-miccional acima de 40 ml, do que do tamanho da próstata, que pode variar de 25 a mais de 80 gramas sem necessariamente causar problemas.**
- A HPB é comum a partir dos 45-50 anos, mas o tamanho da próstata (normalmente 25-30 gramas) não se correlaciona diretamente com a obstrução; próstatas de 28-29 gramas podem ser obstrutivas, enquanto outras de 70-80 gramas não.
- Um indicador chave de obstrução é o resíduo pós-miccional, com volumes acima de 40 ml sendo anormais, e a urofluxometria, onde um fluxo máximo (Qmax) abaixo de 10 ml/s é considerado alterado (normal > 15 ml/s).
- O tratamento medicamentoso para HPB inclui doses como 0,4 mg de Tansulosina ou 2 a 4 mg de Doxazosina para relaxar a musculatura e melhorar o fluxo urinário.

---

### Chunk 19/30
**Article:** The role of red blood cell distribution width (RDW) in cardiovascular risk assessment: useful or hype? (2019)
**Journal:** Annals of Translational Medicine
**Section:** abstract | **Similarity:** 0.474

Red cell distribution width (RDW) has emerged as a prognostic marker across multiple cardiovascular conditions. Reference range 12-15%. Each 1% increase in RDW associates with 1.10-fold higher all-cause mortality risk in heart failure. RDW >15% shows 3-fold increased mortality in CAD, 37% higher stroke risk, and 77% higher atrial fibrillation incidence.

---

### Chunk 20/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.472

actory to treatment≥4 antihypertensive agents• Persistent abnormalities of serum potassium• Acidosis• Anemia• Bone disease• Malnutrition
Figure48|Circumstancesforreferraltospecialistkidneycareservicesandgoalsofthereferral.ACR,albumin-to-creatinineratio;AER,albuminexcretionrate;CKD,chronickidneydisease;eGFR,estimatedglomerularltrationrate;KRT,kidneyreplacementtherapy;PCR,protein-creatinineratio;PER,proteinexcretionrate;RBC,redbloodcells.

---

### Chunk 21/30
**Article:** Confounders in Identification and Analysis of Inflammatory Biomarkers in Cardiovascular Diseases (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.472

atherosclerosis:Resultsfromaprospective,parallel-groupcohortstudy.Clin.Chim.Acta2015,447,16–22.[CrossRef]101.Hijazi,Z.;Lindahl,B.;Oldgren,J.;Andersson,U.;Lindbäck,J.;Granger,C.B.;Alexander,J.H.;Gersh,B.J.;Hanna,M.;Harjola,V.;etal.RepeatedMeasurementsofCardiacBiomarkersinAtrialFibrillationandValidationoftheABCStrokeScoreOverTime.J.Am.HeartAssoc.2017,6,e004851.[CrossRef]102.Abramson,J.L.;Lewis,C.;Murrah,N.V.;Anderson,G.T.;Vaccarino,V.RelationofC-ReactiveProteinandTumorNecrosisFactor-AlphatoAmbulatoryBloodPressureVariabilityinHealthyAdults.Am.J.Cardiol.2006,98,649–652.[CrossRef]103.Simundic,A.-M.;Kackov,S.;Miler,M.;Fraser,C.G.;Petersen,P.H.TermsandSymbolsUsedinStudiesonBiologicalVariation:TheNeedforHarmonization.Clin.Chem.2015,61,438–439.[CrossRef]
	
	
Getulio José Mattos Do Amaral Filho - getfilho@yahoo.com.br - CPF: 034.983.039-88

Biomolecules2021,11,1464
17of17
104.Khuseyinova,N.;Greven,S.;Rückerl,R.;Trischler,G.;Loewel,H.;Peters,A.;

---

### Chunk 22/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.471

–891.01.11.31.51.81.01.31.845–591.31.31.51.72.11.51.72.1Urine albumin-creatinine rg/gm ,oitar eninitaerc-nimubla enirUg/gm ,oitaAll-cause mortality: 82 cohorts26 444 384 participants; 2 604 028 events Myocardial infarction: 64 cohorts 22 838 356 participants; 451 063 events Kidney failure with replacement therapy: 57 cohorts 25 466 956 participants; 158 846 events Heart failure: 61 cohorts24 603 016 participants; 1 132 443 events Cardiovascular mortality: 76 cohorts26 022 346 participants; 776 441 events Stroke: 68 cohorts24 746 436 participants; 461 785 events Acute kidney injury: 49 cohorts23 914 614 participants; 1 408 929 events 
22 886 642 participants; 1 068 701 events Hospitalization: 49 cohorts25 426 722 participants; 8 398 637 events Peripheral artery disease: 54 cohorts24 830 794 participants; 378 924 events 30–441.51.51.61.92.32.01.92.515–291.81.81.92.42.83.33.33.8<152.72.83.03.23.89.19.09.6300–9991000+2.73.82.23.22.23.12.83.73.24.34.25.16.06.03.14.32.43.12.23.02.3

---

### Chunk 23/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.468

.Diagnosticaccuracyofpoint-of-caretestsfordetectingalbuminuria:a
systematicreviewandmeta-analysis.AnnIntMed.2014;160:550–557.318SoFtablesSupplementaryTableS5SearchdateJuly2022
Citationsscreened/includedstudies2184/65
SupplementaryFigureS5Chapter2RiskassessmentinpeoplewithCKDClinicalquestionArekidneyfailurepredictionequationsgoodpredictorsofprogression,kidneyfailure,orend-stagerenaldisease?PopulationAdults,children,andyoungpeoplewithCKDG1-G5PredictorKidneyfailureriskequations(e.g.,Tangriequation[KidneyFailureRiskEquation])OutcomesPrognosticperformance:Calibration(goodnessofmeasures,e.g.,R2,Brierscore,andHosmer-Lemeshowtest)Discrimination(e.g.,sensitivity/specicity;areaunderthecurve[AUC]fromreceiveroperatingcharacteristic[ROC]andareaunderthereceiveroperatingcharacteristiccurve[AUROC];C-statistic)StudydesignSystematicreview
ExistingsystematicreviewNationalInstituteforHealthandCareExcellence.Evidencereviewforthebestcombinationofmeasurestoidentifyincreasedriskofprogressioninadults,childr

---

### Chunk 24/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.468

treatmentforpreventionofmajorcardiovasculardiseasesinpeoplewithandwithouttype2diabetes:anindividualparticipant-leveldatameta-analysis.LancetDiabetesEndocrinol.2022;10:645–654.500.CheungAK,WheltonPK,MuntnerP,etal.InternationalConsensusonStandardizedClinicBloodPressureMeasurement—acalltoaction.AmJMed.2023;136:438–445.501.McManusRJ,MantJ,FranssenM,etal.Efcacyofself-monitoredbloodpressure,withorwithouttelemonitoring,fortitrationof
antihypertensivemedication(TASMINH4):anunmaskedrandomisedcontrolledtrial.Lancet.2018;391:949–959.502.JanseRJ,FuEL,ClaseCM,etal.Stoppingversuscontinuingrenin-angiotensin-systeminhibitorsafteracutekidneyinjuryandadverseclinicaloutcomes:anobservationalstudyfromroutinecaredata.ClinKidneyJ.2022;15:1109–1119.503.LeonSJ,WhitlockR,RigattoC,etal.Hyperkalemia-related
discontinuationofrenin-angiotensin-aldosteronesysteminhibitorsand
clinicaloutcomesinCKD:apopulation-basedcohortstudy.AmJKidneyDis.2022;80:164–173.504.SiewED,ParrSK,Abdel-KaderK,etal.Renin-angiote

---

### Chunk 25/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.467

Rcategory(ml/minper1.73m2)105D90–10475–8960–7445–5930–4415–290–14Serumpotassium$65Female4.1(0.5)4.2(1.3)4.2(0.5)4.3(0.5)4.3(1.3)4.4(0.5)4.5(1.0)4.5(2.0)Male4.2(0.5)4.3(0.6)4.3(1.1)4.4(0.6)4.4(0.7)4.5(1.1)4.6(0.6)4.6(1.6)<65Female4.1(0.7)4.2(1.3)4.3(17.0)4.2(1.0)4.3(0.5)4.3(0.6)4.4(0.6)4.5(1.1)Male4.2(0.4)4.3(0.5)4.3(0.6)4.3(0.4)4.4(0.5)4.5(0.6)4.5(0.7)4.6(0.7)eGFR,estimatedglomerularltrationrate;GFR,glomerularltrationrate.aDatafromtheOptumLabsDataWarehouse,alongitudinal,real-worlddataassetwithdeidentiedadministrativeclaimsandelectronichealthrecorddata.Thedatabasecontainslongitudinalhealthinformationonenrolleesandpatients,representingthediversityofgeographicalregionsacrosstheUnitedStates.

---

### Chunk 26/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** other | **Similarity:** 0.466

-enhancedultrasound(CEUS)andreal-timetissueelastography(RTE).J.Ultrasound2014,17,233–238.[CrossRef]30.Zebari,S.;Huang,D.Y.;Wilkins,C.J.;Sidhu,P.S.AcuteTesticularSegmentalInfarctFollowingEndovascularRepairofaJuxta-renalAbdominalAorticAneurysm:CaseReportandLiteratureReview.Urology2019,126,5–9.[CrossRef]

Cancers2024,16,2309
16of17
31.Kachramanoglou,C.;Rafailidis,V.;Philippidou,M.;Bertolotto,M.;Huang,D.Y.;Deganello,A.;Sellars,M.E.;Sidhu,P.S.MultiparametricSonographyofHematologicMalignanciesoftheTestis:Grayscale,ColorDoppler,andContrast-EnhancedUltrasoundandStrainElastographicAppearancesWithHistologicCorrelation.J.UltrasoundMed.2016,36,409–420.[CrossRef][PubMed]32.Huang,D.Y.;Sidhu,P.S.Focaltesticularlesions:ColourDopplerultrasound,contrast-enhancedultrasoundandtissueelastogra-phyasadjuvantstothediagnosis.Br.J.Radiol.2012,85,S41–S53.[CrossRef][PubMed]33.Huang,D.Y.;Pesapane,F.;Rafailidis,V.;Deganello,A.;Sellars,M.E.;Sidhu,P.S.Theroleofmultiparametricultrasoundinthediagnosisofpaediatr

---

### Chunk 27/30
**Article:** Hipertensão Arterial Sistêmica (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.466

egros).
    *   Excesso de peso, obesidade, sedentarismo, tabagismo.
    *   Ingestão de sódio e álcool.
    *   Fatores genéticos e socioeconômicos.
*   **Causas Primárias vs. Secundárias:**
    *   95% dos casos são classificados como hipertensão primária (essencial), atribuída à genética e estilo de vida.
    *   O palestrante argumenta que causas secundárias são subdiagnosticadas, como apneia obstrutiva do sono (presente em 40% dos hipertensos), disfunções tireoidianas e doença renal crônica. O tratamento da causa base pode curar a hipertensão.
### 2. Fisiopatologia e Diagnóstico
*   **Fisiopatologia Complexa:**
    *   A pressão arterial (PA = Débito Cardíaco x Resistência Vascular Periférica) é regulada por múltiplos sistemas: neural (catecolaminas), hormonal (sistema renina-angiotensina-aldosterona - SRAA), renal, microbiota e endotélio.

---

### Chunk 28/30
**Article:** hs-CRP/HDL-C can predict the risk of all-cause mortality in cardiovascular-kidney-metabolic syndrome stage 1-4 patients (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.466

,HuY,WangW,etal.Thepredictivevalueofthehs-CRP/HDL-Cratio,aninammation-lipidcompositemarker,forcardiovasculardiseaseinmiddle-agedandelderlypeople:evidencefromalargenationalcohortstudy.
LipidsHealthDis.(2024)23:66.doi:10.1186/s12944-024-02055-729.ZhaoY,HuY,SmithJP,StraussJ,YangG.Cohortprole:theChinahealthandretirementlongitudinalstudy(CHARLS).IntJEpidemiol.(2012)43:61–8.doi:10.1093/ije/dys20330.D’AgostinoRB,VasanRS,PencinaMJ,WolfPA,CobainM,MassaroJM,etal.Generalcardiovascularriskproleforuseinprimarycare.Circulation(2008)117:743–53.doi:10.1161/CIRCULATIONAHA.107.69957931.MaY-C,ZuoL,ChenJ-H,LuoQ,YuX-Q,LiY,etal.Modiedglomerularltrationrateestimatingequationforchinesepatientswithchronickidneydisease.JAmSocNephrol(2006)17:2937–44.doi:10.1681/asn.200604036832.JamesPA,OparilS,CarterBL,CushmanWC,Dennison-HimmelfarbC,HandlerJ,etal.Evidence-Basedguidelineforthemanagementofhighbloodpressureinadults:
reportfromthepanelmembersappointedtotheeighthjointnationalcommittee(JNC8).

---

### Chunk 29/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.465

k forrecurrence afterkidney transplantationConditions for whichgenetic testing isrelevant for reproductivecounselingExamples:• GLA (Fabry)• AGXT (primaryhyperoxaluria (PH))• CoQ10 genes (SRNS)• CTNS (cystinosis)• Tubulopathies(Na+, K+ etc.)Conditions amenable to nonspecific renoprotective strategiesExample:• COL4A3/4/5 (Alport)and RAAS blockadeExample:
• Glomerular diseasedue to mutations inAlport genes(COL4A3/4/5)Examples:• (CFH/CFI/C3...): aHUS• (AGXT, GRHPR, HOGA): primary hyperoxaluria (PH)
• Adenine phosphoribo-  syltransferase deficiency (APRT)Conditions amenable to specific screening for extrarenal manifestationsExamples:• HNF1B: diabetes• PKD1/PKD2(ADPKD): intracranialaneurysms • FLCN: renal cellcarcinoma, etc.Example:
• Prenatal/preimplantationdiagnosis
Figure9|Actionablegenesinkidneydisease.Actionabilityreferstothepotentialforgeneticrestresultstoleadtospecicclinicalactionsfrompreventionortreatmentofacondition,supportedbyrecommendationsbasedonevide

---

### Chunk 30/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.460

netheprognosticvalueofcirculatingcardiacbiomarkersinadvancedchronickidneydisease.CanJCardiol.2019;35:1106–1113.663.ColletJP,ThieleH,BarbatoE,etal.2020ESCGuidelinesforthe
managementofacutecoronarysyndromesinpatientspresentingwithoutpersistentST-segmentelevation.EurHeartJ.2021;42:1289–1367.664.KonoK,FujiiH,NakaiK,etal.Compositionandplaquepatternsof
coronaryculpritlesionsandclinicalcharacteristicsofpatientswith
chronickidneydisease.KidneyInt.2012;82:344–351.665.SchiffrinEL,LipmanML,MannJF.Chronickidneydisease:effectsonthecardiovascularsystem.Circulation.2007;116:85–97.666.NationalInstituteforHealthandCareExcellence.Familialhypercholesterolaemia:identicationandmanagement.ClinicalGuideline[CG71].NICE.2008.667.CollinsR,ReithC,EmbersonJ,etal.Interpretationoftheevidencefortheefcacyandsafetyofstatintherapy.Lancet.2016;388:2532–2561.668.CholesterolTreatmentTrialistsCollaboration.Effectofstatintherapyon
musclesymptoms:anindividualparticipantdatameta-analysisoflarge-
scale,randomised,d

---

