# ScoreItem: Litíase

**ID:** `019bf31d-2ef0-79c2-8fc0-7dcd8dcb75bc`
**FullName:** Litíase (Histórico de doenças - Doenças crônicas (Questionar ativamente as doenças crônicas mais comuns, perguntar sobre duração, grau de controle e tratamentos utilizados no passado e atualmente) - Outras doenças renais)

**Preparation Metadata:**
- Quality Grade: **FAIR**
- Total Chunks: 30 de 9 artigos
- Avg Similarity: 0.491

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-79c2-8fc0-7dcd8dcb75bc`.**

```json
{
  "score_item_id": "019bf31d-2ef0-79c2-8fc0-7dcd8dcb75bc",
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

**ScoreItem:** Litíase (Histórico de doenças - Doenças crônicas (Questionar ativamente as doenças crônicas mais comuns, perguntar sobre duração, grau de controle e tratamentos utilizados no passado e atualmente) - Outras doenças renais)

**30 chunks de 9 artigos (avg similarity: 0.491)**

### Chunk 1/30
**Article:** Pathophysiology and Main Molecular Mechanisms of Urinary Stone Formation and Recurrence (2024)
**Journal:** International Journal of Molecular Sciences
**Section:** abstract | **Similarity:** 0.560

Urinary stone disease (nephrolithiasis) is a multifactorial condition affecting 12% of the global population with recurrence rates as high as 50% within 5 years after first stone onset. This comprehensive review examines the pathophysiology and molecular mechanisms underlying urinary stone formation and recurrence. The review covers crystal nucleation, aggregation, and retention processes, highlighting the role of urinary supersaturation, pH imbalances, and inhibitor deficiencies. Special attention is given to calcium oxalate and calcium phosphate stones (80% of cases), uric acid stones, cystine stones, and infection-related struvite crystals. Understanding these mechanisms is essential for developing targeted prevention and treatment strategies for this highly recurrent condition.

---

### Chunk 2/30
**Article:** Outpatient crystalluria: prevalence, crystal types, and associations with comorbidities and urinary tract infections at a provincial hospital (2025)
**Journal:** Iranian Journal of Microbiology
**Section:** abstract | **Similarity:** 0.556

Background and Objectives: Crystalluria refers to the occurrence of crystals in urine resulting from urinary supersaturation, which disrupts the balance between factors that promote and those that inhibit crystal formation in urine. This study examined crystalluria occurrence in 1,025 urine samples from outpatients at a provincial hospital. Results: 22.04% of samples showed crystalluria with mean patient age of 51.3 years. The most common crystal types were calcium oxalate (46.4%), uric acid (23.5%), urates (15.1%) and struvite (9.3%). Diabetes, kidney failure, prostatitis, and nephrotic syndrome were associated with crystal formation. The prevalence of urinary tract infections in patients with urinary crystals was 10.6%. Notably, struvite crystals correlated with bacterial infections. Conclusion: Monitoring crystals is vital for preventing kidney stone development and infection-related complications in susceptible populations.

---

### Chunk 3/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.518

arltrationrate[eGFRcr-cys])(1B).1.1.3EvaluationofchronicityPracticePoint1.1.3.1:Proofofchronicity(durationofaminimumof3months)canbeestablishedby:(i)reviewofpastmeasurements/estimationsofGFR;(ii)reviewofpastmeasurementsofalbuminuriaorproteinuriaandurinemicroscopicexaminations;(iii)imagingndingssuchasreducedkidneysizeandreductionincorticalthickness;(iv)kidneypathologicalndingssuchasbrosisandatrophy;(v)medicalhistory,especiallyconditionsknowntocauseorcontributetoCKD;(vi)repeatmeasurementswithinandbeyondthe3-monthpoint.PracticePoint1.1.3.2:DonotassumechronicitybaseduponasingleabnormallevelforeGFRandACR,asthendingcouldbetheresultofarecentacutekidneyinjury(AKI)eventoracutekidneydisease(AKD).PracticePoint1.1.3.3:ConsiderinitiationoftreatmentsforCKDatrstpresentationofdecreasedGFRorelevatedACRifCKDisdeemedlikelyduetopresenceofotherclinicalindicators.1.1.4Evaluationofcause
PracticePoint1.1.4.1:EstablishthecauseofCKDusingclinicalcontext,personalandfamilyhistory,socialandenvi-ronmentalfactor

