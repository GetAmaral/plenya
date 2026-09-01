# Planos de paciente no EMR — documento vivo

**Atualizado:** 2026-08-31 · **Fase 0 concluída.**
Plano completo aprovado: deck do paciente registrado no EMR, com saída em tela, 16:9 e A4.
Dados clínicos ficam em `pacs/`, que é gitignored. Este arquivo não repete valor de exame.

---

## 1 · Por que

O plano de devolutiva vivia **fora** do EMR: `pacs/<NOME>/deck/build.py` gerava `deck.html` +
`deck.css`, e dois scripts Playwright produziam os PDFs. Três pacientes depois (João, Ana, José
Ricardo) o custo apareceu: os scripts de render tinham o nome de saída hardcoded e **gravaram por
cima do deck de outro paciente duas vezes**; o `reguas.json` era montado à mão a partir do banco; o
verificador de estouro morava em `/tmp`; e nada disso ficava no prontuário nem chegava ao portal.

## 2 · Decisões fechadas

1. **O relatório AGIR vira um modo do deck.** Um só gerador de conteúdo, três renders.
2. **O deck não é assinado com ICP-Brasil.** O modo relatório A4 continua assinado; o deck é peça de
   comunicação. Receita, pedido e relatório seguem no pipeline `signPDF` intocados.
3. **"Em tela" = página no portal em HTML nativo**, responsiva no celular. PDF é para imprimir e
   para apresentar.
4. **Core doc entre as fases**: este arquivo. Cada fase termina aqui e para em checkpoint.
5. Nenhum commit e nenhum deploy sem ordem explícita.

## 3 · Arquitetura alvo

```
patient_plans (JSONB content)          ← fonte única do CONTEÚDO do plano
        ├─ modo deck      → RenderHTML(1920×1080)   → PDF 16:9   (não assinado)
        ├─ modo deck      → RenderHTML(A4 paisagem) → PDF A4     (não assinado)
        ├─ modo relatório → renderDocument (A4)     → PDF assinado ICP
        └─ modo tela      → /patient-portal/plano   → HTML nativo, responsivo
```

`renderDocument` (`pdfdoc/document.go`) é paginador de fluxo A4 e não serve para slide fixo. Quem
serve é `pdfdoc.RenderHTML` (`pdfdoc/render.go`), que já é entrada pública para papel arbitrário com
Chromium singleton, serializado e com timeout de 35s. `care_plan_items` continua sendo a fonte
estruturada das condutas.

---

## 4 · Fase 0 — o dossiê do plano (CONCLUÍDA)

Deriva do prontuário o insumo que antes era montado à mão. Endpoint:

```
GET /api/v1/patients/:id/plan-dossier      (RequireClinician)
```

| Arquivo | O quê |
|---|---|
| `apps/api/internal/dto/patient_plan_dossier.go` | Formato de saída. `PlanRuler` espelha 1:1 o `reguas.json` legado, de propósito: o `build.py` consome sem adaptação. |
| `apps/api/internal/services/patient_plan_dossier_service.go` | A derivação. Só lê. |
| `apps/api/internal/handlers/patient_plan_handler.go` | Handler fino. |
| `apps/api/internal/services/patient_plan_dossier_service_test.go` | 15 testes de unidade, com números tirados do gabarito real. |
| `apps/api/cmd/server/main.go` | Wiring + rota. |

### O que sai

- **`rulers`** — uma régua por exame que o paciente fez e que tem escala no catálogo aplicável a ele.
- **`strong` / `moving`** — achados classificados e **ordenados por pontos perdidos**.
- **`carePlan`**, **`labRequest`**, **`prescriptions`**, **`snapshot`** — o resto do insumo.

### Reuso obrigatório (e por quê)

A seleção de qual `score_item` manda para um código usa **o mesmo caminho do classificador de
resultados** (`lab_result_batch_service.go`): `lab_test_code` → `ScoreItem.AppliesToPatient` →
`pickScoringItem`. Isso não é preferência de estilo: `pickScoringItem` é a função que escolheu o item
que gravou `lab_results.level`. Escolher outra variante aqui faria a régua desenhar uma escala que
**contradiz o nível gravado no próprio resultado que ela plota**. A primeira versão tinha um
`pickMostSpecificItem` próprio; foi removida.

### A regra do eixo (número derivado, registrado aqui)

O eixo da régua não estava documentado em lugar nenhum — foi recuperado por engenharia reversa dos
`reguas.json` de Ricardo e Ana, e agora vive em `rulerAxis()`:

```
span = última_fronteira − primeira_fronteira
base = [primeira − 0,22·span , última + 0,22·span]
eixo = [ max(0, min(base_lo, min(histórico) × 0,96)) ,
              max(base_hi, max(histórico) × 1,04) ]
```

O padding de 22% dá largura desenhável ao segmento aberto das pontas ("≤15", ">300"). A folga de
0,96/1,04 estica o eixo quando o paciente tem valor fora da escala, para o ponto dele não colar na
borda. O piso nunca fica negativo.

