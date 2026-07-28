# Escore — por que dá 93% em quem não está bem, e como apresentar na devolutiva

**Data:** 2026-07-28 · **Estado:** parcialmente aplicado no dev (migration 00064)
**Caso-âncora:** paciente real atendido em 28/07/2026, `Continuum | Médico | Inicial`.
Dados em prod, snapshot `019fa94b-185e-7895-8681-a79e54603264`. Identificação fora do repo.

---

## 1. O que o snapshot do caso-âncora diz

| | |
|---|---|
| Escore global | **93,03%** (2924 / 3143 pontos) |
| Itens avaliados | 205 (mas só **125 distintos** — ver bug §2.1) |
| Itens não avaliados | 1302 |
| Cobertura real | 125 de 1507 itens do escore = **8,3%** |
| Exames | 0 pontos possíveis (sem laboratório) |
| Genética | 0 pontos possíveis |

Por grupo (já corrigindo a duplicação):

| Grupo | Itens | Pontos | Máx | % |
|---|---:|---:|---:|---:|
| Histórico de doenças | 79 | 1394 | 1394 | **100,0** |
| Composição corporal | 2 | 28,4 | 32 | 88,8 |
| Histórico Familiar | 9 | 127,4 | 147 | 86,7 |
| Social | 4 | 28,2 | 33 | 85,5 |
| Sono | 9 | 82,4 | 112 | 73,6 |
| Alimentação | 7 | 18,8 | 26 | 72,3 |
| Stress | 3 | 24,8 | 41 | 60,5 |
| Objetivos | 1 | 8,4 | 14 | 60,0 |
| Cognição | 6 | 43,6 | 74 | 58,9 |
| Vida Sexual | 1 | 3,2 | 8 | 40,0 |
| Movimento | 4 | 3,2 | 44 | **7,3** |
| **Total** | **125** | **1762,4** | **1925** | **91,6** |

Distribuição dos níveis casados: **101 itens em N5**, 2 em N4, 1 em N3, 9 em N2, 3 em N1,
9 em N0. Ou seja, 81% das respostas foram "está tudo bem" — e 72% do denominador veio de um
único grupo que só pergunta se o paciente **tem ou não tem** doença.

O paciente que o Dr. viu no consultório aparece corretamente nos dados: ronca, é sedentário,
falhou nas duas fases do Dubois, dorme mal, tem stress, toma refrigerante. Só que esses 21 itens
ruins somam **162,6 pontos perdidos** contra **1925 possíveis** — 8% do total. A matemática está
funcionando; a **estrutura de pesos** é que está errada.

---

## 2. As cinco causas

### 2.1 BUG — cada item-filho é contado duas vezes

`ScoreSnapshotService.CalculateSnapshot` (`score_snapshot_service.go:150-202`) percorre
`subgroup.Items` **e depois** `item.ChildItems`. Mas `GetAllScoreGroupTrees`
(`score_repository.go:122`) faz `Preload("Subgroups.Items")` **sem filtrar
`parent_item_id IS NULL`** — então `Subgroups.Items` já traz os filhos. Todo item que tem pai
entra duas vezes no snapshot.

No caso-âncora: 205 linhas para 125 itens — 80 filhos duplicados.

Não é discussão de modelagem: **todos os outros três pontos do código que fazem esse mesmo
laço já se protegem com um `seen map[uuid.UUID]bool`** — `anonymous_score_service.go:932`,
`anonymous_score_service.go:1053`, `score_version_service.go:216`. Só o snapshot do paciente
ficou sem. É o único caminho de escore do EMR que erra.

Efeito no caso-âncora: 93,03% → **91,6%**. Pequeno em cima, mas distorce o peso relativo de qualquer
subgrupo que use hierarquia pai→filho (Medicamentos, Doenças crônicas, Consumo de alimentos,
Líquidos no dia…). **Correção obrigatória e independente de qualquer decisão de calibração.**

### 2.2 O checklist de ausência domina o denominador

O grupo "Histórico de doenças" é uma varredura de **presença/ausência**:

| Subgrupo | Itens | Pontos |
|---|---:|---:|
| Doenças crônicas | 25 | 505 |
| Medicamentos | 30 | 503 |
| Cirurgias já realizadas | 19 | 282 |
| Hábitos e vícios | 5 | 104 |

