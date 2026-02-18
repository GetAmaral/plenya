# ScoreItem: USG Próstata - Densidade PSA (PSAD)

**ID:** `019bf31d-2ef0-7f97-8e14-799354166f5e`
**FullName:** USG Próstata - Densidade PSA (PSAD) (Exames - Imagem)
**Unit:** ng/mL/cm³

**Preparation Metadata:**
- Quality Grade: **GOOD**
- Total Chunks: 30 de 14 artigos
- Avg Similarity: 0.598

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-7f97-8e14-799354166f5e`.**

```json
{
  "score_item_id": "019bf31d-2ef0-7f97-8e14-799354166f5e",
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

**ScoreItem:** USG Próstata - Densidade PSA (PSAD) (Exames - Imagem)
**Unidade:** ng/mL/cm³

**30 chunks de 14 artigos (avg similarity: 0.598)**

### Chunk 1/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.719

Específico (PSA)**:
    *   **Função**: Enzima que liquefaz o sêmen.
    *   **Formas**: Complexado (maior parte) e Livre. O PSA total é a soma de ambas.
    *   **Interpretação Clínica**: A relação PSA livre / PSA total é crucial.
        *   **> 0.14 (ou 14%)**: Sugere HPB.
        *   **< 0.14 (ou 14%)**: Aumenta o risco de câncer de próstata.
    *   **Limitações**: Cerca de 1-4% dos cânceres de próstata ocorrem com PSA normal. Em homens com baixa testosterona, esse número pode chegar a 15%.
*   **Exames de Imagem**:
    *   **Ultrassonografia de Próstata com Estudo do Resíduo Pós-Miccional**: Avalia anatomia e função. Um resíduo pós-miccional > 40 ml indica obstrução.
    *   **Urofluxometria**: Indicada para sintomas obstrutivos. Mede a velocidade do fluxo urinário (Qmax). Valores < 10 ml/segundo indicam obstrução.
    *   **Ressonância Magnética Multiparamétrica 3-Tesla (3T)**: Exame de alta especificidade.

---

### Chunk 2/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.712

nsiderado alterado (normal > 15 ml/s).
- O tratamento medicamentoso para HPB inclui doses como 0,4 mg de Tansulosina ou 2 a 4 mg de Doxazosina para relaxar a musculatura e melhorar o fluxo urinário.
### Achados Adicionais Chave
- A ressonância magnética multiparamétrica 3 Tesla é uma ferramenta de alta definição, recomendada a partir dos 40 anos para homens com histórico familiar ou a partir dos 50 anos como rotina para uma avaliação prostática precisa.
- A dosagem hormonal salivar oferece uma medição precisa dos hormônios livres, com faixas de referência para testosterona (47-150), estradiol (0.6-3) e o quociente estrogênico (0.04-1.67), que avalia o equilíbrio hormonal.
- A escala Gleason, que vai de 1 a 10, mede a agressividade do câncer de próstata, com tumores classificados como 8, 9 ou 10 sendo considerados os mais agressivos.
- A hiperplasia prostática é uma condição que afeta apenas 3 espécies: humanos, cães e macacos.

---

### Chunk 3/30
**Article:** Optimal PSA density threshold for prostate biopsy in benign prostatic obstruction patients with elevated PSA levels but negative MRI findings (2025)
**Journal:** BMC Urology
**Section:** abstract | **Similarity:** 0.708

Study identifying optimal PSAD cutoff of 0.30 ng/ml/cm³ for biopsy decision in BPH patients with elevated PSA but negative MRI, demonstrating 93% specificity and 65% sensitivity. ROC analysis showed PSAD achieved AUC 0.848, outperforming PSA alone (0.722) and free/total PSA ratio (0.635).

---

### Chunk 4/30
**Article:** The use of prostate specific antigen density to predict clinically significant prostate cancer (2020)
**Journal:** Scientific Reports
**Section:** abstract | **Similarity:** 0.697

Evaluated 992 men undergoing biopsy, finding PSAD AUC of 0.78 vs PSA AUC of 0.64 for predicting clinically significant cancer. Key thresholds: PSAD <0.09 ng/ml² only 4% risk; PSAD 0.09-0.19 ng/ml² risk increases with smaller prostates; PSAD ≥0.20 ng/ml² optimal cutoff with 70% sensitivity and 79% specificity.

---

### Chunk 5/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.693

tático Específico (PSA)
- **Função e Formas:** O PSA é uma enzima que liquefaz o sêmen. Existe na forma livre e complexada (ligada a proteínas).
- **Relação PSA Livre/Total:** É um cálculo crucial (PSA livre / PSA total).
    - **> 0.14 (ou 0.15):** Sugere HPB.
    - **< 0.14 (ou 0.15):** Aumenta a suspeita de câncer de próstata.
- **PSA Normal e Risco:** Um PSA normal não exclui o risco de câncer, especialmente em homens com deficiência de testosterona, onde a incidência de câncer com PSA normal pode chegar a 15%.
#### 6.2. Exames Subsidiários
- **Ultrassonografia de Próstata com Resíduo Pós-Miccional:** Avalia a anatomia (volume, textura) e a função de esvaziamento da bexiga. Um resíduo acima de 40 ml indica obstrução.
- **Urofluxometria:** Mede a velocidade do jato urinário. Indicado para pacientes com sintomas obstrutivos. Um fluxo máximo (Qmax) abaixo de 10 ml/s é considerado alterado.

---

### Chunk 6/30
**Article:** Integrating PSA Change with PSA Density Enhances Diagnostic Accuracy and Helps Avoid Unnecessary Prostate Biopsies (2025)
**Journal:** Diagnostics (Basel)
**Section:** abstract | **Similarity:** 0.689

Demonstrates that PSA density shows superior diagnostic performance (AUC 0.77-0.81) compared to PSA change alone. Combining both metrics provides optimal results, with >20% PSA decline criterion improving PSAD performance, especially valuable in prostates >80 mL where PSAD accuracy decreases.

---

### Chunk 7/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.652

s a prostatectomia radical foram drasticamente reduzidas: de uma espera de 5 a 10 anos para apenas 3 a 6 meses com PSA indetectável.
**A avaliação da saúde prostática vai além do PSA total, utilizando a relação PSA livre/total como um indicador crucial, onde valores abaixo de 0.14 sugerem maior risco de câncer, enquanto valores acima indicam HPB.**
- O PSA total é composto por uma forma livre (10-30%) e uma complexada (70-90%), e a relação entre elas é diagnóstica.
- A relação PSA livre/total tem um ponto de equilíbrio em 0.14: valores maiores (ex: 0.20, 0.30) são sugestivos de HPB, enquanto valores menores (ex: 0.09, 0.04) aumentam a suspeita de câncer de próstata.
- A concentração normal de PSA total varia de 2.5 a 4.0 ng/ml, mas um aumento de 1.0 ng/ml em um ano ou saltos abruptos (ex: de 2.5 para 4 ou 5 em dois meses) exigem investigação, mesmo dentro da faixa de normalidade.

---

### Chunk 8/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.638

viventes de câncer, realiza anamnese completa, exames físicos, aguarda um ano de PSA indetectável, solicita ressonância magnética de pelve, PET-CT prévio e exames laboratoriais completos.
*   **TRT e Hiperplasia Prostática Benigna (HPB)**:
    *   A TRT não é contraindicada para homens com HPB, a menos que haja descompensação do sistema urinário.
    *   Um aumento mínimo e esperado do volume prostático e do PSA (ex: 0,2 a 0,4) é fisiológico. Saltos abruptos (mais de 1 ponto) exigem investigação.
    *   Para tratar a HPB, sugere-se Tansulosina ou Doxazosina. O uso de Dutasterida e Finasterida é fortemente desaconselhado devido a efeitos colaterais severos (Síndrome Pós-Finasterida).
### 4. Ferramentas Diagnósticas para a Saúde Prostática
*   **Antígeno Prostático Específico (PSA)**:
    *   **Função**: Enzima que liquefaz o sêmen.
    *   **Formas**: Complexado (maior parte) e Livre. O PSA total é a soma de ambas.

---

### Chunk 9/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.633

saltos abruptos (ex: de 2.5 para 4 ou 5 em dois meses) exigem investigação, mesmo dentro da faixa de normalidade.
**O diagnóstico e manejo da Hiperplasia Prostática Benigna (HPB) dependem mais dos sintomas obstrutivos, como resíduo pós-miccional acima de 40 ml, do que do tamanho da próstata, que pode variar de 25 a mais de 80 gramas sem necessariamente causar problemas.**
- A HPB é comum a partir dos 45-50 anos, mas o tamanho da próstata (normalmente 25-30 gramas) não se correlaciona diretamente com a obstrução; próstatas de 28-29 gramas podem ser obstrutivas, enquanto outras de 70-80 gramas não.
- Um indicador chave de obstrução é o resíduo pós-miccional, com volumes acima de 40 ml sendo anormais, e a urofluxometria, onde um fluxo máximo (Qmax) abaixo de 10 ml/s é considerado alterado (normal > 15 ml/s).
- O tratamento medicamentoso para HPB inclui doses como 0,4 mg de Tansulosina ou 2 a 4 mg de Doxazosina para relaxar a musculatura e melhorar o fluxo urinário.

---

### Chunk 10/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.631

urinário (Qmax). Valores < 10 ml/segundo indicam obstrução.
    *   **Ressonância Magnética Multiparamétrica 3-Tesla (3T)**: Exame de alta especificidade. O palestrante solicita de rotina para homens > 50 anos, ou > 40 anos com histórico familiar ou alterações súbitas no PSA.
*   **Dosagem Hormonal Salivar e Quociente Estrogênico**:
    *   **Vantagens da Saliva**: Via não invasiva que mede a fração livre e 100% bioativa dos hormônios (Testosterona, DHT, Estradiol, etc.). Útil quando a clínica do paciente não corresponde aos exames de sangue.
    *   **Quociente Estrogênico**: Fórmula para avaliar o risco de doenças prostáticas.
        *   **Fórmula**: Estriol / (Estradiol + Estrona).
        *   **Valores > 1**: Bom prognóstico (perfil estrogênico protetor).
        *   **Valores < 1**: Mau prognóstico (prevalência de estrogênios proliferativos), sugerindo a necessidade de trabalhar a metilação da estrona.

---

### Chunk 11/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.621

e Doxazosina, e estar ciente da Síndrome Pós-Finasterida para preservar os níveis de DHT.
*   [ ] 4. Para médicos: Ao monitorar pacientes em TRT, entender que pequenos aumentos no PSA podem ser fisiológicos, mas investigar saltos abruptos de mais de um ponto percentual.
*   [ ] 5. Ao avaliar um paciente, calcular a relação PSA livre sobre PSA total para diferenciar risco de HPB e câncer de próstata.
*   [ ] 6. Considerar a solicitação de ressonância magnética 3-Tesla para homens com mais de 50 anos, ou com mais de 40 anos se houver história familiar de câncer de próstata ou alterações significativas no PSA.
*   [ ] 7. Em casos de dissociação entre a clínica do paciente e os exames de sangue, considerar a dosagem hormonal salivar para avaliar as frações livres e bioativas.
*   [ ] 8.

---

### Chunk 12/30
**Article:** Usefulness of free PSA ratio to enhance detection of clinically significant prostate cancer in patients with PI-RADS<3 and PSA≤10 (2024)
**Journal:** Prostate International
**Section:** abstract | **Similarity:** 0.600

Study evaluating free PSA ratio for detecting clinically significant prostate cancer. At cutoff 17.6%, achieved sensitivity 86.5% and specificity 63.7% (AUC 0.757). Clinically significant cancer found in 34% with %fPSA <17.6% vs only 4% with %fPSA ≥17.6%.

---

### Chunk 13/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.598

- A HPB não é uma contraindicação para a TRT, exceto em casos de obstrução urinária severa e descompensada.
- A TRT pode causar um aumento mínimo e esperado do volume prostático e do PSA (ex: 0.2-0.4), o que é considerado normal. Um salto abrupto do PSA (>1.0 ponto em um ano) exige investigação.
- O volume da próstata não se correlaciona diretamente com os sintomas; o mais importante é a direção do crescimento (para dentro ou para fora do canal urinário).
- **Tratamento da HPB:** Medicamentos como Tansulosina e Doxazosina são indicados. Inibidores da 5-alfa-redutase (Finasterida e Dutasterida) foram proscritos pelo palestrante devido ao risco de efeitos colaterais severos (síndrome pós-finasterida).
### 6. Avaliação Diagnóstica da Próstata
#### 6.1. Antígeno Prostático Específico (PSA)
- **Função e Formas:** O PSA é uma enzima que liquefaz o sêmen. Existe na forma livre e complexada (ligada a proteínas).

---

### Chunk 14/30
**Article:** Using the Free-to-total Prostate-specific Antigen Ratio to Detect Prostate Cancer in Men with Nonspecific Elevations of Prostate-specific Antigen Levels (2000)
**Journal:** Journal of General Internal Medicine
**Section:** abstract | **Similarity:** 0.596

Meta-analysis of 21 studies on free/total PSA ratio in PSA 4.0-10.0 ng/ml gray zone. Median likelihood ratio positive 1.76, negative 0.27. At 25% pretest probability, negative test reduced posttest probability to 8%. Modest discriminating power (AUC 0.68).

---

### Chunk 15/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.587

a um risco aumentado de câncer, enquanto a terapia de reposição hormonal se mostra segura e não agrava a hiperplasia prostática benigna (HPB). A avaliação prostática evoluiu para um modelo mais preciso, utilizando a relação PSA livre/total, o monitoramento do resíduo pós-miccional e a ressonância magnética 3 Tesla para um diagnóstico e acompanhamento mais eficazes.
---
### Evidências Chave
**A crença histórica de que a testosterona causa câncer de próstata, baseada em um estudo de 1941 com apenas 2 pacientes, foi desmentida por estudos modernos que mostram uma taxa de câncer de 14% a 18% em homens com baixa testosterona, muito acima da média populacional de 1% a 4,5%.**
- O dogma que associou testosterona ao câncer de próstata perdurou por cerca de 75 a 80 anos, originado de um estudo seminal de 1941 que, no final, contou com apenas 2 pacientes humanos.

---

### Chunk 16/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.587

jato urinário. Indicado para pacientes com sintomas obstrutivos. Um fluxo máximo (Qmax) abaixo de 10 ml/s é considerado alterado.
- **Ressonância Magnética Multiparamétrica 3-Tesla (3T):** Exame de alta definição para detecção, estadiamento e acompanhamento de câncer de próstata, além de avaliação de casos com PSA elevado e biópsias negativas.
### 7. Avaliação Hormonal Avançada
#### 7.1. Dosagem Hormonal Salivar
- É uma ferramenta diagnóstica não invasiva que mede a fração livre e bioativa dos hormônios.
- É particularmente útil quando o quadro clínico do paciente não corresponde aos exames de sangue.
#### 7.2. Hormônios e Conduta Clínica
- **DHT (Diidrotestosterona):** Essencial para a função sexual. O palestrante desaconselha o uso de inibidores da 5-alfa-redutase que bloqueiam sua produção.
- **Estradiol:** Níveis elevados são antagônicos à testosterona e podem anular seus benefícios.

---

### Chunk 17/30
**Article:** Actual Contribution of Free to Total PSA Ratio in Prostate Diseases Differentiation (2016)
**Journal:** Medical Archives
**Section:** abstract | **Similarity:** 0.576

Study of 220 patients examining %fPSA cutoffs. At ≤16%: 72.3% sensitivity, 50.4% specificity. At <7%: 8.4% sensitivity, 97.8% specificity. Cancer patients had significantly lower %fPSA than benign prostatic hyperplasia patients.

---

### Chunk 18/30
**Article:** Management of Lower Urinary Tract Symptoms Attributed to Benign Prostatic Hyperplasia (BPH): AUA Guideline Amendment 2023 (2023)
**Journal:** Journal of Urology
**Section:** abstract | **Similarity:** 0.563

Define prostatic enlargement as volume >30g on imaging. Recommends prostate size assessment via ultrasound prior to intervention. Combination therapy (5-ARI + alpha blocker) for prostates >30g.

---

### Chunk 19/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.558

*   **Valores < 1**: Mau prognóstico (prevalência de estrogênios proliferativos), sugerindo a necessidade de trabalhar a metilação da estrona.
## ❓ Perguntas
*   [Inserir Pergunta/Dúvida]
## 📚 Tarefas
*   [ ] 1. Prestar atenção à explicação sobre a origem histórica da associação entre testosterona e câncer de próstata.
*   [ ] 2. Para médicos: Ao avaliar um paciente para TRT, especialmente um sobrevivente de câncer de próstata, seguir um protocolo rigoroso que inclua anamnese completa, exames físicos, radiológicos (Ressonância Magnética, PET-CT) e laboratoriais, aguardando um período seguro de PSA indetectável.
*   [ ] 3. Para médicos: Evitar a prescrição de Dutasterida e Finasterida para o tratamento de HPB, considerando alternativas como Tansulosina e Doxazosina, e estar ciente da Síndrome Pós-Finasterida para preservar os níveis de DHT.
*   [ ] 4.

---

### Chunk 20/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.554

entes com PSA normal, contra 1-4% na população geral com PSA normal.
**A terapia de reposição de testosterona (TRT) é segura, não aumenta o risco de câncer de próstata e não piora os sintomas da Hiperplasia Prostática Benigna (HPB), com diretrizes modernas permitindo seu início apenas 3 a 6 meses após a prostatectomia, em vez dos antigos 5 a 10 anos.**
- Estudos mostraram que a TRT não "explode" o câncer; em um grupo de 20 pacientes com lesão pré-cancerígena (PIM), apenas 1 (5%) desenvolveu câncer após a terapia.
- Outro estudo com 20 anos de acompanhamento não mostrou aumento no risco de câncer de próstata em usuários de TRT.
- Para pacientes com HPB, um estudo de 2012 demonstrou que a TRT não exacerba os sintomas urinários.
- As diretrizes para iniciar a TRT após a prostatectomia radical foram drasticamente reduzidas: de uma espera de 5 a 10 anos para apenas 3 a 6 meses com PSA indetectável.

---

### Chunk 21/30
**Article:** MFI - Reposição Hormonal - AULA 11 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.548

humor e à saúde óssea.
- 5α-redutase: com Nebido, intervenção raramente necessária; com cipionato, maior probabilidade de precisar saw palmetto 300 mg ou, seletivamente, finasterida/dutasterida conforme sintomas/risco-benefício.
- Pigeum africanum: potencial benefício prostático leve; atenção a variações de PSA na TRT.
### 6. Encaminhamento e limites de atuação em casos urológicos
- Limitar atuação em câncer de próstata e quadros complexos; encaminhar a urologistas (preferencialmente com visão integrativa).
- Em PSA elevado e próstata aumentada, avaliar conjuntamente com urologia; usar rede multiprofissional e trabalhar nos limites de competência.
### 7. Modulação de estrogênios com crucíferas (I3C e DIM)
- Indol-3-carbinol (200–400 mg/dia) e diindolilmetano (100–200 mg/dia) podem modular vias de metabolização de estrogênios (perfil de estronas).

---

### Chunk 22/30
**Article:** Efficacy and Safety of Testosterone Replacement Therapy in Men with Hypogonadism: A Meta-Analysis (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.548

uding the prostate‑specific antigen (PSA) level 
(in 
ng
/
ml)
, International Prostate Symptom Scores (IPSS), 
prostate volume (in ml) and the nature of adverse events (mild 
to
 moderate and serious). Data extraction was independently 
conducted by 2 investigators, and a third reviewer would make 
a judgment when disagreement arose regarding eligibility, as 
described previously (24).
Statistical analysis. 
Statistical analyses were independently 
performed by 
2 authors who were not involved in data extraction. 
Q‑test was used to measure inter‑study heterogeneity. I
2
 metric 
was used to quantify heterogeneity, which is independent of 
the number of studies included in the cumulative analysis. The 
I
2
 values were 0‑100%, with higher values denoting a greater 
degree of heterogeneity. I
2
<25% reflects a small level of inconsis
-
tency, and I
2
>50% reflects significant inconsistency. Data were 
pooled using a fixed‑effect
s model.

---

### Chunk 23/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.546

is para a TRT, protocolos de tratamento, condutas clínicas e o acompanhamento de segurança. Além disso, a sessão cobriu a interpretação do PSA, exames subsidiários como a ressonância magnética 3-Tesla, e a utilidade da dosagem hormonal salivar para uma avaliação clínica completa.
## Conteúdo Abordado
### 1. Introdução e Filosofia da Medicina Preventiva
- **Palestrante e Tema:** O Dr. Wilson Dalla Pasqua Júnior, urologista, andrologista, cirurgião geral e nutrólogo, introduziu o tema da TRT, destacando sua relevância prática para profissionais que trabalham com reposição hormonal e saúde masculina.
- **Abreviações:** Foram definidas as abreviações CAP (Câncer de Próstata) e HPB (Hiperplasia Prostática Benigna).
- **Filosofia Médica:** Citando William James Foley, fundador da Mayo Clinic, a aula foi fundamentada no princípio da medicina preventiva, que busca evitar que a doença se estabeleça.

---

### Chunk 24/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.538

s modernas, lideradas por pesquisadores como Abraham Morgenthaler, que demonstram uma relação inversa: baixos níveis de testosterona estão associados a um maior risco de câncer de próstata agressivo. A palestra detalha as diretrizes atuais para a TRT em sobreviventes de câncer e em pacientes com HPB, concluindo que o tratamento é seguro e benéfico na maioria dos casos, com exceção de cânceres metastáticos ativos. Além disso, são discutidas ferramentas diagnósticas cruciais para a saúde prostática, como a interpretação do Antígeno Prostático Específico (PSA) e sua relação livre/total, exames de imagem como a ressonância magnética 3-Tesla, e a dosagem hormonal salivar para avaliar as frações bioativas de hormônios e o risco prostático através do quociente estrogênico.
## 🔖 Pontos de Conhecimento
### 1.

---

### Chunk 25/30
**Article:** Efficacy and Safety of Testosterone Replacement Therapy in Men with Hypogonadism: A Meta-Analysis (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.535

s in hypogonadal men 
with type 2 diabetes but not in men with coexisting depression: 
The BLAST study. J Sex Med
 11: 840‑856, 2014.
29. 
Roehrborn CG, Siami P, Barkin J
, 
Damião R, Major‑Walker K, 
Nandy I, Morrill BB, Gagnier RP and Montorsi F; CombAT 
Study Group: The effects of combination therapy with dutasteride 
and tamsulosin on clinical outcomes in men with symptomatic 
benign prostatic hyperplasia: 4‑year results from the CombAT 
study. Eur Urol 57: 123‑131, 2010.
30. 
Kaplan SA: Re: The effects of combination therapy with 
dutasteride and tamsulosin on clinical 
outcomes in men with 
symptomatic benign prostatic hyperplasia: 4‑year results from 
the CombAT study. J Urol 185: 1384‑1385, 2011.
31. 
Hsing AW, Reichardt JK and Stanczyk FZ: Hormones and 
prostate cancer: Current perspectives and future directions. 
Prostate 52: 213‑235, 2002.
32.

---

### Chunk 26/30
**Article:** MFI - Reposição Hormonal - AULA 11 (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.531

exames pontuais e relato clínico.
### 7. Saúde prostática, PSA e adjuvantes
- PSA pode subir sem relevância clínica direta em TRT; interpretar no contexto.
- Fitoterápicos: Pygeum africanum pode ajudar em sintomas prostáticos e reduzir discretamente 5-AR.
- Câncer de próstata: há evidências contextuais de benefício potencial de TRT em tipos específicos; manejo reservado a experts.
### 8. Limites profissionais, DUT test e modulação de estrógenos
- Trabalho em equipe e encaminhamento a urologistas funcionais integrativos quando necessário (PSA alto, próstata aumentada).
- DUT test: metabolômica urinária de hormônios sexuais e curva de cortisol; útil, dinâmico, mas não substitui avaliação clínica.
- Modulação de estrógenos:
  - I3C 200–400 mg/dia e DIM 100–200 mg/dia para ajustar vias de estronas conforme DUT e quadro clínico.
  - Evitar dogmas de “valores ideais” fixos; tratar a pessoa, não apenas números.
### 9.

---

### Chunk 27/30
**Article:** Can men with prostates sized 80 mL or larger be managed conservatively? (2017)
**Journal:** Canadian Urological Association Journal
**Section:** abstract | **Similarity:** 0.525

Two-thirds of men with prostates ≥80mL maintained stability with conservative management over 62 months. Clinical progression occurred in 33%. No baseline volume predicted progression.

---

### Chunk 28/30
**Article:** MFI - Reposição Hormonal - AULA 05 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.523

não piora os sintomas do trato urinário inferior em homens com HPB.
*   **Mecanismo de Saturação da Próstata**
    - A próstata possui um número limitado de receptores de testosterona, criando um efeito de "saturação".
    - A reposição de testosterona em níveis adequados "irriga" a próstata até seu estado normal, mas não causa um estímulo de crescimento descontrolado.
    - Um aumento do PSA pode ocorrer devido a esse estímulo inicial, mas não está correlacionado com o desenvolvimento de câncer. O acompanhamento deve ser feito com exames de toque e imagem (ex: ressonância 3D Tesla).
*   **Uso de TRT em Pacientes com Histórico de Câncer de Próstata**
    - Estudos recentes sugerem que a TRT pode ser utilizada em certos tipos de câncer de próstata, durante ou após o tratamento.
    - O instrutor aconselha cautela, não ensina a prática por ser polêmica e recomenda encaminhar esses casos a urologistas com visão funcional/integrativa.
### 3.

---

### Chunk 29/30
**Article:** Lifestyle Medicine: A Brief Review of Its Dramatic Impact on Health and Survival (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.522

2013). 
Prostate Cancer Prostatic Dis 2016 Dec;19(4):395-7. 
DOI: https://doi.org/10.1038/pcan.2016.30.
 96. Carter HB, Albertsen PC, Barry MJ, et al. Early 
detection of prostate cancer: AUA guideline. J 
Urol 2013 Aug;190(2):419-26. DOI: https://doi.
org/10.1016/j.juro.2013.04.119.
 97. Freedland SJ, Aronson WJ. Examining the relationship 
between obesity and prostate cancer. Rev Urol 2004 
Spring;6(2):73-81.
 98. Ornish D, Weidner G, Fair WR, et al. Intensive lifestyle 
changes may affect the progression of prostate cancer. 
J Urol 2005 Sep;174(3):1065-70. DOI: https://doi.
org/10.1097/01.ju.0000169487.49018.73.
 99. Ornish D, Magbanua MJ, Weidner G, et al. Changes 
in prostate gene expression in men undergoing an 
intensive nutrition and lifestyle intervention. Proc Natl 
Acad Sci U S A 2008 Jun 17;105(24):8369-74. DOI: 
https://doi.org/10.1073/pnas.0803080105.
 �100. Yang M, Kenfield SA, Van Blarigan EL, et al.

---

### Chunk 30/30
**Article:** Global trends of cancer: The role of diet, lifestyle, and environmental factors (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.520

0.3390/nu807041936.HyunJS.Prostatecancerandsexualfunction.WorldJMensHealth.2012;30(2):99–107.https://doi.org/10.5534/wjmh.2012.30.2.9937.JianZ,YeD,ChenY,LiH,WangK.Sexualactivityandriskofprostatecancer:adoseresponsemetaanalysis.JSexMed.2018;15(9):1300–9.https://doi.org/10.1016/j.jsxm.2018.07.00438.RapaportL.Doessexprotectmenagainstprostatecancer?ReutersHealthNews.Availablefrom:https://www.reuters.com/article/us-health-ejaculation-prostate-cancer-idUSKCN0
XJ1YC(2016).Accessed22April2016.39.MarignolL,CoffeyM,LawlerM,HollywoodD.Hypoxiainprostatecancer:apowerfulshieldagainsttumourdestruction?CancerTreatRev.2008;34(4):313–27.https://doi.org/10.1016/j.ctrv.2008.01.00640.SinghO,BollaSR.Anatomyabdomenandpelvisprostate[Internet].TreasureIsland(FL):StatPearlsPublishing;2021[updated2021Jul26;cited2021].Availablefrom:https://www.ncbi.nlm.nih.gov/books/NBK540987/.Accessed26Jul2021.41.AhmadF,CherukuriMK,ChoykePL.Metabolicreprogram-minginprostatecancer.BrJCancer.2021;125(9):1185–96.https://d

---

