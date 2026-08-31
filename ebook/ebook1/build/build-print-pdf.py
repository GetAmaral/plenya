#!/usr/bin/env python3
"""Build print-ready PDF for 'ANTES' — Amazon KDP Paperback (6x9").

Usage:
    python3 build-print-pdf.py pt-BR

Pipeline:
    Markdown chapters → consolidate → Pandoc → XeLaTeX → PDF (interior).
    Cover is a SEPARATE artifact built only after pagecount is known.

Differences vs. build-epub.py:
    - PNG figures kept as PNG (300 DPI quality preserved for print).
    - Frontmatter expanded for print (half-title, title, copyright with ISBN, dedication, epigraph, TOC).
    - mainmatter / backmatter switches for proper page numbering.
    - Output is a single PDF ready for KDP miolo upload.
"""

import os
import re
import sys
import shutil
import subprocess
from pathlib import Path
from PIL import Image

BOOK_ROOT = Path(__file__).resolve().parent.parent
LANG = sys.argv[1] if len(sys.argv) > 1 else "pt-BR"

MD_DIR    = BOOK_ROOT / "md" / LANG
FIG_DIR   = BOOK_ROOT / "figuras" / LANG
AUTHOR_PHOTO = BOOK_ROOT / "fotos" / "getulio_bw_halfbody_fullres.jpg"
BUILD_DIR = BOOK_ROOT / "build"
WORK_DIR  = BUILD_DIR / f"work-print-{LANG}"
TEMPLATE  = BUILD_DIR / "print-template.tex"
OUT_PDF   = BUILD_DIR / f"Antes-{LANG}-print-interior.pdf"

