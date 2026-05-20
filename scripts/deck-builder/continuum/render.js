// Continuum Plenya — render HTML → PDF via Playwright headless Chromium
// Output: docs/decks/continuum-plenya-YYYYMMDD.pdf
// Format: 1920x1080 (16:9 Full HD), zero margin, fonts embedded via Google Fonts.

const { chromium } = require("playwright");
const path = require("path");

(async () => {
  const browser = await chromium.launch();
  const context = await browser.newContext({
    viewport: { width: 1920, height: 1080 },
    deviceScaleFactor: 2,
  });
  const page = await context.newPage();

  const fileArg = process.argv.slice(2).find(a => a.startsWith("--file="))?.split("=")[1] || "deck.html";
  const htmlPath = path.resolve(__dirname, fileArg);
  await page.goto(`file://${htmlPath}`, { waitUntil: "networkidle" });

  // Aguarda fontes Google carregarem (evita FOUT no PDF)
  await page.evaluate(() => document.fonts.ready);
  await page.waitForTimeout(800);

  const today = new Date().toISOString().slice(0, 10).replace(/-/g, "");
  const outArg = process.argv.slice(2).find(a => a.startsWith("--out="))?.split("=")[1];
  const outPath = outArg
    ? path.resolve(__dirname, outArg)
    : path.resolve(__dirname, `../../../docs/decks/continuum-plenya-${today}.pdf`);

  // PDF 1920x1080 — formato exato 16:9, margens zero, com background
  await page.pdf({
    path: outPath,
    width: "1920px",
    height: "1080px",
    printBackground: true,
    preferCSSPageSize: false,
    margin: { top: 0, right: 0, bottom: 0, left: 0 },
  });

  await browser.close();
  console.log(`✓ PDF gerado: ${outPath}`);
})();
