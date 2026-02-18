# ScoreItem: Higiene e Acompanhamento

**ID:** `019bf31d-2ef0-76f9-8525-94f68cf574ab`
**FullName:** Higiene e Acompanhamento (Histórico de doenças - Saúde bucal - Situação odontológica atual)

**Preparation Metadata:**
- Quality Grade: **GOOD**
- Total Chunks: 30 de 4 artigos
- Avg Similarity: 0.533

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-76f9-8525-94f68cf574ab`.**

```json
{
  "score_item_id": "019bf31d-2ef0-76f9-8525-94f68cf574ab",
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

**ScoreItem:** Higiene e Acompanhamento (Histórico de doenças - Saúde bucal - Situação odontológica atual)

**30 chunks de 4 artigos (avg similarity: 0.533)**

### Chunk 1/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XX (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.606

eas como alimentação, estilo de vida, exercício e traumas.
    - É crucial que o profissional se torne o "general" da história do paciente, coordenando a abordagem de saúde e buscando conhecimento contínuo em diversas áreas para obter melhores resultados.
### 2. Relação entre Saúde Bucal e Doenças Sistêmicas
*   **Inflamação Crônica e Focos Ocultos**
    - Uma inflamação crônica e silenciosa, que pode desencadear doenças autoimunes ou câncer, pode ter origem em focos bucais não diagnosticados, como doença periodontal, canais maltratados e cavitações.
    - Um caso clínico ilustra como sintomas neurológicos complexos foram resolvidos após o tratamento de uma infecção dentária crônica.
*   **Periodontite e Diabetes Tipo 2**
    - Estudos demonstram uma associação bidirecional: o diabetes piora a doença periodontal, e a doença periodontal piora o controle do diabetes.

---

### Chunk 2/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.558

. J. Dent. Res. 91(2), 161166. https:// doi. org/ 10. 1177/ 00220 34511 431583 (2012). 34. Myllymaki, V. et al. Association between periodontal condition and the development of type 2 diabetes mellitusResults from a 15-year follow-up study. J. Clin. Periodontol. https:// doi. org/ 10. 1111/ jcpe. 13005 (2018). 35. Winning, L., Patterson, C. C., Neville, C. E., Kee, F. & Linden, G. J. Periodontitis and incident type 2 diabetes: A prospective cohort study. J. Clin. Periodontol. 44(3), 266274. https:// doi. org/ 10. 1111/ jcpe. 12691 (2017). 36. Lee, J. H., Choi, J. K., Jeong, S. N. & Choi, S. H. Charlson comorbidity index as a predictor of periodontal disease in elderly par-ticipants. J. Periodontal Implant Sci. 48(2), 92102. https:// doi. org/ 10. 5051/ jpis. 2018. 48.2. 92 (2018). 37. Chiu, S. Y. et al. Temporal sequence of the bidirectional relationship between hyperglycemia and periodontal disease: A community-based study of 5,885 Taiwanese aged 3544 years (KCIS No. 32).

---

### Chunk 3/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XX (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.558

sem base em evidências.
## ❓ Perguntas
- [Inserir Pergunta/Dúvida/Confusão]
## 📚 Tarefas
- [ ] 1. Adotar uma abordagem mais integrada, investigando a saúde bucal de pacientes como possível causa de problemas sistêmicos crônicos.
- [ ] 2. Realizar triagem de diabetes em pacientes com periodontite e, inversamente, avaliar a saúde periodontal em pacientes diabéticos.
- [ ] 3. Buscar conhecimento contínuo e colaboração com profissionais de outras áreas da saúde para uma abordagem holística do paciente.
- [ ] 4. Investigar causas incomuns de inflamação (canais maltratados, cavitações, materiais dentários) ao avaliar pacientes com condições crônicas não resolvidas.
- [ ] 5. Pesquisar os artigos mencionados sobre a relação entre periodontite, AVC (2021) e aterosclerose (2024).
- [ ] 6. Estudar os mecanismos de autoimunidade desencadeados pela disbiose oral (mimetismo molecular, translocação bacteriana, etc.).
- [ ] 7.

---

### Chunk 4/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.555

ssociation was observed for a CPI score ≤ 2, but an increase in the incidence of diabetes was observed aer a CPI score ≥ 3 (score 3: SRR (95% CI): 1.38 (1.02, 1.87); score 4: SRR (95% CI): 2.33 (1.11, 4.87); p for non-linearity:0.090; n = 2) (Fig.2A). e non-linear analysis for the relationship between PPD and the incidence of diabetes showed an increased relative risk for diabetes up to a PPD of 3mm, then the graph reached a plateau with no further increase in risk (PPD 1.0mm: SRR (95% CI): 1.15 (0.83, 1.59); PPD 2.5mm: SRR (95% CI): 1.27 (0.83, 1.95), PPD 3,5mm: SRR (95% CI): 1.30 (0.90, 1.89), and PPD 4,5mm: SRR (95% CI): 1.31 (0.86, 2.01), based on n = 2; p for non-linearity: 0.653) (Fig.2B). e certainty of evidence was judged as moderate for the association between periodontal disease and incidence of diabetes mellitus (TableS6).Diabetes mellitus and incidence of periodontal disease.

---

### Chunk 5/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XX (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.550

ar geral.
---
### Evidências Principais
**O tratamento da periodontite melhora significativamente os marcadores de controle glicêmico e inflamação em pacientes com diabetes tipo 2, reduzindo a hemoglobina glicada em 0,56 e a proteína C reativa em 1,8.**
- Uma meta-análise, que selecionou 9 ensaios clínicos randomizados de um total de 402 estudos potenciais, investigou a relação entre doenças periodontais e diabetes tipo 2.
- A intervenção periodontal resultou em uma redução significativa na hemoglobina glicada, com um intervalo de confiança de 85%.
- O tratamento também levou a uma redução média de 1,8 na proteína C reativa, um marcador de inflamação sistêmica, com um intervalo de confiança de 95%.
- A periodontite também está associada a outras condições sistêmicas, como o diabetes tipo 1.

---

### Chunk 6/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.545

op on the Classication of Peri-odontal and Peri-Implant Diseases and Conditions. J. Periodontol. 89(Suppl 1), S173S182. https:// doi. org/ 10. 1002/ jper. 17- 0721 (2018). 51. Nazir, M. A. et al. e burden of diabetes, its oral complications and their prevention and management. Open Access Macedonian J. Med. Sci. 6(8), 15451553. https:// doi. org/ 10. 3889/ oamjms. 2018. 294 (2018). 52. Salvi, G. E. et al. Inammatory mediator response as a potential risk marker for periodontal diseases in insulin-dependent diabetes mellitus patients. J. Periodontol. 68(2), 127135. https:// doi. org/ 10. 1902/ jop. 1997. 68.2. 127 (1997).
	
	
Getulio José Mattos Do Amaral Filho - getfilho@yahoo.com.br - CPF: 034.983.039-88

9
Scientific Reports |        (2021) 11:13686  | https://doi.org/10.1038/s41598-021-93062-6
 53. Noack, B. et al. Periodontal infections contribute to elevated systemic C-reactive protein level. J. Periodontol. 72(9), 12211227.

---

### Chunk 7/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XX (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.545

# Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XX

**Source:** https://web.plaud.ai/share/d0d71763827796819::YXdzOnVzLXdlc3QtMg

---

## Reasoning Summary

# A Relação entre Saúde Oral e Doenças Sistémicas: Uma Visão Integrativa
## Odontologia Funcional Integrativa e Saúde Sistémica
### Visão e Abordagem Holística
A Odontologia Funcional Integrativa não é uma especialidade, mas sim uma visão e um modo de trabalhar que engloba a odontologia biológica. Nesta abordagem, o profissional de saúde deve abordar o ser humano como um todo, compreendendo diversas áreas como nutrição, estilo de vida e traumas, mesmo que não atue diretamente nelas. O objetivo é identificar a origem dos problemas, mesmo que fora da sua área de atuação direta, para obter melhores resultados e tornar-se o "general" do tratamento do paciente.

---

### Chunk 8/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.544

Higgins JPT TJ, Chandler J, Cumpston M, Li T, Page MJ, Welch VA (editors). Cochrane Handbook for Systematic Reviews of Interventions version 6.0 (updated July 2019). in Cochrane. www. train ing. cochr ane. org/ handb ook (2019). 28. Demmer, R. T., Jacobs, D. R. Jr. & Desvarieux, M. Periodontal disease and incident type 2 diabetes: Results from the First National Health and Nutrition Examination Survey and its epidemiologic follow-up study. Diabetes Care 31(7), 13731379. https:// doi. org/ 10. 2337/ dc08- 0026 (2008). 29. Kebede, T. G. et al. Does periodontitis aect diabetes incidence and haemoglobin A1c change? An 11-year follow-up study. Diabetes Metab. 44(3), 243249. https:// doi. org/ 10. 1016/j. diabet. 2017. 11. 003 (2018). 30. Lee, J. H. et al. Association between periodontal disease and non-communicable diseases: A 12-year longitudinal health-examinee cohort study in South Korea. Medicine 96(26), e7398. https:// doi. org/ 10. 1097/ md. 00000 00000 007398 (2017). 31.

---

### Chunk 9/30
**Article:** Infancy Dietary Patterns, Development, and Health: An Extensive Narrative Review (2022)
**Journal:** Children (Basel)
**Section:** other | **Similarity:** 0.543

tors. J.
Dent. Res. 2017, 96, 380–387. [CrossRef] [PubMed]
121. Dye, B.A.; Afful, J.; Thornton-Evans, G.; Iafolla, T. Overview and quality assurance for the oral health component of the National
Health and Nutrition Examination Survey (NHANES), 2011–2014. BMC Oral Health 2019, 19, 95. [CrossRef] [PubMed]
122. de Castilho, A.R.F.; Mialhe, F.L.; de Souza Barbosa, T.; Puppin-Rontani, R.M. Influence of family environment on children’s oral
health: A systematic review. J. Pediatr. 2013, 89, 116–123. [CrossRef] [PubMed]
123. Singh, A.; Bharathi, M.P.; Sequeira, P.; Acharya, S.; Bhat, M. Oral health status and practices of 5 and 12 year old Indian tribal
children. J. Clin. Pediatr. Dent. 2011, 35, 325–330. [CrossRef]
124. Scardina, G.A.; Messina, P. Good oral health and diet. J. Biomed. Biotechnol. 2012, 2012, 720692. [CrossRef]
125. Chi, D.L.; Scott, J.M. Added Sugar and Dental Caries in Children: A Scientific Update and Future Steps. Dent. Clin. N. Am. 2019,
63, 17–33.

---

### Chunk 10/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.539

s showed that the prevalence for diabetes was much 
Figure2.  Non-linear doseresponse meta-analysis for periodontal disease, dened as (A) CPI-score and (B) dened as PPD, and incidence of diabetes mellitus.
	
	
Getulio José Mattos Do Amaral Filho - getfilho@yahoo.com.br - CPF: 034.983.039-88

6
Scientific Reports |        (2021) 11:13686  | https://doi.org/10.1038/s41598-021-93062-6
lower for self-reported periodontitis compared to clinical periodontal  measurements42, our meta-analyses did not show a signicant dierence here. However, subgroup meta-analyses relied on small numbers of studies and more studies with accurate assessment methods are needed. Although it was rarely applied among the studies, assessment via CAL has been considered the gold standard for the classication of chronic  
periodontitis19.

---

### Chunk 11/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.536

ssing teeth is less dicult and less time-consuming. However, tooth loss is not a reliable marker for  
periodontitis17, as caries is another leading cause for tooth  loss18. us, we decided to exclude tooth loss as exposure/outcome, to minimize bias.If dierent studies reported on similar data (same exposure and outcome), we selected the study with the largest number of participants and cases.Data extraction and risk of bias assessment. e extraction of the data from the studies was con-ducted by one investigator (JS or SS) and checked by two other investigators (JB or MN). Each inconsistency was 
	
	
Getulio José Mattos Do Amaral Filho - getfilho@yahoo.com.br - CPF: 034.983.039-88

3
Scientific Reports |        (2021) 11:13686  | https://doi.org/10.1038/s41598-021-93062-6
debated until agreement was reached.

---

### Chunk 12/30
**Article:** Pedro Neuro - Neurologia Funcional Integrativa 2 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.535

3 min) ou HIIT (4 min/dia).
- [ ] 7. Melhorar a saúde bucal com raspagem de língua (cobre) e probióticos.
- [ ] 8. Avaliar e tratar fontes de inflamação crônica: infecções silenciosas (nasais, bucais), exposição a mofo e metais tóxicos. Investigar CIRS quando aplicável.
- [ ] 9. Para quem vai passar por cirurgia, utilizar o pool de suplementos sugerido para mitigar a neurotoxicidade da anestesia.
- [ ] 10. Discutir com um profissional de saúde a suplementação direcionada com base nos resultados da cognoscopia.

