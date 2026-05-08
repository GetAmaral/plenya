#!/usr/bin/env python3
"""Build EPUB for 'BEFORE' (English edition).

Usage:
    python3 build-epub-en.py
"""

import os
import re
import sys
import shutil
import pypandoc
from pathlib import Path
from PIL import Image

# JPEG quality for PNG → JPG conversion in EPUB.
JPEG_QUALITY = 85

BOOK_ROOT = Path(__file__).resolve().parent.parent
LANG = "en"

# Paths
MD_DIR    = BOOK_ROOT / "md" / LANG
FIG_DIR   = BOOK_ROOT / "figuras" / LANG
COVER     = BOOK_ROOT / "capas" / LANG / "capa.jpg"
AUTHOR_PHOTO = BOOK_ROOT / "fotos" / "getulio_bw_halfbody_1000.jpg"
BUILD_DIR = BOOK_ROOT / "build"
WORK_DIR  = BUILD_DIR / f"work-{LANG}"
OUT_EPUB  = BUILD_DIR / "Before-en.epub"

READING_ORDER = [
    "frontmatter.md",                       # dedication + epigraph
    "00a-credits.md",                       # copyright
    "00b-introduction.md",                  # introduction
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
    """Remove YAML frontmatter block at start of file."""
    if text.startswith("---\n"):
        end = text.find("\n---\n", 4)
        if end != -1:
            return text[end + 5:]
    return text


def strip_meta_sections(text):
    """Extract only the published content using <!-- EPUB-START --> and <!-- EPUB-END --> markers.

    Rules:
      - If EPUB-START is present, the published content starts after it.
      - If EPUB-END is present, the published content ends before it.
      - If neither is present, the entire file is published (except YAML frontmatter, handled elsewhere).
    """
    start_marker = "<!-- EPUB-START -->"
    end_marker = "<!-- EPUB-END -->"
    start_idx = text.find(start_marker)
    if start_idx != -1:
        text = text[start_idx + len(start_marker):]
    end_idx = text.find(end_marker)
    if end_idx != -1:
        text = text[:end_idx]
    return text.strip() + "\n"


def rewrite_image_paths(text, fig_dir_name="images"):
    """Rewrite ../Cap01 Fig01.PNG  ->  images/Cap01_Fig01.jpg
    URL-encoded space (%20) and spaces both handled. Cross-platform case safe.
    PNG references are rewritten to .jpg to match the conversion done in prepare_images."""
    def replace_path(match):
        full = match.group(0)
        # Extract the path between ](  and  )
        path_match = re.search(r'\]\(([^)]+)\)', full)
        if not path_match:
            return full
        old_path = path_match.group(1)
        # Decode %20 and normalize
        decoded = old_path.replace("%20", " ")
        filename = os.path.basename(decoded)
        # Normalize: replace spaces with underscore
        safe_name = filename.replace(" ", "_")
        # Normalize extension — PNG becomes JPG (converted by prepare_images);
        # JPEG/JPG stays JPG; other rare extensions go lowercase.
        base, ext = os.path.splitext(safe_name)
        ext_lower = ext.lower()
        if ext_lower == ".png":
            safe_name = base + ".jpg"
        elif ext_lower in {".jpeg", ".jpg"}:
            safe_name = base + ".jpg"
        else:
            safe_name = base + ext_lower
        return full.replace(old_path, f"{fig_dir_name}/{safe_name}")
    # Match markdown image syntax  ![...](path)
    return re.sub(r'!\[[^\]]*\]\([^)]+\)', replace_path, text)


def prepare_images(images_out_dir):
    """Copy figures from figuras/<lang>/ to images_out_dir/, converting PNG → JPG
    for size reduction. Figures with transparency fall back to PNG to preserve alpha.

    Returns (count, bytes_saved).
    """
    images_out_dir.mkdir(parents=True, exist_ok=True)
    if not FIG_DIR.exists():
        print(f"  ⚠ {FIG_DIR} does not exist — skipping figures")
        return 0, 0
    count = 0
    bytes_in = 0
    bytes_out = 0
    for src in sorted(FIG_DIR.iterdir()):
        if not (src.is_file() and src.suffix.lower() in {".png", ".jpg", ".jpeg"}):
            continue
        base, ext = os.path.splitext(src.name)
        safe_base = base.replace(" ", "_")
        src_size = src.stat().st_size
        bytes_in += src_size

        if ext.lower() == ".png":
            # Try JPEG conversion; fall back to PNG if transparency is needed.
            with Image.open(src) as im:
                has_alpha = (im.mode in ("RGBA", "LA")) or (
                    im.mode == "P" and "transparency" in im.info
                )
                if has_alpha:
                    # Check if alpha is actually used (not just a flag)
                    if im.mode == "P":
                        im = im.convert("RGBA")
                    alpha = im.split()[-1]
                    alpha_min = alpha.getextrema()[0]
                    if alpha_min < 255:
                        # Has real transparency — keep as PNG
                        dst = images_out_dir / (safe_base + ".png")
                        shutil.copy2(src, dst)
                        bytes_out += dst.stat().st_size
                        count += 1
                        continue
                    # Fully opaque alpha channel — discard and convert
                    im = im.convert("RGB")
                elif im.mode != "RGB":
                    im = im.convert("RGB")
                dst = images_out_dir / (safe_base + ".jpg")
                im.save(dst, "JPEG", quality=JPEG_QUALITY, optimize=True)
                bytes_out += dst.stat().st_size
                count += 1
        else:
            # Already JPEG — copy as-is with normalized name
            dst = images_out_dir / (safe_base + ".jpg")
            shutil.copy2(src, dst)
            bytes_out += dst.stat().st_size
            count += 1

    # Also copy author photo
    if AUTHOR_PHOTO.exists():
        shutil.copy2(AUTHOR_PHOTO, images_out_dir / "autor.jpg")
    return count, bytes_in - bytes_out


def main():
    print(f"🛠  Building EPUB for language: {LANG}")

    if not MD_DIR.exists():
        sys.exit(f"❌ Markdown directory not found: {MD_DIR}")
    if not COVER.exists():
        print(f"⚠  Cover not found at {COVER} — EPUB will be built without cover.")

    # Clean work dir
    if WORK_DIR.exists():
        shutil.rmtree(WORK_DIR)
    WORK_DIR.mkdir(parents=True)

    # 1. Prepare images (with PNG → JPG conversion for size)
    print("📸 Processing images (PNG → JPG @ Q=" + str(JPEG_QUALITY) + ")...")
    img_count, bytes_saved = prepare_images(WORK_DIR / "images")
    mb_saved = bytes_saved / (1024 * 1024)
    print(f"   {img_count} figures processed — {mb_saved:.1f} MB saved vs. raw PNGs")

    # 2. Consolidate chapters
    print("📝 Consolidating chapters...")
    full_text = []
    included = 0
    for fname in READING_ORDER:
        src = MD_DIR / fname
        if not src.exists():
            print(f"   ⚠ Missing: {fname}")
            continue
        text = src.read_text(encoding="utf-8")
        text = strip_frontmatter(text)
        text = strip_meta_sections(text)
        text = rewrite_image_paths(text)
        full_text.append(text)
        full_text.append("\n\n")  # chapter separator
        included += 1
    print(f"   {included} chapters included")

    consolidated = "".join(full_text)
    md_file = WORK_DIR / "book.md"
    md_file.write_text(consolidated, encoding="utf-8")
    wc = len(consolidated.split())
    print(f"   {wc:,} words total")

    # 3. Write metadata YAML
    metadata = WORK_DIR / "metadata.yaml"
    metadata_yaml = """---
title: "BEFORE"
subtitle: "The silent window — a decade between normal and optimal where health is decided"
creator:
  - role: author
    text: "Dr. Getulio José Mattos do Amaral Filho"
publisher: "Author's Edition"
rights: "Copyright © 2026 Getulio José Mattos do Amaral Filho. All rights reserved."
language: en-US
date: "2026"
description: |
  Between the "normal" of conventional checkups and the optimal a body can achieve, there is
  a silent window in which disease grows without diagnosis. A nephrologist with 20+ years of
  clinical practice presents the expanded biomarker panel, the ACTS Method, and a quarterly
  Longevity Scorecard. Six real clinical cases run through the entire book.
subject:
  - "Preventive medicine"
  - "Longevity"
  - "Biomarkers"
  - "Health"
  - "Aging"
identifier:
  - scheme: ISBN
    text: "ISBN-EN-PENDING"
...
"""
    metadata.write_text(metadata_yaml, encoding="utf-8")

    # 4. Write stylesheet optimized for Kindle
    css = WORK_DIR / "style.css"
    css.write_text("""/* ANTES — CSS para EPUB Kindle */

@namespace epub "http://www.idpf.org/2007/ops";

body {
  font-family: "Georgia", "Times New Roman", serif;
  font-size: 1em;
  line-height: 1.55;
  margin: 0;
  padding: 0 0.6em;
  color: #111;
  text-align: justify;
  hyphens: auto;
}

h1, h2, h3, h4 {
  font-family: "Georgia", serif;
  font-weight: bold;
  text-align: left;
  page-break-after: avoid;
  margin-top: 1.4em;
  margin-bottom: 0.6em;
  hyphens: none;
}

h1 {
  font-size: 1.8em;
  margin-top: 2em;
  page-break-before: always;
}

h2 {
  font-size: 1.35em;
}

h3 {
  font-size: 1.1em;
  font-style: italic;
  font-weight: normal;
}

h4 {
  font-size: 1em;
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

p {
  margin: 0 0 0.8em 0;
  text-indent: 0;
}

/* Drop cap ornament (primeira letra do primeiro parágrafo de cada capítulo) */
h1 + p::first-letter {
  font-size: 2.8em;
  font-weight: bold;
  float: left;
  line-height: 0.9;
  margin: 0.05em 0.08em 0 0;
  color: #6B4423;
}

blockquote {
  margin: 1em 1.5em;
  padding-left: 0.9em;
  border-left: 3px solid #999;
  font-style: italic;
  color: #333;
}

em, i {
  font-style: italic;
}

strong, b {
  font-weight: bold;
}

img {
  max-width: 100%;
  height: auto;
  display: block;
  margin: 1em auto;
  page-break-inside: avoid;
}

figure {
  margin: 1.2em 0;
  page-break-inside: avoid;
}

figcaption {
  font-size: 0.88em;
  color: #555;
  text-align: center;
  margin-top: 0.4em;
  font-style: italic;
}

table {
  border-collapse: collapse;
  margin: 1em auto;
  font-size: 0.92em;
}

th, td {
  border: 1px solid #888;
  padding: 0.4em 0.7em;
  text-align: left;
}

th {
  background: #f0f0f0;
  font-weight: bold;
}

hr {
  border: none;
  text-align: center;
  margin: 2em 0;
}
hr::before {
  content: "⁂";
  font-size: 1.2em;
  color: #888;
}

ul, ol {
  margin: 0.6em 0 0.8em 1.6em;
  padding: 0;
}

li {
  margin-bottom: 0.3em;
}

/* Page breaks between chapters */
h1 {
  page-break-before: always;
}
""", encoding="utf-8")

    # 5. Run Pandoc
    # IMPORTANT: --toc omitido intencionalmente. Com EPUB3, o nav.xhtml é gerado
    # automaticamente e aparece no painel de navegação do Kindle/leitor — é isso
    # que o leitor usa para pular entre capítulos. O --toc adicionaria uma página
    # "Sumário" no corpo do EPUB, redundante e feia.
    # --toc-depth=2 continua válido para o nav.xhtml (Part → Chapter).
    print("⚙  Running Pandoc...")
    pandoc_args = [
        "--from=markdown+yaml_metadata_block+smart",
        "--to=epub3",
        f"--output={OUT_EPUB}",
        f"--metadata-file={metadata}",
        f"--css={css}",
        "--toc-depth=2",
        "--split-level=1",
        f"--resource-path={WORK_DIR}",
    ]
    if COVER.exists():
        pandoc_args.append(f"--epub-cover-image={COVER}")

    try:
        pypandoc.convert_file(
            str(md_file),
            "epub3",
            outputfile=str(OUT_EPUB),
            extra_args=pandoc_args,
        )
    except Exception as e:
        # pypandoc returns non-empty result by default; ignore if file exists
        if not OUT_EPUB.exists():
            raise
        print(f"   (pandoc warning: {e})")

    # 6. Report
    if OUT_EPUB.exists():
        size_kb = OUT_EPUB.stat().st_size / 1024
        print(f"\n✅ EPUB built: {OUT_EPUB}")
        print(f"   Size: {size_kb:,.0f} KB ({size_kb/1024:.2f} MB)")
    else:
        sys.exit("❌ EPUB not generated")


if __name__ == "__main__":
    main()
