'use client'

import { useParams } from 'next/navigation'
import { useMethodTree, useUnassignedScoreItems, useAssignItemToPillar } from '@/lib/api/method-api'
import { MethodTreeView } from '@/components/methods/MethodTreeView'
import { UnassignedItemsPanel } from '@/components/methods/UnassignedItemsPanel'
import { MethodLetterDialog } from '@/components/methods/MethodLetterDialog'
import { MethodDialog } from '@/components/methods/MethodDialog'
import { DeleteConfirmDialog } from '@/components/methods/DeleteConfirmDialog'
import { DraggableScoreItem } from '@/components/methods/DraggableScoreItem'
import { Button } from '@/components/ui/button'
import { Skeleton } from '@/components/ui/skeleton'
import { Card, CardContent } from '@/components/ui/card'
import { AlertCircle, ArrowLeft, Eye, EyeOff, Plus, Edit, Trash2 } from 'lucide-react'
import Link from 'next/link'
import { useState } from 'react'
import { useAuth } from '@/lib/use-auth'
import { useRouter } from 'next/navigation'
import { useDeleteMethod } from '@/lib/api/method-api'
import { toast } from 'sonner'
import {
  DndContext,
  DragOverlay,
  PointerSensor,
  useSensor,
  useSensors,
  DragEndEvent,
  DragStartEvent,
} from '@dnd-kit/core'
import type { ScoreItem, Method } from '@plenya/types'

