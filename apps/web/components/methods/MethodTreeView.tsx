'use client'

import { useState } from 'react'
import { toast } from 'sonner'
import type { Method, ScoreItem, MethodLetter, MethodPillar } from '@plenya/types'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { ChevronDown, ChevronRight, Plus, Edit, Trash2 } from 'lucide-react'
import { MethodPillarDropZone } from './MethodPillarDropZone'
import { DraggableScoreItem } from './DraggableScoreItem'
import { MethodPillarDialog } from './MethodPillarDialog'
import { MethodLetterDialog } from './MethodLetterDialog'
import { DeleteConfirmDialog } from './DeleteConfirmDialog'
import { useAuth } from '@/lib/use-auth'
import {
  useDeleteMethodLetter,
  useDeleteMethodPillar,
} from '@/lib/api/method-api'

interface MethodTreeViewProps {
  method: Method
  unassignedItems?: ScoreItem[]
}

export function MethodTreeView({ method, unassignedItems = [] }: MethodTreeViewProps) {
  const [expandedLetters, setExpandedLetters] = useState<Set<string>>(new Set())
  const [expandedPillars, setExpandedPillars] = useState<Set<string>>(new Set())

  // Create states
  const [createPillarDialogOpen, setCreatePillarDialogOpen] = useState(false)
  const [selectedLetterId, setSelectedLetterId] = useState<string | null>(null)

  // Edit states
  const [editLetterDialogOpen, setEditLetterDialogOpen] = useState(false)
  const [editingLetter, setEditingLetter] = useState<MethodLetter | null>(null)
  const [editPillarDialogOpen, setEditPillarDialogOpen] = useState(false)
  const [editingPillar, setEditingPillar] = useState<MethodPillar | null>(null)

  // Delete states
  const [deleteConfirmOpen, setDeleteConfirmOpen] = useState(false)
  const [deleteType, setDeleteType] = useState<'letter' | 'pillar' | null>(null)
  const [deleteTarget, setDeleteTarget] = useState<{ id: string; name: string } | null>(null)

  const deleteLetter = useDeleteMethodLetter()
  const deletePillar = useDeleteMethodPillar()
  const { user } = useAuth()

  const isAdmin = user?.roles?.includes('admin')

  const toggleLetter = (letterId: string) => {
    setExpandedLetters((prev) => {
      const newSet = new Set(prev)
      if (newSet.has(letterId)) {
        newSet.delete(letterId)
      } else {
        newSet.add(letterId)
      }
      return newSet
    })
  }

  const togglePillar = (pillarId: string) => {
    setExpandedPillars((prev) => {
      const newSet = new Set(prev)
      if (newSet.has(pillarId)) {
        newSet.delete(pillarId)
      } else {
        newSet.add(pillarId)
      }
      return newSet
    })
  }

  const expandAll = () => {
    const allLetters = new Set(method.letters?.map((l) => l.id) || [])
    const allPillars = new Set(
      method.letters?.flatMap((l) => l.pillars?.map((p) => p.id) || []) || []
    )
    setExpandedLetters(allLetters)
    setExpandedPillars(allPillars)
  }

  const collapseAll = () => {
    setExpandedLetters(new Set())
    setExpandedPillars(new Set())
  }

  // Edit handlers
  const handleEditLetter = (letter: MethodLetter) => {
    setEditingLetter(letter)
    setEditLetterDialogOpen(true)
  }

  const handleEditPillar = (pillar: MethodPillar) => {
    setEditingPillar(pillar)
    setEditPillarDialogOpen(true)
  }

  // Delete handlers
  const handleDeleteLetter = (letter: MethodLetter) => {
    setDeleteType('letter')
    setDeleteTarget({ id: letter.id, name: letter.name })
    setDeleteConfirmOpen(true)
  }

  const handleDeletePillar = (pillar: MethodPillar) => {
    setDeleteType('pillar')
    setDeleteTarget({ id: pillar.id, name: pillar.name })
    setDeleteConfirmOpen(true)
  }

  const confirmDelete = async () => {
    if (!deleteTarget || !deleteType) return

    try {
      if (deleteType === 'letter') {
        await deleteLetter.mutateAsync(deleteTarget.id)
        toast.success('Letra deletada com sucesso')
      } else {
        await deletePillar.mutateAsync(deleteTarget.id)
        toast.success('Pilar deletado com sucesso')
      }
      setDeleteConfirmOpen(false)
      setDeleteTarget(null)
      setDeleteType(null)
    } catch (error: any) {
      toast.error(error?.message || 'Erro ao deletar')
    }
  }

  if (!method.letters || method.letters.length === 0) {
    return (
      <Card>
        <CardContent className="py-8 text-center text-muted-foreground">
          Nenhuma letra cadastrada. Crie letras para começar a organizar os itens.
        </CardContent>
      </Card>
    )
  }

  return (
    <>
      <div className="space-y-4">
        {/* Expand/Collapse Controls */}
        <div className="flex gap-2">
          <Button variant="outline" size="sm" onClick={expandAll}>
            Expandir Tudo
          </Button>
          <Button variant="outline" size="sm" onClick={collapseAll}>
            Recolher Tudo
          </Button>
        </div>

        {/* Letters */}
        {method.letters.map((letter) => (
          <LetterCard
            key={letter.id}
            letter={letter}
            isExpanded={expandedLetters.has(letter.id)}
            onToggle={() => toggleLetter(letter.id)}
            expandedPillars={expandedPillars}
            onTogglePillar={togglePillar}
            isAdmin={isAdmin}
            onCreatePillar={() => {
              setSelectedLetterId(letter.id)
              setCreatePillarDialogOpen(true)
            }}
            onEditLetter={() => handleEditLetter(letter)}
            onDeleteLetter={() => handleDeleteLetter(letter)}
            onEditPillar={handleEditPillar}
            onDeletePillar={handleDeletePillar}
          />
        ))}
      </div>

      {/* Create Pillar Dialog */}
      {selectedLetterId && (
        <MethodPillarDialog
          open={createPillarDialogOpen}
          onOpenChange={setCreatePillarDialogOpen}
          letterId={selectedLetterId}
        />
      )}

      {/* Edit Letter Dialog */}
      {editingLetter && (
        <MethodLetterDialog
          open={editLetterDialogOpen}
          onOpenChange={(open) => {
            setEditLetterDialogOpen(open)
            if (!open) setEditingLetter(null)
          }}
          methodId={method.id}
          letter={editingLetter}
        />
      )}

      {/* Edit Pillar Dialog */}
      {editingPillar && (
        <MethodPillarDialog
          open={editPillarDialogOpen}
          onOpenChange={(open) => {
            setEditPillarDialogOpen(open)
            if (!open) setEditingPillar(null)
          }}
          letterId={editingPillar.letterId}
          pillar={editingPillar}
        />
      )}

      {/* Delete Confirmation Dialog */}
      <DeleteConfirmDialog
        open={deleteConfirmOpen}
        onOpenChange={setDeleteConfirmOpen}
        onConfirm={confirmDelete}
        title={`Deletar ${deleteType === 'letter' ? 'Letra' : 'Pilar'}`}
        description={`Tem certeza que deseja deletar "${deleteTarget?.name}"? ${
          deleteType === 'letter'
            ? 'Todos os pilares desta letra também serão deletados.'
            : 'Os itens atribuídos a este pilar não serão deletados, apenas a associação.'
        }`}
        isDeleting={deleteLetter.isPending || deletePillar.isPending}
      />
    </>
  )
}

