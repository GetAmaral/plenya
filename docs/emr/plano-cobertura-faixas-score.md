# Plano — Cobertura das faixas de score (ScoreItem → ScoreLevels)

> **Status:** investigação concluída (read-only) · **nada implementado** · aguardando aprovação.
> **Origem:** análise dos laudos do paciente João (4 PDFs Lab Oswaldo Cruz) mostrou dezenas de
> exames caindo em "valor fora das faixas de classificação configuradas". Este documento mapeia a
> causa-raiz da não-cobertura e propõe a correção, **sem alterar código nem dados ainda**.

## 1. Método (scan read-only, banco de produção)
- 276 **ScoreItems numéricos** ligados a lab (`lab_test_code` preenchido) + 2.697 `score_levels`.
- 394 itens **categóricos/genéticos** (`GEN_*`, operador `=` sem limites) foram **excluídos** — não
  usam faixa numérica (classificam por `selected_level` da anamnese, ver `emr_score_anamnese_selected_level`).
- Para cada item, montei os intervalos cobertos por `ScoreLevel.EvaluatesTrue` e procurei valores sem nível.
- Scripts da análise em scratchpad (não versionados): `analyze_gaps.py`, `review_checks.py`.

## 2. Resultado headline
**274 de 276 itens numéricos (99%)** têm valores que não classificam. Duas causas-raiz **independentes**
+ um achado clínico paralelo:

| # | Causa | Tipo | Itens | Correção |
|---|-------|------|-------|----------|
| A | Níveis `<` / `<=` mortos | **código** | 198 (252 níveis) | 1 ajuste em `EvaluatesTrue` |
| B | Vãos fracionários entre faixas | **dados** | 194 | regenerar faixas contíguas |
| C | Ordenação clínica não-monotônica | **dados/clínico** | 148 (maioria legítima) | auditoria, gate da Causa A |

Efeito combinado estimado (read-only):

| Cenário | Itens com gap |
|---|---|
| Hoje | 274 / 276 |
| Só corrigindo o código (Causa A) | 202 / 276 |
| Corrigindo código + dados (A + B) | ~6 (extremos pontuais) |

## 3. Causa A — BUG DE CÓDIGO: faixas `<` e `<=` estão mortas
**Evidência.** A convenção dos dados é consistente:
- `>X` / `>=X` → X no **`lower_limit`** (190 + 74 = 264 níveis).
- `<X` / `<=X` → X no **`upper_limit`**, `lower_limit` vazio (248 + 4 = **252 níveis**, **100%** nesse formato).

Mas `ScoreLevel.EvaluatesTrue` (`internal/models/score_level.go`) lê **`LowerLimit`** para `<` e `<=`:

```go
case "<":
    if sl.LowerLimit == nil { return false }
    return value < parseFloat(sl.LowerLimit)   // LowerLimit vazio → 0 → value < 0 → nunca true
```

Como o limite real está no `upper_limit`, **todo nível `<`/`<=` nunca dispara**. Resultado: a faixa
"abaixo de X" está **inoperante em 198 itens**.
- Derrubou os exames do João: **LDL 31** (`<70` morta), **Colesterol Total 111** (`<120` morta),
  **ALT 19** (`<26` morta).
- Outros exemplos: Bilirrubina direta `<0,2`, CEA `<1,3`, Relação CT/HDL `<3,0`, Coenzima Q10 `<0,50`.

**Referência autoritativa (validada no review).** A semântica pretendida já existe **correta** em
`packages/domain/src/score.ts` → `matchLevel`, que para `<`/`<=` lê `maxValue` (upper) e para `>`/`>=`
lê `minValue` (lower). O bug é **só do Go**, que diverge dessa referência. `matchLevel` é usado pelo
`apps/mobile-pro` → ou seja, hoje **mobile-pro e EMR-web classificam `<` de forma diferente** (mesma
config, dois motores). O escore-light do `apps/site` tem tipos próprios (`lib/score-light/types.ts`) —
**verificar o caminho dele à parte** (pode ter avaliador próprio).

**Impacto além do lab.** `EvaluatesTrue` é chamado em **4 lugares**: classificação de lab
(`lab_result_batch_service.go:938`), **snapshots de score** (`score_snapshot_service.go:337,397`) e
**escore anônimo/light** (`anonymous_score_service.go:1195`). O bug mis-classifica em todos; o fix
conserta os 4 de uma vez (fonte única).

**Correção proposta (não aplicada).** Em `EvaluatesTrue`, para `<` e `<=` ler `UpperLimit`
(espelhando `>`/`>=` que leem `LowerLimit` e a referência `matchLevel`), com fallback defensivo para
`LowerLimit` se `UpperLimit` for nulo:

```go
case "<", "<=":
    lim := sl.UpperLimit
    if lim == nil { lim = sl.LowerLimit }   // defensivo; dados hoje têm 100% no upper
    if lim == nil { return false }
    v := parseFloat(lim)
    if sl.Operator == "<" { return value < v }
    return value <= v
```

É **1 função**, sem migration. Fecha o "fundo" de 72 itens.

### 3.1. Blast radius (validado) — por que precisa de gate clínico
O fix **acorda 252 níveis hoje dormentes**. Distribuição por número de nível:

| Nível | Qtd `<`/`<=` | Significado ao acordar |
|---|---|---|
| 5 (melhor) | 78 | "baixo = bom" (LDL, CEA) — **seguro/correto** |
| 0–1 (pior) | 153 | "baixo = ruim" (Vit D <10, PTH <10, Cálcio <8, FC <40, Leucócitos <2,5) — **clinicamente correto**: hoje patologia de extremo baixo cai silenciosamente "fora" |
| 2–4 | 21 | intermediário |

