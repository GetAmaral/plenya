import { useTranslations } from 'next-intl';
import { Link } from '@/lib/i18n/navigation';

export default function NotFound() {
  const t = useTranslations('notFound');

  return (
    <article>
      <header className="editorial-container pt-24 md:pt-32 pb-12">
        <p className="label-meta mb-6">{t('kicker')}</p>
        <h1 className="heading-display text-[clamp(2.2rem,5vw,3.8rem)] max-w-3xl">
          {t('h1')}
        </h1>
        <p className="prose-body mt-8 max-w-2xl text-ink-soft">{t('lead')}</p>
      </header>

      <section className="editorial-container pb-32">
        <div className="flex flex-wrap items-center gap-4 pt-2">
          <Link href="/" className="btn-gold">
            {t('ctaHome')}
          </Link>
          <Link href="/escritos" className="btn-outline">
            {t('ctaWritings')}
          </Link>
          <Link href="/livros/antes" className="link-text font-sans text-sm">
            {t('ctaBook')} →
          </Link>
          <Link href="/contato" className="link-text font-sans text-sm">
            {t('ctaContact')} →
          </Link>
        </div>
      </section>
    </article>
  );
}
