import type { Metadata } from 'next';
import { notFound } from 'next/navigation';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { Breadcrumbs } from '@/components/layout/breadcrumbs';
import { MdxContent } from '@/components/blog/mdx-content';
import { getAllLectures, getLecture, getAudienceLabel, localizedLecture, sortAudience } from '@/lib/lectures';
import { LectureSchema } from '@/components/seo/lecture-schema';
import { BreadcrumbSchema } from '@/components/seo/breadcrumb-schema';
import { EducationalNotice } from '@/components/legal/educational-notice';

export async function generateStaticParams() {
  const lectures = await getAllLectures();
  return lectures.map((l) => ({ slug: l.slug }));
}

export async function generateMetadata({
  params,
}: {
  params: Promise<{ locale: string; slug: string }>;
}): Promise<Metadata> {
  const { locale, slug } = await params;
  const lecture = await getLecture(slug);
  if (!lecture) return {};
  const loc = localizedLecture(lecture, locale);
  const url = locale === 'en' ? `/en/palestras/${slug}` : `/palestras/${slug}`;
  return {
    title: loc.title,
    description: loc.excerpt,
    alternates: {
      canonical: url,
      languages: { 'pt-BR': `/palestras/${slug}`, en: `/en/palestras/${slug}` },
    },
    openGraph: {
      type: 'article',
      url,
      title: loc.title,
      description: loc.excerpt,
      images: ['/images/getulio-square.jpg'],
    },
    twitter: {
      card: 'summary_large_image',
      title: loc.title,
      description: loc.excerpt,
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
  const t = await getTranslations({ locale, namespace: 'palestras' });
  const loc = localizedLecture(lecture, locale);
  const subjectPrefix = t('emailSubjectPrefix');
  const mailto = `mailto:palestras@drgetulioamaralfilho.com.br?subject=${encodeURIComponent(`${subjectPrefix} — ${loc.title}`)}`;

  return (
    <article>
      <LectureSchema
        title={loc.title}
        description={loc.excerpt}
        slug={lecture.slug}
        duration={loc.duration}
        audience={sortAudience(lecture.audience).map((a) => getAudienceLabel(a, locale))}
        locale={locale}
      />
      <BreadcrumbSchema
        items={[
          { name: t('detailBreadcrumbHome'), url: '/' },
          { name: t('detailBreadcrumbList'), url: '/palestras' },
          { name: loc.title },
        ]}
      />
      <header className="editorial-container pt-12 md:pt-16 pb-12">
        <Breadcrumbs
          items={[
            { label: t('detailBreadcrumbHome'), href: '/' },
            { label: t('detailBreadcrumbList'), href: '/palestras' },
            { label: loc.title },
          ]}
        />
      </header>

      <section className="editorial-container pb-16">
        <div className="max-w-3xl space-y-8">
          {lecture.anchor && <p className="label-meta text-bordo">{t('anchorBadge')}</p>}
          <h1 className="heading-display text-[clamp(2.2rem,5vw,3.8rem)]">{loc.title}</h1>
          {loc.subtitle && (
            <p className="font-serif italic text-ink-muted text-xl md:text-2xl">{loc.subtitle}</p>
          )}
          <p className="prose-body">{loc.excerpt}</p>
          {loc.content && (
            <div className="prose-body">
              <MdxContent source={loc.content.trim()} />
            </div>
          )}
        </div>
      </section>

      <section className="border-t border-rule">
        <div className="editorial-container py-16 grid md:grid-cols-3 gap-10">
          <div>
            <p className="label-meta mb-3">{t('durationLabel')}</p>
            <p className="font-serif text-ink-soft">{loc.duration}</p>
          </div>
          <div>
            <p className="label-meta mb-3">{t('formatLabel')}</p>
            <p className="font-serif text-ink-soft">{loc.format}</p>
          </div>
          <div>
            <p className="label-meta mb-3">{t('audienceLabel')}</p>
            <ul className="space-y-1">
              {sortAudience(lecture.audience).map((a) => (
                <li key={a} className="font-serif text-ink-soft">{getAudienceLabel(a, locale)}</li>
              ))}
            </ul>
          </div>
        </div>
      </section>

      <EducationalNotice />

      <section className="border-t border-rule bg-paper">
        <div className="editorial-container py-16 md:py-20 grid md:grid-cols-[1fr_2fr] gap-10 items-start">
          <p className="label-meta">{t('detailInviteKicker')}</p>
          <div className="space-y-3 font-serif text-ink-soft">
            <p>{t('detailInviteBody')}</p>
            <p className="font-sans text-sm">
              <a href={mailto} className="link-text">
                palestras@drgetulioamaralfilho.com.br
              </a>
            </p>
          </div>
        </div>
      </section>
    </article>
  );
}
