#!/usr/bin/env python3
"""Build print-ready PDF for 'BEFORE' (English edition) — Amazon KDP Paperback (6x9").

Usage:
    python3 build-print-pdf-en.py
"""

import os
import re
import sys
import shutil
import subprocess
from pathlib import Path
from PIL import Image

BOOK_ROOT = Path(__file__).resolve().parent.parent
LANG = "en"

MD_DIR    = BOOK_ROOT / "md" / LANG
FIG_DIR   = BOOK_ROOT / "figuras" / LANG
AUTHOR_PHOTO = BOOK_ROOT / "fotos" / "getulio_bw_halfbody_1000.jpg"
BUILD_DIR = BOOK_ROOT / "build"
WORK_DIR  = BUILD_DIR / f"work-print-{LANG}"
TEMPLATE  = BUILD_DIR / "print-template-en.tex"
OUT_PDF   = BUILD_DIR / "Before-en-print-interior.pdf"

# Reading order: SAME as EPUB so book content is identical.
# Only frontmatter / part divisions are typeset differently for print.
READING_ORDER = [
    # frontmatter.md and 00a-credits.md are NOT included here:
    # the print template provides typeset copyright/dedication/epigraph
    # via build_print_frontmatter(). They remain in the EPUB build.
    "00b-introduction.md",
    "01-man-who-almost-died-healthy.md",
    "02-four-that-kill-in-silence.md",
    "03-age-your-body-actually-has.md",
    "04-expanded-panel.md",
    "05-arteries-break-their-silence.md",
    "06-metabolic-health.md",
    "06b-part-iii-intro.md",
    "07-activity.md",
    "08-alimentation-smart-adjuncts.md",
    "09-systems-cardio-renal-hepatic.md",
    "10-biochemical-hormonal-panels.md",
    "11-genomics-exposures.md",
    "12-inner-work.md",
    "13-connection-purpose-meaning.md",
    "14-sleep-rhythm-recovery.md",
    "15-longevity-scorecard.md",
    "16-when-to-see-specialist.md",
    "17-acts-manifesto.md",
    "18-references-resources.md",
    "acknowledgments.md",
    "about-the-author.md",
]


def strip_frontmatter(text):
    if text.startswith("---\n"):
        end = text.find("\n---\n", 4)
        if end != -1:
            return text[end + 5:]
    return text


def strip_meta_sections(text):
    s, e = "<!-- EPUB-START -->", "<!-- EPUB-END -->"
    si = text.find(s)
    if si != -1:
        text = text[si + len(s):]
    ei = text.find(e)
    if ei != -1:
        text = text[:ei]
    return text.strip() + "\n"


# H1 headings that must NOT be numbered chapters in print.
# They become \chapter*{Title} with manual TOC entry, instead of \chapter{Title}.
UNNUMBERED_H1_PATTERNS = [
    re.compile(r'^Introduction\s*$', re.IGNORECASE),
    re.compile(r'^References(?:\s+and\s+Resources)?.*$', re.IGNORECASE),
    re.compile(r'^Acknowledgments\s*$', re.IGNORECASE),
    re.compile(r'^About\s+the\s+[Aa]uthor\s*$', re.IGNORECASE),
]


def is_unnumbered_h1(title):
    return any(p.match(title) for p in UNNUMBERED_H1_PATTERNS)


