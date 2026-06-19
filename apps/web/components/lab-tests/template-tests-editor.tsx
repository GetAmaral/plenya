'use client'

import { useMemo, useState } from 'react'
import {
  DndContext,
  closestCenter,
  PointerSensor,
  KeyboardSensor,
  useSensor,
  useSensors,
  type DragEndEvent,
} from '@dnd-kit/core'
import {
  SortableContext,
  arrayMove,
  useSortable,
  sortableKeyboardCoordinates,
  verticalListSortingStrategy,
} from '@dnd-kit/sortable'
import { CSS } from '@dnd-kit/utilities'
import { Input } from '@/components/ui/input'
import { Card } from '@/components/ui/card'
import { Switch } from '@/components/ui/switch'
import { ScrollArea } from '@/components/ui/scroll-area'
import { Search, Plus, X, GripVertical, ScrollText } from 'lucide-react'
import { LabTestDefinition } from '@/lib/api/lab-request-templates'

// Remove acentos (unaccent) p/ busca.
function normalizeForSearch(str: string): string {
  return str.normalize('NFD').replace(/[̀-ͯ]/g, '').toLowerCase()
}

function matchesSearch(t: LabTestDefinition, q: string): boolean {
  if (!q) return true
  const s = normalizeForSearch(q)
  if (normalizeForSearch(t.name).includes(s)) return true
  if (t.shortName && normalizeForSearch(t.shortName).includes(s)) return true
  if (t.altNames && t.altNames.some((a) => normalizeForSearch(a).includes(s))) return true
  return false
}

interface Props {
  availableTests: LabTestDefinition[]
  // selectedTests: NA ORDEM do template; cada item carrega pageBreakBefore (quebra antes dele).
  selectedTests: LabTestDefinition[]
  onSelectionChange: (selected: LabTestDefinition[]) => void
}

export function TemplateTestsEditor({ availableTests, selectedTests, onSelectionChange }: Props) {
  const [searchLeft, setSearchLeft] = useState('')
  const [searchRight, setSearchRight] = useState('')

  const sensors = useSensors(
    useSensor(PointerSensor, { activationConstraint: { distance: 5 } }),
    useSensor(KeyboardSensor, { coordinateGetter: sortableKeyboardCoordinates }),
  )

  // Disponíveis = catálogo menos os já selecionados, alfabético.
  const unselected = useMemo(() => {
    const ids = new Set(selectedTests.map((t) => t.id))
    return availableTests
      .filter((t) => !ids.has(t.id))
      .filter((t) => matchesSearch(t, searchLeft))
      .sort((a, b) => a.name.localeCompare(b.name, 'pt-BR'))
  }, [availableTests, selectedTests, searchLeft])

  // Selecionados ficam SEMPRE na ordem do array (= ordem do template). Busca só filtra a exibição;
  // se houver busca à direita, desabilita o drag pra não reordenar um subconjunto.
  const rightFiltered = useMemo(
    () => selectedTests.filter((t) => matchesSearch(t, searchRight)),
    [selectedTests, searchRight],
  )
  const dragEnabled = searchRight.trim() === ''

  const addTest = (t: LabTestDefinition) =>
    onSelectionChange([...selectedTests, { ...t, pageBreakBefore: false }])

  const removeTest = (id: string) =>
    onSelectionChange(selectedTests.filter((t) => t.id !== id))

  const togglePageBreak = (id: string) =>
    onSelectionChange(
      selectedTests.map((t) => (t.id === id ? { ...t, pageBreakBefore: !t.pageBreakBefore } : t)),
    )

  const handleDragEnd = (e: DragEndEvent) => {
    const { active, over } = e
    if (!over || active.id === over.id) return
    const from = selectedTests.findIndex((t) => t.id === active.id)
    const to = selectedTests.findIndex((t) => t.id === over.id)
    if (from < 0 || to < 0) return
    onSelectionChange(arrayMove(selectedTests, from, to))
  }

  return (
    <div className="grid grid-cols-[1fr_1.3fr] gap-4">
      {/* Disponíveis */}
      <Card className="p-4">
        <h3 className="text-sm font-medium mb-2">Exames Disponíveis</h3>
        <div className="relative mb-3">
          <Search className="absolute left-2 top-2.5 h-4 w-4 text-muted-foreground" />
          <Input
            placeholder="Buscar exames..."
            value={searchLeft}
            onChange={(e) => setSearchLeft(e.target.value)}
            className="pl-8"
          />
        </div>
        <ScrollArea className="h-[520px]">
          <div className="space-y-1 pr-2">
            {unselected.length === 0 ? (
              <p className="text-sm text-muted-foreground text-center py-8">
                {searchLeft ? 'Nenhum exame encontrado' : 'Todos os exames já foram selecionados'}
              </p>
            ) : (
              unselected.map((t) => (
                <button
                  key={t.id}
                  type="button"
                  onClick={() => addTest(t)}
                  className="w-full text-left p-2 rounded-md hover:bg-accent flex items-start gap-2 group"
                >
                  <Plus className="h-4 w-4 mt-0.5 text-muted-foreground group-hover:text-foreground shrink-0" />
                  <span className="flex-1">
                    <span className="block text-sm font-medium">{t.name}</span>
                    {t.tussCode && (
                      <span className="block text-xs text-muted-foreground">TUSS: {t.tussCode}</span>
                    )}
                  </span>
                </button>
              ))
            )}
          </div>
        </ScrollArea>
        <p className="text-xs text-muted-foreground mt-2">{unselected.length} disponível(is)</p>
      </Card>

      {/* Selecionados — ordenáveis */}
      <Card className="p-4">
        <h3 className="text-sm font-medium mb-1">Exames Selecionados (em ordem)</h3>
        <p className="text-xs text-muted-foreground mb-2">
          Arraste para reordenar. A chave <span className="font-medium">Nova página</span> faz o
          exame começar em uma nova folha no PDF.
        </p>
        <div className="relative mb-3">
          <Search className="absolute left-2 top-2.5 h-4 w-4 text-muted-foreground" />
          <Input
            placeholder="Filtrar selecionados..."
            value={searchRight}
            onChange={(e) => setSearchRight(e.target.value)}
            className="pl-8"
          />
        </div>
        {!dragEnabled && (
          <p className="text-xs text-amber-600 mb-2">
            Limpe o filtro para reordenar (o arraste fica travado durante a busca).
          </p>
        )}
        <ScrollArea className="h-[470px]">
          <DndContext sensors={sensors} collisionDetection={closestCenter} onDragEnd={handleDragEnd}>
            <SortableContext
              items={rightFiltered.map((t) => t.id)}
              strategy={verticalListSortingStrategy}
            >
              <div className="space-y-1 pr-2">
                {rightFiltered.length === 0 ? (
                  <p className="text-sm text-muted-foreground text-center py-8">
                    {searchRight ? 'Nenhum exame encontrado' : 'Nenhum exame selecionado'}
                  </p>
                ) : (
                  rightFiltered.map((t, i) => (
                    <SortableRow
                      key={t.id}
                      test={t}
                      index={selectedTests.findIndex((s) => s.id === t.id)}
                      dragEnabled={dragEnabled}
                      onRemove={() => removeTest(t.id)}
                      onTogglePageBreak={() => togglePageBreak(t.id)}
                    />
                  ))
                )}
              </div>
            </SortableContext>
          </DndContext>
        </ScrollArea>
        <p className="text-xs text-muted-foreground mt-2">
          {selectedTests.length} selecionado(s)
        </p>
      </Card>
    </div>
  )
}

