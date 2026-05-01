import { useTranslations, useLocale } from 'next-intl';
import { brand } from '@plenya/brand';

/**
 * Banner mostrado APENAS no locale `en` no topo das páginas legais
 * (privacidade, termos, lgpd*). Esclarece que a versão portuguesa é
 * a autoridade legal — o site opera sob LGPD/CFM brasileiros.
 */
export function LegalEnNotice() {
  const locale = useLocale();
  const t = useTranslations('legalNotice');
  if (locale !== 'en') return null;

  return (
    <section className="bg-paper border-y border-petrol/15">
      <div className="site-narrow py-6 md:py-8 space-y-2">
        <p className="label-upper text-gold text-[10px]">English Summary</p>
        <p className="heading-section text-petrol text-lg md:text-xl">{t('title')}</p>
        <p className="text-petrol/75 text-sm leading-relaxed max-w-3xl">
          {t('body')}{' '}
          <a
            href={`mailto:${brand.email}`}
            className="text-gold underline underline-offset-4"
          >
            {brand.email}
          </a>
          .
        </p>
      </div>
    </section>
  );
}
