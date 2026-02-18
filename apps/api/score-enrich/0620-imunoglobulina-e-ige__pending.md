# ScoreItem: Imunoglobulina E (IgE)

**ID:** `c77cedd3-2800-797c-864e-c7f5c4b78ae1`
**FullName:** Imunoglobulina E (IgE) (Exames - Laboratoriais)
**Unit:** IU/mL

**Preparation Metadata:**
- Quality Grade: **GOOD**
- Total Chunks: 30 de 11 artigos
- Avg Similarity: 0.552

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `c77cedd3-2800-797c-864e-c7f5c4b78ae1`.**

```json
{
  "score_item_id": "c77cedd3-2800-797c-864e-c7f5c4b78ae1",
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

**ScoreItem:** Imunoglobulina E (IgE) (Exames - Laboratoriais)
**Unidade:** IU/mL

**30 chunks de 11 artigos (avg similarity: 0.552)**

### Chunk 1/30
**Article:** Intolerâncias, Alergias e Hipersensibilidades Alimentares II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.606

- **Marcadores Fecais:**
        - **Calprotectina Fecal:** Valor < 138 exclui alergia não mediada por IgE com alta sensibilidade.
        - **Pesquisa de Sangue Oculto:** Exame simples e sensível para enterocolite.
        - **Alfa-1 Antitripsina Fecal:** Avalia enteropatia perdedora de proteínas.
*   **Testes Específicos e Procedimentos:**
    - **Testes Cutâneos (com alergista):** Prick Test (padrão-ouro para alergia mediada por IgE) e Patch Test (para reações tardias).
    - **Diagnóstico Molecular (RAST, ImunoCAP):** Avalia IgE específica para determinados alérgenos.
    - **Teste de Provocação Oral:** Considerado padrão-ouro para confirmação, mas é arriscado e complexo.
    - **Testes de IgG:** Não devem ser usados de rotina para diagnóstico de alergia, pois podem indicar tolerância e não alergia.
    - **Endoscopia/Colonoscopia:** Podem revelar achados como hiperplasia nodular linfoide.
### 5.

---

### Chunk 2/30
**Article:** Intolerâncias, Alergias e Hipersensibilidades Alimentares II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.584

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

### Chunk 3/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XVI (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.575

ia a diversos alimentos, guiando uma dieta de exclusão. Estudos mostram alta prevalência de IgG positivo em pacientes com urticária, eczema e dermatite. É uma ferramenta pedagógica poderosa para motivar o paciente.
    *   **Teste de Atividade da DAO:** Avalia a capacidade de degradar a histamina.
    *   **Teste de Intolerância à Lactose:** Identifica a má digestão do açúcar do leite.
*   **Estratégia de Tratamento Personalizado:**
    *   Baseia-se na identificação da causa (intolerância à lactose, histamina, reação IgG).
    *   O foco principal é sempre melhorar o bioma intestinal para aumentar a tolerância futura aos alimentos.
    *   Uma dieta de eliminação baseada no teste de IgG mostra alta eficácia, com melhora significativa em quadros de erupção cutânea, prurido, asma, enxaqueca e congestão nasal.

---

### Chunk 4/30
**Article:** Intolerâncias, Alergias e Hipersensibilidades Alimentares II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.569

cruzada entre alérgenos inalantes (pólen) e alimentares (ex: pólen e maçã; látex e banana/kiwi).
### 4. Diagnóstico da Alergia Alimentar
*   **Anamnese e Exame Físico:** A anamnese detalhada é fundamental (história familiar, parto, aleitamento). No exame físico, a dor à palpação da fossa ilíaca direita pode indicar inflamação nas placas de Peyer.
*   **Provas Diagnósticas Laboratoriais:**
    - **Hemograma:** Pode mostrar eosinofilia ou plaquetas elevadas (>400 mil), sugerindo inflamação.
    - **Imunoglobulinas:** IgE aumentada indica reações tipo 1. IgG4 pode estar aumentada na esofagite eosinofílica.
    - **Fenotipagem Linfocitária (CD4/CD8):** Ajuda a diferenciar reações. Relação CD4/CD8 > 3 sugere reação TH2 (humoral); < 1 sugere reação TH1 (celular).
    - **Marcadores Fecais:**
        - **Calprotectina Fecal:** Valor < 138 exclui alergia não mediada por IgE com alta sensibilidade.

---

### Chunk 5/30
**Article:** Intolerâncias, Alergias e Hipersensibilidades Alimentares II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.569

EUA, 90% das alergias alimentares se concentram em leite, trigo, soja, peixes, mariscos, ovos, amendoim, nozes e gergelim, sustentando o foco clínico em grupos alimentares-chave.
- Risco familiar: 30% de chance na criança quando um dos pais é alérgico; 80% quando ambos os pais têm alergia alimentar, reforçando a anamnese como ferramenta crítica.
**O mecanismo imunológico define tempo de resposta e direciona o diagnóstico: IgE (tipo 1) é imediato, enquanto tipos 2–4 são celulares/tardios.**
- Hipersensibilidade tipo 1 (IgE mediada): urticária, angioedema, broncoespasmo, asma, anafilaxia e síndrome alérgica oral; resposta clínica rápida entre 15 minutos e poucas horas (limite inferior 15 minutos; típicos superiores de 30 minutos a 3–4 horas, podendo se estender).
- Hipersensibilidade tipo 2: mecanismo celular não IgE, exemplificado por síndrome da enterocolite induzida por proteína da dieta e proctocolite; diferenciação crucial para manejo.

---

### Chunk 6/30
**Article:** Intolerâncias, Alergias e Hipersensibilidades Alimentares II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.567

o: favorece alergia alimentar celular (perfil TH1).
- CD8 muito baixo: deficiência de tolerância imunológica.
- CD4 aumentado: alergias tipo I (humoral).
- CD19: associado ao componente humoral; CD56: inflamação na parede do TGI.
- Relações: CD4/CD8 >3 aponta para TH2; <1 para TH1; >4 associa falha de tolerância e alergia a múltiplos alimentos.
- CD59: estimulação de linfócito B e produção de IgE.
### 16. Testes cutâneos e sorológicos para alergia tipo IgE
- Teste cutâneo (prick): padrão-ouro para alergia alimentar mediada por IgE; identifica alérgenos alimentares, inalantes e de contato.
- Patch test: útil em crianças pequenas para reações de contato/tardias.
- Teste de provocação oral duplo-cego: padrão-ouro; caro, inconveniente, com risco de reações graves.
- RAST/ImmunoCAP: pedir preferencialmente com IgE elevada; diagnósticos moleculares específicos por alimento; útil também para alérgenos respiratórios.

---

### Chunk 7/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XVI (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.567

- Teste de microbioma intestinal.
    - Teste de tolerância à lactose.
    - Avaliação de resistência à insulina.
    - Avaliação hormonal: diidrotestosterona (DHT), testosterona, SHBG e metabolômica hormonal (metabólitos urinários).
    - Marcadores inflamatórios sistêmicos e avaliação do eixo HPA (estresse).
- **Resultados de Estudos Mencionados:**
    - Um estudo sobre dietas de eliminação baseadas em testes de IgG mostrou melhorias significativas em condições como erupção cutânea, prurido, asma, zumbido, enxaqueca e congestão nasal.
- **Exemplo de Teste de IgG:** Mostrou reatividade (classe 3 ou 4) a alimentos como farelo de aveia, abacaxi, pêssego e leite de vaca.
## Diagnóstico Primário:
- **Avaliação:** O transcrito é uma palestra médica focada na interconexão entre dermatologia, nutrição e saúde metabólica.

---

### Chunk 8/30
**Article:** TDAH - Parte XXI (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.566

térias produtoras de histamina.
    - Testes de reações alimentares, como o teste de IgG, para guiar uma dieta de eliminação.
    - Avaliação de polimorfismos genéticos relacionados ao metabolismo da histamina (ex: HNMT, DAO) e à via da dopamina (ex: DAT1).
    - Avaliação de marcadores inflamatórios e nutrientes.
- **Plano de Tratamento de Acompanhamento:**
    - Implementar uma dieta saudável, rica em frutas e vegetais ("comida de verdade"), e eliminar alimentos processados, corantes e conservantes artificiais.
    - Manipular a microbiota intestinal através de dieta, probióticos e prebióticos com base nos resultados dos testes.
    - Evitar estritamente os antígenos alimentares identificados para pacientes com alergias ou sensibilidades.
    - Considerar a suplementação com cofatores para as vias de degradação da histamina (vitamina B6, vitamina C, cobre) e potencialmente a enzima DAO para intolerância à histamina.

---

### Chunk 9/30
**Article:** pediatria funcional integrativa - parte IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.565

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

### Chunk 10/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.562

desfavorável (poluição, mofo, poeira, pelos, químicos) e interno inflamado (alergias, intoxicações); comorbidades (alergia alimentar, refluxo, obesidade, anemia, rinite).
### 2. Conceitos diagnósticos e fisiopatologia segundo GINA
* Definição e sintomas
  - Asma: inflamação crônica das vias aéreas inferiores com obstrução reversível; sintomas: tosse, sibilância, aperto torácico, dispneia; padrão persistente (leve a grave) ou intermitente.
  - Desencadeantes: IgE (frequente: viroses) vs fenótipo neutrofílico (sem desencadeante claro).
* Confirmação diagnóstica
  - Limitação do fluxo (VEF1 e VEF1/CVF: <0,8 adulto; <0,9 criança).
  - Variabilidade: broncodilatador (≥12% e ≥200 ml adulto; ≥12% criança); PFE 2x/dia por 2 semanas (>10% adulto; >13% criança); resposta a tratamento; testes de desafio.
* Acompanhamento do controle
  - ACT (5 itens; 5–25 pontos) nas versões pediátrica e adulta.

---

### Chunk 11/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.558

intestinal.
-   **SIFO:** Sensação de "blurring" (efeito baiacu - sentir-se distendido mesmo sem distensão objetiva) e dor abdominal.
## Objetivo:
A palestra descreve os exames e achados objetivos para a avaliação de pacientes com suspeita de Síndrome do Intestino Irritável (SII), embora não contenha achados de exame físico de um paciente específico.
-   **Exames Laboratoriais Gerais:** Hemograma e marcadores de atividade inflamatória.
-   **Calprotectina Fecal:** Usada para descartar doença inflamatória intestinal. Um valor abaixo de 100 µg/g tem uma positividade de 98% para o diagnóstico de SII. Entre 100-250 µg/g é uma zona cinzenta. Acima de 250 µg/g requer colonoscopia.
-   **Avaliação para Doença Celíaca:** Dosagem de IgA sérica total e anticorpo antitransglutaminase IgA. Se a IgA for baixa, dosar IgG sérica e antitransglutaminase IgG.
-   **Teste Parasitológico de Fezes:** Recomendado no contexto do Brasil.

---

### Chunk 12/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.556

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

### Chunk 13/30
**Article:** Intolerâncias, Alergias e Hipersensibilidades Alimentares II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.553

:** Mediada por linfócitos T, sem anticorpos.
*   **Manifestações Clínicas:**
    - São variáveis e podem afetar múltiplos sistemas.
    - **Pele:** Prurido, urticária, angioedema, dermatite atópica (mais comuns).
    - **Gastrointestinais:** Refluxo, vômitos, dor abdominal, constipação, diarreia, sangramento oculto.
    - **Respiratórias:** Broncoespasmo, coriza, tosse.
    - **Neurológicas:** Hiperatividade e déficit de atenção.
    - **Outros:** Palidez sem anemia, aftas, língua geográfica.
*   **História Natural:** Alergias a leite e ovos em crianças tendem a desaparecer, enquanto alergias a amendoim, nozes e frutos do mar costumam persistir.
*   **Síndrome da Alergia Oral:** Comum em adultos, com sintomas na orofaringe (coceira, queimação) devido à reatividade cruzada entre alérgenos inalantes (pólen) e alimentares (ex: pólen e maçã; látex e banana/kiwi).
### 4.

---

### Chunk 14/30
**Article:** SAM (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.548

graves):** Omalizumab.
    -   **Inibidores de mastócitos (para mastocitose sistêmica/leucemia mastocítica):** Substâncias específicas não detalhadas.
-   **Próximos Passos/Exame:**
    -   O tratamento deve ser individualizado, seguindo o princípio "comece baixo, vá devagar, mas vá" ("Start low, go slow, but go/grow").
    -   Identificar e eliminar gatilhos, como poluentes ambientais, produtos cosméticos e micotoxinas.
    -   Avaliar a microbiota para disbiose ou supercrescimento bacteriano.
    -   Se o médico não se sentir confortável para tratar, encaminhar o paciente a um especialista.
-   **Plano de Tratamento de Acompanhamento:**
    -   O tratamento é proposto mesmo sem todos os critérios diagnósticos validados, utilizando o teste terapêutico como parte do diagnóstico.
    -   Aumentar as doses dos medicamentos (bloqueadores H1/H2, estabilizadores de mastócitos) até quatro vezes a dose padrão, se necessário, para controle dos sintomas.

---

### Chunk 15/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XVI (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.548

, prurido, asma, enxaqueca e congestão nasal.
*   **Importância da Reposição Hormonal:** Na menopausa, a terapia de reposição hormonal melhora drasticamente a qualidade e a evolução dos tecidos cutâneos, algo que tratamentos externos isolados não conseguem resolver.
*   **Cuidados na Prática:**
    *   A exclusão de lácteos em crianças pequenas deve ser feita por profissionais especializados para garantir a ingestão adequada de cálcio.
    *   Os resultados da dieta de eliminação podem ser potencializados com suplementação, probióticos e fibras para modular o bioma intestinal.
## ❓ Perguntas
- [Inserir Pergunta/Dúvida/Confusão]
## 📚 Tarefas
- [ ] 1. Adotar uma anamnese completa e integrativa, investigando todos os sistemas do corpo (digestivo, hormonal, sono, etc.), mesmo em queixas dermatológicas.
- [ ] 2.

---

### Chunk 16/30
**Article:** MFI - SÍNDROME PÓS COVID - AULA 02 (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.546

*   **Respiratórios:** Rinorreia, congestão nasal, dispneia.
    *   **Neurológicos:** Dores de cabeça, *brain fog*.
    *   **Cardíacos:** Taquicardia, palpitações.
    *   **Gastrointestinais:** Dores abdominais, diarreia, constipação, náuseas.
    *   **Cutâneos:** Urticária, rubor, eczema.

**Diagnóstico e Tratamento:**
*   A suspeita deve ser levantada em pacientes com histórico de alergias ou quadros clínicos muito vastos.
*   **Diagnóstico:**
    1.  **Dosagem de metil-histamina** em urina de 24 horas.
    2.  **Análise da atividade da enzima DAO** (disponível no exame Copromax, que também avalia o *leaky gut*).
*   **Tratamento:**
    1.  **Dieta anti-histamínica:** Restringir por um mês alimentos ricos em histamina (queijos, fermentados), liberadores de histamina ou inibidores da DAO.
    2.  **Medicação:** O uso do anti-histamínico E-Bastel (10 mg, duas vezes ao dia por um mês, seguido de uma vez ao dia por mais um mês) pode ser uma estratégia.

---

### Chunk 17/30
**Article:** Intolerâncias, Alergias e Hipersensibilidades Alimentares II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.543

RAST/ImmunoCAP: pedir preferencialmente com IgE elevada; diagnósticos moleculares específicos por alimento; útil também para alérgenos respiratórios.
- Teste de ativação de basófilos: específico para tipo I; disponível em poucos laboratórios.
- Testes IgG: não solicitar de rotina; positividade pode refletir tolerância; utilidade restrita (p.ex., IgG4 na esofagite eosinofílica).
### 17. Exames complementares gastrointestinais e marcadores fecais
- Endoscopia/colonoscopia com biópsia: avaliação de mucosa e inflamação.
- Alfa-1 antitripsina fecal: marcador de enteropatia perdedora de proteínas.
- Cintilografia/nucleares: avaliação de esvaziamento gástrico.
- Calprotectina fecal: marcador inflamatório gastrointestinal; valor <138 pode excluir alergia não mediada por IgE com alta sensibilidade (~95%) e especificidade moderada (~78,6%).

---

### Chunk 18/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa III (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.543

minase IgA. Se a IgA for baixa, dosar IgG sérica e antitransglutaminase IgG.
-   **Teste Parasitológico de Fezes:** Recomendado no contexto do Brasil.
-   **Avaliação da Permeabilidade Intestinal:** O aumento da permeabilidade (leaky gut) pode ser avaliado pela zonulina (fecal ou sérica). Menciona-se que o estresse (injeção de CRH) pode induzir um aumento nos marcadores de leaky gut.
-   **Avaliação da Microbiota/Metabolômica:** A avaliação isolada da microbiota é considerada de pouco valor. A avaliação da metabolômica (ex: ácidos orgânicos urinários) é mais útil para avaliar a função da microbiota e detectar metabólitos bacterianos e fúngicos. O aumento do D-lactato no sangue pode estar associado ao uso de probióticos e causar "brain fogginess".
-   **Teste Respiratório para SIBO/IMO:** Considerado o método prático padrão, utilizando lactulose ou glicose.
    -   **Critério para SIBO (Hidrogênio):** Elevação acima de 20 ppm nos primeiros 90 minutos.

---

### Chunk 19/30
**Article:** TDAH - Parte XXII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.543

ca e marcadores de hipersensibilidade (IgE, IgG).
   - Entre 184 pacientes com depressão, 66 apresentaram IgE alta (~36%); no grupo controle foram 42, indicando diferença significativa entre grupos.
   - Concentrações médias de IgE: depressão ~49 (quase 50), não depressivos 31,63.
   - Interpretação: hipersensibilidade tipo 1 mediada por IgE pode contribuir para níveis elevados de histamina sérica em adolescentes com depressão.
* DAO (Diamino Oxidase) e intolerância à histamina
   - A atividade de DAO pode se apresentar aumentada como mecanismo compensatório frente ao excesso de histamina advindo do intestino e da alimentação, não por deficiência de produção.
   - Sintomas de histamina com “DAO aumentada” indicam compensação: excreção elevada de histamina (inclusive urinária) e sintomas compatíveis, sem necessariamente deficiência intrínseca de DAO.

---

### Chunk 20/30
**Article:** SAM (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.542

io:** Congestão nasal, espirros, tosse, chiado no peito, dificuldade respiratória.
-   **Cardiovascular:** Taquicardia, hipotensão, síncope.
-   **Neuropsiquiátrico:** Dor de cabeça, confusão mental ("brain fog"), ansiedade, depressão.
-   **Sistêmico:** Anafilaxia, fadiga, dores generalizadas.
As reações podem ser imediatas (segundos a minutos), como na anafilaxia, ou tardias (horas depois da exposição).
## Objetivo:
O diagnóstico é complexo e multifatorial, sem um único teste definitivo. A abordagem diagnóstica inclui:
1.  **Clínica:** Presença de sintomas recorrentes e episódicos em pelo menos dois dos seguintes sistemas: pele, gastrointestinal, respiratório e cardiovascular.
2.  **Marcadores Laboratoriais:**
    -   **Triptase sérica:** Considerado o marcador padrão. O diagnóstico é sugerido por um aumento de 20% + 2 ng/mL acima do valor basal do paciente durante uma crise. No entanto, o palestrante relata que raramente vê resultados positivos.

---

### Chunk 21/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XVI (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.541

eia.
- **Sintomas Neurológicos/Gerais:** Dores de cabeça (relacionadas à sinusite), enxaquecas (migraine), zumbido, fadiga após comer, fadiga crônica.
- **Sintomas de Intolerância:** Coceira após consumir alimentos ricos em histamina (laticínios, pimentão, berinjela, abacate), sintomas de intolerância à lactose.
## Objetivo:
O transcrito é uma palestra médica e não contém os exames de um paciente específico. Discute vários exames e achados objetivos para diagnosticar as causas subjacentes de condições dermatológicas e sistêmicas:
- **Testes Laboratoriais Sugeridos:**
    - Teste de IgG para alimentos para avaliar reações tardias (menciona laboratórios como SYNLAB e Testify).
    - Teste de atividade da DAO (diamina oxidase) para avaliar a intolerância à histamina.
    - Teste de microbioma intestinal.
    - Teste de tolerância à lactose.
    - Avaliação de resistência à insulina.

---

### Chunk 22/30
**Article:** Intolerâncias, Alergias e Hipersensibilidades Alimentares II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.541

.
*   **Medicação:** Em reações agudas e graves (anafilaxia), o uso de anti-histamínicos e corticoides é indispensável.
## ❓ Perguntas
- [Inserir Pergunta/Confusão]
## 📚 Tarefas
- [ ] 1. Ao suspeitar de alergia alimentar, realizar uma anamnese completa, incluindo histórico familiar, tipo de parto, aleitamento e uso de medicamentos.
- [ ] 2. Iniciar a investigação laboratorial com exames básicos como hemograma, dosagem de IgE e marcadores fecais (calprotectina, sangue oculto) se houver sintomas gastrointestinais.
- [ ] 3. Considerar a formação de uma equipe multidisciplinar (nutrólogo, nutricionista, alergista, etc.) para o manejo de casos complexos, especialmente na implementação de dietas de eliminação.
- [ ] 4. Encaminhar o paciente a um alergista para a realização de testes cutâneos (Prick/Patch test) quando indicado para confirmar a sensibilização.
- [ ] 5.

---

### Chunk 23/30
**Article:** TDAH - Parte XXII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.538

ar de distúrbios imunológicos, metabólicos e desequilíbrios nutricionais.
   - Bloquear hipersensibilidade mediada por IgG específica para antígenos alimentares pode ser mecanismo terapêutico; evitar alimentos indutores de alergia é direção futura.
   - Detecção de IgG, IgE, histamina e outros indicadores fornece base objetiva para diagnóstico precoce e avaliação de tratamento.
   - Condições do SNC (Alzheimer, Parkinson, epilepsia, TDAH) associadas ao aumento da permeabilidade da barreira hematoencefálica; hipersensibilidade prolongada mediada por IgE específica a antígenos alimentares pode estar associada à patogênese dessas doenças.

---

### Chunk 24/30
**Article:** Intolerâncias, Alergias e Hipersensibilidades Alimentares II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.537

# Intolerâncias, Alergias e Hipersensibilidades Alimentares II

**Source:** https://web.plaud.ai/share/e7f81765255606372::YXdzOnVzLXdlc3QtMg

---

## Lecture

Data e Hora: 2025-11-20 20:44:42
Local: [Inserir Localização]
Instrutor: Luiz Lox
## 📝 Resumo
Esta aula, ministrada pelo médico gastroenterologista Luiz Lox para a Academia Brasileira de Medicina Funcional Integrativa, oferece uma abordagem abrangente sobre a alergia alimentar, definida como uma reação imunológica adversa a antígenos alimentares. A aula explora a crescente prevalência das alergias, destacando que um pequeno grupo de alimentos (leite, trigo, soja, peixes, ovos, etc.) causa mais de 90% dos casos. São discutidos os mecanismos imunológicos, como a tolerância oral, o papel da microbiota intestinal, a importância da barreira epitelial e os diferentes tipos de reações de hipersensibilidade.

---

### Chunk 25/30
**Article:** Intolerâncias, Alergias e Hipersensibilidades Alimentares II (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.534

imunoglobulinas, fenotipagem linfocitária, testes cutâneos e marcadores fecais), e princípios de manejo como dietas de eliminação, modulação da microbiota, probióticos, nutrientes e compostos fenólicos. Destacou-se a importância da digestibilidade das proteínas, da integridade da barreira intestinal e de equipe multidisciplinar no manejo.
## Conteúdo Não Coberto
1. Testes diagnósticos específicos por tipo de alergia (detalhamento prometido posteriormente)
2. Detalhamento de exames laboratoriais e complementares em protocolos formais
3. Estratégias terapêuticas e modulação intestinal em protocolos práticos padronizados
4. Outros nutrientes além da vitamina A na tolerância oral (serão apresentados futuramente)
5. Discussão aprofundada de hipersensibilidade tipo III e IV aplicadas à alergia alimentar com exemplos
6. Provas dietéticas/terapêuticas com passos práticos e segurança
7.

---

### Chunk 26/30
**Article:** TDAH - Parte XXII (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.533

específica e jejum compõem o conjunto de intervenções.
   - Próxima aula: impacto do exercício físico como regulador essencial, com sustentação para engajamento de pacientes e familiares.
## ❓ Perguntas
- [Insert Question/Confusion]
## 📚 Próximos Arranjos
- [ ] Considerar testes laboratoriais: IgE total e específica, IgG alimentar específica, histamina sérica/urinária, MDA, óxido nítrico, xantinoxidase, vitamina D, ômega 3 (índice ômega-3), zinco e ferritina.
- [ ] Implementar uma dieta de eliminação personalizada, priorizando retirada de potenciais antígenos (ovo, leite, soja, trigo) quando reatividade for sugerida pelos exames.
- [ ] Avaliar suplementação para gestantes e pacientes de risco: ferro, folato, iodo, colina, cobalamina, vitamina D, ômega 3; considerar creatina e sulforafanos conforme evidências e contexto clínico.

---

### Chunk 27/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.532

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

### Chunk 28/30
**Article:** pediatria funcional integrativa - parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.532

oro fisiológico; evitar corticoide e broncodilatador na maioria sem desconforto respiratório significativo.
- APLV (alergia à proteína do leite de vaca) como diferencial em refluxo/cólicas/constipação 0–12 meses; considerar dieta de exclusão antes de medicar.
- Exames sugeridos para avaliação imunológica e nutricional:
  - 25-OH vitamina D, vitamina A.
  - Zinco (idealmente eritrocitário).
  - Perfil de ferro (ferritina, ferro sérico, transferrina/TSAT).
  - Hemograma completo; vitamina B12 opcional.
  - Imunoglobulinas (perfil imunológico) devido a infecções de repetição e múltiplos antibióticos.
  - Prick test para aeroalérgenos (ex.: ácaros).
- Observação clínica em fase aguda (“vir ao consultório quando estiver doente”) para confirmação diagnóstica.

---

### Chunk 29/30
**Article:** Intolerância à Histamina (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.532

# Intolerância à Histamina

**Source:** https://web.plaud.ai/share/08cf1763843274652::YXdzOnVzLXdlc3QtMg

---

## Lecture

Data e Hora: 2025-11-17 17:56:25
Local: [Inserir Local]
Instrutor: Cristiano Ruggi
## 📝 Resumo
Nesta palestra, o Dr. Cristiano Ruggi, médico gastroenterologista, aborda detalhadamente a intolerância à histamina e a síndrome de ativação mastocitária. Ele explica que a histamina é uma molécula neuroimunoendocrinológica com múltiplos receptores e funções, e que a intolerância à histamina resulta de um desequilíbrio entre a histamina acumulada (proveniente da dieta, microbiota e células do corpo) e a capacidade de degradação, principalmente pela enzima diamina oxidase (DAO). A palestra detalha as causas, as diversas manifestações clínicas e a complexidade do diagnóstico, que envolve a exclusão de outras condições, testes sanguíneos para a enzima DAO e análise de polimorfismos genéticos.

---

### Chunk 30/30
**Article:** TDAH - Parte XXII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.529

s.
- Número de pacientes com depressão que apresentaram IgE alta: 66, representando cerca de 36% do grupo depressivo; indica proporção relevante de IgE elevada e evidência de hipersensibilidade tipo 1 mediada por IgE.
- Contagem de indivíduos com IgE alta no grupo controle: 42; sinaliza diferença significativa entre grupos.
- Concentração média de IgE mais alta no grupo com depressão (49) comparada ao grupo não depressivo (31,63), reforçando maior ativação imunológica no grupo depressivo.
**Alimentos cotidianos despontam como possíveis sensibilizadores no grupo analisado, com destaque para ovo e leite.**
- Percentual de reatividade a ovo: 75%; indica alta sensibilização a um alimento amplamente consumido.
- Sensibilização a leite no grupo principal: 47%, versus 10% no grupo comparativo; mostra diferença marcante entre grupos na sensibilização a leite.
- Reatividade a soja: 15%; integra painel de alimentos com sensibilização relevante.

---

