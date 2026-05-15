// Plenya deck base v4 — primitivas hero-photo pra pitch deck premium.
// Régua: cada slide carrega 1 ideia + 1 imagem hero. Copy curta, texto que respira.
// Refs visuais: análise crítica de Function Health (premium clínico) calibrada pra PT-BR.
// Conteúdo: verbatim do site (Regra Zero — apps/site/messages/pt.json).

const path = require("path");

const PLENYA = {
  gold:        "B38645",
  goldSoft:    "D4A86B",
  petrol:      "063B4F",
  petrolDeep:  "041F2A",
  petrolMid:   "0A4A60",
  ocean:       "417E8E",
  sage:        "92B8B4",
  cream:       "EAE7DA",
  white:       "FFFFFF",
  ink:         "0A1F26",
  muted:       "5A6B70",
};

const FONTS = {
  serif: "Cormorant Garamond",
  sans:  "Inter",
};

const SITE_IMG  = path.resolve(__dirname, "../../apps/site/public/images");
const BRAND_IMG = path.resolve(__dirname, "../../apps/site/public/brand");
const DECK_IMG  = path.resolve(__dirname, "../../docs/decks/_assets/style-refs");

const ASSETS = {
  wordmarkCream: path.join(BRAND_IMG, "wordmark/cream.png"),
  wordmarkGold:  path.join(BRAND_IMG, "wordmark/gold.png"),
  symbolGold:    path.join(BRAND_IMG, "symbol/gold.png"),
  symbolOcean:   path.join(BRAND_IMG, "symbol/ocean.png"),
};

const W = 13.333;
const H = 7.5;

function setupPres(pres, { title, author = "Plenya", subject = "Continuum Plenya" } = {}) {
  pres.layout = "LAYOUT_WIDE";
  pres.title = title;
  pres.author = author;
  pres.company = "Plenya";
  pres.subject = subject;
}

// ============================================================
// Backgrounds
// ============================================================

function fillPetrol(slide) {
  slide.background = { color: PLENYA.petrol };
}

function fillPetrolDeep(slide) {
  slide.background = { color: PLENYA.petrolDeep };
}

function fillCream(slide) {
  slide.background = { color: PLENYA.cream };
}

// ============================================================
// Hero photo primitives — coração de cada slide
// ============================================================

// Foto cobrindo o slide inteiro com overlay petrol pra legibilidade
function heroPhotoFullBleed(pres, slide, { path: imgPath, overlayOpacity = 60 } = {}) {
  slide.addImage({ path: imgPath, x: 0, y: 0, w: W, h: H, sizing: { type: "cover", w: W, h: H } });
  // Overlay petrol gradient pra leitura de texto
  slide.addShape(pres.ShapeType.rect, {
    x: 0, y: 0, w: W, h: H,
    fill: { color: PLENYA.petrol, transparency: 100 - overlayOpacity },
    line: { type: "none" },
  });
}

// Foto na metade esquerda, texto na direita
function heroPhotoHalfLeft(pres, slide, { path: imgPath } = {}) {
  slide.background = { color: PLENYA.petrol };
  slide.addImage({ path: imgPath, x: 0, y: 0, w: W * 0.55, h: H, sizing: { type: "cover", w: W * 0.55, h: H } });
  // Gradient suave borda direita pra suavizar transição
  for (let i = 0; i < 8; i++) {
    slide.addShape(pres.ShapeType.rect, {
      x: W * 0.55 - 0.05 * (i + 1), y: 0, w: 0.05, h: H,
      fill: { color: PLENYA.petrol, transparency: 100 - (10 + i * 8) },
      line: { type: "none" },
    });
  }
}

// Foto na metade direita, texto na esquerda
function heroPhotoHalfRight(pres, slide, { path: imgPath } = {}) {
  slide.background = { color: PLENYA.petrol };
  slide.addImage({ path: imgPath, x: W * 0.45, y: 0, w: W * 0.55, h: H, sizing: { type: "cover", w: W * 0.55, h: H } });
  for (let i = 0; i < 8; i++) {
    slide.addShape(pres.ShapeType.rect, {
      x: W * 0.45 + 0.05 * i, y: 0, w: 0.05, h: H,
      fill: { color: PLENYA.petrol, transparency: 100 - (10 + i * 8) },
      line: { type: "none" },
    });
  }
}

// Foto cobrindo metade superior, texto na metade inferior
function heroPhotoTopHalf(pres, slide, { path: imgPath } = {}) {
  slide.background = { color: PLENYA.petrolDeep };
  slide.addImage({ path: imgPath, x: 0, y: 0, w: W, h: H * 0.55, sizing: { type: "cover", w: W, h: H * 0.55 } });
}

// ============================================================
// Logo primitives
// ============================================================

function addWordmark(slide, { x, y, w = 1.6, color = "cream" } = {}) {
  const asset = color === "gold" ? ASSETS.wordmarkGold : ASSETS.wordmarkCream;
  const width = w;
  const height = width / 6.756;
  slide.addImage({
    path: asset,
    x: x ?? 0.6, y: y ?? 0.4, w: width, h: height,
  });
}

function addWordmarkLarge(slide, { y = 2.5, w = 6.5, color = "cream" } = {}) {
  const asset = color === "gold" ? ASSETS.wordmarkGold : ASSETS.wordmarkCream;
  const width = w;
  const height = width / 6.756;
  slide.addImage({
    path: asset,
    x: (W - width) / 2, y, w: width, h: height,
  });
}

