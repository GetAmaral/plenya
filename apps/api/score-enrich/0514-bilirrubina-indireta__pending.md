# ScoreItem: Bilirrubina indireta

**ID:** `019bf31d-2ef0-7b39-9f75-7306ee5c8c46`
**FullName:** Bilirrubina indireta (Exames - Laboratoriais)
**Unit:** mg/dL

**Preparation Metadata:**
- Quality Grade: **EXCELLENT**
- Total Chunks: 30 de 5 artigos
- Avg Similarity: 0.627

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-7b39-9f75-7306ee5c8c46`.**

```json
{
  "score_item_id": "019bf31d-2ef0-7b39-9f75-7306ee5c8c46",
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

**ScoreItem:** Bilirrubina indireta (Exames - Laboratoriais)
**Unidade:** mg/dL

**30 chunks de 5 artigos (avg similarity: 0.627)**

### Chunk 1/30
**Article:** Bilirubin metabolism and jaundice: advances in understanding and clinical implications (2021)
**Journal:** Seminars in Liver Disease
**Section:** abstract | **Similarity:** 0.720

Bilirubin, the end product of heme catabolism, has long been considered merely a waste product.
Recent evidence demonstrates that mildly elevated unconjugated bilirubin levels exert beneficial antioxidant and
anti-inflammatory effects. This review discusses the complete bilirubin metabolism pathway, from heme degradation
through hepatic conjugation and biliary excretion. Understanding the differences between unconjugated (indirect)
and conjugated (direct) hyperbilirubinemia is essential for proper differential diagnosis. Hemolytic disorders,
Gilbert syndrome, and Crigler-Najjar syndrome cause predominantly unconjugated hyperbilirubinemia, while hepatocellular
disease and cholestasis result in conjugated hyperbilirubinemia. The clinical approach includes fractionation of
total bilirubin and assessment of associated liver function tests.

---

### Chunk 2/30
**Article:** Clinical interpretation of bilirubin levels: when and how to measure fractions (2020)
**Journal:** American Family Physician
**Section:** abstract | **Similarity:** 0.696

Total bilirubin measurement is a routine component of liver function testing. However,
fractionation into direct (conjugated) and indirect (unconjugated) bilirubin provides critical diagnostic
information. Total bilirubin levels above 2.5 mg/dL typically warrant fractionation. Indirect hyperbilirubinemia
suggests hemolysis, ineffective erythropoiesis, or inherited disorders like Gilbert syndrome. Direct hyperbilirubinemia
indicates hepatocellular dysfunction or biliary obstruction. This article provides a practical algorithmic approach
to hyperbilirubinemia evaluation, emphasizing cost-effective testing strategies and appropriate specialist referral.

---

### Chunk 3/30
**Article:** Measurement and clinical usefulness of bilirubin in liver disease (2021)
**Journal:** Advances in Laboratory Medicine
**Section:** methods | **Similarity:** 0.686

ReviewArmandoRaúlGuerraRuiz*,JavierCrespo,RosaMariaLópezMartínez,PaulaIruzubieta,GregoriCasalsMercadal,MartaLalanaGarcés,BernardoLavinandManuelMoralesRuizMeasurementandclinicalusefulnessofbilirubininliverdiseasehttps://doi.org/10.1515/almed-2021-0047ReceivedDecember2,2020;acceptedFebruary16,2021;
publishedonlineJuly9,2021Abstract:Elevatedplasmabilirubinlevelsareafrequentclinicalfinding.Itcanbesecondarytoalterationsinanystageofitsmetabolism:(a)excessbilirubinproduction(i.e.,pathologichemolysis);(b)impairedliveruptake,withelevationofindirectbilirubin;(c)impairedconjugation,promptedbyadefectintheUDP-glucuronosyltransferase;and(d)bileclearancedefect,withelevationofdirectbili-rubinsecondarytodefectsinclearanceproteins,orinabilityofthebiletoreachthesmallbowelthroughbileducts.Aliverlesionofanycausereduceshepatocytecellnumberandmayimpairtheuptakeofindirectbilirubinfromplasmaanddiminishdirectbilirubintransportandclearance
throughthebileducts.Variousanalyticalmethodsare
currentlyavailablefor

---

### Chunk 4/30
**Article:** Measurement and clinical usefulness of bilirubin in liver disease (2021)
**Journal:** Advances in Laboratory Medicine
**Section:** other | **Similarity:** 0.668

holia).Bilirubindoesnotincreaseinurine(nocholuria),sinceUCBisnotwater-solubleandisnotfilteredbythekidneys.Medicationssuchasrifampicin,chloramphenicol,andprobenecidmayinduceunconjugatedhyperbilirubinemia,sinceitcompeteswiththetransporterthatcarriesbilirubinintothehepatocyte.TheelevationofUCBisalsoassociatedwithinheriteddisordersaffectingconjugation,ofwhichGilbertsyndromeisthemostcommoninadults,affecting3–10%ofthepopulation.UDPGTactivityislow.Thissyndromedoesnot
requirefollow-uportreatment,sinceitisabenigndisorder.However,itmaybeaconfoundingfactorwhenscreeningforliverdiseaseandisfrequentlymisdiagnosedaschronichepatitis[3].Whenthegeneticdefectdirectlyaffectsenzymepro-duction,itcausesCrigler–Najjarsyndrome(SCN).SCNtype1ischaracterizedbyatotalenzymedeficiencythatdoesnotimprovewithinductiontherapywithpheno-barbital.Thissubtypemaybelife-threateningduetothe
neurologicaltoxicitysecondarytothedepositionofbili-rubininthebasalgangliaandnucleiofthebrainstem(neonatalkernicterus)[36].InSCNtype2,e