interface LetterCardProps {
  letter: MethodLetter
  isExpanded: boolean
  onToggle: () => void
  expandedPillars: Set<string>
  onTogglePillar: (pillarId: string) => void
  isAdmin: boolean
  onCreatePillar: () => void
  onEditLetter: () => void
  onDeleteLetter: () => void
  onEditPillar: (pillar: MethodPillar) => void
  onDeletePillar: (pillar: MethodPillar) => void
}

function LetterCard({ letter, isExpanded, onToggle, expandedPillars, onTogglePillar, isAdmin, onCreatePillar, onEditLetter, onDeleteLetter, onEditPillar, onDeletePillar }: LetterCardProps) {
  const totalItems = letter.pillars?.reduce((sum, p) => sum + (p.scoreItems?.length || 0), 0) || 0

  return (
    <Card>
      <CardHeader className="cursor-pointer hover:bg-accent/50 transition-colors">
        <div className="flex items-start justify-between">
          <div className="flex items-start gap-3 flex-1" onClick={onToggle}>
            <div className="mt-1">
              {isExpanded ? (
                <ChevronDown className="h-5 w-5 text-muted-foreground" />
              ) : (
                <ChevronRight className="h-5 w-5 text-muted-foreground" />
              )}
            </div>
            <div className="space-y-1 flex-1">
              <CardTitle className="flex items-center gap-2">
                {letter.icon && <span className="text-2xl">{letter.icon}</span>}
                <span style={{ color: letter.color || undefined }}>
                  {letter.code} - {letter.name}
                </span>
              </CardTitle>
              {letter.description && (
                <CardDescription>{letter.description}</CardDescription>
              )}
            </div>
          </div>
          <div className="flex items-center gap-2">
            <Badge variant="secondary">
              {letter.pillars?.length || 0} pilares • {totalItems} itens
            </Badge>
            {isAdmin && (
              <div className="flex gap-1">
                <Button
                  variant="ghost"
                  size="sm"
                  onClick={(e) => {
                    e.stopPropagation()
                    onEditLetter()
                  }}
                  className="h-8 w-8 p-0"
                >
                  <Edit className="h-4 w-4" />
                </Button>
                <Button
                  variant="ghost"
                  size="sm"
                  onClick={(e) => {
                    e.stopPropagation()
                    onDeleteLetter()
                  }}
                  className="h-8 w-8 p-0 hover:bg-destructive hover:text-destructive-foreground"
                >
                  <Trash2 className="h-4 w-4" />
                </Button>
              </div>
            )}
          </div>
        </div>
      </CardHeader>

      {isExpanded && (
        <CardContent className="space-y-3 pt-0">
          {isAdmin && (
            <Button
              variant="outline"
              size="sm"
              onClick={(e) => {
                e.stopPropagation()
                onCreatePillar()
              }}
              className="w-full"
            >
              <Plus className="h-4 w-4 mr-1" />
              Novo Pilar
            </Button>
          )}

          {letter.pillars && letter.pillars.length > 0 ? (
            letter.pillars.map((pillar) => (
              <PillarCard
                key={pillar.id}
                pillar={pillar}
                isExpanded={expandedPillars.has(pillar.id)}
                onToggle={() => onTogglePillar(pillar.id)}
                isAdmin={isAdmin}
                onEdit={() => onEditPillar(pillar)}
                onDelete={() => onDeletePillar(pillar)}
              />
            ))
          ) : (
            <p className="text-sm text-muted-foreground text-center py-4">
              Nenhum pilar cadastrado nesta letra
            </p>
          )}
        </CardContent>
      )}
    </Card>
  )
}

