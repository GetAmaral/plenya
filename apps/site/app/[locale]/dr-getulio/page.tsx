import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';

type Params = Promise<{ locale: string }>;

export async function generateMetadata({ params }: { params: Params }): Promise<Metadata> {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: 'drGetulio' });
  return {
    title: t('metaTitle'),
    description: t('metaDescription'),
    alternates: {
      canonical: locale === 'en' ? '/en/dr-getulio' : '/dr-getulio',
      languages: {
        'pt-BR': '/dr-getulio',
        en: '/en/dr-getulio',
        'x-default': '/dr-getulio',
      },
    },
  };
}

export default async function DrGetulioPage({ params }: { params: Params }) {
  const { locale } = await params;
  setRequestLocale(locale);
  const t = await getTranslations('drGetulio');
  const tCta = await getTranslations('cta');

  const formacao = [1, 2, 3, 4, 5].map((n) => ({
    year: t(`edu${n}Year` as 'edu1Year'),
    label: t(`edu${n}Label` as 'edu1Label'),
  }));
  const atuacao = [1, 2, 3, 4, 5, 6].map((n) => t(`activity${n}` as 'activity1'));
  const memberships = [1, 2, 3].map((n) => t(`society${n}` as 'society1'));

  return (
    <>
      {/* Hero */}
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32 grid lg:grid-cols-2 gap-16 items-center">
          <div className="space-y-6">
            <p className="label-upper text-gold">{t('heroLabel')}</p>
            <h1 className="heading-hero text-[clamp(2.5rem,6vw,4.5rem)] text-cream">
              {t('heroNameLine1')}
              <br />
              {t('heroNameLine2')}
            </h1>
            <p className="label-upper text-cream/60">{t('heroCredentials')}</p>
            <p className="text-cream/85 text-lg leading-relaxed max-w-lg">{t('heroP1')}</p>
            <p className="text-cream/70 leading-relaxed max-w-lg">{t('heroP2')}</p>
            <div className="pt-4">
              <Link href="/contato" className="btn-gold">
                {tCta('scheduleConsultation')}
              </Link>
            </div>
            <div className="flex flex-wrap gap-x-8 gap-y-2 pt-6 text-sm">
              <a
                href="https://instagram.com/drGetulioAmaralFilho"
                target="_blank"
                rel="noreferrer"
                className="text-cream/55 hover:text-cream transition underline-offset-4 hover:underline decoration-gold/40"
              >
                @drGetulioAmaralFilho
              </a>
              <a
                href="https://drGetulioAmaralFilho.com.br"
                target="_blank"
                rel="noreferrer"
                className="text-cream/55 hover:text-cream transition underline-offset-4 hover:underline decoration-gold/40"
              >
                drGetulioAmaralFilho.com.br
              </a>
            </div>
          </div>
          <div className="relative aspect-[3/4] overflow-hidden">
            <Image
              src="/images/dr-getulio.jpg"
              alt={`${t('heroNameLine1')} ${t('heroNameLine2')}`}
              fill
              priority
              className="object-cover"
              sizes="(min-width: 1024px) 540px, 100vw"
            />
          </div>
        </div>
      </section>

      {/* Sobre a prática */}
      <section className="bg-cream">
        <div className="site-container section grid lg:grid-cols-[2fr_1fr] gap-16">
          <div className="space-y-6 text-petrol/85 text-lg leading-relaxed">
            <p className="label-upper text-gold">{t('aboutLabel')}</p>
            <p>{t('aboutP1')}</p>
            <p>{t('aboutP2')}</p>
            <p>{t('aboutP3')}</p>
          </div>

          <aside className="space-y-8">
            <div className="space-y-4">
              <p className="label-upper text-gold">{t('activityLabel')}</p>
              <ul className="space-y-3 text-petrol/80 text-sm">
                {atuacao.map((item) => (
                  <li key={item} className="flex gap-3">
                    <span className="text-gold mt-0.5">—</span>
                    <span>{item}</span>
                  </li>
                ))}
              </ul>
            </div>

            <div className="space-y-4">
              <p className="label-upper text-gold">{t('societiesLabel')}</p>
              <ul className="space-y-3 text-petrol/80 text-sm">
                {memberships.map((m) => (
                  <li key={m} className="flex gap-3">
                    <span className="text-gold mt-0.5">—</span>
                    <span>{m}</span>
                  </li>
                ))}
              </ul>
            </div>
          </aside>
        </div>
      </section>

      {/* Formação acadêmica */}
      <section className="bg-paper">
        <div className="site-container section">
          <p className="label-upper text-gold mb-10">{t('educationLabel')}</p>

          <div className="border-t border-petrol/15">
            {formacao.map((item) => (
              <div
                key={`${item.year}-${item.label}`}
                className="grid grid-cols-[80px_1fr] md:grid-cols-[120px_1fr] gap-6 py-6 border-b border-petrol/10"
              >
                <span className="label-upper text-gold pt-1">{item.year}</span>
                <p className="text-petrol/80">{item.label}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Galeria editorial — prática */}
      <section className="bg-cream">
        <div className="site-container section">
          <p className="label-upper text-gold mb-10">{t('practiceLabel')}</p>
          <div className="grid md:grid-cols-2 gap-6 md:gap-8">
            <div className="relative aspect-[3/4] overflow-hidden">
              <Image
                src="/images/getulio-estetoscopio.jpg"
                alt={t('photo1Alt')}
                fill
                className="object-cover"
                sizes="(min-width: 768px) 540px, 100vw"
              />
            </div>
            <div className="relative aspect-[3/4] overflow-hidden">
              <Image
                src="/images/getulio-consulta.jpg"
                alt={t('photo2Alt')}
                fill
                className="object-cover"
                sizes="(min-width: 768px) 540px, 100vw"
              />
            </div>
          </div>
        </div>
      </section>

      {/* Publicação */}
      <section className="bg-paper">
        <div className="site-container section">
          <div className="grid md:grid-cols-[180px_1fr] gap-10 md:gap-14 items-center max-w-3xl">
            <a
              href={locale === 'en' ? 'https://drgetulioamaralfilho.com.br/en/livro' : 'https://drgetulioamaralfilho.com.br/livro'}
              target="_blank"
              rel="noreferrer"
              className="relative aspect-[2/3] w-full max-w-[180px] mx-auto md:mx-0 shadow-xl block"
            >
              <Image
                src="/images/livro-capa.jpg"
                alt={t('publicationCoverAlt')}
                fill
                className="object-cover"
                sizes="180px"
              />
            </a>
            <div className="space-y-4">
              <p className="label-upper text-gold">{t('publicationLabel')}</p>
              <p className="font-serif text-xl md:text-2xl text-petrol leading-snug">
                {t('publicationTitle')}
              </p>
              <p className="text-petrol/75 leading-relaxed">{t('publicationLine')}</p>
              <p>
                <a
                  href={locale === 'en' ? 'https://drgetulioamaralfilho.com.br/en/livro' : 'https://drgetulioamaralfilho.com.br/livro'}
                  target="_blank"
                  rel="noreferrer"
                  className="font-sans text-sm text-gold hover:text-petrol transition underline-offset-4 hover:underline"
                >
                  {t('publicationCta')}
                </a>
              </p>
            </div>
          </div>
        </div>
      </section>

      {/* Cross-links */}
      <section className="bg-cream">
        <div className="site-container section grid md:grid-cols-3 gap-8">
          <Link href="/metodo-agir" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">{t('crossMethodLabel')}</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">
              {t('crossMethodTitle')}
            </p>
          </Link>
          <Link href="/equipe" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">{t('crossTeamLabel')}</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">
              {t('crossTeamTitle')}
            </p>
          </Link>
          <Link href="/continuum" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">{t('crossContinuumLabel')}</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">
              {t('crossContinuumTitle')}
            </p>
          </Link>
        </div>
      </section>
    </>
  );
}
