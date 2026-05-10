import type { Metadata } from 'next';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { ContactForm } from '@/components/contact/contact-form';
import { BreadcrumbSchema } from '@/components/seo/breadcrumb-schema';

export async function generateMetadata({
  params,
}: {
  params: Promise<{ locale: string }>;
}): Promise<Metadata> {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: 'contato' });
  return {
    title: t('metaTitle'),
    description: t('metaDescription'),
    alternates: {
      canonical: locale === 'en' ? '/en/contact' : '/contato',
      languages: { 'pt-BR': '/contato', pt: '/contato', en: '/en/contact' },
    },
  };
}

export default async function ContatoPage({
  params,
}: {
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  setRequestLocale(locale);
  const t = await getTranslations({ locale, namespace: 'contato' });

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

      <section className="editorial-container pb-24 grid lg:grid-cols-[2fr_1fr] gap-16">
        <ContactForm />

        <aside className="space-y-10 lg:border-l lg:border-rule lg:pl-12">
          <div>
            <p className="label-meta mb-4">{t('directKicker')}</p>
            <ul className="space-y-3 font-serif text-ink-soft text-sm">
              <li>
                <p className="text-ink-muted text-xs mb-1">{t('labelGeneral')}</p>
                <a href="mailto:contato@drgetulioamaralfilho.com.br" className="link-text">contato@drgetulioamaralfilho.com.br</a>
              </li>
              <li>
                <p className="text-ink-muted text-xs mb-1">{t('labelPalestras')}</p>
                <a href="mailto:palestras@drgetulioamaralfilho.com.br" className="link-text">palestras@drgetulioamaralfilho.com.br</a>
              </li>
              <li>
                <p className="text-ink-muted text-xs mb-1">{t('labelImprensa')}</p>
                <a href="mailto:imprensa@drgetulioamaralfilho.com.br" className="link-text">imprensa@drgetulioamaralfilho.com.br</a>
              </li>
            </ul>
          </div>

          <div>
            <p className="label-meta mb-4">{t('clinicalKicker')}</p>
            <ul className="space-y-3 font-serif text-ink-soft text-sm">
              <li>
                <p className="text-ink-muted text-xs mb-1">{t('labelPlenya')}</p>
                <a href="https://plenyasaude.com.br" target="_blank" rel="noreferrer" className="link-text">plenyasaude.com.br</a>
              </li>
              <li>
                <p className="text-ink-muted text-xs mb-1">{t('labelNefroclinica')}</p>
                <a href="https://nefroclinica.com" target="_blank" rel="noreferrer" className="link-text">nefroclinica.com</a>
              </li>
            </ul>
          </div>

          <div>
            <p className="label-meta mb-4">{t('localKicker')}</p>
            <p className="font-serif text-ink-soft text-sm">{t('localValue')}</p>
          </div>
        </aside>
      </section>
    </article>
  );
}