---

### Chunk 4/30
**Article:** Fisiologia e Bioquímica do Sistema Imune II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.518

olicitar e registrar parâmetros essenciais (PCR-us, VHS); calcular índices estimáveis (HGI, TAIG); considerar complementares (TNF-α, IL-6, GPx, MDA, antioxidantes totais).
- [ ] 11. Monitorar trimestralmente HbA1c, glicemia de jejum e HGI; estabelecer metas individuais (ex.: reduzir de 6,1 para ~5,5; longo prazo ~5,3).
- [ ] 12. Calcular TG/HDL e integrar com lipidograma/SREBP1c/2; evitar combinação de gordura saturada com açúcar e excesso de saturadas em múltiplas refeições.
- [ ] 13. Avaliar ferro, ferritina, transferrina e saturação (20–50%; evitar <20%); interpretar com hepcidina/SREBP1c e quadro inflamatório.
- [ ] 14. Medir TNF-α (<8,1) e IL-6 (<3,4) para acompanhar atividade inflamatória; relacionar com obesidade inflamada.
- [ ] 15. Calcular HOMA-β (167–175) e HOMA-IR (<2,15); buscar glicemia 60–90 e insulina ~6–7.
- [ ] 16. Monitorar homocisteína (<7,9) e PCR; usar PCR-us; documentar crises (PCR >1.000) e conduzir manejo apropriado.
- [ ] 17.

---

### Chunk 5/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.508

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

### Chunk 6/30
**Article:** Dislipidemias II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.506

dade, inatividade física, pressão arterial e dislipidemia.
- A inflamação é indicada por marcadores como a Proteína C-Reativa (PCR), onde um valor de 5 já é considerado elevado, e a Interleucina 6 (exemplo de paciente com 8.45).
- A resposta anti-inflamatória é medida pela Interleucina 10, com um valor de corte de 3,5 (abaixo disso é um risco) e um exemplo de paciente com 6.44.
- A saúde do endotélio é avaliada pelo óxido nítrico, que deve estar em 8.8; um paciente com 4.8 já apresenta disfunção sistêmica.
- A LDL oxidada é um marcador crítico, com um limite saudável de 133, enquanto pacientes de alto risco podem apresentar valores extremos, como 1000.
**Achados Adicionais**
- Existem cinco parâmetros clássicos para definir a síndrome metabólica.
- O alvo de LDL para pacientes de alto risco, segundo as diretrizes atuais, é de 50.

---

### Chunk 7/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.497

tNZJMed.1992;22:334–337.338.GilbertRE,AkdenizA,JerumsG.Detectionofmicroalbuminuriaindiabeticpatientsbyurinarydipstick.DiabetesResClinPract.1997;35:57–60.339.GrazianiMS,GambaroG,MantovaniL,etal.Diagnosticaccuracyofa
reagentstripforassessingurinaryalbuminexcretioninthegeneralpopulation.NephrolDialTransplant.2008;24:1490–1494.340.GuyM,NewallR,BorzomatoJ,etal.Useofarst-lineurineprotein-to-creatinineratiostriptestonrandomurinestoruleoutproteinuriain
patientswithchronickidneydisease.NephrolDialTransplant.2008;24:1189–1193.341.GuyM,NewallR,BorzomatoJ,etal.Diagnosticaccuracyoftheurinary
albumin:creatinineratiodeterminedbytheCLINITEKMicroalbuminandDCA2000þfortherule-outofalbuminuriainchronickidneydisease.ClinChimActa.2008;399:54–58.342.HasslacherC,MullerP,SchlipfenbacherRL.ResultsofamulticentrestudyforthedeterminationofmicroalbuminuriawithMicral-Test.KlinischesLabor.1995;41:441–447.343.HodelNC,HamadA,ReitherK,etal.Comparisonoftwodifferent
semiquantitativeurinarydipsticktestswithal

