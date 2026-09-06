#!/usr/bin/env python3
"""A porta do prontuário: ler e ESCREVER o registro do paciente a partir desta conversa.

Sem isto, tudo o que a conversa produz (uma conduta nova, a leitura clínica de um achado, um exame
que não veio no PDF, a pressão que ninguém digitou) morre no chat: o deck sai bonito e o prontuário
continua sem a informação. O deck é DERIVADO — conduta que só existe no slide some do prontuário,
não entra no próximo escore e não aparece na próxima consulta.

## Por que pela API, e não por psql

A regra de ouro do projeto manda manipular dado em dev pelo banco direto. Ela continua valendo para
o que ela foi escrita: catálogo, escore, carga em massa, correção de fixture.

**Dado de paciente é diferente e vai por HTTP**, porque só o caminho HTTP tem as quatro coisas que
um prontuário precisa ter: os hooks do model (UUID v7, criptografia de CPF/RG, `LastReview`), a
validação do DTO, o RBAC de cada rota (staff na nota, clínico no lote de exames e na conduta,
admin no catálogo) e a LINHA DE AUDITORIA. `middleware.AuditLog` é middleware de rota: um INSERT por psql entra no banco
sem nenhum registro de quem escreveu, e em produção `RevokeAuditLogMutations` revoga
UPDATE/DELETE/TRUNCATE de `audit_logs` justamente para que essa linha não possa ser apagada depois. Escrever no prontuário por fora do HTTP é apagar a própria pegada.

## O que esta porta NÃO faz, de propósito

  - **assinar**. `POST /clinical-notes/:id/sign` é ato médico com certificado ICP-Brasil. As notas
    criadas aqui nascem `draft`, e quem assina é o médico, na tela dele.
  - **apagar**. Nada aqui chama DELETE.
  - **inventar**. Todo valor vem de você ou do que já está no prontuário.

`glosa` é o único subcomando que exige chave de admin; o resto pede clínico ou staff.

## Uso

    export EMR_API=http://localhost:3001          # dev
    export EMR_TOKEN=dev                          # com DEV_BYPASS_AUTH qualquer valor serve

    # PRODUÇÃO — mesma porta, outra URL e um token de verdade:
    export EMR_API=https://api.plenyasaude.com.br
    export EMR_TOKEN=$(scripts/emr/prod-token.sh)
    # Antes de escrever em prod, leia docs/emr/dados-de-paciente-em-producao.md: confira o paciente
    # pelo UUID (há cadastros duplicados), ensaie o script num paciente descartável na dev, e saiba
    # que nada disto é idempotente.

    emr.py ficha    --paciente <uuid>
    emr.py conduta  --paciente <uuid> --letra A --recomendacao "Musculação três vezes por semana" \\
                    --porque "protege massa magra durante a perda de peso" --meta "3 sessões/semana"
    emr.py nota     --paciente <uuid> --historia historia.md [--conduta conduta.md]
    emr.py vitais   --paciente <uuid> --pa 128/82 --peso 71,4 --altura 1,62
    emr.py exame    --paciente <uuid> --lote lote.json
    emr.py glosa    --codigo PLN3FC5EDA6 --texto "média do açúcar em 3 meses"

O helper HTTP está duplicado do `scripts/plano/plano.py` de propósito: os dois precisam rodar como
arquivo único copiado para dentro do container, e um import quebraria isso.
"""

import argparse
import json
import os
import sys
import urllib.error
import urllib.request
from datetime import datetime, timezone

API = os.environ.get("EMR_API", "http://localhost:3001")
TOKEN = os.environ.get("EMR_TOKEN", "dev")


def chama(metodo, caminho, corpo=None):
    req = urllib.request.Request(
        API.rstrip("/") + caminho,
        method=metodo,
        data=json.dumps(corpo).encode() if corpo is not None else None,
        headers={"Authorization": "Bearer " + TOKEN, "Content-Type": "application/json"},
    )
    try:
        with urllib.request.urlopen(req, timeout=300) as r:
            bruto = r.read()
            return json.loads(bruto) if bruto else None
    except urllib.error.HTTPError as e:
        # O corpo do erro é onde o servidor explica (403 de RBAC, 400 com os campos que falharam na
        # validação). Mostrar só "HTTP 400" transformaria cada recusa em adivinhação.
        sys.exit(f"HTTP {e.code} em {metodo} {caminho}: {e.read().decode('utf-8','replace')[:800]}")


