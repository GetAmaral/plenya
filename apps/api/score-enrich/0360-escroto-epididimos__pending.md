# ScoreItem: Escroto / epidídimos

**ID:** `019bf31d-2ef0-7b42-9d37-487e91411a18`
**FullName:** Escroto / epidídimos (Histórico de doenças - Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual): - Genitália masculina)

**Preparation Metadata:**
- Quality Grade: **GOOD**
- Total Chunks: 30 de 12 artigos
- Avg Similarity: 0.568

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-7b42-9d37-487e91411a18`.**

```json
{
  "score_item_id": "019bf31d-2ef0-7b42-9d37-487e91411a18",
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

**ScoreItem:** Escroto / epidídimos (Histórico de doenças - Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual): - Genitália masculina)

**30 chunks de 12 artigos (avg similarity: 0.568)**

### Chunk 1/30
**Article:** Scrotal Masses (2022)
**Journal:** American Family Physician
**Section:** abstract | **Similarity:** 0.774

Revisão clínica sobre massas escrotais, dividindo condições em dolorosas (torção testicular, torção de apêndice testicular, epididimite) e indolores (hidrocele, varicocele, câncer testicular). Sistema TWIST para estratificação de risco de torção (escore ≥5 indica necessidade de exploração cirúrgica urgente). Epididimite mais comum por Chlamydia trachomatis ou Neisseria gonorrhoeae em homens sexualmente ativos de 14-35 anos. Câncer testicular é tumor sólido mais comum em homens de 15-34 anos, apresentando-se como nódulo unilateral firme.

---

### Chunk 2/30
**Article:** Epididymitis (2023)
**Journal:** StatPearls
**Section:** abstract | **Similarity:** 0.733

Revisão abrangente sobre epididimite, a causa mais comum de dor escrotal aguda em adultos. Aborda epidemiologia, fisiopatologia, diagnóstico diferencial, achados do exame físico (sensibilidade à palpação do epidídimo ao longo do aspecto posterior e superior do testículo), ultrassonografia com avaliação de fluxo vascular, e manejo terapêutico. Enfatiza a importância de excluir torção testicular como emergência cirúrgica. Incidência anual superior a 600.000 casos nos EUA, com pico entre 20-39 anos.

---

### Chunk 3/30
**Article:** Standards for scrotal ultrasonography (2016)
**Journal:** Journal of Ultrasonography
**Section:** abstract | **Similarity:** 0.660

Artigo estabelecendo padrões para ultrassonografia escrotal. Transdutores lineares de alta frequência (≥7 MHz) são essenciais, com capacidades de Doppler colorido e espectral. Avaliação modo-B precisa incluindo volumetria testicular. Na avaliação epididimária, quando aumentado, documentar tamanho da cabeça no plano sagital, corpo e cauda. Manobra de Valsalva é crítica no diagnóstico de varicocele, produzindo aumento do sinal Doppler e reversão do fluxo sanguíneo. Medição volumétrica de hidrocele recomendada pois pode associar-se a tumores embrionários.

---

### Chunk 4/30
**Article:** Testicular Torsion (2023)
**Journal:** StatPearls
**Section:** abstract | **Similarity:** 0.655

Revisão sobre torção testicular, emergência urológica tempo-dependente. A viabilidade testicular diminui significativamente após 6 horas do início dos sintomas. Achados físicos incluem testículo em posição transversa ou elevada. Ultrassonografia com Doppler apresenta sensibilidade de 93% e especificidade de 100%. Taxa de salvamento aproxima-se de 100% quando a intervenção cirúrgica ocorre dentro de 6 horas, mas cai para menos de 50% após 12-24 horas. Mais comum em pacientes abaixo de 25 anos.

---

### Chunk 5/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** discussion | **Similarity:** 0.632

ology2017,285,640–649.[CrossRef][PubMed]19.Bertolotto,M.;Muça,M.;Currò,F.;Bucci,S.;Rocher,L.;Cova,M.A.MultiparametricUSforscrotaldiseases.Abdom.Imaging2018,43,899–917.[CrossRef]20.Pinto,S.P.;Huang,D.Y.;Dinesh,A.A.;Sidhu,P.S.;Ahmed,K.ASystematicReviewontheUseofQualitativeandQuantitativeContrast-enhancedUltrasoundinDiagnosingTesticularAbnormalities.Urology2021,154,16–23.[CrossRef]21.Ager,M.;Donegan,S.;Boeri,L.;deCastro,J.M.;Donaldson,J.F.;Omar,M.I.;Dimitropoulos,K.;Tharakan,T.;Janisch,F.;Muilwijk,T.;etal.Radiologicalfeaturescharacterisingindeterminatetestesmasses:Asystematicreviewandmeta-analysis.BJUInt.2022,131,288–300.[CrossRef][PubMed]22.Patel,K.;Sellars,M.E.;Clarke,J.L.;Sidhu,P.S.;Frcr,K.P.;Mbbs,F.M.E.S.;Msc,J.L.C.;Bsc,M.P.S.S.FeaturesofTesticularEpidermoidCystsonContrast-EnhancedSonographyandReal-timeTissueElastography.J.UltrasoundMed.2012,31,115–122.[CrossRef][PubMed]23.Lung,P.F.C.;Jaffer,O.S.;Sellars,M.E.;Sriprasad,S.;Kooiman,G.G.;Sidhu,P.S.Contrast-EnhancedUltrasoundi