---

### Chunk 8/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.497

lurinarytractdiseaseRecurrentkidneycalculiMultisystemdiseases/chronicinammatoryconditionsSystemiclupuserythematosus
Vasculitis
HIVIatrogenic(relatedtodrugtreatmentsandprocedures)Drug-inducednephrotoxicityandradiationnephritisFamilyhistoryorknowngeneticvariantassociatedwithCKDKidneyfailure,regardlessofidentiedcauseKidneydiseaserecognizedtobeassociatedwithgeneticabnormality(e.g.,PKD,APOL1-mediatedkidneydisease,andAlportsyndrome)GestationalconditionsPretermbirthSmallgestationalsizePre-eclampsia/eclampsiaOccupationalexposuresthatpromoteCKDriskCadmium,lead,andmercuryexposurePolycyclichydrocarbonsPesticidesAKD,acutekidneydisease;AKI,acutekidneyinjury;APOL1,apolipoproteinL1;CKD,chronickidneydisease;CKDu,chronickidneydiseaseofundeterminedorigin;PKD,polycystickidneydisease.

---

### Chunk 9/30
**Article:** Automated urine sediment analyzers underestimate the severity of hematuria in glomerular diseases (2021)
**Journal:** Scientific Reports
**Section:** results | **Similarity:** 0.496

and digital imaging in one system: A multicenter performance evaluation of the cobas 6500 urine work area. Pract. Lab. Med. 17, e00139 (2019). 30. aller, T. R. & Wang, L. P. Evaluation of asymptomatic microscopic hematuria in adults. Am. Fam. Phys. 60, 11431152 (1999). 31. Palsson, R., Srivastava, A. & Waikar, S. S. Performance of the automated urinalysis in diagnosis of proliferative glomerulonephritis. Kidney Int. Rep. 4, 723727 (2019). 32. Orlandi, P. F. et al. Hematuria as a risk factor for progression of chronic kidney disease and death: Findings from the chronic renal insuciency cohort (CRIC) study. BMC Nephrol. 19, 150 (2018). 33. Schulman, G. et al. Risk factors for progression of chronic kidney disease in the EPPIC trials and the eect of AST-120. Clin. Exp. Nephrol. 22, 299308 (2018). 34. Mashitani, T. et al.

---

### Chunk 10/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.495

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
**Section:** other | **Similarity:** 0.493

tospecialistkidneycareservicesinthecircumstanceslistedinFigure48.SpecialconsiderationsPediatricconsiderations.PracticePoint5.1.2:Referchildrenandadolescentstospecialistkidneycareservicesinthefollowingcircumstances:anACRof30mg/g(3mg/mmol)oraPCRof200mg/g(20mg/mmol)ormore,conrmedonarepeatrstmorningvoidsample,whenwellandnotduringmenstruation,persistenthematuria,anysustaineddecreaseineGFR,hypertension,kidneyoutowobstructionoranomaliesofthekidneyandurinarytract,knownorsuspectedCKD,orrecurrenturinarytractinfection.5.2SymptomsinCKD5.2.1Prevalenceandseverityofsymptoms[Norecommendationsandpracticepoints]5.2.2IdenticationandassessmentofsymptomsPracticePoint5.2.2.1:AskpeoplewithprogressiveCKDabouturemicsymptoms(e.g.,reducedappetite,nausea,andleveloffatigue/lethargy)ateachconsultationusingastandardizedvalidatedassessmentofuremicsymptomstool.

---

### Chunk 12/30
**Article:** Proteinuria (2023)
**Journal:** StatPearls Publishing
**Section:** abstract | **Similarity:** 0.491

Comprehensive clinical review of proteinuria covering definition, etiology, diagnosis, and management. Proteinuria serves as a marker of kidney damage assisting with diagnosis, prognosis, and therapy. Normal protein excretion is <150 mg/24 hours. Severity categories: nephritic range (150-3,500 mg/24h) and nephrotic range (>3,500 mg/24h). Urine dipstick provides semi-quantitative screening but has important limitations: false-positives with dehydration, alkaline urine, UTI; inability to detect positively charged proteins. UK CKD guidelines define significant proteinuria as UPCR >45 mg/mmol. Quantitative methods (spot UPCR preferred for convenience and reliability, or 24-hour urine collection for accuracy) are indicated when dipstick is positive or proteinuria >1 g/day is suspected. Early detection and quantification are crucial given associations with progressive renal disease and cardiovascular complications.

