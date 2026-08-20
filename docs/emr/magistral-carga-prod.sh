#!/usr/bin/env bash
# Carga do catálogo magistral em um banco Plenya, na ordem em que os arquivos dependem uns dos
# outros.
#
# PRODUÇÃO NÃO USA ESTE SCRIPT: o conteúdo vai na migration 00081_dados_catalogo_magistral, que
# roda sozinha no deploy. Este arquivo é a FONTE dela — a ordem de dependência mora aqui, e a
# migration é gerada concatenando estes arquivos nesta ordem. Serve para recarregar um banco de
# desenvolvimento e para conferir a ordem antes de regerar a migration.
#
# Todos os arquivos são idempotentes: rodar duas vezes não duplica. A ordem importa porque as
# fórmulas referenciam substâncias e as regras referenciam componentes de fórmula.
#
# Uso:
#   PGDATABASE=... PGUSER=... PGHOST=... ./magistral-carga-prod.sh          # aplica
#   ./magistral-carga-prod.sh --dry-run                                     # só lista a ordem
set -euo pipefail

DIR="$(cd "$(dirname "$0")" && pwd)"
DRY=""
[ "${1:-}" = "--dry-run" ] && DRY=1

# A ordem é a dependência real, não a alfabética.
ARQUIVOS=(
  # 1. substâncias — o catálogo primeiro, porque tudo aponta para ele
  magistral-seed-inicial.sql
  magistral-seed-funcional.sql
  magistral-indicacoes-externas.sql
  magistral-densidades-aproximadas.sql
  magistral-magnesio-formas.sql
  magistral-formas-preferidas-e-correcao.sql
  magistral-substancias-formulario.sql
  magistral-substancias-nomes.sql
  magistral-glp1.sql
  magistral-pentravan.sql
  magistral-peptideos.sql
  magistral-arquitetura-hormonal.sql
  magistral-avulsos-capturados.sql
  magistral-faixas-corrigidas.sql

  # 2. norma e incompatibilidades — referenciam substâncias
  in28-anexo-iv.sql
  magistral-in28-mapa.sql
  magistral-incompat-base.sql

  # 3. fórmulas-base — referenciam substâncias
  magistral-formulas-base-seed.sql
  magistral-formulas-parceiras-curadas.sql
  magistral-formulario-parceiras-completo.sql
  magistral-glp1-formulas.sql
  magistral-pentravan-formulas.sql
  magistral-arquitetura-hormonal-formulas.sql
  magistral-formulario-correcoes.sql
  magistral-conferencia-zerada.sql

  # 4. regras de dose — referenciam componentes das fórmulas
  magistral-regras-dose-dinamica.sql
  magistral-regras-dose-expansao.sql
)

for f in "${ARQUIVOS[@]}"; do
  if [ ! -f "$DIR/$f" ]; then
    echo "FALTA: $f" >&2
    exit 1
  fi
  if [ -n "$DRY" ]; then
    echo "  $f"
    continue
  fi
  echo "→ $f"
  psql -v ON_ERROR_STOP=1 -q -f "$DIR/$f"
done

[ -n "$DRY" ] && { echo "${#ARQUIVOS[@]} arquivos, nesta ordem."; exit 0; }

psql -q -c "SELECT
  (SELECT count(*) FROM magistral_components)                                        AS substancias,
  (SELECT count(*) FROM magistral_formula_templates WHERE deleted_at IS NULL)        AS formulas,
  (SELECT count(*) FROM magistral_formula_template_rules)                            AS regras,
  (SELECT count(*) FROM in28_limits)                                                 AS tetos_in28;"
