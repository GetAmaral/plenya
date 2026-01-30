#!/bin/bash

# ============================================================================
# BATCH FINAL 4 - HISTÓRICO DE DOENÇAS
# Script de execução e validação
# ============================================================================

echo "============================================================================"
echo "BATCH FINAL 4 - HISTÓRICO DE DOENÇAS"
echo "Enriquecimento MFI de 40 items"
echo "============================================================================"
echo ""

# Cores
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Verificar se Docker está rodando
if ! docker compose ps db | grep -q "Up"; then
    echo -e "${RED}❌ Erro: Container do banco de dados não está rodando${NC}"
    echo "Execute: docker compose up -d db"
    exit 1
fi

echo -e "${YELLOW}📋 Informações do Batch${NC}"
echo "Total de items: 40"
echo "Arquivo SQL: scripts/batch_final_4_doencas_EXECUTAVEL.sql"
echo "Tamanho: $(wc -l < scripts/batch_final_4_doencas_EXECUTAVEL.sql) linhas"
echo ""

echo -e "${YELLOW}⏳ Executando SQL no banco de dados...${NC}"
docker compose exec -T db psql -U plenya_user -d plenya_db < scripts/batch_final_4_doencas_EXECUTAVEL.sql

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✅ SQL executado com sucesso!${NC}"
    echo ""

    echo -e "${YELLOW}📊 Validando resultados...${NC}"
    docker compose exec -T db psql -U plenya_user -d plenya_db -c "
    SELECT
      COUNT(*) as total_items,
      COUNT(*) FILTER (WHERE clinical_relevance IS NOT NULL AND clinical_relevance != '') as com_relevancia,
      COUNT(*) FILTER (WHERE interpretation_guide IS NOT NULL AND interpretation_guide != '') as com_interpretacao,
      COUNT(*) FILTER (WHERE jsonb_array_length(recommendations) > 0) as com_recomendacoes,
      COUNT(*) FILTER (WHERE jsonb_array_length(related_markers) > 0) as com_marcadores,
      COUNT(*) FILTER (WHERE jsonb_array_length(articles_suggestions) > 0) as com_artigos
    FROM score_items
    WHERE id IN (
      '1176540d-cefa-4d2c-b5e2-4a992060de4d',
      '360d1e6a-84c5-4763-a743-0fce76fe2686',
      'adcb06a3-4791-4ce2-ae05-c01f8fef9ead',
      '65a28dbf-9466-4a23-9410-9d6034c3e141',
      'b7370158-c34e-449f-8305-537a8fb85d98',
      'e24dae19-4cb0-4d83-a6db-9571aabf9bde',
      'a7fdd0e5-9aa7-4604-b5ae-731c05c90af0',
      '40837ae0-790a-4a1a-a7ea-fb9d0cc21878',
      'b4cda178-b5a9-4193-b496-afe7ce51ed91',
      'deb8e2ee-27a7-4ac0-a743-d6cb400f4b27',
      'fdb0c742-1db1-48ba-95bc-91e454256d84',
      'e375191b-d8be-4044-9d59-b71fa222cbaa',
      '44f501b3-6c1d-4a31-b8c9-0f54b201dbcb',
      '14881a00-579e-4297-bd6d-4c83d8081d2d',
      'a7698447-ddbf-4bee-92b4-81a1559562ad',
      '76578167-9f69-4981-b73d-46aa7c71ec6b',
      '23b90c6c-6f44-481a-8255-108158a64239',
      '4d1372ee-5bee-4ce8-845a-a4e176109ad2',
      'fe938ed9-3cf9-4ddf-a1ae-2f67c03e54a4',
      '5ddc4af9-dd19-4ea1-8117-3d3e30377dab',
      'e010b2a7-6e9e-4b42-9d37-487e91411a18',
      '933b7816-fe10-4e35-94d7-0d232cc258ce',
      '2b01e2f0-76c7-4616-8f7b-33b1c9d11279',
      '33ffce34-e1fb-4e31-b393-f2783549d874',
      'f54dbfaa-599a-40ee-907a-28431ce4858a',
      '30a7764c-4c11-4078-b942-860ee56136a4',
      '1ff1e7f8-5b15-4f0c-b76b-806c1d6ff1fd',
      '4511ae8d-2ad3-40e0-a47d-2cef065d39e9',
      'cbb42389-9a9d-48fe-a64a-61215107d5e3',
      'e65c56dc-5c07-4270-8a8b-017b293ca147',
      'ad36ad5d-e587-43a8-b1ae-3e27305e25d8',
      '8115cf13-eab0-4db2-9844-d04c37701d92',
      '19b1c83e-f2a7-47e9-9045-5c994cfcbf94',
      'fca251a2-5aa2-42f9-b93c-0f0a852d9d51',
      '925c2474-689e-486a-a6b1-3e1b4ca8cc6e',
      'e5ea8f41-05e0-48e0-b6a8-2323b3896449',
      'f691cda1-bd2d-449e-9e66-045319ceaaa3',
      'ec0bbb80-873a-4946-bf90-bc759eddb080',
      '53dbbf23-25bf-4b91-9670-b17fab93c6e9',
      '8c2456de-1ba1-4a1f-ab7c-964f8c8114e6'
    );
    "

    echo ""
    echo -e "${GREEN}✅ BATCH FINAL 4 CONCLUÍDO COM SUCESSO!${NC}"
    echo ""
    echo "📄 Arquivos gerados:"
    echo "  - scripts/batch_final_4_doencas_EXECUTAVEL.sql"
    echo "  - scripts/batch_final_4_doencas_report.json"
    echo "  - BATCH-FINAL-4-EXECUTE.md"
    echo ""
    echo "============================================================================"
else
    echo -e "${RED}❌ Erro na execução do SQL${NC}"
    exit 1
fi
