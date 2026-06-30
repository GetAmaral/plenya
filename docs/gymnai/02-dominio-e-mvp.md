# 02 — Domínio e MVP

> Modelo de domínio, resolução de acesso, escopo do MVP e decisões em aberto. Data: 2026-06-29.
> Wedge inicial: **academia**.

## 1. Modelo de domínio

O Gymnai introduz entidades que o EMR não tem (`Equipment`, `Organization`, `Entitlement`). Há
duas camadas: o **catálogo GLOBAL** (arquétipos + vídeos, compartilhado entre todas as orgs) e a
camada **tenant-scoped** (org-específica). Multi-tenancy via `organization_id` + RLS aplica-se à
camada tenant-scoped — **não** ao catálogo global (ver I3 em [04](04-questoes-abertas.md) p/ o
desenho exato de RLS e o caminho do scan anônimo).

```
── CATÁLOGO GLOBAL (não por org; lido inclusive no scan anônimo) ──
ExerciseArchetype (ex.: "leg press", "supino")   ← arquétipo de movimento
  └─ video_id → Video
Video (POR ARQUÉTIPO; IA ancorada na biomecânica; importado da VPS treinador)
  └─ id opaco + backend (local|bunny) + master original
EquipmentModel (genérico OU específico: "Leg Press"  ou  "Technogym Leg Press 700")
  ├─ name, brand?, model?      (marca/modelo OPCIONAIS — genérico = só o tipo)
  └─ → ExerciseArchetype(s)    (mapeado UMA vez, reusado em todas as academias)

── ATORES & RELAÇÕES (escopados — RLS por contexto, ver I3) ──
Organization (parceiro com lugar: gym | condo)            ← place tenant
  ├─ name, logo_url            (co-branding na página de scan)
  ├─ partnership: none | facility   (dirige o CTA da página de scan)
  └─ Location → Equipment (instância física; → EquipmentModel; qr_token opaco; herda o vídeo do modelo)
Professional (professor / personal trainer)               ← pessoa (hub de relações)
  ├─ ProfessionalOrganization  (N:M — em quais orgs atua)
  └─ ProfessionalClient        (N — quais alunos treina; base do acesso delegado)
Member (aluno; consumer, dono dos próprios dados; signup pelo funil — scan NÃO cria Member)
  └─ WorkoutAssignment   (treino; IA gera + professor pode montar manual OU revisar/ajustar a da IA)

── ASSINATURA & PATROCÍNIO ──
Subscription  → ativa (tier: basico|plus · source: personal|sponsored · asaas_id)
Sponsorship   → patrocinador paga p/ N alunos (sponsor_type: organization|professional ·
                sponsor_id); conciliação roster planilha→API. Generaliza o antigo "facility".
Entitlement   → acesso efetivo, derivado da Subscription ativa
```

> `personal_trainer` **deixou de ser subtipo de Organization** — virou o ator `Professional`.
> Org = lugar com aparelhos (gym/condo); Professional = pessoa. O "facility" (org paga) e o
> patrocínio do professor são o **mesmo mecanismo** (`Sponsorship`).

**Vídeo é por arquétipo, não por marca/modelo** (I1): um vídeo de "leg press" serve qualquer leg
press. Onboarding de academia = **mapear cada aparelho ao arquétipo** (rápido), não gravar vídeo.

Catálogo (arquétipos + vídeos): **fonte de verdade do Gymnai** (ver [00 §5](00-estrategia-spinoff.md)).
Os **exercícios/GIFs** vêm do EMR; os **vídeos por arquétipo** vêm da **VPS treinador** (ver §6).

## 2. A tag é uma isca pública (não concede acesso)

**O QR nunca dá assinatura.** Ele resolve `{organization, location, equipment}` → arquétipo →
**vídeo educativo** (uso correto, variações, erros) + uma **camada contextual** que depende de quem
está escaneando (ver §3). Assinatura vem de **pagar** (pessoal) ou de **estar no roster pago** da
academia (facility) — ver [modelo de negócio em 00 §8](00-estrategia-spinoff.md). O `org`/`location`
na resolução servem para **atribuição/analytics** (de qual academia veio o scan), não para liberar
acesso. Se o usuário **já estiver logado com assinatura válida**, o scan apenas **reconhece** isso
e troca o CTA pelo treino — não concede nada (logo, sem geofence/anti-clonagem).

### Segurança do tag (D6 — atualizado pelo modelo de negócio)

- **Token opaco DB-backed.** O QR codifica um token aleatório (nanoid ~128-bit) numa URL
  `gymnai.app/q/<token>`; o servidor faz lookup em `equipment` → arquétipo → vídeo. **Revogável**
  (rotaciona token / flag) sem reimprimir as outras placas. Não usar token assinado stateless
  (JWT/HMAC): numa placa física que vive anos, revogação > evitar o lookup.
- **Sem geofence, sem anti-clonagem.** Como escanear **nunca** concede assinatura (só vídeo + ad),
  clonar/fotografar o QR não causa dano. Toda a complexidade de geofence + TTL de sessão do plano
  antigo **foi removida** — não existe "virar assinante por estar na academia".
