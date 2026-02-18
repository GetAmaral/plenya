# ScoreItem: Nefrite

**ID:** `019bf31d-2ef0-7103-bda8-a522bc357978`
**FullName:** Nefrite (Histórico de doenças - Doenças crônicas (Questionar ativamente as doenças crônicas mais comuns, perguntar sobre duração, grau de controle e tratamentos utilizados no passado e atualmente) - Outras doenças renais)

**Preparation Metadata:**
- Quality Grade: **GOOD**
- Total Chunks: 30 de 9 artigos
- Avg Similarity: 0.536

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-7103-bda8-a522bc357978`.**

```json
{
  "score_item_id": "019bf31d-2ef0-7103-bda8-a522bc357978",
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

**ScoreItem:** Nefrite (Histórico de doenças - Doenças crônicas (Questionar ativamente as doenças crônicas mais comuns, perguntar sobre duração, grau de controle e tratamentos utilizados no passado e atualmente) - Outras doenças renais)

**30 chunks de 9 artigos (avg similarity: 0.536)**

### Chunk 1/30
**Article:** Dislipidemias II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.598

dade, inatividade física, pressão arterial e dislipidemia.
- A inflamação é indicada por marcadores como a Proteína C-Reativa (PCR), onde um valor de 5 já é considerado elevado, e a Interleucina 6 (exemplo de paciente com 8.45).
- A resposta anti-inflamatória é medida pela Interleucina 10, com um valor de corte de 3,5 (abaixo disso é um risco) e um exemplo de paciente com 6.44.
- A saúde do endotélio é avaliada pelo óxido nítrico, que deve estar em 8.8; um paciente com 4.8 já apresenta disfunção sistêmica.
- A LDL oxidada é um marcador crítico, com um limite saudável de 133, enquanto pacientes de alto risco podem apresentar valores extremos, como 1000.
**Achados Adicionais**
- Existem cinco parâmetros clássicos para definir a síndrome metabólica.
- O alvo de LDL para pacientes de alto risco, segundo as diretrizes atuais, é de 50.

---

### Chunk 2/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.562

arltrationrate[eGFRcr-cys])(1B).1.1.3EvaluationofchronicityPracticePoint1.1.3.1:Proofofchronicity(durationofaminimumof3months)canbeestablishedby:(i)reviewofpastmeasurements/estimationsofGFR;(ii)reviewofpastmeasurementsofalbuminuriaorproteinuriaandurinemicroscopicexaminations;(iii)imagingndingssuchasreducedkidneysizeandreductionincorticalthickness;(iv)kidneypathologicalndingssuchasbrosisandatrophy;(v)medicalhistory,especiallyconditionsknowntocauseorcontributetoCKD;(vi)repeatmeasurementswithinandbeyondthe3-monthpoint.PracticePoint1.1.3.2:DonotassumechronicitybaseduponasingleabnormallevelforeGFRandACR,asthendingcouldbetheresultofarecentacutekidneyinjury(AKI)eventoracutekidneydisease(AKD).PracticePoint1.1.3.3:ConsiderinitiationoftreatmentsforCKDatrstpresentationofdecreasedGFRorelevatedACRifCKDisdeemedlikelyduetopresenceofotherclinicalindicators.1.1.4Evaluationofcause
PracticePoint1.1.4.1:EstablishthecauseofCKDusingclinicalcontext,personalandfamilyhistory,socialandenvi-ronmentalfactor

---

### Chunk 3/30
**Article:** Fisiologia e Bioquímica do Sistema Imune II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.560

