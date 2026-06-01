import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale, getTranslations } from 'next-intl/server';
import { ContactForm } from '@/components/marketing/contact-form';
import { brand } from '@plenya/brand';

type Params = Promise<{ locale: string }>;

export async function generateMetadata({ params }: { params: Params }): Promise<Metadata> {
  const { locale } = await params;
  const t = await getTranslations({ locale, namespace: 'contact' });
  return {
    title: t('metaTitle'),
    description: t('metaDescription'),
    alternates: {
      canonical: locale === 'en' ? '/en/contact' : '/contato',
      languages: {
        'pt-BR': '/contato', pt: '/contato',
        en: '/en/contact',
        'x-default': '/contato',
      },
    },
  };
}

export default async function ContactPage({ params }: { params: Params }) {
  const { locale } = await params;
  setRequestLocale(locale);
  const t = await getTranslations('contact');

  return (
    <>
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32">
          <p className="label-upper text-gold mb-6">{t('heroLabel')}</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,5rem)] text-cream">{t('heroTitle')}</h1>
          <p className="text-cream/70 text-lg mt-6 max-w-lg">{t('heroP')}</p>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-container section grid lg:grid-cols-[5fr_7fr] gap-12 lg:gap-20">
          <div className="space-y-10">
            <div className="space-y-4 max-w-md">
              <p className="label-upper text-gold">{t('whyLabel')}</p>
              <p className="heading-section text-petrol text-2xl md:text-3xl">{t('whyTitle')}</p>
              <p className="text-petrol/70 text-base leading-relaxed">{t('whyDesc')}</p>
            </div>
            <div className="grid sm:grid-cols-2 gap-y-8 gap-x-6 max-w-md">
              <div className="space-y-2">
                <p className="label-upper text-gold">{t('emailLabel')}</p>
                <a
                  href={`mailto:${brand.email}`}
                  className="text-petrol text-base hover:text-gold transition wrap-break-word"
                >
                  {brand.email}
                </a>
              </div>
              <div className="space-y-2">
                <p className="label-upper text-gold">{t('instagramLabel')}</p>
                <a
                  href={brand.social.instagram}
                  target="_blank"
                  rel="noreferrer"
                  className="text-petrol text-base hover:text-gold transition"
                >
                  @plenyaSaude
                </a>
              </div>
              <div className="space-y-2">
                <p className="label-upper text-gold">{t('phoneLabel')}</p>
                <a
                  href={brand.whatsappUrl}
                  target="_blank"
                  rel="noreferrer"
                  className="text-petrol text-base hover:text-gold transition"
                >
                  {brand.phone}
                </a>
              </div>
              <div className="space-y-2">
                <p className="label-upper text-gold">{t('addressLabel')}</p>
                <address className="text-petrol text-base not-italic leading-relaxed">
                  {brand.address.street}<br />
                  {brand.address.complement}<br />
                  {brand.address.neighborhood} · {brand.address.city}/{brand.address.state}<br />
                  {brand.address.postalCode}
                </address>
              </div>
              <div className="space-y-2">
                <p className="label-upper text-gold">{t('hoursLabel')}</p>
                <p className="text-petrol/75 text-base">{t('hoursValue')}</p>
              </div>
            </div>
          </div>
          <ContactForm />
        </div>
      </section>

      <section className="bg-cream-100">
        <div className="site-container pb-24 md:pb-32">
          <figure className="space-y-4">
            <div className="relative aspect-3/2 overflow-hidden">
              <Image
                src="/images/clinic-exterior.jpg"
                alt={t('clinicCaptionValue')}
                fill
                className="object-cover"
                sizes="(min-width: 1024px) 1120px, 100vw"
              />
            </div>
            <figcaption className="flex items-baseline gap-3 text-petrol/60 text-sm">
              <span className="label-upper text-gold">{t('clinicCaptionLabel')}</span>
              <span>{t('clinicCaptionValue')}</span>
            </figcaption>
          </figure>
        </div>
      </section>
    </>
  );
}
