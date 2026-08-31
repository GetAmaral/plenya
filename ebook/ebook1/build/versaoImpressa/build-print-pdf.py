#!/usr/bin/env python3
"""Build print-ready PDF for 'ANTES' — PAPERBACK B&W version (KDP 6x9").

Esta é a variante "versão impressa" — isolada da versão digital e da capa dura.
Diferenças vs. build/build-print-pdf.py:

  1. Overrides de figuras em B&W: antes de copiar do diretório de cor original,
     o script procura uma versão substituta em ``figuras-bw/`` (mesmo basename).
     Vetoriais (.pdf) e raster B&W (.PNG) ambos suportados — o .pdf é
     preservado como vetor pelo XeLaTeX (\\includegraphics nativo).

  2. Substituições nativas em LaTeX para múltiplas figuras. Em vez de só Cap04
     Fig02 hard-coded, uma lista ``NATIVE_REPLACEMENTS`` mapeia basename →
     arquivo ``tabelas-nativas/*.tex``. Cada entrada vira uma tabela ou
     diagrama vetorial nativo no lugar do PNG.

  3. Todos os artefatos (work dir, PDF final, template, filtro Lua) ficam
     dentro de ``build/versaoImpressa/``. Nada vaza pro build/ principal,
     que segue alimentando a versão digital e a capa dura.

Usage:
    python3 build-print-pdf.py        # default: pt-BR
    python3 build-print-pdf.py pt-BR
"""

import os
import re
import sys
import shutil
import subprocess
from pathlib import Path
from PIL import Image

# --- Paths ------------------------------------------------------------------
# Este script vive em build/versaoImpressa/. O book root é dois níveis acima.
BUILD_DIR = Path(__file__).resolve().parent             # build/versaoImpressa/
BOOK_ROOT = BUILD_DIR.parent.parent                     # Ebook 1.../
LANG = sys.argv[1] if len(sys.argv) > 1 else "pt-BR"

MD_DIR       = BOOK_ROOT / "md" / LANG                  # markdown source (compartilhado)
FIG_DIR      = BOOK_ROOT / "figuras" / LANG             # figuras coloridas originais (fallback)
FIG_BW_DIR   = BUILD_DIR / "figuras-bw"                 # overrides B&W (PDF vetorial ou PNG)
TABLES_DIR   = BUILD_DIR / "tabelas-nativas"            # substituições LaTeX vetoriais
AUTHOR_PHOTO = BOOK_ROOT / "fotos" / "getulio_bw_halfbody_fullres.jpg"
WORK_DIR     = BUILD_DIR / f"work-print-{LANG}"
TEMPLATE     = BUILD_DIR / "print-template.tex"
DROPCAP_LUA  = BUILD_DIR / "print-dropcaps.lua"
OUT_PDF      = BUILD_DIR / f"Antes-{LANG}-print-interior.pdf"


# --- Reading order ----------------------------------------------------------
# Mesma ordem do build digital — conteúdo é idêntico, só diagramação muda.
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


# --- Substituições LaTeX nativas -------------------------------------------
# Mapa (basename_seguro, arquivo.tex). O basename usa "_" no lugar de espaços —
# é o mesmo formato em que os PNGs são normalizados ao serem copiados.
# Cada arquivo .tex em tabelas-nativas/ contém o LaTeX vetorial que substitui
# o PNG correspondente. À medida que vetorizarmos novas figuras, é só
# adicionar uma linha aqui e dropar o .tex em tabelas-nativas/.
NATIVE_REPLACEMENTS = [
    ("Cap04_Fig02", "cap04-fig02.tex"),   # Biomarcadores: Normais vs Ótimas
    # (futuro) ("Cap10_Fig01", "cap10-fig01.tex"),
    # (futuro) ("Cap15_Fig01", "cap15-fig01.tex"),
]


