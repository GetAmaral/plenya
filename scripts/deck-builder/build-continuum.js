// Continuum Plenya — deck de venda (companion da VSL).
// Régua editorial: copy verbatim do site (apps/site/messages/pt.json) + zero padrão americano transplantado.
// Régua visual: hero-photo por slide, copy curta, espaço pra respirar, autoridade calma.
//
// Imagens reais do site: apps/site/public/images/* (Getúlio, equipe, lifestyle, clínica)
// Imagens conceituais geradas via gpt-image-2: docs/decks/_assets/style-refs/deck-*.png

const path = require("path");
const pptxgen = require("pptxgenjs");
const {
  PLENYA, FONTS, SITE_IMG, DECK_IMG, W, H,
  setupPres,
  fillPetrol, fillPetrolDeep,
  heroPhotoFullBleed, heroPhotoHalfLeft, heroPhotoHalfRight,
  addWordmark,
  addEyebrow, addPosterHeadline,
  quotePoster,
  addFooter, addPhaseNumber,
} = require("./plenya-base");

const TOTAL = 12;
const pres = new pptxgen();
setupPres(pres, { title: "Continuum Plenya · Programa de acompanhamento contínuo" });

const IMG = {
  capa:       path.join(SITE_IMG, "getulio-consulta-online.jpg"),
  manifesto:  path.join(DECK_IMG, "deck-linha-continua-hero.png"),
  janela:     path.join(DECK_IMG, "deck-janela-silenciosa-hero.png"),
  retratos:   path.join(SITE_IMG, "lifestyle-movement.jpg"),
  equipe:     path.join(SITE_IMG, "team/equipe-formal.jpg"),
  agir:       path.join(SITE_IMG, "metodo-agir-hero.jpg"),
  jornadaPhotoA: path.join(SITE_IMG, "getulio-consulta.jpg"),
  jornadaPhotoB: path.join(SITE_IMG, "getulio-consultorio.jpg"),
  box:        path.join(DECK_IMG, "deck-box-plenya-content.png"),
  escore:     path.join(DECK_IMG, "deck-escore-curve-premium.png"),
  fechamento: path.join(SITE_IMG, "getulio-sorrindo.jpg"),
};

// ============================================================
// 01 — CAPA · foto Getúlio (real) + Continuum
// ============================================================
{
  const s = pres.addSlide();
  heroPhotoHalfLeft(pres, s, { path: IMG.capa });

  // Wordmark canto sup esquerdo
  addWordmark(s, { x: 0.6, y: 0.45, w: 1.4 });

  // Eyebrow gold
  s.addText("PROGRAMA DE ACOMPANHAMENTO CONTÍNUO", {
    x: 7.6, y: 2.4, w: 5.3, h: 0.3,
    fontSize: 11, fontFace: FONTS.sans, bold: true,
    color: PLENYA.gold, charSpacing: 6, margin: 0,
  });

  // Title gigante
  s.addText("Continuum.", {
    x: 7.6, y: 2.9, w: 5.3, h: 1.5,
    fontSize: 88, fontFace: FONTS.serif, bold: true,
    color: PLENYA.cream, margin: 0, valign: "top",
  });

  // Linha gold
  s.addShape(pres.ShapeType.line, {
    x: 7.6, y: 4.65, w: 1, h: 0,
    line: { color: PLENYA.gold, width: 1.5 },
  });

  // Sub-claim (verbatim hero)
  s.addText("Saúde não se resolve em uma consulta. Se constrói no tempo, em conjunto, com método.", {
    x: 7.6, y: 4.9, w: 5.3, h: 2.0,
    fontSize: 18, fontFace: FONTS.serif, italic: true,
    color: PLENYA.sage, margin: 0, valign: "top",
  });
}

