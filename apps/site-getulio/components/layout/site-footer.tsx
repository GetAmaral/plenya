import { useTranslations } from 'next-intl';
import { Link } from '@/lib/i18n/navigation';
import { Wordmark } from '@/components/brand/wordmark';

export function SiteFooter() {
  const t = useTranslations('footer');
  const tNav = useTranslations('nav');
  const year = new Date().getFullYear();

  return (
    <footer className="bg-navy text-paper mt-32">
      {/* Wordmark institucional */}
      <div className="editorial-container pt-14 md:pt-16 pb-10 text-center border-b border-paper/10">
        <Wordmark variant="dark" size="md" />
      </div>

      {/* Identificação clínica + colunas de navegação */}
      <div className="editorial-container py-12 grid md:grid-cols-[2fr_1fr_1fr] gap-10">
        <div className="space-y-3">
          <p className="font-serif text-base text-paper">{t('responsavel')}</p>
          <p className="font-sans text-sm text-paper/70">{t('responsavelTecnico')}</p>
          <p className="font-sans text-sm text-paper/70">{t('local')}</p>
        </div>

        <div className="space-y-3">
          <p className="label-meta text-gold">{t('navigation')}</p>
          <ul className="space-y-2 text-sm text-paper/85">
            <li><Link href="/sobre" className="hover:text-gold transition-colors">{tNav('sobre')}</Link></li>
            <li><Link href="/livros" className="hover:text-gold transition-colors">{tNav('livros')}</Link></li>
            <li><Link href="/escritos" className="hover:text-gold transition-colors">{tNav('escritos')}</Link></li>
            <li><Link href="/palestras" className="hover:text-gold transition-colors">{tNav('palestras')}</Link></li>
            <li><Link href="/ensino" className="hover:text-gold transition-colors">{tNav('ensino')}</Link></li>
            <li><Link href="/onde-atendo" className="hover:text-gold transition-colors">{tNav('ondeAtendo')}</Link></li>
            <li><Link href="/contato" className="hover:text-gold transition-colors">{tNav('contato')}</Link></li>
          </ul>
        </div>

        <div className="space-y-3">
          <p className="label-meta text-gold">{t('contact')}</p>
          <address className="space-y-2 text-sm text-paper/85 not-italic">
            <p>
              <a
                href="mailto:contato@drgetulioamaralfilho.com.br"
                className="hover:text-gold transition-colors"
              >
                contato@drgetulioamaralfilho.com.br
              </a>
            </p>
            <p>
              <a
                href="https://instagram.com/drGetulioAmaralFilho"
                target="_blank"
                rel="noreferrer"
                className="hover:text-gold transition-colors"
              >
                @drGetulioAmaralFilho
              </a>
            </p>
            <p>
              <a
                href="https://instagram.com/plenyaSaude"
                target="_blank"
                rel="noreferrer"
                className="hover:text-gold transition-colors"
              >
                @plenyaSaude {t('clinicLabel')}
              </a>
            </p>
            <p>
              <a
                href="https://plenyasaude.com.br"
                target="_blank"
                rel="noreferrer"
                className="hover:text-gold transition-colors"
              >
                plenyasaude.com.br ↗
              </a>
            </p>
          </address>
        </div>
      </div>

      <div className="editorial-container pb-10 border-t border-paper/10 pt-6">
        <p className="font-sans text-xs text-paper/50">
          {t('disclaimer', { year })}
        </p>
      </div>
    </footer>
  );
}
