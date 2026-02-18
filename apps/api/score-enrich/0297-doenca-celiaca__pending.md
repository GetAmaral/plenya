# ScoreItem: Doença celíaca

**ID:** `019bf31d-2ef0-702c-9375-02d892ce7440`
**FullName:** Doença celíaca (Histórico de doenças - Doenças crônicas (Questionar ativamente as doenças crônicas mais comuns, perguntar sobre duração, grau de controle e tratamentos utilizados no passado e atualmente) - Doenças auto-imunes)

**Preparation Metadata:**
- Quality Grade: **EXCELLENT**
- Total Chunks: 30 de 14 artigos
- Avg Similarity: 0.623

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-702c-9375-02d892ce7440`.**

```json
{
  "score_item_id": "019bf31d-2ef0-702c-9375-02d892ce7440",
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

**ScoreItem:** Doença celíaca (Histórico de doenças - Doenças crônicas (Questionar ativamente as doenças crônicas mais comuns, perguntar sobre duração, grau de controle e tratamentos utilizados no passado e atualmente) - Doenças auto-imunes)

**30 chunks de 14 artigos (avg similarity: 0.623)**

### Chunk 1/30
**Article:** TDAH - Parte XIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.699

e permeabilidade, início de autoimunidade, inflamação sistêmica com potencial de neuroinflamação.
## Indicadores laboratoriais e achados clínicos
- Testes fecais/copro (ex.: Copromax, GI-MAP, Gut Check): calprotectina, zonulina, IgA secretória, elastase mostram integridade/barreiras.
- Zonulina sérica elevada: associada a permeabilidade intestinal e comprometimento de funcionamento social em TDAH, TEA, TOC (meta-análise; 402 participantes).
## Sensibilidade ao glúten não celíaca: perfis e sintomas
- Trato baixo:
  - Diarreia 16,5%; constipação 18,2%; alteração de hábito 27%;
  - Dor/desconforto abdominal 67–83%; distensão 72–87%; perda de peso 25%.
- Trato alto:
  - Dor epigástrica 52%; náusea até 44%; aerofagia 36%; refluxo 32%; estomatite 31%.
- Extraintestinais: dermatite, depressão, brain fog, ansiedade, confusão, cefaleia; fadiga 23–74% (crianças fadigadas tendem a agitar, inclusive à noite).

---

### Chunk 2/30
**Article:** Intolerâncias, Alergias e Hipersensibilidades Alimentares II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.668

peratividade, déficit de atenção.
### 14. Exames laboratoriais básicos e imunológicos
- Hemograma: pode ser normal; eosinofilia sugere esofagite eosinofílica/enterocolopatias; plaquetas >400 mil sugerem enteropatia inflamatória crônica.
- Imunoglobulinas: IgA aumentada na doença celíaca; IgE aumentada em alergias tipo I.
- IgG/IgG4: IgG4 pode modular IgE; pode aumentar na esofagite eosinofílica; uso cauteloso, não diagnóstico isolado.
- Eletroforese de proteínas: alterações em gamaglobulinas indicam cronicidade.
- Enteropatia perdedora de proteínas: pode cursar com hipogamaglobulinemia.
- Anticorpos contra glúten: recomendados na investigação.
### 15. Fenotipagem linfocitária e interpretação (CD4/CD8 e marcadores)
- Relação CD4/CD8 esperada: 1,5–2,5.
- CD8 elevado: favorece alergia alimentar celular (perfil TH1).
- CD8 muito baixo: deficiência de tolerância imunológica.
- CD4 aumentado: alergias tipo I (humoral).

---

### Chunk 3/30
**Article:** TDAH - Parte XIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.665

e pode iniciar autoimunidade. Inflamação intestinal crônica pode afetar o cérebro via neuroinflamação recorrente.
* Evidências em crianças
   - Revisão sistemática/meta-análise: níveis séricos elevados de zonulina associados à hiperpermeabilidade e afetam vias neurais/hormonais/imunológicas; 4 artigos, 402 participantes, em TDAH, TEA e TOC. No TDAH, zonulina elevada associada a pior funcionamento social versus controles.
### 7. Sensibilidade ao glúten não celíaca: sintomas e abordagem
* Sintomas gastrointestinais
   - Diarreia 16,54%; constipação 18,24%; alteração de hábito intestinal 27%; dor/desconforto abdominal 67–83%; distensão abdominal 72–87%; perda de peso 25%.
* Trato digestivo alto
   - Dor epigástrica 52%; náusea até 44%; aerofagia 36%; refluxo 32%; estomatite 31%.
* Extraintestinais
   - Dermatites, depressão, “fog mind/brain fog”, ansiedade, confusão, dores de cabeça; fadiga 23–74%.

---

### Chunk 4/30
**Article:** Intolerâncias, Alergias e Hipersensibilidades Alimentares II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.663

(ex.: intoxicação escombroide em peixes como atum/cavala).
- Não imunológicas:
  - Enzimáticas: intolerância à histamina, intolerância à lactose.
  - Farmacológicas: cafeína, tiramina.
  - Má absorção de frutose: transporte por GLUT5/GLUT2 (não GLUT4).
- Imunológicas:
  - Doença celíaca (autoimune).
  - Tipo I (IgE): urticária, angioedema, broncoespasmo, asma, anafilaxia, síndrome alérgica oral.
  - Não IgE mediadas: FPIES, proctocolite.
  - Mistas: esofagite, gastrite, enterocolite eosinofílica.
  - Tipo III tardia também mencionada.
### 12. Abordagem diagnóstica inicial e achados clínicos
- Anamnese é fundamental; considerar infecções gastrointestinais prévias, resposta TH2 nos primeiros 6 meses.
- História familiar: um dos pais com alergia → risco ~30%; ambos → ~80%.
- Tipo de parto, aleitamento materno exclusivo e uso precoce de mamadeira.
- Exame físico: dor à palpação da fossa ilíaca direita pode sugerir inflamação em placas de Peyer.

---

### Chunk 5/30
**Article:** Glúten (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.657

rância variável.
    3.  **Alergia ao Trigo:** Reação alérgica (mediada por IgE ou não), com tolerância variável.
*   **Biomarcadores e Testes Diagnósticos:**
    *   **Doença Celíaca:** Inicia-se com anticorpos (IgA anti-TTG e IgA total). A positividade geralmente requer confirmação por endoscopia. A genética (HLA-DQ2/DQ8) é útil para exclusão.
    *   **Alergia ao Trigo:** Testes de IgE (cutâneos ou sanguíneos) podem ser usados, mas muitas alergias não são mediadas por IgE, dificultando o diagnóstico.
    *   **SGNC:** O diagnóstico é complexo e de exclusão. Anticorpos anti-gliadina podem ajudar, mas não são definitivos. O futuro aponta para um painel de biomarcadores.
*   **Soberania da Clínica:** Os exames podem ser insuficientes. A dieta de eliminação continua sendo uma ferramenta valiosa, pois a clínica é soberana.
### 5.

---

### Chunk 6/30
**Article:** Glúten (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.646

identificar o problema e classificar o paciente corretamente.
*   **Os biomarcadores variam conforme a desordem.** Para a doença celíaca, os testes de anticorpos (IgA anti-TTG) e genéticos (HLA-DQ2/DQ8) são cruciais. Para alergia, usam-se testes de IgE ou IgG4. Para a sensibilidade não celíaca, os biomarcadores são menos específicos.
*   **A genética tem um papel chave na doença celíaca.** A presença dos genes HLA-DQ2 ou DQ8 é quase 100% necessária para desenvolver a doença, tornando o teste genético útil para excluir o diagnóstico.
*   **Atualmente, não há enzimas eficazes para a digestão do glúten.** Produtos enzimáticos no mercado não têm eficácia comprovada, embora haja endopeptidases bacterianas em ensaios clínicos.
### 3.

---

### Chunk 7/30
**Article:** Intolerâncias, Alergias e Hipersensibilidades Alimentares II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.639

odem indicar tolerância e não alergia.
    - **Endoscopia/Colonoscopia:** Podem revelar achados como hiperplasia nodular linfoide.
### 5. Manejo e Tratamento
*   **Dietas de Eliminação:** Principal abordagem, consiste em retirar o alimento agressor. Deve ser feita com acompanhamento multidisciplinar para evitar déficits nutricionais, especialmente em crianças.
*   **Melhora da Digestão:** Uma digestão inadequada aumenta a carga de antígenos no intestino. O uso de enzimas digestivas pode ajudar a degradar melhor as proteínas e diminuir os sintomas. Fatores como pasteurização e Reação de Maillard podem aumentar a alergenicidade dos alimentos.
*   **Modulação Intestinal:** É o pilar do tratamento.
    - **Microbiota e AGCC:** Uma dieta rica em fibras aumenta a produção de ácidos graxos de cadeia curta (butirato, propionato), que são anti-inflamatórios e fortalecem a barreira intestinal.
    - **Probióticos:** Cepas específicas como *Lactobacillus rhamnosus*, *L.

---

### Chunk 8/30
**Article:** Intolerâncias, Alergias e Hipersensibilidades Alimentares II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.639

arcadores alérgicos.
- Dietas de eliminação graduais: 2 alimentos (laticínios e glúten), 4 alimentos (glúten, laticínios, soja e frutos do mar) e 6 alimentos; maior restrição pode alterar a resposta clínica, orientando estratégias individualizadas.
**Achados de coocorrência e sensibilização cruzada ampliam o escopo clínico da avaliação.**
- Síndrome de alergia alimentar relacionada ao látex ocorre em até 50% dos pacientes com alergia ao látex, indicando alta coocorrência e sensibilização cruzada.
**Outras Constatações Importantes**
- Plaquetas acima de 400.000 podem estar relacionadas à enteropatia inflamatória crônica, servindo como achado laboratorial sugestivo.
- A frutose é descrita como absorvida via GLUT4, explicando possíveis quadros de má absorção e reações não imunológicas que imitam alergia.

---

### Chunk 9/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XVI (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.635

e microbioma intestinal.
    - **Avaliação de Alergias Cutâneas:** Realizar testes de exclusão alimentar (ex: remover laticínios/glúten por um mês) e considerar testes de intolerância alimentar por IgG. Investigar a causa de problemas com laticínios (lactose, proteína ou histamina).
    - **Avaliação Geral:** Avaliar marcadores inflamatórios, eixo HPA (estresse) e realizar uma avaliação hormonal completa.
- **Plano de Tratamento de Acompanhamento:**
    - **Intervenção Dietética:** Implementar uma dieta de eliminação personalizada com base nos resultados dos testes, removendo alimentos reativos (por exemplo, classe 4 no teste de IgG) por 2-3 meses para controlar a inflamação e a resistência à insulina.
    - **Saúde Intestinal:** Melhorar o microbioma e a integridade da barreira intestinal através de dieta, fibras e probióticos.

---

### Chunk 10/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.629

intestinal.
-   **SIFO:** Sensação de "blurring" (efeito baiacu - sentir-se distendido mesmo sem distensão objetiva) e dor abdominal.
## Objetivo:
A palestra descreve os exames e achados objetivos para a avaliação de pacientes com suspeita de Síndrome do Intestino Irritável (SII), embora não contenha achados de exame físico de um paciente específico.
-   **Exames Laboratoriais Gerais:** Hemograma e marcadores de atividade inflamatória.
-   **Calprotectina Fecal:** Usada para descartar doença inflamatória intestinal. Um valor abaixo de 100 µg/g tem uma positividade de 98% para o diagnóstico de SII. Entre 100-250 µg/g é uma zona cinzenta. Acima de 250 µg/g requer colonoscopia.
-   **Avaliação para Doença Celíaca:** Dosagem de IgA sérica total e anticorpo antitransglutaminase IgA. Se a IgA for baixa, dosar IgG sérica e antitransglutaminase IgG.
-   **Teste Parasitológico de Fezes:** Recomendado no contexto do Brasil.

---

### Chunk 11/30
**Article:** Glúten (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.624

nifestar como mudanças de humor, TDAH, perda de memória, fadiga crônica e, em casos mais graves, esquizofrenia ou epilepsia.
*   **Mecanismos do Glúten no Cérebro:**
    *   **Teoria da Gliadorfina:** Fragmentos de glúten (gliadorfinas) podem cruzar a barreira intestinal e se ligar a receptores de endorfina no cérebro, alterando o comportamento.
    *   **Teoria da Zonulina:** Outros fragmentos induzem a liberação de zonulina, aumentando a permeabilidade e iniciando uma cascata inflamatória que pode atingir o cérebro.
### 4. Diagnóstico e Espectro das Desordens Relacionadas ao Glúten
*   **Espectro das Reações ao Glúten:**
    1.  **Doença Celíaca:** Autoimune, exige dieta 100% sem glúten.
    2.  **Sensibilidade ao Glúten Não Celíaca (SGNC):** Inflamatória, com tolerância variável.
    3.  **Alergia ao Trigo:** Reação alérgica (mediada por IgE ou não), com tolerância variável.

---

### Chunk 12/30
**Article:** Remissão do Lúpus Através da Medicina Funcional Integrativa (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.620

erência pelo filtro físico, em detrimento do filtro químico
- [ ] Fazer reaplicações frequentes do protetor solar, quando suar muito ou quando tiver contato com a água
- [ ] Retirar o glúten da dieta, para melhorar os sintomas das doenças autoimunes não celíacas
- [ ] Recuperar a integridade intestinal, para pensar em um processo de remissão
- [ ] Fazer uma hidratação adequada, para recuperar a integridade intestinal
- [ ] Evitar bebida alcoólica, para recuperar a integridade intestinal
- [ ] Excluir glúten e lácteos da dieta, para recuperar a integridade intestinal
- [ ] Reduzir os açúcares, para recuperar a integridade intestinal
- [ ] Optar por carboidratos de baixa carga glicêmica, para recuperar a integridade intestinal
- [ ] Garantir uma ingestão adequada de fibras, frutas e verduras, para recuperar a integridade intestinal
- [ ] Encaminhar o paciente para uma nutricionista funcional, para que haja uma correção na quantidade de carboidratos, proteínas e par

---

### Chunk 13/30
**Article:** Hipertensão Arterial Sistêmica II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.619

Lp(a), APO-B/APO-A, NO) para prevenção e tratamento da DCV.
## ❓ Perguntas
- [Inserir Pergunta/Confusão]
## 📚 Tarefas
- [ ] 1. Solicitar avaliação do índice de ômega 3 e da razão ômega 3:ômega 6; ajustar suplementação de ômega 3 conforme resultados.
- [ ] 2. Dosar vitamina D (25(OH)D) e PTH; estabelecer metas de 80 ng/mL para cardiopatas/hipertensos e considerar >100 ng/mL para autoimunes, com monitorização de segurança.
- [ ] 3. Aplicar protocolo de curva de glicose e insulina (jejum, 30, 60, 90, 120 min) para detectar hiperinsulinemia oculta e resistência à insulina.
- [ ] 4. Introduzir metformina em casos de resistência à insulina, juntamente com plano nutricional coordenado com nutricionista.
- [ ] 5. Avaliar homocisteína e intervir com vitaminas B9, B12 e B6 em casos de elevação, considerando polimorfismos de metilação.
- [ ] 6.

---

### Chunk 14/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte VIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.617

r intervenções alimentares e probióticas/fibras com base em metabolômica.
### 3. Doença Celíaca e Hiperpermeabilidade Intestinal
* Subdiagnóstico de doença celíaca
   - Especialistas (ex.: professor de gastroenterologia infantil em Harvard) defendem subdiagnóstico e necessidade de novos marcadores ou uso mais amplo dos já validados.
   - Compreender leaky gut e sua relação com celíaca é vital para evitar danos crônicos e diagnósticos equivocados como IBS.
* Impacto da hiperpermeabilidade (leaky gut)
   - Perda de muco e ruptura de junções entre enterócitos permitem passagem de fragmentos alimentares, restos bacterianos e LPS, gerando endotoxemia e respostas imunes exacerbadas.
   - Associada a doenças neurológicas (depressão, ansiedade, TDAH), dermatológicas (acne, rosácea, psoríase, eczema), tireoidopatias, colopatias, artrite reumatoide, fibromialgia, cefaleias, fadiga, alergias; ampliada por excesso de fármacos e dieta inadequada.
### 4.

---

### Chunk 15/30
**Article:** TDAH - Parte XIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.616

; estomatite 31%.
* Extraintestinais
   - Dermatites, depressão, “fog mind/brain fog”, ansiedade, confusão, dores de cabeça; fadiga 23–74%. Em crianças, sintomas podem levar a diagnósticos comportamentais (p.ex., TDAH) sem investigar dieta (ex.: ultraprocessados como pão de forma “Seven Boys”).
* Crítica à abordagem convencional
   - Medicina convencional frequentemente só reconhece doença celíaca, negligenciando sensibilidade não celíaca; tende a medicalizar sintomas (simeticona, buscopan) e manter glúten. O instrutor sustenta testar exclusão de glúten por 3 semanas com avaliação clínica.
* Teste e confirmação
   - Exame de alergias tardias ao glúten pode ser considerado, mas um teste prático é exclusão de 3 semanas; para descartar doença celíaca posteriormente, reintroduzir glúten e realizar exames diagnósticos apropriados.
### 8.

---

### Chunk 16/30
**Article:** MFI - SÍNDROME PÓS COVID - AULA 02 (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.615

*   **Respiratórios:** Rinorreia, congestão nasal, dispneia.
    *   **Neurológicos:** Dores de cabeça, *brain fog*.
    *   **Cardíacos:** Taquicardia, palpitações.
    *   **Gastrointestinais:** Dores abdominais, diarreia, constipação, náuseas.
    *   **Cutâneos:** Urticária, rubor, eczema.

**Diagnóstico e Tratamento:**
*   A suspeita deve ser levantada em pacientes com histórico de alergias ou quadros clínicos muito vastos.
*   **Diagnóstico:**
    1.  **Dosagem de metil-histamina** em urina de 24 horas.
    2.  **Análise da atividade da enzima DAO** (disponível no exame Copromax, que também avalia o *leaky gut*).
*   **Tratamento:**
    1.  **Dieta anti-histamínica:** Restringir por um mês alimentos ricos em histamina (queijos, fermentados), liberadores de histamina ou inibidores da DAO.
    2.  **Medicação:** O uso do anti-histamínico E-Bastel (10 mg, duas vezes ao dia por um mês, seguido de uma vez ao dia por mais um mês) pode ser uma estratégia.

---

### Chunk 17/30
**Article:** Microbioma Intestinal III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.614

stinais.
    *   **Lactoferrina Fecal:** Glicoproteína liberada por neutrófilos durante a inflamação, confirmando um quadro inflamatório.
    *   **IgA Secretória (SGA) Fecal:** Marcador da função imunológica da mucosa. Níveis baixos indicam baixa defesa e maior suscetibilidade a infecções e disbiose.
    *   **Zonulina Fecal:** Principal marcador de permeabilidade intestinal. Seu aumento, frequentemente associado ao glúten, é um precursor de inflamação sistêmica e doenças autoimunes.
*   **Função Pancreática**
    *   **Elastase Pancreática Fecal:** Marcador da função pancreática exócrina. Um valor baixo pode indicar insuficiência pancreática, muitas vezes secundária à falta de acidificação estomacal.
### 5. Abordagem Terapêutica
*   **Escala de Prioridades na Consulta**
    *   A avaliação deve seguir a ordem: 1. História Pregressa, 2. História Clínica, 3. Medicamentos, 4. Hábitos Alimentares, 5. Exercícios Físicos.

---

### Chunk 18/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa III (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.613

baixo de 100 µg/g são altamente sugestivos de SII.
*   **Avaliação para Doença Celíaca**
    - É fundamental em todos os pacientes, não apenas naqueles com diarreia. Inclui a dosagem de IgA sérica total e o anticorpo antitransglutaminase IgA.
*   **Avaliação da Microbiota e Metabolômica**
    - A metabolômica (avaliação dos produtos da microbiota) é considerada mais importante que a análise da microbiota isolada. A análise de ácidos orgânicos urinários é uma ferramenta útil.
    - A dosagem de zonulina pode ser um bom marcador para o aumento da permeabilidade intestinal.
*   **Supercrescimento Bacteriano do Intestino Delgado (SIBO)**
    - Incidência 3,7 vezes maior em portadores de SII. O diagnóstico prático é feito pelo teste respiratório com lactulose ou glicose.
*   **Supercrescimento Metanogênico Intestinal (IMO)**
    - Supercrescimento de arqueias produtoras de metano, associado principalmente à constipação intestinal.

---

### Chunk 19/30
**Article:** Glúten (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.611

tinua sendo uma ferramenta valiosa, pois a clínica é soberana.
### 5. Estilo de Vida, Dieta e Prevenção
*   **Importância do Estilo de Vida:** Para minimizar o risco de doenças inflamatórias, é crucial cuidar da saúde intestinal com nutrição equilibrada, manejo do estresse, boa higiene do sono e evitando álcool e fumo.
*   **Dieta e Outros Componentes:**
    *   **Lácteos e Lectinas:** O consumo excessivo, assim como o do glúten, pode causar problemas. A moderação é a chave.
    *   **Dieta Ocidental:** O consumo excessivo e de baixa qualidade de alimentos processados está associado ao aumento de doenças inflamatórias crônicas.
*   **Doenças Autoimunes e Inflamação:**
    *   A tireoidite de Hashimoto tem comorbidade crescente com a doença celíaca. Muitos pacientes melhoram com a restrição de glúten.
    *   Qualquer processo inflamatório pode afetar a fertilidade, que tem se tornado um problema crescente em gerações mais jovens.
### 6.

---

### Chunk 20/30
**Article:** TDAH - Parte XIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.610

: dermatite, depressão, brain fog, ansiedade, confusão, cefaleia; fadiga 23–74% (crianças fadigadas tendem a agitar, inclusive à noite).
- Implicação prática: sintomas “funcionais” frequentemente subestimados; retirar glúten por 3 semanas e avaliar resposta é estratégia diagnóstica simples; depois reintrodução para descartar doença celíaca via exames.
## Condutas práticas em alimentação e suplementação
- Priorizar modulação intestinal antes de ampliar fontes específicas (fibras adequadas ao caso, reduzir fermentação).
- Evitar excessos de lácteos, embutidos e alimentos ricos em histamina em quadros de excitabilidade/alergia/histamina elevada.
- Considerar fitoterápicos com evidência (ex.: erva de São João) conforme perfil do paciente e interações.
- Suplementações basais com lógica bioquímica (magnésio, ômega 3), especialmente em gestação e em contextos de baixa ingestão habitual.

---

### Chunk 21/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte VIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.609

# Abordagem Funcional Integrativa Aplicada a Cada Área - Parte VIII

**Source:** https://web.plaud.ai/share/d1d91763827732043::YXdzOnVzLXdlc3QtMg

---

## Lecture

> Data e Hora: 2025-11-17 14:39:12
> Local: [Inserir Local]
> Instrutor: [Inserir Nome]
## 📝 Resumo
A aula destaca a centralidade do sistema gastrointestinal na saúde integral e na prática clínica em diversas especialidades, propondo medicina personalizada baseada em metabolômica, microbioma e avaliação de nutrientes, em contraste com diagnósticos de exclusão como a síndrome do intestino irritável (IBS). Apresenta evidências de subdiagnóstico (ex.: doença celíaca), correlações entre disbiose, hiperpermeabilidade intestinal (leaky gut) e múltiplas condições crônicas (neurológicas, dermatológicas, autoimunes), e a importância de nutrientes (vitamina D, zinco, selênio, ômega-3, B12) para imunidade inata e adaptativa.

---

### Chunk 22/30
**Article:** Modulação Intestinal I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.606

ade da barreira intestinal.
*   **Papel da Gliadina e da Zonulina**: A gliadina (fração do glúten) mal digerida pode se ligar a um receptor (CXCR3) no enterócito, estimulando a liberação de **zonulina**. A zonulina, por sua vez, bloqueia o receptor de fator de crescimento epidérmico (EGF), prejudicando a renovação celular e causando a ruptura das tight junctions.
*   **Desenvolvimento de Autoimunidade**: A ruptura das junções causa a permeabilidade intestinal ("leaky gut"), permitindo que macromoléculas (como frações de glúten) passem para a corrente sanguínea. Isso pode iniciar um processo de autoimunidade por **mimetismo proteico**, onde o sistema imune confunde proteínas do próprio corpo com as invasoras. Um exemplo clássico é a **tireoidite de Hashimoto**, onde anticorpos são produzidos contra a tireoperoxidase (anti-TPO) devido à semelhança estrutural com frações do glúten. O tratamento deve focar na recuperação da saúde intestinal.

---

### Chunk 23/30
**Article:** Intolerâncias, Alergias e Hipersensibilidades Alimentares II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.603

ambos → ~80%.
- Tipo de parto, aleitamento materno exclusivo e uso precoce de mamadeira.
- Exame físico: dor à palpação da fossa ilíaca direita pode sugerir inflamação em placas de Peyer.
- Colonoscopia: hiperplasia nodular linfóide no íleo terminal (achado associado, não específico).
- Manifestações orais: aftas, prurido faríngeo, dermatite perioral, língua geográfica; palidez facial sem anemia; refluxo, vômitos, dor abdominal crônica recorrente.
### 13. Sinais e sintomas associados à alergia alimentar (ampliados)
- Gastrointestinais: constipação, sangramento oculto, anemia, diarreia, má absorção de gordura.
- Respiratórios: broncoespasmo, coriza, tosse.
- Cutâneos: dermatite atópica, urticária, angioedema, palidez facial.
- Neurológicos/comportamentais: hiperatividade, déficit de atenção.
### 14.

---

### Chunk 24/30
**Article:** Glúten (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.602

*   **A dieta de eliminação é um teste válido, mas com ressalvas.** Experimentar uma dieta sem glúten por algumas semanas pode ajudar a identificar sensibilidades. No entanto, é crucial não demonizar o glúten, pois os grãos que o contêm são fontes importantes de fibras, vitaminas e minerais. O equilíbrio é fundamental.
*   **Fontes ocultas de glúten são comuns.** O glúten é usado como um "adesivo biológico" em produtos como sorvetes, molhos, cosméticos e batons, representando uma fonte significativa de exposição.
### 2. Espectro das Desordens Relacionadas ao Glúten
*   **Existe um espectro de reações.** As desordens incluem a doença celíaca (autoimune), a alergia ao trigo (mediada por IgE ou não) e a sensibilidade ao glúten não celíaca. O desafio clínico é identificar o problema e classificar o paciente corretamente.

---

### Chunk 25/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte VIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.601

l, questionou diagnósticos de exclusão como a Síndrome do Intestino Irritável (SII) e defendeu avaliação causal personalizada (microbioma, metabolômica, intolerâncias, marcadores). Introduziu intolerância à histamina e hiperpermeabilidade intestinal (leaky gut), conectando-as a desfechos imunológicos e sistêmicos. Discutiu o papel dos nutrientes e da vitamina D na imunidade inata e adaptativa, enfatizando a necessidade de monitoramento e personalização dinâmica.
## Conteúdo a Tratar
- Aula detalhada com o professor Alessio (celíaca, hiperpermeabilidade intestinal, novos marcadores)
- Aula do Dr.

---

### Chunk 26/30
**Article:** Infancy Dietary Patterns, Development, and Health: An Extensive Narrative Review (2022)
**Journal:** Children (Basel)
**Section:** other | **Similarity:** 0.601

y representative survey. J. Med. Internet Res. 2014, 16, e128. [CrossRef]
[PubMed]
201. Mowat, A.M. Coeliac disease–a meeting point for genetics, immunology, and protein chemistry. Lancet 2003, 361, 1290–1292.
[CrossRef]
202. Penagini, F.; Dilillo, D.; Meneghin, F.; Mameli, C.; Fabiano, V.; Zuccotti, G.V. Gluten-free diet in children: An approach to a
nutritionally adequate and balanced diet. Nutrients 2013, 5, 4553–4565. [CrossRef]
203. Ciacci, C.; Cirillo, M.; Cavallaro, R.; Mazzacca, G. Long-term follow-up of celiac adults on gluten-free diet: Prevalence and
correlates of intestinal damage. Digestion 2002, 66, 178–185. [CrossRef]
204. Niewinski, M.M. Advances in celiac disease and gluten-free diet. J. Am. Diet. Assoc. 2008, 108, 661–672. [CrossRef]
205. Niland, B.; Cash, B.D. Health Benefits and Adverse Effects of a Gluten-Free Diet in Non-Celiac Disease Patients. Gastroenterol.
Hepatol. 2018, 14, 82–91.
206. Gaesser, G.A.; Angadi, S.S.

---

### Chunk 27/30
**Article:** Microbioma Intestinal III (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.599

reática fecal: 85 — baixa; sugere insuficiência pancreática exócrina leve/moderada, possivelmente secundária a hipocloridria e disfunção digestiva global.
  - Zonulina fecal: 7 (normal < 80) — normal; reduz evidência laboratorial de hiperpermeabilidade via este marcador específico.
- Comentários:
  - Recomendada correlação com parâmetros sanguíneos (PCR, VHS) para reforçar inflamação sistêmica.
  - Colonoscopia citada como método de rastreio em adultos; não indicada para criança neste contexto.
- Mecanismos fisiopatológicos discutidos:
  - Dano a junções estreitas (claudina, ocludina, actina) por dieta (ex.: glúten).
  - Reconhecimento de MAMPs por TLR em células epiteliais; apresentação antigênica por células dendríticas/M e ativação de resposta T.
  - Células de Paneth: estimuladas por IL-22 e beta-glucana; produção de defensinas.
  - Células caliciformes (Goblet): síntese de mucina, principal fator antimicrobiano no cólon.

---

### Chunk 28/30
**Article:** Ácidos Graxos Monoinsaturados II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.594

er carnes pela saciedade/proteína.
- [ ] 5. Em autoimunes: propor teste de 1 mês de dieta vegana com acompanhamento nutricional; posteriormente transicionar para mediterrâneo ajustado com mais peixes/frutos do mar, mantendo nuts.
- [ ] 6. Para vegetarianos com autoimunes refratários que topem: testar dieta carnívora com suporte de enzimas digestivas e medidas para ácido gástrico (espinheira santa, betaína HCl, aloe vera, limão, vinagre), monitorando digestibilidade na primeira semana.
- [ ] 7. Documentar intervenções e anotar para aplicação imediata em consultório, evitando depender de memória.
- [ ] 8. Desenvolver criticismo científico: ao analisar estudos, verificar população, desenho, momento da avaliação de nutrientes e evitar extrapolações para recomendações universais.

---

### Chunk 29/30
**Article:** Intolerâncias, Alergias e Hipersensibilidades Alimentares II (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.593

imunoglobulinas, fenotipagem linfocitária, testes cutâneos e marcadores fecais), e princípios de manejo como dietas de eliminação, modulação da microbiota, probióticos, nutrientes e compostos fenólicos. Destacou-se a importância da digestibilidade das proteínas, da integridade da barreira intestinal e de equipe multidisciplinar no manejo.
## Conteúdo Não Coberto
1. Testes diagnósticos específicos por tipo de alergia (detalhamento prometido posteriormente)
2. Detalhamento de exames laboratoriais e complementares em protocolos formais
3. Estratégias terapêuticas e modulação intestinal em protocolos práticos padronizados
4. Outros nutrientes além da vitamina A na tolerância oral (serão apresentados futuramente)
5. Discussão aprofundada de hipersensibilidade tipo III e IV aplicadas à alergia alimentar com exemplos
6. Provas dietéticas/terapêuticas com passos práticos e segurança
7.

---

### Chunk 30/30
**Article:** Disbiose II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.592

da pressão arterial.
- [ ] 2. Ao avaliar um paciente, investigar o nível de estresse, histórico de uso de medicamentos (antibióticos, prazois, anticoncepcionais), tipo de parto, aleitamento e hábitos alimentares.
- [ ] 3. Considerar o exame coprológico funcional como ferramenta principal para diagnosticar disbiose e problemas de digestibilidade.
- [ ] 4. Priorizar a melhoria da eficiência digestiva (com enzimas, mastigação) e o controle do estresse como primeiros passos no tratamento da disbiose, antes de prescrever probióticos.
- [ ] 5. Monitorar os níveis de vitaminas lipossolúveis (A, D, E, K) e B12 em pacientes com condições que afetam a absorção, como cirurgia bariátrica, doença celíaca ou disbiose.
- [ ] 6. Considerar a suplementação de zinco para otimizar a absorção de ácido fólico, dado que sua hidrólise é dependente deste mineral.
- [ ] 7.

---

