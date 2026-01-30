'use client'

import { useState, useEffect } from 'react'
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Textarea } from '@/components/ui/textarea'
import { Card } from '@/components/ui/card'
import { Plus, Pencil, Trash2, Save, X } from 'lucide-react'
import { DualListSelector } from '@/components/lab-tests/dual-list-selector'
import {
  getAllLabRequestTemplates,
  getRequestableLabTests,
  createLabRequestTemplate,
  updateLabRequestTemplate,
  updateLabRequestTemplateTests,
  deleteLabRequestTemplate,
  getLabRequestTemplateById,
  type LabRequestTemplate,
  type LabTestDefinition
} from '@/lib/api/lab-request-templates'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
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
import { toast } from 'sonner'
import { PageHeader } from '@/components/layout/page-header'

export default function LabRequestTemplatesPage() {
  const [isCreateDialogOpen, setIsCreateDialogOpen] = useState(false)
  const [editingTemplate, setEditingTemplate] = useState<LabRequestTemplate | null>(null)
  const [deleteConfirm, setDeleteConfirm] = useState<LabRequestTemplate | null>(null)

  const queryClient = useQueryClient()

  // Fetch templates
  const { data: templates = [], isLoading: templatesLoading } = useQuery({
    queryKey: ['lab-request-templates'],
    queryFn: () => getAllLabRequestTemplates(false)
  })

  // Fetch available tests
  const { data: availableTests = [] } = useQuery({
    queryKey: ['requestable-lab-tests'],
    queryFn: getRequestableLabTests
  })

  // Create template mutation
  const createMutation = useMutation({
    mutationFn: createLabRequestTemplate,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['lab-request-templates'] })
      setIsCreateDialogOpen(false)
      toast.success('Template criado com sucesso!')
    },
    onError: (error: any) => {
      toast.error(error.response?.data?.error || 'Erro ao criar template')
    }
  })

  // Delete template mutation
  const deleteMutation = useMutation({
    mutationFn: deleteLabRequestTemplate,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['lab-request-templates'] })
      setDeleteConfirm(null)
      toast.success('Template excluído com sucesso!')
    },
    onError: (error: any) => {
      toast.error(error.response?.data?.error || 'Erro ao excluir template')
    }
  })

  const handleCreate = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    const formData = new FormData(e.currentTarget)
    const name = formData.get('name') as string
    const description = formData.get('description') as string

    createMutation.mutate({
      name,
      description: description || undefined
    })
  }

  return (
    <div className="container mx-auto py-8 space-y-8">
      <PageHeader
        breadcrumbs={[{ label: 'Templates de Pedidos' }]}
        title="Templates de Pedidos"
        description={`${templates.length} template${templates.length !== 1 ? 's' : ''} configurado${templates.length !== 1 ? 's' : ''}`}
        actions={[
          {
            label: 'Novo',
            icon: <Plus className="h-4 w-4" />,
            onClick: () => setIsCreateDialogOpen(true),
            variant: 'default',
          },
        ]}
      />

      <Dialog open={isCreateDialogOpen} onOpenChange={setIsCreateDialogOpen}>
        <DialogContent>
          <DialogHeader>
            <DialogTitle>Criar Novo Template</DialogTitle>
            <DialogDescription>
              Crie um template e depois adicione exames a ele
            </DialogDescription>
          </DialogHeader>
          <form onSubmit={handleCreate} className="space-y-4">
            <div>
              <label className="text-sm font-medium">Nome</label>
              <Input
                name="name"
                placeholder="Ex: Check-up Anual, Perfil Tireoidiano"
                required
              />
            </div>
            <div>
              <label className="text-sm font-medium">Descrição (opcional)</label>
              <Textarea
                name="description"
                placeholder="Descrição do template..."
                rows={3}
              />
            </div>
            <div className="flex justify-end gap-2">
              <Button
                type="button"
                variant="outline"
                onClick={() => setIsCreateDialogOpen(false)}
              >
                Cancelar
              </Button>
              <Button type="submit" disabled={createMutation.isPending}>
                {createMutation.isPending ? 'Criando...' : 'Criar Template'}
              </Button>
            </div>
          </form>
        </DialogContent>
      </Dialog>

      {templatesLoading ? (
        <div className="text-center py-12">Carregando templates...</div>
      ) : templates.length === 0 ? (
        <Card className="p-12 text-center">
          <p className="text-muted-foreground mb-4">Nenhum template cadastrado ainda</p>
          <Button onClick={() => setIsCreateDialogOpen(true)}>
            <Plus className="h-4 w-4 mr-2" />
            Criar Primeiro Template
          </Button>
        </Card>
      ) : (
        <div className="grid gap-4">
          {templates.map((template) => (
            <TemplateCard
              key={template.id}
              template={template}
              availableTests={availableTests}
              onEdit={() => setEditingTemplate(template)}
              onDelete={() => setDeleteConfirm(template)}
            />
          ))}
        </div>
      )}

      {/* Edit Dialog */}
      {editingTemplate && (
        <TemplateEditDialog
          template={editingTemplate}
          availableTests={availableTests}
          onClose={() => setEditingTemplate(null)}
        />
      )}

      {/* Delete Confirmation */}
      <AlertDialog open={!!deleteConfirm} onOpenChange={() => setDeleteConfirm(null)}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Confirmar Exclusão</AlertDialogTitle>
            <AlertDialogDescription>
              Tem certeza que deseja excluir o template "{deleteConfirm?.name}"? Esta ação não pode ser desfeita.
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>Cancelar</AlertDialogCancel>
            <AlertDialogAction
              onClick={() => deleteConfirm && deleteMutation.mutate(deleteConfirm.id)}
              className="bg-destructive hover:bg-destructive/90"
            >
              Excluir
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </div>
  )
}

