# `/ads` — Catálogo de erros (causa → fix)

Cada item aqui já custou tempo numa sessão real. **Ler antes de criar qualquer coisa.**
Erros novos: adicionar aqui na hora.

---

## `(#3) Application does not have the capability to make this API call`
Erro de **capability do app**, NÃO de escopo do token. Tem 3 gatilhos distintos:

1. **App em modo Desenvolvimento.** Criar criativo/post de anúncio exige app **Live**.
   Mensagem reveladora (criativo de link): *"post do criativo foi criado por um app em modo
   de desenvolvimento"* (subcode 1885183). → App Dashboard → virar modo **Ativo/Live**.
2. **`call_to_action_type` (scalar) num criativo de reel IG.** QUALQUER CTA (LEARN_MORE,
   VIEW_INSTAGRAM_PROFILE, NO_BUTTON…) dispara `#3`. → **Não passar `call_to_action_type`.**
   Criativo flat sem CTA funciona. (O botão "Ver perfil" da interface não existe na API.)
3. **`object_story_spec` com `page_id` num criativo de reel IG.** → Usar forma flat:
   `instagram_user_id` + `source_instagram_media_id`, sem `object_story_spec`. (A página entra
   no nível do **ad set**, via `promoted_object={page_id}`, não no criativo.)

## `1346001 — A validation error occurred` (na criação do anúncio)
Genérico, mas o gatilho documentado é **`optimization_goal=PROFILE_VISIT`**. "Visitas ao perfil"
é recurso só-da-interface do Meta, sem mapa limpo na API. → Não usar `PROFILE_VISIT` via API.
Para crescer com reel, usar `OUTCOME_ENGAGEMENT` + `THRUPLAY` + `ON_VIDEO` (testado, valida).

## `2490408 — A meta de desempenho não está disponível`
`optimization_goal` incompatível com o `objective` da campanha. Combos válidos testados:
- `OUTCOME_ENGAGEMENT` aceita `THRUPLAY` (com `destination_type=ON_VIDEO`). ✅
- `OUTCOME_ENGAGEMENT` + `POST_ENGAGEMENT`/`PAGE_LIKES`/`THRUPLAY`-sem-destino → recusado.
- `OUTCOME_ENGAGEMENT` + `PROFILE_VISIT` → recusado.
→ Sempre parear objetivo↔otimização↔destino. Na dúvida, testar adset PAUSADO.

## "mesma otimização para seleção de veiculação" (CBO)
Campanha com **orçamento de campanha (CBO)** + lance custo-mais-baixo exige que **todos os
ad sets usem a MESMA `optimization_goal`**. Misturar (ex.: um THRUPLAY, outro REACH) → falha.
→ Padronizar a otimização em todos os ad sets, ou usar orçamento por ad set.

## "is_adset_budget_sharing_enabled é necessário"
Criar campanha **sem** orçamento de campanha exige `is_adset_budget_sharing_enabled=true|false`.
→ Mais simples: usar CBO (`daily_budget` na campanha) e não setar budget no ad set.

## "idade mínima não pode ser > 25 com público Advantage+"
`targeting_automation.advantage_audience=1` trava `age_min<=25`. → Usar `age_min:25` (o Advantage
expande de qualquer forma), ou remover o Advantage se quiser faixa etária estrita.

## Ad reusa criativo IG / criativo com página → "validation error" / 1346001
Anúncio exige identidade de Página. Se o criativo tenta carregar a página e o app não tem a
capability, falha. → Caminho que funciona: criativo flat (só IG) + página no `promoted_object`
do ad set + objetivo/otimização ENGAGEMENT/THRUPLAY/ON_VIDEO.

## `/me/accounts` ou `instagram_accounts` voltam vazios
Token sem `pages_show_list`/`business_management`. Com esses escopos, `/me/accounts` retorna a
página e os `tasks` (confirma MANAGE). O edge `act_.../instagram_accounts` pode voltar `[]` mesmo
funcionando — não é sinal confiável; a identidade IG real está no creative (`instagram_user_id`).

---

## Marketing API Access Tier — REJEITADO (2026-06-11)
Pedido de Standard tier **rejeitado**. Motivo: *"não há número suficiente de chamadas da Ads API
nos últimos 15 dias por este app. É preciso integrar com a Ads API antes de aprovar o Standard
tier."* → É chicken-and-egg: precisa **acumular chamadas reais da Ads API por ~15 dias** e então
**"Solicitar novamente"**. Gerenciar campanhas de verdade já gera essas chamadas.
**NÃO bloqueia nada do nosso fluxo** — o caminho THRUPLAY não precisa de Standard tier. Só
importaria se um dia quiséssemos o objetivo "visitas ao perfil"/botão (recursos só-da-interface).
Recomendação: não perseguir agora; re-submeter em ~2-3 semanas SE decidirmos que vale.

## Token vencendo
`META_MARKETING_TOKEN` é long-lived (~60 dias). Quando perto de vencer: Graph API Explorer →
app PlenyaMarketing → User Token → escopos (`ads_management,ads_read,read_insights,instagram_basic,
pages_show_list,pages_read_engagement,business_management`) → gerar → trocar por long-lived
(`oauth/access_token?grant_type=fb_exchange_token`). Ideal: migrar p/ **System User** (não vence).