---

## SOAP

> Data e Hora: 2025-11-18 14:44:23
> Paciente:
> Diagnóstico:
## Histórico de Diagnóstico:
1. Histórico Médico:
2. Histórico de Medicação: Inserir mais aqui
## Subjetivo:
- Conteúdo educacional/apresentação sobre prevenção e manejo de risco para doença de Alzheimer, sem relato direto de queixas de um paciente específico.

---

### Chunk 13/30
**Article:** Pedro Neuro - Neurologia Funcional Integrativa 2 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.535

ccus salivarius, Lactobacillus sakei), raspador de língua de cobre, evitar dormir de boca aberta; atenção a periodontite/gengivite (Porphyromonas gingivalis).
  - Precauções perioperatórias:
    - Suplementação iniciada 1 semana antes e mantida por 2 semanas após anestesia/cirurgia para mitigar neurotoxicidade (redução de glutationa, risco de hipóxia/hipotensão, uso de antibióticos).
  - Programas de estilo de vida:
    - ReCODE/MAP personalizados conforme cognoscopia: metas de passos, prancha, dieta mediterrânea/Keto Flex e técnicas de respiração.
  - Exercício:
    - Caminhadas diárias: meta ≥5.000 passos, ideal ~10.000.
    - Musculação com ênfase em prancha (até 3 minutos totais/dia).
    - HIIT: protocolos curtos (ex.: 20s forte/10s leve, 8 ciclos, ~4 minutos).
  - Dieta:
    - Ketoflex 12-3 (12 horas de jejum diário, 3 horas entre jantar e sono, abordagem flexitariana com cetose monitorada).

