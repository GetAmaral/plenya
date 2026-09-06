---
name: prontuario-prod
description: Leva um prontuário construído na conversa (exames de PDF, problemas, condutas, nota, receita) para o EMR de PRODUÇÃO, pelos dois caminhos certos - exame por SQL gerado, o resto por HTTP em nome do médico - com os reparos obrigatórios e a conferência contra a dev. Invocar quando o usuário disser "sobe pra prod", "manda pro EMR", "põe no prontuário dele", "migra pra produção" com dado de paciente. NÃO usar para catálogo/escore (isso é migration goose) nem para deploy de código (isso é scripts/deploy/deploy-app.sh).
---

# Skill `/prontuario-prod` — levar prontuário para produção sem erro bobo

> Documento de referência: [docs/emr/dados-de-paciente-em-producao.md](../../../docs/emr/dados-de-paciente-em-producao.md).
> Esta skill é o roteiro executável daquilo. Cada passo aqui existe porque a ausência dele já
> causou um erro real, e o erro está nomeado.

## O que esta skill NÃO faz

- **Não assina** nota nem receita. Ato médico com certificado ICP-Brasil.
- **Não publica** deck no portal.
- **Não apaga** dado clínico.
- **Não toca em catálogo, escore ou schema.** Isso é migration goose, aplicada no deploy do `api`.

## Fase 0 — o alvo, antes de qualquer escrita

**Pelo UUID, nunca pelo nome.** Já houve dois cadastros da mesma paciente em produção, um com erro
de digitação, e só um tinha os pedidos de exames.

```bash
export EMR_API=https://api.plenyasaude.com.br
export EMR_TOKEN=$(scripts/emr/prod-token.sh)

# quem existe com esse nome, INCLUINDO os removidos (deleted_at) — senão você "descobre"
# duplicatas que já foram tratadas e reporta passivo que não existe
ssh plenya "sudo docker exec mb511beqjtgd7nsjlnngh3m6 psql -U plenya_user -d plenya_db -c \"
SELECT id, name, birth_date, deleted_at,
 (SELECT count(*) FROM lab_requests WHERE patient_id=p.id AND deleted_at IS NULL) AS pedidos,
 (SELECT count(*) FROM lab_result_batches WHERE patient_id=p.id AND deleted_at IS NULL) AS lotes
FROM patients p WHERE p.name ILIKE '%SOBRENOME%' ORDER BY p.created_at;\""
```

Regras de desempate: **quem tem os pedidos de exames é o registro bom** (foi por ele que o médico
pediu o painel). Confira a **data de nascimento contra os laudos** — o cadastro de recepção erra, e
três laudos independentes valem mais que uma digitação.

## Fase 1 — os exames, por SQL

Vai por SQL **de propósito**: o PDF já foi lido e conferido na conversa, e remandá-lo pelo
classificador de IA do EMR é pagar de novo pela leitura pior.

### 1.1 Montar o `lote.json`

Mesmo formato do `emr.py exame` (objeto ou lista):
```json
{"laboratorio":"…","coleta":"2026-07-23","observacoes":"…",
 "resultados":[{"codigo":"PLN…","valor":"4,23","unidade":"M/µL","texto":"…"}]}
```

🚨 **Se estiver exportando de um paciente da dev, use o par ORIGINAL onde houve conversão.** Mandar
o valor já convertido com a unidade do laudo faz o serviço converter de novo. Medido: T3 reverso
iria a **2000 ng/dL** em vez de 20; cálcio iônico 4×, PCR 10×, testosterona livre 10×.
```sql
CASE WHEN unit_conversion_status='convertido' THEN result_numeric_original ELSE result_numeric END,
CASE WHEN unit_conversion_status='convertido' THEN unit_original          ELSE unit           END
```

### 1.2 Ensaiar na dev, num paciente descartável

Não pule. Foi o ensaio que pegou a dupla conversão.
```bash
T=$(curl -s -X POST http://localhost:3001/api/v1/patients -H "Authorization: Bearer dev" \
     -H "Content-Type: application/json" -d '{"name":"Ensaio Migracao","birthDate":"…","gender":"…"}' \
     | python3 -c "import sys,json;print(json.load(sys.stdin)['id'])")
scripts/emr/exames-sql.py --paciente "$T" --lote lote.json --autor doctor@plenya.com > ensaio.sql
docker compose exec -T db psql -U plenya_user -d plenya_db -v ON_ERROR_STOP=1 -q < ensaio.sql
```
Compare **valor a valor, casando por código E data** — só por código faz produto cartesiano em
exame com duas coletas e esconde a divergência no ruído. Depois apague o paciente de ensaio.

### 1.3 Aplicar em produção