Itens como `Colecistectomia` e `Anti-inflamatórios (AINEs)…` têm exatamente dois níveis —
N0 "Sim", N5 "Não" — e valem 12 a 22 pontos cada. Ou seja: **"não fiz colecistectomia" vale
14 pontos, o mesmo que "qualidade percebida do sono"**. Quem responde "não" 79 vezes leva
1394 pontos de graça, 72,4% de todo o denominador, com 100% de aproveitamento.

Sem esse grupo, o caso-âncora cai para **69,4%** — que é o número que o Dr. reconheceria como justo.

O problema conceitual: **ausência de doença é pré-requisito, não conquista.** Um escore de
longevidade deveria premiar o que o paciente *faz*, não o que ele *ainda não teve*.

### 2.3 Os domínios modificáveis não têm peso

`score_items.points` NULL ou 0 significa: o item é perguntado, é classificado, aparece na
anamnese — e **não entra na conta**. Por grupo:

| Grupo | Itens | Sem peso | Peso total do grupo |
|---|---:|---:|---:|
| Exames | 327 | 26 | 4442 |
| Histórico de doenças | 182 | 37 | 2023 |
| Genética | 361 | 0 | 851 |
| Composição corporal | 59 | **20** | 529 |
| Movimento | 63 | **20** | 404 |
| Sono | 53 | 12 | 401 |
| Cognição | 38 | 8 | 353 |
| **Alimentação** | 58 | **37** | **236** |
| Histórico Familiar | 24 | 6 | 211 |
| Vida Sexual | 20 | 7 | 142 |
| Social | 23 | 6 | 123 |
| Stress | 6 | 1 | 69 |
| **Objetivos** | 19 | **18** | **14** |

Alimentação inteira pesa 236 pontos — menos do que "Medicamentos" sozinho (523). No snapshot
do caso-âncora, `Consumo de Frutas`, `Consumo de Açúcar`, `Água` e `Álcool` foram avaliados **em N0 e
N1** e contribuíram com `max_points = 0`: as respostas ruins são literalmente invisíveis.
`Café e chás` também está com `points = 0` (relevante para o ponto 1 desta mesma rodada — não
adianta passar a perguntar café no Inicial se ele não pontua).

### 2.4 A curva de níveis é generosa e o topo é chato

`GetLevelMultiplier` é linear: N5=100%, N4=80%, N3=60%, N2=40%, N1=20%, N0=0%. Duas
consequências:

- **N3 ("mediano") entrega 60%.** Numa escala em que o alvo é "ótimo", mediano não deveria
  valer mais que ~45%.
- **N5 e N6 valem os dois 100%.** O nível de excelência não se distingue do "normal".

### 2.5 O total ignora o desenho por pilar — e não sinaliza incompletude

O total é uma soma bruta de pontos entre grupos, então o **tamanho** do grupo é o peso dele.
Comparando as agregações possíveis no mesmo snapshot do caso-âncora:

| Modelo | Resultado |
|---|---:|
| Atual (com duplicação) | 93,0 |
| Só corrigindo a duplicação | 91,6 |
| Média macro dos 36 pilares avaliados | 86,3 |
| Média das 4 letras AGIR (A 74,2 · G 96,3 · I 74,0 · R 76,5) | 80,2 |
| Macro-macro: item→pilar→letra→total | **79,0** |
| Idem, com curva não-linear (§3.3) | **76,1** |
| Só domínios modificáveis (sem checklist de ausência) | 69,4 |

E, em qualquer um desses, **o headline não conta que só 8,3% do escore foi medido**. O caso-âncora não
tem um único exame; o 93% é apresentado com a mesma cara de um paciente com painel completo.

---

## 3. Proposta de recalibração

Cinco frentes, em ordem de segurança. As três primeiras eu recomendo fechar já; as duas últimas
são decisão clínica do Dr.

### 3.1 (Bug) Deduplicar o snapshot — obrigatório

Duas opções, aplicar as duas:

1. `score_repository.go`: filtrar `Preload("Subgroups.Items", … .Where("parent_item_id IS NULL"))`
   — alinha a árvore com o que o nome promete.
2. `score_snapshot_service.go`: `seen map[uuid.UUID]bool` no laço, igual aos outros três call
   sites.

Depois, recalcular os snapshots existentes. Sem release de front.

### 3.2 Rebaixar o checklist de ausência a "gate", não a "pontos"

Recomendação: **o subgrupo de presença/ausência deixa de somar pontos positivos e passa a ser
um redutor.** Concretamente:

