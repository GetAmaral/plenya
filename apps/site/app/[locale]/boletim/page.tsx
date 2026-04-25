import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';
import { NewsletterInline } from '@/components/blog/newsletter-inline';
import { Link } from '@/lib/i18n/navigation';

export const metadata: Metadata = {
  title: 'Boletim Plenya — medicina antecipatória na sua caixa de entrada',
  description:
    'Um e-mail por semana, escrito pela equipe médica Plenya. Leitura curta, embasada e prática. Sem ruído, sem promessa milagrosa, sem venda agressiva.',
};

const promessas = [
  {
    label: 'Frequência',
    title: 'Um e-mail por semana.',
    body: 'Toda quinta-feira pela manhã. Nem mais — porque saúde não se aprende em ruído.',
  },
  {
    label: 'Conteúdo',
    title: 'Escrito por médicos e equipe.',
    body: 'Leitura clínica de um exame, um nutriente, um hábito ou um achado novo da literatura. Sempre com referência.',
  },
  {
    label: 'Tempo',
    title: 'Cinco minutos de leitura.',
    body: 'Chega para você decidir o que fazer no fim da semana — ou trazer pra próxima consulta.',
  },
  {
    label: 'Privacidade',
    title: 'Seu e-mail não vira lista de venda.',
    body: 'Zero compartilhamento com terceiros. Cancelamento em 1 clique. 100% LGPD.',
  },
];

export default async function BoletimPage({ params }: { params: Promise<{ locale: string }> }) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <>
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32 max-w-3xl">
          <p className="label-upper text-gold mb-6">Boletim Plenya</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,4.5rem)] text-cream">
            Medicina antecipatória na sua caixa de entrada.
          </h1>
          <p className="text-cream/75 text-lg leading-relaxed mt-8 max-w-2xl">
            Um recorte semanal escrito pela equipe médica Plenya. Curto, embasado e
            prático — pra quem quer cuidar antes da janela silenciosa fechar.
          </p>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-container section grid md:grid-cols-2 gap-x-12 gap-y-10">
          {promessas.map((p) => (
            <div key={p.label} className="border-t border-petrol/15 pt-6 space-y-2">
              <p className="label-upper text-gold">{p.label}</p>
              <h2 className="heading-section text-petrol text-xl md:text-2xl">{p.title}</h2>
              <p className="text-petrol/75 leading-relaxed">{p.body}</p>
            </div>
          ))}
        </div>
      </section>

      <section className="bg-paper">
        <div className="site-container section max-w-2xl">
          <NewsletterInline />
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-container section text-center space-y-6 max-w-2xl mx-auto">
          <p className="label-upper text-gold">Quer ver antes de assinar?</p>
          <p className="heading-section text-petrol text-2xl md:text-3xl">
            Os artigos mais lidos do Boletim ficam abertos no nosso blog.
          </p>
          <p>
            <Link href="/blog" className="btn-gold">Ler o blog</Link>
          </p>
        </div>
      </section>
    </>
  );
}
