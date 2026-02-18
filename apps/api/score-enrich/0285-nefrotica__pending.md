# ScoreItem: Nefrótica

**ID:** `019bf31d-2ef0-7bdb-b7c8-1c239c373389`
**FullName:** Nefrótica (Histórico de doenças - Doenças crônicas (Questionar ativamente as doenças crônicas mais comuns, perguntar sobre duração, grau de controle e tratamentos utilizados no passado e atualmente) - Outras doenças renais)

**Preparation Metadata:**
- Quality Grade: **GOOD**
- Total Chunks: 30 de 10 artigos
- Avg Similarity: 0.518

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-7bdb-b7c8-1c239c373389`.**

```json
{
  "score_item_id": "019bf31d-2ef0-7bdb-b7c8-1c239c373389",
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

**ScoreItem:** Nefrótica (Histórico de doenças - Doenças crônicas (Questionar ativamente as doenças crônicas mais comuns, perguntar sobre duração, grau de controle e tratamentos utilizados no passado e atualmente) - Outras doenças renais)

**30 chunks de 10 artigos (avg similarity: 0.518)**

### Chunk 1/30
**Article:** Dislipidemias II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.594

dade, inatividade física, pressão arterial e dislipidemia.
- A inflamação é indicada por marcadores como a Proteína C-Reativa (PCR), onde um valor de 5 já é considerado elevado, e a Interleucina 6 (exemplo de paciente com 8.45).
- A resposta anti-inflamatória é medida pela Interleucina 10, com um valor de corte de 3,5 (abaixo disso é um risco) e um exemplo de paciente com 6.44.
- A saúde do endotélio é avaliada pelo óxido nítrico, que deve estar em 8.8; um paciente com 4.8 já apresenta disfunção sistêmica.
- A LDL oxidada é um marcador crítico, com um limite saudável de 133, enquanto pacientes de alto risco podem apresentar valores extremos, como 1000.
**Achados Adicionais**
- Existem cinco parâmetros clássicos para definir a síndrome metabólica.
- O alvo de LDL para pacientes de alto risco, segundo as diretrizes atuais, é de 50.

---

### Chunk 2/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.563

arltrationrate[eGFRcr-cys])(1B).1.1.3EvaluationofchronicityPracticePoint1.1.3.1:Proofofchronicity(durationofaminimumof3months)canbeestablishedby:(i)reviewofpastmeasurements/estimationsofGFR;(ii)reviewofpastmeasurementsofalbuminuriaorproteinuriaandurinemicroscopicexaminations;(iii)imagingndingssuchasreducedkidneysizeandreductionincorticalthickness;(iv)kidneypathologicalndingssuchasbrosisandatrophy;(v)medicalhistory,especiallyconditionsknowntocauseorcontributetoCKD;(vi)repeatmeasurementswithinandbeyondthe3-monthpoint.PracticePoint1.1.3.2:DonotassumechronicitybaseduponasingleabnormallevelforeGFRandACR,asthendingcouldbetheresultofarecentacutekidneyinjury(AKI)eventoracutekidneydisease(AKD).PracticePoint1.1.3.3:ConsiderinitiationoftreatmentsforCKDatrstpresentationofdecreasedGFRorelevatedACRifCKDisdeemedlikelyduetopresenceofotherclinicalindicators.1.1.4Evaluationofcause
PracticePoint1.1.4.1:EstablishthecauseofCKDusingclinicalcontext,personalandfamilyhistory,socialandenvi-ronmentalfactor

---

### Chunk 3/30
**Article:** Proteinuria (2023)
**Journal:** StatPearls Publishing
**Section:** abstract | **Similarity:** 0.554

Comprehensive clinical review of proteinuria covering definition, etiology, diagnosis, and management. Proteinuria serves as a marker of kidney damage assisting with diagnosis, prognosis, and therapy. Normal protein excretion is <150 mg/24 hours. Severity categories: nephritic range (150-3,500 mg/24h) and nephrotic range (>3,500 mg/24h). Urine dipstick provides semi-quantitative screening but has important limitations: false-positives with dehydration, alkaline urine, UTI; inability to detect positively charged proteins. UK CKD guidelines define significant proteinuria as UPCR >45 mg/mmol. Quantitative methods (spot UPCR preferred for convenience and reliability, or 24-hour urine collection for accuracy) are indicated when dipstick is positive or proteinuria >1 g/day is suspected. Early detection and quantification are crucial given associations with progressive renal disease and cardiovascular complications.

---

### Chunk 4/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.534

