# ScoreItem: Ureia

**ID:** `019bf31d-2ef0-7534-b623-80601070d80d`
**FullName:** Ureia (Exames - Laboratoriais)
**Unit:** mg/dL

**Preparation Metadata:**
- Quality Grade: **EXCELLENT**
- Total Chunks: 30 de 4 artigos
- Avg Similarity: 0.624

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-7534-b623-80601070d80d`.**

```json
{
  "score_item_id": "019bf31d-2ef0-7534-b623-80601070d80d",
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

**ScoreItem:** Ureia (Exames - Laboratoriais)
**Unidade:** mg/dL

**30 chunks de 4 artigos (avg similarity: 0.624)**

### Chunk 1/30
**Article:** Proteinuria (2023)
**Journal:** StatPearls Publishing
**Section:** abstract | **Similarity:** 0.660

Comprehensive clinical review of proteinuria covering definition, etiology, diagnosis, and management. Proteinuria serves as a marker of kidney damage assisting with diagnosis, prognosis, and therapy. Normal protein excretion is <150 mg/24 hours. Severity categories: nephritic range (150-3,500 mg/24h) and nephrotic range (>3,500 mg/24h). Urine dipstick provides semi-quantitative screening but has important limitations: false-positives with dehydration, alkaline urine, UTI; inability to detect positively charged proteins. UK CKD guidelines define significant proteinuria as UPCR >45 mg/mmol. Quantitative methods (spot UPCR preferred for convenience and reliability, or 24-hour urine collection for accuracy) are indicated when dipstick is positive or proteinuria >1 g/day is suspected. Early detection and quantification are crucial given associations with progressive renal disease and cardiovascular complications.

---

### Chunk 2/30
**Article:** Urea levels and cardiovascular disease in patients with chronic kidney disease (2022)
**Journal:** Nephrology Dialysis Transplantation
**Section:** discussion | **Similarity:** 0.658

]forpatientsinT2andT3,respectively(Figure5,Supplementarydata,TableS7).DISCUSSIONOuranalysisofalargeprospectivecohortofCKDpatientsnotreceivingmaintenancedialysisattendingnephrologyout-patientfacilitiesshowedthathigherserumurealevelswereassociatedwithahigherincidenceofadverseCVoutcomesandahigherall-causemortalityrate.OurndingsareinlinewiththecurrentbodyofepidemiologicalevidenceonkidneyfunctionbiomarkersandtheriskofCVD,andsuggestthatureashouldbetakenintoaccountwhenseekingtopredictandpreventCVdiseaseinpatientswithCKD.Severalserumbiomarkers[suchasserumcreatinineorurea,orbloodureanitrogen(BUN),whichreectsonlythenitrogencontentofurea],areroutinelyusedinclinicalsettingstoevaluatekidneyfunction.Ureaisthemainmetabolitederivedfromdietaryproteinsandtissueproteinturnover.Thecompoundisalmostexclusivelyexcretedbythekidneysintheurine,afterltrationintheglomerulusandacertaindegreeofreabsorptionfromtheltrate.Althoughseveralnonrenalfactorsaecttheserumureaconcentration[20],reducedurinaryeliminationofu

---

### Chunk 3/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.652

.Diagnosticaccuracyofpoint-of-caretestsfordetectingalbuminuria:a
systematicreviewandmeta-analysis.AnnIntMed.2014;160:550–557.318SoFtablesSupplementaryTableS5SearchdateJuly2022
Citationsscreened/includedstudies2184/65
SupplementaryFigureS5Chapter2RiskassessmentinpeoplewithCKDClinicalquestionArekidneyfailurepredictionequationsgoodpredictorsofprogression,kidneyfailure,orend-stagerenaldisease?PopulationAdults,children,andyoungpeoplewithCKDG1-G5PredictorKidneyfailureriskequations(e.g.,Tangriequation[KidneyFailureRiskEquation])OutcomesPrognosticperformance:Calibration(goodnessofmeasures,e.g.,R2,Brierscore,andHosmer-Lemeshowtest)Discrimination(e.g.,sensitivity/specicity;areaunderthecurve[AUC]fromreceiveroperatingcharacteristic[ROC]andareaunderthereceiveroperatingcharacteristiccurve[AUROC];C-statistic)StudydesignSystematicreview
ExistingsystematicreviewNationalInstituteforHealthandCareExcellence.Evidencereviewforthebestcombinationofmeasurestoidentifyincreasedriskofprogressioninadults,childr

---