---

### Chunk 5/30
**Article:** Measurement and clinical usefulness of bilirubin in liver disease (2021)
**Journal:** Advances in Laboratory Medicine
**Section:** other | **Similarity:** 0.650

acute-on-chronicliverfailure:animportantprognosticindicator.AnnHepatol2014;13:98–104.47.BreimerLH,WannametheeG,EbrahimS,ShaperAG.Serumbilirubinandriskofischemicheartdiseaseinmiddle-agedBritishmen.ClinChem1995;41:1504–8.48.KoGTC,ChanJCN,WooJ,LauE,YeungVTF,ChowC-C,etal.SerumbilirubinandcardiovascularriskfactorsinaChinesepopulation.JCardiovascRisk1996;3:459–63.49.MadhavanM,WattigneyWA,SrinivasanSR,BerensonGS.Serumbilirubindistributionanditsrelationtocardiovascularriskinchildrenandyoungadults.Atherosclerosis1997;131:107–13.50.LinJP.AssociationbetweentheUGT1A1*28allele,bilirubinlevels,andcoronaryheartdiseaseintheFraminghamHeartStudy.Circulation2006;114:1476–81.51.ItekL,JirsaM,BrodanovaM,KalabM,MarecekZ,DanzigV,etal.Gilbertsyndromeandischemicheartdisease:aprotectiveeffectofelevatedbilirubinlevels.Atherosclerosis2002;160:449–56.52.VítekL.Theroleofbilirubinindiabetes,metabolicsyndrome,andcardiovasculardiseases.FrontPharmacol2012;3:1–7.ArticleNote:Theoriginalarticlecanbefoundhere

---

### Chunk 6/30
**Article:** Measurement and clinical usefulness of bilirubin in liver disease (2021)
**Journal:** Advances in Laboratory Medicine
**Section:** methods | **Similarity:** 0.640

humeauxD.Inheriteddisordersofbilirubintransportandconjugation:newinsightsintomolecular
mechanismsandconsequences.Gastroenterology2014;146:1625–38.12.RowlandA,MinersJO,MackenziePI.TheUDP-glucuronosyltransferases:theirroleindrugmetabolismanddetoxication.IntJBiochemCellBiol2013;45:1121–32.13.JemnitzK,Heredi-SzaboK,JanossyJ,IojaE,VereczkeyL,KrajcsiP.ABCC2/Abcc2:amultispecictransporterwithdominantexcretoryfunctions.DrugMetabRev2010;42:402–36.14.PellockSJ,RedinboMR.Glucuronidesinthegut:sugar-drivensymbiosesbetweenmicrobeandhost.JBiolChem2017;292:
8569–76.15.NgashangvaL,BachuV,GoswamiP.Developmentofnewmethodsfordeterminationofbilirubin.JPharmaceutBiomed
Anal2019;162:272–85.16.KwoPY,CohenSM,LimJK.ACGclinicalguideline:evaluationofabnormalliverchemistries.AmJGastroenterol2017;112:18–35.17.ColePG,LatheGH,BillingBH.Thediazoreactingpigmentsofserum,urineandbile.BiochemJ1953;55:xiii.18.KuenzleCC,MaierC,RüttnerJR.Thenatureoffourbilirubinfractionsfromserumandofthreebilirubinfractionsfrom

---

### Chunk 7/30
**Article:** Measurement and clinical usefulness of bilirubin in liver disease (2021)
**Journal:** Advances in Laboratory Medicine
**Section:** other | **Similarity:** 0.638

17;87:e1–294.e8.37.WolkoffAW,KetleyJN,WaggonerJG,BerkPD,JakobyWB.Hepaticaccumulationandintracellularbindingofconjugatedbilirubin.JClinInvest1978;61:142–9.38.Méndez-SánchezN,VítekL,Aguilar-OlivosNE,UribeM.Bilirubinasabiomarkerinliverdisease.In:Biomarkersinliverdisease.Dordrecht:Springer;2017:281–304pp.39.RaymondGD,GalambosJT.Hepaticstorageandexcretionofbilirubininman.AmJGastroenterol1971;55:135–44.40.SticovaE,JirsaM.Newinsightsinbilirubinmetabolismandtheirclinicalimplications.WorldJGastroenterol2013;19:
6398–407.41.RutherfordA,KingLY,HynanLS,VedvyasC,LinW,LeeWM,etal.Developmentofanaccurateindexforpredictingoutcomesof
patientswithacuteliverfailure.Gastroenterology2012;143:
1237–43.42.WlodzimirowKA,EslamiS,Abu-HannaA,NieuwoudtM,ChamuleauRAFM.Asystematicreviewonprognosticindicators360
GuerraRuizetal.:Bilirubin:measurementandclinicalusefulness

ofacuteonchronicliverfailureandtheirpredictivevalueformortality.LiverInt2013;33:40–52.43.HelmkeS,ColmeneroJ,EversonGT.Noninvasiveas

