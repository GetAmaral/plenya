'use client'

import { useState, useMemo } from 'react'
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { ScrollArea } from '@/components/ui/scroll-area'
import { ChevronRight, ChevronLeft, GripVertical } from 'lucide-react'
import { ScoreItem } from '@/lib/api/anamnesis-templates'
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

interface AnamnesisTemplateItemSelectorProps {
  availableScoreItems: ScoreItem[]
  selectedScoreItems: ScoreItem[]
  onSelectionChange: (selected: ScoreItem[]) => void
}

interface SortableItemInBlockProps {
  item: ScoreItem
  onRemove: () => void
}

function SortableItemInBlock({ item, onRemove }: SortableItemInBlockProps) {
  const isChild = !!item.parentItemId

  // Items filhos não são sortable
  const sortable = useSortable({
    id: item.id,
    disabled: isChild,
  })

  const { attributes, listeners, setNodeRef, transform, transition, isDragging } = sortable

  const style = isChild
    ? {}
    : {
        transform: CSS.Transform.toString(transform),
        transition,
        opacity: isDragging ? 0.5 : 1,
      }

  return (
    <div
      ref={setNodeRef}
      style={{
        ...style,
        marginLeft: isChild ? '20px' : undefined,
        borderLeft: isChild ? '4px solid hsl(var(--primary))' : undefined,
        paddingLeft: isChild ? '12px' : undefined,
      }}
      className="flex items-center gap-2 p-3 border rounded-lg bg-card hover:bg-accent cursor-pointer group"
      onClick={onRemove}
    >
      {/* Só mostra drag handle se NÃO for filho */}
      {!isChild && (
        <div
          {...attributes}
          {...listeners}
          className="cursor-grab active:cursor-grabbing flex-shrink-0"
          onClick={(e) => e.stopPropagation()}
        >
          <GripVertical className="h-4 w-4 text-muted-foreground" />
        </div>
      )}
      <div className="flex-1 min-w-0">
        <div className="font-medium break-words">{item.name}</div>
      </div>
    </div>
  )
}

interface SortableBlockProps {
  block: { id: string; groupId: string; subgroupId: string; items: ScoreItem[] }
  onRemoveItem: (item: ScoreItem) => void
}

function SortableBlockHeader({ blockId, type, children }: { blockId: string; type: 'group' | 'subgroup'; children: React.ReactNode }) {
  const { attributes, listeners, setNodeRef, transform, transition, isDragging } = useSortable({
    id: blockId,
  })

  const style = {
    transform: CSS.Transform.toString(transform),
    transition,
    opacity: isDragging ? 0.5 : 1,
  }

  return (
    <div ref={setNodeRef} style={style}>
      {type === 'group' ? (
        <div className="bg-primary text-primary-foreground rounded-lg px-3 py-2">
          <div className="flex items-center gap-2 min-w-0">
            <div {...attributes} {...listeners} className="cursor-grab active:cursor-grabbing flex-shrink-0">
              <GripVertical className="h-4 w-4 text-primary-foreground" />
            </div>
            {children}
          </div>
        </div>
      ) : (
        <div className="bg-gradient-to-br from-muted to-muted/50 rounded-lg p-2 border border-primary/20">
          <div className="flex items-center gap-2 min-w-0">
            <div {...attributes} {...listeners} className="cursor-grab active:cursor-grabbing flex-shrink-0">
              <GripVertical className="h-4 w-4" />
            </div>
            {children}
          </div>
        </div>
      )}
    </div>
  )
}

function SortableBlock({ block, onRemoveItem, showGroupHeader }: SortableBlockProps & { showGroupHeader: boolean }) {
  const firstItem = block.items[0]
  const groupName = firstItem?.subgroup?.group?.name
  const subgroupName = firstItem?.subgroup?.name

  // IDs diferentes para grupo e subgrupo
  const groupHeaderId = `block-group-${block.groupId}`
  const subgroupHeaderId = block.id // block-groupId-subgroupId-index

  return (
    <div className="space-y-2">
      {/* Cabeçalho do Grupo - só mostra se for o primeiro bloco do grupo */}
      {showGroupHeader && groupName && (
        <SortableBlockHeader blockId={groupHeaderId} type="group">
          <h4 className="text-sm font-bold break-words min-w-0 flex-1">{groupName}</h4>
        </SortableBlockHeader>
      )}

      {/* Cabeçalho do Subgrupo */}
      {subgroupName && (
        <SortableBlockHeader blockId={subgroupHeaderId} type="subgroup">
          <h5 className="text-xs font-bold break-words min-w-0 flex-1">{subgroupName}</h5>
        </SortableBlockHeader>
      )}

      {/* Items do bloco */}
      <div className="space-y-2">
        {block.items.map((item) => (
          <SortableItemInBlock
            key={item.id}
            item={item}
            onRemove={() => onRemoveItem(item)}
          />
        ))}
      </div>
    </div>
  )
}

