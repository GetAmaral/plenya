'use client'

import { useState, useEffect } from 'react'
import { Edit, Trash2, Plus, GripVertical, ChevronRight, ChevronLeft } from 'lucide-react'
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from '@/components/ui/accordion'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { ScoreGroup, ScoreSubgroup, ScoreItem, useDeleteScoreGroup, useDeleteScoreSubgroup, useUpdateScoreItem } from '@/lib/api/score-api'
import { ScoreGroupDialog } from './ScoreGroupDialog'
import { ScoreSubgroupDialog } from './ScoreSubgroupDialog'
import { ScoreItemDialog } from './ScoreItemDialog'
import { ScoreItemCard } from './ScoreItemCard'
import { DeleteConfirmDialog } from './DeleteConfirmDialog'
import { toast } from 'sonner'
import {
  DndContext,
  DragEndEvent,
  DragOverlay,
  DragStartEvent,
  closestCorners,
  PointerSensor,
  useSensor,
  useSensors,
} from '@dnd-kit/core'
import {
  SortableContext,
  verticalListSortingStrategy,
  useSortable,
} from '@dnd-kit/sortable'
import { CSS } from '@dnd-kit/utilities'

interface ScoreTreeViewProps {
  groups: ScoreGroup[]
  expandedNodes?: Record<string, boolean>
  expandClinicalTexts?: boolean
  openSubgroups?: string[] // IDs de subgrupos que devem estar abertos
}

// Helper para organizar itens em hierarquia
interface ItemWithChildren extends ScoreItem {
  children?: ItemWithChildren[]
}

// Componente SortableItem para drag and drop
interface SortableItemProps {
  item: ItemWithChildren
  depth: number
  expandedNodes: Record<string, boolean>
  expandClinicalTexts: boolean
  allItems: ScoreItem[]
  onIndent?: (itemId: string) => void
  onOutdent?: (itemId: string) => void
  canIndent: boolean
  canOutdent: boolean
}

function SortableItem({
  item,
  depth,
  expandedNodes,
  expandClinicalTexts,
  allItems,
  onIndent,
  onOutdent,
  canIndent,
  canOutdent
}: SortableItemProps) {
  const {
    attributes,
    listeners,
    setNodeRef,
    transform,
    transition,
    isDragging,
  } = useSortable({ id: item.id })

  const style = {
    transform: CSS.Transform.toString(transform),
    transition,
    opacity: isDragging ? 0.5 : 1,
    marginLeft: `${depth * 20}px`,
    paddingLeft: depth > 0 ? '8px' : '0',
    borderLeft: depth > 0 ? '2px solid hsl(var(--border))' : 'none',
  }

  return (
    <div
      ref={setNodeRef}
      style={style}
      id={`item-${item.id}`}
      className="rounded-lg transition-all flex items-start gap-1 group"
    >
      {/* Coluna de controles à esquerda */}
      <div className="flex-shrink-0 flex flex-col gap-0.5 pt-2.5">
        {/* Drag Handle */}
        <div
          {...attributes}
          {...listeners}
          className="opacity-0 group-hover:opacity-100 transition-opacity cursor-grab active:cursor-grabbing p-0.5"
          title="Arrastar"
        >
          <GripVertical className="h-3.5 w-3.5 text-muted-foreground" />
        </div>

        {/* Indent Button */}
        {canIndent && onIndent && (
          <button
            onClick={() => onIndent(item.id)}
            className="opacity-0 group-hover:opacity-100 transition-opacity p-0.5 hover:bg-accent rounded"
            title="Indentar (tornar filho do item acima)"
          >
            <ChevronRight className="h-3.5 w-3.5 text-muted-foreground" />
          </button>
        )}

        {/* Outdent Button */}
        {canOutdent && onOutdent && (
          <button
            onClick={() => onOutdent(item.id)}
            className="opacity-0 group-hover:opacity-100 transition-opacity p-0.5 hover:bg-accent rounded"
            title="Desindentar (remover hierarquia)"
          >
            <ChevronLeft className="h-3.5 w-3.5 text-muted-foreground" />
          </button>
        )}
      </div>

      {/* Card ocupa o resto do espaço */}
      <div className="flex-1 min-w-0">
        <ScoreItemCard
          item={item}
          isExpanded={expandedNodes[`item-${item.id}`]}
          expandClinicalTexts={expandClinicalTexts}
        />
      </div>
    </div>
  )
}

