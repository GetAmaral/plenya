# Plano — Metadados de material/espécime (e afins) em LabTestDefinition e LabResult

## Motivação
O bug #4 (glicose urinária "+++" classificada como glicemia sanguínea) é de **ambiguidade
de nome**: "Glicose" sozinho não diz se é sangue ou urina. A disambiguação correta vem do
**material/espécime**, que hoje a extração até captura mas o sistema joga fora. Registrar o
material **na definição** (alvo) e **no resultado** (origem) mata essa classe de erro de raiz
e enriquece o prontuário. Tudo opcional (omitido quando o PDF não traz).

## O que já existe (não reinventar)
- **Extração** (`dto.PDFExtractionExam`) já traz, por exame: `nomeExame`, `resultado`,
  `unidade`, **`material`** (ex.: "Urina", "PlasmFluoreto", "Sangue"), **`metodo`**
  (ex.: "Enzimático UV - Hexoquinase"). E no laudo: `laboratorio`, `dataColeta`.
- **`lab_test_definitions`** já tem a coluna **`specimen_type`** (varchar 100) — mas **vazia
  em 662 de 668 itens**. Também tem `result_type`, `category`, `unit`, `unit_conversion`,
  `collection_method`, `fasting_hours`, `loinc_code`, `tuss_code`.
- **`lab_results`** **NÃO** tem material/método/referência/data-por-exame. O
  `createLabResultsFromJSON` ignora `material` e `metodo` (sem coluna onde gravar).
- **Prompt** manda *descartar* "Valores de Referência".

## Campos a registrar (todos opcionais)

### Em `lab_test_definitions` (o alvo do match)
| campo | situação | uso |
|---|---|---|
| `specimen_type` | **existe, popular** | espécime canônico (Sangue/Soro/Plasma/Urina/Fezes/Saliva…). **Chave da disambiguação.** |

(os demais já existem: `unit`, `result_type`, `collection_method`, `fasting_hours`.)

### Em `lab_results` (o registro do resultado) — ADICIONAR colunas nullable
| campo | de onde | valor |
|---|---|---|
| `specimen` | extração `material` (normalizado) | Sangue/Soro/Plasma/Urina/Fezes… — disambiguação + prontuário |
| `method` | extração `metodo` | método analítico (HbA1c por HPLC×imuno muda interpretação; FAN título; etc.) |
| `reference_range` | NOVO no prompt | faixa de referência impressa do lab (ex.: "70 a 99 mg/dL") — exibição + **fallback de normal/alterado quando o exame não é item de score ou não classificou** |
| `collection_date` | extração por-exame ("Coletado em:") | **REQUISITO FIRME (não opcional na prática).** Os laudos trazem data por exame, que pode diferir da do lote. A **tabela de resultados deve renderizar/ordenar por data DO RESULTADO**, não do lote (como é hoje). **Regra de fallback: se nulo, usar a data do lote** (`coalesce(lr.collection_date, batch.collection_date)`). Coluna fica nullable; null = "herda do lote". |

### Na extração (prompt + DTO)
- Já vem `material` e `metodo` — **parar de descartar**, gravar no result.
- ADICIONAR ao DTO/prompt: `valorReferencia` (faixa de ref) e, opcional, `dataColetaExame`.
- O prompt continua "tudo opcional, omitir se não achar".

## Matching ciente de espécime (o coração do conserto)
1. Normalizador `normalizeSpecimen(material)` → categoria canônica: "Urina"; "Sangue"
   (Sangue/Soro/Plasma/PlasmFluoreto/PlasmEDTA/Sangue Total → família sangue); "Fezes";
   "Saliva"; "Líquor"; … (mesmo vocabulário do `specimen_type` da definição).
2. Em `matchTestDefinition(nome, material, defMap)`: quando houver **mais de uma** definição
   candidata (exato/substring/fuzzy) **ou** ambiguidade conhecida, **preferir a definição cujo
   `specimen_type` casa com o `material`** do resultado. Ex.: "Glicose" + Urina → "Glicose
   Urinária"; "Glicose" + Plasma → "Glicemia de jejum".
3. Se a definição não tem `specimen_type` (ainda não populado), cai no comportamento atual
   (compatível com a migração gradual).

## Populando `specimen_type` (o maior esforço)
~668 definições. Estratégia faseada:
- **Heurística automática:** nome contém "urin"/"urinária"/"EAS" → Urina; "fezes"/"parasit" →
  Fezes; genéticos (rs####/genes) → "DNA"/ignorar; default por `category` → Sangue/Soro.
- **Revisão pontual** dos que a heurística não resolver (script gera CSV "definição → palpite",
  revisamos os duvidosos).
- Roda como **migração de catálogo** (o catálogo não é versionado; vai na mesma migração dos
  alt_names dos fixes #1/#2).

## Fases sugeridas
1. **Schema:** migração goose — add `specimen`, `method`, `reference_range`, `collection_date`
   (nullable) em `lab_results`. Atualiza model `lab_result.go` + DTO + `pnpm generate`.
2. **Extração:** prompt extrai `valorReferencia` (+ dataColetaExame); DTO ganha os campos.
3. **Mapeamento:** `createLabResultsFromJSON` grava `specimen` (normalizado), `method`,
   `reference_range`, `collection_date` no result (este último = data por-exame quando houver,
   senão null/herda do lote).
4. **Populando specimen_type:** script heurístico + revisão → migração de catálogo.
5. **Matching:** `matchTestDefinition` passa a receber o material e desempatar por espécime.
6. **Frontend (apps/web):** tabela de resultados passa a **agrupar/ordenar por data do
   RESULTADO** usando `coalesce(lr.collection_date, batch.collection_date)`, no lugar da data
   do lote. (Expor `collectionDate` por result no DTO; `pnpm generate`.)
   Alvo: feature `lab-result-views` (`app/(authenticated)/lab-result-views/` +
   `components/lab-results/`) — a visão longitudinal que hoje organiza por lote. Conferir o
   agrupamento no backend da view também (não só o componente).
7. **Validar** com os 4 laudos do João (glicose urina vs sangue separados; sem vazamento; e a
   tabela mostrando datas por-exame corretas, com fallback pro lote quando ausente).

## Não-objetivos / cuidados
- Não tornar nada obrigatório (laudos variam muito).
- `reference_range` é texto livre do lab (exibição/fallback), **não** substitui as faixas do
  score; é rede de segurança pra exibir normal/alterado quando não há item de score.
- Conversão de unidade (/mm3→k/µL) já existe e é independente disso.
