import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { getAllBooks, localizedBook } from '@/lib/books';
import { BreadcrumbSchema } from '@/components/seo/breadcrumb-schema';

export async function generateMetadata({
  params,
}: {
  params: Promise<{ locale: string }>;
}): Promise<Metadata> {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: 'livros' });
  return {
    title: t('metaTitle'),
    description: t('metaDescription'),
    alternates: {
      canonical: locale === 'en' ? '/en/books' : '/livros',
      languages: { 'pt-BR': '/livros', en: '/en/books' },
    },
  };
}

export default async function LivrosPage({
  params,
}: {
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  setRequestLocale(locale);
  const t = await getTranslations({ locale, namespace: 'livros' });
  const books = await getAllBooks();

  const homeLabel = locale === 'en' ? 'Home' : 'Início';

  return (
    <article>
      <BreadcrumbSchema
        items={[{ name: homeLabel, url: '/' }, { name: t('h1') }]}
      />
      <header className="editorial-container pt-16 md:pt-24 pb-12">
        <p className="label-meta mb-6">{t('kicker')}</p>
        <h1 className="heading-display text-[clamp(2.2rem,5vw,3.8rem)] max-w-3xl">
          {t('h1')}
        </h1>
        <p className="prose-body mt-8 max-w-2xl">{t('lead')}</p>
      </header>

      <section className="editorial-container pb-24">
        <ul className="border-t border-rule">
          {books.map((b) => {
            const loc = localizedBook(b, locale);
            return (
              <li key={b.slug} className="border-b border-rule">
                <Link
                  href={{ pathname: '/livros/[slug]', params: { slug: b.slug } }}
                  className="block py-12 md:py-16 grid md:grid-cols-[200px_1fr] gap-8 md:gap-12 group items-start"
                >
                  <div className="relative aspect-[2/3] w-full max-w-[200px] mx-auto md:mx-0 shadow-xl">
                    <Image
                      src={b.cover}
                      alt={loc.title}
                      fill
                      className="object-cover"
                      sizes="200px"
                    />
                  </div>
                  <div className="space-y-4">
                    <div className="flex items-center gap-4 flex-wrap">
                      <span className="label-meta text-bordo">
                        {t('yearLabel')} · {b.year}
                      </span>
                      <span className="label-meta text-ink-muted">
                        {t('isbnLabel')} {b.isbn}
                      </span>
                    </div>
                    <h2 className="heading-section text-2xl md:text-3xl group-hover:text-bordo transition-colors">
                      {loc.title}
                    </h2>
                    <p className="font-serif italic text-ink-muted text-lg md:text-xl leading-snug">
                      {loc.subtitle}
                    </p>
                    <p className="font-serif text-ink-soft leading-relaxed max-w-2xl">
                      {loc.shortDescription}
                    </p>
                    <p className="label-meta text-bordo pt-2">{t('viewBookCta')}</p>
                  </div>
                </Link>
              </li>
            );
          })}
        </ul>
      </section>
    </article>
  );
}
