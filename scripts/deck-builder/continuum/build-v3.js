// Build deck-v3.html from deck-v2.html: reorder slides + insert 4 new slides + renumber.
// Source of truth for v3 structure. Run: node build-v3.js
const fs = require("fs");
const path = require("path");

const v2 = fs.readFileSync(path.resolve(__dirname, "deck-v2.html"), "utf8");

// --- split head / sections / tail ---
const firstSec = v2.indexOf('<section class="slide');
let head = v2.slice(0, firstSec).replace(/<!--[\s\S]*$/, "");
const lastClose = v2.lastIndexOf("</section>") + "</section>".length;
const tail = v2.slice(lastClose);

const blocks = {};
const re = /<section class="slide[\s\S]*?<\/section>/g;
let m;
while ((m = re.exec(v2)) !== null) {
  const id = (m[0].match(/id="(s\d+)"/) || [])[1];
  if (id) blocks[id] = m[0];
}

// --- chrome helper ---
const brandTL = `  <div style="position: absolute; top: 72px; left: 96px; z-index: 3; display: flex; align-items: center; gap: 18px;">
    <img src="../../../apps/site/public/brand/symbol/gold.png" alt="Plenya" style="width: 36px; height: auto; opacity: 0.95;">
    <span style="font-family: var(--serif); font-size: 24px; font-weight: 500; letter-spacing: 0.34em; color: var(--cream); text-transform: uppercase; opacity: 0.85;">Continuum Plenya</span>
  </div>`;
const eyebrowTR = (t) => `  <div style="position: absolute; top: 78px; right: 96px; z-index: 3; font-family: var(--serif); font-size: 24px; font-weight: 500; letter-spacing: 0.38em; color: var(--gold-soft); text-transform: uppercase;">${t}</div>`;
const pageBL = `  <div style="position: absolute; bottom: 40px; right: 96px; z-index: 3;">
    <span style="font-family: var(--serif); font-size: 24px; font-weight: 400; letter-spacing: 0.32em; color: var(--cream); opacity: 0.75;">90 / 18</span>
  </div>`;

// ============ NEW 1 (s90) — O QUE MUDA · contraste de categoria ============
const epLine = (t) => `        <div style="font-family: var(--serif); font-weight: 400; font-size: 26px; line-height: 1.3; color: var(--sage); opacity: 0.65; padding: 9px 0; border-bottom: 1px solid rgba(146,184,180,0.10);">${t}</div>`;
const loLine = (t) => `        <div style="font-family: var(--serif); font-weight: 400; font-size: 26px; line-height: 1.3; color: var(--cream); opacity: 0.95; padding: 9px 0; border-bottom: 1px solid rgba(234,231,218,0.08);">${t}</div>`;
blocks["s90"] = `<section class="slide slide--deep" id="s90">
${brandTL}
${eyebrowTR("A diferença")}
${pageBL}
  <div class="slide-content" style="position: relative; z-index: 1; justify-content: center; align-items: center; padding: 130px 96px 96px;">
    <h2 style="font-family: var(--serif-display); font-weight: 500; font-size: 58px; line-height: 1.08; letter-spacing: -0.02em; color: var(--cream); text-align: center; max-width: 1500px; margin-bottom: 56px;">
      O que muda quando alguém acompanha a trajetória<span style="color: var(--gold);">.</span>
    </h2>
    <div style="display: grid; grid-template-columns: 1fr 1fr; gap: 0; width: 100%; max-width: 1640px; margin-bottom: 52px;">
      <div style="padding: 0 56px 0 24px; border-right: 1.5px solid rgba(179,134,69,0.35);">
        <div style="font-family: var(--serif); font-size: 24px; font-weight: 500; letter-spacing: 0.30em; color: var(--sage); text-transform: uppercase; opacity: 0.7; margin-bottom: 18px;">O modelo de hoje</div>
        <h3 style="font-family: var(--serif-display); font-weight: 500; font-size: 40px; line-height: 1.05; letter-spacing: -0.02em; color: var(--sage); opacity: 0.85; margin-bottom: 26px;">Cuidado episódico<span style="color: var(--sage);">.</span></h3>
${epLine("Exames que tranquilizam")}
${epLine("Especialistas que não conversam")}
${epLine("Sintomas tratados um a um")}
${epLine("Decisão só quando algo piora")}
      </div>
      <div style="padding: 0 24px 0 56px;">
        <div style="font-family: var(--serif); font-size: 24px; font-weight: 500; letter-spacing: 0.30em; color: var(--gold-soft); text-transform: uppercase; margin-bottom: 18px;">O Continuum</div>
        <h3 style="font-family: var(--serif-display); font-weight: 500; font-size: 40px; line-height: 1.05; letter-spacing: -0.02em; color: var(--cream); margin-bottom: 26px;">Cuidado longitudinal<span style="color: var(--gold);">.</span></h3>
${loLine("Leitura da trajetória inteira")}
${loLine("Uma equipe, um painel")}
${loLine("Ajustes contínuos no tempo")}
${loLine("Decisão antes da perda")}
      </div>
    </div>
    <p style="font-family: var(--serif); font-style: italic; font-weight: 400; font-size: 34px; line-height: 1.3; color: var(--gold-soft); text-align: center; max-width: 1300px; letter-spacing: -0.005em;">
      Continuum não muda um exame. Muda o que acontece entre eles.
    </p>
  </div>
</section>`;

