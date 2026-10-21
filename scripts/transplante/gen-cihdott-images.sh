#!/usr/bin/env bash
# Fotografias da aula "Transplante renal: a volta inteira" (CIHDOTT, 22/10/2026).
# Gente brasileira de verdade, emoção, luz cinematográfica; nunca texto na imagem.
# Uso: gen-cihdott-images.sh            → gera todas que faltam (6 em paralelo)
#      gen-cihdott-images.sh nome ...   → refaz só essas
# Saída: ~/lecture-output/transplante-rede-cihdott/images/<nome>.png
set -euo pipefail
. "$(dirname "${BASH_SOURCE[0]}")/../lib/openai-image.sh"
[ -n "${OPENAI_API_KEY:-}" ] || OPENAI_API_KEY="$(grep -E '^OPENAI_API_KEY=' /home/user/plenya/apps/api/.env | cut -d= -f2-)"
export OPENAI_API_KEY
OUT_DIR="/home/user/lecture-output/transplante-rede-cihdott/images"; mkdir -p "$OUT_DIR"

STYLE=" Cinematic editorial photograph, documentary realism, as if shot by a great photojournalist on a 35mm or 50mm lens. Authentic Brazilian people of mixed ethnicities, real ordinary faces, natural expressions, real emotion, nobody posing for the camera. Brazilian public hospital or community setting, clean and modern but modest, never luxurious. Beautiful motivated light, rich natural colour, deep but soft shadows, shallow depth of field, fine film grain. Compose with clean, uncluttered negative space on the LEFT third of the frame so a title can sit there. Absolutely NO text anywhere: no words, letters, numbers, logos, signage, name badges with writing, labels or readable screens; any monitor is switched off or completely out of focus. Photorealistic, not illustration, not 3D render, not stock-photo cliché."

