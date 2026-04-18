import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { agirLetters, agirTotals } from '@/lib/agir-structure';

export const metadata: Metadata = {
  title: 'Escore Plenya',
  description: `O diagnóstico estruturado que mostra onde sua saúde está — ${agirTotals.items}+ itens organizados em ${agirTotals.pillars} pilares clínicos dentro do Método AGIR.`,
};

const tiers = [
  {
    title: 'Versão Completa',
    desc: 'Aplicada pela equipe Plenya durante o acompanhamento. Inclui avaliação clínica, laboratorial, comportamental e funcional. Gera relatório detalhado com metas e plano de cuidado personalizado.',
  },
  {
    title: 'Versão Intermediária',
    desc: 'Disponível após a primeira consulta. Recorte focado nas metas prioritárias identificadas pela equipe médica.',
  },
  {
    title: 'Versão Light',
    desc: 'Disponível online em breve. Permite começar a entender sua saúde de forma acessível, sem necessidade de consulta prévia.',
  },
];

const lensColumns = [
  {
    title: 'Histórico',
    items: [
      'Pré-natal até hoje',
      'Familiar (3 gerações)',
      'Doenças crônicas e cirurgias',
      'Comportamental e cognitivo',
    ],
  },
  {
    title: 'Bioquímica',
    items: [
      'Hormônios (T livre, estradiol, tireoide)',
      'Metabolismo (glicemia, insulina, ApoB)',
      'Inflamação (PCR-us, IL-6, ferritina)',
      'Vitaminas, minerais e imagem (CAC, RM, densitometria)',
    ],
  },
  {
    title: 'Genética',
    items: [
      'Integrada em cada pilar clínico',
      'APOE e PCSK9 → Risco Cardiovascular',
      'MTHFR e TCF7L2 → Controle Glicêmico',
      'ALDH2 e CYP → Função Hepática',
    ],
  },
];

const letterAccent: Record<string, string> = {
  A: 'border-gold/40',
  G: 'border-ocean/40',
  I: 'border-sage/50',
  R: 'border-petrol/30',
};

