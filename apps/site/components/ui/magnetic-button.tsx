'use client';

import { useMagnetic } from '@/lib/hooks/use-magnetic';
import { cn } from '@/lib/cn';

type Props = React.AnchorHTMLAttributes<HTMLAnchorElement> & {
  variant?: 'primary' | 'outline' | 'outline-cream';
  strength?: number;
  children: React.ReactNode;
};

export function MagneticLink({
  variant = 'primary',
  strength = 0.18,
  className,
  children,
  ...props
}: Props) {
  const ref = useMagnetic<HTMLAnchorElement>(strength);
  const variantClass =
    variant === 'primary'
      ? 'btn-primary'
      : variant === 'outline-cream'
        ? 'btn-outline border-cream/40 text-cream hover:bg-cream hover:text-petrol'
        : 'btn-outline';
  return (
    <a
      ref={ref}
      className={cn(variantClass, 'transition-transform duration-200 ease-out will-change-transform', className)}
      {...props}
    >
      {children}
    </a>
  );
}
