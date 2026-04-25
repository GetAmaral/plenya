import { promises as fs } from 'fs';
import path from 'path';
import matter from 'gray-matter';
import { z } from 'zod';

const plenyaScoreSchema = z
  .object({
    initial: z.number().min(0).max(100).optional(),
    current: z.number().min(0).max(100).optional(),
    months: z.number().min(0).max(120).optional(),
  })
  .optional();

const schema = z.object({
  slug: z.string(),
  patientLabel: z.string(),
  context: z.string().optional(),
  quote: z.string(),
  outcome: z.string().optional(),
  duration: z.string().optional(),
  plenyaScore: plenyaScoreSchema,
  pillar: z
    .enum(['alimentacao-atividade-fisica', 'gestao-metabolica', 'integracao-corpo-mente', 'ritmo-circadiano', 'longevidade'])
    .optional(),
  consent: z.boolean(),
  date: z.union([z.string(), z.date()]).transform((v) => (v instanceof Date ? v.toISOString().slice(0, 10) : v)),
  featured: z.boolean().default(false),
  order: z.number().default(99),
});

export type Testimonial = z.infer<typeof schema>;

const ROOT = path.join(process.cwd(), 'content', 'testimonials');

async function readDirSafe(dir: string) {
  try {
    return await fs.readdir(dir);
  } catch {
    return [];
  }
}

export async function getAllTestimonials(): Promise<Testimonial[]> {
  const files = (await readDirSafe(ROOT)).filter((f) => f.endsWith('.mdx') || f.endsWith('.md'));
  const items = await Promise.all(
    files.map(async (file) => {
      const raw = await fs.readFile(path.join(ROOT, file), 'utf-8');
      const { data } = matter(raw);
      return schema.parse(data);
    }),
  );
  return items
    .filter((t) => t.consent)
    .sort((a, b) => {
      if (a.featured !== b.featured) return a.featured ? -1 : 1;
      if (a.order !== b.order) return a.order - b.order;
      return a.date < b.date ? 1 : -1;
    });
}

export async function getFeaturedTestimonials(limit = 3): Promise<Testimonial[]> {
  const all = await getAllTestimonials();
  return all.filter((t) => t.featured).slice(0, limit);
}