olicitar e registrar parâmetros essenciais (PCR-us, VHS); calcular índices estimáveis (HGI, TAIG); considerar complementares (TNF-α, IL-6, GPx, MDA, antioxidantes totais).
- [ ] 11. Monitorar trimestralmente HbA1c, glicemia de jejum e HGI; estabelecer metas individuais (ex.: reduzir de 6,1 para ~5,5; longo prazo ~5,3).
- [ ] 12. Calcular TG/HDL e integrar com lipidograma/SREBP1c/2; evitar combinação de gordura saturada com açúcar e excesso de saturadas em múltiplas refeições.
- [ ] 13. Avaliar ferro, ferritina, transferrina e saturação (20–50%; evitar <20%); interpretar com hepcidina/SREBP1c e quadro inflamatório.
- [ ] 14. Medir TNF-α (<8,1) e IL-6 (<3,4) para acompanhar atividade inflamatória; relacionar com obesidade inflamada.
- [ ] 15. Calcular HOMA-β (167–175) e HOMA-IR (<2,15); buscar glicemia 60–90 e insulina ~6–7.
- [ ] 16. Monitorar homocisteína (<7,9) e PCR; usar PCR-us; documentar crises (PCR >1.000) e conduzir manejo apropriado.
- [ ] 17.

---

### Chunk 4/30
**Article:** Automated urine sediment analyzers underestimate the severity of hematuria in glomerular diseases (2021)
**Journal:** Scientific Reports
**Section:** results | **Similarity:** 0.558

and digital imaging in one system: A multicenter performance evaluation of the cobas 6500 urine work area. Pract. Lab. Med. 17, e00139 (2019). 30. aller, T. R. & Wang, L. P. Evaluation of asymptomatic microscopic hematuria in adults. Am. Fam. Phys. 60, 11431152 (1999). 31. Palsson, R., Srivastava, A. & Waikar, S. S. Performance of the automated urinalysis in diagnosis of proliferative glomerulonephritis. Kidney Int. Rep. 4, 723727 (2019). 32. Orlandi, P. F. et al. Hematuria as a risk factor for progression of chronic kidney disease and death: Findings from the chronic renal insuciency cohort (CRIC) study. BMC Nephrol. 19, 150 (2018). 33. Schulman, G. et al. Risk factors for progression of chronic kidney disease in the EPPIC trials and the eect of AST-120. Clin. Exp. Nephrol. 22, 299308 (2018). 34. Mashitani, T. et al.

---

### Chunk 5/30
**Article:** Fisiologia e Bioquímica do Sistema Imune II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.556

60 e 90
- [ ] Manter a insulina, o mais baixo possível, 6, 7, estourando 8
- [ ] Avaliar a homocisteína, pois é um marcador inflamatório importante
- [ ] Usar a proteína C-reativa, associado com os níveis de homocisteína
- [ ] Verificar os parâmetros essenciais na avaliação inflamatória
- [ ] Estimar o índice de glicação e o índice TAIG, baseado nos resultados essenciais
- [ ] Complementar a avaliação com TNF-alfa, IL-6, glutationa e malon de aldeído
### Tarefas para @
- [ ] Usar um concentrado de C8 ou um mix de C8 e C10, para estimular mais ainda o CP3 e as UCPs (proteínas desacopladoras), diminuir a produção de espécie reativa de oxigênio e aumentar a oxidação de gordura @
- [ ] Incluir mioinositol, trans-resveratrol e epigalocatequina galato na formulação, para diminuir os compostos de glicação avançada e a hemoglobina glicada @
- [ ] Fazer uma boa distribuição de gordura e trabalhar os ácidos graxos de cadeia curta, para obter o melhor benefício p

---

### Chunk 6/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.548

specickidneyfunctionofglomerularltration.Usethemoregeneralterm“kidneyfunction(s)”whendealingwiththetotalityoffunctionsofthekidney.
PhysicalexamNephrotoxicmedicationsSymptoms and signsof urinary tractabnormalitiesSymptoms and signsof systemic diseasesLaboratory tests, imaging, and tissue sample, such as:• Urinalysis and urine sediment• Urine albumin-to-creatinine ratio• Serologic tests
• Ultrasound• Kidney biopsy• Genetic testingMedicalhistorySocial andenvironmentalhistoryObtain careful family historyfor possible genetic causes,including family pedigree for CKD
Figure8|Evaluationofcauseofchronickidneydisease(CKD).

