# Site Plenya — Plano Mestre

Documento único consolidando todas as decisões de planejamento do site público da Plenya. Fonte de verdade para a execução. Atualizar aqui qualquer mudança de direção.

**Domínio:** `plenyasaude.com.br`
**Stack:** Next.js 16 + React 19 + Tailwind + i18n (PT/EN/ES)
**Hosting:** Hetzner VPS + Coolify (mesmo do EMR)
**DNS:** Cloudflare (proxy ON para web, OFF para email)
**Email:** HostGator (mantido apenas pelo plano de email atual)

---

## 1. Posicionamento e Negócio

### O que é a Plenya
Clínica premium de **saúde funcional integrativa**, focada em pacientes de alto poder aquisitivo (classe A, 35-55 anos), atendimento **particular only**, presencial + online.

### Produtos comerciais
1. **Consulta presencial** (clínica)
2. **Consulta online**
3. **Acompanhamento Plenya com Método AGIR** — trimestral, semestral, anual (preços no site)
4. **Consultas sob demanda**
5. **Tier único no MVP** (previsão de tiers no futuro)

### Diferenciais
- **Método AGIR** (4 pilares: Alimentação/Atividade Física, Gestão Metabólica, Integração Corpo-Mente, Ritmo Circadiano)
- **Escore Plenya** (proprietário): completo (aplicado por profissional) + intermediário + light (paciente faz online)
- **Dr. Getúlio Amaral Filho** como nome principal (nefrologista + medicina funcional integrativa)
- **Equipe multidisciplinar** (médicos, nutricionistas, psicólogos, educadores físicos)

### Hierarquia de protagonismo
1. **Dr. Getúlio** — hero, página dedicada, autoridade
2. **Outros médicos** — segundo nível, página equipe
3. **Multidisciplinar** (nutri, psico, ed. físico) — terceiro nível, mesma página

---

## 2. Branding

Ver detalhes completos em [BRAND_GUIDE.md](./BRAND_GUIDE.md). Resumo:

### Cores
| Nome | Hex | Uso |
|---|---|---|
| Dourado | `#A08456` | CTAs, acentos, hover |
| Azul Petróleo | `#0E3B4F` | Background escuro, header, footer |
| Azul Acinzentado | `#417A86` | Cards, seções alternadas |
| Verde Suavizado | `#C2C9B1` | Backgrounds claros |
| Bege/Off-White | `#E6E1D8` | Background geral claro |

### Tipografia
- **Nalieta** (serif) — headings
- **PolarbandWOO-Front** (mono) — labels, navegação
- **Inter/Geist** (sans) — corpo

### Tom
Segura, Clara, Presente, Consistente, Profundamente comprometida.
**Claim:** "Viva bem. Viva mais."
**Tagline:** "Saúde, Performance & Longevidade"

---

## 3. Arquitetura de Subdomínios

```
plenyasaude.com.br        → site público (Next.js novo)
www.plenyasaude.com.br    → 301 para apex
app.plenyasaude.com.br    → EMR (paciente + profissional, role-based)
api.plenyasaude.com.br    → API Go
staging.plenyasaude.com.br → staging do site (basic auth)
```

### DNS — Cloudflare
- Nameservers no Registro.br trocados para Cloudflare
- Plano Free
- Web records (apex, www, app, api, staging) com **proxy ON** (laranja)
- MX/SPF/DKIM/DMARC do email (HostGator) com **proxy OFF** (cinza)
- Page Rule: `www.plenyasaude.com.br` → 301 para `plenyasaude.com.br`

### Email
- Mantido no HostGator (plano já existente)
- DNS no Cloudflare aponta MX para servidor HostGator
- Sem migração no MVP

---

## 4. Navegação

### Header (desktop)
```
[LOGO PLENYA]              Home  A Plenya  Equipe  Planos  Blog  Contato        [Área do Paciente]
```

### Header (mobile)
```
[LOGO]                                                                          [☰ menu]
```

