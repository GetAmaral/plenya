# ScoreItem: Histórico Periodontal

**ID:** `019bf31d-2ef0-7f92-900b-138dfe299a5f`
**FullName:** Histórico Periodontal (Histórico de doenças - Saúde bucal - Histórico odontológico (principalmente amálgamas antigos com mercúrio))

**Preparation Metadata:**
- Quality Grade: **EXCELLENT**
- Total Chunks: 30 de 3 artigos
- Avg Similarity: 0.631

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-7f92-900b-138dfe299a5f`.**

```json
{
  "score_item_id": "019bf31d-2ef0-7f92-900b-138dfe299a5f",
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

**ScoreItem:** Histórico Periodontal (Histórico de doenças - Saúde bucal - Histórico odontológico (principalmente amálgamas antigos com mercúrio))

**30 chunks de 3 artigos (avg similarity: 0.631)**

### Chunk 1/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XX (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.710

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
**Section:** results | **Similarity:** 0.680

. J. Dent. Res. 91(2), 161166. https:// doi. org/ 10. 1177/ 00220 34511 431583 (2012). 34. Myllymaki, V. et al. Association between periodontal condition and the development of type 2 diabetes mellitusResults from a 15-year follow-up study. J. Clin. Periodontol. https:// doi. org/ 10. 1111/ jcpe. 13005 (2018). 35. Winning, L., Patterson, C. C., Neville, C. E., Kee, F. & Linden, G. J. Periodontitis and incident type 2 diabetes: A prospective cohort study. J. Clin. Periodontol. 44(3), 266274. https:// doi. org/ 10. 1111/ jcpe. 12691 (2017). 36. Lee, J. H., Choi, J. K., Jeong, S. N. & Choi, S. H. Charlson comorbidity index as a predictor of periodontal disease in elderly par-ticipants. J. Periodontal Implant Sci. 48(2), 92102. https:// doi. org/ 10. 5051/ jpis. 2018. 48.2. 92 (2018). 37. Chiu, S. Y. et al. Temporal sequence of the bidirectional relationship between hyperglycemia and periodontal disease: A community-based study of 5,885 Taiwanese aged 3544 years (KCIS No. 32).

---

### Chunk 3/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XX (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.671

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
**Section:** other | **Similarity:** 0.661

021) 11:13686  | https://doi.org/10.1038/s41598-021-93062-6
 53. Noack, B. et al. Periodontal infections contribute to elevated systemic C-reactive protein level. J. Periodontol. 72(9), 12211227. https:// doi. org/ 10. 1902/ jop. 2000. 72.9. 1221 (2001). 54. Genco, R. J., Grossi, S. G., Ho, A., Nishimura, F. & Murayama, Y. A proposed model linking inammation to obesity, diabetes, and periodontal infections. J. Periodontol. 76(11 Suppl), 20752084. https:// doi. org/ 10. 1902/ jop. 2005. 76. 11-S. 2075 (2005). 55. Sanz, M. et al. Scientic evidence on the links between periodontal diseases and diabetes: Consensus report and guidelines of the joint workshop on periodontal diseases and diabetes by the International Diabetes Federation and the European Federation of Periodontology. J. Clin. Periodontol. 45(2), 138149. https:// doi. org/ 10. 1111/ jcpe. 12808 (2018). 56. Genco, R. J. et al. Screening for diabetes mellitus in dental practices: A eld trial. J. Am. Dent. Assoc.

---

### Chunk 5/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.658