interface SortableItemProps {
  item: ScoreItem
  onRemove: () => void
  showGroup: boolean
  showSubgroup: boolean
}

function SortableItem({ item, onRemove, showGroup, showSubgroup }: SortableItemProps) {
  const { attributes, listeners, setNodeRef, transform, transition, isDragging } = useSortable({
    id: item.id,
  })

  const style = {
    transform: CSS.Transform.toString(transform),
    transition,
    opacity: isDragging ? 0.5 : 1,
  }

  const groupName = item.subgroup?.group?.name
  const subgroupName = item.subgroup?.name
  const isChild = !!item.parentItemId

  return (
    <div ref={setNodeRef} style={style} className="space-y-2">
      {/* Cabeçalho do Grupo */}
      {showGroup && groupName && (
        <div className="bg-primary text-primary-foreground rounded-lg px-3 py-2">
          <div className="flex items-center gap-2 min-w-0">
            <div {...attributes} {...listeners} className="cursor-grab active:cursor-grabbing flex-shrink-0">
              <GripVertical className="h-4 w-4 text-primary-foreground" />
            </div>
            <h4 className="text-sm font-bold break-words min-w-0 flex-1">{groupName}</h4>
          </div>
        </div>
      )}

      {/* Cabeçalho do Subgrupo */}
      {showSubgroup && subgroupName && (
        <div className="bg-gradient-to-br from-muted to-muted/50 rounded-lg p-2 border border-primary/20">
          <div className="flex items-center gap-2 min-w-0">
            <h5 className="text-xs font-bold break-words min-w-0 flex-1">{subgroupName}</h5>
          </div>
        </div>
      )}

      {/* Item */}
      <div
        className="flex items-center gap-2 p-3 border rounded-lg bg-card hover:bg-accent cursor-pointer group"
        style={{
          marginLeft: isChild ? '20px' : undefined,
          borderLeft: isChild ? '4px solid hsl(var(--primary))' : undefined,
          paddingLeft: isChild ? '12px' : undefined,
        }}
        onClick={onRemove}
      >
        <div {...attributes} {...listeners} className="cursor-grab active:cursor-grabbing flex-shrink-0">
          <GripVertical className="h-4 w-4 text-muted-foreground" />
        </div>
        <div className="flex-1 min-w-0">
          <div className="font-medium break-words">{item.name}</div>
        </div>
      </div>
    </div>
  )
}

// Organiza items em estrutura hierárquica por grupo > subgrupo > item
interface GroupedItems {
  groupId: string
  groupName: string
  groupOrder: number
  subgroups: {
    subgroupId: string
    subgroupName: string
    subgroupOrder: number
    items: ScoreItem[]
  }[]
}

// Produz lista plana respeitando hierarquia pai→filho:
// items raiz ordenados por order, filhos de cada pai inseridos logo após ele (também por order)
function flattenItemsHierarchy(items: ScoreItem[]): ScoreItem[] {
  const childrenMap = new Map<string, ScoreItem[]>()
  const rootItems: ScoreItem[] = []

  items.forEach((item) => {
    if (item.parentItemId) {
      const children = childrenMap.get(item.parentItemId) || []
      children.push(item)
      childrenMap.set(item.parentItemId, children)
    } else {
      rootItems.push(item)
    }
  })

  childrenMap.forEach((children) => {
    children.sort((a, b) => a.order - b.order || a.name.localeCompare(b.name))
  })

  rootItems.sort((a, b) => a.order - b.order || a.name.localeCompare(b.name))

  const flatten = (list: ScoreItem[]): ScoreItem[] =>
    list.flatMap((item) => [item, ...flatten(childrenMap.get(item.id) || [])])

  return flatten(rootItems)
}