---

### Chunk 8/30
**Article:** Measurement and clinical usefulness of bilirubin in liver disease (2021)
**Journal:** Advances in Laboratory Medicine
**Section:** other | **Similarity:** 0.638

rubinlevels.Atherosclerosis2002;160:449–56.52.VítekL.Theroleofbilirubinindiabetes,metabolicsyndrome,andcardiovasculardiseases.FrontPharmacol2012;3:1–7.ArticleNote:Theoriginalarticlecanbefoundhere:https://doi.org/10.1515/almed-2020-0016.
GuerraRuizetal.:Bilirubin:measurementandclinicalusefulness
361

---

### Chunk 9/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** results | **Similarity:** 0.634

with and 
without hyperbilirubinemia. Dig Dis Sci 2008;53:799-
802.163. Lazo M, Selvin E, Clark JM. Brief communication: 
clinical implications of short-term variability in liver 
function test results. Ann Intern Med 2008;148:348-52.164. Schmidt E, Schmidt FW, Chemnitz G, Kubale R, 
Lobers J. The Szasz-ratio (CK/GOT) as example for the 
diagnostic significance of enzyme ratios in serum. Klin 
Wochenschr 1980;58:709-18.165. Dufour DR. Is it necessary to order aspartate 
aminotransferase with alanine aminotransferase 
in clinical practice? Author’s Reply. Clin Chem 
2001;47:1134-5.

---

### Chunk 10/30
**Article:** Measurement and clinical usefulness of bilirubin in liver disease (2021)
**Journal:** Advances in Laboratory Medicine
**Section:** results | **Similarity:** 0.633

noglucuronide,andbilirubindiglucuronide(indirect-reactionfractions)[18].Afourthfractionresultsfromcovalentbindingofbilirubintoprotein(δ-bilirubin),whichisdifferentfromthebilirubin–albumincomplexinserum[1].DiazomethodThereactionofbilirubinwiththediazoreagentrenderstwocolorazodipyrroles(azopigments)(Figure1(D))thatcanbemeasuredbyspectrophotometry,at530nmtoneutraloracidpH,andat598nmtoalkalinepH(i.e.,bytheadditionofalkalinetartrate).Thisreactionisacceleratedbyalcoholand
avarietyofothercomponents(i.e.,sodiumbenzoate)causingUCBtodissociatefromalbumin[19].Inthepresenceofan‘accelerator’,conjugatedandunconjugatedbilirubinarejointlymeasured(totalbilirubin),whereasintheabsenceofanaccelerator,onlyCBreacts(‘directbilirubin’).ThedifferencebetweenTotalandCByieldsUCBcon-centration(‘indirectbilirubin’).Forthemethodtobeaccurate,itiscrucialthatminimumamountsofUCBreactinthedirectprocedure.ThediazomethoddescribedbyJendrassik&Grofin1938[20]andlatermodiedbyDoumasetal.[21]yieldstotalserumbilir

---

### Chunk 11/30
**Article:** Measurement and clinical usefulness of bilirubin in liver disease (2021)
**Journal:** Advances in Laboratory Medicine
**Section:** results | **Similarity:** 0.633

e-hepatic,hepatic,orpost-hepatic.Pre-hepatichyperbilirubinemiaThetermreferstohyperbilirubinemiasecondarytoexcessbilirubinproduction.Themostcommoncauseisacceler-atedhemolysis.Whenanelevatedbilirubinproduction
rateexceedstheuptakeandexcretioncapacityoftheliver,itresultsinelevatedserumUCBconcentrations,whereasCBcanbenormalorslightlyelevated.Identifyinghemo-lysisasthecauseofhyperbilirubinemiaisnotchallenging,sincethepatientwillexhibitothernumerouscueingsigns(anemia,elevatedreticulocytes,amongothers)[1].Sinceelevatedbilirubinisnotinducedbyliverdamage,testingwillnotshowalterationsinaminotransferases,albumin,orprothrombinactivity.HepatichyperbilirubinemiasThetermreferstoconditionsdirectlyrelatedtoliver
function.Theseconditionsmayaffectbilirubinuptake,metabolism,conjugation,and/orexcretion,andtestswillshowelevatedCBand/orUCBconcentrations.Thesediseasesareassociatedwithconcomitantliverlesionsofdifferentseveritythatmaycompromiseliverfunctionandmightfollowanacuteorchroniccourse.Espe-ciallyintheca

---

### Chunk 12/30
**Article:** Measurement and clinical usefulness of bilirubin in liver disease (2021)
**Journal:** Advances in Laboratory Medicine
**Section:** results | **Similarity:** 0.630

