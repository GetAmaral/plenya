# ScoreItem: Nefrectomia

**ID:** `019bf31d-2ef0-78e0-b6a8-2323b3896449`
**FullName:** Nefrectomia (Histórico de doenças - Cirurgias já realizadas - Cirurgias que interferem diretamente no escore)

**Preparation Metadata:**
- Quality Grade: **GOOD**
- Total Chunks: 30 de 7 artigos
- Avg Similarity: 0.541

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-78e0-b6a8-2323b3896449`.**

```json
{
  "score_item_id": "019bf31d-2ef0-78e0-b6a8-2323b3896449",
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

**ScoreItem:** Nefrectomia (Histórico de doenças - Cirurgias já realizadas - Cirurgias que interferem diretamente no escore)

**30 chunks de 7 artigos (avg similarity: 0.541)**

### Chunk 1/30
**Article:** Clinical Practice Guideline on Follow-up Care for Living Kidney Donors: A Consensus Report from the KDIGO Conference (2017)
**Journal:** Transplantation
**Section:** abstract | **Similarity:** 0.659

KDIGO consensus guidelines for post-nephrectomy monitoring recommend annual assessments including: eGFR calculation, urine albumin-to-creatinine ratio, blood pressure measurement, and screening for diabetes and metabolic syndrome. Baseline assessment at 3 months post-surgery establishes new eGFR baseline. Patients should receive counseling on maintaining healthy BMI, avoiding NSAIDs and nephrotoxins, adequate hydration, and limiting protein intake to 0.8-1.0 g/kg/day. More intensive monitoring required if eGFR <60 mL/min/1.73m², proteinuria develops, or hypertension emerges.

---

### Chunk 2/30
**Article:** Metabolic Consequences of Uninephrectomy: A Comprehensive Review (2018)
**Journal:** Nephrology Dialysis Transplantation
**Section:** abstract | **Similarity:** 0.642

Comprehensive review of metabolic changes following nephrectomy. Unilateral nephrectomy leads to compensatory hyperfiltration in remaining kidney, with GFR typically reaching 60-70% of baseline. Long-term complications include accelerated nephron senescence, increased risk of albuminuria (25-30% at 10 years), metabolic acidosis, and impaired phosphate handling. Hypertension incidence doubles compared to general population. Guidelines recommend annual monitoring of eGFR, urinary albumin, blood pressure, and metabolic parameters. Lifestyle modifications including protein restriction and nephrotoxin avoidance are crucial.

---

### Chunk 3/30
**Article:** Cardiovascular and Renal Outcomes After Nephrectomy: A Population-Based Cohort Study (2016)
**Journal:** Circulation
**Section:** abstract | **Similarity:** 0.570

Population-based cohort study of 3,698 living kidney donors and matched controls followed for median 7.6 years. Nephrectomy was associated with increased risks of end-stage renal disease (HR 11.4), cardiovascular events (HR 1.4), and mortality (HR 1.3). Risk of hypertension increased by 30% compared to controls. Proteinuria developed in 15% of donors. These findings emphasize importance of long-term monitoring for cardiovascular and renal complications in individuals with solitary kidney.

---

### Chunk 4/30
**Article:** Long-term Renal Function Following Nephrectomy for Renal Cell Carcinoma: A Systematic Review and Meta-analysis (2020)
**Journal:** Journal of Clinical Medicine
**Section:** abstract | **Similarity:** 0.567

This systematic review examines long-term renal function outcomes following nephrectomy. Analysis of 45 studies involving over 15,000 patients showed that partial nephrectomy preserves better renal function compared to radical nephrectomy. Post-nephrectomy CKD incidence ranges from 20-40%, with risk factors including age, baseline eGFR, diabetes, and hypertension. The compensatory hypertrophy of the remaining kidney typically occurs within 6-12 months, but does not fully compensate for lost nephron mass. Annual eGFR decline is accelerated compared to general population.

---

### Chunk 5/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.567

e-basedequations.TransplInt.2012;25:527–536.162.JanusN,Launay-VacherV,ByloosE,etal.CancerandrenalinsufciencyresultsoftheBIRMAstudy.BrJCancer.2010;103:1815–1821.163.Launay-VacherV,JanusN,DerayG.Renalinsufciencyandcancertreatments.ESMOOpen.2016;1:e000091.164.NaSY,SungJY,ChangJH,etal.Chronickidneydiseaseincancerpatients:anindependentpredictorofcancer-specicmortality.AmJNephrol.2011;33:121–130.165.RosnerMH,JhaveriKD,McMahonBA,etal.Onconephrology:the
intersectionsbetweenthekidneyandcancer.CACancerJClin.2021;71:47–77.166.SoveriI,BergUB,BjorkJ,etal.MeasuringGFR:asystematicreview.AmJKidneyDis.2014;64:411–424.167.WhiteCA,AkbariA,AllenC,etal.Simultaneousglomerularltrationratedeterminationusinginulin,iohexol,and99mTc-DTPAdemonstratestheneedforcustomizedmeasurementprotocols.KidneyInt.2021;99:957–966.168.XieP,HuangJM,LiuXM,etal.(99m)Tc-DTPArenaldynamicimaging
methodmaybeunsuitabletobeusedasthereferencemethodin
investigatingthevalidityofCDK-EPIequationfordetermining
glomerularltrati

---

### Chunk 6/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.552

eeffectsareshowninTable32.735,736ThereisconsistentevidencethatwithholdingRASiisassociatedwithlowerriskofperioperativehypotensionin
varioustypesofsurgeryandprocedures(noncardiacsurgery,
cardiacsurgery,andcoronaryangiography).737,785,786TheevidencethatwithholdingRASiwouldlowerperioperativeAKIislessconsistentasaffectedbyfewerstudieswithlowsamplesizes.787,788Inthesurgicalcontext,antihyperglycemicagentssuchassulfonylureas,metformin,
andSGLT2iwouldbeheldbecauseoffastingbeforethe
surgery.Casereports,caseseries,andasystematicreviewof
47cases739,740,789supportthecurrentrecommendationsthatSGLT2ishouldbewithheldatleast3–4daysbeforetheelectivesurgery.741,742Temporarydiscontinuationofmedicationstomanageadverseeventsisindicatedinmostcases.However,fearfor
adverseeventrecurrenceoftenresultsinfailuretoresume
treatments.InCKD,hyperkalemiaorAKIarenotuncommon
adverseeffectsofRASitreatment,towhichclinicalguidelines
recommenddiscontinuationofRASiandtherapyreinitiation
atlowdosageswhentheeventisresolved.23

---

### Chunk 7/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.550

tsintheprocessesofdrugstewardshipinpeople
withCKD.Itisbeyondthescopeofthisguidelinetolistall
themedicationsthatmayhavealteredrisks/benetsinpeoplewithCKD.Suchinformationiswidelyavailablein
documentsthatmayexistatlocal,regional,ornationalbodies(e.g.,BritishNationalFormulary:www.bnf.org)andintextbooksofpharmacology.However,wedescribecaseexamplestohighlightthekeyclassesofcommonlypre-
scribedmedicationsinpeoplewithCKD.Thisguidanceis
basedonknowledgeofpharmacologythathasuniversal
relevance.Inmanycases,knowledgeofalteredrisks/benetsofmedicationscomes,however,fromobservationalstudies
andcasereportsfromroutinecare.4.1MedicationchoicesandmonitoringforsafetyAbnormalkidneyfunctionresultsinalterationinpharma-cokineticsandpharmacodynamics,andforpeoplewithCKD,astheGFRworsens,sodoestheprevalenceofpolypharmacy
andcomorbidities.725PeoplewithCKDareatincreasedriskofmedicationerrorsandinappropriateprescribing(noted
tobeupto37%inambulatoryoutpatientstudiesandup
to43%inlong-termcarestudies726,727).Thus,imp

---

### Chunk 8/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.546

studies,adoublingoftheACRwithina2-yeardurationisassociatedwithanincreasein
theriskofprogressiontokidneyfailureby50%–100%.391,392
Albuminuria categoriesDescription and rangeGFR categories (ml/min/1.73 m2)Description and rangeA1G1≥90G260–89G3a45–59G3b30–44G415–29G5<15Kidney failureSeverely decreasedModerately toseverely decreasedMildly tomoderately decreasedMildly decreasedNormal or highA2A3Normal to mildlyincreasedModeratelyincreasedSeverelyincreased<30 mg/g<3 mg/mmolScreen1Screen1Treat1Treat1Treat3Treat3Treat1Treat2Treat2Treat3Treat3Treat3Treat*3Treat4+Treat*3Treat4+Treat4+Treat4+30–299 mg/g3–29 mg/mmol≥300 mg/g≥30 mg/mmolCKD is classified based on:�Cause (C)�GFR (G)� Albuminuria (A)Low risk (if no other markers of kidney disease, no CKD)Moderately increased riskHigh risk
Very high risk
Figure13|Frequencyofmonitoringglomerularltrationrate(GFR)andalbuminuriainpeoplewithchronickidneydisease(CKD).AlbuminuriaandGFRgridreectstheriskofprogressionbyintensityofcolor

---

### Chunk 9/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.540

RifCKDisdeemedlikelyduetopresenceofotherclinicalindicators.1.1.4Evaluationofcause
PracticePoint1.1.4.1:EstablishthecauseofCKDusingclinicalcontext,personalandfamilyhistory,socialandenvi-ronmentalfactors,medications,physicalexamination,laboratorymeasures,imaging,andgeneticandpathologicdiagnosis(Figure8).
www.kidney-international.orgsummaryofrecommendationstatementsandpracticepointsKidneyInternational(2024)105(Suppl4S),S117–S314S149

PracticePoint1.1.4.2:Useteststoestablishacausebasedonresourcesavailable(Table622,98-100).Recommendation1.1.4.1:Wesuggestperformingakidneybiopsyasanacceptable,safe,diagnostictesttoevaluatecauseandguidetreatmentdecisionswhenclinicallyappropriate(2D).1.2EvaluationofGFR1.2.1OtherfunctionsofkidneysbesidesGFR
PracticePoint1.2.1.1:Usetheterm“GFR”whenreferringtothespecickidneyfunctionofglomerularltration.Usethemoregeneralterm“kidneyfunction(s)”whendealingwiththetotalityoffunctionsofthekidney.

---

### Chunk 10/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.536

romnaturalvariation.Lastly,itshouldbenotedthatrestrictingsaltintake
mayhelpensuremaximaleffectsofRASi.PracticePoint3.6.7:ContinueACEiorARBinpeoplewithCKDevenwhentheeGFRfallsbelow30ml/minper
1.73m2.InarecentSTOP-ACEitrialof411participantswithmeaneGFRof13ml/minper1.73m2,apolicyofdiscontinuingRASiinCKDG4–G5didnotresultinanykidneyorcar-diovascularbenets.507Twoobservationalstudieshavealsofoundthatassociationssuggestingoutcomeswereworse
amongparticipantswhostoppedRASiafterreachingan
www.kidney-international.orgchapter3KidneyInternational(2024)105(Suppl4S),S117–S314S213

eGFR<30ml/minper1.73m2,comparedwiththosewhocontinue.508,509Inaddition,arecentindividualpatientleveldatameta-analysisdemonstratedabenetindelayingKRTinpatientswitheGFR<30ml/minper1.73m2.5103.7Sodium-glucosecotransporter-2inhibitors(SGLT2i)TheWorkGroupconcurswiththeKDIGO2022ClinicalPracticeGuidelineforDiabetesManagementinChronicKidneyDisease,whichstated:“Werecommendtreatingpa-tientswithtype2diabetes(T2D),CKD,andaneGFR$20

---

### Chunk 11/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.534

okK,etal.Cross-sectionalandlongitudinal
performanceofcreatinine-andcystatinC-basedestimatingequations
relativetoexogenouslymeasuredglomerularltrationrateinHIV-positiveandHIV-negativepersons.JAcquirImmuneDecSyndr.2020;85:e58–e66.158.DelanayeP,CavalierE,MorelJ,etal.Detectionofdecreasedglomerularltrationrateinintensivecareunits:serumcystatinCversusserumcreatinine.BMCNephrol.2014;15:9.159.CarlierM,DumoulinA,JanssenA,etal.Comparisonofdifferentequationstoassessglomerularltrationincriticallyillpatients.IntensiveCareMed.2015;41:427–435.160.SanglaF,MartiPE,VerissimoT,etal.MeasuredandestimatedglomerularltrationrateintheICU:aprospectivestudy.CritCareMed.2020;48:e1232–e1241.161.WagnerD,KniepeissD,StieglerP,etal.TheassessmentofGFRafter
orthotopiclivertransplantationusingcystatinCandcreatinine-basedequations.TransplInt.2012;25:527–536.162.JanusN,Launay-VacherV,ByloosE,etal.CancerandrenalinsufciencyresultsoftheBIRMAstudy.BrJCancer.2010;103:1815–1821.163.Launay-VacherV,JanusN,DerayG.Re

---

### Chunk 12/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.532

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

### Chunk 13/30
**Article:** Urea levels and cardiovascular disease in patients with chronic kidney disease (2022)
**Journal:** Nephrology Dialysis Transplantation
**Section:** discussion | **Similarity:** 0.531

pressure,historyofcardiovasculardisease,anaemia,serumalbumin,high-sensitivityC-reactiveprotein,historyofacutekidneyinjury,diureticprescription,protonpumpinhibitorprescriptionandprescriptionsofstatinsandantiplateletagentsatbaseline.4.1[3.6;4.6]per100PY.Theincidenceratewas3.5timeshigherinpatientsinT3thaninpatientsinT1(Supplementarydata,TableS2,Figure4).TheadjustedHR[95%CI]formajorCVeventsassociatedwiththebaselineserumurealevelwas1.13[0.77;1.67]forpatientsinT2and1.94[1.27;2.98]forpatientsinT3(Figure4,Supplementarydata,TableS6).DeathbeforeRRTOveramedianfollow-upperiodof4.8[3.3;5.1]yearsintheCKD-REINcohort,407patientsdiedbeforeRRT,leadingtoanincidencerateof4.0[3.6;4.3]per100person-years(Supplementarydata,TableS2,Figure5).TheadjustedHRs[95%CI]fordeathbeforeRRTwere1.31[0.97;1.76]and1.73[1.22;2.45]forpatientsinT2andT3,respectively(Figure5,Supplementarydata,TableS7).DISCUSSIONOuranalysisofalargeprospectivecohortofCKDpatientsnotreceivingmaintenancedialysisattendingnephrologyout-patientfacilities

---

### Chunk 14/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.531

sing/kinetics,etc.Thereisnowarecognitionbymajorregulatoryagenciesthat“anycontemporary,widelyaccepted,andclinicallyapplicableestimatingGFRequationisconsideredreasonable
toassessGFRinpharmacokineticstudies.”770,771PracticePoint4.2.4:Inpeoplewithextremesofbody
weight,eGFRnonindexedforbodysurfacearea(BSA)may
beindicated,especiallyformedicationswithanarrow
therapeuticrangeorrequiringaminimumconcentration
tobeeffective.ForassessmentofCKD,itisrelevanttocomparetheGFRaccordingtoastandardbodysize.Forthisreason,GFRestimatingequationshavebeendevelopedinunitsofml/minper1.73m2.However,becausedrugclearanceismorestronglyassociatedwithnonindexedeGFR(ml/min)than
indexedeGFR(ml/minper1.73m2),inverysmallorlargeindividuals,thiscanresultinover-orunderdosing,respec-
tively,aswellasnoninitiationofcertainmedications.772,773NonindexedeGFRcanbeobtainedbymultiplyingthe
indexedeGFRresultsbytheperson’sBSAanddividingby1.73m2,orbyusinganappropriateonlinecalculator.PracticePoint4.2.5:Considerandadaptdrugdosingin

---

### Chunk 15/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.531

eknowledgegainedmoreeffec-
tively.793–797PracticalimplementationtipsinvolveprintingouttheresultsofthemostrecenteGFRestimationforthe
patienttobringalonginfuturehealthcareconsultationsand/
orwritedownalistofongoingmedicationstoalertotherhealthcareprovidersofmedicationrisksandbenets.AdiagnosisofCKDshouldalwaysbereectedinmedicalrecords,asthiswillalertphysiciansontheneedtoconsideradjustingoravoidingcertainmedicationsorprocedures.Under-recognitionofCKDdiagnosesinmedicalrecordsis
associatedwithmedicationerrors,includingpotentially
inappropriateprescriptionofnephrotoxicmedications.729PracticePoint4.3.1.2:Establishcollaborativerelationships
withotherhealthcareprovidersandpharmacistsand/or
usetoolstoensureandimprovedrugstewardshipinpeople
withCKDtoenhancemanagementoftheircomplex
medicationregimens.Clinicalpharmacistsarehighlyqualiedexpertsinmedi-cinesand,aspartofthemultidisciplinaryteam,canplaya
pivotalroleinimprovingthequalityofcareandensuringpatientsafetyinarangeofways.Thisincludescarryin

---

### Chunk 16/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.531

onrate(GFR)andalbuminuriainpeoplewithchronickidneydisease(CKD)S199Figure14.(a)Predictedriskofkidneyfailureand(b)‡40%declineinestimatedglomerularltrationrate(eGFR)bychronickidneydisease(CKD)eGFR(G1–G5)andalbumin-to-creatinineratio(ACR)(A1–A3)categoriesinOptumLabsDataWarehouseS201Figure15.Transitionfromanestimatedglomerularltrationrate(eGFR)-basedtoarisk-basedapproachtochronickidneydiseasecareS202Figure16.Comparisonofriskofchronickidneydisease(CKD)progression(5-yearprobabilityofestimatedglomerularltrationrate[eGFR]<60ml/minper1.73m2)versuskidneyfailureinadultswithCKDG1–G2calculatedfromtheriskequationavailableathttps://www.ckdpc.org/risk-models.htmlS205Figure17.Chronickidneydisease(CKD)treatmentandriskmodicationS206Figure18.Holisticapproachtochronickidneydisease(CKD)treatmentandriskmodicationS208Figure19.Proteinguidelineforadultswithchronickidneydiseasenottreatedwithdialysis
contentswww.kidney-international.orgS120KidneyInternational(2024)105(Suppl4S),S117–S314

S209Figure2

---

### Chunk 17/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.527

obin,g/dl,mean(SD),n[3,561,622S233Table30.RandomizedcontrolledtrialsinthetreatmentofasymptomatichyperuricemiainpeoplewithCKDS247Table31.Keyexamplesofcommonmedicationswithdocumentednephrotoxicityand,whereavailable,selectednon-nephrotoxicalternativesS252Table32.MedicationsthatshouldbeconsideredfortemporarydiscontinuationbeforeelectivesurgeriesandpotentialperioperativeadverseeventsassociatedwiththeircontinueduseS253Table33.Potentialriskfactorsforcontrast-associatedacutekidneyinjuryS256Table34.BenetsandconsequencesofearlyversuslatereferralS257Table35.Factorsassociatedwithlatereferralforkidneyreplacementtherapyplanning
S257Table36.OutcomesexaminedinasystematicreviewbySmartetal.S258Table37.Recommendedpatient-reportedoutcomemeasurementtoolsforuseinpeoplewithCKD
www.kidney-international.orgcontentsKidneyInternational(2024)105(Suppl4S),S117–S314S119

S259Table38.ManagementstrategiesforcommonsymptomsinCKDS261Table39.Listofvalidatedassessmenttoolsformalnutrition
S262Table40.Keyfeaturesofexisti

---

### Chunk 18/30
**Article:** Urea levels and cardiovascular disease in patients with chronic kidney disease (2022)
**Journal:** Nephrology Dialysis Transplantation
**Section:** other | **Similarity:** 0.526

rosis2017;263:127–1369.KumaA,WangXH,KleinJDetal.Inhibitionofureatransporteramelioratesuremiccardiomyopathyinchronickidneydisease.FASEBJ2020;34:8296–830910.KalimS,KarumanchiSA,ThadhaniRIetal.Proteincarbamylationinkidneydisease:pathogenesisandclinicalimplications.AmJKidneyDis2014;64:793–80311.KalimS,BergAH,KarumanchiSAetal.Proteincarbamylationandchronickidneydiseaseprogressioninthechronicrenalinsuciencycohortstudy.NephrolDialTransplant2021;37:139–14712.MoriD,MatsuiI,ShimomuraAetal.Proteincarbamylationexacerbatesvascularcalcication.KidneyInt2018;94:72–9013.ApostolovEO,RayD,SavenkaAVetal.ChronicuremiastimulatesLDLcarbamylationandatherosclerosis.JAmSocNephrol2010;21:1852–185714.StengelB,MetzgerM,CombeCetal.Riskprole,qualityoflifeandcareofpatientswithmoderateandadvancedCKD:TheFrenchCKD-REINCohortStudy.NephrolDialTransplant2019;34:277–28615.LeveyAS,StevensLA,SchmidCHetal.Anewequationtoestimateglomerularltrationrate.AnnInternMed2009;150:60416.HicksKA,MahaeyKW,MehranRetal.201

---

### Chunk 19/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.526

s206–217,Copyrightª2018,withpermissionfromtheNationalKidneyFoundation,Inc.541
00.20.40.60.81Density2.533.544.555.566.5Potassium, mmol/
eGFR 60+eGFR 30–59eGFR <30
Figure29|Distributionofbloodpotassiumingeneralpopulationandhigh-riskcohortsfromtheChronicKidneyDiseasePrognosisConsortium,byestimatedglomerularltrationrate(eGFR).Densityreferstotheproportionofthepopulationexperiencing
serumpotassiumlevel(e.g.,0.08ofthepopulationwithaGFR>60haveapotassiumof3.8;conversely,0.2ofthepopulationwitha
GFR<30haveapotassiumof5.5).ReproducedfromKovesdyCP,MatsushitaK,SangY,etal.Serumpotassiumandadverseoutcomes
acrosstherangeofkidneyfunction:aCKDPrognosisConsortium
meta-analysis.EuropeanHeartJournal2018;39:1535–1542bypermissionofOxfordUniversityPressonbehalfoftheEuropean
SocietyofCardiology.555Allrightsreserved.ªTheAuthor(s)2018.InclusionunderaCreativeCommonslicenseisprohibited.https://doi.org/10.1093/eurheartj/ehy100
chapter3www.kidney-international.orgS224KidneyInternational(2024)105(Suppl4S),S11

---

### Chunk 20/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.524

unicatedtopeoplewithCKDanddocumentedinthemedicalrecord.4.3.1StrategiestopromotedrugstewardshipPracticePoint4.3.1.1:EducateandinformpeoplewithCKDregardingtheexpectedbenetsandpossiblerisksofmedicationssothattheycanidentifyandreportadverseeventsthatcanbemanaged.Peoplewithkidneydiseasehavearoleindrugstewardship,andgiventhattheymayreceivemedicationsfromnon-nephrologyhealthcareproviders,peoplewithCKDshouldbe
encouragedtoinformthoseprescribersthattheyhavekidney
diseasetofacilitateconsiderationofdosesandpotentialside
effectofmedications.Thus,educationandinformationfor
peoplewithCKDinclusivefortheirpopulation(i.e.,literacylevelandlanguages)areencouraged.Althoughbrochuresandconversationsmaybeuseful,interactiveelectronichealth
applicationshavebeenshowntobeacceptabletopatientsand
mayleadthemtoapplytheknowledgegainedmoreeffec-
tively.793–797PracticalimplementationtipsinvolveprintingouttheresultsofthemostrecenteGFRestimationforthe
patienttobringalonginfuturehealthcareconsultationsand/
orwritedown

---

### Chunk 21/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.523

arltrationrate[eGFRcr-cys])(1B).1.1.3EvaluationofchronicityPracticePoint1.1.3.1:Proofofchronicity(durationofaminimumof3months)canbeestablishedby:(i)reviewofpastmeasurements/estimationsofGFR;(ii)reviewofpastmeasurementsofalbuminuriaorproteinuriaandurinemicroscopicexaminations;(iii)imagingndingssuchasreducedkidneysizeandreductionincorticalthickness;(iv)kidneypathologicalndingssuchasbrosisandatrophy;(v)medicalhistory,especiallyconditionsknowntocauseorcontributetoCKD;(vi)repeatmeasurementswithinandbeyondthe3-monthpoint.PracticePoint1.1.3.2:DonotassumechronicitybaseduponasingleabnormallevelforeGFRandACR,asthendingcouldbetheresultofarecentacutekidneyinjury(AKI)eventoracutekidneydisease(AKD).PracticePoint1.1.3.3:ConsiderinitiationoftreatmentsforCKDatrstpresentationofdecreasedGFRorelevatedACRifCKDisdeemedlikelyduetopresenceofotherclinicalindicators.1.1.4Evaluationofcause
PracticePoint1.1.4.1:EstablishthecauseofCKDusingclinicalcontext,personalandfamilyhistory,socialandenvi-ronmentalfactor

---

### Chunk 22/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.523

etoflifestylechangesMonitorprogressionofGFRdeclineEligibilityforclinicaltrialsTreatmentofAKIMonitoringdrugtoxicityRe-evaluateCKDtreatment
strategiesRiskassessmentRiskofCKDcomplicationsRiskforCKDprogressionRiskofCVDRiskformedicationerrorsRiskforperioperativecomplicationsRiskformortalityFertilityandriskofcomplications
ofpregnancyRiskforCKDprogressionRiskforCVDRiskformortalityFertilityandriskofcomplications
ofpregnancyRiskforkidneyfailureRiskforCVD,HF,andmortalityRiskforadversepregnancy
outcomeAKD,acutekidneydisease;AKI,acutekidneyinjury;CKD,chronickidneydisease;CVD,cardiovasculardisease;GFR,glomerularltrationrate;HF,heartfailure.

---

### Chunk 23/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.521

skmodicationS208Figure19.Proteinguidelineforadultswithchronickidneydiseasenottreatedwithdialysis
contentswww.kidney-international.orgS120KidneyInternational(2024)105(Suppl4S),S117–S314

S209Figure20.AverageproteincontentoffoodsingramsS214Figure21.Algorithmformonitoringofpotassiumandestimatedglomerularltrationrate(eGFR)aftertheinitiationofrenin-angiotensinsysteminhibitorsS215Figure22.Effectofsodium-glucosecotransporter-2inhibitors(SGLT2i)withkidneydiseaseoutcomesbydiabetesstatusS216Figure23.Effectsofsodium-glucosecotransporter-2(SGLT2)inhibitionversusplacebooncardiovascularandmortalityoutcomesbydiabetesstatusandtrialpopulationS217Figure24.Effectsofsodium-glucosecotransporter-2(SGLT2)inhibitionversusplaceboonkidneyfailure(chronickidneydisease[CKD]trials)S219Figure25.Effectsofempagliozinversusplaceboonannualrateofchangeinestimatedglomerularltrationrate(GFR)bykeysubgroupsintheStudyofHeartandKidneyProtectionWithEmpagliozin(EMPA-KIDNEY)S220Figure26.Serumpotassiummonitoringduringtreatm

---

### Chunk 24/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.521

indicatedbyinsulinresistance,higherlevelsofC-reactiveproteinandtumornecrosisfactor,
orlowerlevelsofserumalbumin).129,130,176–185
Table10|IndicationsformeasuredGFRClinicalconditionsinwhicheGFRcr-cysisinaccurateoruncertainduetopotentialnon-GFRdeterminantsofcreatinineandcystatinC.Thismayincludecatabolicstates,suchasseriousinfectionsorinammatorystates,highcellturnoverasinsomecancers,advancedcirrhosisorheartfailure,useofhigh-dosesteroids,ortheveryfrail.SeeFigure12forapproachtoindividualdecision-making.ClinicalsettingsinwhichgreateraccuracyisneededthanisachievedwitheGFRcr-cys.Forexample,decisionsaboutsimultaneouskidneytransplantatthetimeofothersolidorgantransplant,kidneydonorcandidacy,anddrugdosingifthereisanarrowtherapeuticindexorserioustoxicity(e.g.,
chemotherapiesthatareclearedbythekidney).eGFRcr-cys,estimatedGFRbycreatinineandcystatinC;GFR,glomerularltrationrate.