**Aderência ao gabarito: 67/68 no Ricardo.** A única exceção (Alfa-2 Globulina) é ajuste manual, e
na Ana há mais alguns — todos casos de **outlier clampado à mão** (PCR 63,1 num eixo que termina em
15,7; HOMA-IR 13,5 num eixo que termina em 5). Isso é decisão editorial legítima: um valor extremo
esmaga a escala inteira. **Portanto o modelo da Fase 1 precisa de override de eixo por régua.**

### Validação (o teste de aceitação do plano)

Contra os `reguas.json` reais, usando o catálogo do dev (o catálogo é o mesmo de prod; só o dado do
paciente difere):

| | Ricardo (masc, 64) | Ana (fem, 60, pós-menopausa) |
|---|---|---|
| item aplicável escolhido | **68/68** | **36/36** |
| `name` · `points` · `edges` | **68/68** | — |
| `segments` | **68/68** | — |
| `axis` | **67/68** | ajustes manuais (ver acima) |
| `unit` | 67/68 | — |

Rodando a Ana com `menopausa=False` o resultado cai para 35/36, e a diferença é exatamente a
variante pré-menopausa da ferritina: o filtro por sexo/menopausa está agindo de verdade.

Ponta a ponta no dev (João da Silva, 1284 resultados em 21 lotes): 200, **111 réguas**, 23 achados
em "está bem", 9 em "se movendo", ordenados por peso.

### Rotina de qualidade da Fase 0 — o que ela pegou

Build, `go vet`, `gofmt`, suíte completa (8 pacotes verdes) e casos-limite do endpoint. `-race` não
roda no container (não tem gcc); o serviço é leitura pura, sem concorrência.

Os casos-limite acharam **três defeitos reais**, todos corrigidos:

1. **Paciente sem escore calculado dava 500 no dossiê.**
   `ScoreSnapshotRepository.GetLatestByPatientID` engolia o `gorm.ErrRecordNotFound` e devolvia um
   `errors.New("no snapshots found for patient")` solto, que não casa com nenhum `errors.Is`.
   Paciente sem escore é o estado normal de todo paciente novo, não uma falha.
   **Correção na raiz:** sentinela `repository.ErrNoSnapshots`, envelopando também o
   `gorm.ErrRecordNotFound` (`fmt.Errorf("%w: %w", …)`) para os dois testes de erro funcionarem.

2. **O mesmo bug estava latente em produção no relatório AGIR.**
   `care_plan_report_service.go:54` testava `errors.Is(err, gorm.ErrRecordNotFound)` com a intenção
   de responder 400 "sem escore" — e nunca casava, então o paciente sem escore recebia 500
   "falha ao gerar relatório". A correção do item 1 resolve sem tocar nesse arquivo. Verificado:
   agora responde **400**.

3. **`strong` e `moving` saíam como `null`** quando o paciente não tinha achado. Slice nil vira
   `null` em JSON e quebra o `.map()` do front. Agora nascem `[]`, com teste de regressão.

E a própria correção do item 1 **quebrou um terceiro consumidor**, pego antes de sair:
`score_snapshot_handler.go:141` comparava o erro **por string**
(`err.Error() == "no snapshots found for patient"`), e a mensagem mudou. Passou a usar
`errors.Is(err, repository.ErrNoSnapshots)`. Não sobrou nenhuma comparação de erro por string.

Estado dos quatro caminhos, conferido no dev:

| Endpoint | Paciente sem escore | Com escore |
|---|---|---|
| `GET /score-snapshots/latest` | 404 | 200 |
| `POST /care-plan-report` | 400 (era 500) | — |
| `GET /plan-dossier` | 200, listas vazias (era 500) | 200 |

### Revisão de código — mais quatro achados, todos tratados

**A · O dia da coleta repartia uma coleta só em dois pontos da régua.**
A coluna `timestamptz` guarda **duas coisas diferentes**: um dia-calendário digitado, gravado como
meia-noite UTC (`2024-11-05 00:00Z`, que em São Paulo é `2024-11-04 21:00`), e um instante real de
coleta ou importação (`2026-02-06 23:12:29,9-03`). Em `lab_result_batches` são 8 de uma forma e 14
da outra; em `lab_results.collection_date`, 258 de 258 são instantes reais. **Qualquer fuso fixo
erra uma das duas**, e o dia é a chave de deduplicação do histórico: a mesma coleta virava dois
pontos, com um dia de distância. No João da Silva, **1155 dos 1413 resultados** dependiam do
fallback do lote, e a saída mostrava 11 dias de coleta onde existem 7.
`collectionDay()` distingue pelo único sinal confiável — só dia-calendário cai exatamente em
00:00:00 UTC. Resultado: **11 dias → 7 dias**, com fevereiro mantendo o dia local correto.

**B · Os pontos perdidos podiam vir da variante errada.**
`lostBySnapshot` era chaveado pelo código do exame, mas o **IGF-1 (`PLNB153CC32`)** tem um item
guarda-chuva sem recorte etário mais quatro variantes por faixa, e o motor avalia **todas**. A
última do laço sobrescrevia as outras, então os pontos podiam vir de uma variante diferente da que
desenhou a régua — e como eles são a chave de ordenação, mudava a **ordem dos achados que o médico
vê**. Agora a chave é o `ItemID`, e `PlanRuler` expõe `scoreItemId`.

