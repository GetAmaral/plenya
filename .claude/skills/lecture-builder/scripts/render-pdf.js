// Plenya · Lecture Builder — render HTML → PDF via Playwright headless Chromium
// Output: <dir>/lecture.pdf  (1920x1080, 16:9 Full HD, zero margin)
// Uso: node render-pdf.js --dir=/home/user/lecture-output/<slug>

const { chromium } = require("playwright");
const path = require("path");
const fs = require("fs");

const args = process.argv.slice(2);
const dirArg = args.find(a => a.startsWith("--dir="))?.split("=")[1];
const fileArg = args.find(a => a.startsWith("--file="))?.split("=")[1] || "lecture.html";
const outArg = args.find(a => a.startsWith("--out="))?.split("=")[1] || "lecture.pdf";

if (!dirArg) {
  console.error("ERRO: --dir=<output_dir> obrigatório");
  process.exit(1);
}
const baseDir = path.resolve(dirArg);
const htmlPath = path.resolve(baseDir, fileArg);
const outPath = path.resolve(baseDir, outArg);

if (!fs.existsSync(htmlPath)) {
  console.error("ERRO: arquivo não encontrado: " + htmlPath);
  process.exit(1);
}

(async () => {
  const browser = await chromium.launch();
  const context = await browser.newContext({
    viewport: { width: 1920, height: 1080 },
    deviceScaleFactor: 2,
  });
  const page = await context.newPage();

  await page.goto(`file://${htmlPath}`, { waitUntil: "networkidle" });
  await page.evaluate(() => document.fonts.ready);
  await page.waitForTimeout(800);

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
