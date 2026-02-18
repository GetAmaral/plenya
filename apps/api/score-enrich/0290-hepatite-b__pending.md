# ScoreItem: Hepatite B

**ID:** `019bf31d-2ef0-7e76-b57a-c79c8e45a162`
**FullName:** Hepatite B (Histórico de doenças - Doenças crônicas (Questionar ativamente as doenças crônicas mais comuns, perguntar sobre duração, grau de controle e tratamentos utilizados no passado e atualmente) - Doenças virais crônicas)

**Preparation Metadata:**
- Quality Grade: **FAIR**
- Total Chunks: 30 de 10 artigos
- Avg Similarity: 0.445

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-7e76-b57a-c79c8e45a162`.**

```json
{
  "score_item_id": "019bf31d-2ef0-7e76-b57a-c79c8e45a162",
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

**ScoreItem:** Hepatite B (Histórico de doenças - Doenças crônicas (Questionar ativamente as doenças crônicas mais comuns, perguntar sobre duração, grau de controle e tratamentos utilizados no passado e atualmente) - Doenças virais crônicas)

**30 chunks de 10 artigos (avg similarity: 0.445)**

### Chunk 1/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** discussion | **Similarity:** 0.499

h 
JD, Marrero JA, Conjeevaram HS, et al. A simple 
noninvasive index can predict both significant fibrosis 
and cirrhosis in patients with chronic hepatitis C. 
Hepatology 2003;38:518-26.79. Mummadi RR, Petersen JR, Xiao SY, Snyder N. Role of 
simple biomarkers in predicting fibrosis progression in 
HCV infection. World J Gastroenterol 2010;16:5710-5.80. Poynard T, Ngo Y, Perazzo H, Munteanu M, Lebray 
P, Moussalli J, et al. Prognostic value of liver fibrosis 
biomarkers: a meta-analysis. Gastroenterol Hepatol (N 
Y) 2011;7:445-54.81. Okuda M, Li K, Beard MR, Showalter LA, Scholle F, 
Lemon SM, et al. Mitochondrial injury, oxidative stress, 
and antioxidant gene expression are induced by hepatitis 
C virus core protein. Gastroenterology 2002;122:366-
75.82. Nalpas B, Vassault A, Le Guillou A, Lesgourgues B, 
Ferry N, Lacour B, et al. Serum activity of mitochondrial 
aspartate aminotransferase: a sensitive marker of 
alcoholism with or without alcoholic hepatitis.

---

### Chunk 2/30
**Article:** Chronic hepatitis B in 2025: diagnosis, treatment and future directions (2025)
**Journal:** Clinical Medicine (London)
**Section:** abstract | **Similarity:** 0.487

Comprehensive 2025 review on hepatitis B diagnosis emphasizing anti-HBc role in serological pattern interpretation. Highlights common diagnostic pitfall of isolated anti-HBc positivity (typically reflects past exposure, requires management only if immunosuppression planned). Distinguishes anti-HBc IgM (acute infection) from IgG (chronic/resolved). Emphasizes reactivation risk in immunosuppressed anti-HBc-positive patients, particularly with B-cell depleting agents, warranting prophylactic therapy.

---

### Chunk 3/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** discussion | **Similarity:** 0.486

ase (MELD) 
scores.75,76 A raised AST/ALT ratio of only slightly above 1 (1.09) is predictive of the progression of chronic viral hepatitis 
C to cirrhosis.77It is therefore the elevation of AST, rather than ALT, which 
is predictive of fibrosis and other ratios involving AST, such 
as the AST to Platelet Ratio Index (APRI)78 and FIB4 index (which involves the four parameters: AST, ALT, platelets and 
age) are also more predictive.79,80 The reason why the AST is more elevated than ALT with progression of fibrosis is 
uncertain but may be either because of increased production, 
such as mitochondrial release,81,82 or a relatively reduced clearance.83Chronic viral hepatitis may also progress to hepatocellular 
carcinoma, however GGT is the best predictor of this 
complication, while AST is not predictive in multivariate 
analysis and ALT is not predictive at all.84 Alcoholic Hepatitis
The predominance of AST over ALT in alcohol-related liver 
disease was first reported by Harinasuta et a

---

### Chunk 4/30
**Article:** Hepatitis B Core Antibody Level: A Surrogate Marker for Host Antiviral Immunity in Chronic Hepatitis B Virus Infections (2023)
**Journal:** Viruses
**Section:** abstract | **Similarity:** 0.481

Demonstrates quantitative anti-HBc (qAnti-HBc) as universal anti-HBV immune surrogate reflecting host immune response. Patients in immune-active phases show ~10-fold higher qAnti-HBc versus immune-tolerant phases. Correlates positively with ALT activity and histological inflammation severity. Cut-off ~4.5 log10 IU/mL diagnoses moderate-to-severe inflammation in normal-ALT patients. Baseline qAnti-HBc >4.0-4.5 log10 IU/mL independently predicts HBeAg seroconversion with antivirals.

---

### Chunk 5/30
**Article:** Clinical Significance and Remaining Issues of Anti-HBc Antibody and HBV Core-Related Antigen (2024)
**Journal:** Diagnostics (Basel)
**Section:** abstract | **Similarity:** 0.478

Anti-HBc appears early in HBV infection and persists lifelong, crucial for blood safety screening and identifying past exposure. Major challenges: lack of standardization across assay systems (ELISA, CMIA, CLIA, CLEIA) producing results in different units (IU/mL, INH%, S/O), complicating interpretation. Difficulty distinguishing low-titer true positives from false positives with highly sensitive assays. Quantitative anti-HBc predicts treatment response, carcinogenesis risk, reactivation, and HBsAg clearance during therapy.

---

### Chunk 6/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** other | **Similarity:** 0.475

with chronic 
hepatitis C. Dig Dis Sci 1998;43:2156-9.61. Sheth SG, Flamm SL, Gordon FD, Chopra S. AST/ALT 
ratio predicts cirrhosis in patients with chronic hepatitis 
C virus infection. Am J Gastroenterol 1998;93:44-8.62. Anderson FH, Zeng L, Rock NR, Yoshida EM. An 
assessment of the clinical utility of serum ALT and AST 
in chronic hepatitis C. Hepatol Res 2000;18:63-71.63. Assy N, Minuk GY. Serum aspartate but not alanine 
aminotransferase levels help to predict the histological 
features of chronic hepatitis C viral infections in adults. Am J Gastroenterol 2000;95:1545-50.64. Park GJ, Lin BP, Ngu MC, Jones DB, Katelaris PH. 
Aspartate aminotransferase: alanine aminotransferase 
ratio in chronic hepatitis C infection: is it a useful 
predictor of cirrhosis? J Gastroenterol Hepatol 
2000;15:386-90.65. Giannini E, Risso D, Testa R. Transportability and 
reproducibility of the AST/ALT ratio in chronic 
hepatitis C patients. Am J Gastroenterol 2001;96:918-9.66.

---

### Chunk 7/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** other | **Similarity:** 0.465

ALT ratio in acute viral hepatitis survivors 
was 0.3-0.6, while in non survivors the AST/ALT ratio was 
1.2-2.356. The De Ritis ratio therefore reflects the time course of acute viral hepatitis and is generally a vital clue to the 
patient’s prognosis.57Chronic Viral Hepatitis
AST/ALT ratios below 1.0 are also typical of chronic viral 
hepatitis (e.g. hepatitis B and C), however ratios slightly 
above 1.0 may be found in chronic viral hepatitis but this 
is particularly when progression to fibrosis and cirrhosis is 

Clin Biochem Rev Vol 34 November 2013   121
present.58-66 In chronic hepatitis B patients without clinical evidence of cirrhosis, the presence of progressive fibrosis might be predicted using an AST/ALT ratio over 1.0 however 
the ratio does not go above 2.0 in any patient.67 In chronic hepatitis C the raised AST/ALT ratio similarly correlates 
with fibrosis rather than necroinflammatory activity (e.g.

---

### Chunk 8/30
**Article:** Occult Hepatitis B Virus Infection: An Update (2022)
**Journal:** Viruses
**Section:** abstract | **Similarity:** 0.464

Occult HBV infection (OBI): HBsAg-negative with replication-competent HBV DNA in liver. Seropositive OBI (~80%): anti-HBc/anti-HBs detectable; seronegative OBI (~20%): all markers absent. Anti-HBc serves as surrogate marker for OBI screening in blood donors, organ recipients, immunosuppressed patients. Reactivation occurs in up to 40% with immunosuppression (highest risk: anti-CD20 therapies >10%; moderate: chemotherapy/TNFα inhibitors 1-10%). OBI accelerates cirrhosis in concurrent liver disease, maintains oncogenic potential (HBV integration in 75% OBI-HCC). Prophylactic antivirals recommended for high-risk immunosuppression in anti-HBc-positive patients.

---

### Chunk 9/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** other | **Similarity:** 0.463

000;15:386-90.65. Giannini E, Risso D, Testa R. Transportability and 
reproducibility of the AST/ALT ratio in chronic 
hepatitis C patients. Am J Gastroenterol 2001;96:918-9.66. Pohl A, Behling C, Oliver D, Kilani M, Monson P, 
Hassanein T. Serum aminotransferase levels and 
platelet counts as predictors of degree of fibrosis in 
chronic hepatitis C virus infection. Am J Gastroenterol 
2001;96:3142-6.67. Park SY, Kang KH, Park JH, Lee JH, Cho CM, Tak 
WY, et al. [Clinical efficacy of AST/ALT ratio and 
platelet counts as predictors of degree of fibrosis in 
HBV infected patients without clinically evident liver 
cirrhosis]. Korean J Gastroenterol 2004;43:246-51.68. Ustündag Y, Bilezikçi B, Boyacioğlu S, Kayataş M, 
Odemir N. The utility of AST/ALT ratio as a non-
invasive demonstration of the degree of liver fibrosis 
in chronic HCV patients on long-term haemodialysis. 
Nephrol Dial Transplant 2000;15:1716-7.69. Mardini H, Record C.

---

### Chunk 10/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** results | **Similarity:** 0.454

of the degree of liver fibrosis 
in chronic HCV patients on long-term haemodialysis. 
Nephrol Dial Transplant 2000;15:1716-7.69. Mardini H, Record C. Detection assessment and 
monitoring of hepatic fibrosis: biochemistry or biopsy? 
Ann Clin Biochem 2005;42:441-7.70. Zarski JP, Sturm N, Guechot J, Paris A, Zafrani 
ES, Asselah T, et al; ANRS HCEP 23 Fibrostar 
Group. Comparison of nine blood tests and transient 
elastography for liver fibrosis in chronic hepatitis C: the 
ANRS HCEP-23 study. J Hepatol 2012;56:55-62.71. Crisan D, Radu C, Lupsor M, Sparchez Z, Grigorescu 
MD, Grigorescu M. Two or more synchronous 
combination of noninvasive tests to increase accuracy of 
liver fibrosis assessement in chronic hepatitis C; results 
from a cohort of 446 patients. Hepat Mon 2012;12:177-
84.72. Stevenson M, Lloyd-Jones M, Morgan MY, Wong 
R.

---

### Chunk 11/30
**Article:** Emagrecimento - Parte V (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.454

dadeiros fatores de confusão.
### 3. Fisiopatologia e Abordagem da Esteatose Hepática Gordurosa Não Alcoólica (EHGNA/DHGNA)
- A EHGNA é a principal causa de cirrose atualmente, superando as causas alcoólicas, e está associada ao sobrepeso e alimentação inadequada.
- **Fisiopatologia:** Aumento da ingesta calórica -> aumento da insulina e resistência periférica -> aumento da lipogênese -> acúmulo de gordura no fígado, podendo evoluir para inflamação (esteato-hepatite) e cirrose.
- **Marcadores:** A ferritina pode estar elevada em 20-50% dos casos, enquanto TGO e TGP não são fidedignos (normais em até 75% dos casos).
- O diagnóstico (ex: ultrassonografia) pode servir como um "susto bom" para motivar o paciente a mudar o estilo de vida.
### 4. Impacto da Frutose e Intervenções no Estilo de Vida
- O consumo excessivo de frutose, principalmente via xarope de milho (HFCS) em industrializados, é um grande fator para a EHGNA.

---

### Chunk 12/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** other | **Similarity:** 0.453

Stevenson M, Lloyd-Jones M, Morgan MY, Wong 
R. Non-invasive diagnostic assessment tools for the 
detection of liver fibrosis in patients with suspected 
alcohol-related liver disease: a systematic review and 
economic evaluation. Health Technol Assess 2012;16:1-
174.73. Chou R, Wasson N. Blood tests to diagnose fibrosis 
or cirrhosis in patients with chronic hepatitis C virus 
infection: a systematic review. Ann Intern Med 
2013;158:807-20.74. Stránský J, Ryzlová M, Striteský J, Horák J. 
[Aspartate aminotransferase (AST) more than alanine 
aminotransferase (ALT) levels predict the progression 
of liver fibrosis in chronic HCV infection]. Vnitr Lek 
2002;48:924-8.75. Giannini E, Botta F, Testa E, Romagnoli P, Polegato 
S, Malfatti F, et al. The 1-year and 3-month prognostic 
utility of the AST/ALT ratio and model for end-stage 
liver disease score in patients with viral liver cirrhosis. 

Clin Biochem Rev Vol 34 November 2013   127
Am J Gastroenterol 2002;97:2855-60.76.

---

### Chunk 13/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** other | **Similarity:** 0.450

of the AST/ALT ratio and model for end-stage 
liver disease score in patients with viral liver cirrhosis. 

Clin Biochem Rev Vol 34 November 2013   127
Am J Gastroenterol 2002;97:2855-60.76. Giannini E, Risso D, Botta F, Chiarbonello B, Fasoli A, Malfatti F, et al. Validity and clinical utility of the 
aspartate aminotransferase-alanine aminotransferase 
ratio in assessing disease severity and prognosis in 
patients with hepatitis C virus-related chronic liver 
disease. Arch Intern Med 2003;163:218-24.77. Fortunato G, Castaldo G, Oriani G, Cerini R, Intrieri 
M, Molinaro E, et al. Multivariate discriminant function 
based on six biochemical markers in blood can predict 
the cirrhotic evolution of chronic hepatitis. Clin Chem 
2001;47:1696-700.78. Wai CT, Greenson JK, Fontana RJ, Kalbfleisch 
JD, Marrero JA, Conjeevaram HS, et al. A simple 
noninvasive index can predict both significant fibrosis 
and cirrhosis in patients with chronic hepatitis C. 
Hepatology 2003;38:518-26.79.

---

### Chunk 14/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** discussion | **Similarity:** 0.449

u A, Lesgourgues B, 
Ferry N, Lacour B, et al. Serum activity of mitochondrial 
aspartate aminotransferase: a sensitive marker of 
alcoholism with or without alcoholic hepatitis. 
Hepatology 1984;4:893-6.83. Kamimoto Y, Horiuchi S, Tanase S, Morino Y. 
Plasma clearance of intravenously injected aspartate 
aminotransferase isozymes: evidence for preferential 
uptake by sinusoidal liver cells. Hepatology 1985;5:367-
75.84. Hann HW, Wan S, Myers RE, Hann RS, Xing J, Chen 
B, et al. Comprehensive analysis of common serum liver 
enzymes as prospective predictors of hepatocellular 
carcinoma in HBV patients. PLoS One 2012;7:e47687.85. Harinasuta U, Chomet B, Ishak K, Zimmerman 
HJ. Steatonecrosis—Mallory body type. Medicine 
(Baltimore) 1967;46:141-62.86. Cohen JA, Kaplan MM. The SGOT/SGPT ratio—
an indicator of alcoholic liver disease. Dig Dis Sci 
1979;24:835-8.87. Correia JP, Alves PS, Camilo EA. SGOT-SGPT ratios. 
Dig Dis Sci 1981;26:284.88. Alves PS, Camilo EA, Correia JP.

---

### Chunk 15/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** other | **Similarity:** 0.449

.56. Gitlin N. The serum glutamic oxaloacetic transaminase/
serum glutamic pyruvic transaminase ratio as a 
prognostic index in severe acute viral hepatitis. Am J 
Gastroenterol 1982;77:2-4.57. Sunheimer R, Capaldo G, Kashanian F, Finck C, Woo 
J, Korins M, et al. Serum analyte pattern characteristic 
of fulminant hepatic failure. Ann Clin Lab Sci 
1994;24:101-9.58. Williams AL, Hoofnagle JH. Ratio of serum aspartate 
to alanine aminotransferase in chronic hepatitis; 
relationship to cirrhosis. Gastroenterology 1988;95:734-
9.59. Cadiot G, Ink O, Boutron A, Hanny P, Laurent-Puig P, 
Buffet C. Mitochondrial aspartate aminotransferase in 
nonalcoholic cirrhosis. Gastroenterology 1989;97:240-
1.60. Reedy DW, Loo AT, Levine RA. AST/ALT ratio > or = 
1 is not diagnostic of cirrhosis in patients with chronic 
hepatitis C. Dig Dis Sci 1998;43:2156-9.61. Sheth SG, Flamm SL, Gordon FD, Chopra S. AST/ALT 
ratio predicts cirrhosis in patients with chronic hepatitis 
C virus infection.

---

### Chunk 16/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** other | **Similarity:** 0.446

d 20-year risk of metabolic syndrome, diabetes, 
and cardiovascular disease. Gastroenterology 
2008;135:1935-44.152. Uslusoy HS, Nak SG, Gülten M, Bıyıklı Z. Non-
alcoholic steatohepatitis with normal aminotransferase 
values. World J Gastroenterol 2009;15:1863-8.153. Mardini H, Record C. Detection assessment and 
monitoring of hepatic fibrosis: biochemistry or biopsy? 
Ann Clin Biochem 2005;42:441-7.154. Dufour DR, Lott JA, Nolte FS, Gretch DR, Koff RS, 
Seeff LB. Diagnosis and monitoring of hepatic injury. 
II. Recommendations for use of laboratory tests in 
screening, diagnosis, and monitoring. Clin Chem 
2000;46:2050-68.155. Lin CS, Chang CS, Yang SS, Yeh HZ, Lin CW. 
Retrospective evaluation of serum markers APRI and 
AST/ALT for assessing liver fibrosis and cirrhosis in 
chronic hepatitis B and C patients with hepatocellular 
carcinoma. Intern Med 2008;47:569-75.156. Sorbi D, Boynton J, Lindor KD.

---

### Chunk 17/30
**Article:** Cardiologia VIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.431

isco:** Ferramentas como as tabelas de Framingham e MESA, embora matemáticas, são imprecisas por não considerarem uma vasta gama de variáveis metabólicas (ex: sono, hormonas, função mitocondrial, insulina, hemoglobina glicada).
- **Gama GT (Gama-glutamil transferase):** Além de ser um marcador de saúde para rins, pâncreas, fígado e estômago, a Gama GT (GGT) atua como um marcador de significância clínica para desfechos cardiovasculares. Níveis elevados podem indicar toxicidades crônicas (metais pesados, poluentes, defensivos agrícolas) e estão associados a maior risco cardiovascular e mortalidade geral. O objetivo terapêutico é manter o valor no quartil inferior da referência. O uso de simbióticos pode ajudar a melhorar a função hepática.
- **Leucócitos:** São uma "marca individual" e sua análise deve ser comparativa com o histórico do próprio paciente.

---

### Chunk 18/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** other | **Similarity:** 0.431

le 
injury.48Acute Viral Hepatitis
While it is clearly recognised that serum levels of both serum 
ALT and AST are elevated several fold in ‘acute’ viral hepatitis 
(e.g. hepatitis A and E), De Ritis was the first to describe that 
the ALT is usually higher than the AST with the AST/ALT 
ratio usually well below 1.0, and typically in the range 0.5 to 
0.7.49-54 It is important to note that the diagnosis of acute viral hepatitis today (when serology and PCR can clearly diagnose 
a viral cause) is very different to that question 60 years ago. 
Nevertheless, even today, when transaminases are elevated 
several fold, a low De Ritis ratio still provides an important 
clue to aetiology of acute hepatitis.Investigators have however noted that acute viral hepatitis can occasionally cause a high AST/ALT ratio of up to 2.0, 
however these are atypical cases representing fulminant 
hepatitis where there is often a very poor prognosis.

---

### Chunk 19/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** other | **Similarity:** 0.430

tate aminotransferase. Clin Chem 
1983;29:523-6.30. Vanderlinde RE. Review of pyridoxal phosphate and 
the transaminases in liver disease. Ann Clin Lab Sci 
1986;16:79-93.31. Férard G, Imbert-Bismut F, Messous D, Piton A, Abella 
A, Burnat P, et al. [Influence of pyridoxal phosphate in 
measuring aminotransferases activities in patients with 
viral hepatitis]. Ann Biol Clin (Paris) 2004;62:717-20.32. Ning M, Baker H, Leevy CM. Reduction of glutamic 
pyruvic transaminase in pyridoxine deficiency in liver 
disease. Proc Soc Exp Biol Med 1966;121:27-30.33. Matloff DS, Selinger MJ, Kaplan MM. Hepatic 
transaminase activity in alcoholic liver disease. 
Gastroenterology 1980;78:1389-92.34. Diehl AM, Potter J, Boitnott J, Van Duyn MA, Herlong HF, Mezey E. Relationship between pyridoxal 
5’-phosphate deficiency and aminotransferase levels in 
alcoholic hepatitis. Gastroenterology 1984;86:632-6.35. McIntyre N, Heathcote J. The laboratory in the diagnosis 
and management of viral hepatitis.

---

### Chunk 20/30
**Article:** Fisiologia e Bioquímica do Sistema Imune II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.427

olicitar e registrar parâmetros essenciais (PCR-us, VHS); calcular índices estimáveis (HGI, TAIG); considerar complementares (TNF-α, IL-6, GPx, MDA, antioxidantes totais).
- [ ] 11. Monitorar trimestralmente HbA1c, glicemia de jejum e HGI; estabelecer metas individuais (ex.: reduzir de 6,1 para ~5,5; longo prazo ~5,3).
- [ ] 12. Calcular TG/HDL e integrar com lipidograma/SREBP1c/2; evitar combinação de gordura saturada com açúcar e excesso de saturadas em múltiplas refeições.
- [ ] 13. Avaliar ferro, ferritina, transferrina e saturação (20–50%; evitar <20%); interpretar com hepcidina/SREBP1c e quadro inflamatório.
- [ ] 14. Medir TNF-α (<8,1) e IL-6 (<3,4) para acompanhar atividade inflamatória; relacionar com obesidade inflamada.
- [ ] 15. Calcular HOMA-β (167–175) e HOMA-IR (<2,15); buscar glicemia 60–90 e insulina ~6–7.
- [ ] 16. Monitorar homocisteína (<7,9) e PCR; usar PCR-us; documentar crises (PCR >1.000) e conduzir manejo apropriado.
- [ ] 17.

---

### Chunk 21/30
**Article:** Microbioma Intestinal III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.424

glutationa S-transferase) e melhorar o fluxo biliar.
## ❓ Perguntas
- [Inserir Pergunta/Confusão]
## 📚 Tarefas
- [ ] 1. Ao avaliar um paciente com qualquer condição crônica, priorizar a modulação do sistema gastrointestinal como parte fundamental do tratamento.
- [ ] 2. Na anamnese, investigar detalhadamente a história pregressa do paciente (parto, amamentação, uso de antibióticos, doenças, medicamentos).
- [ ] 3. Utilizar ferramentas clínicas como a Escala de Bristol e a observação de distensão abdominal para avaliar a saúde intestinal.
- [ ] 4. Considerar a solicitação de um exame coprológico funcional (como o Copromax) para uma avaliação aprofundada da inflamação e função intestinal.
- [ ] 5. Ao iniciar o uso do exame coprológico funcional, entrar em contato com o laboratório para obter suporte na interpretação dos resultados.
- [ ] 6.

---

### Chunk 22/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** results | **Similarity:** 0.422

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

### Chunk 23/30
**Article:** Emagrecimento - Parte V (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.419

abdominal é uma ferramenta útil para o diagnóstico e para aumentar a consciencialização do paciente.
    *   A ferritina pode ser um marcador de inflamação associada (elevada em 20-50% dos casos, principalmente em homens), enquanto TGO e TGP podem estar normais em até 75% dos pacientes.
*   **Papel do Açúcar e da Frutose:**
    *   O consumo excessivo de frutose, especialmente via xarope de milho rico em frutose (HFCS) em produtos industrializados, é um fator chave.
    *   Um estudo randomizado com adolescentes mostrou que a simples eliminação de açúcares livres (refrigerantes, sumos, etc.) reduziu a gordura no fígado de 25% para 17% em 8 semanas, com diminuição do TGP.
### 3. Estratégias de Tratamento para Esteatose Hepática
*   **Foco do Tratamento:** Combater o estresse oxidativo e a inflamação.
*   **Suplementos e Extratos:**
    *   **Vitamina E:** Combate o estresse oxidativo.

---

### Chunk 24/30
**Article:** Screening and Testing for Hepatitis B Virus Infection: CDC Recommendations (2023)
**Journal:** MMWR Recommendations and Reports
**Section:** abstract | **Similarity:** 0.417

CDC 2023 recommendations for universal HBV screening in adults aged ≥18 years and pregnancy screening.

---

### Chunk 25/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** other | **Similarity:** 0.417

hepatitis B and C patients with hepatocellular 
carcinoma. Intern Med 2008;47:569-75.156. Sorbi D, Boynton J, Lindor KD. The ratio of aspartate 
aminotransferase to alanine aminotransferase: potential 
value in differentiating nonalcoholic steatohepatitis 
from alcoholic liver disease. Am J Gastroenterol 
1999;94:1018-22.157. Neuschwander-Tetri BA, Clark JM, Bass NM, Van 
Natta ML, Unalp-Arida A, Tonascia J, et al; NASH 
Clinical Research Network. Clinical, laboratory and 
histological associations in adults with nonalcoholic 
fatty liver disease. Hepatology 2010;52:913-24.158. Brucknerová I, Benedeková M, Holomán K, Bieliková E, 
Kostrová A, Ujházy E, et al. Delivery as “physiological 
stress” and its influence on liver enzymatic systems 
in asphyxial newborns. Biomed Pap Med Fac Univ 
Palacky Olomouc Czech Repub 2005;149:409-11.159. Rosenthal P, Haight M. Aminotransferase as a 
prognostic index in infants with liver disease. Clin 
Chem 1990;36:346-8.160.

---

### Chunk 26/30
**Article:** Emagrecimento - Parte V (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.416

resumos, na comunicação social ou em interpretações de terceiros. Um exemplo citado é um estudo falacioso que associava pular o pequeno-almoço ao diabetes.
*   **Responsabilidade Individual:** A falta de cuidado com a própria saúde pode onerar financeira e emocionalmente os entes queridos. É importante abordar os pacientes de forma gradual, propondo pequenas mudanças para aumentar a adesão.
### 2. Esteatose Hepática Não Alcoólica (DHGNA)
*   **Fisiopatologia e Diagnóstico:**
    *   A DHGNA é gerada pelo "comer errado", levando ao aumento da insulina, resistência periférica, e consequente aumento da lipogénese (criação de gordura) no fígado.
    *   A sua evolução pode levar à cirrose, sendo atualmente a principal causa desta condição.
    *   A ultrassonografia abdominal é uma ferramenta útil para o diagnóstico e para aumentar a consciencialização do paciente.

---

### Chunk 27/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** other | **Similarity:** 0.416

them, 
which has continued to be useful for almost half a century, 
provides important clinical information which is worth the 
small additional cost.165Competing Interests: None declared.References1. De Ritis F, Coltorti M, Giusti G. An enzymic test for 
the diagnosis of viral hepatitis; the transaminase serum 
activities. Clin Chim Acta 1957;2:70-4.2. Rej R. Aspartate aminotransferase activity and 
isoenzyme proportions in human liver tissues. Clin 
Chem 1978;24:1971-9.3. King J. Practical Clinical Enzymology. London: D. Van 
Nostrand Company Limited; 1965.4. Kamimoto Y, Horiuchi S, Tanase S, Morino Y. 
Plasma clearance of intravenously injected aspartate 
aminotransferase isozymes: evidence for preferential 
uptake by sinusoidal liver cells. Hepatology 1985;5:367-
75.5. Feldstein AE, Canbay A, Angulo P, Taniai M, Burgart 
LJ, Lindor KD, et al. Hepatocyte apoptosis and fas 
expression are prominent features of human nonalcoholic 
steatohepatitis. Gastroenterology 2003;125:437-43.6.

---

### Chunk 28/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** results | **Similarity:** 0.415

l variations, and reference 
values. Clin Chem 1975;21:1077-87.140. Ruhl CE, Everhart JE. Determinants of the 
association of overweight with elevated serum 
alanine aminotransferase activity in the United States. 
Gastroenterology 2003;124:71-9.141. Ioannou GN, Boyko EJ, Lee SP. The prevalence and 
predictors of elevated serum aminotransferase activity 
in the United States in 1999-2002. Am J Gastroenterol 
2006;101:76-82.142. Zamin JI, de Mattos AA, Perin C, Ramos GZ. [The 
importance of AST / ALT rate in nonalcoholic 
steatohepatitis diagnosis]. Arq Gastroenterol 2002;39:22-
6.143. Nanji AA, French SW, Freeman JB. Serum alanine 
aminotransferase to aspartate aminotransferase ratio 
and degree of fatty liver in morbidly obese patients. 
Enzyme 1986;36:266-9.144. Pratt DS, Kaplan MM. Evaluation of abnormal liver-
enzyme results in asymptomatic patients. N Engl J Med 
2000;342:1266-71.145. Josekutty J, Iqbal J, Iwawaki T, Kohno K, Hussain 
MM.

---

### Chunk 29/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** other | **Similarity:** 0.413

c Univ 
Palacky Olomouc Czech Repub 2005;149:409-11.159. Rosenthal P, Haight M. Aminotransferase as a 
prognostic index in infants with liver disease. Clin 
Chem 1990;36:346-8.160. Prati D, Taioli E, Zanella A, Della Torre E, Butelli S, 
Del Vecchio E, et al. Updated definitions of healthy 
ranges for serum alanine aminotransferase levels. Ann 
Intern Med 2002;137:1-10.161. Parise ER, Oliveira AC, Figueiredo-Mendes C, Lanzoni 
V, Martins J, Nader H, et al. Noninvasive serum markers 
in the diagnosis of structural liver damage in chronic 

130   Clin Biochem Rev Vol 34 November 2013
hepatitis C virus infection. Liver Int 2006;26:1095-9.162. Mera JR, Dickson B, Feldman M. Influence of gender on the ratio of serum aspartate aminotransferase (AST) 
to alanine aminotransferase (ALT) in patients with and 
without hyperbilirubinemia. Dig Dis Sci 2008;53:799-
802.163. Lazo M, Selvin E, Clark JM.

---

### Chunk 30/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** other | **Similarity:** 0.409

hepatitis can have ALT greater than AST when several days have elapsed since alcohol exposure. The ratio 
therefore represents the time course and aggressiveness of disease that would be predicted from the relatively short half-life of 
AST (18 h) compared to ALT (36 h). In chronic viral illnesses such as chronic viral hepatitis and chronic alcoholism as well 
as non-alcoholic fatty liver disease, an elevated AST/ALT ratio is predictive of long terms complications including fibrosis 
and cirrhosis. There are methodological issues, particularly whether or not pyridoxal phosphate is used in the transaminase 
assays, and although this can have specific effects when patient samples are deficient in this vitamin, these method differences 
generally have mild effects on the usefulness of the assays or the ratio. Ideally laboratories should be using pyridoxal phosphate 
supplemented assays in alcoholic, elderly and cancer patients who may be pyridoxine deplete.

---

