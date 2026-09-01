'use client';

import type { DeckSlide } from '@plenya/types';

import { Slide } from '@/components/patient-portal/plan-deck';
import { cn } from '@/lib/utils';

/**
 * A miniatura do cartão, reusando o renderizador do portal.
 *
 * Duas coisas precisam estar claras, e as duas são armadilhas.
 *
 * A PRIMEIRA é técnica: a tipografia do renderizador usa `clamp(..., 5vw, ...)`, ancorada na
 * VIEWPORT e não no contêiner. Simplesmente estreitar a caixa daria título gigante dentro de uma
 * miniatura minúscula. Por isso o conteúdo é renderizado numa largura FIXA e depois escalado por
 * transform: o `clamp` resolve uma vez naquela largura e a redução é uniforme.
 *
 * A SEGUNDA é de significado, e é a que engana: isto é a renderização do PORTAL, fluida e
 * empilhada, e NÃO a moldura de 1920×1080 do PDF. Um slide pode parecer perfeito aqui e transbordar
 * na impressão, onde o conteúdo que não cabe some sem erro. Por isso a legenda diz "como o paciente
 * vê na tela", e quem responde pelo encaixe é a medição do servidor.
 */
const LARGURA_BASE = 720;

export function SlideThumbnail({
  slide,
  largura = 240,
  altura = 140,
  className,
}: {
  slide: DeckSlide;
  largura?: number;
  altura?: number;
  className?: string;
}) {
  const escala = largura / LARGURA_BASE;

  return (
    <div
      className={cn('relative shrink-0 overflow-hidden rounded border bg-white', className)}
      style={{ width: largura, height: altura }}
      aria-hidden
    >
      <div
        // Sem eventos: o conteúdo escalado não é clicável nem lido por leitor de tela. A
        // acessibilidade do cartão fica no título, que é texto de verdade.
        className="pointer-events-none origin-top-left select-none"
        style={{ width: LARGURA_BASE, transform: `scale(${escala})` }}
      >
        <Slide slide={slide} />
      </div>
      {/* Esmaece o corte em vez de cortar seco, para ficar claro que há mais embaixo. */}
      <div className="pointer-events-none absolute inset-x-0 bottom-0 h-6 bg-gradient-to-t from-white to-transparent" />
    </div>
  );
}