signicance[MGRS])98Presenceofpersistenthematuriaoralbuminuriaiscriticalindetermining
differentialdiagnosisGenetictestingAPOL1,COL4A3,COL4A4,COL4A5,NPHS1,UMOD,
HNF1B,PKD1,PKD2Evolvingasatoolfordiagnosis,increasedutilizationisexpected.
Recognitionthatgeneticcausesaremorecommonandmaypresent
withoutclassicfamilyhistory99,100ANCA,antineutrophilcytoplasmicantibody;APOL1,apolipoprotein1;COL4A,typeIVcollagenalphachain;CT,computedtomography;GBM,glomerularbasementmembrane;HNF1B,hepatocytenuclearfactor1B;MRI,magneticresonanceimaging;NPHS1,congenitalnephroticsyndrome;PKD1,polycystickidneydisease-1;PKD2,polycystickidneydisease-2;PLA2R,M-typephospholipaseA2receptor;UMOD,uromodulin.

---

### Chunk 5/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.532

andthe“new”stagesbetterreecttheirriskassocia-tions.Forthatreason,whereavailable,cystatinCshouldbe
addedtocreatinineforthepurposeofestimatingGFRfor
CKDdiagnosisandstaging.1.1.3EvaluationofchronicityPracticePoint1.1.3.1:Proofofchronicity(durationofa
minimumof3months)canbeestablishedby:(i)reviewofpastmeasurements/estimationsofGFR;(ii)reviewofpastmeasurementsofalbuminuriaor
proteinuriaandurinemicroscopicexaminations;(iii)imagingndingssuchasreducedkidneysizeandreductionincorticalthickness;(iv)kidneypathologicalndingssuchasbrosisandatrophy;(v)medicalhistory,especiallyconditionsknownto
causeorcontributetoCKD;(vi)repeatmeasurementswithinandbeyondthe3-
monthpoint.PracticePoint1.1.3.2:Donotassumechronicitybased
uponasingleabnormallevelforeGFRandACR,asthendingcouldbetheresultofarecentacutekidneyinjury(AKI)eventoracutekidneydisease(AKD).PracticePoint1.1.3.3:ConsiderinitiationoftreatmentsforCKDatrstpresentationofdecreasedGFRorelevatedACRifCKDisdeemedlikelyduetopresenceofotherclinicalindic

---

### Chunk 6/30
**Article:** Urea levels and cardiovascular disease in patients with chronic kidney disease (2022)
**Journal:** Nephrology Dialysis Transplantation
**Section:** other | **Similarity:** 0.527

20.3(5.1)0%Ageatbaseline(years)69[61–77]68[60–76]69[61–77]69[61–77]0.130%Men,n(%)666566670.710Smoking,n(%)0.030.8Never-smoker,n(%)40.644.539.637.7Currentsmoker,n(%)12.611.711.814.4Formersmoker,n(%)46.843.848.547.9eGFRatbaseline(mL/min/1.73m)33.5(11.6)43.5(9.9)32.6(8.9)24.5(7.0)<0.0010%Albumin-orprotein-to-creatinineratio<0.0018.0A1(normaltomildlyincreased),n(%)28.642.127.016.9A2(moderatelyincreased),n(%)31.831.833.729.7A3(severelyincreased),n(%)39.626.139.253.4Bodymassindex(kg/m)28.8(5.8)28.3(5.2)28.7(5.9)29.5(6.3)<0.0012.0%Diabetes,n(%)44.836.843.953.6<0.0010.2Systolicbloodpressure(mmHg)142(20)142(20)142(21)143(20)0.322.3%Historyofcardiovasculardisease,n(%)53.947.354.659.6<0.0011.3Anaemia,n(%)38.321.135.857.8<0.0010.3Serumbicarbonate(mmol/L)25.0(3.4)25.8(3.1)24.9(3.3)24.1(3.6)<0.0016.9%Serumalbumin(g/L)40.4(4.5)40.6(4.4)40.5(4.2)39.9(4.9)0.00915.2%High-sensitivityC-reactiveprotein(mg/L)2.5[1.1–5.9]2.2[1.1–5.0]2.5[1.1–5.4]2.9[1.2–7.1]<0.00117.6%Historyofacutekidneyinj

---

### Chunk 7/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.523

RifCKDisdeemedlikelyduetopresenceofotherclinicalindicators.1.1.4Evaluationofcause
PracticePoint1.1.4.1:EstablishthecauseofCKDusingclinicalcontext,personalandfamilyhistory,socialandenvi-ronmentalfactors,medications,physicalexamination,laboratorymeasures,imaging,andgeneticandpathologicdiagnosis(Figure8).
www.kidney-international.orgsummaryofrecommendationstatementsandpracticepointsKidneyInternational(2024)105(Suppl4S),S117–S314S149

