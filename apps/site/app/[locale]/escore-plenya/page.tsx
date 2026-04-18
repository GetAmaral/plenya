import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';

export const metadata: Metadata = {
  title: 'Escore Plenya — instrumento de medida do Método AGIR',
  description:
    'Mais de 800 itens em uma pontuação clara, evolutiva e personalizada. Histórico, bioquímica e genética traduzidos em uma medida única do seu estado de saúde.',
};

const steps = [
  {
    n: '01',
    title: 'Avaliação ampla',
    body:
      'Mais de 800 itens são levantados — do histórico pré-natal e familiar aos exames laboratoriais e marcadores genéticos. Não há pergunta solta: cada item alimenta uma decisão clínica.',
  },
  {
    n: '02',
    title: 'Organização AGIR',
    body:
      'Cada item entra em um pilar clínico do Método AGIR — agrupado em uma das quatro letras: Alimentação e Atividade Física, Gestão Metabólica, Integração Mente-Corpo, Ritmo Circadiano.',
  },
  {
    n: '03',
    title: 'Pontuação que evolui',
    body:
      'Você recebe uma pontuação global e uma por letra. A cada reavaliação, a curva mostra o que melhorou, o que estagnou e onde a próxima intervenção precisa entrar.',
  },
];

const lensColumns = [
  {
    title: 'Histórico',
    body:
      'Pré-natal, infância, adolescência, vida adulta. Doenças crônicas, cirurgias, medicamentos, vícios. Histórico familiar em três gerações.',
  },
  {
    title: 'Bioquímica',
    body:
      'Hormônios completos (testosterona livre, estradiol, tireoide, cortisol), metabolismo (insulina, HOMA-IR, ApoB), inflamação (PCR-us, IL-6), vitaminas e minerais, e exames de imagem como CAC e densitometria.',
  },
  {
    title: 'Genética',
    body:
      'Cerca de 80 variantes integradas ao pilar clínico correspondente — APOE no risco cardiovascular, MTHFR no controle glicêmico, ALDH2 na função hepática, ACTN3 na prescrição de exercícios.',
  },
];

const tiers = [
  {
    title: 'Versão Completa',
    desc: 'Aplicada pela equipe Plenya durante o acompanhamento. Avaliação clínica, laboratorial, comportamental e funcional integrais. Gera relatório detalhado com metas e plano personalizado.',
  },
  {
    title: 'Versão Intermediária',
    desc: 'Disponível após a primeira consulta. Recorte focado nas metas prioritárias identificadas pela equipe.',
  },
  {
    title: 'Versão Light',
    desc: 'Disponível online em breve. Permite começar a entender sua saúde de forma acessível, sem necessidade de consulta prévia.',
  },
];

