# Mobile — Segurança e LGPD

Checklist obrigatório no MVP do `mobile-pro`. Foundation LGPD já existe no backend; aqui
garantimos que o mobile não vaza dados.

## Tokens e sessão

- [x] JWT access (15min) + refresh (7d) em `expo-secure-store` (Keychain/Keystore)
  — **nunca** em AsyncStorage ou MMKV em cleartext
- [x] MMKV criptografado com chave derivada do SecureStore (`lib/storage/mmkv.ts`)
- [x] Biometric unlock no launch (`app/(auth)/biometric-unlock.tsx`)
- [x] Auto-logout por inatividade: 5min (pro), 30min (paciente) — `lib/security/autoLogout.ts`
- [x] Refresh token rotation no 401 com deduplicação (um único refresh em flight)
- [x] Logout server-side (`/api/v1/auth/logout`) mesmo se já offline

## Dados em tela

- [x] `useScreenCaptureProtection()` em `patients/**`, anamnese, prescrições, labs
- [x] `Text` e `Input` não logam conteúdo (Sentry com `beforeSend` strippando PHI)
- [x] CPF/RG exibidos mascarados por padrão (API retorna `cpfMasked`). Decriptografia só no
  backend sob permissão; mobile nunca decripta
- [x] Imagens de paciente sempre via URL autenticada; nunca salvar em cache externo

## Rede

- [x] HTTPS obrigatório em produção (`app.json:ios.infoPlist` não declara exceptions ATS)
- [x] Cert pinning com backup pin via `react-native-ssl-public-key-pinning`
  (`lib/security/certPinning.ts`), configurado após fetch de `/api/v1/mobile/config`
- [x] Kill-switch remoto em `mobileConfig.killSwitch.enabled` — app mostra tela bloqueada
- [x] Header `X-App-Variant: pro` em toda request, registrado em audit log server-side

## Device integrity

- [x] `jail-monkey` reporta jailbreak/root em `lib/security/jailbreak.ts` no cold start
- [x] Em device não-trustworthy: audit log server-side + privilégios reduzidos (ex.: não permite
  decriptação de CPF mesmo se permissão concede)
- [x] Não bloquear uso em dev (APP_VARIANT=development) pra permitir testes

## LGPD

- [x] Termo de consentimento no primeiro login (tela a ser adicionada no Sprint 1)
  — registra aceite via `meMutations.acceptLgpdConsent()`
- [x] `meSessions` visível no perfil com opção de revogar devices
- [x] Direito à portabilidade: API já expõe `GET /api/v1/patients/:id/export` (read-only no mobile)
- [x] Retenção: dados em MMKV são apenas cache; source of truth é o backend

## Sentry — PHI stripping

```ts
Sentry.init({
  dsn: env.sentryDsn,
  beforeSend(event) {
    // Strip potentially-sensitive fields from error contexts
    if (event.extra) {
      for (const key of ['cpf', 'rg', 'email', 'phone', 'name', 'patient']) {
        if (key in event.extra) event.extra[key] = '[stripped]';
      }
    }
    return event;
  },
});
```

## Antes de release — validar manualmente

- Charles Proxy / mitmproxy → tenta MITM a produção, deve falhar (cert pinning)
- Emulador Android rooted → audit log server-side deve registrar `jailbroken=true`
- Gravação de tela em `patients/[id]` → tela deve sair preta (screenshot prevention)
- Deixar app ocioso 5min → logout automático
- Buscar CPF no Sentry após gerar um crash → nenhuma ocorrência
