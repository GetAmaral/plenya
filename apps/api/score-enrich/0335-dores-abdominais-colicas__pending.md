# ScoreItem: Dores abdominais / cólicas

**ID:** `019bf31d-2ef0-765b-b291-f80f9dce057b`
**FullName:** Dores abdominais / cólicas (Histórico de doenças - Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual): - Segmento abdominal)

**Preparation Metadata:**
- Quality Grade: **GOOD**
- Total Chunks: 30 de 17 artigos
- Avg Similarity: 0.573

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-765b-b291-f80f9dce057b`.**

```json
{
  "score_item_id": "019bf31d-2ef0-765b-b291-f80f9dce057b",
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

**ScoreItem:** Dores abdominais / cólicas (Histórico de doenças - Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual): - Segmento abdominal)

**30 chunks de 17 artigos (avg similarity: 0.573)**

### Chunk 1/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.660

intestinal.
-   **SIFO:** Sensação de "blurring" (efeito baiacu - sentir-se distendido mesmo sem distensão objetiva) e dor abdominal.
## Objetivo:
A palestra descreve os exames e achados objetivos para a avaliação de pacientes com suspeita de Síndrome do Intestino Irritável (SII), embora não contenha achados de exame físico de um paciente específico.
-   **Exames Laboratoriais Gerais:** Hemograma e marcadores de atividade inflamatória.
-   **Calprotectina Fecal:** Usada para descartar doença inflamatória intestinal. Um valor abaixo de 100 µg/g tem uma positividade de 98% para o diagnóstico de SII. Entre 100-250 µg/g é uma zona cinzenta. Acima de 250 µg/g requer colonoscopia.
-   **Avaliação para Doença Celíaca:** Dosagem de IgA sérica total e anticorpo antitransglutaminase IgA. Se a IgA for baixa, dosar IgG sérica e antitransglutaminase IgG.
-   **Teste Parasitológico de Fezes:** Recomendado no contexto do Brasil.

---

### Chunk 2/30
**Article:** Microbioma Intestinal III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.639

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
**Article:** Microbioma Intestinal III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.609

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

### Chunk 4/30
**Article:** Disbiose I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.601

] 4. Solicitar elastase pancreática fecal para investigar insuficiência exócrina pancreática, especialmente em sintomas de má digestão de carboidratos/gorduras.
- [ ] 5. Investigar e tratar SIBO/SIFO/parasitoses (ex.: giardia) em pacientes com intolerâncias a dissacarídeos (lactose) e sintomas de má absorção; restaurar a integridade da mucosa.
- [ ] 6. Revisar a qualidade da dieta do paciente, enfatizando que energia e nutrientes vêm do alimento; alinhar a ingestão para atender cerca de 30 kcal/kg/dia quando apropriado ao estado basal.
- [ ] 7. Educar sobre a importância da saliva e da fase oral da digestão; evitar comer sob ansiedade/pressa, sentar para as refeições e focar no ato de comer.
- [ ] 8. Implementar estratégias para reduzir inflamação crônica de baixo grau, incluindo melhora da microbiota intestinal e redução de “garbage aging” por meio de suporte digestivo e antioxidante.
- [ ] 9.

---

### Chunk 5/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.594

ração do eixo cérebro-intestino-microbiota. O diagnóstico é clínico, baseado nos critérios de Roma 4, que exigem dor abdominal recorrente associada a alterações no hábito intestinal. A fisiopatologia envolve alterações no SNC, desequilíbrios da microbiota, fatores genéticos/epigenéticos e o papel de neurotransmissores como a serotonina.
- Diagnóstico Suspeito: Nenhum no momento
## Plano:
- Prescrição: Inserir mais aqui
- Próximos Passos/Exames:
    - A apresentação enfatiza a importância de considerar diagnósticos diferenciais, como constipação funcional e diarreia funcional.
    - É crucial investigar sinais de alarme, especialmente em pacientes com mais de 60 anos, para descartar doenças orgânicas como neoplasia de cólon.
    - Menciona abordagens terapêuticas gerais, como o uso de medicamentos que atuam em receptores de serotonina (5-HT) para modular a motilidade e a dor.

---

