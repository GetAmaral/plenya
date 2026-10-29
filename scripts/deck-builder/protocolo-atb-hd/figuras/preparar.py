"""Recorta as ilustrações e exporta PNG com fundo transparente para o card.
O fundo do modelo não é branco puro (~253), e mix-blend-mode não é confiável no PDF:
o quase-branco vira alfa, com rampa curta para não serrilhar o contorno."""
from PIL import Image, ImageChops
import pathlib
HERE = pathlib.Path(__file__).parent
RAW, OUT = HERE / "raw", HERE.parent / "assets"
OUT.mkdir(parents=True, exist_ok=True)
CLARO, RAMPA = 238, 10   # >= 248 vira transparente; abaixo de 238 fica opaco
for f in sorted(RAW.glob("*.png")):
    im = Image.open(f).convert("RGB")
    dif = ImageChops.difference(im, Image.new("RGB", im.size, (255, 255, 255))).convert("L")
    caixa = dif.point(lambda p: 255 if p > 18 else 0).getbbox()
    if caixa:
        m = int(0.02 * max(im.size))
        im = im.crop((max(0, caixa[0] - m), max(0, caixa[1] - m),
                      min(im.width, caixa[2] + m), min(im.height, caixa[3] + m)))
    lum = im.convert("L")
    alfa = lum.point(lambda v: 255 if v < CLARO else max(0, 255 - (v - CLARO) * 255 // RAMPA))
    rgba = im.convert("RGBA"); rgba.putalpha(alfa)
    rgba.thumbnail((800, 800), Image.LANCZOS)
    rgba.save(OUT / f"{f.stem}.png", optimize=True)
    print(f.stem, rgba.size)
