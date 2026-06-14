// Rasteriza o lockup direto (goto no SVG) com deviceScaleFactor alto.
// viewBox 504.125 x 121.916 -> DSF 10 => ~5041 x 1219 px (10 px / unidade).
const { chromium } = require("playwright");
const path = require("path");
const SVG = "file://" + path.resolve(__dirname, "assets/lockup-tagline-petrol.svg");
const DSF = 10;
(async () => {
  const browser = await chromium.launch();
  const ctx = await browser.newContext({
    viewport: { width: 505, height: 122 }, deviceScaleFactor: DSF,
  });
  const page = await ctx.newPage();
  await page.goto(SVG, { waitUntil: "networkidle" });
  await page.waitForTimeout(300);
  await page.screenshot({ path: "/tmp/lockup-measure.png" });
  await browser.close();
  console.log("ok DSF=" + DSF);
})();
