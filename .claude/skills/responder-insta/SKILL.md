# Responder Insta — Instruções do Skill `/responder-insta`

> Referência interna usada por `.claude/commands/responder-insta.md`.
> Conta gerenciada: **@drgetulioamaralfilho** (IG Business/Creator)
> MCP: tools `mcp__composio-plenya__INSTAGRAM_*`

---

## Modo de operação

Analise `$ARGUMENTS`:

- "dms" / "só dm" → **MODO DM** (pula comentários)
- "comentários" / "comments" → **MODO COMMENTS** (pula DMs)
- "post X" (URL/permalink) → **MODO SINGLE_POST** (só esse post)
- Qualquer outra coisa → **MODO COMPLETO**

Independente do modo, segue o workflow abaixo, pulando fases não aplicáveis.

---

## Fase A — Coleta de pendências (uma chamada paralela)

Dispara **em paralelo**, em uma única mensagem com múltiplas tool calls:

1. `mcp__composio-plenya__INSTAGRAM_GET_IG_USER_MEDIA` (ig_user_id="me", limit=10) — últimos 10 posts
2. `mcp__composio-plenya__INSTAGRAM_LIST_ALL_CONVERSATIONS` (limit=25) — últimas conversas DM
3. `mcp__composio-plenya__INSTAGRAM_GET_IG_USER_TAGS` (ig_user_id="me", limit=10) — mentions
4. `mcp__composio-plenya__INSTAGRAM_GET_USER_INFO` (sanity check — só faz se modo COMPLETO)
5. **(A2 — Ads dark do FB)**: ver seção "Fase A2" abaixo