function organizeItemsHierarchy(items: ScoreItem[]): ItemWithChildren[] {
  const itemMap = new Map<string, ItemWithChildren>()
  const rootItems: ItemWithChildren[] = []

  // Primeiro, cria um map de todos os itens
  items.forEach(item => {
    itemMap.set(item.id, { ...item, children: [] })
  })

  // Depois, organiza a hierarquia
  items.forEach(item => {
    const itemWithChildren = itemMap.get(item.id)!

    if (item.parentItemId) {
      // Se tem parent, adiciona como filho
      const parent = itemMap.get(item.parentItemId)
      if (parent) {
        parent.children = parent.children || []
        parent.children.push(itemWithChildren)
      } else {
        // Parent não encontrado, adiciona como raiz
        rootItems.push(itemWithChildren)
      }
    } else {
      // Sem parent, é item raiz
      rootItems.push(itemWithChildren)
    }
  })

  // Ordena todos os níveis por 'order'
  const sortItemsRecursive = (items: ItemWithChildren[]) => {
    items.sort((a, b) => a.order - b.order)
    items.forEach(item => {
      if (item.children && item.children.length > 0) {
        sortItemsRecursive(item.children)
      }
    })
  }

  sortItemsRecursive(rootItems)

  return rootItems
}

export function ScoreTreeView({ groups, expandedNodes = {}, expandClinicalTexts = false, openSubgroups = [] }: ScoreTreeViewProps) {
  const [editingGroup, setEditingGroup] = useState<ScoreGroup | null>(null)
  const [creatingSubgroupFor, setCreatingSubgroupFor] = useState<string | null>(null)
  const [editingSubgroup, setEditingSubgroup] = useState<ScoreSubgroup | null>(null)
  const [creatingItemFor, setCreatingItemFor] = useState<string | null>(null)
  const [deletingGroup, setDeletingGroup] = useState<ScoreGroup | null>(null)
  const [deletingSubgroup, setDeletingSubgroup] = useState<ScoreSubgroup | null>(null)
  const [accordionValues, setAccordionValues] = useState<Record<string, string[]>>({})
  const [activeId, setActiveId] = useState<string | null>(null)

  const deleteGroup = useDeleteScoreGroup()
  const deleteSubgroup = useDeleteScoreSubgroup()
  const updateItem = useUpdateScoreItem()

  // Configurar sensores para drag and drop
  const sensors = useSensors(
    useSensor(PointerSensor, {
      activationConstraint: {
        distance: 8, // 8px de movimento antes de iniciar drag
      },
    })
  )

  // Sincroniza accordionValues quando expandedNodes ou openSubgroups muda (ex: busca, expandir tudo)
  useEffect(() => {
    // Se não há expandedNodes externos nem openSubgroups, não fazer nada (usar estado interno)
    if (Object.keys(expandedNodes).length === 0 && openSubgroups.length === 0) return

    setAccordionValues(prev => {
      const newAccordionValues: Record<string, string[]> = { ...prev }

      groups.forEach(group => {
        const groupKey = group.id
        const expandedSubgroupsFromProps: string[] = []

        group.subgroups?.forEach(subgroup => {
          // Abrir se está em expandedNodes OU em openSubgroups
          if (expandedNodes[`subgroup-${subgroup.id}`] || openSubgroups.includes(subgroup.id)) {
            expandedSubgroupsFromProps.push(subgroup.id)
          }
        })

        // Mescla expandidos de props com os já expandidos manualmente
        const currentExpanded = prev[groupKey] || []
        newAccordionValues[groupKey] = Array.from(
          new Set([...currentExpanded, ...expandedSubgroupsFromProps])
        )
      })

      return newAccordionValues
    })
  }, [expandedNodes, openSubgroups, groups])

  const handleDeleteGroup = async (group: ScoreGroup) => {
    try {
      await deleteGroup.mutateAsync(group.id)
      toast.success('Grupo excluído com sucesso')
      setDeletingGroup(null)
    } catch (error) {
      toast.error('Erro ao excluir grupo')
    }
  }

  const handleDeleteSubgroup = async (subgroup: ScoreSubgroup) => {
    try {
      await deleteSubgroup.mutateAsync(subgroup.id)
      toast.success('Subgrupo excluído com sucesso')
      setDeletingSubgroup(null)
    } catch (error) {
      toast.error('Erro ao excluir subgrupo')
    }
  }

  const handleDragStart = (event: DragStartEvent) => {
    setActiveId(event.active.id as string)
  }

  const handleIndent = async (itemId: string, subgroupId: string, allItems: ScoreItem[]) => {
    const currentItem = allItems.find(i => i.id === itemId)
    if (!currentItem) return

    // Encontrar o item imediatamente acima (mesmo parent level)
    const sameParentItems = allItems
      .filter(i => i.parentItemId === currentItem.parentItemId)
      .sort((a, b) => a.order - b.order)

    const currentIndex = sameParentItems.findIndex(i => i.id === itemId)
    if (currentIndex <= 0) return // É o primeiro, não pode indentar

    const itemAbove = sameParentItems[currentIndex - 1]

    // Encontrar filhos do item atual
    const children = allItems.filter(i => i.parentItemId === itemId)

    try {
      // Atualizar o item para ser filho do item acima
      const updatePromises = [
        updateItem.mutateAsync({
          id: itemId,
          data: {
            parentItemId: itemAbove.id,
          },
        })
      ]

      // Se tem filhos, eles também viram filhos do item acima (não do item atual)
      if (children.length > 0) {
        const childrenPromises = children.map(child =>
          updateItem.mutateAsync({
            id: child.id,
            data: {
              parentItemId: itemAbove.id,
            },
          })
        )
        updatePromises.push(...childrenPromises)
      }

      await Promise.all(updatePromises)

      if (children.length > 0) {
        toast.success(`Item e ${children.length} filho(s) indentados`)
      } else {
        toast.success('Item indentado')
      }
    } catch (error) {
      toast.error('Erro ao indentar item')
    }
  }

  const handleOutdent = async (itemId: string, subgroupId: string, allItems: ScoreItem[]) => {
    const currentItem = allItems.find(i => i.id === itemId)
    if (!currentItem || !currentItem.parentItemId) return

    const parentItem = allItems.find(i => i.id === currentItem.parentItemId)
    if (!parentItem) return

    try {
      // Encontrar irmãos abaixo dele (outros filhos do mesmo pai, excluindo o próprio item)
      const siblings = allItems
        .filter(i =>
          i.id !== itemId && // Não incluir o próprio item
          i.parentItemId === currentItem.parentItemId &&
          i.order >= currentItem.order // >= para pegar itens com mesmo order também
        )
        .sort((a, b) => a.order - b.order)

      // Atualizar o item para ser filho do avô (ou raiz = null)
      const newParentId = parentItem.parentItemId || null

      await updateItem.mutateAsync({
        id: itemId,
        data: { parentItemId: newParentId },
      })

      // Se tem irmãos abaixo, torná-los filhos do item desindentado
      if (siblings.length > 0) {
        await Promise.all(
          siblings.map(sibling =>
            updateItem.mutateAsync({
              id: sibling.id,
              data: {
                parentItemId: itemId,
              },
            })
          )
        )
        toast.success(`Desindentado e capturou ${siblings.length} filho(s)`)
      } else {
        toast.success('Item desindentado')
      }
    } catch (error) {
      toast.error('Erro ao desindentar item')
    }
  }

  // Função helper para coletar todos os descendentes de um item
  const collectDescendants = (itemId: string, allItems: ScoreItem[]): ScoreItem[] => {
    const descendants: ScoreItem[] = []
    const children = allItems.filter(item => item.parentItemId === itemId)

    for (const child of children) {
      descendants.push(child)
      // Recursivamente buscar descendentes do filho
      descendants.push(...collectDescendants(child.id, allItems))
    }

    return descendants
  }

  const handleDragEnd = async (event: DragEndEvent) => {
    const { active, over } = event
    setActiveId(null)

    if (!over || active.id === over.id) return

    // Encontrar o item sendo arrastado e o item sobre o qual foi solto
    let draggedItem: ScoreItem | null = null
    let targetItem: ScoreItem | null = null
    let subgroupId: string | null = null
    let allSubgroupItems: ScoreItem[] = []

    // Buscar nos subgrupos
    for (const group of groups) {
      for (const subgroup of group.subgroups || []) {
        const items = subgroup.items || []

        const dragged = items.find(item => item.id === active.id)
        if (dragged) {
          draggedItem = dragged
          subgroupId = subgroup.id
          allSubgroupItems = items
        }

        const target = items.find(item => item.id === over.id)
        if (target) {
          targetItem = target
        }

        if (draggedItem && targetItem) break
      }
      if (draggedItem && targetItem) break
    }

    if (!draggedItem || !targetItem || !subgroupId) return

    // Normalizar parentItemId (tratar null e undefined como equivalentes)
    const normalizeParentId = (id: string | null | undefined) => id || null
    const draggedParentId = normalizeParentId(draggedItem.parentItemId)
    const targetParentId = normalizeParentId(targetItem.parentItemId)

    // Verificar se estão no mesmo nível hierárquico (mesmo parentItemId)
    if (draggedParentId !== targetParentId) {
      toast.error('Não é possível reordenar itens de níveis hierárquicos diferentes')
      return
    }

    // Obter todos os irmãos (items com o mesmo parentItemId) ordenados
    const siblings = allSubgroupItems
      .filter(item => normalizeParentId(item.parentItemId) === draggedParentId)
      .sort((a, b) => a.order - b.order)

    // Se não houver mudança real de posição, cancelar
    const currentIndex = siblings.findIndex(item => item.id === draggedItem.id)
    const targetIndex = siblings.findIndex(item => item.id === targetItem.id)

    if (currentIndex === targetIndex) return

    // Remover o item arrastado da lista
    const reorderedSiblings = siblings.filter(item => item.id !== draggedItem.id)

    // Inserir na nova posição
    reorderedSiblings.splice(targetIndex, 0, draggedItem)

    // RECALCULAR TODOS OS ORDERS DO SUBGRUPO INTEIRO (não só siblings)
    // Isso garante que não haverá orders duplicados ou gaps
    try {
      // 1. Reassignar orders para os siblings reordenados
      const siblingUpdates = reorderedSiblings.map((item, index) => {
        const newOrder = index + 1
        return updateItem.mutateAsync({
          id: item.id,
          data: { order: newOrder },
        })
      })

      await Promise.all(siblingUpdates)
      toast.success('Itens reordenados')
    } catch (error) {
      toast.error('Erro ao reordenar itens')
    }
  }

  // Função recursiva para renderizar itens com indentação
  const renderItemWithChildren = (
    item: ItemWithChildren,
    depth: number = 0,
    allItems: ScoreItem[] = [],
    subgroupId: string = ''
  ): JSX.Element[] => {
    const elements: JSX.Element[] = []

    // Determinar se pode indentar ou desindentar
    // Tratar null e undefined como equivalentes para parentItemId
    const normalizeParentId = (id: string | null | undefined) => id || null
    const sameParentItems = allItems
      .filter(i => normalizeParentId(i.parentItemId) === normalizeParentId(item.parentItemId))
      .sort((a, b) => a.order - b.order)
    const currentIndex = sameParentItems.findIndex(i => i.id === item.id)

    const canIndent = !item.parentItemId && currentIndex > 0 // Não tem parent e não é o primeiro
    const canOutdent = !!item.parentItemId // Tem parent

    // Renderiza o item atual com indentação (agora com drag and drop)
    // Key inclui parentItemId para forçar re-render quando hierarquia muda
    elements.push(
      <SortableItem
        key={`${item.id}-${item.parentItemId || 'root'}`}
        item={item}
        depth={depth}
        expandedNodes={expandedNodes}
        expandClinicalTexts={expandClinicalTexts}
        allItems={allItems}
        onIndent={(id) => handleIndent(id, subgroupId, allItems)}
        onOutdent={(id) => handleOutdent(id, subgroupId, allItems)}
        canIndent={canIndent}
        canOutdent={canOutdent}
      />
    )

    // Renderiza os filhos recursivamente com indentação aumentada
    if (item.children && item.children.length > 0) {
      item.children.forEach(child => {
        elements.push(...renderItemWithChildren(child, depth + 1, allItems, subgroupId))
      })
    }

    return elements
  }

  return (
    <DndContext
      sensors={sensors}
      collisionDetection={closestCorners}
      onDragStart={handleDragStart}
      onDragEnd={handleDragEnd}
    >
      <div className="space-y-2">
        {groups.map((group) => (
          <div key={group.id} id={`group-${group.id}`} className="rounded-lg border transition-all">
            {/* Group Header */}
            <div className="flex items-center justify-between px-3 py-2 bg-muted/50">
              <div className="flex-1">
                <h2 className="text-sm font-semibold text-left">{group.name}</h2>
                <p className="text-xs text-muted-foreground mt-0.5 text-left">
                  {group.subgroups?.length || 0} subgrupos
                </p>
              </div>

              <div className="flex gap-1.5">
                <Button
                  variant="outline"
                  size="sm"
                  onClick={() => setCreatingSubgroupFor(group.id)}
                >
                  <Plus className="h-4 w-4 mr-1" />
                  Subgrupo
                </Button>
                <Button
                  variant="outline"
                  size="sm"
                  onClick={() => setEditingGroup(group)}
                >
                  <Edit className="h-4 w-4" />
                </Button>
                <Button
                  variant="outline"
                  size="sm"
                  onClick={() => setDeletingGroup(group)}
                >
                  <Trash2 className="h-4 w-4" />
                </Button>
              </div>
            </div>

            {/* Subgroups */}
            {group.subgroups && group.subgroups.length > 0 && (
              <div className="p-2">
                <Accordion
                  type="multiple"
                  className="space-y-1.5"
                  value={accordionValues[group.id] || []}
                  onValueChange={(newValue) => {
                    setAccordionValues(prev => ({
                      ...prev,
                      [group.id]: newValue
                    }))
                  }}
                >
                  {group.subgroups.map((subgroup) => (
                    <AccordionItem
                      key={subgroup.id}
                      value={subgroup.id}
                      id={`subgroup-${subgroup.id}`}
                      className="border rounded-md transition-all"
                    >
                      <div className="flex items-center justify-between px-2.5">
                        <AccordionTrigger className="flex-1 hover:no-underline py-2 text-left">
                          <div className="flex items-center gap-2">
                            <span className="text-sm font-medium text-left">{subgroup.name}</span>
                            <Badge variant="outline" className="text-[10px] px-1.5 py-0">
                              {subgroup.items?.length || 0}
                            </Badge>
                          </div>
                        </AccordionTrigger>
                        <div className="flex gap-0.5 ml-2 flex-shrink-0">
                          <Button
                            variant="ghost"
                            size="sm"
                            onClick={() => setCreatingItemFor(subgroup.id)}
                            className="h-6 text-[10px] px-2"
                          >
                            <Plus className="h-3 w-3 mr-0.5" />
                            Item
                          </Button>
                          <Button
                            variant="ghost"
                            size="sm"
                            onClick={() => setEditingSubgroup(subgroup)}
                            className="h-6 w-6 p-0"
                          >
                            <Edit className="h-3 w-3" />
                          </Button>
                          <Button
                            variant="ghost"
                            size="sm"
                            onClick={() => setDeletingSubgroup(subgroup)}
                            className="h-6 w-6 p-0"
                          >
                            <Trash2 className="h-3 w-3" />
                          </Button>
                        </div>
                      </div>

                      <AccordionContent className="px-2.5 pb-2">
                        {subgroup.items && subgroup.items.length > 0 ? (
                          <SortableContext
                            items={subgroup.items.map(item => item.id)}
                            strategy={verticalListSortingStrategy}
                          >
                            <div className="space-y-1.5 mt-1">
                              {organizeItemsHierarchy(subgroup.items).map((item) =>
                                renderItemWithChildren(item, 0, subgroup.items || [], subgroup.id)
                              )}
                            </div>
                          </SortableContext>
                        ) : (
                          <div className="text-xs text-muted-foreground text-left py-2">
                            Nenhum item cadastrado
                          </div>
                        )}
                      </AccordionContent>
                    </AccordionItem>
                  ))}
                </Accordion>
              </div>
            )}

            {(!group.subgroups || group.subgroups.length === 0) && (
              <div className="p-2 text-left text-xs text-muted-foreground">
                Nenhum subgrupo cadastrado
              </div>
            )}
          </div>
        ))}
      </div>

      {/* Edit Group Dialog */}
      {editingGroup && (
        <ScoreGroupDialog
          open={!!editingGroup}
          onOpenChange={(open) => !open && setEditingGroup(null)}
          group={editingGroup}
        />
      )}

      {/* Create Subgroup Dialog */}
      {creatingSubgroupFor && (
        <ScoreSubgroupDialog
          open={!!creatingSubgroupFor}
          onOpenChange={(open) => !open && setCreatingSubgroupFor(null)}
          groupId={creatingSubgroupFor}
        />
      )}

      {/* Edit Subgroup Dialog */}
      {editingSubgroup && (
        <ScoreSubgroupDialog
          open={!!editingSubgroup}
          onOpenChange={(open) => !open && setEditingSubgroup(null)}
          groupId={editingSubgroup.groupId}
          subgroup={editingSubgroup}
        />
      )}

      {/* Create Item Dialog */}
      {creatingItemFor && (
        <ScoreItemDialog
          open={!!creatingItemFor}
          onOpenChange={(open) => !open && setCreatingItemFor(null)}
          subgroupId={creatingItemFor}
        />
      )}

      {/* Delete Group Confirm Dialog */}
      {deletingGroup && (
        <DeleteConfirmDialog
          open={!!deletingGroup}
          onOpenChange={(open) => !open && setDeletingGroup(null)}
          onConfirm={() => handleDeleteGroup(deletingGroup)}
          title="Excluir grupo de escores?"
          description={`Tem certeza que deseja excluir "${deletingGroup.name}"? Isso também excluirá todos os subgrupos, itens e níveis associados.`}
          isLoading={deleteGroup.isPending}
        />
      )}

      {/* Delete Subgroup Confirm Dialog */}
      {deletingSubgroup && (
        <DeleteConfirmDialog
          open={!!deletingSubgroup}
          onOpenChange={(open) => !open && setDeletingSubgroup(null)}
          onConfirm={() => handleDeleteSubgroup(deletingSubgroup)}
          title="Excluir subgrupo?"
          description={`Tem certeza que deseja excluir "${deletingSubgroup.name}"? Isso também excluirá todos os itens e níveis associados.`}
          isLoading={deleteSubgroup.isPending}
        />
      )}

      {/* Drag Overlay para mostrar o item e seus filhos sendo arrastados */}
      <DragOverlay>
        {activeId ? (() => {
          // Encontrar o item ativo e seus descendentes
          let activeItem: ScoreItem | null = null
          let allItems: ScoreItem[] = []

          for (const group of groups) {
            for (const subgroup of group.subgroups || []) {
              const items = subgroup.items || []
              const found = items.find(item => item.id === activeId)
              if (found) {
                activeItem = found
                allItems = items
                break
              }
            }
            if (activeItem) break
          }

          if (!activeItem) return null

          const descendants = collectDescendants(activeItem.id, allItems)
          const itemsToShow = [activeItem, ...descendants]

          return (
            <div className="bg-background border-2 border-primary rounded-lg shadow-2xl p-2 opacity-90">
              <div className="text-xs font-semibold mb-1">
                Movendo {itemsToShow.length > 1 ? `${itemsToShow.length} itens` : '1 item'}
              </div>
              <div className="space-y-1">
                {itemsToShow.slice(0, 3).map((item, idx) => (
                  <div key={item.id} className="text-xs text-muted-foreground truncate" style={{ paddingLeft: `${idx * 12}px` }}>
                    • {item.name}
                  </div>
                ))}
                {itemsToShow.length > 3 && (
                  <div className="text-xs text-muted-foreground">
                    + {itemsToShow.length - 3} mais
                  </div>
                )}
              </div>
            </div>
          )
        })() : null}
      </DragOverlay>
    </DndContext>
  )
}
