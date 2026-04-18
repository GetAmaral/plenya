import Image from 'next/image';
import { useLocale, useTranslations } from 'next-intl';
import { agirPillars } from '@plenya/brand';
import { Link } from '@/lib/i18n/navigation';

export function AgirPillarsSection() {
  const t = useTranslations('home');
  const locale = useLocale();

  return (
    <section className="relative bg-petrol text-cream overflow-hidden">
      {/* Brand estampa — decorative pattern fades into the right edge */}
      <Image
        src="/brand/pattern/p-block.png"
        alt=""
        aria-hidden="true"
        width={1400}
        height={990}
        className="hidden lg:block absolute -right-40 -top-20 w-[55%] max-w-[760px] opacity-[0.18] pointer-events-none select-none"
      />

      <div className="relative site-container section">
        <p className="label-upper text-gold mb-6">{t('agirTitle')}</p>
        <h2 className="heading-section text-cream text-3xl md:text-5xl max-w-2xl mb-6">
          {t('agirSubtitle')}
        </h2>

        <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-px bg-cream/10 border-y border-cream/15 mt-16">
          {agirPillars.map((pillar) => {
            const name = locale === 'en' ? pillar.nameEn : locale === 'es' ? pillar.nameEs : pillar.name;
            return (
              <div key={pillar.code} className="bg-petrol p-8 space-y-4">
                <span className="heading-section text-gold text-5xl block leading-none">{pillar.code}</span>
                <h3 className="heading-section text-cream text-xl">{name}</h3>
                <p className="text-cream/70 text-sm leading-relaxed">{pillar.idea}</p>
              </div>
            );
          })}
        </div>

        <div className="mt-12">
          <Link href="/metodo-agir" className="btn-gold">Conhecer o Método AGIR</Link>
        </div>
      </div>
    </section>
  );
}
