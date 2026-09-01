'use client';

import { useMemo, useState } from 'react';
import { AlertTriangle, Braces, ChevronDown, Copy, Trash2 } from 'lucide-react';
import type { DeckSlide, PlanDossier } from '@plenya/types';

import { Badge } from '@/components/ui/badge';
import { Button } from '@/components/ui/button';
import { Textarea } from '@/components/ui/textarea';
import { CardsEditor } from '@/components/plan/blocks/cards-editor';
import { RulersEditor, AdicionarReguaDoDossie } from '@/components/plan/blocks/rulers-editor';
import { SlideHeaderFields } from '@/components/plan/blocks/slide-header-fields';
import { SummaryEditor } from '@/components/plan/blocks/summary-editor';
import { TableEditor } from '@/components/plan/blocks/table-editor';
import { TakeawayEditor } from '@/components/plan/blocks/takeaway-editor';
import { SLIDE_SPEC } from '@/lib/plan/slide-spec';
import { orcamentoDoSlide } from '@/lib/plan/budget';
import { cn } from '@/lib/utils';
import { SlideThumbnail } from './slide-thumbnail';

/**
 * Um slide como cartão: miniatura à esquerda, campos do bloco à direita quando expandido.
 *
 * Só um cartão fica expandido por vez, e isso não é só arrumação: com vinte cartões controlados no
 * mesmo estado, manter os inputs de todos vivos faz cada tecla re-renderizar a tela inteira.
 *
 * A escotilha de JSON fica, e é requisito. A tela de hoje SÓ tem JSON; os formulários cobrem menos
 * do que o JSON cobre (`sequence`, campos que venham depois), e entregar sem a escotilha seria
 * regressão de capacidade disfarçada de melhoria.
 */
export interface SlideCardProps {
  slide: DeckSlide;
  indice: number;
  total: number;
  expandido: boolean;
  onExpandir: () => void;
  onChange: (s: DeckSlide) => void;
  onRemover: () => void;
  onDuplicar: () => void;
  dossier?: PlanDossier;
  /** Transbordo medido pelo servidor no último salvamento. É a única verdade geométrica. */
  estouro?: { right?: number; bottom?: number };
  sujo?: boolean;
}

