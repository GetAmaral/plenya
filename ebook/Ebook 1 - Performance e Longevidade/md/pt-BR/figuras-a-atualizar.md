# Figuras — Mapa Completo e Pendências

Este arquivo é interno (não publica) e documenta o estado das figuras do livro após a reestruturação de 14 para 18 capítulos. O texto dos capítulos foi atualizado para usar a numeração nova; a arte física em `figuras/pt-BR/` ainda carrega, em parte, a numeração antiga.

---

## Correções já feitas no texto

- **Cap 8 / Fig 8.1** — referência corrigida de `Cap07 Fig03.PNG` para `Cap08 Fig01.PNG` (conteúdo: Marcos 8 meses depois).
- **Cap 15 / Fig 15.2** — legenda e referência corrigidas de "Figura 11.2 / Cap11 Fig02.PNG" para "Figura 15.2 / Cap15 Fig02.PNG" (conteúdo: rastreamento por década).

## Renames físicos executados em 2026-04-21

Todos os 12 arquivos PNG foram renomeados em `figuras/pt-BR/` em uma única passada decrescente (Cap12→Cap16, Cap11→Cap15, …, Cap07→Cap08). Não houve colisão porque cada destino estava sempre livre quando o rename foi feito.

**Mapa executado:**
- `Cap07 Fig03` → `Cap08 Fig01` (Marcos 8 meses)
- `Cap08 Fig01` → `Cap10 Fig01` (Paulo TT 310)
- `Cap08 Fig02` → `Cap10 Fig02` (IGF-1 em U)
- `Cap08 Fig03` → `Cap11 Fig01` (Paulo 6 meses)
- `Cap09 Fig01` → `Cap12 Fig01` (cascata HPA)
- `Cap09 Fig02` → `Cap13 Fig01` (solidão barras)
- `Cap10 Fig01` → `Cap14 Fig01` (sono arquitetura)
- `Cap10 Fig02` → `Cap14 Fig02` (regularidade vence duração)
- `Cap10 Fig03` → `Cap14 Fig03` (Paulo 4 tempos)
- `Cap11 Fig01` → `Cap15 Fig01` (Marcos placar)
- `Cap11 Fig02` → `Cap15 Fig02` (rastreamento década)
- `Cap12 Fig01` → `Cap16 Fig01` (dois modelos)

Verificação: todas as 30 figuras referenciadas no texto atual estão cobertas pelos arquivos no disco, exceto as 2 figuras novas ainda pendentes de criação (ver abaixo).

## Fig 1.1 — `Cap01 Fig01.PNG`

**Motivo:** os eixos estão rotulados com anos absolutos ("2021", "2026", "2028"). Isso faz o livro envelhecer mal — em 2–3 anos, os números de anos absolutos parecem datados.

**Ação:** regerar a figura com rótulos relativos no eixo temporal.

**Valores (mantidos):**
- Ano −5: HbA1c 4,9%
- Ano −2: HbA1c 5,2%
- Hoje (Ano 0): HbA1c 5,4%
- Projeção Ano +2: 5,7% (em linha tracejada)

**Novos rótulos sugeridos para o eixo:** *"há 5 anos → há 2 anos → hoje → em 2 anos"* ou simplesmente *"Ano −5 | Ano −2 | Hoje | Ano +2"*.

**Caption do capítulo já foi atualizada** para usar linguagem relativa.

---

## Fig 4.1 — `Cap04 Fig01.PNG`

**Motivo:** o caption originalmente dizia que o painel mostra "11 biomarcadores principais + 2 complementares", mas o texto do Cap 4 detalha 16 biomarcadores principais + 2 complementares. O texto ficou descompassado da figura.

**Status:** ajustado no texto via alteração do caption — agora explicita que a figura mostra os 11 mais subpedidos e que outros 5 (microalbuminúria, albumina sérica, colesterol não-HDL, troponina I ultrassensível, NT-proBNP) são detalhados apenas no corpo do capítulo.

**Ação opcional:** se quiser ampliar a figura para 16 + 2 biomarcadores no futuro, o caption pode ser simplificado de novo. Por enquanto, não bloqueia publicação.

---

## Mapa de renomeação pós-reestruturação

A reestruturação AGIR (14 → 18 capítulos) criou descompasso entre nome de arquivo físico e número do capítulo atual em várias figuras. A tabela abaixo descreve o estado atual e o destino final de cada arquivo.

### Figuras sem mudança (conteúdo e número batem)

