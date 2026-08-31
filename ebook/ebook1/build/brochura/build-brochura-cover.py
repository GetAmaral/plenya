#!/usr/bin/env python3
"""Build print-ready cover for 'ANTES' — BROCHURA COM ORELHAS (gráfica nacional).

Usage:
    python3 build-brochura-cover.py pt-BR              # lombada padrão
    python3 build-brochura-cover.py pt-BR 20.4         # lombada em mm
    python3 build-brochura-cover.py pt-BR 20.4 80      # lombada + orelha em mm

Geometria (plano aberto, face externa para cima):

    [sangria] [ORELHA verso] [4ª capa] [lombada] [1ª capa] [ORELHA frente] [sangria]

    Trim por painel : 16 × 23 cm (mesmo formato da capa dura)
    Orelha          : 8 cm por padrão (faixa usual no Brasil: 7 a 10 cm; nunca
                      maior que a capa)
    Lombada         : depende do papel e do número de páginas. 20,4 mm é a
                      estimativa para 350 páginas em Pólen Bold 90 g/m²,
                      calibrada pelo número que a Fábrica do Livro devolveu
                      para a capa dura. O valor definitivo vem da gráfica.
    Sangria         : 3 mm em todos os lados

    Sem marcas de dobra no arquivo de impressão: a gráfica impõe. Para
    conferência visual, o script também grava um PNG *-guias.png com os
    vincos e a área de segurança desenhados.

Conteúdo (decisão editorial de 2026-08):
    ORELHA frente  → gancho/sinopse (texto de md/pt-BR/contracapa.md)
    ORELHA verso   → foto + bio do autor (versão de ~130 palavras)
    4ª capa        → sinopse longa + selo-slogan + código de barras
"""
from __future__ import annotations
import io
import sys
from pathlib import Path
from PIL import Image, ImageDraw, ImageFont
import img2pdf

THIS_DIR  = Path(__file__).resolve().parent
BUILD_DIR = THIS_DIR.parent
BOOK_ROOT = BUILD_DIR.parent
LANG = sys.argv[1] if len(sys.argv) > 1 else "pt-BR"
DEFAULT_SPINE_MM = 20.4
DEFAULT_FLAP_MM = 80.0
SPINE_MM = float(sys.argv[2]) if len(sys.argv) > 2 else DEFAULT_SPINE_MM
FLAP_MM = float(sys.argv[3]) if len(sys.argv) > 3 else DEFAULT_FLAP_MM

# --- Paths ---
CAPA_FRONT    = BOOK_ROOT / "capas" / LANG / "capa.jpg"
BARCODE       = BOOK_ROOT / "capas" / LANG / "isbn-barcode-978-65-02-07691-0.png"
AUTHOR_BW     = BOOK_ROOT / "fotos" / "getulio_bw_halfbody_fullres.jpg"
SYMBOL_GOLD   = BOOK_ROOT.parent.parent / "apps" / "site" / "public" / "brand" / "symbol" / "gold.png"
OUT_PDF       = THIS_DIR / f"Antes-{LANG}-brochura-capa.pdf"
OUT_PNG       = THIS_DIR / f"Antes-{LANG}-brochura-capa.png"
OUT_GUIDES    = THIS_DIR / f"Antes-{LANG}-brochura-capa-guias.png"

# --- Geometry (mm → px) ---
DPI = 600
def mm_to_px(mm): return round(mm * DPI / 25.4)

TRIM_W_MM = 160.0
TRIM_H_MM = 230.0
BLEED_MM  = 3.0

CANVAS_W_MM = BLEED_MM + FLAP_MM + TRIM_W_MM + SPINE_MM + TRIM_W_MM + FLAP_MM + BLEED_MM
CANVAS_H_MM = BLEED_MM + TRIM_H_MM + BLEED_MM
CANVAS_W = mm_to_px(CANVAS_W_MM)
CANVAS_H = mm_to_px(CANVAS_H_MM)

