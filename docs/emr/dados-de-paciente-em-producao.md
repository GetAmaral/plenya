# Escrever dado de PACIENTE em produção

> Escrito em 2026-09-06, depois de uma discussão clínica inteira travar no último metro por falta de
> caminho documentado, e depois de eu afirmar com confiança duas coisas erradas sobre como isso
> vinha sendo feito. A auditoria de produção desmentiu as duas. Este documento é o que ficou de pé.

## São DOIS caminhos, por decisão de projeto

| O que | Caminho | Por quê |
|---|---|---|
| **Resultado de exame** e PDF de laudo | **SQL direto**, gerado por [`scripts/emr/exames-sql.py`](../../scripts/emr/exames-sql.py) | O PDF já foi lido e conferido aqui, com olho clínico. Mandá-lo pelo classificador de IA do EMR é pagar de novo pela mesma leitura, e pela pior das duas. |
| **Anamnese, escore, condutas, pedidos de exame, receitas** | **HTTP normal**, por [`scripts/emr/emr.py`](../../scripts/emr/emr.py), em nome do médico | É pouco volume, é decisão clínica, e é onde RBAC, validação de DTO e auditoria de rota valem o preço. |

Nenhum dos dois assina nada. Assinatura de nota e de receita é ato médico com certificado
ICP-Brasil e não sai de script, nunca.

---

## Caminho 1 — resultado de exame, por SQL

### O que já deu errado, e que o gerador conserta

Cinco cargas entraram em produção por `INSERT` cru entre julho e agosto de 2026 — **645 resultados**
— e ficaram com três defeitos, todos silenciosos:

1. **Sem linha de auditoria.** O prontuário tinha o resultado e nenhum registro de quem o pôs lá.
2. **Sem conversão de unidade.** `unit_original` vazio prova que a camada de serviço nem rodou; o
   número ficou na unidade do laudo e o escore comparou contra uma escala noutra grandeza. Foi o
   que gerou o conversor de 4 camadas e o `reconvert-lab-units`.
3. **Sem nível nos qualitativos.** `lab_results.level` nulo tira sorologia, urina qualitativa e
   BI-RADS do escore sem avisar.

O gerador devolve as três. A trava do catálogo é a quarta: código inexistente **aborta a transação
com o nome do código**, em vez de sumir num JOIN.

### A receita, inteira

```bash
# 1. gerar (o formato de lote.json é o MESMO do `emr.py exame`)
scripts/emr/exames-sql.py --paciente <uuid-de-prod> --lote lote.json \
    --autor getfilho@yahoo.com.br > carga.sql

# 2. conferir a olho, e aplicar em transação única
cat carga.sql | ssh plenya "sudo docker exec -i mb511beqjtgd7nsjlnngh3m6 \
    psql -U plenya_user -d plenya_db -v ON_ERROR_STOP=1"

# 3. os TRÊS reparos, nesta ordem. Cada um repõe algo que a via HTTP faria na ingestão.
API=$(ssh plenya "sudo docker ps --format '{{.Names}}' | grep '^kgcuxgvmnbx6'")
ssh plenya "sudo docker exec $API /app/reconvert-lab-units -aplicar"    # unidade
ssh plenya "sudo docker exec $API /app/classify-all"                    # nível
ssh plenya "sudo docker exec $API /app/recalc-scores -paciente <uuid>"  # escore
```

🚨 **`classify-all` é o que mais some da memória.** Medido numa carga de 147 resultados: sem ele o
escore saiu **74,5% com 110 itens**; com ele, **77,4% com 120**, idêntico à mesma carga feita por
HTTP. Dez itens qualitativos fora do escore, sem nenhum aviso.

### Por que `unit_original` é preenchido no INSERT

Porque `reconvert-lab-units` **parte sempre do valor original** (`result_numeric_original` /
`unit_original` quando existem, senão do valor atual) — é isso que o torna idempotente. Gravar a
unidade do laudo ali é o que faz a conversão sair certa depois, e rodar duas vezes não converter
duas vezes.

### Por que a carga pode escrever a própria auditoria

`RevokeAuditLogMutations` revoga **UPDATE, DELETE e TRUNCATE** de `audit_logs` em produção. **INSERT
não é revogado.** Então não há desculpa para prontuário sem trilha: a carga registra a si mesma,
com `resource='lab-result-batches'` e o `user_agent` dizendo que foi script e qual arquivo.

