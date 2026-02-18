# ScoreItem: Mamas

**ID:** `019bf31d-2ef0-7ad3-9340-a962ffb32c58`
**FullName:** Mamas (Histórico de doenças - Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual): - Segmento torácico)

**Preparation Metadata:**
- Quality Grade: **GOOD**
- Total Chunks: 30 de 8 artigos
- Avg Similarity: 0.571

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-7ad3-9340-a962ffb32c58`.**

```json
{
  "score_item_id": "019bf31d-2ef0-7ad3-9340-a962ffb32c58",
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

**ScoreItem:** Mamas (Histórico de doenças - Histórico de saúde (Questionar ativamente sobre doenças/sintomas - passado ou atual): - Segmento torácico)

**30 chunks de 8 artigos (avg similarity: 0.571)**

### Chunk 1/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.734

dos sintomas gerais a investigar em consulta, como dor mamária e TPM intensa, que podem indicar predominância estrogênica.

## Objetivo:
Não há achados de exame físico de um paciente específico. A abordagem objetiva descrita inclui:
- Avaliação da composição corporal por bioimpedância ou densitometria de corpo total, identificando condições como sarcopenia em mulheres com peso normal.
- Solicitação de exames laboratoriais para avaliar marcadores inflamatórios e verificar se a paciente está metabolicamente doente ou em risco.

## Diagnóstico Primário:
- Avaliação: Discussão educacional sobre estratificação de risco para câncer de mama. 90% dos casos relacionam-se ao estilo de vida e 10% a fatores genéticos conhecidos. Mutações de alta penetrância (ex.: BRCA1, BRCA2, TP53) aumentam significativamente o risco, porém são raras. Epigenética e estilo de vida são cruciais e modificáveis.

---

### Chunk 2/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.648

luto, evitar alarmismo, e definir critérios para exames complementares (ressonância) apenas quando houver fatores de risco adicionais.
- [ ] 4. Atualizar materiais educativos para esclarecer que história familiar, por si só, não contraindica reposição; incorporar achados do Sister Study e WHI.
- [ ] 5. Estabelecer diretriz interna: não indicar reposição hormonal sistêmica em pacientes com histórico de câncer de mama; considerar terapias tópicas para atrofia vaginal após tentativa de métodos não hormonais, com suporte emocional.
- [ ] 6. Criar protocolo de uso criterioso de gestrinona em endometriose e mastalgia refratária, com consentimento informado sobre lacunas de evidência oncológica.
- [ ] 7. Definir critérios de indicação de testosterona por motivos não oncológicos, evitando prescrição para “redução de risco mamário” até que haja validação em guidelines.
- [ ] 8.

---

### Chunk 3/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.640

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
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** methods | **Similarity:** 0.624

; uso de medicamentos/suplementos; hábitos de vida (alimentação, exercício, álcool, tabaco); saúde intestinal; níveis de estresse/ansiedade; rede de apoio.
   - Atenção a sintomas de predominância estrogênica: dor mamária, TPM intensa; lacunas de cuidado em pacientes com alterações fibrocísticas e cistos, frequentemente não acolhidas apesar de sintomas de excesso de estrogênio circulante.
* **Avaliação de composição corporal**
   - Preferência por métodos além da balança: bioimpedância (utilizada pela instrutora) ou densitometria de corpo total.
   - Observação de sarcopenia em mulheres com peso normal e desconhecimento da condição; necessidade de intervenção.
* **Marcadores laboratoriais e metabolicidade**
   - Solicitar marcadores inflamatórios para identificar adoecimento metabólico ou risco, e para monitorar evolução após intervenções.

---

### Chunk 5/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.618

liando polimorfismos que alteram metabolização de drogas.
- [ ] Preparar materiais e plano para a próxima aula focada em dúvidas sobre reposição hormonal.

---

## SOAP

Data e Hora: 2025-11-21 03:11:20
Paciente:
Diagnóstico:

## Histórico de Diagnóstico:
1. Histórico Médico: A transcrição discute fatores de risco para câncer de mama, destacando a importância de uma anamnese detalhada, incluindo patologias mamárias anteriores (biópsias), história oncológica familiar, hábitos de vida (alimentação, exercício, álcool, cigarro), saúde intestinal, níveis de estresse e ansiedade, e sintomas de predominância estrogênica (dor mamária, TPM intensa).
2. Histórico de Medicação: Inserir mais aqui

## Subjetivo:
Não há sintomas de um paciente específico. São mencionados sintomas gerais a investigar em consulta, como dor mamária e TPM intensa, que podem indicar predominância estrogênica.

---

### Chunk 6/30
**Article:** Mastologia I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.596

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

### Chunk 7/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.595

namnese detalhada abrangendo hábitos de vida, saúde intestinal, sintomas hormonais (dor mamária, TPM), história oncológica familiar (ambos os lados) e rede de apoio.
- [ ] Avaliar composição corporal com bioimpedância ou densitometria de corpo total para detectar sarcopenia e orientar intervenção.
- [ ] Solicitar e acompanhar marcadores inflamatórios/metabólicos para monitorar resposta a intervenções de estilo de vida, especialmente em pacientes com diagnóstico prévio de câncer.
- [ ] Utilizar calculadoras de risco (p. ex., Gail e ferramentas que integram PRS e fatores de estilo de vida) com cautela, reconhecendo superestimação em populações brasileiras, e empregá-las para motivar mudanças comportamentais.
- [ ] Considerar farmacogenética ao selecionar terapias, avaliando polimorfismos que alteram metabolização de drogas.
- [ ] Preparar materiais e plano para a próxima aula focada em dúvidas sobre reposição hormonal.

---

### Chunk 8/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.592

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

### Chunk 9/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.590

urgias profiláticas) e a importância do aconselhamento genético. Apresenta calculadoras de risco (Gail como a mais conhecida), observa superestimação fora de populações de origem e sugere uso pedagógico para impulsionar mudanças de estilo de vida. Destaca anamnese detalhada, avaliação de composição corporal e marcadores metabólicos/inflamatórios como base prática de estratificação. Conclui que genética não é destino, introduz epigenética como fator modificável e informa que dúvidas sobre reposição hormonal serão abordadas na próxima aula. Data de criação: 2025-11-21.
## 🔖 Pontos de Conhecimento
### 1. Genética e câncer de mama
* **Proporção de câncer de mama ligado à genética**
   - Aproximadamente 10% dos tumores de mama diagnosticados relacionam-se a mutações genéticas conhecidas; os demais 90% não apresentam achados genéticos e são majoritariamente atribuídos ao estilo de vida.

