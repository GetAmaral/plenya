import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { FaqAccordion } from '@/components/marketing/faq-accordion';
import { FaqSchema } from '@/components/seo/faq-schema';

const renalFaq = [
  {
    q: 'Por que pensar em rim antes de qualquer sintoma?',
    a: 'O rim é o órgão silencioso. Perde função em curva — sem dor, sem sintoma — e só dá sinal quando o dano já é irreversível. Cerca de 10% da população adulta brasileira tem doença renal crônica, e a maioria não sabe. Detectar precocemente, na janela em que o dano ainda é reversível ou estabilizável, muda a trajetória da próxima década.',
  },
  {
    q: 'Quais exames detectam alteração renal precoce?',
    a: 'Creatinina sérica e estimativa de filtração glomerular (eGFR) são o básico — mas insuficientes isoladamente. A combinação que muda decisão clínica: cistatina C (mais sensível em fases iniciais), albuminúria (relação albumina/creatinina urinária), urina tipo I com sedimentoscopia, eletrólitos, pressão arterial em 24h (MAPA) quando indicado. ApoB e perfil metabólico entram porque rim e cardiovascular estão acoplados.',
  },
  {
    q: 'Quem deveria fazer avaliação renal preventiva?',
    a: 'Toda pessoa acima dos 40 anos. Toda pessoa com hipertensão, diabetes, obesidade ou síndrome metabólica em qualquer idade. Toda pessoa com história familiar de doença renal, evento cardiovascular precoce ou rim único. Atletas que usam suplementação proteica intensa ou anti-inflamatórios crônicos. Pacientes em uso prolongado de IBP, lítio, anti-inflamatórios não esteroidais.',
  },
  {
    q: 'Suplementos podem prejudicar o rim?',
    a: 'Em pessoas com função renal normal, a evidência é tranquilizadora para creatina, ômega-3, vitamina D, magnésio. A preocupação concentra-se em doses elevadas de proteína em quem já tem função reduzida, anti-inflamatórios usados cronicamente sem indicação, suplementos importados sem registro, e fórmulas de "detox renal" — que não têm evidência e podem mascarar problema real.',
  },
  {
    q: 'O que muda numa avaliação renal feita por nefrologista vs clínico geral?',
    a: 'O nefrologista lê os mesmos exames com lupa: padrão de declínio da eGFR, interpretação fina da albuminúria por estágio, ajuste medicamentoso por função residual, decisão sobre nefroprotetores em fase pré-clínica. Para o paciente sem queixa, é a diferença entre um carimbo e um plano.',
  },
  {
    q: 'A Plenya faz acompanhamento renal?',
    a: 'Sim. O Dr. Getúlio Amaral Filho é nefrologista com RQE 16.038, professor da Santa Casa de Londrina e coordenador da residência médica em nefrologia. A avaliação renal preventiva é parte estruturada do programa Continuum Plenya, integrada à leitura cardiometabólica e funcional.',
  },
];

export const metadata: Metadata = {
  title: 'Avaliação renal preventiva — nefrologia antes do sintoma',
  description:
    'O rim é o órgão silencioso. Avaliação preventiva com cistatina C, albuminúria, eGFR e leitura nefrológica integrada. Dr. Getúlio Amaral Filho · Plenya · Londrina-PR.',
  alternates: { canonical: '/avaliacao-renal-preventiva' },
  openGraph: {
    title: 'Avaliação renal preventiva — Plenya',
    description:
      'Detectar a janela silenciosa em que o dano renal ainda é reversível.',
    type: 'article',
  },
};