def numero(txt, campo="valor"):
    """Aceita vírgula decimal: quem digita "71,4" está certo, e é assim que o médico escreve."""
    try:
        return float(str(txt).replace(",", "."))
    except ValueError:
        sys.exit(f"{campo}: {txt!r} não é número (use 71,4 ou 71.4)")


def agora():
    return datetime.now(timezone.utc).isoformat()


def data_rfc(txt):
    """'2026-09-01' ou RFC3339 completo. Data solta vira meio-dia UTC, não meia-noite: meia-noite
    em UTC é o dia anterior no Brasil, e a coleta ia parar na véspera."""
    t = str(txt).strip()
    return t if "T" in t else t + "T12:00:00Z"


# ---------------------------------------------------------------------------


def seleciona(patient_id):
    """Aponta o paciente selecionado do USUÁRIO da chave.

    Vários endpoints do EMR (lote de exames, lista de notas clínicas) não recebem o paciente no
    corpo nem na URL: leem `users.selected_patient_id`, porque na tela sempre há um paciente aberto.
    Passar `?patientId=` neles não dá erro, é só ignorado — e a lista volta vazia, ou pior, do
    paciente errado. Selecionar aqui é o que torna cada comando repetível.

    É efeito colateral de verdade: em dev com `DEV_BYPASS_AUTH` a chave é a do MESMO admin que está
    com o EMR aberto no navegador, então selecionar aqui troca o paciente da tela dele. Por isso
    devolve o que estava antes, para o chamador restaurar; e por isso, em prod, use uma chave de
    serviço e não a do médico que está atendendo.
    """
    antes = (chama("GET", "/api/v1/users/me") or {}).get("selectedPatientId")
    chama("PUT", "/api/v1/users/me/selected-patient", {"patientId": patient_id})
    return antes


def restaura(antes):
    """Devolve a seleção de quem estava usando a chave.

    Quando NÃO havia seleção antes, o paciente fica selecionado: `PUT /users/me/selected-patient`
    exige um UUID válido e não existe rota para desselecionar. É o único rastro que estes comandos
    deixam, e ele é visível (o paciente aparece aberto na tela), não silencioso.
    """
    if antes:
        chama("PUT", "/api/v1/users/me/selected-patient", {"patientId": antes})


def ficha(a):
    antes = seleciona(a.paciente)
    p = chama("GET", f"/api/v1/patients/{a.paciente}")
    print(f'{p.get("name")}  ·  {p.get("gender") or "sexo não informado"}  ·  nasc. {p.get("birthDate") or "não informada"}')

    condutas = chama("GET", f"/api/v1/patients/{a.paciente}/care-plan-items") or []
    if isinstance(condutas, dict):
        condutas = condutas.get("data") or condutas.get("items") or []
    print(f"\ncondutas ({len(condutas)}):")
    for c in condutas[:12]:
        print(f'  [{c.get("letterCode")}·{c.get("priority","")}] {c.get("recommendation","")[:90]}')

    # Sem `?patientId=`: este endpoint lê o paciente SELECIONADO (ver `seleciona`).
    notas = chama("GET", "/api/v1/clinical-notes?limit=20") or []
    if isinstance(notas, dict):
        notas = notas.get("data") or notas.get("items") or []
    print(f"\nnotas clínicas ({len(notas)}):")
    for n in notas[:8]:
        print(f'  {n.get("createdAt","")[:10]} · {n.get("status","")} · {(n.get("clinicalHistory") or "")[:70]}')

    vitais = chama("GET", f"/api/v1/patients/{a.paciente}/vitals") or []
    if isinstance(vitais, dict):
        vitais = vitais.get("data") or vitais.get("items") or []
    print(f"\nvitais ({len(vitais)}):")
    for v in vitais[:5]:
        print(f'  {(v.get("measuredAt") or "")[:10]} · PA {v.get("systolicBp")}/{v.get("diastolicBp")} · peso {v.get("weight")}')
    restaura(antes)


def conduta(a):
    corpo = {"letterCode": a.letra, "recommendation": a.recomendacao}
    if a.porque:
        corpo["rationale"] = a.porque
    if a.meta:
        corpo["target"] = a.meta
    if a.prioridade:
        corpo["priority"] = a.prioridade
    if a.exame:
        corpo["labTestCode"] = a.exame
    r = chama("POST", f"/api/v1/patients/{a.paciente}/care-plan-items", corpo)
    print(f'conduta gravada: {r.get("id")} · [{r.get("letterCode")}] {r.get("recommendation","")[:70]}')
    # O deck lê o plano de cuidado, e o dossiê é congelado por plano: quem já tinha um rascunho
    # aberto continua com o dossiê velho até montar de novo.
    print("o rascunho do plano precisa ser remontado para incluir esta conduta")


