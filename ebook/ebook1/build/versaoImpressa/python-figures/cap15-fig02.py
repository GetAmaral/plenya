"""
Cap15 Fig02 (PT-BR, B&W vetorial) — Rastreamento por década.
3 colunas (40s | 50s | 60s+) com 6 seções estruturadas por coluna.
"""
from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle, FancyBboxPatch

rcParams["font.family"] = "sans-serif"
rcParams["font.sans-serif"] = ["Inter", "Open Sans", "DejaVu Sans"]
rcParams["axes.unicode_minus"] = False
rcParams["pdf.fonttype"] = 42
rcParams["ps.fonttype"] = 42

BG       = "#FFFFFF"
INK      = "#000000"
INK_SOFT = "#3A3A3A"
TICK     = "#555555"
FOOT     = "#666666"
HEAD_BG  = "#3A3A3A"
COL_1    = "#F4F4F4"
COL_2    = "#E8E8E8"
COL_3    = "#DCDCDC"
BAND     = "#EDEDED"

fig = plt.figure(figsize=(11.0, 12.5))
fig.patch.set_facecolor(BG)

LEFT = 0.025

fig.text(LEFT, 0.970,
         "Figura 2 — Rastreamento por década — exames e avaliações que se acumulam a cada janela etária.",
         fontsize=13, color=INK, weight="bold")
fig.text(LEFT, 0.952,
         "Cada década acrescenta exames e avaliações à lista anterior, sem subotimizá-los.",
         fontsize=9, color=INK_SOFT, style="italic")

# 3 colunas
COL_X = [0.030, 0.345, 0.660]
COL_W = 0.305

# Header colorido das colunas
HEAD_Y_TOP = 0.920
HEAD_Y_BOT = 0.880
col_heads = [
    ("REVERSIBILIDADE", "40 a 49 anos"),
    ("INTERVENÇÃO",     "50 a 59 anos"),
    ("PRESERVAÇÃO",     "60 anos em diante"),
]
for i, (head, sub) in enumerate(col_heads):
    x = COL_X[i]
    fig.patches.append(Rectangle(
        (x, HEAD_Y_BOT), COL_W, HEAD_Y_TOP - HEAD_Y_BOT,
        facecolor=HEAD_BG, edgecolor="none",
        transform=fig.transFigure, zorder=2
    ))
    fig.text(x + COL_W/2, HEAD_Y_TOP - 0.012, head,
             fontsize=11, color="white", weight="bold", ha="center", va="center")
    fig.text(x + COL_W/2, HEAD_Y_BOT + 0.010, sub,
             fontsize=10, color="white", weight="bold", ha="center", va="center")

# Background colunas
col_bg = [COL_1, COL_2, COL_3]
BODY_TOP = 0.875
BODY_BOT = 0.090
for i in range(3):
    fig.patches.append(Rectangle(
        (COL_X[i], BODY_BOT), COL_W, BODY_TOP - BODY_BOT,
        facecolor=col_bg[i], edgecolor="none",
        transform=fig.transFigure, zorder=0
    ))

# --- Quebra de linha medida -------------------------------------------------
# fig.text não quebra linha sozinho: as quebras só existem onde alguém digitou
# \n no conteúdo. Foi assim que as frases de "FOCO DA DÉCADA" saíram da coluna
# (a de 60+ passava 5,1 mm da borda na página impressa; a de 50-59 ficava com
# 8 px de margem contra os ~35 px das demais). As funções abaixo medem a
# largura real com o renderer e quebram só o que não couber, preservando as
# quebras manuais já existentes.

TEXT_PAD = 0.012                      # recuo interno usado no x das seções
USABLE_W = COL_W - 2 * TEXT_PAD       # largura útil da coluna, fração da figura


def _renderer():
    fig.canvas.draw()
    return fig.canvas.get_renderer()


def text_width(s, fontsize, weight="normal"):
    """Largura de `s` em fração da largura da figura."""
    probe = fig.text(0, 0, s, fontsize=fontsize, weight=weight)
    bb = probe.get_window_extent(renderer=_renderer())
    probe.remove()
    return bb.width / fig.bbox.width


