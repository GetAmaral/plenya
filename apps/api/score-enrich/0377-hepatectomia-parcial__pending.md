# ScoreItem: Hepatectomia parcial

**ID:** `019bf31d-2ef0-749e-9e66-045319ceaaa3`
**FullName:** Hepatectomia parcial (Histórico de doenças - Cirurgias já realizadas - Cirurgias que interferem diretamente no escore)

**Preparation Metadata:**
- Quality Grade: **GOOD**
- Total Chunks: 30 de 15 artigos
- Avg Similarity: 0.568

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-749e-9e66-045319ceaaa3`.**

```json
{
  "score_item_id": "019bf31d-2ef0-749e-9e66-045319ceaaa3",
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

**ScoreItem:** Hepatectomia parcial (Histórico de doenças - Cirurgias já realizadas - Cirurgias que interferem diretamente no escore)

**30 chunks de 15 artigos (avg similarity: 0.568)**

### Chunk 1/30
**Article:** Metabolic changes after hepatectomy: Implications for perioperative management and long-term outcomes (2025)
**Journal:** World Journal of Gastrointestinal Pathophysiology
**Section:** abstract | **Similarity:** 0.691

Hepatectomy initiates intricate metabolic adaptations as the liver adjusts to functional mass loss and begins regeneration. The procedure triggers alterations across carbohydrate, lipid, and protein metabolism, alongside changes in oxygen and energy utilization. Hepatectomy produces a metabolic shift toward gluconeogenesis with transient insulin resistance. Hyperglycemia persisting up to 16 hours postoperatively serves as an indicator of glycogenolysis and gluconeogenesis. Lipid metabolism shifts dramatically, with adipose tissue mobilizing free fatty acids transported to the liver. Protein metabolism becomes imbalanced following hepatectomy, potentially causing hyperammonemia and hepatic encephalopathy.

---

### Chunk 2/30
**Article:** Liver Regeneration after Hepatectomy and Partial Liver Transplantation (2020)
**Journal:** International Journal of Molecular Sciences
**Section:** abstract | **Similarity:** 0.668

Liver regeneration following partial hepatectomy is a unique physiological response that restores hepatic mass and function through tightly orchestrated cellular and molecular events. The process involves both hepatocyte hyperplasia and hypertrophy, triggered primarily by hemodynamic alterations such as increased portal pressure and shear stress. Regeneration capacity is remarkable with healthy livers able to lose up to two-thirds of volume and still recover. The critical timeframe for regaining hepatic function after partial hepatectomy is 5-7 days. Complete restoration of residual liver size occurs within 3-6 months. Factors affecting regeneration include degree of cirrhosis, residual liver volume, and presence of postoperative complications.

---

### Chunk 3/30
**Article:** Nutritional management after hepatopancreatobiliary surgery (2021)
**Journal:** Hepatobiliary Surgery and Nutrition
**Section:** abstract | **Similarity:** 0.662

Nutrition plays a critical role in postoperative recovery after hepatobiliary surgery. Malnutrition is common in 30-50% of hospitalized patients. The liver possesses remarkable regenerative capacity, yet major hepatectomy causes significant biochemical and metabolic changes. Without adequate nutrient supplementation, serious consequences including jaundice, hepatic ascites, and hepatic failure may develop. Hepatectomy patients frequently experience hypoglycemia requiring glucose monitoring. Albumin supplementation requires approximately one week, with levels normalizing by postoperative week three. Liver resection patients benefit from early postoperative diets rich in branched-chain amino acids, plus high-fat diets based on medium-chain triglycerides.

---

### Chunk 4/30
**Article:** Aula 01 Guilherme Sorrentino - Cirurgia e Medicina Funcional Integrativa (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.605

infecções ou inflamação crônica—daí a necessidade de aproximar o fígado da homeostase antes de operar. No pós-operatório, o paciente tipicamente apresenta hiperglicemia, acidose láctica, edema/retenção líquida, resistência periférica à insulina e replanejamento da oferta de glicose a órgãos prioritários (cérebro, coração, rins). A inflamação ultrapassa a barreira cutânea, alcançando intestino e barreira hematoencefálica; vias como NF-κB e o inflamassoma NLRP3 são ativadas, com liberação de citocinas (IL-1, IL-6, TNF-α). A musculatura assume papel imunológico por ser o reservatório de aminoácidos para síntese de células de defesa e componentes estruturais da resposta inflamatória, justificando o foco em composição corporal adequada (bioimpedância, identificação de “falso magro” e gordura visceral, que é mais inflamatória).

---

### Chunk 5/30
**Article:** Aula 01 Guilherme Sorrentino - Cirurgia e Medicina Funcional Integrativa (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.581

s (1–2 semanas), cabe ao cirurgião considerar o momento ideal, em decisão compartilhada com paciente e família, ponderando aptidão física, estado nutricional, bem-estar psicológico e indicação técnica da cirurgia, com eventual adiamento quando o risco é modificável.

------------
## Pilares da Avaliação Pré-Operatória Integrativa

O método do Dr. Sorrentino organiza-se em sete pilares interdependentes: porte cirúrgico, risco cardíaco, função renal, função hepática, estado nutricional, coagulação e perfil inflamatório. A avaliação expande-se além dos exames tradicionais (hemograma, ureia, creatinina, glicemia de jejum, TP, KTTP), incorporando marcadores que refinam o entendimento bioquímico e o risco individual: insulina em jejum, dímero-D, proteína C reativa ultrassensível, homocisteína, e, conforme necessidade, TNF-alfa, CPK e testes relacionados à acidez gástrica e ao metabolismo intestinal.

---

### Chunk 6/30
**Article:** Aula 01 Guilherme Sorrentino - Cirurgia e Medicina Funcional Integrativa (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.579

to: usar frequência cardíaca como guia; intervir se >120 e progressiva apesar de reposição – Intraoperatório contínuo
  - [ ] Evitar exceder 6 horas de tempo cirúrgico e evitar excesso de fluidos – Planejamento e intraoperatório
  - [ ] Assegurar termorregulação do paciente (mantas térmicas, soluções aquecidas quando possível) – Intraoperatório
  - [ ] Aplicar controle de danos e pausar etapas quando sangramento for acima do esperado (p.ex., interromper mastopexia após lipoaspiração com sangramento elevado) – Decisão intraoperatória
  - [ ] Reduzir duração de antibióticos/anti-inflamatórios quando clinicamente seguro; priorizar suplementação e medidas naturais adjuntas – Pós-operatório inicial

- Equipe Multidisciplinar (Nutricionista, Gastroenterologista, Anestesia)
  - [ ] Nutricionista: planejar preabilitação dietética (p.ex., jejum intermitente, dieta cetogênica quando indicada, dieta anti-inflamatória/antifermentativa) – Início imediat

---

### Chunk 7/30
**Article:** Functional assessment of liver regeneration after major hepatectomy (2022)
**Journal:** Hepatobiliary Surgery and Nutrition
**Section:** abstract | **Similarity:** 0.578

This prospective study evaluated early postoperative changes in remnant liver function, volume, and stiffness after major liver resection (≥3 segments). Function and volume of the remnant liver had increased by the 5th postoperative day significantly, with function rising from 6.9 to 9.6 %/min/m² (P=0.004). Functional regeneration plateaued between day 5 and 4-6 weeks, while volumetric increase continued. Patients who had severe postoperative complications did not show significant increase in liver function on the 5th postoperative day despite increase of volume. Functional liver regeneration occurs predominantly within five days post-resection.

---

### Chunk 8/30
**Article:** Aula 02 Guilherme Sorrentino - Suplementação em Cirurgia (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.571

ndo que qualquer discussão sobre suplementos é necessariamente parcial, dado que o corpo requer um espectro completo de nutrientes. Embora a gama de opções seja ampla, sustenta que, com um conjunto “básico” de intervenções, já é possível oferecer ganhos clínicos significativos. Define objetivos operacionais claros: acelerar a cicatrização, reduzir risco de infecção e dar suporte ao metabolismo e à função mitocondrial, inclusive auxiliando o fígado em processos de detoxificação. Defende uma estratégia personalizada, orientada por avaliação das individualidades bioquímicas (ex.: o que é indicado para um paciente pode não ser para outro), pois a demanda metabólica imposta pelo ato cirúrgico supera a capacidade da dieta habitual em suprir necessidades “ótimas”.

---

### Chunk 9/30
**Article:** Aula 01 Guilherme Sorrentino - Cirurgia e Medicina Funcional Integrativa (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.571

o intestino é parte da estratégia de cura. O objetivo clínico é abreviar o estado catabólico, fornecendo macro e micronutrientes (e, em casos selecionados, discutindo uso de hormônios anabólicos como testosterona) para proteger massa muscular e acelerar retorno à homeostase.

------------
## Fatores Adicionais de Risco: Coagulação e Hiperglicemia

A coagulação é mapeada com ferramentas como o score de Caprini, ainda que o cenário pós-pandemia tenha aumentado o risco de trombose por disfunção endotelial, exigindo atenção ampliada—incluindo homocisteína como fator trombogênico, com meta abaixo de 10. A hiperglicemia pré-operatória associa-se consistentemente a piores desfechos: além da inflamação vascular, forma produtos finais de glicação (AGEs) que alteram proteínas críticas da cicatrização (fibroblastos, colágeno, células imunológicas), promovendo excesso de oxidação e complicações como vasculopatia, retinopatia, neuropatia e nefropatia.

---

### Chunk 10/30
**Article:** Aula 01 Guilherme Sorrentino - Cirurgia e Medicina Funcional Integrativa (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.570

tineira: quedas em B6, B9, B12 e betaína prejudicam metilação, elevando homocisteína (objetivo: valores abaixo de 10).

------------
## Avaliação da Função Orgânica e do Perfil Inflamatório Sistêmico

A inflamação sistêmica do contexto cirúrgico impacta diversos sistemas. Renalmente, há maior demanda funcional, redução de eritropoetina e alterações que, junto ao aumento de hepsidina hepática, prejudicam absorção e uso do ferro, promovendo retenção em macrófagos e ferritina. O fígado é descrito como maestro metabólico: conduz gliconeogênese, produz proteínas de fase aguda, sustenta detoxificação e gestão energética. Observa-se, na prática atual, TGO/TGP frequentemente entre 35, 40, 45, 60, indicativos de sobrecarga hepática em muitos pacientes por dieta, infecções ou inflamação crônica—daí a necessidade de aproximar o fígado da homeostase antes de operar.

---

### Chunk 11/30
**Article:** Aula 01 Guilherme Sorrentino - Cirurgia e Medicina Funcional Integrativa (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.569

sta, Anestesia)
  - [ ] Nutricionista: planejar preabilitação dietética (p.ex., jejum intermitente, dieta cetogênica quando indicada, dieta anti-inflamatória/antifermentativa) – Início imediato e ao longo do pré-operatório
  - [ ] Nutricionista: realizar bioimpedância para mapear composição corporal, detectar “falso magro” e gordura visceral – Pré-operatório
  - [ ] Gastroenterologista: avaliar função digestiva (suco gástrico, enzimas pancreáticas), intolerâncias (laticínios, glúten) e sensibilidade a FODMAPs – Pré-operatório
  - [ ] Gastroenterologista: considerar exames avançados (GI-MAP, gut check, nutrigenética) em cirurgias maiores ou casos complexos – Antes da definição do plano cirúrgico
  - [ ] Anestesia: planejar manejo do estresse cirúrgico, manter normotermia e ajustar hemodinâmica visando evitar hipoperfusão e excesso de fluidos – Intraoperatório

- Paciente
  - [ ] Participar da decisão compartilhada sobre timing cirúrgico e a

---

### Chunk 12/30
**Article:** Aula 01 Guilherme Sorrentino - Cirurgia e Medicina Funcional Integrativa (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.565

, manter normotermia e ajustar hemodinâmica visando evitar hipoperfusão e excesso de fluidos – Intraoperatório

- Paciente
  - [ ] Participar da decisão compartilhada sobre timing cirúrgico e aderir ao plano de preabilitação (nutrição, suplementação, manejo do estresse) – Antes do agendamento final
  - [ ] Seguir orientações para otimização metabólica (adesão às estratégias dietéticas e suplementação prescritas) – Pré-operatório contínuo

---

### Chunk 13/30
**Article:** Fisiologia e Bioquímica do Sistema Imune II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.564

60 e 90
- [ ] Manter a insulina, o mais baixo possível, 6, 7, estourando 8
- [ ] Avaliar a homocisteína, pois é um marcador inflamatório importante
- [ ] Usar a proteína C-reativa, associado com os níveis de homocisteína
- [ ] Verificar os parâmetros essenciais na avaliação inflamatória
- [ ] Estimar o índice de glicação e o índice TAIG, baseado nos resultados essenciais
- [ ] Complementar a avaliação com TNF-alfa, IL-6, glutationa e malon de aldeído
### Tarefas para @
- [ ] Usar um concentrado de C8 ou um mix de C8 e C10, para estimular mais ainda o CP3 e as UCPs (proteínas desacopladoras), diminuir a produção de espécie reativa de oxigênio e aumentar a oxidação de gordura @
- [ ] Incluir mioinositol, trans-resveratrol e epigalocatequina galato na formulação, para diminuir os compostos de glicação avançada e a hemoglobina glicada @
- [ ] Fazer uma boa distribuição de gordura e trabalhar os ácidos graxos de cadeia curta, para obter o melhor benefício p

---

### Chunk 14/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XVIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.563

s com baixo índice glicêmico e maior densidade proteica.
- [ ] 5. Estudar viabilidade e protocolos de uso de oxandrolona em contextos selecionados (grandes queimados; suporte pós-operatório de cirurgias de contorno corporal) com doses baixas, monitoramento de perfil hepático e lipídico, e consentimento informado.
- [ ] 6. Desenvolver materiais educativos para pacientes sobre riscos e desfechos de lipoaspiração, ressaltando a necessidade de mudanças de hábitos para evitar aumento compensatório de gordura visceral.
- [ ] 7. Mapear e atualizar diretrizes internas com evidências recentes (New England Journal of Medicine sobre homocisteína e TVP; meta-análise de oxandrolona em queimados) para embasar solicitações de exames e terapias adjuvantes.
- [ ] 8. Preparar conteúdo para a próxima aula: integração de estratégias nutricionais e metabólicas na bariátrica visando manutenção/ganho de massa magra e prevenção de regressão do peso.

---

### Chunk 15/30
**Article:** Aula 01 Guilherme Sorrentino - Cirurgia e Medicina Funcional Integrativa (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.562

va ultrassensível, homocisteína, e, conforme necessidade, TNF-alfa, CPK e testes relacionados à acidez gástrica e ao metabolismo intestinal. Para o rim, não basta ureia e creatinina—é necessário considerar a reserva muscular (que afeta creatinina e risco cardiovascular). Para o fígado, a leitura vai além de TGO/TGP/bilirrubinas, avaliando capacidade de detoxificação e suporte ao metabolismo de fármacos, cicatrização e enzimas alimentares. O estado nutricional é descrito como fator transversal que impacta todos os demais. A coagulação deve ser mapeada tanto para sangramento intraoperatório quanto para trombose no pós-operatório. O perfil inflamatório é eixo crítico de decisão; o cirurgião relata não operar sem ferritina, pelo menos, e defende uma prescrição pré-cirúrgica que inclua suplementação, orientação nutricional e, quando indicado, adiamento planejado.

---

### Chunk 16/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XVIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.558

r quem decora números, abaixo de 6).
   - Após 75 g de glicose, glicose vai a 216 mg/dL; 2 horas depois glicemia 172 mg/dL; insulina 2 horas 70 µU/mL, permanecendo elevada por longo período, acompanhada de fadiga. Indica hiperglicemia pós-prandial significativa e hiperinsulinemia tardia, com risco de avaliação inadequada se apenas jejum for considerado.
* Implicações perioperatórias
   - Estresse cirúrgico eleva cortisol; uso de corticoide pode somar; oferta de lanches hospitalares típicos aumenta estímulo insulinêmico, levando a pior imunidade, cicatrização, inflamação, oxidação e glicação; maior risco de deiscência, necrose, infecção e feridas.
   - Necessidade de triagem e intervenções para reduzir complicações, não apenas negar cirurgia.
### 3. Estado nutricional e cicatrização
* Prevalência e risco
   - Até 25% dos pacientes ambulatoriais de cirurgia plástica estão em risco de desnutrição e não são avaliados; aproximadamente 1 em cada 4.

---

### Chunk 17/30
**Article:** Aula 01 Guilherme Sorrentino - Cirurgia e Medicina Funcional Integrativa (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.556

ia, identificação de “falso magro” e gordura visceral, que é mais inflamatória). A avaliação nutricional é eixo-chave: muitos pacientes—principalmente mulheres—apresentam intestino com funcionamento subótimo, intolerâncias alimentares (laticínios, glúten) e sensibilidade a FODMAPs (fermentação, gases), constipação, diarreia, permeabilidade aumentada e disbiose. Nesses casos, nutricionistas e gastroenterologistas com experiência em metabolômica podem ser decisivos; exames avançados (p.ex., GI-MAP, gut check, nutrigenética) podem elucidar causas de evolução desfavorável em cirurgias maiores. Na sepse, a perda da homeostase intestinal favorece proliferação de patógenos e agrava inflamação sistêmica, dificultando recuperação—por isso, nutrir e restaurar o intestino é parte da estratégia de cura.

---

### Chunk 18/30
**Article:** Fisiologia e Bioquímica do Sistema Imune II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.554

olicitar e registrar parâmetros essenciais (PCR-us, VHS); calcular índices estimáveis (HGI, TAIG); considerar complementares (TNF-α, IL-6, GPx, MDA, antioxidantes totais).
- [ ] 11. Monitorar trimestralmente HbA1c, glicemia de jejum e HGI; estabelecer metas individuais (ex.: reduzir de 6,1 para ~5,5; longo prazo ~5,3).
- [ ] 12. Calcular TG/HDL e integrar com lipidograma/SREBP1c/2; evitar combinação de gordura saturada com açúcar e excesso de saturadas em múltiplas refeições.
- [ ] 13. Avaliar ferro, ferritina, transferrina e saturação (20–50%; evitar <20%); interpretar com hepcidina/SREBP1c e quadro inflamatório.
- [ ] 14. Medir TNF-α (<8,1) e IL-6 (<3,4) para acompanhar atividade inflamatória; relacionar com obesidade inflamada.
- [ ] 15. Calcular HOMA-β (167–175) e HOMA-IR (<2,15); buscar glicemia 60–90 e insulina ~6–7.
- [ ] 16. Monitorar homocisteína (<7,9) e PCR; usar PCR-us; documentar crises (PCR >1.000) e conduzir manejo apropriado.
- [ ] 17.

---

### Chunk 19/30
**Article:** Modulação Intestinal II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.553

do pâncreas antes de prescrever enzimas digestivas.
- [ ] 4. Monitorar marcadores de inflamação de baixo grau, como resistência à insulina (HOMA-IR), homocisteína e proteína C-reativa.
- [ ] 5. Monitorar os níveis de vitamina B12 ao longo da vida, especialmente em pacientes que usam inibidores de bomba de prótons ou bariátricos.
- [ ] 6. Em pacientes com resistência à insulina, avaliar o TMAO sérico para aferir o risco cardiovascular.
- [ ] 7. Para pacientes que utilizam inibidores da bomba de prótons, planejar um desmame cuidadoso para evitar o efeito rebote de hiperacidez.
- [ ] 8. Aplicar o conhecimento sobre os mecanismos de ação (ex: beta-glucana, butirato) para personalizar as intervenções nutricionais de acordo com as necessidades do paciente (ex: horário de administração para controle de saciedade).

---

### Chunk 20/30
**Article:** Aula 01 Guilherme Sorrentino - Cirurgia e Medicina Funcional Integrativa (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.552

gências/transoperatório – Antes da data da cirurgia ou intraoperatório em urgência
  - [ ] Se ferritina 30–100 com transferrina <20% ou PCR >5, manejar anemia/inflamação e considerar adiar cirurgia eletiva – Decisão até o agendamento final
  - [ ] Incluir exames ampliados conforme caso: insulina de jejum, dímero-D, proteína C reativa ultrassensível, homocisteína, TNF-alfa, CPK, testes de acidez gástrica/metabolismo intestinal – Pré-operatório imediato
  - [ ] Avaliar risco cardíaco com ênfase em estresse subclínico e composição corporal (incluindo reserva muscular) – Pré-operatório
  - [ ] Mapear coagulação e risco de trombose; aplicar score de Caprini e considerar fatores pós-pandemia – Pré-operatório
  - [ ] Monitorar intraoperatório para sangramento: usar frequência cardíaca como guia; intervir se >120 e progressiva apesar de reposição – Intraoperatório contínuo
  - [ ] Evitar exceder 6 horas de tempo cirúrgico e evitar excesso de flu

---

### Chunk 21/30
**Article:** Aula 01 Guilherme Sorrentino - Cirurgia e Medicina Funcional Integrativa (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.545

# Aula 01_Guilherme Sorrentino - Cirurgia e Medicina Funcional Integrativa

**Source:** https://web.plaud.ai/share/1d5d1767377464866::YXdzOnVzLXdlc3QtMg

---

# A Abordagem Funcional e Integrativa na Avaliação Pré-Operatória

O Dr. Guilherme Sorrentino apresenta uma abordagem funcional e integrativa para avaliação e preparo pré-operatório, defendendo uma preabilitação sistemática com foco em estado nutricional, perfil inflamatório e função orgânica para reduzir riscos, prevenir complicações e acelerar a recuperação. Ele estrutura a análise em sete pilares, amplia o escopo de exames laboratoriais e descreve condutas práticas para otimização personalizada antes e durante a cirurgia.
------------
## Introdução à Cirurgia Funcional e Integrativa

A apresentação abre com a defesa da medicina funcional integrativa como uma evolução necessária na prática cirúrgica. Segundo o Dr.

---

### Chunk 22/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XVIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.544

mplementar avaliação pré-operatória de resistência insulínica: solicitar glicemia e insulina de jejum para índice HOMA e, quando possível, realizar curva insulinêmica pós-carga de 75 g de glicose.
- [ ] 2. Padronizar triagem nutricional em pacientes de cirurgia plástica: avaliar risco de desnutrição, ingestão proteica, vitamina C, selênio, zinco, cobre e aminoácidos específicos; encaminhar para nutricionista quando necessário.
- [ ] 3. Incluir dosagem de homocisteína na avaliação de risco trombótico, especialmente em cirurgias longas, pacientes de idade avançada, usuárias de anticoncepcionais ou gestantes.
- [ ] 4. Revisar protocolos de alimentação hospitalar pós-operatória para reduzir picos glicêmicos e estímulos insulinêmicos; considerar opções de lanches com baixo índice glicêmico e maior densidade proteica.
- [ ] 5.

---

### Chunk 23/30
**Article:** Microbioma Intestinal III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.544

glutationa S-transferase) e melhorar o fluxo biliar.
## ❓ Perguntas
- [Inserir Pergunta/Confusão]
## 📚 Tarefas
- [ ] 1. Ao avaliar um paciente com qualquer condição crônica, priorizar a modulação do sistema gastrointestinal como parte fundamental do tratamento.
- [ ] 2. Na anamnese, investigar detalhadamente a história pregressa do paciente (parto, amamentação, uso de antibióticos, doenças, medicamentos).
- [ ] 3. Utilizar ferramentas clínicas como a Escala de Bristol e a observação de distensão abdominal para avaliar a saúde intestinal.
- [ ] 4. Considerar a solicitação de um exame coprológico funcional (como o Copromax) para uma avaliação aprofundada da inflamação e função intestinal.
- [ ] 5. Ao iniciar o uso do exame coprológico funcional, entrar em contato com o laboratório para obter suporte na interpretação dos resultados.
- [ ] 6.

---

### Chunk 24/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XVIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.542

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

### Chunk 25/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XIV (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.540

- [ ] 2. Para pacientes com síndrome metabólica, considerar estratégia Mediterrânea organizada; preferir versão com restrição calórica combinada com atividade física, monitorando resultados aos 6 e 12 meses.
- [ ] 3. Solicitar perfil lipídico ampliado em casos selecionados: HDL, triglicerídeos, insulina, PCR, LDL oxidado e, quando indicado, subfracionamento de LDL; evitar decisões baseadas apenas em LDL total.
- [ ] 4. Planejar exames de risco cardiovascular conforme necessidade: score de cálcio coronariano e angiotomografia de coronárias (incluindo avaliação de placas moles) quando o contexto clínico justificar.
- [ ] 5. Revisar e atualizar protocolos internos sobre álcool: remover recomendações de consumo “cardioprotetor”; educar pacientes sobre riscos de câncer e piora do sono; avaliar predisposição/sensibilidade individual.
- [ ] 6.

---

### Chunk 26/30
**Article:** Psiquiatria Metabólica Funcional Integrativa AULA 17 (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.535

se na evidência de meta-análise.
- [ ] 7. Para suspeita de polimorfismo em PGC1-alfa, iniciar jejum intermitente gradualmente, adicionar coenzima Q10, resveratrol, ácido alfa-lipoico, L-carnitina, Rhodiola, e exercícios de resistência antes de avançar para cetogênica.
- [ ] 8. Orientar uso de moduladores de PPAR-γ/α e controle de fome: curcuminoides, ômega-3, antocianinas, ácido hidroxicítrico (500 mg 30 min antes de refeições críticas), chás (verde, hibisco), óleos essenciais cítricos/alecrim (inalação), capsaicina/capsiate.
- [ ] 9. Integrar acompanhamento psicológico que evite vitimização e paternalismo; alinhar expectativas e responsabilidade pessoal no plano terapêutico.
- [ ] 10. Preparar-se para a próxima aula sobre estratégia cetogênica com a Dra. Janaína e para conteúdos sobre estruturação de casos clínicos.

---

### Chunk 27/30
**Article:** Bioquímica Metabólica nos Exercícios - Parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.533

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

### Chunk 28/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XVIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.532

chos cirúrgicos
* Mecanismo e magnitude do risco
   - Resistência insulínica é um dos principais mecanismos que desencadeiam complicações cirúrgicas comuns.
   - Queda da sensibilidade à insulina em 50% após cirurgia aumenta o risco de complicações graves em 5 a 6 vezes e infecções graves em mais de 10 vezes.
* Avaliação adequada
   - Crítica aos protocolos que usam apenas glicemia e nem hemoglobina glicada; muitos não solicitam insulina.
   - Ferramenta sugerida: índice HOMA (Roma, mencionado), solicitando insulina e glicemia em jejum; ideal incluir curva insulinêmica pós-carga de glicose para avaliar resposta dinâmica, não apenas basal.
* Exemplo clínico de curva insulinêmica
   - Caso: glicemia em jejum 101 mg/dL; insulina basal 3 µU/mL (considerada “boa” por quem decora números, abaixo de 6).

---

### Chunk 29/30
**Article:** The De Ritis Ratio: The Test of Time (2013)
**Journal:** Clinical Biochemist Reviews
**Section:** results | **Similarity:** 0.530

e. Curr Opin Clin Nutr 
Metab Care 2013;16:446-52.132. Coss-Bu JA, Sunehag AL, Haymond MW. Contribution 
of galactose and fructose to glucose homeostasis. 
Metabolism 2009;58:1050-8.133. Stanhope KL, Schwarz JM, Havel PJ. Adverse metabolic 
effects of dietary fructose: results from the recent 
epidemiological, clinical, and mechanistic studies. Curr 
Opin Lipidol 2013;24:198-206.134. Brunt EM, Janney CG, Di Bisceglie AM, Neuschwander-
Tetri BA, Bacon BR. Nonalcoholic steatohepatitis: a 
proposal for grading and staging the histological lesions. 
Am J Gastroenterol 1999;94:2467-74.135. Adams LA, Lymp JF, St Sauver J, Sanderson SO, 
Lindor KD, Feldstein A, et al. The natural history of 

Clin Biochem Rev Vol 34 November 2013   129
nonalcoholic fatty liver disease: a population-based cohort study. Gastroenterology 2005;129:113-21.136. Ekstedt M, Franzén LE, Mathiesen UL, Thorelius L, 
Holmqvist M, Bodemar G, et al. Long-term follow-up 
of patients with NAFLD and elevated liver enzymes.

---

### Chunk 30/30
**Article:** Emagrecimento - Parte V (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.530

e 2 da destoxificação hepática.
    - **Silimarina:** Descrita como o mais potente e estudado suplemento para o fígado, com dose de até 300mg.
- **Alimentos e Chás:** Chás (trevo dos prados, dente de leão), suco de repolho, espinafre (rico em ALA), azeite de oliva e broto de brócolis são indicados.
### 6. Ácido Alfa-Lipoico (ALA) no Manejo da DHGNA
- O ALA é chave para o funcionamento hepático, resistência insulínica e diabetes.
- **Funções:** Regenera antioxidantes (Vitamina C, E), aumenta a síntese de glutationa e tem efeito anti-inflamatório.
- **Evidências:** Meta-análises confirmam que o ALA melhora o perfil lipídico (colesterol, triglicerídeos) e reduz marcadores de peroxidação lipídica de forma dose e tempo-dependente.
- **Dosagem:** Prescrever de 300mg (duas vezes ao dia) a 600mg, idealmente em jejum ou em cápsula gastrorresistente.
### 7.

---

