# Plano — Escalas clínicas pergunta-a-pergunta no EMR (PHQ-9, GAD-7, Epworth, …)

**Status:** proposta (não aprovada) · **Autor:** Claude · **Data:** 2026-06-22

## Problema
No site (escore-light) o PHQ-9 tem uma interface ótima: o usuário responde as 9 perguntas
(0–3), o widget soma e classifica automaticamente. No EMR, quando um template de anamnese
inclui o PHQ-9, o profissional só vê os **botões de banda do total** (N0–N5: "20 ou mais",
"15 a 19", …) e precisa calcular a soma de cabeça. Queremos a mesma experiência
pergunta-a-pergunta no EMR, generalizada para GAD-7, Epworth e qualquer escala futura.

## Como funciona hoje (levantamento)

### Site (referência)
- `apps/site/components/escore/phq9-widget.tsx`: as 9 perguntas e os 4 rótulos de resposta
  são **hardcoded** no componente. Soma 0–27, classifica com um `classify()` próprio
  (hardcoded), e grava **só o `selectedLevel`** de um único score item
  (`c77cedd3-2800-721b-94d8-b5515b895753`). As respostas individuais **não** são persistidas.
- Plugado em `escore-light-form.tsx` via `PHQ9_ITEM_IDS` (set de ids que usam o widget).

### EMR
- `apps/web/components/anamnesis/AnamnesisTemplateItemsRenderer.tsx` é o **único** componente
  que renderiza item de escore no preenchimento (usado em compact/mobile e fullscreen/desktop).
- Decide o tipo de input por propriedade do ScoreItem: `unit` → input numérico (com
  `detectLevel()` automático); `levels[]` → botões de banda; sempre um textarea de observação.
- **Já existe** `detectLevel(value, levels)` (L36) que classifica um número contra os
  `levels[]` do item (operador + limites). É exatamente a classificação que o widget precisa.
- Estado/payload por item: `{ scoreItemId, numericValue?, selectedLevel?, textValue?, order }`
  (`AnamnesisItemFormValue` / `AnamnesisItemRequest`). **Um valor por item.**
- Não há special-casing por id no EMR (diferente do site).

### Backend / dados
- `ScoreItem` tem `AnamneseItemCode` (código estável, ex.: `ESCALA_PHQ_9_27`, `GAD_7_21`,
  `ESCALA_DE_SONOLENCIA_DE_EPWORTH_24`), `SiteRenderType` (enum inclui `scale_0_3`, **mas
  nenhum item usa hoje**), `SiteQuestion`, e `ParentItemID`/`ChildItems` (hierarquia, usada
  para categorias de sintoma — não para perguntas de escala).
- `AnamnesisItem` (tabela `anamnesis_items`) guarda **um** valor: `text_value`,
  `numeric_value`, `selected_level`. Não há onde guardar as N respostas individuais.
  (Idem `anonymous_score_items` e `consultation_prep_responses`.)
### Inventário COMPLETO das escalas no score (12 instrumentos, subgrupo "Atual")
Critério: nome com blank de preenchimento (`___`) **e** com níveis de classificação.
Categorizado pela mecânica real (vista nas faixas de nível):

