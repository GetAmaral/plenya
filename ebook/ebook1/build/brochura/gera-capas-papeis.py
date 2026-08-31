#!/usr/bin/env python3
"""Gera uma capa de brochura por papel de miolo, com a lombada certa para cada um.

A gráfica só fecha o papel na hora do orçamento, então a capa é gerada em seis
versões e quem escolhe é ela. O nome do arquivo carrega papel e lombada para
não haver troca na hora de enviar.

    lombada_mm = (páginas / 2) × espessura_mm_por_folha

O número de páginas é lido do próprio miolo — não é parâmetro, para a capa não
poder ficar dessincronizada dele. Mudou o miolo, roda de novo.

Uso:
    ./versaoImpressa/runpy brochura/gera-capas-papeis.py
    ./versaoImpressa/runpy brochura/gera-capas-papeis.py --miolo <arquivo.pdf>
"""

import re
import shutil
import subprocess
import sys
from pathlib import Path

THIS_DIR = Path(__file__).resolve().parent
OUT_DIR = THIS_DIR / "capas-papeis"
COVER_SCRIPT = THIS_DIR / "build-brochura-cover.py"
# build-brochura-cover.py depende de img2pdf, que vive no python do sistema e
# não na .venv-figures (essa é só matplotlib, para as figuras). Este script pode
# ser chamado pelos dois, então o interpretador da capa é resolvido na mão.
PYTHON_CAPA = shutil.which("python3") or sys.executable
MIOLO_PADRAO = THIS_DIR / "Antes-pt-BR-brochura-miolo-cor.pdf"
FLAP_MM = 80.0

# Espessura por folha (mm). Fonte: tabela publicada pela PoloPrinter em
# https://poloprinter.com.br/calcule-lombada/ — a única das três calculadoras
# consultadas que abre os números (Ubaldo e BH Gráfica não publicam).
#
# RESSALVA, ainda em aberto desde 2026-08-25: a tabela dá 0,100 mm/folha para o
# Pólen Bold 90g, o MESMO valor do Couché 115g. Isso implica volume ~1,11, que é
# baixo e contradiz o argumento de venda do papel — a Suzano vende o Pólen Bold
# como papel de alto volume, feito para "aumentar o volume do livro sem aumentar
# o número de páginas". Se o volume real for 1,8, a lombada vai a ~28 mm em vez
# de ~17. Não achei a ficha técnica da Suzano com o número (a loja bloqueia
# acesso automatizado). O Pólen Bold é o papel mais provável para este livro e é
# justamente onde a divergência é maior: PERGUNTAR A ESPESSURA À GRÁFICA antes
# de mandar imprimir.
PAPEIS = [
    ("Polen-Soft-80g",      0.0920),
    ("Polen-Bold-90g",      0.1000),
    ("Couche-Fosco-115g",   0.1000),
    ("Offset-75g",          0.1044),
    ("Couche-Fosco-150g",   0.1200),
    ("Offset-90g",          0.1280),
]


def paginas_do_miolo(pdf: Path) -> int:
    out = subprocess.run(["pdfinfo", str(pdf)], capture_output=True, text=True).stdout
    m = re.search(r"^Pages:\s+(\d+)", out, re.M)
    if not m:
        sys.exit(f"❌ não consegui ler o número de páginas de {pdf}")
    return int(m.group(1))


def main():
    miolo = MIOLO_PADRAO
    if "--miolo" in sys.argv:
        miolo = Path(sys.argv[sys.argv.index("--miolo") + 1])
    if not miolo.exists():
        sys.exit(f"❌ miolo não encontrado: {miolo}")

    paginas = paginas_do_miolo(miolo)
    folhas = paginas / 2
    print(f"Miolo: {miolo.name} — {paginas} páginas ({folhas:.0f} folhas)\n")

    OUT_DIR.mkdir(exist_ok=True)
    # Limpa a rodada anterior: os nomes carregam a lombada, então uma capa velha
    # não é sobrescrita — ficaria lado a lado com a nova, convidando ao engano.
    for velho in OUT_DIR.glob("Antes-capa-brochura_*"):
        velho.unlink()

    for papel, espessura in PAPEIS:
        lombada = round(folhas * espessura, 1)
        r = subprocess.run(
            [PYTHON_CAPA, str(COVER_SCRIPT), "pt-BR", str(lombada), str(FLAP_MM)],
            capture_output=True, text=True, cwd=str(THIS_DIR),
        )
        if r.returncode != 0:
            print(r.stdout[-1500:], r.stderr[-1500:])
            sys.exit(f"❌ falhou em {papel}")

        base = f"Antes-capa-brochura_{papel}_lombada-{lombada}mm"
        for src, ext in ((THIS_DIR / "Antes-pt-BR-brochura-capa.pdf", ".pdf"),
                         (THIS_DIR / "Antes-pt-BR-brochura-capa-guias.png", "-GUIAS.png")):
            if src.exists():
                src.replace(OUT_DIR / (base + ext))
        # O PNG de conferência rápida não vai para capas-papeis: um por papel
        # seriam 100 MB de arquivo idêntico exceto pela largura da lombada.
        solto = THIS_DIR / "Antes-pt-BR-brochura-capa.png"
        if solto.exists():
            solto.unlink()

        largura = FLAP_MM + 160 + lombada + 160 + FLAP_MM + 6
        print(f"  {papel:<20} {espessura:.4f} mm/folha → lombada {lombada:>5} mm"
              f"  (largura total {largura:.1f} mm)")

    print(f"\n✅ {len(PAPEIS)} capas em {OUT_DIR}")
    print("   Confirmar a espessura com a gráfica antes de imprimir — ver a "
          "ressalva do Pólen Bold no topo deste script.")


if __name__ == "__main__":
    main()
