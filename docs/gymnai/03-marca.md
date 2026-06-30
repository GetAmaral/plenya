# 03 — Marca Gymnai (identidade visual)

> Identidade **fechada** (recebida do design 2026-06-29). Resolve D5. Assets em
> [`identidade/`](identidade/). Tom: **premium / performance** — luxo atlético, ouro sobre
> petróleo. Distinta da Plenya, mas da mesma família cromática (gold/petrol/cream).

## Assets de referência

**Fontes PNG (mockups + alta-res):** `gymnai-brandbook.png`, `gymnai-placa-entrada.png`,
`gymnai-tag-aparelho-1/2.png`, e os alta-res `73991D30…PNG` (brandbook+paleta), `A4122B6D…PNG`
(wordmark), `A7B45BBE…PNG` (símbolo isolado).

**Vetores/tokens gerados (a partir dos PNGs):**
- [`identidade/gymnai-symbol.svg`](identidade/gymnai-symbol.svg) — símbolo vetorizado (traçado do alta-res).
- [`identidade/gymnai-logo.svg`](identidade/gymnai-logo.svg) — lockup completo (símbolo + GYMNAI + tagline).
- [`identidade/gymnai-wordmark.svg`](identidade/gymnai-wordmark.svg) — só o lettering (Cinzel).
- [`identidade/gymnai-palette.svg`](identidade/gymnai-palette.svg) · [`gymnai-tokens.css`](identidade/gymnai-tokens.css).
- [`identidade/icons/`](identidade/icons/) — favicon + app-icons (símbolo ouro sobre petróleo), gerados do
  alta-res: `favicon.ico` (16/32/48), `favicon-16/32/48.png`, `apple-touch-icon-180.png`,
  `icon-192.png`, `icon-512.png`, `icon-maskable-512.png` (safe-zone p/ PWA) e
  `symbol-512-transparent.png` (símbolo recortado, fundo transparente).

> O símbolo é **traço automático** (OpenCV) do PNG — usável para dev/web; para o master final, vale
> uma limpeza de curvas no Illustrator/Inkscape ou recuperar o vetor original.

## Nome, símbolo e assinatura

- **Nome / wordmark:** `GYMNAI` (caixa alta, espaçado, ouro).
- **Símbolo:** figura atlética estilizada em ouro (círculo + swoosh) — forma humana dinâmica/em
  movimento. Usável **isolada** como ícone/app-icon/favicon. Vetorizado em
  [`identidade/gymnai-symbol.svg`](identidade/gymnai-symbol.svg) (traçado do PNG alta-res).
- **Assinatura / tagline:** `INTELLIGENCE IN MOTION` — **canônica**. (Um mockup explorou
  "ADAPTIVE TRAINING INTELLIGENCE"; **ignorar** — não é a tagline.)
- **Versão recomendada:** logo em **ouro sobre fundo petróleo escuro** (a marca é dark-first).

## Tipografia

| Papel | Fonte | Uso |
|-------|-------|-----|
| Principal | **Trajan Pro** → na web: **Cinzel** | Logotipo, títulos, frases curtas. Serifa romana clássica, **só caixa alta**. |
| Secundária | **Montserrat** | Corpo de texto, web, aplicações (tem caixa baixa). |

> **Decisão de fonte (M7):** Trajan Pro é **paga** (Adobe). Para web usamos **Cinzel** (Google
> Fonts, **OFL — livre p/ comercial e web**), a alternativa clássica direta ao Trajan (capitulares
> romanas, mesma pegada). Se houver assinatura Adobe CC, o Trajan vem no Adobe Fonts (licenciado p/
> web) e pode ser usado no lugar. **Montserrat** (free) para todo o resto. No **logo**, o lettering
> "GYMNAI" deve ser **outline** (contornos), p/ não depender de fonte instalada — ver
> [`identidade/gymnai-wordmark.svg`](identidade/gymnai-wordmark.svg) (hoje em Cinzel via webfont).

## Paleta

Família **ouro + petróleo + creme** com acentos em azul/teal (mesma DNA cromática da Plenya, com
identidade própria). Fundo **petróleo escuro predominante**.

Paleta **oficial** (5 swatches rotulados do brandbook em alta-res — amostrados 2026-06-29):

| Swatch (rótulo) | Papel | Hex |
|-----|-------|-----|
| **OURO** | primária (logo, títulos, bordas); gradiente até `#D7B975` | `#B4894D` |
| **PETROL** | **fundo predominante** | `#022837` |
| **AZUL CLARO** | acento teal | `#3C6971` |
| **VERDE** | sage / acento secundário | `#ADB29B` |
| **CREME** | texto claro / fundos claros | `#E1D9CC` |

> Valores **amostrados dos swatches rotulados** (não estimativa) — confiáveis salvo leve variação de
> render. Mesma família da Plenya (gold/petrol/ocean/sage/cream). Tokens em
> [`identidade/gymnai-tokens.css`](identidade/gymnai-tokens.css); swatches em
> [`identidade/gymnai-palette.svg`](identidade/gymnai-palette.svg).

## Aplicação no produto (a tag de QR)

O artefato central já está desenhado: **placa petróleo com borda ouro**, contendo de cima para
baixo — símbolo ouro, wordmark `GYMNAI`, `INTELLIGENCE IN MOTION`, o **QR code**, e a chamada
`ESCANEIE E ACESSE SUA EXPERIÊNCIA` com ícone de celular. Há uma versão de **sinalização de
entrada** (maior, com `SUA JORNADA COMEÇA AQUI` + ícones de feature) e a versão **compacta colada
no aparelho**.

Ícones de feature que aparecem na sinalização (pilares de comunicação do produto):
**Treinos adaptativos · Desempenho em tempo real · Saúde & longevidade · Evolução contínua.**
Reforçam o posicionamento "intelligence in motion" (treino adaptativo orientado por dado/IA) — o
que conecta diretamente com o motor de IA de treino do produto.

## Implicações para o código (design tokens)

- **Tema dark-first** no PWA: fundo petróleo, texto creme, ouro como cor de marca/realce.
- Tokens de marca próprios em `packages`/`theme` do Gymnai — **não** reusar tokens da Plenya
  (gold/petrol/ocean/sage/cream da Plenya são parecidos mas são outra marca).
- Fontes: Montserrat via Google Fonts; Trajan Pro (ou fallback) para hero/títulos, confirmando
  licença digital.
- A UI do scan→vídeo deve herdar a estética da placa: petróleo + ouro, sóbria, premium.
