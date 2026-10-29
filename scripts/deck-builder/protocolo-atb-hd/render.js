// Card do protocolo de antibióticos na hemodiálise · uma folha A4 por <section class="page">,
// cada uma no tamanho do seu data-size ("portrait" ou "landscape").
// Uso: node render.js [--png] [--check]
//   sempre gera docs/clinica/protocolo-atb-dialise/card-protocolo-atb-hd.pdf (docs/**/*.pdf não é versionado)
//   --png    PNG por folha em previews/ (aprovação), a partir do PDF (pdftoppm)
//   --check  avisa se alguma folha transborda (rodapé a 6 mm da borda)
const { chromium } = require("/home/user/plenya/scripts/deck-builder/continuum/node_modules/playwright");
const path = require("path"); const fs = require("fs"); const { execFileSync } = require("child_process");
const has = f => process.argv.slice(2).includes(`--${f}`);
const PDF_DIR = path.resolve(__dirname, "../../../docs/clinica/protocolo-atb-dialise");
const PREV = path.resolve(__dirname, "previews");
const PDF = `${PDF_DIR}/card-protocolo-atb-hd.pdf`;
const SIZE = { portrait: { w: 210, h: 297 }, landscape: { w: 297, h: 210 } };
(async () => {
  fs.mkdirSync(PDF_DIR, { recursive: true }); fs.mkdirSync(PREV, { recursive: true });
  const browser = await chromium.launch();
  const page = await browser.newPage({ viewport: { width: 1123, height: 1123 } });
  await page.goto(`file://${path.resolve(__dirname, "card.html")}`, { waitUntil: "networkidle" });
  await page.evaluate(() => document.fonts.ready); await page.waitForTimeout(600);
  const sizes = await page.evaluate(() => [...document.querySelectorAll("section.page")].map(p => p.dataset.size || "landscape"));
  if (has("check")) {
    const res = await page.evaluate(() => [...document.querySelectorAll("section.page")].map((p, i) => {
      const top = p.getBoundingClientRect().top; let max = 0;
      p.querySelectorAll("*:not(.foot):not(.foot *)").forEach(e => { const b = e.getBoundingClientRect().bottom - top; if (b > max) max = b; });
      return { n: i + 1, mm: +(max * 25.4 / 96).toFixed(1), h: p.dataset.size === "portrait" ? 297 : 210 };
    }));
    res.forEach(r => console.log(`p${r.n} fim do conteudo ${r.mm} mm de ${r.h}${r.mm > r.h - 8 ? "  ** TRANSBORDA" : ""}`));
  }
  const parts = [];
  for (let i = 0; i < sizes.length; i++) {
    await page.evaluate(i => document.querySelectorAll("section.page").forEach((p, k) => { p.style.display = k === i ? "" : "none"; }), i);
    const { w, h } = SIZE[sizes[i]];
    const out = path.join(PREV, `.part-${i + 1}.pdf`);
    await page.pdf({ path: out, width: `${w}mm`, height: `${h}mm`, printBackground: true, margin: { top: 0, right: 0, bottom: 0, left: 0 } });
    parts.push(out);
  }
  await browser.close();
  execFileSync("pdfunite", [...parts, PDF]);
  parts.forEach(f => fs.rmSync(f));
  if (has("png")) {
    execFileSync("pdftoppm", ["-png", "-r", "150", PDF, path.join(PREV, "card")]);
    console.log(`ok PNGs em ${PREV}`);
  }
  console.log(`ok ${PDF}`);
})();
