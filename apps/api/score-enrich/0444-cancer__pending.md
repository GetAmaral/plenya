# ScoreItem: Câncer

**ID:** `c77cedd3-2800-7e0c-b44f-ccf9c3b926ef`
**FullName:** Câncer (Histórico Familiar de Doenças - Parentes próximos (pais, irmãos, tios, avós, filhos, netos))

**Preparation Metadata:**
- Quality Grade: **EXCELLENT**
- Total Chunks: 30 de 10 artigos
- Avg Similarity: 0.609

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `c77cedd3-2800-7e0c-b44f-ccf9c3b926ef`.**

```json
{
  "score_item_id": "c77cedd3-2800-7e0c-b44f-ccf9c3b926ef",
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

**ScoreItem:** Câncer (Histórico Familiar de Doenças - Parentes próximos (pais, irmãos, tios, avós, filhos, netos))

**30 chunks de 10 artigos (avg similarity: 0.609)**

### Chunk 1/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.707

íveis a fatores genéticos de alta/média penetrância, enquanto 90% estão ligados a fatores não genéticos e potencialmente modificáveis. Ferramentas de cálculo de risco e critérios clínicos ajudam a identificar quem está em maior risco ao longo da vida (≥20%) ou no curto prazo (Gail 5 anos ≥1,7), orientando prevenção, rastreamento e decisões personalizadas.
---
### Evidências-Chave
**A maior parte dos casos de câncer de mama decorre de fatores não genéticos, mas um subconjunto relevante tem risco hereditário elevado que requer atenção específica.**
- Genética relacionada ao câncer de mama corresponde a cerca de 10% dos tumores diagnosticados, indicando a fração atribuída a fatores genéticos versus outros fatores.
- Os outros 90% dos casos não têm achados genéticos e são relacionados ao estilo de vida, enfatizando a influência de fatores modificáveis.

---

### Chunk 2/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.701

a mutações genéticas conhecidas; os demais 90% não apresentam achados genéticos e são majoritariamente atribuídos ao estilo de vida.
   - Entre as mutações associadas a maior incidência estão BRCA1/2 e TP53; em geral afetam genes supressores tumorais, levando à perda de defesa contra células alteradas e aumento da incidência.
* **Penetrância genética**
   - Alta penetrância: confere chance ≥ 40% de desenvolver câncer de mama ao longo da vida.
   - Penetrância moderada: cerca de 20–25%.
   - Baixa penetrância: < 20%.
   - Nem todas as mutações identificadas implicam mudança prática no acompanhamento; o valor clínico depende da magnitude do risco conferido.
* **Contexto familiar BRCA positivo e decisões clínicas**
   - Em famílias com múltiplos casos e mutação BRCA, o risco é substancial mesmo com intervenções, inclusive cirurgias profiláticas.

---

### Chunk 3/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.683

urgias profiláticas) e a importância do aconselhamento genético. Apresenta calculadoras de risco (Gail como a mais conhecida), observa superestimação fora de populações de origem e sugere uso pedagógico para impulsionar mudanças de estilo de vida. Destaca anamnese detalhada, avaliação de composição corporal e marcadores metabólicos/inflamatórios como base prática de estratificação. Conclui que genética não é destino, introduz epigenética como fator modificável e informa que dúvidas sobre reposição hormonal serão abordadas na próxima aula. Data de criação: 2025-11-21.
## 🔖 Pontos de Conhecimento
### 1. Genética e câncer de mama
* **Proporção de câncer de mama ligado à genética**
   - Aproximadamente 10% dos tumores de mama diagnosticados relacionam-se a mutações genéticas conhecidas; os demais 90% não apresentam achados genéticos e são majoritariamente atribuídos ao estilo de vida.

---

### Chunk 4/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.677

RCA2, TP53) aumentam significativamente o risco, porém são raras. Epigenética e estilo de vida são cruciais e modificáveis. A avaliação individualizada requer anamnese detalhada, composição corporal e marcadores laboratoriais.
- Diagnóstico Suspeito: Nenhum no momento

## Plano:
- Prescrição: Inserir mais aqui
- Próximos Passos:
  - Suspeitar de mutação genética e considerar teste genético em casos de múltiplos cânceres na família (mama, ovário, pâncreas, próstata), câncer de mama triplo negativo, câncer de mama em idade jovem (<45 anos), câncer de mama em homem, ou descendência judaica Ashkenazi.
  - Encaminhar para aconselhamento genético antes do teste em pacientes com alta suspeita de mutação.
  - Utilizar calculadoras de risco (ex.: Gail, Tyrer-Cuzick) para conscientizar sobre mudanças no estilo de vida, reconhecendo limitações na população brasileira.

---

### Chunk 5/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.675

 ão BRCA, o risco é substancial mesmo com intervenções, inclusive cirurgias profiláticas.
   - Decisões sobre cirurgia profilática devem ser individualizadas e respeitar a vivência e história da paciente; medicina humanizada reconhece o contexto familiar e emocional.
* **Critérios de suspeita para mutação genética**
   - Vários casos de câncer na família (especialmente mama), independentemente do lado materno/paterno.
   - Câncer de mama triplo negativo.
   - Câncer de mama em idade jovem (abaixo de 45 anos).
   - Presença de câncer de ovário, pâncreas, próstata, ou câncer de mama em homem na família.
   - Descendência judaica Asquenazita, com maior prevalência de mutações BRCA.
* **Interpretação de resultados genéticos negativos**
   - Resultado negativo não exclui risco familiar elevado; exemplificado por caso de “sétima irmã” com câncer de mama sem mutação detectada.

---

### Chunk 6/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.662

dos sintomas gerais a investigar em consulta, como dor mamária e TPM intensa, que podem indicar predominância estrogênica.

## Objetivo:
Não há achados de exame físico de um paciente específico. A abordagem objetiva descrita inclui:
- Avaliação da composição corporal por bioimpedância ou densitometria de corpo total, identificando condições como sarcopenia em mulheres com peso normal.
- Solicitação de exames laboratoriais para avaliar marcadores inflamatórios e verificar se a paciente está metabolicamente doente ou em risco.

## Diagnóstico Primário:
- Avaliação: Discussão educacional sobre estratificação de risco para câncer de mama. 90% dos casos relacionam-se ao estilo de vida e 10% a fatores genéticos conhecidos. Mutações de alta penetrância (ex.: BRCA1, BRCA2, TP53) aumentam significativamente o risco, porém são raras. Epigenética e estilo de vida são cruciais e modificáveis.

---

### Chunk 7/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.657

eriência pessoal indica superestimação (risco “quase elevado” obtido sem incluir história familiar).
### 4. Aconselhamento genético e solicitação de testes
* **Processo e preparo da paciente**
   - Ao solicitar teste genético, é crucial documentar o motivo e encaminhar para aconselhamento genético quando a suspeita é alta.
   - Resultados positivos alteram a história da família e da descendência; pacientes devem estar preparadas emocional e informacionalmente para receber o resultado.
* **Estratégia de testagem familiar**
   - Quando há mutação identificada no caso índice, faz sentido testar parentes (filhos, irmãs, mãe).
   - Sem mutação identificada, testar familiares pode não trazer valor prático, apesar de alto risco agregado pela história.
### 5.

---

### Chunk 8/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.634

é maior do que 20%, critério de classificação para manejo intensificado.
**Calculadoras e critérios clínicos complementam a genética para identificar risco elevado e orientar prevenção.**
- Na calculadora de Gail, considera-se risco alto quando em cinco anos o risco é de 1,7, limiar utilizado para motivar adesão à mudança de estilo de vida e estratificação de risco.
- Câncer de mama em idade jovem normalmente se estipula abaixo dos 45 anos, critério de suspeita para indicação de investigação de mutação.
**Diversidade genética e experiência prática destacam tanto o poder quanto as limitações dos testes e modelos.**
- Estima-se que cada indivíduo tenha em torno de 50 milhões de variantes (SNPs), contextualizando a diversidade genética usada em bancos de dados para estratificação de risco.
- Relato pessoal: seis pessoas muito próximas na família com mutação BRCA, ilustrando impacto familiar de mutações de alta penetrância.

---

### Chunk 9/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.619

tica e modificabilidade do risco
* **Genética não é destino**
   - O risco pode ser modificado por fatores epigenéticos e de estilo de vida; epigenética explica herança de alterações não mutacionais, influenciadas por ambiente e vivências.
* **Focos de intervenção epigenética**
   - Melhorar processos de metilação e o “contexto” epigenético por meio de intervenções em fases críticas (ex.: impactos de má gestação no futuro).
   - Estratégias epigenéticas são um caminho prático para reduzir risco em cenários não explicados por genética clássica.
### 6. Anamnese, avaliação clínica e laboratorial para estratificação
* **Anamnese detalhada**
   - Coleta de: patologias mamárias prévias, biópsias e seus resultados; história pessoal e familiar oncológica; uso de medicamentos/suplementos; hábitos de vida (alimentação, exercício, álcool, tabaco); saúde intestinal; níveis de estresse/ansiedade; rede de apoio.

---

### Chunk 10/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.618

namnese detalhada abrangendo hábitos de vida, saúde intestinal, sintomas hormonais (dor mamária, TPM), história oncológica familiar (ambos os lados) e rede de apoio.
- [ ] Avaliar composição corporal com bioimpedância ou densitometria de corpo total para detectar sarcopenia e orientar intervenção.
- [ ] Solicitar e acompanhar marcadores inflamatórios/metabólicos para monitorar resposta a intervenções de estilo de vida, especialmente em pacientes com diagnóstico prévio de câncer.
- [ ] Utilizar calculadoras de risco (p. ex., Gail e ferramentas que integram PRS e fatores de estilo de vida) com cautela, reconhecendo superestimação em populações brasileiras, e empregá-las para motivar mudanças comportamentais.
- [ ] Considerar farmacogenética ao selecionar terapias, avaliando polimorfismos que alteram metabolização de drogas.
- [ ] Preparar materiais e plano para a próxima aula focada em dúvidas sobre reposição hormonal.

---

### Chunk 11/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.612

ificar adoecimento metabólico ou risco, e para monitorar evolução após intervenções.
   - Reduzir inflamação crônica e síndrome metabólica, com mudanças de estilo de vida, é crucial na prevenção e ainda mais importante em pacientes já diagnosticadas com câncer.
### 7. Reposição hormonal e dúvidas futuras
* **Reflexão crítica**
   - Questiona se a reposição hormonal é o “verdadeiro problema” no câncer de mama, frente ao conjunto de fatores abordados.
   - Próxima aula tratará das principais dúvidas sobre reposição hormonal, sinalizando continuidade do conteúdo.
## ❓ Perguntas
- [Insert Question/Confusion]
## 📚 Próximos Arranjos
- [ ] Encaminhar pacientes com alta suspeita de mutação para aconselhamento genético antes da testagem.
- [ ] Realizar anamnese detalhada abrangendo hábitos de vida, saúde intestinal, sintomas hormonais (dor mamária, TPM), história oncológica familiar (ambos os lados) e rede de apoio.

---

### Chunk 12/30
**Article:** MFI - Reposição Hormonal - AULA 07 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.612

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

### Chunk 13/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.605

ra conscientizar sobre mudanças no estilo de vida, reconhecendo limitações na população brasileira.
- Acompanhamento e Tratamento:
  - Focar em medicina humanizada e individualizada, respeitando decisões da paciente (ex.: cirurgia profilática).
  - Intervir em fatores de risco modificáveis (estilo de vida) para reduzir inflamação crônica e síndrome metabólica.
  - Realizar anamnese detalhada sobre patologias anteriores, história familiar, uso de medicamentos, hábitos de vida, saúde intestinal, estresse e sintomas hormonais.
  - Monitorar evolução por marcadores laboratoriais e avaliação da composição corporal após intervenções.

---

## Quantitative Data

### Narrativa Quantitativa
A história central mostra que apenas cerca de 10% dos cânceres de mama são atribuíveis a fatores genéticos de alta/média penetrância, enquanto 90% estão ligados a fatores não genéticos e potencialmente modificáveis.

---

### Chunk 14/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.601

néticos e são relacionados ao estilo de vida, enfatizando a influência de fatores modificáveis.
**Penetrância genética estrutura a estratificação de risco: alta (~≥40%), moderada (20–25%) e baixa (<20%), com “alto risco” operacional nas calculadoras definido por ≥20% ao longo da vida.**
- Definição de gene de alta penetrância: confere chance em torno de 40% ou mais de desenvolver câncer de mama ao longo da vida; serve como limiar para classificar risco genético elevado.
- Penetrância moderada definida entre 20 a 25 de risco de câncer de mama, categoria intermediária entre alta e baixa.
- Genes de baixa penetrância: abaixo de 20% de risco ao longo da vida, usado para diferenciar categorias de penetrância.
- Risco ao longo da vida considerado alto nas calculadoras é maior do que 20%, critério de classificação para manejo intensificado.

---

### Chunk 15/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.594

 o exclui risco familiar elevado; exemplificado por caso de “sétima irmã” com câncer de mama sem mutação detectada.
   - A ausência de mutação identificável pode refletir limitações do conhecimento atual; não justifica testar descendentes quando nada foi encontrado no caso índice, ao contrário da situação com mutação identificada.
### 2. SNPs e Score de Risco Poligênico
* **Definição e papel das SNPs**
   - SNPs (polimorfismos de nucleotídeo único) são alterações em uma “letra” do código genético; podem conferir proteção ou aumentar risco para doenças como câncer de mama.
   - Cada indivíduo possui cerca de 50 milhões de variantes; dados populacionais associam variantes específicas ao aumento de risco em doenças complexas (ex.: coronarianas, câncer de mama).
* **Score de risco poligênico (PRS)**
   - Combina múltiplas variantes de pequeno efeito para estimar risco individual; permite classificar pacientes em faixas de risco.

---

### Chunk 16/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.593

para estratificação de risco.
- Relato pessoal: seis pessoas muito próximas na família com mutação BRCA, ilustrando impacto familiar de mutações de alta penetrância.
- Relato pessoal: sete irmãs com câncer de mama na mesma família sem mutação identificada no teste, exemplificando limitações dos testes genéticos atuais e risco familiar não explicado.
- O orador trouxe cinco calculadoras de risco, detalhando três, mostrando a variedade de ferramentas existentes para estimar risco de câncer de mama e critérios adicionais como IMC e score poligênico.
**Constatações Adicionais**
- Número de calculadoras de risco apresentadas: seis.

---

### Chunk 17/30
**Article:** Family history assessment significantly enhances delivery of precision medicine in the genomics era (2020)
**Journal:** Genome Medicine
**Section:** abstract | **Similarity:** 0.592

Family history remains a crucial component of precision medicine, providing information beyond what can be captured by genomic testing alone. This review demonstrates how family history assessment integrates genetic and environmental risk factors to improve disease risk stratification.

---

### Chunk 18/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.588

luto, evitar alarmismo, e definir critérios para exames complementares (ressonância) apenas quando houver fatores de risco adicionais.
- [ ] 4. Atualizar materiais educativos para esclarecer que história familiar, por si só, não contraindica reposição; incorporar achados do Sister Study e WHI.
- [ ] 5. Estabelecer diretriz interna: não indicar reposição hormonal sistêmica em pacientes com histórico de câncer de mama; considerar terapias tópicas para atrofia vaginal após tentativa de métodos não hormonais, com suporte emocional.
- [ ] 6. Criar protocolo de uso criterioso de gestrinona em endometriose e mastalgia refratária, com consentimento informado sobre lacunas de evidência oncológica.
- [ ] 7. Definir critérios de indicação de testosterona por motivos não oncológicos, evitando prescrição para “redução de risco mamário” até que haja validação em guidelines.
- [ ] 8.

---

### Chunk 19/30
**Article:** Trato Gastrointestinal V – Intestino Delgado I (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.583

as e risco biológico: atualização dietética e preservação telomérica são determinantes de desfechos cardiovasculares e infecciosos.**
- A dieta tradicional de diabetes com 60% de carboidratos integrais é criticada como obsoleta, motivando revisões para melhor controle metabólico.
- Telômeros curtos associam-se a aumento de 300% no risco de morte cardíaca e 800% em doenças infecciosas, ressaltando a importância de estratégias protetoras.
**Achados-Chave Adicionais**
- Estudo pediátrico (2016): 174 crianças de 1–4 anos, 12 semanas, randomizado duplo-cego e placebo-controlado com beta-glucana, observando redução de episódios de doenças comuns.
- Idade do primeiro câncer de mama familiar: 35 anos na irmã gêmea da paciente, ilustrando risco familiar e impacto psicológico em decisões de prevenção/terapias.
- Espera inicial de dois meses antes de análogos de GLP-1 serve como janela de avaliação da eficácia de intervenções não farmacológicas.

---

### Chunk 20/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.578

# Mastologia II

**Source:** https://web.plaud.ai/share/c7c71765255680986::YXdzOnVzLXdlc3QtMg

---

## Lecture

> Data e Hora: 2025-11-21 03:11:20
> Local: [Inserir Local]
> Instrutor(a): [Inserir Nome]
## 📝 Resumo
A palestra trata da estratificação do risco de câncer de mama com foco em genética, polimorfismos (SNPs), scores de risco poligênico, calculadoras clínicas de risco, epigenética e estilo de vida. A instrutora enfatiza que cerca de 10% dos tumores de mama estão ligados a mutações de alta penetrância (como BRCA1/2 e TP53), enquanto 90% se associam majoritariamente ao estilo de vida. Explica penetrância, limitações dos scores poligênicos por vieses populacionais e o valor das SNPs na farmacogenética. Defende decisões individualizadas e humanizadas (incluindo cirurgias profiláticas) e a importância do aconselhamento genético.

---

### Chunk 21/30
**Article:** Bases Metabólicas das Doenças Crônicas e Gerenciamento - Inflamação 2 (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.576

 tica de PCR-us e na intervenção proativa. Antecipação: próxima aula focará epigenética/metilação e exames correlatos.
## 🔖 Pontos de Conhecimento
### 1. Interação entre inflamação, imunidade, microbiota e câncer
- Cross-talk em Nature Reviews Cancer: inflamação sustenta comunicação bidirecional entre sistema imune, tumores e micro-organismos.
- Três eixos geradores de inflamação: perda da barreira intestinal (disbiose e ativação de TLR), alimentação mecanística equivocada e inflamação mediada por gordura corporal (inclui desequilíbrio ômega 6/ômega 3).
- Meta-análises: PCR-us como principal marcador de inflamação crônica associada a maior risco de câncer (colorretal, mama) e DCV; IL-6, fibrinogênio e TNF-α também relevantes; pulmão (IL-6/fibrinogênio), próstata/ovário (fibrinogênio/PCR).
- Interpretação prática: medir PCR-us regularmente e integrar prevenção dietética/suplementar e estilo de vida.
### 2.

---

### Chunk 22/30
**Article:** Mastologia I (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.576

** Preocupação existe, porém é menor que a radiação terrestre ou de um escore de cálcio; se fosse principal causa, haveria mais câncer em idades avançadas, mas a idade média vem diminuindo.
    - **Câncer Avançado:** Rastreio não reduziu significativamente a incidência de doença metastática.
    - **Falsos Positivos e Overtreatment:** Aumenta biópsias desnecessárias e diagnósticos de tumores indolentes (como alguns CIS), gerando tratamento excessivo.
*   **Futuro do Diagnóstico**
    - O rastreamento atual tende a mudar.
    - Novas estratégias em estudo: análise de lágrimas, aspirado do mamilo para análise genética, detecção de células tumorais circulantes/DNA tumoral.
    - Medicina de precisão com estratificação de risco individualizada.

### 4. Fatores de Risco e Prevenção
*   **Analogia do Carro**
    - Dirigir alcoolizado e na contramão aumenta risco; dirigir corretamente não elimina acidentes.

---

### Chunk 23/30
**Article:** Emagrecimento XV (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.571

genética: genes modulam risco; estilo de vida determina desfechos.
- Mecanismos citados: acetilação/desacetilação de histonas e metilação (sem detalhamento técnico).
- Fitoquímicos e nutrientes adequados são necessários; evitar excessos e deficiências.
- Takeaways práticos (“não negociações”): sono consistente, fibra diária, proteína adequada, manejo de estresse.
### 2. Interpretação de testes genéticos: alelos de risco, homozigose/heterozigose e RS
- Painéis genéticos apresentam variantes com dois alelos; bancos definem alelos de risco.
- Homozigose para alelo de risco sugere efeito mais forte; heterozigose, risco intermediário.
- “RS” identifica variantes específicas com maior base de evidências; testes muito amplos podem gerar alarmismo e reduzir especificidade.
- Foco nos RS principais de FTO (ex.: rs9939609) e MC4R (RS que “termina com 13”) ao interpretar laudos.

---

### Chunk 24/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.567

câncer de mama; há pelo menos cinco, com três detalhadas pela instrutora.
   - Problema comum: desenvolvidas em outras populações, tendem a superestimar risco em brasileiros.
* **Uso pragmático**
   - Utilizadas para engajar pacientes em mudanças de estilo de vida, demonstrando risco acima do esperado (“assustar” como figura de linguagem para promover adesão).
* **Exemplos e limiares**
   - Calculadora de Gail: a mais famosa e primeira; risco de 1,7% em cinco anos é considerado alto.
   - Outras calculadoras incorporam critérios como IMC e alterações proliferativas mamárias.
   - Uma calculadora online mais recente integra PRS e fatores de estilo de vida; é a preferida da instrutora quando utiliza essas ferramentas.
   - Risco alto ao longo da vida costuma ser > 20%; experiência pessoal indica superestimação (risco “quase elevado” obtido sem incluir história familiar).
### 4.

---

### Chunk 25/30
**Article:** Bases Metabólicas das Doenças Crônicas e Gerenciamento - Inflamação 2 (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.567

resistência insulínica. Apresenta ensaios clínicos e meta-análises que demonstram redução de PCR-us, IL-6 e LDL/triglicerídeos, além de melhora de HDL, FRAP/TRAP, HOMA-IR, adiponectina e BHB. Aborda a anemia da inflamação e suas diferenças laboratoriais em relação à deficiência de ferro. Propõe uma abordagem integrada de prevenção e manejo que combina personalização dietética (low carb, cetogênica, mediterrânea, plant-based), suplementação baseada em evidência (EPA/DHA, curcumina padronizada com piperina ou lipossomada, antocianinas padronizadas, polifenóis diversos), modulação do tônus parassimpático e atividade física para proteção metabólica e imunológica. Destaca a importância do oncologista e do cardiometabologista preventivos na medição sistemática de PCR-us e na intervenção proativa. Antecipação: próxima aula focará epigenética/metilação e exames correlatos.
## 🔖 Pontos de Conhecimento
### 1.

---

### Chunk 26/30
**Article:** Capturing additional genetic risk from family history for improved polygenic risk prediction (2022)
**Journal:** Communications Biology
**Section:** abstract | **Similarity:** 0.565

Study demonstrates that family history captures genetic risk beyond polygenic risk scores, identifying individuals at elevated risk for cancers and cardiovascular diseases more effectively when both approaches are combined.

---

### Chunk 27/30
**Article:** Mastologia I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.563

e estilo de vida.
- Mantenha hormônios em faixa ótima para reduzir risco por desbalanços.
- Avance para estratificação de risco com biomarcadores como DNA tumoral circulante, reduzindo dependência de imagem.

---

### Chunk 28/30
**Article:** Mastologia I (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.563

estão associadas a um aumento drástico no risco de mortalidade em pacientes com câncer.**
- Mulheres com hiperinsulinemia apresentaram um risco 34% maior de desenvolver câncer e um risco 78% maior de morte após o diagnóstico, independentemente do IMC ou da circunferência abdominal.
- Pacientes com sarcopenia (perda de massa muscular) tiveram um aumento de 93% nas mortes por câncer em geral e, especificamente em casos de câncer de mama, a mortalidade foi 41% maior.
- Uma meta-análise também mostrou que a sarcopenia aumentou em 44% as mortes por todas as causas.
**A métrica de "sobrevida em 5 anos", embora comum em oncologia, pode ser enganosa devido a vieses estatísticos relacionados ao momento do diagnóstico.**
- A sobrevida em 5 anos é uma métrica frequentemente usada para avaliar a eficácia percebida do rastreamento mamográfico.

---

### Chunk 29/30
**Article:** Dieta Cetogênica - Parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.562

atamento metabólico do câncer
   - Terapias de pressão (contínuas): dieta cetogênica, cetonas exógenas, suplementos/fitoterápicos/drogas individualizadas, manejo do estresse emocional.
   - Terapias de pulso (intermitentes): inibição de glicose, inibição de glutamina, oxigenoterapia hiperbárica, entre outras.
   - Abordagem integrada e personalizada para maximizar o controle tumoral.
* Ensaio clínico randomizado (2021) em câncer de mama
   - 80 pacientes tratados com quimio; randomização para dois grupos; intervenção cetogênica/metabólica por 12 semanas; exames laboratoriais e de imagem no início e 12 semanas; cirurgia e reestadiamento para doença localmente avançada após quimio.
   - Resultados: redução de TNF-α, IGF-1, insulina; aumento de IL-10; redução significativa do tamanho tumoral no grupo cetogênico.

---

### Chunk 30/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.560

cer de mama).
* **Score de risco poligênico (PRS)**
   - Combina múltiplas variantes de pequeno efeito para estimar risco individual; permite classificar pacientes em faixas de risco.
   - Limitação central: natureza comparativa baseada em bancos de dados predominantemente europeus; aplicabilidade à população brasileira pode ser comprometida e introduzir vieses.
* **Aplicações atuais das SNPs**
   - Farmacogenética: identificação de polimorfismos que alteram metabolização de drogas, permitindo selecionar medicações e doses com maior precisão e menor toxicidade.
   - Já útil na prática e com importância crescente futura.
### 3. Calculadoras clínicas de risco
* **Panorama e finalidade**
   - Calculadoras semelhantes à “escala de Framingham” foram desenvolvidas para câncer de mama; há pelo menos cinco, com três detalhadas pela instrutora.
   - Problema comum: desenvolvidas em outras populações, tendem a superestimar risco em brasileiros.

---