// ============================================================
// 02 — MANIFESTO · linha contínua (hero gerado) + quote do site
// ============================================================
{
  const s = pres.addSlide();
  heroPhotoFullBleed(pres, s, { path: IMG.manifesto, overlayOpacity: 55 });

  addEyebrow(s, "Plenya", { y: 0.45 });

  // Frase verbatim do brandbook/home
  s.addText("Plenitude não é um ponto de chegada.", {
    x: 0.7, y: 3.0, w: 11.9, h: 0.9,
    fontSize: 48, fontFace: FONTS.serif, bold: false,
    color: PLENYA.cream, margin: 0,
  });
  s.addText("É uma linha contínua.", {
    x: 0.7, y: 3.85, w: 11.9, h: 0.9,
    fontSize: 48, fontFace: FONTS.serif, bold: false, italic: true,
    color: PLENYA.gold, margin: 0,
  });

  // Footer mini
  s.addText("CONTINUUM", {
    x: 0.6, y: 7.15, w: 4, h: 0.25,
    fontSize: 8, fontFace: FONTS.sans, bold: true,
    color: PLENYA.sage, charSpacing: 4, margin: 0,
  });
  s.addText("02 · 12", {
    x: 8.7, y: 7.15, w: 4, h: 0.25,
    fontSize: 8, fontFace: FONTS.sans,
    color: PLENYA.gold, charSpacing: 3, align: "right", margin: 0,
  });
}

// ============================================================
// 03 — O PROBLEMA · janela silenciosa (hero gerado) + headline
// ============================================================
{
  const s = pres.addSlide();
  heroPhotoHalfRight(pres, s, { path: IMG.janela });

  addEyebrow(s, "O problema", { x: 0.7, y: 0.55 });

  s.addText("O check-up convencional não procura", {
    x: 0.7, y: 2.0, w: 6.3, h: 0.8,
    fontSize: 32, fontFace: FONTS.serif, bold: true,
    color: PLENYA.cream, margin: 0,
  });
  s.addText("o que vai te tirar do jogo.", {
    x: 0.7, y: 2.75, w: 6.3, h: 0.8,
    fontSize: 32, fontFace: FONTS.serif, bold: true,
    color: PLENYA.cream, margin: 0,
  });
  s.addText("Quando ele encontra,", {
    x: 0.7, y: 3.55, w: 6.3, h: 0.7,
    fontSize: 30, fontFace: FONTS.serif, italic: true,
    color: PLENYA.gold, margin: 0,
  });
  s.addText("você já está doente.", {
    x: 0.7, y: 4.25, w: 6.3, h: 0.7,
    fontSize: 30, fontFace: FONTS.serif, italic: true,
    color: PLENYA.gold, margin: 0,
  });

  // Sub destaque sutil
  s.addShape(pres.ShapeType.line, {
    x: 0.7, y: 5.4, w: 0.6, h: 0,
    line: { color: PLENYA.gold, width: 1.5 },
  });
  s.addText("Entre o exame que diz \"normal\" e o diagnóstico que muda sua vida existe uma janela silenciosa de dez a vinte anos.", {
    x: 0.7, y: 5.6, w: 6.3, h: 1.4,
    fontSize: 12, fontFace: FONTS.sans,
    color: PLENYA.cream, margin: 0,
  });

  addFooter(pres, s, 3, TOTAL);
}

// ============================================================
// 04 — PARA QUEM É · lifestyle full-bleed + retratos curtos
// ============================================================
{
  const s = pres.addSlide();
  heroPhotoFullBleed(pres, s, { path: IMG.retratos, overlayOpacity: 65 });

  addEyebrow(s, "Reconhecimento", { y: 0.55 });

  s.addText("Não é para todo mundo.", {
    x: 0.7, y: 1.1, w: 11.9, h: 0.9,
    fontSize: 44, fontFace: FONTS.serif, bold: true,
    color: PLENYA.cream, margin: 0,
  });

  s.addText("É para quem reconhece um destes momentos.", {
    x: 0.7, y: 2.1, w: 11.9, h: 0.5,
    fontSize: 18, fontFace: FONTS.serif, italic: true,
    color: PLENYA.sage, margin: 0,
  });

  // 6 retratos em grid 3x2 (titles curtos só, copy completa no falado)
  const retratos = [
    "Faz tudo certo, e os sinais não somam.",
    "Recebeu o susto.",
    "Carrega mais do que si mesmo.",
    "Está numa transição que ninguém nomeia.",
    "Informação demais, integração de menos.",
    "Quer estar inteiro para quem ainda vai chegar.",
  ];

  const colW = 4.0, gap = 0.18;
  retratos.forEach((r, i) => {
    const col = i % 3;
    const row = Math.floor(i / 3);
    const x = 0.7 + col * (colW + gap);
    const y = 3.4 + row * 1.65;

    // Número grande gold
    s.addText(String(i + 1).padStart(2, "0"), {
      x, y, w: 0.8, h: 0.5,
      fontSize: 20, fontFace: FONTS.serif, bold: false,
      color: PLENYA.gold, margin: 0,
    });
    // Linha vertical gold thin
    s.addShape(pres.ShapeType.line, {
      x: x + 0.7, y: y + 0.08, w: 0, h: 1.4,
      line: { color: PLENYA.gold, width: 0.5, transparency: 50 },
    });
    // Frase retrato
    s.addText(r, {
      x: x + 0.9, y, w: colW - 1.0, h: 1.5,
      fontSize: 13, fontFace: FONTS.serif, bold: true,
      color: PLENYA.cream, margin: 0,
    });
  });

  addFooter(pres, s, 4, TOTAL);
}

