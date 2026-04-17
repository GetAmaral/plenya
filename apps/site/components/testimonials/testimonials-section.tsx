import { getFeaturedTestimonials } from '@/lib/testimonials';
import { TestimonialCard } from './testimonial-card';

export async function TestimonialsSection() {
  const items = await getFeaturedTestimonials(3);
  if (!items.length) return null;
  return (
    <section className="bg-cream-200">
      <div className="site-container section">
        <p className="label-upper text-gold mb-6">Pacientes Plenya</p>
        <h2 className="heading-section text-petrol text-3xl md:text-5xl max-w-xl mb-16">
          Histórias de quem decidiu cuidar com método.
        </h2>
        <div className="grid md:grid-cols-3 gap-8">
          {items.map((t) => (
            <TestimonialCard key={t.slug} testimonial={t} variant="compact" />
          ))}
        </div>
      </div>
    </section>
  );
}