| Arquivo físico | Usado em | Status |
|---|---|---|
| `Cap01 Fig01.PNG` | Cap 1 Fig 1.1 | ✓ mas precisa regerar (ver Fig 1.1 acima) |
| `Cap01 Fig02.PNG` | Cap 1 Fig 1.2 | ✓ |
| `Cap02 Fig01/02/03.PNG` | Cap 2 Fig 2.1/2.2/2.3 | ✓ |
| `Cap03 Fig01/02.PNG` | Cap 3 Fig 3.1/3.2 | ✓ |
| `Cap04 Fig01/02.PNG` | Cap 4 Fig 4.1/4.2 | ✓ |
| `Cap05 Fig01/02.PNG` | Cap 5 Fig 5.1/5.2 | ✓ |
| `Cap06 Fig01/02/03.PNG` | Cap 6 Fig 6.1/6.2/6.3 | ✓ |
| `Cap07 Fig01.PNG` | Cap 7 Fig 7.1 (VO₂ fator de risco) | ✓ |
| `Cap07 Fig02.PNG` | Cap 7 Fig 7.2 (Primeiro degrau) | ✓ |

### Figuras que precisam de rename físico

O arquivo hoje tem o **nome antigo** (do livro em 14 capítulos). Precisa ser renomeado para o **nome novo** (após restructure em 18 capítulos).

| Conteúdo | Arquivo atual | Arquivo novo | Usado em |
|---|---|---|---|
| Marcos 8 meses depois | `Cap07 Fig03.PNG` | **`Cap08 Fig01.PNG`** | Cap 8 Fig 8.1 |
| Paulo testosterona total 310 | `Cap08 Fig01.PNG` | **`Cap10 Fig01.PNG`** | Cap 10 Fig 10.1 |
| IGF-1 U invertido | `Cap08 Fig02.PNG` | **`Cap10 Fig02.PNG`** | Cap 10 Fig 10.2 |
| Paulo 6 meses (Cap G3) | `Cap08 Fig03.PNG` | **`Cap11 Fig01.PNG`** | Cap 11 Fig 11.1 |
| HPA cascade (ansiedade vira doença) | `Cap09 Fig01.PNG` | **`Cap12 Fig01.PNG`** ⚠️ *conflito: arquivo `Cap12 Fig01.PNG` já existe com tamanho diferente (687 KB vs 795 KB); decidir se é duplicata a apagar, ou se ambos servem* | Cap 12 Fig 12.1 |
| Solidão e mortalidade | `Cap09 Fig02.PNG` | **`Cap13 Fig01.PNG`** | Cap 13 Fig 13.1 |
| Sono arquitetura Paulo | `Cap10 Fig01.PNG` | **`Cap14 Fig01.PNG`** | Cap 14 Fig 14.1 |
| Regularidade vence duração | `Cap10 Fig02.PNG` | **`Cap14 Fig02.PNG`** | Cap 14 Fig 14.2 |
| Paulo 4 tempos (24 meses) | `Cap10 Fig03.PNG` | **`Cap14 Fig03.PNG`** | Cap 14 Fig 14.3 |
| Marcos placar | `Cap11 Fig01.PNG` | **`Cap15 Fig01.PNG`** | Cap 15 Fig 15.1 |
| Rastreamento por década | `Cap11 Fig02.PNG` | **`Cap15 Fig02.PNG`** | Cap 15 Fig 15.2 |

**Procedimento seguro de rename** (para evitar colisões, usar nome temporário e depois o final):

```bash
cd "ebook/Ebook 1 - Performance e Longevidade/figuras/pt-BR"

# 1. Primeiro, renomear para nomes temporários
git mv "Cap07 Fig03.PNG" "tmp_Marcos8m.PNG"
git mv "Cap08 Fig01.PNG" "tmp_PauloTT.PNG"
git mv "Cap08 Fig02.PNG" "tmp_IGF1.PNG"
git mv "Cap08 Fig03.PNG" "tmp_Paulo6m.PNG"
git mv "Cap09 Fig02.PNG" "tmp_Solidao.PNG"
git mv "Cap10 Fig01.PNG" "tmp_SonoArq.PNG"
git mv "Cap10 Fig02.PNG" "tmp_Regularidade.PNG"
git mv "Cap10 Fig03.PNG" "tmp_Paulo4t.PNG"
git mv "Cap11 Fig01.PNG" "tmp_MarcosPlacar.PNG"
git mv "Cap11 Fig02.PNG" "tmp_RastreioDecada.PNG"
# Cap09 Fig01 (HPA) → decidir conflito com Cap12 Fig01 antes de renomear

# 2. Depois, renomear de tmp para nome final
git mv "tmp_Marcos8m.PNG" "Cap08 Fig01.PNG"
git mv "tmp_PauloTT.PNG" "Cap10 Fig01.PNG"
git mv "tmp_IGF1.PNG" "Cap10 Fig02.PNG"
git mv "tmp_Paulo6m.PNG" "Cap11 Fig01.PNG"
git mv "tmp_Solidao.PNG" "Cap13 Fig01.PNG"
git mv "tmp_SonoArq.PNG" "Cap14 Fig01.PNG"
git mv "tmp_Regularidade.PNG" "Cap14 Fig02.PNG"
git mv "tmp_Paulo4t.PNG" "Cap14 Fig03.PNG"
git mv "tmp_MarcosPlacar.PNG" "Cap15 Fig01.PNG"
git mv "tmp_RastreioDecada.PNG" "Cap15 Fig02.PNG"
```