### Chunk 6/30
**Article:** Trato Gastrointestinal II- estômago – hipercloridria e hipocloridria (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.590

echos negativos discutidos nesta sessão.
   - Aprofundamento em estratégias alimentares com participação de Denise e Cristiano.
## ❓ Perguntas
- [Inserir Pergunta/Confusão]
## 📚 Tarefas
- [ ] 1. Revisar e ajustar plano alimentar funcional: reduzir farinha de trigo, café, lácteos, chocolate e ultraprocessados; implementar dieta compatível com digestão adequada.
- [ ] 2. Avaliar necessidade de endoscopia com pesquisa de H. pylori, interpretando resultados à luz dos sintomas e do padrão alimentar.
- [ ] 3. Solicitar exames laboratoriais: ferritina, saturação de transferrina; considerar anticorpos anti-células parietais se suspeita de gastrite atrófica autoimune.
- [ ] 4. Medir B12, folato, magnésio, cálcio, ferro e homocisteína em pacientes com sintomas de hipocloridria ou em uso crônico de IBP.
- [ ] 5. Reavaliar uso de IBP e antagonistas H2, ponderando riscos/benefícios e buscando estratégias não farmacológicas quando possível.
- [ ] 6.

---

### Chunk 7/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.586

o da permeabilidade intestinal, disbiose, inflamação) e comorbidades como distúrbios do sono, SIBO, IMO e SIFO. A abordagem diagnóstica e terapêutica é baseada na medicina funcional e integrativa, enfatizando a individualização do tratamento e a identificação das causas raiz.
-   **Diagnóstico Suspeito:** Nenhum no momento
## Plano:
-   **Prescrição:**
    -   **Neuromoduladores:** Amitriptilina (ação anti-inflamatória) ou Pregabalina (preferência do orador, iniciando com 50 mg/dia para sono, desconforto e distensão).
    -   **Antibióticos/Antifúngicos:** Rifaximina para SIBO; Fluconazol (curso de 2-3 semanas) para SIFO.
    -   **Estabilizadores de Mastócitos/Antialérgicos:** Cetotifeno, Ebastina, Levocetirizina, Montelucaste.
    -   **Suplementos e Nutracêuticos:**
        -   **Controle de Sintomas:** Cápsula de óleo de hortelã-pimenta (dor abdominal).

---

### Chunk 8/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.582

m Síndrome do Intestino Irritável (SII) e condições associadas como SIBO (Supercrescimento Bacteriano do Intestino Delgado) e SIFO (Supercrescimento Fúngico do Intestino Delgado).
-   **Sintomas Gerais:** Dor e distensão abdominal (sensação de "baiacu"), má qualidade do sono, sintomas depressivos, "brain fogginess" (confusão mental), esquecimento e dor abdominal associada ao período menstrual em mulheres.
-   **SII:** Distensão abdominal, dor e desconforto. Os sintomas podem variar, incluindo diarreia ou constipação.
-   **SIBO:** Dor, diarreia, distensão abdominal. Também pode se manifestar com deficiências nutricionais (ex: vitamina B12, ferro baixo) sem sintomas gastrointestinais clássicos.
-   **IMO (Intestinal Methanogenic Overgrowth):** Predomínio de constipação intestinal.
-   **SIFO:** Sensação de "blurring" (efeito baiacu - sentir-se distendido mesmo sem distensão objetiva) e dor abdominal.

---

### Chunk 9/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.582

s sintomas subjetivos e características clínicas gerais de pacientes com Síndrome do Intestino Irritável (SII). De acordo com os critérios de Roma 4, o paciente deve apresentar dor abdominal recorrente (pelo menos uma vez por semana nos últimos três meses, com início dos sintomas há pelo menos seis meses), associada a pelo menos dois dos seguintes critérios: relacionada à evacuação, associada a uma mudança na frequência das fezes ou associada a uma mudança na aparência das fezes. Pacientes com SII podem apresentar traços de personalidade e sintomas neuropsíquicos como ansiedade, hipervigilância aos sintomas, hipersensibilidade visceral, alteração na motilidade intestinal, alexitimia (dificuldade em identificar e descrever emoções) e hipercatastrofização (preocupações excessivas e crenças limitantes sobre comer fora, como medo de passar mal).
## Objetivo:
O conteúdo é uma apresentação médica teórica, não uma consulta.

