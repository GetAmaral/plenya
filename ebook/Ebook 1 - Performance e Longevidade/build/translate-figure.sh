#!/usr/bin/env bash
# Translate an ebook PT figure to EN via gpt-image-2 /v1/images/edits.
# Preserves layout, colors, fonts, icons, all numbers; only translates text.
#
# Usage: translate-figure.sh "<filename>" "<text-replacement-prompt>"
#   filename: e.g. "Cap07 Fig01.PNG" (must exist in figuras/pt-BR/)
#   prompt: explicit PT→EN substitution list
#
# Output: figuras/en/<filename>
set -euo pipefail

FILENAME="${1:?filename required (e.g., 'Cap07 Fig01.PNG')}"
PROMPT="${2:?prompt required}"

EBOOK_DIR="/home/user/plenya/ebook/Ebook 1 - Performance e Longevidade"
SRC="$EBOOK_DIR/figuras/pt-BR/$FILENAME"
OUT="$EBOOK_DIR/figuras/en/$FILENAME"

[ -f "$SRC" ] || { echo "source not found: $SRC" >&2; exit 1; }

if [ -z "${OPENAI_API_KEY:-}" ]; then
  OPENAI_API_KEY="$(grep -E '^OPENAI_API_KEY=' /home/user/plenya/apps/api/.env | cut -d= -f2-)"
fi

INSTRUCTION="Task: translate the inline Portuguese text in this figure to US English using the EXACT English strings provided below.

CRITICAL RULES (read carefully):
1. Use the English strings I give you VERBATIM. Copy each English string EXACTLY as written. Do NOT paraphrase. Do NOT rephrase. Do NOT use synonyms.
2. If I write 'raises death risk', do NOT write 'increases the risk of death'. Use my words.
3. If I write 'Compared with', do NOT write 'Compared to'. Use my words.
4. Do NOT add any English word that I did not write. If I do not include 'behavior', do not add 'behavior'. If I do not include 'no risk factors', do not add 'no risk factors'.
5. Do NOT remove any English word I provided.
6. If I instruct you to REMOVE a parenthetical (e.g., '(do coração)'), do NOT keep its English equivalent. Drop it entirely.
7. Use the Oxford comma EXACTLY where I write it (e.g., 'A, B, or C' — keep the comma before 'or').
8. Use period (.) for English decimals: '5,0x' becomes '5.0x'. Use comma (,) for English thousands: '122.007' becomes '122,007'.

PRESERVE EXACTLY (do not change any of these):
- Layout, every panel/element position
- Every color, every icon, every chart shape
- EVERY numeric value (digit-for-digit; only change comma↔period as decimal/thousands separator)
- Typography style, font, font weight
- Background color and texture
- Proportions, spacing, alignment

DO NOT:
- Add or remove any visual element (badges, legends, icons, lines)
- Invent words that are not in the English strings I gave you
- Translate proper names: keep Portuguese names of patients (Paulo, Ricardo, Marcos, Fernanda, Ana, André, Patrícia) UNCHANGED
- Translate study citations (e.g., 'Mandsager et al.', 'JAMA Network Open, 2018') — keep verbatim except for the literal label 'Fonte:' which becomes 'Source:'

PT→EN substitutions to apply:
${PROMPT}"

TMP_JSON="$(mktemp)"
trap 'rm -f "$TMP_JSON"' EXIT

HTTP_CODE=$(curl -s -o "$TMP_JSON" -w "%{http_code}" \
  https://api.openai.com/v1/images/edits \
  -H "Authorization: Bearer $OPENAI_API_KEY" \
  -F model="gpt-image-2" \
  -F image="@${SRC}" \
  -F "prompt=${INSTRUCTION}" \
  -F size="auto" \
  -F quality="high")

if [ "$HTTP_CODE" != "200" ]; then
  echo "OpenAI API error ($HTTP_CODE):" >&2
  head -c 1500 "$TMP_JSON" >&2
  exit 1
fi

jq -r '.data[0].b64_json' "$TMP_JSON" | base64 -d > "$OUT"
[ -s "$OUT" ] || { echo "empty output file" >&2; exit 1; }

echo "saved → $OUT"
