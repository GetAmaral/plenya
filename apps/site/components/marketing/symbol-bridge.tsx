import { PlenyaInfinity } from '@plenya/brand/logo/PlenyaInfinity';

/**
 * Symbol bridge — mirrors brandbook page 22 mockup exactly.
 * Two pain/desire columns flanking the infinity P (continuidade do cuidado).
 */
export function SymbolBridge() {
  return (
    <section className="bg-cream">
      <div className="site-container section">
        <div className="grid grid-cols-1 md:grid-cols-[1fr_auto_1fr] gap-12 md:gap-20 items-center">
          <p className="heading-section text-petrol text-xl md:text-2xl md:text-right max-w-xs md:ml-auto">
            Ter mais energia e disposição no dia a dia.
          </p>

          <PlenyaInfinity className="h-12 md:h-16 w-auto text-gold mx-auto" />

          <p className="heading-section text-petrol text-xl md:text-2xl max-w-xs">
            Sentir controle e clareza sobre a própria saúde.
          </p>
        </div>
      </div>
    </section>
  );
}