// ============================================================
// 05 — A RESPOSTA · quote poster (sem foto) — frase chave
// ============================================================
{
  const s = pres.addSlide();
  quotePoster(pres, s, {
    eyebrow: "A resposta",
    quote: "Há quem procure saúde quando algo já apareceu. O Continuum existe para quem decidiu o caminho contrário.",
    attribution: "Plenya · Continuum",
  });
  addFooter(pres, s, 5, TOTAL);
}

// ============================================================
// 06 — EQUIPE · foto real + headline + frase oficial
// ============================================================
{
  const s = pres.addSlide();
  heroPhotoFullBleed(pres, s, { path: IMG.equipe, overlayOpacity: 65 });

  addEyebrow(s, "A equipe", { y: 0.55 });

  s.addText("Quatro especialistas.", {
    x: 0.7, y: 1.4, w: 11.9, h: 1.0,
    fontSize: 54, fontFace: FONTS.serif, bold: true,
    color: PLENYA.cream, margin: 0,
  });
  s.addText("Um plano único.", {
    x: 0.7, y: 2.4, w: 11.9, h: 1.0,
    fontSize: 54, fontFace: FONTS.serif, italic: true,
    color: PLENYA.gold, margin: 0,
  });

  // Frase oficial menor
  s.addShape(pres.ShapeType.line, {
    x: 0.7, y: 4.85, w: 0.8, h: 0,
    line: { color: PLENYA.gold, width: 1.5 },
  });
  s.addText("Médico, nutricionista, psicóloga e educador físico se reúnem antes do seu primeiro encontro, discutem o seu caso e desenham um plano integrado. Cada decisão clínica conversa com as outras.", {
    x: 0.7, y: 5.05, w: 8.5, h: 1.8,
    fontSize: 14, fontFace: FONTS.sans,
    color: PLENYA.cream, margin: 0,
  });

  addFooter(pres, s, 6, TOTAL);
}

