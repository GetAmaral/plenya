# ScoreItem: Tireoidectomia

**ID:** `019bf31d-2ef0-7db2-9844-d04c37701d92`
**FullName:** Tireoidectomia (Histórico de doenças - Cirurgias já realizadas - Cirurgias que interferem diretamente no escore)

**Preparation Metadata:**
- Quality Grade: **GOOD**
- Total Chunks: 30 de 9 artigos
- Avg Similarity: 0.567

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-7db2-9844-d04c37701d92`.**

```json
{
  "score_item_id": "019bf31d-2ef0-7db2-9844-d04c37701d92",
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

**ScoreItem:** Tireoidectomia (Histórico de doenças - Cirurgias já realizadas - Cirurgias que interferem diretamente no escore)

**30 chunks de 9 artigos (avg similarity: 0.567)**

### Chunk 1/30
**Article:** Post-thyroidectomy hypoparathyroidism: A clinical surgical dilemma (2024)
**Journal:** Saudi Medical Journal
**Section:** abstract | **Similarity:** 0.638

Thyroid disease represents a prevalent endocrine condition with surgical management remaining the preferred treatment option. Post-thyroidectomy hypoparathyroidism, stemming from inadvertent parathyroid gland injury or removal, constitutes the primary cause of hospitalization following thyroid surgery. This condition results in hypocalcemia and presents significant clinical challenges requiring prompt identification to prevent complications. The article examines clinical presentations, underlying risk factors, and therapeutic strategies for managing this condition. Additionally, surgical techniques designed to minimize occurrence are highlighted, including anatomical knowledge, preservation techniques, near-infrared fluorescence imaging, and parathyroid autotransplantation considerations. The review emphasizes preoperative assessment of biochemical markers and postoperative monitoring protocols to optimize patient outcomes and quality of life following thyroidectomy.

---

### Chunk 2/30
**Article:** Hipotireioidismo (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.634

a pacientes com sintomas persistentes, especialmente aqueles com polimorfismos genéticos (12-14% da população), tireoidectomizados (que perdem 10-20% da produção de T3) ou com doses de T4 acima de 1.2 mcg/kg.
**Achados Adicionais**
- Uma meta-análise de 2017 com 2 milhões de participantes mostrou que o hipotireoidismo é um fator de risco independente para mortalidade cardiovascular.
- Em um estudo com 21 mulheres inférteis com TSH entre 0,5 e 3,5, a otimização da dose de T4 para melhorar o T3 livre resultou em todas engravidando em três meses.
- A levotiroxina foi a segunda droga mais vendida nos EUA em 2019.
- Um estudo de 2001 mostrou que doses suprafisiológicas de hormônio tireoidiano (200-300 microgramas) aliviaram sintomas em pacientes com fibromialgia, uma condição onde 35% podem ter resistência periférica ao hormônio tireoidiano.

---

### Chunk 3/30
**Article:** Hipotireioidismo (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.600

duas doses.
- Guideline (2014): T4 padrão de cuidado; lacunas persistem; necessidade de biomarcadores superiores.
- Consensos recentes (2021–2022): individualizar por etiologia e comorbidades; estudos heterogêneos e curtos.
### 11. Uso do TSH na prática e ajustes
- TSH é útil, mas pode falhar; guias práticos orientam ajuste de dose.
- Recomenda-se conhecer dosagens e percentuais de ajuste; algoritmos baseados em TSH/T4 livre e sintomas.
### 12. Levotiroxina: horário e adesão
- Tomar em jejum pela manhã ou à noite (≥2h após refeição); bedtime pode melhorar TSH/T3 em alguns estudos.
- Ingestão com alimento reduz biodisponibilidade; consistência do horário é essencial.
### 13. Fatores que afetam absorção da levotiroxina
- Absorção 60–80% sob condições ótimas; dependente de pH gástrico e intestino delgado; pico 1–1,5h.
- Redução por: gestação, hipocloridria (IBP), gastrite atrófica, H.

---

### Chunk 4/30
**Article:** Hipotireioidismo (2024)
**Journal:** Pos Graduacao MFI
**Section:** methods | **Similarity:** 0.588

(~10%).
- Impacta conversão T4→T3, receptores periféricos e múltiplos sistemas (intestino, cérebro, cardiovascular, reprodutivo).
- Gatilhos: genéticos, alimentares, estilo de vida, químicos, infecciosos.
- Abordagem integrativa: tratar causas-raiz, desfazer “nós fisiológicos”, considerar T4+T3 em casos selecionados com autoimunidade.
### 30. Mensagens centrais de prática
- Integrar clínica, TSH, T3/T4 (metodologias acuradas), etiologia e biomarcadores teciduais.
- Personalizar metas além do TSH para restaurar função tecidual e qualidade de vida.
- Exercício físico como modulador-chave da sensibilidade do receptor tireoidiano.
- “Não é sobre hormônios; é sobre pessoas que os produzem.” Tratar o sistema antes de apenas repor hormônios.
## ❓ Perguntas
- [Inserir Pergunta/Confusão]
## 📚 Tarefas
- [ ] 1. Revisar protocolos de avaliação incluindo TSH, T4 livre e T3 livre com metodologias mais acuradas (ultrafiltração quando disponível).
- [ ] 2.

---

### Chunk 5/30
**Article:** Hipotireioidismo (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.586

cadas: TSH como absoluto, conversão uniforme T4→T3, normalidade populacional, exclusão do T3 como perigoso, etiologia irrelevante.
- Imunoensaios de T3/T4: variabilidade; ultrafiltração é mais acurada; risco de misclassificação de subclínico vs franco.
- Hipotireoidismo secundário pode cursar com TSH normal/baixo.
- TSH mais alto dentro do “normal” associa-se a pior qualidade de vida (2021).
- Biomarcadores teciduais auxiliares: colesterol total, LDL, Lp(a), SHBG, osteocalcina, N-telopeptídeo urinário, CK, mioglobina, ferritina, inibidor da ECA, G6PD.
- Meta-análise (2021, 99 estudos): T4 visando TSH ~3,3 não normaliza totalmente biomarcadores teciduais.
- Pequenas variações de T4/TSH impactam grande a taxa metabólica de repouso.
### 9. Evolução da terapia e evidências T4/T3
- Pêndulo histórico: clínica→laboratório→individualização com múltiplos marcadores.

---

### Chunk 6/30
**Article:** Hipotireioidismo (2024)
**Journal:** Pos Graduacao MFI
**Section:** methods | **Similarity:** 0.585

ir Pergunta/Confusão]
## 📚 Tarefas
- [ ] 1. Revisar protocolos de avaliação incluindo TSH, T4 livre e T3 livre com metodologias mais acuradas (ultrafiltração quando disponível).
- [ ] 2. Incorporar biomarcadores teciduais (colesterol, LDL, lipoproteína(a), SHBG, osteocalcina, N-telopeptídeo urinário, CK, mioglobina, ferritina, inibidor da ECA, G6PD) na monitorização terapêutica.
- [ ] 3. Investigar etiologia (Hashimoto, hipofisária, pós-cirúrgico) e ajustar conduta conforme causa.
- [ ] 4. Avaliar/corrigir carências nutricionais (ferro, selênio, zinco, vitaminas D/A/B/C/E, iodo, tirosina) e reduzir exposições (flúor excessivo, toxinas).
- [ ] 5. Considerar estresse crônico, cortisol, inflamação de baixo grau e microbioma intestinal na regulação do eixo HHT e no manejo.
- [ ] 6. Prescrever e monitorar exercício físico para melhorar sensibilidade do receptor tireoidiano.
- [ ] 7.

---

### Chunk 7/30
**Article:** Levothyroxine Therapy in Thyrodectomized Patients (2021)
**Journal:** Frontiers in Endocrinology
**Section:** abstract | **Similarity:** 0.583

Administration of optimal levothyroxine dosing is crucial for restoring euthyroidism post-thyroidectomy. Insufficient or excessive dosing may cause hypothyroidism or thyrotoxicosis with associated complications. Most literature addresses primary hypothyroidism treatment rather than post-thyroidectomy replacement. This review examined papers from the last 15 years, identifying eight schemes for calculating proper LT4 dosage. Approximately 75% of patients require dose adjustments, indicating body weight alone is insufficient for determining optimal dosing. Factors affecting LT4 requirements include medication interactions, gastrointestinal conditions, and formulation types. Recently developed liquid and soft gel capsule formulations demonstrate improved absorption compared to traditional tablets, particularly for patients with malabsorption issues.

---

### Chunk 8/30
**Article:** Hipotireioidismo (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.580

manejo.
- [ ] 6. Prescrever e monitorar exercício físico para melhorar sensibilidade do receptor tireoidiano.
- [ ] 7. Ajustar rotinas de tomada de LT4 (manhã em jejum vs noite ≥2 h após refeição) visando aderência e absorção; orientar jejum e evitar coadministração com alimentos/medicações.
- [ ] 8. Revisar medicações/alimentos que reduzem absorção/conversão (IBP, ferro, cálcio, beta-bloqueadores, análogos GLP-1, soja) e planejar espaçamento/alternativas.
- [ ] 9. Investigar má absorção em doses suprafisiológicas (≥200–300 μg): teste respiratório de lactose, endoscopia para H. pylori, triagem para doença celíaca; considerar parasitoses; avaliar SIBO/disbiose.
- [ ] 10. Em casos refratários com TSH “normal” e sintomas persistentes, reavaliar estratégia (ajuste de T4 ± T3), checar conversão periférica e polimorfismos quando possível.
- [ ] 11.

---

### Chunk 9/30
**Article:** 2012 ETA Guidelines: The Use of L-T4 + L-T3 in the Treatment of Hypothyroidism (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.579

evaluation system. J Clin Endo-
crinol Metab 2008;   93:   666–673. 
	
	
Getulio José Mattos Do Amaral Filho - getfilho@yahoo.com.br - CPF: 034.983.039-88

 Wiersinga  /Duntas  /Fadeyev  /Nygaard  /Vanderpump     Eur Thyroid J 2012;1:55–7170
  9 Ladenson PW: Psychological wellbeing in patients. Clin Endocrinol 2002;   57:   575–576.  10 Weetman AP: Whose thyroid hormone re-placement is it anyway? Clin Endocrinol 
2006;   64:   231–233.  11 Saravanan P, Chau W-F, Roberts N, Vedhara K, Greenwood R, Dayan CM: Psychological 
well-being in patients on ‘adequate’ doses of 
L-thyroxine: results of a large, controlled 
community-based questionnaire study. Clin 
Endocrinol 2002;   57:   577–585.  12 Saravanan P, Visser TJ, Dayan CM: Psycho-logical well-being correlates with free thy-
roxine but not free 3,5,3  -triiodothyronine levels in patients on thyroid hormone re-
placement. J Clin Endocrinol Metab 2006;   91:   3389–3393.

---

### Chunk 10/30
**Article:** Hipotireioidismo (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.577

cundário pode ter TSH normal/baixo.
- TSH mais alto dentro da referência associa-se a pior QoL em hipotireoidismo primário (2021).
- Biomarcadores teciduais: colesterol/LDL/Lp(a), SHBG, osteocalcina, N-telopeptídeo urinário, CK, mioglobina, ferritina, inibidor da ECA, G6PD.
- Meta-análise (2021, 99 estudos): T4 com TSH médio ~3,3 não normaliza vários biomarcadores celulares; correção laboratorial nem sempre resolve sintomas.
- Pequenas variações de TSH dentro da normalidade alteram taxa metabólica de repouso.
### 10. Terapia T4 vs. T4+T3: evidências e diretrizes
- Escobar Morreale (1996) propôs que T4+T3 restaura eutiroidismo; meta-análise (2006) não mostrou benefício consistente.
- Diretriz Europeia (2012): considerar combinação; proporção inicial 13:1 a 20:1; T3 em duas doses.
- Guideline (2014): T4 padrão de cuidado; lacunas persistem; necessidade de biomarcadores superiores.

---

### Chunk 11/30
**Article:** Hipotireioidismo (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.574

cias T4/T3
- Pêndulo histórico: clínica→laboratório→individualização com múltiplos marcadores.
- Meta-análises até 2006 sem benefício claro da combinação; guideline europeu (2012) reconhece possíveis benefícios.
- Endocrine Reviews 2022: orientações práticas ainda baseadas em TSH, com reconhecimento de limitações.
- Futuro: incorporar biomarcadores teciduais, genéticos (polimorfismos de deiodinases/receptor TR), metabolômica.
### 10. Prática clínica: ajuste de T4, horários e absorção
- TSH permanece útil para ajustes percentuais, interpretado com clínica e outros marcadores.
- Tomada: manhã em jejum ou à noite (≥2 h após refeição); bedtime pode melhorar TSH/T3 em alguns.
- Absorção: depende de acidez gástrica; IBP/hipocloridria reduzem biodisponibilidade (usuários de IBP precisam ~37% mais dose).

---

### Chunk 12/30
**Article:** 2012 ETA Guidelines: The Use of L-T4 + L-T3 in the Treatment of Hypothyroidism (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.573

se of thyroid hor-
mone replacement  
[66]
 . One should aim at normalization of serum TSH, free T4, free T3 and free T4/free T3 ratio. 
Because there is substantial variation between free T4 and 
free T3 assays, each laboratory should endeavour to es-
tablish its own reference ranges. It seems reasonable to 
check thyroid function tests about 6–8 weeks after start-
ing L-T4 + L-T3 combination therapy, as it is likely that a 
new equilibrium has been established at that time. Ab-
normal test results would qualify for dose adjustment. It 
is suggested to change the dose of just one of the compo-
nents of the combination therapy. If serum free T3 is too 
low, an increase of L-T3 dose is logical, as a higher L-T4 
dose is likely to increase serum free T4. If serum free T4 
is too high, an increase of L-T3 dose will lower serum free 
T4. Therefore in many instances adjustment of the L-T3 
dose will be indicated.

---

### Chunk 13/30
**Article:** Hipotireioidismo (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.572

ões clássicas; amplo impacto sistêmico (“fio de cabelo à unha do pé”).
### 8. História do diagnóstico e tratamento; TSH e limitações
- Do mixedema ao PBI (1909); taxa metabólica basal (1919); T4/T3 identificados (1926/1952).
- Transição 1950–1970: extratos com altas doses; tireotoxicose frequente.
- 1970–1973: conversão periférica; dosagens de TSH/T3/T4; foco em normalização laboratorial.
- Variabilidade histórica de dose/qualidade; até 1997 sem levotiroxina aprovada pelo FDA.
### 9. Armadilhas diagnósticas e biomarcadores teciduais
- TSH reflete função hipofisária; uso isolado é limitado.
- Conversão T4→T3 não é previsível; deiodinases variam por tecido/contexto.
- Imunoensaios de T3 variáveis; ultrafiltração reclassifica casos.
- Hipotireoidismo secundário pode ter TSH normal/baixo.
- TSH mais alto dentro da referência associa-se a pior QoL em hipotireoidismo primário (2021).

---

### Chunk 14/30
**Article:** Hipotireioidismo (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.567

ar T3 livre e segurança para eventual ajuste terapêutico (LT4 e, em casos selecionados, T3), monitorando Pro-BNP e PCR-us.
- [ ] 16. Em obesos, interpretar TSH/T3/T4 à luz da adaptação metabólica e leptina; tratar resistência insulínica e promover perda de peso; monitorar T3 livre em platôs.
- [ ] 17. Na infertilidade feminina/masculina, incluir TSH, T4 livre, T3 livre e prolactina precocemente; otimizar T3 livre em mulheres já em LT4; tratar por 3–12 meses antes de procedimentos.
- [ ] 18. Em depressão, avaliar anti-TPO e considerar T3 como adjuvante com monitorização rigorosa.
- [ ] 19. Em fibromialgia, investigar disfunção tireoidiana/autoimunidade; ponderar ensaios com hormônios tireoidianos com cautela.
- [ ] 20. Educar pacientes sobre fatores que impactam a tireoide (estresse, dieta hipocalórica excessiva, toxinas, infecções), diferenças entre sintéticos e extratos, e limites das formulações disponíveis.
- [ ] 21.

---

### Chunk 15/30
**Article:** 2012 ETA Guidelines: The Use of L-T4 + L-T3 in the Treatment of Hypothyroidism (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.561

n of T4 into T3 makes it possible 
to achieve normal serum T3 concentrations in humans 
treated with L-T4  
[26]
 , but it is argued that the T3 levels ob-served might be lower than the T3 levels in the premorbid 
state. This possibility has been evaluated in a study compar-
ing thyroid function in the same patients before and after 
total thyroidectomy for goiter, benign nodular thyroid dis-
ease or thyroid cancer  
[27]
 . Patients were euthyroid before surgery without medication and after surgery with L-T4 
medication. Pre- and postoperative serum TSH (1.18  8  0.58 vs. 1.30  8  1.89 mU/l) and serum T3 (1.99  8  0.41 vs. 1.96  8  0.43 nmol/l) concentrations were not different.  Another suggestion has been that well-being is opti-mized by fine adjustment of L-T4 dosage, aiming for a 
serum TSH in the lower reference range  
[28]
 .

---

### Chunk 16/30
**Article:** 2012 ETA Guidelines: The Use of L-T4 + L-T3 in the Treatment of Hypothyroidism (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.561

or can be produced lo-
cally from T4. Because of tissue heterogeneity, pituitary 
TSH secretion may not reflect what happens in other tar-
get tissues, and therefore serum TSH alone may not be a 
good marker for the adequacy of thyroid hormone re-
placement  
[67, 68, 70]
 . Theoretically, thyroid hormone re-placement therapy should aim not only at normalization 
of serum TSH but also at normalization of serum free T4, 
free T3 and free T4/free T3 ratio.  The pharmacodynamic equivalence of L-T4 and L-T3 has been recently studied in a randomized, double-blind, 
cross-over study in 10 thyroidectomized patients  
[71]
 . The target (TSH  6 0.5 ^ 1.5 mU/l for at least 30 days) was reached by an average daily dose of 40.3  8  11.3   g L-T3 and 115.2  8  38.5   g L-T4 (L-T3:L-T4 ratio 0.36  8  0.06). It was concluded that therapeutic substitution of L-T3 for 
L-T4 was achieved at approximately 1:   3 ratio.

---

### Chunk 17/30
**Article:** 2012 ETA Guidelines: The Use of L-T4 + L-T3 in the Treatment of Hypothyroidism (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.552

blets (including 
those containing animal thyroid extract) are potentially 
harmful as due to their relatively high T3 content they 
carry a risk of inducing symptoms of thyroid hormone 
excess. It is therefore recommended to use separate L-T4 
and L-T3 tablets in L-T4 + L-T3 combination therapy. 
One may prefer to use brand names for L-T3 replacement 
in order to avoid errors in manufacturing L-T3 tablets of 
just a few micrograms by the local pharmacy. The 12.5-
  g Cytomel tablets can be divided in smaller parts of 3.125   g each.  Monitoring of L-T4 + L-T3 combination therapy should be done by thyroid function tests in blood samples 
withdrawn before morning medication has been taken. 
In doing so, one avoids the risk of measuring relatively 
high free T3 levels in the absorption phase of thyroid hor-
mone replacement  
[66]
 . One should aim at normalization of serum TSH, free T4, free T3 and free T4/free T3 ratio.

---

### Chunk 18/30
**Article:** Hipotireioidismo (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.552

 de cardiovascular
- T3 modula canais iônicos, frequência/contratilidade, débito, vasorrelaxamento, SRAA, oxigenação, mitocôndria.
- Meta-análise (~2 milhões, 2017): hipotireoidismo aumenta mortalidade CV e geral.
- Hipotireoidismo subclínico: disfunção cardíaca leve reversível com T4.
- Baixo T3 em UTI/eventos agudos correlaciona com maior mortalidade; em ICC, menor conversão T4→T3, maior D3/rT3, citocinas; dobutamina aumenta T3 livre.
- T3 em baixa dose pós-IAM/ICC: melhora remodelamento, marcadores (Pro-BNP, PCR-us) e arritmias atriais, com segurança em protocolos selecionados.
### 20. Obesidade e eixo adipotireoidiano
- Obesidade: inflamação crônica, estresse oxidativo, disfunção metabólica.
- Hipotireoidismo franco: ganho de peso modesto (~2–3 kg atribuíveis à tireoide).
- TSH elevado pode ser consequência da adiposidade; leptina elevada influencia TSH e autoimunidade.

---

### Chunk 19/30
**Article:** Hipotireioidismo (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.551

síquicos) com TSH/T4 normalizados.
- Tireoidectomizados (perda de T3 intratireoidiano).
- Doses altas de T4 (>1,0–1,2 mcg/kg/dia) sem controle.
- Preferência/experiência positiva do paciente (~50%).
- Autoimunidade ativa (anti-TPO elevado).
### 28. Como adicionar T3: posologia e monitorização
- Produção diária típica de T3 ~30 mcg; evitar doses altas pela regulação da D2/D3.
- Pico sérico 2–4h pós-dose orienta coleta laboratorial.
- Estratégia: reduzir T4 e adicionar T3; para cada 25 mcg reduzidos de T4, adicionar 2,5–7,5 mcg de T3; fracionar T3 em 2 tomadas.
### 29. Quando evitar combinação e uso isolado de T3
- Não combinar em assintomáticos adequadamente tratados com T4.
- Gestação: evitar T3; feto depende de T4 materno até ~22–24 semanas.
- Evitar em doença cardiovascular instável, malignidade ativa, psiquiatria não controlada.
- T3 isolado: uso específico (p.ex., preparo para radioiodo); meia-vida curta; liberação prolongada sem padronização.

---

### Chunk 20/30
**Article:** MFI - Reposição Hormonal - AULA 06 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.550

- Hipercalemia, aumento de ácido úrico e glicemia (resistência insulínica).
   - Perda de sódio, magnésio e cálcio.
   - Em homens: ginecomastia.
   - Em mulheres: predominância estrogênica, amenorreia; risco de câncer de mama a longo prazo é desconhecido, porém teoricamente possível.
### 4. Princípios da Prática Clínica em Terapia Hormonal
* **Individualização e Cautela**
   - Não existe “fórmula pronta”. Cada paciente aromatiza e responde de forma diferente.
   - Abordagem deve começar “do menos para o mais”, com doses baixas e ajustes conforme resposta e exames.
* **Importância da Avaliação e Acompanhamento**
   - Avaliar de perto, especialmente no início, com exames laboratoriais (estradiol, DHT, testosterona) e avaliação clínica (história, composição corporal, hábitos).
   - Para ganhar experiência, pode-se começar tratando familiares/pessoas próximas, permitindo acompanhamento mais rigoroso.

---

### Chunk 21/30
**Article:** Vitamina D (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.548

ia (já realizada; dose não especificada).
  - Suplementação: vitamina D (inicialmente 30.000 UI/dia), vitaminas B2 e B12, magnésio; possíveis fitoterápicos/antroposóficos (não especificados).
  - Inserir mais aqui.
- Próximos Passos/Exames:
  - Monitorar 25(OH)D visando faixa de 40–100 ng/mL conforme recomendações da ABN, com individualização por resposta clínica e laboratorial.
  - Monitorar PTH para manter próximo ao limite inferior da normalidade, evitando hiperparatireoidismo relativo ou supressão excessiva.
  - Monitorar cálcio sérico total e ionizado, fósforo, função renal; avaliar hipercalciúria periodicamente.
  - Revisar função hepática e medicamentos que interferem nas enzimas do citocromo P450 (corticoides, antiepilépticos).
  - Considerar avaliação de magnésio (preferencialmente estado intracelular), riboflavina (B2), vitamina A, zinco, função tireoidiana, perfil lipídico e hábitos alimentares.

---

### Chunk 22/30
**Article:** Best practices in the laboratory diagnosis, prognostication, prediction, and monitoring of Graves' disease: role of TRAbs (2024)
**Journal:** BMC Endocr Disord
**Section:** results | **Similarity:** 0.547

Tijssen JG, Wiersinga WM. Predict-ing the risk of recurrence before the start of antithyroid drug therapy in patients with Graves hyperthyroidism. J Clin Endocrinol Metab. 2016;101(4):13819. 34. Struja T, Kaeslin M, Boesiger F, Jutzi R, Imahorn N, Kutz A, Bernasconi L, Mundwiler E, Mueller B, Christ-Crain M, et al. External validation of the GREAT score to predict relapse risk in Graves disease: results from a multicenter, retrospective study with 741 patients. Eur J Endocrinol. 2017;176(4):4139. 35. Carvalho GA, Perez CL, Ward LS. The clinical use of thyroid function tests. Arq Bras Endocrinol Metabol. 2013;57(3):193204. 36. Wiersinga WM. Graves Disease: Can It Be Cured? Endocrinol Metab (Seoul). 2019;34(1):2938. 37. Rajput R, Goel V. Indeﬁnite antithyroid drug therapy in toxic Graves disease: What are the cons. Indian J Endocrinol Metab. 2013;17(Suppl 1):S88-92. 38.

---

### Chunk 23/30
**Article:** Hipotireioidismo (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.546

lidade (usuários de IBP precisam ~37% mais dose).
- Fatores: gestação, comorbidades, medicações (ferro/cálcio/soja/GLP-1), idade, massa magra; dissolução gástrica e absorção no intestino delgado.
### 11. Má absorção: causas e investigação
- Redutores de absorção: hipocloridria, IBP, H. pylori, gastrite atrófica, doença celíaca, intolerância à lactose, ferro/cálcio, soja, idade avançada, gastroparesia (diabetes, análogos GLP-1), parasitoses.
- Pistas: necessidade de 200–300 μg de LT4 sugere má absorção.
- Fluxo diagnóstico: teste respiratório de lactose; endoscopia para H. pylori; triagem para doença celíaca; investigar SIBO/disbiose.
### 12. Epidemiologia, sintomas persistentes e satisfação
- Levotiroxina: 2º medicamento mais vendido nos EUA em 2019; aumento de 63% em 20 anos.
- 5–15% dos tratados têm sintomas persistentes; 70–77% relatam insatisfação (2018–2020).

---

### Chunk 24/30
**Article:** 2012 ETA Guidelines: The Use of L-T4 + L-T3 in the Treatment of Hypothyroidism (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.545

64–770.  27 Jonklaas J, Davidson B, Bhagat S, Soldin SJ: Triiodothyronine levels in athyreotic indi-
viduals during levothyroxine therapy. JAMA 
2008;   299:   769–777.  28 Carr D, McLeod DT, Parry G, Thornes HM: Fine adjustment of thyroxine replacement 
dosage: comparison of the thyrotrophin re-
leasing hormone test using a sensitive thyro-
trophin assay with measurement of free thy-
roid hormones and clinical assessment. Clin 
Endocrinol 1988;   28:   325–333.  29 Walsh JP, Ward LC, Burke V, Bhagat CI, Shiels L, Henley D, Gillett MJ, Gilbert R, 
Tanner M, Stuckey BGA: Small changes in 
thyroxine dosage do not produce measurable 
changes in hypothyroid symptoms, well-be-
ing, or quality of life: results of a double-
blind, randomized clinical trial. J Clin Endo-
crinol Metab 2006;   91:   2624–2630.  30 Bianco AC, Kim BW: Deiodinase: implica-tions of the local control of thyroid hormone 
action. J Clin Invest 2006;   116:   2571–2579.

---

### Chunk 25/30
**Article:** Hipotireioidismo (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.543

pendente de pH gástrico e intestino delgado; pico 1–1,5h.
- Redução por: gestação, hipocloridria (IBP), gastrite atrófica, H. pylori, doença celíaca, ferro/cálcio, soja, idade avançada, baixa massa magra, gastroparesia (diabetes; análogos GLP-1), intolerância à lactose (excipientes), parasitoses (giárdia).
- Usuários de IBP podem precisar ~37% mais dose.
- Investigar má absorção em doses >200–300 mcg: lactose, endoscopia, H. pylori, celíaca; considerar formulações líquidas/soft gel.
### 14. Epidemiologia do tratamento e sintomas persistentes
- Levotiroxina: 2ª droga mais vendida nos EUA (2019); uso aumentou 63% em 20 anos.
- 5–15% dos tratados têm sintomas persistentes; ~70–77% relatam insatisfação (2018–2020).
- Sintomas: fadiga, baixa energia, queixas cognitivas, sintomas musculoesqueléticos.
### 15.

---

### Chunk 26/30
**Article:** Hipotireioidismo (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.543

gia (ajuste de T4 ± T3), checar conversão periférica e polimorfismos quando possível.
- [ ] 11. Considerar terapia combinada T4+T3 em candidatos: sintomáticos com LT4 adequado, tireoidectomizados, alta dose de LT4, autoimunidade ativa; seguir proporção 13:1–20:1 e fracionar T3.
- [ ] 12. Evitar T3 em gestantes, cardiopatas instáveis, malignidade ativa, psiquiatria não controlada; usar T3 isolado apenas em indicações específicas.
- [ ] 13. Mapear preferências/experiências dos pacientes com T3 e decidir compartilhadamente.
- [ ] 14. Triar e manejar SIBO/disbiose em hipotireoidismo (especialmente com constipação crônica), usando teste respiratório e terapias específicas; promover saúde do microbioma (“pool” intestinal).
- [ ] 15. Em cardiopatas (ICC, pós-IAM), avaliar T3 livre e segurança para eventual ajuste terapêutico (LT4 e, em casos selecionados, T3), monitorando Pro-BNP e PCR-us.
- [ ] 16.

---

### Chunk 27/30
**Article:** Hipotireioidismo (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.541

ulam produção, conversão e sensibilidade tecidual, e revisa a epidemiologia e etiologia (com ênfase em Hashimoto como causa predominante e sistêmica).
Historicamente, mostra a transição de marcadores clínicos (taxa metabólica basal, PBI, sinais/sintomas) para a centralidade do TSH, apontando limitações dos imunoensaios e do uso exclusivo do TSH como alvo terapêutico. Discute evidências sobre terapia combinada T4/T3 versus monoterapia com T4, polimorfismos de deiodinases, perfis de absorção da levotiroxina e causas de má absorção (hipocloridria, IBP, H. pylori, doença celíaca, SIBO/disbiose), além do papel do microbioma e do “pool” intestinal de hormônios conjugados.
No campo clínico, enfatiza a necessidade de integrar dados laboratoriais com a clínica e biomarcadores teciduais, individualizando metas para restaurar função celular e qualidade de vida.

---

### Chunk 28/30
**Article:** Genética e Epigenética II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.541

a elevar os níveis, seguida de reavaliação sanguínea em dois meses para ajustar a dose de manutenção (geralmente 2.000 a 5.000 UI/dia). O monitoramento é feito com o exame de 25-hidroxivitamina D, e o PTH pode servir como marcador funcional.
### 3. A Importância do Magnésio e da Vitamina K2
- **Magnésio:** A ativação da vitamina D depende de magnésio, sendo crucial prescrevê-los em conjunto. A deficiência de magnésio é generalizada no Brasil, e o exame de sangue sérico não é um bom indicador de seu status corporal. O magnésio atua como um bloqueador natural dos canais de cálcio, sendo vital para a saúde cardiovascular (hipertensão) e para modular a excitotoxicidade no sistema nervoso (ansiedade, depressão). Recomenda-se a suplementação para todos os pacientes.
- **Vitamina K2 (MK7):** Deve ser co-prescrita com a vitamina D para ajudar a direcionar o cálcio para os ossos, otimizando a saúde óssea e cardiovascular.

---

### Chunk 29/30
**Article:** American Thyroid Association Statement on Postoperative Hypoparathyroidism: Diagnosis, Prevention, and Management in Adults (2018)
**Journal:** Thyroid
**Section:** abstract | **Similarity:** 0.540

HypoPT occurs when a low intact parathyroid hormone level accompanies hypocalcemia. Risk factors include bilateral thyroid operations, autoimmune thyroid disease, central neck dissection, and surgeon inexperience. Prevention strategies involve optimizing vitamin D, preserving parathyroid blood supply, and autotransplanting ischemic glands. A postoperative PTH level below 15 pg/mL indicates increased acute hypoPT risk. Management includes oral calcium and vitamin D supplementation with monitoring for rebound hypercalcemia.

---

### Chunk 30/30
**Article:** Hipotireioidismo (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.540

trauma, excesso de flúor, toxinas ambientais, autoimunidade.
- Sensibilidade do receptor: exercício físico aumenta sensibilidade; sedentarismo reduz.
### 6. Epidemiologia e etiologia do hipotireoidismo
- Hashimoto responde por ~90% dos casos; predominância em mulheres 20–60 anos.
- Etiologia orienta desfechos e condutas: hipotireoidismo primário, secundário (hipofisário) ou pós-cirúrgico.
### 7. Histórico do diagnóstico e tratamento
- Do mixedema a extratos animais, PBI e taxa metabólica basal como marcadores.
- Descobertas de T4 (1926), síntese (1949), T3 (1952), e conversão T4→T3 (1970).
- Introdução do TSH (1971) migrou foco para normalização laboratorial; debate sobre alvo ideal persiste.
### 8. Armadilhas diagnósticas e limitações de exames
- Premissas equivocadas: TSH como absoluto, conversão uniforme T4→T3, normalidade populacional, exclusão do T3 como perigoso, etiologia irrelevante.

---

