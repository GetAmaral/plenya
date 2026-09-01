'use client';

import { ArrowRight, Check, ShieldQuestion, X } from 'lucide-react';
import type { PlanSuggestion } from '@plenya/types';

import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';

/**
 * Uma sugestão dentro do cartão do slide que ela altera.
 *
 * A ORIGEM DO NÚMERO fica ao lado do botão, e isso é a decisão de desenho mais importante desta
 * tela. A evidência sobre revisão de conteúdo redigido por IA é dura: mais da metade dos clínicos
 * não pega os erros plantados e uma parte grande envia sem editar, com quase todos declarando
 * confiar na ferramenta. Aceite sem a fonte à vista é carimbo.
 *
 * Também mostra QUANTAS origens candidatas existem. Um número que aparece em cinquenta lugares do
 * dossiê é um número comum, e a origem exibida é palpite — dizer isso vale mais do que esconder.
 */
export function SuggestionStrip({
  sugestao,
  onAceitar,
  onDescartar,
  ocupado,
}: {
  sugestao: PlanSuggestion;
  onAceitar: () => void;
  onDescartar: () => void;
  ocupado?: boolean;
}) {
  const antes = textoDe(sugestao.oldValue);
  const depois = textoDe(sugestao.newValue);
  const provas = sugestao.provenance ?? [];
  const semOrigem = provas.some((p) => !p.found);

  return (
    <div className="rounded-md border-l-2 border-l-amber-500 bg-amber-50/60 p-2 dark:bg-amber-950/20">
      <div className="flex items-center gap-1.5">
        <Badge variant="outline" className="h-5 px-1.5 text-[10px] font-normal">
          {sugestao.class === 'structural' ? 'estrutura' : 'número'}
        </Badge>
        <span className="truncate text-[11px] text-muted-foreground">{sugestao.path || sugestao.op}</span>
      </div>

      <p className="mt-1 text-[11px] text-muted-foreground">{sugestao.rationale}</p>

      {(antes || depois) && (
        <p className="mt-1.5 text-xs leading-relaxed">
          {antes && <span className="text-muted-foreground line-through">{antes}</span>}
          {antes && <ArrowRight className="mx-1 inline h-3 w-3 text-muted-foreground" />}
          <span className="font-medium">{depois}</span>
        </p>
      )}

      {provas.length > 0 && (
        <ul className="mt-1.5 space-y-0.5">
          {provas.map((p, i) => {
            const m = p.matches?.[0];
            const outras = Math.max((p.matches?.length ?? 0) - 1, 0);
            return (
              <li key={i} className="flex items-start gap-1 text-[11px]">
                <ShieldQuestion
                  className={`mt-0.5 h-3 w-3 shrink-0 ${p.found ? 'text-emerald-600' : 'text-destructive'}`}
                />
                <span className="min-w-0">
                  <span className="font-medium tabular-nums">{p.numeral}</span>{' '}
                  {p.found && m ? (
                    <>
                      vem de <span className="text-muted-foreground">{m.label}</span>
                      {outras > 0 && (
                        <span
                          className="text-muted-foreground"
                          title="este número aparece em vários lugares do prontuário; a origem mostrada é a mais provável, não a certa"
                        >
                          {' '}· e mais {outras} possíve{outras > 1 ? 'is' : 'l'}
                        </span>
                      )}
                    </>
                  ) : (
                    <span className="text-destructive">não existe no prontuário deste paciente</span>
                  )}
                </span>
              </li>
            );
          })}
        </ul>
      )}

      {semOrigem && (
        <p className="mt-1 text-[11px] font-medium text-destructive">
          Há número sem origem no prontuário. Confira antes de aceitar.
        </p>
      )}

      <div className="mt-2 flex justify-end gap-1">
        <Button
          type="button"
          size="sm"
          variant="ghost"
          className="h-6 px-2 text-[11px]"
          onClick={onDescartar}
          disabled={ocupado}
        >
          <X className="mr-1 h-3 w-3" />
          Descartar
        </Button>
        <Button type="button" size="sm" className="h-6 px-2 text-[11px]" onClick={onAceitar} disabled={ocupado}>
          <Check className="mr-1 h-3 w-3" />
          Aceitar
        </Button>
      </div>
    </div>
  );
}

function textoDe(v: unknown): string {
  if (v == null) return '';
  if (typeof v === 'string') return v;
  return JSON.stringify(v);
}
