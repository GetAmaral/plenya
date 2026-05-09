import { useTranslations } from 'next-intl';

/**
 * Disclaimer educativo reutilizável para páginas de alta exposição editorial
 * (livro, palestras, escritos). Reforça que o conteúdo do site é educativo e
 * não constitui prescrição — defensividade adicional ao disclaimer do rodapé.
 *
 * Uso: <EducationalNotice /> ao final do conteúdo de uma página, antes do CTA.
 */
export function EducationalNotice() {
  const t = useTranslations('legal');
  return (
    <section className="border-t border-rule bg-paper">
      <div className="editorial-container py-10 md:py-12">
        <div className="grid lg:grid-cols-[200px_1fr] gap-6 lg:gap-12 items-start max-w-3xl">
          <p className="label-meta text-ink-muted">{t('educationalNoticeLabel')}</p>
          <p className="font-serif text-ink-soft text-sm leading-relaxed">
            {t('educationalNoticeText')}
          </p>
        </div>
      </div>
    </section>
  );
}
