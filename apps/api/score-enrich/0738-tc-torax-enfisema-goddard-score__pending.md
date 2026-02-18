# ScoreItem: TC Tórax - Enfisema (Goddard Score)

**ID:** `c77cedd3-2800-7049-a76a-7aee90cce76c`
**FullName:** TC Tórax - Enfisema (Goddard Score) (Exames - Imagem)
**Unit:** score

**Preparation Metadata:**
- Quality Grade: **FAIR**
- Total Chunks: 30 de 14 artigos
- Avg Similarity: 0.478

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `c77cedd3-2800-7049-a76a-7aee90cce76c`.**

```json
{
  "score_item_id": "c77cedd3-2800-7049-a76a-7aee90cce76c",
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

**ScoreItem:** TC Tórax - Enfisema (Goddard Score) (Exames - Imagem)
**Unidade:** score

**30 chunks de 14 artigos (avg similarity: 0.478)**

### Chunk 1/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.553

ra fenótipo de sibilância.
**Corticosteroides inalatórios: efetivos, mas com riscos hormonais, de crescimento e ósseos que exigem vigilância e individualização.**
- Supressão do eixo HPA: 10% sintomática e até 40% bioquímica; risco aumenta 6x em crianças e 4x em adultos com alta dose por 3–6 meses.
- Supressão com corticoide oral: cursos >2 semanas consecutivas ou >3 semanas em 6 meses elevam risco.
- Eixos de monitoramento: cortisol às 8h da manhã; se normal, reavaliar em 6 meses; no teste com ACTH, resposta deve subir 18 µg/dL; preocupação com valores de cortisol tão baixos quanto 3 mg/dL.
- Tratamento de supressão: hidrocortisona base por 6–12 meses; atrofia suprarrenal pode persistir até um ano após suspensão de inalatórios.
- ICS e crescimento: perda final de ~1 cm; diferença anual de ~0,2 cm; achados em revisão com quase 3.400 crianças, por 12–52 semanas; contínuo vs cromoglicato: ~1 cm a menos.

---

### Chunk 2/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.529

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

### Chunk 3/30
**Article:** Dislipidemias II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.513

dade, inatividade física, pressão arterial e dislipidemia.
- A inflamação é indicada por marcadores como a Proteína C-Reativa (PCR), onde um valor de 5 já é considerado elevado, e a Interleucina 6 (exemplo de paciente com 8.45).
- A resposta anti-inflamatória é medida pela Interleucina 10, com um valor de corte de 3,5 (abaixo disso é um risco) e um exemplo de paciente com 6.44.
- A saúde do endotélio é avaliada pelo óxido nítrico, que deve estar em 8.8; um paciente com 4.8 já apresenta disfunção sistêmica.
- A LDL oxidada é um marcador crítico, com um limite saudável de 133, enquanto pacientes de alto risco podem apresentar valores extremos, como 1000.
**Achados Adicionais**
- Existem cinco parâmetros clássicos para definir a síndrome metabólica.
- O alvo de LDL para pacientes de alto risco, segundo as diretrizes atuais, é de 50.

---

### Chunk 4/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.512

esposta a tratamento; testes de desafio.
* Acompanhamento do controle
  - ACT (5 itens; 5–25 pontos) nas versões pediátrica e adulta.
  - Critérios GINA (4 itens/4 semanas): 0 = controlada; 1–2 = parcialmente; 3–4 = não controlada.
### 3. Risco de remodelamento e progressão
* Inflamação subclínica persistente + broncoespasmo levam a destruição epitelial e remodelamento brônquico, com irreversibilidade e evolução para DPOC.
### 4. Terapêutica tradicional por faixa etária (steps GINA) e adesão
* Princípios
  - ICS e broncodilatadores de curta/longa ação conforme steps; doses baixa/média/alta por tabelas; LABA e eventual LAMA.
  - <5 anos: preferência por baixa dose; se necessário, dobrar (alta dose).
* Adesão em adolescentes
  - Dificuldade elevada; responsabilidade deve ser compartilhada e conduzida pelos pais.
### 5.

---

### Chunk 5/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.503

oramento e classificação do controle devem ser sistemáticos e sensíveis à idade.**
- Ponto de corte VEF1/CVF: 0,8 em adultos; 0,9 em crianças.
- Resposta ao broncodilatador (VEF1): aumento >12% e ≥200 ml no adulto.
- Variabilidade do pico de fluxo expiratório: >10% em adultos, 13% em crianças; medir 2x/dia por 2 semanas.
- Reavaliação da função pulmonar: após 4 semanas de tratamento.
- Asthma Control Test: 5 perguntas, pontuação de 1 a 5, escore total de 5 a 25.
- Classificação GINA: controle parcialmente controlado com 1–2 critérios; não controlada com 3–4; sintomas em duas ou mais vezes/semana orientam gravidade.
- Faixas etárias no manejo/diagnóstico pediátrico: manejo 6–11 anos; menor de 5 anos; diagnóstico <6 anos, <2 anos; “menores de três anos” para fenótipo de sibilância.

---

### Chunk 6/30
**Article:** pediatria funcional integrativa - parte IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.501

ção de 50%, possivelmente por diminuir viroses. As falhas nos estudos podem ser devidas a polimorfismos, vieses, doses inadequadas e níveis alvo insuficientes.
- **Recomendação Prática:** Usar doses diárias, por longos períodos, buscando níveis acima de 60 ng/ml, com cautela em crianças.
### 5. Ômega 3, PCR e Asma
- **Mecanismo:** O EPA inibe a via do ácido araquidônico (diminuindo leucotrienos) e o DHA diminui a infiltração eosinofílica.
- **Evidências:** Em gestantes, doses altas reduziram a incidência de asma nas crianças. Pacientes com ômega-3 index > 8% necessitam de menos corticoide. A falha de meta-análises pode ser explicada por polimorfismos (ex: FADES) que determinam a resposta à suplementação.
- **Contexto Clínico:** O PCR está aumentado em asmáticos não controlados, especialmente na asma neutrofílica (comum em obesos), e o ômega 3 pode ajudar a reduzir essa inflamação.
### 6.

---

### Chunk 7/30
**Article:** Pedro Neuro - Neurologia Funcional Integrativa 1 (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.495

> 60 anos):** Associado ao gene APOE.
    -   **APOE2:** Protetor.
    -   **APOE3:** Risco levemente aumentado.
    -   **APOE4:** Risco aumentado de 3 a 15 vezes. Ter um parente próximo com Alzheimer aumenta o risco de 10% para 30%. Uma cópia do alelo E4 aumenta o risco em 3 vezes; duas cópias (E4/E4) aumentam em 15 vezes. 35% dos pacientes com Alzheimer não possuem o alelo de risco APOE4.
**Exames Laboratoriais e de Imagem ("Cognoscopia"):**
-   **Líquor (Líquido Cefalorraquidiano):** Análise das proteínas tau (fosforilada e total) e beta-amiloide.
-   **Imagem:**
    -   **Ressonância Magnética de encéfalo com volumetria de hipocampo:** Útil para excluir outras causas e avaliar atrofia cerebral, especialmente no hipocampo.
    -   **PET Scan (FDG e beta-amiloide):** Focam no metabolismo cerebral e na deposição de proteína beta-amiloide.
-   **Marcadores Sanguíneos (com metas ótimas):**
    -   **Homocisteína:** Meta < 7 micromols/L.

---

### Chunk 8/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XIII (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.488

co da dislipidemia.
- [ ] 2. Analisar criticamente os estudos sobre o Inclisiran (Cibrava), focando na diferença entre desfechos substitutos (redução de LDL) e desfechos clínicos duros (mortalidade, infarto, AVC).
- [ ] 3. Utilizar o site `the-nnt.com` para pesquisar o NNT e NNH de outros medicamentos prescritos na prática clínica.
- [ ] 4. Refletir sobre a influência da indústria farmacêutica e dos interesses financeiros na prescrição de novos medicamentos.
- [ ] 5. Estudar a diferença entre risco relativo e risco absoluto para interpretar criticamente os achados de estudos científicos.
- [ ] 6. Pesquisar sobre os diferentes tipos de ômega-3 (EPA, DHA, ALA) e suas formulações para entender como a qualidade do suplemento afeta os resultados.
- [ ] 7. Investigar os exames de subfracionamento de LDL e a relação ApoA/ApoB como ferramentas de avaliação de risco cardiovascular mais precisas que o LDL total.
- [ ] 8.

---

### Chunk 9/30
**Article:** Long COVID: Complications, Underlying Mechanisms, and Treatment Strategies (2026)
**Journal:** Revista não identificada
**Section:** results | **Similarity:** 0.487

e in PMC 2023 June 29.
Author ManuscriptAuthor ManuscriptAuthor ManuscriptAuthor Manuscript
	
	
Getulio José Mattos Do Amaral Filho - getfilho@yahoo.com.br - CPF: 034.983.039-88

98. McCarthy B, Casey D, Devane D, Murphy K, Murphy E, Lacasse Y. Pulmonary rehabilitation for chronic obstructive pulmonary disease. Cochrane Database Syst Rev (2015): CD003793. [PubMed: 25705944] 
99. Vancheri C, Kreuter M, Richeldi L, Ryerson CJ, Valeyre D, Grutters JC, et al. Nintedanib with Add-on Pirfenidone in Idiopathic Pulmonary Fibrosis. Results of the INJOURNEY Trial. Am J Respir Crit Care Med 197 (2018): 356–63. [PubMed: 28889759] 
100. Gigante A, Aquili A, Farinelli L, Caraffa A, Ronconi G, Enrica Gallenga C, et al. Sodium chromo-glycate and palmitoylethanolamide: A possible strategy to treat mast cell-induced lung inflammation in Covid-19. Med Hypotheses 143 (2020): 109856. [PubMed: 32460208] 
101. Limen RY, Sedono R, Sugiarto A, Hariyanto TI.

---

### Chunk 10/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.481

isol normal, teste de estímulo com ACTH (alvo ≥18 µg/dL); valores muito baixos podem ocorrer (~3 µg/dL); saliva/urina com limitações em crianças.
* Manejo
  - Hidrocortisona base por 6–12 meses até normalizar cortisol basal; doses de estresse em cirurgias/infecções.
## ❓ Perguntas
- [Inserir Pergunta/Confusão]
## 📚 Tarefas
- [ ] 1. Implementar rotina de avaliação de controle em cada consulta usando ACT (versões pediátrica e adulta) e critérios GINA.
- [ ] 2. Padronizar educação de técnica inalatoria com material/link adequado para cada dispositivo, incluindo uso de espaçador e higiene oral pós-ICS.
- [ ] 3. Mapear e intervir nos fatores ambientais do paciente: reduzir mofo, poeira, pelos de animais e produtos químicos (ex.: evitar amaciantes em roupas de cama).
- [ ] 4.

---

### Chunk 11/30
**Article:** Hipertensão Arterial Sistêmica II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.479

em casos de elevação, considerando polimorfismos de metilação.
- [ ] 6. Medir Lp(a) e considerar terapias: otimização de LDL (incluindo PCSK9i), niacina, vitamina C; avaliar elegibilidade para TRH e, quando disponível, terapias específicas (ex.: lepodisirã).
- [ ] 7. Calcular razão APO-B/APO-A e intervir para mantê-la ≤0,7–0,8 por meio de dieta, atividade física e farmacoterapia lipídica quando indicado.
- [ ] 8. Investigar e tratar deficiências hormonais (testosterona, estrogênio, DHEA-S) com abordagem individualizada e considerar TRH para reduzir riscos cardiovasculares e outros desfechos.
- [ ] 9. Implementar plano integrado de estilo de vida: alimentação anti-inflamatória, cessação de fumo, suporte social, manejo de estresse, higiene do sono (redução de resistência à leptina), atividade física regular.
- [ ] 10.

---

### Chunk 12/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.478

liza a magnitude da asma no Brasil e no mundo, discute critérios diagnósticos clínicos e funcionais (GINA), ferramentas de acompanhamento (ACT e controle por critérios GINA), e o impacto da inflamação crônica e remodelamento brônquico na progressão irreversível para DPOC. Detalha os steps terapêuticos do GINA por faixa etária, a adesão (especialmente em adolescentes), e os fenótipos de sibilância em <6 anos (transitório, persistente não atópico, atópico/asmático), com critérios de risco de asma e destaque de que sibilância por viroses em <3 anos demanda imunostimulação e prevenção, não aumento de ICS. Integra fatores farmacológicos, genéticos (SNP RS591118 em PDGFD), dispositivos e características das drogas que aumentam absorção sistêmica de ICS, fornece protocolos de rastreio (cortisol matinal e teste com ACTH) e manejo da supressão adrenal com hidrocortisona e doses de estresse.

---

### Chunk 13/30
**Article:** Abordagem Funcional Integrativa Aplicada a Cada Área - Parte XIX (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.474

g/kg) e prática de exercícios de resistência para preservar massa muscular.
- [ ] 3. Todos os profissionais: Em doenças crônicas sem causa orgânica clara ou com má resposta ao tratamento, investigar ativamente traumas de infância, estresse crônico e questões emocionais não resolvidas como possível "causa primeira".
- [ ] 4. Terapeutas e psicólogos: Adotar "terapia de precisão", utilizando múltiplas ferramentas e combinando diferentes abordagens terapêuticas para personalizar o tratamento e focar em resultados mensuráveis, em vez de seguir uma única linha teórica por longos períodos.
- [ ] 5. Estudo pessoal: Pesquisar o conceito de "causa primeira" de Aristóteles para aprofundar a lógica de buscar a origem dos problemas.
- [ ] 6. Estudo pessoal: Ler o livro de Bruce Lipton sobre a conexão entre mente e doenças físicas.

---

## SOAP

> Data e Hora: 2025-11-17 16:33:53
> Paciente: 
> Diagnóstico:

## Histórico do Diagnóstico:
1. Histórico Médico: 
2.

---

### Chunk 14/30
**Article:** Fisiologia e Bioquímica do Sistema Imune II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.474

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

### Chunk 15/30
**Article:** SAM (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.472

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

### Chunk 16/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.471

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

### Chunk 17/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.468

eduzir carga inflamatória sistêmica.
* Comorbidades
  - Rinite alérgica (vias aéreas unidas), refluxo (associado a alergia alimentar/obesidade), obesidade (fenótipo neutrofílico), anemia/deficiência de ferro.
* Ferramentas práticas
  - Técnica correta de dispositivos; uso de espaçador; ACT em todas as visitas; espirometria e PFE quando disponíveis (PFE 2x/dia por 2 semanas).
### 10. Critérios para suspeita e rastreio de supressão do eixo HPA
* Quando suspeitar
  - Sintomas compatíveis; uso de alta dose 3–6 meses; crescimento monitorado a cada 6 meses mesmo em baixa dose; corticoide oral ≥2 semanas consecutivas ou >3 semanas em 6 meses; uso concomitante de inibidores de CYP3A4.
* Como rastrear
  - Cortisol às 8:00; se normal, reavaliar em 6 meses; se sintomático com cortisol normal, teste de estímulo com ACTH (alvo ≥18 µg/dL); valores muito baixos podem ocorrer (~3 µg/dL); saliva/urina com limitações em crianças.

---

### Chunk 18/30
**Article:** Aula Afonso Salgado - Sistema nervoso autônomo (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.468

bigeminia, trigeminia) etc.;  
  - também permite avaliar se o sistema reage positivamente a intervenções simples (como respiração profunda). Se não houver melhora com essas manobras, ele considera que não é o momento de prescrever exercícios respiratórios intensivos, sendo necessário começar por outras estratégias.

- **Card Check:**  
  - pode ser utilizado com sensor semelhante a oxímetro, inclusive em crianças;  
  - avalia seis funções fisiológicas principais:
    - oxigenação,  
    - ritmo cardíaco,  
    - flexibilidade vascular (índice de pulsatilidade),  
    - índice de resistividade do vaso,  
    - resistência temporal dos vasos e capacidade de reação (flexibilidade),  
    - reservas de energia nervosa e resposta a estressores;  
  - integra essas informações para estimar o estado psíquico–comportamental, via eixo coração–cérebro.

---

### Chunk 19/30
**Article:** Hipertensão Arterial Sistêmica II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.465

o crônica.
    -   Lipoproteína (a) elevada, um fator de risco genético pró-trombótico e pró-inflamatório.
    -   Desequilíbrios hormonais (baixo estrogênio e testosterona), especialmente na menopausa.
-   **Diagnóstico Suspeito:** Nenhum no momento
## Plano:
-   **Prescrição:** Inserir mais aqui
-   **Próximos Passos/Exames:**
    -   O palestrante defende uma avaliação abrangente que vai além dos fatores de risco clássicos, incluindo:
    -   Dosagem das proporções de Ômega-3 e Ômega-6 (Índice Ômega-3).
    -   Medição do Hormônio D (Vitamina D), com metas de níveis ótimos (ex: >80 ng/mL para cardiopatas, controlando com PTH).
    -   Curva glicêmica e de insulina para detectar resistência à insulina precocemente.
    -   Avaliação da homocisteína.
    -   Medição da lipoproteína (a).
    -   Avaliação da relação ApoB/ApoA.
    -   Avaliação dos níveis hormonais (testosterona, estradiol, DHEA).

---

### Chunk 20/30
**Article:** pediatria funcional integrativa - parte III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.462

e (6-12 meses até normalização) e dose de estresse em infecções/cirurgias.
    *   **Suplementação/Prevenção:** Imunoestimulação, Vitamina D, Ômega 3, Carotenoides (foco em antioxidação e inflamação).
    *   **Exacerbações:** Corticoides orais (ex: Prednisolona).
*   **Próximos Passos e Exames:**
    *   **Monitoramento Respiratório:** Espirometria e aplicação do ACT (Teste de Controle da Asma) a cada consulta. Avaliação da técnica inalatória.
    *   **Monitoramento Endócrino/Crescimento:** Acompanhamento da estatura a cada 6 meses (crianças em uso de CI). Dosagem de cortisol às 8h (rastreio) e Teste de ACTH se sintomático com cortisol normal.
    *   **Investigação:** Monitorar interações com inibidores de CYP3A4 e comorbidades (refluxo, apneia).
*   **Plano de Tratamento de Acompanhamento:**
    *   **Controle Ambiental:** Redução de mofo, poeira, pelos, produtos químicos.

---

### Chunk 21/30
**Article:** pediatria funcional integrativa - parte IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.461

o de magnésio inalatório para crianças de 2-6 anos com exacerbação grave. Para maiores de 6 anos e adultos, pode ser usado 2g EV em caso de falha no tratamento inicial para evitar internação.
- **Uso Preventivo:** Um estudo com 330mg de magnésio por 6 meses mostrou melhora na qualidade de vida e controle da doença, mas sem alteração no VEF1 ou nos níveis séricos de magnésio.
### 4. Vitamina D e Asma
- **Mecanismo:** Níveis baixos (< 30 ng/ml) pioram o controle da asma. A Vitamina D melhora a ação do corticoide e modula a resposta imune (diminui citocinas inflamatórias e aumenta a anti-inflamatória IL-10).
- **Evidências:** Apesar da forte plausibilidade, meta-análises falham em demonstrar que a suplementação reduz exacerbações em adultos. Em crianças, há uma redução de 50%, possivelmente por diminuir viroses. As falhas nos estudos podem ser devidas a polimorfismos, vieses, doses inadequadas e níveis alvo insuficientes.

---

### Chunk 22/30
**Article:** 2024 Guidelines of the Polish Society of Laboratory Diagnostics and the Polish Lipid Association on laboratory diagnostics of lipid metabolism disorders (2024)
**Journal:** Archives of Medical Science
**Section:** other | **Similarity:** 0.461

FR 3059 ml/min/1.73 m2; risk ≥ 5% and < 10% according to Pol-SCORE /high risk according to SCORE2 or SCORE-2-OP for gender and ageModerateRisk < 5% according to Pol-SCORE/low and moderate risk according to SCORE2 or SCORE-2-OP for 
gender and ageLowRisk of < 1% according to Pol-SCORE.1This corresponds to a SCORE2 risk > 25%  e.g.

---

### Chunk 23/30
**Article:** Dislipdemia (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.460

vida e identificar fatores de risco para inflamação e disfunção endotelial (dieta, estresse, sedentarismo).
- [ ] 2. Utilizar a calculadora MESA Risk para determinar o risco cardiovascular dos pacientes e considerar a solicitação do Escore de Cálcio Coronariano antes de prescrever estatinas em prevenção primária.
- [ ] 3. Ao avaliar o risco cardiovascular, solicitar exames avançados como a relação ApoB/ApoA, subpartículas de lipoproteínas e anti-LDL oxidado para uma análise mais aprofundada.
- [ ] 4. Em pacientes de prevenção primária, mesmo com colesterol alto, priorizar mudanças no estilo de vida antes de considerar o uso de estatinas.
- [ ] 5. Reservar o uso de estatinas principalmente para pacientes em prevenção secundária ou de altíssimo risco, focando nos seus efeitos pleotrópicos.
- [ ] 6.

---

### Chunk 24/30
**Article:** Pedro Neuro - Neurologia Funcional Integrativa 1 (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.460

e avaliar atrofia cortical; usar PET-FDG/FBB quando indicado, interpretando com cautela.
- [ ] 8. Implementar intervenção de estilo de vida: dieta mediterrânea, redução de açúcar (≤15 g/dia) e de AGEs; manejo de mofo e toxinas.
- [ ] 9. Realizar “cognoscopia” aos ~45 anos: painel com metas ótimas (homocisteína, vitaminas, vitamina D/E, PCR, minerais, hormônios, metais tóxicos, sono/apneia, intestino, glúten, gordura visceral).
- [ ] 10. Triar e tratar apneia do sono; priorizar sono reparador; considerar EEG se houver suspeita de crises parciais complexas.
- [ ] 11. Medir circunferência da cintura e/ou realizar DEXA/bioimpedância; estabelecer metas (mulheres <89 cm; homens <102 cm).
- [ ] 12. Adotar abordagem multimodal (“cartucho de prata”), integrando controle de inflamação, glicose, fatores vasculares, hormônios, nutrição e hábitos.

---

### Chunk 25/30
**Article:** pediatria funcional integrativa - parte IV (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.457

eiros meses de vida. O uso de antibióticos neste período é um fator de risco importante.
    *   Pacientes asmáticos apresentam disbiose pulmonar, com aumento de proteobactérias.
*   **O Papel dos Fungos (Micobioma)**
    *   A exposição a mofo agrava a asma e a colonização fúngica está associada a maior mortalidade e gravidade. É necessário um equilíbrio, pois o uso de antifúngicos pode piorar o quadro.
*   **Inflamação Sistêmica e PCR**
    *   A Proteína C Reativa ultrassensível (PCR-US) está aumentada em pacientes com VEF1 reduzido e se correlaciona com mais eosinófilos no escarro.
    *   Altas doses de corticoide inalatório podem ter efeito sistêmico, reduzindo o PCR.
### 5. Fenótipos, Endótipos e Asma Grave
*   **Heterogeneidade da Asma**
    *   A asma não é uma doença única, mas um conjunto de fenótipos (alérgica, não alérgica, associada à obesidade, etc.) e endótipos (mecanismos fisiopatológicos específicos).

---

### Chunk 26/30
**Article:** Aula Afonso Salgado - Sistema nervoso autônomo (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.457

tresse, estado psíquico).
  - Neurometria funcional (FDA/Anvisa) para casos complexos.
- Classificação: 81 estados fisiológico–patológicos (estresse agudo/crônico, degenerativo, arritmias).
- Interpretação operacional:
  - Se Valsalva/respiração profunda não melhoram o estado, evitar prescrever exercícios respiratórios de imediato; formular hipóteses alternativas e reavaliar.
## Alostase, carga alostática e envelhecimento
- Alostase: reserva energética para enfrentar estressores físicos/químicos/tóxicos/emocionais; metáfora do “combustível do carro”.
- Carga alostática: desgaste longitudinal do envelhecimento e doenças degenerativas; metas terapêuticas para proteger alostase.
## Coerência cardíaca e benefícios do treino de VFC
- Coerência cardíaca: integração de bem-estar físico, mental, emocional e espiritual; base de prescrição clínica nos EUA.

---

### Chunk 27/30
**Article:** Aula Afonso Salgado - Sistema nervoso autônomo (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.455

, registro de medicamentos/estressores; repetição padronizada (3–5).
- Evidências: revisões sistemáticas/meta-análises e colaborações institucionais sustentam interpretação.
- Educação: bibliografia em medicina autonômica; acesso a abstracts via Academia Brasileira; capacitação em teoria polivagal e vias neuroendócrinas/neuroimunes.
## Exemplos e correlações clínicas
- Caso familiar com diabetes gestacional e componente emocional: necessidade de acompanhamento prolongado.
- Exemplo pós-COVID: broncoespasmo e deambulação difícil; proposta de fotobiomodulação em gânglio simpático da 1ª costela com broncodilatação e menor risco cardíaco.
- Perfis com baixa VFC e baixa reserva fisiológica: suspender exercício vigoroso até recuperar alostase.
## 📅 Next Arrangements
- [ ] Implementar protocolo de VFC com repetição padronizada (3–5 medições) em condições controladas.

---

### Chunk 28/30
**Article:** Cardiologia II (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.455

ata que leva à formação de células espumosas via ativação macrofágica, contribuindo para aterogênese.
## Diagnóstico Principal:
- Avaliação: Educação e análise de risco cardiovascular baseada na qualidade do LDL, relação triglicerídeos/HDL, sinais de resistência insulínica, múltiplas modificações da LDL (oxidação, glicação, inflamação) e possíveis influências genéticas.
- Diagnóstico Suspeito: Nenhum no momento.
## Plano:
- Prescrição: Insira mais aqui
- Próximos Passos/Exames:
  - Solicitar perfil lipídico completo: colesterol total, HDL, LDL (com possibilidade de subfracionamento), triglicerídeos.
  - Dosar LDL oxidada direta; considerar anti-LDL oxidada se a direta não for disponível.
  - Avaliar marcadores de metabolismo de glicose: glicemia de jejum, insulina de jejum, hemoglobina glicada.
  - Considerar ApoA1 e ApoB; calcular razão ApoA/ApoB.

---

### Chunk 29/30
**Article:** 2024 Guidelines of the Polish Society of Laboratory Diagnostics and the Polish Lipid Association on laboratory diagnostics of lipid metabolism disorders (2024)
**Journal:** Archives of Medical Science
**Section:** other | **Similarity:** 0.454

l/min/1.73 m2])Very high Cardiovascular disease documented clinically or by imaging examinations; diabetes mellitus with 
organ damage3 or other major risk factors4,5, early onset type 1 diabetes mellitus lasting > 20 years; chronic kidney disease with eGFR < 30 ml/min/1.73 m2; familial hypercholesterolaemia with cardiovas-cular disease or another major risk factor5; risk ≥ 10% and ≤ 20% according to Pol-SCORE/very high risk according to SCORE2 or SCORE-2-OP for gender and ageHighSigniﬁcantly elevated single risk factor, especially TC > 310 mg/dl (> 8 mmol/l), LDL-C > 190 mg/dl  (> 4.9 mmol/l), or blood pressure ≥ 180/110 mm Hg; familial hypercholesterolaemia without other risk factors; diabetes mellitus without organ damage (regardless of duration)6; chronic kidney disease with eGFR 3059 ml/min/1.73 m2; risk ≥ 5% and < 10% according to Pol-SCORE /high risk according to SCORE2 or SCORE-2-OP for gender and ageModerateRisk < 5% according to Pol-SCORE/low and moderate risk acc

---

### Chunk 30/30
**Article:** 2024 Guidelines of the Polish Society of Laboratory Diagnostics and the Polish Lipid Association on laboratory diagnostics of lipid metabolism disorders (2024)
**Journal:** Archives of Medical Science
**Section:** results | **Similarity:** 0.454

 s modiﬁcation. 1Based on the PSC/PoLA 2024 Guidelines [81].Lipid proﬁle  laboratory reportLipid proﬁle includes a battery of blood serum or plasma tests and calculations aimed at identiﬁ-cation of dyslipidemia as a cardiovascular risk fac-tor, deﬁning the recommendations and  treatment  
monitoring, including: total cholesterol (TC) level, HDL cholesterol level (HDL-C), LDL cholesterol level (LDL-C), non-HDL cholesterol level (non-HDL-C),
 triglyceride (TG) level,  lipoprotein (a) level [Lp(a)] (determined at least once in life  see PCS/PoLA 2024 recommenda-
tions [81]), apolipoprotein B (apoB) level  as indicated.In addition  to the results of measurements and calculations, a lipid proﬁle laboratory report  (Table IX), should include information on how the LDL-C level was determined (calculated/deter-mined), as well as the target (desirable) and alarm-Table IX.

---

