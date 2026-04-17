import * as React from 'react';

type SymbolProps = React.SVGAttributes<SVGSVGElement> & {
  title?: string;
};

export function PlenyaSymbol({ title = 'Plenya', ...props }: SymbolProps) {
  return (
    <svg
      viewBox="0 0 120 60"
      fill="none"
      stroke="currentColor"
      strokeWidth="1.5"
      strokeLinecap="round"
      strokeLinejoin="round"
      role="img"
      aria-label={title}
      {...props}
    >
      <title>{title}</title>
      <path d="M30 30 m -22 0 a 22 22 0 1 0 44 0 a 22 22 0 1 0 -44 0" />
      <path d="M90 30 m -22 0 a 22 22 0 1 0 44 0 a 22 22 0 1 0 -44 0" />
      <path d="M30 12 L30 48" />
      <path d="M52 30 L68 30" />
    </svg>
  );
}
