#!/bin/bash
# search-pubmed.sh — Busca artigos no PubMed via E-utilities
# Usa Python3 para JSON (jq não disponível no host)

QUERY="$1"
MAX="${2:-10}"
BASE_URL="https://eutils.ncbi.nlm.nih.gov/entrez/eutils"
EMAIL="plenya@plenya.app"

if [ -z "$QUERY" ]; then
  echo "Uso: $0 'query de busca' [max_results]" >&2
  exit 1
fi

# Encode via Python3
ENCODED_QUERY=$(printf '%s' "$QUERY" | python3 -c "import sys, urllib.parse; print(urllib.parse.quote_plus(sys.stdin.read()))")

echo "=== PubMed: '$QUERY' (max $MAX) ===" >&2

# [Title/Abstract] DEVE ser encoded — colchetes sem encode = curl exit 3
FILTER="%5BTitle%2FAbstract%5D"
SEARCH_URL="${BASE_URL}/esearch.fcgi?db=pubmed&term=${ENCODED_QUERY}${FILTER}&retmax=${MAX}&sort=relevance&retmode=json&email=${EMAIL}"

IDS_JSON=$(curl -s --max-time 20 "$SEARCH_URL" 2>/dev/null)
CURL_EXIT=$?

if [ $CURL_EXIT -ne 0 ] || [ -z "$IDS_JSON" ]; then
  echo "⚠️  PubMed inacessível (curl exit=$CURL_EXIT)." >&2
  exit 1
fi

# Parse IDs com Python3
IDS_AND_TOTAL=$(printf '%s' "$IDS_JSON" | python3 -c "
import json, sys
d = json.load(sys.stdin)
r = d.get('esearchresult', {})
ids = r.get('idlist', [])
total = r.get('count', '0')
print(','.join(ids))
print(total)
" 2>/dev/null)

IDS=$(echo "$IDS_AND_TOTAL" | head -1)
TOTAL=$(echo "$IDS_AND_TOTAL" | tail -1)

# Fallback: All Fields se Title/Abstract não retornou nada
if [ -z "$IDS" ]; then
  echo "   Sem resultados em Title/Abstract. Tentando All Fields..." >&2
  SEARCH_URL2="${BASE_URL}/esearch.fcgi?db=pubmed&term=${ENCODED_QUERY}&retmax=${MAX}&sort=relevance&retmode=json&email=${EMAIL}"
  IDS_JSON2=$(curl -s --max-time 20 "$SEARCH_URL2" 2>/dev/null)
  IDS_AND_TOTAL2=$(printf '%s' "$IDS_JSON2" | python3 -c "
import json, sys
d = json.load(sys.stdin)
r = d.get('esearchresult', {})
ids = r.get('idlist', [])
total = r.get('count', '0')
print(','.join(ids))
print(total)
" 2>/dev/null)
  IDS=$(echo "$IDS_AND_TOTAL2" | head -1)
  TOTAL=$(echo "$IDS_AND_TOTAL2" | tail -1)
fi

if [ -z "$IDS" ]; then
  echo "Nenhum artigo encontrado para '$QUERY'" >&2
  exit 0
fi

echo "Total: $TOTAL artigos. Buscando abstracts dos primeiros $MAX..." >&2

# Buscar abstracts em texto puro
FETCH_URL="${BASE_URL}/efetch.fcgi?db=pubmed&id=${IDS}&rettype=abstract&retmode=text&email=${EMAIL}"

echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
curl -s --max-time 30 "$FETCH_URL" 2>/dev/null
FETCH_EXIT=$?
echo ""
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

if [ $FETCH_EXIT -ne 0 ]; then
  echo "⚠️  Erro ao buscar abstracts. PMIDs: $IDS" >&2
else
  echo "=== PMIDs: $IDS ===" >&2
fi
