# 04 — Questões abertas (tracker da revisão)

> Log vivo da **revisão detalhada do plano** (2026-06-29), feita por 3 revisores independentes
> (técnico, segurança/LGPD, produto/GTM) + síntese. Itens fechados um a um com o usuário. Cada
> item: severidade, status, decisão. Atualizar conforme resolvemos.

## Críticos

### I1 — Produção dos vídeos — ✅ RESOLVIDO
Era o maior risco (produto é "scan → vídeo"; o EMR só tinha GIF). **Resolvido:** já existe uma
biblioteca de vídeos **por arquétipo de movimento**, **gerados por IA ancorada na biomecânica**,
hoje na **VPS treinador**, com **uso comercial liberado**. Serão **importados** para o Gymnai.
- Encaixa no modelo: `Video` é por **arquétipo**, não por marca/modelo de aparelho. `Equipment`
  aponta para o arquétipo. Onboarding de academia = mapear aparelho→arquétipo (rápido), não gravar.
- Moat = a biblioteca que já existe; produção premium/consistente + contexto (QR) + treino
  personalizado (depois) é o diferencial vs YouTube.
- **Subtarefa pendente (precisa de acesso à VPS treinador — usuário buscando chaves novas):**
  inventariar **formato, resolução, duração e quantidade** dos vídeos, e **o que mudou desde o
  último sync VPS→EMR** (o EMR só tem os GIFs antigos; os vídeos novos estão só na VPS treinador).
  - **VPS treinador:** sistema original `App Treinador` (Streamlit/Python + Supabase + Gemini +
    ChromaDB), em `72.62.108.11:/root/app_treinador/` (ref.: `docs/archive/PLANO-MODULO-TREINADOR.md`).
  - **Correção arquitetural (D7):** o **EMR é cópia parcial e downstream** do treinador. A fonte
    mais rica (vídeos + biomecânica + gerador de treino) é a **VPS treinador** → o Gymnai
    provavelmente deve **importar direto do treinador**, não do EMR. Revisar o D7/seed após acesso.

### I2 — Proposta de valor + modelo de negócio — ✅ ESTRUTURA RESOLVIDA (preço + motor pendentes)
**Modelo definido pelo usuário (2026-06-29), documentado em [00 §8](00-estrategia-spinoff.md):**
QR + vídeos são **isca grátis**; o produto pago é a assinatura — **Plano Básico** (avaliação +
treino + periodização = motor do treinador) e **Plano Plus** (avançado + especialistas). Parceria
**facility**: academia/condomínio **paga o Básico** para seus usuários, concedido por **conciliação
de roster** (planilha→API), não por escanear. Isso dissolve o medo "vs YouTube" (o vídeo é isca de
propósito) e responde "por que a academia paga" (perk de retenção, e o QR é grátis pra ela).
- **Lastro do motor:** existe um **gerador de treinos** (`gerador_treinos.py`) na VPS treinador;
  **não** foi portado ao EMR (que só tem periodização IA + recomendação textual + plano manual).
- **Dono do `WorkoutAssignment`: resolvido** — **IA gera + professor** pode montar manual OU
  revisar/ajustar a versão da IA (human-in-the-loop; professor age via acesso delegado — I3/I4).
- **Pendente (precisa VPS):** ler `gerador_treinos.py` (regra simples vs motor) — define o esforço
  da fase 2 (o produto que monetiza).
- **Pendente (decisão do usuário):** **preço** de Básico/Plus e estrutura do patrocínio.

### I3 — RLS × scan + modelo de acesso — ✅ FECHADO (modelo confirmado 2026-06-29)
- **Scan anônimo: resolvido.** Com o modelo de negócio, o scan só toca o **catálogo global**
  (arquétipo→vídeo) + ad; nunca concede acesso nem toca dado tenant. Caminho público sem RLS de
  acesso. Catálogo global não carrega `organization_id`.
- **Scan logado (render contextual):** a mesma página, se houver sessão válida, sobrepõe o
  **treino do usuário** para o arquétipo — caminho **autenticado e RLS-scoped** (path normal de
  auth). Endpoint dual-mode: vídeo público sempre + overlay de treino quando logado. Ver
  [02 §3](02-dominio-e-mvp.md).
- **Modelo de acesso (proposto, 3 contextos) — em validação:**
  - **Consumer (aluno):** dados do próprio usuário (`app.current_user`). Policy permite **dono OU
    professor com vínculo ativo** (`EXISTS ProfessionalClient`) = acesso delegado.
  - **Partner staff (org):** dados da org (`app.current_org`); Better Auth org plugin **só do lado
    parceiro** (aluno NÃO é membro de org).
  - **Professional (professor):** vê seus alunos (delegado) + suas orgs vinculadas.
  - **Plumbing:** `SET LOCAL` dentro de transação (Drizzle `db.transaction`); nunca `SET` de sessão
    solto (pooling). Jobs (conciliação/cobrança) com **role de serviço** que ignora RLS.
  - **Confirmado.** Resta só **implementar** o middleware na Fase 0, antes do scaffold (não é mais
    decisão, é execução).

## Altos

