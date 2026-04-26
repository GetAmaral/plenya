import { promises as fs } from 'fs';
import path from 'path';
import matter from 'gray-matter';
import { z } from 'zod';

const audienceEnum = z.enum(['corporativo', 'medicos', 'residentes', 'aberto', 'congressos']);
export type Audience = z.infer<typeof audienceEnum>;

const audienceLabels: Record<Audience, string> = {
  corporativo: 'Corporativo',
  medicos: 'Médicos',
  residentes: 'Residentes',
  aberto: 'Aberto',
  congressos: 'Congressos',
};

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
});

export type LectureFrontmatter = z.infer<typeof lectureSchema>;
export type Lecture = LectureFrontmatter & { content: string };

const ROOT = path.join(process.cwd(), 'content', 'lectures');

async function readDirSafe(dir: string) {
  try { return await fs.readdir(dir); } catch { return []; }
}

export function getAudienceLabel(a: Audience) {
  return audienceLabels[a];
}

export async function getAllLectures(): Promise<Lecture[]> {
  const files = (await readDirSafe(ROOT)).filter((f) => f.endsWith('.mdx'));
  const lectures = await Promise.all(
    files.map(async (file) => {
      const raw = await fs.readFile(path.join(ROOT, file), 'utf-8');
      const { data, content } = matter(raw);
      const parsed = lectureSchema.parse(data);
      return { ...parsed, content };
    }),
  );
  return lectures.sort((a, b) => a.order - b.order);
}

export async function getLecture(slug: string): Promise<Lecture | null> {
  const all = await getAllLectures();
  return all.find((l) => l.slug === slug) ?? null;
}
