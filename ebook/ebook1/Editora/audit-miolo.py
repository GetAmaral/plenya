#!/usr/bin/env python3
"""Auditoria do miolo — v3. Classes de defeito apontadas pelo editor, no livro inteiro."""
import sys, re, collections
import pymupdf as fitz

PDF = sys.argv[1]
doc = fitz.open(PDF)
REC, VER = (79.37, 416.70), (51.02, 388.34)
IND, LEAD, BODY = 13.09, 14.90, "TeXGyrePagella"
HS = {25.7: "cap", 11.0: "sec", 10.0: "sub"}

def lines_of(p):
    """Agrupa spans em linhas visuais; subscrito/sobrescrito entra na linha-mãe."""
    raw = []
    for b in p.get_text("dict")["blocks"]:
        if b["type"] != 0: continue
        for l in b["lines"]:
            for s in l["spans"]:
                if s["text"].strip(): raw.append(s)
    raw.sort(key=lambda s: (s["origin"][1], s["bbox"][0]))
    rows = []
    for s in raw:
        placed = False
        for r in rows:
            if abs(r["ref"] - s["origin"][1]) < 4.5:
                r["sp"].append(s); placed = True
                if s["size"] > r["maxsz"]: r["maxsz"], r["ref"] = s["size"], s["origin"][1]
                break
        if not placed:
            rows.append(dict(ref=s["origin"][1], maxsz=s["size"], sp=[s]))
    out = []
    for r in sorted(rows, key=lambda r: r["ref"]):
        sp = sorted(r["sp"], key=lambda s: s["bbox"][0])
        big = max(sp, key=lambda s: s["size"])
        out.append(dict(y=r["ref"], x0=min(s["bbox"][0] for s in sp), x1=max(s["bbox"][2] for s in sp),
                        text="".join(s["text"] for s in sp), size=round(big["size"], 1),
                        font=big["font"].split("+")[-1]))
    return out

P = []
for i, p in enumerate(doc):
    n = i + 1
    L, R = REC if n % 2 else VER
    ls = lines_of(p)
    head = [l for l in ls if l["y"] < 60]
    folio = None
    for l in head:
        m = re.search(r"\b(\d+)\b", l["text"])
        if m: folio = int(m.group(1))
    body, para, heads, caps, lst = [], [], [], [], []
    for l in ls:
        if l["y"] < 60: continue
        f, s = l["font"], l["size"]
        if f.startswith(BODY) and abs(s - 10.9) < .35 and L - 2 < l["x0"] < L + 30:
            body.append(l)
            rel = l["x0"] - L
            (para if rel < 2 or abs(rel - IND) < 3 else lst).append(l)
        elif f == "Inter-ExtraBold" and s in HS and l["x0"] < L + 3:
            heads.append(l)
        elif f.startswith("Inter") and abs(s - 9.2) < .2 and L - 2 < l["x0"] < L + 4:
            caps.append(l)
    P.append(dict(n=n, L=L, R=R, ls=ls, head=head, folio=folio,
                  body=body, para=para, lst=lst, heads=heads, caps=caps))

F = collections.defaultdict(list)

# 1 fólio contínuo
for pg in P:
    if pg["folio"] is not None and pg["folio"] != pg["n"]:
        F["folio"].append(f"p.{pg['n']}: cabeço imprime {pg['folio']}")

# 2 página em branco com cabeço/fólio
for pg in P:
    if not pg["body"] and not [l for l in pg["ls"] if l["y"] >= 60] and pg["head"]:
        F["branca_com_cabeco"].append(f"p.{pg['n']}: “{' '.join(l['text'] for l in pg['head'])}”")

# 3 hífen em título/legenda
for pg in P:
    for l in pg["heads"] + pg["caps"]:
        if l["text"].rstrip().endswith(("-", "‐")):
            F["hifen_titulo"].append(f"p.{pg['n']}: “{l['text'].strip()}”")

# 4 capitular
for pg in P:
    ch = [h for h in pg["heads"] if HS[h["size"]] == "cap"]
    if not ch: continue
    drop = [l for l in pg["ls"] if l["font"].startswith(BODY) and l["size"] > 24 and l["y"] > ch[-1]["y"]]
    if not drop:
        first = [l for l in pg["body"] if l["y"] > ch[-1]["y"]]
        F["sem_capitular"].append(f"p.{pg['n']}: “{' '.join(h['text'].strip() for h in ch)}” — 1º §: “{first[0]['text'].strip()[:58] if first else ''}”")

# 5 recuo depois de subtítulo
for pg in P:
    for h in pg["heads"]:
        if HS[h["size"]] == "cap": continue
        aft = [l for l in pg["body"] if l["y"] > h["y"] + 2]
        if aft and aft[0]["x0"] < pg["L"] + IND - 4:
            F["sem_recuo"].append(f"p.{pg['n']}: “{h['text'].strip()[:40]}”")

# 6/7 justificação e transbordo
for pg in P:
    b = pg["body"]
    for j, l in enumerate(b):
        if j + 1 < len(b):
            nx = b[j + 1]
            if abs(nx["y"] - l["y"] - LEAD) < 1.6 and nx["x0"] < pg["L"] + 6 and l["x1"] < pg["R"] - 3:
                F["corpo_ragged"].append(f"p.{pg['n']}: falta {pg['R']-l['x1']:.1f}pt “…{l['text'].strip()[-42:]}”")
        if l["x1"] > pg["R"] + 2:
            F["transborda"].append(f"p.{pg['n']}: +{l['x1']-pg['R']:.1f}pt “…{l['text'].strip()[-42:]}”")
    c = pg["caps"]
    for j, l in enumerate(c):
        if j + 1 < len(c) and 0 < c[j+1]["y"] - l["y"] < 14 and l["x1"] < pg["R"] - 3:
            F["legenda_ragged"].append(f"p.{pg['n']}: “…{l['text'].strip()[-42:]}”")