### Footer
- Coluna 1: Navegação principal
- Coluna 2: Serviços (Consulta, AGIR, Escore Plenya)
- Coluna 3: Conteúdo (Blog, Newsletter)
- Coluna 4: Contato (endereço, telefone, WhatsApp, email, redes sociais)
- Coluna 5: Legal (Privacidade, Termos, LGPD, Acesso profissionais Plenya)
- Linha final: copyright + selo "Plausible Analytics — zero cookies, 100% LGPD"

---

## 5. Sitemap

```
/                              # Home
/dr-getulio                    # Página HERO Dr. Getúlio
/a-plenya                      # Manifesto + Propósito + Método AGIR + Escore
/equipe                        # Lista médicos + multidisciplinar
/equipe/[slug]                 # Página individual (médico/profissional)
/planos                        # Consultas + AGIR
/blog                          # Index do blog
/blog/[slug]                   # Post individual
/blog/categoria/[pilar]        # Filtro por pilar AGIR
/contato                       # Formulários + WhatsApp + endereço
/escore-plenya                 # Conceito + versão light (futuro)
/privacidade                   # Política LGPD
/termos                        # Termos de uso

# i18n via subpath:
/en/...                        # inglês
/es/...                        # espanhol
```

---

## 6. Conversão (Leads → Pacientes)

### Princípio
Premium não é menos atrito — é o **atrito certo**. Audiência classe A interpreta self-checkout para serviço de R$X mil como "produto qualquer". Esperam humano antes de comprar.

### Fluxos

#### Agendar consulta
- Mostra próximos 3-5 horários disponíveis (visível, sem login)
- Clique → formulário curto (5 campos: nome, telefone, email, motivo, janela preferida)
- Confirmação por WhatsApp em até 2h
- **NÃO usar Doctoralia** — vira commodity

#### Comprar pacote AGIR
- Página /planos mostra preços transparentes
- CTA "Conversar com a equipe Plenya" → WhatsApp/call de 15min
- **Sem checkout online** (Mayo, Function Health, Forward fazem assim)

#### WhatsApp
- Bubble flutuante discreto (dourado/petróleo, não verde)
- WhatsApp Business verificado "Plenya Medicina"
- SLA visível: "Resposta em até 2 horas (8h-18h)"

#### Formulários
- **Consulta (5 campos):** nome, telefone, email, motivo (dropdown), janela preferida
- **AGIR (7 campos):** + "como nos conheceu?" + "podemos ligar nas próximas 24h?"

#### Destino dos leads
- Email para `contato@plenyasaude.com.br`
- Sync com **RD Station** (LGPD-native, ~R$300/mês) — só dados não-clínicos
- Webhook futuro para EMR via API Go (após consentimento)

---

## 7. Dr. Getúlio em Destaque

### Página dedicada `/dr-getulio`
Layout HERO completo com:
- Foto profissional grande
- Nome completo: Dr. Getúlio José Mattos do Amaral Filho
- CRM-PR: 21.876 | RQE: 16.038
- Especialidades: Nefrologista, Clínico Médico, Medicina Funcional Integrativa
- Trajetória completa:
  - Medicina UEL 2004
  - Clínica Médica Santa Casa Londrina 2006
  - Nefrologia Santa Casa Londrina 2008
  - Coordenador Residência Nefrologia Santa Casa
  - Fundador Residência Clínica Médica Santa Casa
  - Professor PUC Londrina 2013-2014
  - Pós em Medicina Funcional Integrativa ABMFI 2025-2026
  - Membro SBN + SPN
  - Responsável Técnico DaVita Intra Hospitalar Londrina
- Filosofia (3-4 parágrafos)
- Instagram @drGetulioAmaralFilho
- CTA: "Agendar consulta com Dr. Getúlio"

### Presença na home
Seção dedicada após hero principal:
- Foto + nome + 1 linha de credenciais
- 2-3 frases de filosofia
- CTA "Conhecer Dr. Getúlio" → `/dr-getulio`