interface PillarCardProps {
  pillar: MethodPillar
  isExpanded: boolean
  onToggle: () => void
  isAdmin: boolean
  onEdit: () => void
  onDelete: () => void
}

function PillarCard({ pillar, isExpanded, onToggle, isAdmin, onEdit, onDelete }: PillarCardProps) {
  return (
    <Card className="border-dashed">
      <CardHeader className="pb-3 cursor-pointer hover:bg-accent/30 transition-colors">
        <div className="flex items-start justify-between gap-2">
          <div className="flex items-start gap-2 flex-1" onClick={onToggle}>
            <div className="mt-0.5">
              {isExpanded ? (
                <ChevronDown className="h-4 w-4 text-muted-foreground" />
              ) : (
                <ChevronRight className="h-4 w-4 text-muted-foreground" />
              )}
            </div>
            <div className="flex-1 space-y-1">
              <CardTitle className="text-base">{pillar.name}</CardTitle>
              {pillar.description && (
                <CardDescription className="text-xs">{pillar.description}</CardDescription>
              )}
            </div>
          </div>
          <div className="flex items-center gap-2">
            <Badge variant="outline">{pillar.scoreItems?.length || 0} itens</Badge>
            {isAdmin && (
              <div className="flex gap-1">
                <Button
                  variant="ghost"
                  size="sm"
                  onClick={(e) => {
                    e.stopPropagation()
                    onEdit()
                  }}
                  className="h-7 w-7 p-0"
                >
                  <Edit className="h-3 w-3" />
                </Button>
                <Button
                  variant="ghost"
                  size="sm"
                  onClick={(e) => {
                    e.stopPropagation()
                    onDelete()
                  }}
                  className="h-7 w-7 p-0 hover:bg-destructive hover:text-destructive-foreground"
                >
                  <Trash2 className="h-3 w-3" />
                </Button>
              </div>
            )}
          </div>
        </div>
      </CardHeader>

      {isExpanded && (
        <CardContent className="pt-0">
          <MethodPillarDropZone pillarId={pillar.id}>
            {pillar.scoreItems && pillar.scoreItems.length > 0 ? (
              <div className="space-y-2">
                {pillar.scoreItems.map((item) => (
                  <DraggableScoreItem key={item.id} item={item} pillarId={pillar.id} />
                ))}
              </div>
            ) : (
              <div className="text-sm text-muted-foreground text-center py-6 border-2 border-dashed rounded-lg">
                Arraste itens aqui para atribuí-los a este pilar
              </div>
            )}
          </MethodPillarDropZone>
        </CardContent>
      )}
    </Card>
  )
}
