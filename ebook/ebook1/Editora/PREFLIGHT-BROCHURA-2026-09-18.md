# Preflight da brochura — 18/09/2026

> **Estado:** os itens 1, 2 e 7 foram corrigidos e os arquivos rebuildados; o 3 está
> parcialmente resolvido e aguarda a gráfica. Detalhe no fim, em "O que foi feito".

Checagem de pré-impressão dos arquivos que vão para a gráfica: os três miolos
(`-cor`, `-pb`, `-meio`, 344 páginas cada) e as seis capas de `capas-papeis/`.
Não é revisão de texto — é o que o operador de pré-impressão olha antes de rasterizar.

---

## Resolver antes de enviar

### 1. O miolo não diz à gráfica onde cortar

A página tem **165 × 235 mm**. O projeto é trim 160 × 230 + 2,5 mm de sangria em cada
lado (está no cabeçalho do template). Só que no PDF:

    MediaBox = CropBox = TrimBox = BleedBox = ArtBox = 165 × 235 mm

Todas as caixas iguais, e nenhuma marca de corte. Quem abrir o arquivo vê uma página de
165 × 235 sem instrução nenhuma de refile. O resultado provável é o livro sair 165 × 235
— 5 mm maior em cada dimensão — e a capa, que foi montada para 160 × 230, não fechar.

**E a sangria não é usada por nada.** Varri as 344 páginas: o conteúdo mais próximo da
borda está a **9,5 mm** dela (p.12). Nenhum elemento sangra.

Duas saídas, e a escolha é de quem falou com a gráfica:

- **gerar o miolo já no trim 160 × 230, sem sangria** — é o normal para miolo que não
  sangra, e é só tirar 2,5 mm de cada margem do `geometry` (a mancha de texto não muda
  de lugar); ou
- **manter 165 × 235 e declarar TrimBox de 160 × 230 centralizado**, com marcas de corte.

### 2. O miolo P&B não é preto de uma chapa só

O corpo do texto está certo: sai como `0 g`, DeviceGray, 100% K.

O resto não. Cabeço, títulos de capítulo e seção, rótulo de legenda, "PARTE N" /
"CAPÍTULO N" e capitulares saem em **DeviceRGB**, porque as três cores do template P&B
estão definidas em HTML:

    \definecolor{petrol}{HTML}{111111}
    \definecolor{gold}{HTML}{767676}
    \definecolor{ink}{HTML}{1a1a1a}

São **2.824 operadores RGB em 321 das 344 páginas**. Os valores são neutros (R=G=B, 0%
de pixel cromático), mas o espaço de cor é RGB. Numa conversão para CMYK sem "preservar
preto", isso vira preto de quatro cores — inclusive no cabeço de 8,25 pt que aparece em
toda página. Texto pequeno em quatro chapas é halo de registro garantido.

A correção é trocar `{HTML}{...}` por `{gray}{...}`: petrol → 0.067, gold → 0.463,
ink → 0.102. Uma linha por cor, nenhuma diferença na tela.

### 3. O miolo colorido é inteiramente RGB

Zero operador CMYK nos três arquivos. O petrol da marca é RGB(6, 59, 79) e vai ser
convertido pela gráfica com o perfil que ela usar — a cor desloca e ninguém aqui
controla para onde. Não é erro, mas é uma decisão que precisa ser combinada com a
gráfica em vez de presumida, e é a mesma discussão da papelaria (converter para CMYK só
depois de aprovação visual).

---

## A gráfica vai apontar

| # | Achado | Onde |
|---|---|---|
| 4 | **Duas imagens abaixo de 300 dpi**: 433×375 px em 51,5×44,6 mm = **213 dpi**; e um ícone de 41×52 px, também 213 dpi | p.168 (Fig. 9.1) e p.180 (Fig. 10.3) |
| 5 | **Fontes Type3** (glifos como procedimento, sem arquivo embutido), vindas das figuras feitas em matplotlib | p.210, 234, 242, 249, 259 |
| 6 | **Duas imagens Indexed(DeviceRGB) sobrevivem no miolo P&B** — pixels neutros, espaço de cor RGB | p.168 e p.180 |
| 7 | **A capa também não tem TrimBox nem marcas de dobra.** Capa com orelhas tem quatro dobras e o arquivo não diz onde ficam | as 6 capas |
| 8 | **A capa é um raster único de 600 dpi em RGB**, inclusive o código de barras. 600 dpi dá conta, mas código de barras rasterizado é risco de leitura, e o texto perde o gume do vetor | as 6 capas |

---

## Conferido e correto

- **344 páginas, múltiplo de 4** — fecha caderno sem sobra.
- **As 23 aberturas** (5 partes + 18 capítulos) caem todas em página ímpar.
- **Todas as fontes de texto embutidas** e com subset.
- **Nenhum fio abaixo de 0,25 pt** — nada some na impressão.
- **Nenhuma transparência** (zero ExtGState), nada a achatar.
- **Tamanho e rotação de página uniformes** nas 344.
- **Nada invade a faixa de sangria** — folga mínima de 9,5 mm.
- **ISBN do miolo bate com o do código de barras da capa**: 978-65-02-07691-0.
- **Lombada**: o texto fica com 6,5 mm de folga de cada lado nos 17,2 mm, centrado; o
  petrol cobre a lombada inteira até as duas dobras.
