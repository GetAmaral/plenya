import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { RadarAgir } from '@/components/escore/RadarAgir';

export const metadata: Metadata = {
  title: 'Escore Plenya — instrumento de medida do Método AGIR',
  description:
    'Mais de 800 itens — história, sintomas, exames, hábitos e medicamentos — traduzidos em uma pontuação clara, evolutiva e personalizada. A medida única do seu estado de saúde.',
};

const steps = [
  {
    n: '01',
    title: 'Avaliação ampla',
    body:
      'Mais de 800 itens são levantados — história clínica e familiar, sintomas atuais, exames laboratoriais, hábitos de vida, medicamentos em uso, marcadores genéticos. Não há pergunta solta: cada item alimenta uma decisão clínica.',
  },
  {
    n: '02',
    title: 'Organização AGIR',
    body:
      'Cada item entra em um pilar clínico do Método AGIR — agrupado em uma das quatro letras: Atividade Física, Alimentação e Suplementação Inteligente; Gestão Clínica e Metabólica; Integração Mente-Corpo; Ritmo Circadiano e Repouso.',
  },
  {
    n: '03',
    title: 'Pontuação que evolui',
    body:
      'Você recebe uma pontuação global e uma por letra. A cada reavaliação, a curva mostra o que melhorou, o que estagnou e onde a próxima intervenção precisa entrar.',
  },
];

