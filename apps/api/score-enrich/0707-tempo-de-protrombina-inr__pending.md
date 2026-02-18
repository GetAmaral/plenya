# ScoreItem: Tempo de Protrombina (INR)

**ID:** `019bf31d-2ef0-708f-9735-029dd00e67b6`
**FullName:** Tempo de Protrombina (INR) (Exames - Laboratoriais)
**Unit:** ratio

**Preparation Metadata:**
- Quality Grade: **GOOD**
- Total Chunks: 30 de 10 artigos
- Avg Similarity: 0.550

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-708f-9735-029dd00e67b6`.**

```json
{
  "score_item_id": "019bf31d-2ef0-708f-9735-029dd00e67b6",
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

**ScoreItem:** Tempo de Protrombina (INR) (Exames - Laboratoriais)
**Unidade:** ratio

**30 chunks de 10 artigos (avg similarity: 0.550)**

### Chunk 1/30
**Article:** International Normalized Ratio: Assessment, Monitoring, and Clinical Implications (2025)
**Journal:** StatPearls
**Section:** abstract | **Similarity:** 0.696

Comprehensive review on INR monitoring for anticoagulation therapy, covering assessment methods, therapeutic ranges, and clinical implications. Updated February 2025 with current guidelines for warfarin management and monitoring protocols.

---

### Chunk 2/30
**Article:** Utility of TTR-INR guided warfarin adjustment protocol to improve time in therapeutic range in patients with atrial fibrillation receiving warfarin (2024)
**Journal:** Scientific Reports
**Section:** abstract | **Similarity:** 0.655

Prospective cohort study demonstrating that TTR-INR guided warfarin adjustment protocol significantly improves time in therapeutic range from 65% to 80% in atrial fibrillation patients. Proportion of patients achieving adequate anticoagulation control increased from 47% to 88% over 12 months.

---

### Chunk 3/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.602

stroke,and
CHADS2score.Therewasasuggestionthatmajorbleedingwassignicantlyreducedinpeopleattendingcenterswheretimeintherapeuticinternationalnormalizedratio(INR)range
was<66%comparedwithcenterswith$66%timeinrange(interactionP¼0.02).ThissuggeststhatbenetsofNOACsareinpartaresultoftheirsimplerpharmacokineticproleanddosing.714The2021meta-analysisthatfocusedonCKDsubgroupsfrom7trialsfoundthatbleedingeventswerealso
notsignicantlydifferentamongthoseallocatedNOACsversuswarfarin(HR:0.83;95%CI:0.58–1.18).715DatainCKDG5ondialysiswerelimitedtoobservationalstudies.715Ourevidencereviewwasagainlimitedtoasmallnumberofstudiesreportingsubtypesofbleedingoutcomes,andsoanalysesfoundimpreciseestimatesoftreatmenteffect.The
ndingswerequalitativelyconsistentwiththetotalityoftheevidence(Figure42,SupplementaryTableS15716–722).ThereviewraisedahypothesisthatsomeNOACsmaybemore
likelytoreducetheriskofbleeding.However,giventhe
evidenceofeffectmodicationbytimeintherapeuticrangeinthewarfaringroup,wehavenotprovi

---

### Chunk 4/30
**Article:** Outcomes of Warfarin Home INR Monitoring vs Office-Based Monitoring: a Retrospective Claims-Based Analysis (2024)
**Journal:** Journal of General Internal Medicine
**Section:** abstract | **Similarity:** 0.592

Retrospective claims-based analysis comparing patient self-testing (PST) versus laboratory-based INR monitoring. Demonstrates real-world outcomes and effectiveness of home INR monitoring technologies for warfarin management, with implications for time in therapeutic range and patient safety.

---

### Chunk 5/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.588

04 (0.07, 16.70)0.51 (0.24, 1.09)CrCl 30–50CrCl 30–4946 countries
45 countries2.8 yrs
707 dysEdoxaban 60 mg
Rivaroxaban 20 mgWarfarin
Warfarin0.46 (0.26, 0.82)
0.82 (0.41, 1.60)0.60 (0.34, 1.05)Author, yearAll clinically relevant bleeding†Hijazi, 2021Stanifer, 2020Hori, 2013Fox, 2011Subtotal (I2 = 79.9%, P=0.002)Fatal bleedingBohula, 2016Hori, 2013Subtotal (I2 = 0.0%, P=0.595)Major bleeding†Hijazi, 2021Stanifer, 2020Hijazi, 2018Bohula, 2016Hori, 2013Fox, 2011Subtotal (I2 = 79.0%, P=0.000)Intracranial hemorrhageBohula, 2016Fox, 2011Subtotal (I2 = 38.2%, P=0.203)Hori, 2013Fox, 2011Subtotal (I2 = 14.1%, P=0.281)CrCl 30–49CrCl 30–49Japan
45 countries2.5 yrs
707 dysRivaroxaban 10 mg
Rivaroxaban 20 mgWarfarin
Warfarin1.35 (0.82, 2.22)
1.01 (0.85, 1.20)1.06 (0.86, 1.31)Clinically relevant nonmajor bleeding‡Weighted hazard ratio of bleeding0.20.5512Favors NOACFavors control
Figure42|Pooledhazardratio(HR)comparingnon–vitaminKantagonistoralanticoagulants(NOACs)withwarfarinamongpe

---

### Chunk 6/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.569

International(2024)105(Suppl4S),S117–S314S241

Harms.The2014meta-analysisof4largephaseIIItrialsfoundthatNOACsreducedtheriskofdeathfromanycauseby10%,conrmingnetsafety(RR:0.90;95%CI:0.85–0.95).Comparedwithwarfarin,NOACsreducedtheriskofintracranialhemorrhage(denedashemorrhagicstroke,epidural,subdural,andsubarachnoidhemorrhage)
byaboutone-half(RR:0.48;95%CI:0.39–0.59),andtheriskofgastrointestinalbleedingwasincreasedbyabout
one-quarter(RR:1.25;95%CI:1.01–1.55).Overall,therewasnocleareffectonthecombinationofthese2safety
outcomesreferredtoasmajorbleeding(RR:0.86;95%CI:
0.73–1.00).714TherewerelargeamountsofdataonmajorbleedinginthosewithaCrClof<50ml/min,soreassuringsafetydataclearlyextendedtopeoplewith
CKD.Therewerealsoconsistentsafetydatainsubgroup
analysesbyage,sex,priordiabetes,priorstroke,and
CHADS2score.Therewasasuggestionthatmajorbleedingwassignicantlyreducedinpeopleattendingcenterswheretimeintherapeuticinternationalnormalizedratio(INR)range
was<66%comparedwithcenterswith$66%

---

### Chunk 7/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.565

llationandadvancedchronickidneydisease.Circulation.2020;141:1384–1392.722.HijaziZ,AlexanderJH,LiZ,etal.ApixabanorvitaminKantagonistsandaspirinorplaceboaccordingtokidneyfunctioninpatientswithatrialbrillationafteracutecoronarysyndromeorpercutaneouscoronaryinterventioninsightsfromtheAUGUSTUStrial.Circulation.2021;143:1215–1223.723.AltawalbehSM,AlshogranOY,SmithKJ.Cost-utilityanalysisofapixaban
versuswarfarininatrialbrillationpatientswithchronickidneydisease.ValueHealth.2018;21:1365–1372.724.HeidbuchelH,VerhammeP,AlingsM,etal.UpdatedEuropeanHeart
RhythmAssociationpracticalguideontheuseofnon-vitamin-K
antagonistanticoagulantsinpatientswithnon-valvularatrial
brillation:executivesummary.EurHeartJ.2017;38:2137–2149.725.VondracekSF,TeitelbaumI,KiserTH.Principlesofkidney
pharmacotherapyforthenephrologist:CoreCurriculum2021.AmJKidneyDis.2021;78:442–458.726.DorksM,AllersK,SchmiemannG,etal.Inappropriatemedicationinnon-
hospitalizedpatientswithrenalinsufciency:asystematicreview.JAmGeri

---

### Chunk 8/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.560

0 ml/min≥24 h≥48 h≥24 h≥48 hCrCl 50–80 ml/min≥36 h≥72 h≥24 h≥48 hCrCl 30–50 ml/mina≥48 h≥96 h≥24 h≥48 hCrCl 15–30 ml/minaNo ocial indicationNo ocial indication≥36 h≥48 hCrCl <15 ml/minNo ocial indication for useThere is no need for bridging with LMWH/UFH
Figure44|Adviceonwhentodiscontinuenon–vitaminKantagonistoralanticoagulants(NOACs)beforeprocedures(lowvs.highrisk).Theboldvaluesdeviatefromthecommonstoppingruleof$24-hourlowrisk,$48-hourhighrisk.Lowriskisdenedasalowfrequencyofbleedingand/orminorimpactofableed.Highriskisdenedasahighfrequencyofbleedingand/orimportantclinicalimpact.AdaptedfromHeidbuchelH,VerhammeP,AlingsM,etal.UpdatedEuropeanHeartRhythmAssociationpracticalguideontheuseofnon–vitamin-Kantagonistanticoagulantsinpatientswithnon-valvularatrialbrillation:executivesummary.EurHeartJ.2017;38:2137–2149.724aManyofthesepeoplemaybeonlowerdoseofdabigatran(110mgtwiceperday[b.i.d.])orapixaban(2.5mgb.i.d.),orhavetobeonthelowerdoseofrivaroxaban(

---

### Chunk 9/30
**Article:** Bases Metabólicas das Doenças Crônicas e Gerenciamento - Inflamação 2 (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.559

ejo da inflamação.
**A suplementação com curcuminoides, padronizada a 95%, é uma estratégia anti-inflamatória eficaz, com doses que variam de 500 mg a 2g, mas que exigem cautela em pacientes que usam anticoagulantes.**
- Uma meta-análise de 15 ensaios clínicos randomizados valida o estudo dos efeitos da curcumina.
- A dose inicial recomendada é de 500 mg, podendo chegar a um máximo de 2g por dia.
- Para pacientes que usam anticoagulantes, a dose máxima recomendada é de 500 mg, com uma dose de segurança sugerida de 250 mg (preferencialmente lipossomada para melhor absorção).
- A absorção é frequentemente melhorada pela adição de 10 mg de piperina, que pode, no entanto, causar reações alérgicas em algumas pessoas.

---

### Chunk 10/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.559

Theyofferbenetsintermsofeaseofmonitoring.CKDdoesnotappeartoimportantlymodifythesebenets,atleastdowntoG4.PracticePoint3.16.2:NOACdoseadjustmentforGFRisrequired,withcautionneededatCKDG4–G5.DosesofNOACsmayneedtobemodiedinpeoplewithdecreasedGFRtakingintoconsiderationtheage,weight,andGFRofapersonwithCKD(Figure43710).Consulttherelevantsummariesofproductcharacteristicsforthelatest
informationondosing(Chapter4).PracticePoint3.16.3:DurationofNOACdiscontinuationbeforeelectiveproceduresneedstoconsiderproceduralbleedingrisk,NOACprescribed,andlevelofGFR(Figure44).710,724
eCrCl (ml/min)aWarfarinApixabanbDabigatranEdoxabancRivaroxaban>95Adjusted dose (INR 2–3)5 mg b.i.d.150 mg b.i.d.60 mg QDd20 mg QD51–95Adjusted dose (INR 2–3)5 mg b.i.d.150 mg b.i.d.60 mg QD20 mg QD31–50Adjusted dose (INR 2–3)5 mg b.i.d.
150 mg b.i.d.

---

### Chunk 11/30
**Article:** 2025 Guidelines for direct oral anticoagulants: a practical guidance on the prescription, laboratory testing, peri-operative and bleeding management (2025)
**Journal:** Internal Medicine Journal
**Section:** abstract | **Similarity:** 0.555

Updated guidelines from the Thrombosis and Haemostasis Society of Australia and New Zealand (THANZ) comparing direct oral anticoagulants (DOACs) with warfarin/VKAs. Addresses advantages of DOACs including predictable effect, no routine monitoring requirement, and comparison with INR-based warfarin management.

---

### Chunk 12/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** conclusion | **Similarity:** 0.546

R.eIncountrieswhere110mgb.i.d.isapproved,healthcareprovidersmaypreferthisdoseafterclinicalassessmentofthromboembolicversusbleedingrisk.ThisdosehasnotbeenapprovedforusebytheUSFDA.fNOACdoseslistedinparenthesisaredosesthatdonotcurrentlyhaveanyclinicalorefcacydata.ThedosesofNOACsapixaban5mgb.i.d.,brivaroxaban15mgeveryday,anddabigatran75mgb.i.d.areincludedintheUSFDA–approvedlabelingbasedonlimiteddosepharmacokineticandpharmacodynamicsdatawithnoclinicalsafetydata.Wesuggestconsiderationofthelowerdoseofapixaban2.5mgoralb.i.d.inCKDG5andG5Dtoreducebleedingriskuntilclinicalsafetydataareavailable.gDabigatran75mgavailableonlyintheUnitedStates.hThedosewashalvedifanyofthefollowing:estimatedCrClof30–50ml/min,bodyweightof#60kg,orconcomitantuseofverapamilorquinidine(potentP-glycoproteininhibitors).INR,internationalnormalizedratio;QD,everyday.ReproducedfromTurakhiaMP,BlankestijnPJ,CarreroJJ,etal.Chronickidneydiseaseandarrhythmias:conclusionsfromaKidneyDisease:
ImprovingGlobalOutcomes(KDIGO)Controvers

---

### Chunk 13/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.544

entlevelsofrisk(basedonageandsex)S240Figure39.Meta-analyzedadjustedprevalenceofatrialbrillationfromcohortscontributingtotheChronicKidneyDiseasePrognosisConsortium,bydiabetesstatusS241Figure40.StrategiesforthediagnosisandmanagementofatrialbrillationS242Figure41.Pooledhazardratio(HR)comparingnon–vitaminKantagonistoralanticoagulants(NOACs)withwarfarinamongpeoplewithchronickidneydiseaseintermsofstrokeS243Figure42.Pooledhazardratio(HR)comparingnon–vitaminKantagonistoralanticoagulants(NOACs)withwarfarinamongpeoplewithchronickidneydiseaseintermsofbleedingS244Figure43.Evidencefrom(a)randomizedcontrolledtrials(RCTs)regardingtherapeuticanticoagulationdosebyglomerularltrationrate(GFR)and(b)inareaswhereRCTsarelackingS245Figure44.Adviceonwhentodiscontinuenon–vitaminKantagonistoralanticoagulants(NOACs)beforeprocedures(lowvs.highrisk)
www.kidney-international.orgcontentsKidneyInternational(2024)105(Suppl4S),S117–S314S121

S248Figure45.Selectedherbalremediesanddietarysupplementswithevidence

---

### Chunk 14/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.542

beassociatedwithhighercostsandachievefewerquality-adjustedlife-yearscomparedwithNOACs.723Considerationsforimplementation.Adecisionnottoanti-coagulateforthromboembolicprophylaxisduetolowrisk
wouldideallybere-evaluatedateachconsultationandatleast
every6months.WhenusingantithrombotictherapyinpeoplewithCKD,itisprudenttotreatmodiableriskfactorsforbleeding(e.g.,alcoholintake)andusegastroprophylaxiswithaPPI,particularlywhencombinedwithantiplatelet
therapy.RationaleAnumberoflargeRCTsdemonstratedthatNOACsreduce
theriskofintracranialbleedingcomparedwithwarfarinand,
overall,modestlyreducemortalityinpeoplewithatrial
KidneyfunctionCountryFollow-up
lengthIntervention*ControlHR (95% CI)CrCl 30–50eGFR 25–50eGFR <50CrCl 30–50CrCl 30–49CrCl 30–4933 countries39 countries44 countries46 countriesJapan45 countries6 mos1.8 yrs1.8 yrs2.8 yrs2.5 yrs707 dysApixaban 2.5–5 mgApixaban 2.5–5 mgDabigatran 150 mgEdoxaban 60 mgRivaroxaban 10 mgRivaroxaban 20 mgWarfarinWarfarinWarfarinWarfarinWarfarinWar

---

### Chunk 15/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.540

P=0.980)Ischemic strokeStanifer, 2020Hijazi, 2018Bohula, 2016Hori, 2013Fox, 2011
Subtotal (I2=19.5%, P=0.291)Hemorrhagic strokeBohula, 2016Hori, 2013Fox, 2011Subtotal (I2=0.0%, P=0.619)Weighted hazard ratio of stroke0.20.5512Favors NOACFavors control
Figure41|Pooledhazardratio(HR)comparingnon–vitaminKantagonistoralanticoagulants(NOACs)withwarfarinamongpeoplewithchronickidneydiseaseintermsofstroke.BohulaE,GiuglianoR,RuffC,etal.ImpactofrenalfunctiononoutcomeswithedoxabanintheENGAGEAF-TIMI48trial.Circulation.2016;134:24–36716;FoxKA,PicciniJP,WojdylaD,etal.Preventionofstrokeandsystemicembolismwithrivaroxabancomparedwithwarfarininpatientswithnon-valvularatrialbrillationandmoderaterenalimpairment.EurHeartJ.2011;32:2387–2394718;HijaziZ,HohnloserSH,OldgrenJ,etal.Efcacyandsafetyofdabigatrancomparedwithwarfarininpatientswithatrialbrillationinrelationtorenalfunctionovertime—aRE-LYtrialanalysis.AmHeartJ.2018;198:169–177719;HijaziZ,AlexanderJH,LiZ,etal.ApixabanorvitaminKantagonistsand

---

### Chunk 16/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.540

vant nonmajor bleeding‡Weighted hazard ratio of bleeding0.20.5512Favors NOACFavors control
Figure42|Pooledhazardratio(HR)comparingnon–vitaminKantagonistoralanticoagulants(NOACs)withwarfarinamongpeoplewithchronickidneydiseaseintermsofbleeding.BohulaE,GiuglianoR,RuffC,etal.ImpactofrenalfunctiononoutcomeswithedoxabanintheENGAGEAF-TIMI48trial.Circulation.2016;134:24–36716;FoxKA,PicciniJP,WojdylaD,etal.Preventionofstrokeandsystemicembolismwithrivaroxabancomparedwithwarfarininpatientswithnon-valvularatrialbrillationandmoderaterenalimpairment.EurHeartJ.2011;32:2387–2394718;HijaziZ,HohnloserSH,OldgrenJ,etal.Efcacyandsafetyofdabigatrancomparedwithwarfarininpatientswithatrialbrillationinrelationtorenalfunctionovertime—aRE-LYtrialanalysis.AmHeartJ.2018;198:169–177719;HoriM,MatsumotoM,TanahashiN,etal.SafetyandefcacyofadjusteddoseofrivaroxabaninJapanesepatientswithnon-valvularatrialbrillation:subanalysisofJ-ROCKETAFforpatientswithmoderaterenalimpairment.CircJ.2013;77:632–638720;S

---

### Chunk 17/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.538

co da dislipidemia.
- [ ] 2. Analisar criticamente os estudos sobre o Inclisiran (Cibrava), focando na diferença entre desfechos substitutos (redução de LDL) e desfechos clínicos duros (mortalidade, infarto, AVC).
- [ ] 3. Utilizar o site `the-nnt.com` para pesquisar o NNT e NNH de outros medicamentos prescritos na prática clínica.
- [ ] 4. Refletir sobre a influência da indústria farmacêutica e dos interesses financeiros na prescrição de novos medicamentos.
- [ ] 5. Estudar a diferença entre risco relativo e risco absoluto para interpretar criticamente os achados de estudos científicos.
- [ ] 6. Pesquisar sobre os diferentes tipos de ômega-3 (EPA, DHA, ALA) e suas formulações para entender como a qualidade do suplemento afeta os resultados.
- [ ] 7. Investigar os exames de subfracionamento de LDL e a relação ApoA/ApoB como ferramentas de avaliação de risco cardiovascular mais precisas que o LDL total.
- [ ] 8.

---

### Chunk 18/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.537

an45 countries6 mos1.8 yrs1.8 yrs2.8 yrs2.5 yrs707 dysApixaban 2.5–5 mgApixaban 2.5–5 mgDabigatran 150 mgEdoxaban 60 mgRivaroxaban 10 mgRivaroxaban 20 mgWarfarinWarfarinWarfarinWarfarinWarfarinWarfarin0.51 (0.28, 0.93)0.59 (0.45, 0.77)1.11 (0.87, 1.14)0.76 (0.58, 0.98)0.89 (0.36, 2.18)0.98 (0.73, 1.30)
0.80 (0.61, 1.05)NOTE: Weights are from random eects analysisCrCl 30–50eGFR 25–50CrCl 30–49CrCl 30–4933 countries39 countriesJapan45 countries6 mos1.8 yrs2.5 yrs707 dysApixaban 2.5–5 mgApixaban 2.5–5 mgRivaroxaban 10 mgRivaroxaban 20 mgWarfarinWarfarinWarfarinWarfarin0.59 (0.41, 0.84)0.35 (0.17, 0.72)1.22 (0.78, 1.91)0.98 (0.85, 1.15)0.76 (0.51, 1.14)CrCl 30–50CrCl 30–4946 countries
Japan2.8 yrs
2.5 yrsEdoxaban 60 mg
Rivaroxaban 10 mgWarfarin
Warfarin0.48 (0.22, 1.07)
1.04 (0.07, 16.70)0.51 (0.24, 1.09)CrCl 30–50CrCl 30–4946 countries
45 countries2.8 yrs
707 dysEdoxaban 60 mg
Rivaroxaban 20 mgWarfarin
Warfarin0.46 (0.26, 0.82)
0.82 (0.41, 1.60)0.60 (0.34, 1.05)Au

---

### Chunk 19/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.537

sfor
ischaemicstrokeinatrialbrillationacrossthespectrumofkidneyfunction.EurHeartJ.2021;42:1476–1485.714.RuffCT,GiuglianoRP,BraunwaldE,etal.Comparisonoftheefcacyandsafetyofneworalanticoagulantswithwarfarininpatientswithatrial
brillation:ameta-analysisofrandomisedtrials.Lancet.2014;383:955–962.715.SuX,YanB,WangL,etal.Oralanticoagulantagentsinpatientswith
atrialbrillationandCKD:asystematicreviewandpairwisenetworkmeta-analysis.AmJKidneyDis.2021;78:678–689.e1.716.BohulaE,GiuglianoR,RuffC,etal.Impactofrenalfunctiononoutcomes
withedoxabanintheENGAGEAF-TIMI48trial.Circulation.2016;134:24–36.717.ChashkinaMI,AndreevDA,KozlovskayaNL,etal.[Safetyperformanceof
rivaroxabanversuswarfarininpatientswithatrialbrillationand
www.kidney-international.orgreferencesKidneyInternational(2024)105(Suppl4S),S117–S314S309

advancedchronickidneydisease].Kardiologiia.2020;60:1322[inRussian].718.FoxKA,PicciniJP,WojdylaD,etal.Preventionofstrokeandsystemicembolismwithrivaroxabancomparedwithwarfarininpati

---

### Chunk 20/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.536

tematicreviewandmeta-analysis.EurJPrevCardiol.2022;28:1953–1960.926.KimachiM,FurukawaTA,KimachiK,etal.Directoralanticoagulants
versuswarfarinforpreventingstrokeandsystemicembolicevents
amongatrialbrillationpatientswithchronickidneydisease.CochraneDatabaseSystRev.2017;11:CD011373.927.SterneJAC,SavovicJ,PageMJ,etal.RoB2:arevisedtoolforassessingriskofbiasinrandomisedtrials.BMJ.2019;366:l4898.928.WhitingPF,RutjesAW,WestwoodME,etal.QUADAS-2:arevisedtoolfor
thequalityassessmentofdiagnosticaccuracystudies.AnnInternMed.2011;155:529–536.929.WhitingP,SavovicJ,HigginsJP,etal.ROBIS:anewtooltoassessriskofbiasinsystematicreviewswasdeveloped.JClinEpidemiol.2016;69:225–234.930.DerSimonianR,LairdN.Meta-analysisinclinicaltrials.ControlClinTrials.1986;7:177–188.931.FreemanMF,TukeyJ.Transformationsrelatedtotheangularandthe
squareroot.AnnMathStat.1950;21:601–611.932.NewcombeRG.Two-sidedcondenceintervalsforthesingleproportion:comparisonofsevenmethods.StatMed.1998;17:857–872.933.HigginsJP,Tho

---

### Chunk 21/30
**Article:** Suplementação I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.524

rar interações com anticoagulantes cumarínicos; detalhamento em cardiologia futura.
## Orientações Práticas e Ferramentas Sugeridas
- Mapas e fluxogramas: objetivos da aula; decisão de via (sabor/estabilidade/rapidez); quando considerar K2 com D; checklist de qualidade de suplementos (laudos, certificações, fornecedores, preço, biodisponibilidade).
- Visuals: imagens de cromossomos/telômeros; gráficos de biomarcadores (PCR/IL-6) em resveratrol; comparativos “manipulado confiável vs produto pronto”; mapas bioquímicos de aminoácidos → neurotransmissores.
- Comunicação ética: frases-modelo para recusar infusões sem indicação; foco em necessidade clínica e transparência de custo-benefício.
## Perguntas dos Alunos
Nenhuma pergunta foi levantada pelos alunos.

---

## Concept Insights

Não foram identificados conceitos novos

---

### Chunk 22/30
**Article:** Re-evaluation of Hematocrit as a Determinant of Thrombotic Risk in Erythrocytosis (2019)
**Journal:** Haematologica
**Section:** abstract | **Similarity:** 0.521

Análise crítica do hematócrito como fator de risco trombótico, revisando evidências sobre o limiar terapêutico de 45% e outros fatores de risco associados.

---

### Chunk 23/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.519

ngaCKDparametertothe
CHA2DS2-VAScscore,whichalreadyincludestheseparameters,wouldbeexpectedtohavelittleeffect.Thereis
considerablescopetoimprovethepredictiveperformance
ofthromboprophylaxisriskscoresforuseinCKD.713Recommendation3.16.1:Werecommenduseofnon–vitaminKantagonistoralanticoagulants(NOACs)inpreferencetovitaminKantagonists(e.g.,
warfarin)forthromboprophylaxisinatrialbrilla-tioninpeoplewithCKDG1–G4(1C).ThisrecommendationputshighvalueontheuseofNOACs,alsoreferredtoasdirect-actingoralanticoagulantsorDOACs,inpeoplewithCKDduetotheirsimplerpharmacokineticprole,dosing,andmonitoringthanvitaminKantagonistsandduetotheirimprovedefcacyandrelativelysimilarsafetyprole.AlthoughpeoplewithCKDG4–G5havebeenunderstudiedinRCTs,implementationinsuchgroupscanbeachievedafter
consideringchoiceofNOACanddosing.KeyinformationBalanceofbenetsandharms.Benets.Datafrom42,411participantswhoreceivedNOACsand29,272participants
whoreceivedwarfarinin4phaseIIItrialsweremeta-analyzed
in2014.Suchtrialslargelyex

---

### Chunk 24/30
**Article:** Confounders in Identification and Analysis of Inflammatory Biomarkers in Cardiovascular Diseases (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.517

atherosclerosis:Resultsfromaprospective,parallel-groupcohortstudy.Clin.Chim.Acta2015,447,16–22.[CrossRef]101.Hijazi,Z.;Lindahl,B.;Oldgren,J.;Andersson,U.;Lindbäck,J.;Granger,C.B.;Alexander,J.H.;Gersh,B.J.;Hanna,M.;Harjola,V.;etal.RepeatedMeasurementsofCardiacBiomarkersinAtrialFibrillationandValidationoftheABCStrokeScoreOverTime.J.Am.HeartAssoc.2017,6,e004851.[CrossRef]102.Abramson,J.L.;Lewis,C.;Murrah,N.V.;Anderson,G.T.;Vaccarino,V.RelationofC-ReactiveProteinandTumorNecrosisFactor-AlphatoAmbulatoryBloodPressureVariabilityinHealthyAdults.Am.J.Cardiol.2006,98,649–652.[CrossRef]103.Simundic,A.-M.;Kackov,S.;Miler,M.;Fraser,C.G.;Petersen,P.H.TermsandSymbolsUsedinStudiesonBiologicalVariation:TheNeedforHarmonization.Clin.Chem.2015,61,438–439.[CrossRef]
	
	
Getulio José Mattos Do Amaral Filho - getfilho@yahoo.com.br - CPF: 034.983.039-88

Biomolecules2021,11,1464
17of17
104.Khuseyinova,N.;Greven,S.;Rückerl,R.;Trischler,G.;Loewel,H.;Peters,A.;

---

### Chunk 25/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.514

thCKDwithacuteorunstablecoronarydisease,unacceptablelevelsofangina(e.g.,patientdissatisfaction),leftventricularsystolicdysfunctionattributabletoischemia,orleftmaindisease.
summaryofrecommendationstatementsandpracticepointswww.kidney-international.orgS162KidneyInternational(2024)105(Suppl4S),S117–S314

3.16CKDandatrialbrillationPracticePoint3.16.1:Followestablishedstrategiesforthediagnosisandmanagementofatrialbrillation(Figure40).Recommendation3.16.1:Werecommenduseofnon–vitaminKantagonistoralanticoagulants(NOACs)inpreferencetovitaminKantagonists(e.g.,warfarin)forthromboprophylaxisinatrialbrillationinpeoplewithCKDG1–G4(1C).PracticePoint3.16.2:NOACdoseadjustmentforGFRisrequired,withcautionneededatCKDG4–G5.PracticePoint3.16.3:DurationofNOACdiscontinuationbeforeelectiveproceduresneedstoconsiderproceduralbleedingrisk,NOACprescribed,andlevelofGFR(Figure44).

---

### Chunk 26/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.511

ncomparedwithwarfarininpatientswithatrialbrillationinrelationtorenalfunctionovertime—aRE-LYtrialanalysis.AmHeartJ.2018;198:169–177719;HijaziZ,AlexanderJH,LiZ,etal.ApixabanorvitaminKantagonistsandaspirinorplaceboaccordingtokidneyfunctioninpatientswithatrialbrillationafteracutecoronarysyndromeorpercutaneouscoronaryintervention:insightsfromtheAUGUSTUStrial.Circulation.2021;143:1215–1223722;HoriM,MatsumotoM,TanahashiN,etal.SafetyandefcacyofadjusteddoseofrivaroxabaninJapanesepatientswithnon-valvularatrialbrillation:subanalysisofJ-ROCKETAFforpatientswithmoderaterenalimpairment.CircJ.2013;77:632–638720;StaniferJ,PokorneyS,ChertowG,etal.Apixabanversuswarfarininpatientswithatrialbrillationandadvancedchronickidneydisease.Circulation.2020;141:1384–1392.721CI,condenceinterval;CrCl,creatinineclearance;eGFR,estimatedglomerularltrationrate.

---

### Chunk 27/30
**Article:** Confounders in Identification and Analysis of Inflammatory Biomarkers in Cardiovascular Diseases (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.509

lresearch.Curr.Opin.Clin.Nutr.Metab.Care2010,13,541–547.[CrossRef]92.Levi,M.;VanDerPoll,T.Inﬂammationandcoagulation.Crit.CareMed.2010,38,S26–S34.[CrossRef]93.Cemin,R.;Daves,M.Pre-analyticvariabilityincardiovascularbiomarkertesting.J.Thorac.Dis.2015,7,E395–E401.[CrossRef][PubMed]94.Rodríguez,A.D.;González,P.A.DiurnalVariationsinBiomarkersUsedinCardiovascularMedicine:ClinicalSigniﬁcance.Rev.Esp.Cardiol.2009,62,1340–1341.[CrossRef]95.Rudnicka,A.R.;Rumley,A.;Lowe,G.D.;Strachan,D.P.Diurnal,Seasonal,andBlood-ProcessingPatternsinLevelsofCirculatingFibrinogen,FibrinD-Dimer,C-ReactiveProtein,TissuePlasminogenActivator,andvonWillebrandFactorina45-Year-OldPopulation.Circulation2007,115,996–1003.[CrossRef][PubMed]96.Dominguez-Rodriguez,A.;Tome,M.C.-P.;Abreu-Gonzalez,P.Interrelationbetweenarterialinﬂammationinacutecoronarysyndromeandcircadianvariation.WorldJ.Cardiol.2011,3,57–58.[CrossRef]97.Tirumalai,R.S.;Chan,K.C.;Prieto,D.A.;Issaq,H.J.;Conrads,T.P.;Veenstra,T.D.Characterizati

---

### Chunk 28/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** conclusion | **Similarity:** 0.506

highrisk.Lowriskisdenedasalowfrequencyofbleedingand/orminorimpactofableed.Highriskisdenedasahighfrequencyofbleedingand/orimportantclinicalimpact.AdaptedfromHeidbuchelH,VerhammeP,AlingsM,etal.UpdatedEuropeanHeartRhythmAssociationpracticalguideontheuseofnon–vitamin-Kantagonistanticoagulantsinpatientswithnon-valvularatrialbrillation:executivesummary.EurHeartJ.2017;38:2137–2149.724aManyofthesepeoplemaybeonlowerdoseofdabigatran(110mgtwiceperday[b.i.d.])orapixaban(2.5mgb.i.d.),orhavetobeonthelowerdoseofrivaroxaban(15mgQD)oredoxaban(30mgQD).Dabigatran110mgb.i.d.hasnotbeenapprovedforusebytheUSFoodandDrug
Administration.CrCl,creatinineclearance,LMWH,low-molecular-weightheparin;UFH,unfractionatedheparin.ReproducedfromTurakhiaMP,
BlankestijnPJ,CarreroJJ,etal.Chronickidneydiseaseandarrhythmias:conclusionsfromaKidneyDisease:ImprovingGlobalOutcomes
(KDIGO)ControversiesConference.EurHeartJ.2018;39:2314–2325.710ªTheAuthor(s)2018.PublishedbyOxfordUniversityPressonbehalfoftheEuropeanSocietyofC

---

### Chunk 29/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.505

4S),S117–S314S309

advancedchronickidneydisease].Kardiologiia.2020;60:1322[inRussian].718.FoxKA,PicciniJP,WojdylaD,etal.Preventionofstrokeandsystemicembolismwithrivaroxabancomparedwithwarfarininpatientswith
non-valvularatrialbrillationandmoderaterenalimpairment.EurHeartJ.2011;32:2387–2394.719.HijaziZ,HohnloserSH,OldgrenJ,etal.Efcacyandsafetyofdabigatrancomparedwithwarfarininpatientswithatrialbrillationinrelationtorenalfunctionovertime—aRE-LYtrialanalysis.AmHeartJ.2018;198:169–177.720.HoriM,MatsumotoM,TanahashiN,etal.SafetyandefcacyofadjusteddoseofrivaroxabaninJapanesepatientswithnon-valvularatrial
brillation:subanalysisofJ-ROCKETAFforpatientswithmoderaterenalimpairment.CircJ.2013;77:632–638.721.StaniferJ,PokorneyS,ChertowG,etal.Apixabanversuswarfarinin
patientswithatrialbrillationandadvancedchronickidneydisease.Circulation.2020;141:1384–1392.722.HijaziZ,AlexanderJH,LiZ,etal.ApixabanorvitaminKantagonistsandaspirinorplaceboaccordingtokidneyfunctioninpatientswithatrialb

---

### Chunk 30/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.503

anahashiN,etal.SafetyandefcacyofadjusteddoseofrivaroxabaninJapanesepatientswithnon-valvularatrialbrillation:subanalysisofJ-ROCKETAFforpatientswithmoderaterenalimpairment.CircJ.2013;77:632–638720;StaniferJ,PokorneyS,ChertowG,etal.Apixabanversuswarfarininpatientswithatrialbrillationandadvancedchronickidneydisease.Circulation.2020;141:1384–1392721;HijaziZ,AlexanderJH,LiZ,etal.ApixabanorvitaminKantagonistsandaspirinorplaceboaccordingtokidneyfunctioninpatientswithatrialbrillationafteracutecoronarysyndromeorpercutaneouscoronaryinterventioninsightsfromtheAUGUSTUStrial.Circulation.2021;143:1215–1223.722CI,condenceinterval;CrCl,creatinineclearance;eGFR,estimatedglomerularltrationrate.

---

