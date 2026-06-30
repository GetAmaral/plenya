# 01 — Stack e Infraestrutura

> Decisão de stack greenfield (jul/2026) e infra do MVP. Ver decisão de base do código em
> [00 §4](00-estrategia-spinoff.md). Data: 2026-06-29.

## 1. Princípio

A parte difícil do Gymnai é **velocidade de time pequeno + multi-tenancy + auth + pagamento +
funil consumer** — tudo melhor servido por **um stack TypeScript único da borda ao banco**. 10k+
usuários não é escala onde Go ganha de Node/TS. Logo: **TS full-stack.**

## 2. Stack alvo

### Frontend / PWA
- **Next.js 16** (App Router, RSC, Server Actions) + **React 19.2**, entregue como **PWA
  instalável** (service worker via Serwist). Push web funciona (iOS maduro). É o scan-first:
  QR → vídeo instantâneo.
- **Tailwind v4 + shadcn/ui** + tokens de marca Gymnai (paleta própria, não a gold/petrol Plenya).
- **TanStack Query** no cliente.

### Backend / domínio
- **Monolito modular**: lógica de negócio em **serviços TS puros** (framework-agnostic), exposta
  via Server Actions / route handlers do Next no MVP. Um deploy só.
- Quando chegar app nativo / integração de parceiro, extrair um **API dedicado em Hono** reusando
  os mesmos serviços. **Não microservice cedo.**

### Banco + tenancy
- **PostgreSQL** com **Drizzle ORM** (code-first, SQL-próximo, bundle minúsculo, bom cold-start).
  Prisma 7 é alternativa válida se preferir schema-first.
- **Multi-tenancy via Postgres RLS** (`organization_id`): o banco garante que cada query só vê a
  org do requisitante, mesmo se um handler esquecer o `WHERE`. Desde a migration 00001.

### Auth
- **Better Auth** — biblioteca TS-native, dona dos usuários **no seu próprio Postgres**, com
  **org plugin** (academia↔aluno↔papéis = o modelo B2B2C), passkeys, social login, magic link.
  **Não é serviço externo** — roda dentro do app. (Clerk só se quiser tudo hospedado + SSO
  enterprise — ver D4 em [02 §decisões](02-dominio-e-mvp.md).)

### Vídeo
- Ver §5 (estratégia de vídeo). Na VPS no MVP; alvo de migração ao escalar = **Bunny Stream**
  (D2 — escolhido por custo, ~metade do Cloudflare Stream em VOD de alto volume).

### Pagamento (DECIDIDO — Asaas, D3)
- **Asaas** (PSP brasileiro): Pix Automático recorrente + boleto + cartão, taxas menores, e
  **split de pagamento nativo** — que casa com o modelo de parceria (repasse a academia/personal).
  Pix Automático (recorrência BR) é o must-have. Não é "serviço extra" e sim o PSP — inevitável
  para qualquer produto que cobra. Cobre as duas pontas: recorrência do aluno (B2C) + fatura por
  unidade da academia (B2B). Stripe descartado (taxas maiores no BR, split via Connect mais
  complexo; só valeria com expansão internacional).

### Jobs / observabilidade
- **Jobs**: fila no próprio Postgres (pg-boss / Graphile Worker) + worker na mesma VPS. **Não
  precisa de Inngest no MVP** (é conforto de retry/observabilidade).
- **Observabilidade**: Sentry + PostHog (funil consumer scan→view→signup) — **free tier**, quando
  quiser medir. Pode esperar.

### Mobile (depois)
- PWA-first cobre o wedge academia. App nativo só se houver razão de loja/offline pesado — aí
  **Expo SDK 56** (RN 0.85), reaproveitando skills TS e schemas Zod. **Não no MVP.**

## 3. Shape do stack (DECIDIDO — composto self-hosted)

**Best-of-breed composto, tudo self-hosted na VPS dedicada**: Next + Postgres + Drizzle + Better
Auth + RLS no seu próprio box (Hono extraído só quando chegar app nativo/parceiro). Sem plataforma
gerenciada de DB/auth.

Por quê: coerência com a decisão de **self-host na própria VPS** e de enxugar dependência externa.
Uma plataforma batteries-included tiraria DB e auth da VPS e geraria lock-in — descartado.

Custo aceito: cabear **Better Auth + RLS na mão**. RLS em Postgres puro = políticas por tabela +
setar o tenant por transação (ex.: `SET LOCAL app.current_org = '<uuid>'` no início de cada
request autenticado, lido pelas policies). Trabalho de alguns dias, uma vez.

## 4. Infraestrutura do MVP

