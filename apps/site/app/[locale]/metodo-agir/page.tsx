import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { agirLetters } from '@/lib/agir-structure';

export const metadata: Metadata = {
  title: 'Método AGIR',
  description:
    'O Método AGIR organiza o cuidado clínico Plenya em quatro pilares interdependentes — Atividade Física, Alimentação e Suplementação Inteligente; Gestão Clínica e Metabólica; Integração Mente-Corpo; Ritmo Circadiano e Repouso. Medicina antecipatória com profundidade.',
};

const pilares = [
  {
    code: 'A',
    name: 'Atividade Física, Alimentação e Suplementação Inteligente',
    territory: 'Os três motores que nenhuma medicação substitui',
    intro:
      'Comida não é combustível — é informação. Cada refeição instrui genes, modula hormônios e alimenta o microbioma. O exercício é o medicamento mais poderoso disponível. A suplementação corrige o que o laboratório mostra que falta. Os três motores trabalham juntos.',
    monitoramos: [
      'Composição corporal e VO₂ max',
      'Padrão alimentar e janela alimentar',
      'Força, flexibilidade e capacidade aeróbia',
      'Status de vitamina D, B12, ferro, ômega-3',
    ],
    quote:
      'A medicação cuida do que ela alcança. O resto — insulina, inflamação, sarcopenia, microbioma — depende do que você faz todos os dias.',
  },
  {
    code: 'G',
    name: 'Gestão Clínica e Metabólica',
    territory: 'O painel de controle interno',
    intro:
      'Hormônios, biomarcadores, farmacologia guiada por dados. A parte do trabalho em que o médico deixa de ser conselheiro de estilo de vida e passa a ser piloto de um painel de dados — testar, intervir, retestar, ajustar.',
    monitoramos: [
      'Painel hormonal completo (testosterona livre, SHBG, estradiol, DHEA-S, cortisol, tireoide)',
      'ApoB, lipoproteína(a), CAC quando indicado',
      'Insulina, HOMA-IR, HbA1c, peptídeo C',
      'PCR ultrassensível, homocisteína, ferritina',
      'Marcadores genéticos quando indicados — medicina de precisão',
    ],
    quote:
      'Pedir TSH isolado é como medir a pressão na torneira sem verificar se a água está saindo. A gestão metabólica começa em pedir o exame certo.',
  },
  {
    code: 'I',
    name: 'Integração Mente-Corpo',
    territory: 'O pilar que ninguém prescreve',
    intro:
      'O eixo entre psicologia, sistema imune e inflamação. Estados psicológicos alteram diretamente a função imune, hormonal e inflamatória — não como metáfora, como mecanismo. Quando o eixo HPA está em alerta crônico, nenhum outro pilar funciona plenamente.',
    monitoramos: [
      'Cortisol matinal e padrão diário',
      'PCR e marcadores inflamatórios',
      'Variabilidade da frequência cardíaca (HRV)',
      'Triagem de ansiedade, depressão e burnout',
      'Qualidade dos relacionamentos e suporte social',
    ],
    quote:
      'Ana otimizou tudo o que a bioquímica permitia. A PCR não descia. O que precisava de tratamento era a ansiedade que ela chamava de personalidade.',
  },
  {
    code: 'R',
    name: 'Ritmo Circadiano e Repouso',
    territory: 'O substrato sobre o qual os outros pilares operam',
    intro:
      'Sono, cronobiologia, sincronização dos relógios biológicos. O núcleo supraquiasmático é o maestro; cada célula tem seu relógio. Quando os sinais chegam em horários imprevisíveis, a orquestra desafina — e o nome disso é cronodisrupção.',
    monitoramos: [
      'Duração, eficiência e arquitetura do sono',
      'Sintomas e triagem para apneia obstrutiva',
      'Cronotipo e exposição à luz matinal',
      'Janela alimentar e cronoalimentação',
    ],
    quote:
      'Quando o substrato está comprometido, tudo que foi construído acima começa a rachar. O sono não é um pilar como os outros — é a base.',
  },
] as const;

