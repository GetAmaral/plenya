#!/bin/bash
# Executa atualização do Batch 31 - Fase 1

cd "$(dirname "$0")/.."

echo "==========================================================================="
echo "BATCH 31 - COGNIÇÃO - FASE 1: TESTES NEUROPSICOLÓGICOS (12 items)"
echo "==========================================================================="
echo ""

docker compose exec -T db psql -U plenya_user -d plenya_db -f /tmp/batch31_phase1.sql

echo ""
echo "==========================================================================="
echo "FASE 1 CONCLUÍDA"
echo "==========================================================================="