. J. & Preshaw, P. M. Diabetes and periodontal disease: a two-way relationship. Br. Dent. J. 217(8), 433437. https:// doi. org/ 10. 1038/ sj. bdj. 2014. 907 (2014). 9. Nascimento, G. G., Leite, F. R. M., Vestergaard, P., Scheutz, F. & Lopez, R. Does diabetes increase the risk of periodontitis? A systematic review and meta-regression analysis of longitudinal prospective studies. Acta Diabetol. 55(7), 653667. https:// doi. org/ 10. 1007/ s00592- 018- 1120-4 (2018). 10. Lee, K. S. et al. e relationship between metabolic conditions and prevalence of periodontal disease in rural Korean elderly. Arch. Gerontol. Geriatr. 58(1), 125129. https:// doi. org/ 10. 1016/j. archg er. 2013. 08. 011 (2014). 11. Iwasaki, M. et al. Longitudinal relationship between metabolic syndrome and periodontal disease among Japanese adults aged >/=70 years: e Niigata Study. J. Periodontol. 86(4), 491498. https:// doi. org/ 10. 1902/ jop. 2015. 140398 (2015). 12. Taylor, G. W., Burt, B. A., Becker, M.

---

### Chunk 6/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XX (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.658

ar geral.
---
### Evidências Principais
**O tratamento da periodontite melhora significativamente os marcadores de controle glicêmico e inflamação em pacientes com diabetes tipo 2, reduzindo a hemoglobina glicada em 0,56 e a proteína C reativa em 1,8.**
- Uma meta-análise, que selecionou 9 ensaios clínicos randomizados de um total de 402 estudos potenciais, investigou a relação entre doenças periodontais e diabetes tipo 2.
- A intervenção periodontal resultou em uma redução significativa na hemoglobina glicada, com um intervalo de confiança de 85%.
- O tratamento também levou a uma redução média de 1,8 na proteína C reativa, um marcador de inflamação sistêmica, com um intervalo de confiança de 95%.
- A periodontite também está associada a outras condições sistêmicas, como o diabetes tipo 1.

---

### Chunk 7/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.650

op on the Classication of Peri-odontal and Peri-Implant Diseases and Conditions. J. Periodontol. 89(Suppl 1), S173S182. https:// doi. org/ 10. 1002/ jper. 17- 0721 (2018). 51. Nazir, M. A. et al. e burden of diabetes, its oral complications and their prevention and management. Open Access Macedonian J. Med. Sci. 6(8), 15451553. https:// doi. org/ 10. 3889/ oamjms. 2018. 294 (2018). 52. Salvi, G. E. et al. Inammatory mediator response as a potential risk marker for periodontal diseases in insulin-dependent diabetes mellitus patients. J. Periodontol. 68(2), 127135. https:// doi. org/ 10. 1902/ jop. 1997. 68.2. 127 (1997).
	
	
Getulio José Mattos Do Amaral Filho - getfilho@yahoo.com.br - CPF: 034.983.039-88

9
Scientific Reports |        (2021) 11:13686  | https://doi.org/10.1038/s41598-021-93062-6
 53. Noack, B. et al. Periodontal infections contribute to elevated systemic C-reactive protein level. J. Periodontol. 72(9), 12211227.

---

### Chunk 8/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XX (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.646

sclerótica em pacientes com componentes da síndrome metabólica (obesidade, disglicemia, hipertensão, hiperlipidemia). A inflamação crónica originada na boca pode ser o gatilho para estas condições.
### Disbiose Oral e Doenças Autoimunes
A disbiose oral, especialmente associada à periodontite, pode influenciar respostas autoimunes através da perda da tolerância imunológica. Os mecanismos incluem:
- **Translocação Microbiana:** Bactérias e toxinas da boca entram na corrente sanguínea.
- **Mimetismo Molecular:** O sistema imunitário ataca patógenos orais e, por semelhança estrutural, ataca também tecidos do próprio corpo.
- **Hiperprodução de Citocinas:** A inflamação crónica leva a uma produção excessiva de citocinas (ex: IL-17), desregulando a resposta imune. A bactéria *Porphyromonas gingivalis* é considerada chave neste processo.

---

### Chunk 9/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.645

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

### Chunk 10/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XX (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.645

