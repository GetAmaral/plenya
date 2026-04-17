import type { Metadata } from 'next';
import Image from 'next/image';
import { setRequestLocale } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';

export const metadata: Metadata = {
  title: 'Escore Plenya',
  description: 'Diagnóstico estruturado da sua saúde — base para meta e plano de cuidado personalizado.',
};

const tiers = [
  {
    title: 'Versão Completa',
    desc: 'Aplicada pela equipe Plenya durante o acompanhamento. Inclui avaliação clínica, laboratorial, comportamental e funcional. Gera relatório detalhado com metas e plano de cuidado personalizado.',
  },
  {
    title: 'Versão Intermediária',
    desc: 'Disponível após a primeira consulta. Recorte focado nas metas prioritárias identificadas pela equipe médica.',
  },
  {
    title: 'Versão Light',
    desc: 'Disponível online em breve. Permite começar a entender sua saúde de forma acessível, sem necessidade de consulta prévia.',
  },
];

export default async function ScorePage({ params }: { params: Promise<{ locale: string }> }) {
  const { locale } = await params;
  setRequestLocale(locale);

  return (
    <>
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32">
          <p className="label-upper text-gold mb-6">Diagnóstico</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,5rem)] text-cream max-w-xl">
            Escore Plenya
          </h1>
          <p className="text-cream/70 text-lg mt-6 max-w-lg">
            O Escore Plenya traduz dados clínicos, laboratoriais e de comportamento em um diagnóstico
            estruturado de saúde — base para definir metas e construir seu plano de cuidado.
          </p>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-container section grid lg:grid-cols-2 gap-16 items-center">
          <div className="space-y-8">
            <p className="label-upper text-gold">Como funciona</p>
            <p className="text-petrol/80 text-lg leading-relaxed">
              O escore avalia múltiplas dimensões da saúde — metabólica, cardiovascular, funcional,
              hormonal, comportamental — e integra em uma pontuação única que evolui ao longo do tempo.
              Não é uma nota de aprovação: é um mapa de onde você está e para onde pode ir.
            </p>
            <p className="text-petrol/80 text-lg leading-relaxed">
              A cada reavaliação, o escore reflete as mudanças reais no seu corpo e nos seus hábitos.
              É o jeito Plenya de medir evolução — e de mostrar que o cuidado está funcionando.
            </p>
          </div>

          <div className="relative aspect-[3/4] overflow-hidden">
            <Image
              src="/images/hero-score.jpg"
              alt="Pausa, reflexão e autoavaliação"
              fill
              className="object-cover"
              sizes="(min-width: 1024px) 540px, 100vw"
            />
          </div>
        </div>
      </section>

      <section className="bg-cream-100">
        <div className="site-container section flex justify-center">
          <svg viewBox="0 0 300 300" className="w-64 h-64 md:w-80 md:h-80" aria-hidden>
            <circle cx="150" cy="150" r="140" stroke="#0E3B4F" strokeOpacity="0.08" strokeWidth="1" fill="none" />
            <circle cx="150" cy="150" r="110" stroke="#0E3B4F" strokeOpacity="0.12" strokeWidth="1" fill="none" />
            <circle cx="150" cy="150" r="80" stroke="#0E3B4F" strokeOpacity="0.16" strokeWidth="1" fill="none" />
            <path d="M 150 10 A 140 140 0 1 1 40 260" stroke="#A08456" strokeWidth="2" fill="none" strokeLinecap="round" />
            <text x="150" y="160" textAnchor="middle" fontFamily="'Cormorant Garamond', serif" fontSize="56" fill="#0E3B4F" letterSpacing="-2">78</text>
          </svg>
        </div>
      </section>

      <section className="bg-paper">
        <div className="site-container section">
          <p className="label-upper text-gold mb-10">Versões do Escore</p>
          <div className="grid md:grid-cols-3 gap-8">
            {tiers.map((tier, i) => (
              <div key={tier.title} className="border-t border-petrol/15 pt-8 space-y-4">
                <span className="label-upper text-petrol/40">0{i + 1}</span>
                <h3 className="heading-section text-petrol text-2xl">{tier.title}</h3>
                <p className="text-petrol/70 leading-relaxed">{tier.desc}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-container section grid md:grid-cols-2 gap-8">
          <Link href="/planos" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">Planos</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">Comece seu acompanhamento →</p>
          </Link>
          <Link href="/metodo-agir" className="border-t border-petrol/15 pt-8 space-y-3 group">
            <p className="label-upper text-gold">Método</p>
            <p className="heading-section text-petrol text-xl group-hover:text-gold transition">Conheça o Método AGIR →</p>
          </Link>
        </div>
      </section>
    </>
  );
}
