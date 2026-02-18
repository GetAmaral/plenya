# ScoreItem: Lactose

**ID:** `019bf31d-2ef0-77d9-aaa8-2ca4fad39f47`
**FullName:** Lactose (Alimentação - Histórico - Intolerâncias)

**Preparation Metadata:**
- Quality Grade: **EXCELLENT**
- Total Chunks: 30 de 13 artigos
- Avg Similarity: 0.634

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-77d9-aaa8-2ca4fad39f47`.**

```json
{
  "score_item_id": "019bf31d-2ef0-77d9-aaa8-2ca4fad39f47",
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

**ScoreItem:** Lactose (Alimentação - Histórico - Intolerâncias)

**30 chunks de 13 artigos (avg similarity: 0.634)**

### Chunk 1/30
**Article:** Carboidratos II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.751

Lactose e Intolerância à Lactose
- **Definição:** Açúcar do leite (dissacarídeo de galactose e glicose), cuja digestão depende da enzima lactase.
- **Causa da Intolerância:** Ausência ou insuficiência da lactase, uma condição com alta prevalência em diversas populações.
- **Sintomas e Diagnóstico:** Incluem náuseas, diarreia, gases e enxaquecas. O diagnóstico pode ser feito por teste genético ou teste de tolerância oral.
- **Diferenciação Crucial:** Não confundir intolerância à lactose com alergia às proteínas do leite ou intolerância à histamina.
- **Manejo Clínico:** Aconselhou-se a exclusão total de lácteos em caso de suspeita para avaliar também a sensibilidade às proteínas. A reposição enzimática foi mencionada como opção.
- **Implicações Sistêmicas:** A intolerância não tratada pode levar a disbiose, permeabilidade intestinal, inflamação e impactar condições como ansiedade e acne (eixo intestino-cérebro).
### 4.

---

### Chunk 2/30
**Article:** Carboidratos II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.737

ente de lactase para digestão. Teste oral avalia subida de glicose como proxy de digestão.
- Persistência de lactase: variabilidade por regiões/etnias (valores citados de forma generalizante); não persistência é fenótipo evolutivo comum.
- Sintomas: náuseas, diarreia, dor abdominal, gases, constipação, cefaleias, enxaquecas; testes incluem genética e intolerância oral, mas exclusão de lácteos por 3 semanas (não apenas “sem lactose”) é prática e reveladora.
### 4. Diferenciações críticas: lactose versus outras reações a lácteos
- Intolerância à lactose (carboidrato) é distinta de reatividade à histamina e sensibilidades/alergias às proteínas do leite (caseína etc.). Lácteos ricos em histamina podem causar sintomas GI e extra-GI.
- Má absorção de lactose altera fermentação e microbiota, podendo contribuir para disbiose e impactos sistêmicos (eixo intestino-cérebro e múltiplos sintomas).
### 5.

---

### Chunk 3/30
**Article:** Carboidratos II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.718

e altera fermentação e microbiota, podendo contribuir para disbiose e impactos sistêmicos (eixo intestino-cérebro e múltiplos sintomas).
### 5. Genética, microbioma e evidências relacionadas à lactase
- LCT: forte seleção evolutiva; genótipos interagem intensamente com microbioma; consumo de lactose pode elevar bifidobactérias (efeito bifidogênico).
- Evidências: achados genéticos/fisiológicos sugerem maior DMO/estatura em persistentes, mas ensaios não corroboram. Críticas a interpretações que confundem lactose com SII/FODMAP e subestimam impactos metabólicos/microbiota.
### 6. Manejo clínico da intolerância à lactose
- Estratégias: dieta pobre em lactose e idealmente exclusão de lácteos para avaliar sensibilidades proteicas; reposição de lactase (ex.: “Lacto” da Enzymedica; há equivalentes no Brasil) pode ajudar em casos selecionados.

---

### Chunk 4/30
**Article:** Carboidratos II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.668

s duráveis para manejo de carboidratos e estabilidade glicêmica.
- Priorize carga glicêmica para prever impacto real na glicose e insulina.
- Libere glicose lentamente para reduzir picos de insulina e proteger a homeostase.
- Insulina e glucagon operam em ciclo complementar de armazenamento e liberação de glicose.
### Lactose, Diagnóstico e Tratamento
Reoriente avaliação e manejo de laticínios com foco em fenótipo e sintomas amplos.
- Intolerância à lactose é fenótipo comum por baixa lactase e não doença.
- Não confunda intolerância com reações a histamina ou proteínas do leite.
- Teste exclusão de laticínios por 3 semanas e avalie respostas além do intestino.
- Enzimas de amplo espectro superam lactase isolada em intolerantes multifatoriais.
- Consumir lactose não digerida favorece disbiose e sensibiliza a FODMAPs.
### Açúcares, Frutose e Comportamento
Adapte escolhas para reduzir hiperpalatabilidade e risco hepato-metabólico.

---

### Chunk 5/30
**Article:** Intolerâncias, Alergias e Hipersensibilidades Alimentares II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.654

arcadores alérgicos.
- Dietas de eliminação graduais: 2 alimentos (laticínios e glúten), 4 alimentos (glúten, laticínios, soja e frutos do mar) e 6 alimentos; maior restrição pode alterar a resposta clínica, orientando estratégias individualizadas.
**Achados de coocorrência e sensibilização cruzada ampliam o escopo clínico da avaliação.**
- Síndrome de alergia alimentar relacionada ao látex ocorre em até 50% dos pacientes com alergia ao látex, indicando alta coocorrência e sensibilização cruzada.
**Outras Constatações Importantes**
- Plaquetas acima de 400.000 podem estar relacionadas à enteropatia inflamatória crônica, servindo como achado laboratorial sugestivo.
- A frutose é descrita como absorvida via GLUT4, explicando possíveis quadros de má absorção e reações não imunológicas que imitam alergia.

---

### Chunk 6/30
**Article:** Carboidratos II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.654

reconhecendo diferença pequena e desaconselhando sucos.
- [ ] 5. Fórmulas infantis: preferir lactose em vez de maltodextrina para reduzir doçura e IG e evitar condicionamento hedônico ao doce.
- [ ] 6. Avaliar microbioma e sinais de disbiose em consumidores recorrentes de lácteos/lactose com sintomas sistêmicos (ansiedade, dores crônicas, dermatológicas).
- [ ] 7. Em dietas com leguminosas (veg/vegan): monitorar tolerância a rafinose/estaquiose, ajustar porções e preparo para minimizar fermentação/gases.
- [ ] 8. Educar sobre distinção entre intolerância à lactose, reatividade à histamina e sensibilidades/alergias às proteínas do leite.
- [ ] 9. Estimar ingestão diária de frutose (≤50 g; 50–100 g; >100 g) e ajustar conforme individualidade.
- [ ] 10. Planejar posicionamento de frutas mais calóricas/maior impacto glicêmico (banana, mamão) para horários estratégicos (ex.: final da tarde) e preferir consumo após proteína.
- [ ] 11.

---

### Chunk 7/30
**Article:** Intolerâncias, Alergias e Hipersensibilidades Alimentares II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.644

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
**Article:** Carboidratos II (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.641

” da Enzymedica; há equivalentes no Brasil) pode ajudar em casos selecionados.
- Cautela com “tolerância sintomática”: doses toleradas (≥12 g lactose ≈ 250 ml leite) podem não gerar sintomas GI, mas manter disbiose e efeitos sistêmicos; FODMAP pode aliviar sem remover causa e, em alguns casos, piorar constipação.
### 7. Comparativos: lactose, açúcar de mesa e escolhas para infância
- Lactose: IG ≈ 45, doçura ≈ 15% (bem menos doce) versus açúcar de mesa (IG ≈ 65). Açúcar de coco tem IG inferior, porém diferenças são pequenas no contexto geral.
- Fórmulas infantis: preferir lactose a maltodextrina por menor IG/doçura, evitando condicionamento hedônico ao doce e hiperativação dopaminérgica; alinhar ao padrão do leite materno.
### 8.

---

### Chunk 9/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.641

intestinal.
-   **SIFO:** Sensação de "blurring" (efeito baiacu - sentir-se distendido mesmo sem distensão objetiva) e dor abdominal.
## Objetivo:
A palestra descreve os exames e achados objetivos para a avaliação de pacientes com suspeita de Síndrome do Intestino Irritável (SII), embora não contenha achados de exame físico de um paciente específico.
-   **Exames Laboratoriais Gerais:** Hemograma e marcadores de atividade inflamatória.
-   **Calprotectina Fecal:** Usada para descartar doença inflamatória intestinal. Um valor abaixo de 100 µg/g tem uma positividade de 98% para o diagnóstico de SII. Entre 100-250 µg/g é uma zona cinzenta. Acima de 250 µg/g requer colonoscopia.
-   **Avaliação para Doença Celíaca:** Dosagem de IgA sérica total e anticorpo antitransglutaminase IgA. Se a IgA for baixa, dosar IgG sérica e antitransglutaminase IgG.
-   **Teste Parasitológico de Fezes:** Recomendado no contexto do Brasil.

---

### Chunk 10/30
**Article:** Disbiose I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.630

] 4. Solicitar elastase pancreática fecal para investigar insuficiência exócrina pancreática, especialmente em sintomas de má digestão de carboidratos/gorduras.
- [ ] 5. Investigar e tratar SIBO/SIFO/parasitoses (ex.: giardia) em pacientes com intolerâncias a dissacarídeos (lactose) e sintomas de má absorção; restaurar a integridade da mucosa.
- [ ] 6. Revisar a qualidade da dieta do paciente, enfatizando que energia e nutrientes vêm do alimento; alinhar a ingestão para atender cerca de 30 kcal/kg/dia quando apropriado ao estado basal.
- [ ] 7. Educar sobre a importância da saliva e da fase oral da digestão; evitar comer sob ansiedade/pressa, sentar para as refeições e focar no ato de comer.
- [ ] 8. Implementar estratégias para reduzir inflamação crônica de baixo grau, incluindo melhora da microbiota intestinal e redução de “garbage aging” por meio de suporte digestivo e antioxidante.
- [ ] 9.

