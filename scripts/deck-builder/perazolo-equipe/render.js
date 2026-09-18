// Plenya × Clínica Tatianna Perazolo · deck para a EQUIPE da clínica.
// 16:9 (1920×1080) → PNG por slide + PDF.
// Uso: node render.js [--check] [--png] [--slide=NN] [--pdf]
const { chromium } = require("/home/user/plenya/scripts/deck-builder/continuum/node_modules/playwright");
const path = require("path"); const fs = require("fs");
const args = process.argv.slice(2); const has = f => args.includes(`--${f}`);
const slideArg = (args.find(a => a.startsWith("--slide=")) || "").split("=")[1];
(async () => {
  const browser = await chromium.launch();
  const page = await browser.newPage({ viewport: { width: 1920, height: 1080 }, deviceScaleFactor: 1 });
  await page.goto(`file://${path.resolve(__dirname, "deck.html")}`, { waitUntil: "networkidle" });
  await page.evaluate(() => document.fonts.ready); await page.waitForTimeout(800);
  if (has("check")) {
    const res = await page.evaluate(() => [...document.querySelectorAll("section.slide")].map((s, i) => {
      const top = s.getBoundingClientRect().top; let max = 0;
      s.querySelectorAll(".body *").forEach(e => {
        if (e.classList.contains("content")) return;
        if (!e.textContent.trim() && !e.className.match(/seg|bar|band|rule|mark/)) return;
        const b = e.getBoundingClientRect().bottom - top; if (b > max) max = b;
      });
      const small = [...s.querySelectorAll(".body p,.body li,.body td,.body h2,.body h3")]
        .filter(e => e.textContent.trim() && !e.closest(".note,.caps,.eyebrow,.disc")
          && parseFloat(getComputedStyle(e).fontSize) < 36).length;
      return { n: i + 1, px: Math.round(max), small };
    }));
    res.forEach(r => console.log(`s${String(r.n).padStart(2,"0")} fim ${r.px}px${r.px > 980 ? "  ⚠ TRANSBORDA" : ""}${r.small ? `  ⚠ ${r.small} elem < 36px` : ""}`));
  }
  if (has("png") || slideArg) {
    const out = path.resolve(__dirname, "previews"); fs.mkdirSync(out, { recursive: true });
    const secs = await page.$$("section.slide");
    for (let i = 0; i < secs.length; i++) {
      const n = String(i + 1).padStart(2, "0"); if (slideArg && slideArg !== n) continue;
      await secs[i].screenshot({ path: `${out}/s-${n}.png` });
    }
    console.log(`✓ PNGs em ${out}`);
  }
  if (has("pdf")) {
    const out = path.resolve(__dirname, "../../../docs/parcerias/tatianna-perazolo/apresentacao-equipe-plenya-perazolo.pdf");
    await page.pdf({ path: out, width: "1920px", height: "1080px", printBackground: true, preferCSSPageSize: true, pageRanges: "1-", margin: { top: 0, right: 0, bottom: 0, left: 0 } });
    console.log(`✓ PDF: ${out}`);
  }
  await browser.close();
})();