function organizeItemsByHierarchy(items: ScoreItem[]): GroupedItems[] {
  const grouped = new Map<string, GroupedItems>()

  items.forEach((item) => {
    if (!item.subgroup?.group) return

    const groupId = item.subgroup.group.id
    const subgroupId = item.subgroup.id

    // Cria grupo se não existir
    if (!grouped.has(groupId)) {
      grouped.set(groupId, {
        groupId,
        groupName: item.subgroup.group.name,
        groupOrder: item.subgroup.group.order,
        subgroups: [],
      })
    }

    const group = grouped.get(groupId)!

    // Cria subgrupo se não existir
    let subgroup = group.subgroups.find((s) => s.subgroupId === subgroupId)
    if (!subgroup) {
      subgroup = {
        subgroupId,
        subgroupName: item.subgroup.name,
        subgroupOrder: item.subgroup.order,
        items: [],
      }
      group.subgroups.push(subgroup)
    }

    // Adiciona item ao subgrupo
    subgroup.items.push(item)
  })

  // Ordena grupos, subgrupos e items (respeitando hierarquia pai→filho)
  return Array.from(grouped.values())
    .sort((a, b) => a.groupOrder - b.groupOrder)
    .map((group) => ({
      ...group,
      subgroups: group.subgroups
        .sort((a, b) => a.subgroupOrder - b.subgroupOrder)
        .map((subgroup) => ({
          ...subgroup,
          items: flattenItemsHierarchy(subgroup.items),
        })),
    }))
}

