# ScoreItem: Pancreatectomia

**ID:** `019bf31d-2ef0-7522-94c8-c6f9299e4c59`
**FullName:** Pancreatectomia (Histórico de doenças - Cirurgias já realizadas - Cirurgias que interferem diretamente no escore)

**Preparation Metadata:**
- Quality Grade: **GOOD**
- Total Chunks: 30 de 18 artigos
- Avg Similarity: 0.536

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-7522-94c8-c6f9299e4c59`.**

```json
{
  "score_item_id": "019bf31d-2ef0-7522-94c8-c6f9299e4c59",
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

**ScoreItem:** Pancreatectomia (Histórico de doenças - Cirurgias já realizadas - Cirurgias que interferem diretamente no escore)

**30 chunks de 18 artigos (avg similarity: 0.536)**

### Chunk 1/30
**Article:** Modulação Intestinal II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.606

do pâncreas antes de prescrever enzimas digestivas.
- [ ] 4. Monitorar marcadores de inflamação de baixo grau, como resistência à insulina (HOMA-IR), homocisteína e proteína C-reativa.
- [ ] 5. Monitorar os níveis de vitamina B12 ao longo da vida, especialmente em pacientes que usam inibidores de bomba de prótons ou bariátricos.
- [ ] 6. Em pacientes com resistência à insulina, avaliar o TMAO sérico para aferir o risco cardiovascular.
- [ ] 7. Para pacientes que utilizam inibidores da bomba de prótons, planejar um desmame cuidadoso para evitar o efeito rebote de hiperacidez.
- [ ] 8. Aplicar o conhecimento sobre os mecanismos de ação (ex: beta-glucana, butirato) para personalizar as intervenções nutricionais de acordo com as necessidades do paciente (ex: horário de administração para controle de saciedade).

---

### Chunk 2/30
**Article:** Mapping global new-onset, worsening, and resolution of diabetes following partial pancreatectomy: a systematic review and meta-analysis (2023)
**Journal:** International Journal of Surgery
**Section:** abstract | **Similarity:** 0.572

This systematic review examined metabolic outcomes in 13,257 patients across 82 studies who underwent partial pancreatic resection. The overall prevalence of new-onset diabetes after partial pancreatectomy was 17.1%, with significant geographic variation ranging from 7.6% to 38.0%. Patients with chronic pancreatitis demonstrated markedly higher diabetes rates (30.7%) compared to those with benign lesions (16.4%). Distal pancreatectomy carried the highest risk at 23.7%, while central pancreatectomy showed lower incidence at 9.4%. Approximately 41.1% of patients with preoperative diabetes experienced worsening metabolic control, though 25.8% achieved diabetes resolution postoperatively. Among those developing new-onset diabetes, over half (52.9%) ultimately required insulin therapy. Postoperative diabetes represents a substantial clinical concern warranting attention from healthcare providers.

---

### Chunk 3/30
**Article:** Disbiose I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.561

] 4. Solicitar elastase pancreática fecal para investigar insuficiência exócrina pancreática, especialmente em sintomas de má digestão de carboidratos/gorduras.
- [ ] 5. Investigar e tratar SIBO/SIFO/parasitoses (ex.: giardia) em pacientes com intolerâncias a dissacarídeos (lactose) e sintomas de má absorção; restaurar a integridade da mucosa.
- [ ] 6. Revisar a qualidade da dieta do paciente, enfatizando que energia e nutrientes vêm do alimento; alinhar a ingestão para atender cerca de 30 kcal/kg/dia quando apropriado ao estado basal.
- [ ] 7. Educar sobre a importância da saliva e da fase oral da digestão; evitar comer sob ansiedade/pressa, sentar para as refeições e focar no ato de comer.
- [ ] 8. Implementar estratégias para reduzir inflamação crônica de baixo grau, incluindo melhora da microbiota intestinal e redução de “garbage aging” por meio de suporte digestivo e antioxidante.
- [ ] 9.

---

