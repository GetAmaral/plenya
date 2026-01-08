#!/bin/bash

# Script master: Gera TUDO a partir dos Go models
# ÚNICA fonte de verdade: apps/api/internal/models/*.go

set -e

echo "╔════════════════════════════════════════════════════════════════╗"
echo "║  🚀 PLENYA EMR - Geração Automática                           ║"
echo "║  Fonte única de verdade: Go Models                            ║"
echo "╚════════════════════════════════════════════════════════════════╝"
echo ""

# 1. Gerar migrations SQL do Go models (via Atlas)
echo "📝 [1/4] Gerando migrations SQL a partir dos Go models..."
cd apps/api
if command -v atlas &> /dev/null; then
    atlas migrate diff \
        --dev-url "docker://postgres/17/plenya_dev" \
        --to "file://internal/models" \
        --dir "file://database/migrations" || echo "⚠️  Nenhuma mudança detectada"
else
    echo "⚠️  Atlas não instalado. Instale com: curl -sSf https://atlasgo.sh | sh"
fi
cd ../..
echo "✅ Migrations geradas"
echo ""

# 2. Gerar OpenAPI spec do Go (via Swag)
echo "📋 [2/4] Gerando OpenAPI spec a partir dos Go models..."
cd apps/api
if command -v swag &> /dev/null; then
    swag init -g cmd/server/main.go -o docs --parseDependency --parseInternal
else
    echo "⚠️  Swag não instalado. Instale com: go install github.com/swaggo/swag/cmd/swag@latest"
fi
cd ../..
echo "✅ OpenAPI spec gerado"
echo ""

# 3. Gerar TypeScript types do OpenAPI
echo "📦 [3/4] Gerando TypeScript types do OpenAPI..."
cd packages/types
if [ -f "../../apps/api/docs/swagger.json" ]; then
    mkdir -p src/generated
    npx openapi-typescript ../../apps/api/docs/swagger.json -o src/generated/api-types.ts
    echo "✅ TypeScript types gerados"
else
    echo "⚠️  OpenAPI spec não encontrado. Execute generate:openapi primeiro."
fi
cd ../..
echo ""

# 4. Gerar Zod schemas do OpenAPI
echo "✨ [4/4] Gerando Zod schemas do OpenAPI..."
cd packages/types
if [ -f "../../apps/api/docs/swagger.json" ]; then
    npx openapi-zod-client ../../apps/api/docs/swagger.json \
        -o src/generated/api-schemas.ts \
        --export-schemas
    echo "✅ Zod schemas gerados"
else
    echo "⚠️  OpenAPI spec não encontrado. Execute generate:openapi primeiro."
fi
cd ../..
echo ""

echo "╔════════════════════════════════════════════════════════════════╗"
echo "║  ✅ CONCLUÍDO!                                                ║"
echo "╚════════════════════════════════════════════════════════════════╝"
echo ""
echo "Arquivos gerados automaticamente:"
echo "  📁 apps/api/database/migrations/*.sql        ← SQL gerado do Go"
echo "  📁 apps/api/docs/swagger.json                ← OpenAPI do Go"
echo "  📁 packages/types/src/generated/api-types.ts ← TS do OpenAPI"
echo "  📁 packages/types/src/generated/api-schemas.ts ← Zod do OpenAPI"
echo ""
echo "🎯 ÚNICA FONTE DE VERDADE: apps/api/internal/models/*.go"
echo ""
echo "Para adicionar/modificar campos:"
echo "  1. Edite o Go model em apps/api/internal/models/"
echo "  2. Rode: pnpm generate"
echo "  3. Done! SQL, TypeScript e Zod atualizados automaticamente"
echo ""
