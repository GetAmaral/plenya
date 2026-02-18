# ScoreItem: Proteínas (Qualitativo)

**ID:** `019bf31d-2ef0-73ae-b854-faed2b237b11`
**FullName:** Proteínas (Qualitativo) (Exames - Laboratoriais)
**Unit:** mg/dL

**Preparation Metadata:**
- Quality Grade: **EXCELLENT**
- Total Chunks: 30 de 4 artigos
- Avg Similarity: 0.648

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-73ae-b854-faed2b237b11`.**

```json
{
  "score_item_id": "019bf31d-2ef0-73ae-b854-faed2b237b11",
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

**ScoreItem:** Proteínas (Qualitativo) (Exames - Laboratoriais)
**Unidade:** mg/dL

**30 chunks de 4 artigos (avg similarity: 0.648)**

### Chunk 1/30
**Article:** Proteinuria (2023)
**Journal:** StatPearls Publishing
**Section:** abstract | **Similarity:** 0.737

Comprehensive clinical review of proteinuria covering definition, etiology, diagnosis, and management. Proteinuria serves as a marker of kidney damage assisting with diagnosis, prognosis, and therapy. Normal protein excretion is <150 mg/24 hours. Severity categories: nephritic range (150-3,500 mg/24h) and nephrotic range (>3,500 mg/24h). Urine dipstick provides semi-quantitative screening but has important limitations: false-positives with dehydration, alkaline urine, UTI; inability to detect positively charged proteins. UK CKD guidelines define significant proteinuria as UPCR >45 mg/mmol. Quantitative methods (spot UPCR preferred for convenience and reliability, or 24-hour urine collection for accuracy) are indicated when dipstick is positive or proteinuria >1 g/day is suspected. Early detection and quantification are crucial given associations with progressive renal disease and cardiovascular complications.

---

### Chunk 2/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.708

tNZJMed.1992;22:334–337.338.GilbertRE,AkdenizA,JerumsG.Detectionofmicroalbuminuriaindiabeticpatientsbyurinarydipstick.DiabetesResClinPract.1997;35:57–60.339.GrazianiMS,GambaroG,MantovaniL,etal.Diagnosticaccuracyofa
reagentstripforassessingurinaryalbuminexcretioninthegeneralpopulation.NephrolDialTransplant.2008;24:1490–1494.340.GuyM,NewallR,BorzomatoJ,etal.Useofarst-lineurineprotein-to-creatinineratiostriptestonrandomurinestoruleoutproteinuriain
patientswithchronickidneydisease.NephrolDialTransplant.2008;24:1189–1193.341.GuyM,NewallR,BorzomatoJ,etal.Diagnosticaccuracyoftheurinary
albumin:creatinineratiodeterminedbytheCLINITEKMicroalbuminandDCA2000þfortherule-outofalbuminuriainchronickidneydisease.ClinChimActa.2008;399:54–58.342.HasslacherC,MullerP,SchlipfenbacherRL.ResultsofamulticentrestudyforthedeterminationofmicroalbuminuriawithMicral-Test.KlinischesLabor.1995;41:441–447.343.HodelNC,HamadA,ReitherK,etal.Comparisonoftwodifferent
semiquantitativeurinarydipsticktestswithal

---

### Chunk 3/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.705

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

### Chunk 4/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.699

inmostCKDs.Useof
urinaryalbuminmeasurementasthepreferredtestforpro-teinuriadetectionwillimprovethesensitivity,quality,andconsistencyofapproachtotheearlydetectionandmanage-
mentofkidneydisease.Commonlyusedreagentstripdevicesmeasuringtotalproteinareinsufcientlysensitiveforthereliabledetectionofproteinuria,donotadjustforurinaryconcentration,andare
onlysemiquantitative.Furthermore,thereisnostandardization
betweenmanufacturers.Theuseofsuchstripsshouldbediscouragedinfavorofquantitativelaboratorymeasurementsofalbuminuriaorproteinuria,orvalidatedpoint-of-carede-
vicesforurinealbumin/ACR(Section1.4).Whenused,reagentstripresultsshouldbeconrmedbylaboratorytesting.Althoughthereferencepointremainstheaccuratelytimed24-hourspecimen,itiswidelyacceptedthatthisisadifcultproceduretocontroleffectivelyandthatinaccuraciesinurinary
collectionmaycontributetoerrorsinestimationofalbuminand/orproteinlosses.Inpractice,untimedurinesamplesareareasonablersttestforascertainmentofalbuminuria.Arstmorningvoidsample

---

### Chunk 5/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.686

22:369.318.McTaggartMP,NewallRG,HirstJA,etal.Diagnosticaccuracyofpoint-of-
caretestsfordetectingalbuminuria:asystematicreviewandmeta-analysis.AnnInternMed.2014;160:550–557.319.AbitbolC,ZillerueloG,FreundlichM,etal.Quantitationofproteinuriawithurinaryprotein/creatinineratiosandrandomtestingwithdipsticksinnephroticchildren.JPediatr.1990;116:243–247.320.AgardhCD.Anewsemiquantitativerapidtestforscreeningformicroalbuminuria.PracticalDiabetes.1993;10:146–147.321.AgarwalR,PanesarA,LewisRR.Dipstickproteinuria:canitguide
hypertensionmanagement?AmJKidneyDis.2002;39:1190–1195.322.AroraS,LongT,MenchineM.Testcharacteristicsofurinedipstickfor
identifyingrenalinsufciencyinpatientswithdiabetes.WestJEmergMed.2011;12:250–253.323.ChangCC,SuMJ,HoJL,etal.Theefcacyofsemi-quantitativeurineprotein-to-creatinine(P/C)ratioforthedetectionofsignicantproteinuriainurinespecimensinhealthscreeningsettings.Springerplus.2016;5:1791.324.ChoMC,JiM,KimSY,etal.EvaluationoftheURiSCANsupercassetteACRsemiquantita

---

### Chunk 6/30
**Article:** Methods for Diagnosing Proteinuria—When to Use Which Test and Why: A Review (2024)
**Journal:** American Journal of Kidney Diseases
**Section:** abstract | **Similarity:** 0.674

This comprehensive review examines diagnostic methods for proteinuria, comparing qualitative dipstick testing with quantitative methods. The study found overall concordance of 80.4% between urinalysis (UA) and albumin-to-creatinine ratio (ACR), with 17.2% false-negatives and 2.3% false-positives. Key limitations identified: dipstick tests primarily detect albumin and may miss non-albumin proteins (e.g., immunoglobulins in multiple myeloma). False-positives occur with highly alkaline urine, concentrated urine, gross hematuria, and antiseptic contamination. False-negatives occur in dilute urine. The review emphasizes dipstick reading must be interpreted considering urine concentration (specific gravity). A 1+ dipstick in dilute urine represents more severe proteinuria than the same reading in concentrated urine. Quantitative testing (UACR or UPCR) is required for any positive dipstick to confirm and quantify proteinuria.

---

### Chunk 7/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.669

stClin.2006;58:190–197.328.CroalBL,MutchWJ,ClarkBM,etal.Theclinicalapplicationofaurine
albumin:creatinineratiopoint-of-caredevice.ClinChimActa.2001;307:15–21.329.CurrinSD,GondweMS,MayindiNB,etal.Diagnosticaccuracyof
semiquantitativepointofcareurinealbumintocreatinineratioandurinedipstickanalysisinaprimarycareresourcelimitedsettingin
SouthAfrica.BMCNephrol.2021;22:103.330.DajakM,BonticA,UgnjatovicS,etal.[Evaluationofmethodsforrapidmicroalbuminuriascreeninginkidneydiseasedpatients].SrpArhCelokLek.2012;140:173–178[inSerbian].331.DavidsonMB,BazarganM,BakrisG,etal.ImmunoDip:animproved
screeningmethodformicroalbuminuria.AmJNephrol.2004;24:284–288.332.DavidsonMB,SmileyJF.Relationshipbetweendipstickpositiveproteinuriaandalbumin:creatinineratios.JDiabetesComplications.1999;13:52–55.333.deGrauwWJ,vandeLisdonkEH,vandeHoogenHJ,etal.Screeningfor
microalbuminuriaintype2diabeticpatients:theevaluationofadipsticktestingeneralpractice.DiabetMed.1995;12:657–663.334.FernándezFernándezI,Pá

---

### Chunk 8/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.668

albumindeterminationinpatientswithtype1diabetes.BrazJMedBiolRes.2002;35:337–343.347.KimY,ParkS,KimMH,etal.Canasemi-quantitativemethodreplacethe
currentquantitativemethodfortheannualscreeningofmicroalbuminuriainpatientswithdiabetes?Diagnosticaccuracyandcost-savinganalysisconsideringthepotentialhealthburden.PLoSOne.2020;15:e0227694.348.LeFlochJP,MarreM,RodierM,etal.InterestofClinitekMicroalbumininscreeningformicroalbuminuria:resultsofamulticentrestudyin302diabeticpatients.DiabetesMetab.2001;27:36–39.349.LeongSO,LuiKF,NgWY,etal.Theuseofsemi-quantitativeurinetest-
strip(MicralTest)formicroalbuminuriascreeninginpatientswithdiabetesmellitus.SingaporeMedJ.1998;39:101–103.350.LimD,LeeDY,ChoSH,etal.Diagnosticaccuracyofurinedipstickforproteinuriainolderoutpatients.KidneyResClinPract.2014;33:199–203.351.LimS,YuHJ,LeeS,etal.EvaluationoftheURiSCAN2ACRStripto
estimatetheurinealbumin/creatinineratios.JClinLabAnal.2017;32:e22289.352.LinCJ,ChenHH,PanCF,etal.Thecharacteristicsofnewsemi-
quantit

---

### Chunk 9/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.664

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
**Section:** discussion | **Similarity:** 0.662

inJapan.OpenForumInfectDis.2018;5:ofy216.381.YangCJ,ChenDP,WenYH,etal.Evaluationthediagnosticaccuracyof
albuminuriadetectioninsemi-quantitativeurinalysis.ClinChimActa.2020;510:177–180.382.KouriT,NokelainenP,PelkonenV,etal.EvaluationoftheARKRAYAUTIONElevenreectometerindetectingmicroalbuminuriawithAUTIONScreenteststripsandproteinuriawithAUTIONSticks10PAstrips.ScandJClinLabInvest.2008;69:52–64.383.NagrebetskyA,JinJ,StevensR,etal.Diagnosticaccuracyofurinedipsticktestinginscreeningformicroalbuminuriaintype2diabetes:acohortstudyinprimarycare.FamPract.2012;30:142–152.384.NahEH,ChoS,KimS,etal.Comparisonofurinealbumin-to-creatinineratio(ACR)betweenACRstriptestandquantitativetestinprediabetesanddiabetes.AnnLabMed.2016;37:28–33.385.ShinJI,ChangAR,GramsME,etal.Albuminuriatestinginhypertensionanddiabetes:anindividual-participantdatameta-analysisinaGlobalConsortium.Hypertension.2021;78:1042–1052.386.PantaloneKM,JiX,KongSX,etal.Unmetneedsandopportunitiesforoptimalmanagementofpatientswithty

---

### Chunk 11/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.660

rshallSM,AlbertiKG.Screeningforearlydiabeticnephropathy.AnnClinBiochem.1986;23(Pt2):195–197.290.MillerWG,BrunsDE,HortinGL,etal.Currentissuesinmeasurementandreportingofurinaryalbuminexcretion.ClinChem.2009;55:24–38.291.ChitaliaVC,KothariJ,WellsEJ,etal.Cost-benetanalysisandpredictionof24-hourproteinuriafromthespoturineprotein-creatinineratio.ClinNephrol.2001;55:436–447.292.CoteAM,BrownMA,LamE,etal.Diagnosticaccuracyofurinaryspotprotein:creatinineratioforproteinuriainhypertensivepregnant
women:systematicreview.BMJ.2008;336:1003–1006.293.DysonEH,WillEJ,DavisonAM,etal.Useoftheurinaryproteincreatinine
indextoassessproteinuriainrenaltransplantpatients.NephrolDialTransplant.1992;7:450–452.294.Leanos-MirandaA,Marquez-AcostaJ,Romero-ArauzF,etal.Protein:
creatinineratioinrandomurinesamplesisareliablemarkerof
increased24-hourproteinexcretioninhospitalizedwomenwith
hypertensivedisordersofpregnancy.ClinChem.2007;53:1623–1628.295.LemannJJr,DoumasBT.Proteinuriainhealthanddiseaseassessedby

---

### Chunk 12/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.655

.Diagnosticaccuracyofpoint-of-caretestsfordetectingalbuminuria:a
systematicreviewandmeta-analysis.AnnIntMed.2014;160:550–557.318SoFtablesSupplementaryTableS5SearchdateJuly2022
Citationsscreened/includedstudies2184/65
SupplementaryFigureS5Chapter2RiskassessmentinpeoplewithCKDClinicalquestionArekidneyfailurepredictionequationsgoodpredictorsofprogression,kidneyfailure,orend-stagerenaldisease?PopulationAdults,children,andyoungpeoplewithCKDG1-G5PredictorKidneyfailureriskequations(e.g.,Tangriequation[KidneyFailureRiskEquation])OutcomesPrognosticperformance:Calibration(goodnessofmeasures,e.g.,R2,Brierscore,andHosmer-Lemeshowtest)Discrimination(e.g.,sensitivity/specicity;areaunderthecurve[AUC]fromreceiveroperatingcharacteristic[ROC]andareaunderthereceiveroperatingcharacteristiccurve[AUROC];C-statistic)StudydesignSystematicreview
ExistingsystematicreviewNationalInstituteforHealthandCareExcellence.Evidencereviewforthebestcombinationofmeasurestoidentifyincreasedriskofprogressioninadults,childr

---

### Chunk 13/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.651

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

### Chunk 14/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.648

99–203.351.LimS,YuHJ,LeeS,etal.EvaluationoftheURiSCAN2ACRStripto
estimatetheurinealbumin/creatinineratios.JClinLabAnal.2017;32:e22289.352.LinCJ,ChenHH,PanCF,etal.Thecharacteristicsofnewsemi-
quantitativemethodfordiagnosingproteinuriabyusingrandomurine
samples.JClinLabAnal.2011;25:14–19.353.LloydMM,KuylJ,vanJaarsveldH.Evaluationofpoint-of-caretestsfor
detectingmicroalbuminuriaindiabeticpatients.SAfrFamPract.2011;53:281–286.354.MarshallSM,ShearingPA,AlbertiKG.Micral-teststripsevaluatedfor
screeningforalbuminuria.ClinChem.1992;38:588–591.355.MasimangoMI,HermansMP,MalembakaEB,etal.Impactofruralversus
urbansettingonkidneymarkers:across-sectionalstudyinSouth-Kivu,
DRCongo.BMCNephrol.2021;22:234.356.MasimangoMI,SumailiEK,JadoulM,etal.Prevalenceof
microalbuminuriaanddiagnosticvalueofdipstickproteinuriain
outpatientsfromHIVclinicsinBukavu,theDemocraticRepublicof
Congo.BMCNephrol.2014;15:146.357.McTaggartMP,PriceCP,PinnockRG,etal.Thediagnosticaccuracyofa
urinealbumin-creatinineratiopoint

---

### Chunk 15/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.642

creatinine(P/C)ratioforthedetectionofsignicantproteinuriainurinespecimensinhealthscreeningsettings.Springerplus.2016;5:1791.324.ChoMC,JiM,KimSY,etal.EvaluationoftheURiSCANsupercassetteACRsemiquantitativeurinedipstickformicroalbuminuriascreening.JClinLabAnal.2014;28:281–286.325.CollierG,GreenanMC,BradyJJ,etal.Astudyoftherelationship
betweenalbuminuria,proteinuriaandurinaryreagentstrips.AnnClinBiochem.2009;46:247–249.326.CollinsAC,VincentJ,NewallRG,etal.Anaidtotheearlydetectionand
managementofdiabeticnephropathy:assessmentofanewpointof
caremicroalbuminuriasysteminthediabeticclinic.DiabetMed.2001;18:928–932.327.Cortés-SanabriaL,Martínez-RamírezHR,HernándezJL,etal.Utilityofthe
DipstickMicraltestIIinthescreeningofmicroalbuminuriaofdiabetesmellitustype2andessentialhypertension.RevInvestClin.2006;58:190–197.328.CroalBL,MutchWJ,ClarkBM,etal.Theclinicalapplicationofaurine
albumin:creatinineratiopoint-of-caredevice.ClinChimActa.2001;307:15–21.329.CurrinSD,GondweMS,MayindiNB,etal.D

---

### Chunk 16/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** methods | **Similarity:** 0.639

nsPediatricconsiderations.PracticePoint1.2.4.3:EstimateGFRinchildrenusingvalidatedequationsthathavebeendevelopedorvalidatedincomparablepopulations.1.3Evaluationofalbuminuria1.3.1GuidanceforphysiciansandotherhealthcareprovidersPracticePoint1.3.1.1:Usethefollowingmeasurementsforinitialtestingofalbuminuria(indescendingorderofpref-erence).Inallcases,arstvoidinthemorningmidstreamsampleispreferredinadultsandchildren.(i)urineACR,or(ii)reagentstripurinalysisforalbuminandACRwithautomatedreading.Ifmeasuringurineprotein,usethefollowingmeasurements:(i)urineprotein-to-creatinineratio(PCR),(ii)reagentstripurinalysisfortotalproteinwithautomatedreading,or(iii)reagentstripurinalysisfortotalproteinwithmanualreading.PracticePoint1.3.1.2:Usemoreaccuratemethodswhenalbuminuriaisdetectedusinglessaccuratemethods.Conrmreagentstrippositivealbuminuriaand/orproteinuriabyquantitativelaboratorymea-surementandexpressasaratiotourinecreatininewhereverpossible(i.e.,quantifytheACRorPCRifinitialsemiquantitativetestsar

---

### Chunk 17/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** methods | **Similarity:** 0.636

JExpClinMed.2019;44:118–123.371.SalinasM,López-GarrigósM,FloresE,etal.Urinaryalbuminstripassay
asascreeningtesttoreplacequantitativetechnologyincertain
conditions.ClinChemLabMed.2018;57:204–209.372.SaradisPA,RiehleJ,BogojevicZ,etal.Acomparativeevaluationofvariousmethodsformicroalbuminuriascreening.AmJNephrol.2007;28:324–329.373.ShephardMD,BarrattLJ,Simpson-LyttleW.IstheBayerDCA2000
acceptableasascreeninginstrumentfortheearlydetectionofrenal
disease?AnnClinBiochem.1999;36(Pt3):393–394.374.SiednerMJ,GelberAC,RovinBH,etal.Diagnosticaccuracystudyofurinedipstickinrelationto24-hourmeasurementasascreeningtoolforproteinuriainlupusnephritis.JRheumatol.2007;35:84–90.375.SpoorenPF,LekkerkerkerJF,VermesI.Micral-test:aqualitativedipsticktestformicro-albuminuria.DiabetesResClinPract.1992;18:83–87.376.SzymanowiczA,Blanc-BernardE,RocheC,etal.EvaluationofMicralTestforthescreeningofthemicroalbuminuriainpointofcaretesting.Immuno-AnalyseBiologieSpecialisee.2008;23:109–115.377.TiuSC,LeeSS

---

### Chunk 18/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.632

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

### Chunk 19/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.632

tativelaboratorymeasurementand
expressasaratiotourinecreatininewhereverpossible
(i.e.,quantifytheACRorPCRifinitialsemiquantitative
testsarepositive).ConrmACR‡30mg/g(‡3mg/mmol)onarandomuntimedurinewithasubsequentrstmorningvoidinthemorningmidstreamurinesample.PracticePoint1.3.1.3:Understandfactorsthatmayaffect
interpretationofmeasurementsofurinealbuminand
urinecreatinineandorderconrmatorytestsasindicated(Table16).Thepracticepointadvocatingfortheuseofspotsamplesmeasuringalbuminorproteingreatlyfacilitatesitsincorpo-
rationintoclinicalpracticebyavoidingtheneedfortimed
urinecollections.Suchspotsamplescanover-orunderesti-
mateurinealbuminduetovariationindilution.UseofACR
orprotein-to-creatinineratio(PCR)inspoturinesamples
candecreasethiserror.ACRisanestimateoftotalurineal-
buminloss.Thecreatinineexcretionratevariessubstantiallybetweenpeople.ACRorPCRwilloverestimateurinealbuminlossinpeoplewithlowcreatinineexcretionandwillunder-
estimateurinealbuminorproteinlossinpeoplewithvery
highcrea

---

### Chunk 20/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.629

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

### Chunk 21/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** methods | **Similarity:** 0.624

ey-international.orgS148KidneyInternational(2024)105(Suppl4S),S117–S314

SummaryofrecommendationstatementsandpracticepointsChapter1:EvaluationofCKD1.1DetectionandevaluationofCKD1.1.1DetectionofCKDPracticePoint1.1.1.1:Testpeopleatriskforandwithchronickidneydisease(CKD)usingbothurinealbuminmea-surementandassessmentofglomerularltrationrate(GFR).PracticePoint1.1.1.2:Followingincidentaldetectionofelevatedurinaryalbumin-to-creatinineratio(ACR),hematuria,orlowestimatedGFR(eGFR),repeatteststoconrmpresenceofCKD.1.1.2MethodsforstagingofCKDRecommendation1.1.2.1:InadultsatriskforCKD,werecommendusingcreatinine-basedestimatedglomer-ularltrationrate(eGFRcr).IfcystatinCisavailable,theGFRcategoryshouldbeestimatedfromthecombinationofcreatinineandcystatinC(creatinineandcystatinC–basedestimatedglomerularltrationrate[eGFRcr-cys])(1B).1.1.3EvaluationofchronicityPracticePoint1.1.3.1:Proofofchronicity(durationofaminimumof3months)canbeestablishedby:(i)reviewofpastmeasurements/estimationsofGFR;(ii)revie

---

### Chunk 22/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.620

.Conrmreagentstrippositivealbuminuriaand/orproteinuriabyquantitativelaboratorymea-surementandexpressasaratiotourinecreatininewhereverpossible(i.e.,quantifytheACRorPCRifinitialsemiquantitativetestsarepositive).ConrmACR‡30mg/g(‡3mg/mmol)onarandomuntimedurinewithasubsequentrstmorningvoidinthemorningmidstreamurinesample.PracticePoint1.3.1.3:Understandfactorsthatmayaffectinterpretationofmeasurementsofurinealbuminandurinecreatinineandorderconrmatorytestsasindicated(Table16).

---

### Chunk 23/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.619

diovascularandrenaloutcomesindiabetes.JAmSocNephrol.2009;20:1813–1821.268.ShihabiZK,KonenJC,O’ConnorML.Albuminuriavsurinarytotalproteinfordetectingchronicrenaldisorders.ClinChem.1991;37:621–624.269.MartinH.Laboratorymeasurementofurinealbuminandurinetotalproteininscreeningforproteinuriainchronickidneydisease.ClinBiochemRev.2011;32:97–102.270.WaughJ,BellSC,KilbyM,etal.Effectofconcentrationandbiochemical
assayontheaccuracyofurinedipsticksinhypertensivepregnancies.

---

### Chunk 24/30
**Article:** Automated urine sediment analyzers underestimate the severity of hematuria in glomerular diseases (2021)
**Journal:** Scientific Reports
**Section:** results | **Similarity:** 0.616

and digital imaging in one system: A multicenter performance evaluation of the cobas 6500 urine work area. Pract. Lab. Med. 17, e00139 (2019). 30. aller, T. R. & Wang, L. P. Evaluation of asymptomatic microscopic hematuria in adults. Am. Fam. Phys. 60, 11431152 (1999). 31. Palsson, R., Srivastava, A. & Waikar, S. S. Performance of the automated urinalysis in diagnosis of proliferative glomerulonephritis. Kidney Int. Rep. 4, 723727 (2019). 32. Orlandi, P. F. et al. Hematuria as a risk factor for progression of chronic kidney disease and death: Findings from the chronic renal insuciency cohort (CRIC) study. BMC Nephrol. 19, 150 (2018). 33. Schulman, G. et al. Risk factors for progression of chronic kidney disease in the EPPIC trials and the eect of AST-120. Clin. Exp. Nephrol. 22, 299308 (2018). 34. Mashitani, T. et al.

---

### Chunk 25/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.615

ckproteinuriain
outpatientsfromHIVclinicsinBukavu,theDemocraticRepublicof
Congo.BMCNephrol.2014;15:146.357.McTaggartMP,PriceCP,PinnockRG,etal.Thediagnosticaccuracyofa
urinealbumin-creatinineratiopoint-of-caretestfordetectionof
albuminuriainprimarycare.AmJKidneyDis.2012;60:787–794.358.MeinhardtU,AmmannRA,FlückC,etal.Microalbuminuriaindiabetes
mellitus:efcacyofanewscreeningmethodincomparisonwithtimedovernighturinecollection.JDiabetesComplications.2003;17:254–257.359.MinettiEE,CozziMG,GranataS,etal.Accuracyoftheurinaryalbumin
titratorstick’Micral-Test’inkidney-diseasepatients.NephrolDialTransplant.1997;12:78–80.360.NaruseM,MukoyamaM,MorinagaJ,etal.Usefulnessofthequantitative
measurementofurineproteinatacommunity-basedhealthcheckup:a
cross-sectionalstudy.ClinExpNephrol.2019;24:45–52.361.OlivariusND,MogensenCE.Danishgeneralpractitioners’estimationofurinaryalbuminconcentrationinthedetectionofproteinuriaand
microalbuminuria.BrJGenPract.1995;45:71–73.362.OstaV,NatoliV,Diégue

---

### Chunk 26/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.613

kHM,Begin-HeickN,AcharyaC,etal.Automateddeterminationof
urineandcerebrospinaluidproteinswithCoomassiebrilliantblueandtheAbbottABA-100.ClinBiochem.1980;13:81–83.280.MarshallT,WilliamsKM.Totalproteindeterminationinurine:eliminationofadifferentialresponsebetweentheCoomassieblueandpyrogallolredproteindye-bindingassays.ClinChem.2000;46:392–398.281.GinsbergJM,ChangBS,MatareseRA,etal.Useofsinglevoidedurine
samplestoestimatequantitativeproteinuria.NEnglJMed.1983;309:1543–1546.282.PriceCP,NewallRG,BoydJC.Useofprotein:creatinineratio
measurementsonrandomurinesamplesforpredictionofsignicantproteinuria:asystematicreview.ClinChem.2005;51:1577–1586.283.BeethamR,CattellWR.Proteinuria:pathophysiology,signicanceandrecommendationsformeasurementinclinicalpractice.AnnClinBiochem.1993;30(Pt5):425–434.284.KeaneWF,EknoyanG.Proteinuria,albuminuria,risk,assessment,detection,elimination(PARADE):apositionpaperoftheNational
KidneyFoundation.AmJKidneyDis.1999;33:1004–1010.285.ClaudiT,CooperJG.Compar

---

### Chunk 27/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** methods | **Similarity:** 0.613

wingpage)
www.kidney-international.orgmethodsforguidelinedevelopmentKidneyInternational(2024)105(Suppl4S),S117–S314S275

Table44|(Continued)ClinicalquestionsandsystematicreviewtopicsinPICOSformatChapter1Evaluationofchronickidneydisease(CKD)Intervention(indextest)Machine-readquantitativeorsemiquantitativeproteinoralbuminurinedipsticktestsComparatorLaboratory-basedmethodsformeasuringurinaryproteinoralbumin(e.g.,24-hoururinarysample,spoturineACR,orPCR)OutcomesCriticaloutcomes:measurementbias,analyticalsensitivity(limitofdetection),analyticalvariability(coefcientofvariation),andanalyticspecicity(ornumberstocalculate)Otheroutcomes:probabilityofbeingclassiedineachalbuminuriaorproteinuriastageStudydesignCross-sectionalExistingsystematicreviewsforhandsearchingMcTaggartMP,NewallRG,HirstJA,etal.Diagnosticaccuracyofpoint-of-caretestsfordetectingalbuminuria:a
systematicreviewandmeta-analysis.AnnIntMed.2014;160:550–557.318SoFtablesSupplementaryTableS5SearchdateJuly2022
Citationsscreened/incl

---

### Chunk 28/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.611

:83–87.376.SzymanowiczA,Blanc-BernardE,RocheC,etal.EvaluationofMicralTestforthescreeningofthemicroalbuminuriainpointofcaretesting.Immuno-AnalyseBiologieSpecialisee.2008;23:109–115.377.TiuSC,LeeSS,ChengMW.Comparisonofsixcommercialtechniquesin
themeasurementofmicroalbuminuriaindiabeticpatients.DiabetesCare.1993;16:616–620.378.TsujikawaH,MachiiR,HiratsukaN,etal.[Evaluationofnovelteststripto
measurealbuminandcreatinineinurine].RinshoByori.2005;53:111–117[inJapanese].379.UsuiT,YoshidaY,NishiH,etal.Diagnosticaccuracyofurinedipstickfor
proteinuriacategoryinJapaneseworkers.ClinExpNephrol.2019;24:151–156.380.YanagisawaN,MuramatsuT,KoibuchiT,etal.Prevalenceofchronic
kidneydiseaseandpoordiagnosticaccuracyofdipstickproteinuriainhumanimmunodeciencyvirus-infectedindividuals:amulticenterstudyinJapan.OpenForumInfectDis.2018;5:ofy216.381.YangCJ,ChenDP,WenYH,etal.Evaluationthediagnosticaccuracyof
albuminuriadetectioninsemi-quantitativeurinalysis.ClinChimActa.2020;510:177–180.382.KouriT,Nok

---

### Chunk 29/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.610

inealbuminwhendiscussinggeneral
conceptsandwillrefereithertototalprotein,albumin,orother
specicproteinswhendiscussingthatparameterspecically.1.3.1GuidanceforphysiciansandotherhealthcareprovidersPracticePoint1.3.1.1:Usethefollowingmeasurementsforinitialtestingofalbuminuria(indescendingorderofprefer-ence).Inallcases,arstvoidinthemorningmidstreamsampleispreferredinadultsandchildren.(i)urineACR,or(ii)reagentstripurinalysisforalbuminandACRwith
automatedreading.Ifmeasuringurineprotein,usethefollowingmeasurements:(i)urineprotein-to-creatinineratio(PCR),(ii)reagentstripurinalysisfortotalproteinwithauto-
matedreading,or(iii)reagentstripurinalysisfortotalproteinwith
manualreading.

---

### Chunk 30/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.606

tegyforclinicalinvestigation:novel
usesofdataonbiologicalvariation.ClinChem.1987;33:2034–2038.262.BallantyneFC,GibbonsJ,O’ReillyDS.Urinealbuminshouldreplacetotalproteinfortheassessmentofglomerularproteinuria.AnnClinBiochem.1993;30(Pt1):101–103.263.LambEJ,MacKenzieF,StevensPE.Howshouldproteinuriabedetected
andmeasured?AnnClinBiochem.2009;46:205–217.264.NewmanDJ,ThakkarH,MedcalfEA,etal.Useofurinealbuminmeasurement
asareplacementfortotalprotein.ClinNephrol.1995;43:104–109.265.DawnayA,WilsonAG,LambE,etal.Microalbuminuriainsystemic
sclerosis.AnnRheumDis.1992;51:384–388.266.GrossJL,deAzevedoMJ,SilveiroSP,etal.Diabeticnephropathy:
diagnosis,prevention,andtreatment.DiabetesCare.2005;28:164–176.267.NinomiyaT,PerkovicV,deGalanBE,etal.Albuminuriaandkidney
functionindependentlypredictcardiovascularandrenaloutcomesindiabetes.JAmSocNephrol.2009;20:1813–1821.268.ShihabiZK,KonenJC,O’ConnorML.Albuminuriavsurinarytotalproteinfordetectingchronicrenaldisorders.ClinChem.1991;37:621–624.

---