### Chunk 4/30
**Article:** Urea levels and cardiovascular disease in patients with chronic kidney disease (2022)
**Journal:** Nephrology Dialysis Transplantation
**Section:** other | **Similarity:** 0.646

rosis2017;263:127–1369.KumaA,WangXH,KleinJDetal.Inhibitionofureatransporteramelioratesuremiccardiomyopathyinchronickidneydisease.FASEBJ2020;34:8296–830910.KalimS,KarumanchiSA,ThadhaniRIetal.Proteincarbamylationinkidneydisease:pathogenesisandclinicalimplications.AmJKidneyDis2014;64:793–80311.KalimS,BergAH,KarumanchiSAetal.Proteincarbamylationandchronickidneydiseaseprogressioninthechronicrenalinsuciencycohortstudy.NephrolDialTransplant2021;37:139–14712.MoriD,MatsuiI,ShimomuraAetal.Proteincarbamylationexacerbatesvascularcalcication.KidneyInt2018;94:72–9013.ApostolovEO,RayD,SavenkaAVetal.ChronicuremiastimulatesLDLcarbamylationandatherosclerosis.JAmSocNephrol2010;21:1852–185714.StengelB,MetzgerM,CombeCetal.Riskprole,qualityoflifeandcareofpatientswithmoderateandadvancedCKD:TheFrenchCKD-REINCohortStudy.NephrolDialTransplant2019;34:277–28615.LeveyAS,StevensLA,SchmidCHetal.Anewequationtoestimateglomerularltrationrate.AnnInternMed2009;150:60416.HicksKA,MahaeyKW,MehranRetal.201

---

### Chunk 5/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.644

ta-analysis.ClinJAmSocNephrol.2022;17:1305–1315.193.vanEeghenSA,WiepjesCM,T’SjoenG,etal.CystatinC-basedeGFRchangesduringgender-afrminghormonetherapyintransgenderindividuals.ClinJAmSocNephrol.2023;18:1545–1554.194.PierreC,MarzinkeM,AhmedSB,etal.AACC/NKFguidancedocumenton
improvingequityinchronickidneydiseasecare.JApplLabMed.2023;8:789–816.195.NgDK,FurthSL,WaradyBA,etal.Self-reportedrace,serumcreatinine,
cystatinC,andGFRinchildrenandyoungadultswithpediatrickidney
referenceswww.kidney-international.orgS298KidneyInternational(2024)105(Suppl4S),S117–S314

diseases:areportfromtheChronicKidneyDiseaseinChildren(CKiD)study.AmJKidneyDis.2022;80:174–185.e1.196.Luna-ZaizarH,Virgen-MontelongoM,Cortez-AlvarezCR,etal.Invitro
interferencebyacetaminophen,aspirin,andmetamizoleinserum
measurementsofglucose,urea,andcreatinine.ClinBiochem.2015;48:538–541.197.GreenbergN,RobertsWL,BachmannLM,etal.Specicitycharacteristicsof7commercialcreatininemeasurementproceduresbyenzymaticand
Jaffemethodpri

---

### Chunk 6/30
**Article:** Urea levels and cardiovascular disease in patients with chronic kidney disease (2022)
**Journal:** Nephrology Dialysis Transplantation
**Section:** other | **Similarity:** 0.642

T,GlorieuxG.Ureaandchronickidneydisease:thecomebackofthecentury?(inuraemiaresearch).NephrolDialTransplant2018;33:4–1222.ThompsonS,JamesM,WiebeNetal.Causeofdeathinpatientswithreducedkidneyfunction.JAmSocNephrol2015;26:2504–251123.GreggLP,HedayatiSS.ManagementoftraditionalcardiovascularriskfactorsinCKD:whatarethedata?AmJKidneyDis2018;72:728–744Ureaandcardiovasculardisease191

24.D’ApolitoM,DuX,ZongHetal.Urea-inducedROSgenerationcausesinsulinresistanceinmicewithchronicrenalfailure.JClinInvest2010;120:203–21325.KoppeL,NyamE,VivotKetal.Ureaimpairsβcellglycolysisandinsulinsecretioninchronickidneydisease.JClinInvest2016;126:3598–361226.D’ApolitoM,DuX,PisanelliDetal.Urea-inducedROScauseen-dothelialdysfunctioninchronicrenalfailure.Atherosclerosis2015;239:393–40027.KrenningG,DankersPYW,DrouvenJWetal.Endothelialprogenitorcelldysfunctioninpatientswithprogressivechronickidneydisease.AmJPhysiolRenalPhysiol2009;296:F1314–F132228.WernerN,NickenigG.Endothelialprogenitorcellsinhealtha