const tiers: Array<{ title: string; anchor: string; href?: string; desc: string }> = [
  {
    title: 'Triagem',
    anchor: 'Faça agora — gratuito',
    href: '/escore-plenya/avaliar',
    desc: 'Versão pública e gratuita. Cerca de 35 perguntas respondidas em poucos minutos — sem cadastro. Indica, em primeira leitura, onde a sua saúde está e qual o próximo passo para aprofundar.',
  },
  {
    title: 'Painel',
    anchor: 'Painel ampliado',
    href: '/escore-plenya/painel',
    desc: 'Versão estendida com mais de 80 itens, incluindo exames laboratoriais, imagem e medidas avançadas de composição corporal. Disponibilizada como bônus em ações específicas — preparação útil antes da consulta.',
  },
  {
    title: 'Consulta',
    anchor: 'Consulta Plenya',
    href: '/consultas',
    desc: 'Aplicada pelo médico na consulta avulsa, presencial ou online. Recorte focado nas metas prioritárias identificadas na conduta clínica, com base no painel ampliado de exames.',
  },
  {
    title: 'Continuum',
    anchor: 'Continuum Plenya',
    href: '/continuum',
    desc: 'Versão completa, aplicada pela equipe Plenya ao longo do programa de acompanhamento contínuo. Avaliação clínica, laboratorial, comportamental e funcional integrais. Relatório detalhado com metas e plano personalizado — reavaliado a cada ciclo.',
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
            O instrumento de medida do Método AGIR. Mais de 800 itens — história,
            sintomas, exames, hábitos e medicamentos — em uma pontuação clara,
            evolutiva e personalizada do seu estado de saúde.
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
                <p className="text-gold text-5xl font-light leading-none tabular-nums">{s.n}</p>
                <h3 className="heading-section text-petrol text-xl">{s.title}</h3>
                <p className="text-petrol/75 leading-relaxed">{s.body}</p>
              </li>
            ))}
          </ol>
        </div>
      </section>

      {/* PONTUAÇÃO — radar AGIR como peça visual central */}
      <section className="bg-cream">
        <div className="site-container section grid lg:grid-cols-[1fr_auto] gap-16 items-center">
          <div className="space-y-6 max-w-xl">
            <p className="label-upper text-gold">Pontuação</p>
            <h2 className="heading-section text-petrol text-3xl md:text-5xl leading-tight">
              Não é uma nota. <em className="not-italic text-gold">É um mapa.</em>
            </h2>
            <p className="text-petrol/80 text-lg leading-relaxed">
              Você recebe uma pontuação global e quatro pontuações por letra do AGIR.
              Juntas, elas formam uma constelação — onde está concentrada a força,
              onde mora a lacuna.
            </p>
            <p className="text-petrol/70 leading-relaxed">
              Quando o radar é simétrico, o cuidado está distribuído. Quando uma ponta
              afunda, é ali que a próxima intervenção entra.
            </p>
          </div>

          <RadarAgir />
        </div>
      </section>

      {/* FAIXAS — legenda semântica */}
      <section className="bg-paper border-y border-petrol/10">
        <div className="site-container py-16 md:py-20">
          <div className="max-w-3xl mb-10 space-y-3">
            <p className="label-upper text-gold">Como ler</p>
            <h2 className="heading-section text-petrol text-2xl md:text-3xl">
              Cinco faixas. Direção, não punição.
            </h2>
          </div>

          <div className="grid sm:grid-cols-2 lg:grid-cols-5 gap-3 md:gap-4 max-w-5xl">
            {[
              { label: 'Crítico',  range: '0–30',   color: 'bg-[#c1542d]' },
              { label: 'Atenção',  range: '31–50',  color: 'bg-[#d89345]' },
              { label: 'Regular',  range: '51–70',  color: 'bg-[#caa54a]' },
              { label: 'Bom',      range: '71–85',  color: 'bg-[#7ba07a]' },
              { label: 'Ótimo',    range: '86–100', color: 'bg-[#3f7d6e]' },
            ].map((f) => (
              <div key={f.label} className="flex items-center gap-3 p-4 border border-petrol/10 rounded-sm">
                <span className={`w-3 h-10 rounded-sm ${f.color} flex-shrink-0`} aria-hidden />
                <div>
                  <p className="heading-section text-petrol text-base leading-tight">{f.label}</p>
                  <p className="font-mono text-xs text-petrol/50 mt-0.5">{f.range}</p>
                </div>
              </div>
            ))}
          </div>

          <p className="text-petrol/60 leading-relaxed mt-8 max-w-3xl">
            A faixa não é veredito — é o ponto de partida da próxima conversa clínica.
            Um <strong className="text-petrol">71% em queda</strong> de 85% conta uma história
            diferente de um <strong className="text-petrol">71% subindo</strong> de 60%.
          </p>
        </div>
      </section>

      {/* EVOLUÇÃO — gráfico simbólico */}
      <section className="bg-cream">
        <div className="site-container section grid lg:grid-cols-[1fr_auto] gap-16 items-center">
          <div className="space-y-6 max-w-xl">
            <p className="label-upper text-gold">Evolução</p>
            <h2 className="heading-section text-petrol text-3xl md:text-4xl">
              A curva que mostra se o cuidado está funcionando.
            </h2>
            <p className="text-petrol/80 text-lg leading-relaxed">
              A cada reavaliação, um novo ponto entra na curva. O médico, a equipe e
              você veem o mesmo gráfico: progresso real, estagnação ou queda — sem
              espaço para opinião.
            </p>
          </div>

          <figure className="flex flex-col items-center gap-3">
            <svg viewBox="0 0 420 220" className="w-full max-w-md md:max-w-lg" aria-label="Exemplo de curva de evolução do Escore Plenya">
              {/* Grid horizontal */}
              {[40, 90, 140, 190].map((y) => (
                <line key={y} x1="40" y1={y} x2="400" y2={y} stroke="#063b4f" strokeOpacity="0.06" strokeWidth="1" />
              ))}
              {/* Eixo Y labels */}
              {[
                { y: 40, v: '100' },
                { y: 90, v: '75' },
                { y: 140, v: '50' },
                { y: 190, v: '25' },
              ].map((t) => (
                <text key={t.v} x="32" y={t.y + 4} textAnchor="end" fontFamily="monospace" fontSize="10" fill="#063b4f" fillOpacity="0.4">{t.v}</text>
              ))}
              {/* Linha de tendência exemplo */}
              <polyline
                points="60,160 140,140 220,110 300,85 380,68"
                fill="none"
                stroke="#b38645"
                strokeWidth="2.5"
                strokeLinecap="round"
                strokeLinejoin="round"
              />
              {/* Pontos */}
              {[
                { x: 60, y: 160 },
                { x: 140, y: 140 },
                { x: 220, y: 110 },
                { x: 300, y: 85 },
                { x: 380, y: 68 },
              ].map((p, i) => (
                <circle key={i} cx={p.x} cy={p.y} r="4" fill="#063b4f" stroke="#fbfaf6" strokeWidth="2" />
              ))}
              {/* Eixo X labels */}
              {['T0', '3m', '6m', '9m', '12m'].map((label, i) => (
                <text key={label} x={60 + i * 80} y="210" textAnchor="middle" fontFamily="monospace" fontSize="10" fill="#063b4f" fillOpacity="0.4">{label}</text>
              ))}
            </svg>
            <p className="label-upper text-petrol/40 text-center text-[10px]">
              Exemplo · 12 meses de acompanhamento
            </p>
          </figure>
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
              Quatro níveis de profundidade.
            </h2>
            <p className="text-petrol/70 leading-relaxed">
              Da triagem pública gratuita ao acompanhamento contínuo aplicado pela equipe.
              Cada nível é uma porta de entrada — escolha pelo momento em que você está.
            </p>
          </div>
          <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-8">
            {tiers.map((tier, i) => (
              <div key={tier.title} className="border-t border-petrol/15 pt-8 space-y-3">
                <span className="label-upper text-petrol/40">0{i + 1}</span>
                <h3 className="heading-section text-petrol text-2xl">{tier.title}</h3>
                {tier.href ? (
                  <Link
                    href={tier.href}
                    className="inline-flex items-center gap-1 text-gold text-base font-medium underline decoration-gold/50 underline-offset-4 hover:decoration-gold transition"
                  >
                    {tier.anchor}
                    <span aria-hidden className="transition-transform group-hover:translate-x-1">→</span>
                  </Link>
                ) : (
                  <p className="text-petrol/50 text-base italic">{tier.anchor}</p>
                )}
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
