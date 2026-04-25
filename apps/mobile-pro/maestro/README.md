# Maestro smoke flows — Plenya Pro

Testes E2E mínimos rodados antes de cada release nativo. **Não** rodam em CI
por padrão (precisam device/emulador real); rodam em Maestro Cloud em PRs
que tocam `apps/mobile-pro/`.

## Setup local

```bash
# Instalar Maestro CLI
curl -Ls "https://get.maestro.mobile.dev" | bash

# Variáveis de teste (NÃO commitar credenciais)
cat > ~/.maestro/secrets.yaml <<EOF
MAESTRO_TEST_EMAIL: smoke@plenyasaude.com.br
MAESTRO_TEST_PASSWORD: <pegar-no-1password>
MAESTRO_TEST_2FA_CODE: <gerar-com-oathtool>
EOF

# Iniciar dev client + bundler
pnpm --filter @plenya/mobile-pro start
```

## Rodar

```bash
# Um flow específico
maestro test apps/mobile-pro/maestro/01-login.yaml

# Todos os smoke
maestro test --include-tags smoke apps/mobile-pro/maestro/
```

## Flows

| Arquivo | Tags | O que cobre |
|---|---|---|
| 01-login.yaml | smoke, auth | Login + 2FA + biometric + LGPD consent → tabs |
| 02-patient-detail.yaml | smoke, patient | Lista pacientes + abre detail + verifica todas as sub-rotas |
| 03-anamnesis-create.yaml | smoke, anamnesis, write | Cria anamnese + valida toast e lista |
| 04-logout.yaml | smoke, auth | Logout + retorno à tela de login |

## Antes do release

Rodar `maestro test --include-tags smoke` em **device real** iOS + Android.
Falha = release bloqueada.

## Limitações conhecidas

- Biometria interativa em device real: precisa tester humano no momento ou
  mock via Maestro `setBiometricResult` (Android) / `xcrun simctl biometric`
  (iOS simulator).
- Push notifications não cobertas — exigem servidor de teste; vão num flow
  separado quando integrarmos APNs/FCM nativo.