### Vídeo (v1.5)
90 segundos: credenciais + filosofia, não pitch de venda. Hospedado em Bunny Stream ou Cloudflare Stream.

---

## 8. Blog

### Princípio
Restrição é premium. Menos features, mais autoridade.

### Categorias (espelham AGIR)
1. **Alimentação & Atividade Física** (pilar A)
2. **Gestão Metabólica** (pilar G)
3. **Integração Corpo & Mente** (pilar I)
4. **Ritmo Circadiano** (pilar R)
5. **Longevidade** (cross-cutting)

### Configuração
- **Comentários:** desligados. Substituir por botão "Pergunte ao Dr. Getúlio" → form
- **Frequência:** semanal
- **Autores:** Dr. Getúlio + outros médicos (você escreve no MVP)
- **Tradução:** PT-BR original, EN/ES por IA com revisão humana
- **Compartilhamento:** SÓ WhatsApp (botão único ao final)
- **Newsletter:** Beehiiv "Boletim Plenya", opt-in inline ao final do post (sem popup)

### E-E-A-T (não-negociável YMYL)
- Author box ao final: foto + CRM + RQE + bio 60-80 palavras + link página do autor
- "Revisado por Dr. X" quando coautoria
- 3-5 referências científicas linkadas (PubMed/SciELO)
- "Atualizado em [data]" prominente abaixo do título
- Tempo de leitura: opcional

### Formato dos posts
- 1.200-1.800 palavras
- TL;DR de 100 palavras no topo
- Foco único (1 conceito por post)
- Vídeo/podcast separado (link, não embed)

### Navegação
- Index: latest + filtro por pilar AGIR + 1 featured fixo (explicador AGIR)
- Sem busca no MVP (50 posts não justificam)
- Sem "mais lidos"
- Página de cada autor com bio + lista de posts

### CTA do post
Único, ao final:
> "Quer aplicar o método AGIR na sua vida? Agende uma consulta com Dr. Getúlio."

---

## 9. Internacionalização (i18n)

- **Idiomas:** PT-BR, EN, ES — todos desde o início
- **URL:** subpath (`/en/...`, `/es/...`, raiz = PT-BR)
- **Detecção:** middleware detecta `Accept-Language` do navegador no primeiro acesso → redirect → cookie de preferência depois
- **Tradução:**
  - UI/páginas institucionais: manual revisada
  - Blog: tradução por IA + revisão humana antes de publicar
- **Switcher:** dropdown discreto no header (PT | EN | ES)

---

## 10. Área do Paciente

- **Botão "Área do Paciente"** no header (discreto) → `https://app.plenyasaude.com.br`
- **Mesma URL** para paciente E profissional (sistema roteia pelo role no JWT)
- **Cadastro:** manual pela equipe após primeira consulta (não auto-serviço)
- **App mobile:** seção dedicada na home + Smart App Banner iOS + badges App Store/Play no footer
- **Auth flows** (recovery, 2FA): 100% dentro de `app.plenyasaude.com.br`

---

## 11. Analytics + Privacy + Pixels

### Stack MVP (R$35/mês)
- **Analytics:** Plausible Cloud (~R$35/mês)
- **Search Console:** Google (grátis)
- **Cookie banner:** NENHUM (Plausible-only não exige)
- **Conversões:** Plausible custom events + webhook RD Station
- **Heatmaps:** PULAR
- **A/B testing:** PULAR

### Custom events
- `cta_agendar_clicou`
- `cta_planos_clicou`
- `whatsapp_clicou`
- `form_contato_enviado`
- `form_agir_enviado`
- `newsletter_inscreveu`
- `blog_post_lido` (>60s + scroll 70%)

### Fase 2 (com ads)
- Iubenda (~R$25/mês) → cookie consent + privacy policy auto
- GTM + Meta Pixel + Google Ads tag (deferred load pós-consent)