# --- Full-page figures ------------------------------------------------------
# Figuras que ganham página inteira no impresso (caption na página seguinte).
# Não confundir com NATIVE_REPLACEMENTS — full-page rege POSICIONAMENTO; native
# rege CONTEÚDO (PNG vira LaTeX).
FULLPAGE_FIGURES = [
    ("Cap04_Fig02", "4.2"),   # Tabela Normal vs Ótimo (também native)
    ("Cap12_Fig01", "12.1"),  # HPA cascade
    ("Cap12_Fig03", "12.3"),  # 5 instrumentos psicológicos
    ("Cap16_Fig01", "16.1"),  # 2 modelos de acompanhamento
]


# === Markdown transforms (preservados do build original) ====================

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


def fullpage_figures(text, available_overrides):
    """Full-page treatment para figuras verticais/referência.

    Usa o filename real disponível em images/ (override .pdf ganha prioridade
    sobre PNG original).
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
        # Skip se a figura foi substituída por LaTeX nativo (já não é mais PNG).
        if any(b == safe_basename for b, _ in NATIVE_REPLACEMENTS):
            continue

        # Filename real em images/ (pode ser .PNG do original ou .pdf override)
        final_name = available_overrides.get(safe_basename, f"{safe_basename}.PNG")

        original_name = safe_basename.replace("_", " ")
        pattern = re.compile(
            r'!\[([^\]]*)\]\((?:\.\./)?'
            + re.escape(original_name).replace(r"\ ", r"(?:%20|_| )")
            + r'\.PNG\)',
            re.IGNORECASE
        )

        def make_replacer(label_local, file_local):
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
                    f"{{images/{file_local}}}\n"
                    "\\vspace*{\\fill}\n"
                    "\\end{figure}\n"
                    "\\clearpage\n"
                    f"\\noindent{{\\sffamily\\footnotesize\\textbf{{Figura {label_local}.\\ }}"
                    + clean + "\\par}\n"
                    "\\bigskip\n"
                    "```\n"
                )
            return _r

        text = pattern.sub(make_replacer(label, final_name), text)
    return text


def apply_native_replacements(text):
    """Substitui referências a PNGs por blocos LaTeX nativos vetoriais.

    Para cada (basename, tex_file) em NATIVE_REPLACEMENTS, lê tabelas-nativas/
    <tex_file> e injeta como bloco raw LaTeX no lugar da marcação ![..](X.PNG).
    """
    for safe_basename, tex_file in NATIVE_REPLACEMENTS:
        tex_path = TABLES_DIR / tex_file
        if not tex_path.exists():
            print(f"  ⚠ tabelas-nativas/{tex_file} não encontrado — pulando {safe_basename}")
            continue
        tex_block = tex_path.read_text(encoding="utf-8")

        original_name = safe_basename.replace("_", " ")
        pattern = re.compile(
            r'!\[[^\]]*\]\((?:\.\./)?'
            + re.escape(original_name).replace(r"\ ", r"(?:%20|_| )")
            + r'\.PNG\)',
            re.IGNORECASE
        )
        replacement = "\n```{=latex}\n" + tex_block + "\n```\n"
        text = pattern.sub(lambda m: replacement, text)
    return text


def transform_sobre_o_autor(text):
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


def rewrite_image_paths(text, available_overrides, fig_dir_name="images"):
    """Reescreve caminhos das imagens e strip do prefixo 'Figura X.Y — '.

    Se houver override em figuras-bw/ (PDF ou PNG), usa o filename do override
    (que pode ter extensão diferente do PNG original).
    """
    figure_prefix = re.compile(r'^Figura\s+\d+(?:\.\d+)?\s*[—-]\s*', re.IGNORECASE)

    def replace_image(match):
        full = match.group(0)
        m = re.match(r'!\[([^\]]*)\]\(([^)]+)\)', full)
        if not m:
            return full
        alt, old_path = m.group(1), m.group(2)
        decoded = old_path.replace("%20", " ")
        filename = os.path.basename(decoded)
        # Basename normalizado (espaços→underscore), SEM extensão.
        stem = os.path.splitext(filename)[0].replace(" ", "_")
        # Override determina extensão final; se não, mantém .PNG original.
        final = available_overrides.get(stem, f"{stem}{os.path.splitext(filename)[1]}")
        clean_alt = figure_prefix.sub("", alt)
        return f"![{clean_alt}]({fig_dir_name}/{final})"

    return re.sub(r'!\[[^\]]*\]\([^)]+\)', replace_image, text)


# === Image pipeline ========================================================

def prepare_images(out_dir):
    """Copia figuras pro work dir/images/, preferindo overrides B&W.

    Para cada figura no diretório de cor original, checa se existe um arquivo
    com mesmo basename em ``figuras-bw/`` (com extensão .pdf ou .PNG). Se sim,
    usa o override; senão, copia o original colorido.

    Retorna dict {basename_seguro_sem_ext: filename_final_em_images}.
    Necessário porque PDFs vetoriais entram com extensão .pdf e o rewriter
    precisa saber.
    """
    out_dir.mkdir(parents=True, exist_ok=True)
    overrides_map = {}   # {Cap04_Fig02: "Cap04_Fig02.pdf"} ou {Cap07_Fig03: "Cap07_Fig03.PNG"}

    # Index overrides disponíveis (case-insensitive sobre extensão)
    bw_index = {}        # {stem_lowercase: path}
    if FIG_BW_DIR.exists():
        for f in FIG_BW_DIR.iterdir():
            if f.is_file() and f.suffix.lower() in {".pdf", ".png", ".jpg", ".jpeg"}:
                bw_index[f.stem.lower()] = f

    if not FIG_DIR.exists():
        print(f"  ⚠ {FIG_DIR} não existe — pulando figuras")
        return 0, overrides_map

    count_copy = 0
    count_override = 0
    for src in sorted(FIG_DIR.iterdir()):
        if not (src.is_file() and src.suffix.lower() in {".png", ".jpg", ".jpeg"}):
            continue
        safe_stem = src.stem.replace(" ", "_")
        override = bw_index.get(safe_stem.lower())
        if override:
            dst = out_dir / f"{safe_stem}{override.suffix}"
            shutil.copy2(override, dst)
            overrides_map[safe_stem] = dst.name
            count_override += 1
        else:
            dst = out_dir / f"{safe_stem}{src.suffix}"
            shutil.copy2(src, dst)
            overrides_map[safe_stem] = dst.name  # mantém o nome p/ rewriter
            count_copy += 1

    if AUTHOR_PHOTO.exists():
        shutil.copy2(AUTHOR_PHOTO, out_dir / "autor.jpg")

    print(f"   {count_copy + count_override} figuras copiadas ({count_override} overrides B&W)")
    return count_copy + count_override, overrides_map


# Imagens que NÃO são impressas na largura total do bloco de texto. O valor é a
# fração de \linewidth usada no LaTeX; sem isso o audit acusa DPI baixo em
# imagens que na página saem pequenas (e nítidas).
PRINT_WIDTH_FRACTION = {
    "autor.jpg": 0.24,   # \includegraphics[width=0.24\linewidth] em "Sobre o Autor"
}


def audit_image_dpi(img_dir, text_width_in=4.25, soft_dpi=300, hard_dpi=200):
    """Auditoria DPI vs largura REAL de impressão de cada imagem.
    PDFs vetoriais ficam fora do audit (DPI infinito por definição)."""
    result = {"ok": [], "soft": [], "hard": [], "vector": []}
    for img in sorted(img_dir.iterdir()):
        suf = img.suffix.lower()
        if suf == ".pdf":
            result["vector"].append((img.name, "vetor", "∞"))
            continue
        if suf not in {".png", ".jpg", ".jpeg"}:
            continue
        try:
            with Image.open(img) as im:
                px = im.width
                printed_in = text_width_in * PRINT_WIDTH_FRACTION.get(img.name, 1.0)
                dpi = int(px / printed_in)
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


# === Frontmatter ===========================================================

def build_print_frontmatter(work_dir):
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


# === Main =================================================================

def main():
    print(f"📖 Building PRINT PDF (versaoImpressa) for language: {LANG}")
    print(f"   build root: {BUILD_DIR}")
    for required, label in [(MD_DIR, "markdown"), (TEMPLATE, "template"), (DROPCAP_LUA, "lua filter")]:
        if not required.exists():
            sys.exit(f"❌ {label} não encontrado: {required}")

    if WORK_DIR.exists():
        shutil.rmtree(WORK_DIR)
    WORK_DIR.mkdir(parents=True)

    # 1. Prepare images (color originals + B&W overrides)
    print("📸 Copiando imagens (com overrides B&W quando existem)...")
    _, overrides_map = prepare_images(WORK_DIR / "images")

    # 1b. DPI audit
    audit = audit_image_dpi(WORK_DIR / "images")
    print(f"   DPI audit (a 4.25\" largura de texto):")
    print(f"     ✓ OK     (≥300 DPI): {len(audit['ok'])} figuras")
    if audit["vector"]:
        print(f"     ◆ VETOR  (.pdf):    {len(audit['vector'])} figuras")
        for name, *_ in audit["vector"]:
            print(f"        - {name}")
    if audit["soft"]:
        print(f"     ⚠ SOFT   (200-299 DPI): {len(audit['soft'])} — legível mas mole:")
        for name, px, dpi in audit["soft"]:
            print(f"        - {name}: {px}px → {dpi} DPI")
    if audit["hard"]:
        print(f"     ✗ HARD   (<200 DPI): {len(audit['hard'])} — regerar:")
        for name, px, dpi in audit["hard"]:
            print(f"        - {name}: {px}px → {dpi} DPI")

    # 2. Consolida capítulos
    print("📝 Consolidando capítulos...")
    body = []
    body.append(build_print_frontmatter(WORK_DIR))
    body.append("\n\n")
    included = 0
    backmatter_emitted = False
    for fname in READING_ORDER:
        src = MD_DIR / fname
        if not src.exists():
            print(f"   ⚠ Faltando: {fname}")
            continue
        if fname == "agradecimentos.md" and not backmatter_emitted:
            body.append("\n```{=latex}\n\\backmatter\n```\n\n")
            backmatter_emitted = True
        text = src.read_text(encoding="utf-8")
        text = strip_frontmatter(text)
        text = strip_meta_sections(text)
        text = transform_part_chapter_hierarchy(text)
        # PRIMEIRO: substituições nativas (PNG → LaTeX vetorial)
        text = apply_native_replacements(text)
        # DEPOIS: full-page treatment para os que sobraram como imagem
        text = fullpage_figures(text, overrides_map)
        if fname == "sobre-o-autor.md":
            text = transform_sobre_o_autor(text)
        text = rewrite_image_paths(text, overrides_map)
        body.append(text)
        body.append("\n\n")
        included += 1
    print(f"   {included} capítulos incluídos")

    # Even page count guard
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
    print(f"   {len(consolidated.split()):,} palavras totais")

    # 3. Pandoc → XeLaTeX → PDF
    print("⚙  Rodando Pandoc + XeLaTeX (pode levar 1–3 min)...")
    cmd = [
        "pandoc",
        str(md_file),
        "--from=markdown+yaml_metadata_block+smart+raw_tex+autolink_bare_uris",
        "--pdf-engine=xelatex",
        f"--template={TEMPLATE}",
        f"--resource-path={WORK_DIR}",
        f"--lua-filter={DROPCAP_LUA}",
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
            print("\n❌ Pandoc/XeLaTeX falhou:\n" + err)
            sys.exit(result.returncode)
    except FileNotFoundError as e:
        sys.exit(f"❌ Ferramenta faltando: {e}")

    # 4. Reporta
    if not OUT_PDF.exists():
        sys.exit("❌ PDF não foi gerado")

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
    print(f"✅ PDF gerado: {OUT_PDF}")
    print(f"   Tamanho: {size_mb:.2f} MB")
    print(f"   Páginas: {pagecount}")
    print()
    print("Próximo passo: passar o page count pro cover builder se a largura da lombada mudou.")


if __name__ == "__main__":
    main()
