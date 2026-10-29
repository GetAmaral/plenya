// Ofício (AREM -> DRAS) e card do paciente · biópsia renal · A4 retrato.
// Uso: node render.js [--png] [--check]
//   sempre gera, em docs/clinica/biopsia-renal/ (docs/**/*.pdf não é versionado):
//     card-paciente-biopsia-renal.pdf   card sozinho (1 página fixa)
//     oficio-biopsia-renal.pdf          ofício + card anexo
//   --png    PNG por página em previews/ (aprovação), a partir dos PDFs (pdftoppm)
//   --check  avisa se o card transborda a folha
// O ofício é texto corrido: quem pagina é o Chromium (@page no oficio.html), com número de página
// no rodapé. O card é página fixa (section.page).
const { chromium } = require("/home/user/plenya/scripts/deck-builder/continuum/node_modules/playwright");
const path = require("path"); const fs = require("fs"); const { execFileSync } = require("child_process");
const args = process.argv.slice(2); const has = f => args.includes(`--${f}`);
const PDF_DIR = path.resolve(__dirname, "../../../docs/clinica/biopsia-renal");
const PREV = path.resolve(__dirname, "previews");
const tmp = n => path.join(PREV, `.${n}.pdf`);
(async () => {
  fs.mkdirSync(PDF_DIR, { recursive: true }); fs.mkdirSync(PREV, { recursive: true });
  const browser = await chromium.launch();
  const open = async html => {
    const page = await browser.newPage({ viewport: { width: 794, height: 1123 } });
    await page.goto(`file://${path.resolve(__dirname, html)}`, { waitUntil: "networkidle" });
    await page.evaluate(() => document.fonts.ready); await page.waitForTimeout(600);
    return page;
  };

  const card = await open("card.html");
  if (has("check")) {
    const mm = await card.evaluate(() => {
      const p = document.querySelector("section.page"), top = p.getBoundingClientRect().top; let max = 0;
      p.querySelectorAll("*:not(.foot):not(.foot *)").forEach(e => { const b = e.getBoundingClientRect().bottom - top; if (b > max) max = b; });
      return +(max * 25.4 / 96).toFixed(1);
    });
    console.log(`card: fim do conteudo ${mm} mm${mm > 280 ? "  ** TRANSBORDA" : ""}`);
  }
  await card.pdf({ path: `${PDF_DIR}/card-paciente-biopsia-renal.pdf`, format: "A4", printBackground: true, preferCSSPageSize: true });

  const oficio = await open("oficio.html");
  await oficio.pdf({
    path: tmp("oficio"), format: "A4", printBackground: true, preferCSSPageSize: true,
    displayHeaderFooter: true, headerTemplate: "<span></span>",
    footerTemplate: `<div style="width:100%;font-family:Georgia,serif;font-size:8px;color:#55646A;padding:0 22mm;text-align:right">
      <span class="pageNumber"></span>/<span class="totalPages"></span></div>`,
  });
  await browser.close();

  execFileSync("pdfunite", [tmp("oficio"), `${PDF_DIR}/card-paciente-biopsia-renal.pdf`, `${PDF_DIR}/oficio-biopsia-renal.pdf`]);
  if (has("png")) {
    execFileSync("pdftoppm", ["-png", "-r", "150", tmp("oficio"), path.join(PREV, "oficio")]);
    execFileSync("pdftoppm", ["-png", "-r", "150", "-singlefile", `${PDF_DIR}/card-paciente-biopsia-renal.pdf`, path.join(PREV, "card")]);
    console.log(`ok PNGs em ${PREV}`);
  }
  fs.rmSync(tmp("oficio"));
  console.log(`ok PDFs em ${PDF_DIR}`);
})();
