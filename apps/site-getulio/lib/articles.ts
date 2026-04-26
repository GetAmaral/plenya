import { promises as fs } from 'fs';
import path from 'path';
import matter from 'gray-matter';
import { z } from 'zod';

export const TAGS = ['Nefrologia', 'Longevidade', 'Medicina funcional', 'Casos clínicos'] as const;
export type Tag = (typeof TAGS)[number];

const dateString = z
  .union([z.string(), z.date()])
  .transform((v) => (v instanceof Date ? v.toISOString().slice(0, 10) : v));

const articleSchema = z.object({
  slug: z.string(),
  title: z.string(),
  excerpt: z.string(),
  date: dateString,
  tag: z.enum(TAGS),
  source: z.string().optional(),
  readingMinutes: z.number().optional(),
});

export type ArticleFrontmatter = z.infer<typeof articleSchema>;
export type Article = ArticleFrontmatter & { content: string; readingMinutes: number };

const ROOT = path.join(process.cwd(), 'content', 'articles');

async function readDirSafe(dir: string) {
  try { return await fs.readdir(dir); } catch { return []; }
}

function estimateReadingMinutes(content: string): number {
  const words = content.trim().split(/\s+/).length;
  return Math.max(1, Math.round(words / 220));
}

export async function getAllArticles(): Promise<Article[]> {
  const files = (await readDirSafe(ROOT)).filter((f) => f.endsWith('.mdx'));
  const articles = await Promise.all(
    files.map(async (file) => {
      const raw = await fs.readFile(path.join(ROOT, file), 'utf-8');
      const { data, content } = matter(raw);
      const parsed = articleSchema.parse(data);
      return {
        ...parsed,
        content,
        readingMinutes: parsed.readingMinutes ?? estimateReadingMinutes(content),
      };
    }),
  );
  return articles.sort((a, b) => (a.date < b.date ? 1 : -1));
}

export async function getArticle(slug: string): Promise<Article | null> {
  const all = await getAllArticles();
  return all.find((a) => a.slug === slug) ?? null;
}

export async function getRecentArticles(limit = 3): Promise<Article[]> {
  const all = await getAllArticles();
  return all.slice(0, limit);
}