---

### Chunk 7/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.547

RifCKDisdeemedlikelyduetopresenceofotherclinicalindicators.1.1.4Evaluationofcause
PracticePoint1.1.4.1:EstablishthecauseofCKDusingclinicalcontext,personalandfamilyhistory,socialandenvi-ronmentalfactors,medications,physicalexamination,laboratorymeasures,imaging,andgeneticandpathologicdiagnosis(Figure8).
www.kidney-international.orgsummaryofrecommendationstatementsandpracticepointsKidneyInternational(2024)105(Suppl4S),S117–S314S149

PracticePoint1.1.4.2:Useteststoestablishacausebasedonresourcesavailable(Table622,98-100).Recommendation1.1.4.1:Wesuggestperformingakidneybiopsyasanacceptable,safe,diagnostictesttoevaluatecauseandguidetreatmentdecisionswhenclinicallyappropriate(2D).1.2EvaluationofGFR1.2.1OtherfunctionsofkidneysbesidesGFR
PracticePoint1.2.1.1:Usetheterm“GFR”whenreferringtothespecickidneyfunctionofglomerularltration.Usethemoregeneralterm“kidneyfunction(s)”whendealingwiththetotalityoffunctionsofthekidney.

---

### Chunk 8/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.545

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

### Chunk 9/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.544

k forrecurrence afterkidney transplantationConditions for whichgenetic testing isrelevant for reproductivecounselingExamples:• GLA (Fabry)• AGXT (primaryhyperoxaluria (PH))• CoQ10 genes (SRNS)• CTNS (cystinosis)• Tubulopathies(Na+, K+ etc.)Conditions amenable to nonspecific renoprotective strategiesExample:• COL4A3/4/5 (Alport)and RAAS blockadeExample:
• Glomerular diseasedue to mutations inAlport genes(COL4A3/4/5)Examples:• (CFH/CFI/C3...): aHUS• (AGXT, GRHPR, HOGA): primary hyperoxaluria (PH)
• Adenine phosphoribo-  syltransferase deficiency (APRT)Conditions amenable to specific screening for extrarenal manifestationsExamples:• HNF1B: diabetes• PKD1/PKD2(ADPKD): intracranialaneurysms • FLCN: renal cellcarcinoma, etc.Example:
• Prenatal/preimplantationdiagnosis
Figure9|Actionablegenesinkidneydisease.Actionabilityreferstothepotentialforgeneticrestresultstoleadtospecicclinicalactionsfrompreventionortreatmentofacondition,supportedbyrecommendationsbasedonevide

---

### Chunk 10/30
**Article:** Proteinuria (2023)
**Journal:** StatPearls Publishing
**Section:** abstract | **Similarity:** 0.539

Comprehensive clinical review of proteinuria covering definition, etiology, diagnosis, and management. Proteinuria serves as a marker of kidney damage assisting with diagnosis, prognosis, and therapy. Normal protein excretion is <150 mg/24 hours. Severity categories: nephritic range (150-3,500 mg/24h) and nephrotic range (>3,500 mg/24h). Urine dipstick provides semi-quantitative screening but has important limitations: false-positives with dehydration, alkaline urine, UTI; inability to detect positively charged proteins. UK CKD guidelines define significant proteinuria as UPCR >45 mg/mmol. Quantitative methods (spot UPCR preferred for convenience and reliability, or 24-hour urine collection for accuracy) are indicated when dipstick is positive or proteinuria >1 g/day is suspected. Early detection and quantification are crucial given associations with progressive renal disease and cardiovascular complications.

---

### Chunk 11/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.537

