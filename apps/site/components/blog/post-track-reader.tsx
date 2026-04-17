'use client';

import { useEffect } from 'react';
import { track } from '@/lib/plausible';

export function PostTrackReader({ slug, pillar }: { slug: string; pillar: string }) {
  useEffect(() => {
    let read = false;
    const start = Date.now();

    const onScroll = () => {
      if (read) return;
      const scrolled = window.scrollY + window.innerHeight;
      const total = document.documentElement.scrollHeight;
      if (scrolled / total >= 0.7 && Date.now() - start >= 60_000) {
        read = true;
        track('blog_post_lido', { slug, pillar });
        window.removeEventListener('scroll', onScroll);
      }
    };

    window.addEventListener('scroll', onScroll, { passive: true });
    return () => window.removeEventListener('scroll', onScroll);
  }, [slug, pillar]);

  return null;
}
