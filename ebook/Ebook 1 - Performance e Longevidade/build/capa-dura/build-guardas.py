#!/usr/bin/env python3
"""Build the four endpaper PDFs for 'ANTES' — Brazilian HARDCOVER (capa dura).

Usage:
    python3 build-guardas.py pt-BR

Geometry (from gabarito.pdf — guardas tabs pp. 2-5):
    Each guarda is a BIFOLIO SPREAD — landscape, fold down the middle.
    Page size : 325 × 230 mm   (matches MediaBox of gabarito pp. 2-5)

    Layout when the book is opened at the front endpaper:
        [ inside front cover (glued) | free leaf recto ]   ← Guarda Capa frente
    When the free leaf is turned:
        [ free leaf verso (faces miolo p.1) | glued to cardboard ]
                                                            ← Verso Guarda Capa
    Symmetric on the back:
        [ free leaf recto | inside back cover (glued) ]    ← Guarda Contra Capa
        [ glued to cardboard | free leaf verso (faces last miolo p.) ]
                                                            ← Verso Guarda Contra Capa

    Per the gabarito visuals: the FRENTE sides design both halves; the
    VERSO sides have one half hidden under cardboard glue (gray in
    gabarito), so we keep those halves solid cream.

    The central ~20 mm vertical band IS the fold — content stays clear
    of it ("Evite deixar informações importantes na área de dobra").

Outputs (4 PDFs — exactly the upload slots requested by the printer):
    Antes-pt-BR-guarda-capa.pdf              — front bifolio, frente side
    Antes-pt-BR-verso-guarda-capa.pdf        — front bifolio, verso side
    Antes-pt-BR-guarda-contra-capa.pdf       — back bifolio, frente side
    Antes-pt-BR-verso-guarda-contra-capa.pdf — back bifolio, verso side

Design (Dr. Getúlio Amaral — author imprint, NOT Plenya clinic):
    Palette: paper-cream #f0e8d8, ink #1a1a1a, gold #b48a4a
    Type   : TeX Gyre Pagella (book serif, full Portuguese coverage)

    Front endpaper (frente): italic invocation of the book's thesis,
    centered on the FREE LEAF half (right side of spread). Inside-front-
    cover half (left) kept solid cream.

    Back endpaper (frente): wordmark institucional do Dr. Getúlio Amaral,
    centered on the FREE LEAF half (left side of spread). Inside-back-
    cover half (right) kept solid cream.

    Versos stay solid cream — vector PDF, no raster.
"""
from __future__ import annotations
import sys
import shutil
import subprocess
from pathlib import Path

THIS_DIR  = Path(__file__).resolve().parent
LANG = sys.argv[1] if len(sys.argv) > 1 else "pt-BR"

# --- Geometry (mm) — match gabarito.pdf pages 2-5 exactly ---
PAGE_W_MM = 325.0
PAGE_H_MM = 230.0
HALF_W_MM = PAGE_W_MM / 2   # 162.5 mm — width of each panel of the bifolio


# --- LaTeX preamble shared by all 4 guardas ---
PREAMBLE = r"""
\documentclass[11pt]{article}
\usepackage[
  paperwidth=PAGE_W_MMmm, paperheight=PAGE_H_MMmm,
  margin=0mm
]{geometry}
\usepackage{fontspec}
\usepackage{xcolor}
\usepackage{graphicx}

% Dr. Getúlio Amaral palette (matches /apps/site-getulio brand tokens)
\definecolor{paper}{HTML}{F0E8D8}
\definecolor{ink}{HTML}{1A1A1A}
\definecolor{gold}{HTML}{B48A4A}

\setmainfont{TeX Gyre Pagella}[
  Ligatures={Common,TeX},
  Numbers={Lining,Proportional}
]

\pagestyle{empty}
\setlength{\parindent}{0pt}

% Cream paper across the FULL spread (including 7,5 mm outer sangria) —
% \pagecolor fills paperwidth × paperheight, bleed automatically covered.
\pagecolor{paper}
\color{ink}

\begin{document}
""".strip()

POSTAMBLE = r"\end{document}"


def page_w_h(latex_str: str) -> str:
    return (latex_str
            .replace("PAGE_W_MM", f"{PAGE_W_MM:g}")
            .replace("PAGE_H_MM", f"{PAGE_H_MM:g}")
            .replace("HALF_W_MM", f"{HALF_W_MM:g}"))


# ============== Half-panel content =====================================

# Quote for the FREE LEAF of the front endpaper (book opening invocation).
# Kept distinct from the miolo epigraph — anticipates rather than echoes.
QUOTE_BODY = r"""
\null\vfill
\begin{center}
{\color{gold}\rule{34mm}{0.5pt}}

\vspace{3em}

{\itshape\fontsize{15}{22}\selectfont
Existe uma década entre\\[0.1em]
estar \emph{normal} e estar \emph{ótimo}.\\[1.4em]
É nessa janela silenciosa\\[0.1em]
que a saúde é decidida.}

\vspace{3em}

{\color{gold}\rule{34mm}{0.5pt}}
\end{center}
\vfill\null
""".strip()

# Wordmark for the FREE LEAF of the back endpaper (author colofão).
WORDMARK_BODY = r"""
\null\vfill
\begin{center}
{\color{gold}\rule{14mm}{0.5pt}}\quad
{\fontsize{10.5}{14}\selectfont\addfontfeatures{LetterSpace=32.0} DR.}\quad
{\color{gold}\rule{14mm}{0.5pt}}

\vspace{1.6em}

{\fontsize{28}{32}\selectfont\addfontfeatures{LetterSpace=8.0} GETÚLIO AMARAL}

\vspace{1.6em}

{\itshape\color{ink!70}\fontsize{10.5}{14}\selectfont\addfontfeatures{LetterSpace=12.0} Medicina guiada por raciocínio clínico.}

\vspace{3.2em}

{\color{gold}\rule{34mm}{0.5pt}}

\vspace{2em}

{\fontsize{8.5}{12}\selectfont\addfontfeatures{LetterSpace=20.0}\color{ink!55} EDIÇÃO DO AUTOR \quad·\quad LONDRINA \quad·\quad 2026}
\end{center}
\vfill\null
""".strip()

