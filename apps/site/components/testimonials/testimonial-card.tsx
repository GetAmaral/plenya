import type { Testimonial } from '@/lib/testimonials';

export function TestimonialCard({ testimonial }: { testimonial: Testimonial; variant?: 'default' | 'compact' }) {
  const score = testimonial.plenyaScore;
  const hasScoreEvolution =
    score && typeof score.initial === 'number' && typeof score.current === 'number';

  return (
    <figure className="border-t border-petrol/15 pt-8 space-y-6">
      <blockquote className="heading-section text-petrol text-xl leading-snug">
        “{testimonial.quote}”
      </blockquote>

      {testimonial.outcome && (
        <p className="label-upper text-gold">{testimonial.outcome}</p>
      )}

      {hasScoreEvolution && (
        <div className="border-t border-petrol/10 pt-4 flex items-baseline gap-3">
          <span className="label-upper text-petrol/50">Plenya Score</span>
          <span className="font-mono text-petrol/40 text-lg tabular-nums">{score.initial}</span>
          <span className="text-petrol/30 text-sm" aria-hidden>→</span>
          <span className="font-mono text-gold text-2xl tabular-nums leading-none">
            {score.current}
          </span>
          {score.months && (
            <span className="label-upper text-petrol/50 ml-auto">
              {score.months} {score.months === 1 ? 'mês' : 'meses'}
            </span>
          )}
        </div>
      )}

      <figcaption className="space-y-1">
        <p className="label-upper text-petrol">{testimonial.patientLabel}</p>
        {testimonial.context && <p className="text-petrol/60 text-sm">{testimonial.context}</p>}
      </figcaption>
    </figure>
  );
}
