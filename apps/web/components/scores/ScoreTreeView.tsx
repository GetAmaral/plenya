'use client'

import { useState, useEffect } from 'react'
import { Edit, Trash2, Plus } from 'lucide-react'
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from '@/components/ui/accordion'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { ScoreGroup, ScoreSubgroup, useDeleteScoreGroup, useDeleteScoreSubgroup } from '@/lib/api/score-api'
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
}

export function ScoreTreeView({ groups, expandedNodes = {}, expandClinicalTexts = false }: ScoreTreeViewProps) {
  const [editingGroup, setEditingGroup] = useState<ScoreGroup | null>(null)
  const [creatingSubgroupFor, setCreatingSubgroupFor] = useState<string | null>(null)
  const [editingSubgroup, setEditingSubgroup] = useState<ScoreSubgroup | null>(null)
  const [creatingItemFor, setCreatingItemFor] = useState<string | null>(null)
  const [deletingGroup, setDeletingGroup] = useState<ScoreGroup | null>(null)
  const [deletingSubgroup, setDeletingSubgroup] = useState<ScoreSubgroup | null>(null)
  const [accordionValues, setAccordionValues] = useState<Record<string, string[]>>({})

  const deleteGroup = useDeleteScoreGroup()
  const deleteSubgroup = useDeleteScoreSubgroup()

  // Sincronizar accordionValues com expandedNodes vindos de fora
  useEffect(() => {
    const newAccordionValues: Record<string, string[]> = {}

    groups.forEach(group => {
      const groupKey = group.id
      const expandedSubgroups: string[] = []

      group.subgroups?.forEach(subgroup => {
        if (expandedNodes[`subgroup-${subgroup.id}`]) {
          expandedSubgroups.push(subgroup.id)
        }
      })

      newAccordionValues[groupKey] = expandedSubgroups
    })

    setAccordionValues(newAccordionValues)
  }, [expandedNodes, groups])

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

  return (
    <>
      <div className="space-y-3">
        {groups.map((group) => (
          <div key={group.id} id={`group-${group.id}`} className="rounded-lg border transition-all">
            {/* Group Header */}
            <div className="flex items-center justify-between p-3 bg-muted/50">
              <div className="flex-1">
                <div className="flex items-center gap-2">
                  <h2 className="text-base font-semibold">{group.name}</h2>
                  <Badge variant="secondary" className="text-xs">
                    Ordem: {group.order}
                  </Badge>
                </div>
                <p className="text-sm text-muted-foreground mt-0.5">
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
                      <div className="flex items-center px-3">
                        <AccordionTrigger className="flex-1 hover:no-underline py-3">
                          <div className="flex items-center gap-2">
                            <span className="font-medium">{subgroup.name}</span>
                            <Badge variant="outline" className="text-xs">
                              {subgroup.items?.length || 0} itens
                            </Badge>
                          </div>
                        </AccordionTrigger>
                        <div className="flex gap-1 ml-4">
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
                            {subgroup.items.map((item) => (
                              <div key={item.id} id={`item-${item.id}`} className="rounded-lg transition-all">
                                <ScoreItemCard
                                  item={item}
                                  isExpanded={expandedNodes[`item-${item.id}`]}
                                  expandClinicalTexts={expandClinicalTexts}
                                />
                              </div>
                            ))}
                          </div>
                        ) : (
                          <div className="text-sm text-muted-foreground text-center py-3">
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
              <div className="p-3 text-center text-sm text-muted-foreground">
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