ropriatetreatment.KeyinformationBalanceofbenetsandharms.Thebenetsofkidneybi-opsyintermsofdiagnosis,prognosis,andplanningappro-
priatetreatmentforboththepersonwithCKDandhealthcare
PhysicalexamNephrotoxicmedicationsSymptoms and signsof urinary tractabnormalitiesSymptoms and signsof systemic diseasesLaboratory tests, imaging, and tissue sample, such as:• Urinalysis and urine sediment• Urine albumin-to-creatinine ratio• Serologic tests
• Ultrasound• Kidney biopsy• Genetic testingMedicalhistorySocial andenvironmentalhistoryObtain careful family historyfor possible genetic causes,including family pedigree for CKD
Figure8|Evaluationofcauseofchronickidneydisease(CKD).

---

### Chunk 12/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.535

dneyinjury(AKI)eventoracutekidneydisease(AKD).PracticePoint1.1.3.3:ConsiderinitiationoftreatmentsforCKDatrstpresentationofdecreasedGFRorelevatedACRifCKDisdeemedlikelyduetopresenceofotherclinicalindicators.Kidneydiseasesmaybeacuteorchronic.1,97Weexplicitlyyetarbitrarilydenethedurationofaminimumof3months(>90days)asdelineating“chronic”kidneydisease.TherationalefordeningchronicityistodifferentiateCKDfromAKDs(suchasacuteglomerulonephritis[GN]),includingAKI,whichmayrequiredifferenttimelinesfor
initiationoftreatments,differentinterventions,andhavedifferentetiologiesandoutcomes.Thedurationofkidneydiseasemaybedocumentedorinferredbasedonthe
clinicalcontext.Forexample,apersonwithdecreasedGFR
orkidneydamageduringanacuteillness,withoutprior
documentationofkidneydisease,maybeinferredtohave
AKD.ResolutionoverdaystoweekswouldconrmthediagnosisofAKIfromavarietyofdifferentcauses.A
personwithsimilarndingsintheabsenceofanacuteillnessmaybeinferredtohaveCKD,andiffollowedover
time,wouldbeconrmedtoha

---

### Chunk 13/30
**Article:** Automated urine sediment analyzers underestimate the severity of hematuria in glomerular diseases (2021)
**Journal:** Scientific Reports
**Section:** other | **Similarity:** 0.533

omparison of automated urinalysis systems and manual microscopy. Clin. Chim. Acta 384, 2834 (2007). 6. Cavanaugh, C. & Perazella, M. A. Urine sediment examination in the diagnosis and management of kidney disease: Core cur-riculum 2019. Am. J. Kidney Dis. 73, 258272 (2019). 7. Becker, G. J., Garigali, G. & Fogazzi, G. B. Advances in urine microscopy. Am. J. Kidney Dis. 67, 954964 (2016). 8. Oyaert, M. & Delanghe, J. Progress in automated urinalysis. Ann. Lab. Med. 39, 1522 (2019). 9. Lee, W., Ha, J. S. & Ryoo, N. H. Comparison of the automated cobas u 701 urine microscopy and UF-1000i ow cytometry systems and manual microscopy in the examination of urine sediments. J. Clin. Lab. Anal. 30, 663671 (2016). 10. Moreno, J. A. et al. Glomerular hematuria: Cause or consequence of renal inammation? Int. J. Mol. Sci. 20, 2205 (2019). 11. Kincaid-Smith, P. & Fairley, K. e investigation of hematuria. Semin. Nephrol. 25, 127135 (2005). 12.

---

### Chunk 14/30
**Article:** The Role of Urinary Nitrite in Predicting Bacterial Resistance in Urine Culture Analysis Among Patients With Uncomplicated Urinary Tract Infection (2022)
**Journal:** Cureus
**Section:** results | **Similarity:** 0.532

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

### Chunk 15/30
**Article:** Fisiologia e Bioquímica do Sistema Imune II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.529

