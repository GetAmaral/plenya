# Medical Lecture Builder — Instruções do Skill /aula

> Este arquivo é referência interna usada por `.claude/commands/aula.md`.
> Engine: HTML + CSS + Playwright (mesma stack do deck Continuum), com paleta mista
> petrol-deep (cerimoniais) + cream (conteúdo) e tipografia Cormorant + Fraunces.

---

## Modo de Operação

Analise `$ARGUMENTS`:

- "editar slide N" / "modificar slide N" → **MODO EDIT**
- "regenerar imagem slide N" / "trocar imagem slide N" → **MODO IMAGE**
- "exportar" / "gerar pdf" sem novo tema → **MODO EXPORT**
- Qualquer outra coisa → **MODO CREATE**

---

## MODO CREATE — Nova Aula

### Passo 0: Briefing via AskUserQuestion

**OBRIGATÓRIO.** Use `AskUserQuestion` com as 4 perguntas abaixo em uma única chamada.

```
AskUserQuestion(questions=[
  { "question": "Qual o nível do público?", "header": "Público-alvo", "multiSelect": false, "options":[
      {"label":"Residência médica","description":"R1–R3, conhecimento clínico intermediário"},
      {"label":"Graduação","description":"Estudantes de medicina"},
      {"label":"Pós-graduação / Fellows","description":"Especialistas em formação avançada"},
      {"label":"Congresso / Atualização","description":"Plateia mista"}
  ]},
  { "question": "Qual a duração da aula?", "header": "Duração", "multiSelect": false, "options":[
      {"label":"30 minutos","description":"~10 slides"},
      {"label":"45 minutos","description":"~14–17 slides"},
      {"label":"60 minutos","description":"~18–20 slides"},
      {"label":"90 minutos","description":"~24–28 slides"}
  ]},
  { "question": "Quem está apresentando?", "header": "Apresentador", "multiSelect": false, "options":[
      {"label":"Dr. Getúlio Amaral Filho — Nefrologista","description":"Padrão"},
      {"label":"Outro apresentador","description":"Informe na descrição"}
  ]},
  { "question": "Descreva o que você quer na aula", "header": "Briefing", "multiSelect": false, "options":[
      {"label":"Foco diagnóstico","description":"Reconhecer, investigar, classificar"},
      {"label":"Foco terapêutico","description":"Conduta, protocolos, manejo"},
      {"label":"Fisiopatologia + clínica","description":"Mecanismo + apresentação"},
      {"label":"Descrição livre","description":"Use o campo de notas"}
  ]}
])
```

Após resposta, criar o diretório de saída:

```bash
SLUG="<kebab-case-do-tema>"
mkdir -p ~/lecture-output/${SLUG}/images
```

---

### Passo 1: Busca de literatura

```bash
bash /home/user/plenya/.claude/skills/lecture-builder/scripts/search-rag.sh "TEMA"
bash /home/user/plenya/.claude/skills/lecture-builder/scripts/search-pubmed.sh "TEMA EN" 12
```

Monte lista numerada de referências reais com autores, título, journal, ano e **PMID real**. Se faltar PMID, use WebSearch para confirmar antes de citar. **Nunca alucinar artigo.**

---

### Passo 2: Definir fio condutor

Antes de qualquer slide, defina por escrito:

- **Gancho de abertura**: situação clínica concreta ou dado surpreendente
- **Caso clínico central**: nome, idade, dados objetivos realistas; aparece na abertura, retorna na virada (~⅓-½ do percurso) e fecha no desfecho
- **Arco narrativo**: introdução → aprofundamento → virada → resolução
- **Frase de fechamento**: a única coisa que o público leva embora

---

### Passo 3: Montar outline.json

```json
{
  "title": "...",
  "audience": "residência",
  "duration_min": 45,
  "slug": "tema-kebab",
  "output_dir": "/home/user/lecture-output/tema-kebab",
  "hook": "...",
  "clinical_case": {
    "name": "Lucas, 24 anos",
    "presentation": "...",
    "twist": "...",
    "resolution": "..."
  },
  "closing_message": "...",
  "sections": [
    {
      "slide_number": 4,
      "title": "Fisiopatologia",
      "ceremonial": false,
      "section_eyebrow": "FISIOPATOLOGIA",
      "duration_min": 6,
      "key_points": ["..."],
      "refs": [1, 7],
      "image_prompt": "Diagram of ..."
    }
  ],
  "references": [
    {"index":1,"authors":"Smith RJH et al.","title":"...","journal":"Nat Rev Nephrol","year":2019,"pmid":"30692664"}
  ]
}
```

Convenção: `ceremonial: true` em capa, caso clínico (abertura), virada, desfecho do caso, mensagem final, referências → **slide petrol-deep**. Resto → **cream**.

