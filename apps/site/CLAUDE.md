# apps/site — plenyasaude.com.br (institucional)

Site de marca institucional Plenya. Voz **institucional** (não pessoal do Dr. Getúlio — esse é
`apps/site-getulio`). Regras editoriais invariantes no [CLAUDE.md raiz](../../CLAUDE.md).

> 🚨 **Antes de criar/editar copy ou deck que cite a Plenya, leia a página relevante do site**
> (`app/[locale]/<topico>/page.tsx` + `messages/pt.json`) e `lib/agir-structure.ts`. O site é a
> fonte da marca; não chutar nomes de produto, pilares ou claims.

## Stack
Next 16.2 (Turbopack, porta 3002) · next-intl 4.9 (pt default, en; es legado) · MDX
(next-mdx-remote + gray-matter) · Tailwind + `@plenya/brand` · Plausible · nodemailer.

## Estrutura
```
app/[locale]/        ← home + páginas: a-plenya, continuum, metodo-agir, equipe, dr-getulio,
                       escore-plenya, blog, depoimentos, casos, healthspan, checkup-longevidade,
                       medicina-funcional-integrativa, avaliacao-renal-preventiva, privacidade...
app/api/             ← leads (→ EMR + email + webhook CRM), newsletter (Beehiiv), indexnow
content/blog/{pt,en} ← posts .mdx (frontmatter Zod: title, slug, pillar, references, cover...)
content/doctors/     ← perfis MDX da equipe
content/data/        ← score-light-config.json (sincronizado do EMR — não editar à mão)
lib/agir-structure.ts ← AGIR/ACTS canônico (fonte para escore-plenya/metodo-agir)
lib/blog.ts, authors.ts, team.ts, testimonials.ts, i18n/, plausible.ts
messages/{pt,en,es}.json
```

## Conteúdo
- **Pilares (taxonomia do blog):** alimentacao-atividade-fisica, gestao-metabolica,
  integracao-corpo-mente, ritmo-circadiano, longevidade.
- **Imagens de blog:** geradas via `scripts/blog-generator/` (gpt-image-2). Ficam em
  `public/images/blog/<slug>/`. Ver [.claude/content/images.md](../../.claude/content/images.md).
- **Escore Light:** `content/data/score-light-config.json` vem do EMR via `pnpm sync:score-light`.
- Validação pré-build: `pnpm validate-content`.

## Leads / CRM
`POST /api/leads`: normaliza telefone E.164, posta no EMR (`/api/v1/leads`), envia email
(nodemailer) e dispara webhook (`LEADS_WEBHOOK_URL`). Turnstile + rate-limit por IP.

## Deploy
Coolify 4.0 (Hetzner), Dockerfile multi-stage. Domínios: `plenyasaude.com.br` (master) /
`staging.*`. Ver `COOLIFY.md`. Em monorepo, o `manual_webhook_secret_github` deve ser idêntico
ao secret do GitHub webhook (bug recorrente — memória `coolify_site_getulio_autodeploy_bug`).
