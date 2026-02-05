'use client'

import { use, useEffect, useRef, useState } from 'react'
import { useRouter } from 'next/navigation'
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { z } from 'zod'
import { labResultViewApi } from '@/lib/api/lab-result-view-api'
import { getRequestableLabTests } from '@/lib/api/lab-request-templates'
import { Button } from '@/components/ui/button'
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { Textarea } from '@/components/ui/textarea'
import { Switch } from '@/components/ui/switch'
import { PageHeader } from '@/components/layout/page-header'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { useToast } from '@/hooks/use-toast'
import { useFormNavigation } from '@/lib/use-form-navigation'
import { LabResultViewItemSelector, type LabTestDefinition } from '@/components/lab-results/LabResultViewItemSelector'

const formSchema = z.object({
  name: z.string().min(2, 'Nome deve ter no mínimo 2 caracteres').max(200),
  description: z.string().optional(),
  isActive: z.boolean(),
  displayOrder: z.number().min(0).max(9999),
})

type FormValues = z.infer<typeof formSchema>

export default function LabResultViewEditPage({
  params,
}: {
  params: Promise<{ id: string }>
}) {
  const resolvedParams = use(params)
  const viewId = resolvedParams.id
  const isNew = viewId === 'new'
  const router = useRouter()
  const queryClient = useQueryClient()
  const { toast } = useToast()
  const formRef = useRef<HTMLFormElement>(null)

  const [selectedTests, setSelectedTests] = useState<LabTestDefinition[]>([])

  // Hook de navegação de formulário
  useFormNavigation({ formRef })

  const form = useForm<FormValues>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: '',
      description: '',
      isActive: true,
      displayOrder: 0,
    },
  })

  // Buscar view existente
  const { data: view, isLoading: isLoadingView } = useQuery({
    queryKey: ['lab-result-view', viewId],
    queryFn: () => labResultViewApi.getById(viewId),
    enabled: !isNew,
  })

  // Buscar todos os testes disponíveis
  const { data: allTests = [], isLoading: isLoadingTests } = useQuery({
    queryKey: ['requestable-lab-tests'],
    queryFn: async () => {
      const tests = await getRequestableLabTests()
      return tests.map((test) => ({
        id: test.id,
        name: test.name,
        code: test.code,
        category: test.category || 'Sem Categoria',
      }))
    },
  })

  // Preencher formulário quando view for carregada
  useEffect(() => {
    if (view) {
      form.reset({
        name: view.name,
        description: view.description || '',
        isActive: view.isActive,
        displayOrder: view.displayOrder,
      })

      // Preencher testes selecionados (com order preservado)
      if (view.items) {
        const tests = view.items
          .sort((a, b) => a.order - b.order)
          .map((item) => ({
            id: item.labTestDefinition?.id || item.labTestDefinitionId,
            name: item.labTestDefinition?.name || '',
            code: item.labTestDefinition?.code || '',
            category: (item.labTestDefinition as any)?.category || 'Sem Categoria',
          }))
          .filter((test) => test.name) // Remover testes sem dados

        setSelectedTests(tests)
      }
    }
  }, [view, form])

  // Criar view
  const createMutation = useMutation({
    mutationFn: async (data: FormValues) => {
      const view = await labResultViewApi.create({
        name: data.name,
        description: data.description || undefined,
        displayOrder: data.displayOrder,
      })

      // Atualizar items
      if (selectedTests.length > 0) {
        await labResultViewApi.updateItems(view.id, {
          items: selectedTests.map((test, index) => ({
            labTestDefinitionId: test.id,
            order: index,
          })),
        })
      }

      return view
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['lab-result-views'] })
      toast({
        title: 'View criada',
        description: 'A view foi criada com sucesso.',
      })
      router.push('/lab-result-views')
    },
    onError: () => {
      toast({
        title: 'Erro',
        description: 'Não foi possível criar a view.',
        variant: 'destructive',
      })
    },
  })

  // Atualizar view
  const updateMutation = useMutation({
    mutationFn: async (data: FormValues) => {
      await labResultViewApi.update(viewId, {
        name: data.name,
        description: data.description || undefined,
        isActive: data.isActive,
        displayOrder: data.displayOrder,
      })

      // Atualizar items
      await labResultViewApi.updateItems(viewId, {
        items: selectedTests.map((test, index) => ({
          labTestDefinitionId: test.id,
          order: index,
        })),
      })
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['lab-result-views'] })
      queryClient.invalidateQueries({ queryKey: ['lab-result-view', viewId] })
      toast({
        title: 'View atualizada',
        description: 'A view foi atualizada com sucesso.',
      })
      router.push('/lab-result-views')
    },
    onError: () => {
      toast({
        title: 'Erro',
        description: 'Não foi possível atualizar a view.',
        variant: 'destructive',
      })
    },
  })

  const onSubmit = (data: FormValues) => {
    if (isNew) {
      createMutation.mutate(data)
    } else {
      updateMutation.mutate(data)
    }
  }

  const isLoading = isLoadingView || isLoadingTests
  const isSaving = createMutation.isPending || updateMutation.isPending

  if (isLoading) {
    return (
      <>
        <PageHeader
          title={isNew ? 'Nova View' : 'Editar View'}
          description="Carregando..."
        />
      </>
    )
  }

  return (
    <>
      <PageHeader
        title={isNew ? 'Nova View' : 'Editar View'}
        description={
          isNew
            ? 'Crie uma view customizada para ordenar exames na tabela pivot.'
            : 'Edite a view de ordenação de exames.'
        }
      />

      <Form {...form}>
        <form
          ref={formRef}
          onSubmit={form.handleSubmit(onSubmit)}
          className="space-y-6"
        >
          <Card>
            <CardHeader>
              <CardTitle>Informações Básicas</CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <FormField
                control={form.control}
                name="name"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Nome</FormLabel>
                    <FormControl>
                      <Input placeholder="Ex: View padrão" {...field} />
                    </FormControl>
                    <FormDescription>
                      Nome descritivo da view.
                    </FormDescription>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="description"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Descrição (Opcional)</FormLabel>
                    <FormControl>
                      <Textarea
                        placeholder="Descrição da view..."
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <div className="grid grid-cols-2 gap-4">
                <FormField
                  control={form.control}
                  name="isActive"
                  render={({ field }) => (
                    <FormItem className="flex items-center justify-between rounded-lg border p-4">
                      <div className="space-y-0.5">
                        <FormLabel>Ativa</FormLabel>
                        <FormDescription>
                          View disponível para uso.
                        </FormDescription>
                      </div>
                      <FormControl>
                        <Switch
                          checked={field.value}
                          onCheckedChange={field.onChange}
                        />
                      </FormControl>
                    </FormItem>
                  )}
                />

                <FormField
                  control={form.control}
                  name="displayOrder"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>Ordem de Exibição</FormLabel>
                      <FormControl>
                        <Input
                          type="number"
                          min={0}
                          max={9999}
                          {...field}
                          onChange={(e) =>
                            field.onChange(parseInt(e.target.value) || 0)
                          }
                        />
                      </FormControl>
                      <FormDescription>
                        Ordem na lista de views.
                      </FormDescription>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>Exames</CardTitle>
            </CardHeader>
            <CardContent>
              <LabResultViewItemSelector
                availableTests={allTests}
                selectedTests={selectedTests}
                onSelectionChange={setSelectedTests}
              />
            </CardContent>
          </Card>

          <div className="flex justify-end gap-2">
            <Button
              type="button"
              variant="outline"
              onClick={() => router.push('/lab-result-views')}
              disabled={isSaving}
            >
              Cancelar
            </Button>
            <Button type="submit" disabled={isSaving}>
              {isSaving ? 'Salvando...' : 'Salvar'}
            </Button>
          </div>
        </form>
      </Form>
    </>
  )
}
