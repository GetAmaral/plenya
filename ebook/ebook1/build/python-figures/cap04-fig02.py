"""
Cap04 Fig02 (EN) — Longevity biomarkers: normal vs. optimal ranges.
16 main + 2 complementary biomarkers, organized in 4 thematic groups.

Mirrors PT version. Dense table figure in portrait orientation.
Colors sampled from PT original.
"""

from pathlib import Path
import textwrap
import matplotlib.pyplot as plt
from matplotlib import rcParams
from matplotlib.patches import Rectangle

# DejaVu Sans first because it has ♂/♀ glyphs that Inter is missing
rcParams["font.family"] = "sans-serif"
rcParams["font.sans-serif"] = ["DejaVu Sans", "Inter", "Open Sans", "Roboto"]
rcParams["axes.unicode_minus"] = False

# ---------- palette (sampled from PT) ----------
BG          = "#FCFCFA"   # very slight cream background
TITLE       = "#0B2C3D"   # very dark navy
SUBTITLE    = "#5B6069"
BADGE_RED   = "#C81D1B"
GREEN_FILL  = "#EDF2EE"   # FAIXA ÓTIMA column
PINK_FILL   = "#FAF1EC"   # FAIXA NORMAL column
SECTION_BG  = "#F4EFE5"   # section header background (light tan)
SECTION_TXT = "#0B2C3D"
COL_HDR     = "#5B6069"
ROW_TEXT    = "#1F2A36"
DIVIDER     = "#E0DCD2"
FOOTNOTE    = "#5B6069"

# ---------- data (18 biomarkers in 4 sections) ----------
SECTIONS = [
    ("1. METABOLIC", [
        ("Fasting insulin",
         "< 25 µIU/mL",
         "< 8 µIU/mL",
         "Early insulin resistance — rises 5–10 years before glucose."),
        ("HbA1c",
         "≤ 6.5%",
         "≤ 5.4% (ideal 4.8–5.2%)",
         "Long-term metabolic health — 2–3 month average."),
        ("TG/HDL ratio",
         "no clinical reference",
         "< 2.0 (> 3.5 = risk)",
         "Accessible proxy for insulin resistance."),
    ]),
    ("2. INFLAMMATION AND STRESS", [
        ("hs-CRP",
         "< 3.0 mg/L",
         "< 1.0 mg/L",
         "Low-grade systemic inflammation (inflammaging)."),
        ("Uric acid",
         "3.5–7.2 mg/dL (♂)",
         "< 5.5 mg/dL",
         "Metabolic stress — inhibits endothelial NO."),
        ("GGT",
         "< 60 U/L (♂) / < 40 (♀)",
         "< 25–30 U/L",
         "Oxidative stress, early hepatic stress."),
    ]),
    ("3. LIPIDS AND CARDIOVASCULAR", [
        ("ApoB",
         "< 130 mg/dL",
         "< 90 mg/dL",
         "Real atherosclerotic risk — particle count."),
        ("ApoB/ApoA1 ratio",
         "varies",
         "< 0.6",
         "Balance between arterial attack and defense."),
        ("Lp(a)",
         "< 30 mg/dL (< 75 nmol/L)",
         "< 30 mg/dL (< 75 nmol/L)",
         "Genetic atherogenic risk — measure once per lifetime."),
        ("Non-HDL cholesterol",
         "no single cutoff",
         "< 100 mg/dL",
         "Atherogenic lipid fraction in one number."),
        ("High-sensitivity troponin",
         "< 16 ng/L (♂) / < 24 (♀)",
         "< 14 ng/L",
         "Early subclinical myocardial injury."),
        ("NT-proBNP",
         "age-dependent",
         "< 50 pg/mL (< 50 yrs)",
         "Early cardiac wall stress."),
    ]),
    ("4. REGULATION AND LONGEVITY", [
        ("Vitamin D (25-OH)",
         "> 20 ng/mL",
         "40–60 ng/mL",
         "Optimal immune, bone, and muscle function."),
        ("Homocysteine",
         "< 15 µmol/L",
         "< 10 µmol/L",
         "Vascular and cognitive risk — reversible with B6/B9/B12."),
        ("Ferritin",
         "12–300 ng/mL",
         "40–150 ng/mL",
         "Iron reserve / chronic inflammation."),
        ("eGFR (cystatin C)",
         "> 60 mL/min",
         "> 90 mL/min",
         "Real kidney function — independent of muscle mass."),
        ("Microalbuminuria",
         "< 30 mg/g",
         "< 10 mg/g",
         "Earliest glomerular injury."),
        ("Serum albumin",
         "3.5–5.2 g/dL",
         "> 4.5 g/dL",
         "Hepatic synthesis."),
    ]),
]

COL_HEADERS = ["BIOMARKER", "'NORMAL' RANGE (LAB)", "OPTIMAL RANGE (LONGEVITY)", "WHAT IT REVEALS"]

# ---------- layout (figure coordinates, top-down) ----------
FIG_W, FIG_H = 9.0, 14.5

# Column boundaries (x positions in fig coords): 4 columns
COL_X = [0.04, 0.27, 0.47, 0.69, 0.97]   # left edges + final right edge
COL_WIDTHS = [COL_X[i+1] - COL_X[i] for i in range(4)]

