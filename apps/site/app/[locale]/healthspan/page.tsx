import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { FaqAccordion } from '@/components/marketing/faq-accordion';
import { FaqSchema } from '@/components/seo/faq-schema';
import { BreadcrumbSchema } from '@/components/seo/breadcrumb-schema';
import { MedicalWebPageSchema } from '@/components/seo/medical-webpage-schema';
import { ClinicalReviewBadge } from '@/components/marketing/clinical-review-badge';
import { RelatedBlogPosts } from '@/components/marketing/related-blog-posts';

const healthspanFaq = [
  {
    q: 'Qual a diferença entre lifespan e healthspan?',
    a: 'Lifespan é quanto tempo você vive. Healthspan é quanto desse tempo você vive com autonomia, lucidez e disposição. A medicina convencional otimizou o lifespan no século XX; o desafio do XXI é fechar a distância entre os dois — hoje, em média, os últimos 8 a 12 anos da vida brasileira são vividos com perda funcional significativa.',
  },
  {
    q: 'Healthspan é o mesmo que medicina antienvelhecimento?',
    a: 'Não. Medicina antienvelhecimento promete reverter o relógio, frequentemente sem evidência sólida. Healthspan é uma agenda mensurável: capacidade cardiorrespiratória (VO₂ máx), força e massa muscular, marcadores metabólicos, cognição, sono. O foco é estender a fase de plenitude funcional, não buscar imortalidade.',
  },
  {
    q: 'Quais marcadores realmente importam para healthspan?',
    a: 'Os mais robustos em literatura: VO₂ máximo (preditor mais forte de mortalidade total já estudado), força de preensão, massa muscular apendicular, ApoB, hemoglobina glicada, PCR ultrassensível, vitamina D, ômega-3 índice, qualidade de sono medida (não só duração). O Escore Plenya integra mais de 800 itens nessa direção.',
  },
  {
    q: 'A partir de que idade faz sentido pensar em healthspan?',
    a: 'A janela mais valiosa é entre 35 e 55 anos — antes do dano se instalar, depois que a metabolização hepática, a sensibilidade insulínica e a massa muscular começam a declinar silenciosamente. Mas a abordagem é útil em qualquer idade: aos 30 para construção, aos 60 para reversão do que ainda é reversível.',
  },
  {
    q: 'O programa Continuum Plenya é uma abordagem de healthspan?',
    a: 'Sim. O Continuum é o programa estruturado para healthspan na Plenya — semestral ou anual, com equipe multidisciplinar (médico, nutricionista, psicólogo e educador físico), exames específicos, plano de exercício individualizado e acompanhamento longitudinal mensurado pelo Escore Plenya.',
  },
  {
    q: 'O que healthspan não promete?',
    a: 'Não promete cura, não promete reverter doenças instaladas, não promete imortalidade. Promete dar a você instrumentos baseados em evidência para chegar aos 70, 80 ou 90 anos com mais autonomia do que o esperado para a média da sua geração — o que já é, em si, uma transformação.',
  },
];

export const metadata: Metadata = {
  title: 'Healthspan — viver bem por mais tempo, com evidência',
  description:
    'Healthspan é o tempo de vida com autonomia funcional. Programa Plenya estende a fase de plenitude com VO₂ máx, força, marcadores metabólicos e acompanhamento longitudinal.',
  alternates: { canonical: '/healthspan' },
  openGraph: {
    title: 'Healthspan — viver bem por mais tempo',
    description:
      'A diferença entre quantos anos você vive e quantos vive bem. A agenda da medicina contemporânea.',
    type: 'article',
  },
};

