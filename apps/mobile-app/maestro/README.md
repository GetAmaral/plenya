# Maestro smoke flows — Plenya (paciente)

Testes E2E mínimos rodados antes de cada release nativo. **Não** rodam em CI
por padrão (precisam device/emulador real); rodam em Maestro Cloud em PRs
que tocam `apps/mobile-app/`.

## Setup local

```bash
# Instalar Maestro CLI
curl -Ls "https://get.maestro.mobile.dev" | bash

# Variáveis de teste (NÃO commitar credenciais)
cat > ~/.maestro/secrets.yaml <<EOF
MAESTRO_PATIENT_EMAIL: smoke-patient@plenyasaude.com.br
MAESTRO_PATIENT_PASSWORD: <pegar-no-1password>
EOF

# Iniciar dev client + bundler
pnpm --filter @plenya/mobile-app start
```

## Rodar

```bash
# Um flow específico
maestro test apps/mobile-app/maestro/01-login.yaml

# Todos os smoke
maestro test --include-tags smoke apps/mobile-app/maestro/
```

## Flows

| Arquivo | Tags | O que cobre |
|---|---|---|
| 01-login.yaml | smoke, auth | Login senha → biometric/LGPD opcional → tabs paciente |
| 02-check-in.yaml | smoke, check-in | Home → check-in → preencher 5 escalas → registrar |
| 03-log-workout.yaml | smoke, workout | Treino → abrir plano → iniciar → logar 1 set → finalizar |
| 04-logout.yaml | smoke, auth | Logout + retorno à tela de login |

## Antes do release

Rodar `maestro test --include-tags smoke` em **device real** iOS + Android.
Falha = release bloqueada.

## Limitações conhecidas

- Biometria interativa em device real: precisa tester humano no momento ou
  mock via Maestro `setBiometricResult` (Android) / `xcrun simctl biometric`
  (iOS simulator).
- 02 e 03 usam `tapOn:` por texto / coordenada — quando UI estabilizar, migrar
  pra `id:` com `testID` nos componentes (Scale e Input do logger).
- Push notifications não cobertas — exigem servidor de teste; vão num flow
  separado quando integrarmos APNs/FCM nativo.
