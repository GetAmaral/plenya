import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { HowItWorksSteps } from '@/components/marketing/how-it-works-steps';
import { ComparatorVsConvenio } from '@/components/marketing/comparator-vs-convenio';
import { FaqAccordion } from '@/components/marketing/faq-accordion';

export const metadata: Metadata = {
  title: 'Como funciona — a medicina Plenya em quatro passos',
  description:
    'Avaliar, interpretar, plano personalizado e reavaliar em ciclo. Como o cuidado Plenya entra na sua vida — da Triagem gratuita ao Continuum.',
};

const faq = [
  {
    q: 'Por onde começo?',
    a: (
      <>
        Pela <Link href="/escore-plenya/avaliar" className="underline decoration-gold/60 underline-offset-4">Triagem gratuita</Link>{' '}
        do Escore Plenya — cerca de 35 perguntas em poucos minutos, sem cadastro. O resultado já indica
        onde a sua saúde está hoje e qual o próximo passo: uma Consulta Plenya pontual ou o Continuum.
      </>
    ),
  },
  {
    q: 'A consulta é presencial ou online?',
    a: 'A Consulta Plenya pode ser presencial (Londrina-PR) ou online por telemedicina, com a mesma profundidade clínica. O Continuum Plenya é 100% online — todos os encontros semanais com a equipe acontecem por videochamada.',
  },
  {
    q: 'Vocês trabalham com convênio?',
    a: 'Não. Atendemos exclusivamente em formato particular. Despesas médicas são dedutíveis no Imposto de Renda como gasto com saúde, conforme legislação vigente.',
  },
  {
    q: 'Quanto tempo até ver resultado?',
    a: 'A primeira leitura clínica acontece já na consulta inicial. No Continuum, a curva do Escore tende a mostrar movimento entre o primeiro e o terceiro mês — porque é nesse intervalo que os primeiros ajustes de plano amadurecem.',
  },
  {
    q: 'O que entra no Box Plenya?',
    a: 'No Continuum, após a elaboração do plano, um box personalizado chega até você — mimos selecionados pela equipe e suplementos ou manipulados específicos do seu protocolo. Novos boxes seguem o ritmo do plano.',
  },
  {
    q: 'Posso cancelar o Continuum se mudar de ideia?',
    a: 'Sim. As condições de adesão e cancelamento são apresentadas em conversa direta antes da contratação, conforme a modalidade escolhida (semestral ou anual).',
  },
];

export default async function ComoFuncionaPage({ params }: { params: Promise<{ locale: string }> }) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <>
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32 max-w-3xl">
          <p className="label-upper text-gold mb-6">Como funciona</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,4.5rem)] text-cream">
            Quatro passos. Um cuidado que evolui com você.
          </h1>
          <p className="text-cream/75 text-lg leading-relaxed mt-8 max-w-2xl">
            A medicina Plenya não cabe em uma consulta — cabe em um ciclo. Avaliar amplo,
            interpretar com profundidade, agir com plano escrito, reavaliar com método.
            E começar de novo, no ponto onde os dados pediram.
          </p>
        </div>
      </section>

      <HowItWorksSteps bg="bg-cream" />

      <ComparatorVsConvenio bg="bg-paper" />

      <FaqAccordion title="Perguntas que recebemos toda semana." items={faq} />

      <section className="bg-petrol text-cream">
        <div className="site-container section text-center space-y-6">
          <p className="label-upper text-gold">Próximo passo</p>
          <h2 className="heading-section text-cream text-3xl md:text-4xl max-w-2xl mx-auto">
            Comece pela Triagem. Em cinco minutos você tem a primeira leitura.
          </h2>
          <div className="flex flex-wrap justify-center gap-4 pt-4">
            <Link href="/escore-plenya/avaliar" className="btn-gold">
              Fazer a Triagem
            </Link>
            <Link href="/contato" className="btn-outline-dark border-cream/40 text-cream">
              Falar com a equipe
            </Link>
          </div>
        </div>
      </section>
    </>
  );
}
