import type { Metadata } from 'next';
import { getTranslations, setRequestLocale } from 'next-intl/server';
import { defaultLocale, isLocale, type Locale } from '@/lib/i18n/config';
import { getAllPosts, getFeaturedPost } from '@/lib/blog';
import { PostCard } from '@/components/blog/post-card';
import { PillarFilter } from '@/components/blog/pillar-filter';

export async function generateMetadata({ params }: { params: Promise<{ locale: string }> }): Promise<Metadata> {
  const { locale: rawLocale } = await params;
  const locale: Locale = isLocale(rawLocale) ? rawLocale : defaultLocale;
  const t = await getTranslations({ locale, namespace: 'blogIndex' });
  return { title: t('metaTitle'), description: t('metaDescription') };
}

export default async function BlogIndex({ params }: { params: Promise<{ locale: string }> }) {
  const { locale: rawLocale } = await params;
  const locale: Locale = isLocale(rawLocale) ? rawLocale : defaultLocale;
  setRequestLocale(locale);
  const t = await getTranslations({ locale, namespace: 'blogIndex' });

  const [featured, all] = await Promise.all([getFeaturedPost(locale), getAllPosts(locale)]);
  const rest = all.filter((p) => p.slug !== featured?.slug);

  return (
    <>
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32">
          <p className="label-upper text-gold mb-6">{t('label')}</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,5rem)] text-cream">{t('title')}</h1>
          <p className="text-cream/70 text-lg mt-6 max-w-lg">
            {t('subtitle')}
          </p>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-container section space-y-16">
          <PillarFilter locale={locale} />

          {featured && (
            <div>
              <p className="label-upper text-gold mb-6">{t('featured')}</p>
              <PostCard post={featured} featured locale={locale} />
            </div>
          )}

          {rest.length > 0 && (
            <div>
              <p className="label-upper text-gold mb-8">{t('mostRecent')}</p>
              <div className="grid sm:grid-cols-2 lg:grid-cols-3 gap-x-8 gap-y-12">
                {rest.map((post) => (
                  <PostCard key={post.slug} post={post} locale={locale} />
                ))}
              </div>
            </div>
          )}

          {!featured && !rest.length && (
            <p className="text-petrol/50 text-center label-upper py-16">{t('comingSoon')}</p>
          )}
        </div>
      </section>
    </>
  );
}
