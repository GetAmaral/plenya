import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { getAllDoctors } from '@/lib/team';
import { ComparatorVsConvenio } from '@/components/marketing/comparator-vs-convenio';
import { FaqAccordion } from '@/components/marketing/faq-accordion';
import { FaqSchema } from '@/components/seo/faq-schema';
import { TestimonialsInline } from '@/components/testimonials/testimonials-inline';

const consultasFaq = [
  {
    q: 'Quanto tempo dura uma Consulta Plenya?',
    a: 'Em média 60 a 90 minutos na primeira consulta — tempo necessário para revisar história, exames anteriores e construir uma conduta. Retornos costumam durar 45 a 60 minutos.',
  },
  {
    q: 'É presencial ou online?',
    a: 'Os dois — você escolhe. Presencial em nossa clínica em Londrina-PR ou online por telemedicina, com a mesma profundidade clínica.',
  },
  {
    q: 'Vocês trabalham com convênio?',
    a: 'Não. Atendemos exclusivamente em formato particular. Despesas médicas são dedutíveis no Imposto de Renda como gasto com saúde, conforme legislação vigente.',
  },
  {
    q: 'Preciso levar exames anteriores?',
    a: 'Sim, sempre que possível. Quanto mais histórico tivermos, melhor a leitura — exames laboratoriais, de imagem, laudos antigos e medicamentos em uso atual.',
  },
  {
    q: 'A consulta inclui pedido de novos exames?',
    a: 'Sim, quando há indicação clínica. A solicitação fica com você para realizar no laboratório de sua escolha.',
  },
  {
    q: 'Qual a diferença entre uma Consulta Plenya e o Continuum Plenya?',
    a: 'A Consulta resolve um momento. O Continuum constrói trajetória — equipe multidisciplinar acompanhando junto, plano único e reavaliações em ciclo de 6 ou 12 meses.',
  },
];

export const metadata: Metadata = {
  title: 'Consultas Médicas | Plenya',
  description:
    'Consultas médicas particulares na Plenya — presenciais em Londrina ou online por telemedicina, com Dr. Getúlio Amaral e demais médicos da clínica.',
  alternates: { canonical: '/consultas' },
};