---

### Chunk 14/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.533

our meta-analysis, the characteristics of all included studies are presented in the Supplemental Material (TableS4). A list of the excluded studies and the corresponding reasons are shown in the supplement (TableS3).Periodontal disease and incidence of diabetes mellitus. We identied 10 studies that investigated the association between periodontal disease and incidence of diabetes mellitus with a total of 427,620 par-ticipants and 114,361 identied cases of diabetes mellitus over a mean follow-up period of 9.9years (range 
	
	
Getulio José Mattos Do Amaral Filho - getfilho@yahoo.com.br - CPF: 034.983.039-88

4
Scientific Reports |        (2021) 11:13686  | https://doi.org/10.1038/s41598-021-93062-6
517years)13,2835. Four studies achieved a low, three of them a moderate, and three a high overall risk of bias (TableS5).

---

### Chunk 15/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.532

l cause of tooth extraction in a sample of US male adults. Caries Res. 23(3), 200205. https:// doi. org/ 10. 1159/ 00026 1178 (1989). 19. Armitage, G. C. Development of a classication system for periodontal diseases and conditions. Ann. Periodontol. 4(1), 16. https:// doi. org/ 10. 1902/ annals. 1999.4. 1.1 (1999). 20. Akazawa, H. Periodontitis and diabetes mellitus: Be true to your teeth. Int. Heart J. 59(4), 680682. https:// doi. org/ 10. 1536/ ihj. 18- 410 (2018). 21. Papanas, N. & Ziegler, D. Risk factors and comorbidities in diabetic neuropathy: An update 2015. Rev. Diabet. Stud. 12(12), 4862. https:// doi. org/ 10. 1900/ RDS. 2015. 12. 48 (2015). 22. DerSimonian, R. & Laird, N. Meta-analysis in clinical trials. Control Clin. Trials 7(3), 177188 (1986). 23. Greenland, S. & Longnecker, M. P. Methods for trend estimation from summarized dose-response data, with applications to meta-analysis. Am. J. Epidemiol. 135(11), 13011309 (1992). 24.