---

### Chunk 13/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** methods | **Similarity:** 0.488

alityinchronickidneydisease.AmJNephrol.2015;41:456–463.565.AllonM,ShanklinN.Effectofalbuteroltreatmentonsubsequentdialytic
potassiumremoval.AmJKidneyDis.1995;26:607–613.566.FosterES,JonesWJ,HayslettJP,etal.Roleofaldosteroneanddietarypotassiuminpotassiumadaptationinthedistalcolonoftherat.Gastroenterology.1985;88:41–46.567.GennariFJ,SegalAS.Hyperkalemia:anadaptiveresponseinchronic
renalinsufciency.KidneyInt.2002;62:1–9.568.SandleGI,GaigerE,TapsterS,etal.Evidenceforlargeintestinalcontrol
ofpotassiumhomoeostasisinuraemicpatientsundergoinglong-term
dialysis.ClinSci(Lond).1987;73:247–252.569.RastegarA.Clinicalmethods:thehistory,physical,andlaboratory
examinations.In:WalkerHK,HallWD,HurstJW,eds.SerumPotassium.3rded.Butterworth;1990:731.570.CooperLB,SavareseG,CarreroJJ,etal.Clinicalandresearchimplications
ofserumversusplasmapotassiummeasurements.EurJHeartFail.2019;21:536–537.571.MartinRS,PaneseS,VirginilloM,etal.Increasedsecretionofpotassiumintherectumofhumanswithchronicrenalfailu

---

### Chunk 14/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XIX (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.485

g/kg) e prática de exercícios de resistência para preservar massa muscular.
- [ ] 3. Todos os profissionais: Em doenças crônicas sem causa orgânica clara ou com má resposta ao tratamento, investigar ativamente traumas de infância, estresse crônico e questões emocionais não resolvidas como possível "causa primeira".
- [ ] 4. Terapeutas e psicólogos: Adotar "terapia de precisão", utilizando múltiplas ferramentas e combinando diferentes abordagens terapêuticas para personalizar o tratamento e focar em resultados mensuráveis, em vez de seguir uma única linha teórica por longos períodos.
- [ ] 5. Estudo pessoal: Pesquisar o conceito de "causa primeira" de Aristóteles para aprofundar a lógica de buscar a origem dos problemas.
- [ ] 6. Estudo pessoal: Ler o livro de Bruce Lipton sobre a conexão entre mente e doenças físicas.

---

## SOAP

> Data e Hora: 2025-11-17 16:33:53
> Paciente: 
> Diagnóstico:

## Histórico do Diagnóstico:
1. Histórico Médico: 
2.

---

### Chunk 15/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.484

ropriatetreatment.KeyinformationBalanceofbenetsandharms.Thebenetsofkidneybi-opsyintermsofdiagnosis,prognosis,andplanningappro-
priatetreatmentforboththepersonwithCKDandhealthcare
PhysicalexamNephrotoxicmedicationsSymptoms and signsof urinary tractabnormalitiesSymptoms and signsof systemic diseasesLaboratory tests, imaging, and tissue sample, such as:• Urinalysis and urine sediment• Urine albumin-to-creatinine ratio• Serologic tests
• Ultrasound• Kidney biopsy• Genetic testingMedicalhistorySocial andenvironmentalhistoryObtain careful family historyfor possible genetic causes,including family pedigree for CKD
Figure8|Evaluationofcauseofchronickidneydisease(CKD).

---

### Chunk 16/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.482

