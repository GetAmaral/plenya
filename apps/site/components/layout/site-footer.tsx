import { useTranslations } from 'next-intl';
import { brand } from '@plenya/brand';
import { PlenyaWordmark, PlenyaSymbol } from '@plenya/brand/logo';
import { Link } from '@/lib/i18n/navigation';

export function SiteFooter() {
  const t = useTranslations('footer');
  const tNav = useTranslations('nav');
  const year = new Date().getFullYear();

  return (
    <footer className="bg-petrol text-cream/80 relative overflow-hidden">
      {/* P watermark — selo da marca no canto inferior direito */}
      <PlenyaSymbol
        aria-hidden="true"
        focusable="false"
        className="hidden md:block absolute -bottom-10 -right-10 h-72 w-auto text-cream/[0.03] pointer-events-none"
      />
      <div className="relative site-container py-20 grid gap-12 md:grid-cols-2 lg:grid-cols-5">
        <div className="lg:col-span-2 space-y-4">
          <PlenyaWordmark className="h-6 w-auto text-cream" />
          <p className="text-cream/50 text-sm leading-relaxed max-w-xs">
            Medicina que antecipa. Saúde, performance e longevidade.
          </p>
        </div>

        <div className="space-y-3">
          <p className="label-upper text-gold">{t('navigation')}</p>
          <nav className="flex flex-col gap-2.5 text-sm">
            <Link href="/" className="hover:text-cream transition">{tNav('home')}</Link>
            <Link href="/a-plenya" className="hover:text-cream transition">{tNav('about')}</Link>
            <Link href="/equipe" className="hover:text-cream transition">{tNav('team')}</Link>
            <Link href="/dr-getulio" className="hover:text-cream transition">{tNav('drGetulio')}</Link>
            <Link href="/blog" className="hover:text-cream transition">{tNav('blog')}</Link>
            <Link href="/depoimentos" className="hover:text-cream transition">Depoimentos</Link>
          </nav>
        </div>

        <div className="space-y-3">
          <p className="label-upper text-gold">{t('services')}</p>
          <nav className="flex flex-col gap-2.5 text-sm">
            <Link href="/metodo-agir" className="hover:text-cream transition">Método AGIR</Link>
            <Link href="/escore-plenya" className="hover:text-cream transition">Escore Plenya</Link>
            <Link href="/consultas" className="hover:text-cream transition">{tNav('consultations')}</Link>
            <Link href="/planos" className="hover:text-cream transition">{tNav('plans')}</Link>
            <Link href="/contato" className="hover:text-cream transition">{tNav('contact')}</Link>
            <a href={brand.appUrl} className="hover:text-cream transition">{t('professionals')}</a>
          </nav>
        </div>

        <div className="space-y-3">
          <p className="label-upper text-gold">{t('contact')}</p>
          <div className="flex flex-col gap-2.5 text-sm">
            <a href={`mailto:${brand.email}`} className="hover:text-cream transition">{brand.email}</a>
            <a href={brand.social.instagram} target="_blank" rel="noreferrer" className="hover:text-cream transition">@drGetulioAmaralFilho</a>
            <span className="text-cream/40">Londrina · PR</span>
          </div>
        </div>
      </div>

      <div className="relative border-t border-cream/10">
        <div className="site-container py-6 flex flex-col md:flex-row justify-between gap-4 text-cream/40 text-xs">
          <span>© {year} {brand.legalName}. Todos os direitos reservados.</span>
          <div className="flex gap-6">
            <Link href="/privacidade" className="hover:text-cream/70 transition">{t('privacy')}</Link>
            <Link href="/termos" className="hover:text-cream/70 transition">{t('terms')}</Link>
          </div>
        </div>
      </div>
    </footer>
  );
}
