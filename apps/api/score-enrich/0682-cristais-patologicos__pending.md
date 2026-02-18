# ScoreItem: Cristais Patológicos

**ID:** `019bf31d-2ef0-7754-adf7-50c7b130b286`
**FullName:** Cristais Patológicos (Exames - Laboratoriais)
**Unit:** Descritivo

**Preparation Metadata:**
- Quality Grade: **GOOD**
- Total Chunks: 30 de 8 artigos
- Avg Similarity: 0.539

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-7754-adf7-50c7b130b286`.**

```json
{
  "score_item_id": "019bf31d-2ef0-7754-adf7-50c7b130b286",
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

**ScoreItem:** Cristais Patológicos (Exames - Laboratoriais)
**Unidade:** Descritivo

**30 chunks de 8 artigos (avg similarity: 0.539)**

### Chunk 1/30
**Article:** Outpatient crystalluria: prevalence, crystal types, and associations with comorbidities and urinary tract infections at a provincial hospital (2025)
**Journal:** Iranian Journal of Microbiology
**Section:** abstract | **Similarity:** 0.718

Background and Objectives: Crystalluria refers to the occurrence of crystals in urine resulting from urinary supersaturation, which disrupts the balance between factors that promote and those that inhibit crystal formation in urine. This study examined crystalluria occurrence in 1,025 urine samples from outpatients at a provincial hospital. Results: 22.04% of samples showed crystalluria with mean patient age of 51.3 years. The most common crystal types were calcium oxalate (46.4%), uric acid (23.5%), urates (15.1%) and struvite (9.3%). Diabetes, kidney failure, prostatitis, and nephrotic syndrome were associated with crystal formation. The prevalence of urinary tract infections in patients with urinary crystals was 10.6%. Notably, struvite crystals correlated with bacterial infections. Conclusion: Monitoring crystals is vital for preventing kidney stone development and infection-related complications in susceptible populations.

---

### Chunk 2/30
**Article:** Pathophysiology and Main Molecular Mechanisms of Urinary Stone Formation and Recurrence (2024)
**Journal:** International Journal of Molecular Sciences
**Section:** abstract | **Similarity:** 0.637

Urinary stone disease (nephrolithiasis) is a multifactorial condition affecting 12% of the global population with recurrence rates as high as 50% within 5 years after first stone onset. This comprehensive review examines the pathophysiology and molecular mechanisms underlying urinary stone formation and recurrence. The review covers crystal nucleation, aggregation, and retention processes, highlighting the role of urinary supersaturation, pH imbalances, and inhibitor deficiencies. Special attention is given to calcium oxalate and calcium phosphate stones (80% of cases), uric acid stones, cystine stones, and infection-related struvite crystals. Understanding these mechanisms is essential for developing targeted prevention and treatment strategies for this highly recurrent condition.

---

### Chunk 3/30
**Article:** Amoxicillin crystalluria and amoxicillin-induced crystal nephropathy: a narrative review (2025)
**Journal:** Kidney International
**Section:** abstract | **Similarity:** 0.572

Amoxicillin crystalluria is defined as the precipitation of amoxicillin trihydrate crystals in urine, while amoxicillin-induced crystal nephropathy involves kidney tubule obstruction causing acute kidney injury. This narrative review reveals higher prevalence (24-41%) than previously recognized in patients receiving high-dose intravenous amoxicillin (≥150 mg/kg daily). The condition typically presents with cloudy urine and macroscopic hematuria. Management involves discontinuing the antibiotic and aggressive fluid resuscitation to enhance urine flow. Patients generally recover normal kidney function rapidly after stopping amoxicillin, though 10-40% require renal replacement therapy. No deaths have been directly attributed to this condition.

---

### Chunk 4/30
**Article:** Proteinuria (2023)
**Journal:** StatPearls Publishing
**Section:** abstract | **Similarity:** 0.570

Comprehensive clinical review of proteinuria covering definition, etiology, diagnosis, and management. Proteinuria serves as a marker of kidney damage assisting with diagnosis, prognosis, and therapy. Normal protein excretion is <150 mg/24 hours. Severity categories: nephritic range (150-3,500 mg/24h) and nephrotic range (>3,500 mg/24h). Urine dipstick provides semi-quantitative screening but has important limitations: false-positives with dehydration, alkaline urine, UTI; inability to detect positively charged proteins. UK CKD guidelines define significant proteinuria as UPCR >45 mg/mmol. Quantitative methods (spot UPCR preferred for convenience and reliability, or 24-hour urine collection for accuracy) are indicated when dipstick is positive or proteinuria >1 g/day is suspected. Early detection and quantification are crucial given associations with progressive renal disease and cardiovascular complications.

---

### Chunk 5/30
**Article:** Sulfamethoxazole-induced crystal nephropathy: characterization and prognosis in a case series (2024)
**Journal:** Scientific Reports
**Section:** abstract | **Similarity:** 0.558

This case series examined patients who developed acute kidney injury under sulfamethoxazole treatment to assess the causal relationship between the drug and kidney injury in the presence of N-acetyl-sulfamethoxazole (NASM) crystals in urine. Drug-induced crystalluria represents an important but underrecognized cause of acute kidney injury. The study characterized the clinical presentation, crystal morphology, and prognosis of sulfamethoxazole-induced crystal nephropathy. Findings emphasize the importance of early crystal identification in urine microscopy and prompt discontinuation of the causative medication to prevent irreversible renal dysfunction. Most patients showed renal function recovery after drug cessation, highlighting the generally favorable prognosis with timely intervention.

---

### Chunk 6/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.554

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

### Chunk 7/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.551

tNZJMed.1992;22:334–337.338.GilbertRE,AkdenizA,JerumsG.Detectionofmicroalbuminuriaindiabeticpatientsbyurinarydipstick.DiabetesResClinPract.1997;35:57–60.339.GrazianiMS,GambaroG,MantovaniL,etal.Diagnosticaccuracyofa
reagentstripforassessingurinaryalbuminexcretioninthegeneralpopulation.NephrolDialTransplant.2008;24:1490–1494.340.GuyM,NewallR,BorzomatoJ,etal.Useofarst-lineurineprotein-to-creatinineratiostriptestonrandomurinestoruleoutproteinuriain
patientswithchronickidneydisease.NephrolDialTransplant.2008;24:1189–1193.341.GuyM,NewallR,BorzomatoJ,etal.Diagnosticaccuracyoftheurinary
albumin:creatinineratiodeterminedbytheCLINITEKMicroalbuminandDCA2000þfortherule-outofalbuminuriainchronickidneydisease.ClinChimActa.2008;399:54–58.342.HasslacherC,MullerP,SchlipfenbacherRL.ResultsofamulticentrestudyforthedeterminationofmicroalbuminuriawithMicral-Test.KlinischesLabor.1995;41:441–447.343.HodelNC,HamadA,ReitherK,etal.Comparisonoftwodifferent
semiquantitativeurinarydipsticktestswithal

---

### Chunk 8/30
**Article:** The Role of Urinary Nitrite in Predicting Bacterial Resistance in Urine Culture Analysis Among Patients With Uncomplicated Urinary Tract Infection (2022)
**Journal:** Cureus
**Section:** results | **Similarity:** 0.549

%-90% of cases, 
Escherichia coli
 is the
etiologic agent 
[1,2]
. Early detection and timely and adequate treatment of the disease are critical to
preventing major complications such as renal scarring, hypertension, and chronic renal failure 
[3,4]
.
Regarding the rising treatment resistance of microorganisms causing UTIs, appropriate selection of the
initial antibiotic before collecting urine culture and antibiogram results are critical 
[1,2,5-9]
. The question is
whether it is possible to predict microorganism resistance using markers, such as the nitrite test, prior to
preparing urine cultures and antibiogram results. According to Weisz et al., urine nitrite results can be used
to identify cephalosporin resistance 
[10]
. On the other hand, other studies contradict this viewpoint 
[11,12]
.
Antimicrobial resistance is increasing globally, making infections more difficult to treat, and is related to
increased mortality, morbidity, and cost 
[13,14]
.

---

### Chunk 9/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.547

inmostCKDs.Useof
urinaryalbuminmeasurementasthepreferredtestforpro-teinuriadetectionwillimprovethesensitivity,quality,andconsistencyofapproachtotheearlydetectionandmanage-
mentofkidneydisease.Commonlyusedreagentstripdevicesmeasuringtotalproteinareinsufcientlysensitiveforthereliabledetectionofproteinuria,donotadjustforurinaryconcentration,andare
onlysemiquantitative.Furthermore,thereisnostandardization
betweenmanufacturers.Theuseofsuchstripsshouldbediscouragedinfavorofquantitativelaboratorymeasurementsofalbuminuriaorproteinuria,orvalidatedpoint-of-carede-
vicesforurinealbumin/ACR(Section1.4).Whenused,reagentstripresultsshouldbeconrmedbylaboratorytesting.Althoughthereferencepointremainstheaccuratelytimed24-hourspecimen,itiswidelyacceptedthatthisisadifcultproceduretocontroleffectivelyandthatinaccuraciesinurinary
collectionmaycontributetoerrorsinestimationofalbuminand/orproteinlosses.Inpractice,untimedurinesamplesareareasonablersttestforascertainmentofalbuminuria.Arstmorningvoidsample

---

### Chunk 10/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.537

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
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.536

lurinarytractdiseaseRecurrentkidneycalculiMultisystemdiseases/chronicinammatoryconditionsSystemiclupuserythematosus
Vasculitis
HIVIatrogenic(relatedtodrugtreatmentsandprocedures)Drug-inducednephrotoxicityandradiationnephritisFamilyhistoryorknowngeneticvariantassociatedwithCKDKidneyfailure,regardlessofidentiedcauseKidneydiseaserecognizedtobeassociatedwithgeneticabnormality(e.g.,PKD,APOL1-mediatedkidneydisease,andAlportsyndrome)GestationalconditionsPretermbirthSmallgestationalsizePre-eclampsia/eclampsiaOccupationalexposuresthatpromoteCKDriskCadmium,lead,andmercuryexposurePolycyclichydrocarbonsPesticidesAKD,acutekidneydisease;AKI,acutekidneyinjury;APOL1,apolipoproteinL1;CKD,chronickidneydisease;CKDu,chronickidneydiseaseofundeterminedorigin;PKD,polycystickidneydisease.

---

### Chunk 12/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.533

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

### Chunk 13/30
**Article:** Automated urine sediment analyzers underestimate the severity of hematuria in glomerular diseases (2021)
**Journal:** Scientific Reports
**Section:** results | **Similarity:** 0.531

100 RBCs/HPF).e presence of blood and the pH of the urine were measured via a dipstick test. e urine color and the dipstick test results were read via an automatic reader. e Urisys 2400 automated urine test strip analyzer (Roche Diagnostics GmbH) was used to read the dipstick tests of the urine specimens examined using the UF-1000i urine analyzer, whereas the Cobas 6500 urine analyzer was used to conduct the urine dipstick test as well as the urine RBC counting using two modules, the Cobas u 601 for the urine dipstick and the Cobas u 701 for urine microscopy. Test strips in the Urisys 2400 cassettes and the cobas u packs were used with the Urisys 2400 analyzer and the Cobas u 601 analyzer, respectively. In both analyzers, dipstick hematuria was graded as −, ±, 1+, 2+, 3+, and 4+.

