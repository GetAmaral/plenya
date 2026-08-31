#!/usr/bin/env python3
"""Propõe, por medição, o mapa tom-de-cinza → cor de cada figura vetorial.

Ideia: a figura vetorial B&W é um redesenho da arte colorida original, com
layout aproximadamente correspondente. Então, para cada tom de cinza:

  1. renderiza a figura B&W isolando SÓ aquele tom (pinta o tom de preto e
     todo o resto de branco) → máscara das regiões que aquele tom preenche;
  2. pega as maiores regiões conectadas dessa máscara e o centroide de cada;
  3. converte o centroide para coordenada relativa (0..1);
  4. amostra a arte original naquela mesma coordenada relativa, numa janela
     pequena, e tira a cor dominante;
  5. reporta tom → cor medida, com quanta área cada tom cobre.

O resultado é uma PROPOSTA para conferência visual, não uma verdade — se o
redesenho deslocou um bloco, o centroide cai no lugar errado e isso aparece
na conferência.

Uso: analise-cor.py Cap14_Fig02 [...]   (sem argumento = todas)
"""
import re
import subprocess
import sys
import tempfile
from collections import Counter
from pathlib import Path

import numpy as np
import pikepdf
from PIL import Image
from scipy import ndimage

HERE = Path(__file__).resolve().parent
BW = HERE / "figuras-bw"
ORIG = HERE.parent.parent / "figuras" / "pt-BR"


def tons(pdf_path):
    """Tons de cinza usados, com contagem de operadores."""
    pdf = pikepdf.open(pdf_path)
    c = b"".join(bytes(p.Contents.read_bytes()) for p in pdf.pages)
    vals = Counter()
    for m in re.finditer(rb"([\d.]+)\s+(g|G)\b", c):
        vals[round(float(m.group(1)), 2)] += 1
    for m in re.finditer(rb"([\d.]+\s+[\d.]+\s+[\d.]+)\s+(rg|RG)\b", c):
        p = [float(x) for x in m.group(1).split()]
        if max(p) - min(p) <= 0.04:
            vals[round(sum(p) / 3, 2)] += 1
    return vals


def render_normal(pdf_path, dpi=100):
    """Render simples da figura, sem isolar tom — serve para achar a caixa."""
    with tempfile.TemporaryDirectory() as td:
        subprocess.run(["pdftoppm", "-r", str(dpi), "-png", "-singlefile",
                        str(pdf_path), str(Path(td) / "o")], capture_output=True)
        return np.array(Image.open(Path(td) / "o.png").convert("L"))


def isola(pdf_path, tom, dpi=100):
    """Renderiza a figura com `tom` em preto e todo o resto em branco."""
    pdf = pikepdf.open(pdf_path)

    def sub_g(m):
        v, op = round(float(m.group(1)), 2), m.group(2).decode()
        alvo = "0" if v == tom else "1"
        return f"{alvo} {op}".encode()

    def sub_rgb(m):
        p = [float(x) for x in m.group(1).split()]
        op = m.group(2).decode()
        if max(p) - min(p) > 0.04:
            return b"1 1 1 " + op.encode()
        v = round(sum(p) / 3, 2)
        alvo = "0 0 0" if v == tom else "1 1 1"
        return f"{alvo} {op}".encode()

    for page in pdf.pages:
        d = bytes(page.Contents.read_bytes())
        d = re.sub(rb"([\d.]+)\s+(g|G)\b", sub_g, d)
        d = re.sub(rb"([\d.]+\s+[\d.]+\s+[\d.]+)\s+(rg|RG)\b", sub_rgb, d)
        page.Contents.write(d)
    with tempfile.TemporaryDirectory() as td:
        p = Path(td) / "iso.pdf"
        pdf.save(p)
        subprocess.run(["pdftoppm", "-r", str(dpi), "-png", "-singlefile", str(p),
                        str(Path(td) / "out")], capture_output=True)
        return np.array(Image.open(Path(td) / "out.png").convert("L"))


def bbox_conteudo(arr):
    """Caixa da área com tinta (tudo que não é branco de fundo)."""
    g = arr.mean(axis=2) if arr.ndim == 3 else arr
    tinta = g < 246
    ys, xs = np.where(tinta)
    if not len(ys):
        return 0, 0, g.shape[0], g.shape[1]
    return ys.min(), xs.min(), ys.max() + 1, xs.max() + 1


