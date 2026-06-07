#!/usr/bin/env python3
"""render.py <figure_dir>

Mechanical renderer. Reads <figure_dir>/source/ground_truth.json and emits:
  <figure_dir>/figure.svg

No content is invented here — every text and every position is read from
ground_truth.json, which itself is OCR + visually-verified.
"""
import json
import sys
import html
from pathlib import Path


# k = font-size multiplier of bbox height. ~1.0 for caps (bbox = cap-height);
# ~1.15 for mixed-case (descenders push bbox down past baseline).
SERIF = "'Cormorant Garamond', 'DejaVu Serif', Georgia, serif"
SANS  = "'Inter', 'Liberation Sans', 'DejaVu Sans', sans-serif"

ROLE_STYLES = {
    "title_serif_bold":         {"k": 1.20, "ff": SERIF, "fw": "900",    "fill": "#000"},
    "subtitle_gray":            {"k": 0.95, "ff": SANS,  "fw": "normal", "fill": "#555"},

    "instr_name":               {"k": 1.05, "ff": SANS,  "fw": "900",    "fill": "#000"},
    "instr_diagnosis":          {"k": 1.00, "ff": SANS,  "fw": "bold",   "fill": "#000"},
    "instr_desc":               {"k": 1.00, "ff": SANS,  "fw": "normal", "fill": "#555"},

    "axis_tick":                {"k": 0.95, "ff": SANS,  "fw": "normal", "fill": "#000", "anchor": "middle"},
    "axis_tick_cut":            {"k": 1.00, "ff": SANS,  "fw": "900",    "fill": "#000", "anchor": "middle"},
    "axis_category":            {"k": 0.95, "ff": SANS,  "fw": "normal", "fill": "#666", "fs": "italic"},

    "cut_label":                {"k": 0.85, "ff": SANS,  "fw": "bold",   "fill": "#000"},
    "cut_note":                 {"k": 0.95, "ff": SANS,  "fw": "normal", "fill": "#666", "fs": "italic"},

    "alert_box_text":           {"k": 0.90, "ff": SANS,  "fw": "normal", "fill": "#fff"},

    "sidebar_title_serif_bold": {"k": 1.10, "ff": SERIF, "fw": "bold",   "fill": "#000"},
    "sidebar_item":             {"k": 1.00, "ff": SANS,  "fw": "normal", "fill": "#000"},
    "sidebar_callout":          {"k": 1.10, "ff": SANS,  "fw": "bold",   "fill": "#000", "anchor": "middle"},

    "footer":                   {"k": 0.95, "ff": SANS,  "fw": "normal", "fill": "#666", "fs": "italic"},

    "col_header":               {"k": 1.00, "ff": SANS,  "fw": "900",    "fill": "#666"},
    "group_label":              {"k": 1.05, "ff": SANS,  "fw": "900",    "fill": "#fff"},
    "group_label_plain":        {"k": 0.80, "ff": SANS,  "fw": "900",    "fill": "#555"},
    "small_label_gray":         {"k": 1.00, "ff": SANS,  "fw": "900",    "fill": "#555"},
    "bar_label":                {"k": 1.00, "ff": SANS,  "fw": "900",    "fill": "#000"},
    "bar_label_sub":            {"k": 0.95, "ff": SANS,  "fw": "normal", "fill": "#555"},
    "bar_value":                {"k": 1.00, "ff": SANS,  "fw": "900",    "fill": "#000"},
    "bar_value_big":            {"k": 1.00, "ff": SANS,  "fw": "900",    "fill": "#000"},
    "bar_value_sm":             {"k": 0.95, "ff": SANS,  "fw": "900",    "fill": "#000"},
    "bar_value_soc":            {"k": 1.00, "ff": SANS,  "fw": "900",    "fill": "#000"},
    "bar_value_cla":            {"k": 1.00, "ff": SANS,  "fw": "900",    "fill": "#000"},
    "axis_title":               {"k": 0.85, "ff": SANS,  "fw": "normal", "fill": "#555", "fs": "italic"},
    "side_annotation":          {"k": 0.90, "ff": SANS,  "fw": "bold",   "fill": "#000"},
    "side_annotation_gray":     {"k": 0.90, "ff": SANS,  "fw": "bold",   "fill": "#555"},
    "alert_red_title":          {"k": 0.70, "ff": SANS,  "fw": "900",    "fill": "#000"},
    "alert_red_body":           {"k": 0.70, "ff": SANS,  "fw": "normal", "fill": "#000"},
    "alert_inline_title":       {"k": 0.95, "ff": SANS,  "fw": "900",    "fill": "#000"},
    "alert_inline_text":        {"k": 0.95, "ff": SANS,  "fw": "normal", "fill": "#000"},
    "right_callout":            {"k": 0.95, "ff": SANS,  "fw": "normal", "fill": "#555", "fs": "italic"},

    "time_header":              {"k": 1.05, "ff": SANS,  "fw": "900",    "fill": "#1a1a1a", "anchor": "middle"},
    "time_pill":                {"k": 0.95, "ff": SANS,  "fw": "600",    "fill": "#1a1a1a", "anchor": "middle"},
    "group_block_label":        {"k": 1.05, "ff": SANS,  "fw": "900",    "fill": "#555"},
    "group_block_sub":          {"k": 0.85, "ff": SANS,  "fw": "normal", "fill": "#888", "fs": "italic"},
    "row_label":                {"k": 1.00, "ff": SANS,  "fw": "normal", "fill": "#000"},
    "row_label_sub":            {"k": 0.90, "ff": SANS,  "fw": "normal", "fill": "#555"},
    "cell_text":                {"k": 1.00, "ff": SANS,  "fw": "normal", "fill": "#000"},
    "cell_value_big":           {"k": 1.10, "ff": SANS,  "fw": "900",    "fill": "#000"},
    "quote_serif_italic":       {"k": 1.10, "ff": SERIF, "fw": "500",    "fill": "#000", "fs": "italic", "anchor": "middle"},
    "quote_attrib":             {"k": 1.00, "ff": SANS,  "fw": "normal", "fill": "#555", "anchor": "end"},
    "footer_small":             {"k": 0.95, "ff": SANS,  "fw": "normal", "fill": "#666"},
    "venn_cardinal":            {"k": 1.05, "ff": SANS,  "fw": "700",    "fill": "#1a1a1a", "anchor": "middle"},
    "venn_intersection":        {"k": 1.00, "ff": SANS,  "fw": "900",    "fill": "#000", "anchor": "middle"},
}


