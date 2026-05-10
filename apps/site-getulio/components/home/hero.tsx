import Image from 'next/image';
import { useTranslations } from 'next-intl';
import { Link } from '@/lib/i18n/navigation';
import { Wordmark } from '@/components/brand/wordmark';

/**
 * Hero da home — foto lado a lado com bloco de texto. Foto sobe pra perto
 * do topo (era exilada numa faixa estreita abaixo do hero, ficava deslocada).
 *
 * Hierarquia de CTA:
 *   1. Marcar consulta (gold) → /onde-atendo  (ação principal — clínica)
 *   2. Sobre mim (outline)  → /sobre
 *   3. Conhecer o livro      → /livros/antes  (link sutil, secundário)
 */
export function Hero() {
  const t = useTranslations('home.hero');
  return (
    <section className="border-b border-rule">
      <div className="editorial-container pt-12 md:pt-20 pb-16 md:pb-24">
        <div className="grid md:grid-cols-[1fr_360px] lg:grid-cols-[1fr_400px] gap-10 md:gap-14 lg:gap-20 items-center">
          {/* Bloco texto */}
          <div className="space-y-10 order-2 md:order-1 text-center md:text-left">
            <Wordmark size="xl" as="h1" />

            <p className="label-meta text-ink-muted">{t('credentials')}</p>

            <div className="max-w-xl mx-auto md:mx-0 space-y-5 prose-body text-ink-soft">
              <p>{t('lead1')}</p>
              <p>{t('lead2')}</p>
            </div>

            <div className="flex flex-col sm:flex-row items-center md:items-start sm:items-center gap-4 sm:gap-5 flex-wrap">
              <Link href="/onde-atendo" className="btn-gold">
                {t('ctaConsult')}
              </Link>
              <Link href="/sobre" className="btn-outline">
                {t('ctaAbout')}
              </Link>
              <Link href="/livros/antes" className="link-text font-sans text-sm">
                {t('ctaBook')} →
              </Link>
            </div>
          </div>

          {/* Retrato editorial — agora prominente ao lado do bloco texto */}
          <div className="order-1 md:order-2">
            <div className="relative aspect-[3/4] w-full max-w-[400px] mx-auto md:mx-0 shadow-xl ring-1 ring-gold/20">
              <Image
                src="/images/getulio-hero-bw.jpg"
                alt={t('portraitAlt')}
                fill
                priority
                className="object-cover grayscale"
                sizes="(min-width: 1024px) 400px, (min-width: 768px) 360px, 80vw"
              />
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
