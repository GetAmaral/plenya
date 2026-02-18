# ScoreItem: Hipercalórica

**ID:** `019c534c-0bf6-7d6b-9c60-d650fa5f846f`
**FullName:** Hipercalórica (Alimentação - Atual (últmos 6 meses) - Padrão alimentar atual)

**Preparation Metadata:**
- Quality Grade: **GOOD**
- Total Chunks: 30 de 18 artigos
- Avg Similarity: 0.544

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019c534c-0bf6-7d6b-9c60-d650fa5f846f`.**

```json
{
  "score_item_id": "019c534c-0bf6-7d6b-9c60-d650fa5f846f",
  "clinical_relevance": "Texto técnico para médicos (1000-5000 chars): definição fisiológica precisa, valores de referência e interpretação, fisiopatologia resumida, dados epidemiológicos com números concretos, estratificação de risco baseada em evidências.",
  "points": 0,
  "patient_explanation": "Texto simples para pacientes (500-1000 chars): o que é este parâmetro sem jargões, por que é importante para a saúde, o que valores alterados podem significar. Tom tranquilizador e educativo.",
  "conduct": "Conduta clínica em Markdown (1000-5000 chars): investigação complementar necessária, critérios de encaminhamento a especialistas, intervenções baseadas em evidências. Use bullet points, seções e negrito."
}
```

**Nota:** `points` deve ser `0` neste item — não calcule pontuação.

---

### Contexto Científico

**ScoreItem:** Hipercalórica (Alimentação - Atual (últmos 6 meses) - Padrão alimentar atual)

**30 chunks de 18 artigos (avg similarity: 0.544)**

### Chunk 1/30
**Article:** Mastologia I (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.619

estão associadas a um aumento drástico no risco de mortalidade em pacientes com câncer.**
- Mulheres com hiperinsulinemia apresentaram um risco 34% maior de desenvolver câncer e um risco 78% maior de morte após o diagnóstico, independentemente do IMC ou da circunferência abdominal.
- Pacientes com sarcopenia (perda de massa muscular) tiveram um aumento de 93% nas mortes por câncer em geral e, especificamente em casos de câncer de mama, a mortalidade foi 41% maior.
- Uma meta-análise também mostrou que a sarcopenia aumentou em 44% as mortes por todas as causas.
**A métrica de "sobrevida em 5 anos", embora comum em oncologia, pode ser enganosa devido a vieses estatísticos relacionados ao momento do diagnóstico.**
- A sobrevida em 5 anos é uma métrica frequentemente usada para avaliar a eficácia percebida do rastreamento mamográfico.

---

### Chunk 2/30
**Article:** Jejum Intermitente (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.591

- Prescrição clínica: 5 dias de FMD (dieta que mimetiza o jejum) com quimioterapia no 3º dia; manter FMD por mais 2 dias após.
- Três dias de alta dose de vitamina C pós-quimioterapia atuam como pró-oxidante, visando auxiliar oxidação e potencial clearance tumoral.
**[Achados Adicionais]**
- Restrição calórica (não jejum) em animais costuma variar entre 10–40%, contextualizando efeitos de longevidade fora dos protocolos de jejum.
- Registro histórico extremo: Angus Barberi perdeu 125 kg após 382 dias de jejum, indo de 207 kg para 82 kg; ilustra capacidade de perda de peso, mas não é diretriz clínica.

---

## Concept Insights

### Vulnerabilidade Metabólica Diferencial
**Categoria:** Modelo Mental
**Definição Central:**
A ideia de que células cancerígenas e células normais possuem vulnerabilidades metabólicas distintas que podem ser exploradas terapeuticamente.

---

### Chunk 3/30
**Article:** Bioquímica Metabólica nos Exercícios - Parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.580

a não consensual; dose prática 1 g 3x/dia.
### 13. Hipertrofia: inflamação e modulação
- Hipertrofia depende de sobrecarga mecânica, microlesões, grande processo inflamatório e aumento de síntese proteica.
- IL-6, ERO e lactato são sinalizações úteis; evitar anti-inflamatórios/crioterapia e excesso de antioxidantes imediatamente após.
- Demandas proteicas aumentam com VO2, intensidade e frequência.
### 14. Déficit energético crônico e sinais clínicos
- Indicativos: amônia, ureia, ácido úrico, transaminases, cortisol altos; queda de performance e de massa; desidratação; pior recuperação.
- Sinais: queda de cabelo, unhas quebradiças, imunidade baixa; bioimpedância mostra alterações de água; possível aumento de TSH e queda de T3 por déficit energético (pseudo-hipotireoidismo).
- Hipertrofia inviável sob catabolismo salvo intervenções hormonais não-mTOR com resultados limitados.
### 15.

---

### Chunk 4/30
**Article:** Preparing Patients for Cosmetic Surgery and Aesthetic Procedures: Ensuring an Optimal Nutritional Status for Successful Results (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.574

ithEatingDisorders.Am.J.Clin.Dermatol.2005,6,165–173.[CrossRef]40.Englesbe,M.J.;Lee,J.S.;He,K.;Fan,L.;Schaubel,D.E.;Sheetz,K.H.;Harbaugh,C.M.;Holcombe,S.A.;Campbell,D.A.J.;Sonnenday,C.J.;etal.AnalyticMorphomics,CoreMuscleSize,andSurgicalOutcomes.Ann.Surg.2012,256,255–261.[CrossRef]41.Lee,J.S.;Terjimanian,M.N.;Tishberg,L.M.;Alawieh,A.Z.;Harbaugh,C.M.;Sheetz,K.H.;Holcombe,S.A.;Wang,S.C.;Sonnenday,C.J.;Englesbe,M.J.SurgicalSiteInfectionandAnalyticMorphometricAssessmentofBodyCompositioninPatientsUndergoingMidlineLaparotomy.J.Am.Coll.Surg.2011,213,236–244.[CrossRef]42.Tan,B.H.L.;Birdsell,L.A.;Martin,L.;Baracos,V.E.;Fearon,K.C.H.SarcopeniainanOverweightorObesePatientisanAdversePrognosticFactorinPancreaticCancer.Clin.CancerRes.2009,15,6973–6979.[CrossRef]43.Gillis,C.;Hasil,L.;Kasvis,P.;Bibby,N.;Davies,S.J.;Prado,C.M.;West,M.A.;Shaw,C.NutritionCareProcessModelApproachtoSurgicalPrehabilitationinOncology.Front.Nutr.2021,8,644706.[CrossRef]44.Mazza,E.;Ferro,Y.;Pujia,R.;Mare,R.;Maurotti,S.

---

### Chunk 5/30
**Article:** Preparing Patients for Cosmetic Surgery and Aesthetic Procedures: Ensuring an Optimal Nutritional Status for Successful Results (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.567

ry,J.A.;Cross,K.M.DevelopmentofaUniversalNutritionalScreeningPlatformforPlasticSurgeryPatients.Plast.Reconstr.Surg.Glob.Open2017,5,e1342.[CrossRef][PubMed]80.Detsky,A.S.;McLaughlin,J.R.;Baker,J.P.;Johnston,N.;Whittaker,S.;Mendelson,R.A.;Jeejeebhoy,K.N.WhatisSubjectiveGlobalAssessmentofNutritionalStatus?JPENJ.Parenter.EnteralNutr.1987,11,8–13.[CrossRef]81.Kondrup,J.;Rasmussen,H.H.;Hamberg,O.;Stanga,Z.;AdHocESPENWorkingGroup.NutritionalRiskScreening(NRS2002):ANewMethodBasedonanAnalysisofControlledClinicalTrials.Clin.Nutr.2003,22,321–336.[CrossRef]82.Bakaloudi,D.R.;Halloran,A.;Rippin,H.L.;Oikonomidou,A.C.;Dardavesis,T.I.;Williams,J.;Wickramasinghe,K.;Breda,J.;Chourdakis,M.IntakeandAdequacyoftheVeganDiet.ASystematicReviewoftheEvidence.Clin.Nutr.2021,40,3503–3521.[CrossRef]83.Vitiello,V.;Germani,A.;CapuzzoDolcetta,E.;Donini,L.M.;DelBalzo,V.TheNewModernMediterraneanDietItalianPyramid.Ann.Ig.2016,28,179–186.[PubMed]84.Wong,C.J.InvoluntaryWeightLoss.Med.Clin.NorthAm.2014,98,625–643.[

---

### Chunk 6/30
**Article:** Jejum Intermitente (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.560

quercetina, acetil-L-carnitina) conforme indicação profissional.
- [ ] 9. Em oncologia, considerar FMD ou cetogênica como adjuvante à quimioterapia apenas com aval e acompanhamento de oncologista; seguir protocolos específicos (ex.: quimio no 3º dia de FMD; vitamina C em alta dose por 3 dias).
- [ ] 10. Comparar aderência e resultados entre jejum intermitente e restrição calórica contínua; escolher abordagem com maior probabilidade de manutenção pelo paciente.
- [ ] 11. Educar pacientes sobre mecanismos do jejum (sirtuínas, AMPK, mTOR, BDNF, autofagia) para promover compreensão e adesão informada.

---

## Quantitative Data

### Narrativa Quantitativa
O conjunto de métricas revela que o jejum intermitente, especialmente em janelas de alimentação restritas e modelos 5-2, tende a oferecer benefícios cardiometabólicos e perda de peso comparáveis à restrição calórica contínua, com nuances importantes de horário e adesão.

---

### Chunk 7/30
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

### Chunk 8/30
**Article:** Preparing Patients for Cosmetic Surgery and Aesthetic Procedures: Ensuring an Optimal Nutritional Status for Successful Results (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.555

Germani,A.;CapuzzoDolcetta,E.;Donini,L.M.;DelBalzo,V.TheNewModernMediterraneanDietItalianPyramid.Ann.Ig.2016,28,179–186.[PubMed]84.Wong,C.J.InvoluntaryWeightLoss.Med.Clin.NorthAm.2014,98,625–643.[CrossRef][PubMed]85.Cederholm,T.;Bosaeus,I.;Barazzoni,R.;Bauer,J.;VanGossum,A.;Klek,S.;Muscaritoli,M.;Nyulasi,I.;Ockenga,J.;Schneider,S.M.;etal.DiagnosticCriteriaforMalnutrition—AnESPENConsensusStatement.Clin.Nutr.2015,34,335–340.[CrossRef][PubMed]86.Matory,W.E.J.;O’Sullivan,J.;Fudem,G.;Dunn,R.AbdominalSurgeryinPatientswithSevereMorbidObesity.Plast.Reconstr.Surg.1994,94,976–987.[CrossRef][PubMed]87.Gounden,V.;Vashisht,R.;Jialal,I.Hypoalbuminemia.InStatPearls;Anonymous;StatPearlsPublishingLLC:TreasureIsland,FL,USA,2022.88.Muscaritoli,M.;Arends,J.;Bachmann,P.;Baracos,V.;Barthelemy,N.;Bertz,H.;Bozzetti,F.;Hutterer,E.;Isenring,E.;Kaasa,S.;etal.ESPENPracticalGuideline:ClinicalNutritioninCancer.Clin.Nutr.2021,40,2898–2913.[CrossRef]89.Tuck,C.J.;Biesiekierski,J.R.;Schmid-Grendelmeier,P.

---

### Chunk 9/30
**Article:** Bioquímica Metabólica nos Exercícios - Parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.553

do-hipotireoidismo).
- Hipertrofia inviável sob catabolismo salvo intervenções hormonais não-mTOR com resultados limitados.
### 15. Ferramentas de controle: limiares, zonas e FIT
- Avaliar no esporte real; definir limiar via lactato e prescrever supra-limiar (acidose controlada) ou FatMax (entre 1º e 2º limiar) para mobilização de gordura sem excessiva acidose.
- Framework FIT: frequência, intensidade, tipo e tempo; monitorar FC, estado ácido-base, marcadores de dano muscular, fontes energéticas e risco de overtraining.
### 16. Estratégia clínica integrativa e acompanhamento
- Basear-se na história clínica, nutrição, bioquímica/metabolismo, estilo de vida, equilíbrio hormonal.
- Iniciar com exames simples (sangue, bioimpedância), aplicar intervenções personalizadas e reavaliar em 1–2 meses, mantendo ciclo de melhoria contínua.
## ❓ Perguntas
- [Inserir Pergunta/Confusão]
## 📚 Tarefas/Assignments
- [ ] 1.

---

### Chunk 10/30
**Article:** Bioquímica Metabólica nos Exercícios - Parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.551

avaliação cardiometabólica e a calorimetria indireta para personalizar as estratégias nutricionais (como a indicação ou não de jejum intermitente) e de treinamento.
*   [ ] 5. Para pacientes "glicolíticos", prescrever uma combinação de exercícios aeróbicos de maior volume com treinos de força para aumentar a massa muscular e a capacidade mitocondrial.
*   [ ] 6. Considerar a prescrição de "recuperação ativa" (15-20 minutos de exercício de baixa intensidade pós-treino) para pacientes que treinam em dias consecutivos com menos de 24 horas de intervalo.
*   [ ] 7. Evitar o uso indiscriminado de anti-inflamatórios para dores musculares e buscar alternativas de modulação inflamatória que não prejudiquem a hipertrofia.
*   [ ] 8. Acompanhar a recuperação da frequência cardíaca pós-treino dos pacientes como uma forma subjetiva de monitorar a adaptação e o condicionamento.
*   [ ] 9.

---

### Chunk 11/30
**Article:** Emagrecimento - Parte I (2024)
**Journal:** Pos Graduacao MFI
**Section:** methods | **Similarity:** 0.545

s com peso normal se enquadram nessa categoria.
*   **Métodos de Avaliação Adequados**
    - Composição corporal deve ser avaliada por dobras cutâneas ou bioimpedanciometria.
    - Dois indivíduos com mesmo peso e altura (mesmo IMC) podem ser metabolicamente opostos: um predominância de gordura, outro de músculo.
*   **Cirurgia Bariátrica como Recurso**
    - Válida, porém último recurso após esgotar outras tentativas.
    - Cirurgias aumentaram 85% (2011–2018): 60% bypass e 36% sleeve.
    - Critica prática antiética de orientar ganho de peso para qualificar pelo convênio.
    - Pós-bariátricos enfrentam riscos como alcoolismo, depressão e suicídio; necessitam acompanhamento multidisciplinar e funcional, raramente realizado.

## ❓ Perguntas
- [Inserir Pergunta/Confusão]

## 📚 Tarefas
- [ ] 1. Refletir sobre a prática profissional no emagrecimento e identificar lacunas de conhecimento (fisiologia, intestino, mitocôndrias, inflamação, etc.).
- [ ] 2.

---

### Chunk 12/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XIX (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.543

al para prevenir complicações.
*   **Estratégias para Preservar Massa Magra no Emagrecimento**
    - A prática de exercícios resistidos é fundamental para a manutenção da massa magra.
    - A ingestão de proteína deve ser superior às diretrizes atuais de 0,8 g/kg de peso corporal por dia. Para pacientes em tratamento para obesidade, a recomendação é de mais de 1.5 g/kg de peso para evitar o catabolismo muscular, especialmente ao usar medicamentos que reduzem o apetite.
    - Se a meta proteica não for atingida pela dieta, deve-se considerar a suplementação com whey protein ou shakes de proteína.
*   **Abordagem Pré-Tratamento da Obesidade (Medicamentosa ou Cirúrgica)**
    - Antes de iniciar qualquer tratamento, é essencial discutir o peso, os hábitos alimentares (snacks, refrigerantes), o uso de suplementos e avaliar as causas e complicações da obesidade.

---

### Chunk 13/30
**Article:** Emagrecimento - Parte VII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.540

l e decisão terapêutica
- IMC e percentual de gordura orientam a estratégia: com IMC adequado e % gordura alto, iniciar ajuste alimentar e tentar emagrecer.
- Se progresso lento/estagnação, migrar o foco para ganho de massa muscular.
- Restrições prolongadas podem reduzir metabolismo basal e função tireoidiana; risco de queda de força, cabelo e energia.
- Reversão requer aumento de massa muscular e maior aporte proteico, idealmente com nutricionistas.
- Indicadores para mudar estratégia: 6–8 semanas sem progresso, sinais de baixa força/energia, plateaus persistentes.
### 10. Papel do nutricionista e personalização
- Emagrecimento efetivo demanda parceria com nutricionista; evitar modelos rígidos e “receitas prontas”.
- Centrar no paciente: começar pelo possível, negociar trocas e adaptar à realidade (ex.: doces menos calóricos).
- Fluxos de referência/retorno e entrevista motivacional facilitam adesão.
### 11.

---

### Chunk 14/30
**Article:** Bioquímica Metabólica nos Exercícios - Parte I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.539

bioenergética e a flexibilidade metabólica.
- [ ] 3. Solicitar e interpretar exames dinâmicos (antes/durante/depois) quando possível; monitorar amônia, ureia, ácido úrico, transaminases e lactato para detectar gliconeogênese/proteólise e correlacionar com zonas de treinamento.
- [ ] 4. Implementar suplementação de glutamina em protocolos de alta intensidade ou em sinais de comprometimento imune/anticatabolismo, conforme quadro clínico.
- [ ] 5. Estruturar estratégia pós-exercício: carboidratos para foco em hipertrofia/recuperação rápida; aminoácidos essenciais para emagrecimento com preservação muscular; evitar ausência total de substrato em casos de risco de proteólise.
- [ ] 6. Promover ganho/manutenção de massa muscular em planos de emagrecimento, especialmente em pacientes com obesidade sarcopênica.
- [ ] 7.

---

### Chunk 15/30
**Article:** Jejum Intermitente (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.535

idativo e sensibilidade à quimioterapia; ativam autofagia e reduzem hiperativação de mTOR.
- Efeitos sistêmicos: menor toxicidade da quimioterapia em células normais; maior sensibilidade tumoral; melhora entrega/clearance de quimioterápicos; redução de fatores de crescimento e inflamação.
- Protocolo integrativo (sob oncologista): 5 dias de FMD com quimioterapia no 3º dia; mais 2 dias de FMD pós-quimio; 3 dias de vitamina C em alta dose como pró-oxidante; somente com aval e acompanhamento especializado.
### 7. Evidência científica sobre emagrecimento
- Revisões comparativas: em 11 estudos, pelo menos 9 mostraram perda de peso similar entre jejum e restrição contínua; jejum é ferramenta adicional, não superior por si.
- Revisão guarda-chuva (JAMA): maior força estatística para jejum modificado alternado (~25% do GET no dia) e 5-2; qualidade da evidência majoritariamente baixa/muito baixa; eficácia depende de população e medidas associadas.
### 8.

---

### Chunk 16/30
**Article:** Dieta Cetogênica - Parte I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.535

oterapia: 61 pacientes em dieta convencional vs 20 em cetogênica (com 10 g de aminoácidos nos dias de radioterapia, BCAA/baixa glutamina); cetogênica associada à perda de massa gorda e preservação de massa magra.
- Pós-pancreatectomia: grupo cetogênico (10 pacientes) teve maior adesão, ingestão calórica e satisfação, sem aumento de complicações, comparado ao grupo padrão (9 pacientes).
- Ensaio randomizado em carcinoma gastrointestinal (23 desnutridos): comparação entre dieta isocalórica convencional vs low carb rica em gordura para suporte metabólico.
- Em atletas semiprofissionais (30 dias): intervenção cetogênica comparada à ocidental padrão para composição corporal/força/metabolismo; grupo cetogênico com ingestão proteica de 1,8 g/kg/dia para preservar massa magra.

---

### Chunk 17/30
**Article:** Dieta Cetogênica - Parte I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.535

ertensos e polifarmácia: monitorar PA, ajustar diuréticos/anti-hipertensivos.
### 11. Composição alimentar e exemplos práticos
- Realidade da cetogênica:
  - Refeições com saladas, vegetais, carnes, queijos, ovos, castanhas; evitar ultraprocessados e óleos ricos em ômega 6.
- Fibras e constipação:
  - Fibras não são determinantes; opções compatíveis: abacate (6,7 g/100 g), folhas, brócolis, chia, psyllium, taioba; mamão (1,8 g/100 g) conforme objetivos.
### 12. Evidências clínicas em oncologia e preservação de massa magra
- Estudos:
  - Pancreatobiliar pós-pancreatectomia: maior adesão/satisfação em cetogênica; segurança mantida.
  - Radioterapia (Ketocomp): perda de gordura, preservação de massa magra; em cabeça e pescoço, cetogênica previne perdas induzidas por quimioterapia.
  - Carcinoma GI desnutridos: low-carb rica em gordura melhora qualidade de vida e preserva massa magra.

---

### Chunk 18/30
**Article:** Emagrecimento - Parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.535

eixos hormonais relevantes (HPA/CRH-ACTH, tireoide/TRH, sexuais), resistência insulínica e sinais de disfunção mitocondrial e desnutrição funcional.
- [ ] Implementar planos de nutrição de precisão: mesmo em estratégias de muito baixa caloria, garantir alta densidade de micronutrientes.
- [ ] Educar sobre riscos de anfetaminas (ex.: Venvanse) no controle da fome, especialmente em crianças, e evitar uso para emagrecimento sem indicação clara e acompanhamento rigoroso.
- [ ] Investigar exposição a disruptores endócrinos e toxinas ambientais (incluindo defensivos agrícolas em frutas) e orientar estratégias de redução (ex.: seleção de alimentos, higienização, origem).
- [ ] Avaliar microbiota intestinal e histórico de uso precoce de antibióticos; planejar intervenções para reequilíbrio (quando aplicável).
- [ ] Preparar pacientes para o “platô” do emagrecimento, explicando a adaptação metabólica e ajustando estratégias sem comprometer a saúde.

---

### Chunk 19/30
**Article:** Emagrecimento - Parte VII (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.534

entes com baixo metabolismo (especialmente mulheres), considerar uma estratégia inicial focada no ganho de massa muscular antes de focar na perda de peso.
- [ ] 4. Priorizar a musculação na prescrição de exercícios, mas sempre adaptar à preferência e ao contexto de vida do indivíduo para garantir a adesão.
- [ ] 5. Iniciar o processo de emagrecimento da maioria dos pacientes com uma abordagem low carb baseada em comida de verdade.
- [ ] 6. Implementar variabilidade nas estratégias alimentares, alternando planos (ex: low carb, jejum, mediterrânea) a cada 2-3 meses para evitar estagnação.
- [ ] 7. Abordar a hierarquia da saúde com os pacientes, enfatizando a importância da gestão do stress, sono e relações saudáveis, além da dieta e exercício.
- [ ] 8. Considerar o uso de esteroides como ferramenta terapêutica para restaurar a funcionalidade muscular em casos específicos, como sarcopenia.
- [ ] 9.

---

### Chunk 20/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** other | **Similarity:** 0.531

andinammationusing10parametersincludingdietaryintake,anthropometricmeasurements,laboratoryindices,andfunctionalcapacity.Thescorerangesfrom0(normal)to
30(severemalnutritionandinammation).MiniNutritionAssessment878Includesassessmentofdietaryintake,mobility,neuropsychology,andsomeanthropometricmeasurements,
includingweightandcalfcircumference.Ascoreof12–14pointsindicatesnormalnutritionstatus,8–11indicatesatriskformalnutrition,and0–7pointsindicatesmalnutrition.

---

### Chunk 21/30
**Article:** KDIGO 2024 Clinical Practice Guideline for the Evaluation and Management of Chronic Kidney Disease (2024)
**Journal:** Kidney International
**Section:** results | **Similarity:** 0.531

adultswithadvancedchronickidneydisease:resultsfromthe
EQUALstudy.JRenNutr.2018;28:165–174.876.LimSL,LinXH,DanielsL.Seven-pointsubjectiveglobalassessmentis
moretimesensitivethanconventionalsubjectiveglobalassessmentin
detectingnutritionchanges.JPENJParenterEnteralNutr.2016;40:966–972.877.GraterolTorresF,MolinaM,Soler-MajoralJ,etal.Evolvingconceptson
inammatorybiomarkersandmalnutritioninchronickidneydisease.Nutrients.2022;14:4297.878.PawlaczykW,RogowskiL,KowalskaJ,etal.Assessmentofthenutritional
statusandqualityoflifeinchronickidneydiseaseandkidneytransplant
patients:acomparativeanalysis.Nutrients.2022;14:4814.879.Epping-JordanJE,PruittSD,BengoaR,etal.Improvingthequalityof
healthcareforchronicconditions.QualSafHealthCare.2004;13:299–305.880.NicollR,RobertsonL,GemmellE,etal.Modelsofcareforchronickidneydisease:asystematicreview.Nephrology(Carlton).2018;23:389–396.881.CollisterD,PyneL,CunninghamJ,etal.Multidisciplinarychronickidneydiseaseclinicpractices:ascopingreview.CanJKidneyHea

---

### Chunk 22/30
**Article:** Mastologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.531

namnese detalhada abrangendo hábitos de vida, saúde intestinal, sintomas hormonais (dor mamária, TPM), história oncológica familiar (ambos os lados) e rede de apoio.
- [ ] Avaliar composição corporal com bioimpedância ou densitometria de corpo total para detectar sarcopenia e orientar intervenção.
- [ ] Solicitar e acompanhar marcadores inflamatórios/metabólicos para monitorar resposta a intervenções de estilo de vida, especialmente em pacientes com diagnóstico prévio de câncer.
- [ ] Utilizar calculadoras de risco (p. ex., Gail e ferramentas que integram PRS e fatores de estilo de vida) com cautela, reconhecendo superestimação em populações brasileiras, e empregá-las para motivar mudanças comportamentais.
- [ ] Considerar farmacogenética ao selecionar terapias, avaliando polimorfismos que alteram metabolização de drogas.
- [ ] Preparar materiais e plano para a próxima aula focada em dúvidas sobre reposição hormonal.

---

### Chunk 23/30
**Article:** Dieta Cetogênica - Parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.527

atamento metabólico do câncer
   - Terapias de pressão (contínuas): dieta cetogênica, cetonas exógenas, suplementos/fitoterápicos/drogas individualizadas, manejo do estresse emocional.
   - Terapias de pulso (intermitentes): inibição de glicose, inibição de glutamina, oxigenoterapia hiperbárica, entre outras.
   - Abordagem integrada e personalizada para maximizar o controle tumoral.
* Ensaio clínico randomizado (2021) em câncer de mama
   - 80 pacientes tratados com quimio; randomização para dois grupos; intervenção cetogênica/metabólica por 12 semanas; exames laboratoriais e de imagem no início e 12 semanas; cirurgia e reestadiamento para doença localmente avançada após quimio.
   - Resultados: redução de TNF-α, IGF-1, insulina; aumento de IL-10; redução significativa do tamanho tumoral no grupo cetogênico.

---

### Chunk 24/30
**Article:** Emagrecimento - Parte VII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.526

ais educativos padronizados (folhetos, vídeos curtos).
- Metas educacionais mensuráveis por consulta (ex.: explicar adipogênese em 3 passos).
### 8. Déficit calórico, preservação de massa muscular e adequação proteica
- Em hipocaloria, alguma perda de massa é aceitável; buscar manter turnover proteico adequado.
- Método prático de porções (mãos, peso/tamanho, proporção no prato) para orientar ingestão.
- Preservar/ganhar massa é desafiador; requer proteínas adequadas mesmo em déficit.
- Mulheres com baixa massa e flacidez tendem a metabolismo basal reduzido; foco inicial em ganho de massa pode ser prioritário.
- Caso pós-parto: alinhar expectativas, priorizando recuperação de massa e metabolismo sobre número da balança.
### 9. Avaliação de composição corporal e decisão terapêutica
- IMC e percentual de gordura orientam a estratégia: com IMC adequado e % gordura alto, iniciar ajuste alimentar e tentar emagrecer.

---

### Chunk 25/30
**Article:** Psiquiatria Metabólica Funcional Integrativa AULA 17 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.523

e (polifenóis): faixa de 100 a 150 miligramas como modulador antioxidante/metabólico.
- Capsaicina/capsiate: 5 a 10 miligramas como agente termogênico e modulador do apetite/metabolismo.
**Estratégias comportamentais e nutricionais complementares podem modular apetite e sintomas, mas exigem uso criterioso.**
- Ácido hidroxicítrico (Citrimax/Garcinia cambogia): 500 mg antes das refeições, especialmente meia hora antes do jantar para controle de fome no final da tarde.
- Óleos essenciais cítricos por inalação: três a cinco gotinhas, com instrução de inalação profunda; não ingeríveis e de custo elevado.
**Achados epidemiológicos sugerem papel das vitaminas do complexo B em comportamento, reforçando a importância da qualidade dietética.**
- Análise transversal (2012) relacionou baixa ingestão de B1, B2, B3, B5, B6 e folato a maiores escores de comportamento externalizante.

---

### Chunk 26/30
**Article:** Emagrecimento - Parte VII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.523

tégia inicial pode precisar focar no ganho de massa muscular para elevar o metabolismo, mesmo que isso signifique um ganho de peso inicial na balança.
*   **Riscos da Restrição Calórica Severa:** Dietas muito restritivas levam à queda do metabolismo basal, estagnação na perda de peso, frustração e alterações hormonais. A solução para quebrar esse ciclo é aumentar a massa muscular.
*   **Uso de Esteroides como Ferramenta Terapêutica:** Em certas fases da vida ou em casos de sarcopenia, estratégias hormonais (esteroides) podem ser vistas como "remédios" necessários para restaurar a funcionalidade muscular e permitir a evolução do paciente.
### 3. Exercício Físico e Estratégias Alimentares
*   **Musculação como Prioridade:** Para a maioria das pessoas com excesso de gordura e falta de músculo, a musculação é mais eficaz que o aeróbico, pois aumenta o gasto calórico residual e o metabolismo basal, prevenindo o efeito sanfona.

---

### Chunk 27/30
**Article:** Jejum Intermitente (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.520

 ão; 18 g mostraram tendência de melhora maior, sugerindo dose-resposta.
- Posologia prática: iniciar após pelo menos 12 horas de jejum; começar com 5 g de C8 e titrar até 15 g, com máximo sugerido de 20 g, para equilibrar benefício e tolerabilidade gastrointestinal.
**O horário das refeições importa: jantar cedo favorece o emagrecimento independentemente de dieta e atividade física.**
- Grupos que jantaram às 7h/7h30 perderam 30% mais peso que grupos que jantaram tarde (ex.: 10h30), indicando efeito temporal robusto.
- Este efeito reforça o alinhamento circadiano das janelas alimentares como componente crítico da eficácia do jejum.
**FMD integrada à quimioterapia tem protocolo preciso para tentar proteger células normais e potencialmente aumentar eficácia antitumoral.**
- Prescrição clínica: 5 dias de FMD (dieta que mimetiza o jejum) com quimioterapia no 3º dia; manter FMD por mais 2 dias após.

---

### Chunk 28/30
**Article:** Aula 01 Guilherme Sorrentino - Cirurgia e Medicina Funcional Integrativa (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.520

o intestino é parte da estratégia de cura. O objetivo clínico é abreviar o estado catabólico, fornecendo macro e micronutrientes (e, em casos selecionados, discutindo uso de hormônios anabólicos como testosterona) para proteger massa muscular e acelerar retorno à homeostase.

------------
## Fatores Adicionais de Risco: Coagulação e Hiperglicemia

A coagulação é mapeada com ferramentas como o score de Caprini, ainda que o cenário pós-pandemia tenha aumentado o risco de trombose por disfunção endotelial, exigindo atenção ampliada—incluindo homocisteína como fator trombogênico, com meta abaixo de 10. A hiperglicemia pré-operatória associa-se consistentemente a piores desfechos: além da inflamação vascular, forma produtos finais de glicação (AGEs) que alteram proteínas críticas da cicatrização (fibroblastos, colágeno, células imunológicas), promovendo excesso de oxidação e complicações como vasculopatia, retinopatia, neuropatia e nefropatia.

---

### Chunk 29/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.517

tos colaterais e riscos; priorizar ajustes de estilo de vida antes de escalonar.
- [ ] 14. Definir metas proteicas individuais e estratégias práticas para atingir ingestão adequada, especialmente em pacientes com saciedade elevada e esvaziamento gástrico lentificado.
- [ ] 15. Avaliar criticamente terapias adjuvantes para massa muscular (peptídeos/SARMs): não utilizar sem evidência robusta; acompanhar novos estudos.

---

## Quantitative Data

### Narrativa Quantitativa
O panorama integrado mostra que intervenções nutricionais estruturadas, como o jejum 5x2 com substituição de refeições, podem superar fármacos tradicionais em controle glicêmico e perda de peso no curto prazo, enquanto agonistas de GLP-1 (como tirzepatida) alcançam perdas de peso substanciais em longo prazo, porém com preocupações relevantes sobre perda de massa magra e reganho de peso sem mudanças de estilo de vida.

---

### Chunk 30/30
**Article:** Emagrecimento - Parte IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.516

 ões clínicas da alteração do microbioma; sinalizar nível de evidência.
> - Ajustes práticos: reduzir suplementos de BCAA, priorizar refeições com fibras e vegetais.
### 10. Proteína: metas, segurança renal e benefícios
- Metas diárias: ~1,2–1,6 g/kg favorecem composição corporal, emagrecimento, envelhecimento saudável e desempenho.
- A maioria não atinge as metas por padrão rico em farinha e proteína concentrada no almoço/jantar.
- Segurança renal: em geral, dietas ricas em proteína não são problema com função renal preservada; insuficiência renal grave requer cuidado especializado.
> Sugestões de IA
> - Quadro de conversão g/kg → porções/dia (ovos, carne, laticínios).
> - Planilha de 1 dia com 3–4 distribuições de proteína (café, almoço, lanche, jantar).
> - Delimitar quem não deve aumentar proteína sem supervisão (estágios de DRC).
> - Checklist de triagem renal (eGFR, albuminúria) antes de elevar proteína.
### 11.

---