def esc(s):
    return html.escape(s, quote=True)


def render_text_bbox(el, style):
    x1, y1, x2, y2 = el["bbox"]
    w, h = x2 - x1, y2 - y1
    k = el.get("k", style.get("k", 1.0))
    fs = round(h * k)
    anchor = el.get("anchor", style.get("anchor", "start"))
    tx = x1
    if anchor == "middle":
        tx = (x1 + x2) / 2
    elif anchor == "end":
        tx = x2
    # baseline: place caps within bbox. fs > h because of descender room,
    # so top of caps sits a bit above y1; baseline ≈ y1 + cap_height where
    # cap_height ≈ 0.72 * fs for our serifs/sans.
    ty = y1 + 0.78 * h
    ff = el.get("ff", style["ff"])
    fw = el.get("fw", style["fw"])
    fill = el.get("fill", style["fill"])
    font_style = el.get("fs", style.get("fs"))
    attrs = [
        f'x="{tx:.0f}"', f'y="{ty:.0f}"',
        f'font-size="{fs}"',
        f'font-family="{ff}"',
        f'font-weight="{fw}"',
        f'fill="{fill}"',
        f'text-anchor="{anchor}"',
    ]
    if font_style:
        attrs.append(f'font-style="{font_style}"')
    if el.get("fit_width"):
        attrs.append(f'textLength="{w}"')
        attrs.append('lengthAdjust="spacingAndGlyphs"')
    stroke = el.get("stroke")
    if stroke:
        attrs.append(f'stroke="{stroke}"')
        attrs.append(f'stroke-width="{el.get("stroke_width", 1)}"')
        attrs.append('paint-order="stroke fill"')
        attrs.append('stroke-linejoin="round"')
    return f'<text {" ".join(attrs)}>{esc(el["text"])}</text>'


def render_badge_circle(el, fill_bg="#000", fill_text="#fff"):
    cx, cy, r = el["cx"], el["cy"], el["r"]
    out = [f'<circle cx="{cx}" cy="{cy}" r="{r}" fill="{fill_bg}"/>']
    if "text" in el:
        fs = round(r * 1.2)
        out.append(
            f'<text x="{cx}" y="{cy + r*0.42:.0f}" font-size="{fs}" '
            f'font-family="Inter, sans-serif" font-weight="bold" '
            f'fill="{fill_text}" text-anchor="middle">{esc(el["text"])}</text>'
        )
    return "\n".join(out)