export default function MethodDashboardPage() {
  const params = useParams()
  const methodId = params.methodId as string

  const { data: method, isLoading: isLoadingMethod, error: methodError } = useMethodTree(methodId)
  const { data: unassignedItems, isLoading: isLoadingItems } = useUnassignedScoreItems()
  const { user } = useAuth()
  const router = useRouter()
  const deleteMethod = useDeleteMethod()

  const [showUnassigned, setShowUnassigned] = useState(true)
  const [createLetterDialogOpen, setCreateLetterDialogOpen] = useState(false)
  const [editMethodDialogOpen, setEditMethodDialogOpen] = useState(false)
  const [deleteConfirmOpen, setDeleteConfirmOpen] = useState(false)
  const [activeItem, setActiveItem] = useState<ScoreItem | null>(null)

  const isAdmin = user?.roles?.includes('admin')
  const assignItem = useAssignItemToPillar()

  const sensors = useSensors(
    useSensor(PointerSensor, {
      activationConstraint: { distance: 8 },
    })
  )

  const findItemInMethod = (method: Method, itemId: string, unassignedItems: ScoreItem[]): ScoreItem | null => {
    // Search in method tree
    if (method.letters) {
      for (const letter of method.letters) {
        if (!letter.pillars) continue
        for (const pillar of letter.pillars) {
          if (!pillar.scoreItems) continue
          const item = pillar.scoreItems.find((i) => i.id === itemId)
          if (item) return item
        }
      }
    }

    // Search in unassigned items
    const unassigned = unassignedItems.find((i) => i.id === itemId)
    if (unassigned) return unassigned

    return null
  }

  const handleDragStart = (event: DragStartEvent) => {
    if (!method) return
    const itemId = event.active.id as string
    const item = findItemInMethod(method, itemId, unassignedItems || [])
    setActiveItem(item)
  }

  const handleDragEnd = async (event: DragEndEvent) => {
    const { active, over } = event
    setActiveItem(null)

    if (!over || !method) return

    const itemId = active.id as string
    const targetPillarId = over.id as string

    // Check if item is already in target pillar
    const item = findItemInMethod(method, itemId, unassignedItems || [])
    if (!item) {
      toast.error('Item não encontrado')
      return
    }

    const isAlreadyAssigned = item.methodPillars?.some((p) => p.id === targetPillarId)
    if (isAlreadyAssigned) {
      toast.info('Item já está neste pilar')
      return
    }

    try {
      await assignItem.mutateAsync({ itemId, pillarId: targetPillarId })
      toast.success('Item atribuído ao pilar')
    } catch (error) {
      toast.error('Erro ao atribuir item')
      console.error('Assignment error:', error)
    }
  }

  const handleDeleteMethod = async () => {
    try {
      await deleteMethod.mutateAsync(methodId)
      toast.success('Metodologia deletada com sucesso')
      router.push('/methods')
    } catch (error: any) {
      toast.error(error?.message || 'Erro ao deletar metodologia')
    }
  }

  if (isLoadingMethod) {
    return (
      <div className="space-y-6">
        <div className="flex items-center gap-4">
          <Skeleton className="h-10 w-10" />
          <div className="space-y-2">
            <Skeleton className="h-8 w-[400px]" />
            <Skeleton className="h-4 w-[600px]" />
          </div>
        </div>
        <Skeleton className="h-[400px]" />
      </div>
    )
  }

  if (methodError || !method) {
    return (
      <div className="space-y-6">
        <div className="flex items-center gap-2">
          <Button variant="ghost" size="sm" asChild>
            <Link href="/methods">
              <ArrowLeft className="h-4 w-4 mr-1" />
              Voltar
            </Link>
          </Button>
        </div>
        <Card className="border-destructive">
          <CardContent className="flex items-center gap-3 p-6">
            <AlertCircle className="h-5 w-5 text-destructive" />
            <div>
              <p className="font-medium">Erro ao carregar metodologia</p>
              <p className="text-sm text-muted-foreground">
                A metodologia solicitada não foi encontrada ou ocorreu um erro ao carregá-la.
              </p>
            </div>
          </CardContent>
        </Card>
      </div>
    )
  }

  const unassignedCount = unassignedItems?.length || 0

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="space-y-4">
        <div className="flex items-center gap-2">
          <Button variant="ghost" size="sm" asChild>
            <Link href="/methods">
              <ArrowLeft className="h-4 w-4 mr-1" />
              Voltar
            </Link>
          </Button>
        </div>

        <div className="flex items-start justify-between">
          <div className="space-y-1">
            <h1 className="text-3xl font-bold tracking-tight" style={{ color: method.color || undefined }}>
              {method.name}
            </h1>
            {method.description && (
              <p className="text-muted-foreground max-w-3xl">{method.description}</p>
            )}
            {method.version && (
              <p className="text-sm text-muted-foreground">Versão: {method.version}</p>
            )}
          </div>

          <div className="flex gap-2">
            {isAdmin && (
              <>
                <Button
                  variant="outline"
                  size="sm"
                  onClick={() => setEditMethodDialogOpen(true)}
                >
                  <Edit className="h-4 w-4 mr-1" />
                  Editar
                </Button>
                <Button
                  variant="outline"
                  size="sm"
                  onClick={() => setDeleteConfirmOpen(true)}
                  className="hover:bg-destructive hover:text-destructive-foreground"
                >
                  <Trash2 className="h-4 w-4 mr-1" />
                  Deletar
                </Button>
                <Button size="sm" onClick={() => setCreateLetterDialogOpen(true)}>
                  <Plus className="h-4 w-4 mr-1" />
                  Nova Letra
                </Button>
              </>
            )}
            <Button
              variant="outline"
              size="sm"
              onClick={() => setShowUnassigned(!showUnassigned)}
            >
              {showUnassigned ? (
                <>
                  <EyeOff className="h-4 w-4 mr-1" />
                  Ocultar Não Atribuídos
                </>
              ) : (
                <>
                  <Eye className="h-4 w-4 mr-1" />
                  Mostrar Não Atribuídos
                </>
              )}
              {unassignedCount > 0 && (
                <span className="ml-1">({unassignedCount})</span>
              )}
            </Button>
          </div>
        </div>
      </div>

      {/* Stats */}
      <div className="grid grid-cols-1 md:grid-cols-4 gap-4">
        <Card>
          <CardContent className="pt-6">
            <div className="text-2xl font-bold">{method.letters?.length || 0}</div>
            <p className="text-xs text-muted-foreground">Letras</p>
          </CardContent>
        </Card>
        <Card>
          <CardContent className="pt-6">
            <div className="text-2xl font-bold">
              {method.letters?.reduce((sum, l) => sum + (l.pillars?.length || 0), 0) || 0}
            </div>
            <p className="text-xs text-muted-foreground">Pilares</p>
          </CardContent>
        </Card>
        <Card>
          <CardContent className="pt-6">
            <div className="text-2xl font-bold">
              {method.letters?.reduce(
                (sum, l) =>
                  sum + (l.pillars?.reduce((pSum, p) => pSum + (p.scoreItems?.length || 0), 0) || 0),
                0
              ) || 0}
            </div>
            <p className="text-xs text-muted-foreground">Itens Atribuídos</p>
          </CardContent>
        </Card>
        <Card>
          <CardContent className="pt-6">
            <div className="text-2xl font-bold">{unassignedCount}</div>
            <p className="text-xs text-muted-foreground">Itens Não Atribuídos</p>
          </CardContent>
        </Card>
      </div>

      {/* Drag and Drop Context wrapping both tree and unassigned panel */}
      <DndContext sensors={sensors} onDragStart={handleDragStart} onDragEnd={handleDragEnd}>
        {/* Method Tree */}
        <MethodTreeView method={method} unassignedItems={unassignedItems} />

        {/* Unassigned Items Panel */}
        {showUnassigned && (
          <div>
            {isLoadingItems ? (
              <Skeleton className="h-[300px]" />
            ) : (
              <UnassignedItemsPanel items={unassignedItems || []} />
            )}
          </div>
        )}

        {/* Drag Overlay */}
        <DragOverlay>
          {activeItem && (
            <div className="opacity-80">
              <DraggableScoreItem item={activeItem} isDragging />
            </div>
          )}
        </DragOverlay>
      </DndContext>

      {/* Create Letter Dialog */}
      <MethodLetterDialog
        open={createLetterDialogOpen}
        onOpenChange={setCreateLetterDialogOpen}
        methodId={methodId}
      />

      {/* Edit Method Dialog */}
      {method && (
        <MethodDialog
          open={editMethodDialogOpen}
          onOpenChange={setEditMethodDialogOpen}
          method={method}
        />
      )}

      {/* Delete Method Confirmation */}
      <DeleteConfirmDialog
        open={deleteConfirmOpen}
        onOpenChange={setDeleteConfirmOpen}
        onConfirm={handleDeleteMethod}
        title="Deletar Metodologia"
        description={`Tem certeza que deseja deletar "${method?.name}"? Todas as letras, pilares e associações com itens serão deletados. Esta ação não pode ser desfeita.`}
        isDeleting={deleteMethod.isPending}
      />
    </div>
  )
}
