# KDP — English Editions Upload Sheet

**Audience:** the upload session for `BEFORE` Kindle eBook + `BEFORE` Paperback on Amazon KDP (kdp.amazon.com).
**Source:** translated from `md/pt-BR/kdp-description.md` and `build/work-pt-BR/metadata.yaml`.
**Last revised:** 2026-05-10.

> **Como usar:** abre o KDP, vai em "+ Create" → escolhe Kindle eBook ou Paperback. Cola campo a campo seguindo a ordem aqui. Quando os dois formatos terminarem, a Amazon vai auto-linkar (Kindle + Paperback aparecem na mesma página de produto). Se em 48h não linkar sozinho, abre ticket no KDP Help → "Link my book formats".

---

## A. Shared metadata (use for BOTH Kindle and Paperback)

### Language
```
English
```

### Book Title
```
BEFORE
```

### Subtitle
```
The Silent Window — A Decade Between Normal and Optimal Where Health Is Decided
```

### Edition Number
```
1
```

### Author (Primary)
```
First name: Getulio
Last name: Amaral Filho
Prefix: Dr.
```
*Note: KDP separa First / Last. Coloca "Dr." no campo Prefix se aparecer; senão deixa em branco e adiciona na bio.*

### Contributors
```
(none — author solo)
```

### Description (HTML, max 4000 chars)

Cole exatamente este bloco no campo **Description** (KDP aceita HTML básico):

```html
<b>Ricardo had his annual check-up. "Everything looks normal," the doctor said. Eight months later, he nearly died in a parking lot.</b>
<br><br>
What no one had investigated: between the "normal" your lab prints and the optimal your body can reach, there is a <b>silent window</b>. Ten, fifteen, twenty years in which disease grows without showing up in any routine exam. Health — and longevity — is decided inside that window.
<br><br>
<b>This book is the map of it.</b>
<br><br>
<h4>What you'll find inside:</h4>
<ul>
<li>The <b>expanded biomarker panel</b> conventional check-ups never include</li>
<li>The <b>ACTS Method</b> — four pillars of evidence-based prevention (Assess, Coach, Treat, Sustain)</li>
<li>The <b>quarterly Longevity Scorecard</b> to track your trajectory in real numbers</li>
<li><b>Six real clinical cases</b> woven through the entire book</li>
<li>Practical protocols for physical activity, nutrition, smart supplementation, clinical-metabolic management, mind-body integration, and circadian rhythm</li>
</ul>
<h4>Who this book is for:</h4>
If you are 40, 50, or 60 and unwilling to wait for a diagnosis to act, this book is for you. It is not self-help. It is not a recipe manual. It is the work of a physician who has spent decades caring for patients who, exactly like you, receive "normal" lab results while their bodies tell a different story.
<br><br>
<h4>About the author:</h4>
<b>Dr. Getulio Amaral Filho</b> is a Brazilian nephrologist and internal medicine physician with more than 20 years of clinical practice. He cares for preventive-medicine and longevity patients every day and lectures nationally on health, nephrology, and longevity.
<br><br>
<i>Because living well and living longer doesn't begin after the diagnosis. It begins before.</i>
```

**Char count:** ~2.380 chars (well under the 4.000 KDP limit).
**Why this works for English readers:** the "Ricardo" hook in bold appears in the search-result preview before "Read more" — same conversion device as the PT version. Bullets are short, bold-led, with the ACTS acronym (English equivalent of AGIR) introduced in line.

### Keywords (7 max, ≤50 chars each)

Cole um por campo, sem aspas:

```
1. longevity medicine biomarkers
2. preventive medicine outlive readers
3. healthspan expanded blood panel
4. ApoB cholesterol metabolic health
5. VO2 max strength training longevity
6. quarterly longevity scorecard
7. silent window normal vs optimal
```
*Estratégia: cobrir leitores de Outlive (Attia), médicos preventivos, e termos longtail que diferenciam do mercado wellness genérico. Não repete palavra do título/subtítulo (não vale).*

### Categories (3 max — KDP permite 3)

```
Primary  : Health, Fitness & Dieting > Aging
Secondary: Medical Books > Internal Medicine > Nephrology
Tertiary : Health, Fitness & Dieting > Diseases & Physical Ailments > Heart Disease
```

**BISAC equivalents** (caso peça código):
- HEA049000 — HEALTH & FITNESS / Longevity
- MED059000 — MEDICAL / Internal Medicine
- HEA002000 — HEALTH & FITNESS / Aging

