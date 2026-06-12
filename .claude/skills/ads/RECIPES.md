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

### 1) criativo por reel (com CTA "ver perfil")
```bash
curl -s -X POST "$G/$ACT/adcreatives" \
  --data-urlencode "name=AUT - <nome do reel>" \
  --data-urlencode "instagram_user_id=$IG" \
  --data-urlencode "source_instagram_media_id=<IG_MEDIA_ID_DO_REEL>" \
  --data-urlencode 'call_to_action={"type":"VIEW_INSTAGRAM_PROFILE","value":{"link":"https://www.instagram.com/drgetulioamaralfilho"}}' \
  --data-urlencode "access_token=$T"   # → retorna {"id": creative_id}
```
🚨 CTA SEMPRE no formato OBJETO (`call_to_action={...}`) com `value.link` = URL do PERFIL.
NUNCA o scalar `call_to_action_type` (→ #3, era payload malformado). NÃO usar `object_story_spec`.

### 2) campanha (OUTCOME_TRAFFIC, CBO, PAUSED)
```bash
curl -s -X POST "$G/$ACT/campaigns" \
  --data-urlencode "name=<nome>" \
  --data-urlencode "objective=OUTCOME_TRAFFIC" \
  --data-urlencode "status=PAUSED" \
  --data-urlencode "special_ad_categories=[]" \
  --data-urlencode "daily_budget=3000" \  # em centavos = R$30/dia
  --data-urlencode "bid_strategy=LOWEST_COST_WITHOUT_CAP" \
  --data-urlencode "access_token=$T"
```

### 3) ad set (VISIT_INSTAGRAM_PROFILE / INSTAGRAM_PROFILE, PAUSED) ✅ motor de SEGUIDOR
```bash
curl -s -X POST "$G/$ACT/adsets" \
  --data-urlencode "name=<nome do público>" \
  --data-urlencode "campaign_id=<CAMPAIGN_ID>" \
  --data-urlencode "status=PAUSED" \
  --data-urlencode "optimization_goal=VISIT_INSTAGRAM_PROFILE" \
  --data-urlencode "billing_event=IMPRESSIONS" \
  --data-urlencode "destination_type=INSTAGRAM_PROFILE" \
  --data-urlencode "promoted_object={\"page_id\":\"$PAGE\",\"instagram_profile_id\":\"$IG\"}" \
  --data-urlencode 'attribution_spec=[{"event_type":"CLICK_THROUGH","window_days":1}]' \
  --data-urlencode 'targeting={"age_min":25,"age_max":65,"geo_locations":{"countries":["BR"]},"flexible_spec":[{"interests":[{"id":"6003384248805"},{"id":"6004115167424"}]}],"targeting_automation":{"advantage_audience":1}}' \
  --data-urlencode "access_token=$T"
# 🚨 optimization_goal é VISIT_INSTAGRAM_PROFILE (enum documentado). PROFILE_VISIT (o que a API
#   MOSTRA na leitura) NÃO é aceito na criação → 1346001. Ver ERRORS.md.
# 🚨 promoted_object precisa de instagram_profile_id (= IG account id) além de page_id.
# Esse é o objetivo "visitas ao perfil" = manda pro perfil = seguidor + curtida visível (motor da vencedora).
# NÃO restringir publisher_platforms (o destino-perfil entrega no IG sozinho).
# Interesses úteis (BR): Saúde e boa forma 6003384248805 · Exercício físico 6004115167424 ·
#   Nutrição humana 6002933862573 · Bem Estar 6003147242240 · medicina natural 6003361698660.
# CBO → todos os ad sets com a MESMA optimization_goal (VISIT_INSTAGRAM_PROFILE).
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
## §abo — orçamento por ad set (1 reel por ad set, sem starvation)
Quando um reel forte morre na CBO (cold-start), migrar pra ABO:
```bash
# campanha SEM daily_budget + is_adset_budget_sharing_enabled=false
curl -s -X POST "$G/$ACT/campaigns" --data-urlencode "name=<nome> (ABO)" \
  --data-urlencode "objective=OUTCOME_TRAFFIC" --data-urlencode "status=PAUSED" \
  --data-urlencode "special_ad_categories=[]" \
  --data-urlencode "is_adset_budget_sharing_enabled=false" --data-urlencode "access_token=$T"
# cada ad set carrega SEU daily_budget (≥600 = R$6, mínimo é R$5,10!) + bid_strategy:
#   --data-urlencode "daily_budget=600"  --data-urlencode "bid_strategy=LOWEST_COST_WITHOUT_CAP"
#   (resto igual ao §criar passo 3). 1 anúncio por ad set.
```
🚨 `daily_budget` < R$5,10 em ABO → erro 1885272 ("orçamento muito baixo"). Usar ≥600.
🚨 Targeting da vencedora = fitness direcionado (NÃO amplo): `flexible_spec:[{interests:[
   {id:6003384248805 Saúde e boa forma},{id:6004115167424 Exercício físico}]}]` + Advantage.

## Estado atual (atualizar quando mudar)
- **Campanha ativa: `120246416551100590`** (Autoridade Getúlio — Visitas ao Perfil **ABO**,
  R$30/dia = 5 ad sets × R$6, OUTCOME_TRAFFIC + **VISIT_INSTAGRAM_PROFILE**/INSTAGRAM_PROFILE).
  5 ad sets, 1 reel cada (Ep1 `…198490590` · Ep2 `…201080590` · Ep7 `…203440590` ·
  Proteína `…205500590` · Hidratação `…208640590`), público = fitness direcionado da vencedora.
  Ativada 2026-06-12, 5/5 ACTIVE sem WITH_ISSUES. Motivo da ABO: a CBO anterior sufocou a Proteína.
- **GLP-1 (vencedora original, NÃO MEXER, só status): `120240429386620590`** — PAUSED.
- Histórico de campanhas deletadas/erradas: a CBO `120246351697660590` (sufocou Proteína) foi
  substituída por esta ABO. THRUPLAY/POST_ENGAGEMENT/PROFILE_VISIT foram ERRO. Ver ERRORS.md.
