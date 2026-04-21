import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';

export const metadata: Metadata = {
  title: 'Continuum Plenya — programa de acompanhamento contínuo',
  description:
    'Continuum Plenya é o programa semestral ou anual da Plenya. Acompanhamento 100% online com equipe multidisciplinar — médico, nutricionista, psicólogo e educador físico — organizado pelo Método AGIR e mensurado pelo Escore Plenya.',
};

const cobertura = [
  { code: 'A', name: 'Atividade, alimentação e suplementação', quem: 'Nutricionista + educador físico' },
  { code: 'G', name: 'Gestão clínica e metabólica', quem: 'Médico' },
  { code: 'I', name: 'Integração mente-corpo', quem: 'Psicólogo' },
  { code: 'R', name: 'Ritmo circadiano e repouso', quem: 'Atravessa os quatro, sob coordenação clínica' },
] as const;

const jornada = [
  {
    fase: 'Semana 1',
    titulo: 'Quatro consultas online.',
    body:
      'Uma consulta online com cada um dos quatro profissionais, em momentos diferentes da semana — iniciando pelo médico. Juntas, elas preenchem o Escore Plenya, que define o ponto de partida — e a direção. Quando há indicação, exames complementares são solicitados.',
  },
  {
    fase: 'Semana 2',
    titulo: 'Reunião da equipe e abertura do plano.',
    body:
      'Os quatro profissionais se reúnem, discutem o caso e montam um plano inicial integrado, alinhado aos seus objetivos. Em seguida, em um encontro online de abertura, o médico devolve a leitura clínica e apresenta o plano completo. Ninguém prescreve isolado.',
  },
  {
    fase: 'Semanas seguintes',
    titulo: 'Encontro semanal, em rotação.',
    body:
      'A cada semana, um encontro online com um dos quatro profissionais. A cada quatro semanas, o ciclo se completa — todas as áreas do AGIR são revisadas e reforçadas mês a mês.',
  },
  {
    fase: 'Conforme o plano',
    titulo: 'Box Plenya.',
    body:
      'Após a elaboração do plano, um box personalizado chega até você — mimos selecionados pela equipe e suplementos ou manipulados específicos do seu protocolo. À medida que o cuidado evolui e o plano se ajusta, novos boxes seguem o mesmo ritmo — preparados sempre que houver indicação.',
  },
  {
    fase: 'A cada três meses',
    titulo: 'Reavaliação do Escore.',
    body:
      'Exames de acompanhamento conforme indicação clínica e nova rodada do Escore Plenya. A curva mostra o que avançou, o que estagnou e onde a próxima intervenção entra.',
  },
  {
    fase: 'Final do ciclo',
    titulo: 'Avaliação final e próximo horizonte.',
    body:
      'Ao fim de seis ou doze meses, uma avaliação estruturada fecha o período. A partir dela, decidimos juntos o desenho do próximo ciclo.',
  },
];

const modalidades = [
  {
    name: 'Continuum Semestral',
    period: '6 meses',
    body: 'Ciclo completo: avaliação inicial, plano integrado, encontros semanais, reavaliação trimestral e fechamento.',
  },
  {
    name: 'Continuum Anual',
    period: '12 meses',
    body: 'Mesma estrutura, com mais uma reavaliação trimestral e maior consolidação dos hábitos no tempo.',
    highlight: true,
  },
];

const entra = [
  'Quatro consultas online iniciais — uma com cada profissional',
  'Reunião da equipe e construção do plano integrado',
  'Encontros semanais online durante todo o período',
  'Box Plenya na abertura e sempre que houver indicação — mimos personalizados, suplementos e manipulados',
  'Reavaliações periódicas do Escore e dos objetivos',
  'Avaliação final estruturada',
  'Grupo de WhatsApp com os quatro profissionais — para dúvidas em horário comercial',
];

const naoEntra: React.ReactNode[] = [
  'Exames laboratoriais e de imagem — pedidos pela equipe Plenya, realizados no laboratório de sua escolha',
  <>
    Atendimento presencial em Londrina — caminho avulso em{' '}
    <Link
      href="/consultas"
      className="underline decoration-gold/60 underline-offset-4 hover:text-petrol transition"
    >
      Consultas Plenya
    </Link>
  </>,
  'Urgências e emergências — o grupo de WhatsApp atende dúvidas, não pronto-atendimento',
];