// ============ NEW 2 (s91) — O MÉDICO-GESTOR · hub ============
// hub SVG: centro = médico-gestor, órbita = disciplinas + externos
const cx = 360, cy = 300, R = 215;
const nodes = [
  { a: -90, t: "Seus especialistas" },
  { a: -18, t: "Nutrição" },
  { a: 54,  t: "Psicologia" },
  { a: 126, t: "Educação Física" },
  { a: 198, t: "Exames & dados" },
];
const pt = (a) => [cx + R * Math.cos(a * Math.PI / 180), cy + R * Math.sin(a * Math.PI / 180)];
let lines = "", dots = "", labels = "";
for (const n of nodes) {
  const [x, y] = pt(n.a);
  lines += `    <line x1="${cx}" y1="${cy}" x2="${x.toFixed(0)}" y2="${y.toFixed(0)}" stroke="#B38645" stroke-width="1" opacity="0.45"/>\n`;
  dots += `    <circle cx="${x.toFixed(0)}" cy="${y.toFixed(0)}" r="5" fill="#92B8B4"/>\n`;
  const anchor = x < cx - 20 ? "end" : x > cx + 20 ? "start" : "middle";
  const dx = anchor === "end" ? -14 : anchor === "start" ? 14 : 0;
  const dy = Math.abs(y - cy) < 20 ? 6 : (y < cy ? -16 : 26);
  labels += `    <text x="${(x + dx).toFixed(0)}" y="${(y + dy).toFixed(0)}" text-anchor="${anchor}" font-family="Cormorant Garamond, serif" font-size="25" fill="#92B8B4">${n.t}</text>\n`;
}
const hub = `  <svg viewBox="-70 0 850 600" style="width: 100%; height: auto; max-height: 720px;" aria-hidden="true">
${lines}    <circle cx="${cx}" cy="${cy}" r="86" fill="rgba(179,134,69,0.06)" stroke="#B38645" stroke-width="1.5"/>
    <text x="${cx}" y="${cy - 6}" text-anchor="middle" font-family="Fraunces, Cormorant Garamond, serif" font-size="30" font-weight="500" fill="#D4A86B">Médico-</text>
    <text x="${cx}" y="${cy + 28}" text-anchor="middle" font-family="Fraunces, Cormorant Garamond, serif" font-size="30" font-weight="500" fill="#D4A86B">gestor</text>
${dots}${labels}  </svg>`;
blocks["s91"] = `<section class="slide slide--deep" id="s91">
${brandTL}
${eyebrowTR("O diferencial")}
${pageBL}
  <div class="slide-content" style="position: relative; z-index: 1; justify-content: center; align-items: center; padding: 120px 96px 90px;">
    <div style="display: grid; grid-template-columns: 1fr 1fr; gap: 64px; width: 100%; max-width: 1680px; align-items: center;">
      <div>
        <h2 style="font-family: var(--serif-display); font-weight: 500; font-size: 88px; line-height: 1.0; letter-spacing: -0.03em; color: var(--cream); margin-bottom: 28px;">
          O médico-gestor<span style="color: var(--gold);">.</span>
        </h2>
        <p style="font-family: var(--serif); font-style: italic; font-weight: 400; font-size: 36px; line-height: 1.25; color: var(--gold); letter-spacing: -0.005em; margin-bottom: 44px; max-width: 720px;">
          Alguém precisa acompanhar a trajetória inteira.
        </p>
        <p style="font-family: var(--serif); font-weight: 400; font-size: 28px; line-height: 1.5; color: var(--cream); opacity: 0.92; margin-bottom: 28px; max-width: 720px;">
          Não assume o lugar dos seus médicos. Costura o todo: integra os especialistas, coordena as decisões, traduz exames em contexto e acompanha ao longo do tempo.
        </p>
        <p style="font-family: var(--serif); font-style: italic; font-weight: 400; font-size: 28px; line-height: 1.4; color: var(--gold-soft); max-width: 720px;">
          É o que faz o Continuum existir.
        </p>
      </div>
      <div style="display: flex; justify-content: center; align-items: center;">
${hub}
      </div>
    </div>
  </div>
</section>`;

