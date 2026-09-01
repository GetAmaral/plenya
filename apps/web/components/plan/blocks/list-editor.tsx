'use client';

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
import { GripVertical, Plus, Trash2 } from 'lucide-react';

import { Button } from '@/components/ui/button';
import { Tooltip, TooltipContent, TooltipTrigger } from '@/components/ui/tooltip';
import { cn } from '@/lib/utils';

/**
 * A primitiva que carrega o trabalho repetitivo dos editores de bloco.
 *
 * Sem ela seriam seis componentes quase iguais, cada um reimplementando adicionar, remover,
 * reordenar e o teto. Aqui isso mora num lugar só, e cada editor de bloco fica sendo apenas o
 * formulário de UM item.
 *
 * O teto é DURO, não aviso: os números vêm de contagem nos decks reais e do teste de estouro do
 * backend (oito réguas num slide comprovadamente transbordam). Deixar adicionar a nona e avisar
 * depois seria transferir para o médico um problema que já está resolvido.
 *
 * Segue o padrão de `components/scores/VersionItemsEditor.tsx`, que é o uso canônico de dnd-kit
 * nesta base.
 */
export interface ListEditorProps<T> {
  itens: T[];
  onChange: (itens: T[]) => void;
  /** Chave estável do item. Índice não serve: quebra o arraste e o foco. */
  chave: (item: T, i: number) => string;
  render: (item: T, i: number, atualiza: (novo: T) => void) => React.ReactNode;
  novoItem?: () => T;
  rotuloAdicionar?: string;
  teto?: number;
  /** Explica o teto quando ele é atingido, em vez de só desabilitar. */
  motivoDoTeto?: string;
  /** Alguns blocos têm cardinalidade fixa (`two-cards` é sempre 2). */
  semAdicionar?: boolean;
  semRemover?: boolean;
  semArrastar?: boolean;
  vazio?: React.ReactNode;
}

export function ListEditor<T>({
  itens,
  onChange,
  chave,
  render,
  novoItem,
  rotuloAdicionar = 'Adicionar',
  teto,
  motivoDoTeto,
  semAdicionar,
  semRemover,
  semArrastar,
  vazio,
}: ListEditorProps<T>) {
  const sensores = useSensors(
    useSensor(PointerSensor, { activationConstraint: { distance: 4 } }),
    useSensor(KeyboardSensor, { coordinateGetter: sortableKeyboardCoordinates }),
  );
  const ids = itens.map((it, i) => chave(it, i));
  const noTeto = teto != null && itens.length >= teto;

  const aoSoltar = (e: DragEndEvent) => {
    const { active, over } = e;
    if (!over || active.id === over.id) return;
    const de = ids.indexOf(String(active.id));
    const para = ids.indexOf(String(over.id));
    if (de < 0 || para < 0) return;
    onChange(arrayMove(itens, de, para));
  };

  const atualizaEm = (i: number) => (novo: T) => {
    const copia = itens.slice();
    copia[i] = novo;
    onChange(copia);
  };

  const remove = (i: number) => onChange(itens.filter((_, j) => j !== i));

  const corpo = itens.map((item, i) => (
    <ItemArrastavel
      key={ids[i]}
      id={ids[i]}
      podeArrastar={!semArrastar && itens.length > 1}
      onRemover={semRemover ? undefined : () => remove(i)}
    >
      {render(item, i, atualizaEm(i))}
    </ItemArrastavel>
  ));

  return (
    <div className="space-y-2">
      {itens.length === 0 && vazio}

      {semArrastar || itens.length < 2 ? (
        <div className="space-y-2">{corpo}</div>
      ) : (
        <DndContext sensors={sensores} collisionDetection={closestCenter} onDragEnd={aoSoltar}>
          <SortableContext items={ids} strategy={verticalListSortingStrategy}>
            <div className="space-y-2">{corpo}</div>
          </SortableContext>
        </DndContext>
      )}

      {!semAdicionar && novoItem && (
        <BotaoAdicionar
          desabilitado={noTeto}
          motivo={noTeto ? motivoDoTeto : undefined}
          rotulo={rotuloAdicionar}
          onClick={() => onChange([...itens, novoItem()])}
        />
      )}
    </div>
  );
}

function BotaoAdicionar({
  desabilitado,
  motivo,
  rotulo,
  onClick,
}: {
  desabilitado: boolean;
  motivo?: string;
  rotulo: string;
  onClick: () => void;
}) {
  const botao = (
    <Button
      type="button"
      size="sm"
      variant="outline"
      className="h-7 w-full text-xs"
      disabled={desabilitado}
      onClick={onClick}
    >
      <Plus className="mr-1 h-3 w-3" />
      {rotulo}
    </Button>
  );
  if (!desabilitado || !motivo) return botao;
  return (
    <Tooltip>
      {/* O botão desabilitado não dispara eventos: o span é o alvo do tooltip. */}
      <TooltipTrigger asChild>
        <span className="block">{botao}</span>
      </TooltipTrigger>
      <TooltipContent side="bottom" className="max-w-xs text-xs">
        {motivo}
      </TooltipContent>
    </Tooltip>
  );
}

function ItemArrastavel({
  id,
  podeArrastar,
  onRemover,
  children,
}: {
  id: string;
  podeArrastar: boolean;
  onRemover?: () => void;
  children: React.ReactNode;
}) {
  const { attributes, listeners, setNodeRef, transform, transition, isDragging } = useSortable({ id });
  return (
    <div
      ref={setNodeRef}
      style={{ transform: CSS.Transform.toString(transform), transition }}
      className={cn(
        'flex items-start gap-1 rounded-md border bg-card p-2',
        isDragging && 'opacity-60 shadow-sm',
      )}
    >
      {podeArrastar && (
        <button
          type="button"
          className="mt-1 cursor-grab text-muted-foreground hover:text-foreground"
          aria-label="Reordenar"
          {...attributes}
          {...listeners}
        >
          <GripVertical className="h-3.5 w-3.5" />
        </button>
      )}
      <div className="min-w-0 flex-1">{children}</div>
      {onRemover && (
        <Button
          type="button"
          size="icon"
          variant="ghost"
          className="h-6 w-6 shrink-0 text-muted-foreground hover:text-destructive"
          onClick={onRemover}
          aria-label="Remover"
        >
          <Trash2 className="h-3.5 w-3.5" />
        </Button>
      )}
    </div>
  );
}
