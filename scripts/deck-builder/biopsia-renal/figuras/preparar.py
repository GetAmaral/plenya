"""Recorta o conteúdo das ilustrações e exporta JPEG quadrado para o card."""
from PIL import Image, ImageChops
import pathlib
HERE = pathlib.Path(__file__).parent
RAW, OUT = HERE / "raw", HERE.parent / "assets"
OUT.mkdir(parents=True, exist_ok=True)
BG = (255, 255, 255)
for f in sorted(RAW.glob("*.png")):
    im = Image.open(f).convert("RGB")
    dif = ImageChops.difference(im, Image.new("RGB", im.size, BG)).convert("L")
    caixa = dif.point(lambda p: 255 if p > 18 else 0).getbbox()
    if caixa:
        m = int(0.03 * max(im.size))
        im = im.crop((max(0, caixa[0] - m), max(0, caixa[1] - m),
                      min(im.width, caixa[2] + m), min(im.height, caixa[3] + m)))
    lado = max(im.size)
    tela = Image.new("RGB", (lado, lado), BG)
    tela.paste(im, ((lado - im.width) // 2, (lado - im.height) // 2))
    tela.resize((700, 700), Image.LANCZOS).save(OUT / f"{f.stem}.jpg", "JPEG", quality=92, optimize=True)
    print(f.stem)
