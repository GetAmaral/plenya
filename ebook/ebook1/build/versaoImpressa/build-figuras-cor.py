#!/usr/bin/env python3
"""Gera as figuras vetoriais COLORIDAS a partir das versões em cinza.

Por que recolorir o PDF em vez de reexecutar os geradores: das 33 figuras
vetoriais, 28 vêm dos scripts matplotlib em python-figures/ e 5 vêm do
pipeline v3 (JSON+SVG+Chromium), cujos .py estão desatualizados — reexecutar
os .py delas devolveria uma versão antiga. Trabalhando no PDF pronto, as 33
recebem o mesmo tratamento e nenhuma regride.

Como funciona: cada gerador codifica um papel semântico por tom de cinza
(em Cap14_Fig02, 0.94 = célula ótima, 0.88 = leve, 0.78 = média, 0.66 = alta).
O tom é, portanto, a chave semântica. Este script troca `X g` por `r g b rg`
apenas nos tons listados no mapa da figura; tom não listado sai idêntico ao
de hoje, então o default é "não mexer".

Uso:  build-figuras-cor.py [--only Cap14_Fig02] [--dry-run]
"""
import re
import sys
from pathlib import Path

import pikepdf

BW  = Path(__file__).resolve().parent / "figuras-bw"
COR = Path(__file__).resolve().parent / "figuras-cor"

# --- Mapa por figura: tom de cinza (2 casas) → cor -------------------------
# Tom ausente = fica cinza. Os tons escuros (0.00 texto, 0.23 texto suave,
# 0.33/0.40 eixo e rodapé) ficam SEMPRE neutros: são tipografia e malha, não
# codificam significado.
# --- Mapa por figura: tom de cinza (2 casas) → cor -------------------------
# MEDIDO, não escolhido: para cada tom, a figura B&W é renderizada isolando só
# aquele tom; as regiões que ele pinta são localizadas, alinhadas à arte
# original pela caixa de conteúdo, e a cor dominante da original naquela área é
# amostrada em centenas de pontos (analise-cor.py + gera-mapas.py).
#
# Só tons >= 0.60 entram: os mais escuros são tipografia, régua e marcador, e
# colori-los mudaria o texto. Tom ausente do mapa sai idêntico ao de hoje.
#
# As artes originais usam tints bem mais suaves do que se imagina — em
# Cap01_Fig01 a saturação das faixas é de 8 a 16 em 255. Por isso os valores
# abaixo parecem quase brancos: é o que a arte tem.
MAPS = {
    "Cap01_Fig01": {0.98: "#f0f3eb", 0.93: "#fef6eb", 0.85: "#fdeded"},
    "Cap01_Fig02": {0.96: "#e7eee3", 0.88: "#fef5e0", 0.75: "#f7f7f7"},
    "Cap02_Fig01": {0.84: "#e0e8d8"},
    "Cap02_Fig03": {0.93: "#f0f5ed"},
    "Cap03_Fig01": {0.93: "#e9f0e7"},
    "Cap04_Fig01": {0.96: "#e6ece2", 0.88: "#fef3db", 0.75: "#f2f2f2"},
    "Cap06_Fig04": {0.93: "#e8f0e8", 0.47: "#b81018"},
    "Cap07_Fig01": {0.6: "#939aa2"},
    "Cap07_Fig03": {0.75: "#707070", 0.54: "#d09020"},
    "Cap08_Fig01": {0.74: "#5c717c"},
    "Cap08_Fig02": {0.95: "#f0f0f0", 0.85: "#f8c080", 0.62: "#f0f0f0"},
    "Cap10_Fig01": {0.95: "#f3f5ee", 0.88: "#faefec"},
    "Cap10_Fig02": {0.97: "#f0f3e3", 0.93: "#fef6dc", 0.89: "#fceae5"},
    "Cap10_Fig03": {0.96: "#deeadc", 0.92: "#fef6da", 0.88: "#fcdfde"},
    "Cap11_Fig01": {0.96: "#eaf0e2", 0.93: "#fef3de", 0.85: "#f7e2de"},
    "Cap12_Fig02": {0.91: "#e0e8d8"},
    "Cap12_Fig03": {0.93: "#ebebed"},
    "Cap13_Fig01": {0.87: "#edeef1", 0.61: "#6b6f7b"},
    "Cap13_Fig02": {0.86: "#e1e2e2", 0.75: "#f1f2f7"},
    "Cap14_Fig02": {0.94: "#d8e8c8", 0.88: "#f8e898", 0.78: "#f8d888", 0.66: "#f0b0a8"},
    "Cap14_Fig03": {0.98: "#fef4db"},
    "Cap15_Fig01": {0.96: "#f0f0e8", 0.87: "#f8f8e0", 0.75: "#f8f0f0"},
    "Cap15_Fig02": {0.96: "#f6f8f5", 0.91: "#fefaf3", 0.86: "#f5f6fa", 0.73: "#fefaf3"},
}