ROW_H        = 0.034   # height of each data row (taller for wrapping)
HDR_ROW_H    = 0.020   # column header row
SECTION_H    = 0.028   # section title bar
GAP_AFTER_SEC = 0.005

# Max characters per line for text wrapping in the "What it reveals" column
WRAP_REVEALS = 38

# starting y position (after title block) — pushed down from top
y = 0.875

# ---------- figure ----------
fig = plt.figure(figsize=(FIG_W, FIG_H), dpi=180)
fig.patch.set_facecolor(BG)
ax = fig.add_axes([0, 0, 1, 1])
ax.set_xlim(0, 1)
ax.set_ylim(0, 1)
ax.axis("off")
ax.set_facecolor(BG)

# ---------- title block ----------
fig.text(0.04, 0.965, "FIGURE 1",
         fontsize=10, color=BADGE_RED, weight="bold")
fig.text(0.04, 0.935,
         "Longevity biomarkers: normal vs. optimal ranges.",
         fontsize=18, color=TITLE, weight="bold")
fig.text(0.04, 0.905,
         "16 main + 2 complementary biomarkers, organized in 4 thematic groups.",
         fontsize=10, color=SUBTITLE)

# ---------- helper: draw a filled rectangle in axes coords ----------
def rect(x, y_top, w, h, color):
    ax.add_patch(Rectangle((x, y_top - h), w, h,
                           facecolor=color, edgecolor="none", zorder=1))

def cell_text(x_left, y_center, w, txt, fontsize=9, color=ROW_TEXT,
              weight="normal", style="normal", ha="left", pad=0.008):
    ax.text(x_left + pad, y_center, txt,
            fontsize=fontsize, color=color, weight=weight, style=style,
            ha=ha, va="center", zorder=3)

# ---------- draw sections ----------
for sec_title, rows in SECTIONS:
    # section header bar
    rect(0.04, y, 0.93, SECTION_H, SECTION_BG)
    cell_text(0.04, y - SECTION_H/2, 0.93, sec_title,
              fontsize=11, color=SECTION_TXT, weight="bold", pad=0.012)
    y -= SECTION_H

    # column headers row
    for i, hdr in enumerate(COL_HEADERS):
        cell_text(COL_X[i], y - HDR_ROW_H/2, COL_WIDTHS[i], hdr,
                  fontsize=7.8, color=COL_HDR, weight="bold")
    y -= HDR_ROW_H
    # thin divider under column headers
    ax.plot([0.04, 0.97], [y, y], color=DIVIDER, linewidth=0.6, zorder=2)

    # data rows
    for biomarker, normal, optimal, reveals in rows:
        y_top = y
        # column 2 pink fill (NORMAL range)
        rect(COL_X[1], y_top, COL_WIDTHS[1], ROW_H, PINK_FILL)
        # column 3 green fill (OPTIMAL range)
        rect(COL_X[2], y_top, COL_WIDTHS[2], ROW_H, GREEN_FILL)
        # text
        cell_text(COL_X[0], y_top - ROW_H/2, COL_WIDTHS[0], biomarker,
                  fontsize=9.2, color=ROW_TEXT, weight="bold")
        cell_text(COL_X[1], y_top - ROW_H/2, COL_WIDTHS[1], normal,
                  fontsize=8.5, color=ROW_TEXT)
        cell_text(COL_X[2], y_top - ROW_H/2, COL_WIDTHS[2], optimal,
                  fontsize=8.5, color=ROW_TEXT, weight="bold")
        # wrap the reveals column to avoid overflow
        wrapped_reveals = "\n".join(textwrap.wrap(reveals, WRAP_REVEALS)) or reveals
        cell_text(COL_X[3], y_top - ROW_H/2, COL_WIDTHS[3], wrapped_reveals,
                  fontsize=8.2, color="#3A4250")
        y -= ROW_H
        # row divider
        ax.plot([0.04, 0.97], [y, y], color=DIVIDER, linewidth=0.4, zorder=2)

    y -= GAP_AFTER_SEC

# ---------- footer ----------
y_footer = max(y - 0.025, 0.04)
fig.text(0.04, y_footer + 0.012,
         "♂ men · ♀ women",
         fontsize=8, color=FOOTNOTE, style="italic")
fig.text(0.04, y_footer - 0.005,
         "Reference values adapted from longevity-medicine clinical practice and guidelines (ESC/EAS, ACC/AHA, ABMFI).",
         fontsize=7.5, color=FOOTNOTE)
fig.text(0.04, y_footer - 0.022,
         "On certain biomarkers the laboratory's 'normal' coincides with the threshold of injury; the optimal target is well below that limit.",
         fontsize=7.5, color=FOOTNOTE)
fig.text(0.04, y_footer - 0.039,
         "Chapter 4 — Plenya Expanded Panel — Normal vs. Optimal.",
         fontsize=7.5, color=FOOTNOTE, style="italic")

# ---------- save ----------
out_path = Path(__file__).resolve().parents[2] / "figuras" / "en" / "Cap04 Fig02.PNG"
out_path.parent.mkdir(parents=True, exist_ok=True)
plt.savefig(out_path, dpi=200, facecolor=BG)
print(f"saved → {out_path}")