---

### Chunk 16/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.529

ch as Russells  PI28, or the  CPI13,33,37 to classify the disease. Periodontitis diagnosis based on the PI is critical because this index is only visually assessed and does not include clinical measurements, and it includes gingivitis as an early form of  
periodontitis48. Although the CPI is characterized by its reproducibility and  
simplication49, it is not considered sucient to describe the extent of periodontal  disease50. In summary, the assessment and denition of periodontal disease vary widely across studies and there are no consistent thresholds for CAL/PPD and numbers of aected teeth to determine whether the disease is present or not. e Division of Oral Health at the Centers for Disease Control and Preven-tion in collaboration with the American Academy of Periodontology has provided a denition that combines measurements of CAL and PPD to assess periodontal disease to avoid misinterpretation of the periodontal  status48.

---

### Chunk 17/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XX (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.528

rose (2024).
- [ ] 6. Estudar os mecanismos de autoimunidade desencadeados pela disbiose oral (mimetismo molecular, translocação bacteriana, etc.).
- [ ] 7. Suspender a prescrição de solução de Lugol até aprofundar os estudos sobre o tema, para garantir uma prática baseada em evidências e segurança.

---

## Quantitative Data

### Narrativa Quantitativa
A análise de múltiplos estudos revela uma forte conexão entre a saúde bucal, especificamente a periodontite, e condições sistêmicas graves como diabetes e AVC. O tratamento periodontal demonstra melhorias clínicas significativas, enquanto investigações sobre o flúor na água levantam preocupações sobre seu impacto na função tireoidiana, destacando a interconexão entre exposições ambientais, saúde bucal e bem-estar geral.

