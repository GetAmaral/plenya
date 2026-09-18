# Revisão integral do miolo — 17/09/2026

O editor parou na página 46 de 348 e pediu que o resto fosse visto com o mesmo olho.
Esta revisão faz duas coisas: (1) confere, no livro inteiro, se os oito itens corrigidos
em 26/08 continuam de pé fora do trecho que ele viu; (2) procura, nas 302 páginas que
ninguém revisou, os mesmos tipos de defeito que ele apontou.

Método: varredura geométrica do `Antes-pt-BR-brochura-miolo-cor.pdf` página a página
(script em `Editora/audit-miolo.py`, reexecutável), com conferência
visual de cada achado. As margens são espelhadas — recto 79,37→416,70 pt, verso
51,02→388,34 pt, medida de 337,3 pt nas duas — e a primeira versão do script tratou
tudo como recto, o que gerou 366 falsos positivos de justificação. Os números abaixo
são da segunda passagem.

As três edições (`-cor`, `-pb`, `-meio`) têm 348 páginas e a mesma paginação, então
toda correção vale para as três de uma vez.

---

## 1. Os oito itens de 26/08, conferidos nas 348 páginas

| # | Item | Resultado |
|---|------|-----------|
| 1 | Hífen em título | **0 ocorrências** em títulos, subtítulos, legendas e sumário |
| 2 | `VO□ max` | Subscritos renderizam (conferido em p.124); nenhum `.notdef` |
| 3 | Numeração desde a 1ª folha | **0 divergências**: em 348 páginas o fólio impresso é sempre o físico |
| 4 | Recuo no § depois de subtítulo | **0 faltas** |
| 5 | Capitular | **falha em 5 capítulos** — ver 2.1 |
| 6 | Justificação (corpo e legenda) | **0 faltas** nas duas |
| 7 | Figura isolada / página semivazia | **melhorou, mas sobrou** — ver 2.2 |
| 8 | Sumário | correto: 18 capítulos, 5 partes, fólios batendo |

Seis dos oito estão limpos de ponta a ponta. Os dois que sobraram são os itens 5 e 7, e
os dois falham fora do trecho que o editor leu.

Medidas gerais: 8.529 linhas de corpo, 336 terminando em hífen (**3,9%**); nenhum hífen
na virada de página; uma única escada de três hífens seguidos (p.248); nenhuma viúva e
nenhuma órfã (o template já tem `\widowpenalty=\clubpenalty=10000`). Os 56 "transbordos"
que o script marcou são protrusão do `microtype` (2,2 pt de aspas, 3,3 pt de travessão
na margem) — é margem óptica, está certo.

---

## 2. O que falta dos itens dele

### 2.1 Cinco capítulos ainda abrem sem capitular

| Cap. | Página | Abertura | Por que o filtro não disparou |
|---|---|---|---|
| 3 | 41 | *"Eu tenho 45 anos." Essa foi a primeira…* | § começa com aspas: o pandoc entrega `Quoted`, e o filtro exige `Str` |
| 8 | 139 | *O capítulo anterior tratou do primeiro motor…* | 1ª palavra é **"O"**: o filtro tem `if #word < 2 then return` |
| 9 | 161 | *A gestão clínica e metabólica é a parte…* | 1ª palavra é **"A"**: mesma linha |
| 13 | 227 | *Ricardo voltou ao consultório dezoito meses depois do infarto.* | § de 62 caracteres: `if total < 100 then return` |
| 17 | 309 | *Eu tenho a mesma idade biológica que meus pacientes.* | § de 52 caracteres: mesma linha |

Três causas distintas, todas em `brochura/print-dropcaps.lua`. As duas primeiras são bug
puro (uma abertura com aspas e uma abertura com artigo de uma letra não têm nada de
errado). A terceira é decisão de projeto: um § de uma linha não comporta capitular de
duas linhas sem que o rabo da letra desça por dentro do § seguinte.

Sumário, Referências e Agradecimentos seguem sem capitular de propósito (`SKIP_TITLES`).

### 2.2 Figura isolada: melhorou, mas três páginas ainda ficam com 1/3 a 2/3 em branco

Cinco figuras ocupam página inteira: **p.134, 205, 211, 290, 301**. Em três casos a
página de texto anterior é interrompida cedo:

