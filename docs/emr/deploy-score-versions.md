# Deploy — Score Versions + campos de site (prod)

Runbook para levar ao **prod** as mudanças do workstream *score-versions* (Fases 1–2c +
preenchimento de conteúdo de site). Tudo já está no `master`; falta aplicar no banco de produção.

## O que mudou

| Artefato | Tipo | Aplica em prod via |
|---|---|---|
| `00021_score_site_fields.sql` | migration goose | `cmd/migrate up` (adiciona colunas `site_render_type/site_question/site_explanation` em `score_items` e `site_legend` em `score_levels`; backfill `site_question←light_question`) |
| `00022_score_versions.sql` | migration goose | `cmd/migrate up` (cria `score_versions` + `score_version_items`; **seed** de 2 versions — Triagem 36 itens, Light 83 itens — e 119 vínculos a partir de `is_light_version`/`light_order`) |
| `00023_drop_score_item_light_fields.sql` | migration goose | `cmd/migrate up` (Fase 2e: dropa `is_light_version`/`light_order`/`light_question` de `score_items`). **Roda DEPOIS da 00022** (que ainda lê essas colunas pra semear). O backend passou a resolver o "conjunto Light" pela `score_version` slug="light" (scoring anônimo + PDF de labs). |
| `00024_anonymous_score_item_results.sql` | migration goose | `cmd/migrate up` (Fase 3: cria `anonymous_score_item_results` p/ o radar anônimo por pilar). Aditiva, sem dependência de seed. O radar por pilar no site só "acende" quando a distribuição AGIR (`score_item_method_pillars`) estiver em prod; sem pilares, o resultado cai no fallback por grupo. |
| `00025_snapshot_source.sql` | migration goose | `cmd/migrate up` (Fase 4: `patient_score_snapshots` += `source` + `source_session_id`, p/ marcar e dar idempotência ao import anônimo→prontuário). Aditiva. |
| `docs/emr/site-fields-seed.sql` | seed idempotente | `psql -f` contra o banco prod (preenche `site_explanation` em 1211 itens, `site_legend` em 3824 níveis, em linguagem leiga + acentos) |
| `apps/site/content/data/score-{triagem,light}-config.json` | JSON estático | já no repo — o site **lê do arquivo**, sem dependência de banco no display. Nada a fazer no deploy. |

## Pré-condição crítica: UUIDs dev ≡ prod

O seed (`site-fields-seed.sql`) e a migration `00022` casam linhas **por UUID** de `score_item`
(`score_items` não tem chave natural — só `id`, `name`, `points`). Isso pressupõe que prod e dev
compartilham os mesmos UUIDs de score — o que já é verdade (os 1211 itens vieram do mesmo dump;
a `00022` hardcoda 36 UUIDs de itens e roda nos dois ambientes). **Mesmo assim, faça o pré-flight
abaixo antes do seed.**

## Passos (prod)

```bash
# 0. PRÉ-FLIGHT — confirmar que um UUID de item do dev existe no prod.
#    (pegue qualquer id de score_items no dev e rode no prod)
psql "$PROD_DATABASE_URL" -c "SELECT id, name FROM score_items WHERE id = '<UUID_DE_UM_ITEM_DO_DEV>';"
#    -> tem que retornar 1 linha. Se retornar 0, PARE: os UUIDs divergiram, o seed precisa
#       ser re-chaveado (e a 00022 também já estaria quebrada). Investigar antes de prosseguir.

# 1. Migrations (aplica 00021 + 00022; idempotentes/aditivas).
#    No container api do prod:
go run ./cmd/migrate status
go run ./cmd/migrate up
go run ./cmd/migrate status   # confirmar 00022 aplicada

# 2. Seed de conteúdo de site (idempotente — só preenche colunas IS NULL; reexecutável).
psql "$PROD_DATABASE_URL" -f docs/emr/site-fields-seed.sql

# 3. Verificação.
psql "$PROD_DATABASE_URL" -c "
  SELECT
    (SELECT count(*) FROM score_items  WHERE deleted_at IS NULL AND site_explanation IS NOT NULL) items_site,
    (SELECT count(*) FROM score_levels WHERE deleted_at IS NULL AND site_legend      IS NOT NULL) levels_site,
    (SELECT count(*) FROM score_versions) versions,
    (SELECT count(*) FROM score_version_items) version_items;"
#    Esperado (igual ao dev): items_site=1211, levels_site=3824, versions=2, version_items=119
```

## Backup de segurança

Antes do deploy, há backup completo do **dev** em `backups/` (gitignored, local):

- `plenya_dev_FULL_<ts>.dump` — dump custom-format comprimido (`pg_restore`), disaster recovery.
- `plenya_dev_schema_<ts>.sql` — schema-only (referência).
- `plenya_dev_score_versions_<ts>.sql` — `score_versions`+`score_version_items` (data-only, portável).

> ⚠️ O dump FULL do dev **não** é para restaurar sobre o prod (clobraria dados de paciente).
> Prod recebe as mudanças via **migrations + seed idempotente**, nunca via restore do dev.
> Antes de mexer no prod, tirar um `pg_dump` do **próprio prod**.

## Rollback

- `00021` / `00022` têm `Down` (goose) — revertem colunas/tabelas. Conteúdo de site é aditivo
  (colunas novas); reverter a `00021` derruba as colunas e o conteúdo junto (esperado).
- O seed é idempotente: reexecutar não duplica nada (`WHERE ... IS NULL`).