---

### Chunk 18/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.527

021) 11:13686  | https://doi.org/10.1038/s41598-021-93062-6
 53. Noack, B. et al. Periodontal infections contribute to elevated systemic C-reactive protein level. J. Periodontol. 72(9), 12211227. https:// doi. org/ 10. 1902/ jop. 2000. 72.9. 1221 (2001). 54. Genco, R. J., Grossi, S. G., Ho, A., Nishimura, F. & Murayama, Y. A proposed model linking inammation to obesity, diabetes, and periodontal infections. J. Periodontol. 76(11 Suppl), 20752084. https:// doi. org/ 10. 1902/ jop. 2005. 76. 11-S. 2075 (2005). 55. Sanz, M. et al. Scientic evidence on the links between periodontal diseases and diabetes: Consensus report and guidelines of the joint workshop on periodontal diseases and diabetes by the International Diabetes Federation and the European Federation of Periodontology. J. Clin. Periodontol. 45(2), 138149. https:// doi. org/ 10. 1111/ jcpe. 12808 (2018). 56. Genco, R. J. et al. Screening for diabetes mellitus in dental practices: A eld trial. J. Am. Dent. Assoc.

---

### Chunk 19/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XX (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.524

micas, como o diabetes tipo 1.
**A periodontite dobra o risco de Acidente Vascular Cerebral (AVC), conforme evidenciado por uma análise de 10 estudos envolvendo até 15.792 pacientes acompanhados por até 15 anos.**
- Uma análise de 10 estudos, com publicações recentes em 2021 e 2024, investigou a associação entre periodontite e AVC.
- O número de participantes nesses estudos variou de 80 a 15.792, com um período de acompanhamento que chegou a 15 anos.
- A conclusão central é que indivíduos com periodontite têm o dobro de probabilidade de sofrer um AVC.

