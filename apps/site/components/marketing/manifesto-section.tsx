import { manifesto } from '@plenya/brand/brand';
import { PlenyaInfinity } from '@plenya/brand/logo';

/**
 * Manifesto section — direct rendering of the brandbook manifesto (page 10).
 * Petrol background, gold infinity flourish, generous spacing.
 */
export function ManifestoSection() {
  const lines = manifesto.slice(0, -1);
  const closing = manifesto[manifesto.length - 1];

  return (
    <section className="bg-petrol text-cream relative overflow-hidden">
      <div className="site-container section relative z-10">
        <div className="max-w-3xl mx-auto text-center space-y-8">
          <PlenyaInfinity
            aria-hidden="true"
            focusable="false"
            className="h-14 md:h-20 w-auto text-gold mx-auto"
          />

          <p className="label-upper text-gold">Manifesto</p>

          <div className="space-y-4">
            {lines.map((line, i) => (
              <p
                key={i}
                className="heading-section text-cream/95 text-2xl md:text-3xl leading-relaxed"
              >
                {line}
              </p>
            ))}
          </div>

          <p className="heading-section text-gold text-2xl md:text-3xl pt-6">
            {closing}
          </p>
        </div>
      </div>
    </section>
  );
}
