# Parecer do editor — 26/08/2026 (brochura, miolo em cor)

Quinze áudios e quatro fotos de tela, entregues em `Editora/`. Os áudios foram
transcritos com `gpt-4o-transcribe` (transcrições cruas em `Editora/transcricoes/`);
onde a transcrição saiu confusa, o sentido foi reconstruído pela foto que o editor
mandou no mesmo minuto.

O editor abriu `Antes-pt-BR-brochura-miolo-cor.pdf` e parou na página 46. Ele mesmo
avisa que não passou de lá: *"é um livro bem grande, se eu fosse passar assim e
continuar, eu falaria bastante coisa"*. Então **esta lista não é o parecer completo** —
é o que aparece nos primeiros dois capítulos e meio. Ele sugere imprimir uma prova
física para marcar tudo de uma vez.

> Ele pede duas vezes que o tom não seja lido como crítica: *"não estou criticando o
> teu trabalho... é só alguma coisa que pode melhorar, não é no conteúdo, longe disso"*.

---

## O que já foi corrigido

| # | Apontamento | O que era | Correção |
|---|---|---|---|
| 1 | Título de seção quebrado com hífen (foto 09:23 e 09:33: *"Estão Fa-/lhando"*, *"Não Mos-/tra"*) | 18 subtítulos hifenizados. O ajuste de hifenização de 23/08 mexeu só no corpo; título nunca foi tocado | `\notitlehyphen` em `\part`, `\chapter`, `\section`, `\subsection`, legenda e sumário. Zero hífens de título |
| 2 | `VO□ max` (foto 09:30) | TeX Gyre Pagella não tem os glifos U+2080–U+2089. O XeLaTeX imprimia o `.notdef` box, 31 vezes | `newunicodechar` mapeia `₀..₉` para `\textsubscript` na própria Pagella |
| 3 | *"a numeração do livro é contada desde a primeira página, mesmo que ela seja em branco. Eu vi que não está contando aqui"* | `\frontmatter` romano + `\mainmatter` reiniciando em 1: a Introdução imprimia "1" na 13ª folha física | Contagem contínua. A Introdução agora imprime 13, o Capítulo 1 abre em 21. O pré-textual segue sem fólio impresso |
| 4 | *"faltou um parágrafo ali"* (páginas 2, 4, 5, e de novo no 3.4) | Primeiro parágrafo depois de subtítulo saía sem recuo | `\titlespacing` sem estrela em `\section`/`\subsection`. No `\chapter` continua com estrela: lá quem abre é a capitular |
| 5 | *"talvez a tua intenção era colocar uma letra inicial maior, acho que fica bonito; talvez tenha faltado ali"* | Capitular só nos capítulos numerados. A Introdução chega como `\chapter*` cru e o filtro Lua nunca disparava | Cópia local do filtro (`brochura/print-dropcaps.lua`) com hook de `RawBlock`. Parágrafo de abertura curto agora cai para capitular de 2 linhas em vez de sumir |
| 6 | *"página 32, na parte de baixo, o texto não ficou justificado"* / *"as notas das figuras estão alinhadas à esquerda"* | `justification=raggedright` no `captionsetup` | Legendas justificadas |
| 7 | *"a figura ficou totalmente isolada porque se refere ao texto da página 36"* / *"não ficaria esse desenho melhor no final?"* | Frações de float do `book` (topfraction 0.7, textfraction 0.2) empurravam figura de meia página para uma página de float só dela, com dois terços em branco, duas páginas depois da chamada | Frações afrouxadas + `\floatplacement{figure}{htbp}`. Páginas sem nenhum texto de corpo: **43 → 28**. O miolo encolheu de 352 para 348 páginas |
| 8 | (não apontado, mesmo defeito do #6) | Sumário justificado: título de duas linhas saía com a primeira esticada de margem a margem | `\raggedright` no `titlecontents`, com preenchimento `filll` para o número continuar colado na margem |

Hifenização do corpo, de quebra: **8,2% → 3,4%** das linhas terminando em hífen.
Também se descobriu que `\lefthyphenmin`/`\righthyphenmin` postos no preâmbulo eram
apagados pelo polyglossia no `\begin{document}` — agora ficam depois dele, e não há
mais corte de duas letras (`in-`, `as-`).

---

## Aguardando decisão

### A. Travessão simétrico na capa

> *"na tua capa, 'a janela silenciosa entre o normal e o ótimo', aí você colocou um
> travessão ali, 'onde a saúde é decidida'. Esteticamente, se tivesse um hífen na
> frente e outro no final daquela linha, ficaria mais bonitinho. Fica mais
> centralizada; aquele travessão da esquerda parece que fica um pouco perdido na arte."*

Feito como alternativa em `capas/pt-BR/capa-travessao-simetrico.jpg`. O travessão da
direita é o da esquerda espelhado (mesmos pixels), então a fonte casa exatamente.

**`capa.jpg` é usada por todas as edições** (EPUB, KDP, capa dura e brochura). Trocar
só na brochura deixaria a capa da brochura diferente das outras. Decisão: troca em
todas ou nenhuma.

### B. Espaço duplo no fim de frase

Não foi apontado pelo editor, mas apareceu na verificação. O `\frenchspacing` do
template **não pega**: o polyglossia restaura o `\nonfrenchspacing` dentro do
`\begin{document}`, e o miolo sai com 4,33 pt depois de ponto contra 3,27 pt entre
palavras. Em português a convenção é espaço igual.

Corrigir muda o cinza de todas as páginas e custa hífens (335 → 424), porque aquele
espaço extra é elasticidade que o TeX usa para fechar a linha sem hifenizar. E vale
para as três edições juntas, não só para a brochura.

### C. Espessura do Pólen Bold (pendência antiga, de 25/08)

As seis capas de `brochura/capas-papeis/` foram refeitas para as 348 páginas
(`brochura/gera-capas-papeis.py`, que lê a paginação do próprio miolo):

| papel | mm/folha | lombada 352 pp | lombada 348 pp |
|---|---|---|---|
| Pólen Soft 80 g/m² | 0,0920 | 16,2 mm | **16,0 mm** |
| Pólen Bold 90 g/m² | 0,1000 | 17,6 mm | **17,4 mm** |
| Couché Fosco 115 g/m² | 0,1000 | 17,6 mm | **17,4 mm** |
| Offset 75 g/m² | 0,1044 | 18,4 mm | **18,2 mm** |
| Couché Fosco 150 g/m² | 0,1200 | 21,1 mm | **20,9 mm** |
| Offset 90 g/m² | 0,1280 | 22,5 mm | **22,3 mm** |

Espessuras da tabela publicada pela PoloPrinter (`poloprinter.com.br/calcule-lombada/`),
a única das três calculadoras consultadas que abre os números.

**A ressalva de 25/08 continua de pé:** essa tabela dá 0,100 mm/folha para o Pólen
Bold 90g, o mesmo valor do Couché 115g. Isso implica volume ~1,11, que é baixo e
contradiz o argumento de venda do papel — a Suzano vende o Pólen Bold como papel de
alto volume. Se o volume real for 1,8, a lombada vai a ~28 mm em vez de 17,4. É o
papel mais provável para este livro e é justamente onde a divergência é maior:
**perguntar a espessura à gráfica antes de imprimir.**

---

## Fora do alcance do template

Quatro hífens de fim de linha sobraram, todos dentro da arte das figuras e todos
quebrando num hífen que já existe na palavra (`moderada-severa`, `estresse
pós-traumático`, `cognitivo-comportamental`, `ex-fumante`). Quebra legítima.

O único ruim de verdade era `Neuro-degeneração`, quebrado à mão no script da
Cap12_Fig01; agora vai inteiro, com o corpo encolhendo até caber na caixa.

---

## O que falta

O editor parou na página 46 de 348. As páginas 47 em diante não foram revisadas por
ninguém com olho editorial. A sugestão dele — imprimir uma prova e marcar no papel —
continua de pé, e agora vale mais a pena: a prova sai com os oito itens acima já
corrigidos.
