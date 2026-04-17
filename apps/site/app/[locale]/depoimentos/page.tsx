import type { Metadata } from 'next';
import { setRequestLocale } from 'next-intl/server';
import { isLocale, defaultLocale } from '@/lib/i18n/config';
import { Link } from '@/lib/i18n/navigation';
import { getAllTestimonials } from '@/lib/testimonials';
import { TestimonialCard } from '@/components/testimonials/testimonial-card';

export const metadata: Metadata = {
  title: 'Depoimentos',
  description: 'Histórias reais de pacientes Plenya — todas com consentimento e identificação preservada.',
};

export default async function TestimonialsPage({ params }: { params: Promise<{ locale: string }> }) {
  const { locale: rawLocale } = await params;
  const locale = isLocale(rawLocale) ? rawLocale : defaultLocale;
  setRequestLocale(locale);

  const items = await getAllTestimonials();

  return (
    <>
      <section className="bg-petrol text-cream">
        <div className="site-container pt-32 pb-24 md:pt-40 md:pb-32">
          <p className="label-upper text-gold mb-6">Pacientes Plenya</p>
          <h1 className="heading-hero text-[clamp(2.5rem,6vw,5rem)] text-cream max-w-xl">Depoimentos</h1>
          <p className="text-cream/70 text-lg mt-6 max-w-lg">
            Histórias reais de quem decidiu cuidar com método. Identificações preservadas conforme
            consentimento assinado, em conformidade com a LGPD.
          </p>
        </div>
      </section>

      <section className="bg-cream">
        <div className="site-container section">
          {items.length ? (
            <div className="grid gap-8 md:grid-cols-2 lg:grid-cols-3">
              {items.map((t) => (
                <TestimonialCard key={t.slug} testimonial={t} />
              ))}
            </div>
          ) : (
            <p className="text-petrol/50 text-center label-upper">Em breve, primeiros depoimentos.</p>
          )}
        </div>
      </section>

      <section className="bg-paper">
        <div className="site-container section text-center space-y-6">
          <p className="heading-section text-petrol text-2xl md:text-3xl">
            Quer começar a sua história?
          </p>
          <Link href="/contato" className="btn-gold">Fale com a equipe Plenya</Link>
        </div>
      </section>
    </>
  );
}
