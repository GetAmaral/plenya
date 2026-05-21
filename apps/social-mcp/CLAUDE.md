# apps/social-mcp — Servidor MCP Instagram/Facebook (Python)

Servidor MCP (stdio) que conecta o Claude Code à conta `@drgetulioamaralfilho` via Meta Graph
API. Complementa as actions Composio (`mcp__composio-plenya__*`) usadas pela skill
`/responder-insta`.

## Stack
Python 3.11+ · MCP · httpx · cryptography (Fernet) · python-dotenv.

## Estrutura (`src/plenya_social/`)
```
server.py        ← dispatcher MCP (registra os tools ig_*)
instagram.py     ← bindings Instagram (user info, media, comments, DMs, tags, like/hide/delete)
meta_client.py   ← wrapper Meta Graph API + refresh de token
oauth_setup.py   ← fluxo OAuth one-time (token long-lived 60 dias)
storage.py       ← persistência criptografada dos tokens
config.py
tokens/instagram.json  ← gitignored, criptografado
```

## Tools `ig_*`
Info (`ig_get_user_info`) · media (`ig_list_media`) · comments (`ig_get_media_comments`,
`ig_post_comment_reply`, `ig_like_comment`, `ig_unlike_comment`, `ig_delete_comment`,
`ig_hide_comment`) · DMs (`ig_list_conversations`, `ig_list_messages`, `ig_send_message`,
`ig_mark_seen`) · tags (`ig_get_user_tags`).

## Registro no Claude Code
```bash
claude mcp add plenya-social python -m plenya_social.server \
  --env PLENYA_SOCIAL_TOKENS_DIR=/home/user/plenya/apps/social-mcp/tokens \
  --env PLENYA_SOCIAL_ENCRYPTION_KEY=<Fernet key>
```

## Notas
- Composio mascara `access_token` (`IGAA...`) e não dá proxy HTTP — se a action não estiver
  wrappeada, não dá pra fazer por lá; este MCP cobre o que falta. Memória `composio_token_bloqueado`.
- Mensagens novas de não-seguidores: o Composio deve responder **primeiro** para abrir a janela
  Meta (não responder manual no app antes). Memória `responder_insta_message_requests_first`.
- Workflow de resposta: skill `/responder-insta` — ver [.claude/social/](../../.claude/social/).
