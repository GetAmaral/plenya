import { useTranslations } from 'next-intl';
import { PlenyaInfinity } from '@plenya/brand/logo';
import { Link } from '@/lib/i18n/navigation';

/**
 * Estrutura section — substitui a duplicação do manifesto na home.
 * Usa material do vídeo institucional 05 + o conceito "normal vs ótimo"
 * (vídeo 04). Manifesto integral fica apenas na página /a-plenya.
 */
export function EstruturaSection() {
  const t = useTranslations('home');
  const tCta = useTranslations('cta');
  return (
    <section className="bg-petrol text-cream relative overflow-hidden">
      <div className="site-container section relative z-10">
        <div className="grid lg:grid-cols-[auto_1fr] gap-10 lg:gap-20 items-start max-w-5xl">
          <PlenyaInfinity
            aria-hidden="true"
            focusable="false"
            className="h-14 md:h-20 w-auto text-gold shrink-0"
          />

          <div className="space-y-8">
            <div className="space-y-4">
              <p className="label-upper text-gold">{t('estruturaLabel')}</p>
              <h2 className="heading-section text-cream text-3xl md:text-5xl">
                {t('estruturaTitlePart1')}<em className="not-italic text-gold">{t('estruturaTitleEm')}</em>
              </h2>
            </div>

            <p className="text-cream/80 text-lg leading-relaxed max-w-2xl">
              {t('estruturaP1')}
            </p>

            <p className="text-cream/70 text-lg leading-relaxed max-w-2xl">
              <strong className="text-cream">{t('estruturaP2Strong')}</strong>{t('estruturaP2Rest')}
            </p>

            <div className="pt-4">
              <Link href="/metodo-agir" className="btn-outline-light">
                {tCta('knowMethod')}
              </Link>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
