// Renderiza um documento A4 (HTML → PDF vetorial) via Playwright/Chromium.
// Uso: node render-a4.js --file=planilha-medico.html --out=../../../docs/continuum/arquivo.pdf
const { chromium } = require("playwright");
const path = require("path");

(async () => {
  const arg = n => process.argv.slice(2).find(a => a.startsWith(`--${n}=`))?.split("=")[1];
  const file = arg("file") || "planilha-medico.html";
  const out  = path.resolve(__dirname, arg("out") || "documento.pdf");

  const browser = await chromium.launch();
  const page = await browser.newPage();
  await page.goto(`file://${path.resolve(__dirname, file)}`, { waitUntil: "networkidle" });
  await page.evaluate(() => document.fonts.ready);
  await page.waitForTimeout(600);

  const foot = `<div style="width:100%;font-family:Inter,system-ui,sans-serif;font-size:6.6pt;color:#5A6B70;
      padding:0 18mm;display:flex;justify-content:space-between;align-items:center;">
      <span>Plenya &middot; Continuum M&eacute;dico &middot; estrutura de custos &middot; 08/09/2026</span>
      <span>Documento interno &nbsp;|&nbsp; <span class="pageNumber"></span>/<span class="totalPages"></span></span>
    </div>`;

  await page.pdf({
    path: out,
    format: "A4",
    printBackground: true,
    displayHeaderFooter: true,
    headerTemplate: "<div></div>",
    footerTemplate: foot,
    margin: { top: "16mm", right: "18mm", bottom: "16mm", left: "18mm" },
  });
  await browser.close();
  console.log(`✓ PDF A4 gerado: ${out}`);
})();
