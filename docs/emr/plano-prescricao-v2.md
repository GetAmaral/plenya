# Prescrição v2 — catálogo, ergonomia e manipulados

**Status:** proposta (nada implementado além do desbloqueio do item 1)
**Data:** 2026-08-19

## 1. Por que não salvava (corrigido, commit `1b76c821`)

`apps/web/lib/api/prescriptions.ts` chamava `POST /prescriptions` em vez de
`POST /api/v1/prescriptions`. Toda operação de prescrição batia 404 — listar, abrir, salvar,
assinar, excluir. O log da API em produção mostrava a batida errada repetida
(`GET /prescriptions 404`). O mesmo defeito existia em
`apps/web/lib/api/medication-definitions.ts`, que é por que a busca de medicamento nunca
achava nada. Eram os dois únicos módulos de `lib/api` sem o prefixo; o resto está correto.

## 2. Catálogo de medicamentos: nunca foi importado

`medication_definitions` está com **0 linhas em produção** e **não existe importador no
repositório** — nenhum `cmd/import-medications`, nenhum seed, nenhum CSV. A tela sempre
dependeu de cadastro manual, que nunca aconteceu. Corrigir a URL faz a busca funcionar, mas
ela continua sem resultados até popular.

### Fontes oficiais

| Fonte | O que tem | Serve para |
|---|---|---|
| [Dados Abertos ANVISA — registros válidos](https://dados.anvisa.gov.br) | registro, produto, princípio ativo, empresa, classe terapêutica | nome comercial ↔ princípio ativo, código ANVISA |
| [CMED — listas de preços](https://www.gov.br/anvisa/pt-br/assuntos/medicamentos/cmed/precos) | substância, produto, apresentação, laboratório, EAN, **tarja/lista de controle**, preços | apresentação/concentração e a categoria regulatória |

A **tarja** da CMED é o que permite preencher `category` (`simple` / `c1` / `c5` /
`antibiotic` / `glp1`) automaticamente — e é dela que saem as regras que o modelo já tem
(validade da receita, máximo de substâncias, exigência de assinatura e de SNCR).

### Importador proposto (`apps/api/cmd/import-medications`)

1. Baixar o arquivo oficial e **inspecionar o layout real antes de mapear** (os nomes de
   coluna mudam entre edições; nada de assumir).
2. Normalizar: `commonName` = produto + apresentação, `activeIngredient` = substância,
   `anvisaCode` = registro, `category` = derivada da tarja.
3. Campos novos que a tela precisa e o modelo ainda não tem: **forma farmacêutica**,
   **concentração**, **laboratório**, **apresentação**, **genérico/similar/referência**.
4. Idempotente por (registro + apresentação), rodável de novo a cada atualização da lista.
5. **Dedupe para a busca**: 25 mil apresentações poluem o autocomplete. Duas camadas —
   busca por **princípio ativo** (prescrever genérico pela DCB) e por **produto comercial**.

## 3. Ergonomia da prescrição

Hoje cada medicamento pede **11 campos**, quase todos texto livre: Nome Comercial, Princípio
Ativo, Categoria, Concentração, Via, Dosagem, Frequência, Duração, Quantidade, Quantidade por
Extenso (essa sim automática) e Instruções. Escolher do catálogo preenche só três (nome,
princípio, categoria) — o resto é digitação, toda vez, inclusive para o medicamento que o
médico prescreve todo dia.

| Problema | Proposta |
|---|---|
| Catálogo preenche pouco | Selecionar traz concentração, forma, via padrão, categoria e validade |
| Dose/frequência/duração em texto livre | Posologia estruturada com atalhos ("1 cp", "8/8h", "7 dias") que compõem a frase |
| Quantidade digitada à mão | Calculada de dose × frequência × duração (8/8h por 7 dias = 21), com override |
| "Comprimidos" fixo no extenso | Unidade vem da forma farmacêutica (gotas, mL, cápsulas) |
| Tudo do zero a cada receita | **Favoritos e repetir receita anterior** — o maior ganho real de tempo |
| Regras regulatórias invisíveis | Avisar no ato: C1 limita 3 substâncias/60 dias e exige assinatura |
| Princípio ativo e categoria editáveis | Read-only quando vêm do catálogo; texto livre só no modo manual |

Ordem de maior retorno: **repetir/favoritos > catálogo preenchendo > quantidade calculada >
posologia estruturada**.

## 4. Manipulados — hoje não existe

O modelo atual (`prescription_medications`) é uma linha por medicamento industrializado. Uma
fórmula manipulada é outra coisa: **um item composto de N substâncias**, cada uma com sua
concentração, mais forma farmacêutica, quantidade a aviar e uma posologia única — e vai em
receituário próprio, endereçado à farmácia de manipulação.

### Proposta

- `prescriptions.type`: `commercial` | `compounded`. O switch troca o formulário inteiro,
  como pedido.
- Item manipulado: tabela `prescription_formulas` (forma farmacêutica, quantidade a aviar,
  posologia, uso interno/externo) + `prescription_formula_components` (substância,
  quantidade, unidade). Tabela filha explícita em vez de JSONB: dá para consultar, validar e
  reaproveitar fórmula.
- **Fórmulas favoritas** reaproveitáveis entre pacientes — é o padrão de quem manipula.
- PDF em layout de manipulação: componentes listados, "aviar N cápsulas", posologia, via.
- Regras regulatórias continuam valendo por substância (uma fórmula com substância C1 é
  receituário de controle especial).

## 5. Ordem sugerida

1. ~~Desbloquear produção (URL)~~ — feito.
2. Catálogo ANVISA/CMED + busca que preenche os campos.
3. Manipulados (switch, modelo, formulário, PDF).
4. Ergonomia: repetir receita, favoritos, quantidade calculada, posologia estruturada.

O item 2 é pré-requisito real do 4: sem catálogo, não há o que autopreencher.