| code | instrumento | níveis | mecânica |
|---|---|---|---|
| `ESCALA_PHQ_9_27` | PHQ-9 (humor) /27 | 6 | **soma** (9 itens 0–3) → classifica pelo total |
| `GAD_7_21` | GAD-7 (ansiedade) /21 | 5 | **soma** (7 itens 0–3) → total |
| `ESCALA_DE_SONOLENCIA_DE_EPWORTH_24` | Epworth (sonolência) /24 | 4 | **soma** (8 itens 0–3) → total ⚠️ levels 1/2 sem limites |
| `IIEF_5_25` | IIEF-5 (função erétil) /25 | 6 | **soma** (5 itens 1–5) → total |
| `5_PALAVRAS_DE_DUBOIS_IMEDIATO_5` | Dubois 5 palavras imediato /5 | 3 | **contagem** (0–5) → já é número |
| `5_PALAVRAS_DE_DUBOIS_TARDIO_5` | Dubois 5 palavras tardio /5 | 4 | **contagem** (0–5) |
| `SPAN_DE_DIGITOS_DIRETO_8` | Span de dígitos direto /8 | 4 | **contagem** (maior sequência) |
| `SPAN_DE_DIGITOS_INVERSO_7` | Span de dígitos inverso /7 | 4 | **contagem** |
| `FSS_…` (sem code) `c77cedd3-…-7a2e-…` | Escala de severidade de fadiga (FSS) | 3 | **custom: média** de 9 itens (1–7), limiar |
| `ASEX_25` | ASEX (função sexual) /25 | 4 | **custom: classificação por regra de item** (ex. "1 item ≥5"), não pelo total |
| `FSFI_36` | FSFI (função sexual feminina) /36 | 4 | **custom: domínios ponderados** (19 itens, 6 domínios), total decimal |
| `PSQI_…` (sem code) `c77cedd3-…-722d-…` | Escala de Pittsburgh (PSQI) /21 | 4 | **custom: algoritmo de 7 componentes** (~19 questões) |

Três classes mecânicas:
- **Soma simples** (4): PHQ-9, GAD-7, Epworth, IIEF-5 → `total = Σ respostas`, classifica via
  `detectLevel(total, levels)`.
- **Administração guiada** (4 — testes de memória/cognição): Dubois ×2, Span ×2 → **não** é
  digitar um número. O widget conduz a aplicação ao vivo (apresenta os estímulos, captura o
  desempenho) e daí computa a nota → classifica via `detectLevel`. Ver "Administração guiada"
  abaixo.
- **Custom** (4): FSS (média), ASEX (regra por item), FSFI (domínios ponderados), PSQI
  (componentes) → exigem `score()`/`classify()` próprios. **Forçar soma 0–3 calcularia errado.**

⚠️ **Não há campo único que sirva para todas.** O framework é o mesmo; o que muda por escala é
a função de score e (em ASEX) a classificação.
- Campos `anamneseItemCode`/`siteRenderType`/`siteQuestion` **já estão expostos** no tipo
  TS gerado do front (`packages/types`), então o renderer consegue detectar uma escala.

### Bug de dados encontrado (precisa corrigir)
- PHQ-9 e GAD-7: os `levels` particionam o range inteiro corretamente.
- **Epworth: levels 1 e 2 têm `operator = '='` com `lower_limit`/`upper_limit` VAZIOS**
  (nomes dizem "De 15 a 17" e "De 11 a 14", mas sem limites). Resultado: totais 11–17 não
  casam com nenhum nível → classificação automática falharia. Tem que virar `between` 15-17 e
  11-14.

## Decisão central de design

A forma limpa e genérica (não hardcodar por id como o site):

1. **Registro de escalas** keyed por `anamneseItemCode` (e por id nos 2 sem código, até
   atribuirmos código). Cada `ScaleDef` declara a estrutura E a mecânica — é isso que dá
   uniformidade sem forçar uma fórmula errada:
   ```ts
   type ScaleDef = {
     code: string;                 // anamneseItemCode (chave)
     kind: 'sum' | 'administered' | 'custom';
     title: string;
     instructions?: string;
     questions?: { text: string; options: { value: number; label: string }[] }[]; // sum
     administration?: AdministrationSpec; // administered (memória/cognição) — ver abaixo
     // score: default Σ (sum). administered = computado do desempenho. custom = fórmula
     //   própria (FSS média, FSFI domínios, PSQI componentes).
     score?: (answers) => { total: number; meta?: unknown };
     // classify: default detectLevel(total, item.levels). ASEX sobrescreve (regra por item).
     classify?: (scored, levels) => number | undefined;
   };
   ```
   A classificação **default** reaproveita `detectLevel(total, levels)` (já existe). Adicionar
   escala nova = uma entrada no registro + `levels` corretos no banco. Zero mudança no widget.

