// Card do protocolo de antibióticos na hemodiálise · A4 retrato, 2 páginas fixas.
// Uso: node render.js [--png] [--check]
//   sempre gera docs/clinica/protocolo-atb-dialise/card-protocolo-atb-hd.pdf (docs/**/*.pdf não é versionado)
//   --png    PNG por página em previews/ (aprovação), a partir do PDF (pdftoppm)
//   --check  avisa se alguma página transborda
const { chromium } = require("/home/user/plenya/scripts/deck-builder/continuum/node_modules/playwright");
const path = require("path"); const fs = require("fs"); const { execFileSync } = require("child_process");
const has = f => process.argv.slice(2).includes(`--${f}`);
const PDF_DIR = path.resolve(__dirname, "../../../docs/clinica/protocolo-atb-dialise");
const PREV = path.resolve(__dirname, "previews");
const PDF = `${PDF_DIR}/card-protocolo-atb-hd.pdf`;
(async () => {
  fs.mkdirSync(PDF_DIR, { recursive: true }); fs.mkdirSync(PREV, { recursive: true });
  const browser = await chromium.launch();
  const page = await browser.newPage({ viewport: { width: 794, height: 1123 } });
  await page.goto(`file://${path.resolve(__dirname, "card.html")}`, { waitUntil: "networkidle" });
  await page.evaluate(() => document.fonts.ready); await page.waitForTimeout(600);
  if (has("check")) {
    const res = await page.evaluate(() => [...document.querySelectorAll("section.page")].map((p, i) => {
      const top = p.getBoundingClientRect().top; let max = 0;
      p.querySelectorAll("*:not(.foot):not(.foot *)").forEach(e => { const b = e.getBoundingClientRect().bottom - top; if (b > max) max = b; });
      return { n: i + 1, mm: +(max * 25.4 / 96).toFixed(1) };
    }));
    res.forEach(r => console.log(`p${r.n} fim do conteudo ${r.mm} mm${r.mm > 286 ? "  ** TRANSBORDA (rodapé em 288)" : ""}`));
  }
  await page.pdf({ path: PDF, format: "A4", printBackground: true, preferCSSPageSize: true });
  await browser.close();
  if (has("png")) {
    execFileSync("pdftoppm", ["-png", "-r", "150", PDF, path.join(PREV, "card")]);
    console.log(`ok PNGs em ${PREV}`);
  }
  console.log(`ok ${PDF}`);
})();
