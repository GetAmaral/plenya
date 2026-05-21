# Decks comerciais Plenya

Dois caminhos coexistem; o **HTML/Playwright** é o de produção atual.

## 1. HTML/CSS → PDF (Playwright) — produção
`scripts/deck-builder/continuum/`:
- `deck.html` — deck canônico (single-file, estilos inline). Editar aqui.
- `render-pngs.js` — preview por slide (`node render-pngs.js --slide=NN`, ~5s → `previews/slide-NN.png`).
- `render.js` — PDF final 1920×1080 → `docs/decks/continuum-plenya-YYYYMMDD.pdf`.
- `preview.html` — grid viewer (publicado em `decks.plenyasaude.com.br`).
- `EDITORIAL.md` — **doc viva** (estado por slide, arco narrativo, regras visuais, lições). Ler SEMPRE antes de mexer.

Assets: HD em `docs/site/images/FotosHd/`; específicos do deck em `docs/decks/_assets/`.

## 2. PPTX nativo (skill `/plenya-deck`)
Skill `.claude/skills/plenya-deck/` gera PPTX on-brand via skill `pptx`. Wrapper:
`scripts/deck-builder/plenya-base.js` (paleta/fontes/layouts) + `build-continuum.js`.

## Workflow de ajuste fino (memória `deck_fine_tuning_workflow`)
Mandar PNG de cada slide alterado para aprovação; **não** gerar PDF até a fase fechar.

## Regras invariantes (memórias `plenya_deck_source_of_truth`, regras editoriais)
🚨 Ler a página do site (`apps/site/app/[locale]/<topico>`) ANTES de gerar slide Plenya.
Sem preços · sem marcas comerciais · sem "medicina preditiva" (usar "antecipatório") · sem
casos identificáveis · sem hashtags · sem em-dash · equipe sempre completa · citar "800+ itens"
e "40 anos". Paleta: gold `B38645`, petrol `063B4F`, ocean `417E8E`, sage `92B8B4`, cream `EAE7DA`.
Tipografia: Nalieta (logo) → Cormorant Garamond (títulos) → Inter (corpo).
Imagens: gpt-image-2, petrol-dominante. Ver [images.md](images.md).
