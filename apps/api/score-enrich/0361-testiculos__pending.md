# ScoreItem: Testículos

**ID:** `019bf31d-2ef0-7e35-94d7-0d232cc258ce`
**FullName:** Testículos (Histórico de doenças - Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual): - Genitália masculina)

**Preparation Metadata:**
- Quality Grade: **EXCELLENT**
- Total Chunks: 30 de 8 artigos
- Avg Similarity: 0.633

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-7e35-94d7-0d232cc258ce`.**

```json
{
  "score_item_id": "019bf31d-2ef0-7e35-94d7-0d232cc258ce",
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

**ScoreItem:** Testículos (Histórico de doenças - Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual): - Genitália masculina)

**30 chunks de 8 artigos (avg similarity: 0.633)**

### Chunk 1/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** discussion | **Similarity:** 0.696

F.;Nardone,V.;Guida,C.;Scialpi,M.;Cappabianca,S.DoesmultiparametricUSimprovediagnosticaccuracyinthecharacterizationofsmalltesticularmasses?Gland.Surg.2019,8,S136–S141.[CrossRef][PubMed]42.Schwarze,V.;Marschner,C.;Sabel,B.;deFigueiredo,G.N.;Marcon,J.;Ingrisch,M.;Knösel,T.;Rübenthaler,J.;Clevert,D.-A.Multiparametricultrasonographicanalysisoftesticulartumors:Asingle-centerexperienceinacollectiveof49patients.Scand.J.Urol.2020,54,241–247.[CrossRef][PubMed]43.Eiﬂer,J.B.,Jr.;King,P.;Schlegel,P.N.IncidentalTesticularLesionsFoundDuringInfertilityEvaluationareUsuallyBenignandMaybeManagedConservatively.J.Urol.2008,180,261–265.[CrossRef][PubMed]44.Scandura,G.;Verrill,C.;Protheroe,A.;Joseph,J.;Ansell,W.;Sahdev,A.;Shamash,J.;Berney,D.M.Incidentallydetectedtesticularlesions<10mmindiameter:Canorchidectomybeavoided?BJUInt.2017,121,575–582.[CrossRef][PubMed]45.Luzurier,A.;Maxwell,F.;Correas,J.;Benoit,G.;Izard,V.;Ferlicot,S.;Teglas,J.;Bellin,M.;Rocher,L.Qualitativeandquantitativecontrast-enha

---

### Chunk 2/30
**Article:** Testicular Cancer: Clinical Practice Guidelines for Diagnosis, Treatment and Follow-Up (2015)
**Journal:** European Urology
**Section:** abstract | **Similarity:** 0.693

Context: Testicular cancer is the most common solid malignancy affecting males aged 15-35 years. Early detection through clinical examination and ultrasound is crucial. Objective: To provide evidence-based guidelines for diagnosis and management of testicular cancer. Methods: Systematic review of literature and expert consensus. Results: Physical examination and testicular ultrasound are recommended for suspicious masses. Tumor markers (AFP, beta-hCG, LDH) should be obtained. Self-examination education improves early detection rates.

---

### Chunk 3/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** discussion | **Similarity:** 0.689

ology2017,285,640–649.[CrossRef][PubMed]19.Bertolotto,M.;Muça,M.;Currò,F.;Bucci,S.;Rocher,L.;Cova,M.A.MultiparametricUSforscrotaldiseases.Abdom.Imaging2018,43,899–917.[CrossRef]20.Pinto,S.P.;Huang,D.Y.;Dinesh,A.A.;Sidhu,P.S.;Ahmed,K.ASystematicReviewontheUseofQualitativeandQuantitativeContrast-enhancedUltrasoundinDiagnosingTesticularAbnormalities.Urology2021,154,16–23.[CrossRef]21.Ager,M.;Donegan,S.;Boeri,L.;deCastro,J.M.;Donaldson,J.F.;Omar,M.I.;Dimitropoulos,K.;Tharakan,T.;Janisch,F.;Muilwijk,T.;etal.Radiologicalfeaturescharacterisingindeterminatetestesmasses:Asystematicreviewandmeta-analysis.BJUInt.2022,131,288–300.[CrossRef][PubMed]22.Patel,K.;Sellars,M.E.;Clarke,J.L.;Sidhu,P.S.;Frcr,K.P.;Mbbs,F.M.E.S.;Msc,J.L.C.;Bsc,M.P.S.S.FeaturesofTesticularEpidermoidCystsonContrast-EnhancedSonographyandReal-timeTissueElastography.J.UltrasoundMed.2012,31,115–122.[CrossRef][PubMed]23.Lung,P.F.C.;Jaffer,O.S.;Sellars,M.E.;Sriprasad,S.;Kooiman,G.G.;Sidhu,P.S.Contrast-EnhancedUltrasoundi

---

### Chunk 4/30
**Article:** Scrotal Masses (2022)
**Journal:** American Family Physician
**Section:** abstract | **Similarity:** 0.679