*Por que essas: "Aging" é a categoria mais quente do segmento (Outlive, Lifespan, Why We Sleep ranqueiam aí). "Internal Medicine > Nephrology" usa a especialidade do autor como diferenciação. "Heart Disease" pega o leitor de Atkinson/Esselstyn que busca prevenção cardiovascular.*

### Audience
```
Age range  : 18 and older
Sexually explicit: No
```

### Book Form
```
Original work (not public domain)
```

### Publishing Rights
```
I own the copyright and hold the necessary publishing rights
```

### Translated From (KDP form pergunta isso)
```
Original language: Portuguese (Brazil)
Original title:    ANTES — A Janela Silenciosa entre o Normal e o Ótimo
Original ASIN:     [colar o ASIN da edição PT digital aqui]
Original ISBN:     978-65-02-06742-0
Translator:        Dr. Getulio Amaral Filho (self-translation)
```
*Como pegar o ASIN PT: vai no produto da edição PT no Amazon, copia o B0XXXXXXXX da URL.*

### Series (deixar VAZIO)
```
(no series — see kdp-en-upload.md for rationale)
```
*Decisão: não usa series pra agrupar PT/EN porque series é pra livros sequenciais, não pra traduções. Author Central + A+ Content fazem o trabalho de cross-link sem misuse da feature.*

---

## B. Specific to **Kindle eBook** (digital)

### Step: Kindle eBook Details
*Já preenchidos os campos da seção A. Específico do digital:*

#### Manuscript file
```
File: build/Before-en.epub
Format: EPUB 3
```

#### Cover
```
File: build/Before-en-print-cover-front-only.png  (1600x2560 recommended for Kindle)
```
*Se não tem versão front-only ainda, eu gero a partir do print cover — me avisa.*

#### Reading age (optional)
```
Leave blank
```

#### DRM
```
Yes — Enable Digital Rights Management
```
*Conservador: ativa DRM no Kindle (impede compartilhamento de cópia). Não impacta venda — leitores nem percebem.*

### Step: Pricing & Royalty

#### Royalty plan
```
70% (KDP Select required if you want 70%)
```

#### KDP Select (Kindle Unlimited)
```
RECOMENDADO: Yes — Enroll in KDP Select for the first 90 days
```
*KDP Select dá: (a) 70% royalty em todos os preços qualificados, (b) Kindle Unlimited (leitores assinantes leem grátis, autor recebe por página lida), (c) ferramentas promocionais (Free Days, Countdown). 90 dias com exclusividade Amazon — depois você pode tirar.*

#### Price (USD — primary market)
```
$ 4.99
```
*Royalty: $4,99 × 70% = **$3,49 por venda**. Faixa $2,99–$9,99 = 70% royalty. Estratégia de lançamento agressiva pra autor desconhecido no mercado EN — volume cria ranking, ranking cria descoberta orgânica.*

#### Price by territory
| Territory | Currency | Price |
|---|---|---|
| United States | USD | **$4.99** |
| United Kingdom | GBP | £3.99 |
| Germany / France / Spain / Italy / Netherlands | EUR | €4.99 |
| Japan | JPY | ¥700 |
| Brazil | BRL | R$ 19,90 |
| Canada | CAD | $6.99 |
| Mexico | MXN | $89 |
| Australia | AUD | $7.99 |
| India | INR | ₹299 |

*Brazil a R$ 19,90 = paridade com PT digital (não cria arbitragem entre os 2 idiomas — bilíngue escolhe pelo gosto, não pelo preço).*

#### Pricing escalation roadmap
```
Mês 0 (lançamento)         : $4,99
Semana 2-3                  : 5 dias grátis (KDP Select Free Days)
Semana 5-6                  : Countdown Deal — $0,99 ou $1,99 por 7 dias
Mês 3 (se 30+ reviews ≥4★)  : sobe pra $7,99   (royalty $5,59)
Mês 6 (se 100+ reviews)     : sobe pra $9,99   (royalty $6,99 — TETO 70%)
Acima de $9,99              : ❌ cai pra 35% royalty — não vale
```

#### MatchBook
```
No
```
*MatchBook é descontinuado pela Amazon, deixa em No.*

#### Book Lending
```
Yes
```

---

## C. Specific to **Paperback (Print)**

### Step: Paperback Details
*Campos da seção A já preenchidos. Específico:*

#### Print ISBN
```
Option A (recommended for global distribution): own ISBN
ISBN: 978-65-975814-0-5  (CBL Brazilian agency, English edition)

Option B: Free KDP ISBN
Use this if you only want Amazon distribution and don't have your own ISBN.
```
*Verifica se 978-65-975814-0-5 ainda é válido na CBL pra esta edição. Se sim, usa A. Se a CBL exige tiragem mínima ou contrato editorial, fica com Free KDP ISBN (B).*

