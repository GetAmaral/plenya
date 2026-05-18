/**
 * Valida frontmatter MDX de blog + lectures.
 * Falha com exit 1 se algum arquivo for inválido — bloqueia o build.
 *
 * Rodar manual: pnpm validate-content
 * Roda automático via prebuild antes de `next build`.
 *
 * O bug que motivou esse script: gray-matter falha silenciosamente em YAML
 * inválido (excerpt com aspas internas, `:` não-delimitado). readPostFile
 * captura e retorna null. O artigo some do sistema sem warning. Em maio/2026
 * dois MDXs passaram dias 404 por causa disso.
 */
import { promises as fs } from 'fs';
import path from 'path';
import matter from 'gray-matter';
import { z } from 'zod';

const ROOT = process.cwd();

// ===== Schemas duplicados a propósito do lib/blog.ts e lib/lectures.ts =====
// Manter sincronizado se mudar os originais. Duplicação reduz import-graph
// do script (evita compilar Next).

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

const blogSchema = z.object({
  title: z.string(),
  slug: z.string(),
  excerpt: z.string(),
  date: dateString,
  updated: dateString.optional(),
  author: z.string(),
  pillar: z.enum(PILLARS),
  cover: z.string().optional(),
  tags: z.array(z.string()).default([]),
  references: z
    .array(z.object({ label: z.string(), url: z.string().url().optional() }))
    .default([]),
  hasMirror: z.boolean().default(true),
});

const audienceEnum = z.enum(['corporativo', 'medicos', 'residentes', 'aberto', 'congressos']);

const lectureSchema = z.object({
  slug: z.string(),
  title: z.string(),
  subtitle: z.string().optional(),
  excerpt: z.string(),
  audience: z.array(audienceEnum).min(1),
  duration: z.string(),
  format: z.string(),
  order: z.number().default(99),
  anchor: z.boolean().default(false),
  slugEn: z.string().optional(),
  titleEn: z.string().optional(),
  subtitleEn: z.string().optional(),
  excerptEn: z.string().optional(),
  bodyEn: z.string().optional(),
  durationEn: z.string().optional(),
  formatEn: z.string().optional(),
});

type Failure = { file: string; kind: 'yaml' | 'zod' | 'read'; detail: string };

async function validateDir(
  dir: string,
  schema: z.ZodTypeAny,
  label: string
): Promise<{ ok: number; failures: Failure[] }> {
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
      const parsed = schema.safeParse(data);
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
    { dir: path.join(ROOT, 'content/blog/pt'), schema: blogSchema, label: 'blog/pt' },
    { dir: path.join(ROOT, 'content/blog/en'), schema: blogSchema, label: 'blog/en' },
    { dir: path.join(ROOT, 'content/lectures'), schema: lectureSchema, label: 'lectures' },
  ];

  let totalOk = 0;
  const allFailures: Failure[] = [];

  for (const t of targets) {
    const { ok, failures } = await validateDir(t.dir, t.schema, t.label);
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
