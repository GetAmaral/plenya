import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { FaqAccordion } from '@/components/marketing/faq-accordion';
import { FaqSchema } from '@/components/seo/faq-schema';

const mfiFaq = [
  {
    q: 'O que é medicina funcional integrativa?',
    a: 'É uma abordagem clínica que parte da pergunta "por que esse paciente, por essa via, agora?" — e não apenas "qual o nome do quadro". Olha sistemas (digestivo, hormonal, mitocondrial, imunológico) como rede acoplada, integra dados de hábito, exposição e laboratório, e propõe intervenção em cascata: hábito antes de suplemento, suplemento antes de medicação, medicação antes de procedimento.',
  },
  {
    q: 'Em que se diferencia da medicina convencional?',
    a: 'Não substitui — integra. A medicina convencional é insubstituível em diagnóstico de doença instalada e em tratamento agudo. A funcional integrativa atua antes (prevenção primária), no entremeio (causas funcionais que não fecham diagnóstico tradicional) e depois (acompanhamento metabólico de quem já tem doença). A diferença prática: tempo de consulta, profundidade da história, exames pedidos e linguagem de plano.',
  },
  {
    q: 'É medicina alternativa?',
    a: 'Não. É medicina convencional praticada com mais tempo, mais escuta, mais leitura sistêmica e mais critério na recomendação de hábito, suplementação e medicação. Tudo o que recomendamos passa por evidência publicada. Não há terapias esotéricas, sem registro ou fora do escopo do CFM.',
  },
  {
    q: 'A formação do médico em medicina funcional é reconhecida?',
    a: 'A formação do Dr. Getúlio em medicina funcional integrativa é pela ABMFI (Associação Brasileira de Medicina Funcional Integrativa) — instituição que segue padrões internacionais e exige formação médica prévia + especialidade reconhecida pelo CFM. Funcional não substitui a especialidade; soma a ela. O Dr. Getúlio mantém especialidade primária em nefrologia (RQE 16.038) e clínica médica.',
  },
  {
    q: 'Quais condições mais respondem à abordagem funcional?',
    a: 'Síndrome metabólica e pré-diabetes em fase reversível, fadiga crônica sem causa óbvia, distúrbios de sono e ritmo circadiano, alterações hormonais em homens (testosterona) e mulheres (peri e pós-menopausa), sintomas digestivos funcionais, deficiências nutricionais subclínicas, dor crônica relacionada a inflamação sistêmica de baixo grau. Sempre como complemento ao manejo da doença base, não como substituto.',
  },
  {
    q: 'Quanto tempo até ver resultado?',
    a: 'Depende do alvo. Sono, energia e sintomas digestivos costumam responder em 4 a 8 semanas. Marcadores metabólicos (HbA1c, perfil lipídico) em 12 a 16 semanas. Composição corporal e VO₂ máx em 16 a 24 semanas. Tudo é mensurado em reavaliação programada — sem reavaliação não há medicina funcional, há marketing.',
  },
];

export const metadata: Metadata = {
  title: 'Medicina funcional integrativa — abordagem clínica em Londrina',
  description:
    'Medicina convencional praticada com mais tempo, escuta e leitura sistêmica. Cascata de intervenção baseada em evidência. Plenya, Londrina-PR.',
  alternates: { canonical: '/medicina-funcional-integrativa' },
  openGraph: {
    title: 'Medicina Funcional Integrativa — Plenya',
    description: 'Não é alternativa. É medicina convencional com tempo, escuta e critério.',
    type: 'article',
  },
};

