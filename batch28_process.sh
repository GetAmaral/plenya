#!/bin/bash

# BATCH 28 - Movimento e Atividade Física - Conclusão 100%
# 15 items restantes

API_URL="http://localhost:3001/api/v1"
TOKEN=""

# Login e obter token
echo "=== Obtendo token de autenticação ==="
LOGIN_RESPONSE=$(curl -s -X POST "$API_URL/auth/login" \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@plenya.com","password":"admin123"}')

TOKEN=$(echo $LOGIN_RESPONSE | grep -o '"token":"[^"]*"' | cut -d'"' -f4)

if [ -z "$TOKEN" ]; then
  echo "Erro ao obter token!"
  exit 1
fi

echo "Token obtido com sucesso!"
echo ""

# Array de IDs dos 15 items
ITEM_IDS=(
  "0a6df688-8539-4c6f-a7a9-a1356371984d"
  "0ffb33de-5b9a-4f9c-9bf6-5c053962ca97"
  "2672656c-8680-4b13-adf5-bdca1a504517"
  "2d2f2d91-923a-4dd1-9b13-b3d85979f9f6"
  "2e7dbc4d-9ba3-441e-a22a-423acd1990e1"
  "3c2a9cdf-1239-4f75-9cd7-4a09ae4913f5"
  "482f4b9c-2e18-4e4b-8d8a-237a26edb552"
  "717d09cc-2978-44bd-818c-fddd03fbcf69"
  "87154ec9-6550-4d91-b3c1-3d041bb003d4"
  "8e611b4a-a05e-45f3-9aca-0bddd66c8069"
  "a32cd296-be72-404e-9fe7-08f587ea4710"
  "c3b39567-d51a-43c7-a9e0-a29099d7a710"
  "c9565666-5d45-4ccd-a370-e461d8980883"
  "cee904fc-5813-4f15-940c-a4d47cabb237"
  "f87f5ece-1d1c-46f6-b16d-48e59d91f602"
)

# Nomes dos items
ITEM_NAMES=(
  "Suplementação pré, intra e pós treinos"
  "Situação familiar/amigos de exercícios atual"
  "Restrições de atividades"
  "Quem treina"
  "Vida adulta"
  "Situação familiar/amigos de exercícios atual"
  "Vida adulta"
  "Vida adulta"
  "Situação familiar/amigos de exercícios atual"
  "Quem treina"
  "Restrições de atividades"
  "Quem treina"
  "Restrições de atividades"
  "Suplementação pré, intra e pós treinos"
  "Suplementação pré, intra e pós treinos"
)

echo "=== Total de items a processar: ${#ITEM_IDS[@]} ==="
echo ""

# Contador
COUNT=0

# Processar cada item
for i in "${!ITEM_IDS[@]}"; do
  ITEM_ID="${ITEM_IDS[$i]}"
  ITEM_NAME="${ITEM_NAMES[$i]}"
  COUNT=$((COUNT + 1))

  echo "[$COUNT/15] Processando: $ITEM_NAME"
  echo "ID: $ITEM_ID"

  # Aqui cada item será processado individualmente
  # Os textos serão inseridos manualmente após geração

  echo "Item preparado para atualização"
  echo ""
done

echo "=== Script preparado para processar todos os 15 items ==="
echo "Token válido: $TOKEN"