**C · O nível de cada ponto agora é recalculado sobre a escala desenhada.**
Reusar `pickScoringItem` alinha a ESCALA com o escore, mas não garante que o `lab_results.level`
gravado bata: ele foi calculado com a idade/menopausa do dia da importação, e num item com recorte
etário o paciente troca de faixa. A régua desenha ESTA escala, então o ponto tem que ser
classificado por ela — senão a bolinha cai num segmento e o rótulo diz outro. `levelForValue()` usa
a mesma função do motor (`ScoreLevel.EvaluatesTrue`), sem reimplementar regra.
**Validação: 158 dos 159 pontos do gabarito do Ricardo são reproduzidos exatamente**, e o único
diferente é um que o gabarito trazia como não classificado. O legado já usava a escala.
Efeito colateral bom: 296 resultados do João nunca tinham passado pelo classificador (`level` nulo,
sem `classify_reason`) e agora entram na devolutiva — por isso "está bem" foi de 23 para 71.

**D · O endpoint não existia no OpenAPI.** Faltavam as anotações swag, então nem
`swagger.json` nem `api-types.ts` conheciam `plan-dossier` — contra a regra de rodar `pnpm generate`
no mesmo passo da mudança de DTO. Anotado e regerado; `PlanDossierResponse` e `PlanRuler` agora
estão nos tipos TS que a Fase 1 vai consumir.

Um alerta da revisão foi **verificado e descartado**: ela sugeriu resolver o dia em SQL com
`::date` supondo a sessão do banco em UTC. A sessão é `America/Sao_Paulo`, então `::date` erraria as
datas de lote — o caso A ficaria sem conserto.

### Gaps de DADO encontrados (não são bug de código)

1. **`PLND7C2752F` (Saturação de transferrina) está sem `unit`** em `score_items`; o deck mostra "%".
2. **"Ferritina - Mulheres Pós-Menopausa" tem 4 níveis**, contra 6 da variante pré-menopausa.
3. **Conversão de unidade incompleta em alguns exames.** Ex.: "Leucócitos (WBC) - Sedimento" tem
   escala em células/campo enquanto o laudo reporta em /mL. A régua e o `level` gravado concordam
   (as duas usam a mesma escala), então isso já afeta o escore hoje — o dossiê só tornou visível.

---

---

## 5 · Fase 1 — a entidade e os dois PDFs (CONCLUÍDA)

O deck saiu do `pacs/` e virou entidade do prontuário, com o MESMO conteúdo saindo em três formas.

| Arquivo | O quê |
|---|---|
| `database/migrations/00087_patient_plans.sql` | `patient_plans` com `content jsonb` + CHECK de que é LISTA (a mig 00060 já mostrou o estrago de gravar objeto onde se espera array). |
| `internal/models/patient_plan.go` | Model, UUID v7 no `BeforeCreate`. |
| `internal/pdfdoc/ruler.go` | **Port de `ruler.py` para Go.** Régua grande + mini. Precedente de SVG em Go: `pdfdoc/qr.go`. Sem Python no container. |
| `internal/pdfdoc/deck.go` | `deckCSS` + scaffold do slide + `RenderDeck` nos dois papéis + `CheckDeckOverflow`. |
| `internal/pdfdoc/deck_blocks.go` | Os 8 tipos de slide + `inlineHTML`. |
| `internal/pdfdoc/deck_test.go` | 12 testes, incluindo os de render e o de estouro. |
| `internal/services/patient_plan_service.go` | CRUD, versionamento, prévia e publicação. |
| `internal/handlers/patient_plan_handler.go` + `cmd/server/main.go` | 8 rotas, todas `RequireClinician`. |
| `apps/web/lib/api/patient-plans.ts` + `app/(authenticated)/patients/[id]/plano/page.tsx` | Tela de montagem, prévia, conferência e publicação. |

### O A4 sai do MESMO HTML

`RenderDeck(deck, paper)`: no 16:9 imprime direto em 1920×1080; no A4 aplica `fitSlidesToA4` como
`beforePrint` — envolve cada slide numa folha de 297×210mm, escala para a largura e centraliza. Não
existe segundo HTML nem segundo CSS. O fator de escala sai só de `a4SlideScale()`.

### O que o transbordo virou

`CheckDeckOverflow` mede cada slide no próprio Chromium, depois das webfontes carregarem, e a
publicação **recusa** (HTTP 422) listando slide, título e quantos px passaram. Era a armadilha nº 7
do processo antigo, contornada por um script solto em `/tmp` que dependia de alguém lembrar de
rodar. Conferido no dev: 8 réguas num slide devolvem
`slide 01 (Réguas demais): +282px embaixo`, e a publicação não acontece.

### Decisões tomadas na execução

- **Fraunces foi embutida, não trocada.** Ela carrega TODA a tipografia de display do deck (h1, h2,
  números grandes, doses) — trocar por Cormorant mudaria a cara da peça. A estática pesa 71KB, menos
  que qualquer uma das Inter que já iam embutidas. Fonte: `docs/site/fontes_capa/`.
