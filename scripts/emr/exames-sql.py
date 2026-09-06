#!/usr/bin/env python3
"""Gera o SQL que sobe RESULTADO DE EXAME para produção, direto na tabela.

## Por que existe um caminho por SQL, se a Regra de Ouro 2 manda prontuário por HTTP

Porque resultado de exame é o único dado de prontuário que chega em VOLUME e já interpretado. O
caminho normal do EMR passa o PDF por um classificador de IA para extrair os valores; quando o PDF
já foi lido e conferido aqui, mandá-lo de novo pelo classificador é pagar duas vezes pela mesma
leitura, e é a leitura pior das duas, porque esta teve olho clínico em cima.

Então este caminho é decisão de projeto, não atalho. O que ele NÃO pode fazer é o que aconteceu
antes: 645 resultados entraram em produção por `INSERT` cru, em cinco cargas, e ficaram
**sem uma linha de quem os escreveu** e **sem conversão de unidade** — o número gravado na unidade
do laudo e depois comparado contra uma escala noutra grandeza. Foi o que gerou o conversor de 4
camadas e o `reconvert-lab-units`.

Este gerador devolve as três coisas que faltavam:

1. **Trava no catálogo.** Código inexistente aborta a transação com o nome do código, em vez de
   sumir num JOIN silencioso.
2. **`unit_original` preenchido** com a unidade do laudo. É o que o `reconvert-lab-units` usa para
   converter depois, e ele é idempotente porque parte sempre do original.
3. **Linha de auditoria própria.** `RevokeAuditLogMutations` revoga UPDATE/DELETE/TRUNCATE de
   `audit_logs`, não INSERT — então a carga registra a si mesma, dizendo que foi script e qual
   arquivo. A trilha passa a ser verdadeira em vez de ausente.

## Uso

    scripts/emr/exames-sql.py --paciente <uuid-de-prod> --lote lote.json \\
        --autor getfilho@yahoo.com.br > carga.sql

    # conferir o SQL a olho, e só então:
    cat carga.sql | ssh plenya "sudo docker exec -i mb511beqjtgd7nsjlnngh3m6 \\
        psql -U plenya_user -d plenya_db -v ON_ERROR_STOP=1"

    # SEMPRE os TRÊS, nesta ordem. Nenhum é opcional, e cada um repõe uma coisa que a via HTTP
    # faria sozinha na ingestão:
    ssh plenya "sudo docker exec <api> /app/reconvert-lab-units -aplicar"   # unidade
    POST /api/v1/lab-result-batches/<lote>/classify                         # nível (ver abaixo)
    ssh plenya "sudo docker exec <api> /app/recalc-scores -paciente <uuid>" # escore

A CLASSIFICAÇÃO é o que mais some da memória e o que mais custa. Em produção ela vai pela rota
`POST /lab-result-batches/:id/classify` (um POST por lote, e é preciso selecionar o paciente antes,
senão 403): o binário `classify-all` não está na imagem. Ela atribui `lab_results.level`, e
sem ele **todo resultado QUALITATIVO fica de fora do escore em silêncio** — sorologias, urina
qualitativa, BI-RADS. Medido: numa carga de 147 resultados, sem `classify-all` o escore saiu 74,5%
com 110 itens avaliados; com ele, 77,4% com 120, idêntico à mesma carga feita por HTTP. Dez itens
sumidos e nenhum aviso.

O formato de `lote.json` é o MESMO de `emr.py exame` (um objeto, ou uma lista deles):

    {"laboratorio": "...", "coleta": "2026-07-23", "observacoes": "...",
     "resultados": [{"codigo": "PLN...", "valor": "4,23", "unidade": "M/µL", "texto": "..."}]}

Anamnese, escore, condutas, pedidos e receitas NÃO vêm por aqui: vão pela via normal
(`emr.py`), em nome do médico. Ver docs/emr/dados-de-paciente-em-producao.md.
"""
import argparse
import json
import sys
import uuid


def lit(s):
    """Literal SQL com aspas simples escapadas. Nada aqui vem de entrada não confiável, mas o
    texto dos laudos tem apóstrofo com frequência ('Hb A1c', nomes de método)."""
    if s is None:
        return 'NULL::text'
    return "'" + str(s).replace("'", "''") + "'"


def num(v):
    """Número no formato do laudo brasileiro (vírgula decimal) para literal SQL."""
    # O cast é obrigatório mesmo no NULL: num lote só de resultados textuais, todas as linhas do
    # VALUES teriam NULL sem tipo e o Postgres infere `text`, quebrando o INSERT em result_numeric.
    if v is None or str(v).strip() == '':
        return 'NULL::numeric'
    return "'%s'::numeric" % str(v).replace(',', '.').strip()


