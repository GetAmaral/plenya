import type { Testimonial } from '@/lib/testimonials';

export function TestimonialCard({ testimonial }: { testimonial: Testimonial; variant?: 'default' | 'compact' }) {
  return (
    <figure className="border-t border-petrol/15 pt-8 space-y-6">
      <blockquote className="heading-section text-petrol text-xl leading-snug">
        "{testimonial.quote}"
      </blockquote>
      {testimonial.outcome && (
        <p className="label-upper text-gold">{testimonial.outcome}</p>
      )}
      <figcaption className="space-y-1">
        <p className="label-upper text-petrol">{testimonial.patientLabel}</p>
        {testimonial.context && <p className="text-petrol/60 text-sm">{testimonial.context}</p>}
      </figcaption>
    </figure>
  );
}
