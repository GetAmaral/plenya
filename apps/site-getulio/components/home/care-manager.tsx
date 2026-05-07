import { useTranslations } from 'next-intl';
import { Link } from '@/lib/i18n/navigation';

/**
 * Seção dedicada do conceito de médico-gestor na home.
 * Espelha o que existe no site Plenya — concept-block prominente, não
 * só um pilar entre quatro. Vai entre Hero e Pillars.
 */
export function CareManager() {
  const t = useTranslations('home.manager');
  return (
    <section className="border-t border-rule bg-paper">
      <div className="editorial-container py-20 md:py-28">
        <div className="grid lg:grid-cols-[280px_1fr] gap-10 lg:gap-20 items-start">
          <div className="space-y-3">
            <p className="label-meta-lg text-bordo">{t('kicker')}</p>
          </div>

          <div className="space-y-8 max-w-2xl">
            <h2 className="heading-section text-3xl md:text-4xl lg:text-5xl text-ink leading-snug">
              {t('h2')}
            </h2>

            <p className="font-serif text-lg md:text-xl text-ink-soft leading-relaxed">
              {t('lead')}
            </p>

            {/* Glossário — três conceitos do tripé, marcados sutilmente */}
            <dl className="border-t border-rule pt-6 grid sm:grid-cols-3 gap-6">
              <Footnote
                term={t('footnote1')}
                body={t('footnote1Body')}
              />
              <Footnote
                term={t('footnote2')}
                body={t('footnote2Body')}
              />
              <Footnote
                term={t('footnote3')}
                body={t('footnote3Body')}
              />
            </dl>

            <p className="font-sans text-sm pt-2">
              <Link href="/sobre" className="link-text">
                {t('cta')}
              </Link>
            </p>
          </div>
        </div>
      </div>
    </section>
  );
}

function Footnote({ term, body }: { term: string; body: string }) {
  return (
    <div className="space-y-1">
      <dt className="label-meta text-bordo">{term}</dt>
      <dd className="font-serif text-ink-soft text-sm leading-snug">{body}</dd>
    </div>
  );
}
