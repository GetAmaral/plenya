#!/bin/bash

# Script para executar enriquecimento do item TC Tórax via Docker

set -e

echo "========================================================================"
echo "🔬 ENRIQUECIMENTO: TC Tórax - Maior Nódulo Sólido"
echo "========================================================================"

# Verifica se a chave API está disponível
if [ -z "$ANTHROPIC_API_KEY" ]; then
    echo "❌ ERRO: ANTHROPIC_API_KEY não encontrada"
    echo ""
    echo "Para executar este script, defina a variável de ambiente:"
    echo "export ANTHROPIC_API_KEY='sua-chave-aqui'"
    exit 1
fi

echo "✅ ANTHROPIC_API_KEY encontrada"
echo ""

# Instala dependências Python necessárias no container
echo "📦 Instalando dependências Python..."
docker compose exec -T web bash -c "pip install --quiet anthropic psycopg2-binary"

# Executa o script Python
echo ""
echo "🚀 Executando enriquecimento..."
echo ""

docker compose exec -T \
    -e ANTHROPIC_API_KEY="$ANTHROPIC_API_KEY" \
    web python3 /app/scripts/enrich_tc_torax_nodulo.py

echo ""
echo "✅ Execução concluída!"
