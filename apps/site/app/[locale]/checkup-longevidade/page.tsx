import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { FaqAccordion } from '@/components/marketing/faq-accordion';
import { FaqSchema } from '@/components/seo/faq-schema';
import { BreadcrumbSchema } from '@/components/seo/breadcrumb-schema';
import { MedicalWebPageSchema } from '@/components/seo/medical-webpage-schema';
import { ClinicalReviewBadge } from '@/components/marketing/clinical-review-badge';
import { RelatedBlogPosts } from '@/components/marketing/related-blog-posts';

const checkupFaq = [
  {
    q: 'O que é um checkup de longevidade?',
    a: 'É uma avaliação clínica orientada não apenas a detectar doença instalada, mas a mapear a trajetória de saúde — capacidade cardiorrespiratória, composição corporal, marcadores metabólicos avançados, sono, cognição, hábitos. O objetivo é identificar a janela silenciosa entre o normal e o ótimo, antes que o sintoma apareça.',
  },
  {
    q: 'Em que se diferencia do checkup tradicional?',
    a: 'O checkup tradicional valida normalidade segundo faixas amplas. O checkup de longevidade integra marcadores que predizem trajetória — VO₂ máx, ApoB, Lp(a), composição corporal, ômega-3 índice — e cruza com hábitos, sono e exposições. Não é um carimbo de "normal"; é um plano.',
  },
  {
    q: 'Quais exames fazem parte?',
    a: 'O painel base inclui perfil lipídico avançado (ApoB, Lp(a)), HbA1c, glicemia em jejum e pós-prandial, PCR ultrassensível, vitamina D, ferritina, B12, homocisteína, ômega-3 índice, magnésio, hormônios tireoidianos, vitamina K2 (quando indicado), 25-OH-vitamina D. Avaliação funcional inclui ergoespirometria (VO₂ máx), bioimpedância, dinamometria de preensão. Imagem é dirigida por risco — não usamos varredura sem indicação.',
  },
  {
    q: 'Quanto tempo demora?',
    a: 'O checkup de longevidade Plenya é um processo, não um dia. A coleta de exames acontece em uma manhã. A consulta de retorno, com leitura clínica e plano individualizado, dura 90 minutos. Pacientes que aderem ao programa Continuum têm reavaliação semestral.',
  },
  {
    q: 'É indicado a partir de qual idade?',
    a: 'A janela mais valiosa é entre 35 e 55 anos — antes do dano se instalar, depois das primeiras alterações silenciosas. Acima de 55, ainda é útil para reverter o que é reversível e estabilizar o resto. Abaixo de 35, em pessoas com história familiar de evento cardiovascular precoce, faz sentido.',
  },
  {
    q: 'Plano de saúde cobre?',
    a: 'A consulta Plenya é particular. Os exames podem, em parte, ser feitos pelo seu plano — orientamos a melhor estratégia caso a caso. A maioria dos pacientes opta por concentrar tudo em um laboratório parceiro para garantir padronização das técnicas e comparabilidade longitudinal.',
  },
];

export const metadata: Metadata = {
  title: 'Checkup de longevidade — avaliação clínica baseada em evidência',
  description:
    'Checkup orientado à trajetória de saúde: VO₂ máx, ApoB, composição corporal, sono, cognição. Mapeia a janela silenciosa antes do sintoma. Plenya, Londrina-PR.',
  alternates: { canonical: '/checkup-longevidade' },
  openGraph: {
    title: 'Checkup de longevidade — Plenya',
    description:
      'Avaliação clínica orientada à trajetória de saúde — não só ao diagnóstico de doença instalada.',
    type: 'article',
  },
};

