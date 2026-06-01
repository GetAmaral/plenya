# apps/site-getulio — drgetulioamaralfilho.com.br (pessoal)

Site editorial/pessoal do Dr. Getúlio Amaral Filho. Voz em **primeira pessoa** (não
institucional Plenya). Regras editoriais no [CLAUDE.md raiz](../../CLAUDE.md).

## Stack
Next 16.2 (Turbopack, porta 3003) · next-intl 4.9 (pt/en, todas as rotas com slug traduzido) ·
MDX · Tailwind v4 (CSS-first `@theme`, paleta custom, **não** usa `@plenya/brand`) · nodemailer.

## Estrutura
```
app/[locale]/   ← home, o-medico (/the-physician), artigos (/articles), livros (/books) +
                  [slug]/excertos, palestras (/lectures), ensino (/teaching),
                  onde-atendo (/where-i-practice), contato (/contact)
app/api/        ← contact (nodemailer, roteia por motivo), indexnow
app/artigos.xml, app/llms.txt, app/manifest.ts, sitemap.ts, robots.ts
content/blog/{pt,en} ← mirror do blog da Plenya (lido de apps/site no build; assets reescritos
                       para o CDN plenyasaude.com.br). Voz 1ª pessoa.
content/livros/      ← .mdx (frontmatter Zod: isbn, editions, amazonUrl/amazonUrlEn, excerpts...)
content/lectures/    ← .mdx (audience[], duration, format, slugEn, campos *En)
lib/blog.ts, books.ts, lectures.ts, i18n/config.ts (pathnames traduzidos)
```

## Pontos de atenção
- **Sem CRM/newsletter** — só email via `/api/contact` (roteado por motivo: consulta-plenya,
  consulta-nefroclinica, palestra, imprensa, outro).
- **Imagens** vêm do CDN da Plenya (`plenyasaude.com.br/images/...`), não duplicadas.
- **Redirects** extensos em `next.config.ts` (renomeações de slug pt/en). Validação: `pnpm validate-content`.
- Livro principal: "Antes — Como viver melhor depois dos 40". ISBNs: memória `livro_antes_isbns`.
- Handles/identidade do Dr.: memória `getulio_canonical_handles`. Email/inbox: `email_dr_getulio_acesso`.

## Deploy
Coolify 4.0 (Hetzner). Mesmo cuidado de webhook do monorepo (ver `apps/site/CLAUDE.md`).
Aliases de email Stalwart + Resend SMTP: memória `plenya_vps_site_getulio`.