---

### Chunk 14/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** methods | **Similarity:** 0.529

JExpClinMed.2019;44:118–123.371.SalinasM,López-GarrigósM,FloresE,etal.Urinaryalbuminstripassay
asascreeningtesttoreplacequantitativetechnologyincertain
conditions.ClinChemLabMed.2018;57:204–209.372.SaradisPA,RiehleJ,BogojevicZ,etal.Acomparativeevaluationofvariousmethodsformicroalbuminuriascreening.AmJNephrol.2007;28:324–329.373.ShephardMD,BarrattLJ,Simpson-LyttleW.IstheBayerDCA2000
acceptableasascreeninginstrumentfortheearlydetectionofrenal
disease?AnnClinBiochem.1999;36(Pt3):393–394.374.SiednerMJ,GelberAC,RovinBH,etal.Diagnosticaccuracystudyofurinedipstickinrelationto24-hourmeasurementasascreeningtoolforproteinuriainlupusnephritis.JRheumatol.2007;35:84–90.375.SpoorenPF,LekkerkerkerJF,VermesI.Micral-test:aqualitativedipsticktestformicro-albuminuria.DiabetesResClinPract.1992;18:83–87.376.SzymanowiczA,Blanc-BernardE,RocheC,etal.EvaluationofMicralTestforthescreeningofthemicroalbuminuriainpointofcaretesting.Immuno-AnalyseBiologieSpecialisee.2008;23:109–115.377.TiuSC,LeeSS

