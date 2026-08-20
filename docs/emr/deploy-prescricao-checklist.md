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

## 3. O conteúdo do catálogo vai junto, na migration 00081

O catálogo magistral **não é passo manual**: virou a migration `00081_dados_catalogo_magistral`,
que entra com as outras no start do container. Deploy que depende de alguém lembrar de rodar um
script é deploy que uma hora nasce com a tela vazia.

Ela carrega 290 substâncias, 132 fórmulas-base, 653 componentes, 54 regras de dose, 167 faixas,
os 161 tetos do Anexo IV da IN 28, 12 pares de incompatibilidade e 8 regras de base. Todo bloco é
idempotente. O `Down` **não apaga nada** de propósito: depois que o médico conferir uma fórmula ou
criar a dele, não dá para distinguir o que veio da carga do que é trabalho dele. Para desfazer,
reverte-se até a 00069, que derruba as tabelas — aí a intenção de perder o conteúdo está explícita.

### O catálogo de industrializados também vai junto, na 00082

As 26.001 apresentações da lista ANVISA/CMED (edição 202608) entram como base, com a
classificação deduzida (19.364 simples, 4.367 C1, 1.421 antimicrobianos, 727 tarja preta, 94
GLP-1, 28 C5) e as 5.872 marcadas `needs_review` — que é o sistema dizendo onde não sabe.

Vai junto o `curated_at` das **147 linhas conferidas à mão**: é esse campo que impede o reimport
mensal de sobrescrever a curadoria. Não vai o `curated_by`, que aponta para um usuário do banco
de desenvolvimento inexistente em produção e tem chave estrangeira — perde-se a autoria, não a
curadoria.

A **atualização mensal** continua sendo `cmd/import-cmed`, rodada quando houver vontade: o upsert
casa por `ggrem` e só sobrescreve campo de fonte, preservando o que foi curado.

```bash
docker exec <api> go run ./cmd/import-cmed --file <xls_conformidade_site_AAAAMM.xlsx>
```

## 3.1 Ensaio do deploy, feito

O dump de produção de hoje foi restaurado num banco local e as migrations 66 a 81 rodaram em cima
dele — literalmente o que o container vai fazer:

| | |
|---|---|
| Migration final | 82 |
| Catálogo magistral | 290 substâncias / 132 fórmulas / 653 componentes / 54 regras / 167 faixas |
| Catálogo CMED | 26.001 apresentações, 147 com curadoria preservada |
| Norma e regras | 161 tetos da IN 28, 12 pares, 8 regras de base |
| Dado de produção | 27 pacientes, 1.282 itens de escore, 1.906 resultados, 28 usuários — **intactos** |

O ensaio foi refeito do zero depois de uma correção de idempotência (abaixo) e deu os mesmos
números. Três passadas seguidas da carga no mesmo banco também devolvem sempre 290/132/653/54.

### O defeito que o ensaio achou

Um seed inseria a substância como `Aakg` e um arquivo seguinte renomeava para `AAKG`. Numa segunda
passada o insert recriava o nome antigo e a renomeação colidia com a linha já existente — carga
que só funciona uma vez. Os seeds passaram a inserir **já com o nome canônico**: carga
declarativa, sem etapa de migração de nome no meio.

Antes disso, uma carga limpa num banco vazio achou 9 substâncias e 10 fórmulas que existiam só no
dev, criadas por comandos avulsos e nunca capturadas em arquivo. Estão em
`magistral-avulsos-capturados.sql` e entraram na migration.

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
