#!/usr/bin/env python3
"""extract.py <figure_dir>

Reads <figure_dir>/_reference.png and emits:
  <figure_dir>/source/texts_raw.json   (Tesseract OCR, lang=por, word-level bboxes)
  <figure_dir>/source/shapes.json       (OpenCV: rectangles, circles, long lines)
  <figure_dir>/source/_overlay.png      (full-size: original + bboxes drawn)
  <figure_dir>/source/_overlay_grid.png (full-size: overlay + 400x400 grid)
"""
import sys
import json
from pathlib import Path

import cv2
import numpy as np
import pytesseract
from pytesseract import Output


def run_ocr(img_rgb):
    """Return list of {text, bbox:[x1,y1,x2,y2], conf, level} (word level)."""
    data = pytesseract.image_to_data(
        img_rgb, lang="por", config="--psm 11",
        output_type=Output.DICT,
    )
    out = []
    n = len(data["text"])
    for i in range(n):
        txt = (data["text"][i] or "").strip()
        if not txt:
            continue
        try:
            conf = float(data["conf"][i])
        except ValueError:
            conf = -1.0
        if conf < 30:
            continue
        x, y, w, h = data["left"][i], data["top"][i], data["width"][i], data["height"][i]
        out.append({
            "text": txt,
            "bbox": [int(x), int(y), int(x + w), int(y + h)],
            "conf": round(conf, 1),
            "block": int(data["block_num"][i]),
            "line": int(data["line_num"][i]),
            "word": int(data["word_num"][i]),
        })
    return out


def group_words_into_lines(words):
    """Group consecutive words with same (block,line) into a single line entry."""
    if not words:
        return []
    lines = {}
    order = []
    for w in words:
        key = (w["block"], w["line"])
        if key not in lines:
            lines[key] = []
            order.append(key)
        lines[key].append(w)
    out = []
    for key in order:
        ws = sorted(lines[key], key=lambda w: w["bbox"][0])
        x1 = min(w["bbox"][0] for w in ws)
        y1 = min(w["bbox"][1] for w in ws)
        x2 = max(w["bbox"][2] for w in ws)
        y2 = max(w["bbox"][3] for w in ws)
        text = " ".join(w["text"] for w in ws)
        conf = round(sum(w["conf"] for w in ws) / len(ws), 1)
        out.append({"text": text, "bbox": [x1, y1, x2, y2], "conf": conf})
    return out


def detect_shapes(img_bgr):
    """Heuristic shape detection.

    Strategy: instead of trying to be clever about shape vs text vs noise,
    we detect large connected components in a binarized "non-white" mask.
    The verification pass (me reading _overlay.png) decides which are kept.
    """
    H, W = img_bgr.shape[:2]
    gray = cv2.cvtColor(img_bgr, cv2.COLOR_BGR2GRAY)
    # Anything not near-white counts as ink
    _, mask = cv2.threshold(gray, 235, 255, cv2.THRESH_BINARY_INV)

    rects, circles = [], []
    contours, _ = cv2.findContours(mask, cv2.RETR_EXTERNAL, cv2.CHAIN_APPROX_SIMPLE)
    for c in contours:
        area = cv2.contourArea(c)
        if area < 5000:  # ignore tiny noise
            continue
        x, y, w, h = cv2.boundingRect(c)
        perim = cv2.arcLength(c, True)
        if perim == 0:
            continue
        circularity = 4 * np.pi * area / (perim * perim)
        aspect = w / max(h, 1)
        if circularity > 0.75 and 0.85 < aspect < 1.18 and w < 200:
            circles.append({"cx": int(x + w / 2), "cy": int(y + h / 2), "r": int((w + h) / 4), "area": int(area)})
        else:
            # Filter: shape must be reasonably large to be a structural rect
            if w > 80 and h > 30 and (w > 200 or h > 200):
                rects.append({"bbox": [int(x), int(y), int(x + w), int(y + h)], "area": int(area)})

    # Long thin lines via Hough
    edges = cv2.Canny(gray, 60, 180)
    hlines = cv2.HoughLinesP(edges, 1, np.pi / 180, threshold=200, minLineLength=200, maxLineGap=12)
    lines = []
    if hlines is not None:
        for l in hlines:
            x1, y1, x2, y2 = l[0]
            lines.append({"p1": [int(x1), int(y1)], "p2": [int(x2), int(y2)]})

    return {"rects": rects, "circles": circles, "lines": lines}


