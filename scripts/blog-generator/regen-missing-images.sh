#!/usr/bin/env bash
# Regenerate the blog images that failed during the initial run because
# the OpenAI org hit billing_hard_limit_reached. After topping up at
# https://platform.openai.com/settings/organization/limits, run this:
#
#   bash /home/user/plenya/scripts/blog-generator/regen-missing-images.sh
#
# Each call is idempotent — already-existing files are skipped (rm them
# first if you want to regenerate). All prompts run in parallel batches
# of 4 to avoid hammering the rate limit.

set -u
GEN="/home/user/plenya/scripts/blog-generator/gen-image.sh"
BLOG_IMG_ROOT="/home/user/plenya/apps/site/public/images/blog"

# Each entry: slug | kind | "english prompt"
ITEMS=(
  '12-exames-que-valem-cada-centavo-e-12-que-sao-desperdicio|hero|A weathered wooden physicians desk top-down, with a stack of laboratory request forms tied with twine on one side and a single pencil resting across them, golden hour window light, conveying the careful curation of which tests truly matter, generous negative space upper right.'
  '12-exames-que-valem-cada-centavo-e-12-que-sao-desperdicio|inline|A vintage analog scale with two brass pans, perfectly balanced; on one pan a small stack of coins, on the other a folded paper, on a dark walnut surface with soft shadows, evoking weighing what is worth investing in.'

  'apneia-em-quem-nao-ronca-forma-feminina-invisivel|hero|A single soft linen pillow on dark blue bedsheets at pre-dawn, the only light source a faint cool blue from a window, a partially visible silhouette of a sleeping form barely discernible, conveying invisible disturbance and the silence of unrecognized illness.'
  'apneia-em-quem-nao-ronca-forma-feminina-invisivel|inline|A polished medical tuning fork resting on a folded gauze square next to a small clinical notebook with handwritten time markings, top-down on a slate surface, evoking the precision of finding a hidden signal.'

  'ferritina-30-a-100-normalidade-que-esgota-mulheres|inline|A close-up of a single iron nail and a sprig of fresh spinach beside a small empty glass test tube on a dark linen napkin, top-down, evoking iron stores quietly emptying.'

  'healthspan-vs-lifespan-comprimir-a-morbidade|hero|Two parallel ribbons of fabric on a wooden surface — one long and worn near its end, the other slightly shorter but uniformly intact through its length — illustrating the difference between length of life and length of healthy life, soft directional light from upper left.'
  'healthspan-vs-lifespan-comprimir-a-morbidade|inline|An hourglass with sand mid-fall on a polished oak surface, photographed slightly off-axis, soft warm light, evoking the deliberate compression of decline rather than the lengthening of suffering.'

  'hrv-variabilidade-cardiaca-termometro-do-sistema-nervoso|hero|An abstract macro photograph of soft concentric ripples across still water in a dark stone basin, lit from a single warm side light, conveying the rhythm and variation of an autonomic system, no human elements.'
  'hrv-variabilidade-cardiaca-termometro-do-sistema-nervoso|inline|A vintage seismograph paper strip showing irregular wave patterns, photographed top-down on a dark surface with a brass pen resting at the end, evoking measurement of subtle internal rhythm.'

  'lipoproteina-a-exame-que-cardiologista-nao-pediu|hero|A single slightly curled paper request form for a laboratory test on a dark wooden physicians desk, partially in shadow, with a fountain pen unscrewed beside it, conveying the test that should have been ordered but rarely is, generous negative space upper portion.'
  'lipoproteina-a-exame-que-cardiologista-nao-pediu|inline|A delicate strand of double-helix DNA carved from gold-toned brass on a black velvet surface, lit from one side, evoking inherited cardiovascular risk hidden in genetics.'

  'luz-da-manha-remedio-gratuito-que-reseta-metabolismo|hero|A wide horizontal composition: warm sunrise light pouring through tall east-facing windows onto an empty wooden floor, dust motes visible in the beams, no human figures, conveying the gift of morning light and circadian reset.'
  'luz-da-manha-remedio-gratuito-que-reseta-metabolismo|inline|A small antique brass sundial casting a sharp morning shadow on weathered stone, top-down view, conveying the precise dose of morning light that calibrates the day.'

  'meditacao-que-funciona-em-8-minutos|hero|A small ceramic bowl of clear water on a dark wooden bench at dawn, surface perfectly still, a single soft beam of warm light entering from upper left, conveying mental quietude and the simplicity of a daily practice.'
  'meditacao-que-funciona-em-8-minutos|inline|A vintage brass pocket watch open with second hand frozen, resting on a folded linen cloth, top-down, conveying the small but precise dose of time that produces measurable effect.'

  'pre-diabetes-nao-e-fase-e-janela-de-5-anos|hero|A partially open wooden door at the end of a dim hallway with golden afternoon light streaming through the gap, conveying a closing window of opportunity, no human figures, generous negative space.'
  'pre-diabetes-nao-e-fase-e-janela-de-5-anos|inline|A glass jar half-filled with refined white sugar tipped sideways on a dark slate surface, a few crystals spilled out forming a small mound, top-down, evoking the slow accumulation that tips into disease.'

  'proteina-meta-que-voce-nao-esta-atingindo|hero|A still-life arrangement on a rough linen napkin, top-down: three whole eggs, a portion of grilled chicken, a small bowl of cottage cheese, and a measuring tape coiled to one side, in warm muted tones, evoking quantification of dietary protein.'
  'proteina-meta-que-voce-nao-esta-atingindo|inline|A simple analog kitchen scale with a single piece of grilled meat on its plate, the dial pointing to a precise number, top-down on a wood surface with morning light, conveying the discipline of measurement.'

  'quatro-assassinos-silenciosos-depois-dos-40|hero|Four identical hourglasses arranged in a row on a dark walnut surface, each with sand in different stages of falling, photographed in muted warm light, evoking the four chronic disease processes that progress silently in parallel.'
  'quatro-assassinos-silenciosos-depois-dos-40|inline|A row of four small unlit candles on a stone shelf, equally spaced, with subtle smoke trails as if recently extinguished, evoking the silent diseases that arrive without announcement.'

  'solidao-como-fator-de-risco-cardiovascular|inline|A single empty wooden chair beside a small round table with one cup of tea, in a quiet sunlit room with morning light, photographed from a low angle, conveying the cardiovascular weight of social isolation.'

  'sono-energia-longevidade|inline|A simple ceramic cup of dark coffee untouched on a wooden bedside table next to a closed book, soft cool morning light through curtains, evoking the failure of stimulants to substitute for restorative sleep.'

  'sono-que-nao-recupera|inline|A folded handwritten sleep diary on rumpled linen sheets in early morning light, with circles drawn around several timestamps in the night, evoking the data of fragmented rest.'

  'treinar-muito-e-envelhecer-errado|hero|A pair of well-worn running shoes left abandoned on a stone path at sunset, only one shoe in clear focus the other slightly out of frame, conveying the limits of repetitive volume training, generous negative space upper portion.'
  'treinar-muito-e-envelhecer-errado|inline|A coiled jump rope on a wooden gym floor with a folded towel beside it and a half-empty water bottle, top-down, lit by warm afternoon light, evoking the recovery half of training that is usually neglected.'

  'treinar-para-envelhecer-bem-zona-2-e-forca|hero|A heart rate monitor watch face displaying a steady mid-zone reading, resting on a folded gym towel beside a stainless steel water bottle on a wooden bench, in warm directional light, evoking measured controlled training.'
  'treinar-para-envelhecer-bem-zona-2-e-forca|inline|A barbell loaded with two modest plates resting on a wooden gym platform, soft afternoon light from upper left, no human elements, conveying the simplicity and centrality of resistance training.'
)