// ============ NEW 3 (s92) — O ESCORE ACOMPANHA · trajetória (mock) ============
const ymap = (s) => (400 - (s - 50) / 40 * 340).toFixed(1);
const xs = [160, 460, 760, 1060];
const sc = [58, 63, 67, 71];
const xl = ["Tempo 0", "3 meses", "6 meses", "9 meses"];
const ann = ["", "Sono ajusta", "Inflamação cede", "Humor e energia"];
let poly = "", pts = "", scl = "", xlab = "", anl = "";
xs.forEach((x, i) => {
  const y = ymap(sc[i]);
  poly += `${x},${y} `;
  pts += `    <circle cx="${x}" cy="${y}" r="7" fill="#B38645"/>\n`;
  scl += `    <text x="${x}" y="${(+y - 22).toFixed(0)}" text-anchor="middle" font-family="Fraunces, serif" font-size="34" font-weight="500" fill="#EAE7DA">${sc[i]}</text>\n`;
  xlab += `    <text x="${x}" y="445" text-anchor="middle" font-family="Cormorant Garamond, serif" font-size="24" fill="#92B8B4" opacity="0.85" letter-spacing="2">${xl[i]}</text>\n`;
  if (ann[i]) anl += `    <text x="${x}" y="${(+y + 40).toFixed(0)}" text-anchor="middle" font-family="Cormorant Garamond, serif" font-style="italic" font-size="23" fill="#92B8B4">${ann[i]}</text>\n`;
});
const chart = `  <svg viewBox="0 0 1200 480" style="width: 100%; height: auto; max-width: 1500px;" aria-hidden="true">
    <line x1="120" y1="400" x2="1100" y2="400" stroke="rgba(179,134,69,0.25)" stroke-width="1"/>
    <polyline points="${poly.trim()}" fill="none" stroke="#B38645" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"/>
${pts}${scl}${anl}${xlab}  </svg>`;
blocks["s92"] = `<section class="slide slide--deep" id="s92">
${brandTL}
${eyebrowTR("Longitudinal")}
${pageBL}
  <div class="slide-content" style="position: relative; z-index: 1; justify-content: center; align-items: center; padding: 120px 96px 90px;">
    <h2 style="font-family: var(--serif-display); font-weight: 500; font-size: 76px; line-height: 1.05; letter-spacing: -0.025em; color: var(--cream); text-align: center; margin-bottom: 14px;">
      O Escore não prevê. Acompanha<span style="color: var(--gold);">.</span>
    </h2>
    <p style="font-family: var(--serif); font-style: italic; font-weight: 400; font-size: 30px; line-height: 1.3; color: var(--gold-soft); text-align: center; max-width: 1100px; margin-bottom: 48px; letter-spacing: -0.005em;">
      A cada ciclo, uma leitura nova. A curva mostra a direção.
    </p>
${chart}
    <p style="font-family: var(--serif); font-style: italic; font-weight: 400; font-size: 30px; line-height: 1.3; color: var(--gold-soft); text-align: center; margin-top: 36px;">
      Direção, não punição.
    </p>
    <p style="font-family: var(--serif); font-weight: 400; font-size: 20px; line-height: 1.3; color: var(--sage); opacity: 0.7; text-align: center; margin-top: 14px;">
      Exemplo ilustrativo. A curva real é individual.
    </p>
  </div>
</section>`;

