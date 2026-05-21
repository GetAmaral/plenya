# Social media — Dr. Getúlio / Plenya

Conjunto de skills + tooling para gerir presença social. Regras editoriais (anti-maneirismos de
IA, anti-coach americano, casos genéricos, ≤1 menção Plenya a cada 3-4 posts) no
[CLAUDE.md raiz](../../CLAUDE.md) e nas memórias `linkedin_*` / `plenya_anti_ai_maneirismos`.

## Instagram — `/responder-insta`
Skill `.claude/skills/responder-insta/` — workflow de 5 fases (A-E) para responder comentários,
DMs e comments de ads dark (IG + FB). Carrega `contacts.yaml` (~30 pessoas categorizadas: rel,
tom, apelidos, o-que-evitar) na Fase B e atualiza na Fase E.1.
- Varrer **TODOS** os posts, não só os 10 recentes (reels antigos recebem comments novos).
- DM nova de não-seguidor: Composio responde **primeiro** (abre janela Meta). Memória
  `responder_insta_message_requests_first`.
- Tools: `mcp__composio-plenya__INSTAGRAM_*`/`FACEBOOK_*` + servidor [apps/social-mcp](../../apps/social-mcp/CLAUDE.md).
- Limite de chars da doc Composio está errado (aceita 1180+). Memória `composio_ig_reply_char_limit`.

## LinkedIn — `/linkedin-week`
Skill `.claude/skills/linkedin-week/` — ciclo semanal: lê `scripts/linkedin/queue.yaml`, destila
posts do blog do Getúlio (200-250 palavras, hook + números + fecho seco), escolhe imagem,
apresenta drafts. Após aprovação grava `status=approved` + `scheduled_at` no queue.
- Publisher cron (`scripts/linkedin/publisher.py`, ~15min, idempotente) publica nos horários.
- Sem em-dash; tom prosa clínica conectiva. Memórias `linkedin_no_em_dash`, `linkedin_tom_*`.
- Comments API bloqueada por falta de CNPJ (memória `linkedin_cma_pendente`).

## Pendências de credenciais
NAP Plenya + códigos Search Console/Bing: memória `seo_onda3_pendencias`.
