import * as React from 'react';

type SymbolProps = React.SVGAttributes<SVGSVGElement> & {
  title?: string;
};

/**
 * Plenya monogram — the official "P" symbol from the brandbook.
 * Vertical spine with descender, horizontal cross, rounded bowl on the right.
 * Renders as currentColor so callers control the fill via Tailwind text-* classes.
 */
export function PlenyaSymbol({ title = 'Plenya', ...props }: SymbolProps) {
  return (
    <svg
      viewBox="0 0 100 100"
      fill="none"
      stroke="currentColor"
      strokeWidth={5}
      strokeLinecap="round"
      strokeLinejoin="round"
      role="img"
      aria-label={title}
      {...props}
    >
      <title>{title}</title>
      <line x1="32" y1="6" x2="32" y2="94" />
      <line x1="4" y1="62" x2="96" y2="62" />
      <path d="M32 6 H58 A28 28 0 0 1 58 62 H32" />
    </svg>
  );
}