---

### Chunk 25/30
**Article:** Proteinuria (2023)
**Journal:** StatPearls Publishing
**Section:** abstract | **Similarity:** 0.520

Comprehensive clinical review of proteinuria covering definition, etiology, diagnosis, and management. Proteinuria serves as a marker of kidney damage assisting with diagnosis, prognosis, and therapy. Normal protein excretion is <150 mg/24 hours. Severity categories: nephritic range (150-3,500 mg/24h) and nephrotic range (>3,500 mg/24h). Urine dipstick provides semi-quantitative screening but has important limitations: false-positives with dehydration, alkaline urine, UTI; inability to detect positively charged proteins. UK CKD guidelines define significant proteinuria as UPCR >45 mg/mmol. Quantitative methods (spot UPCR preferred for convenience and reliability, or 24-hour urine collection for accuracy) are indicated when dipstick is positive or proteinuria >1 g/day is suspected. Early detection and quantification are crucial given associations with progressive renal disease and cardiovascular complications.

---

### Chunk 26/30
**Article:** Urea levels and cardiovascular disease in patients with chronic kidney disease (2022)
**Journal:** Nephrology Dialysis Transplantation
**Section:** other | **Similarity:** 0.520

.0016.9%Serumalbumin(g/L)40.4(4.5)40.6(4.4)40.5(4.2)39.9(4.9)0.00915.2%High-sensitivityC-reactiveprotein(mg/L)2.5[1.1–5.9]2.2[1.1–5.0]2.5[1.1–5.4]2.9[1.2–7.1]<0.00117.6%Historyofacutekidneyinjury,n(%)22.520.020.926.60.0047.7Diureticprescriptionatbaseline,n(%)54.039.951.170.7<0.0010.3PPIprescriptionatbaseline,n(%)31.427.933.932.40.020.3RASiprescriptionatbaseline,n(%)76.972.380.278.0<0.0010.3Statinprescriptionatbaseline,n(%)58.951.461.064.1<0.0010.%Antiplateletprescriptionatbaseline,n(%)42.137.643.145.60.0030.3
eGFR,estimatedglomerularltrationrate,basedontheCKD-EPIequation;RASi,renin-angiotensinsysteminhibitor;PPI,protonpumpinhibitor;T,tertile.Dataarequotedasthefrequency(%),mean(standarddeviation)orthemedian[interquartilerange].AnANOVAoraKruskal–WallisorChi-squaredtestwasusedtocomparegroups.

