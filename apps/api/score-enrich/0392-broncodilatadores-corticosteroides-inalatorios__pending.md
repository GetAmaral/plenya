# ScoreItem: Broncodilatadores / Corticosteroides inalatórios

**ID:** `019bf31d-2ef0-7f5b-93c1-4590ed06b91f`
**FullName:** Broncodilatadores / Corticosteroides inalatórios (Histórico de doenças - Medicamentos - Uso atual de medicamentos)

**Preparation Metadata:**
- Quality Grade: **EXCELLENT**
- Total Chunks: 30 de 4 artigos
- Avg Similarity: 0.628

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-7f5b-93c1-4590ed06b91f`.**

```json
{
  "score_item_id": "019bf31d-2ef0-7f5b-93c1-4590ed06b91f",
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

**ScoreItem:** Broncodilatadores / Corticosteroides inalatórios (Histórico de doenças - Medicamentos - Uso atual de medicamentos)

**30 chunks de 4 artigos (avg similarity: 0.628)**

### Chunk 1/30
**Article:** pediatria funcional integrativa - parte IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.683

/ml, com cautela para evitar hipercalcemia.
*   **Ômega-3**
    *   Possui plausibilidade bioquímica, inibindo o NF-κB e a via do ácido araquidônico (reduzindo leucotrienos).
    *   Em gestantes, a suplementação reduziu a incidência de asma na criança.
    *   Ensaios clínicos falham em mostrar diferenças, levando a American Thoracic Society a recomendar contra seu uso.
    *   Polimorfismos no gene FADS podem explicar resultados conflitantes. Pacientes com ômega-índex > 8% necessitam de menos corticoide inalatório, inclusive obesos, contradizendo a recomendação da ATS.
### 3. Fitoterápicos no Tratamento da Asma
*   **Quercetina**
    *   Inibe a liberação de citocinas inflamatórias, a diferenciação de linfócitos TH2 e a liberação de histamina pelos mastócitos.
    *   Em humanos, a forma lipossomada (250-300mg) melhorou sintomas e peak flow. Doses de 500mg são consideradas seguras.
*   **Cúrcuma (Curcuminoides)**
    *   É segura em altas doses.

---

### Chunk 2/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.677

ra fenótipo de sibilância.
**Corticosteroides inalatórios: efetivos, mas com riscos hormonais, de crescimento e ósseos que exigem vigilância e individualização.**
- Supressão do eixo HPA: 10% sintomática e até 40% bioquímica; risco aumenta 6x em crianças e 4x em adultos com alta dose por 3–6 meses.
- Supressão com corticoide oral: cursos >2 semanas consecutivas ou >3 semanas em 6 meses elevam risco.
- Eixos de monitoramento: cortisol às 8h da manhã; se normal, reavaliar em 6 meses; no teste com ACTH, resposta deve subir 18 µg/dL; preocupação com valores de cortisol tão baixos quanto 3 mg/dL.
- Tratamento de supressão: hidrocortisona base por 6–12 meses; atrofia suprarrenal pode persistir até um ano após suspensão de inalatórios.
- ICS e crescimento: perda final de ~1 cm; diferença anual de ~0,2 cm; achados em revisão com quase 3.400 crianças, por 12–52 semanas; contínuo vs cromoglicato: ~1 cm a menos.

---

### Chunk 3/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.672

abilidade > 10% (adultos) e 13% (crianças).
    *   **Testes de Desafio:** Redução na função pulmonar com metacolina, exercício ou frio.
*   **Avaliação Sistêmica/Endócrina (Sinais de Supressão do Eixo HPA):**
    *   **Laboratorial (Triagem):** Eosinofilia periférica (>= 4%). Dosagem de Cortisol às 8h. Teste de estimulação com ACTH (necessário subir 18 mcg/dL; basal < 3 mcg/dL é preocupante).
    *   **Antropometria:** Aumento do IMC (0,07 kg/m²/ano de uso de CI), antecipação do reganho de adiposidade (rebound). Perda na velocidade de crescimento linear (impacto na altura final aprox. 1 cm).
    *   **Ósseo:** Sinais de osteopenia.
## [Diagnóstico Primário e Avaliação:]
*   **Diagnóstico Base:** Asma (Doença inflamatória crônica das vias aéreas).
    *   *Fenótipos:* Sibilante transitório, persistente não atópico, atópico/Asmático clássico (IgE), Asma Neutrofílica (associada à obesidade).

---

### Chunk 4/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.671

e (6-12 meses até normalização) e dose de estresse em infecções/cirurgias.
    *   **Suplementação/Prevenção:** Imunoestimulação, Vitamina D, Ômega 3, Carotenoides (foco em antioxidação e inflamação).
    *   **Exacerbações:** Corticoides orais (ex: Prednisolona).
*   **Próximos Passos e Exames:**
    *   **Monitoramento Respiratório:** Espirometria e aplicação do ACT (Teste de Controle da Asma) a cada consulta. Avaliação da técnica inalatória.
    *   **Monitoramento Endócrino/Crescimento:** Acompanhamento da estatura a cada 6 meses (crianças em uso de CI). Dosagem de cortisol às 8h (rastreio) e Teste de ACTH se sintomático com cortisol normal.
    *   **Investigação:** Monitorar interações com inibidores de CYP3A4 e comorbidades (refluxo, apneia).
*   **Plano de Tratamento de Acompanhamento:**
    *   **Controle Ambiental:** Redução de mofo, poeira, pelos, produtos químicos.

---

### Chunk 5/30
**Article:** pediatria funcional integrativa - parte IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.659

o de magnésio inalatório para crianças de 2-6 anos com exacerbação grave. Para maiores de 6 anos e adultos, pode ser usado 2g EV em caso de falha no tratamento inicial para evitar internação.
- **Uso Preventivo:** Um estudo com 330mg de magnésio por 6 meses mostrou melhora na qualidade de vida e controle da doença, mas sem alteração no VEF1 ou nos níveis séricos de magnésio.
### 4. Vitamina D e Asma
- **Mecanismo:** Níveis baixos (< 30 ng/ml) pioram o controle da asma. A Vitamina D melhora a ação do corticoide e modula a resposta imune (diminui citocinas inflamatórias e aumenta a anti-inflamatória IL-10).
- **Evidências:** Apesar da forte plausibilidade, meta-análises falham em demonstrar que a suplementação reduz exacerbações em adultos. Em crianças, há uma redução de 50%, possivelmente por diminuir viroses. As falhas nos estudos podem ser devidas a polimorfismos, vieses, doses inadequadas e níveis alvo insuficientes.

---

### Chunk 6/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.651

liza a magnitude da asma no Brasil e no mundo, discute critérios diagnósticos clínicos e funcionais (GINA), ferramentas de acompanhamento (ACT e controle por critérios GINA), e o impacto da inflamação crônica e remodelamento brônquico na progressão irreversível para DPOC. Detalha os steps terapêuticos do GINA por faixa etária, a adesão (especialmente em adolescentes), e os fenótipos de sibilância em <6 anos (transitório, persistente não atópico, atópico/asmático), com critérios de risco de asma e destaque de que sibilância por viroses em <3 anos demanda imunostimulação e prevenção, não aumento de ICS. Integra fatores farmacológicos, genéticos (SNP RS591118 em PDGFD), dispositivos e características das drogas que aumentam absorção sistêmica de ICS, fornece protocolos de rastreio (cortisol matinal e teste com ACTH) e manejo da supressão adrenal com hidrocortisona e doses de estresse.

---

### Chunk 7/30
**Article:** pediatria funcional integrativa - parte IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.648

atórios. Além disso, a discussão aprofunda-se na importância do microbioma pulmonar e do eixo intestino-pulmão, a hipótese da higiene e a janela de oportunidade na primeira infância para modular o sistema imune. A apresentação diferencia os fenótipos e endótipos da asma (TH2 e não-TH2), suas respostas distintas aos tratamentos e como a medicina funcional pode oferecer uma abordagem mais eficaz, especialmente para casos de asma grave, focando na remoção de gatilhos, dieta, controle de comorbidades e modulação imunológica.
## 🔖 Pontos de Conhecimento
### 1. Abordagem Funcional no Tratamento da Asma
*   **Vitamina K2 e Saúde Óssea**
    *   Os corticoides reduzem a massa óssea por meio do bloqueio da osteoprotegerina, o que leva a um aumento dos osteoclastos e, consequentemente, à perda de massa óssea.

---

### Chunk 8/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.646

homozigose; NNT estimado 27 para prevenir um caso se testado.
  - Pode aumentar risco de efeitos colaterais de corticoide, inclusive obesidade.
### 7. Obesidade associada ao corticoide inalatório e à asma
* ICS e IMC
  - Aumento do IMC de 0,07 kg/m² por ano; antecipação do reganho de IMC (2 meses/ano de uso); padrão central/visceral.
  - Não dependente da dose (efeito teto); menor em 0–3 anos, maior entre 3–6 anos; possível programação metabólica.
  - Limitações: confusão por gravidade da asma, sedentarismo, maior uso de medicação.
* Asma relacionada à obesidade (fenótipo não-TH2/neutrofílico)
  - Ciclo vicioso entre inflamação sistêmica e pulmonar; abordagem em três pilares: inflamação (vitamina D, ômegas), antioxidação (carotenoides) e metilação (epigenética).
### 8.

---

### Chunk 9/30
**Article:** pediatria funcional integrativa - parte IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.643

ção de 50%, possivelmente por diminuir viroses. As falhas nos estudos podem ser devidas a polimorfismos, vieses, doses inadequadas e níveis alvo insuficientes.
- **Recomendação Prática:** Usar doses diárias, por longos períodos, buscando níveis acima de 60 ng/ml, com cautela em crianças.
### 5. Ômega 3, PCR e Asma
- **Mecanismo:** O EPA inibe a via do ácido araquidônico (diminuindo leucotrienos) e o DHA diminui a infiltração eosinofílica.
- **Evidências:** Em gestantes, doses altas reduziram a incidência de asma nas crianças. Pacientes com ômega-3 index > 8% necessitam de menos corticoide. A falha de meta-análises pode ser explicada por polimorfismos (ex: FADES) que determinam a resposta à suplementação.
- **Contexto Clínico:** O PCR está aumentado em asmáticos não controlados, especialmente na asma neutrofílica (comum em obesos), e o ômega 3 pode ajudar a reduzir essa inflamação.
### 6.

---

### Chunk 10/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.639

rotocolos de rastreio (cortisol matinal e teste com ACTH) e manejo da supressão adrenal com hidrocortisona e doses de estresse. Discute a interface asma-obesidade (fenótipo neutrofílico, inflamação sistêmica, antioxidantes, vitamina D, ômegas, metilação), evidencia impactos de ICS em IMC e crescimento (CAMP Study e revisão Cochrane), e reforça o uso da menor dose eficaz, regimes intermitentes/sob demanda quando apropriado, e manejo integrado de ambiente e comorbidades (rinite, alergia alimentar, refluxo, obesidade, anemia).
## 🔖 Knowledge Points
### 1. Panorama epidemiológico e relevância clínica da asma
* Epidemiologia global e nacional
  - >300 milhões de asmáticos no mundo; no Brasil, >20 milhões; 7º em prevalência.
  - DataSUS: 2021 com 1,3 mi atendimentos na APS; 2022 com 83 mil internações e 574 óbitos.

---

### Chunk 11/30
**Article:** pediatria funcional integrativa - parte IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.638

ças e 30% das mulheres adultas.
- Antagonistas de leucotrienos, usados no tratamento da asma, podem causar sintomas psiquiátricos em até 20% das crianças.
- Pacientes asmáticos em CTI apresentam uma alta taxa de colonização fúngica na pele (54%).

---

## Teaching Note

Data e Hora: 2025-12-09 04:55:32
Local: [Inserir Local]: [Inserir Local]
Aula: [Inserir Nome da Aula]: [Inserir Nome da Aula]
## Visão Geral
A aula abordou a abordagem funcional e integrativa no tratamento da asma, focando em suplementos, fitoterápicos e na modulação do sistema imunológico. Foram discutidos os papéis e evidências da Vitamina K2, Ferro, Magnésio, Vitamina D, Ômega 3, Quercetina, Cúrcuma e Boswellia Serrata, contrastando a plausibilidade bioquímica com os resultados de ensaios clínicos.

---

### Chunk 12/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.634

--

## SOAP

> [Data e Hora: ] 2025-12-09 04:52:19
> [Paciente:]
> [Diagnóstico:]
## [Histórico do Diagnóstico:]
1. **Histórico Médico:**
    *   **Respiratório/Atópico:** Asma (condição base); Rinite alérgica (comorbidade frequente), dermatite atópica, alergia alimentar. Histórico de sibilância na infância (transitória ou persistente).
    *   **Sistêmico/Comorbidades:** Refluxo gastroesofágico, anemia (deficiência de ferro), obesidade (associada à asma neutrofílica e ao uso de corticoides - padrão visceral/maçã).
    *   **Complicações/Predisposições:** Predisposição genética para supressão do eixo HPA (Polimorfismo SNP-RS591118 do gene PDGFD); Histórico de inflamação sistêmica e local (ciclo vicioso asma-obesidade); Osteopenia (associada a ciclos frequentes de corticoide oral).
2. **Histórico de Medicamentos:**
    *   Corticoides inalatórios (Budesonida, Fluticasona, Beclometasona - incluindo apresentação em nanopartículas).

---

### Chunk 13/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.630

ótipos:* Sibilante transitório, persistente não atópico, atópico/Asmático clássico (IgE), Asma Neutrofílica (associada à obesidade).
*   **Complicações/Efeitos Adversos:** Efeitos do uso prolongado de corticosteroides:
    *   Supressão do eixo hipotálamo-pituitária-adrenal (Disfunção Adrenal Secundária).
    *   Obesidade iatrogênica/metabólica.
    *   Retardo de crescimento estatural.
    *   Osteopenia.
## [Plano:]
*   **Prescrição e Medicamentos:**
    *   **Controle da Asma:** Corticoides inalatórios (ajustar para menor dose possível; cuidado com potência da fluticasona vs beclometasona/budesonida).
    *   **Broncodilatadores:** SABA (resgate), LABA (associado ao corticoide), LAMA (se indicado).
    *   **Tratamento da Supressão Adrenal:** Hidrocortisona base (6-12 meses até normalização) e dose de estresse em infecções/cirurgias.

---

### Chunk 14/30
**Article:** pediatria funcional integrativa - parte IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.627

qualidade de vida, controle da doença e redução da hiper-reatividade brônquica.
### 2. Vitamina D e Ômega-3 na Asma
*   **Vitamina D**
    *   Baixos níveis de vitamina D (< 30 ng/ml) estão associados a um pior controle da asma.
    *   A vitamina D melhora a eficácia dos corticoides, diminui citocinas inflamatórias (IL-5, IL-13, IL-17), aumenta a IL-10 (anti-inflamatória) e atua localmente no pulmão.
    *   O GINA sugere manter níveis > 30 ng/ml na gestação para prevenção primária.
    *   Apesar da plausibilidade, meta-análises falham em demonstrar redução de exacerbações em adultos, possivelmente devido a polimorfismos do receptor (VDR), doses inadequadas e outros vieses.
    *   A recomendação é manter a suplementação diária, buscando níveis acima de 60 ng/ml, com cautela para evitar hipercalcemia.
*   **Ômega-3**
    *   Possui plausibilidade bioquímica, inibindo o NF-κB e a via do ácido araquidônico (reduzindo leucotrienos).

---

### Chunk 15/30
**Article:** pediatria funcional integrativa - parte IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.624

cional e Integrativa
*   **Princípio:** Usar a menor dose efetiva de medicação para controle da doença, focando na redução gradual ("step-down").
*   **Intervenções:**
    *   **Remoção de Gatilhos:** Além de alérgenos, inclui produtos químicos (amaciantes), perfumaria e metais pesados (arsênico).
    *   **Dieta e Nutrição:** Dieta anti-inflamatória, livre de alérgenos e contaminantes.
    *   **Atividade Física:** Recomendada, com uso preventivo de SABA se necessário para broncoespasmo induzido por exercício.
    *   **Técnicas Mente-Corpo:** Mindfulness e exercícios respiratórios.
    *   **Controle de Comorbidades:** Manejo de anemia, carências nutricionais, obesidade e efeitos colaterais dos corticoides.
## ❓ Perguntas
- [Inserir Pergunta/Confusão]
## 📚 Tarefas
- [ ] 1. Implementar uma dieta anti-inflamatória, livre de alérgenos, contaminantes e defensivos agrícolas.
- [ ] 2.

---

### Chunk 16/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.618

esposta a tratamento; testes de desafio.
* Acompanhamento do controle
  - ACT (5 itens; 5–25 pontos) nas versões pediátrica e adulta.
  - Critérios GINA (4 itens/4 semanas): 0 = controlada; 1–2 = parcialmente; 3–4 = não controlada.
### 3. Risco de remodelamento e progressão
* Inflamação subclínica persistente + broncoespasmo levam a destruição epitelial e remodelamento brônquico, com irreversibilidade e evolução para DPOC.
### 4. Terapêutica tradicional por faixa etária (steps GINA) e adesão
* Princípios
  - ICS e broncodilatadores de curta/longa ação conforme steps; doses baixa/média/alta por tabelas; LABA e eventual LAMA.
  - <5 anos: preferência por baixa dose; se necessário, dobrar (alta dose).
* Adesão em adolescentes
  - Dificuldade elevada; responsabilidade deve ser compartilhada e conduzida pelos pais.
### 5.

---

### Chunk 17/30
**Article:** pediatria funcional integrativa - parte IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.618

ão controlados, especialmente na asma neutrofílica (comum em obesos), e o ômega 3 pode ajudar a reduzir essa inflamação.
### 6. Fitoterápicos: Quercetina
- **Mecanismo:** Inibe a liberação de citocinas inflamatórias e de histamina pelos mastócitos (ação similar ao cromoglicato), além de regular a atividade da musculatura lisa.
- **Evidências e Segurança:** Estudos mostraram que a quercetina diminui sintomas e aumenta o peak flow. Doses seguras em adultos são de 500mg por até 12 semanas. Faltam estudos de segurança e dose em crianças.
### 7. Fitoterápicos: Cúrcuma na Asma e Rinite
- **Mecanismo:** A cúrcuma é segura e demonstrou diminuir marcadores inflamatórios (IL-4, TNF-alfa) e aumentar os anti-inflamatórios (IL-10).
- **Evidências:** Um estudo brasileiro com crianças mostrou melhora nos sintomas e redução no uso de medicação de resgate. Como 90-95% dos asmáticos têm rinite, tratar a rinite é fundamental.

---

### Chunk 18/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.617

nços em diagnóstico e monitoramento, o uso de corticosteroides inalatórios traz riscos mensuráveis (supressão do eixo HPA, efeitos em crescimento e massa óssea) que exigem vigilância sistemática e ajuste terapêutico. O quadro pede estratégias integradas de controle, educação e segurança terapêutica para reduzir morbidade e óbitos.
---
### Evidências-Chave
**Carga e controle da asma permanecem desproporcionais: alta prevalência, uso intensivo de serviços e controle clínico baixo.**
- Prevalência global: mais de 300 milhões de pessoas com asma; no Brasil, mais de 20 milhões.
- Atenção primária: 1,3 milhões de atendimentos por asma (DataSUS, 2021).
- Internações: 83 mil por asma (DataSUS, 2022).
- Óbitos: 574 por asma em 2022; tendência de redução de 7 óbitos/dia em 2015 para 3/dia até 2020.
- Controle no Brasil (critérios INAH): apenas 10% dos pacientes controlados; em países desenvolvidos, cerca de 50%.

---

### Chunk 19/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.614

Plano de Tratamento de Acompanhamento:**
    *   **Controle Ambiental:** Redução de mofo, poeira, pelos, produtos químicos.
    *   **Estratégia Terapêutica:** Preferir uso intermitente ou doses baixas de corticoide para preservar massa óssea e estatura.
    *   **Saúde Geral:** Tratamento de comorbidades (rinite, refluxo, obesidade) e vigilância para síndrome metabólica. Abordagem de medicina funcional (inflamação/metilação).
    *   **Educação:** Orientação sobre a doença, dispositivos e prevenção primária.

---

## Quantitative Data

### Narrativa Quantitativa
A asma é altamente prevalente globalmente e no Brasil, com grande impacto em crianças, gerando alto volume de atendimentos e internações, enquanto o controle clínico permanece insuficiente.

---

### Chunk 20/30
**Article:** Emagrecimento - Parte VIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.614

elevância clínica.
- Boswellia padronizada entrega mesma eficácia com menos cápsulas, favorecendo adesão.
- Suplementos lipídicos devem ser tomados com refeições para melhor absorção e conforto gástrico.
### Alavancas clínicas complementares
Protocolos simples e personalizados maximizam resultados em dor, inflamação e emagrecimento.
- Inalação direta supera difusão ambiental para efeitos terapêuticos de óleos essenciais.
- Beta-cariofileno da copaíba ativa CB2 e favorece analgesia e modulação inflamatória.
- Otimizar vitamina D melhora resistência insulínica e marcadores inflamatórios, com doses individualizadas por polimorfismos GC/VDR.

---

### Chunk 21/30
**Article:** pediatria funcional integrativa - parte IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.611

Chave
**A suplementação com Vitaminas D e K2 demonstra um potencial significativo no controle da asma e na mitigação de efeitos colaterais, embora a eficácia dependa de doses e níveis séricos adequados.**
- A manutenção de níveis de Vitamina D acima de 30 ng/ml é crucial, com estudos mostrando uma redução de 50% nas exacerbações de asma em crianças e uma diminuição de 25% na sibilância em crianças de até 3 anos quando a mãe suplementa durante a gestação.
- Níveis séricos de Vitamina D acima de 60 ng/ml são considerados excelentes para crianças asmáticas, embora doses baixas (500-1.000 UI/dia) não tenham mostrado eficácia nos estudos.
- A Vitamina K2 demonstrou eficácia marcante em 42% dos pacientes com asma leve e 30% com asma moderada em um estudo de 1975, além de prevenir a perda de massa óssea (de 0,66 para 0,55 g/cm²) em pacientes usando corticoides, com uma dose de 15 mg.

---

### Chunk 22/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.611

entes de corticoide oral).
2. **Histórico de Medicamentos:**
    *   Corticoides inalatórios (Budesonida, Fluticasona, Beclometasona - incluindo apresentação em nanopartículas).
    *   Corticoides orais.
    *   Cromoglicato de sódio e Nedocromil.
    *   Antifúngicos e Antivirais (inibidores de CYP3A4, ex: Ritonavir, antifúngicos azólicos) que causam interação medicamentosa.
