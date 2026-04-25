import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';

export const metadata: Metadata = {
  title: 'Casos clínicos — leitura honesta de pacientes Plenya',
  description:
    'Casos clínicos anonimizados, com consentimento. Como uma escuta longa, uma leitura funcional dos exames e um plano integrado mudam o curso da saúde.',
};

type Caso = {
  slug: string;
  titulo: string;
  resumo: string;
  paciente: string;
  contexto: string;
  achados: string[];
  conduta: string[];
  evolucao: string[];
  tempo: string;
  pillar: 'A' | 'G' | 'I' | 'R';
};

const casos: Caso[] = [
  {
    slug: 'mulher-50-fadiga',
    titulo: 'Mulher, 52 anos — “estou cansada e ninguém acha nada”.',
    resumo:
      'Painel laboratorial dentro da “normalidade”, mas vários marcadores distantes do ótimo funcional. O que parecia idade era estoque baixo somado a sono fragmentado.',
    paciente: 'Mulher · 52 anos · profissional liberal',
    contexto:
      'Procurou a Plenya após dois check-ups anuais que retornaram “tudo normal”. Queixa principal: fadiga persistente, queda de libido, intestino lento e sensação de que “a cabeça não responde igual”.',
    achados: [
      'Ferritina 28 ng/mL (referência > 15, ótimo funcional > 70)',
      'TSH 3,8 µUI/mL (referência < 4,5, ótimo < 2,5)',
      'Vitamina D 24 ng/mL (referência > 20, ótimo 40–60)',
      'Sono fragmentado documentado por wearable — 3 despertares/noite',
    ],
    conduta: [
      'Reposição de ferro com nutricionista, ajustada por absorção',
      'Suplementação de vitamina D em dose calculada por peso',
      'Higiene de sono guiada pela psicóloga + revisão da janela noturna',
      'Reavaliação laboratorial em 12 semanas',
    ],
    evolucao: [
      'Ferritina 78 ng/mL em 4 meses',
      'Vitamina D 52 ng/mL no mesmo período',
      'Despertares noturnos: < 1/noite',
      'Plenya Score: 54 → 78 em 6 meses',
    ],
    tempo: '6 meses · Continuum Anual',
    pillar: 'G',
  },
  {
    slug: 'homem-45-pre-diabetes',
    titulo: 'Homem, 45 anos — pré-diabetes, sem sintomas.',
    resumo:
      'HbA1c em 6,3 detectada em check-up de rotina. Sem queixa, sem peso a perder no espelho. O Continuum reverteu o curso antes da medicação entrar.',
    paciente: 'Homem · 45 anos · executivo',
    contexto:
      'Trouxe exame de empresa com HbA1c 6,3% e glicemia de jejum 108 mg/dL. IMC 26,5. Treinava 2x/semana, dieta “razoável”. Pai diabético tipo 2.',
    achados: [
      'HbA1c 6,3% (pré-diabetes; referência < 5,7)',
      'Glicemia de jejum 108 mg/dL',
      'Triglicérides 188 mg/dL · HDL 38 mg/dL',
      'Composição: 28% gordura corporal',
      'Sono: 5h30 em média, com apneia leve em polissonografia',
    ],
    conduta: [
      'Reorganização alimentar com nutricionista (carga glicêmica + proteína)',
      'Treino de força 3x/semana com educador físico',
      'CPAP indicado e ajustado',
      'Sem medicação oral — decisão compartilhada com paciente',
    ],
    evolucao: [
      'HbA1c 5,4% em 6 meses',
      'Triglicérides 92 mg/dL · HDL 47 mg/dL',
      'Composição: 22% de gordura corporal',
      'Plenya Score: 49 → 81 em 9 meses',
    ],
    tempo: '9 meses · Continuum Anual',
    pillar: 'A',
  },
  {
    slug: 'mulher-38-burnout',
    titulo: 'Mulher, 38 anos — burnout que ninguém nomeou.',
    resumo:
      'Quadro de exaustão extrema atribuído a “estresse”. A leitura integrada mostrou perfil hormonal típico de carga alostática crônica.',
    paciente: 'Mulher · 38 anos · liderança em tecnologia',
    contexto:
      'Chegou após oito meses sem férias completas, alimentação irregular, episódios de palpitação noturna e sensação de que “tudo era demais”. Diagnóstico prévio de ansiedade tratada com escitalopram.',
    achados: [
      'Cortisol matinal alto, perda do ritmo circadiano',
      'DHEA-S na faixa baixa para a idade',
      'Magnésio sérico 1,7 mg/dL (limite inferior)',
      'Sono: latência > 40 min, microdespertares frequentes',
    ],
    conduta: [
      'Plano integrado: psicóloga (cognitivo-comportamental) + médico',
      'Reorganização do ritmo (luz matinal, jantar mais cedo, tela cortada às 22h)',
      'Reposição direcionada — magnésio glicinato + ajuste alimentar',
      'Reavaliação do escitalopram em conjunto com psiquiatra de referência',
    ],
    evolucao: [
      'Cortisol matinal normalizado em 4 meses',
      'Latência de sono < 15 min',
      'Retomada de exercício regular (2x/semana → 4x/semana)',
      'Plenya Score: 41 → 72 em 7 meses',
    ],
    tempo: '7 meses · Continuum Semestral renovado',
    pillar: 'I',
  },
];

