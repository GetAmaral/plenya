import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';
import { isLocale, defaultLocale } from '@/lib/i18n/config';
import { getAllPosts, getFeaturedPost } from '@/lib/blog';
import { PostCard } from '@/components/blog/post-card';
import { PillarFilter } from '@/components/blog/pillar-filter';

export const metadata: Metadata = {
  title: 'Blog Plenya',
  description: 'Conteúdo educacional sobre saúde, performance e longevidade pelos médicos da Plenya.',
};

export default async function BlogIndex({ params }: { params: Promise<{ locale: string }> }) {
  const { locale: rawLocale } = await params;
  const locale = isLocale(rawLocale) ? rawLocale : defaultLocale;
  setRequestLocale(locale);

  const [featured, all] = await Promise.all([getFeaturedPost(locale), getAllPosts(locale)]);
  const rest = all.filter((p) => p.slug !== featured?.slug);

  return (
    <>
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32">
          <p className="label-upper text-gold mb-6">Boletim Plenya</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,5rem)] text-cream">Conteúdo</h1>
          <p className="text-cream/70 text-lg mt-6 max-w-lg">
            Artigos semanais escritos pelos médicos Plenya sobre os pilares do Método AGIR e longevidade.
          </p>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-container section space-y-16">
          <PillarFilter />

          {featured && (
            <div>
              <p className="label-upper text-gold mb-6">Destaque</p>
              <PostCard post={featured} featured />
            </div>
          )}

          {rest.length > 0 && (
            <div>
              <p className="label-upper text-gold mb-8">Mais recentes</p>
              <div className="grid sm:grid-cols-2 lg:grid-cols-3 gap-x-8 gap-y-12">
                {rest.map((post) => (
                  <PostCard key={post.slug} post={post} />
                ))}
              </div>
            </div>
          )}

          {!featured && !rest.length && (
            <p className="text-petrol/50 text-center label-upper py-16">Em breve, novos artigos.</p>
          )}
        </div>
      </section>
    </>
  );
}
