# Plano — Integração dos painéis Bioma (genética) ao Escore Plenya

**Status:** PLANEJAMENTO (2026-06-04). **Nada modificado no banco.** Análise dos 2 painéis reais
de parceria (`docs/emr/Bioma Hair.pdf`, `docs/emr/Bioma - Lifecode.pdf`) cruzada com os 81 genes
do grupo Genética do nosso escore. Dados brutos do cruzamento: `docs/emr/bioma-cruzamento-dados.md`.

## Metodologia (anti-alucinação)

- Texto extraído com `pdftotext -layout`. rs IDs por regex (`rs\d+`); gene = coluna imediatamente
  antes do rs na "TABELA DE GENÓTIPOS RELEVANTES" (Lifecode) / linha-do-gene + `(rs)` (Hair).
- Cobertura de extração: Lifecode 283/285 rs capturados com gene; Hair 28 rs.
- Cada "só nosso" foi **verificado no texto bruto** da Bioma antes de afirmar ausência.
- **Caveats (conferir manualmente, não assumi):**
  - **ACE** aparece como "ECA (ACE)" (Hair) e **HLA-DQ2/DQ8** como tag-SNPs sob "HLA"/"HLA-DRA"
    (Lifecode) — meu parser perdeu pelos nomes; corrigidos manualmente para "em ambos".
  - ⚠️ **CORREÇÃO (Dr. pegou, 2026-06-04):** a 1ª extração por-SNP exigia símbolo de gene PADRÃO
    antes do rs e **perdeu 20 rs** com rótulo PT/composto ("IL1-beta", "Protrombina", "VDR BsmI/
    TaqI/ApaI", "TNF-alfa", "IL-4Ra", "Fator V HR2/Leiden", "FXIII V34L", "ITGB3/HPA1", "HLA-DRA",
    "NFIA-AS2", "ARMS2", "CFH", DHCR7) — **um sub-painel de coagulação inteiro** + SNPs extras de
    VDR/IL1B/TNF. Também perdeu o **readout consolidado APOE ε (E2/E3/E4)** (linha não-rs "E4").
    **TODOS adicionados** (19 variantes + APOE-ε). rs3856806 excluído (só em citação, não é teste).
    Re-verificado: **0 variantes testadas do Lifecode fora do escore.** "307" = nº de catálogo
    declarado; o PDF enumera 285 rs (284 testes reais + 1 citação) + APOE-ε derivado.

## Resultado do cruzamento

| | Genes | Observação |
|---|---|---|
| Nosso escore (Genética) | **81** | 7 subgrupos atuais |
| Bioma (Lifecode 200 + Hair 25, únicos) | **216** | nutrigenética + capilar |
| **Em ambos** | **46** | auditar mutação (Task A) |
| **Só Bioma** | **172** | candidatos a adicionar (Task B) |
| **Só nosso (Bioma não cobre)** | **35** | decisão necessária |

**Achado estrutural:** os painéis Bioma são **nutrigenética (Lifecode Nutri) + tricologia (Hair)**,
não um painel clínico/monogênico. Por isso não cobrem os genes monogênicos do nosso escore (MODY,
Parkinson/Alzheimer familiar, dislipidemias familiares). Isso é por design da Bioma, não falha.

---

## Task A — Auditoria de mutações dos 46 genes em ambos

> **DECIDIDO (Dr., 2026-06-04): ADOTAR BIOMA em tudo** — trocar os 4 SNPs divergentes pelos da
> Bioma, adotar os 6 genes sem-rs com o rs da Bioma, e adicionar os 9 SNPs extras que a Bioma testa.

- **34 com rs idêntico (OK):** ACTN3, AGT, AGTR1, ALDH2, APOA5, CAT, COL1A1, COL5A1, CYP1A2,
  CYP2A6, EPHX1, FABP2, FADS1, FTO, GNB3, GPX1, GSTP1, IGF2BP2, IL1B, IL6, KCNJ11, LDLR, LIPC,
  MC4R, MCM6, MTHFR, PCSK9, PPARG, SLC23A1, SLC30A8, SOD2, TCF7L2, TNF, VDR (+ ACE, HLA por nome).
- **9 desses: a Bioma testa SNP(s) ADICIONAL(is)** que não temos (oportunidade de enriquecer o
  mesmo gene): ex. MTHFR (+rs1801131), FTO (+rs8050136), PCSK9 (+rs11206510,+rs505151), TCF7L2
  (+rs12255372), LDLR (+rs6511720), GSTP1 (+rs1138272), LIPC (+rs2070895), EPHX1 (+rs2234922),
  MC4R (+rs12970134). Lista exata no apêndice.
- **6 onde nosso gene está SEM rs e a Bioma fornece** → adotar o rs da Bioma:
  APOE (rs429358+rs7412), GSTM1 (rs2071487...), GSTT1 (rs2266633...), NAT2 (rs1495741),
  INS (rs3741208), ALPL (rs1697421).
- **4 DIVERGÊNCIAS (nosso SNP ≠ o que a Bioma entrega)** — decisão:
  | Gene | Nosso SNP | SNP da Bioma |
  |---|---|---|
  | APOA1 | rs670 | rs1799837 |
  | FADS2 | rs174575 | rs174616 |
  | IRS1 | rs1801278 | rs2943641 |
  | LPL | rs328 | rs13702 |
  Como a Bioma é a fonte real do dado, a recomendação é **trocar nosso SNP pelo da Bioma** (senão
  o item nunca recebe resultado). Alternativa: manter o nosso e marcar "não disponível via Bioma".

## Task B — Plano para adicionar os 172 genes só-Bioma

**Estrutura proposta (mesmo molde dos genéticos atuais):**

