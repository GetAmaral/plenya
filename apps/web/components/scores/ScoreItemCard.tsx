'use client'

import { useState, memo, useEffect } from 'react'
import { Edit, Trash2, Plus, Info, Calendar } from 'lucide-react'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { Card, CardContent, CardHeader } from '@/components/ui/card'
import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from '@/components/ui/accordion'
import { ScoreItem, ScoreLevel, useDeleteScoreItem, useDeleteScoreLevel } from '@/lib/api/score-api'
import { ScoreLevelBadge } from './ScoreLevelBadge'
import { ScoreItemDialog } from './ScoreItemDialog'
import { ScoreLevelDialog } from './ScoreLevelDialog'
import { DeleteConfirmDialog } from './DeleteConfirmDialog'
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
  const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false)
  const [editingLevel, setEditingLevel] = useState<ScoreLevel | null>(null)
  const [deletingLevel, setDeletingLevel] = useState<ScoreLevel | null>(null)
  // Initialize with empty string to keep it controlled from the start
  const [accordionValue, setAccordionValue] = useState<string>(
    expandClinicalTexts ? 'clinical-info' : ''
  )

  const deleteItem = useDeleteScoreItem()
  const deleteLevel = useDeleteScoreLevel()

  // Sync accordion state with expandClinicalTexts prop
  useEffect(() => {
    if (expandClinicalTexts) {
      setAccordionValue('clinical-info')
    } else {
      // When collapsing all, close the accordion
      setAccordionValue('')
    }
  }, [expandClinicalTexts])

  const handleDeleteItem = async () => {
    try {
      await deleteItem.mutateAsync(item.id)
      toast.success('Item excluído com sucesso')
      setIsDeleteDialogOpen(false)
    } catch (error) {
      toast.error('Erro ao excluir item')
    }
  }

  const handleDeleteLevel = async (level: ScoreLevel) => {
    try {
      await deleteLevel.mutateAsync(level.id)
      toast.success('Nível excluído com sucesso')
      setDeletingLevel(null)
    } catch (error) {
      toast.error('Erro ao excluir nível')
    }
  }

  // Sort levels by level number
  const sortedLevels = [...(item.levels || [])].sort((a, b) => a.level - b.level)

  // Check if item has any clinical information
  const hasClinicalInfo = item.clinicalRelevance || item.patientExplanation || item.conduct

  return (
    <>
      <Card>
        <CardHeader className="pb-3">
          <div className="flex items-start justify-between">
            <div className="flex-1">
              <div className="flex items-center gap-2">
                <h4 className="text-sm font-semibold text-left">{item.name}</h4>
                {item.unit && (
                  <Badge variant="secondary" className="text-xs">
                    {item.unit}
                  </Badge>
                )}
              </div>
              <div className="flex items-center gap-4 mt-0.5 text-sm text-muted-foreground text-left">
                <span>{item.points} pontos</span>
                {item.unitConversion && (
                  <span className="text-xs">{item.unitConversion}</span>
                )}
              </div>
            </div>

            <div className="flex gap-1">
              <Button
                variant="ghost"
                size="sm"
                onClick={() => setIsAddLevelDialogOpen(true)}
                className="h-8 w-8 p-0"
              >
                <Plus className="h-4 w-4" />
              </Button>
              <Button
                variant="ghost"
                size="sm"
                onClick={() => setIsEditDialogOpen(true)}
                className="h-8 w-8 p-0"
              >
                <Edit className="h-4 w-4" />
              </Button>
              <Button
                variant="ghost"
                size="sm"
                onClick={() => setIsDeleteDialogOpen(true)}
                className="h-8 w-8 p-0"
              >
                <Trash2 className="h-4 w-4" />
              </Button>
            </div>
          </div>
        </CardHeader>

        {sortedLevels.length > 0 && (
          <CardContent className="pt-0">
            <div className="flex flex-wrap gap-2">
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

        {sortedLevels.length === 0 && (
          <CardContent className="pt-0">
            <p className="text-sm text-muted-foreground text-left py-2">
              Nenhum nível cadastrado
            </p>
          </CardContent>
        )}

        {hasClinicalInfo && (
          <CardContent className="pt-0">
            <Accordion
              type="single"
              collapsible
              className="border-t"
              value={accordionValue}
              onValueChange={setAccordionValue}
            >
              <AccordionItem value="clinical-info" className="border-0">
                <AccordionTrigger className="py-2 text-sm hover:no-underline text-left">
                  <div className="flex items-center gap-2">
                    <Info className="h-4 w-4" />
                    <span>Informações Clínicas</span>
                    {item.lastReview && (
                      <Badge variant="outline" className="ml-auto mr-4 text-xs">
                        <Calendar className="h-3 w-3 mr-1" />
                        Revisado em {format(new Date(item.lastReview), "dd/MM/yyyy", { locale: ptBR })}
                      </Badge>
                    )}
                  </div>
                </AccordionTrigger>
                <AccordionContent>
                  <div className="space-y-3 text-sm">
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
    </>
  )
}

// Memoize component to prevent unnecessary re-renders during expansion
export const ScoreItemCard = memo(ScoreItemCardComponent, (prevProps, nextProps) => {
  // Return true if props are equal (don't re-render), false otherwise (re-render)
  return (
    prevProps.item.id === nextProps.item.id &&
    prevProps.item.name === nextProps.item.name &&
    prevProps.item.points === nextProps.item.points &&
    prevProps.item.unit === nextProps.item.unit &&
    prevProps.item.unitConversion === nextProps.item.unitConversion &&
    prevProps.item.order === nextProps.item.order &&
    prevProps.item.updatedAt === nextProps.item.updatedAt &&
    prevProps.isExpanded === nextProps.isExpanded &&
    prevProps.expandClinicalTexts === nextProps.expandClinicalTexts &&
    prevProps.item.lastReview === nextProps.item.lastReview
  )
})