- Itens binários N0/N5 de doença, cirurgia e medicamento **saem do denominador** quando a
  resposta é "não" (contribuem 0/0, exatamente como hoje já acontece com item sem dado).
- Quando a resposta é "sim", o item entra com peso cheio no pilar clínico correspondente
  (Cardiovascular, Renal, Metabólico…), puxando aquele pilar para baixo.

Assim o checklist vira o que ele clinicamente é: uma lista de **passivos**. Quem não tem
nenhum não ganha nada por isso; quem tem, perde. O caso-âncora iria para ~69–70% pela porta da frente,
e um paciente com 4 doenças crônicas cairia de verdade.

Alternativa mais conservadora, se o Dr. preferir manter pontos positivos pela ausência:
**teto de contribuição** — o bloco de ausência entra com peso fixo de 25% do total,
independentemente de ter 79 ou 200 itens. Caso-âncora: **77,0%** (74,1% com a curva não-linear).

### 3.3 Curva de níveis não-linear

| Nível | Hoje | Proposta |
|---|---:|---:|
| N6 (excelente) | 100% | 100% |
| N5 (ótimo) | 100% | 92% |
| N4 (bom) | 80% | 72% |
| N3 (mediano) | 60% | 45% |
| N2 (ruim) | 40% | 25% |
| N1 (muito ruim) | 20% | 10% |
| N0 (crítico) | 0% | 0% |

Efeito isolado no caso-âncora: 91,6 → 90,5 (pequeno, porque ele é quase todo N5). O ganho aparece no
paciente mediano, que hoje sai com ~70% tendo respondido "mais ou menos" em tudo — com a curva
nova sai com ~50%. É o que faz o escore **discriminar**.

`GetLevelMultiplier` está em `score_snapshot_service.go:465`; a mesma tabela precisa ir para o
`anonymous_score_service` e para o `matchLevel` do TS, senão escore-light e EMR divergem.

### 3.4 Dar peso aos domínios modificáveis

Passar de "pontos por item" para **orçamento por pilar**: cada pilar AGIR tem um peso-alvo, e
os itens dentro dele dividem esse peso. Consequências:

- Alimentação, Movimento, Sono e Stress param de valer menos que a lista de cirurgias.
- Adicionar um item novo a um pilar **não** infla o pilar (hoje infla) — só redivide.
- Some a classe de bug "item avaliado com `points = 0`" (Café, Água, Álcool, Frutas, Açúcar).

Passo mínimo enquanto isso não existe: atribuir peso aos ~37 itens de Alimentação, 20 de
Movimento e 20 de Composição que hoje estão zerados. É uma curadoria de dados, não código.

### 3.5 Separar "quão bem" de "quanto se sabe"

Dois números, nunca um só:

- **Escore** — a nota, calculada só sobre o que foi medido (como hoje).
- **Completude** — `itens avaliados / itens aplicáveis do plano`. Caso-âncora: 8,3%.