**Decisão pendente:** o arquivo `Cap12 Fig01.PNG` já existe no disco com tamanho diferente do `Cap09 Fig01.PNG`. Preciso saber se:
- (a) `Cap12 Fig01.PNG` existente é uma versão mais nova/aprovada do HPA cascade e deve ficar — nesse caso, deletar `Cap09 Fig01.PNG`;
- (b) `Cap12 Fig01.PNG` existente é outro conteúdo que por engano foi nomeado assim — nesse caso, renomear para outra coisa e promover `Cap09 Fig01.PNG` a `Cap12 Fig01.PNG`.

### Figuras novas aprovadas — 9 em produção

Briefing detalhado em `briefing-figuras-novas.md`. Todas têm referência e legenda já inseridas no texto dos capítulos e estão listadas no frontmatter. Quando os PNGs forem produzidos, basta salvar em `figuras/pt-BR/` com o nome exato.

| Conteúdo | Arquivo | Usado em | Referência no texto |
|---|---|---|---|
| Kraft II de André (curva dupla glicose/insulina) | `Cap06 Fig04.PNG` | Cap 6 Fig 6.4 | ✅ inserida |
| Algoritmo alopecia + PFS + timeline FDA/EMA | `Cap08 Fig02.PNG` | Cap 8 Fig 8.2 | ✅ inserida |
| Eixo cardio-reno-metabólico (Venn) | `Cap09 Fig01.PNG` | Cap 9 Fig 9.1 | ✅ inserida |
| Hipótese da janela TRH (timeline ELITE/KEEPS) | `Cap10 Fig03.PNG` | Cap 10 Fig 10.3 | ✅ inserida |
| Grid de 8 polimorfismos que mudam conduta | `Cap11 Fig02.PNG` | Cap 11 Fig 11.2 | ✅ inserida |
| Ana 6 meses (biomarcadores + PHQ-9/GAD-7) | `Cap12 Fig02.PNG` | Cap 12 Fig 12.2 | ✅ inserida |
| 5 instrumentos de triagem psicológica | `Cap12 Fig03.PNG` | Cap 12 Fig 12.3 | ✅ inserida |
| Ricardo em 3 tempos (Cap 13) | `Cap13 Fig02.PNG` | Cap 13 Fig 13.2 | ✅ inserida |
| Diagrama de Ikigai (4 círculos) | `Cap13 Fig03.PNG` | Cap 13 Fig 13.3 | ✅ inserida |

Total: 9 figuras novas em produção. Livro passará de 30 → 39 figuras (~2,2 por capítulo).

### Figuras já existentes no disco, aguardando referência

Opcional:
| Conteúdo | Arquivo | Uso | Status |
|---|---|---|---|
| Dois modelos de acompanhamento (check-up convencional × preventivo) | `Cap16 Fig01.PNG` | Cap 16 Fig 16.1 | ✅ no disco, referenciada no texto |

### Figuras que ficam obsoletas após o rename

`Cap09 Fig01.PNG` (HPA cascade) — se a decisão sobre o conflito com `Cap12 Fig01.PNG` for no sentido (a), esse arquivo deve ser deletado do repo para evitar confusão.

`Cap12 Fig01.PNG` atual — se a decisão for (b), deve ser renomeado.

---

## Resumo priorizado (atualizado 2026-04-21)

| # | Prioridade | Ação | Status |
|---|---|---|---|
| 1 | Concluído | Rename físico das 12 figuras (14 caps → 18 caps) | ✅ feito em commit 330edd4 |
| 2 | **Em andamento** | Produzir as 9 figuras novas do briefing (Cap06 Fig04, Cap08 Fig02, Cap09 Fig01, Cap10 Fig03, Cap11 Fig02, Cap12 Fig02, Cap12 Fig03, Cap13 Fig02, Cap13 Fig03) | 🎨 em produção |
| 3 | Média | Regerar `Cap01 Fig01.PNG` com eixo temporal relativo | pendente |
| 4 | Baixa | Opcional: ampliar `Cap04 Fig01.PNG` para 16 biomarcadores | pendente |

Enquanto as 9 figuras novas não forem produzidas e salvas em `figuras/pt-BR/`, as respectivas posições no EPUB vão aparecer como links quebrados. O texto já tem as legendas prontas — basta colocar os PNGs com o nome exato e o build funciona.
