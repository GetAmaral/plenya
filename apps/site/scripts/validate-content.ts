/**
 * Valida frontmatter MDX do blog Plenya (PT + EN + ES).
 * Falha com exit 1 se algum arquivo for inválido — bloqueia o build.
 *
 * Rodar manual: pnpm validate-content
 * Roda automático via prebuild antes de `next build`.
 *
 * Schema duplicado a propósito de lib/blog.ts (manter sincronizado).
 */
import { promises as fs } from 'fs';
import path from 'path';
import matter from 'gray-matter';
import { z } from 'zod';

const ROOT = process.cwd();

const dateString = z.union([z.string(), z.date()]).transform((v) =>
  v instanceof Date ? v.toISOString().slice(0, 10) : v
);

const PILLARS = [
  'alimentacao-atividade-fisica',
  'gestao-metabolica',
  'integracao-corpo-mente',
  'ritmo-circadiano',
  'longevidade',
] as const;

const frontmatterSchema = z.object({
  title: z.string(),
  slug: z.string(),
  excerpt: z.string(),
  date: dateString,
  updated: dateString.optional(),
  author: z.string(),
  reviewedBy: z.string().optional(),
  pillar: z.enum(PILLARS),
  tags: z.array(z.string()).default([]),
  cover: z.string().optional(),
  featured: z.boolean().default(false),
  references: z
    .array(z.object({ label: z.string(), url: z.string().url().optional() }))
    .default([]),
  readingMinutes: z.number().optional(),
  cta: z.enum(['default', 'recognition']).default('default'),
});

type Failure = { file: string; kind: 'yaml' | 'zod' | 'read'; detail: string };

async function validateDir(dir: string, label: string): Promise<{ ok: number; failures: Failure[] }> {
  const failures: Failure[] = [];
  let ok = 0;
  let files: string[];
  try {
    files = (await fs.readdir(dir)).filter((f) => f.endsWith('.mdx'));
  } catch {
    console.log(`  ${label}: diretório não existe (${dir}) — skip`);
    return { ok: 0, failures };
  }

  for (const file of files) {
    const full = path.join(dir, file);
    let raw: string;
    try {
      raw = await fs.readFile(full, 'utf-8');
    } catch (err) {
      failures.push({ file: full, kind: 'read', detail: (err as Error).message });
      continue;
    }
    try {
      const { data } = matter(raw);
      const parsed = frontmatterSchema.safeParse(data);
      if (!parsed.success) {
        const detail = parsed.error.issues
          .map((i) => `${i.path.join('.')}: ${i.message}`)
          .join(' | ');
        failures.push({ file: full, kind: 'zod', detail });
      } else {
        ok++;
      }
    } catch (err) {
      failures.push({
        file: full,
        kind: 'yaml',
        detail: (err as Error).message.split('\n')[0],
      });
    }
  }

  return { ok, failures };
}

async function main() {
  const targets = [
    { dir: path.join(ROOT, 'content/blog/pt'), label: 'blog/pt' },
    { dir: path.join(ROOT, 'content/blog/en'), label: 'blog/en' },
    { dir: path.join(ROOT, 'content/blog/es'), label: 'blog/es' },
  ];

  let totalOk = 0;
  const allFailures: Failure[] = [];

  for (const t of targets) {
    const { ok, failures } = await validateDir(t.dir, t.label);
    totalOk += ok;
    allFailures.push(...failures);
    const status = failures.length === 0 ? 'OK' : 'FAIL';
    console.log(`  ${t.label}: ${ok} ok / ${failures.length} bad [${status}]`);
  }

  if (allFailures.length > 0) {
    console.log('\n=== Detalhe das falhas ===');
    for (const f of allFailures) {
      console.log(`\n  [${f.kind.toUpperCase()}] ${path.relative(ROOT, f.file)}`);
      console.log(`    ${f.detail}`);
    }
    console.log(
      `\n${allFailures.length} arquivo(s) inválido(s). Build bloqueado. Corrija e tente de novo.`
    );
    process.exit(1);
  }

  console.log(`\nTodos os ${totalOk} MDX válidos.`);
}

main().catch((err) => {
  console.error('validate-content crashed:', err);
  process.exit(1);
});
