'use client'

import { useState } from 'react'
import { Edit, Trash2, Plus } from 'lucide-react'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { Card, CardContent, CardHeader } from '@/components/ui/card'
import { ScoreItem, ScoreLevel, useDeleteScoreItem, useDeleteScoreLevel } from '@/lib/api/score-api'
import { ScoreLevelBadge } from './ScoreLevelBadge'
import { ScoreItemDialog } from './ScoreItemDialog'
import { ScoreLevelDialog } from './ScoreLevelDialog'
import { DeleteConfirmDialog } from './DeleteConfirmDialog'
import { toast } from 'sonner'

interface ScoreItemCardProps {
  item: ScoreItem
  isExpanded?: boolean
}

export function ScoreItemCard({ item, isExpanded }: ScoreItemCardProps) {
  const [isEditDialogOpen, setIsEditDialogOpen] = useState(false)
  const [isAddLevelDialogOpen, setIsAddLevelDialogOpen] = useState(false)
  const [isDeleteDialogOpen, setIsDeleteDialogOpen] = useState(false)
  const [editingLevel, setEditingLevel] = useState<ScoreLevel | null>(null)
  const [deletingLevel, setDeletingLevel] = useState<ScoreLevel | null>(null)

  const deleteItem = useDeleteScoreItem()
  const deleteLevel = useDeleteScoreLevel()

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

  return (
    <>
      <Card>
        <CardHeader className="pb-3">
          <div className="flex items-start justify-between">
            <div className="flex-1">
              <div className="flex items-center gap-2">
                <h4 className="font-semibold">{item.name}</h4>
                {item.unit && (
                  <Badge variant="secondary" className="text-xs">
                    {item.unit}
                  </Badge>
                )}
              </div>
              <div className="flex items-center gap-4 mt-1 text-sm text-muted-foreground">
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
            <p className="text-sm text-muted-foreground text-center py-2">
              Nenhum nível cadastrado
            </p>
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