// ============================================================
// 07 — MÉTODO AGIR · hero foto + 4 letras
// ============================================================
{
  const s = pres.addSlide();
  heroPhotoFullBleed(pres, s, { path: IMG.agir, overlayOpacity: 70 });

  addEyebrow(s, "Método", { y: 0.55 });

  s.addText("AGIR.", {
    x: 0.7, y: 1.0, w: 11.9, h: 1.3,
    fontSize: 80, fontFace: FONTS.serif, bold: true,
    color: PLENYA.cream, margin: 0,
  });

  s.addText("Quatro pilares interdependentes para uma saúde que se sustenta no tempo.", {
    x: 0.7, y: 2.4, w: 9, h: 0.6,
    fontSize: 18, fontFace: FONTS.serif, italic: true,
    color: PLENYA.sage, margin: 0,
  });

  // Grid 4 pilares mini-cards
  const pilares = [
    { l: "A", nome: "Atividade, alimentação e suplementação", quem: "Nutricionista + educador físico" },
    { l: "G", nome: "Gestão clínica e metabólica", quem: "Médico" },
    { l: "I", nome: "Integração mente-corpo", quem: "Psicóloga" },
    { l: "R", nome: "Ritmo circadiano e repouso", quem: "Atravessa os quatro" },
  ];

  const cellW = 2.95, gap = 0.15;
  pilares.forEach((p, i) => {
    const x = 0.7 + i * (cellW + gap);
    const y = 4.0;

    // Letra gigante gold
    s.addText(p.l, {
      x, y, w: 1.5, h: 1.8,
      fontSize: 96, fontFace: FONTS.serif, bold: true,
      color: PLENYA.gold, margin: 0,
    });
    // Linha gold thin
    s.addShape(pres.ShapeType.line, {
      x, y: y + 1.9, w: 0.6, h: 0,
      line: { color: PLENYA.gold, width: 1 },
    });
    // Nome do pilar
    s.addText(p.nome, {
      x, y: y + 2.05, w: cellW - 0.1, h: 0.7,
      fontSize: 12, fontFace: FONTS.serif, bold: true,
      color: PLENYA.cream, margin: 0,
    });
    // Quem conduz
    s.addText(p.quem, {
      x, y: y + 2.75, w: cellW - 0.1, h: 0.4,
      fontSize: 9, fontFace: FONTS.sans,
      color: PLENYA.sage, italic: true, margin: 0,
    });
  });

  addFooter(pres, s, 7, TOTAL);
}

// ============================================================
// 08 — A JORNADA · 6 fases em timeline editorial
// ============================================================
{
  const s = pres.addSlide();
  fillPetrolDeep(s);

  addEyebrow(s, "A cronologia", { y: 0.55 });

  s.addText("Do tempo zero ao próximo horizonte.", {
    x: 0.7, y: 1.0, w: 11.9, h: 0.9,
    fontSize: 40, fontFace: FONTS.serif, bold: true,
    color: PLENYA.cream, margin: 0,
  });

  const fases = [
    { phase: "Semana 1",          title: "Quatro consultas online.",                  body: "Uma com cada profissional. Juntas, preenchem o Escore Plenya." },
    { phase: "Semana 2",          title: "Reunião da equipe e abertura do plano.",    body: "Os quatro discutem o caso. O médico devolve a leitura clínica integrada." },
    { phase: "Semanas seguintes", title: "Encontro semanal, em rotação.",             body: "A cada quatro semanas o ciclo se completa. Nenhum pilar fica sem revisão." },
    { phase: "Conforme o plano", title: "Box Plenya.",                                body: "Mimos selecionados, suplementos e manipulados do seu protocolo." },
    { phase: "A cada três meses", title: "Reavaliação do Escore.",                    body: "A curva mostra o que avançou, o que estagnou e onde a próxima intervenção entra." },
    { phase: "Final do ciclo",    title: "Avaliação final e próximo horizonte.",      body: "Ao fim do período, decidimos juntos o desenho do próximo ciclo." },
  ];

  const colW = 4.0, gap = 0.15, rowH = 2.05;
  fases.forEach((f, i) => {
    const col = i % 3;
    const row = Math.floor(i / 3);
    const x = 0.7 + col * (colW + gap);
    const y = 2.4 + row * rowH;

    // Número em gold serifa
    s.addText(String(i + 1).padStart(2, "0"), {
      x, y, w: 0.9, h: 0.5,
      fontSize: 24, fontFace: FONTS.serif, bold: false,
      color: PLENYA.gold, margin: 0,
    });
    // Fase tag uppercase
    s.addText(f.phase.toUpperCase(), {
      x: x + 0.95, y: y + 0.13, w: colW - 1.0, h: 0.25,
      fontSize: 9, fontFace: FONTS.sans, bold: true,
      color: PLENYA.gold, charSpacing: 3, margin: 0,
    });
    // Linha gold thin abaixo do número
    s.addShape(pres.ShapeType.line, {
      x, y: y + 0.6, w: colW - 0.1, h: 0,
      line: { color: PLENYA.gold, width: 0.5, transparency: 50 },
    });
    // Title
    s.addText(f.title, {
      x, y: y + 0.7, w: colW - 0.1, h: 0.55,
      fontSize: 15, fontFace: FONTS.serif, bold: true,
      color: PLENYA.cream, margin: 0,
    });
    // Body
    s.addText(f.body, {
      x, y: y + 1.25, w: colW - 0.1, h: 0.75,
      fontSize: 10, fontFace: FONTS.sans,
      color: PLENYA.sage, margin: 0,
    });
  });

  addFooter(pres, s, 8, TOTAL);
}