# X das divisas (px), da esquerda para a direita
FLAPB_X0 = mm_to_px(BLEED_MM)                      # orelha do verso (colada à 4ª capa)
FLAPB_X1 = FLAPB_X0 + mm_to_px(FLAP_MM)
BACK_X0  = FLAPB_X1
BACK_X1  = BACK_X0 + mm_to_px(TRIM_W_MM)
SPINE_X0 = BACK_X1
SPINE_X1 = SPINE_X0 + mm_to_px(SPINE_MM)
FRONT_X0 = SPINE_X1
FRONT_X1 = FRONT_X0 + mm_to_px(TRIM_W_MM)
FLAPF_X0 = FRONT_X1                                # orelha da frente
FLAPF_X1 = FLAPF_X0 + mm_to_px(FLAP_MM)

TRIM_Y0 = mm_to_px(BLEED_MM)
TRIM_Y1 = TRIM_Y0 + mm_to_px(TRIM_H_MM)

# --- Plenya palette ---
PETROL = (6, 59, 79)
GOLD   = (179, 134, 69)
CREAM  = (234, 231, 218)
INK    = (26, 26, 26)


def load_fonts():
    def first(paths):
        for p in paths:
            if Path(p).exists():
                return p
        raise FileNotFoundError(f"No font found among {paths}")
    return {
        "serif_regular": first([
            "/usr/share/texmf/fonts/opentype/public/tex-gyre/texgyrepagella-regular.otf",
            "/usr/share/fonts/truetype/dejavu/DejaVuSerif.ttf"]),
        "serif_italic": first([
            "/usr/share/texmf/fonts/opentype/public/tex-gyre/texgyrepagella-italic.otf",
            "/usr/share/fonts/truetype/dejavu/DejaVuSerif-Italic.ttf"]),
        "sans_bold": first([
            "/usr/share/fonts/opentype/inter/Inter-Bold.otf",
            "/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf"]),
        "sans": first([
            "/usr/share/fonts/opentype/inter/Inter-Regular.otf",
            "/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf"]),
    }

FONTS = load_fonts()

def f(role, size_pt):
    return ImageFont.truetype(FONTS[role], round(size_pt * DPI / 72))


def fit_into_box(img, max_w, max_h):
    img = img.copy()
    img.thumbnail((max_w, max_h), Image.LANCZOS)
    return img


def circular_crop(img, diameter):
    w, h = img.size
    side = min(w, h)
    left = (w - side) // 2
    top = (h - side) // 2
    img = img.crop((left, top, left + side, top + side)).resize((diameter, diameter), Image.LANCZOS)
    if img.mode != "RGBA":
        img = img.convert("RGBA")
    mask = Image.new("L", (diameter, diameter), 0)
    ImageDraw.Draw(mask).ellipse((0, 0, diameter, diameter), fill=255)
    out = Image.new("RGBA", (diameter, diameter), (0, 0, 0, 0))
    out.paste(img, (0, 0), mask)
    return out


def _draw_aligned(draw, line, x, y, w, font, fill, align):
    if align == "left":
        draw.text((x, y), line, font=font, fill=fill)
        return
    bbox = draw.textbbox((0, 0), line, font=font)
    tw = bbox[2] - bbox[0]
    dx = (w - tw) if align == "right" else (w - tw) // 2
    draw.text((x + dx, y), line, font=font, fill=fill)


def draw_wrapped(draw, text, x, y, w, font, fill, line_h_mult=1.35, align="left"):
    line = ""
    cur_y = y
    line_h = round(font.size * line_h_mult)
    for word in text.split():
        trial = (line + " " + word).strip()
        bbox = draw.textbbox((0, 0), trial, font=font)
        if (bbox[2] - bbox[0]) <= w or not line:
            line = trial
        else:
            _draw_aligned(draw, line, x, cur_y, w, font, fill, align)
            cur_y += line_h
            line = word
    if line:
        _draw_aligned(draw, line, x, cur_y, w, font, fill, align)
        cur_y += line_h
    return cur_y


