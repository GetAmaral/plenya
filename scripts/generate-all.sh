#!/bin/bash

set -e

echo "🔄 Gerando OpenAPI e TypeScript types..."

# Migrations NÃO entram aqui: o schema é gerenciado por migrations goose, criadas
# à mão e aplicadas no deploy via cmd/migrate. Ver docs/emr/migrations-decisao.md.
#   docker compose exec -w /app api go run ./cmd/migrate up|status|version

# 1. Gerar OpenAPI docs (swag emite Swagger 2.0). Não há Go no host — roda no container api
# (mesma convenção de migrate:up). swag vem pinado no go.mod; .swaggo mapeia tipos externos
# (datatypes.JSON/JSONMap, fiber.Map) que o swag não resolve sozinho.
echo "📚 Gerando OpenAPI docs (container api)..."
docker compose exec -T -w /app api go run github.com/swaggo/swag/cmd/swag init -g cmd/server/main.go -o docs --overridesFile .swaggo || echo "⚠️  Erro ao gerar OpenAPI docs"

# 2. Converter Swagger 2.0 -> OpenAPI 3.0 (openapi-typescript v7 exige 3.x; swagger2openapi já é dep de apps/api)
echo "🔁 Convertendo Swagger 2.0 -> OpenAPI 3.0..."
cd apps/api
npx swagger2openapi docs/swagger.json -o docs/openapi.json --patch || echo "⚠️  Erro ao converter para OpenAPI 3.0"
cd ../..

# 3. Gerar TypeScript types (a partir do OpenAPI 3.0)
echo "🔨 Gerando TypeScript types..."
cd packages/types
npx openapi-typescript ../../apps/api/docs/openapi.json -o src/generated/api-types.ts || echo "⚠️  Erro ao gerar types"
cd ../..

# 4. Gerar Zod schemas (a partir do OpenAPI 3.0)
echo "✅ Gerando Zod schemas..."
cd packages/types
npx openapi-zod-client ../../apps/api/docs/openapi.json -o src/generated/api-schemas.ts --export-schemas || echo "⚠️  Erro ao gerar schemas"
cd ../..

echo "✅ Geração concluída!"
