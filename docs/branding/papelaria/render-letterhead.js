// Render do papel timbrado: HTML -> PDF (vetorial, tamanho via @page) + PNG de preview.
// Usa o Playwright instalado na skill lecture-builder.
// Uso: node render-letterhead.js <slug>   (ex.: timbrado-branco)
const { chromium } = require("playwright");
const path = require("path");
const fs = require("fs");

const slug = process.argv[2];
if (!slug) { console.error("ERRO: informe o slug (ex.: timbrado-branco)"); process.exit(1); }

const DRAFTS = path.resolve(__dirname, "drafts");
const htmlPath = path.join(DRAFTS, `${slug}.html`);
const pdfPath = path.join(DRAFTS, `${slug}.pdf`);
const pngPath = path.join(DRAFTS, `${slug}.png`);
if (!fs.existsSync(htmlPath)) { console.error("não encontrei " + htmlPath); process.exit(1); }

(async () => {
  const browser = await chromium.launch();
  const context = await browser.newContext({ deviceScaleFactor: 2 });
  const page = await context.newPage();
  await page.goto(`file://${htmlPath}`, { waitUntil: "networkidle" });
  await page.evaluate(() => document.fonts.ready);
  await page.waitForTimeout(400);

  // PDF vetorial: preferCSSPageSize respeita @page { size:222mm 309mm }
  await page.pdf({
    path: pdfPath,
    printBackground: true,
    preferCSSPageSize: true,
    margin: { top: 0, right: 0, bottom: 0, left: 0 },
  });

  // PNG de preview (alta resolução do elemento .sheet)
  const el = await page.$(".sheet");
  await el.screenshot({ path: pngPath, scale: "device" });

  await browser.close();
  console.log(`✓ ${slug}: PDF + PNG`);
})();
