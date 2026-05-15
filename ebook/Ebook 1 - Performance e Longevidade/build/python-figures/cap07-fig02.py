"""
Cap07 Fig02 (EN) — The biggest return on investment is in the first step.
Adjusted all-cause mortality risk by cardiorespiratory fitness percentile.

Source: Mandsager et al., JAMA Network Open, 2018.
"""

from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams

rcParams["font.family"] = "sans-serif"
rcParams["font.sans-serif"] = ["Inter", "Open Sans", "Source Sans 3", "Roboto", "DejaVu Sans"]
rcParams["axes.unicode_minus"] = False

# ---------- palette (sampled from PT) ----------
BG          = "#FEFEFE"
TITLE       = "#071122"
SUBTITLE    = "#5B6069"
BADGE       = "#5B6069"
CURVE_BLUE  = "#0C386C"   # navy curve + markers
GREEN_FILL  = "#F5F8F2"   # subtle green band near 1.0x
GREEN_TEXT  = "#2E5131"
RED_TEXT    = "#B01C23"
TICK        = "#5B6069"
FOOTNOTE    = "#5B6069"

# ---------- data ----------
xs = [1, 2, 3, 4, 5]
risk = [5.04, 2.10, 1.49, 1.29, 1.00]
xlabels = [
    "Low\n(<P25)",
    "Below average\n(P25–P49)",
    "Above average\n(P50–P74)",
    "High\n(P75–P97.6)",
    "Elite\n(≥P97.7)",
]

# ---------- figure ----------
fig = plt.figure(figsize=(10.0, 6.2), dpi=220)
fig.patch.set_facecolor(BG)

ax = fig.add_axes([0.10, 0.27, 0.82, 0.52])
ax.set_facecolor(BG)

# ---------- subtle green band at low-risk zone (1.0x reference) ----------
ax.axhspan(0.5, 1.5, color=GREEN_FILL, zorder=1)

# ---------- 1.41x smoking reference dashed line ----------
ax.axhline(1.41, color=RED_TEXT, linewidth=0.9,
           linestyle=(0, (4, 3)), zorder=2, alpha=0.7)
ax.text(4.6, 1.50, "Smoking-equivalent\nrisk (~1.41x)",
        fontsize=8.3, color=RED_TEXT, style="italic",
        ha="left", va="bottom", zorder=5)

# ---------- 1.0x reference dashed line ----------
ax.axhline(1.0, color="#9AA0A6", linewidth=0.8,
           linestyle=(0, (4, 3)), zorder=2)

# ---------- main descending curve ----------
ax.plot(xs, risk,
        color=CURVE_BLUE, linewidth=2.4, zorder=4,
        marker="o", markersize=10,
        markerfacecolor="white",
        markeredgecolor=CURVE_BLUE, markeredgewidth=2.2)

# value labels above each point
for x, y in zip(xs, risk):
    ax.text(x, y + 0.18, f"{y:.2f}",
            ha="center", va="bottom",
            fontsize=11, color=TITLE, weight="bold", zorder=6)

# ---------- "Biggest gain here" annotation ----------
ax.annotate("Biggest gain here",
            xy=(1.5, 3.5), xytext=(1.5, 4.6),
            fontsize=10, color=GREEN_TEXT, weight="bold",
            ha="center", va="center", style="italic",
            arrowprops=dict(arrowstyle="->", color=GREEN_TEXT,
                            linewidth=1.2),
            zorder=7)

# ---------- axes cosmetics ----------
ax.set_xlim(0.5, 5.5)
ax.set_ylim(0.5, 5.6)
ax.set_xticks(xs)
ax.set_xticklabels(xlabels)
ax.set_yticks([1, 2, 3, 4, 5])
ax.set_yticklabels(["1.0", "2.0", "3.0", "4.0", "5.0"])
ax.tick_params(axis="x", colors=TICK, labelsize=9, length=0, pad=8)
ax.tick_params(axis="y", colors=TICK, labelsize=9, length=0, pad=4)
for spine in ("top", "right"):
    ax.spines[spine].set_visible(False)
ax.spines["left"].set_color("#D6D8DC")
ax.spines["bottom"].set_color("#D6D8DC")

# axis labels
ax.set_xlabel("Cardiorespiratory fitness level (by percentile)",
              fontsize=9.5, color=TICK, labelpad=12)
fig.text(0.10, 0.805, "Death risk (times higher)",
         fontsize=9, color=TICK, weight="bold", ha="left")

# ---------- title block ----------
fig.text(0.04, 0.945, "FIGURE 2",
         fontsize=9, color=BADGE, weight="bold")
fig.text(0.04, 0.895,
         "The biggest return on investment is in the first step.",
         fontsize=15.5, color=TITLE, weight="bold")
fig.text(0.04, 0.855,
         "Adjusted all-cause mortality risk by cardiorespiratory fitness.",
         fontsize=10, color=SUBTITLE)

# ---------- source footer ----------
source_lines = [
    "The largest drop occurs between the lowest-fitness group (<P25) and the next — equivalent to meeting the basic guideline of",
    "150 minutes of moderate activity per week.",
    "Source: Mandsager et al., JAMA Network Open, 2018.",
]
for i, line in enumerate(source_lines):
    fig.text(0.04, 0.078 - i * 0.022, line,
             fontsize=7.4, color=FOOTNOTE)

# save
out_path = Path(__file__).resolve().parents[2] / "figuras" / "en" / "Cap07 Fig02.PNG"
out_path.parent.mkdir(parents=True, exist_ok=True)
plt.savefig(out_path, dpi=220, facecolor=BG)
print(f"saved → {out_path}")
