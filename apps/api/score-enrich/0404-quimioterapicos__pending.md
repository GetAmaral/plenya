# ScoreItem: Quimioterápicos

**ID:** `019bf31d-2ef0-7671-b30a-fbb314fff861`
**FullName:** Quimioterápicos (Histórico de doenças - Medicamentos - Uso atual de medicamentos)

**Preparation Metadata:**
- Quality Grade: **GOOD**
- Total Chunks: 30 de 17 artigos
- Avg Similarity: 0.581

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `019bf31d-2ef0-7671-b30a-fbb314fff861`.**

```json
{
  "score_item_id": "019bf31d-2ef0-7671-b30a-fbb314fff861",
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

**ScoreItem:** Quimioterápicos (Histórico de doenças - Medicamentos - Uso atual de medicamentos)

**30 chunks de 17 artigos (avg similarity: 0.581)**

### Chunk 1/30
**Article:** Jejum Intermitente (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.665

idativo e sensibilidade à quimioterapia; ativam autofagia e reduzem hiperativação de mTOR.
- Efeitos sistêmicos: menor toxicidade da quimioterapia em células normais; maior sensibilidade tumoral; melhora entrega/clearance de quimioterápicos; redução de fatores de crescimento e inflamação.
- Protocolo integrativo (sob oncologista): 5 dias de FMD com quimioterapia no 3º dia; mais 2 dias de FMD pós-quimio; 3 dias de vitamina C em alta dose como pró-oxidante; somente com aval e acompanhamento especializado.
### 7. Evidência científica sobre emagrecimento
- Revisões comparativas: em 11 estudos, pelo menos 9 mostraram perda de peso similar entre jejum e restrição contínua; jejum é ferramenta adicional, não superior por si.
- Revisão guarda-chuva (JAMA): maior força estatística para jejum modificado alternado (~25% do GET no dia) e 5-2; qualidade da evidência majoritariamente baixa/muito baixa; eficácia depende de população e medidas associadas.
### 8.

---

### Chunk 2/30
**Article:** Jejum Intermitente (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.662

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
**Article:** Jejum Intermitente (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.646

quercetina, acetil-L-carnitina) conforme indicação profissional.
- [ ] 9. Em oncologia, considerar FMD ou cetogênica como adjuvante à quimioterapia apenas com aval e acompanhamento de oncologista; seguir protocolos específicos (ex.: quimio no 3º dia de FMD; vitamina C em alta dose por 3 dias).
- [ ] 10. Comparar aderência e resultados entre jejum intermitente e restrição calórica contínua; escolher abordagem com maior probabilidade de manutenção pelo paciente.
- [ ] 11. Educar pacientes sobre mecanismos do jejum (sirtuínas, AMPK, mTOR, BDNF, autofagia) para promover compreensão e adesão informada.

---

## Quantitative Data

### Narrativa Quantitativa
O conjunto de métricas revela que o jejum intermitente, especialmente em janelas de alimentação restritas e modelos 5-2, tende a oferecer benefícios cardiometabólicos e perda de peso comparáveis à restrição calórica contínua, com nuances importantes de horário e adesão.

---

### Chunk 4/30
**Article:** Trato Gastrointestinal VI – Intestino Delgado II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.606

uimioterapia e constipação.
- Durante quimioterapia: evitar suplementação pela possibilidade de alimentar tecidos/células indesejadas; no pós-quimio, pode ajudar na recuperação do trato digestivo.
- Evidências para suporte neuropsiquiátrico (eixo intestino-cérebro, SCFAs, nervo vago, eixo HPA); estudo de 2021 citado. Reconhecimento de evidência clínica robusta limitada para suplementação, mas melhora clínica frequentemente observada.
- Uso prático: componente de sachês como veículo e para adesão.
> **Sugestões de IA**
> Boa ponderação de prós e contras. Ofereça critérios: “indicar quando: constipação, pós-quimio, idosos, SII com mucosa fragilizada; evitar quando: quimioterapia ativa”. Acrescente faixas de dose usuais (5–15 g/dia em doses divididas) e tempo típico de uso, com monitoramento de sintomas gastrointestinais.
## Perguntas dos Alunos
Nenhuma pergunta foi levantada pelos alunos.

---

### Chunk 5/30
**Article:** Dieta Cetogênica - Parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.606

atamento metabólico do câncer
   - Terapias de pressão (contínuas): dieta cetogênica, cetonas exógenas, suplementos/fitoterápicos/drogas individualizadas, manejo do estresse emocional.
   - Terapias de pulso (intermitentes): inibição de glicose, inibição de glutamina, oxigenoterapia hiperbárica, entre outras.
   - Abordagem integrada e personalizada para maximizar o controle tumoral.
* Ensaio clínico randomizado (2021) em câncer de mama
   - 80 pacientes tratados com quimio; randomização para dois grupos; intervenção cetogênica/metabólica por 12 semanas; exames laboratoriais e de imagem no início e 12 semanas; cirurgia e reestadiamento para doença localmente avançada após quimio.
   - Resultados: redução de TNF-α, IGF-1, insulina; aumento de IL-10; redução significativa do tamanho tumoral no grupo cetogênico.

---

### Chunk 6/30
**Article:** Genética e Epigenética II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.604

nas que impactem acetilação de histonas, metilação e reparo/dano do DNA.
- [ ] 2. Implementar estratégias para aumentar AGCC (fibras fermentáveis, modulação da microbiota) com protocolos de prescrição e monitoramento.
- [ ] 3. Avaliar status mitocondrial (sinais clínicos, exames indiretos) e intervir em cofatores (NAD/B3, FAD, alfa-cetoglutarato) conforme necessidade e segurança.
- [ ] 4. Em oncologia (p.ex., quimioterapia), monitorar homocisteína e manter doadores de metil em níveis normais; documentar racional e acompanhamento.
- [ ] 5. Para depressão refratária, considerar metilfolato em doses altas (200–1.000 mcg, podendo 2.000 mcg; em casos específicos, titulação até 15 mg), com monitoramento clínico e laboratorial.
- [ ] 6. Elaborar planos de exercício individualizados: definir faixas de FC, escolher modalidades (força/resistência) que promovam remodelamento muscular e biogênese mitocondrial; ajustar conforme nutrição e status hormonal.
- [ ] 7.

---

### Chunk 7/30
**Article:** Bases Metabólicas das Doenças Crônicas e Gerenciamento - Oxidação 1 (2024)
**Journal:** Pos Graduacao MFI
**Section:** methods | **Similarity:** 0.603

cima.
- Sugestões de IA:
  - Organização: “Folha de bolso” com alimentos-chave, doses, metas e regra Zn:Cu.
  - Métodos: Ensinar leitura de rótulos de suplementos (identificar “quelado” e dosagens).
  - Clareza: Trazer faixas numéricas por laboratório em aula de exames.
  - Melhoria: Alertas sobre selenose e interferências do zinco no cobre (anemia/neutropenia).
### 7. Uso terapêutico do estresse oxidativo e cautelas em oncologia
- Não adianta “entupir” de antioxidantes; o processo oxidativo é fisiológico e necessário.
- Em quimioterapia, geralmente não se prescrevem antioxidantes porque o tratamento é oxidante; decisão deve ficar a critério de especialistas experientes.
- Encaminhar para oncologistas integrativos quando houver dúvida.
- Sugestões de IA:
  - Organização: Quadro “quando evitar antioxidantes” (quimio específica, radioterapia).
  - Métodos: Referências/protocolos de sociedades oncológicas.

---

### Chunk 8/30
**Article:** Bases Metabólicas das Doenças Crônicas e Gerenciamento - Oxidação 1 (2024)
**Journal:** Pos Graduacao MFI
**Section:** methods | **Similarity:** 0.598

Organização: Quadro “quando evitar antioxidantes” (quimio específica, radioterapia).
  - Métodos: Referências/protocolos de sociedades oncológicas.
  - Clareza: Manter nuance de caso a caso com exemplos.
  - Melhoria: Roteiro de comunicação para pacientes que perguntam sobre suplementos durante quimioterapia.
### 8. Consequências vasculares e metabólicas da ROS
- ROS promovem crescimento de musculatura lisa, aumentam inflamação (via NF-κB) e reduzem óxido nítrico, comprometendo vasodilatação/vasoconstrição.
- Hiperglicemia, hiperlipidemia, hiperinsulinemia e hipertensão geram disfunção endotelial → trombose, inflamação, vasoconstrição.
- Desbalanço de ROS causa pane no sistema antioxidante, dano ao DNA, altera HDL e adipocinas; bola de neve para doenças crônicas (cardiovasculares, obesidade, câncer, diabetes; também autoimunes).
- Sugestões de IA:
  - Organização: Mapa “metabólico → endotelial → clínico”.

---

### Chunk 9/30
**Article:** Jejum Intermitente (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.588

 ão; 18 g mostraram tendência de melhora maior, sugerindo dose-resposta.
- Posologia prática: iniciar após pelo menos 12 horas de jejum; começar com 5 g de C8 e titrar até 15 g, com máximo sugerido de 20 g, para equilibrar benefício e tolerabilidade gastrointestinal.
**O horário das refeições importa: jantar cedo favorece o emagrecimento independentemente de dieta e atividade física.**
- Grupos que jantaram às 7h/7h30 perderam 30% mais peso que grupos que jantaram tarde (ex.: 10h30), indicando efeito temporal robusto.
- Este efeito reforça o alinhamento circadiano das janelas alimentares como componente crítico da eficácia do jejum.
**FMD integrada à quimioterapia tem protocolo preciso para tentar proteger células normais e potencialmente aumentar eficácia antitumoral.**
- Prescrição clínica: 5 dias de FMD (dieta que mimetiza o jejum) com quimioterapia no 3º dia; manter FMD por mais 2 dias após.

---

### Chunk 10/30
**Article:** Dieta Cetogênica - Parte I (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.584

a atletas de alta performance (adaptação prolongada, tolerância maior a carboidratos mantendo cetose, referências a Phinney, Volek e estudo FASTER).
No âmbito oncológico, são revisados estudos mostrando viabilidade e segurança da cetogênica durante quimio/radioterapia, com preservação de massa magra, perda de gordura total e visceral, melhor adesão e satisfação. Em cardiometabolismo e microbiota, há melhora de triglicerídeos, HDL, redução de LDL pequeno/denso e ApoB/ApoA, redução de inflamação e otimização da relação Bacteroidetes/Firmicutes. A palestra também destaca o papel do óleo de TCM (preferência por C8/C10, evitar C6), cetonas exógenas, inibidores de SGLT2 e individualização para diferentes objetivos (terapia, emagrecimento, esporte).
## 🔖 Pontos de Conhecimento
### 1. História e evidência da dieta cetogênica
- Jejum em epilepsia (1911–décadas seguintes):
  - Paris (Guelph e Marie): 20 crianças com excelentes resultados.

---

### Chunk 11/30
**Article:** Medicina Baseada em Evidência III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.582

s sob seus cuidados e planejar reavaliação de necessidade e risco/benefício, com foco em redução quando apropriado.
- [ ] 5. Preparar material de consentimento informado que compare riscos e benefícios de opções terapêuticas (p. ex., cirurgia vs nova quimioterapia), incluindo probabilidades de desfechos e incertezas.
- [ ] 6. Implementar intervenções de baixo risco com plausibilidade mecanística e múltiplos benefícios (ex.: curcumina, ômega-3) quando apropriado, monitorando desfechos clínicos (p. ex., dor).
- [ ] 7. Investigar casos clínicos relevantes (ex.: cetogênica e cetose, relato da doutora Janaína) e documentar resultados, contextualizando a ausência de “nível A” formal em abordagens personalizadas.
- [ ] 8. Desenvolver um roteiro de comunicação para pacientes que mitigue o viés de autoridade, promovendo compreensão crítica de estudos e alinhamento com valores e preferências individuais.
- [ ] 9.

---

### Chunk 12/30
**Article:** Microbioma Intestinal IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.581

res
*   **Hidrocolonoterapia e Enema de Café:** O instrutor esclarece que, embora não as pratique rotineiramente, reconhece seus resultados positivos em contextos específicos e com profissionais experientes (cita a Dra. Maria Emília Serra Gadelha). São terapias complexas que exigem preparo e higiene rigorosos.
*   **Terapia de Gerson (Gerson's Therapy):** Abordagem desenvolvida para tratar câncer, conhecida pelo uso de enemas de café. O instrutor relata um caso de paciente com câncer metastático que segue o protocolo há mais de 12 anos com bons resultados. A terapia é descrita como um modelo antigo e pouco individualizado.
*   **Tratamento Metabólico do Câncer:** Abordagem que foca no metabolismo celular como alvo terapêutico, sendo um tema pouco discutido na medicina tradicional.
### 4. Probióticos como Ferramenta Terapêutica
*   **Definições:**
    *   **Probióticos:** Micro-organismos vivos que, em quantidades adequadas, conferem benefício à saúde.

---

### Chunk 13/30
**Article:** Emagrecimento - Parte IX (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.580

l primeiro se asseguram nutrientes essenciais (via exames e individualização) e, a seguir, se avaliam fitoterápicos com evidências específicas e foco em adesão, qualidade e praticidade.
No âmbito dos micronutrientes, discute decisão por exames, doses, formas, marcas e qualidade de manipulação no Brasil versus EUA, destacando vitamina D e K2, iodo (kelp e Ascophyllum nodosum/I-plus), cromo, vanádio, selênio, zinco/cobre, magnésio, além de B12 e B3 (niacinamida/Niagen). Reforça evitar multivitamínicos não individualizados e orientar pacientes sobre padronização, fator de correção e comparação de preços com transparência.
Entre os fitoterápicos prioritários, inclui berberina HCl (especialmente em disbiose), canela do Ceilão, gimnema silvestre, ácido hidroxicítrico padronizado (Citrimax) e ginostema pentaphyllum (gipenosídeos/actiponina), discutindo doses, padronizações e sinergias com cromo e B3.

---

### Chunk 14/30
**Article:** Mitocôndrias - Parte IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** methods | **Similarity:** 0.579

jejum intermitente, exposição ao frio, modulação calórica, além de suplementação específica.
- [ ] 10. Considerar a necessidade geral de suplementação devido à baixa densidade nutricional dos alimentos modernos (dados de 2005–2016); ajustar protocolos dietéticos e suplementares.
- [ ] 11. Em pacientes oncológicos em quimioterapia, evitar suporte antioxidante/nutricional que possa interferir; reavaliar após término da quimioterapia.
- [ ] 12. Preparar para a próxima aula: revisar metabolismo do ferro, métodos de avaliação da homeostase férrica e estratégias de restauração; focar especialmente em mulheres (estimativa de até 90% com estoques inadequados).

---

## Concept Insights

### Dano Mitocondrial Precede a Anemia
**Categoria:** Princípio Diagnóstico
**Definição Central:**
A disfunção mitocondrial e o dano ao DNA são consequências precoces da deficiência de ferro e outros micronutrientes essenciais.

---

### Chunk 15/30
**Article:** Emagrecimento - Parte IX (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.573

valiar aporte e objetivos de médio prazo considerando dieta e adesão.
### 5. Hierarquia terapêutica, disbiose e pré-refeição
- Primeiro corrigir nutrientes essenciais e estratégia alimentar; depois fitoterápicos.
- Em obesos/sobrepeso, disbiose é comum: preferir berberina HCl antes das refeições; adicionar cromo, vanádio; considerar 5-HTP (25–50 mg) e L-teanina (200 mg) para ansiedade, equilibrando número de cápsulas.
- Canela do Ceilão: 1 colher de café no “shot” matinal ou café.
### 6. Evidências de fitoterápicos
- Gimnema silvestre: revisão sistemática e meta-análise (2021, 10 estudos, N=419) mostra redução de glicemias, HbA1c, TG e colesterol em T2DM; dose 200–300 mg antes das refeições.
- Ácido hidroxicítrico (HCA)/Citrimax: usar padronizado; efeitos em leptina e GLUT1/GLUT4; 500 mg antes das refeições; caro e aumenta cápsulas; melhor com B3, cromo e gimnema.

---

### Chunk 16/30
**Article:** MFI PÓS RACHEL GAIGER AULA 01 - OXIGÊNIO HIPERBÁRICO (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.573

cofatores mitocondriais (CoQ10, L-carnitina, ácido alfa-lipoico, glutationa, glicina, taurina, tirosina, MCT, vitaminas B1/B2/B3).
- Integração com terapias padrão e medicina funcional: redução de sintomas para permitir abordagem da causa raiz; sinergia com antibióticos/antifúngicos, esteroides em DII e reabilitação motora/cognitiva em casos neurológicos.
### Segurança, contraindicações e manejo de efeitos adversos
- Contraindicações relativas: doenças pulmonares crônicas (enfisema), pneumotórax recente, cirurgia cardíaca/trauma recente, neurite óptica, uso atual de bleomicina/doxorrubicina, otite média/dificuldade de equalização, implante coclear, marcapasso (avaliação individual), hipercapnia, hipertensão não controlada, transtornos convulsivos.
- Absolutas: pneumotórax, broncoespasmo agudo não resolvido, próteses orbitais específicas.
- Drogas incompatíveis: cisplatina, doxorrubicina; cautela com amiodarona e antiangiogênicos.

---

### Chunk 17/30
**Article:** Dieta Cetogênica - Parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.568

vidas e respostas: resposta parcial/total em 1 ano de 27%; sobrevida média 3 anos de 44%; câncer de mama: 4/5 com sobrevida média de 41 meses; colorretal: 19 meses; pancreático: 10 meses; pulmão: 81 meses e ainda vivo.
   - Comparações de referência: pâncreas estágio 4: sobrevida 5 anos ~3%, usual 8–10 meses; pulmão: 5% e 19 meses com tratamento moderno; resultados sugerem benefício potencial.
   - Efeitos adversos: 1143 eventos relatados; 888 atribuídos à quimio/tumor; 275 à dieta, todos leves; 12 perderam gordura sem perda de massa magra; dieta considerada segura e bem tolerada em terminais.
   - Conclusão: a cetogênica é terapia adjuvante promissora com boa adesão e reprodutibilidade, merecendo ensaios clínicos randomizados controlados.
* Conceito Press-Pulse no tratamento metabólico do câncer
   - Terapias de pressão (contínuas): dieta cetogênica, cetonas exógenas, suplementos/fitoterápicos/drogas individualizadas, manejo do estresse emocional.

---

### Chunk 18/30
**Article:** Genética e Epigenética II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.567

s sempre sob controle.
- Para pacientes em quimioterapia, os níveis de doadores de metil devem ser mantidos normais, nunca elevados.
- A mentalidade de "quanto mais, melhor" é perigosa e incorreta na medicina funcional.
### 5. Exercício Físico como Modificador Epigenético
- O exercício físico é um poderoso modificador epigenético, essencial para a biogênese mitocondrial, remodelação muscular e longevidade, além de ser crucial na prevenção de doenças como o Alzheimer.
- A execução correta é fundamental: o tipo, intensidade, duração e frequência devem ser individualizados.
- A simples recomendação de "fazer musculação" é insuficiente; a forma, o preparo, a alimentação e o estado hormonal são cruciais para obter os benefícios.
### 6.

---

### Chunk 19/30
**Article:** Aula Jéssica Marques - Neurologia Funcional Integrativa 3 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.564

agnésio, vitamina B12, folato (B9), vitamina D e ferritina/ferro; corrigir deficiências conforme resultados.
- [ ] 3. Planejar introdução gradual de suplementação: iniciar uma formulação, aguardar 2–3 dias, adicionar a segunda e, após 2–3 dias, a terceira, monitorando efeitos colaterais.
- [ ] 4. Implementar ciclos de 60–90 dias entre plantas com mecanismos semelhantes, trocando para outra família após cada ciclo para sensibilizar diferentes receptores.
- [ ] 5. Selecionar extratos padronizados com maior biodisponibilidade (ex.: Curcumin C3, Cureit, CurcuVail) e evidência clínica para compor formulações.
- [ ] 6. Incorporar vias alternativas para idosos: avaliar uso de injetáveis, transdérmicos, fotobiomodulação, aromaterapia e tinturas conforme perfil e poder aquisitivo.
- [ ] 7. Estruturar rotina diária do paciente com atividades úteis e significativas (cozinhar, organizar casa, acompanhar netos/filhos), promovendo pertencimento e utilidade.
- [ ] 8.

---

### Chunk 20/30
**Article:** Fisiologia e Bioquímica do Sistema Imune II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.563

 gicos de intervenção
- Ao acordar: “shot” concentrado de ativos.
- Tarde (17:00–18:00): adaptógenos + anti-inflamatórios naturais (Boswellia, cúrcuma).
- Noite: ativos que modulam PPAR-γ (fontes de apigenina) para reduzir inflamação, cravings e favorecer melatonina; jantar cedo recomendado.
### 5. Jejum Intermitente e Time Restricted Feeding (TRF)
- Cetogênese inicia ~12h; janelas de 16–18h geram 4–6h de cetogênese útil com menor pico insulinêmico.
- Insulina alta relaciona-se com IL-6 e COX-2; meta de insulina <6 em autoimunes/inflamatórios.
- Protocolos: 18h de jejum com 2–3 refeições no pós-jejum; janelas TRF como 08:00–14:00 ou 08:00–15:00.
### 6. Fasting Mimicking Diet (FMD)
- Protocolo de 5 dias, 100% vegano, baixa carga glicêmica; modula células dendríticas e interleucinas; aplicável em diabetes, câncer, DCV e anti-aging.
- Periodicidade: cada 1–4 meses conforme estado clínico e crises.
### 7.

---

### Chunk 21/30
**Article:** Dieta Cetogênica - Parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.563

 sicos.
- **Alzheimer (caso da apresentadora):** Acamado, totalmente dependente, dor intensa em membros inferiores por neuropatia, insônia, apatia e sonolência.
- **Câncer (caso 5):** Efeitos colaterais múltiplos e intoleráveis das quimioterapias.
## Objetivo:
Resultados de estudos e relatos clínicos:
- **Acne:** Dietas de alto índice glicêmico elevam insulina e IGF-1, estimulando andrógenos, proliferação de queratinócitos e colonização por *Propionibacterium acnes*.
- **Enxaqueca:** Em um estudo, dieta cetogênica por 1 mês reduziu ataques mensais de 2,9 para 0,7, dias com dor de 5 para 0,9 e uso de analgésicos de 5 para 0,5. Em enxaqueca refratária, dias sintomáticos caíram de 30 para 7,5/mês, com dor significativamente menor.
- **Câncer:**
    - **Caso 1 (Pulmão):** PET limpo 46 meses após rádio, quimio e dieta cetogênica.
    - **Caso 2 (Cólon):** Doença estável 6,5 anos após quimioterapia reduzida + dieta cetogênica.

---

### Chunk 22/30
**Article:** Medicina Baseada em Evidência III (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.562

# Medicina Baseada em Evidência III

**Source:** https://web.plaud.ai/share/e5d61763842373102::YXdzOnVzLXdlc3QtMg

---

## Lecture

> Data e hora: 2025-11-17 17:09:26
> Local: [Inserir Local]
> Instrutor: [Inserir Nome do Instrutor]
## 📝 Resumo
A aula, parte do curso de Medicina Funcional Integrativa da Academia Brasileira de Medicina Funcional Integrativa, discute criticamente a avaliação de evidências científicas e a aprovação de terapias, com foco especial em oncologia. Enfatiza a integração entre julgamento clínico, preferências/valores do paciente, evidência relevante e individualidade, bem como a leitura crítica de meta-análises e revisões sistemáticas, considerando plausibilidade e vieses.

---

### Chunk 23/30
**Article:** Bases Metabólicas das Doenças Crônicas e Gerenciamento - Inflamação 2 (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.561

resistência insulínica. Apresenta ensaios clínicos e meta-análises que demonstram redução de PCR-us, IL-6 e LDL/triglicerídeos, além de melhora de HDL, FRAP/TRAP, HOMA-IR, adiponectina e BHB. Aborda a anemia da inflamação e suas diferenças laboratoriais em relação à deficiência de ferro. Propõe uma abordagem integrada de prevenção e manejo que combina personalização dietética (low carb, cetogênica, mediterrânea, plant-based), suplementação baseada em evidência (EPA/DHA, curcumina padronizada com piperina ou lipossomada, antocianinas padronizadas, polifenóis diversos), modulação do tônus parassimpático e atividade física para proteção metabólica e imunológica. Destaca a importância do oncologista e do cardiometabologista preventivos na medição sistemática de PCR-us e na intervenção proativa. Antecipação: próxima aula focará epigenética/metilação e exames correlatos.
## 🔖 Pontos de Conhecimento
### 1.

---

### Chunk 24/30
**Article:** Dieta Cetogênica - Parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.559

o mensal de analgésicos de 30 para 6 doses.
- A intensidade da dor também diminuiu drasticamente: enquanto 86% dos pacientes inicialmente relatavam dor de nível máximo (3), após a dieta, 55% passaram a relatar dor de nível mínimo (1).
- A eficácia não é nova, com um estudo de 1928 já relatando 40% de melhora em pacientes com enxaqueca.
**Estudos de caso e ensaios clínicos indicam que a terapia cetogênica, como adjuvante, melhora significativamente os resultados em pacientes com câncer avançado, aumentando a sobrevida e a resposta ao tratamento padrão.**
- Em um estudo com 37 pacientes com tumores avançados, a sobrevida média em 3 anos foi de 44%, superando as taxas de sobrevida típicas para cânceres como pâncreas (3% em 5 anos) e pulmão (5% em 5 anos).
- Pacientes no estudo alcançaram sobrevidas notáveis: 81 meses para um paciente com câncer de pulmão e 41 meses para pacientes com câncer de mama.

---

### Chunk 25/30
**Article:** Síndrome dos Ovários Policísticos - Parte II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.557

 mica.
    *   **Dosagem e Efeitos Colaterais:** A dose usual é de até 1.500-1.700 mg/dia. Efeitos gastrointestinais são comuns. O uso crônico pode causar deficiência de vitamina B12, exigindo monitoramento. Contraindicada em insuficiência renal.
*   **Progesterona Micronizada**
    *   Essencial para mulheres com SOP que não produzem progesterona adequadamente.
    *   **Função:** Restabelece a regularidade menstrual, protege contra o câncer de endométrio e é crucial para a fertilidade.
    *   **Protocolo:** Usada de forma cíclica (ex: 10-14 dias por mês), na dose de 200 mg. A via oral pode melhorar o humor e o sono devido à produção de alopregnenolona.
    *   Pode causar acne, que pode ser manejada com antiandrogênicos como espironolactona e saw palmetto.
### 4. Suplementação e Terapias Adjuvantes
*   **Inositol (Mio-inositol e D-chiro-inositol)**
    *   Considerado um tratamento de ponta, atua como segundo mensageiro da insulina.

---

### Chunk 26/30
**Article:** Emagrecimento - Parte VIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.548

elevância clínica.
- Boswellia padronizada entrega mesma eficácia com menos cápsulas, favorecendo adesão.
- Suplementos lipídicos devem ser tomados com refeições para melhor absorção e conforto gástrico.
### Alavancas clínicas complementares
Protocolos simples e personalizados maximizam resultados em dor, inflamação e emagrecimento.
- Inalação direta supera difusão ambiental para efeitos terapêuticos de óleos essenciais.
- Beta-cariofileno da copaíba ativa CB2 e favorece analgesia e modulação inflamatória.
- Otimizar vitamina D melhora resistência insulínica e marcadores inflamatórios, com doses individualizadas por polimorfismos GC/VDR.

---

### Chunk 27/30
**Article:** Bases Metabólicas das Doenças Crônicas e Gerenciamento - Inflamação 2 (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.548

la/maçã/groselhas/batata roxa/berinjela/repolho roxo/rabanete/vagem/cereais.
### 4. Curcumina: evidência, doses e formulações
- Meta-análise de 15 ECR: reduz IL-6, PCR-us e MDA (antioxidante/anti-inflamatória).
- Diferenciar açafrão culinário vs extratos padronizados (95% curcuminoides).
- Formulações/doses: cápsulas 500 mg; 500 mg a 2 g/dia conforme tolerância; piperina 10 mg aumenta biodisponibilidade (avaliar alergia); anticoagulados: ≤500 mg/dia (ou 250 mg lipossomada). Opções lipossomadas/patentes: Cureit, Curveil. Sem piperina quando foco é modulação de microbiota.
### 5. Ômega 3 vs ômega 6: dose e integração dietética
- EPA/DHA são efetores; ALA depende de conversão limitada; preferir óleo de peixe para efeito consistente.
- Doses efetivas frequentemente altas, especialmente se dieta permanece ultraprocessada; integrar antioxidantes e ajustar dieta para incorporação em membranas; individualizar por grau de inflamação/oxidação.
### 6.

---

### Chunk 28/30
**Article:** Aula Jéssica Marques - Neurologia Funcional Integrativa 3 (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.545

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
**Article:** Aula 02 Guilherme Sorrentino - Suplementação em Cirurgia (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.544

ndo que qualquer discussão sobre suplementos é necessariamente parcial, dado que o corpo requer um espectro completo de nutrientes. Embora a gama de opções seja ampla, sustenta que, com um conjunto “básico” de intervenções, já é possível oferecer ganhos clínicos significativos. Define objetivos operacionais claros: acelerar a cicatrização, reduzir risco de infecção e dar suporte ao metabolismo e à função mitocondrial, inclusive auxiliando o fígado em processos de detoxificação. Defende uma estratégia personalizada, orientada por avaliação das individualidades bioquímicas (ex.: o que é indicado para um paciente pode não ser para outro), pois a demanda metabólica imposta pelo ato cirúrgico supera a capacidade da dieta habitual em suprir necessidades “ótimas”.

---

### Chunk 30/30
**Article:** Dieta Cetogênica - Parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.543

, Parkinson, ELA, endometriose, psiquiatria, caquexia) e estabelecer protocolos-piloto com métricas específicas de resultado.
- [ ] 8. Revisar segurança e evidências de ozonioterapia, terapias fotodinâmica/sonodinâmica e hipertermia como adjuvantes, propondo comitê de ética para estudos observacionais.

---

## Quantitative Data

### Narrativa Quantitativa
A análise de múltiplos estudos de caso e ensaios clínicos revela um padrão consistente: intervenções metabólicas, como a dieta cetogênica e a suplementação de cetonas, demonstram um potencial terapêutico significativo em condições neurológicas e oncológicas. Os dados indicam melhorias drásticas em desfechos de enxaqueca, sobrevida e qualidade de vida em pacientes com câncer avançado, e reversão de sintomas em casos de Alzheimer, sugerindo uma via metabólica comum que pode ser explorada para tratamento.

---