- O caminho do scan é **100% público** e toca **só o catálogo global** (sem dado tenant, sem RLS de
  acesso) — ver I3 em [04](04-questoes-abertas.md).

## 3. Acesso = assinatura (não o scan)

**O que destrava treino/avaliação é a assinatura**, em camadas:

1. **Anônimo:** só vídeo + CTA.
2. **Plano Básico (pago):** avaliação física + treino montado + periodização. Logado. É o produto
   que monetiza (= motor do "treinador", ver I2).
3. **Plano Plus (pago):** treinos/periodizações avançados + contato com o time de especialistas.

### A página de scan é contextual (render por sessão + tier + org)

Mesma URL `/q/<token>`. **Co-branding:** se o aparelho está catalogado sob uma org parceira, a
página mostra **nome + logo** da academia/condomínio (não fica genérica).

O conteúdo depende de **estado do usuário × parceria da org**:

| Usuário | Org **sem parceria** | Org **parceira facility** |
|---|---|---|
| Anônimo / **deslogado** (sem sessão) | vídeo + **"assine o Básico"** + **"já é assinante? faça login"** | vídeo + **"fale com a recepção p/ liberar"** + **"faça login"** + "assinar por conta própria" (secundário) |
| **Básico** válido (logado) | vídeo + **seu treino de hoje** + CTA discreto p/ Plus | idem |
| **Plus** válido (logado) | vídeo + treino avançado, sem upsell | idem |

> A primeira linha cobre tanto o visitante novo quanto o **assinante num device deslogado** — por
> isso o "faça login" convive com o "assine". Após login, a página re-renderiza no estado do tier.

"Liberar pela recepção" = entrar no **roster pago** do parceiro (conciliação planilha/API, [00 §8](00-estrategia-spinoff.md));
no MVP é manual. O CTA de recepção só aparece **se o usuário ainda não estiver habilitado**.

**Técnico (reforça I3):** a página tem **duas fontes** — (a) o **vídeo**, sempre, do catálogo
**global/público** (sem auth, sem RLS); (b) **se houver sessão válida**, o **treino do usuário**
para aquele arquétipo, **autenticado e RLS-scoped**. Uma URL, dois caminhos. A camada de treino
chega com a **fase 2** (quando existem Básico + motor de treino); a **fase 1** é só (a) + CTA.

`Subscription` modela a assinatura ativa com `tier: basico | plus` e `source: personal | sponsored`:
- **personal** — o usuário paga (Asaas/Pix).
- **sponsored** — um **patrocinador** paga (via `Sponsorship`): pode ser **academia/condomínio**
  (org parceira) **ou** **professor/personal** patrocinando seus alunos. Concedido por **conciliação
  de roster** (nossa base cruzada com a lista do patrocinador, planilha→API), **não** por escanear.
  Ver [00 §8](00-estrategia-spinoff.md).

`Entitlement` (acesso efetivo) deriva da `Subscription` ativa — "o usuário paga" OU "um patrocinador
paga pelo roster". Mantê-la concentrada, não em `if`s.

## 4. MVP — escopo (wedge academia)

Fases incrementais, cada uma entregável:

1. **Isca grátis — QR → vídeo + ad/CTA** (1 academia piloto). Tag resolve arquétipo → vídeo
   educativo + CTA para assinar. Público, sem login. Valida aquisição. Vídeo servido da VPS
   (ver [01 §5](01-stack-e-infra.md)).
2. **Plano Básico (pago) — o produto que monetiza.** Signup → **avaliação + treino montado +
   periodização** (motor do treinador, ver I2). Conversão B2C via Asaas/Pix. **É o coração do
   negócio** — depende de portar/integrar o motor da VPS treinador.
3. **Patrocínio (`Sponsorship`).** Academia/condomínio **ou professor** paga o Básico para seus
   usuários; **conciliação de roster** (planilha → depois API) cruzando nossa base com a lista do
   patrocinador.
4. **Plano Plus** (avançado + especialistas). Professor/personal entram como ator `Professional`
   (não como `Organization`).

## 5. Decisões (TODAS FECHADAS)

- **D1 — Shape do stack:** composto self-hosted na VPS (Postgres + Drizzle + Better Auth + RLS no
  próprio box). Supabase descartado. Ver [01 §3](01-stack-e-infra.md).
- **D2 — Vídeo (alvo de migração):** Bunny Stream (custo). Migração pós-MVP. Ver [01 §5](01-stack-e-infra.md).
- **D3 — PSP:** Asaas (BR-first, Pix Automático, boleto, split nativo p/ parcerias). Stripe
  descartado. Ver [01 §2](01-stack-e-infra.md).
