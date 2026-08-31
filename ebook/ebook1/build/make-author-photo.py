#!/usr/bin/env python3
"""Gera as fotos P&B do autor a partir do original em alta, sem perda de geração.

Fonte: docs/site/images/FotosHd/Getulio01.png (2148x3223, PNG sem perdas).
Recorte meio-corpo: (320, 0) + 1508x2020 — recorte nativo, sem upscale. As
coordenadas foram recuperadas por correlação com a versão aprovada em 2026-04,
então o enquadramento publicado é preservado exatamente.

Processamento:
  1. luminância ITU-R 709 sobre o PNG (não sobre um JPEG intermediário);
  2. histogram matching contra a versão aprovada, o que reproduz a curva
     editorial já validada sem precisar reconstruí-la à mão;
  3. unsharp discreto aplicado depois de cada redimensionamento.

Saídas (sobrescreve as duas, ambas versionadas no git):
  fotos/getulio_bw_halfbody_fullres.jpg  1508x2020  q96  -> impresso, capas, capa dura, PDF digital
  fotos/getulio_bw_halfbody_1000.jpg     1000x1339  q92  -> EPUB (payload menor na entrega Amazon)

Uso:
    ./versaoImpressa/runpy make-author-photo.py
"""

import shutil
import sys
from pathlib import Path

import numpy as np
from PIL import Image, ImageFilter

BUILD_DIR = Path(__file__).resolve().parent
BOOK_ROOT = BUILD_DIR.parent
REPO_ROOT = BOOK_ROOT.parent.parent

SOURCE = REPO_ROOT / "docs" / "site" / "images" / "FotosHd" / "Getulio01.png"
CROP = (320, 0, 320 + 1508, 0 + 2020)

FOTOS = BOOK_ROOT / "fotos"
OUT_FULL = FOTOS / "getulio_bw_halfbody_fullres.jpg"
OUT_WEB = FOTOS / "getulio_bw_halfbody_1000.jpg"
WEB_WIDTH = 1000

# Referência tonal: a própria versão aprovada, lida antes de ser sobrescrita.
TONE_REFERENCE = OUT_FULL


def to_gray_709(img: Image.Image) -> np.ndarray:
    a = np.asarray(img.convert("RGB"), dtype=np.float32) / 255.0
    lum = 0.2126 * a[..., 0] + 0.7152 * a[..., 1] + 0.0722 * a[..., 2]
    return np.clip(lum * 255.0, 0, 255).astype(np.uint8)


def match_histogram(src: np.ndarray, ref: np.ndarray) -> np.ndarray:
    s_vals, s_idx, s_cnt = np.unique(src.ravel(), return_inverse=True, return_counts=True)
    r_vals, r_cnt = np.unique(ref.ravel(), return_counts=True)
    s_q = np.cumsum(s_cnt).astype(np.float64)
    s_q /= s_q[-1]
    r_q = np.cumsum(r_cnt).astype(np.float64)
    r_q /= r_q[-1]
    mapped = np.interp(s_q, r_q, r_vals)
    return mapped[s_idx].reshape(src.shape).astype(np.uint8)


def sharpness(img: Image.Image) -> float:
    """Variância do laplaciano — proxy de detalhe, só para o relatório."""
    from numpy.lib.stride_tricks import sliding_window_view

    a = np.asarray(img.convert("L"), dtype=np.float32)
    k = np.array([[0, 1, 0], [1, -4, 1], [0, 1, 0]], dtype=np.float32)
    w = sliding_window_view(a, (3, 3))
    return float((w * k).sum(axis=(2, 3)).var())


def main() -> int:
    if not SOURCE.exists():
        print(f"❌ Original não encontrado: {SOURCE}")
        return 1
    if not TONE_REFERENCE.exists():
        print(f"❌ Referência tonal não encontrada: {TONE_REFERENCE}")
        return 1

    print(f"📷 Original: {SOURCE.name}", Image.open(SOURCE).size)

    before_full = Image.open(OUT_FULL); before_full.load()
    before_web = Image.open(OUT_WEB); before_web.load()
    ref_tone = np.asarray(Image.open(TONE_REFERENCE).convert("L"))

    # backup da versão anterior, caso algo precise ser comparado depois
    for path in (OUT_FULL, OUT_WEB):
        shutil.copy2(path, path.with_suffix(".jpg.bak"))

    gray = to_gray_709(Image.open(SOURCE).crop(CROP))
    toned = Image.fromarray(match_histogram(gray, ref_tone), "L")

    full = toned.filter(ImageFilter.UnsharpMask(radius=1.2, percent=55, threshold=3))
    full.save(OUT_FULL, quality=96, subsampling=0, dpi=(300, 300), optimize=True)

    web_h = round(toned.height * WEB_WIDTH / toned.width)
    web = toned.resize((WEB_WIDTH, web_h), Image.LANCZOS)
    web = web.filter(ImageFilter.UnsharpMask(radius=0.8, percent=58, threshold=3))
    web.save(OUT_WEB, quality=92, subsampling=0, dpi=(300, 300), optimize=True)

    for label, before, after, path in (
        ("fullres", before_full, Image.open(OUT_FULL), OUT_FULL),
        ("1000px ", before_web, Image.open(OUT_WEB), OUT_WEB),
    ):
        print(
            f"   {label}: {after.size[0]}x{after.size[1]}"
            f"  {path.stat().st_size/1024:6.0f} KB"
            f"  detalhe {sharpness(before):7.1f} → {sharpness(after):7.1f}"
        )

    print("\n✅ Fotos regeneradas. Backups em *.jpg.bak (apagar quando aprovado).")
    print("   Rebuildar: ./versaoImpressa/runpy versaoImpressa/build-print-pdf.py pt-BR")
    return 0


if __name__ == "__main__":
    sys.exit(main())
