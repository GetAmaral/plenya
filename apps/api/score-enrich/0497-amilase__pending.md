# ScoreItem: Amilase

**ID:** `019bf31d-2ef0-750e-b7f1-03a388ec5e62`
**FullName:** Amilase (Exames - Laboratoriais)
**Unit:** U/L

**Preparation Metadata:**
- Quality Grade: **GOOD**
- Total Chunks: 30 de 16 artigos
- Avg Similarity: 0.581

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-750e-b7f1-03a388ec5e62`.**

```json
{
  "score_item_id": "019bf31d-2ef0-750e-b7f1-03a388ec5e62",
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

**ScoreItem:** Amilase (Exames - Laboratoriais)
**Unidade:** U/L

**30 chunks de 16 artigos (avg similarity: 0.581)**

### Chunk 1/30
**Article:** Disbiose I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.648

] 4. Solicitar elastase pancreática fecal para investigar insuficiência exócrina pancreática, especialmente em sintomas de má digestão de carboidratos/gorduras.
- [ ] 5. Investigar e tratar SIBO/SIFO/parasitoses (ex.: giardia) em pacientes com intolerâncias a dissacarídeos (lactose) e sintomas de má absorção; restaurar a integridade da mucosa.
- [ ] 6. Revisar a qualidade da dieta do paciente, enfatizando que energia e nutrientes vêm do alimento; alinhar a ingestão para atender cerca de 30 kcal/kg/dia quando apropriado ao estado basal.
- [ ] 7. Educar sobre a importância da saliva e da fase oral da digestão; evitar comer sob ansiedade/pressa, sentar para as refeições e focar no ato de comer.
- [ ] 8. Implementar estratégias para reduzir inflamação crônica de baixo grau, incluindo melhora da microbiota intestinal e redução de “garbage aging” por meio de suporte digestivo e antioxidante.
- [ ] 9.

---

### Chunk 2/30
**Article:** Modulação Intestinal II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.637

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

### Chunk 3/30
**Article:** Disbiose I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.624

atletas), estresse crônico.
   - Estresse crônico reduz aldosterona, comprometendo a reabsorção renal de sódio e água, levando à hipossalivação e problemas iniciais de digestão.
* Função pancreática
   - Diminuição da função exócrina pode ocorrer com a idade; esteatose pancreática é possível e afeta a secreção enzimática.
   - Diagnóstico funcional: elastase pancreática fecal reduzida indica insuficiência exócrina.
* Integridade da mucosa e contato enzimático
   - Enzimas de borda em escova precisam de mucosa íntegra e contato com substrato para converter dissacarídeos em monossacarídeos.
   - Hipercrescimento no intestino delgado:
     - SIBO (Small Intestinal Bacterial Overgrowth) e SIFO (fungos), além de parasitas (ex.: giardia), podem formar “tapete” que impede o contato mucosa-substrato, causando intolerâncias (como lactose) por falha de digestão local.

---

### Chunk 4/30
**Article:** Dislipidemias II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.606

dade, inatividade física, pressão arterial e dislipidemia.
- A inflamação é indicada por marcadores como a Proteína C-Reativa (PCR), onde um valor de 5 já é considerado elevado, e a Interleucina 6 (exemplo de paciente com 8.45).
- A resposta anti-inflamatória é medida pela Interleucina 10, com um valor de corte de 3,5 (abaixo disso é um risco) e um exemplo de paciente com 6.44.
- A saúde do endotélio é avaliada pelo óxido nítrico, que deve estar em 8.8; um paciente com 4.8 já apresenta disfunção sistêmica.
- A LDL oxidada é um marcador crítico, com um limite saudável de 133, enquanto pacientes de alto risco podem apresentar valores extremos, como 1000.
**Achados Adicionais**
- Existem cinco parâmetros clássicos para definir a síndrome metabólica.
- O alvo de LDL para pacientes de alto risco, segundo as diretrizes atuais, é de 50.

---

### Chunk 5/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XIX (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.599

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

### Chunk 6/30
**Article:** Cardiologia VIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.599