def amostra_regiao(orig, mask, lbl, idx, _bbox_orig, _bbox_fig, n_pontos=400, rng=None):
    """Cor dominante da arte original sob a região `idx` da máscara.

    Amostra MUITOS pontos dentro da região (erodida, para fugir das bordas) em
    vez de só o centroide: se o redesenho deslocou um pedaço, ou se o centroide
    calha de cair sobre um texto, um ponto só mente. Descarta pixels quase
    pretos e quase brancos (tipografia e fundo) antes de tirar a moda.
    """
    reg = lbl == idx + 1
    ero = ndimage.binary_erosion(reg, iterations=2)
    if ero.sum() < 20:
        ero = reg
    ys, xs = np.where(ero)
    if rng is None:
        rng = np.random.default_rng(7)
    if len(ys) > n_pontos:
        sel = rng.choice(len(ys), n_pontos, replace=False)
        ys, xs = ys[sel], xs[sel]
    # Alinha os dois pela caixa de conteúdo: o redesenho e a arte original têm
    # margens diferentes, então a fração sobre a tela cheia cai no lugar errado.
    my0, mx0, my1, mx1 = _bbox_fig
    oy0, ox0, oy1, ox1 = _bbox_orig
    fy = (ys - my0) / max(my1 - my0, 1)
    fx = (xs - mx0) / max(mx1 - mx0, 1)
    py = np.clip((oy0 + fy * (oy1 - oy0)).astype(int), 0, orig.shape[0] - 1)
    px = np.clip((ox0 + fx * (ox1 - ox0)).astype(int), 0, orig.shape[1] - 1)
    px_vals = orig[py, px]
    lum = px_vals.mean(axis=1)
    keep = (lum > 40) & (lum < 250)          # fora tipografia preta e branco puro
    if keep.sum() < 10:
        keep = np.ones(len(lum), bool)
    # Agrupa por cor quantizada só para ACHAR o grupo dominante; o valor
    # devolvido é a MEDIANA dos pixels crus daquele grupo. Quantizar e devolver
    # o centro do balde arredonda para baixo e deixava toda a paleta ~4
    # unidades mais escura que a arte original.
    sel = px_vals[keep]
    q = (sel // 6 * 6)
    c = Counter(map(tuple, q))
    centro, n = c.most_common(1)[0]
    centro = np.array(centro)
    grupo = sel[(np.abs(sel - centro).max(axis=1) <= 5)]
    if len(grupo) < 5:
        grupo = sel
    r, g, b = np.median(grupo, axis=0)
    return (int(round(r)), int(round(g)), int(round(b))), n / keep.sum()


def analisa(nome, max_regioes=3):
    pdf = BW / f"{nome}.pdf"
    png = ORIG / f"{nome.replace('_', ' ')}.PNG"
    if not pdf.exists() or not png.exists():
        print(f"{nome}: arquivo faltando"); return
    orig = np.array(Image.open(png).convert("RGB"))
    bbox_orig = bbox_conteudo(orig)
    # caixa de conteúdo da figura vetorial: tudo que não é branco em um render
    # com todos os tons visíveis
    bbox_fig = bbox_conteudo(np.dstack([render_normal(pdf)] * 3))
    vals = tons(pdf)
    print(f"\n=== {nome}  ({len(vals)} tons) ===")
    for tom in sorted(vals, reverse=True):
        if tom >= 0.999 or tom <= 0.001:
            continue                      # branco de fundo e preto de texto
        mask = isola(pdf, tom) < 128
        if mask.sum() < 40:
            continue
        lbl, n = ndimage.label(mask)
        if not n:
            continue
        sizes = ndimage.sum(mask, lbl, range(1, n + 1))
        ordem = np.argsort(sizes)[::-1][:max_regioes]
        area_pct = mask.sum() / mask.size * 100
        amostras = []
        for i in ordem:
            if sizes[i] < 30:
                continue
            (r, g, b), pureza = amostra_regiao(orig, mask, lbl, i, bbox_orig, bbox_fig)
            sat = max(r, g, b) - min(r, g, b)
            amostras.append(f"#{r:02x}{g:02x}{b:02x}{'*' if sat > 25 else '·'}"
                            f"({pureza*100:.0f}%)")
        print(f"  tom {tom:.2f}  área {area_pct:5.2f}%  {vals[tom]:>3} ops  "
              f"→ {'  '.join(amostras)}")


if __name__ == "__main__":
    alvos = sys.argv[1:] or [p.stem for p in sorted(BW.glob("Cap*.pdf"))]
    for a in alvos:
        analisa(a)


def paleta_original(nome, top=10, min_pct=0.4):
    """Maiores áreas de cor chapada da arte original, por proporção da tela.

    Cruzar isto com a lista de tons é o que permite dizer 'o tom 0.98 cobre 17%
    da figura, e a original tem uma área de 16% nesta cor' — sem depender de
    acertar a posição.
    """
    png = ORIG / f"{nome.replace('_', ' ')}.PNG"
    a = np.array(Image.open(png).convert("RGB"))
    q = (a // 8 * 8).reshape(-1, 3)
    c = Counter(map(tuple, q))
    tot = len(q)
    out = []
    for (r, g, b), n in c.most_common(60):
        pct = n / tot * 100
        if pct < min_pct:
            break
        sat = max(r, g, b) - min(r, g, b)
        out.append((f"#{r:02x}{g:02x}{b:02x}", pct, sat))
        if len(out) >= top:
            break
    return out


def cruzado(nome):
    print(f"\n### {nome} — paleta da ARTE ORIGINAL (áreas chapadas)")
    for h, pct, sat in paleta_original(nome):
        marca = "COR " if sat > 20 else "cinza"
        print(f"    {h}  {pct:5.2f}% da tela  {marca}")