## [Subjetivo:]
O quadro clínico aborda tanto a sintomatologia respiratória quanto os efeitos adversos do tratamento:
*   **Respiratório:** Episódios recorrentes de tosse (pode ser variante tosse), sibilância, aperto no peito e dispneia. Sintomas podem ser persistentes (leve, moderado, grave) ou intermitentes, com despertar noturno e limitação física.

---

### Chunk 23/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.606

de ~1 cm; diferença anual de ~0,2 cm; achados em revisão com quase 3.400 crianças, por 12–52 semanas; contínuo vs cromoglicato: ~1 cm a menos.
- Massa óssea: 90% incorporada aos 18 anos; pico aos 30; ≥5 ciclos de corticoide oral em 7 anos associaram-se à osteopenia (15% meninos; 22% meninas).
- Acompanhamento do crescimento: acompanhamento linear a cada 6 meses.
- BMI e ICS: aumento do IMC em 0,07 kg/m² por ano de uso; antecipação do reganho de IMC em 2 meses por ano de uso; retomada do IMC por volta dos 6 anos.
- Risco com doses baixas: supressão do eixo HPA pode ocorrer com doses tão baixas quanto 200 mcg de beclometasona (ou equivalente).
- Fluticasona: duas vezes mais absorvida; usar metade da dose da beclometasona/budesonida.
**[Achados Adicionais]**
- Dispositivo e deposição: spray tem 5 vezes mais deposição pulmonar do que pó seco.

