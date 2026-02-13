#!/bin/bash

set -e

echo "🔄 Gerando migrations, OpenAPI e TypeScript types..."

# 1. Gerar migrations com Atlas
echo "📝 Gerando migrations com Atlas..."
cd apps/api
atlas migrate diff --dev-url 'docker://postgres/17/plenya_dev' --to 'file://internal/models' --dir 'file://database/migrations' || echo "⚠️  Nenhuma nova migration necessária"
cd ../..

# 2. Gerar OpenAPI docs
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