---

### Chunk 20/30
**Article:** Infancy Dietary Patterns, Development, and Health: An Extensive Narrative Review (2022)
**Journal:** Children (Basel)
**Section:** other | **Similarity:** 0.522

Biotechnol. 2012, 2012, 720692. [CrossRef]
125. Chi, D.L.; Scott, J.M. Added Sugar and Dental Caries in Children: A Scientific Update and Future Steps. Dent. Clin. N. Am. 2019,
63, 17–33. [CrossRef]
126. Olczak-Kowalczyk, D.; Turska, A.; Gozdowski, D.; Kaczmarek, U. Dental Caries Level and Sugar Consumption in 12-Year-Old
Children from Poland. Adv. Clin. Exp. Med. 2016, 25, 545–550. [CrossRef] [PubMed]
127. Gordon, N. Oral health care for children attending a malnutrition clinic in South Africa. Int. J. Dent. Hyg. 2007, 5, 180–186.
[CrossRef] [PubMed]
128. Peres, M.A.; Sheiham, A.; Liu, P.; Demarco, F.F.; Silva, A.E.R.; Assunção, M.C.; Menezes, A.M.; Barros, F.C.; Peres, K.G. Sugar
Consumption and Changes in Dental Caries from Childhood to Adolescence. J. Dent. Res. 2016, 95, 388–394. [CrossRef] [PubMed]
129. Sheiham, A.; James, W.P.T. A new understanding of the relationship between sugars, dental caries and fluoride use: Implications
for limits on sugars consumption.

---

### Chunk 21/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.521

. J. & Preshaw, P. M. Diabetes and periodontal disease: a two-way relationship. Br. Dent. J. 217(8), 433437. https:// doi. org/ 10. 1038/ sj. bdj. 2014. 907 (2014). 9. Nascimento, G. G., Leite, F. R. M., Vestergaard, P., Scheutz, F. & Lopez, R. Does diabetes increase the risk of periodontitis? A systematic review and meta-regression analysis of longitudinal prospective studies. Acta Diabetol. 55(7), 653667. https:// doi. org/ 10. 1007/ s00592- 018- 1120-4 (2018). 10. Lee, K. S. et al. e relationship between metabolic conditions and prevalence of periodontal disease in rural Korean elderly. Arch. Gerontol. Geriatr. 58(1), 125129. https:// doi. org/ 10. 1016/j. archg er. 2013. 08. 011 (2014). 11. Iwasaki, M. et al. Longitudinal relationship between metabolic syndrome and periodontal disease among Japanese adults aged >/=70 years: e Niigata Study. J. Periodontol. 86(4), 491498. https:// doi. org/ 10. 1902/ jop. 2015. 140398 (2015). 12. Taylor, G. W., Burt, B. A., Becker, M.

---

### Chunk 22/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XX (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.518

e 2,5 mg/L de flúor. O mecanismo proposto é a competição do flúor com o iodo (ambos halogéneos) pela captação na tiroide, o que pode levar ao hipotiroidismo.
Esta prática levanta questões sobre a liberdade individual versus medidas de saúde pública impostas, sendo a fluoretação uma forma de medicação obrigatória iniciada em 1945, que também serviu como forma de descarte de um subproduto industrial.
## Implicações Clínicas e Colaboração Interprofissional
A forte ligação entre a saúde oral e a sistémica exige uma mudança de paradigma na prática clínica.
- **Colaboração Vital:** A interação entre dentistas, médicos e outros profissionais de saúde é crucial para um tratamento eficaz, eliminando "guerras" e subdivisões.
- **Responsabilidade Médica:** Médicos devem estar atentos aos impactos das doenças bucais em condições crónicas, conforme destacado por entidades como a American Heart Association.

---

### Chunk 23/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.517