def draw_overlay(img_bgr, texts, shapes, grid=False):
    out = img_bgr.copy()
    # Texts in green
    for t in texts:
        x1, y1, x2, y2 = t["bbox"]
        cv2.rectangle(out, (x1, y1), (x2, y2), (0, 200, 0), 2)
    # Rects in blue
    for r in shapes["rects"]:
        x1, y1, x2, y2 = r["bbox"]
        cv2.rectangle(out, (x1, y1), (x2, y2), (255, 0, 0), 3)
    # Circles in red
    for c in shapes["circles"]:
        cv2.circle(out, (c["cx"], c["cy"]), c["r"], (0, 0, 255), 3)
    # Lines in magenta
    for ln in shapes["lines"]:
        cv2.line(out, tuple(ln["p1"]), tuple(ln["p2"]), (255, 0, 255), 2)
    if grid:
        H, W = out.shape[:2]
        for x in range(0, W, 400):
            cv2.line(out, (x, 0), (x, H), (200, 200, 200), 1)
            cv2.putText(out, str(x), (x + 4, 30), cv2.FONT_HERSHEY_SIMPLEX, 0.8, (0, 0, 0), 2)
        for y in range(0, H, 400):
            cv2.line(out, (0, y), (W, y), (200, 200, 200), 1)
            cv2.putText(out, str(y), (4, y + 30), cv2.FONT_HERSHEY_SIMPLEX, 0.8, (0, 0, 0), 2)
    return out


def main():
    if len(sys.argv) < 2:
        print("usage: extract.py <figure_dir>", file=sys.stderr)
        sys.exit(1)
    fig_dir = Path(sys.argv[1])
    ref = fig_dir / "_reference.png"
    if not ref.exists():
        print(f"no reference at {ref}", file=sys.stderr)
        sys.exit(1)
    out_dir = fig_dir / "source"
    out_dir.mkdir(exist_ok=True)

    img_bgr = cv2.imread(str(ref))
    H, W = img_bgr.shape[:2]

    words = run_ocr(cv2.cvtColor(img_bgr, cv2.COLOR_BGR2RGB))
    lines = group_words_into_lines(words)
    shapes = detect_shapes(img_bgr)

    (out_dir / "texts_raw.json").write_text(json.dumps(
        {"frame": [W, H], "words": words, "lines": lines},
        ensure_ascii=False, indent=2,
    ))
    (out_dir / "shapes.json").write_text(json.dumps(
        {"frame": [W, H], **shapes}, indent=2,
    ))

    overlay = draw_overlay(img_bgr, lines, shapes, grid=False)
    overlay_grid = draw_overlay(img_bgr, lines, shapes, grid=True)
    cv2.imwrite(str(out_dir / "_overlay.png"), overlay)
    cv2.imwrite(str(out_dir / "_overlay_grid.png"), overlay_grid)

    # Also crop-grid 400x400 (for crop-by-crop inspection)
    grid_dir = out_dir / "_crops"
    grid_dir.mkdir(exist_ok=True)
    for col in range(0, W, 400):
        for row in range(0, H, 400):
            crop = overlay_grid[row:row + 400, col:col + 400]
            cv2.imwrite(str(grid_dir / f"c{col:04d}_r{row:04d}.png"), crop)

    print(f"{len(words)} words → {len(lines)} lines")
    print(f"{len(shapes['rects'])} rects, {len(shapes['circles'])} circles, {len(shapes['lines'])} long lines")
    print(f"see → {out_dir}/_overlay.png")


if __name__ == "__main__":
    main()
