import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';

export const metadata: Metadata = {
  title: 'Política de Privacidade',
  description: 'Política de privacidade da Plenya em conformidade com a LGPD.',
};

export default async function PrivacyPage({ params }: { params: Promise<{ locale: string }> }) {
  const { locale } = await params;
  setRequestLocale(locale);
  return (
    <>
      <section className="bg-petrol text-cream">
        <div className="site-narrow pt-32 pb-20 md:pt-40 md:pb-24">
          <p className="label-upper text-gold mb-6">LGPD</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,4rem)] text-cream">Política de Privacidade</h1>
        </div>
      </section>
      <section className="bg-cream">
        <div className="site-narrow section space-y-6 text-petrol/80 text-lg leading-relaxed">
          <p>
            Sua privacidade é sagrada. Utilizamos analytics sem cookies (Plausible) e nenhum dado pessoal
            é compartilhado com terceiros. Conformidade total com a Lei Geral de Proteção de Dados (LGPD).
          </p>
          <p className="label-upper text-petrol/40">Documento completo em redação.</p>
        </div>
      </section>
    </>
  );
}