// ============================================================
// 09 — BOX PLENYA · hero foto gerada
// ============================================================
{
  const s = pres.addSlide();
  heroPhotoHalfRight(pres, s, { path: IMG.box });

  addEyebrow(s, "Entrega", { y: 0.55 });

  s.addText("Um box chega até você.", {
    x: 0.7, y: 1.6, w: 6.3, h: 1.2,
    fontSize: 42, fontFace: FONTS.serif, bold: true,
    color: PLENYA.cream, margin: 0,
  });

  s.addShape(pres.ShapeType.line, {
    x: 0.7, y: 3.0, w: 0.6, h: 0,
    line: { color: PLENYA.gold, width: 1.5 },
  });

  s.addText("Mimos selecionados pela equipe, suplementos e manipulados específicos do seu protocolo. À medida que o cuidado evolui e o plano se ajusta, novos boxes seguem o mesmo ritmo.", {
    x: 0.7, y: 3.25, w: 6.3, h: 2.6,
    fontSize: 14, fontFace: FONTS.sans,
    color: PLENYA.cream, margin: 0,
  });

  s.addText("Cada caixa carrega o manifesto Plenya em ouro foil.", {
    x: 0.7, y: 6.0, w: 6.3, h: 0.4,
    fontSize: 11, fontFace: FONTS.serif, italic: true,
    color: PLENYA.gold, margin: 0,
  });

  addFooter(pres, s, 9, TOTAL);
}

// ============================================================
// 10 — ESCORE PLENYA · dashboard hero
// ============================================================
{
  const s = pres.addSlide();
  heroPhotoHalfLeft(pres, s, { path: IMG.escore });

  addEyebrow(s, "Instrumento", { x: 7.6, y: 1.0 });

  s.addText("Escore Plenya.", {
    x: 7.6, y: 1.4, w: 5.3, h: 0.9,
    fontSize: 42, fontFace: FONTS.serif, bold: true,
    color: PLENYA.cream, margin: 0,
  });

  s.addText("A curva que mostra se o cuidado está funcionando.", {
    x: 7.6, y: 2.4, w: 5.3, h: 1.1,
    fontSize: 18, fontFace: FONTS.serif, italic: true,
    color: PLENYA.sage, margin: 0,
  });

  s.addShape(pres.ShapeType.line, {
    x: 7.6, y: 3.7, w: 0.6, h: 0,
    line: { color: PLENYA.gold, width: 1.5 },
  });

  s.addText("Mais de 800 itens — história, sintomas, exames, hábitos, medicamentos — consolidados em uma pontuação clara, evolutiva e personalizada. Refeita a cada três meses.", {
    x: 7.6, y: 3.95, w: 5.3, h: 2.5,
    fontSize: 13, fontFace: FONTS.sans,
    color: PLENYA.cream, margin: 0,
  });

  s.addText("Direção, não punição.", {
    x: 7.6, y: 6.3, w: 5.3, h: 0.4,
    fontSize: 12, fontFace: FONTS.serif, italic: true,
    color: PLENYA.gold, margin: 0,
  });

  addFooter(pres, s, 10, TOTAL);
}