---

### Chunk 27/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** methods | **Similarity:** 0.519

ey-international.orgS148KidneyInternational(2024)105(Suppl4S),S117–S314

SummaryofrecommendationstatementsandpracticepointsChapter1:EvaluationofCKD1.1DetectionandevaluationofCKD1.1.1DetectionofCKDPracticePoint1.1.1.1:Testpeopleatriskforandwithchronickidneydisease(CKD)usingbothurinealbuminmea-surementandassessmentofglomerularltrationrate(GFR).PracticePoint1.1.1.2:Followingincidentaldetectionofelevatedurinaryalbumin-to-creatinineratio(ACR),hematuria,orlowestimatedGFR(eGFR),repeatteststoconrmpresenceofCKD.1.1.2MethodsforstagingofCKDRecommendation1.1.2.1:InadultsatriskforCKD,werecommendusingcreatinine-basedestimatedglomer-ularltrationrate(eGFRcr).IfcystatinCisavailable,theGFRcategoryshouldbeestimatedfromthecombinationofcreatinineandcystatinC(creatinineandcystatinC–basedestimatedglomerularltrationrate[eGFRcr-cys])(1B).1.1.3EvaluationofchronicityPracticePoint1.1.3.1:Proofofchronicity(durationofaminimumof3months)canbeestablishedby:(i)reviewofpastmeasurements/estimationsofGFR;(ii)revie