---

### Chunk 11/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XVI (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.629

e microbioma intestinal.
    - **Avaliação de Alergias Cutâneas:** Realizar testes de exclusão alimentar (ex: remover laticínios/glúten por um mês) e considerar testes de intolerância alimentar por IgG. Investigar a causa de problemas com laticínios (lactose, proteína ou histamina).
    - **Avaliação Geral:** Avaliar marcadores inflamatórios, eixo HPA (estresse) e realizar uma avaliação hormonal completa.
- **Plano de Tratamento de Acompanhamento:**
    - **Intervenção Dietética:** Implementar uma dieta de eliminação personalizada com base nos resultados dos testes, removendo alimentos reativos (por exemplo, classe 4 no teste de IgG) por 2-3 meses para controlar a inflamação e a resistência à insulina.
    - **Saúde Intestinal:** Melhorar o microbioma e a integridade da barreira intestinal através de dieta, fibras e probióticos.

---

### Chunk 12/30
**Article:** Trato Gastrointestinal III – estômago – hipocloridria (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.629

de Precisão: FODMAPs e Histamina
- **Dieta FODMAP:** Indicada para pacientes com hipocloridria e excesso de gases, pois a fermentação intestinal pode alterar a produção de ácido gástrico. Polióis (xilitol, eritritol) podem ser gatilhos.
- **Sensibilidade à Histamina:** Um diagnóstico diferencial importante, com sintomas como coceiras, dor de cabeça e rinite. Alimentos ricos em histamina incluem atum, fermentados, lácteos e abacaxi.
### 4. Tratamentos e Suplementos para Hipocloridria
- **Cloridrato de Betaína (Betaína HCL):** Usado para reverter a hipocloridria. A dosagem varia de 300mg a 1500mg, tomada com a primeira garfada da refeição, preferencialmente em comprimidos.
- **Aloe Vera:**
    - **Benefícios:** O gel da *Aloe barbadensis Miller* possui mais de 75 compostos bioativos com ação anti-inflamatória, cicatrizante, antioxidante e imunomoduladora.

---

### Chunk 13/30
**Article:** Carboidratos II (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.628

digerida favorece disbiose e sensibiliza a FODMAPs.
### Açúcares, Frutose e Comportamento
Adapte escolhas para reduzir hiperpalatabilidade e risco hepato-metabólico.
- Prefira frutas inteiras a sucos para modular saciedade e reduzir carga de frutose.
- Excesso de frutose favorece lipogênese hepática e resistência insulínica.
- Lactose é preferível a maltodextrina em fórmulas infantis para evitar hiperpalatabilidade.
- Açúcar de coco tem benefício glicêmico marginal, sobretudo fora de líquidos.

---

## Teaching Note

Data e Hora: 2025-11-17 17:09:49
Local: [Inserir Local]
Aula: Carboidratos e Nutrição Aplicada à Prática Clínica
## Visão Geral
A aula abordou o metabolismo de carboidratos, com foco inicial na regulação hormonal pela insulina e glucagon. Foram analisados açúcares específicos como o açúcar de coco e a lactose, com uma discussão aprofundada sobre a intolerância à lactose, suas implicações clínicas, diagnóstico e manejo.

---

### Chunk 14/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa III (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.627

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

### Chunk 15/30
**Article:** Reevaluating our understanding of lactulose breath tests by incorporating hydrogen sulfide measurements (2019)
**Journal:** JGH Open: An Open Access Journal of Gastroenterology and Hepatology
**Section:** discussion | **Similarity:** 0.624

roenterol.2015;31:130–6.27YaoCK,TuckCJ.Theclinicalvalueofbreathhydrogentesting.J.Gastroenterol.Hepatol.2017;32:20–2.28NewmanA.Breath-analysistestsingastroenterology.Gut.1974;15:308–23.29DiggoryRT,CuschieriA.Theeffectofdoseandosmolalityoflactu-loseontheoral-caecaltransittimedeterminedbythehydrogenbreathtestandthereproducibilityofthetestinnormalsubjects.Ann.Clin.Res.1985;17:331–3.ABirgetal.ReevaluatinglactulosebreathtestsJGHOpen:Anopenaccessjournalofgastroenterologyandhepatology3(2019)228–233©2019TheAuthors.JGHOpen:AnopenaccessjournalofgastroenterologyandhepatologypublishedbyJournalofGastroenterologyandHepatologyFoundationandJohnWiley&SonsAustralia,Ltd.233

---

### Chunk 16/30
**Article:** Intolerâncias, Alergias e Hipersensibilidades Alimentares II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.622

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

### Chunk 17/30
**Article:** Intolerância à Histamina (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.621

plexidade do diagnóstico, que envolve a exclusão de outras condições, testes sanguíneos para a enzima DAO e análise de polimorfismos genéticos. A principal estratégia de tratamento é a "Food First", focando numa dieta baixa em histamina com acompanhamento nutricional, seguida pela suplementação da enzima DAO e, se necessário, o uso de medicamentos bloqueadores de receptores de histamina. A saúde intestinal, incluindo a disbiose e a hiperpermeabilidade (leaky gut), é destacada como um fator crucial que influencia a severidade da intolerância.
## 🔖 Knowledge Points
### 1. Introdução à Histamina e Condições Relacionadas
*   **Relevância Crescente da Histamina**
    *   A palestra aborda dois temas cada vez mais discutidos: a intolerância à histamina e a síndrome de ativação mastocitária.
    *   É crucial diferenciar a intolerância à histamina de outras condições relacionadas.

---

### Chunk 18/30
**Article:** Intolerância à Histamina (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.618

erância à histamina, realizar o diagnóstico diferencial para excluir condições como síndrome de ativação mastocitária e alergias alimentares.
- [ ] 3. Implementar uma dieta baixa em histamina com acompanhamento de um nutricionista como primeira linha de tratamento ("Food First").
- [ ] 4. Considerar a suplementação com a enzima DAO 20 minutos antes das refeições para controle dos sintomas.
- [ ] 5. Avaliar e tratar a saúde intestinal, investigando a presença de hiperpermeabilidade (leaky gut) e disbiose com bactérias estaminogênicas.
- [ ] 6. Avaliar a necessidade de reposição de cofatores da enzima DAO (cobre, vitamina C, vitamina B6).
- [ ] 7. Pausar o vídeo para observar a lista de medicamentos (antidepressivos, anti-hipertensivos, antibióticos) que podem diminuir a atividade da enzima DAO.
- [ ] 8. Utilizar bloqueadores de receptores H1 e H2 como terapia sintomática quando necessário.
- [ ] 9.

---

### Chunk 19/30
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

### Chunk 20/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XVI (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.612

eia.
- **Sintomas Neurológicos/Gerais:** Dores de cabeça (relacionadas à sinusite), enxaquecas (migraine), zumbido, fadiga após comer, fadiga crônica.
- **Sintomas de Intolerância:** Coceira após consumir alimentos ricos em histamina (laticínios, pimentão, berinjela, abacate), sintomas de intolerância à lactose.
## Objetivo:
O transcrito é uma palestra médica e não contém os exames de um paciente específico. Discute vários exames e achados objetivos para diagnosticar as causas subjacentes de condições dermatológicas e sistêmicas:
- **Testes Laboratoriais Sugeridos:**
    - Teste de IgG para alimentos para avaliar reações tardias (menciona laboratórios como SYNLAB e Testify).
    - Teste de atividade da DAO (diamina oxidase) para avaliar a intolerância à histamina.
    - Teste de microbioma intestinal.
    - Teste de tolerância à lactose.
    - Avaliação de resistência à insulina.

---

### Chunk 21/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.612

ratório para SIBO/IMO:** Considerado o método prático padrão, utilizando lactulose ou glicose.
    -   **Critério para SIBO (Hidrogênio):** Elevação acima de 20 ppm nos primeiros 90 minutos.
    -   **Critério para IMO (Metano):** Elevação acima de 10 ppm em qualquer momento do teste.
-   **Diagnóstico de SIFO:** O padrão ouro é o aspirado duodenal com cultura (>10³ UFC/ml de fungos), mas é raramente realizado. A avaliação da micobiota e metabólitos fúngicos pode ser útil.
-   **Outros Achados:** Ativação do eixo HPA e reação de "die-off" (com manifestações cutâneas) durante tratamento antifúngico.
## Diagnóstico Primário:
-   **Avaliação:** O texto é uma discussão abrangente sobre a Síndrome do Intestino Irritável (SII), suas causas subjacentes (aumento da permeabilidade intestinal, disbiose, inflamação) e comorbidades como distúrbios do sono, SIBO, IMO e SIFO.

---

### Chunk 22/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa III (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.610

intestinal (ex: zonulina fecal).
    -   Avaliar a qualidade do sono, histórico de traumas e estresse.
-   **Plano de Tratamento de Acompanhamento:**
    -   **Ajuste Alimentar:** Implementar o protocolo Low FODMAP em três fases (suspensão, reintrodução, personalização) com apoio de nutricionista. Evitar emulsificantes, álcool em excesso, frutose e glicose. Considerar boas fontes de gordura, carotenoides, vitamina D e curcumina.
    -   **Otimização da Digestão:** Avaliar e corrigir a digestão, podendo incluir suplementação de ácido clorídrico, enzimas digestivas ou fibras (com cautela).
    -   **Probióticos:** Usar com cautela (menos cepas, menor quantidade, menor tempo), pois podem piorar sintomas como "brain fogginess" em pacientes com D-lactato elevado.
    -   **Modificações no Estilo de Vida:** Atividade física moderada, técnicas de respiração diafragmática, banhos gelados e otimização do sono.

---

### Chunk 23/30
**Article:** Diagnosis and Management of Fructose Malabsorption (2021)
**Journal:** Journal of Clinical Gastroenterology
**Section:** abstract | **Similarity:** 0.610

Fructose malabsorption affects 30-40% of the population and can cause significant gastrointestinal symptoms. Hydrogen breath testing after 25-50g fructose load is the diagnostic gold standard. A rise in hydrogen ≥20 ppm above baseline indicates fructose malabsorption. The test requires proper preparation including 12-hour fasting, avoiding high-fiber foods the previous day, and no smoking on test day. Positive tests should guide dietary management with low-FODMAP approaches.

---

### Chunk 24/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa III (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.608

minase IgA. Se a IgA for baixa, dosar IgG sérica e antitransglutaminase IgG.
-   **Teste Parasitológico de Fezes:** Recomendado no contexto do Brasil.
-   **Avaliação da Permeabilidade Intestinal:** O aumento da permeabilidade (leaky gut) pode ser avaliado pela zonulina (fecal ou sérica). Menciona-se que o estresse (injeção de CRH) pode induzir um aumento nos marcadores de leaky gut.
-   **Avaliação da Microbiota/Metabolômica:** A avaliação isolada da microbiota é considerada de pouco valor. A avaliação da metabolômica (ex: ácidos orgânicos urinários) é mais útil para avaliar a função da microbiota e detectar metabólitos bacterianos e fúngicos. O aumento do D-lactato no sangue pode estar associado ao uso de probióticos e causar "brain fogginess".
-   **Teste Respiratório para SIBO/IMO:** Considerado o método prático padrão, utilizando lactulose ou glicose.
    -   **Critério para SIBO (Hidrogênio):** Elevação acima de 20 ppm nos primeiros 90 minutos.

---

### Chunk 25/30
**Article:** Carboidratos II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.606

(exemplos práticos, vias), depois avançar para proteínas e, em seguida, lipídios. Materiais de apoio com figuras e PDFs acompanharão as aulas.
## ❓ Questions
- [Inserir Pergunta/Confusão]
## 📚 Assignments
- [ ] 1. Reduzir carga glicêmica nas refeições, priorizando entrada gradual de carboidratos para diminuir picos de glicose/insulina.
- [ ] 2. Suspeita de intolerância à lactose: realizar exclusão de lácteos por 3 semanas, monitorando sintomas GI e extra-GI (rinite, acne, cefaleias, dores).
- [ ] 3. Considerar reposição enzimática (ex.: Lacto, Enzymedica; ou equivalentes nacionais) em casos selecionados, evitando manter consumo apenas “tolerável” sintomaticamente.
- [ ] 4. Revisar adoçantes: se inevitável em bebidas, preferir açúcar de coco pelo IG mais baixo, reconhecendo diferença pequena e desaconselhando sucos.
- [ ] 5. Fórmulas infantis: preferir lactose em vez de maltodextrina para reduzir doçura e IG e evitar condicionamento hedônico ao doce.

---

### Chunk 26/30
**Article:** Ácidos Graxos Saturados de Cadeia Curta (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.603

suplemento CoreBiome (tributirina), mencionando que, apesar do uso prolongado, não percebeu melhora significativa. Ele também observa que, em sua prática clínica, alguns pacientes relataram resultados positivos, enquanto duas pacientes apresentaram piora do quadro, com mal-estar e diarreia.
## Objetivo:
O texto é uma apresentação educacional e não contém achados de exames físicos ou laboratoriais de um paciente específico. A discussão aborda conceitos fisiológicos, resultados de estudos e a experiência clínica com os AGCC e suplementos relacionados, como:
-   Os AGCC (butirato, acetato, propionato) são produzidos pela microbiota intestinal através da fermentação de fibras dietéticas.
-   O butirato possui atividade anti-inflamatória, antimicrobiana, antidiarreica e melhora a hiperpermeabilidade intestinal.

---

### Chunk 27/30
**Article:** Microbioma Intestinal III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.602

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

### Chunk 28/30
**Article:** Intolerâncias, Alergias e Hipersensibilidades Alimentares II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.599

- **Marcadores Fecais:**
        - **Calprotectina Fecal:** Valor < 138 exclui alergia não mediada por IgE com alta sensibilidade.
        - **Pesquisa de Sangue Oculto:** Exame simples e sensível para enterocolite.
        - **Alfa-1 Antitripsina Fecal:** Avalia enteropatia perdedora de proteínas.
*   **Testes Específicos e Procedimentos:**
    - **Testes Cutâneos (com alergista):** Prick Test (padrão-ouro para alergia mediada por IgE) e Patch Test (para reações tardias).
    - **Diagnóstico Molecular (RAST, ImunoCAP):** Avalia IgE específica para determinados alérgenos.
    - **Teste de Provocação Oral:** Considerado padrão-ouro para confirmação, mas é arriscado e complexo.
    - **Testes de IgG:** Não devem ser usados de rotina para diagnóstico de alergia, pois podem indicar tolerância e não alergia.
    - **Endoscopia/Colonoscopia:** Podem revelar achados como hiperplasia nodular linfoide.
### 5.

---

### Chunk 29/30
**Article:** TDAH - Parte XIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.596

nalização e limites
   - Dietas padrão (ex.: Mediterrânea com vinho/queijo/molho de tomate) podem piorar pacientes específicos; personalizar por sintomas, fermentação, intolerâncias e objetivos.
   - Adesão é crucial: citação de Hipócrates “Antes de curar alguém, pergunta-lhe se está disposto a abandonar as coisas que lhe fizeram adoecer.” Sem mudança (ex.: manter vinho com histamina elevada), resultados limitados mesmo com antihistamínicos.
* Suplementos e escolhas
   - Suplementar quando dieta não alcança metas; usar inteligência na escolha de fontes (evitar exacerbar fermentação, histamina ou excitabilidade). Integração multiprofissional é necessária para orientar gestantes e pacientes em risco.

---

### Chunk 30/30
**Article:** Carboidratos II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.596

entre intolerância à lactose, reatividade à histamina e sensibilidades/alergias às proteínas do leite. Compara lactose ao açúcar de mesa (IG e doçura) e aborda escolhas para fórmulas infantis, preferindo lactose a maltodextrina por menor doçura/IG e influência sobre a memória afetiva do doce e dopamina em crianças. Introduz frutose tanto em formas complexas fermentáveis (rafinose, estaquiose em leguminosas) quanto como monossacarídeo de alta doçura/solubilidade (presente em frutas e mel), detalhando sua metabolização predominantemente hepática (frutoquinase, ausência de contrarregulação inicial) e destinos metabólicos (energia, lipogênese, glicose/glicogênio).

---

