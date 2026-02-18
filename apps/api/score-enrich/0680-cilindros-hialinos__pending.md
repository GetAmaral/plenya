# ScoreItem: Cilindros Hialinos

**ID:** `c77cedd3-2800-75bd-a7d7-9393c65ca901`
**FullName:** Cilindros Hialinos (Exames - Laboratoriais)
**Unit:** cilindros/campo

**Preparation Metadata:**
- Quality Grade: **GOOD**
- Total Chunks: 30 de 3 artigos
- Avg Similarity: 0.566

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `c77cedd3-2800-75bd-a7d7-9393c65ca901`.**

```json
{
  "score_item_id": "c77cedd3-2800-75bd-a7d7-9393c65ca901",
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

**ScoreItem:** Cilindros Hialinos (Exames - Laboratoriais)
**Unidade:** cilindros/campo

**30 chunks de 3 artigos (avg similarity: 0.566)**

### Chunk 1/30
**Article:** Automated urine sediment analyzers underestimate the severity of hematuria in glomerular diseases (2021)
**Journal:** Scientific Reports
**Section:** other | **Similarity:** 0.605

hy: A cohort study. Am. J. Kidney Dis. 76, 9099 (2020). 21. Wu, Y. et al. e association of hematuria on kidney clinicopathologic features and renal outcome in patients with diabetic nephropathy: A biopsy-based study. J. Endocrinol. Investig. 43, 12131220 (2020). 22. Imran, S., Eva, G., Christopher, S., Flynn, E. & Henner, D. Is specic gravity a good estimate of urine osmolality? J. Clin. Lab. Anal. 24, 426430 (2010). 23. Sawka, M. N. et al. American College of Sports Medicine position stand. Exercise and uid replacement. Med. Sci. Sports Exerc. 39, 377390 (2007). 24. Sharp, V. J., Barnes, K. T. & Erickson, B. A. Assessment of asymptomatic microscopic hematuria in adults. Am. Fam. Phys. 88, 747754 (2013). 25. Grossfeld, G. D. et al. Evaluation of asymptomatic microscopic hematuria in adults: e American Urological Association best practice policyPart I: Denition, detection, prevalence, and etiology. Urology 57, 599603 (2001). 26. Yuste, C. et al.

---

### Chunk 2/30
**Article:** Automated urine sediment analyzers underestimate the severity of hematuria in glomerular diseases (2021)
**Journal:** Scientific Reports
**Section:** discussion | **Similarity:** 0.596

pH of the urine was categorized as follows; 5, 6, 6.5, 7, 8 and 9.e Urisys 2400 and Cobas u 601 urine analyzers both contain a built-in refractometer, and the specic gravity (SG) was measured via refractometry. e scale of urine SG ranged from 1.000 to 1.050. In dilute urine, the RBCs absorb water, swell, and may rupture. In concentrated urine, the RBCs tend to shrink and become crenated. A urine SG of 1.010 corresponds to approximately 300mOsm/kg, similar to the osmolarity of  plasma22, and a urine SG > 1.020 oen indicates  dehydration23. In this study, therefore, urine samples with SG < 1.010 and > 1.020 were considered dilute and concentrated urine, respectively.Statistical analysis. e data are expressed as the median (interquartile range). e values obtained for the two groups were compared using the Mann–Whitney U test.

---

### Chunk 3/30
**Article:** Proteinuria (2023)
**Journal:** StatPearls Publishing
**Section:** abstract | **Similarity:** 0.593

Comprehensive clinical review of proteinuria covering definition, etiology, diagnosis, and management. Proteinuria serves as a marker of kidney damage assisting with diagnosis, prognosis, and therapy. Normal protein excretion is <150 mg/24 hours. Severity categories: nephritic range (150-3,500 mg/24h) and nephrotic range (>3,500 mg/24h). Urine dipstick provides semi-quantitative screening but has important limitations: false-positives with dehydration, alkaline urine, UTI; inability to detect positively charged proteins. UK CKD guidelines define significant proteinuria as UPCR >45 mg/mmol. Quantitative methods (spot UPCR preferred for convenience and reliability, or 24-hour urine collection for accuracy) are indicated when dipstick is positive or proteinuria >1 g/day is suspected. Early detection and quantification are crucial given associations with progressive renal disease and cardiovascular complications.

---

### Chunk 4/30
**Article:** Automated urine sediment analyzers underestimate the severity of hematuria in glomerular diseases (2021)
**Journal:** Scientific Reports
**Section:** discussion | **Similarity:** 0.593

1
Scientific Reports |        (2021) 11:20981  | https://doi.org/10.1038/s41598-021-00457-6
Automated urine sediment analyzers underestimate the severity of hematuria in glomerular diseasesWon Seok YangHematuria, either glomerular or extraglomerular, is defined as 3 or more red blood cells (RBCs)/high power field. Currently, urinalyses are commonly performed using automated urine sediment analyzers. To assess whether RBC counting by automated urine sediment analyzers is reliable for defining hematuria in glomerular disease, random specimen urinalyses of men with nephritic glomerular disease (7674 urinalyses) and bladder cancer (12,510 urinalyses) were retrospectively reviewed. Urine RBCs were counted by an automated urine sediment analyzer based on flow cytometry (UF-1000i, Sysmex Corporation) or digital image analysis (Cobas 6500, Roche Diagnostics GmbH). In about 20% of urine specimens, the specific gravity was less than 1.010, making the RBC counts unreliable.

---

### Chunk 5/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.586

 52.361.OlivariusND,MogensenCE.Danishgeneralpractitioners’estimationofurinaryalbuminconcentrationinthedetectionofproteinuriaand
microalbuminuria.BrJGenPract.1995;45:71–73.362.OstaV,NatoliV,DiéguezS.[Evaluationoftworapidtestsforthe
determinationofmicroalbuminuriaandtheurinaryalbumin/creatinine
ratio].AnPediatr(Barc).2003;59:131–137[inSpanish].363.OyaertM,DelangheJR.Semiquantitative,fullyautomatedurineteststrip
analysis.JClinLabAnal.2019;33:e22870.364.ParkerJL,KirmizS,NoyesSL,etal.Reliabilityofurinalysisfor
identicationofproteinuriaisreducedinthepresenceofotherabnormalitiesincludinghighspecicgravityandhematuria.UrolOncol.2020;38:853.e859–853.e915.365.ParsonsM,NewmanDJ,PugiaM,etal.Performanceofareagentstrip
deviceforquantitationoftheurinealbumin:creatinineratioinapoint
ofcaresetting.ClinNephrol.1999;51:220–227.366.ParsonsMP,NewmanDJ,NewallRG,etal.Validationofapoint-of-care
assayfortheurinaryalbumin:creatinineratio.ClinChem.1999;45:414–417.367.PendersJ,FiersT,DelangheJR.Quan

---

### Chunk 6/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.586

tNZJMed.1992;22:334–337.338.GilbertRE,AkdenizA,JerumsG.Detectionofmicroalbuminuriaindiabeticpatientsbyurinarydipstick.DiabetesResClinPract.1997;35:57–60.339.GrazianiMS,GambaroG,MantovaniL,etal.Diagnosticaccuracyofa
reagentstripforassessingurinaryalbuminexcretioninthegeneralpopulation.NephrolDialTransplant.2008;24:1490–1494.340.GuyM,NewallR,BorzomatoJ,etal.Useofarst-lineurineprotein-to-creatinineratiostriptestonrandomurinestoruleoutproteinuriain
patientswithchronickidneydisease.NephrolDialTransplant.2008;24:1189–1193.341.GuyM,NewallR,BorzomatoJ,etal.Diagnosticaccuracyoftheurinary
albumin:creatinineratiodeterminedbytheCLINITEKMicroalbuminandDCA2000þfortherule-outofalbuminuriainchronickidneydisease.ClinChimActa.2008;399:54–58.342.HasslacherC,MullerP,SchlipfenbacherRL.ResultsofamulticentrestudyforthedeterminationofmicroalbuminuriawithMicral-Test.KlinischesLabor.1995;41:441–447.343.HodelNC,HamadA,ReitherK,etal.Comparisonoftwodifferent
semiquantitativeurinarydipsticktestswithal

---

### Chunk 7/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.580

nephrologydowiththenewCKD-EPIequation?NephrolDialTransplant.2023;38:1–6.257.NymanU,BjorkJ,BergU,etal.ThemodiedCKiDstudyestimatedGFRequationsforchildrenandyoungadultsunder25yearsofage:
performanceinaeuropeanmulticentercohort.AmJKidneyDis.2022;80:807–810.258.HeathcoteKL,WilsonMP,QuestDW,etal.Prevalenceanddurationofexerciseinducedalbuminuriainhealthypeople.ClinInvestMed.2009;32:E261–E265.259.CarterJL,TomsonCR,StevensPE,etal.Doesurinarytractinfectioncauseproteinuriaormicroalbuminuria?Asystematicreview.NephrolDialTransplant.2006;21:3031–3037.260.McTaggartMP,StevensPE,PriceCP,etal.Investigationofapparentnon-
albuminuricproteinuriainaprimarycarepopulation.ClinChemLabMed.2013;51:1961–1969.261.HoweyJE,BrowningMC,FraserCG.Selectingtheoptimumspecimenfor
assessingslightalbuminuria,andastrategyforclinicalinvestigation:novel
usesofdataonbiologicalvariation.ClinChem.1987;33:2034–2038.262.BallantyneFC,GibbonsJ,O’ReillyDS.Urinealbuminshouldreplacetotalproteinfortheassessmentofglomerular

---

### Chunk 8/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.575

g.ClinNephrol.1999;51:220–227.366.ParsonsMP,NewmanDJ,NewallRG,etal.Validationofapoint-of-care
assayfortheurinaryalbumin:creatinineratio.ClinChem.1999;45:414–417.367.PendersJ,FiersT,DelangheJR.Quantitativeevaluationofurinalysistest
strips.ClinChem.2002;48:2236–2241.368.PoulsenPL,MogensenCE.Clinicalevaluationofatestforimmediateand
quantitativedeterminationofurinaryalbumin-to-creatinineratio.Abrief
report.DiabetesCare.1998;21:97–98.369.PugiaMJ,LottJA,KajimaJ,etal.Screeningschoolchildrenfor
albuminuria,proteinuriaandoccultbloodwithdipsticks.ClinChemLabMed.1999;37:149–157.370.SakaiN,FuchigamiH,IshizukaT,etal.Relationshipbetweena
urineprotein-to-creatinineratioof150mg/gramcreatinineanddipstickgradeinthehealthcheckup:substantialnumberoffalse-
negativeresultsforchronickidneydisease.TokaiJExpClinMed.2019;44:118–123.371.SalinasM,López-GarrigósM,FloresE,etal.Urinaryalbuminstripassay
asascreeningtesttoreplacequantitativetechnologyincertain
conditions.ClinChemLabMed.2018;57:204–209.

---

### Chunk 9/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.572

.Diagnosticaccuracyofpoint-of-caretestsfordetectingalbuminuria:a
systematicreviewandmeta-analysis.AnnIntMed.2014;160:550–557.318SoFtablesSupplementaryTableS5SearchdateJuly2022
Citationsscreened/includedstudies2184/65
SupplementaryFigureS5Chapter2RiskassessmentinpeoplewithCKDClinicalquestionArekidneyfailurepredictionequationsgoodpredictorsofprogression,kidneyfailure,orend-stagerenaldisease?PopulationAdults,children,andyoungpeoplewithCKDG1-G5PredictorKidneyfailureriskequations(e.g.,Tangriequation[KidneyFailureRiskEquation])OutcomesPrognosticperformance:Calibration(goodnessofmeasures,e.g.,R2,Brierscore,andHosmer-Lemeshowtest)Discrimination(e.g.,sensitivity/specicity;areaunderthecurve[AUC]fromreceiveroperatingcharacteristic[ROC]andareaunderthereceiveroperatingcharacteristiccurve[AUROC];C-statistic)StudydesignSystematicreview
ExistingsystematicreviewNationalInstituteforHealthandCareExcellence.Evidencereviewforthebestcombinationofmeasurestoidentifyincreasedriskofprogressioninadults,childr

---

### Chunk 10/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.572

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

### Chunk 11/30
**Article:** Automated urine sediment analyzers underestimate the severity of hematuria in glomerular diseases (2021)
**Journal:** Scientific Reports
**Section:** results | **Similarity:** 0.571

and digital imaging in one system: A multicenter performance evaluation of the cobas 6500 urine work area. Pract. Lab. Med. 17, e00139 (2019). 30. aller, T. R. & Wang, L. P. Evaluation of asymptomatic microscopic hematuria in adults. Am. Fam. Phys. 60, 11431152 (1999). 31. Palsson, R., Srivastava, A. & Waikar, S. S. Performance of the automated urinalysis in diagnosis of proliferative glomerulonephritis. Kidney Int. Rep. 4, 723727 (2019). 32. Orlandi, P. F. et al. Hematuria as a risk factor for progression of chronic kidney disease and death: Findings from the chronic renal insuciency cohort (CRIC) study. BMC Nephrol. 19, 150 (2018). 33. Schulman, G. et al. Risk factors for progression of chronic kidney disease in the EPPIC trials and the eect of AST-120. Clin. Exp. Nephrol. 22, 299308 (2018). 34. Mashitani, T. et al.

---

### Chunk 12/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.569

22:369.318.McTaggartMP,NewallRG,HirstJA,etal.Diagnosticaccuracyofpoint-of-
caretestsfordetectingalbuminuria:asystematicreviewandmeta-analysis.AnnInternMed.2014;160:550–557.319.AbitbolC,ZillerueloG,FreundlichM,etal.Quantitationofproteinuriawithurinaryprotein/creatinineratiosandrandomtestingwithdipsticksinnephroticchildren.JPediatr.1990;116:243–247.320.AgardhCD.Anewsemiquantitativerapidtestforscreeningformicroalbuminuria.PracticalDiabetes.1993;10:146–147.321.AgarwalR,PanesarA,LewisRR.Dipstickproteinuria:canitguide
hypertensionmanagement?AmJKidneyDis.2002;39:1190–1195.322.AroraS,LongT,MenchineM.Testcharacteristicsofurinedipstickfor
identifyingrenalinsufciencyinpatientswithdiabetes.WestJEmergMed.2011;12:250–253.323.ChangCC,SuMJ,HoJL,etal.Theefcacyofsemi-quantitativeurineprotein-to-creatinine(P/C)ratioforthedetectionofsignicantproteinuriainurinespecimensinhealthscreeningsettings.Springerplus.2016;5:1791.324.ChoMC,JiM,KimSY,etal.EvaluationoftheURiSCANsupercassetteACRsemiquantita

---

### Chunk 13/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.569

e-basedequations.TransplInt.2012;25:527–536.162.JanusN,Launay-VacherV,ByloosE,etal.CancerandrenalinsufciencyresultsoftheBIRMAstudy.BrJCancer.2010;103:1815–1821.163.Launay-VacherV,JanusN,DerayG.Renalinsufciencyandcancertreatments.ESMOOpen.2016;1:e000091.164.NaSY,SungJY,ChangJH,etal.Chronickidneydiseaseincancerpatients:anindependentpredictorofcancer-specicmortality.AmJNephrol.2011;33:121–130.165.RosnerMH,JhaveriKD,McMahonBA,etal.Onconephrology:the
intersectionsbetweenthekidneyandcancer.CACancerJClin.2021;71:47–77.166.SoveriI,BergUB,BjorkJ,etal.MeasuringGFR:asystematicreview.AmJKidneyDis.2014;64:411–424.167.WhiteCA,AkbariA,AllenC,etal.Simultaneousglomerularltrationratedeterminationusinginulin,iohexol,and99mTc-DTPAdemonstratestheneedforcustomizedmeasurementprotocols.KidneyInt.2021;99:957–966.168.XieP,HuangJM,LiuXM,etal.(99m)Tc-DTPArenaldynamicimaging
methodmaybeunsuitabletobeusedasthereferencemethodin
investigatingthevalidityofCDK-EPIequationfordetermining
glomerularltrati

---

### Chunk 14/30
**Article:** Automated urine sediment analyzers underestimate the severity of hematuria in glomerular diseases (2021)
**Journal:** Scientific Reports
**Section:** other | **Similarity:** 0.568

omparison of automated urinalysis systems and manual microscopy. Clin. Chim. Acta 384, 2834 (2007). 6. Cavanaugh, C. & Perazella, M. A. Urine sediment examination in the diagnosis and management of kidney disease: Core cur-riculum 2019. Am. J. Kidney Dis. 73, 258272 (2019). 7. Becker, G. J., Garigali, G. & Fogazzi, G. B. Advances in urine microscopy. Am. J. Kidney Dis. 67, 954964 (2016). 8. Oyaert, M. & Delanghe, J. Progress in automated urinalysis. Ann. Lab. Med. 39, 1522 (2019). 9. Lee, W., Ha, J. S. & Ryoo, N. H. Comparison of the automated cobas u 701 urine microscopy and UF-1000i ow cytometry systems and manual microscopy in the examination of urine sediments. J. Clin. Lab. Anal. 30, 663671 (2016). 10. Moreno, J. A. et al. Glomerular hematuria: Cause or consequence of renal inammation? Int. J. Mol. Sci. 20, 2205 (2019). 11. Kincaid-Smith, P. & Fairley, K. e investigation of hematuria. Semin. Nephrol. 25, 127135 (2005). 12.

---

### Chunk 15/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.563

inJapan.OpenForumInfectDis.2018;5:ofy216.381.YangCJ,ChenDP,WenYH,etal.Evaluationthediagnosticaccuracyof
albuminuriadetectioninsemi-quantitativeurinalysis.ClinChimActa.2020;510:177–180.382.KouriT,NokelainenP,PelkonenV,etal.EvaluationoftheARKRAYAUTIONElevenreectometerindetectingmicroalbuminuriawithAUTIONScreenteststripsandproteinuriawithAUTIONSticks10PAstrips.ScandJClinLabInvest.2008;69:52–64.383.NagrebetskyA,JinJ,StevensR,etal.Diagnosticaccuracyofurinedipsticktestinginscreeningformicroalbuminuriaintype2diabetes:acohortstudyinprimarycare.FamPract.2012;30:142–152.384.NahEH,ChoS,KimS,etal.Comparisonofurinealbumin-to-creatinineratio(ACR)betweenACRstriptestandquantitativetestinprediabetesanddiabetes.AnnLabMed.2016;37:28–33.385.ShinJI,ChangAR,GramsME,etal.Albuminuriatestinginhypertensionanddiabetes:anindividual-participantdatameta-analysisinaGlobalConsortium.Hypertension.2021;78:1042–1052.386.PantaloneKM,JiX,KongSX,etal.Unmetneedsandopportunitiesforoptimalmanagementofpatientswithty

---

### Chunk 16/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.560

rshallSM,AlbertiKG.Screeningforearlydiabeticnephropathy.AnnClinBiochem.1986;23(Pt2):195–197.290.MillerWG,BrunsDE,HortinGL,etal.Currentissuesinmeasurementandreportingofurinaryalbuminexcretion.ClinChem.2009;55:24–38.291.ChitaliaVC,KothariJ,WellsEJ,etal.Cost-benetanalysisandpredictionof24-hourproteinuriafromthespoturineprotein-creatinineratio.ClinNephrol.2001;55:436–447.292.CoteAM,BrownMA,LamE,etal.Diagnosticaccuracyofurinaryspotprotein:creatinineratioforproteinuriainhypertensivepregnant
women:systematicreview.BMJ.2008;336:1003–1006.293.DysonEH,WillEJ,DavisonAM,etal.Useoftheurinaryproteincreatinine
indextoassessproteinuriainrenaltransplantpatients.NephrolDialTransplant.1992;7:450–452.294.Leanos-MirandaA,Marquez-AcostaJ,Romero-ArauzF,etal.Protein:
creatinineratioinrandomurinesamplesisareliablemarkerof
increased24-hourproteinexcretioninhospitalizedwomenwith
hypertensivedisordersofpregnancy.ClinChem.2007;53:1623–1628.295.LemannJJr,DoumasBT.Proteinuriainhealthanddiseaseassessedby

---

### Chunk 17/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.559

ely(globalratewas514.9[474.9–558.9]).In2017,CKDindiabetesrepresentedathirdofallDALYs,andtherewere1.4million
Table1|Criteriaforchronickidneydisease(eitherofthefollowingpresentforaminimumof3months)Markersofkidneydamage(1ormore)Albuminuria(ACR$30mg/g[$3mg/mmol])Urinesedimentabnormalities
Persistenthematuria
Electrolyteandotherabnormalitiesdueto
tubulardisorders
Abnormalitiesdetectedbyhistology
Structuralabnormalitiesdetectedbyimaging
HistoryofkidneytransplantationDecreasedGFRGFR<60ml/minper1.73m2(GFRcategoriesG3a–G5)ACR,albumin-to-creatinineratio;GFR,glomerularltrationrate.
Table2|GFRcategoriesinCKDGFR
categoryGFR(ml/minper1.73m2)TermsG1$90NormalorhighG260–89MildlydecreasedaG3a45–59MildlytomoderatelydecreasedG3b30–44ModeratelytoseverelydecreasedG415–29SeverelydecreasedG5<15KidneyfailureCKD,chronickidneydisease;GFR,glomerularltrationrate.aRelativetotheyoungadultlevel.Intheabsenceofevidenceofkidneydamage,neitherG1norG2fulllsthecriteriaforCKD.

---

### Chunk 18/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.557

3.ScheiJ,StefanssonVT,MathisenUD,etal.Residualassociationsof
inammatorymarkerswitheGFRafteraccountingformeasuredGFRinacommunity-basedcohortwithoutCKD.ClinJAmSocNephrol.2016;11:280–286.184.SjostromP,TidmanM,JonesI.Determinationoftheproductionrateand
non-renalclearanceofcystatinCandestimationoftheglomerular
ltrationratefromtheserumconcentrationofcystatinCinhumans.ScandJClinLabInvest.2005;65:111–124.185.XinC,XieJ,FanH,etal.AssociationbetweenserumcystatinCand
thyroiddiseases:asystematicreviewandmeta-analysis.FrontEndocrinol(Lausanne).2021;12:766516.186.AgarwalR,BillsJE,YigazuPM,etal.Assessmentofiothalamateplasma
clearance:durationofstudyaffectsqualityofGFR.ClinJAmSocNephrol.2009;4:77–85.187.ShahKF,StevensPE,LambEJ.Theinuenceofacooked-shmealonestimatedglomerularltrationrate.AnnClinBiochem.2020;57:182–185.188.InkerLA,LeveyAS,CoreshJ.Estimatedglomerularltrationratefromapanelofltrationmarkers-hopeforincreasedaccuracybeyondmeasuredglomerularltrationrate?AdvChronicKidneyDis.2018;

---

### Chunk 19/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.556

e.JAmSocNephrol.2021;32:2994–3015.238.LeveyAS,StevensLA,SchmidCH,etal.Anewequationtoestimate
glomerularltrationrate.AnnInternMed.2009;150:604–612.239.PierceCB,MunozA,NgDK,etal.Age-andsex-dependentclinical
equationstoestimateglomerularltrationratesinchildrenandyoungadultswithchronickidneydisease.KidneyInt.2021;99:948–956.240.PottelH,BjörkJ,CourbebaisseM,etal.Developmentandvalidationofa
modiedfullagespectrumcreatinine-basedequationtoestimateglomerularltrationrate:across-sectionalanalysisofpooleddata.AnnInternMed.2021;174:183–191.241.NymanU,GrubbA,LarssonA,etal.TherevisedLund-MalmoGFR
estimatingequationoutperformsMDRDandCKD-EPIacrossGFR,age
andBMIintervalsinalargeSwedishpopulation.ClinChemLabMed.2014;52:815–824.242.LiuX,GanX,ChenJ,etal.AnewmodiedCKD-EPIequationforChinesepatientswithtype2diabetes.PLoSOne.2014;9:e109743.243.GrubbA,HorioM,HanssonLO,etal.GenerationofanewcystatinC-
basedestimatingequationforglomerularltrationratebyuseof7assaysstandardizedtotheinternationalcali

---

### Chunk 20/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.554

chmannLM.Movingtowardstandardizationofurinealbuminmeasurements.EJIFCC.2017;28:258–267.314.NationalInstituteofStandardsandTechnology.CerticationofStandardReferenceMaterial2925RecombinantHumanSerumAlbuminSolution(PrimaryReferenceCalibratorforUrineAlbumin)(Frozen).U.S.DepartmentofCommerce,NIST;2020.315.CarterJL,ParkerCT,StevensPE,etal.Biologicalvariationofplasmaand
urinarymarkersofacutekidneyinjuryinpatientswithchronickidneydisease.ClinChem.2016;62:876–883.316.NationalInstituteforHealthandCareExcellence.Point-of-carecreatininedevicestoassesskidneyfunctionbeforeCTimagingwithintravenous
contrast.NICEGuideline[NG37].NICE;2019.317.BatteA,MurphyKJ,NamazziR,etal.Evaluatingkidneyfunctionusingapoint-of-carecreatininetestinUgandanchildrenwithseveremalaria:aprospectivecohortstudy.BMCNephrol.2021;22:369.318.McTaggartMP,NewallRG,HirstJA,etal.Diagnosticaccuracyofpoint-of-
caretestsfordetectingalbuminuria:asystematicreviewandmeta-analysis.AnnInternMed.2014;160:550–557.319.AbitbolC,ZillerueloG,F

---

### Chunk 21/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.552

RifCKDisdeemedlikelyduetopresenceofotherclinicalindicators.1.1.4Evaluationofcause
PracticePoint1.1.4.1:EstablishthecauseofCKDusingclinicalcontext,personalandfamilyhistory,socialandenvi-ronmentalfactors,medications,physicalexamination,laboratorymeasures,imaging,andgeneticandpathologicdiagnosis(Figure8).
www.kidney-international.orgsummaryofrecommendationstatementsandpracticepointsKidneyInternational(2024)105(Suppl4S),S117–S314S149

PracticePoint1.1.4.2:Useteststoestablishacausebasedonresourcesavailable(Table622,98-100).Recommendation1.1.4.1:Wesuggestperformingakidneybiopsyasanacceptable,safe,diagnostictesttoevaluatecauseandguidetreatmentdecisionswhenclinicallyappropriate(2D).1.2EvaluationofGFR1.2.1OtherfunctionsofkidneysbesidesGFR
PracticePoint1.2.1.1:Usetheterm“GFR”whenreferringtothespecickidneyfunctionofglomerularltration.Usethemoregeneralterm“kidneyfunction(s)”whendealingwiththetotalityoffunctionsofthekidney.

---

### Chunk 22/30
**Article:** Automated urine sediment analyzers underestimate the severity of hematuria in glomerular diseases (2021)
**Journal:** Scientific Reports
**Section:** other | **Similarity:** 0.552

purpura;11, membranoproliferative glomerulonephritis; and 2, postinfectious glomerulonephritis). 

3
Scientific Reports |        (2021) 11:20981  | https://doi.org/10.1038/s41598-021-00457-6
e median age at diagnosis was 46 (3058) years. e control group consisted of 1087 patients who were diagnosed with bladder cancer (1064, transitional cell carcinoma; 7, urachal carcinoma; 7, neuroendocrine small cell carcinoma; 6, bladder invasion of prostate carcinoma;1, mucinous adenocarcinoma; 1, pleomorphic undierentiated sarcoma; and 1, invasive poorly dierentiated carcinoma). e median age at diagnosis was 69 (6075) years (p < 0.001 compared with glomerular disease). e number of urinalyses was 9 (413) among the patients with glomerular disease and 6 (49) among the patients with bladder cancer.Distributions of urine SG and pH. e SG of urine aects the morphology of RBCs and may thus alter the identication of RBCs via automated urine sediment analyzers.

---

### Chunk 23/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.549

172.LeveyAS,GreeneT,SchluchterMD,etal.Glomerularltrationratemeasurementsinclinicaltrials.ModicationofDietinRenalDiseaseStudyGroupandtheDiabetesControlandComplicationsTrialResearchGroup.JAmSocNephrol.1993;4:1159–1171.173.PerroneRD,SteinmanTI,BeckGJ,etal.Utilityofradioisotopicltrationmarkersinchronicrenalinsufciency:simultaneouscomparisonof125I-iothalamate,169Yb-DTPA,99mTc-DTPA,andinulin.TheModicationofDietinRenalDiseaseStudy.AmJKidneyDis.1990;16:224–235.174.InkerLA,LeveyAS.KnowingyourGFR-whenisthenumbernot(exactly)
thenumber?KidneyInt.2019;96:280–282.175.ShlipakMG,InkerLA,CoreshJ.SerumcystatinCforestimationofGFR.

---

### Chunk 24/30
**Article:** Automated urine sediment analyzers underestimate the severity of hematuria in glomerular diseases (2021)
**Journal:** Scientific Reports
**Section:** results | **Similarity:** 0.548

ability35. Second, the dipstick blood test may also give false positives and false negatives. Bacterial peroxidase, for example, can produce false-positive  results36. However, bacterial peroxidase does not seem to account for the lower numbers of urinary RBCs in glomerular disease because urinary tract infection was not documented in the patients with glomerular disease during the study period. In contrast, ascorbic acid or ascorbic acid-containing so drinks may cause false-negative  results37; however, the patients included in this study were not questioned about this before obtaining urinalysis specimens.In conclusion, our data suggest that urine RBC counting by the UF-1000i urine analyzer or the Cobas 6500 urine analyzer underestimates the severity of hematuria in glomerular disease, possibly because dysmorphic RBCs are susceptible to hemolysis and/or are not properly recognized.Received: 8 July 2021; Accepted: 7 October 2021
References 1. Roberts, J. R.

---

### Chunk 25/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.548

eneral,all3devicesdemonstratedacceptableaccuracyatlowerlevelsofeGFR(<30ml/minper1.73m2).316Resultsshowedthati-STATandABLdevicesmayhavehigherprobabilitiesofcorrectlyclassifyingpeopleinthesameeGFRcategoriesasthelaboratoryreferencethan
StatSensordevices.Foralbumin,theERTidentiedasystematicreviewpub-lishedin2014,byMcTaggartetal.,318thatevaluatedthediagnosticaccuracyofquantitativeandsemiquantitativeproteinoralbuminurinedipsticktestscomparedwithlaboratory-basedtestsamongpeoplewithsuspectedor
diagnosedCKD.TheERTincludedrelevantstudiesfrom
thisreviewandconductedanupdate.Sixty-vestudies(in66articles)319–344,345–368,369–384eval-uatedtheaccuracyofquantitativeandsemiquantitativepro-
teinoralbumindipsticktestsinageneralpopulationnoton
chapter1www.kidney-international.orgS194KidneyInternational(2024)105(Suppl4S),S117–S314

KRTorreceivingend-of-lifecare.Studiesaddressedthefollowingcriticaloutcomes:measurementbias(n¼1),analyticalvariability(n¼5),analyticalsensitivity(n¼2),andanalyticspeci

---

### Chunk 26/30
**Article:** Automated urine sediment analyzers underestimate the severity of hematuria in glomerular diseases (2021)
**Journal:** Scientific Reports
**Section:** other | **Similarity:** 0.548

according to the SG. In most cases, the RBC counts were graded signicantly lower in urine specimens with SG < 1.010 than in those with SG 1.0101.020. High urine SG (> 1.020) did not seem to alter the RBC counting via the UF-1000i urine 
Figure1.  Specic gravities (SG) of urine specimens measured using the UF-1000i urine analyzer ((A) bladder cancer; (B) glomerular disease) and the Cobas 6500 urine analyzer ((C) bladder cancer; (D) glomerular disease).

4
Scientific Reports |        (2021) 11:20981  | https://doi.org/10.1038/s41598-021-00457-6
analyzer, but tended to decrease the RBC counts in patients with glomerular disease when the urinalyses were performed using the Cobas 6500 urine analyzer (Supplementary TablesS1S4, Supplementary Figs.S1S4).Urine pH is another factor that aects the morphology of RBCs. In highly alkaline urine (pH > 9), RBCs may undergo lysis due to swelling of the outer membrane  layers24.

---

### Chunk 27/30
**Article:** Automated urine sediment analyzers underestimate the severity of hematuria in glomerular diseases (2021)
**Journal:** Scientific Reports
**Section:** methods | **Similarity:** 0.547

ge of Medicine, 88 Olympic-ro 43-gil, Songpa-gu, Seoul 05505, Republic of Korea. email: wsyang@amc.seoul.kr

2
Scientific Reports |        (2021) 11:20981  | https://doi.org/10.1038/s41598-021-00457-6
at each positive grade in a dipstick blood test in the urinalyses performed using two automated urine sediment analyzers (UF-1000i urine analyzer and Cobas 6500 urine analyzer).MethodsPatients. is study included adult patients (≥ 18years) who were diagnosed with nephritic glomerular disease via a kidney biopsy and diagnosed with bladder cancer via a bladder biopsy or transurethral resection of bladder tumor at Asan Medical Center (Seoul, Korea), a tertiary hospital; the patients medical records were retrospectively reviewed. Between January 2011 and November 10, 2016, urine sediment examinations were car-ried out using the UF-1000i urine analyzer (Sysmex Corporation). ereaer, the instrument for urinalysis was changed to the Cobas 6500 urine analyzer (Roche Diagnostics GmbH).

---

### Chunk 28/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.547

1;99:957–966.168.XieP,HuangJM,LiuXM,etal.(99m)Tc-DTPArenaldynamicimaging
methodmaybeunsuitabletobeusedasthereferencemethodin
investigatingthevalidityofCDK-EPIequationfordetermining
glomerularltrationrate.PLoSOne.2013;8:e62328.169.RoweC,SitchAJ,BarrattJ,etal.Biologicalvariationofmeasuredand
estimatedglomerularltrationrateinpatientswithchronickidneydisease.KidneyInt.2019;96:429–435.170.DelanayeP,EbertN,MelsomT,etal.Iohexolplasmaclearancefor
measuringglomerularltrationrateinclinicalpracticeandresearch:areview.Part1:howtomeasureglomerularltrationratewithiohexol?ClinKidneyJ.2016;9:682–699.171.KwongYT,StevensLA,SelvinE,etal.Imprecisionofurinaryiothalamate
clearanceasagold-standardmeasureofGFRdecreasesthediagnosticaccuracyofkidneyfunctionestimatingequations.AmJKidneyDis.2010;56:39–49.172.LeveyAS,GreeneT,SchluchterMD,etal.Glomerularltrationratemeasurementsinclinicaltrials.ModicationofDietinRenalDiseaseStudyGroupandtheDiabetesControlandComplicationsTrialResearchGroup.JAmSocNephrol

---

### Chunk 29/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.547

ta-analysis.ClinJAmSocNephrol.2022;17:1305–1315.193.vanEeghenSA,WiepjesCM,T’SjoenG,etal.CystatinC-basedeGFRchangesduringgender-afrminghormonetherapyintransgenderindividuals.ClinJAmSocNephrol.2023;18:1545–1554.194.PierreC,MarzinkeM,AhmedSB,etal.AACC/NKFguidancedocumenton
improvingequityinchronickidneydiseasecare.JApplLabMed.2023;8:789–816.195.NgDK,FurthSL,WaradyBA,etal.Self-reportedrace,serumcreatinine,
cystatinC,andGFRinchildrenandyoungadultswithpediatrickidney
referenceswww.kidney-international.orgS298KidneyInternational(2024)105(Suppl4S),S117–S314

diseases:areportfromtheChronicKidneyDiseaseinChildren(CKiD)study.AmJKidneyDis.2022;80:174–185.e1.196.Luna-ZaizarH,Virgen-MontelongoM,Cortez-AlvarezCR,etal.Invitro
interferencebyacetaminophen,aspirin,andmetamizoleinserum
measurementsofglucose,urea,andcreatinine.ClinBiochem.2015;48:538–541.197.GreenbergN,RobertsWL,BachmannLM,etal.Specicitycharacteristicsof7commercialcreatininemeasurementproceduresbyenzymaticand
Jaffemethodpri

---

### Chunk 30/30
**Article:** Automated urine sediment analyzers underestimate the severity of hematuria in glomerular diseases (2021)
**Journal:** Scientific Reports
**Section:** discussion | **Similarity:** 0.547

ria in adults: e American Urological Association best practice policyPart I: Denition, detection, prevalence, and etiology. Urology 57, 599603 (2001). 26. Yuste, C. et al. Pathogenesis of glomerular haematuria. World J. Nephrol. 4, 185195 (2015). 27. Bataille, A. et al. Evidence of dipstick superiority over urine microscopy analysis for detection of hematuria. BMC Res. Notes 9, 435 (2016). 28. Shayanfar, N., Tobler, U., von Eckardstein, A. & Bestmann, L. Automated urinalysis: First experiences and a comparison between the Iris iQ200 urine microscopy system, the Sysmex UF-100 ow cytometer and manual microscopic particle counting. Clin. Chem. Lab. Med. 45, 12511256 (2007). 29. Cobbaert, C. M. et al. Automated urinalysis combining physicochemical analysis, on-board centrifugation, and digital imaging in one system: A multicenter performance evaluation of the cobas 6500 urine work area. Pract. Lab. Med. 17, e00139 (2019). 30. aller, T. R. & Wang, L. P.

---