---

### Chunk 28/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.519

.65.05.36.07.01.42.03.04.15.4Ref1.31.92.73.61.01.41.72.43.21.41.72.22.83.82.02.32.83.74.63.23.13.55.06.56.16.46.47.38.2
Figure36|Riskofall-causeandcardiovascularmortalitybyestimatedglomerularltrationrate(eGFR)andlevelofalbuminuriafromgeneralpopulationcohortscontributingtotheChronicKidneyDiseasePrognosisConsortium.ACR,albumin-to-creatinineratio;eGFRcr,creatinine-basedestimatedglomerularltrationrate.ReproducedwithpermissionfromJAMA,WritingGroupfortheCKDPrognosisConsortium;GramsME,CoreshJ,MatsushitaK,etal.Estimatedglomerularltrationrate,albuminuria,andadverseoutcomes:anindividual-participantdatameta-analysis.JAMA.2023;330(13):1266–1277.12Copyrightª2023AmericanMedicalAssociation.Allrightsreserved.

---

### Chunk 29/30
**Article:** Urea levels and cardiovascular disease in patients with chronic kidney disease (2022)
**Journal:** Nephrology Dialysis Transplantation
**Section:** results | **Similarity:** 0.518

vel,bloodhaemoglobin,serumalbuminandalbumin-orprotein-to-creatinineratio(Supplementarydata,TableS1)wererecorded.TheGFRwasestimatedusingtheChronicKidneyDiseaseEpidemiologyCollaboration(CKD-EPI)creatinineequationandtheisotopedilutionmassspectrometry-traceablecreatinineconcentrationdeterminedinaJaeassay(in38%ofparticipants)oranenzymaticassay(in57%;theassaytypewasunknownfor5%)[15].Patientswereaskedtobringalltheirdrugprescriptionsforthepreceding3monthstotheenrolmentvisit.TheCRAsusedanelectroniccasereportform[linkedtotheinternationalAnatomicalTherapeuticandChemical(ATC)thesaurus]toenterstandardizedATCcodes.Longitudinalclinicaldata(nephrologyconsultations,hos-pitaladmissions,laboratorytestresults,medicationsandanytransitiontodialysisortransplantation)werecollectedeveryyearfrommedicalrecordsandpatientinterviews.StudyoutcomesCVeventswereassessedcarefullyaccordingtotheCardio-vascularandStrokeEndpointDenitionsforClinicalTrials[16].Theprimaryendpointwastheoccurrenceofarstfatalornonfatalatheroma

