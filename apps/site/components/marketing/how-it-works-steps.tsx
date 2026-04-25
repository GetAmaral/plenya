import { Link } from '@/lib/i18n/navigation';

export type HowStep = {
  n: string;
  title: string;
  body: string;
  detail?: string;
};

const defaultSteps: HowStep[] = [
  {
    n: '01',
    title: 'Avaliar',
    body: 'Você responde ao Escore Plenya — Triagem gratuita ou versão completa em consulta — e a equipe organiza sua história, sintomas, exames, hábitos e medicamentos em um painel único.',
    detail: 'Triagem · Consulta · Continuum',
  },
  {
    n: '02',
    title: 'Interpretar',
    body: 'O médico — e, no Continuum, a equipe inteira — faz a leitura funcional do seu painel. Não é "está normal": é o que está distante do ótimo, o que merece intervenção e o que pode esperar.',
    detail: 'Leitura clínica integrada',
  },
  {
    n: '03',
    title: 'Plano personalizado',
    body: 'Conduta escrita e acordada com você: medicação se houver indicação, ajustes de estilo de vida, suplementação, encaminhamentos. No Continuum, plano único construído pelos quatro profissionais.',
    detail: 'Conduta · Box Plenya · Acompanhamento',
  },
  {
    n: '04',
    title: 'Reavaliar em ciclo',
    body: 'A cada três meses (Continuum) ou conforme indicação, novo Escore. A curva mostra progresso real, estagnação ou queda. O plano se ajusta ao que os dados dizem — não a opinião.',
    detail: 'Curva evolutiva visível para você e a equipe',
  },
];

export function HowItWorksSteps({
  steps = defaultSteps,
  label = 'Como funciona',
  title = 'Quatro passos. Um cuidado que evolui.',
  bg = 'bg-cream',
}: {
  steps?: HowStep[];
  label?: string;
  title?: string;
  bg?: string;
}) {
  return (
    <section className={bg}>
      <div className="site-container section">
        <div className="max-w-3xl mb-14 space-y-4">
          <p className="label-upper text-gold">{label}</p>
          <h2 className="heading-section text-petrol text-3xl md:text-5xl">{title}</h2>
        </div>

        <ol className="grid md:grid-cols-2 lg:grid-cols-4 gap-px bg-petrol/15 border-y border-petrol/15">
          {steps.map((s) => (
            <li key={s.n} className="bg-paper p-8 md:p-10 space-y-5">
              <p className="text-gold text-5xl md:text-6xl font-light leading-none tabular-nums">
                {s.n}
              </p>
              <h3 className="heading-section text-petrol text-2xl">{s.title}</h3>
              <p className="text-petrol/75 leading-relaxed">{s.body}</p>
              {s.detail && (
                <p className="label-upper text-petrol/50 text-[10px] tracking-[0.2em] pt-3 border-t border-petrol/10">
                  {s.detail}
                </p>
              )}
            </li>
          ))}
        </ol>

        <div className="mt-12 flex flex-wrap gap-4">
          <Link href="/escore-plenya/avaliar" className="btn-gold">
            Começar pela Triagem
          </Link>
          <Link href="/contato" className="btn-outline-dark">
            Falar com a equipe
          </Link>
        </div>
      </div>
    </section>
  );
}