---

### Chunk 10/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa I (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.570

 ões excessivas e crenças limitantes sobre comer fora, como medo de passar mal).
## Objetivo:
O conteúdo é uma apresentação médica teórica, não uma consulta. Descreve os achados objetivos e a fisiopatologia da SII. A classificação do padrão intestinal é feita usando a Escala de Bristol (tipos 1-2 para constipação, 3-4 normais, 5-7 para diarreia), subdividindo a SII em subtipos: com constipação (IBS-C), com diarreia (IBS-D), mista (IBS-M) ou indeterminada (IBS-U). Fisiologicamente, a SII envolve alterações na motilidade, sensibilidade visceral, função imune, inflamação de baixo grau, microbiota intestinal e processamento no sistema nervoso central (SNC).
- **Achados de Pesquisa:** Estudos de neurorressonância funcional mostram maior ativação em áreas cerebrais (rede sensorial motora, rede autonômica central) em pacientes com SII.

---

### Chunk 11/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte VIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.568

cálcio iônico e ajustar vitamina D de forma personalizada, com reavaliação em 30–60 dias, especialmente após infecções/estresse.
- [ ] Investigar intolerância à histamina em pacientes com palpitações, arritmias, refluxo, gastrite, ansiedade/alterações do sono; correlacionar com dieta, microbiota e possíveis polimorfismos.
- [ ] Solicitar testes de micotoxinas urinárias em casos de sintomas persistentes sem explicação; revisar fontes alimentares (café, grãos, amendoim) e reforçar suporte hepático e intestinal.
- [ ] Mapear polimorfismos relevantes (ex.: FUT2) em pacientes com IBS ou baixa B12; personalizar dieta e estratégias para reforço da camada de muco.
- [ ] Implementar o protocolo começo–meio–fim: revisar dieta atual; avaliar digestão/absorção/transporte; checar suficiências de nutrientes-chave (vitamina D, C, A, zinco, selênio, ômega-3, B12) e corrigir.

---

### Chunk 12/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa III (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.567

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

### Chunk 13/30
**Article:** Intolerâncias, Alergias e Hipersensibilidades Alimentares II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.567

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

### Chunk 14/30
**Article:** Bases Metabólicas das Doenças Crônicas e Gerenciamento - Inflamação 2 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.567

nto regular de PCR ultra-sensível em pacientes com sobrepeso, sinais de inflamação ou risco oncológico/cardiovascular.
- [ ] 2. Diferenciar anemia por deficiência de ferro de anemia da inflamação usando painel: BCM/HCM/CHr, % eritrócitos hipocrômicos, transferrina, receptor de transferrina plasmático, ferritina; considerar hepcidina em nível acadêmico.
- [ ] 3. Ajustar plano alimentar conforme perfil: iniciar low carb para sobrepeso/inflamação sem constipação; para mulheres constipadas, priorizar regulação intestinal com incremento vegetal cuidadoso e possível redução de carne vermelha.
- [ ] 4. Prescrever fibras não fermentativas para constipação com gases: goma acácia até 5 g/dia e polidextrose até 3 g/dia; avaliar resposta e adaptar.
- [ ] 5. Introduzir ômega 3 (EPA/DHA) com dose individualizada; garantir dieta antioxidante concomitante para evitar oxidação e otimizar incorporação; evitar depender de ALA (linhaça/chia) como única fonte.
- [ ] 6.

---