---

### Passo 3.5: 🛑 ROTEIRO — discutir conteúdo ANTES de gerar slides

**Obrigatório. Nunca pule para os slides sem essa rodada.**

Apresente ao usuário um documento de roteiro em **prosa**, não em formato de slide:

```
# Roteiro — [Título]
**Duração:** [N] min · **Público:** [X] · **Apresentador:** [Y]

## Fio condutor
[1-2 parágrafos descrevendo o arco. Não em bullets.]

## Caso clínico — [Nome, idade]
- Apresentação inicial: [parágrafo]
- Virada (slide ~N): [parágrafo]
- Desfecho (slide ~N): [parágrafo]

## Estrutura proposta
| # | Slide | Min | O que diz | Por que aqui |

## Mensagem de fechamento
> "..."

## Referências âncora (PMID real)
1. ...
```

Termine com **5 perguntas explícitas** ao usuário (caso verossímil? fio condutor afinado? slide faltando ou sobrando? profundidade calibrada? mensagem final afiada?).

**Aguarde aprovação explícita** ("ok, pode gerar"). Refaça quantas rodadas forem necessárias.

---

### Passo 4: Gerar `lecture.html`

Stack: HTML + CSS via Cormorant Garamond + Fraunces (Google Fonts), renderizado por Playwright (mesma engine do `scripts/deck-builder/continuum`).

**Estrutura do arquivo** (em `<output_dir>/lecture.html`):

```html
<!DOCTYPE html>
<html lang="pt-BR">
<head>
  <meta charset="UTF-8">
  <title>{{TITLE}}</title>
  <link rel="stylesheet" href="lecture.css">
</head>
<body>
<!-- slides aqui -->
</body>
</html>
```

Copiar `templates/lecture.css` para `<output_dir>/lecture.css`:

```bash
cp /home/user/plenya/.claude/skills/lecture-builder/templates/lecture.css <output_dir>/lecture.css
```

**Cada slide** é um `<section class="slide ...">` com viewport fixo 1920×1080. Padrões:

#### Slide cerimonial (capa, caso clínico, virada, mensagem, refs)
```html
<section class="slide slide--deep" id="s01">
  <div class="chrome-brand">DR. GETÚLIO AMARAL FILHO · NEFROLOGIA</div>
  <div class="chrome-eyebrow">SEÇÃO</div>
  <div class="chrome-page">01 / 17</div>
  <div class="slide-content">
    <h1 class="t-display">Título<span class="dot">.</span></h1>
    <p class="t-subhead">Subtítulo italic gold</p>
  </div>
</section>
```

Para capa pura e mensagem final use `slide--deep slide--bare` (sem chrome).

#### Slide de conteúdo cream
```html
<section class="slide slide--cream" id="s04">
  <div class="chrome-brand">DR. GETÚLIO AMARAL FILHO · NEFROLOGIA</div>
  <div class="chrome-eyebrow">FISIOPATOLOGIA</div>
  <div class="chrome-page">04 / 17</div>
  <div class="slide-content">
    <h2 class="t-h1">Headline do slide<span class="dot">.</span></h2>
    <p class="t-subhead">Subtítulo opcional.</p>
    <ul class="bullets">
      <li>Bullet sem marcador, hairline gold-soft <sup class="cite">[4]</sup></li>
    </ul>
    <div class="callout callout--concept">
      <span class="callout-title">Conceito-chave</span>
      Pista clínica direta.
    </div>
  </div>
</section>
```

#### Tabela clínica
```html
<table class="table-clinical">
  <thead><tr><th>Regulador</th><th>Onde age</th><th>Função</th></tr></thead>
  <tbody>
    <tr><td class="col-key">CFH</td><td>Fase fluida + superfícies</td><td>Decay-accelerating de C3bBb</td></tr>
  </tbody>
</table>
```

#### Callouts disponíveis
- `.callout--concept` (gold) — conceito-chave
- `.callout--alert` (vermelho) — alerta clínico
- `.callout--evidence` (verde) — guideline/evidência
- `.callout--case` (petrol) — caso clínico reentrando

#### Citações
- Inline: `frase ... <sup class="cite">[4]</sup>`
- Slide final de Referências: `<ol class="refs">` com `<li>Autor et al. Journal. Ano. PMID: <span class="pmid">12345678</span></li>` — PMIDs reais.

#### Imagem
- Hero full-bleed: `<div class="hero-img-fullbleed"><img src="images/section-4.png"></div>` (fade automático para texto à esquerda)
- Hero direita: `<div class="hero-img-right"><img src="..."></div>`
- Inset central: `<img class="hero-img-inset" src="...">`

