# ScoreItem: Lobectomia/pneumectomia

**ID:** `019bf31d-2ef0-7946-bf90-bc759eddb080`
**FullName:** Lobectomia/pneumectomia (Histórico de doenças - Cirurgias já realizadas - Cirurgias que interferem diretamente no escore)

**Preparation Metadata:**
- Quality Grade: **FAIR**
- Total Chunks: 30 de 9 artigos
- Avg Similarity: 0.485

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-7946-bf90-bc759eddb080`.**

```json
{
  "score_item_id": "019bf31d-2ef0-7946-bf90-bc759eddb080",
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

**ScoreItem:** Lobectomia/pneumectomia (Histórico de doenças - Cirurgias já realizadas - Cirurgias que interferem diretamente no escore)

**30 chunks de 9 artigos (avg similarity: 0.485)**

### Chunk 1/30
**Article:** Impact of pulmonary rehabilitation on exercise capacity, health-related quality of life, and cardiopulmonary function in lung surgery patients: a retrospective propensity score-matched analysis (2024)
**Journal:** Frontiers in Medicine
**Section:** abstract | **Similarity:** 0.601

Retrospective study of 420 NSCLC patients who underwent lung surgery (2022-2024). After propensity matching, pulmonary rehabilitation group showed significant improvements: higher FEV1/FVC (64.17% vs 50.87%, p<0.001), enhanced exercise capacity with higher maximal WR percentage (104.76% vs 90.00%, p=0.017), lower cardiac index during exercise, improved peak oxygen uptake, and smaller reductions in muscle cross-sectional area. Pulmonary rehabilitation significantly enhances exercise capacity, quality of life, and cardiopulmonary function while mitigating postoperative muscle loss.

---

### Chunk 2/30
**Article:** Limitations in exercise and functional capacity in long-term postpneumonectomy patients (2015)
**Journal:** Journal of Cardiopulmonary Rehabilitation and Prevention
**Section:** abstract | **Similarity:** 0.597

Study of 17 patients averaging 5.5 years post-pneumonectomy showed decreased exercise capacity limited primarily by the cardiovascular system. Despite exercise deficits, functional walking tests were near normal. Echocardiography showed normal heart function with mildly elevated pulmonary pressure. Long-term survivors generally maintain near-normal daily living activities, though continued cardiopulmonary evaluation remains important.

---

### Chunk 3/30
**Article:** Efficacy of Systemic Postoperative Pulmonary Rehabilitation After Lung Resection Surgery (2015)
**Journal:** Annals of Rehabilitation Medicine
**Section:** abstract | **Similarity:** 0.591

Study of 41 patients post-lung resection: 31 received therapist-supervised rehabilitation (30 min daily during hospitalization) vs 10 with self-directed exercises. Supervised group showed significant improvement on VAS pain scale after three months and improvements in lung function capacity after six months. Systemic pulmonary rehabilitation supervised by a therapist helped improve reduced pulmonary FVC and quality of life in postoperative lung resection patients.

---

### Chunk 4/30
**Article:** Propensity-Matched Analysis Demonstrates Long-Term Risk of Respiratory and Cardiac Mortality Following Pneumonectomy Compared with Lobectomy for Lung Cancer (2022)
**Journal:** Annals of Surgery
**Section:** abstract | **Similarity:** 0.561

Propensity-matched study comparing pneumonectomy and lobectomy outcomes. While 90-day complication rates were similar, pneumonectomy patients experienced higher major complication and mortality rates. The cumulative incidence of nononcologic mortality was substantially higher in the pneumonectomy cohort over five years. Pneumonia and myocardial infarction were leading nononcologic causes of death, revealing excess mortality extends beyond the immediate postoperative period.

---

### Chunk 5/30
**Article:** Exercise as medicine – evidence for prescribing exercise as therapy in 26 different chronic diseases (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.526

thusimprovedﬁtnessinonestudy(Hawkinsetal.,2002),butnotin
another(Wadelletal.,2001).Itisrecommendedthat
oxygenationtherapyshouldbeprovidedattheend
oftrainingifthepatientsarehypoxicorbecome
desaturatedduringthetraining(AmericanThoracic
Society,1999).Trainingtomusicgavebetterresults
thanwithoutmusic(Bauldoffetal.,2002),presum-ablybecausepatientswhorunwithmusicperceivethephysicalexertiontobeless,eventhoughtheyare
doingthesameamountofexercise.Speciﬁctraining
forinspiratorymusclesincreasedthestaminaof
thesemusclesbutdidnotgivethepatientsalower
perceptionofdyspneaorimprovedﬁtness(Scherer
etal.,2000).Thus,strongevidenceexiststhatendur-
ancetrainingaspartofpulmonaryrehabilitationinpatientswithCOPDimprovesexercisecapacityandhealth-relatedqualityoflife.However,dyspnea
limitstheexerciseintensity.Therefore,resistance34Pedersen&Saltin
	
	
Getulio José Mattos Do Amaral Filho - getfilho@yahoo.com.br - CPF: 034.983.039-88

training,whichmaycauselessdysp

---

### Chunk 6/30
**Article:** Exercise as medicine – evidence for prescribing exercise as therapy in 26 different chronic diseases (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.496

TolmunenT,LaukkanenJA,HintikkaJ,KurlS,ViinamakiH,SalonenR,KauhanenJ,KaplanGA,SalonenJT.Lowmaximal
oxygenuptakeisassociatedwithelevateddepressivesymptomsinmiddle-agedmen.EurJEpidemiol
2006:21:701–706.TomlinsonD,DiorioC,BeyeneJ,SungL.Effectofexerciseoncancer-relatedfatigue:ameta-analysis.Am
JPhysMedRehabil2014:93:675–686.TorresLM,YusteSanchezMJ,ZapicoGA,PrietoMD,MayoraldelMoral
O,CerezoTE,MinayoME.Effectivenessofearlyphysiotherapytopreventlymphoedemaafter
surgeryforbreastcancer:
randomised,singleblinded,clinicaltrial.BMJ2010:340:b5396.TranZV,WeltmanA.Differentialeffectsofexerciseonserumlipidand
lipoproteinlevelsseenwithchangesinbodyweight.Ameta-analysis.JAMA1985:254:919–924.TranZV,WeltmanA,GlassGV,MoodDP.Theeffectsofexerciseonbloodlipidsandlipoproteins:ameta-analysisofstudies.MedSci
SportsExerc1983:15:393–402.TroostersT,LangerD,VrijsenB,SegersJ,WoutersK,JanssensW,
GosselinkR,DecramerM,Dupont
L.Skeletalmuscleweakness,exercisetoleranceandphysicalactivityinadultswithcysticﬁbrosis.

---

### Chunk 7/30
**Article:** Exercise as medicine – evidence for prescribing exercise as therapy in 26 different chronic diseases (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.495

oupanalysis
performedtolookatthecomplexityofthepul-
monaryrehabilitationprogramprovidednoevidenceofasigniﬁcantdifferenceintreatmenteffectbetweensubgroupsthatreceivedexerciseonlyandthosethat
receivedexercisecombinedwithmorecomplexinter-
ventions.Studiesshowthatrehabilitationprogramsleadtofewerhospitalizationsandthussaveresources(Grif-
ﬁthsetal.,2000;Grifﬁthsetal.,2001).Moststudies
usehigh-intensitywalkingexercise.Onestudycomparedtheeffectofwalkingorcyclingat80%ofVO2maxvsworkingoutintheformofCallaneticsexercisesandfoundthathigh-intensitytraining
increasedﬁtnesswhiletheworkoutprogram
increasedarmmusclestamina.Bothprogramshada
positiveeffectontheexperienceofdyspnea(Nor-
mandinetal.,2002).Oxygentreatmentinconjunc-
tionwithintensivetrainingforpatientswithCOPDincreasedtrainingintensityandthusimprovedﬁtnessinonestudy(Hawkinsetal.,2002),butnotin
another(Wadelletal.,2001).Itisrecommendedthat
oxygenationtherapyshouldbeprovidedattheend
oftrainingifthepatientsarehypoxicorbecome
desaturate

---

### Chunk 8/30
**Article:** Exercise as medicine – evidence for prescribing exercise as therapy in 26 different chronic diseases (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.494

.Therefore,resistance34Pedersen&Saltin
	
	
Getulio José Mattos Do Amaral Filho - getfilho@yahoo.com.br - CPF: 034.983.039-88

training,whichmaycauselessdyspnea,couldbeanalternative.Moreover,lowmusclemassisassociated
withincreasedriskofmortality(Marquisetal.,
2002).Arecentsystematicreview(Iepsenetal.,2015b)comparedtheeffectofresistanceandendurancetraining.Theauthorsincludedeightrandomizedcontrolledtrials(328participants)andfoundthat
resistancetrainingappearedtoinducethesamebene-
ﬁcialeffectsasendurancetraining.Itwastherefore
recommendedthatresistancetrainingshouldbe
consideredaccordingtopatientpreferenceswhen
designingapulmonaryrehabilitationprogramfor
patientswithCOPD.Thesameauthorsperformedanothersystematicreview(Iepsenetal.,2015a)inwhichtheyassessedtheefﬁciencyofcombiningresis-
tancetrainingwithendurancetrainingcompared
withendurancetrainingalone.Forthisanalysis,they
included11randomizedcontrolledtrials(331partici-
pants)and2previoussys

---

### Chunk 9/30
**Article:** Exercise as medicine – evidence for prescribing exercise as therapy in 26 different chronic diseases (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.488

sical
activityandbackexercisesonlowbackpainandpsychologicaldistress:ﬁndingsfromtheUCLALowBack
PainStudy.AmJPublicHealth
2005:95:1817–1824.HwangR,MarwickT.Efﬁcacyofhome-basedexerciseprogrammesfor
peoplewithchronicheartfailure:a
meta-analysis.EurJCardiovascPrevRehabil2009:16:527–535.IbrahimEM,Al-HomaidhA.Physicalactivityandsurvivalafterbreastcancerdiagnosis:meta-analysisofpublishedstudies.MedOncol2011:
28:753–765.IepsenUW,JorgensenKJ,RingbaekT,HansenH,SkrubbeltrangC,LangeP.Acombinationof
resistanceandendurancetraining
increaseslegmusclestrengthinCOPD:anevidence-basedrecommendationbasedonsystematic
reviewwithmeta-analyses.ChronRespirDis2015a:12:132–145.IepsenUW,JorgensenKJ,RingbaekT,HansenH,SkrubbeltrangC,
LangeP.AsystematicreviewofresistancetrainingversusendurancetraininginCOPD.JCardiopulm
RehabilPrev2015b:35:163–172.IhlebaekC,HanssonTH,LaerumE,BrageS,EriksenHR,HolmSH,SvendsrodR,IndahlA.Prevalence
oflowbackpainandsicknessabsence:a“borderline”studyinNorwayandSweden.Scand

---

### Chunk 10/30
**Article:** Exercise as medicine – evidence for prescribing exercise as therapy in 26 different chronic diseases (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.485

edphysicaltrainingThepositiveimpactofphysicalexerciseforpatients
withCOPDiswelldocumented.A2015Cochrane
Review/meta-analysis(McCarthyetal.,2015)addedtopreviousmeta-analyses(Lacasseetal.,1996;Lacasseetal.,2002;Lacasseetal.,2007;Salman
etal.,2003).The2015updateincludes65RCTs
involving3822participants.Atotalof41ofthepul-
monaryrehabilitationprogramswerehospitalbased,
23werecommunitybasedandonestudyhadbotha
hospitalcomponentandacommunitycomponent.
Mostprogramswereof12-weekor8-weekdurationwithanoverallrangeof4–52weeks.Theauthorsfoundstatisticallysigniﬁcantimprovementforall
includedoutcomes.Infourimportantdomainsof
qualityoflife(ChronicRespiratoryQuestionnaire
(CRQ)scoresfordyspnea,fatigue,emotionalfunc-tion,andmastery),theeffectwaslargerthantheminimalclinicallyimportantdifferenceof0.5units.

---

### Chunk 11/30
**Article:** Exercise as medicine – evidence for prescribing exercise as therapy in 26 different chronic diseases (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.478

eartAssoc2013:2:e004473.CourneyaKS,FriedenreichCM.Physicalexerciseandqualityoflifefollowingcancerdiagnosis:a
literaturereview.AnnBehavMed1999:21:171–179.CourneyaKS,MackeyJR,JonesLW.Copingwithcancerexperience:can
physicalexercisehelp?PhysSportsmed2000:28:49–73.CrawfordDA,JefferyRW,FrenchSA.Televisionviewing,physical
inactivityandobesity.IntJObesRelatMetabDisord1999:23:437–440.CreasyTS,McMillanPJ,FletcherEW,CollinJ,MorrisPJ.Ispercutaneoustransluminalangioplastybetterthan
exerciseforclaudication?Preliminary
resultsfromaprospectiverandomisedtrial.EurJVascSurg1990:4:135–140.CreatsasG,DeligeoroglouE.Polycysticovariansyndromeinadolescents.CurrOpinObstetGynecol2007:19:420–426.CrouseSF,O’BrienBC,GrandjeanPW,LoweRC,RohackJJ,GreenJS,TolsonH.Trainingintensity,bloodlipids,andapolipoproteinsinmenwithhighcholesterol.JAppl
Physiol1997:82:270–277.Cuenca-GarciaM,JagoR,ShieldJP,BurrenCP.Howdoesphysical
activityandﬁtnessinﬂuence
glycaemiccontrolinyoungpeoplewithType1diabetes?DiabetMed201

---

### Chunk 12/30
**Article:** Exercise as medicine – evidence for prescribing exercise as therapy in 26 different chronic diseases (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.477

icalﬁtness,andhypertension.MedSciSportsExerc1993:25:i–x.AmericanDiabetesAssociation.Clinicalpracticerecommendations.DiabetesCare2002:Jan:S1–S147.AmericanThoracicSociety.Pulmonaryrehabilitation–1999.OfﬁcialstatementoftheAmericanThoracicSociety,November1998.AmJRespirCritCareMed1999:1999(159):1666–1682.AndelR,CroweM,PedersenNL,FratiglioniL,JohanssonB,GatzM.

---

### Chunk 13/30
**Article:** Exercise as medicine – evidence for prescribing exercise as therapy in 26 different chronic diseases (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.473

yclingat70–85%ofVO2max(Morganetal.,2001).Supervisedtrainingover7weeksproducedbetterresultswithregardtoa
numberofrespiratoryparametersthana4-weekpro-
gram(Greenetal.,2001).ContraindicationsNogeneralcontraindications.Thetrainingshouldtakecompetingdiseasesintoaccount.Forpatientswithlowoxygensaturation(SaO2<90%)anddyspneawhenatrest,exercisingwithoxygenationis
recommended.BronchialasthmaBackgroundBronchialasthma(asthma)isachronicinﬂamma-torydisordercharacterizedbyepisodicreversibleimpairmentofpulmonaryfunctionandairwayhyper-responsivenesstoavarietyofstimuli
(NationalInstituteofHealth,1995).Allergiesarea
majorcauseofasthmasymptoms,especiallyinchil-
dren,whilemanyadultshaveasthmawithoutan
allergiccomponent.Environmentalfactors,includ-
ingtobaccosmokeandairpollution,contributeto
thedevelopmentofasthma.Physicalexerciseposesaparticularproblemforasthmaticsasphysicalactivitycanprovokebron-
choconstrictioninmostasthmatics(Carlsenand
Carlsen,2002).Regularphysicalactivityisimportant
intherehabili

---

### Chunk 14/30
**Article:** Exercise as medicine – evidence for prescribing exercise as therapy in 26 different chronic diseases (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.472

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

### Chunk 15/30
**Article:** Lifestyle Medicine: A Brief Review of Its Dramatic Impact on Health and Survival (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.469

s for patients 
with breast cancer to improve prognosis and optimize 
overall health. CMAJ 2017 Feb 1;189(7):E268-74. 
DOI: https://doi.org/10.1503/cmaj.160464.
 113. Irwin ML, McTiernan A, Manson JE, et al. Physical 
activity and survival in postmenopausal women with 
breast cancer: Results from the women’s health 
initiative. Cancer Prev Res (Phila) 2011 Apr;4(4):522-9. 
DOI: https://doi.org/10.1158/1940-6207.capr-10-0295.
 114. Chlebowski RT. Nutrition and physical activity influence 
on breast cancer incidence and outcome. Breast 2013 
Aug;22 Suppl 2:S30-7. DOI: https://doi.org/10.1016/j.
breast.2013.07.006.
 115. Meyerhardt JA, Heseltine D, Niedzwiecki D, et al. 
Impact of physical activity on cancer recurrence 
and survival in patients with stage III colon cancer: 
Findings from CALGB 89803. J Clin Oncol 2006 
Aug 1;24(22):3535-41. DOI: https://doi.org/10.1200/
jco.2008.26.15_suppl.4039.
 116. Pierce JP, Stefanick ML, Flatt SW, et al.

---

### Chunk 16/30
**Article:** Exercise as medicine – evidence for prescribing exercise as therapy in 26 different chronic diseases (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.468

cbedrest.ClinSci(Lond)1983:64:537–540.KrotkiewskiM,LonnrothP,MandroukasK,WroblewskiZ,
Rebuffe-ScriveM,HolmG,SmithU,BjorntorpP.Theeffectsofphysicaltrainingoninsulinsecretion
andeffectivenessandonglucosemetabolisminobesityandtype2(non-insulin-dependent)diabetesmellitus.Diabetologia1985:28:881–890.KukJL,KatzmarzykPT,NichamanMZ,ChurchTS,BlairSN,RossR.Visceralfatisanindependent
predictorofall-causemortalityin
men.Obesity(SilverSpring)2006:14:336–341.LaaksonenDE,AtalayM,NiskanenLK,MustonenJ,SenCK,Lakka
TA,UusitupaMI.Aerobicexerciseandthelipidproﬁleintype1diabeticmen:arandomized
controlledtrial.MedSciSports
Exerc2000:32:1541–1548.LacasseY,BrosseauL,MilneS,MartinS,WongE,GuyattGH,
GoldsteinRS.Pulmonaryrehabilitationforchronicobstructivepulmonarydisease.Cochrane
DatabaseSystRev2002:3:
CD003793.LacasseY,MartinS,LassersonTJ,GoldsteinRS.Meta-analysisof
respiratoryrehabilitationinchronic
obstructivepulmonarydisease.ACochranesystematicreview.EuraMedicophys2007:43:
475–485.LacasseY,WongE,G

---

### Chunk 17/30
**Article:** Exercise as medicine – evidence for prescribing exercise as therapy in 26 different chronic diseases (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.465

Statisticallysigniﬁcantimprovementswerenotedin
alldomainsoftheSt.George’sRespiratoryQues-
tionnaire,andimprovementintotalscorewasbetter
than4units.Bothfunctionalexerciseandmaximalexerciseshowedstatisticallysigniﬁcantimprovement.Researchersreportedanincreaseinmaximalexercise
capacity[meanWmax(W)]inparticipantsallocatedtopulmonaryrehabilitationcomparedwithusual
care.Inrelationtofunctionalexercisecapacity,the
6-minwalkdistancemeantreatmenteffectwas
greaterthanthethresholdofclinicalsigniﬁcance.

---

### Chunk 18/30
**Article:** Pedro Neuro - Neurologia Funcional Integrativa 2 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.465

ccus salivarius, Lactobacillus sakei), raspador de língua de cobre, evitar dormir de boca aberta; atenção a periodontite/gengivite (Porphyromonas gingivalis).
  - Precauções perioperatórias:
    - Suplementação iniciada 1 semana antes e mantida por 2 semanas após anestesia/cirurgia para mitigar neurotoxicidade (redução de glutationa, risco de hipóxia/hipotensão, uso de antibióticos).
  - Programas de estilo de vida:
    - ReCODE/MAP personalizados conforme cognoscopia: metas de passos, prancha, dieta mediterrânea/Keto Flex e técnicas de respiração.
  - Exercício:
    - Caminhadas diárias: meta ≥5.000 passos, ideal ~10.000.
    - Musculação com ênfase em prancha (até 3 minutos totais/dia).
    - HIIT: protocolos curtos (ex.: 20s forte/10s leve, 8 ciclos, ~4 minutos).
  - Dieta:
    - Ketoflex 12-3 (12 horas de jejum diário, 3 horas entre jantar e sono, abordagem flexitariana com cetose monitorada).

---

### Chunk 19/30
**Article:** Exercise as medicine – evidence for prescribing exercise as therapy in 26 different chronic diseases (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.462

3793.LacasseY,MartinS,LassersonTJ,GoldsteinRS.Meta-analysisof
respiratoryrehabilitationinchronic
obstructivepulmonarydisease.ACochranesystematicreview.EuraMedicophys2007:43:
475–485.LacasseY,WongE,GuyattGH,KingD,CookDJ,GoldsteinRS.Meta-
analysisofrespiratoryrehabilitation
inchronicobstructivepulmonarydisease.Lancet1996:348:1115–1119.LandinS,HagenfeldtL,SaltinB,WahrenJ.Musclemetabolismduringexerciseinhemipareticpatients.ClinSciMolMed1977:53:
257–269.LaneR,EllisB,WatsonL,LengGC.Exerciseforintermittentclaudication.CochraneDatabase
SystRev2014:7:CD000990.LangbeinWE,CollinsEG,OrebaughC,MaloneyC,WilliamsKJ,Littooy
FN,EdwardsLC.Increasingexercise
toleranceofpersonslimitedbyclaudicationpainusingpolestriding.JVascSurg2002:35:
887–893.LangeAK,VanwanseeleB,FiataroneSinghMA.Strengthtrainingfortreatmentofosteoarthritisofthe
knee:asystematicreview.ArthritisRheum2008:59:1488–1494.LanneforsL,ButtonBM,McIlwaineM.Physiotherapyininfantsand
youngchildrenwithcysticﬁbrosis:currentpracticeandfutu

---

### Chunk 20/30
**Article:** Medicina Baseada em Evidência III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.458

s sob seus cuidados e planejar reavaliação de necessidade e risco/benefício, com foco em redução quando apropriado.
- [ ] 5. Preparar material de consentimento informado que compare riscos e benefícios de opções terapêuticas (p. ex., cirurgia vs nova quimioterapia), incluindo probabilidades de desfechos e incertezas.
- [ ] 6. Implementar intervenções de baixo risco com plausibilidade mecanística e múltiplos benefícios (ex.: curcumina, ômega-3) quando apropriado, monitorando desfechos clínicos (p. ex., dor).
- [ ] 7. Investigar casos clínicos relevantes (ex.: cetogênica e cetose, relato da doutora Janaína) e documentar resultados, contextualizando a ausência de “nível A” formal em abordagens personalizadas.
- [ ] 8. Desenvolver um roteiro de comunicação para pacientes que mitigue o viés de autoridade, promovendo compreensão crítica de estudos e alinhamento com valores e preferências individuais.
- [ ] 9.

---

### Chunk 21/30
**Article:** Exercise as medicine – evidence for prescribing exercise as therapy in 26 different chronic diseases (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.457

rcisecapacity.Therecommended
methodforthisisasymptom-limitedexercisetest,whichcanbecarriedoutbytrainedpersonnel(physiotherapist,nurse,laboratoryassistant)
underadoctor’ssupervision.�Supervisedtrainingwithindividuallyorganized
trainingprogramsafteraninitialexercisetest:2–530–60-minsessionsaweekatanintensityof50–80%ofmaximumexercisecapacity.�Twelveweeksoforganizedaerobictrainingandpossiblyintervaltrainingcombinedwithresistancetraining,especiallyfortheelderlyandpatients
withmuscleweakness.�Dailylow-intensitytraining(walking)over
30min,increasingthelevelunderthesupervision
oftherehabilitationteam.29Exerciseasmedicine–evidenceforprescribingexercise
	
	
Getulio José Mattos Do Amaral Filho - getfilho@yahoo.com.br - CPF: 034.983.039-88

ContraindicationsThefollowingcontraindicationsareinagreementwithaEuropeanWorkingGroup(Gianuzzi&
Tavazzi2001).�AcuteCHU(AMIorunstableangina),untilthe
conditionhasbeenstableforatleast5days�Dyspneaatr

---

### Chunk 22/30
**Article:** Mastologia I (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.456

s metabolizam glicose via glicólise, mesmo com oxigênio, favorecendo rápida multiplicação.
*   **Importância do Exercício de Força**
    - “O sedentarismo será o tabagismo do futuro”.
    - Priorizar exercícios de força, não apenas caminhadas.
    - Meta-análise: sarcopenia associada a 44% mais mortes por todas as causas e 93% mais mortes por câncer.
    - Em câncer de mama, sarcopenia aumenta mortalidade em 41%.

## ❓ Perguntas
- [Inserir Pergunta/Dúvida]

## 📚 Tarefas
- [ ] 1. Estudar fatores de risco para câncer de mama além da genética: alimentação, microbiota, sono, estresse, obesidade e resistência à insulina.
- [ ] 2. Aprender a identificar sinais de resistência à insulina e inflamação crônica, inclusive em pacientes com peso normal.
- [ ] 3. Incorporar na prática clínica a orientação sobre exercícios de força, além de atividades aeróbicas, para prevenção e melhor prognóstico.
- [ ] 4.

---

### Chunk 23/30
**Article:** Exercise as medicine – evidence for prescribing exercise as therapy in 26 different chronic diseases (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.453

gherTNFlevels
inblood(Eidetal.,2001)andmuscletissue(Palacio
etal.,2002).TNF’sbiologicalimpactonmuscletissue
ismanifold.TNFaffectsmyocytedifferentiation,
inducescachexia,andthusapotentialdecreaseinmusclestrength(Li&Reid,2001).ADanishstudyshowedthatsmokinginhibitedproteinsynthesisin
themuscles,whichcanpotentiallyalsoleadtolossof
musclemass(Petersenetal.,2007).Trainingcanpre-
sumablyhaveanimpactonthisprocess.Another
Danishstudyshowedthatphysicaltrainingcounter-
actedtheincreaseinproteindegradationseenin
peoplewithCOPD(Petersenetal.,2008).TypeoftrainingAllpatientswithCOPD,particularlythemoresevere
cases,beneﬁtfromphysicaltraining.Initiallythe
physicaltrainingmustbesupervised,individually
tailoredandincludeacombinationofendurance
trainingandstrengthtraining.Endurancetraining
canbewalkingorcyclingat70–85%ofVO2max(Morganetal.,2001).Supervisedtrainingover7weeksproducedbetterresultswithregardtoa
numberofrespiratoryparametersthana4-weekpro-
gram(Greenetal.,2001).ContraindicationsNogeneralc

---

### Chunk 24/30
**Article:** Exercise as medicine – evidence for prescribing exercise as therapy in 26 different chronic diseases (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.453

onraadsetal.,2002),TNFandFAS-L(Adamopoulosetal.,2002),andthequantityofcirculatingadhesionmolecules
(Adamopoulosetal.,2001)inpatientswithheart
failure.Physicaltraininglowerstheexpressionof
cytokinesintheskeletalmuscle(Gielenetal.,2003)
andinthebloodstream(LeMaitreetal.,2004).TypeoftrainingManystudieshavedemonstratedabeneﬁcialeffect
fromintervaltraining,whichispossiblymoreeffec-
tivethanmoderatecontinualaerobictraining(2001c;
Wisloffetal.,2007).Patientscanstartoninterval
training,beginningonalowexercisecapacity,and
graduallyincreaseduration,frequency,andintensity
(2001c).Practitionersusedtobereluctanttorecommendresistancetrainingoutofconcernthatincreasedvas-
cularresistancewouldincreasecardiacloadmore
thanaerobictraining.Thereisnoevidencethata
combinationofaerobicandresistancetrainingpro-
ducesbetterorworseresultsthanaerobictraining
alone(Haykowskyetal.,2007).Basedontheabove,thefollowingisrecom-mended:�TrainingisrecommendedforallheartfailurepatientsinNYHAfunctionclassII–IIIonafully

---

### Chunk 25/30
**Article:** Exercise as medicine – evidence for prescribing exercise as therapy in 26 different chronic diseases (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.452

J,PayneN,
NewcombeRG,IonescuAA,ThomasJ,TunbridgeJ,LonescuAA.Resultsat1yearofoutpatient
multidisciplinarypulmonaryrehabilitation:arandomisedcontrolledtrial.Lancet2000:355:
362–368.GrifﬁthsTL,PhillipsCJ,DaviesS,BurrML,CampbellIA.Costeffectivenessofanoutpatient
multidisciplinarypulmonary
rehabilitationprogramme.Thorax2001:56:779–784.GrodsteinF,LevineR,TroyL,SpencerT,ColditzGA,StampferMJ.Three-yearfollow-upofparticipantsinacommercialweight
lossprogram.Canyoukeepitoff?

---

### Chunk 26/30
**Article:** Exercise as medicine – evidence for prescribing exercise as therapy in 26 different chronic diseases (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.452

withresistancetraining.Theaer-
obicphysicalexerciseshouldstartatalowintensity
andbegraduallysteppeduptomoderateandﬁnallyhighintensity,graduallyincreasingthedura-tionofthetrainingatthesametime.Theaerobic
trainingshouldbecombinedwithresistancetrain-
ing,whichalsostartsatalowexertionleveland
shortdurations.Itisrecommendedthattrainingshouldbesuper-visedbutthatrelativeandabsolutecontraindications
shouldbeobserved.Evenhospitalizedandbed-rid-denpatientscanproﬁtfromphysicaltraining(Dimeoetal.,1999),butthereissparseinformation
aboutexerciseduringchemotherapyorradiother-
apy.Itisimportanttoemphasizethatthispatient
groupissoheterogeneousthatstandardproposals
makenosenseandformany,especiallyelderlycan-
cerpatients,thefocusoughttobeonretaining
mobilityandphysicalability.ContraindicationsItisadvisedthatpatientsundergoingchemotherapyorradiotherapywithaleukocytecountbelow
0.5910(9)/L,hemoglobinbelow6mmol/L,throm-bocytecountbelow20910(9)/L,temperatureabove38°Cshouldnotexercise.Patientswithbonemeta

---

### Chunk 27/30
**Article:** Exercise as medicine – evidence for prescribing exercise as therapy in 26 different chronic diseases (2026)
**Journal:** Revista não identificada
**Section:** introduction | **Similarity:** 0.451

hereasonwhywalk-
ingexerciseispreferable.Ifbikeexerciseischosen,thepatientmustbeinstructedtopedalusingthe33Exerciseasmedicine–evidenceforprescribingexercise
	
	
Getulio José Mattos Do Amaral Filho - getfilho@yahoo.com.br - CPF: 034.983.039-88

frontofthefootandthesametrainingprinciplesapplyasforwalking.ContraindicationsTherearenogeneralcontraindications.Supervised
exercisetrainingcansafelybeprescribedinpatients
withintermittentclaudicationbecauseanexceed-inglylowall-causecomplicationrateisfound.Rou-tinecardiacscreeningbeforecommencingexercise
trainingisnotrequired.PulmonarydiseasesChronicobstructivepulmonarydiseaseBackgroundChronicobstructivepulmonarydisease(COPD)is
characterizedbyanirreversibledecreaseinlungfunction.Advanced-stageCOPDisalongandpain-fulprocessofgraduallyincreasingandultimately
disablingbreathlessnessasthemainsymptom.

---

### Chunk 28/30
**Article:** Exercise as medicine – evidence for prescribing exercise as therapy in 26 different chronic diseases (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.451

rval,e.g.,walking,running,step
machine,cycling,rowing,orstairwalking.Theaer-
obictrainingpresumablyshouldbebackedupby
strengthtrainingandtrainingintensityshouldbe
increasedasthepatient’sexercisecapacityincreases.Physicaltrainingisalsoadvisableforpatientswithanginapectoriswhoarenotcandidatesforrevascu-
larization(Thompsonetal.,2003).Itisnotclear
howlongthesupervisedtrainingprogramshouldbe,
butthebulkofthestudiesthatwerepartofthe
above-mentionedmeta-analysisinvolvedaduration
of6–24weekswithaweightedaverageofapproxi-mately11weeks.Theeffectwasnotdeterminedbydurationbutbytheoverall“trainingdose”andnodifferencewasfoundinmortalityaftertrainingpro-
gramsinvolvingagenerallylargedoseasopposedto
alowdose(Tayloretal.,2004).Lengthiertraining
programsareaimedatensuringthatthepatient
achievesatrainingeffectandpartlyathelpingto
incorporatenewexercisehabits.Theworkinggroup
assessedthattrainingshouldextendover12weeks,withashorterprogramorlongerprogramforselectedpatientgroupsafterassessment.Inare

---

### Chunk 29/30
**Article:** Exercise as medicine – evidence for prescribing exercise as therapy in 26 different chronic diseases (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.451

ssessedtheefﬁciencyofcombiningresis-
tancetrainingwithendurancetrainingcompared
withendurancetrainingalone.Forthisanalysis,they
included11randomizedcontrolledtrials(331partici-
pants)and2previoussystematicreviews.They
foundequalimprovementsinqualityoflife,walking
distance,andexercisecapacity.However,theyalsofoundmoderateevidenceofasigniﬁcantincreaseinlegmusclestrengthfavoringacombinationofresis-
tanceandendurancetrainingandrecommendthat
resistancetrainingshouldbeincorporatedinrehabil-
itationofCOPDtogetherwithendurancetraining.PossiblemechanismsPhysicalactivitydoesnotimprovelungfunctionin
patientswithCOPDbutincreasesCRFviatheeffectonthemusclesandtheheart.PatientswithCOPDhavechronicinﬂammation,whichmaybeacauseof
thedecreaseinmusclestrengthobservedinCOPD
patients.PatientswithCOPDhavehigherTNFlevels
inblood(Eidetal.,2001)andmuscletissue(Palacio
etal.,2002).TNF’sbiologicalimpactonmuscletissue
ismanifold.TNFaffectsmyocytedifferentiation,
inducescachexia,andthusapotentialdecreaseinmu

---

### Chunk 30/30
**Article:** Exercise as medicine – evidence for prescribing exercise as therapy in 26 different chronic diseases (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.448

ratorymuscletraininginseverelydisabled
multiplesclerosispatients.ArchPhys
MedRehabil2000:81:747–751.GreenJS,StanforthPR,RankinenT,LeonAS,RaoDD,SkinnerJS,
BouchardC,WilmoreJH.The
effectsofexercisetrainingonabdominalvisceralfat,bodycomposition,andindicatorsofthe
metabolicsyndromein54Pedersen&Saltin
	
	
Getulio José Mattos Do Amaral Filho - getfilho@yahoo.com.br - CPF: 034.983.039-88

postmenopausalwomenwithandwithoutestrogenreplacementtherapy:theHERITAGEfamily
study.Metabolism2004:53:1192–1196.GreenRH,SinghSJ,WilliamsJ,MorganMD.Arandomised
controlledtrialoffourweeksversussevenweeksofpulmonaryrehabilitationinchronicobstructive
pulmonarydisease.Thorax2001:56:
143–145.GrifﬁthsTL,BurrML,CampbellIA,Lewis-JenkinsV,MullinsJ,ShielsK,
Turner-LawlorPJ,PayneN,
NewcombeRG,IonescuAA,ThomasJ,TunbridgeJ,LonescuAA.Resultsat1yearofoutpatient
multidisciplinarypulmonaryrehabilitation:arandomisedcontrolledtrial.Lancet2000:355:
362–368.GrifﬁthsTL,Philli

---

