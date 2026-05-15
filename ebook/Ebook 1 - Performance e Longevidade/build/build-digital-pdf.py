#!/usr/bin/env python3
"""Build digital-distribution PDF for 'ANTES' — Hotmart / direct download.

Differs from build-print-pdf.py:
    - A5 trim, oneside, no gutter, no bleed (screen reading, not printing).
    - ISBN matches the EPUB edition (978-65-02-06742-0), not the paperback.
    - Hyperlinks ON (clickable internal cross-refs + URLs).
    - Inline figure layout (no full-page figure-with-deferred-caption editorial move —
      that's a print-only treatment; on screen it creates awkward blank pages).
    - Skips Cap04 Fig02 native-LaTeX-table substitution (PNG is fine on screen).
    - Frontmatter dedication/epigraph kept but without the print blank-page padding.

Usage:
    python3 build-digital-pdf.py pt-BR
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
AUTHOR_PHOTO = BOOK_ROOT / "fotos" / "getulio_bw_halfbody_1000.jpg"
BUILD_DIR = BOOK_ROOT / "build"
WORK_DIR  = BUILD_DIR / f"work-digital-{LANG}"
TEMPLATE  = BUILD_DIR / "digital-template.tex"
OUT_PDF   = BUILD_DIR / f"Antes-{LANG}-digital.pdf"

READING_ORDER = [
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


# Referências is a NUMBERED chapter (Capítulo 18) — the file name and source
# H2 already say "Capítulo 18". Only Introdução / Agradecimentos / Sobre o
# Autor are unnumbered (front/backmatter editorial pieces).
UNNUMBERED_H1_PATTERNS = [
    re.compile(r'^Introdução\s*$', re.IGNORECASE),
    re.compile(r'^Agradecimentos\s*$', re.IGNORECASE),
    re.compile(r'^Sobre\s+o\s+[Aa]utor\s*$', re.IGNORECASE),
]


def is_unnumbered_h1(title):
    return any(p.match(title) for p in UNNUMBERED_H1_PATTERNS)


def transform_part_chapter_hierarchy(text):
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
                in_unnumbered_chapter = False
                out.append(line)
                i += 1
                continue
        m_lift = re.match(r'^(#{3,6})(\s+.+)$', line)
        if m_lift:
            new_hashes = m_lift.group(1)[1:]
            tail = m_lift.group(2)
            if in_unnumbered_chapter and new_hashes == "##":
                if "{.unnumbered" not in tail:
                    tail = tail.rstrip() + " {.unnumbered .unlisted}"
            out.append(new_hashes + tail + "\n")
            i += 1
            continue
        out.append(line)
        i += 1
    return "".join(out)


def transform_sobre_o_autor(text):
    """Special-case the 'Sobre o Autor' chapter to fit on a single page.

    The default \\chapter*{} heading + wide photo + full-leading bio + a 6-line
    contact block separated by paragraph breaks runs to two pages. We replace
    the whole thing with a hand-tuned compact LaTeX block: smaller heading,
    smaller photo, bio in \\small, and the contact info collapsed to a centred
    block with explicit line breaks (no parskip) so it all fits in one page.

    Also fixes the original Pandoc auto-caption bug (figure counter inherits
    from chapter 17 → 'Figura 17.1') by emitting raw \\includegraphics with
    no figure environment.
    """
    # Strip the raw-latex \chapter*{} block emitted upstream by
    # transform_part_chapter_hierarchy — we render our own heading.
    text = re.sub(
        r"```\{=latex\}\n\\chapter\*\{Sobre o Autor\}\n"
        r"\\addcontentsline\{toc\}\{chapter\}\{Sobre o Autor\}\n"
        r"\\markright\{Sobre o Autor\}\n```",
        "",
        text,
    )
    # Strip the photo markdown — emitted from the custom block instead.
    text = re.sub(
        r"!\[[^\]]*\]\(images/autor\.jpg\)\s*\n",
        "",
        text,
        flags=re.IGNORECASE,
    )

    # Split bio from contact at the first '---' horizontal rule.
    parts = re.split(r"\n---+\n", text.strip(), maxsplit=1)
    bio = parts[0].strip()
    contact_md = parts[1].strip() if len(parts) > 1 else ""

    def md_inline_to_tex(s):
        # **bold** → \textbf{}
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
        "{\\sffamily\\fontsize{16}{20}\\selectfont\\bfseries\\color{petrol} Sobre o Autor}\n"
        "\\end{center}\n"
        "\\vspace{0.4em}\n"
        "\\begin{center}\n"
        "\\includegraphics[width=0.24\\linewidth]{images/autor.jpg}\n"
        "\\end{center}\n"
        "\\vspace{0.2em}\n"
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
    """Strip 'Figura X.Y — ' prefix from alt text, rewrite paths to local images/."""
    figure_prefix = re.compile(r'^Figura\s+\d+(?:\.\d+)?\s*[—-]\s*', re.IGNORECASE)

    def replace_image(match):
        full = match.group(0)
        m = re.match(r'!\[([^\]]*)\]\(([^)]+)\)', full)
        if not m:
            return full
        alt, old_path = m.group(1), m.group(2)
        decoded = old_path.replace("%20", " ")
        filename = os.path.basename(decoded)
        safe = filename.replace(" ", "_")
        clean_alt = figure_prefix.sub("", alt)
        return f"![{clean_alt}]({fig_dir_name}/{safe})"

    return re.sub(r'!\[[^\]]*\]\([^)]+\)', replace_image, text)


def prepare_images(out_dir):
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
        # Compress for screen: re-encode PNGs to ~150 DPI equivalent if huge.
        # For digital we don't need 300 DPI — keep file size manageable.
        try:
            with Image.open(src) as im:
                # Cap longest side at 1800px (≈300 DPI at 6" or 200 DPI at 9")
                max_side = max(im.width, im.height)
                if max_side > 1800:
                    ratio = 1800 / max_side
                    new_w = int(im.width * ratio)
                    new_h = int(im.height * ratio)
                    im = im.resize((new_w, new_h), Image.LANCZOS)
                    if dst.suffix.lower() == ".png":
                        im.save(dst, "PNG", optimize=True)
                    else:
                        im.save(dst, "JPEG", quality=85, optimize=True)
                else:
                    shutil.copy2(src, dst)
        except Exception:
            shutil.copy2(src, dst)
        count += 1
    if AUTHOR_PHOTO.exists():
        shutil.copy2(AUTHOR_PHOTO, out_dir / "autor.jpg")
    return count


def build_digital_frontmatter():
    """Digital-only frontmatter — copyright with EPUB ISBN, dedication, epigraph, TOC.
    No blank-page padding (oneside, no recto/verso pretense)."""
    fm = r"""
