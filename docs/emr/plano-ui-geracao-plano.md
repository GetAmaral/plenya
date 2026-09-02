# UI de geração do plano — plano de implantação aprovado

> Aprovado em 01/09/2026. Estudo que o originou: [plano-ui-geracao-estudo.md](plano-ui-geracao-estudo.md).
> Feature de planos já em produção: [plano-planos-paciente.md](plano-planos-paciente.md).
> Documento vivo: cada fase termina atualizando o estado aqui.

## Estado

| Fase | O quê | Estado |
|---|---|---|
| 1 | Identidade do slide + revisões + concorrência otimista | ✅ dev (`0abbfff0`, mig 00092) |
| 2 | Dossiê congelado + fontes fora da prévia + coluna do prontuário | ✅ dev (`7bdff7ba`, `86851e94`, mig 00093) |
| 3 | Ops, classificação e índice numérico (sem IA) | ✅ dev (`28c13928`), 28 testes |
| 4 | Cartões e os 6 editores | ✅ dev (`3ef626ef`) |
| 5 | A conversa, triagem e sugestões | ✅ dev (`98dbec0a`, `53512fe6`, mig 00094) |
| 6 | Auditoria de conteúdo gerado + histórico legível | ✅ dev (mig 00095) |

### Medido na fase 2

- prévia: **1.970.060 → 75.871 bytes** (96% eram as sete fontes em base64), 26× menor
- dossiê congelado do paciente de dev: 107 réguas, 28 se movendo, 73 no bom
- envelhecimento e diff verificados: recalcular o escore acende `stale`, e o diff responde
  "nenhum exame citado mudou" — restrito ao que o deck cita, não ao dossiê inteiro

### Medido na fase 3

Contra o dossiê real de uma paciente de produção (83 réguas, 2.086 fatos indexados):
**197 de 197 valores do histórico reencontrados**, nenhum número inventado casou. O fixture real
não foi versionado (dado clínico mora em `pacs/`); ficaram no teste os onze formatos observados,
com as contagens.

A triagem é toda testável sem chamar modelo, e é o que decide se a feature é segura.

### O que a fase 6 fechou

**A trilha passou a apontar o objeto certo.** Toda rota de plano mora sob `/patients/:id/plans/…`,
e o middleware classificava tudo como acesso a paciente e gravava o id do PACIENTE. A trilha dizia
"alguém acessou o paciente X" para qualquer ação sobre um plano, uma revisão ou uma sugestão:
verdadeiro e inútil. Agora `/plans` é testado antes de `patients`, e o id gravado é o do parâmetro
mais específico da rota (`suggestionId` → `revisionId` → `planId` → `itemId` → `id`). Conferido que
as rotas de paciente puras continuam classificadas como `patients`.

**`ai_touched_paths` na publicação** (migration 00095): no ato de publicar, percorre a cadeia de
revisões e guarda os caminhos cujo ÚLTIMO escritor foi o assistente. "Esta devolutiva tem N campos
gerados que ninguém reescreveu" vira campo em vez de arqueologia em JSONB.

Três decisões que a implementação obrigou a tomar, e o motivo:

- **Salvar à mão NÃO limpa a atribuição.** Edição pela tela reescreve o conteúdo inteiro e não
  declara caminhos. Tratar isso como "o médico passou por tudo" é exatamente a suposição que a
  evidência de viés de automação desmente. O número super-reporta de propósito: errar para cima é o
  lado certo de errar aqui, e a tela diz isso com todas as letras ("sem reescrita registrada
  depois da ferramenta").
- **A revisão de aceite de sugestão passou a gravar as `ops`.** Ela não gravava, e o efeito era o
  pior possível: a sugestão NUMÉRICA aceita, o conteúdo de maior risco do fluxo inteiro, não
  deixava caminho nenhum e sumia da trilha. Um plano cheio de números gerados apareceria como se
  nenhum tivesse vindo da ferramenta.
- **O cálculo percorre a cadeia inteira, não a `plan_version` corrente.** Escopar por versão foi
  erro meu: a partir da segunda publicação só enxergaria o que mudou desde a anterior, e um deck
  escrito pela ferramenta na v1 e republicado apareceria limpo. A pergunta é sobre o conteúdo que o
  paciente tem na mão.

**O histórico ficou legível, e o desfazer existe.** Isto era promessa da fase 1 entregue pela
metade: as revisões eram gravadas desde o começo e nunca houve rota para lê-las nem para restaurar.
`GET .../revisions` devolve a lista sem o `content` (a lista carregaria dezenas de decks inteiros
para desenhar dezenas de linhas) e `POST .../revisions/:revisionId/restore` volta o rascunho.
Restaurar **grava uma revisão nova** com o conteúdo antigo em vez de apagar: desfazer que destrói
histórico é o mesmo defeito que a tabela existe para consertar. Na tela, um `Sheet` com quem
escreveu cada gravação, o que mudou e o selo de IA na publicação.

Verificado de ponta a ponta em dev: 11 revisões listadas com autor e motivo corretos; restaurar a
revisão 1 devolveu o título original e o plano a rascunho; o mesmo plano pedido sob outro paciente
responde 404 nas duas rotas; a auditoria da restauração grava o id da REVISÃO.

E o guarda de estouro pegou a própria ferramenta: as três sugestões aceitas no teste estouraram o
slide em 1099px na horizontal e a publicação recusou com 422, listando o slide.

### A revisão global (2e9204c0) — o que ela pegou que as fases não pegaram

A rotina do projeto é revisão + testes + correções + commit. As seis fases rodaram só a metade de
testes; esta é a revisão que faltou. Onze achados, todos confirmados antes de mexer, e o padrão é o
argumento a favor de rodar a global: **quase tudo estava na costura entre fases**, com cada fase
internamente coerente.

| # | O quê | Efeito real |
|---|---|---|
| 1 | Sugestão estrutural inaceitável por construção | `add`/`reorder` derrubavam o lote inteiro, levando junto as de texto |
| 2 | `published_revision_id` sempre NULL | 7 de 7 planos publicados; a recuperação da v1 nunca existiu |
| 3 | Perda de escrita de 10 a 20 s | PUT do médico durante a chamada era sobrescrito em silêncio |
| 4 | Tipo errado do modelo abortava o turno | a mensagem do médico não era gravada, contra o que o código promete |
| 5 | CSP bloqueava as fontes do deck | prévia em fallback, mentindo sobre o que cabe |
| 6 | Coluna de tabela reordenada não movia as células | dose sob o rótulo errado, sem erro na tela |
| 7 | Chave de idempotência gerada por tentativa | o 409 do servidor era código morto |
| 8 | Salvar não mandava `expectedRevision` | o 409 da concorrência otimista nunca disparava |
| 9 | Proveniência não visitava `When`/`Unit`/`Name` | posologia inventada não acendia o aviso |
| 10 | Classificação do aceite contra o deck já aplicado | `remove` aceito virava "desconhecido" |
| 11 | `proximoSeq` sem trava | dois turnos simultâneos batiam no índice único depois de pagos |

Duas coisas que a correção do #1 revelou e a revisão não tinha visto: `base_hash char(64)` **enche
com espaço**, então o Postgres devolvia 64 espaços onde o Go esperava vazio e o guarda de "o slide
mudou" recusava toda sugestão de `add` (migration 00096); e a trava de linha que conserta o #3
serializa `proximoSeq` e mata o #11 de graça.

### Medido na fase 5, contra o modelo de verdade

Turno real em dev, pedindo para citar um valor de exame:

- 1 operação de texto **aplicada direto** (reescrita sem número novo)
- 4 alterações numéricas viradas **sugestão**, com a origem anexada
- 2 **recusadas**: o modelo tentou reescrever `segments` e `history` da régua, afirmando na
  resposta que os estava "copiando exatamente do prontuário". Reescrita é reescrita.
- cache de prompt: **9.390 tokens lidos** a partir do segundo turno, como projetado
- latência: 5 a 7 segundos por turno

Três defeitos que só o teste de ponta a ponta pegaria, todos corrigidos: o índice de idempotência
filtrando `IS NOT NULL` enquanto o Go grava string vazia (quebrava o segundo turno de toda
conversa); o `down` da migration deixando FK órfã (rollback impossível de re-migrar); e data por
extenso virando prova falsa ("7 de fevereiro de 2026" gerava dois alarmes espúrios).

E uma melhoria vinda de ver a saída real: as origens candidatas são ordenadas por relevância.
"Sua lipase estava em 27 U/L" casava primeiro com o limite do eixo do cortisol, que por acaso
também vale 27. Mostrar a origem errada é pior que mostrar várias, porque parece autoritativo.

### Fase 4 e o débito de QA visual

Instalei o Playwright e consegui **verificar a renderização**: os cartões, as miniaturas reais, os
selos de "alterado" e "não cabe", a borda vermelha e a faixa de aviso aparecem corretos, sem
nenhum erro de console. Foi assim que apareceu um defeito que typecheck e render não pegam: o
cabeçalho do cartão era um `<button>` com mais de vinte `div`s dentro, HTML inválido.

**A interatividade continua não verificada.** No harness (navegador em container, servidor Next em
outro), o React não hidrata: 21 scripts carregam, nenhuma resposta 4xx, e nada fica clicável —
inclusive um `DropdownMenu` do Radix, componente padrão que funciona no resto do app. A causa
provável são as falhas de handshake do HMR, que a partir de uma origem estrangeira abortam a
hidratação. É limitação do ambiente de dev, não do código, mas continua sendo um débito real: o
comportamento de clicar, arrastar e digitar não foi visto funcionando.

Caminho para pagar de verdade: rodar o QA contra um build de produção local (`next build && next
start`), onde não há cliente de HMR.

### Débito consciente

QA visual da coluna com dados reais não foi feita: o eslint do repo está com o config quebrado
(referência circular, falha igual em arquivo intocado) e o Playwright não está instalado. Typecheck
limpo e a página renderiza sem erro de runtime. Fica para a fase 4, onde o editor torna a
verificação visual obrigatória.

---


## Context

O plano de devolutiva já existe como entidade: renderiza 16:9 e A4 do mesmo HTML, mede estouro,
publica no portal e congela `published_content`. O que não existe é a **autoria**. Hoje
`apps/web/app/(authenticated)/patients/[id]/plano/page.tsx` mostra a lista de planos e uma
`<Textarea>` com o JSON cru dos slides. Na prática, quem escreveu os decks reais fui eu, por API,
fora do EMR.

O pedido: a tela que compila o prontuário, monta o rascunho visual e sustenta a discussão até a
versão final.

**O buraco estrutural.** `PatientPlanService.Update` (`patient_plan_service.go:148`) faz
`plan.Content = req.Content` e salva. `patient_plans.version` conta **publicações**, não edições.
O deck da Ana passou por quatro versões nesta sessão e as três primeiras não existem mais. Não se
discute um documento cujos estados anteriores foram apagados.

**A régua de segurança vem de evidência, não de gosto.** Em mensagens de portal redigidas por IA
com erro plantado, mais da metade dos clínicos não pegou todos os erros e 35% a 45% enviaram sem
editar, com 90% declarando confiar na ferramenta
([npj Digital Medicine 2025](https://www.nature.com/articles/s41746-025-01586-2)). Em prescrição
assistida, médicos migraram de conduta certa para errada em 5,2% dos casos por sugestão do sistema
([automation bias em prescrição](https://www.ncbi.nlm.nih.gov/pmc/articles/PMC5356416/)). A
conclusão de desenho é direta: **não confiar na revisão por leitura para o que dá para verificar
por código.** Todo número do deck sai de um conjunto fechado, o dossiê. O servidor verifica em vez
de o prompt pedir.

## Decisões já tomadas com o usuário

1. **Autoridade híbrida da IA**: texto aplica direto e reversível; número ou régua vira sugestão
   aceita slide a slide.
2. **Editor = cartão com formulário por tipo de bloco**, miniatura + campos.
3. **A primeira entrega inclui a conversa.**
4. **Ao aceitar sugestão numérica, a origem do número aparece ao lado do botão** (valor do dossiê,
   exame, data). Um clique para aceitar; sem obrigar a expandir.
5. **O deck publicado não menciona uso de IA.** O documento é do médico, que assina e responde.
   (A trilha `author_kind` nas revisões existe por outra razão: é o que faz desfazer e restaurar
   funcionarem. Não é divulgação, e não aparece para o paciente.)

## O que foi medido, não estimado

| Medida | Valor | Consequência |
|---|---|---|
| `preview` geração | 5 a 95 ms | é só HTML, sem Chromium |
| `preview` payload | **1,97 MB, 96% fontes base64** | prévia fiel não pode ser ao vivo sem separar as fontes |
| deck sem fontes | 75 KB | alvo depois da separação |
| `overflow` | 220 ms a quente, 1,27 s a frio | mede o deck **inteiro**, no mutex global do Chromium |
| `plan-dossier` | 500 ms, **~28-30 queries** | congelar no plano |

**Uso real dos blocos** (2 decks v2 + dev): `table` e `rulers` dominam; `sequence` **nunca foi
usado**; `two-cards` é sempre 2. Campos universais: `eyebrow` 100%, `title` ~90%, `punch` ~85%.
Tetos comprovados: réguas ≤4 (8 estoura, há teste), linhas de tabela ~8, colunas ≤3, linhas de
resumo ≤4, grupos de takeaway ≤3.

---

## As quatro correções que mudaram o desenho

**1. `DeckSlide` não tem identidade.** Só tem `Kind` (`pdfdoc/deck_blocks.go:39`). Endereçar slide
por índice significa que uma sugestão criada às 14h02 sobre "o slide 6" aplica no slide errado se
o médico reordenou às 14h05, em silêncio, escrevendo número de exame errado num documento que o
paciente lê. Mesma classe de falha do `overflow:hidden`. **`DeckSlide.ID` é pré-requisito de tudo**,
com backfill em SQL e preenchimento preguiçoso no `load()`.

**2. A separação texto/número por CAMPO não funciona.** `punch`, `title`, `lede` e célula de tabela
rotineiramente carregam número. A classificação tem que sair do **diff**, não do campo:

```
op de texto:  novos := numerais(depois) \ numerais(antes)
              len(novos) == 0  → aplica direto      ("encurta o título", "tira o jargão")
              caso contrário   → vira sugestão      ("subiu de 48 para 96" no punch)
```

E há um **terceiro nível** que não é nem texto nem número: `add`, `remove`, `reorder`. Perder um
slide em silêncio é pior que um número errado. Estrutural é **sempre** sugestão.

**3. O layout de três colunas não cabe no consultório.** Com a sidebar de 256px e `lg:p-8`, num
laptop de 1366 sobram **310px** para o centro. Não comporta miniatura mais formulário. O padrão
passa a ser **duas colunas** (dossiê + cartões); a conversa é terceira coluna só a partir de `2xl`
e um `Sheet` à direita abaixo disso — que é o padrão da casa para painel auxiliar, não gambiarra de
responsivo (`components/conversations/dossier-panel.tsx` abre assim em tela cheia).

**4. `DeckSummaryLine` não tem `code`.** Sem âncora, o `value` da linha de resumo é string solta e a
proveniência ali é **estruturalmente impossível** de validar. Precisa de `Code` no struct, exigido
pelo schema em qualquer op que crie linha de resumo. Sem isso metade dos números do slide mais denso
fica inauditável.

## Onde os dois designs discordaram, e a decisão

O design de frontend propôs que **a IA nunca escreva no servidor**: devolve um patch, o cliente
aplica no rascunho local, desfazer vira pilha em memória. Elimina concorrência e é mais simples.

**Fica com a escrita no servidor**, pelo motivo que sustenta a feature inteira: se a edição da IA
só existe no cliente até o médico salvar, a revisão registra **o médico** como autor do texto da IA,
e a proveniência — que é o mecanismo central de segurança — se perde exatamente onde mais importa.
Some também a conversa que sobrevive a fechar o notebook.

O risco que o frontend apontou é real e é tratado com o que ele mesmo propôs, mais o lock do backend:
`expectedRevision` no PUT e no turno (409 quando o cliente está velho); **nunca auto-aplicar no
slide que está expandido em foco** (rebaixa para sugestão); e autosave nas fronteiras naturais
(colapsar cartão, trocar de slide), não por timer. "Desfazer" passa a ser restaurar revisão, que é
durável e auditável em vez de morrer com a aba.

## Duas correções aos próprios agentes

- **Os 9 blocos já estão tipados em TS.** Os structs do `pdfdoc` chegam ao swagger e ao
  `packages/types/src/generated/api-types.ts`, com os comentários do Go virando `@description`.
  Transcrever as structs Go à mão violaria a Regra de Ouro nº 1 do projeto. A interface manual em
  `lib/api/patient-plans.ts:17`, com blocos `unknown`, é duplicata a eliminar.
- **A prévia não é "praticamente grátis".** São 5 ms de geração mas **1,97 MB de payload**, 96%
  fontes. `deckHTML` (`pdfdoc/deck.go:98`) concatena `deckFontFaces()` sempre. Um parâmetro que
  troque as fontes embutidas por link (só no caminho do navegador; o PDF continua precisando delas
  inline, porque o Chromium renderiza de string sem rede) derruba para 75 KB e torna a prévia fiel
  barata.

---

## Modelo de dados (migration 00092)

Cinco mudanças, na mesma migration porque são o mesmo problema:

- **`DeckSlide.ID`** + backfill em `content` e `published_content`, e `DeckSummaryLine.Code`.
- **`patient_plan_revisions`** — guarda o estado **resultante**, não o delta (restaurar é copiar uma
  linha; cadeia de patches corrompida no meio mataria o histórico dali para frente). Segue o
  precedente `IntegratedPlanRevision` (`models/patient_continuum.go:169`), com quatro colunas a
  mais: `author_kind` (human/assistant/system), `reason` (edit/ai_apply/suggestion_accept/restore/
  publish), `ops`, `ai_model`+`ai_prompt_version`. `created_by_id` é **sempre o clínico logado**,
  inclusive quando a IA escreve: alguém responde pelo que o paciente leu.
  - `seq` conta **edições**; `patient_plans.version` continua contando **publicações**;
    `plan_version` liga as duas. Na tela: "v2 no portal · rascunho, edição 47".
  - `is_publication` conserta de graça um bug de hoje: `Publish` sobrescreve `published_content`, e
    republicar destrói os bytes exatos da v1 para sempre.
- **Coalescência é o que torna a tabela viável.** Autosave com debounce de 3 s por 90 min daria
  ~1.800 revisões × ~60 KB = 108 MB por plano, ~100 GB/ano. Com janela de 2 min por autor humano:
  ~45 revisões, ~0,5 GB/ano. Save que não muda o hash não vira linha. Revisão de `ai_apply`,
  `suggestion_accept`, `restore` e `publish` **nunca** coalesce.
- **`patient_plan_dossiers`** — tabela, não coluna: `PatientPlanService.load` faz `SELECT *` e é
  chamado por preview, overflow, publish e get; um jsonb de centenas de KB seria detoastado no
  caminho que hoje custa 5 ms. Guarda o payload e um `numeric_index` materializado (é contra ele
  que se valida). Congela no `Create` e em refresh explícito, **nunca automático** — refrescar
  sozinho troca número debaixo do cursor e invalida a base contra a qual as sugestões pendentes
  foram validadas.
  - Envelhecimento é detectado com **uma** query de marcas d'água (`max(created_at)` de snapshots,
    labs e vitais), não remontando o dossiê.
  - `POST .../dossier/refresh` devolve o diff **restrito aos códigos citados no deck**, com onde
    cada um é citado. "O escore mudou 3×" vira duas linhas em vez de "está velho, boa sorte".
- **`patient_plan_messages`** e **`patient_plan_suggestions`** — a conversa persistida (com
  `client_message_id` para idempotência de um POST de ~15 s) e as sugestões com ciclo de vida
  próprio (`pending/accepted/rejected/stale/superseded`), `provenance` e `base_hash`. O `base_hash`
  é o que impede o painel de sugestões de 20 min atrás de apagar o que o médico acabou de escrever.
- **`patient_plans`** ganha `revision_seq` (token de concorrência otimista), `current_dossier_id`,
  `published_revision_id`.

## O contrato com o modelo

Um tool só, forçado (`tool_choice`), no padrão que a casa já usa em cinco lugares
(`InterpretLabResult`, `ai_note_service`), com `strict: true`. Devolve `reply` (prosa em português
para o médico) e `operations[]` (`add`/`edit`/`remove`/`reorder`), com **`numerals[]` obrigatório**:
todo número escrito precisa vir declarado com a origem no dossiê. Modelo que inventa número
normalmente inventa a origem, e origem inventada não existe no índice.

**Vocabulário de caminhos (allowlist no servidor), não JSON Pointer livre** — pointer arbitrário
deixa o modelo escrever em `rulers[0].segments[2].b`. Quatro classes:

| classe | exemplos | tratamento |
|---|---|---|
| autoral-texto | `title`, `punch`, `cards[i].body`, `rulers[i].display/sub/note` | aplica se não introduziu numeral novo |
| numérico | `summary.…lines[j].value/unit`, `takeaway.…dose`, célula de coluna `dose` | sugestão |
| do dossiê | `rulers[i].code/segments/history`, `summary.…lines[j].ruler` | **rejeita**, nem sugere |
| numérico autoral | `rulers[i].axis` | sugestão, com invariante |

O invariante do `axis`: o novo eixo tem que **conter** todos os `segments` e todos os pontos de
`history`. Eixo que corta um ponto do paciente esconde dado sem erro, mesma classe do overflow.

### O limite honesto da validação

> **A validação prova que o número EXISTE no dossiê. Nunca prova que ele significa o que a frase
> diz.** "Sua ferritina está em 96" e "seu colesterol está em 96" passam identicamente se 96
> existir em qualquer lugar. O que ela entrega é **superfície de julgamento**: a op vira sugestão
> com a origem candidata anexada, e o médico julga uma afirmação específica em vez de reler um
> parágrafo.

Não alcança: unidade trocada com o mesmo número, número por extenso ("dobrou", "três meses"),
comparativo calculado ("melhorou 40%"), e toda inferência clínica. Isso precisa estar escrito no
código e visível na tela — a pior versão desta feature é a que parece garantir mais do que garante.

### Chamada

Síncrona, `POST` comum, timeout de 90 s. **Não** `ProcessingJob`: aquela fila tem FK `not null`
para `lab_result_batches`, o ticker de 3 s somaria silêncio a cada turno, e retry automático de um
turno que já aplicou ops re-aplica. O retry certo é o médico reenviar, protegido por
`clientMessageId`. Escape hatch pronto se o p95 desapontar: `readToolUseStream` já existe.

Três detalhes que não podem ser esquecidos:
- **Não mandar `temperature`.** Opus 5 e Sonnet 5 respondem 400. O código atual só funciona porque
  o default é `claude-sonnet-4-6`; trocar o default hoje quebraria `CompleteText` e o
  `ai_note_service`. O comentário em `InterpretLabResult` já registra isso.
- **Prompt caching é o segundo argumento para congelar o dossiê**: byte-idêntico entre turnos, com
  `cache_control` no fim do bloco `system`, os turnos 2..N leem ~90% do contexto a 0,1× do preço.
  Dossiê vivo estouraria o cache a cada turno. Conferir com `cache_read_input_tokens > 0`.
- **Tirar o nome do paciente do dossiê** que vai para a API. `PlanDossierPatient` manda hoje; use
  "o paciente". É de graça.

---

## A tela

```
2xl (≥1536)   [18rem dossiê] [1fr cartões] [22rem conversa]
lg – 2xl      [18rem dossiê] [1fr cartões]   + conversa em Sheet à direita
< lg          uma coluna de cartões          + dossiê e conversa em Sheet
```

Colunas laterais `sticky` com `ScrollArea` própria; o centro é quem rola.

**Coluna esquerda, o prontuário compilado.** Abas: se movendo · está bem · réguas · condutas ·
prescrições · exames pedidos · **lacunas**. A aba de lacunas lista o que trava o plano deste
paciente (sem vitais, sem data de nascimento, resultados marcados para revisão de unidade) — hoje
isso só aparece se alguém rodar SQL.

**Centro, os cartões.** Colapsado mostra miniatura, `eyebrow`, `title`, `punch` e badges
(`alterado`, `IA`, `sugestão`, `não cabe`). Expandido, um por vez, abre o formulário do bloco.
Reordenação com `@dnd-kit`, no padrão de `components/scores/VersionItemsEditor.tsx`.

- **A miniatura reusa `Slide` de `components/patient-portal/plan-deck.tsx`** (exportar; é uma linha).
  A armadilha: a tipografia é `clamp(…, 5vw, …)`, ancorada na viewport — estreitar o contêiner dá
  título gigante em miniatura minúscula. Tem que ser `transform: scale()` sobre largura fixa de
  720px, como `scores/poster/page.tsx:237` já faz.
- **A miniatura é a renderização do portal, não a moldura 1920×1080 do PDF.** Ela pode parecer
  perfeita com o slide estourando na impressão. Rotular como *"como o paciente vê na tela"*, nunca
  como prévia do slide. A prévia fiel é um `Dialog` sob demanda.
- **Manter a escotilha de JSON por slide.** A v1 dos formulários cobre menos que o JSON cobre
  (`sequence`, campos futuros); sem a escotilha a entrega é regressão de capacidade.

**Seis editores, não nove.** `kind` só seleciona quais blocos aparecem; os blocos é que têm editor:
`SlideHeaderFields`, `CardsEditor`, `TableEditor`, `RulersEditor`, `SummaryEditor`,
`TakeawayEditor` — sobre duas primitivas, `ListEditor<T>` (add/remove/reorder/teto) e `Field`
(label + hint + contador). `sequence` é read-only + JSON: kind morto, custa zero. **Não usar o
`rich-text-editor` (Tiptap)**: ele produz HTML de bloco que o deck não aceita e que apareceria
literal para o paciente; a barrinha certa tem três botões (`em`, `strong`, `br`).

**A régua no formulário**: `Popover` + `Command` no padrão de `LabTestDefinitionSelect.tsx`,
listando os `rulers[]` do dossiê congelado (só exame que o paciente tem). Ao escolher, `code`,
`unit`, `segments` e `history` vêm travados com tooltip de origem; `display`, `sub` e `note` são os
únicos editáveis; `axis` tem um "afinar" que mostra o valor do dossiê ao lado e um "voltar ao do
dossiê". Mais uma ação **"ressincronizar do dossiê"** por régua, preservando os campos autorais —
sem ela um rascunho de três dias republica número velho.

**Estouro em duas camadas.** Medir em debounce está errado: a medição é do deck inteiro e passa
pelo mutex do Chromium compartilhado com receita e pedido de exames — o editor poria a impressão da
clínica atrás de si.
- *Camada 1, cliente, 0 ms*: orçamento heurístico a partir das constantes reais de
  `pdfdoc/ruler.go:41-48` e dos tetos conhecidos, alimentando o contador de caracteres e o badge ao
  vivo. É estimativa e diz isso.
- *Camada 2, servidor*: automática **depois de um save bem-sucedido**, no botão explícito, e no
  publish (que já bloqueia com 422). Tetos duros na UI: ao chegar em 4 réguas o botão de adicionar
  desabilita com tooltip "4 é o teto comprovado; 8 estoura".

**A conversa e o aceite.** Resposta chega por POST-e-poll (2 s), espelhando `useSuggestedReply`.
A bolha traz prosa curta mais chips clicáveis: `✎ 3 slides alterados` · `⚠ 1 sugestão` ·
`desfazer tudo desta resposta`. A sugestão aparece **dentro do cartão do slide afetado**, com
antes → depois por campo e a origem do número ao lado do botão (decisão do usuário). Texto
auto-aplicado não vira sugestão: vira **marca de revisão** — ponto âmbar no campo e um
`Collapsible "Alterado pela IA (3)"` com desfazer por item.

**Diff sem biblioteca.** As ops são por campo com `before`/`after`, então `line-through` + seta é
honesto e suficiente. Só `punch`/`lede` longos justificam diff por palavra, e aí são ~30 linhas de
LCS. Para op estrutural o diff é a própria miniatura, lado a lado.

**Estado.** `useReducer` num `PlanDraftProvider`, não react-hook-form como dono do array (o dnd e o
patch da IA brigariam com o cache do RHF). Só o cartão expandido tem inputs vivos; os outros são
memoizados. O dirty-state por string do JSON inteiro sai: precisa ser **por slide**, com
`normalizeSlide()` espelhando o `omitempty` do Go — senão digitar num campo e apagar deixa o slide
sujo para sempre.

---

## Fases

Cada uma entrega sozinha. As três primeiras são a maior parte do trabalho e **não dependem do
modelo**; a quarta é fina se a terceira estiver certa.

| # | O quê | Entrega |
|---|---|---|
| 1 | `DeckSlide.ID` + `DeckSummaryLine.Code` + migration 00092 + `PatientPlanRevisionService` + `Update`/`Publish` transacionais + `expectedRevision` | Histórico e desfazer para o editor manual. Conserta o `published_content` sobrescrito |
| 2 | Dossiê congelado (`Freeze`/`Current`/`Staleness`/`RefreshDiff`) + separar as fontes da prévia + coluna esquerda + hook `usePlanDossier` | Prontuário compilado na tela; a tela para de pagar 28 queries por abertura |
| 3 | `ClassifyPath` + `ApplyOps` + `BuildNumericIndex` + `ExtractNumerals`, com testes contra o dossiê real de um paciente | **É a parte que decide se a feature é segura, e é testável sem uma chamada de API** |
| 4 | Cartões, 6 editores, 2 primitivas, dnd, orçamento de estouro, escotilha JSON | O editor sem JSON cru |
| 5 | `EditPatientPlan` + `SendMessage` + triagem + sugestões + resolve + painel de conversa | A discussão |
| 6 | Middleware de auditoria (resource mais específico) + `ai_touched_paths` no publish | "Esta devolutiva tem 4 frases geradas por IA que ninguém reescreveu" vira campo, não arqueologia |

**Degradação segura**: se a pilha de desfazer ou a triagem escorregarem do escopo, a IA vira 100%
sugestão na v1. Nunca o contrário.

## Verificação

```bash
docker compose exec -w /app api go build ./...
docker compose exec -w /app api go test ./internal/services/... ./internal/pdfdoc/... ./internal/models/...
docker compose exec -w /app api go run ./cmd/migrate up && docker compose exec -w /app api go run ./cmd/migrate status
pnpm generate    # DeckSlide.ID e DeckSummaryLine.Code precisam chegar ao api-types.ts
```

Por fase:

- **1** — editar o mesmo plano duas vezes e conferir que há duas revisões com `seq` crescente e
  autores corretos; publicar duas vezes e conferir que `is_publication` guarda os dois conteúdos;
  dois PUT concorrentes com o mesmo `expectedRevision` → o segundo é 409.
- **2** — congelar, alterar um exame do paciente no banco, e conferir que `dossierStale` acende e
  que o diff do refresh cita **só** os códigos usados no deck. Medir a prévia de novo: tem que
  cair de 1,97 MB para ~75 KB.
- **3** — testes de unidade sobre o dossiê real da Ana (está em prod, 83 réguas): `ExtractNumerals`
  em PT-BR (`1.023`, `1,023`, `112 g`, `12%`); `ClassifyPath` recusando `rulers[0].segments`;
  reescrita de `punch` sem número novo aplicando, e com número novo virando sugestão; `axis` que
  corta um ponto de `history` sendo rejeitado.
- **4** — reconstruir o deck de 20 slides da Ana só pela tela, sem tocar em JSON, e comparar com o
  PDF atual em `pacs/`. Medir a tela em 1366 de largura.
- **5** — QA no Playwright local: pedir "encurta o título do slide 3" (tem que auto-aplicar) e
  "atualiza a ferritina para o valor novo" (tem que virar sugestão com a origem à vista). Conferir
  `cache_read_input_tokens > 0` a partir do segundo turno.
- **6** — publicar e conferir `ai_touched_paths`; conferir que o audit log grava o id do plano, não
  só o do paciente.

Dados clínicos ficam em `pacs/`, que é gitignored. Nenhum deploy sem ordem explícita.

## O que não construir

- **`StepsEditor` para `sequence`** — kind morto nos dois decks reais.
- **Biblioteca de diff** — as ops já são por campo com antes/depois.
- **Preflight de estouro por slide na v1** — a heurística de cliente mais a medição no save cobrem
  o caso e não disputam o Chromium com a receita. Vale na v1.1.
- **Editor visual livre de slide** — a gramática fechada de nove blocos é o que segura a qualidade.
- **Otimizar as ~28 queries do dossiê agora** — congelar já tira isso do caminho quente. Os ~15
  Preloads de `GetLatestByPatientID` servem outra tela; mexer às cegas gera regressão lá. Fica como
  débito com nome: um `GetLatestSummaryByPatientID` sem Preloads, medido antes e depois.
- **Linha de divulgação de IA no deck** — decisão do usuário.