nsulina ideal ~6–7 (até 8 em inflamação).
- Homocisteína: <7,9; correlaciona com PCR; polimorfismos FUT/MTHFR podem elevá-la.
- PCR: desejável <1; risco médio 1–3; alto >3; casos extremos >1.000 em crise; usar PCR-us para sensibilidade; VHS para trajetória da inflamação.
### 12. Preferências e filosofia de ciclo
- Estratégias cíclicas (jejum, low carb, cetogênica limpa, exercício em jejum) para treinar flexibilidade metabólica e melhorar uso de corpos cetônicos.
- Protocolos fundamentados em literatura (PubMed: “Inflammation and Phytochemicals/Bioactive Compounds”) e prática clínica, com resultados consistentes.
### 13. Observações Motivacionais e Agradecimentos
- Inspiração: “A constância dos bons resultados que conduz os homens à felicidade.”
- Agradecimentos e convite à prática contínua; compromisso de apoio.
## ❓ Perguntas
- [Inserir Pergunta/Confusão]
## 📚 Tarefas
- [ ] 1.

---

### Chunk 16/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.528

e-basedequations.TransplInt.2012;25:527–536.162.JanusN,Launay-VacherV,ByloosE,etal.CancerandrenalinsufciencyresultsoftheBIRMAstudy.BrJCancer.2010;103:1815–1821.163.Launay-VacherV,JanusN,DerayG.Renalinsufciencyandcancertreatments.ESMOOpen.2016;1:e000091.164.NaSY,SungJY,ChangJH,etal.Chronickidneydiseaseincancerpatients:anindependentpredictorofcancer-specicmortality.AmJNephrol.2011;33:121–130.165.RosnerMH,JhaveriKD,McMahonBA,etal.Onconephrology:the
intersectionsbetweenthekidneyandcancer.CACancerJClin.2021;71:47–77.166.SoveriI,BergUB,BjorkJ,etal.MeasuringGFR:asystematicreview.AmJKidneyDis.2014;64:411–424.167.WhiteCA,AkbariA,AllenC,etal.Simultaneousglomerularltrationratedeterminationusinginulin,iohexol,and99mTc-DTPAdemonstratestheneedforcustomizedmeasurementprotocols.KidneyInt.2021;99:957–966.168.XieP,HuangJM,LiuXM,etal.(99m)Tc-DTPArenaldynamicimaging
methodmaybeunsuitabletobeusedasthereferencemethodin
investigatingthevalidityofCDK-EPIequationfordetermining
glomerularltrati

---

### Chunk 17/30
**Article:** Automated urine sediment analyzers underestimate the severity of hematuria in glomerular diseases (2021)
**Journal:** Scientific Reports
**Section:** results | **Similarity:** 0.526

re compared using the Mann–Whitney U test. To measure the strength of association between the positive degree on the dipstick blood test and the urine RBC count, Spearmans rank-order correlation was used. e prevalence of categorical variables was compared between the two groups using the chi-squared test or Fishers exact test. Statistical analyses were performed using SPSS version 21 (IBM Co., Armonk, NY, USA). p values less than 0.05 were considered statistically signicant.ResultsPatient characteristics. In the data obtained via the UF-1000i urine analyzer, 330 patients were diagnosed with nephritic glomerular disease (235, IgA nephropathy; 39, pauci-immune crescentic glomerulonephritis; 35, lupus nephritis;13, membranoproliferative glomerulonephritis; 6, Henoch-Schönlein purpura;and 2, postinfec-tious glomerulonephritis). e median age at diagnosis was 47 (2858) years.

---

### Chunk 18/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.526

ionaltechnologies(e.g.,3Dultrasound)KidneybiopsyUltrasound-guidedpercutaneousUsuallyexaminedbylightmicroscopy,immunouorescence,andelectronmicroscopy,and,insomesituations,mayincludemoleculardiagnostics
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
signicance[MGRS])98Presenceofpersistenthematuriaoralbuminuriaiscriticalindetermining
differentialdiagnosisGenetictestingAPOL1,COL4A3,COL4A4,COL4A5,NPHS1,UMOD,
HNF1B,PKD1,PKD2Evolvingasatoolfordiagnos