- **O `@media print` foi a diferença entre 8 e 15 páginas.** Na tela os slides são uma pilha rolável
  com 40px de respiro; na impressão esses 40px somados a 1080px estouravam a página e cada slide
  virava duas folhas, a segunda em branco.
- **Publicação é SÍNCRONA**, contra o que o plano previa. O plano mandava usar job porque o Chromium
  é serializado por mutex global; medido, um deck de 8 slides sai em **1,04 s no 16:9 e 0,18 s no
  A4**, então uma publicação segura a fila da receita por poucos segundos, não por minutos. Uma fila
  de jobs só para isso seria peso sem retorno — e `ProcessingJob`, a única que existe, é acoplada a
  lote de exames (`lab_result_batch_id` NOT NULL). Reavaliar se aparecer deck muito maior ou
  concorrência real. Longe do `renderTimeout` de 35s.
- **O override de eixo que a Fase 0 pediu está resolvido de graça:** o eixo da régua é um campo do
  slide no JSONB, não algo calculado na hora de desenhar. Quem escreve o plano parte do eixo que o
  dossiê derivou e o aperta à mão quando um outlier esmaga a escala (PCR 63,1 num eixo que termina
  em 15,7), que é exatamente o que a Ana precisou.
- **A tela do EMR não é um editor visual de slides.** Ela lista, cria, mostra a prévia, mede e
  publica; o conteúdo é editado como a lista de slides. O caminho de autoria de verdade é a skill da
  Fase 3 — construir um editor visual agora seria escopo que o plano não pediu e que a skill torna
  desnecessário.
- **`inlineHTML` escapa tudo e devolve à mão só `em`, `b`, `strong`, `i`, `small` e `br`.** O texto
  vem do JSONB e entra num HTML que o Chromium executa; o `<em>` do punch precisa funcionar, o resto
  vira texto.

### Verificado ponta a ponta no dev

Criar → conferir (sem transbordo) → publicar → **2 documentos** no portal com nomes legíveis
(`João-Da-Silva_Plano_2026-08-31_01a05ac7.pdf` e `..._PlanoImpressao_...`) e `source_ref`
`patient_plan:<id>:v1:16x9` / `:a4`. Republicar sobe para **v2** e cria documentos novos sem
deduplicar. Slide que estoura devolve 422 com o diagnóstico. `tsc --noEmit`: 0 erros no projeto.

### Rotina de qualidade da Fase 1 — o que ela pegou

Build, `go vet`, `gofmt`, suíte completa, `tsc --noEmit` (0 erros) e casos-limite do endpoint.

**Antes da revisão, os casos-limite acharam um bug meu:** `planId` inválido devolvia **404 "plano
não encontrado"** em vez de 400. Armadilha do Fiber — `c.Status(400).JSON(...)` devolve **nil**
quando consegue escrever a resposta, então testar esse retorno com `if err != nil` (o jeito óbvio)
deixava o handler seguir com `uuid.Nil` depois de já ter escrito o 400. Afetava os 6 endpoints com
`planId`. Corrigido com um `ok` explícito + teste (`patient_plan_handler_test.go`). Varrido: nenhum
outro handler do projeto usa esse retorno como sinal.

**A revisão de código achou 8, e um era grave:**

1. **Escala inteiramente negativa saía pintada de "ótimo" (ALTA).** `rulerAxis` tinha um piso em
   zero, justificado por "exame não tem valor negativo" — falso para **T-score de densitometria**
   (`Densitometria - T-Score Coluna Lombar` e `Colo Femoral`, escala de -2,5 a -1), ECG - Eixo
   Cardíaco e GLS, todos no catálogo. Reproduzido: `edges=[-2,5 -2 -1,5 -1]` virava `axis=[0 1]`,
   todo segmento caía fora e a barra saía como uma faixa única de nível ótimo com a bolinha presa
   na ponta esquerda. **Um paciente com osteoporose (T-score -2,8) receberia um PDF dizendo o
   contrário do exame.** O piso agora só vale quando a própria escala é não-negativa. De quebra, o
   rótulo de faixa negativa era `-2,5--2`, ilegível; virou `-2,5 a -2`.
2. **Notação científica no PDF do paciente (MÉDIA).** `%.4g` fazia `12345,6` sair como
   `"1,235e+04"` e `1234,5` perder a casa decimal em silêncio — CK, triglicerídeo e ferritina passam
   de 1000. Trocado por `'f'` com arredondamento em 4 casas (o que o banco guarda). Herdado do
   `ruler.py`, não introduzido aqui, mas errado do mesmo jeito.
3. **`Ruler`/`RulerPoint`/`RulerSegment` sem tags `json` (MÉDIA).** Gravavam `Display`/`Segments`
   enquanto os tipos TS gerados declaravam `display`/`segments`: qualquer consumidor tipado leria
   `undefined`. Todas as outras structs de slide já eram tagueadas; essas três escaparam. `Dark`
   virou `json:"-"` — é derivado da variante do slide na hora de desenhar, e persistir deixaria
   régua com a cor de um fundo que mudou.
