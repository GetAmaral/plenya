import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale } from 'next-intl/server';
import { ContactForm } from '@/components/marketing/contact-form';
import { brand } from '@plenya/brand';

export const metadata: Metadata = {
  title: 'Contato',
  description: 'Fale com a equipe Plenya — consulta avulsa, Continuum Plenya ou orientação geral.',
};

export default async function ContactPage({ params }: { params: Promise<{ locale: string }> }) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <>
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32">
          <p className="label-upper text-gold mb-6">Atendimento</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,5rem)] text-cream">Vamos conversar.</h1>
          <p className="text-cream/70 text-lg mt-6 max-w-lg">
            Resposta em até 2 horas em dias úteis (8h-18h). Atendemos de forma particular, presencial e online.
          </p>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-container section grid lg:grid-cols-2 gap-16">
          <div className="space-y-8">
            <div className="space-y-2">
              <p className="label-upper text-gold">Email</p>
              <a href={`mailto:${brand.email}`} className="text-petrol text-lg hover:text-gold transition">
                {brand.email}
              </a>
            </div>
            <div className="space-y-2">
              <p className="label-upper text-gold">Instagram</p>
              <a href={brand.social.instagram} target="_blank" rel="noreferrer" className="text-petrol text-lg hover:text-gold transition">
                @plenyaSaude
              </a>
            </div>
            <div className="space-y-2">
              <p className="label-upper text-gold">Endereço</p>
              <p className="text-petrol text-lg">Londrina, PR</p>
            </div>
            <div className="space-y-2">
              <p className="label-upper text-gold">Horário</p>
              <p className="text-petrol/70">Segunda a sexta, 8h às 18h</p>
            </div>
          </div>
          <ContactForm />
        </div>
      </section>

      <section className="bg-cream-100">
        <div className="site-container pb-24 md:pb-32">
          <figure className="space-y-4">
            <div className="relative aspect-[3/2] overflow-hidden">
              <Image
                src="/images/clinic-exterior.jpg"
                alt="Projeto arquitetônico da futura sede Plenya em Londrina"
                fill
                className="object-cover"
                sizes="(min-width: 1024px) 1120px, 100vw"
              />
            </div>
            <figcaption className="flex items-baseline gap-3 text-petrol/60 text-sm">
              <span className="label-upper text-gold">Projeto</span>
              <span>Render arquitetônico da futura sede Plenya em Londrina.</span>
            </figcaption>
          </figure>
        </div>
      </section>
    </>
  );
}