---

### Chunk 24/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.602

eduzir carga inflamatória sistêmica.
* Comorbidades
  - Rinite alérgica (vias aéreas unidas), refluxo (associado a alergia alimentar/obesidade), obesidade (fenótipo neutrofílico), anemia/deficiência de ferro.
* Ferramentas práticas
  - Técnica correta de dispositivos; uso de espaçador; ACT em todas as visitas; espirometria e PFE quando disponíveis (PFE 2x/dia por 2 semanas).
### 10. Critérios para suspeita e rastreio de supressão do eixo HPA
* Quando suspeitar
  - Sintomas compatíveis; uso de alta dose 3–6 meses; crescimento monitorado a cada 6 meses mesmo em baixa dose; corticoide oral ≥2 semanas consecutivas ou >3 semanas em 6 meses; uso concomitante de inibidores de CYP3A4.
* Como rastrear
  - Cortisol às 8:00; se normal, reavaliar em 6 meses; se sintomático com cortisol normal, teste de estímulo com ACTH (alvo ≥18 µg/dL); valores muito baixos podem ocorrer (~3 µg/dL); saliva/urina com limitações em crianças.

---

### Chunk 25/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.602

# pediatria funcional integrativa - parte III

**Source:** https://web.plaud.ai/share/c90e1765417811903::YXdzOnVzLXdlc3QtMg