/ doi. org/ 10. 1111/j. 1600- 0765. 1987. tb020 65.x (1987). 47. Albandar, J. M., Brunelle, J. A. & Kingman, A. Destructive periodontal disease in adults 30 years of age and older in the United States, 19881994. J. Periodontol. 70(1), 1329. https:// doi. org/ 10. 1902/ jop. 1999. 70.1. 13 (1999). 48. Page, R. C. & Eke, P. I. Case denitions for use in population-based surveillance of periodontitis. J. Periodontol. 78(7 Suppl), 13871399. https:// doi. org/ 10. 1902/ jop. 2007. 060264 (2007). 49. Petersen, P. E. & Ogawa, H. Strengthening the prevention of periodontal disease: e WHO approach. J. Periodontol. 76(12), 21872193. https:// doi. org/ 10. 1902/ jop. 2005. 76. 12. 2187 (2005). 50. Papapanou, P. N. et al. Periodontitis: Consensus report of workgroup 2 of the 2017 World Workshop on the Classication of Peri-odontal and Peri-Implant Diseases and Conditions. J. Periodontol. 89(Suppl 1), S173S182. https:// doi. org/ 10. 1002/ jper. 17- 0721 (2018). 51. Nazir, M. A. et al.

---

### Chunk 24/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XX (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.516

ção bidirecional: o diabetes piora a doença periodontal, e a doença periodontal piora o controle do diabetes.
    - O tratamento periodontal demonstrou reduzir significativamente a hemoglobina glicada e a proteína C reativa em pacientes com diabetes tipo 2.
*   **Periodontite e Doenças Cardiovasculares**
    - **Hipertensão:** A periodontite está associada a um maior risco de hipertensão, e seu tratamento pode impactar positivamente o controle da pressão arterial.
    - **AVC:** Uma meta-análise mostrou que indivíduos com periodontite têm o dobro de probabilidade de sofrer um AVC, especialmente o isquêmico.
    - **Aterosclerose:** Um artigo de 2024 relaciona a periodontite ao desenvolvimento de aterosclerose em pacientes com síndrome metabólica.

---

### Chunk 25/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XX (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.515

gnificativas tanto na **hemoglobina glicada (HbA1c)**, com uma diferença média de -0,56%, quanto na **proteína C reativa (PCR)**, um marcador de inflamação sistémica. Isto comprova que tratar a periodontite melhora o controlo metabólico e reduz a inflamação em diabéticos.
### Doença Periodontal, Hipertensão e AVC
A periodontite está associada a um risco aumentado de hipertensão. Além disso, uma meta-análise de dez estudos revelou que indivíduos com periodontite têm o **dobro de probabilidade** de sofrer um Acidente Vascular Cerebral (AVC), incluindo AVC isquémico, em comparação com indivíduos saudáveis.
### Doença Periodontal, Aterosclerose e Síndrome Metabólica
Um estudo de 2024 concluiu que a periodontite promove o desenvolvimento de doença cardiovascular aterosclerótica em pacientes com componentes da síndrome metabólica (obesidade, disglicemia, hipertensão, hiperlipidemia).

---

### Chunk 26/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.514

e quality of reporting in research articles. JMT received salary from Evidence Partners, creator 
	
	
Getulio José Mattos Do Amaral Filho - getfilho@yahoo.com.br - CPF: 034.983.039-88

8
Scientific Reports |        (2021) 11:13686  | https://doi.org/10.1038/s41598-021-93062-6
of DistillerSR soware for systematic reviews; Evidence Partners was not involved in the design or outcomes of the statement, and the views expressed solely represent those of the author. 17. Holmlund, A. & Lind, L. Number of teeth is related to atherosclerotic plaque in the carotid arteries in an elderly population. J. Periodontol. 83(3), 287291. https:// doi. org/ 10. 1902/ jop. 2011. 110100 (2012). 18. Chauncey, H. H., Glass, R. L. & Alman, J. E. Dental caries. Principal cause of tooth extraction in a sample of US male adults. Caries Res. 23(3), 200205. https:// doi. org/ 10. 1159/ 00026 1178 (1989). 19. Armitage, G. C.

---

### Chunk 27/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.513

