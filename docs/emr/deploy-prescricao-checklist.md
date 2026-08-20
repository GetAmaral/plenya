# Deploy da frente de prescrição — checklist

Estado verificado em produção **em 20/08/2026, 04:10**, antes de qualquer coisa.

| | Produção hoje | Depois do deploy |
|---|---|---|
| Migration | **65** | 80 |
| `prescriptions` | **0 linhas** | passa a salvar |
| `medication_definitions` | **0 linhas** | precisa do import CMED |
| Catálogo magistral | não existe | 290 substâncias, 132 fórmulas |

As **zero prescrições** não são coincidência: é o bug que abriu esta frente. A tabela
guardava 12 colunas do modelo antigo, `NOT NULL`, que o código novo não preenchia — todo
INSERT batia 23502. Ninguém nunca conseguiu salvar uma receita. Como consequência prática,
a migration 00067, que apaga essas colunas, **não tem risco de perda de dado**: não há dado.

## 1. Backup — feito

- Dump fresco do banco de produção: `db_plenya_20260820_0405.dump`, 241 MB, gerado na VPS
  antes de qualquer mudança. Conferido com `pg_restore -l`: **118 tabelas com dados**,
  incluindo `patients` e `prescriptions`.
- Espelhado em `/home/user/backups/prod/vps/` (o local canônico), junto com os uploads do
  dia (`uploads_20260820_0330.tar.gz`).
- O cron da VPS continua rodando: banco às 03:00, uploads às 03:30.

## 2. O que o deploy faz sozinho

`scripts/deploy/deploy-app.sh api` dispara o build no Coolify. O container de produção tem
`RUN_MIGRATIONS=true`, então **as migrations 66 a 80 aplicam no start**, incluindo a 00066
(catálogo CMED), que ficou de um trabalho anterior e nunca foi para produção.

Ordem obrigatória: **api primeiro, web depois**. Front novo mandando `formulas` para backend
velho cria receita vazia.

## 3. O que o deploy NÃO faz — e sem isso a tela nasce vazia

As migrations criam as TABELAS. O CONTEÚDO mora em seeds que nada executa:

```bash
# catálogo magistral: 290 substâncias, 132 fórmulas, 54 regras, tetos da IN 28
docs/emr/magistral-carga-prod.sh          # 27 arquivos, na ordem de dependência

# catálogo de industrializados (ANVISA/CMED): ~26 mil apresentações
docker exec <api> go run ./cmd/import-cmed --file <xls_conformidade_site_AAAAMM.xlsx>
```

A ordem do primeiro foi validada carregando **do zero num banco vazio** e comparando com o
dev até bater exatamente: 290 substâncias, 132 fórmulas, 653 componentes, 54 regras, 167
faixas, 161 tetos, 12 pares, 8 regras de base. Esse teste achou 9 substâncias e 10 fórmulas
que existiam só no dev, criadas por comandos avulsos e nunca capturadas em arquivo — estão
em `magistral-avulsos-capturados.sql`.

## 4. Antes de disparar

- [ ] RAM livre na VPS acima de 1.500 MB (o script confere sozinho e aborta se não houver).
      Medido agora: **6.015 MB**.
- [ ] Não deployar às 04:00: a VPS reinicia sozinha para atualização de kernel a cada ~10
      dias nesse horário (29/07, 08/08, 20/08), fora por ~2 minutos, e tudo volta.
- [ ] Canário: emitir uma receita real **em rascunho**, conferir o PDF, e só então assinar.
      Rascunho não publica documento nem dispara WhatsApp.

## 5. Rollback

Reverter a 00067 recria as colunas legadas vazias; como não há prescrição em produção, é
seguro. As migrations do magistral (68 a 80) têm `Down` simétrico e testado (`up`/`down`/`up`).
Depois que a primeira receita magistral for assinada, corrige-se para a frente, não para trás.