declare -A P
P[cihdott]="In a quiet hospital family room at night, a Brazilian ICU nurse in navy scrubs sits close to a grieving middle-aged couple on a sofa, holding the mother's hands, speaking softly; the father looks down; warm lamp light against a dark blue window; compassion, gravity, respect"
P[telefone]="Before dawn in a simple Brazilian bedroom, a 52-year-old Brazilian man sitting on the edge of his bed answers a mobile phone held to his ear, his other hand over his mouth in disbelief and hope; his wife behind him wakes up and touches his shoulder; a single bedside lamp, deep blue window light; a hemodialysis arteriovenous fistula bandage visible on his forearm"
P[ubs]="Morning in a Brazilian public primary care unit (UBS) consultation room: a female family doctor in a white coat measures the blood pressure of a 38-year-old Brazilian man in a work shirt; simple desk, glucometer, a window with tropical plants; warm daylight; ordinary, caring, attentive"
P[ubs-rua]="Exterior of a modest Brazilian public primary health care unit building in a residential neighbourhood of a small city in Paraná at early morning: low white building with a covered entrance, a few patients arriving, a community health agent in a light blue vest walking, trees and red earth sidewalk, golden sunrise light, long shadows; no signs or text on the building"
P[dialise]="A 50-year-old Brazilian man sits in a reclining chair during hemodialysis in a bright dialysis unit, blood lines connected to his forearm fistula, the dialysis machine beside him out of focus with dark screen, he looks out of the window with patient tired eyes; rows of other chairs softly blurred behind; late afternoon light"
P[irmaos]="A Brazilian woman of 34 and her older brother of 38 sit side by side in a hospital outpatient waiting area, the brother has his arm around her shoulders and she leans her head on him, both smiling softly with nervous hope; daylight from large windows; siblings resemblance"
P[equipe]="A multiprofessional Brazilian transplant team in discussion around a table in a bright meeting room: a nephrologist, a surgeon in scrubs, a nurse, a psychologist, a pharmacist and a nutritionist, men and women of different ages and ethnicities, papers and a closed laptop, engaged conversation, one pointing to a document; natural light"
P[caixa]="Motion shot in a long hospital service corridor at dawn: a Brazilian transplant coordinator nurse in scrubs walks fast carrying a plain white medical transport cooler box with no writing, a second professional beside her opening a door; motion blur, cool blue corridor light mixed with warm sunrise from the far end; urgency"
P[cirurgia]="Inside a modern operating room during a kidney transplant: three Brazilian surgeons in blue gowns, caps and masks bent over the patient under a bright round surgical lamp, a scrub nurse passing an instrument; dramatic pool of light in a darker room, teal tones, intense concentration; no blood visible, monitors out of focus"
P[uti]="Night in an intensive care unit: a Brazilian female nurse in scrubs kneels at the side of a hospital bed checking and measuring a urine collection bag with a graduated container, a post-operative middle-aged man resting in the bed; warm low light over the bed, cool blue monitors completely out of focus in the background; vigilance and tenderness"
P[mao]="Close, emotional moment in a hospital room: a Brazilian nurse holds the hand of a worried 52-year-old male patient in bed while his wife sits beside him, the nurse is calm and reassuring, eye contact; soft morning light through a window; empathy"
P[farmacia]="A Brazilian clinical pharmacist in a white coat sits with a transplant patient and his wife at a small table in a hospital room, explaining a weekly pill organizer and medication boxes with no readable labels, the patient attentive; warm daylight"
P[nutricao]="A young male Brazilian nutritionist with short hair and glasses talks with a smiling middle-aged male transplant patient at a table with a colourful plate of rice, beans, cooked vegetables and grilled chicken and a bowl of fruit; bright friendly hospital consultation room, daylight"
P[psicologia]="A Brazilian psychologist, woman in her forties, listens with full attention to a young adult woman patient sitting across from her in two armchairs in a calm hospital room with a plant; the patient speaks with emotion; soft side light; trust"
P[social]="A grey-haired male Brazilian social worker in a cardigan sits beside an elderly patient and his daughter in a hospital corridor bench, helping them with documents in a folder, the daughter nodding; warm natural light; practical care, kindness"
P[alta]="Hospital discharge day: a 52-year-old Brazilian man walks out through the glass doors of a hospital into bright sunshine, carrying a small bag, his wife holding his arm, a nurse at the door waving goodbye; joy, relief, back light flare"
P[acs]="A Brazilian community health agent woman in a light blue vest stands at the front gate of a simple colorful house in a neighbourhood of Paraná, talking warmly with a middle-aged woman and her teenage son on the porch, bougainvillea, morning sun; community, trust"
P[um-ano]="One year later: a 53-year-old Brazilian man laughing as he plays with his small granddaughter in the backyard of a simple house, red soil and a mango tree, family lunch table in the background, late afternoon golden light; life restored"
P[maos]="Top-down close shot of many hands of a Brazilian healthcare team stacked together in the center, different skin tones, some with scrub sleeves in different colours (navy, green, white coat, blue vest), warm light, celebration of teamwork"
P[anatomia]="Elegant, accurate medical illustration of a kidney transplant, anterior view of the lower abdomen and pelvis in a refined anatomical textbook style: the two native kidneys remain in place high in the back, the transplanted kidney sits low in the right iliac fossa, its renal artery joined to the right external iliac artery, its renal vein joined to the right external iliac vein, and its ureter joined to the top of the bladder. Soft cream background, delicate watercolour shading in muted red, blue, gold and sage, no labels, no text, no letters, no arrows"

names=("$@"); if [ ${#names[@]} -gt 0 ]; then FORCE=1; else names=("${!P[@]}"); fi
gen(){ local n="$1" out="$OUT_DIR/$1.png" style="$STYLE"
  [ "$n" = anatomia ] && style=" Medical textbook illustration, refined and accurate. Absolutely NO text, labels, letters or numbers."
  [ -s "$out" ] && [ -z "${FORCE:-}" ] && { echo "skip $n"; return; }
  local tmp; tmp="$(mktemp)"
  local payload; payload=$(jq -n --arg m "$OPENAI_IMAGE_MODEL" --arg p "${P[$n]}.${style}" '{model:$m,prompt:$p,size:"1536x1024",quality:"xhigh",n:1}')
  local code; code=$(curl -s --max-time 400 -o "$tmp" -w "%{http_code}" https://api.openai.com/v1/images/generations \
    -H "Authorization: Bearer $OPENAI_API_KEY" -H "Content-Type: application/json" -d "$payload")
  if [ "$code" = 200 ]; then jq -r '.data[0].b64_json' "$tmp" | base64 -d > "$out"; echo "ok $n"; else echo "ERRO $code $n: $(head -c 300 "$tmp")"; fi
  rm -f "$tmp"; }
i=0; for n in "${names[@]}"; do gen "$n" & i=$((i+1)); [ $((i % 6)) -eq 0 ] && wait; done; wait
