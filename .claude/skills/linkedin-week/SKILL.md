---
name: linkedin-week
description: Ciclo semanal de aprovação de posts LinkedIn do Dr. Getúlio. Lê queue.yaml, identifica próximos N posts em status=draft, lê fonte do blog Getúlio, destila para formato LinkedIn (200-250 palavras com hook, números concretos, fecho seco), seleciona imagem do blog, apresenta drafts pra aprovação. Após aprovação atualiza queue.yaml com status=approved + scheduled_at + commentary + image_path. Cron publisher publica nos horários certos. Invocar uma vez por semana (domingo/segunda à noite). Usar para alinhar batch de 2-3 posts antes do início da semana de publicação.
---

# Skill: `/linkedin-week` — Ciclo semanal LinkedIn Dr. Getúlio

## Quando usar

Uma vez por semana (idealmente domingo à noite ou segunda de manhã). O usuário invoca pra:
- Revisar próximos 2-3 posts que vão ao ar na semana corrente/seguinte
- Aprovar, editar ou pular cada um
- Marcar como `approved` na queue para o cron publisher executar nos horários certos

## Onde está o estado

- **Queue**: `scripts/linkedin/queue.yaml` — fonte única, versionada no git
- **Fonte dos blogs**: `apps/site-getulio/content/blog/pt/<slug>.mdx` (sempre essa versão — voz pessoal Getúlio, não a Plenya institucional)
- **Imagens dos blogs**: `apps/site/public/images/blog/<slug>/figura-*.webp` ou `hero.webp`
- **Cron publisher**: `scripts/linkedin/publisher.py` (não mexer — só consome)
- **Token + author URN**: `scripts/linkedin/.env` (não exibir, gitignored)

## Workflow

### 1. Diagnóstico inicial

Leia `scripts/linkedin/queue.yaml` e identifique:
- Qual o último post `status: published` (data + slug)
- Quantos posts em `status: approved` ainda não publicados
- Próximos N posts em `status: draft` (em ordem por `order`)
- Bloco corrente (manifesto / exames / síndromes / método)

Reporte ao usuário: "Último publicado: X em DD/MM. Já aprovados aguardando: N. Próximos drafts: lista." Pergunte quantos quer alinhar (default 2 nas primeiras 4 semanas warmup, 3 depois).

### 2. Cálculo de slots de publicação

Leia `config.cadence` do queue.yaml. Default: `[Tue, Wed, Thu]` às `07:30` BRT.

Calcule os próximos N slots disponíveis a partir de **agora + 1 dia** (nunca agendar pra menos de 24h pra dar buffer de aprovação tardia). Pule slots que já têm post `approved` ou `published`.

Exemplo: hoje é dom 18/05, próximos 3 slots = ter 20/05 07:30, qua 21/05 07:30, qui 22/05 07:30.

### 3. Geração de drafts

Para cada slug a alinhar:

1. **Leia o fonte completo**: `apps/site-getulio/content/blog/pt/<slug>.mdx` (corpo, ignorar frontmatter)

2. **Destile para LinkedIn** (~1300-1800 chars, 200-300 palavras):
   - **Hook na 1ª linha** (uma frase que prende — ideal: punch line do próprio post)
   - **Observação clínica concreta** (primeira pessoa, "vinte anos", "vi paciente", "atendi")
   - **2-3 números específicos** (insulina 22, ferritina 38, ApoB 110, etc — extrair dos dados do post)
   - **Insight central destacado** (em parágrafo curto, isolado)
   - **Menção orgânica do livro Antes e/ou Plenya** (não sales pitch)
   - **Fecho curto e duro** (1 linha, signature line)
   - **NÃO usar** "medicina preditiva" nem martelar "antecipa". Voz Getúlio = primeira pessoa observacional. Voz Plenya = institucional (não usar aqui).
   - **NÃO usar** hashtags (LinkedIn 2026 algorithm não premia; soa amador pra audiência médica)
   - **NÃO incluir links externos** no corpo (algoritmo penaliza). Link Amazon/blog vai em comentário manual depois.

