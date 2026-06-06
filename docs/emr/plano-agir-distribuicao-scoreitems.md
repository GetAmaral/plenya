# Plano — Distribuição dos ScoreItems nas letras/pilares AGIR

**Iniciado:** 2026-06-04 · **Status:** Fase 2a APLICADA NO DEV — 42 pilares construídos (SQL direto
`docs/emr/agir-taxonomia-pilares.sql`, UUIDs fixos, letras renomeadas, M2M zerado). NÃO commitado/
deployado. Próximo: Fase 2b (distribuição dos 916 itens). dev-only, sem prod.

## Objetivo

Redesenhar a taxonomia de pilares do método AGIR e popular o M2M
`score_item_method_pillars`, ligando os **916 score_items ativos** ao(s) pilar(es) que
fizer(em) sentido — primeiro tudo no dev, depois migrar pro prod. Hoje o dev tem **30
vínculos exploratórios**; o prod tem **0**.

## Fase 1 — Verificação dev ≡ prod ✅ CONCLUÍDA (2026-06-04)

Hashing de identidade (id-set + conteúdo, ignorando timestamps voláteis) nas duas bases:
todo o esqueleto é **byte-idêntico com os mesmos UUIDs**: methods(1), method_letters(4),
method_pillars(13), score_groups(13), score_subgroups(40), score_items(961/916 ativos),
score_levels(3521). Única divergência: o próprio alvo `score_item_method_pillars`
(dev=30, prod=0). SQL: `/tmp/score_compare.sql`.

## Decisões transversais (2026-06-04, definidas pelo Dr.)

- **D1 — M2M multi-pilar:** um score_item PODE pertencer a >1 pilar quando relevante. Sem
  regra de "1 pilar primário".
- **D2 — Órfãos permitidos:** item pode ficar sem pilar. **Objetivos** (grupo 1) NÃO precisa
  de pilar. Mas **Medicamentos, Hábitos e vícios nocivos, Saúde bucal, Cirurgias** PRECISAM
  de pilar (definir depois). **Acompanhamento médico/multiprofissional** — decidir depois.
- **D3 — Genética distribuída:** cada gene entra no(s) pilar(es) específico(s) que fizer(em)
  sentido (Cardiovascular, Metabólico, Neurológico, Imune…), NÃO num bucket único "Genético".
- **D4 — Redesenho completo:** os 13 pilares atuais foram só pra teste. Reescrever TODA a
  taxonomia AGIR, letra a letra. `agir-structure.ts` é input/referência, não palavra final.
  **Ordem: A → G → I → R.** Migration de reconciliação só depois da taxonomia fechada em dev.
  **Nada vai pro prod nesta fase.**

## Fase 2a — Redesenho da taxonomia (letra a letra) — EM DISCUSSÃO

### Letra A — "Atividade Física, Alimentação e Suplementação Inteligente" — ✅ FECHADA (Dr., 2026-06-04)

**Decidido:** 7 pilares (tabela abaixo). #7 Trajetória e Origens PONTUA (não é só contexto).
Defaults aceitos: burpee fica em Capacidade Física (sem vértice cardiorrespiratório próprio);
água/ângulo de fase dentro de Massa Muscular; lesões/cirurgias = modificadores em Capacidade
Física; contexto social/familiar = órfão em A (reavaliar cruzamento p/ letra I quando chegar no I).
Nomes em framing de domínio medido; nome público da letra A no site não muda.

Dados de score que alimentam A (~180 itens):
- **Alimentação** (grupo 2, 58): histórico (23: alimentação por fase da vida, aleitamento,
  intolerâncias lactose/glúten/proteína-leite/histamina, alergias, restrições) + atual (35:
  água, frutas, açúcar, álcool, proteína, calorias, padrões dietéticos, recordatório,
  **suplementação utilizada/prescrita**).
- **Movimento e atividade física** (grupo 3, 63): histórico (23) + atual (8, inclui
  **suplementação pré/intra/pós-treino**) + **testes práticos (32)**: resistência neuromuscular
  (abdominal, flexão, prancha) e cardiorrespiratória (burpee) por sexo/idade.
- **Composição corporal** (grupo 8, 59): histórico (9) + atual (3: DEXA/bioimpedância) +
  **medidas objetivas (47)**: peso, altura, IMC, BRI, cintura, quadril, razões, % gordura,
  gordura visceral, massa muscular, ASMI, água corporal, ângulo de fase.

**Proposta (estudo multi-lente 2026-06-04 — 4 lentes + síntese):** 7 pilares para A.