### Chunk 4/30
**Article:** Aula 01 Guilherme Sorrentino - Cirurgia e Medicina Funcional Integrativa (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.557

o intestino é parte da estratégia de cura. O objetivo clínico é abreviar o estado catabólico, fornecendo macro e micronutrientes (e, em casos selecionados, discutindo uso de hormônios anabólicos como testosterona) para proteger massa muscular e acelerar retorno à homeostase.

------------
## Fatores Adicionais de Risco: Coagulação e Hiperglicemia

A coagulação é mapeada com ferramentas como o score de Caprini, ainda que o cenário pós-pandemia tenha aumentado o risco de trombose por disfunção endotelial, exigindo atenção ampliada—incluindo homocisteína como fator trombogênico, com meta abaixo de 10. A hiperglicemia pré-operatória associa-se consistentemente a piores desfechos: além da inflamação vascular, forma produtos finais de glicação (AGEs) que alteram proteínas críticas da cicatrização (fibroblastos, colágeno, células imunológicas), promovendo excesso de oxidação e complicações como vasculopatia, retinopatia, neuropatia e nefropatia.

---

### Chunk 5/30
**Article:** Carboidratos IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.555

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

### Chunk 6/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XVIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.554

r quem decora números, abaixo de 6).
   - Após 75 g de glicose, glicose vai a 216 mg/dL; 2 horas depois glicemia 172 mg/dL; insulina 2 horas 70 µU/mL, permanecendo elevada por longo período, acompanhada de fadiga. Indica hiperglicemia pós-prandial significativa e hiperinsulinemia tardia, com risco de avaliação inadequada se apenas jejum for considerado.
* Implicações perioperatórias
   - Estresse cirúrgico eleva cortisol; uso de corticoide pode somar; oferta de lanches hospitalares típicos aumenta estímulo insulinêmico, levando a pior imunidade, cicatrização, inflamação, oxidação e glicação; maior risco de deiscência, necrose, infecção e feridas.
   - Necessidade de triagem e intervenções para reduzir complicações, não apenas negar cirurgia.
### 3. Estado nutricional e cicatrização
* Prevalência e risco
   - Até 25% dos pacientes ambulatoriais de cirurgia plástica estão em risco de desnutrição e não são avaliados; aproximadamente 1 em cada 4.

---

### Chunk 7/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XIX (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.552

ós-Cirurgia Bariátrica**
    - Pacientes pós-bariátricos frequentemente apresentam dificuldades na absorção de nutrientes, sendo as deficiências de ferro e complexo B (especialmente B12) as mais comuns.
    - A deficiência de B12 ocorre devido à separação anatômica do estômago, que impede a união do fator intrínseco (produzido no fundo gástrico) com o fator extrínseco (a vitamina B12), e também pelo desvio de um segmento do intestino delgado crucial para sua absorção.
    - É crucial uma avaliação nutricional completa e a correção de déficits pré-existentes antes da cirurgia, algo que, segundo o instrutor, raramente é feito.
    - O acompanhamento nutricional contínuo ao longo da vida, com suplementação personalizada de polivitamínicos e minerais, é essencial para prevenir complicações.
*   **Estratégias para Preservar Massa Magra no Emagrecimento**
    - A prática de exercícios resistidos é fundamental para a manutenção da massa magra.

---

### Chunk 8/30
**Article:** Aula 01 Guilherme Sorrentino - Cirurgia e Medicina Funcional Integrativa (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.547

sta, Anestesia)
  - [ ] Nutricionista: planejar preabilitação dietética (p.ex., jejum intermitente, dieta cetogênica quando indicada, dieta anti-inflamatória/antifermentativa) – Início imediato e ao longo do pré-operatório
  - [ ] Nutricionista: realizar bioimpedância para mapear composição corporal, detectar “falso magro” e gordura visceral – Pré-operatório
  - [ ] Gastroenterologista: avaliar função digestiva (suco gástrico, enzimas pancreáticas), intolerâncias (laticínios, glúten) e sensibilidade a FODMAPs – Pré-operatório
  - [ ] Gastroenterologista: considerar exames avançados (GI-MAP, gut check, nutrigenética) em cirurgias maiores ou casos complexos – Antes da definição do plano cirúrgico
  - [ ] Anestesia: planejar manejo do estresse cirúrgico, manter normotermia e ajustar hemodinâmica visando evitar hipoperfusão e excesso de fluidos – Intraoperatório

- Paciente
  - [ ] Participar da decisão compartilhada sobre timing cirúrgico e a

---