export default async function ConsultasPage({ params }: { params: Promise<{ locale: string }> }) {
  const { locale } = await params;
  setRequestLocale(locale);

  const doctors = await getAllDoctors();
  const getulio = doctors.find((d) => d.slug === 'getulio-amaral');
  const outros = doctors.filter((d) => d.category === 'medico' && d.slug !== 'getulio-amaral');

  return (
    <>
      <FaqSchema items={consultasFaq.map((f) => ({ q: f.q, a: typeof f.a === 'string' ? f.a : '' }))} />
      {/* Hero */}
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32 max-w-3xl">
          <p className="label-upper text-gold mb-6">Consultas</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,4.5rem)] text-cream">
            Consulta médica.
          </h1>
          <p className="text-cream/75 text-lg leading-relaxed mt-8 max-w-2xl">
            Saúde não é sobre reagir. É sobre antecipar. A consulta médica na Plenya
            é o ponto em que essa medicina começa — escuta longa, leitura cuidadosa
            dos dados e uma conduta clara, ainda que o exame esteja "normal".
          </p>
        </div>
      </section>

      {/* Sobre a consulta */}
      <section className="bg-cream">
        <div className="site-container section grid lg:grid-cols-[2fr_1fr] gap-16">
          <div className="space-y-6 text-petrol/85 text-lg leading-relaxed">
            <p className="label-upper text-gold">Sobre a consulta</p>
            <p>
              Reservamos tempo para escutar. Uma consulta médica na Plenya não cabe em
              quinze minutos — e não foi pensada para caber. Revisamos sua história,
              seus exames anteriores, seu sono, sua rotina, sua composição corporal,
              o que mudou desde a última vez que alguém olhou de fato.
            </p>
            <p>
              Trabalhamos com a leitura funcional dos exames: o que é normal de
              referência laboratorial e o que é ótimo do ponto de vista clínico
              raramente são a mesma coisa. Quando há indicação, solicitamos exames
              complementares e construímos uma conduta — medicação, ajuste de
              estilo de vida, encaminhamento à equipe multidisciplinar, ou todos
              eles juntos.
            </p>
            <p>
              Cada Consulta Plenya entrega leitura honesta dos seus dados,
              conduta clara e o que precisa entrar em prática agora — pode ser
              pontual ou se repetir no seu tempo, conforme a sua saúde pedir.
              Quando o cuidado pede continuidade — equipe acompanhando junto,
              plano único e revisões em ciclo — o próximo passo se chama{' '}
              <Link href="/continuum" className="underline decoration-gold/60 underline-offset-4 hover:text-petrol transition">
                Continuum Plenya
              </Link>
              .
            </p>
          </div>

          <aside className="space-y-8">
            <div className="space-y-3">
              <p className="label-upper text-gold">Modalidades</p>
              <p className="text-petrol/80 leading-relaxed">
                Presencial, em nossa clínica em Londrina (PR), ou online por
                telemedicina — com a mesma profundidade clínica.
              </p>
            </div>

            <div className="space-y-3">
              <p className="label-upper text-gold">Atendimento</p>
              <p className="text-petrol/80 leading-relaxed">
                Particular. Não trabalhamos com convênios.
              </p>
            </div>

            <div className="space-y-3">
              <p className="label-upper text-gold">Agendamento</p>
              <p className="text-petrol/80 leading-relaxed">
                Pela equipe da clínica, conforme disponibilidade do médico
                escolhido.
              </p>
            </div>
          </aside>
        </div>
      </section>

      {/* Modalidades — visual */}
      <section className="bg-paper">
        <div className="site-container section-sm">
          <div className="grid md:grid-cols-2 gap-6 md:gap-8">
            <figure className="space-y-3">
              <div className="relative aspect-[4/5] overflow-hidden">
                <Image
                  src="/images/getulio-consulta.jpg"
                  alt="Consulta médica presencial"
                  fill
                  className="object-cover"
                  sizes="(min-width: 768px) 540px, 100vw"
                />
              </div>
              <figcaption className="label-upper text-gold">Presencial · Londrina</figcaption>
            </figure>
            <figure className="space-y-3">
              <div className="relative aspect-[4/5] overflow-hidden">
                <Image
                  src="/images/getulio-consulta-online.jpg"
                  alt="Consulta médica online por telemedicina"
                  fill
                  className="object-cover"
                  sizes="(min-width: 768px) 540px, 100vw"
                />
              </div>
              <figcaption className="label-upper text-gold">Online · Telemedicina</figcaption>
            </figure>
          </div>
        </div>
      </section>

      {/* Quem atende */}
      <section className="bg-petrol text-cream">
        <div className="site-container section">
          <div className="grid lg:grid-cols-12 gap-12 mb-20">
            <div className="lg:col-span-5">
              <p className="label-upper text-gold mb-6">Quem atende</p>
              <h2 className="heading-section text-cream text-3xl md:text-5xl">
                Médicos da equipe Plenya.
              </h2>
            </div>
            <p className="lg:col-span-7 text-cream/75 text-lg leading-relaxed self-end max-w-2xl">
              Cada médico que integra a Plenya foi escolhido pelo alinhamento à
              mesma medicina — clínica rigorosa, antecipatória, integrada à
              equipe multidisciplinar. Você pode escolher o profissional com
              quem deseja consultar.
            </p>
          </div>

          <div className="grid lg:grid-cols-2 gap-12">
            {getulio && (
              <article className="border-t border-gold pt-8 space-y-5">
                <p className="label-upper text-gold">Direção clínica</p>
                <h3 className="heading-section text-cream text-3xl md:text-4xl">
                  {getulio.name}
                </h3>
                <p className="label-upper text-cream/50">
                  {getulio.credentials}
                  {getulio.rqe ? ` · RQE ${getulio.rqe}` : ''}
                </p>
                <p className="text-cream/75 leading-relaxed">
                  {getulio.shortBio}
                </p>
                <p>
                  <Link
                    href="/dr-getulio"
                    className="text-gold hover:text-cream transition border-b border-gold/40 pb-0.5"
                  >
                    Sobre Dr. Getúlio →
                  </Link>
                </p>
              </article>
            )}

            <article className="border-t border-cream/30 pt-8 space-y-5">
              <p className="label-upper text-cream/60">Equipe médica</p>
              <h3 className="heading-section text-cream text-3xl md:text-4xl">
                Demais médicos da clínica
              </h3>
              <p className="text-cream/75 leading-relaxed">
                Médicos formados em diferentes especialidades, todos com prática
                ancorada em medicina funcional integrativa e no Método AGIR.
                Atuam em conjunto com a equipe multidisciplinar — nutricionista,
                psicólogo e educador físico — para que cada plano terapêutico
                seja construído de forma coordenada.
              </p>
              {outros.length > 0 && (
                <ul className="space-y-3 pt-2">
                  {outros.map((d) => (
                    <li key={d.slug}>
                      <Link
                        href={`/equipe/${d.slug}`}
                        className="block group"
                      >
                        <span className="block text-cream group-hover:text-gold transition">
                          {d.name}
                        </span>
                        <span className="block text-cream/50 text-sm mt-0.5">
                          {d.role}
                        </span>
                      </Link>
                    </li>
                  ))}
                </ul>
              )}
              <p>
                <Link
                  href="/equipe"
                  className="text-gold hover:text-cream transition border-b border-gold/40 pb-0.5"
                >
                  Conhecer a equipe completa →
                </Link>
              </p>
            </article>
          </div>
        </div>
      </section>

      <ComparatorVsConvenio
        title="O que muda entre uma Consulta Plenya e a consulta convencional."
        bg="bg-paper"
      />

      <TestimonialsInline
        bg="bg-cream"
        label="Histórias Plenya"
        title="O que pacientes contam depois da consulta."
        limit={3}
      />

      <FaqAccordion title="Perguntas que recebemos toda semana sobre a Consulta Plenya." items={consultasFaq} />

      {/* Agendamento + cross-links */}
      <section className="bg-cream">
        <div className="site-container section grid md:grid-cols-12 gap-12">
          <div className="md:col-span-6 space-y-6">
            <p className="label-upper text-gold">Agendamento</p>
            <h2 className="heading-section text-petrol text-3xl md:text-4xl">
              Marcar uma consulta.
            </h2>
            <p className="text-petrol/75 leading-relaxed max-w-md">
              Entre em contato com a equipe da clínica para verificar
              disponibilidade e confirmar o agendamento.
            </p>
            <p>
              <Link href="/contato" className="btn-gold">Entrar em contato</Link>
            </p>
          </div>

          <div className="md:col-span-6 grid sm:grid-cols-2 gap-6 content-end">
            <Link href="/continuum" className="border-t border-petrol/15 pt-6 space-y-2 group">
              <p className="label-upper text-gold">Continuidade</p>
              <p className="heading-section text-petrol text-lg group-hover:text-gold transition">
                Continuum Plenya →
              </p>
            </Link>
            <Link href="/escore-plenya" className="border-t border-petrol/15 pt-6 space-y-2 group">
              <p className="label-upper text-gold">Diagnóstico</p>
              <p className="heading-section text-petrol text-lg group-hover:text-gold transition">
                Escore Plenya →
              </p>
            </Link>
          </div>
        </div>
      </section>
    </>
  );
}