def transform_part_chapter_hierarchy(text):
    """Print-only transform: lift PARTE headings to LaTeX \part{}, promote chapter h2 to h1.

    Source pattern (legacy from EPUB build):
      # PARTE I — O DESPERTAR
      ## Capítulo 1 — O Homem que Quase Morreu Saudável

    Target (print):
      ```{=latex}
      \part{O Despertar}
      ```
      # Capítulo 1 — O Homem que Quase Morreu Saudável

    Files that don't have the PARTE header pass through with chapter h1 untouched.
    EPUB build remains unaffected — this runs only in the print pipeline.
    """
    lines = text.splitlines(keepends=True)
    out = []
    i = 0
    n = len(lines)
    in_unnumbered_chapter = False
    while i < n:
        line = lines[i]
        m_part = re.match(r'^#\s+PART\s+([IVX]+)\s*[—-]\s*(.+?)\s*$', line)
        if m_part:
            title = m_part.group(2).strip()
            pretty = title.title() if title.isupper() else title
            out.append("```{=latex}\n")
            out.append(f"\\part{{{pretty}}}\n")
            out.append("```\n\n")
            i += 1
            if i < n and lines[i].strip() == "":
                i += 1
            continue
        # Promote "## Capítulo N — Título" H2 to H1 AND strip the "Capítulo N — "
        # prefix — LaTeX template auto-prints "CAPÍTULO N" via title format.
        m_ch = re.match(r'^##\s+Chapter\s+\d+\s*[—-]\s*(.+)$', line)
        if m_ch:
            title = m_ch.group(1).strip()
            if is_unnumbered_h1(title):
                out.append("```{=latex}\n")
                out.append(f"\\chapter*{{{title}}}\n")
                out.append(f"\\addcontentsline{{toc}}{{chapter}}{{{title}}}\n")
                out.append(f"\\markboth{{{title}}}{{}}\n")
                out.append("```\n\n")
                in_unnumbered_chapter = True
            else:
                out.append(f"# {title}\n")
                in_unnumbered_chapter = False
            i += 1
            continue
        # Plain H1 — Introdução / Agradecimentos / Sobre o Autor → \chapter*{}
        m_h1 = re.match(r'^#\s+(.+)$', line)
        if m_h1:
            title = m_h1.group(1).strip()
            if is_unnumbered_h1(title):
                out.append("```{=latex}\n")
                out.append(f"\\chapter*{{{title}}}\n")
                out.append(f"\\addcontentsline{{toc}}{{chapter}}{{{title}}}\n")
                out.append(f"\\markboth{{{title}}}{{}}\n")
                out.append("```\n\n")
                in_unnumbered_chapter = True
                i += 1
                continue
            else:
                # Numbered H1 — chapters that already arrived as plain # Title.
                in_unnumbered_chapter = False
                out.append(line)
                i += 1
                continue
        # Once chapters are at H1, sub-sections must shift up one level too:
        # ### → ##, #### → ###. Only headings, never inline #.
        m_lift = re.match(r'^(#{3,6})(\s+.+)$', line)
        if m_lift:
            new_hashes = m_lift.group(1)[1:]
            tail = m_lift.group(2)
            # Inside unnumbered chapters, suppress section numbering & TOC entries
            if in_unnumbered_chapter and new_hashes == "##":
                # Append {.unnumbered .unlisted} attribute if not already present
                if "{.unnumbered" not in tail:
                    tail = tail.rstrip() + " {.unnumbered .unlisted}"
            out.append(new_hashes + tail + "\n")
            i += 1
            continue
        out.append(line)
        i += 1
    return "".join(out)


# Figures that get full-page editorial treatment in print:
# image lives alone on its page, caption deferred to next page.
# (file_basename, figure_label) pairs — matched against ![Figura X.Y — ...](...PNG).
FULLPAGE_FIGURES = [
    ("Cap04_Fig02", "4.2"),  # Tabela Normal vs Ótimo
    ("Cap12_Fig01", "12.1"), # HPA cascade
    ("Cap12_Fig03", "12.3"), # 5 instrumentos psicológicos
    ("Cap16_Fig01", "16.1"), # 2 modelos de acompanhamento
]


