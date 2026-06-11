Você gerencia as campanhas de anúncios do Dr. Getúlio / Plenya no Meta (Instagram/Facebook) via Marketing API direta (curl), para crescer a conta @drgetulioamaralfilho com mídia paga.

Execute conforme as instruções em `/home/user/plenya/.claude/skills/ads/SKILL.md`.

**Argumentos recebidos:** $ARGUMENTS

**ANTES de criar/alterar qualquer coisa, leia:**
- `/home/user/plenya/.claude/skills/ads/SKILL.md` (workflow + regras)
- `/home/user/plenya/.claude/skills/ads/ERRORS.md` (erros já cometidos → fix — não repetir)
- `/home/user/plenya/.claude/skills/ads/RECIPES.md` (curl prontos e testados)

Lembretes críticos:
- **NUNCA editar a campanha GLP-1 `120240429386620590`** — só mudar status (PAUSED/ACTIVE).
- **Nada que gere gasto sem "pode ativar" explícito.** Montar PAUSED → preview → aprovação → ativar.
- Caminho que funciona p/ impulsionar reel: `OUTCOME_ENGAGEMENT` + `THRUPLAY` + `ON_VIDEO` + criativo flat sem CTA. `PROFILE_VISIT` e CTAs quebram via API (ver ERRORS.md).
- Token no cofre `~/.secrets/plenya-meta.env` (`META_MARKETING_TOKEN`) — **nunca imprimir o valor**.
- Copy obedece CFM (sem superlativo/promessa/preço de procedimento) + regras editoriais Plenya.
- Ao terminar, atualizar a memória `[[ads_campanha_autoridade_getulio]]`, o plano em `docs/marketing/`, e o §"Estado atual" do RECIPES.md.
