import { promises as fs } from 'fs';
import path from 'path';
import matter from 'gray-matter';
import { z } from 'zod';

// O blog Plenya é a fonte canônica. Aqui no site pessoal espelhamos
// o conteúdo completo dos posts de autoria do Dr. Getúlio, mas todo
// SEO sinal (canonical, mainEntityOfPage) aponta para plenyasaude.com.br.

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

export const PLENYA_PILLAR_LABELS_EN: Record<(typeof PILLARS)[number], string> = {
  'alimentacao-atividade-fisica': 'Activity · Nutrition',
  'gestao-metabolica': 'Clinical Management',
  'integracao-corpo-mente': 'Mind-Body',
  'ritmo-circadiano': 'Circadian Rhythm',
  longevidade: 'Longevity',
};

export function pillarLabels(locale: string): Record<(typeof PILLARS)[number], string> {
  return locale === 'en' ? PLENYA_PILLAR_LABELS_EN : PLENYA_PILLAR_LABELS;
}

const dateString = z
  .union([z.string(), z.date()])
  .transform((v) => (v instanceof Date ? v.toISOString().slice(0, 10) : v));

const schema = z.object({
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
});

export type PlenyaPost = z.infer<typeof schema>;
export type PlenyaPostFull = PlenyaPost & { content: string; readingMinutes: number };

export const PLENYA_BASE = 'https://plenyasaude.com.br';
export const PLENYA_BLOG_BASE = `${PLENYA_BASE}/blog`;

// Caminho relativo a apps/site-getulio (cwd no build)
const PLENYA_BLOG_ROOT = path.join(process.cwd(), '..', 'site', 'content', 'blog', 'pt');

function estimateReadingMinutes(content: string): number {
  const words = content.trim().split(/\s+/).length;
  return Math.max(1, Math.round(words / 220));
}

/**
 * Reescreve paths relativos de imagem (`/images/blog/...`, `/images/...`)
 * para apontar absolutamente para o domínio canônico da Plenya.
 * Os assets vivem em apps/site/public — o site Getúlio não os hospeda.
 */
export function rewriteAssetPaths(content: string): string {
  return content
    .replace(/]\(\s*\/images\//g, `](${PLENYA_BASE}/images/`)
    .replace(/src="\s*\/images\//g, `src="${PLENYA_BASE}/images/`);
}

// Capa do post mora em apps/site/public — não é hospedada aqui.
// Reescreve `/images/...` para URL absoluta no domínio canônico Plenya.
export function absoluteCover(cover: string | undefined): string | undefined {
  if (!cover) return undefined;
  if (/^https?:\/\//.test(cover)) return cover;
  if (cover.startsWith('/')) return `${PLENYA_BASE}${cover}`;
  return cover;
}

async function readPostFile(file: string): Promise<PlenyaPostFull | null> {
  try {
    const raw = await fs.readFile(path.join(PLENYA_BLOG_ROOT, file), 'utf-8');
    const { data, content } = matter(raw);
    const parsed = schema.safeParse(data);
    if (!parsed.success) return null;
    return {
      ...parsed.data,
      content: rewriteAssetPaths(content),
      readingMinutes: estimateReadingMinutes(content),
    };
  } catch {
    return null;
  }
}

async function listPostFiles(): Promise<string[]> {
  try {
    return (await fs.readdir(PLENYA_BLOG_ROOT)).filter((f) => f.endsWith('.mdx'));
  } catch {
    return [];
  }
}

export async function getPlenyaPostsByGetulio(): Promise<PlenyaPost[]> {
  const files = await listPostFiles();
  const posts = await Promise.all(files.map(readPostFile));
  return posts
    .filter((p): p is PlenyaPostFull => p !== null && p.author === 'getulio-amaral')
    .sort((a, b) => (a.date < b.date ? 1 : -1))
    .map(({ content: _content, readingMinutes: _rm, ...meta }) => meta);
}

export async function getAllPlenyaPostsFull(): Promise<PlenyaPostFull[]> {
  const files = await listPostFiles();
  const posts = await Promise.all(files.map(readPostFile));
  return posts
    .filter((p): p is PlenyaPostFull => p !== null && p.author === 'getulio-amaral')
    .sort((a, b) => (a.date < b.date ? 1 : -1));
}

export async function getPlenyaPost(slug: string): Promise<PlenyaPostFull | null> {
  const post = await readPostFile(`${slug}.mdx`);
  if (!post || post.author !== 'getulio-amaral') return null;
  return post;
}