def fullpage_figures(text):
    """Print-only: list of figures to render on their own page, caption deferred
    to the next page. Editorial move for vertical/reference figures that lose
    impact when squeezed inline (Penguin/FSG style).

    EPUB build keeps the inline image+caption — only print is restructured.
    """
    figure_prefix = re.compile(r'^Figure\s+\d+(?:\.\d+)?\s*[—-]\s*', re.IGNORECASE)

    def escape_tex(s):
        return (s
            .replace("\\", r"\textbackslash{}")
            .replace("&", r"\&")
            .replace("%", r"\%")
            .replace("#", r"\#")
            .replace("_", r"\_")
            .replace("$", r"\$")
            .replace("{", r"\{")
            .replace("}", r"\}"))

    for safe_basename, label in FULLPAGE_FIGURES:
        # Pattern accepts space, %20, or _ in the filename (markdown source uses spaces)
        original_name = safe_basename.replace("_", " ")
        pattern = re.compile(
            r'!\[([^\]]*)\]\((?:\.\./)?'
            + re.escape(original_name).replace(r"\ ", r"(?:%20|_| )")
            + r'\.PNG\)',
            re.IGNORECASE
        )

        def make_replacer(label_local):
            def _r(match):
                alt = match.group(1)
                clean = figure_prefix.sub("", alt).strip()
                clean = escape_tex(clean)
                # \stepcounter{figure} keeps the LaTeX figure counter in sync
                # with the pandoc-rendered figures so subsequent auto-numbered
                # figures in the same chapter get the correct N.M label.
                return (
                    "\n```{=latex}\n"
                    "\\clearpage\n"
                    "\\thispagestyle{empty}\n"
                    "\\stepcounter{figure}\n"
                    "\\begin{figure}[p]\n"
                    "\\centering\n"
                    "\\vspace*{\\fill}\n"
                    f"\\includegraphics[height=0.95\\textheight, width=\\linewidth, keepaspectratio]"
                    f"{{images/{safe_basename}.PNG}}\n"
                    "\\vspace*{\\fill}\n"
                    "\\end{figure}\n"
                    "\\clearpage\n"
                    "\\noindent{\\sffamily\\footnotesize\\textbf{Figure \\thefigure.\\ }"
                    + clean + "\\par}\n"
                    "\\bigskip\n"
                    "```\n"
                )
            return _r

        text = pattern.sub(make_replacer(label), text)
    return text


def fix_author_photo(text):
    """Print-only: render the author headshot as an unnumbered centered image
    instead of letting pandoc wrap it in a numbered figure environment
    (which currently produces an unwanted "Figure 17.1" caption)."""
    pattern = re.compile(r'!\[Getulio Amaral Filho\]\((?:\.\./)?(?:images/)?autor\.jpg\)')
    replacement = (
        "\n```{=latex}\n"
        "\\begin{center}\n"
        "\\includegraphics[width=0.45\\linewidth]{images/autor.jpg}\n"
        "\\end{center}\n"
        "```\n"
    )
    return pattern.sub(lambda m: replacement, text)


def replace_fig02_with_native_table(text):
    """Print-only: replace Cap04 Fig02 (Tabela Normal vs Ótimo) PNG ref with a
    native LaTeX longtable. Vector typography → infinitely sharp at any DPI.
    EPUB build keeps the PNG. Note: this is the SECOND figure in chapter 4 (4.2);
    the first (4.1, caso Fernanda) keeps its PNG since it carries patient data
    that should remain a single visual unit.
    """
    table_path = BUILD_DIR / "cap04-fig02.tex"
    if not table_path.exists():
        return text
    table_tex = table_path.read_text(encoding="utf-8")
    pattern = re.compile(
        r'!\[[^\]]*\]\((?:\.\./)?Cap04(?:%20|_| )Fig02\.PNG\)',
        re.IGNORECASE
    )
    replacement = "\n```{=latex}\n" + table_tex + "\n```\n"
    return pattern.sub(lambda m: replacement, text)


def rewrite_image_paths(text, fig_dir_name="images"):
    """Rewrite image paths AND strip the 'Figura X.Y — ' prefix from alt text.

    Pandoc auto-prefixes 'Figura X.Y.' to every figure caption in LaTeX. Source alt
    text starts with 'Figura X.Y — ...' for EPUB readability — but in print this
    duplicates ('Figura 4.1. Figura 4.1 — ...'). Print-only strip keeps EPUB intact.
    """
    figure_prefix = re.compile(r'^Figure\s+\d+(?:\.\d+)?\s*[—-]\s*', re.IGNORECASE)

    def replace_image(match):
        full = match.group(0)
        # Pull alt + path out
        m = re.match(r'!\[([^\]]*)\]\(([^)]+)\)', full)
        if not m:
            return full
        alt, old_path = m.group(1), m.group(2)
        decoded = old_path.replace("%20", " ")
        filename = os.path.basename(decoded)
        safe = filename.replace(" ", "_")
        # Strip the 'Figura X.Y — ' prefix from alt text in print only
        clean_alt = figure_prefix.sub("", alt)
        return f"![{clean_alt}]({fig_dir_name}/{safe})"

    return re.sub(r'!\[[^\]]*\]\([^)]+\)', replace_image, text)