# --- Tons que só ganham cor FORA de bloco de texto -------------------------
# São os escuros (linha do gráfico, marcador, seta) que na arte original não
# são pretos. Como compartilham o tom com a tipografia, só podem ser trocados
# onde não há letra — daí a separação.
# Primeiro token relevante depois de um operador de cor. `BT` abre bloco de
# texto e `Tj`/`TJ` mostram texto; `f`/`S`/`B` e afins pintam caminho e `Do`
# desenha um XObject (marcador). O que vier primeiro decide se aquela cor é de
# tipografia ou de gráfico.
PROX_TOKEN = re.compile(
    rb"(?<![A-Za-z0-9*])(BT|Tj|TJ|Do|f\*|f|F|S|s|B\*|B|b\*|b)(?![A-Za-z0-9*])")
TOKENS_TEXTO = {b"BT", b"Tj", b"TJ"}

# --- Tons que só ganham cor em TEXTO ---------------------------------------
# Usado quando o tom codifica uma série e não a tipografia do livro.
# --- Figuras que a versão colorida vem do PRÓPRIO gerador -------------------
# Nestas, o matplotlib agrupa elementos de cores diferentes num mesmo operador
# ("FASE 1" junto com o título; as réguas das duas colunas juntas), então o
# recolorir por tom não alcança. O gerador tem modo FIG_COR=1 e escreve direto
# em figuras-cor/ — este script não deve sobrescrevê-las.
#   python-figures/capNN-figNN.py  com FIG_COR=1
DO_GERADOR = {"Cap06_Fig03", "Cap03_Fig02", "Cap12_Fig01", "Cap08_Fig02"}


MAPS_TEXTO = {
    "Cap06_Fig04": {
        0.0: {16: "#003870", 19: "#003870", 22: "#003870"},
        0.33: {5: "#003870", 8: "#b81018"},
        0.23: {5: "#0a3a0a"},
    },
}


MAPS_GRAFICO = {
    "Cap01_Fig01": {0.0: "#003060"},
    "Cap01_Fig02": {0.0: "#183048"},
    "Cap04_Fig01": {0.0: "#082850"},
    "Cap05_Fig02": {0.0: "#b01010"},
    "Cap06_Fig04": {0.0: {5: "#003870", 7: "#003870", 8: "#003870",
                          10: "#003870", 30: "#b81018", 31: "#b81018"},
                    0.53: {0: "#003870", 2: "#b81018"}},
    "Cap07_Fig01": {0.0: {"f": "#e04020"}},
    "Cap07_Fig02": {0.0: "#001020"},
    "Cap07_Fig03": {0.0: "#103050", 0.23: "#306030"},
    "Cap08_Fig01": {0.0: {"*": "#003040", 3: "#c02820"}},
    "Cap09_Fig01": {0.0: {5: "#0c3ea1", 8: "#70a170", 11: "#d3a10c", 2: "#b00010"}},
    "Cap10_Fig01": {0.0: "#001020"},
    "Cap10_Fig02": {0.0: "#002030"},
    # Só o selo "FIGURA 3" (ocorrência 2) é vermelho. O vermelho cego no tom
    # 0.00 pintava também o eixo do tempo (ocorrência 17) e o cartão da
    # Fernanda, que na arte original são preto e verde.
    "Cap10_Fig03": {0.0: {2: "#b81010"}},
    "Cap11_Fig01": {0.0: "#001020"},
    "Cap11_Fig02": {
        0.62: {0: "#105830", 1: "#105830", 2: "#105830",
               3: "#083868", 4: "#083868", 5: "#083868",
               6: "#b86800", 7: "#b86800"},
        0.0: {3: "#105830", 4: "#083868", 5: "#b86800"},
    },
    "Cap12_Fig03": {0.0: {0: "#103050", 1: "#c00010"}},
    "Cap13_Fig01": {0.0: "#c01010", 0.1: "#c81818"},
    "Cap14_Fig02": {0.0: "#002040"},
    "Cap14_Fig03": {0.0: "#001030"},
}


def rgb(hexstr):
    h = hexstr.lstrip("#")
    return tuple(int(h[i:i + 2], 16) / 255 for i in (0, 2, 4))