4. **Publicação parcial deixava documento órfão no portal.** O laço criava o documento 16:9 e só
   então renderizava o A4; falha no A4 deixava o 16:9 visível para o paciente enquanto o médico via
   "falha ao publicar" e o plano continuava rascunho. Agora renderiza os dois antes de publicar
   qualquer um.
5. **`.rez-legenda` ficou sem regra de CSS** ao portar o `deck.css`: a legenda do resumo saía no
   tamanho padrão do navegador (16px) dentro de um slide onde tudo tem 26-36px.
6. **Editar um plano publicado deixava os ponteiros de documento apontando para o conteúdo
   antigo.** Agora saem junto com a volta para rascunho; a data de publicação fica como registro, e
   a tela mostra "rascunho · vN no portal desde <data>".
7. **`handleCreate`/`handleSave`/`handleDelete` sem `try/catch`** — um save que falha ficava
   silencioso e o médico seguia achando que gravou.
8. **A tela não era alcançável por nenhum link.** Entrou como ação "Plano" na página do paciente,
   ao lado de Prontuário e Continuum.

Revalidado depois de tudo: gabarito do Ricardo segue em **67/68 no eixo e 68/68 nos segmentos**, o
dossiê do João segue com 111 réguas, e o round-trip de uma régua pelo banco volta com as chaves em
minúsculas e sem `Dark`.

---

## 6 · Fase 2 — o relatório como modo, a tela do portal e as três fontes (CONCLUÍDA)

### O relatório virou o terceiro modo

`pdfdoc/deck_report.go` achata os MESMOS slides no documento A4 fluido da papelaria, assinado com
ICP-Brasil. Nenhum texto é reescrito e nenhuma régua é redesenhada: a régua é SVG com `viewBox`,
então encolher de 1714px para os 170mm do miolo é só largura. Rota
`POST /patients/:id/plans/:planId/report` (`RequireDoctor` — relatório assinado é ato médico).

O deck 16:9 e o A4 paisagem continuam **sem assinatura**: são peça de comunicação. O relatório é o
documento clínico. Verificado: o mesmo plano gerou `v2:16x9`, `v2:a4` e o relatório, os três a
partir do mesmo conteúdo.

Limpeza que veio junto: `care_plan_report_service.go` foi de 288 para 165 linhas
(`buildReportHTML` morto, o mapa `reportLevelColor` que só vivia nele, e o campo
`pdf *ScorePDFService` injetado e nunca lido). O fluxo "assina ou degrada" virou `signOrDegrade`,
compartilhado pelos dois relatórios.

### A tela do portal é HTML nativo

`components/patient-portal/plan-deck.tsx` redesenha os 8 blocos para tela vertical: tipografia em
`clamp`, grades que caem para uma coluna, e a régua com o nome **acima** da barra em vez de numa
coluna lateral — no celular aquela coluna deixaria a barra com dois centímetros. Não há
`dangerouslySetInnerHTML`: a ênfase é reconstruída como elemento React, com a MESMA lista de tags do
servidor (`pdfdoc.inlineAllowed`).

O paciente só vê plano **publicado**, e o filtro está na consulta: um rascunho responde "não
encontrado", nunca "existe mas você não pode ver".

### As três fontes (consulta, anamnese, exames)

O dossiê lia só exames. Passou a ler as três, **inclusive o que está em rascunho**:

- **Exames** — já lia, e não filtra status do lote. Isso é load-bearing: no dev os 22 lotes estão em
  `pending`, e um filtro por `completed` devolveria um dossiê vazio.
- **Anamnese** — `anamnesis_items` liga direto ao `score_item_id`. O nível vem de `selected_level`,
  NUNCA de `numeric_value` (ler o número cru no lugar do nível já zerou o escore antes). O mesmo
  item respondido em consultas diferentes vira **histórico**, então tem direção igual a um exame.
  Não há filtro de "finalizada": a tabela nem tem status, e a anamnese da consulta de hoje é
  justamente a que interessa.
- **Consulta** — `consultation_vitals` entra como bloco factual (pressão, frequência, peso, altura,
  cintura, IMC), com as duas últimas medidas para dar direção. O deck cita esses números direto
  ("a pressão está em 120 por 70, sem remédio nenhum") e eles não vêm de exame nenhum.

Medido no dev: 102 achados de exame + 3 de anamnese + vitais, todos na mesma lista ordenada.

### A ordenação de "está bem" mudou por causa disso

Ligar a anamnese fez o paciente demo devolver **99 achados** em "está bem", quase todos
"Adrenalectomia: não" em nível 5 com zero pontos perdidos — a mesma inflação de checklist de
ausência já conhecida no escore ([[escore_inflacao_e_devolutiva_estudo]]).

A causa é de ordenação: em nível 4-5 ninguém perdeu ponto, então ordenar "está bem" por pontos
perdidos deixa a lista inteira empatada em zero e sem sinal. Agora **"se movendo" ordena por pontos
perdidos** (o que se está deixando na mesa) e **"está bem" ordena pelo peso do item** (o marcador
pesado que está no ótimo). O topo do demo passou a ser doença cardiovascular ausente (25pts),
câncer (20), tempo de sono (20) — que é o que uma devolutiva de fato diz.