- **D4 — Auth:** Better Auth (self-hosted, org plugin). Clerk descartado. Ver [01 §2](01-stack-e-infra.md).
- **D5 — Marca:** identidade fechada (ouro/petróleo, Trajan Pro + Montserrat, "Intelligence in
  Motion"). Ver [03-marca.md](03-marca.md).
- **D6 — `qr_token`:** token opaco DB-backed + revogável. **Atualizado pelo modelo de negócio:** o
  scan é isca pública e **não concede acesso** → geofence/TTL/anti-clonagem **removidos**. Ver §2.
- **D7 — Seeding do catálogo:** ver §6.

## 6. Seeding do catálogo (D7) — duas fontes

> ⚠️ **Em revisão (I1/I2):** descobriu-se que o **EMR é cópia parcial** do sistema original
> **App Treinador** (VPS `72.62.108.11`). A fonte mais rica (vídeos novos + biomecânica + gerador
> de treino `gerador_treinos.py`) está **na VPS treinador, não no EMR**. O seed provavelmente
> deve **importar direto do treinador**. Confirmar após acesso à VPS. O abaixo é o entendimento atual.

Como o Gymnai vira **dono do catálogo** (e o EMR congela), é **import único, não sync vivo**. São
**duas origens distintas**:

- **Exercícios + GIFs antigos → do EMR.** Script Go no EMR dumpa `exercises` para JSON
  (`ExternalId`, nomes PT/EN, músculos, body parts, equipamentos, instruções, + JSONB de
  enriquecimento NSCA/biomecânica) e copia os GIFs de `/uploads/exercises/`. Import no Gymnai.
- **Vídeos por arquétipo → da VPS treinador** (I1). A biblioteca IA/biomecânica vive na **VPS
  treinador** (não no EMR — o EMR só recebeu os GIFs antigos num sync anterior). Importar
  direto VPS treinador → Gymnai. **Uso comercial liberado.**
  - **Pendente (precisa acesso à VPS treinador — chaves novas):** inventariar formato/resolução/
    duração/quantidade e mapear o gap desde o último sync VPS→EMR.
  - **Normalização no import (M3):** o import roda `ffprobe` e transcodifica **uma vez** para web
    (H.264/`+faststart`/720p), fora da VPS de prod — qualquer que seja o formato de origem.
- **Não vem no import** (é nativo do Gymnai): `Equipment`, `qr_token`, mapeamento aparelho→arquétipo,
  `Member`, `Entitlement`.
- **Timing:** os **vídeos** são necessários já na **fase 1** (scan→vídeo). O **catálogo de
  exercícios** (treino pessoal) só na **fase 3** — mas os formatos de import ficam fixados agora.

## 7. Checklist de início (`/home/user/gymnai`)

**Fase 0 — setup + design (antes do scaffold) (M6):**
- [ ] Criar `/home/user/gymnai` + `git init` (repo próprio, fora do monorepo Plenya).
- [ ] Copiar [CLAUDE-seed.md](CLAUDE-seed.md) → `/home/user/gymnai/CLAUDE.md` + semear memória.
- [ ] Provisionar **VPS dedicada dimensionada** (~**4GB** RAM+, não 1GB) + Cloudflare na frente +
      **HTTPS** (Let's Encrypt + auto-renew, ou Cloudflare SSL).
- [ ] Desenhar o **middleware auth+RLS** (3 contextos: `current_user`/`current_org`, `SET LOCAL` em
      transação, role de serviço p/ jobs) — Item 3.
- [ ] Definir a **interface do `VideoService`** (local→Bunny; URL versionada — M1/M4).
- [ ] Lado EMR: **congelar** o módulo de treino (só bug fix). Ver [00 §6](00-estrategia-spinoff.md).

**Fase 1 — build:**
- [ ] Scaffold Next.js 16 (PWA) + Postgres + Drizzle + Better Auth.
- [ ] Migration 00001: atores escopados + **RLS** ligado desde o início.
- [ ] Fluxo **QR → vídeo + CTA** (`Equipment` → `EquipmentModel` → arquétipo → `Video`).
- [ ] **Admin UI de onboarding** (ver §8) — cadastro + geração de QRs.

## 8. Onboarding de academia + admin (I6)

**Admin UI (deliverable do MVP)** — equipe Gymnai opera:
- Cadastrar **academias** (`Organization` + `Location`s, com nome/logo/partnership).
- Cadastrar **modelos de equipamento** (`EquipmentModel`) e **mapeá-los a arquétipo(s)** — uma vez,
  reusado em todas as academias. Pode ser **genérico** (só "leg press") ou **específico**
  (marca/modelo); marca e modelo são opcionais.
- Cadastrar **aparelhos** (`Equipment`, instância numa Location, referenciando um modelo).
- **Gerar os links + QR Codes** (um `qr_token` por aparelho) para impressão.

**Fluxo operacional do MVP (equipe Gymnai faz tudo):**
1. Gymnai fecha parceria com a academia.
2. Academia envia a **lista de aparelhos** (modelos).
3. Gymnai **cadastra** (academia, modelos→arquétipo, aparelhos), **gera e imprime os QRs**.
4. Gymnai vai à academia e **aplica** as placas.

Depois (pós-MVP): produtizar um **kit auto-instalável** (placas impressas + instruções) para a
academia colar sozinha. Validação de mundo-real (disposição a pagar, qual ponta converte) é
tarefa de campo do usuário, não bloqueia o código.