export function AnamnesisTemplateItemSelector({
  availableScoreItems,
  selectedScoreItems,
  onSelectionChange,
}: AnamnesisTemplateItemSelectorProps) {
  const [searchAvailable, setSearchAvailable] = useState('')
  const [searchSelected, setSearchSelected] = useState('')

  const sensors = useSensors(
    useSensor(PointerSensor),
    useSensor(KeyboardSensor, {
      coordinateGetter: sortableKeyboardCoordinates,
    })
  )

  // Calcula itens não selecionados
  const unselectedItems = useMemo(() => {
    const selectedIds = new Set(selectedScoreItems.map((item) => item.id))
    return availableScoreItems.filter((item) => !selectedIds.has(item.id))
  }, [availableScoreItems, selectedScoreItems])

  // Filtra e organiza itens disponíveis
  const filteredAndGroupedAvailable = useMemo(() => {
    let items = unselectedItems

    // Aplica filtro de busca
    if (searchAvailable) {
      const search = searchAvailable.toLowerCase()
      items = items.filter((item) => item.name.toLowerCase().includes(search))
    }

    return organizeItemsByHierarchy(items)
  }, [unselectedItems, searchAvailable])

  // Filtra itens selecionados
  const filteredSelected = useMemo(() => {
    if (!searchSelected) return selectedScoreItems
    const search = searchSelected.toLowerCase()
    return selectedScoreItems.filter((item) => item.name.toLowerCase().includes(search))
  }, [selectedScoreItems, searchSelected])

  const addItem = (item: ScoreItem) => {
    const itemsToAdd: ScoreItem[] = []

    // Se for filho, garantir que o pai seja adicionado primeiro
    if (item.parentItemId) {
      const parent = availableScoreItems.find((i) => i.id === item.parentItemId)
      if (parent && !selectedScoreItems.some((s) => s.id === parent.id)) {
        itemsToAdd.push(parent)
      }
    }

    itemsToAdd.push(item)

    // Buscar filhos do item (se for pai)
    const children = availableScoreItems.filter(
      (availableItem) => availableItem.parentItemId === item.id
    )
    if (children.length > 0) {
      itemsToAdd.push(...children)
    }

    // Filtrar items que já estão selecionados
    const newItems = itemsToAdd.filter(
      (itemToAdd) => !selectedScoreItems.some((selected) => selected.id === itemToAdd.id)
    )

    onSelectionChange([...selectedScoreItems, ...newItems])
  }

  const addGroup = (group: GroupedItems) => {
    // Adicionar todos os items de todos os subgrupos do grupo
    const allGroupItems = group.subgroups.flatMap((subgroup) => subgroup.items)
    const itemsToAdd = allGroupItems.filter(
      (item) => !selectedScoreItems.some((selected) => selected.id === item.id)
    )
    onSelectionChange([...selectedScoreItems, ...itemsToAdd])
  }

  const addSubgroup = (subgroupItems: ScoreItem[]) => {
    // Adicionar todos os items do subgrupo
    const itemsToAdd = subgroupItems.filter(
      (item) => !selectedScoreItems.some((selected) => selected.id === item.id)
    )
    onSelectionChange([...selectedScoreItems, ...itemsToAdd])
  }

  const removeItem = (item: ScoreItem) => {
    onSelectionChange(selectedScoreItems.filter((i) => i.id !== item.id))
  }

  // Agrupa items consecutivos por grupo/subgrupo em blocos
  const groupItemsIntoBlocks = (items: ScoreItem[]) => {
    const blocks: { id: string; groupId: string; subgroupId: string; items: ScoreItem[] }[] = []

    items.forEach((item) => {
      const groupId = item.subgroup?.group?.id || 'no-group'
      const subgroupId = item.subgroup?.id || 'no-subgroup'

      // Verificar se o último bloco tem o mesmo grupo/subgrupo
      const lastBlock = blocks[blocks.length - 1]
      if (lastBlock && lastBlock.groupId === groupId && lastBlock.subgroupId === subgroupId) {
        lastBlock.items.push(item)
      } else {
        // Criar novo bloco com prefixo 'block-'
        blocks.push({
          id: `block-${groupId}-${subgroupId}-${blocks.length}`,
          groupId,
          subgroupId,
          items: [item],
        })
      }
    })

    return blocks
  }

  const handleDragEnd = (event: DragEndEvent) => {
    const { active, over } = event

    if (!over || active.id === over.id) return

    const activeId = active.id.toString()
    const overId = over.id.toString()

    const isDraggingGroup = activeId.startsWith('block-group-')
    const isDraggingSubgroup = activeId.startsWith('block-') && !isDraggingGroup
    const isOverBlock = overId.startsWith('block-')

    if (isDraggingGroup) {
      // Arrastar grupo inteiro (todos os subgrupos daquele grupo)
      const groupId = activeId.replace('block-group-', '')
      const blocks = groupItemsIntoBlocks(selectedScoreItems)

      // Encontrar todos os blocos daquele grupo
      const draggedBlocks = blocks.filter((block) => block.groupId === groupId)
      const draggedItems = draggedBlocks.flatMap((block) => block.items)

      if (draggedItems.length === 0) return

      // Encontrar índice de inserção
      let insertIndex = 0
      if (isOverBlock) {
        const overBlockId = overId.startsWith('block-group-')
          ? blocks.find(b => `block-group-${b.groupId}` === overId)?.id
          : overId
        const overBlock = blocks.find((block) => block.id === overBlockId || `block-group-${block.groupId}` === overId)
        if (overBlock) {
          insertIndex = selectedScoreItems.findIndex((item) => item.id === overBlock.items[0].id)
        }
      } else {
        insertIndex = selectedScoreItems.findIndex((item) => item.id === overId)
      }

      // Remover e reinserir
      const draggedItemIds = new Set(draggedItems.map(item => item.id))
      const withoutDragged = selectedScoreItems.filter(item => !draggedItemIds.has(item.id))

      const adjustedIndex = insertIndex > selectedScoreItems.findIndex(item => item.id === draggedItems[0].id)
        ? insertIndex - draggedItems.length
        : insertIndex

      const newOrder = [
        ...withoutDragged.slice(0, adjustedIndex),
        ...draggedItems,
        ...withoutDragged.slice(adjustedIndex),
      ]

      onSelectionChange(newOrder)
    } else if (isDraggingSubgroup) {
      // Arrastar subgrupo (um bloco específico)
      const blocks = groupItemsIntoBlocks(selectedScoreItems)
      const draggedBlock = blocks.find((block) => block.id === activeId)

      if (!draggedBlock) return

      let insertIndex = 0
      if (isOverBlock) {
        const overBlockId = overId.startsWith('block-group-')
          ? blocks.find(b => `block-group-${b.groupId}` === overId)?.id
          : overId
        const overBlock = blocks.find((block) => block.id === overBlockId || `block-group-${block.groupId}` === overId)
        if (overBlock) {
          insertIndex = selectedScoreItems.findIndex((item) => item.id === overBlock.items[0].id)
        }
      } else {
        insertIndex = selectedScoreItems.findIndex((item) => item.id === overId)
      }

      const draggedItemIds = new Set(draggedBlock.items.map(item => item.id))
      const withoutDragged = selectedScoreItems.filter(item => !draggedItemIds.has(item.id))

      const adjustedIndex = insertIndex > selectedScoreItems.findIndex(item => item.id === draggedBlock.items[0].id)
        ? insertIndex - draggedBlock.items.length
        : insertIndex

      const newOrder = [
        ...withoutDragged.slice(0, adjustedIndex),
        ...draggedBlock.items,
        ...withoutDragged.slice(adjustedIndex),
      ]

      onSelectionChange(newOrder)
    } else {
      // Arrastar item individual
      const draggedItem = selectedScoreItems.find((item) => item.id === activeId)
      if (!draggedItem) return

      // Verificar se o item arrastado é um pai (tem filhos)
      const children = selectedScoreItems.filter((item) => item.parentItemId === draggedItem.id)
      const itemsToMove = [draggedItem, ...children]

      // Encontrar índice de inserção
      let insertIndex = selectedScoreItems.findIndex((item) => item.id === overId)

      if (isOverBlock) {
        const blocks = groupItemsIntoBlocks(selectedScoreItems)
        const overBlockId = overId.startsWith('block-group-')
          ? blocks.find(b => `block-group-${b.groupId}` === overId)?.id
          : overId
        const targetBlock = blocks.find((block) => block.id === overBlockId || `block-group-${block.groupId}` === overId)
        if (targetBlock) {
          insertIndex = selectedScoreItems.findIndex((item) => item.id === targetBlock.items[0].id)
        }
      }

      if (insertIndex !== -1) {
        // Remover item pai e filhos da lista
        const itemIdsToMove = new Set(itemsToMove.map(item => item.id))
        const withoutMoved = selectedScoreItems.filter(item => !itemIdsToMove.has(item.id))

        // Ajustar índice se necessário
        const oldIndex = selectedScoreItems.findIndex(item => item.id === draggedItem.id)
        const adjustedIndex = insertIndex > oldIndex
          ? insertIndex - itemsToMove.length
          : insertIndex

        // Inserir item pai e filhos na nova posição
        const newOrder = [
          ...withoutMoved.slice(0, adjustedIndex),
          ...itemsToMove,
          ...withoutMoved.slice(adjustedIndex),
        ]

        onSelectionChange(newOrder)
      }
    }
  }

  const totalAvailable = filteredAndGroupedAvailable.reduce(
    (acc, group) =>
      acc + group.subgroups.reduce((subAcc, subgroup) => subAcc + subgroup.items.length, 0),
    0
  )

  return (
    <div className="grid grid-cols-[1fr_auto_1fr] gap-4">
      {/* Coluna esquerda: Itens disponíveis organizados por grupo/subgrupo */}
      <div className="space-y-2 min-w-0">
        <div className="flex items-center justify-between">
          <h3 className="text-sm font-medium">Score Items Disponíveis</h3>
          <span className="text-xs text-muted-foreground">{totalAvailable} items</span>
        </div>
        <Input
          placeholder="Buscar..."
          value={searchAvailable}
          onChange={(e) => setSearchAvailable(e.target.value)}
          className="h-9"
        />
        <ScrollArea className="h-[500px] border rounded-lg">
          <div className="p-2 space-y-3">
            {totalAvailable === 0 ? (
              <div className="text-center text-sm text-muted-foreground py-8">
                Nenhum item disponível
              </div>
            ) : (
              filteredAndGroupedAvailable.map((group, groupIndex) => (
                <div key={group.groupId} className="space-y-2">
                  {/* Cabeçalho do Grupo */}
                  <div
                    className="bg-primary text-primary-foreground rounded-lg px-3 py-2 cursor-pointer hover:bg-primary/90 transition-colors"
                    onClick={() => addGroup(group)}
                    onKeyDown={(e) => {
                      if (e.key === 'Enter') {
                        addGroup(group)
                      }
                    }}
                    tabIndex={0}
                  >
                    <div className="flex items-center gap-2 min-w-0">
                      <div className="flex items-center justify-center w-6 h-6 rounded-full bg-primary-foreground/20 text-xs font-bold flex-shrink-0">
                        {groupIndex + 1}
                      </div>
                      <h4 className="text-sm font-bold break-words min-w-0 flex-1">{group.groupName}</h4>
                    </div>
                  </div>

                  {/* Subgrupos */}
                  {group.subgroups.map((subgroup, subgroupIndex) => (
                    <div
                      key={subgroup.subgroupId}
                      className="bg-gradient-to-br from-muted to-muted/50 rounded-lg p-2 border border-primary/20"
                    >
                      {/* Cabeçalho do Subgrupo */}
                      <div
                        className="mb-2 pb-1 border-b border-primary/30 cursor-pointer hover:bg-primary/5 transition-colors rounded px-2 py-1"
                        onClick={() => addSubgroup(subgroup.items)}
                        onKeyDown={(e) => {
                          if (e.key === 'Enter') {
                            addSubgroup(subgroup.items)
                          }
                        }}
                        tabIndex={0}
                      >
                        <div className="flex items-center gap-2 min-w-0">
                          <div className="flex items-center justify-center w-5 h-5 rounded-full bg-primary/10 text-xs font-bold text-primary flex-shrink-0">
                            {groupIndex + 1}.{subgroupIndex + 1}
                          </div>
                          <h5 className="text-xs font-bold break-words min-w-0 flex-1">{subgroup.subgroupName}</h5>
                        </div>
                      </div>

                      {/* Items do Subgrupo (selecionáveis) */}
                      <div className="space-y-1">
                        {subgroup.items.map((item) => {
                          const isChild = !!item.parentItemId

                          return (
                            <div
                              key={item.id}
                              className="p-2 border rounded-lg bg-card hover:bg-accent cursor-pointer transition-colors"
                              style={{
                                marginLeft: isChild ? '20px' : undefined,
                                borderLeft: isChild ? '4px solid hsl(var(--primary))' : undefined,
                                paddingLeft: isChild ? '12px' : undefined,
                              }}
                              onClick={() => addItem(item)}
                              onKeyDown={(e) => {
                                if (e.key === 'Enter') {
                                  addItem(item)
                                }
                              }}
                              tabIndex={0}
                            >
                              <div className="flex items-start justify-between gap-2 min-w-0">
                                <div className="flex-1 min-w-0">
                                  <div className="text-sm font-medium break-words">{item.name}</div>
                                </div>
                              </div>
                            </div>
                          )
                        })}
                      </div>
                    </div>
                  ))}
                </div>
              ))
            )}
          </div>
        </ScrollArea>
      </div>

      {/* Coluna central: Botões */}
      <div className="flex flex-col items-center justify-center gap-2">
        <Button variant="outline" size="icon" disabled>
          <ChevronRight className="h-4 w-4" />
        </Button>
        <Button variant="outline" size="icon" disabled>
          <ChevronLeft className="h-4 w-4" />
        </Button>
      </div>

      {/* Coluna direita: Itens selecionados (com drag-drop) */}
      <div className="space-y-2 min-w-0">
        <div className="flex items-center justify-between">
          <h3 className="text-sm font-medium">Score Items Selecionados</h3>
          <span className="text-xs text-muted-foreground">{filteredSelected.length} items</span>
        </div>
        <Input
          placeholder="Buscar..."
          value={searchSelected}
          onChange={(e) => setSearchSelected(e.target.value)}
          className="h-9"
        />
        <ScrollArea className="h-[500px] border rounded-lg">
          <div className="p-2 space-y-4">
            {filteredSelected.length === 0 ? (
              <div className="text-center text-sm text-muted-foreground py-8">
                Nenhum item selecionado
              </div>
            ) : (
              <DndContext
                sensors={sensors}
                collisionDetection={closestCenter}
                onDragEnd={handleDragEnd}
              >
                <SortableContext
                  items={[
                    ...groupItemsIntoBlocks(filteredSelected).map((block) => block.id),
                    ...groupItemsIntoBlocks(filteredSelected).map((block) => `block-group-${block.groupId}`),
                    ...filteredSelected.map((item) => item.id),
                  ]}
                  strategy={verticalListSortingStrategy}
                >
                  {groupItemsIntoBlocks(filteredSelected).map((block, index, blocks) => {
                    // Verificar se deve mostrar cabeçalho de grupo (primeiro bloco do grupo)
                    const prevBlock = index > 0 ? blocks[index - 1] : null
                    const showGroupHeader = !prevBlock || prevBlock.groupId !== block.groupId

                    return (
                      <SortableBlock
                        key={block.id}
                        block={block}
                        onRemoveItem={(item) => removeItem(item)}
                        showGroupHeader={showGroupHeader}
                      />
                    )
                  })}
                </SortableContext>
              </DndContext>
            )}
          </div>
        </ScrollArea>
        <div className="text-xs text-muted-foreground">
          💡 Arraste para reordenar • Clique para remover
        </div>
      </div>
    </div>
  )
}
