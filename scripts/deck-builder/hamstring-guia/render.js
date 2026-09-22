// Guia clínico · Tendinopatia proximal dos isquiotibiais · A4 retrato -> PDF + PNG por página.
// Uso: node render.js [--check] [--png] [--page=NN] [--pdf]
//                     [--paciente="Nome Completo"] [--data="22 de setembro de 2026"]
// Com --paciente o nome e a data entram na capa e o PDF sai com o nome do paciente no
// arquivo. Sem, sai a versão em branco, para preencher à mão.
// O PDF nunca é versionado (docs/**/*.pdf está no .gitignore), que é onde ele precisa
// ficar: nome de paciente não entra no repositório.
const { chromium } = require("/home/user/plenya/scripts/deck-builder/continuum/node_modules/playwright");
const path = require("path"); const fs = require("fs");
const args = process.argv.slice(2); const has = f => args.includes(`--${f}`);
const pageArg = (args.find(a => a.startsWith("--page=")) || "").split("=")[1];
const arg = f => { const a = args.find(x => x.startsWith(`--${f}=`)); return a ? a.slice(f.length + 3).replace(/^["']|["']$/g, "") : ""; };
const paciente = arg("paciente");
const dataDoc = arg("data") || new Date().toLocaleDateString("pt-BR", { day: "numeric", month: "long", year: "numeric" });
const slug = s => s.normalize("NFD").replace(/[\u0300-\u036f]/g, "").toLowerCase().replace(/[^a-z0-9]+/g, "-").replace(/^-|-$/g, "");
(async () => {
  const browser = await chromium.launch();
  const page = await browser.newPage({ viewport: { width: 794, height: 1123 }, deviceScaleFactor: 2 });
  await page.goto(`file://${path.resolve(__dirname, "guia.html")}`, { waitUntil: "networkidle" });
  if (paciente) {
    await page.evaluate(([nome, quando]) => {
      document.querySelector("#paciente").textContent = nome;
      document.querySelector("#data").textContent = quando;
    }, [paciente, dataDoc]);
  }
  await page.evaluate(() => document.fonts.ready); await page.waitForTimeout(800);
  if (has("check")) {
    const res = await page.evaluate(() => [...document.querySelectorAll("section.page")].map((p, i) => {
      const top = p.getBoundingClientRect().top; let max = 0;
      p.querySelectorAll(".body *").forEach(e => { const b = e.getBoundingClientRect().bottom - top; if (b > max) max = b; });
      return { n: i + 1, mm: +(max * 25.4 / 96).toFixed(1) };
    }));
    res.forEach(r => console.log(`p${String(r.n).padStart(2,"0")} fim do conteudo ${r.mm} mm${r.mm > 281 ? "  ** TRANSBORDA" : ""}`));
  }
  if (has("png") || pageArg) {
    const out = path.resolve(__dirname, "previews"); fs.mkdirSync(out, { recursive: true });
    const secs = await page.$$("section.page");
    for (let i = 0; i < secs.length; i++) {
      const n = String(i + 1).padStart(2, "0"); if (pageArg && pageArg !== n) continue;
      await secs[i].screenshot({ path: `${out}/p-${n}.png` });
    }
    console.log(`ok PNGs em ${out}`);
  }
  if (has("pdf")) {
    const nomeArq = paciente
      ? `tendinopatia-isquiotibiais-${slug(paciente)}.pdf`
      : "tendinopatia-isquiotibiais-proximal.pdf";
    const out = path.resolve(__dirname, "../../../docs/clinica", nomeArq);
    await page.pdf({ path: out, format: "A4", printBackground: true, preferCSSPageSize: true, margin: { top: 0, right: 0, bottom: 0, left: 0 } });
    console.log(`ok PDF: ${out}`);
  }
  await browser.close();
})();
