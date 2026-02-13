'use client'

import { useState, memo, useEffect } from 'react'
import { Edit, Trash2, Plus, Info, Calendar, Copy, FileText, MoreVertical } from 'lucide-react'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { Card, CardContent, CardHeader } from '@/components/ui/card'
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from '@/components/ui/accordion'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { ScoreItem, ScoreLevel, useDeleteScoreItem, useDeleteScoreLevel, useCreateScoreItem, useCreateScoreLevel } from '@/lib/api/score-api'
import { ScoreLevelBadge } from './ScoreLevelBadge'
import { ScoreItemDialog } from './ScoreItemDialog'
import { ScoreLevelDialog } from './ScoreLevelDialog'
import { DeleteConfirmDialog } from './DeleteConfirmDialog'
import { ScoreItemArticlesSheet } from './ScoreItemArticlesSheet'
import { toast } from 'sonner'
import { format } from 'date-fns'
import { ptBR } from 'date-fns/locale'

interface ScoreItemCardProps {
  item: ScoreItem
  isExpanded?: boolean
  expandClinicalTexts?: boolean
}

function ScoreItemCardComponent({ item, isExpanded, expandClinicalTexts = false }: ScoreItemCardProps) {
  const [isEditDialogOpen, setIsEditDialogOpen] = useState(false)
  const [isAddLevelDialogOpen, setIsAddLevelDialogOpen] = useState(false)
  const [isAddChildItemDialogOpen, setIsAddChildItemDialogOpen] = useState(false)
  const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false)
  const [isArticlesSheetOpen, setIsArticlesSheetOpen] = useState(false)
  const [editingLevel, setEditingLevel] = useState<ScoreLevel | null>(null)
  const [deletingLevel, setDeletingLevel] = useState<ScoreLevel | null>(null)
  // Initialize with empty string to keep it controlled from the start
  const [accordionValue, setAccordionValue] = useState<string>(
    expandClinicalTexts ? 'clinical-info' : ''
  )

  const deleteItem = useDeleteScoreItem()
  const deleteLevel = useDeleteScoreLevel()
  const createItem = useCreateScoreItem()
  const createLevel = useCreateScoreLevel()

  // Sync accordion state with expandClinicalTexts prop
  useEffect(() => {
    if (expandClinicalTexts) {
      setAccordionValue('clinical-info')
    } else {
      // When collapsing all, close the accordion
      setAccordionValue('')
    }
  }, [expandClinicalTexts])

  const handleDuplicate = async () => {
    try {
      // Criar o item duplicado
      const newItem = await createItem.mutateAsync({
        name: `${item.name} (Cópia)`,
        unit: item.unit,
        unitConversion: item.unitConversion,
        clinicalRelevance: item.clinicalRelevance,
        patientExplanation: item.patientExplanation,
        conduct: item.conduct,
        points: item.points,
        order: item.order + 1,
        subgroupId: item.subgroupId,
        parentItemId: item.parentItemId,
      })

      // Duplicar os níveis se existirem
      if (item.levels && item.levels.length > 0) {
        for (const level of item.levels) {
          await createLevel.mutateAsync({
            itemId: newItem.id,
            level: level.level,
            name: level.name,
            lowerLimit: level.lowerLimit,
            upperLimit: level.upperLimit,
            operator: level.operator,
            clinicalRelevance: level.clinicalRelevance,
            patientExplanation: level.patientExplanation,
            conduct: level.conduct,
          })
        }
        toast.success(`Item e ${item.levels.length} ${item.levels.length === 1 ? 'nível duplicado' : 'níveis duplicados'} com sucesso`)
      } else {
        toast.success('Item duplicado com sucesso')
      }
    } catch (error: any) {
      // Extract error message from API error
      const errorMessage = error?.message || 'Erro ao duplicar item'
      toast.error(errorMessage)
    }
  }

  const handleDeleteItem = async () => {
    try {
      await deleteItem.mutateAsync(item.id)
      toast.success('Item excluído com sucesso')
      setIsDeleteDialogOpen(false)
    } catch (error: any) {
      // Extract error message from API error
      const errorMessage = error?.message || 'Erro ao excluir item'
      toast.error(errorMessage)
    }
  }

  const handleDeleteLevel = async (level: ScoreLevel) => {
    try {
      await deleteLevel.mutateAsync(level.id)
      toast.success('Nível excluído com sucesso')
      setDeletingLevel(null)
    } catch (error: any) {
      // Extract error message from API error
      const errorMessage = error?.message || 'Erro ao excluir nível'
      toast.error(errorMessage)
    }
  }

  // Sort levels by level number
  const sortedLevels = [...(item.levels || [])].sort((a, b) => a.level - b.level)

  // Check if item has any clinical information
  const hasClinicalInfo = item.clinicalRelevance || item.patientExplanation || item.conduct

  return (
    <>
      <Card>
        <CardHeader className="pb-2 pt-2.5 px-3">
          <div className="flex items-start justify-between">
            <div className="flex-1">
              <div className="flex items-center gap-2">
                <h4 className="text-xs font-semibold text-left">{item.name}</h4>
                {item.unit && (
                  <Badge variant="secondary" className="text-[10px] px-1.5 py-0">
                    {item.unit}
                  </Badge>
                )}
              </div>
              <div className="flex items-center gap-3 mt-0.5 text-xs text-muted-foreground text-left">
                {item.points !== undefined && item.points > 0 && (
                  <span>{item.points} pts</span>
                )}
                {item.unitConversion && (
                  <span className="text-[10px]">{item.unitConversion}</span>
                )}
              </div>
            </div>

            <div className="flex gap-1">
              {/* Editar - Ação primária destacada */}
              <Button
                variant="outline"
                size="sm"
                onClick={() => setIsEditDialogOpen(true)}
                className="h-6 px-2 text-xs"
                title="Editar item"
              >
                <Edit className="h-3 w-3 mr-1" />
                Editar
              </Button>

              {/* Adicionar nível - Ação secundária */}
              <Button
                variant="ghost"
                size="sm"
                onClick={() => setIsAddLevelDialogOpen(true)}
                className="h-6 w-6 p-0"
                title="Adicionar nível"
              >
                <Plus className="h-3 w-3" />
              </Button>

              {/* Adicionar item filho - Só aparece em itens raiz (sem pai) */}
              {!item.parentItemId && (
                <Button
                  variant="ghost"
                  size="sm"
                  onClick={() => setIsAddChildItemDialogOpen(true)}
                  className="h-6 px-2 text-xs"
                  title="Adicionar item filho"
                >
                  <Plus className="h-3 w-3 mr-0.5" />
                  Item
                </Button>
              )}

              {/* Dropdown com outras ações */}
              <DropdownMenu>
                <DropdownMenuTrigger asChild>
                  <Button
                    variant="ghost"
                    size="sm"
                    className="h-6 w-6 p-0"
                    title="Mais opções"
                  >
                    <MoreVertical className="h-3 w-3" />
                  </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent align="end">
                  <DropdownMenuItem onClick={() => setIsArticlesSheetOpen(true)}>
                    <FileText className="h-4 w-4 mr-2" />
                    Ver artigos científicos
                  </DropdownMenuItem>
                  <DropdownMenuItem
                    onClick={handleDuplicate}
                    disabled={createItem.isPending}
                  >
                    <Copy className="h-4 w-4 mr-2" />
                    Duplicar item
                  </DropdownMenuItem>
                  <DropdownMenuSeparator />
                  <DropdownMenuItem
                    onClick={() => setIsDeleteDialogOpen(true)}
                    className="text-destructive focus:text-destructive"
                  >
                    <Trash2 className="h-4 w-4 mr-2" />
                    Excluir item
                  </DropdownMenuItem>
                </DropdownMenuContent>
              </DropdownMenu>
            </div>
          </div>
        </CardHeader>

        {sortedLevels.length > 0 && (
          <CardContent className="pt-0 px-3 pb-2">
            <div className="flex flex-wrap gap-1.5">
              {sortedLevels.map((level) => (
                <div key={level.id} id={`level-${level.id}`} className="rounded transition-all">
                  <ScoreLevelBadge
                    level={level}
                    size="sm"
                    onEdit={() => setEditingLevel(level)}
                    onDelete={() => setDeletingLevel(level)}
                  />
                </div>
              ))}
            </div>
          </CardContent>
        )}


        {hasClinicalInfo && (
          <CardContent className="pt-0 px-3 pb-2">
            <Accordion
              type="single"
              collapsible
              className="border-t"
              value={accordionValue}
              onValueChange={setAccordionValue}
            >
              <AccordionItem value="clinical-info" className="border-0">
                <AccordionTrigger className="py-1.5 text-xs hover:no-underline text-left">
                  <div className="flex items-center gap-2">
                    <Info className="h-3 w-3" />
                    <span>Informações Clínicas</span>
                    {item.lastReview && (
                      <Badge variant="outline" className="ml-auto mr-4 text-[10px] px-1.5 py-0">
                        <Calendar className="h-2.5 w-2.5 mr-0.5" />
                        {format(new Date(item.lastReview), "dd/MM/yy", { locale: ptBR })}
                      </Badge>
                    )}
                  </div>
                </AccordionTrigger>
                <AccordionContent>
                  <div className="space-y-2 text-xs">
                    {item.clinicalRelevance && (
                      <div>
                        <h4 className="font-semibold mb-1 text-foreground">
                          Relevância Clínica
                        </h4>
                        <p className="text-muted-foreground whitespace-pre-wrap">
                          {item.clinicalRelevance}
                        </p>
                      </div>
                    )}
                    {item.patientExplanation && (
                      <div>
                        <h4 className="font-semibold mb-1 text-foreground">
                          Explicação para o Paciente
                        </h4>
                        <p className="text-muted-foreground whitespace-pre-wrap">
                          {item.patientExplanation}
                        </p>
                      </div>
                    )}
                    {item.conduct && (
                      <div>
                        <h4 className="font-semibold mb-1 text-foreground">
                          Conduta Clínica
                        </h4>
                        <p className="text-muted-foreground whitespace-pre-wrap">
                          {item.conduct}
                        </p>
                      </div>
                    )}
                  </div>
                </AccordionContent>
              </AccordionItem>
            </Accordion>
          </CardContent>
        )}
      </Card>

      {/* Edit Item Dialog */}
      <ScoreItemDialog
        open={isEditDialogOpen}
        onOpenChange={setIsEditDialogOpen}
        subgroupId={item.subgroupId}
        item={item}
      />

      {/* Add Level Dialog */}
      <ScoreLevelDialog
        open={isAddLevelDialogOpen}
        onOpenChange={setIsAddLevelDialogOpen}
        itemId={item.id}
      />

      {/* Add Child Item Dialog */}
      <ScoreItemDialog
        open={isAddChildItemDialogOpen}
        onOpenChange={setIsAddChildItemDialogOpen}
        subgroupId={item.subgroupId}
        initialParentItemId={item.id}
      />

      {/* Edit Level Dialog */}
      {editingLevel && (
        <ScoreLevelDialog
          open={!!editingLevel}
          onOpenChange={(open) => !open && setEditingLevel(null)}
          itemId={item.id}
          level={editingLevel}
        />
      )}

      {/* Delete Item Dialog */}
      <DeleteConfirmDialog
        open={isDeleteDialogOpen}
        onOpenChange={setIsDeleteDialogOpen}
        onConfirm={handleDeleteItem}
        title="Excluir item de escore?"
        description={`Tem certeza que deseja excluir "${item.name}"? Isso também excluirá todos os níveis associados.`}
        isLoading={deleteItem.isPending}
      />

      {/* Delete Level Dialog */}
      {deletingLevel && (
        <DeleteConfirmDialog
          open={!!deletingLevel}
          onOpenChange={(open) => !open && setDeletingLevel(null)}
          onConfirm={() => handleDeleteLevel(deletingLevel)}
          title="Excluir nível?"
          description={`Tem certeza que deseja excluir o nível "${deletingLevel.name}"?`}
          isLoading={deleteLevel.isPending}
        />
      )}

      {/* Articles Sheet */}
      <ScoreItemArticlesSheet
        scoreItemId={item.id}
        scoreItemName={item.name}
        open={isArticlesSheetOpen}
        onOpenChange={setIsArticlesSheetOpen}
      />
    </>
  )
}

// Memoize component to prevent unnecessary re-renders during expansion
// CRITICAL: We use default shallow comparison now because the custom comparison
// was not checking item.levels array, causing badges to not update when levels changed.
// Default shallow comparison checks all top-level properties including object/array references.
export const ScoreItemCard = memo(ScoreItemCardComponent)
