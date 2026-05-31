"""
Cap07 Fig01 (EN) — Extreme sedentary behavior raises death risk ~5x
Mirror of the PT version. White background, sans-serif, sampled colors.

Source: Mandsager et al., JAMA Network Open, 2018.
"""

from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams

# typography (matches PT — humanist sans)
rcParams["font.family"] = "sans-serif"
rcParams["font.sans-serif"] = ["Inter", "Open Sans", "Source Sans 3", "Source Sans Pro", "Roboto", "DejaVu Sans"]
rcParams["axes.unicode_minus"] = False

# ---------- palette (sampled directly from PT original) ----------
BG       = "#FEFEFE"   # near-white background (NOT cream)
TITLE    = "#071122"   # very dark navy (almost black)
SUBTITLE = "#5B6069"   # cool dark gray
BADGE    = "#5B6069"   # FIGURE 1 label — dark gray (NOT red)
RED      = "#E73F2B"   # vibrant red-orange highlight bar
GRAY_BAR = "#9399A1"   # cool/bluish gray for non-highlight bars
GRAY_REF = "#9399A1"   # dashed reference line + ref label
TICK     = "#5B6069"
FOOTNOTE = "#5B6069"

# ---------- data ----------
labels = [
    "Coronary heart disease",
    "Diabetes",
    "Smoking",
    "End-stage kidney disease",
    "Low cardiorespiratory fitness\n(extreme sedentary)",
]
values = [1.3, 1.4, 1.4, 2.0, 5.0]
highlight = [False, False, False, False, True]

# ---------- figure ----------
FIG_W, FIG_H = 9.0, 5.6
fig = plt.figure(figsize=(FIG_W, FIG_H), dpi=220)
fig.patch.set_facecolor(BG)

ax = fig.add_axes([0.34, 0.32, 0.62, 0.50])
ax.set_facecolor(BG)

# bars
bar_colors = [RED if h else GRAY_BAR for h in highlight]
bars = ax.barh(labels, values, color=bar_colors, height=0.55, zorder=3,
               edgecolor="none")

# value labels at right end of bars
for bar, val, h in zip(bars, values, highlight):
    color = RED if h else TITLE
    ax.text(val + 0.10, bar.get_y() + bar.get_height() / 2,
            f"{val:.1f}x",
            va="center", ha="left",
            fontsize=12 if h else 10.5,
            color=color, weight="bold", zorder=4)

# 1.0x reference dashed line
ax.axvline(1.0, color=GRAY_REF, linestyle=(0, (4, 3)), linewidth=1.0,
           zorder=2)

# axes cosmetics — no vertical gridlines, clean look
ax.set_xlim(0, 5.7)
ax.set_xlabel("Death risk (times higher)", fontsize=10, color=TICK,
              labelpad=8)
ax.tick_params(axis="x", colors=TICK, labelsize=9)
ax.tick_params(axis="y", colors=TITLE, labelsize=10.5, length=0, pad=6)
for spine in ("top", "right", "left"):
    ax.spines[spine].set_visible(False)
ax.spines["bottom"].set_color("#D6D8DC")
ax.set_xticks([0, 1, 2, 3, 4, 5])
ax.grid(False)

# bold + red on highlighted y label
for tick, h in zip(ax.get_yticklabels(), highlight):
    if h:
        tick.set_color(RED)
        tick.set_weight("bold")

# 1.0x reference marker (anchored under the dashed line)
# axes occupy x=0.34..0.96 fig coords, data x=0..5.7. So x=1.0 → 0.34 + (1.0/5.7)*0.62 ≈ 0.4488
ref_x = 0.34 + (1.0 / 5.7) * 0.62
fig.text(ref_x, 0.245, "1.0x",
         ha="center", va="top",
         fontsize=9, color=GRAY_REF, weight="bold")
fig.text(ref_x, 0.215, "Reference risk\n(people in excellent fitness)",
         ha="center", va="top",
         fontsize=7.8, color=GRAY_REF, style="italic")

# title block
fig.text(0.04, 0.945, "FIGURE 1",
         fontsize=9, color=BADGE, weight="bold")
fig.text(0.04, 0.895, "Extreme sedentary behavior raises death risk ~5x",
         fontsize=16, color=TITLE, weight="bold")
fig.text(0.04, 0.855, "Compared with other established risk factors",
         fontsize=10.5, color=SUBTITLE)

# callout
fig.text(0.5, 0.155,
         "More than smoking, diabetes, or heart disease.",
         fontsize=11, color=TITLE, weight="bold", ha="center", style="italic")

# source footer
source_lines = [
    "Source: Mandsager et al., JAMA Network Open, 2018. Study of 122,007 adults followed for a median of 8.4 years.",
    "Low fitness defined as below the 25th percentile for age and sex. Reference: elite group (above the 97.7th percentile).",
]
for i, line in enumerate(source_lines):
    fig.text(0.04, 0.060 - i * 0.022, line,
             fontsize=7.4, color=FOOTNOTE)

# save
out_path = Path(__file__).resolve().parents[2] / "figuras" / "en" / "Cap07 Fig01.PNG"
out_path.parent.mkdir(parents=True, exist_ok=True)
plt.savefig(out_path, dpi=220, facecolor=BG)
print(f"saved → {out_path}")