# A completely blank half (solid cream — the page background).
BLANK_BODY = r"""
\null
""".strip()


# ============== Spread assembly =========================================
#
# Two side-by-side minipages, each = half the paper width. The fold runs
# down the seam between them. We keep ~12 mm of padding on the inside edge
# of each half so content sits clear of the dobra (gabarito warning:
# "Evite deixar informações importantes na área de dobra").

# Padding on the inside edge (next to the fold) of each half.
INSIDE_PAD_MM = 12.0

SPREAD_TEMPLATE = r"""
\noindent
\begin{minipage}[c][\paperheight][c]{HALF_W_MMmm}
__LEFT_INNER__
\end{minipage}\begin{minipage}[c][\paperheight][c]{HALF_W_MMmm}
__RIGHT_INNER__
\end{minipage}
""".strip()


def spread(left_body: str, right_body: str) -> str:
    """Wrap two half-bodies into a spread, with inside-edge padding."""
    pad = f"{INSIDE_PAD_MM:g}mm"
    left_wrapped = (
        f"\\begin{{minipage}}[c][\\paperheight][c]{{\\dimexpr\\linewidth-{pad}\\relax}}\n"
        f"{left_body}\n"
        f"\\end{{minipage}}\\hspace{{{pad}}}"
    )
    right_wrapped = (
        f"\\hspace{{{pad}}}\\begin{{minipage}}[c][\\paperheight][c]{{\\dimexpr\\linewidth-{pad}\\relax}}\n"
        f"{right_body}\n"
        f"\\end{{minipage}}"
    )
    return (SPREAD_TEMPLATE
            .replace("__LEFT_INNER__", left_wrapped)
            .replace("__RIGHT_INNER__", right_wrapped))


# ============== The four pages ==========================================
#
# Page 1 — GUARDA CAPA (frente):
#   LEFT half  = inside front cover (visible, glued surface side facing reader)
#   RIGHT half = free leaf recto (first visible page of the book block)
#   → Quote on the FREE LEAF (right), inside-cover (left) solid cream.

# Page 2 — VERSO GUARDA CAPA:
#   LEFT half  = free leaf verso (visible, faces miolo p.1)
#   RIGHT half = glued to cardboard (hidden under horlle)
#   → Both halves blank cream. The visible left half stays a calm
#     transition page into the book block.

# Page 3 — GUARDA CONTRA CAPA (frente):
#   LEFT half  = free leaf recto (visible, last page before back cover)
#   RIGHT half = inside back cover (visible, glued surface side)
#   → Wordmark/colofão on the FREE LEAF (left), inside-cover (right)
#     solid cream.

# Page 4 — VERSO GUARDA CONTRA CAPA:
#   LEFT half  = glued to cardboard (hidden)
#   RIGHT half = free leaf verso (visible, faces last miolo page)
#   → Both halves blank cream.

PAGES = [
    ("guarda-capa",              spread(BLANK_BODY,    QUOTE_BODY)),
    ("verso-guarda-capa",        spread(BLANK_BODY,    BLANK_BODY)),
    ("guarda-contra-capa",       spread(WORDMARK_BODY, BLANK_BODY)),
    ("verso-guarda-contra-capa", spread(BLANK_BODY,    BLANK_BODY)),
]


# ============== Compilation =============================================

def compile_one(slug: str, body: str, work_dir: Path) -> Path:
    tex = page_w_h(PREAMBLE) + "\n" + page_w_h(body) + "\n" + POSTAMBLE
    tex_path = work_dir / f"{slug}.tex"
    tex_path.write_text(tex, encoding="utf-8")
    proc = subprocess.run(
        ["xelatex", "-interaction=nonstopmode", "-halt-on-error",
         f"-output-directory={work_dir}", tex_path.name],
        capture_output=True, text=True, cwd=str(work_dir),
    )
    pdf_path = work_dir / f"{slug}.pdf"
    if proc.returncode != 0 or not pdf_path.exists():
        log = (proc.stdout or "") + (proc.stderr or "")
        sys.exit(f"❌ xelatex failed for {slug}:\n{log[-2500:]}")
    return pdf_path


def main():
    print(f"📑 Building HARDCOVER guardas ({LANG}) — vector / xelatex")
    print(f"   Page size : {PAGE_W_MM:.0f} × {PAGE_H_MM:.0f} mm  (bifolio spread, matches gabarito pp. 2-5)")
    print(f"   Each half : {HALF_W_MM:.1f} mm wide")
    print(f"   Inside pad: {INSIDE_PAD_MM:.0f} mm each side of central fold")
    print(f"   Branding  : Dr. Getúlio Amaral (cream / ink / gold)")
    print()

    work_dir = THIS_DIR / "work-guardas"
    if work_dir.exists():
        shutil.rmtree(work_dir)
    work_dir.mkdir(parents=True)

    for slug, body in PAGES:
        pdf = compile_one(slug, body, work_dir)
        out = THIS_DIR / f"Antes-{LANG}-{slug}.pdf"
        shutil.copy2(pdf, out)
        size_kb = out.stat().st_size / 1024
        print(f"  ✅ {out.name}  —  {size_kb:.0f} KB  (vector)")

    print()
    print("✅ Guardas ready (lossless vector PDFs, 325 × 230 mm spread).")


if __name__ == "__main__":
    main()