| Página | Texto acaba em | Branco | O que vem depois |
|---|---|---|---|
| 204 | 10 linhas | **365 pt** (≈ 2/3 da página) | p.205, figura de página inteira |
| 210 | 18 linhas | 216 pt | p.211, figura de página inteira |
| 300 | 30 linhas | 141 pt | p.301, figura de página inteira |

É o mesmo desenho que ele descreveu em 09:22 e 09:32. Nas outras 14 figuras a distância
entre arte e texto está resolvida.

**E um caso pior, que não é de figura:** a **p.160** fecha o capítulo 8 com quatro linhas
(três da transição em itálico e a linha *"Pilar G — Gestão Clínica e Metabólica"*) e 490 pt
de branco embaixo. Uma página inteira gasta para duas frases.

---

## 3. Defeitos novos, da p.47 em diante, na mesma família

### 3.1 Toda figura é rotulada duas vezes, com números que não batem

Este é o mais visível dos achados e o editor teria batido nele na primeira arte que
encontrasse (p.77). A arte traz, dentro do próprio desenho, uma tarja preta **"FIGURA N"**
mais um título; logo abaixo, a legenda do LaTeX repete o título com **outro número**:

| Página | Legenda impressa | Rótulo dentro da arte |
|---|---|---|
| 154 | Figura **8.1**. Finasterida: quando sim, quando não. | FIGURA **2** — Finasterida: quando sim, quando não. |
| 158 | Figura **8.2**. Marcos, 8 meses depois… | FIGURA **1** — Marcos, 8 meses depois… |
| 195 | Figura **11.1**. Oito genes, oito decisões… | FIGURA **2** — Oito genes. Oito decisões… |
| 197 | Figura **11.2**. Paulo: 6 Meses Depois. | FIGURA **3** — Paulo: 6 meses sem reposição… |
| 223 | Figura **12.1**. Ana, dois marcadores duplos… | FIGURA **2** — Quando o pilar psicológico entra… |
| 236 | Figura **13.1**. O preço da solidão… | FIGURA **2** — O preço da solidão… |
| 244 | Figura **13.2**. Ikigai… | FIGURA **3** — Ikigai… |
| 261 | Figura **14.2**. A arquitetura de uma noite… | FIGURA **1** — A arquitetura de uma noite… |

São **21 figuras com tarja dentro da arte** contra **38 legendas de LaTeX**. Onde os dois
existem, o leitor vê o mesmo título duas vezes, e em oito casos com números diferentes.
O número da tarja é sequencial por capítulo mas não corresponde ao da legenda.

### 3.2 ~~A legenda impressa é o briefing de geração da figura~~ — retirado

Este item foi levantado por mim e **não procede**. As legendas vêm do texto alternativo
do markdown e eu li isso como sobra do pipeline de geração das figuras. O autor corrigiu
em 18/09: a legenda sempre esteve ali, é escrita de propósito e carrega conteúdo que o
desenho não dá sozinho (na 8.1, a regra de decisão, as alternativas com evidência e a
linha do tempo regulatória de 2011, 2022 e 2025). Fica como está.

Chegaram a ser apagadas num build intermediário, por uma decisão minha que não estava
autorizada; foram devolvidas ao estado original no mesmo dia.

### 3.3 Três legendas ficaram numa página, a arte na outra

| Arte | Legenda |
|---|---|
| p.205 (figura de página inteira) | p.206 — "Figura 12.1. Como a ansiedade vira doença…" |
| p.211 | p.212 — "Figura 12.3. Cinco instrumentos, cinco minutos…" |
| p.301 | p.302 — "Figura 16.1. Dois modelos de acompanhamento médico…" |

Essas três legendas são LaTeX cru (`\noindent{\sffamily\footnotesize\textbf{Figura …}}`)
escrito à mão no markdown, fora do `figure`, então nada garante que fiquem junto da arte.

### 3.4 Numeração do capítulo 12 quebrada

`Figura 12.1` aparece **duas vezes** (p.206 e p.223) e **`Figura 12.2` não existe** —
há 12.1, 12.3 e 12.1. A causa é a mistura do 3.3: as legendas escritas à mão não
incrementam o contador do LaTeX, e a figura flutuante de verdade do capítulo pegou o
número 12.1 que a legenda manual já tinha usado.

### 3.5 Onze páginas em branco saem com cabeço e fólio impressos

