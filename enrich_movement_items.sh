#!/bin/bash

# Script para enriquecer Score Items de Movimento e Atividade Física

API_URL="http://localhost:3001/api/v1"
EMAIL="import@plenya.com"
PASSWORD="Import@123456"

# Login e obter token
echo "Fazendo login..."
TOKEN=$(curl -s -X POST "$API_URL/auth/login" \
  -H "Content-Type: application/json" \
  -d "{\"email\":\"$EMAIL\",\"password\":\"$PASSWORD\"}" | \
  grep -o '"accessToken":"[^"]*' | cut -d'"' -f4)

if [ -z "$TOKEN" ]; then
  echo "Erro ao obter token"
  exit 1
fi

echo "Token obtido com sucesso"
echo "Token: ${TOKEN:0:50}..."

# IDs dos items a processar
ITEMS=(
  "1c1659ac-6511-46b4-a6e4-085fb2a7e934"
  "49c880ac-1c90-47ce-8296-6d488eefc152"
  "0470945d-d26c-4fe8-b2f8-983e189ee327"
  "d33e6d45-936f-4e8a-ad02-789abdf15ae6"
  "bc23dde1-1bc3-4b99-8f31-e862e74e14c6"
  "fd055f4d-13fc-4d78-9988-31336e100104"
  "24fa6cf8-ca10-46fa-89d5-e4b5917dc239"
  "2d73c62a-1e87-40cd-ac6e-434546f5fe3d"
  "e7f6e213-02ee-498f-956c-aabf08171b29"
  "703586f6-7cb2-4ca5-bf39-afcf89a87cad"
)

# Exportar token para uso em outros scripts
echo "$TOKEN" > /tmp/plenya_token.txt
echo "Token salvo em /tmp/plenya_token.txt"