lirubin’).Forthemethodtobeaccurate,itiscrucialthatminimumamountsofUCBreactinthedirectprocedure.ThediazomethoddescribedbyJendrassik&Grofin1938[20]andlatermodiedbyDoumasetal.[21]yieldstotalserumbilirubinresultswhicharereproducibleandreliable.In
thismethod,theacceleratorisacaffeineandsodiumbenzoatesolution.Thismethodhasacceptableinter-laboratorytrans-ferabilityandiscurrentlythegold-standardmethod[15,21–23].ItstruenesstomeasuretotalanddirectbilirubinhasbeenassessedbycomparingwithUCBandbilirubindiglu-curonidequantiedbynuclearmagneticresonance.ChromatographyHighperformanceliquidchromatographyhasbeenusedtomeasuretotalserumbilirubin(TB)aftertheadditionofthefourindividualfractionsorspeciesmentionedabove(unconjugated,monoanddiglucuronide,anddelta-bilirubin).TBvaluesmeasuredbyHPLCareconsistentwiththoseobtainedusingtheJendrassik–Grofmethod[24].HPLCenabledtheidenticationofthetypeofbilirubinthatenduresonceinitialliverdiseasehasbeensolved(delta-bilirubin,whichhasalongerhalf-lifethantheotherf

---

### Chunk 13/30
**Article:** Measurement and clinical usefulness of bilirubin in liver disease (2021)
**Journal:** Advances in Laboratory Medicine
**Section:** methods | **Similarity:** 0.625

ismandrecirculation.354
GuerraRuizetal.:Bilirubin:measurementandclinicalusefulness

lowergastrointestinaltract,thethreeurobilinogensspon-taneouslyoxidizetoproducetheequivalentbrownish-yellowbilepigmentsstercobilinogen,urobilinogen,andmesobilinogen,givingstoolitsdistinctivecolor.Upto20%oftheurobilinogenproduceddailyisreabsorbedfromtheintestineandundergoesenterohepaticrecirculation.Mostofthisurobilinogenistakenupbytheliver(viatheportalvein)andisreexcretedintobile,whereasasmallfraction(2–5%)isexcretedintothegeneralcirculationandlteredbythekidneys,beingdetectableinurine.TestingmethodsVariousanalyticalmethodsarecurrentlyavailableformeasuringbilirubinanditsmetabolitesinserum,urine,andfeces.Serumbilirubinisdeterminedby:(i)diazotransferreaction,currently,thegold-standardmeasurementmethod;(ii)high-performanceliquidchromatography(HPLC);(iii)oxidative,enzymatic,andchemicalmethods;(iv)directspectrophotometry;and(v)transcutaneousmethods[15].Sinceitwasfirstdescribedbytheendofthe19thcentury,themos

---

### Chunk 14/30
**Article:** Measurement and clinical usefulness of bilirubin in liver disease (2021)
**Journal:** Advances in Laboratory Medicine
**Section:** other | **Similarity:** 0.625

tedresponsibilityfortheentirecontentofthismanuscriptandapproveditssubmission.Competinginterests:Authorsstatenoconictofinterest.References1.CappelliniMD,LoSF,SwinkelsDW.38–Hemoglobin,iron,bilirubin.In:Tietztextbookofclinicalchemistryandmoleculardiagnostics,6thed.St.Louis,MO,USA:ElsevierInc.;2017.https://doi.org/10.1016/B978-0-323-35921-4.00038-7.2.Méndez-SánchezN,QiX,VitekL,ArreseM.Evaluatinganoutpatientwithanelevatedbilirubin.AmJGastroenterol2019;114:1185–8.3.DufourDR,LottJA,NolteFS,GretchDR,KoffRS,SeeffLB.Diagnosisandmonitoringofhepaticinjury.II.Recommendationsforuseoflaboratorytestsinscreening,diagnosis,andmonitoring.Clin
Chem2000;46:2050–68.4.FeveryJ.Bilirubininclinicalpractice:areview.LiverInt2008;28:592–605.5.VítekL.Bilirubinasapredictorofdiseasesofcivilization.Isittimetoestablishdecisionlimitsforserumbilirubinconcentrations?ArchBiochemBiophys2019;672:108062.

---

### Chunk 15/30
**Article:** Measurement and clinical usefulness of bilirubin in liver disease (2021)
**Journal:** Advances in Laboratory Medicine
**Section:** introduction | **Similarity:** 0.622

ubin;cholestasis;diazomethod;liverdisease.IntroductionBilirubinisanorange-yellowpigmentofbilethatresultsfromthedegradationofvariousheme-containingproteins,
especiallyfromhemoglobincatabolism.Hemeisbroken
downintobiliverdin,whichisconvertedintounconju-
gatedorindirectbilirubin(UCB).UCBiswater-insoluble
andenterscirculationboundtoalbumin.Intheliver,
glucuronicacidisaddedtoUCB(conjugation)torenderit
water-soluble(directbilirubin);finally,itiseitherexcreted
intobileorrecirculatedbacktothebloodstream,whereitis
filtratedbythekidneysandexcretedthroughurine[1].Elevationofplasmabilirubinlevelsisafrequentfindingbothinprimary[2]andhospitalcare.Allliverle-
sionsinduceadecreaseinthehepatocytecellcount,which
maycausehyperbilirubinemia[3].Hyperbilirubinemiacan
originatefromanalterationinanystageofbilirubinmetabolism:excessproduction,impairedliveruptake,conjugationdefects,orbiliaryexcretiondefects[4].Bilirubinisawell-establishedmarkerthatisroutinelyincludedinbiochemicaltestsforpatientswithliver
dysfun

---

### Chunk 16/30
**Article:** Measurement and clinical usefulness of bilirubin in liver disease (2021)
**Journal:** Advances in Laboratory Medicine
**Section:** other | **Similarity:** 0.621

thantherateofbili-rubinproduction[39].However,hyperbilirubinemiaisalong-establishedmarkerofliverandbilealterations,andhasprognosticvalueincertainliverdiseases[3,40].Inthehyperacutestageofacuteliverfailure,bilirubinconcentrationisrelativelylowascomparedtothesub-stantialelevationofplasmaaminotransferaseconcentra-tionsinplasma.However,inthesubacutestage,the358
GuerraRuizetal.:Bilirubin:measurementandclinicalusefulness

situationreverses[41].Inthiscase,elevatedlevelsofbili-rubininplasmaareanindicatorofpoorprognosisandmortality[42].Hyperbilirubinemiadoesnothaveaprognosticvalueinpatientswithacutehepatitisinducedbyparacetamol,butitdoesinacuteandsubacutehepatitisinducedbyothercauses[43].Bilirubinconcentrations>17.6mg/dLisanindicationforhospitalizationinpatientswithacute
hepatitisunrelatedtotheintakeofparacetamol[44].Hepaticcirrhosiscanbeaccompaniedbyprogressivebilirubinelevations.Increasedbilirubinconcentrationsarearelativelylateeventinchronicliverdiseaseandindicatesevereliverdysfunction[4].In

---

### Chunk 17/30
**Article:** Measurement and clinical usefulness of bilirubin in liver disease (2021)
**Journal:** Advances in Laboratory Medicine
**Section:** abstract | **Similarity:** 0.620

Elevatedplasmabilirubinlevelsareafrequentclinicalfinding.Itcanbesecondarytoalterationsinanystageofitsmetabolism:(a)excessbilirubinproduction(i.e.,pathologichemolysis);(b)impairedliveruptake,withelevationofindirectbilirubin;(c)impairedconjugation,promptedbyadefectintheUDP-glucuronosyltransferase;and(d)bileclearancedefect,withelevationofdirectbili-rubinsecondarytodefectsinclearanceproteins,orinabilityofthebiletoreachthesmallbowelthroughbileducts.Aliverlesionofanycausereduceshepatocytecellnumberandmayimpairtheuptakeofindirectbilirubinfromplasmaanddiminishdirectbilirubintransportandclearance
throughthebileducts.Variousanalyticalmethodsare
currentlyavailableformeasuringbilirubinanditsmetabo-
litesinserum,urineandfeces.Serumbilirubinisdeter-
minedby(1)diazotransferreaction,currently,thegold-
standard;(2)high-performanceliquidchromatography
(HPLC);(3)oxidative,enzymatic,andchemicalmethods;(4)
directspectrophotometry;and(5)transcutaneousmethods.
Althoughbilirubinisawell-establishedmarkerofliver
function,itdoesnotalwaysidentifyalesioninthisorgan.Therefore,foraccuratediagnosis,alterationsinbilirubinconcentrationsshouldbeassessedinrelationtopatient
anamnesis,thedegreeofthealteration,andthepatternof
concurrentbiochemicalalterations.

---

### Chunk 18/30
**Article:** BilR is a gut microbial enzyme that reduces bilirubin to urobilinogen (2024)
**Journal:** Nature Microbiology
**Section:** results | **Similarity:** 0.618

_12Bacteroides cellulosilyticus CL02T12C19Bacteroides finegoldii CL09T03C10E. coli BL21E. coli 5­alphaClostridium citroniae WAL­17109Clostridium innocuum 6_1_30L. reuteri CF48­3ARuminococcus gnavus CC55_001CClostridium bolteae CC43_001BClostridioides diicile P11Clostridioides diicile Isolate 7Clostridioides diicile CD3Clostridioides diicile P720406080Sample­to­blank ratioStercobilinUrobilinConjugatedbilirubinBilirubinUrobilinogenHMOX1BLVRABiliverdinUGT1A1
Fig. 1 | Identification of bilirubin-reducing bacterial strains. a, Illustrated representation of the haem degradation pathway. Key human enzymes are labelled with grey text. b, Diagram of the structures of bilirubin and urobilinogen. The bonds reduced during bilirubin reduction are shown in red. c, Results of fluorescence assay screening of bacterial strains. Measurements from n=3 independent biological replicates are shown as black points.

---

### Chunk 19/30
**Article:** Measurement and clinical usefulness of bilirubin in liver disease (2021)
**Journal:** Advances in Laboratory Medicine
**Section:** other | **Similarity:** 0.613

ia)whereasurinehasexcesscolor(choluria)andurineurobilinogenisdecreased[1].Extra-hepaticcholestasiscanbecausedbyatotalorpartialphysicalobstructionofextrahepaticbileducts.Themostcommoncausesinclude:choledocholithiasis,extrinsiccompressionsofthebileduct(pancreaticneoplasia,Mir-
izzisyndrome),disordersoftheextrahepaticbileducts(cholangiocarcinoma,primaryorsecondarysclerosingcholangitis),andinfections(CMV,parasites).BilirubinasadiagnosticandprognosticmarkerInliverdiseaseAsmentionedabove,elevationofbilirubinconcentrationscanbeinducedbynumerouscausesandhence,itisanonspecificmarkerofliverdysfunction.Itisnotasensitivemarkerofliverinjuryeither:ahealthylivercanconjugatedailyUCBproductionuptotwotimeswithoutcausinganincreaseintotalbilirubinconcentrations.Also,therateof
bilirubinexcretionis10timeshigherthantherateofbili-rubinproduction[39].However,hyperbilirubinemiaisalong-establishedmarkerofliverandbilealterations,andhasprognosticvalueincertainliverdiseases[3,40].Inthehyperacutestageofacuteliverfai