**p.18, 40, 72, 98, 170, 188, 200, 254, 278, 306, 314.** Cada uma tem *só* o cabeço
("40 · ANTES") e mais nada. São as folhas que o `\cleardoublepage` insere para o capítulo
abrir em página ímpar. A página de cortesia conta na numeração (é o que ele falou às
09:15) mas não leva fólio impresso nem cabeço. O `\cleardoublepage` do LaTeX base não
aplica `\thispagestyle{empty}`; quem resolve é o pacote `emptypage`.

### 3.6 Hífen repetido na virada de linha onde a regra do português não vale

O polyglossia repete o hífen no início da linha seguinte quando a palavra composta quebra
no próprio hífen. **Em português isso está certo** e acontece bem em seis lugares
(`compará-/-lo`, `pré-/-diabetes`, `hipotálamo-hipófise-/-adrenal`, `cardio-/-reno-metabólicos`,
`europeu-/-asiático`, `cognitivo-/-comportamental`).

Só que a mesma regra dispara dentro de DOI, de nome de ensaio clínico e de título em
inglês, onde não tem cabimento — **12 ocorrências**, quase todas nas Referências:

| Página | Sai impresso |
|---|---|
| 326 | `10.1186/s12958-023-` ⏎ `-01055-z.` (DOI com hífen duplo) |
| 322 | `OCEAN(a)-` ⏎ `-DOSE` (nome do ensaio partido) |
| 319, 327, 336 | `meta-` ⏎ `-analysis` (3×) |
| 336 | `double-` ⏎ `-blind` · `treatment-` ⏎ `-resistant` |
| 318 | `DXA-` ⏎ `-Derived` |
| 323 | `Low-` ⏎ `-Dose` |
| 155 | `off-` ⏎ `-label` |
| 107 | `diabetes in-` ⏎ `-situ` |

### 3.7 Rabo de palavra hifenizada sozinho na última linha do parágrafo

Dez ocorrências. As duas piores são a mesma frase, repetida:

- **p.120 e p.138**: *"…Atividade Física, Alimentação e Suplementação Inteli-"* ⏎ **"gente"**.
  A palavra "gente" sozinha, como última linha do parágrafo.
- p.330: *"…não-hormo-"* ⏎ **"nais:"**
- outras: p.82 `hiperten-/sivos`, p.233 `anticon-/vulsivantes`, p.271 `internacional-/mente`,
  p.321 `adicio-/nada`, p.331 `Gyneco-/logy`.

### 3.8 Quatro cabeços de recto saem com um "·" solto

**p.15, 17, 119, 277** imprimem "· 15" — o separador sem nada à esquerda. São as páginas
sem `\chaptermark`: a Introdução (`\chapter*`) e os textos de abertura das partes III e IV.

### 3.9 Cabeço do capítulo 4 é o dobro dos outros

`\chaptermark` corta o título no primeiro travessão. Os quatro capítulos cujo título usa
dois-pontos em vez de travessão (4, 6, 15, 17) escapam do corte e vão inteiros para o
cabeço. O do capítulo 4 ocupa 280 pt dos 337 pt da medida: *"O Painel Ampliado: O Mapa Que
o Check-up Básico Não Entrega · 69"*.

### 3.10 A arte da Figura 12.3 (p.211) tem texto sobreposto

Na escala AUDIT, *"Sobretudo em paciente que diz "só bebo socialmente""* atravessa a régua
vertical que separa a coluna da direita. Na escala UCLA-3, o "o" de *"pertencimento"* encosta
na seta de *"→ ≥ 6 → solidão relevante"*, e *"Mesmo em paciente que diz ter "muitos amigos""*
toca a borda da caixa. A arte também usa aspas retas (`"`) onde o miolo usa aspas curvas.

---

## 4. Continua pendente de decisão sua (do parecer de 26/08)

- **A. Travessão simétrico na capa.** Alternativa pronta em `capas/pt-BR/capa-travessao-simetrico.jpg`.
  `capa.jpg` serve EPUB, KDP, capa dura e brochura: troca nas quatro ou em nenhuma.
- **B. Espaço duplo no fim de frase.** Medido de novo agora: **4,76 pt depois de ponto
  contra 3,46 pt entre palavras — 1,38×**. Em português a convenção é espaço igual.
  Corrigir custa hífens (o espaço extra é a elasticidade que o TeX usa para fechar linha)
  e vale para as três edições.
