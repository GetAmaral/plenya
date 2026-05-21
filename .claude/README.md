# Documentação técnica Plenya

Documentação modular do monorepo. O índice raiz e as regras invariantes ficam no
[`CLAUDE.md`](../CLAUDE.md) da raiz; cada app tem seu próprio `apps/<x>/CLAUDE.md`. Esta pasta
guarda os detalhes transversais.

## Estrutura
```
.claude/
├── 01-overview.md  02-stack.md  03-architecture.md   # fundação
├── backend/    models · database 🔥 · hooks · service-layer · api-endpoints
├── frontend/   form-navigation · patient-context · tanstack-query
├── domain/     score-system 🎯 · patients · security
├── workflows/  development · database-ops 🔥 · enrichment-automation 🤖 ·
│               adding-features · dev-bypass-auth
├── content/    decks · ebooks · images           # geração de conteúdo
├── social/     README (responder-insta, linkedin-week, social-mcp)
├── mobile/     setup · security · deploy · ota-policy · release-checklist
├── skills/     plenya-deck · linkedin-week · lecture-builder · ebook-builder · responder-insta · pptx
└── commands/   aula · ebook · responder-insta
```

## Como navegar (por contexto)
- **Manipular dados / escores** → `workflows/database-ops.md` + `domain/score-system.md`
- **Nova feature backend** → `apps/api/CLAUDE.md` + `backend/models.md` + `workflows/adding-features.md`
- **Frontend** → `apps/web/CLAUDE.md` + `frontend/*.md`
- **Segurança/LGPD** → `domain/security.md` + `backend/hooks.md`
- **Deck/eBook/imagem/post** → `content/*.md` + `social/README.md`
- **Mobile** → `apps/mobile-pro/CLAUDE.md` + `mobile/*.md`

## Princípios
Arquivos pequenos e focados · links relativos entre docs · não duplicar o que código/git já
registra · **nunca chutar dados verificáveis** (versões, contagens, nomes saem do código/site).
Ao mudar versões ou contagens, conferir na fonte antes de escrever.