ersbladder,nuclearmedicinestudies,MRIAssesskidneystructure(i.e.,kidneyshape,size,symmetry,andevidenceofobstruction)forcysticdiseaseandreuxdisease.Evolvingroleofadditionaltechnologies(e.g.,3Dultrasound)KidneybiopsyUltrasound-guidedpercutaneousUsuallyexaminedbylightmicroscopy,immunouorescence,andelectronmicroscopy,and,insomesituations,mayincludemoleculardiagnostics
Usedforexactdiagnosis,planningtreatment,assessingactivityand
chronicityofdisease,andlikelihoodoftreatmentresponse;mayalsobe
usedtoassessgeneticdiseaseLaboratorytests:
serologic,urine
testsChemistryincludingacid-baseandelectrolytes,
serologictestssuchasanti-PLA2R,ANCA,anti-GBM
antibodies
Serum-freelightchains,serum,andurineprotein
electrophoresis/immunoxationUrinalysisandurinesedimentexaminationRefertoKDIGO2021ClinicalPracticeGuidelinefortheManagementofGlomerularDiseases22Increasingrecognitionoftheroleoflightchainsinkidneydiseaseevenin
theabsenceofmultiplemyeloma(monoclonalgammopathyofrenal
signicance[MGRS])98Presenceofpers

---

### Chunk 17/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.481

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
**Section:** other | **Similarity:** 0.480

specickidneyfunctionofglomerularltration.Usethemoregeneralterm“kidneyfunction(s)”whendealingwiththetotalityoffunctionsofthekidney.
PhysicalexamNephrotoxicmedicationsSymptoms and signsof urinary tractabnormalitiesSymptoms and signsof systemic diseasesLaboratory tests, imaging, and tissue sample, such as:• Urinalysis and urine sediment• Urine albumin-to-creatinine ratio• Serologic tests
• Ultrasound• Kidney biopsy• Genetic testingMedicalhistorySocial andenvironmentalhistoryObtain careful family historyfor possible genetic causes,including family pedigree for CKD
Figure8|Evaluationofcauseofchronickidneydisease(CKD).

---

### Chunk 19/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.479

3.ScheiJ,StefanssonVT,MathisenUD,etal.Residualassociationsof
inammatorymarkerswitheGFRafteraccountingformeasuredGFRinacommunity-basedcohortwithoutCKD.ClinJAmSocNephrol.2016;11:280–286.184.SjostromP,TidmanM,JonesI.Determinationoftheproductionrateand
non-renalclearanceofcystatinCandestimationoftheglomerular
ltrationratefromtheserumconcentrationofcystatinCinhumans.ScandJClinLabInvest.2005;65:111–124.185.XinC,XieJ,FanH,etal.AssociationbetweenserumcystatinCand
thyroiddiseases:asystematicreviewandmeta-analysis.FrontEndocrinol(Lausanne).2021;12:766516.186.AgarwalR,BillsJE,YigazuPM,etal.Assessmentofiothalamateplasma
clearance:durationofstudyaffectsqualityofGFR.ClinJAmSocNephrol.2009;4:77–85.187.ShahKF,StevensPE,LambEJ.Theinuenceofacooked-shmealonestimatedglomerularltrationrate.AnnClinBiochem.2020;57:182–185.188.InkerLA,LeveyAS,CoreshJ.Estimatedglomerularltrationratefromapanelofltrationmarkers-hopeforincreasedaccuracybeyondmeasuredglomerularltrationrate?AdvChronicKidneyDis.2018;

---

### Chunk 20/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.479

andthe“new”stagesbetterreecttheirriskassocia-tions.Forthatreason,whereavailable,cystatinCshouldbe
addedtocreatinineforthepurposeofestimatingGFRfor
CKDdiagnosisandstaging.1.1.3EvaluationofchronicityPracticePoint1.1.3.1:Proofofchronicity(durationofa
minimumof3months)canbeestablishedby:(i)reviewofpastmeasurements/estimationsofGFR;(ii)reviewofpastmeasurementsofalbuminuriaor
proteinuriaandurinemicroscopicexaminations;(iii)imagingndingssuchasreducedkidneysizeandreductionincorticalthickness;(iv)kidneypathologicalndingssuchasbrosisandatrophy;(v)medicalhistory,especiallyconditionsknownto
causeorcontributetoCKD;(vi)repeatmeasurementswithinandbeyondthe3-
monthpoint.PracticePoint1.1.3.2:Donotassumechronicitybased
uponasingleabnormallevelforeGFRandACR,asthendingcouldbetheresultofarecentacutekidneyinjury(AKI)eventoracutekidneydisease(AKD).PracticePoint1.1.3.3:ConsiderinitiationoftreatmentsforCKDatrstpresentationofdecreasedGFRorelevatedACRifCKDisdeemedlikelyduetopresenceofotherclinicalindic

