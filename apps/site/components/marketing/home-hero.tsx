import { useTranslations } from 'next-intl';
import Image from 'next/image';
import { PlenyaSymbol } from '@plenya/brand/logo/PlenyaSymbol';
import { Link } from '@/lib/i18n/navigation';

export function HomeHero() {
  const t = useTranslations('home');
  const tCta = useTranslations('cta');

  return (
    <section className="relative min-h-[100svh] flex items-center overflow-hidden">
      {/* Photo background */}
      <Image
        src="/images/hero.jpg"
        alt=""
        fill
        priority
        className="object-cover"
        sizes="100vw"
      />
      {/* Warm overlay for text legibility */}
      <div className="absolute inset-0 bg-gradient-to-b from-petrol/50 via-petrol/20 to-petrol/60" />

      {/* Decorative P monogram — top-right watermark */}
      <PlenyaSymbol
        aria-hidden="true"
        focusable="false"
        className="hidden md:block absolute top-32 right-12 lg:right-24 w-40 lg:w-56 h-40 lg:h-56 text-gold/20 pointer-events-none"
      />

      <div className="relative site-container py-40 text-cream">
        <h1 className="heading-hero text-[clamp(2.8rem,7vw,5.5rem)] max-w-[14ch]">
          Viver bem{' '}
          <em className="not-italic text-gold">é o reflexo</em>{' '}
          do que você escolhe todos os dias.
        </h1>

        <div className="mt-10 max-w-md space-y-1">
          <p className="heading-section text-cream/90 text-xl">
            Plenitude não é um ponto de chegada.
          </p>
          <p className="heading-section text-cream/90 text-xl">
            É uma linha contínua.
          </p>
        </div>

        <div className="mt-14">
          <Link href="/equipe" className="btn-gold">
            {tCta('knowTeam')}
          </Link>
        </div>
      </div>
    </section>
  );
}
