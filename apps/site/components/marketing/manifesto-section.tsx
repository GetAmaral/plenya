import { PlenyaInfinity } from '@plenya/brand/logo';

const manifesto = [
  'Tudo está conectado.',
  'Corpo, mente, tempo e escolhas.',
  'Viver bem é o reflexo do que você escolhe todos os dias.',
  'É consciência. É liberdade.',
  'Plenitude não é um ponto de chegada.',
  'É uma linha contínua.',
] as const;

const closing = 'Plenya. Viva bem, viva mais.';

/**
 * Manifesto section — direct rendering of the brandbook manifesto (page 10).
 */
export function ManifestoSection() {
  return (
    <section className="bg-petrol text-cream relative overflow-hidden">
      <div className="site-container section relative z-10">
        <div className="max-w-3xl space-y-8">
          <PlenyaInfinity
            aria-hidden="true"
            focusable="false"
            className="h-14 md:h-20 w-auto text-gold"
          />

          <p className="label-upper text-gold">Manifesto</p>

          <div className="space-y-4">
            {manifesto.map((line, i) => (
              <p
                key={i}
                className="heading-section text-cream/95 text-2xl md:text-3xl leading-relaxed"
              >
                {line}
              </p>
            ))}
          </div>

          <p className="heading-section text-gold text-2xl md:text-3xl pt-6">{closing}</p>
        </div>
      </div>
    </section>
  );
}
