# Campanha — Série Longevidade & Copa (Autoridade + Seguidores)

Status: **criada PAUSADA** em 2026-06-25 via Marketing API. Zero gasto até ativação manual.
Objetivo declarado pelo Dr: **criar autoridade no assunto + aumentar seguidores no IG** (@drgetulioamaralfilho).
Modelo herdado da campanha vencedora **GLP1** (mesma estrutura validada: PROFILE_VISIT → perfil).

## IDs (Ad Account `act_912683771498112`)

| Objeto | ID | Status |
|---|---|---|
| Campanha | `120247313439350590` | PAUSED · pronta |
| Conjunto (ad set) | `120247313440700590` | PAUSED · pronta · **sem anúncio** |
| Anúncio | — | **não criado** — vídeo bloqueado p/ ads (ver abaixo) |

## ⛔ Bloqueio: o vídeo da série não pode ser impulsionado como está (2026-06-25)

Tentar criar o anúncio com o Reel `18108147745987522` (`reel/DZ8s-GARns6`, "Última dança / CR7")
dá **erro de validação duro** (`error_code 1346001`, HARD_ERROR), sempre com o mesmo `mid`
`04e91cc5bc96dba2fd6f1ad29f37628e` — sinal de bloqueio em **nível de mídia**, não de config.

Diagnóstico provado por A/B: o **mesmo conjunto + mesmo CTA** com outro Reel do Dr
(talking-head "5000 seguidores" `18016844516683101`) valida **limpo, sem issues**. Só este vídeo falha.
→ Config (campanha/público/CTA `VIEW_INSTAGRAM_PROFILE` com `app_link`+`link`) está **100% correta**.

**Causa quase certa:** áudio/música não licenciada para anúncios (Meta bloqueia impulsionar Reels
com som de uso não comercial). Causa secundária possível: uso da imagem/nome do Cristiano Ronaldo.

**Como destravar (qualquer um):**
1. Reexportar o vídeo com **áudio liberado p/ ads** (voz, trilha original, ou música da Sound
   Collection da Meta) e republicar → impulsionar a versão nova.
2. Promover um **episódio da série com áudio próprio** (talking-head valida sem problema).
3. Subir o vídeo direto pelo Ads Manager como mídia nova (sem o som bloqueado).

Quando houver um asset promovível, anexar leva segundos (criativo + anúncio no conjunto já pronto):
o passo do CTA correto é `call_to_action={"type":"VIEW_INSTAGRAM_PROFILE","value":{"app_link":"instagram://user?username=drgetulioamaralfilho&userid=70180144108","link":"http://instagram.com/drgetulioamaralfilho"}}`.

## Configuração

- **Objetivo:** OUTCOME_TRAFFIC (é o único que habilita PROFILE_VISIT; a GLP1 usa o mesmo)
- **Otimização:** `PROFILE_VISIT` · **Destino:** `INSTAGRAM_PROFILE` · **Cobrança:** por impressão · **Lance:** menor custo
- **Página promovida:** `1046561478538408` (Clínica médica Dr Getulio)
- **Criativo:** Reel orgânico da série, sem alteração — `source_instagram_media_id=18108147745987522`
  (post `reel/DZ8s-GARns6` "Última dança / CR7 + Copa"); IG ad id `17841470083252518`; CTA `VIEW_INSTAGRAM_PROFILE`
- **Público:** Brasil · 25–65 · interesses `6003384248805` (Saúde e boa forma) + `6004115167424` (Exercício físico) · **Advantage audience ON**
- **Orçamento:** R$30/dia (`daily_budget=3000`). Teste sugerido: 5–7 dias (~R$150–210).

## Como ativar (quando o Dr autorizar)

```bash
source ~/.secrets/plenya-meta.env
# ativa conjunto e anúncio (campanha já sobe junto):
curl -s -X POST "https://graph.facebook.com/v21.0/120247313440700590" -F "status=ACTIVE" -F "access_token=$META_MARKETING_TOKEN"
curl -s -X POST "https://graph.facebook.com/v21.0/120247313561880590" -F "status=ACTIVE" -F "access_token=$META_MARKETING_TOKEN"
curl -s -X POST "https://graph.facebook.com/v21.0/120247313439350590" -F "status=ACTIVE" -F "access_token=$META_MARKETING_TOKEN"
```

Ou no Gerenciador de Anúncios: localizar "Serie Longevidade & Copa", revisar e ativar.

## Métricas-alvo do teste (benchmark GLP1)

- CTR > 5% · CPC < R$0,16 · custo por ThruPlay ~R$0,04 · alcance ~80k/episódio.
- KPI principal aqui: **visitas ao perfil** e **novos seguidores** (não cliques em site).
- Decidir escala/seguir série pelos episódios que mais puxam seguidor por real gasto.

## Próximas fases (não criadas ainda)

1. **Retargeting da série:** público de quem assistiu ≥50% do Ep1 → impulsionar Ep2 (vínculo sequencial). Exige custom audience de vídeo (montar quando houver views pagos).
2. **Gancho dos 3s:** Ep1 tem skip rate 59,4% e watch 11,2s (vs 14–16s da GLP1). Recortar abertura antes de escalar gasto — Dr optou por **não mexer no vídeo por enquanto** (2026-06-25).

## Avisos

- **Token de Ads do IG vence 2026-08-10** — rodar o teste antes disso.
- **Não ativar sem ordem explícita do Dr** (envolve gasto).