### I4 — Modelo Subscription/Entitlement/Patrocínio — ✅ MODELADO (refinar no schema)
Definido em [02 §1/§3](02-dominio-e-mvp.md): `Subscription(tier: basico|plus, source:
personal|sponsored, asaas_id)`; **`Sponsorship`** generaliza o antigo "facility" — patrocinador é
`organization | professional` (academia/condomínio **ou** professor patrocinando seus alunos),
conciliado por roster (planilha→API), **não** geofence. Novo ator **`Professional`** com
`ProfessionalOrganization` (N:M) e `ProfessionalClient` (alunos). `Entitlement` deriva da Subscription
ativa. Falta detalhar schema/edge cases (aluno patrocinado por 2 fontes, expiração, prioridade).

### I5 — Fundação LGPD — ✅ FECHADO
**Decisão:** avaliação **completa** (ACSM, com PA/labs/risco CV = **dado sensível** art. 11) +
**18+**. Avaliação/treino **100% por IA**; **educador físico do Gymnai (CREF) = responsável técnico**
que assina (revisão opcional); disclaimer **screening/orientação, não diagnóstico**. Requisitos
completos em [00 §9](00-estrategia-spinoff.md): consentimento específico p/ sensível, base legal
(contrato + consentimento), criptografia at-rest + audit, retenção, responsável por privacidade,
fronteira com o prontuário EMR. Geolocalização saiu (sem geofence).
- **A produzir (execução, não decisão):** política de privacidade, fluxos de consentimento, prazos
  de retenção, disclaimer.

### I6 — Onboarding físico + wedge — ✅ FECHADO
**Fluxo MVP (equipe Gymnai faz tudo):** academia envia lista de aparelhos → Gymnai cadastra
(academia, `EquipmentModel`→arquétipo, `Equipment`), **gera/imprime os QRs** e vai **aplicar** as
placas. **Admin UI** é deliverable (cadastro + geração de QRs). Ver [02 §8](02-dominio-e-mvp.md).
Modelo de equipamento mapeado a arquétipo **uma vez**, reusado entre academias. Pós-MVP: kit
auto-instalável. **Wedge:** academia (colocação + traz usuários que pagam pelas 3 fontes já
modeladas); a validação de campo (disposição a pagar) é tarefa do usuário, não bloqueia código.

## Médios

- **M1** — ✅ **Resolvido:** versionar a URL do vídeo (`/videos/<id>/v<n>.mp4`/hash) para invalidar
  sem purge-by-key. Ver [01 §5](01-stack-e-infra.md).
- **M2** — ✅ **Resolvido:** gatilho de migração por **métrica** (banda >~50%, CPU/RAM >80%, streams
  concorrentes), 1ª academia estável antes. Removido o "segunda academia". Ver [01 §5](01-stack-e-infra.md).
- **M3** — ✅ **Resolvido:** validação/normalização acontece **no import** (não upload contínuo):
  `ffprobe` valida + transcodifica uma vez p/ web (H.264/faststart/720p), fora da VPS de prod. Casa
  com o inventário do I1 (normaliza o que vier da VPS treinador). Ver [02 §6](02-dominio-e-mvp.md).
- **M4** — ✅ **Resolvido:** manter vídeo **público + cache CDN** (Cloudflare absorve repetições →
  origin barato); proteger com **Referer/hotlink + rate-limit por IP + alerta de banda**. URL
  assinada **não** (mataria o cache; vídeo é isca pública) — só se houver abuso real.
- **M5** — ✅ **Resolvido/encolhido:** IP é **seu** (treinador+EMR), sem licenciamento de terceiro.
  Fica só: `version_id`/carimbo no import (rastreabilidade) + disclaimer clínico já coberto pelo I5.
- **M6** — ✅ **Resolvido:** checklist agora tem **Fase 0** (VPS ~4GB + HTTPS, middleware auth+RLS,
  interface `VideoService`, freeze do EMR) antes da Fase 1. Ver [02 §7](02-dominio-e-mvp.md).
- **M7** — ✅ **Resolvido (com alta-res):** paleta **oficial** dos swatches rotulados (OURO `#B4894D`,
  PETROL `#022837`, AZUL CLARO `#3C6971`, VERDE/sage `#ADB29B`, CREME `#E1D9CC`) em `gymnai-tokens.css`.
  Fonte: **Cinzel** (OFL) no lugar do Trajan. **Símbolo vetorizado** (OpenCV) → `gymnai-symbol.svg`;
  lockup `gymnai-logo.svg`. Tagline canônica = **INTELLIGENCE IN MOTION** (ignorar "Adaptive…").
  **Favicon + app-icons** gerados em `identidade/icons/` (favicon.ico + PNGs 16→512, apple-touch,
  maskable PWA, símbolo transparente). Único refino: limpeza fina de curvas do símbolo no master (opcional).

## Calibragens (revisores ajustados, não acatados ao pé da letra)

- **Geofence:** segue **leve** (decisão D6); o conserto é higiene de privacidade (I5), não endurecer.
- **Extrair Hono "na fase 1.5":** não é bloqueador pré-código; revisitar se load test pedir.
- **DPO formal agora:** prematuro p/ piloto; basta **responsável por privacidade + política + consentimento**.
- **URLs assinadas de vídeo:** só se houver abuso (conflitam com cache barato e com vídeo público de propósito).