### Rotina de qualidade da Fase 2

Build, `go vet`, `gofmt`, suíte completa, `tsc --noEmit` (0 erros) e casos-limite. A revisão achou
**9**, todos corrigidos. Os que valem registro:

1. **Republicar o relatório entregava o PDF antigo com o QR errado (ALTA).** O `source_ref` era
   `patient_plan:<id>:v<n>:report`, mas `Update` não sobe a versão — só `Publish`. Editar os slides
   e gerar o relatório de novo casava com o ref antigo: `CreateFromBytes` devolvia o PDF velho, o
   render novo ia fora, e o paciente ficava com um arquivo cujo QR aponta para um `IssuedDocument`
   diferente do que o banco registrou. A identidade do relatório é o `IssuedDocument`, não a versão
   do plano — a chave passou a ser `patient_plan_report:<docID>`.
2. **A régua saía a ~5px no relatório assinado.** As classes `.rg-*` só existiam no `deckCSS`; sem
   elas o `<text>` herda o corpo do documento, e como o valor é lido em unidades do `viewBox`, 13
   unidades viram menos de 5px impressos. Repetidas no `reportDeckCSS`, com teste.
3. **A tela do EMR usava o paciente do SELETOR, não o da rota.** Abrir `/patients/X/plano` com o
   paciente Y selecionado mostraria os planos de Y — e "Publicar no portal" entregaria a devolutiva
   de Y. Agora usa o parâmetro da rota, como a tela irmã de `subscriptions`; o seletor continua
   como guarda de sessão.
4. **`<br>` e `<small>` apareciam literais na tela** enquanto quebravam linha nos dois PDFs: a lista
   de tags do portal divergia da do servidor.
5. **As faixas da régua do portal eram empilhadas em fluxo**, o que assume que cobrem o eixo sem
   buraco; um nível com limite ilegível é descartado na derivação e todas as faixas seguintes
   escorregariam, deixando a bolinha do paciente sobre a cor errada. Agora cada faixa é posicionada
   pelo próprio valor, como no Go.
6. **A folga do eixo encolhia em valor negativo** (multiplicar -3,5 por 0,96 o aproxima do zero), e
   o eixo nunca continha o ponto. Passou a crescer pela magnitude.
7. Publicar duas vezes um plano intocado criava uma versão e um par de PDFs idênticos.
8. `signOrDegrade` com serviço de assinatura ausente devolvia um PDF que **afirma** ser assinado.
9. Prévia de rascunho vazio dava 500 em vez de 400.

---

## 7 · Fase 3 — a skill de autoria (CONCLUÍDA)

`.claude/skills/plano-paciente/` + o comando `/plano`. É onde mora o que o EMR não deriva: o arco,
os títulos e as frases de fechamento.

| Arquivo | O quê |
|---|---|
| `.claude/skills/plano-paciente/SKILL.md` | Os três modos (CRIAR / REVISAR / PUBLICAR), o gate do dossiê, a voz e as regras editoriais. |
| `.claude/skills/plano-paciente/referencias/gramatica-v2.md` | Os 8 blocos com o JSON de cada um, e as 4 regras de desenho com a evidência por trás. |
| `.claude/commands/plano.md` | O disparo, no padrão de `/aula` e `/ebook`. |
| `CLAUDE.md` | Skill registrada no índice do monorepo. |

### O gate que define a skill

**É proibido escrever um slide antes de ler o dossiê.** Todo número, nome de exame, data e valor sai
de `GET /patients/{id}/plan-dossier`. Dado que não está no dossiê não entra no plano: ou se confirma
na fonte, ou o slide muda. É a regra que separa esta skill de um gerador de texto plausível, e ela
existe porque aqui inventar um número é o pior erro possível.

O arco é proposto e **espera aprovação** antes de virar slide. A skill também documenta o que o
dossiê NÃO resolve (o porquê clínico de um achado, conduta ainda não registrada, contexto de vida) e
manda perguntar ao médico em vez de preencher com plausibilidade.

### Validação

Os exemplos JSON da referência foram submetidos à API exatamente como estão documentados: os **8
blocos** fazem round-trip preservando estrutura (régua com 3 faixas, 2 pontos e nota; takeaway com
destaque e 2 grupos), a conferência de estouro passa, e os três modos publicam (16:9 e A4 em 200,
relatório assinado em 201). Conferido visualmente: o resumo renderiza com a mini-régua aninhada.

Isso importa porque a referência é um contrato: se o JSON documentado não casasse com
`pdfdoc.DeckSlide`, a skill produziria planos que o servidor recusa.

O texto de exemplo que vira conteúdo de paciente está limpo de travessão. A prosa interna dos
arquivos usa 9, o menor número entre as 8 skills do projeto (a maior tem 71).

### Rotina de qualidade da Fase 3 — o que ela pegou

A revisão achou **8**, e três mudaram comportamento de verdade:

