'use client';

import { useReveal } from '@/lib/hooks/use-reveal';
import { cn } from '@/lib/cn';

type Props = {
  as?: 'div' | 'section' | 'article' | 'p' | 'h1' | 'h2' | 'h3';
  delay?: number;
  className?: string;
  children: React.ReactNode;
};

export function Reveal({ as: Tag = 'div', delay = 0, className, children }: Props) {
  const { ref, visible } = useReveal<HTMLDivElement>();
  return (
    <Tag
      ref={ref as never}
      style={{ transitionDelay: `${delay}ms` }}
      className={cn(
        'transition-all duration-700 will-change-transform',
        'motion-reduce:transition-none',
        visible
          ? 'opacity-100 translate-y-0 blur-0'
          : 'opacity-0 translate-y-4 blur-[2px]',
        className,
      )}
    >
      {children}
    </Tag>
  );
}