run_one() {
  local item="$1"
  IFS='|' read -r slug kind prompt <<<"$item"
  local out="$BLOG_IMG_ROOT/$slug/$kind.webp"
  if [ -f "$out" ]; then
    echo "SKIP  $slug/$kind.webp (exists)"
    return 0
  fi
  if "$GEN" "$slug" "$kind" "$prompt" >/dev/null 2>&1; then
    echo "OK    $slug/$kind.webp"
  else
    echo "FAIL  $slug/$kind.webp"
    return 1
  fi
}

export -f run_one
export GEN BLOG_IMG_ROOT

printf "%s\n" "${ITEMS[@]}" | xargs -I{} -P 4 bash -c 'run_one "$@"' _ {}

# Post-step: 3 MDX files had cover/inline refs stripped by the original agent
# because the images had failed mid-run. If their images now exist, patch the
# MDX to point at them.
PATCH_DIR="/home/user/plenya/apps/site/content/blog/pt"
patch_mdx() {
  local slug="$1" inline_caption="$2"
  local file="$PATCH_DIR/$slug.mdx"
  local hero="$BLOG_IMG_ROOT/$slug/hero.webp"
  local inline="$BLOG_IMG_ROOT/$slug/inline.webp"
  [ -f "$hero" ] && [ -f "$inline" ] || { echo "PATCH SKIP $slug (images missing)"; return 0; }
  grep -q "^cover:" "$file" || sed -i "/^tags:/a cover: /images/blog/$slug/hero.webp" "$file"
  grep -q "/images/blog/$slug/inline.webp" "$file" || {
    # Insert before the last H2 (## ...) in the body — common "Plenya bridge" / closing section
    awk -v slug="$slug" -v cap="$inline_caption" '
      BEGIN { last_h2 = 0 }
      /^## / { last_h2 = NR }
      { lines[NR] = $0 }
      END {
        for (i = 1; i <= NR; i++) {
          if (i == last_h2) {
            print "![" cap "](/images/blog/" slug "/inline.webp)"
            print ""
          }
          print lines[i]
        }
      }
    ' "$file" > "$file.tmp" && mv "$file.tmp" "$file"
  }
  echo "PATCH OK   $slug"
}
patch_mdx "luz-da-manha-remedio-gratuito-que-reseta-metabolismo" "A dose precisa de luz matinal — o relógio interno que define o resto do dia"
patch_mdx "meditacao-que-funciona-em-8-minutos" "O tempo medido — pequena dose, efeito mensurável"
patch_mdx "apneia-em-quem-nao-ronca-forma-feminina-invisivel" "A diretriz silenciosa que escapa do protocolo padrão"

echo
echo "Done. Inventory:"
for d in "$BLOG_IMG_ROOT"/*/; do
  slug=$(basename "$d")
  n=$(ls "$d" 2>/dev/null | wc -l)
  echo "  $n  $slug"
done

echo
echo "Validate MDX:"
cp /home/user/plenya/scripts/blog-generator/validate-blog.mjs /home/user/plenya/apps/site/check_blog.mjs \
  && cd /home/user/plenya/apps/site && node check_blog.mjs && rm -f check_blog.mjs
