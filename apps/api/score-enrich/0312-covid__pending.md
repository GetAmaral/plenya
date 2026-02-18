# ScoreItem: COVID

**ID:** `019bf31d-2ef0-734a-aef8-dfd8a07933e8`
**FullName:** COVID (Histórico de doenças - Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual): - Doenças infecciosas)

**Preparation Metadata:**
- Quality Grade: **GOOD**
- Total Chunks: 30 de 4 artigos
- Avg Similarity: 0.583

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-734a-aef8-dfd8a07933e8`.**

```json
{
  "score_item_id": "019bf31d-2ef0-734a-aef8-dfd8a07933e8",
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

**ScoreItem:** COVID (Histórico de doenças - Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual): - Doenças infecciosas)

**30 chunks de 4 artigos (avg similarity: 0.583)**

### Chunk 1/30
**Article:** MFI - SÍNDROME PÓS COVID - AULA 01 (1) (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.688

l por tecido (cérebro, articulações, intestino, estômago).
- Categorias de manifestações:
  - Neurológicas, renais, hepáticas, gastrointestinais, tromboembólicas, cardíacas, endócrinas, dermatológicas.
- Base inicial para todos os casos:
  - Foco em inflamação sistêmica e suporte mitocondrial.
  - Personalização adicional conforme achados clínicos e laboratoriais.

## Sintomas Comuns e Fatos Epidemiológicos
- Exemplos de sintomas de COVID longo:
  - Fadiga, cefaleia, desatenção, alopecia, dispneia, agueusia, anosmia, polipneia, dores articulares.
- Mais de 50 efeitos possíveis:
  - Necessidade de mapear o perfil individual; não padronizar tratamentos sem avaliação.

## Eixo Endócrino-Imune e Mecanismos Hormonais
- Interações esperadas entre resposta endócrina e imunológica:
  - Ativação de cascatas inflamatórias e disfunções de eixos hormonais.
- Três mecanismos principais de interação com o sistema endócrino:
  1.

---

### Chunk 2/30
**Article:** Long COVID: Complications, Underlying Mechanisms, and Treatment Strategies (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.645

can raise the likelihood of long Covid. However, subsequent studies suggest long Covid can occur regardless of prior comorbidities or severity of acute Covid-19.
Zadeh et al.Page 38
Arch Microbiol Immunol. Author manuscript; available in PMC 2023 June 29.
Author ManuscriptAuthor ManuscriptAuthor ManuscriptAuthor Manuscript
	
	
Getulio José Mattos Do Amaral Filho - getfilho@yahoo.com.br - CPF: 034.983.039-88

Figure 2: 
Covid-19 infection affects almost all organs and organ systems are affected resulting in different pathophysiology. Few of the key symptoms and outcome results are shown. This is primarily due to the sequelae of cytokine storm.
Zadeh et al.Page 39
Arch Microbiol Immunol. Author manuscript; available in PMC 2023 June 29.

---

### Chunk 3/30
**Article:** Long COVID: Complications, Underlying Mechanisms, and Treatment Strategies (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.640

plications for follow-up: results from a prospective UK cohort. Thorax 76 (2021): 399–401. [PubMed: 33273026] 
29. Alnefeesi Y, Siegel A, Lui LMW, Teopiz KM, Ho RCM, Lee Y, et al. Impact of SARS-CoV-2 Infection on Cognitive Function: A Systematic Review. Front Psychiatry 11 (2020): 621773. [PubMed: 33643083] 
30. Schultheiss C, Willscher E, Paschold L, Gottschick C, Klee B, Henkes SS, et al. The IL-1beta, IL-6, and TNF cytokine triad is associated with post-acute sequelae of Covid-19. Cell Rep Med 3 (2022): 100663. [PubMed: 35732153] 
31. VanderVeen BN, Fix DK, Montalvo RN, Counts BR, Smuder AJ, Murphy EA, et al. The regulation of skeletal muscle fatigability and mitochondrial function by chronically elevated interleukin-6. Exp Physiol 104 (2019): 385–97. [PubMed: 30576589] 
32. Motta-Santos D, Dos Santos RA, Oliveira M, Qadri F, Poglitsch M, Mosienko V, et al.

---

### Chunk 4/30
**Article:** Long COVID: Complications, Underlying Mechanisms, and Treatment Strategies (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.616

Metab 11 (2022): e0261. [PubMed: 35441129] 
25. Vollbracht C, Kraft K. Feasibility of Vitamin C in the Treatment of Post Viral Fatigue with Focus on Long Covid, Based on a Systematic Review of IV Vitamin C on Fatigue. Nutrients 13 (2021).
26. Santana K, Franca E, Sato J, Silva A, Queiroz M, de Farias J, et al. Non-invasive brain stimulation for fatigue in post-acute sequelae of SARS-CoV-2 (PASC). Brain Stimul 16 (2023): 100–7. [PubMed: 36693536] 
27. Khraisat B, Toubasi A, AlZoubi L, Al-Sayegh T, Mansour A. Meta-analysis of prevalence: the psychological sequelae among Covid-19 survivors. Int J Psychiatry Clin Pract 26 (2022): 234–43. [PubMed: 34709105] 
28. Arnold DT, Hamilton FW, Milne A, Morley AJ, Viner J, Attwood M, et al. Patient outcomes after hospitalisation with Covid-19 and implications for follow-up: results from a prospective UK cohort. Thorax 76 (2021): 399–401. [PubMed: 33273026] 
29. Alnefeesi Y, Siegel A, Lui LMW, Teopiz KM, Ho RCM, Lee Y, et al.

---

### Chunk 5/30
**Article:** Long COVID: Complications, Underlying Mechanisms, and Treatment Strategies (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.601

lable in PMC 2023 June 29.
Author ManuscriptAuthor ManuscriptAuthor ManuscriptAuthor Manuscript
	
	
Getulio José Mattos Do Amaral Filho - getfilho@yahoo.com.br - CPF: 034.983.039-88

References
1. Davis HE, Assaf GS, McCorkell L, Wei H, Low RJ, Re'em Y, et al. Characterizing long Covid in an international cohort: 7 months of symptoms and their impact. EClinicalMedicine 38 (2021): 101019. [PubMed: 34308300] 
2. Logue JK, Franko NM, McCulloch DJ, McDonald D, Magedson A, Wolf CR, et al. Sequelae in Adults at 6 Months After Covid -19 Infection. JAMA Netw Open 4 (2021): e210830. [PubMed: 33606031] 
3. Callard F. Epidemic Time: Thinking from the Sickbed. Bull Hist Med 94 (2020): 727–43. [PubMed: 33775950] 
4. Yong SJ, Liu S. Proposed subtypes of post-Covid-19 syndrome (or long- Covid) and their respective potential therapies. Rev Med Virol 32 (2022): e2315. [PubMed: 34888989] 
5. Wong MC, Huang J, Wong YY, Wong GL, Yip TC, Chan RN, et al.

---

### Chunk 6/30
**Article:** Long COVID: Complications, Underlying Mechanisms, and Treatment Strategies (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.597

are due to other contributing factors (such as pulmonary fibrosis, medications used during acute treatment, intensive care syndrome, or exacerbation of a currently existing medical condition) (
4
) (Figure 1). Finding a clear boundary between such processes is often impossible in real patient scenarios. However, in this article, we regard long Covid as a diagnosis of exclusion to be made only when the elicited symptoms of the disease cannot be explained by another diagnosis such as post-intensive care syndrome or exacerbation of pre-existing health conditions.
It has been shown that Covid-19 infection can result in long Covid symptoms, regardless of prior health status or severity of initial acute infection (
5
). Many studies have shown that Covid -19 can persist in certain body tissues long past acute infection.

---

### Chunk 7/30
**Article:** MFI - SÍNDROME PÓS COVID - AULA 01 (1) (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.597

rtisolismo sem corticoide prévio:
  - Suspeitar mimetismo anti-ACTH e dano adrenal via ACE2.
- Piora glicêmica pós-quadros gripais incomuns:
  - Considerar COVID subclínica; evitar medicalização precipitada; focar em mecanismos (cortisol, resistência à insulina).
- Tosse persistente:
  - Foco em nervo vago; quercetina fitossômica; ebastina quando histamina suspeita; exames de histamina/DAO; microfisioterapia/Miltapod.

## Reforços Didáticos do Curso Aplicados ao Pós-COVID
- Não há bases fisiológicas novas:
  - Mesmos princípios aplicados de forma conjugada.
- Casos clínicos multifatoriais:
  - Intestino + dor de cabeça; tireoide + “leaky gut”; ansiedade + disbiose; fadiga + HPA; neuroinflamação + excitotoxicidade.
- Combinação de ferramentas:
  - Do intestino às mitocôndrias, do eixo HPA à modulação do SNA.

## 📅 Próximos Passos
- [ ] Mapear sintomas e sistemas afetados em cada paciente (checklist multissistêmico).

---

### Chunk 8/30
**Article:** Long COVID: Complications, Underlying Mechanisms, and Treatment Strategies (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.591

, neutrophil extracellular traps (NET)s were found to be increased in plasma of severe acute respiratory syndrome coronavirus and Covid-19 patients (
9
, 
183
, 
184
, 
185
). The persistence of these factors in the tissue can contribute to symptoms of long Covid.
Treatment: 
Small studies have shown that use of complement inhibitors during acute Covid-19 can reduce Covid-19-associated complement activation (
186
). Such treatments can also be considered in the hypercoagulable states associated with long Covid.
2B. 
Microvascular and endothelial injury, and platelet activation: 
Severe endothelial damage in the vasculature of lung, heart, kidney, liver, and intestine of Covid-19 patients (
8
, 
96
, 
128
, 
129
). Such findings causally suggest disruption of ACE2 receptor downstream cascade in these vascular beds. These impaired microcirculatory findings are collectively known as Covid-19-endothelitis (
8
, 
96
, 
128
, 
187
).

---

### Chunk 9/30
**Article:** Long COVID: Complications, Underlying Mechanisms, and Treatment Strategies (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.590

such complications can persist while others can reoccur (Figure 2). Some general symptoms of long Covid include fatigue (not improved with rest), heart palpitations, shortness 
Zadeh et al.Page 2
Arch Microbiol Immunol. Author manuscript; available in PMC 2023 June 29.
Author ManuscriptAuthor ManuscriptAuthor ManuscriptAuthor Manuscript
	
	
Getulio José Mattos Do Amaral Filho - getfilho@yahoo.com.br - CPF: 034.983.039-88

of breadth, cough, anosmia, hyposmia, headache, cognitive dysfunction, parkinsonism, dementia delirium, “brain fog” (poor short-term memory, concentration, problem-solving, and executive function), mental fatigue, dizziness, vertigo, and anxiety/depression (Figure 3). Fatigue is the most reported symptom of long Covid irrespective of the severity of the infection (
11
, 
12
).

---

### Chunk 10/30
**Article:** Long COVID: Complications, Underlying Mechanisms, and Treatment Strategies (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.590

189
).
Zadeh et al.Page 22
Arch Microbiol Immunol. Author manuscript; available in PMC 2023 June 29.
Author ManuscriptAuthor ManuscriptAuthor ManuscriptAuthor Manuscript
	
	
Getulio José Mattos Do Amaral Filho - getfilho@yahoo.com.br - CPF: 034.983.039-88

Treatment: 
Use of antioxidant and cholesterol-lowering therapies, ACE inhibitors, and anti-TNF-α during acute phase of Covid-19 has been proposed asemans to stabilize endothelium (
9
, 
190
). Such considerations are especially important in comorbidities that can lead to endothelial dysfunction; i.e., smoking, hypertension, diabetes, obesity, male sex, and history of cardiovascular diseases (
128
).
3.
 
SARS-COV-2 specific and autoreactive immune responses:
 Covid-19 can evolve 
into long Covid through long-term activation of host immune response (
191
).

---

### Chunk 11/30
**Article:** Long COVID: Complications, Underlying Mechanisms, and Treatment Strategies (2026)
**Journal:** Revista não identificada
**Section:** introduction | **Similarity:** 0.590

José Mattos Do Amaral Filho - getfilho@yahoo.com.br - CPF: 034.983.039-88

1. Introduction
Following the emergence of the Covid-19 pandemic, the disease complications persisted in many post-Covid survivors long after acute infection (
1
, 
2
). This phenomenon is known as long Covid or post-Covid-19 and soon spawned a series of studies on the long-term symptoms following acute Covid. The term “long Covid” was first generated in early 2020 by patients who suffered from symptoms of Covid-19 for weeks to months (
3
). This condition later led to the initiation of many support groups and campaigns for greater understanding, recognition, and treatment of "long-term Covid”.

---

### Chunk 12/30
**Article:** Aula Afonso Salgado - Sistema nervoso autônomo (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.587

rtante da frequência cardíaca ao passar de deitado para em pé;  
     - **hipotensão neurogênica** – queda da pressão arterial ao assumir a posição ortostática;  
     - fadiga intensa, dor difusa, brain fog, distúrbios de sono, ansiedade, intolerância ao exercício, sintomas gastrointestinais, cefaleias, dor temporomandibular, dores articulares, fibromialgia.

   Ele destaca que:

   - valores de **SDNN ~ 40** (na soma de deitado + em pé) são referência para boa saúde;  
   - em muitos pacientes pós-COVID, encontra SDNN de **9–11**, o que indica prognóstico ruim;  
   - o COVID é, em essência, um estado de **desequilíbrio autonômico**;  
   - sequelas pós-COVID em crianças, mesmo em casos sem sintomas graves na fase aguda, associam-se a queixas de TDAH, memória, fadiga e comprometimento mitocondrial.

   Essa associação reforça a necessidade de incluir a avaliação da VFC como biomarcador central no manejo do long COVID.

8.

---

### Chunk 13/30
**Article:** MFI - SÍNDROME PÓS COVID - AULA 01 (1) (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.582

abólico: glicemia, hemoglobina glicada.
    - Tireóide: TSH, FT4, anticorpos tireoidianos.
    - Prolactina e macroprolactina.
    - IGF-1, quando pertinente.
    - Histamina/DAO em suspeita mastocitária.
  - Diagnósticos diferenciais e comorbidades (autoimunidade, neurodegeneração, distúrbios do SNA).
- Tratamento baseado em mecanismos:
  - Personalizar anti-inflamatórios, antioxidantes, suporte mitocondrial, moduladores do intestino, adaptógenos, hormônios quando necessário.
  - Integrar técnicas somáticas (p. ex., microfisioterapia) para desautonomia/tosse.
  - Monitorar marcadores e ajustar.

## Pontos Críticos e Exemplos Concretos
- Prolactina elevada pós-COVID:
  - Frequentemente benigna; conduta expectante com macroprolactina; intervir apenas se sintomática.
- Hipocortisolismo sem corticoide prévio:
  - Suspeitar mimetismo anti-ACTH e dano adrenal via ACE2.

---

### Chunk 14/30
**Article:** MFI - SÍNDROME PÓS COVID - AULA 01 (1) (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.577

para tratar pós-COVID:
  - Manejo requer abordagem funcional integrativa, caso a caso, baseada em clínica, exames e mecanismos.
  - O conteúdo do curso já fornece as bases para tratar pós-COVID ao integrar módulos de mitocôndria, intestino, hormônios, neuroinflamação, eixos neuroendócrinos, etc.
- Dois detalhes a enfatizar (nuances clínicas):
  - Desautonomia (disfunções do sistema nervoso autônomo).
  - Hiperativação mastocitária e intolerância à histamina (ponte com aula prévia do Dr. Cristiano Rude).

## Impacto Multissistêmico do SARS-CoV-2 e Bases Fisiopatológicas
- Tropismo multiorgânico via ACE2:
  - SNC, coração, trato gastrointestinal e cascata de coagulação.
- Inflamação persistente pós-infecção:
  - “Rastro inflamatório” com impacto variável por tecido (cérebro, articulações, intestino, estômago).

---

### Chunk 15/30
**Article:** Long COVID: Complications, Underlying Mechanisms, and Treatment Strategies (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.577

-Covid-19 syndrome (or long- Covid) and their respective potential therapies. Rev Med Virol 32 (2022): e2315. [PubMed: 34888989] 
5. Wong MC, Huang J, Wong YY, Wong GL, Yip TC, Chan RN, et al. Epidemiology, Symptomatology, and Risk Factors for Long Covid Symptoms: Population-Based, Multicenter Study. JMIR Public Health Surveill 9 (2023): e42315. [PubMed: 36645453] 
6. Davis G, Li K, Thankam FG, Wilson DR, Agrawal DK: Ocular Transmissibility of COVID-19: Possibilities and Perspectives. Mol Cell Biochem 477 (2022): 849–864. [PubMed: 35066705] 
7. Wais T, Hasan M, Rai V, Agrawal DK: Gut-brain communication in COVID-19: Potential biomarkers and interventional strategies. Expert Review of Clin Immunol 18 (2022): 947–960.
8. Thankam FG, Agrawal DK: Molecular chronicles of cytokine burst in COVID-19 patients with cardiovascular diseases. J Thoracic Cardiovasc Surg 161 (2021): e217–e226.
9.

---

### Chunk 16/30
**Article:** MFI - SÍNDROME PÓS COVID - AULA 02 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.573

tamina e Ativação Mastocitária

Para pacientes com sintomas persistentes, multissistêmicos e aparentemente inexplicáveis, uma hipótese diagnóstica fundamental é a **intolerância à histamina** ou a **síndrome de ativação mastocitária**, que podem ser exacerbadas pela infecção por COVID-19 ou pela vacinação.

**Mecanismos e Sintomas:**
*   A histamina é degradada por duas vias principais: a enzima **DAO (diamina oxidase)** e a **HNMT (histamina N-metiltransferase)**. Polimorfismos ou disfunções nessas enzimas podem levar ao acúmulo de histamina.
*   A condição de *leaky gut* (intestino permeável) potencializa os efeitos da histamina.
*   Os sintomas são variados devido à ampla distribuição de receptores de histamina (H1, H2, H3, H4) no corpo, podendo incluir:
    *   **Respiratórios:** Rinorreia, congestão nasal, dispneia.
    *   **Neurológicos:** Dores de cabeça, *brain fog*.
    *   **Cardíacos:** Taquicardia, palpitações.

---

### Chunk 17/30
**Article:** Long COVID: Complications, Underlying Mechanisms, and Treatment Strategies (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.571

32279441] 
78. Franke C, Ferse C, Kreye J, Reincke SM, Sanchez-Sendin E, Rocco A, et al. High frequency of cerebrospinal fluid autoantibodies in Covid-19 patients with neurological symptoms. Brain Behav Immun 93 (2021): 415–9. [PubMed: 33359380] 
79. Saniasiaya J, Narayanan P. Parosmia post Covid-19: an unpleasant manifestation of long Covidsyndrome. Postgrad Med J 98 (2022): e96. [PubMed: 35232851] 
Zadeh et al.Page 29
Arch Microbiol Immunol. Author manuscript; available in PMC 2023 June 29.
Author ManuscriptAuthor ManuscriptAuthor ManuscriptAuthor Manuscript
	
	
Getulio José Mattos Do Amaral Filho - getfilho@yahoo.com.br - CPF: 034.983.039-88

80. De Luca P, Camaioni A, Marra P, Salzano G, Carriere G, Ricciardi L, et al. Effect of Ultra-Micronized Palmitoylethanolamide and Luteolin on Olfaction and Memory in Patients with Long Covid: Results of a Longitudinal Study. Cells 11 (2022).
81. Bawazeer MA, Theoharides TC.

---

### Chunk 18/30
**Article:** MFI - SÍNDROME PÓS COVID - AULA 01 (1) (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.568

consequências. O vírus SARS-CoV-2, por sua afinidade com os receptores ACE2, pode afetar múltiplos órgãos (sistema nervoso central, coração, sistema gastrointestinal), e a inflamação resultante pode se manifestar de forma diferente em cada indivíduo, afetando predominantemente o cérebro, as articulações ou o intestino.

O tratamento inicial deve, portanto, focar no controle da inflamação e na correção do distúrbio mitocondrial. O palestrante detalha que sintomas como fadiga estão frequentemente ligados à disfunção do eixo HPA, enquanto sintomas neurológicos (depressão, ansiedade, pânico) derivam da neuroinflamação e excitotoxicidade glutamatérgica. Duas condições específicas são destacadas:
1.  **Desautonomia:** Uma disfunção do sistema nervoso autônomo, cada vez mais comum.
2.  **Hiperativação Mastocitária:** Uma liberação excessiva de histamina, levando a sintomas como tosse irritativa.

---

### Chunk 19/30
**Article:** Long COVID: Complications, Underlying Mechanisms, and Treatment Strategies (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.567

Long COVID: Complications, Underlying Mechanisms, and Treatment Strategies
Farigol Hakem Zadeh,
Daniel R. Wilson,
Devendra K. Agrawal
*
Department of Translational Research, College of Osteopathic Medicine of the Pacific Western University of Health Sciences, Pomona, California 91766, USA
Abstract
Long Covid is one of the most prevalent and puzzling conditions that arose with the Covid pandemic. Covid-19 infection generally resolves within several weeks but some experience new or lingering symptoms. Though there is no formal definition for such lingering symptoms the CDC boadly describes long Covid as persons having a wide range of new, recurring or sustained health issues four or more weeks after first being infected with SARS-CoV2. The WHO defines long Covid as the manifestation of symptoms from a “probable or confirmed” Covid-19 infection that start approximately 3 months after the onset of the acute infection and last for more than 2 months.

---

### Chunk 20/30
**Article:** Long COVID: Complications, Underlying Mechanisms, and Treatment Strategies (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.566

ost-viral syndrome post Covid-19. Med Hypotheses 144 (2020): 110055. [PubMed: 32758891] 
43. Parvizi J, Damasio AR. Neuroanatomical correlates of brainstem coma. Brain 126 (2003): 1524–36. [PubMed: 12805123] 
44. Zimniak M, Kirschner L, Hilpert H, Geiger N, Danov O, Oberwinkler H, et al. The serotonin reuptake inhibitor Fluoxetine inhibits SARS-CoV-2 in human lung tissue. Sci Rep 11 (2021): 5890. [PubMed: 33723270] 
45. Hoertel N, Sanchez-Rico M, Vernet R, Beeker N, Jannot AS, Neuraz A, et al. Association between antidepressant use and reduced risk of intubation or death in hospitalized patients with Covid-19: results from an observational study. Mol Psychiatry 26 (2021): 5199–212. [PubMed: 33536545] 
46. Davis HE, McCorkell L, Vogel JM, Topol EJ: Long COVID: major findings, mechanisms and recommendations. Nature Reviews Microbiology 21 (2023): 133–146. [PubMed: 36639608] 
47.

---

### Chunk 21/30
**Article:** Long COVID: Complications, Underlying Mechanisms, and Treatment Strategies (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.561

effects long Covid has on quality of life, future health and life expectancy are required to better understand and eventually prevent or treat the disease. We acknowledge the effects of long Covid are not limited to those in this article but as it may affect the health of future offspring and therefore, we deem it important to identify more prognostic and therapeutic targets to control this condition.
Keywords
Anxiety; Brain fog; Chronic fatigue syndrome; Cognitive dysfunction; Covid-19; Depression; Long Covid; Post Covid-19 syndrome; Post SARS-CoV2; Postural orthostatic tachycardia syndrome; Pulmonary fibrosis
This article is an open access article distributed under the terms and conditions of the Creative Commons Attribution (CC-BY) license 4.0
*
Corresponding author: Devendra K. Agrawal, Department of Translational Research, College of Osteopathic Medicine of the Pacific Western University of Health Sciences, Pomona, California 91766, USA.

---

### Chunk 22/30
**Article:** MFI - SÍNDROME PÓS COVID - AULA 01 (1) (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.561

o:
      - Metil-histamina urinária de 24 horas.
      - Atividade de DAO (diaminoxidase) sanguínea.

## Antivirais e Observações de Prática Clínica
- Ivermectina:
  - Padrão empírico adotado pelo docente; comparação com estabilização do uso de oseltamivir para H1N1.
  - Posologia sugerida: 1 comprimido de 1 mg por cada 30 kg de peso, por 5 dias, com refeição rica em gordura para melhor absorção.
  - Racional observado:
    - Diferença clínica percebida no pós-COVID entre pacientes que usaram e não usaram, correlacionada à replicação viral.
    - Sugestão: testar na prática e observar evolução do “pós”.
  - Nota: respeitar divergências e crenças clínicas; ponderar riscos/benefícios.
- Contexto de gestantes, autismo e medicamentos:
  - Cautela com exposições (p. ex., paracetamol) e programação metabólica fetal; considerar modulação inflamatória segura.

---

### Chunk 23/30
**Article:** MFI - SÍNDROME PÓS COVID - AULA 01 (1) (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.560

cada vez mais comum.
2.  **Hiperativação Mastocitária:** Uma liberação excessiva de histamina, levando a sintomas como tosse irritativa. Para esses casos, sugere-se quercetina em doses altas (pelo menos 500 mg/dia) e, em situações específicas, o uso temporário de antialérgicos (ex: ebastina 10mg duas vezes ao dia). Para confirmação diagnóstica, recomenda-se a dosagem de metil-histamina urinária ou da atividade da enzima DAO.
------------
## O Impacto Viral no Sistema Endócrino e Imunológico

A aula aprofunda a íntima relação entre as respostas imunológicas e endócrinas durante e após a infecção por COVID-19. A disfunção hormonal ocorre por três mecanismos principais:
1.  **Infecção Viral Direta:** O vírus pode infectar glândulas como a pituitária e a adrenal através dos receptores ACE2, causando dano celular (edema, necrose) e hipofisite (inflamação da hipófise).
2.

---

### Chunk 24/30
**Article:** MFI - SÍNDROME PÓS COVID - AULA 01 (1) (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.560

ndido ao longo do curso.
------------
## Plano de Ação

### Para: Palestrante
- [ ] Solicitar exame de metil-histamina urinária de 24 horas ou análise da atividade da DAO sanguínea em pacientes com suspeita de hiperativação mastocitária. - Sem prazo definido

---

## Teaching Note

## Módulo Final: Síndrome Pós-COVID – Visão Geral e Princípios de Manejo
- Condição nova, porém baseada em mecanismos já conhecidos:
  - Inflamação persistente como eixo central.
  - Dano mitocondrial e disfunção do eixo HPA (hipotálamo–pituitária–adrenal).
  - Disfunções hormonais e desautonomia (sistema nervoso autônomo).
  - Variabilidade individual elevada: não há uma única “síndrome”; há clusters de sintomas que variam por pessoa.
- Não existe “fórmula pronta” para tratar pós-COVID:
  - Manejo requer abordagem funcional integrativa, caso a caso, baseada em clínica, exames e mecanismos.

---

### Chunk 25/30
**Article:** Long COVID: Complications, Underlying Mechanisms, and Treatment Strategies (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.558

is frequent in patients with Covid-19. J Thromb Haemost 18 (2020): 2064–5. [PubMed: 32324958] 
116. Sollini M, Ciccarelli M, Cecconi M, Aghemo A, Morelli P, Gelardi F, et al. Vasculitis changes in Covid-19 survivors with persistent symptoms: an [(18)F]FDG-PET/CT study. Eur J Nucl Med Mol Imaging 48 (2021): 1460–6. [PubMed: 33123760] 
117. Wallukat G, Hohberger B, Wenzel K, Furst J, Schulze-Rothe S, Wallukat A, et al. Functional autoantibodies against G-protein coupled receptors in patients with persistent Long-Covid-19 symptoms. J Transl Autoimmun 4 (2021): 100100. [PubMed: 33880442] 
118. Sinagra G, Anzini M, Pereira NL, Bussani R, Finocchiaro G, Bartunek J, et al. Myocarditis in Clinical Practice. Mayo Clin Proc 91 (2016): 1256–66. [PubMed: 27489051] 
119. Shaw BH, Stiles LE, Bourne K, Green EA, Shibao CA, Okamoto LE, et al. The face of postural tachycardia syndrome - insights from a large cross-sectional online community-based survey. J Intern Med 286 (2019): 438–48.

---

### Chunk 26/30
**Article:** Long COVID: Complications, Underlying Mechanisms, and Treatment Strategies (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.556

1. Okada Y, Yoshimura K, Toya S, Tsuchimochi M. Pathogenesis of taste impairment and salivary dysfunction in Covid-19 patients. Jpn Dent Sci Rev 57 (2021): 111–22. [PubMed: 34257762] 
92. Mart MF, Ware LB. The long-lasting effects of the acute respiratory distress syndrome. Expert Rev Respir Med 14 (2020): 577–86. [PubMed: 32168460] 
93. Motiejunaite J, Balagny P, Arnoult F, Mangin L, Bancal C, d'Ortho MP, et al. Hyperventilation: A Possible Explanation for Long-Lasting Exercise Intolerance in Mild Covid-19 Survivors? Front Physiol 11 (2020): 614590. [PubMed: 33536937] 
94. Carsana L, Sonzogni A, Nasr A, Rossi RS, Pellegrinelli A, Zerbi P, et al. Pulmonary post-mortem findings in a series of Covid-19 cases from northern Italy: a two-centre descriptive study. Lancet Infect Dis 20 (2020): 1135–40. [PubMed: 32526193] 
95. McElvaney OJ, McEvoy NL, McElvaney OF, Carroll TP, Murphy MP, Dunlea DM, et al. Characterization of the Inflammatory Response to Severe Covid-19 Illness.

---

### Chunk 27/30
**Article:** Long COVID: Complications, Underlying Mechanisms, and Treatment Strategies (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.555

ng, and managing the severity of Covid-19 and long Covid.
Nonetheless, more studies of long Covid are needed, particularly to enhance prevention, promote recovery, optimize quality of life, and reduce future patient morbidity and mortality. Moreover, it is important to further investigate the role of long Covid in women’s health, pregnancy and the impact on future generations.
Funding:
The research work of DKA is supported by the research grants R01 HL144125 and R01HL147662 from the National Heart, Lung, and Blood Institute, National Institutes of Health, USA. The contents of this article are solely the responsibility of the authors and do not necessarily represent the official views of the National Institutes of Health.
Zadeh et al.Page 25
Arch Microbiol Immunol. Author manuscript; available in PMC 2023 June 29.

---

### Chunk 28/30
**Article:** Long COVID: Complications, Underlying Mechanisms, and Treatment Strategies (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.554

beta-1a plus remdesivir compared with remdesivir alone in hospitalised adults with Covid-19: a double-bind, randomised, placebo-controlled, phase 3 trial. Lancet Respir Med 9 (2021): 1365–76. [PubMed: 34672949] 
114. Consortium WHOST, Pan H, Peto R, Henao-Restrepo AM, Preziosi MP, Sathiyamoorthy V, et al. Repurposed Antiviral Drugs for Covid-19 - Interim WHO Solidarity Trial Results. N Engl J Med 384 (2021): 497–511. [PubMed: 33264556] 
Zadeh et al.Page 31
Arch Microbiol Immunol. Author manuscript; available in PMC 2023 June 29.
Author ManuscriptAuthor ManuscriptAuthor ManuscriptAuthor Manuscript
	
	
Getulio José Mattos Do Amaral Filho - getfilho@yahoo.com.br - CPF: 034.983.039-88

115. Harzallah I, Debliquis A, Drenou B. Lupus anticoagulant is frequent in patients with Covid-19. J Thromb Haemost 18 (2020): 2064–5. [PubMed: 32324958] 
116. Sollini M, Ciccarelli M, Cecconi M, Aghemo A, Morelli P, Gelardi F, et al.

---

### Chunk 29/30
**Article:** Long COVID: Complications, Underlying Mechanisms, and Treatment Strategies (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.553

exertional 
Zadeh et al.Page 11
Arch Microbiol Immunol. Author manuscript; available in PMC 2023 June 29.
Author ManuscriptAuthor ManuscriptAuthor ManuscriptAuthor Manuscript
	
	
Getulio José Mattos Do Amaral Filho - getfilho@yahoo.com.br - CPF: 034.983.039-88

hyperventilation, and breathlessness can partly originate from the autonomic dysregulations that impede ventilatory control. Currently, a clinical trial is looking at the prognostic of dysfunctional breathing in long Covid (
NCT05217875
).
III.
 
Immunothrombosis and elevated D-dimer: 
 While the immunothrombosis nature of 
acute severe Covid-19 is known, studies are investigating mechanisms underlying exertional breathlessness of long Covid. Studies have shown that while many coagulation parameters downtrend during Covid-19, D-dimer remains elevated in patients with symptoms of breathlessness with no signs of pulmonary embolism upon CT angiogram (
97
).

---

### Chunk 30/30
**Article:** MFI - SÍNDROME PÓS COVID - AULA 02 (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.552

ssores

Para combater a neuroinflamação e a excitotoxicidade glutamatérgica, sintomas centrais da síndrome pós-COVID, é recomendada uma abordagem combinada. O objetivo é modular a resposta inflamatória e reequilibrar os neurotransmissores, principalmente através da abertura de canais de GABA para reduzir a excitabilidade neuronal.

**Suplementação e Intervenções:**
*   **Ômega-3:** Uma meta-análise de ensaios clínicos randomizados demonstrou sua eficácia.
*   **Curcuminoides:** A dose sugerida é de pelo menos 500 mg por dia, associada a 5 mg de piperina.
*   **Magnésio:** Prescrição de magnésio treonato, glicina ou uma combinação de formas como treonato e dimalato. É possível prescrever 1 g de glicina separadamente.
*   **L-teanina:** Recomendada em doses de até 600 mg por dia.
*   **5-HTP:** A administração sublingual é uma opção.

---