---

### Chunk 15/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.527

tospecialistkidneycareservicesinthecircumstanceslistedinFigure48.SpecialconsiderationsPediatricconsiderations.PracticePoint5.1.2:Referchildrenandadolescentstospecialistkidneycareservicesinthefollowingcircumstances:anACRof30mg/g(3mg/mmol)oraPCRof200mg/g(20mg/mmol)ormore,conrmedonarepeatrstmorningvoidsample,whenwellandnotduringmenstruation,persistenthematuria,anysustaineddecreaseineGFR,hypertension,kidneyoutowobstructionoranomaliesofthekidneyandurinarytract,knownorsuspectedCKD,orrecurrenturinarytractinfection.5.2SymptomsinCKD5.2.1Prevalenceandseverityofsymptoms[Norecommendationsandpracticepoints]5.2.2IdenticationandassessmentofsymptomsPracticePoint5.2.2.1:AskpeoplewithprogressiveCKDabouturemicsymptoms(e.g.,reducedappetite,nausea,andleveloffatigue/lethargy)ateachconsultationusingastandardizedvalidatedassessmentofuremicsymptomstool.

---

### Chunk 16/30
**Article:** The Role of Urinary Nitrite in Predicting Bacterial Resistance in Urine Culture Analysis Among Patients With Uncomplicated Urinary Tract Infection (2022)
**Journal:** Cureus
**Section:** results | **Similarity:** 0.522

