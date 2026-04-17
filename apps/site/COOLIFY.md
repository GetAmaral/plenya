# Deploy do `apps/site` no Coolify

Site público da Plenya — `plenyasaude.com.br` — rodando no mesmo VPS Hetzner do EMR.

## Setup inicial no Coolify

### 1. Aplicação nova

- **Type:** Dockerfile (não Nixpacks)
- **Source:** Git (repo do monorepo Plenya)
- **Branch:**
  - `master` → produção (`plenyasaude.com.br`)
  - `staging` → staging (`staging.plenyasaude.com.br`)
- **Base directory:** `/` (raiz do monorepo — Dockerfile precisa do contexto completo)
- **Dockerfile path:** `apps/site/Dockerfile`
- **Build context:** `.` (raiz)
- **Port:** `3002`

### 2. Domínio

Produção:
- `plenyasaude.com.br` (apex)
- `www.plenyasaude.com.br` (Coolify configura redirect 301 para apex via labels Caddy)

Staging:
- `staging.plenyasaude.com.br` + Basic Auth no Coolify

### 3. Variáveis de ambiente

```
NEXT_PUBLIC_SITE_URL=https://plenyasaude.com.br
NEXT_PUBLIC_APP_URL=https://app.plenyasaude.com.br
NEXT_PUBLIC_API_URL=https://api.plenyasaude.com.br
NEXT_PUBLIC_PLAUSIBLE_DOMAIN=plenyasaude.com.br
NEXT_PUBLIC_WHATSAPP_NUMBER=55XXXXXXXXXXX
LEADS_WEBHOOK_URL=<RD Station ou outro CRM>
LEADS_WEBHOOK_SECRET=<secret bearer>
LEADS_NOTIFY_EMAIL=contato@plenyasaude.com.br
SMTP_HOST=mail.hostgator.com.br
SMTP_PORT=587
SMTP_USER=contato@plenyasaude.com.br
SMTP_PASSWORD=<vault>
NEWSLETTER_PROVIDER_API_KEY=<beehiiv>
NEWSLETTER_LIST_ID=<beehiiv publication id>
```

Staging usa as mesmas variáveis com URLs trocadas para `staging.plenyasaude.com.br` e webhook apontando para conta de teste.

### 4. Auto-deploy

- **Master push** → deploy produção (com health check em `/`)
- **Staging push** → deploy staging
- **PR preview** (opcional): habilitar para testar conteúdo antes do merge

### 5. Health check

- Path: `/`
- Esperado: `200`
- Coolify default: 30s startup grace, 10s interval

### 6. Build cache

A imagem multi-stage cacheia o layer `deps` por hash de `pnpm-lock.yaml`. Mudanças só de código (não dependências) reusam esse layer (~30s vs ~3min).

### 7. Rollback

Coolify mantém histórico de deploys. Em caso de regressão, basta clicar "Redeploy" na revisão anterior.

## Smoke test pós-deploy

```bash
curl -sS -o /dev/null -w "%{http_code}\n" https://plenyasaude.com.br
curl -sS https://plenyasaude.com.br/sitemap.xml | head -c 200
curl -sS https://plenyasaude.com.br/robots.txt
```

## Logs

Coolify expõe logs em tempo real na UI. Para acesso CLI:

```bash
ssh user@hetzner-vps
docker logs -f $(docker ps --filter name=plenya-site --format '{{.ID}}')
```

## Limites

- Plausible Cloud: ilimitado em pageviews
- Beehiiv free: até 2.500 subscribers — monitorar quando aproximar

## Observações sobre o Cloudflare

- Cloudflare na frente com proxy ON (laranja)
- SSL: Full (strict) — Coolify gera cert válido via Let's Encrypt
- Page Rule: `www.plenyasaude.com.br` → 301 para `plenyasaude.com.br`
- Cache Rule: HTML com TTL 4h, assets `/_next/static/*` com TTL 1 ano
