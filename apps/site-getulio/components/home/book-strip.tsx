import Image from 'next/image';
import { useTranslations } from 'next-intl';
import { Link } from '@/lib/i18n/navigation';

export function BookStrip() {
  const t = useTranslations('home.book');
  return (
    <section className="bg-navy text-paper">
      <div className="editorial-container py-20 md:py-28">
        <div className="grid md:grid-cols-[280px_1fr] gap-10 md:gap-16 items-center">
          <div className="relative aspect-[2/3] w-full max-w-[280px] mx-auto md:mx-0 shadow-2xl">
            <Image
              src="/images/livro-capa.jpg"
              alt={t('coverAlt')}
              fill
              className="object-cover"
              sizes="(min-width: 768px) 280px, 70vw"
            />
          </div>

          <div className="space-y-6">
            <div className="flex items-center gap-3">
              <span className="filete-gold" aria-hidden="true" />
              <p className="label-meta text-gold">{t('kicker')}</p>
            </div>
            <h2 className="heading-section text-paper text-3xl md:text-5xl">
              {t('titleLine1')}<br />{t('titleLine2')}
            </h2>
            <p className="font-serif italic text-paper/75 text-xl md:text-2xl leading-snug -mt-2">
              {t('titleSubtitle')}
            </p>
            <blockquote className="font-serif text-paper/85 text-xl md:text-2xl italic leading-relaxed border-l-2 border-gold/60 pl-6">
              {t('quote')}
            </blockquote>
            <p className="font-sans text-sm text-paper/60">{t('isbn')}</p>
            <Link href={{ pathname: '/livros/[slug]', params: { slug: 'antes' } }} className="link-text-light inline-block font-sans text-sm">
              {t('cta')}
            </Link>
          </div>
        </div>
      </div>
    </section>
  );
}
