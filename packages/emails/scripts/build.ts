/**
 * Renderiza cada template React Email para HTML estático em dist/.
 * Os arquivos .html mantêm placeholders {{VAR}} que o backend Go substitui
 * em runtime via strings.ReplaceAll após carregar via go:embed.
 *
 * Uso: pnpm --filter @plenya/emails build
 */
import { render } from '@react-email/render';
import { mkdirSync, writeFileSync } from 'node:fs';
import { dirname, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';

import MagicLink from '../src/templates/MagicLink.js';
import BoasVindas from '../src/templates/BoasVindas.js';
import FollowUp30Dias from '../src/templates/FollowUp30Dias.js';

const __dirname = dirname(fileURLToPath(import.meta.url));
const distDir = resolve(__dirname, '..', 'dist');

mkdirSync(distDir, { recursive: true });

const templates = {
  magic_link: MagicLink,
  boas_vindas: BoasVindas,
  follow_up_30_dias: FollowUp30Dias,
} as const;

async function build() {
  for (const [name, Template] of Object.entries(templates)) {
    const html = await render(Template());
    const path = resolve(distDir, `${name}.html`);
    writeFileSync(path, html, 'utf8');
    console.log(`✓ ${name}.html (${html.length} bytes)`);
  }
}

build().catch((err) => {
  console.error('build failed:', err);
  process.exit(1);
});