def prepare_images(out_dir):
    """Copy figures preserving original quality.
    Print PDF must keep PNG/JPEG at original resolution; KDP requires ≥300 DPI in print size.
    """
    out_dir.mkdir(parents=True, exist_ok=True)
    if not FIG_DIR.exists():
        print(f"  ⚠ {FIG_DIR} does not exist — skipping figures")
        return 0
    count = 0
    for src in sorted(FIG_DIR.iterdir()):
        if not (src.is_file() and src.suffix.lower() in {".png", ".jpg", ".jpeg"}):
            continue
        safe = src.name.replace(" ", "_")
        dst = out_dir / safe
        shutil.copy2(src, dst)
        count += 1
    if AUTHOR_PHOTO.exists():
        shutil.copy2(AUTHOR_PHOTO, out_dir / "autor.jpg")
    return count


def audit_image_dpi(img_dir, text_width_in=4.25, soft_dpi=300, hard_dpi=200):
    """Audit figures vs print DPI.

    Three categories:
      OK    : DPI >= soft_dpi at full text-width (4.25\")           — clean print
      SOFT  : DPI between hard_dpi and soft_dpi at text-width        — readable but slightly soft
      HARD  : DPI below hard_dpi at text-width                       — visible pixelation; regenerate

    Returns dict {ok, soft, hard} of lists with (name, width_px, dpi_at_text_width).
    """
    result = {"ok": [], "soft": [], "hard": []}
    for img in sorted(img_dir.iterdir()):
        if img.suffix.lower() not in {".png", ".jpg", ".jpeg"}:
            continue
        try:
            with Image.open(img) as im:
                px = im.width
                dpi = int(px / text_width_in)
                entry = (img.name, px, dpi)
                if dpi >= soft_dpi:
                    result["ok"].append(entry)
                elif dpi >= hard_dpi:
                    result["soft"].append(entry)
                else:
                    result["hard"].append(entry)
        except Exception:
            pass
    return result


def build_print_frontmatter(work_dir):
    """Print-only frontmatter (copyright/ISBN/dedication/epigraph/TOC).
    Wrapped in pandoc raw-latex block so backslash sequences pass through verbatim.
    The template emits half-title (i) and title page (iii); this block emits iv onwards.
    """
    fm = r"""
```{=latex}
% --- Copyright page (iv) ---
\thispagestyle{empty}
\vspace*{1cm}
\begingroup
\sffamily\footnotesize\color{ink}\raggedright
\noindent BEFORE — The silent window — a decade between normal and optimal where health is decided \par\smallskip
\noindent Copyright \copyright\ 2026 Getulio José Mattos do Amaral Filho. \par
\noindent All rights reserved. \par\bigskip
\noindent \textbf{1\textsuperscript{st} English edition} — Author's Edition — 2026 \par\smallskip
\noindent Translated from the Brazilian Portuguese original \textit{ANTES} (1\textsuperscript{st} edition, 2026). \par\smallskip
\noindent ISBN: [ISBN-EN-PENDING] \par\bigskip
\noindent No part of this work may be reproduced, stored in a retrieval system, or
transmitted, in any form or by any means — electronic, mechanical, photocopy,
recording, or otherwise — without the prior written permission of the author,
except in brief quotations for reviews, academic articles, and other uses
permitted by law. \par\bigskip
\noindent \textbf{Important medical notice.} This book is for educational purposes only.
The information presented reflects the clinical practice of the author at the date
of publication and does not replace consultation with a qualified health professional.
Decisions about tests, medications, supplements, or lifestyle changes should be
made together with your treating physician. \par\bigskip
\noindent \textit{Cover, layout, and design:} Plenya Saúde. \par
\noindent \textit{Body font:} TeX Gyre Pagella. \quad \textit{Display font:} Inter. \par\bigskip
\noindent \textcolor{gold}{\textbf{PLENYA}}\quad plenyasaude.com.br \par
\noindent drgetulioamaralfilho.com.br
\endgroup
\clearpage

% --- (v) blank ---
\thispagestyle{empty}\null\clearpage

% --- (vi) Dedication ---
\thispagestyle{empty}
\vspace*{6cm}
\begin{center}
\itshape\large
To my patients — \par
those who stayed, \par
and those who left too soon.
\end{center}
\clearpage

% --- (vii) blank ---
\thispagestyle{empty}\null\clearpage

% --- (viii) Epigraph ---
\thispagestyle{empty}
\vspace*{6cm}
\begin{center}
\itshape\normalsize
Between knowing and doing, there is an abyss.
\end{center}
\clearpage

% --- (ix) blank ---
\thispagestyle{empty}\null\clearpage

% --- (x) TOC ---
\tableofcontents
\clearpage

\mainmatter
\pagestyle{fancy}
```
"""
    return fm


