# ScoreItem: USG Próstata - Volume Prostático

**ID:** `019bf31d-2ef0-7a1d-95f4-cfe9183295e0`
**FullName:** USG Próstata - Volume Prostático (Exames - Imagem)
**Unit:** cm³

**Preparation Metadata:**
- Quality Grade: **EXCELLENT**
- Total Chunks: 30 de 11 artigos
- Avg Similarity: 0.608

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-7a1d-95f4-cfe9183295e0`.**

```json
{
  "score_item_id": "019bf31d-2ef0-7a1d-95f4-cfe9183295e0",
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

**ScoreItem:** USG Próstata - Volume Prostático (Exames - Imagem)
**Unidade:** cm³

**30 chunks de 11 artigos (avg similarity: 0.608)**

### Chunk 1/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.738

nsiderado alterado (normal > 15 ml/s).
- O tratamento medicamentoso para HPB inclui doses como 0,4 mg de Tansulosina ou 2 a 4 mg de Doxazosina para relaxar a musculatura e melhorar o fluxo urinário.
### Achados Adicionais Chave
- A ressonância magnética multiparamétrica 3 Tesla é uma ferramenta de alta definição, recomendada a partir dos 40 anos para homens com histórico familiar ou a partir dos 50 anos como rotina para uma avaliação prostática precisa.
- A dosagem hormonal salivar oferece uma medição precisa dos hormônios livres, com faixas de referência para testosterona (47-150), estradiol (0.6-3) e o quociente estrogênico (0.04-1.67), que avalia o equilíbrio hormonal.
- A escala Gleason, que vai de 1 a 10, mede a agressividade do câncer de próstata, com tumores classificados como 8, 9 ou 10 sendo considerados os mais agressivos.
- A hiperplasia prostática é uma condição que afeta apenas 3 espécies: humanos, cães e macacos.

---

### Chunk 2/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.724

saltos abruptos (ex: de 2.5 para 4 ou 5 em dois meses) exigem investigação, mesmo dentro da faixa de normalidade.
**O diagnóstico e manejo da Hiperplasia Prostática Benigna (HPB) dependem mais dos sintomas obstrutivos, como resíduo pós-miccional acima de 40 ml, do que do tamanho da próstata, que pode variar de 25 a mais de 80 gramas sem necessariamente causar problemas.**
- A HPB é comum a partir dos 45-50 anos, mas o tamanho da próstata (normalmente 25-30 gramas) não se correlaciona diretamente com a obstrução; próstatas de 28-29 gramas podem ser obstrutivas, enquanto outras de 70-80 gramas não.
- Um indicador chave de obstrução é o resíduo pós-miccional, com volumes acima de 40 ml sendo anormais, e a urofluxometria, onde um fluxo máximo (Qmax) abaixo de 10 ml/s é considerado alterado (normal > 15 ml/s).
- O tratamento medicamentoso para HPB inclui doses como 0,4 mg de Tansulosina ou 2 a 4 mg de Doxazosina para relaxar a musculatura e melhorar o fluxo urinário.

---

### Chunk 3/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.707

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

### Chunk 4/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.698

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

### Chunk 5/30
**Article:** Management of Lower Urinary Tract Symptoms Attributed to Benign Prostatic Hyperplasia (BPH): AUA Guideline Amendment 2023 (2023)
**Journal:** Journal of Urology
**Section:** abstract | **Similarity:** 0.682

Define prostatic enlargement as volume >30g on imaging. Recommends prostate size assessment via ultrasound prior to intervention. Combination therapy (5-ARI + alpha blocker) for prostates >30g.

---

### Chunk 6/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.672

- A HPB não é uma contraindicação para a TRT, exceto em casos de obstrução urinária severa e descompensada.
- A TRT pode causar um aumento mínimo e esperado do volume prostático e do PSA (ex: 0.2-0.4), o que é considerado normal. Um salto abrupto do PSA (>1.0 ponto em um ano) exige investigação.
- O volume da próstata não se correlaciona diretamente com os sintomas; o mais importante é a direção do crescimento (para dentro ou para fora do canal urinário).
- **Tratamento da HPB:** Medicamentos como Tansulosina e Doxazosina são indicados. Inibidores da 5-alfa-redutase (Finasterida e Dutasterida) foram proscritos pelo palestrante devido ao risco de efeitos colaterais severos (síndrome pós-finasterida).
### 6. Avaliação Diagnóstica da Próstata
#### 6.1. Antígeno Prostático Específico (PSA)
- **Função e Formas:** O PSA é uma enzima que liquefaz o sêmen. Existe na forma livre e complexada (ligada a proteínas).

---

### Chunk 7/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.660

urinário (Qmax). Valores < 10 ml/segundo indicam obstrução.
    *   **Ressonância Magnética Multiparamétrica 3-Tesla (3T)**: Exame de alta especificidade. O palestrante solicita de rotina para homens > 50 anos, ou > 40 anos com histórico familiar ou alterações súbitas no PSA.
*   **Dosagem Hormonal Salivar e Quociente Estrogênico**:
    *   **Vantagens da Saliva**: Via não invasiva que mede a fração livre e 100% bioativa dos hormônios (Testosterona, DHT, Estradiol, etc.). Útil quando a clínica do paciente não corresponde aos exames de sangue.
    *   **Quociente Estrogênico**: Fórmula para avaliar o risco de doenças prostáticas.
        *   **Fórmula**: Estriol / (Estradiol + Estrona).
        *   **Valores > 1**: Bom prognóstico (perfil estrogênico protetor).
        *   **Valores < 1**: Mau prognóstico (prevalência de estrogênios proliferativos), sugerindo a necessidade de trabalhar a metilação da estrona.

---

### Chunk 8/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.654

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
**Article:** Age-stratified normal values for prostate volume, PSA, maximum urinary flow rate, IPSS, and other LUTS/BPH indicators in the German male community-dwelling population aged 50 years or older (2011)
**Journal:** World Journal of Urology
**Section:** abstract | **Similarity:** 0.633

Normal prostate volume ranges 20-30 mL. Volume increases with age. Classification: mild enlargement 30-50mL, moderate 50-70mL, marked >70mL.

---

### Chunk 10/30
**Article:** Can men with prostates sized 80 mL or larger be managed conservatively? (2017)
**Journal:** Canadian Urological Association Journal
**Section:** abstract | **Similarity:** 0.627

Two-thirds of men with prostates ≥80mL maintained stability with conservative management over 62 months. Clinical progression occurred in 33%. No baseline volume predicted progression.

---

### Chunk 11/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.625

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

### Chunk 12/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.601

e Doxazosina, e estar ciente da Síndrome Pós-Finasterida para preservar os níveis de DHT.
*   [ ] 4. Para médicos: Ao monitorar pacientes em TRT, entender que pequenos aumentos no PSA podem ser fisiológicos, mas investigar saltos abruptos de mais de um ponto percentual.
*   [ ] 5. Ao avaliar um paciente, calcular a relação PSA livre sobre PSA total para diferenciar risco de HPB e câncer de próstata.
*   [ ] 6. Considerar a solicitação de ressonância magnética 3-Tesla para homens com mais de 50 anos, ou com mais de 40 anos se houver história familiar de câncer de próstata ou alterações significativas no PSA.
*   [ ] 7. Em casos de dissociação entre a clínica do paciente e os exames de sangue, considerar a dosagem hormonal salivar para avaliar as frações livres e bioativas.
*   [ ] 8.

---

### Chunk 13/30
**Article:** Lichen sclerosus: The 2023 update (2023)
**Journal:** Frontiers in Medicine
**Section:** other | **Similarity:** 0.601

al dilator for 13 months prior to referral 
to a urologist.
GPP
Oer male patients with urinary symptoms, urethral stricture due to LS or who have failed to topical steroids and/or circumcision referral for a urology opinion and 
further investigation and management of lower urinary tract symptoms and treatment with other surgical options.
GPP
Advise obese male patients with LS and a buried penis to lose weight
Level of recommendation: 
⬆⬆
: Strong recommendation for the use of an intervention. 
⬆
: Weak recommendation for the use of an intervention. GPP: Good practice point recommendations are 
derived from informal consensus.

DeLuca et al. 
10.3389/fmed.2023.1106318
Frontiers in 
Medicine
11
frontiersin.org
therapies failed to control the disease (
171
,
 
172
). A list of systemic 
treatments with the recommended doses are enumerated in 
Table9
 
(
170

181
).
10.4.

---

### Chunk 14/30
**Article:** MFI - Reposição Hormonal - AULA 05 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.593

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

### Chunk 15/30
**Article:** International Urogynecology consultation chapter 2 committee 3: the clinical evaluation of pelvic organ prolapse including investigations into associated morbidity/pelvic floor dysfunction (2023)
**Journal:** International Urogynecology Journal
**Section:** results | **Similarity:** 0.590

function in the studies reviewed; however, there was no agreement on the cut-oﬀ value at which retention was diagnosed ranging from 50 to 200 ml, as shown in Table5.To ﬁnd a cut-oﬀ value for PVR that could predict postop-erative voiding trial results more accurately than a predeter-mined value of 100 ml, one study used a receiver operating curve, but no PVR value was better than 100 ml (the prede-termined value used in the study) [49].The accuracy of translabial ultrasound scan formulae used for PVR measurement in patients with prolapse was examined in one paper [54]. It found that the results obtained by the three published formulae correlated with the catheter-measured PVR.Urine ﬂow studies These included free-ﬂow studies (non-instrumented ﬂow studies) and pressure-ﬂow studies (instru-mented urodynamic ﬂow studies).

---

### Chunk 16/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.588

entes com PSA normal, contra 1-4% na população geral com PSA normal.
**A terapia de reposição de testosterona (TRT) é segura, não aumenta o risco de câncer de próstata e não piora os sintomas da Hiperplasia Prostática Benigna (HPB), com diretrizes modernas permitindo seu início apenas 3 a 6 meses após a prostatectomia, em vez dos antigos 5 a 10 anos.**
- Estudos mostraram que a TRT não "explode" o câncer; em um grupo de 20 pacientes com lesão pré-cancerígena (PIM), apenas 1 (5%) desenvolveu câncer após a terapia.
- Outro estudo com 20 anos de acompanhamento não mostrou aumento no risco de câncer de próstata em usuários de TRT.
- Para pacientes com HPB, um estudo de 2012 demonstrou que a TRT não exacerba os sintomas urinários.
- As diretrizes para iniciar a TRT após a prostatectomia radical foram drasticamente reduzidas: de uma espera de 5 a 10 anos para apenas 3 a 6 meses com PSA indetectável.

---

### Chunk 17/30
**Article:** MFI - Reposição Hormonal - AULA 11 (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.576

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

### Chunk 18/30
**Article:** Efficacy and Safety of Testosterone Replacement Therapy in Men with Hypogonadism: A Meta-Analysis (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.571

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

### Chunk 19/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.570

is para a TRT, protocolos de tratamento, condutas clínicas e o acompanhamento de segurança. Além disso, a sessão cobriu a interpretação do PSA, exames subsidiários como a ressonância magnética 3-Tesla, e a utilidade da dosagem hormonal salivar para uma avaliação clínica completa.
## Conteúdo Abordado
### 1. Introdução e Filosofia da Medicina Preventiva
- **Palestrante e Tema:** O Dr. Wilson Dalla Pasqua Júnior, urologista, andrologista, cirurgião geral e nutrólogo, introduziu o tema da TRT, destacando sua relevância prática para profissionais que trabalham com reposição hormonal e saúde masculina.
- **Abreviações:** Foram definidas as abreviações CAP (Câncer de Próstata) e HPB (Hiperplasia Prostática Benigna).
- **Filosofia Médica:** Citando William James Foley, fundador da Mayo Clinic, a aula foi fundamentada no princípio da medicina preventiva, que busca evitar que a doença se estabeleça.

---

### Chunk 20/30
**Article:** MFI - Reposição Hormonal - AULA 11 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.568

humor e à saúde óssea.
- 5α-redutase: com Nebido, intervenção raramente necessária; com cipionato, maior probabilidade de precisar saw palmetto 300 mg ou, seletivamente, finasterida/dutasterida conforme sintomas/risco-benefício.
- Pigeum africanum: potencial benefício prostático leve; atenção a variações de PSA na TRT.
### 6. Encaminhamento e limites de atuação em casos urológicos
- Limitar atuação em câncer de próstata e quadros complexos; encaminhar a urologistas (preferencialmente com visão integrativa).
- Em PSA elevado e próstata aumentada, avaliar conjuntamente com urologia; usar rede multiprofissional e trabalhar nos limites de competência.
### 7. Modulação de estrogênios com crucíferas (I3C e DIM)
- Indol-3-carbinol (200–400 mg/dia) e diindolilmetano (100–200 mg/dia) podem modular vias de metabolização de estrogênios (perfil de estronas).

---

### Chunk 21/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.565

a um risco aumentado de câncer, enquanto a terapia de reposição hormonal se mostra segura e não agrava a hiperplasia prostática benigna (HPB). A avaliação prostática evoluiu para um modelo mais preciso, utilizando a relação PSA livre/total, o monitoramento do resíduo pós-miccional e a ressonância magnética 3 Tesla para um diagnóstico e acompanhamento mais eficazes.
---
### Evidências Chave
**A crença histórica de que a testosterona causa câncer de próstata, baseada em um estudo de 1941 com apenas 2 pacientes, foi desmentida por estudos modernos que mostram uma taxa de câncer de 14% a 18% em homens com baixa testosterona, muito acima da média populacional de 1% a 4,5%.**
- O dogma que associou testosterona ao câncer de próstata perdurou por cerca de 75 a 80 anos, originado de um estudo seminal de 1941 que, no final, contou com apenas 2 pacientes humanos.

---

### Chunk 22/30
**Article:** MFI - Reposição Hormonal - AULA 05 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.564

levados de testosterona influenciem na etiologia do câncer de próstata.
    - Um estudo mostrou que homens com níveis mais altos de testosterona têm menos chance de câncer de próstata, enquanto níveis mais elevados de estradiol estão associados a uma maior incidência.
    - Uma pesquisa com 3.000 homens com câncer de próstata e mais de 6.100 controles não encontrou relação entre o câncer e os níveis de testosterona total, livre ou outros andrógenos.
    - O urologista Abraham Morgenthaler é citado, afirmando que nunca existiu base científica para acreditar que a testosterona aumenta o risco de câncer de próstata.
*   **TRT e Hiperplasia Prostática Benigna (HPB)**
    - A reposição transdérmica não demonstrou aumento no risco de HPB.
    - Estudos mostram que a TRT não piora os sintomas do trato urinário inferior em homens com HPB.

---

### Chunk 23/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.562

*   **Valores < 1**: Mau prognóstico (prevalência de estrogênios proliferativos), sugerindo a necessidade de trabalhar a metilação da estrona.
## ❓ Perguntas
*   [Inserir Pergunta/Dúvida]
## 📚 Tarefas
*   [ ] 1. Prestar atenção à explicação sobre a origem histórica da associação entre testosterona e câncer de próstata.
*   [ ] 2. Para médicos: Ao avaliar um paciente para TRT, especialmente um sobrevivente de câncer de próstata, seguir um protocolo rigoroso que inclua anamnese completa, exames físicos, radiológicos (Ressonância Magnética, PET-CT) e laboratoriais, aguardando um período seguro de PSA indetectável.
*   [ ] 3. Para médicos: Evitar a prescrição de Dutasterida e Finasterida para o tratamento de HPB, considerando alternativas como Tansulosina e Doxazosina, e estar ciente da Síndrome Pós-Finasterida para preservar os níveis de DHT.
*   [ ] 4.

---

### Chunk 24/30
**Article:** Optimal PSA density threshold for prostate biopsy in benign prostatic obstruction patients with elevated PSA levels but negative MRI findings (2025)
**Journal:** BMC Urology
**Section:** abstract | **Similarity:** 0.562

Study identifying optimal PSAD cutoff of 0.30 ng/ml/cm³ for biopsy decision in BPH patients with elevated PSA but negative MRI, demonstrating 93% specificity and 65% sensitivity. ROC analysis showed PSAD achieved AUC 0.848, outperforming PSA alone (0.722) and free/total PSA ratio (0.635).

---

### Chunk 25/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.560

s a prostatectomia radical foram drasticamente reduzidas: de uma espera de 5 a 10 anos para apenas 3 a 6 meses com PSA indetectável.
**A avaliação da saúde prostática vai além do PSA total, utilizando a relação PSA livre/total como um indicador crucial, onde valores abaixo de 0.14 sugerem maior risco de câncer, enquanto valores acima indicam HPB.**
- O PSA total é composto por uma forma livre (10-30%) e uma complexada (70-90%), e a relação entre elas é diagnóstica.
- A relação PSA livre/total tem um ponto de equilíbrio em 0.14: valores maiores (ex: 0.20, 0.30) são sugestivos de HPB, enquanto valores menores (ex: 0.09, 0.04) aumentam a suspeita de câncer de próstata.
- A concentração normal de PSA total varia de 2.5 a 4.0 ng/ml, mas um aumento de 1.0 ng/ml em um ano ou saltos abruptos (ex: de 2.5 para 4 ou 5 em dois meses) exigem investigação, mesmo dentro da faixa de normalidade.

---

### Chunk 26/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.556

flare" (surto de testosterona), acelerando o crescimento das metástases.
### 4. Conduta Clínica em Pacientes com Histórico de Câncer de Próstata
- A TRT é considerada segura para pacientes que trataram o câncer de próstata e estão curados.
- **Pós-Prostatectomia Radical ou Braquiterapia:** A TRT pode ser iniciada após um período de segurança, geralmente quando o PSA se torna indetectável (a partir de 3-6 meses, segundo a literatura).
- **Protocolo Conservador do Palestrante:** Recomenda-se esperar pelo menos 1 ano de PSA indetectável, realizar exames de imagem avançados (Ressonância Magnética de pelve e PET-CT) para descartar recidiva local ou metástases, e fazer uma avaliação hormonal completa antes de iniciar a TRT.
### 5. Hiperplasia Prostática Benigna (HPB) e TRT
- A HPB não é uma contraindicação para a TRT, exceto em casos de obstrução urinária severa e descompensada.

---

### Chunk 27/30
**Article:** MFI - Reposição Hormonal - AULA 06 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.555

tico de baixas doses de Saw Palmetto (ex.: 150 mg) e monitorar resposta.
- [ ] 3. Ao prescrever gestrinona, iniciar com creme vaginal para testar tolerância antes de considerar implantes.
- [ ] 4. Em homens que necessitam inibição da aromatase, começar com doses baixas de anastrozol (ex.: 0,1 mg) e reavaliar estradiol após 1–2 meses para ajustar a dose.
- [ ] 5. Evitar espironolactona como primeira linha para queixas androgênicas; priorizar inibidores específicos e avaliar toda a cascata hormonal.
- [ ] 6. Estudar para a próxima aula a cadeia hormonal, focando em enzimas, genes e substâncias que influenciam as vias.

---

## SOAP

> Data e Hora: 2025-11-21 04:14:16
> Paciente: 
> Diagnóstico:
## Histórico de Diagnóstico:
1. Histórico Médico: 
2. Histórico de Medicação: Inserir mais aqui
## Subjetivo:
Não há conteúdo de entrevista clínica entre médico e paciente.

---

### Chunk 28/30
**Article:** Efficacy and Safety of Testosterone Replacement Therapy in Men with Hypogonadism: A Meta-Analysis (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.555

:  853-863,  2016
859
Prostate volume. 
A total of 8 RCTs, involving 663 partici
-
pants (355 in the testosterone group and 308 in the control 
group), included the results of prostate volume. The data 
revealed that the average increase in prostate volume was 1.58 
(95% CI, 0.6 to 2.56; P=0.002; Fig. 6B) following TRT. However, 
no significant difference in prostate volume was reported 
following long‑term TRT (MD, ‑0.55; 
95% CI, ‑2.27 to 1.17; 
P=0.96).
Mild to moderate adverse events. 
A total of 6 RCTs, 
involving 1,351 participants (775 in the testosterone group 
and 576 in the control group), included details of 
mild to 
moderate adverse events (Table II). Analysis demonstrated 
that the frequency of mild to moderate adverse events in 
the testosterone group was higher than in the control group 
(MD, 1.58; 95% CI, 1.07 to 
2.33; P=0.02; Fig. 6C), particu
-
larly in patients undergoing long‑term treatment (MD, 3.10; 
95% CI, 1.14 to 8.41; P=0.03).

---

### Chunk 29/30
**Article:** Efficacy and Safety of Testosterone Replacement Therapy in Men with Hypogonadism: A Meta-Analysis (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.550

A, prostate volume
2007 (12)       2 weeks 
Emmelot‑Vonk
 Netherlands 113 110 TT<13.7 nmol/l Y=60‑69, 6 months 160 mg/day TU, PO BMI, body weight,
et al
, 2008 (7)    Y=60‑80 >4.5 µg/l   total c
holesterol, total
     Y>70,    lean body mass, total
     >6.5 µg/l   fat mass, IPSS, BMD,
        adverse event, PSA,
        prostate volume
Aversa 
et al
,  Italy   40   10 TT<3.0 ng/ml, with Age‑adjusted 24 months 1,000 mg TU IM, every  BMI, total cholesterol
2010 (13)    metabolic syndrome, elevated PSA  12 weeks 
    Y=45‑65    
Aversa 
et al
, 
 Italy   32   10 TT<3.20 ng/ml, with  Age‑adjusted 12 months 1,000 mg TU IM, every BMI, total cholesterol,
2010 (14)    metabolic syndrome, elevated PSA  12 weeks AMS score, IPSS, 
    Y=50-65    prostate volume
Idan 
et al
, USA   55   55 H
ealthy, Y>50 >4.0 g/l 24 months 70 mg DHT daily, BMI, total cholesterol,
2010 (15)       transdermal prostate volume, IPSS,
        PSA, mild to moderate
        adverse event, total lean

---

### Chunk 30/30
**Article:** Efficacy and Safety of Testosterone Replacement Therapy in Men with Hypogonadism: A Meta-Analysis (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.546

A, Saglam M, Turan M, 
Corakci A and Ali Gundogan M: Effects of gonadotropin and 
testosterone treatments on prostate volume and serum prostate 
specific antigen levels in male hypogonadism. Endocr J 44: 
719‑724, 1997.
36. 
Calof OM, Singh AB, Lee ML
, 
Kenny AM, Urban RJ, Tenover JL 
and Bhasin S: Adverse events associated with testosterone 
replacement in middle‑aged and older men: A meta‑analysis 
of randomized, placebo‑controlled trials. J Gerontol A Biol Sci 
Med Sci 60: 1451‑1457, 2005.
	
	
Getulio José Mattos Do Amaral Filho - getfilho@yahoo.com.br - CPF: 034.983.039-88

---

