# @plenya/site

Site público institucional da Plenya (`plenyasaude.com.br`).

Stack: **Next.js 16 + React 19 + Tailwind + next-intl (PT/EN/ES) + MDX**.

## Desenvolvimento

```bash
# Do root do monorepo
pnpm install
pnpm --filter @plenya/site dev   # http://localhost:3002
```

Via Docker Compose (similar ao apps/web):

```bash
docker compose up site
```

## Estrutura

```
app/
  [locale]/                 # rotas i18n (pt, en, es)
    page.tsx                # Home
    dr-getulio/page.tsx     # HERO Dr. Getúlio
    a-plenya/page.tsx
    equipe/page.tsx
    planos/page.tsx
    blog/page.tsx
    contato/page.tsx
    escore-plenya/page.tsx
    privacidade/page.tsx
    termos/page.tsx
  api/
    leads/route.ts          # webhook para CRM/RD Station
    newsletter/route.ts     # webhook para Beehiiv
  sitemap.ts
  robots.ts
components/
  layout/                   # Header, Footer, LocaleSwitcher
  marketing/                # Hero, AGIR pillars, cards, forms
  seo/                      # Structured data components
content/
  blog/                     # MDX posts
  doctors/                  # MDX perfis de equipe
lib/
  i18n/                     # next-intl config + navigation
  cn.ts                     # Tailwind merge helper
  plausible.ts              # Custom events
messages/
  pt.json / en.json / es.json
```

## Variáveis de ambiente

Ver `.env.example`. Em produção, os valores reais vivem no Coolify.

## SEO

- Metadata API do Next 16 em todas as páginas
- `sitemap.ts` inclui todas as rotas × locales com `hreflang`
- `robots.ts` bloqueia staging automaticamente
- `components/seo/organization-schema.tsx` injeta `MedicalOrganization` JSON-LD

## Analytics

Plausible Cloud via `lib/plausible.ts` — eventos custom listados em `PlenyaEvent`. Sem cookies, sem banner LGPD necessário.

## Branding

Tokens vêm de `@plenya/brand` (preset Tailwind + logo + marca). Nunca hardcodar cores/fontes aqui.

## Roadmap

Ver [docs/site/SITE_PLAN.md](../../docs/site/SITE_PLAN.md) para o plano faseado (MVP v1, v1.5, Fase 2, Fase 3).