### Chunk 15/30
**Article:** Intolerância à Histamina (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.566

erância à histamina, realizar o diagnóstico diferencial para excluir condições como síndrome de ativação mastocitária e alergias alimentares.
- [ ] 3. Implementar uma dieta baixa em histamina com acompanhamento de um nutricionista como primeira linha de tratamento ("Food First").
- [ ] 4. Considerar a suplementação com a enzima DAO 20 minutos antes das refeições para controle dos sintomas.
- [ ] 5. Avaliar e tratar a saúde intestinal, investigando a presença de hiperpermeabilidade (leaky gut) e disbiose com bactérias estaminogênicas.
- [ ] 6. Avaliar a necessidade de reposição de cofatores da enzima DAO (cobre, vitamina C, vitamina B6).
- [ ] 7. Pausar o vídeo para observar a lista de medicamentos (antidepressivos, anti-hipertensivos, antibióticos) que podem diminuir a atividade da enzima DAO.
- [ ] 8. Utilizar bloqueadores de receptores H1 e H2 como terapia sintomática quando necessário.
- [ ] 9.

---

### Chunk 16/30
**Article:** Disbiose II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.566

da pressão arterial.
- [ ] 2. Ao avaliar um paciente, investigar o nível de estresse, histórico de uso de medicamentos (antibióticos, prazois, anticoncepcionais), tipo de parto, aleitamento e hábitos alimentares.
- [ ] 3. Considerar o exame coprológico funcional como ferramenta principal para diagnosticar disbiose e problemas de digestibilidade.
- [ ] 4. Priorizar a melhoria da eficiência digestiva (com enzimas, mastigação) e o controle do estresse como primeiros passos no tratamento da disbiose, antes de prescrever probióticos.
- [ ] 5. Monitorar os níveis de vitaminas lipossolúveis (A, D, E, K) e B12 em pacientes com condições que afetam a absorção, como cirurgia bariátrica, doença celíaca ou disbiose.
- [ ] 6. Considerar a suplementação de zinco para otimizar a absorção de ácido fólico, dado que sua hidrólise é dependente deste mineral.
- [ ] 7.

---

### Chunk 17/30
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

### Chunk 18/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.563

pecialmente em indivíduos com mais de 60 anos ou com sinais de alarme, para descartar doenças orgânicas como neoplasia de cólon.
- [ ] 3. Utilizar a Escala de Bristol para obter uma descrição detalhada do padrão intestinal do paciente, indo além de perguntas genéricas.
- [ ] 4. Acolher e validar as queixas de hipercatastrofização e hipervigilância dos pacientes com SII, reconhecendo que são manifestações de alterações neurológicas reais.
- [ ] 5. Considerar a indicação de abordagens complementares, como a osteopatia, para pacientes com sintomas digestivos que possam ter relação com a inervação tóraco-lombo-sacral.

---

## SOAP

Data e Hora: 2025-11-17 17:55:53
Paciente: 
Diagnóstico:
## Histórico de Diagnóstico:
1. Histórico Médico: O conteúdo é uma apresentação médica sobre a Síndrome do Intestino Irritável (SII) e o eixo cérebro-intestino-microbiota, não um registro de paciente.

---

### Chunk 19/30
**Article:** Trato Gastrointestinal V – Intestino Delgado I (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.563

 " antes de mergulhar no "como". A narrativa do caso clínico foi envolvente e clara. Para aprimorar ainda mais, você poderia, ao final da discussão do caso, resumir em 2-3 pontos-chave a diferença entre a abordagem que você critica e a abordagem que você defende, para solidificar o aprendizado antes de passar para o próximo tópico.
### 2. Fisiologia e Importância do Intestino Delgado
- O intestino é fundamental, pois é onde ocorrem as principais absorções e sinalizações do corpo.
- Enzimas presentes: sacarase, maltase, isomaltase, lactase, peptidase, lipase. A intolerância à lactose ocorre pela deficiência da enzima lactase.
- O muco no intestino delgado é alcalino para proteger a mucosa e otimizar a função enzimática.
- A modulação da digestão depende do bolo fecal e de vias aferentes vagais (sistema parassimpático).
- A importância de estar relaxado e presente (mindfulness) durante as refeições para ativar o parassimpático e garantir uma boa digestão.

---

### Chunk 20/30
**Article:** Trato Gastrointestinal III – estômago – hipocloridria (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.562

de Precisão: FODMAPs e Histamina
- **Dieta FODMAP:** Indicada para pacientes com hipocloridria e excesso de gases, pois a fermentação intestinal pode alterar a produção de ácido gástrico. Polióis (xilitol, eritritol) podem ser gatilhos.
- **Sensibilidade à Histamina:** Um diagnóstico diferencial importante, com sintomas como coceiras, dor de cabeça e rinite. Alimentos ricos em histamina incluem atum, fermentados, lácteos e abacaxi.
### 4. Tratamentos e Suplementos para Hipocloridria
- **Cloridrato de Betaína (Betaína HCL):** Usado para reverter a hipocloridria. A dosagem varia de 300mg a 1500mg, tomada com a primeira garfada da refeição, preferencialmente em comprimidos.
- **Aloe Vera:**
    - **Benefícios:** O gel da *Aloe barbadensis Miller* possui mais de 75 compostos bioativos com ação anti-inflamatória, cicatrizante, antioxidante e imunomoduladora.

---

### Chunk 21/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.559

dos sintomas gerais a investigar em consulta, como dor mamária e TPM intensa, que podem indicar predominância estrogênica.

## Objetivo:
Não há achados de exame físico de um paciente específico. A abordagem objetiva descrita inclui:
- Avaliação da composição corporal por bioimpedância ou densitometria de corpo total, identificando condições como sarcopenia em mulheres com peso normal.
- Solicitação de exames laboratoriais para avaliar marcadores inflamatórios e verificar se a paciente está metabolicamente doente ou em risco.

## Diagnóstico Primário:
- Avaliação: Discussão educacional sobre estratificação de risco para câncer de mama. 90% dos casos relacionam-se ao estilo de vida e 10% a fatores genéticos conhecidos. Mutações de alta penetrância (ex.: BRCA1, BRCA2, TP53) aumentam significativamente o risco, porém são raras. Epigenética e estilo de vida são cruciais e modificáveis.

---

### Chunk 22/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XIX (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.557

g/kg) e prática de exercícios de resistência para preservar massa muscular.
- [ ] 3. Todos os profissionais: Em doenças crônicas sem causa orgânica clara ou com má resposta ao tratamento, investigar ativamente traumas de infância, estresse crônico e questões emocionais não resolvidas como possível "causa primeira".
- [ ] 4. Terapeutas e psicólogos: Adotar "terapia de precisão", utilizando múltiplas ferramentas e combinando diferentes abordagens terapêuticas para personalizar o tratamento e focar em resultados mensuráveis, em vez de seguir uma única linha teórica por longos períodos.
- [ ] 5. Estudo pessoal: Pesquisar o conceito de "causa primeira" de Aristóteles para aprofundar a lógica de buscar a origem dos problemas.
- [ ] 6. Estudo pessoal: Ler o livro de Bruce Lipton sobre a conexão entre mente e doenças físicas.

---

## SOAP

> Data e Hora: 2025-11-17 16:33:53
> Paciente: 
> Diagnóstico:

## Histórico do Diagnóstico:
1. Histórico Médico: 
2.

---

### Chunk 23/30
**Article:** Modulação Intestinal II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.555

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

### Chunk 24/30
**Article:** Microbioma Intestinal III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.555

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

### Chunk 25/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.555

tação médica sobre a Síndrome do Intestino Irritável (SII) e o eixo cérebro-intestino-microbiota, não um registro de paciente. A apresentação aborda a prevalência da SII no Brasil, que pode chegar a 4,7% da população (critérios de Roma 4) ou 8,3% (critérios de Roma 3). A SII é mais frequente em mulheres (relação de 2-3 para 1 homem), ocorrendo geralmente entre a terceira e a quarta década de vida. O aparecimento de sintomas após os 60 anos é um sinal de alerta para outras doenças, como neoplasia de cólon. Fatores socioeconômicos mais baixos, urbanização e falta de suporte social estão associados a uma maior incidência.
2. Histórico de Medicação: Inserir mais aqui
## Subjetivo:
O conteúdo é uma apresentação médica, não uma consulta. No entanto, descreve os sintomas subjetivos e características clínicas gerais de pacientes com Síndrome do Intestino Irritável (SII).

---

### Chunk 26/30
**Article:** SAM (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.553

car os gatilhos individuais do paciente, incluindo dieta, exposições ambientais (poluentes, micotoxinas) e desequilíbrios internos (disbiose).
- [ ] 5. Para profissionais de saúde: Se não houver conforto ou especialização para tratar a SAM, encaminhar o paciente a um profissional qualificado após levantar a suspeita diagnóstica.
- [ ] 6. Para profissionais de saúde e pacientes: Considerar a suplementação com vitaminas (C, D, E), minerais (magnésio), probióticos e flavonoides (quercetina, luteolina) como parte de um plano de tratamento integrativo.
- [ ] 7. Para o público: Se não assistiu, pausar e assistir à aula sobre "Intolerância à Histamina" e "Síndrome do Intestino Irritável" para uma compreensão mais profunda dos conceitos relacionados à SAM.

---

## SOAP

Data e Hora: 2025-11-17 17:56:34
Paciente: 
Diagnóstico:
## Histórico de Diagnóstico:
1.

---

### Chunk 27/30
**Article:** Microbioma Intestinal III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.551

to com o laboratório para obter suporte na interpretação dos resultados.
- [ ] 6. Desenvolver um plano alimentar personalizado, evitando abordagens genéricas, especialmente se houver sinais de fermentação excessiva.
- [ ] 7. Ao prescrever Biointestil, alertar o paciente sobre a possível reação de Herxheimer e considerar uma introdução gradual.
- [ ] 8. Em casos de insuficiência pancreática funcional (elastase baixa), investigar a função gástrica como causa primária.
- [ ] 9. Estudar para a próxima aula, que abordará a prescrição de fibras, probióticos (cepas específicas) e o conceito de paraprobióticos.

---

## SOAP

> Data e Hora: 2025-11-17 17:48:32
> Paciente: [Speaker 1]
Diagnóstico:
## Histórico do Diagnóstico:
1. Histórico Médico:
   - Criança de 6 anos com quadro crônico de inflamação intestinal compatível com disbiose.
   - Nascimento por cesariana; alimentação inicial com fórmulas infantis.

---

### Chunk 28/30
**Article:** Microbioma Intestinal III (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.550

  do paciente).
   - Sem história de fibrose cística; sem indicação atual de neoplasia, pólipos ou doença celíaca confirmada, apenas discutidas como diferenciais.
   - Encaminhado à pediatra da equipe; quadro referido como “resolvido” após intervenções multifatoriais.
   - Discussão ampla sobre microbioma intestinal, homeostase versus disbiose, integridade de mucosas e sistema imunológico, com potenciais impactos sistêmicos (ossos, cérebro, saúde mental, distúrbios cognitivos, autoimunidade, obesidade, transtornos metabólicos, asma, alergias).
2. Histórico de Medicação:
   - Uso prévio de múltiplos medicamentos (antibióticos, corticoides; antidiarreicos em consulta com gastroenterologista).
   - Suplementos/intervenções discutidas: lactoferrina 500 mg, colostro, Biointestil (geraniol + gengibre), berberina.
   - Inserir mais aqui.
## Subjetivo:
- Distensão abdominal pós-prandial (estufamento), sugerindo fermentação inadequada.

---

### Chunk 29/30
**Article:** Trato Gastrointestinal IV – Pâncreas e Vesícula Biliar (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.550

efas
- [ ] 1. Avaliar hábitos de mastigação e educar sobre mastigação lenta/eficaz para melhorar digestibilidade.
- [ ] 2. Revisar uso de inibidores de ácido e considerar estratégias para restaurar acidez gástrica adequada quando indicado.
- [ ] 3. Investigar sinais de putrefação proteica (estufamento vespertino, gases, fezes fétidas) e correlacionar com dieta.
- [ ] 4. Avaliar ferro (hemograma, ferritina, saturação de transferrina) e suportar com vitamina C para otimizar CYPs e síntese biliar.
- [ ] 5. Considerar suplementação de taurina e glicina para suporte à destoxificação e potencial redução de gama-GT.
- [ ] 6. Implementar estratégias dietéticas que estimulem CCK/secretina (gorduras de boa qualidade e proteínas bem preparadas) para melhorar secreção pancreática e ejeção biliar.
- [ ] 7. Aumentar ingestão de fibras prebióticas e alimentos coloridos; incluir chás ricos em polifenóis e um shot matinal, monitorando sintomas e bem-estar.
- [ ] 8.

---

### Chunk 30/30
**Article:** Disbiose II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.549

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

