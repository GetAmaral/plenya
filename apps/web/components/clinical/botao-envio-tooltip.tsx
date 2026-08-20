'use client';

/**
 * Envolve um botão de envio para que, quando ele estiver desabilitado por falta de dado do
 * paciente, o motivo apareça em tooltip.
 *
 * Existe porque botão desabilitado NÃO dispara evento de mouse: o tooltip precisa ficar num
 * elemento em volta, senão nunca abre — que é justamente quando ele mais importa.
 */

import type { ReactNode } from 'react';
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from '@/components/ui/tooltip';

interface Props {
  /** Quando vazio, o botão está habilitado e não há nada a explicar. */
  motivo?: string;
  children: ReactNode;
}

export function BotaoEnvioTooltip({ motivo, children }: Props) {
  if (!motivo) return <>{children}</>;

  return (
    <TooltipProvider>
      <Tooltip>
        <TooltipTrigger asChild>
          {/* O span recebe o hover que o botão desabilitado não recebe. */}
          <span tabIndex={0} className="inline-flex">
            {children}
          </span>
        </TooltipTrigger>
        <TooltipContent side="top">{motivo}</TooltipContent>
      </Tooltip>
    </TooltipProvider>
  );
}
