# ScoreItem: Ecodopplercardiograma - TAPSE

**ID:** `019bf31d-2ef0-7b94-8e72-dc35cd889b1d`
**FullName:** Ecodopplercardiograma - TAPSE (Exames - Imagem)
**Unit:** mm

**Preparation Metadata:**
- Quality Grade: **GOOD**
- Total Chunks: 30 de 15 artigos
- Avg Similarity: 0.526

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-7b94-8e72-dc35cd889b1d`.**

```json
{
  "score_item_id": "019bf31d-2ef0-7b94-8e72-dc35cd889b1d",
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

**ScoreItem:** Ecodopplercardiograma - TAPSE (Exames - Imagem)
**Unidade:** mm

**30 chunks de 15 artigos (avg similarity: 0.526)**

### Chunk 1/30
**Article:** Relationship of TAPSE Normalized by Right Ventricular Area With Pulmonary Compliance, Exercise Capacity, and Clinical Outcomes (2024)
**Journal:** Circulation: Heart Failure
**Section:** abstract | **Similarity:** 0.662

Study introducing normalized TAPSE values (TAPSE/RVA-D and TAPSE/RVA-S) for improved prognostic assessment. TAPSE/RVA-D <1.1 and TAPSE/RVA-S <1.5 predicted adverse cardiovascular outcomes, providing better discrimination than traditional TAPSE alone. Demonstrates relationship between normalized TAPSE and pulmonary compliance and exercise capacity in heart failure patients.

---

### Chunk 2/30
**Article:** Tricuspid Annular Displacement Predicts Survival in Pulmonary Hypertension (2006)
**Journal:** American Journal of Respiratory and Critical Care Medicine
**Section:** abstract | **Similarity:** 0.659

Foundational study establishing TAPSE as a powerful prognostic marker in pulmonary hypertension. Demonstrated that for every 1-mm decrease in TAPSE, the unadjusted risk of death increased by 17%. TAPSE <1.8 cm was associated with greater RV systolic dysfunction and worse outcomes. Established TAPSE as a simple, reproducible measure of RV function with important clinical implications.

---

### Chunk 3/30
**Article:** Guidelines for the Echocardiographic Assessment of the Right Heart in Adults: 2025 ASE Update (2025)
**Journal:** Journal of the American Society of Echocardiography
**Section:** abstract | **Similarity:** 0.657

Updated guidelines for right heart echocardiographic assessment introducing graded severity classification for TAPSE. Normal TAPSE values are ≥2.5 cm, with graded severity ranges allowing reporting as normal, mild, moderate, or severely reduced. Emphasizes multiparametric approach including TAPSE, S', FAC, 3D RVEF, and RV-PA coupling for comprehensive right ventricular function assessment.

---

### Chunk 4/30
**Article:** Tricuspid Annular Plane Systolic Excursion and Pulmonary Arterial Systolic Pressure Relationship in Heart Failure (2013)
**Journal:** Chest
**Section:** abstract | **Similarity:** 0.618

Landmark study demonstrating that TAPSE/PASP ratio improves prognostic resolution in heart failure patients. Shows that TAPSE powerfully reflects RV function and prognosis. Unlike LVEF, TAPSE was an independent predictor of outcome in chronic heart failure. Establishes the importance of RV-PA coupling assessment beyond simple TAPSE measurement.

---

### Chunk 5/30
**Article:** Recommendations for Cardiac Chamber Quantification by Echocardiography in Adults: An Update from the American Society of Echocardiography (2015)
**Journal:** Journal of the American Society of Echocardiography
**Section:** abstract | **Similarity:** 0.559

Updated guidelines for echocardiographic assessment. Normal LVEF ≥54% in women by Simpson method. Sex-specific reference ranges are essential for accurate diagnosis of systolic dysfunction.

---

### Chunk 6/30
**Article:** Testosterone deficiency independently predicts mortality in women with HFrEF: insights from the T.O.S.CA. registry (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.555

TDshowedlowerleveloftricuspidannularplaneexcursion(TAPSE)topulmonaryarterialsystolicpressurePASPratio(TAPSE/PASP)(0.53±0.22and0.68±0.32mm/mmHg
inTDandTD,respectively,P=0.008),whichisasurrogateparameterofrightventricular-arterialuncoupling.Furthermore,HFrEFwomenwithTDalsodisplayedmoreim-
pairedcardiopulmonaryperformanceasshownbylowervaluesofpeakoxygenconsumption(VO2peak)thanthosewithoutTD(18.89±4.92vs.21.31±4.92ml*kg1*min1P=0.03).WomenwithTDalsohadworserenalfunctionasshownbylowervaluesofeGFRascomparedwiththoseTD(54.27±23.02vs.75.55±36.76ml/min,P<0.001).ClinicaloutcomesAfterameanfollow-upof36months,atotalof54eventswererecorded,with18deathsand36cardiovascularhospi-talizations.NinewomenoftheTDgroup(30%)andninewomenoftheTDgroup(14%)diedduringfollow-up.Fif-teenwomenoftheTDgroup(50%)underwenthospitaliza-tionduetocardiovascularreasonswhereas21cardiovascularhospitalizationswererecordedintheTDgroup(32%).AsdepictedinFigure1A,TDwasassociatedwithasignicantdifferenceintheoccurrenceofthe

---

### Chunk 7/30
**Article:** Recommendations for the Evaluation of Left Ventricular Diastolic Function by Echocardiography: An Update From the American Society of Echocardiography (2025)
**Journal:** Journal of the American Society of Echocardiography
**Section:** abstract | **Similarity:** 0.543

Updated guidelines for the evaluation of left ventricular diastolic function incorporating LAVI as a key parameter in the assessment algorithm. LAVI is an integral part of routine evaluation of patients with dyspnea or heart failure, providing powerful predictive marker of LV diastolic dysfunction.

---

### Chunk 8/30
**Article:** Testosterone deficiency independently predicts mortality in women with HFrEF: insights from the T.O.S.CA. registry (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.526

N-terminalproB-typenatriureticpeptide;BSA,bodysurfacearea;SBP,systolicbloodpressure;DBP,diastolicbloodpressure;eGFR,estimatedglomelurarltrationrate;LAVi,leftatrialvolumeindex;e,earlydiastolictransmitralowvelocity;E′,earlydiastolicmitralannularvelocity;TAPSE,tricuspidannularplaneexcursion;PASP,pulmonaryarterialsystolicpressure;RVOT-AT,rightventricularoutowtractaccelerationtime;VO2peak,peakoxygenconsumption;NYHA,NewYorkHeartAssociation;AF,atrialbrillation;ACE-I,angiotensin-converting-enzymeinhibitors;ARBs,angiotensin-receptorblockers;MRA,mineralocorticoid-receptorantagonist;ICD,implantablecardioverter-debrillator;CRT,cardiacresynchronizationtherapy;IGF-1D,IGF-1deciency;DHEA-SD,DHEA-Sdeciency.

---

### Chunk 9/30
**Article:** Testosterone deficiency independently predicts mortality in women with HFrEF: insights from the T.O.S.CA. registry (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.521

±SDMean±SDMean±SDAge(years)66.18±10.2668.93±9.74564.89±10.3250.12NTproBNP(pg/ml)2597.05±4093.042794.47±3380.142504.52±4409.550.72
BSA(m2)1.70±0.161.68±0.11.71±0.190.34SBP(mmHg)123.29±17.82121.32±16.86124.24±18.370.53DBP(mmHg)75.46±12.3975.05±9.5175.65±13.650.85
eGFR(ml/min)68.61±34.2854.27±23.0275.55±36.760.001
Haemoglobin(g/dl)12.85±1.5312.50±1.2113.01±1.650.13Albumin(%)56.8±5.256.6±7.756.9±3.70.9Testosterone(ng/dl)48.80±61.8613.60±6.6362.55±68.16<0.001Ejectionfraction(%)33.81±6.4633.20±7.6634.10±5.850.53
LAVi(ml/m2)49.46±25.5755.51±30.2446.63±22.770.17e/E′14.68±9.9613.04±4.9512.53±6.400.16TAPSE(mm)18.82±4.5117.78±4.2619.31±4.560.12
PASP(mmHg)34.38±13.2537.69±13.7232.84±12.840.98
RVOT-AT(ms)103.27±36.47112.11±34.7497.15±37.730.35TAPSE/PASP(mm/mmHg)0.63±0.290.53±0.220.68±0.320.008Pericardialeffusion14(15.20%)7(23.30%)7(10.93%)0.11
VO2peak(ml*kg1*min1)(n=67)20.61±4.7018.89±4.9221.31±4.920.03NYHAclassII66(70.21%)21(70.00%)45(70.31

---

### Chunk 10/30
**Article:** Testosterone deficiency independently predicts mortality in women with HFrEF: insights from the T.O.S.CA. registry (2026)
**Journal:** Revista não identificada
**Section:** introduction | **Similarity:** 0.519


Getulio José Mattos Do Amaral Filho - getfilho@yahoo.com.br - CPF: 034.983.039-88

denedbythecapabilityoftherightventricletoincreaseitscontractilityinresponsetoincreasedafterload.23TAPSE/PASPisindeedavalidatedmarkerofrightventricu-
lar-pulmonaryarterialuncoupling24andisalsoitselfastrongandindependentpredictorofmortalityinHFrEF.25,26AstudyperformedbyVentetuoloetal.elegantlydemonstratedhowmenopauseandwaningtestosteronelevelshaveadramatic
impactonlunghaemodynamicsandrightventricularfunction.27Interestingly,whencombined,bothVO2peakaswellasTAPSE/PASPareabletopredictpoorclinical
outcomesinpatientswithHFrEF.28AnotherndingfromourstudyisthatHFrEFforwomenwithTDalsodisplayworserenalfunctionwhichisafurtherindependentpredictorof
mortalityinHFpatients.29Mountingevidencesuggestastrongimpactoftestosteroneonrenalfunction.30Indeed,testosteronereplacementhasbeenalsoproposedaspotentialtherapyinrenaltransplantrecipientsinordertopreservekidneyfunction.31Testosteronedeciencyinthecontextofm

---

### Chunk 11/30
**Article:** Testosterone deficiency independently predicts mortality in women with HFrEF: insights from the T.O.S.CA. registry (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.513

riaforhearttransplantation:A
10-yearupdate.JHeartLungTransplant.2016;35:1–23.22.HarmsCA.Doesgenderaffectpulmo-naryfunctionandexercisecapacity?
RespirPhysiolNeurobiol.2006;151:124–131.23.VonkNoordegraafA,WesterhofBE,WesterhofN.Therelationshipbetween
therightventricleanditsloadinpulmo-naryhypertension.JAmCollCardiol.2017;69:236–243.24.RichterMJ,YogeswaranA,Husain-SyedF,VadászI,RakoZ,MohajeraniE,GhofraniHA,NaeijeR,SeegerW,HerbergU,RiethA,TedfordRJ,
GrimmingerF,GallH,TelloK.Anovel
non-invasiveandechocardiography-derivedmethodforquanticationofrightventricularpressure-volumeloops.

---

### Chunk 12/30
**Article:** Aula Afonso Salgado - Sistema nervoso autônomo (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.507

or prognóstico.  
- **Baixa variabilidade** → baixa atividade parassimpática, maior carga alostática, pior prognóstico.

Ele introduz dois conceitos centrais:

1. **Alostase:**  
   - é a capacidade do organismo de mobilizar energia para enfrentar os estressores;  
   - na metáfora de Afonso, é o “combustível do carro”: sem alostase, o paciente não tem “gasolina” para reagir;  
   - a avaliação da VFC mede, na prática, o nível de alostase.

2. **Carga alostática:**  
   - é o desgaste acumulado ao longo do tempo decorrente do esforço crônico para manter a homeostase;  
   - conecta estresse crônico a doenças degenerativas e crônicas não transmissíveis;  
   - idosos, por exemplo, tendem a ter **baixa VFC** e alta carga alostática.

O protocolo ideal de avaliação descrito inclui:

- Exame realizado preferencialmente pela manhã, em jejum, evitando a ingestão de medicamentos naquele momento (pois alteram a leitura).

---

### Chunk 13/30
**Article:** Capnodynamics - noninvasive cardiac output and mixed venous oxygen saturation monitoring in children (2023)
**Journal:** Frontiers in Pediatrics
**Section:** results | **Similarity:** 0.507

esasintheclinicalstudy.Thistimetheinvasivegoldstandard,transpulmonaryarteryowprobe,wasusedasreference.Theagreementofabsolutevaluesshowedameanpercentageerrorof31%,indicatingacceptableagreement.Inaddition,thecapnodynamicmethodhada100%concordanceratewiththegold
standardindicatingitsabilitytodetectthesamechangesinCOasthereferencemethodin100%ofthecases.TheresultspointstowardsamorereliableperformanceofthecapnodynamicmethodinassessingCOandchangeinCOthansuprasternalDopplerperformedbyanexperiencedcardiacsonographer.Figure3belowshowsatimeplotofthecapnodynamicmethodagainstCOTS(A)intheexperimentalstudyandagainstsuprasternal2DDopplerintheclinicalstudy(B)inresponsetoPEEPexposureandatropine.Anadditionalclinicalvalidationstudyininfantshasbeenperformed,usingtransesophagealDoppler(CardioQ-ODM+®,KarlssonandLönnqvist10.3389/fped.2023.1111270FrontiersinPediatrics04frontiersin.org

DeltexMedicalLtd,Chichester,UnitedKingdom)asreference.15infants(medianageandweight8monthsand9kgrespectively,smallest3.7kg)we

---

### Chunk 14/30
**Article:** Exercise as medicine – evidence for prescribing exercise as therapy in 26 different chronic diseases (2026)
**Journal:** Revista não identificada
**Section:** introduction | **Similarity:** 0.506

ontraindicationsThefollowingcontraindicationsareinagreementwithaEuropeanWorkingGroup(Gianuzzi&
Tavazzi2001).�AcuteCHU(AMIorunstableangina),untilthe
conditionhasbeenstableforatleast5days�Dyspneaatrest�Pericarditis,myocarditis,endocarditis�Symptomaticaorticstenosis�Severehypertension.Thereisnoestablished,docu-
mentedborderlinebloodpressurevaluedeemedto
bethecut-offpointforincreasedrisk.Generallyit
isrecommendedthatdemandingphysicalexercise
beavoidedinthecaseofsystolicBP>180ordias-
tolicBP>105mmHg�Fever�Seriousnon-cardiacdiseaseHeartfailureBackgroundHeartfailureisaconditionwheretheheartisunabletopumpsufﬁcientlytomaintainbloodﬂowtomeetthemetabolicneedsoftheperipheraltissue(Braun-
wald&Libby,2008).Heartfailureisaclinicalsyn-
dromewithsymptomssuchasﬂuidretention,
breathlessness,orexcessivetirednesswhenrestingor
exercising,andwithobjectivesymptomsofreduced
systolicfunctionoftheleftventricleatrest.Asymptomaticleftventriculardysfunctionisoftentheprecursorofthissyndrome.Sympt

---

### Chunk 15/30
**Article:** Aula Afonso Salgado - Sistema nervoso autônomo (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.506

iação descrito inclui:

- Exame realizado preferencialmente pela manhã, em jejum, evitando a ingestão de medicamentos naquele momento (pois alteram a leitura).  
- Repetição do exame em **3 a 5 ocasiões** em condições semelhantes, para obter dados de “padrão ouro” (maior confiabilidade).  

A partir do ECG, softwares especializados analisam a VFC tanto no **domínio do tempo** quanto no **domínio da frequência**:

- No domínio do tempo, o parâmetro mais citado é o **SDNN** (desvio padrão dos intervalos NN), que é uma raiz quadrada aplicada à distribuição dos intervalos.  
- SDNN mais alto indica maior variabilidade; SDNN baixo indica rigidez do ritmo, associada a pior prognóstico.

No domínio da frequência, embora Afonso não detalhe numericamente, ele indica o uso de técnicas matemáticas como:

- **Rápida transformada de Fourier (FFT)**,  
- **wavelet transform**,  
- **ritmogramas** (conceito de origem russa).

---

### Chunk 16/30
**Article:** Testosterone deficiency independently predicts mortality in women with HFrEF: insights from the T.O.S.CA. registry (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.505

(mm/mmHg)0.63±0.290.53±0.220.68±0.320.008Pericardialeffusion14(15.20%)7(23.30%)7(10.93%)0.11
VO2peak(ml*kg1*min1)(n=67)20.61±4.7018.89±4.9221.31±4.920.03NYHAclassII66(70.21%)21(70.00%)45(70.31%)0.98III23(24.46%)7(23.33%)16(25.00%)0.87
IV5(5.33%)2(6.67%)3(4.69%)0.65Aetiology(ischaemic)27(28.7%)17(56.7%)10(15.6%)0.01Yearofdisease5[2–10.75]4[1–9.75]7[2.75–12]0.1AF11(22.4%)4(13.33%)7(10.90%)0.73
Beta-blockers66(91.7%)24(72.00%)42(65.60%)0.15
ACE-I31(43.1%)12(40.00%)19(29.70%)0.22ARBs26(36.1%)9(30.00%)17(26.60%)0.72MRA26(36.1%)8(26.70%)18(28.12%)0.88
ICD26(36.6%)8(26.66%)18(28.10%)0.88
CRT10(14.1%)1(3.33%)9(14.10%)0.15IGF-1D50(53.2%)16(53.3%)34(53.1%)0.98DHEA-SD73(77.7%)29(97.7%)44(68.7%)0.01
TypeIIdiabetes26(27.7%)6(20%)20(31.2)0.37Abbreviations:TD,testosteronedeciency;NTproBNP,N-terminalproB-typenatriureticpeptide;BSA,bodysurfacearea;SBP,systolicbloodpressure;DBP,diastolicbloodpressure;eGFR,estimatedglomelurarltrationrate;LAVi,leftatrialvolumeindex;e,earlydiastolictransmit

---

### Chunk 17/30
**Article:** Exercise as medicine – evidence for prescribing exercise as therapy in 26 different chronic diseases (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.505

ducesbetterorworseresultsthanaerobictraining
alone(Haykowskyetal.,2007).Basedontheabove,thefollowingisrecom-mended:�TrainingisrecommendedforallheartfailurepatientsinNYHAfunctionclassII–IIIonafullytitratedmedicationregimenandwellcompensatedfor3weeks.�Allpatientsshouldbeassessedbyacardiologist
beforeembarkingonatrainingprogram.�Forthesakeofcautionandinordertodetermine
individualexercisecapacity,thetrainingshouldbe
precededbyasymptom-limitedexercisetest.31Exerciseasmedicine–evidenceforprescribingexercise
	
	
Getulio José Mattos Do Amaral Filho - getfilho@yahoo.com.br - CPF: 034.983.039-88

�Asupervised,tailoredtrainingprogramisrecom-mendedafteraninitialexercisetest:Trainingprogramsforheartfailurepatientswithverylowexercisecapacitymustbestructuredwithshortdailysessionsoflow-intensityexercise,gradu-
allyincreasingdurationastheprogramprogresses.

---

### Chunk 18/30
**Article:** Monitoring of Tissue Oxygenation: an Everyday Clinical Challenge (2018)
**Journal:** Frontiers in Medicine
**Section:** other | **Similarity:** 0.503

ction imbalance that could be 
detected by low ScvO2 (24). In a clinical study aer myocardial infarction in patients with heart failure and cardiogenic shock, SvO2 was 43%, while in patients with heart failure without shock, it was 56% compared to patients without heart failure with an 
SvO2 of 70% (25). It may also be useful in patients with cardio-genic shock requiring the support by intraaortic balloon counter 
pulsation. In a study during weaning period, intraaortic balloon 

4
Molnar and Nemeth
Frontiers in Medicine | www.frontiersin.orgJanuary 2018 | Volume 4 | Article 247
pump assist ratio was decreased from 1:1 to 1:3. In the weaning failure group, decreased support was accompanied by a drop in ScvO2, while it remained constant in the successful group (26). In patients with chronic heart failure, ScvO2 can be chronically low.

---

### Chunk 19/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.500

herauthorshavenothing
todisclose.References1.McDonaghTA,MetraM,AdamoM,GardnerRS,BaumbachA,BöhmM,etal.;ESCScienticDocumentGroup.2021ESCGuidelinesforthediagnosisandtreatmentofacuteandchronicheartfailure:DevelopedbytheTaskForcefor
thediagnosisandtreatmentofacuteandchronicheartfailureoftheEuropean
SocietyofCardiology(ESC).WiththespecialcontributionoftheHeartFailure
Association(HFA)oftheESC.EurJHeartFail2022;24:4–131.https://doi.org/10.1002/ejhf.23332.Guidelinesforthediagnosisofheartfailure.TheTaskForceonHeartFailureoftheEuropeanSocietyofCardiology.EurHeartJ1995;16:741–751.3.HeidenreichPA,BozkurtB,AguilarD,AllenLA,ByunJJ,ColvinMM,etal.AHA/ACC/HFSAGuidelineforthemanagementofheartfailure:Executive
summary:AreportoftheAmericanCollegeofCardiology/AmericanHeart
AssociationJointCommitteeonClinicalPracticeGuidelines.JAmCollCardiol2022;79:1757–1780.https://doi.org/10.1016/j.jacc.2021.12.0114.Revuelta-LópezE,BarallatJ,CserkóováA,Gálvez-MontónC,JaffeAS,JanuzziJL,etal.Pre-analyticalconsid

---

### Chunk 20/30
**Article:** Echocardiographic Reference Ranges of Global Longitudinal Strain for All Cardiac Chambers Using Guideline-Directed Dedicated Views (2024)
**Journal:** JACC: Cardiovascular Imaging
**Section:** abstract | **Similarity:** 0.496

Estudo estabeleceu valores de referência para strain longitudinal global de todas câmaras cardíacas usando imagens obtidas conforme recomendações da ASE/EACVI. Incluiu 1.329 participantes saudáveis do estudo HUNT4Echo. Valores de referência para GLS do ventrículo esquerdo: -24% a -16%. Todos os strains associaram-se com idade, e GLS ventricular esquerdo associou-se com sexo. World Alliance Societies of Echocardiography propôs valores normais: -17% a -24% para homens e -18% a -26% para mulheres. Define-se GLS normal como ≤-18% e anormal como ≥-16%.

---

### Chunk 21/30
**Article:** Longitudinal strain by speckle tracking and echocardiographic parameters as predictors of adverse cardiovascular outcomes in chronic Chagas cardiomyopathy (2022)
**Journal:** International Journal of Cardiovascular Imaging
**Section:** abstract | **Similarity:** 0.495

Este estudo prospectivo com 177 pacientes examinou o valor prognóstico do strain longitudinal global (GLS) por speckle tracking em cardiomiopatia chagásica. Durante 42 meses, documentou-se desfechos adversos em 22,6% dos participantes. O GLS combinado com fração de ejeção e razão E/e' demonstrou valor prognóstico incremental (AUC 0,76 para 0,79). Pacientes com anormalidades nos três parâmetros apresentaram 60% de eventos adversos, comparado a 3,2% naqueles com valores normais, indicando que a avaliação ecocardiográfica multimodal melhora substancialmente a estratificação de risco.

---

### Chunk 22/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.494

congestiveheartfailureinanurgent-caresetting.JAmCollCardiol2001;37:379–385.https://doi.org/10.1016/s0735-1097(00)01156-619.CollinsSP,LindsellCJ,StorrowAB,AbrahamWT;ADHEREScienticAdvisoryCommittee,InvestigatorsandStudyGroup.Prevalenceofnegativechestradio-graphyresultsintheemergencydepartmentpatientwithdecompensatedheart
failure.AnnEmergMed2006;47:13–18.https://doi.org/10.1016/j.annemergmed.2005.04.00320.Bayés-GenísA,Santaló-BelM,Zapico-MuñizE,LópezL,CotesC,BellidoJ,etal.N-terminalprobrainnatriureticpeptide(NT-proBNP)intheemergencydiagnosisandin-hospitalmonitoringofpatientswithdyspnoeaandventriculardysfunction.EurJHeartFail2004;6:301–308.https://doi.org/10.1016/j.ejheart.2003.12.01321.LamLL,CameronPA,SchneiderHG,AbramsonMJ,MüllerC,KrumH.Meta-analysis:EffectofB-typenatriureticpeptidetestingonclinicaloutcomes
inpatientswithacutedyspneaintheemergencysetting.AnnInternMed2010;153:728–735.https://doi.org/10.7326/0003-4819-153-11-201012070-0000622.MoeGW,HowlettJ,JanuzziJL,ZowallH

---

### Chunk 23/30
**Article:** Testosterone deficiency independently predicts mortality in women with HFrEF: insights from the T.O.S.CA. registry (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.494

derwenthospitaliza-tionduetocardiovascularreasonswhereas21cardiovascularhospitalizationswererecordedintheTDgroup(32%).AsdepictedinFigure1A,TDwasassociatedwithasignicantdifferenceintheoccurrenceoftheprimaryendpoint
(unadjustedHR:2.31,95%CI1.20–4.10,P=0.004).Further-more,TDwasalsoassociatedwithhigherall-causemortality(unadjustedHR:2.87,95%CI1.11–7.40,P=0.03)(FigureS1A)andcardiovascularhospitalization(unadjustedHR:2.37,95%CI1.20–7.40,P=0.01)(FigureS1B).MultivariableanalysisTable1showsthatTDresultedanindependentpredictorofthecombinedendpointofall-causemortality/cardiovascular
hospitalization(adjusted,adj-HR:10.45,95%CI3.54–17.01,P=0.001).Sameresultwasfoundconsideringtheindividualcomponentsofthecompositeendpoint(adj-HR:8.33,95%CI
5.36–15.11P=0.039forall-causemortality;adj-HR:2.41,95%CI1.13–4.50;P=0.02forcardiovascularhospitalization)(TableS1).Similarresultswereobtainedafterexcluding(n=8)patientswithLVEFbetween40and45%(Table2).DiscussionThemainndingsofthecurrentstudyarethat(1)

---

### Chunk 24/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.491

sensus
documentalignwiththeNICEguidelinesonchronicheartfailure,
whichrecommendacut-offvalueofNT-proBNP>2000pg/ml.40InananalysisofprimarycaredatafromEngland,anNT-proBNPvalue
of>2000pg/mlwasassociatedwithamorethantwo-foldhigherriskofheartfailurehospitalizationand50%higherriskofmortalityascomparedwithanNT-proBNPof400–2000pg/ml.47Wesug-gestthat,irrespectiveofageandsex,patientswithanNT-proBNP
>2000pg/mlshouldbeprioritizedforechocardiographyandclini-calevaluationwithin2weeksofdiagnosis(Figure2).NT-proBNPinasymptomaticpatientswithriskfactors:heart
stressVariousriskfactors,suchashypertension,atheroscleroticcardio-vasculardisease,diabetes,obesity,andothers,contributetoanincreasedsusceptibilitytothedevelopmentofheartfailure.Intheabsenceofsymptomsofheartfailure,patientswithriskfactorsmay
exhibiteitherhearthealthorheartstress.Hearthealthrefersto
individualswhohaveastructurallynormalheartandnormalplasma
concentrationsofNPsandtroponins.©2023EuropeanSocietyofCardiology.

---

### Chunk 25/30
**Article:** Monitoring of Tissue Oxygenation: an Everyday Clinical Challenge (2018)
**Journal:** Frontiers in Medicine
**Section:** results | **Similarity:** 0.491

gan dysfunction 
syndrome, and thus improve the outcome of septic patients. In patients with early phase of severe sepsis, septic shock, early goal-
directed intervention guided by continuous monitoring of ScvO2, central venous pressure and mean arterial pressure (MAP), with target values of CVP 8 to 12mmHg, MAP>65mmHg and ScvO2>70%, reduced mortality from 46.5 to 30.5% at the 28th day (19).Although this study has been criticized for several reasons and these results could never be repeated, there is international 
consensus that that low ScvO2 values are very important warning signs of inadequate DO2 and can prognosticate complications and poor outcome. However, recent data suggest that high ScvO2 values may also have adverse outcomes in septic patients (20).

---

### Chunk 26/30
**Article:** Exercise as medicine – evidence for prescribing exercise as therapy in 26 different chronic diseases (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.491

ss,orexcessivetirednesswhenrestingor
exercising,andwithobjectivesymptomsofreduced
systolicfunctionoftheleftventricleatrest.Asymptomaticleftventriculardysfunctionisoftentheprecursorofthissyndrome.Symptomsvaryfromverylightrestrictionoffunctiontoseriouslydis-
ablingsymptoms.Heartfailureisoftendividedinto
left-sided(themostfrequentandbestresearched
type)andright-sidedheartfailureandintoacute
(pulmonaryedema,cardiogenicshock)andchronic
heartfailure.Heartfailureisoftencausedby
ischemicdisease,butcanalsobecausedby,forexample,hypertensionorvalvularheartdisease(Braunwald&Libby,2008).Maximumoxygenuptake(VO2max)islowerinpatientswithheartfailure(Sullivanetal.,1989b;
Cohen-Solaletal.,1990;2001c).Thisiscaused,
amongotherthings,bythereducedpumpingfunc-
tionoftheheartandbyperipheralconditionsinthe
muscles(Massieetal.,1988;Sullivanetal.,1990;2001c).Acommonsymptomwithheartfailurepatientsismuscleatrophy,tiredness,andreduced
musclestrength(Wilsonetal.,1993;Ankeretal.,
1997;Harringtonetal.,1997).Heartfailu

---

### Chunk 27/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.491

accausesofdyspnoea.ForrulinginheartfailureintheED,age-adjustedcut-pointshavebeenestablished:450pg/mlforpatientsunder50years,©2023EuropeanSocietyofCardiology.

---

### Chunk 28/30
**Article:** Aula Afonso Salgado - Sistema nervoso autônomo (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.485

bigeminia, trigeminia) etc.;  
  - também permite avaliar se o sistema reage positivamente a intervenções simples (como respiração profunda). Se não houver melhora com essas manobras, ele considera que não é o momento de prescrever exercícios respiratórios intensivos, sendo necessário começar por outras estratégias.

- **Card Check:**  
  - pode ser utilizado com sensor semelhante a oxímetro, inclusive em crianças;  
  - avalia seis funções fisiológicas principais:
    - oxigenação,  
    - ritmo cardíaco,  
    - flexibilidade vascular (índice de pulsatilidade),  
    - índice de resistividade do vaso,  
    - resistência temporal dos vasos e capacidade de reação (flexibilidade),  
    - reservas de energia nervosa e resposta a estressores;  
  - integra essas informações para estimar o estado psíquico–comportamental, via eixo coração–cérebro.

---

### Chunk 29/30
**Article:** Predictive value of electrocardiographic left ventricular hypertrophy in the general population: A meta-analysis (2020)
**Journal:** Journal of Electrocardiology
**Section:** abstract | **Similarity:** 0.484

This meta-analysis examined electrocardiographic left ventricular hypertrophy (LVH) as a predictor of adverse cardiac outcomes in 58,400 participants across 10 studies. The Sokolow-Lyon voltage, Cornell voltage, and Cornell product criteria showed comparable ability in predicting major adverse cardiovascular events (MACE), with risk ratios ranging from 1.56 to 1.70. The pooled risk ratio of MACE was 1.62 (95% CI 1.40-1.89) for Sokolow-Lyon voltage criteria. The pooled risk ratio of all-cause mortality was 1.47 (95% CI 1.10-1.97), and cardiovascular mortality was 1.38 (95% CI 1.19-1.60) for Sokolow-Lyon criteria. Cornell voltage demonstrated stronger predictive value for cardiovascular and all-cause mortality outcomes.

---

### Chunk 30/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.484

oninpri-marycare.EurHeartJ2010;31:1881–1889.https://doi.org/10.1093/eurheartj/ehq16346.ClelandJGF,PfefferMA,ClarkAL,JanuzziJL,McMurrayJJV,MuellerC,etal.Thestruggletowardsauniversaldenitionofheartfailure–howtoproceed?EurHeartJ2021;42:2331–2343.https://doi.org/10.1093/eurheartj/ehab08247.TaylorCJ,Lay-FlurrieSL,Ordóñez-MenaJM,GoyderCR,JonesNR,RoalfeAK,etal.NatriureticpeptidelevelatheartfailurediagnosisandriskofhospitalisationanddeathinEngland2004-2018.Heart2022;108:543–549.https://doi.org/10.1136/heartjnl-2021-31919648.NatriureticPeptidesStudiesCollaboration;WilleitP,KaptogeS,WelshP,Butter-worthAS,ChowdhuryR,SpackmanSA,etal.,Natriureticpeptidesandintegratedriskassessmentforcardiovasculardisease:Anindividual-participant-data©2023EuropeanSocietyofCardiology.

1898A.Bayes-Genisetal.

---