# 8 hifenização
tot = hy = 0
for pg in P:
    run = 0
    for l in pg["body"]:
        tot += 1
        if l["text"].rstrip().endswith(("-", "‐")):
            hy += 1; run += 1
            w = l["text"].rstrip().split()[-1].rstrip("-")
            if len(w) <= 2: F["hifen_curto"].append(f"p.{pg['n']}: “…{l['text'].strip()[-24:]}”")
        else:
            if run >= 3: F["escada_hifens"].append(f"p.{pg['n']}: {run} seguidos")
            run = 0
    if run >= 3: F["escada_hifens"].append(f"p.{pg['n']}: {run} seguidos")
    if pg["body"] and pg["body"][-1]["text"].rstrip().endswith("-"):
        F["hifen_virada"].append(f"p.{pg['n']}: “…{pg['body'][-1]['text'].strip()[-28:]}”")

# 9 viúvas / órfãs (só parágrafos de texto corrido, listas de fora)
for k, pg in enumerate(P):
    pa = pg["para"]
    if len(pa) < 2: continue
    if pa[-1]["x0"] - pg["L"] > IND - 3 and pa[-1]["y"] > 545:
        F["orfa"].append(f"p.{pg['n']}: “{pa[-1]['text'].strip()[:58]}…”")
    prev = P[k-1]["para"] if k else []
    if prev and abs(pa[0]["x0"] - pg["L"]) < 3 and pa[1]["x0"] - pg["L"] > IND - 3 \
       and prev[-1]["x1"] > P[k-1]["R"] - 4 and pa[0]["y"] < 130:
        F["viuva"].append(f"p.{pg['n']}: “{pa[0]['text'].strip()[:58]}”")

# 10 linha final de parágrafo muito curta (linha-clube)
for pg in P:
    pa = pg["para"]
    for j, l in enumerate(pa):
        nxt_new = (j + 1 >= len(pa)) or (pa[j+1]["x0"] - pg["L"] > IND - 3) or (pa[j+1]["y"] - l["y"] > 20)
        if nxt_new and j > 0 and (l["x1"] - l["x0"]) < 42:
            F["linha_clube"].append(f"p.{pg['n']}: {l['x1']-l['x0']:.0f}pt “{l['text'].strip()[:28]}”")

# 11 páginas sem corpo
for pg in P:
    if pg["body"]: continue
    cont = [l for l in pg["ls"] if l["y"] >= 60]
    has_img = any(b["type"] != 0 for b in doc[pg["n"]-1].get_text("dict")["blocks"])
    if not cont and not has_img: kind = "branca"
    elif any(abs(l["size"] - 34.8) < 1 for l in cont): kind = "abertura de parte"
    elif has_img or len(cont) > 15: kind = "FIGURA DE PÁGINA INTEIRA"
    else: kind = "outro"
    F["pagina_sem_corpo"].append(f"p.{pg['n']} — {kind}")

# 12 mancha curta (texto acaba muito acima do pé)
for pg in P:
    b = [l for l in pg["ls"] if l["y"] >= 60]
    if not pg["body"] or len(pg["body"]) < 3: continue
    ybot = max(l["y"] for l in b)
    has_img = any(bb["type"] != 0 for bb in doc[pg["n"]-1].get_text("dict")["blocks"])
    if ybot < 500 and not has_img:
        F["mancha_curta"].append(f"p.{pg['n']}: texto acaba em y={ybot:.0f} (pé ≈ 604) — {604-ybot:.0f}pt de branco")

# 13 figura x chamada
figp, refp = {}, collections.defaultdict(list)
for pg in P:
    for l in pg["caps"]:
        m = re.match(r"\s*Figura\s+(\d+\.\d+)", l["text"])
        if m: figp.setdefault(m.group(1), pg["n"])
    t = " ".join(l["text"] for l in pg["body"])
    for m in re.finditer(r"[Ff]igura\s+(\d+\.\d+)", t):
        refp[m.group(1)].append(pg["n"])
for fig in sorted(figp, key=lambda k: figp[k]):
    fp, rp = figp[fig], refp.get(fig)
    if not rp: F["figura_sem_chamada"].append(f"Figura {fig} — arte na p.{fp}")
    else:
        d = min(abs(fp - r) for r in rp)
        if d >= 2: F["figura_longe"].append(f"Figura {fig}: arte p.{fp}, chamada p.{rp}")

print(f"### {PDF} — {doc.page_count} pp · {tot} linhas de corpo · hífen no fim: {hy} ({100*hy/tot:.1f}%)\n")
for k in ["folio","branca_com_cabeco","hifen_titulo","sem_capitular","sem_recuo","corpo_ragged",
          "legenda_ragged","transborda","escada_hifens","hifen_curto","hifen_virada",
          "viuva","orfa","linha_clube","mancha_curta","pagina_sem_corpo","figura_longe","figura_sem_chamada"]:
    v = F.get(k, [])
    print(f"== {k}: {len(v)}")
    for x in v[:70]: print("   ", x)
    if len(v) > 70: print(f"    … +{len(v)-70}")
    print()
