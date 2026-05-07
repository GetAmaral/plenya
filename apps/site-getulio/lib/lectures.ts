import { promises as fs } from 'fs';
import path from 'path';
import matter from 'gray-matter';
import { z } from 'zod';

const audienceEnum = z.enum(['corporativo', 'medicos', 'residentes', 'aberto', 'congressos']);
export type Audience = z.infer<typeof audienceEnum>;

const audienceLabelsPt: Record<Audience, string> = {
  corporativo: 'Corporativo',
  medicos: 'Médicos',
  residentes: 'Residentes',
  aberto: 'Aberto',
  congressos: 'Congressos',
};

const audienceLabelsEn: Record<Audience, string> = {
  corporativo: 'Corporate',
  medicos: 'Physicians',
  residentes: 'Residents',
  aberto: 'Open',
  congressos: 'Conferences',
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
  // Campos EN opcionais — fallback automático pra PT quando ausentes.
  titleEn: z.string().optional(),
  subtitleEn: z.string().optional(),
  excerptEn: z.string().optional(),
  bodyEn: z.string().optional(),
  durationEn: z.string().optional(),
  formatEn: z.string().optional(),
});

export type LectureFrontmatter = z.infer<typeof lectureSchema>;
export type Lecture = LectureFrontmatter & { content: string };

const ROOT = path.join(process.cwd(), 'content', 'lectures');

async function readDirSafe(dir: string) {
  try { return await fs.readdir(dir); } catch { return []; }
}

export function getAudienceLabel(a: Audience, locale: string = 'pt') {
  return locale === 'en' ? audienceLabelsEn[a] : audienceLabelsPt[a];
}

/**
 * Prioridade de exibição: público leigo primeiro, corporativo depois,
 * técnico (médicos/residentes/congressos) por último. Reflete o foco
 * editorial do site.
 */
const AUDIENCE_PRIORITY: Record<Audience, number> = {
  aberto: 1,
  corporativo: 2,
  congressos: 3,
  medicos: 4,
  residentes: 5,
};

export function sortAudience(audience: Audience[]): Audience[] {
  return [...audience].sort((a, b) => AUDIENCE_PRIORITY[a] - AUDIENCE_PRIORITY[b]);
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

/**
 * Resolve campos exibíveis no locale pedido, com fallback PT quando EN
 * está ausente. O `bodyEn` (quando existe) substitui o corpo MDX.
 */
export function localizedLecture(l: Lecture, locale: string) {
  const isEn = locale === 'en';
  return {
    title: isEn && l.titleEn ? l.titleEn : l.title,
    subtitle: isEn && l.subtitleEn ? l.subtitleEn : l.subtitle,
    excerpt: isEn && l.excerptEn ? l.excerptEn : l.excerpt,
    duration: isEn && l.durationEn ? l.durationEn : l.duration,
    format: isEn && l.formatEn ? l.formatEn : l.format,
    content: isEn && l.bodyEn ? l.bodyEn : l.content,
  };
}
