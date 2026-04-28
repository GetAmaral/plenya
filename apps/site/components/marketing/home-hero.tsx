import { useTranslations } from 'next-intl';
import Image from 'next/image';
import { PlenyaInfinity } from '@plenya/brand/logo';
import { Link } from '@/lib/i18n/navigation';

export function HomeHero() {
  const t = useTranslations('home');
  const tCta = useTranslations('cta');

  return (
    <section className="relative min-h-[100svh] flex items-end overflow-hidden">
      {/* Photo background */}
      <Image
        src="/images/hero.jpg"
        alt=""
        fill
        priority
        fetchPriority="high"
        className="object-cover"
        sizes="100vw"
        quality={75}
      />
      {/* Overlay mais escuro embaixo pra ancorar o título contra a foto e
         garantir contraste do .heading-hero (line-height 0.95). */}
      <div className="absolute inset-0 bg-gradient-to-b from-petrol/30 via-petrol/40 to-petrol/85" />

      {/* Decorative infinity-P watermark — símbolo de continuidade do cuidado */}
      <PlenyaInfinity
        aria-hidden="true"
        focusable="false"
        className="hidden md:block absolute top-28 right-8 lg:right-20 w-56 lg:w-80 h-auto text-gold/20 pointer-events-none"
      />

      <div className="relative site-container pb-24 md:pb-32 pt-40 text-cream w-full">
        <div className="max-w-[820px]">
          <h1 className="heading-hero text-[clamp(2.8rem,7vw,5.5rem)]">
            Viver bem{' '}
            <em className="not-italic text-gold">é o reflexo</em>{' '}
            do que você escolhe todos os dias.
          </h1>

          <div className="mt-10 max-w-md space-y-0.5">
            <p className="heading-section text-cream/90 text-xl md:text-2xl">
              Plenitude não é um ponto de chegada.
            </p>
            <p className="heading-section text-cream/90 text-xl md:text-2xl">
              É uma linha contínua.
            </p>
          </div>

          <div className="mt-14">
            <Link href="/equipe" className="btn-gold">
              {tCta('knowTeam')}
            </Link>
          </div>

          <p className="mt-12 label-upper text-cream/60 tracking-[0.3em]">
            Plenya · Viva bem, viva mais.
          </p>
        </div>
      </div>
    </section>
  );
}
