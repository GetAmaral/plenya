#!/usr/bin/env bash
# Regenera TODAS as imagens da palestra "Cuidados com o Paciente Transplantado Renal".
# Os prompts ficam aqui para que qualquer imagem possa ser refeita igual, ou ajustada
# num lugar só. Saída: docs/transplante/images/<nome>.webp
#
# Uso:  gen-all-images.sh            regenera todas
#       gen-all-images.sh 14-uti     regenera só uma
#
# A anatomia (05-anatomia.svg) NÃO está aqui de propósito: é desenho vetorial feito à
# mão. Modelo de imagem por IA não acerta topografia anatômica, e um esquema errado
# numa aula de enfermagem é pior que nenhum.
set -euo pipefail
G="$(dirname "$0")/gen-slide-image.sh"
ONLY="${1:-}"

run(){ [ -z "$ONLY" ] || [ "$ONLY" = "$1" ] || return 0; "$G" "$@"; }

run 04-caso "A dark bedside table at 2am lit only by the glow of a mobile phone that is ringing, face up beside a glass of water and a folded pair of reading glasses. The bed is unmade in the background, out of focus. The mood is the middle of the night, a call that changes everything. Empty space on the left third of the frame"

run 10-medicacao "A weekly pill organiser lying open on a pale kitchen counter beside two amber prescription bottles with completely blank unlabelled sides, a few white tablets loose on the surface, and a glass of water. Morning light from a window on the right. Quiet domestic order, the discipline of taking medicine at a fixed hour"

run 13-preparo "An empty hospital gurney with folded white linen waiting in a quiet pre-operative corridor, double doors to the operating theatre closed at the end of the hall, polished floor reflecting early morning light from a high window. Nobody in frame"

run 14-uti "Close view of a stack of three infusion pumps on a pole in an intensive care bay, their displays switched off and dark, clear tubing looped and clipped in tidy order, a pale blue blanket edge and bed rail out of focus behind. Cool clinical light"

run 16-enfermaria "A tidy hospital ward room in full afternoon daylight, bed freshly made, a walking frame parked beside it and an armchair by the window with a folded blanket. Warm light across the floor, a sense of someone about to get up and walk. Nobody in frame"

run 20-conversa "Two cups of coffee cooling on a worn wooden kitchen table seen from above and slightly to the side, four hands resting near them belonging to two people sitting across from each other, faces out of frame. An ordinary family conversation in afternoon light"