wis LM: 
Can urinary nitrite results be used to guide antimicrobial choice
for urinary tract infection?
. J Emerg Med. 1997, 15:435-8. 
10.1016/s0736-4679(97)00069-3
13
. 
Walker E, Lyman A, Gupta K, Mahoney MV, Snyder GM, Hirsch EB: 
Clinical management of an increasing
threat: outpatient urinary tract infections due to multidrug-resistant uropathogens
. Clin Infect Dis. 2016,
63:960-5. 
10.1093/cid/ciw396
14
. 
Prakash V, Lewis JS 2nd, Herrera ML, Wickes BL, Jorgensen JH: 
Oral and parenteral therapeutic options for
outpatient urinary infections caused by enterobacteriaceae producing CTX-M extended-spectrum beta-
lactamases
. Antimicrob Agents Chemother. 2009, 53:1278-80. 
10.1128/AAC.01519-08
15
. 
Spellberg B, Guidos R, Gilbert D, et al.: 
The epidemic of antibiotic-resistant infections: a call to action for
the medical community from the Infectious Diseases Society of America
. Clin Infect Dis. 2008, 46:155-64.
10.1086/524891
16
. 
Bono MJ, Reygaert WC: 
Urinary Tract Infection
.

---

### Chunk 17/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.519

rshallSM,AlbertiKG.Screeningforearlydiabeticnephropathy.AnnClinBiochem.1986;23(Pt2):195–197.290.MillerWG,BrunsDE,HortinGL,etal.Currentissuesinmeasurementandreportingofurinaryalbuminexcretion.ClinChem.2009;55:24–38.291.ChitaliaVC,KothariJ,WellsEJ,etal.Cost-benetanalysisandpredictionof24-hourproteinuriafromthespoturineprotein-creatinineratio.ClinNephrol.2001;55:436–447.292.CoteAM,BrownMA,LamE,etal.Diagnosticaccuracyofurinaryspotprotein:creatinineratioforproteinuriainhypertensivepregnant
women:systematicreview.BMJ.2008;336:1003–1006.293.DysonEH,WillEJ,DavisonAM,etal.Useoftheurinaryproteincreatinine
indextoassessproteinuriainrenaltransplantpatients.NephrolDialTransplant.1992;7:450–452.294.Leanos-MirandaA,Marquez-AcostaJ,Romero-ArauzF,etal.Protein:
creatinineratioinrandomurinesamplesisareliablemarkerof
increased24-hourproteinexcretioninhospitalizedwomenwith
hypertensivedisordersofpregnancy.ClinChem.2007;53:1623–1628.295.LemannJJr,DoumasBT.Proteinuriainhealthanddiseaseassessedby

