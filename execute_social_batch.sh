#!/bin/bash

# Script executor para batch social enrichment
# Verifica ambiente, executa enriquecimento e gera relatório

set -e

echo "=============================================="
echo "SOCIAL ITEMS BATCH ENRICHMENT"
echo "=============================================="
echo ""

# Verificar ANTHROPIC_API_KEY
if [ -z "$ANTHROPIC_API_KEY" ]; then
    echo "❌ ANTHROPIC_API_KEY não encontrada no ambiente"
    echo "   Configure com: export ANTHROPIC_API_KEY='sua-chave'"
    exit 1
fi

echo "✓ ANTHROPIC_API_KEY encontrada"

# Verificar se API está rodando
echo "🔍 Verificando API..."
if curl -s http://localhost:3001/health > /dev/null 2>&1; then
    echo "✓ API está rodando"
else
    echo "❌ API não está respondendo em http://localhost:3001"
    echo "   Inicie com: docker compose up -d"
    exit 1
fi

# Verificar dependências Python
echo "🔍 Verificando dependências Python..."
python3 -c "import anthropic, requests" 2>/dev/null || {
    echo "❌ Dependências faltando. Instalando..."
    pip install anthropic requests
}

echo "✓ Dependências OK"
echo ""

# Executar enriquecimento
echo "=============================================="
echo "INICIANDO ENRIQUECIMENTO DE 30 ITEMS SOCIAL"
echo "=============================================="
echo ""

python3 /home/user/plenya/scripts/batch_social_enrichment.py

# Verificar resultado
if [ $? -eq 0 ]; then
    echo ""
    echo "=============================================="
    echo "✅ ENRIQUECIMENTO CONCLUÍDO COM SUCESSO!"
    echo "=============================================="
    echo ""
    echo "Próximos passos:"
    echo "1. Revisar relatório: SOCIAL-BATCH-REPORT.json"
    echo "2. Verificar items atualizados no banco"
    echo "3. Testar no frontend"
    echo ""
else
    echo ""
    echo "=============================================="
    echo "❌ ENRIQUECIMENTO FALHOU"
    echo "=============================================="
    echo ""
    echo "Verifique os logs acima para detalhes"
    exit 1
fi