export default async function ScorePage({ params }: { params: Promise<{ locale: string }> }) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <>
      {/* HERO */}
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32">
          <p className="label-upper text-gold mb-6">Instrumento</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,5rem)] text-cream max-w-2xl">
            Escore Plenya
          </h1>
          <p className="text-cream/75 text-xl md:text-2xl mt-8 max-w-2xl leading-relaxed font-light">
            O instrumento de medida do Método AGIR. Mais de 800 itens em uma
            pontuação clara, evolutiva e personalizada do seu estado de saúde.
          </p>
        </div>
      </section>

      {/* O QUE É — texto âncora curto */}
      <section className="bg-cream">
        <div className="site-container section">
          <div className="max-w-3xl space-y-6">
            <p className="label-upper text-gold">O que é</p>
            <p className="heading-section text-petrol text-2xl md:text-4xl leading-tight">
              O Escore Plenya não substitui um exame. Substitui o vácuo entre eles.
            </p>
            <p className="text-petrol/80 text-lg leading-relaxed">
              É a medida que o cuidado precisava ter para deixar de ser uma sequência
              desconexa de consultas e virar um processo com direção e progresso visível.
              Onde antes havia laudos isolados, agora há uma pontuação que sintetiza —
              e que evolui a cada reavaliação.
            </p>
          </div>
        </div>
      </section>

      {/* COMO FUNCIONA — 3 etapas */}
      <section className="bg-paper">
        <div className="site-container section">
          <div className="max-w-3xl mb-14 space-y-4">
            <p className="label-upper text-gold">Como funciona</p>
            <h2 className="heading-section text-petrol text-3xl md:text-4xl">
              Três etapas. Uma pontuação que faz sentido.
            </h2>
          </div>

          <ol className="grid md:grid-cols-3 gap-10 max-w-5xl">
            {steps.map((s) => (
              <li key={s.n} className="space-y-4">
                <p className="heading-section text-gold text-5xl leading-none">{s.n}</p>
                <h3 className="heading-section text-petrol text-xl">{s.title}</h3>
                <p className="text-petrol/75 leading-relaxed">{s.body}</p>
              </li>
            ))}
          </ol>
        </div>
      </section>

      {/* GAUGE EXEMPLO — peça visual central */}
      <section className="bg-cream">
        <div className="site-container section grid lg:grid-cols-2 gap-16 items-center">
          <div className="space-y-6">
            <p className="label-upper text-gold">Pontuação</p>
            <h2 className="heading-section text-petrol text-3xl md:text-4xl">
              Um número que diz onde você está. Uma curva que diz para onde está indo.
            </h2>
            <p className="text-petrol/80 text-lg leading-relaxed">
              A pontuação global resume o estado atual em uma escala de 0 a 100. As
              pontuações por letra mostram em quais frentes você está mais forte e
              quais precisam de atenção. A cada reavaliação, a curva no tempo torna o
              progresso (ou a estagnação) impossível de ignorar.
            </p>
          </div>

          <div className="flex flex-col items-center gap-4">
            <svg viewBox="0 0 300 300" className="w-64 h-64 md:w-80 md:h-80" aria-label="Exemplo de pontuação do Escore Plenya: 78">
              <circle cx="150" cy="150" r="140" stroke="#063b4f" strokeOpacity="0.08" strokeWidth="1" fill="none" />
              <circle cx="150" cy="150" r="110" stroke="#063b4f" strokeOpacity="0.12" strokeWidth="1" fill="none" />
              <circle cx="150" cy="150" r="80" stroke="#063b4f" strokeOpacity="0.16" strokeWidth="1" fill="none" />
              <path d="M 150 10 A 140 140 0 1 1 40 260" stroke="#b38645" strokeWidth="2" fill="none" strokeLinecap="round" />
              <text x="150" y="160" textAnchor="middle" fontFamily="'Cormorant Garamond', serif" fontSize="56" fill="#063b4f" letterSpacing="-2">78</text>
            </svg>
            <p className="label-upper text-petrol/50 text-center text-xs">
              Exemplo · escala 0–100
            </p>
          </div>
        </div>
      </section>

      {/* O QUE O ESCORE ENXERGA — 3 lentes */}
      <section className="bg-paper">
        <div className="site-container section">
          <div className="max-w-3xl mb-12 space-y-4">
            <p className="label-upper text-gold">O que o Escore enxerga</p>
            <h2 className="heading-section text-petrol text-3xl md:text-4xl">
              Três lentes sobre uma vida inteira.
            </h2>
            <p className="text-petrol/70 leading-relaxed">
              Para cada pilar do Método AGIR, o Escore combina três dimensões de evidência —
              não apenas o exame do dia.
            </p>
          </div>

          <div className="grid md:grid-cols-3 gap-10 max-w-5xl">
            {lensColumns.map((col) => (
              <div key={col.title} className="space-y-4 border-t-2 border-gold pt-6">
                <h3 className="heading-section text-petrol text-xl">{col.title}</h3>
                <p className="text-petrol/75 leading-relaxed">{col.body}</p>
              </div>
            ))}
          </div>

          <p className="label-upper text-petrol/50 mt-12 text-xs max-w-3xl">
            Lista parcial — a versão completa cobre 800+ marcadores específicos por caso.
          </p>
        </div>
      </section>

      {/* IMAGEM ÂNCORA */}
      <section className="bg-cream">
        <div className="site-container">
          <div className="relative aspect-[16/7] overflow-hidden">
            <Image
              src="/images/hero-score.jpg"
              alt="Pausa, reflexão e autoavaliação"
              fill
              className="object-cover"
              sizes="100vw"
            />
          </div>
        </div>
      </section>

      {/* VERSÕES */}
      <section className="bg-cream">
        <div className="site-container section">
          <div className="max-w-3xl mb-12 space-y-4">
            <p className="label-upper text-gold">Versões</p>
            <h2 className="heading-section text-petrol text-3xl md:text-4xl">
              Três níveis de profundidade.
            </h2>
          </div>
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

      {/* SAÍDA — link único pro método */}
      <section className="bg-petrol text-cream">
        <div className="site-container section text-center space-y-8">
          <div className="max-w-2xl mx-auto space-y-4">
            <p className="label-upper text-gold">Próximo passo</p>
            <h2 className="heading-section text-cream text-3xl md:text-4xl">
              O Escore mostra onde você está. <br className="hidden md:block" />
              O Método AGIR é como atuamos a partir daí.
            </h2>
          </div>
          <Link href="/metodo-agir" className="btn-gold">
            Conhecer o Método AGIR
          </Link>
        </div>
      </section>
    </>
  );
}