---

### Chunk 18/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.518

wever,currentcommerciallyavailableassaysforurinealbuminarenotstandardizedagainstthisreferencematerial.Laboratoriesshouldensurethattheyareenrolledanddemonstrate
satisfactoryperformancein,anEQAschemeforurine
albumin,creatinine,andACR.Urinealbumin(andprotein)concentrationsinurineshouldbereportedasaratiotocreatinine—ACR(orPCR).Reportingasaratiotocreatininecorrectsforvariationsin
urinaryowrateandenablesreportingonuntimed,spotsamples,obviatingtheneedfortimed,including24-hour,
collections,whicharepronetocollectionerrorandtedious
forpeopletoundertake.Reportingalbuminasaratioto
creatininereducestheintraindividualvariabilityinalbumin-
uriacomparedwithreportingasalbuminconcentrationalone
(mg/mmolormg/g).315Toaidclarityinreportingacrossandwithinhealthcaresystems,andtoprovideguidanceregardingthenumberofmeaningfuldigitsinaresult,astandardizedapproachshouldbeusedinrelationtoreportingunitsofACRandPCR.ACR
resultsshouldbeexpressedtoonedecimalplace(mg/mmol)
orwholenumbers(mg/g).BothenzymaticandJaffeas

---

### Chunk 19/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.517

ollectionundersupervisedcondi-
tionsasrecommendedabove.SpecialconsiderationsPediatricconsiderations.PracticePoint1.3.1.4:Inchildren,obtainarstmorningurinesampleforinitialtestingofalbuminuriaandpro-
teinuria(indescendingorderofpreference):(i)BothurinePCRandurineACR,(ii)Reagentstripurinalysisfortotalproteinandfor
albuminwithautomatedreading,or(iii)Reagentstripurinalysisfortotalproteinandfor
albuminwithmanualreading.ConsistentwiththeKDIGO2012ClinicalPracticeGuidelinefortheEvaluationandManagementofChronic
KidneyDisease,1PCRisadvisedandpreferredasinitialscreeningforchildrenasthemajorityofchildrenhave
underlyingdevelopmentalabnormalitiesoftenreferredtoas
CAKUTandamuchhigherproportionofchildrenthanadultshavetubularpathology.303TestingforACRmaymisstubularproteinuria.However,testingexclusivelyforproteinuriadoesnotallowcharacterizationofthesource.If
urinePCRisused,urineACRshouldalsobemeasuredto
chapter1www.kidney-international.orgS192KidneyInternational(2024)105(Suppl4S),S117–S314

betterchar

---

### Chunk 20/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.517

