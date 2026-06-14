#!/bin/bash
# Geração de OpenAPI + tipos TS/Zod a partir dos Go models (fonte única de verdade).
# FALHA-DURA: qualquer passo que quebrar aborta o script (sem artefato parcial/defasado).
# Migrations NÃO entram aqui (schema = goose, ver docs/emr/migrations-decisao.md).
set -euo pipefail

echo "🔄 Gerando OpenAPI e TypeScript types..."

# 1. OpenAPI docs (swag emite Swagger 2.0). Não há Go no host — roda no container api.
#    .swaggo mapeia tipos externos (datatypes.JSON/JSONMap, fiber.Map) que o swag não resolve.
echo "📚 [1/4] swag init (Swagger 2.0)..."
docker compose exec -T -w /app api go tool swag init \
  -g cmd/server/main.go -o docs --overridesFile .swaggo

# 2. Swagger 2.0 -> OpenAPI 3.0 (openapi-typescript v7 exige 3.x).
echo "🔁 [2/4] swagger2openapi (2.0 -> 3.0)..."
( cd apps/api && npx swagger2openapi docs/swagger.json -o docs/openapi.json --patch )

# 3. TypeScript types (única saída consumida pelo front; via @plenya/types).
echo "🔨 [3/3] openapi-typescript..."
( cd packages/types && npx openapi-typescript ../../apps/api/docs/openapi.json -o src/generated/api-types.ts )

# Zod gerado (openapi-zod-client) foi APOSENTADO: nunca foi consumido em runtime (forms usam
# schemas Zod à mão em apps/web/lib/validations) e o gerador emitia TS inválido. Ver
# docs/emr/plano-tipos-gerados-migracao.md.

echo "✅ Geração concluída."
