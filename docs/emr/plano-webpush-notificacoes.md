# Plano — Notificações Web Push no EMR

**Aprovado:** 2026-06-08 · **Decisão:** Web Push (e-mail descartado — "só enche a caixa").

## Objetivo

A equipe recebe um aviso do sistema (notificação nativa do navegador/PWA) quando chega
mensagem de WhatsApp (e demais eventos que já geram notificação), **mesmo com o EMR fechado** e
o celular bloqueado — sem depender do app nativo (Pro/App), que ainda não funciona.

## Princípio: reaproveitar o gatilho que já existe

O inbound de WhatsApp já cria notificação in-app via `NotifyTeamOfInboundMessage` →
`NotificationService.CreateConversationNotification` → `dispatchPush` → `PushSender.Send`.
Hoje o único `PushSender` é o **Expo** (mobile). Web Push entra como **segundo canal**, em
paralelo, sem mexer no fluxo de quem cria a notificação.

```
inbound WA → Notification (in-app, sino) ──┐
                                            ├─ dispatchPush (goroutine, fire-and-forget)
                                            └─→ PushSender:
                                                  ├─ Expo (mobile, já existe)
                                                  └─ WebPush (NOVO)  ← este plano
```

### Limitação honesta (Apple)
- **Desktop** (Chrome/Edge/Safari/Firefox): funciona aceitando "Permitir notificações" uma vez.
  Aviso chega com aba/navegador fechados. Zero instalação.
- **iPhone**: iOS só entrega Web Push se o site for **"Adicionar à Tela de Início"** (PWA
  standalone). Não é o app da loja — é um passo único. Em aba normal do Safari, a Apple bloqueia.
  → A UI detecta iOS-não-standalone e mostra a dica "adicione à tela de início" em vez do botão.

## Backend (`apps/api`)

1. **Dependência:** `github.com/SherClockHolmes/webpush-go` (criptografia aes128gcm + JWT VAPID).
2. **Config/env (VAPID):** `VAPID_PUBLIC_KEY`, `VAPID_PRIVATE_KEY`, `VAPID_SUBJECT`
   (`mailto:contato@plenyasaude.com.br`). Vazio → Web Push desligado (graceful, não quebra).
   Chaves geradas uma vez; privada vai pro cofre + Coolify (api).
3. **Model + migration `00035_web_push_subscriptions.sql`** (goose; próxima após 00034):
   `id`, `user_id` (idx), `endpoint` (unique), `p256dh`, `auth`, `device_label`, `user_agent`,
   `last_seen_at`, `created_at`, `updated_at`. Hard-delete quando o endpoint volta 404/410 (gone).
4. **`WebPushService`** (`internal/services/web_push_service.go`), implementa `PushSender`:
   - `Send(userID, PushPayload)` — carrega subs do user, envia cada uma; 404/410 → apaga sub.
   - `Subscribe(userID, sub, label, ua)` (upsert por endpoint) · `Unsubscribe(endpoint)` · `PublicKey()`.
5. **Composite sender** — `CompositePushSender{senders []PushSender}` com `Send` que faz fan-out.
   Em `main.go`: `notificationService.SetPushSender(NewCompositePushSender(pushService, webPushService))`.
6. **Handler + rotas** (`web_push_handler.go`, autenticado, qualquer staff logado):
   - `GET  /api/v1/web-push/vapid-public-key`
   - `POST /api/v1/web-push/subscribe`
   - `POST /api/v1/web-push/unsubscribe`
   - `POST /api/v1/web-push/test` — dispara push de teste pro próprio user (verificação).

## Frontend (`apps/web`)

1. **PWA:** `public/manifest.webmanifest` (nome, display standalone, theme petrol, ícones
   192/512 do `logo_infinity`) + `<link rel="manifest">`/apple-touch-icon na metadata do layout raiz.
2. **Service worker:** `public/sw.js` — evento `push` → `showNotification`; `notificationclick`
   → foca aba existente ou abre a `url` da notificação.
3. **Hook** `lib/web-push.ts` (`useWebPush()`): registra SW, lê VAPID public key, pede permissão,
   `pushManager.subscribe`, POST `/web-push/subscribe`. Estado:
   `unsupported | default | granted | denied | subscribed` + `enable()`/`disable()`/`sendTest()`.
4. **UI:** toggle "Ativar avisos neste aparelho" no dropdown do sino (top-bar). iOS-não-standalone
   → dica "Adicione à Tela de Início".
5. **Detecção iOS standalone** via `display-mode: standalone` / `navigator.standalone`.

## Fora de escopo (depois)
- Preferências por tipo (só WA × tudo) — começa com todas as notificações da equipe.
- Web Push no portal do paciente.

## Rollout
- Dev: build, registra SW em localhost, testa com `/web-push/test`.
- Env: gerar VAPID → cofre + Coolify (api).
- **Sem deploy até ordem explícita.**

## Status
- [x] Backend: dep (`SherClockHolmes/webpush-go` v1.4.0) + config `WebPushConfig` + model
  `WebPushSubscription` + migration `00035` (aplicada, goose v35) + `WebPushService` (implementa
  `PushSender`) + `CompositePushSender` (Expo+Web) + handler/rotas `/web-push/*`. Build verde.
- [x] Frontend: `manifest.webmanifest` + ícones 192/512 + `public/sw.js` + hook `lib/web-push.ts`
  + `WebPushToggle` no painel do sino + metadata/viewport no layout raiz. Typecheck verde.
- [x] VAPID dev gerado (`apps/api/.env`) — endpoint `/web-push/vapid-public-key` devolve
  `enabled:true`. Subscribe persiste + unsubscribe remove (verificado via API+DB). Assets 200.
- [x] **DEPLOYADO em prod 2026-06-08** (commit `604ae0d1`): api+web no ar; goose v35 +
  tabela `web_push_subscriptions` confirmadas; rota `/web-push/*` responde 401 (existe);
  assets PWA (`sw.js`/`manifest`/ícones) servindo 200 em app.plenyasaude.com.br. VAPID de
  **prod** (par próprio) gerado, salvo no cofre `webpush-prod.env` e setado nas envs do Coolify
  (app `plenya-api`). Dev usa par próprio em `apps/api/.env`.
- [ ] Teste real no navegador (clicar "Ativar" → push de teste chega) — fica pro usuário.

### Notas de prod (quando liberar)
- Gerar par VAPID **próprio de produção** (não reusar o de dev) e setar
  `VAPID_PUBLIC_KEY`/`VAPID_PRIVATE_KEY`/`VAPID_SUBJECT` nas envs do app `plenya-api` no Coolify.
- Web Push exige HTTPS (já temos) e, no iPhone, "Adicionar à Tela de Início" (PWA standalone).
