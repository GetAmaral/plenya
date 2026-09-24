#!/usr/bin/env bash
# Fonte única do modelo de geração de imagem da OpenAI usado no monorepo.
#
# Todo script que gera imagem dá `source` aqui em vez de escrever o nome do modelo
# na mão. Quando a OpenAI lançar a próxima geração, muda-se UMA linha, aqui.
#
# Família atual: gpt-image-2.5
#   sunburst -> qualidade (padrão)
#   flare    -> velocidade
# gpt-image-2 é a geração anterior; gpt-image-1 e dall-e-3 são legado. Não usar.
#
# Override pontual:  OPENAI_IMAGE_MODEL=gpt-image-2.5-flare ./gen-image.sh ...
#
# Conferir o que existe antes de assumir:
#   curl -s https://api.openai.com/v1/models -H "Authorization: Bearer $OPENAI_API_KEY" \
#     | jq -r '.data[].id' | grep image

OPENAI_IMAGE_MODEL="${OPENAI_IMAGE_MODEL:-gpt-image-2.5-sunburst}"
OPENAI_IMAGE_MODEL_FAST="${OPENAI_IMAGE_MODEL_FAST:-gpt-image-2.5-flare}"
export OPENAI_IMAGE_MODEL OPENAI_IMAGE_MODEL_FAST