erofmeaningfuldigitsinaresult,astandardizedapproachshouldbeusedinrelationtoreportingunitsofACRandPCR.ACR
resultsshouldbeexpressedtoonedecimalplace(mg/mmol)
orwholenumbers(mg/g).BothenzymaticandJaffeassays
aregenerallysuitableforthemeasurementofcreatininein
urine,althoughhighconcentrationsofglucosecaninterfere
inJaffeurinecreatininemeasurementandproduceclinically
meaningfulerrorsinACR.1.4Point-of-caretestingRecommendation1.4.1:Wesuggestthatpoint-of-caretesting(POCT)maybeusedforcreatinineand
urinealbuminmeasurementwhereaccesstoalab-oratoryislimitedorprovidingatestatthepoint-of-carefacilitatestheclinicalpathway(2C).PracticePoint1.4.1:WheneveraPOCTdeviceisusedforcreatinineandurinealbumintesting,ensurethatthesame
preanalytical,analytical,andpostanalyticalqualitycriteria
relatingtothespecimencollectionandperformanceofthe
device,includingexternalqualityassessment,andtheinterpretationoftheresultisused.PracticePoint1.4.2:WhereaPOCTdeviceforcreatininetestingisbeingused,generateanestimateofGFR.Us

---

### Chunk 21/30
**Article:** Automated urine sediment analyzers underestimate the severity of hematuria in glomerular diseases (2021)
**Journal:** Scientific Reports
**Section:** results | **Similarity:** 0.516

and digital imaging in one system: A multicenter performance evaluation of the cobas 6500 urine work area. Pract. Lab. Med. 17, e00139 (2019). 30. aller, T. R. & Wang, L. P. Evaluation of asymptomatic microscopic hematuria in adults. Am. Fam. Phys. 60, 11431152 (1999). 31. Palsson, R., Srivastava, A. & Waikar, S. S. Performance of the automated urinalysis in diagnosis of proliferative glomerulonephritis. Kidney Int. Rep. 4, 723727 (2019). 32. Orlandi, P. F. et al. Hematuria as a risk factor for progression of chronic kidney disease and death: Findings from the chronic renal insuciency cohort (CRIC) study. BMC Nephrol. 19, 150 (2018). 33. Schulman, G. et al. Risk factors for progression of chronic kidney disease in the EPPIC trials and the eect of AST-120. Clin. Exp. Nephrol. 22, 299308 (2018). 34. Mashitani, T. et al.

---

### Chunk 22/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.514

22:369.318.McTaggartMP,NewallRG,HirstJA,etal.Diagnosticaccuracyofpoint-of-
caretestsfordetectingalbuminuria:asystematicreviewandmeta-analysis.AnnInternMed.2014;160:550–557.319.AbitbolC,ZillerueloG,FreundlichM,etal.Quantitationofproteinuriawithurinaryprotein/creatinineratiosandrandomtestingwithdipsticksinnephroticchildren.JPediatr.1990;116:243–247.320.AgardhCD.Anewsemiquantitativerapidtestforscreeningformicroalbuminuria.PracticalDiabetes.1993;10:146–147.321.AgarwalR,PanesarA,LewisRR.Dipstickproteinuria:canitguide
hypertensionmanagement?AmJKidneyDis.2002;39:1190–1195.322.AroraS,LongT,MenchineM.Testcharacteristicsofurinedipstickfor
identifyingrenalinsufciencyinpatientswithdiabetes.WestJEmergMed.2011;12:250–253.323.ChangCC,SuMJ,HoJL,etal.Theefcacyofsemi-quantitativeurineprotein-to-creatinine(P/C)ratioforthedetectionofsignicantproteinuriainurinespecimensinhealthscreeningsettings.Springerplus.2016;5:1791.324.ChoMC,JiM,KimSY,etal.EvaluationoftheURiSCANsupercassetteACRsemiquantita

---

### Chunk 23/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.514

creasedwitheGFRvalues>105ml/minper1.73m2(Figure712).IncontrasteGFRcr-cysdemonstratedmuchmorelinearassociationswitheachofthesecomplications
throughoutitsdistribution.Thesedatademonstratethatthe
combinedeGFRcr-cysequationissuperiorfordistinguishing
GFRriskstagescomparedwitheGFRcr.Certaintyofevidence.Thisrecommendationisbasedon2broadlydifferenttypesofdata—datacomparingtheaccuracy(P30)ofequationsfromacombinationofcreatinineandcystatinCasltrationmarkersandcreatinineandcystatinC
Table5|RiskfactorsforCKDDomainsExampleconditionsCommonriskfactorsHypertensionDiabetesCardiovasculardisease(includingheartfailure)
PriorAKI/AKDPeoplewholiveingeographicalareaswithhighprevalenceofCKDAreaswithendemicCKDu
AreaswiththehighprevalenceofAPOL1geneticvariantsEnvironmentalexposuresGenitourinarydisordersStructuralurinarytractdiseaseRecurrentkidneycalculiMultisystemdiseases/chronicinammatoryconditionsSystemiclupuserythematosus
Vasculitis
HIVIatrogenic(relatedtodrugtreatmentsandprocedures)Drug-inducednephrotox