def bloco(pid, lote, autor, origem):
    lab = lote.get('laboratorio') or 'Informado em consulta'
    dia = lote['coleta']
    obs = lote.get('observacoes')
    itens = lote['resultados']
    codigos = sorted({r['codigo'] for r in itens})

    valores = ',\n      '.join(
        '(%s, %s, %s, %s)' % (lit(r['codigo']), num(r.get('valor')),
                              lit(r.get('unidade')), lit(r.get('texto')))
        for r in itens)
    lista_codigos = ', '.join('(%s)' % lit(c) for c in codigos)

    return f"""
-- ---------------------------------------------------------------------------
-- {dia} · {lab} · {len(itens)} resultados
DO $carga$
DECLARE
  v_lote  uuid;
  v_falta text;
  v_n     int;
  v_autor uuid;
BEGIN
  -- 1. Trava: código fora do catálogo aborta com nome, em vez de sumir no JOIN.
  SELECT string_agg(t.c, ', ') INTO v_falta
    FROM (VALUES {lista_codigos}) AS t(c)
   WHERE NOT EXISTS (SELECT 1 FROM lab_test_definitions d
                      WHERE d.code = t.c AND d.deleted_at IS NULL);
  IF v_falta IS NOT NULL THEN
    RAISE EXCEPTION 'códigos ausentes no catálogo: %', v_falta;
  END IF;

  SELECT id INTO v_autor FROM users WHERE email = {lit(autor)} AND deleted_at IS NULL;
  IF v_autor IS NULL THEN
    RAISE EXCEPTION 'autor % não encontrado em users', {lit(autor)};
  END IF;

  INSERT INTO lab_result_batches
    (id, patient_id, laboratory_name, collection_date, status, observations, created_at, updated_at)
  VALUES
    (uuid_generate_v7(), {lit(pid)}, {lit(lab)}, {lit(dia)}::date, 'completed', {lit(obs)}, now(), now())
  RETURNING id INTO v_lote;

  -- 2. unit_original = a unidade DO LAUDO. É dela que o reconvert-lab-units parte, e é por isso
  --    que ele é idempotente. Sem isto o valor fica cru e o escore compara noutra grandeza.
  INSERT INTO lab_results
    (id, lab_result_batch_id, lab_test_definition_id, test_name, test_type,
     result_numeric, result_text, unit, unit_original, matched, source, created_at, updated_at)
  SELECT uuid_generate_v7(), v_lote, d.id, d.name, coalesce(d.category, 'other'),
         i.valor, i.texto, coalesce(i.unidade, d.unit), coalesce(i.unidade, d.unit),
         true, 'manual', now(), now()
    FROM (VALUES
      {valores}
    ) AS i(codigo, valor, unidade, texto)
    JOIN lab_test_definitions d ON d.code = i.codigo AND d.deleted_at IS NULL;
  GET DIAGNOSTICS v_n = ROW_COUNT;

  IF v_n <> {len(itens)} THEN
    RAISE EXCEPTION 'esperava {len(itens)} resultados, inseriu %', v_n;
  END IF;

  -- 3. A carga registra a si mesma. INSERT em audit_logs é permitido (só UPDATE/DELETE/TRUNCATE
  --    são revogados em produção), então não há desculpa para prontuário sem trilha.
  INSERT INTO audit_logs
    (id, user_id, action, resource, resource_id, ip_address, user_agent, success, created_at)
  VALUES
    (uuid_generate_v7(), v_autor, 'create', 'lab-result-batches', v_lote, '127.0.0.1',
     {lit('carga-sql/exames-sql.py · %s · %d resultados' % (origem, len(itens)))}, true, now());

  RAISE NOTICE 'lote % · % resultados · %', v_lote, v_n, {lit(dia)};
END
$carga$;
"""


def main():
    ap = argparse.ArgumentParser(description='SQL de carga de resultados de exame para produção.')
    ap.add_argument('--paciente', required=True, help='UUID do paciente EM PRODUÇÃO')
    ap.add_argument('--lote', required=True, help='lote.json (objeto ou lista), formato do emr.py')
    ap.add_argument('--autor', required=True, help='e-mail do usuário a quem a auditoria atribui a carga')
    a = ap.parse_args()

    try:
        uuid.UUID(a.paciente)
    except ValueError:
        sys.exit('--paciente precisa ser UUID; conferir pelo NOME já colocou dado no paciente errado')

    dados = json.load(open(a.lote, encoding='utf-8'))
    lotes = dados if isinstance(dados, list) else [dados]
    total = sum(len(l['resultados']) for l in lotes)

    print('-- Carga de resultados de exame — gerada por scripts/emr/exames-sql.py')
    print(f'-- paciente {a.paciente} · {len(lotes)} lote(s) · {total} resultados · origem {a.lote}')
    print('-- Rodar com -v ON_ERROR_STOP=1. Tudo numa transação: ou entra inteiro, ou nada.')
    print('-- DEPOIS, obrigatoriamente: reconvert-lab-units -aplicar  e  recalc-scores.')
    print('BEGIN;')
    for lote in lotes:
        print(bloco(a.paciente, lote, a.autor, a.lote))
    print('COMMIT;')


if __name__ == '__main__':
    main()
