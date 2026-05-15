"""
Cap01 Fig01 (EN) — Three 'normal' check-ups — a dangerous trend.
Ricardo's HbA1c trajectory across four check-ups.

Mirrors PT version exactly. Colors sampled from PT original.
NO figure badge — PT version has none.
"""

from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams

rcParams["font.family"] = "sans-serif"
rcParams["font.sans-serif"] = ["Inter", "Open Sans", "Source Sans 3", "Roboto", "DejaVu Sans"]
rcParams["axes.unicode_minus"] = False

# ---------- palette (sampled from PT) ----------
BG          = "#FEFEFE"
TITLE       = "#0B1942"   # dark navy (not pure black)
SUBTITLE    = "#5B6069"
CURVE_BLUE  = "#0C386C"   # navy curve
GREEN_TEXT  = "#2E5131"
ORANGE_TEXT = "#A35101"   # "Normal — mas subótimo" right label
RED_TEXT    = "#B01C23"   # pre-diabetes label + "5.7%" value + "in 2 years"
GREEN_FILL  = "#F0F3EB"   # zona ótima band
YELLOW_FILL = "#FEF7EB"   # normal-subótimo band
PINK_FILL   = "#FDECEC"   # pre-diabetes band
TICK        = "#5B6069"
FOOTNOTE    = "#5B6069"

# ---------- data ----------
xs = [1, 2, 3, 4]
hba1c = [4.9, 5.2, 5.4, 5.7]

# ---------- figure ----------
fig = plt.figure(figsize=(11.0, 6.2), dpi=220)
fig.patch.set_facecolor(BG)

ax = fig.add_axes([0.085, 0.22, 0.78, 0.58])
ax.set_facecolor(BG)

# ---------- background bands (correct boundaries) ----------
# green: HbA1c < 5.0
# yellow: 5.0 - 5.7
# pink: > 5.7
ax.axhspan(5.7, 6.1, color=PINK_FILL,   zorder=1)
ax.axhspan(5.0, 5.7, color=YELLOW_FILL, zorder=1)
ax.axhspan(4.5, 5.0, color=GREEN_FILL,  zorder=1)

# ---------- reference dashed line at 5.7% ("Limite de normalidade") ----------
ax.axhline(5.7, color=CURVE_BLUE, linewidth=0.9,
           linestyle=(0, (4, 3)), zorder=2, alpha=0.7)
ax.text(0.6, 5.74, "Lab cutoff for normal",
        fontsize=8.5, color=TICK, style="italic", zorder=5)

# ---------- curve: solid for points 1-3, dashed for projection 3-4 ----------
ax.plot(xs[:3], hba1c[:3],
        color=CURVE_BLUE, linewidth=2.4, zorder=4,
        marker="o", markersize=10,
        markerfacecolor=CURVE_BLUE,           # filled markers for real measurements
        markeredgecolor=CURVE_BLUE, markeredgewidth=2)

# dashed projection line from point 3 to point 4
ax.plot(xs[2:], hba1c[2:],
        color=CURVE_BLUE, linewidth=2.0, zorder=4,
        linestyle=(0, (5, 3)),
        marker="None")

# point 4 marker — HOLLOW (white-filled, navy outline) for projection
ax.plot([xs[3]], [hba1c[3]],
        marker="o", markersize=11,
        markerfacecolor="white",              # hollow
        markeredgecolor=CURVE_BLUE, markeredgewidth=2.2,
        linestyle="None", zorder=5)

# value annotations above each point
# first 3: dark text, point 4: red
for x, y in zip(xs[:3], hba1c[:3]):
    ax.text(x, y + 0.08, f"{y:.1f}%",
            ha="center", va="bottom",
            fontsize=11, color=TITLE, weight="bold", zorder=6)

ax.text(xs[3], hba1c[3] + 0.08, f"{hba1c[3]:.1f}%",
        ha="center", va="bottom",
        fontsize=11, color=RED_TEXT, weight="bold", zorder=6)

# "✓ Normal" below first 3 points
for x, y in zip(xs[:3], hba1c[:3]):
    ax.text(x, y - 0.10, "✓ Normal",
            ha="center", va="top",
            fontsize=9, color=GREEN_TEXT, zorder=6)

