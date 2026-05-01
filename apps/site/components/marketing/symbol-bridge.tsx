import { useTranslations } from 'next-intl';
import { PlenyaInfinity } from '@plenya/brand/logo';

/**
 * Symbol bridge — mirrors brandbook page 22 mockup exactly.
 * Two pain/desire columns flanking the infinity P (continuidade do cuidado).
 */
export function SymbolBridge() {
  const t = useTranslations('home');
  const left = t('symbolBridgeLeft').split('\n');
  const right = t('symbolBridgeRight').split('\n');
  return (
    <section className="bg-cream">
      <div className="site-container section">
        <div className="grid grid-cols-1 md:grid-cols-[1fr_auto_1fr] gap-12 md:gap-20 items-center">
          <p className="heading-section text-petrol text-2xl md:text-3xl md:text-right max-w-[14ch] md:ml-auto">
            {left.map((line, i) => (
              <span key={i}>
                {line}
                {i < left.length - 1 && <br />}
              </span>
            ))}
          </p>

          <PlenyaInfinity className="h-12 md:h-16 w-auto text-gold mx-auto" />

          <p className="heading-section text-petrol text-2xl md:text-3xl max-w-[14ch]">
            {right.map((line, i) => (
              <span key={i}>
                {line}
                {i < right.length - 1 && <br />}
              </span>
            ))}
          </p>
        </div>
      </div>
    </section>
  );
}
