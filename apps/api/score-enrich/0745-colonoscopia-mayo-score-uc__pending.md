# ScoreItem: Colonoscopia - Mayo Score UC

**ID:** `c77cedd3-2800-7ab7-a81f-2cf21706bcee`
**FullName:** Colonoscopia - Mayo Score UC (Exames - Imagem)
**Unit:** score

**Preparation Metadata:**
- Quality Grade: **GOOD**
- Total Chunks: 30 de 21 artigos
- Avg Similarity: 0.525

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

**O JSON deve obrigatoriamente conter o campo `score_item_id` com o valor `c77cedd3-2800-7ab7-a81f-2cf21706bcee`.**

```json
{
  "score_item_id": "c77cedd3-2800-7ab7-a81f-2cf21706bcee",
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

**ScoreItem:** Colonoscopia - Mayo Score UC (Exames - Imagem)
**Unidade:** score

**30 chunks de 21 artigos (avg similarity: 0.525)**

### Chunk 1/30
**Article:** Assessment of Endoscopic Disease Activity in Ulcerative Colitis: Is Simplicity the Ultimate Sophistication? (2021)
**Journal:** Inflammatory Intestinal Diseases
**Section:** abstract | **Similarity:** 0.737

Review of endoscopic scoring systems for ulcerative colitis, comparing Mayo Endoscopic Score (MES) with newer indices like UCEIS and UCCIS. Highlights limitations of MES in depicting segmental healing and measuring partial therapeutic responses.

---

### Chunk 2/30
**Article:** Mayo Score/Disease Activity Index (DAI) for Ulcerative Colitis (2024)
**Journal:** MDCalc Clinical Decision Tools
**Section:** abstract | **Similarity:** 0.676

Clinical calculator and reference guide for Mayo Score assessment in ulcerative colitis. Provides scoring criteria (0-12 scale), endoscopic subscore definitions, and clinical interpretation guidelines for disease activity monitoring.

---

### Chunk 3/30
**Article:** The Ulcerative Colitis Endoscopic Index of Severity More Accurately Reflects Clinical Outcomes and Long-term Prognosis than the Mayo Endoscopic Score (2015)
**Journal:** Journal of Crohn's & Colitis
**Section:** abstract | **Similarity:** 0.617

Comparative study demonstrating UCEIS superior ability to detect early mucosal healing and predict long-term outcomes compared to Mayo Endoscopic Score in 41 UC patients receiving tacrolimus therapy.

---

### Chunk 4/30
**Article:** Intolerâncias, Alergias e Hipersensibilidades Alimentares II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.567

odem indicar tolerância e não alergia.
    - **Endoscopia/Colonoscopia:** Podem revelar achados como hiperplasia nodular linfoide.
### 5. Manejo e Tratamento
*   **Dietas de Eliminação:** Principal abordagem, consiste em retirar o alimento agressor. Deve ser feita com acompanhamento multidisciplinar para evitar déficits nutricionais, especialmente em crianças.
*   **Melhora da Digestão:** Uma digestão inadequada aumenta a carga de antígenos no intestino. O uso de enzimas digestivas pode ajudar a degradar melhor as proteínas e diminuir os sintomas. Fatores como pasteurização e Reação de Maillard podem aumentar a alergenicidade dos alimentos.
*   **Modulação Intestinal:** É o pilar do tratamento.
    - **Microbiota e AGCC:** Uma dieta rica em fibras aumenta a produção de ácidos graxos de cadeia curta (butirato, propionato), que são anti-inflamatórios e fortalecem a barreira intestinal.
    - **Probióticos:** Cepas específicas como *Lactobacillus rhamnosus*, *L.

---

### Chunk 5/30
**Article:** Intolerâncias, Alergias e Hipersensibilidades Alimentares II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.545

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

### Chunk 6/30
**Article:** Breastfeeding Beyond Six Months: Evidence of Child Health Benefits (2024)
**Journal:** Nutrients
**Section:** other | **Similarity:** 0.542

hinol. Allergy 2020, 34, 632–641.
[CrossRef] [PubMed]
Ng, S.C.; Shi, H.Y.; Hamidi, N.; Underwood, F.E.; Tang, W.; Benchimol, E.I.; Panaccione, R.; Ghosh, S.; Wu, J.C.Y.; Chan,
F.K.L.; et al. Worldwide Incidence and Prevalence of Inflammatory Bowel Disease in the 21st Century: A Systematic Review of
Population-Based Studies. Lancet 2017, 390, 2769–2778. [CrossRef]
Turner, D.; Ruemmele, F.M.; Orlanski-Meyer, E.; Griffiths, A.M.; De Carpi, J.M.; Bronsky, J.; Veres, G.; Aloi, M.; Strisciuglio, C.;
Braegger, C.P.; et al. Management of Paediatric Ulcerative Colitis, Part 1: Ambulatory Care-An Evidence-Based Guideline from
European Crohn’s and Colitis Organization and European Society of Paediatric Gastroenterology, Hepatology and Nutrition.
J. Pediatr. Gastroenterol. Nutr. 2018, 67, 257–291. [CrossRef] [PubMed]
Veauthier, B.; Hornecker, J.R. Crohn’s Disease: Diagnosis and Management. Am. Fam. Physician 2018, 98, 661–669.
Iyengar, S.R.; Walker, W.A.

---

### Chunk 7/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.537

intestinal.
-   **SIFO:** Sensação de "blurring" (efeito baiacu - sentir-se distendido mesmo sem distensão objetiva) e dor abdominal.
## Objetivo:
A palestra descreve os exames e achados objetivos para a avaliação de pacientes com suspeita de Síndrome do Intestino Irritável (SII), embora não contenha achados de exame físico de um paciente específico.
-   **Exames Laboratoriais Gerais:** Hemograma e marcadores de atividade inflamatória.
-   **Calprotectina Fecal:** Usada para descartar doença inflamatória intestinal. Um valor abaixo de 100 µg/g tem uma positividade de 98% para o diagnóstico de SII. Entre 100-250 µg/g é uma zona cinzenta. Acima de 250 µg/g requer colonoscopia.
-   **Avaliação para Doença Celíaca:** Dosagem de IgA sérica total e anticorpo antitransglutaminase IgA. Se a IgA for baixa, dosar IgG sérica e antitransglutaminase IgG.
-   **Teste Parasitológico de Fezes:** Recomendado no contexto do Brasil.

---

### Chunk 8/30
**Article:** (Dr Otávio Freitas) Aula 02 - Vitamina D - Doenças Autoimunes (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.537

A apresentação expande para doença de Crohn e retocolite ulcerativa (RCU), alinhando observações clínicas do consultório a evidências publicadas: meta-análise de 55 estudos observacionais relaciona deficiência de vitamina D com essas condições; estudos sugerem que a vitamina D atenua a inflamação na RCU por ativar o receptor de vitamina D e modular a resposta NL-RPC; há menções sobre possíveis relações entre níveis de vitamina D e a extensão da doença. O orador cita um paciente acompanhado por cerca de sete anos com colonoscopia normal após tratamento. O depoimento de Juliano ilustra um percurso de 15 anos desde o diagnóstico por exames e cirurgia, com uma década de tratamentos convencionais e dor/desconforto persistentes.

---

### Chunk 9/30
**Article:** Microbioma Intestinal III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.529

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

### Chunk 10/30
**Article:** Microbioma Intestinal III (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.528

reática fecal: 85 — baixa; sugere insuficiência pancreática exócrina leve/moderada, possivelmente secundária a hipocloridria e disfunção digestiva global.
  - Zonulina fecal: 7 (normal < 80) — normal; reduz evidência laboratorial de hiperpermeabilidade via este marcador específico.
- Comentários:
  - Recomendada correlação com parâmetros sanguíneos (PCR, VHS) para reforçar inflamação sistêmica.
  - Colonoscopia citada como método de rastreio em adultos; não indicada para criança neste contexto.
- Mecanismos fisiopatológicos discutidos:
  - Dano a junções estreitas (claudina, ocludina, actina) por dieta (ex.: glúten).
  - Reconhecimento de MAMPs por TLR em células epiteliais; apresentação antigênica por células dendríticas/M e ativação de resposta T.
  - Células de Paneth: estimuladas por IL-22 e beta-glucana; produção de defensinas.
  - Células caliciformes (Goblet): síntese de mucina, principal fator antimicrobiano no cólon.

---

### Chunk 11/30
**Article:** SAM (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.519

ão. O diagnóstico é sugerido por um aumento de 20% + 2 ng/mL acima do valor basal do paciente durante uma crise. No entanto, o palestrante relata que raramente vê resultados positivos.
    -   **N-metil-histamina urinária (urina de 24h):** Considerado um marcador útil. Um valor acima de 60 microgramas por grama em 24 horas é sugestivo.
    -   **Outros mediadores:** Cromogranina A (pode estar elevada pelo uso de inibidores da bomba de prótons), heparina (potencialmente o melhor marcador, mas ainda não validado), prostaglandinas e leucotrienos podem estar elevados, mas não são validados para diagnóstico.
3.  **Biópsia do Trato Gastrointestinal:** A endoscopia ou colonoscopia com biópsias e análise por imuno-histoquímica pode revelar um aumento no número de mastócitos (>20 por campo de grande aumento), o que apoia o diagnóstico.
4.

---

### Chunk 12/30
**Article:** Glúten (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.514

, alergia ao trigo), utilizando a abordagem clínica como soberana.
- [ ] 3. Adotar e educar os pacientes sobre a importância de um estilo de vida anti-inflamatório (dieta equilibrada, manejo do estresse, sono) como primeira linha de intervenção para a saúde intestinal e geral.
- [ ] 4. Ao prescrever antibióticos, considerar a recomendação de probióticos (preferencialmente de múltiplas cepas ou fontes naturais) para mitigar danos ao microbioma.
- [ ] 5. Manter-se atualizado sobre as pesquisas em medicina personalizada e análise do microbioma para futuras aplicações clínicas.

---

### Chunk 13/30
**Article:** Fisiologia e Bioquímica do Sistema Imune I (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.513

vel apresentou 7,94%, indicando inflamação ativa.
- A presença de 21% da bactéria Prausnitzi, combinada com a ausência de Akkermansia, sugere uma dieta com carga glicêmica muito alta.
- A estratégia de tratamento inclui um protocolo de "sete passos para a reprogramação intestinal", que pode envolver dietas como a low FODMAP por um a dois meses.
- Para a modulação intestinal, são sugeridas tinturas em proporções específicas, como 50% de alcaçuz e 50% de cúrcuma.
**O tratamento é altamente personalizado, utilizando suplementos em doses específicas e protocolos dietéticos faseados para controlar a inflamação e modular a resposta imune.**
- Suplementos de curcumina devem ter alta concentração de curcuminoides (95% a 99%) para garantir eficácia.
- Para o controle de TH2, doses de N-acetilcisteína variam de 400 mg a 1000 mg, enquanto para a modulação de TH17, a berberina é usada em doses de 100 mg a 300 mg.

---

### Chunk 14/30
**Article:** Levels of Evidence Supporting American College of Cardiology/American Heart Association and European Society of Cardiology Guidelines, 2008-2018 (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.512

.com.br - CPF: 034.983.039-88

horse.JClinEndocrinolMetab.2013;98(8):3246-3252.doi:10.1210/jc.2013-181497.FeuersteinJD,AkbariM,GiffordAE,etal.Systematicreview:thequalityofthescientific
evidenceandconflictsofinterestininternational
inflammatoryboweldiseasepracticeguidelines.
AlimentPharmacolTher.2013;37(10):937-946.doi:10.1111/apt.1229098.FeuersteinJD,AkbariM,GiffordAE,etal.Systematicanalysisunderlyingthequalityofthe
scientificevidenceandconflictsofinterestin
interventionalmedicinesubspecialtyguidelines.

---

### Chunk 15/30
**Article:** Fisiologia e Bioquímica do Sistema Imune II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.510

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

### Chunk 16/30
**Article:** Functional Disease, Dysbiosis, and Dyspepsia: How Helpful Is Rifaximin? (2026)
**Journal:** Revista não identificada
**Section:** discussion | **Similarity:** 0.505

non-invasive diagnosis of small intestinal bacterial overgrowth: a sys-tematic review with meta-analysis. J Neurogastroenterol Motil. 2020;26:1628. 6. Black CJ, Burr NE, Camilleri M etal. Eﬃcacy of pharmaco-logical therapies in patients with IBS with diarrhoea or mixed stool pattern: systematic review and network meta-analysis. Gut. 2020;69:7482. 7. Shah A, Gurusamy SR, Hansen T, Callaghan G, Talley NJ, Koloski N, Walker MM, Jones MP, Morrison M, Holtmann GJ. Concomitant Irritable Bowel Syndrome does not inﬂuence the response to antimicrobial therapy in patients with functional dyspepsia. Dig Dis Sci. (Epub ahead of print). https:// doi. org/ 10. 1007/ s10620- 021- 07149-1. 8. Koloski NA, Jones M, Hammer J etal. The validity of a new struc-tured assessment of gastrointestinal symptoms scale (SAGIS) for evaluating symptoms in the clinical setting [published correction appears in Dig Dis Sci 2017 Jul 8]. Dig Dis Sci 2017;62:19131922. https:// doi. org/ 10.

---

### Chunk 17/30
**Article:** Vitamina D (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.504

ntil, estresse, deficiência severa de vitamina D com nível de 19 ng/mL).
*   **Tratamento:** Após pulsoterapia com corticoides, a paciente recusou as medicações alopáticas convencionais e optou por um tratamento integrativo com altas doses de vitamina D (30.000 UI/dia), cofatores (B2, B12, magnésio) e mudanças no estilo de vida.
*   **Resultados:** Em três meses, a ressonância magnética de controle mostrou uma redução "importantíssima" das lesões, sem novas lesões e sem captação de contraste, indicando ausência de atividade inflamatória.
*   **Conclusão do Caso:** O caso ilustra o potencial da abordagem integrativa, que combina o melhor da medicina convencional (ex: corticoides em surtos) com terapias complementares. Enfatiza-se a corresponsabilidade do paciente, que deve aderir a uma dieta com restrição de cálcio, hidratação adequada e atividade física.
## ❓ Perguntas
- [Inserir Pergunta/Confusão]
## 📚 Tarefas
- [ ] 1.

---

### Chunk 18/30
**Article:** Modulação Intestinal II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.501

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

### Chunk 19/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa I (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.500

 ões excessivas e crenças limitantes sobre comer fora, como medo de passar mal).
## Objetivo:
O conteúdo é uma apresentação médica teórica, não uma consulta. Descreve os achados objetivos e a fisiopatologia da SII. A classificação do padrão intestinal é feita usando a Escala de Bristol (tipos 1-2 para constipação, 3-4 normais, 5-7 para diarreia), subdividindo a SII em subtipos: com constipação (IBS-C), com diarreia (IBS-D), mista (IBS-M) ou indeterminada (IBS-U). Fisiologicamente, a SII envolve alterações na motilidade, sensibilidade visceral, função imune, inflamação de baixo grau, microbiota intestinal e processamento no sistema nervoso central (SNC).
- **Achados de Pesquisa:** Estudos de neurorressonância funcional mostram maior ativação em áreas cerebrais (rede sensorial motora, rede autonômica central) em pacientes com SII.

---

### Chunk 20/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa III (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.500

ratório para SIBO/IMO:** Considerado o método prático padrão, utilizando lactulose ou glicose.
    -   **Critério para SIBO (Hidrogênio):** Elevação acima de 20 ppm nos primeiros 90 minutos.
    -   **Critério para IMO (Metano):** Elevação acima de 10 ppm em qualquer momento do teste.
-   **Diagnóstico de SIFO:** O padrão ouro é o aspirado duodenal com cultura (>10³ UFC/ml de fungos), mas é raramente realizado. A avaliação da micobiota e metabólitos fúngicos pode ser útil.
-   **Outros Achados:** Ativação do eixo HPA e reação de "die-off" (com manifestações cutâneas) durante tratamento antifúngico.
## Diagnóstico Primário:
-   **Avaliação:** O texto é uma discussão abrangente sobre a Síndrome do Intestino Irritável (SII), suas causas subjacentes (aumento da permeabilidade intestinal, disbiose, inflamação) e comorbidades como distúrbios do sono, SIBO, IMO e SIFO.

---

### Chunk 21/30
**Article:** Microbioma Intestinal II (2024)
**Journal:** Pos Graduacao MFI
**Section:** results | **Similarity:** 0.497

  intestino permeável” por uma arquitetura clara. Conforme evolui, transforma diagnóstico e tratamento em engenharia de sistemas biológicos: identificar qual camada falhou e qual intervenção a restaura (p. ex., modulação da microbiota; suporte de SCFAs para muco; fortalecimento de claudinas/occludinas para junções; redução de disparadores e reequilíbrio imune). Em sua forma mais avançada, ele integra-se à métrica de homeostase, oferecendo um mapa topográfico da função intestinal que aumenta a precisão terapêutica e possibilita experimentação dirigida camada a camada. O resultado é um protocolo iterativo e mensurável, capaz de converter achados complexos em linhas práticas de ação com feedback claro.
**Trilha de evidências:**
> “Existem quatro camadas para a barreira do sangue... A microbiota... A camada de mucosa fina... O epifilium... e, finalmente, a tinta linfática associada ao sangue... Uma das estruturas mais essenciais...

---

### Chunk 22/30
**Article:** Fisiologia e Bioquímica do Sistema Imune II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.494

ular IL-6/COX-2 e reduzir picos.
- [ ] 5. Programar FMD vegano por 5 dias consecutivos; definir periodicidade (mensal, bimestral, trimestral) conforme estado clínico.
- [ ] 6. Integrar low carb + cetogênica limpa + jejum + atividade física em jejum visando biogênese mitocondrial; monitorar AMPK, PGC-1α, NRF2 quando possível.
- [ ] 7. Criar plano alimentar de baixa carga glicêmica (abacate, amêndoas, brócolis, etc.); incluir exemplos de café, almoço, lanches e jantar com otimizadores (C8/MCT, CoQ10, PQQ, curcumina, BHB, magnésio inositol).
- [ ] 8. Ajustar tubérculos (batata-doce 50–80 g) conforme nível de atividade física em estratégia low carb/cetogênica limpa.
- [ ] 9. Educar sobre PPAR-γ–melatonina–cravings; reforçar jantar cedo e apigenina à noite.
- [ ] 10. Solicitar e registrar parâmetros essenciais (PCR-us, VHS); calcular índices estimáveis (HGI, TAIG); considerar complementares (TNF-α, IL-6, GPx, MDA, antioxidantes totais).
- [ ] 11.

---

### Chunk 23/30
**Article:** Microbioma Intestinal III (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.488

gastroenterologia pediátrica; colonoscopia apenas se evolução/idade/indicação justificarem.
- Plano de Tratamento de Seguimento:
  - Modulação contínua do sistema gastrointestinal como base para manejo de condições crônicas (ansiedade, autoimunidade, metabolismo).
  - Estímulo à via AHR e à produção de IL-22/defensinas por meio de dieta rica em fibras e compostos como beta-glucana, dentro da tolerância individual.
  - Intervenções de estilo de vida: incremento gradual de atividade física adequada à idade.
  - Educação familiar sobre impacto do microbioma na saúde mental e sistêmica; suporte à adesão terapêutica.
  - Reavaliação clínica e laboratorial periódica para ajustar dieta, suplementos e intervenções conforme resposta; acompanhar sintomas cutâneos, respiratórios e gastrointestinais.
  - Evitar uso indiscriminado de antibióticos e corticoides; racionalizar terapias conforme diagnóstico diferencial e evolução.

---

### Chunk 24/30
**Article:** Síndrome do Intestino Irritável sob o olhar da Medicina Funcional Integrativa III (2024)
**Journal:** Pos Graduacao MFI
**Section:** introduction | **Similarity:** 0.488

lam para melhorar sintomas na maioria dos pacientes. A narrativa central é um percurso clínico: estratificar risco inflamatório com calprotectina, investigar supercrescimento microbiano com critérios padronizados, intervir com dieta e terapias complementares, e monitorar desfechos ao longo de semanas a meses.
---
### Evidências-Chave
**Low FODMAP proporciona melhora sintomática robusta em SII, estruturada em eliminação e reintrodução, com seguimento prolongado.**
- Duração da Fase 1 do Protocolo Low FODMAP: duas a seis semanas (duração mínima e máxima da fase 1).
- Duração da Fase de Reintrodução do Low FODMAP: 8 a 12 semanas (duração típica de reintrodução).
- Taxa de Melhora Sintomática no Low FODMAP: 75% (taxa aproximada de melhora).
- Duração de Seguimento em Estudo Low FODMAP: 12 meses (período de acompanhamento).

---

### Chunk 25/30
**Article:** Lifestyle Medicine: A Brief Review of Its Dramatic Impact on Health and Survival (2026)
**Journal:** Revista não identificada
**Section:** other | **Similarity:** 0.488

c inammation 
to the progression of cancer. e risk of 
colorectal cancer developing increases with 
the duration and extent of inammatory 
bowel disease.
87
 e microbiome of the gut 
has also been implicated in the develop
-
ment of sporadic colorectal carcinoma.
25,88
 
Risk factors for the development of 
colorectal cancer include a sedentary 
lifestyle, obesity, and the dietary compo
-
nents that form the basis of the standard 
American diet (large consumption of red 
meats and highly processed foods and low 
amounts of fruit, vegetables, legumes, and 
ber intake).
89
 Low-ber diets, such as the 
standard American or Westernized diet 
that promotes inammation, have been 
linked to the increased risk and develop
-
ment of colorectal cancer.
90
 In addition, 
patients with colorectal cancers appear to 
have more comorbidities at the time of di
-
agnosis than patients with other malignan
-
cies.

---

### Chunk 26/30
**Article:** Microbioma Intestinal III (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.487

ias metabólicas como a AHR e a do triptofano na saúde mental e imunológica. A segunda parte aprofunda a análise de exames fecais, como o Copromax, para diagnosticar a saúde intestinal, detalhando marcadores como alfa-1-antitripsina, calprotectina, lactoferrina, IgA secretória e elastase pancreática. Utilizando o caso de uma criança com inflamação severa, o instrutor ilustra como esses marcadores indicam permeabilidade intestinal (leaky gut), inflamação crônica e desequilíbrios digestivos. A palestra conclui enfatizando uma abordagem clínica personalizada, que inclui a história do paciente, ferramentas como a Escala de Bristol, e intervenções terapêuticas como o suplemento Biointestil e terapias alternativas (hidrocolonoterapia, enemas de café), antecipando a próxima aula sobre fibras, probióticos e paraprobióticos.
## 🔖 Knowledge Points
### 1.

---

### Chunk 27/30
**Article:** Intolerâncias, Alergias e Hipersensibilidades Alimentares II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.482

RAST/ImmunoCAP: pedir preferencialmente com IgE elevada; diagnósticos moleculares específicos por alimento; útil também para alérgenos respiratórios.
- Teste de ativação de basófilos: específico para tipo I; disponível em poucos laboratórios.
- Testes IgG: não solicitar de rotina; positividade pode refletir tolerância; utilidade restrita (p.ex., IgG4 na esofagite eosinofílica).
### 17. Exames complementares gastrointestinais e marcadores fecais
- Endoscopia/colonoscopia com biópsia: avaliação de mucosa e inflamação.
- Alfa-1 antitripsina fecal: marcador de enteropatia perdedora de proteínas.
- Cintilografia/nucleares: avaliação de esvaziamento gástrico.
- Calprotectina fecal: marcador inflamatório gastrointestinal; valor <138 pode excluir alergia não mediada por IgE com alta sensibilidade (~95%) e especificidade moderada (~78,6%).

---

### Chunk 28/30
**Article:** Dislipdemia (2024)
**Journal:** Pos Graduacao MFI
**Section:** discussion | **Similarity:** 0.480

vida e identificar fatores de risco para inflamação e disfunção endotelial (dieta, estresse, sedentarismo).
- [ ] 2. Utilizar a calculadora MESA Risk para determinar o risco cardiovascular dos pacientes e considerar a solicitação do Escore de Cálcio Coronariano antes de prescrever estatinas em prevenção primária.
- [ ] 3. Ao avaliar o risco cardiovascular, solicitar exames avançados como a relação ApoB/ApoA, subpartículas de lipoproteínas e anti-LDL oxidado para uma análise mais aprofundada.
- [ ] 4. Em pacientes de prevenção primária, mesmo com colesterol alto, priorizar mudanças no estilo de vida antes de considerar o uso de estatinas.
- [ ] 5. Reservar o uso de estatinas principalmente para pacientes em prevenção secundária ou de altíssimo risco, focando nos seus efeitos pleotrópicos.
- [ ] 6.

---

### Chunk 29/30
**Article:** International Urogynecology consultation chapter 2 committee 3: the clinical evaluation of pelvic organ prolapse including investigations into associated morbidity/pelvic floor dysfunction (2023)
**Journal:** International Urogynecology Journal
**Section:** results | **Similarity:** 0.479

g peritone-ocele, full-thickness, and internal rectal prolapse, possibly because of more physiological positioning for DCPLienemann etal. [125]Casecontrol study, Germanyn=66Physical examinationDiagnosis of enteroceleExaminationMR-CCRG MR-CCRG was better than DCP at diagnosing enterocelesPresent4353Absent122Diagnosis of enteroceleExaminationDCPPresent2314Absent1120Diagnosis of enteroceleMR-CCRG DCP55 patients with POPDCPPresent291411 controls without POPMR-CCRG Absent520MR-CCRG detected enteroceles missed on examination

2682
 International Urogynecology Journal (2023) 34:2657–2688
Table 9  (continued)
ReferenceStudy designPopulationMethod(s) of clinical assessmentResultsDiscussionAnal physiology testingGroenendijk etal.

---

### Chunk 30/30
**Article:** Fisiologia e Bioquímica do Sistema Imune II (2024)
**Journal:** Pos Graduacao MFI
**Section:** other | **Similarity:** 0.478

c/2, ferro/ferritina/transferrina, TNF-α, IL-6, HOMA-β/IR, homocisteína, PCR. Monitoramento a cada 3–5 meses, paciente como próprio controle.
### 10. Estresse oxidativo, glicação e vias pró-inflamatórias
- ROS elevam NF-κB, AP-1; LPS/PAMPs/DAMPs ativam caspases e IL-1β/IL-18/IL-6.
- Reação de Maillard: açúcar redutor + aminoácidos + gordura → AGEs; hiperglicemia aumenta HbA1c; autoimunes demandam baixa carga glicêmica.
- Polióis (sorbitol, maltitol, xilitol) geram AGEs por via frutose.
- Impactos: resistência à insulina, T2D, DCV, pulmonares e neurológicos.
- Exemplo crítico: churros (gordura + açúcar + leite) maximiza AGEs.
- Antiglicação: EGCG, trans-resveratrol, mio-inositol.
### 11. Marcadores e metas de acompanhamento
- HbA1c: meia-vida ~120 dias; metas integrativas ~5,3–5,2; diagnóstico ≥6,5; risco alto ≥5,6. Evoluções podem levar 2–3 anos.
- Frutosamina: ~20 dias; complementar.

---