---

### Chunk 7/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** methods | **Similarity:** 0.631

torytodietaryintervention,orcognitiveimpairment
summaryofrecommendationstatementsandpracticepointswww.kidney-international.orgS168KidneyInternational(2024)105(Suppl4S),S117–S314

Chapter1:EvaluationofCKD1.1DetectionandevaluationofCKDBothdecreasedGFRandincreasedalbuminuriaorothermarkersofkidneydamageareoftensilentandnotapparentto
thepersonatriskofCKDorthehealthcareproviderunless
laboratorytestsareperformed.ThecauseofthedecreasedGFR
orincreasedalbuminuriamayalsonotbeapparent.InthedecadesincethepublicationofthepreviousKDIGOClinicalPracticeGuidelinefortheEvaluationandManagementof
ChronicKidneyDisease,1therehavebeensubstantialadvancesintreatmentforCKDofallcauses(Chapter3),targeted
therapiesforspeciccausesofCKD(e.g.,KDIGO2021ClinicalPracticeGuidelinefortheManagementof
GlomerularDiseases22),aswellasunderstandingofandmethodstodeterminetheetiologyofCKD.Alltogether,theseadvanceshavethepotentialtoslowandpossiblypreventtheprogressionofkidneydisease.Thus,inthissectionofChapter
1,weemphasizetheim

---

### Chunk 8/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** methods | **Similarity:** 0.630

ey-international.orgS148KidneyInternational(2024)105(Suppl4S),S117–S314

SummaryofrecommendationstatementsandpracticepointsChapter1:EvaluationofCKD1.1DetectionandevaluationofCKD1.1.1DetectionofCKDPracticePoint1.1.1.1:Testpeopleatriskforandwithchronickidneydisease(CKD)usingbothurinealbuminmea-surementandassessmentofglomerularltrationrate(GFR).PracticePoint1.1.1.2:Followingincidentaldetectionofelevatedurinaryalbumin-to-creatinineratio(ACR),hematuria,orlowestimatedGFR(eGFR),repeatteststoconrmpresenceofCKD.1.1.2MethodsforstagingofCKDRecommendation1.1.2.1:InadultsatriskforCKD,werecommendusingcreatinine-basedestimatedglomer-ularltrationrate(eGFRcr).IfcystatinCisavailable,theGFRcategoryshouldbeestimatedfromthecombinationofcreatinineandcystatinC(creatinineandcystatinC–basedestimatedglomerularltrationrate[eGFRcr-cys])(1B).1.1.3EvaluationofchronicityPracticePoint1.1.3.1:Proofofchronicity(durationofaminimumof3months)canbeestablishedby:(i)reviewofpastmeasurements/estimationsofGFR;(ii)revie

---

### Chunk 9/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.627

eneral,all3devicesdemonstratedacceptableaccuracyatlowerlevelsofeGFR(<30ml/minper1.73m2).316Resultsshowedthati-STATandABLdevicesmayhavehigherprobabilitiesofcorrectlyclassifyingpeopleinthesameeGFRcategoriesasthelaboratoryreferencethan
StatSensordevices.Foralbumin,theERTidentiedasystematicreviewpub-lishedin2014,byMcTaggartetal.,318thatevaluatedthediagnosticaccuracyofquantitativeandsemiquantitativeproteinoralbuminurinedipsticktestscomparedwithlaboratory-basedtestsamongpeoplewithsuspectedor
diagnosedCKD.TheERTincludedrelevantstudiesfrom
thisreviewandconductedanupdate.Sixty-vestudies(in66articles)319–344,345–368,369–384eval-uatedtheaccuracyofquantitativeandsemiquantitativepro-
teinoralbumindipsticktestsinageneralpopulationnoton
chapter1www.kidney-international.orgS194KidneyInternational(2024)105(Suppl4S),S117–S314

KRTorreceivingend-of-lifecare.Studiesaddressedthefollowingcriticaloutcomes:measurementbias(n¼1),analyticalvariability(n¼5),analyticalsensitivity(n¼2),andanalyticspeci

---

### Chunk 10/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.625

