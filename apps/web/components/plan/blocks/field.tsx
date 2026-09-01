'use client';

import { Label } from '@/components/ui/label';
import { cn } from '@/lib/utils';

/**
 * Um campo do formulário do bloco: rótulo, controle, ajuda e contador.
 *
 * O contador não é enfeite. O slide tem altura fixa e `overflow:hidden`, então texto demais SOME do
 * PDF sem erro nenhum — o contador é o aviso que chega antes de a medição rodar.
 */
export interface FieldProps {
  label: string;
  htmlFor?: string;
  hint?: string;
  /** A partir daqui o contador fica âmbar. Vem dos tetos observados nos decks reais. */
  limite?: number;
  valor?: string;
  children: React.ReactNode;
}

export function Field({ label, htmlFor, hint, limite, valor, children }: FieldProps) {
  const usado = valor?.length ?? 0;
  const apertado = limite != null && usado > limite;

  return (
    <div className="space-y-1">
      <div className="flex items-baseline justify-between gap-2">
        <Label htmlFor={htmlFor} className="text-xs font-medium">
          {label}
        </Label>
        {limite != null && (
          <span
            className={cn(
              'text-[10px] tabular-nums',
              apertado ? 'font-medium text-amber-700' : 'text-muted-foreground',
            )}
            title={
              apertado
                ? 'acima do que costuma caber; a conferência ao salvar dá a palavra final'
                : undefined
            }
          >
            {usado}/{limite}
          </span>
        )}
      </div>
      {children}
      {hint && <p className="text-[11px] leading-snug text-muted-foreground">{hint}</p>}
    </div>
  );
}