PracticePoint1.1.4.2:Useteststoestablishacausebasedonresourcesavailable(Table622,98-100).Recommendation1.1.4.1:Wesuggestperformingakidneybiopsyasanacceptable,safe,diagnostictesttoevaluatecauseandguidetreatmentdecisionswhenclinicallyappropriate(2D).1.2EvaluationofGFR1.2.1OtherfunctionsofkidneysbesidesGFR
PracticePoint1.2.1.1:Usetheterm“GFR”whenreferringtothespecickidneyfunctionofglomerularltration.Usethemoregeneralterm“kidneyfunction(s)”whendealingwiththetotalityoffunctionsofthekidney.

---

### Chunk 8/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.521

tNZJMed.1992;22:334–337.338.GilbertRE,AkdenizA,JerumsG.Detectionofmicroalbuminuriaindiabeticpatientsbyurinarydipstick.DiabetesResClinPract.1997;35:57–60.339.GrazianiMS,GambaroG,MantovaniL,etal.Diagnosticaccuracyofa
reagentstripforassessingurinaryalbuminexcretioninthegeneralpopulation.NephrolDialTransplant.2008;24:1490–1494.340.GuyM,NewallR,BorzomatoJ,etal.Useofarst-lineurineprotein-to-creatinineratiostriptestonrandomurinestoruleoutproteinuriain
patientswithchronickidneydisease.NephrolDialTransplant.2008;24:1189–1193.341.GuyM,NewallR,BorzomatoJ,etal.Diagnosticaccuracyoftheurinary
albumin:creatinineratiodeterminedbytheCLINITEKMicroalbuminandDCA2000þfortherule-outofalbuminuriainchronickidneydisease.ClinChimActa.2008;399:54–58.342.HasslacherC,MullerP,SchlipfenbacherRL.ResultsofamulticentrestudyforthedeterminationofmicroalbuminuriawithMicral-Test.KlinischesLabor.1995;41:441–447.343.HodelNC,HamadA,ReitherK,etal.Comparisonoftwodifferent
semiquantitativeurinarydipsticktestswithal

---

### Chunk 9/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.521

ropriatetreatment.KeyinformationBalanceofbenetsandharms.Thebenetsofkidneybi-opsyintermsofdiagnosis,prognosis,andplanningappro-
priatetreatmentforboththepersonwithCKDandhealthcare
PhysicalexamNephrotoxicmedicationsSymptoms and signsof urinary tractabnormalitiesSymptoms and signsof systemic diseasesLaboratory tests, imaging, and tissue sample, such as:• Urinalysis and urine sediment• Urine albumin-to-creatinine ratio• Serologic tests
• Ultrasound• Kidney biopsy• Genetic testingMedicalhistorySocial andenvironmentalhistoryObtain careful family historyfor possible genetic causes,including family pedigree for CKD
Figure8|Evaluationofcauseofchronickidneydisease(CKD).

---

### Chunk 10/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.520

dneyinjury(AKI)eventoracutekidneydisease(AKD).PracticePoint1.1.3.3:ConsiderinitiationoftreatmentsforCKDatrstpresentationofdecreasedGFRorelevatedACRifCKDisdeemedlikelyduetopresenceofotherclinicalindicators.Kidneydiseasesmaybeacuteorchronic.1,97Weexplicitlyyetarbitrarilydenethedurationofaminimumof3months(>90days)asdelineating“chronic”kidneydisease.TherationalefordeningchronicityistodifferentiateCKDfromAKDs(suchasacuteglomerulonephritis[GN]),includingAKI,whichmayrequiredifferenttimelinesfor
initiationoftreatments,differentinterventions,andhavedifferentetiologiesandoutcomes.Thedurationofkidneydiseasemaybedocumentedorinferredbasedonthe
clinicalcontext.Forexample,apersonwithdecreasedGFR
orkidneydamageduringanacuteillness,withoutprior
documentationofkidneydisease,maybeinferredtohave
AKD.ResolutionoverdaystoweekswouldconrmthediagnosisofAKIfromavarietyofdifferentcauses.A
personwithsimilarndingsintheabsenceofanacuteillnessmaybeinferredtohaveCKD,andiffollowedover
time,wouldbeconrmedtoha

---

### Chunk 11/30
**Article:** Automated urine sediment analyzers underestimate the severity of hematuria in glomerular diseases (2021)
**Journal:** Scientific Reports
**Section:** results | **Similarity:** 0.519

and digital imaging in one system: A multicenter performance evaluation of the cobas 6500 urine work area. Pract. Lab. Med. 17, e00139 (2019). 30. aller, T. R. & Wang, L. P. Evaluation of asymptomatic microscopic hematuria in adults. Am. Fam. Phys. 60, 11431152 (1999). 31. Palsson, R., Srivastava, A. & Waikar, S. S. Performance of the automated urinalysis in diagnosis of proliferative glomerulonephritis. Kidney Int. Rep. 4, 723727 (2019). 32. Orlandi, P. F. et al. Hematuria as a risk factor for progression of chronic kidney disease and death: Findings from the chronic renal insuciency cohort (CRIC) study. BMC Nephrol. 19, 150 (2018). 33. Schulman, G. et al. Risk factors for progression of chronic kidney disease in the EPPIC trials and the eect of AST-120. Clin. Exp. Nephrol. 22, 299308 (2018). 34. Mashitani, T. et al.