---

## Lecture

> Data e Hora: 2025-12-09 04:52:19
> Local: [Inserir Local]
> Instrutor(a): Dra. Rita [Inserir Nome]
## 📝 Resumo
A aula, criada em 2025-12-09, apresenta uma visão abrangente da asma com foco pediátrico, integrando a abordagem pneumológica tradicional (GINA) e estratégias funcionais/integrativas, com ênfase nos riscos e manejo dos corticoides inalatórios (ICS), incluindo supressão do eixo hipotálamo-hipófise-adrenal (HPA), efeitos sobre crescimento, massa óssea e obesidade.

---

### Chunk 26/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.601

critérios; reconhecer viroses como principal causa em <3 anos; foco em imunostimulação.
* Prevenção primária e secundária
  - Primária: gestação, tipo de parto, aleitamento; vitamina D e ômega-3 podem reduzir sibilância transitória.
  - Secundária: reduzir aeroalérgenos (mofo), eosinofilia; exposição a outras crianças/pets/ambiente de fazenda pode reduzir risco (microbiota).
### 6. Corticoide inalatório: benefícios, riscos e manejo do eixo HPA
* Benefícios do ICS
  - Reduz sintomas, hiperresponsividade e exacerbações; melhora função; diminui uso de corticoide oral e mortalidade.
