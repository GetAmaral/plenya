# ScoreItem: Delta Hidrogênio (Δ H₂)

**ID:** `c77cedd3-2800-7dac-93ef-95da8fbee038`
**FullName:** Delta Hidrogênio (Δ H₂) (Exames - Imagem)
**Unit:** ppm

**Preparation Metadata:**
- Quality Grade: **EXCELLENT**
- Total Chunks: 30 de 9 artigos
- Avg Similarity: 0.639

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `c77cedd3-2800-7dac-93ef-95da8fbee038`.**

```json
{
  "score_item_id": "c77cedd3-2800-7dac-93ef-95da8fbee038",
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

**ScoreItem:** Delta Hidrogênio (Δ H₂) (Exames - Imagem)
**Unidade:** ppm

**30 chunks de 9 artigos (avg similarity: 0.639)**

### Chunk 1/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.760

ratório para SIBO/IMO:** Considerado o método prático padrão, utilizando lactulose ou glicose.
    -   **Critério para SIBO (Hidrogênio):** Elevação acima de 20 ppm nos primeiros 90 minutos.
    -   **Critério para IMO (Metano):** Elevação acima de 10 ppm em qualquer momento do teste.
-   **Diagnóstico de SIFO:** O padrão ouro é o aspirado duodenal com cultura (>10³ UFC/ml de fungos), mas é raramente realizado. A avaliação da micobiota e metabólitos fúngicos pode ser útil.
-   **Outros Achados:** Ativação do eixo HPA e reação de "die-off" (com manifestações cutâneas) durante tratamento antifúngico.
## Diagnóstico Primário:
-   **Avaliação:** O texto é uma discussão abrangente sobre a Síndrome do Intestino Irritável (SII), suas causas subjacentes (aumento da permeabilidade intestinal, disbiose, inflamação) e comorbidades como distúrbios do sono, SIBO, IMO e SIFO.

---

### Chunk 2/30
**Article:** Understanding Our Tests: Hydrogen-Methane Breath Testing to Diagnose Small Intestinal Bacterial Overgrowth (2023)
**Journal:** Clinical and Translational Gastroenterology
**Section:** abstract | **Similarity:** 0.703

Revisão detalhada sobre testes respiratórios hidrogênio-metano para diagnóstico de SIBO/IMO. Introduz o conceito de supercrescimento de metanógenos intestinais (IMO) para distinguir padrões predominantes em metano do SIBO clássico. Múltiplos estudos identificaram que níveis elevados de metano estão positivamente associados à constipação, com o metano demonstrando inibir diretamente o trânsito intestinal em 59% em modelos animais. Um nível de metano ≥10 ppm em jejum ou em qualquer momento durante o teste define IMO positivo. Abordagem simplificada com medição única de metano em jejum (SMM) ≥10 ppm demonstrou alta performance diagnóstica comparável aos protocolos padrão de 2 horas.

---

### Chunk 3/30
**Article:** ACG Clinical Guideline: Small Intestinal Bacterial Overgrowth (2020)
**Journal:** American Journal of Gastroenterology
**Section:** abstract | **Similarity:** 0.699

This ACG clinical guideline provides evidence-based recommendations for diagnosing and managing SIBO. Glucose and lactulose breath tests are recommended diagnostic tools, with interpretation based on consensus criteria. A rise in hydrogen ≥20 ppm from baseline within 90 minutes (glucose) or 90-120 minutes (lactulose) is considered positive. The guideline emphasizes the importance of proper test preparation, including dietary restrictions and medication avoidance, to minimize false-positive and false-negative results.

---

### Chunk 4/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa III (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.688

baixo de 100 µg/g são altamente sugestivos de SII.
*   **Avaliação para Doença Celíaca**
    - É fundamental em todos os pacientes, não apenas naqueles com diarreia. Inclui a dosagem de IgA sérica total e o anticorpo antitransglutaminase IgA.
*   **Avaliação da Microbiota e Metabolômica**
    - A metabolômica (avaliação dos produtos da microbiota) é considerada mais importante que a análise da microbiota isolada. A análise de ácidos orgânicos urinários é uma ferramenta útil.
    - A dosagem de zonulina pode ser um bom marcador para o aumento da permeabilidade intestinal.
*   **Supercrescimento Bacteriano do Intestino Delgado (SIBO)**
    - Incidência 3,7 vezes maior em portadores de SII. O diagnóstico prático é feito pelo teste respiratório com lactulose ou glicose.
*   **Supercrescimento Metanogênico Intestinal (IMO)**
    - Supercrescimento de arqueias produtoras de metano, associado principalmente à constipação intestinal.

---

### Chunk 5/30
**Article:** Functional Disease, Dysbiosis, and Dyspepsia: How Helpful Is Rifaximin? (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.675

t meta-analysis showed the prevalence of SIBO in FD was 2.8 times higher than in con-trols [4]. The prevalence of SIBO is higher in patients with overlap compared to isolated FD and IBS. The conventional diagnosis of SIBO is based on jejunal aspirate culture, with colony count   105CFU/ml diagnostic of SIBO. Non-inva-sive tests for diagnosis of SIBO used primarily in clinical practice are the glucose hydrogen breath test (GHBT) and lactulose hydrogen breath test (LHBT) [5]. Though popular, their accuracy is doubtful, with a recent meta-analysis show-ing both tests to have low sensitivity with high speciﬁcity. Rifaximin is the one of ﬁrst-choice therapies in SIBO, where it acts as a poorly absorbed oral antibiotic that alters the composition of the gut microbiota. This eﬀect is also use-ful in patients with IBS and FD where disturbances in gut microbial composition are common.

---

### Chunk 6/30
**Article:** Reevaluating our understanding of lactulose breath tests by incorporating hydrogen sulfide measurements (2019)
**Journal:** JGH Open: An Open Access Journal of Gastroenterology and Hepatology
**Section:** results | **Similarity:** 0.673

-
rently,hydrogenandmethanegasconcentrationsaremeasuredinbreathtestingand
evaluatedagainstspeciccut-offvaluesforinterpretationasnormalorabnormal.However,microbialgaskineticsisacomplexprocessthatisnotcurrentlyfullycon-
sideredwheninterpretingbreathgasresults.Gasexchangebetweenhydrogenpro-
ducersandhydrogenconsumers(methanogensandsulfate-reducingbacteria)isa
processwherebyhydrogenavailabilityisdeterminedbybothitsproductionand
removal.Hydrogensuldeisacrucialgasinvolvedinthisprocessasitisamajorhydrogen-consumptivepathwayinvolvedinenergyexchange.
Methods:Thisisacross-sectionalstudyevaluatinglactulosebreathtestingwiththeinclusionofhydrogensuldemeasurementsinpatientsreferredforbreathtestingforgastrointestinalsymptomsofbloating,excessivegas,and/orabdominalpain.

---

### Chunk 7/30
**Article:** Reevaluating our understanding of lactulose breath tests by incorporating hydrogen sulfide measurements (2019)
**Journal:** JGH Open: An Open Access Journal of Gastroenterology and Hepatology
**Section:** results | **Similarity:** 0.667

etween
hydrogenproductionandconsumptionbymethanogenesisorsul-
fatereduction.Scintigraphiccecalarrivalcouldindicatethe
arrivaloftheheadoflactulosebolusintothececum.However,
breathhydrogenwouldonlyrisewhenhydrogenproductionhas
exceededthehydrogenconsumptiveprocesses.Thus,thetimeto
riseforbreathhydrogenwouldalwaysbelaterintimethanscin-
tigraphiccecalentry.InastudybyYuetal.,orocecalscintigraphywascom-paredtoLBTresults;theseauthorsfoundthat,inamajorityof
cases,timetoriseofbreathhydrogenoccurredaftercecalarrival
byscintigraphy.17Thisstudyconcludedthat,giventhetemporalrelationshipbetweenscintigraphyandbreathtesting,LBTwasnotreliableforthediagnosisofSIBO.17Thisdiscrepancycouldbeexplainedbytheworkofhydrogen-consumingmicrobes.As
hydrogengasisrapidlyusedupbyhydrogenconsumersin
methanogenesisorsulfatereduction,the“delayed”timetoriseofbreathhydrogen,whencomparedtoscintigraphiccecalarrival,
Figure4Graphicalrepresentationoflactulosebreathtestingresultsforpatientsconsidered“hydrogennonproducers 

---

### Chunk 8/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.656

m Estudo Low FODMAP: 12 meses (período de acompanhamento).
**Biomarcadores fecais e testes respiratórios orientam a diferenciação entre SII e doença orgânica, e a detecção de SIBO/SIFO.**
- Ponto de Corte Inferior de Calprotectina (fecal): 100 (ponto de corte inferior para avaliação de SII).
- Probabilidade de SII com Calprotectina <100: 98% (alta probabilidade de diagnóstico funcional).
- Faixa de Zona Cinzenta da Calprotectina: 100–250 (necessita interpretação cautelosa).
- Limite de Calprotectina para Colonoscopia: 250 (exige colonoscopia).
- Critério Diagnóstico de SIBO por Aspirado: 10^3 UFC/ml (critério diagnóstico via aspirado).
- Critério de Positividade no Teste Respiratório de Hidrogênio: >20 partes por milhão em até 90 minutos (positividade de H2).
- Critério de Positividade no Teste Respiratório de Metano: >10 partes por milhão em qualquer momento (positividade de CH4).

---

### Chunk 9/30
**Article:** Reevaluating our understanding of lactulose breath tests by incorporating hydrogen sulfide measurements (2019)
**Journal:** JGH Open: An Open Access Journal of Gastroenterology and Hepatology
**Section:** discussion | **Similarity:** 0.650

ig.Dis.Sci.1998;43:2080–5.16LevittMD.Productionandexcretionofhydrogengasinman.N.Engl.J.Med.1969;281:122–7.17YuD,CheesemanF,VannerS.Combinedoro-caecalscintigraphyandlactulosehydrogenbreathtestingdemonstratethatbreathtestingdetectsoro-caecaltransit,notsmallintestinalbacterialovergrowthinpatientswithIBS.Gut.2011;60:334–40.18KajsTM,FitzgeraldJA,BucknerRYetal.InuenceofamethanogenicoraonthebreathH2andsymptomresponsetoingestionofsorbitoloroatber.Am.J.Gastroenterol.1997;92:89–94.19NingY,LouC,HuangZetal.Clinicalvalueofradionuclidesmallintestinetransittimemeasurementcombinedwithlactulosehydrogenbreathtestforthediagnosisofbacterialovergrowthinirritablebowel
syndrome.Hell.J.Nucl.Med.2016;19:124–9.20KokuboT,MatsuiS,IshiguroM.Meta-analysisoforo-cecaltransittimeinfastingsubjects.Pharm.Res.2013;30:402–11.21LinHC,PratherC,FisherRSetal.Measurementofgastrointestinaltransit.Dig.Dis.Sci.2005;50:989–1004.22ScarpelliniE,AbenavoliL,BalsanoC,GabrielliM,LuzzaF,TackJ.Breathtestsfortheassessmento

---

### Chunk 10/30
**Article:** Methodology and indications of H2-breath testing in gastrointestinal diseases: the Rome Consensus Conference (2009)
**Journal:** Alimentary Pharmacology & Therapeutics
**Section:** abstract | **Similarity:** 0.640

The Rome Consensus Conference established standardized methodology for hydrogen breath testing in gastrointestinal diseases. Key recommendations include: 12-hour fasting, avoidance of antibiotics for 4 weeks, low-fermentable diet the day before, and baseline breath samples. The test is indicated for SIBO diagnosis, lactose intolerance, fructose malabsorption, and oro-cecal transit time assessment. Proper patient preparation and standardized interpretation criteria are essential for diagnostic accuracy.

---

### Chunk 11/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.638

intestinal.
-   **SIFO:** Sensação de "blurring" (efeito baiacu - sentir-se distendido mesmo sem distensão objetiva) e dor abdominal.
## Objetivo:
A palestra descreve os exames e achados objetivos para a avaliação de pacientes com suspeita de Síndrome do Intestino Irritável (SII), embora não contenha achados de exame físico de um paciente específico.
-   **Exames Laboratoriais Gerais:** Hemograma e marcadores de atividade inflamatória.
-   **Calprotectina Fecal:** Usada para descartar doença inflamatória intestinal. Um valor abaixo de 100 µg/g tem uma positividade de 98% para o diagnóstico de SII. Entre 100-250 µg/g é uma zona cinzenta. Acima de 250 µg/g requer colonoscopia.
-   **Avaliação para Doença Celíaca:** Dosagem de IgA sérica total e anticorpo antitransglutaminase IgA. Se a IgA for baixa, dosar IgG sérica e antitransglutaminase IgG.
-   **Teste Parasitológico de Fezes:** Recomendado no contexto do Brasil.

---

### Chunk 12/30
**Article:** Hydrogen Sulfide and Methane on Breath Test Correlate with Human Small Intestinal Hydrogen Sulfide Producers and Methanogens (2025)
**Journal:** Digestive Diseases and Sciences
**Section:** other | **Similarity:** 0.637

studies identiﬁed SIBO by culturing small bowel aspirates [2, 3]. Subsequently, breath testing (BT) became an estab-lished indirect technique for diagnosing SIBO based on levels of hydrogen  (H2) on the breath [4]. Originally, BT onlymeasured  H2, but later testing also incorporated meas-urement of methane  (CH4). This allowed researchers to elu-cidate the importance of  CH4, which is produced in the gut 
 * Mark Pimentel  pimentelm@cshs.org1 Medically Associated Science andTechnology (MAST) Program, Cedars-Sinai, 770 N.

---

### Chunk 13/30
**Article:** Reevaluating our understanding of lactulose breath tests by incorporating hydrogen sulfide measurements (2019)
**Journal:** JGH Open: An Open Access Journal of Gastroenterology and Hepatology
**Section:** other | **Similarity:** 0.635

hedouble-peakphenom-enonwheretherearetwodistinctriseandfallpatternsofbreath
hydrogenconcentration,withtherstriserepresentingsmallbowelfermentationandthesecondriserepresentingcolonicfermenta-tion.29ThispatternhasbeenusedasacriterionfordiagnosingSIBO.4Ourmeanhydrogenproledoesshowapatternconsistentwithdouble-peak,withahigherSEMnotedatthetwopeaks
(Fig.5).Wefoundacontinuousriseofbreathhydrogenthroughout
theentiretestingperiod,suggestingthatthedouble-peakphenome-
noncouldbebetterexplainedonthebasisofadynamicprocess
wherebytheamountofhydrogenproducedbyfermentationinter-
mittentlyexceedsthehydrogen-consumptivecapacitytodrivea
spikeinhydrogenconcentration.Thus,breathhydrogenconcentra-
tionriseswhentheamountofhydrogenexceedsthehydrogen-
consumptivecapacitybutfallswhenthehydrogenproducediscon-
sumedastheamountofavailablehydrogendropsbelowthesatura-
tionpointforitsconversiontoeithermethaneorhydrogensulde.Inaddition,itisnotsurprisingthatithasbeenreportedthatthe
“smallbowel”and“largebowel  

---

### Chunk 14/30
**Article:** Hydrogen Sulfide and Methane on Breath Test Correlate with Human Small Intestinal Hydrogen Sulfide Producers and Methanogens (2025)
**Journal:** Digestive Diseases and Sciences
**Section:** results | **Similarity:** 0.632

s also include the small intestine. These results have important conse-quences. First, they reinforce the importance of breath testing and its interpretation. Excessive methane and meth-anogens are associated with constipation and the absence of excessive  CH4 allows the determination that higher  H2S producers mean looser stool. More importantly, identify-ing the organisms responsible for these gases, their precise gut locations, and their consequences for human health will help develop a new generation of approaches for treat-ing SIBO, IMO and sulﬁde overproduction.Supplementary Information The online version contains supplemen-tary material available at https:// doi. org/ 10. 1007/ s10620- 025- 09156-y.Acknowledgments The authors thank the REIMAGINE Study Group for their assistance in obtaining samples.

---

### Chunk 15/30
**Article:** Pros and Cons of Breath Testing for Small Intestinal Bacterial Overgrowth and Intestinal Methanogen Overgrowth (2023)
**Journal:** Gastroenterology & Hepatology
**Section:** abstract | **Similarity:** 0.631

This review examines diagnostic breath testing for SIBO and IMO, highlighting its advantages as a safe, noninvasive, widely accessible test that distinguishes hydrogen-predominant SIBO from IMO for tailored antibiotic therapy. Limitations include indirect measurement of microbial overgrowth, variability in orocecal transit time, risk of false-positive and false-negative results, and requirement for strict patient protocol compliance.

---

### Chunk 16/30
**Article:** Hydrogen and Methane-Based Breath Testing in Gastrointestinal Disorders: The North American Consensus (2017)
**Journal:** American Journal of Gastroenterology
**Section:** abstract | **Similarity:** 0.628

Este consenso norte-americano estabelece diretrizes para testes respiratórios com hidrogênio e metano em distúrbios gastrointestinais. Define que níveis de metano ≥10 ppm devem ser considerados positivos, estabelecendo a correlação entre metano e constipação. Pacientes com supercrescimento bacteriano predominante em metano têm 5 vezes mais probabilidade de apresentar constipação comparado aos casos predominantes em hidrogênio, com a gravidade correlacionando-se diretamente aos níveis de metano.

---

### Chunk 17/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.622

ério de Positividade no Teste Respiratório de Metano: >10 partes por milhão em qualquer momento (positividade de CH4).
- Prevalência de SIBO em Pacientes com SII: 78% (positividade em estudo de 2000).
- Risco Relativo de SIBO em SII: 3,7 vezes (incidência maior em SII).
- Critério Diagnóstico de SIFO: 10^3 UFC/ml (critério diagnóstico para SIFO).
**Terapias complementares e neuromodulação apoiam o controle de sintomas e comorbidades (sono, dor), especialmente quando há disbiose micótica ou hipersensibilidade.**
- Duração do Curso de Fluconazol para SIFO: duas a três semanas (curso antifúngico).
- Taxa de Resposta a Antifúngicos: 100% (resposta a fluconazol ou caspofungina).
- Posologia de Saccharomyces boulardii: 250 mg, duas vezes ao dia (adjuvante probiótico).
- Dose Inicial de Pregabalina: 200/50 miligramas (doses iniciais referidas).
- Tempo de Estimulação do Nervo Vago em Cada Orelha: 10 minutos (tempo na orelha direita e esquerda).

---

### Chunk 18/30
**Article:** Hydrogen Sulfide and Methane on Breath Test Correlate with Human Small Intestinal Hydrogen Sulfide Producers and Methanogens (2025)
**Journal:** Digestive Diseases and Sciences
**Section:** methods | **Similarity:** 0.621

mirabilis, D. oligo-trophicus, and D. widdelii. These data support the concept that breath gas proﬁles are inﬂuenced, at least in part, by small bowel microbial composition.Over the last half century, SIBO and BT have been increasingly understood. Despite the long history of BT as an indirect technique for assessing SIBO, and more recently IMO, the methodology to validate BT was not available until recently. An important question was whether BT ﬁndings were related to the levels of microorganisms producing these gases. The lack of data supported criticisms of BT, such as suggestions that BT, and particularly lactulose BT, is merely a marker of transit [36, 37]. Moreover, the measurement of breath  H2S, and the understanding that overgrowth of  H2S producers may contribute to patient symptoms [14], is still evolving. Here, we show that breath  CH4 and  H2S levels are in fact inﬂuenced by small bowel microbial composi-tion.

---

### Chunk 19/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa III (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.620

minase IgA. Se a IgA for baixa, dosar IgG sérica e antitransglutaminase IgG.
-   **Teste Parasitológico de Fezes:** Recomendado no contexto do Brasil.
-   **Avaliação da Permeabilidade Intestinal:** O aumento da permeabilidade (leaky gut) pode ser avaliado pela zonulina (fecal ou sérica). Menciona-se que o estresse (injeção de CRH) pode induzir um aumento nos marcadores de leaky gut.
-   **Avaliação da Microbiota/Metabolômica:** A avaliação isolada da microbiota é considerada de pouco valor. A avaliação da metabolômica (ex: ácidos orgânicos urinários) é mais útil para avaliar a função da microbiota e detectar metabólitos bacterianos e fúngicos. O aumento do D-lactato no sangue pode estar associado ao uso de probióticos e causar "brain fogginess".
-   **Teste Respiratório para SIBO/IMO:** Considerado o método prático padrão, utilizando lactulose ou glicose.
    -   **Critério para SIBO (Hidrogênio):** Elevação acima de 20 ppm nos primeiros 90 minutos.

---

### Chunk 20/30
**Article:** Hydrogen Sulfide and Methane on Breath Test Correlate with Human Small Intestinal Hydrogen Sulfide Producers and Methanogens (2025)
**Journal:** Digestive Diseases and Sciences
**Section:** discussion | **Similarity:** 0.620

hensive review. Gastroenterol Hepatol (N Y). 2007;3:112122. 4. Rangan V, Nee J, Lembo AJ. Small intestinal bacterial overgrowth breath testing in gastroenterology: clinical utility and pitfalls. Clin Gastroenterol Hepatol. 2022;20:14501453. 5. Miller TL, Wolin MJ. Enumeration of Methanobrevibacter smithii in human feces. Arch Microbiol. 1982;131:1418.

3855Digestive Diseases and Sciences (2025) 70:3846–3856 
 6. Pimentel M, Lin HC, Enayati P etal. Methane, a gas produced by enteric bacteria, slows intestinal transit and augments small intes-tinal contractile activity. Am J Physiol Gastrointest Liver Physiol. 2006;290:G1089-1095. 7. Kunkel D, Basseri RJ, Makhani MD, Chong K, Chang C, Pimentel M. Methane on breath testing is associated with constipation: a systematic review and meta-analysis. Dig Dis Sci. 2011;56:16121618 https:// doi. org/ 10. 1007/ s10620- 011- 1590-5. 8. Mehravar S, Takakura W, Wang J, Pimentel M, Nasser J, Rezaie A.

---

### Chunk 21/30
**Article:** Reevaluating our understanding of lactulose breath tests by incorporating hydrogen sulfide measurements (2019)
**Journal:** JGH Open: An Open Access Journal of Gastroenterology and Hepatology
**Section:** conclusion | **Similarity:** 0.616

alyzedbetweenOctober2016andJune2017.Meanhydrogenconcentrationswithapositivetrendthrougha3-hperiod(R2=0.97),meanmethaneconcentrationswithapositivetrend(R2=0.69),andmeanhydrogensuldeconcentrationswithanegativetrend(R2=−0.71)wereobserved.Conclusion:Byincorporatingenergyexchangeintheinterpretationofthelactulosebreathtest,wereevaluatedspecicbreathgasproles,includingthosecommonlydescribedas“hydrogennonproducers”andthe“double-peak”phenomenon.IntroductionLactulosebreathtesting(LBT)hasbeenusedasadiagnostictoolforgastrointestinalconditionsinvolvingalteredmicrobialfer-
mentation,includingsmallintestinalbacterialovergrowth(SIBO)
andmaldigestion/malabsorptionsyndromes.1Despitebreathtest-ingbecomingwidelyused,thereislittleagreementontheinter-pretationofresults.2Levittandcolleaguesmeasuredhydrogen(H2),carbondioxide(CO2),methane(CH4),oxygen(O2),nitro-gen(N2),hydrogensulde(H2S),andammonia(NH3)inbreathandintestinalgassamplescollectedfromhealthysubjects.3Thesourceofbreathgasesthatareexclusi

---

### Chunk 22/30
**Article:** Reevaluating our understanding of lactulose breath tests by incorporating hydrogen sulfide measurements (2019)
**Journal:** JGH Open: An Open Access Journal of Gastroenterology and Hepatology
**Section:** results | **Similarity:** 0.613

dastheamountofavailablehydrogendropsbelowthesatura-
tionpointforitsconversiontoeithermethaneorhydrogensulde.Inaddition,itisnotsurprisingthatithasbeenreportedthatthe
“smallbowel”and“largebowel”breathhydrogenpeaksdidnotmatchscintigraphicradionuclidelocations.4ThisstudyanalyzedanalternativeinterpretationbasedonLBTwithconcurrenthydrogen,methane,andhydrogensuldeconcentrationresults,allavailablefromthesamepatient.Theideathatbreathhydrogenconcentrationmaydependontheinter-
actionofhydrogenproducersandconsumersprovidesanovel
conceptualframeworkforunderstandingsomeofthepuzzling
ndingsobservedduringalactulosebreathtestandinseveralpublishedstudiesinvolvingLBTandsimultaneousscintigraphy.

---

### Chunk 23/30
**Article:** Hydrogen Sulfide and Methane on Breath Test Correlate with Human Small Intestinal Hydrogen Sulfide Producers and Methanogens (2025)
**Journal:** Digestive Diseases and Sciences
**Section:** other | **Similarity:** 0.613

ces (2025) 70:3846–3856 
by methanogenic archaea[5], in delayed gastrointestinal (GI) transit and constipation, including constipation-predominant irritable bowel syndrome (IBS-C) [69]; increased intesti-nal colonization with methanogens is now known as intes-tinal methanogen overgrowth (IMO) [10, 11]. Neither  H2 nor  CH4 are produced by human cells, so measuring these gases provides exclusive indirect markers of gut microbial composition. Adopting BT for the diagnosis of SIBO has facilitated an ever-increasing understanding of the role of SIBO in conditions such as irritable bowel syndrome (IBS) [12], celiac disease [13], and others.Recently, a third gas, hydrogen sulﬁde  (H2S) has been introduced to BT [14].  H2S is a gasotransmitter that plays important roles in inﬂammation and mucosal repair in the GI tract [15] and has been linked to diarrhea-predominant irritable bowel syndrome (IBS-D) [16, 17].

---

### Chunk 24/30
**Article:** Reevaluating our understanding of lactulose breath tests by incorporating hydrogen sulfide measurements (2019)
**Journal:** JGH Open: An Open Access Journal of Gastroenterology and Hepatology
**Section:** discussion | **Similarity:** 0.611

roenterol.2015;31:130–6.27YaoCK,TuckCJ.Theclinicalvalueofbreathhydrogentesting.J.Gastroenterol.Hepatol.2017;32:20–2.28NewmanA.Breath-analysistestsingastroenterology.Gut.1974;15:308–23.29DiggoryRT,CuschieriA.Theeffectofdoseandosmolalityoflactu-loseontheoral-caecaltransittimedeterminedbythehydrogenbreathtestandthereproducibilityofthetestinnormalsubjects.Ann.Clin.Res.1985;17:331–3.ABirgetal.ReevaluatinglactulosebreathtestsJGHOpen:Anopenaccessjournalofgastroenterologyandhepatology3(2019)228–233©2019TheAuthors.JGHOpen:AnopenaccessjournalofgastroenterologyandhepatologypublishedbyJournalofGastroenterologyandHepatologyFoundationandJohnWiley&SonsAustralia,Ltd.233

---

### Chunk 25/30
**Article:** Hydrogen Sulfide and Methane on Breath Test Correlate with Human Small Intestinal Hydrogen Sulfide Producers and Methanogens (2025)
**Journal:** Digestive Diseases and Sciences
**Section:** discussion | **Similarity:** 0.609

symptoms [14], is still evolving. Here, we show that breath  CH4 and  H2S levels are in fact inﬂuenced by small bowel microbial composi-tion. Enhanced understanding of diﬀerent types of microbial overgrowth and potential competition for gas utilization will improve how we interpret breath testing.Methane is an important gas in breath testing. Substantial evidence now suggests that  CH4 produced in the gut alters gut neuromuscular function [6] and is associated with con-stipation [7
]. Further, constipation severity is proportional to 

3853Digestive Diseases and Sciences (2025) 70:3846–3856 
breath  CH4 levels [38] and stool M. smithii levels. In a rand-omized-controlled study, reducing  CH4 with a combination of antibiotics improved constipation [8, 39]. Recently, stool microbiome analysis in IBS-C subjects found direct corre-lations between breath  
CH4, constipation, and M. smithii levels [11].

---

### Chunk 26/30
**Article:** Reevaluating our understanding of lactulose breath tests by incorporating hydrogen sulfide measurements (2019)
**Journal:** JGH Open: An Open Access Journal of Gastroenterology and Hepatology
**Section:** other | **Similarity:** 0.608

gbac-
teriabutalsobymultiplehostdetoxicationmechanisms.Recordingmethanegasasthesolerouteofhydrogenconsump-
tiononLBTleadstoanincompleteinterpretationofthecomplex
interactionsinvolved.Wehopethatanappreciationandabetter
understandingofthisdynamicsystem,consideringhydrogenpro-
ductionaswellasmultiplepathwaysofhydrogenconsumptions,
willprovideresearcherswithamorecompleteapproachto
reviewinglactulosebreathtestsandshouldofferthenecessary
toolstocorrectlyinterpretlactulosebreathtestsinthesettingof
diseasessuchasSIBOandirritablebowelsyndrome.AcknowledgmentThisstudywassupported,inpart,bytheWinklerBacterialOver-
growthResearchFund.References1RezaieA,BuresiM,LemboAetal.Hydrogenandmethane-basedbreathtestingingastrointestinaldisorders:theNorthAmericanCon-
sensus.Am.J.Gastroenterol.2017;112:775–84.2DiStefanoM,MengoliC,BergonziMetal.Breathmethaneexcre-tionisnotanaccuratemarkerofcolonicmethaneproductioninirrita-
blebowelsyndrome.Am.J.Gastroenterol.2015;110:891–8.3NakamuraN,LinHC,McSweeneyCS,Macki

---

### Chunk 27/30
**Article:** Reevaluating our understanding of lactulose breath tests by incorporating hydrogen sulfide measurements (2019)
**Journal:** JGH Open: An Open Access Journal of Gastroenterology and Hepatology
**Section:** results | **Similarity:** 0.607

ereduction,withH2Sproductionbeingthemainhydrogen-consumptivepathway.6However,hydro-gensuldeconcentrationisnotroutinelymeasuredinbreathtests.Inthisstudy,wewilldescribethepotentialrelationshipbetweenconcentrationvaluesofhydrogenandthatofhydrogensuldeand/ormethaneonLBT.Wewilltestthehypothesisthatbreathtestinggasresultscouldbeinterpretedintermsofadynamicbalancebetweenhydrogenproductionbyfermentationandhydrogenconsumptionthroughthecompetingprocessesofmethanogenesisversussulfatereduction,sulfategaselimination,andreactionsaturation.MethodsPatientselection.ThisstudywasreviewedandapprovedbytheinstitutionalreviewboardoftheNewMexicoVAHealthCareSystem.BreathtestingdatafromOctober2016toJune
2017ofconsecutivepatientsreferredtotheGIlabattheVet-eransAffairsMedicalCenterinAlbuquerque,NM,USAwere
doi:10.1002/jgh3.12145228JGHOpen:Anopenaccessjournalofgastroenterologyandhepatology3(2019)228–233©2019TheAuthors.JGHOpen:AnopenaccessjournalofgastroenterologyandhepatologypublishedbyJournalofGastroenterolog

---

### Chunk 28/30
**Article:** Reevaluating our understanding of lactulose breath tests by incorporating hydrogen sulfide measurements (2019)
**Journal:** JGH Open: An Open Access Journal of Gastroenterology and Hepatology
**Section:** results | **Similarity:** 0.601

13;30:402–11.21LinHC,PratherC,FisherRSetal.Measurementofgastrointestinaltransit.Dig.Dis.Sci.2005;50:989–1004.22ScarpelliniE,AbenavoliL,BalsanoC,GabrielliM,LuzzaF,TackJ.Breathtestsfortheassessmentoftheorocecaltransittime.Eur.Rev.Med.Pharmacol.Sci.2013;17(Suppl.2):39–44.23LinEC,MasseyBT.Scintigraphydemonstrateshighrateoffalse-positiveresultsfromglucosebreathtestsforsmallbowelbacterialovergrowth.Clin.Gastroenterol.Hepatol.2016;14:203–8.24JandhyalaSM,TalukdarR,SubramanyamC,VuyyuruH,SasikalaM,NageshwarReddyD.Roleofthenormalgutmicrobiota.

---

### Chunk 29/30
**Article:** Reevaluating our understanding of lactulose breath tests by incorporating hydrogen sulfide measurements (2019)
**Journal:** JGH Open: An Open Access Journal of Gastroenterology and Hepatology
**Section:** discussion | **Similarity:** 0.601

ghthe
meanvaluesdidnotreach10ppmdifferencebetweenpeakand
trough,thereisasubstantialindividualpatientvariability.DiscussionTheimpactofhydrogenproducersandhydrogenconsumers(methanogensandsulfate-reducingbacteria)changingtheavailabilityofhydrogenintheexhaledbreathhasnotbeenade-quatelyconsideredorevaluated.Inturn,theinterpretationof
LBTresultshasbeenbasedonanincompletepicture.Ashydro-
genconsumersconverthydrogentomethaneandhydrogensul-
de,theamountofH2remainingandenteringthecirculationandappearingintheexhaledbreathdecreases.Ascommercially
availablegaschromatographsmeasureonlyH2andCH4,couldthecurrentapproachininterpretingbreathgasresultsbeawedwhenonlyapartofthegasexchangeisseen?
Figure2Meanmethane(CH4)concentrationinpartspermillion(ppm)overa3-hlactulosebreathtesting.R2of0.71.
Figure3Meanvaluesofhydrogensulde(H2S)gasinpartsperbil-lion(ppb)overa3-hlactulosebreathtesting.R2of0.69.

---

### Chunk 30/30
**Article:** Hydrogen Sulfide and Methane on Breath Test Correlate with Human Small Intestinal Hydrogen Sulfide Producers and Methanogens (2025)
**Journal:** Digestive Diseases and Sciences
**Section:** discussion | **Similarity:** 0.601

= 0.316, P = 0.001), Halar-chaeum sp. CBA1220 (R = 0.200, P = 0.037), and Candida-tus Nitrosotenuis cloacae (R = 0.201, P = 0.036). In IMO-negative subjects, P. mirabilis (R = 0.281, P = 0.021) and D. oligotrophicus (R = 0.353, P = 0.003) RAs correlated with looser stool.DiscussionHere, we demonstrate that  
CH4 levels on breath testing (BT) exhibit correlations with levels of methanogens in the small intestine. Importantly, the predominant methanogen M. smithii and other methanogens are present in the small intestine and appear to play additive roles in contributing to breath  CH4 levels, in addition to contributions from colonic methanogens. We also show that  H2S levels on BT exhibit correlations with levels of  
H2S-producing bacteria in the small intestine, including speciﬁcally P. mirabilis, D. oligo-trophicus, and D. widdelii.

---