def wrap_to_column(text, fontsize, weight="normal", max_w=USABLE_W):
    """Quebra as linhas que estouram `max_w`, mantendo as quebras manuais.

    Continuação de item de lista ("• ...") herda o recuo de dois espaços que o
    conteúdo já usa, pra alinhar sob o texto do bullet e não sob a bolinha.
    """
    out = []
    for line in text.split("\n"):
        if text_width(line, fontsize, weight) <= max_w:
            out.append(line)
            continue
        stripped = line.lstrip()
        lead = line[: len(line) - len(stripped)]
        cont = "  " if stripped.startswith("•") else lead
        cur, prefix = "", lead
        for word in stripped.split():
            cand = word if not cur else f"{cur} {word}"
            if cur and text_width(prefix + cand, fontsize, weight) > max_w:
                out.append(prefix + cur)
                prefix, cur = cont, word
            else:
                cur = cand
        if cur:
            out.append(prefix + cur)
    return "\n".join(out)


# Conteúdo em texto único por coluna, com seções marcadas
content = [
    # Coluna 1: 40-49
    [
        ("FOCO DA DÉCADA",
         "Detecção precoce. Identificar fatores de risco antes do dano."),
        ("O QUE ENTRA NESTA DÉCADA",
         "• Painel ampliado anual (ApoB, PCR-us, HbA1c,\n  insulina jejum, homocisteína, NT-proBNP,\n  ácido úrico, ferritina, TFGe, microalbuminúria)\n• CAC\n• MTHFR, APOE\n• Composição corporal (bioimpedância,\n  DEXA bianual)\n• Tireoide completo (TSH, T4 livre, anti-TPO)\n• Colonoscopia a partir dos 45 anos\n• Mamografia a partir de 45 anos (M)\n• PSA basal a partir de 45 anos (H)"),
        ("DIFERENCIAIS DO FOCO",
         "Maior impacto das intervenções: cada\nmudança de hábito ainda reverte trajetória."),
        ("O QUE SE ACUMULA",
         "(início da lista)"),
        ("PACIENTE-ÂNCORA",
         "Fernanda, 41 anos: rastreio detectou\nalterações sutis 5 anos antes do que o\ncheck-up básico veria."),
    ],
    # Coluna 2: 50-59
    [
        ("FOCO DA DÉCADA",
         "Ação máxima. Intervir intensamente para reduzir risco e impacto."),
        ("O QUE ENTRA NESTA DÉCADA",
         "• CAC novamente (se primeira década\n  não fez)\n• TC tórax baixa dose (DR 80kg, ex-\n  tabagistas, fumantes ou exposição prévia)\n• Endoscopia digestiva alta\n• Painel hormonal masculino (H):\n  testosterona, DHEA-S, SHBG\n• Avaliação cognitiva basal\n• Audiometria\n• Avaliação ginecológica (M): USG\n  transvaginal (peri-menopausa)\n• Avaliação inicial óssea — DEXA"),
        ("DIFERENCIAIS DO FOCO",
         "Maior impacto das estatinas, CPAP,\ntreino de força. Janela curta\npara reverter."),
        ("O QUE SE ACUMULA",
         "Tudo do anterior continua."),
        ("PACIENTE-ÂNCORA",
         "Marcos, 57 anos: CAC 412 foi seu\nmarcador estrutural de mudança\nde era da vida."),
    ],
    # Coluna 3: 60+
    [
        ("FOCO DA DÉCADA",
         "Preservação funcional. Manter capacidade, autonomia e qualidade de vida."),
        ("O QUE ENTRA NESTA DÉCADA",
         "• Avaliação funcional (handgrip, SPPB,\n  marcha cronometrada)\n• Composição corporal (DEXA)\n• Avaliação cognitiva periódica\n  (MoCA, p-tau217 quando indicado)\n• Avaliação geriátrica funcional\n• Revisão de medicações com\n  desprescrição planejada\n• Avaliação de risco de quedas\n• Vacinas atualizadas (zóster,\n  pneumocócica, COVID, gripe)"),
        ("DIFERENCIAIS DO FOCO",
         "Preservar independência, mobilidade,\nfunção. Prevenir fragilidade."),
        ("O QUE SE ACUMULA",
         "Tudo dos blocos anteriores."),
        ("PACIENTE-ÂNCORA",
         "A década que mostra se o programa\nfuncionou."),
    ],
]

