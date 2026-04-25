# Mobile — Build e deploy (EAS)

Operamos dois apps com projetos EAS distintos. Ambos reutilizam os packages do monorepo.

## Projetos EAS

| App | Bundle ID | Canal prod | Canal staging |
|---|---|---|---|
| `mobile-pro` | `com.plenya.pro` | `production` | `preview` |
| `mobile-app` | `com.plenya.app` | `production` | `preview` |

## Primeira configuração (por app)

```bash
cd apps/mobile-pro
eas init                     # cria projectId — guardar em app.json:extra.eas.projectId
eas build:configure           # gera eas.json (já existe — revisar)
eas credentials              # gerar/importar signing keys iOS e Android
```

## Build profiles

- **development** — dev client, distribuição interna, simulador iOS habilitado
- **preview** — APK interno para QA, aponta para `api-staging`
- **production** — AAB/IPA para stores, aponta para `api` prod, auto-increment de build number

## Fluxo de release nativo (stores)

1. Bump `version` em `app.json` (semver)
2. `pnpm --filter @plenya/mobile-pro build:preview` → smoke test TestFlight/Play Internal
3. Validar checklist de segurança (ver `security.md`)
4. `pnpm --filter @plenya/mobile-pro build:production`
5. `eas submit --platform ios --profile production` → App Store Connect
6. `eas submit --platform android --profile production` → Play Console
7. Expor gradualmente:
   - iOS: TestFlight → phased release (App Store Connect)
   - Android: Play Internal → Closed → Open → Production (faixa de 10% → 50% → 100%)

## Secrets e env

- Signing keys via EAS Credentials (nunca commitar)
- `.env.*.local` nunca commitado; EAS usa `env` em `eas.json` ou `eas secret` para CI
- `secrets/google-play-service-account.json` — gitignored; obter em Play Console

## Review da Apple

- `mobile-pro` → categoria **Medical** (rigoroso — espere 48-72h)
- `mobile-app` → categoria **Health & Fitness** (mais rápido)
- ASO separado por app (screenshots, descrições, keywords por persona)

## Monitoramento

- Sentry → crash-free sessions > 99.5%
- PostHog → cold start p75 < 2s, engajamento por feature
- EAS Update dashboard → adoção por versão, rollout status
