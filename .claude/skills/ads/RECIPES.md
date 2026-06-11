# `/ads` — Receitas (curl testados que FUNCIONAM)

Sempre começar com:
```bash
source ~/.secrets/plenya-meta.env   # define META_MARKETING_TOKEN (nunca imprimir o valor)
T="$META_MARKETING_TOKEN"
ACT=act_912683771498112; PAGE=1046561478538408; IG=17841470083252518
G=https://graph.facebook.com/v21.0
```
Padrão: `--data-urlencode` em tudo; parsear com `python3 -c`; criar SEMPRE `status=PAUSED`.

---

## §avaliar — eficiência das campanhas
```bash
# listar campanhas
curl -s "$G/$ACT/campaigns?access_token=$T&fields=id,name,status,effective_status,objective,daily_budget&limit=50" | python3 -m json.tool
# insights de uma campanha (lifetime)
curl -s "$G/<CAMPAIGN_ID>/insights?access_token=$T&date_preset=maximum&fields=spend,impressions,reach,frequency,cpm,cpc,ctr,clicks,actions,cost_per_action_type" | python3 -m json.tool
# (date_preset alternativos: last_7d, last_30d)
```
Ler: CTR (bom >2%), CPC, CPM, frequência (>2 = fadiga), e se o objetivo casa com a meta atual.

## §reels — rankear potencial (via Composio)
```
INSTAGRAM_GET_IG_USER_MEDIA(ig_user_id="me", fields="id,caption,media_type,media_product_type,permalink,timestamp,like_count,comments_count", limit=40)
   → paginar via `after` até esgotar
INSTAGRAM_GET_IG_MEDIA_INSIGHTS(ig_media_id=<id>, metric=["reach","views","shares","saved","total_interactions"])
```
Métrica-chave = **shares/1k de alcance** (>3% = viralizando por DM). Watch time = sinal #1.
Separar educativo (cresce) de pessoal/lifestyle (não impulsionar pra crescer).

## §criar — campanha de crescimento (impulsionar reels) ✅ caminho que funciona

