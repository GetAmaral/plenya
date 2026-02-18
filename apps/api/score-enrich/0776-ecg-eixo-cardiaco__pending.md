# ScoreItem: ECG - Eixo Cardíaco

**ID:** `c77cedd3-2800-79b6-ae60-61aca75de8f8`
**FullName:** ECG - Eixo Cardíaco (Exames - Imagem)
**Unit:** graus

**Preparation Metadata:**
- Quality Grade: **POOR**
- Total Chunks: 30 de 21 artigos
- Avg Similarity: 0.390

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `c77cedd3-2800-79b6-ae60-61aca75de8f8`.**

```json
{
  "score_item_id": "c77cedd3-2800-79b6-ae60-61aca75de8f8",
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

**ScoreItem:** ECG - Eixo Cardíaco (Exames - Imagem)
**Unidade:** graus

**30 chunks de 21 artigos (avg similarity: 0.390)**

### Chunk 1/30
**Article:** Cardiologia VIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.488

# Cardiologia VIII

**Source:** https://web.plaud.ai/share/43b41764908850761::YXdzOnVzLXdlc3QtMg

---

## Reasoning Summary

## Análise de Exames, Fatores de Risco e Estratégias Terapêuticas em Cardiologia
### 1. Interpretação de Exames e Marcadores de Risco Cardiovascular
- **Princípio da Probabilidade e Contexto Clínico:** A análise de exames laboratoriais deve seguir o princípio de que a medicina é uma "ciência da probabilidade". Os valores devem ser interpretados dentro do contexto clínico do paciente (história, exame físico, idade) e não como números isolados a serem "corrigidos". É crucial não se fixar em atingir valores "ótimos" em todos os exames, pois o exame é um desfecho substituto e a prioridade é o paciente como um todo.

---

### Chunk 2/30
**Article:** Aula Afonso Salgado - Sistema nervoso autônomo (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.439

iação descrito inclui:

- Exame realizado preferencialmente pela manhã, em jejum, evitando a ingestão de medicamentos naquele momento (pois alteram a leitura).  
- Repetição do exame em **3 a 5 ocasiões** em condições semelhantes, para obter dados de “padrão ouro” (maior confiabilidade).  

A partir do ECG, softwares especializados analisam a VFC tanto no **domínio do tempo** quanto no **domínio da frequência**:

- No domínio do tempo, o parâmetro mais citado é o **SDNN** (desvio padrão dos intervalos NN), que é uma raiz quadrada aplicada à distribuição dos intervalos.  
- SDNN mais alto indica maior variabilidade; SDNN baixo indica rigidez do ritmo, associada a pior prognóstico.

No domínio da frequência, embora Afonso não detalhe numericamente, ele indica o uso de técnicas matemáticas como:

- **Rápida transformada de Fourier (FFT)**,  
- **wavelet transform**,  
- **ritmogramas** (conceito de origem russa).

---

### Chunk 3/30
**Article:** Aula Afonso Salgado - Sistema nervoso autônomo (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.436

## Avaliação Funcional e Diagnóstico via Variabilidade da Frequência Cardíaca (VFC)

No eixo diagnóstico, Afonso apresenta a **variabilidade da frequência cardíaca (VFC)** como o principal **biomarcador funcional** da integridade do SNA. A VFC é medida a partir de um eletrocardiograma simples e não invasivo, analisando-se os intervalos entre batimentos (intervalos NN). As variações naturais desses intervalos refletem a flexibilidade neurocardíaca.

Segundo a definição adotada pela Associação Americana de Cardiologia, a VFC é a **medida da função neurocardíaca** resultante da interação reflexa entre coração e cérebro, fornecendo dados dinâmicos do estado do SNA. Afonso resume:

- **Alta variabilidade** → alta atividade parassimpática, melhor resiliência, melhor prognóstico.  
- **Baixa variabilidade** → baixa atividade parassimpática, maior carga alostática, pior prognóstico.

Ele introduz dois conceitos centrais:

1.

---

### Chunk 4/30
**Article:** Echocardiographic Reference Ranges of Global Longitudinal Strain for All Cardiac Chambers Using Guideline-Directed Dedicated Views (2024)
**Journal:** JACC: Cardiovascular Imaging
**Section:** abstract | **Similarity:** 0.421

Estudo estabeleceu valores de referência para strain longitudinal global de todas câmaras cardíacas usando imagens obtidas conforme recomendações da ASE/EACVI. Incluiu 1.329 participantes saudáveis do estudo HUNT4Echo. Valores de referência para GLS do ventrículo esquerdo: -24% a -16%. Todos os strains associaram-se com idade, e GLS ventricular esquerdo associou-se com sexo. World Alliance Societies of Echocardiography propôs valores normais: -17% a -24% para homens e -18% a -26% para mulheres. Define-se GLS normal como ≤-18% e anormal como ≥-16%.

---

### Chunk 5/30
**Article:** Aula Afonso Salgado - Sistema nervoso autônomo (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.416

culados ao eixo hipotálamo–hipófise–adrenais.  
- **Trajeto neuroimune:** envolvendo macrófagos, múltiplas interleucinas e outros mediadores inflamatórios.

Ele enfatiza que há hoje grande volume de evidências (revisões sistemáticas e meta-análises) comprovando a relevância do SNA em diversas áreas: cardiologia, endocrinologia, imunologia, psiquiatria, neurologia, sono, nutrição, entre outras.

O SNA é entendido como um **exame biofísico**, porque sua avaliação se dá por meio da captação de sinais biológicos – sobretudo o eletrocardiograma (ECG). A partir dos intervalos entre batimentos cardíacos (intervalos NN), algoritmos matemáticos processam esses dados, resultando em parâmetros que permitem:

- interpretar o estado funcional do organismo;  
- distinguir **disfunção reversível** de **patologia instalada**;  
- comparar a importância diagnóstica do exame com a de exames clássicos, como o hemograma.

---

### Chunk 6/30
**Article:** Hipertensão Arterial Sistêmica II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.403

o de Diagnóstico:
1.  **Histórico Médico:** O conteúdo é uma palestra médica sobre fatores de risco cardiovascular contemporâneos, não uma consulta com um paciente específico.
2.  **Histórico de Medicação:** Inserir mais aqui
## Subjetivo:
O conteúdo é uma palestra médica e não contém queixas subjetivas de um paciente.
## Objetivo:
O conteúdo é uma palestra médica e não contém achados de exames de um paciente específico. O palestrante menciona seus próprios resultados de exames como exemplo:
-   **Índice de Ômega-3:** 6.7 (ideal entre 3 e 14).
-   **Relação Ômega-6 para Ômega-3:** 5:1 (ideal de 2:1 a 3:1), apesar da suplementação.
## Diagnóstico Primário:
-   **Avaliação:** O texto é uma apresentação educacional sobre fatores de risco inflamatórios e metabólicos para doenças vasculares, que são frequentemente negligenciados na cardiologia tradicional. Os principais fatores discutidos incluem:
    -   Desequilíbrio entre Ômega-3 e Ômega-6.

---

### Chunk 7/30
**Article:** Cardiologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** abstract | **Similarity:** 0.396

Conteúdo da aula: Cardiologia III...

---

### Chunk 8/30
**Article:** Aula Afonso Salgado - Sistema nervoso autônomo (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.393

bigeminia, trigeminia) etc.;  
  - também permite avaliar se o sistema reage positivamente a intervenções simples (como respiração profunda). Se não houver melhora com essas manobras, ele considera que não é o momento de prescrever exercícios respiratórios intensivos, sendo necessário começar por outras estratégias.

- **Card Check:**  
  - pode ser utilizado com sensor semelhante a oxímetro, inclusive em crianças;  
  - avalia seis funções fisiológicas principais:
    - oxigenação,  
    - ritmo cardíaco,  
    - flexibilidade vascular (índice de pulsatilidade),  
    - índice de resistividade do vaso,  
    - resistência temporal dos vasos e capacidade de reação (flexibilidade),  
    - reservas de energia nervosa e resposta a estressores;  
  - integra essas informações para estimar o estado psíquico–comportamental, via eixo coração–cérebro.

---

### Chunk 9/30
**Article:** Cardiologia VI (2024)
**Journal:** Pos Graduacao MFI
**Section:** abstract | **Similarity:** 0.390

Conteúdo da aula: Cardiologia VI...

---

### Chunk 10/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.390

dechocardiograms.........................................................................................................................................................................

---

### Chunk 11/30
**Article:** Sex-Specific Thresholds of High-Sensitivity Troponin in Patients With Suspected Acute Coronary Syndrome (2019)
**Journal:** Journal of the American College of Cardiology
**Section:** results | **Similarity:** 0.390

dline).Patientsaregroupedaccordingtowhethermyocardialinjurywaspresent(red)orabsent(gray).Pairedlog-ranktestresultsarep¼0.40forwomenwithmyocardialinjuryandp¼0.08forwomenwithoutmyocardialinjury.JACCVOL.74,NO.16,2019Leeetal.OCTOBER22,2019:2032–43Sex-SpecicThresholdsofhs-cTnI2039

infarctioninwomenandmen(15).Theimpactofsex-specicthresholdsonthediagnosisofmyocardialinfarctionhasbeenevaluatedinanumberofobser-vationalstudieswithdivergentndings(7,16–20).Mostofthesestudiesenrolledselectedpatientswith
acutecoronarysyndrome,ofwhomthemajoritywere
men.Furthermore,sex-specicthresholdswerenotusedtoguideclinicalcareorsubsequentinvestiga-
tionforcoronaryarterydisease.Here,weimple-
mentedsex-specicthresholdsintoroutineclinicalcareinarandomizedcontrolledtrialandevaluatedtheirimpactinconsecutivepatientspresentingwith
suspectedacutecoronarysyndrome.Wefoundthat
useofsex-specicthresholdsidentiedproportion-atelymorewomen,suchthattheoverallpercentages
ofwomenandmenidentiedashavingmyocardialinjury

---

### Chunk 12/30
**Article:** Cardiologia VII (2024)
**Journal:** Pos Graduacao MFI
**Section:** abstract | **Similarity:** 0.388

Conteúdo da aula: Cardiologia VII...

---

### Chunk 13/30
**Article:** Cardiologia IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** abstract | **Similarity:** 0.387

Conteúdo da aula: Cardiologia IV...

---

### Chunk 14/30
**Article:** Dislipidemias II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.385

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
**Article:** Cardiologia VIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** abstract | **Similarity:** 0.383

Conteúdo da aula: Cardiologia VIII...

---

### Chunk 16/30
**Article:** Ritmo Circadiano Eixo HPA - Parte I (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.382

das do cortisol (prometidos para aulas futuras)
7. Medição salivar de cortisol e melatonina (procedimentos e interpretação)
8. Módulo de saúde mental: teoria do cérebro trino e psiquiatria nutricional/metabólica
9. Modulação prática de parâmetros do ritmo circadiano (protocolos)
## Conteúdo Coberto
### 1. Introdução ao eixo HPA e sua centralidade clínica
- O eixo HPA regula e influencia todos os outros eixos hormonais sem exceção.
- Alta prevalência de sobrecarga do eixo de estresse em pacientes com queixas crônicas.
- Objetivo do módulo: compreender o eixo HPA em detalhes para transformar a prática clínica.
- Reconhecimento da heterogeneidade dos pacientes e importância de abordagem funcional integrativa.
> **Sugestões de IA**
> Você estabeleceu bem a relevância clínica; mantenha esse enquadramento. Para reforçar a retenção, você poderia abrir com um mapa conceitual visual do “panorama dos eixos” e onde o HPA se encaixa.

---

### Chunk 17/30
**Article:** Absolute Lymphocyte Count as a Surrogate Marker of CD4 Count in Monitoring HIV Infected Individuals: A Prospective Study (2016)
**Journal:** Journal of Clinical and Diagnostic Research
**Section:** other | **Similarity:** 0.382

blood in the pseudoaneurysm. ECG ﬁndings 
are non-speciﬁc [6]. X-ray chest is inconclusive as cardiomegaly 
is the sole ﬁnding. However detection of a para cardiac mass 
in a post myocardial infarct patient should alert us about the 
probability of a pseudoaneurysm. Diagnosis is possible by 
echocardiography but it is a highly operator dependent modality. 
Although echocardiography is the ﬁrst investigation performed 
in patients with cardiac murmur but it is sometimes limited by a 
poor echo window. Transesophageal echocardiography allows a 
much better visualization of the cardiac anatomy but has limited 
availability. CECT is the modality of choice as it offers a quick 
and easily available option for deﬁnitive diagnosis of this entity 
[7]. Catheter angiography is an invasive technique and runs the 
risk of causing rupture of the pseudoaneurysm or dislodging the 
thrombus.

---

### Chunk 18/30
**Article:** Cardiologia I (2024)
**Journal:** Pos Graduacao MFI
**Section:** abstract | **Similarity:** 0.381

Conteúdo da aula: Cardiologia I...

---

### Chunk 19/30
**Article:** Psiquiatria Metabólica Funcional Integrativa 22 - Frederico Porto (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.381

po em partes isoladas. A sua evolução reside na aplicação prática desta interconexão, que gera uma hierarquia terapêutica. A importância estratégica do modelo não é apenas reconhecer que os eixos estão ligados, mas usar essa ligação para priorizar intervenções. Ele sugere uma abordagem que começa com o físico e fundamental (nutrientes para o eixo cérebro-intestino) e progride para o mental e emocional (gestão do stress para o eixo cérebro-suprarrenal; respiração e coerência para o eixo cérebro-coração). Desta forma, o conceito evolui de um princípio filosófico ("não separar o corpo") para uma estratégia operacional clara que trata o paciente como um sistema unificado e interdependente.
**Trilha de Evidências:**
> Eu entendo que o eixo, a noção de eixo, ela te obriga a pensar integrativamente. Você é obrigado a não separar o corpo em pedaços. Então, eu enxergo três eixos.

---

### Chunk 20/30
**Article:** Aula Afonso Salgado - Sistema nervoso autônomo (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.377

formações para estimar o estado psíquico–comportamental, via eixo coração–cérebro.

- **Neurometria funcional (equipamento brasileiro):**  
  - aprovado pela Anvisa e com registro em órgãos internacionais (FDA mencionado);  
  - avalia VFC e múltiplos sistemas fisiológicos simultaneamente;  
  - é utilizado, segundo Afonso, para casos clínicos complexos, em que uma avaliação mais ampla do organismo é necessária.

Em termos de interpretação, Afonso exemplifica:

- Um traçado normal mostra grande variação entre batimentos (alta variabilidade);  
- Em pacientes com disfunções crônicas, o traçado mostra poucos picos e vales (baixa variabilidade), refletindo rigidez autonômica;  
- Em pós-COVID graves, SDNN abaixo de 20, ou mesmo na casa de 9–11, é um forte sinal de prognóstico ruim, enquanto SDNN em torno de 40 (soma de deitado + em pé) é considerado faixa saudável.

---

### Chunk 21/30
**Article:** Cardiologia V (2024)
**Journal:** Pos Graduacao MFI
**Section:** abstract | **Similarity:** 0.374

Conteúdo da aula: Cardiologia V...

---

### Chunk 22/30
**Article:** Left ventricular hypertrophy determined by Sokolow-Lyon criteria: a different predictor in women than in men? (2006)
**Journal:** Journal of Human Hypertension
**Section:** abstract | **Similarity:** 0.371

This prospective study examined 3,338 women and 3,330 men with hypertension over 11.2 years to assess whether ECG left ventricular hypertrophy (LVH) by Sokolow-Lyon voltage criteria predicted cardiovascular outcomes differently by gender. Increasing voltage independently predicted CVD mortality in both men and women. The risk of stroke, coronary heart disease (CHD) and cardiovascular disease (CVD) mortality increased significantly for each quantitative 0.1 mV increase in baseline ECG voltage, in women within the range of 1.6-3.9% and in men 1.4-3.0%. Women tended to have a high risk of stroke mortality owing to LVH, while men demonstrated stronger associations between voltage and coronary heart disease mortality.

---

### Chunk 23/30
**Article:** Cardiologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** abstract | **Similarity:** 0.371

Conteúdo da aula: Cardiologia II...

---

### Chunk 24/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.370

accausesofdyspnoea.ForrulinginheartfailureintheED,age-adjustedcut-pointshavebeenestablished:450pg/mlforpatientsunder50years,©2023EuropeanSocietyofCardiology.

---

### Chunk 25/30
**Article:** Predictive value of electrocardiographic left ventricular hypertrophy in the general population: A meta-analysis (2020)
**Journal:** Journal of Electrocardiology
**Section:** abstract | **Similarity:** 0.368

This meta-analysis examined electrocardiographic left ventricular hypertrophy (LVH) as a predictor of adverse cardiac outcomes in 58,400 participants across 10 studies. The Sokolow-Lyon voltage, Cornell voltage, and Cornell product criteria showed comparable ability in predicting major adverse cardiovascular events (MACE), with risk ratios ranging from 1.56 to 1.70. The pooled risk ratio of MACE was 1.62 (95% CI 1.40-1.89) for Sokolow-Lyon voltage criteria. The pooled risk ratio of all-cause mortality was 1.47 (95% CI 1.10-1.97), and cardiovascular mortality was 1.38 (95% CI 1.19-1.60) for Sokolow-Lyon criteria. Cornell voltage demonstrated stronger predictive value for cardiovascular and all-cause mortality outcomes.

---

### Chunk 26/30
**Article:** Guidelines for the Echocardiographic Assessment of the Right Heart in Adults: 2025 ASE Update (2025)
**Journal:** Journal of the American Society of Echocardiography
**Section:** abstract | **Similarity:** 0.368

Updated guidelines for right heart echocardiographic assessment introducing graded severity classification for TAPSE. Normal TAPSE values are ≥2.5 cm, with graded severity ranges allowing reporting as normal, mild, moderate, or severely reduced. Emphasizes multiparametric approach including TAPSE, S', FAC, 3D RVEF, and RV-PA coupling for comprehensive right ventricular function assessment.

---

### Chunk 27/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.368

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

### Chunk 28/30
**Article:** Cardiologia VIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.367

6 através de fitoquímicos (chás, shots) e jejum intermitente.
- [ ] 7. Abandonar a recomendação de consumo moderado de álcool, educando os pacientes sobre seus riscos metabólicos, genéticos e sobre a qualidade do sono.
- [ ] 8. Estudar e ter em mãos os estudos que embasam a abordagem funcional para argumentar contra dogmas médicos estabelecidos, encaminhando a outros profissionais quando necessário.
- [ ] 9. Ficar atento às aulas do Dr. Túlio Sperber, que complementarão o conteúdo deste módulo de cardiologia.

---

## Teaching Note

Data e Hora: 2025-11-20 20:42:21
Local: [Inserir Local]
Aula: [Inserir Nome da Aula]: Módulo de Cardiologia
## Visão Geral
A aula abordou a interpretação de exames laboratoriais e marcadores genéticos na cardiologia, enfatizando a individualização do tratamento em detrimento do foco exclusivo em valores de referência.

---

### Chunk 29/30
**Article:** Ritmo Circadiano Eixo HPA - Parte V (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.366

m paciente específico; conteúdo é descritivo e didático.
- Recomenda-se curva de cortisol salivar (idealmente domiciliar; laboratório especializado citado: Lemos, Juiz de Fora) para avaliação funcional do eixo HPA.
- Observações laboratoriais gerais:
  - ACTH frequentemente normal em disfunção do eixo; cortisol sérico matinal pode estar normal/alto por estresse da coleta; cortisol matinal muito baixo aumenta suspeita e indica curva salivar.
  - Em estresse militar de 5 dias: cortisol aumentou (ex.: ~542 para ~860 no 3º dia; ~550 para ~698 no 4º dia); testosterona total reduziu (~32 para ~5,3 nmol/L); testosterona livre reduziu (~127 para ~28); estradiol aumentou (~128 para ~158); DHEA reduziu (~27 para ~6). Ritmo circadiano permaneceu alterado após 5 dias de descanso.
- Curvas de cortisol didáticas:
  - Estresse agudo: cortisol elevado mantendo ritmo circadiano.
  - Fase adaptativa: pico matinal atenuado, vespertino/noturno elevados.

---

### Chunk 30/30
**Article:** EuropeanJournalofHeartFailure(2023)25,1891–1898POSITIONPAPERdoi:10.1002/ejhf.3036 (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.363

herhearthealthorheartstress.Hearthealthrefersto
individualswhohaveastructurallynormalheartandnormalplasma
concentrationsofNPsandtroponins.©2023EuropeanSocietyofCardiology.

---

