import { useTranslations } from 'next-intl';
import { Link } from '@/lib/i18n/navigation';
import { Wordmark } from '@/components/brand/wordmark';

export function SiteFooter() {
  const t = useTranslations('footer');
  const year = new Date().getFullYear();

  return (
    <footer className="bg-navy text-paper mt-32">
      {/* Bloco institucional — wordmark dark sobre navy + tagline */}
      <div className="editorial-container py-16 md:py-20 text-center border-b border-paper/10">
        <Wordmark variant="dark" size="lg" />
      </div>

      {/* Identificação clínica + colunas de navegação */}
      <div className="editorial-container py-14 grid md:grid-cols-[2fr_1fr_1fr] gap-10">
        <div className="space-y-3">
          <p className="font-serif text-base text-paper">{t('responsavel')}</p>
          <p className="font-sans text-sm text-paper/65">{t('responsavelTecnico')}</p>
          <p className="font-sans text-sm text-paper/65">{t('local')}</p>
        </div>

        <div className="space-y-3">
          <p className="label-meta text-gold">Navegação</p>
          <ul className="space-y-2 text-sm text-paper/75">
            <li><Link href="/sobre" className="hover:text-gold transition-colors">Sobre</Link></li>
            <li><Link href="/livro" className="hover:text-gold transition-colors">Livro</Link></li>
            <li><Link href="/escritos" className="hover:text-gold transition-colors">Escritos</Link></li>
            <li><Link href="/palestras" className="hover:text-gold transition-colors">Palestras</Link></li>
            <li><Link href="/ensino" className="hover:text-gold transition-colors">Ensino</Link></li>
            <li><Link href="/onde-atendo" className="hover:text-gold transition-colors">Onde atendo</Link></li>
            <li><Link href="/contato" className="hover:text-gold transition-colors">Contato</Link></li>
          </ul>
        </div>

        <div className="space-y-3">
          <p className="label-meta text-gold">Contato</p>
          <address className="space-y-2 text-sm text-paper/75 not-italic">
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
          </address>
        </div>
      </div>

      <div className="editorial-container pb-10 border-t border-paper/10 pt-6">
        <p className="font-sans text-xs text-paper/50">
          © {year} · Conteúdo educativo. Não substitui consulta médica.
        </p>
      </div>
    </footer>
  );
}