#### Regras editoriais (igual deck Continuum)
- **Aspas «» francesas** ou curly `"…"`, nunca `"…"` retas
- **Sem em-dashes** em texto narrativo (vírgula/ponto)
- **Sem ícones decorativos**, sem "Não é X. É Y." empilhado
- Tom: prosa clínica conectiva, direto colega-para-colega

#### 🚨 Regras editoriais INVARIANTES (não esquecer)

**1. Footer de referências em TODO slide com claim de literatura.**
Toda afirmação clínica (epidemiologia, mecanismo, definição, conduta, dado de estudo) precisa de fonte visível no slide, não só inline. Use `.chrome-refs` no canto inferior esquerdo:
```html
<div class="chrome-refs">
  <span class="ref-tag">[4]</span>Fakhouri F et al. N Engl J Med 2025 — VALIANT
</div>
```
Múltiplas refs separadas por `<span class="ref-sep">·</span>`. Apenas slides cerimoniais puros (capa, headline de seção, caso clínico) podem omitir o footer.

**2. Acrônimos expandidos na primeira aparição — SEMPRE no fluxo do texto, NUNCA em legenda separada.**
A expansão entra na própria frase, em parênteses, na primeira vez que a sigla aparece:
- ✅ `A glomerulonefrite membranoproliferativa (GNMP) é um padrão, não um diagnóstico.`
- ✅ `Deposição de C3 ≥ 2 ordens de magnitude maior do que qualquer imunoglobulina (Ig).`
- ❌ NUNCA criar uma linha-legenda separada listando siglas (`GNMP — ... · MBG — ...`). É redundante e feio. Proibido.

Regras práticas:
- Se a sigla está rotulada **dentro de uma imagem/diagrama** (ex.: a cascata já mostra "MAC — Complexo de Ataque à Membrana"), a expansão já está feita — não repetir no texto.
- **Símbolos de gene** (CFH, CFI, CFB, C3, CFHR1-5, MCP, DGKE) são identificadores padrão — não precisam de expansão quando aparecem como lista de painel genético.
- Prefira escrever o termo por extenso no headline/subtítulo quando couber (`Fator H — o freio principal` em vez de `CFH — ...`), deixando a sigla só onde o espaço exige.
- Mantenha rastreabilidade slide-a-slide do que já foi introduzido; não reexpandir.

**3. Convenções de unidade clínica.**
- Proteinúria em **mg/g** (não g/g)
- Creatinina em mg/dL
- eGFR em mL/min/1,73 m²
- Pressão em mmHg

---

#### Régua tipográfica
- `.t-display` 96px → headlines cerimoniais
- `.t-h1` 64px → headline de slide conteúdo
- `.t-h2` 48px → headline secundária
- `.t-h3` 34px → subtítulo de coluna
- `.t-subhead` 36px italic → frase-âncora
- `.t-body` 26px / `.t-body-lg` 30px → texto corrido
- Tabelas: 26px (linhas), 20px caps (cabeçalho)
- Bullets: 28px com hairline gold-soft entre itens

---

### Passo 5: Ilustrações — DOIS sistemas

**Regra de ouro:** o residente precisa **ler rótulos** nos diagramas. Duas alternativas válidas, escolher por contexto:

- **gpt-image-2 com rótulos no prompt** — o modelo renderiza texto técnico bem (inclusive em diagramas científicos). Pedir explicitamente: "label each molecule with its name (C3, C3b, Bb, Factor H, etc.)". Vantagem: imagem rica, anatomicamente sofisticada.
- **SVG inline** — controle absoluto, paleta perfeita, editável sem regerar. Vantagem: zero risco de rótulo torto, custo zero, iteração rápida.

Use SVG quando o diagrama é **estritamente esquemático** (charts, heredogramas, score gauges, mecanismos com formas simples). Use gpt-image-2 quando quer **textura visual** (cascatas anatômicas, micrografias didáticas com call-outs, biomoléculas em estrutura terciária, vias com órgãos).

#### 5a · SVG inline (esquemas geométricos puros)

Para charts, heredogramas, score gauges, diagramas com retângulos+círculos+setas — escreva SVG inline direto no HTML do slide.

Padrão (igual `scripts/deck-builder/continuum/build-v3.js` linhas 88-92, 135-138):
- `viewBox` fixo (ex.: `0 0 1600 720`)
- Paleta: cream/petrol/gold/sage como `fill`/`stroke`
- Tipografia: `font-family="Inter, sans-serif"` (corpo) ou `"Fraunces, serif"` (display)
- Setas: `<marker>` reutilizável com `marker-end="url(#arr)"`
- Loops/curvas: `<path d="M ... Q ..." stroke-dasharray="8 4">`
- Rótulos: `<text>` com `font-weight="700"` para nomes de moléculas, `font-style="italic"` para qualificadores