1. **Editar um plano publicado fazia ele SUMIR da tela do paciente (migration 00088).** O plano
   guardava um `content` só; salvar devolvia o registro para rascunho, e como a consulta do portal
   filtrava `status = 'published'`, corrigir um typo tirava do paciente a devolutiva que ele já
   tinha recebido. A instrução da skill promete o contrário, em texto. Agora publicar congela uma
   cópia em `published_content`, e é ela que o portal lê: o médico edita o rascunho à vontade sem
   mexer no que está na mão do paciente. O portão da consulta deixou de ser `status` e passou a ser
   `published_content IS NOT NULL`.
2. **O render 16:9 era o único do pacote que não esperava as webfontes.** A conferência de estouro
   mede COM Fraunces e Inter carregadas e diz que cabe; o PDF imprimia com as métricas da fonte de
   fallback e o `overflow:hidden` cortava o excesso em silêncio — exatamente o que a conferência
   existe para impedir. Dava também para o A4 e o 16:9 do mesmo conteúdo saírem diferentes.
3. **A tela do EMR conferia e publicava o conteúdo SALVO enquanto o médico editava a caixa.** Colar
   slides novos, clicar em "conferir se cabe" e receber "todos os slides cabem" media o conteúdo
   antigo; publicar em seguida entregava ao paciente o deck de antes da edição, com o toast dizendo
   que deu certo. Agora há marca de alteração não salva, com aviso e os três botões travados.

Os outros cinco: `Publish` segurava o Chromium por **três** renders e não dois (a conferência
gerava um PDF inteiro só para rodar o script de medição — agora mede na mesma passada do 16:9, e o
comentário que orçava o tempo estava subestimando); um nível com operador `<` era rotulado `≤`,
dizendo ao paciente que o valor da fronteira pertence à faixa errada; `PublishReport` deixava um
`IssuedDocument` órfão a cada tentativa que falhava; `GetPublished` recebia os argumentos na ordem
inversa das irmãs, num caminho onde a troca silenciosa buscaria o plano do paciente errado; e a
migration 00087 misturava `CREATE TABLE IF NOT EXISTS` com um `ADD CONSTRAINT` sem guarda.

O teste-guarda do portal **caiu sozinho** quando mudei o filtro, que é o trabalho dele; foi
reescrito para o invariante novo.

---

## 8 · Os testes de aceitação do plano (EXECUTADOS)

O plano aprovado exigia quatro comparações contra material real. Foram feitas, e três acharam
defeito.

### Fase 1 — reconstruir o deck do Ricardo pelo motor Go e comparar slide a slide

O `deck.html` dele foi convertido para o contrato `pdfdoc.DeckSlide` (20 slides, 22 réguas, 35
linhas de tabela), renderizado pelo motor Go e comparado em PNG contra `deck-jose-ricardo.pdf`.
**Passou**: réguas, rampa, seta de direção, ticks, legenda, punch e numeração saem equivalentes.

**Achou o maior buraco da gramática:** o bloco `table` não existia. Tabela densa é **8 dos 20
slides do Ricardo e 9 dos 21 da Ana** — 40% do deck real, o bloco MAIS usado, e o motor não sabia
desenhá-lo. Criado, com modelo por COLUNA e não por campos fixos: a primeira tentativa (item/dose/
porquê) jogava a prosa da terceira coluna no campo de dose, que tem `white-space:nowrap`, e o slide
vazava 2472px para fora da moldura.

### Fase 2 — relatório novo × relatório antigo, mesmo paciente

**Passou**: mesma papelaria, sem regressão. E fechou a unificação que estava pela metade.

### Fase 2 — QA visual do portal em viewport de celular (Playwright)

**Passou**: em 375, 768 e 1280px o `scrollWidth` bate com o viewport, sem rolagem horizontal; 21
seções renderizadas; menor fonte 12px. Conferido na foto: régua com o nome acima da barra,
"antes 1,00 · agora 0,99", cartão de passos em coluna única.

Duas armadilhas de infra no caminho, que valem registro: `pnpm start` serve o `.next` ANTIGO (a
rota nova dava 404 até subir em modo dev), e o `NEXT_PUBLIC_DEV_BYPASS_AUTH` injeta um admin falso
que sobrepõe a sessão do paciente, tornando o portal impossível de testar como paciente sem
desligá-lo.

### Fase 3 — a gramática da skill contra o deck real da Ana

**Passou, e melhor do que eu esperava**: os 21 slides da Ana, feitos à mão antes da skill existir,
encaixam **100%** nos blocos documentados, sem sobra.

Mas o teste também mostrou que **o bloco `sequence` que eu inventei nunca foi usado**: nos dois
decks reais, "Os próximos três meses, em ordem" é uma TABELA. A gramática passou a dizer isso.

### A unificação do relatório, que estava pela metade

O plano pedia que `RenderCarePlanReport` consumisse o mesmo conteúdo; eu tinha criado um render
paralelo e deixado dois geradores. Agora a derivação do escore produz **slides** e passa pelo MESMO
`RenderPlanReport`. `pdfdoc/care_plan_report.go` foi apagado. Existe um renderizador de relatório.

