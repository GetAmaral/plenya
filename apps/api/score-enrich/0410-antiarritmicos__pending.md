# ScoreItem: Antiarrítmicos

**ID:** `c77cedd3-2800-7dec-9ee4-e8ed1335b026`
**FullName:** Antiarrítmicos (Histórico de doenças - Medicamentos)

**Preparation Metadata:**
- Quality Grade: **GOOD**
- Total Chunks: 30 de 20 artigos
- Avg Similarity: 0.564

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `c77cedd3-2800-7dec-9ee4-e8ed1335b026`.**

```json
{
  "score_item_id": "c77cedd3-2800-7dec-9ee4-e8ed1335b026",
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

**ScoreItem:** Antiarrítmicos (Histórico de doenças - Medicamentos)

**30 chunks de 20 artigos (avg similarity: 0.564)**

### Chunk 1/30
**Article:** TDAH - Parte XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.600

res de bomba de prótons e diuréticos) e sintomas comportamentais devem ser avaliados.
   - **Combinações Eficazes:** Bons resultados com magnésio, ômega-3, vitamina D e zinco.
* **Fisiopatologia Multifatorial do TDAH**
   - Fatores de risco: baixo peso ao nascer, exposição pré-natal a toxinas (álcool, nicotina, chumbo, pesticidas), questões educacionais e familiares.
   - Fatores maternos e genéticos: dieta materna, toxinas, polimorfismos genéticos e hereditariedade.
   - Fatores individuais e neurológicos: idade, gênero, status socioeconômico, predisposições de neurotransmissores, inflamação e desordens associadas (distúrbios do sono, depressão, ansiedade).
* **Crítica à Abordagem Médica Convencional**
   - O instrutor critica veementemente os médicos que negam a influência da alimentação no TDAH e em outras condições de saúde, classificando tal atitude como irresponsável e antiética.

---

### Chunk 2/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte VIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.590

cálcio iônico e ajustar vitamina D de forma personalizada, com reavaliação em 30–60 dias, especialmente após infecções/estresse.
- [ ] Investigar intolerância à histamina em pacientes com palpitações, arritmias, refluxo, gastrite, ansiedade/alterações do sono; correlacionar com dieta, microbiota e possíveis polimorfismos.
- [ ] Solicitar testes de micotoxinas urinárias em casos de sintomas persistentes sem explicação; revisar fontes alimentares (café, grãos, amendoim) e reforçar suporte hepático e intestinal.
- [ ] Mapear polimorfismos relevantes (ex.: FUT2) em pacientes com IBS ou baixa B12; personalizar dieta e estratégias para reforço da camada de muco.
- [ ] Implementar o protocolo começo–meio–fim: revisar dieta atual; avaliar digestão/absorção/transporte; checar suficiências de nutrientes-chave (vitamina D, C, A, zinco, selênio, ômega-3, B12) e corrigir.

---

### Chunk 3/30
**Article:** Ritmo Circadiano Eixo HPA - Parte VII (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.582

e fitoterápicos.
- **Adesão do Paciente:** Alguns pacientes têm dificuldade com o sabor dos sachês; orientar sobre a necessidade do tratamento é essencial.
> **Sugestões da IA**
> A seção sobre magnésio foi extremamente prática. A distinção diurno (malato) vs. noturno (treonato) é uma dica clínica valiosa. A tabela com as formas de magnésio é um recurso excelente. A discussão sobre formulação em sachês e adesão ("tem gente que é fresco demais") foi realista e divertida, conectando com os desafios do consultório. A organização foi impecável, da fisiopatologia à aplicação clínica.
### 5. Sugestão de Fórmula Básica de Vitaminas e Minerais
- **Componentes Sugeridos:** Tiamina, Riboflavina, Niacinamida, Ácido Pantotênico, Piridoxina (P5P como alternativa), Biotina (atenção à interferência no TSH), Metilfolato, B12, Magnésio (glicina, treonato, malato), Selênio, Manganês, Zinco, Cobre, Vitamina D e Vitamina K2/K7.

---

### Chunk 4/30
**Article:** TDAH - Parte XII (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.582

agnésio sérico e capilar mais baixos em indivíduos com TDAH.
    - Estudo de coorte (2010): Melhora de sintomas com a combinação de magnésio, ômega-3 e zinco.
    - Ensaio clínico randomizado (2021): Magnésio e Vitamina D melhoraram escores emocionais e sociais em TDAH.
> **Sugestões da IA**
> A compilação de estudos foi excelente. Como a tabela não foi exibida, destaque verbalmente um ou dois achados por estudo para fixar a relevância clínica. Ex.: “No estudo de 2017 nos EUA, o ponto-chave foi a rapidez do efeito: melhora em duas semanas, sugerindo impacto direto e rápido do magnésio.”
### 3. Mecanismos de Ação do Magnésio e a Relação com o Sono
- Modula a tirosina hidroxilase, enzima essencial para a síntese de dopamina a partir da tirosina.
- Atua como antagonista dos receptores NMDA, reduzindo a excitotoxicidade do glutamato.
- Reduz citocinas inflamatórias (IL-6 e TNF-alfa).
- Estabiliza a regulação do GABA, o ritmo circadiano e o eixo HPA.

---

### Chunk 5/30
**Article:** MFI - Psiquiatria 13 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.581

ico de uso de medicamentos psiquiátricos (o que usaram, como, por quanto tempo) e sintomas relacionados para mapear sistemas de neurotransmissores (glutamato, GABA, dopamina, etc.).
- [ ] Indicar terapias comportamentais aos pacientes com ansiedade, sugerindo troca de abordagem caso a terapia atual não esteja gerando resultados práticos.
- [ ] Recomendar uso de óleo de lavanda (ex.: 5 gotas na palma da mão para inalação, 3 vezes ao dia) como intervenção simples e eficaz para pacientes com ansiedade.
- [ ] Considerar associação de suplementos como magnésio, zinco, L-teanina, probióticos e adaptógenos no plano de tratamento da ansiedade.
- [ ] Investigar e abordar saúde do eixo intestino-cérebro em todos os pacientes com transtornos de humor, considerando neuroinflamação como fator causal.

---

### Chunk 6/30
**Article:** MFI - Psiquiatria Metabólica Funcional Integrativa - Aula 09 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.579

prescrever a pacientes em antidepressivos ou ansiolíticos devido a possíveis interações desconhecidas.
*   **Mucuna Pruriens**
    - Fitoterápico ayurvédico com L-Dopa (levodopa), precursor direto da dopamina que atravessa a barreira hematoencefálica.
    - L-Dopa é convertida em dopamina pela Dopa descarboxilase.
    - Estudos focam em doença de Parkinson; também investigada em Alzheimer, ELA e AVC por ação neuroprotetora.
    - O instrutor relata ausência de grandes resultados em uso pessoal.
*   **Selegilina**
    - Fármaco antigo, inibidor de MAO, usado em Parkinson e considerado nootrópico.
    - Inibe degradação de dopamina; combinação com fenilalanina melhorou escores de depressão em estudo.
    - Doses baixas (2–2,5 mg) podem auxiliar memória, foco e atenção, sem os efeitos colaterais ou restrições alimentares (queijos, cerveja) típicos de doses altas de IMAO.

---

### Chunk 7/30
**Article:** MFI - Psiquiatria 13 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.575

como Silexan) é apresentado como alternativa eficaz e segura aos benzodiazepínicos (como Lorazepam) para transtorno de ansiedade generalizada, conforme estudo clínico de 2010.
    - Outras estratégias: magnésio, probióticos, zinco, adaptógenos e L-teanina.
*   **Uso Correto de Medicamentos**
    - Medicamentos são úteis e devem ser prescritos quando necessário, especialmente em casos graves de depressão.
    - A medicação isolada raramente resolve a causa raiz; em casos graves pode “zumbificar” o paciente.
    - “Remédio bom é aquele que entra e sai”: remedeia a situação e depois é descontinuado.
## ❓ Perguntas
- [Inserir Pergunta/Confusão]
## 📚 Próximas Providências
- [ ] A partir de 19 de novembro de 2025, começar a perguntar aos pacientes sobre histórico de uso de medicamentos psiquiátricos (o que usaram, como, por quanto tempo) e sintomas relacionados para mapear sistemas de neurotransmissores (glutamato, GABA, dopamina, etc.).

---

### Chunk 8/30
**Article:** Psiquiatria Metabólica Funcional Integrativa AULA 20 (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.575

s/hipotireoidismo; considerar prova terapêutica de T3, integrada a outras intervenções.
* Arsenal terapêutico além de fármacos
   - Anti-inflamatórios, complexo B, análise de polimorfismos, dieta, exercício físico, regulação do eixo HPA, suporte mitocondrial e correção de micronutrientes essenciais para deiodinases (selênio, zinco, cobre, ferro).
* Educação e mudança de paradigma
   - Incentivar leitura de autores críticos (Kirsch, Frances) e divulgar evidências para colegas; reconhecer o desconforto em abandonar práticas consolidadas, priorizando qualidade de vida e resultados clínicos.
### 9. Aulas futuras e continuidade do módulo
* Próximos conteúdos
   - Aulas restantes do Dr. Frederico.
   - Aulas da “Ju” sobre tratamento da tireoide e tomada de decisão terapêutica (prova terapêutica).
   - Aula da Dra. Janaína sobre dieta cetogênica.
   - Continuidade sobre jejum intermitente e tireoide.

---

### Chunk 9/30
**Article:** Emagrecimento - Parte IX (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.570

valiar aporte e objetivos de médio prazo considerando dieta e adesão.
### 5. Hierarquia terapêutica, disbiose e pré-refeição
- Primeiro corrigir nutrientes essenciais e estratégia alimentar; depois fitoterápicos.
- Em obesos/sobrepeso, disbiose é comum: preferir berberina HCl antes das refeições; adicionar cromo, vanádio; considerar 5-HTP (25–50 mg) e L-teanina (200 mg) para ansiedade, equilibrando número de cápsulas.
- Canela do Ceilão: 1 colher de café no “shot” matinal ou café.
### 6. Evidências de fitoterápicos
- Gimnema silvestre: revisão sistemática e meta-análise (2021, 10 estudos, N=419) mostra redução de glicemias, HbA1c, TG e colesterol em T2DM; dose 200–300 mg antes das refeições.
- Ácido hidroxicítrico (HCA)/Citrimax: usar padronizado; efeitos em leptina e GLUT1/GLUT4; 500 mg antes das refeições; caro e aumenta cápsulas; melhor com B3, cromo e gimnema.

---

### Chunk 10/30
**Article:** MFI PÓS RACHEL GAIGER AULA 01 - OXIGÊNIO HIPERBÁRICO (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.569

cofatores mitocondriais (CoQ10, L-carnitina, ácido alfa-lipoico, glutationa, glicina, taurina, tirosina, MCT, vitaminas B1/B2/B3).
- Integração com terapias padrão e medicina funcional: redução de sintomas para permitir abordagem da causa raiz; sinergia com antibióticos/antifúngicos, esteroides em DII e reabilitação motora/cognitiva em casos neurológicos.
### Segurança, contraindicações e manejo de efeitos adversos
- Contraindicações relativas: doenças pulmonares crônicas (enfisema), pneumotórax recente, cirurgia cardíaca/trauma recente, neurite óptica, uso atual de bleomicina/doxorrubicina, otite média/dificuldade de equalização, implante coclear, marcapasso (avaliação individual), hipercapnia, hipertensão não controlada, transtornos convulsivos.
- Absolutas: pneumotórax, broncoespasmo agudo não resolvido, próteses orbitais específicas.
- Drogas incompatíveis: cisplatina, doxorrubicina; cautela com amiodarona e antiangiogênicos.

---

### Chunk 11/30
**Article:** Hipotireioidismo (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.568

gia (ajuste de T4 ± T3), checar conversão periférica e polimorfismos quando possível.
- [ ] 11. Considerar terapia combinada T4+T3 em candidatos: sintomáticos com LT4 adequado, tireoidectomizados, alta dose de LT4, autoimunidade ativa; seguir proporção 13:1–20:1 e fracionar T3.
- [ ] 12. Evitar T3 em gestantes, cardiopatas instáveis, malignidade ativa, psiquiatria não controlada; usar T3 isolado apenas em indicações específicas.
- [ ] 13. Mapear preferências/experiências dos pacientes com T3 e decidir compartilhadamente.
- [ ] 14. Triar e manejar SIBO/disbiose em hipotireoidismo (especialmente com constipação crônica), usando teste respiratório e terapias específicas; promover saúde do microbioma (“pool” intestinal).
- [ ] 15. Em cardiopatas (ICC, pós-IAM), avaliar T3 livre e segurança para eventual ajuste terapêutico (LT4 e, em casos selecionados, T3), monitorando Pro-BNP e PCR-us.
- [ ] 16.

---

### Chunk 12/30
**Article:** Hipertensão Arterial Sistêmica (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.568

vamente causas de hipertensão secundária, como apneia do sono (polissonografia) e disfunções da tireoide (TSH).
- [ ] 3. Implementar a mudança de estilo de vida como primeira linha de tratamento, focando em perda de peso, dieta low-carb anti-inflamatória e atividade física.
- [ ] 4. Evitar a prescrição de medicamentos para pacientes em pré-hipertensão (130-139 / 85-89 mmHg) sem comorbidades ou alto risco cardiovascular.
- [ ] 5. Educar os pacientes sobre a diferença entre sal integral e refinado, orientando o foco na eliminação de produtos industrializados e açúcar.
- [ ] 6. Implementar estratégias de suplementação (ex: alho envelhecido, beterraba, CoQ10, magnésio) como parte de uma abordagem integrativa, especialmente em pré-hipertensos ou hipertensos de baixo risco.
- [ ] 7. Revisar a medicação de pacientes que usam betabloqueadores como terapia primária, considerando a substituição por classes mais eficazes.
- [ ] 8.

---

### Chunk 13/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.566

co da dislipidemia.
- [ ] 2. Analisar criticamente os estudos sobre o Inclisiran (Cibrava), focando na diferença entre desfechos substitutos (redução de LDL) e desfechos clínicos duros (mortalidade, infarto, AVC).
- [ ] 3. Utilizar o site `the-nnt.com` para pesquisar o NNT e NNH de outros medicamentos prescritos na prática clínica.
- [ ] 4. Refletir sobre a influência da indústria farmacêutica e dos interesses financeiros na prescrição de novos medicamentos.
- [ ] 5. Estudar a diferença entre risco relativo e risco absoluto para interpretar criticamente os achados de estudos científicos.
- [ ] 6. Pesquisar sobre os diferentes tipos de ômega-3 (EPA, DHA, ALA) e suas formulações para entender como a qualidade do suplemento afeta os resultados.
- [ ] 7. Investigar os exames de subfracionamento de LDL e a relação ApoA/ApoB como ferramentas de avaliação de risco cardiovascular mais precisas que o LDL total.
- [ ] 8.

---

### Chunk 14/30
**Article:** Emagrecimento - Parte IX (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.565

[ ] 9. Decidir entre kelp (200–400 mcg) e Ascophyllum nodosum/I-plus (200–400 mg) para iodo, considerando potência, detox e carga de cápsulas.
- [ ] 10. Prescrever K2 (MK-7) 80–200 mcg com as refeições, especialmente quando suplementar vitamina D, exceto em usuários regulares de natto.
- [ ] 11. Em disbiose/hiperpermeabilidade, introduzir berberina HCl pré-refeição (250–500 mg) e considerar cromo e vanádio; avaliar 5-HTP (25–50 mg) e L-teanina (200 mg) para ansiedade, balanceando cápsulas.
- [ ] 12. Considerar gimnema silvestre 200–300 mg antes das refeições para suporte glicêmico e lipídico.
- [ ] 13. Avaliar custo-benefício do HCA (Citrimax) 500 mg antes das refeições; preferir sinergia com B3, cromo e gimnema; monitorar adesão.
- [ ] 14. Considerar ginostema: padronizar 80% de gipenosídeos (150–300 mg antes das refeições) ou actiponina 400 mg/dia; aplicar fator de correção e documentar.
- [ ] 15.

---

### Chunk 15/30
**Article:** Trato Gastrointestinal II- estômago – hipercloridria e hipocloridria (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.562

ou em uso crônico de IBP.
- [ ] 5. Reavaliar uso de IBP e antagonistas H2, ponderando riscos/benefícios e buscando estratégias não farmacológicas quando possível.
- [ ] 6. Considerar suporte com nutracêuticos e fitoterápicos apropriados (ex.: espinheira-santa), integrados ao plano alimentar, conforme avaliação individual.
- [ ] 7. Educar pacientes sobre mecanismos da hipocloridria e impactos sistêmicos, promovendo adesão a mudanças de hábitos.
- [ ] 8. Preparar para a próxima aula: coletar dados clínicos e laboratoriais para discussão de casos e estratégias de tratamento da hipocloridria.

---

## Teaching Note

Data e Hora: 2025-11-17 17:44:53
Local: [Inserir Local]
Aula: Medicina Funcional Integrativa - Sistema Gastrointestinal (Aula 2)
## Visão Geral
A aula abordou a hipocloridria, detalhando suas causas, sinais, sintomas e a importância do histórico alimentar. Foi feita uma análise crítica sobre o tratamento convencional do H.

---

### Chunk 16/30
**Article:** Hipotireioidismo (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.562

 de cardiovascular
- T3 modula canais iônicos, frequência/contratilidade, débito, vasorrelaxamento, SRAA, oxigenação, mitocôndria.
- Meta-análise (~2 milhões, 2017): hipotireoidismo aumenta mortalidade CV e geral.
- Hipotireoidismo subclínico: disfunção cardíaca leve reversível com T4.
- Baixo T3 em UTI/eventos agudos correlaciona com maior mortalidade; em ICC, menor conversão T4→T3, maior D3/rT3, citocinas; dobutamina aumenta T3 livre.
- T3 em baixa dose pós-IAM/ICC: melhora remodelamento, marcadores (Pro-BNP, PCR-us) e arritmias atriais, com segurança em protocolos selecionados.
### 20. Obesidade e eixo adipotireoidiano
- Obesidade: inflamação crônica, estresse oxidativo, disfunção metabólica.
- Hipotireoidismo franco: ganho de peso modesto (~2–3 kg atribuíveis à tireoide).
- TSH elevado pode ser consequência da adiposidade; leptina elevada influencia TSH e autoimunidade.

---

### Chunk 17/30
**Article:** Hipotireioidismo (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.561

idianos (inclusive eutireoidianos em tratamento).
- Hipotireoidismo altera motilidade, ácidos biliares, diversidade bacteriana; disfunções GI geram hiperpermeabilidade, inflamação, baixa absorção de nutrientes; impacto na autoimunidade.
- Microbioma afeta circulação entero-hepática de hormônios, biodisponibilidade de levotiroxina e metabolismo de antitireoidianos; atua como “tanque reserva” de conjugados.
- Relevância clínica: constipação crônica, hipocloridria, intolerâncias alimentares.
### 18. Paradigmas e limitações do TSH
- TSH isolado insuficiente; conversão T4→T3 é variável; valores populacionais não refletem set point individual.
- Fluxos de decisão devem incluir T3 livre, T4 livre, rT3 e anticorpos quando apropriado.
### 19. Função tireoidiana e saúde cardiovascular
- T3 modula canais iônicos, frequência/contratilidade, débito, vasorrelaxamento, SRAA, oxigenação, mitocôndria.

---

### Chunk 18/30
**Article:** TDAH - Parte XXVI (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.561

gação e correção de causas subjacentes — dieta, sono, exercício e saúde intestinal — antes do uso de medicação. O objetivo é promover um cuidado mais responsável e consciente, evitando a patologização da infância e a dependência de soluções farmacológicas.
## 🔖 Pontos de Conhecimento
### 1. Efeitos Adversos de Medicamentos Estimulantes para TDAH
*   **Efeitos Cardiovasculares**
    - Estimulantes aumentam, em média, a frequência cardíaca em 12 bpm e a pressão arterial (sistólica/diastólica) de 1 a 4 mmHg, geralmente sem relevância clínica.
    - Em 5% a 15% dos indivíduos, podem ocorrer elevações mais significativas de pressão e frequência cardíaca, exigindo monitoramento médico.
    - Recomenda-se histórico cardíaco individual e familiar abrangente antes do tratamento. Havendo fatores de risco (morte súbita na família, condução cardíaca anormal, anormalidade estrutural), indicar avaliação adicional.

---

### Chunk 19/30
**Article:** Psiquiatria Metabólica Funcional Integrativa AULA 20 (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.556

dor crônica, exposição a toxinas) e implementar estratégias para reduzi-los.
- [ ] 4. Integrar suporte metabólico: metilfolato, complexo B, NAC, L-acetilcarnitina, ácido alfa-lipoico, triptofano, zinco, magnésio, ômega-3 e CoQ10, conforme perfil do paciente.
- [ ] 5. Avaliar e modular eixo HPA: rotina de sono, manejo de estresse, exercício físico regular e intervenções de estilo de vida.
- [ ] 6. Investigar sinais de desbiose e leaky gut; considerar estratégias para saúde intestinal e barreira hematoencefálica.
- [ ] 7. Ler e discutir com a equipe os trabalhos de Irving Kirsch e Allen Frances; revisar dados do STAR*D e do painel 2005–2015 sobre antidepressivos e qualidade de vida.
- [ ] 8. Planejar acompanhamento estruturado de sintomas e biomarcadores durante a introdução de T3, com monitorização de efeitos e ajuste de doses.
- [ ] 9. Preparar-se para conteúdos futuros: estudar diretrizes sobre dieta cetogênica (Dra.

---

### Chunk 20/30
**Article:** Genética e Epigenética I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.555

adequados.
- [ ] 15. Considerar suplementação de colina (incluindo gestantes) e TMG como suporte ao ciclo de um carbono; evitar confundir com betaína HCl.
- [ ] 16. Planejar intervenções de curto prazo perceptíveis (ex.: manejo de ansiedade) enquanto estrutura modulações epigenéticas de longo prazo.
- [ ] 17. Mapear pacientes autoimunes e coordenar cuidado com reumatologista funcional integrativo; evitar retirada súbita de medicações.
- [ ] 18. Identificar pacientes com consumo elevado de café (>5/dia) e oferecer plano de redução gradual.
- [ ] 19. Orientar redução/cessação de álcool e seus riscos; evitar “remendos” pós-excesso.
- [ ] 20. Triar usuárias de anticoncepcional para possível deficiência de B9, B6, B12; planejar suporte nutricional/suplementação.
- [ ] 21. Auditar complexos B com ácido fólico em doses altas; racionalizar escolhas conforme necessidade e condição financeira.
- [ ] 22.

---

### Chunk 21/30
**Article:** Cardiologia I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.554

conforme PA e função renal.
## Conteúdo a Cobrir (Restante)
1. Revisão aprofundada de colesterol (ex.: “The Great Cholesterol Myth”, “The Cholesterol Myths and The Sulfics”).
2. Arritmias e suas abordagens dentro da medicina funcional integrativa.
3. Aprofundamentos por especialidade (cardiologia avançada, gastroenterologia, psiquiatria, neurologia, reumatologia).
4. Detalhamento prático de reposição hormonal na prevenção cardiovascular e manejo da resistência insulínica.
5. Protocolo de suplementação (ex.: ômega-3) com critérios e dosagens.
6. Estratégias estruturadas de emagrecimento específicas para risco cardiovascular.
7. Continuação sobre LDL: estratégias dietéticas e clínicas para modulação de subtipos de LDL.
8. Questionamento detalhado de estudos recentes sobre colesterol e aterosclerose.
9. Protocolos práticos de ajuste de medicação (estatinas) e alimentação individualizada com base em exames.

---

### Chunk 22/30
**Article:** Aula Jéssica Marques - Neurologia Funcional Integrativa 3 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.554

agnésio, vitamina B12, folato (B9), vitamina D e ferritina/ferro; corrigir deficiências conforme resultados.
- [ ] 3. Planejar introdução gradual de suplementação: iniciar uma formulação, aguardar 2–3 dias, adicionar a segunda e, após 2–3 dias, a terceira, monitorando efeitos colaterais.
- [ ] 4. Implementar ciclos de 60–90 dias entre plantas com mecanismos semelhantes, trocando para outra família após cada ciclo para sensibilizar diferentes receptores.
- [ ] 5. Selecionar extratos padronizados com maior biodisponibilidade (ex.: Curcumin C3, Cureit, CurcuVail) e evidência clínica para compor formulações.
- [ ] 6. Incorporar vias alternativas para idosos: avaliar uso de injetáveis, transdérmicos, fotobiomodulação, aromaterapia e tinturas conforme perfil e poder aquisitivo.
- [ ] 7. Estruturar rotina diária do paciente com atividades úteis e significativas (cozinhar, organizar casa, acompanhar netos/filhos), promovendo pertencimento e utilidade.
- [ ] 8.

---

### Chunk 23/30
**Article:** Ritmo Circadiano Eixo HPA - Parte IX (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.551

a livre, agravando falta de energia e libido em disfunção do eixo HPA.
    - Associados a depressão, pânico e suicídio.
*   **Impacto na Função Tiroideia e Nutricional**
    - Podem aumentar T3 reverso (forma inativa que bloqueia recetores de T3), reduzindo metabolismo basal.
    - Em disfunção do eixo HPA, com conversão T4→T3 já reduzida, o T3 reverso agrava o quadro metabólico.
    - Diminuem absorção de folato, vitamina B12 e vitamina B6.
## ❓ Perguntas
- [Inserir Pergunta/Dúvida]
## 📚 Próximos Passos
- [ ] Solicitar curva de cortisol salivar em pacientes com fadiga extrema e sintomas sugestivos para avaliar o eixo HPA.
- [ ] Em curvas “flat”, considerar hidrocortisona em baixas doses (ex.: 10 mg manhã, 5 mg tarde) como terapia de curto prazo, com monitorização e revisão em 2–4 meses; reduzir e retirar conforme melhoria.

---

### Chunk 24/30
**Article:** Hipotireioidismo (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.550

cias T4/T3
- Pêndulo histórico: clínica→laboratório→individualização com múltiplos marcadores.
- Meta-análises até 2006 sem benefício claro da combinação; guideline europeu (2012) reconhece possíveis benefícios.
- Endocrine Reviews 2022: orientações práticas ainda baseadas em TSH, com reconhecimento de limitações.
- Futuro: incorporar biomarcadores teciduais, genéticos (polimorfismos de deiodinases/receptor TR), metabolômica.
### 10. Prática clínica: ajuste de T4, horários e absorção
- TSH permanece útil para ajustes percentuais, interpretado com clínica e outros marcadores.
- Tomada: manhã em jejum ou à noite (≥2 h após refeição); bedtime pode melhorar TSH/T3 em alguns.
- Absorção: depende de acidez gástrica; IBP/hipocloridria reduzem biodisponibilidade (usuários de IBP precisam ~37% mais dose).

---

### Chunk 25/30
**Article:** TDAH - Parte XXVI (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.550

antes do tratamento. Havendo fatores de risco (morte súbita na família, condução cardíaca anormal, anormalidade estrutural), indicar avaliação adicional.
    - Publicação no *Journal of the American College of Cardiology* (JACC) afirma que fármacos para TDAH são potentes estimulantes do SNC, associados a eventos cardiovasculares adversos, devendo ser prescritos após opções mais seguras (exercício, ômega 3).
    - Estudo longitudinal de 14 anos no *JAMA Psychiatry* (2024) sugere que uso prolongado de medicamentos para TDAH está associado a maior risco de doenças cardiovasculares, especialmente hipertensão e doença arterial, mais evidente com estimulantes e em doses altas.
*   **Efeitos Adversos Psiquiátricos**
    - Adolescentes podem relatar maior retraimento social e sentimento de inibição.
    - Irritabilidade sustentada ou aumento da ansiedade podem indicar ajuste de dose ou troca de medicação (ex: de Venvanse para atomoxetina).

---

### Chunk 26/30
**Article:** Base Metabólica das Doenças Crônicas e Gerenciamento - Oxidação 2 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.549

ados.
- **Funções:** Tratamento de osteoporose, anemia hipocrômica, prevenção de doenças cardiovasculares.
### 3. Suplementação de Magnésio
- **Fontes alimentares:** Sementes, leguminosas, folhas verdes escuras. A quantidade nos alimentos é pequena devido ao solo brasileiro pobre em magnésio.
- **Posologia:** Idealmente suplementar de manhã e à noite.
- **Funções:** Mais de 300 funções enzimáticas, melhora do humor, insônia, depressão, cãibras, estresse, enxaqueca, peristaltismo intestinal, sensibilidade à insulina.
- **Formas de suplementação:** A escolha depende da queixa do paciente (ex: citrato para intestino, treonato para memória). A quantidade de magnésio elementar varia conforme a forma (ex: 500 mg de magnésio glicina contêm 150 mg de magnésio elementar).
- **Óxido de magnésio:** Alta biodisponibilidade, mas com efeito antiácido geralmente indesejado.
### 4.

---

### Chunk 27/30
**Article:** Hipotireioidismo (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.548

a pacientes com sintomas persistentes, especialmente aqueles com polimorfismos genéticos (12-14% da população), tireoidectomizados (que perdem 10-20% da produção de T3) ou com doses de T4 acima de 1.2 mcg/kg.
**Achados Adicionais**
- Uma meta-análise de 2017 com 2 milhões de participantes mostrou que o hipotireoidismo é um fator de risco independente para mortalidade cardiovascular.
- Em um estudo com 21 mulheres inférteis com TSH entre 0,5 e 3,5, a otimização da dose de T4 para melhorar o T3 livre resultou em todas engravidando em três meses.
- A levotiroxina foi a segunda droga mais vendida nos EUA em 2019.
- Um estudo de 2001 mostrou que doses suprafisiológicas de hormônio tireoidiano (200-300 microgramas) aliviaram sintomas em pacientes com fibromialgia, uma condição onde 35% podem ter resistência periférica ao hormônio tireoidiano.

---

### Chunk 28/30
**Article:** Aula Jéssica Marques - Neurologia Funcional Integrativa 3 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.548

-teanina, Huperzine, Ginseng.
- Adaptação individual
  - Ajustar doses e frequência conforme resposta; introduzir um composto por vez.
### 9. Caso Prático e Abordagem Integrativa
- Aromaterapia e dieta
  - Óleo de gergelim com óleos essenciais, caldo enriquecido com colágeno; respeitar paladar e otimizar dieta para funcionalidade.
- Continuidade terapêutica
  - Uso de fitoterápicos, suplementos e, em próxima sessão, óleo de cannabis para otimização neurológica.
## ❓ Perguntas
- [Inserir Pergunta/Confusão]
## 📚 Tarefas
- [ ] 1. Implementar rastreio precoce em pacientes com queixas sutis (humor, sono, preferência por doces), incluindo PET-CT/FDG PET, ressonância funcional e biomarcadores no líquor quando indicado.
- [ ] 2. Solicitar exames laboratoriais para avaliar magnésio, vitamina B12, folato (B9), vitamina D e ferritina/ferro; corrigir deficiências conforme resultados.
- [ ] 3.

---

### Chunk 29/30
**Article:** MFI - Psiquiatria Metabólica Funcional Integrativa - Aula 05 (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.547

Revisão sistemática (2021) em pacientes de cirurgia cardíaca aberta: recomendada suplementação oral para reduzir ansiedade/depressão e melhorar sono no pós-operatório.
     - Revisões/meta-análises em desordens neurológicas: enxaqueca (31 revisões, 2 meta-análises), depressão (15 revisões, 2 meta-análises), epilepsia (3 revisões, 1 meta-análise), dor crônica (5 revisões), ansiedade (1 meta-análise, 8 revisões), AVC (22 revisões, 6 meta-análises), Alzheimer e Parkinson.
   - Formas e doses práticas:
     - Magnésio treonato favorece passagem hematoencefálica; iniciar em 500 mg a 1 g/dia de treonato.
     - Combinações: treonato 500 mg + glicina 200 mg + malato 250 mg para suporte mitocondrial e modulação com glicina.
     - Faixa geral de magnésio total: 500 mg a 2 g/dia, ajustando à tolerância.

---

### Chunk 30/30
**Article:** Ácidos Graxos Saturados de Cadeia Curta (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.547

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

