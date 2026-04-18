import * as React from 'react';

type SymbolProps = React.SVGAttributes<SVGSVGElement> & {
  title?: string;
};

/**
 * Plenya monogram — vector geometry extracted from the official brandbook PDF.
 * Two filled paths (descender + spine/bowl/cross). Renders as currentColor so
 * callers control the fill via Tailwind text-* classes.
 */
export function PlenyaSymbol({ title = 'Plenya', ...props }: SymbolProps) {
  return (
    <svg
      viewBox="0 0 157.828 122.696"
      fill="currentColor"
      role="img"
      aria-label={title}
      {...props}
    >
      <title>{title}</title>
      <path d="M 71.148438 109.097656 L 71.148438 116.183594 C 71.148438 119.621094 68.375 122.394531 64.9375 122.394531 C 61.515625 122.394531 58.742188 119.621094 58.742188 116.183594 L 58.742188 96.105469 L 71.121094 96.105469 L 71.121094 108.464844 C 71.132812 108.675781 71.148438 108.886719 71.148438 109.097656" />
      <path d="M 157.546875 44.984375 C 157.546875 82.199219 128.09375 89.984375 104.480469 89.984375 L 71.144531 89.984375 L 71.144531 90.007812 L 0.109375 90.007812 L 0.109375 85.363281 L 78.808594 85.363281 L 78.808594 85.351562 L 104.480469 85.351562 C 135.722656 85.351562 145.285156 64.421875 145.285156 44.984375 C 145.285156 25.703125 135.722656 4.632812 104.480469 4.632812 L 71.144531 4.632812 L 71.144531 79.332031 L 58.742188 79.34375 L 58.742188 0 L 104.480469 0 C 128.09375 0 157.546875 7.773438 157.546875 44.984375" />
    </svg>
  );
}