def render_badge_dark(el):
    """Rectangular dark badge with centered light text. Pill style with rounded corners + padding."""
    x1, y1, x2, y2 = el["bbox"]
    w, h = x2 - x1, y2 - y1
    fs = round(h * el.get("k", 0.45))
    fw = el.get("fw", "700")
    rx = el.get("rx", 10)
    return (
        f'<rect x="{x1}" y="{y1}" width="{w}" height="{h}" rx="{rx}" ry="{rx}" fill="#000"/>\n'
        f'<text x="{(x1+x2)/2:.0f}" y="{y1+h*0.68:.0f}" font-size="{fs}" '
        f'font-family="Inter, sans-serif" font-weight="{fw}" letter-spacing="0.5" '
        f'fill="#fff" text-anchor="middle">{esc(el["text"])}</text>'
    )


def render_shape(s):
    t = s["type"]
    if t == "rect_fill":
        x1, y1, x2, y2 = s["bbox"]
        color = s.get("color", "#000")
        rx = s.get("rx", 0)
        rx_attr = f' rx="{rx}" ry="{rx}"' if rx else ""
        return f'<rect x="{x1}" y="{y1}" width="{x2-x1}" height="{y2-y1}" fill="{color}"{rx_attr}/>'
    if t == "rect_stroke":
        x1, y1, x2, y2 = s["bbox"]
        sw = s.get("stroke", 2)
        color = s.get("color", "#000")
        rx = s.get("rx", 0)
        rx_attr = f' rx="{rx}" ry="{rx}"' if rx else ""
        return f'<rect x="{x1}" y="{y1}" width="{x2-x1}" height="{y2-y1}" fill="none" stroke="{color}" stroke-width="{sw}"{rx_attr}/>'
    if t == "line_h":
        sw = s.get("stroke", 2)
        color = s.get("color", "#000")
        if color == "gray": color = "#888"
        if color == "lightgray": color = "#ddd"
        return f'<line x1="{s["x1"]}" y1="{s["y"]}" x2="{s["x2"]}" y2="{s["y"]}" stroke="{color}" stroke-width="{sw}"/>'
    if t == "line_v":
        sw = s.get("stroke", 2)
        color = s.get("color", "#000")
        if color == "gray": color = "#888"
        if color == "lightgray": color = "#ddd"
        return f'<line x1="{s["x"]}" y1="{s["y1"]}" x2="{s["x"]}" y2="{s["y2"]}" stroke="{color}" stroke-width="{sw}"/>'
    if t == "circle_fill":
        color = s.get("color", "#000")
        return f'<circle cx="{s["cx"]}" cy="{s["cy"]}" r="{s["r"]}" fill="{color}"/>'
    if t == "circle_stroke":
        color = s.get("color", "#000")
        sw = s.get("stroke", 2)
        return f'<circle cx="{s["cx"]}" cy="{s["cy"]}" r="{s["r"]}" fill="none" stroke="{color}" stroke-width="{sw}"/>'
    if t == "circle_half":
        # Left half filled, right half empty (with outline)
        cx, cy, r = s["cx"], s["cy"], s["r"]
        color = s.get("color", "#000")
        sw = s.get("stroke", 2)
        return (
            f'<path d="M {cx} {cy-r} A {r} {r} 0 0 0 {cx} {cy+r} Z" fill="{color}"/>'
            f'<circle cx="{cx}" cy="{cy}" r="{r}" fill="none" stroke="{color}" stroke-width="{sw}"/>'
        )
    if t == "downward_brace":
        # Curly brace pointing UP at center (attaches to box above),
        # arms curve out horizontally to chart edges, with small INWARD-curling tips at ends.
        # Center has a sharp "peak" going UP (like a "}" tip).
        x_l = s["x_left"]
        x_r = s["x_right"]
        y_peak = s["y_peak"]    # highest point (smallest Y) at center
        y_arms = s["y_arms"]    # arm baseline Y
        y_tick = s.get("y_tick", y_arms + 6)
        cx = s.get("center_x", (x_l + x_r) / 2)
        sw = s.get("stroke", 1.2)
        color = s.get("color", "#888")
        # Path: small inward-curl at left, horizontal arm with subtle curve up to peak, then down to right arm, end with curl
        # Inset distance for the curls
        curl = s.get("curl", 12)
        return (
            f'<path d="M {x_l} {y_tick} '
            f'Q {x_l} {y_arms} {x_l + curl} {y_arms} '                # left curl (corner)
            f'L {cx - curl} {y_arms} '                                 # horizontal arm to near center
            f'Q {cx} {y_arms} {cx} {y_peak} '                          # curve up to peak
            f'Q {cx} {y_arms} {cx + curl} {y_arms} '                   # curve back down
            f'L {x_r - curl} {y_arms} '                                # horizontal arm
            f'Q {x_r} {y_arms} {x_r} {y_tick}" '                       # right curl
            f'fill="none" stroke="{color}" stroke-width="{sw}" stroke-linecap="round"/>'
        )
    if t == "polyline":
        pts = s["points"]
        color = s.get("color", "#000")
        sw = s.get("stroke", 2)
        fill = s.get("fill", "none")
        ptstr = " ".join(f"{p[0]},{p[1]}" for p in pts)
        return f'<polyline points="{ptstr}" fill="{fill}" stroke="{color}" stroke-width="{sw}" stroke-linejoin="miter"/>'
    if t == "dashed_line_h":
        sw = s.get("stroke", 1)
        color = s.get("color", "#ccc")
        return f'<line x1="{s["x1"]}" y1="{s["y"]}" x2="{s["x2"]}" y2="{s["y"]}" stroke="{color}" stroke-width="{sw}" stroke-dasharray="4 4"/>'
    if t == "heavy_quote_open":
        # Pair of heavy round opening quotes (❝ style).
        # Each glyph: filled disk + small filled wedge protruding from bottom-left.
        cx, cy, r = s["cx"], s["cy"], s["r"]
        gap = s.get("gap", r*1.7)
        color = s.get("color", "#bbb")
        def comma(ox):
            # Bowl as full circle, then a small triangular tail attached at SW.
            tx1 = ox - r*0.55      # tail attach left
            ty1 = cy + r*0.30      # tail attach top (intersects bowl)
            tx2 = ox + r*0.05      # tail attach right (intersects bowl)
            ty2 = cy + r*0.95
            tipx = ox - r*0.45     # tail tip
            tipy = cy + r*1.45
            return (
                f'<circle cx="{ox}" cy="{cy}" r="{r}" fill="{color}"/>'
                f'<path d="M {tx1} {ty1} '
                f'Q {ox-r*0.65} {cy+r*1.15} {tipx} {tipy} '
                f'Q {ox-r*0.10} {cy+r*1.10} {tx2} {ty2} Z" fill="{color}"/>'
            )
        return comma(cx - gap/2) + comma(cx + gap/2)
    if t == "arrow":
        # Generic arrow from (x1,y1) to (x2,y2)
        import math
        x1, y1 = s["x1"], s["y1"]
        x2, y2 = s["x2"], s["y2"]
        sw = s.get("stroke", 1)
        color = s.get("color", "#000")
        head = s.get("head", 5)
        # Angle of line
        ang = math.atan2(y2 - y1, x2 - x1)
        # Arrowhead points
        hx = x2 - head * math.cos(ang)
        hy = y2 - head * math.sin(ang)
        h1x = hx - head * 0.5 * math.sin(ang)
        h1y = hy + head * 0.5 * math.cos(ang)
        h2x = hx + head * 0.5 * math.sin(ang)
        h2y = hy - head * 0.5 * math.cos(ang)
        return (
            f'<line x1="{x1}" y1="{y1}" x2="{hx:.1f}" y2="{hy:.1f}" stroke="{color}" stroke-width="{sw}"/>'
            f'<polygon points="{x2},{y2} {h1x:.1f},{h1y:.1f} {h2x:.1f},{h2y:.1f}" fill="{color}"/>'
        )
    if t == "arrow_h":
        # Horizontal arrow from (x1,y) to (x2,y), arrowhead at x2 end
        x1, x2, y = s["x1"], s["x2"], s["y"]
        sw = s.get("stroke", 1.5)
        color = s.get("color", "#888")
        if color == "gray": color = "#888"
        if color == "lightgray": color = "#bbb"
        head = s.get("head", 6)
        return (
            f'<line x1="{x1}" y1="{y}" x2="{x2-head}" y2="{y}" stroke="{color}" stroke-width="{sw}"/>'
            f'<polygon points="{x2},{y} {x2-head},{y-head/1.7:.1f} {x2-head},{y+head/1.7:.1f}" fill="{color}"/>'
        )
    if t == "glyph_centered":
        return (
            f'<text x="{s["cx"]}" y="{s["cy"]+s["size"]*0.35:.0f}" '
            f'font-size="{s["size"]}" font-family="Inter, sans-serif" '
            f'font-weight="{s.get("weight","normal")}" '
            f'fill="{s.get("fill","#000")}" text-anchor="middle">{esc(s["text"])}</text>'
        )
    if t == "bracket":
        # Square bracket [ or ]: 3 line segments
        x, y1, y2, d = s["x"], s["y1"], s["y2"], s["depth"]
        sw = s.get("stroke", 3)
        if s["side"] == "left":
            x_in = x + d
        else:
            x_in = x - d
        return (
            f'<line x1="{x_in}" y1="{y1}" x2="{x}" y2="{y1}" stroke="#000" stroke-width="{sw}" stroke-linecap="round"/>'
            f'<line x1="{x}" y1="{y1}" x2="{x}" y2="{y2}" stroke="#000" stroke-width="{sw}" stroke-linecap="round"/>'
            f'<line x1="{x_in}" y1="{y2}" x2="{x}" y2="{y2}" stroke="#000" stroke-width="{sw}" stroke-linecap="round"/>'
        )
    if t == "shield":
        # Simple shield approximation: rounded rect with bottom point
        cx, cy, w, h = s["cx"], s["cy"], s["w"], s["h"]
        color = s.get("color", "#000")
        x1, x2 = cx - w/2, cx + w/2
        y1 = cy - h/2
        y2_flat = cy + h/4   # where straight sides end
        y2_tip = cy + h/2    # bottom point
        return (
            f'<path d="M {x1} {y1+6} '
            f'Q {x1} {y1} {x1+6} {y1} '
            f'L {x2-6} {y1} '
            f'Q {x2} {y1} {x2} {y1+6} '
            f'L {x2} {y2_flat} '
            f'L {cx} {y2_tip} '
            f'L {x1} {y2_flat} Z" fill="{color}"/>'
        )
    if t == "shield_stroke":
        # Heraldic shield outline: V-notch top, vertical upper sides, diagonal taper to sharp point.
        cx, cy, w, h = s["cx"], s["cy"], s["w"], s["h"]
        color = s.get("color", "#000")
        sw = s.get("stroke", 3)
        x1, x2 = cx - w/2, cx + w/2
        y_top = cy - h/2
        y_notch = y_top + h*0.10
        y_shoulder = y_top + h*0.04
        y_taper = cy - h*0.05    # taper begins above center
        y_curve = cy + h*0.25    # curve toward tip starts here
        y_tip = cy + h/2
        # V-notch top, vertical sides for upper portion, gradual taper, ending with quick curve to sharp tip
        return (
            f'<path d="M {cx} {y_notch} '
            f'L {x1+3} {y_top} '
            f'L {x1} {y_shoulder} '
            f'L {x1} {y_taper} '
            f'L {x1+w*0.15} {y_curve} '
            f'L {cx} {y_tip} '
            f'L {x2-w*0.15} {y_curve} '
            f'L {x2} {y_taper} '
            f'L {x2} {y_shoulder} '
            f'L {x2-3} {y_top} Z" '
            f'fill="none" stroke="{color}" stroke-width="{sw}" stroke-linejoin="round" stroke-linecap="round"/>'
        )
    if t == "star":
        # 5-pointed star, filled
        import math
        cx, cy, r = s["cx"], s["cy"], s["r"]
        color = s.get("color", "#000")
        inner_ratio = s.get("inner_ratio", 0.5)
        r_inner = r * inner_ratio
        points = []
        for i in range(10):
            angle = -math.pi/2 + i * math.pi/5
            rr = r if i % 2 == 0 else r_inner
            px = cx + rr * math.cos(angle)
            py = cy + rr * math.sin(angle)
            points.append(f"{px:.1f},{py:.1f}")
        return f'<polygon points="{" ".join(points)}" fill="{color}"/>'
    return ""