---

### Chunk 19/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** discussion | **Similarity:** 0.525

ta-analysis.ClinJAmSocNephrol.2022;17:1305–1315.193.vanEeghenSA,WiepjesCM,T’SjoenG,etal.CystatinC-basedeGFRchangesduringgender-afrminghormonetherapyintransgenderindividuals.ClinJAmSocNephrol.2023;18:1545–1554.194.PierreC,MarzinkeM,AhmedSB,etal.AACC/NKFguidancedocumenton
improvingequityinchronickidneydiseasecare.JApplLabMed.2023;8:789–816.195.NgDK,FurthSL,WaradyBA,etal.Self-reportedrace,serumcreatinine,
cystatinC,andGFRinchildrenandyoungadultswithpediatrickidney
referenceswww.kidney-international.orgS298KidneyInternational(2024)105(Suppl4S),S117–S314

diseases:areportfromtheChronicKidneyDiseaseinChildren(CKiD)study.AmJKidneyDis.2022;80:174–185.e1.196.Luna-ZaizarH,Virgen-MontelongoM,Cortez-AlvarezCR,etal.Invitro
interferencebyacetaminophen,aspirin,andmetamizoleinserum
measurementsofglucose,urea,andcreatinine.ClinBiochem.2015;48:538–541.197.GreenbergN,RobertsWL,BachmannLM,etal.Specicitycharacteristicsof7commercialcreatininemeasurementproceduresbyenzymaticand
Jaffemethodpri

---

### Chunk 20/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.525

solutionoverdaystoweekswouldconrmthediagnosisofAKIfromavarietyofdifferentcauses.A
personwithsimilarndingsintheabsenceofanacuteillnessmaybeinferredtohaveCKD,andiffollowedover
time,wouldbeconrmedtohaveCKD.Inbothcases,repeatascertainmentofGFRandkidneydamageis
recommendedforaccuratediagnosisandstaging.The
timingoftheevaluationdependsonclinicaljudgment,with
earlierevaluationforthosesuspectedofhavingAKIand
laterevaluationforthosesuspectedofhavingCKD.ForpeoplewithriskfactorsforCKD,delayingdiag-nosisforthesakeofconrmingchronicitycandelaycare.Manypeoplemaynotrecognizetheimportanceof
arepeatvisitiftreatmenthadnotbeeninitiated.Thus,
initiatingtreatmentbothallowsforearlierintervention
andalsoindicatestopeopletheimportanceofthe
disease.

---

### Chunk 21/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** methods | **Similarity:** 0.524

ey-international.orgS148KidneyInternational(2024)105(Suppl4S),S117–S314

SummaryofrecommendationstatementsandpracticepointsChapter1:EvaluationofCKD1.1DetectionandevaluationofCKD1.1.1DetectionofCKDPracticePoint1.1.1.1:Testpeopleatriskforandwithchronickidneydisease(CKD)usingbothurinealbuminmea-surementandassessmentofglomerularltrationrate(GFR).PracticePoint1.1.1.2:Followingincidentaldetectionofelevatedurinaryalbumin-to-creatinineratio(ACR),hematuria,orlowestimatedGFR(eGFR),repeatteststoconrmpresenceofCKD.1.1.2MethodsforstagingofCKDRecommendation1.1.2.1:InadultsatriskforCKD,werecommendusingcreatinine-basedestimatedglomer-ularltrationrate(eGFRcr).IfcystatinCisavailable,theGFRcategoryshouldbeestimatedfromthecombinationofcreatinineandcystatinC(creatinineandcystatinC–basedestimatedglomerularltrationrate[eGFRcr-cys])(1B).1.1.3EvaluationofchronicityPracticePoint1.1.3.1:Proofofchronicity(durationofaminimumof3months)canbeestablishedby:(i)reviewofpastmeasurements/estimationsofGFR;(ii)revie

