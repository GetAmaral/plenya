# ScoreItem: NT-proBNP (<50 anos)

**ID:** `019bf31d-2ef0-7d19-8b60-64765b6fc8f0`
**FullName:** NT-proBNP (<50 anos) (Exames - Laboratoriais)
**Unit:** pg/mL
**Age Min:** 0 anos
**Age Max:** 50 anos

**Preparation Metadata:**
- Quality Grade: **EXCELLENT**
- Total Chunks: 30 de 3 artigos
- Avg Similarity: 0.626

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-7d19-8b60-64765b6fc8f0`.**

```json
{
  "score_item_id": "019bf31d-2ef0-7d19-8b60-64765b6fc8f0",
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

**ScoreItem:** NT-proBNP (<50 anos) (Exames - Laboratoriais)
**Unidade:** pg/mL
**Faixa Etária:** 0-50 anos

**30 chunks de 3 artigos (avg similarity: 0.626)**

### Chunk 1/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.717

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

### Chunk 2/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.674

eAssociationexaminesthepracticalusesofN-terminalpro-B-typenatriureticpeptide(NT-proBNP)invariousclinicalscenarios.TheconcentrationsofNT-proBNPvaryaccordingtothepatientproleandtheclinicalscenario,thereforevaluesshouldbeinterpreted
withcautiontoensureappropriatediagnosis.Validatedcut-pointsareprovidedtoruleinorruleoutacuteheartfailureintheemergencydepartmentandtodiagnosedenovoheartfailureintheoutpatientsetting.Wealsocointheconceptof‘heartstress’whenNT-proBNPlevelsareelevatedinanasymptomaticpatientwithriskfactorsforheartfailure(i.e.diabetes,hypertension,coronaryarterydisease),
underlyingthedevelopmentofcardiacdysfunctionandfurtherincreasedrisk.Weproposeasimpleacronymforhealthcareprofessionalsandpatients,FIND-HF,whichservesasaprompttoconsiderheartfailure:Fatigue,Increasedwateraccumulation,Natriureticpeptidetesting,
*Correspondingauthor.DepartmentofMedicine,UAB,Head,HeartInstitute.HospitalUniversitariGermansTriasiPujol,CarreteradelCanyets/n,08916Badalona,Spain.Email:abayesgenis@gmail.c

---

### Chunk 3/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.673

from125to250pg/ml,whileforthoseover75years,itextendsfrom125to500pg/ml.Itiscrucialtoconductathoroughevaluationofpatientswithinthegreyzone,consideringfactorssuchasobesity,race-basedvariations,andongoingtreatment(asmanypatientswithahistory
ofhypertensionmayalreadybeondiuretics,renin–angiotensin
systeminhibitors,ormineralocorticoidreceptorantagonists).Intheoutpatientsetting,theextentofelevationinNPconcen-trationsatthetimeofheartfailurediagnosisiscloselylinkedtothe
riskofsubsequenthospitalizationandmortality.16Asaresult,therehasbeenasuggestiontoutilizeNT-proBNPconcentrationsatthe
timeofacommunity-basedheartfailurediagnosisasatriagingtooltoprioritizeaccesstoexpediteddiagnosticechocardiogra-phyandtosetupafollow-upplanforindividualswiththehighest
short-termriskofadverseevents.Theauthorsofthisconsensus
documentalignwiththeNICEguidelinesonchronicheartfailure,
whichrecommendacut-offvalueofNT-proBNP>2000pg/ml.40InananalysisofprimarycaredatafromEngland,anNT-proBNPvalue
of>2000pg/mlwasassociatedwi

---

### Chunk 4/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.667

edvariationsinNT-proBNPconcentrations,suggesting
thatasinglerule-inthresholdintheoutpatientsettingcould
resultinunnecessaryreferralsforadditionalinvestigationssuch
asechocardiography.41–45Additionally,comorbiditiessuchasobesityandrenaldysfunctionmayalsoaffectNT-proBNPrule-in
concentrations.31Inordertoensuresimplicityandeaseofuse,thisconsensusdocumentadoptsacompromiseapproachbyuti-
lizingage-specicrule-incut-pointsexclusively.Thesecut-points
aredeterminedbasedonthe95thpercentileoftheGeneration
Scotlandcohort,41followingthesameagestrataasthoseemployedintheED.Theage-adjustedrule-incut-pointsintheoutpatientset-
tinghavebeenestablished:125pg/mlforpatientsunder50years,250pg/mlforpatientsaged50-75years,and500pg/mlforpatientsover75years.IfNT-proBNPconcentrationssurpass©2023EuropeanSocietyofCardiology.

1894A.Bayes-Genisetal.

---

### Chunk 5/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.659

yJ,Bayes-GenisA,Ordonez-LlanosJ,Santalo-BelM,etal.NT-proBNPtestingfordiagnosisandshort-termprognosisinacutedestabilizedheartfailure:Aninternationalpooledanalysisof1256patients:TheInternationalCollaborativeofNT-proBNPstudy.EurHeartJ2006;27:330–337.https://doi.org/10.1093/eurheartj/ehi63128.JanuzziJLJr,Chen-TournouxAA,ChristensonRH,DorosG,HollanderJE,LevyPD,etal.;ICON-RELOADEDInvestigators.N-terminalpro-B-typenatriureticpeptideintheemergencydepartment:TheICON-RELOADEDstudy.JAmCollCardiol2018;71:1191–1200.https://doi.org/10.1016/j.jacc.2018.01.02129.Bayes-GenisA,Lloyd-JonesDM,vanKimmenadeRR,LainchburyJG,RichardsAM,Ordoñez-LlanosJ,etal.Effectofbodymassindexondiagnosticandprognosticusefulnessofamino-terminalpro-brainnatriureticpeptideinpatientswithacutedyspnea.ArchInternMed2007;167:400–407.https://doi.org/10.1001/archinte.167.4.400...........................................................................................................................................................

---

### Chunk 6/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.647

Forpatientsaged50–74years,acut-pointof150pg/ml,andforthoseover75yearsofage,acut-pointof300pg/mlisproposed
(Figure3).Increasedconcentrationsindicatethelikelihoodofheartstress,promptingcarefulre-evaluationofthepatientforproblemssuchasatrialbrillationandchronickidneydisease,lifestyleadvice
ScreeningforHeart Stressin Asymptomatic patients with T2D(or otherrisk factors forCVD)NT-proBNPHeartStressVery UnlikelyRepeat NT-proBNP in one yearHeartStressNotLikelyRepeat NT-proBNP in 6 monthsHeartStressLikelyArrange EchocardiographyAssessment by Heart Failure team if cardiac dysfunction presentRule-out≤ 50pg/mLGreyzoneAge-adjustedRule-in<50y: ≥ 75pg/mL50-74y: ≥ 150pg/mL≥75y: ≥ 300pg/mL
Figure3NT-proBNPfordiagnosisofheartstressinasymp-tomaticpatientswithdiabetes.CVD,cardiovasculardisease;
NT-proBNP,N-terminalpro-B-typenatriureticpeptide;T2D,type2diabetes.......................................................................................................................................

---

### Chunk 7/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.644

de,Serbia;19CardioPulmonaryDepartment,IRCCSSanRaffaele,Rome,Italy;20ExerciseScienceandMedicine,SanRaffaeleOpenUniversity,Rome,Italy;21CardiovascularDivision,BrighamandWomen’sHospital,HarvardMedicalSchool,Boston,MA,USA;22CardiologyandCardiacCatheterizationLaboratory,Cardio-ThoracicDepartment,CivilHospitals;DepartmentofMedicalandSurgicalSpecialties,RadiologicalSciences,andPublicHealth,UniversityofBrescia,Brescia,Italy;and23IRCCSSanRaffaele,Rome,ItalyReceived19June2023;revised13September2023;accepted13September2023;onlinepublish-ahead-of-print26September2023
Diagnosingheartfailureisoftendifcultduetothenon-specicnatureofsymptoms,whichcanbecausedbyarangeofmedicalconditions.Natriureticpeptides(NPs)havebeenrecognizedasimportantbiomarkersfordiagnosingheartfailure.ThisdocumentfromtheHeartFailureAssociationexaminesthepracticalusesofN-terminalpro-B-typenatriureticpeptide(NT-proBNP)invariousclinicalscenarios.TheconcentrationsofNT-proBNPvaryaccordingtothepatientproleandtheclinicalscenario,ther

---

### Chunk 8/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.643

stigationstodiagnoseheartfailure(e.g.measurementofNPconcentrations).16AlthoughmeasurementofNPsasadiagnostictoolhasbeenendorsedinguidelinesformorethan20years,their
usehasonlyincreasedslightlyovertime.38Asignicantdecitper-sistsintheuptakeofNPtestingtodiagnoseorexcludeheartfailure
inthecommunity39;thisdocumentshouldserveasacalltoaction.Asinglerule-outcut-pointof125pg/mlforNT-proBNPhasbeenconsistentlyrecommendedbyguidelinesintheoutpatient
settingtoexcludeadiagnosisofheartfailure,regardlessofthe
patient’sage(Figure2).1,3Incontrast,theoutpatientsettinglackswell-denedrule-invaluesforNT-proBNP.TheguidelinesfromtheUnitedKingdom
NationalInstituteforHealthandCareExcellence(NICE)pro-
poseasinglerule-inthresholdof400pg/mlforNT-proBNP.40However,populationstudieshaverevealedsignicantage-and
sex-basedvariationsinNT-proBNPconcentrations,suggesting
thatasinglerule-inthresholdintheoutpatientsettingcould
resultinunnecessaryreferralsforadditionalinvestigationssuch
asechocardiography.41–45Additional

---

### Chunk 9/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.636

ureticpeptidetestingforheartfailureinUKprimarycare:Acohortstudy.EurHeartJ2021;43:881–891.https://doi.org/10.1093/eurheartj/ehab78139.Bayes-GenisA,CoatsAJS.’PeptideforLife’inprimarycare:Workinprogress.EurHeartJ2022;43:892–894.https://doi.org/10.1093/eurheartj/ehab82940.NationalInstituteforHealthandCareExcellence.Heartfailureguidance.www.nice.org.uk/Guidance/Conditions-and-diseases/Cardiovascular-conditions/
Heart-failure.Accessed17September2023.41.WelshP,CampbellRT,MooneyL,KimenaiDM,HaywardC,CampbellA,etal.ReferencerangesforNT-proBNP(N-terminalpro-B-typenatriureticpeptide)andriskfactorsforhigherNT-proBNPconcentrationsinalargegeneralpopulationcohort.CircHeartFail2022;15:e009427.https://doi.org/10.1161/CIRCHEARTFAILURE.121.00942742.MuS,Echouffo-TcheuguiJB,NdumeleCE,CoreshJ,JuraschekS,BradyT,etal.NT-proBNPreferenceintervalsinhealthyU.S.children,adolescents,andadults.

---

### Chunk 10/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.629

ruleinofNT-proBNPlevelswereinitiallyestab-
lishedbytheInternationalCollaborativeofNT-proBNP(ICON)
studyin2006.27Thesecut-pointswerefurthervalidatedintheICON-RELOADEDstudypublishedin2018,whichincludedamorecontemporarypopulationofolderandmorecomorbid
patients.28Theconsistencyintheuseofthesecut-pointsacrossdifferentstudieshighlightstheirutilityinclinicalpractice.33The‘greyzone’inNT-proBNPreferstoarangeofplasmaconcentrationsthatliebetweentherule-outandrule-invalues.InpatientswithacutedyspnoeaintheED,thegreyzoneforNT-proBNPconcentrationsrangesbetweentheage-independent
rule-outvalueandtheage-adjustedrule-invalue.Forexample,
amongindividualsunder50years,thegreyzoneisbetween300
and450pg/ml,forthoseaged50and75years,itisbetween300
and900pg/ml,andforthoseover75years,itisbetween300and
1800pg/ml.Patientsfallingwithinthegreyzonerequirefurtherdiagnostictesting,suchasechocardiographyorcardiacimaging,
andconsiderationofotherclinicalfactors,todeterminewhether
theyhaveheartfailureoranotherunderlyingc

---

### Chunk 11/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.628

congestiveheartfailureinanurgent-caresetting.JAmCollCardiol2001;37:379–385.https://doi.org/10.1016/s0735-1097(00)01156-619.CollinsSP,LindsellCJ,StorrowAB,AbrahamWT;ADHEREScienticAdvisoryCommittee,InvestigatorsandStudyGroup.Prevalenceofnegativechestradio-graphyresultsintheemergencydepartmentpatientwithdecompensatedheart
failure.AnnEmergMed2006;47:13–18.https://doi.org/10.1016/j.annemergmed.2005.04.00320.Bayés-GenísA,Santaló-BelM,Zapico-MuñizE,LópezL,CotesC,BellidoJ,etal.N-terminalprobrainnatriureticpeptide(NT-proBNP)intheemergencydiagnosisandin-hospitalmonitoringofpatientswithdyspnoeaandventriculardysfunction.EurJHeartFail2004;6:301–308.https://doi.org/10.1016/j.ejheart.2003.12.01321.LamLL,CameronPA,SchneiderHG,AbramsonMJ,MüllerC,KrumH.Meta-analysis:EffectofB-typenatriureticpeptidetestingonclinicaloutcomes
inpatientswithacutedyspneaintheemergencysetting.AnnInternMed2010;153:728–735.https://doi.org/10.7326/0003-4819-153-11-201012070-0000622.MoeGW,HowlettJ,JanuzziJL,ZowallH

---

### Chunk 12/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.628

tientsfallingwithinthegreyzonerequirefurtherdiagnostictesting,suchasechocardiographyorcardiacimaging,
andconsiderationofotherclinicalfactors,todeterminewhether
theyhaveheartfailureoranotherunderlyingcondition.34........................................................................................................................................................................PatientspresentingtotheEDwithveryhighNT-proBNPcon-centrations,particularlythoseabove5000pg/ml,haveapoorprog-
nosis.27,35PatientswithNT-proBNPconcentrationsabovethisthresholdrequirehospitaladmission,oftenincriticalcare,urgent
investigationandclosemonitoring.Treatmentofcongestion,usu-
allyadministeredintravenously,isnecessary(Figure1).NT-proBNPisalsoavaluablediagnostictoolforidentifyingworseningheartfailureinothercaresettings.Toidentifypatients
withknownandtreatedheartfailure,whoarefreefromcon-
gestion,weutilizetheterm‘dry’NT-proBNP.Consequently,an
NT-proBNPincreasebymorethan25%comparedtothe‘dry’
NT-proBNPva

---

### Chunk 13/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.626

tidetestingonclinicaloutcomes
inpatientswithacutedyspneaintheemergencysetting.AnnInternMed2010;153:728–735.https://doi.org/10.7326/0003-4819-153-11-201012070-0000622.MoeGW,HowlettJ,JanuzziJL,ZowallH;CanadianMulticenterImprovedManagementofPatientswithCongestiveHeartFailure(IMPROVE-CHF)Study
Investigators.N-terminalpro-B-typenatriureticpeptidetestingimprovesthemanagementofpatientswithsuspectedacuteheartfailure:PrimaryresultsoftheCanadianprospectiverandomizedmulticenterIMPROVE-CHFstudy.

---

### Chunk 14/30
**Article:** Dislipidemias II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.626

dade, inatividade física, pressão arterial e dislipidemia.
- A inflamação é indicada por marcadores como a Proteína C-Reativa (PCR), onde um valor de 5 já é considerado elevado, e a Interleucina 6 (exemplo de paciente com 8.45).
- A resposta anti-inflamatória é medida pela Interleucina 10, com um valor de corte de 3,5 (abaixo disso é um risco) e um exemplo de paciente com 6.44.
- A saúde do endotélio é avaliada pelo óxido nítrico, que deve estar em 8.8; um paciente com 4.8 já apresenta disfunção sistêmica.
- A LDL oxidada é um marcador crítico, com um limite saudável de 133, enquanto pacientes de alto risco podem apresentar valores extremos, como 1000.
**Achados Adicionais**
- Existem cinco parâmetros clássicos para definir a síndrome metabólica.
- O alvo de LDL para pacientes de alto risco, segundo as diretrizes atuais, é de 50.

---

### Chunk 15/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.625

accausesofdyspnoea.ForrulinginheartfailureintheED,age-adjustedcut-pointshavebeenestablished:450pg/mlforpatientsunder50years,©2023EuropeanSocietyofCardiology.

---

### Chunk 16/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.623

   5,000pg/mL
Figure1NT-proBNPfordiagnosisofheartfailureintheemer-gencydepartment.AF,atrialbrillation/utter;BMI,bodymass
index;CXR,chestx-ray;ECGelectrocardiogram;LVEF,leftven-tricularejectionfraction;NT-proBNP,N-terminalpro-B-typenatriureticpeptide.900pg/mlforpatientsaged50–75years,and1800pg/mlforpatientsover75years27,28(onlinesupplementaryTableS1).IfNT-proBNPconcentrationsexceedthesecut-points,acuteheart
failureislikely,andadmission,furtherinvestigation,andtreatment
forheartfailureshouldbeadvised.Theage-adjustedrule-invalues
forNT-proBNPintheEDdonotnecessitateadditionaladjust-
mentsforfactorsthatcaninuenceNPlevels,suchasrenaldys-
function(lowestimatedglomerularltrationrate[eGFR]),obesity,
diabetes,oratrialbrillation.29–32Thecut-pointsfortheage-independentruleoutandage-adjustedruleinofNT-proBNPlevelswereinitiallyestab-
lishedbytheInternationalCollaborativeofNT-proBNP(ICON)
studyin2006.27Thesecut-pointswerefurthervalidatedintheICON-RELOADEDstudypublishedin2018,whichincludeda

---

### Chunk 17/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.622

T-proBNP,N-terminalpro-B-typenatriureticpeptide;T2D,type2diabetes.........................................................................................................................................................................(e.g.dietarysaltintake,exercise,smoking,etc.)andthetreatmentofriskfactorssuchashypertension.NT-proBNPconcentrationsmaybere-evaluatedwithinthefollowing6–12monthstodeter-minetheresponsetoanyinterventionandtheneedforfurtherinvestigation(Figure3).Thisadvice,basedonaconsensusdecision,requiresconrmationbyprospectivestudies.Itisimportanttoval-idatetheheartstressalgorithmusingclinicaltrialdata;conductingpost-hocvalidationshouldbefeasiblefromexistingtrialdataor
registries(suchastheUKBiobank).FIND-HF–TheHFAacronymforearlydiagnosisofheartfailureTopromoteearlydiagnosisofheartfailureandassisthealthcareprofessionalsandpatients,wesuggestthemnemonicacronymFIND-HF(Fatigue,Increasedwateraccumulation,Natriureticpep-tidetesting,andDyspnoea-HeartFailure),whichservesasa
re

---

### Chunk 18/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.619

nitionofheartfailure6alsounder-scorestheimportanceofmeasuringNPsfordiagnosingheart
failure.ThisparadigmshiftrecognizesNPsasakeycomponentintheearlydetectionofheartfailure,benetingbothspecialistsandnon-specialistsalike.ThisclinicalconsensusstatementwillfocusonNT-proBNP,asitisthemostutilizedpeptidefordiagnosingandmanagingheartfailureinEurope.7WhilethereisastoichiometricreleaseofBNPandNT-proBNP,theirhalf-livesdiffer.Theestimatedhalf-lifefor
BNPis21min,whereasforNT-proBNP,itisextendedtoaround70min.Consequently,concentrationsofNT-proBNParehigherthanthoseofBNP.8–10StabilityatroomtemperaturefacilitatesthehandlingofspecimensforNT-proBNPmeasurementinusuallybusyclinicallaboratories;inthisregard,NT-proBNPisamore
convenientmoleculetoworkwiththanBNP,whosestabilityisdependentonthespecicassayandislargelyunstableatroomtemperature.11Moreover,itisnotaffectedbytreatmentsthatalterdegradationofBNP(e.g.sacubitril/valsartan).12,13OthersmaywishtodevelopdedicatedalgorithmsforBNPandMR-proANP.Toaddressthevar

---

### Chunk 19/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.615

ulnessofaminoterminalpro-brainnatriureticpeptidetestingforthediagnosticandprognosticevaluationofdyspneicpatientswithdiabetesmellitusseenintheemergencydepartment(fromthePRIDEstudy).AmJCardiol2007;100:1336–1340.https://doi.org/10.1016/j.amjcard.2007.06.02033.LeeKK,DoudesisD,AnwarM,AstengoF,Chenevier-GobeauxC,ClaessensYE,etal.CoDE-HFinvestigators.Developmentandvalidationofadecisionsupporttoolforthediagnosisofacuteheartfailure:Systematicreview,meta-analysis,andmodellingstudy.BMJ2022;377:e068424.https://doi.org/10.1136/bmj-2021-06842434.vanKimmenadeRR,PintoYM,JanuzziJLJr.Importanceandinterpretationofintermediate(grayzone)amino-terminalpro-B-typenatriureticpeptideconcen-trations.AmJCardiol2008;101:39–42.https://doi.org/10.1016/j.amjcard.2007.11.01835.SalahK,KokWE,EurlingsLW,BettencourtP,PimentaJM,MetraM,etal.Anoveldischargeriskmodelforpatientshospitalisedforacutedecompensated
heartfailureincorporatingN-terminalpro-B-typenatriureticpeptidelevels:AEuropeancoLlaborationonAcutedecompeNsatedH

---

### Chunk 20/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.615

3–815.https://doi.org/10.1002/ejhf.147124.MuellerC,ScholerA,Laule-KilianK,MartinaB,SchindlerC,BuserP,etal.UseofB-typenatriureticpeptideintheevaluationandmanagementofacutedyspnea.NEnglJMed2004;350:647–654.https://doi.org/10.1056/NEJMoa03168125.SiebertU,MilevS,ZouD,LitkiewiczM,GagginHK,TirapelleL,etal.EconomicevaluationofanN-terminalproB-typenatriureticpeptide-supporteddiagnos-ticstrategyamongdyspneicpatientssuspectedofacuteheartfailureinthe
emergencydepartment.AmJCardiol2021;147:61–69.https://doi.org/10.1016/j.amjcard.2021.01.03626.MuellerC,Laule-KilianK,SchindlerC,KlimaT,FranaB,RodriguezD,etal.Cost-effectivenessofB-typenatriureticpeptidetestinginpatientswithacutedys-pnea.ArchInternMed2006;166:1081–1087.https://doi.org/10.1001/archinte.166.10.108127.JanuzziJL,vanKimmenadeR,LainchburyJ,Bayes-GenisA,Ordonez-LlanosJ,Santalo-BelM,etal.NT-proBNPtestingfordiagnosisandshort-termprognosisinacutedestabilizedheartfailure:Aninternationalpooledanalysisof1256patients:TheInternationalCollabor

---

### Chunk 21/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.610

ttings.Toidentifypatients
withknownandtreatedheartfailure,whoarefreefromcon-
gestion,weutilizetheterm‘dry’NT-proBNP.Consequently,an
NT-proBNPincreasebymorethan25%comparedtothe‘dry’
NT-proBNPvaluesuggestsworseningheartfailure.Acompre-hensiveclinicalevaluationisessentialtodiagnoseandmanagesuspectedworseningheartfailure,aselevatedNT-proBNP
concentrationscanbeobservedinotherconditions.15,36,37NT-proBNPforrulingoutandrulingindenovoheartfailureintheoutpatientsettingThediagnosisofheartfailureintheoutpatientsettingcanbechallengingduetothewiderangeofnon-specicsymptomswhich
patientscanpresenttotheirgeneralpractitioner(GP).Many
patientswhopresenttohospitalandarediagnosedwithheart
failureforthersttimehavingpreviouslypresentedtotheirGP
withsymptomssuggestiveofheartfailureyetdidnotreceive
investigationstodiagnoseheartfailure(e.g.measurementofNPconcentrations).16AlthoughmeasurementofNPsasadiagnostictoolhasbeenendorsedinguidelinesformorethan20years,their
usehasonlyincreasedslightlyovertime.3

---

### Chunk 22/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.609

cedoniareportedthelowest
use(none),followedbytheRussianFederation(0.02hospitalspermillionpeople).Certainly,differencesinlocalhealthcaresystemorganization,delivery,localcosts,andfundingcancontributeto
discrepanciesintheuseofNT-proBNPinED.Practicalalgorithms
fortheuseofNT-proBNPinEDareneededtoensureconsistencyandstandardizationinthediagnosisofheartfailure.IntheED,higherNPcut-pointsareutilizedcomparedtotheoutpatientsetting,reectingthepresenceof‘wet’NPscharac-terizedbyacutelyelevatedBNP/NT-proBNPlevelsduetoseverecongestionleadingtogreatermyocardialstretchingandincreased
NPsecretion.Asinglecut-pointof300pg/mlNT-proBNPisconsideredtoruleoutthediagnosisofheartfailure(Figure1),regardlessofthepatient’sage.27,28WhenplasmaNT-proBNPisbelowthisvalue,cliniciansshouldevaluatethepatientfor
non-cardiaccausesofdyspnoea.ForrulinginheartfailureintheED,age-adjustedcut-pointshavebeenestablished:450pg/mlforpatientsunder50years,©2023EuropeanSocietyofCardiology.

---

### Chunk 23/30
**Article:** Cardiologia VIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.605

# Cardiologia VIII

**Source:** https://web.plaud.ai/share/43b41764908850761::YXdzOnVzLXdlc3QtMg

---

## Reasoning Summary

## Análise de Exames, Fatores de Risco e Estratégias Terapêuticas em Cardiologia
### 1. Interpretação de Exames e Marcadores de Risco Cardiovascular
- **Princípio da Probabilidade e Contexto Clínico:** A análise de exames laboratoriais deve seguir o princípio de que a medicina é uma "ciência da probabilidade". Os valores devem ser interpretados dentro do contexto clínico do paciente (história, exame físico, idade) e não como números isolados a serem "corrigidos". É crucial não se fixar em atingir valores "ótimos" em todos os exames, pois o exame é um desfecho substituto e a prioridade é o paciente como um todo.

---

### Chunk 24/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.604

EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036
PracticalalgorithmsforearlydiagnosisofheartfailureandheartstressusingNT-proBNP:AclinicalconsensusstatementfromtheHeartFailureAssociationoftheESCAntoniBayes-Genis1*
,KieranF.Docherty2,MarkC.Petrie2,JamesL.Januzzi3,ChristianMueller4,LisaAnderson5,BiykemBozkurt6,JavedButler7,OvidiuChioncel8,JohnG.F.Cleland9,RuxandraChristodorescu10,StefanoDelPrato11,FinnGustafsson12,CarolynS.P.Lam13,BrendaMoura14,15,RodicaPop-Busui16,PetarSeferovic17,18,MaurizioVolterrani19,20,MuthiahVaduganathan21,MarcoMetra22,andGiuseppeRosano231HeartInstitute,HospitalUnbiversitariGermasnTriasiPujol,UniversitatAutonomadeBarcelona,CIBERCV,Barcelona,Spain;2SchoolofCardiovascularandMetabolicHealth,UniversityofGlasgow,Glasgow,UK;3CardiologyDivision,MassachusettsGeneralHospital,HarvardMedicalSchool,Boston,MA,USA;4DepartmentofCardiologyandCardiovascularResearchInstituteBasel(CRIB),UniversityHospitalBasel,UniversityofBasel,Basel,Switzerland;5Ca

---

### Chunk 25/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.597

tteronbaselineECG.46Thefol-lowingsuggestedmodicationsforcomorbidconditionsarebased
moreonexpertopinionratherthanonstrongevidenceandshould
berenedasmoreinformationbecomesavailable.WheneGFRis<30ml/min/1.73m2,thecut-pointforNT-proBNPshouldbeincreasedby35%;foreGFRbetween30and45ml/min/1.73m2by25%andforeGFR45-60ml/min/1.73m2by15%.WhenBMIisbetween30and35kg/m2,theNT-proBNPcut-offshouldbereducedby25%;forBMIbetween35and40kg/m2by30%andover40kg/m2by40%.Foratrialbrillationorutter,theNT-proBNPcut-pointshouldbeincreasedby50%whentheventricularrate
is90bpmatthetimeoftheblooddraworby100%whentheventricularrateis>90bpm.Furtherinvestigationisrequiredtoascertainwhichofthetwoapproaches,thesimplerage-adjustedorthemoresophisticatedfullyadjustedrule-incut-points,offeragreaterreductioninunnec-
essaryreferralsandechocardiograms.........................................................................................................................................................................

---

### Chunk 26/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.594

onwithmajorandadverseeffectsonqualityoflife,survival,andhealthcarecosts.1Asheartfailuretreatmentsreducehospitalizationsanddeath,earlydiagnosisisessential.ThemostrecentEuropeanSocietyofCardiology(ESC)guidelinesonheartfailurerecommendmeasuringconcentrationsofN-terminalpro-B-typenatriureticpeptide(NT-proBNP),B-typenatriuretic
peptide(BNP),ormid-regionalpro-atrialnatriureticpeptide(MR-proANP)forthediagnosticevaluationofheartfailure.1TheveryrstESCguidelineonheartfailurepublishedin1995explicitlyrecommendedtheuseofnatriureticpeptides(NPs)toruleoutadiagnosisofheartfailure.2However,themostrecentguidelinesfromtheESC,theAmericanHeartAssociation/AmericanCollegeofCardiologyandtheAmericanDiabetesAssociationnowsuggestthatNPsshouldalsobeconsideredtodiagnosis(i.e.rulein)heartfailure.1,3–5Theuniversaldenitionofheartfailure6alsounder-scorestheimportanceofmeasuringNPsfordiagnosingheart
failure.ThisparadigmshiftrecognizesNPsasakeycomponentintheearlydetectionofheartfailure,benetingbothspecialistsandnon

---

### Chunk 27/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.593

osisofheartfailureandheartstress.Thisdocument..............................................................................................................................................isanupdatetothe2019HFAdocumentof‘Cardiologypracticalguidanceontheuseofnatriureticpeptideconcentrations’.15NT-proBNPforrulingoutandrulinginheartfailureintheemergencydepartmentApproximately60–80%ofheartfailurediagnosesaremadeforthersttimeintheemergencydepartment(ED)asreportedina
NationalHealthServiceaudit.16Itisimportanttonotethatthesestatisticsmaynotnecessarilymirrorthesituationindifferentglobalregions.Whenevaluatingpatientswithpossibleheartfailure,clini-
cianuncertaintyisassociatedwithincreasedmorbidityandmortal-ityandprolongshospitallengthofstay.17StudieshaveconsistentlydemonstratedthattheuseofNPtestingintheEDresultsinmore
accurateandtimelydiagnosisofheartfailure,leadingtofasterinitia-
tionoflife-savingtherapiesandearlierdischarge,withresultingclin-icalandeconomicbenets.18–24TheuseofNT-proBN

---

### Chunk 28/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.589

ntaJM,MetraM,etal.Anoveldischargeriskmodelforpatientshospitalisedforacutedecompensated
heartfailureincorporatingN-terminalpro-B-typenatriureticpeptidelevels:AEuropeancoLlaborationonAcutedecompeNsatedHeartFailure:ELAN-HFscore.Heart2014;100:115–125.https://doi.org/10.1136/heartjnl-2013-30363236.SchouM,GustafssonF,KjaerA,HildebrandtPR.Long-termclinicalvariationofNT-proBNPinstablechronicheartfailurepatients.EurHeartJ2007;28:177–182.https://doi.org/10.1093/eurheartj/ehl44937.JanuzziJL,TroughtonR.AreserialBNPmeasurementsusefulinheartfailuremanagement?Serialnatriureticpeptidemeasurementsareusefulinheart
failuremanagement.Circulation2013;127:500–508.https://doi.org/10.1161/CIRCULATIONAHA.112.12048538.RoalfeAK,Lay-FlurrieSL,Ordóñez-MenaJM,GoyderCR,JonesNR,HobbsFDR,etal.LongtermtrendsinnatriureticpeptidetestingforheartfailureinUKprimarycare:Acohortstudy.EurHeartJ2021;43:881–891.https://doi.org/10.1093/eurheartj/ehab78139.Bayes-GenisA,CoatsAJS.’PeptideforLife’inprimarycare:Workinpro

---

### Chunk 29/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.587

ceptorantagonist;NT-proBNP,N-terminal
pro-B-typenatriureticpeptide;RASi,renin–angiotensinsysteminhibitor.thesecut-points,itisindicativeoflikelyheartfailure,andarrangingechocardiographywithin6weeksisadvised(Figure2).Complexcut-pointschemesthatincorporatemultiplecomor-biditiesinadditiontoagecanbetakenintoconsideration,especially
iftheycanbeintegratedintoelectronichealthrecordsthattrig-
geranalarmwhenNT-proBNPconcentrationsexceedspeciedthresholds.Theseshouldincludeage-andsex-specicNT-proBNPcut-pointsforthediagnosisofHFwithadjustmentsforrenal
dysfunction,atrialbrillationorutterincludingventricularrate
onthebaselinerestingelectrocardiogram(ECG)andbodymass
index(BMI).Table1showssuggestedNT-proBNPcut-offsstrat-iedbyageandgenderfornon-obesepatientswithoutkidneyfailureandatrialbrillation/utteronbaselineECG.46Thefol-lowingsuggestedmodicationsforcomorbidconditionsarebased
moreonexpertopinionratherthanonstrongevidenceandshould
berenedasmoreinformationbecomesavailable.WheneGFRis<30ml/min

---

### Chunk 30/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.583

oninpri-marycare.EurHeartJ2010;31:1881–1889.https://doi.org/10.1093/eurheartj/ehq16346.ClelandJGF,PfefferMA,ClarkAL,JanuzziJL,McMurrayJJV,MuellerC,etal.Thestruggletowardsauniversaldenitionofheartfailure–howtoproceed?EurHeartJ2021;42:2331–2343.https://doi.org/10.1093/eurheartj/ehab08247.TaylorCJ,Lay-FlurrieSL,Ordóñez-MenaJM,GoyderCR,JonesNR,RoalfeAK,etal.NatriureticpeptidelevelatheartfailurediagnosisandriskofhospitalisationanddeathinEngland2004-2018.Heart2022;108:543–549.https://doi.org/10.1136/heartjnl-2021-31919648.NatriureticPeptidesStudiesCollaboration;WilleitP,KaptogeS,WelshP,Butter-worthAS,ChowdhuryR,SpackmanSA,etal.,Natriureticpeptidesandintegratedriskassessmentforcardiovasculardisease:Anindividual-participant-data©2023EuropeanSocietyofCardiology.

1898A.Bayes-Genisetal.

---