Revisão clínica sobre massas escrotais, dividindo condições em dolorosas (torção testicular, torção de apêndice testicular, epididimite) e indolores (hidrocele, varicocele, câncer testicular). Sistema TWIST para estratificação de risco de torção (escore ≥5 indica necessidade de exploração cirúrgica urgente). Epididimite mais comum por Chlamydia trachomatis ou Neisseria gonorrhoeae em homens sexualmente ativos de 14-35 anos. Câncer testicular é tumor sólido mais comum em homens de 15-34 anos, apresentando-se como nódulo unilateral firme.

---

### Chunk 5/30
**Article:** A Call to Action to Review the USPSTF's Recommendation for Testicular Self-Examination (2022)
**Journal:** American Journal of Men's Health
**Section:** results | **Similarity:** 0.669

9-2Lin, K., & Sharangpani, R. (2010). Screening for testicular can-cer: An evidence review for the U.S. Preventive Services 
Task Force. Annals of Internal Medicine, 153, 396–399. https://doi.org/10.7326/0003-4819-153-6-201009210-00007; https://journals.sagepub.com/author-instructions/
JMH?_gl=1*u3rvu6*_ga*MTQ3OTE5National Cancer Institute. (2021). Surveillance, epidemiology and end results program—Cancer stat facts: Testicular 
cancer. https://seer.cancer.gov/statfacts/html/testis.htmlRovito, M. J., Adams, W. B., Craycraft, M., Gooljar, C., Maresca, M., Guelmes, J., & Gallelli, A. (2022). The asso-ciation between testicular self-examination and stages of testicular cancer diagnosis: A cross-sectional analysis. 
Journal of Adolescent and Young Adult Oncology, 11(1), 41–47.Rovito, M. J., Leone, J. E., & Cavayero, C. T. (2018). “Off-label” usage of testicular self-examination (TSE): Benefits 
beyond cancer detection. American Journal of Men’s 
Health, 12(3), 505–513.

---

### Chunk 6/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** other | **Similarity:** 0.668

hnique.Eur.J.Radiol.2021,145,110000.[CrossRef][PubMed]2.Dogra,V.S.;Gottlieb,R.H.;Oka,M.;Rubens,D.J.SonographyoftheScrotum.Radiology2003,227,18–36.[CrossRef][PubMed]3.Bhatt,S.;Jafri,S.Z.H.;Wasserman,N.;Dogra,V.S.Nonneoplasticintratesticularmasses.Diagn.Interv.Radiol.2009,17,52–63.[CrossRef][PubMed]4.Bhatt,S.;Dogra,V.S.RoleofUSinTesticularandScrotalTrauma.RadioGraphics2008,28,1617–1629.[CrossRef][PubMed]5.Woodward,P.J.;Sohaey,R.;O’donoghue,M.J.;Green,D.E.FromtheArchivesoftheAFIP.RadioGraphics2002,22,189–216.[CrossRef][PubMed]6.McDonald,M.W.;Reed,A.B.;Tran,P.T.;Evans,L.A.TesticularTumorUltrasoundCharacteristicsandAssociationwithHistopathology.Urol.Int.2011,89,196–202.[CrossRef][PubMed]

Cancers2024,16,2309
15of17
7.Mirochnik,B.;Bhargava,P.;Dighe,M.K.;Kanth,N.UltrasoundEvaluationofScrotalPathology.Radiol.Clin.NorthAm.2012,50,317–332.[CrossRef][PubMed]8.Altaffer,L.F.;Steele,S.M.ScrotalExplorationsNegativeforMalignancy.J.Urol.1980,124,617–619.[CrossRef][PubMed]9.Toren,P.J.;Ro

---

### Chunk 7/30
**Article:** Standards for scrotal ultrasonography (2016)
**Journal:** Journal of Ultrasonography
**Section:** abstract | **Similarity:** 0.665

Artigo estabelecendo padrões para ultrassonografia escrotal. Transdutores lineares de alta frequência (≥7 MHz) são essenciais, com capacidades de Doppler colorido e espectral. Avaliação modo-B precisa incluindo volumetria testicular. Na avaliação epididimária, quando aumentado, documentar tamanho da cabeça no plano sagital, corpo e cauda. Manobra de Valsalva é crítica no diagnóstico de varicocele, produzindo aumento do sinal Doppler e reversão do fluxo sanguíneo. Medição volumétrica de hidrocele recomendada pois pode associar-se a tumores embrionários.

---

### Chunk 8/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** results | **Similarity:** 0.655