// ============ NEW 4 (s93) — O QUE 20 ANOS ENSINARAM ============
const lessons = [
  ["I", "Doença grave raramente começa no diagnóstico."],
  ["II", "Quem acompanha de perto decide mais cedo."],
  ["III", "Corpo lido em partes perde o todo."],
  ["IV", "Acompanhamento contínuo muda trajetórias."],
];
let lcols = "";
lessons.forEach(([r, t], i) => {
  const border = i < 3 ? "border-right: 1.5px solid rgba(179,134,69,0.32);" : "";
  lcols += `      <div style="padding: 0 36px; ${border} text-align: left;">
        <div style="font-family: var(--serif-display); font-style: italic; font-weight: 400; font-size: 44px; line-height: 1; color: var(--gold-soft); margin-bottom: 28px;">${r}</div>
        <h3 style="font-family: var(--serif-display); font-weight: 500; font-size: 34px; line-height: 1.15; letter-spacing: -0.015em; color: var(--cream);">${t}</h3>
      </div>\n`;
});
blocks["s93"] = `<section class="slide slide--deep" id="s93">
${brandTL}
${eyebrowTR("Vinte anos")}
${pageBL}
  <div class="slide-content" style="position: relative; z-index: 1; justify-content: center; align-items: center; padding: 140px 96px 100px;">
    <h2 style="font-family: var(--serif-display); font-weight: 500; font-size: 64px; line-height: 1.08; letter-spacing: -0.02em; color: var(--cream); text-align: center; max-width: 1500px; margin-bottom: 80px;">
      O que vinte anos de prática ensinaram<span style="color: var(--gold);">.</span>
    </h2>
    <div style="display: grid; grid-template-columns: repeat(4, 1fr); gap: 0; width: 100%; max-width: 1700px; margin-bottom: 80px;">
${lcols}    </div>
    <p style="font-family: var(--serif); font-style: italic; font-weight: 400; font-size: 34px; line-height: 1.3; color: var(--gold-soft); text-align: center; letter-spacing: -0.005em;">
      O Continuum nasceu daqui.
    </p>
  </div>
</section>`;

// --- v3 order (by content) ---
const order = ["s01","s02","s03","s04","s05","s06","s90","s91","s09","s10","s11","s92","s12","s08","s93","s07","s14","s13","s15","s16","s17","s18"];
const titles = {s01:"HOOK",s02:"JANELA SILENCIOSA",s03:"ORIGEM",s04:"CONTINUUM REVEAL",s05:"TRÊS VERBOS",s06:"TRÊS CRENÇAS",s90:"O QUE MUDA (novo)",s91:"MÉDICO-GESTOR (novo)",s09:"MÉTODO AGIR",s10:"APROFUNDAMENTO G",s11:"ESCORE",s92:"ESCORE ACOMPANHA (novo)",s12:"JORNADA",s08:"EQUIPE",s93:"20 ANOS (novo)",s07:"PARA QUEM É",s14:"PARA QUEM NÃO É",s13:"BOX",s15:"ESCOPO",s16:"VALOR + COMPROMISSO",s17:"INVESTIMENTO",s18:"FECHAMENTO"};

const TT = order.length; // 22
let body = "";
order.forEach((id, idx) => {
  const pos = String(idx + 1).padStart(2, "0");
  let b = blocks[id];
  if (!b) throw new Error("missing block " + id);
  b = b.replace(/id="s\d+"/, `id="s${pos}"`);
  b = b.replace(/>(\d{2}) \/ \d{2}</g, `>${pos} / ${TT}<`);
  body += `<!-- ============================================================\n     SLIDE ${pos} — ${titles[id]}\n     ============================================================ -->\n${b}\n\n`;
});

fs.writeFileSync(path.resolve(__dirname, "deck-v3.html"), head + body + tail.replace(/^\s*/, "\n"));
console.log(`✓ deck-v3.html — ${TT} slides`);
