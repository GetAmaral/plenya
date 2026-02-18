# ScoreItem: Gama GT

**ID:** `019bf31d-2ef0-7465-91d2-cab643ae08d2`
**FullName:** Gama GT (Exames - Laboratoriais)
**Unit:** U/L

**Preparation Metadata:**
- Quality Grade: **EXCELLENT**
- Total Chunks: 30 de 10 artigos
- Avg Similarity: 0.622

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-7465-91d2-cab643ae08d2`.**

```json
{
  "score_item_id": "019bf31d-2ef0-7465-91d2-cab643ae08d2",
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

**ScoreItem:** Gama GT (Exames - Laboratoriais)
**Unidade:** U/L

**30 chunks de 10 artigos (avg similarity: 0.622)**

### Chunk 1/30
**Article:** Gamma-Glutamyltransferase: A Predictive Biomarker of Cellular Antioxidant Inadequacy and Disease Risk (2015)
**Journal:** Disease Markers
**Section:** abstract | **Similarity:** 0.752

Esta revisão examina a gama-glutamiltransferase (GGT) como marcador sérico superior para predizer múltiplas doenças crônicas e risco de mortalidade. Embora tradicionalmente reconhecida para identificar doença hepática relacionada ao álcool, evidências demonstram utilidade preditiva mais ampla da GGT em condições como doença cardiovascular, síndrome metabólica, diabetes tipo 2 e câncer. Estudos epidemiológicos mostram que níveis moderadamente elevados de GGT - frequentemente medidos anos antes do início da doença - predizem independentemente complicações metabólicas e cardiovasculares. O mecanismo envolve o papel da GGT na metabolização de xenobióticos glutatiolados, contribuindo para estresse oxidativo e nitrosativo. Criticamente, GGT elevada correlaciona-se inversamente com status antioxidante, particularmente depleção de glutationa. Aplicações clínicas incluem triagem de diabetes gestacional e monitoramento de progressão de doença renal crônica.

---

### Chunk 2/30
**Article:** Cardiologia VIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.745

isco:** Ferramentas como as tabelas de Framingham e MESA, embora matemáticas, são imprecisas por não considerarem uma vasta gama de variáveis metabólicas (ex: sono, hormonas, função mitocondrial, insulina, hemoglobina glicada).
- **Gama GT (Gama-glutamil transferase):** Além de ser um marcador de saúde para rins, pâncreas, fígado e estômago, a Gama GT (GGT) atua como um marcador de significância clínica para desfechos cardiovasculares. Níveis elevados podem indicar toxicidades crônicas (metais pesados, poluentes, defensivos agrícolas) e estão associados a maior risco cardiovascular e mortalidade geral. O objetivo terapêutico é manter o valor no quartil inferior da referência. O uso de simbióticos pode ajudar a melhorar a função hepática.
- **Leucócitos:** São uma "marca individual" e sua análise deve ser comparativa com o histórico do próprio paciente.

---

### Chunk 3/30
**Article:** Characteristics of peripheral blood Gamma-glutamyl transferase in different liver diseases (2022)
**Journal:** Medicine (Baltimore)
**Section:** abstract | **Similarity:** 0.730

Esta investigação examinou padrões de GGT em quatro condições hepáticas distintas para esclarecer seu significado diagnóstico além da avaliação de gravidade. A pesquisa envolveu 408 pacientes com cirrose biliar primária (CBP), lesão hepática induzida por drogas (DILI), doença hepática alcoólica (DHA) e doença hepática gordurosa não alcoólica (DHGNA). Achados revelaram diferenças marcantes em padrões de elevação da GGT. Em CBP e DILI, anormalidades da GGT ocorreram em aproximadamente 93-96% dos casos, correlacionando fortemente com marcadores de colestase como fosfatase alcalina. DHA demonstrou GGT elevada em 78% dos pacientes, refletindo tanto estresse oxidativo quanto colestase. DHGNA mostrou menor taxa de anormalidade (54%), com níveis máximos atingindo apenas 200 U/L comparados a mais de 2000 U/L na DHA. GGT correlacionou-se distintamente com diferentes biomarcadores em cada doença: conexões mais fortes com fosfatase alcalina e colesterol total em condições colestáticas, versus associações com triglicerídeos em distúrbios metabólicos.

---

### Chunk 4/30
**Article:** Usefulness of Gamma Glutamyl Transferase as Reliable Biological Marker in Objective Corroboration of Relapse in Alcohol Dependent Patients (2015)
**Journal:** Journal of Clinical and Diagnostic Research
**Section:** abstract | **Similarity:** 0.691

Este estudo prospectivo avaliou a gama-glutamiltransferase (GGT) como ferramenta diagnóstica para detectar recaída na dependência alcoólica em 52 pacientes das forças armadas ao longo de 12 meses. Os pesquisadores mediram níveis séricos de GGT na admissão e em intervalos de acompanhamento (3, 6, 9 e 12 meses), comparando resultados com avaliações psiquiátricas como padrão-ouro. A investigação determinou valores de corte ótimos usando análise ROC. Com 50 UI/L, a GGT demonstrou especificidade perfeita (100%) com sensibilidade variando de 56-100% nos diferentes momentos. A GGT exibiu diferenças estatisticamente significativas entre grupos com recaída e abstinentes em todos os intervalos. A meia-vida da enzima de 14-26 dias permite normalização em 4-5 semanas de abstinência, tornando-a adequada para monitorar progresso do tratamento. Taxa de sucesso terapêutico de 73-78% sugere que feedback objetivo com biomarcadores fortalece motivação do paciente para abstinência sustentada.

---

### Chunk 5/30
**Article:** Cardiologia VIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.684

dos biomarcadores de inflamação subclínica, como Gama GT e leucócitos, e estratégias de modulação genética (genes SIRT1/SIRT6) através de fitoquímicos e jejum. Por fim, o instrutor critica dogmas médicos, como a recomendação do consumo de álcool, e incentiva os profissionais a questionarem paradigmas estabelecidos com base em evidências atualizadas.
## 🔖 Pontos de Conhecimento
### 1. Interpretação de Exames e Abordagem Clínica
*   **Cuidado na Análise de Exames Laboratoriais:**
    *   Não se deve tentar enquadrar os exames em valores "ótimos" a todo custo, pois nem sempre é possível ou necessário. A medicina é a "ciência da probabilidade".
    *   Ferramentas como as tabelas de Framingham e MESA são imprecisas por não considerarem múltiplos fatores (sono, hormônios, função mitocondrial).
    *   O exame é um "desfecho substituto" e não deve sobrepor-se à avaliação do paciente como um todo.

---

### Chunk 6/30
**Article:** Cardiologia VIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.661

e metilcobalamina.
### 4. Marcadores Bioquímicos e Modulação Genética
- **Gama GT (GGT):** Quando elevado, pode indicar toxicidades crônicas e está associado a risco cardiovascular. O objetivo é mantê-lo no quartil inferior.
- **Leucócitos:** Um aumento no padrão individual pode indicar inflamação subclínica crônica, associada a lesão vascular.
- **Genes SIRT1 e SIRT6:** São importantes para a proteção cardiovascular. A má gestão de sua expressão pode levar a dano oxidativo e aterosclerose. Fitoquímicos (chás, shots) e o jejum intermitente são formas eficazes de modular positivamente esses genes.
### 5. Análise Crítica de Dogmas Médicos
- **Consumo de Álcool:** A recomendação de consumo moderado para saúde cardiovascular é problemática. O álcool interfere na metilação, seu metabólito (acetaldeído) é tóxico, e polimorfismos (ALDH2) podem intensificar o dano.

---

### Chunk 7/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** results | **Similarity:** 0.659

uation of abnormal liver-
enzyme results in asymptomatic patients. N Engl J Med 
2000;342:1266-71.145. Josekutty J, Iqbal J, Iwawaki T, Kohno K, Hussain 
MM. Microsomal triglyceride transfer protein inhibition 
induces endoplasmic reticulum stress and increases �gene transcription via Ire1α/cJun to enhance plasma 
ALT/AST. J Biol Chem 2013;288:14372-83.146. Feldstein AE, Wieckowska A, Lopez AR, Liu YC, 
Zein NN, McCullough AJ. Cytokeratin-18 fragment 
levels as noninvasive biomarkers for nonalcoholic 
steatohepatitis: a multicenter validation study. 
Hepatology 2009;50:1072-8.147. Kawamoto R, Kohara K, Kusunoki T, Tabara Y, 
Abe M, Miki T. Alanine aminotransferase/aspartate 
aminotransferase ratio is the best surrogate marker 
for insulin resistance in non-obese Japanese adults. 
Cardiovasc Diabetol 2012;11:117.148. Sookoian S, Pirola CJ. Alanine and aspartate 
aminotransferase and glutamine-cycling pathway: their 
roles in pathogenesis of metabolic syndrome.

---

### Chunk 8/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** other | **Similarity:** 0.648

io—
an indicator of alcoholic liver disease. Dig Dis Sci 
1979;24:835-8.87. Correia JP, Alves PS, Camilo EA. SGOT-SGPT ratios. 
Dig Dis Sci 1981;26:284.88. Alves PS, Camilo EA, Correia JP. The SGOT/
SGPT ratio in alcoholic liver disease. Acta Med Port 
1981;3:255-60.89. Salaspuro M. Use of enzymes for the diagnosis of 
alcohol-related organ damage. Enzyme 1987;37:87-
107.90. Takahashi A, Sekiya C, Yazaki Y, Ono M, Sato H, 
Hasebe C, et al. [Hepatic GOT and GPT activities in 
patients with various liver diseases—especially alcoholic 
liver disease]. Hokkaido Igaku Zasshi 1986;61:431-6.91. Sharpe PC. Biochemical detection and monitoring 
of alcohol abuse and abstinence. Ann Clin Biochem 
2001;38:652-64.92. Hietala J, Puukka K, Koivisto H, Anttila P, Niemelä 
O. Serum gamma-glutamyl transferase in alcoholics, 
moderate drinkers and abstainers: effect on gt reference 
intervals at population level. Alcohol Alcohol 
2005;40:511-4.93. Rosman AS, Lieber CS.

---

### Chunk 9/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** other | **Similarity:** 0.638

sferase in alcoholics, 
moderate drinkers and abstainers: effect on gt reference 
intervals at population level. Alcohol Alcohol 
2005;40:511-4.93. Rosman AS, Lieber CS. Diagnostic utility of laboratory 
tests in alcoholic liver disease. Clin Chem 1994;40:1641-
51.94. Matloff DS, Selinger MJ, Kaplan MM. Hepatic 
transaminase activity in alocholic liver disease. 
Gastroenterology 1980;78:1389-92.95. Diehl AM, Potter J, Boitnott J, Van Duyn MA, 
Herlong HF, Mezey E. Relationship between pyridoxal �5′-phosphate deficiency and aminotransferase levels in 
alcoholic hepatitis. Gastroenterology 1984;86:632-6.96. Nalpas B, Vassault A, Le Guillou A, Lesgourgues B, 
Ferry N, Lacour B, et al. Serum activity of mitochondrial 
aspartate aminotransferase: a sensitive marker of 
alcoholism with or without alcoholic hepatitis. 
Hepatology 1984;4:893-6.97. Rej R. Aspartate aminotransferase activity and 
isoenzyme proportions in human liver tissues. Clin 
Chem 1978;24:1971-9.98.

---

### Chunk 10/30
**Article:** Cardiologia VIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.631

  o padrão-ouro para diagnóstico. Níveis séricos podem ser falsamente elevados por algas ou levedura nutricional. O polimorfismo no gene FUT2 pode prejudicar sua absorção intestinal.
- **Homocisteína:** Seu aumento eleva a mortalidade por todas as causas, não apenas o risco cardiovascular, causando lesão endotelial e trombogênese. O valor ideal buscado é entre 4, 5 e 8. A elevação pode ser causada por deficiência de B12, folato, B6, colina ou por fatores como excesso de cafeína.
- **Folato e MTHFR:** O ácido fólico (sintético) é diferente do folato (natural). O polimorfismo no gene MTHFR é comum e está associado a níveis mais altos de homocisteína e maior risco de doenças. A suplementação deve ser feita com formas ativas como metilfolato, piridoxal-5-fosfato (P5P) e metilcobalamina.
### 4. Marcadores Bioquímicos e Modulação Genética
- **Gama GT (GGT):** Quando elevado, pode indicar toxicidades crônicas e está associado a risco cardiovascular.

---

### Chunk 11/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** results | **Similarity:** 0.626

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

### Chunk 12/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** other | **Similarity:** 0.615

logy 1984;4:893-6.101. Okuno F, Ishii H, Kashiwazaki K, Takagi S, Shigeta Y, 
Arai M, et al. Increase in mitochondrial GOT (m-GOT) 
activity after chronic alcohol consumption: clinical and 
experimental observations. Alcohol 1988;5:49-53.102. Hourigan KJ, Bowling FG. Alcoholic liver disease: 
a clinical series in an Australian private practice. J 
Gastroenterol Hepatol 2001;16:1138-43.103. Nyblom H, Berggren U, Balldin J, Olsson R. High 
AST/ALT ratio may indicate advanced alcoholic liver 
disease rather than heavy drinking. Alcohol Alcohol 
2004;39:336-9.104. Larsson A, Tryding N. Is it necessary to order aspartate 
aminotransferase with alanine aminotransferase in 
clinical practice? Clin Chem 2001;47:1133-5.

128   Clin Biochem Rev Vol 34 November 2013
105. Nyblom H, Berggren U, Balldin J, Olsson R. High AST/ALT ratio may indicate advanced alcoholic liver 
disease rather than heavy drinking. Alcohol Alcohol 
2004;39:336-9.106. Liangpunsakul S, Qi R, Crabb DW, Witzmann F.

---

### Chunk 13/30
**Article:** MFI - PROGRAMAÇÃO METABÓLICA - AULA 05 (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.610

masculino e estradiol elevado no cordão.
- Implicações: potenciais alterações neurocomportamentais e desenvolvimento sexual, com necessidade de atenção clínica diferenciada.
### 4. Gamma GT como marcador de toxicidade ambiental
- Frequentemente negligenciada em gestantes e pré-concepção; pode indicar exposição/toxicidade ambiental além de causas hepatobiliares.
- Interpretação contextual por quartis e ausência de causa aparente; nem sempre estará elevada.
- Observada associação com risco de câncer de mama em alguns contextos; considerar co-marcadores e histórico ocupacional.
### 5. Ferramentas (aplicativos) para avaliar toxinas em produtos
- Apps com leitura por código de barras úteis para produtos de beleza/infantis/gestantes; aplicabilidade maior para itens dos EUA.

---

### Chunk 14/30
**Article:** Dislipidemias II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.605

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
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** results | **Similarity:** 0.603

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

### Chunk 16/30
**Article:** Microbioma Intestinal III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.600

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

### Chunk 17/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** other | **Similarity:** 0.598

and value of 
determination of glutamic acid dehydrogenase activity 
in the serum. A contribution to the importance of 
examination of enzyme relations in the serum]. Klin 
Wochenschr 1962;40:962-9.11. Schmidt E, Schmidt FW. [Enzyme determinations in the 
serum in liver diseases. Function patterns as a means of 
diagnosis]. Enzymol Biol Clin (Basel) 1963;79:1-52.12. Forster G, Filippa G, Landolt M. [The significance of 
glutamate dehydrogenase for the differential diagnosis 
of jaundices]. Helv Med Acta 1963;30:672-84.13. Konttinen A, Härtel G, Louhija A. Multiple serum 
enzyme analyses in chronic alcoholics. Acta Med Scand 
1970;188:257-64.14. Aronsen KF, Hanson A, Nosslin B. The value of gamma 
glutamyl transpeptidase in differentiating viral hepatitis 
from obstructive jaundice. A statistical comparison with 
alkaline phosphatase. Acta Chir Scand 1965;130:92-9.15. Aronsen KF, Nosslin B, Pihl B. The value of gamma-
glutamyl transpeptidase as a screen test for liver tumour.

---

### Chunk 18/30
**Article:** Cardiologia VIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.596

itamina B6.
*   **Suplementação e Fatores Confundidores:**
    *   Quando a homocisteína está alta apesar de B12 e folato normais, investigar deficiência de B6, colina, betaína ou consumo excessivo de cafeína.
    *   A suplementação deve ser feita com as formas ativas: metilcobalamina (B12), piridoxal-5-fosfato (P5P, para B6) e metilfolato.
### 3. Biomarcadores Inflamatórios e Modulação Genética
*   **Gama GT (GGT) e Leucócitos:**
    *   A Gama GT elevada, mesmo dentro da referência, é um marcador de toxicidade crônica e risco cardiovascular. O objetivo é mantê-la no quartil inferior.
    *   Um aumento nos leucócitos, mesmo dentro da normalidade, pode indicar inflamação subclínica crônica, associada à lesão vascular.
*   **Modulação Genética (Sirtuínas):**
    *   Os genes SIRT1 e SIRT6 sinalizam vias de proteção cardiovascular e longevidade.

---

### Chunk 19/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** other | **Similarity:** 0.592

ratio above 1 but a 
high AST/ALT ratio is suggestive of either recent exposure or 
advanced alcoholic liver disease.105 The transaminases alone, or in combination, were not helpful in identifying heavy 
drinking in the NHANES study106 and others have found that the AST/ALT ratio may fall with increasing consumption.107Another argument that the association of an AST/ALT ratio 
of over 2.0 with alcoholic cirrhosis is more to do with recent 
alcohol exposure rather than cirrhosis per se is the fact that other 
causes of liver related death such as primary biliary cirrhosis108 and primary sclerosing cholangitis109 are associated with AST/

122   Clin Biochem Rev Vol 34 November 2013
ALT ratios of above 1.0 but not 2.0. Other non-hepatic alcohol related diseases such as oesophageal cancer also have AST/
ALT ratios >2.0 as a risk factor.110 Furthermore other acute hepatic toxicities, e.g.

---

### Chunk 20/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** other | **Similarity:** 0.591

vasc Diabetol 2012;11:117.148. Sookoian S, Pirola CJ. Alanine and aspartate 
aminotransferase and glutamine-cycling pathway: their 
roles in pathogenesis of metabolic syndrome. World J 
Gastroenterol 2012;18:3775-81.149. Khedmat H, Fallahian F, Abolghasemi H, 
Hajibeigi B, Attarchi Z, Alaeddini F, et al. Serum �γ-glutamyltransferase, alanine aminotransferase, and 
aspartate aminotransferase activity in Iranian healthy 
blood donor men. World J Gastroenterol 2007;13:889-
94.150. Hsu CH, Wang JY, Chen YL, Liu CC, Chang YL, 
Chen HS, et al. Relationships between alanine 
aminotransferase levels, abnormal liver echogenicity, 
and metabolic syndrome. J Am Board Fam Med 
2011;24:407-14.151. Goessling W, Massaro JM, Vasan RS, D’Agostino 
RB Sr, Ellison RC, Fox CS. Aminotransferase levels 
and 20-year risk of metabolic syndrome, diabetes, 
and cardiovascular disease. Gastroenterology 
2008;135:1935-44.152. Uslusoy HS, Nak SG, Gülten M, Bıyıklı Z.

---

### Chunk 21/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** discussion | **Similarity:** 0.585

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

### Chunk 22/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** other | **Similarity:** 0.584

GN, Boyko EJ, Lee SP. The prevalence and 
predictors of elevated serum aminotransferase activity 
in the United States in 1999-2002. Am J Gastroenterol 
2006;101:76-82.118. Fuchs CS, Stampfer MJ, Colditz GA, Giovannucci EL, 
Manson JE, Kawachi I, et al. Alcohol consumption and 
mortality among women. N Engl J Med 1995;332:1245-50.119. Jackson R, Beaglehole R. Alcohol consumption 
guidelines: relative safety vs absolute risks and benefits. 
Lancet 1995;346:716.120. Loomba R, Bettencourt R, Barrett-Connor E. Synergistic 
association between alcohol intake and body mass index 
with serum alanine and aspartate aminotransferase 
levels in older adults: the Rancho Bernardo Study. 
Aliment Pharmacol Ther 2009;30:1137-49.121. Adams LA, Knuiman MW, Divitini ML, Olynyk JK. 
Body mass index is a stronger predictor of alanine 
aminotransaminase levels than alcohol consumption. J 
Gastroenterol Hepatol 2008;23:1089-93.122. Park EY, Lim MK, Oh JK, Cho H, Bae MJ, Yun 
EH, et al.

---

### Chunk 23/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** methods | **Similarity:** 0.580

anbay A, Angulo P, Taniai M, Burgart 
LJ, Lindor KD, et al. Hepatocyte apoptosis and fas 
expression are prominent features of human nonalcoholic 
steatohepatitis. Gastroenterology 2003;125:437-43.6.  Wroblewski F. The significance of alterations in 
transaminase activities of serum and other body fluids. 
Adv Clin Chem 1958;1:313-51.7. Goldberg DM. Enzymes in the diagnosis of myocardial 
infarction and liver disease. Ann Clin Biochem 
1971;8:195-200.8. Goldberg DM. Clinical Enzymology. In: Ellis GP, 
West GB, editors. Progress in Medicinal Chemistry. 
Amsterdam: North Holland; 1976;13:1-158.9. Goldberg DM, Fletcher MJ, Watts C. Serum adenosine 
deaminase activity in hepatic disease: a comparative 
enzymological evaluation. Clin Chim Acta 1966;14:720-
8.10. Schmidt E, Schmidt FW. [Methods and value of 
determination of glutamic acid dehydrogenase activity 
in the serum. A contribution to the importance of 
examination of enzyme relations in the serum].

---

### Chunk 24/30
**Article:** MFI - PROGRAMAÇÃO METABÓLICA - AULA 05 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.578

ma GT como Marcador de Toxicidade**
    *   O Gama GT (GGT) não é apenas um marcador hepático, mas também um indicador de exposição a toxinas ambientais e um marcador para câncer de mama.
    *   É um marcador chave para suspeitar de toxicidades, mesmo que não esteja sempre aumentado. Estar no quartil superior já é um sinal de alerta.
*   **Aplicativos para Identificação de Toxinas em Produtos**
    *   São mencionados dois aplicativos (nomes não especificados) que ajudam a identificar toxinas em produtos de beleza e skincare, escaneando o código de barras.
    *   São mais eficazes para produtos americanos e úteis para pacientes que fazem enxoval no exterior.
### 3. Riscos de Medicamentos na Gestação
*   **Aumento do Uso de Medicamentos**
    *   O número de mulheres que tomam medicamentos durante a gravidez mais do que dobrou nos últimos 30 anos.
    *   Atualmente, 9 em cada 10 mulheres tomam pelo menos um medicamento durante a gestação.

---

### Chunk 25/30
**Article:** Bases Metabólicas das Doenças Crônicas e Gerenciamento - Glicação 2 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.577

e jejum ideal ≤6 µU/mL; aceitável até 10 µU/mL. Paciente refere 4–5 µU/mL.
  - Hemoglobina glicada: melhor indicador de glicação do que glicemia de jejum; pode estar elevada mesmo com glicemia normal (compatível com picos glicêmicos e alta carga glicêmica).
  - Glicemia de jejum isolada é exame “pobre” para avaliar glicação.
  - Recomendada curva insulinêmica-glicêmica para avaliar resistência insulínica e resposta a carboidratos.
- Genética:
  - Polimorfismos em FTO (5 variantes), MC4R (múltiplas), PPAR-γ, TCF7L2, SOD, CETP, BDNF, CCK, LEP e LEP-R associados a resistência insulínica, obesidade, menor saciedade, risco de diabetes e transporte lipídico desfavorável.
- Sinais e correlações fisiopatológicas:
  - Interação AGE-RAGE ativa NF-κB, aumenta citocinas e ROS, perpetuando inflamação crônica.
  - Excesso de ultraprocessados e preparos em alta temperatura eleva produtos de glicação avançada (AGEs).

---

### Chunk 26/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** other | **Similarity:** 0.576

istical comparison with 
alkaline phosphatase. Acta Chir Scand 1965;130:92-9.15. Aronsen KF, Nosslin B, Pihl B. The value of gamma-
glutamyl transpeptidase as a screen test for liver tumour. 
Acta Chir Scand 1970;136:17-22.16. Betro MG, Oon RC, Edwards JB. Gamma-glutamyl 
transpeptidase in diseases of the liver and bone. Am J 
Clin Pathol 1973;60:672-8.17. Goldberg DM, Ellis G. Mathematical and computer-
assisted procedures in the diagnosis of liver and biliary 
tract disorders. Adv Clin Chem 1978;20:49-128.18. Karmen A, Wroblewski F, Ladue JS. Transaminase 
activity in human blood. J Clin Invest 1955;34:126-31.19. Kay HD. Plasma phosphatase: method of determination, 
some properties of the enzyme. J Biol Chem 
1930;89:235-47.20. Gutman EB, Sproul EE, Gutman AB. Significance of 
increased phosphatase activity of bone at the site of 
osteoplastic metastases secondary to carcinoma of the 
prostate gland. Am J Cancer 1936;28:485-95.21. Rej R.

---

### Chunk 27/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** other | **Similarity:** 0.575

y glutamate.148 Only ALT levels predict the progression to metabolic syndrome149,150 while both ALT and AST predict the progression to diabetes.151In NASH, fibrosis may be present with normal transaminase 
levels152 particularly if high transaminase reference limits (>>40 IU/L) are used.

---

### Chunk 28/30
**Article:** Emagrecimento - Parte V (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.574

dadeiros fatores de confusão.
### 3. Fisiopatologia e Abordagem da Esteatose Hepática Gordurosa Não Alcoólica (EHGNA/DHGNA)
- A EHGNA é a principal causa de cirrose atualmente, superando as causas alcoólicas, e está associada ao sobrepeso e alimentação inadequada.
- **Fisiopatologia:** Aumento da ingesta calórica -> aumento da insulina e resistência periférica -> aumento da lipogênese -> acúmulo de gordura no fígado, podendo evoluir para inflamação (esteato-hepatite) e cirrose.
- **Marcadores:** A ferritina pode estar elevada em 20-50% dos casos, enquanto TGO e TGP não são fidedignos (normais em até 75% dos casos).
- O diagnóstico (ex: ultrassonografia) pode servir como um "susto bom" para motivar o paciente a mudar o estilo de vida.
### 4. Impacto da Frutose e Intervenções no Estilo de Vida
- O consumo excessivo de frutose, principalmente via xarope de milho (HFCS) em industrializados, é um grande fator para a EHGNA.

---

### Chunk 29/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** other | **Similarity:** 0.574

onger predictor of alanine 
aminotransaminase levels than alcohol consumption. J 
Gastroenterol Hepatol 2008;23:1089-93.122. Park EY, Lim MK, Oh JK, Cho H, Bae MJ, Yun 
EH, et al. Independent and supra-additive effects of 
alcohol consumption, cigarette smoking, and metabolic 
syndrome on the elevation of serum liver enzyme levels. 
PLoS One 2013;8:e63439.123. Lee DH, Ha MH, Christiani DC. Body weight, alcohol 
consumption and liver enzyme activity—a 4-year 
follow-up study. Int J Epidemiol 2001;30:766-70.124. Nagata K, Suzuki H, Sakaguchi S. Common pathogenic 
mechanism in development progression of liver injury 
caused by non-alcoholic or alcoholic steatohepatitis. J 
Toxicol Sci 2007;32:453-68.125. Lieber CS. Alcoholic fatty liver: its pathogenesis and 
mechanism of progression to inflammation and fibrosis. 
Alcohol 2004;34:9-19.126. Tappy L, Lê KA. Does fructose consumption contribute 
to non-alcoholic fatty liver disease? Clin Res Hepatol 
Gastroenterol 2012;36:554-60.127.

---

### Chunk 30/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** other | **Similarity:** 0.573

thout alcoholic hepatitis. 
Hepatology 1984;4:893-6.97. Rej R. Aspartate aminotransferase activity and 
isoenzyme proportions in human liver tissues. Clin 
Chem 1978;24:1971-9.98. Ishii H, Okuno F, Shigeta Y, Tsuchiya M. Enhanced 
serum glutamic oxaloacetic transaminase activity of 
mitochondrial origin in chronic alcoholics. Curr Alcohol 
1979;5:101-8.99. Nalpas B, Vassault A, Charpin S, Lacour B, Berthelot 
P. Serum mitochondrial aspartate aminotransferase as 
a marker of chronic alcoholism: diagnostic value and 
interpretation in a liver unit. Hepatology 1986;6:608-14.100. Nalpas B, Vassault A, Le Guillou A, Lesgourgues B, 
Ferry N, Lacour B, et al. Serum activity of mitochondrial 
aspartate aminotransferase: a sensitive marker of 
alcoholism with or without alcoholic hepatitis. 
Hepatology 1984;4:893-6.101. Okuno F, Ishii H, Kashiwazaki K, Takagi S, Shigeta Y, 
Arai M, et al.

---

