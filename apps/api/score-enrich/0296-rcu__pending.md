# ScoreItem: RCU

**ID:** `019bf31d-2ef0-729d-a9df-4b53569231dc`
**FullName:** RCU (Histórico de doenças - Doenças crônicas (Questionar ativamente as doenças crônicas mais comuns, perguntar sobre duração, grau de controle e tratamentos utilizados no passado e atualmente) - Doenças auto-imunes)

**Preparation Metadata:**
- Quality Grade: **GOOD**
- Total Chunks: 30 de 14 artigos
- Avg Similarity: 0.577

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-729d-a9df-4b53569231dc`.**

```json
{
  "score_item_id": "019bf31d-2ef0-729d-a9df-4b53569231dc",
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

**ScoreItem:** RCU (Histórico de doenças - Doenças crônicas (Questionar ativamente as doenças crônicas mais comuns, perguntar sobre duração, grau de controle e tratamentos utilizados no passado e atualmente) - Doenças auto-imunes)

**30 chunks de 14 artigos (avg similarity: 0.577)**

### Chunk 1/30
**Article:** (Dr Otávio Freitas) Aula 02 - Vitamina D - Doenças Autoimunes (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.672

A apresentação expande para doença de Crohn e retocolite ulcerativa (RCU), alinhando observações clínicas do consultório a evidências publicadas: meta-análise de 55 estudos observacionais relaciona deficiência de vitamina D com essas condições; estudos sugerem que a vitamina D atenua a inflamação na RCU por ativar o receptor de vitamina D e modular a resposta NL-RPC; há menções sobre possíveis relações entre níveis de vitamina D e a extensão da doença. O orador cita um paciente acompanhado por cerca de sete anos com colonoscopia normal após tratamento. O depoimento de Juliano ilustra um percurso de 15 anos desde o diagnóstico por exames e cirurgia, com uma década de tratamentos convencionais e dor/desconforto persistentes.

---

### Chunk 2/30
**Article:** Microbioma Intestinal III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.637

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

### Chunk 3/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa III (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.635

**
        -   **Controle de Sintomas:** Cápsula de óleo de hortelã-pimenta (dor abdominal).
        -   **Tratamento SIFO/Die-off:** Saccharomyces boulardii (250 mg 2x/dia durante tratamento antifúngico), Cúrcuma longa (Golden Milk), Ácido Caprílico (Óleo de Coco).
        -   **Integridade Intestinal:** Zinco-carnosina, glutamina, pectina, beta-glucana, butirato.
        -   **Motilidade:** Magnésio, Trífala.
-   **Próximos Passos/Exames:**
    -   Realizar uma avaliação laboratorial completa (hemograma, marcadores inflamatórios, calprotectina fecal, testes para doença celíaca, parasitológico de fezes).
    -   Considerar testes funcionais como teste respiratório para SIBO/IMO e análise de ácidos orgânicos urinários (metabolômica).
    -   Avaliar a permeabilidade intestinal (ex: zonulina fecal).
    -   Avaliar a qualidade do sono, histórico de traumas e estresse.

---

### Chunk 4/30
**Article:** Microbioma Intestinal III (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.619

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

### Chunk 5/30
**Article:** Microbioma Intestinal III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.604

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

### Chunk 6/30
**Article:** Fisiologia e Bioquímica do Sistema Imune I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.590

vel apresentou 7,94%, indicando inflamação ativa.
- A presença de 21% da bactéria Prausnitzi, combinada com a ausência de Akkermansia, sugere uma dieta com carga glicêmica muito alta.
- A estratégia de tratamento inclui um protocolo de "sete passos para a reprogramação intestinal", que pode envolver dietas como a low FODMAP por um a dois meses.
- Para a modulação intestinal, são sugeridas tinturas em proporções específicas, como 50% de alcaçuz e 50% de cúrcuma.
**O tratamento é altamente personalizado, utilizando suplementos em doses específicas e protocolos dietéticos faseados para controlar a inflamação e modular a resposta imune.**
- Suplementos de curcumina devem ter alta concentração de curcuminoides (95% a 99%) para garantir eficácia.
- Para o controle de TH2, doses de N-acetilcisteína variam de 400 mg a 1000 mg, enquanto para a modulação de TH17, a berberina é usada em doses de 100 mg a 300 mg.

---

### Chunk 7/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.587

ratório para SIBO/IMO:** Considerado o método prático padrão, utilizando lactulose ou glicose.
    -   **Critério para SIBO (Hidrogênio):** Elevação acima de 20 ppm nos primeiros 90 minutos.
    -   **Critério para IMO (Metano):** Elevação acima de 10 ppm em qualquer momento do teste.
-   **Diagnóstico de SIFO:** O padrão ouro é o aspirado duodenal com cultura (>10³ UFC/ml de fungos), mas é raramente realizado. A avaliação da micobiota e metabólitos fúngicos pode ser útil.
-   **Outros Achados:** Ativação do eixo HPA e reação de "die-off" (com manifestações cutâneas) durante tratamento antifúngico.
## Diagnóstico Primário:
-   **Avaliação:** O texto é uma discussão abrangente sobre a Síndrome do Intestino Irritável (SII), suas causas subjacentes (aumento da permeabilidade intestinal, disbiose, inflamação) e comorbidades como distúrbios do sono, SIBO, IMO e SIFO.

---

### Chunk 8/30
**Article:** Microbioma Intestinal III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.581

s na Consulta**
    *   A avaliação deve seguir a ordem: 1. História Pregressa, 2. História Clínica, 3. Medicamentos, 4. Hábitos Alimentares, 5. Exercícios Físicos. A modulação gastrointestinal deve ser priorizada.
*   **Biointestil (Suplemento)**
    *   Composto por óleo essencial de *Cymbopogon martinii* e gengibre, com ação antimicrobiana seletiva, anti-inflamatória e carminativa, liberado principalmente no cólon.
    *   Pode causar a reação de Jarisch-Herxheimer (piora inicial dos sintomas).
*   **Terapias Alternativas para o Intestino**
    *   **Hidrocolonoterapia:** Limpeza do intestino grosso com água ozonizada, mencionada como benéfica para constipação crônica e inflamação.
    *   **Enema de Café:** Terapia que visa ativar a desintoxicação hepática (glutationa S-transferase) e melhorar o fluxo biliar.
## ❓ Perguntas
- [Inserir Pergunta/Confusão]
## 📚 Tarefas
- [ ] 1.

---

### Chunk 9/30
**Article:** Ácidos Graxos Saturados de Cadeia Curta (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.573

o de sódio.
    *   **Posologia:** A dose sugerida é de 3mg, de uma a três vezes ao dia, junto às refeições.
    *   **Experiência Clínica e Custo:** É um suplemento caro com resultados variáveis. Alguns pacientes melhoram, mas outros podem apresentar piora (mal-estar, diarreia).
    *   **Recomendação de Uso:** Deve ser considerado após tentativas de modulação endógena. A prescrição deve incluir um período de teste (ex: dois meses) com monitoramento clínico para avaliar a real eficácia e justificar a manutenção. O objetivo é usá-lo como uma ferramenta temporária, não para dependência.
*   **Probióticos:** A prescrição deve ser individualizada, pois são considerados um "band-aid". O ideal é modular o sistema para que não sejam necessários.
## ❓ Perguntas
- [Inserir Pergunta/Confusão]
## 📚 Tarefas
- [ ] 1. Estudar como individualizar planos alimentares e tipos de fibras para otimizar a produção de AGCC.
- [ ] 2.

---

### Chunk 10/30
**Article:** Remissão do Lúpus Através da Medicina Funcional Integrativa (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.573

s, frutas e verduras, para recuperar a integridade intestinal
- [ ] Encaminhar o paciente para uma nutricionista funcional, para que haja uma correção na quantidade de carboidratos, proteínas e para ajudar na suplementação
- [ ] Pesquisar estressores crônicos comuns na consulta, com o paciente
- [ ] Revisar todos os itens de gestão do estresse, para não sobrecarregar a prescrição de medicamentos e suplementos
- [ ] Verificar o sono do paciente, para gestão do estresse
- [ ] Verificar se o paciente está realizando atividade física regularmente, para gestão do estresse
- [ ] Verificar se o paciente está socialmente ativo, para gestão do estresse
- [ ] Verificar se o paciente tem momentos de lazer, para gestão do estresse
- [ ] Verificar se o paciente consegue manter uma rotina organizada, para gestão do estresse
- [ ] Verificar se o paciente consegue gerenciar o tempo, para gestão do estresse
- [ ] Verificar se o paciente pratica técnicas de relaxamento, para gestão

---

### Chunk 11/30
**Article:** Disbiose II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.572

da pressão arterial.
- [ ] 2. Ao avaliar um paciente, investigar o nível de estresse, histórico de uso de medicamentos (antibióticos, prazois, anticoncepcionais), tipo de parto, aleitamento e hábitos alimentares.
- [ ] 3. Considerar o exame coprológico funcional como ferramenta principal para diagnosticar disbiose e problemas de digestibilidade.
- [ ] 4. Priorizar a melhoria da eficiência digestiva (com enzimas, mastigação) e o controle do estresse como primeiros passos no tratamento da disbiose, antes de prescrever probióticos.
- [ ] 5. Monitorar os níveis de vitaminas lipossolúveis (A, D, E, K) e B12 em pacientes com condições que afetam a absorção, como cirurgia bariátrica, doença celíaca ou disbiose.
- [ ] 6. Considerar a suplementação de zinco para otimizar a absorção de ácido fólico, dado que sua hidrólise é dependente deste mineral.
- [ ] 7.

---

### Chunk 12/30
**Article:** Disbiose II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.570

ta de aleitamento materno, excesso de sanitarização.
    - **Contaminantes:** Agrotóxicos, metais pesados, cloro, flúor.
    - **Envelhecimento:** Redução do ácido gástrico (*inflammaging*).
*   **Manifestações**
    - **Digestivas:** Distensão, refluxo, síndrome do intestino irritável, alteração do hábito intestinal.
    - **Extra-digestivas:** Alergias, doenças autoimunes, problemas de saúde mental, alterações hormonais.
### 5. Abordagem Terapêutica e Diagnóstico
*   **Diagnóstico**
    - Primariamente clínico, baseado nos sintomas e exame físico.
    - O **exame coprológico funcional** é uma ferramenta chave para avaliar a digestibilidade e o comportamento da microbiota.
*   **Estratégias de Tratamento (Hierarquia)**
    1.  **Melhorar a Eficiência Digestiva:** É o primeiro passo. Inclui mastigação, *mindful eating* e uso de enzimas digestivas (suplementos como pancreatina ou alimentos como mamão e abacaxi).
    2.

---

### Chunk 13/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.570

intestinal.
-   **SIFO:** Sensação de "blurring" (efeito baiacu - sentir-se distendido mesmo sem distensão objetiva) e dor abdominal.
## Objetivo:
A palestra descreve os exames e achados objetivos para a avaliação de pacientes com suspeita de Síndrome do Intestino Irritável (SII), embora não contenha achados de exame físico de um paciente específico.
-   **Exames Laboratoriais Gerais:** Hemograma e marcadores de atividade inflamatória.
-   **Calprotectina Fecal:** Usada para descartar doença inflamatória intestinal. Um valor abaixo de 100 µg/g tem uma positividade de 98% para o diagnóstico de SII. Entre 100-250 µg/g é uma zona cinzenta. Acima de 250 µg/g requer colonoscopia.
-   **Avaliação para Doença Celíaca:** Dosagem de IgA sérica total e anticorpo antitransglutaminase IgA. Se a IgA for baixa, dosar IgG sérica e antitransglutaminase IgG.
-   **Teste Parasitológico de Fezes:** Recomendado no contexto do Brasil.

---

### Chunk 14/30
**Article:** Ácidos Graxos Saturados de Cadeia Curta (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.569

suplemento CoreBiome (tributirina), mencionando que, apesar do uso prolongado, não percebeu melhora significativa. Ele também observa que, em sua prática clínica, alguns pacientes relataram resultados positivos, enquanto duas pacientes apresentaram piora do quadro, com mal-estar e diarreia.
## Objetivo:
O texto é uma apresentação educacional e não contém achados de exames físicos ou laboratoriais de um paciente específico. A discussão aborda conceitos fisiológicos, resultados de estudos e a experiência clínica com os AGCC e suplementos relacionados, como:
-   Os AGCC (butirato, acetato, propionato) são produzidos pela microbiota intestinal através da fermentação de fibras dietéticas.
-   O butirato possui atividade anti-inflamatória, antimicrobiana, antidiarreica e melhora a hiperpermeabilidade intestinal.

---

### Chunk 15/30
**Article:** Modulação Intestinal II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.569

homocisteína, proteína C-reativa.
    - **Inflamação intestinal:** Calprotectina fecal.
    - **Risco cardiovascular:** TMAO sérico (em pacientes com resistência à insulina).
    - **Saúde geral:** Níveis de vitamina B12, cuja absorção depende de um pH gástrico adequado.
## ❓ Perguntas
- [Inserir Pergunta/Confusão]
## 📚 Tarefas
- [ ] 1. Ao solicitar o exame de calprotectina fecal, justificar como "suspeita de doença inflamatória intestinal" para aumentar a chance de aprovação pelo plano de saúde.
- [ ] 2. Orientar pacientes adultos a coletar a calprotectina fecal em um dia de rotina alimentar normal (ex: quarta-feira), evitando períodos pós-excessos para não gerar falsos positivos.
- [ ] 3. Considerar a dosagem de elastase fecal para avaliar a função exócrina do pâncreas antes de prescrever enzimas digestivas.
- [ ] 4. Monitorar marcadores de inflamação de baixo grau, como resistência à insulina (HOMA-IR), homocisteína e proteína C-reativa.

---

### Chunk 16/30
**Article:** Disbiose II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.564

ociar pancreatina com betaína HCl na mesma cápsula; timing: betaína HCl durante a refeição (liberação gástrica), pancreatina antes (T–15 min).
- Enzimas vegetais: maior labilidade; uso próximo à betaína; sachês misturados ao alimento (não diluir em água).
- Integração com nutricionistas e individualização dietética.
- Controle de estresse: psicoterapia e terapias complementares (ex.: privação sensorial).
- Suplementação conforme necessidade: aminoácidos, lipídios, complexo B, magnésio.
- Manejo da constipação e atividade física para motilidade.
### 17. Diagnóstico clínico e exames funcionais
- Valorização da queixa e exame físico: distensão, ruídos hidroaéreos, massas.
- Rastreio de deficiências nutricionais (ex.: ferro), doença celíaca e SIBO.
- Exame coprológico funcional: avaliação de digestibilidade, sobras alimentares, comportamento microbiano, produção de amônia e ácidos; interpretação integrada com quadro clínico.

---

### Chunk 17/30
**Article:** Microbioma Intestinal III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.564

tada.
  - Insuficiência pancreática exócrina funcional (elastase fecal baixa), provavelmente secundária a disfunção digestiva global/hipocloridria.
  - Zonulina normal no momento.
- Diagnóstico Suspeito:
  - DII (doença de Crohn/colite ulcerosa) versus infecção entérica bacteriana; doença celíaca como diferencial; perda proteica intestinal.
  - Nenhum adicional no momento.
## Plano:
- Prescrição:
  - Inserir mais aqui.
  - Biointestil: 600 mg/dia à noite; em fases iniciais considerar até 600 mg 3x/dia (manhã, antes do almoço e antes do jantar), com cautela pelo custo e possível reação de Jarisch-Herxheimer.
  - Gengibre: uso em extrato/shot e chás com ação anti-inflamatória e carminativa.
  - Berberina: alternativa com menor risco de piora inicial.
  - Lactoferrina: 500 mg para suporte imunológico, conforme avaliação pediátrica.

---

### Chunk 18/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa III (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.564

intestinal (ex: zonulina fecal).
    -   Avaliar a qualidade do sono, histórico de traumas e estresse.
-   **Plano de Tratamento de Acompanhamento:**
    -   **Ajuste Alimentar:** Implementar o protocolo Low FODMAP em três fases (suspensão, reintrodução, personalização) com apoio de nutricionista. Evitar emulsificantes, álcool em excesso, frutose e glicose. Considerar boas fontes de gordura, carotenoides, vitamina D e curcumina.
    -   **Otimização da Digestão:** Avaliar e corrigir a digestão, podendo incluir suplementação de ácido clorídrico, enzimas digestivas ou fibras (com cautela).
    -   **Probióticos:** Usar com cautela (menos cepas, menor quantidade, menor tempo), pois podem piorar sintomas como "brain fogginess" em pacientes com D-lactato elevado.
    -   **Modificações no Estilo de Vida:** Atividade física moderada, técnicas de respiração diafragmática, banhos gelados e otimização do sono.

---

### Chunk 19/30
**Article:** Microbioma Intestinal III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.563

ino e no cérebro.
    *   Essa redução pode causar ansiedade, depressão e distúrbios do sono.
### 4. Avaliação Clínica e Marcadores Fecais
*   **Avaliação Clínica**
    *   A **Escala de Bristol** avalia a forma das fezes (tipo 4 é o ideal).
    *   A distensão abdominal pós-prandial é um sinal clínico importante de fermentação excessiva.
*   **Marcadores de Inflamação e Permeabilidade (Exame Coprológico Funcional)**
    *   **Alfa-1-antitripsina:** Marcador de aumento da permeabilidade intestinal (leaky gut) e perda proteica. Níveis elevados indicam inflamação crônica.
    *   **Calprotectina Fecal:** Marcador de inflamação intestinal aguda ou crônica (valor ideal < 50). Níveis muito altos são comuns em distúrbios neurológicos e doenças inflamatórias intestinais.
    *   **Lactoferrina Fecal:** Glicoproteína liberada por neutrófilos durante a inflamação, confirmando um quadro inflamatório.

---

### Chunk 20/30
**Article:** (Dr Otávio Freitas) Aula 02 - Vitamina D - Doenças Autoimunes (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.563

independente nas atividades diárias, voltou à escola, faz educação física, anda a cavalo; recuperou força nas mãos para escovar os dentes e pintar.
#### **Doença de Crohn e Retocolite Ulcerativa**
* **Base científica:** Meta-análise de 55 estudos observacionais confirma associação entre deficiência de vitamina D e essas doenças.
* **Relato (Juliano):**
    * **Diagnóstico:** Doença de Crohn há 15 anos.
    * **Antes:** 10 anos de tratamentos convencionais sem sucesso, dor e desconforto constantes.
    * **Após iniciar o protocolo:** Melhora a partir do terceiro mês; vida normal, sem dor, recuperou peso; colonoscopia atual normal.
#### **Dermatite Atópica**
* **Base científica:** Revisões sistemáticas e meta-análises indicam correlação entre baixos níveis de vitamina D e maior gravidade, especialmente em crianças.
* **Relato (Letícia):**
    * **Diagnóstico:** Dermatite atópica desde o nascimento (32 anos atualmente).

---

### Chunk 21/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.563

ração do eixo cérebro-intestino-microbiota. O diagnóstico é clínico, baseado nos critérios de Roma 4, que exigem dor abdominal recorrente associada a alterações no hábito intestinal. A fisiopatologia envolve alterações no SNC, desequilíbrios da microbiota, fatores genéticos/epigenéticos e o papel de neurotransmissores como a serotonina.
- Diagnóstico Suspeito: Nenhum no momento
## Plano:
- Prescrição: Inserir mais aqui
- Próximos Passos/Exames:
    - A apresentação enfatiza a importância de considerar diagnósticos diferenciais, como constipação funcional e diarreia funcional.
    - É crucial investigar sinais de alarme, especialmente em pacientes com mais de 60 anos, para descartar doenças orgânicas como neoplasia de cólon.
    - Menciona abordagens terapêuticas gerais, como o uso de medicamentos que atuam em receptores de serotonina (5-HT) para modular a motilidade e a dor.

---

### Chunk 22/30
**Article:** Fisiologia e Bioquímica do Sistema Imune I (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.562

*   **Análise da Microbiota Intestinal**
    *   A diversidade bacteriana muda conforme a doença. Um exemplo prático de um paciente com Síndrome do Intestino Irritável mostrou predomínio de Firmicutes, alta contagem de Protobactérias, e ausência de *Akkermansia muciniphila*, indicando inflamação, risco de obesidade e diabetes tipo 2.
    *   A falta de bactérias como *Blautia* e *Coprococcus* (metabolizadoras de fibras) explica sintomas como distensão abdominal.
### 3. Gestão da Microbiota e Saúde Intestinal
*   **Disbiose, LPS e Permeabilidade Intestinal**
    *   A disbiose leva ao excesso de lipopolissacarídeos (LPS), que ao contactar a parede intestinal, deflagra uma resposta inflamatória via células dendríticas, aumentando a IL-6 e desequilibrando as células T (TREG, TH1, TH17).
    *   A baixa quantidade de *Akkermansia* enfraquece a barreira intestinal, causando dano direto aos enterócitos e facilitando a translocação de LPS.

---

### Chunk 23/30
**Article:** Intolerâncias, Alergias e Hipersensibilidades Alimentares II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.560

odem indicar tolerância e não alergia.
    - **Endoscopia/Colonoscopia:** Podem revelar achados como hiperplasia nodular linfoide.
### 5. Manejo e Tratamento
*   **Dietas de Eliminação:** Principal abordagem, consiste em retirar o alimento agressor. Deve ser feita com acompanhamento multidisciplinar para evitar déficits nutricionais, especialmente em crianças.
*   **Melhora da Digestão:** Uma digestão inadequada aumenta a carga de antígenos no intestino. O uso de enzimas digestivas pode ajudar a degradar melhor as proteínas e diminuir os sintomas. Fatores como pasteurização e Reação de Maillard podem aumentar a alergenicidade dos alimentos.
*   **Modulação Intestinal:** É o pilar do tratamento.
    - **Microbiota e AGCC:** Uma dieta rica em fibras aumenta a produção de ácidos graxos de cadeia curta (butirato, propionato), que são anti-inflamatórios e fortalecem a barreira intestinal.
    - **Probióticos:** Cepas específicas como *Lactobacillus rhamnosus*, *L.

---

### Chunk 24/30
**Article:** Emagrecimento - Parte IX (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.558

ra desequilíbrios como inflamação sistêmica e apoio metabólico para discussão na próxima aula.
- [ ] 4. Preparar uma lista de suplementos com evidências para emagrecimento e modulação de inflamação, com mecanismos e segurança.
- [ ] 5. Elaborar um plano alimentar focado em “alimento como remédio”, integrando abordagens anti-inflamatórias.
- [ ] 6. Solicitar exames de B12, vitamina D, zinco e cobre (cobre sérico com altas doses de zinco) e avaliar necessidade de selênio com base no consumo de castanhas-do-Pará.
- [ ] 7. Ajustar cromo para 200–300 mcg por refeição principal, priorizando adesão (permitir durante as refeições).
- [ ] 8. Implementar magnésio 200 mg à noite, preferencialmente com inositol e L-triptofano, visando relaxamento e suporte metabólico.
- [ ] 9. Decidir entre kelp (200–400 mcg) e Ascophyllum nodosum/I-plus (200–400 mg) para iodo, considerando potência, detox e carga de cápsulas.
- [ ] 10.

---

### Chunk 25/30
**Article:** Fisiologia e Bioquímica do Sistema Imune II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.557

60 e 90
- [ ] Manter a insulina, o mais baixo possível, 6, 7, estourando 8
- [ ] Avaliar a homocisteína, pois é um marcador inflamatório importante
- [ ] Usar a proteína C-reativa, associado com os níveis de homocisteína
- [ ] Verificar os parâmetros essenciais na avaliação inflamatória
- [ ] Estimar o índice de glicação e o índice TAIG, baseado nos resultados essenciais
- [ ] Complementar a avaliação com TNF-alfa, IL-6, glutationa e malon de aldeído
### Tarefas para @
- [ ] Usar um concentrado de C8 ou um mix de C8 e C10, para estimular mais ainda o CP3 e as UCPs (proteínas desacopladoras), diminuir a produção de espécie reativa de oxigênio e aumentar a oxidação de gordura @
- [ ] Incluir mioinositol, trans-resveratrol e epigalocatequina galato na formulação, para diminuir os compostos de glicação avançada e a hemoglobina glicada @
- [ ] Fazer uma boa distribuição de gordura e trabalhar os ácidos graxos de cadeia curta, para obter o melhor benefício p

---

### Chunk 26/30
**Article:** Remissão do Lúpus Através da Medicina Funcional Integrativa (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.555

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

### Chunk 27/30
**Article:** Fisiologia e Bioquímica do Sistema Imune I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.555

*   **Estratégias:** N-acetilcisteína (NAC), quercetina. Curcumina e resveratrol atuam como "coringas" no equilíbrio TH1/TH2.
*   **Modulação da Resposta TH17 (Fibromialgia, esclerose múltipla, artrite reumatoide)**
    *   **Citocinas:** IL-17, IL-22. O álcool aumenta a IL-17 e deve ser evitado.
    *   **Estratégias:** Boswellia (padronizada em AKBA), curcumina, berberina e ácido ursólico.
### 5. Estratégias Práticas e Plano de Tratamento Cíclico
*   **Os Sete Passos para a Reprogramação Intestinal**
    *   1. **Tomar Consciência:** Entender o diagnóstico e gatilhos.
    *   2. **Definir Estratégia:** Escolher abordagens (Low FODMAP, jejum).
    *   3. **Potencializar:** Usar tinturas, fitoterápicos, butirato.
    *   4. **Lapidar:** Reavaliar e ajustar.
    *   5. **Reconectar:** Integrar atividade física, meditação.
    *   6. **Fatores Psicossomáticos:** Trabalhar a conexão mente-intestino.
    *   7. **Comemorar:** Celebrar o progresso.

---

### Chunk 28/30
**Article:** Aula Jéssica Marques - Neurologia Funcional Integrativa 3 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.555

agnésio, vitamina B12, folato (B9), vitamina D e ferritina/ferro; corrigir deficiências conforme resultados.
- [ ] 3. Planejar introdução gradual de suplementação: iniciar uma formulação, aguardar 2–3 dias, adicionar a segunda e, após 2–3 dias, a terceira, monitorando efeitos colaterais.
- [ ] 4. Implementar ciclos de 60–90 dias entre plantas com mecanismos semelhantes, trocando para outra família após cada ciclo para sensibilizar diferentes receptores.
- [ ] 5. Selecionar extratos padronizados com maior biodisponibilidade (ex.: Curcumin C3, Cureit, CurcuVail) e evidência clínica para compor formulações.
- [ ] 6. Incorporar vias alternativas para idosos: avaliar uso de injetáveis, transdérmicos, fotobiomodulação, aromaterapia e tinturas conforme perfil e poder aquisitivo.
- [ ] 7. Estruturar rotina diária do paciente com atividades úteis e significativas (cozinhar, organizar casa, acompanhar netos/filhos), promovendo pertencimento e utilidade.
- [ ] 8.

---

### Chunk 29/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte VIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.554

cálcio iônico e ajustar vitamina D de forma personalizada, com reavaliação em 30–60 dias, especialmente após infecções/estresse.
- [ ] Investigar intolerância à histamina em pacientes com palpitações, arritmias, refluxo, gastrite, ansiedade/alterações do sono; correlacionar com dieta, microbiota e possíveis polimorfismos.
- [ ] Solicitar testes de micotoxinas urinárias em casos de sintomas persistentes sem explicação; revisar fontes alimentares (café, grãos, amendoim) e reforçar suporte hepático e intestinal.
- [ ] Mapear polimorfismos relevantes (ex.: FUT2) em pacientes com IBS ou baixa B12; personalizar dieta e estratégias para reforço da camada de muco.
- [ ] Implementar o protocolo começo–meio–fim: revisar dieta atual; avaliar digestão/absorção/transporte; checar suficiências de nutrientes-chave (vitamina D, C, A, zinco, selênio, ômega-3, B12) e corrigir.

---

### Chunk 30/30
**Article:** Intolerâncias, Alergias e Hipersensibilidades Alimentares II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.553

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

