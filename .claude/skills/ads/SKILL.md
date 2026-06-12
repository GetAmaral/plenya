---
name: ads
description: Criar, avaliar e gerenciar campanhas de anúncios do Dr. Getúlio / Plenya no Meta (Instagram/Facebook) via Marketing API. Avalia eficiência das campanhas rodando, rankeia reels orgânicos por potencial (share/watch), monta campanha de crescimento (impulsionar reels), gera previews, ativa com aprovação e monitora performance. Invocar quando o usuário pedir "anúncio", "campanha", "impulsionar/turbinar reel", "Meta Ads", "avaliar campanha", "crescer seguidores com mídia paga", "boost". Conta @drgetulioamaralfilho. NÃO é pra responder comentários/DM (isso é /responder-insta).
---

# Skill: `/ads` — Gestão de campanhas Meta (Dr. Getúlio / Plenya)

> Referência interna usada por `.claude/commands/ads.md`.
> Conta: **@drgetulioamaralfilho** · Anúncios via **Marketing API direta (curl)**, NÃO Composio.
> Memória relacionada: `[[ads_campanha_autoridade_getulio]]` · Plano: `docs/marketing/plano-campanha-autoridade-getulio.md`

---

## ⚠️ Leia ANTES de qualquer ação

1. **`RECIPES.md`** — snippets de curl prontos e testados (o caminho que FUNCIONA).
2. **`ERRORS.md`** — catálogo de erros já cometidos → causa → fix. **Ler antes de criar qualquer
   coisa.** Cada erro ali já custou tempo; não repetir.

A API do Meta tem armadilhas não-óbvias (objetivo×otimização, capability do app, recursos
só-da-interface). Não improvisar: seguir as receitas. Se um caminho novo for necessário,
**testar em objeto PAUSADO** primeiro (não gera gasto) e, se der erro, consultar/atualizar `ERRORS.md`.

---

## 🚨 Regras de ouro (invioláveis)

- **NUNCA editar/alterar a campanha GLP-1 que funciona** (`120240429386620590`). A ÚNICA ação
  permitida nela é mudar `status` (PAUSED/ACTIVE). Nunca mexer em budget, criativo, público.
- **Nada que gere gasto sem ordem explícita do usuário.** Montar SEMPRE pausado → preview →
  aprovação → só então ativar. "Segue/faz" não é autorização de ativar.
- **Criar tudo PAUSED.** Ativar é um passo separado e explícito (liga ad + ad set + campanha).
- **Não chutar dados** (budget, IDs, otimização): seguir `RECIPES.md`/`ERRORS.md` ou testar pausado.
- **CFM (Res. 2.336/2023)** na copy: sem superlativo ("o melhor/único"), sem promessa de
  resultado, sem preço de procedimento, sem endosso de marca. Assinar com nome + especialidade
  + `CRM-PR 21.876 · RQE 16.038` quando houver texto.
- **Sem maneirismos de IA / regras editoriais Plenya** (ver `[[plenya_anti_ai_maneirismos]]`).

---

## Ativos e acesso (fonte única)

| Item | Valor |
|---|---|
| Token | `~/.secrets/plenya-meta.env` → `META_MARKETING_TOKEN` (long-lived, **vence 2026-08-10**) |
| Escopos do token | `ads_management, ads_read, read_insights, instagram_basic, pages_show_list, pages_read_engagement, business_management` |
| Ad account | `act_912683771498112` (BRL) |
| Página FB (identidade) | `1046561478538408` ("Clínica médica Dr Getulio") |
| IG Business id | `17841470083252518` (@drgetulioamaralfilho) |
| App Meta | `1296933188599217` (PlenyaMarketing, Live) |
| Graph API | `https://graph.facebook.com/v21.0` |

`source ~/.secrets/plenya-meta.env` no início de cada script. **Nunca imprimir o valor do token.**
Se o token vencer: gerar novo via Graph API Explorer (mesmos escopos) ou migrar p/ System User
(não vence) — ver `ERRORS.md` §token.

---

## Modo de operação (analisa `$ARGUMENTS`)

- "avaliar" / "como estão as campanhas" → **MODO AVALIAÇÃO** (só Fase 1-2, relatório)
- "montar" / "nova campanha" / "impulsionar X" → **MODO CRIAR** (Fase 1→5)
- "ativar" → **MODO ATIVAR** (Fase 6, só com aprovação)
- "performance" / "como tá indo" → **MODO MONITORAR** (Fase 7)
- vazio → perguntar o que o usuário quer

---

## Fases

### Fase 1 — Avaliar campanhas em curso
- `GET act_.../campaigns` (name, status, objective, daily_budget).
- Pra cada ativa: `GET {campaign}/insights?date_preset=maximum&fields=spend,impressions,reach,frequency,cpm,cpc,ctr,clicks,actions,cost_per_action_type`.
- Ler: CTR (bom >2%), CPC, CPM, frequência (>2 = fadiga), e **se o objetivo bate com a meta atual**
  (tráfego ≠ seguidor). Snippets em `RECIPES.md §avaliar`.