---

### Chunk 21/30
**Article:** Emagrecimento - Parte VII (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.477

3.  **Exercício Físico:** Potencializa os resultados.
    4.  **Movimento e Relações Saudáveis:** Incluindo a necessidade de terapia.
    5.  **Conexão com a Natureza:** Contato com o ambiente natural para saúde mental e espiritual.
*   **Colaboração Multidisciplinar:** O emagrecimento eficaz exige a colaboração com um nutricionista. Os pacientes devem ser incentivados a investir nesse acompanhamento profissional.
## ❓ Perguntas
- [Inserir Pergunta/Confusão]
## 📚 Tarefas
- [ ] 1. Educar os pacientes sobre a adipogênese e a "memória corporal" para o ganho de peso, usando analogias como a do balão.
- [ ] 2. Solicitar o exame de Proteína C Reativa ultrassensível (PCR-us) como marcador de inflamação sistêmica, independentemente da especialidade.
- [ ] 3. Para pacientes com baixo metabolismo (especialmente mulheres), considerar uma estratégia inicial focada no ganho de massa muscular antes de focar na perda de peso.
- [ ] 4.

---

### Chunk 22/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.477

k forrecurrence afterkidney transplantationConditions for whichgenetic testing isrelevant for reproductivecounselingExamples:• GLA (Fabry)• AGXT (primaryhyperoxaluria (PH))• CoQ10 genes (SRNS)• CTNS (cystinosis)• Tubulopathies(Na+, K+ etc.)Conditions amenable to nonspecific renoprotective strategiesExample:• COL4A3/4/5 (Alport)and RAAS blockadeExample:
• Glomerular diseasedue to mutations inAlport genes(COL4A3/4/5)Examples:• (CFH/CFI/C3...): aHUS• (AGXT, GRHPR, HOGA): primary hyperoxaluria (PH)
• Adenine phosphoribo-  syltransferase deficiency (APRT)Conditions amenable to specific screening for extrarenal manifestationsExamples:• HNF1B: diabetes• PKD1/PKD2(ADPKD): intracranialaneurysms • FLCN: renal cellcarcinoma, etc.Example:
• Prenatal/preimplantationdiagnosis
Figure9|Actionablegenesinkidneydisease.Actionabilityreferstothepotentialforgeneticrestresultstoleadtospecicclinicalactionsfrompreventionortreatmentofacondition,supportedbyrecommendationsbasedonevide

---

### Chunk 23/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.476

Cause of CKD is uncertain• Hereditary kidney disease• Recurrent extensive nephrolithiasis• A >3%–5% 5-year risk of requiring KRT measuredusing a validated risk equation• eGFR <30 ml/min per 1.73 m2• A sustained fall in GFR of >20% or >30% in thosepeople initiating hemodynamically active therapies• Consistent nding of signicant albuminuria(ACR ≥300 mg/g [≥30 mg/mmol] or AER ≥300 mg/24 hours, approximately equivalent to PCR ≥500 mg/g [≥50 mg/mmol] or PER ≥500 mg/24 h) in combination with hematuria• ≥2-fold increase in albuminuria in people with signicant albuminuria undergoing monitoring• A consistent nding of ACR >700 mg/g [>70 mg/mmol]• Urinary red cell casts, RBC >20 per high power eld sustained and not readily explained• CKD and hypertension refractory to treatment≥4 antihypertensive agents• Persistent abnormalities of serum potassium• Acidosis• Anemia• Bone disease• Malnutrition
Figure48|Circumstancesforreferraltospecialistkidney

---

### Chunk 24/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.474

chronickidney
disease.AmJKidneyDis.2005;46:S1–S122.555.KovesdyCP,MatsushitaK,SangY,etal.Serumpotassiumandadverse
outcomesacrosstherangeofkidneyfunction:aCKDPrognosisConsortiummeta-analysis.EurHeartJ.2018;39:1535–1542.556.ClaseCM,CarreroJJ,EllisonDH,etal.Potassiumhomeostasisandmanagementofdyskalemiainkidneydiseases:conclusionsfroma
KidneyDisease:ImprovingGlobalOutcomes(KDIGO)ControversiesConference.KidneyInt.2020;97:42–61.557.LewisEJ,HunsickerLG,ClarkeWR,etal.Renoprotectiveeffectoftheangiotensin-receptorantagonistirbesartaninpatientswith
nephropathyduetotype2diabetes.NEnglJMed.2001;345:851–860.558.CollinsAJ,PittB,ReavenN,etal.Associationofserumpotassiumwithall-causemortalityinpatientswithandwithoutheartfailure,chronic
kidneydisease,and/ordiabetes.AmJNephrol.2017;46:213–221.559.KorgaonkarS,TileaA,GillespieBW,etal.Serumpotassiumand
outcomesinCKD:insightsfromtheRRI-CKDcohortstudy.ClinJAmSocNephrol.2010;5:762–769.560.EinhornLM,ZhanM,HsuVD,etal.Thefrequencyofhyperkalemiaand
itssi

---

### Chunk 25/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.474

potassiumbalancebyincreasingtheexcretionofpotassium571,572Medications:blockingtheRAASpathwayandothermedicationresultingintheinabilitytoexcreteexcessivepotassium(Table26)569,573DiurnalvariationinpotassiumexcretionwithmostexcretioninhumansoccurringclosetonoonCircadianexcretionofkidneyelectrolyteshavebeenwelldocumented.574ClinicalrelevanceisyettobeunderstoodNotethe0.24–0.73mmol/lvariationinKþvalueswithinindividualsovera24-hourperiodPlasmavs.serumpotassiumvaluesPotassiumvaluesdifferbetweenserumandplasmavalueswithserumvaluesbeingtypicallyhigher.Healthcareprovidersneedtobeawareoftherightreference
valuesforthesample570PostprandialhyperkalemiaAskidneyfunctiondeclinesinCKD,thereisacorrespondingdeclineintheabilityofthekidneystoincreasekaliuresispostprandially,eventuallybecoming
insufcienttomaintainexternalpotassiumbalance575CKD,chronickidneydisease;GFR,glomerularltrationrate;Kþ,potassium;RAAS,renin-angiotensin-aldosteronesystem.

---

### Chunk 26/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.472

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

### Chunk 27/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.472

e-basedequations.TransplInt.2012;25:527–536.162.JanusN,Launay-VacherV,ByloosE,etal.CancerandrenalinsufciencyresultsoftheBIRMAstudy.BrJCancer.2010;103:1815–1821.163.Launay-VacherV,JanusN,DerayG.Renalinsufciencyandcancertreatments.ESMOOpen.2016;1:e000091.164.NaSY,SungJY,ChangJH,etal.Chronickidneydiseaseincancerpatients:anindependentpredictorofcancer-specicmortality.AmJNephrol.2011;33:121–130.165.RosnerMH,JhaveriKD,McMahonBA,etal.Onconephrology:the
intersectionsbetweenthekidneyandcancer.CACancerJClin.2021;71:47–77.166.SoveriI,BergUB,BjorkJ,etal.MeasuringGFR:asystematicreview.AmJKidneyDis.2014;64:411–424.167.WhiteCA,AkbariA,AllenC,etal.Simultaneousglomerularltrationratedeterminationusinginulin,iohexol,and99mTc-DTPAdemonstratestheneedforcustomizedmeasurementprotocols.KidneyInt.2021;99:957–966.168.XieP,HuangJM,LiuXM,etal.(99m)Tc-DTPArenaldynamicimaging
methodmaybeunsuitabletobeusedasthereferencemethodin
investigatingthevalidityofCDK-EPIequationfordetermining
glomerularltrati

---

### Chunk 28/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.471

rshallSM,AlbertiKG.Screeningforearlydiabeticnephropathy.AnnClinBiochem.1986;23(Pt2):195–197.290.MillerWG,BrunsDE,HortinGL,etal.Currentissuesinmeasurementandreportingofurinaryalbuminexcretion.ClinChem.2009;55:24–38.291.ChitaliaVC,KothariJ,WellsEJ,etal.Cost-benetanalysisandpredictionof24-hourproteinuriafromthespoturineprotein-creatinineratio.ClinNephrol.2001;55:436–447.292.CoteAM,BrownMA,LamE,etal.Diagnosticaccuracyofurinaryspotprotein:creatinineratioforproteinuriainhypertensivepregnant
women:systematicreview.BMJ.2008;336:1003–1006.293.DysonEH,WillEJ,DavisonAM,etal.Useoftheurinaryproteincreatinine
indextoassessproteinuriainrenaltransplantpatients.NephrolDialTransplant.1992;7:450–452.294.Leanos-MirandaA,Marquez-AcostaJ,Romero-ArauzF,etal.Protein:
creatinineratioinrandomurinesamplesisareliablemarkerof
increased24-hourproteinexcretioninhospitalizedwomenwith
hypertensivedisordersofpregnancy.ClinChem.2007;53:1623–1628.295.LemannJJr,DoumasBT.Proteinuriainhealthanddiseaseassessedby

---

### Chunk 29/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.471

kHM,Begin-HeickN,AcharyaC,etal.Automateddeterminationof
urineandcerebrospinaluidproteinswithCoomassiebrilliantblueandtheAbbottABA-100.ClinBiochem.1980;13:81–83.280.MarshallT,WilliamsKM.Totalproteindeterminationinurine:eliminationofadifferentialresponsebetweentheCoomassieblueandpyrogallolredproteindye-bindingassays.ClinChem.2000;46:392–398.281.GinsbergJM,ChangBS,MatareseRA,etal.Useofsinglevoidedurine
samplestoestimatequantitativeproteinuria.NEnglJMed.1983;309:1543–1546.282.PriceCP,NewallRG,BoydJC.Useofprotein:creatinineratio
measurementsonrandomurinesamplesforpredictionofsignicantproteinuria:asystematicreview.ClinChem.2005;51:1577–1586.283.BeethamR,CattellWR.Proteinuria:pathophysiology,signicanceandrecommendationsformeasurementinclinicalpractice.AnnClinBiochem.1993;30(Pt5):425–434.284.KeaneWF,EknoyanG.Proteinuria,albuminuria,risk,assessment,detection,elimination(PARADE):apositionpaperoftheNational
KidneyFoundation.AmJKidneyDis.1999;33:1004–1010.285.ClaudiT,CooperJG.Compar

---

### Chunk 30/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.469

223

levelsofGFR,thusunderstandingpotassiumphysiologyanditsimpactingfactorsareimportantineffectivepatientcare.HyperkalemiainpeoplewithpreservedGFRislessprev-alent.Anacuteepisodeofhyperkalemiaisapotassiumresult
abovetheupperlimitofnormalthatisnotknowntobe
chronic.Atthecurrenttime,thereisnoconsensusonthe
magnitude,duration,andfrequencyofelevatedpotassium
valuesthatdenechronicity.556InadditiontodecreasedeGFR,otherriskfactorsforhyperkalemiaincludedhigherACRandpriordiabetes,hyperglycemia,constipation,
RASi,557andMRA.536NotethatSGLT2idonotappeartoincreaseserumpotassiumvalues.403,517StudieshavedemonstratedacontinuousU-shapedrela-tionshipbetweenserumpotassiumandall-causemortality
inarangeofdifferentpopulations(Figure31).555,558Ithasalsobeenassociatedwithworsekidneyprognosis.559Observationally,theriskofdeathfromthesamedegreeof
hyperkalemiaislowerinmoreadvancedCKDstages.560–564Thismaysuggestthatthereareadaptivemechanismsthat
renderbettertolerancetoelevatedlevelsofpotassiumin
circulation.561,5

---