- **C. Espessura do Pólen Bold 90 g.** A tabela da PoloPrinter dá 0,100 mm/folha (volume
  ~1,11), que contradiz o argumento de venda da Suzano. Se o volume real for 1,8, a lombada
  vai de 17,4 mm para ~28 mm. **Perguntar à gráfica antes de imprimir.**

---

## 5. O que já foi corrigido (mesmo dia)

O miolo saiu de **348 para 344 páginas** e as três variantes (`-cor`, `-pb`, `-meio`)
continuam com paginação idêntica.

### Figuras (3.1, 3.3, 3.4 e 2.2)

**O número saiu de dentro da arte.** Quem numera é o contador do LaTeX, que segue a
ordem de leitura. Dois scripts novos, ambos idempotentes (o original vai para
`_com-numero/` na primeira passagem e é sempre de lá que se lê, então rodar de novo não
empilha efeito e apagar a pasta restaura):

- `build/versaoImpressa/strip-figure-numbers.py` — 62 PDFs vetoriais (31 em
  `figuras-cor/`, 31 em `figuras-bw/`). Nas 44 de tarja, redação do texto e do chip
  colorido atrás dele. Nas 16 de número soldado no título (`Figura 1 — Os 20 Anos
  Silenciosos`), redação da linha inteira e o título redesenhado sem o prefixo, no mesmo
  ponto de origem, no mesmo Inter-Bold, tamanho e cor. Cobrir só o prefixo deixaria o
  título 90 pt adentro, porque esses títulos são alinhados à esquerda.
- `build/versaoImpressa/strip-figure-numbers-raster.py` — as 4 que só existem em raster
  (mais as 4 cópias B&W). Sem texto para redigir: a tarja é repintada com o fundo
  amostrado **nas mesmas linhas**, à direita (amostrar acima/abaixo deixava um retângulo
  mais claro, porque o cabeçalho fica sobre um painel cinza que começa dentro da faixa);
  e no caso de número soldado o bloco de pixels do título é recortado e colado na margem
  esquerda, o que evita ter de adivinhar a fonte.

**A legenda virou `\caption` de verdade — com o texto que sempre teve.** O tratamento de
página inteira em `build-brochura-miolo.py` foi reescrito: a legenda fica DENTRO do
float, e os dois `\clearpage` saíram. Isso resolveu de uma vez a legenda que caía na
página seguinte à arte, a `Figura 12.1` duplicada com a `12.2` faltando, e as páginas de
texto cortadas no meio (p.204 tinha 10 linhas e dois terços em branco). A tabela nativa
do Cap. 4 passou de número escrito à mão para `\captionof`, mantendo o texto original.

Resultado: **38 legendas, com o texto de sempre, numeração contínua por capítulo, sem
duplicata e sem buraco**; figuras de página inteira caíram de 5 para 3.

### Template (3.5, 3.6, 3.7, 3.8, 3.9 e 2.1)

| Item | Correção | Antes → depois |
|---|---|---|
| 3.5 Página em branco com cabeço | `\usepackage{emptypage}` | 11 → **0** |
| 3.8 Cabeço com “·” solto | `\plenyaheadmark` testa se `\leftmark` está vazio | 4 → **0** |
| 3.9 Cabeço longo do cap. 4 | `\chaptermark` corta no travessão **e** no dois-pontos | 280 pt → 223 pt (o maior agora é o do cap. 8) |
| 2.1 Capitular | `print-dropcaps.lua`: desmonta o `Quoted` e joga as aspas no `ante`; aceita primeira palavra de uma letra; § de abertura curto passa a vez para o § seguinte | 5 capítulos sem → **0** |
| 3.6 Hífen repetido | bloco `{=latex}` no topo de `18-referencias-recursos.md` solta o hífen só ali (`\XeTeXcharclass` + `\exhyphenpenalty=2000`) | 18 → **9**, e os 9 são compostos portugueses, onde repetir é a regra |
| 3.7 Rabo de palavra | `brochura/linha-de-destaque.lua` trata `***Pilar X — …***` como título e não hifeniza; `\finalhyphendemerits` de 10000 para 100000 | 11 → **8**, e “Inteli-/gente” saiu |

Desligar o `splithyphens` do polyglossia globalmente foi tentado e descartado: sem o
`\discretionary` o `\exhyphenpenalty=10000` proíbe qualquer quebra em hífen existente, e
um DOI passou a furar 34 pt para fora da margem. O ajuste local nas Referências resolve
os 12 casos errados sem tirar a regra do português do resto do livro.