ingwithsymptomssuchasscrotalpainorsubfertility.Includingbothpalpableandimpalpablelesionsallowsustoassesstheeffectivenessofultrasoundinarealisticclinicalsettingandreﬂecttheheterogeneityobservedinroutinepractice.However,thisbroadinclusioncomplicatestheextractionofpreciseguidanceforpopulationswithspeciﬁcclinicalpresentations.Theseriessizewasmodest,coveringawidespectrumoftesticularpathologiesbutwithlimitedcasesperspeciﬁccondition,suchasmalignantlesions<10mm.Thissuggestsaneedforfuturestudiesonalargerscaletovalidatetheseresultswithgreaterstatisticalpower.Additionally,thisstudyfo-cusedonaselectgroupofpatientswithfocaltesticularabnormalitiesidentiﬁedthroughgreyscaleultrasound,riskingselectionbias.Theoperator-dependentnatureofultrasoundalsopresentsachallengeinensuringconsistentandreproducibleresultsacrossdifferentexaminers.Lastly,non-surgicalcasesmanagedasbenignwerenotconﬁrmedhistologically,leavingroomfordiagnosticuncertainty.5.ConclusionsOurdecade-longexperienceindicatesthatintegrati

---

### Chunk 9/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** other | **Similarity:** 0.653

ScrotalPathology.Radiol.Clin.NorthAm.2012,50,317–332.[CrossRef][PubMed]8.Altaffer,L.F.;Steele,S.M.ScrotalExplorationsNegativeforMalignancy.J.Urol.1980,124,617–619.[CrossRef][PubMed]9.Toren,P.J.;Roberts,M.;Lecker,I.;Grober,E.D.;Jarvi,K.;Lo,K.C.SmallIncidentallyDiscoveredTesticularMassesinInfertileMen—IsActiveSurveillancetheNewStandardofCare?J.Urol.2010,183,1373–1377.[CrossRef]10.Carmignani,L.;Gadda,F.;Mancini,M.;Gazzano,G.;Nerva,F.;Rocco,F.;Colpi,G.M.DETECTIONOFTESTICULARULTRA-SONOGRAPHICLESIONSINSEVEREMALEINFERTILITY.J.Urol.2004,172,1045–1047.[CrossRef]11.Carmignani,L.;Gadda,F.;Gazzano,G.;Nerva,F.;Mancini,M.;Ferruti,M.;Bulfamante,G.;Bosari,S.;Coggi,G.;Rocco,F.;etal.HighIncidenceofBenignTesticularNeoplasmsDiagnosedbyUltrasound.J.Urol.2003,170,1783–1786.[CrossRef][PubMed]12.Zuniga,A.;Lawrentschuk,N.;Jewett,M.A.S.Organ-sparingapproachesfortesticularmasses.Nat.Rev.Urol.2010,7,454–464.[CrossRef][PubMed]13.Tsili,A.C.;Bertolotto,M.;Turgut,A.T.;Dogra,V.;Freeman,S.;Rocher,L.;Bel  

---

### Chunk 10/30
**Article:** The Male Genital Examination: A Critical Component of Primary Care (2018)
**Journal:** American Family Physician
**Section:** abstract | **Similarity:** 0.641

Common male genital conditions include erectile dysfunction, varicocele, hydrocele, testicular masses, phimosis, and sexually transmitted infections. A thorough genital examination should be part of routine health maintenance for men. Key components include inspection for lesions, palpation of testes for masses, assessment for varicocele and hydrocele, and evaluation of urethral discharge. Early detection of testicular cancer through examination and patient education about self-examination can be life-saving. Erectile dysfunction warrants cardiovascular risk assessment. Clinical decision-making should guide when to order scrotal ultrasound, hormonal studies, or refer to urology.

---

### Chunk 11/30
**Article:** A Call to Action to Review the USPSTF's Recommendation for Testicular Self-Examination (2022)
**Journal:** American Journal of Men's Health
**Section:** discussion | **Similarity:** 0.639

ty. (2022). Cancer facts and figures 2022. 4. https://www.cancer.org/content/dam/cancer-org/research/cancer-facts-and-statistics/annual-cancer-facts-and-figures/2022/2022-cancer-facts-and-figures.pdfAmerican Urological Association. (2014). AUA Men’s Health Checklist. https://www.google.com/url?sa=t&rct=j&q=&esrc=s&source=web&cd=&ved=2ahUKEwimrKqY65r3AhV8l2oFHZarAtIQFnoECAkQAQ&ur
l=https%3A%2F%2Fwww.auanet.org%2Fdocuments%2Feducation%2Fclinical-guidance%2FMens-Health-
Checklist.pdf&usg=AOvVaw37zmiamYQgrxnHFMZcr-y8&cshid=1650189706876783Barbonetti, A., Martorella, A., Minaldi, E., D’Andrea, S., Bardhi, D., Castellini, C., Francavilla, F., & Francavilla, S. (2019). Testicular cancer in infertile men with and without testicular microlithiasis: A systematic review and meta-analysis of case-control studies. Frontiers in Endocrinology, 10, Article 164.Calonge, N. (2005). Screening for testicular cancer: Recommendation statement. U.S. Preventive Services  Task Force.

---

### Chunk 12/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** discussion | **Similarity:** 0.638

.[59]reportedtheirexperienceintestis-sparingsurgeryforincidentallydetectedtesticularlesions<5mm.Allpatientswhooptedfortestis-sparingsurgeryinsteadofradicalorchiectomyfollowingMPUSassessmentswereconﬁrmedtohavebenignconditionsduringourdecade-longclinicalexperience.Ourexperienceshowedtheutilityofadvancedultrasoundtechniques

