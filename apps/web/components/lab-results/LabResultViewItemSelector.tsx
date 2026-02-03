'use client'

import { useState, useMemo } from 'react'
import { Input } from '@/components/ui/input'
import { ScrollArea } from '@/components/ui/scroll-area'
import { GripVertical, Search } from 'lucide-react'
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

export interface LabTestDefinition {
  id: string
  name: string
  code: string
  category?: string
}

interface LabResultViewItemSelectorProps {
  availableTests: LabTestDefinition[]
  selectedTests: LabTestDefinition[]
  onSelectionChange: (selected: LabTestDefinition[]) => void
}

interface SortableItemProps {
  test: LabTestDefinition
  onClick: () => void
}

function SortableItem({ test, onClick }: SortableItemProps) {
  const { attributes, listeners, setNodeRef, transform, transition, isDragging } = useSortable({
    id: test.id,
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
      className="flex items-center gap-2 p-3 border rounded-lg bg-card hover:bg-accent cursor-pointer group"
      onClick={onClick}
    >
      <div
        {...attributes}
        {...listeners}
        className="cursor-grab active:cursor-grabbing flex-shrink-0"
        onClick={(e) => e.stopPropagation()}
      >
        <GripVertical className="h-4 w-4 text-muted-foreground" />
      </div>
      <div className="flex-1 min-w-0">
        <div className="font-medium break-words">{test.name}</div>
        {test.code && (
          <div className="text-xs text-muted-foreground">{test.code}</div>
        )}
      </div>
    </div>
  )
}

/**
 * Agrupa testes por categoria
 */
function organizeTestsByCategory(tests: LabTestDefinition[]) {
  const categories = new Map<string, LabTestDefinition[]>()

  tests.forEach((test) => {
    const category = test.category || 'Sem Categoria'
    if (!categories.has(category)) {
      categories.set(category, [])
    }
    categories.get(category)!.push(test)
  })

  // Ordenar testes dentro de cada categoria
  categories.forEach((tests) => {
    tests.sort((a, b) => a.name.localeCompare(b.name))
  })

  // Converter para array e ordenar categorias
  return Array.from(categories.entries())
    .sort(([a], [b]) => a.localeCompare(b))
    .map(([category, tests]) => ({ category, tests }))
}

export function LabResultViewItemSelector({
  availableTests,
  selectedTests,
  onSelectionChange,
}: LabResultViewItemSelectorProps) {
  const [searchAvailable, setSearchAvailable] = useState('')
  const [searchSelected, setSearchSelected] = useState('')

  const sensors = useSensors(
    useSensor(PointerSensor),
    useSensor(KeyboardSensor, {
      coordinateGetter: sortableKeyboardCoordinates,
    })
  )

  // Filtrar testes disponíveis (não selecionados)
  const available = useMemo(() => {
    const selectedIds = new Set(selectedTests.map((t) => t.id))
    return availableTests.filter((test) => !selectedIds.has(test.id))
  }, [availableTests, selectedTests])

  // Busca nos testes disponíveis
  const filteredAvailable = useMemo(() => {
    if (!searchAvailable) return available
    const query = searchAvailable.toLowerCase()
    return available.filter(
      (test) =>
        test.name.toLowerCase().includes(query) ||
        test.code?.toLowerCase().includes(query)
    )
  }, [available, searchAvailable])

  // Busca nos testes selecionados
  const filteredSelected = useMemo(() => {
    if (!searchSelected) return selectedTests
    const query = searchSelected.toLowerCase()
    return selectedTests.filter(
      (test) =>
        test.name.toLowerCase().includes(query) ||
        test.code?.toLowerCase().includes(query)
    )
  }, [selectedTests, searchSelected])

  // Organizar testes disponíveis por categoria
  const categorizedAvailable = useMemo(
    () => organizeTestsByCategory(filteredAvailable),
    [filteredAvailable]
  )

  // Adicionar teste
  const handleAddTest = (test: LabTestDefinition) => {
    onSelectionChange([...selectedTests, test])
  }

  // Adicionar categoria inteira
  const handleAddCategory = (tests: LabTestDefinition[]) => {
    onSelectionChange([...selectedTests, ...tests])
  }

  // Remover teste
  const handleRemoveTest = (testId: string) => {
    onSelectionChange(selectedTests.filter((t) => t.id !== testId))
  }

  // Reordenar testes selecionados
  const handleDragEnd = (event: DragEndEvent) => {
    const { active, over } = event

    if (over && active.id !== over.id) {
      const oldIndex = selectedTests.findIndex((t) => t.id === active.id)
      const newIndex = selectedTests.findIndex((t) => t.id === over.id)

      onSelectionChange(arrayMove(selectedTests, oldIndex, newIndex))
    }
  }

  return (
    <div className="grid grid-cols-2 gap-4">
      {/* Coluna Esquerda: Testes Disponíveis */}
      <div className="space-y-3">
        <div>
          <h3 className="text-sm font-semibold mb-2">Exames Disponíveis</h3>
          <div className="relative">
            <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
            <Input
              placeholder="Buscar exames..."
              value={searchAvailable}
              onChange={(e) => setSearchAvailable(e.target.value)}
              className="pl-10"
            />
          </div>
        </div>

        <ScrollArea className="h-[500px] rounded-lg border p-3">
          <div className="space-y-4">
            {categorizedAvailable.map(({ category, tests }) => (
              <div key={category} className="space-y-2">
                <div
                  className="font-semibold text-sm text-primary cursor-pointer hover:underline"
                  onClick={() => handleAddCategory(tests)}
                >
                  {category} ({tests.length})
                </div>
                <div className="space-y-2">
                  {tests.map((test) => (
                    <div
                      key={test.id}
                      className="p-2 border rounded-lg bg-card hover:bg-accent cursor-pointer"
                      onClick={() => handleAddTest(test)}
                    >
                      <div className="font-medium text-sm">{test.name}</div>
                      {test.code && (
                        <div className="text-xs text-muted-foreground">
                          {test.code}
                        </div>
                      )}
                    </div>
                  ))}
                </div>
              </div>
            ))}
            {categorizedAvailable.length === 0 && (
              <div className="text-center text-sm text-muted-foreground py-8">
                {searchAvailable
                  ? 'Nenhum exame encontrado'
                  : 'Todos os exames foram selecionados'}
              </div>
            )}
          </div>
        </ScrollArea>
      </div>

      {/* Coluna Direita: Testes Selecionados (Drag-and-Drop) */}
      <div className="space-y-3">
        <div>
          <h3 className="text-sm font-semibold mb-2">
            Exames Selecionados ({selectedTests.length})
          </h3>
          <div className="relative">
            <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
            <Input
              placeholder="Buscar selecionados..."
              value={searchSelected}
              onChange={(e) => setSearchSelected(e.target.value)}
              className="pl-10"
            />
          </div>
        </div>

        <ScrollArea className="h-[500px] rounded-lg border p-3">
          <DndContext
            sensors={sensors}
            collisionDetection={closestCenter}
            onDragEnd={handleDragEnd}
          >
            <SortableContext
              items={filteredSelected.map((t) => t.id)}
              strategy={verticalListSortingStrategy}
            >
              <div className="space-y-2">
                {filteredSelected.map((test) => (
                  <SortableItem
                    key={test.id}
                    test={test}
                    onClick={() => handleRemoveTest(test.id)}
                  />
                ))}
              </div>
            </SortableContext>
          </DndContext>

          {filteredSelected.length === 0 && (
            <div className="text-center text-sm text-muted-foreground py-8">
              {searchSelected
                ? 'Nenhum exame encontrado'
                : 'Nenhum exame selecionado'}
            </div>
          )}
        </ScrollArea>
      </div>
    </div>
  )
}
