import { promises as fs } from 'fs';
import path from 'path';
import matter from 'gray-matter';
import { z } from 'zod';

// Conteúdo editorial do Dr. Getúlio. Antes (até 2026-05) os artigos viviam
// em apps/site/content/blog (Plenya, canonical) e este site servia mirror.
// Agora (cisão por audiência) os artigos vivem aqui — versão editorial em 1a
// pessoa — e o Plenya mantém versão paralela com tom institucional.
// Cross-links recíprocos no rodapé de cada post.

const PILLARS = [
  'alimentacao-atividade-fisica',
  'gestao-metabolica',
  'integracao-corpo-mente',
  'ritmo-circadiano',
  'longevidade',
] as const;

export const PILLAR_LABELS_PT: Record<(typeof PILLARS)[number], string> = {
  'alimentacao-atividade-fisica': 'Atividade · Alimentação',
  'gestao-metabolica': 'Gestão Clínica',
  'integracao-corpo-mente': 'Mente-Corpo',
  'ritmo-circadiano': 'Ritmo Circadiano',
  longevidade: 'Longevidade',
};

export const PILLAR_LABELS_EN: Record<(typeof PILLARS)[number], string> = {
  'alimentacao-atividade-fisica': 'Activity · Nutrition',
  'gestao-metabolica': 'Clinical Management',
  'integracao-corpo-mente': 'Mind-Body',
  'ritmo-circadiano': 'Circadian Rhythm',
  longevidade: 'Longevity',
};

export function pillarLabels(locale: string): Record<(typeof PILLARS)[number], string> {
  return locale === 'en' ? PILLAR_LABELS_EN : PILLAR_LABELS_PT;
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
  // Quando true, mostra cross-link pra versão Plenya do mesmo post.
  hasMirror: z.boolean().default(true),
});

export type Post = z.infer<typeof schema>;
export type PostFull = Post & { content: string; readingMinutes: number };

// Cross-link helpers — ainda apontam pra Plenya como destino, mas agora
// como "versão clínica para profissionais", não como canonical.
export const PLENYA_BASE = 'https://plenyasaude.com.br';
export const PLENYA_BLOG_BASE = `${PLENYA_BASE}/blog`;
export const PLENYA_BLOG_BASE_EN = `${PLENYA_BASE}/en/blog`;

export function plenyaMirrorUrl(slug: string, locale: string): string {
  const base = locale === 'en' ? PLENYA_BLOG_BASE_EN : PLENYA_BLOG_BASE;
  return `${base}/${slug}`;
}

function blogRoot(locale: string): string {
  const lang = locale === 'en' ? 'en' : 'pt';
  return path.join(process.cwd(), 'content', 'blog', lang);
}

function estimateReadingMinutes(content: string): number {
  const words = content.trim().split(/\s+/).length;
  return Math.max(1, Math.round(words / 220));
}

/**
 * Reescreve paths relativos de imagem (`/images/blog/...`, `/images/...`)
 * para apontar absolutamente para o domínio Plenya. Os assets vivem em
 * apps/site/public — o site Getúlio reusa as mesmas figuras (pedido do autor:
 * "mantendo as mesmas figuras"). Evita duplicação de ~100MB de imagens.
 */
export function rewriteAssetPaths(content: string): string {
  return content
    .replace(/]\(\s*\/images\//g, `](${PLENYA_BASE}/images/`)
    .replace(/src="\s*\/images\//g, `src="${PLENYA_BASE}/images/`);
}

export function absoluteCover(cover: string | undefined): string | undefined {
  if (!cover) return undefined;
  if (/^https?:\/\//.test(cover)) return cover;
  if (cover.startsWith('/')) return `${PLENYA_BASE}${cover}`;
  return cover;
}

async function readPostFile(root: string, file: string): Promise<PostFull | null> {
  try {
    const raw = await fs.readFile(path.join(root, file), 'utf-8');
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

async function listPostFiles(root: string): Promise<string[]> {
  try {
    return (await fs.readdir(root)).filter((f) => f.endsWith('.mdx'));
  } catch {
    return [];
  }
}

export async function getAllPosts(locale: string = 'pt'): Promise<Post[]> {
  const root = blogRoot(locale);
  const files = await listPostFiles(root);
  const posts = await Promise.all(files.map((f) => readPostFile(root, f)));
  return posts
    .filter((p): p is PostFull => p !== null)
    .sort((a, b) => (a.date < b.date ? 1 : -1))
    .map(({ content: _content, readingMinutes: _rm, ...meta }) => meta);
}

export async function getAllPostsFull(locale: string = 'pt'): Promise<PostFull[]> {
  const root = blogRoot(locale);
  const files = await listPostFiles(root);
  const posts = await Promise.all(files.map((f) => readPostFile(root, f)));
  return posts
    .filter((p): p is PostFull => p !== null)
    .sort((a, b) => (a.date < b.date ? 1 : -1));
}

export async function getPost(slug: string, locale: string = 'pt'): Promise<PostFull | null> {
  const root = blogRoot(locale);
  let post = await readPostFile(root, `${slug}.mdx`);
  // Fallback: se locale=en não tiver, devolve PT.
  if (!post && locale !== 'pt') {
    post = await readPostFile(blogRoot('pt'), `${slug}.mdx`);
  }
  return post;
}

// ===== Backward-compat shims (apagar após confirmar todos os imports atualizados) =====
export const getAllPlenyaPostsFull = getAllPostsFull;
export const getPlenyaPostsByGetulio = getAllPosts;
export const getPlenyaPost = getPost;
export const plenyaBlogBase = (locale: string) =>
  locale === 'en' ? PLENYA_BLOG_BASE_EN : PLENYA_BLOG_BASE;
