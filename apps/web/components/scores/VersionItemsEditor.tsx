'use client'

import { useMemo, useState } from 'react'
import { Input } from '@/components/ui/input'
import { ScrollArea } from '@/components/ui/scroll-area'
import { GripVertical, Search, X, Plus } from 'lucide-react'
import {
  DndContext,
  closestCenter,
  KeyboardSensor,
  PointerSensor,
  useSensor,
  useSensors,
  DragEndEvent,
} from '@dnd-kit/core'
import {
  arrayMove,
  SortableContext,
  sortableKeyboardCoordinates,
  useSortable,
  verticalListSortingStrategy,
} from '@dnd-kit/sortable'
import { CSS } from '@dnd-kit/utilities'
import type { ScoreItem } from '@/lib/api/score-api'

interface VersionItemsEditorProps {
  availableItems: ScoreItem[]
  selectedItems: ScoreItem[]
  onSelectionChange: (selected: ScoreItem[]) => void
  disabled?: boolean
}

function categoryLabel(item: ScoreItem): string {
  const group = item.subgroup?.group?.name
  const subgroup = item.subgroup?.name
  if (group && subgroup) return `${group} › ${subgroup}`
  return group || subgroup || 'Sem grupo'
}

// Mantém a ordem em que o backend devolve (group.order → subgroup.order → item.order).
function organizeByCategory(items: ScoreItem[]) {
  const categories = new Map<string, ScoreItem[]>()
  for (const item of items) {
    const cat = categoryLabel(item)
    if (!categories.has(cat)) categories.set(cat, [])
    categories.get(cat)!.push(item)
  }
  return Array.from(categories.entries()).map(([category, items]) => ({ category, items }))
}

interface SortableSelectedProps {
  item: ScoreItem
  index: number
  onRemove: () => void
  disabled?: boolean
}

function SortableSelected({ item, index, onRemove, disabled }: SortableSelectedProps) {
  const { attributes, listeners, setNodeRef, transform, transition, isDragging } = useSortable({
    id: item.id,
    disabled,
  })

  const style = {
    transform: CSS.Transform.toString(transform),
    transition,
    opacity: isDragging ? 0.5 : 1,
  }

  return (
    <div
      ref={setNodeRef}
      style={style}
      className="flex items-center gap-2 p-2 border rounded-lg bg-card group"
    >
      <div
        {...attributes}
        {...listeners}
        className={disabled ? 'shrink-0 text-muted-foreground/40' : 'cursor-grab active:cursor-grabbing shrink-0'}
      >
        <GripVertical className="h-4 w-4 text-muted-foreground" />
      </div>
      <span className="w-6 shrink-0 text-xs tabular-nums text-muted-foreground">{index + 1}</span>
      <div className="flex-1 min-w-0">
        <div className="text-sm font-medium wrap-break-word">{item.name}</div>
        <div className="text-xs text-muted-foreground truncate">{categoryLabel(item)}</div>
      </div>
      {!disabled && (
        <button
          type="button"
          onClick={onRemove}
          className="shrink-0 rounded p-1 text-muted-foreground hover:bg-destructive/10 hover:text-destructive"
          title="Remover da versão"
        >
          <X className="h-4 w-4" />
        </button>
      )}
    </div>
  )
}

