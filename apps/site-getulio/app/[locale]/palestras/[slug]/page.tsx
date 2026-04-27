import type { Metadata } from 'next';
import { notFound } from 'next/navigation';
import { setRequestLocale } from 'next-intl/server';
import { Breadcrumbs } from '@/components/layout/breadcrumbs';
import { getAllLectures, getLecture, getAudienceLabel } from '@/lib/lectures';
import { LectureSchema } from '@/components/seo/lecture-schema';
import { BreadcrumbSchema } from '@/components/seo/breadcrumb-schema';

export async function generateStaticParams() {
  const lectures = await getAllLectures();
  return lectures.map((l) => ({ slug: l.slug }));
}

export async function generateMetadata({
  params,
}: {
  params: Promise<{ slug: string }>;
}): Promise<Metadata> {
  const { slug } = await params;
  const lecture = await getLecture(slug);
  if (!lecture) return {};
  const url = `/palestras/${slug}`;
  return {
    title: lecture.title,
    description: lecture.excerpt,
    alternates: { canonical: url },
    openGraph: {
      type: 'article',
      url,
      title: lecture.title,
      description: lecture.excerpt,
      images: ['/images/getulio-square.jpg'],
    },
    twitter: {
      card: 'summary_large_image',
      title: lecture.title,
      description: lecture.excerpt,
      images: ['/images/getulio-square.jpg'],
    },
  };
}

export default async function LecturePage({
  params,
}: {
  params: Promise<{ locale: string; slug: string }>;
}) {
  const { locale, slug } = await params;
  setRequestLocale(locale);
  const lecture = await getLecture(slug);
  if (!lecture) notFound();

  return (
    <article>
      <LectureSchema
        title={lecture.title}
        description={lecture.excerpt}
        slug={lecture.slug}
        duration={lecture.duration}
        audience={lecture.audience.map(getAudienceLabel)}
      />
      <BreadcrumbSchema
        items={[
          { name: 'Início', url: '/' },
          { name: 'Palestras', url: '/palestras' },
          { name: lecture.title },
        ]}
      />
      <header className="editorial-container pt-12 md:pt-16 pb-12">
        <Breadcrumbs
          items={[
            { label: 'Início', href: '/' },
            { label: 'Palestras', href: '/palestras' },
            { label: lecture.title },
          ]}
        />
      </header>

      <section className="editorial-container pb-16">
        <div className="max-w-3xl space-y-8">
          {lecture.anchor && <p className="label-meta text-bordo">Palestra-âncora</p>}
          <h1 className="heading-display text-[clamp(2.2rem,5vw,3.8rem)]">{lecture.title}</h1>
          {lecture.subtitle && (
            <p className="font-serif italic text-ink-muted text-xl md:text-2xl">{lecture.subtitle}</p>
          )}
          <p className="prose-body">{lecture.excerpt}</p>
          {lecture.content && (
            <div className="prose-body whitespace-pre-line">{lecture.content.trim()}</div>
          )}
        </div>
      </section>

      <section className="border-t border-rule">
        <div className="editorial-container py-16 grid md:grid-cols-3 gap-10">
          <div>
            <p className="label-meta mb-3">Duração</p>
            <p className="font-serif text-ink-soft">{lecture.duration}</p>
          </div>
          <div>
            <p className="label-meta mb-3">Formato</p>
            <p className="font-serif text-ink-soft">{lecture.format}</p>
          </div>
          <div>
            <p className="label-meta mb-3">Público</p>
            <ul className="space-y-1">
              {lecture.audience.map((a) => (
                <li key={a} className="font-serif text-ink-soft">{getAudienceLabel(a)}</li>
              ))}
            </ul>
          </div>
        </div>
      </section>

      <section className="border-t border-rule bg-paper">
        <div className="editorial-container py-16 md:py-20 grid md:grid-cols-[1fr_2fr] gap-10 items-start">
          <p className="label-meta">Convidar</p>
          <div className="space-y-3 font-serif text-ink-soft">
            <p>
              Para convidar para esta palestra, envie proposta com nome do evento, data, local,
              público estimado e formato (presencial ou remoto).
            </p>
            <p className="font-sans text-sm">
              <a href={`mailto:palestras@drgetulioamaralfilho.com.br?subject=Palestra%20—%20${encodeURIComponent(lecture.title)}`} className="link-text">
                palestras@drgetulioamaralfilho.com.br
              </a>
            </p>
          </div>
        </div>
      </section>
    </article>
  );
}
