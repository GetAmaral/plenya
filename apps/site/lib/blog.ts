import { promises as fs } from 'fs';
import path from 'path';
import matter from 'gray-matter';
import { z } from 'zod';
import type { Locale } from './i18n/config';

const PILLARS = ['alimentacao-atividade-fisica', 'gestao-metabolica', 'integracao-corpo-mente', 'ritmo-circadiano', 'longevidade'] as const;
export type Pillar = (typeof PILLARS)[number];
export const pillars = PILLARS;

export const pillarLabels: Record<Pillar, string> = {
  'alimentacao-atividade-fisica': 'Alimentação & Atividade Física',
  'gestao-metabolica': 'Gestão Metabólica',
  'integracao-corpo-mente': 'Integração Corpo & Mente',
  'ritmo-circadiano': 'Ritmo Circadiano',
  longevidade: 'Longevidade',
};

const dateString = z
  .union([z.string(), z.date()])
  .transform((value) => (value instanceof Date ? value.toISOString().slice(0, 10) : value));

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
    .array(z.object({ label: z.string(), url: z.string().url() }))
    .default([]),
  readingMinutes: z.number().optional(),
});

export type PostFrontmatter = z.infer<typeof frontmatterSchema>;

export type Post = PostFrontmatter & {
  locale: Locale;
  content: string;
  readingMinutes: number;
};

const ROOT = path.join(process.cwd(), 'content', 'blog');

async function readDirSafe(dir: string) {
  try {
    return await fs.readdir(dir);
  } catch {
    return [];
  }
}

function estimateReadingMinutes(content: string): number {
  const words = content.trim().split(/\s+/).length;
  return Math.max(1, Math.round(words / 220));
}

export async function getAllPosts(locale: Locale): Promise<Post[]> {
  const dir = path.join(ROOT, locale);
  const files = (await readDirSafe(dir)).filter((f) => f.endsWith('.mdx'));
  const posts = await Promise.all(
    files.map(async (file) => {
      const raw = await fs.readFile(path.join(dir, file), 'utf-8');
      const { data, content } = matter(raw);
      const parsed = frontmatterSchema.parse(data);
      return {
        ...parsed,
        locale,
        content,
        readingMinutes: parsed.readingMinutes ?? estimateReadingMinutes(content),
      } satisfies Post;
    }),
  );
  return posts.sort((a, b) => (a.date < b.date ? 1 : -1));
}

export async function getPost(locale: Locale, slug: string): Promise<Post | null> {
  const posts = await getAllPosts(locale);
  return posts.find((p) => p.slug === slug) ?? null;
}

export async function getPostsByPillar(locale: Locale, pillar: Pillar): Promise<Post[]> {
  const posts = await getAllPosts(locale);
  return posts.filter((p) => p.pillar === pillar);
}

export async function getFeaturedPost(locale: Locale): Promise<Post | null> {
  const posts = await getAllPosts(locale);
  return posts.find((p) => p.featured) ?? posts[0] ?? null;
}