---

### Chunk 24/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.512

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

### Chunk 25/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.512

eyDiseases,volume73,issue2,InkerLA,GramsME,LeveyAS,etal.RelationshipofestimatedGFRandalbuminuriatoconcurrentlaboratoryabnormalities:anindividualparticipantdatameta-
analysisinaGlobalConsortium,pages206–217,Copyrightª2018,withpermissionfromtheNationalKidneyFoundation,Inc.541
chapter3www.kidney-international.orgS230KidneyInternational(2024)105(Suppl4S),S117–S314

withaneGFRof<60ml/minper1.73m2atbaseline(amongwhom71primaryoutcomesaccrued)werecomparedwiththe5181peoplewithaneGFRof$60ml/minper1.73m2(568outcomes).611Certaintyofevidence.Theoverallcertaintyoftheevidenceforuricacid–loweringtherapyamongpeoplewithCKDandhyperuricemiaisverylow(seeSupplementaryTableS11150,612–614).ThecriticaloutcomeofdelayingprogressionofCKDwasaddressedby7RCTs.150,612,615–619The2largestRCTswereconsideredtohavealowriskofbias.615,616Thecertaintyoftheevidencewasdowngradedforinconsistency
becausetherewassubstantialstatisticalheterogeneitydetectedinthemeta-analysis(I2¼50%)andtheestimatedRRsrangedfrom0.05to2.96

---

### Chunk 26/30
**Article:** The Role of Urinary Nitrite in Predicting Bacterial Resistance in Urine Culture Analysis Among Patients With Uncomplicated Urinary Tract Infection (2022)
**Journal:** Cureus
**Section:** results | **Similarity:** 0.511

ian to choose
an effective empiric agent. There has been little research on whether the absence of urine nitrites reflects
resistance to popular antibiotics used to treat uncomplicated UTIs. Furthermore, the findings of the few
studies that have looked into this correlation are contradictory 
[15]
. This study aimed to investigate if nitrite
findings in urinalysis could be utilized to guide bacterial isolate and resistance when treating UTIs.
Materials And Methods
The urine samples of 59 adult outpatients with a mean age of 37 years and a diagnosis of UTI were reviewed.
The Urology Department at Tbilisi State Medical University (TSMU) the First University Clinic provided
samples throughout a six-month period (April-October 2020). Patients who were below the age of 18 years,
who did not have a urine culture sent to the laboratory, or those with the diagnosis of UTI but a negative
culture were excluded from the study.

---

### Chunk 27/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.511

logy
traceabletointernationalstandardsusingastandardreference
material.Thisiscurrentlynotthecase,andresultsmayvary
bygreaterthan40%betweenlaboratoriesdependingonthe
methodologyusedwithattendantimpactontheinterpreta-
tionofreportedresults.Thetypeofurinecollectionandtheanalyticalmethodinuenceresultinterpretation.Twenty-four-hoururinecol-lectionspresentproblemsintermsofcompletenessof
collection,specimenstorage,andtimingaccuracy.Therefore,
theassessmentofACRfromasinglevoidisacommonand
convenientclinicalpractice.TheACRaccountsforhydration
andhassimilardiagnosticperformanceto24-hoururine
AER.Thecollectionmethodshouldremainconsistent,pref-
erablyusingtherstmorningvoidspecimen.Ifspecimensarebeingstoredforfutureanalysis,carefulattentionmustbepaidtothestorageconditionstoavoid
degradationofalbuminleadingtoquanticationerror.There-portedeffectsoffrozenstorageonurinealbuminaresomewhat
inconsistent.Albuminisgenerallystableinurinestoredat2C–8Cfor7days.However,lossesofalbuminhavebeenreportedwhen