export function SlideCard({
  slide,
  indice,
  total,
  expandido,
  onExpandir,
  onChange,
  onRemover,
  onDuplicar,
  dossier,
  estouro,
  sujo,
}: SlideCardProps) {
  const [jsonAberto, setJsonAberto] = useState(false);
  const [jsonTexto, setJsonTexto] = useState('');
  const [jsonErro, setJsonErro] = useState('');

  const spec = SLIDE_SPEC[slide.kind];
  const orcamento = useMemo(() => orcamentoDoSlide(slide), [slide]);
  const patch = (p: Partial<DeckSlide>) => onChange({ ...slide, ...p });

  const abreJson = () => {
    setJsonTexto(JSON.stringify(slide, null, 2));
    setJsonErro('');
    setJsonAberto(true);
  };
  const aplicaJson = () => {
    try {
      const novo = JSON.parse(jsonTexto) as DeckSlide;
      if (!novo.kind) throw new Error('slide sem "kind"');
      // O id não é editável por aqui: é o alvo de toda operação e sugestão.
      onChange({ ...novo, id: slide.id });
      setJsonAberto(false);
      setJsonErro('');
    } catch (e) {
      setJsonErro(e instanceof Error ? e.message : 'JSON inválido');
    }
  };

  const transborda = (estouro?.bottom ?? 0) > 0 || (estouro?.right ?? 0) > 0;

  return (
    <div
      className={cn(
        'rounded-lg border bg-card',
        transborda && 'border-l-2 border-l-destructive',
        expandido && 'ring-1 ring-primary/30',
      )}
    >
      {/*
        Região clicável como div com `role="button"`, e não como <button>.
        Um <button> não pode conter conteúdo de fluxo, e o cabeçalho do cartão tem miniatura e
        selos — mais de vinte divs aninhadas. O HTML inválido fazia o clique não chegar ao
        manipulador, e o cartão simplesmente não abria. Só apareceu na verificação visual: o
        typecheck e o render estavam limpos.
      */}
      <div
        role="button"
        tabIndex={0}
        onClick={onExpandir}
        onKeyDown={(e) => {
          if (e.key === 'Enter' || e.key === ' ') {
            e.preventDefault();
            onExpandir();
          }
        }}
        className="flex w-full cursor-pointer items-start gap-3 p-3 text-left focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring"
        aria-expanded={expandido}
      >
        <span className="w-6 shrink-0 pt-1 text-right text-[11px] tabular-nums text-muted-foreground">
          {String(indice + 1).padStart(2, '0')}
        </span>
        <SlideThumbnail slide={slide} largura={168} altura={96} />
        <span className="min-w-0 flex-1">
          <span className="block truncate text-[10px] font-semibold uppercase tracking-wide text-muted-foreground">
            {slide.eyebrow || spec?.label || slide.kind}
          </span>
          <span className="mt-0.5 block line-clamp-2 text-sm font-medium">
            {slide.title || <span className="text-muted-foreground">sem título</span>}
          </span>
          {slide.punch && (
            <span className="mt-0.5 block truncate text-[11px] italic text-muted-foreground">
              {slide.punch.replace(/<[^>]+>/g, '')}
            </span>
          )}
        </span>
        <span className="flex shrink-0 flex-col items-end gap-1">
          <Badge variant="outline" className="h-5 px-1.5 text-[10px] font-normal">
            {spec?.label ?? slide.kind}
          </Badge>
          {sujo && (
            <Badge variant="secondary" className="h-5 px-1.5 text-[10px] font-normal">
              alterado
            </Badge>
          )}
          {transborda && (
            <Badge variant="destructive" className="h-5 gap-1 px-1.5 text-[10px] font-normal">
              <AlertTriangle className="h-2.5 w-2.5" />
              não cabe
            </Badge>
          )}
          {!transborda && orcamento.nivel !== 'ok' && (
            <span
              className="text-[10px] text-amber-700"
              title={`estimativa: ~${orcamento.usado}px de ${orcamento.disponivel}px, o maior é ${orcamento.maiorBloco}. A conferência ao salvar dá a palavra final.`}
            >
              {orcamento.nivel === 'apertado' ? 'apertado' : 'talvez não caiba'}
            </span>
          )}
          <ChevronDown
            className={cn('h-3.5 w-3.5 text-muted-foreground transition-transform', expandido && 'rotate-180')}
          />
        </span>
      </div>

      {expandido && (
        <div className="space-y-4 border-t p-3">
          {spec?.hint && <p className="text-[11px] text-muted-foreground">{spec.hint}</p>}

          {spec?.readOnly ? (
            <p className="rounded border border-dashed p-3 text-[11px] text-muted-foreground">
              Este tipo de slide não tem formulário: os dois decks reais nunca o usaram. Edite pelo
              JSON abaixo se precisar dele.
            </p>
          ) : (
            <>
              {spec?.blocks.includes('header') && <SlideHeaderFields slide={slide} onChange={patch} />}

              {spec?.blocks.includes('rulers') && (
                <div className="space-y-2">
                  <p className="text-xs font-medium">Réguas</p>
                  <RulersEditor
                    rulers={slide.rulers ?? []}
                    onChange={(rulers) => patch({ rulers })}
                    dossier={dossier}
                    teto={spec.ceilings?.rulers}
                  />
                  <AdicionarReguaDoDossie
                    dossier={dossier}
                    desabilitado={(slide.rulers?.length ?? 0) >= (spec.ceilings?.rulers ?? 4)}
                    motivo="Quatro é o teto comprovado: com oito o slide transborda."
                    onAdd={(r) => patch({ rulers: [...(slide.rulers ?? []), r] })}
                  />
                </div>
              )}

              {spec?.blocks.includes('summary') && (
                <SummaryEditor
                  summary={slide.summary ?? {}}
                  onChange={(summary) => patch({ summary })}
                  dossier={dossier}
                  tetoCartoes={spec.ceilings?.cards}
                  tetoLinhas={spec.ceilings?.lines}
                />
              )}

              {spec?.blocks.includes('cards') && (
                <div className="space-y-2">
                  <p className="text-xs font-medium">Cartões</p>
                  <CardsEditor
                    cards={slide.cards ?? []}
                    onChange={(cards) => patch({ cards })}
                    teto={spec.ceilings?.cards}
                    fixo={slide.kind === 'two-cards'}
                  />
                </div>
              )}

              {spec?.blocks.includes('table') && (
                <div className="space-y-2">
                  <p className="text-xs font-medium">Tabela</p>
                  <TableEditor
                    table={slide.table ?? {}}
                    onChange={(table) => patch({ table })}
                    tetoLinhas={spec.ceilings?.rows}
                    tetoColunas={spec.ceilings?.cols}
                  />
                </div>
              )}

              {spec?.blocks.includes('takeaway') && (
                <div className="space-y-2">
                  <p className="text-xs font-medium">Para levar</p>
                  <TakeawayEditor
                    take={slide.takeaway ?? {}}
                    onChange={(takeaway) => patch({ takeaway })}
                    tetoGrupos={spec.ceilings?.groups}
                  />
                </div>
              )}
            </>
          )}

          <div className="space-y-1 border-t pt-3">
            <Button
              type="button"
              size="sm"
              variant="ghost"
              className="h-6 px-1 text-[11px] text-muted-foreground"
              onClick={() => (jsonAberto ? setJsonAberto(false) : abreJson())}
            >
              <Braces className="mr-1 h-3 w-3" />
              JSON deste slide
            </Button>
            {jsonAberto && (
              <div className="space-y-1">
                <Textarea
                  value={jsonTexto}
                  onChange={(e) => setJsonTexto(e.target.value)}
                  rows={12}
                  className="font-mono text-[11px]"
                />
                {jsonErro && <p className="text-[11px] text-destructive">{jsonErro}</p>}
                <Button type="button" size="sm" className="h-7 text-xs" onClick={aplicaJson}>
                  Aplicar
                </Button>
              </div>
            )}
          </div>

          <div className="flex justify-end gap-1">
            <Button type="button" size="sm" variant="ghost" className="h-7 text-xs" onClick={onDuplicar}>
              <Copy className="mr-1 h-3 w-3" />
              Duplicar
            </Button>
            <Button
              type="button"
              size="sm"
              variant="ghost"
              className="h-7 text-xs text-destructive hover:text-destructive"
              onClick={onRemover}
              disabled={total <= 1}
            >
              <Trash2 className="mr-1 h-3 w-3" />
              Remover
            </Button>
          </div>
        </div>
      )}
    </div>
  );
}
