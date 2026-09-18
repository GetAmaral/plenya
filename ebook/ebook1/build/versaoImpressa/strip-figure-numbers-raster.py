#!/usr/bin/env python3
"""Tira o número de dentro das quatro figuras que só existem em raster.

Companheiro de `strip-figure-numbers.py`, que cuida das vetoriais. Aqui não há
texto para redigir, só pixel, então:

  tarja    detecta o bloco colorido (ou escuro, no B&W) no canto superior
           esquerdo e repinta com o fundo amostrado linha a linha, para não
           deixar costura em figura com gradiente. O título fica onde está.
  inline   "Figura 2 — Da Estria Gordurosa ao Infarto": em vez de redesenhar
           texto (que exigiria acertar a fonte), RECORTA o bloco de pixels do
           título depois do travessão e cola na margem esquerda. Os glifos são
           os originais, não há fonte a adivinhar.

Idempotente pelo mesmo mecanismo: original em `_com-numero/`, sempre lido de lá.

Uso:  ./runpy strip-figure-numbers-raster.py [--dry]
"""
import sys, shutil
from pathlib import Path
import numpy as np
from PIL import Image

HERE = Path(__file__).resolve().parent
BOOK = HERE.parent.parent
BACKUP = "_com-numero"
DRY = "--dry" in sys.argv

# (arquivo, modo).  Os mesmos quatro desenhos existem em cor (figuras/pt-BR, com
# espaço no nome) e em B&W (figuras-bw, com underscore).
ALVOS = [
    (BOOK / "figuras/pt-BR/Cap02 Fig02.PNG", "inline"),
    (BOOK / "figuras/pt-BR/Cap05 Fig01.PNG", "tarja"),
    (BOOK / "figuras/pt-BR/Cap06 Fig01.PNG", "tarja"),
    (BOOK / "figuras/pt-BR/Cap06 Fig02.PNG", "tarja"),
    (HERE / "figuras-bw/Cap02_Fig02.png", "inline"),
    (HERE / "figuras-bw/Cap05_Fig01.png", "tarja"),
    (HERE / "figuras-bw/Cap06_Fig01.png", "tarja"),
    (HERE / "figuras-bw/Cap06_Fig02.png", "tarja"),
]


def backup(p: Path) -> Path:
    bdir = p.parent / BACKUP
    bdir.mkdir(exist_ok=True)
    o = bdir / p.name
    if not o.exists():
        shutil.copy2(p, o)
    return o


def text_rows(gray: np.ndarray, limit: int):
    """Faixas de linhas com pixel escuro, dentro dos primeiros `limit` pixels."""
    dark = (gray[:limit, :] < 128).sum(axis=1)
    runs, s = [], None
    for y, v in enumerate(dark):
        if v > 0 and s is None:
            s = y
        elif v == 0 and s is not None:
            runs.append((s, y - 1)); s = None
    if s is not None:
        runs.append((s, limit - 1))
    return runs


def preenche(arr: np.ndarray, gray: np.ndarray, y0: int, y1: int, x0: int, x1: int):
    """Repinta o retângulo com o fundo amostrado NAS MESMAS LINHAS, à direita.

    Interpolar entre a linha de cima e a de baixo não serve aqui: o cabeçalho
    dessas figuras fica sobre um painel cinza-claro que começa e termina dentro
    da faixa, e a amostra vinha do branco de fora do painel, deixando um
    retângulo mais claro no lugar da tarja. Amostrando na mesma altura, o tom
    do painel é o mesmo.
    """
    W = arr.shape[1]
    largura = 16
    janela = None
    x = x1 + 6
    while x + largura < W:
        if (gray[y0:y1 + 1, x:x + largura] < 170).sum() == 0:
            janela = (x, x + largura)
            break
        x += 4
    if janela is None:                     # nada limpo à direita: tenta à esquerda
        x = max(0, x0 - 6 - largura)
        janela = (x, x + largura)
    fundo = np.median(arr[y0:y1 + 1, janela[0]:janela[1]], axis=1)
    for i, y in enumerate(range(y0, y1 + 1)):
        arr[y, x0:x1 + 1] = fundo[i]


