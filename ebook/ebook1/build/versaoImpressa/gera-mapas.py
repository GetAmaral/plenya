#!/usr/bin/env python3
"""Gera o bloco MAPS de build-figuras-cor.py a partir da medição.

Para cada figura e cada tom de cinza, adota a cor medida na arte original
(amostra de maior pureza). Descarta o que claramente não é preenchimento:
branco de fundo, preto de tipografia, e amostras que caíram no fundo.

O resultado é uma proposta — a conferência visual figura a figura vem depois
e sobrescreve o que estiver errado via OVERRIDES em build-figuras-cor.py.
"""
import importlib.util
import sys
from pathlib import Path

import numpy as np
from scipy import ndimage

spec = importlib.util.spec_from_file_location("ac", "analise-cor.py")
ac = importlib.util.module_from_spec(spec)
sys.modules["ac"] = ac
spec.loader.exec_module(ac)

AREA_MIN = 0.10      # % da figura — abaixo disso é traço/régua, não preenchimento


def medir(nome):
    pdf = ac.BW / f"{nome}.pdf"
    png = ac.ORIG / f"{nome.replace('_', ' ')}.PNG"
    if not (pdf.exists() and png.exists()):
        return {}
    from PIL import Image
    orig = np.array(Image.open(png).convert("RGB"))
    bbox_orig = ac.bbox_conteudo(orig)
    bbox_fig = ac.bbox_conteudo(np.dstack([ac.render_normal(pdf)] * 3))
    out = {}
    for tom in sorted(ac.tons(pdf), reverse=True):
        if tom >= 0.999 or tom <= 0.01:
            continue
        mask = ac.isola(pdf, tom) < 128
        area = mask.sum() / mask.size * 100
        if area < AREA_MIN or mask.sum() < 40:
            continue
        lbl, n = ndimage.label(mask)
        if not n:
            continue
        sizes = ndimage.sum(mask, lbl, range(1, n + 1))
        melhor = None
        for i in np.argsort(sizes)[::-1][:4]:
            if sizes[i] < 30:
                continue
            (r, g, b), pureza = ac.amostra_regiao(orig, mask, lbl, i, bbox_orig, bbox_fig)
            if (r, g, b) == (248, 248, 248):      # caiu no fundo branco
                continue
            if melhor is None or pureza > melhor[1]:
                melhor = ((r, g, b), pureza)
        if melhor is None:
            continue
        (r, g, b), pureza = melhor
        if pureza < 0.30:
            continue
        out[tom] = (f"#{r:02x}{g:02x}{b:02x}", area, pureza)
    return out


if __name__ == "__main__":
    alvos = sys.argv[1:] or [p.stem for p in sorted(ac.BW.glob("Cap*.pdf"))]
    print("MEDIDO = {")
    for nome in alvos:
        m = medir(nome)
        if not m:
            print(f'    "{nome}": {{}},')
            continue
        itens = ", ".join(f'{t}: "{c}"' for t, (c, a, p) in sorted(m.items(), reverse=True))
        print(f'    "{nome}": {{{itens}}},')
        for t, (c, a, p) in sorted(m.items(), reverse=True):
            print(f"        # tom {t:.2f} → {c}  (área {a:.2f}%, pureza {p*100:.0f}%)",
                  file=sys.stderr)
    print("}")