**Tipos de slide que devem usar SVG:**
- Cascatas de complemento/coagulação/inflamação
- Esquemas de molécula + alvo (drogas)
- Gráficos de barras, linhas, radar
- Heredogramas (4 figuras genéticas básicas)
- Quadros comparativos (atividade vs cronicidade do C3G-HI)
- Score gauges
- Dipstick

#### 5b · gpt-image-2 (ilustração visual rica — pode incluir rótulos)

Para ilustrações ricas em textura/anatomia/profundidade, mecanismos com forma orgânica, micrografias, fotos editoriais. **gpt-image-2 renderiza texto bem em diagramas científicos** — peça explicitamente os rótulos no prompt.

```bash
bash /home/user/plenya/.claude/skills/lecture-builder/scripts/generate-image.sh \
  "PROMPT com labels especificados" "section-N.png" "<output_dir>"
```

**Padrão de prompt para diagrama labeled:**
- Descrever a cena/mecanismo com precisão
- **Listar explicitamente os rótulos exigidos**: `Label each molecule with its name: "C3", "C3b", "Bb", "Factor B", "Factor D", "Properdin", "C5", "C5a", "C5b", "MAC"`
- Pedir tipografia editorial: `clean sans-serif labels, scientific textbook style`
- Especificar paleta: `restrained color palette — cream, deep teal, muted gold, sage accents`

**Tipos de slide bons para gpt-image-2:**
- Capa atmosférica (rim em corte, paisagem editorial)
- Cascatas anatômicas com biomoléculas em forma terciária (labeled)
- Micrografia eletrônica didática
- Imunofluorescência de glomérulo
- Foto editorial de microscópio/consultório
- Cortes anatômicos labeled (glomérulo + MBG + podócitos)
- Imagem contemplativa para mensagem final (sunrise, horizonte)

Saída em `<output_dir>/images/section-N.png`. Referencie no HTML com `<img src="images/section-N.png">` ou via `.hero-img-*` do CSS.

---

### Passo 6: 🖼 Iteração visual — preview slide-a-slide

**Obrigatório. Não pular para PDF.**

```bash
cd /home/user/plenya/.claude/skills/lecture-builder/scripts
node render-pngs.js --dir=<output_dir> --slide=NN
```

→ Gera `<output_dir>/previews/slide-NN.png` (1920×1080 @2x).

**Para cada slide gerado, mandar o PNG ao usuário via `SendUserFile`** antes de prosseguir. Aceitar correções (texto, layout, imagem) e re-renderizar. Apenas depois que todos os slides estiverem aprovados, gerar o PDF.

(Regra paralela à do deck: ver memória `deck_fine_tuning_workflow`.)

Para gerar PNG de **todos** os slides de uma vez:

```bash
node render-pngs.js --dir=<output_dir>
```

---

### Passo 7: PDF final

```bash
cd /home/user/plenya/.claude/skills/lecture-builder/scripts
node render-pdf.js --dir=<output_dir>
```

→ `<output_dir>/lecture.pdf` (1920×1080, multi-slide, fonts embedded).

---

### Passo 8: Resumo ao usuário

```
✅ <output_dir>/
   lecture.html | lecture.css | lecture.pdf | outline.json
   images/ (gpt-image-2)
   previews/slide-NN.png (1920×1080)

/aula editar slide N "instrução"
/aula regenerar imagem slide N
/aula exportar
```

---

## MODO EDIT

```bash
ls ~/lecture-output/*/outline.json 2>/dev/null
```

1. Leia `outline.json` → `output_dir`
2. Leia `lecture.html`
3. Localize `<section ... id="sNN">` (numeração com pad-zero)
4. **Edit** o conteúdo do `<section>` mantendo classes do design system
5. Re-render PNG do slide editado e envie:
   ```bash
   node /home/user/plenya/.claude/skills/lecture-builder/scripts/render-pngs.js --dir=<output_dir> --slide=NN
   ```
   `SendUserFile <output_dir>/previews/slide-NN.png`
6. PDF só após aprovação:
   ```bash
   node /home/user/plenya/.claude/skills/lecture-builder/scripts/render-pdf.js --dir=<output_dir>
   ```

---

## MODO IMAGE

1. Leia `outline.json` → `image_prompt` da seção N
2. Gere:
   ```bash
   bash /home/user/plenya/.claude/skills/lecture-builder/scripts/generate-image.sh \
     "PROMPT" "section-N.png" "<output_dir>"
   ```
3. Re-render PNG do slide e envie via `SendUserFile`
4. PDF só após aprovação

---

## MODO EXPORT

```bash
node /home/user/plenya/.claude/skills/lecture-builder/scripts/render-pdf.js --dir=<output_dir>
```
