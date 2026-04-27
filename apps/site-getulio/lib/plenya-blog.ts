import { promises as fs } from 'fs';
import path from 'path';
import matter from 'gray-matter';
import { z } from 'zod';

// O blog Plenya é a fonte única; lemos os MDX no build time e listamos
// só os artigos de autoria do Dr. Getúlio aqui no site pessoal,
// linkando out para a URL canônica em plenyasaude.com.br.

const PILLARS = [
  'alimentacao-atividade-fisica',
  'gestao-metabolica',
  'integracao-corpo-mente',
  'ritmo-circadiano',
  'longevidade',
] as const;

export const PLENYA_PILLAR_LABELS: Record<(typeof PILLARS)[number], string> = {
  'alimentacao-atividade-fisica': 'Atividade · Alimentação',
  'gestao-metabolica': 'Gestão Clínica',
  'integracao-corpo-mente': 'Mente-Corpo',
  'ritmo-circadiano': 'Ritmo Circadiano',
  longevidade: 'Longevidade',
};

const dateString = z
  .union([z.string(), z.date()])
  .transform((v) => (v instanceof Date ? v.toISOString().slice(0, 10) : v));

const schema = z.object({
  title: z.string(),
  slug: z.string(),
  excerpt: z.string(),
  date: dateString,
  author: z.string(),
  pillar: z.enum(PILLARS),
  cover: z.string().optional(),
});

export type PlenyaPost = z.infer<typeof schema>;

export const PLENYA_BLOG_BASE = 'https://plenyasaude.com.br/blog';

// Caminho relativo a apps/site-getulio (cwd no build)
const PLENYA_BLOG_ROOT = path.join(process.cwd(), '..', 'site', 'content', 'blog', 'pt');

export async function getPlenyaPostsByGetulio(): Promise<PlenyaPost[]> {
  let files: string[] = [];
  try {
    files = (await fs.readdir(PLENYA_BLOG_ROOT)).filter((f) => f.endsWith('.mdx'));
  } catch {
    return [];
  }
  const posts = await Promise.all(
    files.map(async (file) => {
      try {
        const raw = await fs.readFile(path.join(PLENYA_BLOG_ROOT, file), 'utf-8');
        const { data } = matter(raw);
        const parsed = schema.safeParse(data);
        return parsed.success ? parsed.data : null;
      } catch {
        return null;
      }
    }),
  );
  return posts
    .filter((p): p is PlenyaPost => p !== null && p.author === 'getulio-amaral')
    .sort((a, b) => (a.date < b.date ? 1 : -1));
}