### Privacy como posicionamento
Footer + página privacidade:
> "Sua privacidade é sagrada. Usamos analytics sem cookies e nenhum dado pessoal é compartilhado. Conformidade total com LGPD."

---

## 12. Estrutura Técnica (Monorepo)

### Adições no monorepo
```
packages/
├── ui/                      # mantém — componentes EMR
├── brand/                   # NOVO
│   ├── tokens.ts            # cores, fontes, spacing
│   ├── tailwind-preset.ts   # preset compartilhável
│   ├── logo/                # SVGs do logo
│   └── fonts/               # Nalieta + Polarband
├── types/                   # mantém

apps/
├── api/                     # mantém
├── web/                     # mantém — EMR
├── mobile/                  # mantém
└── site/                    # NOVO — site público
    ├── app/
    │   ├── [locale]/
    │   │   ├── layout.tsx
    │   │   ├── page.tsx                # Home
    │   │   ├── dr-getulio/page.tsx
    │   │   ├── a-plenya/page.tsx
    │   │   ├── equipe/
    │   │   │   ├── page.tsx
    │   │   │   └── [slug]/page.tsx
    │   │   ├── planos/page.tsx
    │   │   ├── blog/
    │   │   │   ├── page.tsx
    │   │   │   ├── [slug]/page.tsx
    │   │   │   └── categoria/[pilar]/page.tsx
    │   │   ├── contato/page.tsx
    │   │   ├── escore-plenya/page.tsx
    │   │   ├── privacidade/page.tsx
    │   │   └── termos/page.tsx
    │   ├── api/
    │   │   ├── leads/route.ts          # webhook RD Station + EMR
    │   │   └── newsletter/route.ts     # webhook Beehiiv
    │   ├── sitemap.ts
    │   ├── robots.ts
    │   └── opengraph-image.tsx
    ├── components/
    │   ├── ui/              # primitivos próprios leves
    │   ├── marketing/       # Hero, AGIRPillars, TestimonialCard
    │   ├── layout/          # Header, Footer, Nav, MobileMenu
    │   └── blog/            # MDXContent, AuthorBox
    ├── content/
    │   ├── blog/            # .mdx
    │   ├── doctors/         # .mdx (perfis equipe)
    │   └── plans/           # .mdx ou TS
    ├── lib/
    │   ├── i18n/
    │   ├── blog.ts          # parsing MDX
    │   ├── plausible.ts
    │   └── seo.ts
    ├── messages/            # pt-BR.json, en.json, es.json
    ├── public/
    │   ├── images/
    │   └── og/
    ├── middleware.ts        # i18n detection
    ├── package.json
    ├── next.config.ts
    ├── tailwind.config.ts   # estende packages/brand/tailwind-preset
    └── tsconfig.json
```

### Stack mínima `apps/site/package.json`
```
next 16.2.x
react 19.2.x
tailwindcss 3.4.x
next-mdx-remote
gray-matter
zod
react-hook-form
next-intl
plausible-tracker
lucide-react
```

**NÃO incluir:** TanStack Query, axios, dnd-kit, recharts, tiptap, todos Radix do EMR. Site não precisa.

### Performance budget
- Lighthouse target: 95+ em todas as métricas
- Bundle JS inicial: <100KB gzipped
- Imagens: `next/image` AVIF/WebP, blur placeholder
- Fontes: self-hosted com `font-display: swap` + preload críticas

---

## 13. SEO

### Day-1 obrigatório
- Google Search Console configurado
- Sitemap XML dinâmico (todos os posts/médicos/idiomas)
- robots.txt (permite tudo prod, bloqueia staging)
- Metadata API Next 16 em todas as páginas
- JSON-LD: `MedicalOrganization`, `Physician` (Dr. Getúlio + equipe), `Article` (blog), `MedicalBusiness`
- Open Graph images dinâmicas por rota (`opengraph-image.tsx`)
- Canonical URLs corretos com hreflang para PT/EN/ES

### Local SEO (v1.5)
- Google Meu Negócio otimizado
- Schema `LocalBusiness` com endereço + horários
- Reviews integration quando houver