def do_tarja(src: Path, dst: Path):
    im = Image.open(src)
    arr = np.array(im)
    gray = np.array(im.convert("L"))
    H, W = gray.shape
    # a tarja é a PRIMEIRA faixa de pixels não-brancos da coluna da esquerda.
    # Sem separar por faixa, a janela pega também a primeira linha de rótulos do
    # gráfico, que às vezes começa dentro dos mesmos 22% de largura.
    strip = gray[:int(H * 0.30), :int(W * 0.22)]
    linhas = (strip < 200).sum(axis=1)
    faixas, s_ = [], None
    for y, v in enumerate(linhas):
        if v > 0 and s_ is None:
            s_ = y
        elif v == 0 and s_ is not None:
            faixas.append((s_, y - 1)); s_ = None
    if s_ is not None:
        faixas.append((s_, strip.shape[0] - 1))
    if not faixas:
        return "tarja não achada", None
    y0, y1 = faixas[0]
    # dentro da faixa, a tarja é o PRIMEIRO bloco de colunas; o título vem depois
    # de um vão. Sem cortar no vão, a repintura comeria o começo do título.
    faixa = gray[y0:y1 + 1, :int(W * 0.45)]
    xs = np.where((faixa < 200).sum(axis=0) > 0)[0]
    if len(xs) == 0:
        return "tarja não achada", None
    x0 = xs[0]; x1 = xs[0]
    for c in xs[1:]:
        if c - x1 > 12:
            break
        x1 = c
    if DRY:
        return "tarja", f"bloco x{x0}-{x1} y{y0}-{y1}"
    # folga também na vertical: a borda arredondada da tarja deixa uma linha
    # anti-serrilhada quase invisível uma linha acima do bloco detectado.
    pad = 3
    y0 = max(0, y0 - pad); y1 = min(arr.shape[0] - 1, y1 + pad)
    preenche(arr, gray, y0, y1, max(0, x0 - pad), min(arr.shape[1] - 1, x1 + pad))
    Image.fromarray(arr).save(dst)
    return "tarja", f"bloco x{x0}-{x1} y{y0}-{y1} repintado"


def do_inline(src: Path, dst: Path):
    im = Image.open(src)
    arr = np.array(im)
    gray = np.array(im.convert("L"))
    rows = text_rows(gray, 200)
    if not rows:
        return "título não achado", None
    y0, y1 = rows[0]
    band = gray[y0:y1 + 1, :]
    cols = np.where((band < 128).sum(axis=0) > 0)[0]
    # gaps de palavra: "Figura" | "2" | "—" | resto do título
    gaps, prev = [], cols[0]
    for c in cols[1:]:
        if c - prev > 10:
            gaps.append((prev, c))
        prev = c
    if len(gaps) < 3:
        return "prefixo não achado", None
    left = cols.min()
    rest_x0 = gaps[2][1]          # começo da palavra depois do travessão
    rest_x1 = cols.max()
    if DRY:
        return "inline", f"título y{y0}-{y1}; prefixo x{left}-{gaps[2][0]}; resto x{rest_x0}-{rest_x1} → x{left}"
    # o recorte usa o limiar duro (<128), mas a borda anti-serrilhada do último
    # glifo fica em tom médio e escapava da limpeza, deixando um risco solto na
    # coluna onde o título terminava. Recorta e limpa com folga.
    folga = 5
    b0 = max(0, rest_x0 - 2)
    solto = np.where((gray[y0:y1 + 1, :] < 235).sum(axis=0) > 0)[0]
    b1 = min(arr.shape[1] - 1, int(solto.max()) + folga)
    bloco = arr[y0:y1 + 1, b0:rest_x1 + 1 + 2].copy()
    preenche(arr, gray, y0, y1, max(0, left - folga), b1)
    largura = bloco.shape[1]
    arr[y0:y1 + 1, left:left + largura] = bloco
    Image.fromarray(arr).save(dst)
    return "inline", f"prefixo removido, título deslocado {b0 - left}px à esquerda"


def main():
    for p, modo in ALVOS:
        if not p.exists():
            print(f"  ⚠ {p} não existe"); continue
        o = backup(p)
        kind, det = (do_inline if modo == "inline" else do_tarja)(o, p)
        print(f"  {p.name:22s} {kind:20s} {det or ''}")
    if DRY:
        print("(--dry: nada foi escrito)")


if __name__ == "__main__":
    main()