def recolor_stream(data: bytes, mapping: dict, mapping_grafico: dict = None,
                   mapping_texto: dict = None) -> tuple[bytes, int]:
    """Troca operadores de cor cinza pelos RGB correspondentes.

    Três mapas, do mais amplo para o mais restrito:
      `mapping`          vale sempre (preenchimentos de área);
      `mapping_grafico`  vale só onde a cor pinta traçado, nunca letra;
      `mapping_texto`    vale só onde a cor pinta letra.

    A separação existe porque em algumas figuras o tom escuro é tipografia
    comum (e colori-lo estragaria o texto do livro) e em outras o tom É a
    codificação da série — em Cap06_Fig04 o tom 0.47 é a insulina inteira,
    curva e rótulos, e ali colorir o texto é justamente o certo.

    Qualquer entrada pode ser uma cor só, um dict por ocorrência
    ({"*": padrão, 3: "#outra"}) ou um dict por tipo de operador
    ({"f": preenchimento, "s": traço}). O índice de ocorrência conta TODOS os
    operadores daquele tom no stream, na ordem em que aparecem — é o mesmo
    índice que a sonda reporta.
    """
    hits = 0
    mapping_grafico = mapping_grafico or {}
    mapping_texto = mapping_texto or {}
    ocorrencia = {}

    def eh_grafico(pos):
        """True se a cor definida em `pos` vai pintar traçado, não letra."""
        m = PROX_TOKEN.search(data, pos)
        return True if not m else m.group(1) not in TOKENS_TEXTO

    def escolhe(chave, pos):
        """Qual mapa vale para este operador."""
        if chave in mapping:
            return mapping[chave]
        if eh_grafico(pos):
            return mapping_grafico.get(chave)
        return mapping_texto.get(chave)

    def concreta(cor, chave, is_fill):
        n = ocorrencia[chave]
        if isinstance(cor, dict) and ({"f", "s"} & set(cor)):
            return cor.get("f" if is_fill else "s")
        if isinstance(cor, dict):
            return cor.get(n, cor.get("*"))
        return cor

    def fmt(c, op):
        r, g, b = c
        return f"{r:.6g} {g:.6g} {b:.6g} {op}".encode()

    def aplica(m, chave, op, is_fill, pos):
        nonlocal hits
        # o contador avança para TODO operador do tom, mapeado ou não, para o
        # índice bater com o da sonda
        ocorrencia[chave] = ocorrencia.get(chave, 0)
        cor = escolhe(chave, pos)
        cor = concreta(cor, chave, is_fill) if cor is not None else None
        ocorrencia[chave] += 1
        if cor is None:
            return m.group(0)
        hits += 1
        return fmt(rgb(cor), op)

    def sub_gray(m):
        val, op = float(m.group(1)), m.group(2).decode()
        return aplica(m, round(val, 2), "rg" if op == "g" else "RG", op == "g", m.end())

    def sub_rgb(m):
        parts = [float(x) for x in m.group(1).split()]
        op = m.group(2).decode()
        if max(parts) - min(parts) > 0.04:
            return m.group(0)
        return aplica(m, round(sum(parts) / 3, 2), op, op == "rg", m.end())

    # O separador entre o número e o operador pode ser espaço OU quebra de
    # linha — o matplotlib usa os dois. Exigir espaço literal fazia parte dos
    # operadores passar batido; e como a cor de preenchimento é emitida duas
    # vezes seguidas, o cinza não convertido sobrescrevia a cor recém-posta.
    # ORDEM IMPORTA: os operadores RGB originais primeiro. Se o cinza fosse
    # convertido antes, a passada seguinte reprocessaria o que a primeira
    # acabou de escrever — e como as cores medidas são quase neutras, duas
    # faixas diferentes acabariam com a mesma cor.
    data = re.sub(rb"([\d.]+\s+[\d.]+\s+[\d.]+)\s+(rg|RG)\b", sub_rgb, data)
    data = re.sub(rb"([\d.]+)\s+(g|G)\b", sub_gray, data)
    return data, hits


def main():
    dry = "--dry-run" in sys.argv
    only = None
    if "--only" in sys.argv:
        only = sys.argv[sys.argv.index("--only") + 1]
    COR.mkdir(exist_ok=True)

    total_col = total_neutral = 0
    for src in sorted(BW.glob("Cap*.pdf")):
        name = src.stem
        if only and name != only:
            continue
        if name in DO_GERADOR:
            print(f"  {name:<14} pulada — vem do gerador (FIG_COR=1)")
            continue
        mapping = MAPS.get(name, {})
        mapping_grafico = MAPS_GRAFICO.get(name, {})
        mapping_texto = MAPS_TEXTO.get(name, {})
        pdf = pikepdf.open(src)
        hits = 0
        for page in pdf.pages:
            data = bytes(page.Contents.read_bytes())
            new, h = recolor_stream(data, mapping, mapping_grafico, mapping_texto)
            hits += h
            if h:
                page.Contents.write(new)
        dst = COR / src.name
        if not dry:
            pdf.save(dst)
        if hits:
            total_col += 1
            extra = f" +{len(mapping_grafico)} no gráfico" if mapping_grafico else ""
            print(f"  {name:<14} {len(mapping)} tom(ns) → cor{extra}, {hits} operadores")
        else:
            total_neutral += 1
            print(f"  {name:<14} neutra (sem área com significado de cor)")

    # os 4 rasters não têm override colorido: o build cai no PNG original
    print(f"\n{total_col} figuras coloridas, {total_neutral} mantidas neutras")
    print("As 4 raster (Cap02_Fig02, Cap05_Fig01, Cap06_Fig01, Cap06_Fig02) não")
    print("entram aqui — sem override em figuras-cor/, o build usa o PNG colorido.")


if __name__ == "__main__":
    main()
