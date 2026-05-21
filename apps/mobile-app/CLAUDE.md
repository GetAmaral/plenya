# apps/mobile-app — App do Paciente (Plenya)

App nativo para pacientes — espelho do portal do paciente. Estado: **inicial** (v0.1.0,
estrutura/sprints iniciais; menos construído que `mobile-pro`). Regras de Ouro no
[CLAUDE.md raiz](../../CLAUDE.md). Plano de execução: memória `mobile_app_patient_plan`.

## Identidade
- Nome: **Plenya** · slug `plenya-app` · scheme `plenya-app://`
- Bundle: `com.plenya.app` (iOS + Android)
- Universal link: `https://app.plenyasaude.com.br` (intentFilter Android para deep link)

## Stack
Igual ao `mobile-pro`: Expo SDK 52 · RN 0.76.5 · Expo Router · TanStack Query v5 · Zustand 5 ·
NativeWind 4. Mesmos packages compartilhados.

## Estrutura (alvo)
```
app/(auth)/   ← login (suporta magic-link), 2FA opcional
app/(tabs)/   ← home, my-score, my-records, my-workouts, messages
app/appointments/
```

## Diferenças vs mobile-pro
- 2FA **opcional**; auto-logout **30min**; login por **magic-link** suportado.
- Sem câmera; screen capture bloqueado só em /exams e /profile.
- Check-in diário de bem-estar (`HealthCheckIn` no backend) + push configurável.

## Docs
Mesmos arquivos em [.claude/mobile/](../../.claude/mobile/). Decisões de produto (WorkoutSession,
check-in MVP, deep link custom scheme): memória `mobile_app_patient_plan`.