const pillarLabels: Record<Caso['pillar'], string> = {
  A: 'Atividade · Alimentação',
  G: 'Gestão clínica e metabólica',
  I: 'Integração mente-corpo',
  R: 'Ritmo circadiano e repouso',
};

export default async function CasosPage({ params }: { params: Promise<{ locale: string }> }) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <>
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32 max-w-3xl">
          <p className="label-upper text-gold mb-6">Casos clínicos</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,4.5rem)] text-cream">
            Como a medicina Plenya muda o curso.
          </h1>
          <p className="text-cream/75 text-lg leading-relaxed mt-8 max-w-2xl">
            Casos reais, anonimizados, com consentimento por escrito. Mostram em achado, conduta e
            evolução o que diferencia uma escuta longa de uma consulta corrida — e por que medir
            de novo importa tanto quanto medir uma vez.
          </p>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-container section space-y-20">
          {casos.map((c, i) => (
            <article key={c.slug} className="grid lg:grid-cols-12 gap-10 border-t border-petrol/15 pt-10">
              <header className="lg:col-span-4 space-y-4">
                <p className="label-upper text-gold">Caso 0{i + 1} · {pillarLabels[c.pillar]}</p>
                <h2 className="heading-section text-petrol text-2xl md:text-3xl leading-tight">
                  {c.titulo}
                </h2>
                <p className="text-petrol/70 leading-relaxed">{c.resumo}</p>
                <p className="label-upper text-petrol/55 pt-2 border-t border-petrol/10">{c.paciente}</p>
                <p className="label-upper text-gold">{c.tempo}</p>
              </header>

              <div className="lg:col-span-8 grid md:grid-cols-3 gap-8">
                <div className="space-y-3">
                  <p className="label-upper text-petrol/55">Contexto</p>
                  <p className="text-petrol/80 leading-relaxed text-sm">{c.contexto}</p>
                </div>
                <div className="space-y-3">
                  <p className="label-upper text-petrol/55">Achados</p>
                  <ul className="space-y-2">
                    {c.achados.map((a) => (
                      <li key={a} className="flex gap-2 text-petrol/80 text-sm leading-relaxed">
                        <span className="text-gold mt-1.5 leading-none">—</span>
                        <span>{a}</span>
                      </li>
                    ))}
                  </ul>
                </div>
                <div className="space-y-3">
                  <p className="label-upper text-petrol/55">Conduta</p>
                  <ul className="space-y-2">
                    {c.conduta.map((a) => (
                      <li key={a} className="flex gap-2 text-petrol/80 text-sm leading-relaxed">
                        <span className="text-gold mt-1.5 leading-none">—</span>
                        <span>{a}</span>
                      </li>
                    ))}
                  </ul>
                </div>

                <div className="md:col-span-3 bg-paper p-6 md:p-8 mt-2 space-y-3">
                  <p className="label-upper text-gold">Evolução</p>
                  <ul className="grid sm:grid-cols-2 gap-x-8 gap-y-2">
                    {c.evolucao.map((a) => (
                      <li key={a} className="flex gap-2 text-petrol/85 text-sm leading-relaxed">
                        <span className="text-gold mt-1.5 leading-none">→</span>
                        <span>{a}</span>
                      </li>
                    ))}
                  </ul>
                </div>
              </div>
            </article>
          ))}
        </div>
      </section>

      <section className="bg-paper">
        <div className="site-container section max-w-3xl space-y-4">
          <p className="label-upper text-petrol/55">Nota de transparência</p>
          <p className="text-petrol/75 leading-relaxed text-sm">
            Todos os casos publicados nesta página são reais, anonimizados e divulgados após
            consentimento por escrito do paciente, em conformidade com a LGPD e com o Código de
            Ética Médica. Achados laboratoriais foram preservados em fidelidade clínica;
            elementos identificadores (profissão, cidade, idade exata, contexto familiar) podem
            ter sido alterados para preservar a privacidade. Resultados individuais variam.
          </p>
        </div>
      </section>

      <section className="bg-petrol text-cream">
        <div className="site-container section text-center space-y-6 max-w-2xl mx-auto">
          <p className="label-upper text-gold">Quer entender o seu caso?</p>
          <p className="heading-section text-cream text-2xl md:text-3xl">
            Comece pela Triagem ou converse com a equipe.
          </p>
          <div className="flex flex-wrap justify-center gap-4 pt-2">
            <Link href="/escore-plenya/avaliar" className="btn-gold">Fazer a Triagem</Link>
            <Link href="/contato" className="btn-outline-dark border-cream/40 text-cream">
              Falar com a equipe
            </Link>
          </div>
        </div>
      </section>
    </>
  );
}