---

## Caminho 2 — anamnese, escore, condutas, pedidos e receitas, por HTTP

Vai pela porta do prontuário, em nome do médico. Só o caminho HTTP tem os hooks do model (UUID v7,
cripto de CPF/RG), a validação do DTO, o RBAC de cada rota e a auditoria de rota.

```bash
export EMR_API=https://api.plenyasaude.com.br
export EMR_TOKEN=$(scripts/emr/prod-token.sh)

scripts/emr/emr.py ficha   --paciente <uuid-de-prod>
scripts/emr/emr.py conduta --paciente <uuid> --letra A --recomendacao "…" --porque "…"
scripts/emr/emr.py nota    --paciente <uuid> --historia historia.md
scripts/plano/plano.py --api "$EMR_API" ler --paciente <uuid> --plano <uuid>
```

### Papel por rota

Não existe API key nem conta de serviço no código: auth é usuário + senha → JWT
(`POST /api/v1/auth/login`; 2FA só entra se estiver ligado naquele usuário).

| Escrita | Rota | Papel mínimo |
|---|---|---|
| Conduta do plano | `POST /patients/:id/care-plan-items` | `RequireClinician` |
| Nota clínica | `POST /clinical-notes` | `RequireAnyStaff` |
| Problemas, medicações, vitais | `POST /patients/:id/{problems,medications,vitals}` | `RequireAnyStaff` |
| Pedido de exames | `POST /lab-requests` | `RequireClinician` |
| **Receita** | `POST /prescriptions` | **`RequireDoctor`** |
| Glosa do catálogo | `PUT /lab-tests/definitions/:id` | `RequireAdmin` |

`doctor` cobre tudo. `nurse` cobre tudo menos receita.

### O token

`scripts/emr/prod-token.sh` loga e imprime só o `accessToken`, para caber em substituição de
comando. Lê de `~/.plenya-vps-secrets/emr-prod-api.env`, **separado de propósito** do
`emr-prod.env`, que guarda chave de criptografia e senha do banco e não deve ser aberto para obter
um JWT.

Hoje o arquivo aponta para a conta do médico (`ADMIN_INITIAL_EMAIL`), e é por isso que a auditoria
das cargas antigas diz "Getulio Amaral Filho · curl/8.5.0". Uma **conta de serviço dedicada** é
melhor, e não por burocracia: a trilha passa a dizer que foi um script, e a assinatura do médico
continua sendo o que transforma rascunho em ato.

---

## Disciplina, nos dois caminhos

Cada item corresponde a um erro que já aconteceu.

1. **Paciente pelo UUID, nunca pelo nome.** Em produção havia DOIS cadastros da mesma pessoa, um
   com erro de digitação, e só um tinha os pedidos de exames. O gerador SQL recusa `--paciente` que
   não seja UUID por isso.
2. **Ensaie na dev, num paciente descartável, com o script exato**, e compare valor a valor casando
   por **código E data da coleta**. Casar só por código faz produto cartesiano em exame com duas
   coletas e esconde a divergência no ruído.
3. **Ao exportar da dev, use o par ORIGINAL quando houve conversão.** Mandar o valor já convertido
   com a unidade do laudo faz o serviço converter de novo. Medido: T3 reverso iria a 2000 ng/dL em
   vez de 20; cálcio iônico 4×, PCR 10×, testosterona livre 10×.
   ```sql
   CASE WHEN unit_conversion_status='convertido' THEN result_numeric_original ELSE result_numeric END,
   CASE WHEN unit_conversion_status='convertido' THEN unit_original          ELSE unit           END
   ```
4. **Nada disto é idempotente.** Rodar duas vezes duplica. O SQL vai em transação única, então
   erro no meio não deixa carga pela metade — mas sucesso duas vezes deixa tudo em dobro.
5. **O lote por HTTP e a nota lêem `users.selected_patient_id`**, não a URL. `emr.py` seleciona
   antes e aborta se o lote cair no paciente errado.
6. **Nunca assine, nunca publique.**

Relacionado: [migrations-decisao.md](migrations-decisao.md) ·
[.claude/workflows/database-ops.md](../../.claude/workflows/database-ops.md) ·
[scripts/emr/exames-sql.py](../../scripts/emr/exames-sql.py) ·
[scripts/emr/emr.py](../../scripts/emr/emr.py)