```bash
scripts/emr/exames-sql.py --paciente <uuid-prod> --lote lote.json \
    --autor getfilho@yahoo.com.br > carga.sql
cat carga.sql | ssh plenya "sudo docker exec -i mb511beqjtgd7nsjlnngh3m6 \
    psql -U plenya_user -d plenya_db -v ON_ERROR_STOP=1 -q"
```
Transação única: erro no meio não deixa carga pela metade. **Mas sucesso duas vezes duplica** —
nada aqui é idempotente.

### 1.4 Os DOIS reparos obrigatórios

```bash
A=$(ssh plenya "sudo docker ps --format '{{.Names}}' | grep '^kgcuxgvmnbx6'")
ssh plenya "sudo docker exec $A /app/reconvert-lab-units -aplicar"   # unidade
```

E a **classificação, que vai pela ROTA e não pelo binário**: `classify-all` NÃO está na imagem de
produção (o Dockerfile só compila server, migrate, backfill-patient-phones, import-cmed,
reconvert-lab-units e recalc-scores). A rota é melhor de qualquer forma, porque é auditada:

```bash
# selecionar o paciente ANTES, senão 403: a rota confere users.selected_patient_id
curl -s -X PUT "$EMR_API/api/v1/users/me/selected-patient" -H "Authorization: Bearer $EMR_TOKEN" \
     -H "Content-Type: application/json" -d '{"patientId":"<uuid-prod>"}'
for L in $(ssh plenya "sudo docker exec mb511beqjtgd7nsjlnngh3m6 psql -U plenya_user -d plenya_db -t -A -c \
    \"SELECT id FROM lab_result_batches WHERE patient_id='<uuid-prod>' AND deleted_at IS NULL;\""); do
  curl -s -o /dev/null -w "$L %{http_code}\n" -X POST "$EMR_API/api/v1/lab-result-batches/$L/classify" \
       -H "Authorization: Bearer $EMR_TOKEN" -H "Content-Type: application/json" -d '{}'
done
```

🚨 **Sem a classificação, TODO resultado qualitativo fica fora do escore em silêncio** — sorologias,
urina qualitativa, BI-RADS. Medido em 147 resultados: **74,5% com 110 itens** sem ela, **77,4% com
120** com ela.

## Fase 2 — o resto, por HTTP em nome do médico

Problemas, medicações em uso, condutas, nota e receita. `doctor` cobre tudo; `nurse` cobre tudo
menos receita (`POST /prescriptions` é `RequireDoctor`).

```bash
scripts/emr/emr.py conduta --paciente <uuid-prod> --letra A --recomendacao "…" --porque "…"
scripts/emr/emr.py nota    --paciente <uuid-prod> --historia historia.md
```

Campos do paciente que destravam itens de escore inteiros e são fáceis de esquecer:
`menopause`, `gender`, `birthDate`. **`menopause` nulo derruba os três itens de ferritina**; ancorar
em FSH/estradiol/progesterona do próprio painel, nunca chutar.

## Fase 3 — escore, e a conferência que fecha

```bash
curl -s -X POST "$EMR_API/api/v1/patients/<uuid-prod>/score-snapshots" \
     -H "Authorization: Bearer $EMR_TOKEN" -H "Content-Type: application/json" -d '{}'
```

**Compare prod contra dev.** Contagens iguais e escore igual é o critério de pronto; escore
diferente com contagem igual significa valor diferente, e aí volte à Fase 1.

```
exames · problemas · condutas · medicações · escore · itens avaliados
```

## Fase 4 — o deck, se houver

Montar em prod dá um deck DIFERENTE do da dev, porque prod tem o pedido de exames e ganha o slide
"o que ainda falta". Se o deck já foi curado na conversa, **porte o conteúdo** em vez de reescrever:
`GET` o plano da dev, remova os `id` de slide, e `PUT` no plano de prod com o `expectedRevision`
dele. Depois **confira o overflow**, que muda com o slide novo.

## Erros que já aconteceram, em uma lista

1. **Afirmar como isso foi feito antes sem olhar a auditoria.** Eu disse "nunca escrevi paciente em
   prod" e a `audit_logs` mostrou `curl/8.5.0` em meu próprio rastro.
2. **Consultar `patients` sem filtrar `deleted_at`** e reportar duplicata que já estava removida.
3. **Exportar valor convertido com unidade do laudo** (dupla conversão).
4. **Esquecer a classificação** e perder os qualitativos do escore.
5. **Esquecer de selecionar o paciente** antes de rotas que leem `selected_patient_id` (403).
6. **`NULL` sem cast** em lote só de resultado textual: o VALUES infere `text` e o INSERT quebra.
7. **`set -euo pipefail` + `grep` sem match** derrubando script de token em silêncio.
