# Plenya (paciente)

App React Native (Expo SDK 52, New Architecture) para pacientes Plenya.
Bundle: `com.plenya.app` · Scheme: `plenya-app://` · Universal link: `https://app.plenyasaude.com.br`.

Espelha o portal web (`minha.plenyasaude.com.br`) com adições nativas:
log de treino, check-in diário de bem-estar, push configurável.

## Início rápido

```bash
pnpm install
pnpm --filter @plenya/mobile-app start
```

## Variáveis de ambiente

Definidas em `app.json` (extra) ou via `.env` com prefixo `EXPO_PUBLIC_`:

- `EXPO_PUBLIC_API_BASE_URL` — ex: `http://localhost:3001`
- `EXPO_PUBLIC_SENTRY_DSN`
- `EXPO_PUBLIC_POSTHOG_API_KEY`
- `EXPO_PUBLIC_EAS_PROJECT_ID`

## Diferenças vs `mobile-pro`

| Aspecto | mobile-pro | mobile-app (paciente) |
|---|---|---|
| 2FA | obrigatório | opcional |
| Auto-logout | 5min | 30min |
| Login | senha + 2FA | senha OR magic-link |
| Screen capture | bloqueado em prontuário | bloqueado em /exames e /perfil |
| Permissões | câmera + biometria + foto | biometria + foto |

70%+ de código vem de `packages/{api-client, brand, ui-mobile, domain}` —
mesma stack do mobile-pro, telas e fluxos diferentes.

## Estrutura

```
app/
├── _layout.tsx               providers (QueryClient, MMKV persist, theme)
├── (auth)/
│   ├── login.tsx             email + senha
│   ├── magic-link.tsx        request magic link
│   └── consume-link.tsx      deep link → JWT
└── (tabs)/                   placeholder Sprint 2
features/                     domain hooks (auth, workouts, check-ins, etc)
lib/
├── security/                 cert pin, jailbreak, screenCapture, autoLogout(30min)
├── storage/                  secure.ts (Keychain), mmkv.ts
├── observability/            sentry, posthog
└── push/                     registerDevice, handler (deep link router)
```

## Comandos

```bash
pnpm --filter @plenya/mobile-app start            # dev client
pnpm --filter @plenya/mobile-app build:dev:ios    # EAS dev iOS
pnpm --filter @plenya/mobile-app build:preview    # EAS preview ambos
```
