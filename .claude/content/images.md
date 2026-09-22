# Geração de imagens (OpenAI gpt-image-2.5)

**Modelo correto:** família `gpt-image-2.5` — `gpt-image-2.5-sunburst` (qualidade) e `gpt-image-2.5-flare` (velocidade), snapshot `-2026-09-08`. `gpt-image-2` é a geração anterior.
NÃO usar `gpt-image-1` nem `dall-e-3` (legado). Memória `openai_image_model`. **Figura de
exercício por IA: só com o nome canônico do exercício no prompt e conferência de cada uma**; memória `ai_imagem_nao_serve_para_figura_de_exercicio`.

Endpoint: `POST https://api.openai.com/v1/images/generations` com `"model": "gpt-image-2.5-sunburst"`.
Imagem de referência: `POST /v1/images/edits` (multipart, `-F image=@ref.png`).
Retorna base64 em `data[0].b64_json`. API key em `apps/api/.env` (`OPENAI_API_KEY`).

## Wrappers em `scripts/blog-generator/`
- `gen-figure.sh` — infográficos (charts/diagramas com dados reais, paleta Plenya, fundo cream). 1024×1024.
- `gen-image.sh` — hero/editorial atmosférico (tons quentes, DOF raso, sem texto). 1536×1024 ou 1024×1024.
- `gen-illust.sh` — ilustrações conceituais (paleta mais livre).
- Os três wrappers ainda passam `gpt-image-2` no `--arg model`; trocar ao regerar.
- Variantes EN: `gen-figure-en.sh`, `gen-figure-translate.sh`; `regen-missing-images.sh` (lote);
  `validate-blog.mjs` (integridade).

Saída: `apps/site/public/images/blog/<slug>/*.webp`. Guias: `STYLE-GUIDE.md`, `TOPICS.md`.

## Imagens de deck
Petrol-dominante, estética quiet-luxury (refs Aman/Brunello), sem ícones/charts, sensibilidade
brasileira. Regras detalhadas em `scripts/deck-builder/continuum/EDITORIAL.md` e [decks.md](decks.md).
