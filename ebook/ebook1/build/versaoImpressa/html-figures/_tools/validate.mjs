#!/usr/bin/env node
// validate.mjs <figure_dir>
//   renders <figure_dir>/figure.svg with Playwright →
//     <figure_dir>/preview.png      (just my render)
//     <figure_dir>/overlay.png      (mine + original at 35%)
//     <figure_dir>/diff.png         (mix-blend-mode difference)
//
// Also writes the final vector PDF to ../../figuras-bw/CapNN_FigNN.pdf

import { chromium } from 'playwright';
import { pathToFileURL } from 'node:url';
import path from 'node:path';
import fs from 'node:fs';

const figArg = process.argv[2];
if (!figArg) { console.error('usage: node validate.mjs <figure_dir>'); process.exit(1); }
const figDir = path.resolve(figArg);
const svgPath = path.join(figDir, 'figure.svg');
const refPath = path.join(figDir, '_reference.png');
if (!fs.existsSync(svgPath)) { console.error('no figure.svg'); process.exit(1); }

const svg = fs.readFileSync(svgPath, 'utf-8');
const m = svg.match(/viewBox="0 0 (\d+) (\d+)"/);
const W = +m[1], H = +m[2];

// Encode reference as data URL so the file:// protocol restrictions don't bite
const refDataUrl = `data:image/png;base64,${fs.readFileSync(refPath).toString('base64')}`;
const svgDataUrl = `data:image/svg+xml;base64,${Buffer.from(svg).toString('base64')}`;

// Three HTML pages: clean, overlay, diff
// Inline the SVG (not <img>) so @font-face actually loads.
const svgInline = svg.replace(/<\?xml[^>]*\?>\s*/, '');

// Build @font-face block in HTML <style> so document.fonts can pick it up.
function fontFaceCSS() {
  const home = process.env.HOME;
  const cands = [
    { path: `${home}/.local/share/fonts/playfair/PlayfairDisplay-VF.ttf`, family: 'Playfair Display', weight: 900 },
    { path: `${home}/.local/share/fonts/cormorant/CormorantGaramond-Bold.otf`, family: 'Cormorant Garamond', weight: 700 },
  ];
  const out = [];
  for (const c of cands) {
    if (!fs.existsSync(c.path)) continue;
    const b = fs.readFileSync(c.path).toString('base64');
    const fmt = c.path.endsWith('.otf') ? 'opentype' : 'truetype';
    out.push(`@font-face { font-family: '${c.family}'; font-weight: ${c.weight}; src: url(data:font/${fmt};base64,${b}) format('${fmt}'); }`);
  }
  return out.join('\n');
}
const fontCSS = fontFaceCSS();

function makeHtml(mode) {
  const refOpacity = mode === 'clean' ? 0 : (mode === 'overlay' ? 0.45 : 1);
  const blend = mode === 'diff' ? 'mix-blend-mode: difference;' : '';
  const frameBg = mode === 'clean' ? '#fff' : 'transparent';
  return `<!doctype html><html><head><meta charset="utf-8"><style>
    ${fontCSS}
    body { margin: 0; padding: 0; background: #fff; }
    .frame { position: relative; width: ${W}px; height: ${H}px; background: ${frameBg}; }
    .ref { position: absolute; inset: 0; width: ${W}px; height: ${H}px;
           opacity: ${refOpacity}; ${blend} z-index: 1; }
    .me  { position: absolute; inset: 0; width: ${W}px; height: ${H}px;
           z-index: 2; }
    .me svg { width: ${W}px; height: ${H}px; display: block; }
  </style></head><body>
  <div class="frame">
    <div class="me">${svgInline}</div>
    <img class="ref" src="${refDataUrl}">
  </div></body></html>`;
}

const browser = await chromium.launch();
const ctx = await browser.newContext({ viewport: { width: W, height: H }, deviceScaleFactor: 1 });
const page = await ctx.newPage();

for (const mode of ['clean', 'overlay', 'diff']) {
  await page.setContent(makeHtml(mode), { waitUntil: 'networkidle' });
  await page.evaluate(() => document.fonts.ready);
  const out = mode === 'clean' ? 'preview.png' : `${mode}.png`;
  await page.locator('.frame').screenshot({ path: path.join(figDir, out) });
  console.log(`→ ${out}`);
}

// Vector PDF (clean mode, inline SVG so @font-face applies)
const pdfHtml = `<!doctype html><html><head><style>${fontCSS}</style></head><body style="margin:0;padding:0;background:#fff;">
  <div style="width:${W}px;height:${H}px;">${svgInline}</div>
</body></html>`;
await page.setContent(pdfHtml, { waitUntil: 'networkidle' });
await page.evaluate(() => document.fonts.ready);

const figName = path.basename(figDir); // e.g. cap12-fig03
const m2 = figName.match(/^cap(\d+)-fig(\d+)$/);
if (m2) {
  const outName = `Cap${m2[1].padStart(2, '0')}_Fig${m2[2].padStart(2, '0')}`;
  const pdfPath = path.join(figDir, '..', '..', 'figuras-bw', `${outName}.pdf`);
  await page.pdf({ path: pdfPath, width: `${W}px`, height: `${H}px`,
                   margin: { top: 0, right: 0, bottom: 0, left: 0 },
                   printBackground: true });
  console.log(`→ ${pdfPath}`);
}

await browser.close();