- **Geometria da capa**: 3 + 80 (orelha) + 160 + lombada + 160 + 80 (orelha) + 3 mm,
  altura 236 = 230 + 3 + 3. As seis batem com a lombada do nome do arquivo.
- **Sangria da capa** cobre as quatro bordas, sem nenhum pixel branco.
- **Código de barras** a 12,5 mm do corte inferior e 12,2 mm da dobra da lombada.
- **"1ª edição"** sai com o indicador ordinal certo (a extração de texto mostra "1a",
  mas o glifo impresso é ª).

---

## Falta, mas não é da gráfica

A página de créditos não traz **ficha catalográfica (CIP)**. Não impede a impressão e não
é obrigatória para edição do autor, mas é o que se espera de um livro com ISBN, e quem
elabora é bibliotecário com CRB. Vale decidir antes de fechar o miolo, porque ela ocupa
espaço na p.4.

---

## O que foi feito (18/09, depois da checagem)

### 1 — TrimBox declarado, corrigido

Passo novo no fim do build (`set_print_boxes`, em `build-brochura-miolo.py`): as 344
páginas dos três miolos saem agora com

    MediaBox  165 × 235 mm     (papel, com a sangria)
    TrimBox   160 × 230 mm     a 2,5 mm de cada borda
    BleedBox  165 × 235 mm
    ArtBox    ausente de propósito

Escolhi este caminho, e não o de gerar o miolo já em 160 × 230, porque ele **não mexe em
nada do projeto** — a mancha de texto, a paginação e as margens ficam idênticas, e o
arquivo só passa a dizer o que antes estava implícito. Se a gráfica pedir o miolo no
formato final sem sangria, é trocar quatro números no `geometry` e rebuildar; a mancha
não sai do lugar. Marcas de corte impressas não entram: elas precisariam de mais espaço
do que os 2,5 mm de sangria disponíveis.

**ArtBox fica de fora nos dois.** A primeira versão gravava ArtBox igual ao TrimBox, e
a revisão pegou: a ISO 15930 (PDF/X) manda a página ter TrimBox **ou** ArtBox, nunca os
dois, e Acrobat Preflight e PitStop reportam a coexistência como erro — não como aviso.
Era o tipo de coisa que volta da gráfica.

**A capa foi junto** (item 7): as seis passaram a declarar TrimBox da capa aberta sem
sangria — 497,2 × 230 mm no Pólen Bold, e assim por diante conforme a lombada. As quatro
dobras continuam só no PNG de guias, porque marca impressa não cabe em 3 mm.

### 2 — Miolo P&B agora é de uma chapa só, corrigido

As três cores do template P&B passaram de `{HTML}` para `{gray}`:

    petrol  #111111 → gray 0.067
    gold    #767676 → gray 0.463
    ink     #1a1a1a → gray 0.102

Os tons são os mesmos (conferido no PDF: o título do cap. 1 continua saindo RGB
17,17,17). O que muda é a chapa:

| | antes | depois |
|---|---|---|
| operadores nos streams de página | 5.648 `g/G` + **2.824 `rg/RG`** | **10.916 `g/G`, zero `rg/RG`** |

Cabeço, títulos, rótulo de legenda, "PARTE N"/"CAPÍTULO N" e capitulares agora saem em
DeviceGray.

### 3 — Miolo colorido: metade resolvida, metade esperando a gráfica

O `ink` (#1a1a1a, neutro) virou `gray` também no template colorido — é a cor do cabeço
de toda página, e não havia motivo para ela virar preto de quatro cores. Sobraram
**817 operadores RGB nos streams de página**, que são só o petrol e o gold, as duas
cores cromáticas da marca.

Essas duas eu **não converti**. Escolher os valores CMYK da marca muda a cor impressa e
depende do perfil que a gráfica usa; é a mesma regra da papelaria, converter só depois de
aprovação visual. Quando eles disserem o perfil, faço a conversão e uma prova para
comparar.

### Resíduo conhecido

Dentro das figuras (XObjects) ainda há operadores RGB: 371 no P&B, dos quais **366 são
neutros e 5 têm um resto de cor** (o mais "colorido" é RGB 0,604 / 0,604 / 0,627 — seis
pontos em 255 de diferença, invisível). Vêm das artes em `figuras-bw/`, que são
convertidas e não regeradas em cinza. Se a gráfica rodar o P&B como uma chapa só, isso
desaparece na conversão dela. Regerar as artes em DeviceGray é mexer no pipeline das
figuras e fica para quando valer a pena.

### Correções vindas da revisão do próprio passo

- **ArtBox saiu** dos nove arquivos (três miolos, seis capas): TrimBox e ArtBox juntos
  são erro de PDF/X.
- **O passo agora falha alto.** Antes, na capa, ele era um `try/except` que só imprimia
  `⚠` — e o wrapper `gera-capas-papeis.py` descarta o stdout quando o retorno é zero, de
  modo que as seis capas podiam sair sem TrimBox com um `✅ 6 capas` na tela. Agora a
  falta de pikepdf, ou um erro no `save()` que deixasse o PDF truncado, para o build.
- **O log diz o que foi gravado**, lendo de volta o TrimBox do arquivo, em vez de repetir
  o valor que se esperava gravar. Se um dia o miolo for gerado em 160 × 230 sem sangria, o
  log mostra o valor real em vez de afirmar 160 × 230 por hábito.
- **O fallback de interpretador** usa `shutil.which("python3")` e trata `OSError`, como já
  fazia o `pad_to_signature`.

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