export default async function CheckupPage({
  params,
}: {
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <>
      <FaqSchema items={checkupFaq} />
      <MedicalWebPageSchema
        name="Checkup de longevidade — Plenya"
        description="Avaliação clínica orientada à trajetória de saúde — VO₂ máx, ApoB, composição corporal, sono, cognição. Mapeia a janela silenciosa antes do sintoma."
        path="/checkup-longevidade"
        about="Preventive medical checkup"
      />
      <BreadcrumbSchema
        items={[
          { name: 'Home', url: '/' },
          { name: 'Checkup de Longevidade' },
        ]}
      />

      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32 max-w-3xl">
          <p className="label-upper text-gold mb-6">Checkup de Longevidade</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,4.5rem)] text-cream">
            O exame que olha pra frente.
          </h1>
          <p className="text-cream/75 text-lg leading-relaxed mt-8 max-w-2xl">
            Um checkup tradicional confirma se você está doente hoje. O checkup de longevidade
            mapeia para onde sua saúde está indo nos próximos vinte anos — e o que ainda dá tempo
            de mudar.
          </p>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-container section grid lg:grid-cols-[2fr_1fr] gap-16">
          <div className="space-y-6 text-petrol/85 text-lg leading-relaxed">
            <p className="label-upper text-gold">Sobre o checkup</p>
            <p>
              A medicina convencional valida normalidade. A medicina de longevidade lê trajetória.
              São perguntas diferentes — e exigem instrumentos diferentes.
            </p>
            <p>
              No Brasil, a maior parte dos checkups anuais ainda é desenhada para identificar
              doença já instalada. Hemograma, glicemia, colesterol total, ultrassom abdominal — é
              útil, mas insuficiente. ApoB prediz risco cardiovascular melhor que LDL. VO₂ máximo
              prediz mortalidade total melhor que qualquer outro marcador isolado. Composição
              corporal por bioimpedância distingue quem perdeu peso de gordura de quem perdeu de
              músculo. Esses são os exames que mudam decisão.
            </p>
            <p>
              O checkup de longevidade Plenya é estruturado em três camadas. A primeira, exames
              laboratoriais avançados — perfil lipídico ApoB-cêntrico, marcadores inflamatórios,
              minerais e vitaminas que regulam função celular. A segunda, avaliação funcional —
              ergoespirometria para VO₂ máx, dinamometria, bioimpedância segmentar. A terceira,
              avaliação de hábito e ambiente — sono medido, exposição à luz, padrão alimentar,
              atividade física registrada por wearable.
            </p>
            <p>
              O resultado não é um envelope com marcações em vermelho. É uma conversa de 90
              minutos com leitura clínica integrada e um plano nominal — substância, dose,
              frequência, prazo de reavaliação.
            </p>
          </div>

          <aside className="space-y-6">
            <div className="border border-petrol/10 p-6 space-y-3">
              <p className="label-upper text-gold">Camada 1 — Exames</p>
              <ul className="text-petrol/75 text-sm space-y-1.5 leading-relaxed">
                <li>· ApoB, Lp(a), perfil lipídico avançado</li>
                <li>· HbA1c, glicemia jejum e pós-prandial</li>
                <li>· PCR ultrassensível, ferritina</li>
                <li>· 25-OH-vit D, B12, ômega-3 índice, magnésio</li>
                <li>· Hormônios tireoidianos completos</li>
              </ul>
            </div>
            <div className="border border-petrol/10 p-6 space-y-3">
              <p className="label-upper text-gold">Camada 2 — Função</p>
              <ul className="text-petrol/75 text-sm space-y-1.5 leading-relaxed">
                <li>· Ergoespirometria (VO₂ máx)</li>
                <li>· Bioimpedância segmentar</li>
                <li>· Dinamometria de preensão</li>
              </ul>
            </div>
            <div className="border border-petrol/10 p-6 space-y-3">
              <p className="label-upper text-gold">Camada 3 — Hábito</p>
              <ul className="text-petrol/75 text-sm space-y-1.5 leading-relaxed">
                <li>· Sono medido (questionário + wearable)</li>
                <li>· Diário alimentar de 7 dias</li>
                <li>· Cronotipo + exposição à luz</li>
              </ul>
            </div>
          </aside>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-container">
          <ClinicalReviewBadge />
        </div>
      </section>

      <FaqAccordion title="Perguntas que recebemos toda semana." items={checkupFaq} />

      <RelatedBlogPosts
        title="Do blog Plenya"
        pillars={['gestao-metabolica', 'longevidade']}
        limit={3}
      />

      <section className="bg-cream">
        <div className="site-container section grid md:grid-cols-12 gap-12">
          <div className="md:col-span-6 space-y-6">
            <p className="label-upper text-gold">Agendamento</p>
            <h2 className="heading-section text-petrol text-3xl md:text-4xl">
              Marcar o checkup.
            </h2>
            <p className="text-petrol/75 leading-relaxed max-w-md">
              Entre em contato com a equipe Plenya para entender o melhor formato para o seu
              momento — consulta isolada ou programa Continuum.
            </p>
            <p>
              <Link href="/contato" className="btn-gold">
                Falar com a equipe
              </Link>
            </p>
          </div>
          <div className="md:col-span-6 grid sm:grid-cols-2 gap-6 content-end">
            <Link
              href="/healthspan"
              className="border border-petrol/10 p-6 hover:bg-petrol hover:text-cream transition group"
            >
              <p className="label-upper text-gold mb-2">Healthspan</p>
              <p className="text-petrol/85 group-hover:text-cream/85 text-sm">
                A agenda por trás do checkup.
              </p>
            </Link>
            <Link
              href="/continuum"
              className="border border-petrol/10 p-6 hover:bg-petrol hover:text-cream transition group"
            >
              <p className="label-upper text-gold mb-2">Continuum</p>
              <p className="text-petrol/85 group-hover:text-cream/85 text-sm">
                Acompanhamento longitudinal.
              </p>
            </Link>
          </div>
        </div>
      </section>
    </>
  );
}