#### Publication Date
```
Leave blank (KDP fills automatically when you click Publish)
```

#### Publisher
```
Independently Published
```
*Ou "Author's Edition" se preferir. KDP aceita custom publisher name se você tiver registrado.*

### Step: Print Options

#### Ink and Paper Type
```
Black ink with cream paper      ← RECOMENDADO (premium feel, easier on eyes)
```
*Alternativa "Black ink with white paper" se quiser a vibe mais "manual técnico". Cream paper é o padrão de bestsellers (Outlive, Lifespan, Why We Sleep todos cream).*

#### Trim Size
```
6 x 9 in   (15.24 x 22.86 cm)
```

#### Bleed Settings
```
No bleed
```
*Nosso PDF não tem sangria (KDP padrão).*

#### Cover Finish
```
Matte                ← RECOMENDADO (premium, anti-glare, segura digital fingerprints melhor)
```
*Glossy é mais barata mas reflete luz e mostra dedo — pra livro premium, matte é padrão.*

### Step: Manuscript

#### Manuscript file
```
File: build/Before-en-print-interior.pdf
Pages: 410
```

#### Cover file
```
File: build/Before-en-print-cover-410pp.pdf
```
*Cover já tem a lombada calculada pra 410 páginas — verificar se o page count atual ainda é 410 antes de subir; se mudou, regenerar a capa.*

### Step: Pricing & Royalty (Paperback)

#### Royalty plan (only 60% available for paperback)
```
60% royalty option (only one available)
```

#### Royalty math (paperback)
```
KDP paperback = 60% royalty FIXA (não tem 70%)
Royalty = (Preço × 0,60) − Printing Cost
Printing cost (US, 410 pp, cream paper, 6×9): $5,92 (KDP confirmou)

Mínimo viável  : $9,99   → royalty $0,07     (inviável)
Sweet spot     : $19,99  → royalty $6,07     ⭐ recomendado
Premium        : $24,99  → royalty $9,07     (Outlive paperback territory)
```

#### Price (USD — primary market)
```
$ 19.99
```
*Por quê $19,99: matches Sinclair/Walker/Bryson paperback ($18,99-19,99). Te coloca **no clube dos bestsellers** sem cobrar premium de autor estabelecido. Royalty $6,07 por venda — saudável, não chase de volume mínimo.*

#### Pricing
| Territory | Currency | Price | Approx Print Cost | Approx Royalty |
|---|---|---|---|---|
| Amazon.com | USD | **$19.99** | $5,92 | **$6,07** |
| Amazon.co.uk | GBP | £15.99 | ~£4,80 | £4,79 |
| Amazon.de / .fr / .it / .es / .nl | EUR | €19.99 | ~€5,80 | €6,19 |
| Amazon.com.br | BRL | **R$ 89,90** | ~R$ 28-35 | ~R$ 19-23 |
| Amazon.co.jp | JPY | ¥2.800 | ~¥1.000 | ¥680 |
| Amazon.ca | CAD | $25.99 | ~$8,30 | $7,29 |
| Amazon.com.au | AUD | $29.99 | ~$10 | $7,99 |

*KDP mostra o "Printing Cost" exato por território quando você preenche o page count (410 pp). Royalty efetiva pode variar ±15% do estimado.*

#### Cross-format ratio sanity check
```
Kindle    $4,99   →  royalty $3,49
Paperback $19,99  →  royalty $6,07

Ratio paperback/kindle = 4,0:1
Industry standard nonfiction = 3,8-4,2:1 (Sinclair, Walker, Bryson)
✅ Em paridade com o segmento — leitor digital tem entrada barata,
   leitor físico paga premium esperado, ninguém sente arbitragem.
```

#### Pricing escalation roadmap (paperback)
```
Mês 0 (lançamento)         : $19,99
Mês 3 (se 30+ reviews ≥4★)  : sobe pra $22,99  (royalty $7,87)
Mês 6 (se 100+ reviews)     : sobe pra $24,99  (royalty $9,07 — match Outlive paperback)
Black Friday / Natal        : Countdown Deal $14,99 por 7 dias
TETO recomendado            : $24,99 (acima vira hardcover territory; tu já tem capa dura via UICLAP)
```

