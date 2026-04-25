# Release checklist — Plenya Pro

Workflow obrigatório antes de cada **release nativo** (App Store / Play Store).
Para releases OTA, ver [ota-policy.md](./ota-policy.md).

## 1. Pré-build — qualidade e segurança

- [ ] `pnpm --filter @plenya/mobile-pro lint` limpo
- [ ] `pnpm --filter @plenya/mobile-pro tsc --noEmit` zero erros
- [ ] `pnpm --filter @plenya/mobile-pro test` (se houver testes unit) passa
- [ ] CHANGELOG.md atualizado com versão nova + categorizado (feat/fix/breaking)
- [ ] Bump `version` em `app.json` (semver) e `versionCode` (Android)
- [ ] Backend `mobileMinVersion` em `mobile_config_service.go` revisado
      (bump SOMENTE se este release tem mudança incompatível com versões antigas)

## 2. Maestro smoke (device real iOS + Android)

```bash
maestro test --include-tags smoke apps/mobile-pro/maestro/
```

- [ ] 01-login passou em iPhone real
- [ ] 01-login passou em Android real (mínimo Android 9)
- [ ] 02-patient-detail OK em ambos
- [ ] 03-anamnesis-create OK em ambos (verificar que registro aparece no web)
- [ ] 04-logout OK em ambos

## 3. Checklist de segurança manual

- [ ] **Cert pinning ativo** — testar com Charles Proxy / mitmproxy ⇒ requests
      devem falhar com SSL handshake error
- [ ] **Jailbreak detection** — testar em emulador rooted ⇒ audit log
      server-side registra `jailbroken=true`
- [ ] **Screenshot prevention** — gravar tela em `patients/[id]` ⇒ tela preta
- [ ] **Auto-logout** — deixar app idle 5min ⇒ retorna pra login
- [ ] **Biometric unlock** — Face ID/Touch ID/digital funciona
- [ ] **MMKV encryption** — `cat /data/data/com.plenya.pro/...mmkv` ⇒
      lixo binário, sem strings legíveis
- [ ] **Sentry strip** — gerar crash com paciente carregado, conferir Sentry
      Issues: nome, CPF, dados clínicos NÃO aparecem em context/extra/breadcrumbs

## 4. LGPD

- [ ] Termo de consentimento aparece no **primeiro** login do device
- [ ] Não-aceite bloqueia entrada (volta pra login)
- [ ] Aceite registra `LGPDConsentedAt` no backend
- [ ] `meSessions` no perfil mostra device atual + permite revogar outros

## 5. Privacy manifests

- [ ] `apps/mobile-pro/PrivacyInfo.xcprivacy` reflete features atuais
- [ ] `apps/mobile-pro/play-store/data-safety.md` está atualizado
- [ ] Se algum **data type novo** foi adicionado: atualizar ambos + LGPD consent

## 6. Build

```bash
# Preview interno (TestFlight + Play Internal)
pnpm --filter @plenya/mobile-pro build:preview

# Produção (após preview validado)
pnpm --filter @plenya/mobile-pro build:production
```

- [ ] Preview build subiu sem erro no EAS
- [ ] Distribuído pra time interno (mínimo 3 testers)
- [ ] 24h sem crashes/regressões reportadas

## 7. Submit

```bash
# iOS
pnpm --filter @plenya/mobile-pro exec eas submit --platform ios --profile production --latest

# Android
pnpm --filter @plenya/mobile-pro exec eas submit --platform android --profile production --latest
```

- [ ] iOS submetido — TestFlight build aparece em ~10min, full review 24-72h
- [ ] Android submetido pra **Internal track** primeiro (não production direto)
- [ ] App Store Connect: preencher "What's New in This Version"
- [ ] Play Console: preencher "Release notes" PT-BR

## 8. ASO (App Store Optimization)

Reusa de release pra release a menos que hajam screenshots novos:

- [ ] Screenshots iOS 6.7" (iPhone 15 Pro Max) — login, dashboard, paciente,
      escore, anamnese (5 telas mínimo)
- [ ] Screenshots iOS 6.5" (compatibilidade)
- [ ] Screenshots Android phone (mínimo 4)
- [ ] Screenshots tablet — `apps/mobile-pro/app.json:ios.supportsTablet=true`
      exige
- [ ] App description PT-BR atualizada (reusar do site)
- [ ] Keywords iOS (100 chars): `prontuário,EMR,médico,telemedicina,saúde`
- [ ] Privacy Policy URL: https://plenyasaude.com.br/privacidade
- [ ] Support URL: https://plenyasaude.com.br/suporte
- [ ] Marketing URL: https://plenyasaude.com.br

## 9. Pós-submit

- [ ] Tag git: `git tag -a mobile-pro-vX.Y.Z -m "..." && git push origin --tags`
- [ ] Apple Review: monitorar status no App Store Connect; se rejeitar,
      responder em até 24h
- [ ] Play: rollout faseado — Internal → Closed (10 testers externos) →
      Open (100) → Production (10% → 50% → 100% com gate de 24h cada)
- [ ] Monitorar Sentry: crash-free sessions > 99.5% nas primeiras 48h
- [ ] PostHog: cold start p75 < 2s, API error rate < 1%
- [ ] Atualizar `mobileMinVersion` no backend SOMENTE depois de adoção > 80%
      da nova versão (Play Console > Statistics) — protege quem ainda não atualizou

## 10. Comunicação

- [ ] Email pra clínicas usuárias com novidades
- [ ] In-app banner via `mobile/config.featureFlags` se feature destaque

## Rollback

Se um release nativo der ruim:
- iOS: NÃO dá pra remover do App Store; força submit de patch (≥7 dias para
  Apple aprovar). Workaround: `mobileConfig.killSwitch.enabled=true` bloqueia
  a versão problemática até patch sair.
- Android: Play Console > "Halt rollout" (instant para a faixa atual).