Conclusão: o fix é **ganho líquido** — passa a classificar corretamente o extremo baixo. O risco fica
**restrito aos itens mal-ordenados** (Causa C), que ao acordar virariam falso crítico.

## 4. Causa B — DADOS: vãos fracionários entre faixas (194 itens)
As faixas `between` usam **fronteiras inteiras consecutivas**: `[10–19]` e `[20–39]` deixam a fração
**(19, 20)** descoberta. Como `between` é inclusivo nos dois lados, qualquer decimal na fenda não classifica.
- Exemplos: Vit D **19,5** · Colesterol Total **159,5** · Albumina **3,45** · ApoB/ApoA1 **0,695**.
- **Permanece mesmo corrigida a Causa A** (são buracos entre dois `between`): 194 itens.

**Correção proposta (não aplicada).** Tornar as faixas **contíguas** — borda superior de uma = borda
inferior da próxima. Duas opções:
- **(b1)** Repetir o número na junção (`[120–160]`, `[160–200]`): mínimo esforço; o `between` inclusivo
  aceita o ponto de junção (empate resolvido pelo loop "primeiro nível por `level ASC`").
- **(b2)** Padronizar os níveis como **partição por threshold** (só `>=`/`<` ordenados, sem `between`):
  elimina por construção sobreposição e buraco; mais limpo, porém mexe em mais linhas.

As 194 seguem o mesmo padrão → **regenerável em massa por script de dados** (migration goose ou
script Go one-off no padrão `cmd/`). Recomendo **(b1)** pelo menor risco.

## 5. Causa C — ordenação clínica não-monotônica (148 itens) — auditoria
148 itens têm nível × valor **não-monotônico**. A **maioria é legítima** (U-shaped: extremos alto E
baixo são ruins — Vit D toxicidade, cálcio, sódio). Mas alguns parecem **mal-centrados**:
- **Colesterol Total**: `<120` = nível 0 (pior) e `200–249` = nível 5 (melhor). Clinicamente, desejável
  é `<200`; `<120` não deveria ser o pior. **Provável erro de dados.**

**Gate obrigatório:** antes de ativar a Causa A em produção, **auditar clinicamente** os 153 níveis
`<`/`<=` em nível 0/1 e os itens não-monotônicos mal-centrados (Colesterol Total no mínimo), senão o fix
de código transforma os erros de ordenação em **falsos críticos**. A auditoria é do Dr. Getúlio (ou lista
priorizada por nós para ele validar).

## 6. Achados menores
- **3 sobreposições reais** de faixas (RDW, Ecodopplercardiograma-GLS, Doppler Carótidas-PSV) — o loop
  "primeiro nível que casa" resolve, mas pode escolher nível subótimo. Revisar caso a caso.
- **~6 itens** sem faixa de extremo (3 sem topo aberto, 3 sem fundo) após corrigir A+B — caso a caso.

## 7. Ordem de execução proposta (a aprovar — NÃO executar ainda)
1. **Auditoria clínica (gate)** — lista dos 153 níveis `<`/`<=` em L0/L1 + itens mal-centrados → validação do Dr. Getúlio. Corrigir os comprovadamente errados (ex.: Colesterol Total) **antes** do fix de código.
2. **Fix A (código)** — `EvaluatesTrue` lê `UpperLimit` para `<`/`<=`. + teste unitário cobrindo `<`, `<=`, `between`, `>`, `>=`.
3. **Fix B (dados)** — script que torna as 194 faixas contíguas (opção b1). Idempotente, dry-run primeiro.
4. **Fix C menor** — 6 extremos + 3 sobreposições, caso a caso.
5. **Re-scan de cobertura** → esperar ~0 gaps.

## 8. Verificação
- Teste unitário de `EvaluatesTrue` (todos os operadores; caso `<` com limite no upper).
- Build no container: `docker compose exec -w /app api go build ./...`.
- Re-rodar o scan de cobertura (scripts do scratchpad) → confirmar queda de 274 → ~6.
- **Dev:** re-classificar os 4 lotes do João → LDL/Colesterol/ALT passam a classificar **e** sem
  falso crítico (após auditoria da Causa C). QA via Playwright (`emr_qa_visual_playwright`).
- **Deploy:** manual por-app, só `api` (Fix A) e dados (Fix B via migration/script), **com ordem
  explícita** (`plenya_no_deploy_sem_ordem`).

## 9. Riscos
- **Falso crítico** se ativar Causa A sem o gate clínico (Causa C). → mitigado pela ordem (1 antes de 2).
- **Empates** na opção b1 (junção repetida) resolvidos de forma determinística pelo loop `level ASC`;
  validar que o nível escolhido na junção é o desejado.
- **Dados existentes não se re-classificam sozinhos.** Mudar `EvaluatesTrue`/faixas afeta classificações
  **futuras**; os `lab_results` já gravados continuam com `level` nulo (caíram fora). Para os existentes
  se beneficiarem, é preciso **re-rodar `ClassifyBatchResults`** nos lotes afetados (e re-snapshot dos
  scores que usaram faixas `<`/`<=`). Operacionalmente: re-classificar em massa pós-fix, depois
  `recomputeBatchCriticality`. Avaliar custo/escopo (quantos lotes/pacientes).
- **Inconsistência cross-plataforma** durante a transição: até o Go ser corrigido e os dados
  re-classificados, mobile-pro (matchLevel) e EMR-web (Go) podem mostrar níveis diferentes para o mesmo
  exame `<`/`<=`.

---
**Não faz parte deste escopo** (rastreado no report da importação de laudos): exames de urina/sedimento
e diferencial leucocitário casados com definição de sangue → falsos críticos por **matching/material**,
não por cobertura de faixa.