* Efeitos adversos e mitigação
  - Candidíase (higiene oral), disfonia (espaçador), sistêmicos (obesidade, crescimento, massa óssea, supressão HPA).
  - Após certo ponto, aumentar dose eleva efeitos sistêmicos sem ganho proporcional; objetivo: menor dose eficaz.

---

### Chunk 27/30
**Article:** pediatria funcional integrativa - parte IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.598

Tarefas
- [ ] 1. Implementar uma dieta anti-inflamatória, livre de alérgenos, contaminantes e defensivos agrícolas.
- [ ] 2. Reduzir a exposição a gatilhos ambientais, incluindo poluentes, produtos químicos domésticos (ex: amaciantes), perfumaria e mofo.
- [ ] 3. Investigar e tratar possíveis intoxicações por metais pesados, como o arsênico.
- [ ] 4. Avaliar e corrigir os níveis de ferro, evitando tanto a deficiência (que mimetiza sintomas de asma) quanto o excesso (que é pró-inflamatório).
- [ ] 5. Considerar a suplementação de Vitamina K2 em pacientes em uso crônico de corticoides para prevenir a perda de massa óssea.
- [ ] 6. Manter os níveis de Vitamina D acima de 60 ng/ml através de suplementação diária, com atenção especial a crianças.
- [ ] 7. Avaliar o ômega-índex e suplementar ômega-3 para atingir níveis > 8%, especialmente em pacientes obesos, para reduzir a inflamação.
- [ ] 8.

