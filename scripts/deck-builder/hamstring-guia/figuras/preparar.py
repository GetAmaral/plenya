"""Recorta o conteudo, normaliza o enquadramento e exporta os JPEG que o guia.html usa.
Quadrado (900) para os cartoes de 3 colunas, 3:2 (1100x733) para os de 2 colunas."""
from PIL import Image, ImageChops
import pathlib
HERE = pathlib.Path(__file__).parent
RAW, OUT = HERE / "raw", HERE.parent / "assets" / "fig"
OUT.mkdir(parents=True, exist_ok=True)
BG = (245, 241, 232)
QUADRADO = {"d1", "d2", "d3", "e5", "e6", "e7"}      # cartoes de 3 colunas
TRES_DOIS = {"e1", "e2", "e3", "e4", "e8", "e9", "e10"}

def preparar(nome, tw, th):
    im = Image.open(RAW / f"{nome}.png").convert("RGB")
    dif = ImageChops.difference(im, Image.new("RGB", im.size, BG)).convert("L")
    caixa = dif.point(lambda p: 255 if p > 18 else 0).getbbox()
    if caixa:
        m = int(0.025 * max(im.size))
        caixa = (max(0, caixa[0] - m), max(0, caixa[1] - m),
                 min(im.width, caixa[2] + m), min(im.height, caixa[3] + m))
        im = im.crop(caixa)
    w, h = im.size
    if w / h > tw / th: nw, nh = tw, max(1, round(tw * h / w))
    else:               nh, nw = th, max(1, round(th * w / h))
    im = im.resize((nw, nh), Image.LANCZOS)
    tela = Image.new("RGB", (tw, th), BG)
    tela.paste(im, ((tw - nw) // 2, (th - nh) // 2))
    # JPEG de proposito: o Chromium repassa o stream direto para o PDF e o arquivo fica leve.
    tela.save(OUT / f"{nome}.jpg", "JPEG", quality=92, optimize=True, subsampling=0)
    print(nome, tela.size)

for n in sorted(QUADRADO):  preparar(n, 900, 900)
for n in sorted(TRES_DOIS): preparar(n, 1100, 733)

# O esquema anatômico entra inteiro: os dois painéis já vêm na mesma escala e
# o recorte automático quebraria o alinhamento entre eles. As setas de pressão
# são um overlay SVG no guia.html, posicionado em coordenadas desta imagem.
if (RAW / "anat.png").exists():
    Image.open(RAW / "anat.png").convert("RGB").save(
        OUT / "anat.jpg", "JPEG", quality=92, optimize=True, subsampling=0)
    print("anat (inteiro)")
