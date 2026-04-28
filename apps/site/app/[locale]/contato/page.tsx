import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale } from 'next-intl/server';
import { ContactForm } from '@/components/marketing/contact-form';
import { brand } from '@plenya/brand';

export const metadata: Metadata = {
  title: 'Contato',
  description: 'Fale com a equipe Plenya — Consulta Plenya, Continuum Plenya ou orientação geral.',
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
        <div className="site-container section grid lg:grid-cols-[5fr_7fr] gap-12 lg:gap-20">
          <div className="space-y-10">
            <div className="space-y-4 max-w-md">
              <p className="label-upper text-gold">Por que falar com a Plenya</p>
              <p className="heading-section text-petrol text-2xl md:text-3xl">
                Escuta longa antes da conduta.
              </p>
              <p className="text-petrol/70 text-base leading-relaxed">
                Cada mensagem é lida pela equipe — e a primeira resposta orienta o
                próximo passo, seja Consulta avulsa, Continuum ou apenas uma
                dúvida clínica.
              </p>
            </div>
            <div className="grid sm:grid-cols-2 gap-y-8 gap-x-6 max-w-md">
              <div className="space-y-2">
                <p className="label-upper text-gold">Email</p>
                <a href={`mailto:${brand.email}`} className="text-petrol text-base hover:text-gold transition break-words">
                  {brand.email}
                </a>
              </div>
              <div className="space-y-2">
                <p className="label-upper text-gold">Instagram</p>
                <a href={brand.social.instagram} target="_blank" rel="noreferrer" className="text-petrol text-base hover:text-gold transition">
                  @plenyaSaude
                </a>
              </div>
              <div className="space-y-2">
                <p className="label-upper text-gold">Endereço</p>
                <p className="text-petrol text-base">Londrina, PR</p>
              </div>
              <div className="space-y-2">
                <p className="label-upper text-gold">Horário</p>
                <p className="text-petrol/75 text-base">Seg a sex · 8h–18h</p>
              </div>
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
