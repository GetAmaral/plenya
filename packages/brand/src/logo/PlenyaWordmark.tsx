import * as React from 'react';

type WordmarkProps = React.HTMLAttributes<HTMLSpanElement> & {
  withTagline?: boolean;
};

export function PlenyaWordmark({ withTagline = false, className, ...props }: WordmarkProps) {
  return (
    <span className={['inline-flex flex-col leading-none', className].filter(Boolean).join(' ')} {...props}>
      <span className="font-heading text-2xl tracking-[0.4em] uppercase">PLENYA</span>
      {withTagline ? (
        <span className="font-mono text-[10px] tracking-[0.25em] uppercase mt-1 text-current/70">
          Saúde, Performance &amp; Longevidade
        </span>
      ) : null}
    </span>
  );
}