---

### Chunk 6/30
**Article:** The Male Genital Examination: A Critical Component of Primary Care (2018)
**Journal:** American Family Physician
**Section:** abstract | **Similarity:** 0.629

Common male genital conditions include erectile dysfunction, varicocele, hydrocele, testicular masses, phimosis, and sexually transmitted infections. A thorough genital examination should be part of routine health maintenance for men. Key components include inspection for lesions, palpation of testes for masses, assessment for varicocele and hydrocele, and evaluation of urethral discharge. Early detection of testicular cancer through examination and patient education about self-examination can be life-saving. Erectile dysfunction warrants cardiovascular risk assessment. Clinical decision-making should guide when to order scrotal ultrasound, hormonal studies, or refer to urology.

---

### Chunk 7/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** other | **Similarity:** 0.596

hnique.Eur.J.Radiol.2021,145,110000.[CrossRef][PubMed]2.Dogra,V.S.;Gottlieb,R.H.;Oka,M.;Rubens,D.J.SonographyoftheScrotum.Radiology2003,227,18–36.[CrossRef][PubMed]3.Bhatt,S.;Jafri,S.Z.H.;Wasserman,N.;Dogra,V.S.Nonneoplasticintratesticularmasses.Diagn.Interv.Radiol.2009,17,52–63.[CrossRef][PubMed]4.Bhatt,S.;Dogra,V.S.RoleofUSinTesticularandScrotalTrauma.RadioGraphics2008,28,1617–1629.[CrossRef][PubMed]5.Woodward,P.J.;Sohaey,R.;O’donoghue,M.J.;Green,D.E.FromtheArchivesoftheAFIP.RadioGraphics2002,22,189–216.[CrossRef][PubMed]6.McDonald,M.W.;Reed,A.B.;Tran,P.T.;Evans,L.A.TesticularTumorUltrasoundCharacteristicsandAssociationwithHistopathology.Urol.Int.2011,89,196–202.[CrossRef][PubMed]

Cancers2024,16,2309
15of17
7.Mirochnik,B.;Bhargava,P.;Dighe,M.K.;Kanth,N.UltrasoundEvaluationofScrotalPathology.Radiol.Clin.NorthAm.2012,50,317–332.[CrossRef][PubMed]8.Altaffer,L.F.;Steele,S.M.ScrotalExplorationsNegativeforMalignancy.J.Urol.1980,124,617–619.[CrossRef][PubMed]9.Toren,P.J.;Ro

---

### Chunk 8/30
**Article:** A Call to Action to Review the USPSTF's Recommendation for Testicular Self-Examination (2022)
**Journal:** American Journal of Men's Health
**Section:** results | **Similarity:** 0.582

9-2Lin, K., & Sharangpani, R. (2010). Screening for testicular can-cer: An evidence review for the U.S. Preventive Services 
Task Force. Annals of Internal Medicine, 153, 396–399. https://doi.org/10.7326/0003-4819-153-6-201009210-00007; https://journals.sagepub.com/author-instructions/
JMH?_gl=1*u3rvu6*_ga*MTQ3OTE5National Cancer Institute. (2021). Surveillance, epidemiology and end results program—Cancer stat facts: Testicular 
cancer. https://seer.cancer.gov/statfacts/html/testis.htmlRovito, M. J., Adams, W. B., Craycraft, M., Gooljar, C., Maresca, M., Guelmes, J., & Gallelli, A. (2022). The asso-ciation between testicular self-examination and stages of testicular cancer diagnosis: A cross-sectional analysis. 
Journal of Adolescent and Young Adult Oncology, 11(1), 41–47.Rovito, M. J., Leone, J. E., & Cavayero, C. T. (2018). “Off-label” usage of testicular self-examination (TSE): Benefits 
beyond cancer detection. American Journal of Men’s 
Health, 12(3), 505–513.

---

