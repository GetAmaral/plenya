# Mobile — Política de OTA (EAS Update)

Contexto: EMR é regulado. OTA sem disciplina cria risco clínico e regulatório.

## Regras gerais

- `runtimeVersion.policy = "appVersion"` — atualizações JS só se aplicam a builds nativos
  com a mesma `version`. Bump em `app.json:version` = novo binário, não OTA
- Canais: `development`, `preview`, `production`
- `fallbackToCacheTimeout: 0` — se update demora, app usa cache (offline-safe)

## O que PODE ir via OTA

- Bugfix JS puro (ex.: loop infinito, render error, null deref)
- Copy, textos, tradução
- Ajustes de estilo / layout
- Telas de marketing / promoção
- Ajuste de feature flag default

## O que NÃO PODE ir via OTA (exige release nativo)

- Mudança em cálculo de escore, dose, pontuação clínica
- Mudança em lógica de decisão clínica (ex.: flag de risco)
- Mudança de schema de request/response (pode divergir de API)
- Features que pedem nova permissão nativa (câmera, biometria, push)
- Atualização de dependências nativas
- Mudanças em segurança (cert pinning, biometria, screenshot prevention)

## Procedimento de release OTA

1. Criar branch `fix/...` a partir de `main`
2. Commit + PR + review
3. Merge em `main`
4. `pnpm --filter @plenya/mobile-pro update:preview` → smoke test interno 24h
5. Monitor Sentry crash-free > 99.5% no canal preview
6. Promover pra `production` com rollout gradual:
   - `eas update --branch production --rollout 10` — 10%
   - Aguardar 4h → métricas OK → `--rollout 50`
   - Aguardar 4h → métricas OK → `--rollout 100`
7. Documentar no changelog (`apps/mobile-pro/CHANGELOG.md`)

## Rollback

- `eas update --branch production --republish <previous-update-id>`
- Ou kill-switch via `GET /api/v1/mobile/config:killSwitch.enabled = true` (bloqueia app até
  release nativo)

## Validação antes de qualquer OTA

- [ ] `pnpm --filter @plenya/mobile-pro lint` limpo
- [ ] Smoke Maestro passou em iOS + Android (preview build)
- [ ] Nenhum dos itens "NÃO PODE" tocado
- [ ] CHANGELOG atualizado
