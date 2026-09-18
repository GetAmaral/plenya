#!/usr/bin/env python3
"""Tira o número de dentro da arte das figuras.

Quem numera figura no livro é o contador do LaTeX (`\\caption`), que segue a ordem
de leitura. O número gravado dentro do desenho é a ordem em que a figura foi FEITA,
não bate com a de leitura (no cap. 8 o livro mostra a "FIGURA 2" antes da "FIGURA 1"),
não bate nem com o nome do arquivo (Cap11_Fig01 traz "FIGURA 3") e não tem capítulo,
então não dá para citar. Sai.

Dois casos:

  tarja    "FIGURA 2" num chip isolado → redação do texto + do chip.
  inline   "Figura 1 — Os 20 Anos Silenciosos" → redação da linha inteira e o título
           redesenhado sem o prefixo, no mesmo ponto de origem, no mesmo Inter-Bold,
           tamanho e cor. Só cobrir o prefixo deixaria o título 90pt adentro, porque
           esses títulos são alinhados à esquerda.

Idempotente: o original vai para `_com-numero/` na primeira passagem e é sempre ele
que é lido. Rodar de novo não empilha efeito, e apagar o `_com-numero/` restaura.

Uso:  ./runpy strip-figure-numbers.py [--dry]
"""
import re, sys, shutil
from pathlib import Path
import pymupdf as fitz

HERE = Path(__file__).resolve().parent
DIRS = [HERE / "figuras-cor", HERE / "figuras-bw"]
BACKUP = "_com-numero"
INTER_BOLD = "/usr/share/fonts/opentype/inter/Inter-Bold.otf"

PAT_BADGE  = re.compile(r"^F\s*I\s*G\s*U\s*R\s*A\s*\d+\s*\.?$", re.I)
PAT_INLINE = re.compile(r"^Figura\s+\d+\s*[—–-]\s*(.+)$")

DRY = "--dry" in sys.argv


def header_line(page):
    """Devolve (kind, line, texto) da linha que carrega o número, ou None."""
    for b in page.get_text("dict")["blocks"]:
        if b["type"] != 0:
            continue
        for l in b["lines"]:
            t = "".join(s["text"] for s in l["spans"]).strip()
            if not t:
                continue
            if PAT_BADGE.match(t):
                return "tarja", l, t
            m = PAT_INLINE.match(t)
            if m:
                return "inline", l, t
    return None


def other_line_boxes(page, skip_bbox):
    out = []
    for b in page.get_text("dict")["blocks"]:
        if b["type"] != 0:
            continue
        for l in b["lines"]:
            if tuple(l["bbox"]) == tuple(skip_bbox):
                continue
            if "".join(s["text"] for s in l["spans"]).strip():
                out.append(l["bbox"])
    return out


def chip_rect(page, text_rect):
    """Une o retângulo do texto ao chip colorido atrás dele, se houver."""
    r = fitz.Rect(text_rect)
    page_area = page.rect.get_area()
    for d in page.get_drawings():
        dr = d["rect"]
        if dr.get_area() > page_area * 0.03:
            continue                      # fundo ou moldura: não é chip
        # o chip é o retângulo que fica ATRÁS do texto: tem de conter a caixa do
        # texto (com folga) e não ser muito maior que ela, senão é outra coisa do
        # desenho e a redação levaria arte junto.
        # o fundo da página também "contém" o texto; o chip é da ordem de 3x a
        # caixa do texto e o fundo passa de 500x, então 12x separa com folga.
        if dr.get_area() > r.get_area() * 12.0:
            continue
        if dr.x0 <= r.x0 + 3 and dr.y0 <= r.y0 + 3 and dr.x1 >= r.x1 - 3 and dr.y1 >= r.y1 - 3:
            r |= dr
    return r


def strip(path: Path):
    bdir = path.parent / BACKUP
    bdir.mkdir(exist_ok=True)
    orig = bdir / path.name
    if not orig.exists():
        shutil.copy2(path, orig)

    doc = fitz.open(orig)
    page = doc[0]
    found = header_line(page)
    if not found:
        doc.close()
        return "sem número", ""
    kind, line, text = found
    span = line["spans"][0]

    if kind == "tarja":
        rect = chip_rect(page, line["bbox"])
        rect += (-1.5, -1.5, 1.5, 1.5)
        novo = None
    else:
        rect = fitz.Rect(line["bbox"])
        below = [b for b in other_line_boxes(page, line["bbox"]) if b[1] >= rect.y1 - 1.0]
        if below:
            rect.y1 = min(rect.y1, min(b[1] for b in below) - 0.5)
        novo = PAT_INLINE.match(text).group(1).strip()

    if DRY:
        doc.close()
        return kind, f"{text!r} → {novo!r} em {tuple(round(v,1) for v in rect)}"

    page.add_redact_annot(rect)
    page.apply_redactions(images=fitz.PDF_REDACT_IMAGE_NONE)

    if novo:
        c = span["color"]
        rgb = ((c >> 16) & 255) / 255, ((c >> 8) & 255) / 255, (c & 255) / 255
        page.insert_text(fitz.Point(*span["origin"]), novo,
                         fontname="FInterBold", fontfile=INTER_BOLD,
                         fontsize=span["size"], color=rgb)

    doc.save(path, garbage=3, deflate=True)
    doc.close()
    return kind, (novo or text)


def main():
    tot = {"tarja": 0, "inline": 0, "sem número": 0}
    for d in DIRS:
        if not d.is_dir():
            print(f"  ⚠ {d} não existe"); continue
        print(f"\n{d.name}/")
        for f in sorted(d.glob("*.pdf")):
            kind, det = strip(f)
            tot[kind] = tot.get(kind, 0) + 1
            print(f"  {f.name:22s} {kind:10s} {det}")
    print("\n" + " · ".join(f"{k}: {v}" for k, v in tot.items()))
    if DRY:
        print("(--dry: nada foi escrito)")


if __name__ == "__main__":
    main()