Cancers2024,16,2309
14of17
inpreoperativecharacterisationcouldaidinfacilitatingtheformulationofoptimalclinicalmanagementstrategies.Thisstudyissubjecttosomelimitations.Primarily,itwasstructuredasaretrospectiveanalysis,whichmayimpacttheprospectiveapplicabilityoftheﬁndings.Ourstudyencom-passesawiderangeofcasesencounteredinroutineclinicalpracticeatourtertiarycarecentreoveraten-yearperiod.Ourdatasetincludesclinicallyimpalpablelesionsinciden-tallydiscoveredinadultmenpresentingwithsymptomssuchasscrotalpainorsubfertility.Includingbothpalpableandimpalpablelesionsallowsustoassesstheeffectivenessofultrasoundinarealisticclinicalsettingandreﬂecttheheterogeneityobservedinrout

---

### Chunk 13/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** other | **Similarity:** 0.638

resEnUrol.2015,25,75–82.[CrossRef]52.Ma,W.;Sarasohn,D.;Zheng,J.;Vargas,H.A.;Bach,A.CausesofAvascularHypoechoicTesticularLesionsDetectedatScrotalUltrasound:CanTheyBeConsideredBenign?Am.J.Roentgenol.2017,209,110–115.[CrossRef][PubMed]53.Patrikidou,A.;Cazzaniga,W.;Berney,D.;Boormans,J.;deAngst,I.;DiNardo,D.;Fankhauser,C.;Fischer,S.;Gravina,C.;Gremmels,H.;etal.EuropeanAssociationofUrologyGuidelinesonTesticularCancer:2023Update.Eur.Urol.2023,84,289–301.[CrossRef][PubMed]54.Maxwell,F.;Savignac,A.;Bekdache,O.;Calvez,S.;Lebacle,C.;Arama,E.;Garrouche,N.;Rocher,L.LeydigCellTumorsoftheTestis:AnUpdateoftheImagingCharacteristicsofaNotSoRareLesion.Cancers2022,14,3652.[CrossRef]

Cancers2024,16,2309
17of17
55.Nicolai,N.;Necchi,A.;Raggi,D.;Biasoni,D.;Catanzaro,M.;Piva,L.;Stagni,S.;Maffezzini,M.;Torelli,T.;Faré,E.;etal.ClinicalOutcomeinTesticularSexCordStromalTumors:TestisSparingvsRadicalOrchiectomyandManagementofAdvancedDisease.Urology2015,85,402–406.[CrossRef][PubMed]56.Westlander,G.;Ekerhov

---

### Chunk 14/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** discussion | **Similarity:** 0.637

nignandthemalignantgroupsinourcohortofpatients.InastudybyEiﬂeretal.[43]basedon49lesionsin145menreferredforazoospermiawhounderwentultrasonographicanalysis,theinvestigatorsproposedanalgorithmbasedontumourmarkersandthesizeandvascularityofthelesions.Theysuggestedthatalesion<5mm,characterisedbyanabsenceofvascularityandnegativetumourmarkers,couldbefollowedbyserialUSmonitoring.AfurtherstudybyScanduraetal.[44]reportedthemajorityoftesticularlesions<10mmidentiﬁedbyradiologywerebenign.Inourstudy,itwasshownthatforanenhancinglesion,havingasizelargerthan10mmincreasestheoddsofmalignancybynearlytentimes.However,wefoundthatgreyscaleultrasounddidnotdemonstratestatisticallysigniﬁcantdifferencesinfeaturessuchasmargin,hypoechoicnature,orthepresenceofmicrolithiasisbetweenbenignandmalignanttesticularlesions.Thelackofsigniﬁcantdifferencesintheseobservedgreyscalesonographicfeaturessuggeststhattheseparametersaloneareinsufﬁcientforreliabledifferentiation,underscoringthelimitationofgreyscaleUS.Thedescri

---

### Chunk 15/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** introduction | **Similarity:** 0.636

fadvancedultrasoundtechniquesinenhancingtheevaluationoffocaltesticularabnormalitiesinclinicalpracticeandcouldaidashifttowardstestis-sparingmanagementstrategies.Keywords:multiparametric;ultrasound;testicularcancer;testis-sparingsurgery;orchiectomy
1.IntroductionConventionalultrasonography(US),includinggreyscaleimagingandcolourDopplerUS(CDUS),standsasthecornerstoneforevaluatingscrotalpathologiesduetoitshighresolution,availability,cost-effectiveness,andabsenceofionizingradiation[1–4].Despite
Cancers2024,16,2309.https://doi.org/10.3390/cancers16132309https://www.mdpi.com/journal/cancers

Cancers2024,16,2309
2of17
itswidespreaduse,thespeciﬁcityofgreyscaleultrasoundincharacterisingscrotalmassesremainslimited,oftenleavingthenatureofsuchlesionsambiguous[5–7].Traditionally,solidtesticularlesions,especiallythosepresentingaspalpablelumps,haveledtoradi-calorchidectomy[8,9].However,thelandscapeofscrotalultrasonographyhasevolvedsigniﬁcantlywithadvancementsintechnologyandtechnique,includinghi

