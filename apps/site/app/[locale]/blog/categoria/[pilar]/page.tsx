import type { Metadata } from 'next';
import { notFound } from 'next/navigation';
import { getTranslations, setRequestLocale } from 'next-intl/server';
import { defaultLocale, isLocale, locales, type Locale } from '@/lib/i18n/config';
import { Link } from '@/lib/i18n/navigation';
import { getPostsByPillar, getPillarLabels, pillars, type Pillar } from '@/lib/blog';
import { PostCard } from '@/components/blog/post-card';
import { PillarFilter } from '@/components/blog/pillar-filter';

export function generateStaticParams() {
  return locales.flatMap((locale) => pillars.map((pilar) => ({ locale, pilar })));
}

function isPillar(value: string): value is Pillar {
  return (pillars as readonly string[]).includes(value);
}

export async function generateMetadata({ params }: { params: Promise<{ locale: string; pilar: string }> }): Promise<Metadata> {
  const { locale: rawLocale, pilar } = await params;
  if (!isPillar(pilar)) return {};
  const locale: Locale = isLocale(rawLocale) ? rawLocale : defaultLocale;
  const labels = getPillarLabels(locale);
  const t = await getTranslations({ locale, namespace: 'blogIndex' });
  return {
    title: `${labels[pilar]} — ${t('metaTitle')}`,
    description: `${t('categoryMetaPrefix')} ${labels[pilar].toLowerCase()} ${t('categoryMetaSuffix')}`,
    alternates: { canonical: `/blog/categoria/${pilar}` },
  };
}

export default async function BlogPillarPage({ params }: { params: Promise<{ locale: string; pilar: string }> }) {
  const { locale: rawLocale, pilar } = await params;
  const locale: Locale = isLocale(rawLocale) ? rawLocale : defaultLocale;
  if (!isPillar(pilar)) notFound();
  setRequestLocale(locale);

  const posts = await getPostsByPillar(locale, pilar);
  const labels = getPillarLabels(locale);
  const t = await getTranslations({ locale, namespace: 'blogIndex' });

  return (
    <>
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32">
          <Link href="/blog" className="label-upper text-cream/50 hover:text-gold transition">← Blog</Link>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,5rem)] text-cream mt-6">{labels[pilar]}</h1>
          <p className="text-cream/70 text-lg mt-4">
            {posts.length} {posts.length === 1 ? t('categoryArticleSingular') : t('categoryArticlePlural')}
          </p>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-container section space-y-16">
          <PillarFilter active={pilar} locale={locale} />

          {posts.length ? (
            <div className="grid sm:grid-cols-2 lg:grid-cols-3 gap-x-8 gap-y-16 md:gap-y-24">
              {posts.map((post) => (
                <PostCard key={post.slug} post={post} locale={locale} />
              ))}
            </div>
          ) : (
            <p className="text-petrol/50 text-center label-upper py-16">{t('categoryEmpty')}</p>
          )}
        </div>
      </section>
    </>
  );
}
