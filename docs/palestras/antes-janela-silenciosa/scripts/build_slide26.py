#!/usr/bin/env python3
"""Build slide 26 — TOTG de André (faithful reconstruction of book Fig 6.4).

Data from ebook Cap 6: Glicose 92/148/162/154/131 and Insulina 8.5/78/124/118/89
at [0, 30, 60, 90, 120] min. Dual Y-axes, normal post-prandial glucose band,
annotations on insulin peak (Kraft II, 14x basal).
"""
from PIL import Image, ImageDraw, ImageFont
from pathlib import Path
import sys
sys.path.insert(0, "/home/user/plenya/docs/palestras/antes-janela-silenciosa/scripts")
from composite_logo import composite_logo

OUT = Path("/home/user/plenya/docs/palestras/antes-janela-silenciosa/slides/slide-26-totg-andre-base.png")
FINAL = Path("/home/user/plenya/docs/palestras/antes-janela-silenciosa/slides/slide-26-totg-andre-final.png")

W, H = 2048, 1152
PETROL = (6, 59, 79)
CREAM = (234, 231, 218)
CREAM_DIM = (234, 231, 218, 150)
GOLD = (179, 134, 69)
SAGE = (146, 184, 180)
OCEAN = (65, 126, 142)
AMBER_SOFT = (200, 130, 60, 60)

canvas = Image.new("RGB", (W, H), PETROL)
draw = ImageDraw.Draw(canvas, "RGBA")

serif = "/usr/share/fonts/truetype/liberation/LiberationSerif-Regular.ttf"
serif_bold = "/usr/share/fonts/truetype/liberation/LiberationSerif-Bold.ttf"
serif_italic = "/usr/share/fonts/truetype/liberation/LiberationSerif-Italic.ttf"
sans = "/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf"
sans_bold = "/usr/share/fonts/truetype/dejavu/DejaVuSans-Bold.ttf"

f_eyebrow = ImageFont.truetype(sans, 22)
f_title = ImageFont.truetype(serif_bold, 44)
f_subtitle = ImageFont.truetype(serif_italic, 28)
f_axis = ImageFont.truetype(sans, 20)
f_value = ImageFont.truetype(sans_bold, 22)
f_legend = ImageFont.truetype(sans, 22)
f_note = ImageFont.truetype(serif_italic, 22)
f_source = ImageFont.truetype(sans, 16)

# ===== Header =====
draw.text((150, 80), "FIGURA 6.4  ·  TOTG DE ANDRÉ", font=f_eyebrow, fill=GOLD)
draw.text((150, 115), "Quando o jejum mente: glicose aparentemente normal", font=f_title, fill=CREAM)
draw.text((150, 175), "com resposta insulínica 14× o basal.", font=f_subtitle, fill=CREAM)

# ===== Chart area =====
CX0, CY0 = 190, 290        # top-left of plot
CX1, CY1 = 1720, 900       # bottom-right of plot
PW, PH = CX1 - CX0, CY1 - CY0

# Y-axis scales
# Left: Glicose 0-180 mg/dL
# Right: Insulina 0-140 µIU/mL
GLUC_MAX = 180
INS_MAX = 140

# X-axis: time points
times = [0, 30, 60, 90, 120]
x_labels = ["jejum", "30 min", "60 min", "90 min", "120 min"]

def x(t):
    return CX0 + (t / 120) * PW

def y_gluc(v):
    return CY1 - (v / GLUC_MAX) * PH

def y_ins(v):
    return CY1 - (v / INS_MAX) * PH

# ===== Normal post-prandial glucose band (<140 mg/dL) =====
y_140 = y_gluc(140)
draw.rectangle([CX0, y_140, CX1, CY1], fill=(146, 184, 180, 28))
draw.text((CX0 + 20, y_140 + 10),
          "Faixa normal pós-prandial — glicose < 140 mg/dL",
          font=f_note, fill=(146, 184, 180, 180))

# ===== Grid =====
for v in [30, 60, 90, 120, 150, 180]:
    yg = y_gluc(v)
    draw.line([(CX0, yg), (CX1, yg)], fill=(234, 231, 218, 20), width=1)
for t in times:
    xt = x(t)
    draw.line([(xt, CY0), (xt, CY1)], fill=(234, 231, 218, 20), width=1)

# ===== Axes =====
# Left Y axis label
draw.text((70, CY0 - 40), "Glicose\n(mg/dL)", font=f_axis, fill=CREAM)
for v in [0, 30, 60, 90, 120, 150, 180]:
    yg = y_gluc(v)
    draw.text((130, yg - 10), str(v), font=f_axis, fill=CREAM_DIM, anchor="la")
    draw.line([(175, yg), (CX0, yg)], fill=CREAM_DIM, width=1)

