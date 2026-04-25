'use client';

import { useState } from 'react';
import { ChevronDown } from 'lucide-react';
import { cn } from '@/lib/cn';

export type FaqItem = {
  q: string;
  a: React.ReactNode;
};

export function FaqAccordion({
  label = 'Perguntas frequentes',
  title,
  items,
}: {
  label?: string;
  title: string;
  items: FaqItem[];
}) {
  const [open, setOpen] = useState<number | null>(0);

  return (
    <section className="bg-cream">
      <div className="site-container section">
        <div className="max-w-3xl mb-12 space-y-4">
          <p className="label-upper text-gold">{label}</p>
          <h2 className="heading-section text-petrol text-3xl md:text-4xl">{title}</h2>
        </div>

        <ul className="max-w-3xl border-t border-petrol/15">
          {items.map((it, i) => {
            const isOpen = open === i;
            return (
              <li key={it.q} className="border-b border-petrol/15">
                <button
                  type="button"
                  onClick={() => setOpen(isOpen ? null : i)}
                  aria-expanded={isOpen}
                  className="w-full text-left flex items-center justify-between gap-6 py-6 group"
                >
                  <span
                    className={cn(
                      'heading-section text-petrol text-lg md:text-xl leading-snug transition group-hover:text-gold',
                      isOpen && 'text-gold',
                    )}
                  >
                    {it.q}
                  </span>
                  <ChevronDown
                    size={20}
                    className={cn(
                      'flex-shrink-0 text-petrol/50 transition-transform duration-300',
                      isOpen && 'rotate-180 text-gold',
                    )}
                  />
                </button>
                <div
                  className={cn(
                    'grid transition-all duration-300 ease-out',
                    isOpen ? 'grid-rows-[1fr] pb-8' : 'grid-rows-[0fr]',
                  )}
                >
                  <div className="overflow-hidden">
                    <div className="text-petrol/75 text-base leading-relaxed max-w-2xl pr-8">
                      {it.a}
                    </div>
                  </div>
                </div>
              </li>
            );
          })}
        </ul>
      </div>
    </section>
  );
}