tswithtype2diabetes.PLoSOne.2014;9:e109743.243.GrubbA,HorioM,HanssonLO,etal.GenerationofanewcystatinC-
basedestimatingequationforglomerularltrationratebyuseof7assaysstandardizedtotheinternationalcalibrator.ClinChem.2014;60:974–986.244.ZhaoL,LiHL,LiuHJ,etal.ValidationoftheEKFCequationforglomerular
ltrationrateestimationandcomparisonwiththeAsian-modiedCKD-EPIequationinChinesechronickidneydiseasepatientsinanexternal
study.RenalFail.2023;45:2150217.
www.kidney-international.orgreferencesKidneyInternational(2024)105(Suppl4S),S117–S314S299

245.EneanyaND,AdingwupuOM,KostelanetzS,etal.SocialdeterminantsofhealthandtheirimpactontheBlackracecoefcientinserumcreatinine-basedestimationofGFR:secondaryanalysisofMDRDandCRICstudies.

---

### Chunk 11/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.625

e-basedequations.TransplInt.2012;25:527–536.162.JanusN,Launay-VacherV,ByloosE,etal.CancerandrenalinsufciencyresultsoftheBIRMAstudy.BrJCancer.2010;103:1815–1821.163.Launay-VacherV,JanusN,DerayG.Renalinsufciencyandcancertreatments.ESMOOpen.2016;1:e000091.164.NaSY,SungJY,ChangJH,etal.Chronickidneydiseaseincancerpatients:anindependentpredictorofcancer-specicmortality.AmJNephrol.2011;33:121–130.165.RosnerMH,JhaveriKD,McMahonBA,etal.Onconephrology:the
intersectionsbetweenthekidneyandcancer.CACancerJClin.2021;71:47–77.166.SoveriI,BergUB,BjorkJ,etal.MeasuringGFR:asystematicreview.AmJKidneyDis.2014;64:411–424.167.WhiteCA,AkbariA,AllenC,etal.Simultaneousglomerularltrationratedeterminationusinginulin,iohexol,and99mTc-DTPAdemonstratestheneedforcustomizedmeasurementprotocols.KidneyInt.2021;99:957–966.168.XieP,HuangJM,LiuXM,etal.(99m)Tc-DTPArenaldynamicimaging
methodmaybeunsuitabletobeusedasthereferencemethodin
investigatingthevalidityofCDK-EPIequationfordetermining
glomerularltrati

---

### Chunk 12/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.624

S(i)renin-angiotensin-aldosteronesystem(inhibitor)RBCredbloodcellRCTrandomizedcontrolledtrialRRrelativeriskSCrserumcreatinineSBPsystolicbloodpressureSESsocioeconomicstatusSGLT2isodium-glucosecotransporter-2inhibitor(s)T1DType1diabetesT2DType2diabetesUKUnitedKingdomUSUnitedStatesUSRDSUnitedStatesRenalDataSystem
WHOWorldHealthOrganization
abbreviationsandacronymswww.kidney-international.orgS128KidneyInternational(2024)105(Suppl4S),S117–S314

