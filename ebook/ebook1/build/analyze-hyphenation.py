#!/usr/bin/env python3
"""Mede hifenização (e glifos faltando) num miolo já compilado.

Lê o PDF via ``mutool draw -F stext`` para ter fonte + corpo de cada linha, o
que separa o que é TEXTO CORRIDO do que é TÍTULO. Hífen de fim de linha em
título é erro tipográfico duro (títulos são ragged-right: nunca deveriam
hifenizar); no corpo justificado é só uma taxa a manter baixa.

Também conta U+FFFD, que é como o extrator representa o .notdef box — o
"quadrado" que aparece quando a fonte não tem o glifo (ex.: TeX Gyre Pagella
não tem os subscritos U+2082..U+2084, então "VO₂" imprimia "VO□").

Uso:
    ./analyze-hyphenation.py caminho/para/miolo.pdf [--exemplos N]
"""

import collections
import html
import re
import subprocess
import sys
import tempfile
from pathlib import Path

RE_PAGE = re.compile(r'<page id="(\w+)"')
RE_FONT = re.compile(r'<font name="([^"]+)" size="([\d.]+)"')
RE_CHAR = re.compile(r'\sc="(.*?)"/>')
RE_QUAD = re.compile(r'quad="([-\d.eE ]+)"')
RE_HYPH_END = re.compile(r"[A-Za-zÀ-ÿ][-‐]$")

# Corpo do texto = a serifada. Qualquer coisa em Inter num corpo >= 10pt é
# título/rótulo — ali hífen de quebra é inaceitável.
BODY_FONT_PREFIX = "TeXGyrePagella"


def extract_lines(pdf: Path):
    """[(page, font, size, texto)] — uma entrada por linha desenhada."""
    with tempfile.TemporaryDirectory() as td:
        xml = Path(td) / "stext.xml"
        subprocess.run(
            ["mutool", "draw", "-F", "stext", "-o", str(xml), str(pdf)],
            check=True, capture_output=True,
        )
        page = None
        fonts, chars = set(), []
        for raw in xml.open(encoding="utf-8", errors="replace"):
            m = RE_PAGE.search(raw)
            if m:
                page = m.group(1)
                continue
            if "<line " in raw:
                fonts, chars = set(), []
                continue
            m = RE_FONT.search(raw)
            if m:
                fonts.add((m.group(1), round(float(m.group(2)), 1)))
                continue
            m = RE_CHAR.search(raw)
            if m:
                ch = html.unescape(m.group(1))
                q = RE_QUAD.search(raw)
                x0 = x1 = None
                if q:
                    v = [float(n) for n in q.group(1).split()]
                    x0, x1 = v[0], v[2]
                chars.append((ch, x0, x1))
                continue
            if "</line>" in raw:
                font, size = sorted(fonts)[0] if fonts else ("?", 0.0)
                yield page, font, size, chars


def main():
    args = [a for a in sys.argv[1:] if not a.startswith("--")]
    if not args:
        sys.exit(__doc__)
    pdf = Path(args[0])
    n_ex = 6
    for a in sys.argv[1:]:
        if a.startswith("--exemplos="):
            n_ex = int(a.split("=", 1)[1])

    total = collections.Counter()
    hyph = collections.Counter()
    examples = collections.defaultdict(list)
    tofu = []

    spaces = []   # largura mediana do espaço de cada linha, em em
    right = []    # borda direita de cada linha do corpo, em pt

    for page, font, size, chars in extract_lines(pdf):
        text = "".join(c for c, _, _ in chars)
        key = (font, size)
        if font.startswith(BODY_FONT_PREFIX) and size > 9:
            # Largura do próprio glifo de espaço. O espaço natural da Pagella
            # é 0,25 em; acima de 1,8x já se enxerga rio branco na página.
            sp = sorted(x1 - x0 for c, x0, x1 in chars
                        if c == " " and x0 is not None and x1 > x0)
            if len(sp) >= 4:
                spaces.append(sp[len(sp) // 2] / size / 0.25)
            xs = [x1 for _, _, x1 in chars if x1 is not None]
            if xs:
                right.append(max(xs))
        total[key] += 1
        if RE_HYPH_END.search(text):
            hyph[key] += 1
            if len(examples[key]) < n_ex:
                examples[key].append((page, text[-48:]))


    # O .notdef box não vem marcado no stext do mutool; o pdftotext o entrega
    # como U+FFFD. Contamos por lá.
    txt = subprocess.run(["pdftotext", str(pdf), "-"],
                         capture_output=True, text=True).stdout
    tofu = [ln.strip()[:70] for ln in txt.splitlines() if "\ufffd" in ln]

    body_lines = sum(v for k, v in total.items() if k[0].startswith(BODY_FONT_PREFIX))
    body_hyph = sum(v for k, v in hyph.items() if k[0].startswith(BODY_FONT_PREFIX))
    title_hyph = {k: v for k, v in hyph.items() if not k[0].startswith(BODY_FONT_PREFIX)}

    print(f"== {pdf.name}")
    pct = 100 * body_hyph / body_lines if body_lines else 0
    print(f"corpo (Pagella): {body_hyph} hífens em {body_lines} linhas ({pct:.1f}%)")
    print(f"títulos/rótulos: {sum(title_hyph.values())} hífens  "
          f"{'← DEVE SER 0' if title_hyph else '✅'}")
    print(f"glifos faltando (□): {len(tofu)} linhas  "
          f"{'← DEVE SER 0' if tofu else '✅'}")
    if spaces:
        s = sorted(spaces)
        n = len(s)
        frouxa = sum(1 for v in s if v > 1.8)
        pessima = sum(1 for v in s if v > 2.5)
        print(f"espaço entre palavras (x natural): mediana {s[n//2]:.2f}  "
              f"p95 {s[int(n*0.95)]:.2f}  p99 {s[int(n*0.99)]:.2f}")
        print(f"linhas frouxas: {100*frouxa/n:.1f}% >1,8x   "
              f"{100*pessima/n:.1f}% >2,5x  (de {n} linhas justificadas)")

    if title_hyph:
        print("\n-- hífens em títulos/rótulos --")
        for k, v in sorted(title_hyph.items(), key=lambda kv: -kv[1]):
            print(f"  {k[0]} {k[1]}pt: {v}")
            for p, t in examples[k]:
                print(f"      {p}  {t!r}")
    if tofu:
        print("\n-- linhas com glifo faltando --")
        for ln in tofu[:20]:
            print(f"  {ln!r}")

    if right:
        r = sorted(right)
        margem = r[int(len(r) * 0.90)]      # borda direita do bloco de texto
        over = [v for v in r if v > margem + 2]
        print(f"overfull (texto passando da margem +2pt): {len(over)} linhas"
              + (f"  pior +{max(over)-margem:.1f}pt" if over else "  ✅"))

    print("\n-- hífens no corpo, por variante --")
    for k, v in sorted(hyph.items(), key=lambda kv: -kv[1]):
        if k[0].startswith(BODY_FONT_PREFIX):
            print(f"  {k[0]} {k[1]}pt: {v}/{total[k]} ({100*v/total[k]:.1f}%)")


if __name__ == "__main__":
    main()