# Right Y axis
draw.text((CX1 + 30, CY0 - 40), "Insulina\n(µIU/mL)", font=f_axis, fill=GOLD)
for v in [0, 20, 40, 60, 80, 100, 120, 140]:
    yi = y_ins(v)
    draw.text((CX1 + 15, yi - 10), str(v), font=f_axis, fill=(179, 134, 69, 180), anchor="la")
    draw.line([(CX1, yi), (CX1 + 10, yi)], fill=(179, 134, 69, 180), width=1)

# X axis baseline
draw.line([(CX0, CY1), (CX1, CY1)], fill=CREAM, width=2)
# X tick labels
for t, label in zip(times, x_labels):
    xt = x(t)
    draw.text((xt, CY1 + 20), label, font=f_axis, fill=CREAM, anchor="ma")

# ===== Data series =====
gluc_pts = [92, 148, 162, 154, 131]
ins_pts = [8.5, 78, 124, 118, 89]

# Glicose line (cream solid)
gpoints = [(x(t), y_gluc(v)) for t, v in zip(times, gluc_pts)]
for i in range(len(gpoints) - 1):
    draw.line([gpoints[i], gpoints[i + 1]], fill=CREAM, width=3)

# Insulina line (gold dashed feel — solid thicker with circles)
ipoints = [(x(t), y_ins(v)) for t, v in zip(times, ins_pts)]
for i in range(len(ipoints) - 1):
    draw.line([ipoints[i], ipoints[i + 1]], fill=GOLD, width=4)

# Data points (glicose)
for (px, py), v in zip(gpoints, gluc_pts):
    draw.ellipse([px - 8, py - 8, px + 8, py + 8], fill=CREAM, outline=PETROL, width=2)
    # Value label above (small)
    draw.text((px, py - 28), str(v), font=f_value, fill=CREAM, anchor="ma")

# Data points (insulina)
for (px, py), v in zip(ipoints, ins_pts):
    draw.ellipse([px - 9, py - 9, px + 9, py + 9], fill=GOLD, outline=PETROL, width=2)
    val_s = str(v).replace(".0", "").replace(".", ",")
    draw.text((px + 18, py - 8), val_s, font=f_value, fill=GOLD)

# ===== Highlight peak insulin at 60 min =====
peak_x, peak_y = ipoints[2]
# Halo around peak
for r in range(28, 16, -3):
    draw.ellipse([peak_x - r, peak_y - r, peak_x + r, peak_y + r],
                 outline=(179, 134, 69, 40), width=2)

# Annotation box to the right of the peak
ann_x = peak_x + 60
ann_y = peak_y - 70
draw.text((ann_x, ann_y),
          "pico 124 µIU/mL", font=ImageFont.truetype(sans_bold, 26), fill=GOLD)
draw.text((ann_x, ann_y + 36),
          "14× o valor basal", font=ImageFont.truetype(sans, 22), fill=CREAM)
draw.text((ann_x, ann_y + 66),
          "Kraft II · hiperinsulinemia compensatória", font=f_note, fill=(234, 231, 218, 200))

# Line from annotation to peak
draw.line([(peak_x + 12, peak_y - 10), (ann_x - 10, ann_y + 18)],
          fill=(179, 134, 69, 140), width=1)

# ===== Annotate basal insulin "normal lab" =====
basal_x, basal_y = ipoints[0]
draw.text((basal_x - 90, basal_y + 18),
          "8,5 — 'normal'\npelo laboratório",
          font=ImageFont.truetype(sans, 18), fill=(234, 231, 218, 170))

# ===== Legend (top-right corner of chart area) =====
legend_x = CX1 - 340
legend_y = CY0 - 0
# Glicose sample
draw.line([(legend_x, legend_y + 12), (legend_x + 40, legend_y + 12)], fill=CREAM, width=3)
draw.ellipse([legend_x + 16, legend_y + 6, legend_x + 26, legend_y + 16], fill=CREAM)
draw.text((legend_x + 55, legend_y), "Glicose (mg/dL)", font=f_legend, fill=CREAM)

# Insulina sample
draw.line([(legend_x, legend_y + 48), (legend_x + 40, legend_y + 48)], fill=GOLD, width=4)
draw.ellipse([legend_x + 15, legend_y + 41, legend_x + 27, legend_y + 53], fill=GOLD)
draw.text((legend_x + 55, legend_y + 36), "Insulina (µIU/mL)", font=f_legend, fill=GOLD)

# ===== Source footer =====
draw.text((150, 1060),
          "Padrão de referência: Kraft JR (Detection of Diabetes Mellitus In Situ, 2008). Dados reconstruídos a partir do caso apresentado no Capítulo 6.",
          font=f_source, fill=(234, 231, 218, 140))

canvas.save(OUT, "PNG", optimize=True)
print(f"Base: {OUT.stat().st_size:,} bytes")

composite_logo(OUT, FINAL, size_pct=5, position="footer-right")