NoticeSECTIONI:USEOFTHECLINICALPRACTICEGUIDELINEThisClinicalPracticeGuidelinedocumentisbaseduponliteraturesearchesconductedfromJuly2022throughApril2023andupdatedinJuly2023.Itisdesignedtoassistdecision-making.Itisnotintendedtodeneastandardofcareandshouldnotbeinterpretedasprescribinganexclusivecourseofmanagement.Variationsinpracticewillinevitablyandappropriatelyoccur
whencliniciansconsidertheneedsofindividualpatients,availableresources,andlimitationsuniquetoaninstitutionortypeof
practice.Healthcareprovidersusingthestatementsinthisdocument(bothpracti

---

### Chunk 13/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.619

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

### Chunk 14/30
**Article:** Urea levels and cardiovascular disease in patients with chronic kidney disease (2022)
**Journal:** Nephrology Dialysis Transplantation
**Section:** other | **Similarity:** 0.619

dheartfailure.AmJMed2004;116:466–47333.KirtaneAJ,LederDM,WaikarSSetal.Serumbloodureanitrogenasanindependentmarkerofsubsequentmortalityamongpatientswithacutecoronarysyndromesandnormaltomildlyreducedglomerularltrationrates.JAmCollCardiol2005;45:1781–178634.CauthenCA,LipinskiMJ,AbbateAetal.Relationofbloodureanitrogentolong-termmortalityinpatientswithheartfailure.AmJCardiol2008;101:1643–164735.DiIorioBR,MarzoccoS,BellasiAetal.Nutritionaltherapyreducesproteincarbamylationthroughurealoweringinchronickidneydisease.NephrolDialTransplant2018;33:804–813Received:26.10.2021;Editorialdecision:2.2.2022192S.M.Lavilleetal.

---

### Chunk 15/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.618

RifCKDisdeemedlikelyduetopresenceofotherclinicalindicators.1.1.4Evaluationofcause
PracticePoint1.1.4.1:EstablishthecauseofCKDusingclinicalcontext,personalandfamilyhistory,socialandenvi-ronmentalfactors,medications,physicalexamination,laboratorymeasures,imaging,andgeneticandpathologicdiagnosis(Figure8).
www.kidney-international.orgsummaryofrecommendationstatementsandpracticepointsKidneyInternational(2024)105(Suppl4S),S117–S314S149

PracticePoint1.1.4.2:Useteststoestablishacausebasedonresourcesavailable(Table622,98-100).Recommendation1.1.4.1:Wesuggestperformingakidneybiopsyasanacceptable,safe,diagnostictesttoevaluatecauseandguidetreatmentdecisionswhenclinicallyappropriate(2D).1.2EvaluationofGFR1.2.1OtherfunctionsofkidneysbesidesGFR
PracticePoint1.2.1.1:Usetheterm“GFR”whenreferringtothespecickidneyfunctionofglomerularltration.Usethemoregeneralterm“kidneyfunction(s)”whendealingwiththetotalityoffunctionsofthekidney.

---

### Chunk 16/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.618

onrate(GFR)andalbuminuriainpeoplewithchronickidneydisease(CKD)S199Figure14.(a)Predictedriskofkidneyfailureand(b)‡40%declineinestimatedglomerularltrationrate(eGFR)bychronickidneydisease(CKD)eGFR(G1–G5)andalbumin-to-creatinineratio(ACR)(A1–A3)categoriesinOptumLabsDataWarehouseS201Figure15.Transitionfromanestimatedglomerularltrationrate(eGFR)-basedtoarisk-basedapproachtochronickidneydiseasecareS202Figure16.Comparisonofriskofchronickidneydisease(CKD)progression(5-yearprobabilityofestimatedglomerularltrationrate[eGFR]<60ml/minper1.73m2)versuskidneyfailureinadultswithCKDG1–G2calculatedfromtheriskequationavailableathttps://www.ckdpc.org/risk-models.htmlS205Figure17.Chronickidneydisease(CKD)treatmentandriskmodicationS206Figure18.Holisticapproachtochronickidneydisease(CKD)treatmentandriskmodicationS208Figure19.Proteinguidelineforadultswithchronickidneydiseasenottreatedwithdialysis
contentswww.kidney-international.orgS120KidneyInternational(2024)105(Suppl4S),S117–S314

S209Figure2

---

### Chunk 17/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.618

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

### Chunk 18/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.618

edwithmeasured
values.VerylowlevelsofSCroftenrepresentpoorhealthstatus,suchasfrailtyorsarcopenia,whichlimitstheproductionof
creatinine.Thisbiologicalfeatureofcreatinine(i.e.,relation
tomusclemass)haslimiteditsprognosticutilityandresults
inreducingtheriskassociationsforeGFRcr45–60ml/minper1.73m2andelevatingrisksforeGFRcr>110ml/minper1.73m2.TheselimitationsarenotobservedwhenriskisestimatedusingeGFRcr-cysorcystatinC–basedeGFR(eGFRcys)(Figure7).WhencomparingGFRestimatesusingthese2ltrationmarkers,riskgradientsareconsistentlystrongerformost
outcomesforeGFRcysincomparisonwitheGFRcr.Therefore,
forthepurposeofevaluatingtheassociationofeGFRwith
outcomes(i.e.,projectingprognosisforpeoplewithCKD),the
eGFRcysoreGFRcr-cyscanbeconsideredmoreaccurate.

---

### Chunk 19/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.617

okK,etal.Cross-sectionalandlongitudinal
performanceofcreatinine-andcystatinC-basedestimatingequations
relativetoexogenouslymeasuredglomerularltrationrateinHIV-positiveandHIV-negativepersons.JAcquirImmuneDecSyndr.2020;85:e58–e66.158.DelanayeP,CavalierE,MorelJ,etal.Detectionofdecreasedglomerularltrationrateinintensivecareunits:serumcystatinCversusserumcreatinine.BMCNephrol.2014;15:9.159.CarlierM,DumoulinA,JanssenA,etal.Comparisonofdifferentequationstoassessglomerularltrationincriticallyillpatients.IntensiveCareMed.2015;41:427–435.160.SanglaF,MartiPE,VerissimoT,etal.MeasuredandestimatedglomerularltrationrateintheICU:aprospectivestudy.CritCareMed.2020;48:e1232–e1241.161.WagnerD,KniepeissD,StieglerP,etal.TheassessmentofGFRafter
orthotopiclivertransplantationusingcystatinCandcreatinine-basedequations.TransplInt.2012;25:527–536.162.JanusN,Launay-VacherV,ByloosE,etal.CancerandrenalinsufciencyresultsoftheBIRMAstudy.BrJCancer.2010;103:1815–1821.163.Launay-VacherV,JanusN,DerayG.Re

---

### Chunk 20/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** introduction | **Similarity:** 0.617

VOLUME 105 | ISSUE 4S | APRIL 2024
www.kidney-international.org
KDIGO 2024 Clinical Practice Guideline for the 
 
Evaluation and Management of Chronic Kidney Disease
SUPPLEMENT TO

KDIGO2024CLINICALPRACTICEGUIDELINEFORTHEEVALUATIONANDMANAGEMENTOFCHRONICKIDNEYDISEASEKidneyInternational(2024)105(Suppl4S),S117–S314S117

KDIGO2024ClinicalPracticeGuidelinefortheEvaluationandManagementofChronicKidneyDiseaseS118Tables,gures,andsupplementarymaterialS124KDIGOExecutiveCommitteeS125Referencekeys
S126CKDnomenclature
S127Conversionfactors
S128Abbreviationsandacronyms
S129Notice
S130ForewordS131WorkGroupmembershipS133Abstract
S134Patientforeword
S135Introduction,qualifyingstatements,andkeyconcepts
S141Specialconsiderations
S144SummaryofrelativeandabsoluterisksrelevanttoCKDfrommeta-analysisoflargemultinationalpopulationstudiesintheCKDPrognosisConsortium(CKD-PC)S149SummaryofrecommendationstatementsandpracticepointsS169Chapter1:EvaluationofCKD
S196Chapter2:RiskassessmentinpeoplewithCKD
S205Chapter3:

---

### Chunk 21/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.615

national(2024)105(Suppl4S),S117–S314S201

referralwaittimes,particularlyforhigh-riskindividuals.Inotherclinicalsettingswithrelativelyscarceaccessto
nephrologycare,thesethresholdsshouldbeadjustedto
ensurethatwaittimesareacceptableforlocalstandards.421Discussionofriskshouldalsoconsidertheindividual
person,theircomorbidities,andtheirriskofdeathfrom
othercauses.PracticePoint2.2.2:A2-yearkidneyfailureriskof>10%canbeusedtodeterminethetimingofmultidisciplinarycareinadditiontoeGFR-basedcriteriaandotherclinical
considerations.PeoplewithCKDG4–G5aremorelikelytodevelopcon-currentcomplicationsofCKDincludinganemia,hyper-
kalemia,bonemineraldisorders,and/ormetabolicacidosis
andprotein-energywasting.Inaddition,theyremainathigh
riskforadverseeventsincludingAKI,emergencydepartment
visits,andhospitalizations.Assuch,inmanycountriesand
healthcaresettings,thesepeoplemaybeenrolledininter-
disciplinarycareclinicsorreceivecaremanagementresourcestoreducemorbidityandhealthcarecosts,andtoavoidun-planneddialys

---

### Chunk 22/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.612

tNZJMed.1992;22:334–337.338.GilbertRE,AkdenizA,JerumsG.Detectionofmicroalbuminuriaindiabeticpatientsbyurinarydipstick.DiabetesResClinPract.1997;35:57–60.339.GrazianiMS,GambaroG,MantovaniL,etal.Diagnosticaccuracyofa
reagentstripforassessingurinaryalbuminexcretioninthegeneralpopulation.NephrolDialTransplant.2008;24:1490–1494.340.GuyM,NewallR,BorzomatoJ,etal.Useofarst-lineurineprotein-to-creatinineratiostriptestonrandomurinestoruleoutproteinuriain
patientswithchronickidneydisease.NephrolDialTransplant.2008;24:1189–1193.341.GuyM,NewallR,BorzomatoJ,etal.Diagnosticaccuracyoftheurinary
albumin:creatinineratiodeterminedbytheCLINITEKMicroalbuminandDCA2000þfortherule-outofalbuminuriainchronickidneydisease.ClinChimActa.2008;399:54–58.342.HasslacherC,MullerP,SchlipfenbacherRL.ResultsofamulticentrestudyforthedeterminationofmicroalbuminuriawithMicral-Test.KlinischesLabor.1995;41:441–447.343.HodelNC,HamadA,ReitherK,etal.Comparisonoftwodifferent
semiquantitativeurinarydipsticktestswithal

---

### Chunk 23/30
**Article:** Urea levels and cardiovascular disease in patients with chronic kidney disease (2022)
**Journal:** Nephrology Dialysis Transplantation
**Section:** results | **Similarity:** 0.611

ntsintherstureatertileinourstudycanbeconsideredashavinganormallevel,thoseinthesecondtertilehaveaslightlyelevatedlevelandthoseinthethirdtertilehaveamoderatelyelevatedlevel.DuringthecourseoftheCKD,serumurealevelscaneasilyreachorexceed10timestheuppernormallimit—especiallywhenkidneyfailureoccurs[21].Untilrecently,ureawasconsideredtobeabiologicallyinertmarker.However,studiesofanimalmodelsofCKDhaveshownthattheaccumulationofureaistoxic[21].CKDisacommondiseaseanditsprevalencewillcontinuetoincreaseinthecomingyears[2].CVDisoneoftheleadingcontributorstomorbidityandmortalityinpatientswithCKD[22].ThelimitedeectivenessofmostconventionalriskfactormodicationstrategiesinpatientswithCKDmaysuggestthatvariousmetabolicpathwaysunderliethedevelopmentofCVDinthispopulation[23].TheresultsofanimalstudiessuggestthatseveraldirectandindirectpathophysiologicalmechanismsunderlietherelationshipbetweenurealevelsandCVadverseeventsinCKD[5].Uraemicconditionsleadtotheoverexpressionofproapoptoticgenesinanimaltissuesand

---

### Chunk 24/30
**Article:** hs-CRP/HDL-C can predict the risk of all-cause mortality in cardiovascular-kidney-metabolic syndrome stage 1-4 patients (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.611

)levelsserveassignicantpredictorsfortheriskandprognosisofcardiovasculardiseases(CVD)(4–6).Furthermore,thetriglyceride-glucose(TyG)indexandtheestimatedglucosedisposalrate(eGDR)havealsobeenshowntoincreasetheriskofcardiovasculardiseasesinpatientswithCKMsyndromestages0-3,subsequentlyelevatingtheirriskofmortality(7,8).Incontrast,high-densitylipoproteincholesterol(HDL-C)hasbeenconventionallyviewedasacontributorto
cholesterolefux,endothelialprotection,andantioxidantfunctions,withitslevelsnegativelycorrelatedtotheonsetofcardiovasculardisease(CVD)(9,10).AstudypublishedintheAmericanJournalofKidneyDiseasespresentsaparadox:Inchronickidneydisease(CKD),asignicantdeclineintheactivityofkeyhigh-densitylipoprotein(HDL)-relatedenzymessuchasparaoxonase1(PON1),nitricoxide(NO)synthase(NOS),andLCATresultsinareductionofHDL-C’sfunctionalactivity,subsequentlytransitioningitintoaproinammatorystate.ThisstatenotonlyacceleratestheoxidationofphospholipidsbutalsopromotestheaccumulationofserumamyloidAandC-reac

---

### Chunk 25/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.611

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

### Chunk 26/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.610

eyDiseases,volume73,issue2,InkerLA,GramsME,LeveyAS,etal.RelationshipofestimatedGFRandalbuminuriatoconcurrentlaboratoryabnormalities:anindividualparticipantdatameta-
analysisinaGlobalConsortium,pages206–217,Copyrightª2018,withpermissionfromtheNationalKidneyFoundation,Inc.541
chapter3www.kidney-international.orgS230KidneyInternational(2024)105(Suppl4S),S117–S314

withaneGFRof<60ml/minper1.73m2atbaseline(amongwhom71primaryoutcomesaccrued)werecomparedwiththe5181peoplewithaneGFRof$60ml/minper1.73m2(568outcomes).611Certaintyofevidence.Theoverallcertaintyoftheevidenceforuricacid–loweringtherapyamongpeoplewithCKDandhyperuricemiaisverylow(seeSupplementaryTableS11150,612–614).ThecriticaloutcomeofdelayingprogressionofCKDwasaddressedby7RCTs.150,612,615–619The2largestRCTswereconsideredtohavealowriskofbias.615,616Thecertaintyoftheevidencewasdowngradedforinconsistency
becausetherewassubstantialstatisticalheterogeneitydetectedinthemeta-analysis(I2¼50%)andtheestimatedRRsrangedfrom0.05to2.96

---

### Chunk 27/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.610

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

### Chunk 28/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.609

chmannLM.Movingtowardstandardizationofurinealbuminmeasurements.EJIFCC.2017;28:258–267.314.NationalInstituteofStandardsandTechnology.CerticationofStandardReferenceMaterial2925RecombinantHumanSerumAlbuminSolution(PrimaryReferenceCalibratorforUrineAlbumin)(Frozen).U.S.DepartmentofCommerce,NIST;2020.315.CarterJL,ParkerCT,StevensPE,etal.Biologicalvariationofplasmaand
urinarymarkersofacutekidneyinjuryinpatientswithchronickidneydisease.ClinChem.2016;62:876–883.316.NationalInstituteforHealthandCareExcellence.Point-of-carecreatininedevicestoassesskidneyfunctionbeforeCTimagingwithintravenous
contrast.NICEGuideline[NG37].NICE;2019.317.BatteA,MurphyKJ,NamazziR,etal.Evaluatingkidneyfunctionusingapoint-of-carecreatininetestinUgandanchildrenwithseveremalaria:aprospectivecohortstudy.BMCNephrol.2021;22:369.318.McTaggartMP,NewallRG,HirstJA,etal.Diagnosticaccuracyofpoint-of-
caretestsfordetectingalbuminuria:asystematicreviewandmeta-analysis.AnnInternMed.2014;160:550–557.319.AbitbolC,ZillerueloG,F

---

### Chunk 29/30
**Article:** Urea levels and cardiovascular disease in patients with chronic kidney disease (2022)
**Journal:** Nephrology Dialysis Transplantation
**Section:** discussion | **Similarity:** 0.608

pressure,historyofcardiovasculardisease,anaemia,serumalbumin,high-sensitivityC-reactiveprotein,historyofacutekidneyinjury,diureticprescription,protonpumpinhibitorprescriptionandprescriptionsofstatinsandantiplateletagentsatbaseline.4.1[3.6;4.6]per100PY.Theincidenceratewas3.5timeshigherinpatientsinT3thaninpatientsinT1(Supplementarydata,TableS2,Figure4).TheadjustedHR[95%CI]formajorCVeventsassociatedwiththebaselineserumurealevelwas1.13[0.77;1.67]forpatientsinT2and1.94[1.27;2.98]forpatientsinT3(Figure4,Supplementarydata,TableS6).DeathbeforeRRTOveramedianfollow-upperiodof4.8[3.3;5.1]yearsintheCKD-REINcohort,407patientsdiedbeforeRRT,leadingtoanincidencerateof4.0[3.6;4.3]per100person-years(Supplementarydata,TableS2,Figure5).TheadjustedHRs[95%CI]fordeathbeforeRRTwere1.31[0.97;1.76]and1.73[1.22;2.45]forpatientsinT2andT3,respectively(Figure5,Supplementarydata,TableS7).DISCUSSIONOuranalysisofalargeprospectivecohortofCKDpatientsnotreceivingmaintenancedialysisattendingnephrologyout-patientfacilities

---

### Chunk 30/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** methods | **Similarity:** 0.606

edmethodstoestimateglomerularltrationrate(eGFR),seeSection1.2.†MarkersofkidneydamageotherthanalbuminuriamayalsobeusedtodiagnoseCKD,butalbumin-to-creatinineratio(ACR)andGFRarestillrequiredtodeterminestageandestimateriskofprogression.Acutekidneydisease(AKD)isdenedbytheabnormalitiesofkidneyfunctionand/orstructurewithimplicationsforhealthandwithadurationof#3months.28TheorangeboxesindicateactionsinpeopleatriskforCKDandinwhomtestingshouldbeperformed.Theblueboxesindicatetestingsteps.ThegreenboxesindicatetheidenticationofCKDanditsstagesandtheinitiationoftreatment.ThepurpleboxindicatestheidenticationofAKD/acutekidneyinjury(AKI).PleasealsoseetheKidneyDisease:ImprovingGlobalOutcomes(KDIGO)ClinicalPracticeGuidelineforAcuteKidneyInjury.97
www.kidney-international.orgintroduction,qualifyingstatements,andkeyconceptsKidneyInternational(2024)105(Suppl4S),S117–S314S139

Importantly,slowingCKDprogressionatearlystagesshouldprovideeconomicbenetsandpreventthedevelop-mentofkidneyfailureandcardiovasc

---