Quando voltar:
- Pra **TODOS os posts** (não só os 10 mais recentes — comentários novos PODEM cair em Reels antigos!): `INSTAGRAM_GET_IG_MEDIA_COMMENTS` (media_id=$id, limit=100, **fields="id,text,from{id,username},timestamp,like_count,replies{id,text,from{id,username},timestamp,parent_id},parent_id"**) — paralelo. Pague o custo de varrer tudo: pagina via `after` cursor em `GET_IG_USER_MEDIA` até esgotar. **Confirmado em produção (2026-05-12)**: comentário novo da @robertamonteirobo no Reel de Jejum Intermitente (post #18) passou batido com varredura limitada aos 10 mais recentes. Nunca mais.
- Pra **cada conversa** DM: `INSTAGRAM_LIST_ALL_MESSAGES` (conversation_id=$id, limit=10) — paralelo

Em todas as chamadas, peça campos mínimos necessários. Não infle payload.

**🚨 CRÍTICO — bug do username (resolvido 2026-05-13):**
NUNCA peça `username` flat no `fields=`. O Meta IG Graph API só popula `username` flat quando o autor do comment é dono da media — pra demais usuários, o username vem ANINHADO em `from{username}`. Sempre peça `from{id,username}` em vez de só `username`. Isso vale pra `GET_IG_MEDIA_COMMENTS`, `GET_IG_COMMENT_REPLIES` e qualquer call que retorne comments. Sem isso, o autor vem como `null` em ~80% dos comments e a triagem fica inutilizável.

Pra acessar o handle nos resultados: `comment.from.username` (não `comment.username`).

---

## Fase A2 — Ads dark da FB Page (Marketing API)

**Confirmado em produção (2026-05-13):** ads dark do Facebook (criados no Ads Manager, não publicados na timeline) têm **comments próprios**, separados dos Reels orgânicos do IG. **2 lados** distintos a varrer pra cada ad:

- **Lado IG** (dark IG post clone): `effective_instagram_media_id` → `/media_id/comments`
- **Lado FB** (dark FB post clone): `effective_object_story_id` → `pageId_postId` → comments via Composio FB

### Setup

Credenciais Meta no `~/.secrets/plenya-meta.env`:
```bash
source ~/.secrets/plenya-meta.env
# $META_MARKETING_TOKEN — User token com ads_read, ads_management
# Ad Account: act_912683771498112 (Getulio José Mattos Do Amaral Filho)
# Page: 1046561478538408 (Clínica médica Dr Getulio)
```

### Fluxo

1. **Listar ads ativos** (curl direto, Marketing API):
   ```bash
   curl -s "https://graph.facebook.com/v21.0/act_912683771498112/ads?access_token=$META_MARKETING_TOKEN&fields=id,name,status,effective_status" | jq
   ```

2. **Pegar creative IDs de cada ad** (para descobrir os dark posts):
   ```bash
   curl -s "https://graph.facebook.com/v21.0/$AD_ID?access_token=$META_MARKETING_TOKEN&fields=name,creative%7Beffective_object_story_id%2Ceffective_instagram_media_id%2Cinstagram_permalink_url%7D"
   ```
   Retorna: `effective_object_story_id` (FB) e `effective_instagram_media_id` (IG).

3. **Comments do lado IG** (mesmo token):
   ```bash
   curl -s "https://graph.facebook.com/v21.0/$IG_MEDIA_ID/comments?access_token=$META_MARKETING_TOKEN&fields=id,text,timestamp,parent_id&limit=50"
   ```
   ⚠️ IG side dos ads dark **não retorna `from{username}`** mesmo pedindo — Meta esconde identidade nos dark posts. Apresenta como "anônimo".

4. **Comments do lado FB** (via Composio FB toolkit, **NÃO** via curl direto — token Marketing API é user, não page):
   - Use `FACEBOOK_GET_COMMENTS(object_id=$pageId_postId, fields="id,message,from{id,name},created_time,comment_count,like_count,permalink_url")`
   - Lado FB **retorna `from{id,name}`** quando autor é identificável.

5. **CRÍTICO — Verificar replies de cada comment**:
   - Se `comment_count > 0` no comment FB → chamar `FACEBOOK_GET_COMMENTS(object_id=$comment_id)` pra buscar as replies
   - Replies do Dr aparecem com `from.id = 1046561478538408` (Page ID Clínica médica Dr Getulio)
   - **Sem essa verificação, vai re-responder comments já tratados pelo Dr.**

### Postagem de replies

- **Comment IG** (mesmo do lado dark do ad): `INSTAGRAM_POST_IG_COMMENT_REPLIES(ig_comment_id=$id, message=...)` — **funciona normalmente** pelas connections IG já existentes.
- **Comment FB**: `FACEBOOK_CREATE_COMMENT(object_id=$comment_id_SOMENTE, message=...)`. **NÃO passar `pageId_postId_commentId`** — Composio interpreta `pageId` errado e dá erro `page_id:XXXX not found in your managed pages`. **Sempre só o `comment_id` numérico isolado.**

### Quando rodar Fase A2

- Toda checagem completa, junto com IG
- Tools obrigatórias (ToolSearch primeiro se não estiverem carregadas):
  - `FACEBOOK_GET_COMMENTS`
  - `FACEBOOK_CREATE_COMMENT`

---

## Fase A3 — FB Page Messenger DMs

A Page "Clínica médica Dr Getulio" (`1046561478538408`) pode receber DMs via Messenger. Em 2026-05-13 a Page tinha 0 conversas, mas a varredura completa deve incluir mesmo assim — DMs novas podem chegar via ads.

### Fluxo

1. **Listar conversas Messenger da Page:**
   ```
   FACEBOOK_GET_PAGE_CONVERSATIONS(page_id="1046561478538408", limit=25, fields="id,updated_time,unread_count,participants,message_count")
   ```

2. **Pra cada conversa pendente** (last message do user, não do Dr):
   ```
   FACEBOOK_GET_CONVERSATION_MESSAGES(page_id="1046561478538408", conversation_id="$id", limit=10)
   ```

3. **Postar resposta:**
   - `FACEBOOK_SEND_MESSAGE(recipient_id=$psid, message=...)` (carregar via ToolSearch se necessário)

### Tools necessárias (ToolSearch se não carregadas)
- `FACEBOOK_GET_PAGE_CONVERSATIONS`
- `FACEBOOK_GET_CONVERSATION_MESSAGES`
- `FACEBOOK_SEND_MESSAGE`
- `FACEBOOK_MARK_MESSAGE_SEEN`

---

## 📋 Inventário completo de superfícies (atualizado 2026-05-13)

### ✅ COBERTAS pela skill

| Superfície | Fase | Tool/Método |
|---|---|---|
| IG posts orgânicos — comments | A | `INSTAGRAM_GET_IG_MEDIA_COMMENTS` em TODOS posts |
| IG DMs (inclui story replies, story mentions, reel shares) | A | `INSTAGRAM_LIST_ALL_CONVERSATIONS` + `INSTAGRAM_LIST_ALL_MESSAGES` |
| IG mentions (tags) | A | `INSTAGRAM_GET_IG_USER_TAGS` |
| **Ads dark FB Page — comments lado IG** | A2 | Marketing API curl direto |
| **Ads dark FB Page — comments lado FB** | A2 | `FACEBOOK_GET_COMMENTS` Composio |
| **FB Page Messenger DMs** | A3 | `FACEBOOK_GET_PAGE_CONVERSATIONS` |

### ❌ NÃO COBERTAS (limitações conhecidas)

| Superfície | Motivo | Severidade |
|---|---|---|
| FB Page posts orgânicos — comments | Page tem 0 posts orgânicos (só roda ads) — não aplicável | N/A |
| FB Page mentions (alguém marca a Page num post próprio) | Sem tool fácil no Composio; precisa Graph API direta | Baixa (uso real raro) |
| FB Stories | API limitada; tipicamente vira DM Messenger se replying | Baixa |
| Threads | Threads API existe mas não está no Composio | Baixa (Dr não usa Threads) |
| Mentions hashtag (#drgetulioamaralfilho em posts de terceiros) | Hashtag search API restrita a Business Discovery | Baixa |

Pra cada superfície não-coberta, é decisão consciente baseada em valor/custo.

---

## Fase B — Filtragem e classificação

**Antes de filtrar, carregue o catálogo de contatos E o log de já-respondidos:**

```bash
cat /home/user/plenya/.claude/skills/responder-insta/contacts.yaml
cat /home/user/plenya/.claude/skills/responder-insta/responded-log.yaml
```

O `responded-log.yaml` registra itens **já tratados** — em especial DMs que o Dr respondeu
**pessoalmente no app** (que podem reaparecer como "pendentes" se a API não capturou a resposta).
Cruze cada item pendente com esse log: se já consta como respondido (e não há mensagem nova do
usuário depois), **pule sem perguntar**. Atualize o log na Fase E.1.

Esse arquivo (YAML) contém ~30+ contatos categorizados com:
- `rel` — tipo de relação (esposa, irmã, primo, sogra, paciente, etc.)
- `name` — primeiro nome conhecido
- `call_dr` — apelido que esse contato usa pro Dr ("Get", "Getao", "Getinho", "Monstro")
- `tone` — calibragem de tom recomendada
- `notes` — contexto não-óbvio (ex.: "Tiago é influente em política — sem bajulação")
- `avoid` — coisas a NÃO escrever pra esse contato (palavras, temas)
- `last` — última interação registrada

**Use o catálogo agressivamente.** Quando um `@handle` aparecer, cruze com o catálogo:
- ✅ Achou: já sabe nome, relação, apelidos, tom certo, coisas a evitar. Use tudo.
- ❌ Não achou: trate como novo, classifique pelo conteúdo do comentário. Anote pra adicionar no fim da sessão (Fase E).

Para cada item coletado, determinar:

### Comentários
- **JÁ RESPONDIDO?** → comentário tem reply do @drgetulioamaralfilho na lista `replies`? Pular se sim.
- **É sub-reply (parent_id existe)?** → Geralmente pular, a menos que seja uma resposta nova do usuário a uma resposta do Dr (ex.: "obrigada doutor" após resposta dele). Esses casos viram **fechamento gentil**.
- **Spam óbvio?** → Promotional/foreign/non-sequitur. Marcar como descarte.
- **Senão**: classificar tipo via [PLAYBOOK.md](PLAYBOOK.md) e incluir na fila.

### DMs
- **Última mensagem é do usuário?** (campo `from.username` ≠ "drgetulioamaralfilho") → pendente
- **Última mensagem é do Dr?** → pular (já está aguardando resposta do usuário)
- **Conversation `updated_time` > 24h atrás?** → cuidado: fora da janela livre de DM. Mencionar isso ao gerar draft.

### Mentions (tags)
- Sempre opcional. Pular a menos que o usuário pediu explicitamente. Tags geralmente são posts de família/colegas — não precisam resposta no IG do Dr.

---

## Fase C — Apresentar triagem

Mostre tabela compacta:

```
[N] TIPO | @USER | preview... | há X horas | classificação
```

Exemplo:
```
[1] DM   @maria_silva    "Doutor, sobre creatina pós-menopausa..."  6h   PERGUNTA CLÍNICA
[2] COM  @joao           "Posso treinar em jejum?"                  2h   PERGUNTA CLÍNICA
[3] COM  @ana            "Que aula maravilhosa!"                    1d   ELOGIO
[4] COM  @bot123         "Como ganhar dinheiro online..."           3h   SPAM (descartar)
[5] DM   @paciente_x     "Quero remarcar consulta"                  4h   AGENDAMENTO
```

**Pergunte ao usuário:**
- "Rodo todas (pulando spam), ou filtra por tipo/prioridade?"
- "Algum item pra eu pular já agora?"

Se ele responder algo como "só clínicas" ou "vai todas", siga.

---

## Fase D — Iterar item por item

Para cada item aprovado pela triagem, fazer **inline approval loop**:

### Passo D.1 — Carregar contexto completo

- **Se DM:** já tem últimos 10 messages. Identifique a thread completa (queixa do usuário → eventual histórico).
- **Se COMENTÁRIO:** carregar o post pai (caption do reel/post + tema central). Importante pra responder relevante ao contexto do conteúdo.

### Passo D.2 — Buscar evidência (se PERGUNTA CLÍNICA)

Veja [RAG.md](RAG.md) para procedimento. Resumo:

1. **Sempre primeiro**: `bash .claude/skills/responder-insta/scripts/search-rag.sh "<termo da pergunta>"` — busca semântica nos artigos da base científica Plenya
2. **Se tópico exige evidência ≥2024**: `WebSearch` em inglês com termos científicos + PubMed/NEJM/Lancet/BMJ
3. **NUNCA invente referência.** Se não achar fonte, melhor escrever menos do que citar errado.

### Passo D.3 — Escrever draft

**Primeiro: consultar `contacts.yaml`** pelo @handle do item. Se houver entrada:
- Use `name` no draft (não fica "querida" genérico — fica "Mari querida", "Tiago")
- Aplique `tone` recomendado
- Respeite TODA a lista `avoid` (palavras, temas, frases)
- Use `notes` pra contextualizar (ex.: Michell tem 3 meninas → não chamar de "tropa"; Tiago é figura política → sem bajulação/marketing)

Aplicar o **tom Dr Getúlio** detalhado em [TONE.md](TONE.md). Resumo:

- Abertura: "Pergunta importante", "Pergunta relevante", ou variante reconhecida
- 1ª pessoa, direta, sem rodeios
- Linguagem **leiga** — sem jargão técnico não-traduzido (ver TONE.md lista de termos a evitar)
- Cita evidência (estudos, sociedades — KDIGO, ESC, ACSM, etc.) mas em prosa, não Vancouver
- Recusa diagnóstico/prescrição em rede social: redireciona pra consulta com nefrologista/médico
- Fechamento aberto e colaborativo, NUNCA imperativo. Ex.: "**No seu caso, a princípio vale a pena considerar — conversa com seu nefrologista e decide junto com ele**"
- Comprimento típico: 600–1.200 caracteres (sim, o limite real é >1.180 chars, confirmado em produção em 2026-05-11)

Para tipos não-clínicos, ver [PLAYBOOK.md](PLAYBOOK.md).

### Passo D.4 — Apresentar pacote pra aprovação

Formato exato:

```
[N/total] TIPO @username — preview do contexto

CONTEXTO:
<thread DM completa OU caption do post + comentário original>

DRAFT:
> <texto do draft>

FONTES USADAS:
- RAG Plenya: <título do artigo se aplicável>
- WebSearch: <citação curta>
- (ou: "nenhuma — resposta de tom/recepção")

CARACTERES: ~N | PALAVRAS: ~N

[a]provar e postar | [e]ditar | [p]ular | [r]ejeitar definitivamente | [s]top
```

### Passo D.5 — Reagir à escolha

- **[a]** → postar via MCP correspondente:
  - Comentário: `mcp__composio-plenya__INSTAGRAM_POST_IG_COMMENT_REPLIES` (ig_comment_id=$id, message=<draft>)
  - DM: `mcp__composio-plenya__INSTAGRAM_SEND_TEXT_MESSAGE` (recipient=$user_id, text=<draft>) + `INSTAGRAM_MARK_SEEN` na conversação depois
  - Registrar resultado: comment_id ou message_id da resposta

- **[e]** → pedir ao usuário o que ajustar, refazer o draft mantendo as fontes, apresentar v2. Loop até [a] ou [r].

- **[p]** → pular este item, ir pro próximo

- **[r]** → registrar que ele recusou esse item (não tentar de novo na próxima sessão da skill)

- **[s]** → parar tudo, fazer resumo parcial e sair

### Passo D.6 — Pós-post

Após postar com sucesso, guardar em memória de sessão:
- `comment_id` original
- `permalink` do post pai (necessário pro link de like manual depois)
- timestamp do post da resposta
- categoria

---

## Fase E — Resumo final + atualizar catálogo + likes manuais

### E.1 — Atualizar contacts.yaml + responded-log.yaml

**Sempre** atualizar os DOIS arquivos no fim da sessão.

No `responded-log.yaml`: registrar tudo que foi respondido nesta sessão (DM, comentário, ad),
com `by` (`dr_app` / `dr_api` / `page`), data e `re` curto. Se o Dr disse "essa eu já respondi
pessoalmente", grave como `by: dr_app` — é exatamente o caso que o log existe pra cobrir. Atualizar
`last_updated` e `comments_baseline_swept_through`.

No `contacts.yaml`. Para cada contato respondido nesta sessão:

- **Já existe no catálogo?** → atualizar o campo `last` com a data + breve contexto
  ```
  last: "2026-05-12 - reply em post Creatina (pergunta clínica)"
  ```
- **Não existe ainda?** → adicionar nova entrada com pelo menos `rel` (mesmo que "seguidor" genérico) e `last`
- **Aprendeu algo novo nesta sessão sobre alguém?** (ex.: descobriu que é tia em vez de irmã; que tem trigêmeos; que tem filhas e não filhos) → **OBRIGATÓRIO atualizar `notes` e `avoid`**

Usuário corrigir alguma coisa na sessão (ex.: "não é irmã, é tia"; "Michell tem meninas, não chame de tropa") = sinal forte de que precisa virar entrada permanente no `contacts.yaml`. Anote sem perguntar.

Também atualizar campo `last_updated` no topo do arquivo.

### E.2 — Resumo da sessão

Quando terminar todos os itens (ou usuário deu `[s]`), mostre:

```
✅ N respondidos (X comentários + Y DMs)
⏭  N pulados ([p])
🚫 N rejeitados ([r])

📋 LIKES MANUAIS (bate ❤️ nesses comentários — Composio/Meta não automatiza ainda):

  [post 1] https://www.instagram.com/reel/ABC123/
    ↳ comment de @user1 "..." (respondido agora)
    ↳ comment de @user2 "..." (respondido agora)
  
  [post 2] https://www.instagram.com/reel/DEF456/
    ↳ comment de @user3 "..." (respondido agora)
```

Abrir os permalinks NÃO é necessário — apenas listar. Ele que decide se vai bater ❤️ agora ou depois.

---

## Restrições absolutas

- **NUNCA postar sem aprovação `[a]` explícita.** Mesmo replies "óbvios" como agradecimento. O usuário aprova TODOS.
- **NUNCA inventar dose, diagnóstico, posologia.** Se a pergunta exige prescrição, sempre redirecionar pra consulta.
- **NUNCA citar fonte sem ter visto o conteúdo dela.** Hallucination de PMID/artigo é falha de skill.
- **NUNCA agir em escopo fora da @drgetulioamaralfilho.** A skill **não posta** na conta @plenyaSaude nem na Página FB. Outra skill futura se precisar.
- **NUNCA usar `INSTAGRAM_DELETE_COMMENT`** sem pedido EXPLÍCITO do usuário (irreversível).
- **NUNCA marcar `MARK_SEEN`** antes de o usuário aprovar a resposta — manter o "não lido" preserva o lembrete visual no app dele.
- **Limite de DM:** se uma conversa tem última mensagem do usuário >24h, mencionar antes de gerar draft: a Meta tem janela de 24h pra DM livre — fora dela, só "human agent" reply (que tem regras próprias). Pedir confirmação se quer prosseguir.

---

## Compose vs. Composio — uso de tools

Use **somente** os tools `mcp__composio-plenya__INSTAGRAM_*`. Se algum tool não estiver disponível na sessão (deferred), carregue via ToolSearch com `select:<TOOL_NAME>`.

Tools obrigatórios pra essa skill:
- INSTAGRAM_GET_IG_USER_MEDIA
- INSTAGRAM_GET_IG_MEDIA_COMMENTS
- INSTAGRAM_POST_IG_COMMENT_REPLIES
- INSTAGRAM_LIST_ALL_CONVERSATIONS
- INSTAGRAM_LIST_ALL_MESSAGES
- INSTAGRAM_SEND_TEXT_MESSAGE
- INSTAGRAM_MARK_SEEN
- INSTAGRAM_GET_IG_USER_TAGS
- INSTAGRAM_GET_USER_INFO (opcional, sanity)

Se o usuário pedir, há também `mcp__composio-plenya__FACEBOOK_*` mas a Página FB tem volume zero orgânico (verificado em 2026-05-11) — não rodar a menos que ele peça explicitamente.

---

## Lembretes operacionais

- Tom de Dr Getúlio: ver [TONE.md](TONE.md) com 6 exemplos verbatim dos replies históricos dele.
- Categorias de resposta: ver [PLAYBOOK.md](PLAYBOOK.md) com 7 tipos + política pra cada.
- RAG access: ver [RAG.md](RAG.md) — usa `scripts/search-rag.sh` que hita o backend Plenya (`localhost:3001`).
- **Catálogo de contatos:** [contacts.yaml](contacts.yaml) — ~30 contatos categorizados (família, amigos, colegas, pacientes). Carregar SEMPRE no início (Fase B). Atualizar SEMPRE no fim (Fase E.1).
- Brand essence completa: `~/.claude/projects/-home-user-plenya/memory/plenya_brand_essence.md` (consultar se a pergunta tocar em posicionamento Plenya).