### Administração guiada (testes de memória/cognição) — requisito do Getúlio
Os testes cognitivos têm que ser **aplicáveis dentro do EMR durante o preenchimento**: o
widget apresenta o estímulo e captura o desempenho; a nota é derivada (não digitada).

- **Dubois 5 palavras (imediato e tardio)** — `administration.type = 'word_recall'`:
  1. Codificação: o widget mostra as **5 palavras** (com a categoria semântica de cada uma)
     para o profissional apresentar e checar o registro.
  2. Evocação imediata (item *imediato*): checklist das 5 palavras → marcar evocação espontânea
     vs. com dica. Conta automática 0–5.
  3. Evocação tardia (item *tardio*): mesma checklist reaberta depois → 0–5.
  - Persiste quais palavras foram lembradas (espontânea/dica), não só o total.
- **Span de dígitos (direto e inverso)** — `administration.type = 'digit_span'`:
  - O widget apresenta sequências de dígitos de tamanho crescente (2 tentativas por tamanho);
    o profissional lê em voz alta (botão revela/avança) e marca cada tentativa acerto/erro.
  - A nota = **maior sequência** repetida corretamente (direto = ordem; inverso = ordem reversa).
  - Persiste o detalhe das tentativas + o span final.

`AdministrationSpec` é parte do `ScaleDef` (em `@plenya/domain`), então o **conteúdo** (lista de
palavras de Dubois validada em PT-BR; banco de sequências de dígitos) é dado, não código —
revisado pelo Dr. Getúlio. O mesmo `<ScaleWidget>` renderiza esse modo (sub-componentes
`WordRecallAdmin` / `DigitSpanAdmin`); a saída cai no mesmo caminho (`selectedLevel` +
`scale_responses` JSONB).

2. **Widget genérico `<ScaleWidget>`** no EMR: renderiza `questions[].options`, chama
   `score()` → `classify()`, mostra "X/N respondidas → total → classificação", e grava
   `selectedLevel` (caminho existente) + o JSON das respostas. Um único componente cobre as 3
   mecânicas (sum/count/custom); `count` renderiza um campo único com instrução. Funciona em
   compact e fullscreen.

3. **Fonte única das definições**: colocar o registro em `@plenya/domain` (lógica de domínio
   pura) e fazer **EMR e site** consumirem o mesmo — elimina o PHQ-9 duplicado no site e já
   habilita GAD-7/Epworth como widget no site se quisermos.

### Decisões travadas (2026-06-22)

**A) Persistir as respostas individuais? → SIM.** 1 coluna `JSONB` nullable em
`anamnesis_items`. Motor de score inalterado (lê `selected_level`); JSONB é detalhe para
exibição/histórico, incluindo o flag de **PHQ-9 Q9 (ideação suicida)**. Fase 3 incluída.

**B) Onde definir as escalas? → `@plenya/domain` (fonte única).** Site migra para consumir o
mesmo registro na mesma PR. Fase 4 (site) deixa de ser opcional para a parte de dedup.

## Fases

### Fase 0 — Higiene de dados (rápida, pré-requisito)
- Corrigir Epworth levels 1 e 2 → `operator='between'` + limites 15-17 e 11-14 (dev via psql +
  goose seed/migration para prod).
- **Atribuir `anamnese_item_code`** aos 2 sem código: FSS (`c77cedd3-…-7a2e-…`) e
  PSQI (`c77cedd3-…-722d-…`).
- Validar, para os **12**, que os `levels` particionam o range sem buraco/sobreposição (com a
  função de score de cada um). Marcar explicitamente as que classificam por regra (ASEX) e por
  total decimal (FSFI).