Na tela e no PDF, escore com cobertura baixa aparece com faixa de incerteza e a frase de que
é parcial. Isso resolve o desconforto de mostrar 93% para quem não fez um exame sequer, sem
precisar mexer na matemática — e cria a alavanca comercial natural ("seu escore está
incompleto; faltam os exames").

### 3.6 Simulação combinada

Aplicando 3.1 + 3.2 (versão "gate") + 3.3 + o desenho por pilar:

| Paciente | Hoje | Proposto |
|---|---:|---:|
| Caso-âncora (real) | 93,0% | **~66–70%** |
| Caso-âncora com o teto de 25% em vez do gate | 93,0% | **~74%** |

Recomendo mirar a faixa **65–75% para esse perfil**: alto o bastante para não desmotivar, baixo o
bastante para a conversa "olha onde dá para melhorar" fazer sentido.

---

## 4. Devolutiva visual — radar + drill por pilar e subpilar

Hoje a `/health-scores/[id]` tem o radar AGIR e um acordeão de 4 níveis
(letra → pilar → grupo → subgrupo → itens). O acordeão é ferramenta de auditoria, não de
consulta: para mostrar ao paciente é preciso abrir cinco caixas até chegar num item.

Proposta: um **modo Devolutiva** na mesma página (toggle "Apresentar"), tela cheia, fonte
grande, navegável com setas — cinco telas.

### Tela 1 — Panorama
Radar AGIR atual (já existe) + escore global + **medidor de completude**. Nada mais.

### Tela 2 — Mapa de pilares
Grid dos 43 pilares como tiles, agrupados pelas 4 letras e coloridos por faixa
(crítico / atenção / bom / ótimo). Pilar sem dado fica hachurado, não colorido — hoje o
`buildAgir` deixa pilar com `max = 0` aparecer como **0%**, que lê como "péssimo" quando é
"sem peso" (aconteceu com "Massa Muscular e Hidratação Celular" no caso-âncora). Bug a corrigir junto.

O paciente vê num relance onde estão os buracos. No caso-âncora: Atividade Física (7,3%) acende
vermelho sozinho, cercado de verde.

### Tela 3 — Foco: os 5 maiores ganhos
Ranking por **pontos deixados na mesa** (`max - actual`), que é a métrica que responde
"onde eu ganho mais mexendo". No caso-âncora sai exatamente a conversa clínica certa:

| Perda | Nível | Item | Pilar |
|---:|---|---|---|
| 14,0 | N0 | Roncos | Distúrbios do Sono |
| 14,0 | N0 | Estratégia macro atual | Atividade Física |
| 14,0 | N0 | 5 palavras de Dubois — tardio | Função Cognitiva |
| 14,0 | N0 | 5 palavras de Dubois — imediato | Função Cognitiva |
| 12,0 | N0 | Divisão das atividades | Atividade Física |
| 11,2 | N1 | Obesidade (familiar) | Histórico Familiar |
| 10,0 | N0 | Lesões relacionadas ao exercício | Atividade Física |

Cada linha já tem texto pronto no banco: `score_levels.patient_explanation` e
`score_levels.conduct` do nível casado. A tela é montagem, não redação.

### Tela 4 — Um pilar por vez
Escolhido um pilar (do radar ou do mapa), abre uma tela só dele:
barra de faixa (crítico→ótimo) com a posição do paciente, os subpilares como barras
horizontais ordenadas do pior para o melhor, e os itens como chips coloridos por nível.
Clicar no chip abre a explicação do nível.

Isso é o "detalhe visual de cada pilar e subpilar" pedido — o dado já existe
(`score_item_method_pillars` → pilar; `score_groups`/`score_subgroups` → subpilar).

### Tela 5 — Plano
Os 3 a 5 focos escolhidos na tela 3 viram itens do plano, com o `conduct` de cada nível.
Sai como PDF pelo pipeline de assinatura já existente.

### Correções de front que entram junto

1. `buildAgir` (`packages/ui/src/score-radar/build-agir.ts:60`): pilar com `max = 0` deve ser
   "sem dado", não `score: 0`.
2. Faixas de cor precisam ser as mesmas em radar, mapa e chips — hoje o `LEVEL_STYLES` da
   página usa uma escala e o radar usa a paleta de marca por letra.
3. O modo Devolutiva reusa `RadarAgir`; nada de segundo componente de radar.

---

## 4b. O que já foi aplicado no dev (2026-07-28)

Migration `00064_escore_equilibrio_cafe_aines_alimentacao.sql` + fix do laço em
`score_snapshot_service.go`:

| Frente | Estado |
|---|---|
| §3.1 dedup do snapshot | ✅ `seen map[uuid.UUID]bool` no laço (121 linhas = 121 itens) |
| §3.4 peso na alimentação | ✅ 8 itens `level_choice` zerados → 10 pts; grupo 236 → 326 |
| §3.2 (versão parcial) histórico de doenças | ✅ peso de cada item dividido por 2; grupo 2023 → 1023,5 |
| Café | ✅ 10 pts, migrou de Nutri Inicial para Médico Inicial |
| AINEs | ✅ item próprio, 12 pts, 4 níveis por frequência; analgésicos ficam com 6 pts |
| Modo Devolutiva | ✅ protótipo em `components/health-scores/ScoreDevolutiva.tsx` |

Efeito medido no mesmo padrão de respostas do caso-âncora: **93,0% → 83,9%**.
Alimentação saiu de 72,3% para 40,5% (as respostas ruins agora custam). Histórico de doenças
caiu de 72% para 53% do denominador.

Pendências abertas para chegar na faixa 65-75%: §3.2 na versão "gate" e §3.3 (curva não-linear).

### Segunda rodada (migration 00065)

1. ✅ **`CETOGENICA` zerado.** Valia 8 pontos com níveis "Não"=N0 / "Sim"=N5, ou seja, quem não
   faz dieta cetogênica perdia 8 pontos. Agora acompanha os outros 10 booleanos de padrão
   alimentar (Livre, Vegana, Carnívora…), todos em 0.
2. ✅ **`patient_explanation` + `conduct` preenchidos nos 569 níveis** dos itens do
   Continuum | Médico | Inicial (antes: 569 de 569 vazios). Texto autoral nos itens de
   comportamento e estilo de vida; texto derivado por regra nos checklists de presença/ausência
   de doença, cirurgia e medicamento. A migration só grava onde o campo está vazio.
   **Conteúdo de primeira passada, precisa da revisão do Dr. Getúlio antes de prod.**
3. ⚠️ **Nomes de nível técnicos permanecem** — o Dubois ainda mostra "≤8" e "≤7" como rótulo.
   Contornado na tela: a Devolutiva passou a exibir o `patient_explanation` embaixo do rótulo,
   então o paciente lê a frase e não o corte numérico. Renomear os níveis continua desejável.
4. Restam **3.987 níveis vazios** fora desse template (Exames 1.593, Genética 1.041, e o resto
   do Histórico de doenças). São os que o pipeline de enrichment (`scripts/enrichment/`,
   `cmd/enrich-score-items`) foi feito para cobrir.

### Terceira rodada — revisão da exibição

Sete problemas encontrados relendo as telas, todos corrigidos em
`components/health-scores/ScoreDevolutiva.tsx`:

1. **Faixa de cor generosa demais.** Com o corte antigo em 80, o caso-âncora (sedentário,
   Dubois alterado, ronco) era rotulado **"Ótimo"** — exatamente o erro que a devolutiva
   existe para não cometer. Cortes novos: 85 Ótimo · 70 Bom · 50 Atenção · abaixo disso
   Prioridade. O mesmo paciente agora lê "Bom", com a ressalva de resultado parcial.
2. **Item isolado colorido pela razão de pontos.** "Insônia: Raramente" (nível 2) saía
   **vermelho**, porque nível 2 vale 40% e 40% caía na faixa de prioridade. Item discreto
   passa a ser colorido pelo **nível**, não pela fração de pontos.
3. **Histórico familiar entrava no ranking de focos.** "Obesidade em 2 parentes" aparecia como
   foco de plano, e o paciente não muda a família dele. Histórico familiar, genética, cirurgias
   já feitas e histórico de infância viraram uma seção separada, **"Contexto, não meta"**.
4. **Pilar sem dado não aparecia em lugar nenhum.** O paciente via 37 pilares e não sabia que
   existem outros esperando exame. Nova seção **"Ainda não medidos"** no mapa, que é também a
   ponte natural para pedir os exames.
5. **Nota alta sustentada por um item só parecia igual a nota alta com 19 itens.** Cada tile
   agora mostra `x de y itens medidos`, e marca **"base pequena"** quando dois ou menos itens
   sustentam a nota (ex.: "Tolerância Alimentar 100%" com 1 item).
6. **Régua de faixa sem marcador legível** e **chips de item sem o nível visível** (só no
   hover, inútil em consulta). Marcador com rótulo, e o nível casado aparece em cada item.
   Subpilar passou a ser rotulado `Grupo · Subgrupo` — "Atual" sozinho não diz nada.
7. **Faltava navegação de apresentação.** Setas ←/→ trocam de passo e Esc sai. O seletor de
   pilar, que era um paredão de 37 chips, ficou agrupado por letra.

Bônus fora do componente: `buildAgir` (`packages/ui/src/score-radar/build-agir.ts`) filtrava
só por `filled > 0`, então pilar com peso zero entrava no radar como **0%**, que lê "péssimo"
quando é "sem peso medido". Passou a exigir `max > 0`. Corrige o radar em todas as telas,
não só na devolutiva.

## 5. O que decidir

1. **§3.2** — checklist de ausência vira gate (recomendado) ou teto de 25%?
2. **§3.3** — aprova a curva não-linear proposta ou prefere outra?
3. **§3.4** — vale investir no orçamento por pilar, ou só curar os pesos zerados por enquanto?
4. **§3.5** — mostrar completude junto do escore em todas as superfícies (EMR, portal, PDF)?
5. **§4** — modo Devolutiva na `/health-scores/[id]` ou tela separada?

Independente disso: **§3.1 (dedup) e o bug do pilar sem peso no radar podem ir já.**
