/**
 * Cores estáveis por médico — derivadas via hash do UUID, então o mesmo médico
 * mantém a mesma cor entre sessões/dispositivos sem precisar persistir nada.
 *
 * Paleta escolhida pra contraste em fundo claro/escuro com Tailwind utilities.
 */

const PALETTE = [
  {
    badge: 'border-rose-300 bg-rose-50 text-rose-700 dark:bg-rose-950/40 dark:text-rose-200',
    block: 'border-l-rose-500 bg-rose-50/70 text-rose-900 dark:bg-rose-950/40 dark:text-rose-100',
    dot: 'bg-rose-500',
  },
  {
    badge: 'border-amber-300 bg-amber-50 text-amber-700 dark:bg-amber-950/40 dark:text-amber-200',
    block: 'border-l-amber-500 bg-amber-50/70 text-amber-900 dark:bg-amber-950/40 dark:text-amber-100',
    dot: 'bg-amber-500',
  },
  {
    badge: 'border-emerald-300 bg-emerald-50 text-emerald-700 dark:bg-emerald-950/40 dark:text-emerald-200',
    block: 'border-l-emerald-500 bg-emerald-50/70 text-emerald-900 dark:bg-emerald-950/40 dark:text-emerald-100',
    dot: 'bg-emerald-500',
  },
  {
    badge: 'border-sky-300 bg-sky-50 text-sky-700 dark:bg-sky-950/40 dark:text-sky-200',
    block: 'border-l-sky-500 bg-sky-50/70 text-sky-900 dark:bg-sky-950/40 dark:text-sky-100',
    dot: 'bg-sky-500',
  },
  {
    badge: 'border-violet-300 bg-violet-50 text-violet-700 dark:bg-violet-950/40 dark:text-violet-200',
    block: 'border-l-violet-500 bg-violet-50/70 text-violet-900 dark:bg-violet-950/40 dark:text-violet-100',
    dot: 'bg-violet-500',
  },
  {
    badge: 'border-fuchsia-300 bg-fuchsia-50 text-fuchsia-700 dark:bg-fuchsia-950/40 dark:text-fuchsia-200',
    block: 'border-l-fuchsia-500 bg-fuchsia-50/70 text-fuchsia-900 dark:bg-fuchsia-950/40 dark:text-fuchsia-100',
    dot: 'bg-fuchsia-500',
  },
  {
    badge: 'border-teal-300 bg-teal-50 text-teal-700 dark:bg-teal-950/40 dark:text-teal-200',
    block: 'border-l-teal-500 bg-teal-50/70 text-teal-900 dark:bg-teal-950/40 dark:text-teal-100',
    dot: 'bg-teal-500',
  },
  {
    badge: 'border-orange-300 bg-orange-50 text-orange-700 dark:bg-orange-950/40 dark:text-orange-200',
    block: 'border-l-orange-500 bg-orange-50/70 text-orange-900 dark:bg-orange-950/40 dark:text-orange-100',
    dot: 'bg-orange-500',
  },
] as const;

function hashString(s: string): number {
  let h = 5381;
  for (let i = 0; i < s.length; i++) {
    h = ((h << 5) + h) ^ s.charCodeAt(i);
  }
  return h >>> 0;
}

export function doctorColorClasses(doctorId: string) {
  return PALETTE[hashString(doctorId) % PALETTE.length];
}

export function doctorBlockClass(doctorId: string) {
  return doctorColorClasses(doctorId).block;
}

export function doctorDotClass(doctorId: string) {
  return doctorColorClasses(doctorId).dot;
}
