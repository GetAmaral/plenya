# Mobile — Setup de desenvolvimento

Guia para desenvolver os apps `apps/mobile-pro` (profissional) e `apps/mobile-app` (paciente,
Fase 2) no monorepo.

## Pré-requisitos

- **Node >=24** (o workspace exige em `package.json:engines`)
- **pnpm 10.28.1** (exato — `packageManager` pin)
- **Watchman** (recomendado no macOS/Linux)
- **Xcode 16+** para iOS, **Android Studio Ladybug+** para Android
- **EAS CLI** — `npm i -g eas-cli` e `eas login`
- API Go rodando localmente (`docker compose up api`) em `http://localhost:3001`

## Workspace

Os dois apps consomem os packages via `workspace:*`:

- `@plenya/types` — tipos + Zod schemas gerados dos Go models
- `@plenya/brand` — tokens, fontes, logos
- `@plenya/api-client` — fetcher + `queryOptions()` + `queryKeys` TanStack v5
- `@plenya/domain` — cálculo de escore, validações CPF, formatadores
- `@plenya/ui-mobile` — primitivos RN + NativeWind

Regra: **não duplicar lógica de domínio ou rotas de API dentro de `apps/mobile-pro/`** — migra pra
um package primeiro.

## Primeiro run (local)

```bash
# Na raiz do monorepo
pnpm install

# Em outra aba: subir API
docker compose up -d api db

# Em outra aba: iniciar metro com dev client
pnpm --filter @plenya/mobile-pro start
```

Dev client precisa de build nativo ao menos uma vez:

```bash
pnpm --filter @plenya/mobile-pro build:dev:ios     # via EAS (remoto)
pnpm --filter @plenya/mobile-pro build:dev:android # via EAS (remoto)
```

Ou local (`expo run:ios` / `expo run:android`) se tiver o toolchain instalado.

## Variáveis de ambiente

Definidas em `app.json:extra` e/ou `.env.local` com prefixo `EXPO_PUBLIC_`:

| Variável | Uso |
|---|---|
| `EXPO_PUBLIC_API_BASE_URL` | URL da API (default `http://localhost:3001`) |
| `EXPO_PUBLIC_SENTRY_DSN` | DSN Sentry para crash reporting com strip de PHI |
| `EXPO_PUBLIC_POSTHOG_API_KEY` | Feature flags e analytics |
| `EXPO_PUBLIC_POSTHOG_HOST` | Host PostHog (default app.posthog.com) |
| `EXPO_PUBLIC_EAS_PROJECT_ID` | Project ID EAS para push tokens |

## Convergência com web/backend

Ver `.claude/workflows/development.md` e a seção **"Convergência com web/backend em evolução"**
do plano-mestre. Resumo:

1. Mexeu em Go model → rodar `pnpm generate` → commit dos artefatos no mesmo PR
2. Breaking change em endpoint → bumpa `mobileMinVersion` em `GET /mobile/config`
3. Features instáveis ficam no web; mobile pega quando estabilizou
4. Cadência mobile: sprint de 2 semanas com release nativo; entre releases, só OTA

## Problemas comuns

- **Metro não resolve `workspace:*`** — reinicie com `--reset-cache`. Se persistir, conferir
  `metro.config.js` — `watchFolders` deve apontar pra raiz do monorepo.
- **NativeWind classes não aplicam** — confirme que `global.css` está importado em `app/_layout.tsx`
  e que `tailwind.config.js` tem `content` cobrindo `features/` e `components/`.
- **SecureStore não funciona em simulador Android** — use device real (Keystore precisa de HW).
