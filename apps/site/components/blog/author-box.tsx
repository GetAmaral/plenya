import Image from 'next/image';
import { getTranslations } from 'next-intl/server';
import type { Author } from '@/lib/authors';
import { defaultLocale, type Locale } from '@/lib/i18n/config';

function formatLocaleDate(iso: string, locale: Locale) {
  const tag = locale === 'en' ? 'en-US' : 'pt-BR';
  return new Date(iso).toLocaleDateString(tag, { month: 'long', year: 'numeric' });
}

export async function AuthorBox({
  author,
  reviewedBy,
  reviewedAt,
  locale = defaultLocale,
}: {
  author: Author;
  reviewedBy?: Author | null;
  reviewedAt?: string;
  locale?: Locale;
}) {
  const isMedicalAuthor = author.credentials?.includes('CRM');
  const reviewer = reviewedBy ?? (isMedicalAuthor ? author : null);
  const t = await getTranslations({ locale, namespace: 'blogIndex' });

  return (
    <aside className="border-y border-petrol/10 py-8 my-12 grid gap-6 sm:grid-cols-[80px_1fr] items-start">
      <div className="relative h-20 w-20 rounded-full overflow-hidden bg-cream-200 border border-petrol/10">
        {author.photo ? (
          <Image src={author.photo} alt={author.name} fill className="object-cover" sizes="80px" />
        ) : (
          <div className="absolute inset-0 flex items-center justify-center label-upper text-petrol/40">FOTO</div>
        )}
      </div>
      <div className="space-y-2">
        <p className="label-upper text-gold">{t('authorLabel')}</p>
        <p className="heading-section text-petrol text-xl">{author.name}</p>
        <p className="label-upper text-petrol/55">{author.credentials}</p>
        <p className="text-petrol/75 text-sm leading-relaxed max-w-prose">{author.bio}</p>
        {reviewer && (
          <div className="pt-3 mt-2 border-t border-petrol/5 space-y-1">
            <p className="label-upper text-petrol/45">{t('clinicalReviewLabel')}</p>
            <p className="text-petrol/65 text-xs leading-relaxed">
              {t('clinicalReviewBy')} {reviewer.name} · {reviewer.credentials}
              {reviewedAt && <> · {formatLocaleDate(reviewedAt, locale)}</>}
            </p>
          </div>
        )}
      </div>
    </aside>
  );
}