---

### Chunk 12/30
**Article:** Treatment of hyponatremia: comprehension and best clinical practice (2025)
**Journal:** Clinical and Experimental Nephrology
**Section:** results | **Similarity:** 0.516

e plasma sodium levels in patients with the syndrome of inappropriate anti-diuresis. J Am Soc Nephrol. 2020;31:61524. 
32. Refardt J, Imber C, Nobbenhuis R, Sailer CO, Haslbauer A, Mon-nerat S, Bathelt C, Vogt DR, Berres M, Winzeler B, Bridenbaugh SA, Christ-Crain M. Treatment eﬀect of the SGLT2 inhibitor empagliﬂozin on chronic syndrome of inappropriate antidiuresis: results of a randomized, double-blind, placebo-controlled, crosso-ver trial. J Am Soc Nephrol. 2023;34:32232. 
33. Hamblin PS, Wong R, Ekinci EI, Fourlanos S, Shah S, Jones AR, Hare MJL, Calder GL, Epa DS, George EM, Giri R, Kotowicz MA, Kyi M, Lafontaine N, MacIsaac RJ, Nolan BJ, ONeal DN, Renouf D, Varadarajan S, Wong J, Xu S, Bach LA. SGLT2 inhibi-tors increase the risk of diabetic ketoacidosis developing in the community and during hospital admission. J Clin Endocrinol Metab. 2019;104:307787. 
34. Schrier RW, Gross P, Gheorghiade M, Berl T, Verbalis JG, Czer-wiec FS, Orlandi C, SALT Investigators.

---

### Chunk 13/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.515

eoplewithriskfactorsforCKD,thereisahigherprobabilitythattheperson
doeshaveCKDevenwithanunexpectednding.Subsequenttestingshouldbeperformedtoconrmthediagnosisandtocompletetheevaluation,asisrequired.Timingoftherepeat
sampleshouldbedeterminedbasedonclinicalsetting
includingriskfactorsforCKDaswellasconcernforAKI/AKD.Hematuriaiscommonandassociatedwithriskforsub-sequentdevelopmentofCKD.65Thereareseveralcausesof
chapter1www.kidney-international.orgS170KidneyInternational(2024)105(Suppl4S),S117–S314

transienthematuria.Persistenthematuriamayindicateglomerulardisease,otherkidneydiseases,orurologic
diseaseincludinggenitourinarymalignancy.Thus,persistenthematuriashouldpromptfurtherinvestigation.66,67SpecialconsiderationsPediatricconsiderations.Peoplewhoarebornpreterm,especiallyifalsosmallforgestationalage,areatincreasedrisk
forCKDandkidneyfailure.Thisislargelyrelatedtodecreasednephronnumber.68–70AdditionalinsultsafterbirthsuchasneonatalAKIandchildhoodobesitycanfurtherin-
creasetheriskofCKD.7

---

### Chunk 14/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.512

anginglistof“atrisk”groups.27TestingforCKDwithoutregardtoagegeneratescontro-versy.Thoseinolderagegroupsexperiencethegreatest
burdenofCKDandarealsoatthehighestriskforcardio-
vascularcomplications.Aswithotherdetectionprogramslike
cancerdetection,CKDdetectioneffortsshouldbeindividu-
alizedbasedontheperson’sgoalsofcareandsuitabilityfortreatment.PracticePoint1.1.1.2:Followingincidentaldetectionofelevatedurinaryalbumin-to-creatinineratio(ACR),he-
maturia,orlowestimatedGFR(eGFR),repeatteststo
conrmpresenceofCKD.ThereisknownbiologicalandanalyticalvariabilityinSCrandinurinealbuminorurineproteinnotrelatedtotheir
propertiesasmarkersofkidneydisease.Inpeoplewithout
riskfactorsforCKD,thereisalowpretestprobabilityforCKD.Thus,anyunexpectedresultsshouldbeveriedbeforediagnosingapersonashavingCKD.InpeoplewithriskfactorsforCKD,thereisahigherprobabilitythattheperson
doeshaveCKDevenwithanunexpectednding.Subsequenttestingshouldbeperformedtoconrmthediagnosisandtocompletetheevaluation,asisrequired.Ti

---

### Chunk 15/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.511

