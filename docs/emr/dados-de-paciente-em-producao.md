# Escrever dado de PACIENTE em produção

> Escrito em 2026-09-06, depois de uma discussão clínica inteira (uma paciente com carcinoma de
> mama, cirurgia na semana seguinte) travar no último metro: o prontuário estava pronto na dev, a
> ordem de subir foi dada, e **não havia caminho**. Este documento existe para isso não se repetir.

## O mal-entendido que gerou este documento

"Você já ajustou prod várias vezes." Verdade, e é por isso que a confusão é natural. Só que **o que
eu ajustei nunca foi dado de paciente**:

| O que | Como | Onde está documentado |
|---|---|---|
| Catálogo, escore, pilares, níveis, seeds | Bundle SQL por `docker exec -i <DB> psql`, depois migrations goose no deploy | `docs/emr/agir-bioma-PROD-bundle.sql`, [[plenya_dado_catalogo_vai_por_migration]] |
| Schema | Migration goose, aplicada no deploy do `api` (`RUN_MIGRATIONS=true`) | [docs/emr/migrations-decisao.md](migrations-decisao.md) |
| Leitura do banco de prod | `ssh plenya "sudo docker exec mb511beq… psql …"` | memória `plenya_prod_psql_receita` |
| **Dado de paciente em prod** | **não existia** | este documento |

A porta do prontuário (`scripts/emr/emr.py`, commit `9d2a64ef`) foi construída para isso e o
cabeçalho dela diz "prod é outra URL e outro token". **O token nunca foi criado.** A porta existe;
ninguém cortou a chave.

## Por que não pode ser por psql, mesmo sendo mais fácil

A Regra de Ouro 2 do CLAUDE.md manda dado de dev ir por banco direto, e abre exceção para
prontuário. A exceção não é preciosismo: só o caminho HTTP tem as quatro coisas que um registro
clínico precisa ter.

1. **Hooks do model** — UUID v7, criptografia de CPF/RG, `LastReview`.
2. **Validação do DTO** — um `resultNumeric` fora de faixa ou uma prioridade inválida param aqui.
3. **RBAC de rota** — ver a tabela de papéis abaixo.
4. **Linha de auditoria** — `middleware.AuditLog` é middleware de ROTA. Um `INSERT` por psql entra
   no banco sem nenhum registro de quem escreveu, e em produção `RevokeAuditLogMutations` revoga
   UPDATE/DELETE/TRUNCATE de `audit_logs` justamente para que essa linha não possa ser apagada
   depois. **Escrever prontuário por fora do HTTP é apagar a própria pegada.**

O mesmo vale para a conversão de unidade: o serviço converte na ingestão e grava
`unit_original` + `unit_conversion_status`. Por psql o valor entra cru e o escore compara na
unidade errada, em silêncio (é a 4ª causa da memória `emr_conversao_unidade_exames`).

## O que a API exige, por rota

Não existe conceito de API key nem de conta de serviço no código: autenticação é
usuário + senha → JWT (`POST /api/v1/auth/login` devolve `accessToken` e `refreshToken`; 2FA só
entra se `TwoFactorEnabled` estiver ligado naquele usuário).

| Escrita | Rota | Papel mínimo |
|---|---|---|
| Lote de exames | `POST /lab-result-batches` | `RequireClinician` |
| Conduta do plano | `POST /patients/:id/care-plan-items` | `RequireClinician` |
| Nota clínica | `POST /clinical-notes` | `RequireAnyStaff` |
| Problemas, medicações em uso, vitais | `POST /patients/:id/{problems,medications,vitals}` | `RequireAnyStaff` |
| **Receita** | `POST /prescriptions` | **`RequireDoctor`** |
| Glosa do catálogo | `PUT /lab-tests/definitions/:id` | `RequireAdmin` |

Ou seja: **`doctor` cobre tudo**; `nurse`/`nutritionist` cobrem tudo menos receita.

## O caminho: conta de serviço dedicada

Não reutilize a conta do médico. Uma conta própria é melhor por um motivo que não é burocrático:
**a linha de auditoria passa a dizer a verdade.** Se o script escrever como "Getulio A", o
prontuário afirma que o médico digitou aquilo. Se escrever como "Assistente (serviço)", o registro
diz o que aconteceu, e a assinatura do médico continua sendo o que transforma rascunho em ato.

