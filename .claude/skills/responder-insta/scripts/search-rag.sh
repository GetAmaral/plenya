#!/bin/bash
# search-rag.sh — Busca semântica no RAG Plenya
# Usa Python3 para JSON (jq não disponível no host)

QUERY="$1"
LIMIT="${2:-10}"
API_URL="http://localhost:3001"

if [ -z "$QUERY" ]; then
  echo "Uso: $0 'query de busca' [limit]" >&2
  exit 1
fi

# Encode via Python3 (disponível no host)
ENCODED_QUERY=$(printf '%s' "$QUERY" | python3 -c "import sys, urllib.parse; print(urllib.parse.quote_plus(sys.stdin.read()))")

echo "=== RAG Plenya: '$QUERY' ===" >&2

RESULT=$(curl -s --max-time 20 \
  "${API_URL}/api/v1/articles/semantic-search?q=${ENCODED_QUERY}&limit=${LIMIT}")
CURL_EXIT=$?

if [ $CURL_EXIT -ne 0 ] || [ -z "$RESULT" ]; then
  echo "⚠️  API não respondeu (curl exit=$CURL_EXIT)." >&2
  RESULT=""
fi

# Parse com Python3
RESULT_COUNT=$(printf '%s' "$RESULT" | python3 -c "
import json, sys
try:
    d = json.load(sys.stdin)
    print(len(d.get('results', [])))
except:
    print(0)
" 2>/dev/null)

if [ "${RESULT_COUNT:-0}" -gt 0 ] 2>/dev/null; then
  printf '%s' "$RESULT" | python3 -c "
import json, sys
d = json.load(sys.stdin)
for r in d.get('results', []):
    a = r.get('article', {})
    pub = a.get('publishDate', '')[:4]
    abstract = (a.get('abstract') or '')[:500]
    sim = round(r.get('similarity', 0) * 100)
    print('---')
    print('Título:', a.get('title', 'N/A'))
    print('Autores:', a.get('authors', 'N/A'))
    print('Journal:', a.get('journal', 'N/A'))
    print('Ano:', pub)
    print('PMID:', a.get('pmid', 'N/A'))
    print('DOI:', a.get('doi', 'N/A'))
    print(f'Similaridade: {sim}%')
    print('Abstract:', abstract)
    print()
"
  echo "=== $RESULT_COUNT artigos no RAG ===" >&2
else
  echo "⚠️  Sem resultados via API. Tentando PostgreSQL direto..." >&2

  cd /home/user/plenya 2>/dev/null || true

  SQL="SELECT
    COALESCE(authors, 'N/A'),
    title,
    COALESCE(journal, 'N/A'),
    COALESCE(EXTRACT(YEAR FROM publish_date)::text, 'N/A'),
    COALESCE(pmid, 'N/A'),
    LEFT(COALESCE(abstract,''), 400)
  FROM articles
  WHERE title ILIKE '%${QUERY}%' OR abstract ILIKE '%${QUERY}%'
  ORDER BY publish_date DESC NULLS LAST
  LIMIT ${LIMIT};"

  FALLBACK=$(docker compose exec -T db psql -U plenya_user -d plenya_db \
    --no-align --field-separator='|' -t -c "$SQL" 2>/dev/null)

  if [ -n "$FALLBACK" ]; then
    echo "$FALLBACK" | while IFS='|' read -r authors title journal year pmid abstract; do
      printf -- "---\nTítulo: %s\nAutores: %s\nJournal: %s\nAno: %s\nPMID: %s\nAbstract: %s\n\n" \
        "$title" "$authors" "$journal" "$year" "$pmid" "$abstract"
    done
    echo "=== Resultados via PostgreSQL ===" >&2
  else
    echo "⚠️  RAG indisponível (API + banco inacessíveis)." >&2
  fi
fi
