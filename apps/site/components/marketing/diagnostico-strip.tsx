import { Link } from '@/lib/i18n/navigation';

export function DiagnosticoStrip() {
  return (
    <section className="bg-paper border-b border-petrol/5">
      <div className="site-container py-8 flex flex-col md:flex-row md:items-center md:justify-between gap-4">
        <div className="space-y-1">
          <p className="label-upper text-gold">Por onde começar</p>
          <p className="text-petrol/85 leading-relaxed">
            Não sabe qual caminho da Plenya é o seu?
            <span className="text-petrol/60"> Cinco perguntas e devolvemos a leitura.</span>
          </p>
        </div>
        <Link
          href="/diagnostico"
          className="inline-flex items-center self-start md:self-auto px-5 py-2.5 border border-petrol/30 text-petrol hover:bg-petrol hover:text-cream transition label-upper whitespace-nowrap"
        >
          Fazer o diagnóstico
        </Link>
      </div>
    </section>
  );
}
