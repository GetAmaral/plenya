#!/bin/bash

set -e

echo "=================================="
echo "BATCH LAB PART 5 - 20 Items"
echo "=================================="
echo ""

# Verificar se API key está definida
if [ -z "$ANTHROPIC_API_KEY" ]; then
    echo "ERRO: ANTHROPIC_API_KEY não definida"
    echo "Execute: export ANTHROPIC_API_KEY=your_key"
    exit 1
fi

# Verificar se containers estão rodando
if ! docker compose ps | grep -q "plenya-api.*Up"; then
    echo "ERRO: Container API não está rodando"
    echo "Execute: docker compose up -d"
    exit 1
fi

# Executar o script Python
python3 /home/user/plenya/scripts/batch_lab_part5_20items.py

echo ""
echo "=================================="
echo "BATCH COMPLETO"
echo "=================================="
echo ""
echo "Relatório salvo em: /home/user/plenya/batch_lab_part5_report.json"
