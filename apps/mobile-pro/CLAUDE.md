# apps/mobile-pro — App Profissional (Plenya Pro)

App nativo para profissionais (médicos, enfermagem, nutrição, psicologia, ed. física).
Estado: **MVP em curso** (v0.1.0) — auth e navegação prontos, telas de dados em construção.
Regras de Ouro no [CLAUDE.md raiz](../../CLAUDE.md).

## Identidade
- Nome: **Plenya Pro** · slug `plenya-pro` · scheme `plenyapro://`
- Bundle: `com.plenya.pro` (iOS + Android)

## Stack
Expo SDK 52 · React Native 0.76.5 · React 19 · Expo Router (file-based) · TanStack Query v5 ·
Zustand 5 · NativeWind 4.

## Packages compartilhados (não há reuso de componente web↔mobile — só tokens/lógica)
`@plenya/types` · `@plenya/api-client` · `@plenya/brand` (tokens) · `@plenya/domain` ·
`@plenya/ui-mobile` (primitivos RN + NativeWind).

## Estrutura
```
app/(auth)/    ← login, 2FA
app/(tabs)/    ← patients, agenda, leads, conversations, notifications, training, search, profile
lib/security/  ← boot, biometric, screenCapture, certPinning, jailbreak, autoLogout
features/      ← hooks de auth, seleção de paciente, uploads
```

## Segurança (mobile-pro é o perfil mais restrito)
2FA obrigatório · auto-logout 5min · SecureStore (tokens) · MMKV criptografado (cache) ·
prevenção de screen capture · cert pinning · detecção de jailbreak · Sentry com PHI stripping.
Checklist completo: [.claude/mobile/security.md](../../.claude/mobile/security.md).

## Docs
[setup](../../.claude/mobile/setup.md) · [deploy](../../.claude/mobile/deploy.md) (EAS) ·
[ota-policy](../../.claude/mobile/ota-policy.md) · [release-checklist](../../.claude/mobile/release-checklist.md).
Plano-mestre da fase mobile: `/home/user/.claude/plans/vivid-shimmying-glacier.md`.
