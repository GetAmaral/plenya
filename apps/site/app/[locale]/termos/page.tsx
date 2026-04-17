import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';

export const metadata: Metadata = {
  title: 'Termos de Uso',
  description: 'Termos de uso do site Plenya.',
};

export default async function TermsPage({ params }: { params: Promise<{ locale: string }> }) {
  const { locale } = await params;
  setRequestLocale(locale);
  return (
    <>
      <section className="bg-petrol text-cream">
        <div className="site-narrow pt-32 pb-20 md:pt-40 md:pb-24">
          <p className="label-upper text-gold mb-6">Legal</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,4rem)] text-cream">Termos de Uso</h1>
        </div>
      </section>
      <section className="bg-cream">
        <div className="site-narrow section text-petrol/80 text-lg">
          <p className="label-upper text-petrol/40">Documento em redação.</p>
        </div>
      </section>
    </>
  );
}