export default async function ContinuumPage({ params }: { params: Promise<{ locale: string }> }) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <>
      {/* HERO — compacto, padrão Escore/AGIR */}
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32">
          <p className="label-upper text-gold mb-6">Programa de acompanhamento contínuo</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,5rem)] text-cream max-w-3xl">
            Continuum Plenya
          </h1>
          <p className="text-cream/80 text-lg md:text-xl mt-8 max-w-2xl leading-relaxed">
            Saúde não se resolve em uma consulta. Se constrói no tempo —
            em conjunto, com método. Seis ou doze meses conduzidos pela{' '}
            <Link href="/equipe" className="underline decoration-gold/60 underline-offset-4 hover:text-cream/100 transition">
              Equipe Plenya
            </Link>
            , guiados pelo{' '}
            <Link href="/escore-plenya" className="underline decoration-gold/60 underline-offset-4 hover:text-cream/100 transition">
              Escore Plenya
            </Link>
            , organizados pelo{' '}
            <Link href="/metodo-agir" className="underline decoration-gold/60 underline-offset-4 hover:text-cream/100 transition">
              Método AGIR
            </Link>
            .
          </p>
        </div>
      </section>

      {/* SOBRE O PROGRAMA */}
      <section className="bg-cream">
        <div className="site-container section">
          <div className="max-w-3xl space-y-6">
            <p className="label-upper text-gold">Sobre o programa</p>
            <p className="heading-section text-petrol text-2xl md:text-4xl leading-tight">
              Há quem procure saúde quando algo já apareceu. O Continuum Plenya
              existe para quem decidiu o caminho contrário.
            </p>
            <p className="text-petrol/80 text-lg leading-relaxed">
              Uma consulta avulsa resolve um momento. O Continuum constrói trajetória.
              Vinte anos de consultório ensinaram uma verdade simples: viver mais não
              é viver bem. Saúde se constrói no tempo — em conjunto, com método e com
              presença — ou não se constrói.
            </p>
            <p className="text-petrol/75 leading-relaxed">
              Tudo começa por uma leitura ampla — o Escore Plenya — e segue em ritmo
              contínuo: encontro semanal com um dos quatro profissionais, ciclo que se
              completa a cada quatro semanas, reavaliações trimestrais e fechamento
              estruturado ao fim do período. Nenhum pilar fica sem revisão.
            </p>
          </div>
        </div>
      </section>

      {/* PARA QUEM — Reconhecimento (mantém) */}
      <section className="bg-paper">
        <div className="site-container section">
          <div className="grid lg:grid-cols-12 gap-12 mb-16">
            <div className="lg:col-span-5">
              <p className="label-upper text-gold mb-6">Reconhecimento</p>
              <h2 className="heading-section text-petrol text-3xl md:text-5xl">
                O Continuum Plenya<br />não é para todo mundo.
              </h2>
            </div>
            <p className="lg:col-span-7 text-petrol/75 text-lg leading-relaxed self-end max-w-2xl">
              É para quem reconhece um destes momentos. Não como diagnóstico —
              como o ponto da vida em que faz sentido construir saúde no tempo,
              em conjunto, com método.
            </p>
          </div>

          <div className="grid sm:grid-cols-2 lg:grid-cols-3 gap-x-12 gap-y-12">
            <div className="space-y-3">
              <h3 className="heading-section text-petrol text-xl">
                Faz tudo certo — e os sinais não somam.
              </h3>
              <p className="text-petrol/75 leading-relaxed">
                Treina, come, dorme. Mesmo assim a disposição não volta, o
                intestino bagunça, a pele perde viço, três viroses no ano.
                Sintomas pequenos que, juntos, descrevem um terreno em desgaste.
              </p>
            </div>

            <div className="space-y-3">
              <h3 className="heading-section text-petrol text-xl">
                Recebeu o susto.
              </h3>
              <p className="text-petrol/75 leading-relaxed">
                Um exame que assustou, um diagnóstico de família, um aniversário
                que doeu. A janela parecia distante e ficou perto demais.
              </p>
            </div>

            <div className="space-y-3">
              <h3 className="heading-section text-petrol text-xl">
                Carrega mais do que si mesmo.
              </h3>
              <p className="text-petrol/75 leading-relaxed">
                Sócios, equipe, filhos, pais — gente depende de você inteiro.
                Saúde virou variável crítica do que você sustenta.
              </p>
            </div>

            <div className="space-y-3">
              <h3 className="heading-section text-petrol text-xl">
                Está numa transição que ninguém nomeia.
              </h3>
              <p className="text-petrol/75 leading-relaxed">
                Sono, humor, corpo — algo está mudando e a resposta clínica é
                "ainda não é nada" ou "é da idade". Falta um mapa.
              </p>
            </div>

            <div className="space-y-3">
              <h3 className="heading-section text-petrol text-xl">
                Informação demais, integração de menos.
              </h3>
              <p className="text-petrol/75 leading-relaxed">
                Personal, nutri, terapeuta, médico — cada um no seu pedaço.
                Wearables, painéis, relatórios — dado sobrando, decisão
                faltando.
              </p>
            </div>

            <div className="space-y-3">
              <h3 className="heading-section text-petrol text-xl">
                Quer estar inteiro para quem ainda vai chegar.
              </h3>
              <p className="text-petrol/75 leading-relaxed">
                Filho pequeno aos 45, projetos que não cabem em dez anos. A
                pergunta deixou de ser quanto vou viver — virou como vou estar
                quando precisarem.
              </p>
            </div>
          </div>

          <div className="mt-20 grid lg:grid-cols-[2fr_1fr] gap-10 items-center border-t border-petrol/10 pt-10">
            <div className="space-y-3">
              <p className="label-upper text-gold">Em dúvida?</p>
              <p className="text-petrol text-xl md:text-2xl heading-section max-w-xl">
                Não tem certeza se algum desses retratos é o seu?
              </p>
              <p className="text-petrol/70 leading-relaxed max-w-xl">
                Cinco perguntas curtas devolvem quais retratos do programa mais
                se aproximam do seu momento. Sem diagnóstico, sem captura de
                dados.
              </p>
            </div>
            <div className="flex lg:justify-end">
              <Link href="/diagnostico" className="btn-gold">
                Fazer o diagnóstico
              </Link>
            </div>
          </div>
        </div>
      </section>

      {/* EQUIPE — referência curta (não section completa) */}
      <section className="bg-petrol text-cream">
        <div className="site-container section grid lg:grid-cols-2 gap-16 items-start">
          <div className="space-y-6">
            <p className="label-upper text-gold">Equipe</p>
            <h2 className="heading-section text-cream text-3xl md:text-5xl">
              Quatro especialistas. Um plano único.
            </h2>
            <p className="text-cream/80 text-lg leading-relaxed max-w-xl">
              Médico, nutricionista, psicólogo e educador físico se reúnem antes
              do seu primeiro encontro, discutem o seu caso e desenham um plano
              integrado. Cada decisão clínica conversa com as outras.
            </p>
          </div>
          <div className="lg:pt-20">
            <Link href="/equipe" className="btn-outline-dark border-cream/40 text-cream">
              Conhecer a equipe completa
            </Link>
          </div>
        </div>
      </section>

      {/* COMO ACONTECE — fusão Cobertura AGIR + Jornada */}
      <section className="bg-paper">
        <div className="site-container section">
          <div className="max-w-3xl space-y-4 mb-16">
            <p className="label-upper text-gold">Como acontece</p>
            <h2 className="heading-section text-petrol text-3xl md:text-5xl">
              Uma frente por semana. Todas presentes todo mês.
            </h2>
            <p className="text-petrol/70 text-lg leading-relaxed">
              No ciclo de quatro semanas, cada pilar do Método AGIR é abordado por
              um profissional responsável — sob coordenação clínica do médico.
              Nenhum pilar fica sem revisão mensal.
            </p>
          </div>

          {/* Mini-grid 4 letras AGIR */}
          <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-px bg-petrol/15 border-y border-petrol/15 mb-20">
            {cobertura.map((c) => (
              <div key={c.code} className="bg-paper p-8 space-y-4">
                <span className="heading-section text-gold text-5xl block leading-none">{c.code}</span>
                <h3 className="heading-section text-petrol text-lg">{c.name}</h3>
                <p className="label-upper text-petrol/55 text-[10px] tracking-[0.2em] pt-2 border-t border-petrol/10">
                  {c.quem}
                </p>
              </div>
            ))}
          </div>

          {/* Cronologia */}
          <div className="max-w-3xl mb-12">
            <p className="label-upper text-gold mb-3">A cronologia</p>
            <h3 className="heading-section text-petrol text-2xl md:text-3xl">
              Do tempo zero ao próximo horizonte.
            </h3>
          </div>

          <div className="border-t border-petrol/15">
            {jornada.map((j, i) => (
              <article
                key={j.fase}
                className="grid md:grid-cols-[180px_1fr] gap-6 md:gap-12 py-10 border-b border-petrol/10"
              >
                <div className="space-y-2">
                  <p className="label-upper text-gold">{String(i + 1).padStart(2, '0')}</p>
                  <p className="label-upper text-petrol/60">{j.fase}</p>
                </div>
                <div className="space-y-3 max-w-2xl">
                  <h3 className="heading-section text-petrol text-2xl md:text-3xl leading-tight">
                    {j.titulo}
                  </h3>
                  <p className="text-petrol/75 leading-relaxed">{j.body}</p>
                </div>
              </article>
            ))}
          </div>
        </div>
      </section>

      {/* MODALIDADES */}
      <section className="bg-cream">
        <div className="site-container section">
          <div className="max-w-3xl space-y-4 mb-16">
            <p className="label-upper text-gold">Modalidades</p>
            <h2 className="heading-section text-petrol text-3xl md:text-5xl">
              Dois horizontes de compromisso.
            </h2>
          </div>

          <div className="grid md:grid-cols-2 gap-12 mb-12">
            {modalidades.map((m) => (
              <article
                key={m.name}
                className={
                  m.highlight
                    ? 'border-t-2 border-gold pt-8 space-y-5'
                    : 'border-t border-petrol/20 pt-8 space-y-5'
                }
              >
                <h3 className="heading-section text-petrol text-3xl md:text-4xl">{m.name}</h3>
                <p className="label-upper text-petrol/60">{m.period}</p>
                <p className="text-petrol/75 leading-relaxed">{m.body}</p>
              </article>
            ))}
          </div>

          <div className="border-t border-petrol/10 pt-8 max-w-3xl space-y-2">
            <p className="label-upper text-gold">Valores</p>
            <p className="text-petrol/75 leading-relaxed">
              Sob consulta. A equipe apresenta as condições em conversa direta,
              conforme a modalidade escolhida.
            </p>
          </div>
        </div>
      </section>

      {/* ENTRA · NÃO ENTRA */}
      <section className="bg-paper">
        <div className="site-container section">
          <div className="max-w-3xl space-y-4 mb-16">
            <p className="label-upper text-gold">Escopo</p>
            <h2 className="heading-section text-petrol text-3xl md:text-5xl">
              O que entra. O que fica de fora.
            </h2>
          </div>

          <div className="grid md:grid-cols-2 gap-12 lg:gap-16">
            <div className="space-y-6">
              <p className="label-upper text-petrol/55">Entra no programa</p>
              <ul className="space-y-4">
                {entra.map((item) => (
                  <li key={item} className="flex gap-4 text-petrol/85 leading-relaxed">
                    <span className="text-gold mt-1.5 leading-none">—</span>
                    <span>{item}</span>
                  </li>
                ))}
              </ul>
            </div>

            <div className="space-y-6 md:border-l md:border-petrol/15 md:pl-12">
              <p className="label-upper text-petrol/55">Fica de fora</p>
              <ul className="space-y-4">
                {naoEntra.map((item, i) => (
                  <li key={i} className="flex gap-4 text-petrol/75 leading-relaxed">
                    <span className="text-petrol/40 mt-1.5 leading-none">—</span>
                    <span>{item}</span>
                  </li>
                ))}
              </ul>
            </div>
          </div>
        </div>
      </section>

      {/* CTA + cross-links no padrão Plenya */}
      <section className="bg-petrol text-cream">
        <div className="site-container section">
          <div className="max-w-2xl space-y-6 mb-20">
            <p className="label-upper text-gold">Próximo passo</p>
            <h2 className="heading-section text-cream text-3xl md:text-4xl">
              Conversar com a equipe Plenya.
            </h2>
            <p className="text-cream/75 leading-relaxed">
              Entre em contato para entender se o Continuum faz sentido pro seu
              momento e desenhar o início do acompanhamento.
            </p>
            <p>
              <Link href="/contato" className="btn-gold">Entrar em contato</Link>
            </p>
            <p className="text-cream/60 text-sm pt-2 leading-relaxed">
              Ainda em dúvida se é pra você?{' '}
              <Link
                href="/diagnostico"
                className="underline decoration-gold/60 underline-offset-4 hover:text-cream transition"
              >
                Faça o diagnóstico em cinco perguntas
              </Link>
              .
            </p>
          </div>

          <div className="grid md:grid-cols-3 gap-8">
            <Link href="/escore-plenya" className="border-t border-cream/20 pt-8 space-y-3 group">
              <p className="label-upper text-gold">Instrumento</p>
              <p className="heading-section text-cream text-xl group-hover:text-gold transition">Escore Plenya →</p>
            </Link>
            <Link href="/metodo-agir" className="border-t border-cream/20 pt-8 space-y-3 group">
              <p className="label-upper text-gold">Método</p>
              <p className="heading-section text-cream text-xl group-hover:text-gold transition">Método AGIR →</p>
            </Link>
            <Link href="/consultas" className="border-t border-cream/20 pt-8 space-y-3 group">
              <p className="label-upper text-gold">Avulso</p>
              <p className="heading-section text-cream text-xl group-hover:text-gold transition">Consultas Plenya →</p>
            </Link>
          </div>
        </div>
      </section>
    </>
  );
}
