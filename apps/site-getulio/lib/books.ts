import { promises as fs } from 'fs';
import path from 'path';
import matter from 'gray-matter';
import { z } from 'zod';

const excerptSchema = z.object({
  cap: z.string(),
  citacao: z.string(),
});

const editionSchema = z.object({
  id: z.string(),
  format: z.enum(['EBook', 'Paperback', 'Hardcover', 'AudiobookFormat']),
  language: z.string(),
  isbn: z.string(),
  label: z.string(),
  labelEn: z.string().optional(),
  purchaseUrl: z.string().url().optional(),
});

export type Edition = z.infer<typeof editionSchema>;

const bookSchema = z.object({
  slug: z.string(),
  order: z.number().default(99),
  year: z.number(),
  isbn: z.string(),
  cover: z.string(),
  amazonUrl: z.string().url(),
  amazonUrlEn: z.string().url().optional(),
  hotmartUrl: z.string().url(),
  editions: z.array(editionSchema).optional(),

  title: z.string(),
  subtitle: z.string(),
  shortDescription: z.string(),
  edition: z.string(),
  lead1: z.string(),
  lead2: z.string(),
  excerpts: z.array(excerptSchema).min(1),

  titleEn: z.string().optional(),
  subtitleEn: z.string().optional(),
  shortDescriptionEn: z.string().optional(),
  editionEn: z.string().optional(),
  lead1En: z.string().optional(),
  lead2En: z.string().optional(),
  excerptsEn: z.array(excerptSchema).optional(),
});

export type BookFrontmatter = z.infer<typeof bookSchema>;
export type Book = BookFrontmatter & { content: string };

const ROOT = path.join(process.cwd(), 'content', 'livros');

async function readDirSafe(dir: string) {
  try { return await fs.readdir(dir); } catch { return []; }
}

export async function getAllBooks(): Promise<Book[]> {
  const files = (await readDirSafe(ROOT)).filter((f) => f.endsWith('.mdx'));
  const books = await Promise.all(
    files.map(async (file) => {
      const raw = await fs.readFile(path.join(ROOT, file), 'utf-8');
      const { data, content } = matter(raw);
      const parsed = bookSchema.parse(data);
      return { ...parsed, content };
    }),
  );
  return books.sort((a, b) => a.order - b.order);
}

export async function getBook(slug: string): Promise<Book | null> {
  const all = await getAllBooks();
  return all.find((b) => b.slug === slug) ?? null;
}

/**
 * Resolve fields in the requested locale, falling back to PT when the
 * EN counterpart is missing. Amazon URL has a locale-specific variant
 * (PT → amazon.com.br, EN → amazon.com); Hotmart is BR-only.
 */
export function localizedBook(b: Book, locale: string) {
  const isEn = locale === 'en';
  return {
    title: isEn && b.titleEn ? b.titleEn : b.title,
    subtitle: isEn && b.subtitleEn ? b.subtitleEn : b.subtitle,
    shortDescription: isEn && b.shortDescriptionEn ? b.shortDescriptionEn : b.shortDescription,
    edition: isEn && b.editionEn ? b.editionEn : b.edition,
    lead1: isEn && b.lead1En ? b.lead1En : b.lead1,
    lead2: isEn && b.lead2En ? b.lead2En : b.lead2,
    excerpts: isEn && b.excerptsEn ? b.excerptsEn : b.excerpts,
    amazonUrl: isEn && b.amazonUrlEn ? b.amazonUrlEn : b.amazonUrl,
  };
}