isco:** Ferramentas como as tabelas de Framingham e MESA, embora matemáticas, são imprecisas por não considerarem uma vasta gama de variáveis metabólicas (ex: sono, hormonas, função mitocondrial, insulina, hemoglobina glicada).
- **Gama GT (Gama-glutamil transferase):** Além de ser um marcador de saúde para rins, pâncreas, fígado e estômago, a Gama GT (GGT) atua como um marcador de significância clínica para desfechos cardiovasculares. Níveis elevados podem indicar toxicidades crônicas (metais pesados, poluentes, defensivos agrícolas) e estão associados a maior risco cardiovascular e mortalidade geral. O objetivo terapêutico é manter o valor no quartil inferior da referência. O uso de simbióticos pode ajudar a melhorar a função hepática.
- **Leucócitos:** São uma "marca individual" e sua análise deve ser comparativa com o histórico do próprio paciente.

---

### Chunk 7/30
**Article:** Microbioma Intestinal III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.598

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

### Chunk 8/30
**Article:** Fisiologia e Bioquímica do Sistema Imune II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.594

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

### Chunk 9/30
**Article:** Microbioma Intestinal III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.588

(elevada) — glicoproteína que inibe elastase neutrofílica; marcador de atividade inflamatória crônica intestinal. Valor elevado sugere inflamação intestinal.
  - Referências educacionais: pH fecal, estercobilina, bilirrubina presentes no relatório (sem valores descritos).
- Marcadores adicionais:
  - Calprotectina fecal: 1.428 (ideal < 50) — muito elevada; correlaciona com atividade de doença inflamatória intestinal (DII).
  - Lactoferrina fecal: 9.330 — muito elevada; associada a neutrófilos fecais; diferencial inclui DII (Crohn/colite ulcerosa) e infecção entérica bacteriana (Shigella, Salmonella, Campylobacter, C. difficile, E. coli enteroinvasiva).
  - IgA secretória fecal: aumentada (sem valor numérico) — resposta imunológica mucosal elevada.
  - Elastase pancreática fecal: 85 — baixa; sugere insuficiência pancreática exócrina leve/moderada, possivelmente secundária a hipocloridria e disfunção digestiva global.

---