---

### Chunk 16/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** discussion | **Similarity:** 0.625

dinourdepartmentforthefollowingindications:evaluationandlocationofpalpablescrotalmasses,detectionofprimarytumours,follow-upofpa-tientswithtesticularmicrolithiasis,follow-upofpatientswithpreviouslymphoma,acutescrotum,scrotaltrauma,localisationoftheundescendedtestis,detectionofvaricocelesininfertilemen,andevaluationoftesticularischaemia.Fromaninitialdatasetofallscrotalultrasoundexaminations,124consecutivecasesoffocaltesticularabnormalitiesinvestigatedbyMPUSwereselectedforanalysis.InstitutionalReviewBoardapprovalwasobtained,withallproceduresperformedinaccordancewithethicalstandardsandpatientconﬁdentialityguidelines.2.2.PatientCohortandDataAcquisitionEligiblecaseswereidentiﬁedfromthedepartmentalultrasounddatabasebasedontheinclusioncriteriaofhavingundergoneMPUS,comprisinggreyscaleUS,CDUS,CEUS,andSE.Comprehensiveclinicaldata,includingthepatients’ages,clinicalpresentations,tumourmarkers,histopathologicalreports,andfollow-upoutcomes,wereextractedfromelectronicpatientrecords.Imagingdatawe

---

### Chunk 17/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** results | **Similarity:** 0.617

  t.TheWilcoxonsigned-ranktestwasusedforcomparingperfusionparametersbetweeneachlesionanditsadjacentnormaltesticulartissue.AllstatisticalanalyseswereperformedusingSPSSsoftware(version29,IBM),withasigniﬁcancethresholdsetatp<0.05.3.ResultsAtotalof124MPUSexaminationsconductedtoevaluateintratesticularfocalabnormalitieswereincluded.Thiscohortcomprised78benignand46malignantdiagnoses.Thedemographicanalysis(Table1)demonstratednosignificantagedifferencebetweenpatientswithbenignandthosewithmalignanttesticularabnormalities(p=0.27).Asignificantdifference(p=0.005)wasobservedinthepalpabilityofthefocalabnormalities,withagreaterprevalenceinmalignantcasescomparedtobenigncases.Elevatedtumourmarkerswerenotedonlyin9outof46patientswithmalignantlesionsandnoneinpatientswithabenigndiagnosis.Histopathologicalanalysiswasavailablefor76cases(Table2),identifying46asmalignantand30asbenign.Amongthese,16patientsunderwenttestis-sparingsurgery,withallabnormalitiesinthissubsetconﬁrmedasbenign.Fortheremaining48casesno

---

### Chunk 18/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** other | **Similarity:** 0.613

ntstothediagnosis.Br.J.Radiol.2012,85,S41–S53.[CrossRef][PubMed]33.Huang,D.Y.;Pesapane,F.;Rafailidis,V.;Deganello,A.;Sellars,M.E.;Sidhu,P.S.Theroleofmultiparametricultrasoundinthediagnosisofpaediatricscrotalpathology.Br.J.Radiol.2020,93,20200063.[CrossRef][PubMed]34.Mansoor,N.M.;Huang,D.Y.;Sidhu,P.S.Multiparametricultrasoundimagingcharacteristicsofmultipletesticularadrenalresttumoursincongenitaladrenalhyperplasia.Ultrasound2022,30,80–84.[CrossRef][PubMed]35.Sidhu,P.S.MultiparametricUltrasound(MPUS)Imaging:TerminologyDescribingtheManyAspectsofUltrasonography.UltraschallMed.2015,36,315–317.[CrossRef]36.Itoh,A.;Ueno,E.;Tohno,E.;Kamma,H.;Takahashi,H.;Shiina,T.;Yamakawa,M.;Matsumura,T.BreastDisease:ClinicalApplicationofUSElastographyforDiagnosis.Radiology2006,239,341–350.[CrossRef]37.Kim,I.;Young,R.H.;Scully,R.E.Leydigcelltumorsofthetestis.Am.J.Surg.Pathol.1985,9,177–192.[CrossRef]38.Farkas,L.M.;Székely,J.G.;Pusztai,C.;Baki,M.HighFrequencyofMetastaticLeydigCellTesticularTumours.O

---

### Chunk 19/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** discussion | **Similarity:** 0.613

