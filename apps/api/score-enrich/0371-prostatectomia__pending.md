# ScoreItem: Prostatectomia

**ID:** `019bf31d-2ef0-73a8-b1ae-3e27305e25d8`
**FullName:** Prostatectomia (Histórico de doenças - Cirurgias já realizadas - Cirurgias que interferem diretamente no escore)

**Preparation Metadata:**
- Quality Grade: **GOOD**
- Total Chunks: 30 de 11 artigos
- Avg Similarity: 0.579

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-73a8-b1ae-3e27305e25d8`.**

```json
{
  "score_item_id": "019bf31d-2ef0-73a8-b1ae-3e27305e25d8",
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

**ScoreItem:** Prostatectomia (Histórico de doenças - Cirurgias já realizadas - Cirurgias que interferem diretamente no escore)

**30 chunks de 11 artigos (avg similarity: 0.579)**

### Chunk 1/30
**Article:** Clinically Localized Prostate Cancer: AUA/ASTRO Guideline, Part II: Principles of Active Surveillance, Principles of Surgery, and Follow-Up (2022)
**Journal:** Journal of Urology
**Section:** abstract | **Similarity:** 0.640

This clinical practice guideline provides evidence-based recommendations for the management of clinically localized prostate cancer. Part II focuses on principles of active surveillance, surgical management including radical prostatectomy, and post-treatment follow-up protocols. The guideline emphasizes PSA measurement as the cornerstone of follow-up after local treatment, with PSA testing recommended every 6-12 months for 5 years and annually thereafter. Biochemical recurrence is defined as PSA levels reaching 0.2-0.4 ng/mL after prostatectomy. The guideline addresses assessment of treatment-related complications including urinary, bowel, and sexual dysfunction.

---

### Chunk 2/30
**Article:** DISFUNÇÃO ERÉTIL (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.636

[ ] Medir circunferência abdominal; se >94, reforçar intervenção; se >102, considerar alto risco e intensificar manejo da síndrome metabólica.
- [ ] Exame físico genital completo: testículos, ginecomastia, placas/curvatura peniana; investigar cicatrizes/cirurgias prévias.
- [ ] Solicitar exames básicos: painel hormonal (incluindo testosterona total/livre), PSA quando indicado, função renal/hepática, inflamatórios, lipidograma; complementar conforme caso.
- [ ] Solicitar ecografia abdominal total (próstata, fígado/esteatose, rins) e, conforme risco, tomografia com escore de cálcio coronariano; considerar teste ergométrico/ecocardiograma.
- [ ] Investigar sono com polissonografia domiciliar em presença de ronco, sonolência, despertares ou redução de ereções matinais.
- [ ] Revisar medicações: 5-ARIs, psicotrópicos, estatinas, anti-hipertensivos, lisdexanfetamina; discutir alternativas e risco/benefício.

---

### Chunk 3/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.630

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

### Chunk 4/30
**Article:** MFI - Reposição Hormonal - AULA 05 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.613

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

### Chunk 5/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.606

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

### Chunk 6/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.602

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

### Chunk 7/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.601

flare" (surto de testosterona), acelerando o crescimento das metástases.
### 4. Conduta Clínica em Pacientes com Histórico de Câncer de Próstata
- A TRT é considerada segura para pacientes que trataram o câncer de próstata e estão curados.
- **Pós-Prostatectomia Radical ou Braquiterapia:** A TRT pode ser iniciada após um período de segurança, geralmente quando o PSA se torna indetectável (a partir de 3-6 meses, segundo a literatura).
- **Protocolo Conservador do Palestrante:** Recomenda-se esperar pelo menos 1 ano de PSA indetectável, realizar exames de imagem avançados (Ressonância Magnética de pelve e PET-CT) para descartar recidiva local ou metástases, e fazer uma avaliação hormonal completa antes de iniciar a TRT.
### 5. Hiperplasia Prostática Benigna (HPB) e TRT
- A HPB não é uma contraindicação para a TRT, exceto em casos de obstrução urinária severa e descompensada.

---

### Chunk 8/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.600

nsiderado alterado (normal > 15 ml/s).
- O tratamento medicamentoso para HPB inclui doses como 0,4 mg de Tansulosina ou 2 a 4 mg de Doxazosina para relaxar a musculatura e melhorar o fluxo urinário.
### Achados Adicionais Chave
- A ressonância magnética multiparamétrica 3 Tesla é uma ferramenta de alta definição, recomendada a partir dos 40 anos para homens com histórico familiar ou a partir dos 50 anos como rotina para uma avaliação prostática precisa.
- A dosagem hormonal salivar oferece uma medição precisa dos hormônios livres, com faixas de referência para testosterona (47-150), estradiol (0.6-3) e o quociente estrogênico (0.04-1.67), que avalia o equilíbrio hormonal.
- A escala Gleason, que vai de 1 a 10, mede a agressividade do câncer de próstata, com tumores classificados como 8, 9 ou 10 sendo considerados os mais agressivos.
- A hiperplasia prostática é uma condição que afeta apenas 3 espécies: humanos, cães e macacos.

---

### Chunk 9/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.595

a um risco aumentado de câncer, enquanto a terapia de reposição hormonal se mostra segura e não agrava a hiperplasia prostática benigna (HPB). A avaliação prostática evoluiu para um modelo mais preciso, utilizando a relação PSA livre/total, o monitoramento do resíduo pós-miccional e a ressonância magnética 3 Tesla para um diagnóstico e acompanhamento mais eficazes.
---
### Evidências Chave
**A crença histórica de que a testosterona causa câncer de próstata, baseada em um estudo de 1941 com apenas 2 pacientes, foi desmentida por estudos modernos que mostram uma taxa de câncer de 14% a 18% em homens com baixa testosterona, muito acima da média populacional de 1% a 4,5%.**
- O dogma que associou testosterona ao câncer de próstata perdurou por cerca de 75 a 80 anos, originado de um estudo seminal de 1941 que, no final, contou com apenas 2 pacientes humanos.

---

### Chunk 10/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.593

s a prostatectomia radical foram drasticamente reduzidas: de uma espera de 5 a 10 anos para apenas 3 a 6 meses com PSA indetectável.
**A avaliação da saúde prostática vai além do PSA total, utilizando a relação PSA livre/total como um indicador crucial, onde valores abaixo de 0.14 sugerem maior risco de câncer, enquanto valores acima indicam HPB.**
- O PSA total é composto por uma forma livre (10-30%) e uma complexada (70-90%), e a relação entre elas é diagnóstica.
- A relação PSA livre/total tem um ponto de equilíbrio em 0.14: valores maiores (ex: 0.20, 0.30) são sugestivos de HPB, enquanto valores menores (ex: 0.09, 0.04) aumentam a suspeita de câncer de próstata.
- A concentração normal de PSA total varia de 2.5 a 4.0 ng/ml, mas um aumento de 1.0 ng/ml em um ano ou saltos abruptos (ex: de 2.5 para 4 ou 5 em dois meses) exigem investigação, mesmo dentro da faixa de normalidade.

---

### Chunk 11/30
**Article:** MFI - Reposição Hormonal - AULA 04 (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.580

tosterona livre” calculada do soro como decisor; incluir, quando indicado, painel salivar (testosterona, DHT, estradiol; progesterona no D22–D24) juntamente com sangue total e SHBG.
- [ ] 5. Implementar triagem de fatores ambientais/ocupacionais que elevem temperatura escrotal (vestimenta apertada, longos períodos sentado, dormir de cueca, ambientes quentes) e orientar medidas corretivas.
- [ ] 6. Estabelecer protocolo para avaliação pós-ciclo de testosterona (endógena/exógena), reconhecendo períodos de LH/FSH inibidos e evitando interpretações equivocadas de queda transitória.
- [ ] 7. Preparar leitura dos estudos recomendados sobre obesidade e hipogonadismo, bariátrica e reversão hormonal, e relações entre obesidade e andropausa, para discussão na próxima aula.
- [ ] 8. Educar equipes clínicas sobre a inadequação de prescrever inibidores de PDE5 (Viagra/Cialis) sem avaliação hormonal quando há suspeita de disfunção androgênica.
- [ ] 9.

---

### Chunk 12/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.580

is baixos de testosterona estão associados a tumores mais agressivos (Gleason mais alto), enquanto níveis elevados parecem ser protetores.
*   **A Exceção: Câncer de Próstata Metastático**: A TRT é contraindicada ou exige extremo cuidado em homens com câncer de próstata metastático ativo, devido ao risco de "testosterone flare" (aumento das metástases).
### 3. Diretrizes de Tratamento e Acompanhamento com TRT
*   **TRT Pós-Tratamento de Câncer de Próstata**:
    *   **Prostatectomia Radical**: A TRT pode ser iniciada após o PSA se tornar indetectável (o palestrante adota uma conduta conservadora de esperar 1 ano).
    *   **Braquiterapia**: A TRT pode ser iniciada após a queda do PSA (em média, 3 meses).
    *   **Conduta do Palestrante**: Antes de iniciar a TRT em sobreviventes de câncer, realiza anamnese completa, exames físicos, aguarda um ano de PSA indetectável, solicita ressonância magnética de pelve, PET-CT prévio e exames laboratoriais completos.

---

### Chunk 13/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.578

entes com PSA normal, contra 1-4% na população geral com PSA normal.
**A terapia de reposição de testosterona (TRT) é segura, não aumenta o risco de câncer de próstata e não piora os sintomas da Hiperplasia Prostática Benigna (HPB), com diretrizes modernas permitindo seu início apenas 3 a 6 meses após a prostatectomia, em vez dos antigos 5 a 10 anos.**
- Estudos mostraram que a TRT não "explode" o câncer; em um grupo de 20 pacientes com lesão pré-cancerígena (PIM), apenas 1 (5%) desenvolveu câncer após a terapia.
- Outro estudo com 20 anos de acompanhamento não mostrou aumento no risco de câncer de próstata em usuários de TRT.
- Para pacientes com HPB, um estudo de 2012 demonstrou que a TRT não exacerba os sintomas urinários.
- As diretrizes para iniciar a TRT após a prostatectomia radical foram drasticamente reduzidas: de uma espera de 5 a 10 anos para apenas 3 a 6 meses com PSA indetectável.

---

### Chunk 14/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.577

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

### Chunk 15/30
**Article:** DISFUNÇÃO ERÉTIL (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.573

ixa tolerância a esforço correlaciona-se com pior desempenho sexual; predomínio simpático (estresse) prejudica ereção.
- Sono e hormônios: apneia obstrutiva do sono reduz testosterona, aumenta endotelina e piora o IIEF; sono é crucial para produção hormonal.
- Exame físico direcionado: testículos (atrofia), ginecomastia (predominância estrogênica), cicatrizes e cirurgias prévias, doença de Peyronie (placas/fibroses), composição corporal (bioimpedância/ISAK; circunferência abdominal >94 e >102 como pontos de risco).
- Exames laboratoriais e imagem: painel hormonal, inflamatório, renal/hepático, lipidograma, PSA quando indicado; ecografia abdominal; risco cardiovascular (teste ergométrico, ecocardiograma, tomografia com escore de cálcio coronariano); polissonografia domiciliar para sono.
### 4.

---

### Chunk 16/30
**Article:** Recovery of Social Continence and Sexual Function in Men With High-risk Prostate Cancer After Radical Prostatectomy: Results From a Statewide Collaborative (2024)
**Journal:** Urology
**Section:** abstract | **Similarity:** 0.573

This study examined postoperative urinary and sexual outcomes in high-risk prostate cancer patients undergoing radical prostatectomy within the Michigan Urological Surgery Improvement Collaborative (2014-2023). Among 1,323 patients, 58% achieved social continence (≤1 pad daily) at 3 months and 86% at 12 months. Only 15% recovered sexual function at 12 months. Continence recovery was associated with higher baseline EPIC-26 urinary continence scores and negatively correlated with advancing age. Sexual function recovery was linked to nerve-sparing techniques, lower preoperative PSA levels, and absence of adjuvant therapy.

---

### Chunk 17/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.573

e Doxazosina, e estar ciente da Síndrome Pós-Finasterida para preservar os níveis de DHT.
*   [ ] 4. Para médicos: Ao monitorar pacientes em TRT, entender que pequenos aumentos no PSA podem ser fisiológicos, mas investigar saltos abruptos de mais de um ponto percentual.
*   [ ] 5. Ao avaliar um paciente, calcular a relação PSA livre sobre PSA total para diferenciar risco de HPB e câncer de próstata.
*   [ ] 6. Considerar a solicitação de ressonância magnética 3-Tesla para homens com mais de 50 anos, ou com mais de 40 anos se houver história familiar de câncer de próstata ou alterações significativas no PSA.
*   [ ] 7. Em casos de dissociação entre a clínica do paciente e os exames de sangue, considerar a dosagem hormonal salivar para avaliar as frações livres e bioativas.
*   [ ] 8.

---

### Chunk 18/30
**Article:** Efficacy and Safety of Testosterone Replacement Therapy in Men with Hypogonadism: A Meta-Analysis (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.572

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
**Section:** other | **Similarity:** 0.572

urinário (Qmax). Valores < 10 ml/segundo indicam obstrução.
    *   **Ressonância Magnética Multiparamétrica 3-Tesla (3T)**: Exame de alta especificidade. O palestrante solicita de rotina para homens > 50 anos, ou > 40 anos com histórico familiar ou alterações súbitas no PSA.
*   **Dosagem Hormonal Salivar e Quociente Estrogênico**:
    *   **Vantagens da Saliva**: Via não invasiva que mede a fração livre e 100% bioativa dos hormônios (Testosterona, DHT, Estradiol, etc.). Útil quando a clínica do paciente não corresponde aos exames de sangue.
    *   **Quociente Estrogênico**: Fórmula para avaliar o risco de doenças prostáticas.
        *   **Fórmula**: Estriol / (Estradiol + Estrona).
        *   **Valores > 1**: Bom prognóstico (perfil estrogênico protetor).
        *   **Valores < 1**: Mau prognóstico (prevalência de estrogênios proliferativos), sugerindo a necessidade de trabalhar a metilação da estrona.

---

### Chunk 20/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XVII (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.570

nal.
- Necessidade de farmacovigilância prolongada, consentimento informado e comunicação de riscos ao paciente.
> Sugestões de IA
> - Estruture riscos em categorias (neuropsiquiátricos, metabólicos, hepatorrenais, sexuais).
> - Slide de referências-chave (título, ano, achados).
> - Diferencie associação vs. causalidade com exemplos.
> - Protocolo de monitorização trimestral (PHQ-9/GAD-7, função sexual, testes hepáticos, perfil hormonal) e estratégias de mitigação (nutrição, exercício, sono, silimarina/alcachofra quando aplicável).
> - Tabela comparativa finasterida vs. dutasterida (enzimas, potência, perfil de risco).
> - Consentimento informado padronizado com sinais de alarme.
### 5. Mecanismo Hormonal: Bloqueio da 5-Alfa-Redutase e Consequências
- Finasterida/dutasterida reduzem conversão de testosterona em DHT; aumentar testosterona não resolve se a via de conversão está bloqueada.

---

### Chunk 21/30
**Article:** Global trends of cancer: The role of diet, lifestyle, and environmental factors (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.566

0.3390/nu807041936.HyunJS.Prostatecancerandsexualfunction.WorldJMensHealth.2012;30(2):99–107.https://doi.org/10.5534/wjmh.2012.30.2.9937.JianZ,YeD,ChenY,LiH,WangK.Sexualactivityandriskofprostatecancer:adoseresponsemetaanalysis.JSexMed.2018;15(9):1300–9.https://doi.org/10.1016/j.jsxm.2018.07.00438.RapaportL.Doessexprotectmenagainstprostatecancer?ReutersHealthNews.Availablefrom:https://www.reuters.com/article/us-health-ejaculation-prostate-cancer-idUSKCN0
XJ1YC(2016).Accessed22April2016.39.MarignolL,CoffeyM,LawlerM,HollywoodD.Hypoxiainprostatecancer:apowerfulshieldagainsttumourdestruction?CancerTreatRev.2008;34(4):313–27.https://doi.org/10.1016/j.ctrv.2008.01.00640.SinghO,BollaSR.Anatomyabdomenandpelvisprostate[Internet].TreasureIsland(FL):StatPearlsPublishing;2021[updated2021Jul26;cited2021].Availablefrom:https://www.ncbi.nlm.nih.gov/books/NBK540987/.Accessed26Jul2021.41.AhmadF,CherukuriMK,ChoykePL.Metabolicreprogram-minginprostatecancer.BrJCancer.2021;125(9):1185–96.https://d

---

### Chunk 22/30
**Article:** MFI - Reposição Hormonal - AULA 11 (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.560

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

### Chunk 23/30
**Article:** DISFUNÇÃO ERÉTIL (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.559

elatar redução na libido, frequência das relações e na rigidez da ereção, especialmente em casos de hipogonadismo. A ausência de ereções matinais (tumescência peniana noturna) também é um sintoma importante, frequentemente associado à apneia do sono.
## Objetivo:
*   **Exame Físico:**
    *   Avaliação da composição corporal (bioimpedância, antropometria ou medição da circunferência abdominal).
    *   Exame genital para avaliar atrofia testicular, palpação do pênis para identificar calcificações ou fibroses (sugestivo de Doença de Peyronie).
    *   Verificação de ginecomastia.
    *   Busca por cicatrizes de cirurgias prévias na região perineal, inguinal e baixo ventre.
*   **Questionários:** Uso do questionário validado "Índice Internacional de Função Erétil" para estratificar o grau da disfunção (leve, moderada ou severa).

---

### Chunk 24/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.557

*   **Valores < 1**: Mau prognóstico (prevalência de estrogênios proliferativos), sugerindo a necessidade de trabalhar a metilação da estrona.
## ❓ Perguntas
*   [Inserir Pergunta/Dúvida]
## 📚 Tarefas
*   [ ] 1. Prestar atenção à explicação sobre a origem histórica da associação entre testosterona e câncer de próstata.
*   [ ] 2. Para médicos: Ao avaliar um paciente para TRT, especialmente um sobrevivente de câncer de próstata, seguir um protocolo rigoroso que inclua anamnese completa, exames físicos, radiológicos (Ressonância Magnética, PET-CT) e laboratoriais, aguardando um período seguro de PSA indetectável.
*   [ ] 3. Para médicos: Evitar a prescrição de Dutasterida e Finasterida para o tratamento de HPB, considerando alternativas como Tansulosina e Doxazosina, e estar ciente da Síndrome Pós-Finasterida para preservar os níveis de DHT.
*   [ ] 4.

---

### Chunk 25/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.550

- A HPB não é uma contraindicação para a TRT, exceto em casos de obstrução urinária severa e descompensada.
- A TRT pode causar um aumento mínimo e esperado do volume prostático e do PSA (ex: 0.2-0.4), o que é considerado normal. Um salto abrupto do PSA (>1.0 ponto em um ano) exige investigação.
- O volume da próstata não se correlaciona diretamente com os sintomas; o mais importante é a direção do crescimento (para dentro ou para fora do canal urinário).
- **Tratamento da HPB:** Medicamentos como Tansulosina e Doxazosina são indicados. Inibidores da 5-alfa-redutase (Finasterida e Dutasterida) foram proscritos pelo palestrante devido ao risco de efeitos colaterais severos (síndrome pós-finasterida).
### 6. Avaliação Diagnóstica da Próstata
#### 6.1. Antígeno Prostático Específico (PSA)
- **Função e Formas:** O PSA é uma enzima que liquefaz o sêmen. Existe na forma livre e complexada (ligada a proteínas).

---

### Chunk 26/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.548

s modernas, lideradas por pesquisadores como Abraham Morgenthaler, que demonstram uma relação inversa: baixos níveis de testosterona estão associados a um maior risco de câncer de próstata agressivo. A palestra detalha as diretrizes atuais para a TRT em sobreviventes de câncer e em pacientes com HPB, concluindo que o tratamento é seguro e benéfico na maioria dos casos, com exceção de cânceres metastáticos ativos. Além disso, são discutidas ferramentas diagnósticas cruciais para a saúde prostática, como a interpretação do Antígeno Prostático Específico (PSA) e sua relação livre/total, exames de imagem como a ressonância magnética 3-Tesla, e a dosagem hormonal salivar para avaliar as frações bioativas de hormônios e o risco prostático através do quociente estrogênico.
## 🔖 Pontos de Conhecimento
### 1.

---

### Chunk 27/30
**Article:** DISFUNÇÃO ERÉTIL (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.546

os Passos/Exames:**
    *   Realizar anamnese ampla e exame físico completo.
    *   Aplicar o questionário "Índice Internacional de Função Erétil".
    *   Solicitar exames laboratoriais (perfil hormonal, Vitamina D, ácido fólico, marcadores inflamatórios, etc.).
    *   Solicitar ecografia abdominal total.
    *   Considerar tomografia com score de cálcio coronariano e polissonografia.
    *   Em caso de falha no tratamento de primeira linha, referenciar a um especialista para tratamentos de segunda linha (medicamentos injetáveis).
*   **Plano de Tratamento de Acompanhamento:**
    *   **Mudanças no Estilo de Vida:**
        *   **Dieta:** Adotar uma dieta baseada em proteínas e gorduras boas, com vegetais de alta qualidade, evitando alimentos ultraprocessados e carboidratos refinados.
        *   **Atividade Física:** Incentivar exercícios aeróbicos de intensidade leve a vigorosa, pelo menos 40 minutos, 4 vezes por semana (total de 160 min/semana).

---

### Chunk 28/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XVII (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.543

ções no tamanho e curvatura do pênis. Os efeitos podem persistir indefinidamente após a descontinuação.
*   **Psicológicos:** Problemas de memória, ideação suicida, insônia, ataques de pânico, ansiedade, depressão, "brain fog" (enevoamento do cérebro), anedonia, sentimentos de desesperança. Depressão clinicamente significativa foi relatada em 50% dos pacientes com a síndrome.
*   **Físicos:** Fadiga, perda de massa muscular e mal-estar geral.
## Objetivo:
A transcrição combina os achados de um paciente específico com informações de estudos e discussões médicas gerais.
**Achados do Paciente Específico ([Speaker 1]):**
*   **Exame de Metabolômica Hormonal:**
    *   Testosterona: Nível zero (ou próximo de zero).
    *   Diidrotestosterona (DHT): Nível zero, devido ao bloqueio da enzima 5-alfa-redutase.
    *   16-hidroxiestrona e 4-hidroxiestrona: Elevadas, indicando desvio do metabolismo hormonal.

---

### Chunk 29/30
**Article:** Terapia de Reposição Hormonal com Testosterona XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.542

saltos abruptos (ex: de 2.5 para 4 ou 5 em dois meses) exigem investigação, mesmo dentro da faixa de normalidade.
**O diagnóstico e manejo da Hiperplasia Prostática Benigna (HPB) dependem mais dos sintomas obstrutivos, como resíduo pós-miccional acima de 40 ml, do que do tamanho da próstata, que pode variar de 25 a mais de 80 gramas sem necessariamente causar problemas.**
- A HPB é comum a partir dos 45-50 anos, mas o tamanho da próstata (normalmente 25-30 gramas) não se correlaciona diretamente com a obstrução; próstatas de 28-29 gramas podem ser obstrutivas, enquanto outras de 70-80 gramas não.
- Um indicador chave de obstrução é o resíduo pós-miccional, com volumes acima de 40 ml sendo anormais, e a urofluxometria, onde um fluxo máximo (Qmax) abaixo de 10 ml/s é considerado alterado (normal > 15 ml/s).
- O tratamento medicamentoso para HPB inclui doses como 0,4 mg de Tansulosina ou 2 a 4 mg de Doxazosina para relaxar a musculatura e melhorar o fluxo urinário.

---

### Chunk 30/30
**Article:** Patient reported outcomes and health related quality of life in localized prostate cancer: A review of current evidence (2022)
**Journal:** Urology Oncology
**Section:** abstract | **Similarity:** 0.541

This review examines how different treatment approaches for early-stage prostate cancer affect patient quality of life. Analysis of 4 randomized trials and 15 prospective studies (2010-2021) shows that surgery has the largest short and long-term negative effect on sexual function and incontinence but advantages with regards to bowel function. Radiation therapy primarily impacts urinary symptoms and bowel issues, while active surveillance shows favorable short-term outcomes. Long-term global quality of life impact regarding anxiety, mental, emotional well-being, and fatigue seem to be equivalent between treatment modalities.

---