---

### Chunk 28/30
**Article:** pediatria funcional integrativa - parte IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.597

 mega-índex e suplementar ômega-3 para atingir níveis > 8%, especialmente em pacientes obesos, para reduzir a inflamação.
- [ ] 8. Utilizar sulfato de magnésio como terapia coadjuvante em exacerbações graves e considerar seu papel na prevenção.
- [ ] 9. Incorporar fitoterápicos como cúrcuma, quercetina e Boswellia serrata como coadjuvantes, respeitando as doses seguras.
- [ ] 10. Focar na manutenção de uma microbiota saudável, especialmente em crianças na "janela de oportunidade", promovendo parto normal, amamentação e evitando o uso excessivo de antibióticos.
- [ ] 11. Investigar a presença de inflamação sistêmica (ex: medir PCR-US) em pacientes com asma de difícil controle.
- [ ] 12. Incorporar práticas de mindfulness e exercícios respiratórios para melhorar o controle da asma.

---

### Chunk 29/30
**Article:** MFI PÓS RACHEL GAIGER AULA 01 - OXIGÊNIO HIPERBÁRICO (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.595

cofatores mitocondriais (CoQ10, L-carnitina, ácido alfa-lipoico, glutationa, glicina, taurina, tirosina, MCT, vitaminas B1/B2/B3).
- Integração com terapias padrão e medicina funcional: redução de sintomas para permitir abordagem da causa raiz; sinergia com antibióticos/antifúngicos, esteroides em DII e reabilitação motora/cognitiva em casos neurológicos.
### Segurança, contraindicações e manejo de efeitos adversos
- Contraindicações relativas: doenças pulmonares crônicas (enfisema), pneumotórax recente, cirurgia cardíaca/trauma recente, neurite óptica, uso atual de bleomicina/doxorrubicina, otite média/dificuldade de equalização, implante coclear, marcapasso (avaliação individual), hipercapnia, hipertensão não controlada, transtornos convulsivos.
- Absolutas: pneumotórax, broncoespasmo agudo não resolvido, próteses orbitais específicas.
- Drogas incompatíveis: cisplatina, doxorrubicina; cautela com amiodarona e antiangiogênicos.

---

### Chunk 30/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.592

es em uso prolongado (3-6 meses) ou altas doses de corticoides (incluindo ciclos orais frequentes), há risco de supressão adrenal manifestada por fadiga, fraqueza, mal-estar, náusea, vômito, diarreia e dor abdominal. Relatos de preocupação com ganho de peso e retardo no crescimento. O paciente pode estar assintomático da asma, mas apresentando sinais sistêmicos do tratamento.
## [Objetivo:]
*   **Avaliação Pulmonar:**
    *   **Sinais Clínicos:** Sibilância à ausculta, sinais de desconforto respiratório em exacerbações.
    *   **Espirometria:** Redução do VEF1 pré-broncodilatador. Relação VEF1/CVF reduzida (<0,8 adultos; <0,9 crianças).
    *   **Prova Broncodilatadora:** Positiva se aumento do VEF1 > 12% e 200 ml (adultos) ou > 12% (crianças).
    *   **PFE:** Variabilidade > 10% (adultos) e 13% (crianças).
    *   **Testes de Desafio:** Redução na função pulmonar com metacolina, exercício ou frio.

---