---

### Chunk 22/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.523

tNZJMed.1992;22:334–337.338.GilbertRE,AkdenizA,JerumsG.Detectionofmicroalbuminuriaindiabeticpatientsbyurinarydipstick.DiabetesResClinPract.1997;35:57–60.339.GrazianiMS,GambaroG,MantovaniL,etal.Diagnosticaccuracyofa
reagentstripforassessingurinaryalbuminexcretioninthegeneralpopulation.NephrolDialTransplant.2008;24:1490–1494.340.GuyM,NewallR,BorzomatoJ,etal.Useofarst-lineurineprotein-to-creatinineratiostriptestonrandomurinestoruleoutproteinuriain
patientswithchronickidneydisease.NephrolDialTransplant.2008;24:1189–1193.341.GuyM,NewallR,BorzomatoJ,etal.Diagnosticaccuracyoftheurinary
albumin:creatinineratiodeterminedbytheCLINITEKMicroalbuminandDCA2000þfortherule-outofalbuminuriainchronickidneydisease.ClinChimActa.2008;399:54–58.342.HasslacherC,MullerP,SchlipfenbacherRL.ResultsofamulticentrestudyforthedeterminationofmicroalbuminuriawithMicral-Test.KlinischesLabor.1995;41:441–447.343.HodelNC,HamadA,ReitherK,etal.Comparisonoftwodifferent
semiquantitativeurinarydipsticktestswithal

---

### Chunk 23/30
**Article:** Cardiologia VIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.523

"marca individual" e sua análise deve ser comparativa com o histórico do próprio paciente. Um aumento, mesmo dentro da faixa de normalidade, pode indicar inflamação subclínica crônica, que está ligada a lesões vasculares e ao desenvolvimento de doenças cardiovasculares em todo o corpo.
### 2. Metabolismo de Um Carbono e Homocisteína
- **Importância da Vitamina B12:** A deficiência é prevalente, afetando cerca de 20% da população com polimorfismos genéticos no transporte de B12 e 20% dos idosos (frequentemente por baixa acidez gástrica). O polimorfismo no gene FUT2 também pode reduzir sua absorção. O padrão-ouro para diagnóstico é o ácido metilmalónico (preferencialmente na urina). Níveis elevados de B12 sérica podem ser falsos, causados pelo consumo de algas ou levedura nutricional.
- **Ciclo da Homocisteína e Doadores de Metil:** A homocisteína deve ser mantida idealmente entre 4, 5 e 8 mg/dL.

---