### Chunk 9/30
**Article:** DISFUNÇÃO ERÉTIL (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.564

elatar redução na libido, frequência das relações e na rigidez da ereção, especialmente em casos de hipogonadismo. A ausência de ereções matinais (tumescência peniana noturna) também é um sintoma importante, frequentemente associado à apneia do sono.
## Objetivo:
*   **Exame Físico:**
    *   Avaliação da composição corporal (bioimpedância, antropometria ou medição da circunferência abdominal).
    *   Exame genital para avaliar atrofia testicular, palpação do pênis para identificar calcificações ou fibroses (sugestivo de Doença de Peyronie).
    *   Verificação de ginecomastia.
    *   Busca por cicatrizes de cirurgias prévias na região perineal, inguinal e baixo ventre.
*   **Questionários:** Uso do questionário validado "Índice Internacional de Função Erétil" para estratificar o grau da disfunção (leve, moderada ou severa).

---

### Chunk 10/30
**Article:** Testicular Cancer: Clinical Practice Guidelines for Diagnosis, Treatment and Follow-Up (2015)
**Journal:** European Urology
**Section:** abstract | **Similarity:** 0.564

Context: Testicular cancer is the most common solid malignancy affecting males aged 15-35 years. Early detection through clinical examination and ultrasound is crucial. Objective: To provide evidence-based guidelines for diagnosis and management of testicular cancer. Methods: Systematic review of literature and expert consensus. Results: Physical examination and testicular ultrasound are recommended for suspicious masses. Tumor markers (AFP, beta-hCG, LDH) should be obtained. Self-examination education improves early detection rates.

---

### Chunk 11/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** discussion | **Similarity:** 0.563

F.;Nardone,V.;Guida,C.;Scialpi,M.;Cappabianca,S.DoesmultiparametricUSimprovediagnosticaccuracyinthecharacterizationofsmalltesticularmasses?Gland.Surg.2019,8,S136–S141.[CrossRef][PubMed]42.Schwarze,V.;Marschner,C.;Sabel,B.;deFigueiredo,G.N.;Marcon,J.;Ingrisch,M.;Knösel,T.;Rübenthaler,J.;Clevert,D.-A.Multiparametricultrasonographicanalysisoftesticulartumors:Asingle-centerexperienceinacollectiveof49patients.Scand.J.Urol.2020,54,241–247.[CrossRef][PubMed]43.Eiﬂer,J.B.,Jr.;King,P.;Schlegel,P.N.IncidentalTesticularLesionsFoundDuringInfertilityEvaluationareUsuallyBenignandMaybeManagedConservatively.J.Urol.2008,180,261–265.[CrossRef][PubMed]44.Scandura,G.;Verrill,C.;Protheroe,A.;Joseph,J.;Ansell,W.;Sahdev,A.;Shamash,J.;Berney,D.M.Incidentallydetectedtesticularlesions<10mmindiameter:Canorchidectomybeavoided?BJUInt.2017,121,575–582.[CrossRef][PubMed]45.Luzurier,A.;Maxwell,F.;Correas,J.;Benoit,G.;Izard,V.;Ferlicot,S.;Teglas,J.;Bellin,M.;Rocher,L.Qualitativeandquantitativecontrast-enha

---

### Chunk 12/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** other | **Similarity:** 0.558

ectadministration,D.Y.H.;fundingacquisition,notapplicable.Allauthorshavereadandagreedtothepublishedversionofthemanuscript.Funding:Thisresearchreceivednoexternalfunding.InstitutionalReviewBoardStatement:ThisstudywasconductedinaccordancewiththeDeclarationofHelsinki,andtheprotocolwasapprovedbytheEthicsCommitteeoftheInstitutionalReviewBoardofKing’sCollegeHospitalNHSFoundationTrust(projectidentiﬁcationcode:KCH14-102(IRAS148856).Dateofapproval:12October2018).InformedConsentStatement:Waivedduetotheretrospectivenatureofthisstudy.DataAvailabilityStatement:Thedatasetspresentedinthisarticleareunavailableduetoprivacyrestrictions.ConﬂictsofInterest:Theauthorsdeclarenoconﬂictsofinterest.References1.Tsili,A.C.;Bougia,C.K.;Pappa,O.;Argyropoulou,M.I.Ultrasonographyofthescrotum:Revisitingaclassictechnique.Eur.J.Radiol.2021,145,110000.[CrossRef][PubMed]2.Dogra,V.S.;Gottlieb,R.H.;Oka,M.;Rubens,D.J.SonographyoftheScrotum.Radiology2003,227,18–36.[CrossRef][PubMed]3.Bhatt,S.;Jafri,S.Z.H.;Wasserman,

---

### Chunk 13/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** other | **Similarity:** 0.557

ScrotalPathology.Radiol.Clin.NorthAm.2012,50,317–332.[CrossRef][PubMed]8.Altaffer,L.F.;Steele,S.M.ScrotalExplorationsNegativeforMalignancy.J.Urol.1980,124,617–619.[CrossRef][PubMed]9.Toren,P.J.;Roberts,M.;Lecker,I.;Grober,E.D.;Jarvi,K.;Lo,K.C.SmallIncidentallyDiscoveredTesticularMassesinInfertileMen—IsActiveSurveillancetheNewStandardofCare?J.Urol.2010,183,1373–1377.[CrossRef]10.Carmignani,L.;Gadda,F.;Mancini,M.;Gazzano,G.;Nerva,F.;Rocco,F.;Colpi,G.M.DETECTIONOFTESTICULARULTRA-SONOGRAPHICLESIONSINSEVEREMALEINFERTILITY.J.Urol.2004,172,1045–1047.[CrossRef]11.Carmignani,L.;Gadda,F.;Gazzano,G.;Nerva,F.;Mancini,M.;Ferruti,M.;Bulfamante,G.;Bosari,S.;Coggi,G.;Rocco,F.;etal.HighIncidenceofBenignTesticularNeoplasmsDiagnosedbyUltrasound.J.Urol.2003,170,1783–1786.[CrossRef][PubMed]12.Zuniga,A.;Lawrentschuk,N.;Jewett,M.A.S.Organ-sparingapproachesfortesticularmasses.Nat.Rev.Urol.2010,7,454–464.[CrossRef][PubMed]13.Tsili,A.C.;Bertolotto,M.;Turgut,A.T.;Dogra,V.;Freeman,S.;Rocher,L.;Bel  

---

### Chunk 14/30
**Article:** A Call to Action to Review the USPSTF's Recommendation for Testicular Self-Examination (2022)
**Journal:** American Journal of Men's Health
**Section:** discussion | **Similarity:** 0.548

ty. (2022). Cancer facts and figures 2022. 4. https://www.cancer.org/content/dam/cancer-org/research/cancer-facts-and-statistics/annual-cancer-facts-and-figures/2022/2022-cancer-facts-and-figures.pdfAmerican Urological Association. (2014). AUA Men’s Health Checklist. https://www.google.com/url?sa=t&rct=j&q=&esrc=s&source=web&cd=&ved=2ahUKEwimrKqY65r3AhV8l2oFHZarAtIQFnoECAkQAQ&ur
l=https%3A%2F%2Fwww.auanet.org%2Fdocuments%2Feducation%2Fclinical-guidance%2FMens-Health-
Checklist.pdf&usg=AOvVaw37zmiamYQgrxnHFMZcr-y8&cshid=1650189706876783Barbonetti, A., Martorella, A., Minaldi, E., D’Andrea, S., Bardhi, D., Castellini, C., Francavilla, F., & Francavilla, S. (2019). Testicular cancer in infertile men with and without testicular microlithiasis: A systematic review and meta-analysis of case-control studies. Frontiers in Endocrinology, 10, Article 164.Calonge, N. (2005). Screening for testicular cancer: Recommendation statement. U.S. Preventive Services  Task Force.

---

### Chunk 15/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** discussion | **Similarity:** 0.547

dinourdepartmentforthefollowingindications:evaluationandlocationofpalpablescrotalmasses,detectionofprimarytumours,follow-upofpa-tientswithtesticularmicrolithiasis,follow-upofpatientswithpreviouslymphoma,acutescrotum,scrotaltrauma,localisationoftheundescendedtestis,detectionofvaricocelesininfertilemen,andevaluationoftesticularischaemia.Fromaninitialdatasetofallscrotalultrasoundexaminations,124consecutivecasesoffocaltesticularabnormalitiesinvestigatedbyMPUSwereselectedforanalysis.InstitutionalReviewBoardapprovalwasobtained,withallproceduresperformedinaccordancewithethicalstandardsandpatientconﬁdentialityguidelines.2.2.PatientCohortandDataAcquisitionEligiblecaseswereidentiﬁedfromthedepartmentalultrasounddatabasebasedontheinclusioncriteriaofhavingundergoneMPUS,comprisinggreyscaleUS,CDUS,CEUS,andSE.Comprehensiveclinicaldata,includingthepatients’ages,clinicalpresentations,tumourmarkers,histopathologicalreports,andfollow-upoutcomes,wereextractedfromelectronicpatientrecords.Imagingdatawe

---

### Chunk 16/30
**Article:** DISFUNÇÃO ERÉTIL (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.544

ixa tolerância a esforço correlaciona-se com pior desempenho sexual; predomínio simpático (estresse) prejudica ereção.
- Sono e hormônios: apneia obstrutiva do sono reduz testosterona, aumenta endotelina e piora o IIEF; sono é crucial para produção hormonal.
- Exame físico direcionado: testículos (atrofia), ginecomastia (predominância estrogênica), cicatrizes e cirurgias prévias, doença de Peyronie (placas/fibroses), composição corporal (bioimpedância/ISAK; circunferência abdominal >94 e >102 como pontos de risco).
- Exames laboratoriais e imagem: painel hormonal, inflamatório, renal/hepático, lipidograma, PSA quando indicado; ecografia abdominal; risco cardiovascular (teste ergométrico, ecocardiograma, tomografia com escore de cálcio coronariano); polissonografia domiciliar para sono.
### 4.

---

### Chunk 17/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** results | **Similarity:** 0.533

ingwithsymptomssuchasscrotalpainorsubfertility.Includingbothpalpableandimpalpablelesionsallowsustoassesstheeffectivenessofultrasoundinarealisticclinicalsettingandreﬂecttheheterogeneityobservedinroutinepractice.However,thisbroadinclusioncomplicatestheextractionofpreciseguidanceforpopulationswithspeciﬁcclinicalpresentations.Theseriessizewasmodest,coveringawidespectrumoftesticularpathologiesbutwithlimitedcasesperspeciﬁccondition,suchasmalignantlesions<10mm.Thissuggestsaneedforfuturestudiesonalargerscaletovalidatetheseresultswithgreaterstatisticalpower.Additionally,thisstudyfo-cusedonaselectgroupofpatientswithfocaltesticularabnormalitiesidentiﬁedthroughgreyscaleultrasound,riskingselectionbias.Theoperator-dependentnatureofultrasoundalsopresentsachallengeinensuringconsistentandreproducibleresultsacrossdifferentexaminers.Lastly,non-surgicalcasesmanagedasbenignwerenotconﬁrmedhistologically,leavingroomfordiagnosticuncertainty.5.ConclusionsOurdecade-longexperienceindicatesthatintegrati

---

### Chunk 18/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** discussion | **Similarity:** 0.533

onofthesurgicalteamtoaidinimmediatehistopathologicalevaluation.2.5.DataAnalysisTwoexperiencedreviewers,eachwithoverﬁveyearsofexpertiseinscrotalultra-sonographyandblindedtotheother’sevaluations,independentlyrecordedtheultrasoundfeaturesofeachlesion.Toensuretheobjectivityandintegrityofourdataanalysis,bothreviewerswerecompletelyblindedtoallclinicalinformation,includingthepresenceorabsenceofraisedtumourmarkersanddistantmetastasis.Thereviewers’assessmentswerebasedsolelyontheimagingdatapresentedtothem,devoidofanypreconceivednotionsaboutthepatients’clinicalstatus.Discrepancieswereresolvedthroughjointdiscussionsbetweenthereviewerstoreachaconsensus.Theassessmentfocusedondocumentingessentialsonographiccharacteristics,suchassize,echogenicity,vascularpatterns,contrastenhancementproperties,andstrainelastographyﬁndings.Strainelastographyresultswereanalysedusingacolour-codedschemetodelineatevaryingdegreesoftissuestiffnessinaccordancewithestablishedcriteriafortesticularstrainelastography[26,

---

### Chunk 19/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** other | **Similarity:** 0.531

ntstothediagnosis.Br.J.Radiol.2012,85,S41–S53.[CrossRef][PubMed]33.Huang,D.Y.;Pesapane,F.;Rafailidis,V.;Deganello,A.;Sellars,M.E.;Sidhu,P.S.Theroleofmultiparametricultrasoundinthediagnosisofpaediatricscrotalpathology.Br.J.Radiol.2020,93,20200063.[CrossRef][PubMed]34.Mansoor,N.M.;Huang,D.Y.;Sidhu,P.S.Multiparametricultrasoundimagingcharacteristicsofmultipletesticularadrenalresttumoursincongenitaladrenalhyperplasia.Ultrasound2022,30,80–84.[CrossRef][PubMed]35.Sidhu,P.S.MultiparametricUltrasound(MPUS)Imaging:TerminologyDescribingtheManyAspectsofUltrasonography.UltraschallMed.2015,36,315–317.[CrossRef]36.Itoh,A.;Ueno,E.;Tohno,E.;Kamma,H.;Takahashi,H.;Shiina,T.;Yamakawa,M.;Matsumura,T.BreastDisease:ClinicalApplicationofUSElastographyforDiagnosis.Radiology2006,239,341–350.[CrossRef]37.Kim,I.;Young,R.H.;Scully,R.E.Leydigcelltumorsofthetestis.Am.J.Surg.Pathol.1985,9,177–192.[CrossRef]38.Farkas,L.M.;Székely,J.G.;Pusztai,C.;Baki,M.HighFrequencyofMetastaticLeydigCellTesticularTumours.O

---

### Chunk 20/30
**Article:** A Call to Action to Review the USPSTF's Recommendation for Testicular Self-Examination (2022)
**Journal:** American Journal of Men's Health
**Section:** other | **Similarity:** 0.529

J., Leone, J. E., & Cavayero, C. T. (2018). “Off-label” usage of testicular self-examination (TSE): Benefits 
beyond cancer detection. American Journal of Men’s 
Health, 12(3), 505–513.

6 American Journal of Men’s Health Rovito, M. J., Manjelievskaia, J., Leone, J. E., Lutz, M. J., & Nangia, A. (2016). From “D” to “I”: A critique of the current United States preventive services task force rec-ommendation for testicular cancer screening. Preventive Medicine Reports, 3, 361–366. https://doi.org/10.1016/j.pmedr.2016.04.006Rovito, M. J., & Nangia, A. K. (2019, February 14). Testis CA screening: Why USPSTF “D” grade is misguided. Urology 
Times. https://www.urologytimes.com/view/testis-ca-
screening-why-uspstf-d-grade-misguidedSaab, M. M., Hegarty, J., & Landers, M. (2019). Testicular aware-ness: The what, the why, and the how. International Journal of Mens Social and Community Health, 2(1), e1–e10.Saab, M. M., Landers, M., & Hegarty, J. (2018).

---

### Chunk 21/30
**Article:** MFI - Reposição Hormonal - AULA 04 (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.525

tosterona livre” calculada do soro como decisor; incluir, quando indicado, painel salivar (testosterona, DHT, estradiol; progesterona no D22–D24) juntamente com sangue total e SHBG.
- [ ] 5. Implementar triagem de fatores ambientais/ocupacionais que elevem temperatura escrotal (vestimenta apertada, longos períodos sentado, dormir de cueca, ambientes quentes) e orientar medidas corretivas.
- [ ] 6. Estabelecer protocolo para avaliação pós-ciclo de testosterona (endógena/exógena), reconhecendo períodos de LH/FSH inibidos e evitando interpretações equivocadas de queda transitória.
- [ ] 7. Preparar leitura dos estudos recomendados sobre obesidade e hipogonadismo, bariátrica e reversão hormonal, e relações entre obesidade e andropausa, para discussão na próxima aula.
- [ ] 8. Educar equipes clínicas sobre a inadequação de prescrever inibidores de PDE5 (Viagra/Cialis) sem avaliação hormonal quando há suspeita de disfunção androgênica.
- [ ] 9.

---

### Chunk 22/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** introduction | **Similarity:** 0.524

fadvancedultrasoundtechniquesinenhancingtheevaluationoffocaltesticularabnormalitiesinclinicalpracticeandcouldaidashifttowardstestis-sparingmanagementstrategies.Keywords:multiparametric;ultrasound;testicularcancer;testis-sparingsurgery;orchiectomy
1.IntroductionConventionalultrasonography(US),includinggreyscaleimagingandcolourDopplerUS(CDUS),standsasthecornerstoneforevaluatingscrotalpathologiesduetoitshighresolution,availability,cost-effectiveness,andabsenceofionizingradiation[1–4].Despite
Cancers2024,16,2309.https://doi.org/10.3390/cancers16132309https://www.mdpi.com/journal/cancers

Cancers2024,16,2309
2of17
itswidespreaduse,thespeciﬁcityofgreyscaleultrasoundincharacterisingscrotalmassesremainslimited,oftenleavingthenatureofsuchlesionsambiguous[5–7].Traditionally,solidtesticularlesions,especiallythosepresentingaspalpablelumps,haveledtoradi-calorchidectomy[8,9].However,thelandscapeofscrotalultrasonographyhasevolvedsigniﬁcantlywithadvancementsintechnologyandtechnique,includinghi

---

### Chunk 23/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** other | **Similarity:** 0.524

resEnUrol.2015,25,75–82.[CrossRef]52.Ma,W.;Sarasohn,D.;Zheng,J.;Vargas,H.A.;Bach,A.CausesofAvascularHypoechoicTesticularLesionsDetectedatScrotalUltrasound:CanTheyBeConsideredBenign?Am.J.Roentgenol.2017,209,110–115.[CrossRef][PubMed]53.Patrikidou,A.;Cazzaniga,W.;Berney,D.;Boormans,J.;deAngst,I.;DiNardo,D.;Fankhauser,C.;Fischer,S.;Gravina,C.;Gremmels,H.;etal.EuropeanAssociationofUrologyGuidelinesonTesticularCancer:2023Update.Eur.Urol.2023,84,289–301.[CrossRef][PubMed]54.Maxwell,F.;Savignac,A.;Bekdache,O.;Calvez,S.;Lebacle,C.;Arama,E.;Garrouche,N.;Rocher,L.LeydigCellTumorsoftheTestis:AnUpdateoftheImagingCharacteristicsofaNotSoRareLesion.Cancers2022,14,3652.[CrossRef]

Cancers2024,16,2309
17of17
55.Nicolai,N.;Necchi,A.;Raggi,D.;Biasoni,D.;Catanzaro,M.;Piva,L.;Stagni,S.;Maffezzini,M.;Torelli,T.;Faré,E.;etal.ClinicalOutcomeinTesticularSexCordStromalTumors:TestisSparingvsRadicalOrchiectomyandManagementofAdvancedDisease.Urology2015,85,402–406.[CrossRef][PubMed]56.Westlander,G.;Ekerhov

---

### Chunk 24/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.522

Específico (PSA)**:
    *   **Função**: Enzima que liquefaz o sêmen.
    *   **Formas**: Complexado (maior parte) e Livre. O PSA total é a soma de ambas.
    *   **Interpretação Clínica**: A relação PSA livre / PSA total é crucial.
        *   **> 0.14 (ou 14%)**: Sugere HPB.
        *   **< 0.14 (ou 14%)**: Aumenta o risco de câncer de próstata.
    *   **Limitações**: Cerca de 1-4% dos cânceres de próstata ocorrem com PSA normal. Em homens com baixa testosterona, esse número pode chegar a 15%.
*   **Exames de Imagem**:
    *   **Ultrassonografia de Próstata com Estudo do Resíduo Pós-Miccional**: Avalia anatomia e função. Um resíduo pós-miccional > 40 ml indica obstrução.
    *   **Urofluxometria**: Indicada para sintomas obstrutivos. Mede a velocidade do fluxo urinário (Qmax). Valores < 10 ml/segundo indicam obstrução.
    *   **Ressonância Magnética Multiparamétrica 3-Tesla (3T)**: Exame de alta especificidade.

---

### Chunk 25/30
**Article:** A Call to Action to Review the USPSTF's Recommendation for Testicular Self-Examination (2022)
**Journal:** American Journal of Men's Health
**Section:** other | **Similarity:** 0.522

control studies. Frontiers in Endocrinology, 10, Article 164.Calonge, N. (2005). Screening for testicular cancer: Recommendation statement. U.S. Preventive Services  Task Force. American Family Physician, 72(10), 2069–2070.Fadich, A., Giorgianni, S. J., Rovito, M. J., Pecchia, G. A., Bonhomme, J. J., Adams, W. B., Stephenson, C. L., Mesa-Morales, F. E., & Sparkes, J. S. (2018). USPSTF testicular examination nomination-self-exam-
inations and examinations in a clinical setting. American 
Journal of Men's Health, 12(5), 1510–1516. https://doi.org/10.1177/1557988318768597Ghazarian, A. A., & McGlynn, K. A. (2020). Increasing inci-dence of testicular germ cell tumors among racial/ethnic 
minorities in the United States. Cancer Epidemiology and 
Prevention Biomarkers, 29(6), 1237–1245. https://doi.org/10.1158/1055-9965.EPI-20-0107Institute of Medicine. (2011). Clinical practice guidelines we can trust. National Academies Press. https://doi.
org/10.17226/13058Kim, C., McGlynn, K.

---

### Chunk 26/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** other | **Similarity:** 0.520

IntratesticularLesionsWithStrainElastographyUsingStrainRatioandColorMapVisualGrading:DifferentiationofNeoplasticandNonneoplasticLesions.J.UltrasoundMed.2018,38,223–232.[CrossRef][PubMed]27.Yusuf,G.;Sellars,M.E.;Kooiman,G.G.;Diaz-Cano,S.;Sidhu,P.S.GlobalTesticularInfarctioninthePresenceofEpididymitis.J.UltrasoundMed.2013,32,175–180.[CrossRef][PubMed]28.Rafailidis,V.;Robbie,H.;Konstantatou,E.;Huang,D.Y.;Deganello,A.;ESellars,M.;Cantisani,V.;Isidori,A.M.;Sidhu,P.S.Sonographicimagingofextra-testicularfocallesions:Comparisonofgrey-scale,colourDopplerandcontrast-enhancedultrasound.Ultrasound2016,24,23–33.[CrossRef][PubMed]29.Patel,K.V.;Huang,D.Y.;Sidhu,P.S.Metachronousbilateralsegmentaltesticularinfarction:Multi-parametricultrasoundimagingwithgrey-scaleultrasound,Dopplerultrasound,contrast-enhancedultrasound(CEUS)andreal-timetissueelastography(RTE).J.Ultrasound2014,17,233–238.[CrossRef]30.Zebari,S.;Huang,D.Y.;Wilkins,C.J.;Sidhu,P.S.AcuteTesticularSegmentalInfarctFollowingEndovascular

---

### Chunk 27/30
**Article:** DISFUNÇÃO ERÉTIL (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.519

[ ] Medir circunferência abdominal; se >94, reforçar intervenção; se >102, considerar alto risco e intensificar manejo da síndrome metabólica.
- [ ] Exame físico genital completo: testículos, ginecomastia, placas/curvatura peniana; investigar cicatrizes/cirurgias prévias.
- [ ] Solicitar exames básicos: painel hormonal (incluindo testosterona total/livre), PSA quando indicado, função renal/hepática, inflamatórios, lipidograma; complementar conforme caso.
- [ ] Solicitar ecografia abdominal total (próstata, fígado/esteatose, rins) e, conforme risco, tomografia com escore de cálcio coronariano; considerar teste ergométrico/ecocardiograma.
- [ ] Investigar sono com polissonografia domiciliar em presença de ronco, sonolência, despertares ou redução de ereções matinais.
- [ ] Revisar medicações: 5-ARIs, psicotrópicos, estatinas, anti-hipertensivos, lisdexanfetamina; discutir alternativas e risco/benefício.

---

### Chunk 28/30
**Article:** Varicocele and Male Infertility: Current Concepts and Future Perspectives (2017)
**Journal:** Nature Reviews Urology
**Section:** abstract | **Similarity:** 0.519

Varicocele is present in approximately 15% of the general male population and up to 35% of men with primary infertility. Physical examination remains the gold standard for diagnosis, supplemented by scrotal ultrasound in equivocal cases. Varicocele can impair spermatogenesis through multiple mechanisms including testicular hyperthermia, oxidative stress, and hormonal imbalances. Varicocelectomy is recommended for men with palpable varicocele, abnormal semen parameters, and infertility. Treatment can improve sperm parameters in 60-80% of cases and pregnancy rates in couples.

---

### Chunk 29/30
**Article:** Multiparametric Ultrasound for Focal Testicular Pathology: A Ten-Year Retrospective Review (2024)
**Journal:** Cancers (Basel)
**Section:** other | **Similarity:** 0.518

,M.;Bob,F.;Bojunga,J.;Brock,M.;etal.TheEFSUMBGuidelinesandRecommendationsfortheClinicalPracticeofElastographyinNon-HepaticApplications:Update2018.UltraschallMed.2019,40,425–453.[CrossRef][PubMed]16.Isidori,A.M.;Pozza,C.;Gianfrilli,D.;Giannetta,E.;Lemma,A.;Poﬁ,R.;Barbagallo,F.;Manganaro,L.;Martino,G.;Lombardo,F.;etal.DifferentialDiagnosisofNonpalpableTesticularLesions:QualitativeandQuantitativeContrast-enhancedUSofBenignandMalignantTesticularTumors.Radiology2014,273,606–618.[CrossRef][PubMed]17.Fang,C.;Huang,D.Y.;Sidhu,P.S.Elastographyoffocaltesticularlesions:Currentconceptsandutility.Ultrasonography2019,38,302–310.[CrossRef][PubMed]18.Auer,T.;DeZordo,T.;Dejaco,C.;Gruber,L.;Pichler,R.;Jaschke,W.;Dogra,V.S.;Aigner,F.ValueofMultiparametricUSintheAssessmentofIntratesticularLesions.Radiology2017,285,640–649.[CrossRef][PubMed]19.Bertolotto,M.;Muça,M.;Currò,F.;Bucci,S.;Rocher,L.;Cova,M.A.MultiparametricUSforscrotaldiseases.Abdom.Imaging2018,43,899–917.[CrossRef]20.Pinto,S.P.;Hua

---

### Chunk 30/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.515

urinário (Qmax). Valores < 10 ml/segundo indicam obstrução.
    *   **Ressonância Magnética Multiparamétrica 3-Tesla (3T)**: Exame de alta especificidade. O palestrante solicita de rotina para homens > 50 anos, ou > 40 anos com histórico familiar ou alterações súbitas no PSA.
*   **Dosagem Hormonal Salivar e Quociente Estrogênico**:
    *   **Vantagens da Saliva**: Via não invasiva que mede a fração livre e 100% bioativa dos hormônios (Testosterona, DHT, Estradiol, etc.). Útil quando a clínica do paciente não corresponde aos exames de sangue.
    *   **Quociente Estrogênico**: Fórmula para avaliar o risco de doenças prostáticas.
        *   **Fórmula**: Estriol / (Estradiol + Estrona).
        *   **Valores > 1**: Bom prognóstico (perfil estrogênico protetor).
        *   **Valores < 1**: Mau prognóstico (prevalência de estrogênios proliferativos), sugerindo a necessidade de trabalhar a metilação da estrona.

---