def nota(a):
    # `ClinicalNoteService.resolvePatient` recusa quando não há paciente selecionado, e recusa de
    # novo quando o selecionado é OUTRO — com 500, não com uma mensagem que ajude. Selecionar antes
    # é o que faz o comando funcionar em chave nova e logo depois de escrever noutro paciente.
    antes = seleciona(a.paciente)
    corpo = {
        "patientId": a.paciente,
        "clinicalHistory": open(a.historia, encoding="utf-8").read().strip(),
        "visibility": a.visibilidade,
        # NUNCA true: assinar é ato médico com certificado, e não sai daqui.
        "sign": False,
    }
    if a.conduta:
        corpo["conduct"] = open(a.conduta, encoding="utf-8").read().strip()
    r = chama("POST", "/api/v1/clinical-notes", corpo)
    restaura(antes)
    print(f'nota criada: {r.get("id")} · status {r.get("status")}')
    print("nasce como rascunho: quem assina é o médico, na tela dele")


def vitais(a):
    corpo = {"measuredAt": data_rfc(a.quando) if a.quando else agora()}
    if a.pa:
        sis, barra, dia = a.pa.partition("/")
        if not barra or not sis.strip().isdigit() or not dia.strip().isdigit():
            sys.exit(f"--pa {a.pa!r}: use sistólica/diastólica, ex.: 128/82")
        corpo["systolicBp"], corpo["diastolicBp"] = int(sis), int(dia)
    for arg, campo in (("fc", "heartRate"), ("spo2", "spo2")):
        if getattr(a, arg) is not None:
            corpo[campo] = int(getattr(a, arg))
    for arg, campo in (("peso", "weight"), ("cintura", "waistCircumference"),
                       ("temperatura", "temperature")):
        if getattr(a, arg) is not None:
            corpo[campo] = numero(getattr(a, arg), "--" + arg)
    if a.altura is not None:
        # A coluna é em CENTÍMETROS (o `BeforeSave` divide por 100 para o IMC), e "1,62" é como se
        # fala altura em português. Digitado em metros, o IMC dava 272.061 e o banco devolvia
        # "numeric field overflow" — erro que não diz nada sobre a unidade. Ninguém tem 3 cm nem
        # 3 metros, então o valor diz sozinho em que unidade veio.
        h = numero(a.altura, "--altura")
        corpo["height"] = h * 100 if h < 3 else h
    if len(corpo) == 1:
        sys.exit("nada para gravar: passe ao menos uma medida")
    r = chama("POST", f"/api/v1/patients/{a.paciente}/vitals", corpo)
    print(f'vitais gravados: {r.get("id")} · PA {r.get("systolicBp")}/{r.get("diastolicBp")} · peso {r.get("weight")}')


def exame(a):
    """Grava um lote de resultados, ligado ao CATÁLOGO.

    O caminho antigo (`POST /lab-results`, com nome de exame em texto livre) grava numa tabela que o
    escore e o dossiê não leem: o resultado entra no banco e não aparece em régua nenhuma. O que o
    dossiê lê é `lab_results` dentro de um lote, com `lab_test_definition_id` preenchido e
    `result_numeric` não nulo. Resolver o código do catálogo aqui é o que garante isso.
    """
    lote = json.load(open(a.lote, encoding="utf-8"))
    resultados = []
    for r in lote["resultados"]:
        d = chama("GET", f'/api/v1/lab-tests/definitions/code/{r["codigo"]}')
        item = {
            "labTestDefinitionId": d["id"],
            "testName": d["name"],
            "testType": d.get("category") or "other",
            "unit": r.get("unidade") or d.get("unit"),
            "matched": True,
            "source": "manual",
        }
        # Número quando é número. Sem `resultNumeric` a régua não desenha e o escore não classifica.
        if "valor" in r:
            item["resultNumeric"] = numero(r["valor"], r["codigo"])
        if "texto" in r:
            item["resultText"] = str(r["texto"])
        if "referencia" in r:
            item["referenceRange"] = r["referencia"]
        resultados.append(item)
        print(f'  {d["name"]}  ←  {r.get("valor", r.get("texto"))} {item["unit"] or ""}')

    antes = seleciona(a.paciente)
    corpo = {
        "laboratoryName": lote.get("laboratorio") or "Informado em consulta",
        "collectionDate": data_rfc(lote["coleta"]),
        "status": lote.get("status") or "completed",
        "labResults": resultados,
    }
    if lote.get("observacoes"):
        corpo["observations"] = lote["observacoes"]
    r = chama("POST", "/api/v1/lab-result-batches", corpo)
    restaura(antes)
    # O lote não leva paciente no corpo nem na URL: ele cai em quem estiver SELECIONADO. Se algo
    # tiver mexido na seleção entre a linha de cima e esta, um lote inteiro de exames aterrissa no
    # prontuário errado, e sem erro nenhum. Conferir o dono na resposta é barato.
    if r.get("patientId") and r["patientId"] != a.paciente:
        sys.exit(f'ABORTA: o lote {r.get("id")} caiu no paciente {r["patientId"]}, não em {a.paciente}')
    print(f'\nlote gravado: {r.get("id")} · {len(resultados)} resultados em {lote["coleta"]}')
    print("rode `recalc-scores` se quiser o escore refeito antes de remontar o plano")