# Renderizar cada coluna com seções empilhadas
SECTION_START_Y = BODY_TOP - 0.020
SECTION_GAP = 0.012

for col_idx, col_content in enumerate(content):
    x = COL_X[col_idx] + 0.012
    y_current = SECTION_START_Y

    for section_idx, (section_title, section_text) in enumerate(col_content):
        # Header da seção
        fig.text(x, y_current, section_title,
                 fontsize=7.5, color=INK, weight="bold",
                 va="top")
        y_current -= 0.018

        # Texto da seção (quebrado à largura da coluna antes de medir a altura)
        section_text = wrap_to_column(section_text, 7.5)
        n_lines = section_text.count("\n") + 1
        fig.text(x, y_current, section_text,
                 fontsize=7.5, color=INK, va="top",
                 linespacing=1.3)
        y_current -= n_lines * 0.012 + SECTION_GAP

        # Linha separadora entre seções
        if section_idx < len(col_content) - 1:
            fig.lines.append(plt.Line2D(
                [COL_X[col_idx] + 0.008, COL_X[col_idx] + COL_W - 0.008],
                [y_current + 0.005, y_current + 0.005],
                color="#BBBBBB", linewidth=0.4,
                transform=fig.transFigure, zorder=2
            ))

# Footer
fig.text(LEFT, 0.060,
         "M = mulheres; H = homens; CAC = cálcio coronariano; ASMI = índice de massa muscular apendicular.",
         fontsize=7, color=FOOT, style="italic")
fig.text(LEFT, 0.040,
         "A lógica: cada década adiciona exames à lista anterior — ela se aprofunda, não diminui.",
         fontsize=8, color=INK, weight="bold", style="italic")

# --- Verificação: nenhum texto pode passar da borda da sua coluna -----------
# Roda antes de salvar. Se falhar, o conteúdo mudou e a figura NÃO é gravada —
# melhor quebrar o build do que mandar pro miolo uma caixa estourada.
_r = _renderer()
_falhas = []
for _t in fig.texts:
    _bb = _t.get_window_extent(renderer=_r)
    _x0, _x1 = _bb.x0 / fig.bbox.width, _bb.x1 / fig.bbox.width
    for _i, _cx in enumerate(COL_X):
        if _cx <= _x0 < _cx + COL_W:
            _limite = _cx + COL_W - TEXT_PAD
            if _x1 > _limite + 1e-4:
                _falhas.append((_i + 1, _t.get_text().split("\n")[0][:52],
                                (_x1 - _limite) * fig.get_size_inches()[0] * 25.4))
if _falhas:
    for _c, _txt, _mm in _falhas:
        print(f"  ✗ coluna {_c}: estoura {_mm:.1f} mm — {_txt!r}")
    raise SystemExit("Texto fora da coluna — ajuste o conteúdo ou COL_W.")
print("✓ bbox: todo texto dentro das colunas")

out_dir = Path(__file__).resolve().parents[1] / "figuras-bw"
out_dir.mkdir(parents=True, exist_ok=True)
pdf_path = out_dir / "Cap15_Fig02.pdf"
png_path = out_dir / "_preview_Cap15_Fig02.png"
plt.savefig(pdf_path, facecolor=BG, bbox_inches="tight", pad_inches=0.15)
plt.savefig(png_path, dpi=160, facecolor=BG, bbox_inches="tight", pad_inches=0.15)
print(f"saved → {pdf_path}")
print(f"preview → {png_path}")