### Chunk 24/30
**Article:** Aula 01 Guilherme Sorrentino - Cirurgia e Medicina Funcional Integrativa (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.523

teínas críticas da cicatrização (fibroblastos, colágeno, células imunológicas), promovendo excesso de oxidação e complicações como vasculopatia, retinopatia, neuropatia e nefropatia. Tais achados reforçam a necessidade de avaliação personalizada, com seleção de exames e intervenções conforme o histórico e os achados iniciais de cada paciente.

---

### Chunk 25/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.522

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

### Chunk 26/30
**Article:** Hipertensão Arterial Sistêmica (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.521

nal (sistema renina-angiotensina-aldosterona - SRAA), renal, microbiota e endotélio.
    *   **Mecanismos de Controle:** Rápidos (neurais), médio prazo (hormonais, alvo principal dos fármacos) e longo prazo (controle de volemia pelo rim).
*   **Diagnóstico e Classificação:**
    *   A medição em consultório é criticada; recomenda-se MAPA (Monitorização Ambulatorial) ou MRPA (Monitorização Residencial) para um diagnóstico preciso.
    *   **Valores de Referência para Diagnóstico:** ≥ 140/90 mmHg (consultório), ≥ 130/80 mmHg (MAPA 24h), ≥ 135/85 mmHg (MRPA).
    *   **Classificação (a partir de 18 anos):**
        *   **Ótima:** < 120/80 mmHg.
        *   **Normal:** 120-129 / 80-84 mmHg.
        *   **Pré-hipertensão:** 130-139 / 85-89 mmHg.
        *   **Hipertensão Estágio 1:** 140-159 / 90-99 mmHg.
        *   **Hipertensão Estágio 2:** 160-179 / 100-109 mmHg.
        *   **Hipertensão Estágio 3:** ≥ 180/110 mmHg.
### 3.

---

### Chunk 27/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.521

signicance[MGRS])98Presenceofpersistenthematuriaoralbuminuriaiscriticalindetermining
differentialdiagnosisGenetictestingAPOL1,COL4A3,COL4A4,COL4A5,NPHS1,UMOD,
HNF1B,PKD1,PKD2Evolvingasatoolfordiagnosis,increasedutilizationisexpected.
Recognitionthatgeneticcausesaremorecommonandmaypresent
withoutclassicfamilyhistory99,100ANCA,antineutrophilcytoplasmicantibody;APOL1,apolipoprotein1;COL4A,typeIVcollagenalphachain;CT,computedtomography;GBM,glomerularbasementmembrane;HNF1B,hepatocytenuclearfactor1B;MRI,magneticresonanceimaging;NPHS1,congenitalnephroticsyndrome;PKD1,polycystickidneydisease-1;PKD2,polycystickidneydisease-2;PLA2R,M-typephospholipaseA2receptor;UMOD,uromodulin.

---

### Chunk 28/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.521

andthe“new”stagesbetterreecttheirriskassocia-tions.Forthatreason,whereavailable,cystatinCshouldbe
addedtocreatinineforthepurposeofestimatingGFRfor
CKDdiagnosisandstaging.1.1.3EvaluationofchronicityPracticePoint1.1.3.1:Proofofchronicity(durationofa
minimumof3months)canbeestablishedby:(i)reviewofpastmeasurements/estimationsofGFR;(ii)reviewofpastmeasurementsofalbuminuriaor
proteinuriaandurinemicroscopicexaminations;(iii)imagingndingssuchasreducedkidneysizeandreductionincorticalthickness;(iv)kidneypathologicalndingssuchasbrosisandatrophy;(v)medicalhistory,especiallyconditionsknownto
causeorcontributetoCKD;(vi)repeatmeasurementswithinandbeyondthe3-
monthpoint.PracticePoint1.1.3.2:Donotassumechronicitybased
uponasingleabnormallevelforeGFRandACR,asthendingcouldbetheresultofarecentacutekidneyinjury(AKI)eventoracutekidneydisease(AKD).PracticePoint1.1.3.3:ConsiderinitiationoftreatmentsforCKDatrstpresentationofdecreasedGFRorelevatedACRifCKDisdeemedlikelyduetopresenceofotherclinicalindic

---

### Chunk 29/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.521

lurinarytractdiseaseRecurrentkidneycalculiMultisystemdiseases/chronicinammatoryconditionsSystemiclupuserythematosus
Vasculitis
HIVIatrogenic(relatedtodrugtreatmentsandprocedures)Drug-inducednephrotoxicityandradiationnephritisFamilyhistoryorknowngeneticvariantassociatedwithCKDKidneyfailure,regardlessofidentiedcauseKidneydiseaserecognizedtobeassociatedwithgeneticabnormality(e.g.,PKD,APOL1-mediatedkidneydisease,andAlportsyndrome)GestationalconditionsPretermbirthSmallgestationalsizePre-eclampsia/eclampsiaOccupationalexposuresthatpromoteCKDriskCadmium,lead,andmercuryexposurePolycyclichydrocarbonsPesticidesAKD,acutekidneydisease;AKI,acutekidneyinjury;APOL1,apolipoproteinL1;CKD,chronickidneydisease;CKDu,chronickidneydiseaseofundeterminedorigin;PKD,polycystickidneydisease.

---

### Chunk 30/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.518

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