---

### Chunk 20/30
**Article:** Measurement and clinical usefulness of bilirubin in liver disease (2021)
**Journal:** Advances in Laboratory Medicine
**Section:** methods | **Similarity:** 0.613

bratedirectbilirubinmeasurementmethods.Theproteinmatrixofthesecalibratorsishumanserum,bovineserum,oracombinationofthetwo.UCBinhumanserumreactstotallywiththegold-standardmethodandotherdiazo-basedsystemsavailableinclinicalanalyzers.However,itsreactioninbovineserumis
incompleteandunpredictable.Thislimitationmakesitimpossibletodetermineexactbilirubinvaluesinthecali-brationmaterialwhentheproteinmatrixisbovineserum[34].Accordingtotheliterature,ditaurobilirubininhumanserumwasunderestimatedbytwoofthesevenclinicalan-alyzerstested,whichusedcalibratorsbasedonbovineserum.Ditaurobilirubininbovineserumwasunder-estimatedbyallanalyzersandbythemethodofreference[33].Asaresult,calibratorsofbilirubinbasedonamatrixof356
GuerraRuizetal.:Bilirubin:measurementandclinicalusefulness

bovineserumshouldnotbeused,sincetheycompromisetheaccuracyofbilirubindeterminationtests.ClinicalsignificanceDiseasesordisordersinterferingwithbilirubinmetabolism
maycauseanelevationofserumbilirubinconcentrations.Elevatedbilirubinlev

