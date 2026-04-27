import type { Metadata } from 'next';
import { notFound } from 'next/navigation';
import { setRequestLocale } from 'next-intl/server';
import { defaultLocale, isLocale, locales } from '@/lib/i18n/config';
import { Link } from '@/lib/i18n/navigation';
import { getPostsByPillar, pillarLabels, pillars, type Pillar } from '@/lib/blog';
import { PostCard } from '@/components/blog/post-card';
import { PillarFilter } from '@/components/blog/pillar-filter';

export function generateStaticParams() {
  return locales.flatMap((locale) => pillars.map((pilar) => ({ locale, pilar })));
}

function isPillar(value: string): value is Pillar {
  return (pillars as readonly string[]).includes(value);
}

export async function generateMetadata({ params }: { params: Promise<{ pilar: string }> }): Promise<Metadata> {
  const { pilar } = await params;
  if (!isPillar(pilar)) return {};
  return {
    title: `${pillarLabels[pilar]} — Blog Plenya`,
    description: `Artigos sobre ${pillarLabels[pilar].toLowerCase()} no Método AGIR.`,
    alternates: { canonical: `/blog/categoria/${pilar}` },
  };
}

export default async function BlogPillarPage({ params }: { params: Promise<{ locale: string; pilar: string }> }) {
  const { locale: rawLocale, pilar } = await params;
  const locale = isLocale(rawLocale) ? rawLocale : defaultLocale;
  if (!isPillar(pilar)) notFound();
  setRequestLocale(locale);

  const posts = await getPostsByPillar(locale, pilar);

  return (
    <>
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32">
          <Link href="/blog" className="label-upper text-cream/50 hover:text-gold transition">← Blog</Link>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,5rem)] text-cream mt-6">{pillarLabels[pilar]}</h1>
          <p className="text-cream/70 text-lg mt-4">
            {posts.length} artigo{posts.length === 1 ? '' : 's'} nesta categoria.
          </p>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-container section space-y-16">
          <PillarFilter active={pilar} />

          {posts.length ? (
            <div>
              {posts.map((post) => (
                <PostCard key={post.slug} post={post} />
              ))}
            </div>
          ) : (
            <p className="text-petrol/50 text-center label-upper py-16">Nenhum artigo nesta categoria ainda.</p>
          )}
        </div>
      </section>
    </>
  );
}
