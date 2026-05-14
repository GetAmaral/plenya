import type { Metadata } from 'next';
import { notFound } from 'next/navigation';
import { getTranslations, setRequestLocale } from 'next-intl/server';
import { brand } from '@plenya/brand';
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

// Links institucionais correspondentes a cada pilar. Cria caminho navegacional
// (e sinal de PageRank interno) entre o cluster de artigos e a página-pilar
// canônica do método. Hardcoded inline porque é mapa estável e curto.
type DeepenLink = { href: string; key: 'method' | 'healthspan' | 'score' | 'checkup' | 'kidney' };
const DEEPEN_MAP: Record<Pillar, DeepenLink[]> = {
  'alimentacao-atividade-fisica': [
    { href: '/healthspan', key: 'healthspan' },
    { href: '/metodo-agir', key: 'method' },
  ],
  'gestao-metabolica': [
    { href: '/checkup-longevidade', key: 'checkup' },
    { href: '/avaliacao-renal-preventiva', key: 'kidney' },
    { href: '/metodo-agir', key: 'method' },
  ],
  'integracao-corpo-mente': [
    { href: '/metodo-agir', key: 'method' },
    { href: '/escore-plenya', key: 'score' },
  ],
  'ritmo-circadiano': [
    { href: '/metodo-agir', key: 'method' },
    { href: '/healthspan', key: 'healthspan' },
  ],
  longevidade: [
    { href: '/healthspan', key: 'healthspan' },
    { href: '/escore-plenya', key: 'score' },
    { href: '/metodo-agir', key: 'method' },
  ],
};

const DEEPEN_LABELS: Record<'pt' | 'en', Record<DeepenLink['key'], { name: string; line: string }>> = {
  pt: {
    method: { name: 'Método AGIR', line: 'Os quatro pilares que organizam o cuidado contínuo' },
    healthspan: { name: 'Healthspan', line: 'Comprimir a morbidade — a década entre o normal e o ótimo' },
    score: { name: 'Escore Plenya', line: 'Estratificação longitudinal com mais de 800 itens' },
    checkup: { name: 'Check-up de longevidade', line: 'O check-up que enxerga a década seguinte' },
    kidney: { name: 'Avaliação renal preventiva', line: 'Função renal antes da janela fechar' },
  },
  en: {
    method: { name: 'ACTS Method', line: 'The four pillars that structure continuous care' },
    healthspan: { name: 'Healthspan', line: 'Compressing morbidity — the decade between normal and optimal' },
    score: { name: 'Plenya Score', line: 'Longitudinal risk stratification with 800+ items' },
    checkup: { name: 'Longevity check-up', line: 'A check-up that sees the next decade' },
    kidney: { name: 'Preventive kidney assessment', line: 'Kidney function before the silent window closes' },
  },
};

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

  // CollectionPage + ItemList: sinal explícito pro Google de que esta é
  // uma página-cluster contendo N artigos relacionados pelo pilar.
  const localePrefix = locale === 'en' ? '/en' : '';
  const pageUrl = `${brand.url}${localePrefix}/blog/categoria/${pilar}`;
  const collectionSchema = {
    '@context': 'https://schema.org',
    '@type': 'CollectionPage',
    '@id': `${pageUrl}#collection`,
    url: pageUrl,
    name: labels[pilar],
    inLanguage: locale === 'en' ? 'en' : 'pt-BR',
    isPartOf: { '@type': 'WebSite', '@id': `${brand.url}#website` },
    about: { '@type': 'Thing', name: labels[pilar] },
    mainEntity: {
      '@type': 'ItemList',
      numberOfItems: posts.length,
      itemListElement: posts.map((post, i) => ({
        '@type': 'ListItem',
        position: i + 1,
        url: `${brand.url}${localePrefix}/blog/${post.slug}`,
        name: post.title,
      })),
    },
  };

  const deepen = DEEPEN_MAP[pilar];
  const deepenLabels = DEEPEN_LABELS[locale === 'en' ? 'en' : 'pt'];
  const deepenHeading = locale === 'en' ? 'Go deeper into the method' : 'Aprofunde no método';
  const deepenSub = locale === 'en'
    ? 'Where this pillar lives inside Plenya — clinical pages, score, assessments.'
    : 'Onde este pilar mora dentro da Plenya — páginas clínicas, escore, avaliações.';

  return (
    <>
      <script
        type="application/ld+json"
        dangerouslySetInnerHTML={{ __html: JSON.stringify(collectionSchema) }}
      />
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

          <div className="border-t border-petrol/15 pt-12">
            <p className="label-upper text-gold">{deepenHeading}</p>
            <p className="text-petrol/70 mt-2 max-w-2xl">{deepenSub}</p>
            <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-8 mt-8">
              {deepen.map((d) => (
                <Link
                  key={d.href}
                  href={d.href as Parameters<typeof Link>[0]['href']}
                  className="border-t border-petrol/15 pt-6 space-y-2 group"
                >
                  <p className="heading-section text-petrol text-xl group-hover:text-gold transition">
                    {deepenLabels[d.key].name}
                  </p>
                  <p className="text-petrol/70 text-sm leading-relaxed">{deepenLabels[d.key].line}</p>
                </Link>
              ))}
            </div>
          </div>
        </div>
      </section>
    </>
  );
}
