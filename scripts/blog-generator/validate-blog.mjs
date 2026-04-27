import { promises as fs } from 'fs';
import path from 'path';
import matter from 'gray-matter';
import { z } from 'zod';

const PILLARS = ['alimentacao-atividade-fisica', 'gestao-metabolica', 'integracao-corpo-mente', 'ritmo-circadiano', 'longevidade'];
const dateString = z.union([z.string(), z.date()]).transform((v) => v instanceof Date ? v.toISOString().slice(0,10) : v);
const schema = z.object({
  title: z.string(), slug: z.string(), excerpt: z.string(),
  date: dateString, updated: dateString.optional(),
  author: z.string(), reviewedBy: z.string().optional(),
  pillar: z.enum(PILLARS), tags: z.array(z.string()).default([]),
  cover: z.string().optional(), featured: z.boolean().default(false),
  references: z.array(z.object({ label: z.string(), url: z.string().url() })).default([]),
  readingMinutes: z.number().optional(),
  cta: z.enum(['default','recognition']).default('default'),
});

const dir = '/home/user/plenya/apps/site/content/blog/pt';
const files = (await fs.readdir(dir)).filter(f => f.endsWith('.mdx'));
let bad = 0;
for (const f of files) {
  const raw = await fs.readFile(path.join(dir, f), 'utf-8');
  const { data } = matter(raw);
  try {
    const p = schema.parse(data);
    console.log(`OK  ${f}  (${p.references.length} refs)`);
  } catch (e) {
    bad++;
    console.log(`ERR ${f}: ${e.message}`);
  }
}
process.exit(bad ? 1 : 0);
