#!/bin/bash

set -e

echo "🔄 Gerando OpenAPI e TypeScript types..."

# Migrations NÃO entram aqui: o schema é gerenciado por migrations goose, criadas
# à mão e aplicadas no deploy via cmd/migrate. Ver docs/emr/migrations-decisao.md.
#   docker compose exec -w /app api go run ./cmd/migrate up|status|version

# 1. Gerar OpenAPI docs
echo "📚 Gerando OpenAPI docs..."
cd apps/api
swag init -g cmd/server/main.go -o docs || echo "⚠️  Erro ao gerar OpenAPI docs"
cd ../..

# 3. Gerar TypeScript types
echo "🔨 Gerando TypeScript types..."
cd packages/types
npx openapi-typescript ../../apps/api/docs/swagger.json -o src/generated/api-types.ts || echo "⚠️  Erro ao gerar types"
cd ../..

# 4. Gerar Zod schemas
echo "✅ Gerando Zod schemas..."
cd packages/types
npx openapi-zod-client ../../apps/api/docs/swagger.json -o src/generated/api-schemas.ts --export-schemas || echo "⚠️  Erro ao gerar schemas"
cd ../..

echo "✅ Geração concluída!"
