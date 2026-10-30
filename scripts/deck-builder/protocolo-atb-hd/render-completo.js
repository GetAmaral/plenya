// PDF único do protocolo: as 4 folhas do card + a fundamentação científica (TEORICO.md) no mesmo visual.
// Uso: node render.js && node render-completo.js [--png]
//   gera docs/clinica/protocolo-atb-dialise/protocolo-atb-hd-completo.pdf (docs/**/*.pdf não é versionado)
//   --png    PNG das páginas do texto em previews/texto-N.png (conferência)
const { chromium } = require("/home/user/plenya/scripts/deck-builder/continuum/node_modules/playwright");
const { marked } = require("/home/user/plenya/node_modules/marked");
const path = require("path"); const fs = require("fs"); const { execFileSync } = require("child_process");
const DOCS = path.resolve(__dirname, "../../../docs/clinica/protocolo-atb-dialise");
const CARD = `${DOCS}/card-protocolo-atb-hd.pdf`;
const OUT = `${DOCS}/protocolo-atb-hd-completo.pdf`;
const PREV = path.resolve(__dirname, "previews");

const md = fs.readFileSync(`${DOCS}/TEORICO.md`, "utf8");
const [titleLine, ...rest] = md.split("\n");
const body = marked.parse(rest.join("\n"))
  .replace(/\[(\d+(?:, \d+)*)\]/g, '<span class="rf">[$1]</span>');   // números de referência discretos
const refsAt = body.indexOf('<h2 id="referências">') >= 0 ? body.indexOf('<h2 id="referências">') : body.lastIndexOf("<h2");

const html = `<!DOCTYPE html><html lang="pt-BR"><head><meta charset="UTF-8">
<style>
@import url('https://fonts.googleapis.com/css2?family=Atkinson+Hyperlegible:ital,wght@0,400;0,700;1,400;1,700&display=swap');
:root{--pet:#0E4F5C;--pet2:#2C7A86;--ink:#17323A;--mut:#5C7379;--rule:#DCE7E9;--bg:#F5F9F9}
@page{size:A4;margin:16mm 18mm 18mm}
*{box-sizing:border-box}
body{margin:0;font-family:'Atkinson Hyperlegible',Arial,sans-serif;color:var(--ink);font-size:10.4pt;line-height:1.5}
.hd{border-bottom:.35mm solid var(--rule);padding-bottom:3mm;margin-bottom:6mm}
.hd h1{font-size:19pt;color:var(--pet);margin:0;line-height:1.15}
.hd .k{font-size:9.5pt;color:var(--mut);margin-top:1.6mm}
h2{font-size:13pt;color:var(--pet);margin:7mm 0 2mm;break-after:avoid}
h3{font-size:11pt;color:var(--pet2);margin:5mm 0 1.5mm;break-after:avoid}
p{margin:0 0 2.6mm;text-align:left;orphans:3;widows:3}
hr{border:0;border-top:.3mm solid var(--rule);margin:5mm 0}
.rf{font-size:.78em;color:var(--pet2);white-space:nowrap}
code{font-family:inherit}
.refs ol{padding-left:7mm;margin:0}
.refs li{font-size:8.8pt;line-height:1.38;margin:0 0 1.4mm;break-inside:avoid}
.refs p{margin:0}
</style></head><body>
<div class="hd"><h1>${titleLine.replace(/^#\s*/, "")}</h1>
<div class="k">DaVita Londrina | Bandeirantes | Lago Parque | Rolândia · Acompanha o card do protocolo (folhas 1 a 4)</div></div>
${body.slice(0, refsAt)}<div class="refs">${body.slice(refsAt)}</div>
</body></html>`;

(async () => {
  if (!fs.existsSync(CARD)) throw new Error("rode antes: node render.js");
  const tmp = path.join(PREV, ".texto.html"); fs.writeFileSync(tmp, html);
  const browser = await chromium.launch();
  const page = await browser.newPage();
  await page.goto(`file://${tmp}`, { waitUntil: "networkidle" });
  await page.evaluate(() => document.fonts.ready);
  const textPdf = path.join(PREV, ".texto.pdf");
  await page.pdf({
    path: textPdf, format: "A4", printBackground: true, displayHeaderFooter: true,
    headerTemplate: "<span></span>",
    footerTemplate: `<div style="width:100%;font-family:Arial,sans-serif;font-size:7.5pt;color:#5C7379;padding:0 18mm;display:flex;justify-content:space-between">
      <span>Protocolo de antibióticos na hemodiálise · Fundamentação científica</span><span><span class="pageNumber"></span>/<span class="totalPages"></span></span></div>`,
    margin: { top: "16mm", bottom: "18mm", left: "18mm", right: "18mm" },
  });
  await browser.close();
  fs.rmSync(tmp);
  execFileSync("pdfunite", [CARD, textPdf, OUT]);
  if (process.argv.includes("--png")) execFileSync("pdftoppm", ["-png", "-r", "110", textPdf, path.join(PREV, "texto")]);
  fs.rmSync(textPdf);
  console.log(`ok ${OUT}`);
})();