function SortableRow({
  test,
  index,
  dragEnabled,
  onRemove,
  onTogglePageBreak,
}: {
  test: LabTestDefinition
  index: number
  dragEnabled: boolean
  onRemove: () => void
  onTogglePageBreak: () => void
}) {
  const { attributes, listeners, setNodeRef, transform, transition, isDragging } = useSortable({
    id: test.id,
    disabled: !dragEnabled,
  })
  const style = { transform: CSS.Transform.toString(transform), transition }
  const just = (test.requestJustification || '').trim()

  return (
    <div
      ref={setNodeRef}
      style={style}
      className={`rounded-md border bg-card p-2 ${isDragging ? 'opacity-60 shadow-lg' : ''} ${
        test.pageBreakBefore ? 'border-t-2 border-t-ocean' : ''
      }`}
    >
      <div className="flex items-start gap-2">
        <button
          type="button"
          className={`mt-0.5 shrink-0 text-muted-foreground ${dragEnabled ? 'cursor-grab' : 'cursor-not-allowed opacity-40'}`}
          {...attributes}
          {...listeners}
          aria-label="Arrastar para reordenar"
        >
          <GripVertical className="h-4 w-4" />
        </button>
        <span className="text-xs text-muted-foreground tabular-nums mt-0.5 w-6 shrink-0">
          {index + 1}.
        </span>
        <div className="flex-1 min-w-0">
          <div className="text-sm font-medium">{test.name}</div>
          {test.tussCode && (
            <div className="text-xs text-muted-foreground">TUSS: {test.tussCode}</div>
          )}
          {just ? (
            <div className="mt-1 flex items-start gap-1 text-xs text-muted-foreground">
              <ScrollText className="h-3 w-3 mt-0.5 shrink-0" />
              <span className="line-clamp-2" title={just}>
                {just}
              </span>
            </div>
          ) : (
            <div className="mt-1 text-xs text-muted-foreground/60 italic">sem justificativa</div>
          )}
        </div>
        <div className="flex flex-col items-end gap-1 shrink-0">
          <button
            type="button"
            onClick={onRemove}
            className="text-muted-foreground hover:text-destructive"
            aria-label="Remover"
          >
            <X className="h-4 w-4" />
          </button>
          <label className="flex items-center gap-1 text-[11px] text-muted-foreground cursor-pointer select-none">
            Nova página
            <Switch checked={!!test.pageBreakBefore} onCheckedChange={onTogglePageBreak} />
          </label>
        </div>
      </div>
    </div>
  )
}