// ============================================================
// Typography
// ============================================================

function addEyebrow(slide, text, { x = 0.6, y = 0.4, w = 12, color = PLENYA.gold } = {}) {
  slide.addText(text.toUpperCase(), {
    x, y, w, h: 0.3,
    fontSize: 10, fontFace: FONTS.sans, bold: true,
    color, charSpacing: 6, margin: 0,
  });
}

// Headline gigante (poster style) — para 1 frase principal do slide
function addPosterHeadline(slide, text, { x = 0.7, y = 2.5, w = 12, color = PLENYA.cream, fontSize = 52, align = "left" } = {}) {
  const charsPerLine = Math.floor(w * 96 / (fontSize * 0.5));
  const lines = Math.ceil(text.length / charsPerLine);
  const h = lines * (fontSize / 72) * 1.15;
  slide.addText(text, {
    x, y, w, h: Math.max(1, h),
    fontSize, fontFace: FONTS.serif, bold: true,
    color, align, margin: 0, valign: "top",
  });
  return y + h + 0.15;
}

// Sub-headline italic serif
function addPosterSubhead(slide, text, { x = 0.7, y, w = 12, color = PLENYA.sage, fontSize = 20, align = "left" } = {}) {
  slide.addText(text, {
    x, y, w, h: 1.5,
    fontSize, fontFace: FONTS.serif, italic: true,
    color, align, margin: 0, valign: "top",
  });
  return y + 1.5;
}

// Body curta — só pra slides que precisam de mais contexto
function addBodyShort(slide, text, { x = 0.7, y, w = 7, color = PLENYA.cream, fontSize = 15 } = {}) {
  slide.addText(text, {
    x, y, w, h: 3,
    fontSize, fontFace: FONTS.sans,
    color, paraSpaceAfter: 8, margin: 0, valign: "top",
  });
}

// ============================================================
// Quote-as-poster (frase gigante, sem foto)
// ============================================================

function quotePoster(pres, slide, { eyebrow, quote, attribution, accent = "left" } = {}) {
  fillPetrol(slide);

  if (eyebrow) {
    slide.addText(eyebrow.toUpperCase(), {
      x: 0.7, y: 1.5, w: 12, h: 0.3,
      fontSize: 11, fontFace: FONTS.sans, bold: true,
      color: PLENYA.gold, charSpacing: 6, margin: 0,
    });
  }

  // Linha gold thin
  slide.addShape(pres.ShapeType.line, {
    x: 0.7, y: 2.1, w: 0.8, h: 0,
    line: { color: PLENYA.gold, width: 1.2 },
  });

  // Quote gigante serif
  slide.addText(quote, {
    x: 0.7, y: 2.4, w: 11.9, h: 4.0,
    fontSize: 60, fontFace: FONTS.serif, bold: false,
    color: PLENYA.cream, margin: 0, valign: "top",
  });

  if (attribution) {
    slide.addText(attribution, {
      x: 0.7, y: 6.4, w: 11.9, h: 0.4,
      fontSize: 14, fontFace: FONTS.sans, italic: true,
      color: PLENYA.sage, margin: 0,
    });
  }
}

// ============================================================
// Footer minimal
// ============================================================

function addFooter(pres, slide, slideNum, totalSlides, { onLight = false } = {}) {
  const color = onLight ? PLENYA.muted : PLENYA.sage;
  const goldColor = onLight ? PLENYA.gold : PLENYA.gold;
  slide.addText("CONTINUUM PLENYA", {
    x: 0.6, y: 7.15, w: 4, h: 0.25,
    fontSize: 8, fontFace: FONTS.sans, bold: true,
    color, charSpacing: 4, margin: 0,
  });
  slide.addText(`${String(slideNum).padStart(2, "0")} · ${String(totalSlides).padStart(2, "0")}`, {
    x: 8.7, y: 7.15, w: 4, h: 0.25,
    fontSize: 8, fontFace: FONTS.sans,
    color: goldColor, charSpacing: 3, align: "right", margin: 0,
  });
}

// ============================================================
// Number markers (timeline minimalista)
// ============================================================

function addPhaseNumber(slide, { x, y, n, color = PLENYA.gold }) {
  slide.addText(String(n).padStart(2, "0"), {
    x, y, w: 0.9, h: 0.6,
    fontSize: 36, fontFace: FONTS.serif, bold: false,
    color, margin: 0,
  });
}

// ============================================================
// Stat micro (não trust-strip — usar com critério)
// ============================================================

function addStatLine(slide, { x, y, w, big, label, color = PLENYA.cream }) {
  slide.addText(big, {
    x, y, w: 1.5, h: 0.5,
    fontSize: 26, fontFace: FONTS.serif, bold: true,
    color: PLENYA.gold, margin: 0,
  });
  slide.addText(label, {
    x: x + 1.6, y: y + 0.05, w: w - 1.7, h: 0.5,
    fontSize: 13, fontFace: FONTS.sans,
    color, margin: 0,
  });
}

module.exports = {
  PLENYA, FONTS, ASSETS, SITE_IMG, BRAND_IMG, DECK_IMG, W, H,
  setupPres,
  fillPetrol, fillPetrolDeep, fillCream,
  heroPhotoFullBleed, heroPhotoHalfLeft, heroPhotoHalfRight, heroPhotoTopHalf,
  addWordmark, addWordmarkLarge,
  addEyebrow, addPosterHeadline, addPosterSubhead, addBodyShort,
  quotePoster,
  addFooter,
  addPhaseNumber, addStatLine,
};
