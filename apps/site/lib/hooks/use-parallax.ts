'use client';

import { useEffect, useRef } from 'react';

/**
 * Subtle parallax on element (translateY based on scroll progress through viewport).
 * Strength: 0.1-0.3 = subtle editorial. Avoid >0.5 (too much, feels cheap).
 */
export function useParallax<T extends HTMLElement>(strength = 0.15) {
  const ref = useRef<T | null>(null);

  useEffect(() => {
    const el = ref.current;
    if (!el) return;
    if (window.matchMedia('(prefers-reduced-motion: reduce)').matches) return;

    let raf = 0;
    const update = () => {
      const rect = el.getBoundingClientRect();
      const center = rect.top + rect.height / 2;
      const distance = center - window.innerHeight / 2;
      el.style.transform = `translate3d(0, ${distance * -strength}px, 0)`;
    };
    const onScroll = () => {
      cancelAnimationFrame(raf);
      raf = requestAnimationFrame(update);
    };

    update();
    window.addEventListener('scroll', onScroll, { passive: true });
    return () => {
      cancelAnimationFrame(raf);
      window.removeEventListener('scroll', onScroll);
    };
  }, [strength]);

  return ref;
}
