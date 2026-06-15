// Plenya · Lecture Builder — render per-slide PNG previews via Playwright
// Para cada <section class="slide"> em lecture.html, screenshota 1920x1080 → previews/slide-XX.png
// Uso: node render-pngs.js --dir=/home/user/lecture-output/<slug> [--slide=04]

const { chromium } = require("playwright");
const path = require("path");
const fs = require("fs");

const args = process.argv.slice(2);
const dirArg = args.find(a => a.startsWith("--dir="))?.split("=")[1];
const onlySlide = args.find(a => a.startsWith("--slide="))?.split("=")[1];
const fileArg = args.find(a => a.startsWith("--file="))?.split("=")[1] || "lecture.html";

if (!dirArg) {
  console.error("ERRO: --dir=<output_dir> obrigatório");
  process.exit(1);
}
const baseDir = path.resolve(dirArg);
const htmlPath = path.resolve(baseDir, fileArg);
if (!fs.existsSync(htmlPath)) {
  console.error("ERRO: arquivo não encontrado: " + htmlPath);
  process.exit(1);
}

const outDir = path.join(baseDir, "previews");
if (!fs.existsSync(outDir)) fs.mkdirSync(outDir, { recursive: true });

(async () => {
  const browser = await chromium.launch();
  const context = await browser.newContext({
    viewport: { width: 1920, height: 1080 },
    deviceScaleFactor: 2,
  });
  const page = await context.newPage();

  await page.goto(`file://${htmlPath}`, { waitUntil: "networkidle" });
  await page.evaluate(() => document.fonts.ready);
  await page.waitForTimeout(600);

  const slideIds = await page.$$eval("section.slide", sections =>
    sections.map(s => s.id || "")
  );

  const target = onlySlide ? `s${String(onlySlide).padStart(2, "0")}` : null;

  for (const id of slideIds) {
    if (!id) continue;
    if (target && id !== target) continue;
    const num = id.replace("s", "");
    const outPath = path.join(outDir, `slide-${num}.png`);
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