### Palavras-chave alvo (a refinar)
- "medicina funcional integrativa Londrina"
- "longevidade premium"
- "saúde preventiva personalizada"
- "Dr. Getúlio Amaral Filho"
- "acompanhamento médico integrado"
- "método AGIR"
- "escore de saúde"

---

## 14. MVP — Fases e Timeline

### Fase 0 — Infra (3 dias)
- Cloudflare DNS migrado
- `packages/brand` criado
- `apps/site/` scaffold no monorepo
- Coolify staging com basic auth
- Plausible + Search Console configurados

### MVP v1 — Go Live (semanas 1-3)
**Páginas:** Home, Dr. Getúlio, A Plenya, Equipe, Planos, Contato, Blog index + 1-2 posts, Privacidade, Termos
**Features:** i18n PT-BR completo, WhatsApp bubble, formulários funcionais, sitemap, structured data, OG images, Plausible events
**Não inclui:** Blog cheio, Newsletter, área paciente conectada, pixels Meta/Google, cookie banner, A/B, heatmaps, vídeo Dr. Getúlio (se não pronto)

### MVP v1.5 — Autoridade (semanas 4-6)
- 8-10 posts publicados
- Beehiiv + opt-in newsletter
- Vídeo 90s Dr. Getúlio
- 3-5 depoimentos com consentimento
- Página `/escore-plenya` com versão light
- EN e ES revisados
- Google Meu Negócio otimizado

### Fase 2 — Aquisição paga (mês 2-3)
- Iubenda cookie consent
- Meta Pixel + Google Ads via GTM
- Primeiras campanhas
- Landing pages para campanhas
- Schema reviews

### Fase 3 — Otimização (mês 4+)
- Microsoft Clarity (só landing)
- A/B testing
- Calculadoras médicas (lead magnets)
- SSO área paciente integrada com EMR

### Princípio
"Go live embarrassingly simple". Melhor 6 páginas sólidas em 3 semanas que 20 incompletas em 3 meses. SEO conta do dia 1.

---

## 15. Custos Mensais Recurring

### MVP v1
| Item | Custo |
|---|---|
| Cloudflare Free | R$0 |
| Plausible Cloud | R$35 |
| Hetzner VPS (já pago) | — |
| HostGator email (já pago) | — |
| **Total novo** | **R$35/mês** |

### MVP v1.5 (+ newsletter)
| Item | Custo |
|---|---|
| Beehiiv Free (até 2.500 subs) | R$0 |
| **Total** | **R$35/mês** |

### Fase 2 (+ ads)
| Item | Custo |
|---|---|
| Iubenda | R$25 |
| RD Station (se ainda não tiver) | R$300 |
| **Total** | **~R$360/mês** |

---

## 16. Decisões em Aberto

Itens a definir antes/durante execução:

- [ ] Foto profissional do Dr. Getúlio (alta resolução, fundo neutro)
- [ ] Lista completa da equipe (nomes, CRM/CRN/CREF, fotos, bios)
- [ ] Endereço da clínica em Londrina (para footer + Google Maps)
- [ ] Horários de atendimento (presencial + online)
- [ ] Telefone fixo + WhatsApp Business número
- [ ] Preços do AGIR (trimestral, semestral, anual)
- [ ] Preço "a partir de" para consultas (Dr. Getúlio vs outros)
- [ ] 3-5 textos prontos para primeiros posts (você mencionou ter material inicial)
- [ ] Filosofia/valores adicionais do Dr. Getúlio para página dedicada
- [ ] Confirmação se RD Station entra no MVP v1 ou v2
- [ ] Política de privacidade redigida (LGPD) — usar Iubenda ou redigir custom?
- [ ] Termos de uso redigidos
- [ ] Configurar conta Cloudflare + iniciar migração DNS

---

## Changelog

- **2026-04-16** — Documento inicial consolidando 11 perguntas de planejamento
