import { promises as fs } from 'fs';
import path from 'path';
import matter from 'gray-matter';
import { z } from 'zod';

const CATEGORIES = ['medico', 'nutricao', 'psicologia', 'educacao_fisica'] as const;
export type Category = (typeof CATEGORIES)[number];

export const categoryLabels: Record<Category, string> = {
  medico: 'Médicos',
  nutricao: 'Nutrição',
  psicologia: 'Psicologia',
  educacao_fisica: 'Educação Física',
};

export const categoryOrder: Category[] = ['medico', 'nutricao', 'psicologia', 'educacao_fisica'];

const frontmatterSchema = z.object({
  name: z.string(),
  slug: z.string(),
  role: z.string(),
  category: z.enum(CATEGORIES),
  order: z.number().default(99),
  credentials: z.string(),
  rqe: z.union([z.string(), z.number()]).transform(String).optional(),
  specialties: z.array(z.string()).default([]),
  shortBio: z.string(),
  photo: z.string().optional(),
  instagram: z.string().url().optional(),
  linkedin: z.string().url().optional(),
  featured: z.boolean().default(false),
});

export type DoctorFrontmatter = z.infer<typeof frontmatterSchema>;
export type Doctor = DoctorFrontmatter & { bio: string };

const ROOT = path.join(process.cwd(), 'content', 'doctors');

async function readDirSafe(dir: string) {
  try {
    return await fs.readdir(dir);
  } catch {
    return [];
  }
}

export async function getAllDoctors(): Promise<Doctor[]> {
  const files = (await readDirSafe(ROOT)).filter((f) => f.endsWith('.mdx'));
  const doctors = await Promise.all(
    files.map(async (file) => {
      const raw = await fs.readFile(path.join(ROOT, file), 'utf-8');
      const { data, content } = matter(raw);
      const parsed = frontmatterSchema.parse(data);
      return { ...parsed, bio: content } satisfies Doctor;
    }),
  );
  return doctors.sort((a, b) => {
    const ca = categoryOrder.indexOf(a.category);
    const cb = categoryOrder.indexOf(b.category);
    if (ca !== cb) return ca - cb;
    return a.order - b.order;
  });
}

export async function getDoctor(slug: string): Promise<Doctor | null> {
  const all = await getAllDoctors();
  return all.find((d) => d.slug === slug) ?? null;
}

export async function getDoctorsByCategory(): Promise<Record<Category, Doctor[]>> {
  const all = await getAllDoctors();
  return categoryOrder.reduce(
    (acc, cat) => {
      acc[cat] = all.filter((d) => d.category === cat);
      return acc;
    },
    {} as Record<Category, Doctor[]>,
  );
}