export default async function MfiPage({
  params,
}: {
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <>
      <FaqSchema items={mfiFaq} />

      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32 max-w-3xl">
          <p className="label-upper text-gold mb-6">Medicina Funcional Integrativa</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,4.5rem)] text-cream">
            Por que esse paciente, por essa via, agora.
          </h1>
          <p className="text-cream/75 text-lg leading-relaxed mt-8 max-w-2xl">
            A pergunta da medicina convencional é <em>qual o diagnóstico</em>. A pergunta da
            medicina funcional integrativa é <em>por que aqui, por que agora, por que assim</em>.
            Não substitui — soma.
          </p>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-container section grid lg:grid-cols-[2fr_1fr] gap-16">
          <div className="space-y-6 text-petrol/85 text-lg leading-relaxed">
            <p className="label-upper text-gold">A abordagem</p>
            <p>
              A medicina funcional integrativa não é alternativa nem oposta à convencional. É
              medicina convencional praticada com tempo de consulta diferente — 60 a 90 minutos —,
              escuta sistêmica, exames que vão além do mínimo, e linguagem de plano que cabe na
              vida do paciente.
            </p>
            <p>
              A lógica é uma cascata. Primeiro hábito — sono, exposição à luz, alimentação,
              movimento, conexão. Depois suplementação dirigida por exame — não chute, não pacote
              comercial. Depois medicação, quando o sinal pede e o hábito não basta. Procedimento e
              cirurgia, quando há indicação clara. A ordem importa porque inverter custa mais e
              entrega menos.
            </p>
            <p>
              O Dr. Getúlio Amaral Filho mantém especialidade primária em nefrologia e clínica
              médica (RQE 16.038), e formação em medicina funcional integrativa pela ABMFI.
              Funcional não substitui a especialidade — soma. Por isso o Continuum Plenya tem
              equipe multidisciplinar: médico, nutricionista, psicólogo e educador físico, cada um
              com leitura própria do mesmo paciente.
            </p>
            <p>
              Tudo o que recomendamos passa por evidência publicada. Não há terapias esotéricas,
              sem registro CFM ou fora de escopo. O que diferencia é tempo, profundidade e
              critério — não ferramenta exótica.
            </p>
          </div>

          <aside className="space-y-6">
            <div className="border border-petrol/10 p-6 space-y-3">
              <p className="label-upper text-gold">A cascata</p>
              <ol className="text-petrol/75 text-sm space-y-1.5 leading-relaxed list-decimal pl-5">
                <li>Hábito (sono, luz, alimentação, movimento)</li>
                <li>Suplementação dirigida por exame</li>
                <li>Medicação, quando o hábito não basta</li>
                <li>Procedimento, quando há indicação clara</li>
              </ol>
            </div>
            <div className="border border-petrol/10 p-6 space-y-3">
              <p className="label-upper text-gold">Quando pode ajudar</p>
              <ul className="text-petrol/75 text-sm space-y-1.5 leading-relaxed">
                <li>· Pré-diabetes em fase reversível</li>
                <li>· Fadiga crônica sem causa óbvia</li>
                <li>· Sono fragmentado e baixa energia</li>
                <li>· Peri e pós-menopausa</li>
                <li>· Andropausa funcional</li>
                <li>· Dor crônica de baixo grau</li>
                <li>· Sintomas digestivos funcionais</li>
              </ul>
            </div>
          </aside>
        </div>
      </section>

      <FaqAccordion
        title="Perguntas frequentes sobre medicina funcional integrativa."
        items={mfiFaq}
      />

      <section className="bg-cream">
        <div className="site-container section grid md:grid-cols-12 gap-12">
          <div className="md:col-span-6 space-y-6">
            <p className="label-upper text-gold">Agendamento</p>
            <h2 className="heading-section text-petrol text-3xl md:text-4xl">
              Marcar consulta.
            </h2>
            <p className="text-petrol/75 leading-relaxed max-w-md">
              Consulta médica de 60-90 minutos com o Dr. Getúlio — presencial em Londrina ou online.
            </p>
            <p>
              <Link href="/contato" className="btn-gold">
                Falar com a equipe
              </Link>
            </p>
          </div>
          <div className="md:col-span-6 grid sm:grid-cols-2 gap-6 content-end">
            <Link
              href="/metodo-agir"
              className="border border-petrol/10 p-6 hover:bg-petrol hover:text-cream transition group"
            >
              <p className="label-upper text-gold mb-2">Método AGIR</p>
              <p className="text-petrol/85 group-hover:text-cream/85 text-sm">
                Os quatro pilares.
              </p>
            </Link>
            <Link
              href="/continuum"
              className="border border-petrol/10 p-6 hover:bg-petrol hover:text-cream transition group"
            >
              <p className="label-upper text-gold mb-2">Continuum</p>
              <p className="text-petrol/85 group-hover:text-cream/85 text-sm">
                Programa estruturado.
              </p>
            </Link>
          </div>
        </div>
      </section>
    </>
  );
}
