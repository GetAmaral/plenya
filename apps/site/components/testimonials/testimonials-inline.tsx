import { getTranslations } from 'next-intl/server';
import { Link } from '@/lib/i18n/navigation';
import { getFeaturedTestimonials } from '@/lib/testimonials';
import { TestimonialCard } from './testimonial-card';

export async function TestimonialsInline({
  label,
  title,
  bg = 'bg-cream',
  limit = 3,
}: {
  label?: string;
  title?: string;
  bg?: string;
  limit?: number;
}) {
  const items = await getFeaturedTestimonials(limit);
  if (!items.length) return null;

  const tHome = await getTranslations('home');
  const tCta = await getTranslations('cta');
  const resolvedLabel = label ?? tHome('testimonialsLabel');
  const resolvedTitle = title ?? tHome('testimonialsTitle');

  return (
    <section className={bg}>
      <div className="site-container section">
        <div className="flex items-baseline justify-between gap-6 flex-wrap mb-12">
          <div className="space-y-3">
            <p className="label-upper text-gold">{resolvedLabel}</p>
            <h2 className="heading-section text-petrol text-2xl md:text-4xl">{resolvedTitle}</h2>
          </div>
          <Link
            href="/depoimentos"
            className="label-upper text-gold border-b border-gold/40 pb-1 hover:text-petrol transition"
          >
            {tCta('seeAll')}
          </Link>
        </div>

        <div className="grid gap-8 md:grid-cols-2 lg:grid-cols-3">
          {items.map((t) => (
            <TestimonialCard key={t.slug} testimonial={t} />
          ))}
        </div>
      </div>
    </section>
  );
}