hortalsorevealedthatlatehyperenhancementonCEUSinbenignenhancingtumours,suchasinLeydigcelltumourswithlowmalignantpotential,offersapotentialfeaturefordistinguishingthesefrommalignanttumours,suchasseminomas.ThisstudyalsoshowedthatbyintegratingCEUSandthecombinationofCEUSandSE,thediagnosticperformanceofultrasoundimagingindifferentiatingbenignandmalignantfocaltesticularabnormalitieswasimproved.Speciﬁcally,CEUSshowedahighersensitivitywhencomparedtoCDUS,andCEUS+SEshowedhigherspeciﬁcitywhencomparedtoCDUS.Notably,duringourdecade-longclinicalexperience,allpa-tientswhooptedfortestis-sparingsurgeryinsteadofradicalorchiectomyfollowingMPUSassessmentswereconﬁrmedtohavebenignconditions.Inourstudy,astatisticallysigniﬁcantdifferencewasobservedintheassociationoflesionsizeandﬁnaldiagnosisbetweenthebenignandthemalignantgroupsinourcohortofpatients.InastudybyEiﬂeretal.[43]basedon49lesionsin145menreferredforazoospermiawhounderwentultrasonographicanalysis,theinvestigatorsproposedanalgorithmbasedontu

---

### Chunk 20/30
**Article:** Testicular Torsion (2023)
**Journal:** StatPearls
**Section:** abstract | **Similarity:** 0.612

Revisão sobre torção testicular, emergência urológica tempo-dependente. A viabilidade testicular diminui significativamente após 6 horas do início dos sintomas. Achados físicos incluem testículo em posição transversa ou elevada. Ultrassonografia com Doppler apresenta sensibilidade de 93% e especificidade de 100%. Taxa de salvamento aproxima-se de 100% quando a intervenção cirúrgica ocorre dentro de 6 horas, mas cai para menos de 50% após 12-24 horas. Mais comum em pacientes abaixo de 25 anos.

---

### Chunk 21/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** other | **Similarity:** 0.612

,M.;Bob,F.;Bojunga,J.;Brock,M.;etal.TheEFSUMBGuidelinesandRecommendationsfortheClinicalPracticeofElastographyinNon-HepaticApplications:Update2018.UltraschallMed.2019,40,425–453.[CrossRef][PubMed]16.Isidori,A.M.;Pozza,C.;Gianfrilli,D.;Giannetta,E.;Lemma,A.;Poﬁ,R.;Barbagallo,F.;Manganaro,L.;Martino,G.;Lombardo,F.;etal.DifferentialDiagnosisofNonpalpableTesticularLesions:QualitativeandQuantitativeContrast-enhancedUSofBenignandMalignantTesticularTumors.Radiology2014,273,606–618.[CrossRef][PubMed]17.Fang,C.;Huang,D.Y.;Sidhu,P.S.Elastographyoffocaltesticularlesions:Currentconceptsandutility.Ultrasonography2019,38,302–310.[CrossRef][PubMed]18.Auer,T.;DeZordo,T.;Dejaco,C.;Gruber,L.;Pichler,R.;Jaschke,W.;Dogra,V.S.;Aigner,F.ValueofMultiparametricUSintheAssessmentofIntratesticularLesions.Radiology2017,285,640–649.[CrossRef][PubMed]19.Bertolotto,M.;Muça,M.;Currò,F.;Bucci,S.;Rocher,L.;Cova,M.A.MultiparametricUSforscrotaldiseases.Abdom.Imaging2018,43,899–917.[CrossRef]20.Pinto,S.P.;Hua

---

### Chunk 22/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** other | **Similarity:** 0.610

ons,especiallythosepresentingaspalpablelumps,haveledtoradi-calorchidectomy[8,9].However,thelandscapeofscrotalultrasonographyhasevolvedsigniﬁcantlywithadvancementsintechnologyandtechnique,includinghighfrequency,tissueharmonic,andcompoundimaging.Thisevolution,alongsideabroaderspectrumofclinicalapplications,hasincreasedthedetectionofsmall,incidentalfocaltesticularlesions,manyofwhicharebenign.Indeed,recentliteraturesuggestsapredominanceofbenignityinthesecases,withLeydigcelltumourswithlowmalignantpotential(LCT-LMP)consti-tutingasigniﬁcantfractionamongsmall,impalpable,incidentallydiscoveredtesticularnodules[10,11].Thisshiftinthediagnosticlandscapenecessitatesareconsiderationofradicalorchiec-tomyforfocaltesticularabnormalities,pivotingtowardsmoreorgan-sparingapproacheswhenthereisahighlikelihoodofbenignity[12].Yet,despiteimprovedimagingmodalitiesanddiagnosticaids,includingtumourmarkersandsecond-lineMRIasrecommendedbytheEuropeanSocietyofUrogenitalRadiology(ESUR)[13],signiﬁcantdiagnosticam

---

### Chunk 23/30
**Article:** A Call to Action to Review the USPSTF's Recommendation for Testicular Self-Examination (2022)
**Journal:** American Journal of Men's Health
**Section:** other | **Similarity:** 0.609

