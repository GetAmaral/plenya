#!/usr/bin/env python3
"""Gera o radar AGIR do slide do Escore com a MESMA geometria do componente da casa
(packages/ui/src/score-radar/RadarAgir.tsx): letras como arcos externos com largura
proporcional ao nº de pilares (piso 10°), A no topo, ponto por pilar em raio = score/100,
polígono dourado ligando os pontos. Pilares vêm de apps/site/lib/agir-structure.ts.

Saída: fragmento SVG em stdout, para colar no slide.
"""
import json, math, sys

VIEWBOX, CX, CY, RMAX = 400, 200, 200, 150
ARC_R = RMAX + 18
MIN_LETTER_DEG = 10.0
POLY = "#b38645"
INK_12, INK_35 = "rgba(6,59,79,0.12)", "rgba(6,59,79,0.35)"
COLORS = {"A": "#2b9c6d", "G": "#c9822f", "I": "#b85e93", "R": "#4585bc"}  # tons dark, fundo petrol
NAMES = {"A": "Atividade", "G": "Gestão", "I": "Integração", "R": "Ritmo"}

def polar(a, r):
    rad = math.radians(a - 90)
    return round(CX + r * math.cos(rad), 1), round(CY + r * math.sin(rad), 1)

def arc_path(r, a0, a1):
    x0, y0 = polar(a0, r); x1, y1 = polar(a1, r)
    large = 1 if (a1 - a0) % 360 > 180 else 0
    return f"M {x0} {y0} A {r} {r} 0 {large} 1 {x1} {y1}"

def build(pillars_by_letter, scores):
    letters = ["A", "G", "I", "R"]
    total = sum(len(pillars_by_letter[l]) for l in letters)
    remaining = 360 - len(letters) * MIN_LETTER_DEG
    per = remaining / total
    span = lambda vp: MIN_LETTER_DEG + vp * per
    inner = MIN_LETTER_DEG / 2
    rot = -span(len(pillars_by_letter["A"])) / 2

    pts, arcs, cursor = [], [], 0.0
    for l in letters:
        ps = pillars_by_letter[l]
        s0, s1 = cursor + rot, cursor + span(len(ps)) + rot
        for i, nome in enumerate(ps):
            ang = s0 + inner + per * (i + 0.5)
            sc = max(0, min(100, scores[l][i]))
            pts.append((nome, l, ang, *polar(ang, sc / 100 * RMAX)))
        arcs.append((l, s0, s1))
        cursor += span(len(ps))

    o = [f'<svg viewBox="0 0 {VIEWBOX} {VIEWBOX}" style="width:100%">']
    for f in (0.25, 0.5, 0.75, 1.0):   # anéis
        o.append(f'<circle cx="{CX}" cy="{CY}" r="{RMAX*f:.0f}" fill="none" stroke="rgba(234,231,218,.13)" stroke-width="1"/>')
    for _, l, ang, x, y in pts:        # raios
        ex, ey = polar(ang, RMAX)
        o.append(f'<line x1="{CX}" y1="{CY}" x2="{ex}" y2="{ey}" stroke="rgba(234,231,218,.08)" stroke-width="1"/>')
    o.append('<polygon points="' + " ".join(f"{x},{y}" for *_, x, y in pts) +
             f'" fill="{POLY}" fill-opacity=".30" stroke="{POLY}" stroke-width="2.5" stroke-linejoin="round"/>')
    for _, l, _, x, y in pts:
        o.append(f'<circle cx="{x}" cy="{y}" r="3.4" fill="{COLORS[l]}"/>')
    for l, a0, a1 in arcs:             # arcos das letras
        o.append(f'<path d="{arc_path(ARC_R, a0+1.25, a1-1.25)}" fill="none" stroke="{COLORS[l]}" stroke-width="9" stroke-linecap="round"/>')
        mx, my = polar((a0 + a1) / 2, ARC_R + 26)
        o.append(f'<text x="{mx}" y="{my+7}" text-anchor="middle" fill="{COLORS[l]}" '
                 f'style="font-family:Fraunces,serif;font-size:26px">{l}</text>')
    o.append('</svg>')
    return "\n".join(o)

if __name__ == "__main__":
    dados = json.load(open(sys.argv[1], encoding="utf-8"))
    print(build(dados["pillars"], dados["scores"]))
