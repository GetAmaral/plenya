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
| `00026_snapshot_source_session_unique.sql` | migration goose | `cmd/migrate up` (hardening Fase 4: troca o índice de `source_session_id` por ÚNICO PARCIAL `uq_snapshot_source_session` — no máx. 1 import vivo por sessão; fecha corrida de import concorrente). |
| `00027_fix_version_items_prod.sql` | migration goose | `cmd/migrate up` (**hardening pré-deploy, BLOCKER**: a 00022 deixaria a Light VAZIA em prod — `is_light_version`/`light_order` são dev-only. A 00027 faz upsert dos 119 vínculos por UUID+ordem, EXISTS-filtered/idempotente, e corrige a ordem da Triagem. Sem ela, o Escore Light público dá 500). |
| `apps/web/Dockerfile` + `apps/site/Dockerfile` | build de app | Passaram a copiar `packages/ui` (deps + runner). Sem isso, `pnpm install --frozen-lockfile` quebra no Coolify (a Fase 0 adicionou `@plenya/ui` como workspace member de web e site). Já corrigidos no repo — sobem com o deploy normal. |
| `docs/emr/site-fields-seed.sql` | seed idempotente | `psql -f` contra o banco prod (preenche `site_explanation` em 1211 itens, `site_legend` em 3824 níveis, em linguagem leiga + acentos) |
| `apps/site/content/data/score-{triagem,light}-config.json` | JSON estático | já no repo — o site **lê do arquivo**, sem dependência de banco no display. Nada a fazer no deploy. |

## Pré-condição crítica: UUIDs dev ≡ prod

O seed (`site-fields-seed.sql`) e a migration `00022` casam linhas **por UUID** de `score_item`
(`score_items` não tem chave natural — só `id`, `name`, `points`). Isso pressupõe que prod e dev
compartilham os mesmos UUIDs de score — o que já é verdade (os 1211 itens vieram do mesmo dump;
a `00022` hardcoda 36 UUIDs de itens e roda nos dois ambientes). **Mesmo assim, faça o pré-flight
abaixo antes do seed.**

## Passos (prod)

> **Por que a 00027 existe:** a 00022 semeia a version **Light** via `is_light_version=true` e ordena
> ambas por `light_order` — MAS esses flags foram setados só no **dev** (script avulso
> `docs/escore-light/light-curation-v5.sql`, nunca migrado pra prod) e a coluna some na 00023. Sem a
> **00027** a Light entraria VAZIA em prod → o Escore Light público daria **500** em toda submissão
> (`GetVersionItemIDsBySlug("light")` vazio). A 00027 torna as duas versions self-contained (119 itens
> por UUID + ordem, idempotente, EXISTS-filtered). Por isso o gate de version_items abaixo é CRÍTICO.

```bash
# 0. PRÉ-FLIGHT — confirmar que os UUIDs de score_item do dev existem no prod (dev≡prod).
#    Pegue alguns ids no dev e rode no prod; todos têm que existir.
psql "$PROD_DATABASE_URL" -c "SELECT count(*) FROM score_items WHERE id IN ('<uuid1>','<uuid2>','<uuid3>');"
#    -> precisa retornar 3. Se < 3, PARE: UUIDs divergiram; a 00022/00027 e o seed precisam re-chavear.

# 1. Migrations — aplica TODAS as pendentes (00021→00027) em ordem, idempotentes/aditivas.
go run ./cmd/migrate status
go run ./cmd/migrate up
go run ./cmd/migrate status   # confirmar 00027 aplicada (a ÚLTIMA)

# 2. Seed de conteúdo de site (idempotente — só preenche colunas IS NULL; reexecutável).
psql "$PROD_DATABASE_URL" -f docs/emr/site-fields-seed.sql

# 3. VERIFICAÇÃO (gate de go/no-go).
psql "$PROD_DATABASE_URL" -c "
  SELECT
    (SELECT count(*) FROM score_version_items svi JOIN score_versions sv ON sv.id=svi.version_id WHERE sv.slug='light')   light_items,
    (SELECT count(*) FROM score_version_items svi JOIN score_versions sv ON sv.id=svi.version_id WHERE sv.slug='triagem') triagem_items,
    (SELECT count(*) FROM score_items  WHERE deleted_at IS NULL AND site_explanation IS NOT NULL) items_site,
    (SELECT count(*) FROM score_levels WHERE deleted_at IS NULL AND site_legend      IS NOT NULL) levels_site,
    (SELECT to_regclass('public.anonymous_score_item_results')) anon_results_tbl,
    (SELECT count(*) FROM score_item_method_pillars) agir_links;"
#  GATE CRÍTICO: light_items=83 e triagem_items=36. Se light_items=0 → a 00027 não casou
#    (UUID mismatch) → o Escore Light está QUEBRADO. NÃO liberar tráfego; investigar os UUIDs.
#  anon_results_tbl deve ser 'anonymous_score_item_results' (não NULL).
#  items_site/levels_site: esperado 1211/3824 SE a distribuição AGIR+Bioma estiver em prod. Se o
#    Bioma (+297 itens b10) NÃO foi deployado, esperar ~914/3488 — NÃO é erro, só menos itens.
#  agir_links: se 0 → o radar anônimo por pilar (Fase 3) cai no FALLBACK por grupo (esperado até a
#    distribuição AGIR — workstream agir_distribuicao_bioma — ir a prod). Não bloqueia.

# 4. Configs estáticos do site: já estão commitados em apps/site/content/data/score-{light,triagem}-
#    config.json (gerados da fonte). O deploy do site (next build) os serve direto — sem passo de DB.
```

> **Dependência cross-workstream (Fase 3):** o radar por PILAR no site/EMR precisa dos vínculos
> `score_item_method_pillars` (distribuição AGIR), que é um workstream SEPARADO e ainda **dev-only**.
> Sem eles, o resultado renderiza o radar por **grupo** (fallback, sem erro). Para "acender" o radar
> por pilar em prod, deployar a distribuição AGIR antes/junto (ver memória `agir_distribuicao_bioma_status`).

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