gnificativas tanto na **hemoglobina glicada (HbA1c)**, com uma diferença média de -0,56%, quanto na **proteína C reativa (PCR)**, um marcador de inflamação sistémica. Isto comprova que tratar a periodontite melhora o controlo metabólico e reduz a inflamação em diabéticos.
### Doença Periodontal, Hipertensão e AVC
A periodontite está associada a um risco aumentado de hipertensão. Além disso, uma meta-análise de dez estudos revelou que indivíduos com periodontite têm o **dobro de probabilidade** de sofrer um Acidente Vascular Cerebral (AVC), incluindo AVC isquémico, em comparação com indivíduos saudáveis.
### Doença Periodontal, Aterosclerose e Síndrome Metabólica
Um estudo de 2024 concluiu que a periodontite promove o desenvolvimento de doença cardiovascular aterosclerótica em pacientes com componentes da síndrome metabólica (obesidade, disglicemia, hipertensão, hiperlipidemia).

---

### Chunk 11/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.634

Higgins JPT TJ, Chandler J, Cumpston M, Li T, Page MJ, Welch VA (editors). Cochrane Handbook for Systematic Reviews of Interventions version 6.0 (updated July 2019). in Cochrane. www. train ing. cochr ane. org/ handb ook (2019). 28. Demmer, R. T., Jacobs, D. R. Jr. & Desvarieux, M. Periodontal disease and incident type 2 diabetes: Results from the First National Health and Nutrition Examination Survey and its epidemiologic follow-up study. Diabetes Care 31(7), 13731379. https:// doi. org/ 10. 2337/ dc08- 0026 (2008). 29. Kebede, T. G. et al. Does periodontitis aect diabetes incidence and haemoglobin A1c change? An 11-year follow-up study. Diabetes Metab. 44(3), 243249. https:// doi. org/ 10. 1016/j. diabet. 2017. 11. 003 (2018). 30. Lee, J. H. et al. Association between periodontal disease and non-communicable diseases: A 12-year longitudinal health-examinee cohort study in South Korea. Medicine 96(26), e7398. https:// doi. org/ 10. 1097/ md. 00000 00000 007398 (2017). 31.

---

### Chunk 12/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.633

sease and non-communicable diseases: A 12-year longitudinal health-examinee cohort study in South Korea. Medicine 96(26), e7398. https:// doi. org/ 10. 1097/ md. 00000 00000 007398 (2017). 31. Lin, S. Y. et al. Association between periodontitis needing surgical treatment and subsequent diabetes risk: A population-based cohort study. J. Periodontol. 85(6), 779786. https:// doi. org/ 10. 1902/ jop. 2013. 130357 (2014). 32. Miyawaki, A., Toyokawa, S., Inoue, K., Miyoshi, Y. & Kobayashi, Y. Self-reported periodontitis and incident type 2 diabetes among male workers from a 5-year follow-up to my health up study. PLoS ONE 11(4), e0153464. https:// doi. org/ 10. 1371/ journ al. pone. 01534 64 (2016). 33. Morita, I. et al. Relationship between periodontal status and levels of glycated hemoglobin. J. Dent. Res. 91(2), 161166. https:// doi. org/ 10. 1177/ 00220 34511 431583 (2012). 34. Myllymaki, V. et al.

---

### Chunk 13/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.633

l cause of tooth extraction in a sample of US male adults. Caries Res. 23(3), 200205. https:// doi. org/ 10. 1159/ 00026 1178 (1989). 19. Armitage, G. C. Development of a classication system for periodontal diseases and conditions. Ann. Periodontol. 4(1), 16. https:// doi. org/ 10. 1902/ annals. 1999.4. 1.1 (1999). 20. Akazawa, H. Periodontitis and diabetes mellitus: Be true to your teeth. Int. Heart J. 59(4), 680682. https:// doi. org/ 10. 1536/ ihj. 18- 410 (2018). 21. Papanas, N. & Ziegler, D. Risk factors and comorbidities in diabetic neuropathy: An update 2015. Rev. Diabet. Stud. 12(12), 4862. https:// doi. org/ 10. 1900/ RDS. 2015. 12. 48 (2015). 22. DerSimonian, R. & Laird, N. Meta-analysis in clinical trials. Control Clin. Trials 7(3), 177188 (1986). 23. Greenland, S. & Longnecker, M. P. Methods for trend estimation from summarized dose-response data, with applications to meta-analysis. Am. J. Epidemiol. 135(11), 13011309 (1992). 24.