- Confirmar com o Dr. Getúlio os textos PT-BR validados e a fórmula de score de cada escala
  (PHQ-9 já existe no widget do site). Conteúdo clínico é pré-requisito de cada tier abaixo.

### Fase 1 — Framework de definições compartilhadas (`@plenya/domain`)
- `SCALE_REGISTRY: Record<code, ScaleDef>` com `kind` + `questions[]` + `score?`/`classify?`
  (ver tipo acima). Helper `classifyByLevels(total, levels)` reaproveitando a semântica do
  `detectLevel` do EMR, para o site não depender de `classify()` hardcoded.
- Entregar o framework + as escalas do **tier A** (soma simples), que cobrem o pedido original.

### Fase 2 — Widget no EMR + rollout por tier
- `apps/web/components/anamnesis/ScaleWidget.tsx` (genérico: 3 mecânicas). Plugar em
  `AnamnesisTemplateItemsRenderer.tsx`: se `SCALE_REGISTRY[scoreItem.anamneseItemCode]` existir,
  renderiza `<ScaleWidget>` no lugar dos botões de banda — compact e fullscreen.
- **Rollout incremental do conteúdo** (framework já suporta todos desde o dia 1):
  - **Tier A — soma simples** (entrega o pedido): PHQ-9, GAD-7, Epworth, IIEF-5.
  - **Tier B — administração guiada**: Dubois ×2 (`WordRecallAdmin`), Span ×2
    (`DigitSpanAdmin`) — aplicáveis ao vivo no EMR. Conteúdo (palavras/sequências) validado
    pelo Dr. Getúlio antes de ativar.
  - **Tier C — custom** (score próprio + validação clínica pesada): FSS (média), ASEX
    (regra por item), FSFI (domínios ponderados), PSQI (componentes). Cada uma revisada
    individualmente com o Dr. Getúlio antes de ativar.
- Sem mudança de backend para a paridade. **Ao fim do tier A o EMR já iguala o site.**

### Fase 3 — Persistir respostas individuais (recomendado; se decisão A = SIM)
- Migration goose (próxima, ~00038): `ALTER TABLE anamnesis_items ADD COLUMN scale_responses JSONB`.
- Model `AnamnesisItem` + DTO request/response: campo `ScaleResponses`. `pnpm generate`
  (lembrar do caveat Swagger 2.x parcial).
- Widget grava `{answers:{...}, total}`; ao editar, reidrata do `scale_responses` (fallback:
  começa vazio). Motor de score inalterado.
- Exibição: na leitura da anamnese e no histórico longitudinal, mostrar o detalhe por pergunta;
  **destacar PHQ-9 Q9 ≥ 1 (ideação suicida)**.

### Fase 4 — Reuso (opcional)
- Site passa a consumir `@plenya/domain` (dedup do PHQ-9; habilita GAD-7/Epworth widget no site).
- Mesmo widget nos formulários pré-consulta do paciente (`ConsultationPrep`) e sessões anônimas.

## Arquivos-chave
- Renderer EMR: `apps/web/components/anamnesis/AnamnesisTemplateItemsRenderer.tsx`
  (decisão de tipo ~L843–930; `detectLevel` L36; compact L450–667; fullscreen L672–966).
- Tipos de form/payload: `AnamnesisTemplateItemsForm.tsx` (L7–14), `lib/api/anamnesis.ts` (L7–28).
- Widget site (referência/migrar): `apps/site/components/escore/phq9-widget.tsx`.
- Models: `apps/api/internal/models/{score_item,score_level,anamnesis_item}.go`.

## Riscos / notas
- O renderer tem lógica duplicada compact vs fullscreen; o widget precisa ser autocontido para
  entrar nas duas sem retrabalho.
- `scale_0_3` no enum sugere intenção original; vamos usar `anamneseItemCode` como gatilho
  (mais robusto, já populado) e, se quiser, marcar os itens com `scale_0_3` por consistência.
- Não tocar no motor de score: `selected_level` continua sendo a fonte da pontuação.