export default async function ScorePage({ params }: { params: Promise<{ locale: string }> }) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <>
      {/* HERO */}
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32">
          <p className="label-upper text-gold mb-6">Diagnóstico</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,5rem)] text-cream max-w-2xl">
            Escore Plenya
          </h1>
          <p className="text-cream/70 text-lg mt-6 max-w-2xl">
            <strong className="text-cream">800+ itens em {agirTotals.pillars} pilares clínicos</strong>,
            organizados pelas 4 letras AGIR. O diagnóstico que mostra onde sua saúde está hoje e
            define para onde ir.
          </p>
        </div>
      </section>

      {/* ESCORE × AGIR */}
      <section className="bg-cream">
        <div className="site-container section">
          <div className="grid md:grid-cols-2 gap-10 max-w-4xl">
            <div className="space-y-3 border-l-4 border-gold pl-6">
              <p className="label-upper text-gold">Escore Plenya</p>
              <h2 className="heading-section text-petrol text-2xl md:text-3xl">Onde você está.</h2>
              <p className="text-petrol/75 leading-relaxed">
                A ferramenta de diagnóstico. Mapeia mais de 800 itens — do histórico familiar ao
                DNA, do sono à bioquímica — em uma pontuação única e por pilar.
              </p>
            </div>
            <div className="space-y-3 border-l-4 border-petrol/30 pl-6">
              <p className="label-upper text-petrol/60">Método AGIR</p>
              <h2 className="heading-section text-petrol text-2xl md:text-3xl">Para onde vamos.</h2>
              <p className="text-petrol/75 leading-relaxed">
                O framework de atuação. Toma o que o Escore revela e organiza o cuidado em quatro
                frentes interdependentes — com equipe e plano único.
              </p>
            </div>
          </div>
        </div>
      </section>

      {/* COMO FUNCIONA — diagrama enxuto */}
      <section className="bg-paper">
        <div className="site-container section">
          <div className="max-w-3xl mb-10 space-y-4">
            <p className="label-upper text-gold">Como funciona</p>
            <h2 className="heading-section text-petrol text-3xl md:text-4xl">
              Cada item entra em um pilar. Cada pilar pertence a uma letra.
            </h2>
            <p className="text-petrol/70 leading-relaxed max-w-2xl">
              Você recebe três níveis de pontuação — global, por letra e por pilar — para ver com
              precisão onde estão as forças e as lacunas.
            </p>
          </div>

          <div className="flex flex-wrap items-center gap-3 md:gap-6 max-w-3xl text-petrol/85">
            <span className="px-5 py-3 border border-petrol/20 rounded-full label-upper text-xs">
              Item
            </span>
            <span className="text-gold text-xl">→</span>
            <span className="px-5 py-3 border border-petrol/20 rounded-full label-upper text-xs">
              Pilar clínico
            </span>
            <span className="text-gold text-xl">→</span>
            <span className="px-5 py-3 border border-petrol/20 rounded-full label-upper text-xs">
              Letra AGIR
            </span>
            <span className="text-gold text-xl">→</span>
            <span className="px-5 py-3 bg-petrol text-cream rounded-full label-upper text-xs">
              Escore
            </span>
          </div>
        </div>
      </section>

      {/* AS 4 LETRAS — visão agregada */}
      <section className="bg-cream">
        <div className="site-container section">
          <div className="max-w-3xl mb-12 space-y-4">
            <p className="label-upper text-gold">Cobertura por letra</p>
            <h2 className="heading-section text-petrol text-3xl md:text-4xl">
              {agirTotals.pillars} pilares em 4 letras.
            </h2>
            <p className="text-petrol/70 leading-relaxed">
              Cada letra concentra o conjunto de pilares clínicos da sua área. O detalhe completo
              de cada pilar está na página do Método AGIR.
            </p>
          </div>

          <div className="grid sm:grid-cols-2 lg:grid-cols-4 gap-6">
            {agirLetters.map((letter) => (
              <Link
                key={letter.code}
                href="/metodo-agir"
                className={`group p-6 border-2 rounded-sm bg-paper space-y-4 transition hover:shadow-lg hover:-translate-y-px ${letterAccent[letter.code]}`}
              >
                <div className="flex items-baseline justify-between">
                  <span className="heading-section text-petrol text-6xl leading-none">{letter.code}</span>
                  <span className="label-upper text-petrol/40 text-[10px]">
                    {letter.pillarCount} pilares
                  </span>
                </div>
                <h3 className="heading-section text-petrol text-base leading-tight">
                  {letter.name}
                </h3>
                <p className="label-upper text-gold text-[10px] pt-2 border-t border-petrol/10">
                  ~{letter.itemCount} itens · ver detalhe →
                </p>
              </Link>
            ))}
          </div>
        </div>
      </section>

      {/* O QUE OLHAMOS — 3 colunas-âncora */}
      <section className="bg-paper">
        <div className="site-container section">
          <div className="max-w-3xl mb-12 space-y-4">
            <p className="label-upper text-gold">O que o Escore enxerga</p>
            <h2 className="heading-section text-petrol text-3xl md:text-4xl">
              Três lentes sobre uma vida inteira.
            </h2>
          </div>

          <div className="grid md:grid-cols-3 gap-10 max-w-5xl">
            {lensColumns.map((col) => (
              <div key={col.title} className="space-y-4 border-t-2 border-gold pt-6">
                <h3 className="heading-section text-petrol text-xl">{col.title}</h3>
                <ul className="space-y-2.5">
                  {col.items.map((item) => (
                    <li key={item} className="text-petrol/75 leading-relaxed flex gap-3">
                      <span className="text-gold mt-1.5 leading-none text-xs">●</span>
                      <span>{item}</span>
                    </li>
                  ))}
                </ul>
              </div>
            ))}
          </div>

          <p className="label-upper text-petrol/50 mt-12 max-w-3xl text-xs">
            Lista parcial · o painel completo cobre 800+ marcadores específicos por caso.
          </p>
        </div>
      </section>

      {/* PONTUAÇÃO EXEMPLO */}
      <section className="bg-cream-100">
        <div className="site-container section flex flex-col items-center gap-4">
          <svg viewBox="0 0 300 300" className="w-64 h-64 md:w-80 md:h-80" aria-label="Exemplo de pontuação do Escore Plenya: 78">
            <circle cx="150" cy="150" r="140" stroke="#063b4f" strokeOpacity="0.08" strokeWidth="1" fill="none" />
            <circle cx="150" cy="150" r="110" stroke="#063b4f" strokeOpacity="0.12" strokeWidth="1" fill="none" />
            <circle cx="150" cy="150" r="80" stroke="#063b4f" strokeOpacity="0.16" strokeWidth="1" fill="none" />
            <path d="M 150 10 A 140 140 0 1 1 40 260" stroke="#b38645" strokeWidth="2" fill="none" strokeLinecap="round" />
            <text x="150" y="160" textAnchor="middle" fontFamily="'Cormorant Garamond', serif" fontSize="56" fill="#063b4f" letterSpacing="-2">78</text>
          </svg>
          <p className="label-upper text-petrol/50 text-center">
            Exemplo de pontuação global · escala 0–100
          </p>
        </div>
      </section>

      {/* VERSÕES */}
      <section className="bg-paper">
        <div className="site-container section">
          <p className="label-upper text-gold mb-10">Versões do Escore</p>
          <div className="grid md:grid-cols-3 gap-8">
            {tiers.map((tier, i) => (
              <div key={tier.title} className="border-t border-petrol/15 pt-8 space-y-4">
                <span className="label-upper text-petrol/40">0{i + 1}</span>
                <h3 className="heading-section text-petrol text-2xl">{tier.title}</h3>
                <p className="text-petrol/70 leading-relaxed">{tier.desc}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* CROSS-LINKS */}
      <section className="bg-cream">
        <div className="site-container section grid md:grid-cols-2 gap-8">
          <Link href="/metodo-agir" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">Método</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">Como agimos a partir do Escore →</p>
          </Link>
          <Link href="/planos" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">Planos</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">Comece seu acompanhamento →</p>
          </Link>
        </div>
      </section>
    </>
  );
}
