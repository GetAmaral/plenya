#!/bin/bash

# Script para aplicar a migration do sistema de escores
# Data: 2026-01-21

echo "🔄 Aplicando migration do sistema de escores..."

# Aplicar migration usando docker exec
cat apps/api/database/migrations/20260121_create_score_tables.sql | \
  docker exec -i plenya-db psql -U plenya_user -d plenya_db

if [ $? -eq 0 ]; then
    echo "✅ Migration aplicada com sucesso!"

    # Verificar tabelas criadas
    echo ""
    echo "📋 Verificando tabelas criadas..."
    docker exec -i plenya-db psql -U plenya_user -d plenya_db -c "\dt score_*"

    echo ""
    echo "📊 Contando registros..."
    docker exec -i plenya-db psql -U plenya_user -d plenya_db << EOF
SELECT
    'score_groups' as tabela, COUNT(*) as registros FROM score_groups
UNION ALL
SELECT 'score_subgroups', COUNT(*) FROM score_subgroups
UNION ALL
SELECT 'score_items', COUNT(*) FROM score_items
UNION ALL
SELECT 'score_levels', COUNT(*) FROM score_levels;
EOF
else
    echo "❌ Erro ao aplicar migration"
    exit 1
fi