19). Testicular aware-ness: The what, the why, and the how. International Journal of Mens Social and Community Health, 2(1), e1–e10.Saab, M. M., Landers, M., & Hegarty, J. (2018). Males’ aware-ness of benign testicular disorders: An integrative review . American Journal of Men’s Health, 12(3), 556–566.Schepisi, G., De Padova, S., De Lisi, D., Casadei, C., Meggiolaro, E., Ruffilli, F., Rosti, G., Lolli, C., Ravaglia, 
G., Conteduca, V., Farolfi, A., Grassi, L., & De Giorgi, U. (2019). Psychosocial issues in long-term survivors of tes-
ticular cancer. Frontiers in Endocrinology, 10, Article 113. https://doi.org/10.3389/fendo.2019.00113Sevilla, M. B., & Gonzales, D. G. (2022). Association between testicular cancer and microlithiasis. Actas Urológicas 
Españolas (English Edition). Epub ahead of print 29 July. 
DOI: 10.1016/j.acuroe.2022.07.002.U.S Preventive Services Task Force. (1996). January 01). Testicular cancer: Screening.

---

### Chunk 24/30
**Article:** A Call to Action to Review the USPSTF's Recommendation for Testicular Self-Examination (2022)
**Journal:** American Journal of Men's Health
**Section:** other | **Similarity:** 0.609

control studies. Frontiers in Endocrinology, 10, Article 164.Calonge, N. (2005). Screening for testicular cancer: Recommendation statement. U.S. Preventive Services  Task Force. American Family Physician, 72(10), 2069–2070.Fadich, A., Giorgianni, S. J., Rovito, M. J., Pecchia, G. A., Bonhomme, J. J., Adams, W. B., Stephenson, C. L., Mesa-Morales, F. E., & Sparkes, J. S. (2018). USPSTF testicular examination nomination-self-exam-
inations and examinations in a clinical setting. American 
Journal of Men's Health, 12(5), 1510–1516. https://doi.org/10.1177/1557988318768597Ghazarian, A. A., & McGlynn, K. A. (2020). Increasing inci-dence of testicular germ cell tumors among racial/ethnic 
minorities in the United States. Cancer Epidemiology and 
Prevention Biomarkers, 29(6), 1237–1245. https://doi.org/10.1158/1055-9965.EPI-20-0107Institute of Medicine. (2011). Clinical practice guidelines we can trust. National Academies Press. https://doi.
org/10.17226/13058Kim, C., McGlynn, K.

---

### Chunk 25/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** other | **Similarity:** 0.604

IntratesticularLesionsWithStrainElastographyUsingStrainRatioandColorMapVisualGrading:DifferentiationofNeoplasticandNonneoplasticLesions.J.UltrasoundMed.2018,38,223–232.[CrossRef][PubMed]27.Yusuf,G.;Sellars,M.E.;Kooiman,G.G.;Diaz-Cano,S.;Sidhu,P.S.GlobalTesticularInfarctioninthePresenceofEpididymitis.J.UltrasoundMed.2013,32,175–180.[CrossRef][PubMed]28.Rafailidis,V.;Robbie,H.;Konstantatou,E.;Huang,D.Y.;Deganello,A.;ESellars,M.;Cantisani,V.;Isidori,A.M.;Sidhu,P.S.Sonographicimagingofextra-testicularfocallesions:Comparisonofgrey-scale,colourDopplerandcontrast-enhancedultrasound.Ultrasound2016,24,23–33.[CrossRef][PubMed]29.Patel,K.V.;Huang,D.Y.;Sidhu,P.S.Metachronousbilateralsegmentaltesticularinfarction:Multi-parametricultrasoundimagingwithgrey-scaleultrasound,Dopplerultrasound,contrast-enhancedultrasound(CEUS)andreal-timetissueelastography(RTE).J.Ultrasound2014,17,233–238.[CrossRef]30.Zebari,S.;Huang,D.Y.;Wilkins,C.J.;Sidhu,P.S.AcuteTesticularSegmentalInfarctFollowingEndovascular

---

### Chunk 26/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** other | **Similarity:** 0.602

chidectomybeavoided?BJUInt.2017,121,575–582.[CrossRef][PubMed]45.Luzurier,A.;Maxwell,F.;Correas,J.;Benoit,G.;Izard,V.;Ferlicot,S.;Teglas,J.;Bellin,M.;Rocher,L.Qualitativeandquantitativecontrast-enhancedultrasonographyforthecharacterisationofnon-palpabletesticulartumours.Clin.Radiol.2018,73,322.e1–322.e9.[CrossRef][PubMed]46.Schröder,C.;Lock,G.;Schmidt,C.;Löning,T.;Dieckmann,K.-P.Real-TimeElastographyandContrast-EnhancedUltrasonographyintheEvaluationofTesticularMasses:AComparativeProspectiveStudy.UltrasoundMed.Biol.2016,42,1807–1815.[CrossRef][PubMed]47.Aigner,F.;DeZordo,T.;Pallwein-Prettner,L.;Junker,D.;Schäfer,G.;Pichler,R.;Leonhartsberger,N.;Pinggera,G.;Dogra,V.S.;Frauscher,F.Real-timeSonoelastographyfortheEvaluationofTesticularLesions.Radiology2012,263,584–589.[CrossRef][PubMed]48.Goddi,A.;Sacchi,A.;Magistretti,G.;Almolla,J.;Salvadore,M.Real-timetissueelastographyfortesticularlesionassessment.Eur.Radiol.2011,22,721–730.[CrossRef][PubMed]49.Pozza,C.;Gianfrilli,D.;Fattori