### Chunk 10/30
**Article:** Trato Gastrointestinal IV – Pâncreas e Vesícula Biliar (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.587

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

### Chunk 11/30
**Article:** Fisiologia e Bioquímica do Sistema Imune II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.586

olicitar e registrar parâmetros essenciais (PCR-us, VHS); calcular índices estimáveis (HGI, TAIG); considerar complementares (TNF-α, IL-6, GPx, MDA, antioxidantes totais).
- [ ] 11. Monitorar trimestralmente HbA1c, glicemia de jejum e HGI; estabelecer metas individuais (ex.: reduzir de 6,1 para ~5,5; longo prazo ~5,3).
- [ ] 12. Calcular TG/HDL e integrar com lipidograma/SREBP1c/2; evitar combinação de gordura saturada com açúcar e excesso de saturadas em múltiplas refeições.
- [ ] 13. Avaliar ferro, ferritina, transferrina e saturação (20–50%; evitar <20%); interpretar com hepcidina/SREBP1c e quadro inflamatório.
- [ ] 14. Medir TNF-α (<8,1) e IL-6 (<3,4) para acompanhar atividade inflamatória; relacionar com obesidade inflamada.
- [ ] 15. Calcular HOMA-β (167–175) e HOMA-IR (<2,15); buscar glicemia 60–90 e insulina ~6–7.
- [ ] 16. Monitorar homocisteína (<7,9) e PCR; usar PCR-us; documentar crises (PCR >1.000) e conduzir manejo apropriado.
- [ ] 17.

---

### Chunk 12/30
**Article:** Microbioma Intestinal III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.586

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

### Chunk 13/30
**Article:** Carboidratos IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.585

a faixa de referência, já indicam um problema.
*   **Hipoglicemia de Rebote (ou Reativa)**
    - Em pessoas com resistência à insulina, o pâncreas libera uma quantidade desproporcional de insulina, que, após baixar a glicose, continua alta e causa uma queda excessiva (hipoglicemia).
    - Essa hipoglicemia gera um desejo desesperado por comida, criando um ciclo vicioso de picos de glicose e insulina.
### 3. Análise de Casos Clínicos e Risco Cardiovascular
*   **Caso 1: Homem, 42 anos**
    - Paciente com 101 kg, IMC de 32. Glicemia de jejum de 89, mas insulina basal de 13.
    - A curva insulinêmica mostrou picos absurdos de insulina (ex: 81 em 60 minutos), confirmando a resistência à insulina severa.
*   **Caso 2: Mulher, 71 anos**
    - Paciente com 87 kg, múltiplas queixas (dores, depressão, hipertensão). Glicemia de jejum de 90 e insulina de 10.

---

### Chunk 14/30
**Article:** Disbiose II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.580

da pressão arterial.
- [ ] 2. Ao avaliar um paciente, investigar o nível de estresse, histórico de uso de medicamentos (antibióticos, prazois, anticoncepcionais), tipo de parto, aleitamento e hábitos alimentares.
- [ ] 3. Considerar o exame coprológico funcional como ferramenta principal para diagnosticar disbiose e problemas de digestibilidade.
- [ ] 4. Priorizar a melhoria da eficiência digestiva (com enzimas, mastigação) e o controle do estresse como primeiros passos no tratamento da disbiose, antes de prescrever probióticos.
- [ ] 5. Monitorar os níveis de vitaminas lipossolúveis (A, D, E, K) e B12 em pacientes com condições que afetam a absorção, como cirurgia bariátrica, doença celíaca ou disbiose.
- [ ] 6. Considerar a suplementação de zinco para otimizar a absorção de ácido fólico, dado que sua hidrólise é dependente deste mineral.
- [ ] 7.

---

### Chunk 15/30
**Article:** Trato Gastrointestinal IV – Pâncreas e Vesícula Biliar (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.578

de LPS, dor e inflamação. Conclui com postura crítica e aberta sobre práticas como “limpeza do fígado/vesícula”, recomendando avaliação individual, evidência clínica e integridade profissional.
## 🔖 Pontos de Conhecimento
### 1. Fisiologia do suco pancreático e composição
- Produção diária de 2–3 L, regulada dinamicamente por estímulos digestivos.
- Enzimas: proenzimas proteolíticas (tripsinogênio, quimiotripsinogênio, procarboxipolipeptidase), amilase/dissacaridases/trisacaridases para carboidratos, lipases/esterases/fosfolipases para gorduras.
- Bicarbonato alcaliniza o quimo ácido para atividade ótima enzimática; regulado por acetilcolina e CCK, dependente de acidez gástrica preservada.
### 2. Regulação exócrina do pâncreas via duodeno
- Função exócrina: secreções “para fora” no trato GI.
- Gorduras e ácido no duodeno estimulam células I (CCK) e S (secretina).

---

### Chunk 16/30
**Article:** Disbiose I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.577

utilidade social/familiar.
* Microbiota como “órgão esquecido”
   - Reconhecimento crescente na medicina tradicional de que interações entre sistema digestivo e microbiota impactam saúde sistêmica e envelhecimento.
## ❓ Perguntas
- [Inserir Pergunta/Confusão]
## 📚 Tarefas
- [ ] 1. Orientar pacientes a mastigar até formar pasta antes de engolir, especialmente alimentos fibrosos ou carnes, para reduzir partículas >2 mm e melhorar a digestão.
- [ ] 2. Avaliar sinais de hipossalivação: investigar estresse crônico, hidratação, histórico de sialite/Sjögren/cálculos salivares; orientar hidratação e manejo do estresse.
- [ ] 3. Solicitar hemograma com diferencial para monitorar monócitos; considerar monócitos >8% como pista de inflamação crônica/inflammaging.
- [ ] 4. Solicitar elastase pancreática fecal para investigar insuficiência exócrina pancreática, especialmente em sintomas de má digestão de carboidratos/gorduras.
- [ ] 5.

---

### Chunk 17/30
**Article:** Disbiose II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.573

:** É o primeiro passo. Inclui mastigação, *mindful eating* e uso de enzimas digestivas (suplementos como pancreatina ou alimentos como mamão e abacaxi).
    2.  **Ajustes na Dieta:** Individualizar a dieta, fracionar refeições se necessário.
    3.  **Controle do Estresse:** Encaminhar para psicoterapia ou terapias complementares.
    4.  **Suplementação Adicional:** Avaliar a necessidade de aminoácidos, vitaminas e minerais.
    5.  **Atividade Física:** Melhora a motilidade intestinal.
    *   A suplementação com probióticos não deve ser a primeira linha de tratamento, mas sim uma etapa posterior, se necessária.
## ❓ Perguntas
- [Inserir Pergunta/Dúvida]
## 📚 Tarefas
- [ ] 1. Aumentar a ingestão hídrica para auxiliar na fluidificação das fezes e na manutenção da pressão arterial.
- [ ] 2.

---

### Chunk 18/30
**Article:** Bases Metabólicas das Doenças Crônicas e Gerenciamento - Glicação 2 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.573

e jejum ideal ≤6 µU/mL; aceitável até 10 µU/mL. Paciente refere 4–5 µU/mL.
  - Hemoglobina glicada: melhor indicador de glicação do que glicemia de jejum; pode estar elevada mesmo com glicemia normal (compatível com picos glicêmicos e alta carga glicêmica).
  - Glicemia de jejum isolada é exame “pobre” para avaliar glicação.
  - Recomendada curva insulinêmica-glicêmica para avaliar resistência insulínica e resposta a carboidratos.
- Genética:
  - Polimorfismos em FTO (5 variantes), MC4R (múltiplas), PPAR-γ, TCF7L2, SOD, CETP, BDNF, CCK, LEP e LEP-R associados a resistência insulínica, obesidade, menor saciedade, risco de diabetes e transporte lipídico desfavorável.
- Sinais e correlações fisiopatológicas:
  - Interação AGE-RAGE ativa NF-κB, aumenta citocinas e ROS, perpetuando inflamação crônica.
  - Excesso de ultraprocessados e preparos em alta temperatura eleva produtos de glicação avançada (AGEs).

---

### Chunk 19/30
**Article:** Pedro Neuro - Neurologia Funcional Integrativa 1 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.572

12 (avaliar ácido metilmalônico).
  - Vitamina B1 (tiamina; considerar pirofosfato em hemácias).
  - Vitamina E 12–20 μg/mL (preferir fontes alimentares).
  - Resistência insulínica: reduzir açúcar para ≤15 g/dia; EDI compete com degradação de amiloide.
  - AGEs: reduzir frituras, assados e grelhados em alta temperatura.
  - Inflamação: PCR <0,9 mg/L (ideal <0,7); ferritina, ácido úrico, VSG, RDW; causas incluem intestino, boca e estresse/ruminação.
  - Vitamina D 50–80 ng/mL.
  - Tireoide: otimizar TSH/T4/T3.
  - Hormônios sexuais: estradiol/progesterona/testosterona; mulheres mais afetadas (menopausa vs andropausa).
  - Eixo adrenal: cortisol (alto/baixo), pregnenolona meta 50–100, DHEA com metas por sexo.
  - Minerais: zinco/cobre na proporção adequada; magnésio (idealmente RBC), suplementar mesmo com sérico normal; selênio; glutationa.
  - Metais tóxicos: mercúrio, chumbo, cádmio, arsênico; dosagem anual.

---

### Chunk 20/30
**Article:** Disbiose II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.568

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

### Chunk 21/30
**Article:** Microbioma Intestinal III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.567

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

### Chunk 22/30
**Article:** Trato Gastrointestinal I- boca e esôfago (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.566

s e impactos metabólicos.
## Conteúdo Remanescente
1. O papel do fígado e das enzimas pancreáticas na digestão.
2. Causas da hipocloridria e exames relacionados.
3. Tratamentos funcionais para a hipocloridria, incluindo estratégias terapêuticas.
4. Detalhes sobre o intestino, conforme mencionado na discussão sobre o nervo vago.
5. Terapias digestivas detalhadas que serão abordadas por outros especialistas.
## Conteúdo Abordado
### 1. Introdução à Importância do Trato Digestório
- A avaliação do ritmo circadiano do sistema digestório é o ponto de partida para todas as consultas, pois melhorar as questões digestivas é fundamental para a saúde geral.
- Breve menção aos componentes do trato digestório: boca (amilase, muco), esôfago (muco), fígado, estômago (suco gástrico) e pâncreas (enzimas).
### 2.

---

### Chunk 23/30
**Article:** Hipertensão Arterial Sistêmica II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.565

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

### Chunk 24/30
**Article:** Trato Gastrointestinal I- boca e esôfago (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.561

ausando problemas digestivos.
- Mastigar um alimento trinta vezes até que se torne líquido é considerado o ideal, mas é uma prática rara.
**O ambiente estomacal opera dentro de faixas de pH estritamente reguladas, sendo que a acidez é essencial para a digestão de proteínas e a esterilização dos alimentos, um equilíbrio que é frequentemente perturbado por intervenções médicas inadequadas.**
- Em jejum, o estômago mantém um pH altamente ácido entre 1 e 3, essencial para suas funções.
- Após uma refeição, o pH sobe para cerca de 4.5 antes de retornar a níveis ácidos (1.9 a 2.5) em aproximadamente 3 horas.
- A enzima pepsina, crucial para a digestão de proteínas, é ativada em um pH entre 1.8 e 3.5, mas torna-se inativa acima de um pH de 4 ou 5.
- O uso de "prazois" pode elevar o pH do estômago para 4, inativando a pepsina e prejudicando a digestão, o que questiona a lógica de alcalinizar um ambiente que precisa ser ácido.

---

### Chunk 25/30
**Article:** Fisiologia e Bioquímica do Sistema Imune II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.558

tegrativas ~5,3–5,2; diagnóstico ≥6,5; risco alto ≥5,6. Evoluções podem levar 2–3 anos.
- Frutosamina: ~20 dias; complementar.
- HGI: diferença entre HbA1c observada e predita da glicemia; estratos de risco orientam acompanhamento trimestral.
- MDA: <4,8; GPx: >400 (ideal 800–1000); antioxidantes totais: 560–1120.
- TAIG: TG/(glicose/2); meta <8; TG/HDL: mulheres <1,4; homens <1,2.
- Lipidograma/SREBP1c/2: excesso de saturadas + açúcar eleva SREBP1c, VLDL e LDL ox; aumenta hepcidina e altera ferro.
- Ferro/ferritina/transferrina: saturação 20–50% (evitar <20%); hiperferritinemia inflamatória (“Serum Ferritin Lacking Iron”).
- TNF-α: meta <8,1; IL-6: meta <3,4; relação direta em obesidade inflamada.
- HOMA-β: 167–175; HOMA-IR: <2,15; glicemia alvo 60–90; insulina ideal ~6–7 (até 8 em inflamação).
- Homocisteína: <7,9; correlaciona com PCR; polimorfismos FUT/MTHFR podem elevá-la.

---

### Chunk 26/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** results | **Similarity:** 0.555

with and 
without hyperbilirubinemia. Dig Dis Sci 2008;53:799-
802.163. Lazo M, Selvin E, Clark JM. Brief communication: 
clinical implications of short-term variability in liver 
function test results. Ann Intern Med 2008;148:348-52.164. Schmidt E, Schmidt FW, Chemnitz G, Kubale R, 
Lobers J. The Szasz-ratio (CK/GOT) as example for the 
diagnostic significance of enzyme ratios in serum. Klin 
Wochenschr 1980;58:709-18.165. Dufour DR. Is it necessary to order aspartate 
aminotransferase with alanine aminotransferase 
in clinical practice? Author’s Reply. Clin Chem 
2001;47:1134-5.

---

### Chunk 27/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa III (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.554

minase IgA. Se a IgA for baixa, dosar IgG sérica e antitransglutaminase IgG.
-   **Teste Parasitológico de Fezes:** Recomendado no contexto do Brasil.
-   **Avaliação da Permeabilidade Intestinal:** O aumento da permeabilidade (leaky gut) pode ser avaliado pela zonulina (fecal ou sérica). Menciona-se que o estresse (injeção de CRH) pode induzir um aumento nos marcadores de leaky gut.
-   **Avaliação da Microbiota/Metabolômica:** A avaliação isolada da microbiota é considerada de pouco valor. A avaliação da metabolômica (ex: ácidos orgânicos urinários) é mais útil para avaliar a função da microbiota e detectar metabólitos bacterianos e fúngicos. O aumento do D-lactato no sangue pode estar associado ao uso de probióticos e causar "brain fogginess".
-   **Teste Respiratório para SIBO/IMO:** Considerado o método prático padrão, utilizando lactulose ou glicose.
    -   **Critério para SIBO (Hidrogênio):** Elevação acima de 20 ppm nos primeiros 90 minutos.

---

### Chunk 28/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.553

intestinal.
-   **SIFO:** Sensação de "blurring" (efeito baiacu - sentir-se distendido mesmo sem distensão objetiva) e dor abdominal.
## Objetivo:
A palestra descreve os exames e achados objetivos para a avaliação de pacientes com suspeita de Síndrome do Intestino Irritável (SII), embora não contenha achados de exame físico de um paciente específico.
-   **Exames Laboratoriais Gerais:** Hemograma e marcadores de atividade inflamatória.
-   **Calprotectina Fecal:** Usada para descartar doença inflamatória intestinal. Um valor abaixo de 100 µg/g tem uma positividade de 98% para o diagnóstico de SII. Entre 100-250 µg/g é uma zona cinzenta. Acima de 250 µg/g requer colonoscopia.
-   **Avaliação para Doença Celíaca:** Dosagem de IgA sérica total e anticorpo antitransglutaminase IgA. Se a IgA for baixa, dosar IgG sérica e antitransglutaminase IgG.
-   **Teste Parasitológico de Fezes:** Recomendado no contexto do Brasil.

---

### Chunk 29/30
**Article:** Modulação Intestinal II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.553

do pâncreas antes de prescrever enzimas digestivas.
- [ ] 4. Monitorar marcadores de inflamação de baixo grau, como resistência à insulina (HOMA-IR), homocisteína e proteína C-reativa.
- [ ] 5. Monitorar os níveis de vitamina B12 ao longo da vida, especialmente em pacientes que usam inibidores de bomba de prótons ou bariátricos.
- [ ] 6. Em pacientes com resistência à insulina, avaliar o TMAO sérico para aferir o risco cardiovascular.
- [ ] 7. Para pacientes que utilizam inibidores da bomba de prótons, planejar um desmame cuidadoso para evitar o efeito rebote de hiperacidez.
- [ ] 8. Aplicar o conhecimento sobre os mecanismos de ação (ex: beta-glucana, butirato) para personalizar as intervenções nutricionais de acordo com as necessidades do paciente (ex: horário de administração para controle de saciedade).

---

### Chunk 30/30
**Article:** Trato Gastrointestinal IV – Pâncreas e Vesícula Biliar (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.552

 tica e ejeção biliar.
- [ ] 7. Aumentar ingestão de fibras prebióticas e alimentos coloridos; incluir chás ricos em polifenóis e um shot matinal, monitorando sintomas e bem-estar.
- [ ] 8. Avaliar resistência insulínica e saciedade; priorizar aumento fisiológico de GLP-1 (fibra, microbiota, otimização biliar) antes de análogos farmacológicos.
- [ ] 9. Monitorar sintomas e marcadores inflamatórios intestinais; promover microbiota saudável para formação de ácidos biliares secundários e ativação de TGR5.
- [ ] 10. Para disbiose: considerar probióticos (L. acidophilus, B. lactis, E. faecium); acompanhar lipídios e inflamação.
- [ ] 11. Em H. pylori confirmada: priorizar tratamento medicamentoso, usando probióticos como adjuvantes.
- [ ] 12. Considerar Astragalus membranaceus (100–200 mg/dia, até 500 mg) para modulação de LPS em disbiose/inflamação, com acompanhamento.
- [ ] 13.

---