#### Expanded Distribution
```
Yes — Enable Expanded Distribution
```
*Free. Distribui pra Barnes & Noble e indies americanas via Ingram. Royalty cai pra 40% nessas vendas:*
```
Expanded royalty = $19,99 × 0,40 − $5,92 = $2,07/venda
```
*Ainda positivo, e exposição que não compra com nada. Vale.*

---

## D. Author Central (depois de publicar)

Quando os 2 EN forem publicados, **claim ambos no Author Central US**:

```
URL: https://author.amazon.com (US — separate account from .com.br)
Steps:
1. Sign in (or create) with the same email used for KDP
2. "Books" → "Add more books"
3. Search "BEFORE Getulio Amaral" → claim both Kindle and Paperback
4. Bio → paste English bio (below)
5. Photo → upload getulio_color_halfbody_1200.jpg from /fotos/
```

### Author Central Bio (English, ~700 chars)
```
Dr. Getulio Amaral Filho is a nephrologist and internal medicine physician with over 20 years of clinical practice. He earned his medical degree from the State University of Londrina (UEL) in 2004 and completed both internal medicine and nephrology training at Santa Casa de Londrina, where he now coordinates the Nephrology Residency program. He serves as Medical Director of the DaVita in-hospital dialysis unit in Londrina and lectures nationally on health, nephrology, and longevity. In 2026 he completed graduate training in Integrative Functional Medicine. From two decades of caring for patients who arrive with "normal" lab results yet are quietly aging fast, he developed the ACTS Method that structures BEFORE — his first book.
```

### Profile links
```
Website: https://drgetulioamaralfilho.com.br
Instagram: @drGetulioAmaralFilho
```

---

## E. A+ Content (Premium content blocks — depois de publicar)

KDP libera A+ Content quando o livro é approved. **Faz pra ambos os formatos**:

### Module 1 — "Cross-language banner" (módulo "Image with Text")
- Image: simple banner com a capa PT ao lado da capa EN
- Headline: `Also available in Portuguese`
- Body: `If you prefer reading in Portuguese, the original edition "ANTES" is available on Amazon.com.br.`
- Link: ASIN do PT digital

### Module 2 — "What you'll learn" (módulo "Single Image with Sidebar")
- Image: infographic do método ACTS (4 pilares)
- Body: bullets do que tem dentro do livro

### Module 3 — "About the author" (módulo "Premium Image and Text")
- Image: foto do Dr. Getulio HD
- Body: versão estendida da bio + credenciais (CRM-PR 21,876, Instagram, Plenya)

*Quando quiser, eu gero os PNGs nos tamanhos certos do A+ Content (970x600 hero, 300x300 thumbs).*

---

## F. Pre-flight checklist antes de clicar "Publish"

| Item | Onde verificar | OK? |
|---|---|---|
| EPUB passa no Kindle Previewer? | https://kdp.amazon.com/help/topic/G201283750 — baixar Previewer 3 e abrir o `Before-en.epub` | ☐ |
| PDF impresso passa no KDP Print Previewer? | Aba "Launch Previewer" depois de subir o PDF | ☐ |
| Cover impressa fits 410pp spine? | Verificar se o filename é `Before-en-print-cover-410pp.pdf` (não `400pp` etc.) | ☐ |
| ISBN válido (se usando próprio)? | https://www.cblservicos.org.br/ — buscar 978-65-975814-0-5 | ☐ |
| Description preview shows hook before "Read more"? | Após salvar o draft, ver na pré-visualização do KDP | ☐ |
| Categorias mostram em browse path "Aging"? | KDP mostra após selecionar | ☐ |
| Translated From field aponta pro ASIN PT correto? | Pegar ASIN do produto PT no Amazon | ☐ |
| Pricing USD vs BRL não causa arbitrage cross-border? | Comparar com PT Kindle ($/BRL) — diferença ≤30% é seguro | ☐ |

---

## G. Tempo estimado da sessão de upload

| Task | Tempo |
|---|---|
| Kindle eBook — preencher metadata | 12 min |
| Kindle eBook — upload + previewer | 8 min |
| Kindle eBook — pricing | 4 min |
| Kindle eBook — Submit | 1 min |
| Paperback — preencher metadata | 8 min (já vem pré-preenchido do Kindle) |
| Paperback — upload PDFs + previewer | 10 min |
| Paperback — pricing + distribution | 5 min |
| Paperback — Submit | 1 min |
| **Total** | **~50 min** |

Approval da Amazon: 24-72h pra Kindle, 72h pra Paperback.

---

*Última revisão: 2026-05-10. Quando atualizar a description ou keywords, atualizar também `md/pt-BR/kdp-description.md` (PT) pra manter os dois idiomas em paridade editorial.*