### Provisionar (uma vez)

Criar o usuário é ato de administração, não de prontuário, então vai pela tela ou pela rota de
usuários — **não** por psql, para o próprio usuário nascer com os hooks certos.

1. No EMR de produção, **Configurações → Usuários → novo**, com:
   - nome: `Assistente Plenya (serviço)`
   - e-mail: `servico@plenyasaude.com.br` (caixa real, para recuperação de senha)
   - papel: `doctor` se a conta precisar criar receitas; `nurse` se não precisar
   - 2FA: **desligado** nesta conta (com 2FA ligado não há login não-interativo)
2. Guardar as credenciais em `~/.plenya-vps-secrets/emr-prod-api.env`, no formato:
   ```
   EMR_API=https://api.plenyasaude.com.br
   EMR_USER=servico@plenyasaude.com.br
   EMR_PASSWORD=<senha forte, só desta conta>
   ```
3. Revisar em `/configuracoes` periodicamente: é uma conta com poder de escrita em prontuário.

### Usar

```bash
export EMR_API=https://api.plenyasaude.com.br
export EMR_TOKEN=$(scripts/emr/prod-token.sh)      # login → accessToken, ~expira, refaça quando 401

scripts/emr/emr.py ficha   --paciente <uuid-de-prod>
scripts/emr/emr.py exame   --paciente <uuid-de-prod> --lote lote.json
scripts/plano/plano.py --api "$EMR_API" ler --paciente <uuid> --plano <uuid>
```

`emr.py` e `plano.py` já leem `EMR_API`/`EMR_TOKEN` do ambiente: **não é preciso mudar código**, só
apontar as variáveis.

## Disciplina obrigatória antes de qualquer escrita em prod

Isto não é zelo. Cada item abaixo corresponde a um erro que já aconteceu.

1. **Confira o paciente alvo pelo UUID, nunca pelo nome.** Em prod havia DOIS cadastros para a
   mesma pessoa, um com erro de digitação, e só um tinha os pedidos de exames. Nome não desempata.
2. **Ensaie num paciente descartável, na dev, com o script exato.** Foi assim que se descobriu uma
   **dupla conversão de unidade** que teria gravado T3 reverso como 2000 ng/dL em vez de 20, além de
   cálcio iônico 4×, PCR 10× e testosterona livre 10× errados. A causa: exportar o valor **já
   convertido** rotulado com a unidade **do laudo**, fazendo o serviço converter de novo.
   Ao exportar da dev, use o par ORIGINAL quando houve conversão:
   ```sql
   CASE WHEN unit_conversion_status='convertido' THEN result_numeric_original ELSE result_numeric END,
   CASE WHEN unit_conversion_status='convertido' THEN unit_original          ELSE unit           END
   ```
3. **Compare valor a valor depois**, casando por código **e data da coleta**. Casar só por código
   faz produto cartesiano em exame com duas coletas e esconde a divergência real no meio do ruído.
4. **Nenhum script destes é idempotente.** Rodar duas vezes duplica tudo. Não há `ON CONFLICT` numa
   API REST.
5. **O lote de exames e a nota clínica leem `users.selected_patient_id`**, não a URL. Selecione o
   paciente antes e confira o dono na resposta (`emr.py` já faz e aborta se cair no errado).
6. **Nunca assine.** `sign: false` sempre. Assinatura de nota e de receita é ato médico com
   certificado ICP-Brasil, e não sai de script.

## O que continua sem caminho, de propósito

- **Assinar** nota, receita ou relatório.
- **Publicar** o deck para o portal do paciente.
- **Apagar** dado clínico. O `DELETE` de paciente é *soft* (`Patient.DeletedAt` é
  `gorm.DeletedAt`), mas continua sendo decisão do médico, não do script.

Relacionado: [scripts/emr/emr.py](../../scripts/emr/emr.py) ·
[docs/emr/migrations-decisao.md](migrations-decisao.md) ·
[.claude/workflows/database-ops.md](../../.claude/workflows/database-ops.md)
