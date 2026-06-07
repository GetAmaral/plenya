#!/usr/bin/env python3
"""diff.py <figure_dir>

Compares <figure_dir>/preview.png with <figure_dir>/_reference.png
and writes <figure_dir>/diff.png — pixel difference (red highlight where mismatch).
Both images are reduced to grayscale before comparison.
"""
import sys
from pathlib import Path
import numpy as np
from PIL import Image

fig_dir = Path(sys.argv[1])
ref = np.array(Image.open(fig_dir / "_reference.png").convert("L"))
me  = np.array(Image.open(fig_dir / "preview.png").convert("L"))
if ref.shape != me.shape:
    me_pil = Image.open(fig_dir / "preview.png").convert("L").resize(ref.shape[::-1])
    me = np.array(me_pil)

# Mask: I am ink (dark) where original is paper (light), or vice versa
ref_ink = ref < 200
me_ink  = me  < 200

# Mismatch: pixels where one is ink and the other isn't
miss_in_me  = ref_ink & ~me_ink   # original has ink, mine doesn't (RED)
extra_in_me = me_ink & ~ref_ink   # mine has ink, original doesn't (BLUE)
match_ink   = ref_ink & me_ink

H, W = ref.shape
out = np.full((H, W, 3), 255, dtype=np.uint8)  # white background
out[match_ink]   = [0, 0, 0]       # both ink → black
out[miss_in_me]  = [255, 80, 80]   # missing → soft red
out[extra_in_me] = [80, 80, 255]   # extra → soft blue

Image.fromarray(out).save(fig_dir / "diff.png")
print(f"diff: ref-ink={ref_ink.sum()}  me-ink={me_ink.sum()}  match={match_ink.sum()}  miss={miss_in_me.sum()}  extra={extra_in_me.sum()}")