sease and non-communicable diseases: A 12-year longitudinal health-examinee cohort study in South Korea. Medicine 96(26), e7398. https:// doi. org/ 10. 1097/ md. 00000 00000 007398 (2017). 31. Lin, S. Y. et al. Association between periodontitis needing surgical treatment and subsequent diabetes risk: A population-based cohort study. J. Periodontol. 85(6), 779786. https:// doi. org/ 10. 1902/ jop. 2013. 130357 (2014). 32. Miyawaki, A., Toyokawa, S., Inoue, K., Miyoshi, Y. & Kobayashi, Y. Self-reported periodontitis and incident type 2 diabetes among male workers from a 5-year follow-up to my health up study. PLoS ONE 11(4), e0153464. https:// doi. org/ 10. 1371/ journ al. pone. 01534 64 (2016). 33. Morita, I. et al. Relationship between periodontal status and levels of glycated hemoglobin. J. Dent. Res. 91(2), 161166. https:// doi. org/ 10. 1177/ 00220 34511 431583 (2012). 34. Myllymaki, V. et al.

---

### Chunk 28/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** conclusion | **Similarity:** 0.512

eriodontal disease and that glycemic control is more dif-cult in this case. us, every initial examination should include a periodontal  
evaluation55. It has been shown, that 40.7% of the dental patients without a diagnosis of diabetes (< 45years) had HbA1c values around 5,7% or  higher56, thus, screening for diabetes mellitus in the dental oce is as important.In conclusion, there was a bidirectional association between periodontal disease and diabetes mellitus, even aer stratifying for major risk of bias. However, only few studies with low risk of bias were available. To strengthen these ndings more studies with valid assessment of periodontal diseases and diabetes are needed.

---

### Chunk 29/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.511

periodontitis in Taiwan: A nationwide cohort study. Diabetes Res. Clin. Pract. 150, 245252. https:// doi. org/ 10. 1016/j. diabr es. 2019. 03. 019 (2019). 41. Alshihayb, T. S., Kaye, E. A., Zhao, Y., Leone, C. W. & Heaton, B. A quantitative bias analysis to assess the impact of unmeasured confounding on associations between diabetes and periodontitis. J. Clin. Periodontol. 48(1), 5160. https:// doi. org/ 10. 1111/ jcpe. 13386 (2021). 42. Ziukaite, L., Slot, D. E. & Van der Weijden, F. A. Prevalence of diabetes mellitus in people clinically diagnosed with periodontitis: A systematic review and meta-analysis of epidemiologic studies. J. Clin. Periodontol. 45(6), 650662. https:// doi. org/ 10. 1111/ jcpe. 12839 (2018). 43. Tonetti, M. S., Greenwell, H. & Kornman, K. S. Staging and grading of periodontitis: Framework and proposal of a new classica-tion and case denition. J. Clin. Periodontol. 45(Suppl 20), S149S161. https:// doi. org/ 10. 1111/ jcpe. 12945 (2018). 44.

---

### Chunk 30/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.510

isease among Japanese adults aged >/=70 years: e Niigata Study. J. Periodontol. 86(4), 491498. https:// doi. org/ 10. 1902/ jop. 2015. 140398 (2015). 12. Taylor, G. W., Burt, B. A., Becker, M. P., Genco, R. J. & Shlossman, M. Glycemic control and alveolar bone loss progression in type 2 diabetes. Ann. Periodontol. 3(1), 3039. https:// doi. org/ 10. 1902/ annals. 1998.3. 1. 30 (1998). 13. Ide, R., Hoshuyama, T., Wilson, D., Takahashi, K. & Higashi, T. Periodontal disease and incident diabetes: A seven-year study. J. Dent. Res. 90(1), 4146. https:// doi. org/ 10. 1177/ 00220 34510 381902 (2011). 14. Engebretson, S. P. et al. Gingival crevicular uid levels of interleukin-1beta and glycemic control in patients with chronic peri-odontitis and type 2 diabetes. J. Periodontol. 75(9), 12031208. https:// doi. org/ 10. 1902/ jop. 2004. 75.9. 1203 (2004). 15. Hayden, J. A., van der Windt, D. A., Cartwright, J. L., Cote, P. & Bombardier, C.

---

