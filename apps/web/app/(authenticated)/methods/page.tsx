'use client'

import { useState } from 'react'
import { useAllMethods, useDeleteMethod } from '@/lib/api/method-api'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { Skeleton } from '@/components/ui/skeleton'
import { AlertCircle, Plus, Edit, Trash2 } from 'lucide-react'
import Link from 'next/link'
import { MethodDialog } from '@/components/methods/MethodDialog'
import { DeleteConfirmDialog } from '@/components/methods/DeleteConfirmDialog'
import { useAuth } from '@/lib/use-auth'
import { toast } from 'sonner'
import type { Method } from '@plenya/types'

export default function MethodsListPage() {
  const { data: methods, isLoading, error } = useAllMethods()
  const { user } = useAuth()
  const deleteMethod = useDeleteMethod()

  const [createDialogOpen, setCreateDialogOpen] = useState(false)
  const [editDialogOpen, setEditDialogOpen] = useState(false)
  const [editingMethod, setEditingMethod] = useState<Method | null>(null)
  const [deleteConfirmOpen, setDeleteConfirmOpen] = useState(false)
  const [deletingMethod, setDeletingMethod] = useState<Method | null>(null)

  const isAdmin = user?.roles?.includes('admin')

  const handleEdit = (method: Method, e: React.MouseEvent) => {
    e.preventDefault()
    e.stopPropagation()
    setEditingMethod(method)
    setEditDialogOpen(true)
  }

  const handleDelete = (method: Method, e: React.MouseEvent) => {
    e.preventDefault()
    e.stopPropagation()
    setDeletingMethod(method)
    setDeleteConfirmOpen(true)
  }

  const confirmDelete = async () => {
    if (!deletingMethod) return

    try {
      await deleteMethod.mutateAsync(deletingMethod.id)
      toast.success('Metodologia deletada com sucesso')
      setDeleteConfirmOpen(false)
      setDeletingMethod(null)
    } catch (error: any) {
      toast.error(error?.message || 'Erro ao deletar metodologia')
    }
  }

  if (isLoading) {
    return (
      <div className="space-y-6">
        <div className="space-y-2">
          <Skeleton className="h-10 w-[300px]" />
          <Skeleton className="h-5 w-[500px]" />
        </div>
        <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
          <Skeleton className="h-[200px]" />
          <Skeleton className="h-[200px]" />
          <Skeleton className="h-[200px]" />
        </div>
      </div>
    )
  }

  if (error) {
    return (
      <div className="space-y-6">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Metodologias Clínicas</h1>
          <p className="text-muted-foreground">
            Organize itens de escore em metodologias clínicas estruturadas
          </p>
        </div>
        <Card className="border-destructive">
          <CardContent className="flex items-center gap-3 p-6">
            <AlertCircle className="h-5 w-5 text-destructive" />
            <div>
              <p className="font-medium">Erro ao carregar metodologias</p>
              <p className="text-sm text-muted-foreground">
                Ocorreu um erro ao carregar as metodologias. Tente novamente mais tarde.
              </p>
            </div>
          </CardContent>
        </Card>
      </div>
    )
  }

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex items-start justify-between">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Metodologias Clínicas</h1>
          <p className="text-muted-foreground">
            Organize itens de escore em metodologias clínicas estruturadas
          </p>
        </div>
        {isAdmin && (
          <Button onClick={() => setCreateDialogOpen(true)}>
            <Plus className="h-4 w-4 mr-1" />
            Nova Metodologia
          </Button>
        )}
      </div>

      {/* Methods Grid */}
      {methods && methods.length > 0 ? (
        <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
          {methods.map((method) => (
            <Card key={method.id} className="h-full transition-all hover:shadow-lg cursor-pointer group">
              <Link href={`/methods/${method.id}`}>
                <CardHeader>
                  <div className="flex items-start justify-between gap-2">
                    <div className="flex-1">
                      <div className="flex items-start justify-between gap-2">
                        <CardTitle style={{ color: method.color || undefined }}>
                          {method.shortName}
                        </CardTitle>
                        <div className="flex gap-1.5">
                          {method.isDefault && (
                            <Badge variant="default" className="bg-primary">
                              Padrão
                            </Badge>
                          )}
                          <Badge variant="outline">v{method.version || '1.0'}</Badge>
                        </div>
                      </div>
                      <CardDescription className="text-base font-medium mt-1">
                        {method.name}
                      </CardDescription>
                    </div>
                  </div>
                  {isAdmin && (
                    <div className="flex gap-1 mt-2 opacity-0 group-hover:opacity-100 transition-opacity">
                      <Button
                        variant="outline"
                        size="sm"
                        onClick={(e) => handleEdit(method, e)}
                        className="h-8 flex-1"
                      >
                        <Edit className="h-3 w-3 mr-1" />
                        Editar
                      </Button>
                      <Button
                        variant="outline"
                        size="sm"
                        onClick={(e) => handleDelete(method, e)}
                        className="h-8 flex-1 hover:bg-destructive hover:text-destructive-foreground"
                      >
                        <Trash2 className="h-3 w-3 mr-1" />
                        Deletar
                      </Button>
                    </div>
                  )}
                </CardHeader>
                <CardContent>
                  {method.description && (
                    <p className="text-sm text-muted-foreground line-clamp-3">
                      {method.description}
                    </p>
                  )}
                </CardContent>
              </Link>
            </Card>
          ))}
        </div>
      ) : (
        <Card>
          <CardContent className="flex flex-col items-center justify-center py-12">
            <p className="text-muted-foreground text-center">
              Nenhuma metodologia cadastrada ainda.
            </p>
            <p className="text-sm text-muted-foreground text-center mt-2">
              Entre em contato com o administrador para criar metodologias.
            </p>
          </CardContent>
        </Card>
      )}

      {/* Create Method Dialog */}
      <MethodDialog
        open={createDialogOpen}
        onOpenChange={setCreateDialogOpen}
      />

      {/* Edit Method Dialog */}
      {editingMethod && (
        <MethodDialog
          open={editDialogOpen}
          onOpenChange={(open) => {
            setEditDialogOpen(open)
            if (!open) setEditingMethod(null)
          }}
          method={editingMethod}
        />
      )}

      {/* Delete Confirmation Dialog */}
      <DeleteConfirmDialog
        open={deleteConfirmOpen}
        onOpenChange={setDeleteConfirmOpen}
        onConfirm={confirmDelete}
        title="Deletar Metodologia"
        description={`Tem certeza que deseja deletar "${deletingMethod?.name}"? Todas as letras, pilares e associações serão deletados. Esta ação não pode ser desfeita.`}
        isDeleting={deleteMethod.isPending}
      />
    </div>
  )
}