---

### Chunk 30/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.517

eneral,all3devicesdemonstratedacceptableaccuracyatlowerlevelsofeGFR(<30ml/minper1.73m2).316Resultsshowedthati-STATandABLdevicesmayhavehigherprobabilitiesofcorrectlyclassifyingpeopleinthesameeGFRcategoriesasthelaboratoryreferencethan
StatSensordevices.Foralbumin,theERTidentiedasystematicreviewpub-lishedin2014,byMcTaggartetal.,318thatevaluatedthediagnosticaccuracyofquantitativeandsemiquantitativeproteinoralbuminurinedipsticktestscomparedwithlaboratory-basedtestsamongpeoplewithsuspectedor
diagnosedCKD.TheERTincludedrelevantstudiesfrom
thisreviewandconductedanupdate.Sixty-vestudies(in66articles)319–344,345–368,369–384eval-uatedtheaccuracyofquantitativeandsemiquantitativepro-
teinoralbumindipsticktestsinageneralpopulationnoton
chapter1www.kidney-international.orgS194KidneyInternational(2024)105(Suppl4S),S117–S314

KRTorreceivingend-of-lifecare.Studiesaddressedthefollowingcriticaloutcomes:measurementbias(n¼1),analyticalvariability(n¼5),analyticalsensitivity(n¼2),andanalyticspeci

---