def glosa(a):
    """Cura a glosa do exame: o que ele MEDE, em até cinco palavras, para o paciente ler na régua.

    Único comando do arquivo que exige chave de ADMIN: `PUT /lab-tests/definitions/:id` é
    `RequireAdmin()`, enquanto o resto pede clínico ou staff. Em dev o bypass já injeta um admin;
    em prod, uma chave de serviço de clínico leva 403 aqui e passa em todo o resto.
    """
    d = chama("GET", f"/api/v1/lab-tests/definitions/code/{a.codigo}")
    chama("PUT", f'/api/v1/lab-tests/definitions/{d["id"]}', {**d, "patientGloss": a.texto})
    print(f'{d["name"]}: "{a.texto}"')


# ---------------------------------------------------------------------------


def main():
    p = argparse.ArgumentParser(description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    sub = p.add_subparsers(dest="cmd", required=True)

    s = sub.add_parser("ficha", help="o que já existe no prontuário")
    s.add_argument("--paciente", required=True)
    s.set_defaults(fn=ficha)

    s = sub.add_parser("conduta", help="grava uma recomendação no plano de cuidado")
    s.add_argument("--paciente", required=True)
    s.add_argument("--letra", required=True, choices=list("AGIR"), help="pilar AGIR")
    s.add_argument("--recomendacao", required=True)
    s.add_argument("--porque", help="rationale: por que entra agora")
    s.add_argument("--meta", help="target: o que se espera")
    s.add_argument("--prioridade", choices=["high", "medium", "low"])
    s.add_argument("--exame", help="código do catálogo que esta conduta persegue")
    s.set_defaults(fn=conduta)

    s = sub.add_parser("nota", help="cria a nota clínica em RASCUNHO (nunca assina)")
    s.add_argument("--paciente", required=True)
    s.add_argument("--historia", required=True, help="arquivo com a história clínica")
    s.add_argument("--conduta", help="arquivo com a conduta")
    s.add_argument("--visibilidade", default="medicalOnly",
                   choices=["all", "medicalOnly", "psychOnly", "authorOnly"])
    s.set_defaults(fn=nota)

    s = sub.add_parser("vitais", help="pressão, peso, altura")
    s.add_argument("--paciente", required=True)
    s.add_argument("--pa", help="sistólica/diastólica, ex.: 128/82")
    s.add_argument("--fc", type=int)
    s.add_argument("--spo2", type=int)
    s.add_argument("--peso")
    s.add_argument("--altura", help="em cm (162) ou em metros (1,62)")
    s.add_argument("--cintura")
    s.add_argument("--temperatura")
    s.add_argument("--quando", help="AAAA-MM-DD; default agora")
    s.set_defaults(fn=vitais)

    s = sub.add_parser("exame", help="grava um lote de resultados ligado ao catálogo")
    s.add_argument("--paciente", required=True)
    s.add_argument("--lote", required=True, help='JSON: {"laboratorio","coleta","resultados":[{"codigo","valor","unidade"}]}')
    s.set_defaults(fn=exame)

    s = sub.add_parser("glosa", help="cura a glosa do exame no catálogo")
    s.add_argument("--codigo", required=True)
    s.add_argument("--texto", required=True)
    s.set_defaults(fn=glosa)

    a = p.parse_args()
    a.fn(a)


if __name__ == "__main__":
    main()
