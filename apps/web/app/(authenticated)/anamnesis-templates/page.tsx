'use client'

import { useState, useRef, useEffect } from 'react'
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog'
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from '@/components/ui/alert-dialog'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import {
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from '@/components/ui/collapsible'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Loader2, Plus, Pencil, Trash2, ChevronDown } from 'lucide-react'
import { toast } from 'sonner'
import { useFormNavigation } from '@/lib/use-form-navigation'
import {
  getAllAnamnesisTemplates,
  getAnamnesisTemplateById,
  createAnamnesisTemplate,
  updateAnamnesisTemplate,
  updateAnamnesisTemplateItems,
  deleteAnamnesisTemplate,
  getAllScoreItems,
  AnamnesisTemplate,
  AnamnesisTemplateArea,
  ScoreItem,
} from '@/lib/api/anamnesis-templates'
import { AnamnesisTemplateItemSelector } from '@/components/anamnesis/anamnesis-template-item-selector'

const AREAS: AnamnesisTemplateArea[] = ['Medicina', 'Nutricao', 'Psicologia', 'Educacao Fisica']

export default function AnamnesisTemplatesPage() {
  const queryClient = useQueryClient()

  // State
  const [isCreateDialogOpen, setIsCreateDialogOpen] = useState(false)
  const [editingTemplate, setEditingTemplate] = useState<AnamnesisTemplate | null>(null)
  const [deleteConfirm, setDeleteConfirm] = useState<AnamnesisTemplate | null>(null)

  // Form refs
  const createFormRef = useRef<HTMLFormElement>(null)
  useFormNavigation({ formRef: createFormRef, disabled: !isCreateDialogOpen })

  // Queries
  const { data: templates = [], isLoading: templatesLoading } = useQuery({
    queryKey: ['anamnesis-templates'],
    queryFn: () => getAllAnamnesisTemplates(true),
  })

  const { data: availableScoreItems = [] } = useQuery({
    queryKey: ['score-items'],
    queryFn: getAllScoreItems,
  })

  // Mutations
  const createMutation = useMutation({
    mutationFn: createAnamnesisTemplate,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['anamnesis-templates'] })
      setIsCreateDialogOpen(false)
      toast.success('Template criado com sucesso!')
    },
    onError: (error: any) => {
      toast.error(error.response?.data?.error || 'Erro ao criar template')
    },
  })

  const deleteMutation = useMutation({
    mutationFn: deleteAnamnesisTemplate,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['anamnesis-templates'] })
      setDeleteConfirm(null)
      toast.success('Template deletado com sucesso!')
    },
    onError: (error: any) => {
      toast.error(error.response?.data?.error || 'Erro ao deletar template')
    },
  })

  // Handlers
  const handleCreate = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    const formData = new FormData(e.currentTarget)
    const name = formData.get('name') as string
    const area = formData.get('area') as AnamnesisTemplateArea

    createMutation.mutate({ name, area })
  }

  if (templatesLoading) {
    return (
      <div className="flex items-center justify-center h-96">
        <Loader2 className="h-8 w-8 animate-spin text-primary" />
      </div>
    )
  }

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Templates de Anamnese</h1>
          <p className="text-muted-foreground">
            Gerencie templates de anamnese por área médica
          </p>
        </div>
        <Button onClick={() => setIsCreateDialogOpen(true)}>
          <Plus className="mr-2 h-4 w-4" />
          Novo Template
        </Button>
      </div>

      {/* Templates Grid */}
      {templates.length === 0 ? (
        <Card>
          <CardContent className="flex flex-col items-center justify-center py-12">
            <p className="text-muted-foreground mb-4">
              Nenhum template criado ainda
            </p>
            <Button onClick={() => setIsCreateDialogOpen(true)}>
              <Plus className="mr-2 h-4 w-4" />
              Criar Primeiro Template
            </Button>
          </CardContent>
        </Card>
      ) : (
        <div className="space-y-4">
          {templates.map((template) => (
            <TemplateCard
              key={template.id}
              template={template}
              availableScoreItems={availableScoreItems}
              onEdit={() => setEditingTemplate(template)}
              onDelete={() => setDeleteConfirm(template)}
            />
          ))}
        </div>
      )}

      {/* Create Dialog */}
      <Dialog open={isCreateDialogOpen} onOpenChange={setIsCreateDialogOpen}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Novo Template de Anamnese</DialogTitle>
            <DialogDescription>
              Crie um template para facilitar o preenchimento de anamneses
            </DialogDescription>
          </DialogHeader>
          <form ref={createFormRef} onSubmit={handleCreate}>
            <div className="space-y-4 py-4">
              <div className="space-y-2">
                <label className="text-sm font-medium">Nome do Template</label>
                <Input name="name" placeholder="Ex: Anamnese Clínica Geral" required />
              </div>
              <div className="space-y-2">
                <label className="text-sm font-medium">Área</label>
                <Select name="area" required>
                  <SelectTrigger>
                    <SelectValue placeholder="Selecione a área" />
                  </SelectTrigger>
                  <SelectContent>
                    {AREAS.map((area) => (
                      <SelectItem key={area} value={area}>
                        {area}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
              </div>
            </div>
            <DialogFooter>
              <Button
                type="button"
                variant="outline"
                onClick={() => setIsCreateDialogOpen(false)}
              >
                Cancelar
              </Button>
              <Button type="submit" disabled={createMutation.isPending}>
                {createMutation.isPending && (
                  <Loader2 className="mr-2 h-4 w-4 animate-spin" />
                )}
                Criar Template
              </Button>
            </DialogFooter>
          </form>
        </DialogContent>
      </Dialog>

      {/* Edit Dialog */}
      {editingTemplate && (
        <TemplateEditDialog
          template={editingTemplate}
          availableScoreItems={availableScoreItems}
          onClose={() => setEditingTemplate(null)}
        />
      )}

      {/* Delete Confirmation */}
      <AlertDialog open={!!deleteConfirm} onOpenChange={() => setDeleteConfirm(null)}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Confirmar exclusão</AlertDialogTitle>
            <AlertDialogDescription>
              Tem certeza que deseja deletar o template &quot;{deleteConfirm?.name}&quot;? Esta
              ação não pode ser desfeita.
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>Cancelar</AlertDialogCancel>
            <AlertDialogAction
              onClick={() => deleteConfirm && deleteMutation.mutate(deleteConfirm.id)}
              className="bg-destructive text-destructive-foreground hover:bg-destructive/90"
            >
              {deleteMutation.isPending && (
                <Loader2 className="mr-2 h-4 w-4 animate-spin" />
              )}
              Deletar
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </div>
  )
}

// Template Card Component
interface TemplateCardProps {
  template: AnamnesisTemplate
  availableScoreItems: ScoreItem[]
  onEdit: () => void
  onDelete: () => void
}

function TemplateCard({ template, availableScoreItems, onEdit, onDelete }: TemplateCardProps) {
  const [isOpen, setIsOpen] = useState(false)

  const { data: fullTemplate } = useQuery({
    queryKey: ['anamnesis-template', template.id],
    queryFn: () => getAnamnesisTemplateById(template.id),
    enabled: isOpen,
  })

  const itemCount = template.items?.length || 0

  return (
    <Card>
      <CardHeader>
        <div className="flex items-start justify-between">
          <div className="space-y-1 flex-1">
            <CardTitle>{template.name}</CardTitle>
            <CardDescription>{template.area}</CardDescription>
          </div>
          <div className="flex gap-2">
            <Button variant="ghost" size="icon" onClick={onEdit}>
              <Pencil className="h-4 w-4" />
            </Button>
            <Button variant="ghost" size="icon" onClick={onDelete}>
              <Trash2 className="h-4 w-4 text-destructive" />
            </Button>
          </div>
        </div>
      </CardHeader>
      <CardContent>
        <Collapsible open={isOpen} onOpenChange={setIsOpen}>
          <CollapsibleTrigger asChild>
            <Button variant="outline" className="w-full justify-between">
              <span>{itemCount} Score Items</span>
              <ChevronDown
                className={`h-4 w-4 transition-transform ${isOpen ? 'rotate-180' : ''}`}
              />
            </Button>
          </CollapsibleTrigger>
          <CollapsibleContent className="mt-4">
            {((fullTemplate?.items || template.items) && (fullTemplate?.items || template.items).length > 0) ? (
              <div className="space-y-3">
                {(fullTemplate?.items || template.items)
                  .sort((a, b) => a.order - b.order)
                  .map((item, index, sortedItems) => {
                    // Buscar o scoreItem completo com relações do availableScoreItems
                    const fullScoreItem = availableScoreItems.find(
                      (si) => si.id === item.scoreItemId
                    )

                    if (!fullScoreItem) return null

                    // Verificar se deve mostrar grupo/subgrupo comparando com item anterior
                    const prevItem = index > 0 ? sortedItems[index - 1] : null
                    const prevScoreItem = prevItem
                      ? availableScoreItems.find((si) => si.id === prevItem.scoreItemId)
                      : null

                    const groupId = fullScoreItem.subgroup?.group?.id
                    const subgroupId = fullScoreItem.subgroup?.id
                    const prevGroupId = prevScoreItem?.subgroup?.group?.id
                    const prevSubgroupId = prevScoreItem?.subgroup?.id

                    const showGroup = groupId && groupId !== prevGroupId
                    const showSubgroup = subgroupId && (showGroup || subgroupId !== prevSubgroupId)

                    const groupName = fullScoreItem.subgroup?.group?.name
                    const subgroupName = fullScoreItem.subgroup?.name
                    const isChild = !!fullScoreItem.parentItemId

                    return (
                      <div key={item.id} className="space-y-2">
                        {/* Cabeçalho do Grupo */}
                        {showGroup && groupName && (
                          <div className="bg-primary text-primary-foreground rounded-lg px-3 py-2">
                            <div className="flex items-center gap-2 min-w-0">
                              <div className="flex items-center justify-center w-6 h-6 rounded-full bg-primary-foreground/20 text-xs font-bold flex-shrink-0">
                                {index + 1}
                              </div>
                              <h4 className="text-sm font-bold break-words min-w-0 flex-1">
                                {groupName}
                              </h4>
                            </div>
                          </div>
                        )}

                        {/* Cabeçalho do Subgrupo */}
                        {showSubgroup && subgroupName && (
                          <div className="bg-gradient-to-br from-muted to-muted/50 rounded-lg p-2 border border-primary/20">
                            <div className="flex items-center gap-2 min-w-0">
                              <h5 className="text-xs font-bold break-words min-w-0 flex-1">
                                {subgroupName}
                              </h5>
                            </div>
                          </div>
                        )}

                        {/* Item */}
                        <div
                          className="p-2 border rounded-lg bg-card text-sm"
                          style={{
                            marginLeft: isChild ? '20px' : undefined,
                            borderLeft: isChild ? '4px solid hsl(var(--primary))' : undefined,
                            paddingLeft: isChild ? '12px' : undefined,
                          }}
                        >
                          <div className="font-medium break-words">{fullScoreItem.name}</div>
                        </div>
                      </div>
                    )
                  })}
              </div>
            ) : (
              <p className="text-sm text-muted-foreground text-center py-4">
                Nenhum item adicionado
              </p>
            )}
          </CollapsibleContent>
        </Collapsible>
      </CardContent>
    </Card>
  )
}

// Template Edit Dialog Component
interface TemplateEditDialogProps {
  template: AnamnesisTemplate
  availableScoreItems: ScoreItem[]
  onClose: () => void
}

function TemplateEditDialog({
  template,
  availableScoreItems,
  onClose,
}: TemplateEditDialogProps) {
  const queryClient = useQueryClient()
  const [name, setName] = useState(template.name)
  const [area, setArea] = useState<AnamnesisTemplateArea>(template.area)
  const [selectedScoreItems, setSelectedScoreItems] = useState<ScoreItem[]>([])

  // Buscar template completo
  const { data: fullTemplate, isLoading } = useQuery({
    queryKey: ['anamnesis-template', template.id],
    queryFn: () => getAnamnesisTemplateById(template.id),
  })

  // Inicializar items selecionados quando carregar
  useEffect(() => {
    if (fullTemplate?.items && availableScoreItems.length > 0) {
      const items = fullTemplate.items
        .sort((a, b) => a.order - b.order)
        .map((item) => {
          // Buscar o scoreItem completo com relações do availableScoreItems
          return availableScoreItems.find((si) => si.id === item.scoreItemId)
        })
        .filter((item): item is ScoreItem => item !== undefined)
      setSelectedScoreItems(items)
    }
  }, [fullTemplate, availableScoreItems])

  // Mutation para atualizar items
  const updateItemsMutation = useMutation({
    mutationFn: (items: { scoreItemId: string; order: number }[]) =>
      updateAnamnesisTemplateItems(template.id, { items }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['anamnesis-template', template.id] })
      toast.success('Items atualizados com sucesso!')
    },
    onError: (error: any) => {
      console.error('Erro ao atualizar items:', error)
      toast.error(error?.message || 'Erro ao atualizar items do template')
    },
  })

  // Mutation para atualizar info
  const updateInfoMutation = useMutation({
    mutationFn: () => updateAnamnesisTemplate(template.id, { name, area }),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['anamnesis-templates'] })
      queryClient.invalidateQueries({ queryKey: ['anamnesis-template', template.id] })
      onClose()
      toast.success('Template atualizado com sucesso!')
    },
    onError: (error: any) => {
      toast.error(error.response?.data?.error || 'Erro ao atualizar template')
    },
  })

  const handleSave = async () => {
    try {
      // 1. Atualizar items primeiro
      const items = selectedScoreItems.map((item, index) => ({
        scoreItemId: item.id,
        order: index,
      }))
      await updateItemsMutation.mutateAsync(items)

      // 2. Depois atualizar info
      updateInfoMutation.mutate()
    } catch (error) {
      toast.error('Erro ao salvar template')
    }
  }

  if (isLoading) {
    return (
      <Dialog open onOpenChange={onClose}>
        <DialogContent className="max-w-6xl max-h-[90vh]">
          <DialogHeader>
            <DialogTitle>Carregando Template</DialogTitle>
            <DialogDescription>Aguarde enquanto carregamos os dados...</DialogDescription>
          </DialogHeader>
          <div className="flex items-center justify-center h-96">
            <Loader2 className="h-8 w-8 animate-spin text-primary" />
          </div>
        </DialogContent>
      </Dialog>
    )
  }

  return (
    <Dialog open onOpenChange={onClose}>
      <DialogContent className="w-[90vw] max-w-none max-h-[90vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle>Editar Template</DialogTitle>
          <DialogDescription>
            Altere o nome, área e score items do template
          </DialogDescription>
        </DialogHeader>

        <div className="space-y-4 py-4">
          <div className="grid grid-cols-2 gap-4">
            <div className="space-y-2">
              <label className="text-sm font-medium">Nome do Template</label>
              <Input
                value={name}
                onChange={(e) => setName(e.target.value)}
                placeholder="Ex: Anamnese Clínica Geral"
              />
            </div>
            <div className="space-y-2">
              <label className="text-sm font-medium">Área</label>
              <Select value={area} onValueChange={(v) => setArea(v as AnamnesisTemplateArea)}>
                <SelectTrigger>
                  <SelectValue />
                </SelectTrigger>
                <SelectContent>
                  {AREAS.map((a) => (
                    <SelectItem key={a} value={a}>
                      {a}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>
          </div>

          <div className="space-y-2">
            <label className="text-sm font-medium">Score Items</label>
            <AnamnesisTemplateItemSelector
              availableScoreItems={availableScoreItems}
              selectedScoreItems={selectedScoreItems}
              onSelectionChange={setSelectedScoreItems}
            />
          </div>
        </div>

        <DialogFooter>
          <Button type="button" variant="outline" onClick={onClose}>
            Cancelar
          </Button>
          <Button
            onClick={handleSave}
            disabled={updateItemsMutation.isPending || updateInfoMutation.isPending}
          >
            {(updateItemsMutation.isPending || updateInfoMutation.isPending) && (
              <Loader2 className="mr-2 h-4 w-4 animate-spin" />
            )}
            Salvar Alterações
          </Button>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  )
}
