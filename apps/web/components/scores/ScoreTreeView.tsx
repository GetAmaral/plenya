'use client'

import { useState, useEffect } from 'react'
import { Edit, Trash2, Plus } from 'lucide-react'
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from '@/components/ui/accordion'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { ScoreGroup, ScoreSubgroup, ScoreItem, useDeleteScoreGroup, useDeleteScoreSubgroup } from '@/lib/api/score-api'
import { ScoreGroupDialog } from './ScoreGroupDialog'
import { ScoreSubgroupDialog } from './ScoreSubgroupDialog'
import { ScoreItemDialog } from './ScoreItemDialog'
import { ScoreItemCard } from './ScoreItemCard'
import { DeleteConfirmDialog } from './DeleteConfirmDialog'
import { toast } from 'sonner'

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

  const deleteGroup = useDeleteScoreGroup()
  const deleteSubgroup = useDeleteScoreSubgroup()

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

  // Função recursiva para renderizar itens com indentação
  const renderItemWithChildren = (item: ItemWithChildren, depth: number = 0): JSX.Element[] => {
    const elements: JSX.Element[] = []

    // Renderiza o item atual com indentação
    elements.push(
      <div
        key={item.id}
        id={`item-${item.id}`}
        className="rounded-lg transition-all"
        style={{
          marginLeft: `${depth * 32}px`,
          paddingLeft: depth > 0 ? '12px' : '0',
          borderLeft: depth > 0 ? '2px solid hsl(var(--border))' : 'none'
        }}
      >
        <ScoreItemCard
          item={item}
          isExpanded={expandedNodes[`item-${item.id}`]}
          expandClinicalTexts={expandClinicalTexts}
        />
      </div>
    )

    // Renderiza os filhos recursivamente com indentação aumentada
    if (item.children && item.children.length > 0) {
      item.children.forEach(child => {
        elements.push(...renderItemWithChildren(child, depth + 1))
      })
    }

    return elements
  }

  return (
    <>
      <div className="space-y-3">
        {groups.map((group) => (
          <div key={group.id} id={`group-${group.id}`} className="rounded-lg border transition-all">
            {/* Group Header */}
            <div className="flex items-center justify-between p-3 bg-muted/50">
              <div className="flex-1">
                <h2 className="text-base font-semibold text-left">{group.name}</h2>
                <p className="text-sm text-muted-foreground mt-0.5 text-left">
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
              <div className="p-3">
                <Accordion
                  type="multiple"
                  className="space-y-2"
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
                      <div className="flex items-center justify-between px-3">
                        <AccordionTrigger className="flex-1 hover:no-underline py-3 text-left">
                          <div className="flex items-center gap-2">
                            <span className="font-medium text-left">{subgroup.name}</span>
                            <Badge variant="outline" className="text-xs">
                              {subgroup.items?.length || 0} itens
                            </Badge>
                          </div>
                        </AccordionTrigger>
                        <div className="flex gap-1 ml-4 flex-shrink-0">
                          <Button
                            variant="ghost"
                            size="sm"
                            onClick={() => setCreatingItemFor(subgroup.id)}
                            className="h-7 text-xs"
                          >
                            <Plus className="h-3 w-3 mr-1" />
                            Item
                          </Button>
                          <Button
                            variant="ghost"
                            size="sm"
                            onClick={() => setEditingSubgroup(subgroup)}
                            className="h-7 w-7 p-0"
                          >
                            <Edit className="h-3 w-3" />
                          </Button>
                          <Button
                            variant="ghost"
                            size="sm"
                            onClick={() => setDeletingSubgroup(subgroup)}
                            className="h-7 w-7 p-0"
                          >
                            <Trash2 className="h-3 w-3" />
                          </Button>
                        </div>
                      </div>

                      <AccordionContent className="px-3 pb-3">
                        {subgroup.items && subgroup.items.length > 0 ? (
                          <div className="space-y-2 mt-2">
                            {organizeItemsHierarchy(subgroup.items).map((item) =>
                              renderItemWithChildren(item, 0)
                            )}
                          </div>
                        ) : (
                          <div className="text-sm text-muted-foreground text-left py-3">
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
              <div className="p-3 text-left text-sm text-muted-foreground">
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
    </>
  )
}
