# ScoreItem: Mastectomia

**ID:** `019bf31d-2ef0-7270-8a8b-017b293ca147`
**FullName:** Mastectomia (Histórico de doenças - Cirurgias já realizadas - Cirurgias que interferem diretamente no escore)

**Preparation Metadata:**
- Quality Grade: **GOOD**
- Total Chunks: 30 de 12 artigos
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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-7270-8a8b-017b293ca147`.**

```json
{
  "score_item_id": "019bf31d-2ef0-7270-8a8b-017b293ca147",
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

**ScoreItem:** Mastectomia (Histórico de doenças - Cirurgias já realizadas - Cirurgias que interferem diretamente no escore)

**30 chunks de 12 artigos (avg similarity: 0.573)**

### Chunk 1/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.681

dos sintomas gerais a investigar em consulta, como dor mamária e TPM intensa, que podem indicar predominância estrogênica.

## Objetivo:
Não há achados de exame físico de um paciente específico. A abordagem objetiva descrita inclui:
- Avaliação da composição corporal por bioimpedância ou densitometria de corpo total, identificando condições como sarcopenia em mulheres com peso normal.
- Solicitação de exames laboratoriais para avaliar marcadores inflamatórios e verificar se a paciente está metabolicamente doente ou em risco.

## Diagnóstico Primário:
- Avaliação: Discussão educacional sobre estratificação de risco para câncer de mama. 90% dos casos relacionam-se ao estilo de vida e 10% a fatores genéticos conhecidos. Mutações de alta penetrância (ex.: BRCA1, BRCA2, TP53) aumentam significativamente o risco, porém são raras. Epigenética e estilo de vida são cruciais e modificáveis.

---

### Chunk 2/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.642

namnese detalhada abrangendo hábitos de vida, saúde intestinal, sintomas hormonais (dor mamária, TPM), história oncológica familiar (ambos os lados) e rede de apoio.
- [ ] Avaliar composição corporal com bioimpedância ou densitometria de corpo total para detectar sarcopenia e orientar intervenção.
- [ ] Solicitar e acompanhar marcadores inflamatórios/metabólicos para monitorar resposta a intervenções de estilo de vida, especialmente em pacientes com diagnóstico prévio de câncer.
- [ ] Utilizar calculadoras de risco (p. ex., Gail e ferramentas que integram PRS e fatores de estilo de vida) com cautela, reconhecendo superestimação em populações brasileiras, e empregá-las para motivar mudanças comportamentais.
- [ ] Considerar farmacogenética ao selecionar terapias, avaliando polimorfismos que alteram metabolização de drogas.
- [ ] Preparar materiais e plano para a próxima aula focada em dúvidas sobre reposição hormonal.

---

### Chunk 3/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** methods | **Similarity:** 0.630

; uso de medicamentos/suplementos; hábitos de vida (alimentação, exercício, álcool, tabaco); saúde intestinal; níveis de estresse/ansiedade; rede de apoio.
   - Atenção a sintomas de predominância estrogênica: dor mamária, TPM intensa; lacunas de cuidado em pacientes com alterações fibrocísticas e cistos, frequentemente não acolhidas apesar de sintomas de excesso de estrogênio circulante.
* **Avaliação de composição corporal**
   - Preferência por métodos além da balança: bioimpedância (utilizada pela instrutora) ou densitometria de corpo total.
   - Observação de sarcopenia em mulheres com peso normal e desconhecimento da condição; necessidade de intervenção.
* **Marcadores laboratoriais e metabolicidade**
   - Solicitar marcadores inflamatórios para identificar adoecimento metabólico ou risco, e para monitorar evolução após intervenções.

---

### Chunk 4/30
**Article:** Mastologia I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.606

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

### Chunk 5/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.604

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

### Chunk 6/30
**Article:** Lifestyle Medicine: A Brief Review of Its Dramatic Impact on Health and Survival (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.602

s for patients 
with breast cancer to improve prognosis and optimize 
overall health. CMAJ 2017 Feb 1;189(7):E268-74. 
DOI: https://doi.org/10.1503/cmaj.160464.
 113. Irwin ML, McTiernan A, Manson JE, et al. Physical 
activity and survival in postmenopausal women with 
breast cancer: Results from the women’s health 
initiative. Cancer Prev Res (Phila) 2011 Apr;4(4):522-9. 
DOI: https://doi.org/10.1158/1940-6207.capr-10-0295.
 114. Chlebowski RT. Nutrition and physical activity influence 
on breast cancer incidence and outcome. Breast 2013 
Aug;22 Suppl 2:S30-7. DOI: https://doi.org/10.1016/j.
breast.2013.07.006.
 115. Meyerhardt JA, Heseltine D, Niedzwiecki D, et al. 
Impact of physical activity on cancer recurrence 
and survival in patients with stage III colon cancer: 
Findings from CALGB 89803. J Clin Oncol 2006 
Aug 1;24(22):3535-41. DOI: https://doi.org/10.1200/
jco.2008.26.15_suppl.4039.
 116. Pierce JP, Stefanick ML, Flatt SW, et al.

---

### Chunk 7/30
**Article:** Long-Term Effects of Breast Cancer Surgery, Treatment, and Survivor Care (2019)
**Journal:** Journal of Midwifery and Womens Health
**Section:** abstract | **Similarity:** 0.596

Comprehensive review showing that 90% of breast cancer survivors experience lasting physical, metabolic, and psychosocial complications requiring multidisciplinary long-term care.

---

### Chunk 8/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.592

luto, evitar alarmismo, e definir critérios para exames complementares (ressonância) apenas quando houver fatores de risco adicionais.
- [ ] 4. Atualizar materiais educativos para esclarecer que história familiar, por si só, não contraindica reposição; incorporar achados do Sister Study e WHI.
- [ ] 5. Estabelecer diretriz interna: não indicar reposição hormonal sistêmica em pacientes com histórico de câncer de mama; considerar terapias tópicas para atrofia vaginal após tentativa de métodos não hormonais, com suporte emocional.
- [ ] 6. Criar protocolo de uso criterioso de gestrinona em endometriose e mastalgia refratária, com consentimento informado sobre lacunas de evidência oncológica.
- [ ] 7. Definir critérios de indicação de testosterona por motivos não oncológicos, evitando prescrição para “redução de risco mamário” até que haja validação em guidelines.
- [ ] 8.

---

### Chunk 9/30
**Article:** Mastologia I (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.590

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

### Chunk 10/30
**Article:** Mastologia I (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.588

orporar na prática clínica a orientação sobre exercícios de força, além de atividades aeróbicas, para prevenção e melhor prognóstico.
- [ ] 4. Ler artigo indicado sobre novas tecnologias e estratégias de diagnóstico do câncer de mama para atualização.
- [ ] 5. Usar analogias, como a do “carro” e do “acidente”, para explicar às pacientes a diferença entre redução de risco (prevenção) e diagnóstico precoce.

---

## Quantitative Data

### Narrativa Quantitativa
O cenário do câncer de mama no Brasil revela uma tendência alarmante de crescimento, com uma projeção de 74 mil novos casos anuais e um aumento futuro de 40% na incidência e 50% na mortalidade até 2040. Fatores metabólicos como hiperinsulinemia e sarcopenia agravam drasticamente este quadro, aumentando significativamente o risco de morte. A análise estatística, como a sobrevida em 5 anos, requer cautela, pois pode ser influenciada por vieses de diagnóstico precoce.

---

### Chunk 11/30
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

### Chunk 12/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.581

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

### Chunk 13/30
**Article:** Mastologia I (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.581

estão associadas a um aumento drástico no risco de mortalidade em pacientes com câncer.**
- Mulheres com hiperinsulinemia apresentaram um risco 34% maior de desenvolver câncer e um risco 78% maior de morte após o diagnóstico, independentemente do IMC ou da circunferência abdominal.
- Pacientes com sarcopenia (perda de massa muscular) tiveram um aumento de 93% nas mortes por câncer em geral e, especificamente em casos de câncer de mama, a mortalidade foi 41% maior.
- Uma meta-análise também mostrou que a sarcopenia aumentou em 44% as mortes por todas as causas.
**A métrica de "sobrevida em 5 anos", embora comum em oncologia, pode ser enganosa devido a vieses estatísticos relacionados ao momento do diagnóstico.**
- A sobrevida em 5 anos é uma métrica frequentemente usada para avaliar a eficácia percebida do rastreamento mamográfico.

---

### Chunk 14/30
**Article:** Metabolic syndrome is associated with breast cancer recurrence and breast cancer-specific mortality (2025)
**Journal:** Journal of Internal Medicine
**Section:** abstract | **Similarity:** 0.575

Meta-analysis of 42,135 survivors showing metabolic syndrome associated with 69% higher recurrence risk (HR 1.69) and 83% higher breast cancer mortality (HR 1.83).

---

### Chunk 15/30
**Article:** Fisiologia Endócrina Feminina (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.561

o de risco individual antes de terapia hormonal: histórico pessoal/familiar de câncer de mama, trombose, risco cardiovascular; densidade mineral óssea.
    - Diferenciar fogachos de outras causas de “calor” (carcinoide, mastocitose, fármacos, ansiedade, etc.).
    - Considerar perfil lipídico, marcadores inflamatórios, saúde óssea (densitometria), saúde urogenital e qualidade do sono.
    - Considerar intervenções graduais na transição menopausal (reposição de progesterona, estradiol, testosterona) conforme deficiência, indicação e riscos.
    - Educação da paciente para adesão terapêutica informada e tomada de decisão compartilhada.
- Plano de Tratamento de Seguimento:
  - Mudanças de estilo de vida:
    - Atividade física regular, com ênfase em treino de resistência (~250 minutos semanais) para saúde óssea, muscular e geral.
    - Higiene do sono (priorizar sono profundo entre ~22h–5h).

---

### Chunk 16/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.559

dos não hormonais.
   - Abordagem em etapas: iniciar com não hormonais; se falhar e houver queixas significativas, considerar tópico. Importante manejo emocional/psicológico, acolhimento de sintomas, e ausência de pressa, individualizando.
### 7. Princípios de acompanhamento e cuidado integral
* **Evitar sobrevigília desnecessária**
   - Reposição hormonal não implica necessidade de acompanhar mamas “de três em três meses” para segurança; tal prática pode aumentar ansiedade sem base em evidência.
* **Saúde integral reduz risco**
   - Tudo que se faz em prol da saúde integral da paciente tende a reduzir chances de câncer de mama; a mama deve ser considerada no contexto do corpo inteiro, não isoladamente.

---

### Chunk 17/30
**Article:** MFI - Reposição Hormonal - AULA 07 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.558

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
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.557

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

### Chunk 19/30
**Article:** Mastologia I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.553

ltidisciplinar.
- Converta conhecimento em mudança sustentada com acompanhamento coordenado e metas claras.
### Prevenção Metabólica
Prevenção eficaz mira inflamação, resistência à insulina e composição corporal.  
- Mamografia é detecção precoce e não prevenção, com benefícios limitados em câncer avançado e risco de sobretratamento.
- Rastreie e corrija hiperinsulinemia e inflamação crônica como alvos primários de risco e mortalidade.
- Combata obesidade e sarcopenia com treino de força para remodelar o microambiente mamário e reduzir risco.
### Ambiente, Hormônios e Precisão
Risco é modelado por ambiente, hábitos e regulação hormonal, guiando rastreio de precisão.  
- Migração para países de alto IDH eleva incidência, destacando impacto ambiental e de estilo de vida.
- Mantenha hormônios em faixa ótima para reduzir risco por desbalanços.

---

### Chunk 20/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XIV (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.552

- [ ] 2. Para pacientes com síndrome metabólica, considerar estratégia Mediterrânea organizada; preferir versão com restrição calórica combinada com atividade física, monitorando resultados aos 6 e 12 meses.
- [ ] 3. Solicitar perfil lipídico ampliado em casos selecionados: HDL, triglicerídeos, insulina, PCR, LDL oxidado e, quando indicado, subfracionamento de LDL; evitar decisões baseadas apenas em LDL total.
- [ ] 4. Planejar exames de risco cardiovascular conforme necessidade: score de cálcio coronariano e angiotomografia de coronárias (incluindo avaliação de placas moles) quando o contexto clínico justificar.
- [ ] 5. Revisar e atualizar protocolos internos sobre álcool: remover recomendações de consumo “cardioprotetor”; educar pacientes sobre riscos de câncer e piora do sono; avaliar predisposição/sensibilidade individual.
- [ ] 6.

---

### Chunk 21/30
**Article:** Dieta Cetogênica - Parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.550

atamento metabólico do câncer
   - Terapias de pressão (contínuas): dieta cetogênica, cetonas exógenas, suplementos/fitoterápicos/drogas individualizadas, manejo do estresse emocional.
   - Terapias de pulso (intermitentes): inibição de glicose, inibição de glutamina, oxigenoterapia hiperbárica, entre outras.
   - Abordagem integrada e personalizada para maximizar o controle tumoral.
* Ensaio clínico randomizado (2021) em câncer de mama
   - 80 pacientes tratados com quimio; randomização para dois grupos; intervenção cetogênica/metabólica por 12 semanas; exames laboratoriais e de imagem no início e 12 semanas; cirurgia e reestadiamento para doença localmente avançada após quimio.
   - Resultados: redução de TNF-α, IGF-1, insulina; aumento de IL-10; redução significativa do tamanho tumoral no grupo cetogênico.

---

### Chunk 22/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.547

urgias profiláticas) e a importância do aconselhamento genético. Apresenta calculadoras de risco (Gail como a mais conhecida), observa superestimação fora de populações de origem e sugere uso pedagógico para impulsionar mudanças de estilo de vida. Destaca anamnese detalhada, avaliação de composição corporal e marcadores metabólicos/inflamatórios como base prática de estratificação. Conclui que genética não é destino, introduz epigenética como fator modificável e informa que dúvidas sobre reposição hormonal serão abordadas na próxima aula. Data de criação: 2025-11-21.
## 🔖 Pontos de Conhecimento
### 1. Genética e câncer de mama
* **Proporção de câncer de mama ligado à genética**
   - Aproximadamente 10% dos tumores de mama diagnosticados relacionam-se a mutações genéticas conhecidas; os demais 90% não apresentam achados genéticos e são majoritariamente atribuídos ao estilo de vida.

---

### Chunk 23/30
**Article:** Lifestyle Medicine: A Brief Review of Its Dramatic Impact on Health and Survival (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.547

h 
reported in 2014, from worldwide data, that 
diet, physical activity, and weight control are 
major contributors to long-term survival 
after a diagnosis of breast cancer.
110
 Fur
-
thermore, a 2011 meta-analysis of postdiag
-
nosis exercise in patients with breast cancer 
involving more than 12,000 patients dem
-
onstrated a 34% decrease in risk of death 
caused by breast cancer, a 24% decrease in 
recurrence, and a 41% decrease in the risk of 
all-cause mortality.
111
 is conclusion is the 
result of the review of 67 published articles 
addressing lifestyle changes as they relate to 
the reduction of breast cancer recurrence.
112
 
Additional studies have documented 
that physical activity not only increases 
survival and decreases recurrence but also 
improves overall quality of life in patients 
with breast cancer and in patients with 
colon cancer.
113-115
 Another study followed 
nearly 1500 women diagnosed with early-
stage breast cancer.

---

### Chunk 24/30
**Article:** Mastologia I (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.544

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

### Chunk 25/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.543

eres jovens em idade reprodutiva sofrem complicações do tratamento; uma paciente relatou “pagar o preço da cura”, destacando a necessidade urgente de soluções para qualidade de vida pós-câncer.
## ❓ Perguntas
- [Inserir Pergunta/Confusão]
## 📚 Tarefas
- [ ] 1. Revisar protocolos internos para reposição hormonal, diferenciando claramente progestágenos sintéticos de progesterona micronizada e ajustando práticas conforme evidências de ECRs.
- [ ] 2. Implementar fluxos de manejo de nódulos BI-RADS 3: investigação adequada de nódulo novo na menopausa, seguimento semestral no primeiro ano, redução após, e uso de biópsia a vácuo quando indicado.
- [ ] 3. Padronizar comunicação aos pacientes sobre densidade mamária: explicar diferença entre risco relativo e absoluto, evitar alarmismo, e definir critérios para exames complementares (ressonância) apenas quando houver fatores de risco adicionais.
- [ ] 4.

---

### Chunk 26/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.542

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

### Chunk 27/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.536

RCA2, TP53) aumentam significativamente o risco, porém são raras. Epigenética e estilo de vida são cruciais e modificáveis. A avaliação individualizada requer anamnese detalhada, composição corporal e marcadores laboratoriais.
- Diagnóstico Suspeito: Nenhum no momento

## Plano:
- Prescrição: Inserir mais aqui
- Próximos Passos:
  - Suspeitar de mutação genética e considerar teste genético em casos de múltiplos cânceres na família (mama, ovário, pâncreas, próstata), câncer de mama triplo negativo, câncer de mama em idade jovem (<45 anos), câncer de mama em homem, ou descendência judaica Ashkenazi.
  - Encaminhar para aconselhamento genético antes do teste em pacientes com alta suspeita de mutação.
  - Utilizar calculadoras de risco (ex.: Gail, Tyrer-Cuzick) para conscientizar sobre mudanças no estilo de vida, reconhecendo limitações na população brasileira.

---

### Chunk 28/30
**Article:** Cardiologia VI (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.536

es com insuficiência cardíaca congestiva, conforme um estudo de 2010 com acompanhamento de 6 meses.
**A obesidade e a resistência insulínica emergem como fatores de risco independentes e severos, aumentando a chance de hipertensão em 3,5 vezes e reduzindo a expectativa de vida em até 20 anos em casos de obesidade mórbida.**
- Pacientes obesos com síndrome metabólica apresentam um risco de mortalidade 18% maior e uma redução na expectativa de vida de 5 a 20 anos.
- Um estudo de 2017 mostrou que pacientes obesos com síndrome metabólica têm 3,5 vezes mais chances de se tornarem hipertensos.
- A resistência insulínica foi identificada como um fator de risco para doença cardiovascular já em 1996 e foi prevalente em 26% das pacientes não diabéticas com câncer de mama em um estudo de 2012-2014 com 760 mulheres.

---

### Chunk 29/30
**Article:** Trato Gastrointestinal V – Intestino Delgado I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.535

(como análogos de GLP-1), mas explicando o metabolismo e oferecendo estratégias alternativas.
    - A medicina funcional integrativa exige uma compreensão profunda dos sistemas do corpo, pois a saúde metabólica impacta tudo, incluindo a metabolização de hormônios.
*   **Exemplo Clínico: Reposição Hormonal**
    - É apresentado o caso de uma paciente com histórico familiar de câncer de mama que buscava reposição hormonal. Médicos anteriores prescreveram hormônios sem abordar sua saúde metabólica, o que é crucial para garantir que os hormônios sejam metabolizados de forma segura.
    - A abordagem correta seria organizar a vida da paciente, incluindo a alimentação, mesmo ela sendo magra, pois magreza não é sinônimo de saúde.
### 4.

---

### Chunk 30/30
**Article:** Mastologia III (2024)
**Journal:** Pos Graduacao MFI
**Section:** methods | **Similarity:** 0.531

rar exames adicionais (p.ex., ressonância magnética) conforme caso, não de rotina.
    - Atrofia vaginal: abordagem escalonada, iniciando com métodos não hormonais e, se necessário, terapia hormonal tópica.
- Plano de Tratamento de Acompanhamento:
    - RH deve ser individualizada, considerando particularidades de cada paciente, em vez de protocolo único.
    - Aconselhamento e educação para desmistificar medos e alinhar expectativas, especialmente sobre nódulos benignos e manejo de mamas densas.
    - RH não recomendada atualmente para pacientes pós-câncer de mama devido a ensaios clínicos que mostram aumento de recorrência.
    - Testosterona não deve ser indicada com objetivo de reduzir risco de câncer de mama, pois isso não é validado.
    - Enfatizar aconselhamento sobre estilo de vida e saúde geral para reduzir risco de câncer de mama.

---