---

### Chunk 28/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.510

chmannLM.Movingtowardstandardizationofurinealbuminmeasurements.EJIFCC.2017;28:258–267.314.NationalInstituteofStandardsandTechnology.CerticationofStandardReferenceMaterial2925RecombinantHumanSerumAlbuminSolution(PrimaryReferenceCalibratorforUrineAlbumin)(Frozen).U.S.DepartmentofCommerce,NIST;2020.315.CarterJL,ParkerCT,StevensPE,etal.Biologicalvariationofplasmaand
urinarymarkersofacutekidneyinjuryinpatientswithchronickidneydisease.ClinChem.2016;62:876–883.316.NationalInstituteforHealthandCareExcellence.Point-of-carecreatininedevicestoassesskidneyfunctionbeforeCTimagingwithintravenous
contrast.NICEGuideline[NG37].NICE;2019.317.BatteA,MurphyKJ,NamazziR,etal.Evaluatingkidneyfunctionusingapoint-of-carecreatininetestinUgandanchildrenwithseveremalaria:aprospectivecohortstudy.BMCNephrol.2021;22:369.318.McTaggartMP,NewallRG,HirstJA,etal.Diagnosticaccuracyofpoint-of-
caretestsfordetectingalbuminuria:asystematicreviewandmeta-analysis.AnnInternMed.2014;160:550–557.319.AbitbolC,ZillerueloG,F

---

### Chunk 29/30
**Article:** The Role of Urinary Nitrite in Predicting Bacterial Resistance in Urine Culture Analysis Among Patients With Uncomplicated Urinary Tract Infection (2022)
**Journal:** Cureus
**Section:** discussion | **Similarity:** 0.510

BLE
 3: Frequency of isolated uropathogens
Discussion
The purpose of this study was to see if the presence of urine nitrite could help clinicians choose the most
appropriate empirical antibiotic before the bacteriological analysis findings. Commonly seen pathogens in
complicated UTIs which do not convert nitrates to nitrites include 
Enterococcus
, 
Pseudomonas
, and
Acinetobacter 
[16]
.
Weisz et al. stated that urine nitrite results are a useful sign of resistance to empiric therapy with first- and
third-generation cephalosporins 
[10]
. In our study, we subjected nitrite-positive and nitrite-negative groups
to second and third-generation cephalosporins including cefuroxime, cefotaxime, ceftriaxone, and
ceftazidime. In our study, only cefuroxime, cefotaxime, and ceftriaxone revealed a statistically significant
difference (p ≤ 0.05) among these four antibiotics.
Larson et al. compared the sensitivity of nitrite-positive and nitrite-negative cultures to TMP-SMX.

---

### Chunk 30/30
**Article:** The Role of Urinary Nitrite in Predicting Bacterial Resistance in Urine Culture Analysis Among Patients With Uncomplicated Urinary Tract Infection (2022)
**Journal:** Cureus
**Section:** discussion | **Similarity:** 0.509

resence of urine nitrite were
recorded. Resistance rates to the antibiotics were compared between the positive and negative nitrite
groups. Chi-squared test was used to perform the statistical analysis using Prism software version 9.3.1
(GraphPad Software, Inc., San Diego, California).
Results
We examined the correlation between the nitrite-positive and -negative groups with the resistance pattern
to ceftriaxone, trimethoprim/sulfamethoxazole (TMP-SMX), ampicillin-sulbactam, fosfomycin, amikacin,
doxycycline, cefuroxime, cefotaxime, ceftazidime, and nitrofurantoin.
A total of 59 outpatients with a mean age of 37 years met the inclusion criteria between April and October
2020. In the positive and negative nitrite groups, there were 23 and 36 patients, respectively. Three (17.6%)
of the 17 gram-positive organisms and 20 (62.5%) of the 42 gram-negative organisms yielded positive nitrite
results.

---