---

### Chunk 10/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.578

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

### Chunk 11/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.572

íveis a fatores genéticos de alta/média penetrância, enquanto 90% estão ligados a fatores não genéticos e potencialmente modificáveis. Ferramentas de cálculo de risco e critérios clínicos ajudam a identificar quem está em maior risco ao longo da vida (≥20%) ou no curto prazo (Gail 5 anos ≥1,7), orientando prevenção, rastreamento e decisões personalizadas.
---
### Evidências-Chave
**A maior parte dos casos de câncer de mama decorre de fatores não genéticos, mas um subconjunto relevante tem risco hereditário elevado que requer atenção específica.**
- Genética relacionada ao câncer de mama corresponde a cerca de 10% dos tumores diagnosticados, indicando a fração atribuída a fatores genéticos versus outros fatores.
- Os outros 90% dos casos não têm achados genéticos e são relacionados ao estilo de vida, enfatizando a influência de fatores modificáveis.

---

### Chunk 12/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.567

100%, resolvendo muitas questões e muitas vezes removendo a lesão sem necessidade de sala cirúrgica.
### 3. Densidade mamária
* **Definição e critérios**
   - Densidade mamária é critério mamográfico; não pode ser diagnosticada sem a primeira mamografia.
   - Mamas densas em pacientes jovens são esperadas; critério torna-se mais relevante em rastreio, geralmente a partir dos 50 anos.
* **Risco relativo versus absoluto**
   - Estudos associam densidade aumentada a maior incidência de câncer (com razão de risco frequentemente citada como 4 a 6 vezes ao comparar mama muito densa com lipossubstituída), mas o risco absoluto é baixo: exemplo citado, de 10% para 10,6%.
   - O problema maior é a dificuldade diagnóstica em mamas muito densas; exames complementares, como ressonância magnética, podem ser úteis em pacientes com mamas densas e fatores de risco adicionais.

---

### Chunk 13/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** methods | **Similarity:** 0.566

rar exames adicionais (p.ex., ressonância magnética) conforme caso, não de rotina.
    - Atrofia vaginal: abordagem escalonada, iniciando com métodos não hormonais e, se necessário, terapia hormonal tópica.
