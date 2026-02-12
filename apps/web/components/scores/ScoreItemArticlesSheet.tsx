'use client'

import { useState, useEffect, useMemo } from 'react'
import { FileText, Edit, X } from 'lucide-react'
import { Button } from '@/components/ui/button'
import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
} from '@/components/ui/sheet'
import { ScoreItemArticles } from './ScoreItemArticles'
import { ScoreItemArticlesSelector } from './ScoreItemArticlesSelector'
import { useArticlesForScoreItem } from '@/lib/api/article-api'
import { useAddScoreItems, useRemoveScoreItems } from '@/lib/api/article-api'
import { useQueryClient } from '@tanstack/react-query'
import { toast } from 'sonner'

interface ScoreItemArticlesSheetProps {
  scoreItemId: string
  scoreItemName: string
  open: boolean
  onOpenChange: (open: boolean) => void
}

export function ScoreItemArticlesSheet({
  scoreItemId,
  scoreItemName,
  open,
  onOpenChange,
}: ScoreItemArticlesSheetProps) {
  const queryClient = useQueryClient()
  const [isEditing, setIsEditing] = useState(false)
  const [selectedArticleIds, setSelectedArticleIds] = useState<string[]>([])

  // Fetch articles for this score item (only when sheet is open)
  const { data: articles = [], isLoading } = useArticlesForScoreItem(scoreItemId, open)

  // Mutations
  const addScoreItemsMutation = useAddScoreItems()
  const removeScoreItemsMutation = useRemoveScoreItems()

  // Initialize selected IDs when articles load and sheet opens
  // Use a stable key to prevent infinite loops (array reference changes on every render)
  const articlesIdsKey = useMemo(
    () => articles.map((a) => a.id).sort().join(','),
    [articles]
  )

  useEffect(() => {
    // Only update when sheet is open
    if (open) {
      if (articles.length > 0) {
        setSelectedArticleIds(articles.map((a) => a.id))
      } else {
        setSelectedArticleIds([])
      }
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [articlesIdsKey, open])

  // Reset editing state when sheet closes
  useEffect(() => {
    if (!open) {
      setIsEditing(false)
    }
  }, [open])

  const handleSave = async () => {
    try {
      const currentIds = articles.map((a) => a.id)
      const toAdd = selectedArticleIds.filter((id) => !currentIds.includes(id))
      const toRemove = currentIds.filter((id) => !selectedArticleIds.includes(id))

      // Add new articles (calling from article side, with score item ID)
      for (const articleId of toAdd) {
        await addScoreItemsMutation.mutateAsync({
          id: articleId,
          scoreItemIds: [scoreItemId],
        })
      }

      // Remove articles
      for (const articleId of toRemove) {
        await removeScoreItemsMutation.mutateAsync({
          id: articleId,
          scoreItemIds: [scoreItemId],
        })
      }

      // Invalidate caches
      queryClient.invalidateQueries({ queryKey: ['score-items', scoreItemId, 'articles'] })
      queryClient.invalidateQueries({ queryKey: ['articles'] })
      queryClient.invalidateQueries({ queryKey: ['score-groups'] })

      toast.success('Artigos atualizados com sucesso')
      setIsEditing(false)
    } catch (error) {
      toast.error('Erro ao atualizar artigos')
    }
  }

  const handleCancel = () => {
    // Reset to current state
    setSelectedArticleIds(articles.map((a) => a.id))
    setIsEditing(false)
  }

  const isSaving = addScoreItemsMutation.isPending || removeScoreItemsMutation.isPending

  return (
    <Sheet open={open} onOpenChange={onOpenChange}>
      <SheetContent className="w-full sm:max-w-2xl overflow-y-auto">
        <SheetHeader>
          <div className="flex items-start justify-between pr-6">
            <div>
              <SheetTitle className="flex items-center gap-2">
                <FileText className="h-5 w-5" />
                Artigos Científicos
              </SheetTitle>
              <SheetDescription className="mt-1">
                {scoreItemName}
              </SheetDescription>
            </div>
          </div>
        </SheetHeader>

        <div className="mt-6 space-y-4">
          {!isEditing ? (
            <>
              {/* View Mode */}
              <div className="flex justify-end">
                <Button
                  variant="outline"
                  size="sm"
                  onClick={() => setIsEditing(true)}
                  className="gap-2"
                >
                  <Edit className="h-4 w-4" />
                  Editar Artigos
                </Button>
              </div>

              <ScoreItemArticles articles={articles} isLoading={isLoading} />
            </>
          ) : (
            <>
              {/* Edit Mode */}
              <div className="flex justify-end gap-2">
                <Button
                  variant="outline"
                  size="sm"
                  onClick={handleCancel}
                  disabled={isSaving}
                >
                  Cancelar
                </Button>
                <Button
                  variant="default"
                  size="sm"
                  onClick={handleSave}
                  disabled={isSaving}
                >
                  {isSaving ? 'Salvando...' : 'Salvar'}
                </Button>
              </div>

              <ScoreItemArticlesSelector
                selectedArticleIds={selectedArticleIds}
                onChange={setSelectedArticleIds}
              />
            </>
          )}
        </div>
      </SheetContent>
    </Sheet>
  )
}
