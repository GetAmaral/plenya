import { PlenyaSymbol } from '@plenya/brand/logo/PlenyaSymbol';

export function SymbolBridge() {
  return (
    <section className="bg-cream">
      <div className="site-container section">
        <div className="grid grid-cols-1 md:grid-cols-[1fr_auto_1fr] gap-12 md:gap-20 items-center">
          <p className="heading-section text-petrol text-xl md:text-2xl md:text-right max-w-xs md:ml-auto">
            Ter mais energia e disposição no dia a dia.
          </p>

          <PlenyaSymbol className="h-16 w-16 text-gold mx-auto" />

          <p className="heading-section text-petrol text-xl md:text-2xl max-w-xs">
            Sentir controle e clareza sobre a própria saúde.
          </p>
        </div>
      </div>
    </section>
  );
}
