# ScoreItem: Doppler Aorta - RAR (Renal-Aortic Ratio)

**ID:** `c77cedd3-2800-7a21-b23d-8edf321f7874`
**FullName:** Doppler Aorta - RAR (Renal-Aortic Ratio) (Exames - Imagem)
**Unit:** ratio

**Preparation Metadata:**
- Quality Grade: **GOOD**
- Total Chunks: 30 de 16 artigos
- Avg Similarity: 0.521

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `c77cedd3-2800-7a21-b23d-8edf321f7874`.**

```json
{
  "score_item_id": "c77cedd3-2800-7a21-b23d-8edf321f7874",
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

**ScoreItem:** Doppler Aorta - RAR (Renal-Aortic Ratio) (Exames - Imagem)
**Unidade:** ratio

**30 chunks de 16 artigos (avg similarity: 0.521)**

### Chunk 1/30
**Article:** Critical analysis of renal duplex ultrasound parameters in detecting significant renal artery stenosis (2012)
**Journal:** Journal of Vascular Surgery
**Section:** abstract | **Similarity:** 0.779

Large study of 313 patients evaluating Doppler parameters for renal artery stenosis. Mean renal-aortic ratios for normal, <60%, and ≥60% stenosis were 2.2, 2.9, and 4.5 respectively. RAR >3.5 demonstrated high diagnostic accuracy for detecting hemodynamically significant stenosis.

---

### Chunk 2/30
**Article:** Doppler Renal Assessment, Protocols, and Interpretation (2024)
**Journal:** StatPearls
**Section:** abstract | **Similarity:** 0.777

Comprehensive review of renal Doppler ultrasound techniques, including renal-aortic ratio (RAR) for detecting renal artery stenosis. The RAR compares intrastenotic flow velocity in renal arteries with aortic reference values, with RAR >3.5 predicting ≥60% stenosis with 84-91% sensitivity and 95-97% specificity.

---

### Chunk 3/30
**Article:** Doppler ultrasound and renal artery stenosis: An overview (2013)
**Journal:** Journal of Ultrasound
**Section:** abstract | **Similarity:** 0.769

Overview of Doppler ultrasound techniques for diagnosing renal artery stenosis, the most common cause of secondary hypertension. Discusses renal-aortic ratio as a reliable parameter that normalizes individual hemodynamic variations, improving diagnostic specificity compared to peak systolic velocity alone.

---

### Chunk 4/30
**Article:** Optimal Peak Systolic Velocity Thresholds for Predicting Internal Carotid Artery Stenosis Greater than or Equal to 50%, 60%, 70%, and 80% (2016)
**Journal:** Journal of Stroke and Cerebrovascular Diseases
**Section:** abstract | **Similarity:** 0.550

The research established optimal peak systolic velocity measurements for detecting various degrees of internal carotid artery stenosis. Testing 127 arterial specimens, researchers identified specific thresholds: 130 cm/s, 160 cm/s, 200 cm/s, and 270 cm/s for detecting stenosis at increasing severity levels (≥50%, ≥60%, ≥70%, and ≥80% respectively). The findings demonstrated high diagnostic accuracies across all measured thresholds, with sensitivity and specificity values exceeding 85% for each threshold category.

---

### Chunk 5/30
**Article:** Correlation between Ultrasound Peak Systolic Velocity and Angiography for Grading Internal Carotid Artery Stenosis (2024)
**Journal:** Journal of Clinical Medicine
**Section:** abstract | **Similarity:** 0.535

The study evaluated how peak systolic velocity (PSV) measured via duplex ultrasound correlates with angiography findings for assessing internal carotid artery stenosis. Researchers analyzed 47 stenotic lesions using both ultrasound and digital subtraction angiography, applying NASCET and ECST classification protocols. Key findings indicated that a PSV threshold of 200 cm/s was found to be the best criterion for identifying severe NASCET stenoses (≥70%), while a threshold of 180 cm/s was the best for ECST stenoses (≥80%). However, PSV demonstrated limited reliability for moderate stenoses, suggesting complementary imaging techniques should be employed in such cases.

---

### Chunk 6/30
**Article:** Hipertensão Arterial Sistêmica (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.531

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

### Chunk 7/30
**Article:** Peak systolic velocity ratio for evaluation of internal carotid artery stenosis correlated with plaque morphology: substudy results of the ANTIQUE study (2023)
**Journal:** Frontiers in Neurology
**Section:** abstract | **Similarity:** 0.525

This study evaluated how effectively four duplex sonography measurements assess carotid artery narrowing severity when compared against computed tomography angiography. The research examined 143 patients with significant stenosis, analyzing peak systolic velocity (PSV), velocity ratios, end-diastolic velocity, and B-mode imaging. Results demonstrated that the PSV ICA/CCA ratio showed the highest correlation with CTA, followed by PSV and other parameters. The study found that plaque composition significantly influenced measurement accuracy, with calcified plaques producing substantially weaker correlations than softer plaque types with smooth surfaces.

---

### Chunk 8/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** other | **Similarity:** 0.520

-enhancedultrasound(CEUS)andreal-timetissueelastography(RTE).J.Ultrasound2014,17,233–238.[CrossRef]30.Zebari,S.;Huang,D.Y.;Wilkins,C.J.;Sidhu,P.S.AcuteTesticularSegmentalInfarctFollowingEndovascularRepairofaJuxta-renalAbdominalAorticAneurysm:CaseReportandLiteratureReview.Urology2019,126,5–9.[CrossRef]

Cancers2024,16,2309
16of17
31.Kachramanoglou,C.;Rafailidis,V.;Philippidou,M.;Bertolotto,M.;Huang,D.Y.;Deganello,A.;Sellars,M.E.;Sidhu,P.S.MultiparametricSonographyofHematologicMalignanciesoftheTestis:Grayscale,ColorDoppler,andContrast-EnhancedUltrasoundandStrainElastographicAppearancesWithHistologicCorrelation.J.UltrasoundMed.2016,36,409–420.[CrossRef][PubMed]32.Huang,D.Y.;Sidhu,P.S.Focaltesticularlesions:ColourDopplerultrasound,contrast-enhancedultrasoundandtissueelastogra-phyasadjuvantstothediagnosis.Br.J.Radiol.2012,85,S41–S53.[CrossRef][PubMed]33.Huang,D.Y.;Pesapane,F.;Rafailidis,V.;Deganello,A.;Sellars,M.E.;Sidhu,P.S.Theroleofmultiparametricultrasoundinthediagnosisofpaediatr

---

### Chunk 9/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.504

ardiovasculardiseaseandantithrombotictherapy.EurHeartJ.2013;34:1708–1713,1713a–1713b.697.JamesS,BudajA,AylwardP,etal.Ticagrelorversusclopidogrelinacutecoronarysyndromesinrelationtorenalfunction:resultsfromthePlateletInhibitionandPatientOutcomes(PLATO)trial.Circulation.2010;122:1056–1067.698.HerringtonWG,StaplinN.InpatientswithcoronarydiseaseandCKD,
addinganinvasivestrategytoMTdidnotimproveoutcomes.AnnInternMed.2020;173:JC16.699.SarnakMJ,AmannK,BangaloreS,etal.Chronickidneydiseaseand
coronaryarterydisease:JACCstate-of-the-artreview.JAmCollCardiol.2019;74:1823–1838.700.CharytanDM,WallentinL,LagerqvistB,etal.Earlyangiographyin
patientswithchronickidneydisease:acollaborativesystematicreview.ClinJAmSocNephrol.2009;4:1032–1043.701.HastingsRS,HochmanJS,DzavikV,etal.Effectoflaterevascularization
ofatotallyoccludedcoronaryarteryaftermyocardialinfarctiononmortalityratesinpatientswithrenalimpairment.AmJCardiol.2012;110:954–960.702.JohnstonN,JernbergT,LagerqvistB,etal.Earlyinvasivetrea

---

### Chunk 10/30
**Article:** Hipertensão Arterial Sistêmica (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.499

causa secundária identificável.
**O diagnóstico e a classificação da hipertensão seguem limiares específicos que variam conforme o método de medição, com estágios progressivos que orientam a terapia.**
- A pressão arterial é classificada como ótima (abaixo de 120/80 mmHg), normal (até 129/84 mmHg) e pré-hipertensão (130-139 / 85-89 mmHg).
- O diagnóstico de hipertensão é estabelecido a partir de 14 por 9 mmHg em medições de consultório, aplicável a indivíduos com mais de 18 anos.
- Os estágios da hipertensão são definidos como: Estágio 1 (a partir de 14/9), Estágio 2 (a partir de 16/10) e Estágio 3 (acima de 18/11).
- Para exames fora do consultório, os limiares são mais baixos: 13 por 8 mmHg para o MAPA (24 horas) e 13 por 8,5 mmHg para o MRPA.

---

### Chunk 11/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.496

urinário (Qmax). Valores < 10 ml/segundo indicam obstrução.
    *   **Ressonância Magnética Multiparamétrica 3-Tesla (3T)**: Exame de alta especificidade. O palestrante solicita de rotina para homens > 50 anos, ou > 40 anos com histórico familiar ou alterações súbitas no PSA.
*   **Dosagem Hormonal Salivar e Quociente Estrogênico**:
    *   **Vantagens da Saliva**: Via não invasiva que mede a fração livre e 100% bioativa dos hormônios (Testosterona, DHT, Estradiol, etc.). Útil quando a clínica do paciente não corresponde aos exames de sangue.
    *   **Quociente Estrogênico**: Fórmula para avaliar o risco de doenças prostáticas.
        *   **Fórmula**: Estriol / (Estradiol + Estrona).
        *   **Valores > 1**: Bom prognóstico (perfil estrogênico protetor).
        *   **Valores < 1**: Mau prognóstico (prevalência de estrogênios proliferativos), sugerindo a necessidade de trabalhar a metilação da estrona.

---

### Chunk 12/30
**Article:** Hipertensão Arterial Sistêmica (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.496

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

### Chunk 13/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.493

nsiderado alterado (normal > 15 ml/s).
- O tratamento medicamentoso para HPB inclui doses como 0,4 mg de Tansulosina ou 2 a 4 mg de Doxazosina para relaxar a musculatura e melhorar o fluxo urinário.
### Achados Adicionais Chave
- A ressonância magnética multiparamétrica 3 Tesla é uma ferramenta de alta definição, recomendada a partir dos 40 anos para homens com histórico familiar ou a partir dos 50 anos como rotina para uma avaliação prostática precisa.
- A dosagem hormonal salivar oferece uma medição precisa dos hormônios livres, com faixas de referência para testosterona (47-150), estradiol (0.6-3) e o quociente estrogênico (0.04-1.67), que avalia o equilíbrio hormonal.
- A escala Gleason, que vai de 1 a 10, mede a agressividade do câncer de próstata, com tumores classificados como 8, 9 ou 10 sendo considerados os mais agressivos.
- A hiperplasia prostática é uma condição que afeta apenas 3 espécies: humanos, cães e macacos.

---

### Chunk 14/30
**Article:** Hipertensão Arterial Sistêmica (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.493

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
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.490

.Diagnosticaccuracyofpoint-of-caretestsfordetectingalbuminuria:a
systematicreviewandmeta-analysis.AnnIntMed.2014;160:550–557.318SoFtablesSupplementaryTableS5SearchdateJuly2022
Citationsscreened/includedstudies2184/65
SupplementaryFigureS5Chapter2RiskassessmentinpeoplewithCKDClinicalquestionArekidneyfailurepredictionequationsgoodpredictorsofprogression,kidneyfailure,orend-stagerenaldisease?PopulationAdults,children,andyoungpeoplewithCKDG1-G5PredictorKidneyfailureriskequations(e.g.,Tangriequation[KidneyFailureRiskEquation])OutcomesPrognosticperformance:Calibration(goodnessofmeasures,e.g.,R2,Brierscore,andHosmer-Lemeshowtest)Discrimination(e.g.,sensitivity/specicity;areaunderthecurve[AUC]fromreceiveroperatingcharacteristic[ROC]andareaunderthereceiveroperatingcharacteristiccurve[AUROC];C-statistic)StudydesignSystematicreview
ExistingsystematicreviewNationalInstituteforHealthandCareExcellence.Evidencereviewforthebestcombinationofmeasurestoidentifyincreasedriskofprogressioninadults,childr

---

### Chunk 16/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.487

aticreviews,asavailable.DetailsofthePICOSforthesequestionsarealsoprovidedinTable44.Literaturesearchesandarticleselection.SearchesforRCTswereconductedonPubMed,Embase,andtheCochrane
CentralRegisterofControlledTrials(CENTRAL),and
searchesfordiagnosis/prognosisstudieswereconductedon
PubMed,Embase,andCINAHL.Fortopicswithavailable
existingreviews,thereviewwasusedandanupdatedsearchwasconducted.ThesearchstrategiesareprovidedinAppendixA:SupplementaryTableS1.Toimproveefciencyandaccuracyinthetitle/abstractscreeningprocessandtomanagetheprocess,searchresults
wereuploadedtoaweb-basedscreeningtool,PICOPortal
(www.picoportal.net).PICOPortalusesmachinelearningtosortandpresentthosecitationsmostlikelytobepromotedto
full-textscreeningrst.Thetitlesandabstractsresultingfrom
methodsforguidelinedevelopmentwww.kidney-international.orgS274KidneyInternational(2024)105(Suppl4S),S117–S314

Table44|ClinicalquestionsandsystematicreviewtopicsinPICOSformatChapter1Evaluationofchronickidneydisease(CKD)Clinicalquesti

---

### Chunk 17/30
**Article:** Dislipidemias II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.485

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
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.483

e-basedequations.TransplInt.2012;25:527–536.162.JanusN,Launay-VacherV,ByloosE,etal.CancerandrenalinsufciencyresultsoftheBIRMAstudy.BrJCancer.2010;103:1815–1821.163.Launay-VacherV,JanusN,DerayG.Renalinsufciencyandcancertreatments.ESMOOpen.2016;1:e000091.164.NaSY,SungJY,ChangJH,etal.Chronickidneydiseaseincancerpatients:anindependentpredictorofcancer-specicmortality.AmJNephrol.2011;33:121–130.165.RosnerMH,JhaveriKD,McMahonBA,etal.Onconephrology:the
intersectionsbetweenthekidneyandcancer.CACancerJClin.2021;71:47–77.166.SoveriI,BergUB,BjorkJ,etal.MeasuringGFR:asystematicreview.AmJKidneyDis.2014;64:411–424.167.WhiteCA,AkbariA,AllenC,etal.Simultaneousglomerularltrationratedeterminationusinginulin,iohexol,and99mTc-DTPAdemonstratestheneedforcustomizedmeasurementprotocols.KidneyInt.2021;99:957–966.168.XieP,HuangJM,LiuXM,etal.(99m)Tc-DTPArenaldynamicimaging
methodmaybeunsuitabletobeusedasthereferencemethodin
investigatingthevalidityofCDK-EPIequationfordetermining
glomerularltrati

---

### Chunk 19/30
**Article:** Hipertensão Arterial Sistêmica II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.481

o de Diagnóstico:
1.  **Histórico Médico:** O conteúdo é uma palestra médica sobre fatores de risco cardiovascular contemporâneos, não uma consulta com um paciente específico.
2.  **Histórico de Medicação:** Inserir mais aqui
## Subjetivo:
O conteúdo é uma palestra médica e não contém queixas subjetivas de um paciente.
## Objetivo:
O conteúdo é uma palestra médica e não contém achados de exames de um paciente específico. O palestrante menciona seus próprios resultados de exames como exemplo:
-   **Índice de Ômega-3:** 6.7 (ideal entre 3 e 14).
-   **Relação Ômega-6 para Ômega-3:** 5:1 (ideal de 2:1 a 3:1), apesar da suplementação.
## Diagnóstico Primário:
-   **Avaliação:** O texto é uma apresentação educacional sobre fatores de risco inflamatórios e metabólicos para doenças vasculares, que são frequentemente negligenciados na cardiologia tradicional. Os principais fatores discutidos incluem:
    -   Desequilíbrio entre Ômega-3 e Ômega-6.

---

### Chunk 20/30
**Article:** Hipertensão Arterial Sistêmica (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.480

*Classes Preferenciais:** IECA (Inibidores da Enzima Conversora de Angiotensina), BRA (Bloqueadores do Receptor de Angiotensina), diuréticos tiazídicos e bloqueadores do canal de cálcio. A combinação entre eles é a melhor estratégia. A associação de IECA com BRA é proibida.
*   **Hierarquia Terapêutica:**
    1.  Mudança de estilo de vida.
    2.  IECA/BRA, bloqueador de canal de cálcio ou diurético tiazídico.
    3.  Espironolactona (4ª opção).
    4.  Betabloqueador (5ª opção).
*   **Betabloqueadores:** Não são mais primeira linha. Têm menor proteção cardiovascular, aumentam o risco de diabetes e causam efeitos adversos (disfunção sexual, ganho de peso). São considerados remédios de exceção.
*   **Metas Terapêuticas:**
    *   **Alto risco:** Manter PA entre 120/70 e 130/80 mmHg.
    *   **Baixo/moderado risco e idosos hígidos:** Manter PA até 140/90 mmHg.

---

### Chunk 21/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.477

treatmentforpreventionofmajorcardiovasculardiseasesinpeoplewithandwithouttype2diabetes:anindividualparticipant-leveldatameta-analysis.LancetDiabetesEndocrinol.2022;10:645–654.500.CheungAK,WheltonPK,MuntnerP,etal.InternationalConsensusonStandardizedClinicBloodPressureMeasurement—acalltoaction.AmJMed.2023;136:438–445.501.McManusRJ,MantJ,FranssenM,etal.Efcacyofself-monitoredbloodpressure,withorwithouttelemonitoring,fortitrationof
antihypertensivemedication(TASMINH4):anunmaskedrandomisedcontrolledtrial.Lancet.2018;391:949–959.502.JanseRJ,FuEL,ClaseCM,etal.Stoppingversuscontinuingrenin-angiotensin-systeminhibitorsafteracutekidneyinjuryandadverseclinicaloutcomes:anobservationalstudyfromroutinecaredata.ClinKidneyJ.2022;15:1109–1119.503.LeonSJ,WhitlockR,RigattoC,etal.Hyperkalemia-related
discontinuationofrenin-angiotensin-aldosteronesysteminhibitorsand
clinicaloutcomesinCKD:apopulation-basedcohortstudy.AmJKidneyDis.2022;80:164–173.504.SiewED,ParrSK,Abdel-KaderK,etal.Renin-angiote

---

### Chunk 22/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.477

saltos abruptos (ex: de 2.5 para 4 ou 5 em dois meses) exigem investigação, mesmo dentro da faixa de normalidade.
**O diagnóstico e manejo da Hiperplasia Prostática Benigna (HPB) dependem mais dos sintomas obstrutivos, como resíduo pós-miccional acima de 40 ml, do que do tamanho da próstata, que pode variar de 25 a mais de 80 gramas sem necessariamente causar problemas.**
- A HPB é comum a partir dos 45-50 anos, mas o tamanho da próstata (normalmente 25-30 gramas) não se correlaciona diretamente com a obstrução; próstatas de 28-29 gramas podem ser obstrutivas, enquanto outras de 70-80 gramas não.
- Um indicador chave de obstrução é o resíduo pós-miccional, com volumes acima de 40 ml sendo anormais, e a urofluxometria, onde um fluxo máximo (Qmax) abaixo de 10 ml/s é considerado alterado (normal > 15 ml/s).
- O tratamento medicamentoso para HPB inclui doses como 0,4 mg de Tansulosina ou 2 a 4 mg de Doxazosina para relaxar a musculatura e melhorar o fluxo urinário.

---

### Chunk 23/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.477

 10460–8945–5930–44<30105+90–104
60–8945–5930–44<30105+90–10460–8945–5930–44<30105+90–10460–89
45–59Urine albumin-creatinine ratio, mg/gUrine albumin-creatinine ratio, mg/gAll-cause mortality: 11 cohorts692 802 participants; 97 006 eventsMyocardial infarction: 10 cohorts649 365 participants; 17 926 eventsKidney failure with replacement therapy: 5 cohorts 630 370 participants; 4306 eventsHeart failure: 9 cohorts641 298 participants; 27 406 eventsCardiovascular mortality: 11 cohorts692 322 participants, 25 322 eventsStroke: 9 cohorts662 605 participants; 16 909 eventsAcute kidney injury: 5 cohorts630 370 participants; 24 062 events
607 102 participants; 37 278 eventsHospitalization: 3 cohorts630 489 participants; 464 894 eventsPeripheral artery disease: 6 cohorts642 624 participants; 3943 events30–44<30300+1.01.31.62.50.91.21.42.8ref1.31.52.0ref1.21.41.81.21.51.92.51.21.41.51.91.72.22.53.31.61.92.33.3
2.32.63.44.42.12.63.13.33.64.05.57.15.13.04.95.01.01.41.8

---

### Chunk 24/30
**Article:** Dislipdemia (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.476

quais ~250 mil ocorrem antes dos 60 anos.
- Brasil: 410 mil mortes/ano por DCV; 14 milhões com alguma DCV; 36% dos óbitos ≥55 anos decorrem de doença cardio-circulatória.
- Fisiopatologia: fluxo arterial torna-se turbilhonado ~50% de estenose, aumentando estresse de parede e risco de ruptura de placa.
**Achados Adicionais**
- LDL alvo em baixo risco: diretriz sugere <130 mg/dL (por vezes <100), mas o número isolado não determina benefício sem DCV prévia e sem avaliação de partículas/cálcio.
- LDL basal em estudo: 190 mg/dL; colesterol total 275 e HDL 31, ilustrando perfis iniciais elevados; em pacientes com cálcio zero, LDL >240 não alterou mortalidade/infarto com estatina.

---

### Chunk 25/30
**Article:** Urea levels and cardiovascular disease in patients with chronic kidney disease (2022)
**Journal:** Nephrology Dialysis Transplantation
**Section:** other | **Similarity:** 0.475

20.3(5.1)0%Ageatbaseline(years)69[61–77]68[60–76]69[61–77]69[61–77]0.130%Men,n(%)666566670.710Smoking,n(%)0.030.8Never-smoker,n(%)40.644.539.637.7Currentsmoker,n(%)12.611.711.814.4Formersmoker,n(%)46.843.848.547.9eGFRatbaseline(mL/min/1.73m)33.5(11.6)43.5(9.9)32.6(8.9)24.5(7.0)<0.0010%Albumin-orprotein-to-creatinineratio<0.0018.0A1(normaltomildlyincreased),n(%)28.642.127.016.9A2(moderatelyincreased),n(%)31.831.833.729.7A3(severelyincreased),n(%)39.626.139.253.4Bodymassindex(kg/m)28.8(5.8)28.3(5.2)28.7(5.9)29.5(6.3)<0.0012.0%Diabetes,n(%)44.836.843.953.6<0.0010.2Systolicbloodpressure(mmHg)142(20)142(20)142(21)143(20)0.322.3%Historyofcardiovasculardisease,n(%)53.947.354.659.6<0.0011.3Anaemia,n(%)38.321.135.857.8<0.0010.3Serumbicarbonate(mmol/L)25.0(3.4)25.8(3.1)24.9(3.3)24.1(3.6)<0.0016.9%Serumalbumin(g/L)40.4(4.5)40.6(4.4)40.5(4.2)39.9(4.9)0.00915.2%High-sensitivityC-reactiveprotein(mg/L)2.5[1.1–5.9]2.2[1.1–5.0]2.5[1.1–5.4]2.9[1.2–7.1]<0.00117.6%Historyofacutekidneyinj

---

### Chunk 26/30
**Article:** Hipertensão Arterial Sistêmica (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.474

a PA sistólica em 5.6 mmHg e a diastólica em 2.8 mmHg. Potencializa o efeito de anti-hipertensivos. A forma taurato é a mais indicada.
*   **Cúrcuma Longa:** Uso por mais de 12 semanas mostrou redução média de 8 mmHg na PA sistólica.
*   **Outros:** Potássio (com cautela), quercetina, arginina, cacau, resveratrol e piquenogenol também podem auxiliar.
*   **Abordagem Integrativa:** A suplementação melhora vias metabólicas e inflamatórias, auxiliando no controle da pressão. É crucial saber quando usar medicação se as metas não forem atingidas apenas com estilo de vida e suplementos.
## ❓ Perguntas
- [Inserir Pergunta/Confusão]
## 📚 Tarefas
- [ ] 1. Para pacientes com suspeita de hipertensão, solicitar MAPA ou MRPA para um diagnóstico preciso.
- [ ] 2. Rastrear ativamente causas de hipertensão secundária, como apneia do sono (polissonografia) e disfunções da tireoide (TSH).
- [ ] 3.

---

### Chunk 27/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.473

S(i)renin-angiotensin-aldosteronesystem(inhibitor)RBCredbloodcellRCTrandomizedcontrolledtrialRRrelativeriskSCrserumcreatinineSBPsystolicbloodpressureSESsocioeconomicstatusSGLT2isodium-glucosecotransporter-2inhibitor(s)T1DType1diabetesT2DType2diabetesUKUnitedKingdomUSUnitedStatesUSRDSUnitedStatesRenalDataSystem
WHOWorldHealthOrganization
abbreviationsandacronymswww.kidney-international.orgS128KidneyInternational(2024)105(Suppl4S),S117–S314

NoticeSECTIONI:USEOFTHECLINICALPRACTICEGUIDELINEThisClinicalPracticeGuidelinedocumentisbaseduponliteraturesearchesconductedfromJuly2022throughApril2023andupdatedinJuly2023.Itisdesignedtoassistdecision-making.Itisnotintendedtodeneastandardofcareandshouldnotbeinterpretedasprescribinganexclusivecourseofmanagement.Variationsinpracticewillinevitablyandappropriatelyoccur
whencliniciansconsidertheneedsofindividualpatients,availableresources,andlimitationsuniquetoaninstitutionortypeof
practice.Healthcareprovidersusingthestatementsinthisdocument(bothpracti

---

### Chunk 28/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.473

e resistência insulínica e sua conexão com a síndrome metabólica e as doenças cardiovasculares.
- [ ] 2. Comparar as diretrizes da dieta DASH com as de uma dieta focada na correção da resistência insulínica (ex: baixo carboidrato) para avaliar qual abordagem é mais adequada pessoalmente.
- [ ] 3. Investigar a aplicação do jejum intermitente (TRE) como estratégia complementar no manejo da hipertensão, considerando seus efeitos na resistência insulínica.
- [ ] 4. Estudar os mecanismos fisiopatológicos do processo aterosclerótico para além da hipótese lipídica, focando em inflamação, estresse oxidativo e saúde endotelial.
- [ ] 5. Ao avaliar o risco cardiovascular, utilizar marcadores mais abrangentes do que apenas o colesterol LDL, como a relação ApoB/ApoA e fatores de risco psicossociais.
- [ ] 6.

---

### Chunk 29/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.471

Específico (PSA)**:
    *   **Função**: Enzima que liquefaz o sêmen.
    *   **Formas**: Complexado (maior parte) e Livre. O PSA total é a soma de ambas.
    *   **Interpretação Clínica**: A relação PSA livre / PSA total é crucial.
        *   **> 0.14 (ou 14%)**: Sugere HPB.
        *   **< 0.14 (ou 14%)**: Aumenta o risco de câncer de próstata.
    *   **Limitações**: Cerca de 1-4% dos cânceres de próstata ocorrem com PSA normal. Em homens com baixa testosterona, esse número pode chegar a 15%.
*   **Exames de Imagem**:
    *   **Ultrassonografia de Próstata com Estudo do Resíduo Pós-Miccional**: Avalia anatomia e função. Um resíduo pós-miccional > 40 ml indica obstrução.
    *   **Urofluxometria**: Indicada para sintomas obstrutivos. Mede a velocidade do fluxo urinário (Qmax). Valores < 10 ml/segundo indicam obstrução.
    *   **Ressonância Magnética Multiparamétrica 3-Tesla (3T)**: Exame de alta especificidade.

---

### Chunk 30/30
**Article:** Resistência Insulínica (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.470

g/dL, HbA1c <5,7%.
- Pré-diabetes: jejum ≥100 e <126 mg/dL; 2h OGTT ≥140 e <200 mg/dL; HbA1c ≥5,7 e <6,5%; qualquer positividade confirma.
- Diabetes: jejum ≥126 mg/dL; 2h OGTT ≥200 mg/dL; glicemia aleatória ≥200 mg/dL com sintomas típicos; HbA1c ≥6,5%.
- Repetir exames na ausência de correlação clínica/sintomas antes de confirmar diagnóstico.
## Síndrome Metabólica: Definição e Critérios
- Evolução da RI para síndrome metabólica: hipertensão, DM2, risco cardiovascular (AVC/infarto).
- Definição prática: insuficiência do tecido adiposo para lidar com supernutrição.
- Critérios (ATP III/IDF): circunferência abdominal elevada (cortes variáveis por etnia), TG >150 mg/dL, HDL baixo, PA elevada, glicemia alterada; tratamento medicamentoso conta ponto.
- Condições associadas: SOP, lipodistrofias, história familiar, obesidade visceral.

---