- **Grupo:** `Genética` (existente, `g.order=13`).
- **Subgrupos (espelhando os domínios do Lifecode + Hair):** distribuição aproximada dos 172
  (refino por-gene na execução):

  | Subgrupo proposto | ~genes |
  |---|---|
  | Outros / a classificar (refinar) | 72 |
  | Neuro, Humor & Comportamento | 15 |
  | Pele, Estética & Capilar (Hair) | 15 |
  | Vitaminas & Micronutrientes | 15 |
  | Lipídico & Cardiovascular | 14 |
  | Metabolismo, Apetite & Energia | 13 |
  | Circadiano & Sono | 7 |
  | Cardiovascular & Coagulação | 6 |
  | Imune & Inflamatório | 6 |
  | Detoxificação | 5 |
  | Aptidão Física & Lesão | 4 |

  (Listas por subgrupo no apêndice. "Tricologia/Capilar" é domínio NOVO, vindo do Hair.)

- **Levels:** espelhar o modelo atual — **1 `ScoreLevel` por gene** (`level=0`, `operator='='`),
  `name = "Genótipo | <alelo-risco>=risco / <alelo-proteção>=proteção"`, derivado do **alelo de
  risco que a Bioma informa em cada linha** (está no PDF; não inventar). Opcional futuro: modelar
  3 genótipos (normal/heterozigoto/homozigoto-risco) se quisermos pontuação graduada — hoje o
  escore não faz isso para genes.

- **Pillars (M2M):** cada gene novo → **Genético** (automático, regra D3) **+ 1-2 áreas clínicas**
  pelo domínio (Lipídico→Cardiovascular/Metabólico; Vitaminas→Nutrologia; Detox→Toxicológico;
  Neuro/Humor→Neurológico/Humor/Função Cognitiva; Aptidão→Capacidade Física(A)/Osteomuscular;
  Capilar→Dermatológico; Circadiano→pilares de R + Neurológico). Mesma lógica/cap-3 já usada.

- **Códigos:** cada item novo recebe `anamnese_item_code` estável (ex.: `GENE_rsNNNN`).

## 35 genes só-nossos (Bioma não cobre)

> **DECIDIDO (Dr., 2026-06-04): MANTER como estão** no escore (não dropar), aguardando o Dr.
> encontrar um fornecedor/laboratório complementar que os cubra. Ficam sem fonte de dado por ora.

ABCA1, ABCC8, ADD1, ADH1B, APOC2, APP, BCO1, C9orf72, CDKAL1, CRP, CYP11B2, CYP1A1, ESR1, GCK,
GPIHBP1, GRN, HHEX, HNF1A, HNF1B, HNF4A, LCAT, LEPR, LMF1, LRRK2, MAPT, NOS3, PARK2, PARK7/DJ-1,
PINK1, POMC, PPARA, PPARGC1A, PSEN1, PSEN2, SNCA.

Maioria é monogênica/clínica (MODY, Parkinson/Alzheimer familiar, dislipidemia familiar) — fora do
escopo nutrigenético da Bioma. Opções:
- (a) Manter como **entrada manual** (laudo de outro laboratório quando houver indicação clínica);
- (b) Marcar como **"não disponível via Bioma"** e despriorizar no fluxo padrão;
- (c) Buscar **painel clínico complementar** (exoma/targeted) para esses genes.

## Execução decidida (Dr., 2026-06-04) — dev-only, sem prod

- **Task A — ADOTAR BIOMA:** trocar os 4 SNPs divergentes pelos da Bioma; adotar os 6 sem-rs;
  adicionar os 9 SNPs extras dos genes em comum.
- **Task B — ADICIONAR os genes da Bioma** com os grupos/subgrupos propostos. **Unidade = 1 item
  por variante (gene+rs)**, espelhando o padrão atual "GENE rsXXX (traço)". Nome do item:
  `{GENE} {rs} ({função})`; `anamnese_item_code = {GENE}_{rs}`; 1 `ScoreLevel` (level=0,
  `Genótipo | {RISCO}=risco`) a partir da coluna RISCO da Bioma. Sizing: 234 variantes Lifecode
  (172 genes) + variantes do Hair.
- **35 só-nossos:** MANTER no escore, aguardando o Dr. achar fornecedor complementar.
- **Sequência:** (1) extrair variantes; (2) classificar subgrupo + pilares; (3) gravar itens+levels
  no dev; (4) aplicar Task A; (5) **depois, distribuir TODOS os genes (81 + novos) no AGIR de uma vez.**

## Execução — andamento

- ✅ **Task B FEITO (dev):** 10 subgrupos + **247 itens + 247 níveis** (alelo de risco da Bioma).
  SQL `docs/emr/agir-bioma-itens.sql`. Genética: 81 → **328 itens**. Pilares de cada variante já
  computados (Genético + área), prontos pra etapa AGIR. UUIDs `b10…` fixos, idempotente.
- ✅ **Task A FEITO (dev):** 4 trocas (APOA1→rs1799837, FADS2→rs174616, IRS1→rs2943641, LPL→rs13702)
  + 6 adoções (APOE, GSTM1, GSTT1, NAT2, INS, ALPL) + 13 novos itens (10 SNPs extras + 3 segundo-rs).
  SQL `docs/emr/agir-bioma-taskA.sql`. Genética: 328 → **341 itens**.
- ✅ **AGIR FEITO (dev):** todos os **341 genes** → Genético + áreas (932 vínculos, 0 não-resolvidos).
  SQL `docs/emr/agir-dist-genetica.sql`. Genético pegou os 341; áreas coerentes (Metabólico 94,
  Cardiovascular 79, Neurológico 50…).