// ============================================================
// 11 — MODALIDADES
// ============================================================
{
  const s = pres.addSlide();
  fillPetrol(s);

  addEyebrow(s, "Modalidades", { y: 0.55 });

  s.addText("Dois horizontes de compromisso.", {
    x: 0.7, y: 1.0, w: 11.9, h: 1.0,
    fontSize: 44, fontFace: FONTS.serif, bold: true,
    color: PLENYA.cream, margin: 0,
  });

  const mods = [
    { period: "6 meses",  name: "Continuum Semestral", body: "Ciclo completo: avaliação inicial, plano integrado, encontros semanais, reavaliação trimestral e fechamento.", highlight: false },
    { period: "12 meses", name: "Continuum Anual",     body: "Mesma estrutura, com mais uma reavaliação trimestral e maior consolidação dos hábitos no tempo.",              highlight: true  },
  ];

  const cardW = 5.85, gap = 0.4, cardH = 3.6;
  const startX = (W - (cardW * 2 + gap)) / 2;
  mods.forEach((m, i) => {
    const x = startX + i * (cardW + gap);
    const y = 2.85;
    const accent = m.highlight ? PLENYA.gold : PLENYA.ocean;

    // Border-top
    s.addShape(pres.ShapeType.line, {
      x, y, w: cardW, h: 0,
      line: { color: accent, width: m.highlight ? 2 : 0.75 },
    });

    s.addText(m.period.toUpperCase(), {
      x, y: y + 0.25, w: cardW, h: 0.3,
      fontSize: 10, fontFace: FONTS.sans, bold: true,
      color: accent, charSpacing: 5, margin: 0,
    });

    s.addText(m.name, {
      x, y: y + 0.65, w: cardW, h: 1.0,
      fontSize: 42, fontFace: FONTS.serif, bold: true,
      color: PLENYA.cream, margin: 0,
    });

    s.addText(m.body, {
      x, y: y + 1.85, w: cardW, h: 1.4,
      fontSize: 13, fontFace: FONTS.sans,
      color: PLENYA.cream, margin: 0,
    });
  });

  // Valores
  s.addShape(pres.ShapeType.line, {
    x: 0.7, y: 6.5, w: 11.9, h: 0,
    line: { color: PLENYA.gold, width: 0.5, transparency: 70 },
  });
  s.addText("Valores · sob consulta. A equipe apresenta as condições em conversa direta.", {
    x: 0.7, y: 6.65, w: 11.9, h: 0.4,
    fontSize: 11, fontFace: FONTS.serif, italic: true,
    color: PLENYA.sage, align: "center", margin: 0,
  });

  addFooter(pres, s, 11, TOTAL);
}

// ============================================================
// 12 — CTA · fechamento + claim
// ============================================================
{
  const s = pres.addSlide();
  heroPhotoFullBleed(pres, s, { path: IMG.fechamento, overlayOpacity: 75 });

  // Wordmark gold grande centralizado
  const wordmarkW = 5.5;
  const wordmarkH = wordmarkW / 6.756;
  s.addImage({
    path: path.join(SITE_IMG, "../brand/wordmark/gold.png"),
    x: (W - wordmarkW) / 2, y: 1.9, w: wordmarkW, h: wordmarkH,
  });

  s.addShape(pres.ShapeType.line, {
    x: W / 2 - 1, y: 3.0, w: 2, h: 0,
    line: { color: PLENYA.gold, width: 1.5 },
  });

  s.addText("Viva bem, viva mais.", {
    x: 0.7, y: 3.25, w: 11.9, h: 1.1,
    fontSize: 56, fontFace: FONTS.serif, italic: true,
    color: PLENYA.cream, align: "center", margin: 0,
  });

  s.addText("Conversar com a equipe Plenya.", {
    x: 0.7, y: 4.7, w: 11.9, h: 0.7,
    fontSize: 22, fontFace: FONTS.serif, bold: true,
    color: PLENYA.gold, align: "center", margin: 0,
  });

  s.addText("plenyasaude.com.br   ·   Londrina · PR   ·   contato@plenyasaude.com.br", {
    x: 0.7, y: 6.4, w: 11.9, h: 0.3,
    fontSize: 11, fontFace: FONTS.sans,
    color: PLENYA.cream, charSpacing: 3, align: "center", margin: 0,
  });
}

// ============================================================
// Salvar
// ============================================================
const today = new Date().toISOString().slice(0, 10).replace(/-/g, "");
const outPath = path.resolve(__dirname, `../../docs/decks/continuum-plenya-${today}.pptx`);

pres.writeFile({ fileName: outPath })
  .then(name => {
    console.log(`✓ Deck gerado: ${name}`);
    console.log(`  Slides: ${TOTAL}`);
    console.log(`  Estilo: hero-photo pitch deck premium`);
  })
  .catch(err => {
    console.error("✗ Erro:", err);
    process.exit(1);
  });