e-basedequations.TransplInt.2012;25:527–536.162.JanusN,Launay-VacherV,ByloosE,etal.CancerandrenalinsufciencyresultsoftheBIRMAstudy.BrJCancer.2010;103:1815–1821.163.Launay-VacherV,JanusN,DerayG.Renalinsufciencyandcancertreatments.ESMOOpen.2016;1:e000091.164.NaSY,SungJY,ChangJH,etal.Chronickidneydiseaseincancerpatients:anindependentpredictorofcancer-specicmortality.AmJNephrol.2011;33:121–130.165.RosnerMH,JhaveriKD,McMahonBA,etal.Onconephrology:the
intersectionsbetweenthekidneyandcancer.CACancerJClin.2021;71:47–77.166.SoveriI,BergUB,BjorkJ,etal.MeasuringGFR:asystematicreview.AmJKidneyDis.2014;64:411–424.167.WhiteCA,AkbariA,AllenC,etal.Simultaneousglomerularltrationratedeterminationusinginulin,iohexol,and99mTc-DTPAdemonstratestheneedforcustomizedmeasurementprotocols.KidneyInt.2021;99:957–966.168.XieP,HuangJM,LiuXM,etal.(99m)Tc-DTPArenaldynamicimaging
methodmaybeunsuitabletobeusedasthereferencemethodin
investigatingthevalidityofCDK-EPIequationfordetermining
glomerularltrati

---