export function VersionItemsEditor({
  availableItems,
  selectedItems,
  onSelectionChange,
  disabled = false,
}: VersionItemsEditorProps) {
  const [searchAvailable, setSearchAvailable] = useState('')
  const [searchSelected, setSearchSelected] = useState('')

  const sensors = useSensors(
    useSensor(PointerSensor),
    useSensor(KeyboardSensor, { coordinateGetter: sortableKeyboardCoordinates })
  )

  // Disponíveis = todos menos os já selecionados.
  const available = useMemo(() => {
    const selectedIds = new Set(selectedItems.map((i) => i.id))
    return availableItems.filter((i) => !selectedIds.has(i.id))
  }, [availableItems, selectedItems])

  const filteredAvailable = useMemo(() => {
    if (!searchAvailable) return available
    const q = searchAvailable.toLowerCase()
    return available.filter(
      (i) => i.name.toLowerCase().includes(q) || categoryLabel(i).toLowerCase().includes(q)
    )
  }, [available, searchAvailable])

  const categorizedAvailable = useMemo(() => organizeByCategory(filteredAvailable), [filteredAvailable])

  const filteredSelected = useMemo(() => {
    if (!searchSelected) return selectedItems
    const q = searchSelected.toLowerCase()
    return selectedItems.filter(
      (i) => i.name.toLowerCase().includes(q) || categoryLabel(i).toLowerCase().includes(q)
    )
  }, [selectedItems, searchSelected])

  const addItem = (item: ScoreItem) => onSelectionChange([...selectedItems, item])
  const addCategory = (items: ScoreItem[]) => onSelectionChange([...selectedItems, ...items])
  const removeItem = (id: string) => onSelectionChange(selectedItems.filter((i) => i.id !== id))

  // Reordena só quando a busca dos selecionados está vazia (índices batem com a lista real).
  const handleDragEnd = (event: DragEndEvent) => {
    const { active, over } = event
    if (!over || active.id === over.id) return
    const oldIndex = selectedItems.findIndex((i) => i.id === active.id)
    const newIndex = selectedItems.findIndex((i) => i.id === over.id)
    if (oldIndex === -1 || newIndex === -1) return
    onSelectionChange(arrayMove(selectedItems, oldIndex, newIndex))
  }

  const reorderDisabled = disabled || !!searchSelected

  return (
    <div className="grid gap-4 lg:grid-cols-2">
      {/* Disponíveis */}
      <div className="space-y-3">
        <div>
          <h3 className="text-sm font-semibold mb-2">Itens disponíveis ({available.length})</h3>
          <div className="relative">
            <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
            <Input
              placeholder="Buscar itens..."
              value={searchAvailable}
              onChange={(e) => setSearchAvailable(e.target.value)}
              className="pl-10"
            />
          </div>
        </div>

        <ScrollArea className="h-[560px] rounded-lg border p-3">
          <div className="space-y-4">
            {categorizedAvailable.map(({ category, items }) => (
              <div key={category} className="space-y-2">
                <div className="flex items-center justify-between gap-2">
                  <span className="text-sm font-semibold text-primary">{category}</span>
                  {!disabled && (
                    <button
                      type="button"
                      onClick={() => addCategory(items)}
                      className="shrink-0 inline-flex items-center gap-1 text-xs text-muted-foreground hover:text-primary"
                      title="Adicionar todos deste grupo"
                    >
                      <Plus className="h-3 w-3" /> {items.length}
                    </button>
                  )}
                </div>
                <div className="space-y-1.5">
                  {items.map((item) => (
                    <button
                      key={item.id}
                      type="button"
                      disabled={disabled}
                      onClick={() => addItem(item)}
                      className="w-full text-left p-2 border rounded-lg bg-card hover:bg-accent disabled:cursor-not-allowed disabled:opacity-60"
                    >
                      <div className="text-sm font-medium">{item.name}</div>
                      {typeof item.points === 'number' && (
                        <div className="text-xs text-muted-foreground">{item.points} pts</div>
                      )}
                    </button>
                  ))}
                </div>
              </div>
            ))}
            {categorizedAvailable.length === 0 && (
              <div className="text-center text-sm text-muted-foreground py-8">
                {searchAvailable ? 'Nenhum item encontrado' : 'Todos os itens já estão na versão'}
              </div>
            )}
          </div>
        </ScrollArea>
      </div>

      {/* Selecionados (ordem de exibição) */}
      <div className="space-y-3">
        <div>
          <h3 className="text-sm font-semibold mb-2">Itens da versão ({selectedItems.length})</h3>
          <div className="relative">
            <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
            <Input
              placeholder="Buscar selecionados..."
              value={searchSelected}
              onChange={(e) => setSearchSelected(e.target.value)}
              className="pl-10"
            />
          </div>
          {reorderDisabled && !!searchSelected && (
            <p className="mt-1 text-xs text-muted-foreground">Limpe a busca para reordenar por arraste.</p>
          )}
        </div>

        <ScrollArea className="h-[560px] rounded-lg border p-3">
          <DndContext sensors={sensors} collisionDetection={closestCenter} onDragEnd={handleDragEnd}>
            <SortableContext
              items={filteredSelected.map((i) => i.id)}
              strategy={verticalListSortingStrategy}
            >
              <div className="space-y-1.5">
                {filteredSelected.map((item) => (
                  <SortableSelected
                    key={item.id}
                    item={item}
                    index={selectedItems.findIndex((i) => i.id === item.id)}
                    onRemove={() => removeItem(item.id)}
                    disabled={reorderDisabled}
                  />
                ))}
              </div>
            </SortableContext>
          </DndContext>

          {filteredSelected.length === 0 && (
            <div className="text-center text-sm text-muted-foreground py-8">
              {searchSelected ? 'Nenhum item encontrado' : 'Nenhum item nesta versão. Adicione pela coluna ao lado.'}
            </div>
          )}
        </ScrollArea>
      </div>
    </div>
  )
}
