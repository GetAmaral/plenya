"""sidebyside.py <figure_dir> — gera <figure_dir>/side-by-side.png: original | meu | diff"""
import sys
from pathlib import Path
from PIL import Image

fd = Path(sys.argv[1])
orig = Image.open(fd / "_reference.png").convert("RGB")
me   = Image.open(fd / "preview.png").convert("RGB")
diff = Image.open(fd / "diff.png").convert("RGB")

W, H = orig.size
out = Image.new("RGB", (W * 3 + 40, H + 40), "white")
out.paste(orig, (0, 20))
out.paste(me,   (W + 20, 20))
out.paste(diff, (2 * W + 40, 20))
out.thumbnail((2400, 1600))
out.save(fd / "side-by-side.png")
print("saved")
