/**
 * sync-score-versions — gera os configs ESTÁTICOS do site a partir da fonte de verdade do EMR
 * (score_versions). Substitui o antigo sync-score-light.ts. Para cada version, busca
 * GET /api/v1/score-light/config/<slug> e grava apps/site/content/data/<file>.
 *
 * Uso (a partir de apps/site):  pnpm sync:score-versions
 * API: INTERNAL_API_URL | NEXT_PUBLIC_API_URL | http://localhost:3001
 */
import { writeFileSync } from 'node:fs';
import { join } from 'node:path';

const API = process.env.INTERNAL_API_URL || process.env.NEXT_PUBLIC_API_URL || 'http://localhost:3001';
const OUT_DIR = join(process.cwd(), 'content', 'data');

const VERSIONS: { slug: string; file: string }[] = [
  { slug: 'triagem', file: 'score-triagem-config.json' },
  { slug: 'light', file: 'score-light-config.json' },
];

async function main() {
  for (const v of VERSIONS) {
    const url = `${API}/api/v1/score-light/config/${v.slug}`;
    const res = await fetch(url);
    if (!res.ok) throw new Error(`${v.slug}: HTTP ${res.status} em ${url}`);
    const cfg = await res.json();
    const path = join(OUT_DIR, v.file);
    writeFileSync(path, JSON.stringify(cfg, null, 2) + '\n', 'utf8');
    console.log(`✓ ${v.file} — ${cfg.itemCount} itens (${cfg.version})`);
  }
  console.log('Configs do site sincronizados da fonte EMR (score_versions).');
}

main().catch((e) => {
  console.error('✗ sync-score-versions falhou:', e.message);
  process.exit(1);
});