# Reading order: SAME as EPUB so book content is identical.
# Only frontmatter / part divisions are typeset differently for print.
READING_ORDER = [
    # frontmatter.md and 00a-creditos.md are NOT included here:
    # the print template provides typeset copyright/dedication/epigraph
    # via build_print_frontmatter(). They remain in the EPUB build.
    "00b-introducao.md",
    "01-homem-que-morreu-saudavel.md",
    "02-quatro-assassinos-silenciosos.md",
    "03-seu-corpo-esta-envelhecendo.md",
    "04-mapa-dos-biomarcadores.md",
    "05-seu-coracao-esta-falando.md",
    "06-saude-metabolica.md",
    "06b-parte-iii-intro.md",
    "07-atividade-fisica.md",
    "08-alimentacao-suplementacao.md",
    "09-sistemas-cardio-renal-hepatico.md",
    "10-paineis-bioquimicos.md",
    "11-genomica-exposicoes.md",
    "12-integracao-mente-corpo.md",
    "13-conexao-proposito-sentido.md",
    "14-ritmo-circadiano-e-repouso.md",
    "15-seu-placar-de-longevidade.md",
    "16-quando-procurar-especialista.md",
    "17-manifesto-agir.md",
    "18-referencias-recursos.md",
    "agradecimentos.md",
    "sobre-o-autor.md",
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
# Referências = Capítulo 18 (numbered) — only the front/backmatter pieces are unnumbered.
UNNUMBERED_H1_PATTERNS = [
    re.compile(r'^Introdução\s*$', re.IGNORECASE),
    re.compile(r'^Agradecimentos\s*$', re.IGNORECASE),
    re.compile(r'^Sobre\s+o\s+[Aa]utor\s*$', re.IGNORECASE),
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
        m_part = re.match(r'^#\s+PARTE\s+([IVX]+)\s*[—-]\s*(.+?)\s*$', line)
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
        m_ch = re.match(r'^##\s+Capítulo\s+\d+\s*[—-]\s*(.+)$', line)
        if m_ch:
            title = m_ch.group(1).strip()
            if is_unnumbered_h1(title):
                out.append("```{=latex}\n")
                out.append(f"\\chapter*{{{title}}}\n")
                out.append(f"\\addcontentsline{{toc}}{{chapter}}{{{title}}}\n")
                out.append(f"\\markright{{{title}}}\n")
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
                out.append(f"\\markright{{{title}}}\n")
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
    figure_prefix = re.compile(r'^Figura\s+\d+(?:\.\d+)?\s*[—-]\s*', re.IGNORECASE)

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
                return (
                    "\n```{=latex}\n"
                    "\\clearpage\n"
                    "\\thispagestyle{empty}\n"
                    "\\begin{figure}[p]\n"
                    "\\centering\n"
                    "\\vspace*{\\fill}\n"
                    f"\\includegraphics[height=0.95\\textheight, width=\\linewidth, keepaspectratio]"
                    f"{{images/{safe_basename}.PNG}}\n"
                    "\\vspace*{\\fill}\n"
                    "\\end{figure}\n"
                    "\\clearpage\n"
                    f"\\noindent{{\\sffamily\\footnotesize\\textbf{{Figura {label_local}.\\ }}"
                    + clean + "\\par}\n"
                    "\\bigskip\n"
                    "```\n"
                )
            return _r

        text = pattern.sub(make_replacer(label), text)
    return text


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


def transform_sobre_o_autor(text):
    """Special-case the 'Sobre o Autor' chapter to fit on a single page.

    Same approach as the digital build: custom heading, smaller photo, bio in
    \\small, and a centred contact block with explicit line breaks instead of
    paragraph breaks. Fixes the Pandoc auto-caption bug too.
    """
    text = re.sub(
        r"```\{=latex\}\n\\chapter\*\{Sobre o Autor\}\n"
        r"\\addcontentsline\{toc\}\{chapter\}\{Sobre o Autor\}\n"
        r"\\markright\{Sobre o Autor\}\n```",
        "",
        text,
    )
    text = re.sub(
        r"!\[[^\]]*\]\(images/autor\.jpg\)\s*\n",
        "",
        text,
        flags=re.IGNORECASE,
    )

    parts = re.split(r"\n---+\n", text.strip(), maxsplit=1)
    bio = parts[0].strip()
    contact_md = parts[1].strip() if len(parts) > 1 else ""

    def md_inline_to_tex(s):
        return re.sub(r'\*\*(.+?)\*\*', r'\\textbf{\1}', s)

    contact_lines = [md_inline_to_tex(ln.strip())
                     for ln in contact_md.split("\n") if ln.strip()]
    contact_tex = " \\\\\n".join(contact_lines)

    head = (
        "```{=latex}\n"
        "\\clearpage\n"
        "\\thispagestyle{plain}\n"
        "\\addcontentsline{toc}{chapter}{Sobre o Autor}\n"
        "\\markright{Sobre o Autor}\n"
        "\\vspace*{0.2em}\n"
        "\\begin{center}\n"
        "{\\sffamily\\fontsize{18}{22}\\selectfont\\bfseries\\color{petrol} Sobre o Autor}\n"
        "\\end{center}\n"
        "\\vspace{0.5em}\n"
        "\\begin{center}\n"
        "\\includegraphics[width=0.24\\linewidth]{images/autor.jpg}\n"
        "\\end{center}\n"
        "\\vspace{0.3em}\n"
        "\\begingroup\n"
        "\\small\\linespread{0.96}\\selectfont\n"
        "```\n"
    )
    tail = (
        "\n```{=latex}\n"
        "\\par\\smallskip\n"
        "\\begin{center}\n"
        "\\sffamily\\footnotesize\\color{ink}\\linespread{1.0}\\selectfont\n"
        + contact_tex + "\n"
        "\\end{center}\n"
        "\\endgroup\n"
        "```\n"
    )
    return head + bio + tail


def transform_author_photo(text):
    """Backwards-compatible alias."""
    return transform_sobre_o_autor(text)


def rewrite_image_paths(text, fig_dir_name="images"):
    """Rewrite image paths AND strip the 'Figura X.Y — ' prefix from alt text.

    Pandoc auto-prefixes 'Figura X.Y.' to every figure caption in LaTeX. Source alt
    text starts with 'Figura X.Y — ...' for EPUB readability — but in print this
    duplicates ('Figura 4.1. Figura 4.1 — ...'). Print-only strip keeps EPUB intact.
    """
    figure_prefix = re.compile(r'^Figura\s+\d+(?:\.\d+)?\s*[—-]\s*', re.IGNORECASE)

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
\noindent ANTES — A Janela Silenciosa entre o Normal e o Ótimo — onde a saúde é decidida \par\smallskip
\noindent Copyright \copyright\ 2026 Getúlio José Mattos do Amaral Filho. \par
\noindent Todos os direitos reservados. \par\bigskip
\noindent \textbf{1\textsuperscript{a} edição} — Edição do Autor — 2026 \par\smallskip
\noindent ISBN 978-65-02-07691-0 \par\bigskip
\noindent Nenhuma parte desta obra pode ser reproduzida, armazenada ou
transmitida por qualquer meio (eletrônico, mecânico, fotocópia, gravação ou
outro) sem a autorização prévia e por escrito do autor, exceto para citações
breves em resenhas, estudos e materiais didáticos, com indicação da fonte. \par\bigskip
\noindent \textbf{Aviso médico.} Este livro tem caráter educacional. As
informações apresentadas refletem a visão clínica do autor à data da
publicação e não substituem a consulta presencial com profissional habilitado.
Decisões sobre exames, medicamentos, suplementos ou mudanças de hábito devem
ser tomadas em conjunto com seu médico assistente. \par\bigskip
\noindent \textit{Projeto gráfico, capa e diagramação:} Plenya Saúde. \par
\noindent \textit{Fonte de corpo:} TeX Gyre Pagella. \quad \textit{Fonte de títulos:} Inter. \par\bigskip
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
Para meus pacientes — \par
os que ficaram, \par
e os que partiram cedo demais.
\end{center}
\clearpage

% --- (vii) blank ---
\thispagestyle{empty}\null\clearpage

% --- (viii) Epigraph ---
\thispagestyle{empty}
\vspace*{6cm}
\begin{center}
\itshape\normalsize
Entre saber e fazer, há um abismo.
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
    backmatter_emitted = False
    for fname in READING_ORDER:
        src = MD_DIR / fname
        if not src.exists():
            print(f"   ⚠ Missing: {fname}")
            continue
        # Switch to backmatter once we hit Agradecimentos. \backmatter turns
        # off chapter numbering and gives the closing pieces a clean
        # final-pages feel (no leftover headers from the last numbered chapter).
        if fname == "agradecimentos.md" and not backmatter_emitted:
            body.append("\n```{=latex}\n\\backmatter\n```\n\n")
            backmatter_emitted = True
        text = src.read_text(encoding="utf-8")
        text = strip_frontmatter(text)
        text = strip_meta_sections(text)
        text = transform_part_chapter_hierarchy(text)
        # Editorial full-page treatment for the 4 vertical/reference figures
        # (4.2, 12.1, 12.3, 16.1). Must run BEFORE rewrite_image_paths so the
        # image markdown is captured intact with the original alt text.
        text = fullpage_figures(text)
        if fname == "sobre-o-autor.md":
            text = transform_sobre_o_autor(text)
        text = rewrite_image_paths(text)
        body.append(text)
        body.append("\n\n")
        included += 1
    print(f"   {included} chapters included")

    # Industry standard: print books must have an even page count (each leaf
    # is a recto+verso pair). After \clearpage ships the last content page,
    # \value{page} holds the number of the *next* page — so if the last shipped
    # page was odd, the counter is even, and we need to inject a blank verso.
    # We test the page-1 expression so the comparison is on the page that was
    # actually printed. KDP and UICLAP both reject odd page counts.
    body.append(
        "\n```{=latex}\n"
        "\\clearpage\n"
        "\\ifodd\\numexpr\\value{page}-1\\relax\n"
        "  \\null\\thispagestyle{empty}\\clearpage\n"
        "\\fi\n"
        "```\n"
    )

    consolidated = "".join(body)
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
        "-V", "lang=pt-BR",
        "-V", "title=ANTES",
        "-V", "subtitle=A Janela Silenciosa entre o Normal e o Ótimo — onde a saúde é decidida",
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