- Plano de Tratamento de Acompanhamento:
    - RH deve ser individualizada, considerando particularidades de cada paciente, em vez de protocolo único.
    - Aconselhamento e educação para desmistificar medos e alinhar expectativas, especialmente sobre nódulos benignos e manejo de mamas densas.
    - RH não recomendada atualmente para pacientes pós-câncer de mama devido a ensaios clínicos que mostram aumento de recorrência.
    - Testosterona não deve ser indicada com objetivo de reduzir risco de câncer de mama, pois isso não é validado.
    - Enfatizar aconselhamento sobre estilo de vida e saúde geral para reduzir risco de câncer de mama.

---

### Chunk 14/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.566

RCA2, TP53) aumentam significativamente o risco, porém são raras. Epigenética e estilo de vida são cruciais e modificáveis. A avaliação individualizada requer anamnese detalhada, composição corporal e marcadores laboratoriais.
- Diagnóstico Suspeito: Nenhum no momento

## Plano:
- Prescrição: Inserir mais aqui
- Próximos Passos:
  - Suspeitar de mutação genética e considerar teste genético em casos de múltiplos cânceres na família (mama, ovário, pâncreas, próstata), câncer de mama triplo negativo, câncer de mama em idade jovem (<45 anos), câncer de mama em homem, ou descendência judaica Ashkenazi.
  - Encaminhar para aconselhamento genético antes do teste em pacientes com alta suspeita de mutação.
  - Utilizar calculadoras de risco (ex.: Gail, Tyrer-Cuzick) para conscientizar sobre mudanças no estilo de vida, reconhecendo limitações na população brasileira.

---

### Chunk 15/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.563

clínicos randomizados (p.ex., WHI) não mostraram o mesmo aumento de risco de câncer de mama observado em estudos observacionais.
- Progestágeno (ex.: medroxiprogesterona) foi associado a aumento do risco mamário, enquanto progesterona micronizada não demonstrou o mesmo efeito.
- Nódulos BI-RADS 3 têm risco muito baixo de malignidade (~0,8% no ACRIN 666); acompanhamento semestral no primeiro ano e depois regular é o padrão.
- Densidade mamária é critério mamográfico com ligeiro aumento de incidência de câncer de mama, mas risco absoluto baixo.
- Biópsia a vácuo é eficaz para diagnosticar nódulos, com precisão comparável à biópsia cirúrgica.
- Histórico familiar aumenta risco pessoal, mas estudos (Sister Study, WHI) indicam que RH não aumenta adicionalmente esse risco em pacientes com histórico familiar positivo.

---

### Chunk 16/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.560

 , gerando orientações como acompanhamento mamário trimestral ou responsabilização do paciente (“por sua conta e risco”).
* **Natureza proliferativa dos hormônios e sua relevância**
   - Hormônios femininos (estrógeno e progesterona) são proliferativos e isso é fisiologicamente benéfico (ossos, vasos, desenvolvimento mamário), não podendo esperar proliferação seletiva apenas onde “convém”.
   - Exemplifica com a menarca: aumento de hormônios circulantes promove desenvolvimento mamário; a mama é um órgão endócrino com função de produzir leite, regulada por um conjunto hormonal.
   - Outros hormônios influenciam a mama: insulina e IGF-1, entre outros; portanto, não apenas “hormônios femininos” são proliferativos na glândula mamária.
* **Tipos de estudos e interpretação crítica**
   - Há estudos observacionais e ensaios clínicos randomizados (ECR); é preciso pautar decisões pelo melhor nível de evidência.

---

### Chunk 17/30
**Article:** MFI - Reposição Hormonal - AULA 07 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.550

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

### Chunk 18/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.549

 ncia magnética, podem ser úteis em pacientes com mamas densas e fatores de risco adicionais.
* **Conduta e reposição hormonal**
   - Não há indicação formal de adicionar exames em toda paciente com mama densa; evitar exageros.
   - Uso de reposição hormonal não justifica acompanhamento mais intenso apenas por si; não se deve aumentar vigilância por conta da reposição se a intenção é tranquilizar a paciente.
### 4. História familiar e reposição
* **Impacto da história familiar**
   - História familiar (mãe/irmã com câncer de mama) aumenta discretamente o risco pessoal, influenciada por fatores genéticos, epigenéticos e ambientais compartilhados.
* **Evidência sobre reposição em quem tem história familiar**
   - Estudos (Sister Study, WHI) incluíram pacientes com história familiar positiva; reposição hormonal, conforme realizada nesses estudos, não aumentou adicionalmente o risco em relação ao já conferido pela história familiar.

