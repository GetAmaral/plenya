import { PlenyaInfinity } from '@plenya/brand/logo';
import { Link } from '@/lib/i18n/navigation';

/**
 * Estrutura section — substitui a duplicação do manifesto na home.
 * Usa material do vídeo institucional 05 + o conceito "normal vs ótimo"
 * (vídeo 04). Manifesto integral fica apenas na página /a-plenya.
 */
export function EstruturaSection() {
  return (
    <section className="bg-petrol text-cream relative overflow-hidden">
      <div className="site-container section relative z-10">
        <div className="grid lg:grid-cols-[auto_1fr] gap-10 lg:gap-20 items-start max-w-5xl">
          <PlenyaInfinity
            aria-hidden="true"
            focusable="false"
            className="h-14 md:h-20 w-auto text-gold shrink-0"
          />

          <div className="space-y-8">
            <div className="space-y-4">
              <p className="label-upper text-gold">O que nos move</p>
              <h2 className="heading-section text-cream text-3xl md:text-5xl">
                Saúde não é sorte. <em className="not-italic text-gold">É estrutura.</em>
              </h2>
            </div>

            <p className="text-cream/80 text-lg leading-relaxed max-w-2xl">
              Quem melhora de verdade não é quem tem mais força de vontade — é quem tem método.
              E saúde não melhora com intenção: melhora com direção.
            </p>

            <p className="text-cream/70 text-lg leading-relaxed max-w-2xl">
              <strong className="text-cream">Normal não é o mesmo que ótimo.</strong> Atendemos
              quem não está doente, mas também não está bem — quem ouviu que os exames
              estavam &ldquo;normais&rdquo; e mesmo assim sente que algo não fecha.
              Antecipar é mais inteligente que reagir.
            </p>

            <div className="pt-4">
              <Link href="/metodo-agir" className="btn-outline-light">
                Conhecer o Método AGIR
              </Link>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