const conviccoes = [
  {
    n: '01',
    text: 'O corpo não espera pelo diagnóstico para adoecer.',
    sub: 'A doença começa dez, quinze, vinte anos antes do rótulo. A janela de intervenção existe — mas ela fecha. E fecha em silêncio.',
  },
  {
    n: '02',
    text: 'Ninguém muda tudo ao mesmo tempo.',
    sub: 'Dois focos por trimestre, com medida e prazo. Quem tenta mudar dez coisas na segunda-feira geralmente não mudou nenhuma na sexta.',
  },
  {
    n: '03',
    text: 'Os pilares não funcionam isolados.',
    sub: 'Alimentação sem sono é desgaste. Exercício sem recuperação é overtraining. Hormônio sem saúde mental é otimização sobre areia.',
  },
  {
    n: '04',
    text: 'A parceria importa.',
    sub: 'Nenhum paciente faz sozinho. Medicina preventiva é trabalho em rede — médico, equipe multidisciplinar e o paciente no centro.',
  },
  {
    n: '05',
    text: 'O melhor momento para começar era dez anos atrás. O segundo melhor é hoje.',
    sub: 'Não é frase motivacional. É dado clínico. A reversibilidade diminui com o tempo. O custo da inação aumenta com a idade.',
  },
];

export default async function MethodPage({ params }: { params: Promise<{ locale: string }> }) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <>
      {/* Hero — fotográfico */}
      <section className="relative bg-petrol text-cream overflow-hidden">
        <Image
          src="/images/hero-about.jpg"
          alt=""
          fill
          priority
          className="object-cover opacity-40"
          sizes="100vw"
        />
        <div className="absolute inset-0 bg-gradient-to-b from-petrol/70 via-petrol/60 to-petrol" />
        <div className="relative site-container pt-32 pb-24 md:pt-40 md:pb-32">
          <p className="label-upper text-gold mb-6">Método</p>
          <h1 className="heading-hero text-[clamp(2.8rem,7vw,5.5rem)] text-cream max-w-3xl">
            AGIR — quatro pilares para uma saúde que se sustenta no tempo.
          </h1>
          <p className="text-cream/80 text-lg md:text-xl mt-8 max-w-2xl leading-relaxed">
            O framework clínico Plenya organiza o cuidado em quatro dimensões interdependentes da
            saúde. Não é protocolo. É a forma como decidimos, monitoramos e acompanhamos —
            com a profundidade que a longevidade exige.
          </p>
        </div>
      </section>

      {/* Premissa — por que AGIR */}
      <section className="bg-cream">
        <div className="site-container section">
          <div className="grid lg:grid-cols-[1fr_2fr] gap-16 items-start">
            <p className="label-upper text-gold lg:pt-3">Premissa</p>
            <div className="space-y-6 text-petrol/85 text-lg leading-relaxed">
              <p className="heading-section text-petrol text-2xl md:text-4xl leading-tight">
                O corpo não espera pelo diagnóstico para adoecer.
              </p>
              <p>
                Quando a glicemia sobe, a resistência insulínica já está há uma década.
                Quando o infarto acontece, a placa cresce há vinte anos. A janela de intervenção
                existe — mas fecha em silêncio.
              </p>
              <p>
                O Método AGIR é a forma de operar dentro dessa janela. Quatro pilares
                interdependentes que, juntos, organizam o cuidado preventivo com o mesmo
                rigor que se aplica à medicina hospitalar — exames, dados, decisão estruturada —
                voltado para o que vem antes da doença.
              </p>
            </div>
          </div>
        </div>
      </section>

      {/* Os 4 pilares — em profundidade */}
      <section className="bg-paper">
        <div className="site-container section space-y-24 md:space-y-32">
          {pilares.map((pilar, idx) => (
            <article
              key={pilar.code}
              className="grid lg:grid-cols-[180px_1fr] gap-8 lg:gap-16 border-t border-petrol/15 pt-12"
            >
              <div className="space-y-3">
                <p className="label-upper text-gold">Pilar 0{idx + 1}</p>
                <span className="heading-hero text-gold text-[clamp(6rem,14vw,11rem)] leading-[0.85] block">
                  {pilar.code}
                </span>
              </div>

              <div className="space-y-10">
                <header className="space-y-3">
                  <h2 className="heading-section text-petrol text-3xl md:text-4xl lg:text-5xl">
                    {pilar.name}
                  </h2>
                  <p className="label-upper text-petrol/55">{pilar.territory}</p>
                </header>

                <p className="text-petrol/85 text-lg leading-relaxed max-w-2xl">{pilar.intro}</p>

                <blockquote className="border-l-2 border-gold pl-6 py-2 text-petrol/80 text-lg italic max-w-2xl">
                  "{pilar.quote}"
                </blockquote>

                {(() => {
                  const letter = agirLetters.find((l) => l.code === pilar.code);
                  if (!letter) return null;
                  return (
                    <div className="space-y-4">
                      <p className="label-upper text-petrol/55">
                        {letter.pillarCount} pilares clínicos
                      </p>
                      <div className="space-y-5">
                        {letter.groups.map((group, gi) => (
                          <div key={gi} className="space-y-2">
                            {group.label && (
                              <p className="label-upper text-petrol/40 text-[10px]">
                                {group.label}
                              </p>
                            )}
                            <div className="flex flex-wrap gap-2">
                              {group.pillars.map((sp) => (
                                <span
                                  key={sp}
                                  className="inline-flex items-center px-3.5 py-1.5 rounded-full border border-petrol/20 text-petrol/85 text-[13px]"
                                >
                                  {sp}
                                </span>
                              ))}
                            </div>
                          </div>
                        ))}
                      </div>
                    </div>
                  );
                })()}

                <div className="space-y-4">
                  <p className="label-upper text-petrol/55">O que monitoramos · exemplos</p>
                  <ul className="grid sm:grid-cols-2 gap-x-10 gap-y-3 max-w-3xl">
                    {pilar.monitoramos.map((item) => (
                      <li key={item} className="flex gap-3 text-petrol/80 text-base">
                        <span className="text-gold mt-1.5 leading-none">—</span>
                        <span className="leading-relaxed">{item}</span>
                      </li>
                    ))}
                  </ul>
                  <p className="text-petrol/50 text-sm pt-2 max-w-3xl">
                    Lista parcial — o monitoramento real cobre dezenas de marcadores adicionais
                    por pilar, ajustados ao seu caso.
                  </p>
                </div>
              </div>
            </article>
          ))}
        </div>
      </section>

      {/* As 5 convicções */}
      <section className="bg-petrol text-cream">
        <div className="site-container section">
          <div className="max-w-2xl space-y-4 mb-16">
            <p className="label-upper text-gold">Convicções</p>
            <h2 className="heading-section text-cream text-3xl md:text-5xl">
              Cinco frases que levaram uma década para virar convicção.
            </h2>
          </div>

          <div className="border-t border-cream/15">
            {conviccoes.map((c) => (
              <div
                key={c.n}
                className="grid md:grid-cols-[80px_1fr] gap-6 md:gap-12 py-10 border-b border-cream/10"
              >
                <span className="label-upper text-gold pt-1">{c.n}</span>
                <div className="space-y-3 max-w-3xl">
                  <p className="heading-section text-cream text-2xl md:text-3xl leading-tight">
                    {c.text}
                  </p>
                  <p className="text-cream/70 leading-relaxed">{c.sub}</p>
                </div>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Como aplicamos */}
      <section className="bg-cream">
        <div className="site-container section grid lg:grid-cols-2 gap-16 items-start">
          <div className="space-y-6">
            <p className="label-upper text-gold">Como aplicamos</p>
            <h2 className="heading-section text-petrol text-3xl md:text-4xl">
              Aplicado por uma equipe. Mensurado pelo Escore Plenya.
            </h2>
            <p className="text-petrol/80 text-lg leading-relaxed">
              O Método AGIR é o framework que organiza nossas decisões clínicas e nossa comunicação
              com você. O Escore Plenya traduz o estado atual em uma pontuação que evolui ao longo
              do acompanhamento — para que o progresso seja claro, mensurável e compartilhado.
            </p>
            <p className="text-petrol/80 text-lg leading-relaxed">
              A regra dos dois opera aqui: começamos por dois focos, com critério clínico, e
              avançamos no ritmo do seu corpo — não em listas de boas intenções.
            </p>
          </div>

          <div className="space-y-4 lg:pt-12">
            <Link href="/escore-plenya" className="btn-outline-dark w-full sm:w-auto">
              Conhecer o Escore Plenya
            </Link>
            <Link href="/continuum" className="btn-gold w-full sm:w-auto sm:ml-3">
              Conhecer o Continuum Plenya
            </Link>
          </div>
        </div>
      </section>

      {/* Cross-links */}
      <section className="bg-paper">
        <div className="site-container section grid md:grid-cols-3 gap-8">
          <Link href="/dr-getulio" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">Direção Clínica</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">Dr. Getúlio Amaral Filho →</p>
          </Link>
          <Link href="/equipe" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">Equipe</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">Quem aplica o Método →</p>
          </Link>
          <Link href="/a-plenya" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">Plenya</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">Manifesto e propósito →</p>
          </Link>
        </div>
      </section>
    </>
  );
}