# --- Copy -------------------------------------------------------------------
BLURB = (
    "Ricardo tinha 52 anos, fazia check-up todo ano, "
    "corria três vezes por semana. Oito meses depois "
    "do último exame \"dentro da normalidade\", caiu "
    "no estacionamento do escritório com um infarto.\n\n"
    "Entre o que o seu laboratório chama de normal e "
    "o que a ciência da longevidade chama de ótimo "
    "existe uma janela silenciosa de dez a vinte anos. "
    "É nela que a doença é construída, e é nela que "
    "ainda dá para mudar a trajetória.\n\n"
    "Em ANTES, Dr. Getulio Amaral Filho reúne vinte "
    "anos de prática clínica num mapa que o check-up "
    "convencional não desenha: 16 biomarcadores em 4 "
    "grupos que separam o normal do ótimo, o método "
    "AGIR de quatro pilares e a Regra dos Dois (a "
    "cada trimestre, dois focos com maior potencial "
    "de mover trajetória).\n\n"
    "Não é wellness. É medicina baseada em evidência, "
    "vinte anos de consultório, casos reais. Para "
    "quem ainda tem a década inteira pela frente, e "
    "para quem só tem a próxima."
)

# Orelha da frente — texto de md/pt-BR/contracapa.md ("Para a orelha do livro")
FLAP_FRONT = (
    "Se você tem 40, 50 ou 60 anos e não está "
    "disposto a esperar um diagnóstico para agir, "
    "este livro é para você.\n\n"
    "Não é self-help. Não é manual de receitas. É o "
    "trabalho de um médico que há décadas acompanha "
    "pacientes que, exatamente como você, recebem "
    "exames \"normais\" enquanto o corpo conta outra "
    "história."
)
FLAP_FRONT_KICKER = "A janela ainda está aberta."

# Orelha do verso — bio de ~130 palavras (md/pt-BR/sobre-o-autor.md)
FLAP_BACK_BIO = (
    "Médico nefrologista e clínico com mais de 20 anos "
    "de prática, formado em Medicina pela Universidade "
    "Estadual de Londrina em 2004. Especializou-se em "
    "Clínica Médica e em Nefrologia pela Santa Casa de "
    "Londrina, onde hoje coordena a Residência Médica de "
    "Nefrologia e fundou a Residência de Clínica Médica. "
    "É Responsável Técnico da DaVita Intra-hospitalar de "
    "Londrina e concluiu em 2026 pós-graduação em "
    "Medicina Funcional Integrativa pela ABMFI.\n\n"
    "É palestrante nacional em saúde, nefrologia e "
    "longevidade. Ao longo de duas décadas atendeu "
    "milhares de pacientes, e foi dessa prática que "
    "nasceram o Método AGIR e a convicção de que a saúde "
    "se decide na janela silenciosa entre o normal e o "
    "ótimo."
)
FLAP_BACK_FOOT = ["CRM-PR 21.876 · RQE 16.038",
                  "drgetulioamaralfilho.com.br",
                  "@drGetulioAmaralFilho"]


def draw_flap_front(canvas, draw):
    """Orelha da frente: gancho editorial."""
    safe = mm_to_px(8.0)
    x0 = FLAPF_X0 + safe
    x1 = FLAPF_X1 - mm_to_px(10.0)      # margem maior no lado que vai para a dobra externa
    w = x1 - x0
    y = TRIM_Y0 + mm_to_px(52.0)      # bloco curto: desce para equilibrar a mancha

    rule_w = mm_to_px(18)
    draw.line((x0, y, x0 + rule_w, y), fill=GOLD, width=mm_to_px(0.4))
    y += mm_to_px(9)

    for para in FLAP_FRONT.split("\n\n"):
        y = draw_wrapped(draw, para, x0, y, w, f("serif_regular", 10), CREAM, line_h_mult=1.5)
        y += mm_to_px(3.5)

    y += mm_to_px(6)
    draw.line((x0, y, x0 + rule_w, y), fill=GOLD, width=mm_to_px(0.4))
    y += mm_to_px(8)
    draw_wrapped(draw, FLAP_FRONT_KICKER, x0, y, w, f("serif_italic", 11.5), GOLD, line_h_mult=1.4)