### Chunk 16/30
**Article:** SAM (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.511

ão. O diagnóstico é sugerido por um aumento de 20% + 2 ng/mL acima do valor basal do paciente durante uma crise. No entanto, o palestrante relata que raramente vê resultados positivos.
    -   **N-metil-histamina urinária (urina de 24h):** Considerado um marcador útil. Um valor acima de 60 microgramas por grama em 24 horas é sugestivo.
    -   **Outros mediadores:** Cromogranina A (pode estar elevada pelo uso de inibidores da bomba de prótons), heparina (potencialmente o melhor marcador, mas ainda não validado), prostaglandinas e leucotrienos podem estar elevados, mas não são validados para diagnóstico.
3.  **Biópsia do Trato Gastrointestinal:** A endoscopia ou colonoscopia com biópsias e análise por imuno-histoquímica pode revelar um aumento no número de mastócitos (>20 por campo de grande aumento), o que apoia o diagnóstico.
4.

---

### Chunk 17/30
**Article:** Urea levels and cardiovascular disease in patients with chronic kidney disease (2022)
**Journal:** Nephrology Dialysis Transplantation
**Section:** other | **Similarity:** 0.511

.0016.9%Serumalbumin(g/L)40.4(4.5)40.6(4.4)40.5(4.2)39.9(4.9)0.00915.2%High-sensitivityC-reactiveprotein(mg/L)2.5[1.1–5.9]2.2[1.1–5.0]2.5[1.1–5.4]2.9[1.2–7.1]<0.00117.6%Historyofacutekidneyinjury,n(%)22.520.020.926.60.0047.7Diureticprescriptionatbaseline,n(%)54.039.951.170.7<0.0010.3PPIprescriptionatbaseline,n(%)31.427.933.932.40.020.3RASiprescriptionatbaseline,n(%)76.972.380.278.0<0.0010.3Statinprescriptionatbaseline,n(%)58.951.461.064.1<0.0010.%Antiplateletprescriptionatbaseline,n(%)42.137.643.145.60.0030.3
eGFR,estimatedglomerularltrationrate,basedontheCKD-EPIequation;RASi,renin-angiotensinsysteminhibitor;PPI,protonpumpinhibitor;T,tertile.Dataarequotedasthefrequency(%),mean(standarddeviation)orthemedian[interquartilerange].AnANOVAoraKruskal–WallisorChi-squaredtestwasusedtocomparegroups.

---

### Chunk 18/30
**Article:** Resistência Insulínica (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.509

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

### Chunk 19/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.508

kHM,Begin-HeickN,AcharyaC,etal.Automateddeterminationof
urineandcerebrospinaluidproteinswithCoomassiebrilliantblueandtheAbbottABA-100.ClinBiochem.1980;13:81–83.280.MarshallT,WilliamsKM.Totalproteindeterminationinurine:eliminationofadifferentialresponsebetweentheCoomassieblueandpyrogallolredproteindye-bindingassays.ClinChem.2000;46:392–398.281.GinsbergJM,ChangBS,MatareseRA,etal.Useofsinglevoidedurine
samplestoestimatequantitativeproteinuria.NEnglJMed.1983;309:1543–1546.282.PriceCP,NewallRG,BoydJC.Useofprotein:creatinineratio
measurementsonrandomurinesamplesforpredictionofsignicantproteinuria:asystematicreview.ClinChem.2005;51:1577–1586.283.BeethamR,CattellWR.Proteinuria:pathophysiology,signicanceandrecommendationsformeasurementinclinicalpractice.AnnClinBiochem.1993;30(Pt5):425–434.284.KeaneWF,EknoyanG.Proteinuria,albuminuria,risk,assessment,detection,elimination(PARADE):apositionpaperoftheNational
KidneyFoundation.AmJKidneyDis.1999;33:1004–1010.285.ClaudiT,CooperJG.Compar

---

### Chunk 20/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.507

specickidneyfunctionofglomerularltration.Usethemoregeneralterm“kidneyfunction(s)”whendealingwiththetotalityoffunctionsofthekidney.
PhysicalexamNephrotoxicmedicationsSymptoms and signsof urinary tractabnormalitiesSymptoms and signsof systemic diseasesLaboratory tests, imaging, and tissue sample, such as:• Urinalysis and urine sediment• Urine albumin-to-creatinine ratio• Serologic tests
• Ultrasound• Kidney biopsy• Genetic testingMedicalhistorySocial andenvironmentalhistoryObtain careful family historyfor possible genetic causes,including family pedigree for CKD
Figure8|Evaluationofcauseofchronickidneydisease(CKD).

---

### Chunk 21/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.507

22:369.318.McTaggartMP,NewallRG,HirstJA,etal.Diagnosticaccuracyofpoint-of-
caretestsfordetectingalbuminuria:asystematicreviewandmeta-analysis.AnnInternMed.2014;160:550–557.319.AbitbolC,ZillerueloG,FreundlichM,etal.Quantitationofproteinuriawithurinaryprotein/creatinineratiosandrandomtestingwithdipsticksinnephroticchildren.JPediatr.1990;116:243–247.320.AgardhCD.Anewsemiquantitativerapidtestforscreeningformicroalbuminuria.PracticalDiabetes.1993;10:146–147.321.AgarwalR,PanesarA,LewisRR.Dipstickproteinuria:canitguide
hypertensionmanagement?AmJKidneyDis.2002;39:1190–1195.322.AroraS,LongT,MenchineM.Testcharacteristicsofurinedipstickfor
identifyingrenalinsufciencyinpatientswithdiabetes.WestJEmergMed.2011;12:250–253.323.ChangCC,SuMJ,HoJL,etal.Theefcacyofsemi-quantitativeurineprotein-to-creatinine(P/C)ratioforthedetectionofsignicantproteinuriainurinespecimensinhealthscreeningsettings.Springerplus.2016;5:1791.324.ChoMC,JiM,KimSY,etal.EvaluationoftheURiSCANsupercassetteACRsemiquantita

---

### Chunk 22/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XIV (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.506

- [ ] 2. Para pacientes com síndrome metabólica, considerar estratégia Mediterrânea organizada; preferir versão com restrição calórica combinada com atividade física, monitorando resultados aos 6 e 12 meses.
- [ ] 3. Solicitar perfil lipídico ampliado em casos selecionados: HDL, triglicerídeos, insulina, PCR, LDL oxidado e, quando indicado, subfracionamento de LDL; evitar decisões baseadas apenas em LDL total.
- [ ] 4. Planejar exames de risco cardiovascular conforme necessidade: score de cálcio coronariano e angiotomografia de coronárias (incluindo avaliação de placas moles) quando o contexto clínico justificar.
- [ ] 5. Revisar e atualizar protocolos internos sobre álcool: remover recomendações de consumo “cardioprotetor”; educar pacientes sobre riscos de câncer e piora do sono; avaliar predisposição/sensibilidade individual.
- [ ] 6.

---

### Chunk 23/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.506

solutionoverdaystoweekswouldconrmthediagnosisofAKIfromavarietyofdifferentcauses.A
personwithsimilarndingsintheabsenceofanacuteillnessmaybeinferredtohaveCKD,andiffollowedover
time,wouldbeconrmedtohaveCKD.Inbothcases,repeatascertainmentofGFRandkidneydamageis
recommendedforaccuratediagnosisandstaging.The
timingoftheevaluationdependsonclinicaljudgment,with
earlierevaluationforthosesuspectedofhavingAKIand
laterevaluationforthosesuspectedofhavingCKD.ForpeoplewithriskfactorsforCKD,delayingdiag-nosisforthesakeofconrmingchronicitycandelaycare.Manypeoplemaynotrecognizetheimportanceof
arepeatvisitiftreatmenthadnotbeeninitiated.Thus,
initiatingtreatmentbothallowsforearlierintervention
andalsoindicatestopeopletheimportanceofthe
disease.

---

### Chunk 24/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** methods | **Similarity:** 0.506

ey-international.orgS148KidneyInternational(2024)105(Suppl4S),S117–S314

SummaryofrecommendationstatementsandpracticepointsChapter1:EvaluationofCKD1.1DetectionandevaluationofCKD1.1.1DetectionofCKDPracticePoint1.1.1.1:Testpeopleatriskforandwithchronickidneydisease(CKD)usingbothurinealbuminmea-surementandassessmentofglomerularltrationrate(GFR).PracticePoint1.1.1.2:Followingincidentaldetectionofelevatedurinaryalbumin-to-creatinineratio(ACR),hematuria,orlowestimatedGFR(eGFR),repeatteststoconrmpresenceofCKD.1.1.2MethodsforstagingofCKDRecommendation1.1.2.1:InadultsatriskforCKD,werecommendusingcreatinine-basedestimatedglomer-ularltrationrate(eGFRcr).IfcystatinCisavailable,theGFRcategoryshouldbeestimatedfromthecombinationofcreatinineandcystatinC(creatinineandcystatinC–basedestimatedglomerularltrationrate[eGFRcr-cys])(1B).1.1.3EvaluationofchronicityPracticePoint1.1.3.1:Proofofchronicity(durationofaminimumof3months)canbeestablishedby:(i)reviewofpastmeasurements/estimationsofGFR;(ii)revie

---

### Chunk 25/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.503

hohavemetabolicdisordersandareonaverylow–proteindiet,acystatinC–basedequationislikelymoreappropriate.1.3EvaluationofalbuminuriaAlbuminuriareferstoabnormallossofalbuminintheurine(urineACR$30mg/gor$3mg/mmol).Albuminisonetypeofplasmaproteinfoundintheurineinnormalsubjectsand
inlargerquantityinpeoplewithkidneydisease.IntheKDIGO
2012ClinicalPracticeGuidelinefortheEvaluationand
ManagementofChronicKidneyDisease,1clinicalterminologywaschangedtofocusonalbuminuriaratherthanproteinuria
asalbuministheprincipalcomponentofurinaryproteininmostkidneydiseases.Epidemiologicdatademonstrateastrongrelationshipbetweenthequantityofurinealbuminwithboth
kidneyandCVDriskandobservedCVDevenatverylow
levels,andassaystomeasurealbuminaremorepreciseand
sensitivethanassaystomeasureurineprotein.Wereferto
albuminuriaorurinealbuminwhendiscussinggeneral
conceptsandwillrefereithertototalprotein,albumin,orother
specicproteinswhendiscussingthatparameterspecically.1.3.1GuidanceforphysiciansandotherhealthcareprovidersPrac

---

### Chunk 26/30
**Article:** Fisiologia e Bioquímica do Sistema Imune II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.500

olicitar e registrar parâmetros essenciais (PCR-us, VHS); calcular índices estimáveis (HGI, TAIG); considerar complementares (TNF-α, IL-6, GPx, MDA, antioxidantes totais).
- [ ] 11. Monitorar trimestralmente HbA1c, glicemia de jejum e HGI; estabelecer metas individuais (ex.: reduzir de 6,1 para ~5,5; longo prazo ~5,3).
- [ ] 12. Calcular TG/HDL e integrar com lipidograma/SREBP1c/2; evitar combinação de gordura saturada com açúcar e excesso de saturadas em múltiplas refeições.
- [ ] 13. Avaliar ferro, ferritina, transferrina e saturação (20–50%; evitar <20%); interpretar com hepcidina/SREBP1c e quadro inflamatório.
- [ ] 14. Medir TNF-α (<8,1) e IL-6 (<3,4) para acompanhar atividade inflamatória; relacionar com obesidade inflamada.
- [ ] 15. Calcular HOMA-β (167–175) e HOMA-IR (<2,15); buscar glicemia 60–90 e insulina ~6–7.
- [ ] 16. Monitorar homocisteína (<7,9) e PCR; usar PCR-us; documentar crises (PCR >1.000) e conduzir manejo apropriado.
- [ ] 17.

---

### Chunk 27/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.500

rshallSM,AlbertiKG.Screeningforearlydiabeticnephropathy.AnnClinBiochem.1986;23(Pt2):195–197.290.MillerWG,BrunsDE,HortinGL,etal.Currentissuesinmeasurementandreportingofurinaryalbuminexcretion.ClinChem.2009;55:24–38.291.ChitaliaVC,KothariJ,WellsEJ,etal.Cost-benetanalysisandpredictionof24-hourproteinuriafromthespoturineprotein-creatinineratio.ClinNephrol.2001;55:436–447.292.CoteAM,BrownMA,LamE,etal.Diagnosticaccuracyofurinaryspotprotein:creatinineratioforproteinuriainhypertensivepregnant
women:systematicreview.BMJ.2008;336:1003–1006.293.DysonEH,WillEJ,DavisonAM,etal.Useoftheurinaryproteincreatinine
indextoassessproteinuriainrenaltransplantpatients.NephrolDialTransplant.1992;7:450–452.294.Leanos-MirandaA,Marquez-AcostaJ,Romero-ArauzF,etal.Protein:
creatinineratioinrandomurinesamplesisareliablemarkerof
increased24-hourproteinexcretioninhospitalizedwomenwith
hypertensivedisordersofpregnancy.ClinChem.2007;53:1623–1628.295.LemannJJr,DoumasBT.Proteinuriainhealthanddiseaseassessedby

---

### Chunk 28/30
**Article:** Urea levels and cardiovascular disease in patients with chronic kidney disease (2022)
**Journal:** Nephrology Dialysis Transplantation
**Section:** results | **Similarity:** 0.499

vel,bloodhaemoglobin,serumalbuminandalbumin-orprotein-to-creatinineratio(Supplementarydata,TableS1)wererecorded.TheGFRwasestimatedusingtheChronicKidneyDiseaseEpidemiologyCollaboration(CKD-EPI)creatinineequationandtheisotopedilutionmassspectrometry-traceablecreatinineconcentrationdeterminedinaJaeassay(in38%ofparticipants)oranenzymaticassay(in57%;theassaytypewasunknownfor5%)[15].Patientswereaskedtobringalltheirdrugprescriptionsforthepreceding3monthstotheenrolmentvisit.TheCRAsusedanelectroniccasereportform[linkedtotheinternationalAnatomicalTherapeuticandChemical(ATC)thesaurus]toenterstandardizedATCcodes.Longitudinalclinicaldata(nephrologyconsultations,hos-pitaladmissions,laboratorytestresults,medicationsandanytransitiontodialysisortransplantation)werecollectedeveryyearfrommedicalrecordsandpatientinterviews.StudyoutcomesCVeventswereassessedcarefullyaccordingtotheCardio-vascularandStrokeEndpointDenitionsforClinicalTrials[16].Theprimaryendpointwastheoccurrenceofarstfatalornonfatalatheroma

---

### Chunk 29/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.499

studies,adoublingoftheACRwithina2-yeardurationisassociatedwithanincreasein
theriskofprogressiontokidneyfailureby50%–100%.391,392
Albuminuria categoriesDescription and rangeGFR categories (ml/min/1.73 m2)Description and rangeA1G1≥90G260–89G3a45–59G3b30–44G415–29G5<15Kidney failureSeverely decreasedModerately toseverely decreasedMildly tomoderately decreasedMildly decreasedNormal or highA2A3Normal to mildlyincreasedModeratelyincreasedSeverelyincreased<30 mg/g<3 mg/mmolScreen1Screen1Treat1Treat1Treat3Treat3Treat1Treat2Treat2Treat3Treat3Treat3Treat*3Treat4+Treat*3Treat4+Treat4+Treat4+30–299 mg/g3–29 mg/mmol≥300 mg/g≥30 mg/mmolCKD is classified based on:�Cause (C)�GFR (G)� Albuminuria (A)Low risk (if no other markers of kidney disease, no CKD)Moderately increased riskHigh risk
Very high risk
Figure13|Frequencyofmonitoringglomerularltrationrate(GFR)andalbuminuriainpeoplewithchronickidneydisease(CKD).AlbuminuriaandGFRgridreectstheriskofprogressionbyintensityofcolor

---

### Chunk 30/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.499

chronickidney
disease.AmJKidneyDis.2005;46:S1–S122.555.KovesdyCP,MatsushitaK,SangY,etal.Serumpotassiumandadverse
outcomesacrosstherangeofkidneyfunction:aCKDPrognosisConsortiummeta-analysis.EurHeartJ.2018;39:1535–1542.556.ClaseCM,CarreroJJ,EllisonDH,etal.Potassiumhomeostasisandmanagementofdyskalemiainkidneydiseases:conclusionsfroma
KidneyDisease:ImprovingGlobalOutcomes(KDIGO)ControversiesConference.KidneyInt.2020;97:42–61.557.LewisEJ,HunsickerLG,ClarkeWR,etal.Renoprotectiveeffectoftheangiotensin-receptorantagonistirbesartaninpatientswith
nephropathyduetotype2diabetes.NEnglJMed.2001;345:851–860.558.CollinsAJ,PittB,ReavenN,etal.Associationofserumpotassiumwithall-causemortalityinpatientswithandwithoutheartfailure,chronic
kidneydisease,and/ordiabetes.AmJNephrol.2017;46:213–221.559.KorgaonkarS,TileaA,GillespieBW,etal.Serumpotassiumand
outcomesinCKD:insightsfromtheRRI-CKDcohortstudy.ClinJAmSocNephrol.2010;5:762–769.560.EinhornLM,ZhanM,HsuVD,etal.Thefrequencyofhyperkalemiaand
itssi

---