---

### Chunk 14/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.628

xpected7. e disease is characterized by a chronic inammation of the entire periodontium that can irreparably destruct the tooth-surrounding tissue and result in the resorption of the alveolar bone. Consequences such as gingival bleeding, increased tooth mobility and tooth loss can be  expected8. Recently, a meta-analysis summarized ndings on glucose disturbance, including diabetes, and peri-odontal disease and indicated a positive association between these two  factors9. However, in this meta-analysis, studies with dierent exposures and outcomes were mixed. For example, the authors combined studies on dia-betes, prediabetes and diabetes  severity10,11. In addition, the outcome was a mixture of periodontal disease and progression of the  disease12. Moreover, there is indication that periodontal disease is a risk factor for diabetes  mellitus13. Both conditions are driven by inammatory processes, which might be a possible explanation for this bidirectional  association14.

---

### Chunk 15/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.628

isease among Japanese adults aged >/=70 years: e Niigata Study. J. Periodontol. 86(4), 491498. https:// doi. org/ 10. 1902/ jop. 2015. 140398 (2015). 12. Taylor, G. W., Burt, B. A., Becker, M. P., Genco, R. J. & Shlossman, M. Glycemic control and alveolar bone loss progression in type 2 diabetes. Ann. Periodontol. 3(1), 3039. https:// doi. org/ 10. 1902/ annals. 1998.3. 1. 30 (1998). 13. Ide, R., Hoshuyama, T., Wilson, D., Takahashi, K. & Higashi, T. Periodontal disease and incident diabetes: A seven-year study. J. Dent. Res. 90(1), 4146. https:// doi. org/ 10. 1177/ 00220 34510 381902 (2011). 14. Engebretson, S. P. et al. Gingival crevicular uid levels of interleukin-1beta and glycemic control in patients with chronic peri-odontitis and type 2 diabetes. J. Periodontol. 75(9), 12031208. https:// doi. org/ 10. 1902/ jop. 2004. 75.9. 1203 (2004). 15. Hayden, J. A., van der Windt, D. A., Cartwright, J. L., Cote, P. & Bombardier, C.

---

### Chunk 16/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XX (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.627

**Tipo de material** utilizado em implantes.
- **Desequilíbrios no microbioma oral (disbiose).**
Estes problemas podem manifestar-se como dores crónicas, desequilíbrios de força entre os lados do corpo e problemas em órgãos específicos.
## O Impacto da Saúde Bucal em Condições Sistémicas Específicas
### Doença Periodontal e Diabetes Tipo 2
Uma revisão sistemática e meta-análise confirma uma associação positiva e bidirecional entre periodontite e diabetes. O diabetes agrava a doença periodontal, e a doença periodontal, por sua vez, piora o controlo glicémico e a resistência à insulina, criando um ciclo vicioso.
Uma meta-análise de nove ensaios clínicos demonstrou que o tratamento periodontal não cirúrgico em pacientes com diabetes tipo 2 resultou em reduções significativas tanto na **hemoglobina glicada (HbA1c)**, com uma diferença média de -0,56%, quanto na **proteína C reativa (PCR)**, um marcador de inflamação sistémica.

---

### Chunk 17/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XX (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.627

ção bidirecional: o diabetes piora a doença periodontal, e a doença periodontal piora o controle do diabetes.
    - O tratamento periodontal demonstrou reduzir significativamente a hemoglobina glicada e a proteína C reativa em pacientes com diabetes tipo 2.
*   **Periodontite e Doenças Cardiovasculares**
    - **Hipertensão:** A periodontite está associada a um maior risco de hipertensão, e seu tratamento pode impactar positivamente o controle da pressão arterial.
    - **AVC:** Uma meta-análise mostrou que indivíduos com periodontite têm o dobro de probabilidade de sofrer um AVC, especialmente o isquêmico.
    - **Aterosclerose:** Um artigo de 2024 relaciona a periodontite ao desenvolvimento de aterosclerose em pacientes com síndrome metabólica.