function TemplateCard({
  template,
  availableTests,
  onEdit,
  onDelete
}: {
  template: LabRequestTemplate
  availableTests: LabTestDefinition[]
  onEdit: () => void
  onDelete: () => void
}) {
  // Fetch template details with tests
  const { data: templateWithTests } = useQuery({
    queryKey: ['lab-request-template', template.id],
    queryFn: () => getLabRequestTemplateById(template.id)
  })

  const testsCount = templateWithTests?.labTests?.length || 0

  return (
    <Card className="p-4">
      <div className="flex items-start justify-between">
        <div className="flex-1">
          <h3 className="font-semibold text-lg">{template.name}</h3>
          {template.description && (
            <p className="text-sm text-muted-foreground mt-1">{template.description}</p>
          )}
          <p className="text-sm text-muted-foreground mt-2">
            {testsCount} exame{testsCount !== 1 ? 's' : ''} configurado{testsCount !== 1 ? 's' : ''}
          </p>
        </div>
        <div className="flex gap-2">
          <Button variant="outline" size="sm" onClick={onEdit}>
            <Pencil className="h-4 w-4" />
          </Button>
          <Button variant="outline" size="sm" onClick={onDelete}>
            <Trash2 className="h-4 w-4 text-destructive" />
          </Button>
        </div>
      </div>
    </Card>
  )
}

function TemplateEditDialog({
  template,
  availableTests,
  onClose
}: {
  template: LabRequestTemplate
  availableTests: LabTestDefinition[]
  onClose: () => void
}) {
  const queryClient = useQueryClient()
  const [name, setName] = useState(template.name)
  const [description, setDescription] = useState(template.description || '')
  const [selectedTests, setSelectedTests] = useState<LabTestDefinition[]>([])

  // Fetch template with tests
  const { data: templateWithTests, isLoading } = useQuery({
    queryKey: ['lab-request-template', template.id],
    queryFn: () => getLabRequestTemplateById(template.id)
  })

  // Update selected tests when template data loads
  useEffect(() => {
    if (templateWithTests?.labTests) {
      setSelectedTests(templateWithTests.labTests)
    }
  }, [templateWithTests])

  // Update template info
  const updateInfoMutation = useMutation({
    mutationFn: (data: { name: string; description?: string }) =>
      updateLabRequestTemplate(template.id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['lab-request-templates'] })
      queryClient.invalidateQueries({ queryKey: ['lab-request-template', template.id] })
      toast.success('Template atualizado!')
    }
  })

  // Update tests
  const updateTestsMutation = useMutation({
    mutationFn: (testIds: string[]) =>
      updateLabRequestTemplateTests(template.id, testIds),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['lab-request-template', template.id] })
      toast.success('Exames atualizados!')
    }
  })

  const handleSave = () => {
    // Update basic info
    updateInfoMutation.mutate({
      name,
      description: description || undefined
    })

    // Update tests
    const testIds = selectedTests.map(t => t.id)
    updateTestsMutation.mutate(testIds)
  }

  return (
    <Dialog open={true} onOpenChange={onClose}>
      <DialogContent className="max-w-6xl max-h-[90vh] overflow-y-auto">
        <DialogHeader>
          <DialogTitle>Editar Template</DialogTitle>
          <DialogDescription>
            Configure o nome, descrição e exames do template
          </DialogDescription>
        </DialogHeader>

        {isLoading ? (
          <div className="py-12 text-center">Carregando...</div>
        ) : (
          <div className="space-y-6">
            <div className="grid gap-4">
              <div>
                <label className="text-sm font-medium">Nome</label>
                <Input
                  value={name}
                  onChange={(e) => setName(e.target.value)}
                  placeholder="Ex: Check-up Anual"
                />
              </div>
              <div>
                <label className="text-sm font-medium">Descrição</label>
                <Textarea
                  value={description}
                  onChange={(e) => setDescription(e.target.value)}
                  placeholder="Descrição do template..."
                  rows={2}
                />
              </div>
            </div>

            <div>
              <h4 className="font-medium mb-4">Selecionar Exames</h4>
              <DualListSelector
                availableTests={availableTests}
                selectedTests={selectedTests}
                onSelectionChange={setSelectedTests}
              />
            </div>

            <div className="flex justify-end gap-2 pt-4 border-t">
              <Button variant="outline" onClick={onClose}>
                <X className="h-4 w-4 mr-2" />
                Cancelar
              </Button>
              <Button
                onClick={handleSave}
                disabled={updateInfoMutation.isPending || updateTestsMutation.isPending}
              >
                <Save className="h-4 w-4 mr-2" />
                {updateInfoMutation.isPending || updateTestsMutation.isPending ? 'Salvando...' : 'Salvar'}
              </Button>
            </div>
          </div>
        )}
      </DialogContent>
    </Dialog>
  )
}