### Chunk 9/30
**Article:** Trato Gastrointestinal V – Intestino Delgado I (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.545

as e risco biológico: atualização dietética e preservação telomérica são determinantes de desfechos cardiovasculares e infecciosos.**
- A dieta tradicional de diabetes com 60% de carboidratos integrais é criticada como obsoleta, motivando revisões para melhor controle metabólico.
- Telômeros curtos associam-se a aumento de 300% no risco de morte cardíaca e 800% em doenças infecciosas, ressaltando a importância de estratégias protetoras.
**Achados-Chave Adicionais**
- Estudo pediátrico (2016): 174 crianças de 1–4 anos, 12 semanas, randomizado duplo-cego e placebo-controlado com beta-glucana, observando redução de episódios de doenças comuns.
- Idade do primeiro câncer de mama familiar: 35 anos na irmã gêmea da paciente, ilustrando risco familiar e impacto psicológico em decisões de prevenção/terapias.
- Espera inicial de dois meses antes de análogos de GLP-1 serve como janela de avaliação da eficácia de intervenções não farmacológicas.

---

### Chunk 10/30
**Article:** MFI - SÍNDROME PÓS COVID - AULA 02 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.545

, antioxidação e regeneração pancreática.
*   [ ] Tratar a hiperglicemia com uma abordagem combinada: metformina, berberina, cromo, zinco, magnésio, ômega-3, curcuminoides e, preferencialmente, FMD.
*   [ ] Elaborar planos alimentares ricos em nutrientes para combater a resistência à insulina.
*   [ ] Educar os pacientes sobre a importância da carga glicêmica para promover autonomia e melhores escolhas.
*   [ ] Realizar uma curva insulinêmica glicêmica para personalizar o tratamento da resistência à insulina.
*   [ ] Focar em nutrir a tireoide e tratar a inflamação como primeira linha de abordagem para disfunções tireoidianas, antes de considerar a prescrição hormonal.
*   [ ] Excluir alimentos ricos em histamina por um período para avaliar a resposta clínica em pacientes com suspeita de intolerância.

---

### Chunk 11/30
**Article:** Preparing Patients for Cosmetic Surgery and Aesthetic Procedures: Ensuring an Optimal Nutritional Status for Successful Results (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.544

overalleffectsimproveinﬂammatorybiomarkersandthelevelofoxidativestress[63].EarlyidentiﬁcationinandtreatmentofpatientswhomaybeatriskofvitaminDdeﬁciencyiscritical,especiallyinpatientswhohaveundergonebariatricsurgeryandwhoarereferredforplasticprocedures.VitaminC:VitaminCisanessentialcofactorforvariousenzymaticreactionsandhasstrongantioxidantproperties.Duringthehydroxylationofprolineandlysine,vitaminCisimportantforcollagenformation[60].Italsoaccelerateswoundhealingandcontributestobedsorehealing.Thecombinationofsurgeryprocedureswithpre-existinginsufﬁcientvitaminCstatusmayleadtosigniﬁcantalterationsinwoundhealing.PreclinicalstudieshaveshownthatvitaminCsupplementationresultsinhigherexpressionofwoundrepairmediatorsandreducedexpressionofpro-inﬂammatorymediatorsfortheearlyresolutionoftissueremodelingandinﬂammation[64,65].VitaminCdeﬁcitcanleadtocapillaryfragility,disturbancesintheproductionofcollagen,slowerwoundhealing,andreducedresistancetoinfection,aswellasscurvy[66].VitaminE:Du

---

### Chunk 12/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XVIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.541

mplementar avaliação pré-operatória de resistência insulínica: solicitar glicemia e insulina de jejum para índice HOMA e, quando possível, realizar curva insulinêmica pós-carga de 75 g de glicose.
- [ ] 2. Padronizar triagem nutricional em pacientes de cirurgia plástica: avaliar risco de desnutrição, ingestão proteica, vitamina C, selênio, zinco, cobre e aminoácidos específicos; encaminhar para nutricionista quando necessário.
- [ ] 3. Incluir dosagem de homocisteína na avaliação de risco trombótico, especialmente em cirurgias longas, pacientes de idade avançada, usuárias de anticoncepcionais ou gestantes.
- [ ] 4. Revisar protocolos de alimentação hospitalar pós-operatória para reduzir picos glicêmicos e estímulos insulinêmicos; considerar opções de lanches com baixo índice glicêmico e maior densidade proteica.
- [ ] 5.

---

### Chunk 13/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XVIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.538

eito de risco calculado. A organização foi lógica, começando com o contexto geral antes de mergulhar nos detalhes técnicos.
### 2. Resistência Insulínica e Desfechos Cirúrgicos
- A resistência insulínica é um mecanismo chave para complicações cirúrgicas.
- Uma queda de 50% na sensibilidade à insulina pós-cirurgia pode aumentar o risco de complicações graves em 5–6 vezes e de infecção grave em mais de 10 vezes.
- A avaliação padrão (glicemia, hemoglobina glicada) é insuficiente; é crucial medir a insulina para calcular o índice HOMA e, idealmente, solicitar uma curva insulinêmica.
- Caso clínico: paciente com insulina basal de 3 e glicemia de 101 (HOMA normal), mas com picos de glicose (216) e insulina (70) após estímulo, revelando estado pré-diabético oculto.
- Operar um paciente nessa condição aumenta risco de inflamação, piora na cicatrização e baixa imunidade, especialmente com o estresse metabólico da cirurgia e a dieta hospitalar.

---

### Chunk 14/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XVIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.537

to.
- Operar um paciente nessa condição aumenta risco de inflamação, piora na cicatrização e baixa imunidade, especialmente com o estresse metabólico da cirurgia e a dieta hospitalar.
> **Sugestões da IA**
> O uso de um caso clínico real para demonstrar a insuficiência do índice HOMA isolado foi excelente e muito didático. Você explicou de forma clara como os exames convencionais podem ser enganosos. Para reforçar ainda mais, ao mostrar os números do exame (glicose 216, insulina 70), você poderia usar um slide simples com um gráfico da curva glicêmica e insulinêmica do paciente. Isso ajudaria a visualizar o "pico" e a "queda lenta" que você descreveu, tornando o conceito ainda mais impactante e fácil de memorizar.
### 3. Desnutrição e Cicatrização
- Até 25% dos pacientes de cirurgia plástica ambulatorial estão em risco de desnutrição.
- O déficit nutricional diminui a imunidade e a capacidade de superar estressores como infecções.

---

### Chunk 15/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XVIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.535

ng Note

Data e Hora: 2025-11-17 16:33:35
Local: [Inserir Local]: [Inserir Local]
Aula: [Inserir Nome da Aula]: [Inserir Nome da Aula]
## Visão Geral
A aula abordou a aplicação da visão funcional e integrativa na cirurgia plástica para reduzir riscos e melhorar os resultados. Foram discutidos os impactos da resistência insulínica, desnutrição e hiper-homocisteinemia nos desfechos cirúrgicos. Também foi explorado o uso da oxandrolona no tratamento de queimaduras e seu potencial em cirurgias plásticas, além de uma análise crítica sobre os riscos e a subnotificação de complicações em procedimentos estéticos.
## Conteúdo Remanescente
1. Análise aprofundada da cirurgia bariátrica e perda de peso.
2. Estratégias para otimizar respostas metabólicas e desfechos pós-cirurgia bariátrica.
3. Manutenção da composição corporal em processos de emagrecimento (aplicável a ganho de massa magra, envelhecimento saudável e recuperação de cirurgia plástica).

---

### Chunk 16/30
**Article:** Emagrecimento XIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.533

C10) + goma acácia + proteína a partir de 16h30, consumida gradualmente.
- [ ] 9. Avaliar risco de refluxo e hipocloridria antes/depois de iniciar agonistas de GLP‑1; modular secreção gástrica e esvaziamento conforme necessário.
- [ ] 10. Mapear perfil hedônico (estresse/humor/depressão) dos pacientes antes de prescrever agonistas de GLP‑1 e prover suporte dopaminérgico/psicossocial.
- [ ] 11. Revisar módulo de sistema gastrointestinal (EDA, DEP, hormônios GI) para consolidar base fisiológica.
- [ ] 12. Monitorar marcadores metabólicos em DM2 ao testar intervenções com gorduras; priorizar monoinsaturadas.
- [ ] 13. Preparar materiais práticos (lista de compras, receitas rápidas, guia de uso de shakeira) para facilitar adesão.
- [ ] 14. Aguardar e participar da próxima aula sobre incretinas e medicações, comparando eficácia e custo‑benefício para protocolos de emagrecimento.

---

### Chunk 17/30
**Article:** Fisiologia e Bioquímica do Sistema Imune II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.530

olicitar e registrar parâmetros essenciais (PCR-us, VHS); calcular índices estimáveis (HGI, TAIG); considerar complementares (TNF-α, IL-6, GPx, MDA, antioxidantes totais).
- [ ] 11. Monitorar trimestralmente HbA1c, glicemia de jejum e HGI; estabelecer metas individuais (ex.: reduzir de 6,1 para ~5,5; longo prazo ~5,3).
- [ ] 12. Calcular TG/HDL e integrar com lipidograma/SREBP1c/2; evitar combinação de gordura saturada com açúcar e excesso de saturadas em múltiplas refeições.
- [ ] 13. Avaliar ferro, ferritina, transferrina e saturação (20–50%; evitar <20%); interpretar com hepcidina/SREBP1c e quadro inflamatório.
- [ ] 14. Medir TNF-α (<8,1) e IL-6 (<3,4) para acompanhar atividade inflamatória; relacionar com obesidade inflamada.
- [ ] 15. Calcular HOMA-β (167–175) e HOMA-IR (<2,15); buscar glicemia 60–90 e insulina ~6–7.
- [ ] 16. Monitorar homocisteína (<7,9) e PCR; usar PCR-us; documentar crises (PCR >1.000) e conduzir manejo apropriado.
- [ ] 17.

---

### Chunk 18/30
**Article:** Preparing Patients for Cosmetic Surgery and Aesthetic Procedures: Ensuring an Optimal Nutritional Status for Successful Results (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.523

immunityandinabilitytotolerateandovercomestressors,suchasinfections[10,11].Nutritionalstatuscanmarkedlyaffectwoundhealingandtissuerepairfollowingsurgicalinterventions[12].Individualswhoundergocosmeticplasticsurgeryandaestheticmedicinespanthenutritionalspectrumfrombeingpatientswhoaregenerallyhealthyandnutritionallyadequatetopatientswhoareinherentlycatabolic,withchronicwounds,andnutritionallydeﬁcient[6].Ithasbeenshownthatupto25%ofplasticsurgeryoutpatientsareatriskofmalnutrition[13].Professionalsworkinginthisindustry,therefore,needtoconsiderthenutritionalaspectsoftheirpatientsinordertoobtainthebestsurgicalresults.Theaimofthisnarrativereviewistoidentifythenutritionaldeﬁcitsorexcessesassociatedwiththemajorcomplicationsofreconstructivesurgery,aestheticsurgery,andmini-invasiveaestheticmedicalprocedures.Anotheraimistoprovideachecklistorbundleofactionsforprofessionalsworkinginthisindustry,sothattheycanreducetherisksassociatedwithaestheticprocedures.2.MajorComplicationsofReconstructiveSurger

---

### Chunk 19/30
**Article:** Modulação Intestinal II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.521

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

### Chunk 20/30
**Article:** Hipertensão Arterial Sistêmica II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.521

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

### Chunk 21/30
**Article:** Aula 02 Guilherme Sorrentino - Suplementação em Cirurgia (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.521

ndo que qualquer discussão sobre suplementos é necessariamente parcial, dado que o corpo requer um espectro completo de nutrientes. Embora a gama de opções seja ampla, sustenta que, com um conjunto “básico” de intervenções, já é possível oferecer ganhos clínicos significativos. Define objetivos operacionais claros: acelerar a cicatrização, reduzir risco de infecção e dar suporte ao metabolismo e à função mitocondrial, inclusive auxiliando o fígado em processos de detoxificação. Defende uma estratégia personalizada, orientada por avaliação das individualidades bioquímicas (ex.: o que é indicado para um paciente pode não ser para outro), pois a demanda metabólica imposta pelo ato cirúrgico supera a capacidade da dieta habitual em suprir necessidades “ótimas”.

---

### Chunk 22/30
**Article:** Emagrecimento XIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.520

UA: responsabilidade recai na escolha do paciente; bons resultados validam práticas interdisciplinares.
- Postura recomendada: fazer bem feito, construir comunidade do bem e evitar interferências desnecessárias.
### 4. Regulação neuroendócrina gastrointestinal de curto prazo
- Incretinas (GLP‑1, GLP‑2, GIP) reduzem fome, modulam prazer e homeostase; receptores no hipotálamo e sistema límbico com componente dopaminérgico.
- Efeitos: esvaziamento gástrico mais lento, possível hipocloridria e refluxo; necessidade de modulação digestiva integrada.
- DPP‑4 inativa incretinas rapidamente; medicações devem ser ponderadas e acompanhadas de correções sistêmicas.
### 5. Agonistas de GLP‑1: histórico e cautelas
- Exenatida (BYETTA) como primeiro análogo, resistente ao DPP‑4; desenvolvidos para DM2 com efeito adicional em peso.
- Benefícios: insulina dependente da glicose, redução de glucagon, suporte às células beta.

---

### Chunk 23/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XVIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.519

s com baixo índice glicêmico e maior densidade proteica.
- [ ] 5. Estudar viabilidade e protocolos de uso de oxandrolona em contextos selecionados (grandes queimados; suporte pós-operatório de cirurgias de contorno corporal) com doses baixas, monitoramento de perfil hepático e lipídico, e consentimento informado.
- [ ] 6. Desenvolver materiais educativos para pacientes sobre riscos e desfechos de lipoaspiração, ressaltando a necessidade de mudanças de hábitos para evitar aumento compensatório de gordura visceral.
- [ ] 7. Mapear e atualizar diretrizes internas com evidências recentes (New England Journal of Medicine sobre homocisteína e TVP; meta-análise de oxandrolona em queimados) para embasar solicitações de exames e terapias adjuvantes.
- [ ] 8. Preparar conteúdo para a próxima aula: integração de estratégias nutricionais e metabólicas na bariátrica visando manutenção/ganho de massa magra e prevenção de regressão do peso.

---

### Chunk 24/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XIV (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.519

- [ ] 2. Para pacientes com síndrome metabólica, considerar estratégia Mediterrânea organizada; preferir versão com restrição calórica combinada com atividade física, monitorando resultados aos 6 e 12 meses.
- [ ] 3. Solicitar perfil lipídico ampliado em casos selecionados: HDL, triglicerídeos, insulina, PCR, LDL oxidado e, quando indicado, subfracionamento de LDL; evitar decisões baseadas apenas em LDL total.
- [ ] 4. Planejar exames de risco cardiovascular conforme necessidade: score de cálcio coronariano e angiotomografia de coronárias (incluindo avaliação de placas moles) quando o contexto clínico justificar.
- [ ] 5. Revisar e atualizar protocolos internos sobre álcool: remover recomendações de consumo “cardioprotetor”; educar pacientes sobre riscos de câncer e piora do sono; avaliar predisposição/sensibilidade individual.
- [ ] 6.

---

### Chunk 25/30
**Article:** Preparing Patients for Cosmetic Surgery and Aesthetic Procedures: Ensuring an Optimal Nutritional Status for Successful Results (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.518

itionaldeﬁcitsorexcessesassociatedwiththemajorcomplicationsofreconstructivesurgery,aestheticsurgery,andmini-invasiveaestheticprocedures.Anadditionalgoalistoprovideabundleofactionsforprofessionalsworkingintheindustryinordertoreducetherisksofaestheticproceduresandimprovetheclinicaloutcomes.Granulomas,hypertrophicscarsandkeloids,seromas,infectionsandxerosis,hyperpigmentation,petechiae,livedoreticularis,slowerwoundhealing,andotherpooroutcomesarefrequentlyassociatedwithnutritionaldeﬁciencies.Nutritionalstatuscanmarkedlyaffectwoundhealingandtissuerepairfollowingsurgicalinterventions,aswellastheoutcomesofaestheticandcosmeticmedicalpractices.Professionalsworkinginthisindustry,therefore,needtoconsiderthenutritionalaspectsoftheirpatientstoobtainthebestresults.Keywords:nutritionalstate;nutritionaldeﬁcit;woundhealing;cosmeticsurgery;aestheticprocedures
1.IntroductionAestheticmedicalpracticesareboomingglobally.WithincreasedprevalenceandtheaccessibilityandevolutionofcosmeticsurgeryinWesternsoc

---

### Chunk 26/30
**Article:** Emagrecimento - Parte I (2024)
**Journal:** Pos Graduacao MFI
**Section:** methods | **Similarity:** 0.518

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

### Chunk 27/30
**Article:** Dieta Cetogênica - Parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.515

um intermitente zeraram o score de severidade dos sintomas.
### Achados Adicionais Chave
- Uma meta-análise de 14 ensaios clínicos incluiu 734 participantes com sobrepeso e obesidade, dos quais 444 eram diabéticos e 290 não diabéticos.
- Outra revisão sistemática e meta-análise sobre diabetes tipo 2 avaliou 13 estudos com um total de 567 participantes.
- O efeito térmico da proteína, que converte 25% de suas calorias em calor, é um mecanismo que contribui para a perda de peso em dietas com maior teor proteico.

---

## SOAP

Data e Hora: 2025-11-18 17:49:32
Paciente: 
Diagnóstico:
## Histórico de Diagnóstico:
1. Histórico Médico: Apresentação médica sobre benefícios da dieta cetogênica e low-carb para diabetes tipo 2, obesidade, psoríase e esclerose múltipla; não há informações de um paciente específico.
2.

---

### Chunk 28/30
**Article:** Preparing Patients for Cosmetic Surgery and Aesthetic Procedures: Ensuring an Optimal Nutritional Status for Successful Results (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.515

ithEatingDisorders.Am.J.Clin.Dermatol.2005,6,165–173.[CrossRef]40.Englesbe,M.J.;Lee,J.S.;He,K.;Fan,L.;Schaubel,D.E.;Sheetz,K.H.;Harbaugh,C.M.;Holcombe,S.A.;Campbell,D.A.J.;Sonnenday,C.J.;etal.AnalyticMorphomics,CoreMuscleSize,andSurgicalOutcomes.Ann.Surg.2012,256,255–261.[CrossRef]41.Lee,J.S.;Terjimanian,M.N.;Tishberg,L.M.;Alawieh,A.Z.;Harbaugh,C.M.;Sheetz,K.H.;Holcombe,S.A.;Wang,S.C.;Sonnenday,C.J.;Englesbe,M.J.SurgicalSiteInfectionandAnalyticMorphometricAssessmentofBodyCompositioninPatientsUndergoingMidlineLaparotomy.J.Am.Coll.Surg.2011,213,236–244.[CrossRef]42.Tan,B.H.L.;Birdsell,L.A.;Martin,L.;Baracos,V.E.;Fearon,K.C.H.SarcopeniainanOverweightorObesePatientisanAdversePrognosticFactorinPancreaticCancer.Clin.CancerRes.2009,15,6973–6979.[CrossRef]43.Gillis,C.;Hasil,L.;Kasvis,P.;Bibby,N.;Davies,S.J.;Prado,C.M.;West,M.A.;Shaw,C.NutritionCareProcessModelApproachtoSurgicalPrehabilitationinOncology.Front.Nutr.2021,8,644706.[CrossRef]44.Mazza,E.;Ferro,Y.;Pujia,R.;Mare,R.;Maurotti,S.

---

### Chunk 29/30
**Article:** Disbiose II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.515

oridria.
### 15. Manifestações clínicas da disbiose
- Digestivas: distensão abdominal, meteorismo, DRGE, DII, SII, alterações do hábito intestinal.
- Extraintestinais: alergias, autoimunidade, câncer, saúde mental e hormonal.
### 16. Estratégias terapêuticas: otimizar digestão e intervenções
- Priorizar otimização digestiva (enzimas) antes de probióticos.
- Alimentos ricos em enzimas: kiwi, mamão, limão, abacaxi antes das refeições.
- Mindful eating, mastigação adequada, fracionamento de volumes conforme tolerância.
- Cautela com janelas alimentares curtas/jejum intermitente em certos pacientes.
- Pancreatina (Creon): origem porcina; dose adulta ≥20.000 UI; opções 10.000/25.000 UI; preferir cápsulas gastro-resistentes (ação em pH duodenal básico).
- Não associar pancreatina com betaína HCl na mesma cápsula; timing: betaína HCl durante a refeição (liberação gástrica), pancreatina antes (T–15 min).

---

### Chunk 30/30
**Article:** Aula 01 Guilherme Sorrentino - Cirurgia e Medicina Funcional Integrativa (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.514

, manter normotermia e ajustar hemodinâmica visando evitar hipoperfusão e excesso de fluidos – Intraoperatório

- Paciente
  - [ ] Participar da decisão compartilhada sobre timing cirúrgico e aderir ao plano de preabilitação (nutrição, suplementação, manejo do estresse) – Antes do agendamento final
  - [ ] Seguir orientações para otimização metabólica (adesão às estratégias dietéticas e suplementação prescritas) – Pré-operatório contínuo

---