---

### Chunk 18/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.624

periodontitis in Taiwan: A nationwide cohort study. Diabetes Res. Clin. Pract. 150, 245252. https:// doi. org/ 10. 1016/j. diabr es. 2019. 03. 019 (2019). 41. Alshihayb, T. S., Kaye, E. A., Zhao, Y., Leone, C. W. & Heaton, B. A quantitative bias analysis to assess the impact of unmeasured confounding on associations between diabetes and periodontitis. J. Clin. Periodontol. 48(1), 5160. https:// doi. org/ 10. 1111/ jcpe. 13386 (2021). 42. Ziukaite, L., Slot, D. E. & Van der Weijden, F. A. Prevalence of diabetes mellitus in people clinically diagnosed with periodontitis: A systematic review and meta-analysis of epidemiologic studies. J. Clin. Periodontol. 45(6), 650662. https:// doi. org/ 10. 1111/ jcpe. 12839 (2018). 43. Tonetti, M. S., Greenwell, H. & Kornman, K. S. Staging and grading of periodontitis: Framework and proposal of a new classica-tion and case denition. J. Clin. Periodontol. 45(Suppl 20), S149S161. https:// doi. org/ 10. 1111/ jcpe. 12945 (2018). 44.

---

### Chunk 19/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.618

ssociation was observed for a CPI score ≤ 2, but an increase in the incidence of diabetes was observed aer a CPI score ≥ 3 (score 3: SRR (95% CI): 1.38 (1.02, 1.87); score 4: SRR (95% CI): 2.33 (1.11, 4.87); p for non-linearity:0.090; n = 2) (Fig.2A). e non-linear analysis for the relationship between PPD and the incidence of diabetes showed an increased relative risk for diabetes up to a PPD of 3mm, then the graph reached a plateau with no further increase in risk (PPD 1.0mm: SRR (95% CI): 1.15 (0.83, 1.59); PPD 2.5mm: SRR (95% CI): 1.27 (0.83, 1.95), PPD 3,5mm: SRR (95% CI): 1.30 (0.90, 1.89), and PPD 4,5mm: SRR (95% CI): 1.31 (0.86, 2.01), based on n = 2; p for non-linearity: 0.653) (Fig.2B). e certainty of evidence was judged as moderate for the association between periodontal disease and incidence of diabetes mellitus (TableS6).Diabetes mellitus and incidence of periodontal disease.

---

### Chunk 20/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XX (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.617

micas, como o diabetes tipo 1.
**A periodontite dobra o risco de Acidente Vascular Cerebral (AVC), conforme evidenciado por uma análise de 10 estudos envolvendo até 15.792 pacientes acompanhados por até 15 anos.**
- Uma análise de 10 estudos, com publicações recentes em 2021 e 2024, investigou a associação entre periodontite e AVC.
- O número de participantes nesses estudos variou de 80 a 15.792, com um período de acompanhamento que chegou a 15 anos.
- A conclusão central é que indivíduos com periodontite têm o dobro de probabilidade de sofrer um AVC.

---

### Chunk 21/30
**Article:** Pedro Neuro - Neurologia Funcional Integrativa 2 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.610

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

### Chunk 22/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.608

our meta-analysis, the characteristics of all included studies are presented in the Supplemental Material (TableS4). A list of the excluded studies and the corresponding reasons are shown in the supplement (TableS3).Periodontal disease and incidence of diabetes mellitus. We identied 10 studies that investigated the association between periodontal disease and incidence of diabetes mellitus with a total of 427,620 par-ticipants and 114,361 identied cases of diabetes mellitus over a mean follow-up period of 9.9years (range 
	
	
Getulio José Mattos Do Amaral Filho - getfilho@yahoo.com.br - CPF: 034.983.039-88

4
Scientific Reports |        (2021) 11:13686  | https://doi.org/10.1038/s41598-021-93062-6
517years)13,2835. Four studies achieved a low, three of them a moderate, and three a high overall risk of bias (TableS5).