| # | Pilar | ~itens | mede | cruza p/ G |
|---|---|---|---|---|
| 1 | Alimentação (Qualidade e Padrão) | 26 | o que/como come hoje: dieta, açúcar, álcool, hidratação comportamental, recordatório, adesão | Metabólico, Hepático (álcool), GI, Nutrologia |
| 2 | Tolerância Alimentar e Suplementação | 9 | intolerâncias/alergias/restrições + suplementos (dose/adesão) | GI (lactose/glúten), Imune (histamina/alergia), Nutrologia/Hormonal |
| 3 | Atividade Física (Hábito) | 13 | comportamento de mover-se (volume/frequência/intensidade, sedentarismo) — autorrelato | Cardiovascular, Metabólico, Osteomuscular |
| 4 | Capacidade Física (Testes) | 36 | aptidão MEDIDA com tabela normativa (abdominal/flexão/prancha + burpee) | Cardiovascular/Pulmonar (burpee), Osteomuscular |
| 5 | Adiposidade e Risco Cardiometabólico | 22 | gordura e distribuição central (cintura, visceral, IMC, % gordura) — ponte A→G prioritária | Metabólico, Cardiovascular, Hormonal |
| 6 | Massa Muscular e Hidratação Celular | 23 | músculo/sarcopenia (ASMI/MME) + água corporal/ângulo de fase | Osteomuscular, Hormonal, Metabólico, Renal |
| 7 | Trajetória e Origens (Histórico/DOHaD) | 48 | cluster temporal: perinatal/fases da vida/histórico familiar — risco programado | Genético, Metabólico |

Órfãos por D2: metas/objetivos, satisfação, preferências (modalidades, alimentos que gosta),
contexto social/familiar (quem prepara/quem treina → ou cruza p/ letra I).

**Convergência forte das 4 lentes (baixo risco de travar):**
1. **Relato vs Medida** — separar hábito autorrelatado da capacidade medida (no movimento E na composição).
2. **Adiposidade ≠ Massa Magra** — não fundir num "Composição Corporal" único (apagaria sinais opostos: gordura↑ vs músculo↓).
3. **Framing de domínio medido** — "Alimentação"/"Atividade Física" (não "Avaliação Nutricional"/"Prescrição de Exercícios"); casa 1:1 com os score groups do EMR; **nome da letra A no site NÃO muda**.
4. **Suplementação é transversal** e fina demais sozinha → fundida com Tolerância.

**Decisões abertas pro Dr. (divergências):**
- (D-A1) Granularidade: **7** (recomendado) vs 6 vs 8-11.
- (D-A2) Cardiorrespiratório (burpee) vira vértice próprio? Recomendado **não** (só 1 movimento, sem VO2) — fica em Capacidade Física.
- (D-A3) Água/ângulo de fase: dentro de Massa Muscular (recomendado) ou 3º vértice "Integridade Celular"?
- (D-A4) **O maior:** os ~48 itens de Histórico/DOHaD viram o pilar #7 "Trajetória e Origens" (recomendado) OU ficam órfãos/cruzam direto pra G (radar só com estado atual modificável)?
- (D-A5) Lesões/cirurgias/restrições (4) = modificadores em Capacidade Física (recomendado) ou pilar próprio?
- (D-A6) Contexto social/familiar: órfão em A ou cruza pra letra I (Vínculos)?

### Letra G — "Gestão Clínica e Metabólica" — ✅ FECHADA (Dr., 2026-06-04) — 23 pilares

**Achado (estudo dos ~580 itens reais, 2026-06-04):** o canônico de 14 sistemas é
**lab-centric**. Os labs (245) e a genética (81, via D3) caem limpo nos 14. Mas o **histórico
de doenças (163)**, **imagem (67)** e **cirurgias (22)** trazem domínios clínicos SEM casa nos
14. Buracos identificados:
- **Infeccioso/Infectologia** — sorologias (Hep B/C, HIV, VDRL) + infecções (dengue, COVID, TB, herpes, EBV/CMV…).
- **Toxicológico e Detoxificação** — metais pesados (As, Pb, Hg, Cd, Al) + genética CYP/GST/NAT2 + amálgamas.
- **Dermatológico/Tegumentar** — pele, cabelo, unha, pelos.
- **Urológico/Andrológico** — genital masc., próstata, PSA, USG próstata.
- **Ginecológico/Saúde da Mulher** — genital fem., mama, USG transvaginal, mamografia.
- **Órgãos dos Sentidos** — oftalmo (fundoscopia/retinopatia) + ORL (rinite, sinusite, zumbido, visão).
- **Odontológico/Saúde Bucal** — já sinalizado (D2): amálgamas, periodontal, mastigação, Rx panorâmico.

**TAXONOMIA FINAL DE G — 23 pilares (Dr., 2026-06-04):**

*① Sistemas (14):* Cardiovascular · Metabólico · Renal · Hepático · Gastrointestinal · Pulmonar ·
Hematológico (+coagulação) · Hormonal · Imune e Inflamatório · Neurológico · Osteomuscular ·
Nutrologia e Micronutrientes · Rastreio Oncológico · **Genético**.

