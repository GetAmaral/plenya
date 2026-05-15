#!/usr/bin/env python3
"""Build EPUB for AGORA · Vol. 1 — Energia e Disposição (pt-BR)."""

import os
import re
import sys
import shutil
import subprocess
from pathlib import Path

BOOK_ROOT = Path(__file__).resolve().parent.parent
MD_DIR    = BOOK_ROOT / "rascunhos"
FIG_DIR   = BOOK_ROOT / "figuras"
COVER     = BOOK_ROOT / "capas" / "pt-BR" / "capa-v7-final.jpg"
BUILD_DIR = BOOK_ROOT / "build"
WORK_DIR  = BUILD_DIR / "work-pt-BR"
OUT_EPUB  = BUILD_DIR / "AGORA-Vol1-Energia-Disposicao.epub"


def strip_frontmatter(text):
    if text.startswith("---\n"):
        end = text.find("\n---\n", 4)
        if end != -1:
            return text[end + 5:]
    return text


def rewrite_image_paths(text, fig_dir_name="images"):
    def replace_path(match):
        full = match.group(0)
        path_match = re.search(r'\]\(([^)]+)\)', full)
        if not path_match:
            return full
        old_path = path_match.group(1)
        decoded = old_path.replace("%20", " ")
        filename = os.path.basename(decoded)
        safe_name = filename.replace(" ", "_")
        return full.replace(old_path, f"{fig_dir_name}/{safe_name}")
    return re.sub(r'!\[[^\]]*\]\([^)]+\)', replace_path, text)


def prepare_images(images_out_dir):
    images_out_dir.mkdir(parents=True, exist_ok=True)
    if not FIG_DIR.exists():
        return 0
    count = 0
    for src in sorted(FIG_DIR.rglob("*")):
        if not (src.is_file() and src.suffix.lower() in {".png", ".jpg", ".jpeg", ".webp"}):
            continue
        safe_name = src.name.replace(" ", "_")
        dst = images_out_dir / safe_name
        shutil.copy2(src, dst)
        count += 1
    return count


def main():
    print("Building EPUB: AGORA Vol. 1")

    if not MD_DIR.exists():
        sys.exit(f"Markdown dir not found: {MD_DIR}")

    if WORK_DIR.exists():
        shutil.rmtree(WORK_DIR)
    WORK_DIR.mkdir(parents=True)

    img_count = prepare_images(WORK_DIR / "images")
    print(f"  {img_count} figures copied")

    chapters = sorted([p for p in MD_DIR.glob("*.md") if p.is_file()])
    full_text = []
    for src in chapters:
        text = src.read_text(encoding="utf-8")
        text = strip_frontmatter(text)
        text = rewrite_image_paths(text)
        full_text.append(text.strip())
        full_text.append("\n\n")
    print(f"  {len(chapters)} chapters")

    consolidated = "".join(full_text)
    md_file = WORK_DIR / "book.md"
    md_file.write_text(consolidated, encoding="utf-8")
    wc = len(consolidated.split())
    print(f"  {wc:,} words")

    metadata = WORK_DIR / "metadata.yaml"
    metadata.write_text("""---
title: "AGORA · Vol. 1"
subtitle: "Do normal ao ótimo — Energia e disposição: o cansaço que não passa"
creator:
  - role: author
    text: "Dr. Getúlio José Mattos do Amaral Filho"
publisher: "Edição do Autor"
rights: "Copyright © 2026 Getúlio José Mattos do Amaral Filho. Todos os direitos reservados."
language: pt-BR
date: "2026"
description: |
  Volume 1 da Série AGORA. O cansaço crônico inexplicado — o que vem antes do
  frasco e o atlas dirigido dos suplementos. Auditoria das doze perguntas,
  painel laboratorial específico e o atlas da fadiga (ferro, B12, D, magnésio,
  CoQ10, NAD+, adaptógenos, tireoide) com formas, doses e timing.
subject:
  - "Medicina preventiva"
  - "Fadiga crônica"
  - "Suplementação"
  - "Longevidade"
...
""", encoding="utf-8")

    css = WORK_DIR / "style.css"
    css.write_text("""body {
  font-family: "Georgia", serif;
  font-size: 1em;
  line-height: 1.55;
  margin: 0;
  padding: 0 0.6em;
  color: #111;
  text-align: justify;
  hyphens: auto;
}
h1, h2, h3, h4 { font-family: "Georgia", serif; font-weight: bold; text-align: left;
  page-break-after: avoid; margin-top: 1.4em; margin-bottom: 0.6em; hyphens: none; }
h1 { font-size: 1.8em; margin-top: 2em; page-break-before: always; }
h2 { font-size: 1.35em; }
h3 { font-size: 1.1em; font-style: italic; font-weight: normal; }
h4 { font-size: 1em; text-transform: uppercase; letter-spacing: 0.05em; }
p { margin: 0 0 0.8em 0; text-indent: 0; }
h1 + p::first-letter { font-size: 2.8em; font-weight: bold; float: left;
  line-height: 0.9; margin: 0.05em 0.08em 0 0; color: #6B4423; }
blockquote { margin: 1em 1.5em; padding-left: 0.9em; border-left: 3px solid #999;
  font-style: italic; color: #333; }
img { max-width: 100%; height: auto; display: block; margin: 1em auto;
  page-break-inside: avoid; }
table { border-collapse: collapse; margin: 1em auto; font-size: 0.92em; }
th, td { border: 1px solid #888; padding: 0.4em 0.7em; text-align: left; }
th { background: #f0f0f0; font-weight: bold; }
hr { border: none; text-align: center; margin: 2em 0; }
hr::before { content: "⁂"; font-size: 1.2em; color: #888; }
ul, ol { margin: 0.6em 0 0.8em 1.6em; padding: 0; }
li { margin-bottom: 0.3em; }
""", encoding="utf-8")

    print("  Running pandoc...")
    cmd = [
        "pandoc",
        str(md_file),
        "--from=markdown+yaml_metadata_block+smart",
        "--to=epub3",
        "-o", str(OUT_EPUB),
        f"--metadata-file={metadata}",
        f"--css={css}",
        "--toc-depth=2",
        "--split-level=1",
        f"--resource-path={WORK_DIR}",
    ]
    if COVER.exists():
        cmd.append(f"--epub-cover-image={COVER}")
    subprocess.run(cmd, check=True)

    if OUT_EPUB.exists():
        size_kb = OUT_EPUB.stat().st_size / 1024
        print(f"\n✓ {OUT_EPUB}")
        print(f"  {size_kb:,.0f} KB ({size_kb/1024:.2f} MB)")


if __name__ == "__main__":
    main()
