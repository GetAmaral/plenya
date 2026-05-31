"""
Cap10 Fig01 (EN) — The testosterone of a 48-year-old — and a man of 80.
Population mean total testosterone (ng/dL) across adult life.

Source: MMAS (1988-1994), European Male Ageing Study,
Travison et al., JCEM, 2007; Wu et al., PLoS ONE, 2010.
"""

from pathlib import Path
import matplotlib.pyplot as plt
from matplotlib import rcParams

rcParams["font.family"] = "sans-serif"
rcParams["font.sans-serif"] = ["Inter", "Open Sans", "Source Sans 3", "Roboto", "DejaVu Sans"]
rcParams["axes.unicode_minus"] = False

# ---------- palette (sampled from PT original) ----------
BG          = "#FDFDFD"   # near-white
TITLE       = "#071122"   # very dark navy
SUBTITLE    = "#5B6069"   # cool dark gray
BADGE       = "#5B6069"   # FIGURE 1 badge
CURVE_GRAY  = "#8D9298"   # cool gray for curve + markers
GREEN_FILL  = "#F3F5EE"   # zona ótima fill
GREEN_TEXT  = "#526C3A"   # zona ótima label
PINK_FILL   = "#FAEFEC"   # hipogonadismo fill
RED_ACCENT  = "#C81D1B"   # red highlight (Paulo + +30 anos arrow)
TICK        = "#5B6069"
FOOTNOTE    = "#5B6069"

# ---------- data: population mean total T by age ----------
ages = [20, 30, 40, 50, 60, 70, 80, 85]
test = [680, 610, 550, 470, 410, 350, 310, 310]

PAULO_AGE = 48
PAULO_T   = 310

# ---------- figure ----------
fig = plt.figure(figsize=(10.0, 6.0), dpi=220)
fig.patch.set_facecolor(BG)

ax = fig.add_axes([0.085, 0.20, 0.86, 0.58])
ax.set_facecolor(BG)

# ---------- background bands ----------
ax.axhspan(500, 760, color=GREEN_FILL, zorder=1)
ax.axhspan(0,   300, color=PINK_FILL,  zorder=1)

# ---------- horizontal grid (subtle) ----------
for y in (100, 200, 300, 400, 500, 600, 700):
    ax.axhline(y, color="#E5E6E8", linewidth=0.7, zorder=2)

# ---------- main population curve ----------
ax.plot(ages, test,
        color=CURVE_GRAY, linewidth=2.0, zorder=4,
        marker="o", markersize=8,
        markerfacecolor="white",
        markeredgecolor=CURVE_GRAY, markeredgewidth=2)

# value labels above each point
for x, y in zip(ages, test):
    ax.annotate(f"{y}", xy=(x, y), xytext=(0, 12),
                textcoords="offset points",
                ha="center", fontsize=9.5,
                color=TITLE, weight="bold", zorder=6)

# ---------- Paulo highlight (red dot + dashed projection) ----------
# red dot at (48, 310)
ax.plot([PAULO_AGE], [PAULO_T],
        marker="o", markersize=11,
        markerfacecolor=RED_ACCENT, markeredgecolor=RED_ACCENT,
        zorder=7, linestyle="None")

# horizontal red dashed line from (48, 310) to (80, 310) showing equivalence
ax.annotate("", xy=(80, PAULO_T), xytext=(PAULO_AGE + 1, PAULO_T),
            arrowprops=dict(arrowstyle="->", color=RED_ACCENT,
                            linewidth=1.4,
                            linestyle=(0, (5, 3))),
            zorder=6)

# Paulo callout box (thin red border, white fill)
from matplotlib.patches import FancyBboxPatch
ax.text(54, 215, "Paulo, 48 anos: 310 ng/dL",
        fontsize=9.5, color=TITLE, weight="bold",
        ha="left", va="top", zorder=9,
        bbox=dict(boxstyle="round,pad=0.45",
                  facecolor="white", edgecolor=RED_ACCENT,
                  linewidth=1.0))
ax.text(54, 178, "Typical value at age 80",
        fontsize=8.5, color=TITLE, ha="left", va="top",
        style="italic", zorder=9)

# "+30 years of hormonal aging" annotation along the red arrow
ax.text(64, 285, "+30 years of hormonal aging",
        fontsize=9, color=RED_ACCENT, ha="center", va="top",
        style="italic", weight="bold", zorder=8)

# ---------- band labels ----------
ax.text(20.5, 720,
        "OPTIMAL ZONE\n>500 ng/dL",
        fontsize=9.5, color=GREEN_TEXT, weight="bold",
        va="top", ha="left", zorder=5)
ax.text(20.5, 690,
        "Associated with lower mortality in longevity studies.",
        fontsize=8.3, color=GREEN_TEXT, va="top", ha="left",
        style="italic", zorder=5)

ax.text(20.5, 280,
        "LABORATORY HYPOGONADISM\n<300 ng/dL",
        fontsize=9.5, color=RED_ACCENT, weight="bold",
        va="top", ha="left", zorder=5)

# ---------- axes cosmetics ----------
ax.set_xlim(18, 87)
ax.set_ylim(0, 760)
ax.set_xticks([20, 30, 40, 48, 50, 60, 70, 80, 85])
ax.set_yticks([0, 100, 200, 300, 400, 500, 600, 700])
ax.tick_params(axis="x", colors=TICK, labelsize=9, length=0, pad=4)
ax.tick_params(axis="y", colors=TICK, labelsize=9, length=0, pad=4)
ax.set_xlabel("AGE (YEARS)", fontsize=9, color=TICK,
              labelpad=10, weight="bold")
for spine in ("top", "right"):
    ax.spines[spine].set_visible(False)
ax.spines["left"].set_color("#D6D8DC")
ax.spines["bottom"].set_color("#D6D8DC")

# Y-axis label (vertical, small caps style)
fig.text(0.085, 0.795, "TOTAL TESTOSTERONE (ng/dL)",
         fontsize=9, color=TICK, weight="bold", ha="left")

# ---------- title block ----------
fig.text(0.04, 0.945, "FIGURE 1",
         fontsize=9, color=BADGE, weight="bold")
fig.text(0.04, 0.900,
         "The testosterone of a 48-year-old — and a man of 80.",
         fontsize=15.5, color=TITLE, weight="bold")
fig.text(0.04, 0.860, "The same number.",
         fontsize=15.5, color=TITLE, weight="bold")
fig.text(0.04, 0.825,
         "Population mean total testosterone (ng/dL) across adult life.",
         fontsize=10, color=SUBTITLE)

# ---------- source footer ----------
source_lines = [
    "Mean total testosterone by age range in adult men (combined population cohort data) — 1–2% annual decline after age 40. The green band marks values",
    "associated with lower mortality in longevity studies; the lower band marks formal laboratory hypogonadism.",
    "Sources: MMAS (1988–1994); European Male Ageing Study; Travison et al., JCEM, 2007; Wu et al., PLoS ONE, 2010.",
]
for i, line in enumerate(source_lines):
    fig.text(0.04, 0.085 - i * 0.020, line,
             fontsize=7.4, color=FOOTNOTE)

# save
out_path = Path(__file__).resolve().parents[2] / "figuras" / "en" / "Cap10 Fig01.PNG"
out_path.parent.mkdir(parents=True, exist_ok=True)
plt.savefig(out_path, dpi=220, facecolor=BG)
print(f"saved → {out_path}")