3. **Selecione imagem (FLUXO PENSANTE — não copy-paste):**

   a. **Liste TODAS as figuras disponíveis**:
      ```bash
      ls apps/site/public/images/blog/<slug>/
      ```
      Pode ter `hero.webp`, `figura-1.webp`, `figura-2.webp`, `figura-3.webp` etc. (Ignore variantes `*-en.webp` — são versões em inglês.)

   b. **Veja cada uma visualmente** com a tool Read (Claude lê webp). Não confie só no alt-text do MDX.

   c. **Decida em função do TEXTO CONDENSADO do LinkedIn que você escreveu**, não do post original do blog. O LinkedIn condensa o argumento — a imagem precisa servir ao argumento condensado.

   d. **Critérios de escolha**:
      - **Carrega a tese central** do texto condensado? (Ex: post sobre "1 em 5 com Lp(a) alta" → histograma de distribuição encaixa; post sobre "metáfora alarme vs detector" → quadro comparativo encaixa.)
      - **Mobile readable**? Tabelas/gráficos com texto pequeno ficam ilegíveis no feed mobile. Considerar se vale o risco. Hero atmosférico nunca tem esse problema.
      - **Variedade visual ao longo do bloco**? Evitar 3 hero atmosféricos seguidos OU 3 tabelas seguidas. Calibre olhando o que foi usado nos posts anteriores do bloco.
      - **Coerência com paleta Plenya**? Heros costumam estar 100% na paleta. Figuras técnicas variam.

   e. **Se NENHUMA das figuras existentes encaixar bem com o texto condensado:**
      - **NÃO force uma figura ruim**. Avise o usuário com mensagem clara: "Nenhuma das N figuras disponíveis serve ao texto. As que tem são: ... Posso pedir geração de figura nova com `scripts/blog-generator/gen-figure.sh` (gpt-image-2) — você aprova?"
      - Aguarde aprovação antes de gerar.

   f. **Documente a escolha** ao apresentar pro usuário: "Imagem: <path> (<tipo: figura clínica didática | hero editorial | tabela | diagrama>). Por que essa: <1 linha justificando>."

### 4. Apresentação ao usuário

Mostre os N drafts em sequência clara:

```
═══ POST 1/N — <slug> — agendamento ter DD/MM 07:30 ═══
🖼 Imagem: <caminho> (<tipo: figura clínica | hero | sem imagem>)

[TEXTO COMPLETO DO POST]

—
═══ POST 2/N — <slug> — agendamento qua DD/MM 07:30 ═══
...
```

Pergunte: "Aprovar tudo? Ou ajustes em algum?"

### 5. Iteração

Aceite ajustes em qualquer post:
- "Post 2: troca a abertura por X" → regenere SÓ aquele
- "Post 1: ok. Post 3: pula esse essa semana" → marque post 3 como `skipped: true` e pegue o próximo do bloco
- "Tudo ok" → segue pra próxima fase

### 6. Persistência

Para cada post aprovado, atualize a entry no `queue.yaml`:

```yaml
- slug: <slug>
  block: N
  order: N
  status: approved              # era: draft
  scheduled_at: "YYYY-MM-DD HH:MM:SS-03:00"
  commentary: |
    <texto final>
  image_path: apps/site/public/images/blog/<slug>/figura-1.webp
  image_title: "<título da figura>"
```

**Preserve a estrutura YAML** (use Edit tool, não Write — não reescrever queue inteira). Posts não-alinhados permanecem em `status: draft`.

### 7. Confirmação final

Reporte ao usuário:
- N posts agendados
- Datas/horários exatos
- "Cron publisher executa nos horários. Você não precisa fazer mais nada até a próxima sessão `/linkedin-week`."
- Pergunte se quer fazer commit das mudanças no queue.yaml (recomendado: sim — versiona o que foi aprovado).

## Anti-padrões a evitar

- **Não publicar imediatamente** via `post.py` — cron faz isso. Você só atualiza estado.
- **Não usar a versão Plenya do blog** (tem "nós/equipe" — voz institucional não combina com perfil pessoal). Sempre `apps/site-getulio/`.
- **Não exceder 2 posts/semana nas primeiras 4 semanas** (warmup — algoritmo LinkedIn). Depois ramp pra 3.
- **Não agendar fim de semana** (sáb/dom têm engagement 50% menor pra healthcare).
- **Não cobrar comentário manual com link Amazon** no corpo da skill — está em memória `linkedin_cma_pendente.md`, é workaround conhecido.
- **Não tentar atualizar via API LinkedIn algo já agendado** — queue é fonte. API só pra publish.

## Comandos úteis durante o ciclo

```bash
# Ver estado atual da queue
python3 -c "import yaml; q=yaml.safe_load(open('scripts/linkedin/queue.yaml')); from collections import Counter; print(Counter(p.get('status','draft') for p in q['posts']))"

# Próximos drafts em ordem
python3 -c "import yaml; q=yaml.safe_load(open('scripts/linkedin/queue.yaml')); [print(p['order'], p['slug']) for p in q['posts'] if p.get('status')=='draft'][:5]"

# Ver log do cron publisher (últimas 20 linhas)
tail -20 scripts/linkedin/publisher.log

# Forçar uma rodada do publisher (debugging)
python3 scripts/linkedin/publisher.py
```

## Referências cruzadas

- `docs/seo/08-linkedin-upgrade.md` — pacote inicial de upgrade do perfil
- `~/.claude/projects/-home-user-plenya/memory/linkedin_cma_pendente.md` — comments via API bloqueado (manual)
- `~/.claude/projects/-home-user-plenya/memory/plenya_brand_voice_no_preditiva.md` — vocabulário a evitar
- `~/.claude/projects/-home-user-plenya/memory/linkedin_perfil_pendencias.md` — itens ainda pendentes do perfil
