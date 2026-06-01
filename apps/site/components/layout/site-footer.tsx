import { useTranslations, useLocale } from 'next-intl';
import { brand } from '@plenya/brand';
import { PlenyaWordmark, PlenyaSymbol } from '@plenya/brand/logo';
import { Link } from '@/lib/i18n/navigation';

export function SiteFooter() {
  const t = useTranslations('footer');
  const tNav = useTranslations('nav');
  const tMenu = useTranslations('navMenu');
  const locale = useLocale();
  const bookExternalUrl =
    locale === 'en'
      ? 'https://drgetulioamaralfilho.com.br/en/books/antes'
      : 'https://drgetulioamaralfilho.com.br/livros/antes';
  const year = new Date().getFullYear();

  return (
    <footer className="bg-petrol text-cream/90 relative overflow-hidden">
      {/* P watermark — selo da marca no canto inferior direito */}
      <PlenyaSymbol
        aria-hidden="true"
        focusable="false"
        className="hidden md:block absolute -bottom-10 -right-10 h-72 w-auto text-cream/3 pointer-events-none"
      />
      <div className="relative site-container py-20 grid gap-12 md:grid-cols-2 lg:grid-cols-6">
        <div className="lg:col-span-2 space-y-4">
          <PlenyaWordmark className="h-6 w-auto text-cream" />
          <p className="text-cream/65 text-base leading-relaxed max-w-xs">
            {t('tagline')}
          </p>
          <div className="flex flex-col gap-3 text-base pt-4">
            <a href={`mailto:${brand.email}`} className="text-cream/85 hover:text-cream transition">{brand.email}</a>
            <a href="https://wa.me/5543999748899" target="_blank" rel="noreferrer" className="text-cream/85 hover:text-cream transition">{t('whatsapp')}</a>
            <a href={brand.social.instagram} target="_blank" rel="noreferrer" className="text-cream/85 hover:text-cream transition">@plenyaSaude</a>
            <a
              href="https://drgetulioamaralfilho.com.br"
              target="_blank"
              rel="noreferrer"
              className="text-cream/85 hover:text-cream transition"
            >
              {t('drGetulioSiteLink')}
            </a>
            <a
              href={bookExternalUrl}
              target="_blank"
              rel="noreferrer"
              className="text-cream/85 hover:text-cream transition"
            >
              {t('bookExternalLink')}
            </a>
            <address className="text-cream/55 not-italic leading-relaxed">
              {brand.address.street}<br />
              {brand.address.complement}<br />
              {brand.address.neighborhood} · {brand.address.city}/{brand.address.state} · {brand.address.postalCode}
            </address>
          </div>
        </div>

        <div className="space-y-3">
          <p className="label-upper text-gold">{tMenu('groupAbout')}</p>
          <nav className="flex flex-col gap-3 text-base text-cream/80">
            <Link href="/a-plenya" className="hover:text-cream transition">{tNav('about')}</Link>
            <Link href="/dr-getulio" className="hover:text-cream transition">{tNav('drGetulio')}</Link>
            <Link href="/equipe" className="hover:text-cream transition">{tNav('team')}</Link>
            <Link href="/depoimentos" className="hover:text-cream transition">{tMenu('testimonials')}</Link>
          </nav>
        </div>

        <div className="space-y-3">
          <p className="label-upper text-gold">{tMenu('groupCare')}</p>
          <nav className="flex flex-col gap-3 text-base text-cream/80">
            <Link href="/metodo-agir" className="hover:text-cream transition">{tMenu('method')}</Link>
            <Link href="/escore-plenya" className="hover:text-cream transition">{tMenu('scoreFull')}</Link>
          </nav>
        </div>

        <div className="space-y-3">
          <p className="label-upper text-gold">{tMenu('groupStart')}</p>
          <nav className="flex flex-col gap-3 text-base text-cream/80">
            <Link href="/diagnostico" className="hover:text-cream transition">{tMenu('diagnostic')}</Link>
            <Link href="/consultas" className="hover:text-cream transition">{tNav('consultations')}</Link>
            <Link href="/continuum" className="hover:text-cream transition">{tNav('plans')}</Link>
            <Link href="/contato" className="hover:text-cream transition">{tNav('contact')}</Link>
          </nav>
        </div>

        <div className="space-y-3">
          <p className="label-upper text-gold">{tMenu('groupLearn')}</p>
          <nav className="flex flex-col gap-3 text-base text-cream/80">
            <Link href="/blog" className="hover:text-cream transition">{tNav('blog')}</Link>
            <a href={brand.appUrl} className="hover:text-cream transition">{t('professionals')}</a>
          </nav>
        </div>
      </div>

      <div className="relative border-t border-cream/10">
        <div className="site-container py-5 text-cream/55 text-xs leading-relaxed">
          <p>
            <span className="text-cream/70">{t('directorPrefix')}</span> {t('directorName')}
          </p>
          <p className="mt-1 text-cream/40">
            {t('disclaimer')}
          </p>
        </div>
      </div>

      <div className="relative border-t border-cream/10">
        <div className="site-container py-6 flex flex-col md:flex-row justify-between gap-4 text-cream/40 text-xs">
          <span>© {year} {brand.companyName} · CNPJ {brand.cnpj}. {t('rights')}</span>
          <div className="flex gap-6 flex-wrap">
            <Link href="/privacidade" className="hover:text-cream/70 transition">{t('privacy')}</Link>
            <Link href="/termos" className="hover:text-cream/70 transition">{t('terms')}</Link>
            <Link href="/lgpd/direitos" className="hover:text-cream/70 transition">{t('lgpdRights')}</Link>
            <Link href="/lgpd/encarregado" className="hover:text-cream/70 transition">{t('lgpdOfficer')}</Link>
          </div>
        </div>
      </div>
    </footer>
  );
}
