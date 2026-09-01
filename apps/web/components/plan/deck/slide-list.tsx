'use client';

import { useState } from 'react';
import {
  DndContext,
  KeyboardSensor,
  PointerSensor,
  closestCenter,
  useSensor,
  useSensors,
  type DragEndEvent,
} from '@dnd-kit/core';
import {
  SortableContext,
  arrayMove,
  sortableKeyboardCoordinates,
  useSortable,
  verticalListSortingStrategy,
} from '@dnd-kit/sortable';
import { CSS } from '@dnd-kit/utilities';
import { GripVertical, Plus } from 'lucide-react';
import type {
  DeckOverflow,
  DeckSlide,
  DeckSlideKind,
  PlanDossier,
  PlanSuggestion,
} from '@plenya/types';

import { Button } from '@/components/ui/button';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';
import { SLIDE_SPEC } from '@/lib/plan/slide-spec';
import { cn } from '@/lib/utils';
import { SlideCard } from './slide-card';

/**
 * A lista de slides do rascunho: arrastar para reordenar, um cartão expandido por vez.
 *
 * O transbordo vem do servidor (`overflow`), não de estimativa: é a única medida geométrica de
 * verdade, e ela chega indexada por POSIÇÃO no deck, que é como o backend a devolve.
 */
export function SlideList({
  slides,
  onChange,
  dossier,
  overflow = [],
  sujos,
  sugestoes = [],
  onAceitarSugestao,
  onDescartarSugestao,
  resolvendo,
}: {
  slides: DeckSlide[];
  onChange: (s: DeckSlide[]) => void;
  dossier?: PlanDossier;
  overflow?: DeckOverflow[];
  sujos?: Set<string>;
  sugestoes?: PlanSuggestion[];
  onAceitarSugestao?: (id: string) => void;
  onDescartarSugestao?: (id: string) => void;
  resolvendo?: boolean;
}) {
  const [expandido, setExpandido] = useState<string | null>(null);

  const sensores = useSensors(
    useSensor(PointerSensor, { activationConstraint: { distance: 4 } }),
    useSensor(KeyboardSensor, { coordinateGetter: sortableKeyboardCoordinates }),
  );
  const ids = slides.map((s, i) => s.id || `pos-${i}`);

  const aoSoltar = (e: DragEndEvent) => {
    const { active, over } = e;
    if (!over || active.id === over.id) return;
    const de = ids.indexOf(String(active.id));
    const para = ids.indexOf(String(over.id));
    if (de < 0 || para < 0) return;
    onChange(arrayMove(slides, de, para));
  };

  const troca = (i: number) => (novo: DeckSlide) => {
    const copia = slides.slice();
    copia[i] = novo;
    onChange(copia);
  };

  const estouroDe = (i: number) => overflow.find((o) => o.slide === i + 1);

  const naoCabem = overflow.length;

  return (
    <div className="space-y-3">
      {naoCabem > 0 && (
        <p className="rounded-md border border-destructive/40 bg-destructive/5 px-3 py-2 text-xs text-destructive">
          {naoCabem} slide{naoCabem > 1 ? 's' : ''} não cabe{naoCabem > 1 ? 'm' : ''} na moldura
          impressa. O que passa some do PDF sem aviso, então a publicação fica bloqueada até caber.
        </p>
      )}

      <DndContext sensors={sensores} collisionDetection={closestCenter} onDragEnd={aoSoltar}>
        <SortableContext items={ids} strategy={verticalListSortingStrategy}>
          <div className="space-y-2">
            {slides.map((s, i) => (
              <SlideArrastavel key={ids[i]} id={ids[i]}>
                <SlideCard
                  slide={s}
                  indice={i}
                  total={slides.length}
                  expandido={expandido === ids[i]}
                  onExpandir={() => setExpandido(expandido === ids[i] ? null : ids[i])}
                  onChange={troca(i)}
                  onRemover={() => onChange(slides.filter((_, j) => j !== i))}
                  onDuplicar={() => {
                    const copia = slides.slice();
                    // Sem id: o servidor atribui um novo ao salvar. Duplicar mantendo o id faria
                    // duas linhas disputarem o mesmo alvo de operação.
                    const { id: _ignora, ...semId } = s;
                    copia.splice(i + 1, 0, semId as DeckSlide);
                    onChange(copia);
                  }}
                  dossier={dossier}
                  estouro={estouroDe(i)}
                  sujo={sujos?.has(ids[i])}
                  sugestoes={sugestoes.filter((g) => g.slideId === s.id)}
                  onAceitarSugestao={onAceitarSugestao}
                  onDescartarSugestao={onDescartarSugestao}
                  resolvendo={resolvendo}
                />
              </SlideArrastavel>
            ))}
          </div>
        </SortableContext>
      </DndContext>

      <AdicionarSlide onAdd={(kind) => onChange([...slides, { kind } as DeckSlide])} />
    </div>
  );
}

function AdicionarSlide({ onAdd }: { onAdd: (kind: DeckSlideKind) => void }) {
  const tipos = Object.entries(SLIDE_SPEC) as [DeckSlideKind, (typeof SLIDE_SPEC)[DeckSlideKind]][];
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button type="button" variant="outline" size="sm" className="w-full text-xs">
          <Plus className="mr-1 h-3 w-3" />
          Adicionar slide
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="start" className="w-56">
        {tipos
          .filter(([, spec]) => !spec.readOnly)
          .map(([kind, spec]) => (
            <DropdownMenuItem key={kind} onSelect={() => onAdd(kind)} className="text-xs">
              {spec.label}
            </DropdownMenuItem>
          ))}
      </DropdownMenuContent>
    </DropdownMenu>
  );
}

function SlideArrastavel({ id, children }: { id: string; children: React.ReactNode }) {
  const { attributes, listeners, setNodeRef, transform, transition, isDragging } = useSortable({ id });
  return (
    <div
      ref={setNodeRef}
      style={{ transform: CSS.Transform.toString(transform), transition }}
      className={cn('relative', isDragging && 'z-10 opacity-70')}
    >
      <button
        type="button"
        className="absolute -left-5 top-4 hidden cursor-grab text-muted-foreground hover:text-foreground lg:block"
        aria-label="Reordenar slide"
        {...attributes}
        {...listeners}
      >
        <GripVertical className="h-4 w-4" />
      </button>
      {children}
    </div>
  );
}