---

### Chunk 21/30
**Article:** Measurement and clinical usefulness of bilirubin in liver disease (2021)
**Journal:** Advances in Laboratory Medicine
**Section:** methods | **Similarity:** 0.612

ycausereduceshepatocytecellnumberandmayimpairtheuptakeofindirectbilirubinfromplasmaanddiminishdirectbilirubintransportandclearance
throughthebileducts.Variousanalyticalmethodsare
currentlyavailableformeasuringbilirubinanditsmetabo-
litesinserum,urineandfeces.Serumbilirubinisdeter-
minedby(1)diazotransferreaction,currently,thegold-
standard;(2)high-performanceliquidchromatography
(HPLC);(3)oxidative,enzymatic,andchemicalmethods;(4)
directspectrophotometry;and(5)transcutaneousmethods.

---

### Chunk 22/30
**Article:** Measurement and clinical usefulness of bilirubin in liver disease (2021)
**Journal:** Advances in Laboratory Medicine
**Section:** results | **Similarity:** 0.612

hbilirubinandurobilinogendeterminationinserumandurine.Inthehospitalcontext,bilirubinconcen-trationsareveryusefulforprognosisofacuteliverdiseaseandmonitoringchronicliverdisease.Theseresultsmustbeinterpretedinthecontextofpatientanamnesis,degreeofalteration,andotherclinicallaboratoryparameters.Acknowledgments:TheauthorsthanktheSpanishSocietyofLaboratoryMedicine(SEQC-ML)andtheSpanishSocietyofDigestiveDiseases(SEPD)andtheirscientificboardsfortheirsupport.Researchfunding:1)MinisteriodeEconomíayCom-petitividad(PID2019-105502RB-I00aMM-R).2)ProyectoPI19/00774,nanciadoporelInstitutodeSaludCarlosIIIyconanciadoporlaUniónEuropea(FEDER).3)ProyectoPI15/02138,nanciadoporelInstitutodeSaludCarlosIII.4)ProyectoFISPI18/01304,nanciadoporelInstitutodeSaludCarlosIII.

---

### Chunk 23/30
**Article:** Measurement and clinical usefulness of bilirubin in liver disease (2021)
**Journal:** Advances in Laboratory Medicine
**Section:** other | **Similarity:** 0.606

inmetabolism:excessproduction,impairedliveruptake,conjugationdefects,orbiliaryexcretiondefects[4].Bilirubinisawell-establishedmarkerthatisroutinelyincludedinbiochemicaltestsforpatientswithliver
dysfunctionoranyothercondition.However,bilirubinisnot
*Correspondingauthor:ArmandoRaúlGuerraRuiz,CommissiononBiochemistryofLiverDisease,SEQCML,Barcelona,Spain;andServiceofClinicalBiochemistry,HospitalUniversitarioMarquésdeValdecilla,AvenidaValdecilla,s/n.CP39008,Santander,Spain,E-mail:a.raulguerra@gmail.com.https://orcid.org/0000-0001-
8896-8611JavierCrespoandPaulaIruzubieta,ServiceofGastroenterology,MarquésdeValdecillaUniversityHospital,Santander,Spain;and
ClinicalandTranslationalResearchGrouponDigestiveDiseases,IDIVAL.Santander,SpainRosaMariaLópezMartínez,CommissiononBiochemistryofLiverDisease,SEQCML,Barcelona,Spain;andUnitofLiverDisease,ServicesofBiochemistryandMicrobiology,HospitalUniversitariValld’Hebron,UniversitatAutònomadeBarcelona,Barcelona,SpainGregoriCasalsMercadal,Commissiono

---