---

### Chunk 23/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.607

s provided a denition that combines measurements of CAL and PPD to assess periodontal disease to avoid misinterpretation of the periodontal  status48. In order to make clear statements about the association between periodontitis and diabetes mellitus, it should be ensured in the future studies that established assessment of periodontitis is applied, and thus, enables comparability between studies.ere are many possible explanations for the observed bidirectional associations between periodontitis and diabetes, which are related to inammatory processes. For example, on the one hand, untreated diabetes mel-litus, both type 1 or 2 diabetes, lead to metabolic disorders caused by  hyperglycemia51. Poor glycemic control in patients with diabetes has been shown to raise the level of systemic inammation markers, e.g. interleukin-1ß, in the gingival crevicular uid of a periodontal  
pocket14, which is associated with the onset and severity of peri-odontal  disease52.

---

### Chunk 24/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** conclusion | **Similarity:** 0.607

iodontal disease is a risk factor for diabetes  mellitus13. Both conditions are driven by inammatory processes, which might be a possible explanation for this bidirectional  association14. To draw clear conclusions on these associations, a systematic review and meta-analysis is needed that considers methodological challenges when combining the existing data from primary studies. First, the time sequence of exposure and outcome needs to be taken into account to obtain the direction of the association. Second, the measurement of periodontal disease diers between the studies. In observational studies, periodontal disease has been assessed as self-reported periodontitis, clinical measurements attained from oral examinations (e.g.

---

### Chunk 25/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.604

/ doi. org/ 10. 1111/j. 1600- 0765. 1987. tb020 65.x (1987). 47. Albandar, J. M., Brunelle, J. A. & Kingman, A. Destructive periodontal disease in adults 30 years of age and older in the United States, 19881994. J. Periodontol. 70(1), 1329. https:// doi. org/ 10. 1902/ jop. 1999. 70.1. 13 (1999). 48. Page, R. C. & Eke, P. I. Case denitions for use in population-based surveillance of periodontitis. J. Periodontol. 78(7 Suppl), 13871399. https:// doi. org/ 10. 1902/ jop. 2007. 060264 (2007). 49. Petersen, P. E. & Ogawa, H. Strengthening the prevention of periodontal disease: e WHO approach. J. Periodontol. 76(12), 21872193. https:// doi. org/ 10. 1902/ jop. 2005. 76. 12. 2187 (2005). 50. Papapanou, P. N. et al. Periodontitis: Consensus report of workgroup 2 of the 2017 World Workshop on the Classication of Peri-odontal and Peri-Implant Diseases and Conditions. J. Periodontol. 89(Suppl 1), S173S182. https:// doi. org/ 10. 1002/ jper. 17- 0721 (2018). 51. Nazir, M. A. et al.

---

### Chunk 26/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.604

ch as Russells  PI28, or the  CPI13,33,37 to classify the disease. Periodontitis diagnosis based on the PI is critical because this index is only visually assessed and does not include clinical measurements, and it includes gingivitis as an early form of  
periodontitis48. Although the CPI is characterized by its reproducibility and  
simplication49, it is not considered sucient to describe the extent of periodontal  disease50. In summary, the assessment and denition of periodontal disease vary widely across studies and there are no consistent thresholds for CAL/PPD and numbers of aected teeth to determine whether the disease is present or not. e Division of Oral Health at the Centers for Disease Control and Preven-tion in collaboration with the American Academy of Periodontology has provided a denition that combines measurements of CAL and PPD to assess periodontal disease to avoid misinterpretation of the periodontal  status48.

---

### Chunk 27/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.603

ssing teeth is less dicult and less time-consuming. However, tooth loss is not a reliable marker for  
periodontitis17, as caries is another leading cause for tooth  loss18. us, we decided to exclude tooth loss as exposure/outcome, to minimize bias.If dierent studies reported on similar data (same exposure and outcome), we selected the study with the largest number of participants and cases.Data extraction and risk of bias assessment. e extraction of the data from the studies was con-ducted by one investigator (JS or SS) and checked by two other investigators (JB or MN). Each inconsistency was 
	
	
Getulio José Mattos Do Amaral Filho - getfilho@yahoo.com.br - CPF: 034.983.039-88

3
Scientific Reports |        (2021) 11:13686  | https://doi.org/10.1038/s41598-021-93062-6
debated until agreement was reached.

---

### Chunk 28/30
**Article:** Bidirectional association between periodontal disease and diabetes mellitus: a systematic review and meta‑analysis of cohort studies (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.597

ng of periodontitis: Framework and proposal of a new classica-tion and case denition. J. Clin. Periodontol. 45(Suppl 20), S149S161. https:// doi. org/ 10. 1111/ jcpe. 12945 (2018). 44. Robinson, P. J. & Vitek, R. M. e relationship between gingival inammation and resistance to probe penetration. J. Periodontal Res. 14(3), 239243. https:// doi. org/ 10. 1111/j. 1600- 0765. 1979. tb002 29.x (1979). 45. Haajee, A. D. & Socransky, S. S. Attachment level changes in destructive periodontal diseases. J. Clin. Periodontol. 13(5), 461475. https:// doi. org/ 10. 1111/j. 1600- 051x. 1986. tb014 91.x (1986). 46. Carlos, J. P., Brunelle, J. A. & Wolfe, M. D. Attachment loss vs. pocket depth as indicators of periodontal disease: A methodologic note. J. Periodontal Res. 22(6), 524525. https:// doi. org/ 10. 1111/j. 1600- 0765. 1987. tb020 65.x (1987). 47. Albandar, J. M., Brunelle, J. A. & Kingman, A.

---

### Chunk 29/30
**Article:** Pedro Neuro - Neurologia Funcional Integrativa 2 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.595

C auxiliam na desintoxicação.
*   **Saúde Bucal**
    - Bactérias como a *Porphyromonas gingivalis* estão implicadas no Alzheimer.
    - Recomenda-se o uso de probióticos bucais, raspagem da língua com raspador de cobre e evitar dormir de boca aberta.
*   **Agentes Anestésicos**
    - A anestesia geral contribui para o declínio cognitivo. Recomenda-se um pool de suplementos antes e após cirurgias para minimizar os efeitos neurotóxicos.
### 3. Programas de Intervenção e Estilo de Vida
*   **Programa Recode**
    - Desenvolvido por Dale Bredesen, é um programa personalizado baseado nos resultados da cognoscopia.
    - É um "norte" para uma visão multifacetada do paciente, incluindo dieta Keto Flex, sono, estresse, suplementação e avaliação da síndrome da resposta inflamatória crônica (CIRS).
*   **Programa MAP (Movimento, Alimento, Pensamento)**
    - Desenvolvido pelo instrutor, foca em 10 itens essenciais.

---

### Chunk 30/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XX (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.593

tes se deslocam para a corrente sanguínea.
        - **Mimetismo Molecular:** O sistema imune ataca estruturas próprias por semelhança com antígenos bacterianos.
        - **Hiperprodução de Citocinas Inflamatórias:** A inflamação crônica desregula o sistema imune.
        - **Bactérias Específicas:** A *Porphyromonas gingivalis* é citada como crucial no desencadeamento da autoimunidade.
*   **Implicações Clínicas**
    - A compreensão dessa relação pode levar a novas terapias focadas na restauração da homeostase oral e sistêmica, exigindo integração entre práticas médicas e odontológicas.
### 4. Colaboração Interprofissional e Identificação de Causas
*   **Importância da Colaboração**
    - A interação entre todos os profissionais de saúde é vital para o bem-estar do paciente. Dentistas precisam entender os mecanismos inflamatórios, e médicos devem estar atentos aos impactos das doenças bucais.

---