---

### Chunk 19/30
**Article:** Mastologia I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.548

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

### Chunk 20/30
**Article:** Mastologia I (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.543

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

### Chunk 21/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.540

 gicos, evitando prescrição para “redução de risco mamário” até que haja validação em guidelines.
- [ ] 8. Treinar equipe para comunicação que desestimule “acompanhamento trimestral” de mamas por reposição hormonal sem indicação, reduzindo ansiedade e sobre-exames.

---

## SOAP

Data e Hora: 2025-11-21 03:31:38
Paciente: 
Diagnóstico:
## Histórico de Diagnóstico:
1. Histórico Médico: A transcrição é uma apresentação geral sobre reposição hormonal e não detalha o histórico médico de um paciente específico.
2. Histórico de Medicação: Inserir mais aqui
## Subjetivo:
A transcrição é uma apresentação médica e não contém queixas subjetivas de um paciente. Os tópicos discutidos incluem dúvidas frequentes sobre reposição hormonal:
- Uso de RH em pacientes com nódulos mamários benignos (BI-RADS 3 ou fibroadenomas).
- Preocupação com surgimento de novos nódulos durante RH.
- RH em pacientes com mamas densas.

---

### Chunk 22/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.536

a a vácuo oferece uma alternativa quase 100% eficaz à cirurgia.**
- Nódulos benignos são comumente classificados como BI-RADS 3.
- O estudo ACRIM 666, que acompanhou pacientes de alto risco, revelou que a taxa de erro em que um nódulo BI-RADS 3 era, na verdade, câncer, foi de apenas 0,8%.
- A biópsia a vácuo demonstra uma eficácia comparável à da biópsia cirúrgica (padrão ouro), aproximando-se de 100%.
**Embora mulheres com mamas densas tenham um risco relativo de 4 a 6 vezes maior de desenvolver câncer de mama, o risco absoluto aumenta de forma modesta (de 10% para 10,6%), reforçando a importância do rastreio mamográfico a partir dos 50 anos.**
- Mulheres com mamas muito densas apresentam uma incidência de câncer de mama de 4 a 6 vezes maior em comparação com aquelas com mamas lipossubstituídas.
- Apesar do risco relativo elevado, o risco absoluto de uma paciente com mama densa desenvolver câncer ao longo da vida aumenta de uma base de 10% para apenas 10,6%.

---

### Chunk 23/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.536

s com história familiar positiva; reposição hormonal, conforme realizada nesses estudos, não aumentou adicionalmente o risco em relação ao já conferido pela história familiar.
   - Exemplo: duas irmãs com mãe com câncer; uma usa reposição e a outra não; chance de desenvolver câncer é semelhante, reforçando que história familiar, por si, não contraindica automaticamente a reposição.
### 5. Fármacos específicos: gestrinona e testosterona
* **Gestrinona**
   - Uso antigo; poucos estudos sobre mama. Ensaio clínico comparando gestrinona e danazol mostrou melhora de mastalgia, com redução de estradiol e progesterona durante o tratamento.
   - Não há estudos robustos com desfechos oncológicos de longo prazo; especula-se possível proteção futura com base em mecanismos e melhora de mastalgia, mas não é possível afirmar.

---

### Chunk 24/30
**Article:** Global trends of cancer: The role of diet, lifestyle, and environmental factors (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.535

uetomammographiccompression.AcademicRadiol.2014;21(2):151–61.https://doi.org/10.1016/j.acra.2013.10.00925.SoroushA,FarshchianN,KomasiS,IzadiN,AmirifardN,ShahmohammadiA.TheroleoforalcontraceptivepillsonincreasedriskofbreastcancerinIranianpopulations:ametaanalysis.JCancerPrev.2016;21(4):294–301.https://doi.org/10.15430/JCP.2016.21.4.29426.RiosS,ChenAC.Wearingatightbraformanyhoursadayisassociatedwithincreasedriskofbreastcancer.JOncolResTreat.2016;1:1–5.https://doi.org/10.13140/RG.2.2.10742.1952527.FuY,GuoF,ChenH,LinY,FuX,ZhangH,etal.Coreneedlebiopsypromoteslungmetastasisofbreastcancer:anexperi-mentalstudy.MolClinOncol.2018;10(2):253–60.https://doi.org/10.3892/mco.2018.178428.BrunicardiFC,AndersenDK.Schwartzprinciplesofsurgery.10thed.USA:McGrawHillProfessional;2014.29.AbadiAT,IerardiE,LeeYY.WhydowestillhaveHelicobacterpyloriinourstomachs.MalaysJMedSci.2015;22(5):70–5.30.AlexanderSM,RetnakumarRJ,ChouhanD,DeviTNB,DharmaseelanS,DevadasK,etal.Helicobacterpyloriinhumanstomach:theinco

---

### Chunk 25/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.534

dos não hormonais.
   - Abordagem em etapas: iniciar com não hormonais; se falhar e houver queixas significativas, considerar tópico. Importante manejo emocional/psicológico, acolhimento de sintomas, e ausência de pressa, individualizando.
### 7. Princípios de acompanhamento e cuidado integral
* **Evitar sobrevigília desnecessária**
   - Reposição hormonal não implica necessidade de acompanhar mamas “de três em três meses” para segurança; tal prática pode aumentar ansiedade sem base em evidência.
* **Saúde integral reduz risco**
   - Tudo que se faz em prol da saúde integral da paciente tende a reduzir chances de câncer de mama; a mama deve ser considerada no contexto do corpo inteiro, não isoladamente.

---

### Chunk 26/30
**Article:** Fisiologia Endócrina Feminina (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.533

incluindo:
  - Fogachos (calores) e sudoreses noturnas; distúrbios do sono/insônia.
  - Ansiedade, depressão, baixa disposição/energia; redução de memória e vitalidade.
  - Dor articular; olho seco; ressecamento vaginal; dispareunia; baixa libido.
  - Alterações menstruais (polimenorreia com ciclos encurtados; períodos de amenorreia).
  - Síndrome geniturinária da pós-menopausa; sintomas urogenitais (incontinência, risco de prolapso).
  - Sarcopenia/perda de massa e tônus muscular.
  - Alterações cutâneas (queda de colágeno; pele mais flácida).
  - Mudanças metabólicas (↑ triglicérides, colesterol, leptina).
- Discussão sobre declínio da fertilidade com a idade, janela ótima de maternidade entre 20–30 anos e postergação da maternidade.
## Objetivo:
- Não há exame físico, laboratoriais ou de imagem de uma paciente específica.

---

### Chunk 27/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.532

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

### Chunk 28/30
**Article:** Ritmo Circadiano Eixo HPA - Parte X (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.531

udo sobre Trabalho Noturno e Câncer de Mama**
    - Meta-análise de 61 estudos (≈115 mil mulheres): trabalho em regime noturno aumenta o risco de câncer de mama em 32% na população geral.
    - Em enfermeiras, o risco sobe a 58%, possivelmente por alto consumo de café, alimentação inadequada (pizza, hambúrguer, doces) e estresse elevado do ambiente noturno.
*   **Higiene do Sono e Rotinas Matinais**
    - Orientação de higiene do sono é fundamental para todos os pacientes, mesmo sem queixas, pois muitos não percebem a má qualidade do descanso.
    - Evitar eletrônicos perto da cama à noite (celulares — especialmente carregando — e relógios eletrônicos).
    - Exposição à luz natural logo ao acordar é essencial para regular o ritmo circadiano, pois as células são fotossensíveis.
    - Rotina matinal sugerida: abrir a janela para luz natural, orar/conectar-se com uma força maior, agradecer e pedir por um dia iluminado antes de olhar o celular.
### 2.

---

### Chunk 29/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.527

com mamas lipossubstituídas.
- Apesar do risco relativo elevado, o risco absoluto de uma paciente com mama densa desenvolver câncer ao longo da vida aumenta de uma base de 10% para apenas 10,6%.
- O rastreio mamográfico é geralmente iniciado a partir dos 50 anos de idade.
**A percepção de risco da reposição hormonal foi moldada por estudos observacionais, como um de 2019, mas avanços nos últimos 20 anos permitem um acompanhamento mais seguro, como a monitorização da mama a cada três meses.**
- Um estudo observacional publicado em 2019 mostrou um aumento na incidência de câncer de mama, o que gerou receio entre os médicos.
- O material complementar deste estudo, com 50 páginas, continha um resumo de ensaios clínicos randomizados que ajudavam a contextualizar os achados.
- Nos últimos 20 anos, surgiram novos estudos que melhoraram o entendimento sobre a reposição hormonal.

---

### Chunk 30/30
**Article:** Terapia de Reposição Hormonal Feminina III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.527

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

