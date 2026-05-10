import type { Metadata } from 'next';
import { notFound } from 'next/navigation';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { Breadcrumbs } from '@/components/layout/breadcrumbs';
import { EducationalNotice } from '@/components/legal/educational-notice';
import { BreadcrumbSchema } from '@/components/seo/breadcrumb-schema';
import { getAllBooks, getBook, localizedBook } from '@/lib/books';

export async function generateStaticParams() {
  const books = await getAllBooks();
  return books.map((b) => ({ slug: b.slug }));
}

export async function generateMetadata({
  params,
}: {
  params: Promise<{ locale: string; slug: string }>;
}): Promise<Metadata> {
  const { locale, slug } = await params;
  const book = await getBook(slug);
  if (!book) return {};
  const loc = localizedBook(book, locale);
  const t = await getTranslations({ locale, namespace: 'book' });
  return {
    title: `${t('excerptsKicker')} — ${loc.title}`,
    description: loc.shortDescription,
    alternates: {
      canonical: locale === 'en' ? `/en/books/${slug}/excerpts` : `/livros/${slug}/excertos`,
      languages: {
        'pt-BR': `/livros/${slug}/excertos`, pt: `/livros/${slug}/excertos`,
        en: `/en/books/${slug}/excerpts`,
      },
    },
  };
}

export default async function ExcertosPage({
  params,
}: {
  params: Promise<{ locale: string; slug: string }>;
}) {
  const { locale, slug } = await params;
  setRequestLocale(locale);
  const book = await getBook(slug);
  if (!book) notFound();
  const t = await getTranslations({ locale, namespace: 'book' });
  const loc = localizedBook(book, locale);
  const hotmartUrl = `${book.hotmartUrl}?src=excertos`;
  const homeLabel = locale === 'en' ? 'Home' : 'Início';
  const booksLabel = locale === 'en' ? 'Books' : 'Livros';

  return (
    <article>
      <BreadcrumbSchema
        items={[
          { name: homeLabel, url: '/' },
          { name: booksLabel, url: '/livros' },
          { name: loc.title, url: `/livros/${slug}` },
          { name: t('excerptsKicker') },
        ]}
      />

      <header className="editorial-container pt-12 md:pt-16 pb-8">
        <Breadcrumbs
          items={[
            { label: homeLabel, href: '/' },
            { label: booksLabel, href: '/livros' },
            { label: loc.title, href: `/livros/${slug}` },
            { label: t('excerptsKicker') },
          ]}
        />
      </header>

      <section className="editorial-narrow pb-16">
        <div className="space-y-8">
          <p className="label-meta-lg text-bordo">{t('excerptsKicker')} · {loc.title}</p>
          <h1 className="heading-display text-[clamp(2.4rem,5vw,4rem)]">
            {t('excerptsH1')}
          </h1>
          <p className="font-serif italic text-ink-soft text-xl md:text-2xl leading-relaxed">
            {t('excerptsLead')}
          </p>
        </div>
      </section>

      <section className="border-t border-rule bg-paper">
        <div className="editorial-narrow py-20 md:py-28 space-y-24 md:space-y-32">
          {loc.excerpts.map((tr, i) => (
            <figure key={i} className="space-y-6">
              <p className="label-meta-lg text-bordo">{tr.cap}</p>
              <blockquote className="font-serif text-2xl md:text-3xl leading-snug text-ink italic border-l-2 border-bordo pl-6 md:pl-10">
                {tr.citacao}
              </blockquote>
            </figure>
          ))}
        </div>
      </section>

      <EducationalNotice />

      <section className="border-t border-rule bg-paper">
        <div className="editorial-container py-20 md:py-24">
          <div className="grid md:grid-cols-[1fr_2fr] gap-12 items-start">
            <p className="label-meta-lg text-bordo">{t('excerptsCloserKicker')}</p>
            <div className="space-y-6 max-w-xl">
              <h2 className="heading-section text-ink text-2xl md:text-3xl leading-snug">
                {t('excerptsCloserH2')}
              </h2>
              <p className="font-serif text-ink-soft leading-relaxed">{t('excerptsCloserBody')}</p>
              <div className="flex flex-wrap items-center gap-4 pt-2">
                <a
                  href={loc.amazonUrl}
                  target="_blank"
                  rel="noreferrer"
                  className="btn-gold"
                >
                  {t('excerptsCtaAmazon')}
                </a>
                <a
                  href={hotmartUrl}
                  target="_blank"
                  rel="noreferrer"
                  className="font-sans text-sm link-text text-ink-soft hover:text-bordo"
                >
                  {t('excerptsCtaHotmart')}
                </a>
              </div>
              <p className="font-sans text-sm pt-4">
                <Link href={{ pathname: '/livros/[slug]', params: { slug } }} className="link-text text-ink-muted">
                  {t('excerptsBackToBook')}
                </Link>
              </p>
            </div>
          </div>
        </div>
      </section>
    </article>
  );
}
