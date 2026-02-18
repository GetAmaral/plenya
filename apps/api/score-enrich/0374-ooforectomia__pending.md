# ScoreItem: Ooforectomia

**ID:** `019bf31d-2ef0-72f9-b93c-0f0a852d9d51`
**FullName:** Ooforectomia (Histórico de doenças - Cirurgias já realizadas - Cirurgias que interferem diretamente no escore)

**Preparation Metadata:**
- Quality Grade: **GOOD**
- Total Chunks: 30 de 19 artigos
- Avg Similarity: 0.578

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-72f9-b93c-0f0a852d9d51`.**

```json
{
  "score_item_id": "019bf31d-2ef0-72f9-b93c-0f0a852d9d51",
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

**ScoreItem:** Ooforectomia (Histórico de doenças - Cirurgias já realizadas - Cirurgias que interferem diretamente no escore)

**30 chunks de 19 artigos (avg similarity: 0.578)**

### Chunk 1/30
**Article:** Trajectories of metabolic parameters after bilateral oophorectomy in premenopausal women (2022)
**Journal:** Maturitas
**Section:** abstract | **Similarity:** 0.666

This population-based cohort study examined metabolic changes in premenopausal women undergoing bilateral ovary removal over a decade. The research compared women without surgery (270), those receiving hormone therapy (163), and those without therapy (107). The three groups had significantly different mean values of diastolic blood pressure, weight, body mass index, and cholesterol markers. Weight and BMI showed the most pronounced trajectory shifts, with changes occurring primarily in the initial 4-5 years. Women receiving estrogen therapy were comparable to referent women with respect to weight and BMI trends, and they experienced an increase in HDL-C over time, suggesting hormone replacement mitigated some adverse effects. Surgical removal of ovaries before natural menopause produces unfavorable metabolic changes possibly increasing cardiovascular risk, though hormone therapy appeared protective for certain parameters.

---

### Chunk 2/30
**Article:** Long-term health consequences of hysterectomy with bilateral oophorectomy: A systematic review (2018)
**Journal:** Maturitas
**Section:** abstract | **Similarity:** 0.627

Hysterectomy with bilateral oophorectomy before natural menopause causes an abrupt surgical menopause and long-lasting estrogen deficiency. This systematic review examined cardiovascular, metabolic, cognitive, mental health, sexual function, and skeletal effects of bilateral oophorectomy. Women who underwent bilateral oophorectomy before age 45-50 years showed increased risks of cardiovascular disease, cognitive impairment, dementia, parkinsonism, osteoporosis, and overall mortality. These risks were attenuated but not eliminated by estrogen therapy.

---

### Chunk 3/30
**Article:** Treatment of Women After Bilateral Salpingo-oophorectomy Performed Prior to Natural Menopause (2021)
**Journal:** JAMA
**Section:** abstract | **Similarity:** 0.626

This clinical insight examines treatment approaches for women undergoing surgical removal of both ovaries before natural menopause. The authors note that surgical menopause produces more rapid hormone decline than natural menopause, resulting in severe vasomotor symptoms and increased risks of mood disorders, sleep disturbances, and sexual dysfunction. Long-term health consequences include elevated rates of cardiovascular disease, cognitive decline, osteoporosis, and increased overall mortality. The article recommends menopausal hormone therapy for symptomatic relief and potential mitigation of adverse long-term effects, along with lifestyle modifications and cardiometabolic assessment.

---

### Chunk 4/30
**Article:** Hysterectomy and risk of osteoporosis: a longitudinal cohort study (2003)
**Journal:** American Journal of Epidemiology
**Section:** abstract | **Similarity:** 0.604

This longitudinal cohort study examined bone health outcomes following hysterectomy. Women who underwent hysterectomy, particularly with bilateral oophorectomy, showed accelerated bone loss and increased fracture risk compared to age-matched controls. Even hysterectomy with ovarian preservation was associated with earlier onset osteoporosis, likely due to disrupted ovarian function. The findings support routine bone density monitoring and preventive interventions including calcium, vitamin D supplementation, and consideration of hormone replacement therapy in appropriate candidates.

---

### Chunk 5/30
**Article:** Risk of cardiovascular disease after preventive salpingo-oophorectomy (2020)
**Journal:** International Journal of Gynecological Cancer
**Section:** abstract | **Similarity:** 0.603

This Norwegian cohort study compared cardiovascular disease risk between 134 women who underwent preventive oophorectomy (median age 47) and 268 age-matched reference women. The researchers used the NORRISK 2 risk assessment tool and examined cardiometabolic factors including triglycerides, C-reactive protein, BMI, and waist circumference. The findings indicated that ten year cardiovascular risk was similar in women after preventive oophorectomy and reference women (1.15% versus 1.25%, p=0.4). Women in the surgery group demonstrated lower BMI and waist circumference.

---

### Chunk 6/30
**Article:** Fisiologia Endócrina Feminina (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.601

o de risco individual antes de terapia hormonal: histórico pessoal/familiar de câncer de mama, trombose, risco cardiovascular; densidade mineral óssea.
    - Diferenciar fogachos de outras causas de “calor” (carcinoide, mastocitose, fármacos, ansiedade, etc.).
    - Considerar perfil lipídico, marcadores inflamatórios, saúde óssea (densitometria), saúde urogenital e qualidade do sono.
    - Considerar intervenções graduais na transição menopausal (reposição de progesterona, estradiol, testosterona) conforme deficiência, indicação e riscos.
    - Educação da paciente para adesão terapêutica informada e tomada de decisão compartilhada.
- Plano de Tratamento de Seguimento:
  - Mudanças de estilo de vida:
    - Atividade física regular, com ênfase em treino de resistência (~250 minutos semanais) para saúde óssea, muscular e geral.
    - Higiene do sono (priorizar sono profundo entre ~22h–5h).

---

### Chunk 7/30
**Article:** MFI - Reposição Hormonal - AULA 07 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.597

sidade e inflamação é fundamental.
*   **Modulação dos metabólitos do estrogênio (estronas)**
   - Crucíferas (brócolis, couve-flor, couve) ajudam a tornar estronas menos proliferativas; consumo moderado (≤3–4x/semana) por serem goitrogênicas.
   - Suplementação:
     - **Indol-3-carbinol (I3C):** 200–400 mg/dia; mais fraco e mais barato.
     - **Di-indolilmetano (DIM):** 100–200 mg/dia; estrutura dupla, mais potente.
*   **Acompanhamento avançado com o DUTCH Test**
   - Ideal para acompanhamento assertivo: metabolômica dos hormônios esteroides via DUTCH Test (D-U-T-C-H).
   - Permite visualizar todos os metabólitos hormonais.
   - Exame caro, pouco acessível e complexo; requer estudo prévio do profissional antes de discutir resultados com o paciente.

---

### Chunk 8/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.587

luto, evitar alarmismo, e definir critérios para exames complementares (ressonância) apenas quando houver fatores de risco adicionais.
- [ ] 4. Atualizar materiais educativos para esclarecer que história familiar, por si só, não contraindica reposição; incorporar achados do Sister Study e WHI.
- [ ] 5. Estabelecer diretriz interna: não indicar reposição hormonal sistêmica em pacientes com histórico de câncer de mama; considerar terapias tópicas para atrofia vaginal após tentativa de métodos não hormonais, com suporte emocional.
- [ ] 6. Criar protocolo de uso criterioso de gestrinona em endometriose e mastalgia refratária, com consentimento informado sobre lacunas de evidência oncológica.
- [ ] 7. Definir critérios de indicação de testosterona por motivos não oncológicos, evitando prescrição para “redução de risco mamário” até que haja validação em guidelines.
- [ ] 8.

---

### Chunk 9/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** methods | **Similarity:** 0.585

; uso de medicamentos/suplementos; hábitos de vida (alimentação, exercício, álcool, tabaco); saúde intestinal; níveis de estresse/ansiedade; rede de apoio.
   - Atenção a sintomas de predominância estrogênica: dor mamária, TPM intensa; lacunas de cuidado em pacientes com alterações fibrocísticas e cistos, frequentemente não acolhidas apesar de sintomas de excesso de estrogênio circulante.
* **Avaliação de composição corporal**
   - Preferência por métodos além da balança: bioimpedância (utilizada pela instrutora) ou densitometria de corpo total.
   - Observação de sarcopenia em mulheres com peso normal e desconhecimento da condição; necessidade de intervenção.
* **Marcadores laboratoriais e metabolicidade**
   - Solicitar marcadores inflamatórios para identificar adoecimento metabólico ou risco, e para monitorar evolução após intervenções.

---

### Chunk 10/30
**Article:** Síndrome dos Ovários Policísticos - Parte I (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.579

*
- Quase 60% das mulheres com SOP reportam insatisfação com os cuidados médicos atuais.
- O dado sinaliza falhas em educação, rastreio precoce e personalização terapêutica.
**Achados Metabólicos e de Sinalização sustentam o risco para diabetes tipo 2.**
- A SOP está associada à diabetes mellitus tipo 2, ligada à resistência à insulina e hiperinsulinemia.
- Elementos da via de sinalização da insulina (PI3K) e falhas na translocação do GLUT4 comprometem a captação de glicose, contribuindo para resistência à insulina.
**Achados Adicionais**
- 17-hidroxiprogesterona é citada no contexto de excesso androgênico e desequilíbrio com androstenediona, como possível marcador hormonal específico.
- O risco de diabetes tipo 2 está aumentado em mulheres com SOP devido ao descontrole metabólico associado.

---

### Chunk 11/30
**Article:** The Optimal Age for Oophorectomy in Women with Benign Conditions: A Narrative Review (2025)
**Journal:** Journal of Personalized Medicine
**Section:** abstract | **Similarity:** 0.575

The study examines when ovarian removal should occur during surgery for benign gynecological conditions. Researchers reviewed literature from 2000-2025 comparing outcomes between keeping versus removing ovaries. Their findings indicate that ovarian conservation in average-risk women reduces cardiovascular disease, osteoporosis, and death risks. However, removal benefits high-risk populations, particularly those with BRCA mutations facing significantly lower cancer risk. The authors conclude that surgical decisions should be personalized, balancing patient-specific factors rather than relying solely on age-based recommendations.

---

### Chunk 12/30
**Article:** Psiquiatria Metabólica Funcional Integrativa AULA 19 (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.574

ado em aula posterior com a Dra. Juliana Bica.
### 8. Avaliação e manejo prático antecipados
* Testosterona salivar e cortisol em mulheres
   - Para maior acurácia, preferir avaliação salivar de testosterona e cortisol versus sanguínea, especialmente em mulheres.
* Retirada de contraceptivos orais
   - O processo pode desencadear queda de cabelo e acne; é necessário preparo, aviso e suporte para manejar os efeitos de retirada (“tem que estar bem amparado”).
* Avaliação hormonal ampla
   - Em homens e mulheres, além de testosterona, observar estradiol e diidrotestosterona; entender metabólitos ativos e a “dança” hormonal para decisões terapêuticas equilibradas.
## ❓ Perguntas
- [Inserir Pergunta/Confusão]
## 📚 Tarefas
- [ ] 1. Ler a revisão sistemática e meta-análise recente citada pelo instrutor sobre contraceptivos orais e suas implicações (incluindo recomendação nível A para L-metilfolato).
- [ ] 2.

---

### Chunk 13/30
**Article:** Terapia de Reposição Hormonal Feminina III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.574

e administração.**
- A "janela de oportunidade" para iniciar a TRH com a melhor relação risco-benefício é para mulheres com menos de 60 anos ou nos primeiros 10 anos de menopausa.
- Para o estradiol, o objetivo é atingir níveis séricos entre 60 e 90 pg/ml para máxima proteção cardiovascular, com doses de gel variando de 0,3 a 2,5 mg e adesivos de 50 a 100 mcg.
- Para a progesterona natural micronizada, as doses variam de 100 mg (contínua) a 200-300 mg (sequencial, a cada 12-15 dias).
- Para a testosterona, o objetivo é manter os níveis fisiológicos abaixo de 100 ng/dl para evitar efeitos colaterais.
- O monitoramento inclui mamografia anual a partir dos 40 anos e avaliação da espessura endometrial, que não deve ultrapassar 10 mm.
### Achados Adicionais Chave
- O risco de mortalidade cardiovascular é seis vezes maior que o risco de mortalidade por câncer, ressaltando a importância da proteção cardiovascular oferecida pela TRH.

---

### Chunk 14/30
**Article:** MFI - Reposição Hormonal - AULA 06 (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.573

sânimo, perda de memória em homens).
  - Em mulheres com endometriose/hiperestrogenismo: considerar gestrinona vaginal antes de implantes; avaliar risco de exacerbação de sinais androgênicos, especialmente em história de acne/alopecia/ovário policístico.
  - Ética e custo: preferir testar creme vaginal antes de implantes; monitorar exames; individualizar conforme história clínica e fenótipo (oleosidade, pilificação, acne, queda de cabelo).
## Diagnóstico Principal:
- Avaliação: Não há diagnóstico clínico individual; conteúdo instrucional sem caso específico.
- Diagnóstico Suspeito: Nenhum no momento
## Plano:
- Prescrição: Inserir mais aqui
- Próximos Passos/Exames:
  - Monitorar, em contextos de terapia hormonal, estradiol, testosterona total e livre, DHT, SHBG antes, durante e após intervenções.
  - Ajustar doses de inibidores de aromatase e agentes antiandrogênicos conforme resposta clínica e laboratorial.

---

### Chunk 15/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.570

namnese detalhada abrangendo hábitos de vida, saúde intestinal, sintomas hormonais (dor mamária, TPM), história oncológica familiar (ambos os lados) e rede de apoio.
- [ ] Avaliar composição corporal com bioimpedância ou densitometria de corpo total para detectar sarcopenia e orientar intervenção.
- [ ] Solicitar e acompanhar marcadores inflamatórios/metabólicos para monitorar resposta a intervenções de estilo de vida, especialmente em pacientes com diagnóstico prévio de câncer.
- [ ] Utilizar calculadoras de risco (p. ex., Gail e ferramentas que integram PRS e fatores de estilo de vida) com cautela, reconhecendo superestimação em populações brasileiras, e empregá-las para motivar mudanças comportamentais.
- [ ] Considerar farmacogenética ao selecionar terapias, avaliando polimorfismos que alteram metabolização de drogas.
- [ ] Preparar materiais e plano para a próxima aula focada em dúvidas sobre reposição hormonal.

---

### Chunk 16/30
**Article:** Síndrome dos Ovários Policísticos - Parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.570

; em descendência asiática, ≥4.
- Exames laboratoriais para diferenciais:
  - Prolactina (hiperprolactinemia).
  - 17-OHP (HAC não clássica).
  - TSH, T4 (± T3) para disfunção tireoidiana.
  - Testosterona total/livre, DHEA-S (tumores secretores/uso exógeno).
  - USG pélvica; RM/TC se suspeita de tumores.
  - Síndrome de Cushing: cortisol salivar noturno ou teste de supressão com dexametasona 1 mg (se suspeita clínica).
- Achados clínicos gerais:
  - Irregularidade menstrual frequente; ciclos <21 dias, oligomenorreia >35 dias, amenorreia ≥3 meses ou <8 menstruações/ano.
  - Sangramento uterino anormal de causa ovulatória (não estrutural) pode ocorrer.
  - Fenótipo A (três critérios presentes) com maior risco de complicações metabólicas.

---

### Chunk 17/30
**Article:** Síndrome dos Ovários Policísticos - Parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.567

Dosar progesterona (dia 21–24; qualquer dia em amenorreia).
  - Monitorar vitamina D sérica; ajustar para 45–60 ng/mL.
  - Avaliar B12, B6 (PLP) e folato (metilfolato); considerar suplementação ativa (metilcobalamina, metilfolato, P5P) se necessário.
  - Acompanhar função hepática (esteatose) e marcadores inflamatórios; considerar polimorfismos do receptor de melatonina em casos refratários.
- Plano de Tratamento de Seguimento:
  - Monitorar resposta clínica/metabólica ao regime escolhido (COCs versus alternativas, metformina, suplementações).
  - Reavaliar efeitos colaterais dos COCs; ajustar estrogênio/progestágeno conforme perfil e impacto na insulina.
  - Considerar espironolactona temporária ao descontinuar COCs para controle de hirsutismo/acne; reavaliar periodicamente.
  - Manter intervenções de estilo de vida de longo prazo (exercícios ≥250 min/semana moderados ou 150 min intensos; dieta estruturada; sono; manejo de estresse).

---

### Chunk 18/30
**Article:** Fisiologia Endócrina Feminina (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.566

sporinas, agentes antiestrogênicos, uso prolongado de anticoncepcionais, tabagismo (↑ metabolização de estrogênios), exposição à radiação, certas drogas, menopausa precoce.
  - Janela de oportunidade terapêutica: primeiros 10 anos pós-menopausa; “janela ótima” sugerida nos 10 anos que antecedem a menopausa para iniciar intervenções.
  - História da terapia hormonal:
    - Premarin (estrógeno equino conjugado) aprovado em 1942 para fogachos; combinação com acetato de medroxiprogesterona (Prempro) no WHI (2002) associou ↑ risco relativo de câncer de mama e eventos tromboembólicos; Million Women Study (2003) com achados semelhantes.
    - Reavaliações posteriores (p.ex., 2018, Rhodes et al.) indicam nuances e potenciais efeitos neutros/protetores dependendo de via, tipo, tempo e perfil da paciente. Ênfase em reposição hormonal personalizada.

---

### Chunk 19/30
**Article:** Bayesian Meta-analysis of Hormone Therapy and Mortality in Younger Postmenopausal Women (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.564

y in the relationship to
carcinoma and cardiovascular and metabolic problems. Obstet Gynecol. 1979;54:74-79.
42. PEPI Trial Writing Group. Effects of estrogen or estrogen/progestin
regimens on heart disease risk factors in postmenopausal women. The
Postmenopausal Estrogen/Progestin Interventions (PEPI) Trial. The
Writing Group for the PEPI Trial. JAMA. 1995;273:199-208.
43. Perez-Jaraiz MD, Revilla M, Alvarez de los Heros JI, et al. Prophylaxis of osteoporosis with calcium, estrogens and/or eelcatonin: comparative longitudinal study of bone mass. Maturitas. 1996;23:327-332.
44. Ravn P, Bidstrup M, Wasnich RD, et al. Alendronate and estrogenprogestin in the long-term prevention of bone loss: four-year results
from the early postmenopausal intervention cohort study. A randomized, controlled trial. Ann Intern Med. 1999;131:935-942.
45. Watts NB, Nolan JC, Brennan JJ, Yang HM. Esterified estrogen
therapy in postmenopausal women.

---

### Chunk 20/30
**Article:** Psiquiatria Metabólica Funcional Integrativa AULA 19 (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.563

-análise recente citada pelo instrutor sobre contraceptivos orais e suas implicações (incluindo recomendação nível A para L-metilfolato).
- [ ] 2. Implementar triagem em usuárias de contraceptivos orais: solicitar folato, homocisteína, B12, B6 e avaliar necessidade de L-metilfolato preventivo.
- [ ] 3. Avaliar sintomas em pacientes usuárias de anticoncepcionais relacionados a humor, energia, cognição e densidade óssea; considerar impacto de SHBG e testosterona livre.
- [ ] 4. Planejar protocolos de retirada de contraceptivos orais com orientação prévia sobre possíveis efeitos (queda de cabelo, acne) e estratégias de manejo.
- [ ] 5. Adotar avaliação salivar de testosterona e cortisol em mulheres com sintomas compatíveis, especialmente sem uso de anticoncepcionais.
- [ ] 6. Realizar rastreio tireoidiano em pacientes com sintomas depressivos antes de firmar diagnóstico de depressão: TSH, T4, T3; considerar T3 reverso quando clinicamente pertinente.
- [ ] 7.

---

### Chunk 21/30
**Article:** Testosterone in women—the clinical significance (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.560

double-blind, placebo-controlled study. 
J Clin Endocrinol Metab 2006; 91: 1683–90.80 Popat VB, Calis KA, Kalantaridou SN, et al. Bone mineral density in young women with primary ovarian insu  ciency: results of a 
three-year randomized controlled trial of physiological transdermal 
estradiol and testosterone replacement. 
J Clin Endocrinol Metab 2014; 99: 3418–26.81 Dobs AS, Nguyen T, Pace C, Roberts CP. Di erential e ects of oral estrogen versus oral estrogen-androgen replacement therapy on 
body composition in postmenopausal women. 
J Clin Endocrinol Metab 2002; 87: 1509–16.82 Key TJ, Verkasalo PK, Banks E. Epidemiology of breast cancer. Lancet Oncol 2001; 2: 133–40.83 Davis SR. Cardiovascular and cancer safety of testosterone in women. Curr Opin Endocrinol Diabetes Obes 2011; 18: 198–203.84 Peters AA, Buchanan G, Ricciardelli C, et al. Androgen receptor inhibits estrogen receptor-alpha activity and is prognostic in breast 
cancer.

---

### Chunk 22/30
**Article:** MFI - Reposição Hormonal - AULA 06 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.559

antes, durante e após intervenções.
  - Ajustar doses de inibidores de aromatase e agentes antiandrogênicos conforme resposta clínica e laboratorial.
- Acompanhamento e Tratamento:
  - Individualizar modulação hormonal; iniciar com menor dose eficaz; reavaliar em 1–2 meses.
  - Considerar intervenções de estilo de vida: dieta anti-inflamatória, modulação intestinal em endometriose; adequar adesão e hábitos.
  - Em mulheres, ponderar uso de gestrinona vaginal e evitar implantes até confirmar tolerância; em homens, preferir anastrozol em doses baixas quando necessário e saw palmetto para manejo de DHT, sempre com acompanhamento.

---

## Teaching Note

Data e Hora: 2025-11-21 04:14:16
Local: [Inserir Local]: [Inserir Local]
Aula: [Inserir Nome da Aula]: Terapia de Reposição Hormonal com Testosterona
## Visão Geral
A aula abordou estratégias para modular a conversão de testosterona, focando em diminuir o estradiol e o DHT.

---

### Chunk 23/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.559

com mamas lipossubstituídas.
- Apesar do risco relativo elevado, o risco absoluto de uma paciente com mama densa desenvolver câncer ao longo da vida aumenta de uma base de 10% para apenas 10,6%.
- O rastreio mamográfico é geralmente iniciado a partir dos 50 anos de idade.
**A percepção de risco da reposição hormonal foi moldada por estudos observacionais, como um de 2019, mas avanços nos últimos 20 anos permitem um acompanhamento mais seguro, como a monitorização da mama a cada três meses.**
- Um estudo observacional publicado em 2019 mostrou um aumento na incidência de câncer de mama, o que gerou receio entre os médicos.
- O material complementar deste estudo, com 50 páginas, continha um resumo de ensaios clínicos randomizados que ajudavam a contextualizar os achados.
- Nos últimos 20 anos, surgiram novos estudos que melhoraram o entendimento sobre a reposição hormonal.

---

### Chunk 24/30
**Article:** Síndrome dos Ovários Policísticos - Parte I (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.559

m: obesidade, sedentarismo, má alimentação, dislipidemia, esteatose hepática, hiperandrogenismo, resistência à insulina, inflamação crônica, disbiose intestinal, estresse oxidativo, disfunção mitocondrial e exposição a desreguladores endócrinos.
2.  **Histórico de Medicação:** Inserir mais aqui
## Subjetivo:
O conteúdo é uma palestra informativa sobre a Síndrome dos Ovários Policísticos (SOP) e não um registro de um paciente específico. O texto não contém queixas subjetivas de um paciente.
## Objetivo:
O conteúdo é uma palestra informativa sobre a Síndrome dos Ovários Policísticos (SOP) e não contém achados de exames de um paciente específico.

---

### Chunk 25/30
**Article:** Terapia de Reposição Hormonal Feminina I (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.559

adesão da paciente.
*   É crucial alinhar as expectativas da paciente, informando que a melhora clínica pode levar de 2 a 3 meses.
## Diagnóstico Primário:
*   **Avaliação:** O foco principal é a abordagem e manejo da terapia de reposição hormonal (TRH) em mulheres na menopausa. A discussão enfatiza a importância de iniciar a TRH o mais próximo possível da menopausa, idealmente começando a otimização hormonal 10 anos antes (janela de otimização).
*   **Diagnóstico Suspeito:** Nenhum no momento.
## Plano:
*   **Prescrição:** [Não aplicável]
*   **Próximos Passos/Exames:**
    *   Avaliar o perfil da paciente, incluindo estilo de vida, composição corporal (bioimpedanciometria), qualidade do sono e perfil lipídico.
    *   Avaliar a função intestinal e o estroboloma.
    *   Considerar a dosagem de vitaminas e minerais essenciais para a metabolização hormonal (ex: ferro, vitamina B12).

---

### Chunk 26/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.558

dos sintomas gerais a investigar em consulta, como dor mamária e TPM intensa, que podem indicar predominância estrogênica.

## Objetivo:
Não há achados de exame físico de um paciente específico. A abordagem objetiva descrita inclui:
- Avaliação da composição corporal por bioimpedância ou densitometria de corpo total, identificando condições como sarcopenia em mulheres com peso normal.
- Solicitação de exames laboratoriais para avaliar marcadores inflamatórios e verificar se a paciente está metabolicamente doente ou em risco.

## Diagnóstico Primário:
- Avaliação: Discussão educacional sobre estratificação de risco para câncer de mama. 90% dos casos relacionam-se ao estilo de vida e 10% a fatores genéticos conhecidos. Mutações de alta penetrância (ex.: BRCA1, BRCA2, TP53) aumentam significativamente o risco, porém são raras. Epigenética e estilo de vida são cruciais e modificáveis.

---

### Chunk 27/30
**Article:** Síndrome dos Ovários Policísticos - Parte I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.557

atório com controle de carboidratos; programa de exercícios com musculação, aeróbicos e HIIT para ganho de massa muscular e melhora da sensibilidade à insulina.
- [ ] 6. Reduzir hiperinsulinemia antes de induções de ovulação em pacientes que desejam engravidar; corrigir peso/percentual de gordura e abordar estresse oxidativo.
- [ ] 7. Reavaliar uso de anticoncepcionais: considerar uso pontual quando indicado; evitar abordagem rotineira e discutir alternativas focadas nas causas.
- [ ] 8. Educar pacientes sobre os riscos e complicações do “iceberg” da SOP e sobre a natureza bidirecional dos mecanismos (insulina–desbiose–andrógenos–inflamação); promover adesão ao tratamento causal.
- [ ] 9. Estruturar atendimento multidisciplinar contínuo (médico, nutricionista, educador físico) com plano individualizado por fenótipo/subfenótipo.
- [ ] 10. Revisar e adotar, quando possível, critérios diagnósticos uniformizados para SOP na prática clínica local.

---

### Chunk 28/30
**Article:** Terapia de Reposição Hormonal Feminina I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.557

loma.
    *   Considerar a dosagem de vitaminas e minerais essenciais para a metabolização hormonal (ex: ferro, vitamina B12).
    *   Realizar reavaliações periódicas a cada três a seis meses após o início da TRH para ajustar a dose e monitorar a condição da paciente.
    *   Avaliação contínua e estratificação de riscos para garantir que as condições que justificaram o início da terapia permaneçam válidas.
*   **Plano de Tratamento de Acompanhamento:**
    *   Individualizar a terapia de reposição hormonal (TRH) escolhendo a via, a dose e o tempo corretos, por um profissional qualificado.
    *   Utilizar a menor dose eficaz para controlar os sintomas, seguindo a abordagem "comece devagar e siga devagar" (Start slow, go slow).
    *   A progesterona deve ser administrada pela via oral para garantir proteção endometrial e outros benefícios sistêmicos.

---

### Chunk 29/30
**Article:** Síndrome dos Ovários Policísticos - Parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.557

1.000–2.000 mg/dia com enhancers de biodisponibilidade.
  - Coenzima Q10: suporte à foliculogênese e resposta a indutores.
## Diagnóstico Principal:
- Avaliação: SOP em discussão, com foco em resistência à insulina, hiperandrogenismo, irregularidade menstrual, riscos e impactos de COCs, e estratégias de manejo metabólico e reprodutivo. Conteúdo educacional; sem confirmação em paciente específico.
- Diagnóstico Suspeito: Nenhum no momento.
## Plano:
- Prescrição:
  - Inserir de acordo com o caso individual.
  - Opções discutidas:
    - Metformina XR até 1.500 mg/dia (avaliar renal, B12/B6/folato).
    - Mio-inositol 2 g + ácido fólico 200 mcg, 2x/dia; considerar alfa-lactoalbumina 50 mg, 2x/dia.
    - Progesterona micronizada oral em esquema cíclico: 200 mg/dia por 10–14 dias do ciclo (ex.: dias 15–24); monitorar humor/sono.
    - Antiandrogênios em sintomas: espironolactona, finasterida; considerar Serenoa repens 400 mg em acne/hirsutismo.

---

### Chunk 30/30
**Article:** Hipertensão Arterial Sistêmica II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.556

-   Medição da lipoproteína (a).
    -   Avaliação da relação ApoB/ApoA.
    -   Avaliação dos níveis hormonais (testosterona, estradiol, DHEA).
-   **Plano de Tratamento de Acompanhamento:**
    -   O plano de tratamento é conceitual, focado em abordar os fatores de risco identificados:
    -   Suplementação para corrigir deficiências (ex: Ômega-3, Vitamina D, complexo B para homocisteína).
    -   Manejo da resistência à insulina através de dieta (com apoio de nutricionista), estilo de vida e medicamentos como metformina.
    -   Terapia de reposição hormonal (estrogênio, testosterona) quando indicado, para proteção cardiovascular.
    -   Uso de novas terapias como análogos de GLP-1 (ex: semaglutida) para obesidade e insuficiência cardíaca, e medicamentos para reduzir a lipoproteína (a) (ex: lepodisiran).

---

