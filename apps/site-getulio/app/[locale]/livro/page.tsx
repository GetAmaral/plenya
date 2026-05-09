import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { EducationalNotice } from '@/components/legal/educational-notice';
import { BookSchema } from '@/components/seo/book-schema';

type Trecho = { cap: string; citacao: string };

export async function generateMetadata({
  params,
}: {
  params: Promise<{ locale: string }>;
}): Promise<Metadata> {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: 'book' });
  return {
    title: t('metaTitle'),
    description: t('metaDescription'),
    alternates: {
      canonical: locale === 'en' ? '/en/livro' : '/livro',
      languages: { 'pt-BR': '/livro', en: '/en/livro' },
    },
  };
}

export default async function LivroPage({
  params,
}: {
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  setRequestLocale(locale);
  const t = await getTranslations({ locale, namespace: 'book' });
  const trechos = t.raw('trechos') as Trecho[];
  const amazonUrl = locale === 'en' ? t('amazonUrlEn') : t('amazonUrl');
  const hotmartBase = t('hotmartUrl');
  const hotmartHero = `${hotmartBase}?src=site-hero`;
  const hotmartBuy = `${hotmartBase}?src=site-buy`;

  return (
    <article>
      <BookSchema
        title={t('metaTitle')}
        description={t('metaDescription')}
        isbn="978-65-02-06742-0"
        amazonUrl={amazonUrl}
        hotmartUrl={hotmartBase}
        coverUrl="https://drgetulioamaralfilho.com.br/images/livro-capa.jpg"
        locale={locale}
      />
      {/* Hero */}
      <section className="editorial-container pt-16 md:pt-24 pb-20">
        <div className="grid lg:grid-cols-[320px_1fr] gap-12 lg:gap-20 items-start">
          <div className="relative aspect-[2/3] w-full max-w-[320px] mx-auto lg:mx-0 shadow-2xl">
            <Image
              src="/images/livro-capa.jpg"
              alt={t('coverAlt')}
              fill
              priority
              className="object-cover"
              sizes="(min-width: 1024px) 320px, 80vw"
            />
          </div>

          <div className="space-y-8">
            <p className="label-meta">{t('kicker')}</p>
            <h1 className="heading-display text-[clamp(2.2rem,5vw,4rem)]">
              {t('h1Title')}
              <span className="block font-serif font-normal text-2xl md:text-3xl text-ink-muted mt-4 italic">
                {t('h1Subtitle')}
              </span>
            </h1>

            <div className="prose-body max-w-xl">
              <p>{t('lead1')}</p>
              <p>{t('lead2')}</p>
            </div>

            <div className="pt-2 space-y-3">
              <div>
                <a
                  href={amazonUrl}
                  target="_blank"
                  rel="noreferrer"
                  className="btn-gold"
                >
                  {t('heroBuyCta')}
                </a>
              </div>
              <p className="font-sans text-sm">
                <a
                  href={hotmartHero}
                  target="_blank"
                  rel="noreferrer"
                  className="link-text text-ink-muted hover:text-bordo"
                >
                  {t('heroBuyCtaHotmart')}
                </a>
              </p>
            </div>

            <div className="space-y-1 font-sans text-sm text-ink-muted">
              <p>{t('isbnLine')}</p>
              <p>{t('editionLine')}</p>
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
                <Link href="/sobre" className="link-text">{t('authorCta')}</Link>
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

      {/* Trechos selecionados */}
      <section className="border-t border-rule bg-paper">
        <div className="editorial-container py-20 md:py-28">
          <p className="label-meta mb-14">{t('trechosKicker')}</p>

          <div className="space-y-20 max-w-3xl">
            {trechos.map((tr, i) => (
              <figure key={i} className="space-y-5">
                <p className="label-meta text-bordo">{tr.cap}</p>
                <blockquote className="font-serif text-2xl md:text-3xl leading-snug text-ink italic border-l-2 border-bordo pl-6 md:pl-10">
                  {tr.citacao}
                </blockquote>
              </figure>
            ))}
          </div>

          <p className="font-sans text-sm pt-16 max-w-3xl">
            <Link href="/livro/excertos" className="link-text text-bordo">
              {t('trechosViewAll')}
            </Link>
          </p>
        </div>
      </section>

      <EducationalNotice />

      {/* Onde comprar — Amazon (impresso + Kindle) */}
      <section className="border-t border-rule bg-paper">
        <div className="editorial-container py-20 md:py-24">
          <div className="grid md:grid-cols-[1fr_2fr] gap-12 items-start">
            <p className="label-meta-lg text-bordo">{t('buyKicker')}</p>
            <div className="space-y-6 max-w-xl">
              <h2 className="heading-section text-ink text-2xl md:text-3xl leading-snug">
                {t('buyH2')}
              </h2>
              <p className="font-serif text-ink-soft leading-relaxed">
                {t('buyBody')}
              </p>
              <div className="flex flex-wrap items-center gap-4 pt-2">
                <a
                  href={amazonUrl}
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
                  className="font-sans text-sm link-text text-ink-soft hover:text-bordo"
                >
                  {t('buyCtaHotmart')}
                </a>
              </div>
            </div>
          </div>
        </div>
      </section>
    </article>
  );
}
