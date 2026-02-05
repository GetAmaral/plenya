'use client'

import { useState } from 'react'
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { Plus, Pencil, Trash2 } from 'lucide-react'
import { useRouter } from 'next/navigation'
import { toast } from 'sonner'

import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Input } from '@/components/ui/input'
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@/components/ui/table'
import { Badge } from '@/components/ui/badge'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
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
  listMedicationDefinitions,
  deleteMedicationDefinition,
  type MedicationDefinition,
} from '@/lib/api/medication-definitions'

export default function MedicationDefinitionsPage() {
  const router = useRouter()
  const queryClient = useQueryClient()

  const [searchQuery, setSearchQuery] = useState('')
  const [categoryFilter, setCategoryFilter] = useState<string>('all')
  const [deleteId, setDeleteId] = useState<string | null>(null)

  const { data, isLoading } = useQuery({
    queryKey: ['medication-definitions', categoryFilter],
    queryFn: () =>
      listMedicationDefinitions({
        category: categoryFilter === 'all' ? undefined : categoryFilter,
        limit: 100,
      }),
  })

  const deleteMutation = useMutation({
    mutationFn: deleteMedicationDefinition,
    onSuccess: () => {
      toast.success('Medicamento deletado com sucesso')
      queryClient.invalidateQueries({ queryKey: ['medication-definitions'] })
      setDeleteId(null)
    },
    onError: () => {
      toast.error('Erro ao deletar medicamento')
    },
  })

  const filteredMedications =
    data?.data.filter((med) =>
      searchQuery
        ? med.commonName.toLowerCase().includes(searchQuery.toLowerCase()) ||
          med.activeIngredient.toLowerCase().includes(searchQuery.toLowerCase())
        : true
    ) || []

  const getCategoryLabel = (category: string) => {
    const labels: Record<string, string> = {
      simple: 'Simples',
      c1: 'C1 (Controle Especial)',
      c5: 'C5 (Psicotrópico)',
      antibiotic: 'Antibiótico',
      glp1: 'GLP-1',
    }
    return labels[category] || category
  }

  const getCategoryVariant = (
    category: string
  ): 'default' | 'secondary' | 'destructive' | 'outline' => {
    const variants: Record<string, 'default' | 'secondary' | 'destructive' | 'outline'> = {
      simple: 'outline',
      c1: 'default',
      c5: 'destructive',
      antibiotic: 'secondary',
      glp1: 'default',
    }
    return variants[category] || 'outline'
  }

  return (
    <div className="container mx-auto py-8">
      <Card>
        <CardHeader className="flex flex-row items-center justify-between">
          <CardTitle>Definições de Medicamentos</CardTitle>
          <Button onClick={() => router.push('/admin/medication-definitions/new')}>
            <Plus className="mr-2 h-4 w-4" />
            Novo Medicamento
          </Button>
        </CardHeader>

        <CardContent className="space-y-4">
          <div className="flex gap-4">
            <div className="flex-1">
              <Input
                placeholder="Buscar por nome ou princípio ativo..."
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
              />
            </div>

            <Select value={categoryFilter} onValueChange={setCategoryFilter}>
              <SelectTrigger className="w-[200px]">
                <SelectValue placeholder="Categoria" />
              </SelectTrigger>
              <SelectContent>
                <SelectItem value="all">Todas as categorias</SelectItem>
                <SelectItem value="simple">Simples</SelectItem>
                <SelectItem value="c1">C1 (Controle Especial)</SelectItem>
                <SelectItem value="c5">C5 (Psicotrópico)</SelectItem>
                <SelectItem value="antibiotic">Antibiótico</SelectItem>
                <SelectItem value="glp1">GLP-1</SelectItem>
              </SelectContent>
            </Select>
          </div>

          {isLoading ? (
            <div className="text-center py-8 text-muted-foreground">Carregando...</div>
          ) : filteredMedications.length === 0 ? (
            <div className="text-center py-8 text-muted-foreground">
              Nenhum medicamento encontrado
            </div>
          ) : (
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead>Nome Comercial</TableHead>
                  <TableHead>Princípio Ativo</TableHead>
                  <TableHead>Categoria</TableHead>
                  <TableHead className="text-center">Validade</TableHead>
                  <TableHead className="text-center">SNCR</TableHead>
                  <TableHead className="text-center">Assinatura</TableHead>
                  <TableHead className="text-right">Ações</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {filteredMedications.map((medication) => (
                  <TableRow key={medication.id}>
                    <TableCell className="font-medium">{medication.commonName}</TableCell>
                    <TableCell className="text-muted-foreground">
                      {medication.activeIngredient}
                    </TableCell>
                    <TableCell>
                      <Badge variant={getCategoryVariant(medication.category)}>
                        {getCategoryLabel(medication.category)}
                      </Badge>
                    </TableCell>
                    <TableCell className="text-center">
                      {medication.validityDays} dias
                    </TableCell>
                    <TableCell className="text-center">
                      {medication.requiresSNCR ? (
                        <Badge variant="outline" className="bg-yellow-50 text-yellow-700">
                          Sim
                        </Badge>
                      ) : (
                        <span className="text-muted-foreground">-</span>
                      )}
                    </TableCell>
                    <TableCell className="text-center">
                      {medication.requiresDigitalSignature ? (
                        <Badge variant="outline" className="bg-blue-50 text-blue-700">
                          Sim
                        </Badge>
                      ) : (
                        <span className="text-muted-foreground">-</span>
                      )}
                    </TableCell>
                    <TableCell className="text-right">
                      <div className="flex justify-end gap-2">
                        <Button
                          variant="ghost"
                          size="sm"
                          onClick={() =>
                            router.push(`/admin/medication-definitions/${medication.id}/edit`)
                          }
                        >
                          <Pencil className="h-4 w-4" />
                        </Button>
                        <Button
                          variant="ghost"
                          size="sm"
                          onClick={() => setDeleteId(medication.id)}
                        >
                          <Trash2 className="h-4 w-4 text-destructive" />
                        </Button>
                      </div>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          )}
        </CardContent>
      </Card>

      <AlertDialog open={!!deleteId} onOpenChange={() => setDeleteId(null)}>
        <AlertDialogContent>
          <AlertDialogHeader>
            <AlertDialogTitle>Confirmar exclusão</AlertDialogTitle>
            <AlertDialogDescription>
              Tem certeza que deseja deletar este medicamento? Esta ação não pode ser
              desfeita.
            </AlertDialogDescription>
          </AlertDialogHeader>
          <AlertDialogFooter>
            <AlertDialogCancel>Cancelar</AlertDialogCancel>
            <AlertDialogAction
              onClick={() => deleteId && deleteMutation.mutate(deleteId)}
              className="bg-destructive text-destructive-foreground hover:bg-destructive/90"
            >
              Deletar
            </AlertDialogAction>
          </AlertDialogFooter>
        </AlertDialogContent>
      </AlertDialog>
    </div>
  )
}