def _embed_fonts():
    """Return <defs><style> block embedding fonts as data URLs."""
    import base64
    import os
    css = []
    home = os.path.expanduser("~")
    candidates = [
        (f"{home}/.local/share/fonts/playfair/PlayfairDisplay-VF.ttf", "Playfair Display", "900", "normal"),
        (f"{home}/.local/share/fonts/cormorant/CormorantGaramond-Bold.otf", "Cormorant Garamond", "700", "normal"),
    ]
    for path, family, weight, style in candidates:
        if not os.path.exists(path):
            continue
        with open(path, "rb") as f:
            b = base64.b64encode(f.read()).decode("ascii")
        fmt = "opentype" if path.endswith(".otf") else "truetype"
        css.append(
            f"@font-face {{ font-family: '{family}'; font-weight: {weight}; "
            f"font-style: {style}; src: url(data:font/{fmt};base64,{b}) format('{fmt}'); }}"
        )
    if not css:
        return ""
    return f'<defs><style type="text/css">{chr(10).join(css)}</style></defs>'


def main():
    if len(sys.argv) < 2:
        print("usage: render.py <figure_dir>", file=sys.stderr)
        sys.exit(1)
    fig_dir = Path(sys.argv[1])
    gt = json.load(open(fig_dir / "source" / "ground_truth.json"))
    W, H = gt["frame"]

    out = [
        '<?xml version="1.0" encoding="UTF-8"?>',
        f'<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 {W} {H}" '
        f'width="{W}" height="{H}">',
        # @font-face is injected at the HTML level by validate.mjs, not here.
        # No opaque background fill — the SVG layer must be transparent so
        # overlay/diff with the original is honest (white fills mask reality).
    ]

    # Layer 1: shapes that are background (rect_fill, sidebar_box)
    bg_shapes = [s for s in gt["shapes"] if s["type"] in ("rect_fill", "rect_stroke")]
    for s in bg_shapes:
        out.append(render_shape(s))

    # Layer 2: lines (axes, dividers, grid)
    for s in gt["shapes"]:
        if s["type"] in ("line_h", "line_v"):
            out.append(render_shape(s))

    # Tick marks (regular small ticks + emphasized cut ticks)
    for tm in gt.get("tick_marks", []):
        y = tm["axis_y"]
        cut_x = set(tm.get("cut_x", []))
        for x in tm["ticks_x"]:
            if x in cut_x:
                # Cut tick: slightly thicker, same height as regular ticks
                out.append(f'<line x1="{x}" y1="{y-14}" x2="{x}" y2="{y+14}" stroke="#000" stroke-width="5"/>')
            else:
                out.append(f'<line x1="{x}" y1="{y-12}" x2="{x}" y2="{y+12}" stroke="#000" stroke-width="3"/>')

    # Layer 2b: arrows + dashed gridlines + polylines (line graphs)
    for s in gt["shapes"]:
        if s["type"] in ("arrow_h", "arrow", "dashed_line_h", "polyline", "downward_brace"):
            out.append(render_shape(s))
    # Heavy quote ornament (between layers — sits in margin area)
    for s in gt["shapes"]:
        if s["type"] == "heavy_quote_open":
            out.append(render_shape(s))

    # Layer 3: filled circles (markers, alert icon bg)
    for s in gt["shapes"]:
        if s["type"] in ("circle_fill", "circle_stroke", "circle_half"):
            out.append(render_shape(s))
    # Layer 3c: foreground rectangles drawn over circles (e.g., center boxes)
    for s in gt["shapes"]:
        if s["type"] == "rect_fg_fill":
            ss = dict(s); ss["type"] = "rect_fill"
            out.append(render_shape(ss))
        elif s["type"] == "rect_fg_stroke":
            ss = dict(s); ss["type"] = "rect_stroke"
            out.append(render_shape(ss))
    # Layer 3b: brackets and shields (paths)
    for s in gt["shapes"]:
        if s["type"] in ("bracket", "shield", "shield_stroke", "star"):
            out.append(render_shape(s))
    for s in gt["shapes"]:
        if s["type"] == "glyph_centered":
            out.append(render_shape(s))

    # Layer 4: elements (text + badges)
    for el in gt["elements"]:
        role = el["role"]
        if role == "badge_dark":
            out.append(render_badge_dark(el))
        elif role in ("badge_circle", "badge_circle_star", "sidebar_num"):
            out.append(render_badge_circle(el))
        else:
            style = ROLE_STYLES.get(role)
            if style is None:
                print(f"warn: unknown role '{role}' for {el['id']}", file=sys.stderr)
                continue
            out.append(render_text_bbox(el, style))

    out.append("</svg>")
    (fig_dir / "figure.svg").write_text("\n".join(out), encoding="utf-8")
    print(f"wrote {fig_dir/'figure.svg'}")


if __name__ == "__main__":
    main()
