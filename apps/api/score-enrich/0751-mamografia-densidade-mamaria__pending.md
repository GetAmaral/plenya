# ScoreItem: Mamografia - Densidade Mamária

**ID:** `c77cedd3-2800-78bc-b316-71e29954eedd`
**FullName:** Mamografia - Densidade Mamária (Exames - Imagem)
**Unit:** Categoria

**Preparation Metadata:**
- Quality Grade: **EXCELLENT**
- Total Chunks: 30 de 6 artigos
- Avg Similarity: 0.625

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `c77cedd3-2800-78bc-b316-71e29954eedd`.**

```json
{
  "score_item_id": "c77cedd3-2800-78bc-b316-71e29954eedd",
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

**ScoreItem:** Mamografia - Densidade Mamária (Exames - Imagem)
**Unidade:** Categoria

**30 chunks de 6 artigos (avg similarity: 0.625)**

### Chunk 1/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.714

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
**Section:** other | **Similarity:** 0.710

100%, resolvendo muitas questões e muitas vezes removendo a lesão sem necessidade de sala cirúrgica.
### 3. Densidade mamária
* **Definição e critérios**
   - Densidade mamária é critério mamográfico; não pode ser diagnosticada sem a primeira mamografia.
   - Mamas densas em pacientes jovens são esperadas; critério torna-se mais relevante em rastreio, geralmente a partir dos 50 anos.
* **Risco relativo versus absoluto**
   - Estudos associam densidade aumentada a maior incidência de câncer (com razão de risco frequentemente citada como 4 a 6 vezes ao comparar mama muito densa com lipossubstituída), mas o risco absoluto é baixo: exemplo citado, de 10% para 10,6%.
   - O problema maior é a dificuldade diagnóstica em mamas muito densas; exames complementares, como ressonância magnética, podem ser úteis em pacientes com mamas densas e fatores de risco adicionais.

---

### Chunk 3/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.701

com mamas lipossubstituídas.
- Apesar do risco relativo elevado, o risco absoluto de uma paciente com mama densa desenvolver câncer ao longo da vida aumenta de uma base de 10% para apenas 10,6%.
- O rastreio mamográfico é geralmente iniciado a partir dos 50 anos de idade.
**A percepção de risco da reposição hormonal foi moldada por estudos observacionais, como um de 2019, mas avanços nos últimos 20 anos permitem um acompanhamento mais seguro, como a monitorização da mama a cada três meses.**
- Um estudo observacional publicado em 2019 mostrou um aumento na incidência de câncer de mama, o que gerou receio entre os médicos.
- O material complementar deste estudo, com 50 páginas, continha um resumo de ensaios clínicos randomizados que ajudavam a contextualizar os achados.
- Nos últimos 20 anos, surgiram novos estudos que melhoraram o entendimento sobre a reposição hormonal.

---

### Chunk 4/30
**Article:** Breast Density and Risk of Breast Cancer: Understanding the BI-RADS Classification (2021)
**Journal:** JAMA Oncology
**Section:** abstract | **Similarity:** 0.690

Comprehensive review of breast density classification using BI-RADS categories (A-D) and association with breast cancer risk. Women with dense breasts (categories C and D) have 1.5-2 times higher risk of developing breast cancer compared to women with fatty breasts.

---

### Chunk 5/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.688

a a vácuo oferece uma alternativa quase 100% eficaz à cirurgia.**
- Nódulos benignos são comumente classificados como BI-RADS 3.
- O estudo ACRIM 666, que acompanhou pacientes de alto risco, revelou que a taxa de erro em que um nódulo BI-RADS 3 era, na verdade, câncer, foi de apenas 0,8%.
- A biópsia a vácuo demonstra uma eficácia comparável à da biópsia cirúrgica (padrão ouro), aproximando-se de 100%.
**Embora mulheres com mamas densas tenham um risco relativo de 4 a 6 vezes maior de desenvolver câncer de mama, o risco absoluto aumenta de forma modesta (de 10% para 10,6%), reforçando a importância do rastreio mamográfico a partir dos 50 anos.**
- Mulheres com mamas muito densas apresentam uma incidência de câncer de mama de 4 a 6 vezes maior em comparação com aquelas com mamas lipossubstituídas.
- Apesar do risco relativo elevado, o risco absoluto de uma paciente com mama densa desenvolver câncer ao longo da vida aumenta de uma base de 10% para apenas 10,6%.

---

### Chunk 6/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.671

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

### Chunk 7/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.664

luto, evitar alarmismo, e definir critérios para exames complementares (ressonância) apenas quando houver fatores de risco adicionais.
- [ ] 4. Atualizar materiais educativos para esclarecer que história familiar, por si só, não contraindica reposição; incorporar achados do Sister Study e WHI.
- [ ] 5. Estabelecer diretriz interna: não indicar reposição hormonal sistêmica em pacientes com histórico de câncer de mama; considerar terapias tópicas para atrofia vaginal após tentativa de métodos não hormonais, com suporte emocional.
- [ ] 6. Criar protocolo de uso criterioso de gestrinona em endometriose e mastalgia refratária, com consentimento informado sobre lacunas de evidência oncológica.
- [ ] 7. Definir critérios de indicação de testosterona por motivos não oncológicos, evitando prescrição para “redução de risco mamário” até que haja validação em guidelines.
- [ ] 8.

---

### Chunk 8/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.663

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
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.654

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

### Chunk 10/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.644

clínicos randomizados (p.ex., WHI) não mostraram o mesmo aumento de risco de câncer de mama observado em estudos observacionais.
- Progestágeno (ex.: medroxiprogesterona) foi associado a aumento do risco mamário, enquanto progesterona micronizada não demonstrou o mesmo efeito.
- Nódulos BI-RADS 3 têm risco muito baixo de malignidade (~0,8% no ACRIN 666); acompanhamento semestral no primeiro ano e depois regular é o padrão.
- Densidade mamária é critério mamográfico com ligeiro aumento de incidência de câncer de mama, mas risco absoluto baixo.
- Biópsia a vácuo é eficaz para diagnosticar nódulos, com precisão comparável à biópsia cirúrgica.
- Histórico familiar aumenta risco pessoal, mas estudos (Sister Study, WHI) indicam que RH não aumenta adicionalmente esse risco em pacientes com histórico familiar positivo.

---

### Chunk 11/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.637

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

### Chunk 12/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.636

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

### Chunk 13/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.626

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

### Chunk 14/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.622

urgias profiláticas) e a importância do aconselhamento genético. Apresenta calculadoras de risco (Gail como a mais conhecida), observa superestimação fora de populações de origem e sugere uso pedagógico para impulsionar mudanças de estilo de vida. Destaca anamnese detalhada, avaliação de composição corporal e marcadores metabólicos/inflamatórios como base prática de estratificação. Conclui que genética não é destino, introduz epigenética como fator modificável e informa que dúvidas sobre reposição hormonal serão abordadas na próxima aula. Data de criação: 2025-11-21.
## 🔖 Pontos de Conhecimento
### 1. Genética e câncer de mama
* **Proporção de câncer de mama ligado à genética**
   - Aproximadamente 10% dos tumores de mama diagnosticados relacionam-se a mutações genéticas conhecidas; os demais 90% não apresentam achados genéticos e são majoritariamente atribuídos ao estilo de vida.

---

### Chunk 15/30
**Article:** Clinical Implications of Breast Density Assessment in Mammographic Screening (2020)
**Journal:** Radiology
**Section:** abstract | **Similarity:** 0.618

Evidence-based guidelines for clinical management of women with dense breasts. Discusses screening strategies, supplemental imaging modalities, and patient communication strategies.

---

### Chunk 16/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.613

íveis a fatores genéticos de alta/média penetrância, enquanto 90% estão ligados a fatores não genéticos e potencialmente modificáveis. Ferramentas de cálculo de risco e critérios clínicos ajudam a identificar quem está em maior risco ao longo da vida (≥20%) ou no curto prazo (Gail 5 anos ≥1,7), orientando prevenção, rastreamento e decisões personalizadas.
---
### Evidências-Chave
**A maior parte dos casos de câncer de mama decorre de fatores não genéticos, mas um subconjunto relevante tem risco hereditário elevado que requer atenção específica.**
- Genética relacionada ao câncer de mama corresponde a cerca de 10% dos tumores diagnosticados, indicando a fração atribuída a fatores genéticos versus outros fatores.
- Os outros 90% dos casos não têm achados genéticos e são relacionados ao estilo de vida, enfatizando a influência de fatores modificáveis.

---

### Chunk 17/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.610

# Mastologia III

**Source:** https://web.plaud.ai/share/e6aa1765255687340::YXdzOnVzLXdlc3QtMg

---

## Lecture

> Data e Hora: 2025-11-21 03:31:38
> Local: [Inserir Local]
> Instrutora: [Inserir Nome]
## 📝 Resumo
A palestrante discute dúvidas frequentes sobre reposição hormonal feminina, especialmente em relação ao risco de câncer de mama, contextualizando a evolução das evidências científicas nos últimos 20 anos. Ela explica a natureza proliferativa dos hormônios, diferencia progesterona de progestágenos, aborda estudos como o WHI e o grande observacional de 2019, e esclarece condutas frente a nódulos benignos (BI-RADS 3), mamas densas, história familiar, uso de gestrinona, testosterona e terapias tópicas pós-câncer. Enfatiza a individualização da terapia, a necessidade de não sobrevigiar injustificadamente pacientes em reposição, e a importância de cuidar da saúde integral para reduzir riscos. Data de criação do conteúdo: 2025-11-21.

---

### Chunk 18/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.607

eres jovens em idade reprodutiva sofrem complicações do tratamento; uma paciente relatou “pagar o preço da cura”, destacando a necessidade urgente de soluções para qualidade de vida pós-câncer.
## ❓ Perguntas
- [Inserir Pergunta/Confusão]
## 📚 Tarefas
- [ ] 1. Revisar protocolos internos para reposição hormonal, diferenciando claramente progestágenos sintéticos de progesterona micronizada e ajustando práticas conforme evidências de ECRs.
- [ ] 2. Implementar fluxos de manejo de nódulos BI-RADS 3: investigação adequada de nódulo novo na menopausa, seguimento semestral no primeiro ano, redução após, e uso de biópsia a vácuo quando indicado.
- [ ] 3. Padronizar comunicação aos pacientes sobre densidade mamária: explicar diferença entre risco relativo e absoluto, evitar alarmismo, e definir critérios para exames complementares (ressonância) apenas quando houver fatores de risco adicionais.
- [ ] 4.

---

### Chunk 19/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** methods | **Similarity:** 0.605

rar exames adicionais (p.ex., ressonância magnética) conforme caso, não de rotina.
    - Atrofia vaginal: abordagem escalonada, iniciando com métodos não hormonais e, se necessário, terapia hormonal tópica.
- Plano de Tratamento de Acompanhamento:
    - RH deve ser individualizada, considerando particularidades de cada paciente, em vez de protocolo único.
    - Aconselhamento e educação para desmistificar medos e alinhar expectativas, especialmente sobre nódulos benignos e manejo de mamas densas.
    - RH não recomendada atualmente para pacientes pós-câncer de mama devido a ensaios clínicos que mostram aumento de recorrência.
    - Testosterona não deve ser indicada com objetivo de reduzir risco de câncer de mama, pois isso não é validado.
    - Enfatizar aconselhamento sobre estilo de vida e saúde geral para reduzir risco de câncer de mama.

---

### Chunk 20/30
**Article:** MFI - Reposição Hormonal - AULA 07 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.601

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

### Chunk 21/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.592

RCA2, TP53) aumentam significativamente o risco, porém são raras. Epigenética e estilo de vida são cruciais e modificáveis. A avaliação individualizada requer anamnese detalhada, composição corporal e marcadores laboratoriais.
- Diagnóstico Suspeito: Nenhum no momento

## Plano:
- Prescrição: Inserir mais aqui
- Próximos Passos:
  - Suspeitar de mutação genética e considerar teste genético em casos de múltiplos cânceres na família (mama, ovário, pâncreas, próstata), câncer de mama triplo negativo, câncer de mama em idade jovem (<45 anos), câncer de mama em homem, ou descendência judaica Ashkenazi.
  - Encaminhar para aconselhamento genético antes do teste em pacientes com alta suspeita de mutação.
  - Utilizar calculadoras de risco (ex.: Gail, Tyrer-Cuzick) para conscientizar sobre mudanças no estilo de vida, reconhecendo limitações na população brasileira.

---

### Chunk 22/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.583

# 2. Manejo de nódulos mamários benignos e BI-RADS 3
* **Diferenciação de nódulos**
   - Nódulo antigo (estável, já acompanhado/biopsiado): classificado como BI-RADS 2; não requer intervenção.
   - Nódulo novo na menopausa: deve ser esclarecido, não por causa da reposição em si, mas pelo fator idade (envelhecimento é um grande risco). Em pacientes com tendência prévia a nódulos (p. ex., fibroadenoma), reposição pode manter o surgimento de novos nódulos, devendo a paciente ser previamente orientada para evitar susto.
* **ACRIN 6666 e risco de BI-RADS 3**
   - Estudo ACRIN 6666 acompanhou pacientes de alto risco (lesões proliferativas, história familiar em mãe/irmã e mamas densas). Em 6 meses, apenas um caso evoluiu; estimou-se 0,8% de câncer não diagnosticado entre BI-RADS 3, menor que o esperado.

---

### Chunk 23/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** methods | **Similarity:** 0.581

; uso de medicamentos/suplementos; hábitos de vida (alimentação, exercício, álcool, tabaco); saúde intestinal; níveis de estresse/ansiedade; rede de apoio.
   - Atenção a sintomas de predominância estrogênica: dor mamária, TPM intensa; lacunas de cuidado em pacientes com alterações fibrocísticas e cistos, frequentemente não acolhidas apesar de sintomas de excesso de estrogênio circulante.
* **Avaliação de composição corporal**
   - Preferência por métodos além da balança: bioimpedância (utilizada pela instrutora) ou densitometria de corpo total.
   - Observação de sarcopenia em mulheres com peso normal e desconhecimento da condição; necessidade de intervenção.
* **Marcadores laboratoriais e metabolicidade**
   - Solicitar marcadores inflamatórios para identificar adoecimento metabólico ou risco, e para monitorar evolução após intervenções.

---

### Chunk 24/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.580

 , gerando orientações como acompanhamento mamário trimestral ou responsabilização do paciente (“por sua conta e risco”).
* **Natureza proliferativa dos hormônios e sua relevância**
   - Hormônios femininos (estrógeno e progesterona) são proliferativos e isso é fisiologicamente benéfico (ossos, vasos, desenvolvimento mamário), não podendo esperar proliferação seletiva apenas onde “convém”.
   - Exemplifica com a menarca: aumento de hormônios circulantes promove desenvolvimento mamário; a mama é um órgão endócrino com função de produzir leite, regulada por um conjunto hormonal.
   - Outros hormônios influenciam a mama: insulina e IGF-1, entre outros; portanto, não apenas “hormônios femininos” são proliferativos na glândula mamária.
* **Tipos de estudos e interpretação crítica**
   - Há estudos observacionais e ensaios clínicos randomizados (ECR); é preciso pautar decisões pelo melhor nível de evidência.

---

### Chunk 25/30
**Article:** Mastologia I (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.580

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

### Chunk 26/30
**Article:** Mastologia I (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.579

orporar na prática clínica a orientação sobre exercícios de força, além de atividades aeróbicas, para prevenção e melhor prognóstico.
- [ ] 4. Ler artigo indicado sobre novas tecnologias e estratégias de diagnóstico do câncer de mama para atualização.
- [ ] 5. Usar analogias, como a do “carro” e do “acidente”, para explicar às pacientes a diferença entre redução de risco (prevenção) e diagnóstico precoce.

---

## Quantitative Data

### Narrativa Quantitativa
O cenário do câncer de mama no Brasil revela uma tendência alarmante de crescimento, com uma projeção de 74 mil novos casos anuais e um aumento futuro de 40% na incidência e 50% na mortalidade até 2040. Fatores metabólicos como hiperinsulinemia e sarcopenia agravam drasticamente este quadro, aumentando significativamente o risco de morte. A análise estatística, como a sobrevida em 5 anos, requer cautela, pois pode ser influenciada por vieses de diagnóstico precoce.

---

### Chunk 27/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.574

é maior do que 20%, critério de classificação para manejo intensificado.
**Calculadoras e critérios clínicos complementam a genética para identificar risco elevado e orientar prevenção.**
- Na calculadora de Gail, considera-se risco alto quando em cinco anos o risco é de 1,7, limiar utilizado para motivar adesão à mudança de estilo de vida e estratificação de risco.
- Câncer de mama em idade jovem normalmente se estipula abaixo dos 45 anos, critério de suspeita para indicação de investigação de mutação.
**Diversidade genética e experiência prática destacam tanto o poder quanto as limitações dos testes e modelos.**
- Estima-se que cada indivíduo tenha em torno de 50 milhões de variantes (SNPs), contextualizando a diversidade genética usada em bancos de dados para estratificação de risco.
- Relato pessoal: seis pessoas muito próximas na família com mutação BRCA, ilustrando impacto familiar de mutações de alta penetrância.

---

### Chunk 28/30
**Article:** Mastologia I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.574

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

### Chunk 29/30
**Article:** Mastologia I (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.573

ntando significativamente o risco de morte. A análise estatística, como a sobrevida em 5 anos, requer cautela, pois pode ser influenciada por vieses de diagnóstico precoce.
---
### Evidências Chave
**A incidência e mortalidade por câncer de mama apresentam uma projeção de crescimento alarmante, com a estimativa de 74 mil novos casos anuais no Brasil para o próximo triênio.**
- Este número representa um aumento em relação ao ano anterior, que era de 67 mil casos esperados.
- As projeções de longo prazo indicam que, até 2040, a incidência da doença pode aumentar em 40% e a mortalidade em 50%.
- Em contraste, dados dos Estados Unidos sugerem uma desaceleração na mortalidade por câncer de mama a partir de 2021.

---

### Chunk 30/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.572

re BI-RADS 3, menor que o esperado.
   - Pode haver exagero em exames muito frequentes para BI-RADS 3; orientação: em nódulo novo na menopausa, investigar, acompanhar semestralmente no primeiro ano e depois reduzir intensidade do seguimento, evitando impacto emocional desnecessário.
* **Biópsias: tipos e acurácia**
   - Punção aspirativa por agulha fina (PAAF): excelente para esvaziar líquidos; fornece análise citológica apenas, limitada para caracterização completa.
   - Biópsia de fragmento (core) dá nome e sobrenome ao nódulo, devendo ser preferida em nódulos novos na menopausa.
   - Padrão-ouro comparativo é a cirurgia (maior volume de tecido à patologista); biópsia a vácuo produz fragmentos maiores e, em estudos, tem comparação com biópsia cirúrgica de quase 100%, resolvendo muitas questões e muitas vezes removendo a lesão sem necessidade de sala cirúrgica.
### 3.

---