*② Novos domínios clínicos (6):* Infeccioso · Toxicológico/Detox · Dermatológico ·
**Saúde Urogenital e Reprodutiva** (uro+gineco unidos) · Órgãos dos Sentidos (oftalmo+ORL) · Odontológico.

*③ Transversais (3):* Medicamentos · História Cirúrgica · Hábitos e Vícios.

**Regras de vínculo:**
- **Genético (atualiza D3):** TODOS os 81 genes entram no pilar Genético (visão de painel completo)
  E cada gene relevante TAMBÉM cruza pro(s) pilar(es) da área (Cardiovascular, Metabólico,
  Neurológico, Imune, Toxicológico…). Multi-pilar nativo (D1).
- **Transversais:** Medicamentos, Cirurgias e Hábitos/Vícios têm **pilar próprio E cruzamento**
  pro sistema (anti-HAS→Cardiovascular, nefrectomia→Renal, tabaco→Pulmonar/Cardiovascular…).
- **Acompanhamento médico (2 itens):** órfão/meta (não é sistema).

### Letra I — "Integração Mente-Corpo" — ✅ FECHADA (Dr., 2026-06-04) — 7 pilares

**Achados (itens reais: Cognição 38, Stress 6, Vida Sexual 21, Social 23):**
- O grupo **"Cognição" carrega 3 construtos distintos:** (a) **Saúde Mental/psiquiátrico**
  (depressão, ansiedade, bipolar, TDAH, TEA, esquizofrenia, PHQ-9, GAD-7, psicotrópicos);
  (b) **Função Cognitiva/neurocognição** (memória, foco, 5-palavras Dubois, span de dígitos);
  (c) **Vitalidade/Disposição** (energia/disposição p/ atividades). Fundir tudo apaga sinais.
- O grupo **"Social" carrega exposoma ambiental** que não é "mente-corpo": Qualidade do Ar
  Interior, Exposição Ambiental, Ambiente Sonoro, Qualidade da Água, Luminosidade Natural,
  Espaço p/ atividade. Esses vazam: luz→R, ar/água→G(Toxicológico/Pulmonar), espaço→A.
- **Vida Sexual** aqui é o lado experiencial/funcional (libido, ciclo, escalas ASEX/IIEF-5/FSFI)
  — cruza com G·Saúde Urogenital (lado órgão/diagnóstico).
- "Técnicas de Relaxamento" (canônico) some como pilar → vira o manejo dentro de Estresse.

**TAXONOMIA FINAL DE I — 7 pilares (Dr., 2026-06-04):** Humor · Função Cognitiva ·
Vitalidade e Disposição · Estresse, Trauma e Resiliência · Vida Sexual · Vínculos Sociais e
Suporte (inclui Propósito/Espiritualidade) · Contexto e Determinantes de Vida.
- "Cognição" quebrada em 3 (Humor + Função Cognitiva + Vitalidade).
- "Saúde Mental" renomeado **Humor**. Propósito/Espiritualidade **fundido** em Vínculos.
- Exposoma ambiental NÃO é pilar de I → **distribuído**: luz→R, ruído→R, ar/água→G
  (Toxicológico/Pulmonar), espaço p/ atividade→A.
- Cruzamentos: Função Cognitiva→G·Neurológico; Vida Sexual→G·Saúde Urogenital.
- TDAH/TEA/esquizofrenia (não são humor) → Função Cognitiva + cruzam G·Neurológico (recomendado).

### Letra R — "Ritmo Circadiano e Repouso" — ✅ FECHADA (Dr., 2026-06-04) — 5 pilares

**TAXONOMIA FINAL DE R:** Qualidade e Arquitetura do Sono · Cronobiologia e Ritmo ·
Exposição à Luz · Higiene e Ambiente do Sono · Distúrbios do Sono.
- Canônico 3 + Higiene/Ambiente (1 pilar, merge de higiene+quarto físico) + Distúrbios (próprio).
- Recebe do exposoma de I: luz→Exposição à Luz; ruído→Higiene e Ambiente.
- **Distúrbios do Sono** cruza forte p/ G: apneia/roncos/CPAP→Pulmonar+Cardiovascular;
  bruxismo→Odontológico; sudorese noturna→Hormonal/Infeccioso. Estimulantes noturnos→cruza A.

## Taxonomia AGIR completa — 42 pilares (2026-06-04)

- **A — Atividade Física, Alimentação e Suplementação Inteligente (7):** Alimentação (Qualidade e
  Padrão) · Tolerância Alimentar e Suplementação · Atividade Física (Hábito) · Capacidade Física
  (Testes) · Adiposidade e Risco Cardiometabólico · Massa Muscular e Hidratação Celular ·
  Trajetória e Origens.
