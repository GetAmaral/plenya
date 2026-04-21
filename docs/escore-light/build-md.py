#!/usr/bin/env python3
"""Convert /tmp/score_dump.psv → docs/escore-light/items-with-levels.md
One-shot script for Light curation reference. Safe to delete after use.
"""
from collections import defaultdict
from pathlib import Path

SRC = Path("/tmp/score_dump.psv")
OUT = Path("/home/user/plenya/docs/escore-light/items-with-levels.md")

# tree[group][subgroup][item_id] = {meta..., 'levels': [...]}
tree = defaultdict(lambda: defaultdict(dict))

for raw in SRC.read_text(encoding="utf-8").splitlines():
    parts = raw.split("|")
    if len(parts) < 15:
        continue
    g, sg, iid, iname, gender, unit, points, amin, amax, pmeno, lvl, op, lo, hi, lname = parts[:15]
    item = tree[g][sg].setdefault(iid, {
        "name": iname, "gender": gender, "unit": unit, "points": points,
        "amin": amin, "amax": amax, "pmeno": pmeno, "levels": [],
    })
    if lvl != "":
        item["levels"].append({"level": lvl, "op": op, "lo": lo, "hi": hi, "lname": lname})

# Write MD
out = []
out.append("# Escore Plenya — Items com Levels (referência para curadoria do Light)\n")
out.append("Gerado a partir de `score_items` + `score_levels`. **Items sem levels são omitidos** (não podem entrar no Light — exigem julgamento profissional).\n")
out.append("Estrutura: `Grupo → Subgrupo → Item (gênero, max_points) → Levels (nível, operador, faixa, label)`\n\n")
out.append("---\n\n")

n_items_with_levels = 0
n_items_without = 0

for g in sorted(tree.keys()):
    out.append(f"## {g}\n\n")
    for sg in sorted(tree[g].keys()):
        # Filter items: only those with levels
        items_with = [(iid, it) for iid, it in tree[g][sg].items() if it["levels"]]
        items_without = [it for iid, it in tree[g][sg].items() if not it["levels"]]
        n_items_with_levels += len(items_with)
        n_items_without += len(items_without)
        if not items_with:
            continue
        out.append(f"### {sg}\n\n")
        # Sort by points desc
        items_with.sort(key=lambda x: -float(x[1]["points"] or 0))
        for iid, it in items_with:
            tags = []
            if it["gender"] and it["gender"] != "not_applicable":
                tags.append(it["gender"])
            if it["amin"] or it["amax"]:
                tags.append(f"idade {it['amin']}-{it['amax']}")
            if it["pmeno"] == "true":
                tags.append("pós-menopausa")
            if it["unit"]:
                tags.append(it["unit"])
            tag_str = f" _({', '.join(tags)})_" if tags else ""
            out.append(f"- **{it['name']}**{tag_str} — `{it['points']} pts` · `{iid}`\n")
            for lv in sorted(it["levels"], key=lambda x: int(x["level"]) if x["level"].isdigit() else 99):
                rng = ""
                if lv["op"] == "between":
                    rng = f"{lv['lo']}–{lv['hi']}"
                elif lv["op"] in ("=", ""):
                    rng = lv["lo"] or lv["hi"] or ""
                else:
                    rng = f"{lv['op']} {lv['lo'] or lv['hi']}"
                out.append(f"  - `L{lv['level']}` {rng:>16} — {lv['lname']}\n")
            out.append("\n")

out.append(f"\n---\n\n## Estatísticas\n\n")
out.append(f"- Items com levels: **{n_items_with_levels}**\n")
out.append(f"- Items sem levels (excluídos do Light): **{n_items_without}**\n")

OUT.write_text("".join(out), encoding="utf-8")
print(f"Wrote {OUT} — {OUT.stat().st_size} bytes, {n_items_with_levels} items with levels, {n_items_without} omitted")