---

### Chunk 27/30
**Article:** Epididymitis (2023)
**Journal:** StatPearls
**Section:** abstract | **Similarity:** 0.599

Revisão abrangente sobre epididimite, a causa mais comum de dor escrotal aguda em adultos. Aborda epidemiologia, fisiopatologia, diagnóstico diferencial, achados do exame físico (sensibilidade à palpação do epidídimo ao longo do aspecto posterior e superior do testículo), ultrassonografia com avaliação de fluxo vascular, e manejo terapêutico. Enfatiza a importância de excluir torção testicular como emergência cirúrgica. Incidência anual superior a 600.000 casos nos EUA, com pico entre 20-39 anos.

---

### Chunk 28/30
**Article:** A Call to Action to Review the USPSTF's Recommendation for Testicular Self-Examination (2022)
**Journal:** American Journal of Men's Health
**Section:** other | **Similarity:** 0.596

https://doi.org/10.1177/15579883221130186
American Journal of Men’s HealthSeptember-October 1 –6© The Author(s) 2022Article reuse guidelines: sagepub.com/journals-permissionsDOI: 10.1177/15579883221130186journals.sagepub.com/home/jmh
Creative Commons Non Commercial CC BY-NC: This article is distributed under the terms of the Creative Commons Attribution-NonCommercial 4.0 License (https://creativecommons.org/licenses/by-nc/4.0/) which permits non-commercial use, reproduction and distribution of the work without further permission provided the original work is attributed as specified on the SAGE and Open Access pages (https://us.sagepub.com/en-us/nam/open-access-at-sage).Urological Cancer - Opinion PieceOverview of the IssueTesticular cancer (TC) is the most common malignancy in males ages 20 to 34 years, with a median age of diag-
nosis of 33 ( National Cancer Institute [NCI], 2021). The 
American Cancer Society (ACS, 2022) predicts 9,910 
new cases and 460 deaths from TC in 2022.

---

### Chunk 29/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** discussion | **Similarity:** 0.592

s.Thelackofsigniﬁcantdifferencesintheseobservedgreyscalesonographicfeaturessuggeststhattheseparametersaloneareinsufﬁcientforreliabledifferentiation,underscoringthelimitationofgreyscaleUS.Thedescriptivestatisticalanalysisofvariousultrasoundmodalities,includingCDUS,CEUS,andSE,demonstratedhighsensitivityacrossthesetechniques,aligningwithﬁndingsreportedinexistingliterature[16,18,45,46].However,allexaminedmodalitiesexhibitedlowspeciﬁcityinourstudy,withSEidentiﬁedastheleastspeciﬁctechnique.ExistingresearchonthediagnosticaccuracyofSEforevaluatingtesticularlesionspresentsdivergentoutcomes[46–49].Forinstance,Grassoetal.[50]comparedB-modepluscolourDopplerultrasoundtoreal-timeelastography(RTE)in41patientsandnotedtheinabilitytodifferen-tiatemalignantfrombenignlesionsbasedsolelyonelastography.Goddietal.[48]reportedanSEsensitivityof87.5%withaspeciﬁcityof98.2%inalargeseriesof144testicularlesionsbutcomprisingaclearmajorityofbenignlesions(112of144,77%).Aigneretal.[47]re-portedsimilarsen

---

### Chunk 30/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** other | **Similarity:** 0.592

cedSonographyandReal-timeTissueElastography.J.UltrasoundMed.2012,31,115–122.[CrossRef][PubMed]23.Lung,P.F.C.;Jaffer,O.S.;Sellars,M.E.;Sriprasad,S.;Kooiman,G.G.;Sidhu,P.S.Contrast-EnhancedUltrasoundintheEvaluationofFocalTesticularComplicationsSecondarytoEpididymitis.Am.J.Roentgenol.2012,199,W345–W354.[CrossRef][PubMed]24.Hedayati,V.;ESellars,M.;Sharma,D.M.;Sidhu,P.S.Contrast-enhancedultrasoundintesticulartrauma:Roleindirectingexploration,debridementandorgansalvage.Br.J.Radiol.2012,85,e65–e68.[CrossRef][PubMed]25.Yusuf,G.;Konstantatou,E.;Sellars,M.E.;Huang,D.Y.;Sidhu,P.S.MultiparametricSonographyofTesticularHematomas.J.UltrasoundMed.2015,34,1319–1328.[CrossRef][PubMed]26.Konstantatou,E.;Fang,C.;Romanos,O.;Derchi,L.E.;Bertolotto,M.;Valentino,M.;Kalogeropoulou,C.;Sidhu,P.S.EvaluationofIntratesticularLesionsWithStrainElastographyUsingStrainRatioandColorMapVisualGrading:DifferentiationofNeoplasticandNonneoplasticLesions.J.UltrasoundMed.2018,38,223–232.[CrossRef][PubMed]27.Yusuf,G.;

---