# "Prediabetes" below point 4 (red, italic)
ax.text(xs[3], hba1c[3] - 0.08, "Prediabetes",
        ha="center", va="top",
        fontsize=8.5, color=RED_TEXT, weight="bold", style="italic", zorder=6)

# ---------- annotation BELOW curve with curved arrow up to point 3 ----------
ax.annotate(
    "Each result was 'normal'.\nThe trend was not.",
    xy=(2.85, 5.35),                     # arrow tip near point 3
    xytext=(2.5, 4.78),                  # text position below curve
    fontsize=10, color=TITLE, ha="center", va="center",
    style="italic",
    arrowprops=dict(
        arrowstyle="->",
        color=TICK,
        linewidth=1.0,
        connectionstyle="arc3,rad=-0.35",
    ),
    zorder=7,
)

# ---------- right-side band labels ----------
ax.text(4.55, 5.85, "Prediabetes",
        fontsize=10, color=RED_TEXT, weight="bold",
        va="center", ha="left", zorder=5)
ax.text(4.55, 5.35, "Normal —\nbut suboptimal",
        fontsize=10, color=ORANGE_TEXT, weight="bold",
        va="center", ha="left", zorder=5)
ax.text(4.55, 4.75, "Optimal zone\nfor longevity",
        fontsize=10, color=GREEN_TEXT, weight="bold",
        va="center", ha="left", zorder=5)

# ---------- axes cosmetics ----------
ax.set_xlim(0.5, 4.5)
ax.set_ylim(4.5, 6.05)
ax.set_xticks(xs)

# X-axis labels — multi-line, with "in 2 years" of point 4 in red
xtick_main = ["Check-up 1", "Check-up 2", "Check-up 3", "Check-up 4"]
xtick_sub  = ["(5 years ago)", "(2 years ago)", "(today)", "(projection"]
ax.set_xticklabels([f"{m}\n{s}" for m, s in zip(xtick_main, xtick_sub)])
# manually place red second-line for point 4 below the standard tick label
ax.text(xs[3], 4.40, "in 2 years)",
        ha="center", va="top",
        fontsize=9, color=RED_TEXT,
        transform=ax.transData, clip_on=False, zorder=8)

ax.set_yticks([4.5, 5.0, 5.2, 5.5, 5.7, 6.0])
ax.set_yticklabels(["4.5%", "5.0%", "5.2%", "5.5%", "5.7%", "6.0%"])
ax.tick_params(axis="x", colors=TICK, labelsize=9, length=0, pad=8)
ax.tick_params(axis="y", colors=TICK, labelsize=9, length=0, pad=4)
for spine in ("top", "right"):
    ax.spines[spine].set_visible(False)
ax.spines["left"].set_color("#D6D8DC")
ax.spines["bottom"].set_color("#D6D8DC")

# Y-axis label
fig.text(0.04, 0.49, "HbA1c (%)",
         fontsize=10, color=TITLE, weight="bold",
         rotation=90, va="center", ha="center")

# ---------- title (NO badge — PT version has none) ----------
fig.text(0.04, 0.910,
         "Three 'normal' check-ups — a dangerous trend",
         fontsize=20, color=TITLE, weight="bold")

# ---------- source footer (3 lines) ----------
source_lines = [
    "HbA1c (glycated hemoglobin) reflects average blood sugar over the past 2–3 months.",
    "The lab 'normal' range (below 5.7%) includes values that, while not meeting the prediabetes threshold,",
    "are far from optimal for longevity. The trend matters more than any single number.",
]
for i, line in enumerate(source_lines):
    fig.text(0.04, 0.080 - i * 0.020, line,
             fontsize=8, color=FOOTNOTE)

# save
out_path = Path(__file__).resolve().parents[2] / "figuras" / "en" / "Cap01 Fig01.PNG"
out_path.parent.mkdir(parents=True, exist_ok=True)
plt.savefig(out_path, dpi=220, facecolor=BG)
print(f"saved → {out_path}")