def draw_flap_back(canvas, draw):
    """Orelha do verso: foto + bio do autor."""
    safe = mm_to_px(10.0)               # margem maior no lado da dobra externa
    x0 = FLAPB_X0 + safe
    x1 = FLAPB_X1 - mm_to_px(8.0)
    w = x1 - x0
    y = TRIM_Y0 + mm_to_px(22.0)

    dia = mm_to_px(46)
    photo = circular_crop(Image.open(AUTHOR_BW), dia)
    canvas.paste(photo, (x0 + (w - dia) // 2, y), photo)
    y += dia + mm_to_px(7)

    y = draw_wrapped(draw, "Dr. Getulio", x0, y, w, f("sans_bold", 12), CREAM, 1.25, align="center")
    y = draw_wrapped(draw, "Amaral Filho", x0, y, w, f("sans_bold", 12), CREAM, 1.25, align="center")
    y += mm_to_px(5)

    for para in FLAP_BACK_BIO.split("\n\n"):
        y = draw_wrapped(draw, para, x0, y, w, f("serif_regular", 8.6), CREAM, line_h_mult=1.45)
        y += mm_to_px(3)

    y = TRIM_Y1 - mm_to_px(26)
    draw.line((x0, y, x0 + mm_to_px(14), y), fill=GOLD, width=mm_to_px(0.35))
    y += mm_to_px(5)
    for ln in FLAP_BACK_FOOT:
        y = draw_wrapped(draw, ln, x0, y, w, f("sans", 7.5), GOLD, line_h_mult=1.45)


def draw_back_cover(canvas, draw):
    safe = mm_to_px(12.0)
    x0 = BACK_X0 + safe
    x1 = BACK_X1 - safe
    w = x1 - x0
    y = TRIM_Y0 + mm_to_px(20.0)

    for para in BLURB.split("\n\n"):
        y = draw_wrapped(draw, para, x0, y, w, f("serif_regular", 10.5), CREAM, line_h_mult=1.45)
        y += mm_to_px(4)

    # selo-slogan entre filetes dourados
    y += mm_to_px(6)
    rule_w = mm_to_px(36)
    rx = x0 + (w - rule_w) // 2
    draw.line((rx, y, rx + rule_w, y), fill=GOLD, width=mm_to_px(0.35))
    y += mm_to_px(7)
    for ln in ("O melhor momento para começar",
               "era dez anos atrás.",
               "O segundo melhor é hoje."):
        y = draw_wrapped(draw, ln, x0, y, w, f("serif_italic", 12), GOLD, 1.3, align="center")
    y += mm_to_px(6)
    draw.line((rx, y, rx + rule_w, y), fill=GOLD, width=mm_to_px(0.35))

    # código de barras + símbolo
    margin = mm_to_px(12.0)
    bc = Image.open(BARCODE).convert("RGB")
    bc_w = round(1.6 * DPI)
    bc = bc.resize((bc_w, round(bc_w * bc.height / bc.width)), Image.LANCZOS)
    pad = round(0.08 * DPI)
    px0 = BACK_X1 - margin - bc.width - 2 * pad
    py0 = TRIM_Y1 - margin - bc.height - 2 * pad
    draw.rectangle((px0, py0, px0 + bc.width + 2 * pad, py0 + bc.height + 2 * pad), fill=(255, 255, 255))
    canvas.paste(bc, (px0 + pad, py0 + pad))

    if SYMBOL_GOLD.exists():
        sym = fit_into_box(Image.open(SYMBOL_GOLD).convert("RGBA"), round(0.65 * DPI), round(0.65 * DPI))
        canvas.paste(sym, (BACK_X0 + margin, TRIM_Y1 - margin - sym.height), sym)


def draw_spine(canvas):
    spine_w = SPINE_X1 - SPINE_X0
    spine_h = TRIM_Y1 - TRIM_Y0
    if spine_w < round(0.25 * DPI):
        return
    strip = Image.new("RGB", (spine_h, spine_w), PETROL)
    sd = ImageDraw.Draw(strip)
    title_size = max(11, min(int(SPINE_MM * 0.95), 26))
    author_size = max(9, min(int(SPINE_MM * 0.55), 16))
    ft = f("sans_bold", title_size)
    fa = f("sans", author_size)
    bt = sd.textbbox((0, 0), "ANTES", font=ft)
    ba = sd.textbbox((0, 0), "Getulio Amaral Filho", font=fa)
    sd.text((round(0.6 * DPI), (spine_w - (bt[3] - bt[1])) // 2 - bt[1]), "ANTES", font=ft, fill=CREAM)
    sd.text((spine_h - (ba[2] - ba[0]) - round(0.6 * DPI),
             (spine_w - (ba[3] - ba[1])) // 2 - ba[1]), "Getulio Amaral Filho", font=fa, fill=GOLD)
    canvas.paste(strip.rotate(-90, expand=True), (SPINE_X0, TRIM_Y0))


def save_guides(canvas):
    """PNG de conferência com vincos, corte e área de segurança."""
    g = canvas.copy()
    d = ImageDraw.Draw(g)
    dash = mm_to_px(3)
    for x in (FLAPB_X1, BACK_X1, SPINE_X1, FRONT_X1):        # vincos de dobra
        for y in range(0, CANVAS_H, dash * 2):
            d.line((x, y, x, y + dash), fill=(255, 60, 60), width=mm_to_px(0.5))
    d.rectangle((FLAPB_X0, TRIM_Y0, FLAPF_X1, TRIM_Y1), outline=(60, 220, 120), width=mm_to_px(0.4))
    safe = mm_to_px(8)
    d.rectangle((FLAPB_X0 + safe, TRIM_Y0 + safe, FLAPF_X1 - safe, TRIM_Y1 - safe),
                outline=(255, 210, 80), width=mm_to_px(0.3))
    fnt = f("sans_bold", 14)
    d.text((FLAPB_X0 + mm_to_px(4), mm_to_px(1)), "vermelho: vinco  ·  verde: corte  ·  amarelo: segurança",
           font=fnt, fill=(255, 255, 255))
    g.save(OUT_GUIDES, "PNG", optimize=True)


def main():
    print(f"📗 Building BROCHURA COM ORELHAS cover: {LANG}")
    print(f"   Trim por painel : {TRIM_W_MM:.0f} × {TRIM_H_MM:.0f} mm")
    print(f"   Orelha          : {FLAP_MM:.0f} mm cada")
    print(f"   Lombada         : {SPINE_MM:.2f} mm")
    print(f"   Sangria         : {BLEED_MM:.0f} mm")
    print(f"   Canvas total    : {CANVAS_W_MM:.1f} × {CANVAS_H_MM:.1f} mm → {CANVAS_W} × {CANVAS_H} px @ {DPI} DPI")

    canvas = Image.new("RGB", (CANVAS_W, CANVAS_H), PETROL)

    # 1ª capa: a arte ocupa o trim e é estendida para a sangria por replicação de borda
    front_w = FRONT_X1 - FRONT_X0
    front_h = TRIM_Y1 - TRIM_Y0
    front = Image.open(CAPA_FRONT).convert("RGB").resize((front_w, front_h), Image.LANCZOS)
    canvas.paste(front, (FRONT_X0, TRIM_Y0))
    top = front.crop((0, 0, front_w, 8)).resize((front_w, TRIM_Y0), Image.LANCZOS)
    canvas.paste(top, (FRONT_X0, 0))
    bot = front.crop((0, front_h - 8, front_w, front_h)).resize((front_w, CANVAS_H - TRIM_Y1), Image.LANCZOS)
    canvas.paste(bot, (FRONT_X0, TRIM_Y1))

    draw = ImageDraw.Draw(canvas)
    draw_back_cover(canvas, draw)
    draw_flap_front(canvas, draw)
    draw_flap_back(canvas, draw)
    draw_spine(canvas)

    canvas.save(OUT_PNG, "PNG", optimize=True)
    save_guides(canvas)

    buf = io.BytesIO()
    canvas.save(buf, "PNG", optimize=True)
    buf.seek(0)
    layout = img2pdf.get_layout_fun(pagesize=(img2pdf.mm_to_pt(CANVAS_W_MM), img2pdf.mm_to_pt(CANVAS_H_MM)))
    with open(OUT_PDF, "wb") as out:
        out.write(img2pdf.convert(buf.read(), layout_fun=layout))

    print(f"\n✅ Capa (lossless, {DPI} DPI): {OUT_PDF}  —  {OUT_PDF.stat().st_size/1048576:.1f} MB")
    print(f"✅ Preview: {OUT_PNG}")
    print(f"✅ Conferência com guias: {OUT_GUIDES}")


if __name__ == "__main__":
    main()
