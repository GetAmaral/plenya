# Seed para `/home/user/gymnai`

Este arquivo prepara o contexto de Claude Code do novo projeto. Quando criar `/home/user/gymnai`:

1. Copie o bloco **CLAUDE.md** abaixo para `/home/user/gymnai/CLAUDE.md`.
2. Semeie a memória conforme **"O que semear de memória"**.
3. Não copie skills/agents da Plenya (deck, responder-insta, linkedin etc. são irrelevantes).

> Por que seed e não herança: a base de conhecimento da Plenya é ~90% específica dela (EMR, deploy,
> social, clínico). Transferível para o Gymnai é só um subconjunto pequeno — convenções de dev e o
> mapa do domínio de treino. Seed único = sinal alto, sem ruído. Ver [00 §3](00-estrategia-spinoff.md).

---

## CLAUDE.md (copiar para `/home/user/gymnai/CLAUDE.md`)

```markdown
# Gymnai

PWA scan-first de academia: o usuário lê um QR code no aparelho e vê o vídeo explicativo; se logado
com assinatura ativa, vê o treino dele para aquele aparelho no dia. B2B2C — **academias e
condomínios** são organizações/tenants (lugares com aparelhos); **personal trainers** são o ator
`Professional` (pessoa). Spinoff do módulo de treino do EMR Plenya, mas **produto totalmente
separado**.

> Planejamento detalhado nasceu em `plenya/docs/gymnai/` (estratégia, stack, domínio). Replicar o
> essencial aqui conforme o projeto andar.

## Marca

Identidade **fechada** (assets em `plenya/docs/gymnai/identidade/`: SVGs de símbolo/logo/wordmark,
`gymnai-tokens.css`, e `icons/` com favicon + app-icons PWA). Tom **premium / performance** — luxo
atlético, **ouro sobre petróleo escuro** (dark-first). Símbolo: figura atlética em ouro
(vetorizado). Assinatura: **"INTELLIGENCE IN MOTION"**. Tipografia: **Trajan Pro** (paga) → na web
**Cinzel** (Google Fonts, OFL, livre) para títulos + **Montserrat** para o corpo. Paleta oficial:
OURO `#B4894D`, PETROL `#022837`, AZUL CLARO `#3C6971`, VERDE `#ADB29B`, CREME `#E1D9CC` (gradiente
do ouro até `#D7B975`). Tema do PWA = petróleo + creme + ouro de realce. Tokens próprios em
`gymnai-tokens.css`, **não** reusar os da Plenya.

## Regras invariantes

1. **Multi-tenant nativo.** Toda tabela tem `organization_id` desde a migration 00001. Isolamento
   garantido por **Postgres RLS**, não só por código. Nunca adiar tenancy para "V2".
2. **Stack TypeScript full-stack.** Next.js 16 (App Router, PWA) + Postgres + Drizzle + Better Auth
   (org plugin) + RLS. Sem Go.
3. **Dev ≡ Prod paridade.** Banco direto via psql/scripts em dev; nunca manipular dado manual via
   API HTTP em dev. Migrations à mão (não auto-derivar schema dos models).
4. **Catálogo de treino/vídeo é a fonte de verdade do Gymnai.** Semeado do EMR no início; daqui pra
   frente o Gymnai é o dono. Nada de prontuário aqui — dado clínico ligado a paciente vive no EMR.
5. **Identidade separada da Plenya.** Aluno do Gymnai não é Patient do EMR. Integração só com
   consentimento explícito, via API, nunca por JOIN.
6. **Vídeo atrás de `VideoService` trocável.** `Equipment` referencia `video_id`, nunca um path.
   Backend de vídeo (`local`|`bunny`|`cloudflare`) é detalhe de implementação. Nunca transcodificar
   sob demanda; sempre pré-assar o MP4.
7. **Nunca deploy/commit/push sem ordem explícita.** (Herdado da cultura Plenya.)

## Domínio (núcleo)

```
Catálogo GLOBAL (sem org_id; lido até no scan anônimo):
  ExerciseArchetype → Video        (vídeo por ARQUÉTIPO; importado da VPS treinador)
  EquipmentModel (genérico OU marca/modelo) → arquétipo(s)   (mapeado 1x, reusado entre orgs)
Tenant (RLS por org):
  Organization (gym|condo; partnership: none|facility) → Location → Equipment (qr_token opaco → EquipmentModel)
Atores:
  Member       (aluno; dono dos próprios dados)
  Professional (personal/professor; N orgs + N alunos; base do acesso delegado)
Assinatura & patrocínio:
  Subscription (tier: basico|plus · source: personal|sponsored)
  Sponsorship  (org OU professor paga p/ N alunos; conciliação de roster planilha→API)
  Entitlement  (acesso efetivo; deriva da Subscription ativa)
  WorkoutAssignment (IA gera + professor pode montar manual OU revisar/ajustar)
```

Scan = **isca pública**: `/q/<token>` resolve org/location/equipment → arquétipo → **vídeo**
(sempre, sem auth, só catálogo global). Se houver **sessão válida**, sobrepõe o **treino do
usuário** (autenticado, RLS-scoped). **Escanear nunca concede assinatura** — sem geofence, sem TTL.
Acesso = Subscription ativa (pessoal OU patrocinada por roster). `org`/`location` no scan servem só
para co-branding e atribuição.

## Stack

- Frontend/PWA: Next.js 16 + React 19.2 + Tailwind v4 + shadcn + TanStack Query + Serwist (SW)
- Backend: monolito modular (serviços TS puros via Server Actions/route handlers) → Hono depois
- Banco: Postgres + Drizzle + RLS
- Auth: Better Auth (org plugin, passkeys, social, magic link)
- Vídeo: VPS no MVP (pré-assado + Cloudflare cache) → Bunny Stream ao escalar
- Pagamento: Pix Automático via Asaas (split nativo p/ parcerias)
- Jobs: fila no Postgres (pg-boss/Graphile) + worker
- Infra: VPS dedicada própria (NÃO a VPS da Plenya) + Cloudflare grátis na frente

## Comandos (preencher quando o scaffold existir)

# (pnpm dev, drizzle migrate, etc.)
```

---

## O que semear de memória

Copiar para o `memory/` do Gymnai (reescrito no contexto Gymnai), e **só** isto:

- **Convenções de dev/prod** (banco direto em dev, migrations à mão, paridade dev≡prod) — adaptar
  da cultura Plenya.
- **Mapa do domínio de treino** (estrutura WorkoutPlan/Session/Exercise, logs set-a-set/RPE,
  periodização) como referência do que reescrever em TS — origem: exploração do módulo de treino do
  EMR registrada em `plenya/docs/gymnai/`.
- **As decisões fechadas** (greenfield TS, repo separado, fronteira prontuário-vs-catálogo, vídeo na
  VPS com gatilho de migração, wedge academia) — uma memória `project` apontando para
  `plenya/docs/gymnai/`.

**Não** semear: deploy por-app Plenya, regras editoriais/voz de marca Plenya, social/IG/LinkedIn,
decks, escore clínico, VPS/Coolify da Plenya, qualquer coisa de paciente/prontuário.

---

## Lembrete lado-Plenya

Ao iniciar o Gymnai, **congelar o módulo de treino do EMR** para correções apenas (ver
[00 §6](00-estrategia-spinoff.md)) e considerar gravar uma memória Plenya registrando o freeze e o
ponteiro para `docs/gymnai/`.
