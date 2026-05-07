/**
 * Wordmark institucional do Dr. Getúlio Amaral — réplica fiel do
 * branding oficial (PNGs em /docs/site/drgetulio/).
 *
 * Estrutura:
 *   ─── DR. ───
 *   GETÚLIO AMARAL
 *   <tagline localizada>
 *
 * Variantes: light (sobre cream) | dark (sobre navy/preto).
 * Tamanhos: sm (header), md, lg, xl (hero).
 * Tagline opcional (default true).
 */
import { useTranslations } from 'next-intl';

type Variant = 'light' | 'dark';
type Size = 'sm' | 'md' | 'lg' | 'xl';

const NAME_SIZE: Record<Size, string> = {
  sm: 'text-xl md:text-2xl',
  md: 'text-3xl md:text-5xl',
  lg: 'text-5xl md:text-7xl',
  xl: 'text-6xl md:text-8xl',
};

const PREFIX_SIZE: Record<Size, string> = {
  sm: 'text-[0.55rem]',
  md: 'text-xs',
  lg: 'text-sm',
  xl: 'text-base',
};

const TAGLINE_SIZE: Record<Size, string> = {
  sm: 'text-[0.6rem]',
  md: 'text-[0.7rem] md:text-xs',
  lg: 'text-xs md:text-sm',
  xl: 'text-sm md:text-base',
};

const FILETE_WIDTH: Record<Size, string> = {
  sm: 'w-6',
  md: 'w-10',
  lg: 'w-16',
  xl: 'w-20',
};

export function Wordmark({
  variant = 'light',
  size = 'md',
  tagline = true,
  className = '',
}: {
  variant?: Variant;
  size?: Size;
  tagline?: boolean;
  className?: string;
}) {
  const isDark = variant === 'dark';
  const ink = isDark ? 'text-paper' : 'text-ink';
  const inkSoft = isDark ? 'text-paper/70' : 'text-ink-soft/85';
  const gold = 'bg-gold';
  const t = useTranslations('wordmark');

  return (
    <div className={`inline-flex flex-col items-center gap-2 ${className}`}>
      <div className="flex items-center gap-3">
        <span className={`${gold} h-px ${FILETE_WIDTH[size]}`} aria-hidden="true" />
        <span className={`wordmark-prefix ${ink} ${PREFIX_SIZE[size]}`}>DR.</span>
        <span className={`${gold} h-px ${FILETE_WIDTH[size]}`} aria-hidden="true" />
      </div>
      <h1 className={`wordmark-name ${ink} ${NAME_SIZE[size]} font-normal m-0`}>
        GETÚLIO AMARAL
      </h1>
      {tagline && (
        <p className={`wordmark-tagline ${inkSoft} ${TAGLINE_SIZE[size]} mt-1`}>
          {t('tagline')}
        </p>
      )}
    </div>
  );
}
