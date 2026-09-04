#!/usr/bin/env python3
"""Ler e escrever a devolutiva do paciente pela API, do lado de cá da conversa.

O EMR monta o rascunho mecânico de graça (`POST .../plans/assemble`). O que sobra é a leitura
clínica: o punch de cada slide, o título como afirmação, o tema da seção. Isso nasce aqui, na
conversa com o médico, e precisa VOLTAR para o EMR — senão o plano bom mora no chat e o plano do
prontuário continua descritivo.

Escrever o deck inteiro à mão a cada ajuste é o que torna esse ciclo caro: são 18 slides de JSON
para trocar três frases, e um `expectedRevision` errado sobrescreve em silêncio quem escreveu antes.
Este utilitário faz o ciclo curto: lê o plano, aplica um punhado de campos por slide, devolve com a
revisão esperada e mede o estouro.

    # ver o deck, numerado, para a conversa poder dizer "slide 7"
    scripts/plano/plano.py ler --paciente <uuid> --plano <uuid>

    # escrever a leitura clínica em três slides
    scripts/plano/plano.py escrever --paciente <uuid> --plano <uuid> --edicoes leitura.json

`leitura.json` é `{"<slide>": {"campo": "valor"}}`, onde `<slide>` é a POSIÇÃO (1, 2, 3…) ou o id:

    {
      "7":  {"punch": "A PCR saiu de 0,37 para 63,10 em quatro meses. <em>Isso precisa de nome antes do resto.</em>"},
      "10": {"title": "A musculação entra junto com a medicação, não depois"}
    }

Campos aceitos: title, punch, lede, kicker, eyebrow. Nada mais: o resto do slide (régua, tabela,
cartão, dose) é dado do prontuário, e dado se corrige no prontuário, não no deck.
"""

import argparse
import json
import sys
import urllib.error
import urllib.request

CAMPOS = ("title", "punch", "lede", "kicker", "eyebrow")


def chama(base, token, metodo, caminho, corpo=None):
    req = urllib.request.Request(
        base.rstrip("/") + caminho,
        method=metodo,
        data=json.dumps(corpo).encode() if corpo is not None else None,
        headers={
            "Authorization": "Bearer " + token,
            "Content-Type": "application/json",
        },
    )
    try:
        with urllib.request.urlopen(req, timeout=300) as r:
            bruto = r.read()
            return json.loads(bruto) if bruto else None
    except urllib.error.HTTPError as e:
        # O corpo do erro é onde o servidor explica (409 de revisão, 422 de publicação com
        # estouro). Engolir isso e mostrar só "HTTP 409" transformaria cada falha em adivinhação.
        detalhe = e.read().decode("utf-8", "replace")[:600]
        sys.exit(f"HTTP {e.code} em {metodo} {caminho}: {detalhe}")


def base_do(args):
    return f"/api/v1/patients/{args.paciente}/plans"


def ler(args):
    plano = chama(args.api, args.token, "GET", f"{base_do(args)}/{args.plano}")
    slides = plano.get("content") or []
    print(f'{plano["title"]}  ·  revisão {plano["revisionSeq"]}  ·  {len(slides)} slides  ·  {plano["status"]}')
    for i, s in enumerate(slides, 1):
        print(f'\n{i:2}. [{s.get("kind","")}] {s.get("title","")}')
        if s.get("eyebrow"):
            print(f'    olho:  {s["eyebrow"]}')
        for campo in ("lede", "punch"):
            if s.get(campo):
                print(f'    {campo}: {s[campo]}')
        for r in s.get("rulers") or []:
            print(f'    régua: {r.get("display","")} — {r.get("note","")}')
        if not s.get("punch") and s.get("kind") not in ("cover", "closing", "takeaway"):
            print("    ⚠ sem punch")


def escrever(args):
    edicoes = json.load(sys.stdin if args.edicoes == "-" else open(args.edicoes, encoding="utf-8"))
    plano = chama(args.api, args.token, "GET", f"{base_do(args)}/{args.plano}")
    slides = plano.get("content") or []
    porID = {s.get("id"): i for i, s in enumerate(slides) if s.get("id")}

    aplicadas = 0
    for chave, campos in edicoes.items():
        # Posição OU id. A conversa fala em "slide 7"; o id é o que sobrevive a uma reordenação.
        if chave.isdigit():
            idx = int(chave) - 1
            if not 0 <= idx < len(slides):
                sys.exit(f"slide {chave} não existe: o deck tem {len(slides)}")
        elif chave in porID:
            idx = porID[chave]
        else:
            sys.exit(f"slide {chave!r} não encontrado")
        for campo, valor in campos.items():
            if campo not in CAMPOS:
                sys.exit(f"campo {campo!r} não é de texto; corrija o dado no prontuário, não no deck")
            slides[idx][campo] = valor
            aplicadas += 1
        print(f'{idx+1:2}. {slides[idx].get("title","")[:52]}  ←  {", ".join(campos)}')

    salvo = chama(
        args.api, args.token, "PUT", f"{base_do(args)}/{args.plano}",
        # `expectedRevision` não é opcional: sem ele o PUT cai na saída de emergência do servidor e
        # sobrescreve em silêncio quem escreveu entre o GET e o PUT (outra aba, outro clínico).
        {"title": plano["title"], "content": slides, "expectedRevision": plano["revisionSeq"]},
    )
    print(f'\n{aplicadas} campos aplicados · revisão {plano["revisionSeq"]} → {salvo["revisionSeq"]}')

    # O slide tem altura fixa e `overflow:hidden`: o que não cabe some do PDF sem erro nenhum.
    # Medir depois de escrever é o único jeito de saber que o texto novo coube.
    # A resposta é `{"slides":[...]}`, não uma lista solta.
    estouro = (chama(args.api, args.token, "GET", f"{base_do(args)}/{args.plano}/overflow") or {}).get("slides") or []
    if not estouro:
        print("cabe: nenhum slide estourando")
        return
    for o in estouro:
        print(f'⚠ slide {o["slide"]} estoura {o.get("bottom",0):.0f}px de altura, {o.get("right",0):.0f}px de largura')
    sys.exit(1)


def main():
    p = argparse.ArgumentParser(description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    p.add_argument("--api", default="http://localhost:3001", help="base da API (default: dev local)")
    p.add_argument("--token", default="dev", help="Bearer; em dev com DEV_BYPASS_AUTH qualquer valor serve")
    sub = p.add_subparsers(dest="cmd", required=True)
    for nome, fn in (("ler", ler), ("escrever", escrever)):
        s = sub.add_parser(nome)
        s.add_argument("--paciente", required=True)
        s.add_argument("--plano", required=True)
        if nome == "escrever":
            s.add_argument("--edicoes", required=True, help='arquivo JSON, ou "-" para stdin')
        s.set_defaults(fn=fn)
    args = p.parse_args()
    args.fn(args)


if __name__ == "__main__":
    main()