### Decisão: VPS dedicada própria — **não** a VPS da Plenya
A VPS da Plenya (8GB, Coolify) **já dá OOM** rebuildando os 3 apps (por isso o auto-deploy é off).
Não tem folga para um quarto app + vídeo. Pior: um pico do Gymnai (academia cheia às 19h, streams
simultâneos) poderia **derrubar o EMR clínico em produção**. Some isolamento LGPD e a história de
spinoff.

A escolha real não é "VPS vs nuvem" e sim **"mesma máquina da prod clínica" vs "máquina dedicada
barata"** → uma **VPS pequena e dedicada só do Gymnai** (ex.: Hetzner, poucos dólares/mês). Self-
host, custo baixo, isolamento desde o dia 1, sem nunca desacoplar depois.

### Footprint enxuto
```
VPS dedicada → app (Next PWA + API) + Postgres + worker + Better Auth
Cloudflare   → CDN grátis na frente da VPS (cache de PWA/estáticos e vídeo)
Vídeo        → na VPS no MVP (§5); provedor gerenciado quando escalar
PSP          → Asaas (D3; Stripe descartado)
```
Quase tudo é **biblioteca rodando no app** (Better Auth, Drizzle) ou **auto-hospedado** (Postgres,
worker). Não é sprawl de SaaS.

## 5. Estratégia de vídeo: na VPS no MVP, migrável por swap

Decisão: **vídeo na VPS no MVP**, com plano de migrar para provedor gerenciado ao escalar. Para a
migração ser um swap de config e não retrabalho, seguir 4 regras:

1. **Abstrair atrás de um `VideoService`.** Tabela `videos` com **id opaco** + backend atual
   (`local` agora, `bunny`/`cloudflare` depois) + chave/URL + ponteiro para o **master original**.
   `Equipment` referencia `video_id`, **nunca** um path. (É o padrão `GifUrl`/`GifUrlFallback` que
   o EMR já usa.) Migrar = trocar a implementação do serviço e re-apontar linhas.
2. **Nunca transcodificar sob demanda.** Transcodificar cada vídeo **uma vez**, de preferência
   **fora da VPS** (na máquina de dev), e subir só o MP4 final (H.264, `+faststart`, 720p, bitrate
   sensato). A VPS só serve bytes estáticos com HTTP Range → baixo risco de CPU/OOM.
3. **Cloudflare grátis na frente, cache agressivo.** Vídeos de aparelho são estáticos e imutáveis
   → caso perfeito de CDN. `Cache-Control: public, max-age=31536000, immutable`. A VPS serve cada
   vídeo uma vez; repetições saem do edge.
   - **Versionar a URL (M1):** Cloudflare grátis não faz purge por chave. Então a URL inclui versão
     (`/videos/<id>/v<n>.mp4` ou hash de conteúdo). Trocar um vídeo = nova URL = cache antigo
     irrelevante, sem precisar de purge. O `VideoService` gera a URL versionada.
4. **Gatilho de migração por MÉTRICA, não por contagem (M2):** banda de saída cruzando ~50% do
   plano da VPS, OU streams concorrentes acima de N, OU pressão sustentada de CPU/RAM >80% — com a
   1ª academia já estável por algumas semanas. (Removido o "segunda academia" — vem cedo demais.)
   Qualquer um dispara o swap **para Bunny Stream** (D2). Como os **masters** ficam guardados
   (+ backup), o re-encode no provedor é lossless.
5. **Vídeo é público, protegido sem URL assinada (M4):** mantém público + cache CDN (o vídeo é
   isca de propósito; URL assinada mataria o cache). Proteção: regra **Referer/hotlink** + **rate-
   limit por IP** no Cloudflare + **alerta de banda**. O CDN absorve repetições → origin barato.
   URL assinada só se houver abuso real.

## 6. O que diverge do stack Plenya (referência)

| Camada | Plenya (EMR) | Gymnai |
|--------|--------------|--------|
| Linguagem | Go | **TypeScript** |
| Backend | Fiber v2 | Next (monolito modular) → Hono depois |
| ORM | GORM | **Drizzle** |
| Banco | PG18 (self-host) | Postgres (+ **RLS** multi-tenant) |
| Auth | JWT custom | **Better Auth** (org plugin) |
| Frontend | Next (web app) | Next como **PWA** scan-first |
| Mobile | Expo SDK 52 | PWA-first; Expo 56 depois |
| Vídeo | só GIF | **vídeo** (VPS→gerenciado) |
| Pagamento | — | **Pix Automático** (Asaas) |
| Tenancy | single-tenant | **multi-tenant** (org_id + RLS) |
| Infra | VPS Plenya (Coolify) | **VPS dedicada própria** |
