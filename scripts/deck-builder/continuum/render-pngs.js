// Continuum Plenya — render per-slide PNG previews via Playwright
// Para cada <section class="slide"> em deck.html, screenshota 1920x1080
// pra previews/slide-XX.png. Workflow de iteração rápida — sem PDF.
//
// Uso: node render-pngs.js [--slide=03]   (ou todos)

const { chromium } = require("playwright");
const path = require("path");
const fs = require("fs");

const args = process.argv.slice(2);
const onlySlide = args.find(a => a.startsWith("--slide="))?.split("=")[1];
const fileArg = args.find(a => a.startsWith("--file="))?.split("=")[1] || "deck.html";
const outArg = args.find(a => a.startsWith("--out="))?.split("=")[1] || "previews";

(async () => {
  const browser = await chromium.launch();
  const context = await browser.newContext({
    viewport: { width: 1920, height: 1080 },
    deviceScaleFactor: 2,
  });
  const page = await context.newPage();

  const htmlPath = path.resolve(__dirname, fileArg);
  await page.goto(`file://${htmlPath}`, { waitUntil: "networkidle" });
  await page.evaluate(() => document.fonts.ready);
  await page.waitForTimeout(500);

  // Descobre todos os slides via DOM
  const slideIds = await page.$$eval("section.slide", sections =>
    sections.map(s => s.id || "")
  );

  const outDir = path.resolve(__dirname, outArg);
  if (!fs.existsSync(outDir)) fs.mkdirSync(outDir, { recursive: true });

  const target = onlySlide ? `s${String(onlySlide).padStart(2, "0")}` : null;

  for (const id of slideIds) {
    if (!id) continue;
    if (target && id !== target) continue;

    const num = id.replace("s", "");
    const outPath = path.join(outDir, `slide-${num}.png`);

    // Scrolla pro slide e captura só ele
    const section = await page.$(`section#${id}`);
    if (!section) {
      console.warn(`✗ section#${id} não encontrada`);
      continue;
    }

    await section.screenshot({ path: outPath, omitBackground: false });
    console.log(`✓ ${outPath}`);
  }

  await browser.close();
})();