Hifenização do corpo: 3,9% → **3,8%**. Transbordos acima de 5 pt: **0** (os 52 que a
varredura ainda marca são protrusão do `microtype`, 2,2 pt de aspas e 3,3 pt de
travessão na margem óptica).

### O que isto faz nas outras edições

As artes são compartilhadas, então o número saiu delas também para a KDP 6×9, a capa dura
e o EPUB — o que nessas edições já é ganho: some o conflito entre o número da tarja e o
da legenda. Mas duas coisas ficam pela metade lá, porque os filtros novos estão só no
build da brochura:

- o número das figuras de página inteira dessas edições continua escrito à mão em
  `FULLPAGE_FIGURES`, com o mesmo risco de colidir com o contador;
- `versaoImpressa/tabelas-nativas/cap04-fig02.tex` é compartilhado e passou a usar
  `\captionof`; o texto da legenda é o mesmo, mas o número passa a vir do contador
  também lá quando essas edições forem rebuildadas.

São dois `--lua-filter` e o mesmo reescrito de `fullpage_figures()` em
`versaoImpressa/build-print-pdf.py`. Não foram aplicados aqui porque essas edições não
foram rebuildadas nem conferidas nesta rodada.

### O que continua aberto

- **Os capítulos 13 e 17 abrem com uma linha de entrada.** Baixar a capitular para duas
  linhas foi tentado e descartado: num § de uma linha o rabo da letra desce por dentro do
  § seguinte e come o recuo dele. A capitular passa para o primeiro § que comporta, e a
  linha curta ("Ricardo voltou ao consultório dezoito meses depois do infarto.") fica por
  cima, como deixa. Vale conferir na prova se é o efeito que o autor quer.
- **A `Figura 12.2` e a `12.3` trocaram de número.** No markdown do cap. 12 a arte
  `Cap12_Fig03` aparece ANTES da `Cap12_Fig02`, e agora quem numera é a ordem de leitura:
  a arte feita como "12.3" imprime 12.2, e vice-versa. Não quebra nenhuma citação (não há
  citação de figura no livro), mas enquanto as outras edições não receberem o mesmo
  tratamento, essas duas figuras terão números diferentes em edições diferentes.
- **Nenhuma figura é citada no corpo.** As 38 legendas trazem número, mas o texto nunca
  diz "ver Figura 8.1". Não é defeito — a legenda se sustenta sozinha —, mas se em algum
  momento o autor quiser amarrar figura e texto, a numeração agora é confiável para isso.
- **O título aparece duas vezes em 21 figuras** (dentro da arte e no começo da legenda).
  Foi o item 3.1; com as legendas mantidas como estão, a repetição continua. Resolver
  significa encurtar a abertura da legenda, e isso é escrita, não diagramação.
- **Quatro capítulos terminam com três ou quatro linhas numa página quase vazia**
  (p.268, 288, 195, 299). Não há ajuste global para isso: é `\looseness` caso a caso, na
  prova.
- A arte da Figura 12.3 (p.206) continua com texto sobreposto à régua (3.10).
- As três decisões antigas: capa, espaço duplo, Pólen Bold.

## 6. Próximo passo

Com os itens acima corrigidos, a sugestão do editor volta a ser a certa: imprimir uma
prova física e marcar no papel. As páginas que terminam curtas e o texto sobreposto na
arte da 12.3 são exatamente o tipo de coisa que se resolve com a prova na mão.

A varredura é reexecutável a cada build:

    build/.venv-figures/bin/python3 Editora/audit-miolo.py \
        build/brochura/Antes-pt-BR-brochura-miolo-cor.pdf

---

## Escopo: só a brochura

Decidido em 18/09: **as outras edições ficam fora desta rodada.** EPUB, KDP 6×9 e capa
dura seguem como estão — com `capa.jpg`, com a legenda de figura vinda do texto
alternativo e com o mesmo esticamento de 11,3% na 1ª capa que foi corrigido aqui
(`build-print-cover.py`, `build-print-cover-en.py`, `capa-dura/build-hardcover-cover.py`).
As menções a elas neste arquivo ficam como registro do que existe, não como pendência.

O que muda nelas de qualquer jeito, porque os arquivos são compartilhados: as artes das
figuras (o número saiu de dentro do desenho) e `versaoImpressa/tabelas-nativas/cap04-fig02.tex`
(o número da legenda passou a vir do contador).