- **G — Gestão Clínica e Metabólica (23):** *Sistemas (14):* Cardiovascular · Metabólico · Renal ·
  Hepático · Gastrointestinal · Pulmonar · Hematológico · Hormonal · Imune e Inflamatório ·
  Neurológico · Osteomuscular · Nutrologia e Micronutrientes · Rastreio Oncológico · Genético.
  *Domínios (6):* Infeccioso · Toxicológico/Detox · Dermatológico · Saúde Urogenital e Reprodutiva ·
  Órgãos dos Sentidos · Odontológico. *Transversais (3):* Medicamentos · História Cirúrgica ·
  Hábitos e Vícios.
- **I — Integração Mente-Corpo (7):** Humor · Função Cognitiva · Vitalidade e Disposição ·
  Estresse, Trauma e Resiliência · Vida Sexual · Vínculos Sociais e Suporte · Contexto e
  Determinantes de Vida.
- **R — Ritmo Circadiano e Repouso (5):** Qualidade e Arquitetura do Sono · Cronobiologia e Ritmo ·
  Exposição à Luz · Higiene e Ambiente do Sono · Distúrbios do Sono.

Também atualizar os **nomes das 4 letras** (method_letters) para os canônicos.

## Fase 2b — Distribuição dos itens (dev) — ✅ CONCLUÍDA (2026-06-04)

**Concluída:** 1155 itens distribuídos · **2585 vínculos M2M**. Por letra: A 290 · G 1007 · I 199 ·
R 71 (itens cruzam letras). Todos os grupos: labs, imagem, doenças, genética (341), alimentação,
movimento, composição, sono, cognição, stress, vida sexual, social, hist. familiar.
SQLs `docs/emr/agir-dist-*.sql`. **Tudo dev-only, nada commitado.** Falta: **Fase 3 (prod)**.

**Órfãos intencionais (sem pilar, por decisão):** 19 Objetivos + **17 Acompanhamento externo**
(painel de equipe de apoio externo — registro sem pontuar; substituiu os 2 itens genéricos).
**Acompanhamento externo (Dr., 2026-06-04):** subgrupos "Acompanhamento — Especialistas médicos
externos" (11: incl. Ginecologista ♀≥18 e Urologista ♂≥40 como rotina) + "Equipe multiprofissional
externa" (6). 3 níveis de status (não acompanha / em dia / atrasado) + nome do profissional como
texto. SQL `docs/emr/agir-acompanhamento-externo.sql`.


**Critério calibrado pelo Dr. (vale pra todos os grupos):**
- Cada item: 1 pilar **primário** + cruzamentos. **Cap = 3 pilares por item** (corta o cruzamento mais fraco).
- Cruzamento liberal/abrangente, **inclusive cross-letter** (G→A/I/R) quando clinicamente acionável.
- Diretrizes fixas: urinálise/sedimento = só Renal · eletrólitos cruzam · hormônios sexuais→Urogenital+Vida Sexual ·
  cortisol/ACTH/DHEA-S→Estresse+Vitalidade (I) · magnésio→Sono (R) · CPK→Osteomuscular · leucócitos/globulinas primário Imune.

**Progresso (M2M dev):**
- ✅ **Exames/Laboratoriais (245)** → 587 vínculos. SQL `docs/emr/agir-dist-labs.sql`.
- ✅ **Exames/Imagem (67)** → 109 vínculos. SQL `docs/emr/agir-dist-imagem.sql`.
- ✅ **Histórico de doenças (161)** → 318 vínculos (transversais roteados 100%: Cirurgias→História
  Cirúrgica, Medicamentos→Medicamentos, Hábitos→Hábitos e Vícios, Saúde bucal→Odontológico).
  SQL `docs/emr/agir-dist-doencas.sql`. **2 itens "Acompanhamento médico" = órfãos (decidir).**
- ✅ **Genética (341)** → 932 vínculos (Genético + áreas). Inclui os 247 da Bioma + Task A. SQL
  `docs/emr/agir-dist-genetica.sql`. Ver `plano-bioma-genetica.md`. (Grupo Genética cresceu 81→341.)
- ⏳ Pendentes: Histórico Familiar (24) · grupos lifestyle A (Alimentação/Movimento/Composição) ·
  R (Sono) · I (Cognição/Stress/Vida Sexual/Social) · Objetivos (órfão).
- **M2M dev atual: 1946** (814 itens). Escore ativo cresceu ~916 → ~1176 itens (Bioma+TaskA).

## Fase 3 — Migrar pro prod — PENDENTE
Taxonomia + distribuição via migration(s) goose com UUIDs hardcoded (dev≡prod). Decidir o
formato (migration de dados vs sync banco-direto) ao fechar a 2b.

## Convenções
Commit direto no master. Migrations goose à mão, aditivas/reversíveis. Banco direto em dev
(nunca API HTTP pra manipular dado). `migrate status` antes de `up`.
