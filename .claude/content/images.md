# Geração de imagens (OpenAI gpt-image-2)

**Modelo correto:** `gpt-image-2` (lançado 21/04/2026; snapshot `gpt-image-2-2026-04-21`).
NÃO usar `gpt-image-1` (gerava texto PT errado) nem `dall-e-3` (legado). Memória `openai_image_model`.

Endpoint: `POST https://api.openai.com/v1/images/generations` com `"model": "gpt-image-2"`.
Retorna base64 em `data[0].b64_json`. API key em `apps/api/.env` (`OPENAI_API_KEY`).

## Wrappers em `scripts/blog-generator/`
- `gen-figure.sh` — infográficos (charts/diagramas com dados reais, paleta Plenya, fundo cream). 1024×1024.
- `gen-image.sh` — hero/editorial atmosférico (tons quentes, DOF raso, sem texto). 1536×1024 ou 1024×1024.
- `gen-illust.sh` — ilustrações conceituais (paleta mais livre).
- Variantes EN: `gen-figure-en.sh`, `gen-figure-translate.sh`; `regen-missing-images.sh` (lote);
  `validate-blog.mjs` (integridade).

Saída: `apps/site/public/images/blog/<slug>/*.webp`. Guias: `STYLE-GUIDE.md`, `TOPICS.md`.

## Imagens de deck
Petrol-dominante, estética quiet-luxury (refs Aman/Brunello), sem ícones/charts, sensibilidade
brasileira. Regras detalhadas em `scripts/deck-builder/continuum/EDITORIAL.md` e [decks.md](decks.md).