### Fase 2 — Rankear reels orgânicos por potencial
- `INSTAGRAM_GET_IG_USER_MEDIA` (Composio, ig_user_id="me", paginar via `after` até esgotar).
- Pra cada reel: `INSTAGRAM_GET_IG_MEDIA_INSIGHTS` metric `["reach","views","shares","saved","total_interactions"]`.
- **Métrica-chave = shares por 1k de alcance** (espalha pra não-seguidor; sinal #2 do algoritmo).
  >3% de shares/reach = viralizando por DM. Watch time é sinal #1.
- Separar educativo (cresce) de pessoal/lifestyle (relacionamento, NÃO impulsionar pra crescer).
- Entregar tabela rankeada. Snippets em `RECIPES.md §reels`.

### Fase 3 — Montar campanha (PAUSADA) — receita que funciona
Para crescer SEGUIDOR (visitas ao perfil — motor da campanha vencedora), usar **exatamente** este
caminho, validado entregando ao vivo (ver `RECIPES.md §criar` + `ERRORS.md`):
1. **Criativo por reel:** `instagram_user_id` + `source_instagram_media_id` (o reel) +
   `call_to_action={"type":"VIEW_INSTAGRAM_PROFILE","value":{"link":"<URL do perfil>"}}`
   (formato OBJETO; o scalar `call_to_action_type` dá #3).
2. **Campanha:** `objective=OUTCOME_TRAFFIC`, `status=PAUSED`, `special_ad_categories=[]`,
   `daily_budget=<cents>` (CBO), `bid_strategy=LOWEST_COST_WITHOUT_CAP`.
3. **Ad sets:** `optimization_goal=VISIT_INSTAGRAM_PROFILE` (NÃO `PROFILE_VISIT` → 1346001),
   `billing_event=IMPRESSIONS`, `destination_type=INSTAGRAM_PROFILE`,
   `promoted_object={"page_id":"1046561478538408","instagram_profile_id":"17841470083252518"}`,
   `attribution_spec=[{event_type:CLICK_THROUGH,window_days:1}]`, targeting (Advantage trava idade mín=25),
   `status=PAUSED`. Com CBO, todos os ad sets com a MESMA otimização.
3. **Anúncios:** `creative={"creative_id":...}`, `status=PAUSED`. 1 anúncio por (reel × ad set).

### Fase 4 — Previews
- `GET {ad}/?fields=preview_shareable_link` (link logado) **e** baixar a capa do reel:
  `GET {media}/?fields=thumbnail_url` → SendUserFile. Mostrar ao usuário. `RECIPES.md §preview`.

### Fase 5 — Aprovação
- Apresentar resumo (campanha/budget/ad sets/anúncios) + previews. **Esperar aprovação explícita.**
  Ajustar (budget/reels/público) se pedido.

### Fase 6 — Ativar (só com "pode ativar")
- Setar `status=ACTIVE` em CADA anúncio, CADA ad set, e a campanha (3 níveis).
- Se for realocação: **pausar a campanha antiga no mesmo momento** (só status).
- Verificar `effective_status` (ads novos ficam `IN_PROCESS`/review por minutos-horas; normal).

### Fase 7 — Monitorar / otimizar
- Após ~3-7 dias: `GET {campaign}/insights` + por ad set/ad. Ler custo por resultado,
  qual reel/público vence. Escalar o que performa, pausar o que não. Atualizar a memória/plano.

---

## Glossário de objetivo×otimização (ODAX, o que a API aceita)
- Crescer SEGUIDOR (visitas ao perfil) → `OUTCOME_TRAFFIC` + `VISIT_INSTAGRAM_PROFILE` +
  `destination_type=INSTAGRAM_PROFILE` + `promoted_object{page_id, instagram_profile_id}` +
  criativo com CTA objeto VIEW_INSTAGRAM_PROFILE. ✅ É o motor da campanha vencedora, 100% via API.
- ⚠️ Enum: a API MOSTRA `PROFILE_VISIT` ao ler, mas na CRIAÇÃO só aceita `VISIT_INSTAGRAM_PROFILE`.
- ❌ THRUPLAY e POST_ENGAGEMENT **NÃO geram seguidor** (engajam o dark post invisível) — foram erro.
- Detalhes e o histórico dos erros em `ERRORS.md`.

---

## Ao terminar
- Atualizar `[[ads_campanha_autoridade_getulio]]` (memória) e `docs/marketing/plano-campanha-autoridade-getulio.md`
  com IDs novos, estado e qualquer erro/aprendizado novo (→ também em `ERRORS.md`).
