// Card do paciente · biópsia renal · A4 retrato -> PNG (aprovação) e PDF.
// Uso: node render.js [--png] [--pdf] [--check]
// O PDF vai para docs/clinica/biopsia-renal/ (docs/**/*.pdf não é versionado).
const { chromium } = require("/home/user/plenya/scripts/deck-builder/continuum/node_modules/playwright");
const path = require("path"); const fs = require("fs");
const args = process.argv.slice(2); const has = f => args.includes(`--${f}`);
(async () => {
  const browser = await chromium.launch();
  const page = await browser.newPage({ viewport: { width: 794, height: 1123 }, deviceScaleFactor: 2 });
  await page.goto(`file://${path.resolve(__dirname, "card.html")}`, { waitUntil: "networkidle" });
  await page.evaluate(() => document.fonts.ready); await page.waitForTimeout(600);
  if (has("check")) {
    const r = await page.evaluate(() => {
      const p = document.querySelector("section.page"), top = p.getBoundingClientRect().top; let max = 0;
      p.querySelectorAll("*:not(.foot):not(.foot *)").forEach(e => { const b = e.getBoundingClientRect().bottom - top; if (b > max) max = b; });
      return +(max * 25.4 / 96).toFixed(1);
    });
    console.log(`fim do conteudo ${r} mm${r > 287 ? "  ** TRANSBORDA (rodapé em 290)" : ""}`);
  }
  if (has("png")) {
    const out = path.resolve(__dirname, "previews"); fs.mkdirSync(out, { recursive: true });
    await (await page.$("section.page")).screenshot({ path: `${out}/card.png` });
    console.log(`ok ${out}/card.png`);
  }
  if (has("pdf")) {
    const dir = path.resolve(__dirname, "../../../docs/clinica/biopsia-renal"); fs.mkdirSync(dir, { recursive: true });
    await page.pdf({ path: `${dir}/card-paciente-biopsia-renal.pdf`, format: "A4", printBackground: true, preferCSSPageSize: true });
    console.log(`ok ${dir}/card-paciente-biopsia-renal.pdf`);
  }
  await browser.close();
})();
