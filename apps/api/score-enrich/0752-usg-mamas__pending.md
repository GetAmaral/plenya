# ScoreItem: USG mamas

**ID:** `c77cedd3-2800-71a2-b317-07fb2e53253e`
**FullName:** USG mamas (Exames - Imagem)
**Unit:** Categoria

**Preparation Metadata:**
- Quality Grade: **EXCELLENT**
- Total Chunks: 30 de 8 artigos
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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `c77cedd3-2800-71a2-b317-07fb2e53253e`.**

```json
{
  "score_item_id": "c77cedd3-2800-71a2-b317-07fb2e53253e",
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

**ScoreItem:** USG mamas (Exames - Imagem)
**Unidade:** Categoria

**30 chunks de 8 artigos (avg similarity: 0.609)**

### Chunk 1/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.724

dos sintomas gerais a investigar em consulta, como dor mamária e TPM intensa, que podem indicar predominância estrogênica.

## Objetivo:
Não há achados de exame físico de um paciente específico. A abordagem objetiva descrita inclui:
- Avaliação da composição corporal por bioimpedância ou densitometria de corpo total, identificando condições como sarcopenia em mulheres com peso normal.
- Solicitação de exames laboratoriais para avaliar marcadores inflamatórios e verificar se a paciente está metabolicamente doente ou em risco.

## Diagnóstico Primário:
- Avaliação: Discussão educacional sobre estratificação de risco para câncer de mama. 90% dos casos relacionam-se ao estilo de vida e 10% a fatores genéticos conhecidos. Mutações de alta penetrância (ex.: BRCA1, BRCA2, TP53) aumentam significativamente o risco, porém são raras. Epigenética e estilo de vida são cruciais e modificáveis.

---

### Chunk 2/30
**Article:** MFI - Reposição Hormonal - AULA 07 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.688

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

### Chunk 3/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.647

m pacientes com nódulos mamários benignos (BI-RADS 3 ou fibroadenomas).
- Preocupação com surgimento de novos nódulos durante RH.
- RH em pacientes com mamas densas.
- RH em pacientes com histórico familiar de câncer de mama.
- Uso de gestrinona, testosterona e terapia hormonal tópica.
- RH após tratamento de câncer de mama.
- Queixas de atrofia vaginal em pacientes pós-câncer.
## Objetivo:
A transcrição não contém achados de exame físico individual. O médico aborda conceitos e evidências de estudos:
- Estrogênio e progesterona têm efeito proliferativo, responsável por benefícios da RH e desenvolvimento mamário; insulina e IGF-1 também influenciam a glândula mamária.
- Estudos como WHI (Women's Health Initiative) e One Million Study geraram controvérsia; ensaios clínicos randomizados (p.ex., WHI) não mostraram o mesmo aumento de risco de câncer de mama observado em estudos observacionais.

---

### Chunk 4/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.640

luto, evitar alarmismo, e definir critérios para exames complementares (ressonância) apenas quando houver fatores de risco adicionais.
- [ ] 4. Atualizar materiais educativos para esclarecer que história familiar, por si só, não contraindica reposição; incorporar achados do Sister Study e WHI.
- [ ] 5. Estabelecer diretriz interna: não indicar reposição hormonal sistêmica em pacientes com histórico de câncer de mama; considerar terapias tópicas para atrofia vaginal após tentativa de métodos não hormonais, com suporte emocional.
- [ ] 6. Criar protocolo de uso criterioso de gestrinona em endometriose e mastalgia refratária, com consentimento informado sobre lacunas de evidência oncológica.
- [ ] 7. Definir critérios de indicação de testosterona por motivos não oncológicos, evitando prescrição para “redução de risco mamário” até que haja validação em guidelines.
- [ ] 8.

---

### Chunk 5/30
**Article:** Mastologia I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.625

e estilo de vida.
- Mantenha hormônios em faixa ótima para reduzir risco por desbalanços.
- Avance para estratificação de risco com biomarcadores como DNA tumoral circulante, reduzindo dependência de imagem.

---

### Chunk 6/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** methods | **Similarity:** 0.623

; uso de medicamentos/suplementos; hábitos de vida (alimentação, exercício, álcool, tabaco); saúde intestinal; níveis de estresse/ansiedade; rede de apoio.
   - Atenção a sintomas de predominância estrogênica: dor mamária, TPM intensa; lacunas de cuidado em pacientes com alterações fibrocísticas e cistos, frequentemente não acolhidas apesar de sintomas de excesso de estrogênio circulante.
* **Avaliação de composição corporal**
   - Preferência por métodos além da balança: bioimpedância (utilizada pela instrutora) ou densitometria de corpo total.
   - Observação de sarcopenia em mulheres com peso normal e desconhecimento da condição; necessidade de intervenção.
* **Marcadores laboratoriais e metabolicidade**
   - Solicitar marcadores inflamatórios para identificar adoecimento metabólico ou risco, e para monitorar evolução após intervenções.

---

### Chunk 7/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.621

RCA2, TP53) aumentam significativamente o risco, porém são raras. Epigenética e estilo de vida são cruciais e modificáveis. A avaliação individualizada requer anamnese detalhada, composição corporal e marcadores laboratoriais.
- Diagnóstico Suspeito: Nenhum no momento

## Plano:
- Prescrição: Inserir mais aqui
- Próximos Passos:
  - Suspeitar de mutação genética e considerar teste genético em casos de múltiplos cânceres na família (mama, ovário, pâncreas, próstata), câncer de mama triplo negativo, câncer de mama em idade jovem (<45 anos), câncer de mama em homem, ou descendência judaica Ashkenazi.
  - Encaminhar para aconselhamento genético antes do teste em pacientes com alta suspeita de mutação.
  - Utilizar calculadoras de risco (ex.: Gail, Tyrer-Cuzick) para conscientizar sobre mudanças no estilo de vida, reconhecendo limitações na população brasileira.

---

### Chunk 8/30
**Article:** Terapia de Reposição Hormonal Feminina III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.615

om transvaginal para monitorização.
*   **Abordagem Holística:**
    - A saúde da mulher depende de múltiplos pilares: nutrição, atividade física, gestão do stress, sono, saúde intestinal, desintoxicação e equilíbrio hormonal.
## ❓ Perguntas
- [Inserir Pergunta/Dúvida/Confusão]
## 📚 Tarefas
- [ ] 1. Avaliar e modular fatores de risco modificáveis: dieta (aumentar crucíferas), exercício, peso corporal, consumo de álcool e tabaco.
- [ ] 2. Ao prescrever TRH, optar pela combinação de 17-beta-estradiol transdérmico e progesterona natural micronizada oral.
- [ ] 3. Para mulheres em TRH, realizar monitorização anual com mamografia, ultrassom de mamas e ultrassom transvaginal.
- [ ] 4. Melhorar a saúde intestinal através de uma dieta rica em fibras, hidratação e, se necessário, pré/probióticos para otimizar a eliminação de hormonas.
- [ ] 5.

---

### Chunk 9/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.612

clínicos randomizados (p.ex., WHI) não mostraram o mesmo aumento de risco de câncer de mama observado em estudos observacionais.
- Progestágeno (ex.: medroxiprogesterona) foi associado a aumento do risco mamário, enquanto progesterona micronizada não demonstrou o mesmo efeito.
- Nódulos BI-RADS 3 têm risco muito baixo de malignidade (~0,8% no ACRIN 666); acompanhamento semestral no primeiro ano e depois regular é o padrão.
- Densidade mamária é critério mamográfico com ligeiro aumento de incidência de câncer de mama, mas risco absoluto baixo.
- Biópsia a vácuo é eficaz para diagnosticar nódulos, com precisão comparável à biópsia cirúrgica.
- Histórico familiar aumenta risco pessoal, mas estudos (Sister Study, WHI) indicam que RH não aumenta adicionalmente esse risco em pacientes com histórico familiar positivo.

---

### Chunk 10/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.612

urgias profiláticas) e a importância do aconselhamento genético. Apresenta calculadoras de risco (Gail como a mais conhecida), observa superestimação fora de populações de origem e sugere uso pedagógico para impulsionar mudanças de estilo de vida. Destaca anamnese detalhada, avaliação de composição corporal e marcadores metabólicos/inflamatórios como base prática de estratificação. Conclui que genética não é destino, introduz epigenética como fator modificável e informa que dúvidas sobre reposição hormonal serão abordadas na próxima aula. Data de criação: 2025-11-21.
## 🔖 Pontos de Conhecimento
### 1. Genética e câncer de mama
* **Proporção de câncer de mama ligado à genética**
   - Aproximadamente 10% dos tumores de mama diagnosticados relacionam-se a mutações genéticas conhecidas; os demais 90% não apresentam achados genéticos e são majoritariamente atribuídos ao estilo de vida.

---

### Chunk 11/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.610

100%, resolvendo muitas questões e muitas vezes removendo a lesão sem necessidade de sala cirúrgica.
### 3. Densidade mamária
* **Definição e critérios**
   - Densidade mamária é critério mamográfico; não pode ser diagnosticada sem a primeira mamografia.
   - Mamas densas em pacientes jovens são esperadas; critério torna-se mais relevante em rastreio, geralmente a partir dos 50 anos.
* **Risco relativo versus absoluto**
   - Estudos associam densidade aumentada a maior incidência de câncer (com razão de risco frequentemente citada como 4 a 6 vezes ao comparar mama muito densa com lipossubstituída), mas o risco absoluto é baixo: exemplo citado, de 10% para 10,6%.
   - O problema maior é a dificuldade diagnóstica em mamas muito densas; exames complementares, como ressonância magnética, podem ser úteis em pacientes com mamas densas e fatores de risco adicionais.

---

### Chunk 12/30
**Article:** Mastologia I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.609

s em crescimento causam hipóxia, necrose e liberação de citocinas, atraindo macrófagos e sustentando inflamação mamária.
*   **Resistência à Insulina e Síndrome Metabólica**
    - Mulheres com IMC normal mas com síndrome metabólica têm maior incidência.
    - Hiperinsulinemia aumenta risco de câncer em 34% e risco de morte após câncer em 78%, independentemente de IMC/circunferência abdominal.
    - Inflamação crônica em mulheres eutróficas é um fator chave.
    - A insulina é um hormônio de atenção; é inadmissível que mastologistas não identifiquem resistência à insulina nas pacientes.
*   **Câncer como Doença Metabólica**
    - Base metabólica comum com Alzheimer, doenças cardiovasculares e diabetes.
    - “Efeito Warburg” (1920): células tumorais metabolizam glicose via glicólise, mesmo com oxigênio, favorecendo rápida multiplicação.
*   **Importância do Exercício de Força**
    - “O sedentarismo será o tabagismo do futuro”.

---

### Chunk 13/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.609

isco em pacientes com histórico familiar positivo.
## Diagnóstico Primário:
- Avaliação: Apresentação informativa sobre reposição hormonal, riscos de câncer de mama e manejo de condições relacionadas; não se aplica a diagnóstico individual. Enfatiza individualização do tratamento, desmistifica medos comuns sobre RH com base em evidências e discute manejo de nódulos, mamas densas e uso de hormônios como gestrinona e testosterona.
- Diagnóstico Suspeito: Nenhum no momento
## Plano:
- Prescrição: Inserir mais aqui
- Próximos Passos/Exames:
    - Nódulos novos na menopausa: investigar, geralmente com biópsia de fragmento.
    - Nódulos BI-RADS 3: acompanhamento semestral no primeiro ano e depois anual; RH não exige seguimento mais intenso.
    - Mamas densas: considerar exames adicionais (p.ex., ressonância magnética) conforme caso, não de rotina.

---

### Chunk 14/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.602

namnese detalhada abrangendo hábitos de vida, saúde intestinal, sintomas hormonais (dor mamária, TPM), história oncológica familiar (ambos os lados) e rede de apoio.
- [ ] Avaliar composição corporal com bioimpedância ou densitometria de corpo total para detectar sarcopenia e orientar intervenção.
- [ ] Solicitar e acompanhar marcadores inflamatórios/metabólicos para monitorar resposta a intervenções de estilo de vida, especialmente em pacientes com diagnóstico prévio de câncer.
- [ ] Utilizar calculadoras de risco (p. ex., Gail e ferramentas que integram PRS e fatores de estilo de vida) com cautela, reconhecendo superestimação em populações brasileiras, e empregá-las para motivar mudanças comportamentais.
- [ ] Considerar farmacogenética ao selecionar terapias, avaliando polimorfismos que alteram metabolização de drogas.
- [ ] Preparar materiais e plano para a próxima aula focada em dúvidas sobre reposição hormonal.

---

### Chunk 15/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.601

a a vácuo oferece uma alternativa quase 100% eficaz à cirurgia.**
- Nódulos benignos são comumente classificados como BI-RADS 3.
- O estudo ACRIM 666, que acompanhou pacientes de alto risco, revelou que a taxa de erro em que um nódulo BI-RADS 3 era, na verdade, câncer, foi de apenas 0,8%.
- A biópsia a vácuo demonstra uma eficácia comparável à da biópsia cirúrgica (padrão ouro), aproximando-se de 100%.
**Embora mulheres com mamas densas tenham um risco relativo de 4 a 6 vezes maior de desenvolver câncer de mama, o risco absoluto aumenta de forma modesta (de 10% para 10,6%), reforçando a importância do rastreio mamográfico a partir dos 50 anos.**
- Mulheres com mamas muito densas apresentam uma incidência de câncer de mama de 4 a 6 vezes maior em comparação com aquelas com mamas lipossubstituídas.
- Apesar do risco relativo elevado, o risco absoluto de uma paciente com mama densa desenvolver câncer ao longo da vida aumenta de uma base de 10% para apenas 10,6%.

---

### Chunk 16/30
**Article:** Mastologia I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.600

# Mastologia I

**Source:** https://web.plaud.ai/share/e07f1765255675478::YXdzOnVzLXdlc3QtMg

---

## Lecture

Data e Hora: 2025-11-21 03:11:05
Local: [Inserir Local]
Instrutor: Marina Ávila
## 📝 Resumo
A mastologista Marina Ávila apresenta uma visão integrativa da saúde mamária e do câncer de mama, criticando a abordagem tradicional centrada no diagnóstico precoce via mamografia, considerada insuficiente diante do aumento alarmante da incidência. A palestra enfatiza cuidar “da paciente por trás das mamas”, abordando fatores de risco frequentemente negligenciados como estilo de vida, obesidade, inflamação crônica e resistência à insulina. Defende a prevenção por meio da educação das pacientes para “dirigir melhor o carro” de suas vidas, com foco em alimentação, exercício de força, sono e manejo do estresse. Discute limitações e vieses do rastreamento mamográfico e antecipa um futuro com tecnologias de diagnóstico mais precisas e individualizadas.

---

### Chunk 17/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte VI (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.598

canismos epigenéticos subjacentes. Foi apresentada uma crítica contundente ao tratamento convencional, que mascara sintomas, em contraste com uma abordagem integrativa que busca tratar as causas metabólicas e inflamatórias.
## Conteúdo Abordado
### 1. Anticoncepcionais Hormonais e Risco de Câncer de Mama
- A visão funcional integrativa busca uma análise isenta sobre o uso de anticoncepcionais, avaliando a diferença entre risco relativo e risco absoluto.
- Um estudo da Universidade de Oxford e outro do New England (1,8 milhão de mulheres) apontam um aumento no risco relativo de câncer de mama (cerca de 20%) com o uso de contraceptivos hormonais orais.
- O risco persiste por cinco anos ou mais após a descontinuação do uso.
- O DIU hormonal (Mirena) também demonstrou um risco aumentado (1.4 vezes maior em um estudo dinamarquês), contrariando a ideia de que sua ação é apenas local.

---

### Chunk 18/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.597

 , gerando orientações como acompanhamento mamário trimestral ou responsabilização do paciente (“por sua conta e risco”).
* **Natureza proliferativa dos hormônios e sua relevância**
   - Hormônios femininos (estrógeno e progesterona) são proliferativos e isso é fisiologicamente benéfico (ossos, vasos, desenvolvimento mamário), não podendo esperar proliferação seletiva apenas onde “convém”.
   - Exemplifica com a menarca: aumento de hormônios circulantes promove desenvolvimento mamário; a mama é um órgão endócrino com função de produzir leite, regulada por um conjunto hormonal.
   - Outros hormônios influenciam a mama: insulina e IGF-1, entre outros; portanto, não apenas “hormônios femininos” são proliferativos na glândula mamária.
* **Tipos de estudos e interpretação crítica**
   - Há estudos observacionais e ensaios clínicos randomizados (ECR); é preciso pautar decisões pelo melhor nível de evidência.

---

### Chunk 19/30
**Article:** Trato Gastrointestinal V – Intestino Delgado I (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.597

secretina e a colecistoquinina.
2. Explicação aprofundada sobre os diferentes tipos de alergias e intolerâncias alimentares.
3. Discussão sobre a microbiota, probióticos, o equilíbrio do microbioma e a reação dos pacientes às fibras.
4. Análise aprofundada do sistema imunológico e do módulo de imunidade.
5. Explicação sobre os diferentes tipos de fibras.
6. Discussão detalhada sobre o TMAO (óxido de trimetilamina) e sua relevância clínica.
7. Continuação dos estímulos no intestino delgado.
## Conteúdo Abordado
### 1. A Abordagem Funcional Integrativa vs. Convencional
- A medicina funcional integrativa foca em organizar a vida do paciente, não apenas em prescrever medicamentos.
- Um caso clínico foi apresentado: uma paciente em reposição hormonal com histórico familiar de câncer de mama, cujos médicos anteriores prescreveram hormônios sem abordar o metabolismo hormonal ou o estilo de vida.

---

### Chunk 20/30
**Article:** Mastologia I (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.591

avaliar a eficácia percebida do rastreamento mamográfico.
- Um exemplo ilustra o viés: uma paciente diagnosticada aos 61 anos que falece aos 70 tem uma "sobrevida em 5 anos" de 100%, enquanto outra diagnosticada aos 67 que falece na mesma idade (70) tem uma sobrevida de 0%, mostrando como o diagnóstico precoce pode inflar artificialmente a estatística sem necessariamente alterar o desfecho final.
- Em um cenário hipotético, de 1000 pacientes com câncer progressivo, se 600 morressem, a taxa de sobrevida em cinco anos seria de 40%.

---

## Meeting Highlights

### Paradigma de Cuidado
Foco em medicina integrativa centrada na paciente para tratar causas e contexto além da imagem.  
- Trate a “dona das mamas” com plano multifatorial que prioriza fatores controláveis e suporte multidisciplinar.
- Converta conhecimento em mudança sustentada com acompanhamento coordenado e metas claras.

---

### Chunk 21/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** methods | **Similarity:** 0.590

rar exames adicionais (p.ex., ressonância magnética) conforme caso, não de rotina.
    - Atrofia vaginal: abordagem escalonada, iniciando com métodos não hormonais e, se necessário, terapia hormonal tópica.
- Plano de Tratamento de Acompanhamento:
    - RH deve ser individualizada, considerando particularidades de cada paciente, em vez de protocolo único.
    - Aconselhamento e educação para desmistificar medos e alinhar expectativas, especialmente sobre nódulos benignos e manejo de mamas densas.
    - RH não recomendada atualmente para pacientes pós-câncer de mama devido a ensaios clínicos que mostram aumento de recorrência.
    - Testosterona não deve ser indicada com objetivo de reduzir risco de câncer de mama, pois isso não é validado.
    - Enfatizar aconselhamento sobre estilo de vida e saúde geral para reduzir risco de câncer de mama.

---

### Chunk 22/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.589

sobre estilo de vida e saúde geral para reduzir risco de câncer de mama.

---

## Quantitative Data

### Narrativa Quantitativa
A discussão sobre reposição hormonal e câncer de mama é complexa, marcada por estudos históricos que geraram receio, como o de 2019. No entanto, a análise moderna, apoiada por ferramentas de diagnóstico precisas como a biópsia a vácuo (quase 100% eficaz) e a estratificação de risco (BI-RADS 3 com erro de apenas 0,8%), permite uma abordagem mais segura e individualizada, mesmo para pacientes com fatores de risco como mamas densas.
---
### Evidências Principais
**A avaliação de nódulos mamários benignos (BI-RADS 3) é altamente confiável, com um estudo de referência (ACRIM 666) mostrando um risco de erro diagnóstico de apenas 0,8%, e a biópsia a vácuo oferece uma alternativa quase 100% eficaz à cirurgia.**
- Nódulos benignos são comumente classificados como BI-RADS 3.

---

### Chunk 23/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.588

dos não hormonais.
   - Abordagem em etapas: iniciar com não hormonais; se falhar e houver queixas significativas, considerar tópico. Importante manejo emocional/psicológico, acolhimento de sintomas, e ausência de pressa, individualizando.
### 7. Princípios de acompanhamento e cuidado integral
* **Evitar sobrevigília desnecessária**
   - Reposição hormonal não implica necessidade de acompanhar mamas “de três em três meses” para segurança; tal prática pode aumentar ansiedade sem base em evidência.
* **Saúde integral reduz risco**
   - Tudo que se faz em prol da saúde integral da paciente tende a reduzir chances de câncer de mama; a mama deve ser considerada no contexto do corpo inteiro, não isoladamente.

---

### Chunk 24/30
**Article:** Mastologia I (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.586

orporar na prática clínica a orientação sobre exercícios de força, além de atividades aeróbicas, para prevenção e melhor prognóstico.
- [ ] 4. Ler artigo indicado sobre novas tecnologias e estratégias de diagnóstico do câncer de mama para atualização.
- [ ] 5. Usar analogias, como a do “carro” e do “acidente”, para explicar às pacientes a diferença entre redução de risco (prevenção) e diagnóstico precoce.

---

## Quantitative Data

### Narrativa Quantitativa
O cenário do câncer de mama no Brasil revela uma tendência alarmante de crescimento, com uma projeção de 74 mil novos casos anuais e um aumento futuro de 40% na incidência e 50% na mortalidade até 2040. Fatores metabólicos como hiperinsulinemia e sarcopenia agravam drasticamente este quadro, aumentando significativamente o risco de morte. A análise estatística, como a sobrevida em 5 anos, requer cautela, pois pode ser influenciada por vieses de diagnóstico precoce.

---

### Chunk 25/30
**Article:** Mastologia I (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.586

s metabolizam glicose via glicólise, mesmo com oxigênio, favorecendo rápida multiplicação.
*   **Importância do Exercício de Força**
    - “O sedentarismo será o tabagismo do futuro”.
    - Priorizar exercícios de força, não apenas caminhadas.
    - Meta-análise: sarcopenia associada a 44% mais mortes por todas as causas e 93% mais mortes por câncer.
    - Em câncer de mama, sarcopenia aumenta mortalidade em 41%.

## ❓ Perguntas
- [Inserir Pergunta/Dúvida]

## 📚 Tarefas
- [ ] 1. Estudar fatores de risco para câncer de mama além da genética: alimentação, microbiota, sono, estresse, obesidade e resistência à insulina.
- [ ] 2. Aprender a identificar sinais de resistência à insulina e inflamação crônica, inclusive em pacientes com peso normal.
- [ ] 3. Incorporar na prática clínica a orientação sobre exercícios de força, além de atividades aeróbicas, para prevenção e melhor prognóstico.
- [ ] 4.

---

### Chunk 26/30
**Article:** Mastologia I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.585

o, exercício de força, sono e manejo do estresse. Discute limitações e vieses do rastreamento mamográfico e antecipa um futuro com tecnologias de diagnóstico mais precisas e individualizadas.

## 🔖 Pontos de Conhecimento
### 1. Apresentação e Filosofia de Atendimento
*   **Formação e Trajetória da Palestrante**
    - Marina Ávila, mastologista.
    - Formação na Universidade Federal de Pernambuco, residência em ginecologia/obstetrícia e, posteriormente, em mastologia.
    - A residência em mastologia no Brasil foi a primeira do mundo; a especialidade se separou da ginecologia devido ao avanço do conhecimento sobre o câncer de mama.
*   **Abordagem Integrativa da Mastologia**
    - Incomodada com a abordagem de apenas diagnosticar e tratar.
    - Filosofia atual: cuidar “da paciente atrás das mamas, da dona das mamas”.
    - A medicina integrativa amplia as ferramentas para ajudar as pacientes, abordando fatores de risco geralmente ignorados.

### 2.

---

### Chunk 27/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.584

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

### Chunk 28/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.582

íveis a fatores genéticos de alta/média penetrância, enquanto 90% estão ligados a fatores não genéticos e potencialmente modificáveis. Ferramentas de cálculo de risco e critérios clínicos ajudam a identificar quem está em maior risco ao longo da vida (≥20%) ou no curto prazo (Gail 5 anos ≥1,7), orientando prevenção, rastreamento e decisões personalizadas.
---
### Evidências-Chave
**A maior parte dos casos de câncer de mama decorre de fatores não genéticos, mas um subconjunto relevante tem risco hereditário elevado que requer atenção específica.**
- Genética relacionada ao câncer de mama corresponde a cerca de 10% dos tumores diagnosticados, indicando a fração atribuída a fatores genéticos versus outros fatores.
- Os outros 90% dos casos não têm achados genéticos e são relacionados ao estilo de vida, enfatizando a influência de fatores modificáveis.

---

### Chunk 29/30
**Article:** Dieta Cetogênica - Parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.581

atamento metabólico do câncer
   - Terapias de pressão (contínuas): dieta cetogênica, cetonas exógenas, suplementos/fitoterápicos/drogas individualizadas, manejo do estresse emocional.
   - Terapias de pulso (intermitentes): inibição de glicose, inibição de glutamina, oxigenoterapia hiperbárica, entre outras.
   - Abordagem integrada e personalizada para maximizar o controle tumoral.
* Ensaio clínico randomizado (2021) em câncer de mama
   - 80 pacientes tratados com quimio; randomização para dois grupos; intervenção cetogênica/metabólica por 12 semanas; exames laboratoriais e de imagem no início e 12 semanas; cirurgia e reestadiamento para doença localmente avançada após quimio.
   - Resultados: redução de TNF-α, IGF-1, insulina; aumento de IL-10; redução significativa do tamanho tumoral no grupo cetogênico.

---

### Chunk 30/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte VI (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.581

# Abordagem Funcional Integrativa Aplicada a Cada Área - Parte VI

**Source:** https://web.plaud.ai/share/10021763827755332::YXdzOnVzLXdlc3QtMg

---

## Lecture

Data e Hora: 2025-11-17 14:39:53
Local: [Inserir Local]
Instrutores: [Inserir Nome], Vitória
## 📝 Resumo
Esta palestra aborda temas da ginecologia sob uma perspectiva funcional e integrativa, com foco crítico no uso de anticoncepcionais hormonais e na gestão da Síndrome dos Ovários Policísticos (SOP). Os instrutores analisam estudos que associam vários tipos de contraceptivos hormonais a um aumento no risco de câncer de mama, enfatizando a importância de interpretar corretamente os dados estatísticos (risco relativo vs. absoluto). A abordagem convencional para a SOP, baseada no uso de anticoncepcionais, é criticada por mascarar os sintomas sem tratar as causas subjacentes, como a resistência à insulina e a inflamação crônica.

---

