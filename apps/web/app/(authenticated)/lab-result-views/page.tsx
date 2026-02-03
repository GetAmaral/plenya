'use client'

import { useState } from 'react'
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { useRouter } from 'next/navigation'
import { Plus, Pencil, Trash2, Search } from 'lucide-react'
import { labResultViewApi } from '@/lib/api/lab-result-view-api'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
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
import { Badge } from '@/components/ui/badge'
import { PageHeader } from '@/components/layout/page-header'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { useToast } from '@/hooks/use-toast'

export default function LabResultViewsPage() {
  const router = useRouter()
  const queryClient = useQueryClient()
  const { toast } = useToast()
  const [searchQuery, setSearchQuery] = useState('')
  const [deleteViewId, setDeleteViewId] = useState<string | null>(null)

  // Buscar views
  const { data: views, isLoading } = useQuery({
    queryKey: ['lab-result-views'],
    queryFn: () => labResultViewApi.list(false, true),
  })

  // Deletar view
  const deleteMutation = useMutation({
    mutationFn: labResultViewApi.delete,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['lab-result-views'] })
      toast({
        title: 'View deletada',
        description: 'A view foi deletada com sucesso.',
      })
      setDeleteViewId(null)
    },
    onError: () => {
      toast({
        title: 'Erro',
        description: 'Não foi possível deletar a view.',
        variant: 'destructive',
      })
    },
  })

  // Filtrar views
  const filteredViews = views?.filter((view) => {
    if (!searchQuery) return true
    const query = searchQuery.toLowerCase()
    return (
      view.name.toLowerCase().includes(query) ||
      view.description?.toLowerCase().includes(query)
    )
  })

  return (
    <>
      <PageHeader
        title="Views de Resultados Laboratoriais"
        description="Gerencie views customizadas para ordenação de exames na tabela pivot."
      >
        <Button onClick={() => router.push('/lab-result-views/new/edit')}>
          <Plus className="mr-2 h-4 w-4" />
          Nova View
        </Button>
      </PageHeader>

      <Card>
        <CardHeader>
          <CardTitle>Views</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="mb-4">
            <div className="relative">
              <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-muted-foreground" />
              <Input
                placeholder="Buscar por nome ou descrição..."
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
                className="pl-10"
              />
            </div>
          </div>

          {isLoading ? (
            <div className="py-8 text-center text-muted-foreground">
              Carregando views...
            </div>
          ) : filteredViews && filteredViews.length > 0 ? (
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead>Nome</TableHead>
                  <TableHead>Descrição</TableHead>
                  <TableHead>Status</TableHead>
                  <TableHead className="text-right">Ordem</TableHead>
                  <TableHead className="text-right">Ações</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {filteredViews.map((view) => (
                  <TableRow key={view.id}>
                    <TableCell className="font-medium">{view.name}</TableCell>
                    <TableCell className="text-muted-foreground">
                      {view.description || '-'}
                    </TableCell>
                    <TableCell>
                      <Badge variant={view.isActive ? 'default' : 'secondary'}>
                        {view.isActive ? 'Ativa' : 'Inativa'}
                      </Badge>
                    </TableCell>
                    <TableCell className="text-right">
                      {view.displayOrder}
                    </TableCell>
                    <TableCell className="text-right">
                      <div className="flex justify-end gap-2">
                        <Button
                          variant="ghost"
                          size="sm"
                          onClick={() =>
                            router.push(`/lab-result-views/${view.id}/edit`)
                          }
                        >
                          <Pencil className="h-4 w-4" />
                        </Button>
                        <Button
                          variant="ghost"
                          size="sm"
                          onClick={() => setDeleteViewId(view.id)}
                        >
                          <Trash2 className="h-4 w-4 text-destructive" />
                        </Button>
                      </div>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          ) : (
            <div className="py-8 text-center text-muted-foreground">
              {searchQuery
                ? 'Nenhuma view encontrada.'
                : 'Nenhuma view cadastrada.'}
            </div>
          )}
        </CardContent>
      </Card>

      {/* Dialog de confirmação de exclusão */}
      <AlertDialog
        open={!!deleteViewId}
        onOpenChange={() => setDeleteViewId(null)}
      >
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Confirmar exclusão</AlertDialogTitle>
            <AlertDialogDescription>
              Tem certeza que deseja deletar esta view? Esta ação não pode ser
              desfeita.
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>Cancelar</AlertDialogCancel>
            <AlertDialogAction
              onClick={() => deleteViewId && deleteMutation.mutate(deleteViewId)}
            >
              Deletar
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </>
  )
}