export default async function RenalPage({
  params,
}: {
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <>
      <FaqSchema items={renalFaq} />

      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32 max-w-3xl">
          <p className="label-upper text-gold mb-6">Avaliação Renal Preventiva</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,4.5rem)] text-cream">
            O rim é o órgão silencioso.
          </h1>
          <p className="text-cream/75 text-lg leading-relaxed mt-8 max-w-2xl">
            Perde função em curva, sem sintoma. Quando dói, o dano já passou da janela. Avaliação
            renal preventiva é olhar o rim na fase em que ainda há tempo de mudar a trajetória.
          </p>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-container section grid lg:grid-cols-[2fr_1fr] gap-16">
          <div className="space-y-6 text-petrol/85 text-lg leading-relaxed">
            <p className="label-upper text-gold">Por que olhar antes</p>
            <p>
              Aproximadamente 10% da população adulta brasileira tem algum grau de doença renal
              crônica. Mais de 90% não sabe. O rim só dá sinal quando perdeu mais da metade da
              função — e a esse ponto, muito do que se perde não volta.
            </p>
            <p>
              Por isso a abordagem nefrológica preventiva muda agenda. Não esperamos a creatinina
              subir; lemos cistatina C, albuminúria, sedimento urinário e padrão pressórico em 24
              horas. Cruzamos com perfil cardiometabólico — porque rim e coração compartilham via
              de dano. Ajustamos medicação que sobrecarrega rim antes de gerar problema. Tratamos
              hipertensão por nefroproteção, não só por número.
            </p>
            <p>
              A avaliação renal preventiva Plenya é conduzida pelo <strong>Dr. Getúlio Amaral
              Filho</strong>, nefrologista (CRM-PR 21.876 · RQE 16.038), professor e coordenador da
              residência médica em nefrologia da Santa Casa de Londrina. A leitura é específica:
              cada exame entra em contexto, cada decisão tem prazo de reavaliação.
            </p>
          </div>

          <aside className="space-y-6">
            <div className="border border-petrol/10 p-6 space-y-3">
              <p className="label-upper text-gold">Painel renal preventivo</p>
              <ul className="text-petrol/75 text-sm space-y-1.5 leading-relaxed">
                <li>· Creatinina sérica + eGFR</li>
                <li>· Cistatina C (sensibilidade precoce)</li>
                <li>· Albuminúria (relação A/C urinária)</li>
                <li>· Urina tipo I + sedimentoscopia</li>
                <li>· Eletrólitos completos</li>
                <li>· MAPA (quando indicado)</li>
                <li>· Ácido úrico, vitamina D, PTH</li>
              </ul>
            </div>
            <div className="border border-petrol/10 p-6 space-y-3">
              <p className="label-upper text-gold">Quem deve avaliar</p>
              <ul className="text-petrol/75 text-sm space-y-1.5 leading-relaxed">
                <li>· Adultos &gt; 40 anos (rotina)</li>
                <li>· Hipertensos, diabéticos, obesos</li>
                <li>· História familiar de DRC</li>
                <li>· Uso crônico de AINE, IBP, lítio</li>
                <li>· Atletas com suplementação intensa</li>
              </ul>
            </div>
          </aside>
        </div>
      </section>

      <FaqAccordion
        title="Perguntas frequentes sobre avaliação renal."
        items={renalFaq}
      />

      <section className="bg-cream">
        <div className="site-container section grid md:grid-cols-12 gap-12">
          <div className="md:col-span-6 space-y-6">
            <p className="label-upper text-gold">Agendamento</p>
            <h2 className="heading-section text-petrol text-3xl md:text-4xl">
              Avaliar com nefrologista.
            </h2>
            <p className="text-petrol/75 leading-relaxed max-w-md">
              Consulta com o Dr. Getúlio Amaral Filho — presencial em Londrina ou online.
            </p>
            <p>
              <Link href="/contato" className="btn-gold">
                Marcar consulta
              </Link>
            </p>
          </div>
          <div className="md:col-span-6 grid sm:grid-cols-2 gap-6 content-end">
            <Link
              href="/dr-getulio"
              className="border border-petrol/10 p-6 hover:bg-petrol hover:text-cream transition group"
            >
              <p className="label-upper text-gold mb-2">Dr. Getúlio</p>
              <p className="text-petrol/85 group-hover:text-cream/85 text-sm">
                Trajetória e formação.
              </p>
            </Link>
            <Link
              href="/checkup-longevidade"
              className="border border-petrol/10 p-6 hover:bg-petrol hover:text-cream transition group"
            >
              <p className="label-upper text-gold mb-2">Checkup completo</p>
              <p className="text-petrol/85 group-hover:text-cream/85 text-sm">
                Avaliação integrada.
              </p>
            </Link>
          </div>
        </div>
      </section>
    </>
  );
}
