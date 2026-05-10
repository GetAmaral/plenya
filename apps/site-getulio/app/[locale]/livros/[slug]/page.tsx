import type { Metadata } from 'next';
import Image from 'next/image';
import { notFound } from 'next/navigation';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { EducationalNotice } from '@/components/legal/educational-notice';
import { BookSchema } from '@/components/seo/book-schema';
import { Breadcrumbs } from '@/components/layout/breadcrumbs';
import { BreadcrumbSchema } from '@/components/seo/breadcrumb-schema';
import { getAllBooks, getBook, localizedBook } from '@/lib/books';

const BASE = 'https://drgetulioamaralfilho.com.br';

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
  const fullTitle = `${loc.title}: ${loc.subtitle}`;
  const url = locale === 'en' ? `/en/books/${slug}` : `/livros/${slug}`;
  return {
    title: fullTitle,
    description: loc.shortDescription,
    alternates: {
      canonical: url,
      languages: { 'pt-BR': `/livros/${slug}`, pt: `/livros/${slug}`, en: `/en/books/${slug}` },
    },
    openGraph: {
      type: 'book',
      url,
      title: fullTitle,
      description: loc.shortDescription,
      images: [`${BASE}${book.cover}`],
    },
  };
}

export default async function BookDetailPage({
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
  const hotmartHero = `${book.hotmartUrl}?src=site-hero`;
  const hotmartBuy = `${book.hotmartUrl}?src=site-buy`;
  const fullTitle = `${loc.title}: ${loc.subtitle}`;
  const homeLabel = locale === 'en' ? 'Home' : 'Início';
  const booksLabel = locale === 'en' ? 'Books' : 'Livros';

  return (
    <article>
      <BookSchema
        title={fullTitle}
        description={loc.shortDescription}
        isbn={book.isbn}
        slug={book.slug}
        amazonUrl={loc.amazonUrl}
        hotmartUrl={book.hotmartUrl}
        coverUrl={`${BASE}${book.cover}`}
        locale={locale}
      />
      <BreadcrumbSchema
        items={[
          { name: homeLabel, url: '/' },
          { name: booksLabel, url: '/livros' },
          { name: loc.title },
        ]}
      />

      <header className="editorial-container pt-12 md:pt-16 pb-8">
        <Breadcrumbs
          items={[
            { label: homeLabel, href: '/' },
            { label: booksLabel, href: '/livros' },
            { label: loc.title },
          ]}
        />
      </header>

      {/* Hero */}
      <section className="editorial-container pb-20">
        <div className="grid lg:grid-cols-[320px_1fr] gap-12 lg:gap-20 items-start">
          <div className="relative aspect-[2/3] w-full max-w-[320px] mx-auto lg:mx-0 shadow-2xl">
            <Image
              src={book.cover}
              alt={loc.title}
              fill
              priority
              className="object-cover"
              sizes="(min-width: 1024px) 320px, 80vw"
            />
          </div>

          <div className="space-y-8">
            <p className="label-meta">{t('kickerPrefix')} · {book.year}</p>
            <h1 className="heading-display text-[clamp(2.2rem,5vw,4rem)]">
              {loc.title}
              <span className="block font-serif font-normal text-2xl md:text-3xl text-ink-muted mt-4 italic">
                {loc.subtitle}
              </span>
            </h1>

            <div className="prose-body max-w-xl">
              <p>{loc.lead1}</p>
              <p>{loc.lead2}</p>
            </div>

            <div className="pt-2 flex flex-wrap items-center gap-4">
              <a
                href={loc.amazonUrl}
                target="_blank"
                rel="noreferrer"
                className="btn-gold"
              >
                {t('heroBuyCta')}
              </a>
              <a
                href={hotmartHero}
                target="_blank"
                rel="noreferrer"
                className="btn-outline"
              >
                {t('heroBuyCtaHotmart')}
              </a>
            </div>

            <div className="space-y-1 font-sans text-sm text-ink-muted">
              <p>{t('isbnPrefix')} {book.isbn} · {book.year}</p>
              <p>{loc.edition}</p>
            </div>
          </div>
        </div>
      </section>

      {/* Autor */}
      <section className="border-t border-rule">
        <div className="editorial-container py-20 md:py-24">
          <div className="grid md:grid-cols-[1fr_320px] gap-12 lg:gap-20 items-center">
            <div className="space-y-5 max-w-xl order-2 md:order-1">
              <p className="label-meta-lg text-bordo">{t('authorKicker')}</p>
              <p className="font-serif text-lg md:text-xl text-ink-soft leading-relaxed">
                {t('authorBody')}
              </p>
              <p className="font-sans text-sm">
                <Link href="/o-medico" className="link-text">{t('authorCta')}</Link>
              </p>
            </div>
            <div className="relative aspect-[3/4] w-full max-w-[320px] mx-auto md:mx-0 order-1 md:order-2">
              <Image
                src="/images/getulio-autor.jpg"
                alt={t('authorPortraitAlt')}
                fill
                className="object-cover"
                sizes="(min-width: 768px) 320px, 70vw"
              />
            </div>
          </div>
        </div>
      </section>

      {/* Trechos */}
      <section className="border-t border-rule bg-paper">
        <div className="editorial-container py-20 md:py-28">
          <p className="label-meta mb-14">{t('trechosKicker')}</p>

          <div className="space-y-20 max-w-3xl">
            {loc.excerpts.map((tr, i) => (
              <figure key={i} className="space-y-5">
                <p className="label-meta text-bordo">{tr.cap}</p>
                <blockquote className="font-serif text-2xl md:text-3xl leading-snug text-ink italic border-l-2 border-bordo pl-6 md:pl-10">
                  {tr.citacao}
                </blockquote>
              </figure>
            ))}
          </div>

          <p className="font-sans text-sm pt-16 max-w-3xl">
            <Link href={{ pathname: '/livros/[slug]/excertos', params: { slug: book.slug } }} className="link-text text-bordo">
              {t('trechosViewAll')}
            </Link>
          </p>
        </div>
      </section>

      <EducationalNotice />

      {/* Onde comprar */}
      <section className="border-t border-rule bg-paper">
        <div className="editorial-container py-20 md:py-24">
          <div className="grid md:grid-cols-[1fr_2fr] gap-12 items-start">
            <p className="label-meta-lg text-bordo">{t('buyKicker')}</p>
            <div className="space-y-6 max-w-xl">
              <h2 className="heading-section text-ink text-2xl md:text-3xl leading-snug">
                {t('buyH2')}
              </h2>
              <p className="font-serif text-ink-soft leading-relaxed">
                {t('buyBody')} {loc.edition}.
              </p>
              <div className="flex flex-wrap items-center gap-4 pt-2">
                <a
                  href={loc.amazonUrl}
                  target="_blank"
                  rel="noreferrer"
                  className="btn-gold"
                >
                  {t('buyCta')}
                </a>
                <a
                  href={hotmartBuy}
                  target="_blank"
                  rel="noreferrer"
                  className="btn-outline"
                >
                  {t('buyCtaHotmart')}
                </a>
              </div>
              <p className="font-sans text-sm pt-4">
                <Link href="/livros" className="link-text text-ink-muted">
                  {t('backToList')}
                </Link>
              </p>
            </div>
          </div>
        </div>
      </section>
    </article>
  );
}