Isso expôs mais dois defeitos, os dois corrigidos:

- **O achatador do relatório não conhecia o bloco `table`**, adicionado depois. O A4 assinado saía
  com o título da seção e nenhuma linha — e os 8 slides de tabela do Ricardo sairiam vazios.
- **Tabela longa invadia o rodapé e a assinatura.** O paginador da papelaria move blocos inteiros
  entre páginas mas não sabe partir um; 21 linhas viravam um bloco mais alto que a área útil. Agora
  a tabela é fatiada em blocos de 12 linhas, cada um repetindo o cabeçalho.

### Rotina desta rodada — 5 achados, e o nº1 dizia que meu conserto não funcionava

1. **O fatiamento da tabela não estava funcionando (ALTA).** As fatias eram emitidas DENTRO da
   seção que `flattenSlide` já tinha aberto. `paginateDoc` só pagina filhos DIRETOS de `#src`, então
   viravam netos e o slide continuava um bloco único mais alto que a página — exatamente o que o
   comentário do conserto dizia ter resolvido. Eu tinha até atribuído o rodapé apertado ao "design
   do paginador"; era o meu bug. Agora as fatias são irmãs. Provado: tabela de **60 linhas ocupa 5
   páginas** em vez de estourar uma.
   Ao mover o ramo para o topo da função apareceu um segundo erro meu, com crash de verdade
   (`HierarchyRequestError: the new child contains the parent`): o ramo rodava depois de a função já
   ter aberto a seção, deixando uma `<div>` sem fechar. O teste que eu tinha escrito **contava**
   seções e por isso não pegou; passou a checar o balanceamento das tags.
2. **O relatório usava os cartões do SLIDE** (`.two`/`.card`), classes que só existem no `deckCSS`:
   no A4 assinado sairiam sem moldura e, com 3+, numa grade dimensionada para 1920px numa folha de
   170mm.
3. **A tabela sumia da tela do paciente.** O portal não tinha ramo para o bloco `table`, então o
   PDF que o paciente baixa e a tela que ele abre diziam coisas diferentes — no bloco mais usado do
   deck. No celular ela vira lista (três colunas em 375px não se lê). E as células passaram a ir por
   `RichText`: um `<b>` no conteúdo funcionava no PDF e aparecia literal na tela.
4. **`sujo` nunca voltava a falso depois de salvar.** O Go normaliza a ordem das chaves e derruba
   campo vazio (`omitempty`), então o texto colado nunca voltava a bater com o salvo: Prévia,
   Conferir e Publicar ficavam travados até trocar de plano e voltar. Agora a caixa é reescrita com
   o que o servidor gravou.
5. **`RenderDeckMeasured` estava escrito e não ligado**: `Publish` seguia fazendo três passagens
   pelo Chromium. Ligado — são duas.

---

## 9 · O que fica em aberto

- **Semear o plano a partir do dossiê na tela do EMR.** Hoje "Novo plano" nasce com capa e fecho
  vazios; a autoria real é pela skill. Um semeador mecânico (capa + réguas dos top achados) é
  possível, mas produziria deck sem arco — por isso ficou com a skill.
- **A tela do EMR não é editor visual de slides.** Lista, cria, mostra prévia, mede e publica; o
  conteúdo é editado como a lista de slides.
- **Rodar a skill num paciente real** (Ana ou José Ricardo, que estão em prod) e comparar o arco
  proposto contra o deck que já existe. A validação até aqui usou o dossiê do dev.
- **Migration 00087 só está no dev.** Prod segue no goose 86.
- Os gaps de dado do capítulo 4 continuam abertos (unidade da saturação de transferrina, níveis da
  ferritina pós-menopausa, conversão de unidade do sedimento).

### Regra editorial que a evidência impõe

Barra horizontal colorida bate tabela e número cru em compreensão, e mostrar a **faixa-meta no lugar
da faixa de referência** do laboratório (não somada a ela) melhora o entendimento — a régua já faz as
duas coisas. Mas o melhor desempenho medido é da barra **com rótulo avaliativo**, e `ruler.py` foi
escrito de propósito "sem rótulo escrito, sem numeral de nível". Hoje o rótulo vem de fora (título da
seção, `punch`, `note`). Vira regra: **nenhuma régua entra num slide sem rótulo avaliativo visível no
mesmo slide.** É também o alívio de acessibilidade que o próprio `deck.css` documenta para o
contraste do dourado (2,6 = WARN).

Fontes: [revisão sistemática JMIR 2024](https://www.jmir.org/2024/1/e53993) ·
[faixa-meta substituindo referência](https://www.ncbi.nlm.nih.gov/pmc/articles/PMC6231727/) ·
[harm anchors / Zikmund-Fisher](https://www.ncbi.nlm.nih.gov/pmc/articles/PMC5891666/) ·
[direção da mudança](https://www.advancesinpro.org/article/S3050-6964(26)00018-2/fulltext) ·
[AHRQ teach-back](https://www.ahrq.gov/health-literacy/improve/precautions/tool5.html).
