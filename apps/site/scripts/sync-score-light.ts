/**
 * sync-score-light.ts — busca a configuração atual do Escore Light no EMR
 * e grava em content/data/score-light-config.json (commitado no repo).
 *
 * Uso:
 *   pnpm sync:score-light
 *   pnpm sync:score-light --api=https://emr.plenyasaude.com.br
 *
 * Quando o admin marca/desmarca/edita items Light no EMR, basta rodar
 * este script + commit para o site refletir.
 */

import { writeFileSync, mkdirSync } from 'node:fs';
import { dirname, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';

const __dirname = dirname(fileURLToPath(import.meta.url));

const API_URL =
  process.argv.find((a) => a.startsWith('--api='))?.split('=')[1] ??
  process.env.PLENYA_API_URL ??
  'http://localhost:3001';

const OUTPUT_PATH = resolve(__dirname, '..', 'content', 'data', 'score-light-config.json');

async function main() {
  const url = `${API_URL.replace(/\/$/, '')}/api/v1/score-light/config`;
  console.log(`→ Buscando ${url}`);

  const res = await fetch(url, { headers: { 'cache-control': 'no-cache' } });
  if (!res.ok) {
    console.error(`✗ Erro ${res.status}: ${res.statusText}`);
    process.exit(1);
  }
  const data = await res.json();

  if (!data || typeof data !== 'object' || !('groups' in data)) {
    console.error('✗ Resposta inesperada — esperava objeto com .groups');
    process.exit(1);
  }

  mkdirSync(dirname(OUTPUT_PATH), { recursive: true });
  writeFileSync(OUTPUT_PATH, JSON.stringify(data, null, 2) + '\n', 'utf-8');

  console.log(
    `✓ ${data.itemCount ?? 0} items Light em ${data.groups.length} grupos → ${OUTPUT_PATH}`
  );
  console.log(`  Versão: ${data.version}`);
}

main().catch((err) => {
  console.error('✗ Falha no sync:', err);
  process.exit(1);
});
