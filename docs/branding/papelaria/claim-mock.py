# -*- coding: utf-8 -*-
"""Mockups de posicionamento do claim 'Viva bem, viva mais.' como marca d'água.
Reusa header/footer/css do letterhead. Saída: drafts/claim-{A,B,C}-*.html (preview RGB)."""
import letterhead as L

CLAIM = "Viva bem, viva mais."
SYM = L.a("symbol-petrol.svg")

CLAIM_CSS = """
/* A — lateral vertical, margem esquerda */
.claim-left { position:absolute; left:9mm; top:50%; transform-origin:left center;
   transform:rotate(-90deg) translateX(-50%);
   font-family:'Cormorant',serif; font-weight:500; font-size:19.2px; letter-spacing:3.6px;
   color:var(--petrol); opacity:.25; white-space:nowrap; }
/* B — rodapé centralizado, acima do NAP */
.claim-foot { position:absolute; left:0; right:0; bottom:30mm; text-align:center;
   font-family:'Cormorant',serif; font-weight:500; font-size:15px; letter-spacing:2.5px;
   color:var(--petrol); opacity:.20; }
/* C — central grande (substitui o P∞) */
.claim-center { position:absolute; left:50%; top:50%; transform:translate(-50%,-50%);
   font-family:'Cormorant',serif; font-weight:500; font-size:48px; letter-spacing:2px;
   color:var(--petrol); opacity:.08; white-space:nowrap; }
/* D — logo abaixo do P∞ central */
.claim-belowP { position:absolute; left:50%; top:calc(50% + 47mm); transform:translateX(-50%);
   font-family:'Cormorant',serif; font-weight:500; font-size:25px; letter-spacing:2px;
   color:var(--petrol); opacity:.13; white-space:nowrap; text-align:center; }
/* E/F — abaixo do P∞, alinhado à borda DIREITA / ESQUERDA do glifo (P vai de 47.7 a 162.2mm trim) */
.claim-Ralign { position:absolute; right:47.8mm; top:calc(50% + 47mm + 15px);
   font-family:'Cormorant',serif; font-weight:500; font-size:25px; letter-spacing:2px;
   color:var(--petrol); opacity:.10; white-space:nowrap; }
.claim-Lalign { position:absolute; left:47.7mm; top:calc(50% + 47mm);
   font-family:'Cormorant',serif; font-weight:500; font-size:25px; letter-spacing:2px;
   color:var(--petrol); opacity:.13; white-space:nowrap; }
"""

def page(claim_html, with_pinf=True):
    wm = f'<img class="wm" src="{SYM}">' if with_pinf else ''
    inner = (f'{wm}{claim_html}<div class="frame">{L.header()}'
             f'<div class="body"></div><div class="foot">{L.footer_nap()}</div></div>')
    sheet = (f'<div class="sheet"><div class="bleed"></div>'
             f'<div class="trim">{inner}</div>{L.crop_marks()}</div>')
    return (f"<!doctype html><html><head><meta charset='utf-8'>"
            f"<style>{L.css('#fff')}{CLAIM_CSS}</style></head><body>{sheet}</body></html>")

def page_estampado(claim_html, paper):
    pat = f'<img class="estampa" src="{L.a("pattern.svg")}">'
    inner = (f'{claim_html}<div class="frame">{L.header()}'
             f'<div class="body"></div><div class="foot">{L.footer_nap()}</div></div>')
    sheet = (f'<div class="sheet"><div class="bleed">{pat}</div>'
             f'<div class="trim">{inner}</div>{L.crop_marks()}</div>')
    return (f"<!doctype html><html><head><meta charset='utf-8'>"
            f"<style>{L.css(paper)}{CLAIM_CSS}</style></head><body>{sheet}</body></html>")

mocks = {
    "claim-A-lateral": page(f'<div class="claim-left">{CLAIM}</div>', with_pinf=True),
    "claim-B-rodape":  page(f'<div class="claim-foot">{CLAIM}</div>', with_pinf=True),
    "claim-C-central": page(f'<div class="claim-center">{CLAIM}</div>', with_pinf=False),
    "claim-D-belowP":  page(f'<div class="claim-belowP">{CLAIM}</div>', with_pinf=True),
    "claim-E-direita": page(f'<div class="claim-Ralign">{CLAIM}</div>', with_pinf=True),
    "claim-F-esquerda":page(f'<div class="claim-Lalign">{CLAIM}</div>', with_pinf=True),
    "claim-estampado-lateral": page_estampado(f'<div class="claim-left">{CLAIM}</div>', "var(--cream)"),
}
for slug, html in mocks.items():
    (L.DRAFTS / f"{slug}.html").write_text(html, encoding="utf-8")
    print("wrote", slug)