### Chunk 24/30
**Article:** Measurement and clinical usefulness of bilirubin in liver disease (2021)
**Journal:** Advances in Laboratory Medicine
**Section:** other | **Similarity:** 0.605

etheycompromisetheaccuracyofbilirubindeterminationtests.ClinicalsignificanceDiseasesordisordersinterferingwithbilirubinmetabolism
maycauseanelevationofserumbilirubinconcentrations.Elevatedbilirubinlevelsinblood(>1mg/dL)[35]causebilirubindepositionintissues,especiallyinthosecontainingalargevolumeofelasticbers(palate,conjunctiva,amongothers).Substantialaccumulation(generallyabove2.5mg/dL)givesmucosaandskinayellowishcolor,whichisknownasjaundice.Hyperbilirubinemiabyitselfdoesnothaveapoorprognosis[5],sinceourbodyhaseffectivedetoxica-tionmechanisms(exceptforneonates).However,itisasign
ofimpairedbilirubinproductionormetabolism.Differentapproacheshavebeenadoptedfortheclassi-ficationofdisorderscausinghyperbilirubinemia.Basedontheirlocation,thedisordersinducinghyperbilirubinemiaareclassifiedintopre-hepatic,hepatic,orpost-hepatic.Pre-hepatichyperbilirubinemiaThetermreferstohyperbilirubinemiasecondarytoexcessbilirubinproduction.Themostcommoncauseisacceler-atedhemolysis.Whenanelevatedbilirubinpro

---

### Chunk 25/30
**Article:** Measurement and clinical usefulness of bilirubin in liver disease (2021)
**Journal:** Advances in Laboratory Medicine
**Section:** other | **Similarity:** 0.602

owsuppression.Advanceddiseasecausesincreasedbilir-ubinemia,generallyconjugated[16,43].Itshouldnotbeforgottenthatelevationofserumbili-rubindoesnotnecessarilyindicateliverfunctionstatus.Indeed,theearliestandmostaccuratemarkerofliverfailureisprothrombintimemeasuredusingtheinternationalnormalizedratio(INR),whichshouldalwaysbeincludedintheevaluationofacuteorchronicliverdisease.IncardiovasculardiseaseInthe1990s,solidevidencewaspublishedofastrongnegativecorrelationbetweenplasmabilirubinconcentra-tionsandtheriskforcoronaryarterydisease[47–49].Inthesameline,aslightincreaseinbilirubinconcentrationswasfoundtobeassociatedwithalowerriskforatherosclerosis
inthecohortoftheFraminghamstudy[50]andinacohortofpatientswithGilbertsyndrome[51].Thissuggeststhatbili-rubinisaprotectivefactoragainstcardiovasculardiseaseindependentfromstandardcardiovascularriskfactors.Recentclinicalstudiesdemonstratethatslightlyelevatedbilirubinconcentrationsexertprotectiveeffectsagainstava-rietyofoxidativestress-induceddisease

---

### Chunk 26/30
**Article:** Measurement and clinical usefulness of bilirubin in liver disease (2021)
**Journal:** Advances in Laboratory Medicine
**Section:** other | **Similarity:** 0.594

henexposedtolightandstoredatroomtemperatureandcanoxidizetobiliverdin(negativediazo)atthenormallyacidpHofurine.Iftestingisdelayed,thesamplemustbekeptinadarkplaceandstoredat2–8°Cforamaximumof24h.Thismethodmaybeaffectedbypositive(ascorbicacid,nitrites)andnegative(substancesthatgiveurineabrownish/redcolorsuchasdrugsormetabolites,i.e.,rifampicin)interferences[29].Itisworthmentioningthatconcurrenturobilinogendeterminationwiththesamereactivestripprovidesguid-anceaboutthenatureofthebilirubinmetabolismdisorder.Elevatedurobilinogenlevelswithconcomitantincreasedornormalbilirubinaresuggestiveofincreasedhemolysisorliverdisease,withelevatedenterohepaticcirculation.Incontrast,elevatedbilirubinconcentrationswithnormalurobilinogenlevelsindicatereducedCBsecretionintotheintestine,asincasesofbileductobstruction.RequirementsforbilirubintestingTherequirementsfortestingbilirubinvaryslightlydepend-
ingonitsuse.AccordingtoCLIAguidelines,atotalerrorof20%orlessisrequired.Also,thequalitystandardsestab-lishedby

---

### Chunk 27/30
**Article:** Measurement and clinical usefulness of bilirubin in liver disease (2021)
**Journal:** Advances in Laboratory Medicine
**Section:** methods | **Similarity:** 0.594

.17.ColePG,LatheGH,BillingBH.Thediazoreactingpigmentsofserum,urineandbile.BiochemJ1953;55:xiii.18.KuenzleCC,MaierC,RüttnerJR.Thenatureoffourbilirubinfractionsfromserumandofthreebilirubinfractionsfrombile.J
LabClinMed1966;67:294–306.19.LoSF,DoumasBT,AshwoodER.PerformanceofbilirubindeterminationsinUSlaboratories-revisited.ClinChem2004;50:
190–4.20.JendrassikG.VereinfachtephotometrischeMethodenzurBestimmungdesBlutbilirubins.BiochemZ1938;297:81–9.21.DoumasBT,Kwok-CheungPP,PerryBW.Candidatereferencemethodfordeterminationoftotalbilirubininserum:developmentandvalidation.ClinChem1985;31:1779–89.22.PerryBW,DoumasBT,BayseDD,ButlerT,CohenA,FellowsW,etal.Acandidatereferencemethodfordeterminationofbilirubininserum.Testfortransferability.ClinChem1983;29:297–301.23.SchlebuschH,AxerK,SchneiderC,LiappisN,RöhleG.Comparisonofveroutinemethodswiththecandidatereferencemethodforthedeterminationofbilirubininneonatalserum.ClinChemLabMed1990;28:203–10.24.``OsawaS,SugoS,YoshidaT,YamaokaT,NomuraF.A