### 1) criativo flat por reel (SEM CTA)
```bash
curl -s -X POST "$G/$ACT/adcreatives" \
  --data-urlencode "name=AUT - <nome do reel>" \
  --data-urlencode "instagram_user_id=$IG" \
  --data-urlencode "source_instagram_media_id=<IG_MEDIA_ID_DO_REEL>" \
  --data-urlencode "access_token=$T"   # → retorna {"id": creative_id}
```
NÃO passar `call_to_action_type` nem `object_story_spec` (→ erro #3; ver ERRORS.md).

### 2) campanha (OUTCOME_ENGAGEMENT, CBO, PAUSED)
```bash
curl -s -X POST "$G/$ACT/campaigns" \
  --data-urlencode "name=<nome>" \
  --data-urlencode "objective=OUTCOME_ENGAGEMENT" \
  --data-urlencode "status=PAUSED" \
  --data-urlencode "special_ad_categories=[]" \
  --data-urlencode "daily_budget=3000" \  # em centavos = R$30/dia
  --data-urlencode "bid_strategy=LOWEST_COST_WITHOUT_CAP" \
  --data-urlencode "access_token=$T"
```

### 3) ad set (POST_ENGAGEMENT / ON_POST, **IG-only**, PAUSED)
```bash
curl -s -X POST "$G/$ACT/adsets" \
  --data-urlencode "name=<nome do público>" \
  --data-urlencode "campaign_id=<CAMPAIGN_ID>" \
  --data-urlencode "status=PAUSED" \
  --data-urlencode "optimization_goal=POST_ENGAGEMENT" \
  --data-urlencode "billing_event=IMPRESSIONS" \
  --data-urlencode "destination_type=ON_POST" \
  --data-urlencode "promoted_object={\"page_id\":\"$PAGE\"}" \
  --data-urlencode 'targeting={"age_min":25,"age_max":65,"geo_locations":{"countries":["BR"]},"publisher_platforms":["instagram"],"targeting_automation":{"advantage_audience":1}}' \
  --data-urlencode "access_token=$T"
# 🚨 publisher_platforms=["instagram"] é OBRIGATÓRIO p/ crescer o IG — sem isso, placement
#   automático vaza ~80% do budget pro Facebook (ver ERRORS.md). THRUPLAY/ON_VIDEO compra view
#   passiva barata; POST_ENGAGEMENT/ON_POST otimiza engajamento real (preferir).
# com interesses: adicionar "flexible_spec":[{"interests":[{"id":"6003384248805"},...]}] no targeting.
# Interesses úteis (BR): Saúde e boa forma 6003384248805 · Nutrição humana 6002933862573 ·
#   Bem Estar 6003147242240 · medicina natural 6003361698660. Buscar mais: GET /search?type=adinterest&q=...
# CBO → TODOS os ad sets com a MESMA optimization_goal (POST_ENGAGEMENT).
```

### 4) anúncio (PAUSED) — 1 por (reel × ad set)
```bash
curl -s -X POST "$G/$ACT/ads" \
  --data-urlencode "name=<reel> [<público>]" \
  --data-urlencode "adset_id=<ADSET_ID>" \
  --data-urlencode "creative={\"creative_id\":\"<CREATIVE_ID>\"}" \
  --data-urlencode "status=PAUSED" \
  --data-urlencode "access_token=$T"
```

## §preview
```bash
# link de preview (abre logado no FB)
curl -s "$G/<AD_ID>?access_token=$T&fields=preview_shareable_link"
# capa do reel pra mostrar no chat (SendUserFile)
URL=$(curl -s "$G/<IG_MEDIA_ID>?access_token=$T&fields=thumbnail_url,media_url" | python3 -c "import sys,json;d=json.load(sys.stdin);print(d.get('thumbnail_url') or d.get('media_url'))")
curl -s -L "$URL" -o /tmp/reel_cover.jpg
```

## §ativar — só com "pode ativar" (liga os 3 níveis)
```bash
act(){ curl -s -X POST "$G/$1" --data-urlencode "status=ACTIVE" --data-urlencode "access_token=$T" >/dev/null; }
# cada anúncio, cada ad set, e a campanha:
for x in <AD_IDS...> <ADSET_IDS...> <CAMPAIGN_ID>; do act "$x"; done
# se realocação: pausar a campanha antiga NO MESMO momento (só status):
curl -s -X POST "$G/<CAMPANHA_ANTIGA>" --data-urlencode "status=PAUSED" --data-urlencode "access_token=$T"
# verificar (ads novos ficam IN_PROCESS/review por minutos-horas — normal):
curl -s "$G/<CAMPAIGN_ID>/ads?access_token=$T&fields=effective_status&limit=50" | python3 -m json.tool
```

## §monitorar
```bash
curl -s "$G/<CAMPAIGN_ID>/insights?access_token=$T&date_preset=last_7d&fields=spend,reach,impressions,actions,cost_per_action_type&breakdowns=" | python3 -m json.tool
# por ad set / por ad: trocar o id e usar /insights
```

---
## Estado atual (atualizar quando mudar)
- Campanha ativa: `120246291720670590` (Autoridade Getúlio — Engajamento/Seguidor, R$30/dia,
  POST_ENGAGEMENT/ON_POST, **IG-only**). Ad sets `120246338862850590` (Interesses) +
  `120246338863220590` (Aberto). 10 anúncios. (v1 THRUPLAY+placement-auto foi refeita em 2026-06-11
  após vazar 81% pro FB — ver ERRORS.md.)
- Criativos: Ep1 `1881768479183674` · Ep2 `2250980205643279` · Ep7 `1337845195077603` ·
  Proteína `1576050790745622` · Hidratação `1763630304623063`.
- GLP-1 (NÃO MEXER, só status): `120240429386620590` — PAUSED desde 2026-06-10.
