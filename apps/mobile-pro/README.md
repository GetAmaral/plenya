# Plenya Pro

App React Native (Expo SDK 52, New Architecture) para profissionais de saúde Plenya.

Este app consome a API Go do Plenya e espelha as funcionalidades do web `apps/web`. Ver
`.claude/mobile/setup.md` na raiz para fluxo de desenvolvimento, e o plano-mestre em
`/home/user/.claude/plans/vivid-shimmying-glacier.md`.

## Início rápido

```bash
# Requer Node >=24 e pnpm 10.28.1
pnpm install

# Dev client (precisa de build nativo em device real; simulador para iOS OK)
pnpm --filter @plenya/mobile-pro start
```

## Variáveis de ambiente

Definidas em `app.json` (extra) ou via `.env` com prefixo `EXPO_PUBLIC_`:

- `EXPO_PUBLIC_API_BASE_URL` — ex: `http://localhost:3001`
- `EXPO_PUBLIC_SENTRY_DSN`
- `EXPO_PUBLIC_POSTHOG_API_KEY`
- `EXPO_PUBLIC_EAS_PROJECT_ID`

## Segurança (camada `lib/security/`)

| Módulo | Responsabilidade |
|---|---|
| `biometric.ts` | Face ID/fingerprint via expo-local-authentication |
| `screenCapture.ts` | Hook `useScreenCaptureProtection` — bloquear screenshot em prontuário |
| `certPinning.ts` | SSL pinning com backup pin (configurado em cold start) |
| `jailbreak.ts` | Detecção jail/root + sinal de confiança |
| `autoLogout.ts` | Logout por inatividade (5min) |

## Armazenamento

- `lib/storage/secure.ts` — SecureStore (Keychain/Keystore) para tokens JWT
- `lib/storage/mmkv.ts` — MMKV criptografado (chave derivada do SecureStore) para cache