export default async function HealthspanPage({
  params,
}: {
  params: Promise<{ locale: string }>;
}) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <>
      <FaqSchema items={healthspanFaq} />
      <MedicalWebPageSchema
        name="Healthspan — viver bem por mais tempo"
        description="A diferença entre lifespan e healthspan. Programa Plenya estende a fase de plenitude com VO₂ máx, força, marcadores metabólicos e acompanhamento longitudinal."
        path="/healthspan"
        about="Healthspan"
      />
      <BreadcrumbSchema
        items={[
          { name: 'Home', url: '/' },
          { name: 'Healthspan' },
        ]}
      />

      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32 max-w-3xl">
          <p className="label-upper text-gold mb-6">Healthspan</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,4.5rem)] text-cream">
            Viver bem por mais tempo.
          </h1>
          <p className="text-cream/75 text-lg leading-relaxed mt-8 max-w-2xl">
            Lifespan é quantos anos você vive. <strong className="text-cream">Healthspan</strong> é
            quantos desses anos você vive com autonomia, lucidez e disposição. A medicina do século
            XX estendeu o primeiro. A medicina contemporânea responde pelo segundo.
          </p>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-container section grid lg:grid-cols-[2fr_1fr] gap-16">
          <div className="space-y-6 text-petrol/85 text-lg leading-relaxed">
            <p className="label-upper text-gold">A janela do meio</p>
            <p>
              No Brasil, a expectativa de vida saudável fica em torno de 65 anos — a expectativa
              total, em torno de 76. A diferença é de mais de uma década vivida com perda funcional:
              dor, dependência, fragilidade cognitiva, sarcopenia. Encurtar essa distância é a
              agenda do healthspan.
            </p>
            <p>
              Não se faz isso com promessa antienvelhecimento, suplementos exóticos ou exames
              proprietários sem validação. Faz-se com poucos marcadores que a literatura mostrou
              serem preditivos: capacidade cardiorrespiratória (VO₂ máx), força de preensão, massa
              muscular apendicular, ApoB, glicemia em jejum e pós-prandial, PCR, sono medido.
            </p>
            <p>
              A janela mais valiosa para intervir é entre 35 e 55 anos. É quando a sensibilidade
              insulínica começa a cair, a massa muscular declina ~1% ao ano sem treino, o sono
              fragmenta e os primeiros marcadores cardiovasculares mudam — mas tudo ainda é
              reversível com baixo custo. Esperar os sintomas é trocar uma intervenção barata por
              uma cara.
            </p>
            <p>
              Healthspan, na Plenya, é uma medida — não um slogan. Cada decisão clínica passa pelo
              Escore Plenya: mais de 800 itens organizados em quatro pilares (Atividade,
              Alimentação, Gestão metabólica, Integração mente-corpo, Ritmo circadiano). O
              acompanhamento é longitudinal — comparamos você com você mesmo, semestre a semestre.
            </p>
          </div>

          <aside className="space-y-6 lg:sticky lg:top-28 self-start">
            <div className="bg-paper border-l-2 border-gold p-7 space-y-4">
              <p className="label-upper text-gold">O que medimos</p>
              <ul className="text-petrol/85 text-base space-y-2 leading-relaxed">
                <li>· VO₂ máximo (ergoespirometria)</li>
                <li>· Força de preensão e massa muscular</li>
                <li>· ApoB, Lp(a), HbA1c, glicemia pós-prandial</li>
                <li>· PCR ultrassensível, ômega-3 índice</li>
                <li>· 25-OH-vitamina D, magnésio</li>
                <li>· Sono (questionários + wearable)</li>
                <li>· Cognição (rastreio funcional)</li>
              </ul>
            </div>
            <div className="bg-petrol text-cream p-7 space-y-4">
              <p className="label-upper text-gold">Programa Continuum</p>
              <p className="heading-section text-cream text-2xl leading-tight">
                Healthspan aplicado, mensurado a cada ciclo.
              </p>
              <p className="text-cream/75 text-base leading-relaxed">
                Semestral ou anual, equipe multidisciplinar, plano único —
                Escore Plenya como medida.
              </p>
              <Link href="/continuum" className="btn-outline-light w-full text-center">
                Conhecer o Continuum →
              </Link>
            </div>
          </aside>
        </div>
      </section>

      <section className="bg-paper">
        <div className="site-container section max-w-4xl space-y-8">
          <p className="label-upper text-gold">A leitura do Continuum</p>
          <h2 className="heading-section text-petrol text-3xl md:text-4xl max-w-2xl">
            O que healthspan <em>não</em> promete.
          </h2>
          <div className="grid md:grid-cols-2 gap-8 text-petrol/80 leading-relaxed">
            <div className="space-y-3">
              <p className="label-upper text-petrol/60">Não promete</p>
              <ul className="space-y-2 text-base">
                <li>· Reverter o envelhecimento</li>
                <li>· Cura para doenças instaladas</li>
                <li>· Garantia de longevidade</li>
                <li>· Substituir tratamento médico convencional</li>
              </ul>
            </div>
            <div className="space-y-3">
              <p className="label-upper text-petrol/60">Promete</p>
              <ul className="space-y-2 text-base">
                <li>· Marcadores baseados em evidência</li>
                <li>· Acompanhamento longitudinal mensurado</li>
                <li>· Plano individualizado e revisado a cada semestre</li>
                <li>· Equipe multidisciplinar com responsabilidade clínica</li>
              </ul>
            </div>
          </div>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-container">
          <ClinicalReviewBadge />
        </div>
      </section>

      <FaqAccordion title="Perguntas frequentes sobre healthspan." items={healthspanFaq} />

      <RelatedBlogPosts
        title="Do blog Plenya"
        pillars={['longevidade', 'gestao-metabolica']}
        limit={3}
      />

      <section className="bg-cream">
        <div className="site-container section grid md:grid-cols-12 gap-12">
          <div className="md:col-span-6 space-y-6">
            <p className="label-upper text-gold">Próximo passo</p>
            <h2 className="heading-section text-petrol text-3xl md:text-4xl">
              Saber onde você está.
            </h2>
            <p className="text-petrol/75 leading-relaxed max-w-md">
              O Escore Plenya é o instrumento de medida do Método AGIR. Comece com uma avaliação
              gratuita e veja seu mapa de healthspan.
            </p>
            <div className="flex gap-4 flex-wrap pt-2">
              <Link href="/escore-plenya" className="btn-gold">
                Avaliar gratuitamente
              </Link>
              <Link href="/contato" className="btn-outline-petrol">
                Falar com a equipe
              </Link>
            </div>
          </div>
          <div className="md:col-span-6 grid sm:grid-cols-2 gap-6 content-end">
            <Link
              href="/continuum"
              className="border border-petrol/10 p-6 hover:bg-petrol hover:text-cream transition group"
            >
              <p className="label-upper text-gold mb-2">Continuum</p>
              <p className="text-petrol/85 group-hover:text-cream/85 text-sm">
                Programa semestral ou anual.
              </p>
            </Link>
            <Link
              href="/metodo-agir"
              className="border border-petrol/10 p-6 hover:bg-petrol hover:text-cream transition group"
            >
              <p className="label-upper text-gold mb-2">Método AGIR</p>
              <p className="text-petrol/85 group-hover:text-cream/85 text-sm">
                Os quatro pilares.
              </p>
            </Link>
          </div>
        </div>
      </section>
    </>
  );
}
