import { useTranslations } from 'next-intl';
import { Link } from '@/lib/i18n/navigation';

export function SiteFooter() {
  const t = useTranslations('footer');
  const year = new Date().getFullYear();

  return (
    <footer className="border-t border-rule mt-32">
      <div className="editorial-container py-14 grid md:grid-cols-[2fr_1fr_1fr] gap-10">
        <div className="space-y-3">
          <p className="font-serif text-base text-ink">{t('responsavel')}</p>
          <p className="font-sans text-sm text-ink-muted">{t('responsavelTecnico')}</p>
          <p className="font-sans text-sm text-ink-muted">{t('local')}</p>
        </div>

        <div className="space-y-3">
          <p className="label-meta">Navegação</p>
          <ul className="space-y-2 text-sm text-ink-soft">
            <li><Link href="/sobre" className="hover:text-bordo">Sobre</Link></li>
            <li><Link href="/livro" className="hover:text-bordo">Livro</Link></li>
            <li><Link href="/palestras" className="hover:text-bordo">Palestras</Link></li>
            <li><Link href="/onde-atendo" className="hover:text-bordo">Onde atendo</Link></li>
            <li><Link href="/contato" className="hover:text-bordo">Contato</Link></li>
          </ul>
        </div>

        <div className="space-y-3">
          <p className="label-meta">Contato</p>
          <ul className="space-y-2 text-sm text-ink-soft">
            <li><a href="mailto:contato@drgetulioamaralfilho.com.br" className="hover:text-bordo">contato@drgetulioamaralfilho.com.br</a></li>
            <li><a href="https://instagram.com/drGetulioAmaralFilho" target="_blank" rel="noreferrer" className="hover:text-bordo">@drGetulioAmaralFilho</a></li>
          </ul>
        </div>
      </div>

      <div className="editorial-container pb-8">
        <p className="font-sans text-xs text-ink-muted/70">© {year} · Conteúdo educativo. Não substitui consulta médica.</p>
      </div>
    </footer>
  );
}