def main():
    print(f"📖 Building PRINT PDF for language: {LANG}")
    if not MD_DIR.exists():
        sys.exit(f"❌ Markdown dir not found: {MD_DIR}")
    if not TEMPLATE.exists():
        sys.exit(f"❌ Template not found: {TEMPLATE}")

    if WORK_DIR.exists():
        shutil.rmtree(WORK_DIR)
    WORK_DIR.mkdir(parents=True)

    # 1. Prepare images
    print("📸 Copying images at original resolution...")
    img_count = prepare_images(WORK_DIR / "images")
    print(f"   {img_count} figures copied")

    # 1b. DPI audit (realistic threshold at 4.25" text width = full body)
    audit = audit_image_dpi(WORK_DIR / "images")
    print(f"   DPI audit (at 4.25\" text width):")
    print(f"     ✓ OK    (≥300 DPI): {len(audit['ok'])} figures")
    if audit["soft"]:
        print(f"     ⚠ SOFT  (200-299 DPI): {len(audit['soft'])} — readable, slightly soft:")
        for name, px, dpi in audit["soft"]:
            print(f"        - {name}: {px}px → {dpi} DPI")
    if audit["hard"]:
        print(f"     ✗ HARD  (<200 DPI): {len(audit['hard'])} — should regenerate:")
        for name, px, dpi in audit["hard"]:
            print(f"        - {name}: {px}px → {dpi} DPI")

    # 2. Consolidate chapters
    print("📝 Consolidating chapters...")
    body = []
    body.append(build_print_frontmatter(WORK_DIR))
    body.append("\n\n")
    included = 0
    for fname in READING_ORDER:
        src = MD_DIR / fname
        if not src.exists():
            print(f"   ⚠ Missing: {fname}")
            continue
        text = src.read_text(encoding="utf-8")
        text = strip_frontmatter(text)
        text = strip_meta_sections(text)
        text = transform_part_chapter_hierarchy(text)
        # Editorial full-page treatment for the 4 vertical/reference figures
        # (4.2, 12.1, 12.3, 16.1). Must run BEFORE rewrite_image_paths so the
        # image markdown is captured intact with the original alt text.
        text = fullpage_figures(text)
        text = fix_author_photo(text)
        text = rewrite_image_paths(text)
        body.append(text)
        body.append("\n\n")
        included += 1
    print(f"   {included} chapters included")

    consolidated = "".join(body)
    # Map Unicode subscript/superscript digits to LaTeX commands so the
    # embedded font keeps a valid /ToUnicode entry (search & copy/paste work).
    # The Pagella subscript codepoints (U+2080..U+2089) lack a toUnicode map.
    # Inside inline code / fenced code, replace with plain digits since
    # monospace doesn't render subscripts and raw \textsubscript would print literally.
    SCRIPT_MAP_TEX = {
        0x2080: r'\textsubscript{0}', 0x2081: r'\textsubscript{1}',
        0x2082: r'\textsubscript{2}', 0x2083: r'\textsubscript{3}',
        0x2084: r'\textsubscript{4}', 0x2085: r'\textsubscript{5}',
        0x2086: r'\textsubscript{6}', 0x2087: r'\textsubscript{7}',
        0x2088: r'\textsubscript{8}', 0x2089: r'\textsubscript{9}',
        0x2070: r'\textsuperscript{0}', 0x00B9: r'\textsuperscript{1}',
        0x00B2: r'\textsuperscript{2}', 0x00B3: r'\textsuperscript{3}',
        0x2074: r'\textsuperscript{4}', 0x2075: r'\textsuperscript{5}',
        0x2076: r'\textsuperscript{6}', 0x2077: r'\textsuperscript{7}',
        0x2078: r'\textsuperscript{8}', 0x2079: r'\textsuperscript{9}',
    }
    SCRIPT_MAP_PLAIN = {k: v.split('{')[1].rstrip('}') for k, v in SCRIPT_MAP_TEX.items()}
    code_pattern = re.compile(r'(```[\s\S]*?```|`[^`\n]+`)')
    parts = code_pattern.split(consolidated)
    for idx in range(len(parts)):
        if parts[idx].startswith('`'):
            parts[idx] = parts[idx].translate(SCRIPT_MAP_PLAIN)
        else:
            parts[idx] = parts[idx].translate(SCRIPT_MAP_TEX)
    consolidated = ''.join(parts)
    md_file = WORK_DIR / "book.md"
    md_file.write_text(consolidated, encoding="utf-8")
    print(f"   {len(consolidated.split()):,} words total")

    # 3. Run Pandoc → XeLaTeX → PDF
    print("⚙  Running Pandoc + XeLaTeX (this can take 1–3 min on first run)...")
    dropcap_filter = BUILD_DIR / "print-dropcaps.lua"
    cmd = [
        "pandoc",
        str(md_file),
        "--from=markdown+yaml_metadata_block+smart+raw_tex+autolink_bare_uris",
        "--pdf-engine=xelatex",
        f"--template={TEMPLATE}",
        f"--resource-path={WORK_DIR}",
        f"--lua-filter={dropcap_filter}",
        "--top-level-division=chapter",
        # TOC depth controlled by \setcounter{tocdepth}{0} in print-template.tex.
        # Pandoc's --toc-depth must be ≥1; we keep it at 1 and the template overrides.
        "--toc-depth=1",
        "-V", "documentclass=book",
        "-V", "lang=en-US",
        "-V", "title=BEFORE",
        "-V", "subtitle=The silent window — a decade between normal and optimal where health is decided",
        "-V", "author=Dr. Getulio Amaral Filho",
        "-o", str(OUT_PDF),
    ]
    try:
        result = subprocess.run(cmd, capture_output=True, text=True, cwd=str(WORK_DIR))
        if result.returncode != 0:
            err = (result.stderr or "")[-3000:]
            print("\n❌ Pandoc/XeLaTeX failed:\n" + err)
            sys.exit(result.returncode)
    except FileNotFoundError as e:
        sys.exit(f"❌ Required tool missing: {e}")

    # 4. Report
    if not OUT_PDF.exists():
        sys.exit("❌ PDF was not generated")

    size_mb = OUT_PDF.stat().st_size / (1024 * 1024)
    # Get page count via pdfinfo if available
    pagecount = "?"
    try:
        info = subprocess.run(["pdfinfo", str(OUT_PDF)], capture_output=True, text=True)
        for line in info.stdout.splitlines():
            if line.startswith("Pages:"):
                pagecount = line.split(":")[1].strip()
                break
    except FileNotFoundError:
        pass

    print()
    print(f"✅ PDF built: {OUT_PDF}")
    print(f"   Size: {size_mb:.2f} MB")
    print(f"   Pages: {pagecount}")
    print()
    print("Next step: pass the page count to the cover builder so spine width is correct.")


if __name__ == "__main__":
    main()
