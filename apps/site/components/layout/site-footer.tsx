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
      <div className="relative site-container py-20 grid gap-12 md:grid-cols-2 lg:grid-cols-6">
        <div className="lg:col-span-2 space-y-4">
          <PlenyaWordmark className="h-6 w-auto text-cream" />
          <p className="text-cream/50 text-sm leading-relaxed max-w-xs">
            Medicina que antecipa. Saúde, performance e longevidade.
          </p>
          <div className="flex flex-col gap-2.5 text-sm pt-4">
            <a href={`mailto:${brand.email}`} className="text-cream/70 hover:text-cream transition">{brand.email}</a>
            <a href={brand.social.instagram} target="_blank" rel="noreferrer" className="text-cream/70 hover:text-cream transition">@plenyaSaude</a>
            <span className="text-cream/40">Londrina · PR</span>
          </div>
        </div>

        <div className="space-y-3">
          <p className="label-upper text-gold">Conheça a Plenya</p>
          <nav className="flex flex-col gap-2.5 text-sm">
            <Link href="/a-plenya" className="hover:text-cream transition">{tNav('about')}</Link>
            <Link href="/dr-getulio" className="hover:text-cream transition">{tNav('drGetulio')}</Link>
            <Link href="/equipe" className="hover:text-cream transition">{tNav('team')}</Link>
            <Link href="/depoimentos" className="hover:text-cream transition">Depoimentos</Link>
          </nav>
        </div>

        <div className="space-y-3">
          <p className="label-upper text-gold">Como cuidamos</p>
          <nav className="flex flex-col gap-2.5 text-sm">
            <Link href="/metodo-agir" className="hover:text-cream transition">Método AGIR</Link>
            <Link href="/escore-plenya" className="hover:text-cream transition">Escore Plenya</Link>
          </nav>
        </div>

        <div className="space-y-3">
          <p className="label-upper text-gold">Comece</p>
          <nav className="flex flex-col gap-2.5 text-sm">
            <Link href="/diagnostico" className="hover:text-cream transition">Diagnóstico — é para mim?</Link>
            <Link href="/consultas" className="hover:text-cream transition">{tNav('consultations')}</Link>
            <Link href="/continuum" className="hover:text-cream transition">{tNav('plans')}</Link>
            <Link href="/contato" className="hover:text-cream transition">{tNav('contact')}</Link>
          </nav>
        </div>

        <div className="space-y-3">
          <p className="label-upper text-gold">Aprenda</p>
          <nav className="flex flex-col gap-2.5 text-sm">
            <Link href="/blog" className="hover:text-cream transition">{tNav('blog')}</Link>
            <a href={brand.appUrl} className="hover:text-cream transition">{t('professionals')}</a>
          </nav>
        </div>
      </div>

      <div className="relative border-t border-cream/10">
        <div className="site-container py-6 flex flex-col md:flex-row justify-between gap-4 text-cream/40 text-xs">
          <span>© {year} {brand.legalName}. Todos os direitos reservados.</span>
          <div className="flex gap-6 flex-wrap">
            <Link href="/privacidade" className="hover:text-cream/70 transition">{t('privacy')}</Link>
            <Link href="/termos" className="hover:text-cream/70 transition">{t('terms')}</Link>
            <Link href="/lgpd/direitos" className="hover:text-cream/70 transition">Direitos LGPD</Link>
            <Link href="/lgpd/encarregado" className="hover:text-cream/70 transition">Encarregado</Link>
          </div>
        </div>
      </div>
    </footer>
  );
}