---

### Chunk 28/30
**Article:** Measurement and clinical usefulness of bilirubin in liver disease (2021)
**Journal:** Advances in Laboratory Medicine
**Section:** other | **Similarity:** 0.585

013;132:871–81.28.El-BeshbishiSN,ShattuckKE,MohammadAA,PetersenJR.Hyperbilirubinemiaandtranscutaneousbilirubinometry.Clin
Chem2009;55:1280–7.29.RojoVizcainoI,AntojaRiboF,GalimanySoleR.Interferenciasenelanalisisdeorinacontirasmultirreactivas:comisionde
interferenciasyefectosdelosmedicamentosenbioquimicaclinica.QuimClin2000;19:34–40.30.MoranchoJ,PradaE,Gutiérrez-BassiniG,SalasA,BlázquezR,JouJM,etal.Actualizacióndelasespecicacionesdelacalidadanalítica2014.ConsensodelasSociedadesCientícasnacionales.RevDelLabClínico2014;7:3–8.31.RicósC,PerichC,DoménechM,FernándezP,BioscaC,MinchinelaJ,etal.Variación;biológica.Revisióndesdeunaperspectivapráctica.RevDelLabClin2010;3:192–200.32.KlaukeR,KytziaHJ,WeberF,Grote-KoskaD,BrandK,SchumannG.Referencemeasurementprocedurefortotalbilirubininserumre-evaluatedandmeasurementuncertaintydetermined.ClinChimActa2018;481:115–20.33.LoSF,DoumasBT.ThestatusofbilirubinmeasurementsinU.S.Laboratories:whyisaccuracyelusive?SeminPerinatol2011;35:1

---

### Chunk 29/30
**Article:** Measurement and clinical usefulness of bilirubin in liver disease (2021)
**Journal:** Advances in Laboratory Medicine
**Section:** other | **Similarity:** 0.585

ors360
GuerraRuizetal.:Bilirubin:measurementandclinicalusefulness

ofacuteonchronicliverfailureandtheirpredictivevalueformortality.LiverInt2013;33:40–52.43.HelmkeS,ColmeneroJ,EversonGT.Noninvasiveassessmentofliverfunction.CurrOpinGastroenterol2015;31:199–208.44.EuropeanAssociationfortheStudyoftheLiver,ClinicalPracticeGuidelinesPanel,WendonJ,PanelMembers,CordobaJ,Dhawan
A,etal.EASLClinicalPracticalGuidelinesonthemanagementofacute(fulminant)liverfailure.JHepatol2017;66:1047–81.45.FengD,WangM,HuJ,LiS,ZhaoS,LiH,etal.Prognosticvalueofthealbumin–bilirubingradeinpatientswithhepatocellularcarcinomaandotherliverdiseases.AnnTranslMed2020;8:553.46.López-VelázquezJA,Chávez-TapiaNC,Ponciano-RodríguezG,Sánchez-ValleV,CaldwellSH,UribeM,etal.Bilirubinaloneasa
biomarkerforshort-termmortalityinacute-on-chronicliverfailure:animportantprognosticindicator.AnnHepatol2014;13:98–104.47.BreimerLH,WannametheeG,EbrahimS,ShaperAG.Serumbilirubinandriskofischemicheartdiseaseinmiddle-agedBritishmen.Cl

---

### Chunk 30/30
**Article:** Measurement and clinical usefulness of bilirubin in liver disease (2021)
**Journal:** Advances in Laboratory Medicine
**Section:** conclusion | **Similarity:** 0.584

iseaseindependentfromstandardcardiovascularriskfactors.Recentclinicalstudiesdemonstratethatslightlyelevatedbilirubinconcentrationsexertprotectiveeffectsagainstava-rietyofoxidativestress-induceddiseases,ofwhichathero-scleroticdiseasesarethemostclinicallyrelevant.Thisissuehasbeenthesubjectofanexcellentreview[52].ConclusionsBilirubinispartofthebasicstudyofliverfunction.Therearenumerousmeasurementplatformsandmethods,beingthediazomethodthegold-standard.Thesamplemostcommonlyusedisserumorplasma,andalsourine,forwhichoptimalpre-analyticalconditionsarerequired.Despiteitslimitedsensitivityandspecificity,bilirubinisfrequentlymeasuredfortheevaluationofdifferentpathologiesrelatedtoliverandbilefunction.Totalandconjugatedbilirubinconcentrationsprovideguidanceabouttheoriginofthealteration.Thesame
occurswithbilirubinandurobilinogendeterminationinserumandurine.Inthehospitalcontext,bilirubinconcen-trationsareveryusefulforprognosisofacuteliverdiseaseandmonitoringchronicliverdisease.Theseresultsmustbeinterp

---