```{=latex}
% --- Copyright page ---
\thispagestyle{empty}
\vspace*{0.6cm}
\begingroup
\sffamily\footnotesize\color{ink}\raggedright
\noindent ANTES — A Janela Silenciosa entre o Normal e o Ótimo — onde a saúde é decidida \par\smallskip
\noindent Copyright \copyright\ 2026 Getúlio José Mattos do Amaral Filho. \par
\noindent Todos os direitos reservados. \par\bigskip
\noindent \textbf{1\textsuperscript{a} edição digital} — Edição do Autor — 2026 \par\smallskip
\noindent ISBN 978-65-02-06742-0 \par\bigskip
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

% --- Dedication ---
\thispagestyle{empty}
\vspace*{4cm}
\begin{center}
\itshape\large
Para meus pacientes — \par
os que ficaram, \par
e os que partiram cedo demais.
\end{center}
\clearpage

% --- Epigraph ---
\thispagestyle{empty}
\vspace*{4cm}
\begin{center}
\itshape\normalsize
Entre saber e fazer, há um abismo.
\end{center}
\clearpage

% --- TOC ---
\tableofcontents
\clearpage

\mainmatter
\pagestyle{fancy}
```
"""
    return fm


def main():
    print(f"📖 Building DIGITAL PDF (Hotmart / direct download) for: {LANG}")
    if not MD_DIR.exists():
        sys.exit(f"❌ Markdown dir not found: {MD_DIR}")
    if not TEMPLATE.exists():
        sys.exit(f"❌ Template not found: {TEMPLATE}")

    if WORK_DIR.exists():
        shutil.rmtree(WORK_DIR)
    WORK_DIR.mkdir(parents=True)

    print("📸 Preparing images (resize → ≤1800px longest side)...")
    img_count = prepare_images(WORK_DIR / "images")
    print(f"   {img_count} figures prepared")

    print("📝 Consolidating chapters...")
    body = []
    body.append(build_digital_frontmatter())
    body.append("\n\n")
    included = 0
    backmatter_emitted = False
    for fname in READING_ORDER:
        src = MD_DIR / fname
        if not src.exists():
            print(f"   ⚠ Missing: {fname}")
            continue
        # Switch to backmatter once we hit Agradecimentos. \backmatter turns
        # off chapter numbering, suppresses any leftover \chaptermark headers,
        # and gives the closing pieces a clean final-pages feel.
        if fname == "agradecimentos.md" and not backmatter_emitted:
            body.append("\n```{=latex}\n\\backmatter\n```\n\n")
            backmatter_emitted = True
        text = src.read_text(encoding="utf-8")
        text = strip_frontmatter(text)
        text = strip_meta_sections(text)
        text = transform_part_chapter_hierarchy(text)
        # NO fullpage_figures() — digital keeps figures inline with captions.
        # NO replace_fig02_with_native_table() — PNG is fine at A5 screen size.
        if fname == "sobre-o-autor.md":
            text = transform_sobre_o_autor(text)
        text = rewrite_image_paths(text)
        body.append(text)
        body.append("\n\n")
        included += 1
    print(f"   {included} chapters included")

    consolidated = "".join(body)
    md_file = WORK_DIR / "book.md"
    md_file.write_text(consolidated, encoding="utf-8")
    print(f"   {len(consolidated.split()):,} words total")

    print("⚙  Running Pandoc + XeLaTeX...")
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

    if not OUT_PDF.exists():
        sys.exit("❌ PDF was not generated")

    size_mb = OUT_PDF.stat().st_size / (1024 * 1024)
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
    print(f"✅ Digital PDF built: {OUT_PDF}")
    print(f"   Size: {size_mb:.2f} MB")
    print(f"   Pages: {pagecount}")
    print(f"   ISBN: 978-65-02-06742-0 (matches EPUB edition)")


if __name__ == "__main__":
    main()
