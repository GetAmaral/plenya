# eBooks do Dr. Getúlio

Skill: `/ebook` (`.claude/skills/ebook-builder/`). 4 modos: BRIEFING → CAPITULO → REVISAR →
EXPORTAR (PDF). Cada capítulo passa por aprovação humana antes de seguir. Diretório de trabalho:
`ebook/build/<slug>/` (`00-briefing.md`, `_context.md`, `referencias.md`, `NN-<slug>.md` +
`NN-<slug>-figuras.md`). Busca de literatura: RAG Plenya + PubMed (sem alucinar PMIDs).

## Séries em `ebook/`
- **Série AGORA** (`ebook/SerieAgora/`) — 10 eBooks derivados do livro "Antes". Briefing-mestre:
  `00-briefing-mestre.md`; mapa de migração e redistribuição de personagens nos `00-*.md`.
  Livro 1 (Energia e Disposição) em progresso; demais planejados.
- **Série Bases** (`ebook/SerieBases/`) — 17 monografias de suplementos (creatina, magnésio,
  ômega-3, vitamina D, NAD, etc.). `00-briefing-mestre.md`.
- **Performance e Longevidade** — eBook legado já compilado (14 capítulos em PDF).

## Regras
Voz e regras editoriais do [CLAUDE.md raiz](../../CLAUDE.md). ISBNs do livro "Antes": memória
`livro_antes_isbns`. Figuras via gpt-image-2 — ver [images.md](images.md).
